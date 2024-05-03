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
	_ datasource.DataSource = &K8SNginxOrgVirtualServerV1Manifest{}
)

func NewK8SNginxOrgVirtualServerV1Manifest() datasource.DataSource {
	return &K8SNginxOrgVirtualServerV1Manifest{}
}

type K8SNginxOrgVirtualServerV1Manifest struct{}

type K8SNginxOrgVirtualServerV1ManifestData struct {
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
		Dos         *string `tfsdk:"dos" json:"dos,omitempty"`
		ExternalDNS *struct {
			Enable           *bool              `tfsdk:"enable" json:"enable,omitempty"`
			Labels           *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			ProviderSpecific *[]struct {
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"provider_specific" json:"providerSpecific,omitempty"`
			RecordTTL  *int64  `tfsdk:"record_ttl" json:"recordTTL,omitempty"`
			RecordType *string `tfsdk:"record_type" json:"recordType,omitempty"`
		} `tfsdk:"external_dns" json:"externalDNS,omitempty"`
		Gunzip           *bool   `tfsdk:"gunzip" json:"gunzip,omitempty"`
		Host             *string `tfsdk:"host" json:"host,omitempty"`
		Http_snippets    *string `tfsdk:"http_snippets" json:"http-snippets,omitempty"`
		IngressClassName *string `tfsdk:"ingress_class_name" json:"ingressClassName,omitempty"`
		InternalRoute    *bool   `tfsdk:"internal_route" json:"internalRoute,omitempty"`
		Listener         *struct {
			Http  *string `tfsdk:"http" json:"http,omitempty"`
			Https *string `tfsdk:"https" json:"https,omitempty"`
		} `tfsdk:"listener" json:"listener,omitempty"`
		Policies *[]struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"policies" json:"policies,omitempty"`
		Routes *[]struct {
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
		} `tfsdk:"routes" json:"routes,omitempty"`
		Server_snippets *string `tfsdk:"server_snippets" json:"server-snippets,omitempty"`
		Tls             *struct {
			Cert_manager *struct {
				Cluster_issuer  *string `tfsdk:"cluster_issuer" json:"cluster-issuer,omitempty"`
				Common_name     *string `tfsdk:"common_name" json:"common-name,omitempty"`
				Duration        *string `tfsdk:"duration" json:"duration,omitempty"`
				Issue_temp_cert *bool   `tfsdk:"issue_temp_cert" json:"issue-temp-cert,omitempty"`
				Issuer          *string `tfsdk:"issuer" json:"issuer,omitempty"`
				Issuer_group    *string `tfsdk:"issuer_group" json:"issuer-group,omitempty"`
				Issuer_kind     *string `tfsdk:"issuer_kind" json:"issuer-kind,omitempty"`
				Renew_before    *string `tfsdk:"renew_before" json:"renew-before,omitempty"`
				Usages          *string `tfsdk:"usages" json:"usages,omitempty"`
			} `tfsdk:"cert_manager" json:"cert-manager,omitempty"`
			Redirect *struct {
				BasedOn *string `tfsdk:"based_on" json:"basedOn,omitempty"`
				Code    *int64  `tfsdk:"code" json:"code,omitempty"`
				Enable  *bool   `tfsdk:"enable" json:"enable,omitempty"`
			} `tfsdk:"redirect" json:"redirect,omitempty"`
			Secret *string `tfsdk:"secret" json:"secret,omitempty"`
		} `tfsdk:"tls" json:"tls,omitempty"`
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

func (r *K8SNginxOrgVirtualServerV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_k8s_nginx_org_virtual_server_v1_manifest"
}

func (r *K8SNginxOrgVirtualServerV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "VirtualServer defines the VirtualServer resource.",
		MarkdownDescription: "VirtualServer defines the VirtualServer resource.",
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
				Description:         "VirtualServerSpec is the spec of the VirtualServer resource.",
				MarkdownDescription: "VirtualServerSpec is the spec of the VirtualServer resource.",
				Attributes: map[string]schema.Attribute{
					"dos": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"external_dns": schema.SingleNestedAttribute{
						Description:         "ExternalDNS defines externaldns sub-resource of a virtual server.",
						MarkdownDescription: "ExternalDNS defines externaldns sub-resource of a virtual server.",
						Attributes: map[string]schema.Attribute{
							"enable": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"labels": schema.MapAttribute{
								Description:         "Labels stores labels defined for the Endpoint",
								MarkdownDescription: "Labels stores labels defined for the Endpoint",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"provider_specific": schema.ListNestedAttribute{
								Description:         "ProviderSpecific stores provider specific config",
								MarkdownDescription: "ProviderSpecific stores provider specific config",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name of the property",
											MarkdownDescription: "Name of the property",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "Value of the property",
											MarkdownDescription: "Value of the property",
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

							"record_ttl": schema.Int64Attribute{
								Description:         "TTL for the record",
								MarkdownDescription: "TTL for the record",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"record_type": schema.StringAttribute{
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

					"gunzip": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"host": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"http_snippets": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ingress_class_name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"internal_route": schema.BoolAttribute{
						Description:         "InternalRoute allows for the configuration of internal routing.",
						MarkdownDescription: "InternalRoute allows for the configuration of internal routing.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"listener": schema.SingleNestedAttribute{
						Description:         "VirtualServerListener references a custom http and/or https listener defined in GlobalConfiguration.",
						MarkdownDescription: "VirtualServerListener references a custom http and/or https listener defined in GlobalConfiguration.",
						Attributes: map[string]schema.Attribute{
							"http": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"https": schema.StringAttribute{
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

					"policies": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"routes": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"action": schema.SingleNestedAttribute{
									Description:         "Action defines an action.",
									MarkdownDescription: "Action defines an action.",
									Attributes: map[string]schema.Attribute{
										"pass": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"proxy": schema.SingleNestedAttribute{
											Description:         "ActionProxy defines a proxy in an Action.",
											MarkdownDescription: "ActionProxy defines a proxy in an Action.",
											Attributes: map[string]schema.Attribute{
												"request_headers": schema.SingleNestedAttribute{
													Description:         "ProxyRequestHeaders defines the request headers manipulation in an ActionProxy.",
													MarkdownDescription: "ProxyRequestHeaders defines the request headers manipulation in an ActionProxy.",
													Attributes: map[string]schema.Attribute{
														"pass": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"set": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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
													Description:         "ProxyResponseHeaders defines the response headers manipulation in an ActionProxy.",
													MarkdownDescription: "ProxyResponseHeaders defines the response headers manipulation in an ActionProxy.",
													Attributes: map[string]schema.Attribute{
														"add": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"always": schema.BoolAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"ignore": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"pass": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
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
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"upstream": schema.StringAttribute{
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

										"redirect": schema.SingleNestedAttribute{
											Description:         "ActionRedirect defines a redirect in an Action.",
											MarkdownDescription: "ActionRedirect defines a redirect in an Action.",
											Attributes: map[string]schema.Attribute{
												"code": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"url": schema.StringAttribute{
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

										"return": schema.SingleNestedAttribute{
											Description:         "ActionReturn defines a return in an Action.",
											MarkdownDescription: "ActionReturn defines a return in an Action.",
											Attributes: map[string]schema.Attribute{
												"body": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"code": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"headers": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"value": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
									Required: false,
									Optional: true,
									Computed: false,
								},

								"dos": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"error_pages": schema.ListNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"codes": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"redirect": schema.SingleNestedAttribute{
												Description:         "ErrorPageRedirect defines a redirect for an ErrorPage.",
												MarkdownDescription: "ErrorPageRedirect defines a redirect for an ErrorPage.",
												Attributes: map[string]schema.Attribute{
													"code": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"url": schema.StringAttribute{
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

											"return": schema.SingleNestedAttribute{
												Description:         "ErrorPageReturn defines a return for an ErrorPage.",
												MarkdownDescription: "ErrorPageReturn defines a return for an ErrorPage.",
												Attributes: map[string]schema.Attribute{
													"body": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"code": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"headers": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"value": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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

								"location_snippets": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"matches": schema.ListNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"action": schema.SingleNestedAttribute{
												Description:         "Action defines an action.",
												MarkdownDescription: "Action defines an action.",
												Attributes: map[string]schema.Attribute{
													"pass": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"proxy": schema.SingleNestedAttribute{
														Description:         "ActionProxy defines a proxy in an Action.",
														MarkdownDescription: "ActionProxy defines a proxy in an Action.",
														Attributes: map[string]schema.Attribute{
															"request_headers": schema.SingleNestedAttribute{
																Description:         "ProxyRequestHeaders defines the request headers manipulation in an ActionProxy.",
																MarkdownDescription: "ProxyRequestHeaders defines the request headers manipulation in an ActionProxy.",
																Attributes: map[string]schema.Attribute{
																	"pass": schema.BoolAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"set": schema.ListNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"name": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"value": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
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
																Description:         "ProxyResponseHeaders defines the response headers manipulation in an ActionProxy.",
																MarkdownDescription: "ProxyResponseHeaders defines the response headers manipulation in an ActionProxy.",
																Attributes: map[string]schema.Attribute{
																	"add": schema.ListNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"always": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"name": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"value": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
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
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"ignore": schema.ListAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"pass": schema.ListAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"upstream": schema.StringAttribute{
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

													"redirect": schema.SingleNestedAttribute{
														Description:         "ActionRedirect defines a redirect in an Action.",
														MarkdownDescription: "ActionRedirect defines a redirect in an Action.",
														Attributes: map[string]schema.Attribute{
															"code": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"url": schema.StringAttribute{
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

													"return": schema.SingleNestedAttribute{
														Description:         "ActionReturn defines a return in an Action.",
														MarkdownDescription: "ActionReturn defines a return in an Action.",
														Attributes: map[string]schema.Attribute{
															"body": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"code": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"headers": schema.ListNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"value": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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
												Required: false,
												Optional: true,
												Computed: false,
											},

											"conditions": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"argument": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"cookie": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"header": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"value": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"variable": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
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
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"action": schema.SingleNestedAttribute{
															Description:         "Action defines an action.",
															MarkdownDescription: "Action defines an action.",
															Attributes: map[string]schema.Attribute{
																"pass": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"proxy": schema.SingleNestedAttribute{
																	Description:         "ActionProxy defines a proxy in an Action.",
																	MarkdownDescription: "ActionProxy defines a proxy in an Action.",
																	Attributes: map[string]schema.Attribute{
																		"request_headers": schema.SingleNestedAttribute{
																			Description:         "ProxyRequestHeaders defines the request headers manipulation in an ActionProxy.",
																			MarkdownDescription: "ProxyRequestHeaders defines the request headers manipulation in an ActionProxy.",
																			Attributes: map[string]schema.Attribute{
																				"pass": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"set": schema.ListNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"value": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
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
																			Description:         "ProxyResponseHeaders defines the response headers manipulation in an ActionProxy.",
																			MarkdownDescription: "ProxyResponseHeaders defines the response headers manipulation in an ActionProxy.",
																			Attributes: map[string]schema.Attribute{
																				"add": schema.ListNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"always": schema.BoolAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"value": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
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
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"ignore": schema.ListAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"pass": schema.ListAttribute{
																					Description:         "",
																					MarkdownDescription: "",
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
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"upstream": schema.StringAttribute{
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

																"redirect": schema.SingleNestedAttribute{
																	Description:         "ActionRedirect defines a redirect in an Action.",
																	MarkdownDescription: "ActionRedirect defines a redirect in an Action.",
																	Attributes: map[string]schema.Attribute{
																		"code": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"url": schema.StringAttribute{
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

																"return": schema.SingleNestedAttribute{
																	Description:         "ActionReturn defines a return in an Action.",
																	MarkdownDescription: "ActionReturn defines a return in an Action.",
																	Attributes: map[string]schema.Attribute{
																		"body": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"code": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"headers": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"value": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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
															Required: false,
															Optional: true,
															Computed: false,
														},

														"weight": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
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
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"policies": schema.ListNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
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
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"splits": schema.ListNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"action": schema.SingleNestedAttribute{
												Description:         "Action defines an action.",
												MarkdownDescription: "Action defines an action.",
												Attributes: map[string]schema.Attribute{
													"pass": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"proxy": schema.SingleNestedAttribute{
														Description:         "ActionProxy defines a proxy in an Action.",
														MarkdownDescription: "ActionProxy defines a proxy in an Action.",
														Attributes: map[string]schema.Attribute{
															"request_headers": schema.SingleNestedAttribute{
																Description:         "ProxyRequestHeaders defines the request headers manipulation in an ActionProxy.",
																MarkdownDescription: "ProxyRequestHeaders defines the request headers manipulation in an ActionProxy.",
																Attributes: map[string]schema.Attribute{
																	"pass": schema.BoolAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"set": schema.ListNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"name": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"value": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
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
																Description:         "ProxyResponseHeaders defines the response headers manipulation in an ActionProxy.",
																MarkdownDescription: "ProxyResponseHeaders defines the response headers manipulation in an ActionProxy.",
																Attributes: map[string]schema.Attribute{
																	"add": schema.ListNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"always": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"name": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"value": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
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
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"ignore": schema.ListAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"pass": schema.ListAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"upstream": schema.StringAttribute{
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

													"redirect": schema.SingleNestedAttribute{
														Description:         "ActionRedirect defines a redirect in an Action.",
														MarkdownDescription: "ActionRedirect defines a redirect in an Action.",
														Attributes: map[string]schema.Attribute{
															"code": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"url": schema.StringAttribute{
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

													"return": schema.SingleNestedAttribute{
														Description:         "ActionReturn defines a return in an Action.",
														MarkdownDescription: "ActionReturn defines a return in an Action.",
														Attributes: map[string]schema.Attribute{
															"body": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"code": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"headers": schema.ListNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"value": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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
												Required: false,
												Optional: true,
												Computed: false,
											},

											"weight": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
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

					"server_snippets": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tls": schema.SingleNestedAttribute{
						Description:         "TLS defines TLS configuration for a VirtualServer.",
						MarkdownDescription: "TLS defines TLS configuration for a VirtualServer.",
						Attributes: map[string]schema.Attribute{
							"cert_manager": schema.SingleNestedAttribute{
								Description:         "CertManager defines a cert manager config for a TLS.",
								MarkdownDescription: "CertManager defines a cert manager config for a TLS.",
								Attributes: map[string]schema.Attribute{
									"cluster_issuer": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"common_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"duration": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"issue_temp_cert": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"issuer": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"issuer_group": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"issuer_kind": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"renew_before": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"usages": schema.StringAttribute{
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

							"redirect": schema.SingleNestedAttribute{
								Description:         "TLSRedirect defines a redirect for a TLS.",
								MarkdownDescription: "TLSRedirect defines a redirect for a TLS.",
								Attributes: map[string]schema.Attribute{
									"based_on": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"code": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"enable": schema.BoolAttribute{
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

					"upstreams": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"backup": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"backup_port": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"buffer_size": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"buffering": schema.BoolAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"buffers": schema.SingleNestedAttribute{
									Description:         "UpstreamBuffers defines Buffer Configuration for an Upstream.",
									MarkdownDescription: "UpstreamBuffers defines Buffer Configuration for an Upstream.",
									Attributes: map[string]schema.Attribute{
										"number": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"size": schema.StringAttribute{
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

								"client_max_body_size": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"connect_timeout": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"fail_timeout": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"health_check": schema.SingleNestedAttribute{
									Description:         "HealthCheck defines the parameters for active Upstream HealthChecks.",
									MarkdownDescription: "HealthCheck defines the parameters for active Upstream HealthChecks.",
									Attributes: map[string]schema.Attribute{
										"connect_timeout": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"enable": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"fails": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"grpc_service": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"grpc_status": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"headers": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"value": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
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
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"jitter": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"keepalive_time": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"mandatory": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"passes": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"persistent": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"port": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"read_timeout": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"send_timeout": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"status_match": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tls": schema.SingleNestedAttribute{
											Description:         "UpstreamTLS defines a TLS configuration for an Upstream.",
											MarkdownDescription: "UpstreamTLS defines a TLS configuration for an Upstream.",
											Attributes: map[string]schema.Attribute{
												"enable": schema.BoolAttribute{
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
									Required: false,
									Optional: true,
									Computed: false,
								},

								"keepalive": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"lb_method": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"max_conns": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"max_fails": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"next_upstream": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"next_upstream_timeout": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"next_upstream_tries": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ntlm": schema.BoolAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"port": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"queue": schema.SingleNestedAttribute{
									Description:         "UpstreamQueue defines Queue Configuration for an Upstream.",
									MarkdownDescription: "UpstreamQueue defines Queue Configuration for an Upstream.",
									Attributes: map[string]schema.Attribute{
										"size": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"timeout": schema.StringAttribute{
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

								"read_timeout": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"send_timeout": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"service": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"session_cookie": schema.SingleNestedAttribute{
									Description:         "SessionCookie defines the parameters for session persistence.",
									MarkdownDescription: "SessionCookie defines the parameters for session persistence.",
									Attributes: map[string]schema.Attribute{
										"domain": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"enable": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"expires": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"http_only": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"samesite": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secure": schema.BoolAttribute{
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

								"slow_start": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"subselector": schema.MapAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"tls": schema.SingleNestedAttribute{
									Description:         "UpstreamTLS defines a TLS configuration for an Upstream.",
									MarkdownDescription: "UpstreamTLS defines a TLS configuration for an Upstream.",
									Attributes: map[string]schema.Attribute{
										"enable": schema.BoolAttribute{
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

								"type": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"use_cluster_ip": schema.BoolAttribute{
									Description:         "",
									MarkdownDescription: "",
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

func (r *K8SNginxOrgVirtualServerV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_k8s_nginx_org_virtual_server_v1_manifest")

	var model K8SNginxOrgVirtualServerV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("k8s.nginx.org/v1")
	model.Kind = pointer.String("VirtualServer")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
