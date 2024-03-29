/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package networking_istio_io_v1alpha3

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &NetworkingIstioIoSidecarV1Alpha3Manifest{}
)

func NewNetworkingIstioIoSidecarV1Alpha3Manifest() datasource.DataSource {
	return &NetworkingIstioIoSidecarV1Alpha3Manifest{}
}

type NetworkingIstioIoSidecarV1Alpha3Manifest struct{}

type NetworkingIstioIoSidecarV1Alpha3ManifestData struct {
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
		Egress *[]struct {
			Bind        *string   `tfsdk:"bind" json:"bind,omitempty"`
			CaptureMode *string   `tfsdk:"capture_mode" json:"captureMode,omitempty"`
			Hosts       *[]string `tfsdk:"hosts" json:"hosts,omitempty"`
			Port        *struct {
				Name       *string `tfsdk:"name" json:"name,omitempty"`
				Number     *int64  `tfsdk:"number" json:"number,omitempty"`
				Protocol   *string `tfsdk:"protocol" json:"protocol,omitempty"`
				TargetPort *int64  `tfsdk:"target_port" json:"targetPort,omitempty"`
			} `tfsdk:"port" json:"port,omitempty"`
		} `tfsdk:"egress" json:"egress,omitempty"`
		InboundConnectionPool *struct {
			Http *struct {
				H2UpgradePolicy          *string `tfsdk:"h2_upgrade_policy" json:"h2UpgradePolicy,omitempty"`
				Http1MaxPendingRequests  *int64  `tfsdk:"http1_max_pending_requests" json:"http1MaxPendingRequests,omitempty"`
				Http2MaxRequests         *int64  `tfsdk:"http2_max_requests" json:"http2MaxRequests,omitempty"`
				IdleTimeout              *string `tfsdk:"idle_timeout" json:"idleTimeout,omitempty"`
				MaxConcurrentStreams     *int64  `tfsdk:"max_concurrent_streams" json:"maxConcurrentStreams,omitempty"`
				MaxRequestsPerConnection *int64  `tfsdk:"max_requests_per_connection" json:"maxRequestsPerConnection,omitempty"`
				MaxRetries               *int64  `tfsdk:"max_retries" json:"maxRetries,omitempty"`
				UseClientProtocol        *bool   `tfsdk:"use_client_protocol" json:"useClientProtocol,omitempty"`
			} `tfsdk:"http" json:"http,omitempty"`
			Tcp *struct {
				ConnectTimeout        *string `tfsdk:"connect_timeout" json:"connectTimeout,omitempty"`
				IdleTimeout           *string `tfsdk:"idle_timeout" json:"idleTimeout,omitempty"`
				MaxConnectionDuration *string `tfsdk:"max_connection_duration" json:"maxConnectionDuration,omitempty"`
				MaxConnections        *int64  `tfsdk:"max_connections" json:"maxConnections,omitempty"`
				TcpKeepalive          *struct {
					Interval *string `tfsdk:"interval" json:"interval,omitempty"`
					Probes   *int64  `tfsdk:"probes" json:"probes,omitempty"`
					Time     *string `tfsdk:"time" json:"time,omitempty"`
				} `tfsdk:"tcp_keepalive" json:"tcpKeepalive,omitempty"`
			} `tfsdk:"tcp" json:"tcp,omitempty"`
		} `tfsdk:"inbound_connection_pool" json:"inboundConnectionPool,omitempty"`
		Ingress *[]struct {
			Bind           *string `tfsdk:"bind" json:"bind,omitempty"`
			CaptureMode    *string `tfsdk:"capture_mode" json:"captureMode,omitempty"`
			ConnectionPool *struct {
				Http *struct {
					H2UpgradePolicy          *string `tfsdk:"h2_upgrade_policy" json:"h2UpgradePolicy,omitempty"`
					Http1MaxPendingRequests  *int64  `tfsdk:"http1_max_pending_requests" json:"http1MaxPendingRequests,omitempty"`
					Http2MaxRequests         *int64  `tfsdk:"http2_max_requests" json:"http2MaxRequests,omitempty"`
					IdleTimeout              *string `tfsdk:"idle_timeout" json:"idleTimeout,omitempty"`
					MaxConcurrentStreams     *int64  `tfsdk:"max_concurrent_streams" json:"maxConcurrentStreams,omitempty"`
					MaxRequestsPerConnection *int64  `tfsdk:"max_requests_per_connection" json:"maxRequestsPerConnection,omitempty"`
					MaxRetries               *int64  `tfsdk:"max_retries" json:"maxRetries,omitempty"`
					UseClientProtocol        *bool   `tfsdk:"use_client_protocol" json:"useClientProtocol,omitempty"`
				} `tfsdk:"http" json:"http,omitempty"`
				Tcp *struct {
					ConnectTimeout        *string `tfsdk:"connect_timeout" json:"connectTimeout,omitempty"`
					IdleTimeout           *string `tfsdk:"idle_timeout" json:"idleTimeout,omitempty"`
					MaxConnectionDuration *string `tfsdk:"max_connection_duration" json:"maxConnectionDuration,omitempty"`
					MaxConnections        *int64  `tfsdk:"max_connections" json:"maxConnections,omitempty"`
					TcpKeepalive          *struct {
						Interval *string `tfsdk:"interval" json:"interval,omitempty"`
						Probes   *int64  `tfsdk:"probes" json:"probes,omitempty"`
						Time     *string `tfsdk:"time" json:"time,omitempty"`
					} `tfsdk:"tcp_keepalive" json:"tcpKeepalive,omitempty"`
				} `tfsdk:"tcp" json:"tcp,omitempty"`
			} `tfsdk:"connection_pool" json:"connectionPool,omitempty"`
			DefaultEndpoint *string `tfsdk:"default_endpoint" json:"defaultEndpoint,omitempty"`
			Port            *struct {
				Name       *string `tfsdk:"name" json:"name,omitempty"`
				Number     *int64  `tfsdk:"number" json:"number,omitempty"`
				Protocol   *string `tfsdk:"protocol" json:"protocol,omitempty"`
				TargetPort *int64  `tfsdk:"target_port" json:"targetPort,omitempty"`
			} `tfsdk:"port" json:"port,omitempty"`
			Tls *struct {
				CaCertificates        *string   `tfsdk:"ca_certificates" json:"caCertificates,omitempty"`
				CaCrl                 *string   `tfsdk:"ca_crl" json:"caCrl,omitempty"`
				CipherSuites          *[]string `tfsdk:"cipher_suites" json:"cipherSuites,omitempty"`
				CredentialName        *string   `tfsdk:"credential_name" json:"credentialName,omitempty"`
				HttpsRedirect         *bool     `tfsdk:"https_redirect" json:"httpsRedirect,omitempty"`
				MaxProtocolVersion    *string   `tfsdk:"max_protocol_version" json:"maxProtocolVersion,omitempty"`
				MinProtocolVersion    *string   `tfsdk:"min_protocol_version" json:"minProtocolVersion,omitempty"`
				Mode                  *string   `tfsdk:"mode" json:"mode,omitempty"`
				PrivateKey            *string   `tfsdk:"private_key" json:"privateKey,omitempty"`
				ServerCertificate     *string   `tfsdk:"server_certificate" json:"serverCertificate,omitempty"`
				SubjectAltNames       *[]string `tfsdk:"subject_alt_names" json:"subjectAltNames,omitempty"`
				VerifyCertificateHash *[]string `tfsdk:"verify_certificate_hash" json:"verifyCertificateHash,omitempty"`
				VerifyCertificateSpki *[]string `tfsdk:"verify_certificate_spki" json:"verifyCertificateSpki,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
		} `tfsdk:"ingress" json:"ingress,omitempty"`
		OutboundTrafficPolicy *struct {
			EgressProxy *struct {
				Host *string `tfsdk:"host" json:"host,omitempty"`
				Port *struct {
					Number *int64 `tfsdk:"number" json:"number,omitempty"`
				} `tfsdk:"port" json:"port,omitempty"`
				Subset *string `tfsdk:"subset" json:"subset,omitempty"`
			} `tfsdk:"egress_proxy" json:"egressProxy,omitempty"`
			Mode *string `tfsdk:"mode" json:"mode,omitempty"`
		} `tfsdk:"outbound_traffic_policy" json:"outboundTrafficPolicy,omitempty"`
		WorkloadSelector *struct {
			Labels *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		} `tfsdk:"workload_selector" json:"workloadSelector,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *NetworkingIstioIoSidecarV1Alpha3Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_networking_istio_io_sidecar_v1alpha3_manifest"
}

func (r *NetworkingIstioIoSidecarV1Alpha3Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "Configuration affecting network reachability of a sidecar. See more details at: https://istio.io/docs/reference/config/networking/sidecar.html",
				MarkdownDescription: "Configuration affecting network reachability of a sidecar. See more details at: https://istio.io/docs/reference/config/networking/sidecar.html",
				Attributes: map[string]schema.Attribute{
					"egress": schema.ListNestedAttribute{
						Description:         "Egress specifies the configuration of the sidecar for processing outbound traffic from the attached workload instance to other services in the mesh.",
						MarkdownDescription: "Egress specifies the configuration of the sidecar for processing outbound traffic from the attached workload instance to other services in the mesh.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"bind": schema.StringAttribute{
									Description:         "The IP(IPv4 or IPv6) or the Unix domain socket to which the listener should be bound to.",
									MarkdownDescription: "The IP(IPv4 or IPv6) or the Unix domain socket to which the listener should be bound to.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"capture_mode": schema.StringAttribute{
									Description:         "When the bind address is an IP, the captureMode option dictates how traffic to the listener is expected to be captured (or not).Valid Options: DEFAULT, IPTABLES, NONE",
									MarkdownDescription: "When the bind address is an IP, the captureMode option dictates how traffic to the listener is expected to be captured (or not).Valid Options: DEFAULT, IPTABLES, NONE",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("DEFAULT", "IPTABLES", "NONE"),
									},
								},

								"hosts": schema.ListAttribute{
									Description:         "One or more service hosts exposed by the listener in 'namespace/dnsName' format.",
									MarkdownDescription: "One or more service hosts exposed by the listener in 'namespace/dnsName' format.",
									ElementType:         types.StringType,
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"port": schema.SingleNestedAttribute{
									Description:         "The port associated with the listener.",
									MarkdownDescription: "The port associated with the listener.",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Label assigned to the port.",
											MarkdownDescription: "Label assigned to the port.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"number": schema.Int64Attribute{
											Description:         "A valid non-negative integer port number.",
											MarkdownDescription: "A valid non-negative integer port number.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"protocol": schema.StringAttribute{
											Description:         "The protocol exposed on the port.",
											MarkdownDescription: "The protocol exposed on the port.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"target_port": schema.Int64Attribute{
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

					"inbound_connection_pool": schema.SingleNestedAttribute{
						Description:         "Settings controlling the volume of connections Envoy will accept from the network.",
						MarkdownDescription: "Settings controlling the volume of connections Envoy will accept from the network.",
						Attributes: map[string]schema.Attribute{
							"http": schema.SingleNestedAttribute{
								Description:         "HTTP connection pool settings.",
								MarkdownDescription: "HTTP connection pool settings.",
								Attributes: map[string]schema.Attribute{
									"h2_upgrade_policy": schema.StringAttribute{
										Description:         "Specify if http1.1 connection should be upgraded to http2 for the associated destination.Valid Options: DEFAULT, DO_NOT_UPGRADE, UPGRADE",
										MarkdownDescription: "Specify if http1.1 connection should be upgraded to http2 for the associated destination.Valid Options: DEFAULT, DO_NOT_UPGRADE, UPGRADE",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("DEFAULT", "DO_NOT_UPGRADE", "UPGRADE"),
										},
									},

									"http1_max_pending_requests": schema.Int64Attribute{
										Description:         "Maximum number of requests that will be queued while waiting for a ready connection pool connection.",
										MarkdownDescription: "Maximum number of requests that will be queued while waiting for a ready connection pool connection.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"http2_max_requests": schema.Int64Attribute{
										Description:         "Maximum number of active requests to a destination.",
										MarkdownDescription: "Maximum number of active requests to a destination.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"idle_timeout": schema.StringAttribute{
										Description:         "The idle timeout for upstream connection pool connections.",
										MarkdownDescription: "The idle timeout for upstream connection pool connections.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_concurrent_streams": schema.Int64Attribute{
										Description:         "The maximum number of concurrent streams allowed for a peer on one HTTP/2 connection.",
										MarkdownDescription: "The maximum number of concurrent streams allowed for a peer on one HTTP/2 connection.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_requests_per_connection": schema.Int64Attribute{
										Description:         "Maximum number of requests per connection to a backend.",
										MarkdownDescription: "Maximum number of requests per connection to a backend.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_retries": schema.Int64Attribute{
										Description:         "Maximum number of retries that can be outstanding to all hosts in a cluster at a given time.",
										MarkdownDescription: "Maximum number of retries that can be outstanding to all hosts in a cluster at a given time.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"use_client_protocol": schema.BoolAttribute{
										Description:         "If set to true, client protocol will be preserved while initiating connection to backend.",
										MarkdownDescription: "If set to true, client protocol will be preserved while initiating connection to backend.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"tcp": schema.SingleNestedAttribute{
								Description:         "Settings common to both HTTP and TCP upstream connections.",
								MarkdownDescription: "Settings common to both HTTP and TCP upstream connections.",
								Attributes: map[string]schema.Attribute{
									"connect_timeout": schema.StringAttribute{
										Description:         "TCP connection timeout.",
										MarkdownDescription: "TCP connection timeout.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"idle_timeout": schema.StringAttribute{
										Description:         "The idle timeout for TCP connections.",
										MarkdownDescription: "The idle timeout for TCP connections.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_connection_duration": schema.StringAttribute{
										Description:         "The maximum duration of a connection.",
										MarkdownDescription: "The maximum duration of a connection.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_connections": schema.Int64Attribute{
										Description:         "Maximum number of HTTP1 /TCP connections to a destination host.",
										MarkdownDescription: "Maximum number of HTTP1 /TCP connections to a destination host.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tcp_keepalive": schema.SingleNestedAttribute{
										Description:         "If set then set SO_KEEPALIVE on the socket to enable TCP Keepalives.",
										MarkdownDescription: "If set then set SO_KEEPALIVE on the socket to enable TCP Keepalives.",
										Attributes: map[string]schema.Attribute{
											"interval": schema.StringAttribute{
												Description:         "The time duration between keep-alive probes.",
												MarkdownDescription: "The time duration between keep-alive probes.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"probes": schema.Int64Attribute{
												Description:         "Maximum number of keepalive probes to send without response before deciding the connection is dead.",
												MarkdownDescription: "Maximum number of keepalive probes to send without response before deciding the connection is dead.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"time": schema.StringAttribute{
												Description:         "The time duration a connection needs to be idle before keep-alive probes start being sent.",
												MarkdownDescription: "The time duration a connection needs to be idle before keep-alive probes start being sent.",
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

					"ingress": schema.ListNestedAttribute{
						Description:         "Ingress specifies the configuration of the sidecar for processing inbound traffic to the attached workload instance.",
						MarkdownDescription: "Ingress specifies the configuration of the sidecar for processing inbound traffic to the attached workload instance.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"bind": schema.StringAttribute{
									Description:         "The IP(IPv4 or IPv6) to which the listener should be bound.",
									MarkdownDescription: "The IP(IPv4 or IPv6) to which the listener should be bound.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"capture_mode": schema.StringAttribute{
									Description:         "The captureMode option dictates how traffic to the listener is expected to be captured (or not).Valid Options: DEFAULT, IPTABLES, NONE",
									MarkdownDescription: "The captureMode option dictates how traffic to the listener is expected to be captured (or not).Valid Options: DEFAULT, IPTABLES, NONE",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("DEFAULT", "IPTABLES", "NONE"),
									},
								},

								"connection_pool": schema.SingleNestedAttribute{
									Description:         "Settings controlling the volume of connections Envoy will accept from the network.",
									MarkdownDescription: "Settings controlling the volume of connections Envoy will accept from the network.",
									Attributes: map[string]schema.Attribute{
										"http": schema.SingleNestedAttribute{
											Description:         "HTTP connection pool settings.",
											MarkdownDescription: "HTTP connection pool settings.",
											Attributes: map[string]schema.Attribute{
												"h2_upgrade_policy": schema.StringAttribute{
													Description:         "Specify if http1.1 connection should be upgraded to http2 for the associated destination.Valid Options: DEFAULT, DO_NOT_UPGRADE, UPGRADE",
													MarkdownDescription: "Specify if http1.1 connection should be upgraded to http2 for the associated destination.Valid Options: DEFAULT, DO_NOT_UPGRADE, UPGRADE",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("DEFAULT", "DO_NOT_UPGRADE", "UPGRADE"),
													},
												},

												"http1_max_pending_requests": schema.Int64Attribute{
													Description:         "Maximum number of requests that will be queued while waiting for a ready connection pool connection.",
													MarkdownDescription: "Maximum number of requests that will be queued while waiting for a ready connection pool connection.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"http2_max_requests": schema.Int64Attribute{
													Description:         "Maximum number of active requests to a destination.",
													MarkdownDescription: "Maximum number of active requests to a destination.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"idle_timeout": schema.StringAttribute{
													Description:         "The idle timeout for upstream connection pool connections.",
													MarkdownDescription: "The idle timeout for upstream connection pool connections.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"max_concurrent_streams": schema.Int64Attribute{
													Description:         "The maximum number of concurrent streams allowed for a peer on one HTTP/2 connection.",
													MarkdownDescription: "The maximum number of concurrent streams allowed for a peer on one HTTP/2 connection.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"max_requests_per_connection": schema.Int64Attribute{
													Description:         "Maximum number of requests per connection to a backend.",
													MarkdownDescription: "Maximum number of requests per connection to a backend.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"max_retries": schema.Int64Attribute{
													Description:         "Maximum number of retries that can be outstanding to all hosts in a cluster at a given time.",
													MarkdownDescription: "Maximum number of retries that can be outstanding to all hosts in a cluster at a given time.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"use_client_protocol": schema.BoolAttribute{
													Description:         "If set to true, client protocol will be preserved while initiating connection to backend.",
													MarkdownDescription: "If set to true, client protocol will be preserved while initiating connection to backend.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"tcp": schema.SingleNestedAttribute{
											Description:         "Settings common to both HTTP and TCP upstream connections.",
											MarkdownDescription: "Settings common to both HTTP and TCP upstream connections.",
											Attributes: map[string]schema.Attribute{
												"connect_timeout": schema.StringAttribute{
													Description:         "TCP connection timeout.",
													MarkdownDescription: "TCP connection timeout.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"idle_timeout": schema.StringAttribute{
													Description:         "The idle timeout for TCP connections.",
													MarkdownDescription: "The idle timeout for TCP connections.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"max_connection_duration": schema.StringAttribute{
													Description:         "The maximum duration of a connection.",
													MarkdownDescription: "The maximum duration of a connection.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"max_connections": schema.Int64Attribute{
													Description:         "Maximum number of HTTP1 /TCP connections to a destination host.",
													MarkdownDescription: "Maximum number of HTTP1 /TCP connections to a destination host.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"tcp_keepalive": schema.SingleNestedAttribute{
													Description:         "If set then set SO_KEEPALIVE on the socket to enable TCP Keepalives.",
													MarkdownDescription: "If set then set SO_KEEPALIVE on the socket to enable TCP Keepalives.",
													Attributes: map[string]schema.Attribute{
														"interval": schema.StringAttribute{
															Description:         "The time duration between keep-alive probes.",
															MarkdownDescription: "The time duration between keep-alive probes.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"probes": schema.Int64Attribute{
															Description:         "Maximum number of keepalive probes to send without response before deciding the connection is dead.",
															MarkdownDescription: "Maximum number of keepalive probes to send without response before deciding the connection is dead.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"time": schema.StringAttribute{
															Description:         "The time duration a connection needs to be idle before keep-alive probes start being sent.",
															MarkdownDescription: "The time duration a connection needs to be idle before keep-alive probes start being sent.",
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

								"default_endpoint": schema.StringAttribute{
									Description:         "The IP endpoint or Unix domain socket to which traffic should be forwarded to.",
									MarkdownDescription: "The IP endpoint or Unix domain socket to which traffic should be forwarded to.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"port": schema.SingleNestedAttribute{
									Description:         "The port associated with the listener.",
									MarkdownDescription: "The port associated with the listener.",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Label assigned to the port.",
											MarkdownDescription: "Label assigned to the port.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"number": schema.Int64Attribute{
											Description:         "A valid non-negative integer port number.",
											MarkdownDescription: "A valid non-negative integer port number.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"protocol": schema.StringAttribute{
											Description:         "The protocol exposed on the port.",
											MarkdownDescription: "The protocol exposed on the port.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"target_port": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},

								"tls": schema.SingleNestedAttribute{
									Description:         "Set of TLS related options that will enable TLS termination on the sidecar for requests originating from outside the mesh.",
									MarkdownDescription: "Set of TLS related options that will enable TLS termination on the sidecar for requests originating from outside the mesh.",
									Attributes: map[string]schema.Attribute{
										"ca_certificates": schema.StringAttribute{
											Description:         "REQUIRED if mode is 'MUTUAL' or 'OPTIONAL_MUTUAL'.",
											MarkdownDescription: "REQUIRED if mode is 'MUTUAL' or 'OPTIONAL_MUTUAL'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ca_crl": schema.StringAttribute{
											Description:         "OPTIONAL: The path to the file containing the certificate revocation list (CRL) to use in verifying a presented client side certificate.",
											MarkdownDescription: "OPTIONAL: The path to the file containing the certificate revocation list (CRL) to use in verifying a presented client side certificate.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"cipher_suites": schema.ListAttribute{
											Description:         "Optional: If specified, only support the specified cipher list.",
											MarkdownDescription: "Optional: If specified, only support the specified cipher list.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"credential_name": schema.StringAttribute{
											Description:         "For gateways running on Kubernetes, the name of the secret that holds the TLS certs including the CA certificates.",
											MarkdownDescription: "For gateways running on Kubernetes, the name of the secret that holds the TLS certs including the CA certificates.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"https_redirect": schema.BoolAttribute{
											Description:         "If set to true, the load balancer will send a 301 redirect for all http connections, asking the clients to use HTTPS.",
											MarkdownDescription: "If set to true, the load balancer will send a 301 redirect for all http connections, asking the clients to use HTTPS.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"max_protocol_version": schema.StringAttribute{
											Description:         "Optional: Maximum TLS protocol version.Valid Options: TLS_AUTO, TLSV1_0, TLSV1_1, TLSV1_2, TLSV1_3",
											MarkdownDescription: "Optional: Maximum TLS protocol version.Valid Options: TLS_AUTO, TLSV1_0, TLSV1_1, TLSV1_2, TLSV1_3",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("TLS_AUTO", "TLSV1_0", "TLSV1_1", "TLSV1_2", "TLSV1_3"),
											},
										},

										"min_protocol_version": schema.StringAttribute{
											Description:         "Optional: Minimum TLS protocol version.Valid Options: TLS_AUTO, TLSV1_0, TLSV1_1, TLSV1_2, TLSV1_3",
											MarkdownDescription: "Optional: Minimum TLS protocol version.Valid Options: TLS_AUTO, TLSV1_0, TLSV1_1, TLSV1_2, TLSV1_3",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("TLS_AUTO", "TLSV1_0", "TLSV1_1", "TLSV1_2", "TLSV1_3"),
											},
										},

										"mode": schema.StringAttribute{
											Description:         "Optional: Indicates whether connections to this port should be secured using TLS.Valid Options: PASSTHROUGH, SIMPLE, MUTUAL, AUTO_PASSTHROUGH, ISTIO_MUTUAL, OPTIONAL_MUTUAL",
											MarkdownDescription: "Optional: Indicates whether connections to this port should be secured using TLS.Valid Options: PASSTHROUGH, SIMPLE, MUTUAL, AUTO_PASSTHROUGH, ISTIO_MUTUAL, OPTIONAL_MUTUAL",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("PASSTHROUGH", "SIMPLE", "MUTUAL", "AUTO_PASSTHROUGH", "ISTIO_MUTUAL", "OPTIONAL_MUTUAL"),
											},
										},

										"private_key": schema.StringAttribute{
											Description:         "REQUIRED if mode is 'SIMPLE' or 'MUTUAL'.",
											MarkdownDescription: "REQUIRED if mode is 'SIMPLE' or 'MUTUAL'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"server_certificate": schema.StringAttribute{
											Description:         "REQUIRED if mode is 'SIMPLE' or 'MUTUAL'.",
											MarkdownDescription: "REQUIRED if mode is 'SIMPLE' or 'MUTUAL'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"subject_alt_names": schema.ListAttribute{
											Description:         "A list of alternate names to verify the subject identity in the certificate presented by the client.",
											MarkdownDescription: "A list of alternate names to verify the subject identity in the certificate presented by the client.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"verify_certificate_hash": schema.ListAttribute{
											Description:         "An optional list of hex-encoded SHA-256 hashes of the authorized client certificates.",
											MarkdownDescription: "An optional list of hex-encoded SHA-256 hashes of the authorized client certificates.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"verify_certificate_spki": schema.ListAttribute{
											Description:         "An optional list of base64-encoded SHA-256 hashes of the SPKIs of authorized client certificates.",
											MarkdownDescription: "An optional list of base64-encoded SHA-256 hashes of the SPKIs of authorized client certificates.",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"outbound_traffic_policy": schema.SingleNestedAttribute{
						Description:         "Configuration for the outbound traffic policy.",
						MarkdownDescription: "Configuration for the outbound traffic policy.",
						Attributes: map[string]schema.Attribute{
							"egress_proxy": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"host": schema.StringAttribute{
										Description:         "The name of a service from the service registry.",
										MarkdownDescription: "The name of a service from the service registry.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"port": schema.SingleNestedAttribute{
										Description:         "Specifies the port on the host that is being addressed.",
										MarkdownDescription: "Specifies the port on the host that is being addressed.",
										Attributes: map[string]schema.Attribute{
											"number": schema.Int64Attribute{
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

									"subset": schema.StringAttribute{
										Description:         "The name of a subset within the service.",
										MarkdownDescription: "The name of a subset within the service.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"mode": schema.StringAttribute{
								Description:         "Valid Options: REGISTRY_ONLY, ALLOW_ANY",
								MarkdownDescription: "Valid Options: REGISTRY_ONLY, ALLOW_ANY",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("REGISTRY_ONLY", "ALLOW_ANY"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"workload_selector": schema.SingleNestedAttribute{
						Description:         "Criteria used to select the specific set of pods/VMs on which this 'Sidecar' configuration should be applied.",
						MarkdownDescription: "Criteria used to select the specific set of pods/VMs on which this 'Sidecar' configuration should be applied.",
						Attributes: map[string]schema.Attribute{
							"labels": schema.MapAttribute{
								Description:         "One or more labels that indicate a specific set of pods/VMs on which the configuration should be applied.",
								MarkdownDescription: "One or more labels that indicate a specific set of pods/VMs on which the configuration should be applied.",
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
	}
}

func (r *NetworkingIstioIoSidecarV1Alpha3Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_networking_istio_io_sidecar_v1alpha3_manifest")

	var model NetworkingIstioIoSidecarV1Alpha3ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("networking.istio.io/v1alpha3")
	model.Kind = pointer.String("Sidecar")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
