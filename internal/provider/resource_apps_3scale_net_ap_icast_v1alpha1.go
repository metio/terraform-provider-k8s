/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

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

type Apps3ScaleNetAPIcastV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*Apps3ScaleNetAPIcastV1Alpha1Resource)(nil)
)

type Apps3ScaleNetAPIcastV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type Apps3ScaleNetAPIcastV1Alpha1GoModel struct {
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
		AdminPortalCredentialsRef *struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"admin_portal_credentials_ref" yaml:"adminPortalCredentialsRef,omitempty"`

		AllProxy *string `tfsdk:"all_proxy" yaml:"allProxy,omitempty"`

		CacheConfigurationSeconds *int64 `tfsdk:"cache_configuration_seconds" yaml:"cacheConfigurationSeconds,omitempty"`

		CacheMaxTime *string `tfsdk:"cache_max_time" yaml:"cacheMaxTime,omitempty"`

		CacheStatusCodes *string `tfsdk:"cache_status_codes" yaml:"cacheStatusCodes,omitempty"`

		ConfigurationLoadMode *string `tfsdk:"configuration_load_mode" yaml:"configurationLoadMode,omitempty"`

		CustomEnvironments *[]struct {
			SecretRef *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`
		} `tfsdk:"custom_environments" yaml:"customEnvironments,omitempty"`

		CustomPolicies *[]struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			SecretRef *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

			Version *string `tfsdk:"version" yaml:"version,omitempty"`
		} `tfsdk:"custom_policies" yaml:"customPolicies,omitempty"`

		DeploymentEnvironment *string `tfsdk:"deployment_environment" yaml:"deploymentEnvironment,omitempty"`

		DnsResolverAddress *string `tfsdk:"dns_resolver_address" yaml:"dnsResolverAddress,omitempty"`

		EmbeddedConfigurationSecretRef *struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"embedded_configuration_secret_ref" yaml:"embeddedConfigurationSecretRef,omitempty"`

		EnabledServices *[]string `tfsdk:"enabled_services" yaml:"enabledServices,omitempty"`

		ExposedHost *struct {
			Host *string `tfsdk:"host" yaml:"host,omitempty"`

			Tls *[]struct {
				Hosts *[]string `tfsdk:"hosts" yaml:"hosts,omitempty"`

				SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
			} `tfsdk:"tls" yaml:"tls,omitempty"`
		} `tfsdk:"exposed_host" yaml:"exposedHost,omitempty"`

		ExtendedMetrics *bool `tfsdk:"extended_metrics" yaml:"extendedMetrics,omitempty"`

		HttpProxy *string `tfsdk:"http_proxy" yaml:"httpProxy,omitempty"`

		HttpsCertificateSecretRef *struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"https_certificate_secret_ref" yaml:"httpsCertificateSecretRef,omitempty"`

		HttpsPort *int64 `tfsdk:"https_port" yaml:"httpsPort,omitempty"`

		HttpsProxy *string `tfsdk:"https_proxy" yaml:"httpsProxy,omitempty"`

		HttpsVerifyDepth *int64 `tfsdk:"https_verify_depth" yaml:"httpsVerifyDepth,omitempty"`

		Image *string `tfsdk:"image" yaml:"image,omitempty"`

		LoadServicesWhenNeeded *bool `tfsdk:"load_services_when_needed" yaml:"loadServicesWhenNeeded,omitempty"`

		LogLevel *string `tfsdk:"log_level" yaml:"logLevel,omitempty"`

		ManagementAPIScope *string `tfsdk:"management_api_scope" yaml:"managementAPIScope,omitempty"`

		NoProxy *string `tfsdk:"no_proxy" yaml:"noProxy,omitempty"`

		OidcLogLevel *string `tfsdk:"oidc_log_level" yaml:"oidcLogLevel,omitempty"`

		OpenSSLPeerVerificationEnabled *bool `tfsdk:"open_ssl_peer_verification_enabled" yaml:"openSSLPeerVerificationEnabled,omitempty"`

		OpenTracing *struct {
			Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

			TracingConfigSecretRef *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"tracing_config_secret_ref" yaml:"tracingConfigSecretRef,omitempty"`

			TracingLibrary *string `tfsdk:"tracing_library" yaml:"tracingLibrary,omitempty"`
		} `tfsdk:"open_tracing" yaml:"openTracing,omitempty"`

		PathRoutingEnabled *bool `tfsdk:"path_routing_enabled" yaml:"pathRoutingEnabled,omitempty"`

		Replicas *int64 `tfsdk:"replicas" yaml:"replicas,omitempty"`

		Resources *struct {
			Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

			Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
		} `tfsdk:"resources" yaml:"resources,omitempty"`

		ResponseCodesIncluded *bool `tfsdk:"response_codes_included" yaml:"responseCodesIncluded,omitempty"`

		ServiceAccount *string `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

		ServiceConfigurationVersionOverride *map[string]string `tfsdk:"service_configuration_version_override" yaml:"serviceConfigurationVersionOverride,omitempty"`

		ServicesFilterByURL *string `tfsdk:"services_filter_by_url" yaml:"servicesFilterByURL,omitempty"`

		Timezone *string `tfsdk:"timezone" yaml:"timezone,omitempty"`

		UpstreamRetryCases *string `tfsdk:"upstream_retry_cases" yaml:"upstreamRetryCases,omitempty"`

		Workers *int64 `tfsdk:"workers" yaml:"workers,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewApps3ScaleNetAPIcastV1Alpha1Resource() resource.Resource {
	return &Apps3ScaleNetAPIcastV1Alpha1Resource{}
}

func (r *Apps3ScaleNetAPIcastV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_apps_3scale_net_ap_icast_v1alpha1"
}

func (r *Apps3ScaleNetAPIcastV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "APIcast is the Schema for the apicasts API.",
		MarkdownDescription: "APIcast is the Schema for the apicasts API.",
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
				Description:         "APIcastSpec defines the desired state of APIcast.",
				MarkdownDescription: "APIcastSpec defines the desired state of APIcast.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"admin_portal_credentials_ref": {
						Description:         "Secret reference to a Kubernetes Secret containing the admin portal endpoint URL. The Secret must be located in the same namespace.",
						MarkdownDescription: "Secret reference to a Kubernetes Secret containing the admin portal endpoint URL. The Secret must be located in the same namespace.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
								MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

					"all_proxy": {
						Description:         "AllProxy specifies a HTTP(S) proxy to be used for connecting to services if a protocol-specific proxy is not specified. Authentication is not supported. Format is <scheme>://<host>:<port>",
						MarkdownDescription: "AllProxy specifies a HTTP(S) proxy to be used for connecting to services if a protocol-specific proxy is not specified. Authentication is not supported. Format is <scheme>://<host>:<port>",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"cache_configuration_seconds": {
						Description:         "The period (in seconds) that the APIcast configuration will be stored in APIcast's cache.",
						MarkdownDescription: "The period (in seconds) that the APIcast configuration will be stored in APIcast's cache.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"cache_max_time": {
						Description:         "CacheMaxTime indicates the maximum time to be cached. If cache-control header is not set, the time to be cached will be the defined one.",
						MarkdownDescription: "CacheMaxTime indicates the maximum time to be cached. If cache-control header is not set, the time to be cached will be the defined one.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"cache_status_codes": {
						Description:         "CacheStatusCodes defines the status codes for which the response content will be cached.",
						MarkdownDescription: "CacheStatusCodes defines the status codes for which the response content will be cached.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"configuration_load_mode": {
						Description:         "ConfigurationLoadMode can be used to set APIcast's configuration load mode.",
						MarkdownDescription: "ConfigurationLoadMode can be used to set APIcast's configuration load mode.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("boot", "lazy"),
						},
					},

					"custom_environments": {
						Description:         "CustomEnvironments specifies an array of defined custome environments to be loaded",
						MarkdownDescription: "CustomEnvironments specifies an array of defined custome environments to be loaded",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"secret_ref": {
								Description:         "LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.",
								MarkdownDescription: "LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"custom_policies": {
						Description:         "CustomPolicies specifies an array of defined custome policies to be loaded",
						MarkdownDescription: "CustomPolicies specifies an array of defined custome policies to be loaded",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "Name specifies the name of the custom policy",
								MarkdownDescription: "Name specifies the name of the custom policy",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"secret_ref": {
								Description:         "SecretRef specifies the secret holding the custom policy metadata and lua code",
								MarkdownDescription: "SecretRef specifies the secret holding the custom policy metadata and lua code",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: true,
								Optional: false,
								Computed: false,
							},

							"version": {
								Description:         "Version specifies the name of the custom policy",
								MarkdownDescription: "Version specifies the name of the custom policy",

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

					"deployment_environment": {
						Description:         "DeploymentEnvironment is the environment for which the configuration will be downloaded from 3scale (Staging or Production), when using APIcast. The value will also be used in the header X-3scale-User-Agent in the authorize/report requests made to 3scale Service Management API. It is used by 3scale for statistics.",
						MarkdownDescription: "DeploymentEnvironment is the environment for which the configuration will be downloaded from 3scale (Staging or Production), when using APIcast. The value will also be used in the header X-3scale-User-Agent in the authorize/report requests made to 3scale Service Management API. It is used by 3scale for statistics.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"dns_resolver_address": {
						Description:         "DNSResolverAddress can be used to specify a custom DNS resolver address to be used by OpenResty.",
						MarkdownDescription: "DNSResolverAddress can be used to specify a custom DNS resolver address to be used by OpenResty.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"embedded_configuration_secret_ref": {
						Description:         "Secret reference to a Kubernetes secret containing the gateway configuration. The Secret must be located in the same namespace.",
						MarkdownDescription: "Secret reference to a Kubernetes secret containing the gateway configuration. The Secret must be located in the same namespace.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
								MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

					"enabled_services": {
						Description:         "EnabledServices can be used to specify a list of service IDs used to filter the configured services.",
						MarkdownDescription: "EnabledServices can be used to specify a list of service IDs used to filter the configured services.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"exposed_host": {
						Description:         "ExposedHost is the domain name used for external access. By default no external access is configured.",
						MarkdownDescription: "ExposedHost is the domain name used for external access. By default no external access is configured.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"host": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"tls": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"hosts": {
										Description:         "Hosts are a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.",
										MarkdownDescription: "Hosts are a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret_name": {
										Description:         "SecretName is the name of the secret used to terminate TLS traffic on port 443. Field is left optional to allow TLS routing based on SNI hostname alone. If the SNI host in a listener conflicts with the 'Host' header field used by an IngressRule, the SNI host is used for termination and value of the Host header is used for routing.",
										MarkdownDescription: "SecretName is the name of the secret used to terminate TLS traffic on port 443. Field is left optional to allow TLS routing based on SNI hostname alone. If the SNI host in a listener conflicts with the 'Host' header field used by an IngressRule, the SNI host is used for termination and value of the Host header is used for routing.",

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

					"extended_metrics": {
						Description:         "ExtendedMetrics enables additional information on Prometheus metrics; some labels will be used with specific information that will provide more in-depth details about APIcast.",
						MarkdownDescription: "ExtendedMetrics enables additional information on Prometheus metrics; some labels will be used with specific information that will provide more in-depth details about APIcast.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"http_proxy": {
						Description:         "HTTPProxy specifies a HTTP(S) Proxy to be used for connecting to HTTP services. Authentication is not supported. Format is <scheme>://<host>:<port>",
						MarkdownDescription: "HTTPProxy specifies a HTTP(S) Proxy to be used for connecting to HTTP services. Authentication is not supported. Format is <scheme>://<host>:<port>",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"https_certificate_secret_ref": {
						Description:         "HTTPSCertificateSecretRef references secret containing the X.509 certificate in the PEM format and the X.509 certificate secret key.",
						MarkdownDescription: "HTTPSCertificateSecretRef references secret containing the X.509 certificate in the PEM format and the X.509 certificate secret key.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
								MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

					"https_port": {
						Description:         "HttpsPort controls on which port APIcast should start listening for HTTPS connections. If this clashes with HTTP port it will be used only for HTTPS.",
						MarkdownDescription: "HttpsPort controls on which port APIcast should start listening for HTTPS connections. If this clashes with HTTP port it will be used only for HTTPS.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"https_proxy": {
						Description:         "HTTPSProxy specifies a HTTP(S) Proxy to be used for connecting to HTTPS services. Authentication is not supported. Format is <scheme>://<host>:<port>",
						MarkdownDescription: "HTTPSProxy specifies a HTTP(S) Proxy to be used for connecting to HTTPS services. Authentication is not supported. Format is <scheme>://<host>:<port>",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"https_verify_depth": {
						Description:         "HTTPSVerifyDepth defines the maximum length of the client certificate chain.",
						MarkdownDescription: "HTTPSVerifyDepth defines the maximum length of the client certificate chain.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							int64validator.AtLeast(0),
						},
					},

					"image": {
						Description:         "Image allows overriding the default APIcast gateway container image. This setting should only be used for dev/testing purposes. Setting this disables automated upgrades of the image.",
						MarkdownDescription: "Image allows overriding the default APIcast gateway container image. This setting should only be used for dev/testing purposes. Setting this disables automated upgrades of the image.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"load_services_when_needed": {
						Description:         "LoadServicesWhenNeeded makes the configurations to be loaded lazily. APIcast will only load the ones configured for the host specified in the host header of the request.",
						MarkdownDescription: "LoadServicesWhenNeeded makes the configurations to be loaded lazily. APIcast will only load the ones configured for the host specified in the host header of the request.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"log_level": {
						Description:         "LogLevel controls the log level of APIcast's OpenResty logs.",
						MarkdownDescription: "LogLevel controls the log level of APIcast's OpenResty logs.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("debug", "info", "notice", "warn", "error", "crit", "alert", "emerg"),
						},
					},

					"management_api_scope": {
						Description:         "ManagementAPIScope controls APIcast Management API scope. The Management API is powerful and can control the APIcast configuration. debug level should only be enabled for debugging purposes.",
						MarkdownDescription: "ManagementAPIScope controls APIcast Management API scope. The Management API is powerful and can control the APIcast configuration. debug level should only be enabled for debugging purposes.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("disabled", "status", "policies", "debug"),
						},
					},

					"no_proxy": {
						Description:         "NoProxy specifies a comma-separated list of hostnames and domain names for which the requests should not be proxied. Setting to a single * character, which matches all hosts, effectively disables the proxy.",
						MarkdownDescription: "NoProxy specifies a comma-separated list of hostnames and domain names for which the requests should not be proxied. Setting to a single * character, which matches all hosts, effectively disables the proxy.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"oidc_log_level": {
						Description:         "OidcLogLevel allows to set the log level for the logs related to OpenID Connect integration.",
						MarkdownDescription: "OidcLogLevel allows to set the log level for the logs related to OpenID Connect integration.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("debug", "info", "notice", "warn", "error", "crit", "alert", "emerg"),
						},
					},

					"open_ssl_peer_verification_enabled": {
						Description:         "OpenSSLPeerVerificationEnabled controls OpenSSL peer verification.",
						MarkdownDescription: "OpenSSLPeerVerificationEnabled controls OpenSSL peer verification.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"open_tracing": {
						Description:         "OpenTracingSpec contains the OpenTracing integration configuration with APIcast.",
						MarkdownDescription: "OpenTracingSpec contains the OpenTracing integration configuration with APIcast.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"enabled": {
								Description:         "Enabled controls whether OpenTracing integration with APIcast is enabled. By default it is not enabled.",
								MarkdownDescription: "Enabled controls whether OpenTracing integration with APIcast is enabled. By default it is not enabled.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"tracing_config_secret_ref": {
								Description:         "TracingConfigSecretRef contains a Secret reference the OpenTracing configuration. Each supported tracing library provides a default configuration file that is used if TracingConfig is not specified.",
								MarkdownDescription: "TracingConfigSecretRef contains a Secret reference the OpenTracing configuration. Each supported tracing library provides a default configuration file that is used if TracingConfig is not specified.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

							"tracing_library": {
								Description:         "TracingLibrary controls which OpenTracing library is loaded. At the moment the only supported tracer is 'jaeger'. If not set, 'jaeger' will be used.",
								MarkdownDescription: "TracingLibrary controls which OpenTracing library is loaded. At the moment the only supported tracer is 'jaeger'. If not set, 'jaeger' will be used.",

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

					"path_routing_enabled": {
						Description:         "PathRoutingEnabled can be used to enable APIcast's path-based routing in addition to to the default host-based routing.",
						MarkdownDescription: "PathRoutingEnabled can be used to enable APIcast's path-based routing in addition to to the default host-based routing.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"replicas": {
						Description:         "Number of replicas of the APIcast Deployment.",
						MarkdownDescription: "Number of replicas of the APIcast Deployment.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"resources": {
						Description:         "Resources can be used to set custom compute Kubernetes Resource Requirements for the APIcast deployment.",
						MarkdownDescription: "Resources can be used to set custom compute Kubernetes Resource Requirements for the APIcast deployment.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"limits": {
								Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
								MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"requests": {
								Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
								MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",

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

					"response_codes_included": {
						Description:         "ResponseCodesIncluded can be set to log the response codes of the responses in Apisonator, so they can then be visualized in the 3scale admin portal.",
						MarkdownDescription: "ResponseCodesIncluded can be set to log the response codes of the responses in Apisonator, so they can then be visualized in the 3scale admin portal.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"service_account": {
						Description:         "Kubernetes Service Account name to be used for the APIcast Deployment. The Service Account must exist beforehand.",
						MarkdownDescription: "Kubernetes Service Account name to be used for the APIcast Deployment. The Service Account must exist beforehand.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"service_configuration_version_override": {
						Description:         "ServiceConfigurationVersionOverride contains service configuration version map to prevent it from auto-updating.",
						MarkdownDescription: "ServiceConfigurationVersionOverride contains service configuration version map to prevent it from auto-updating.",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"services_filter_by_url": {
						Description:         "ServicesFilterByURL is used to filter the service configured in the 3scale API Manager, the filter matches with the public base URL (Staging or production).",
						MarkdownDescription: "ServicesFilterByURL is used to filter the service configured in the 3scale API Manager, the filter matches with the public base URL (Staging or production).",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"timezone": {
						Description:         "Timezone specifies the local timezone of the APIcast deployment pods. A timezone value available in the TZ database must be set.",
						MarkdownDescription: "Timezone specifies the local timezone of the APIcast deployment pods. A timezone value available in the TZ database must be set.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"upstream_retry_cases": {
						Description:         "UpstreamRetryCases Used only when the retry policy is configured. Specified in which cases a request to the upstream API should be retried.",
						MarkdownDescription: "UpstreamRetryCases Used only when the retry policy is configured. Specified in which cases a request to the upstream API should be retried.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("error", "timeout", "invalid_header", "http_500", "http_502", "http_503", "http_504", "http_403", "http_404", "http_429", "non_idempotent", "off"),
						},
					},

					"workers": {
						Description:         "Workers defines the number of APIcast's worker processes per pod.",
						MarkdownDescription: "Workers defines the number of APIcast's worker processes per pod.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							int64validator.AtLeast(1),
						},
					},
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *Apps3ScaleNetAPIcastV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_apps_3scale_net_ap_icast_v1alpha1")

	var state Apps3ScaleNetAPIcastV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel Apps3ScaleNetAPIcastV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("apps.3scale.net/v1alpha1")
	goModel.Kind = utilities.Ptr("APIcast")

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

func (r *Apps3ScaleNetAPIcastV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_apps_3scale_net_ap_icast_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *Apps3ScaleNetAPIcastV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_apps_3scale_net_ap_icast_v1alpha1")

	var state Apps3ScaleNetAPIcastV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel Apps3ScaleNetAPIcastV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("apps.3scale.net/v1alpha1")
	goModel.Kind = utilities.Ptr("APIcast")

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

func (r *Apps3ScaleNetAPIcastV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_apps_3scale_net_ap_icast_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
