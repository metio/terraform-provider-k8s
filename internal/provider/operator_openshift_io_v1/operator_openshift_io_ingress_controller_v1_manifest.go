/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package operator_openshift_io_v1

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
	_ datasource.DataSource = &OperatorOpenshiftIoIngressControllerV1Manifest{}
)

func NewOperatorOpenshiftIoIngressControllerV1Manifest() datasource.DataSource {
	return &OperatorOpenshiftIoIngressControllerV1Manifest{}
}

type OperatorOpenshiftIoIngressControllerV1Manifest struct{}

type OperatorOpenshiftIoIngressControllerV1ManifestData struct {
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
		ClientTLS *struct {
			AllowedSubjectPatterns *[]string `tfsdk:"allowed_subject_patterns" json:"allowedSubjectPatterns,omitempty"`
			ClientCA               *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"client_ca" json:"clientCA,omitempty"`
			ClientCertificatePolicy *string `tfsdk:"client_certificate_policy" json:"clientCertificatePolicy,omitempty"`
		} `tfsdk:"client_tls" json:"clientTLS,omitempty"`
		DefaultCertificate *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"default_certificate" json:"defaultCertificate,omitempty"`
		Domain                     *string `tfsdk:"domain" json:"domain,omitempty"`
		EndpointPublishingStrategy *struct {
			HostNetwork *struct {
				HttpPort  *int64  `tfsdk:"http_port" json:"httpPort,omitempty"`
				HttpsPort *int64  `tfsdk:"https_port" json:"httpsPort,omitempty"`
				Protocol  *string `tfsdk:"protocol" json:"protocol,omitempty"`
				StatsPort *int64  `tfsdk:"stats_port" json:"statsPort,omitempty"`
			} `tfsdk:"host_network" json:"hostNetwork,omitempty"`
			LoadBalancer *struct {
				AllowedSourceRanges *[]string `tfsdk:"allowed_source_ranges" json:"allowedSourceRanges,omitempty"`
				DnsManagementPolicy *string   `tfsdk:"dns_management_policy" json:"dnsManagementPolicy,omitempty"`
				ProviderParameters  *struct {
					Aws *struct {
						ClassicLoadBalancer *struct {
							ConnectionIdleTimeout *string `tfsdk:"connection_idle_timeout" json:"connectionIdleTimeout,omitempty"`
						} `tfsdk:"classic_load_balancer" json:"classicLoadBalancer,omitempty"`
						NetworkLoadBalancer *map[string]string `tfsdk:"network_load_balancer" json:"networkLoadBalancer,omitempty"`
						Type                *string            `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"aws" json:"aws,omitempty"`
					Gcp *struct {
						ClientAccess *string `tfsdk:"client_access" json:"clientAccess,omitempty"`
					} `tfsdk:"gcp" json:"gcp,omitempty"`
					Ibm *struct {
						Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
					} `tfsdk:"ibm" json:"ibm,omitempty"`
					Type *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"provider_parameters" json:"providerParameters,omitempty"`
				Scope *string `tfsdk:"scope" json:"scope,omitempty"`
			} `tfsdk:"load_balancer" json:"loadBalancer,omitempty"`
			NodePort *struct {
				Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
			} `tfsdk:"node_port" json:"nodePort,omitempty"`
			Private *struct {
				Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
			} `tfsdk:"private" json:"private,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"endpoint_publishing_strategy" json:"endpointPublishingStrategy,omitempty"`
		HttpCompression *struct {
			MimeTypes *[]string `tfsdk:"mime_types" json:"mimeTypes,omitempty"`
		} `tfsdk:"http_compression" json:"httpCompression,omitempty"`
		HttpEmptyRequestsPolicy *string `tfsdk:"http_empty_requests_policy" json:"httpEmptyRequestsPolicy,omitempty"`
		HttpErrorCodePages      *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"http_error_code_pages" json:"httpErrorCodePages,omitempty"`
		HttpHeaders *struct {
			Actions *struct {
				Request *[]struct {
					Action *struct {
						Set *struct {
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"set" json:"set,omitempty"`
						Type *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"action" json:"action,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"request" json:"request,omitempty"`
				Response *[]struct {
					Action *struct {
						Set *struct {
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"set" json:"set,omitempty"`
						Type *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"action" json:"action,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"response" json:"response,omitempty"`
			} `tfsdk:"actions" json:"actions,omitempty"`
			ForwardedHeaderPolicy     *string   `tfsdk:"forwarded_header_policy" json:"forwardedHeaderPolicy,omitempty"`
			HeaderNameCaseAdjustments *[]string `tfsdk:"header_name_case_adjustments" json:"headerNameCaseAdjustments,omitempty"`
			UniqueId                  *struct {
				Format *string `tfsdk:"format" json:"format,omitempty"`
				Name   *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"unique_id" json:"uniqueId,omitempty"`
		} `tfsdk:"http_headers" json:"httpHeaders,omitempty"`
		Logging *struct {
			Access *struct {
				Destination *struct {
					Container *struct {
						MaxLength *int64 `tfsdk:"max_length" json:"maxLength,omitempty"`
					} `tfsdk:"container" json:"container,omitempty"`
					Syslog *struct {
						Address   *string `tfsdk:"address" json:"address,omitempty"`
						Facility  *string `tfsdk:"facility" json:"facility,omitempty"`
						MaxLength *int64  `tfsdk:"max_length" json:"maxLength,omitempty"`
						Port      *int64  `tfsdk:"port" json:"port,omitempty"`
					} `tfsdk:"syslog" json:"syslog,omitempty"`
					Type *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"destination" json:"destination,omitempty"`
				HttpCaptureCookies *[]struct {
					MatchType  *string `tfsdk:"match_type" json:"matchType,omitempty"`
					MaxLength  *int64  `tfsdk:"max_length" json:"maxLength,omitempty"`
					Name       *string `tfsdk:"name" json:"name,omitempty"`
					NamePrefix *string `tfsdk:"name_prefix" json:"namePrefix,omitempty"`
				} `tfsdk:"http_capture_cookies" json:"httpCaptureCookies,omitempty"`
				HttpCaptureHeaders *struct {
					Request *[]struct {
						MaxLength *int64  `tfsdk:"max_length" json:"maxLength,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"request" json:"request,omitempty"`
					Response *[]struct {
						MaxLength *int64  `tfsdk:"max_length" json:"maxLength,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"response" json:"response,omitempty"`
				} `tfsdk:"http_capture_headers" json:"httpCaptureHeaders,omitempty"`
				HttpLogFormat    *string `tfsdk:"http_log_format" json:"httpLogFormat,omitempty"`
				LogEmptyRequests *string `tfsdk:"log_empty_requests" json:"logEmptyRequests,omitempty"`
			} `tfsdk:"access" json:"access,omitempty"`
		} `tfsdk:"logging" json:"logging,omitempty"`
		NamespaceSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
		NodePlacement *struct {
			NodeSelector *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
			Tolerations *[]struct {
				Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
				Key               *string `tfsdk:"key" json:"key,omitempty"`
				Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
				TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
				Value             *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"tolerations" json:"tolerations,omitempty"`
		} `tfsdk:"node_placement" json:"nodePlacement,omitempty"`
		Replicas       *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
		RouteAdmission *struct {
			NamespaceOwnership *string `tfsdk:"namespace_ownership" json:"namespaceOwnership,omitempty"`
			WildcardPolicy     *string `tfsdk:"wildcard_policy" json:"wildcardPolicy,omitempty"`
		} `tfsdk:"route_admission" json:"routeAdmission,omitempty"`
		RouteSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"route_selector" json:"routeSelector,omitempty"`
		TlsSecurityProfile *struct {
			Custom *struct {
				Ciphers       *[]string `tfsdk:"ciphers" json:"ciphers,omitempty"`
				MinTLSVersion *string   `tfsdk:"min_tls_version" json:"minTLSVersion,omitempty"`
			} `tfsdk:"custom" json:"custom,omitempty"`
			Intermediate *map[string]string `tfsdk:"intermediate" json:"intermediate,omitempty"`
			Modern       *map[string]string `tfsdk:"modern" json:"modern,omitempty"`
			Old          *map[string]string `tfsdk:"old" json:"old,omitempty"`
			Type         *string            `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"tls_security_profile" json:"tlsSecurityProfile,omitempty"`
		TuningOptions *struct {
			ClientFinTimeout            *string `tfsdk:"client_fin_timeout" json:"clientFinTimeout,omitempty"`
			ClientTimeout               *string `tfsdk:"client_timeout" json:"clientTimeout,omitempty"`
			HeaderBufferBytes           *int64  `tfsdk:"header_buffer_bytes" json:"headerBufferBytes,omitempty"`
			HeaderBufferMaxRewriteBytes *int64  `tfsdk:"header_buffer_max_rewrite_bytes" json:"headerBufferMaxRewriteBytes,omitempty"`
			HealthCheckInterval         *string `tfsdk:"health_check_interval" json:"healthCheckInterval,omitempty"`
			MaxConnections              *int64  `tfsdk:"max_connections" json:"maxConnections,omitempty"`
			ReloadInterval              *string `tfsdk:"reload_interval" json:"reloadInterval,omitempty"`
			ServerFinTimeout            *string `tfsdk:"server_fin_timeout" json:"serverFinTimeout,omitempty"`
			ServerTimeout               *string `tfsdk:"server_timeout" json:"serverTimeout,omitempty"`
			ThreadCount                 *int64  `tfsdk:"thread_count" json:"threadCount,omitempty"`
			TlsInspectDelay             *string `tfsdk:"tls_inspect_delay" json:"tlsInspectDelay,omitempty"`
			TunnelTimeout               *string `tfsdk:"tunnel_timeout" json:"tunnelTimeout,omitempty"`
		} `tfsdk:"tuning_options" json:"tuningOptions,omitempty"`
		UnsupportedConfigOverrides *map[string]string `tfsdk:"unsupported_config_overrides" json:"unsupportedConfigOverrides,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OperatorOpenshiftIoIngressControllerV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_operator_openshift_io_ingress_controller_v1_manifest"
}

func (r *OperatorOpenshiftIoIngressControllerV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "IngressController describes a managed ingress controller for the cluster. The controller can service OpenShift Route and Kubernetes Ingress resources.  When an IngressController is created, a new ingress controller deployment is created to allow external traffic to reach the services that expose Ingress or Route resources. Updating this resource may lead to disruption for public facing network connections as a new ingress controller revision may be rolled out.  https://kubernetes.io/docs/concepts/services-networking/ingress-controllers  Whenever possible, sensible defaults for the platform are used. See each field for more details.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		MarkdownDescription: "IngressController describes a managed ingress controller for the cluster. The controller can service OpenShift Route and Kubernetes Ingress resources.  When an IngressController is created, a new ingress controller deployment is created to allow external traffic to reach the services that expose Ingress or Route resources. Updating this resource may lead to disruption for public facing network connections as a new ingress controller revision may be rolled out.  https://kubernetes.io/docs/concepts/services-networking/ingress-controllers  Whenever possible, sensible defaults for the platform are used. See each field for more details.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
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
				Description:         "spec is the specification of the desired behavior of the IngressController.",
				MarkdownDescription: "spec is the specification of the desired behavior of the IngressController.",
				Attributes: map[string]schema.Attribute{
					"client_tls": schema.SingleNestedAttribute{
						Description:         "clientTLS specifies settings for requesting and verifying client certificates, which can be used to enable mutual TLS for edge-terminated and reencrypt routes.",
						MarkdownDescription: "clientTLS specifies settings for requesting and verifying client certificates, which can be used to enable mutual TLS for edge-terminated and reencrypt routes.",
						Attributes: map[string]schema.Attribute{
							"allowed_subject_patterns": schema.ListAttribute{
								Description:         "allowedSubjectPatterns specifies a list of regular expressions that should be matched against the distinguished name on a valid client certificate to filter requests.  The regular expressions must use PCRE syntax.  If this list is empty, no filtering is performed.  If the list is nonempty, then at least one pattern must match a client certificate's distinguished name or else the ingress controller rejects the certificate and denies the connection.",
								MarkdownDescription: "allowedSubjectPatterns specifies a list of regular expressions that should be matched against the distinguished name on a valid client certificate to filter requests.  The regular expressions must use PCRE syntax.  If this list is empty, no filtering is performed.  If the list is nonempty, then at least one pattern must match a client certificate's distinguished name or else the ingress controller rejects the certificate and denies the connection.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"client_ca": schema.SingleNestedAttribute{
								Description:         "clientCA specifies a configmap containing the PEM-encoded CA certificate bundle that should be used to verify a client's certificate.  The administrator must create this configmap in the openshift-config namespace.",
								MarkdownDescription: "clientCA specifies a configmap containing the PEM-encoded CA certificate bundle that should be used to verify a client's certificate.  The administrator must create this configmap in the openshift-config namespace.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "name is the metadata.name of the referenced config map",
										MarkdownDescription: "name is the metadata.name of the referenced config map",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: true,
								Optional: false,
								Computed: false,
							},

							"client_certificate_policy": schema.StringAttribute{
								Description:         "clientCertificatePolicy specifies whether the ingress controller requires clients to provide certificates.  This field accepts the values 'Required' or 'Optional'.  Note that the ingress controller only checks client certificates for edge-terminated and reencrypt TLS routes; it cannot check certificates for cleartext HTTP or passthrough TLS routes.",
								MarkdownDescription: "clientCertificatePolicy specifies whether the ingress controller requires clients to provide certificates.  This field accepts the values 'Required' or 'Optional'.  Note that the ingress controller only checks client certificates for edge-terminated and reencrypt TLS routes; it cannot check certificates for cleartext HTTP or passthrough TLS routes.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("", "Required", "Optional"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"default_certificate": schema.SingleNestedAttribute{
						Description:         "defaultCertificate is a reference to a secret containing the default certificate served by the ingress controller. When Routes don't specify their own certificate, defaultCertificate is used.  The secret must contain the following keys and data:  tls.crt: certificate file contents tls.key: key file contents  If unset, a wildcard certificate is automatically generated and used. The certificate is valid for the ingress controller domain (and subdomains) and the generated certificate's CA will be automatically integrated with the cluster's trust store.  If a wildcard certificate is used and shared by multiple HTTP/2 enabled routes (which implies ALPN) then clients (i.e., notably browsers) are at liberty to reuse open connections. This means a client can reuse a connection to another route and that is likely to fail. This behaviour is generally known as connection coalescing.  The in-use certificate (whether generated or user-specified) will be automatically integrated with OpenShift's built-in OAuth server.",
						MarkdownDescription: "defaultCertificate is a reference to a secret containing the default certificate served by the ingress controller. When Routes don't specify their own certificate, defaultCertificate is used.  The secret must contain the following keys and data:  tls.crt: certificate file contents tls.key: key file contents  If unset, a wildcard certificate is automatically generated and used. The certificate is valid for the ingress controller domain (and subdomains) and the generated certificate's CA will be automatically integrated with the cluster's trust store.  If a wildcard certificate is used and shared by multiple HTTP/2 enabled routes (which implies ALPN) then clients (i.e., notably browsers) are at liberty to reuse open connections. This means a client can reuse a connection to another route and that is likely to fail. This behaviour is generally known as connection coalescing.  The in-use certificate (whether generated or user-specified) will be automatically integrated with OpenShift's built-in OAuth server.",
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

					"domain": schema.StringAttribute{
						Description:         "domain is a DNS name serviced by the ingress controller and is used to configure multiple features:  * For the LoadBalancerService endpoint publishing strategy, domain is used to configure DNS records. See endpointPublishingStrategy.  * When using a generated default certificate, the certificate will be valid for domain and its subdomains. See defaultCertificate.  * The value is published to individual Route statuses so that end-users know where to target external DNS records.  domain must be unique among all IngressControllers, and cannot be updated.  If empty, defaults to ingress.config.openshift.io/cluster .spec.domain.",
						MarkdownDescription: "domain is a DNS name serviced by the ingress controller and is used to configure multiple features:  * For the LoadBalancerService endpoint publishing strategy, domain is used to configure DNS records. See endpointPublishingStrategy.  * When using a generated default certificate, the certificate will be valid for domain and its subdomains. See defaultCertificate.  * The value is published to individual Route statuses so that end-users know where to target external DNS records.  domain must be unique among all IngressControllers, and cannot be updated.  If empty, defaults to ingress.config.openshift.io/cluster .spec.domain.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"endpoint_publishing_strategy": schema.SingleNestedAttribute{
						Description:         "endpointPublishingStrategy is used to publish the ingress controller endpoints to other networks, enable load balancer integrations, etc.  If unset, the default is based on infrastructure.config.openshift.io/cluster .status.platform:  AWS:          LoadBalancerService (with External scope) Azure:        LoadBalancerService (with External scope) GCP:          LoadBalancerService (with External scope) IBMCloud:     LoadBalancerService (with External scope) AlibabaCloud: LoadBalancerService (with External scope) Libvirt:      HostNetwork  Any other platform types (including None) default to HostNetwork.  endpointPublishingStrategy cannot be updated.",
						MarkdownDescription: "endpointPublishingStrategy is used to publish the ingress controller endpoints to other networks, enable load balancer integrations, etc.  If unset, the default is based on infrastructure.config.openshift.io/cluster .status.platform:  AWS:          LoadBalancerService (with External scope) Azure:        LoadBalancerService (with External scope) GCP:          LoadBalancerService (with External scope) IBMCloud:     LoadBalancerService (with External scope) AlibabaCloud: LoadBalancerService (with External scope) Libvirt:      HostNetwork  Any other platform types (including None) default to HostNetwork.  endpointPublishingStrategy cannot be updated.",
						Attributes: map[string]schema.Attribute{
							"host_network": schema.SingleNestedAttribute{
								Description:         "hostNetwork holds parameters for the HostNetwork endpoint publishing strategy. Present only if type is HostNetwork.",
								MarkdownDescription: "hostNetwork holds parameters for the HostNetwork endpoint publishing strategy. Present only if type is HostNetwork.",
								Attributes: map[string]schema.Attribute{
									"http_port": schema.Int64Attribute{
										Description:         "httpPort is the port on the host which should be used to listen for HTTP requests. This field should be set when port 80 is already in use. The value should not coincide with the NodePort range of the cluster. When the value is 0 or is not specified it defaults to 80.",
										MarkdownDescription: "httpPort is the port on the host which should be used to listen for HTTP requests. This field should be set when port 80 is already in use. The value should not coincide with the NodePort range of the cluster. When the value is 0 or is not specified it defaults to 80.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
											int64validator.AtMost(65535),
										},
									},

									"https_port": schema.Int64Attribute{
										Description:         "httpsPort is the port on the host which should be used to listen for HTTPS requests. This field should be set when port 443 is already in use. The value should not coincide with the NodePort range of the cluster. When the value is 0 or is not specified it defaults to 443.",
										MarkdownDescription: "httpsPort is the port on the host which should be used to listen for HTTPS requests. This field should be set when port 443 is already in use. The value should not coincide with the NodePort range of the cluster. When the value is 0 or is not specified it defaults to 443.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
											int64validator.AtMost(65535),
										},
									},

									"protocol": schema.StringAttribute{
										Description:         "protocol specifies whether the IngressController expects incoming connections to use plain TCP or whether the IngressController expects PROXY protocol.  PROXY protocol can be used with load balancers that support it to communicate the source addresses of client connections when forwarding those connections to the IngressController.  Using PROXY protocol enables the IngressController to report those source addresses instead of reporting the load balancer's address in HTTP headers and logs.  Note that enabling PROXY protocol on the IngressController will cause connections to fail if you are not using a load balancer that uses PROXY protocol to forward connections to the IngressController.  See http://www.haproxy.org/download/2.2/doc/proxy-protocol.txt for information about PROXY protocol.  The following values are valid for this field:  * The empty string. * 'TCP'. * 'PROXY'.  The empty string specifies the default, which is TCP without PROXY protocol.  Note that the default is subject to change.",
										MarkdownDescription: "protocol specifies whether the IngressController expects incoming connections to use plain TCP or whether the IngressController expects PROXY protocol.  PROXY protocol can be used with load balancers that support it to communicate the source addresses of client connections when forwarding those connections to the IngressController.  Using PROXY protocol enables the IngressController to report those source addresses instead of reporting the load balancer's address in HTTP headers and logs.  Note that enabling PROXY protocol on the IngressController will cause connections to fail if you are not using a load balancer that uses PROXY protocol to forward connections to the IngressController.  See http://www.haproxy.org/download/2.2/doc/proxy-protocol.txt for information about PROXY protocol.  The following values are valid for this field:  * The empty string. * 'TCP'. * 'PROXY'.  The empty string specifies the default, which is TCP without PROXY protocol.  Note that the default is subject to change.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("", "TCP", "PROXY"),
										},
									},

									"stats_port": schema.Int64Attribute{
										Description:         "statsPort is the port on the host where the stats from the router are published. The value should not coincide with the NodePort range of the cluster. If an external load balancer is configured to forward connections to this IngressController, the load balancer should use this port for health checks. The load balancer can send HTTP probes on this port on a given node, with the path /healthz/ready to determine if the ingress controller is ready to receive traffic on the node. For proper operation the load balancer must not forward traffic to a node until the health check reports ready. The load balancer should also stop forwarding requests within a maximum of 45 seconds after /healthz/ready starts reporting not-ready. Probing every 5 to 10 seconds, with a 5-second timeout and with a threshold of two successful or failed requests to become healthy or unhealthy respectively, are well-tested values. When the value is 0 or is not specified it defaults to 1936.",
										MarkdownDescription: "statsPort is the port on the host where the stats from the router are published. The value should not coincide with the NodePort range of the cluster. If an external load balancer is configured to forward connections to this IngressController, the load balancer should use this port for health checks. The load balancer can send HTTP probes on this port on a given node, with the path /healthz/ready to determine if the ingress controller is ready to receive traffic on the node. For proper operation the load balancer must not forward traffic to a node until the health check reports ready. The load balancer should also stop forwarding requests within a maximum of 45 seconds after /healthz/ready starts reporting not-ready. Probing every 5 to 10 seconds, with a 5-second timeout and with a threshold of two successful or failed requests to become healthy or unhealthy respectively, are well-tested values. When the value is 0 or is not specified it defaults to 1936.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
											int64validator.AtMost(65535),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"load_balancer": schema.SingleNestedAttribute{
								Description:         "loadBalancer holds parameters for the load balancer. Present only if type is LoadBalancerService.",
								MarkdownDescription: "loadBalancer holds parameters for the load balancer. Present only if type is LoadBalancerService.",
								Attributes: map[string]schema.Attribute{
									"allowed_source_ranges": schema.ListAttribute{
										Description:         "allowedSourceRanges specifies an allowlist of IP address ranges to which access to the load balancer should be restricted.  Each range must be specified using CIDR notation (e.g. '10.0.0.0/8' or 'fd00::/8'). If no range is specified, '0.0.0.0/0' for IPv4 and '::/0' for IPv6 are used by default, which allows all source addresses.  To facilitate migration from earlier versions of OpenShift that did not have the allowedSourceRanges field, you may set the service.beta.kubernetes.io/load-balancer-source-ranges annotation on the 'router-<ingresscontroller name>' service in the 'openshift-ingress' namespace, and this annotation will take effect if allowedSourceRanges is empty on OpenShift 4.12.",
										MarkdownDescription: "allowedSourceRanges specifies an allowlist of IP address ranges to which access to the load balancer should be restricted.  Each range must be specified using CIDR notation (e.g. '10.0.0.0/8' or 'fd00::/8'). If no range is specified, '0.0.0.0/0' for IPv4 and '::/0' for IPv6 are used by default, which allows all source addresses.  To facilitate migration from earlier versions of OpenShift that did not have the allowedSourceRanges field, you may set the service.beta.kubernetes.io/load-balancer-source-ranges annotation on the 'router-<ingresscontroller name>' service in the 'openshift-ingress' namespace, and this annotation will take effect if allowedSourceRanges is empty on OpenShift 4.12.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"dns_management_policy": schema.StringAttribute{
										Description:         "dnsManagementPolicy indicates if the lifecycle of the wildcard DNS record associated with the load balancer service will be managed by the ingress operator. It defaults to Managed. Valid values are: Managed and Unmanaged.",
										MarkdownDescription: "dnsManagementPolicy indicates if the lifecycle of the wildcard DNS record associated with the load balancer service will be managed by the ingress operator. It defaults to Managed. Valid values are: Managed and Unmanaged.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("Managed", "Unmanaged"),
										},
									},

									"provider_parameters": schema.SingleNestedAttribute{
										Description:         "providerParameters holds desired load balancer information specific to the underlying infrastructure provider.  If empty, defaults will be applied. See specific providerParameters fields for details about their defaults.",
										MarkdownDescription: "providerParameters holds desired load balancer information specific to the underlying infrastructure provider.  If empty, defaults will be applied. See specific providerParameters fields for details about their defaults.",
										Attributes: map[string]schema.Attribute{
											"aws": schema.SingleNestedAttribute{
												Description:         "aws provides configuration settings that are specific to AWS load balancers.  If empty, defaults will be applied. See specific aws fields for details about their defaults.",
												MarkdownDescription: "aws provides configuration settings that are specific to AWS load balancers.  If empty, defaults will be applied. See specific aws fields for details about their defaults.",
												Attributes: map[string]schema.Attribute{
													"classic_load_balancer": schema.SingleNestedAttribute{
														Description:         "classicLoadBalancerParameters holds configuration parameters for an AWS classic load balancer. Present only if type is Classic.",
														MarkdownDescription: "classicLoadBalancerParameters holds configuration parameters for an AWS classic load balancer. Present only if type is Classic.",
														Attributes: map[string]schema.Attribute{
															"connection_idle_timeout": schema.StringAttribute{
																Description:         "connectionIdleTimeout specifies the maximum time period that a connection may be idle before the load balancer closes the connection.  The value must be parseable as a time duration value; see <https://pkg.go.dev/time#ParseDuration>.  A nil or zero value means no opinion, in which case a default value is used.  The default value for this field is 60s.  This default is subject to change.",
																MarkdownDescription: "connectionIdleTimeout specifies the maximum time period that a connection may be idle before the load balancer closes the connection.  The value must be parseable as a time duration value; see <https://pkg.go.dev/time#ParseDuration>.  A nil or zero value means no opinion, in which case a default value is used.  The default value for this field is 60s.  This default is subject to change.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"network_load_balancer": schema.MapAttribute{
														Description:         "networkLoadBalancerParameters holds configuration parameters for an AWS network load balancer. Present only if type is NLB.",
														MarkdownDescription: "networkLoadBalancerParameters holds configuration parameters for an AWS network load balancer. Present only if type is NLB.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "type is the type of AWS load balancer to instantiate for an ingresscontroller.  Valid values are:  * 'Classic': A Classic Load Balancer that makes routing decisions at either the transport layer (TCP/SSL) or the application layer (HTTP/HTTPS). See the following for additional details:  https://docs.aws.amazon.com/AmazonECS/latest/developerguide/load-balancer-types.html#clb  * 'NLB': A Network Load Balancer that makes routing decisions at the transport layer (TCP/SSL). See the following for additional details:  https://docs.aws.amazon.com/AmazonECS/latest/developerguide/load-balancer-types.html#nlb",
														MarkdownDescription: "type is the type of AWS load balancer to instantiate for an ingresscontroller.  Valid values are:  * 'Classic': A Classic Load Balancer that makes routing decisions at either the transport layer (TCP/SSL) or the application layer (HTTP/HTTPS). See the following for additional details:  https://docs.aws.amazon.com/AmazonECS/latest/developerguide/load-balancer-types.html#clb  * 'NLB': A Network Load Balancer that makes routing decisions at the transport layer (TCP/SSL). See the following for additional details:  https://docs.aws.amazon.com/AmazonECS/latest/developerguide/load-balancer-types.html#nlb",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("Classic", "NLB"),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"gcp": schema.SingleNestedAttribute{
												Description:         "gcp provides configuration settings that are specific to GCP load balancers.  If empty, defaults will be applied. See specific gcp fields for details about their defaults.",
												MarkdownDescription: "gcp provides configuration settings that are specific to GCP load balancers.  If empty, defaults will be applied. See specific gcp fields for details about their defaults.",
												Attributes: map[string]schema.Attribute{
													"client_access": schema.StringAttribute{
														Description:         "clientAccess describes how client access is restricted for internal load balancers.  Valid values are: * 'Global': Specifying an internal load balancer with Global client access allows clients from any region within the VPC to communicate with the load balancer.  https://cloud.google.com/kubernetes-engine/docs/how-to/internal-load-balancing#global_access  * 'Local': Specifying an internal load balancer with Local client access means only clients within the same region (and VPC) as the GCP load balancer can communicate with the load balancer. Note that this is the default behavior.  https://cloud.google.com/load-balancing/docs/internal#client_access",
														MarkdownDescription: "clientAccess describes how client access is restricted for internal load balancers.  Valid values are: * 'Global': Specifying an internal load balancer with Global client access allows clients from any region within the VPC to communicate with the load balancer.  https://cloud.google.com/kubernetes-engine/docs/how-to/internal-load-balancing#global_access  * 'Local': Specifying an internal load balancer with Local client access means only clients within the same region (and VPC) as the GCP load balancer can communicate with the load balancer. Note that this is the default behavior.  https://cloud.google.com/load-balancing/docs/internal#client_access",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("Global", "Local"),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"ibm": schema.SingleNestedAttribute{
												Description:         "ibm provides configuration settings that are specific to IBM Cloud load balancers.  If empty, defaults will be applied. See specific ibm fields for details about their defaults.",
												MarkdownDescription: "ibm provides configuration settings that are specific to IBM Cloud load balancers.  If empty, defaults will be applied. See specific ibm fields for details about their defaults.",
												Attributes: map[string]schema.Attribute{
													"protocol": schema.StringAttribute{
														Description:         "protocol specifies whether the load balancer uses PROXY protocol to forward connections to the IngressController. See 'service.kubernetes.io/ibm-load-balancer-cloud-provider-enable-features: 'proxy-protocol'' at https://cloud.ibm.com/docs/containers?topic=containers-vpc-lbaas'  PROXY protocol can be used with load balancers that support it to communicate the source addresses of client connections when forwarding those connections to the IngressController.  Using PROXY protocol enables the IngressController to report those source addresses instead of reporting the load balancer's address in HTTP headers and logs.  Note that enabling PROXY protocol on the IngressController will cause connections to fail if you are not using a load balancer that uses PROXY protocol to forward connections to the IngressController.  See http://www.haproxy.org/download/2.2/doc/proxy-protocol.txt for information about PROXY protocol.  Valid values for protocol are TCP, PROXY and omitted. When omitted, this means no opinion and the platform is left to choose a reasonable default, which is subject to change over time. The current default is TCP, without the proxy protocol enabled.",
														MarkdownDescription: "protocol specifies whether the load balancer uses PROXY protocol to forward connections to the IngressController. See 'service.kubernetes.io/ibm-load-balancer-cloud-provider-enable-features: 'proxy-protocol'' at https://cloud.ibm.com/docs/containers?topic=containers-vpc-lbaas'  PROXY protocol can be used with load balancers that support it to communicate the source addresses of client connections when forwarding those connections to the IngressController.  Using PROXY protocol enables the IngressController to report those source addresses instead of reporting the load balancer's address in HTTP headers and logs.  Note that enabling PROXY protocol on the IngressController will cause connections to fail if you are not using a load balancer that uses PROXY protocol to forward connections to the IngressController.  See http://www.haproxy.org/download/2.2/doc/proxy-protocol.txt for information about PROXY protocol.  Valid values for protocol are TCP, PROXY and omitted. When omitted, this means no opinion and the platform is left to choose a reasonable default, which is subject to change over time. The current default is TCP, without the proxy protocol enabled.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("", "TCP", "PROXY"),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"type": schema.StringAttribute{
												Description:         "type is the underlying infrastructure provider for the load balancer. Allowed values are 'AWS', 'Azure', 'BareMetal', 'GCP', 'IBM', 'Nutanix', 'OpenStack', and 'VSphere'.",
												MarkdownDescription: "type is the underlying infrastructure provider for the load balancer. Allowed values are 'AWS', 'Azure', 'BareMetal', 'GCP', 'IBM', 'Nutanix', 'OpenStack', and 'VSphere'.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("AWS", "Azure", "BareMetal", "GCP", "Nutanix", "OpenStack", "VSphere", "IBM"),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"scope": schema.StringAttribute{
										Description:         "scope indicates the scope at which the load balancer is exposed. Possible values are 'External' and 'Internal'.",
										MarkdownDescription: "scope indicates the scope at which the load balancer is exposed. Possible values are 'External' and 'Internal'.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("Internal", "External"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"node_port": schema.SingleNestedAttribute{
								Description:         "nodePort holds parameters for the NodePortService endpoint publishing strategy. Present only if type is NodePortService.",
								MarkdownDescription: "nodePort holds parameters for the NodePortService endpoint publishing strategy. Present only if type is NodePortService.",
								Attributes: map[string]schema.Attribute{
									"protocol": schema.StringAttribute{
										Description:         "protocol specifies whether the IngressController expects incoming connections to use plain TCP or whether the IngressController expects PROXY protocol.  PROXY protocol can be used with load balancers that support it to communicate the source addresses of client connections when forwarding those connections to the IngressController.  Using PROXY protocol enables the IngressController to report those source addresses instead of reporting the load balancer's address in HTTP headers and logs.  Note that enabling PROXY protocol on the IngressController will cause connections to fail if you are not using a load balancer that uses PROXY protocol to forward connections to the IngressController.  See http://www.haproxy.org/download/2.2/doc/proxy-protocol.txt for information about PROXY protocol.  The following values are valid for this field:  * The empty string. * 'TCP'. * 'PROXY'.  The empty string specifies the default, which is TCP without PROXY protocol.  Note that the default is subject to change.",
										MarkdownDescription: "protocol specifies whether the IngressController expects incoming connections to use plain TCP or whether the IngressController expects PROXY protocol.  PROXY protocol can be used with load balancers that support it to communicate the source addresses of client connections when forwarding those connections to the IngressController.  Using PROXY protocol enables the IngressController to report those source addresses instead of reporting the load balancer's address in HTTP headers and logs.  Note that enabling PROXY protocol on the IngressController will cause connections to fail if you are not using a load balancer that uses PROXY protocol to forward connections to the IngressController.  See http://www.haproxy.org/download/2.2/doc/proxy-protocol.txt for information about PROXY protocol.  The following values are valid for this field:  * The empty string. * 'TCP'. * 'PROXY'.  The empty string specifies the default, which is TCP without PROXY protocol.  Note that the default is subject to change.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("", "TCP", "PROXY"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"private": schema.SingleNestedAttribute{
								Description:         "private holds parameters for the Private endpoint publishing strategy. Present only if type is Private.",
								MarkdownDescription: "private holds parameters for the Private endpoint publishing strategy. Present only if type is Private.",
								Attributes: map[string]schema.Attribute{
									"protocol": schema.StringAttribute{
										Description:         "protocol specifies whether the IngressController expects incoming connections to use plain TCP or whether the IngressController expects PROXY protocol.  PROXY protocol can be used with load balancers that support it to communicate the source addresses of client connections when forwarding those connections to the IngressController.  Using PROXY protocol enables the IngressController to report those source addresses instead of reporting the load balancer's address in HTTP headers and logs.  Note that enabling PROXY protocol on the IngressController will cause connections to fail if you are not using a load balancer that uses PROXY protocol to forward connections to the IngressController.  See http://www.haproxy.org/download/2.2/doc/proxy-protocol.txt for information about PROXY protocol.  The following values are valid for this field:  * The empty string. * 'TCP'. * 'PROXY'.  The empty string specifies the default, which is TCP without PROXY protocol.  Note that the default is subject to change.",
										MarkdownDescription: "protocol specifies whether the IngressController expects incoming connections to use plain TCP or whether the IngressController expects PROXY protocol.  PROXY protocol can be used with load balancers that support it to communicate the source addresses of client connections when forwarding those connections to the IngressController.  Using PROXY protocol enables the IngressController to report those source addresses instead of reporting the load balancer's address in HTTP headers and logs.  Note that enabling PROXY protocol on the IngressController will cause connections to fail if you are not using a load balancer that uses PROXY protocol to forward connections to the IngressController.  See http://www.haproxy.org/download/2.2/doc/proxy-protocol.txt for information about PROXY protocol.  The following values are valid for this field:  * The empty string. * 'TCP'. * 'PROXY'.  The empty string specifies the default, which is TCP without PROXY protocol.  Note that the default is subject to change.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("", "TCP", "PROXY"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"type": schema.StringAttribute{
								Description:         "type is the publishing strategy to use. Valid values are:  * LoadBalancerService  Publishes the ingress controller using a Kubernetes LoadBalancer Service.  In this configuration, the ingress controller deployment uses container networking. A LoadBalancer Service is created to publish the deployment.  See: https://kubernetes.io/docs/concepts/services-networking/service/#loadbalancer  If domain is set, a wildcard DNS record will be managed to point at the LoadBalancer Service's external name. DNS records are managed only in DNS zones defined by dns.config.openshift.io/cluster .spec.publicZone and .spec.privateZone.  Wildcard DNS management is currently supported only on the AWS, Azure, and GCP platforms.  * HostNetwork  Publishes the ingress controller on node ports where the ingress controller is deployed.  In this configuration, the ingress controller deployment uses host networking, bound to node ports 80 and 443. The user is responsible for configuring an external load balancer to publish the ingress controller via the node ports.  * Private  Does not publish the ingress controller.  In this configuration, the ingress controller deployment uses container networking, and is not explicitly published. The user must manually publish the ingress controller.  * NodePortService  Publishes the ingress controller using a Kubernetes NodePort Service.  In this configuration, the ingress controller deployment uses container networking. A NodePort Service is created to publish the deployment. The specific node ports are dynamically allocated by OpenShift; however, to support static port allocations, user changes to the node port field of the managed NodePort Service will preserved.",
								MarkdownDescription: "type is the publishing strategy to use. Valid values are:  * LoadBalancerService  Publishes the ingress controller using a Kubernetes LoadBalancer Service.  In this configuration, the ingress controller deployment uses container networking. A LoadBalancer Service is created to publish the deployment.  See: https://kubernetes.io/docs/concepts/services-networking/service/#loadbalancer  If domain is set, a wildcard DNS record will be managed to point at the LoadBalancer Service's external name. DNS records are managed only in DNS zones defined by dns.config.openshift.io/cluster .spec.publicZone and .spec.privateZone.  Wildcard DNS management is currently supported only on the AWS, Azure, and GCP platforms.  * HostNetwork  Publishes the ingress controller on node ports where the ingress controller is deployed.  In this configuration, the ingress controller deployment uses host networking, bound to node ports 80 and 443. The user is responsible for configuring an external load balancer to publish the ingress controller via the node ports.  * Private  Does not publish the ingress controller.  In this configuration, the ingress controller deployment uses container networking, and is not explicitly published. The user must manually publish the ingress controller.  * NodePortService  Publishes the ingress controller using a Kubernetes NodePort Service.  In this configuration, the ingress controller deployment uses container networking. A NodePort Service is created to publish the deployment. The specific node ports are dynamically allocated by OpenShift; however, to support static port allocations, user changes to the node port field of the managed NodePort Service will preserved.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("LoadBalancerService", "HostNetwork", "Private", "NodePortService"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"http_compression": schema.SingleNestedAttribute{
						Description:         "httpCompression defines a policy for HTTP traffic compression. By default, there is no HTTP compression.",
						MarkdownDescription: "httpCompression defines a policy for HTTP traffic compression. By default, there is no HTTP compression.",
						Attributes: map[string]schema.Attribute{
							"mime_types": schema.ListAttribute{
								Description:         "mimeTypes is a list of MIME types that should have compression applied. This list can be empty, in which case the ingress controller does not apply compression.  Note: Not all MIME types benefit from compression, but HAProxy will still use resources to try to compress if instructed to.  Generally speaking, text (html, css, js, etc.) formats benefit from compression, but formats that are already compressed (image, audio, video, etc.) benefit little in exchange for the time and cpu spent on compressing again. See https://joehonton.medium.com/the-gzip-penalty-d31bd697f1a2",
								MarkdownDescription: "mimeTypes is a list of MIME types that should have compression applied. This list can be empty, in which case the ingress controller does not apply compression.  Note: Not all MIME types benefit from compression, but HAProxy will still use resources to try to compress if instructed to.  Generally speaking, text (html, css, js, etc.) formats benefit from compression, but formats that are already compressed (image, audio, video, etc.) benefit little in exchange for the time and cpu spent on compressing again. See https://joehonton.medium.com/the-gzip-penalty-d31bd697f1a2",
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

					"http_empty_requests_policy": schema.StringAttribute{
						Description:         "httpEmptyRequestsPolicy describes how HTTP connections should be handled if the connection times out before a request is received. Allowed values for this field are 'Respond' and 'Ignore'.  If the field is set to 'Respond', the ingress controller sends an HTTP 400 or 408 response, logs the connection (if access logging is enabled), and counts the connection in the appropriate metrics.  If the field is set to 'Ignore', the ingress controller closes the connection without sending a response, logging the connection, or incrementing metrics.  The default value is 'Respond'.  Typically, these connections come from load balancers' health probes or Web browsers' speculative connections ('preconnect') and can be safely ignored.  However, these requests may also be caused by network errors, and so setting this field to 'Ignore' may impede detection and diagnosis of problems.  In addition, these requests may be caused by port scans, in which case logging empty requests may aid in detecting intrusion attempts.",
						MarkdownDescription: "httpEmptyRequestsPolicy describes how HTTP connections should be handled if the connection times out before a request is received. Allowed values for this field are 'Respond' and 'Ignore'.  If the field is set to 'Respond', the ingress controller sends an HTTP 400 or 408 response, logs the connection (if access logging is enabled), and counts the connection in the appropriate metrics.  If the field is set to 'Ignore', the ingress controller closes the connection without sending a response, logging the connection, or incrementing metrics.  The default value is 'Respond'.  Typically, these connections come from load balancers' health probes or Web browsers' speculative connections ('preconnect') and can be safely ignored.  However, these requests may also be caused by network errors, and so setting this field to 'Ignore' may impede detection and diagnosis of problems.  In addition, these requests may be caused by port scans, in which case logging empty requests may aid in detecting intrusion attempts.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Respond", "Ignore"),
						},
					},

					"http_error_code_pages": schema.SingleNestedAttribute{
						Description:         "httpErrorCodePages specifies a configmap with custom error pages. The administrator must create this configmap in the openshift-config namespace. This configmap should have keys in the format 'error-page-<error code>.http', where <error code> is an HTTP error code. For example, 'error-page-503.http' defines an error page for HTTP 503 responses. Currently only error pages for 503 and 404 responses can be customized. Each value in the configmap should be the full response, including HTTP headers. Eg- https://raw.githubusercontent.com/openshift/router/fadab45747a9b30cc3f0a4b41ad2871f95827a93/images/router/haproxy/conf/error-page-503.http If this field is empty, the ingress controller uses the default error pages.",
						MarkdownDescription: "httpErrorCodePages specifies a configmap with custom error pages. The administrator must create this configmap in the openshift-config namespace. This configmap should have keys in the format 'error-page-<error code>.http', where <error code> is an HTTP error code. For example, 'error-page-503.http' defines an error page for HTTP 503 responses. Currently only error pages for 503 and 404 responses can be customized. Each value in the configmap should be the full response, including HTTP headers. Eg- https://raw.githubusercontent.com/openshift/router/fadab45747a9b30cc3f0a4b41ad2871f95827a93/images/router/haproxy/conf/error-page-503.http If this field is empty, the ingress controller uses the default error pages.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "name is the metadata.name of the referenced config map",
								MarkdownDescription: "name is the metadata.name of the referenced config map",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"http_headers": schema.SingleNestedAttribute{
						Description:         "httpHeaders defines policy for HTTP headers.  If this field is empty, the default values are used.",
						MarkdownDescription: "httpHeaders defines policy for HTTP headers.  If this field is empty, the default values are used.",
						Attributes: map[string]schema.Attribute{
							"actions": schema.SingleNestedAttribute{
								Description:         "actions specifies options for modifying headers and their values. Note that this option only applies to cleartext HTTP connections and to secure HTTP connections for which the ingress controller terminates encryption (that is, edge-terminated or reencrypt connections).  Headers cannot be modified for TLS passthrough connections. Setting the HSTS ('Strict-Transport-Security') header is not supported via actions. 'Strict-Transport-Security' may only be configured using the 'haproxy.router.openshift.io/hsts_header' route annotation, and only in accordance with the policy specified in Ingress.Spec.RequiredHSTSPolicies. Any actions defined here are applied after any actions related to the following other fields: cache-control, spec.clientTLS, spec.httpHeaders.forwardedHeaderPolicy, spec.httpHeaders.uniqueId, and spec.httpHeaders.headerNameCaseAdjustments. In case of HTTP request headers, the actions specified in spec.httpHeaders.actions on the Route will be executed after the actions specified in the IngressController's spec.httpHeaders.actions field. In case of HTTP response headers, the actions specified in spec.httpHeaders.actions on the IngressController will be executed after the actions specified in the Route's spec.httpHeaders.actions field. Headers set using this API cannot be captured for use in access logs. The following header names are reserved and may not be modified via this API: Strict-Transport-Security, Proxy, Host, Cookie, Set-Cookie. Note that the total size of all net added headers *after* interpolating dynamic values must not exceed the value of spec.tuningOptions.headerBufferMaxRewriteBytes on the IngressController. Please refer to the documentation for that API field for more details.",
								MarkdownDescription: "actions specifies options for modifying headers and their values. Note that this option only applies to cleartext HTTP connections and to secure HTTP connections for which the ingress controller terminates encryption (that is, edge-terminated or reencrypt connections).  Headers cannot be modified for TLS passthrough connections. Setting the HSTS ('Strict-Transport-Security') header is not supported via actions. 'Strict-Transport-Security' may only be configured using the 'haproxy.router.openshift.io/hsts_header' route annotation, and only in accordance with the policy specified in Ingress.Spec.RequiredHSTSPolicies. Any actions defined here are applied after any actions related to the following other fields: cache-control, spec.clientTLS, spec.httpHeaders.forwardedHeaderPolicy, spec.httpHeaders.uniqueId, and spec.httpHeaders.headerNameCaseAdjustments. In case of HTTP request headers, the actions specified in spec.httpHeaders.actions on the Route will be executed after the actions specified in the IngressController's spec.httpHeaders.actions field. In case of HTTP response headers, the actions specified in spec.httpHeaders.actions on the IngressController will be executed after the actions specified in the Route's spec.httpHeaders.actions field. Headers set using this API cannot be captured for use in access logs. The following header names are reserved and may not be modified via this API: Strict-Transport-Security, Proxy, Host, Cookie, Set-Cookie. Note that the total size of all net added headers *after* interpolating dynamic values must not exceed the value of spec.tuningOptions.headerBufferMaxRewriteBytes on the IngressController. Please refer to the documentation for that API field for more details.",
								Attributes: map[string]schema.Attribute{
									"request": schema.ListNestedAttribute{
										Description:         "request is a list of HTTP request headers to modify. Actions defined here will modify the request headers of all requests passing through an ingress controller. These actions are applied to all Routes i.e. for all connections handled by the ingress controller defined within a cluster. IngressController actions for request headers will be executed before Route actions. Currently, actions may define to either 'Set' or 'Delete' headers values. Actions are applied in sequence as defined in this list. A maximum of 20 request header actions may be configured. Sample fetchers allowed are 'req.hdr' and 'ssl_c_der'. Converters allowed are 'lower' and 'base64'. Example header values: '%[req.hdr(X-target),lower]', '%{+Q}[ssl_c_der,base64]'.",
										MarkdownDescription: "request is a list of HTTP request headers to modify. Actions defined here will modify the request headers of all requests passing through an ingress controller. These actions are applied to all Routes i.e. for all connections handled by the ingress controller defined within a cluster. IngressController actions for request headers will be executed before Route actions. Currently, actions may define to either 'Set' or 'Delete' headers values. Actions are applied in sequence as defined in this list. A maximum of 20 request header actions may be configured. Sample fetchers allowed are 'req.hdr' and 'ssl_c_der'. Converters allowed are 'lower' and 'base64'. Example header values: '%[req.hdr(X-target),lower]', '%{+Q}[ssl_c_der,base64]'.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"action": schema.SingleNestedAttribute{
													Description:         "action specifies actions to perform on headers, such as setting or deleting headers.",
													MarkdownDescription: "action specifies actions to perform on headers, such as setting or deleting headers.",
													Attributes: map[string]schema.Attribute{
														"set": schema.SingleNestedAttribute{
															Description:         "set specifies how the HTTP header should be set. This field is required when type is Set and forbidden otherwise.",
															MarkdownDescription: "set specifies how the HTTP header should be set. This field is required when type is Set and forbidden otherwise.",
															Attributes: map[string]schema.Attribute{
																"value": schema.StringAttribute{
																	Description:         "value specifies a header value. Dynamic values can be added. The value will be interpreted as an HAProxy format string as defined in http://cbonte.github.io/haproxy-dconv/2.6/configuration.html#8.2.6  and may use HAProxy's %[] syntax and otherwise must be a valid HTTP header value as defined in https://datatracker.ietf.org/doc/html/rfc7230#section-3.2. The value of this field must be no more than 16384 characters in length. Note that the total size of all net added headers *after* interpolating dynamic values must not exceed the value of spec.tuningOptions.headerBufferMaxRewriteBytes on the IngressController.",
																	MarkdownDescription: "value specifies a header value. Dynamic values can be added. The value will be interpreted as an HAProxy format string as defined in http://cbonte.github.io/haproxy-dconv/2.6/configuration.html#8.2.6  and may use HAProxy's %[] syntax and otherwise must be a valid HTTP header value as defined in https://datatracker.ietf.org/doc/html/rfc7230#section-3.2. The value of this field must be no more than 16384 characters in length. Note that the total size of all net added headers *after* interpolating dynamic values must not exceed the value of spec.tuningOptions.headerBufferMaxRewriteBytes on the IngressController.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtLeast(1),
																		stringvalidator.LengthAtMost(16384),
																	},
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"type": schema.StringAttribute{
															Description:         "type defines the type of the action to be applied on the header. Possible values are Set or Delete. Set allows you to set HTTP request and response headers. Delete allows you to delete HTTP request and response headers.",
															MarkdownDescription: "type defines the type of the action to be applied on the header. Possible values are Set or Delete. Set allows you to set HTTP request and response headers. Delete allows you to delete HTTP request and response headers.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("Set", "Delete"),
															},
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},

												"name": schema.StringAttribute{
													Description:         "name specifies the name of a header on which to perform an action. Its value must be a valid HTTP header name as defined in RFC 2616 section 4.2. The name must consist only of alphanumeric and the following special characters, '-!#$%&'*+.^_''. The following header names are reserved and may not be modified via this API: Strict-Transport-Security, Proxy, Host, Cookie, Set-Cookie. It must be no more than 255 characters in length. Header name must be unique.",
													MarkdownDescription: "name specifies the name of a header on which to perform an action. Its value must be a valid HTTP header name as defined in RFC 2616 section 4.2. The name must consist only of alphanumeric and the following special characters, '-!#$%&'*+.^_''. The following header names are reserved and may not be modified via this API: Strict-Transport-Security, Proxy, Host, Cookie, Set-Cookie. It must be no more than 255 characters in length. Header name must be unique.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
														stringvalidator.LengthAtMost(255),
														stringvalidator.RegexMatches(regexp.MustCompile(`^[-!#$%&'*+.0-9A-Z^_`+"`"+`a-z|~]+$`), ""),
													},
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"response": schema.ListNestedAttribute{
										Description:         "response is a list of HTTP response headers to modify. Actions defined here will modify the response headers of all requests passing through an ingress controller. These actions are applied to all Routes i.e. for all connections handled by the ingress controller defined within a cluster. IngressController actions for response headers will be executed after Route actions. Currently, actions may define to either 'Set' or 'Delete' headers values. Actions are applied in sequence as defined in this list. A maximum of 20 response header actions may be configured. Sample fetchers allowed are 'res.hdr' and 'ssl_c_der'. Converters allowed are 'lower' and 'base64'. Example header values: '%[res.hdr(X-target),lower]', '%{+Q}[ssl_c_der,base64]'.",
										MarkdownDescription: "response is a list of HTTP response headers to modify. Actions defined here will modify the response headers of all requests passing through an ingress controller. These actions are applied to all Routes i.e. for all connections handled by the ingress controller defined within a cluster. IngressController actions for response headers will be executed after Route actions. Currently, actions may define to either 'Set' or 'Delete' headers values. Actions are applied in sequence as defined in this list. A maximum of 20 response header actions may be configured. Sample fetchers allowed are 'res.hdr' and 'ssl_c_der'. Converters allowed are 'lower' and 'base64'. Example header values: '%[res.hdr(X-target),lower]', '%{+Q}[ssl_c_der,base64]'.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"action": schema.SingleNestedAttribute{
													Description:         "action specifies actions to perform on headers, such as setting or deleting headers.",
													MarkdownDescription: "action specifies actions to perform on headers, such as setting or deleting headers.",
													Attributes: map[string]schema.Attribute{
														"set": schema.SingleNestedAttribute{
															Description:         "set specifies how the HTTP header should be set. This field is required when type is Set and forbidden otherwise.",
															MarkdownDescription: "set specifies how the HTTP header should be set. This field is required when type is Set and forbidden otherwise.",
															Attributes: map[string]schema.Attribute{
																"value": schema.StringAttribute{
																	Description:         "value specifies a header value. Dynamic values can be added. The value will be interpreted as an HAProxy format string as defined in http://cbonte.github.io/haproxy-dconv/2.6/configuration.html#8.2.6  and may use HAProxy's %[] syntax and otherwise must be a valid HTTP header value as defined in https://datatracker.ietf.org/doc/html/rfc7230#section-3.2. The value of this field must be no more than 16384 characters in length. Note that the total size of all net added headers *after* interpolating dynamic values must not exceed the value of spec.tuningOptions.headerBufferMaxRewriteBytes on the IngressController.",
																	MarkdownDescription: "value specifies a header value. Dynamic values can be added. The value will be interpreted as an HAProxy format string as defined in http://cbonte.github.io/haproxy-dconv/2.6/configuration.html#8.2.6  and may use HAProxy's %[] syntax and otherwise must be a valid HTTP header value as defined in https://datatracker.ietf.org/doc/html/rfc7230#section-3.2. The value of this field must be no more than 16384 characters in length. Note that the total size of all net added headers *after* interpolating dynamic values must not exceed the value of spec.tuningOptions.headerBufferMaxRewriteBytes on the IngressController.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtLeast(1),
																		stringvalidator.LengthAtMost(16384),
																	},
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"type": schema.StringAttribute{
															Description:         "type defines the type of the action to be applied on the header. Possible values are Set or Delete. Set allows you to set HTTP request and response headers. Delete allows you to delete HTTP request and response headers.",
															MarkdownDescription: "type defines the type of the action to be applied on the header. Possible values are Set or Delete. Set allows you to set HTTP request and response headers. Delete allows you to delete HTTP request and response headers.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("Set", "Delete"),
															},
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},

												"name": schema.StringAttribute{
													Description:         "name specifies the name of a header on which to perform an action. Its value must be a valid HTTP header name as defined in RFC 2616 section 4.2. The name must consist only of alphanumeric and the following special characters, '-!#$%&'*+.^_''. The following header names are reserved and may not be modified via this API: Strict-Transport-Security, Proxy, Host, Cookie, Set-Cookie. It must be no more than 255 characters in length. Header name must be unique.",
													MarkdownDescription: "name specifies the name of a header on which to perform an action. Its value must be a valid HTTP header name as defined in RFC 2616 section 4.2. The name must consist only of alphanumeric and the following special characters, '-!#$%&'*+.^_''. The following header names are reserved and may not be modified via this API: Strict-Transport-Security, Proxy, Host, Cookie, Set-Cookie. It must be no more than 255 characters in length. Header name must be unique.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
														stringvalidator.LengthAtMost(255),
														stringvalidator.RegexMatches(regexp.MustCompile(`^[-!#$%&'*+.0-9A-Z^_`+"`"+`a-z|~]+$`), ""),
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

							"forwarded_header_policy": schema.StringAttribute{
								Description:         "forwardedHeaderPolicy specifies when and how the IngressController sets the Forwarded, X-Forwarded-For, X-Forwarded-Host, X-Forwarded-Port, X-Forwarded-Proto, and X-Forwarded-Proto-Version HTTP headers.  The value may be one of the following:  * 'Append', which specifies that the IngressController appends the headers, preserving existing headers.  * 'Replace', which specifies that the IngressController sets the headers, replacing any existing Forwarded or X-Forwarded-* headers.  * 'IfNone', which specifies that the IngressController sets the headers if they are not already set.  * 'Never', which specifies that the IngressController never sets the headers, preserving any existing headers.  By default, the policy is 'Append'.",
								MarkdownDescription: "forwardedHeaderPolicy specifies when and how the IngressController sets the Forwarded, X-Forwarded-For, X-Forwarded-Host, X-Forwarded-Port, X-Forwarded-Proto, and X-Forwarded-Proto-Version HTTP headers.  The value may be one of the following:  * 'Append', which specifies that the IngressController appends the headers, preserving existing headers.  * 'Replace', which specifies that the IngressController sets the headers, replacing any existing Forwarded or X-Forwarded-* headers.  * 'IfNone', which specifies that the IngressController sets the headers if they are not already set.  * 'Never', which specifies that the IngressController never sets the headers, preserving any existing headers.  By default, the policy is 'Append'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Append", "Replace", "IfNone", "Never"),
								},
							},

							"header_name_case_adjustments": schema.ListAttribute{
								Description:         "headerNameCaseAdjustments specifies case adjustments that can be applied to HTTP header names.  Each adjustment is specified as an HTTP header name with the desired capitalization.  For example, specifying 'X-Forwarded-For' indicates that the 'x-forwarded-for' HTTP header should be adjusted to have the specified capitalization.  These adjustments are only applied to cleartext, edge-terminated, and re-encrypt routes, and only when using HTTP/1.  For request headers, these adjustments are applied only for routes that have the haproxy.router.openshift.io/h1-adjust-case=true annotation.  For response headers, these adjustments are applied to all HTTP responses.  If this field is empty, no request headers are adjusted.",
								MarkdownDescription: "headerNameCaseAdjustments specifies case adjustments that can be applied to HTTP header names.  Each adjustment is specified as an HTTP header name with the desired capitalization.  For example, specifying 'X-Forwarded-For' indicates that the 'x-forwarded-for' HTTP header should be adjusted to have the specified capitalization.  These adjustments are only applied to cleartext, edge-terminated, and re-encrypt routes, and only when using HTTP/1.  For request headers, these adjustments are applied only for routes that have the haproxy.router.openshift.io/h1-adjust-case=true annotation.  For response headers, these adjustments are applied to all HTTP responses.  If this field is empty, no request headers are adjusted.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"unique_id": schema.SingleNestedAttribute{
								Description:         "uniqueId describes configuration for a custom HTTP header that the ingress controller should inject into incoming HTTP requests. Typically, this header is configured to have a value that is unique to the HTTP request.  The header can be used by applications or included in access logs to facilitate tracing individual HTTP requests.  If this field is empty, no such header is injected into requests.",
								MarkdownDescription: "uniqueId describes configuration for a custom HTTP header that the ingress controller should inject into incoming HTTP requests. Typically, this header is configured to have a value that is unique to the HTTP request.  The header can be used by applications or included in access logs to facilitate tracing individual HTTP requests.  If this field is empty, no such header is injected into requests.",
								Attributes: map[string]schema.Attribute{
									"format": schema.StringAttribute{
										Description:         "format specifies the format for the injected HTTP header's value. This field has no effect unless name is specified.  For the HAProxy-based ingress controller implementation, this format uses the same syntax as the HTTP log format.  If the field is empty, the default value is '%{+X}o %ci:%cp_%fi:%fp_%Ts_%rt:%pid'; see the corresponding HAProxy documentation: http://cbonte.github.io/haproxy-dconv/2.0/configuration.html#8.2.3",
										MarkdownDescription: "format specifies the format for the injected HTTP header's value. This field has no effect unless name is specified.  For the HAProxy-based ingress controller implementation, this format uses the same syntax as the HTTP log format.  If the field is empty, the default value is '%{+X}o %ci:%cp_%fi:%fp_%Ts_%rt:%pid'; see the corresponding HAProxy documentation: http://cbonte.github.io/haproxy-dconv/2.0/configuration.html#8.2.3",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(0),
											stringvalidator.LengthAtMost(1024),
											stringvalidator.RegexMatches(regexp.MustCompile(`^(%(%|(\{[-+]?[QXE](,[-+]?[QXE])*\})?([A-Za-z]+|\[[.0-9A-Z_a-z]+(\([^)]+\))?(,[.0-9A-Z_a-z]+(\([^)]+\))?)*\]))|[^%[:cntrl:]])*$`), ""),
										},
									},

									"name": schema.StringAttribute{
										Description:         "name specifies the name of the HTTP header (for example, 'unique-id') that the ingress controller should inject into HTTP requests.  The field's value must be a valid HTTP header name as defined in RFC 2616 section 4.2.  If the field is empty, no header is injected.",
										MarkdownDescription: "name specifies the name of the HTTP header (for example, 'unique-id') that the ingress controller should inject into HTTP requests.  The field's value must be a valid HTTP header name as defined in RFC 2616 section 4.2.  If the field is empty, no header is injected.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(0),
											stringvalidator.LengthAtMost(1024),
											stringvalidator.RegexMatches(regexp.MustCompile(`^$|^[-!#$%&'*+.0-9A-Z^_`+"`"+`a-z|~]+$`), ""),
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

					"logging": schema.SingleNestedAttribute{
						Description:         "logging defines parameters for what should be logged where.  If this field is empty, operational logs are enabled but access logs are disabled.",
						MarkdownDescription: "logging defines parameters for what should be logged where.  If this field is empty, operational logs are enabled but access logs are disabled.",
						Attributes: map[string]schema.Attribute{
							"access": schema.SingleNestedAttribute{
								Description:         "access describes how the client requests should be logged.  If this field is empty, access logging is disabled.",
								MarkdownDescription: "access describes how the client requests should be logged.  If this field is empty, access logging is disabled.",
								Attributes: map[string]schema.Attribute{
									"destination": schema.SingleNestedAttribute{
										Description:         "destination is where access logs go.",
										MarkdownDescription: "destination is where access logs go.",
										Attributes: map[string]schema.Attribute{
											"container": schema.SingleNestedAttribute{
												Description:         "container holds parameters for the Container logging destination. Present only if type is Container.",
												MarkdownDescription: "container holds parameters for the Container logging destination. Present only if type is Container.",
												Attributes: map[string]schema.Attribute{
													"max_length": schema.Int64Attribute{
														Description:         "maxLength is the maximum length of the log message.  Valid values are integers in the range 480 to 8192, inclusive.  When omitted, the default value is 1024.",
														MarkdownDescription: "maxLength is the maximum length of the log message.  Valid values are integers in the range 480 to 8192, inclusive.  When omitted, the default value is 1024.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(480),
															int64validator.AtMost(8192),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"syslog": schema.SingleNestedAttribute{
												Description:         "syslog holds parameters for a syslog endpoint.  Present only if type is Syslog.",
												MarkdownDescription: "syslog holds parameters for a syslog endpoint.  Present only if type is Syslog.",
												Attributes: map[string]schema.Attribute{
													"address": schema.StringAttribute{
														Description:         "address is the IP address of the syslog endpoint that receives log messages.",
														MarkdownDescription: "address is the IP address of the syslog endpoint that receives log messages.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"facility": schema.StringAttribute{
														Description:         "facility specifies the syslog facility of log messages.  If this field is empty, the facility is 'local1'.",
														MarkdownDescription: "facility specifies the syslog facility of log messages.  If this field is empty, the facility is 'local1'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("kern", "user", "mail", "daemon", "auth", "syslog", "lpr", "news", "uucp", "cron", "auth2", "ftp", "ntp", "audit", "alert", "cron2", "local0", "local1", "local2", "local3", "local4", "local5", "local6", "local7"),
														},
													},

													"max_length": schema.Int64Attribute{
														Description:         "maxLength is the maximum length of the log message.  Valid values are integers in the range 480 to 4096, inclusive.  When omitted, the default value is 1024.",
														MarkdownDescription: "maxLength is the maximum length of the log message.  Valid values are integers in the range 480 to 4096, inclusive.  When omitted, the default value is 1024.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(480),
															int64validator.AtMost(4096),
														},
													},

													"port": schema.Int64Attribute{
														Description:         "port is the UDP port number of the syslog endpoint that receives log messages.",
														MarkdownDescription: "port is the UDP port number of the syslog endpoint that receives log messages.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(1),
															int64validator.AtMost(65535),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"type": schema.StringAttribute{
												Description:         "type is the type of destination for logs.  It must be one of the following:  * Container  The ingress operator configures the sidecar container named 'logs' on the ingress controller pod and configures the ingress controller to write logs to the sidecar.  The logs are then available as container logs.  The expectation is that the administrator configures a custom logging solution that reads logs from this sidecar.  Note that using container logs means that logs may be dropped if the rate of logs exceeds the container runtime's or the custom logging solution's capacity.  * Syslog  Logs are sent to a syslog endpoint.  The administrator must specify an endpoint that can receive syslog messages.  The expectation is that the administrator has configured a custom syslog instance.",
												MarkdownDescription: "type is the type of destination for logs.  It must be one of the following:  * Container  The ingress operator configures the sidecar container named 'logs' on the ingress controller pod and configures the ingress controller to write logs to the sidecar.  The logs are then available as container logs.  The expectation is that the administrator configures a custom logging solution that reads logs from this sidecar.  Note that using container logs means that logs may be dropped if the rate of logs exceeds the container runtime's or the custom logging solution's capacity.  * Syslog  Logs are sent to a syslog endpoint.  The administrator must specify an endpoint that can receive syslog messages.  The expectation is that the administrator has configured a custom syslog instance.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("Container", "Syslog"),
												},
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"http_capture_cookies": schema.ListNestedAttribute{
										Description:         "httpCaptureCookies specifies HTTP cookies that should be captured in access logs.  If this field is empty, no cookies are captured.",
										MarkdownDescription: "httpCaptureCookies specifies HTTP cookies that should be captured in access logs.  If this field is empty, no cookies are captured.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"match_type": schema.StringAttribute{
													Description:         "matchType specifies the type of match to be performed on the cookie name.  Allowed values are 'Exact' for an exact string match and 'Prefix' for a string prefix match.  If 'Exact' is specified, a name must be specified in the name field.  If 'Prefix' is provided, a prefix must be specified in the namePrefix field.  For example, specifying matchType 'Prefix' and namePrefix 'foo' will capture a cookie named 'foo' or 'foobar' but not one named 'bar'.  The first matching cookie is captured.",
													MarkdownDescription: "matchType specifies the type of match to be performed on the cookie name.  Allowed values are 'Exact' for an exact string match and 'Prefix' for a string prefix match.  If 'Exact' is specified, a name must be specified in the name field.  If 'Prefix' is provided, a prefix must be specified in the namePrefix field.  For example, specifying matchType 'Prefix' and namePrefix 'foo' will capture a cookie named 'foo' or 'foobar' but not one named 'bar'.  The first matching cookie is captured.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("Exact", "Prefix"),
													},
												},

												"max_length": schema.Int64Attribute{
													Description:         "maxLength specifies a maximum length of the string that will be logged, which includes the cookie name, cookie value, and one-character delimiter.  If the log entry exceeds this length, the value will be truncated in the log message.  Note that the ingress controller may impose a separate bound on the total length of HTTP headers in a request.",
													MarkdownDescription: "maxLength specifies a maximum length of the string that will be logged, which includes the cookie name, cookie value, and one-character delimiter.  If the log entry exceeds this length, the value will be truncated in the log message.  Note that the ingress controller may impose a separate bound on the total length of HTTP headers in a request.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(1),
														int64validator.AtMost(1024),
													},
												},

												"name": schema.StringAttribute{
													Description:         "name specifies a cookie name.  Its value must be a valid HTTP cookie name as defined in RFC 6265 section 4.1.",
													MarkdownDescription: "name specifies a cookie name.  Its value must be a valid HTTP cookie name as defined in RFC 6265 section 4.1.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(0),
														stringvalidator.LengthAtMost(1024),
														stringvalidator.RegexMatches(regexp.MustCompile(`^[-!#$%&'*+.0-9A-Z^_`+"`"+`a-z|~]*$`), ""),
													},
												},

												"name_prefix": schema.StringAttribute{
													Description:         "namePrefix specifies a cookie name prefix.  Its value must be a valid HTTP cookie name as defined in RFC 6265 section 4.1.",
													MarkdownDescription: "namePrefix specifies a cookie name prefix.  Its value must be a valid HTTP cookie name as defined in RFC 6265 section 4.1.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(0),
														stringvalidator.LengthAtMost(1024),
														stringvalidator.RegexMatches(regexp.MustCompile(`^[-!#$%&'*+.0-9A-Z^_`+"`"+`a-z|~]*$`), ""),
													},
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"http_capture_headers": schema.SingleNestedAttribute{
										Description:         "httpCaptureHeaders defines HTTP headers that should be captured in access logs.  If this field is empty, no headers are captured.  Note that this option only applies to cleartext HTTP connections and to secure HTTP connections for which the ingress controller terminates encryption (that is, edge-terminated or reencrypt connections).  Headers cannot be captured for TLS passthrough connections.",
										MarkdownDescription: "httpCaptureHeaders defines HTTP headers that should be captured in access logs.  If this field is empty, no headers are captured.  Note that this option only applies to cleartext HTTP connections and to secure HTTP connections for which the ingress controller terminates encryption (that is, edge-terminated or reencrypt connections).  Headers cannot be captured for TLS passthrough connections.",
										Attributes: map[string]schema.Attribute{
											"request": schema.ListNestedAttribute{
												Description:         "request specifies which HTTP request headers to capture.  If this field is empty, no request headers are captured.",
												MarkdownDescription: "request specifies which HTTP request headers to capture.  If this field is empty, no request headers are captured.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"max_length": schema.Int64Attribute{
															Description:         "maxLength specifies a maximum length for the header value.  If a header value exceeds this length, the value will be truncated in the log message.  Note that the ingress controller may impose a separate bound on the total length of HTTP headers in a request.",
															MarkdownDescription: "maxLength specifies a maximum length for the header value.  If a header value exceeds this length, the value will be truncated in the log message.  Note that the ingress controller may impose a separate bound on the total length of HTTP headers in a request.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(1),
															},
														},

														"name": schema.StringAttribute{
															Description:         "name specifies a header name.  Its value must be a valid HTTP header name as defined in RFC 2616 section 4.2.",
															MarkdownDescription: "name specifies a header name.  Its value must be a valid HTTP header name as defined in RFC 2616 section 4.2.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.RegexMatches(regexp.MustCompile(`^[-!#$%&'*+.0-9A-Z^_`+"`"+`a-z|~]+$`), ""),
															},
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"response": schema.ListNestedAttribute{
												Description:         "response specifies which HTTP response headers to capture.  If this field is empty, no response headers are captured.",
												MarkdownDescription: "response specifies which HTTP response headers to capture.  If this field is empty, no response headers are captured.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"max_length": schema.Int64Attribute{
															Description:         "maxLength specifies a maximum length for the header value.  If a header value exceeds this length, the value will be truncated in the log message.  Note that the ingress controller may impose a separate bound on the total length of HTTP headers in a request.",
															MarkdownDescription: "maxLength specifies a maximum length for the header value.  If a header value exceeds this length, the value will be truncated in the log message.  Note that the ingress controller may impose a separate bound on the total length of HTTP headers in a request.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(1),
															},
														},

														"name": schema.StringAttribute{
															Description:         "name specifies a header name.  Its value must be a valid HTTP header name as defined in RFC 2616 section 4.2.",
															MarkdownDescription: "name specifies a header name.  Its value must be a valid HTTP header name as defined in RFC 2616 section 4.2.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.RegexMatches(regexp.MustCompile(`^[-!#$%&'*+.0-9A-Z^_`+"`"+`a-z|~]+$`), ""),
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

									"http_log_format": schema.StringAttribute{
										Description:         "httpLogFormat specifies the format of the log message for an HTTP request.  If this field is empty, log messages use the implementation's default HTTP log format.  For HAProxy's default HTTP log format, see the HAProxy documentation: http://cbonte.github.io/haproxy-dconv/2.0/configuration.html#8.2.3  Note that this format only applies to cleartext HTTP connections and to secure HTTP connections for which the ingress controller terminates encryption (that is, edge-terminated or reencrypt connections).  It does not affect the log format for TLS passthrough connections.",
										MarkdownDescription: "httpLogFormat specifies the format of the log message for an HTTP request.  If this field is empty, log messages use the implementation's default HTTP log format.  For HAProxy's default HTTP log format, see the HAProxy documentation: http://cbonte.github.io/haproxy-dconv/2.0/configuration.html#8.2.3  Note that this format only applies to cleartext HTTP connections and to secure HTTP connections for which the ingress controller terminates encryption (that is, edge-terminated or reencrypt connections).  It does not affect the log format for TLS passthrough connections.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"log_empty_requests": schema.StringAttribute{
										Description:         "logEmptyRequests specifies how connections on which no request is received should be logged.  Typically, these empty requests come from load balancers' health probes or Web browsers' speculative connections ('preconnect'), in which case logging these requests may be undesirable.  However, these requests may also be caused by network errors, in which case logging empty requests may be useful for diagnosing the errors.  In addition, these requests may be caused by port scans, in which case logging empty requests may aid in detecting intrusion attempts.  Allowed values for this field are 'Log' and 'Ignore'.  The default value is 'Log'.",
										MarkdownDescription: "logEmptyRequests specifies how connections on which no request is received should be logged.  Typically, these empty requests come from load balancers' health probes or Web browsers' speculative connections ('preconnect'), in which case logging these requests may be undesirable.  However, these requests may also be caused by network errors, in which case logging empty requests may be useful for diagnosing the errors.  In addition, these requests may be caused by port scans, in which case logging empty requests may aid in detecting intrusion attempts.  Allowed values for this field are 'Log' and 'Ignore'.  The default value is 'Log'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("Log", "Ignore"),
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

					"namespace_selector": schema.SingleNestedAttribute{
						Description:         "namespaceSelector is used to filter the set of namespaces serviced by the ingress controller. This is useful for implementing shards.  If unset, the default is no filtering.",
						MarkdownDescription: "namespaceSelector is used to filter the set of namespaces serviced by the ingress controller. This is useful for implementing shards.  If unset, the default is no filtering.",
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
											Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"values": schema.ListAttribute{
											Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

					"node_placement": schema.SingleNestedAttribute{
						Description:         "nodePlacement enables explicit control over the scheduling of the ingress controller.  If unset, defaults are used. See NodePlacement for more details.",
						MarkdownDescription: "nodePlacement enables explicit control over the scheduling of the ingress controller.  If unset, defaults are used. See NodePlacement for more details.",
						Attributes: map[string]schema.Attribute{
							"node_selector": schema.SingleNestedAttribute{
								Description:         "nodeSelector is the node selector applied to ingress controller deployments.  If set, the specified selector is used and replaces the default.  If unset, the default depends on the value of the defaultPlacement field in the cluster config.openshift.io/v1/ingresses status.  When defaultPlacement is Workers, the default is:  kubernetes.io/os: linux node-role.kubernetes.io/worker: ''  When defaultPlacement is ControlPlane, the default is:  kubernetes.io/os: linux node-role.kubernetes.io/master: ''  These defaults are subject to change.  Note that using nodeSelector.matchExpressions is not supported.  Only nodeSelector.matchLabels may be used.  This is a limitation of the Kubernetes API: the pod spec does not allow complex expressions for node selectors.",
								MarkdownDescription: "nodeSelector is the node selector applied to ingress controller deployments.  If set, the specified selector is used and replaces the default.  If unset, the default depends on the value of the defaultPlacement field in the cluster config.openshift.io/v1/ingresses status.  When defaultPlacement is Workers, the default is:  kubernetes.io/os: linux node-role.kubernetes.io/worker: ''  When defaultPlacement is ControlPlane, the default is:  kubernetes.io/os: linux node-role.kubernetes.io/master: ''  These defaults are subject to change.  Note that using nodeSelector.matchExpressions is not supported.  Only nodeSelector.matchLabels may be used.  This is a limitation of the Kubernetes API: the pod spec does not allow complex expressions for node selectors.",
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
													Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
													MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"values": schema.ListAttribute{
													Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
													MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

							"tolerations": schema.ListNestedAttribute{
								Description:         "tolerations is a list of tolerations applied to ingress controller deployments.  The default is an empty list.  See https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/",
								MarkdownDescription: "tolerations is a list of tolerations applied to ingress controller deployments.  The default is an empty list.  See https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"effect": schema.StringAttribute{
											Description:         "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
											MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"key": schema.StringAttribute{
											Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
											MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"operator": schema.StringAttribute{
											Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
											MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"toleration_seconds": schema.Int64Attribute{
											Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
											MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
											MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
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

					"replicas": schema.Int64Attribute{
						Description:         "replicas is the desired number of ingress controller replicas. If unset, the default depends on the value of the defaultPlacement field in the cluster config.openshift.io/v1/ingresses status.  The value of replicas is set based on the value of a chosen field in the Infrastructure CR. If defaultPlacement is set to ControlPlane, the chosen field will be controlPlaneTopology. If it is set to Workers the chosen field will be infrastructureTopology. Replicas will then be set to 1 or 2 based whether the chosen field's value is SingleReplica or HighlyAvailable, respectively.  These defaults are subject to change.",
						MarkdownDescription: "replicas is the desired number of ingress controller replicas. If unset, the default depends on the value of the defaultPlacement field in the cluster config.openshift.io/v1/ingresses status.  The value of replicas is set based on the value of a chosen field in the Infrastructure CR. If defaultPlacement is set to ControlPlane, the chosen field will be controlPlaneTopology. If it is set to Workers the chosen field will be infrastructureTopology. Replicas will then be set to 1 or 2 based whether the chosen field's value is SingleReplica or HighlyAvailable, respectively.  These defaults are subject to change.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"route_admission": schema.SingleNestedAttribute{
						Description:         "routeAdmission defines a policy for handling new route claims (for example, to allow or deny claims across namespaces).  If empty, defaults will be applied. See specific routeAdmission fields for details about their defaults.",
						MarkdownDescription: "routeAdmission defines a policy for handling new route claims (for example, to allow or deny claims across namespaces).  If empty, defaults will be applied. See specific routeAdmission fields for details about their defaults.",
						Attributes: map[string]schema.Attribute{
							"namespace_ownership": schema.StringAttribute{
								Description:         "namespaceOwnership describes how host name claims across namespaces should be handled.  Value must be one of:  - Strict: Do not allow routes in different namespaces to claim the same host.  - InterNamespaceAllowed: Allow routes to claim different paths of the same host name across namespaces.  If empty, the default is Strict.",
								MarkdownDescription: "namespaceOwnership describes how host name claims across namespaces should be handled.  Value must be one of:  - Strict: Do not allow routes in different namespaces to claim the same host.  - InterNamespaceAllowed: Allow routes to claim different paths of the same host name across namespaces.  If empty, the default is Strict.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("InterNamespaceAllowed", "Strict"),
								},
							},

							"wildcard_policy": schema.StringAttribute{
								Description:         "wildcardPolicy describes how routes with wildcard policies should be handled for the ingress controller. WildcardPolicy controls use of routes [1] exposed by the ingress controller based on the route's wildcard policy.  [1] https://github.com/openshift/api/blob/master/route/v1/types.go  Note: Updating WildcardPolicy from WildcardsAllowed to WildcardsDisallowed will cause admitted routes with a wildcard policy of Subdomain to stop working. These routes must be updated to a wildcard policy of None to be readmitted by the ingress controller.  WildcardPolicy supports WildcardsAllowed and WildcardsDisallowed values.  If empty, defaults to 'WildcardsDisallowed'.",
								MarkdownDescription: "wildcardPolicy describes how routes with wildcard policies should be handled for the ingress controller. WildcardPolicy controls use of routes [1] exposed by the ingress controller based on the route's wildcard policy.  [1] https://github.com/openshift/api/blob/master/route/v1/types.go  Note: Updating WildcardPolicy from WildcardsAllowed to WildcardsDisallowed will cause admitted routes with a wildcard policy of Subdomain to stop working. These routes must be updated to a wildcard policy of None to be readmitted by the ingress controller.  WildcardPolicy supports WildcardsAllowed and WildcardsDisallowed values.  If empty, defaults to 'WildcardsDisallowed'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("WildcardsAllowed", "WildcardsDisallowed"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"route_selector": schema.SingleNestedAttribute{
						Description:         "routeSelector is used to filter the set of Routes serviced by the ingress controller. This is useful for implementing shards.  If unset, the default is no filtering.",
						MarkdownDescription: "routeSelector is used to filter the set of Routes serviced by the ingress controller. This is useful for implementing shards.  If unset, the default is no filtering.",
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
											Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"values": schema.ListAttribute{
											Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

					"tls_security_profile": schema.SingleNestedAttribute{
						Description:         "tlsSecurityProfile specifies settings for TLS connections for ingresscontrollers.  If unset, the default is based on the apiservers.config.openshift.io/cluster resource.  Note that when using the Old, Intermediate, and Modern profile types, the effective profile configuration is subject to change between releases. For example, given a specification to use the Intermediate profile deployed on release X.Y.Z, an upgrade to release X.Y.Z+1 may cause a new profile configuration to be applied to the ingress controller, resulting in a rollout.",
						MarkdownDescription: "tlsSecurityProfile specifies settings for TLS connections for ingresscontrollers.  If unset, the default is based on the apiservers.config.openshift.io/cluster resource.  Note that when using the Old, Intermediate, and Modern profile types, the effective profile configuration is subject to change between releases. For example, given a specification to use the Intermediate profile deployed on release X.Y.Z, an upgrade to release X.Y.Z+1 may cause a new profile configuration to be applied to the ingress controller, resulting in a rollout.",
						Attributes: map[string]schema.Attribute{
							"custom": schema.SingleNestedAttribute{
								Description:         "custom is a user-defined TLS security profile. Be extremely careful using a custom profile as invalid configurations can be catastrophic. An example custom profile looks like this:  ciphers:  - ECDHE-ECDSA-CHACHA20-POLY1305  - ECDHE-RSA-CHACHA20-POLY1305  - ECDHE-RSA-AES128-GCM-SHA256  - ECDHE-ECDSA-AES128-GCM-SHA256  minTLSVersion: VersionTLS11",
								MarkdownDescription: "custom is a user-defined TLS security profile. Be extremely careful using a custom profile as invalid configurations can be catastrophic. An example custom profile looks like this:  ciphers:  - ECDHE-ECDSA-CHACHA20-POLY1305  - ECDHE-RSA-CHACHA20-POLY1305  - ECDHE-RSA-AES128-GCM-SHA256  - ECDHE-ECDSA-AES128-GCM-SHA256  minTLSVersion: VersionTLS11",
								Attributes: map[string]schema.Attribute{
									"ciphers": schema.ListAttribute{
										Description:         "ciphers is used to specify the cipher algorithms that are negotiated during the TLS handshake.  Operators may remove entries their operands do not support.  For example, to use DES-CBC3-SHA  (yaml):  ciphers: - DES-CBC3-SHA",
										MarkdownDescription: "ciphers is used to specify the cipher algorithms that are negotiated during the TLS handshake.  Operators may remove entries their operands do not support.  For example, to use DES-CBC3-SHA  (yaml):  ciphers: - DES-CBC3-SHA",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"min_tls_version": schema.StringAttribute{
										Description:         "minTLSVersion is used to specify the minimal version of the TLS protocol that is negotiated during the TLS handshake. For example, to use TLS versions 1.1, 1.2 and 1.3 (yaml):  minTLSVersion: VersionTLS11  NOTE: currently the highest minTLSVersion allowed is VersionTLS12",
										MarkdownDescription: "minTLSVersion is used to specify the minimal version of the TLS protocol that is negotiated during the TLS handshake. For example, to use TLS versions 1.1, 1.2 and 1.3 (yaml):  minTLSVersion: VersionTLS11  NOTE: currently the highest minTLSVersion allowed is VersionTLS12",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("VersionTLS10", "VersionTLS11", "VersionTLS12", "VersionTLS13"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"intermediate": schema.MapAttribute{
								Description:         "intermediate is a TLS security profile based on:  https://wiki.mozilla.org/Security/Server_Side_TLS#Intermediate_compatibility_.28recommended.29  and looks like this (yaml):  ciphers:  - TLS_AES_128_GCM_SHA256  - TLS_AES_256_GCM_SHA384  - TLS_CHACHA20_POLY1305_SHA256  - ECDHE-ECDSA-AES128-GCM-SHA256  - ECDHE-RSA-AES128-GCM-SHA256  - ECDHE-ECDSA-AES256-GCM-SHA384  - ECDHE-RSA-AES256-GCM-SHA384  - ECDHE-ECDSA-CHACHA20-POLY1305  - ECDHE-RSA-CHACHA20-POLY1305  - DHE-RSA-AES128-GCM-SHA256  - DHE-RSA-AES256-GCM-SHA384  minTLSVersion: VersionTLS12",
								MarkdownDescription: "intermediate is a TLS security profile based on:  https://wiki.mozilla.org/Security/Server_Side_TLS#Intermediate_compatibility_.28recommended.29  and looks like this (yaml):  ciphers:  - TLS_AES_128_GCM_SHA256  - TLS_AES_256_GCM_SHA384  - TLS_CHACHA20_POLY1305_SHA256  - ECDHE-ECDSA-AES128-GCM-SHA256  - ECDHE-RSA-AES128-GCM-SHA256  - ECDHE-ECDSA-AES256-GCM-SHA384  - ECDHE-RSA-AES256-GCM-SHA384  - ECDHE-ECDSA-CHACHA20-POLY1305  - ECDHE-RSA-CHACHA20-POLY1305  - DHE-RSA-AES128-GCM-SHA256  - DHE-RSA-AES256-GCM-SHA384  minTLSVersion: VersionTLS12",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"modern": schema.MapAttribute{
								Description:         "modern is a TLS security profile based on:  https://wiki.mozilla.org/Security/Server_Side_TLS#Modern_compatibility  and looks like this (yaml):  ciphers:  - TLS_AES_128_GCM_SHA256  - TLS_AES_256_GCM_SHA384  - TLS_CHACHA20_POLY1305_SHA256  minTLSVersion: VersionTLS13",
								MarkdownDescription: "modern is a TLS security profile based on:  https://wiki.mozilla.org/Security/Server_Side_TLS#Modern_compatibility  and looks like this (yaml):  ciphers:  - TLS_AES_128_GCM_SHA256  - TLS_AES_256_GCM_SHA384  - TLS_CHACHA20_POLY1305_SHA256  minTLSVersion: VersionTLS13",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"old": schema.MapAttribute{
								Description:         "old is a TLS security profile based on:  https://wiki.mozilla.org/Security/Server_Side_TLS#Old_backward_compatibility  and looks like this (yaml):  ciphers:  - TLS_AES_128_GCM_SHA256  - TLS_AES_256_GCM_SHA384  - TLS_CHACHA20_POLY1305_SHA256  - ECDHE-ECDSA-AES128-GCM-SHA256  - ECDHE-RSA-AES128-GCM-SHA256  - ECDHE-ECDSA-AES256-GCM-SHA384  - ECDHE-RSA-AES256-GCM-SHA384  - ECDHE-ECDSA-CHACHA20-POLY1305  - ECDHE-RSA-CHACHA20-POLY1305  - DHE-RSA-AES128-GCM-SHA256  - DHE-RSA-AES256-GCM-SHA384  - DHE-RSA-CHACHA20-POLY1305  - ECDHE-ECDSA-AES128-SHA256  - ECDHE-RSA-AES128-SHA256  - ECDHE-ECDSA-AES128-SHA  - ECDHE-RSA-AES128-SHA  - ECDHE-ECDSA-AES256-SHA384  - ECDHE-RSA-AES256-SHA384  - ECDHE-ECDSA-AES256-SHA  - ECDHE-RSA-AES256-SHA  - DHE-RSA-AES128-SHA256  - DHE-RSA-AES256-SHA256  - AES128-GCM-SHA256  - AES256-GCM-SHA384  - AES128-SHA256  - AES256-SHA256  - AES128-SHA  - AES256-SHA  - DES-CBC3-SHA  minTLSVersion: VersionTLS10",
								MarkdownDescription: "old is a TLS security profile based on:  https://wiki.mozilla.org/Security/Server_Side_TLS#Old_backward_compatibility  and looks like this (yaml):  ciphers:  - TLS_AES_128_GCM_SHA256  - TLS_AES_256_GCM_SHA384  - TLS_CHACHA20_POLY1305_SHA256  - ECDHE-ECDSA-AES128-GCM-SHA256  - ECDHE-RSA-AES128-GCM-SHA256  - ECDHE-ECDSA-AES256-GCM-SHA384  - ECDHE-RSA-AES256-GCM-SHA384  - ECDHE-ECDSA-CHACHA20-POLY1305  - ECDHE-RSA-CHACHA20-POLY1305  - DHE-RSA-AES128-GCM-SHA256  - DHE-RSA-AES256-GCM-SHA384  - DHE-RSA-CHACHA20-POLY1305  - ECDHE-ECDSA-AES128-SHA256  - ECDHE-RSA-AES128-SHA256  - ECDHE-ECDSA-AES128-SHA  - ECDHE-RSA-AES128-SHA  - ECDHE-ECDSA-AES256-SHA384  - ECDHE-RSA-AES256-SHA384  - ECDHE-ECDSA-AES256-SHA  - ECDHE-RSA-AES256-SHA  - DHE-RSA-AES128-SHA256  - DHE-RSA-AES256-SHA256  - AES128-GCM-SHA256  - AES256-GCM-SHA384  - AES128-SHA256  - AES256-SHA256  - AES128-SHA  - AES256-SHA  - DES-CBC3-SHA  minTLSVersion: VersionTLS10",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "type is one of Old, Intermediate, Modern or Custom. Custom provides the ability to specify individual TLS security profile parameters. Old, Intermediate and Modern are TLS security profiles based on:  https://wiki.mozilla.org/Security/Server_Side_TLS#Recommended_configurations  The profiles are intent based, so they may change over time as new ciphers are developed and existing ciphers are found to be insecure.  Depending on precisely which ciphers are available to a process, the list may be reduced.  Note that the Modern profile is currently not supported because it is not yet well adopted by common software libraries.",
								MarkdownDescription: "type is one of Old, Intermediate, Modern or Custom. Custom provides the ability to specify individual TLS security profile parameters. Old, Intermediate and Modern are TLS security profiles based on:  https://wiki.mozilla.org/Security/Server_Side_TLS#Recommended_configurations  The profiles are intent based, so they may change over time as new ciphers are developed and existing ciphers are found to be insecure.  Depending on precisely which ciphers are available to a process, the list may be reduced.  Note that the Modern profile is currently not supported because it is not yet well adopted by common software libraries.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Old", "Intermediate", "Modern", "Custom"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"tuning_options": schema.SingleNestedAttribute{
						Description:         "tuningOptions defines parameters for adjusting the performance of ingress controller pods. All fields are optional and will use their respective defaults if not set. See specific tuningOptions fields for more details.  Setting fields within tuningOptions is generally not recommended. The default values are suitable for most configurations.",
						MarkdownDescription: "tuningOptions defines parameters for adjusting the performance of ingress controller pods. All fields are optional and will use their respective defaults if not set. See specific tuningOptions fields for more details.  Setting fields within tuningOptions is generally not recommended. The default values are suitable for most configurations.",
						Attributes: map[string]schema.Attribute{
							"client_fin_timeout": schema.StringAttribute{
								Description:         "clientFinTimeout defines how long a connection will be held open while waiting for the client response to the server/backend closing the connection.  If unset, the default timeout is 1s",
								MarkdownDescription: "clientFinTimeout defines how long a connection will be held open while waiting for the client response to the server/backend closing the connection.  If unset, the default timeout is 1s",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"client_timeout": schema.StringAttribute{
								Description:         "clientTimeout defines how long a connection will be held open while waiting for a client response.  If unset, the default timeout is 30s",
								MarkdownDescription: "clientTimeout defines how long a connection will be held open while waiting for a client response.  If unset, the default timeout is 30s",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"header_buffer_bytes": schema.Int64Attribute{
								Description:         "headerBufferBytes describes how much memory should be reserved (in bytes) for IngressController connection sessions. Note that this value must be at least 16384 if HTTP/2 is enabled for the IngressController (https://tools.ietf.org/html/rfc7540). If this field is empty, the IngressController will use a default value of 32768 bytes.  Setting this field is generally not recommended as headerBufferBytes values that are too small may break the IngressController and headerBufferBytes values that are too large could cause the IngressController to use significantly more memory than necessary.",
								MarkdownDescription: "headerBufferBytes describes how much memory should be reserved (in bytes) for IngressController connection sessions. Note that this value must be at least 16384 if HTTP/2 is enabled for the IngressController (https://tools.ietf.org/html/rfc7540). If this field is empty, the IngressController will use a default value of 32768 bytes.  Setting this field is generally not recommended as headerBufferBytes values that are too small may break the IngressController and headerBufferBytes values that are too large could cause the IngressController to use significantly more memory than necessary.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(16384),
								},
							},

							"header_buffer_max_rewrite_bytes": schema.Int64Attribute{
								Description:         "headerBufferMaxRewriteBytes describes how much memory should be reserved (in bytes) from headerBufferBytes for HTTP header rewriting and appending for IngressController connection sessions. Note that incoming HTTP requests will be limited to (headerBufferBytes - headerBufferMaxRewriteBytes) bytes, meaning headerBufferBytes must be greater than headerBufferMaxRewriteBytes. If this field is empty, the IngressController will use a default value of 8192 bytes.  Setting this field is generally not recommended as headerBufferMaxRewriteBytes values that are too small may break the IngressController and headerBufferMaxRewriteBytes values that are too large could cause the IngressController to use significantly more memory than necessary.",
								MarkdownDescription: "headerBufferMaxRewriteBytes describes how much memory should be reserved (in bytes) from headerBufferBytes for HTTP header rewriting and appending for IngressController connection sessions. Note that incoming HTTP requests will be limited to (headerBufferBytes - headerBufferMaxRewriteBytes) bytes, meaning headerBufferBytes must be greater than headerBufferMaxRewriteBytes. If this field is empty, the IngressController will use a default value of 8192 bytes.  Setting this field is generally not recommended as headerBufferMaxRewriteBytes values that are too small may break the IngressController and headerBufferMaxRewriteBytes values that are too large could cause the IngressController to use significantly more memory than necessary.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(4096),
								},
							},

							"health_check_interval": schema.StringAttribute{
								Description:         "healthCheckInterval defines how long the router waits between two consecutive health checks on its configured backends.  This value is applied globally as a default for all routes, but may be overridden per-route by the route annotation 'router.openshift.io/haproxy.health.check.interval'.  Expects an unsigned duration string of decimal numbers, each with optional fraction and a unit suffix, eg '300ms', '1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs' U+00B5 or 'μs' U+03BC), 'ms', 's', 'm', 'h'.  Setting this to less than 5s can cause excess traffic due to too frequent TCP health checks and accompanying SYN packet storms.  Alternatively, setting this too high can result in increased latency, due to backend servers that are no longer available, but haven't yet been detected as such.  An empty or zero healthCheckInterval means no opinion and IngressController chooses a default, which is subject to change over time. Currently the default healthCheckInterval value is 5s.  Currently the minimum allowed value is 1s and the maximum allowed value is 2147483647ms (24.85 days).  Both are subject to change over time.",
								MarkdownDescription: "healthCheckInterval defines how long the router waits between two consecutive health checks on its configured backends.  This value is applied globally as a default for all routes, but may be overridden per-route by the route annotation 'router.openshift.io/haproxy.health.check.interval'.  Expects an unsigned duration string of decimal numbers, each with optional fraction and a unit suffix, eg '300ms', '1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs' U+00B5 or 'μs' U+03BC), 'ms', 's', 'm', 'h'.  Setting this to less than 5s can cause excess traffic due to too frequent TCP health checks and accompanying SYN packet storms.  Alternatively, setting this too high can result in increased latency, due to backend servers that are no longer available, but haven't yet been detected as such.  An empty or zero healthCheckInterval means no opinion and IngressController chooses a default, which is subject to change over time. Currently the default healthCheckInterval value is 5s.  Currently the minimum allowed value is 1s and the maximum allowed value is 2147483647ms (24.85 days).  Both are subject to change over time.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(0|([0-9]+(\.[0-9]+)?(ns|us|µs|μs|ms|s|m|h))+)$`), ""),
								},
							},

							"max_connections": schema.Int64Attribute{
								Description:         "maxConnections defines the maximum number of simultaneous connections that can be established per HAProxy process. Increasing this value allows each ingress controller pod to handle more connections but at the cost of additional system resources being consumed.  Permitted values are: empty, 0, -1, and the range 2000-2000000.  If this field is empty or 0, the IngressController will use the default value of 50000, but the default is subject to change in future releases.  If the value is -1 then HAProxy will dynamically compute a maximum value based on the available ulimits in the running container. Selecting -1 (i.e., auto) will result in a large value being computed (~520000 on OpenShift >=4.10 clusters) and therefore each HAProxy process will incur significant memory usage compared to the current default of 50000.  Setting a value that is greater than the current operating system limit will prevent the HAProxy process from starting.  If you choose a discrete value (e.g., 750000) and the router pod is migrated to a new node, there's no guarantee that that new node has identical ulimits configured. In such a scenario the pod would fail to start. If you have nodes with different ulimits configured (e.g., different tuned profiles) and you choose a discrete value then the guidance is to use -1 and let the value be computed dynamically at runtime.  You can monitor memory usage for router containers with the following metric: 'container_memory_working_set_bytes{container='router',namespace='openshift-ingress'}'.  You can monitor memory usage of individual HAProxy processes in router containers with the following metric: 'container_memory_working_set_bytes{container='router',namespace='openshift-ingress'}/container_processes{container='router',namespace='openshift-ingress'}'.",
								MarkdownDescription: "maxConnections defines the maximum number of simultaneous connections that can be established per HAProxy process. Increasing this value allows each ingress controller pod to handle more connections but at the cost of additional system resources being consumed.  Permitted values are: empty, 0, -1, and the range 2000-2000000.  If this field is empty or 0, the IngressController will use the default value of 50000, but the default is subject to change in future releases.  If the value is -1 then HAProxy will dynamically compute a maximum value based on the available ulimits in the running container. Selecting -1 (i.e., auto) will result in a large value being computed (~520000 on OpenShift >=4.10 clusters) and therefore each HAProxy process will incur significant memory usage compared to the current default of 50000.  Setting a value that is greater than the current operating system limit will prevent the HAProxy process from starting.  If you choose a discrete value (e.g., 750000) and the router pod is migrated to a new node, there's no guarantee that that new node has identical ulimits configured. In such a scenario the pod would fail to start. If you have nodes with different ulimits configured (e.g., different tuned profiles) and you choose a discrete value then the guidance is to use -1 and let the value be computed dynamically at runtime.  You can monitor memory usage for router containers with the following metric: 'container_memory_working_set_bytes{container='router',namespace='openshift-ingress'}'.  You can monitor memory usage of individual HAProxy processes in router containers with the following metric: 'container_memory_working_set_bytes{container='router',namespace='openshift-ingress'}/container_processes{container='router',namespace='openshift-ingress'}'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"reload_interval": schema.StringAttribute{
								Description:         "reloadInterval defines the minimum interval at which the router is allowed to reload to accept new changes. Increasing this value can prevent the accumulation of HAProxy processes, depending on the scenario. Increasing this interval can also lessen load imbalance on a backend's servers when using the roundrobin balancing algorithm. Alternatively, decreasing this value may decrease latency since updates to HAProxy's configuration can take effect more quickly.  The value must be a time duration value; see <https://pkg.go.dev/time#ParseDuration>. Currently, the minimum value allowed is 1s, and the maximum allowed value is 120s. Minimum and maximum allowed values may change in future versions of OpenShift. Note that if a duration outside of these bounds is provided, the value of reloadInterval will be capped/floored and not rejected (e.g. a duration of over 120s will be capped to 120s; the IngressController will not reject and replace this disallowed value with the default).  A zero value for reloadInterval tells the IngressController to choose the default, which is currently 5s and subject to change without notice.  This field expects an unsigned duration string of decimal numbers, each with optional fraction and a unit suffix, e.g. '300ms', '1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs' U+00B5 or 'μs' U+03BC), 'ms', 's', 'm', 'h'.  Note: Setting a value significantly larger than the default of 5s can cause latency in observing updates to routes and their endpoints. HAProxy's configuration will be reloaded less frequently, and newly created routes will not be served until the subsequent reload.",
								MarkdownDescription: "reloadInterval defines the minimum interval at which the router is allowed to reload to accept new changes. Increasing this value can prevent the accumulation of HAProxy processes, depending on the scenario. Increasing this interval can also lessen load imbalance on a backend's servers when using the roundrobin balancing algorithm. Alternatively, decreasing this value may decrease latency since updates to HAProxy's configuration can take effect more quickly.  The value must be a time duration value; see <https://pkg.go.dev/time#ParseDuration>. Currently, the minimum value allowed is 1s, and the maximum allowed value is 120s. Minimum and maximum allowed values may change in future versions of OpenShift. Note that if a duration outside of these bounds is provided, the value of reloadInterval will be capped/floored and not rejected (e.g. a duration of over 120s will be capped to 120s; the IngressController will not reject and replace this disallowed value with the default).  A zero value for reloadInterval tells the IngressController to choose the default, which is currently 5s and subject to change without notice.  This field expects an unsigned duration string of decimal numbers, each with optional fraction and a unit suffix, e.g. '300ms', '1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs' U+00B5 or 'μs' U+03BC), 'ms', 's', 'm', 'h'.  Note: Setting a value significantly larger than the default of 5s can cause latency in observing updates to routes and their endpoints. HAProxy's configuration will be reloaded less frequently, and newly created routes will not be served until the subsequent reload.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(0|([0-9]+(\.[0-9]+)?(ns|us|µs|μs|ms|s|m|h))+)$`), ""),
								},
							},

							"server_fin_timeout": schema.StringAttribute{
								Description:         "serverFinTimeout defines how long a connection will be held open while waiting for the server/backend response to the client closing the connection.  If unset, the default timeout is 1s",
								MarkdownDescription: "serverFinTimeout defines how long a connection will be held open while waiting for the server/backend response to the client closing the connection.  If unset, the default timeout is 1s",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"server_timeout": schema.StringAttribute{
								Description:         "serverTimeout defines how long a connection will be held open while waiting for a server/backend response.  If unset, the default timeout is 30s",
								MarkdownDescription: "serverTimeout defines how long a connection will be held open while waiting for a server/backend response.  If unset, the default timeout is 30s",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"thread_count": schema.Int64Attribute{
								Description:         "threadCount defines the number of threads created per HAProxy process. Creating more threads allows each ingress controller pod to handle more connections, at the cost of more system resources being used. HAProxy currently supports up to 64 threads. If this field is empty, the IngressController will use the default value.  The current default is 4 threads, but this may change in future releases.  Setting this field is generally not recommended. Increasing the number of HAProxy threads allows ingress controller pods to utilize more CPU time under load, potentially starving other pods if set too high. Reducing the number of threads may cause the ingress controller to perform poorly.",
								MarkdownDescription: "threadCount defines the number of threads created per HAProxy process. Creating more threads allows each ingress controller pod to handle more connections, at the cost of more system resources being used. HAProxy currently supports up to 64 threads. If this field is empty, the IngressController will use the default value.  The current default is 4 threads, but this may change in future releases.  Setting this field is generally not recommended. Increasing the number of HAProxy threads allows ingress controller pods to utilize more CPU time under load, potentially starving other pods if set too high. Reducing the number of threads may cause the ingress controller to perform poorly.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(64),
								},
							},

							"tls_inspect_delay": schema.StringAttribute{
								Description:         "tlsInspectDelay defines how long the router can hold data to find a matching route.  Setting this too short can cause the router to fall back to the default certificate for edge-terminated or reencrypt routes even when a better matching certificate could be used.  If unset, the default inspect delay is 5s",
								MarkdownDescription: "tlsInspectDelay defines how long the router can hold data to find a matching route.  Setting this too short can cause the router to fall back to the default certificate for edge-terminated or reencrypt routes even when a better matching certificate could be used.  If unset, the default inspect delay is 5s",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tunnel_timeout": schema.StringAttribute{
								Description:         "tunnelTimeout defines how long a tunnel connection (including websockets) will be held open while the tunnel is idle.  If unset, the default timeout is 1h",
								MarkdownDescription: "tunnelTimeout defines how long a tunnel connection (including websockets) will be held open while the tunnel is idle.  If unset, the default timeout is 1h",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"unsupported_config_overrides": schema.MapAttribute{
						Description:         "unsupportedConfigOverrides allows specifying unsupported configuration options.  Its use is unsupported.",
						MarkdownDescription: "unsupportedConfigOverrides allows specifying unsupported configuration options.  Its use is unsupported.",
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
	}
}

func (r *OperatorOpenshiftIoIngressControllerV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_operator_openshift_io_ingress_controller_v1_manifest")

	var model OperatorOpenshiftIoIngressControllerV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("operator.openshift.io/v1")
	model.Kind = pointer.String("IngressController")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
