/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package traefik_io_v1alpha1

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
	_ datasource.DataSource = &TraefikIoIngressRouteTcpV1Alpha1Manifest{}
)

func NewTraefikIoIngressRouteTcpV1Alpha1Manifest() datasource.DataSource {
	return &TraefikIoIngressRouteTcpV1Alpha1Manifest{}
}

type TraefikIoIngressRouteTcpV1Alpha1Manifest struct{}

type TraefikIoIngressRouteTcpV1Alpha1ManifestData struct {
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
		EntryPoints *[]string `tfsdk:"entry_points" json:"entryPoints,omitempty"`
		Routes      *[]struct {
			Match       *string `tfsdk:"match" json:"match,omitempty"`
			Middlewares *[]struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"middlewares" json:"middlewares,omitempty"`
			Priority *int64 `tfsdk:"priority" json:"priority,omitempty"`
			Services *[]struct {
				Name          *string `tfsdk:"name" json:"name,omitempty"`
				Namespace     *string `tfsdk:"namespace" json:"namespace,omitempty"`
				NativeLB      *bool   `tfsdk:"native_lb" json:"nativeLB,omitempty"`
				NodePortLB    *bool   `tfsdk:"node_port_lb" json:"nodePortLB,omitempty"`
				Port          *string `tfsdk:"port" json:"port,omitempty"`
				ProxyProtocol *struct {
					Version *int64 `tfsdk:"version" json:"version,omitempty"`
				} `tfsdk:"proxy_protocol" json:"proxyProtocol,omitempty"`
				ServersTransport *string `tfsdk:"servers_transport" json:"serversTransport,omitempty"`
				TerminationDelay *int64  `tfsdk:"termination_delay" json:"terminationDelay,omitempty"`
				Tls              *bool   `tfsdk:"tls" json:"tls,omitempty"`
				Weight           *int64  `tfsdk:"weight" json:"weight,omitempty"`
			} `tfsdk:"services" json:"services,omitempty"`
			Syntax *string `tfsdk:"syntax" json:"syntax,omitempty"`
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
			Passthrough *bool   `tfsdk:"passthrough" json:"passthrough,omitempty"`
			SecretName  *string `tfsdk:"secret_name" json:"secretName,omitempty"`
			Store       *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"store" json:"store,omitempty"`
		} `tfsdk:"tls" json:"tls,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *TraefikIoIngressRouteTcpV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_traefik_io_ingress_route_tcp_v1alpha1_manifest"
}

func (r *TraefikIoIngressRouteTcpV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "IngressRouteTCP is the CRD implementation of a Traefik TCP Router.",
		MarkdownDescription: "IngressRouteTCP is the CRD implementation of a Traefik TCP Router.",
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
				Description:         "IngressRouteTCPSpec defines the desired state of IngressRouteTCP.",
				MarkdownDescription: "IngressRouteTCPSpec defines the desired state of IngressRouteTCP.",
				Attributes: map[string]schema.Attribute{
					"entry_points": schema.ListAttribute{
						Description:         "EntryPoints defines the list of entry point names to bind to.Entry points have to be configured in the static configuration.More info: https://doc.traefik.io/traefik/v3.0/routing/entrypoints/Default: all.",
						MarkdownDescription: "EntryPoints defines the list of entry point names to bind to.Entry points have to be configured in the static configuration.More info: https://doc.traefik.io/traefik/v3.0/routing/entrypoints/Default: all.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"routes": schema.ListNestedAttribute{
						Description:         "Routes defines the list of routes.",
						MarkdownDescription: "Routes defines the list of routes.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"match": schema.StringAttribute{
									Description:         "Match defines the router's rule.More info: https://doc.traefik.io/traefik/v3.0/routing/routers/#rule_1",
									MarkdownDescription: "Match defines the router's rule.More info: https://doc.traefik.io/traefik/v3.0/routing/routers/#rule_1",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"middlewares": schema.ListNestedAttribute{
									Description:         "Middlewares defines the list of references to MiddlewareTCP resources.",
									MarkdownDescription: "Middlewares defines the list of references to MiddlewareTCP resources.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name defines the name of the referenced Traefik resource.",
												MarkdownDescription: "Name defines the name of the referenced Traefik resource.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "Namespace defines the namespace of the referenced Traefik resource.",
												MarkdownDescription: "Namespace defines the namespace of the referenced Traefik resource.",
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

								"priority": schema.Int64Attribute{
									Description:         "Priority defines the router's priority.More info: https://doc.traefik.io/traefik/v3.0/routing/routers/#priority_1",
									MarkdownDescription: "Priority defines the router's priority.More info: https://doc.traefik.io/traefik/v3.0/routing/routers/#priority_1",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"services": schema.ListNestedAttribute{
									Description:         "Services defines the list of TCP services.",
									MarkdownDescription: "Services defines the list of TCP services.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name defines the name of the referenced Kubernetes Service.",
												MarkdownDescription: "Name defines the name of the referenced Kubernetes Service.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "Namespace defines the namespace of the referenced Kubernetes Service.",
												MarkdownDescription: "Namespace defines the namespace of the referenced Kubernetes Service.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"native_lb": schema.BoolAttribute{
												Description:         "NativeLB controls, when creating the load-balancer,whether the LB's children are directly the pods IPs or if the only child is the Kubernetes Service clusterIP.The Kubernetes Service itself does load-balance to the pods.By default, NativeLB is false.",
												MarkdownDescription: "NativeLB controls, when creating the load-balancer,whether the LB's children are directly the pods IPs or if the only child is the Kubernetes Service clusterIP.The Kubernetes Service itself does load-balance to the pods.By default, NativeLB is false.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"node_port_lb": schema.BoolAttribute{
												Description:         "NodePortLB controls, when creating the load-balancer,whether the LB's children are directly the nodes internal IPs using the nodePort when the service type is NodePort.It allows services to be reachable when Traefik runs externally from the Kubernetes cluster but within the same network of the nodes.By default, NodePortLB is false.",
												MarkdownDescription: "NodePortLB controls, when creating the load-balancer,whether the LB's children are directly the nodes internal IPs using the nodePort when the service type is NodePort.It allows services to be reachable when Traefik runs externally from the Kubernetes cluster but within the same network of the nodes.By default, NodePortLB is false.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"port": schema.StringAttribute{
												Description:         "Port defines the port of a Kubernetes Service.This can be a reference to a named port.",
												MarkdownDescription: "Port defines the port of a Kubernetes Service.This can be a reference to a named port.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"proxy_protocol": schema.SingleNestedAttribute{
												Description:         "ProxyProtocol defines the PROXY protocol configuration.More info: https://doc.traefik.io/traefik/v3.0/routing/services/#proxy-protocol",
												MarkdownDescription: "ProxyProtocol defines the PROXY protocol configuration.More info: https://doc.traefik.io/traefik/v3.0/routing/services/#proxy-protocol",
												Attributes: map[string]schema.Attribute{
													"version": schema.Int64Attribute{
														Description:         "Version defines the PROXY Protocol version to use.",
														MarkdownDescription: "Version defines the PROXY Protocol version to use.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"servers_transport": schema.StringAttribute{
												Description:         "ServersTransport defines the name of ServersTransportTCP resource to use.It allows to configure the transport between Traefik and your servers.Can only be used on a Kubernetes Service.",
												MarkdownDescription: "ServersTransport defines the name of ServersTransportTCP resource to use.It allows to configure the transport between Traefik and your servers.Can only be used on a Kubernetes Service.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"termination_delay": schema.Int64Attribute{
												Description:         "TerminationDelay defines the deadline that the proxy sets, after one of its connected peers indicatesit has closed the writing capability of its connection, to close the reading capability as well,hence fully terminating the connection.It is a duration in milliseconds, defaulting to 100.A negative value means an infinite deadline (i.e. the reading capability is never closed).Deprecated: TerminationDelay is not supported APIVersion traefik.io/v1, please use ServersTransport to configure the TerminationDelay instead.",
												MarkdownDescription: "TerminationDelay defines the deadline that the proxy sets, after one of its connected peers indicatesit has closed the writing capability of its connection, to close the reading capability as well,hence fully terminating the connection.It is a duration in milliseconds, defaulting to 100.A negative value means an infinite deadline (i.e. the reading capability is never closed).Deprecated: TerminationDelay is not supported APIVersion traefik.io/v1, please use ServersTransport to configure the TerminationDelay instead.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"tls": schema.BoolAttribute{
												Description:         "TLS determines whether to use TLS when dialing with the backend.",
												MarkdownDescription: "TLS determines whether to use TLS when dialing with the backend.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"weight": schema.Int64Attribute{
												Description:         "Weight defines the weight used when balancing requests between multiple Kubernetes Service.",
												MarkdownDescription: "Weight defines the weight used when balancing requests between multiple Kubernetes Service.",
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

								"syntax": schema.StringAttribute{
									Description:         "Syntax defines the router's rule syntax.More info: https://doc.traefik.io/traefik/v3.0/routing/routers/#rulesyntax_1",
									MarkdownDescription: "Syntax defines the router's rule syntax.More info: https://doc.traefik.io/traefik/v3.0/routing/routers/#rulesyntax_1",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"tls": schema.SingleNestedAttribute{
						Description:         "TLS defines the TLS configuration on a layer 4 / TCP Route.More info: https://doc.traefik.io/traefik/v3.0/routing/routers/#tls_1",
						MarkdownDescription: "TLS defines the TLS configuration on a layer 4 / TCP Route.More info: https://doc.traefik.io/traefik/v3.0/routing/routers/#tls_1",
						Attributes: map[string]schema.Attribute{
							"cert_resolver": schema.StringAttribute{
								Description:         "CertResolver defines the name of the certificate resolver to use.Cert resolvers have to be configured in the static configuration.More info: https://doc.traefik.io/traefik/v3.0/https/acme/#certificate-resolvers",
								MarkdownDescription: "CertResolver defines the name of the certificate resolver to use.Cert resolvers have to be configured in the static configuration.More info: https://doc.traefik.io/traefik/v3.0/https/acme/#certificate-resolvers",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"domains": schema.ListNestedAttribute{
								Description:         "Domains defines the list of domains that will be used to issue certificates.More info: https://doc.traefik.io/traefik/v3.0/routing/routers/#domains",
								MarkdownDescription: "Domains defines the list of domains that will be used to issue certificates.More info: https://doc.traefik.io/traefik/v3.0/routing/routers/#domains",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"main": schema.StringAttribute{
											Description:         "Main defines the main domain name.",
											MarkdownDescription: "Main defines the main domain name.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"sans": schema.ListAttribute{
											Description:         "SANs defines the subject alternative domain names.",
											MarkdownDescription: "SANs defines the subject alternative domain names.",
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

							"options": schema.SingleNestedAttribute{
								Description:         "Options defines the reference to a TLSOption, that specifies the parameters of the TLS connection.If not defined, the 'default' TLSOption is used.More info: https://doc.traefik.io/traefik/v3.0/https/tls/#tls-options",
								MarkdownDescription: "Options defines the reference to a TLSOption, that specifies the parameters of the TLS connection.If not defined, the 'default' TLSOption is used.More info: https://doc.traefik.io/traefik/v3.0/https/tls/#tls-options",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name defines the name of the referenced Traefik resource.",
										MarkdownDescription: "Name defines the name of the referenced Traefik resource.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
										Description:         "Namespace defines the namespace of the referenced Traefik resource.",
										MarkdownDescription: "Namespace defines the namespace of the referenced Traefik resource.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"passthrough": schema.BoolAttribute{
								Description:         "Passthrough defines whether a TLS router will terminate the TLS connection.",
								MarkdownDescription: "Passthrough defines whether a TLS router will terminate the TLS connection.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secret_name": schema.StringAttribute{
								Description:         "SecretName is the name of the referenced Kubernetes Secret to specify the certificate details.",
								MarkdownDescription: "SecretName is the name of the referenced Kubernetes Secret to specify the certificate details.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"store": schema.SingleNestedAttribute{
								Description:         "Store defines the reference to the TLSStore, that will be used to store certificates.Please note that only 'default' TLSStore can be used.",
								MarkdownDescription: "Store defines the reference to the TLSStore, that will be used to store certificates.Please note that only 'default' TLSStore can be used.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name defines the name of the referenced Traefik resource.",
										MarkdownDescription: "Name defines the name of the referenced Traefik resource.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
										Description:         "Namespace defines the namespace of the referenced Traefik resource.",
										MarkdownDescription: "Namespace defines the namespace of the referenced Traefik resource.",
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
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *TraefikIoIngressRouteTcpV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_traefik_io_ingress_route_tcp_v1alpha1_manifest")

	var model TraefikIoIngressRouteTcpV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("traefik.io/v1alpha1")
	model.Kind = pointer.String("IngressRouteTCP")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
