/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"
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

type TraefikContainoUsTraefikServiceV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*TraefikContainoUsTraefikServiceV1Alpha1Resource)(nil)
)

type TraefikContainoUsTraefikServiceV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type TraefikContainoUsTraefikServiceV1Alpha1GoModel struct {
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
		Mirroring *struct {
			Port *string `tfsdk:"port" yaml:"port,omitempty"`

			Weight *int64 `tfsdk:"weight" yaml:"weight,omitempty"`

			MaxBodySize *int64 `tfsdk:"max_body_size" yaml:"maxBodySize,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

			PassHostHeader *bool `tfsdk:"pass_host_header" yaml:"passHostHeader,omitempty"`

			ResponseForwarding *struct {
				FlushInterval *string `tfsdk:"flush_interval" yaml:"flushInterval,omitempty"`
			} `tfsdk:"response_forwarding" yaml:"responseForwarding,omitempty"`

			Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`

			ServersTransport *string `tfsdk:"servers_transport" yaml:"serversTransport,omitempty"`

			Sticky *struct {
				Cookie *struct {
					HttpOnly *bool `tfsdk:"http_only" yaml:"httpOnly,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					SameSite *string `tfsdk:"same_site" yaml:"sameSite,omitempty"`

					Secure *bool `tfsdk:"secure" yaml:"secure,omitempty"`
				} `tfsdk:"cookie" yaml:"cookie,omitempty"`
			} `tfsdk:"sticky" yaml:"sticky,omitempty"`

			Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

			Mirrors *[]struct {
				Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`

				Sticky *struct {
					Cookie *struct {
						SameSite *string `tfsdk:"same_site" yaml:"sameSite,omitempty"`

						Secure *bool `tfsdk:"secure" yaml:"secure,omitempty"`

						HttpOnly *bool `tfsdk:"http_only" yaml:"httpOnly,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`
					} `tfsdk:"cookie" yaml:"cookie,omitempty"`
				} `tfsdk:"sticky" yaml:"sticky,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

				PassHostHeader *bool `tfsdk:"pass_host_header" yaml:"passHostHeader,omitempty"`

				Port *string `tfsdk:"port" yaml:"port,omitempty"`

				ResponseForwarding *struct {
					FlushInterval *string `tfsdk:"flush_interval" yaml:"flushInterval,omitempty"`
				} `tfsdk:"response_forwarding" yaml:"responseForwarding,omitempty"`

				ServersTransport *string `tfsdk:"servers_transport" yaml:"serversTransport,omitempty"`

				Strategy *string `tfsdk:"strategy" yaml:"strategy,omitempty"`

				Weight *int64 `tfsdk:"weight" yaml:"weight,omitempty"`

				Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Percent *int64 `tfsdk:"percent" yaml:"percent,omitempty"`
			} `tfsdk:"mirrors" yaml:"mirrors,omitempty"`

			Strategy *string `tfsdk:"strategy" yaml:"strategy,omitempty"`
		} `tfsdk:"mirroring" yaml:"mirroring,omitempty"`

		Weighted *struct {
			Services *[]struct {
				Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`

				Sticky *struct {
					Cookie *struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						SameSite *string `tfsdk:"same_site" yaml:"sameSite,omitempty"`

						Secure *bool `tfsdk:"secure" yaml:"secure,omitempty"`

						HttpOnly *bool `tfsdk:"http_only" yaml:"httpOnly,omitempty"`
					} `tfsdk:"cookie" yaml:"cookie,omitempty"`
				} `tfsdk:"sticky" yaml:"sticky,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

				ResponseForwarding *struct {
					FlushInterval *string `tfsdk:"flush_interval" yaml:"flushInterval,omitempty"`
				} `tfsdk:"response_forwarding" yaml:"responseForwarding,omitempty"`

				ServersTransport *string `tfsdk:"servers_transport" yaml:"serversTransport,omitempty"`

				Strategy *string `tfsdk:"strategy" yaml:"strategy,omitempty"`

				Weight *int64 `tfsdk:"weight" yaml:"weight,omitempty"`

				Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

				PassHostHeader *bool `tfsdk:"pass_host_header" yaml:"passHostHeader,omitempty"`

				Port *string `tfsdk:"port" yaml:"port,omitempty"`
			} `tfsdk:"services" yaml:"services,omitempty"`

			Sticky *struct {
				Cookie *struct {
					SameSite *string `tfsdk:"same_site" yaml:"sameSite,omitempty"`

					Secure *bool `tfsdk:"secure" yaml:"secure,omitempty"`

					HttpOnly *bool `tfsdk:"http_only" yaml:"httpOnly,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"cookie" yaml:"cookie,omitempty"`
			} `tfsdk:"sticky" yaml:"sticky,omitempty"`
		} `tfsdk:"weighted" yaml:"weighted,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewTraefikContainoUsTraefikServiceV1Alpha1Resource() resource.Resource {
	return &TraefikContainoUsTraefikServiceV1Alpha1Resource{}
}

func (r *TraefikContainoUsTraefikServiceV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_traefik_containo_us_traefik_service_v1alpha1"
}

func (r *TraefikContainoUsTraefikServiceV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "TraefikService is the CRD implementation of a Traefik Service. TraefikService object allows to: - Apply weight to Services on load-balancing - Mirror traffic on services More info: https://doc.traefik.io/traefik/v2.9/routing/providers/kubernetes-crd/#kind-traefikservice",
		MarkdownDescription: "TraefikService is the CRD implementation of a Traefik Service. TraefikService object allows to: - Apply weight to Services on load-balancing - Mirror traffic on services More info: https://doc.traefik.io/traefik/v2.9/routing/providers/kubernetes-crd/#kind-traefikservice",
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
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.RequiresReplace(),
						},
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
				Description:         "TraefikServiceSpec defines the desired state of a TraefikService.",
				MarkdownDescription: "TraefikServiceSpec defines the desired state of a TraefikService.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"mirroring": {
						Description:         "Mirroring defines the Mirroring service configuration.",
						MarkdownDescription: "Mirroring defines the Mirroring service configuration.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"port": {
								Description:         "Port defines the port of a Kubernetes Service. This can be a reference to a named port.",
								MarkdownDescription: "Port defines the port of a Kubernetes Service. This can be a reference to a named port.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"weight": {
								Description:         "Weight defines the weight and should only be specified when Name references a TraefikService object (and to be precise, one that embeds a Weighted Round Robin).",
								MarkdownDescription: "Weight defines the weight and should only be specified when Name references a TraefikService object (and to be precise, one that embeds a Weighted Round Robin).",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"max_body_size": {
								Description:         "MaxBodySize defines the maximum size allowed for the body of the request. If the body is larger, the request is not mirrored. Default value is -1, which means unlimited size.",
								MarkdownDescription: "MaxBodySize defines the maximum size allowed for the body of the request. If the body is larger, the request is not mirrored. Default value is -1, which means unlimited size.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "Name defines the name of the referenced Kubernetes Service or TraefikService. The differentiation between the two is specified in the Kind field.",
								MarkdownDescription: "Name defines the name of the referenced Kubernetes Service or TraefikService. The differentiation between the two is specified in the Kind field.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"namespace": {
								Description:         "Namespace defines the namespace of the referenced Kubernetes Service or TraefikService.",
								MarkdownDescription: "Namespace defines the namespace of the referenced Kubernetes Service or TraefikService.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pass_host_header": {
								Description:         "PassHostHeader defines whether the client Host header is forwarded to the upstream Kubernetes Service. By default, passHostHeader is true.",
								MarkdownDescription: "PassHostHeader defines whether the client Host header is forwarded to the upstream Kubernetes Service. By default, passHostHeader is true.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"response_forwarding": {
								Description:         "ResponseForwarding defines how Traefik forwards the response from the upstream Kubernetes Service to the client.",
								MarkdownDescription: "ResponseForwarding defines how Traefik forwards the response from the upstream Kubernetes Service to the client.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"flush_interval": {
										Description:         "FlushInterval defines the interval, in milliseconds, in between flushes to the client while copying the response body. A negative value means to flush immediately after each write to the client. This configuration is ignored when ReverseProxy recognizes a response as a streaming response; for such responses, writes are flushed to the client immediately. Default: 100ms",
										MarkdownDescription: "FlushInterval defines the interval, in milliseconds, in between flushes to the client while copying the response body. A negative value means to flush immediately after each write to the client. This configuration is ignored when ReverseProxy recognizes a response as a streaming response; for such responses, writes are flushed to the client immediately. Default: 100ms",

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

							"scheme": {
								Description:         "Scheme defines the scheme to use for the request to the upstream Kubernetes Service. It defaults to https when Kubernetes Service port is 443, http otherwise.",
								MarkdownDescription: "Scheme defines the scheme to use for the request to the upstream Kubernetes Service. It defaults to https when Kubernetes Service port is 443, http otherwise.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"servers_transport": {
								Description:         "ServersTransport defines the name of ServersTransport resource to use. It allows to configure the transport between Traefik and your servers. Can only be used on a Kubernetes Service.",
								MarkdownDescription: "ServersTransport defines the name of ServersTransport resource to use. It allows to configure the transport between Traefik and your servers. Can only be used on a Kubernetes Service.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"sticky": {
								Description:         "Sticky defines the sticky sessions configuration. More info: https://doc.traefik.io/traefik/v2.9/routing/services/#sticky-sessions",
								MarkdownDescription: "Sticky defines the sticky sessions configuration. More info: https://doc.traefik.io/traefik/v2.9/routing/services/#sticky-sessions",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"cookie": {
										Description:         "Cookie defines the sticky cookie configuration.",
										MarkdownDescription: "Cookie defines the sticky cookie configuration.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"http_only": {
												Description:         "HTTPOnly defines whether the cookie can be accessed by client-side APIs, such as JavaScript.",
												MarkdownDescription: "HTTPOnly defines whether the cookie can be accessed by client-side APIs, such as JavaScript.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "Name defines the Cookie name.",
												MarkdownDescription: "Name defines the Cookie name.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"same_site": {
												Description:         "SameSite defines the same site policy. More info: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie/SameSite",
												MarkdownDescription: "SameSite defines the same site policy. More info: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie/SameSite",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secure": {
												Description:         "Secure defines whether the cookie can only be transmitted over an encrypted connection (i.e. HTTPS).",
												MarkdownDescription: "Secure defines whether the cookie can only be transmitted over an encrypted connection (i.e. HTTPS).",

												Type: types.BoolType,

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

							"kind": {
								Description:         "Kind defines the kind of the Service.",
								MarkdownDescription: "Kind defines the kind of the Service.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"mirrors": {
								Description:         "Mirrors defines the list of mirrors where Traefik will duplicate the traffic.",
								MarkdownDescription: "Mirrors defines the list of mirrors where Traefik will duplicate the traffic.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"scheme": {
										Description:         "Scheme defines the scheme to use for the request to the upstream Kubernetes Service. It defaults to https when Kubernetes Service port is 443, http otherwise.",
										MarkdownDescription: "Scheme defines the scheme to use for the request to the upstream Kubernetes Service. It defaults to https when Kubernetes Service port is 443, http otherwise.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"sticky": {
										Description:         "Sticky defines the sticky sessions configuration. More info: https://doc.traefik.io/traefik/v2.9/routing/services/#sticky-sessions",
										MarkdownDescription: "Sticky defines the sticky sessions configuration. More info: https://doc.traefik.io/traefik/v2.9/routing/services/#sticky-sessions",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"cookie": {
												Description:         "Cookie defines the sticky cookie configuration.",
												MarkdownDescription: "Cookie defines the sticky cookie configuration.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"same_site": {
														Description:         "SameSite defines the same site policy. More info: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie/SameSite",
														MarkdownDescription: "SameSite defines the same site policy. More info: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie/SameSite",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"secure": {
														Description:         "Secure defines whether the cookie can only be transmitted over an encrypted connection (i.e. HTTPS).",
														MarkdownDescription: "Secure defines whether the cookie can only be transmitted over an encrypted connection (i.e. HTTPS).",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"http_only": {
														Description:         "HTTPOnly defines whether the cookie can be accessed by client-side APIs, such as JavaScript.",
														MarkdownDescription: "HTTPOnly defines whether the cookie can be accessed by client-side APIs, such as JavaScript.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": {
														Description:         "Name defines the Cookie name.",
														MarkdownDescription: "Name defines the Cookie name.",

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

									"namespace": {
										Description:         "Namespace defines the namespace of the referenced Kubernetes Service or TraefikService.",
										MarkdownDescription: "Namespace defines the namespace of the referenced Kubernetes Service or TraefikService.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"pass_host_header": {
										Description:         "PassHostHeader defines whether the client Host header is forwarded to the upstream Kubernetes Service. By default, passHostHeader is true.",
										MarkdownDescription: "PassHostHeader defines whether the client Host header is forwarded to the upstream Kubernetes Service. By default, passHostHeader is true.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"port": {
										Description:         "Port defines the port of a Kubernetes Service. This can be a reference to a named port.",
										MarkdownDescription: "Port defines the port of a Kubernetes Service. This can be a reference to a named port.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"response_forwarding": {
										Description:         "ResponseForwarding defines how Traefik forwards the response from the upstream Kubernetes Service to the client.",
										MarkdownDescription: "ResponseForwarding defines how Traefik forwards the response from the upstream Kubernetes Service to the client.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"flush_interval": {
												Description:         "FlushInterval defines the interval, in milliseconds, in between flushes to the client while copying the response body. A negative value means to flush immediately after each write to the client. This configuration is ignored when ReverseProxy recognizes a response as a streaming response; for such responses, writes are flushed to the client immediately. Default: 100ms",
												MarkdownDescription: "FlushInterval defines the interval, in milliseconds, in between flushes to the client while copying the response body. A negative value means to flush immediately after each write to the client. This configuration is ignored when ReverseProxy recognizes a response as a streaming response; for such responses, writes are flushed to the client immediately. Default: 100ms",

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

									"servers_transport": {
										Description:         "ServersTransport defines the name of ServersTransport resource to use. It allows to configure the transport between Traefik and your servers. Can only be used on a Kubernetes Service.",
										MarkdownDescription: "ServersTransport defines the name of ServersTransport resource to use. It allows to configure the transport between Traefik and your servers. Can only be used on a Kubernetes Service.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"strategy": {
										Description:         "Strategy defines the load balancing strategy between the servers. RoundRobin is the only supported value at the moment.",
										MarkdownDescription: "Strategy defines the load balancing strategy between the servers. RoundRobin is the only supported value at the moment.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"weight": {
										Description:         "Weight defines the weight and should only be specified when Name references a TraefikService object (and to be precise, one that embeds a Weighted Round Robin).",
										MarkdownDescription: "Weight defines the weight and should only be specified when Name references a TraefikService object (and to be precise, one that embeds a Weighted Round Robin).",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"kind": {
										Description:         "Kind defines the kind of the Service.",
										MarkdownDescription: "Kind defines the kind of the Service.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "Name defines the name of the referenced Kubernetes Service or TraefikService. The differentiation between the two is specified in the Kind field.",
										MarkdownDescription: "Name defines the name of the referenced Kubernetes Service or TraefikService. The differentiation between the two is specified in the Kind field.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"percent": {
										Description:         "Percent defines the part of the traffic to mirror. Supported values: 0 to 100.",
										MarkdownDescription: "Percent defines the part of the traffic to mirror. Supported values: 0 to 100.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"strategy": {
								Description:         "Strategy defines the load balancing strategy between the servers. RoundRobin is the only supported value at the moment.",
								MarkdownDescription: "Strategy defines the load balancing strategy between the servers. RoundRobin is the only supported value at the moment.",

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

					"weighted": {
						Description:         "Weighted defines the Weighted Round Robin configuration.",
						MarkdownDescription: "Weighted defines the Weighted Round Robin configuration.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"services": {
								Description:         "Services defines the list of Kubernetes Service and/or TraefikService to load-balance, with weight.",
								MarkdownDescription: "Services defines the list of Kubernetes Service and/or TraefikService to load-balance, with weight.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"scheme": {
										Description:         "Scheme defines the scheme to use for the request to the upstream Kubernetes Service. It defaults to https when Kubernetes Service port is 443, http otherwise.",
										MarkdownDescription: "Scheme defines the scheme to use for the request to the upstream Kubernetes Service. It defaults to https when Kubernetes Service port is 443, http otherwise.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"sticky": {
										Description:         "Sticky defines the sticky sessions configuration. More info: https://doc.traefik.io/traefik/v2.9/routing/services/#sticky-sessions",
										MarkdownDescription: "Sticky defines the sticky sessions configuration. More info: https://doc.traefik.io/traefik/v2.9/routing/services/#sticky-sessions",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"cookie": {
												Description:         "Cookie defines the sticky cookie configuration.",
												MarkdownDescription: "Cookie defines the sticky cookie configuration.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "Name defines the Cookie name.",
														MarkdownDescription: "Name defines the Cookie name.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"same_site": {
														Description:         "SameSite defines the same site policy. More info: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie/SameSite",
														MarkdownDescription: "SameSite defines the same site policy. More info: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie/SameSite",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"secure": {
														Description:         "Secure defines whether the cookie can only be transmitted over an encrypted connection (i.e. HTTPS).",
														MarkdownDescription: "Secure defines whether the cookie can only be transmitted over an encrypted connection (i.e. HTTPS).",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"http_only": {
														Description:         "HTTPOnly defines whether the cookie can be accessed by client-side APIs, such as JavaScript.",
														MarkdownDescription: "HTTPOnly defines whether the cookie can be accessed by client-side APIs, such as JavaScript.",

														Type: types.BoolType,

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

									"name": {
										Description:         "Name defines the name of the referenced Kubernetes Service or TraefikService. The differentiation between the two is specified in the Kind field.",
										MarkdownDescription: "Name defines the name of the referenced Kubernetes Service or TraefikService. The differentiation between the two is specified in the Kind field.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"namespace": {
										Description:         "Namespace defines the namespace of the referenced Kubernetes Service or TraefikService.",
										MarkdownDescription: "Namespace defines the namespace of the referenced Kubernetes Service or TraefikService.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"response_forwarding": {
										Description:         "ResponseForwarding defines how Traefik forwards the response from the upstream Kubernetes Service to the client.",
										MarkdownDescription: "ResponseForwarding defines how Traefik forwards the response from the upstream Kubernetes Service to the client.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"flush_interval": {
												Description:         "FlushInterval defines the interval, in milliseconds, in between flushes to the client while copying the response body. A negative value means to flush immediately after each write to the client. This configuration is ignored when ReverseProxy recognizes a response as a streaming response; for such responses, writes are flushed to the client immediately. Default: 100ms",
												MarkdownDescription: "FlushInterval defines the interval, in milliseconds, in between flushes to the client while copying the response body. A negative value means to flush immediately after each write to the client. This configuration is ignored when ReverseProxy recognizes a response as a streaming response; for such responses, writes are flushed to the client immediately. Default: 100ms",

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

									"servers_transport": {
										Description:         "ServersTransport defines the name of ServersTransport resource to use. It allows to configure the transport between Traefik and your servers. Can only be used on a Kubernetes Service.",
										MarkdownDescription: "ServersTransport defines the name of ServersTransport resource to use. It allows to configure the transport between Traefik and your servers. Can only be used on a Kubernetes Service.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"strategy": {
										Description:         "Strategy defines the load balancing strategy between the servers. RoundRobin is the only supported value at the moment.",
										MarkdownDescription: "Strategy defines the load balancing strategy between the servers. RoundRobin is the only supported value at the moment.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"weight": {
										Description:         "Weight defines the weight and should only be specified when Name references a TraefikService object (and to be precise, one that embeds a Weighted Round Robin).",
										MarkdownDescription: "Weight defines the weight and should only be specified when Name references a TraefikService object (and to be precise, one that embeds a Weighted Round Robin).",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"kind": {
										Description:         "Kind defines the kind of the Service.",
										MarkdownDescription: "Kind defines the kind of the Service.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"pass_host_header": {
										Description:         "PassHostHeader defines whether the client Host header is forwarded to the upstream Kubernetes Service. By default, passHostHeader is true.",
										MarkdownDescription: "PassHostHeader defines whether the client Host header is forwarded to the upstream Kubernetes Service. By default, passHostHeader is true.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"port": {
										Description:         "Port defines the port of a Kubernetes Service. This can be a reference to a named port.",
										MarkdownDescription: "Port defines the port of a Kubernetes Service. This can be a reference to a named port.",

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

							"sticky": {
								Description:         "Sticky defines whether sticky sessions are enabled. More info: https://doc.traefik.io/traefik/v2.9/routing/providers/kubernetes-crd/#stickiness-and-load-balancing",
								MarkdownDescription: "Sticky defines whether sticky sessions are enabled. More info: https://doc.traefik.io/traefik/v2.9/routing/providers/kubernetes-crd/#stickiness-and-load-balancing",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"cookie": {
										Description:         "Cookie defines the sticky cookie configuration.",
										MarkdownDescription: "Cookie defines the sticky cookie configuration.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"same_site": {
												Description:         "SameSite defines the same site policy. More info: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie/SameSite",
												MarkdownDescription: "SameSite defines the same site policy. More info: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie/SameSite",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secure": {
												Description:         "Secure defines whether the cookie can only be transmitted over an encrypted connection (i.e. HTTPS).",
												MarkdownDescription: "Secure defines whether the cookie can only be transmitted over an encrypted connection (i.e. HTTPS).",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"http_only": {
												Description:         "HTTPOnly defines whether the cookie can be accessed by client-side APIs, such as JavaScript.",
												MarkdownDescription: "HTTPOnly defines whether the cookie can be accessed by client-side APIs, such as JavaScript.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "Name defines the Cookie name.",
												MarkdownDescription: "Name defines the Cookie name.",

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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}, nil
}

func (r *TraefikContainoUsTraefikServiceV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_traefik_containo_us_traefik_service_v1alpha1")

	var state TraefikContainoUsTraefikServiceV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel TraefikContainoUsTraefikServiceV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("traefik.containo.us/v1alpha1")
	goModel.Kind = utilities.Ptr("TraefikService")

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

func (r *TraefikContainoUsTraefikServiceV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_traefik_containo_us_traefik_service_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *TraefikContainoUsTraefikServiceV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_traefik_containo_us_traefik_service_v1alpha1")

	var state TraefikContainoUsTraefikServiceV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel TraefikContainoUsTraefikServiceV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("traefik.containo.us/v1alpha1")
	goModel.Kind = utilities.Ptr("TraefikService")

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

func (r *TraefikContainoUsTraefikServiceV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_traefik_containo_us_traefik_service_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
