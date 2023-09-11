/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package traefik_io_v1alpha1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &TraefikIoIngressRouteV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &TraefikIoIngressRouteV1Alpha1DataSource{}
)

func NewTraefikIoIngressRouteV1Alpha1DataSource() datasource.DataSource {
	return &TraefikIoIngressRouteV1Alpha1DataSource{}
}

type TraefikIoIngressRouteV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type TraefikIoIngressRouteV1Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		EntryPoints *[]string `tfsdk:"entry_points" json:"entryPoints,omitempty"`
		Routes      *[]struct {
			Kind        *string `tfsdk:"kind" json:"kind,omitempty"`
			Match       *string `tfsdk:"match" json:"match,omitempty"`
			Middlewares *[]struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"middlewares" json:"middlewares,omitempty"`
			Priority *int64 `tfsdk:"priority" json:"priority,omitempty"`
			Services *[]struct {
				Kind               *string `tfsdk:"kind" json:"kind,omitempty"`
				Name               *string `tfsdk:"name" json:"name,omitempty"`
				Namespace          *string `tfsdk:"namespace" json:"namespace,omitempty"`
				NativeLB           *bool   `tfsdk:"native_lb" json:"nativeLB,omitempty"`
				PassHostHeader     *bool   `tfsdk:"pass_host_header" json:"passHostHeader,omitempty"`
				Port               *string `tfsdk:"port" json:"port,omitempty"`
				ResponseForwarding *struct {
					FlushInterval *string `tfsdk:"flush_interval" json:"flushInterval,omitempty"`
				} `tfsdk:"response_forwarding" json:"responseForwarding,omitempty"`
				Scheme           *string `tfsdk:"scheme" json:"scheme,omitempty"`
				ServersTransport *string `tfsdk:"servers_transport" json:"serversTransport,omitempty"`
				Sticky           *struct {
					Cookie *struct {
						HttpOnly *bool   `tfsdk:"http_only" json:"httpOnly,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						SameSite *string `tfsdk:"same_site" json:"sameSite,omitempty"`
						Secure   *bool   `tfsdk:"secure" json:"secure,omitempty"`
					} `tfsdk:"cookie" json:"cookie,omitempty"`
				} `tfsdk:"sticky" json:"sticky,omitempty"`
				Strategy *string `tfsdk:"strategy" json:"strategy,omitempty"`
				Weight   *int64  `tfsdk:"weight" json:"weight,omitempty"`
			} `tfsdk:"services" json:"services,omitempty"`
		} `tfsdk:"routes" json:"routes,omitempty"`
		Tls *struct {
			CertResolver *string `tfsdk:"cert_resolver" json:"certResolver,omitempty"`
			Domains      *[]struct {
				Main *string   `tfsdk:"main" json:"main,omitempty"`
				Sans *[]string `tfsdk:"sans" json:"sans,omitempty"`
			} `tfsdk:"domains" json:"domains,omitempty"`
			Options *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"options" json:"options,omitempty"`
			SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
			Store      *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"store" json:"store,omitempty"`
		} `tfsdk:"tls" json:"tls,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *TraefikIoIngressRouteV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_traefik_io_ingress_route_v1alpha1"
}

func (r *TraefikIoIngressRouteV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "IngressRoute is the CRD implementation of a Traefik HTTP Router.",
		MarkdownDescription: "IngressRoute is the CRD implementation of a Traefik HTTP Router.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"api_version": schema.StringAttribute{
				Description:         "The API group of the requested resource.",
				MarkdownDescription: "The API group of the requested resource.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"kind": schema.StringAttribute{
				Description:         "The type of the requested resource.",
				MarkdownDescription: "The type of the requested resource.",
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
						Optional:            false,
						Computed:            true,
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "IngressRouteSpec defines the desired state of IngressRoute.",
				MarkdownDescription: "IngressRouteSpec defines the desired state of IngressRoute.",
				Attributes: map[string]schema.Attribute{
					"entry_points": schema.ListAttribute{
						Description:         "EntryPoints defines the list of entry point names to bind to. Entry points have to be configured in the static configuration. More info: https://doc.traefik.io/traefik/v3.0/routing/entrypoints/ Default: all.",
						MarkdownDescription: "EntryPoints defines the list of entry point names to bind to. Entry points have to be configured in the static configuration. More info: https://doc.traefik.io/traefik/v3.0/routing/entrypoints/ Default: all.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"routes": schema.ListNestedAttribute{
						Description:         "Routes defines the list of routes.",
						MarkdownDescription: "Routes defines the list of routes.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"kind": schema.StringAttribute{
									Description:         "Kind defines the kind of the route. Rule is the only supported kind.",
									MarkdownDescription: "Kind defines the kind of the route. Rule is the only supported kind.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"match": schema.StringAttribute{
									Description:         "Match defines the router's rule. More info: https://doc.traefik.io/traefik/v3.0/routing/routers/#rule",
									MarkdownDescription: "Match defines the router's rule. More info: https://doc.traefik.io/traefik/v3.0/routing/routers/#rule",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"middlewares": schema.ListNestedAttribute{
									Description:         "Middlewares defines the list of references to Middleware resources. More info: https://doc.traefik.io/traefik/v3.0/routing/providers/kubernetes-crd/#kind-middleware",
									MarkdownDescription: "Middlewares defines the list of references to Middleware resources. More info: https://doc.traefik.io/traefik/v3.0/routing/providers/kubernetes-crd/#kind-middleware",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name defines the name of the referenced Middleware resource.",
												MarkdownDescription: "Name defines the name of the referenced Middleware resource.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"namespace": schema.StringAttribute{
												Description:         "Namespace defines the namespace of the referenced Middleware resource.",
												MarkdownDescription: "Namespace defines the namespace of the referenced Middleware resource.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"priority": schema.Int64Attribute{
									Description:         "Priority defines the router's priority. More info: https://doc.traefik.io/traefik/v3.0/routing/routers/#priority",
									MarkdownDescription: "Priority defines the router's priority. More info: https://doc.traefik.io/traefik/v3.0/routing/routers/#priority",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"services": schema.ListNestedAttribute{
									Description:         "Services defines the list of Service. It can contain any combination of TraefikService and/or reference to a Kubernetes Service.",
									MarkdownDescription: "Services defines the list of Service. It can contain any combination of TraefikService and/or reference to a Kubernetes Service.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"kind": schema.StringAttribute{
												Description:         "Kind defines the kind of the Service.",
												MarkdownDescription: "Kind defines the kind of the Service.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"name": schema.StringAttribute{
												Description:         "Name defines the name of the referenced Kubernetes Service or TraefikService. The differentiation between the two is specified in the Kind field.",
												MarkdownDescription: "Name defines the name of the referenced Kubernetes Service or TraefikService. The differentiation between the two is specified in the Kind field.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"namespace": schema.StringAttribute{
												Description:         "Namespace defines the namespace of the referenced Kubernetes Service or TraefikService.",
												MarkdownDescription: "Namespace defines the namespace of the referenced Kubernetes Service or TraefikService.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"native_lb": schema.BoolAttribute{
												Description:         "NativeLB controls, when creating the load-balancer, whether the LB's children are directly the pods IPs or if the only child is the Kubernetes Service clusterIP. The Kubernetes Service itself does load-balance to the pods. By default, NativeLB is false.",
												MarkdownDescription: "NativeLB controls, when creating the load-balancer, whether the LB's children are directly the pods IPs or if the only child is the Kubernetes Service clusterIP. The Kubernetes Service itself does load-balance to the pods. By default, NativeLB is false.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"pass_host_header": schema.BoolAttribute{
												Description:         "PassHostHeader defines whether the client Host header is forwarded to the upstream Kubernetes Service. By default, passHostHeader is true.",
												MarkdownDescription: "PassHostHeader defines whether the client Host header is forwarded to the upstream Kubernetes Service. By default, passHostHeader is true.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"port": schema.StringAttribute{
												Description:         "Port defines the port of a Kubernetes Service. This can be a reference to a named port.",
												MarkdownDescription: "Port defines the port of a Kubernetes Service. This can be a reference to a named port.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"response_forwarding": schema.SingleNestedAttribute{
												Description:         "ResponseForwarding defines how Traefik forwards the response from the upstream Kubernetes Service to the client.",
												MarkdownDescription: "ResponseForwarding defines how Traefik forwards the response from the upstream Kubernetes Service to the client.",
												Attributes: map[string]schema.Attribute{
													"flush_interval": schema.StringAttribute{
														Description:         "FlushInterval defines the interval, in milliseconds, in between flushes to the client while copying the response body. A negative value means to flush immediately after each write to the client. This configuration is ignored when ReverseProxy recognizes a response as a streaming response; for such responses, writes are flushed to the client immediately. Default: 100ms",
														MarkdownDescription: "FlushInterval defines the interval, in milliseconds, in between flushes to the client while copying the response body. A negative value means to flush immediately after each write to the client. This configuration is ignored when ReverseProxy recognizes a response as a streaming response; for such responses, writes are flushed to the client immediately. Default: 100ms",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"scheme": schema.StringAttribute{
												Description:         "Scheme defines the scheme to use for the request to the upstream Kubernetes Service. It defaults to https when Kubernetes Service port is 443, http otherwise.",
												MarkdownDescription: "Scheme defines the scheme to use for the request to the upstream Kubernetes Service. It defaults to https when Kubernetes Service port is 443, http otherwise.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"servers_transport": schema.StringAttribute{
												Description:         "ServersTransport defines the name of ServersTransport resource to use. It allows to configure the transport between Traefik and your servers. Can only be used on a Kubernetes Service.",
												MarkdownDescription: "ServersTransport defines the name of ServersTransport resource to use. It allows to configure the transport between Traefik and your servers. Can only be used on a Kubernetes Service.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"sticky": schema.SingleNestedAttribute{
												Description:         "Sticky defines the sticky sessions configuration. More info: https://doc.traefik.io/traefik/v3.0/routing/services/#sticky-sessions",
												MarkdownDescription: "Sticky defines the sticky sessions configuration. More info: https://doc.traefik.io/traefik/v3.0/routing/services/#sticky-sessions",
												Attributes: map[string]schema.Attribute{
													"cookie": schema.SingleNestedAttribute{
														Description:         "Cookie defines the sticky cookie configuration.",
														MarkdownDescription: "Cookie defines the sticky cookie configuration.",
														Attributes: map[string]schema.Attribute{
															"http_only": schema.BoolAttribute{
																Description:         "HTTPOnly defines whether the cookie can be accessed by client-side APIs, such as JavaScript.",
																MarkdownDescription: "HTTPOnly defines whether the cookie can be accessed by client-side APIs, such as JavaScript.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "Name defines the Cookie name.",
																MarkdownDescription: "Name defines the Cookie name.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"same_site": schema.StringAttribute{
																Description:         "SameSite defines the same site policy. More info: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie/SameSite",
																MarkdownDescription: "SameSite defines the same site policy. More info: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie/SameSite",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secure": schema.BoolAttribute{
																Description:         "Secure defines whether the cookie can only be transmitted over an encrypted connection (i.e. HTTPS).",
																MarkdownDescription: "Secure defines whether the cookie can only be transmitted over an encrypted connection (i.e. HTTPS).",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"strategy": schema.StringAttribute{
												Description:         "Strategy defines the load balancing strategy between the servers. RoundRobin is the only supported value at the moment.",
												MarkdownDescription: "Strategy defines the load balancing strategy between the servers. RoundRobin is the only supported value at the moment.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"weight": schema.Int64Attribute{
												Description:         "Weight defines the weight and should only be specified when Name references a TraefikService object (and to be precise, one that embeds a Weighted Round Robin).",
												MarkdownDescription: "Weight defines the weight and should only be specified when Name references a TraefikService object (and to be precise, one that embeds a Weighted Round Robin).",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"tls": schema.SingleNestedAttribute{
						Description:         "TLS defines the TLS configuration. More info: https://doc.traefik.io/traefik/v3.0/routing/routers/#tls",
						MarkdownDescription: "TLS defines the TLS configuration. More info: https://doc.traefik.io/traefik/v3.0/routing/routers/#tls",
						Attributes: map[string]schema.Attribute{
							"cert_resolver": schema.StringAttribute{
								Description:         "CertResolver defines the name of the certificate resolver to use. Cert resolvers have to be configured in the static configuration. More info: https://doc.traefik.io/traefik/v3.0/https/acme/#certificate-resolvers",
								MarkdownDescription: "CertResolver defines the name of the certificate resolver to use. Cert resolvers have to be configured in the static configuration. More info: https://doc.traefik.io/traefik/v3.0/https/acme/#certificate-resolvers",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"domains": schema.ListNestedAttribute{
								Description:         "Domains defines the list of domains that will be used to issue certificates. More info: https://doc.traefik.io/traefik/v3.0/routing/routers/#domains",
								MarkdownDescription: "Domains defines the list of domains that will be used to issue certificates. More info: https://doc.traefik.io/traefik/v3.0/routing/routers/#domains",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"main": schema.StringAttribute{
											Description:         "Main defines the main domain name.",
											MarkdownDescription: "Main defines the main domain name.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"sans": schema.ListAttribute{
											Description:         "SANs defines the subject alternative domain names.",
											MarkdownDescription: "SANs defines the subject alternative domain names.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"options": schema.SingleNestedAttribute{
								Description:         "Options defines the reference to a TLSOption, that specifies the parameters of the TLS connection. If not defined, the 'default' TLSOption is used. More info: https://doc.traefik.io/traefik/v3.0/https/tls/#tls-options",
								MarkdownDescription: "Options defines the reference to a TLSOption, that specifies the parameters of the TLS connection. If not defined, the 'default' TLSOption is used. More info: https://doc.traefik.io/traefik/v3.0/https/tls/#tls-options",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name defines the name of the referenced TLSOption. More info: https://doc.traefik.io/traefik/v3.0/routing/providers/kubernetes-crd/#kind-tlsoption",
										MarkdownDescription: "Name defines the name of the referenced TLSOption. More info: https://doc.traefik.io/traefik/v3.0/routing/providers/kubernetes-crd/#kind-tlsoption",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"namespace": schema.StringAttribute{
										Description:         "Namespace defines the namespace of the referenced TLSOption. More info: https://doc.traefik.io/traefik/v3.0/routing/providers/kubernetes-crd/#kind-tlsoption",
										MarkdownDescription: "Namespace defines the namespace of the referenced TLSOption. More info: https://doc.traefik.io/traefik/v3.0/routing/providers/kubernetes-crd/#kind-tlsoption",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"secret_name": schema.StringAttribute{
								Description:         "SecretName is the name of the referenced Kubernetes Secret to specify the certificate details.",
								MarkdownDescription: "SecretName is the name of the referenced Kubernetes Secret to specify the certificate details.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"store": schema.SingleNestedAttribute{
								Description:         "Store defines the reference to the TLSStore, that will be used to store certificates. Please note that only 'default' TLSStore can be used.",
								MarkdownDescription: "Store defines the reference to the TLSStore, that will be used to store certificates. Please note that only 'default' TLSStore can be used.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name defines the name of the referenced TLSStore. More info: https://doc.traefik.io/traefik/v3.0/routing/providers/kubernetes-crd/#kind-tlsstore",
										MarkdownDescription: "Name defines the name of the referenced TLSStore. More info: https://doc.traefik.io/traefik/v3.0/routing/providers/kubernetes-crd/#kind-tlsstore",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"namespace": schema.StringAttribute{
										Description:         "Namespace defines the namespace of the referenced TLSStore. More info: https://doc.traefik.io/traefik/v3.0/routing/providers/kubernetes-crd/#kind-tlsstore",
										MarkdownDescription: "Namespace defines the namespace of the referenced TLSStore. More info: https://doc.traefik.io/traefik/v3.0/routing/providers/kubernetes-crd/#kind-tlsstore",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *TraefikIoIngressRouteV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedDataSourceDataError(request.ProviderData))
	}
}

func (r *TraefikIoIngressRouteV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_traefik_io_ingress_route_v1alpha1")

	var data TraefikIoIngressRouteV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "traefik.io", Version: "v1alpha1", Resource: "ingressroutes"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetNamespacedResourceError(err, data.Metadata.Name, data.Metadata.Namespace))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse TraefikIoIngressRouteV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("traefik.io/v1alpha1")
	data.Kind = pointer.String("IngressRoute")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
