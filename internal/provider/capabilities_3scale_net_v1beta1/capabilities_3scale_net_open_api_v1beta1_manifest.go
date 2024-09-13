/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package capabilities_3scale_net_v1beta1

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
	_ datasource.DataSource = &Capabilities3ScaleNetOpenApiV1Beta1Manifest{}
)

func NewCapabilities3ScaleNetOpenApiV1Beta1Manifest() datasource.DataSource {
	return &Capabilities3ScaleNetOpenApiV1Beta1Manifest{}
}

type Capabilities3ScaleNetOpenApiV1Beta1Manifest struct{}

type Capabilities3ScaleNetOpenApiV1Beta1ManifestData struct {
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
		OpenapiRef *struct {
			SecretRef *struct {
				ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
				FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
				Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
				Name            *string `tfsdk:"name" json:"name,omitempty"`
				Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
				ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
				Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
			} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
			Url *string `tfsdk:"url" json:"url,omitempty"`
		} `tfsdk:"openapi_ref" json:"openapiRef,omitempty"`
		PrefixMatching          *bool   `tfsdk:"prefix_matching" json:"prefixMatching,omitempty"`
		PrivateAPIHostHeader    *string `tfsdk:"private_api_host_header" json:"privateAPIHostHeader,omitempty"`
		PrivateAPISecretToken   *string `tfsdk:"private_api_secret_token" json:"privateAPISecretToken,omitempty"`
		PrivateBaseURL          *string `tfsdk:"private_base_url" json:"privateBaseURL,omitempty"`
		ProductSystemName       *string `tfsdk:"product_system_name" json:"productSystemName,omitempty"`
		ProductionPublicBaseURL *string `tfsdk:"production_public_base_url" json:"productionPublicBaseURL,omitempty"`
		ProviderAccountRef      *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"provider_account_ref" json:"providerAccountRef,omitempty"`
		StagingPublicBaseURL *string `tfsdk:"staging_public_base_url" json:"stagingPublicBaseURL,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *Capabilities3ScaleNetOpenApiV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_capabilities_3scale_net_open_api_v1beta1_manifest"
}

func (r *Capabilities3ScaleNetOpenApiV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "OpenAPI is the Schema for the openapis API",
		MarkdownDescription: "OpenAPI is the Schema for the openapis API",
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
				Description:         "OpenAPISpec defines the desired state of OpenAPI",
				MarkdownDescription: "OpenAPISpec defines the desired state of OpenAPI",
				Attributes: map[string]schema.Attribute{
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
								Description:         "IssuerEndpointRef is the reference to OIDC issuer Secret that contains IssuerEndpoint",
								MarkdownDescription: "IssuerEndpointRef is the reference to OIDC issuer Secret that contains IssuerEndpoint",
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

					"openapi_ref": schema.SingleNestedAttribute{
						Description:         "OpenAPIRef Reference to the OpenAPI Specification",
						MarkdownDescription: "OpenAPIRef Reference to the OpenAPI Specification",
						Attributes: map[string]schema.Attribute{
							"secret_ref": schema.SingleNestedAttribute{
								Description:         "SecretRef refers to the secret object that contains the OpenAPI Document",
								MarkdownDescription: "SecretRef refers to the secret object that contains the OpenAPI Document",
								Attributes: map[string]schema.Attribute{
									"api_version": schema.StringAttribute{
										Description:         "API version of the referent.",
										MarkdownDescription: "API version of the referent.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"field_path": schema.StringAttribute{
										Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
										MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"kind": schema.StringAttribute{
										Description:         "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
										MarkdownDescription: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
										Description:         "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
										MarkdownDescription: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"resource_version": schema.StringAttribute{
										Description:         "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
										MarkdownDescription: "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"uid": schema.StringAttribute{
										Description:         "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
										MarkdownDescription: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"url": schema.StringAttribute{
								Description:         "URL Remote URL from where to fetch the OpenAPI Document",
								MarkdownDescription: "URL Remote URL from where to fetch the OpenAPI Document",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^https?:\/\/.*$`), ""),
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"prefix_matching": schema.BoolAttribute{
						Description:         "PrefixMatching Use prefix matching instead of strict matching on mapping rules derived from openapi operations",
						MarkdownDescription: "PrefixMatching Use prefix matching instead of strict matching on mapping rules derived from openapi operations",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"private_api_host_header": schema.StringAttribute{
						Description:         "PrivateAPIHostHeader Custom host header sent by the API gateway to the private API",
						MarkdownDescription: "PrivateAPIHostHeader Custom host header sent by the API gateway to the private API",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"private_api_secret_token": schema.StringAttribute{
						Description:         "PrivateAPISecretToken Custom secret token sent by the API gateway to the private API",
						MarkdownDescription: "PrivateAPISecretToken Custom secret token sent by the API gateway to the private API",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"private_base_url": schema.StringAttribute{
						Description:         "PrivateBaseURL Custom private base URL",
						MarkdownDescription: "PrivateBaseURL Custom private base URL",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"product_system_name": schema.StringAttribute{
						Description:         "ProductSystemName 3scale product system name",
						MarkdownDescription: "ProductSystemName 3scale product system name",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"production_public_base_url": schema.StringAttribute{
						Description:         "ProductionPublicBaseURL Custom public production URL",
						MarkdownDescription: "ProductionPublicBaseURL Custom public production URL",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^https?:\/\/.*$`), ""),
						},
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

					"staging_public_base_url": schema.StringAttribute{
						Description:         "StagingPublicBaseURL Custom public staging URL",
						MarkdownDescription: "StagingPublicBaseURL Custom public staging URL",
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
	}
}

func (r *Capabilities3ScaleNetOpenApiV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_capabilities_3scale_net_open_api_v1beta1_manifest")

	var model Capabilities3ScaleNetOpenApiV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("capabilities.3scale.net/v1beta1")
	model.Kind = pointer.String("OpenAPI")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
