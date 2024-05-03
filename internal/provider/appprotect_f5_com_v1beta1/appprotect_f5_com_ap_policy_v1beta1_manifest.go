/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package appprotect_f5_com_v1beta1

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
	_ datasource.DataSource = &AppprotectF5ComAppolicyV1Beta1Manifest{}
)

func NewAppprotectF5ComAppolicyV1Beta1Manifest() datasource.DataSource {
	return &AppprotectF5ComAppolicyV1Beta1Manifest{}
}

type AppprotectF5ComAppolicyV1Beta1Manifest struct{}

type AppprotectF5ComAppolicyV1Beta1ManifestData struct {
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
		Modifications          *[]map[string]string `tfsdk:"modifications" json:"modifications,omitempty"`
		ModificationsReference *struct {
			Link *string `tfsdk:"link" json:"link,omitempty"`
		} `tfsdk:"modifications_reference" json:"modificationsReference,omitempty"`
		Policy *struct {
			ApplicationLanguage *string `tfsdk:"application_language" json:"applicationLanguage,omitempty"`
			Blocking_settings   *struct {
				Evasions *[]struct {
					Description       *string `tfsdk:"description" json:"description,omitempty"`
					Enabled           *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
					MaxDecodingPasses *int64  `tfsdk:"max_decoding_passes" json:"maxDecodingPasses,omitempty"`
				} `tfsdk:"evasions" json:"evasions,omitempty"`
				Http_protocols *[]struct {
					Description *string `tfsdk:"description" json:"description,omitempty"`
					Enabled     *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
					MaxCookies  *int64  `tfsdk:"max_cookies" json:"maxCookies,omitempty"`
					MaxHeaders  *int64  `tfsdk:"max_headers" json:"maxHeaders,omitempty"`
					MaxParams   *int64  `tfsdk:"max_params" json:"maxParams,omitempty"`
				} `tfsdk:"http_protocols" json:"http-protocols,omitempty"`
				Violations *[]struct {
					Alarm       *bool   `tfsdk:"alarm" json:"alarm,omitempty"`
					Block       *bool   `tfsdk:"block" json:"block,omitempty"`
					Description *string `tfsdk:"description" json:"description,omitempty"`
					Name        *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"violations" json:"violations,omitempty"`
			} `tfsdk:"blocking_settings" json:"blocking-settings,omitempty"`
			BlockingSettingReference *struct {
				Link *string `tfsdk:"link" json:"link,omitempty"`
			} `tfsdk:"blocking_setting_reference" json:"blockingSettingReference,omitempty"`
			Bot_defense *struct {
				Mitigations *struct {
					Anomalies *[]struct {
						Dollaraction   *string `tfsdk:"dollaraction" json:"$action,omitempty"`
						Action         *string `tfsdk:"action" json:"action,omitempty"`
						Name           *string `tfsdk:"name" json:"name,omitempty"`
						ScoreThreshold *string `tfsdk:"score_threshold" json:"scoreThreshold,omitempty"`
					} `tfsdk:"anomalies" json:"anomalies,omitempty"`
					Browsers *[]struct {
						Dollaraction *string `tfsdk:"dollaraction" json:"$action,omitempty"`
						Action       *string `tfsdk:"action" json:"action,omitempty"`
						MaxVersion   *int64  `tfsdk:"max_version" json:"maxVersion,omitempty"`
						MinVersion   *int64  `tfsdk:"min_version" json:"minVersion,omitempty"`
						Name         *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"browsers" json:"browsers,omitempty"`
					Classes *[]struct {
						Action *string `tfsdk:"action" json:"action,omitempty"`
						Name   *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"classes" json:"classes,omitempty"`
					Signatures *[]struct {
						Dollaraction *string `tfsdk:"dollaraction" json:"$action,omitempty"`
						Action       *string `tfsdk:"action" json:"action,omitempty"`
						Name         *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"signatures" json:"signatures,omitempty"`
				} `tfsdk:"mitigations" json:"mitigations,omitempty"`
				Settings *struct {
					CaseSensitiveHttpHeaders *bool `tfsdk:"case_sensitive_http_headers" json:"caseSensitiveHttpHeaders,omitempty"`
					IsEnabled                *bool `tfsdk:"is_enabled" json:"isEnabled,omitempty"`
				} `tfsdk:"settings" json:"settings,omitempty"`
			} `tfsdk:"bot_defense" json:"bot-defense,omitempty"`
			Browser_definitions *[]struct {
				Dollaraction  *string `tfsdk:"dollaraction" json:"$action,omitempty"`
				IsUserDefined *bool   `tfsdk:"is_user_defined" json:"isUserDefined,omitempty"`
				MatchRegex    *string `tfsdk:"match_regex" json:"matchRegex,omitempty"`
				MatchString   *string `tfsdk:"match_string" json:"matchString,omitempty"`
				Name          *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"browser_definitions" json:"browser-definitions,omitempty"`
			CaseInsensitive *bool `tfsdk:"case_insensitive" json:"caseInsensitive,omitempty"`
			Character_sets  *[]struct {
				CharacterSet *[]struct {
					IsAllowed *bool   `tfsdk:"is_allowed" json:"isAllowed,omitempty"`
					Metachar  *string `tfsdk:"metachar" json:"metachar,omitempty"`
				} `tfsdk:"character_set" json:"characterSet,omitempty"`
				CharacterSetType *string `tfsdk:"character_set_type" json:"characterSetType,omitempty"`
			} `tfsdk:"character_sets" json:"character-sets,omitempty"`
			CharacterSetReference *struct {
				Link *string `tfsdk:"link" json:"link,omitempty"`
			} `tfsdk:"character_set_reference" json:"characterSetReference,omitempty"`
			Cookie_settings *struct {
				MaximumCookieHeaderLength *string `tfsdk:"maximum_cookie_header_length" json:"maximumCookieHeaderLength,omitempty"`
			} `tfsdk:"cookie_settings" json:"cookie-settings,omitempty"`
			CookieReference *struct {
				Link *string `tfsdk:"link" json:"link,omitempty"`
			} `tfsdk:"cookie_reference" json:"cookieReference,omitempty"`
			CookieSettingsReference *struct {
				Link *string `tfsdk:"link" json:"link,omitempty"`
			} `tfsdk:"cookie_settings_reference" json:"cookieSettingsReference,omitempty"`
			Cookies *[]struct {
				Dollaraction                         *string `tfsdk:"dollaraction" json:"$action,omitempty"`
				AccessibleOnlyThroughTheHttpProtocol *bool   `tfsdk:"accessible_only_through_the_http_protocol" json:"accessibleOnlyThroughTheHttpProtocol,omitempty"`
				AttackSignaturesCheck                *bool   `tfsdk:"attack_signatures_check" json:"attackSignaturesCheck,omitempty"`
				DecodeValueAsBase64                  *string `tfsdk:"decode_value_as_base64" json:"decodeValueAsBase64,omitempty"`
				EnforcementType                      *string `tfsdk:"enforcement_type" json:"enforcementType,omitempty"`
				InsertSameSiteAttribute              *string `tfsdk:"insert_same_site_attribute" json:"insertSameSiteAttribute,omitempty"`
				MaskValueInLogs                      *bool   `tfsdk:"mask_value_in_logs" json:"maskValueInLogs,omitempty"`
				Name                                 *string `tfsdk:"name" json:"name,omitempty"`
				SecuredOverHttpsConnection           *bool   `tfsdk:"secured_over_https_connection" json:"securedOverHttpsConnection,omitempty"`
				SignatureOverrides                   *[]struct {
					Enabled     *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
					Name        *string `tfsdk:"name" json:"name,omitempty"`
					SignatureId *int64  `tfsdk:"signature_id" json:"signatureId,omitempty"`
					Tag         *string `tfsdk:"tag" json:"tag,omitempty"`
				} `tfsdk:"signature_overrides" json:"signatureOverrides,omitempty"`
				Type          *string `tfsdk:"type" json:"type,omitempty"`
				WildcardOrder *int64  `tfsdk:"wildcard_order" json:"wildcardOrder,omitempty"`
			} `tfsdk:"cookies" json:"cookies,omitempty"`
			Csrf_protection *struct {
				Enabled                 *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
				ExpirationTimeInSeconds *string `tfsdk:"expiration_time_in_seconds" json:"expirationTimeInSeconds,omitempty"`
				SslOnly                 *bool   `tfsdk:"ssl_only" json:"sslOnly,omitempty"`
			} `tfsdk:"csrf_protection" json:"csrf-protection,omitempty"`
			Csrf_urls *[]struct {
				Dollaraction      *string `tfsdk:"dollaraction" json:"$action,omitempty"`
				EnforcementAction *string `tfsdk:"enforcement_action" json:"enforcementAction,omitempty"`
				Method            *string `tfsdk:"method" json:"method,omitempty"`
				Url               *string `tfsdk:"url" json:"url,omitempty"`
				WildcardOrder     *int64  `tfsdk:"wildcard_order" json:"wildcardOrder,omitempty"`
			} `tfsdk:"csrf_urls" json:"csrf-urls,omitempty"`
			Data_guard *struct {
				CreditCardNumbers             *bool     `tfsdk:"credit_card_numbers" json:"creditCardNumbers,omitempty"`
				CustomPatterns                *bool     `tfsdk:"custom_patterns" json:"customPatterns,omitempty"`
				CustomPatternsList            *[]string `tfsdk:"custom_patterns_list" json:"customPatternsList,omitempty"`
				Enabled                       *bool     `tfsdk:"enabled" json:"enabled,omitempty"`
				EnforcementMode               *string   `tfsdk:"enforcement_mode" json:"enforcementMode,omitempty"`
				EnforcementUrls               *[]string `tfsdk:"enforcement_urls" json:"enforcementUrls,omitempty"`
				FirstCustomCharactersToExpose *int64    `tfsdk:"first_custom_characters_to_expose" json:"firstCustomCharactersToExpose,omitempty"`
				LastCcnDigitsToExpose         *int64    `tfsdk:"last_ccn_digits_to_expose" json:"lastCcnDigitsToExpose,omitempty"`
				LastCustomCharactersToExpose  *int64    `tfsdk:"last_custom_characters_to_expose" json:"lastCustomCharactersToExpose,omitempty"`
				LastSsnDigitsToExpose         *int64    `tfsdk:"last_ssn_digits_to_expose" json:"lastSsnDigitsToExpose,omitempty"`
				MaskData                      *bool     `tfsdk:"mask_data" json:"maskData,omitempty"`
				UsSocialSecurityNumbers       *bool     `tfsdk:"us_social_security_numbers" json:"usSocialSecurityNumbers,omitempty"`
			} `tfsdk:"data_guard" json:"data-guard,omitempty"`
			DataGuardReference *struct {
				Link *string `tfsdk:"link" json:"link,omitempty"`
			} `tfsdk:"data_guard_reference" json:"dataGuardReference,omitempty"`
			Description       *string `tfsdk:"description" json:"description,omitempty"`
			EnablePassiveMode *bool   `tfsdk:"enable_passive_mode" json:"enablePassiveMode,omitempty"`
			EnforcementMode   *string `tfsdk:"enforcement_mode" json:"enforcementMode,omitempty"`
			Enforcer_settings *struct {
				EnforcerStateCookies *struct {
					HttpOnlyAttribute *bool   `tfsdk:"http_only_attribute" json:"httpOnlyAttribute,omitempty"`
					SameSiteAttribute *string `tfsdk:"same_site_attribute" json:"sameSiteAttribute,omitempty"`
					SecureAttribute   *string `tfsdk:"secure_attribute" json:"secureAttribute,omitempty"`
				} `tfsdk:"enforcer_state_cookies" json:"enforcerStateCookies,omitempty"`
			} `tfsdk:"enforcer_settings" json:"enforcer-settings,omitempty"`
			FiletypeReference *struct {
				Link *string `tfsdk:"link" json:"link,omitempty"`
			} `tfsdk:"filetype_reference" json:"filetypeReference,omitempty"`
			Filetypes *[]struct {
				Dollaraction           *string `tfsdk:"dollaraction" json:"$action,omitempty"`
				Allowed                *bool   `tfsdk:"allowed" json:"allowed,omitempty"`
				CheckPostDataLength    *bool   `tfsdk:"check_post_data_length" json:"checkPostDataLength,omitempty"`
				CheckQueryStringLength *bool   `tfsdk:"check_query_string_length" json:"checkQueryStringLength,omitempty"`
				CheckRequestLength     *bool   `tfsdk:"check_request_length" json:"checkRequestLength,omitempty"`
				CheckUrlLength         *bool   `tfsdk:"check_url_length" json:"checkUrlLength,omitempty"`
				Name                   *string `tfsdk:"name" json:"name,omitempty"`
				PostDataLength         *int64  `tfsdk:"post_data_length" json:"postDataLength,omitempty"`
				QueryStringLength      *int64  `tfsdk:"query_string_length" json:"queryStringLength,omitempty"`
				RequestLength          *int64  `tfsdk:"request_length" json:"requestLength,omitempty"`
				ResponseCheck          *bool   `tfsdk:"response_check" json:"responseCheck,omitempty"`
				Type                   *string `tfsdk:"type" json:"type,omitempty"`
				UrlLength              *int64  `tfsdk:"url_length" json:"urlLength,omitempty"`
				WildcardOrder          *int64  `tfsdk:"wildcard_order" json:"wildcardOrder,omitempty"`
			} `tfsdk:"filetypes" json:"filetypes,omitempty"`
			FullPath *string `tfsdk:"full_path" json:"fullPath,omitempty"`
			General  *struct {
				AllowedResponseCodes           *[]string `tfsdk:"allowed_response_codes" json:"allowedResponseCodes,omitempty"`
				CustomXffHeaders               *[]string `tfsdk:"custom_xff_headers" json:"customXffHeaders,omitempty"`
				MaskCreditCardNumbersInRequest *bool     `tfsdk:"mask_credit_card_numbers_in_request" json:"maskCreditCardNumbersInRequest,omitempty"`
				TrustXff                       *bool     `tfsdk:"trust_xff" json:"trustXff,omitempty"`
			} `tfsdk:"general" json:"general,omitempty"`
			GeneralReference *struct {
				Link *string `tfsdk:"link" json:"link,omitempty"`
			} `tfsdk:"general_reference" json:"generalReference,omitempty"`
			Graphql_profiles *[]struct {
				Dollaraction          *string `tfsdk:"dollaraction" json:"$action,omitempty"`
				AttackSignaturesCheck *bool   `tfsdk:"attack_signatures_check" json:"attackSignaturesCheck,omitempty"`
				DefenseAttributes     *struct {
					AllowIntrospectionQueries *bool   `tfsdk:"allow_introspection_queries" json:"allowIntrospectionQueries,omitempty"`
					MaximumBatchedQueries     *string `tfsdk:"maximum_batched_queries" json:"maximumBatchedQueries,omitempty"`
					MaximumQueryCost          *string `tfsdk:"maximum_query_cost" json:"maximumQueryCost,omitempty"`
					MaximumStructureDepth     *string `tfsdk:"maximum_structure_depth" json:"maximumStructureDepth,omitempty"`
					MaximumTotalLength        *string `tfsdk:"maximum_total_length" json:"maximumTotalLength,omitempty"`
					MaximumValueLength        *string `tfsdk:"maximum_value_length" json:"maximumValueLength,omitempty"`
					TolerateParsingWarnings   *bool   `tfsdk:"tolerate_parsing_warnings" json:"tolerateParsingWarnings,omitempty"`
				} `tfsdk:"defense_attributes" json:"defenseAttributes,omitempty"`
				Description          *string `tfsdk:"description" json:"description,omitempty"`
				MetacharElementCheck *bool   `tfsdk:"metachar_element_check" json:"metacharElementCheck,omitempty"`
				MetacharOverrides    *[]struct {
					IsAllowed *bool   `tfsdk:"is_allowed" json:"isAllowed,omitempty"`
					Metachar  *string `tfsdk:"metachar" json:"metachar,omitempty"`
				} `tfsdk:"metachar_overrides" json:"metacharOverrides,omitempty"`
				Name                *string `tfsdk:"name" json:"name,omitempty"`
				ResponseEnforcement *struct {
					BlockDisallowedPatterns *bool     `tfsdk:"block_disallowed_patterns" json:"blockDisallowedPatterns,omitempty"`
					DisallowedPatterns      *[]string `tfsdk:"disallowed_patterns" json:"disallowedPatterns,omitempty"`
				} `tfsdk:"response_enforcement" json:"responseEnforcement,omitempty"`
				SensitiveData *[]struct {
					ParameterName *string `tfsdk:"parameter_name" json:"parameterName,omitempty"`
				} `tfsdk:"sensitive_data" json:"sensitiveData,omitempty"`
				SignatureOverrides *[]struct {
					Enabled     *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
					Name        *string `tfsdk:"name" json:"name,omitempty"`
					SignatureId *int64  `tfsdk:"signature_id" json:"signatureId,omitempty"`
					Tag         *string `tfsdk:"tag" json:"tag,omitempty"`
				} `tfsdk:"signature_overrides" json:"signatureOverrides,omitempty"`
			} `tfsdk:"graphql_profiles" json:"graphql-profiles,omitempty"`
			Grpc_profiles *[]struct {
				Dollaraction               *string `tfsdk:"dollaraction" json:"$action,omitempty"`
				AssociateUrls              *bool   `tfsdk:"associate_urls" json:"associateUrls,omitempty"`
				AttackSignaturesCheck      *bool   `tfsdk:"attack_signatures_check" json:"attackSignaturesCheck,omitempty"`
				DecodeStringValuesAsBase64 *string `tfsdk:"decode_string_values_as_base64" json:"decodeStringValuesAsBase64,omitempty"`
				DefenseAttributes          *struct {
					AllowUnknownFields *bool   `tfsdk:"allow_unknown_fields" json:"allowUnknownFields,omitempty"`
					MaximumDataLength  *string `tfsdk:"maximum_data_length" json:"maximumDataLength,omitempty"`
				} `tfsdk:"defense_attributes" json:"defenseAttributes,omitempty"`
				Description *string `tfsdk:"description" json:"description,omitempty"`
				HasIdlFiles *bool   `tfsdk:"has_idl_files" json:"hasIdlFiles,omitempty"`
				IdlFiles    *[]struct {
					IdlFile *struct {
						Contents *string `tfsdk:"contents" json:"contents,omitempty"`
						FileName *string `tfsdk:"file_name" json:"fileName,omitempty"`
						IsBase64 *bool   `tfsdk:"is_base64" json:"isBase64,omitempty"`
					} `tfsdk:"idl_file" json:"idlFile,omitempty"`
					ImportUrl          *string `tfsdk:"import_url" json:"importUrl,omitempty"`
					IsPrimary          *bool   `tfsdk:"is_primary" json:"isPrimary,omitempty"`
					PrimaryIdlFileName *string `tfsdk:"primary_idl_file_name" json:"primaryIdlFileName,omitempty"`
				} `tfsdk:"idl_files" json:"idlFiles,omitempty"`
				MetacharCheck        *bool   `tfsdk:"metachar_check" json:"metacharCheck,omitempty"`
				MetacharElementCheck *bool   `tfsdk:"metachar_element_check" json:"metacharElementCheck,omitempty"`
				Name                 *string `tfsdk:"name" json:"name,omitempty"`
				SignatureOverrides   *[]struct {
					Enabled     *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
					Name        *string `tfsdk:"name" json:"name,omitempty"`
					SignatureId *int64  `tfsdk:"signature_id" json:"signatureId,omitempty"`
					Tag         *string `tfsdk:"tag" json:"tag,omitempty"`
				} `tfsdk:"signature_overrides" json:"signatureOverrides,omitempty"`
			} `tfsdk:"grpc_profiles" json:"grpc-profiles,omitempty"`
			Header_settings *struct {
				MaximumHttpHeaderLength *string `tfsdk:"maximum_http_header_length" json:"maximumHttpHeaderLength,omitempty"`
			} `tfsdk:"header_settings" json:"header-settings,omitempty"`
			HeaderReference *struct {
				Link *string `tfsdk:"link" json:"link,omitempty"`
			} `tfsdk:"header_reference" json:"headerReference,omitempty"`
			HeaderSettingsReference *struct {
				Link *string `tfsdk:"link" json:"link,omitempty"`
			} `tfsdk:"header_settings_reference" json:"headerSettingsReference,omitempty"`
			Headers *[]struct {
				Dollaraction             *string `tfsdk:"dollaraction" json:"$action,omitempty"`
				AllowRepeatedOccurrences *bool   `tfsdk:"allow_repeated_occurrences" json:"allowRepeatedOccurrences,omitempty"`
				Base64Decoding           *bool   `tfsdk:"base64_decoding" json:"base64Decoding,omitempty"`
				CheckSignatures          *bool   `tfsdk:"check_signatures" json:"checkSignatures,omitempty"`
				DecodeValueAsBase64      *string `tfsdk:"decode_value_as_base64" json:"decodeValueAsBase64,omitempty"`
				HtmlNormalization        *bool   `tfsdk:"html_normalization" json:"htmlNormalization,omitempty"`
				Mandatory                *bool   `tfsdk:"mandatory" json:"mandatory,omitempty"`
				MaskValueInLogs          *bool   `tfsdk:"mask_value_in_logs" json:"maskValueInLogs,omitempty"`
				Name                     *string `tfsdk:"name" json:"name,omitempty"`
				NormalizationViolations  *bool   `tfsdk:"normalization_violations" json:"normalizationViolations,omitempty"`
				PercentDecoding          *bool   `tfsdk:"percent_decoding" json:"percentDecoding,omitempty"`
				SignatureOverrides       *[]struct {
					Enabled     *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
					Name        *string `tfsdk:"name" json:"name,omitempty"`
					SignatureId *int64  `tfsdk:"signature_id" json:"signatureId,omitempty"`
					Tag         *string `tfsdk:"tag" json:"tag,omitempty"`
				} `tfsdk:"signature_overrides" json:"signatureOverrides,omitempty"`
				Type             *string `tfsdk:"type" json:"type,omitempty"`
				UrlNormalization *bool   `tfsdk:"url_normalization" json:"urlNormalization,omitempty"`
				WildcardOrder    *int64  `tfsdk:"wildcard_order" json:"wildcardOrder,omitempty"`
			} `tfsdk:"headers" json:"headers,omitempty"`
			Host_names *[]struct {
				Dollaraction      *string `tfsdk:"dollaraction" json:"$action,omitempty"`
				IncludeSubdomains *bool   `tfsdk:"include_subdomains" json:"includeSubdomains,omitempty"`
				Name              *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"host_names" json:"host-names,omitempty"`
			Idl_files *[]struct {
				Contents *string `tfsdk:"contents" json:"contents,omitempty"`
				FileName *string `tfsdk:"file_name" json:"fileName,omitempty"`
				IsBase64 *bool   `tfsdk:"is_base64" json:"isBase64,omitempty"`
			} `tfsdk:"idl_files" json:"idl-files,omitempty"`
			Json_profiles *[]struct {
				Dollaraction          *string `tfsdk:"dollaraction" json:"$action,omitempty"`
				AttackSignaturesCheck *bool   `tfsdk:"attack_signatures_check" json:"attackSignaturesCheck,omitempty"`
				DefenseAttributes     *struct {
					MaximumArrayLength           *string `tfsdk:"maximum_array_length" json:"maximumArrayLength,omitempty"`
					MaximumStructureDepth        *string `tfsdk:"maximum_structure_depth" json:"maximumStructureDepth,omitempty"`
					MaximumTotalLengthOfJSONData *string `tfsdk:"maximum_total_length_of_json_data" json:"maximumTotalLengthOfJSONData,omitempty"`
					MaximumValueLength           *string `tfsdk:"maximum_value_length" json:"maximumValueLength,omitempty"`
					TolerateJSONParsingWarnings  *bool   `tfsdk:"tolerate_json_parsing_warnings" json:"tolerateJSONParsingWarnings,omitempty"`
				} `tfsdk:"defense_attributes" json:"defenseAttributes,omitempty"`
				Description                  *string `tfsdk:"description" json:"description,omitempty"`
				HandleJsonValuesAsParameters *bool   `tfsdk:"handle_json_values_as_parameters" json:"handleJsonValuesAsParameters,omitempty"`
				HasValidationFiles           *bool   `tfsdk:"has_validation_files" json:"hasValidationFiles,omitempty"`
				MetacharOverrides            *[]struct {
					IsAllowed *bool   `tfsdk:"is_allowed" json:"isAllowed,omitempty"`
					Metachar  *string `tfsdk:"metachar" json:"metachar,omitempty"`
				} `tfsdk:"metachar_overrides" json:"metacharOverrides,omitempty"`
				Name               *string `tfsdk:"name" json:"name,omitempty"`
				SignatureOverrides *[]struct {
					Enabled     *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
					Name        *string `tfsdk:"name" json:"name,omitempty"`
					SignatureId *int64  `tfsdk:"signature_id" json:"signatureId,omitempty"`
					Tag         *string `tfsdk:"tag" json:"tag,omitempty"`
				} `tfsdk:"signature_overrides" json:"signatureOverrides,omitempty"`
				ValidationFiles *[]struct {
					ImportUrl          *string `tfsdk:"import_url" json:"importUrl,omitempty"`
					IsPrimary          *bool   `tfsdk:"is_primary" json:"isPrimary,omitempty"`
					JsonValidationFile *struct {
						Dollaraction *string `tfsdk:"dollaraction" json:"$action,omitempty"`
						Contents     *string `tfsdk:"contents" json:"contents,omitempty"`
						FileName     *string `tfsdk:"file_name" json:"fileName,omitempty"`
						IsBase64     *bool   `tfsdk:"is_base64" json:"isBase64,omitempty"`
					} `tfsdk:"json_validation_file" json:"jsonValidationFile,omitempty"`
				} `tfsdk:"validation_files" json:"validationFiles,omitempty"`
			} `tfsdk:"json_profiles" json:"json-profiles,omitempty"`
			Json_validation_files *[]struct {
				Dollaraction *string `tfsdk:"dollaraction" json:"$action,omitempty"`
				Contents     *string `tfsdk:"contents" json:"contents,omitempty"`
				FileName     *string `tfsdk:"file_name" json:"fileName,omitempty"`
				IsBase64     *bool   `tfsdk:"is_base64" json:"isBase64,omitempty"`
			} `tfsdk:"json_validation_files" json:"json-validation-files,omitempty"`
			JsonProfileReference *struct {
				Link *string `tfsdk:"link" json:"link,omitempty"`
			} `tfsdk:"json_profile_reference" json:"jsonProfileReference,omitempty"`
			JsonValidationFileReference *struct {
				Link *string `tfsdk:"link" json:"link,omitempty"`
			} `tfsdk:"json_validation_file_reference" json:"jsonValidationFileReference,omitempty"`
			MethodReference *struct {
				Link *string `tfsdk:"link" json:"link,omitempty"`
			} `tfsdk:"method_reference" json:"methodReference,omitempty"`
			Methods *[]struct {
				Dollaraction *string `tfsdk:"dollaraction" json:"$action,omitempty"`
				Name         *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"methods" json:"methods,omitempty"`
			Name           *string `tfsdk:"name" json:"name,omitempty"`
			Open_api_files *[]struct {
				Link *string `tfsdk:"link" json:"link,omitempty"`
			} `tfsdk:"open_api_files" json:"open-api-files,omitempty"`
			ParameterReference *struct {
				Link *string `tfsdk:"link" json:"link,omitempty"`
			} `tfsdk:"parameter_reference" json:"parameterReference,omitempty"`
			Parameters *[]struct {
				Dollaraction               *string `tfsdk:"dollaraction" json:"$action,omitempty"`
				AllowEmptyValue            *bool   `tfsdk:"allow_empty_value" json:"allowEmptyValue,omitempty"`
				AllowRepeatedParameterName *bool   `tfsdk:"allow_repeated_parameter_name" json:"allowRepeatedParameterName,omitempty"`
				ArraySerializationFormat   *string `tfsdk:"array_serialization_format" json:"arraySerializationFormat,omitempty"`
				AttackSignaturesCheck      *bool   `tfsdk:"attack_signatures_check" json:"attackSignaturesCheck,omitempty"`
				CheckMaxValue              *bool   `tfsdk:"check_max_value" json:"checkMaxValue,omitempty"`
				CheckMaxValueLength        *bool   `tfsdk:"check_max_value_length" json:"checkMaxValueLength,omitempty"`
				CheckMetachars             *bool   `tfsdk:"check_metachars" json:"checkMetachars,omitempty"`
				CheckMinValue              *bool   `tfsdk:"check_min_value" json:"checkMinValue,omitempty"`
				CheckMinValueLength        *bool   `tfsdk:"check_min_value_length" json:"checkMinValueLength,omitempty"`
				CheckMultipleOfValue       *bool   `tfsdk:"check_multiple_of_value" json:"checkMultipleOfValue,omitempty"`
				ContentProfile             *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"content_profile" json:"contentProfile,omitempty"`
				DataType                        *string `tfsdk:"data_type" json:"dataType,omitempty"`
				DecodeValueAsBase64             *string `tfsdk:"decode_value_as_base64" json:"decodeValueAsBase64,omitempty"`
				DisallowFileUploadOfExecutables *bool   `tfsdk:"disallow_file_upload_of_executables" json:"disallowFileUploadOfExecutables,omitempty"`
				EnableRegularExpression         *bool   `tfsdk:"enable_regular_expression" json:"enableRegularExpression,omitempty"`
				ExclusiveMax                    *bool   `tfsdk:"exclusive_max" json:"exclusiveMax,omitempty"`
				ExclusiveMin                    *bool   `tfsdk:"exclusive_min" json:"exclusiveMin,omitempty"`
				IsBase64                        *bool   `tfsdk:"is_base64" json:"isBase64,omitempty"`
				IsCookie                        *bool   `tfsdk:"is_cookie" json:"isCookie,omitempty"`
				IsHeader                        *bool   `tfsdk:"is_header" json:"isHeader,omitempty"`
				Level                           *string `tfsdk:"level" json:"level,omitempty"`
				Mandatory                       *bool   `tfsdk:"mandatory" json:"mandatory,omitempty"`
				MaximumLength                   *int64  `tfsdk:"maximum_length" json:"maximumLength,omitempty"`
				MaximumValue                    *int64  `tfsdk:"maximum_value" json:"maximumValue,omitempty"`
				MetacharsOnParameterValueCheck  *bool   `tfsdk:"metachars_on_parameter_value_check" json:"metacharsOnParameterValueCheck,omitempty"`
				MinimumLength                   *int64  `tfsdk:"minimum_length" json:"minimumLength,omitempty"`
				MinimumValue                    *int64  `tfsdk:"minimum_value" json:"minimumValue,omitempty"`
				MultipleOf                      *int64  `tfsdk:"multiple_of" json:"multipleOf,omitempty"`
				Name                            *string `tfsdk:"name" json:"name,omitempty"`
				NameMetacharOverrides           *[]struct {
					IsAllowed *bool   `tfsdk:"is_allowed" json:"isAllowed,omitempty"`
					Metachar  *string `tfsdk:"metachar" json:"metachar,omitempty"`
				} `tfsdk:"name_metachar_overrides" json:"nameMetacharOverrides,omitempty"`
				ObjectSerializationStyle *string   `tfsdk:"object_serialization_style" json:"objectSerializationStyle,omitempty"`
				ParameterEnumValues      *[]string `tfsdk:"parameter_enum_values" json:"parameterEnumValues,omitempty"`
				ParameterLocation        *string   `tfsdk:"parameter_location" json:"parameterLocation,omitempty"`
				RegularExpression        *string   `tfsdk:"regular_expression" json:"regularExpression,omitempty"`
				SensitiveParameter       *bool     `tfsdk:"sensitive_parameter" json:"sensitiveParameter,omitempty"`
				SignatureOverrides       *[]struct {
					Enabled     *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
					Name        *string `tfsdk:"name" json:"name,omitempty"`
					SignatureId *int64  `tfsdk:"signature_id" json:"signatureId,omitempty"`
					Tag         *string `tfsdk:"tag" json:"tag,omitempty"`
				} `tfsdk:"signature_overrides" json:"signatureOverrides,omitempty"`
				StaticValues *string `tfsdk:"static_values" json:"staticValues,omitempty"`
				Type         *string `tfsdk:"type" json:"type,omitempty"`
				Url          *struct {
					Method   *string `tfsdk:"method" json:"method,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
					Type     *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"url" json:"url,omitempty"`
				ValueMetacharOverrides *[]struct {
					IsAllowed *bool   `tfsdk:"is_allowed" json:"isAllowed,omitempty"`
					Metachar  *string `tfsdk:"metachar" json:"metachar,omitempty"`
				} `tfsdk:"value_metachar_overrides" json:"valueMetacharOverrides,omitempty"`
				ValueType     *string `tfsdk:"value_type" json:"valueType,omitempty"`
				WildcardOrder *int64  `tfsdk:"wildcard_order" json:"wildcardOrder,omitempty"`
			} `tfsdk:"parameters" json:"parameters,omitempty"`
			Response_pages *[]struct {
				AjaxActionType      *string `tfsdk:"ajax_action_type" json:"ajaxActionType,omitempty"`
				AjaxCustomContent   *string `tfsdk:"ajax_custom_content" json:"ajaxCustomContent,omitempty"`
				AjaxEnabled         *bool   `tfsdk:"ajax_enabled" json:"ajaxEnabled,omitempty"`
				AjaxPopupMessage    *string `tfsdk:"ajax_popup_message" json:"ajaxPopupMessage,omitempty"`
				AjaxRedirectUrl     *string `tfsdk:"ajax_redirect_url" json:"ajaxRedirectUrl,omitempty"`
				GrpcStatusCode      *string `tfsdk:"grpc_status_code" json:"grpcStatusCode,omitempty"`
				GrpcStatusMessage   *string `tfsdk:"grpc_status_message" json:"grpcStatusMessage,omitempty"`
				ResponseActionType  *string `tfsdk:"response_action_type" json:"responseActionType,omitempty"`
				ResponseContent     *string `tfsdk:"response_content" json:"responseContent,omitempty"`
				ResponseHeader      *string `tfsdk:"response_header" json:"responseHeader,omitempty"`
				ResponsePageType    *string `tfsdk:"response_page_type" json:"responsePageType,omitempty"`
				ResponseRedirectUrl *string `tfsdk:"response_redirect_url" json:"responseRedirectUrl,omitempty"`
			} `tfsdk:"response_pages" json:"response-pages,omitempty"`
			ResponsePageReference *struct {
				Link *string `tfsdk:"link" json:"link,omitempty"`
			} `tfsdk:"response_page_reference" json:"responsePageReference,omitempty"`
			Sensitive_parameters *[]struct {
				Dollaraction *string `tfsdk:"dollaraction" json:"$action,omitempty"`
				Name         *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"sensitive_parameters" json:"sensitive-parameters,omitempty"`
			SensitiveParameterReference *struct {
				Link *string `tfsdk:"link" json:"link,omitempty"`
			} `tfsdk:"sensitive_parameter_reference" json:"sensitiveParameterReference,omitempty"`
			Server_technologies *[]struct {
				Dollaraction         *string `tfsdk:"dollaraction" json:"$action,omitempty"`
				ServerTechnologyName *string `tfsdk:"server_technology_name" json:"serverTechnologyName,omitempty"`
			} `tfsdk:"server_technologies" json:"server-technologies,omitempty"`
			ServerTechnologyReference *struct {
				Link *string `tfsdk:"link" json:"link,omitempty"`
			} `tfsdk:"server_technology_reference" json:"serverTechnologyReference,omitempty"`
			Signature_requirements *[]struct {
				Dollaraction *string `tfsdk:"dollaraction" json:"$action,omitempty"`
				Tag          *string `tfsdk:"tag" json:"tag,omitempty"`
			} `tfsdk:"signature_requirements" json:"signature-requirements,omitempty"`
			Signature_sets     *[]map[string]string `tfsdk:"signature_sets" json:"signature-sets,omitempty"`
			Signature_settings *struct {
				AttackSignatureFalsePositiveMode      *string `tfsdk:"attack_signature_false_positive_mode" json:"attackSignatureFalsePositiveMode,omitempty"`
				MinimumAccuracyForAutoAddedSignatures *string `tfsdk:"minimum_accuracy_for_auto_added_signatures" json:"minimumAccuracyForAutoAddedSignatures,omitempty"`
			} `tfsdk:"signature_settings" json:"signature-settings,omitempty"`
			SignatureReference *struct {
				Link *string `tfsdk:"link" json:"link,omitempty"`
			} `tfsdk:"signature_reference" json:"signatureReference,omitempty"`
			SignatureSetReference *struct {
				Link *string `tfsdk:"link" json:"link,omitempty"`
			} `tfsdk:"signature_set_reference" json:"signatureSetReference,omitempty"`
			SignatureSettingReference *struct {
				Link *string `tfsdk:"link" json:"link,omitempty"`
			} `tfsdk:"signature_setting_reference" json:"signatureSettingReference,omitempty"`
			Signatures *[]struct {
				Enabled     *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
				Name        *string `tfsdk:"name" json:"name,omitempty"`
				SignatureId *int64  `tfsdk:"signature_id" json:"signatureId,omitempty"`
				Tag         *string `tfsdk:"tag" json:"tag,omitempty"`
			} `tfsdk:"signatures" json:"signatures,omitempty"`
			SoftwareVersion *string `tfsdk:"software_version" json:"softwareVersion,omitempty"`
			Template        *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"template" json:"template,omitempty"`
			Threat_campaigns *[]struct {
				IsEnabled *bool   `tfsdk:"is_enabled" json:"isEnabled,omitempty"`
				Name      *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"threat_campaigns" json:"threat-campaigns,omitempty"`
			ThreatCampaignReference *struct {
				Link *string `tfsdk:"link" json:"link,omitempty"`
			} `tfsdk:"threat_campaign_reference" json:"threatCampaignReference,omitempty"`
			UrlReference *struct {
				Link *string `tfsdk:"link" json:"link,omitempty"`
			} `tfsdk:"url_reference" json:"urlReference,omitempty"`
			Urls *[]struct {
				Dollaraction                        *string `tfsdk:"dollaraction" json:"$action,omitempty"`
				AllowRenderingInFrames              *string `tfsdk:"allow_rendering_in_frames" json:"allowRenderingInFrames,omitempty"`
				AllowRenderingInFramesOnlyFrom      *string `tfsdk:"allow_rendering_in_frames_only_from" json:"allowRenderingInFramesOnlyFrom,omitempty"`
				AttackSignaturesCheck               *bool   `tfsdk:"attack_signatures_check" json:"attackSignaturesCheck,omitempty"`
				ClickjackingProtection              *bool   `tfsdk:"clickjacking_protection" json:"clickjackingProtection,omitempty"`
				Description                         *string `tfsdk:"description" json:"description,omitempty"`
				DisallowFileUploadOfExecutables     *bool   `tfsdk:"disallow_file_upload_of_executables" json:"disallowFileUploadOfExecutables,omitempty"`
				Html5CrossOriginRequestsEnforcement *struct {
					AllowOriginsEnforcementMode *string `tfsdk:"allow_origins_enforcement_mode" json:"allowOriginsEnforcementMode,omitempty"`
					CheckAllowedMethods         *bool   `tfsdk:"check_allowed_methods" json:"checkAllowedMethods,omitempty"`
					CrossDomainAllowedOrigin    *[]struct {
						IncludeSubDomains *bool   `tfsdk:"include_sub_domains" json:"includeSubDomains,omitempty"`
						OriginName        *string `tfsdk:"origin_name" json:"originName,omitempty"`
						OriginPort        *string `tfsdk:"origin_port" json:"originPort,omitempty"`
						OriginProtocol    *string `tfsdk:"origin_protocol" json:"originProtocol,omitempty"`
					} `tfsdk:"cross_domain_allowed_origin" json:"crossDomainAllowedOrigin,omitempty"`
					EnforcementMode *string `tfsdk:"enforcement_mode" json:"enforcementMode,omitempty"`
				} `tfsdk:"html5_cross_origin_requests_enforcement" json:"html5CrossOriginRequestsEnforcement,omitempty"`
				IsAllowed         *bool `tfsdk:"is_allowed" json:"isAllowed,omitempty"`
				MandatoryBody     *bool `tfsdk:"mandatory_body" json:"mandatoryBody,omitempty"`
				MetacharOverrides *[]struct {
					IsAllowed *bool   `tfsdk:"is_allowed" json:"isAllowed,omitempty"`
					Metachar  *string `tfsdk:"metachar" json:"metachar,omitempty"`
				} `tfsdk:"metachar_overrides" json:"metacharOverrides,omitempty"`
				MetacharsOnUrlCheck *bool   `tfsdk:"metachars_on_url_check" json:"metacharsOnUrlCheck,omitempty"`
				Method              *string `tfsdk:"method" json:"method,omitempty"`
				MethodOverrides     *[]struct {
					Allowed *bool   `tfsdk:"allowed" json:"allowed,omitempty"`
					Method  *string `tfsdk:"method" json:"method,omitempty"`
				} `tfsdk:"method_overrides" json:"methodOverrides,omitempty"`
				MethodsOverrideOnUrlCheck *bool   `tfsdk:"methods_override_on_url_check" json:"methodsOverrideOnUrlCheck,omitempty"`
				Name                      *string `tfsdk:"name" json:"name,omitempty"`
				OperationId               *string `tfsdk:"operation_id" json:"operationId,omitempty"`
				PositionalParameters      *[]struct {
					Parameter *struct {
						Dollaraction               *string `tfsdk:"dollaraction" json:"$action,omitempty"`
						AllowEmptyValue            *bool   `tfsdk:"allow_empty_value" json:"allowEmptyValue,omitempty"`
						AllowRepeatedParameterName *bool   `tfsdk:"allow_repeated_parameter_name" json:"allowRepeatedParameterName,omitempty"`
						ArraySerializationFormat   *string `tfsdk:"array_serialization_format" json:"arraySerializationFormat,omitempty"`
						AttackSignaturesCheck      *bool   `tfsdk:"attack_signatures_check" json:"attackSignaturesCheck,omitempty"`
						CheckMaxValue              *bool   `tfsdk:"check_max_value" json:"checkMaxValue,omitempty"`
						CheckMaxValueLength        *bool   `tfsdk:"check_max_value_length" json:"checkMaxValueLength,omitempty"`
						CheckMetachars             *bool   `tfsdk:"check_metachars" json:"checkMetachars,omitempty"`
						CheckMinValue              *bool   `tfsdk:"check_min_value" json:"checkMinValue,omitempty"`
						CheckMinValueLength        *bool   `tfsdk:"check_min_value_length" json:"checkMinValueLength,omitempty"`
						CheckMultipleOfValue       *bool   `tfsdk:"check_multiple_of_value" json:"checkMultipleOfValue,omitempty"`
						ContentProfile             *struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"content_profile" json:"contentProfile,omitempty"`
						DataType                        *string `tfsdk:"data_type" json:"dataType,omitempty"`
						DecodeValueAsBase64             *string `tfsdk:"decode_value_as_base64" json:"decodeValueAsBase64,omitempty"`
						DisallowFileUploadOfExecutables *bool   `tfsdk:"disallow_file_upload_of_executables" json:"disallowFileUploadOfExecutables,omitempty"`
						EnableRegularExpression         *bool   `tfsdk:"enable_regular_expression" json:"enableRegularExpression,omitempty"`
						ExclusiveMax                    *bool   `tfsdk:"exclusive_max" json:"exclusiveMax,omitempty"`
						ExclusiveMin                    *bool   `tfsdk:"exclusive_min" json:"exclusiveMin,omitempty"`
						IsBase64                        *bool   `tfsdk:"is_base64" json:"isBase64,omitempty"`
						IsCookie                        *bool   `tfsdk:"is_cookie" json:"isCookie,omitempty"`
						IsHeader                        *bool   `tfsdk:"is_header" json:"isHeader,omitempty"`
						Level                           *string `tfsdk:"level" json:"level,omitempty"`
						Mandatory                       *bool   `tfsdk:"mandatory" json:"mandatory,omitempty"`
						MaximumLength                   *int64  `tfsdk:"maximum_length" json:"maximumLength,omitempty"`
						MaximumValue                    *int64  `tfsdk:"maximum_value" json:"maximumValue,omitempty"`
						MetacharsOnParameterValueCheck  *bool   `tfsdk:"metachars_on_parameter_value_check" json:"metacharsOnParameterValueCheck,omitempty"`
						MinimumLength                   *int64  `tfsdk:"minimum_length" json:"minimumLength,omitempty"`
						MinimumValue                    *int64  `tfsdk:"minimum_value" json:"minimumValue,omitempty"`
						MultipleOf                      *int64  `tfsdk:"multiple_of" json:"multipleOf,omitempty"`
						Name                            *string `tfsdk:"name" json:"name,omitempty"`
						NameMetacharOverrides           *[]struct {
							IsAllowed *bool   `tfsdk:"is_allowed" json:"isAllowed,omitempty"`
							Metachar  *string `tfsdk:"metachar" json:"metachar,omitempty"`
						} `tfsdk:"name_metachar_overrides" json:"nameMetacharOverrides,omitempty"`
						ObjectSerializationStyle *string   `tfsdk:"object_serialization_style" json:"objectSerializationStyle,omitempty"`
						ParameterEnumValues      *[]string `tfsdk:"parameter_enum_values" json:"parameterEnumValues,omitempty"`
						ParameterLocation        *string   `tfsdk:"parameter_location" json:"parameterLocation,omitempty"`
						RegularExpression        *string   `tfsdk:"regular_expression" json:"regularExpression,omitempty"`
						SensitiveParameter       *bool     `tfsdk:"sensitive_parameter" json:"sensitiveParameter,omitempty"`
						SignatureOverrides       *[]struct {
							Enabled     *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
							Name        *string `tfsdk:"name" json:"name,omitempty"`
							SignatureId *int64  `tfsdk:"signature_id" json:"signatureId,omitempty"`
							Tag         *string `tfsdk:"tag" json:"tag,omitempty"`
						} `tfsdk:"signature_overrides" json:"signatureOverrides,omitempty"`
						StaticValues *string `tfsdk:"static_values" json:"staticValues,omitempty"`
						Type         *string `tfsdk:"type" json:"type,omitempty"`
						Url          *struct {
							Method   *string `tfsdk:"method" json:"method,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
							Type     *string `tfsdk:"type" json:"type,omitempty"`
						} `tfsdk:"url" json:"url,omitempty"`
						ValueMetacharOverrides *[]struct {
							IsAllowed *bool   `tfsdk:"is_allowed" json:"isAllowed,omitempty"`
							Metachar  *string `tfsdk:"metachar" json:"metachar,omitempty"`
						} `tfsdk:"value_metachar_overrides" json:"valueMetacharOverrides,omitempty"`
						ValueType     *string `tfsdk:"value_type" json:"valueType,omitempty"`
						WildcardOrder *int64  `tfsdk:"wildcard_order" json:"wildcardOrder,omitempty"`
					} `tfsdk:"parameter" json:"parameter,omitempty"`
					UrlSegmentIndex *int64 `tfsdk:"url_segment_index" json:"urlSegmentIndex,omitempty"`
				} `tfsdk:"positional_parameters" json:"positionalParameters,omitempty"`
				Protocol           *string `tfsdk:"protocol" json:"protocol,omitempty"`
				SignatureOverrides *[]struct {
					Enabled     *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
					Name        *string `tfsdk:"name" json:"name,omitempty"`
					SignatureId *int64  `tfsdk:"signature_id" json:"signatureId,omitempty"`
					Tag         *string `tfsdk:"tag" json:"tag,omitempty"`
				} `tfsdk:"signature_overrides" json:"signatureOverrides,omitempty"`
				Type               *string `tfsdk:"type" json:"type,omitempty"`
				UrlContentProfiles *[]struct {
					ContentProfile *struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"content_profile" json:"contentProfile,omitempty"`
					HeaderName  *string `tfsdk:"header_name" json:"headerName,omitempty"`
					HeaderOrder *string `tfsdk:"header_order" json:"headerOrder,omitempty"`
					HeaderValue *string `tfsdk:"header_value" json:"headerValue,omitempty"`
					Name        *string `tfsdk:"name" json:"name,omitempty"`
					Type        *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"url_content_profiles" json:"urlContentProfiles,omitempty"`
				WildcardOrder *int64 `tfsdk:"wildcard_order" json:"wildcardOrder,omitempty"`
			} `tfsdk:"urls" json:"urls,omitempty"`
			Whitelist_ips *[]struct {
				Dollaraction     *string `tfsdk:"dollaraction" json:"$action,omitempty"`
				BlockRequests    *string `tfsdk:"block_requests" json:"blockRequests,omitempty"`
				IpAddress        *string `tfsdk:"ip_address" json:"ipAddress,omitempty"`
				IpMask           *string `tfsdk:"ip_mask" json:"ipMask,omitempty"`
				NeverLogRequests *bool   `tfsdk:"never_log_requests" json:"neverLogRequests,omitempty"`
			} `tfsdk:"whitelist_ips" json:"whitelist-ips,omitempty"`
			WhitelistIpReference *struct {
				Link *string `tfsdk:"link" json:"link,omitempty"`
			} `tfsdk:"whitelist_ip_reference" json:"whitelistIpReference,omitempty"`
			Xml_profiles *[]struct {
				Dollaraction          *string `tfsdk:"dollaraction" json:"$action,omitempty"`
				AttackSignaturesCheck *bool   `tfsdk:"attack_signatures_check" json:"attackSignaturesCheck,omitempty"`
				DefenseAttributes     *struct {
					AllowCDATA                  *bool   `tfsdk:"allow_cdata" json:"allowCDATA,omitempty"`
					AllowDTDs                   *bool   `tfsdk:"allow_dt_ds" json:"allowDTDs,omitempty"`
					AllowExternalReferences     *bool   `tfsdk:"allow_external_references" json:"allowExternalReferences,omitempty"`
					AllowProcessingInstructions *bool   `tfsdk:"allow_processing_instructions" json:"allowProcessingInstructions,omitempty"`
					MaximumAttributeValueLength *string `tfsdk:"maximum_attribute_value_length" json:"maximumAttributeValueLength,omitempty"`
					MaximumAttributesPerElement *string `tfsdk:"maximum_attributes_per_element" json:"maximumAttributesPerElement,omitempty"`
					MaximumChildrenPerElement   *string `tfsdk:"maximum_children_per_element" json:"maximumChildrenPerElement,omitempty"`
					MaximumDocumentDepth        *string `tfsdk:"maximum_document_depth" json:"maximumDocumentDepth,omitempty"`
					MaximumDocumentSize         *string `tfsdk:"maximum_document_size" json:"maximumDocumentSize,omitempty"`
					MaximumElements             *string `tfsdk:"maximum_elements" json:"maximumElements,omitempty"`
					MaximumNSDeclarations       *string `tfsdk:"maximum_ns_declarations" json:"maximumNSDeclarations,omitempty"`
					MaximumNameLength           *string `tfsdk:"maximum_name_length" json:"maximumNameLength,omitempty"`
					MaximumNamespaceLength      *string `tfsdk:"maximum_namespace_length" json:"maximumNamespaceLength,omitempty"`
					TolerateCloseTagShorthand   *bool   `tfsdk:"tolerate_close_tag_shorthand" json:"tolerateCloseTagShorthand,omitempty"`
					TolerateLeadingWhiteSpace   *bool   `tfsdk:"tolerate_leading_white_space" json:"tolerateLeadingWhiteSpace,omitempty"`
					TolerateNumericNames        *bool   `tfsdk:"tolerate_numeric_names" json:"tolerateNumericNames,omitempty"`
				} `tfsdk:"defense_attributes" json:"defenseAttributes,omitempty"`
				Description        *string `tfsdk:"description" json:"description,omitempty"`
				EnableWss          *bool   `tfsdk:"enable_wss" json:"enableWss,omitempty"`
				FollowSchemaLinks  *bool   `tfsdk:"follow_schema_links" json:"followSchemaLinks,omitempty"`
				Name               *string `tfsdk:"name" json:"name,omitempty"`
				SignatureOverrides *[]struct {
					Enabled     *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
					Name        *string `tfsdk:"name" json:"name,omitempty"`
					SignatureId *int64  `tfsdk:"signature_id" json:"signatureId,omitempty"`
					Tag         *string `tfsdk:"tag" json:"tag,omitempty"`
				} `tfsdk:"signature_overrides" json:"signatureOverrides,omitempty"`
				UseXmlResponsePage *bool `tfsdk:"use_xml_response_page" json:"useXmlResponsePage,omitempty"`
			} `tfsdk:"xml_profiles" json:"xml-profiles,omitempty"`
			Xml_validation_files *[]struct {
				Dollaraction *string `tfsdk:"dollaraction" json:"$action,omitempty"`
				Contents     *string `tfsdk:"contents" json:"contents,omitempty"`
				FileName     *string `tfsdk:"file_name" json:"fileName,omitempty"`
				IsBase64     *bool   `tfsdk:"is_base64" json:"isBase64,omitempty"`
			} `tfsdk:"xml_validation_files" json:"xml-validation-files,omitempty"`
			XmlProfileReference *struct {
				Link *string `tfsdk:"link" json:"link,omitempty"`
			} `tfsdk:"xml_profile_reference" json:"xmlProfileReference,omitempty"`
			XmlValidationFileReference *struct {
				Link *string `tfsdk:"link" json:"link,omitempty"`
			} `tfsdk:"xml_validation_file_reference" json:"xmlValidationFileReference,omitempty"`
		} `tfsdk:"policy" json:"policy,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AppprotectF5ComAppolicyV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_appprotect_f5_com_ap_policy_v1beta1_manifest"
}

func (r *AppprotectF5ComAppolicyV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "APPolicyConfig is the Schema for the APPolicyconfigs API",
		MarkdownDescription: "APPolicyConfig is the Schema for the APPolicyconfigs API",
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
				Description:         "APPolicySpec defines the desired state of APPolicy",
				MarkdownDescription: "APPolicySpec defines the desired state of APPolicy",
				Attributes: map[string]schema.Attribute{
					"modifications": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.MapType{ElemType: types.StringType},
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"modifications_reference": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"link": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^http`), ""),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"policy": schema.SingleNestedAttribute{
						Description:         "Defines the App Protect policy",
						MarkdownDescription: "Defines the App Protect policy",
						Attributes: map[string]schema.Attribute{
							"application_language": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("iso-8859-10", "iso-8859-6", "windows-1255", "auto-detect", "koi8-r", "gb18030", "iso-8859-8", "windows-1250", "iso-8859-9", "windows-1252", "iso-8859-16", "gb2312", "iso-8859-2", "iso-8859-5", "windows-1257", "windows-1256", "iso-8859-13", "windows-874", "windows-1253", "iso-8859-3", "euc-jp", "utf-8", "gbk", "windows-1251", "big5", "iso-8859-1", "shift_jis", "euc-kr", "iso-8859-4", "iso-8859-7", "iso-8859-15"),
								},
							},

							"blocking_settings": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"evasions": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("%u decoding", "Apache whitespace", "Bad unescape", "Bare byte decoding", "Directory traversals", "IIS backslashes", "IIS Unicode codepoints", "Multiple decoding", "Multiple slashes", "Semicolon path parameters", "Trailing dot", "Trailing slash"),
													},
												},

												"enabled": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"max_decoding_passes": schema.Int64Attribute{
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

									"http_protocols": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("Unescaped space in URL", "Unparsable request content", "Several Content-Length headers", "POST request with Content-Length: 0", "Null in request", "No Host header in HTTP/1.1 request", "Multiple host headers", "Host header contains IP address", "High ASCII characters in headers", "Header name with no header value", "CRLF characters before request start", "Content length should be a positive number", "Chunked request with Content-Length header", "Check maximum number of cookies", "Check maximum number of parameters", "Check maximum number of headers", "Body in GET or HEAD requests", "Bad multipart/form-data request parsing", "Bad multipart parameters parsing", "Bad HTTP version", "Bad host header value"),
													},
												},

												"enabled": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"max_cookies": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(1),
														int64validator.AtMost(100),
													},
												},

												"max_headers": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(1),
														int64validator.AtMost(150),
													},
												},

												"max_params": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(1),
														int64validator.AtMost(5000),
													},
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"violations": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"alarm": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"block": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
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
													Validators: []validator.String{
														stringvalidator.OneOf("VIOL_ACCESS_INVALID", "VIOL_ACCESS_MALFORMED", "VIOL_ACCESS_MISSING", "VIOL_ACCESS_UNAUTHORIZED", "VIOL_ASM_COOKIE_HIJACKING", "VIOL_ASM_COOKIE_MODIFIED", "VIOL_BLACKLISTED_IP", "VIOL_COOKIE_EXPIRED", "VIOL_COOKIE_LENGTH", "VIOL_COOKIE_MALFORMED", "VIOL_COOKIE_MODIFIED", "VIOL_CSRF", "VIOL_DATA_GUARD", "VIOL_ENCODING", "VIOL_EVASION", "VIOL_FILE_UPLOAD", "VIOL_FILE_UPLOAD_IN_BODY", "VIOL_FILETYPE", "VIOL_GRAPHQL_ERROR_RESPONSE", "VIOL_GRAPHQL_FORMAT", "VIOL_GRAPHQL_INTROSPECTION_QUERY", "VIOL_GRAPHQL_MALFORMED", "VIOL_GRPC_FORMAT", "VIOL_GRPC_MALFORMED", "VIOL_GRPC_METHOD", "VIOL_HEADER_LENGTH", "VIOL_HEADER_METACHAR", "VIOL_HEADER_REPEATED", "VIOL_HTTP_PROTOCOL", "VIOL_HTTP_RESPONSE_STATUS", "VIOL_JSON_FORMAT", "VIOL_JSON_MALFORMED", "VIOL_JSON_SCHEMA", "VIOL_MANDATORY_HEADER", "VIOL_MANDATORY_PARAMETER", "VIOL_MANDATORY_REQUEST_BODY", "VIOL_METHOD", "VIOL_PARAMETER", "VIOL_PARAMETER_ARRAY_VALUE", "VIOL_PARAMETER_DATA_TYPE", "VIOL_PARAMETER_EMPTY_VALUE", "VIOL_PARAMETER_LOCATION", "VIOL_PARAMETER_MULTIPART_NULL_VALUE", "VIOL_PARAMETER_NAME_METACHAR", "VIOL_PARAMETER_NUMERIC_VALUE", "VIOL_PARAMETER_REPEATED", "VIOL_PARAMETER_STATIC_VALUE", "VIOL_PARAMETER_VALUE_BASE64", "VIOL_PARAMETER_VALUE_LENGTH", "VIOL_PARAMETER_VALUE_METACHAR", "VIOL_PARAMETER_VALUE_REGEXP", "VIOL_POST_DATA_LENGTH", "VIOL_QUERY_STRING_LENGTH", "VIOL_RATING_NEED_EXAMINATION", "VIOL_RATING_THREAT", "VIOL_REQUEST_LENGTH", "VIOL_REQUEST_MAX_LENGTH", "VIOL_THREAT_CAMPAIGN", "VIOL_URL", "VIOL_URL_CONTENT_TYPE", "VIOL_URL_LENGTH", "VIOL_URL_METACHAR", "VIOL_XML_FORMAT", "VIOL_XML_MALFORMED"),
													},
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

							"blocking_setting_reference": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"link": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^http`), ""),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"bot_defense": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"mitigations": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"anomalies": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"dollaraction": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("delete"),
															},
														},

														"action": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("alarm", "block", "default", "detect", "ignore"),
															},
														},

														"name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"score_threshold": schema.StringAttribute{
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

											"browsers": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"dollaraction": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("delete"),
															},
														},

														"action": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("alarm", "block", "detect"),
															},
														},

														"max_version": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(0),
																int64validator.AtMost(2.147483647e+09),
															},
														},

														"min_version": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(0),
																int64validator.AtMost(2.147483647e+09),
															},
														},

														"name": schema.StringAttribute{
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

											"classes": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"action": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("alarm", "block", "detect", "ignore"),
															},
														},

														"name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("browser", "malicious-bot", "suspicious-browser", "trusted-bot", "unknown", "untrusted-bot"),
															},
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"signatures": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"dollaraction": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("delete"),
															},
														},

														"action": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("alarm", "block", "detect", "ignore"),
															},
														},

														"name": schema.StringAttribute{
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
										Required: false,
										Optional: true,
										Computed: false,
									},

									"settings": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"case_sensitive_http_headers": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"is_enabled": schema.BoolAttribute{
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

							"browser_definitions": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"dollaraction": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("delete"),
											},
										},

										"is_user_defined": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"match_regex": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"match_string": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"case_insensitive": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"character_sets": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"character_set": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"is_allowed": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"metachar": schema.StringAttribute{
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

										"character_set_type": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("gwt-content", "header", "json-content", "parameter-name", "parameter-value", "plain-text-content", "url", "xml-content"),
											},
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"character_set_reference": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"link": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^http`), ""),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"cookie_settings": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"maximum_cookie_header_length": schema.StringAttribute{
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

							"cookie_reference": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"link": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^http`), ""),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"cookie_settings_reference": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"link": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^http`), ""),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"cookies": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"dollaraction": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("delete"),
											},
										},

										"accessible_only_through_the_http_protocol": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"attack_signatures_check": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"decode_value_as_base64": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("enabled", "disabled", "required"),
											},
										},

										"enforcement_type": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"insert_same_site_attribute": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("lax", "none", "none-value", "strict"),
											},
										},

										"mask_value_in_logs": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
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

										"secured_over_https_connection": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"signature_overrides": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"enabled": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
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

													"signature_id": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"tag": schema.StringAttribute{
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

										"type": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("explicit", "wildcard"),
											},
										},

										"wildcard_order": schema.Int64Attribute{
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

							"csrf_protection": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"expiration_time_in_seconds": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`disabled|\d+`), ""),
										},
									},

									"ssl_only": schema.BoolAttribute{
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

							"csrf_urls": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"dollaraction": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("delete"),
											},
										},

										"enforcement_action": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("verify-origin", "none"),
											},
										},

										"method": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("GET", "POST", "any"),
											},
										},

										"url": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"wildcard_order": schema.Int64Attribute{
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

							"data_guard": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"credit_card_numbers": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"custom_patterns": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"custom_patterns_list": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"enabled": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"enforcement_mode": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("ignore-urls-in-list", "enforce-urls-in-list"),
										},
									},

									"enforcement_urls": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"first_custom_characters_to_expose": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"last_ccn_digits_to_expose": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"last_custom_characters_to_expose": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"last_ssn_digits_to_expose": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"mask_data": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"us_social_security_numbers": schema.BoolAttribute{
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

							"data_guard_reference": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"link": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^http`), ""),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"description": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_passive_mode": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enforcement_mode": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("transparent", "blocking"),
								},
							},

							"enforcer_settings": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"enforcer_state_cookies": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"http_only_attribute": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"same_site_attribute": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("lax", "none", "none-value", "strict"),
												},
											},

											"secure_attribute": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("always", "never"),
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

							"filetype_reference": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"link": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^http`), ""),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"filetypes": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"dollaraction": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("delete"),
											},
										},

										"allowed": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_post_data_length": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_query_string_length": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_request_length": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_url_length": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
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

										"post_data_length": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"query_string_length": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"request_length": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"response_check": schema.BoolAttribute{
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
											Validators: []validator.String{
												stringvalidator.OneOf("explicit", "wildcard"),
											},
										},

										"url_length": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"wildcard_order": schema.Int64Attribute{
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

							"full_path": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"general": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"allowed_response_codes": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"custom_xff_headers": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"mask_credit_card_numbers_in_request": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"trust_xff": schema.BoolAttribute{
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

							"general_reference": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"link": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^http`), ""),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"graphql_profiles": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"dollaraction": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("delete"),
											},
										},

										"attack_signatures_check": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"defense_attributes": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"allow_introspection_queries": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"maximum_batched_queries": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"maximum_query_cost": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"maximum_structure_depth": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"maximum_total_length": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"maximum_value_length": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"tolerate_parsing_warnings": schema.BoolAttribute{
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

										"description": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"metachar_element_check": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"metachar_overrides": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"is_allowed": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"metachar": schema.StringAttribute{
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

										"name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"response_enforcement": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"block_disallowed_patterns": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"disallowed_patterns": schema.ListAttribute{
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

										"sensitive_data": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"parameter_name": schema.StringAttribute{
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

										"signature_overrides": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"enabled": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
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

													"signature_id": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"tag": schema.StringAttribute{
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

							"grpc_profiles": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"dollaraction": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("delete"),
											},
										},

										"associate_urls": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"attack_signatures_check": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"decode_string_values_as_base64": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("disabled", "enabled"),
											},
										},

										"defense_attributes": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"allow_unknown_fields": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"maximum_data_length": schema.StringAttribute{
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

										"description": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"has_idl_files": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"idl_files": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"idl_file": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"contents": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"file_name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"is_base64": schema.BoolAttribute{
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

													"import_url": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"is_primary": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"primary_idl_file_name": schema.StringAttribute{
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

										"metachar_check": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"metachar_element_check": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
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

										"signature_overrides": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"enabled": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
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

													"signature_id": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"tag": schema.StringAttribute{
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

							"header_settings": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"maximum_http_header_length": schema.StringAttribute{
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

							"header_reference": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"link": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^http`), ""),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"header_settings_reference": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"link": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^http`), ""),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"headers": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"dollaraction": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("delete"),
											},
										},

										"allow_repeated_occurrences": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"base64_decoding": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_signatures": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"decode_value_as_base64": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("enabled", "disabled", "required"),
											},
										},

										"html_normalization": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"mandatory": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"mask_value_in_logs": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
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

										"normalization_violations": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"percent_decoding": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"signature_overrides": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"enabled": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
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

													"signature_id": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"tag": schema.StringAttribute{
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

										"type": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("explicit", "wildcard"),
											},
										},

										"url_normalization": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"wildcard_order": schema.Int64Attribute{
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

							"host_names": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"dollaraction": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("delete"),
											},
										},

										"include_subdomains": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"idl_files": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"contents": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"file_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"is_base64": schema.BoolAttribute{
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

							"json_profiles": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"dollaraction": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("delete"),
											},
										},

										"attack_signatures_check": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"defense_attributes": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"maximum_array_length": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"maximum_structure_depth": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"maximum_total_length_of_json_data": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"maximum_value_length": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"tolerate_json_parsing_warnings": schema.BoolAttribute{
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

										"description": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"handle_json_values_as_parameters": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"has_validation_files": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"metachar_overrides": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"is_allowed": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"metachar": schema.StringAttribute{
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

										"name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"signature_overrides": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"enabled": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
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

													"signature_id": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"tag": schema.StringAttribute{
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

										"validation_files": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"import_url": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"is_primary": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"json_validation_file": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"dollaraction": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("delete"),
																},
															},

															"contents": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"file_name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"is_base64": schema.BoolAttribute{
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

							"json_validation_files": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"dollaraction": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("delete"),
											},
										},

										"contents": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"file_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"is_base64": schema.BoolAttribute{
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

							"json_profile_reference": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"link": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^http`), ""),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"json_validation_file_reference": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"link": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^http`), ""),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"method_reference": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"link": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^http`), ""),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"methods": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"dollaraction": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("delete"),
											},
										},

										"name": schema.StringAttribute{
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

							"name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"open_api_files": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"link": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^http`), ""),
											},
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"parameter_reference": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"link": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^http`), ""),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"parameters": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"dollaraction": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("delete"),
											},
										},

										"allow_empty_value": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"allow_repeated_parameter_name": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"array_serialization_format": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("csv", "form", "label", "matrix", "multi", "multipart", "pipe", "ssv", "tsv"),
											},
										},

										"attack_signatures_check": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_max_value": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_max_value_length": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_metachars": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_min_value": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_min_value_length": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_multiple_of_value": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"content_profile": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
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

										"data_type": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("alpha-numeric", "binary", "boolean", "decimal", "email", "integer", "none", "phone"),
											},
										},

										"decode_value_as_base64": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("enabled", "disabled", "required"),
											},
										},

										"disallow_file_upload_of_executables": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"enable_regular_expression": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclusive_max": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclusive_min": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"is_base64": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"is_cookie": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"is_header": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"level": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("global", "url"),
											},
										},

										"mandatory": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"maximum_length": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"maximum_value": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"metachars_on_parameter_value_check": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"minimum_length": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"minimum_value": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"multiple_of": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
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

										"name_metachar_overrides": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"is_allowed": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"metachar": schema.StringAttribute{
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

										"object_serialization_style": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"parameter_enum_values": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"parameter_location": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("any", "cookie", "form-data", "header", "path", "query"),
											},
										},

										"regular_expression": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"sensitive_parameter": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"signature_overrides": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"enabled": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
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

													"signature_id": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"tag": schema.StringAttribute{
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

										"static_values": schema.StringAttribute{
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
											Validators: []validator.String{
												stringvalidator.OneOf("explicit", "wildcard"),
											},
										},

										"url": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"method": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("ACL", "BCOPY", "BDELETE", "BMOVE", "BPROPFIND", "BPROPPATCH", "CHECKIN", "CHECKOUT", "CONNECT", "COPY", "DELETE", "GET", "HEAD", "LINK", "LOCK", "MERGE", "MKCOL", "MKWORKSPACE", "MOVE", "NOTIFY", "OPTIONS", "PATCH", "POLL", "POST", "PROPFIND", "PROPPATCH", "PUT", "REPORT", "RPC_IN_DATA", "RPC_OUT_DATA", "SEARCH", "SUBSCRIBE", "TRACE", "TRACK", "UNLINK", "UNLOCK", "UNSUBSCRIBE", "VERSION_CONTROL", "X-MS-ENUMATTS", "*"),
													},
												},

												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"protocol": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("http", "https"),
													},
												},

												"type": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("explicit", "wildcard"),
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"value_metachar_overrides": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"is_allowed": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"metachar": schema.StringAttribute{
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

										"value_type": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("array", "auto-detect", "dynamic-content", "dynamic-parameter-name", "ignore", "json", "object", "openapi-array", "static-content", "user-input", "xml"),
											},
										},

										"wildcard_order": schema.Int64Attribute{
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

							"response_pages": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"ajax_action_type": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("alert-popup", "custom", "redirect"),
											},
										},

										"ajax_custom_content": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ajax_enabled": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ajax_popup_message": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ajax_redirect_url": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"grpc_status_code": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`ABORTED|ALREADY_EXISTS|CANCELLED|DATA_LOSS|DEADLINE_EXCEEDED|FAILED_PRECONDITION|INTERNAL|INVALID_ARGUMENT|NOT_FOUND|OK|OUT_OF_RANGE|PERMISSION_DENIED|RESOURCE_EXHAUSTED|UNAUTHENTICATED|UNAVAILABLE|UNIMPLEMENTED|UNKNOWN|d+`), ""),
											},
										},

										"grpc_status_message": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"response_action_type": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("custom", "default", "erase-cookies", "redirect", "soap-fault"),
											},
										},

										"response_content": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"response_header": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"response_page_type": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("ajax", "ajax-login", "captcha", "captcha-fail", "default", "failed-login-honeypot", "failed-login-honeypot-ajax", "hijack", "leaked-credentials", "leaked-credentials-ajax", "mobile", "persistent-flow", "xml", "grpc"),
											},
										},

										"response_redirect_url": schema.StringAttribute{
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

							"response_page_reference": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"link": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^http`), ""),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"sensitive_parameters": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"dollaraction": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("delete"),
											},
										},

										"name": schema.StringAttribute{
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

							"sensitive_parameter_reference": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"link": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^http`), ""),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"server_technologies": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"dollaraction": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("delete"),
											},
										},

										"server_technology_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Jenkins", "SharePoint", "Oracle Application Server", "Python", "Oracle Identity Manager", "Spring Boot", "CouchDB", "SQLite", "Handlebars", "Mustache", "Prototype", "Zend", "Redis", "Underscore.js", "Ember.js", "ZURB Foundation", "ef.js", "Vue.js", "UIKit", "TYPO3 CMS", "RequireJS", "React", "MooTools", "Laravel", "GraphQL", "Google Web Toolkit", "Express.js", "CodeIgniter", "Backbone.js", "AngularJS", "JavaScript", "Nginx", "Jetty", "Joomla", "JavaServer Faces (JSF)", "Ruby", "MongoDB", "Django", "Node.js", "Citrix", "JBoss", "Elasticsearch", "Apache Struts", "XML", "PostgreSQL", "IBM DB2", "Sybase/ASE", "CGI", "Proxy Servers", "SSI (Server Side Includes)", "Cisco", "Novell", "Macromedia JRun", "BEA Systems WebLogic Server", "Lotus Domino", "MySQL", "Oracle", "Microsoft SQL Server", "PHP", "Outlook Web Access", "Apache/NCSA HTTP Server", "Apache Tomcat", "WordPress", "Macromedia ColdFusion", "Unix/Linux", "Microsoft Windows", "ASP.NET", "Front Page Server Extensions (FPSE)", "IIS", "WebDAV", "ASP", "Java Servlets/JSP", "jQuery"),
											},
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"server_technology_reference": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"link": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^http`), ""),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"signature_requirements": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"dollaraction": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("delete"),
											},
										},

										"tag": schema.StringAttribute{
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

							"signature_sets": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.MapType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"signature_settings": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"attack_signature_false_positive_mode": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("detect", "detect-and-allow", "disabled"),
										},
									},

									"minimum_accuracy_for_auto_added_signatures": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("high", "low", "medium"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"signature_reference": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"link": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^http`), ""),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"signature_set_reference": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"link": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^http`), ""),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"signature_setting_reference": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"link": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^http`), ""),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"signatures": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"enabled": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
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

										"signature_id": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tag": schema.StringAttribute{
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

							"software_version": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"template": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
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

							"threat_campaigns": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"is_enabled": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"threat_campaign_reference": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"link": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^http`), ""),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"url_reference": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"link": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^http`), ""),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"urls": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"dollaraction": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("delete"),
											},
										},

										"allow_rendering_in_frames": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("never", "only-same"),
											},
										},

										"allow_rendering_in_frames_only_from": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"attack_signatures_check": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"clickjacking_protection": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"description": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"disallow_file_upload_of_executables": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"html5_cross_origin_requests_enforcement": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"allow_origins_enforcement_mode": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("replace-with", "unmodified"),
													},
												},

												"check_allowed_methods": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"cross_domain_allowed_origin": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"include_sub_domains": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"origin_name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"origin_port": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"origin_protocol": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("http", "http/https", "https"),
																},
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"enforcement_mode": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("disabled", "enforce"),
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"is_allowed": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"mandatory_body": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"metachar_overrides": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"is_allowed": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"metachar": schema.StringAttribute{
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

										"metachars_on_url_check": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"method": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("ACL", "BCOPY", "BDELETE", "BMOVE", "BPROPFIND", "BPROPPATCH", "CHECKIN", "CHECKOUT", "CONNECT", "COPY", "DELETE", "GET", "HEAD", "LINK", "LOCK", "MERGE", "MKCOL", "MKWORKSPACE", "MOVE", "NOTIFY", "OPTIONS", "PATCH", "POLL", "POST", "PROPFIND", "PROPPATCH", "PUT", "REPORT", "RPC_IN_DATA", "RPC_OUT_DATA", "SEARCH", "SUBSCRIBE", "TRACE", "TRACK", "UNLINK", "UNLOCK", "UNSUBSCRIBE", "VERSION_CONTROL", "X-MS-ENUMATTS", "*"),
											},
										},

										"method_overrides": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"allowed": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"method": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("ACL", "BCOPY", "BDELETE", "BMOVE", "BPROPFIND", "BPROPPATCH", "CHECKIN", "CHECKOUT", "CONNECT", "COPY", "DELETE", "GET", "HEAD", "LINK", "LOCK", "MERGE", "MKCOL", "MKWORKSPACE", "MOVE", "NOTIFY", "OPTIONS", "PATCH", "POLL", "POST", "PROPFIND", "PROPPATCH", "PUT", "REPORT", "RPC_IN_DATA", "RPC_OUT_DATA", "SEARCH", "SUBSCRIBE", "TRACE", "TRACK", "UNLINK", "UNLOCK", "UNSUBSCRIBE", "VERSION_CONTROL", "X-MS-ENUMATTS"),
														},
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"methods_override_on_url_check": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
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

										"operation_id": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"positional_parameters": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"parameter": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"dollaraction": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("delete"),
																},
															},

															"allow_empty_value": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"allow_repeated_parameter_name": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"array_serialization_format": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("csv", "form", "label", "matrix", "multi", "multipart", "pipe", "ssv", "tsv"),
																},
															},

															"attack_signatures_check": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"check_max_value": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"check_max_value_length": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"check_metachars": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"check_min_value": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"check_min_value_length": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"check_multiple_of_value": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"content_profile": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
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

															"data_type": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("alpha-numeric", "binary", "boolean", "decimal", "email", "integer", "none", "phone"),
																},
															},

															"decode_value_as_base64": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("enabled", "disabled", "required"),
																},
															},

															"disallow_file_upload_of_executables": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"enable_regular_expression": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"exclusive_max": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"exclusive_min": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"is_base64": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"is_cookie": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"is_header": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"level": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("global", "url"),
																},
															},

															"mandatory": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"maximum_length": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"maximum_value": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"metachars_on_parameter_value_check": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"minimum_length": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"minimum_value": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"multiple_of": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
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

															"name_metachar_overrides": schema.ListNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"is_allowed": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"metachar": schema.StringAttribute{
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

															"object_serialization_style": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"parameter_enum_values": schema.ListAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"parameter_location": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("any", "cookie", "form-data", "header", "path", "query"),
																},
															},

															"regular_expression": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"sensitive_parameter": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"signature_overrides": schema.ListNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"enabled": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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

																		"signature_id": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"tag": schema.StringAttribute{
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

															"static_values": schema.StringAttribute{
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
																Validators: []validator.String{
																	stringvalidator.OneOf("explicit", "wildcard"),
																},
															},

															"url": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"method": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("ACL", "BCOPY", "BDELETE", "BMOVE", "BPROPFIND", "BPROPPATCH", "CHECKIN", "CHECKOUT", "CONNECT", "COPY", "DELETE", "GET", "HEAD", "LINK", "LOCK", "MERGE", "MKCOL", "MKWORKSPACE", "MOVE", "NOTIFY", "OPTIONS", "PATCH", "POLL", "POST", "PROPFIND", "PROPPATCH", "PUT", "REPORT", "RPC_IN_DATA", "RPC_OUT_DATA", "SEARCH", "SUBSCRIBE", "TRACE", "TRACK", "UNLINK", "UNLOCK", "UNSUBSCRIBE", "VERSION_CONTROL", "X-MS-ENUMATTS", "*"),
																		},
																	},

																	"name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"protocol": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("http", "https"),
																		},
																	},

																	"type": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("explicit", "wildcard"),
																		},
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"value_metachar_overrides": schema.ListNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"is_allowed": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"metachar": schema.StringAttribute{
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

															"value_type": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("array", "auto-detect", "dynamic-content", "dynamic-parameter-name", "ignore", "json", "object", "openapi-array", "static-content", "user-input", "xml"),
																},
															},

															"wildcard_order": schema.Int64Attribute{
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

													"url_segment_index": schema.Int64Attribute{
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

										"protocol": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("http", "https"),
											},
										},

										"signature_overrides": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"enabled": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
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

													"signature_id": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"tag": schema.StringAttribute{
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

										"type": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("explicit", "wildcard"),
											},
										},

										"url_content_profiles": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"content_profile": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
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

													"header_name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"header_order": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"header_value": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
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

													"type": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("apply-content-signatures", "apply-value-and-content-signatures", "disallow", "do-nothing", "form-data", "gwt", "json", "xml", "grpc"),
														},
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"wildcard_order": schema.Int64Attribute{
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

							"whitelist_ips": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"dollaraction": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("delete"),
											},
										},

										"block_requests": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("always", "never", "policy-default"),
											},
										},

										"ip_address": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}`), ""),
											},
										},

										"ip_mask": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}`), ""),
											},
										},

										"never_log_requests": schema.BoolAttribute{
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

							"whitelist_ip_reference": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"link": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^http`), ""),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"xml_profiles": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"dollaraction": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("delete"),
											},
										},

										"attack_signatures_check": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"defense_attributes": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"allow_cdata": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"allow_dt_ds": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"allow_external_references": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"allow_processing_instructions": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"maximum_attribute_value_length": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"maximum_attributes_per_element": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"maximum_children_per_element": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"maximum_document_depth": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"maximum_document_size": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"maximum_elements": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"maximum_ns_declarations": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"maximum_name_length": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"maximum_namespace_length": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"tolerate_close_tag_shorthand": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"tolerate_leading_white_space": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"tolerate_numeric_names": schema.BoolAttribute{
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

										"description": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"enable_wss": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"follow_schema_links": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
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

										"signature_overrides": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"enabled": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
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

													"signature_id": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"tag": schema.StringAttribute{
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

										"use_xml_response_page": schema.BoolAttribute{
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

							"xml_validation_files": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"dollaraction": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("delete"),
											},
										},

										"contents": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"file_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"is_base64": schema.BoolAttribute{
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

							"xml_profile_reference": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"link": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^http`), ""),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"xml_validation_file_reference": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"link": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^http`), ""),
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
		},
	}
}

func (r *AppprotectF5ComAppolicyV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_appprotect_f5_com_ap_policy_v1beta1_manifest")

	var model AppprotectF5ComAppolicyV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("appprotect.f5.com/v1beta1")
	model.Kind = pointer.String("APPolicy")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
