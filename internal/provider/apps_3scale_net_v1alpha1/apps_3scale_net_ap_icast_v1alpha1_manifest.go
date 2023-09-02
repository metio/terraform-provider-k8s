/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package apps_3scale_net_v1alpha1

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &Apps3ScaleNetAPIcastV1Alpha1Manifest{}
)

func NewApps3ScaleNetAPIcastV1Alpha1Manifest() datasource.DataSource {
	return &Apps3ScaleNetAPIcastV1Alpha1Manifest{}
}

type Apps3ScaleNetAPIcastV1Alpha1Manifest struct{}

type Apps3ScaleNetAPIcastV1Alpha1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
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
		AdminPortalCredentialsRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"admin_portal_credentials_ref" json:"adminPortalCredentialsRef,omitempty"`
		AllProxy                  *string `tfsdk:"all_proxy" json:"allProxy,omitempty"`
		CacheConfigurationSeconds *int64  `tfsdk:"cache_configuration_seconds" json:"cacheConfigurationSeconds,omitempty"`
		CacheMaxTime              *string `tfsdk:"cache_max_time" json:"cacheMaxTime,omitempty"`
		CacheStatusCodes          *string `tfsdk:"cache_status_codes" json:"cacheStatusCodes,omitempty"`
		ConfigurationLoadMode     *string `tfsdk:"configuration_load_mode" json:"configurationLoadMode,omitempty"`
		CustomEnvironments        *[]struct {
			SecretRef *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
		} `tfsdk:"custom_environments" json:"customEnvironments,omitempty"`
		CustomPolicies *[]struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			SecretRef *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
			Version *string `tfsdk:"version" json:"version,omitempty"`
		} `tfsdk:"custom_policies" json:"customPolicies,omitempty"`
		DeploymentEnvironment          *string `tfsdk:"deployment_environment" json:"deploymentEnvironment,omitempty"`
		DnsResolverAddress             *string `tfsdk:"dns_resolver_address" json:"dnsResolverAddress,omitempty"`
		EmbeddedConfigurationSecretRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"embedded_configuration_secret_ref" json:"embeddedConfigurationSecretRef,omitempty"`
		EnabledServices *[]string `tfsdk:"enabled_services" json:"enabledServices,omitempty"`
		ExposedHost     *struct {
			Host *string `tfsdk:"host" json:"host,omitempty"`
			Tls  *[]struct {
				Hosts      *[]string `tfsdk:"hosts" json:"hosts,omitempty"`
				SecretName *string   `tfsdk:"secret_name" json:"secretName,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
		} `tfsdk:"exposed_host" json:"exposedHost,omitempty"`
		ExtendedMetrics           *bool   `tfsdk:"extended_metrics" json:"extendedMetrics,omitempty"`
		HttpProxy                 *string `tfsdk:"http_proxy" json:"httpProxy,omitempty"`
		HttpsCertificateSecretRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"https_certificate_secret_ref" json:"httpsCertificateSecretRef,omitempty"`
		HttpsPort                      *int64  `tfsdk:"https_port" json:"httpsPort,omitempty"`
		HttpsProxy                     *string `tfsdk:"https_proxy" json:"httpsProxy,omitempty"`
		HttpsVerifyDepth               *int64  `tfsdk:"https_verify_depth" json:"httpsVerifyDepth,omitempty"`
		Image                          *string `tfsdk:"image" json:"image,omitempty"`
		LoadServicesWhenNeeded         *bool   `tfsdk:"load_services_when_needed" json:"loadServicesWhenNeeded,omitempty"`
		LogLevel                       *string `tfsdk:"log_level" json:"logLevel,omitempty"`
		ManagementAPIScope             *string `tfsdk:"management_api_scope" json:"managementAPIScope,omitempty"`
		NoProxy                        *string `tfsdk:"no_proxy" json:"noProxy,omitempty"`
		OidcLogLevel                   *string `tfsdk:"oidc_log_level" json:"oidcLogLevel,omitempty"`
		OpenSSLPeerVerificationEnabled *bool   `tfsdk:"open_ssl_peer_verification_enabled" json:"openSSLPeerVerificationEnabled,omitempty"`
		OpenTelemetry                  *struct {
			Enabled                *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
			TracingConfigSecretKey *string `tfsdk:"tracing_config_secret_key" json:"tracingConfigSecretKey,omitempty"`
			TracingConfigSecretRef *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"tracing_config_secret_ref" json:"tracingConfigSecretRef,omitempty"`
		} `tfsdk:"open_telemetry" json:"openTelemetry,omitempty"`
		OpenTracing *struct {
			Enabled                *bool `tfsdk:"enabled" json:"enabled,omitempty"`
			TracingConfigSecretRef *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"tracing_config_secret_ref" json:"tracingConfigSecretRef,omitempty"`
			TracingLibrary *string `tfsdk:"tracing_library" json:"tracingLibrary,omitempty"`
		} `tfsdk:"open_tracing" json:"openTracing,omitempty"`
		PathRoutingEnabled *bool  `tfsdk:"path_routing_enabled" json:"pathRoutingEnabled,omitempty"`
		Replicas           *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
		Resources          *struct {
			Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
			Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
		} `tfsdk:"resources" json:"resources,omitempty"`
		ResponseCodesIncluded               *bool              `tfsdk:"response_codes_included" json:"responseCodesIncluded,omitempty"`
		ServiceAccount                      *string            `tfsdk:"service_account" json:"serviceAccount,omitempty"`
		ServiceCacheSize                    *int64             `tfsdk:"service_cache_size" json:"serviceCacheSize,omitempty"`
		ServiceConfigurationVersionOverride *map[string]string `tfsdk:"service_configuration_version_override" json:"serviceConfigurationVersionOverride,omitempty"`
		ServicesFilterByURL                 *string            `tfsdk:"services_filter_by_url" json:"servicesFilterByURL,omitempty"`
		Timezone                            *string            `tfsdk:"timezone" json:"timezone,omitempty"`
		UpstreamRetryCases                  *string            `tfsdk:"upstream_retry_cases" json:"upstreamRetryCases,omitempty"`
		Workers                             *int64             `tfsdk:"workers" json:"workers,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *Apps3ScaleNetAPIcastV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_apps_3scale_net_ap_icast_v1alpha1_manifest"
}

func (r *Apps3ScaleNetAPIcastV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "APIcast is the Schema for the apicasts API.",
		MarkdownDescription: "APIcast is the Schema for the apicasts API.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

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
				Description:         "APIcastSpec defines the desired state of APIcast.",
				MarkdownDescription: "APIcastSpec defines the desired state of APIcast.",
				Attributes: map[string]schema.Attribute{
					"admin_portal_credentials_ref": schema.SingleNestedAttribute{
						Description:         "Secret reference to a Kubernetes Secret containing the admin portal endpoint URL. The Secret must be located in the same namespace.",
						MarkdownDescription: "Secret reference to a Kubernetes Secret containing the admin portal endpoint URL. The Secret must be located in the same namespace.",
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

					"all_proxy": schema.StringAttribute{
						Description:         "AllProxy specifies a HTTP(S) proxy to be used for connecting to services if a protocol-specific proxy is not specified. Authentication is not supported. Format is <scheme>://<host>:<port>",
						MarkdownDescription: "AllProxy specifies a HTTP(S) proxy to be used for connecting to services if a protocol-specific proxy is not specified. Authentication is not supported. Format is <scheme>://<host>:<port>",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"cache_configuration_seconds": schema.Int64Attribute{
						Description:         "The period (in seconds) that the APIcast configuration will be stored in APIcast's cache.",
						MarkdownDescription: "The period (in seconds) that the APIcast configuration will be stored in APIcast's cache.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"cache_max_time": schema.StringAttribute{
						Description:         "CacheMaxTime indicates the maximum time to be cached. If cache-control header is not set, the time to be cached will be the defined one.",
						MarkdownDescription: "CacheMaxTime indicates the maximum time to be cached. If cache-control header is not set, the time to be cached will be the defined one.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"cache_status_codes": schema.StringAttribute{
						Description:         "CacheStatusCodes defines the status codes for which the response content will be cached.",
						MarkdownDescription: "CacheStatusCodes defines the status codes for which the response content will be cached.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"configuration_load_mode": schema.StringAttribute{
						Description:         "ConfigurationLoadMode can be used to set APIcast's configuration load mode.",
						MarkdownDescription: "ConfigurationLoadMode can be used to set APIcast's configuration load mode.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("boot", "lazy"),
						},
					},

					"custom_environments": schema.ListNestedAttribute{
						Description:         "CustomEnvironments specifies an array of defined custome environments to be loaded",
						MarkdownDescription: "CustomEnvironments specifies an array of defined custome environments to be loaded",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"secret_ref": schema.SingleNestedAttribute{
									Description:         "LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.",
									MarkdownDescription: "LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
											MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"custom_policies": schema.ListNestedAttribute{
						Description:         "CustomPolicies specifies an array of defined custome policies to be loaded",
						MarkdownDescription: "CustomPolicies specifies an array of defined custome policies to be loaded",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name specifies the name of the custom policy",
									MarkdownDescription: "Name specifies the name of the custom policy",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"secret_ref": schema.SingleNestedAttribute{
									Description:         "SecretRef specifies the secret holding the custom policy metadata and lua code",
									MarkdownDescription: "SecretRef specifies the secret holding the custom policy metadata and lua code",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
											MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},

								"version": schema.StringAttribute{
									Description:         "Version specifies the name of the custom policy",
									MarkdownDescription: "Version specifies the name of the custom policy",
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

					"deployment_environment": schema.StringAttribute{
						Description:         "DeploymentEnvironment is the environment for which the configuration will be downloaded from 3scale (Staging or Production), when using APIcast. The value will also be used in the header X-3scale-User-Agent in the authorize/report requests made to 3scale Service Management API. It is used by 3scale for statistics.",
						MarkdownDescription: "DeploymentEnvironment is the environment for which the configuration will be downloaded from 3scale (Staging or Production), when using APIcast. The value will also be used in the header X-3scale-User-Agent in the authorize/report requests made to 3scale Service Management API. It is used by 3scale for statistics.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"dns_resolver_address": schema.StringAttribute{
						Description:         "DNSResolverAddress can be used to specify a custom DNS resolver address to be used by OpenResty.",
						MarkdownDescription: "DNSResolverAddress can be used to specify a custom DNS resolver address to be used by OpenResty.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"embedded_configuration_secret_ref": schema.SingleNestedAttribute{
						Description:         "Secret reference to a Kubernetes secret containing the gateway configuration. The Secret must be located in the same namespace.",
						MarkdownDescription: "Secret reference to a Kubernetes secret containing the gateway configuration. The Secret must be located in the same namespace.",
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

					"enabled_services": schema.ListAttribute{
						Description:         "EnabledServices can be used to specify a list of service IDs used to filter the configured services.",
						MarkdownDescription: "EnabledServices can be used to specify a list of service IDs used to filter the configured services.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"exposed_host": schema.SingleNestedAttribute{
						Description:         "ExposedHost is the domain name used for external access. By default no external access is configured.",
						MarkdownDescription: "ExposedHost is the domain name used for external access. By default no external access is configured.",
						Attributes: map[string]schema.Attribute{
							"host": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"tls": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"hosts": schema.ListAttribute{
											Description:         "Hosts are a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.",
											MarkdownDescription: "Hosts are a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secret_name": schema.StringAttribute{
											Description:         "SecretName is the name of the secret used to terminate TLS traffic on port 443. Field is left optional to allow TLS routing based on SNI hostname alone. If the SNI host in a listener conflicts with the 'Host' header field used by an IngressRule, the SNI host is used for termination and value of the Host header is used for routing.",
											MarkdownDescription: "SecretName is the name of the secret used to terminate TLS traffic on port 443. Field is left optional to allow TLS routing based on SNI hostname alone. If the SNI host in a listener conflicts with the 'Host' header field used by an IngressRule, the SNI host is used for termination and value of the Host header is used for routing.",
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

					"extended_metrics": schema.BoolAttribute{
						Description:         "ExtendedMetrics enables additional information on Prometheus metrics; some labels will be used with specific information that will provide more in-depth details about APIcast.",
						MarkdownDescription: "ExtendedMetrics enables additional information on Prometheus metrics; some labels will be used with specific information that will provide more in-depth details about APIcast.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"http_proxy": schema.StringAttribute{
						Description:         "HTTPProxy specifies a HTTP(S) Proxy to be used for connecting to HTTP services. Authentication is not supported. Format is <scheme>://<host>:<port>",
						MarkdownDescription: "HTTPProxy specifies a HTTP(S) Proxy to be used for connecting to HTTP services. Authentication is not supported. Format is <scheme>://<host>:<port>",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"https_certificate_secret_ref": schema.SingleNestedAttribute{
						Description:         "HTTPSCertificateSecretRef references secret containing the X.509 certificate in the PEM format and the X.509 certificate secret key.",
						MarkdownDescription: "HTTPSCertificateSecretRef references secret containing the X.509 certificate in the PEM format and the X.509 certificate secret key.",
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

					"https_port": schema.Int64Attribute{
						Description:         "HttpsPort controls on which port APIcast should start listening for HTTPS connections. If this clashes with HTTP port it will be used only for HTTPS.",
						MarkdownDescription: "HttpsPort controls on which port APIcast should start listening for HTTPS connections. If this clashes with HTTP port it will be used only for HTTPS.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"https_proxy": schema.StringAttribute{
						Description:         "HTTPSProxy specifies a HTTP(S) Proxy to be used for connecting to HTTPS services. Authentication is not supported. Format is <scheme>://<host>:<port>",
						MarkdownDescription: "HTTPSProxy specifies a HTTP(S) Proxy to be used for connecting to HTTPS services. Authentication is not supported. Format is <scheme>://<host>:<port>",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"https_verify_depth": schema.Int64Attribute{
						Description:         "HTTPSVerifyDepth defines the maximum length of the client certificate chain.",
						MarkdownDescription: "HTTPSVerifyDepth defines the maximum length of the client certificate chain.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},

					"image": schema.StringAttribute{
						Description:         "Image allows overriding the default APIcast gateway container image. This setting should only be used for dev/testing purposes. Setting this disables automated upgrades of the image.",
						MarkdownDescription: "Image allows overriding the default APIcast gateway container image. This setting should only be used for dev/testing purposes. Setting this disables automated upgrades of the image.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"load_services_when_needed": schema.BoolAttribute{
						Description:         "LoadServicesWhenNeeded makes the configurations to be loaded lazily. APIcast will only load the ones configured for the host specified in the host header of the request.",
						MarkdownDescription: "LoadServicesWhenNeeded makes the configurations to be loaded lazily. APIcast will only load the ones configured for the host specified in the host header of the request.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"log_level": schema.StringAttribute{
						Description:         "LogLevel controls the log level of APIcast's OpenResty logs.",
						MarkdownDescription: "LogLevel controls the log level of APIcast's OpenResty logs.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("debug", "info", "notice", "warn", "error", "crit", "alert", "emerg"),
						},
					},

					"management_api_scope": schema.StringAttribute{
						Description:         "ManagementAPIScope controls APIcast Management API scope. The Management API is powerful and can control the APIcast configuration. debug level should only be enabled for debugging purposes.",
						MarkdownDescription: "ManagementAPIScope controls APIcast Management API scope. The Management API is powerful and can control the APIcast configuration. debug level should only be enabled for debugging purposes.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("disabled", "status", "policies", "debug"),
						},
					},

					"no_proxy": schema.StringAttribute{
						Description:         "NoProxy specifies a comma-separated list of hostnames and domain names for which the requests should not be proxied. Setting to a single * character, which matches all hosts, effectively disables the proxy.",
						MarkdownDescription: "NoProxy specifies a comma-separated list of hostnames and domain names for which the requests should not be proxied. Setting to a single * character, which matches all hosts, effectively disables the proxy.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"oidc_log_level": schema.StringAttribute{
						Description:         "OidcLogLevel allows to set the log level for the logs related to OpenID Connect integration.",
						MarkdownDescription: "OidcLogLevel allows to set the log level for the logs related to OpenID Connect integration.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("debug", "info", "notice", "warn", "error", "crit", "alert", "emerg"),
						},
					},

					"open_ssl_peer_verification_enabled": schema.BoolAttribute{
						Description:         "OpenSSLPeerVerificationEnabled controls OpenSSL peer verification.",
						MarkdownDescription: "OpenSSLPeerVerificationEnabled controls OpenSSL peer verification.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"open_telemetry": schema.SingleNestedAttribute{
						Description:         "OpenTelemetry contains the gateway instrumentation configuration with APIcast.",
						MarkdownDescription: "OpenTelemetry contains the gateway instrumentation configuration with APIcast.",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "Enabled controls whether OpenTelemetry integration with APIcast is enabled. By default it is not enabled.",
								MarkdownDescription: "Enabled controls whether OpenTelemetry integration with APIcast is enabled. By default it is not enabled.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tracing_config_secret_key": schema.StringAttribute{
								Description:         "TracingConfigSecretKey contains the key of the secret to select the configuration from. if unspecified, the first secret key in lexicographical order will be selected.",
								MarkdownDescription: "TracingConfigSecretKey contains the key of the secret to select the configuration from. if unspecified, the first secret key in lexicographical order will be selected.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tracing_config_secret_ref": schema.SingleNestedAttribute{
								Description:         "TracingConfigSecretRef contains a Secret reference the Opentelemetry configuration. The configuration file specification is defined in the Nginx instrumentation library repo https://github.com/open-telemetry/opentelemetry-cpp-contrib/tree/main/instrumentation/nginx",
								MarkdownDescription: "TracingConfigSecretRef contains a Secret reference the Opentelemetry configuration. The configuration file specification is defined in the Nginx instrumentation library repo https://github.com/open-telemetry/opentelemetry-cpp-contrib/tree/main/instrumentation/nginx",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"open_tracing": schema.SingleNestedAttribute{
						Description:         "OpenTracingSpec contains the OpenTracing integration configuration with APIcast. Deprecated",
						MarkdownDescription: "OpenTracingSpec contains the OpenTracing integration configuration with APIcast. Deprecated",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "Enabled controls whether OpenTracing integration with APIcast is enabled. By default it is not enabled.",
								MarkdownDescription: "Enabled controls whether OpenTracing integration with APIcast is enabled. By default it is not enabled.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tracing_config_secret_ref": schema.SingleNestedAttribute{
								Description:         "TracingConfigSecretRef contains a Secret reference the OpenTracing configuration. Each supported tracing library provides a default configuration file that is used if TracingConfig is not specified.",
								MarkdownDescription: "TracingConfigSecretRef contains a Secret reference the OpenTracing configuration. Each supported tracing library provides a default configuration file that is used if TracingConfig is not specified.",
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

							"tracing_library": schema.StringAttribute{
								Description:         "TracingLibrary controls which OpenTracing library is loaded. At the moment the only supported tracer is 'jaeger'. If not set, 'jaeger' will be used.",
								MarkdownDescription: "TracingLibrary controls which OpenTracing library is loaded. At the moment the only supported tracer is 'jaeger'. If not set, 'jaeger' will be used.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"path_routing_enabled": schema.BoolAttribute{
						Description:         "PathRoutingEnabled can be used to enable APIcast's path-based routing in addition to to the default host-based routing.",
						MarkdownDescription: "PathRoutingEnabled can be used to enable APIcast's path-based routing in addition to to the default host-based routing.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"replicas": schema.Int64Attribute{
						Description:         "Number of replicas of the APIcast Deployment.",
						MarkdownDescription: "Number of replicas of the APIcast Deployment.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"resources": schema.SingleNestedAttribute{
						Description:         "Resources can be used to set custom compute Kubernetes Resource Requirements for the APIcast deployment.",
						MarkdownDescription: "Resources can be used to set custom compute Kubernetes Resource Requirements for the APIcast deployment.",
						Attributes: map[string]schema.Attribute{
							"limits": schema.MapAttribute{
								Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"requests": schema.MapAttribute{
								Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

					"response_codes_included": schema.BoolAttribute{
						Description:         "ResponseCodesIncluded can be set to log the response codes of the responses in Apisonator, so they can then be visualized in the 3scale admin portal.",
						MarkdownDescription: "ResponseCodesIncluded can be set to log the response codes of the responses in Apisonator, so they can then be visualized in the 3scale admin portal.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"service_account": schema.StringAttribute{
						Description:         "Kubernetes Service Account name to be used for the APIcast Deployment. The Service Account must exist beforehand.",
						MarkdownDescription: "Kubernetes Service Account name to be used for the APIcast Deployment. The Service Account must exist beforehand.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"service_cache_size": schema.Int64Attribute{
						Description:         "ServiceCacheSize specifies the number of services that APICast can store in the internal cache",
						MarkdownDescription: "ServiceCacheSize specifies the number of services that APICast can store in the internal cache",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"service_configuration_version_override": schema.MapAttribute{
						Description:         "ServiceConfigurationVersionOverride contains service configuration version map to prevent it from auto-updating.",
						MarkdownDescription: "ServiceConfigurationVersionOverride contains service configuration version map to prevent it from auto-updating.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"services_filter_by_url": schema.StringAttribute{
						Description:         "ServicesFilterByURL is used to filter the service configured in the 3scale API Manager, the filter matches with the public base URL (Staging or production).",
						MarkdownDescription: "ServicesFilterByURL is used to filter the service configured in the 3scale API Manager, the filter matches with the public base URL (Staging or production).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"timezone": schema.StringAttribute{
						Description:         "Timezone specifies the local timezone of the APIcast deployment pods. A timezone value available in the TZ database must be set.",
						MarkdownDescription: "Timezone specifies the local timezone of the APIcast deployment pods. A timezone value available in the TZ database must be set.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"upstream_retry_cases": schema.StringAttribute{
						Description:         "UpstreamRetryCases Used only when the retry policy is configured. Specified in which cases a request to the upstream API should be retried.",
						MarkdownDescription: "UpstreamRetryCases Used only when the retry policy is configured. Specified in which cases a request to the upstream API should be retried.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("error", "timeout", "invalid_header", "http_500", "http_502", "http_503", "http_504", "http_403", "http_404", "http_429", "non_idempotent", "off"),
						},
					},

					"workers": schema.Int64Attribute{
						Description:         "Workers defines the number of APIcast's worker processes per pod.",
						MarkdownDescription: "Workers defines the number of APIcast's worker processes per pod.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
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

func (r *Apps3ScaleNetAPIcastV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_apps_3scale_net_ap_icast_v1alpha1_manifest")

	var model Apps3ScaleNetAPIcastV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("apps.3scale.net/v1alpha1")
	model.Kind = pointer.String("APIcast")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"YAML Error: "+err.Error(),
		)
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
