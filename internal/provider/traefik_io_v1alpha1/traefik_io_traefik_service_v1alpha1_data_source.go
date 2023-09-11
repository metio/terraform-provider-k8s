/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package traefik_io_v1alpha1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"net/http"
)

var (
	_ datasource.DataSource              = &TraefikIoTraefikServiceV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &TraefikIoTraefikServiceV1Alpha1DataSource{}
)

func NewTraefikIoTraefikServiceV1Alpha1DataSource() datasource.DataSource {
	return &TraefikIoTraefikServiceV1Alpha1DataSource{}
}

type TraefikIoTraefikServiceV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type TraefikIoTraefikServiceV1Alpha1DataSourceData struct {
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
		Mirroring *struct {
			Kind        *string `tfsdk:"kind" json:"kind,omitempty"`
			MaxBodySize *int64  `tfsdk:"max_body_size" json:"maxBodySize,omitempty"`
			Mirrors     *[]struct {
				Kind               *string `tfsdk:"kind" json:"kind,omitempty"`
				Name               *string `tfsdk:"name" json:"name,omitempty"`
				Namespace          *string `tfsdk:"namespace" json:"namespace,omitempty"`
				NativeLB           *bool   `tfsdk:"native_lb" json:"nativeLB,omitempty"`
				PassHostHeader     *bool   `tfsdk:"pass_host_header" json:"passHostHeader,omitempty"`
				Percent            *int64  `tfsdk:"percent" json:"percent,omitempty"`
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
			} `tfsdk:"mirrors" json:"mirrors,omitempty"`
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
		} `tfsdk:"mirroring" json:"mirroring,omitempty"`
		Weighted *struct {
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
			Sticky *struct {
				Cookie *struct {
					HttpOnly *bool   `tfsdk:"http_only" json:"httpOnly,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					SameSite *string `tfsdk:"same_site" json:"sameSite,omitempty"`
					Secure   *bool   `tfsdk:"secure" json:"secure,omitempty"`
				} `tfsdk:"cookie" json:"cookie,omitempty"`
			} `tfsdk:"sticky" json:"sticky,omitempty"`
		} `tfsdk:"weighted" json:"weighted,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *TraefikIoTraefikServiceV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_traefik_io_traefik_service_v1alpha1"
}

func (r *TraefikIoTraefikServiceV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "TraefikService is the CRD implementation of a Traefik Service. TraefikService object allows to: - Apply weight to Services on load-balancing - Mirror traffic on services More info: https://doc.traefik.io/traefik/v3.0/routing/providers/kubernetes-crd/#kind-traefikservice",
		MarkdownDescription: "TraefikService is the CRD implementation of a Traefik Service. TraefikService object allows to: - Apply weight to Services on load-balancing - Mirror traffic on services More info: https://doc.traefik.io/traefik/v3.0/routing/providers/kubernetes-crd/#kind-traefikservice",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
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
				Description:         "TraefikServiceSpec defines the desired state of a TraefikService.",
				MarkdownDescription: "TraefikServiceSpec defines the desired state of a TraefikService.",
				Attributes: map[string]schema.Attribute{
					"mirroring": schema.SingleNestedAttribute{
						Description:         "Mirroring defines the Mirroring service configuration.",
						MarkdownDescription: "Mirroring defines the Mirroring service configuration.",
						Attributes: map[string]schema.Attribute{
							"kind": schema.StringAttribute{
								Description:         "Kind defines the kind of the Service.",
								MarkdownDescription: "Kind defines the kind of the Service.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"max_body_size": schema.Int64Attribute{
								Description:         "MaxBodySize defines the maximum size allowed for the body of the request. If the body is larger, the request is not mirrored. Default value is -1, which means unlimited size.",
								MarkdownDescription: "MaxBodySize defines the maximum size allowed for the body of the request. If the body is larger, the request is not mirrored. Default value is -1, which means unlimited size.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"mirrors": schema.ListNestedAttribute{
								Description:         "Mirrors defines the list of mirrors where Traefik will duplicate the traffic.",
								MarkdownDescription: "Mirrors defines the list of mirrors where Traefik will duplicate the traffic.",
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

										"percent": schema.Int64Attribute{
											Description:         "Percent defines the part of the traffic to mirror. Supported values: 0 to 100.",
											MarkdownDescription: "Percent defines the part of the traffic to mirror. Supported values: 0 to 100.",
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
						Required: false,
						Optional: false,
						Computed: true,
					},

					"weighted": schema.SingleNestedAttribute{
						Description:         "Weighted defines the Weighted Round Robin configuration.",
						MarkdownDescription: "Weighted defines the Weighted Round Robin configuration.",
						Attributes: map[string]schema.Attribute{
							"services": schema.ListNestedAttribute{
								Description:         "Services defines the list of Kubernetes Service and/or TraefikService to load-balance, with weight.",
								MarkdownDescription: "Services defines the list of Kubernetes Service and/or TraefikService to load-balance, with weight.",
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

							"sticky": schema.SingleNestedAttribute{
								Description:         "Sticky defines whether sticky sessions are enabled. More info: https://doc.traefik.io/traefik/v3.0/routing/providers/kubernetes-crd/#stickiness-and-load-balancing",
								MarkdownDescription: "Sticky defines whether sticky sessions are enabled. More info: https://doc.traefik.io/traefik/v3.0/routing/providers/kubernetes-crd/#stickiness-and-load-balancing",
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

func (r *TraefikIoTraefikServiceV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *TraefikIoTraefikServiceV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_traefik_io_traefik_service_v1alpha1")

	var data TraefikIoTraefikServiceV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "traefik.io", Version: "v1alpha1", Resource: "traefikservices"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		var statusError *k8sErrors.StatusError
		if errors.As(err, &statusError) {
			if statusError.Status().Code == http.StatusNotFound {
				response.Diagnostics.AddError(
					"Unable to find resource",
					fmt.Sprintf("The requested resource cannot be found. "+
						"Make sure that it does exist in your cluster and you have set the correct name and namespace configured.\n\n"+
						"Namespace: %s\n"+
						"Name: %s", data.Metadata.Namespace, data.Metadata.Name),
				)
				return
			}
		} else {
			response.Diagnostics.AddError(
				"Unable to GET resource",
				fmt.Sprintf("An unexpected error occurred while reading the resource. "+
					"Please report this issue to the provider developers.\n\n"+
					"GET Error (%T): %s", err, err.Error()),
			)
		}
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse TraefikIoTraefikServiceV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("traefik.io/v1alpha1")
	data.Kind = pointer.String("TraefikService")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
