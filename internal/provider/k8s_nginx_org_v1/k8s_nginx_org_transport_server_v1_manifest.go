/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package k8s_nginx_org_v1

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
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &K8SNginxOrgTransportServerV1Manifest{}
)

func NewK8SNginxOrgTransportServerV1Manifest() datasource.DataSource {
	return &K8SNginxOrgTransportServerV1Manifest{}
}

type K8SNginxOrgTransportServerV1Manifest struct{}

type K8SNginxOrgTransportServerV1ManifestData struct {
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
		Action *struct {
			Pass *string `tfsdk:"pass" json:"pass,omitempty"`
		} `tfsdk:"action" json:"action,omitempty"`
		Host             *string `tfsdk:"host" json:"host,omitempty"`
		IngressClassName *string `tfsdk:"ingress_class_name" json:"ingressClassName,omitempty"`
		Listener         *struct {
			Name     *string `tfsdk:"name" json:"name,omitempty"`
			Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
		} `tfsdk:"listener" json:"listener,omitempty"`
		ServerSnippets    *string `tfsdk:"server_snippets" json:"serverSnippets,omitempty"`
		SessionParameters *struct {
			Timeout *string `tfsdk:"timeout" json:"timeout,omitempty"`
		} `tfsdk:"session_parameters" json:"sessionParameters,omitempty"`
		StreamSnippets *string `tfsdk:"stream_snippets" json:"streamSnippets,omitempty"`
		Tls            *struct {
			Secret *string `tfsdk:"secret" json:"secret,omitempty"`
		} `tfsdk:"tls" json:"tls,omitempty"`
		UpstreamParameters *struct {
			ConnectTimeout      *string `tfsdk:"connect_timeout" json:"connectTimeout,omitempty"`
			NextUpstream        *bool   `tfsdk:"next_upstream" json:"nextUpstream,omitempty"`
			NextUpstreamTimeout *string `tfsdk:"next_upstream_timeout" json:"nextUpstreamTimeout,omitempty"`
			NextUpstreamTries   *int64  `tfsdk:"next_upstream_tries" json:"nextUpstreamTries,omitempty"`
			UdpRequests         *int64  `tfsdk:"udp_requests" json:"udpRequests,omitempty"`
			UdpResponses        *int64  `tfsdk:"udp_responses" json:"udpResponses,omitempty"`
		} `tfsdk:"upstream_parameters" json:"upstreamParameters,omitempty"`
		Upstreams *[]struct {
			Backup      *string `tfsdk:"backup" json:"backup,omitempty"`
			BackupPort  *int64  `tfsdk:"backup_port" json:"backupPort,omitempty"`
			FailTimeout *string `tfsdk:"fail_timeout" json:"failTimeout,omitempty"`
			HealthCheck *struct {
				Enable   *bool   `tfsdk:"enable" json:"enable,omitempty"`
				Fails    *int64  `tfsdk:"fails" json:"fails,omitempty"`
				Interval *string `tfsdk:"interval" json:"interval,omitempty"`
				Jitter   *string `tfsdk:"jitter" json:"jitter,omitempty"`
				Match    *struct {
					Expect *string `tfsdk:"expect" json:"expect,omitempty"`
					Send   *string `tfsdk:"send" json:"send,omitempty"`
				} `tfsdk:"match" json:"match,omitempty"`
				Passes  *int64  `tfsdk:"passes" json:"passes,omitempty"`
				Port    *int64  `tfsdk:"port" json:"port,omitempty"`
				Timeout *string `tfsdk:"timeout" json:"timeout,omitempty"`
			} `tfsdk:"health_check" json:"healthCheck,omitempty"`
			LoadBalancingMethod *string `tfsdk:"load_balancing_method" json:"loadBalancingMethod,omitempty"`
			MaxConns            *int64  `tfsdk:"max_conns" json:"maxConns,omitempty"`
			MaxFails            *int64  `tfsdk:"max_fails" json:"maxFails,omitempty"`
			Name                *string `tfsdk:"name" json:"name,omitempty"`
			Port                *int64  `tfsdk:"port" json:"port,omitempty"`
			Service             *string `tfsdk:"service" json:"service,omitempty"`
		} `tfsdk:"upstreams" json:"upstreams,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *K8SNginxOrgTransportServerV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_k8s_nginx_org_transport_server_v1_manifest"
}

func (r *K8SNginxOrgTransportServerV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "TransportServer defines the TransportServer resource.",
		MarkdownDescription: "TransportServer defines the TransportServer resource.",
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
				Description:         "TransportServerSpec is the spec of the TransportServer resource.",
				MarkdownDescription: "TransportServerSpec is the spec of the TransportServer resource.",
				Attributes: map[string]schema.Attribute{
					"action": schema.SingleNestedAttribute{
						Description:         "The action to perform for a request.",
						MarkdownDescription: "The action to perform for a request.",
						Attributes: map[string]schema.Attribute{
							"pass": schema.StringAttribute{
								Description:         "Passes connections/datagrams to an upstream. The upstream with that name must be defined in the resource.",
								MarkdownDescription: "Passes connections/datagrams to an upstream. The upstream with that name must be defined in the resource.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"host": schema.StringAttribute{
						Description:         "The host (domain name) of the server. Must be a valid subdomain as defined in RFC 1123, such as my-app or hello.example.com. When using a wildcard domain like *.example.com the domain must be contained in double quotes. The host value needs to be unique among all Ingress and VirtualServer resources.",
						MarkdownDescription: "The host (domain name) of the server. Must be a valid subdomain as defined in RFC 1123, such as my-app or hello.example.com. When using a wildcard domain like *.example.com the domain must be contained in double quotes. The host value needs to be unique among all Ingress and VirtualServer resources.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ingress_class_name": schema.StringAttribute{
						Description:         "Specifies which Ingress Controller must handle the VirtualServer resource.",
						MarkdownDescription: "Specifies which Ingress Controller must handle the VirtualServer resource.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"listener": schema.SingleNestedAttribute{
						Description:         "Sets a custom HTTP and/or HTTPS listener. Valid fields are listener.http and listener.https. Each field must reference the name of a valid listener defined in a GlobalConfiguration resource",
						MarkdownDescription: "Sets a custom HTTP and/or HTTPS listener. Valid fields are listener.http and listener.https. Each field must reference the name of a valid listener defined in a GlobalConfiguration resource",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "The name of a listener defined in a GlobalConfiguration resource.",
								MarkdownDescription: "The name of a listener defined in a GlobalConfiguration resource.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"protocol": schema.StringAttribute{
								Description:         "The protocol of the listener.",
								MarkdownDescription: "The protocol of the listener.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"server_snippets": schema.StringAttribute{
						Description:         "Sets a custom snippet in server context. Overrides the server-snippets ConfigMap key.",
						MarkdownDescription: "Sets a custom snippet in server context. Overrides the server-snippets ConfigMap key.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"session_parameters": schema.SingleNestedAttribute{
						Description:         "The parameters of the session to be used for the Server context",
						MarkdownDescription: "The parameters of the session to be used for the Server context",
						Attributes: map[string]schema.Attribute{
							"timeout": schema.StringAttribute{
								Description:         "The timeout between two successive read or write operations on client or proxied server connections. The default is 10m.",
								MarkdownDescription: "The timeout between two successive read or write operations on client or proxied server connections. The default is 10m.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"stream_snippets": schema.StringAttribute{
						Description:         "Sets a custom snippet in the stream context. Overrides the stream-snippets ConfigMap key.",
						MarkdownDescription: "Sets a custom snippet in the stream context. Overrides the stream-snippets ConfigMap key.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tls": schema.SingleNestedAttribute{
						Description:         "The TLS termination configuration.",
						MarkdownDescription: "The TLS termination configuration.",
						Attributes: map[string]schema.Attribute{
							"secret": schema.StringAttribute{
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

					"upstream_parameters": schema.SingleNestedAttribute{
						Description:         "UpstreamParameters defines parameters for an upstream.",
						MarkdownDescription: "UpstreamParameters defines parameters for an upstream.",
						Attributes: map[string]schema.Attribute{
							"connect_timeout": schema.StringAttribute{
								Description:         "The timeout for establishing a connection with a proxied server. The default is 60s.",
								MarkdownDescription: "The timeout for establishing a connection with a proxied server. The default is 60s.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"next_upstream": schema.BoolAttribute{
								Description:         "If a connection to the proxied server cannot be established, determines whether a client connection will be passed to the next server. The default is true.",
								MarkdownDescription: "If a connection to the proxied server cannot be established, determines whether a client connection will be passed to the next server. The default is true.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"next_upstream_timeout": schema.StringAttribute{
								Description:         "The time allowed to pass a connection to the next server. The default is 0.",
								MarkdownDescription: "The time allowed to pass a connection to the next server. The default is 0.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"next_upstream_tries": schema.Int64Attribute{
								Description:         "The number of tries for passing a connection to the next server. The default is 0.",
								MarkdownDescription: "The number of tries for passing a connection to the next server. The default is 0.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"udp_requests": schema.Int64Attribute{
								Description:         "The number of datagrams, after receiving which, the next datagram from the same client starts a new session. The default is 0.",
								MarkdownDescription: "The number of datagrams, after receiving which, the next datagram from the same client starts a new session. The default is 0.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"udp_responses": schema.Int64Attribute{
								Description:         "The number of datagrams expected from the proxied server in response to a client datagram. By default, the number of datagrams is not limited.",
								MarkdownDescription: "The number of datagrams expected from the proxied server in response to a client datagram. By default, the number of datagrams is not limited.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"upstreams": schema.ListNestedAttribute{
						Description:         "A list of upstreams.",
						MarkdownDescription: "A list of upstreams.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"backup": schema.StringAttribute{
									Description:         "The name of the backup service of type ExternalName. This will be used when the primary servers are unavailable. Note: The parameter cannot be used along with the random, hash or ip_hash load balancing methods.",
									MarkdownDescription: "The name of the backup service of type ExternalName. This will be used when the primary servers are unavailable. Note: The parameter cannot be used along with the random, hash or ip_hash load balancing methods.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"backup_port": schema.Int64Attribute{
									Description:         "The port of the backup service. The backup port is required if the backup service name is provided. The port must fall into the range 1..65535.",
									MarkdownDescription: "The port of the backup service. The backup port is required if the backup service name is provided. The port must fall into the range 1..65535.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"fail_timeout": schema.StringAttribute{
									Description:         "Sets the number of unsuccessful attempts to communicate with the server that should happen in the duration set by the failTimeout parameter to consider the server unavailable. The default is 1.",
									MarkdownDescription: "Sets the number of unsuccessful attempts to communicate with the server that should happen in the duration set by the failTimeout parameter to consider the server unavailable. The default is 1.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"health_check": schema.SingleNestedAttribute{
									Description:         "The health check configuration for the Upstream. Note: this feature is supported only in NGINX Plus.",
									MarkdownDescription: "The health check configuration for the Upstream. Note: this feature is supported only in NGINX Plus.",
									Attributes: map[string]schema.Attribute{
										"enable": schema.BoolAttribute{
											Description:         "Enables a health check for an upstream server. The default is false.",
											MarkdownDescription: "Enables a health check for an upstream server. The default is false.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"fails": schema.Int64Attribute{
											Description:         "The number of consecutive failed health checks of a particular upstream server after which this server will be considered unhealthy. The default is 1.",
											MarkdownDescription: "The number of consecutive failed health checks of a particular upstream server after which this server will be considered unhealthy. The default is 1.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"interval": schema.StringAttribute{
											Description:         "The interval between two consecutive health checks. The default is 5s.",
											MarkdownDescription: "The interval between two consecutive health checks. The default is 5s.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"jitter": schema.StringAttribute{
											Description:         "The time within which each health check will be randomly delayed. By default, there is no delay.",
											MarkdownDescription: "The time within which each health check will be randomly delayed. By default, there is no delay.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"match": schema.SingleNestedAttribute{
											Description:         "Controls the data to send and the response to expect for the healthcheck.",
											MarkdownDescription: "Controls the data to send and the response to expect for the healthcheck.",
											Attributes: map[string]schema.Attribute{
												"expect": schema.StringAttribute{
													Description:         "A literal string or a regular expression that the data obtained from the server should match. The regular expression is specified with the preceding ~* modifier (for case-insensitive matching), or the ~ modifier (for case-sensitive matching). NGINX Ingress Controller validates a regular expression using the RE2 syntax.",
													MarkdownDescription: "A literal string or a regular expression that the data obtained from the server should match. The regular expression is specified with the preceding ~* modifier (for case-insensitive matching), or the ~ modifier (for case-sensitive matching). NGINX Ingress Controller validates a regular expression using the RE2 syntax.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"send": schema.StringAttribute{
													Description:         "A string to send to an upstream server.",
													MarkdownDescription: "A string to send to an upstream server.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"passes": schema.Int64Attribute{
											Description:         "The number of consecutive passed health checks of a particular upstream server after which the server will be considered healthy. The default is 1.",
											MarkdownDescription: "The number of consecutive passed health checks of a particular upstream server after which the server will be considered healthy. The default is 1.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"port": schema.Int64Attribute{
											Description:         "The port used for health check requests. By default, the server port is used. Note: in contrast with the port of the upstream, this port is not a service port, but a port of a pod.",
											MarkdownDescription: "The port used for health check requests. By default, the server port is used. Note: in contrast with the port of the upstream, this port is not a service port, but a port of a pod.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"timeout": schema.StringAttribute{
											Description:         "This overrides the timeout set by proxy_timeout which is set in SessionParameters for health checks. The default value is 5s.",
											MarkdownDescription: "This overrides the timeout set by proxy_timeout which is set in SessionParameters for health checks. The default value is 5s.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"load_balancing_method": schema.StringAttribute{
									Description:         "The method used to load balance the upstream servers. By default, connections are distributed between the servers using a weighted round-robin balancing method.",
									MarkdownDescription: "The method used to load balance the upstream servers. By default, connections are distributed between the servers using a weighted round-robin balancing method.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"max_conns": schema.Int64Attribute{
									Description:         "Sets the time during which the specified number of unsuccessful attempts to communicate with the server should happen to consider the server unavailable and the period of time the server will be considered unavailable. The default is 10s.",
									MarkdownDescription: "Sets the time during which the specified number of unsuccessful attempts to communicate with the server should happen to consider the server unavailable and the period of time the server will be considered unavailable. The default is 10s.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"max_fails": schema.Int64Attribute{
									Description:         "Sets the number of maximum connections to the proxied server. Default value is zero, meaning there is no limit. The default is 0.",
									MarkdownDescription: "Sets the number of maximum connections to the proxied server. Default value is zero, meaning there is no limit. The default is 0.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "The name of the upstream. Must be a valid DNS label as defined in RFC 1035. For example, hello and upstream-123 are valid. The name must be unique among all upstreams of the resource.",
									MarkdownDescription: "The name of the upstream. Must be a valid DNS label as defined in RFC 1035. For example, hello and upstream-123 are valid. The name must be unique among all upstreams of the resource.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"port": schema.Int64Attribute{
									Description:         "The port of the service. If the service doesn’t define that port, NGINX will assume the service has zero endpoints and close client connections/ignore datagrams. The port must fall into the range 1..65535.",
									MarkdownDescription: "The port of the service. If the service doesn’t define that port, NGINX will assume the service has zero endpoints and close client connections/ignore datagrams. The port must fall into the range 1..65535.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"service": schema.StringAttribute{
									Description:         "The name of a service. The service must belong to the same namespace as the resource. If the service doesn’t exist, NGINX will assume the service has zero endpoints and close client connections/ignore datagrams.",
									MarkdownDescription: "The name of a service. The service must belong to the same namespace as the resource. If the service doesn’t exist, NGINX will assume the service has zero endpoints and close client connections/ignore datagrams.",
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

func (r *K8SNginxOrgTransportServerV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_k8s_nginx_org_transport_server_v1_manifest")

	var model K8SNginxOrgTransportServerV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("k8s.nginx.org/v1")
	model.Kind = pointer.String("TransportServer")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
