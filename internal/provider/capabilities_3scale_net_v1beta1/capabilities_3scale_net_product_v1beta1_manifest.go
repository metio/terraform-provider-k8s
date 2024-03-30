/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package capabilities_3scale_net_v1beta1

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
	_ datasource.DataSource = &Capabilities3ScaleNetProductV1Beta1Manifest{}
)

func NewCapabilities3ScaleNetProductV1Beta1Manifest() datasource.DataSource {
	return &Capabilities3ScaleNetProductV1Beta1Manifest{}
}

type Capabilities3ScaleNetProductV1Beta1Manifest struct{}

type Capabilities3ScaleNetProductV1Beta1ManifestData struct {
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
		ApplicationPlans *struct {
			AppsRequireApproval *bool   `tfsdk:"apps_require_approval" json:"appsRequireApproval,omitempty"`
			CostMonth           *string `tfsdk:"cost_month" json:"costMonth,omitempty"`
			Limits              *[]struct {
				MetricMethodRef *struct {
					Backend    *string `tfsdk:"backend" json:"backend,omitempty"`
					SystemName *string `tfsdk:"system_name" json:"systemName,omitempty"`
				} `tfsdk:"metric_method_ref" json:"metricMethodRef,omitempty"`
				Period *string `tfsdk:"period" json:"period,omitempty"`
				Value  *int64  `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"limits" json:"limits,omitempty"`
			Name         *string `tfsdk:"name" json:"name,omitempty"`
			PricingRules *[]struct {
				From            *int64 `tfsdk:"from" json:"from,omitempty"`
				MetricMethodRef *struct {
					Backend    *string `tfsdk:"backend" json:"backend,omitempty"`
					SystemName *string `tfsdk:"system_name" json:"systemName,omitempty"`
				} `tfsdk:"metric_method_ref" json:"metricMethodRef,omitempty"`
				PricePerUnit *string `tfsdk:"price_per_unit" json:"pricePerUnit,omitempty"`
				To           *int64  `tfsdk:"to" json:"to,omitempty"`
			} `tfsdk:"pricing_rules" json:"pricingRules,omitempty"`
			Published   *bool   `tfsdk:"published" json:"published,omitempty"`
			SetupFee    *string `tfsdk:"setup_fee" json:"setupFee,omitempty"`
			TrialPeriod *int64  `tfsdk:"trial_period" json:"trialPeriod,omitempty"`
		} `tfsdk:"application_plans" json:"applicationPlans,omitempty"`
		BackendUsages *struct {
			Path *string `tfsdk:"path" json:"path,omitempty"`
		} `tfsdk:"backend_usages" json:"backendUsages,omitempty"`
		Deployment *struct {
			ApicastHosted *struct {
				Authentication *struct {
					AppKeyAppID *struct {
						AppID           *string `tfsdk:"app_id" json:"appID,omitempty"`
						AppKey          *string `tfsdk:"app_key" json:"appKey,omitempty"`
						Credentials     *string `tfsdk:"credentials" json:"credentials,omitempty"`
						GatewayResponse *struct {
							ErrorAuthFailed            *string `tfsdk:"error_auth_failed" json:"errorAuthFailed,omitempty"`
							ErrorAuthMissing           *string `tfsdk:"error_auth_missing" json:"errorAuthMissing,omitempty"`
							ErrorHeadersAuthFailed     *string `tfsdk:"error_headers_auth_failed" json:"errorHeadersAuthFailed,omitempty"`
							ErrorHeadersAuthMissing    *string `tfsdk:"error_headers_auth_missing" json:"errorHeadersAuthMissing,omitempty"`
							ErrorHeadersLimitsExceeded *string `tfsdk:"error_headers_limits_exceeded" json:"errorHeadersLimitsExceeded,omitempty"`
							ErrorHeadersNoMatch        *string `tfsdk:"error_headers_no_match" json:"errorHeadersNoMatch,omitempty"`
							ErrorLimitsExceeded        *string `tfsdk:"error_limits_exceeded" json:"errorLimitsExceeded,omitempty"`
							ErrorNoMatch               *string `tfsdk:"error_no_match" json:"errorNoMatch,omitempty"`
							ErrorStatusAuthFailed      *int64  `tfsdk:"error_status_auth_failed" json:"errorStatusAuthFailed,omitempty"`
							ErrorStatusAuthMissing     *int64  `tfsdk:"error_status_auth_missing" json:"errorStatusAuthMissing,omitempty"`
							ErrorStatusLimitsExceeded  *int64  `tfsdk:"error_status_limits_exceeded" json:"errorStatusLimitsExceeded,omitempty"`
							ErrorStatusNoMatch         *int64  `tfsdk:"error_status_no_match" json:"errorStatusNoMatch,omitempty"`
						} `tfsdk:"gateway_response" json:"gatewayResponse,omitempty"`
						Security *struct {
							HostHeader  *string `tfsdk:"host_header" json:"hostHeader,omitempty"`
							SecretToken *string `tfsdk:"secret_token" json:"secretToken,omitempty"`
						} `tfsdk:"security" json:"security,omitempty"`
					} `tfsdk:"app_key_app_id" json:"appKeyAppID,omitempty"`
					Oidc *struct {
						AuthenticationFlow *struct {
							DirectAccessGrantsEnabled *bool `tfsdk:"direct_access_grants_enabled" json:"directAccessGrantsEnabled,omitempty"`
							ImplicitFlowEnabled       *bool `tfsdk:"implicit_flow_enabled" json:"implicitFlowEnabled,omitempty"`
							ServiceAccountsEnabled    *bool `tfsdk:"service_accounts_enabled" json:"serviceAccountsEnabled,omitempty"`
							StandardFlowEnabled       *bool `tfsdk:"standard_flow_enabled" json:"standardFlowEnabled,omitempty"`
						} `tfsdk:"authentication_flow" json:"authenticationFlow,omitempty"`
						Credentials     *string `tfsdk:"credentials" json:"credentials,omitempty"`
						GatewayResponse *struct {
							ErrorAuthFailed            *string `tfsdk:"error_auth_failed" json:"errorAuthFailed,omitempty"`
							ErrorAuthMissing           *string `tfsdk:"error_auth_missing" json:"errorAuthMissing,omitempty"`
							ErrorHeadersAuthFailed     *string `tfsdk:"error_headers_auth_failed" json:"errorHeadersAuthFailed,omitempty"`
							ErrorHeadersAuthMissing    *string `tfsdk:"error_headers_auth_missing" json:"errorHeadersAuthMissing,omitempty"`
							ErrorHeadersLimitsExceeded *string `tfsdk:"error_headers_limits_exceeded" json:"errorHeadersLimitsExceeded,omitempty"`
							ErrorHeadersNoMatch        *string `tfsdk:"error_headers_no_match" json:"errorHeadersNoMatch,omitempty"`
							ErrorLimitsExceeded        *string `tfsdk:"error_limits_exceeded" json:"errorLimitsExceeded,omitempty"`
							ErrorNoMatch               *string `tfsdk:"error_no_match" json:"errorNoMatch,omitempty"`
							ErrorStatusAuthFailed      *int64  `tfsdk:"error_status_auth_failed" json:"errorStatusAuthFailed,omitempty"`
							ErrorStatusAuthMissing     *int64  `tfsdk:"error_status_auth_missing" json:"errorStatusAuthMissing,omitempty"`
							ErrorStatusLimitsExceeded  *int64  `tfsdk:"error_status_limits_exceeded" json:"errorStatusLimitsExceeded,omitempty"`
							ErrorStatusNoMatch         *int64  `tfsdk:"error_status_no_match" json:"errorStatusNoMatch,omitempty"`
						} `tfsdk:"gateway_response" json:"gatewayResponse,omitempty"`
						IssuerEndpoint    *string `tfsdk:"issuer_endpoint" json:"issuerEndpoint,omitempty"`
						IssuerEndpointRef *struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"issuer_endpoint_ref" json:"issuerEndpointRef,omitempty"`
						IssuerType               *string `tfsdk:"issuer_type" json:"issuerType,omitempty"`
						JwtClaimWithClientID     *string `tfsdk:"jwt_claim_with_client_id" json:"jwtClaimWithClientID,omitempty"`
						JwtClaimWithClientIDType *string `tfsdk:"jwt_claim_with_client_id_type" json:"jwtClaimWithClientIDType,omitempty"`
						Security                 *struct {
							HostHeader  *string `tfsdk:"host_header" json:"hostHeader,omitempty"`
							SecretToken *string `tfsdk:"secret_token" json:"secretToken,omitempty"`
						} `tfsdk:"security" json:"security,omitempty"`
					} `tfsdk:"oidc" json:"oidc,omitempty"`
					Userkey *struct {
						AuthUserKey     *string `tfsdk:"auth_user_key" json:"authUserKey,omitempty"`
						Credentials     *string `tfsdk:"credentials" json:"credentials,omitempty"`
						GatewayResponse *struct {
							ErrorAuthFailed            *string `tfsdk:"error_auth_failed" json:"errorAuthFailed,omitempty"`
							ErrorAuthMissing           *string `tfsdk:"error_auth_missing" json:"errorAuthMissing,omitempty"`
							ErrorHeadersAuthFailed     *string `tfsdk:"error_headers_auth_failed" json:"errorHeadersAuthFailed,omitempty"`
							ErrorHeadersAuthMissing    *string `tfsdk:"error_headers_auth_missing" json:"errorHeadersAuthMissing,omitempty"`
							ErrorHeadersLimitsExceeded *string `tfsdk:"error_headers_limits_exceeded" json:"errorHeadersLimitsExceeded,omitempty"`
							ErrorHeadersNoMatch        *string `tfsdk:"error_headers_no_match" json:"errorHeadersNoMatch,omitempty"`
							ErrorLimitsExceeded        *string `tfsdk:"error_limits_exceeded" json:"errorLimitsExceeded,omitempty"`
							ErrorNoMatch               *string `tfsdk:"error_no_match" json:"errorNoMatch,omitempty"`
							ErrorStatusAuthFailed      *int64  `tfsdk:"error_status_auth_failed" json:"errorStatusAuthFailed,omitempty"`
							ErrorStatusAuthMissing     *int64  `tfsdk:"error_status_auth_missing" json:"errorStatusAuthMissing,omitempty"`
							ErrorStatusLimitsExceeded  *int64  `tfsdk:"error_status_limits_exceeded" json:"errorStatusLimitsExceeded,omitempty"`
							ErrorStatusNoMatch         *int64  `tfsdk:"error_status_no_match" json:"errorStatusNoMatch,omitempty"`
						} `tfsdk:"gateway_response" json:"gatewayResponse,omitempty"`
						Security *struct {
							HostHeader  *string `tfsdk:"host_header" json:"hostHeader,omitempty"`
							SecretToken *string `tfsdk:"secret_token" json:"secretToken,omitempty"`
						} `tfsdk:"security" json:"security,omitempty"`
					} `tfsdk:"userkey" json:"userkey,omitempty"`
				} `tfsdk:"authentication" json:"authentication,omitempty"`
			} `tfsdk:"apicast_hosted" json:"apicastHosted,omitempty"`
			ApicastSelfManaged *struct {
				Authentication *struct {
					AppKeyAppID *struct {
						AppID           *string `tfsdk:"app_id" json:"appID,omitempty"`
						AppKey          *string `tfsdk:"app_key" json:"appKey,omitempty"`
						Credentials     *string `tfsdk:"credentials" json:"credentials,omitempty"`
						GatewayResponse *struct {
							ErrorAuthFailed            *string `tfsdk:"error_auth_failed" json:"errorAuthFailed,omitempty"`
							ErrorAuthMissing           *string `tfsdk:"error_auth_missing" json:"errorAuthMissing,omitempty"`
							ErrorHeadersAuthFailed     *string `tfsdk:"error_headers_auth_failed" json:"errorHeadersAuthFailed,omitempty"`
							ErrorHeadersAuthMissing    *string `tfsdk:"error_headers_auth_missing" json:"errorHeadersAuthMissing,omitempty"`
							ErrorHeadersLimitsExceeded *string `tfsdk:"error_headers_limits_exceeded" json:"errorHeadersLimitsExceeded,omitempty"`
							ErrorHeadersNoMatch        *string `tfsdk:"error_headers_no_match" json:"errorHeadersNoMatch,omitempty"`
							ErrorLimitsExceeded        *string `tfsdk:"error_limits_exceeded" json:"errorLimitsExceeded,omitempty"`
							ErrorNoMatch               *string `tfsdk:"error_no_match" json:"errorNoMatch,omitempty"`
							ErrorStatusAuthFailed      *int64  `tfsdk:"error_status_auth_failed" json:"errorStatusAuthFailed,omitempty"`
							ErrorStatusAuthMissing     *int64  `tfsdk:"error_status_auth_missing" json:"errorStatusAuthMissing,omitempty"`
							ErrorStatusLimitsExceeded  *int64  `tfsdk:"error_status_limits_exceeded" json:"errorStatusLimitsExceeded,omitempty"`
							ErrorStatusNoMatch         *int64  `tfsdk:"error_status_no_match" json:"errorStatusNoMatch,omitempty"`
						} `tfsdk:"gateway_response" json:"gatewayResponse,omitempty"`
						Security *struct {
							HostHeader  *string `tfsdk:"host_header" json:"hostHeader,omitempty"`
							SecretToken *string `tfsdk:"secret_token" json:"secretToken,omitempty"`
						} `tfsdk:"security" json:"security,omitempty"`
					} `tfsdk:"app_key_app_id" json:"appKeyAppID,omitempty"`
					Oidc *struct {
						AuthenticationFlow *struct {
							DirectAccessGrantsEnabled *bool `tfsdk:"direct_access_grants_enabled" json:"directAccessGrantsEnabled,omitempty"`
							ImplicitFlowEnabled       *bool `tfsdk:"implicit_flow_enabled" json:"implicitFlowEnabled,omitempty"`
							ServiceAccountsEnabled    *bool `tfsdk:"service_accounts_enabled" json:"serviceAccountsEnabled,omitempty"`
							StandardFlowEnabled       *bool `tfsdk:"standard_flow_enabled" json:"standardFlowEnabled,omitempty"`
						} `tfsdk:"authentication_flow" json:"authenticationFlow,omitempty"`
						Credentials     *string `tfsdk:"credentials" json:"credentials,omitempty"`
						GatewayResponse *struct {
							ErrorAuthFailed            *string `tfsdk:"error_auth_failed" json:"errorAuthFailed,omitempty"`
							ErrorAuthMissing           *string `tfsdk:"error_auth_missing" json:"errorAuthMissing,omitempty"`
							ErrorHeadersAuthFailed     *string `tfsdk:"error_headers_auth_failed" json:"errorHeadersAuthFailed,omitempty"`
							ErrorHeadersAuthMissing    *string `tfsdk:"error_headers_auth_missing" json:"errorHeadersAuthMissing,omitempty"`
							ErrorHeadersLimitsExceeded *string `tfsdk:"error_headers_limits_exceeded" json:"errorHeadersLimitsExceeded,omitempty"`
							ErrorHeadersNoMatch        *string `tfsdk:"error_headers_no_match" json:"errorHeadersNoMatch,omitempty"`
							ErrorLimitsExceeded        *string `tfsdk:"error_limits_exceeded" json:"errorLimitsExceeded,omitempty"`
							ErrorNoMatch               *string `tfsdk:"error_no_match" json:"errorNoMatch,omitempty"`
							ErrorStatusAuthFailed      *int64  `tfsdk:"error_status_auth_failed" json:"errorStatusAuthFailed,omitempty"`
							ErrorStatusAuthMissing     *int64  `tfsdk:"error_status_auth_missing" json:"errorStatusAuthMissing,omitempty"`
							ErrorStatusLimitsExceeded  *int64  `tfsdk:"error_status_limits_exceeded" json:"errorStatusLimitsExceeded,omitempty"`
							ErrorStatusNoMatch         *int64  `tfsdk:"error_status_no_match" json:"errorStatusNoMatch,omitempty"`
						} `tfsdk:"gateway_response" json:"gatewayResponse,omitempty"`
						IssuerEndpoint    *string `tfsdk:"issuer_endpoint" json:"issuerEndpoint,omitempty"`
						IssuerEndpointRef *struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"issuer_endpoint_ref" json:"issuerEndpointRef,omitempty"`
						IssuerType               *string `tfsdk:"issuer_type" json:"issuerType,omitempty"`
						JwtClaimWithClientID     *string `tfsdk:"jwt_claim_with_client_id" json:"jwtClaimWithClientID,omitempty"`
						JwtClaimWithClientIDType *string `tfsdk:"jwt_claim_with_client_id_type" json:"jwtClaimWithClientIDType,omitempty"`
						Security                 *struct {
							HostHeader  *string `tfsdk:"host_header" json:"hostHeader,omitempty"`
							SecretToken *string `tfsdk:"secret_token" json:"secretToken,omitempty"`
						} `tfsdk:"security" json:"security,omitempty"`
					} `tfsdk:"oidc" json:"oidc,omitempty"`
					Userkey *struct {
						AuthUserKey     *string `tfsdk:"auth_user_key" json:"authUserKey,omitempty"`
						Credentials     *string `tfsdk:"credentials" json:"credentials,omitempty"`
						GatewayResponse *struct {
							ErrorAuthFailed            *string `tfsdk:"error_auth_failed" json:"errorAuthFailed,omitempty"`
							ErrorAuthMissing           *string `tfsdk:"error_auth_missing" json:"errorAuthMissing,omitempty"`
							ErrorHeadersAuthFailed     *string `tfsdk:"error_headers_auth_failed" json:"errorHeadersAuthFailed,omitempty"`
							ErrorHeadersAuthMissing    *string `tfsdk:"error_headers_auth_missing" json:"errorHeadersAuthMissing,omitempty"`
							ErrorHeadersLimitsExceeded *string `tfsdk:"error_headers_limits_exceeded" json:"errorHeadersLimitsExceeded,omitempty"`
							ErrorHeadersNoMatch        *string `tfsdk:"error_headers_no_match" json:"errorHeadersNoMatch,omitempty"`
							ErrorLimitsExceeded        *string `tfsdk:"error_limits_exceeded" json:"errorLimitsExceeded,omitempty"`
							ErrorNoMatch               *string `tfsdk:"error_no_match" json:"errorNoMatch,omitempty"`
							ErrorStatusAuthFailed      *int64  `tfsdk:"error_status_auth_failed" json:"errorStatusAuthFailed,omitempty"`
							ErrorStatusAuthMissing     *int64  `tfsdk:"error_status_auth_missing" json:"errorStatusAuthMissing,omitempty"`
							ErrorStatusLimitsExceeded  *int64  `tfsdk:"error_status_limits_exceeded" json:"errorStatusLimitsExceeded,omitempty"`
							ErrorStatusNoMatch         *int64  `tfsdk:"error_status_no_match" json:"errorStatusNoMatch,omitempty"`
						} `tfsdk:"gateway_response" json:"gatewayResponse,omitempty"`
						Security *struct {
							HostHeader  *string `tfsdk:"host_header" json:"hostHeader,omitempty"`
							SecretToken *string `tfsdk:"secret_token" json:"secretToken,omitempty"`
						} `tfsdk:"security" json:"security,omitempty"`
					} `tfsdk:"userkey" json:"userkey,omitempty"`
				} `tfsdk:"authentication" json:"authentication,omitempty"`
				ProductionPublicBaseURL *string `tfsdk:"production_public_base_url" json:"productionPublicBaseURL,omitempty"`
				StagingPublicBaseURL    *string `tfsdk:"staging_public_base_url" json:"stagingPublicBaseURL,omitempty"`
			} `tfsdk:"apicast_self_managed" json:"apicastSelfManaged,omitempty"`
		} `tfsdk:"deployment" json:"deployment,omitempty"`
		Description  *string `tfsdk:"description" json:"description,omitempty"`
		MappingRules *[]struct {
			HttpMethod      *string `tfsdk:"http_method" json:"httpMethod,omitempty"`
			Increment       *int64  `tfsdk:"increment" json:"increment,omitempty"`
			Last            *bool   `tfsdk:"last" json:"last,omitempty"`
			MetricMethodRef *string `tfsdk:"metric_method_ref" json:"metricMethodRef,omitempty"`
			Pattern         *string `tfsdk:"pattern" json:"pattern,omitempty"`
		} `tfsdk:"mapping_rules" json:"mappingRules,omitempty"`
		Methods *struct {
			Description  *string `tfsdk:"description" json:"description,omitempty"`
			FriendlyName *string `tfsdk:"friendly_name" json:"friendlyName,omitempty"`
		} `tfsdk:"methods" json:"methods,omitempty"`
		Metrics *struct {
			Description  *string `tfsdk:"description" json:"description,omitempty"`
			FriendlyName *string `tfsdk:"friendly_name" json:"friendlyName,omitempty"`
			Unit         *string `tfsdk:"unit" json:"unit,omitempty"`
		} `tfsdk:"metrics" json:"metrics,omitempty"`
		Name     *string `tfsdk:"name" json:"name,omitempty"`
		Policies *[]struct {
			Configuration    *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
			ConfigurationRef *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"configuration_ref" json:"configurationRef,omitempty"`
			Enabled *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
			Name    *string `tfsdk:"name" json:"name,omitempty"`
			Version *string `tfsdk:"version" json:"version,omitempty"`
		} `tfsdk:"policies" json:"policies,omitempty"`
		ProviderAccountRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"provider_account_ref" json:"providerAccountRef,omitempty"`
		SystemName *string `tfsdk:"system_name" json:"systemName,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *Capabilities3ScaleNetProductV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_capabilities_3scale_net_product_v1beta1_manifest"
}

func (r *Capabilities3ScaleNetProductV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Product is the Schema for the products API",
		MarkdownDescription: "Product is the Schema for the products API",
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
				Description:         "ProductSpec defines the desired state of Product",
				MarkdownDescription: "ProductSpec defines the desired state of Product",
				Attributes: map[string]schema.Attribute{
					"application_plans": schema.SingleNestedAttribute{
						Description:         "Application Plans Map: system_name -> Application Plan Spec",
						MarkdownDescription: "Application Plans Map: system_name -> Application Plan Spec",
						Attributes: map[string]schema.Attribute{
							"apps_require_approval": schema.BoolAttribute{
								Description:         "Set whether or not applications can be created on demand or if approval is required from you before they are activated.",
								MarkdownDescription: "Set whether or not applications can be created on demand or if approval is required from you before they are activated.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cost_month": schema.StringAttribute{
								Description:         "Cost per Month (USD)",
								MarkdownDescription: "Cost per Month (USD)",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(\.\d{2})?$`), ""),
								},
							},

							"limits": schema.ListNestedAttribute{
								Description:         "Limits",
								MarkdownDescription: "Limits",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"metric_method_ref": schema.SingleNestedAttribute{
											Description:         "Metric or Method Reference",
											MarkdownDescription: "Metric or Method Reference",
											Attributes: map[string]schema.Attribute{
												"backend": schema.StringAttribute{
													Description:         "BackendSystemName identifies uniquely the backend Backend reference must be used by the product",
													MarkdownDescription: "BackendSystemName identifies uniquely the backend Backend reference must be used by the product",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"system_name": schema.StringAttribute{
													Description:         "SystemName identifies uniquely the metric or methods",
													MarkdownDescription: "SystemName identifies uniquely the metric or methods",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"period": schema.StringAttribute{
											Description:         "Limit Period",
											MarkdownDescription: "Limit Period",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("eternity", "year", "month", "week", "day", "hour", "minute"),
											},
										},

										"value": schema.Int64Attribute{
											Description:         "Limit Value",
											MarkdownDescription: "Limit Value",
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

							"name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pricing_rules": schema.ListNestedAttribute{
								Description:         "Pricing Rules",
								MarkdownDescription: "Pricing Rules",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"from": schema.Int64Attribute{
											Description:         "Range From",
											MarkdownDescription: "Range From",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"metric_method_ref": schema.SingleNestedAttribute{
											Description:         "Metric or Method Reference",
											MarkdownDescription: "Metric or Method Reference",
											Attributes: map[string]schema.Attribute{
												"backend": schema.StringAttribute{
													Description:         "BackendSystemName identifies uniquely the backend Backend reference must be used by the product",
													MarkdownDescription: "BackendSystemName identifies uniquely the backend Backend reference must be used by the product",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"system_name": schema.StringAttribute{
													Description:         "SystemName identifies uniquely the metric or methods",
													MarkdownDescription: "SystemName identifies uniquely the metric or methods",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"price_per_unit": schema.StringAttribute{
											Description:         "Price per unit (USD)",
											MarkdownDescription: "Price per unit (USD)",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(\.\d{2})?$`), ""),
											},
										},

										"to": schema.Int64Attribute{
											Description:         "Range To",
											MarkdownDescription: "Range To",
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

							"published": schema.BoolAttribute{
								Description:         "Controls whether the application plan is published. If not specified it is hidden by default",
								MarkdownDescription: "Controls whether the application plan is published. If not specified it is hidden by default",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"setup_fee": schema.StringAttribute{
								Description:         "Setup fee (USD)",
								MarkdownDescription: "Setup fee (USD)",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(\.\d{2})?$`), ""),
								},
							},

							"trial_period": schema.Int64Attribute{
								Description:         "Trial Period (days)",
								MarkdownDescription: "Trial Period (days)",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"backend_usages": schema.SingleNestedAttribute{
						Description:         "Backend usage will be a map of Map: system_name -> BackendUsageSpec Having system_name as the index, the structure ensures one backend is not used multiple times.",
						MarkdownDescription: "Backend usage will be a map of Map: system_name -> BackendUsageSpec Having system_name as the index, the structure ensures one backend is not used multiple times.",
						Attributes: map[string]schema.Attribute{
							"path": schema.StringAttribute{
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

					"deployment": schema.SingleNestedAttribute{
						Description:         "Deployment defined 3scale product deployment mode",
						MarkdownDescription: "Deployment defined 3scale product deployment mode",
						Attributes: map[string]schema.Attribute{
							"apicast_hosted": schema.SingleNestedAttribute{
								Description:         "ApicastHostedSpec defines the desired state of Product Apicast Hosted",
								MarkdownDescription: "ApicastHostedSpec defines the desired state of Product Apicast Hosted",
								Attributes: map[string]schema.Attribute{
									"authentication": schema.SingleNestedAttribute{
										Description:         "AuthenticationSpec defines the desired state of Product Authentication",
										MarkdownDescription: "AuthenticationSpec defines the desired state of Product Authentication",
										Attributes: map[string]schema.Attribute{
											"app_key_app_id": schema.SingleNestedAttribute{
												Description:         "AppKeyAppIDAuthenticationSpec defines the desired state of AppKey&AppId Authentication",
												MarkdownDescription: "AppKeyAppIDAuthenticationSpec defines the desired state of AppKey&AppId Authentication",
												Attributes: map[string]schema.Attribute{
													"app_id": schema.StringAttribute{
														Description:         "AppID is the name of the parameter that acts of behalf of app id",
														MarkdownDescription: "AppID is the name of the parameter that acts of behalf of app id",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"app_key": schema.StringAttribute{
														Description:         "AppKey is the name of the parameter that acts of behalf of app key",
														MarkdownDescription: "AppKey is the name of the parameter that acts of behalf of app key",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"credentials": schema.StringAttribute{
														Description:         "CredentialsLoc available options: headers: As HTTP Headers query: As query parameters (GET) or body parameters (POST/PUT/DELETE) authorization: As HTTP Basic Authentication",
														MarkdownDescription: "CredentialsLoc available options: headers: As HTTP Headers query: As query parameters (GET) or body parameters (POST/PUT/DELETE) authorization: As HTTP Basic Authentication",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("headers", "query", "authorization"),
														},
													},

													"gateway_response": schema.SingleNestedAttribute{
														Description:         "GatewayResponseSpec defines the desired gateway response configuration",
														MarkdownDescription: "GatewayResponseSpec defines the desired gateway response configuration",
														Attributes: map[string]schema.Attribute{
															"error_auth_failed": schema.StringAttribute{
																Description:         "ErrorAuthFailed specifies the response body when authentication fails",
																MarkdownDescription: "ErrorAuthFailed specifies the response body when authentication fails",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_auth_missing": schema.StringAttribute{
																Description:         "ErrorAuthMissing specifies the response body when authentication is missing",
																MarkdownDescription: "ErrorAuthMissing specifies the response body when authentication is missing",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_headers_auth_failed": schema.StringAttribute{
																Description:         "ErrorHeadersAuthFailed specifies the Content-Type header when authentication fails",
																MarkdownDescription: "ErrorHeadersAuthFailed specifies the Content-Type header when authentication fails",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_headers_auth_missing": schema.StringAttribute{
																Description:         "ErrorHeadersAuthMissing specifies the Content-Type header when authentication is missing",
																MarkdownDescription: "ErrorHeadersAuthMissing specifies the Content-Type header when authentication is missing",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_headers_limits_exceeded": schema.StringAttribute{
																Description:         "ErrorHeadersLimitsExceeded specifies the Content-Type header when usage limit exceeded",
																MarkdownDescription: "ErrorHeadersLimitsExceeded specifies the Content-Type header when usage limit exceeded",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_headers_no_match": schema.StringAttribute{
																Description:         "ErrorHeadersNoMatch specifies the Content-Type header when no match error",
																MarkdownDescription: "ErrorHeadersNoMatch specifies the Content-Type header when no match error",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_limits_exceeded": schema.StringAttribute{
																Description:         "ErrorLimitsExceeded specifies the response body when usage limit exceeded",
																MarkdownDescription: "ErrorLimitsExceeded specifies the response body when usage limit exceeded",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_no_match": schema.StringAttribute{
																Description:         "ErrorNoMatch specifies the response body when no match error",
																MarkdownDescription: "ErrorNoMatch specifies the response body when no match error",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_status_auth_failed": schema.Int64Attribute{
																Description:         "ErrorStatusAuthFailed specifies the response code when authentication fails",
																MarkdownDescription: "ErrorStatusAuthFailed specifies the response code when authentication fails",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_status_auth_missing": schema.Int64Attribute{
																Description:         "ErrorStatusAuthMissing specifies the response code when authentication is missing",
																MarkdownDescription: "ErrorStatusAuthMissing specifies the response code when authentication is missing",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_status_limits_exceeded": schema.Int64Attribute{
																Description:         "ErrorStatusLimitsExceeded specifies the response code when usage limit exceeded",
																MarkdownDescription: "ErrorStatusLimitsExceeded specifies the response code when usage limit exceeded",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_status_no_match": schema.Int64Attribute{
																Description:         "ErrorStatusNoMatch specifies the response code when no match error",
																MarkdownDescription: "ErrorStatusNoMatch specifies the response code when no match error",
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
														Description:         "SecuritySpec defines the desired state of Authentication Security",
														MarkdownDescription: "SecuritySpec defines the desired state of Authentication Security",
														Attributes: map[string]schema.Attribute{
															"host_header": schema.StringAttribute{
																Description:         "HostHeader Lets you define a custom Host request header. This is needed if your API backend only accepts traffic from a specific host.",
																MarkdownDescription: "HostHeader Lets you define a custom Host request header. This is needed if your API backend only accepts traffic from a specific host.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_token": schema.StringAttribute{
																Description:         "SecretToken Enables you to block any direct developer requests to your API backend; each 3scale API gateway call to your API backend contains a request header called X-3scale-proxy-secret-token. The value of this header can be set by you here. It's up to you ensure your backend only allows calls with this secret header.",
																MarkdownDescription: "SecretToken Enables you to block any direct developer requests to your API backend; each 3scale API gateway call to your API backend contains a request header called X-3scale-proxy-secret-token. The value of this header can be set by you here. It's up to you ensure your backend only allows calls with this secret header.",
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

											"oidc": schema.SingleNestedAttribute{
												Description:         "OIDCSpec defines the desired configuration of OpenID Connect Authentication",
												MarkdownDescription: "OIDCSpec defines the desired configuration of OpenID Connect Authentication",
												Attributes: map[string]schema.Attribute{
													"authentication_flow": schema.SingleNestedAttribute{
														Description:         "AuthenticationFlow specifies OAuth2.0 authorization grant type",
														MarkdownDescription: "AuthenticationFlow specifies OAuth2.0 authorization grant type",
														Attributes: map[string]schema.Attribute{
															"direct_access_grants_enabled": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"implicit_flow_enabled": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"service_accounts_enabled": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"standard_flow_enabled": schema.BoolAttribute{
																Description:         "OIDCIssuer is the OIDC issuer",
																MarkdownDescription: "OIDCIssuer is the OIDC issuer",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"credentials": schema.StringAttribute{
														Description:         "Credentials Location available options: headers: As HTTP Headers query: As query parameters (GET) or body parameters (POST/PUT/DELETE) authorization: As HTTP Basic Authentication",
														MarkdownDescription: "Credentials Location available options: headers: As HTTP Headers query: As query parameters (GET) or body parameters (POST/PUT/DELETE) authorization: As HTTP Basic Authentication",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("headers", "query", "authorization"),
														},
													},

													"gateway_response": schema.SingleNestedAttribute{
														Description:         "GatewayResponseSpec defines the desired gateway response configuration",
														MarkdownDescription: "GatewayResponseSpec defines the desired gateway response configuration",
														Attributes: map[string]schema.Attribute{
															"error_auth_failed": schema.StringAttribute{
																Description:         "ErrorAuthFailed specifies the response body when authentication fails",
																MarkdownDescription: "ErrorAuthFailed specifies the response body when authentication fails",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_auth_missing": schema.StringAttribute{
																Description:         "ErrorAuthMissing specifies the response body when authentication is missing",
																MarkdownDescription: "ErrorAuthMissing specifies the response body when authentication is missing",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_headers_auth_failed": schema.StringAttribute{
																Description:         "ErrorHeadersAuthFailed specifies the Content-Type header when authentication fails",
																MarkdownDescription: "ErrorHeadersAuthFailed specifies the Content-Type header when authentication fails",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_headers_auth_missing": schema.StringAttribute{
																Description:         "ErrorHeadersAuthMissing specifies the Content-Type header when authentication is missing",
																MarkdownDescription: "ErrorHeadersAuthMissing specifies the Content-Type header when authentication is missing",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_headers_limits_exceeded": schema.StringAttribute{
																Description:         "ErrorHeadersLimitsExceeded specifies the Content-Type header when usage limit exceeded",
																MarkdownDescription: "ErrorHeadersLimitsExceeded specifies the Content-Type header when usage limit exceeded",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_headers_no_match": schema.StringAttribute{
																Description:         "ErrorHeadersNoMatch specifies the Content-Type header when no match error",
																MarkdownDescription: "ErrorHeadersNoMatch specifies the Content-Type header when no match error",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_limits_exceeded": schema.StringAttribute{
																Description:         "ErrorLimitsExceeded specifies the response body when usage limit exceeded",
																MarkdownDescription: "ErrorLimitsExceeded specifies the response body when usage limit exceeded",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_no_match": schema.StringAttribute{
																Description:         "ErrorNoMatch specifies the response body when no match error",
																MarkdownDescription: "ErrorNoMatch specifies the response body when no match error",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_status_auth_failed": schema.Int64Attribute{
																Description:         "ErrorStatusAuthFailed specifies the response code when authentication fails",
																MarkdownDescription: "ErrorStatusAuthFailed specifies the response code when authentication fails",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_status_auth_missing": schema.Int64Attribute{
																Description:         "ErrorStatusAuthMissing specifies the response code when authentication is missing",
																MarkdownDescription: "ErrorStatusAuthMissing specifies the response code when authentication is missing",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_status_limits_exceeded": schema.Int64Attribute{
																Description:         "ErrorStatusLimitsExceeded specifies the response code when usage limit exceeded",
																MarkdownDescription: "ErrorStatusLimitsExceeded specifies the response code when usage limit exceeded",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_status_no_match": schema.Int64Attribute{
																Description:         "ErrorStatusNoMatch specifies the response code when no match error",
																MarkdownDescription: "ErrorStatusNoMatch specifies the response code when no match error",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"issuer_endpoint": schema.StringAttribute{
														Description:         "Issuer is the OIDC issuer",
														MarkdownDescription: "Issuer is the OIDC issuer",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"issuer_endpoint_ref": schema.SingleNestedAttribute{
														Description:         "IssuerEndpointRef  is the reference to OIDC issuer Secret that contains IssuerEndpoint",
														MarkdownDescription: "IssuerEndpointRef  is the reference to OIDC issuer Secret that contains IssuerEndpoint",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "name is unique within a namespace to reference a secret resource.",
																MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "namespace defines the space within which the secret name must be unique.",
																MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"issuer_type": schema.StringAttribute{
														Description:         "IssuerType is the type of the OIDC issuer",
														MarkdownDescription: "IssuerType is the type of the OIDC issuer",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("keycloak", "rest"),
														},
													},

													"jwt_claim_with_client_id": schema.StringAttribute{
														Description:         "JwtClaimWithClientID is the JSON Web Token (JWT) Claim with ClientID that contains the clientID. Defaults to 'azp'.",
														MarkdownDescription: "JwtClaimWithClientID is the JSON Web Token (JWT) Claim with ClientID that contains the clientID. Defaults to 'azp'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"jwt_claim_with_client_id_type": schema.StringAttribute{
														Description:         "JwtClaimWithClientIDType sets to process the ClientID Token Claim value as a string or as a liquid template.",
														MarkdownDescription: "JwtClaimWithClientIDType sets to process the ClientID Token Claim value as a string or as a liquid template.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("plain", "liquid"),
														},
													},

													"security": schema.SingleNestedAttribute{
														Description:         "SecuritySpec defines the desired state of Authentication Security",
														MarkdownDescription: "SecuritySpec defines the desired state of Authentication Security",
														Attributes: map[string]schema.Attribute{
															"host_header": schema.StringAttribute{
																Description:         "HostHeader Lets you define a custom Host request header. This is needed if your API backend only accepts traffic from a specific host.",
																MarkdownDescription: "HostHeader Lets you define a custom Host request header. This is needed if your API backend only accepts traffic from a specific host.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_token": schema.StringAttribute{
																Description:         "SecretToken Enables you to block any direct developer requests to your API backend; each 3scale API gateway call to your API backend contains a request header called X-3scale-proxy-secret-token. The value of this header can be set by you here. It's up to you ensure your backend only allows calls with this secret header.",
																MarkdownDescription: "SecretToken Enables you to block any direct developer requests to your API backend; each 3scale API gateway call to your API backend contains a request header called X-3scale-proxy-secret-token. The value of this header can be set by you here. It's up to you ensure your backend only allows calls with this secret header.",
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

											"userkey": schema.SingleNestedAttribute{
												Description:         "UserKeyAuthenticationSpec defines the desired state of User Key Authentication",
												MarkdownDescription: "UserKeyAuthenticationSpec defines the desired state of User Key Authentication",
												Attributes: map[string]schema.Attribute{
													"auth_user_key": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"credentials": schema.StringAttribute{
														Description:         "Credentials Location available options: headers: As HTTP Headers query: As query parameters (GET) or body parameters (POST/PUT/DELETE) authorization: As HTTP Basic Authentication",
														MarkdownDescription: "Credentials Location available options: headers: As HTTP Headers query: As query parameters (GET) or body parameters (POST/PUT/DELETE) authorization: As HTTP Basic Authentication",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("headers", "query", "authorization"),
														},
													},

													"gateway_response": schema.SingleNestedAttribute{
														Description:         "GatewayResponseSpec defines the desired gateway response configuration",
														MarkdownDescription: "GatewayResponseSpec defines the desired gateway response configuration",
														Attributes: map[string]schema.Attribute{
															"error_auth_failed": schema.StringAttribute{
																Description:         "ErrorAuthFailed specifies the response body when authentication fails",
																MarkdownDescription: "ErrorAuthFailed specifies the response body when authentication fails",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_auth_missing": schema.StringAttribute{
																Description:         "ErrorAuthMissing specifies the response body when authentication is missing",
																MarkdownDescription: "ErrorAuthMissing specifies the response body when authentication is missing",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_headers_auth_failed": schema.StringAttribute{
																Description:         "ErrorHeadersAuthFailed specifies the Content-Type header when authentication fails",
																MarkdownDescription: "ErrorHeadersAuthFailed specifies the Content-Type header when authentication fails",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_headers_auth_missing": schema.StringAttribute{
																Description:         "ErrorHeadersAuthMissing specifies the Content-Type header when authentication is missing",
																MarkdownDescription: "ErrorHeadersAuthMissing specifies the Content-Type header when authentication is missing",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_headers_limits_exceeded": schema.StringAttribute{
																Description:         "ErrorHeadersLimitsExceeded specifies the Content-Type header when usage limit exceeded",
																MarkdownDescription: "ErrorHeadersLimitsExceeded specifies the Content-Type header when usage limit exceeded",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_headers_no_match": schema.StringAttribute{
																Description:         "ErrorHeadersNoMatch specifies the Content-Type header when no match error",
																MarkdownDescription: "ErrorHeadersNoMatch specifies the Content-Type header when no match error",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_limits_exceeded": schema.StringAttribute{
																Description:         "ErrorLimitsExceeded specifies the response body when usage limit exceeded",
																MarkdownDescription: "ErrorLimitsExceeded specifies the response body when usage limit exceeded",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_no_match": schema.StringAttribute{
																Description:         "ErrorNoMatch specifies the response body when no match error",
																MarkdownDescription: "ErrorNoMatch specifies the response body when no match error",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_status_auth_failed": schema.Int64Attribute{
																Description:         "ErrorStatusAuthFailed specifies the response code when authentication fails",
																MarkdownDescription: "ErrorStatusAuthFailed specifies the response code when authentication fails",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_status_auth_missing": schema.Int64Attribute{
																Description:         "ErrorStatusAuthMissing specifies the response code when authentication is missing",
																MarkdownDescription: "ErrorStatusAuthMissing specifies the response code when authentication is missing",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_status_limits_exceeded": schema.Int64Attribute{
																Description:         "ErrorStatusLimitsExceeded specifies the response code when usage limit exceeded",
																MarkdownDescription: "ErrorStatusLimitsExceeded specifies the response code when usage limit exceeded",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_status_no_match": schema.Int64Attribute{
																Description:         "ErrorStatusNoMatch specifies the response code when no match error",
																MarkdownDescription: "ErrorStatusNoMatch specifies the response code when no match error",
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
														Description:         "SecuritySpec defines the desired state of Authentication Security",
														MarkdownDescription: "SecuritySpec defines the desired state of Authentication Security",
														Attributes: map[string]schema.Attribute{
															"host_header": schema.StringAttribute{
																Description:         "HostHeader Lets you define a custom Host request header. This is needed if your API backend only accepts traffic from a specific host.",
																MarkdownDescription: "HostHeader Lets you define a custom Host request header. This is needed if your API backend only accepts traffic from a specific host.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_token": schema.StringAttribute{
																Description:         "SecretToken Enables you to block any direct developer requests to your API backend; each 3scale API gateway call to your API backend contains a request header called X-3scale-proxy-secret-token. The value of this header can be set by you here. It's up to you ensure your backend only allows calls with this secret header.",
																MarkdownDescription: "SecretToken Enables you to block any direct developer requests to your API backend; each 3scale API gateway call to your API backend contains a request header called X-3scale-proxy-secret-token. The value of this header can be set by you here. It's up to you ensure your backend only allows calls with this secret header.",
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

							"apicast_self_managed": schema.SingleNestedAttribute{
								Description:         "ApicastSelfManagedSpec defines the desired state of Product Apicast Self Managed",
								MarkdownDescription: "ApicastSelfManagedSpec defines the desired state of Product Apicast Self Managed",
								Attributes: map[string]schema.Attribute{
									"authentication": schema.SingleNestedAttribute{
										Description:         "AuthenticationSpec defines the desired state of Product Authentication",
										MarkdownDescription: "AuthenticationSpec defines the desired state of Product Authentication",
										Attributes: map[string]schema.Attribute{
											"app_key_app_id": schema.SingleNestedAttribute{
												Description:         "AppKeyAppIDAuthenticationSpec defines the desired state of AppKey&AppId Authentication",
												MarkdownDescription: "AppKeyAppIDAuthenticationSpec defines the desired state of AppKey&AppId Authentication",
												Attributes: map[string]schema.Attribute{
													"app_id": schema.StringAttribute{
														Description:         "AppID is the name of the parameter that acts of behalf of app id",
														MarkdownDescription: "AppID is the name of the parameter that acts of behalf of app id",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"app_key": schema.StringAttribute{
														Description:         "AppKey is the name of the parameter that acts of behalf of app key",
														MarkdownDescription: "AppKey is the name of the parameter that acts of behalf of app key",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"credentials": schema.StringAttribute{
														Description:         "CredentialsLoc available options: headers: As HTTP Headers query: As query parameters (GET) or body parameters (POST/PUT/DELETE) authorization: As HTTP Basic Authentication",
														MarkdownDescription: "CredentialsLoc available options: headers: As HTTP Headers query: As query parameters (GET) or body parameters (POST/PUT/DELETE) authorization: As HTTP Basic Authentication",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("headers", "query", "authorization"),
														},
													},

													"gateway_response": schema.SingleNestedAttribute{
														Description:         "GatewayResponseSpec defines the desired gateway response configuration",
														MarkdownDescription: "GatewayResponseSpec defines the desired gateway response configuration",
														Attributes: map[string]schema.Attribute{
															"error_auth_failed": schema.StringAttribute{
																Description:         "ErrorAuthFailed specifies the response body when authentication fails",
																MarkdownDescription: "ErrorAuthFailed specifies the response body when authentication fails",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_auth_missing": schema.StringAttribute{
																Description:         "ErrorAuthMissing specifies the response body when authentication is missing",
																MarkdownDescription: "ErrorAuthMissing specifies the response body when authentication is missing",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_headers_auth_failed": schema.StringAttribute{
																Description:         "ErrorHeadersAuthFailed specifies the Content-Type header when authentication fails",
																MarkdownDescription: "ErrorHeadersAuthFailed specifies the Content-Type header when authentication fails",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_headers_auth_missing": schema.StringAttribute{
																Description:         "ErrorHeadersAuthMissing specifies the Content-Type header when authentication is missing",
																MarkdownDescription: "ErrorHeadersAuthMissing specifies the Content-Type header when authentication is missing",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_headers_limits_exceeded": schema.StringAttribute{
																Description:         "ErrorHeadersLimitsExceeded specifies the Content-Type header when usage limit exceeded",
																MarkdownDescription: "ErrorHeadersLimitsExceeded specifies the Content-Type header when usage limit exceeded",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_headers_no_match": schema.StringAttribute{
																Description:         "ErrorHeadersNoMatch specifies the Content-Type header when no match error",
																MarkdownDescription: "ErrorHeadersNoMatch specifies the Content-Type header when no match error",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_limits_exceeded": schema.StringAttribute{
																Description:         "ErrorLimitsExceeded specifies the response body when usage limit exceeded",
																MarkdownDescription: "ErrorLimitsExceeded specifies the response body when usage limit exceeded",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_no_match": schema.StringAttribute{
																Description:         "ErrorNoMatch specifies the response body when no match error",
																MarkdownDescription: "ErrorNoMatch specifies the response body when no match error",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_status_auth_failed": schema.Int64Attribute{
																Description:         "ErrorStatusAuthFailed specifies the response code when authentication fails",
																MarkdownDescription: "ErrorStatusAuthFailed specifies the response code when authentication fails",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_status_auth_missing": schema.Int64Attribute{
																Description:         "ErrorStatusAuthMissing specifies the response code when authentication is missing",
																MarkdownDescription: "ErrorStatusAuthMissing specifies the response code when authentication is missing",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_status_limits_exceeded": schema.Int64Attribute{
																Description:         "ErrorStatusLimitsExceeded specifies the response code when usage limit exceeded",
																MarkdownDescription: "ErrorStatusLimitsExceeded specifies the response code when usage limit exceeded",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_status_no_match": schema.Int64Attribute{
																Description:         "ErrorStatusNoMatch specifies the response code when no match error",
																MarkdownDescription: "ErrorStatusNoMatch specifies the response code when no match error",
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
														Description:         "SecuritySpec defines the desired state of Authentication Security",
														MarkdownDescription: "SecuritySpec defines the desired state of Authentication Security",
														Attributes: map[string]schema.Attribute{
															"host_header": schema.StringAttribute{
																Description:         "HostHeader Lets you define a custom Host request header. This is needed if your API backend only accepts traffic from a specific host.",
																MarkdownDescription: "HostHeader Lets you define a custom Host request header. This is needed if your API backend only accepts traffic from a specific host.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_token": schema.StringAttribute{
																Description:         "SecretToken Enables you to block any direct developer requests to your API backend; each 3scale API gateway call to your API backend contains a request header called X-3scale-proxy-secret-token. The value of this header can be set by you here. It's up to you ensure your backend only allows calls with this secret header.",
																MarkdownDescription: "SecretToken Enables you to block any direct developer requests to your API backend; each 3scale API gateway call to your API backend contains a request header called X-3scale-proxy-secret-token. The value of this header can be set by you here. It's up to you ensure your backend only allows calls with this secret header.",
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

											"oidc": schema.SingleNestedAttribute{
												Description:         "OIDCSpec defines the desired configuration of OpenID Connect Authentication",
												MarkdownDescription: "OIDCSpec defines the desired configuration of OpenID Connect Authentication",
												Attributes: map[string]schema.Attribute{
													"authentication_flow": schema.SingleNestedAttribute{
														Description:         "AuthenticationFlow specifies OAuth2.0 authorization grant type",
														MarkdownDescription: "AuthenticationFlow specifies OAuth2.0 authorization grant type",
														Attributes: map[string]schema.Attribute{
															"direct_access_grants_enabled": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"implicit_flow_enabled": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"service_accounts_enabled": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"standard_flow_enabled": schema.BoolAttribute{
																Description:         "OIDCIssuer is the OIDC issuer",
																MarkdownDescription: "OIDCIssuer is the OIDC issuer",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"credentials": schema.StringAttribute{
														Description:         "Credentials Location available options: headers: As HTTP Headers query: As query parameters (GET) or body parameters (POST/PUT/DELETE) authorization: As HTTP Basic Authentication",
														MarkdownDescription: "Credentials Location available options: headers: As HTTP Headers query: As query parameters (GET) or body parameters (POST/PUT/DELETE) authorization: As HTTP Basic Authentication",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("headers", "query", "authorization"),
														},
													},

													"gateway_response": schema.SingleNestedAttribute{
														Description:         "GatewayResponseSpec defines the desired gateway response configuration",
														MarkdownDescription: "GatewayResponseSpec defines the desired gateway response configuration",
														Attributes: map[string]schema.Attribute{
															"error_auth_failed": schema.StringAttribute{
																Description:         "ErrorAuthFailed specifies the response body when authentication fails",
																MarkdownDescription: "ErrorAuthFailed specifies the response body when authentication fails",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_auth_missing": schema.StringAttribute{
																Description:         "ErrorAuthMissing specifies the response body when authentication is missing",
																MarkdownDescription: "ErrorAuthMissing specifies the response body when authentication is missing",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_headers_auth_failed": schema.StringAttribute{
																Description:         "ErrorHeadersAuthFailed specifies the Content-Type header when authentication fails",
																MarkdownDescription: "ErrorHeadersAuthFailed specifies the Content-Type header when authentication fails",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_headers_auth_missing": schema.StringAttribute{
																Description:         "ErrorHeadersAuthMissing specifies the Content-Type header when authentication is missing",
																MarkdownDescription: "ErrorHeadersAuthMissing specifies the Content-Type header when authentication is missing",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_headers_limits_exceeded": schema.StringAttribute{
																Description:         "ErrorHeadersLimitsExceeded specifies the Content-Type header when usage limit exceeded",
																MarkdownDescription: "ErrorHeadersLimitsExceeded specifies the Content-Type header when usage limit exceeded",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_headers_no_match": schema.StringAttribute{
																Description:         "ErrorHeadersNoMatch specifies the Content-Type header when no match error",
																MarkdownDescription: "ErrorHeadersNoMatch specifies the Content-Type header when no match error",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_limits_exceeded": schema.StringAttribute{
																Description:         "ErrorLimitsExceeded specifies the response body when usage limit exceeded",
																MarkdownDescription: "ErrorLimitsExceeded specifies the response body when usage limit exceeded",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_no_match": schema.StringAttribute{
																Description:         "ErrorNoMatch specifies the response body when no match error",
																MarkdownDescription: "ErrorNoMatch specifies the response body when no match error",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_status_auth_failed": schema.Int64Attribute{
																Description:         "ErrorStatusAuthFailed specifies the response code when authentication fails",
																MarkdownDescription: "ErrorStatusAuthFailed specifies the response code when authentication fails",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_status_auth_missing": schema.Int64Attribute{
																Description:         "ErrorStatusAuthMissing specifies the response code when authentication is missing",
																MarkdownDescription: "ErrorStatusAuthMissing specifies the response code when authentication is missing",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_status_limits_exceeded": schema.Int64Attribute{
																Description:         "ErrorStatusLimitsExceeded specifies the response code when usage limit exceeded",
																MarkdownDescription: "ErrorStatusLimitsExceeded specifies the response code when usage limit exceeded",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_status_no_match": schema.Int64Attribute{
																Description:         "ErrorStatusNoMatch specifies the response code when no match error",
																MarkdownDescription: "ErrorStatusNoMatch specifies the response code when no match error",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"issuer_endpoint": schema.StringAttribute{
														Description:         "Issuer is the OIDC issuer",
														MarkdownDescription: "Issuer is the OIDC issuer",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"issuer_endpoint_ref": schema.SingleNestedAttribute{
														Description:         "IssuerEndpointRef  is the reference to OIDC issuer Secret that contains IssuerEndpoint",
														MarkdownDescription: "IssuerEndpointRef  is the reference to OIDC issuer Secret that contains IssuerEndpoint",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "name is unique within a namespace to reference a secret resource.",
																MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "namespace defines the space within which the secret name must be unique.",
																MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"issuer_type": schema.StringAttribute{
														Description:         "IssuerType is the type of the OIDC issuer",
														MarkdownDescription: "IssuerType is the type of the OIDC issuer",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("keycloak", "rest"),
														},
													},

													"jwt_claim_with_client_id": schema.StringAttribute{
														Description:         "JwtClaimWithClientID is the JSON Web Token (JWT) Claim with ClientID that contains the clientID. Defaults to 'azp'.",
														MarkdownDescription: "JwtClaimWithClientID is the JSON Web Token (JWT) Claim with ClientID that contains the clientID. Defaults to 'azp'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"jwt_claim_with_client_id_type": schema.StringAttribute{
														Description:         "JwtClaimWithClientIDType sets to process the ClientID Token Claim value as a string or as a liquid template.",
														MarkdownDescription: "JwtClaimWithClientIDType sets to process the ClientID Token Claim value as a string or as a liquid template.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("plain", "liquid"),
														},
													},

													"security": schema.SingleNestedAttribute{
														Description:         "SecuritySpec defines the desired state of Authentication Security",
														MarkdownDescription: "SecuritySpec defines the desired state of Authentication Security",
														Attributes: map[string]schema.Attribute{
															"host_header": schema.StringAttribute{
																Description:         "HostHeader Lets you define a custom Host request header. This is needed if your API backend only accepts traffic from a specific host.",
																MarkdownDescription: "HostHeader Lets you define a custom Host request header. This is needed if your API backend only accepts traffic from a specific host.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_token": schema.StringAttribute{
																Description:         "SecretToken Enables you to block any direct developer requests to your API backend; each 3scale API gateway call to your API backend contains a request header called X-3scale-proxy-secret-token. The value of this header can be set by you here. It's up to you ensure your backend only allows calls with this secret header.",
																MarkdownDescription: "SecretToken Enables you to block any direct developer requests to your API backend; each 3scale API gateway call to your API backend contains a request header called X-3scale-proxy-secret-token. The value of this header can be set by you here. It's up to you ensure your backend only allows calls with this secret header.",
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

											"userkey": schema.SingleNestedAttribute{
												Description:         "UserKeyAuthenticationSpec defines the desired state of User Key Authentication",
												MarkdownDescription: "UserKeyAuthenticationSpec defines the desired state of User Key Authentication",
												Attributes: map[string]schema.Attribute{
													"auth_user_key": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"credentials": schema.StringAttribute{
														Description:         "Credentials Location available options: headers: As HTTP Headers query: As query parameters (GET) or body parameters (POST/PUT/DELETE) authorization: As HTTP Basic Authentication",
														MarkdownDescription: "Credentials Location available options: headers: As HTTP Headers query: As query parameters (GET) or body parameters (POST/PUT/DELETE) authorization: As HTTP Basic Authentication",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("headers", "query", "authorization"),
														},
													},

													"gateway_response": schema.SingleNestedAttribute{
														Description:         "GatewayResponseSpec defines the desired gateway response configuration",
														MarkdownDescription: "GatewayResponseSpec defines the desired gateway response configuration",
														Attributes: map[string]schema.Attribute{
															"error_auth_failed": schema.StringAttribute{
																Description:         "ErrorAuthFailed specifies the response body when authentication fails",
																MarkdownDescription: "ErrorAuthFailed specifies the response body when authentication fails",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_auth_missing": schema.StringAttribute{
																Description:         "ErrorAuthMissing specifies the response body when authentication is missing",
																MarkdownDescription: "ErrorAuthMissing specifies the response body when authentication is missing",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_headers_auth_failed": schema.StringAttribute{
																Description:         "ErrorHeadersAuthFailed specifies the Content-Type header when authentication fails",
																MarkdownDescription: "ErrorHeadersAuthFailed specifies the Content-Type header when authentication fails",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_headers_auth_missing": schema.StringAttribute{
																Description:         "ErrorHeadersAuthMissing specifies the Content-Type header when authentication is missing",
																MarkdownDescription: "ErrorHeadersAuthMissing specifies the Content-Type header when authentication is missing",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_headers_limits_exceeded": schema.StringAttribute{
																Description:         "ErrorHeadersLimitsExceeded specifies the Content-Type header when usage limit exceeded",
																MarkdownDescription: "ErrorHeadersLimitsExceeded specifies the Content-Type header when usage limit exceeded",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_headers_no_match": schema.StringAttribute{
																Description:         "ErrorHeadersNoMatch specifies the Content-Type header when no match error",
																MarkdownDescription: "ErrorHeadersNoMatch specifies the Content-Type header when no match error",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_limits_exceeded": schema.StringAttribute{
																Description:         "ErrorLimitsExceeded specifies the response body when usage limit exceeded",
																MarkdownDescription: "ErrorLimitsExceeded specifies the response body when usage limit exceeded",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_no_match": schema.StringAttribute{
																Description:         "ErrorNoMatch specifies the response body when no match error",
																MarkdownDescription: "ErrorNoMatch specifies the response body when no match error",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_status_auth_failed": schema.Int64Attribute{
																Description:         "ErrorStatusAuthFailed specifies the response code when authentication fails",
																MarkdownDescription: "ErrorStatusAuthFailed specifies the response code when authentication fails",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_status_auth_missing": schema.Int64Attribute{
																Description:         "ErrorStatusAuthMissing specifies the response code when authentication is missing",
																MarkdownDescription: "ErrorStatusAuthMissing specifies the response code when authentication is missing",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_status_limits_exceeded": schema.Int64Attribute{
																Description:         "ErrorStatusLimitsExceeded specifies the response code when usage limit exceeded",
																MarkdownDescription: "ErrorStatusLimitsExceeded specifies the response code when usage limit exceeded",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"error_status_no_match": schema.Int64Attribute{
																Description:         "ErrorStatusNoMatch specifies the response code when no match error",
																MarkdownDescription: "ErrorStatusNoMatch specifies the response code when no match error",
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
														Description:         "SecuritySpec defines the desired state of Authentication Security",
														MarkdownDescription: "SecuritySpec defines the desired state of Authentication Security",
														Attributes: map[string]schema.Attribute{
															"host_header": schema.StringAttribute{
																Description:         "HostHeader Lets you define a custom Host request header. This is needed if your API backend only accepts traffic from a specific host.",
																MarkdownDescription: "HostHeader Lets you define a custom Host request header. This is needed if your API backend only accepts traffic from a specific host.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_token": schema.StringAttribute{
																Description:         "SecretToken Enables you to block any direct developer requests to your API backend; each 3scale API gateway call to your API backend contains a request header called X-3scale-proxy-secret-token. The value of this header can be set by you here. It's up to you ensure your backend only allows calls with this secret header.",
																MarkdownDescription: "SecretToken Enables you to block any direct developer requests to your API backend; each 3scale API gateway call to your API backend contains a request header called X-3scale-proxy-secret-token. The value of this header can be set by you here. It's up to you ensure your backend only allows calls with this secret header.",
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

									"production_public_base_url": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^https?:\/\/.*$`), ""),
										},
									},

									"staging_public_base_url": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^https?:\/\/.*$`), ""),
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

					"description": schema.StringAttribute{
						Description:         "Description is a human readable text of the product",
						MarkdownDescription: "Description is a human readable text of the product",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"mapping_rules": schema.ListNestedAttribute{
						Description:         "Mapping Rules Array: MappingRule Spec",
						MarkdownDescription: "Mapping Rules Array: MappingRule Spec",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"http_method": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS", "TRACE", "PATCH", "CONNECT"),
									},
								},

								"increment": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"last": schema.BoolAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metric_method_ref": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"pattern": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"methods": schema.SingleNestedAttribute{
						Description:         "Methods Map: system_name -> MethodSpec system_name attr is unique for all metrics AND methods In other words, if metric's system_name is A, there is no metric or method with system_name A.",
						MarkdownDescription: "Methods Map: system_name -> MethodSpec system_name attr is unique for all metrics AND methods In other words, if metric's system_name is A, there is no metric or method with system_name A.",
						Attributes: map[string]schema.Attribute{
							"description": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"friendly_name": schema.StringAttribute{
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

					"metrics": schema.SingleNestedAttribute{
						Description:         "Metrics Map: system_name -> MetricSpec system_name attr is unique for all metrics AND methods In other words, if metric's system_name is A, there is no metric or method with system_name A.",
						MarkdownDescription: "Metrics Map: system_name -> MetricSpec system_name attr is unique for all metrics AND methods In other words, if metric's system_name is A, there is no metric or method with system_name A.",
						Attributes: map[string]schema.Attribute{
							"description": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"friendly_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"unit": schema.StringAttribute{
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

					"name": schema.StringAttribute{
						Description:         "Name is human readable name for the product",
						MarkdownDescription: "Name is human readable name for the product",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"policies": schema.ListNestedAttribute{
						Description:         "Policies holds the product's policy chain",
						MarkdownDescription: "Policies holds the product's policy chain",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"configuration": schema.MapAttribute{
									Description:         "Configuration defines the policy configuration",
									MarkdownDescription: "Configuration defines the policy configuration",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"configuration_ref": schema.SingleNestedAttribute{
									Description:         "ConfigurationRef Secret reference containing policy configuration",
									MarkdownDescription: "ConfigurationRef Secret reference containing policy configuration",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "name is unique within a namespace to reference a secret resource.",
											MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"namespace": schema.StringAttribute{
											Description:         "namespace defines the space within which the secret name must be unique.",
											MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"enabled": schema.BoolAttribute{
									Description:         "Enabled defines activation state",
									MarkdownDescription: "Enabled defines activation state",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name defines the policy unique name",
									MarkdownDescription: "Name defines the policy unique name",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"version": schema.StringAttribute{
									Description:         "Version defines the policy version",
									MarkdownDescription: "Version defines the policy version",
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

					"provider_account_ref": schema.SingleNestedAttribute{
						Description:         "ProviderAccountRef references account provider credentials",
						MarkdownDescription: "ProviderAccountRef references account provider credentials",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
								MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"system_name": schema.StringAttribute{
						Description:         "SystemName identifies uniquely the product within the account provider Default value will be sanitized Name",
						MarkdownDescription: "SystemName identifies uniquely the product within the account provider Default value will be sanitized Name",
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
	}
}

func (r *Capabilities3ScaleNetProductV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_capabilities_3scale_net_product_v1beta1_manifest")

	var model Capabilities3ScaleNetProductV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("capabilities.3scale.net/v1beta1")
	model.Kind = pointer.String("Product")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
