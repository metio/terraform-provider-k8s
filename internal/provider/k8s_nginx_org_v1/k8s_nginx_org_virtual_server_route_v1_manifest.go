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
	_ datasource.DataSource = &K8SNginxOrgVirtualServerRouteV1Manifest{}
)

func NewK8SNginxOrgVirtualServerRouteV1Manifest() datasource.DataSource {
	return &K8SNginxOrgVirtualServerRouteV1Manifest{}
}

type K8SNginxOrgVirtualServerRouteV1Manifest struct{}

type K8SNginxOrgVirtualServerRouteV1ManifestData struct {
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
		Host             *string `tfsdk:"host" json:"host,omitempty"`
		IngressClassName *string `tfsdk:"ingress_class_name" json:"ingressClassName,omitempty"`
		Subroutes        *[]struct {
			Action *struct {
				Pass  *string `tfsdk:"pass" json:"pass,omitempty"`
				Proxy *struct {
					RequestHeaders *struct {
						Pass *bool `tfsdk:"pass" json:"pass,omitempty"`
						Set  *[]struct {
							Name  *string `tfsdk:"name" json:"name,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"set" json:"set,omitempty"`
					} `tfsdk:"request_headers" json:"requestHeaders,omitempty"`
					ResponseHeaders *struct {
						Add *[]struct {
							Always *bool   `tfsdk:"always" json:"always,omitempty"`
							Name   *string `tfsdk:"name" json:"name,omitempty"`
							Value  *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"add" json:"add,omitempty"`
						Hide   *[]string `tfsdk:"hide" json:"hide,omitempty"`
						Ignore *[]string `tfsdk:"ignore" json:"ignore,omitempty"`
						Pass   *[]string `tfsdk:"pass" json:"pass,omitempty"`
					} `tfsdk:"response_headers" json:"responseHeaders,omitempty"`
					RewritePath *string `tfsdk:"rewrite_path" json:"rewritePath,omitempty"`
					Upstream    *string `tfsdk:"upstream" json:"upstream,omitempty"`
				} `tfsdk:"proxy" json:"proxy,omitempty"`
				Redirect *struct {
					Code *int64  `tfsdk:"code" json:"code,omitempty"`
					Url  *string `tfsdk:"url" json:"url,omitempty"`
				} `tfsdk:"redirect" json:"redirect,omitempty"`
				Return *struct {
					Body    *string `tfsdk:"body" json:"body,omitempty"`
					Code    *int64  `tfsdk:"code" json:"code,omitempty"`
					Headers *[]struct {
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"headers" json:"headers,omitempty"`
					Type *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"return" json:"return,omitempty"`
			} `tfsdk:"action" json:"action,omitempty"`
			Dos        *string `tfsdk:"dos" json:"dos,omitempty"`
			ErrorPages *[]struct {
				Codes    *[]string `tfsdk:"codes" json:"codes,omitempty"`
				Redirect *struct {
					Code *int64  `tfsdk:"code" json:"code,omitempty"`
					Url  *string `tfsdk:"url" json:"url,omitempty"`
				} `tfsdk:"redirect" json:"redirect,omitempty"`
				Return *struct {
					Body    *string `tfsdk:"body" json:"body,omitempty"`
					Code    *int64  `tfsdk:"code" json:"code,omitempty"`
					Headers *[]struct {
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"headers" json:"headers,omitempty"`
					Type *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"return" json:"return,omitempty"`
			} `tfsdk:"error_pages" json:"errorPages,omitempty"`
			Location_snippets *string `tfsdk:"location_snippets" json:"location-snippets,omitempty"`
			Matches           *[]struct {
				Action *struct {
					Pass  *string `tfsdk:"pass" json:"pass,omitempty"`
					Proxy *struct {
						RequestHeaders *struct {
							Pass *bool `tfsdk:"pass" json:"pass,omitempty"`
							Set  *[]struct {
								Name  *string `tfsdk:"name" json:"name,omitempty"`
								Value *string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"set" json:"set,omitempty"`
						} `tfsdk:"request_headers" json:"requestHeaders,omitempty"`
						ResponseHeaders *struct {
							Add *[]struct {
								Always *bool   `tfsdk:"always" json:"always,omitempty"`
								Name   *string `tfsdk:"name" json:"name,omitempty"`
								Value  *string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"add" json:"add,omitempty"`
							Hide   *[]string `tfsdk:"hide" json:"hide,omitempty"`
							Ignore *[]string `tfsdk:"ignore" json:"ignore,omitempty"`
							Pass   *[]string `tfsdk:"pass" json:"pass,omitempty"`
						} `tfsdk:"response_headers" json:"responseHeaders,omitempty"`
						RewritePath *string `tfsdk:"rewrite_path" json:"rewritePath,omitempty"`
						Upstream    *string `tfsdk:"upstream" json:"upstream,omitempty"`
					} `tfsdk:"proxy" json:"proxy,omitempty"`
					Redirect *struct {
						Code *int64  `tfsdk:"code" json:"code,omitempty"`
						Url  *string `tfsdk:"url" json:"url,omitempty"`
					} `tfsdk:"redirect" json:"redirect,omitempty"`
					Return *struct {
						Body    *string `tfsdk:"body" json:"body,omitempty"`
						Code    *int64  `tfsdk:"code" json:"code,omitempty"`
						Headers *[]struct {
							Name  *string `tfsdk:"name" json:"name,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"headers" json:"headers,omitempty"`
						Type *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"return" json:"return,omitempty"`
				} `tfsdk:"action" json:"action,omitempty"`
				Conditions *[]struct {
					Argument *string `tfsdk:"argument" json:"argument,omitempty"`
					Cookie   *string `tfsdk:"cookie" json:"cookie,omitempty"`
					Header   *string `tfsdk:"header" json:"header,omitempty"`
					Value    *string `tfsdk:"value" json:"value,omitempty"`
					Variable *string `tfsdk:"variable" json:"variable,omitempty"`
				} `tfsdk:"conditions" json:"conditions,omitempty"`
				Splits *[]struct {
					Action *struct {
						Pass  *string `tfsdk:"pass" json:"pass,omitempty"`
						Proxy *struct {
							RequestHeaders *struct {
								Pass *bool `tfsdk:"pass" json:"pass,omitempty"`
								Set  *[]struct {
									Name  *string `tfsdk:"name" json:"name,omitempty"`
									Value *string `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"set" json:"set,omitempty"`
							} `tfsdk:"request_headers" json:"requestHeaders,omitempty"`
							ResponseHeaders *struct {
								Add *[]struct {
									Always *bool   `tfsdk:"always" json:"always,omitempty"`
									Name   *string `tfsdk:"name" json:"name,omitempty"`
									Value  *string `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"add" json:"add,omitempty"`
								Hide   *[]string `tfsdk:"hide" json:"hide,omitempty"`
								Ignore *[]string `tfsdk:"ignore" json:"ignore,omitempty"`
								Pass   *[]string `tfsdk:"pass" json:"pass,omitempty"`
							} `tfsdk:"response_headers" json:"responseHeaders,omitempty"`
							RewritePath *string `tfsdk:"rewrite_path" json:"rewritePath,omitempty"`
							Upstream    *string `tfsdk:"upstream" json:"upstream,omitempty"`
						} `tfsdk:"proxy" json:"proxy,omitempty"`
						Redirect *struct {
							Code *int64  `tfsdk:"code" json:"code,omitempty"`
							Url  *string `tfsdk:"url" json:"url,omitempty"`
						} `tfsdk:"redirect" json:"redirect,omitempty"`
						Return *struct {
							Body    *string `tfsdk:"body" json:"body,omitempty"`
							Code    *int64  `tfsdk:"code" json:"code,omitempty"`
							Headers *[]struct {
								Name  *string `tfsdk:"name" json:"name,omitempty"`
								Value *string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"headers" json:"headers,omitempty"`
							Type *string `tfsdk:"type" json:"type,omitempty"`
						} `tfsdk:"return" json:"return,omitempty"`
					} `tfsdk:"action" json:"action,omitempty"`
					Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
				} `tfsdk:"splits" json:"splits,omitempty"`
			} `tfsdk:"matches" json:"matches,omitempty"`
			Path     *string `tfsdk:"path" json:"path,omitempty"`
			Policies *[]struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"policies" json:"policies,omitempty"`
			Route  *string `tfsdk:"route" json:"route,omitempty"`
			Splits *[]struct {
				Action *struct {
					Pass  *string `tfsdk:"pass" json:"pass,omitempty"`
					Proxy *struct {
						RequestHeaders *struct {
							Pass *bool `tfsdk:"pass" json:"pass,omitempty"`
							Set  *[]struct {
								Name  *string `tfsdk:"name" json:"name,omitempty"`
								Value *string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"set" json:"set,omitempty"`
						} `tfsdk:"request_headers" json:"requestHeaders,omitempty"`
						ResponseHeaders *struct {
							Add *[]struct {
								Always *bool   `tfsdk:"always" json:"always,omitempty"`
								Name   *string `tfsdk:"name" json:"name,omitempty"`
								Value  *string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"add" json:"add,omitempty"`
							Hide   *[]string `tfsdk:"hide" json:"hide,omitempty"`
							Ignore *[]string `tfsdk:"ignore" json:"ignore,omitempty"`
							Pass   *[]string `tfsdk:"pass" json:"pass,omitempty"`
						} `tfsdk:"response_headers" json:"responseHeaders,omitempty"`
						RewritePath *string `tfsdk:"rewrite_path" json:"rewritePath,omitempty"`
						Upstream    *string `tfsdk:"upstream" json:"upstream,omitempty"`
					} `tfsdk:"proxy" json:"proxy,omitempty"`
					Redirect *struct {
						Code *int64  `tfsdk:"code" json:"code,omitempty"`
						Url  *string `tfsdk:"url" json:"url,omitempty"`
					} `tfsdk:"redirect" json:"redirect,omitempty"`
					Return *struct {
						Body    *string `tfsdk:"body" json:"body,omitempty"`
						Code    *int64  `tfsdk:"code" json:"code,omitempty"`
						Headers *[]struct {
							Name  *string `tfsdk:"name" json:"name,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"headers" json:"headers,omitempty"`
						Type *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"return" json:"return,omitempty"`
				} `tfsdk:"action" json:"action,omitempty"`
				Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
			} `tfsdk:"splits" json:"splits,omitempty"`
		} `tfsdk:"subroutes" json:"subroutes,omitempty"`
		Upstreams *[]struct {
			Backup      *string `tfsdk:"backup" json:"backup,omitempty"`
			BackupPort  *int64  `tfsdk:"backup_port" json:"backupPort,omitempty"`
			Buffer_size *string `tfsdk:"buffer_size" json:"buffer-size,omitempty"`
			Buffering   *bool   `tfsdk:"buffering" json:"buffering,omitempty"`
			Buffers     *struct {
				Number *int64  `tfsdk:"number" json:"number,omitempty"`
				Size   *string `tfsdk:"size" json:"size,omitempty"`
			} `tfsdk:"buffers" json:"buffers,omitempty"`
			Client_max_body_size *string `tfsdk:"client_max_body_size" json:"client-max-body-size,omitempty"`
			Connect_timeout      *string `tfsdk:"connect_timeout" json:"connect-timeout,omitempty"`
			Fail_timeout         *string `tfsdk:"fail_timeout" json:"fail-timeout,omitempty"`
			HealthCheck          *struct {
				Connect_timeout *string `tfsdk:"connect_timeout" json:"connect-timeout,omitempty"`
				Enable          *bool   `tfsdk:"enable" json:"enable,omitempty"`
				Fails           *int64  `tfsdk:"fails" json:"fails,omitempty"`
				GrpcService     *string `tfsdk:"grpc_service" json:"grpcService,omitempty"`
				GrpcStatus      *int64  `tfsdk:"grpc_status" json:"grpcStatus,omitempty"`
				Headers         *[]struct {
					Name  *string `tfsdk:"name" json:"name,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"headers" json:"headers,omitempty"`
				Interval       *string `tfsdk:"interval" json:"interval,omitempty"`
				Jitter         *string `tfsdk:"jitter" json:"jitter,omitempty"`
				Keepalive_time *string `tfsdk:"keepalive_time" json:"keepalive-time,omitempty"`
				Mandatory      *bool   `tfsdk:"mandatory" json:"mandatory,omitempty"`
				Passes         *int64  `tfsdk:"passes" json:"passes,omitempty"`
				Path           *string `tfsdk:"path" json:"path,omitempty"`
				Persistent     *bool   `tfsdk:"persistent" json:"persistent,omitempty"`
				Port           *int64  `tfsdk:"port" json:"port,omitempty"`
				Read_timeout   *string `tfsdk:"read_timeout" json:"read-timeout,omitempty"`
				Send_timeout   *string `tfsdk:"send_timeout" json:"send-timeout,omitempty"`
				StatusMatch    *string `tfsdk:"status_match" json:"statusMatch,omitempty"`
				Tls            *struct {
					Enable *bool `tfsdk:"enable" json:"enable,omitempty"`
				} `tfsdk:"tls" json:"tls,omitempty"`
			} `tfsdk:"health_check" json:"healthCheck,omitempty"`
			Keepalive             *int64  `tfsdk:"keepalive" json:"keepalive,omitempty"`
			Lb_method             *string `tfsdk:"lb_method" json:"lb-method,omitempty"`
			Max_conns             *int64  `tfsdk:"max_conns" json:"max-conns,omitempty"`
			Max_fails             *int64  `tfsdk:"max_fails" json:"max-fails,omitempty"`
			Name                  *string `tfsdk:"name" json:"name,omitempty"`
			Next_upstream         *string `tfsdk:"next_upstream" json:"next-upstream,omitempty"`
			Next_upstream_timeout *string `tfsdk:"next_upstream_timeout" json:"next-upstream-timeout,omitempty"`
			Next_upstream_tries   *int64  `tfsdk:"next_upstream_tries" json:"next-upstream-tries,omitempty"`
			Ntlm                  *bool   `tfsdk:"ntlm" json:"ntlm,omitempty"`
			Port                  *int64  `tfsdk:"port" json:"port,omitempty"`
			Queue                 *struct {
				Size    *int64  `tfsdk:"size" json:"size,omitempty"`
				Timeout *string `tfsdk:"timeout" json:"timeout,omitempty"`
			} `tfsdk:"queue" json:"queue,omitempty"`
			Read_timeout  *string `tfsdk:"read_timeout" json:"read-timeout,omitempty"`
			Send_timeout  *string `tfsdk:"send_timeout" json:"send-timeout,omitempty"`
			Service       *string `tfsdk:"service" json:"service,omitempty"`
			SessionCookie *struct {
				Domain   *string `tfsdk:"domain" json:"domain,omitempty"`
				Enable   *bool   `tfsdk:"enable" json:"enable,omitempty"`
				Expires  *string `tfsdk:"expires" json:"expires,omitempty"`
				HttpOnly *bool   `tfsdk:"http_only" json:"httpOnly,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Path     *string `tfsdk:"path" json:"path,omitempty"`
				Samesite *string `tfsdk:"samesite" json:"samesite,omitempty"`
				Secure   *bool   `tfsdk:"secure" json:"secure,omitempty"`
			} `tfsdk:"session_cookie" json:"sessionCookie,omitempty"`
			Slow_start  *string            `tfsdk:"slow_start" json:"slow-start,omitempty"`
			Subselector *map[string]string `tfsdk:"subselector" json:"subselector,omitempty"`
			Tls         *struct {
				Enable *bool `tfsdk:"enable" json:"enable,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
			Type           *string `tfsdk:"type" json:"type,omitempty"`
			Use_cluster_ip *bool   `tfsdk:"use_cluster_ip" json:"use-cluster-ip,omitempty"`
		} `tfsdk:"upstreams" json:"upstreams,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *K8SNginxOrgVirtualServerRouteV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_k8s_nginx_org_virtual_server_route_v1_manifest"
}

func (r *K8SNginxOrgVirtualServerRouteV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "VirtualServerRoute defines the VirtualServerRoute resource.",
		MarkdownDescription: "VirtualServerRoute defines the VirtualServerRoute resource.",
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
				Description:         "VirtualServerRouteSpec is the spec of the VirtualServerRoute resource.",
				MarkdownDescription: "VirtualServerRouteSpec is the spec of the VirtualServerRoute resource.",
				Attributes: map[string]schema.Attribute{
					"host": schema.StringAttribute{
						Description:         "The host (domain name) of the server. Must be a valid subdomain as defined in RFC 1123, such as my-app or hello.example.com. When using a wildcard domain like *.example.com the domain must be contained in double quotes. Must be the same as the host of the VirtualServer that references this resource.",
						MarkdownDescription: "The host (domain name) of the server. Must be a valid subdomain as defined in RFC 1123, such as my-app or hello.example.com. When using a wildcard domain like *.example.com the domain must be contained in double quotes. Must be the same as the host of the VirtualServer that references this resource.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ingress_class_name": schema.StringAttribute{
						Description:         "Specifies which Ingress Controller must handle the VirtualServerRoute resource. Must be the same as the ingressClassName of the VirtualServer that references this resource.",
						MarkdownDescription: "Specifies which Ingress Controller must handle the VirtualServerRoute resource. Must be the same as the ingressClassName of the VirtualServer that references this resource.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"subroutes": schema.ListNestedAttribute{
						Description:         "A list of subroutes.",
						MarkdownDescription: "A list of subroutes.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"action": schema.SingleNestedAttribute{
									Description:         "The default action to perform for a request.",
									MarkdownDescription: "The default action to perform for a request.",
									Attributes: map[string]schema.Attribute{
										"pass": schema.StringAttribute{
											Description:         "Passes requests to an upstream. The upstream with that name must be defined in the resource.",
											MarkdownDescription: "Passes requests to an upstream. The upstream with that name must be defined in the resource.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"proxy": schema.SingleNestedAttribute{
											Description:         "Passes requests to an upstream with the ability to modify the request/response (for example, rewrite the URI or modify the headers).",
											MarkdownDescription: "Passes requests to an upstream with the ability to modify the request/response (for example, rewrite the URI or modify the headers).",
											Attributes: map[string]schema.Attribute{
												"request_headers": schema.SingleNestedAttribute{
													Description:         "The request headers modifications.",
													MarkdownDescription: "The request headers modifications.",
													Attributes: map[string]schema.Attribute{
														"pass": schema.BoolAttribute{
															Description:         "Passes the original request headers to the proxied upstream server. Default is true.",
															MarkdownDescription: "Passes the original request headers to the proxied upstream server. Default is true.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"set": schema.ListNestedAttribute{
															Description:         "Allows redefining or appending fields to present request headers passed to the proxied upstream servers.",
															MarkdownDescription: "Allows redefining or appending fields to present request headers passed to the proxied upstream servers.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "The name of the header.",
																		MarkdownDescription: "The name of the header.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
																		Description:         "The value of the header.",
																		MarkdownDescription: "The value of the header.",
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

												"response_headers": schema.SingleNestedAttribute{
													Description:         "The response headers modifications.",
													MarkdownDescription: "The response headers modifications.",
													Attributes: map[string]schema.Attribute{
														"add": schema.ListNestedAttribute{
															Description:         "Adds headers to the response to the client.",
															MarkdownDescription: "Adds headers to the response to the client.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"always": schema.BoolAttribute{
																		Description:         "If set to true, add the header regardless of the response status code**. Default is false.",
																		MarkdownDescription: "If set to true, add the header regardless of the response status code**. Default is false.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "The name of the header.",
																		MarkdownDescription: "The name of the header.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
																		Description:         "The value of the header.",
																		MarkdownDescription: "The value of the header.",
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

														"hide": schema.ListAttribute{
															Description:         "The headers that will not be passed* in the response to the client from a proxied upstream server.",
															MarkdownDescription: "The headers that will not be passed* in the response to the client from a proxied upstream server.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"ignore": schema.ListAttribute{
															Description:         "Disables processing of certain headers** to the client from a proxied upstream server.",
															MarkdownDescription: "Disables processing of certain headers** to the client from a proxied upstream server.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"pass": schema.ListAttribute{
															Description:         "Allows passing the hidden header fields* to the client from a proxied upstream server.",
															MarkdownDescription: "Allows passing the hidden header fields* to the client from a proxied upstream server.",
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

												"rewrite_path": schema.StringAttribute{
													Description:         "The rewritten URI. If the route path is a regular expression – starts with ~ – the rewritePath can include capture groups with $1-9. For example $1 for the first group, and so on. For more information, check the rewrite example.",
													MarkdownDescription: "The rewritten URI. If the route path is a regular expression – starts with ~ – the rewritePath can include capture groups with $1-9. For example $1 for the first group, and so on. For more information, check the rewrite example.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"upstream": schema.StringAttribute{
													Description:         "The name of the upstream which the requests will be proxied to. The upstream with that name must be defined in the resource.",
													MarkdownDescription: "The name of the upstream which the requests will be proxied to. The upstream with that name must be defined in the resource.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"redirect": schema.SingleNestedAttribute{
											Description:         "Redirects requests to a provided URL.",
											MarkdownDescription: "Redirects requests to a provided URL.",
											Attributes: map[string]schema.Attribute{
												"code": schema.Int64Attribute{
													Description:         "The status code of a redirect. The allowed values are: 301, 302, 307 or 308. The default is 301.",
													MarkdownDescription: "The status code of a redirect. The allowed values are: 301, 302, 307 or 308. The default is 301.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"url": schema.StringAttribute{
													Description:         "The URL to redirect the request to. Supported NGINX variables: $scheme, $http_x_forwarded_proto, $request_uri or $host. Variables must be enclosed in curly braces. For example: ${host}${request_uri}.",
													MarkdownDescription: "The URL to redirect the request to. Supported NGINX variables: $scheme, $http_x_forwarded_proto, $request_uri or $host. Variables must be enclosed in curly braces. For example: ${host}${request_uri}.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"return": schema.SingleNestedAttribute{
											Description:         "Returns a preconfigured response.",
											MarkdownDescription: "Returns a preconfigured response.",
											Attributes: map[string]schema.Attribute{
												"body": schema.StringAttribute{
													Description:         "The body of the response. Supports NGINX variables*. Variables must be enclosed in curly brackets. For example: Request is ${request_uri}n.",
													MarkdownDescription: "The body of the response. Supports NGINX variables*. Variables must be enclosed in curly brackets. For example: Request is ${request_uri}n.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"code": schema.Int64Attribute{
													Description:         "The status code of the response. The allowed values are: 2XX, 4XX or 5XX. The default is 200.",
													MarkdownDescription: "The status code of the response. The allowed values are: 2XX, 4XX or 5XX. The default is 200.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"headers": schema.ListNestedAttribute{
													Description:         "The custom headers of the response.",
													MarkdownDescription: "The custom headers of the response.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "The name of the header.",
																MarkdownDescription: "The name of the header.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"value": schema.StringAttribute{
																Description:         "The value of the header.",
																MarkdownDescription: "The value of the header.",
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

												"type": schema.StringAttribute{
													Description:         "The MIME type of the response. The default is text/plain.",
													MarkdownDescription: "The MIME type of the response. The default is text/plain.",
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

								"dos": schema.StringAttribute{
									Description:         "A reference to a DosProtectedResource, setting this enables DOS protection of the VirtualServer route.",
									MarkdownDescription: "A reference to a DosProtectedResource, setting this enables DOS protection of the VirtualServer route.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"error_pages": schema.ListNestedAttribute{
									Description:         "The custom responses for error codes. NGINX will use those responses instead of returning the error responses from the upstream servers or the default responses generated by NGINX. A custom response can be a redirect or a canned response. For example, a redirect to another URL if an upstream server responded with a 404 status code.",
									MarkdownDescription: "The custom responses for error codes. NGINX will use those responses instead of returning the error responses from the upstream servers or the default responses generated by NGINX. A custom response can be a redirect or a canned response. For example, a redirect to another URL if an upstream server responded with a 404 status code.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"codes": schema.ListAttribute{
												Description:         "A list of error status codes.",
												MarkdownDescription: "A list of error status codes.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"redirect": schema.SingleNestedAttribute{
												Description:         "The canned response action for the given status codes.",
												MarkdownDescription: "The canned response action for the given status codes.",
												Attributes: map[string]schema.Attribute{
													"code": schema.Int64Attribute{
														Description:         "The status code of a redirect. The allowed values are: 301, 302, 307 or 308. The default is 301.",
														MarkdownDescription: "The status code of a redirect. The allowed values are: 301, 302, 307 or 308. The default is 301.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"url": schema.StringAttribute{
														Description:         "The URL to redirect the request to. Supported NGINX variables: $scheme, $http_x_forwarded_proto, $request_uri or $host. Variables must be enclosed in curly braces. For example: ${host}${request_uri}.",
														MarkdownDescription: "The URL to redirect the request to. Supported NGINX variables: $scheme, $http_x_forwarded_proto, $request_uri or $host. Variables must be enclosed in curly braces. For example: ${host}${request_uri}.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"return": schema.SingleNestedAttribute{
												Description:         "The redirect action for the given status codes.",
												MarkdownDescription: "The redirect action for the given status codes.",
												Attributes: map[string]schema.Attribute{
													"body": schema.StringAttribute{
														Description:         "The body of the response. Supports NGINX variables*. Variables must be enclosed in curly brackets. For example: Request is ${request_uri}n.",
														MarkdownDescription: "The body of the response. Supports NGINX variables*. Variables must be enclosed in curly brackets. For example: Request is ${request_uri}n.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"code": schema.Int64Attribute{
														Description:         "The status code of the response. The allowed values are: 2XX, 4XX or 5XX. The default is 200.",
														MarkdownDescription: "The status code of the response. The allowed values are: 2XX, 4XX or 5XX. The default is 200.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"headers": schema.ListNestedAttribute{
														Description:         "The custom headers of the response.",
														MarkdownDescription: "The custom headers of the response.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "The name of the header.",
																	MarkdownDescription: "The name of the header.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"value": schema.StringAttribute{
																	Description:         "The value of the header.",
																	MarkdownDescription: "The value of the header.",
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

													"type": schema.StringAttribute{
														Description:         "The MIME type of the response. The default is text/plain.",
														MarkdownDescription: "The MIME type of the response. The default is text/plain.",
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

								"location_snippets": schema.StringAttribute{
									Description:         "Sets a custom snippet in the location context. Overrides the location-snippets ConfigMap key.",
									MarkdownDescription: "Sets a custom snippet in the location context. Overrides the location-snippets ConfigMap key.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"matches": schema.ListNestedAttribute{
									Description:         "The matching rules for advanced content-based routing. Requires the default Action or Splits. Unmatched requests will be handled by the default Action or Splits.",
									MarkdownDescription: "The matching rules for advanced content-based routing. Requires the default Action or Splits. Unmatched requests will be handled by the default Action or Splits.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"action": schema.SingleNestedAttribute{
												Description:         "The action to perform for a request.",
												MarkdownDescription: "The action to perform for a request.",
												Attributes: map[string]schema.Attribute{
													"pass": schema.StringAttribute{
														Description:         "Passes requests to an upstream. The upstream with that name must be defined in the resource.",
														MarkdownDescription: "Passes requests to an upstream. The upstream with that name must be defined in the resource.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"proxy": schema.SingleNestedAttribute{
														Description:         "Passes requests to an upstream with the ability to modify the request/response (for example, rewrite the URI or modify the headers).",
														MarkdownDescription: "Passes requests to an upstream with the ability to modify the request/response (for example, rewrite the URI or modify the headers).",
														Attributes: map[string]schema.Attribute{
															"request_headers": schema.SingleNestedAttribute{
																Description:         "The request headers modifications.",
																MarkdownDescription: "The request headers modifications.",
																Attributes: map[string]schema.Attribute{
																	"pass": schema.BoolAttribute{
																		Description:         "Passes the original request headers to the proxied upstream server. Default is true.",
																		MarkdownDescription: "Passes the original request headers to the proxied upstream server. Default is true.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"set": schema.ListNestedAttribute{
																		Description:         "Allows redefining or appending fields to present request headers passed to the proxied upstream servers.",
																		MarkdownDescription: "Allows redefining or appending fields to present request headers passed to the proxied upstream servers.",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"name": schema.StringAttribute{
																					Description:         "The name of the header.",
																					MarkdownDescription: "The name of the header.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"value": schema.StringAttribute{
																					Description:         "The value of the header.",
																					MarkdownDescription: "The value of the header.",
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

															"response_headers": schema.SingleNestedAttribute{
																Description:         "The response headers modifications.",
																MarkdownDescription: "The response headers modifications.",
																Attributes: map[string]schema.Attribute{
																	"add": schema.ListNestedAttribute{
																		Description:         "Adds headers to the response to the client.",
																		MarkdownDescription: "Adds headers to the response to the client.",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"always": schema.BoolAttribute{
																					Description:         "If set to true, add the header regardless of the response status code**. Default is false.",
																					MarkdownDescription: "If set to true, add the header regardless of the response status code**. Default is false.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"name": schema.StringAttribute{
																					Description:         "The name of the header.",
																					MarkdownDescription: "The name of the header.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"value": schema.StringAttribute{
																					Description:         "The value of the header.",
																					MarkdownDescription: "The value of the header.",
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

																	"hide": schema.ListAttribute{
																		Description:         "The headers that will not be passed* in the response to the client from a proxied upstream server.",
																		MarkdownDescription: "The headers that will not be passed* in the response to the client from a proxied upstream server.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"ignore": schema.ListAttribute{
																		Description:         "Disables processing of certain headers** to the client from a proxied upstream server.",
																		MarkdownDescription: "Disables processing of certain headers** to the client from a proxied upstream server.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"pass": schema.ListAttribute{
																		Description:         "Allows passing the hidden header fields* to the client from a proxied upstream server.",
																		MarkdownDescription: "Allows passing the hidden header fields* to the client from a proxied upstream server.",
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

															"rewrite_path": schema.StringAttribute{
																Description:         "The rewritten URI. If the route path is a regular expression – starts with ~ – the rewritePath can include capture groups with $1-9. For example $1 for the first group, and so on. For more information, check the rewrite example.",
																MarkdownDescription: "The rewritten URI. If the route path is a regular expression – starts with ~ – the rewritePath can include capture groups with $1-9. For example $1 for the first group, and so on. For more information, check the rewrite example.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"upstream": schema.StringAttribute{
																Description:         "The name of the upstream which the requests will be proxied to. The upstream with that name must be defined in the resource.",
																MarkdownDescription: "The name of the upstream which the requests will be proxied to. The upstream with that name must be defined in the resource.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"redirect": schema.SingleNestedAttribute{
														Description:         "Redirects requests to a provided URL.",
														MarkdownDescription: "Redirects requests to a provided URL.",
														Attributes: map[string]schema.Attribute{
															"code": schema.Int64Attribute{
																Description:         "The status code of a redirect. The allowed values are: 301, 302, 307 or 308. The default is 301.",
																MarkdownDescription: "The status code of a redirect. The allowed values are: 301, 302, 307 or 308. The default is 301.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"url": schema.StringAttribute{
																Description:         "The URL to redirect the request to. Supported NGINX variables: $scheme, $http_x_forwarded_proto, $request_uri or $host. Variables must be enclosed in curly braces. For example: ${host}${request_uri}.",
																MarkdownDescription: "The URL to redirect the request to. Supported NGINX variables: $scheme, $http_x_forwarded_proto, $request_uri or $host. Variables must be enclosed in curly braces. For example: ${host}${request_uri}.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"return": schema.SingleNestedAttribute{
														Description:         "Returns a preconfigured response.",
														MarkdownDescription: "Returns a preconfigured response.",
														Attributes: map[string]schema.Attribute{
															"body": schema.StringAttribute{
																Description:         "The body of the response. Supports NGINX variables*. Variables must be enclosed in curly brackets. For example: Request is ${request_uri}n.",
																MarkdownDescription: "The body of the response. Supports NGINX variables*. Variables must be enclosed in curly brackets. For example: Request is ${request_uri}n.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"code": schema.Int64Attribute{
																Description:         "The status code of the response. The allowed values are: 2XX, 4XX or 5XX. The default is 200.",
																MarkdownDescription: "The status code of the response. The allowed values are: 2XX, 4XX or 5XX. The default is 200.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"headers": schema.ListNestedAttribute{
																Description:         "The custom headers of the response.",
																MarkdownDescription: "The custom headers of the response.",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Description:         "The name of the header.",
																			MarkdownDescription: "The name of the header.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"value": schema.StringAttribute{
																			Description:         "The value of the header.",
																			MarkdownDescription: "The value of the header.",
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

															"type": schema.StringAttribute{
																Description:         "The MIME type of the response. The default is text/plain.",
																MarkdownDescription: "The MIME type of the response. The default is text/plain.",
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

											"conditions": schema.ListNestedAttribute{
												Description:         "A list of conditions. Must include at least 1 condition.",
												MarkdownDescription: "A list of conditions. Must include at least 1 condition.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"argument": schema.StringAttribute{
															Description:         "The name of an argument. Must consist of alphanumeric characters or _.",
															MarkdownDescription: "The name of an argument. Must consist of alphanumeric characters or _.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"cookie": schema.StringAttribute{
															Description:         "The name of a cookie. Must consist of alphanumeric characters or _.",
															MarkdownDescription: "The name of a cookie. Must consist of alphanumeric characters or _.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"header": schema.StringAttribute{
															Description:         "The name of a header. Must consist of alphanumeric characters or -.",
															MarkdownDescription: "The name of a header. Must consist of alphanumeric characters or -.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"value": schema.StringAttribute{
															Description:         "The value to match the condition against.",
															MarkdownDescription: "The value to match the condition against.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"variable": schema.StringAttribute{
															Description:         "The name of an NGINX variable. Must start with $.",
															MarkdownDescription: "The name of an NGINX variable. Must start with $.",
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

											"splits": schema.ListNestedAttribute{
												Description:         "The splits configuration for traffic splitting. Must include at least 2 splits.",
												MarkdownDescription: "The splits configuration for traffic splitting. Must include at least 2 splits.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"action": schema.SingleNestedAttribute{
															Description:         "The action to perform for a request.",
															MarkdownDescription: "The action to perform for a request.",
															Attributes: map[string]schema.Attribute{
																"pass": schema.StringAttribute{
																	Description:         "Passes requests to an upstream. The upstream with that name must be defined in the resource.",
																	MarkdownDescription: "Passes requests to an upstream. The upstream with that name must be defined in the resource.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"proxy": schema.SingleNestedAttribute{
																	Description:         "Passes requests to an upstream with the ability to modify the request/response (for example, rewrite the URI or modify the headers).",
																	MarkdownDescription: "Passes requests to an upstream with the ability to modify the request/response (for example, rewrite the URI or modify the headers).",
																	Attributes: map[string]schema.Attribute{
																		"request_headers": schema.SingleNestedAttribute{
																			Description:         "The request headers modifications.",
																			MarkdownDescription: "The request headers modifications.",
																			Attributes: map[string]schema.Attribute{
																				"pass": schema.BoolAttribute{
																					Description:         "Passes the original request headers to the proxied upstream server. Default is true.",
																					MarkdownDescription: "Passes the original request headers to the proxied upstream server. Default is true.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"set": schema.ListNestedAttribute{
																					Description:         "Allows redefining or appending fields to present request headers passed to the proxied upstream servers.",
																					MarkdownDescription: "Allows redefining or appending fields to present request headers passed to the proxied upstream servers.",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"name": schema.StringAttribute{
																								Description:         "The name of the header.",
																								MarkdownDescription: "The name of the header.",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"value": schema.StringAttribute{
																								Description:         "The value of the header.",
																								MarkdownDescription: "The value of the header.",
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

																		"response_headers": schema.SingleNestedAttribute{
																			Description:         "The response headers modifications.",
																			MarkdownDescription: "The response headers modifications.",
																			Attributes: map[string]schema.Attribute{
																				"add": schema.ListNestedAttribute{
																					Description:         "Adds headers to the response to the client.",
																					MarkdownDescription: "Adds headers to the response to the client.",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"always": schema.BoolAttribute{
																								Description:         "If set to true, add the header regardless of the response status code**. Default is false.",
																								MarkdownDescription: "If set to true, add the header regardless of the response status code**. Default is false.",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"name": schema.StringAttribute{
																								Description:         "The name of the header.",
																								MarkdownDescription: "The name of the header.",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"value": schema.StringAttribute{
																								Description:         "The value of the header.",
																								MarkdownDescription: "The value of the header.",
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

																				"hide": schema.ListAttribute{
																					Description:         "The headers that will not be passed* in the response to the client from a proxied upstream server.",
																					MarkdownDescription: "The headers that will not be passed* in the response to the client from a proxied upstream server.",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"ignore": schema.ListAttribute{
																					Description:         "Disables processing of certain headers** to the client from a proxied upstream server.",
																					MarkdownDescription: "Disables processing of certain headers** to the client from a proxied upstream server.",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"pass": schema.ListAttribute{
																					Description:         "Allows passing the hidden header fields* to the client from a proxied upstream server.",
																					MarkdownDescription: "Allows passing the hidden header fields* to the client from a proxied upstream server.",
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

																		"rewrite_path": schema.StringAttribute{
																			Description:         "The rewritten URI. If the route path is a regular expression – starts with ~ – the rewritePath can include capture groups with $1-9. For example $1 for the first group, and so on. For more information, check the rewrite example.",
																			MarkdownDescription: "The rewritten URI. If the route path is a regular expression – starts with ~ – the rewritePath can include capture groups with $1-9. For example $1 for the first group, and so on. For more information, check the rewrite example.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"upstream": schema.StringAttribute{
																			Description:         "The name of the upstream which the requests will be proxied to. The upstream with that name must be defined in the resource.",
																			MarkdownDescription: "The name of the upstream which the requests will be proxied to. The upstream with that name must be defined in the resource.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"redirect": schema.SingleNestedAttribute{
																	Description:         "Redirects requests to a provided URL.",
																	MarkdownDescription: "Redirects requests to a provided URL.",
																	Attributes: map[string]schema.Attribute{
																		"code": schema.Int64Attribute{
																			Description:         "The status code of a redirect. The allowed values are: 301, 302, 307 or 308. The default is 301.",
																			MarkdownDescription: "The status code of a redirect. The allowed values are: 301, 302, 307 or 308. The default is 301.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"url": schema.StringAttribute{
																			Description:         "The URL to redirect the request to. Supported NGINX variables: $scheme, $http_x_forwarded_proto, $request_uri or $host. Variables must be enclosed in curly braces. For example: ${host}${request_uri}.",
																			MarkdownDescription: "The URL to redirect the request to. Supported NGINX variables: $scheme, $http_x_forwarded_proto, $request_uri or $host. Variables must be enclosed in curly braces. For example: ${host}${request_uri}.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"return": schema.SingleNestedAttribute{
																	Description:         "Returns a preconfigured response.",
																	MarkdownDescription: "Returns a preconfigured response.",
																	Attributes: map[string]schema.Attribute{
																		"body": schema.StringAttribute{
																			Description:         "The body of the response. Supports NGINX variables*. Variables must be enclosed in curly brackets. For example: Request is ${request_uri}n.",
																			MarkdownDescription: "The body of the response. Supports NGINX variables*. Variables must be enclosed in curly brackets. For example: Request is ${request_uri}n.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"code": schema.Int64Attribute{
																			Description:         "The status code of the response. The allowed values are: 2XX, 4XX or 5XX. The default is 200.",
																			MarkdownDescription: "The status code of the response. The allowed values are: 2XX, 4XX or 5XX. The default is 200.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"headers": schema.ListNestedAttribute{
																			Description:         "The custom headers of the response.",
																			MarkdownDescription: "The custom headers of the response.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "The name of the header.",
																						MarkdownDescription: "The name of the header.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"value": schema.StringAttribute{
																						Description:         "The value of the header.",
																						MarkdownDescription: "The value of the header.",
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

																		"type": schema.StringAttribute{
																			Description:         "The MIME type of the response. The default is text/plain.",
																			MarkdownDescription: "The MIME type of the response. The default is text/plain.",
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

														"weight": schema.Int64Attribute{
															Description:         "The weight of an action. Must fall into the range 0..100. The sum of the weights of all splits must be equal to 100.",
															MarkdownDescription: "The weight of an action. Must fall into the range 0..100. The sum of the weights of all splits must be equal to 100.",
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
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"path": schema.StringAttribute{
									Description:         "The path of the route. NGINX will match it against the URI of a request. Possible values are: a prefix ( / , /path ), an exact match ( =/exact/match ), a case insensitive regular expression ( ~*^/Bar.*.jpg ) or a case sensitive regular expression ( ~^/foo.*.jpg ). In the case of a prefix (must start with / ) or an exact match (must start with = ), the path must not include any whitespace characters, { , } or ;. In the case of the regex matches, all double quotes ' must be escaped and the match can’t end in an unescaped backslash . The path must be unique among the paths of all routes of the VirtualServer. Check the location directive for more information.",
									MarkdownDescription: "The path of the route. NGINX will match it against the URI of a request. Possible values are: a prefix ( / , /path ), an exact match ( =/exact/match ), a case insensitive regular expression ( ~*^/Bar.*.jpg ) or a case sensitive regular expression ( ~^/foo.*.jpg ). In the case of a prefix (must start with / ) or an exact match (must start with = ), the path must not include any whitespace characters, { , } or ;. In the case of the regex matches, all double quotes ' must be escaped and the match can’t end in an unescaped backslash . The path must be unique among the paths of all routes of the VirtualServer. Check the location directive for more information.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"policies": schema.ListNestedAttribute{
									Description:         "A list of policies. The policies override the policies of the same type defined in the spec of the VirtualServer.",
									MarkdownDescription: "A list of policies. The policies override the policies of the same type defined in the spec of the VirtualServer.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "The name of a policy. If the policy doesn’t exist or invalid, NGINX will respond with an error response with the 500 status code.",
												MarkdownDescription: "The name of a policy. If the policy doesn’t exist or invalid, NGINX will respond with an error response with the 500 status code.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "The namespace of a policy. If not specified, the namespace of the VirtualServer resource is used.",
												MarkdownDescription: "The namespace of a policy. If not specified, the namespace of the VirtualServer resource is used.",
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

								"route": schema.StringAttribute{
									Description:         "The name of a VirtualServerRoute resource that defines this route. If the VirtualServerRoute belongs to a different namespace than the VirtualServer, you need to include the namespace. For example, tea-namespace/tea.",
									MarkdownDescription: "The name of a VirtualServerRoute resource that defines this route. If the VirtualServerRoute belongs to a different namespace than the VirtualServer, you need to include the namespace. For example, tea-namespace/tea.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"splits": schema.ListNestedAttribute{
									Description:         "The default splits configuration for traffic splitting. Must include at least 2 splits.",
									MarkdownDescription: "The default splits configuration for traffic splitting. Must include at least 2 splits.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"action": schema.SingleNestedAttribute{
												Description:         "The action to perform for a request.",
												MarkdownDescription: "The action to perform for a request.",
												Attributes: map[string]schema.Attribute{
													"pass": schema.StringAttribute{
														Description:         "Passes requests to an upstream. The upstream with that name must be defined in the resource.",
														MarkdownDescription: "Passes requests to an upstream. The upstream with that name must be defined in the resource.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"proxy": schema.SingleNestedAttribute{
														Description:         "Passes requests to an upstream with the ability to modify the request/response (for example, rewrite the URI or modify the headers).",
														MarkdownDescription: "Passes requests to an upstream with the ability to modify the request/response (for example, rewrite the URI or modify the headers).",
														Attributes: map[string]schema.Attribute{
															"request_headers": schema.SingleNestedAttribute{
																Description:         "The request headers modifications.",
																MarkdownDescription: "The request headers modifications.",
																Attributes: map[string]schema.Attribute{
																	"pass": schema.BoolAttribute{
																		Description:         "Passes the original request headers to the proxied upstream server. Default is true.",
																		MarkdownDescription: "Passes the original request headers to the proxied upstream server. Default is true.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"set": schema.ListNestedAttribute{
																		Description:         "Allows redefining or appending fields to present request headers passed to the proxied upstream servers.",
																		MarkdownDescription: "Allows redefining or appending fields to present request headers passed to the proxied upstream servers.",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"name": schema.StringAttribute{
																					Description:         "The name of the header.",
																					MarkdownDescription: "The name of the header.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"value": schema.StringAttribute{
																					Description:         "The value of the header.",
																					MarkdownDescription: "The value of the header.",
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

															"response_headers": schema.SingleNestedAttribute{
																Description:         "The response headers modifications.",
																MarkdownDescription: "The response headers modifications.",
																Attributes: map[string]schema.Attribute{
																	"add": schema.ListNestedAttribute{
																		Description:         "Adds headers to the response to the client.",
																		MarkdownDescription: "Adds headers to the response to the client.",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"always": schema.BoolAttribute{
																					Description:         "If set to true, add the header regardless of the response status code**. Default is false.",
																					MarkdownDescription: "If set to true, add the header regardless of the response status code**. Default is false.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"name": schema.StringAttribute{
																					Description:         "The name of the header.",
																					MarkdownDescription: "The name of the header.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"value": schema.StringAttribute{
																					Description:         "The value of the header.",
																					MarkdownDescription: "The value of the header.",
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

																	"hide": schema.ListAttribute{
																		Description:         "The headers that will not be passed* in the response to the client from a proxied upstream server.",
																		MarkdownDescription: "The headers that will not be passed* in the response to the client from a proxied upstream server.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"ignore": schema.ListAttribute{
																		Description:         "Disables processing of certain headers** to the client from a proxied upstream server.",
																		MarkdownDescription: "Disables processing of certain headers** to the client from a proxied upstream server.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"pass": schema.ListAttribute{
																		Description:         "Allows passing the hidden header fields* to the client from a proxied upstream server.",
																		MarkdownDescription: "Allows passing the hidden header fields* to the client from a proxied upstream server.",
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

															"rewrite_path": schema.StringAttribute{
																Description:         "The rewritten URI. If the route path is a regular expression – starts with ~ – the rewritePath can include capture groups with $1-9. For example $1 for the first group, and so on. For more information, check the rewrite example.",
																MarkdownDescription: "The rewritten URI. If the route path is a regular expression – starts with ~ – the rewritePath can include capture groups with $1-9. For example $1 for the first group, and so on. For more information, check the rewrite example.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"upstream": schema.StringAttribute{
																Description:         "The name of the upstream which the requests will be proxied to. The upstream with that name must be defined in the resource.",
																MarkdownDescription: "The name of the upstream which the requests will be proxied to. The upstream with that name must be defined in the resource.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"redirect": schema.SingleNestedAttribute{
														Description:         "Redirects requests to a provided URL.",
														MarkdownDescription: "Redirects requests to a provided URL.",
														Attributes: map[string]schema.Attribute{
															"code": schema.Int64Attribute{
																Description:         "The status code of a redirect. The allowed values are: 301, 302, 307 or 308. The default is 301.",
																MarkdownDescription: "The status code of a redirect. The allowed values are: 301, 302, 307 or 308. The default is 301.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"url": schema.StringAttribute{
																Description:         "The URL to redirect the request to. Supported NGINX variables: $scheme, $http_x_forwarded_proto, $request_uri or $host. Variables must be enclosed in curly braces. For example: ${host}${request_uri}.",
																MarkdownDescription: "The URL to redirect the request to. Supported NGINX variables: $scheme, $http_x_forwarded_proto, $request_uri or $host. Variables must be enclosed in curly braces. For example: ${host}${request_uri}.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"return": schema.SingleNestedAttribute{
														Description:         "Returns a preconfigured response.",
														MarkdownDescription: "Returns a preconfigured response.",
														Attributes: map[string]schema.Attribute{
															"body": schema.StringAttribute{
																Description:         "The body of the response. Supports NGINX variables*. Variables must be enclosed in curly brackets. For example: Request is ${request_uri}n.",
																MarkdownDescription: "The body of the response. Supports NGINX variables*. Variables must be enclosed in curly brackets. For example: Request is ${request_uri}n.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"code": schema.Int64Attribute{
																Description:         "The status code of the response. The allowed values are: 2XX, 4XX or 5XX. The default is 200.",
																MarkdownDescription: "The status code of the response. The allowed values are: 2XX, 4XX or 5XX. The default is 200.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"headers": schema.ListNestedAttribute{
																Description:         "The custom headers of the response.",
																MarkdownDescription: "The custom headers of the response.",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Description:         "The name of the header.",
																			MarkdownDescription: "The name of the header.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"value": schema.StringAttribute{
																			Description:         "The value of the header.",
																			MarkdownDescription: "The value of the header.",
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

															"type": schema.StringAttribute{
																Description:         "The MIME type of the response. The default is text/plain.",
																MarkdownDescription: "The MIME type of the response. The default is text/plain.",
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

											"weight": schema.Int64Attribute{
												Description:         "The weight of an action. Must fall into the range 0..100. The sum of the weights of all splits must be equal to 100.",
												MarkdownDescription: "The weight of an action. Must fall into the range 0..100. The sum of the weights of all splits must be equal to 100.",
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

								"buffer_size": schema.StringAttribute{
									Description:         "Sets the size of the buffer used for reading the first part of a response received from the upstream server. The default is set in the proxy-buffer-size ConfigMap key.",
									MarkdownDescription: "Sets the size of the buffer used for reading the first part of a response received from the upstream server. The default is set in the proxy-buffer-size ConfigMap key.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"buffering": schema.BoolAttribute{
									Description:         "Enables buffering of responses from the upstream server. The default is set in the proxy-buffering ConfigMap key.",
									MarkdownDescription: "Enables buffering of responses from the upstream server. The default is set in the proxy-buffering ConfigMap key.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"buffers": schema.SingleNestedAttribute{
									Description:         "Configures the buffers used for reading a response from the upstream server for a single connection.",
									MarkdownDescription: "Configures the buffers used for reading a response from the upstream server for a single connection.",
									Attributes: map[string]schema.Attribute{
										"number": schema.Int64Attribute{
											Description:         "Configures the number of buffers. The default is set in the proxy-buffers ConfigMap key.",
											MarkdownDescription: "Configures the number of buffers. The default is set in the proxy-buffers ConfigMap key.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"size": schema.StringAttribute{
											Description:         "Configures the size of a buffer. The default is set in the proxy-buffers ConfigMap key.",
											MarkdownDescription: "Configures the size of a buffer. The default is set in the proxy-buffers ConfigMap key.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"client_max_body_size": schema.StringAttribute{
									Description:         "Sets the maximum allowed size of the client request body. The default is set in the client-max-body-size ConfigMap key.",
									MarkdownDescription: "Sets the maximum allowed size of the client request body. The default is set in the client-max-body-size ConfigMap key.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"connect_timeout": schema.StringAttribute{
									Description:         "The timeout for establishing a connection with an upstream server. The default is specified in the proxy-connect-timeout ConfigMap key.",
									MarkdownDescription: "The timeout for establishing a connection with an upstream server. The default is specified in the proxy-connect-timeout ConfigMap key.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"fail_timeout": schema.StringAttribute{
									Description:         "The time during which the specified number of unsuccessful attempts to communicate with an upstream server should happen to consider the server unavailable. The default is set in the fail-timeout ConfigMap key.",
									MarkdownDescription: "The time during which the specified number of unsuccessful attempts to communicate with an upstream server should happen to consider the server unavailable. The default is set in the fail-timeout ConfigMap key.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"health_check": schema.SingleNestedAttribute{
									Description:         "The health check configuration for the Upstream. Note: this feature is supported only in NGINX Plus.",
									MarkdownDescription: "The health check configuration for the Upstream. Note: this feature is supported only in NGINX Plus.",
									Attributes: map[string]schema.Attribute{
										"connect_timeout": schema.StringAttribute{
											Description:         "The timeout for establishing a connection with an upstream server. By default, the connect-timeout of the upstream is used.",
											MarkdownDescription: "The timeout for establishing a connection with an upstream server. By default, the connect-timeout of the upstream is used.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

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

										"grpc_service": schema.StringAttribute{
											Description:         "The gRPC service to be monitored on the upstream server. Only valid on gRPC type upstreams.",
											MarkdownDescription: "The gRPC service to be monitored on the upstream server. Only valid on gRPC type upstreams.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"grpc_status": schema.Int64Attribute{
											Description:         "The expected gRPC status code of the upstream server response to the Check method. Configure this field only if your gRPC services do not implement the gRPC health checking protocol. For example, configure 12 if the upstream server responds with 12 (UNIMPLEMENTED) status code. Only valid on gRPC type upstreams.",
											MarkdownDescription: "The expected gRPC status code of the upstream server response to the Check method. Configure this field only if your gRPC services do not implement the gRPC health checking protocol. For example, configure 12 if the upstream server responds with 12 (UNIMPLEMENTED) status code. Only valid on gRPC type upstreams.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"headers": schema.ListNestedAttribute{
											Description:         "The request headers used for health check requests. NGINX Plus always sets the Host, User-Agent and Connection headers for health check requests.",
											MarkdownDescription: "The request headers used for health check requests. NGINX Plus always sets the Host, User-Agent and Connection headers for health check requests.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "The name of the header.",
														MarkdownDescription: "The name of the header.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"value": schema.StringAttribute{
														Description:         "The value of the header.",
														MarkdownDescription: "The value of the header.",
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

										"keepalive_time": schema.StringAttribute{
											Description:         "Enables keepalive connections for health checks and specifies the time during which requests can be processed through one keepalive connection. The default is 60s.",
											MarkdownDescription: "Enables keepalive connections for health checks and specifies the time during which requests can be processed through one keepalive connection. The default is 60s.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"mandatory": schema.BoolAttribute{
											Description:         "Require every newly added server to pass all configured health checks before NGINX Plus sends traffic to it. If this is not specified, or is set to false, the server will be initially considered healthy. When combined with slow-start, it gives a new server more time to connect to databases and “warm up” before being asked to handle their full share of traffic.",
											MarkdownDescription: "Require every newly added server to pass all configured health checks before NGINX Plus sends traffic to it. If this is not specified, or is set to false, the server will be initially considered healthy. When combined with slow-start, it gives a new server more time to connect to databases and “warm up” before being asked to handle their full share of traffic.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"passes": schema.Int64Attribute{
											Description:         "The number of consecutive passed health checks of a particular upstream server after which the server will be considered healthy. The default is 1.",
											MarkdownDescription: "The number of consecutive passed health checks of a particular upstream server after which the server will be considered healthy. The default is 1.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"path": schema.StringAttribute{
											Description:         "The path used for health check requests. The default is /. This is not configurable for gRPC type upstreams.",
											MarkdownDescription: "The path used for health check requests. The default is /. This is not configurable for gRPC type upstreams.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"persistent": schema.BoolAttribute{
											Description:         "Set the initial “up” state for a server after reload if the server was considered healthy before reload. Enabling persistent requires that the mandatory parameter is also set to true.",
											MarkdownDescription: "Set the initial “up” state for a server after reload if the server was considered healthy before reload. Enabling persistent requires that the mandatory parameter is also set to true.",
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

										"read_timeout": schema.StringAttribute{
											Description:         "The timeout for reading a response from an upstream server. By default, the read-timeout of the upstream is used.",
											MarkdownDescription: "The timeout for reading a response from an upstream server. By default, the read-timeout of the upstream is used.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"send_timeout": schema.StringAttribute{
											Description:         "The timeout for transmitting a request to an upstream server. By default, the send-timeout of the upstream is used.",
											MarkdownDescription: "The timeout for transmitting a request to an upstream server. By default, the send-timeout of the upstream is used.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"status_match": schema.StringAttribute{
											Description:         "The expected response status codes of a health check. By default, the response should have status code 2xx or 3xx. Examples: '200', '! 500', '301-303 307'. This not supported for gRPC type upstreams.",
											MarkdownDescription: "The expected response status codes of a health check. By default, the response should have status code 2xx or 3xx. Examples: '200', '! 500', '301-303 307'. This not supported for gRPC type upstreams.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tls": schema.SingleNestedAttribute{
											Description:         "The TLS configuration used for health check requests. By default, the tls field of the upstream is used.",
											MarkdownDescription: "The TLS configuration used for health check requests. By default, the tls field of the upstream is used.",
											Attributes: map[string]schema.Attribute{
												"enable": schema.BoolAttribute{
													Description:         "Enables HTTPS for requests to upstream servers. The default is False , meaning that HTTP will be used. Note: by default, NGINX will not verify the upstream server certificate. To enable the verification, configure an EgressMTLS Policy.",
													MarkdownDescription: "Enables HTTPS for requests to upstream servers. The default is False , meaning that HTTP will be used. Note: by default, NGINX will not verify the upstream server certificate. To enable the verification, configure an EgressMTLS Policy.",
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

								"keepalive": schema.Int64Attribute{
									Description:         "Configures the cache for connections to upstream servers. The value 0 disables the cache. The default is set in the keepalive ConfigMap key.",
									MarkdownDescription: "Configures the cache for connections to upstream servers. The value 0 disables the cache. The default is set in the keepalive ConfigMap key.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"lb_method": schema.StringAttribute{
									Description:         "The load balancing method. To use the round-robin method, specify round_robin. The default is specified in the lb-method ConfigMap key.",
									MarkdownDescription: "The load balancing method. To use the round-robin method, specify round_robin. The default is specified in the lb-method ConfigMap key.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"max_conns": schema.Int64Attribute{
									Description:         "The maximum number of simultaneous active connections to an upstream server. By default there is no limit. Note: if keepalive connections are enabled, the total number of active and idle keepalive connections to an upstream server may exceed the max_conns value.",
									MarkdownDescription: "The maximum number of simultaneous active connections to an upstream server. By default there is no limit. Note: if keepalive connections are enabled, the total number of active and idle keepalive connections to an upstream server may exceed the max_conns value.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"max_fails": schema.Int64Attribute{
									Description:         "The number of unsuccessful attempts to communicate with an upstream server that should happen in the duration set by the fail-timeout to consider the server unavailable. The default is set in the max-fails ConfigMap key.",
									MarkdownDescription: "The number of unsuccessful attempts to communicate with an upstream server that should happen in the duration set by the fail-timeout to consider the server unavailable. The default is set in the max-fails ConfigMap key.",
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

								"next_upstream": schema.StringAttribute{
									Description:         "Specifies in which cases a request should be passed to the next upstream server. The default is error timeout.",
									MarkdownDescription: "Specifies in which cases a request should be passed to the next upstream server. The default is error timeout.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"next_upstream_timeout": schema.StringAttribute{
									Description:         "The time during which a request can be passed to the next upstream server. The 0 value turns off the time limit. The default is 0.",
									MarkdownDescription: "The time during which a request can be passed to the next upstream server. The 0 value turns off the time limit. The default is 0.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"next_upstream_tries": schema.Int64Attribute{
									Description:         "The number of possible tries for passing a request to the next upstream server. The 0 value turns off this limit. The default is 0.",
									MarkdownDescription: "The number of possible tries for passing a request to the next upstream server. The 0 value turns off this limit. The default is 0.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ntlm": schema.BoolAttribute{
									Description:         "Allows proxying requests with NTLM Authentication. In order for NTLM authentication to work, it is necessary to enable keepalive connections to upstream servers using the keepalive field. Note: this feature is supported only in NGINX Plus.",
									MarkdownDescription: "Allows proxying requests with NTLM Authentication. In order for NTLM authentication to work, it is necessary to enable keepalive connections to upstream servers using the keepalive field. Note: this feature is supported only in NGINX Plus.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"port": schema.Int64Attribute{
									Description:         "The port of the service. If the service doesn’t define that port, NGINX will assume the service has zero endpoints and return a 502 response for requests for this upstream. The port must fall into the range 1..65535.",
									MarkdownDescription: "The port of the service. If the service doesn’t define that port, NGINX will assume the service has zero endpoints and return a 502 response for requests for this upstream. The port must fall into the range 1..65535.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"queue": schema.SingleNestedAttribute{
									Description:         "Configures a queue for an upstream. A client request will be placed into the queue if an upstream server cannot be selected immediately while processing the request. By default, no queue is configured. Note: this feature is supported only in NGINX Plus.",
									MarkdownDescription: "Configures a queue for an upstream. A client request will be placed into the queue if an upstream server cannot be selected immediately while processing the request. By default, no queue is configured. Note: this feature is supported only in NGINX Plus.",
									Attributes: map[string]schema.Attribute{
										"size": schema.Int64Attribute{
											Description:         "The size of the queue.",
											MarkdownDescription: "The size of the queue.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"timeout": schema.StringAttribute{
											Description:         "The timeout of the queue. A request cannot be queued for a period longer than the timeout. The default is 60s.",
											MarkdownDescription: "The timeout of the queue. A request cannot be queued for a period longer than the timeout. The default is 60s.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"read_timeout": schema.StringAttribute{
									Description:         "The timeout for reading a response from an upstream server. The default is specified in the proxy-read-timeout ConfigMap key.",
									MarkdownDescription: "The timeout for reading a response from an upstream server. The default is specified in the proxy-read-timeout ConfigMap key.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"send_timeout": schema.StringAttribute{
									Description:         "The timeout for transmitting a request to an upstream server. The default is specified in the proxy-send-timeout ConfigMap key.",
									MarkdownDescription: "The timeout for transmitting a request to an upstream server. The default is specified in the proxy-send-timeout ConfigMap key.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"service": schema.StringAttribute{
									Description:         "The name of a service. The service must belong to the same namespace as the resource. If the service doesn’t exist, NGINX will assume the service has zero endpoints and return a 502 response for requests for this upstream. For NGINX Plus only, services of type ExternalName are also supported .",
									MarkdownDescription: "The name of a service. The service must belong to the same namespace as the resource. If the service doesn’t exist, NGINX will assume the service has zero endpoints and return a 502 response for requests for this upstream. For NGINX Plus only, services of type ExternalName are also supported .",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"session_cookie": schema.SingleNestedAttribute{
									Description:         "The SessionCookie field configures session persistence which allows requests from the same client to be passed to the same upstream server. The information about the designated upstream server is passed in a session cookie generated by NGINX Plus.",
									MarkdownDescription: "The SessionCookie field configures session persistence which allows requests from the same client to be passed to the same upstream server. The information about the designated upstream server is passed in a session cookie generated by NGINX Plus.",
									Attributes: map[string]schema.Attribute{
										"domain": schema.StringAttribute{
											Description:         "The domain for which the cookie is set.",
											MarkdownDescription: "The domain for which the cookie is set.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"enable": schema.BoolAttribute{
											Description:         "Enables session persistence with a session cookie for an upstream server. The default is false.",
											MarkdownDescription: "Enables session persistence with a session cookie for an upstream server. The default is false.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"expires": schema.StringAttribute{
											Description:         "The time for which a browser should keep the cookie. Can be set to the special value max, which will cause the cookie to expire on 31 Dec 2037 23:55:55 GMT.",
											MarkdownDescription: "The time for which a browser should keep the cookie. Can be set to the special value max, which will cause the cookie to expire on 31 Dec 2037 23:55:55 GMT.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"http_only": schema.BoolAttribute{
											Description:         "Adds the HttpOnly attribute to the cookie.",
											MarkdownDescription: "Adds the HttpOnly attribute to the cookie.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "The name of the cookie.",
											MarkdownDescription: "The name of the cookie.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"path": schema.StringAttribute{
											Description:         "The path for which the cookie is set.",
											MarkdownDescription: "The path for which the cookie is set.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"samesite": schema.StringAttribute{
											Description:         "Adds the SameSite attribute to the cookie. The allowed values are: strict, lax, none",
											MarkdownDescription: "Adds the SameSite attribute to the cookie. The allowed values are: strict, lax, none",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secure": schema.BoolAttribute{
											Description:         "Adds the Secure attribute to the cookie.",
											MarkdownDescription: "Adds the Secure attribute to the cookie.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"slow_start": schema.StringAttribute{
									Description:         "The slow start allows an upstream server to gradually recover its weight from 0 to its nominal value after it has been recovered or became available or when the server becomes available after a period of time it was considered unavailable. By default, the slow start is disabled. Note: The parameter cannot be used along with the random, hash or ip_hash load balancing methods and will be ignored.",
									MarkdownDescription: "The slow start allows an upstream server to gradually recover its weight from 0 to its nominal value after it has been recovered or became available or when the server becomes available after a period of time it was considered unavailable. By default, the slow start is disabled. Note: The parameter cannot be used along with the random, hash or ip_hash load balancing methods and will be ignored.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"subselector": schema.MapAttribute{
									Description:         "Selects the pods within the service using label keys and values. By default, all pods of the service are selected. Note: the specified labels are expected to be present in the pods when they are created. If the pod labels are updated, NGINX Ingress Controller will not see that change until the number of the pods is changed.",
									MarkdownDescription: "Selects the pods within the service using label keys and values. By default, all pods of the service are selected. Note: the specified labels are expected to be present in the pods when they are created. If the pod labels are updated, NGINX Ingress Controller will not see that change until the number of the pods is changed.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"tls": schema.SingleNestedAttribute{
									Description:         "The TLS configuration for the Upstream.",
									MarkdownDescription: "The TLS configuration for the Upstream.",
									Attributes: map[string]schema.Attribute{
										"enable": schema.BoolAttribute{
											Description:         "Enables HTTPS for requests to upstream servers. The default is False , meaning that HTTP will be used. Note: by default, NGINX will not verify the upstream server certificate. To enable the verification, configure an EgressMTLS Policy.",
											MarkdownDescription: "Enables HTTPS for requests to upstream servers. The default is False , meaning that HTTP will be used. Note: by default, NGINX will not verify the upstream server certificate. To enable the verification, configure an EgressMTLS Policy.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"type": schema.StringAttribute{
									Description:         "The type of the upstream. Supported values are http and grpc. The default is http. For gRPC, it is necessary to enable HTTP/2 in the ConfigMap and configure TLS termination in the VirtualServer.",
									MarkdownDescription: "The type of the upstream. Supported values are http and grpc. The default is http. For gRPC, it is necessary to enable HTTP/2 in the ConfigMap and configure TLS termination in the VirtualServer.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"use_cluster_ip": schema.BoolAttribute{
									Description:         "Enables using the Cluster IP and port of the service instead of the default behavior of using the IP and port of the pods. When this field is enabled, the fields that configure NGINX behavior related to multiple upstream servers (like lb-method and next-upstream) will have no effect, as NGINX Ingress Controller will configure NGINX with only one upstream server that will match the service Cluster IP.",
									MarkdownDescription: "Enables using the Cluster IP and port of the service instead of the default behavior of using the IP and port of the pods. When this field is enabled, the fields that configure NGINX behavior related to multiple upstream servers (like lb-method and next-upstream) will have no effect, as NGINX Ingress Controller will configure NGINX with only one upstream server that will match the service Cluster IP.",
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

func (r *K8SNginxOrgVirtualServerRouteV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_k8s_nginx_org_virtual_server_route_v1_manifest")

	var model K8SNginxOrgVirtualServerRouteV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("k8s.nginx.org/v1")
	model.Kind = pointer.String("VirtualServerRoute")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
