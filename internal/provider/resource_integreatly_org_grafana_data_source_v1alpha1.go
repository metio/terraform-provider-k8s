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

type IntegreatlyOrgGrafanaDataSourceV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*IntegreatlyOrgGrafanaDataSourceV1Alpha1Resource)(nil)
)

type IntegreatlyOrgGrafanaDataSourceV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type IntegreatlyOrgGrafanaDataSourceV1Alpha1GoModel struct {
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
		Datasources *[]struct {
			Access *string `tfsdk:"access" yaml:"access,omitempty"`

			BasicAuth *bool `tfsdk:"basic_auth" yaml:"basicAuth,omitempty"`

			BasicAuthPassword *string `tfsdk:"basic_auth_password" yaml:"basicAuthPassword,omitempty"`

			BasicAuthUser *string `tfsdk:"basic_auth_user" yaml:"basicAuthUser,omitempty"`

			CustomJsonData utilities.Dynamic `tfsdk:"custom_json_data" yaml:"customJsonData,omitempty"`

			CustomSecureJsonData utilities.Dynamic `tfsdk:"custom_secure_json_data" yaml:"customSecureJsonData,omitempty"`

			Database *string `tfsdk:"database" yaml:"database,omitempty"`

			Editable *bool `tfsdk:"editable" yaml:"editable,omitempty"`

			IsDefault *bool `tfsdk:"is_default" yaml:"isDefault,omitempty"`

			JsonData *struct {
				AddCorsHeader *bool `tfsdk:"add_cors_header" yaml:"addCorsHeader,omitempty"`

				AlertmanagerUid *string `tfsdk:"alertmanager_uid" yaml:"alertmanagerUid,omitempty"`

				AllowInfraExplore *bool `tfsdk:"allow_infra_explore" yaml:"allowInfraExplore,omitempty"`

				ApiToken *string `tfsdk:"api_token" yaml:"apiToken,omitempty"`

				AppInsightsAppId *string `tfsdk:"app_insights_app_id" yaml:"appInsightsAppId,omitempty"`

				AssumeRoleArn *string `tfsdk:"assume_role_arn" yaml:"assumeRoleArn,omitempty"`

				AuthType *string `tfsdk:"auth_type" yaml:"authType,omitempty"`

				AuthenticationType *string `tfsdk:"authentication_type" yaml:"authenticationType,omitempty"`

				AzureLogAnalyticsSameAs *string `tfsdk:"azure_log_analytics_same_as" yaml:"azureLogAnalyticsSameAs,omitempty"`

				ClientEmail *string `tfsdk:"client_email" yaml:"clientEmail,omitempty"`

				ClientId *string `tfsdk:"client_id" yaml:"clientId,omitempty"`

				CloudName *string `tfsdk:"cloud_name" yaml:"cloudName,omitempty"`

				ClusterUrl *string `tfsdk:"cluster_url" yaml:"clusterUrl,omitempty"`

				ConnMaxLifetime *int64 `tfsdk:"conn_max_lifetime" yaml:"connMaxLifetime,omitempty"`

				CustomMetricsNamespaces *string `tfsdk:"custom_metrics_namespaces" yaml:"customMetricsNamespaces,omitempty"`

				CustomQueryParameters *string `tfsdk:"custom_query_parameters" yaml:"customQueryParameters,omitempty"`

				DefaultBucket *string `tfsdk:"default_bucket" yaml:"defaultBucket,omitempty"`

				DefaultDatabase *string `tfsdk:"default_database" yaml:"defaultDatabase,omitempty"`

				DefaultProject *string `tfsdk:"default_project" yaml:"defaultProject,omitempty"`

				DefaultRegion *string `tfsdk:"default_region" yaml:"defaultRegion,omitempty"`

				DerivedFields *[]struct {
					DatasourceUid *string `tfsdk:"datasource_uid" yaml:"datasourceUid,omitempty"`

					MatcherRegex *string `tfsdk:"matcher_regex" yaml:"matcherRegex,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Url *string `tfsdk:"url" yaml:"url,omitempty"`
				} `tfsdk:"derived_fields" yaml:"derivedFields,omitempty"`

				Encrypt *string `tfsdk:"encrypt" yaml:"encrypt,omitempty"`

				EsVersion *string `tfsdk:"es_version" yaml:"esVersion,omitempty"`

				ExemplarTraceIdDestinations *[]struct {
					DatasourceUid *string `tfsdk:"datasource_uid" yaml:"datasourceUid,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Url *string `tfsdk:"url" yaml:"url,omitempty"`

					UrlDisplayLabel *string `tfsdk:"url_display_label" yaml:"urlDisplayLabel,omitempty"`
				} `tfsdk:"exemplar_trace_id_destinations" yaml:"exemplarTraceIdDestinations,omitempty"`

				GithubUrl *string `tfsdk:"github_url" yaml:"githubUrl,omitempty"`

				GraphiteVersion *string `tfsdk:"graphite_version" yaml:"graphiteVersion,omitempty"`

				HttpHeaderName1 *string `tfsdk:"http_header_name1" yaml:"httpHeaderName1,omitempty"`

				HttpHeaderName2 *string `tfsdk:"http_header_name2" yaml:"httpHeaderName2,omitempty"`

				HttpHeaderName3 *string `tfsdk:"http_header_name3" yaml:"httpHeaderName3,omitempty"`

				HttpHeaderName4 *string `tfsdk:"http_header_name4" yaml:"httpHeaderName4,omitempty"`

				HttpHeaderName5 *string `tfsdk:"http_header_name5" yaml:"httpHeaderName5,omitempty"`

				HttpHeaderName6 *string `tfsdk:"http_header_name6" yaml:"httpHeaderName6,omitempty"`

				HttpHeaderName7 *string `tfsdk:"http_header_name7" yaml:"httpHeaderName7,omitempty"`

				HttpHeaderName8 *string `tfsdk:"http_header_name8" yaml:"httpHeaderName8,omitempty"`

				HttpHeaderName9 *string `tfsdk:"http_header_name9" yaml:"httpHeaderName9,omitempty"`

				HttpMethod *string `tfsdk:"http_method" yaml:"httpMethod,omitempty"`

				HttpMode *string `tfsdk:"http_mode" yaml:"httpMode,omitempty"`

				Implementation *string `tfsdk:"implementation" yaml:"implementation,omitempty"`

				Interval *string `tfsdk:"interval" yaml:"interval,omitempty"`

				LogAnalyticsClientId *string `tfsdk:"log_analytics_client_id" yaml:"logAnalyticsClientId,omitempty"`

				LogAnalyticsDefaultWorkspace *string `tfsdk:"log_analytics_default_workspace" yaml:"logAnalyticsDefaultWorkspace,omitempty"`

				LogAnalyticsSubscriptionId *string `tfsdk:"log_analytics_subscription_id" yaml:"logAnalyticsSubscriptionId,omitempty"`

				LogAnalyticsTenantId *string `tfsdk:"log_analytics_tenant_id" yaml:"logAnalyticsTenantId,omitempty"`

				LogLevelField *string `tfsdk:"log_level_field" yaml:"logLevelField,omitempty"`

				LogMessageField *string `tfsdk:"log_message_field" yaml:"logMessageField,omitempty"`

				ManageAlerts *bool `tfsdk:"manage_alerts" yaml:"manageAlerts,omitempty"`

				MaxIdleConns *int64 `tfsdk:"max_idle_conns" yaml:"maxIdleConns,omitempty"`

				MaxLines *int64 `tfsdk:"max_lines" yaml:"maxLines,omitempty"`

				MaxOpenConns *int64 `tfsdk:"max_open_conns" yaml:"maxOpenConns,omitempty"`

				NodeGraph *struct {
					Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`
				} `tfsdk:"node_graph" yaml:"nodeGraph,omitempty"`

				OauthPassThru *bool `tfsdk:"oauth_pass_thru" yaml:"oauthPassThru,omitempty"`

				Organization *string `tfsdk:"organization" yaml:"organization,omitempty"`

				Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

				PostgresVersion *int64 `tfsdk:"postgres_version" yaml:"postgresVersion,omitempty"`

				QueryTimeout *string `tfsdk:"query_timeout" yaml:"queryTimeout,omitempty"`

				Search *struct {
					Hide *bool `tfsdk:"hide" yaml:"hide,omitempty"`
				} `tfsdk:"search" yaml:"search,omitempty"`

				Server *string `tfsdk:"server" yaml:"server,omitempty"`

				ServiceMap *struct {
					DatasourceUid *string `tfsdk:"datasource_uid" yaml:"datasourceUid,omitempty"`
				} `tfsdk:"service_map" yaml:"serviceMap,omitempty"`

				ShowOffline *bool `tfsdk:"show_offline" yaml:"showOffline,omitempty"`

				SigV4AssumeRoleArn *string `tfsdk:"sig_v4_assume_role_arn" yaml:"sigV4AssumeRoleArn,omitempty"`

				SigV4Auth *bool `tfsdk:"sig_v4_auth" yaml:"sigV4Auth,omitempty"`

				SigV4AuthType *string `tfsdk:"sig_v4_auth_type" yaml:"sigV4AuthType,omitempty"`

				SigV4ExternalId *string `tfsdk:"sig_v4_external_id" yaml:"sigV4ExternalId,omitempty"`

				SigV4Profile *string `tfsdk:"sig_v4_profile" yaml:"sigV4Profile,omitempty"`

				SigV4Region *string `tfsdk:"sig_v4_region" yaml:"sigV4Region,omitempty"`

				Sslmode *string `tfsdk:"sslmode" yaml:"sslmode,omitempty"`

				SubscriptionId *string `tfsdk:"subscription_id" yaml:"subscriptionId,omitempty"`

				TenantId *string `tfsdk:"tenant_id" yaml:"tenantId,omitempty"`

				TimeField *string `tfsdk:"time_field" yaml:"timeField,omitempty"`

				TimeInterval *string `tfsdk:"time_interval" yaml:"timeInterval,omitempty"`

				Timeout *int64 `tfsdk:"timeout" yaml:"timeout,omitempty"`

				Timescaledb *bool `tfsdk:"timescaledb" yaml:"timescaledb,omitempty"`

				Timezone *string `tfsdk:"timezone" yaml:"timezone,omitempty"`

				TlsAuth *bool `tfsdk:"tls_auth" yaml:"tlsAuth,omitempty"`

				TlsAuthWithCACert *bool `tfsdk:"tls_auth_with_ca_cert" yaml:"tlsAuthWithCACert,omitempty"`

				TlsSkipVerify *bool `tfsdk:"tls_skip_verify" yaml:"tlsSkipVerify,omitempty"`

				TokenUri *string `tfsdk:"token_uri" yaml:"tokenUri,omitempty"`

				TracesToLogs *struct {
					DatasourceUid *string `tfsdk:"datasource_uid" yaml:"datasourceUid,omitempty"`

					FilterBySpanID *bool `tfsdk:"filter_by_span_id" yaml:"filterBySpanID,omitempty"`

					FilterByTraceID *bool `tfsdk:"filter_by_trace_id" yaml:"filterByTraceID,omitempty"`

					LokiSearch *bool `tfsdk:"loki_search" yaml:"lokiSearch,omitempty"`

					SpanEndTimeShift *string `tfsdk:"span_end_time_shift" yaml:"spanEndTimeShift,omitempty"`

					SpanStartTimeShift *string `tfsdk:"span_start_time_shift" yaml:"spanStartTimeShift,omitempty"`

					Tags *[]string `tfsdk:"tags" yaml:"tags,omitempty"`
				} `tfsdk:"traces_to_logs" yaml:"tracesToLogs,omitempty"`

				TsdbResolution *string `tfsdk:"tsdb_resolution" yaml:"tsdbResolution,omitempty"`

				TsdbVersion *string `tfsdk:"tsdb_version" yaml:"tsdbVersion,omitempty"`

				Url *string `tfsdk:"url" yaml:"url,omitempty"`

				UsePOST *bool `tfsdk:"use_post" yaml:"usePOST,omitempty"`

				UseProxy *bool `tfsdk:"use_proxy" yaml:"useProxy,omitempty"`

				UseYandexCloudAuthorization *bool `tfsdk:"use_yandex_cloud_authorization" yaml:"useYandexCloudAuthorization,omitempty"`

				Username *string `tfsdk:"username" yaml:"username,omitempty"`

				Version *string `tfsdk:"version" yaml:"version,omitempty"`

				XHeaderKey *string `tfsdk:"x_header_key" yaml:"xHeaderKey,omitempty"`

				XHeaderUser *string `tfsdk:"x_header_user" yaml:"xHeaderUser,omitempty"`
			} `tfsdk:"json_data" yaml:"jsonData,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			OrgId *int64 `tfsdk:"org_id" yaml:"orgId,omitempty"`

			Password *string `tfsdk:"password" yaml:"password,omitempty"`

			SecureJsonData *struct {
				AccessKey *string `tfsdk:"access_key" yaml:"accessKey,omitempty"`

				AccessToken *string `tfsdk:"access_token" yaml:"accessToken,omitempty"`

				AppInsightsApiKey *string `tfsdk:"app_insights_api_key" yaml:"appInsightsApiKey,omitempty"`

				BasicAuthPassword *string `tfsdk:"basic_auth_password" yaml:"basicAuthPassword,omitempty"`

				ClientSecret *string `tfsdk:"client_secret" yaml:"clientSecret,omitempty"`

				HttpHeaderValue1 *string `tfsdk:"http_header_value1" yaml:"httpHeaderValue1,omitempty"`

				HttpHeaderValue2 *string `tfsdk:"http_header_value2" yaml:"httpHeaderValue2,omitempty"`

				HttpHeaderValue3 *string `tfsdk:"http_header_value3" yaml:"httpHeaderValue3,omitempty"`

				HttpHeaderValue4 *string `tfsdk:"http_header_value4" yaml:"httpHeaderValue4,omitempty"`

				HttpHeaderValue5 *string `tfsdk:"http_header_value5" yaml:"httpHeaderValue5,omitempty"`

				HttpHeaderValue6 *string `tfsdk:"http_header_value6" yaml:"httpHeaderValue6,omitempty"`

				HttpHeaderValue7 *string `tfsdk:"http_header_value7" yaml:"httpHeaderValue7,omitempty"`

				HttpHeaderValue8 *string `tfsdk:"http_header_value8" yaml:"httpHeaderValue8,omitempty"`

				HttpHeaderValue9 *string `tfsdk:"http_header_value9" yaml:"httpHeaderValue9,omitempty"`

				LogAnalyticsClientSecret *string `tfsdk:"log_analytics_client_secret" yaml:"logAnalyticsClientSecret,omitempty"`

				Password *string `tfsdk:"password" yaml:"password,omitempty"`

				PrivateKey *string `tfsdk:"private_key" yaml:"privateKey,omitempty"`

				SecretKey *string `tfsdk:"secret_key" yaml:"secretKey,omitempty"`

				SigV4AccessKey *string `tfsdk:"sig_v4_access_key" yaml:"sigV4AccessKey,omitempty"`

				SigV4SecretKey *string `tfsdk:"sig_v4_secret_key" yaml:"sigV4SecretKey,omitempty"`

				TlsCACert *string `tfsdk:"tls_ca_cert" yaml:"tlsCACert,omitempty"`

				TlsClientCert *string `tfsdk:"tls_client_cert" yaml:"tlsClientCert,omitempty"`

				TlsClientKey *string `tfsdk:"tls_client_key" yaml:"tlsClientKey,omitempty"`

				Token *string `tfsdk:"token" yaml:"token,omitempty"`
			} `tfsdk:"secure_json_data" yaml:"secureJsonData,omitempty"`

			Type *string `tfsdk:"type" yaml:"type,omitempty"`

			Uid *string `tfsdk:"uid" yaml:"uid,omitempty"`

			Url *string `tfsdk:"url" yaml:"url,omitempty"`

			User *string `tfsdk:"user" yaml:"user,omitempty"`

			Version *int64 `tfsdk:"version" yaml:"version,omitempty"`

			WithCredentials *bool `tfsdk:"with_credentials" yaml:"withCredentials,omitempty"`
		} `tfsdk:"datasources" yaml:"datasources,omitempty"`

		Name *string `tfsdk:"name" yaml:"name,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewIntegreatlyOrgGrafanaDataSourceV1Alpha1Resource() resource.Resource {
	return &IntegreatlyOrgGrafanaDataSourceV1Alpha1Resource{}
}

func (r *IntegreatlyOrgGrafanaDataSourceV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_integreatly_org_grafana_data_source_v1alpha1"
}

func (r *IntegreatlyOrgGrafanaDataSourceV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "GrafanaDataSource is the Schema for the grafanadatasources API",
		MarkdownDescription: "GrafanaDataSource is the Schema for the grafanadatasources API",
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
				Description:         "GrafanaDataSourceSpec defines the desired state of GrafanaDataSource",
				MarkdownDescription: "GrafanaDataSourceSpec defines the desired state of GrafanaDataSource",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"datasources": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"access": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"basic_auth": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"basic_auth_password": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"basic_auth_user": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"custom_json_data": {
								Description:         "CustomJsonData will be used in place of jsonData, if present, and supports arbitrary JSON, not just those of official datasources",
								MarkdownDescription: "CustomJsonData will be used in place of jsonData, if present, and supports arbitrary JSON, not just those of official datasources",

								Type: utilities.DynamicType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"custom_secure_json_data": {
								Description:         "SecureCustomJsonData will be used in place of secureJsonData, if present, and supports arbitrary JSON, not just those of official datasources",
								MarkdownDescription: "SecureCustomJsonData will be used in place of secureJsonData, if present, and supports arbitrary JSON, not just those of official datasources",

								Type: utilities.DynamicType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"database": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"editable": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"is_default": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"json_data": {
								Description:         "GrafanaDataSourceJsonData contains the most common json options See https://grafana.com/docs/administration/provisioning/#datasources",
								MarkdownDescription: "GrafanaDataSourceJsonData contains the most common json options See https://grafana.com/docs/administration/provisioning/#datasources",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"add_cors_header": {
										Description:         "Useful fields for clickhouse datasource See https://github.com/Vertamedia/clickhouse-grafana/tree/master/dist/README.md#configure-the-datasource-with-provisioning See https://github.com/Vertamedia/clickhouse-grafana/tree/master/src/datasource.ts#L44",
										MarkdownDescription: "Useful fields for clickhouse datasource See https://github.com/Vertamedia/clickhouse-grafana/tree/master/dist/README.md#configure-the-datasource-with-provisioning See https://github.com/Vertamedia/clickhouse-grafana/tree/master/src/datasource.ts#L44",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"alertmanager_uid": {
										Description:         "AlertManagerUID if null use the internal grafana alertmanager",
										MarkdownDescription: "AlertManagerUID if null use the internal grafana alertmanager",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"allow_infra_explore": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"api_token": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"app_insights_app_id": {
										Description:         "Fields for Azure data sources",
										MarkdownDescription: "Fields for Azure data sources",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"assume_role_arn": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"auth_type": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"authentication_type": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"azure_log_analytics_same_as": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"client_email": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"client_id": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"cloud_name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"cluster_url": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"conn_max_lifetime": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"custom_metrics_namespaces": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"custom_query_parameters": {
										Description:         "Fields for Prometheus data sources",
										MarkdownDescription: "Fields for Prometheus data sources",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"default_bucket": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"default_database": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"default_project": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"default_region": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"derived_fields": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"datasource_uid": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"matcher_regex": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

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

									"encrypt": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"es_version": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"exemplar_trace_id_destinations": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"datasource_uid": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

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

											"url": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"url_display_label": {
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

									"github_url": {
										Description:         "Fields for Github data sources",
										MarkdownDescription: "Fields for Github data sources",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"graphite_version": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"http_header_name1": {
										Description:         "Custom HTTP headers for datasources See https://grafana.com/docs/grafana/latest/administration/provisioning/#datasources",
										MarkdownDescription: "Custom HTTP headers for datasources See https://grafana.com/docs/grafana/latest/administration/provisioning/#datasources",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"http_header_name2": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"http_header_name3": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"http_header_name4": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"http_header_name5": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"http_header_name6": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"http_header_name7": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"http_header_name8": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"http_header_name9": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"http_method": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"http_mode": {
										Description:         "Fields for InfluxDB data sources",
										MarkdownDescription: "Fields for InfluxDB data sources",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"implementation": {
										Description:         "Fields for Alertmanager data sources",
										MarkdownDescription: "Fields for Alertmanager data sources",

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

									"log_analytics_client_id": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"log_analytics_default_workspace": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"log_analytics_subscription_id": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"log_analytics_tenant_id": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"log_level_field": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"log_message_field": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"manage_alerts": {
										Description:         "ManageAlerts turns on alert management from UI",
										MarkdownDescription: "ManageAlerts turns on alert management from UI",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"max_idle_conns": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"max_lines": {
										Description:         "Fields for Loki data sources",
										MarkdownDescription: "Fields for Loki data sources",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"max_open_conns": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"node_graph": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"enabled": {
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

									"oauth_pass_thru": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"organization": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

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

									"postgres_version": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"query_timeout": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"search": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"hide": {
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

									"server": {
										Description:         "Fields for Grafana Clickhouse data sources",
										MarkdownDescription: "Fields for Grafana Clickhouse data sources",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"service_map": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"datasource_uid": {
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

									"show_offline": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"sig_v4_assume_role_arn": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"sig_v4_auth": {
										Description:         "Fields for AWS Prometheus data sources",
										MarkdownDescription: "Fields for AWS Prometheus data sources",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"sig_v4_auth_type": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"sig_v4_external_id": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"sig_v4_profile": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"sig_v4_region": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"sslmode": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"subscription_id": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tenant_id": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"time_field": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"time_interval": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"timeout": {
										Description:         "HTTP Request timeout in seconds. Overrides dataproxy.timeout option",
										MarkdownDescription: "HTTP Request timeout in seconds. Overrides dataproxy.timeout option",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"timescaledb": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"timezone": {
										Description:         "Extra field for MySQL data source",
										MarkdownDescription: "Extra field for MySQL data source",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls_auth": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls_auth_with_ca_cert": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls_skip_verify": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"token_uri": {
										Description:         "Fields for Stackdriver data sources",
										MarkdownDescription: "Fields for Stackdriver data sources",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"traces_to_logs": {
										Description:         "Fields for tracing data sources",
										MarkdownDescription: "Fields for tracing data sources",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"datasource_uid": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"filter_by_span_id": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"filter_by_trace_id": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"loki_search": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"span_end_time_shift": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"span_start_time_shift": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"tags": {
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

									"tsdb_resolution": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tsdb_version": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"url": {
										Description:         "Fields for Instana data sources See https://github.com/instana/instana-grafana-datasource/blob/main/provisioning/datasources/datasource.yml",
										MarkdownDescription: "Fields for Instana data sources See https://github.com/instana/instana-grafana-datasource/blob/main/provisioning/datasources/datasource.yml",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"use_post": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"use_proxy": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"use_yandex_cloud_authorization": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"username": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"version": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"x_header_key": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"x_header_user": {
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

							"name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"org_id": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"password": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"secure_json_data": {
								Description:         "GrafanaDataSourceSecureJsonData contains the most common secure json options See https://grafana.com/docs/administration/provisioning/#datasources",
								MarkdownDescription: "GrafanaDataSourceSecureJsonData contains the most common secure json options See https://grafana.com/docs/administration/provisioning/#datasources",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"access_key": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"access_token": {
										Description:         "Fields for Github data sources",
										MarkdownDescription: "Fields for Github data sources",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"app_insights_api_key": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"basic_auth_password": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"client_secret": {
										Description:         "Fields for Azure data sources",
										MarkdownDescription: "Fields for Azure data sources",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"http_header_value1": {
										Description:         "Custom HTTP headers for datasources See https://grafana.com/docs/grafana/latest/administration/provisioning/#datasources",
										MarkdownDescription: "Custom HTTP headers for datasources See https://grafana.com/docs/grafana/latest/administration/provisioning/#datasources",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"http_header_value2": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"http_header_value3": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"http_header_value4": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"http_header_value5": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"http_header_value6": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"http_header_value7": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"http_header_value8": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"http_header_value9": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"log_analytics_client_secret": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"password": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"private_key": {
										Description:         "Fields for Stackdriver data sources",
										MarkdownDescription: "Fields for Stackdriver data sources",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret_key": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"sig_v4_access_key": {
										Description:         "Fields for AWS data sources",
										MarkdownDescription: "Fields for AWS data sources",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"sig_v4_secret_key": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls_ca_cert": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls_client_cert": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls_client_key": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"token": {
										Description:         "Fields for InfluxDB data sources",
										MarkdownDescription: "Fields for InfluxDB data sources",

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

							"type": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"uid": {
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

							"user": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"version": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"with_credentials": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"name": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: true,
						Optional: false,
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

func (r *IntegreatlyOrgGrafanaDataSourceV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_integreatly_org_grafana_data_source_v1alpha1")

	var state IntegreatlyOrgGrafanaDataSourceV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel IntegreatlyOrgGrafanaDataSourceV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("integreatly.org/v1alpha1")
	goModel.Kind = utilities.Ptr("GrafanaDataSource")

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

func (r *IntegreatlyOrgGrafanaDataSourceV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_integreatly_org_grafana_data_source_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *IntegreatlyOrgGrafanaDataSourceV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_integreatly_org_grafana_data_source_v1alpha1")

	var state IntegreatlyOrgGrafanaDataSourceV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel IntegreatlyOrgGrafanaDataSourceV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("integreatly.org/v1alpha1")
	goModel.Kind = utilities.Ptr("GrafanaDataSource")

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

func (r *IntegreatlyOrgGrafanaDataSourceV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_integreatly_org_grafana_data_source_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
