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
	_ datasource.DataSource = &TraefikIoTraefikServiceV1Alpha1Manifest{}
)

func NewTraefikIoTraefikServiceV1Alpha1Manifest() datasource.DataSource {
	return &TraefikIoTraefikServiceV1Alpha1Manifest{}
}

type TraefikIoTraefikServiceV1Alpha1Manifest struct{}

type TraefikIoTraefikServiceV1Alpha1ManifestData struct {
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
		Mirroring *struct {
			HealthCheck *struct {
				FollowRedirects *bool              `tfsdk:"follow_redirects" json:"followRedirects,omitempty"`
				Headers         *map[string]string `tfsdk:"headers" json:"headers,omitempty"`
				Hostname        *string            `tfsdk:"hostname" json:"hostname,omitempty"`
				Interval        *string            `tfsdk:"interval" json:"interval,omitempty"`
				Method          *string            `tfsdk:"method" json:"method,omitempty"`
				Mode            *string            `tfsdk:"mode" json:"mode,omitempty"`
				Path            *string            `tfsdk:"path" json:"path,omitempty"`
				Port            *int64             `tfsdk:"port" json:"port,omitempty"`
				Scheme          *string            `tfsdk:"scheme" json:"scheme,omitempty"`
				Status          *int64             `tfsdk:"status" json:"status,omitempty"`
				Timeout         *string            `tfsdk:"timeout" json:"timeout,omitempty"`
			} `tfsdk:"health_check" json:"healthCheck,omitempty"`
			Kind        *string `tfsdk:"kind" json:"kind,omitempty"`
			MaxBodySize *int64  `tfsdk:"max_body_size" json:"maxBodySize,omitempty"`
			MirrorBody  *bool   `tfsdk:"mirror_body" json:"mirrorBody,omitempty"`
			Mirrors     *[]struct {
				HealthCheck *struct {
					FollowRedirects *bool              `tfsdk:"follow_redirects" json:"followRedirects,omitempty"`
					Headers         *map[string]string `tfsdk:"headers" json:"headers,omitempty"`
					Hostname        *string            `tfsdk:"hostname" json:"hostname,omitempty"`
					Interval        *string            `tfsdk:"interval" json:"interval,omitempty"`
					Method          *string            `tfsdk:"method" json:"method,omitempty"`
					Mode            *string            `tfsdk:"mode" json:"mode,omitempty"`
					Path            *string            `tfsdk:"path" json:"path,omitempty"`
					Port            *int64             `tfsdk:"port" json:"port,omitempty"`
					Scheme          *string            `tfsdk:"scheme" json:"scheme,omitempty"`
					Status          *int64             `tfsdk:"status" json:"status,omitempty"`
					Timeout         *string            `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"health_check" json:"healthCheck,omitempty"`
				Kind               *string `tfsdk:"kind" json:"kind,omitempty"`
				Name               *string `tfsdk:"name" json:"name,omitempty"`
				Namespace          *string `tfsdk:"namespace" json:"namespace,omitempty"`
				NativeLB           *bool   `tfsdk:"native_lb" json:"nativeLB,omitempty"`
				NodePortLB         *bool   `tfsdk:"node_port_lb" json:"nodePortLB,omitempty"`
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
						MaxAge   *int64  `tfsdk:"max_age" json:"maxAge,omitempty"`
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
			NodePortLB         *bool   `tfsdk:"node_port_lb" json:"nodePortLB,omitempty"`
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
					MaxAge   *int64  `tfsdk:"max_age" json:"maxAge,omitempty"`
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
				HealthCheck *struct {
					FollowRedirects *bool              `tfsdk:"follow_redirects" json:"followRedirects,omitempty"`
					Headers         *map[string]string `tfsdk:"headers" json:"headers,omitempty"`
					Hostname        *string            `tfsdk:"hostname" json:"hostname,omitempty"`
					Interval        *string            `tfsdk:"interval" json:"interval,omitempty"`
					Method          *string            `tfsdk:"method" json:"method,omitempty"`
					Mode            *string            `tfsdk:"mode" json:"mode,omitempty"`
					Path            *string            `tfsdk:"path" json:"path,omitempty"`
					Port            *int64             `tfsdk:"port" json:"port,omitempty"`
					Scheme          *string            `tfsdk:"scheme" json:"scheme,omitempty"`
					Status          *int64             `tfsdk:"status" json:"status,omitempty"`
					Timeout         *string            `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"health_check" json:"healthCheck,omitempty"`
				Kind               *string `tfsdk:"kind" json:"kind,omitempty"`
				Name               *string `tfsdk:"name" json:"name,omitempty"`
				Namespace          *string `tfsdk:"namespace" json:"namespace,omitempty"`
				NativeLB           *bool   `tfsdk:"native_lb" json:"nativeLB,omitempty"`
				NodePortLB         *bool   `tfsdk:"node_port_lb" json:"nodePortLB,omitempty"`
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
						MaxAge   *int64  `tfsdk:"max_age" json:"maxAge,omitempty"`
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
					MaxAge   *int64  `tfsdk:"max_age" json:"maxAge,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					SameSite *string `tfsdk:"same_site" json:"sameSite,omitempty"`
					Secure   *bool   `tfsdk:"secure" json:"secure,omitempty"`
				} `tfsdk:"cookie" json:"cookie,omitempty"`
			} `tfsdk:"sticky" json:"sticky,omitempty"`
		} `tfsdk:"weighted" json:"weighted,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *TraefikIoTraefikServiceV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_traefik_io_traefik_service_v1alpha1_manifest"
}

func (r *TraefikIoTraefikServiceV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "TraefikService is the CRD implementation of a Traefik Service. TraefikService object allows to: - Apply weight to Services on load-balancing - Mirror traffic on services More info: https://doc.traefik.io/traefik/v3.2/routing/providers/kubernetes-crd/#kind-traefikservice",
		MarkdownDescription: "TraefikService is the CRD implementation of a Traefik Service. TraefikService object allows to: - Apply weight to Services on load-balancing - Mirror traffic on services More info: https://doc.traefik.io/traefik/v3.2/routing/providers/kubernetes-crd/#kind-traefikservice",
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
				Description:         "TraefikServiceSpec defines the desired state of a TraefikService.",
				MarkdownDescription: "TraefikServiceSpec defines the desired state of a TraefikService.",
				Attributes: map[string]schema.Attribute{
					"mirroring": schema.SingleNestedAttribute{
						Description:         "Mirroring defines the Mirroring service configuration.",
						MarkdownDescription: "Mirroring defines the Mirroring service configuration.",
						Attributes: map[string]schema.Attribute{
							"health_check": schema.SingleNestedAttribute{
								Description:         "Healthcheck defines health checks for ExternalName services.",
								MarkdownDescription: "Healthcheck defines health checks for ExternalName services.",
								Attributes: map[string]schema.Attribute{
									"follow_redirects": schema.BoolAttribute{
										Description:         "FollowRedirects defines whether redirects should be followed during the health check calls. Default: true",
										MarkdownDescription: "FollowRedirects defines whether redirects should be followed during the health check calls. Default: true",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"headers": schema.MapAttribute{
										Description:         "Headers defines custom headers to be sent to the health check endpoint.",
										MarkdownDescription: "Headers defines custom headers to be sent to the health check endpoint.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"hostname": schema.StringAttribute{
										Description:         "Hostname defines the value of hostname in the Host header of the health check request.",
										MarkdownDescription: "Hostname defines the value of hostname in the Host header of the health check request.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"interval": schema.StringAttribute{
										Description:         "Interval defines the frequency of the health check calls. Default: 30s",
										MarkdownDescription: "Interval defines the frequency of the health check calls. Default: 30s",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"method": schema.StringAttribute{
										Description:         "Method defines the healthcheck method.",
										MarkdownDescription: "Method defines the healthcheck method.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"mode": schema.StringAttribute{
										Description:         "Mode defines the health check mode. If defined to grpc, will use the gRPC health check protocol to probe the server. Default: http",
										MarkdownDescription: "Mode defines the health check mode. If defined to grpc, will use the gRPC health check protocol to probe the server. Default: http",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"path": schema.StringAttribute{
										Description:         "Path defines the server URL path for the health check endpoint.",
										MarkdownDescription: "Path defines the server URL path for the health check endpoint.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"port": schema.Int64Attribute{
										Description:         "Port defines the server URL port for the health check endpoint.",
										MarkdownDescription: "Port defines the server URL port for the health check endpoint.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"scheme": schema.StringAttribute{
										Description:         "Scheme replaces the server URL scheme for the health check endpoint.",
										MarkdownDescription: "Scheme replaces the server URL scheme for the health check endpoint.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"status": schema.Int64Attribute{
										Description:         "Status defines the expected HTTP status code of the response to the health check request.",
										MarkdownDescription: "Status defines the expected HTTP status code of the response to the health check request.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"timeout": schema.StringAttribute{
										Description:         "Timeout defines the maximum duration Traefik will wait for a health check request before considering the server unhealthy. Default: 5s",
										MarkdownDescription: "Timeout defines the maximum duration Traefik will wait for a health check request before considering the server unhealthy. Default: 5s",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind defines the kind of the Service.",
								MarkdownDescription: "Kind defines the kind of the Service.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Service", "TraefikService"),
								},
							},

							"max_body_size": schema.Int64Attribute{
								Description:         "MaxBodySize defines the maximum size allowed for the body of the request. If the body is larger, the request is not mirrored. Default value is -1, which means unlimited size.",
								MarkdownDescription: "MaxBodySize defines the maximum size allowed for the body of the request. If the body is larger, the request is not mirrored. Default value is -1, which means unlimited size.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"mirror_body": schema.BoolAttribute{
								Description:         "MirrorBody defines whether the body of the request should be mirrored. Default value is true.",
								MarkdownDescription: "MirrorBody defines whether the body of the request should be mirrored. Default value is true.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"mirrors": schema.ListNestedAttribute{
								Description:         "Mirrors defines the list of mirrors where Traefik will duplicate the traffic.",
								MarkdownDescription: "Mirrors defines the list of mirrors where Traefik will duplicate the traffic.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"health_check": schema.SingleNestedAttribute{
											Description:         "Healthcheck defines health checks for ExternalName services.",
											MarkdownDescription: "Healthcheck defines health checks for ExternalName services.",
											Attributes: map[string]schema.Attribute{
												"follow_redirects": schema.BoolAttribute{
													Description:         "FollowRedirects defines whether redirects should be followed during the health check calls. Default: true",
													MarkdownDescription: "FollowRedirects defines whether redirects should be followed during the health check calls. Default: true",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"headers": schema.MapAttribute{
													Description:         "Headers defines custom headers to be sent to the health check endpoint.",
													MarkdownDescription: "Headers defines custom headers to be sent to the health check endpoint.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"hostname": schema.StringAttribute{
													Description:         "Hostname defines the value of hostname in the Host header of the health check request.",
													MarkdownDescription: "Hostname defines the value of hostname in the Host header of the health check request.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"interval": schema.StringAttribute{
													Description:         "Interval defines the frequency of the health check calls. Default: 30s",
													MarkdownDescription: "Interval defines the frequency of the health check calls. Default: 30s",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"method": schema.StringAttribute{
													Description:         "Method defines the healthcheck method.",
													MarkdownDescription: "Method defines the healthcheck method.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"mode": schema.StringAttribute{
													Description:         "Mode defines the health check mode. If defined to grpc, will use the gRPC health check protocol to probe the server. Default: http",
													MarkdownDescription: "Mode defines the health check mode. If defined to grpc, will use the gRPC health check protocol to probe the server. Default: http",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"path": schema.StringAttribute{
													Description:         "Path defines the server URL path for the health check endpoint.",
													MarkdownDescription: "Path defines the server URL path for the health check endpoint.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"port": schema.Int64Attribute{
													Description:         "Port defines the server URL port for the health check endpoint.",
													MarkdownDescription: "Port defines the server URL port for the health check endpoint.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"scheme": schema.StringAttribute{
													Description:         "Scheme replaces the server URL scheme for the health check endpoint.",
													MarkdownDescription: "Scheme replaces the server URL scheme for the health check endpoint.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"status": schema.Int64Attribute{
													Description:         "Status defines the expected HTTP status code of the response to the health check request.",
													MarkdownDescription: "Status defines the expected HTTP status code of the response to the health check request.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"timeout": schema.StringAttribute{
													Description:         "Timeout defines the maximum duration Traefik will wait for a health check request before considering the server unhealthy. Default: 5s",
													MarkdownDescription: "Timeout defines the maximum duration Traefik will wait for a health check request before considering the server unhealthy. Default: 5s",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"kind": schema.StringAttribute{
											Description:         "Kind defines the kind of the Service.",
											MarkdownDescription: "Kind defines the kind of the Service.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Service", "TraefikService"),
											},
										},

										"name": schema.StringAttribute{
											Description:         "Name defines the name of the referenced Kubernetes Service or TraefikService. The differentiation between the two is specified in the Kind field.",
											MarkdownDescription: "Name defines the name of the referenced Kubernetes Service or TraefikService. The differentiation between the two is specified in the Kind field.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"namespace": schema.StringAttribute{
											Description:         "Namespace defines the namespace of the referenced Kubernetes Service or TraefikService.",
											MarkdownDescription: "Namespace defines the namespace of the referenced Kubernetes Service or TraefikService.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"native_lb": schema.BoolAttribute{
											Description:         "NativeLB controls, when creating the load-balancer, whether the LB's children are directly the pods IPs or if the only child is the Kubernetes Service clusterIP. The Kubernetes Service itself does load-balance to the pods. By default, NativeLB is false.",
											MarkdownDescription: "NativeLB controls, when creating the load-balancer, whether the LB's children are directly the pods IPs or if the only child is the Kubernetes Service clusterIP. The Kubernetes Service itself does load-balance to the pods. By default, NativeLB is false.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"node_port_lb": schema.BoolAttribute{
											Description:         "NodePortLB controls, when creating the load-balancer, whether the LB's children are directly the nodes internal IPs using the nodePort when the service type is NodePort. It allows services to be reachable when Traefik runs externally from the Kubernetes cluster but within the same network of the nodes. By default, NodePortLB is false.",
											MarkdownDescription: "NodePortLB controls, when creating the load-balancer, whether the LB's children are directly the nodes internal IPs using the nodePort when the service type is NodePort. It allows services to be reachable when Traefik runs externally from the Kubernetes cluster but within the same network of the nodes. By default, NodePortLB is false.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"pass_host_header": schema.BoolAttribute{
											Description:         "PassHostHeader defines whether the client Host header is forwarded to the upstream Kubernetes Service. By default, passHostHeader is true.",
											MarkdownDescription: "PassHostHeader defines whether the client Host header is forwarded to the upstream Kubernetes Service. By default, passHostHeader is true.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"percent": schema.Int64Attribute{
											Description:         "Percent defines the part of the traffic to mirror. Supported values: 0 to 100.",
											MarkdownDescription: "Percent defines the part of the traffic to mirror. Supported values: 0 to 100.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"port": schema.StringAttribute{
											Description:         "Port defines the port of a Kubernetes Service. This can be a reference to a named port.",
											MarkdownDescription: "Port defines the port of a Kubernetes Service. This can be a reference to a named port.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"response_forwarding": schema.SingleNestedAttribute{
											Description:         "ResponseForwarding defines how Traefik forwards the response from the upstream Kubernetes Service to the client.",
											MarkdownDescription: "ResponseForwarding defines how Traefik forwards the response from the upstream Kubernetes Service to the client.",
											Attributes: map[string]schema.Attribute{
												"flush_interval": schema.StringAttribute{
													Description:         "FlushInterval defines the interval, in milliseconds, in between flushes to the client while copying the response body. A negative value means to flush immediately after each write to the client. This configuration is ignored when ReverseProxy recognizes a response as a streaming response; for such responses, writes are flushed to the client immediately. Default: 100ms",
													MarkdownDescription: "FlushInterval defines the interval, in milliseconds, in between flushes to the client while copying the response body. A negative value means to flush immediately after each write to the client. This configuration is ignored when ReverseProxy recognizes a response as a streaming response; for such responses, writes are flushed to the client immediately. Default: 100ms",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"scheme": schema.StringAttribute{
											Description:         "Scheme defines the scheme to use for the request to the upstream Kubernetes Service. It defaults to https when Kubernetes Service port is 443, http otherwise.",
											MarkdownDescription: "Scheme defines the scheme to use for the request to the upstream Kubernetes Service. It defaults to https when Kubernetes Service port is 443, http otherwise.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"servers_transport": schema.StringAttribute{
											Description:         "ServersTransport defines the name of ServersTransport resource to use. It allows to configure the transport between Traefik and your servers. Can only be used on a Kubernetes Service.",
											MarkdownDescription: "ServersTransport defines the name of ServersTransport resource to use. It allows to configure the transport between Traefik and your servers. Can only be used on a Kubernetes Service.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"sticky": schema.SingleNestedAttribute{
											Description:         "Sticky defines the sticky sessions configuration. More info: https://doc.traefik.io/traefik/v3.2/routing/services/#sticky-sessions",
											MarkdownDescription: "Sticky defines the sticky sessions configuration. More info: https://doc.traefik.io/traefik/v3.2/routing/services/#sticky-sessions",
											Attributes: map[string]schema.Attribute{
												"cookie": schema.SingleNestedAttribute{
													Description:         "Cookie defines the sticky cookie configuration.",
													MarkdownDescription: "Cookie defines the sticky cookie configuration.",
													Attributes: map[string]schema.Attribute{
														"http_only": schema.BoolAttribute{
															Description:         "HTTPOnly defines whether the cookie can be accessed by client-side APIs, such as JavaScript.",
															MarkdownDescription: "HTTPOnly defines whether the cookie can be accessed by client-side APIs, such as JavaScript.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"max_age": schema.Int64Attribute{
															Description:         "MaxAge indicates the number of seconds until the cookie expires. When set to a negative number, the cookie expires immediately. When set to zero, the cookie never expires.",
															MarkdownDescription: "MaxAge indicates the number of seconds until the cookie expires. When set to a negative number, the cookie expires immediately. When set to zero, the cookie never expires.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name defines the Cookie name.",
															MarkdownDescription: "Name defines the Cookie name.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"same_site": schema.StringAttribute{
															Description:         "SameSite defines the same site policy. More info: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie/SameSite",
															MarkdownDescription: "SameSite defines the same site policy. More info: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie/SameSite",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"secure": schema.BoolAttribute{
															Description:         "Secure defines whether the cookie can only be transmitted over an encrypted connection (i.e. HTTPS).",
															MarkdownDescription: "Secure defines whether the cookie can only be transmitted over an encrypted connection (i.e. HTTPS).",
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

										"strategy": schema.StringAttribute{
											Description:         "Strategy defines the load balancing strategy between the servers. RoundRobin is the only supported value at the moment.",
											MarkdownDescription: "Strategy defines the load balancing strategy between the servers. RoundRobin is the only supported value at the moment.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"weight": schema.Int64Attribute{
											Description:         "Weight defines the weight and should only be specified when Name references a TraefikService object (and to be precise, one that embeds a Weighted Round Robin).",
											MarkdownDescription: "Weight defines the weight and should only be specified when Name references a TraefikService object (and to be precise, one that embeds a Weighted Round Robin).",
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

							"name": schema.StringAttribute{
								Description:         "Name defines the name of the referenced Kubernetes Service or TraefikService. The differentiation between the two is specified in the Kind field.",
								MarkdownDescription: "Name defines the name of the referenced Kubernetes Service or TraefikService. The differentiation between the two is specified in the Kind field.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace defines the namespace of the referenced Kubernetes Service or TraefikService.",
								MarkdownDescription: "Namespace defines the namespace of the referenced Kubernetes Service or TraefikService.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"native_lb": schema.BoolAttribute{
								Description:         "NativeLB controls, when creating the load-balancer, whether the LB's children are directly the pods IPs or if the only child is the Kubernetes Service clusterIP. The Kubernetes Service itself does load-balance to the pods. By default, NativeLB is false.",
								MarkdownDescription: "NativeLB controls, when creating the load-balancer, whether the LB's children are directly the pods IPs or if the only child is the Kubernetes Service clusterIP. The Kubernetes Service itself does load-balance to the pods. By default, NativeLB is false.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"node_port_lb": schema.BoolAttribute{
								Description:         "NodePortLB controls, when creating the load-balancer, whether the LB's children are directly the nodes internal IPs using the nodePort when the service type is NodePort. It allows services to be reachable when Traefik runs externally from the Kubernetes cluster but within the same network of the nodes. By default, NodePortLB is false.",
								MarkdownDescription: "NodePortLB controls, when creating the load-balancer, whether the LB's children are directly the nodes internal IPs using the nodePort when the service type is NodePort. It allows services to be reachable when Traefik runs externally from the Kubernetes cluster but within the same network of the nodes. By default, NodePortLB is false.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pass_host_header": schema.BoolAttribute{
								Description:         "PassHostHeader defines whether the client Host header is forwarded to the upstream Kubernetes Service. By default, passHostHeader is true.",
								MarkdownDescription: "PassHostHeader defines whether the client Host header is forwarded to the upstream Kubernetes Service. By default, passHostHeader is true.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port": schema.StringAttribute{
								Description:         "Port defines the port of a Kubernetes Service. This can be a reference to a named port.",
								MarkdownDescription: "Port defines the port of a Kubernetes Service. This can be a reference to a named port.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"response_forwarding": schema.SingleNestedAttribute{
								Description:         "ResponseForwarding defines how Traefik forwards the response from the upstream Kubernetes Service to the client.",
								MarkdownDescription: "ResponseForwarding defines how Traefik forwards the response from the upstream Kubernetes Service to the client.",
								Attributes: map[string]schema.Attribute{
									"flush_interval": schema.StringAttribute{
										Description:         "FlushInterval defines the interval, in milliseconds, in between flushes to the client while copying the response body. A negative value means to flush immediately after each write to the client. This configuration is ignored when ReverseProxy recognizes a response as a streaming response; for such responses, writes are flushed to the client immediately. Default: 100ms",
										MarkdownDescription: "FlushInterval defines the interval, in milliseconds, in between flushes to the client while copying the response body. A negative value means to flush immediately after each write to the client. This configuration is ignored when ReverseProxy recognizes a response as a streaming response; for such responses, writes are flushed to the client immediately. Default: 100ms",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"scheme": schema.StringAttribute{
								Description:         "Scheme defines the scheme to use for the request to the upstream Kubernetes Service. It defaults to https when Kubernetes Service port is 443, http otherwise.",
								MarkdownDescription: "Scheme defines the scheme to use for the request to the upstream Kubernetes Service. It defaults to https when Kubernetes Service port is 443, http otherwise.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"servers_transport": schema.StringAttribute{
								Description:         "ServersTransport defines the name of ServersTransport resource to use. It allows to configure the transport between Traefik and your servers. Can only be used on a Kubernetes Service.",
								MarkdownDescription: "ServersTransport defines the name of ServersTransport resource to use. It allows to configure the transport between Traefik and your servers. Can only be used on a Kubernetes Service.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"sticky": schema.SingleNestedAttribute{
								Description:         "Sticky defines the sticky sessions configuration. More info: https://doc.traefik.io/traefik/v3.2/routing/services/#sticky-sessions",
								MarkdownDescription: "Sticky defines the sticky sessions configuration. More info: https://doc.traefik.io/traefik/v3.2/routing/services/#sticky-sessions",
								Attributes: map[string]schema.Attribute{
									"cookie": schema.SingleNestedAttribute{
										Description:         "Cookie defines the sticky cookie configuration.",
										MarkdownDescription: "Cookie defines the sticky cookie configuration.",
										Attributes: map[string]schema.Attribute{
											"http_only": schema.BoolAttribute{
												Description:         "HTTPOnly defines whether the cookie can be accessed by client-side APIs, such as JavaScript.",
												MarkdownDescription: "HTTPOnly defines whether the cookie can be accessed by client-side APIs, such as JavaScript.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_age": schema.Int64Attribute{
												Description:         "MaxAge indicates the number of seconds until the cookie expires. When set to a negative number, the cookie expires immediately. When set to zero, the cookie never expires.",
												MarkdownDescription: "MaxAge indicates the number of seconds until the cookie expires. When set to a negative number, the cookie expires immediately. When set to zero, the cookie never expires.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name defines the Cookie name.",
												MarkdownDescription: "Name defines the Cookie name.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"same_site": schema.StringAttribute{
												Description:         "SameSite defines the same site policy. More info: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie/SameSite",
												MarkdownDescription: "SameSite defines the same site policy. More info: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie/SameSite",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"secure": schema.BoolAttribute{
												Description:         "Secure defines whether the cookie can only be transmitted over an encrypted connection (i.e. HTTPS).",
												MarkdownDescription: "Secure defines whether the cookie can only be transmitted over an encrypted connection (i.e. HTTPS).",
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

							"strategy": schema.StringAttribute{
								Description:         "Strategy defines the load balancing strategy between the servers. RoundRobin is the only supported value at the moment.",
								MarkdownDescription: "Strategy defines the load balancing strategy between the servers. RoundRobin is the only supported value at the moment.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"weight": schema.Int64Attribute{
								Description:         "Weight defines the weight and should only be specified when Name references a TraefikService object (and to be precise, one that embeds a Weighted Round Robin).",
								MarkdownDescription: "Weight defines the weight and should only be specified when Name references a TraefikService object (and to be precise, one that embeds a Weighted Round Robin).",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
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
										"health_check": schema.SingleNestedAttribute{
											Description:         "Healthcheck defines health checks for ExternalName services.",
											MarkdownDescription: "Healthcheck defines health checks for ExternalName services.",
											Attributes: map[string]schema.Attribute{
												"follow_redirects": schema.BoolAttribute{
													Description:         "FollowRedirects defines whether redirects should be followed during the health check calls. Default: true",
													MarkdownDescription: "FollowRedirects defines whether redirects should be followed during the health check calls. Default: true",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"headers": schema.MapAttribute{
													Description:         "Headers defines custom headers to be sent to the health check endpoint.",
													MarkdownDescription: "Headers defines custom headers to be sent to the health check endpoint.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"hostname": schema.StringAttribute{
													Description:         "Hostname defines the value of hostname in the Host header of the health check request.",
													MarkdownDescription: "Hostname defines the value of hostname in the Host header of the health check request.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"interval": schema.StringAttribute{
													Description:         "Interval defines the frequency of the health check calls. Default: 30s",
													MarkdownDescription: "Interval defines the frequency of the health check calls. Default: 30s",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"method": schema.StringAttribute{
													Description:         "Method defines the healthcheck method.",
													MarkdownDescription: "Method defines the healthcheck method.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"mode": schema.StringAttribute{
													Description:         "Mode defines the health check mode. If defined to grpc, will use the gRPC health check protocol to probe the server. Default: http",
													MarkdownDescription: "Mode defines the health check mode. If defined to grpc, will use the gRPC health check protocol to probe the server. Default: http",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"path": schema.StringAttribute{
													Description:         "Path defines the server URL path for the health check endpoint.",
													MarkdownDescription: "Path defines the server URL path for the health check endpoint.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"port": schema.Int64Attribute{
													Description:         "Port defines the server URL port for the health check endpoint.",
													MarkdownDescription: "Port defines the server URL port for the health check endpoint.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"scheme": schema.StringAttribute{
													Description:         "Scheme replaces the server URL scheme for the health check endpoint.",
													MarkdownDescription: "Scheme replaces the server URL scheme for the health check endpoint.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"status": schema.Int64Attribute{
													Description:         "Status defines the expected HTTP status code of the response to the health check request.",
													MarkdownDescription: "Status defines the expected HTTP status code of the response to the health check request.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"timeout": schema.StringAttribute{
													Description:         "Timeout defines the maximum duration Traefik will wait for a health check request before considering the server unhealthy. Default: 5s",
													MarkdownDescription: "Timeout defines the maximum duration Traefik will wait for a health check request before considering the server unhealthy. Default: 5s",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"kind": schema.StringAttribute{
											Description:         "Kind defines the kind of the Service.",
											MarkdownDescription: "Kind defines the kind of the Service.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Service", "TraefikService"),
											},
										},

										"name": schema.StringAttribute{
											Description:         "Name defines the name of the referenced Kubernetes Service or TraefikService. The differentiation between the two is specified in the Kind field.",
											MarkdownDescription: "Name defines the name of the referenced Kubernetes Service or TraefikService. The differentiation between the two is specified in the Kind field.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"namespace": schema.StringAttribute{
											Description:         "Namespace defines the namespace of the referenced Kubernetes Service or TraefikService.",
											MarkdownDescription: "Namespace defines the namespace of the referenced Kubernetes Service or TraefikService.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"native_lb": schema.BoolAttribute{
											Description:         "NativeLB controls, when creating the load-balancer, whether the LB's children are directly the pods IPs or if the only child is the Kubernetes Service clusterIP. The Kubernetes Service itself does load-balance to the pods. By default, NativeLB is false.",
											MarkdownDescription: "NativeLB controls, when creating the load-balancer, whether the LB's children are directly the pods IPs or if the only child is the Kubernetes Service clusterIP. The Kubernetes Service itself does load-balance to the pods. By default, NativeLB is false.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"node_port_lb": schema.BoolAttribute{
											Description:         "NodePortLB controls, when creating the load-balancer, whether the LB's children are directly the nodes internal IPs using the nodePort when the service type is NodePort. It allows services to be reachable when Traefik runs externally from the Kubernetes cluster but within the same network of the nodes. By default, NodePortLB is false.",
											MarkdownDescription: "NodePortLB controls, when creating the load-balancer, whether the LB's children are directly the nodes internal IPs using the nodePort when the service type is NodePort. It allows services to be reachable when Traefik runs externally from the Kubernetes cluster but within the same network of the nodes. By default, NodePortLB is false.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"pass_host_header": schema.BoolAttribute{
											Description:         "PassHostHeader defines whether the client Host header is forwarded to the upstream Kubernetes Service. By default, passHostHeader is true.",
											MarkdownDescription: "PassHostHeader defines whether the client Host header is forwarded to the upstream Kubernetes Service. By default, passHostHeader is true.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"port": schema.StringAttribute{
											Description:         "Port defines the port of a Kubernetes Service. This can be a reference to a named port.",
											MarkdownDescription: "Port defines the port of a Kubernetes Service. This can be a reference to a named port.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"response_forwarding": schema.SingleNestedAttribute{
											Description:         "ResponseForwarding defines how Traefik forwards the response from the upstream Kubernetes Service to the client.",
											MarkdownDescription: "ResponseForwarding defines how Traefik forwards the response from the upstream Kubernetes Service to the client.",
											Attributes: map[string]schema.Attribute{
												"flush_interval": schema.StringAttribute{
													Description:         "FlushInterval defines the interval, in milliseconds, in between flushes to the client while copying the response body. A negative value means to flush immediately after each write to the client. This configuration is ignored when ReverseProxy recognizes a response as a streaming response; for such responses, writes are flushed to the client immediately. Default: 100ms",
													MarkdownDescription: "FlushInterval defines the interval, in milliseconds, in between flushes to the client while copying the response body. A negative value means to flush immediately after each write to the client. This configuration is ignored when ReverseProxy recognizes a response as a streaming response; for such responses, writes are flushed to the client immediately. Default: 100ms",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"scheme": schema.StringAttribute{
											Description:         "Scheme defines the scheme to use for the request to the upstream Kubernetes Service. It defaults to https when Kubernetes Service port is 443, http otherwise.",
											MarkdownDescription: "Scheme defines the scheme to use for the request to the upstream Kubernetes Service. It defaults to https when Kubernetes Service port is 443, http otherwise.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"servers_transport": schema.StringAttribute{
											Description:         "ServersTransport defines the name of ServersTransport resource to use. It allows to configure the transport between Traefik and your servers. Can only be used on a Kubernetes Service.",
											MarkdownDescription: "ServersTransport defines the name of ServersTransport resource to use. It allows to configure the transport between Traefik and your servers. Can only be used on a Kubernetes Service.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"sticky": schema.SingleNestedAttribute{
											Description:         "Sticky defines the sticky sessions configuration. More info: https://doc.traefik.io/traefik/v3.2/routing/services/#sticky-sessions",
											MarkdownDescription: "Sticky defines the sticky sessions configuration. More info: https://doc.traefik.io/traefik/v3.2/routing/services/#sticky-sessions",
											Attributes: map[string]schema.Attribute{
												"cookie": schema.SingleNestedAttribute{
													Description:         "Cookie defines the sticky cookie configuration.",
													MarkdownDescription: "Cookie defines the sticky cookie configuration.",
													Attributes: map[string]schema.Attribute{
														"http_only": schema.BoolAttribute{
															Description:         "HTTPOnly defines whether the cookie can be accessed by client-side APIs, such as JavaScript.",
															MarkdownDescription: "HTTPOnly defines whether the cookie can be accessed by client-side APIs, such as JavaScript.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"max_age": schema.Int64Attribute{
															Description:         "MaxAge indicates the number of seconds until the cookie expires. When set to a negative number, the cookie expires immediately. When set to zero, the cookie never expires.",
															MarkdownDescription: "MaxAge indicates the number of seconds until the cookie expires. When set to a negative number, the cookie expires immediately. When set to zero, the cookie never expires.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name defines the Cookie name.",
															MarkdownDescription: "Name defines the Cookie name.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"same_site": schema.StringAttribute{
															Description:         "SameSite defines the same site policy. More info: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie/SameSite",
															MarkdownDescription: "SameSite defines the same site policy. More info: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie/SameSite",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"secure": schema.BoolAttribute{
															Description:         "Secure defines whether the cookie can only be transmitted over an encrypted connection (i.e. HTTPS).",
															MarkdownDescription: "Secure defines whether the cookie can only be transmitted over an encrypted connection (i.e. HTTPS).",
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

										"strategy": schema.StringAttribute{
											Description:         "Strategy defines the load balancing strategy between the servers. RoundRobin is the only supported value at the moment.",
											MarkdownDescription: "Strategy defines the load balancing strategy between the servers. RoundRobin is the only supported value at the moment.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"weight": schema.Int64Attribute{
											Description:         "Weight defines the weight and should only be specified when Name references a TraefikService object (and to be precise, one that embeds a Weighted Round Robin).",
											MarkdownDescription: "Weight defines the weight and should only be specified when Name references a TraefikService object (and to be precise, one that embeds a Weighted Round Robin).",
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

							"sticky": schema.SingleNestedAttribute{
								Description:         "Sticky defines whether sticky sessions are enabled. More info: https://doc.traefik.io/traefik/v3.2/routing/providers/kubernetes-crd/#stickiness-and-load-balancing",
								MarkdownDescription: "Sticky defines whether sticky sessions are enabled. More info: https://doc.traefik.io/traefik/v3.2/routing/providers/kubernetes-crd/#stickiness-and-load-balancing",
								Attributes: map[string]schema.Attribute{
									"cookie": schema.SingleNestedAttribute{
										Description:         "Cookie defines the sticky cookie configuration.",
										MarkdownDescription: "Cookie defines the sticky cookie configuration.",
										Attributes: map[string]schema.Attribute{
											"http_only": schema.BoolAttribute{
												Description:         "HTTPOnly defines whether the cookie can be accessed by client-side APIs, such as JavaScript.",
												MarkdownDescription: "HTTPOnly defines whether the cookie can be accessed by client-side APIs, such as JavaScript.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_age": schema.Int64Attribute{
												Description:         "MaxAge indicates the number of seconds until the cookie expires. When set to a negative number, the cookie expires immediately. When set to zero, the cookie never expires.",
												MarkdownDescription: "MaxAge indicates the number of seconds until the cookie expires. When set to a negative number, the cookie expires immediately. When set to zero, the cookie never expires.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name defines the Cookie name.",
												MarkdownDescription: "Name defines the Cookie name.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"same_site": schema.StringAttribute{
												Description:         "SameSite defines the same site policy. More info: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie/SameSite",
												MarkdownDescription: "SameSite defines the same site policy. More info: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie/SameSite",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"secure": schema.BoolAttribute{
												Description:         "Secure defines whether the cookie can only be transmitted over an encrypted connection (i.e. HTTPS).",
												MarkdownDescription: "Secure defines whether the cookie can only be transmitted over an encrypted connection (i.e. HTTPS).",
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
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *TraefikIoTraefikServiceV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_traefik_io_traefik_service_v1alpha1_manifest")

	var model TraefikIoTraefikServiceV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("traefik.io/v1alpha1")
	model.Kind = pointer.String("TraefikService")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
