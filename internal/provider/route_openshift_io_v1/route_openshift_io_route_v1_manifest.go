/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package route_openshift_io_v1

import (
	"context"
	"fmt"
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
	_ datasource.DataSource = &RouteOpenshiftIoRouteV1Manifest{}
)

func NewRouteOpenshiftIoRouteV1Manifest() datasource.DataSource {
	return &RouteOpenshiftIoRouteV1Manifest{}
}

type RouteOpenshiftIoRouteV1Manifest struct{}

type RouteOpenshiftIoRouteV1ManifestData struct {
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
		AlternateBackends *[]struct {
			Kind   *string `tfsdk:"kind" json:"kind,omitempty"`
			Name   *string `tfsdk:"name" json:"name,omitempty"`
			Weight *int64  `tfsdk:"weight" json:"weight,omitempty"`
		} `tfsdk:"alternate_backends" json:"alternateBackends,omitempty"`
		Host        *string `tfsdk:"host" json:"host,omitempty"`
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
		} `tfsdk:"http_headers" json:"httpHeaders,omitempty"`
		Path *string `tfsdk:"path" json:"path,omitempty"`
		Port *struct {
			TargetPort *string `tfsdk:"target_port" json:"targetPort,omitempty"`
		} `tfsdk:"port" json:"port,omitempty"`
		Subdomain *string `tfsdk:"subdomain" json:"subdomain,omitempty"`
		Tls       *struct {
			CaCertificate                 *string `tfsdk:"ca_certificate" json:"caCertificate,omitempty"`
			Certificate                   *string `tfsdk:"certificate" json:"certificate,omitempty"`
			DestinationCACertificate      *string `tfsdk:"destination_ca_certificate" json:"destinationCACertificate,omitempty"`
			InsecureEdgeTerminationPolicy *string `tfsdk:"insecure_edge_termination_policy" json:"insecureEdgeTerminationPolicy,omitempty"`
			Key                           *string `tfsdk:"key" json:"key,omitempty"`
			Termination                   *string `tfsdk:"termination" json:"termination,omitempty"`
		} `tfsdk:"tls" json:"tls,omitempty"`
		To *struct {
			Kind   *string `tfsdk:"kind" json:"kind,omitempty"`
			Name   *string `tfsdk:"name" json:"name,omitempty"`
			Weight *int64  `tfsdk:"weight" json:"weight,omitempty"`
		} `tfsdk:"to" json:"to,omitempty"`
		WildcardPolicy *string `tfsdk:"wildcard_policy" json:"wildcardPolicy,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *RouteOpenshiftIoRouteV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_route_openshift_io_route_v1_manifest"
}

func (r *RouteOpenshiftIoRouteV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "A route allows developers to expose services through an HTTP(S) aware load balancing and proxy layer via a public DNS entry. The route may further specify TLS options and a certificate, or specify a public CNAME that the router should also accept for HTTP and HTTPS traffic. An administrator typically configures their router to be visible outside the cluster firewall, and may also add additional security, caching, or traffic controls on the service content. Routers usually talk directly to the service endpoints.  Once a route is created, the 'host' field may not be changed. Generally, routers use the oldest route with a given host when resolving conflicts.  Routers are subject to additional customization and may support additional controls via the annotations field.  Because administrators may configure multiple routers, the route status field is used to return information to clients about the names and states of the route under each router. If a client chooses a duplicate name, for instance, the route status conditions are used to indicate the route cannot be chosen.  To enable HTTP/2 ALPN on a route it requires a custom (non-wildcard) certificate. This prevents connection coalescing by clients, notably web browsers. We do not support HTTP/2 ALPN on routes that use the default certificate because of the risk of connection re-use/coalescing. Routes that do not have their own custom certificate will not be HTTP/2 ALPN-enabled on either the frontend or the backend.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		MarkdownDescription: "A route allows developers to expose services through an HTTP(S) aware load balancing and proxy layer via a public DNS entry. The route may further specify TLS options and a certificate, or specify a public CNAME that the router should also accept for HTTP and HTTPS traffic. An administrator typically configures their router to be visible outside the cluster firewall, and may also add additional security, caching, or traffic controls on the service content. Routers usually talk directly to the service endpoints.  Once a route is created, the 'host' field may not be changed. Generally, routers use the oldest route with a given host when resolving conflicts.  Routers are subject to additional customization and may support additional controls via the annotations field.  Because administrators may configure multiple routers, the route status field is used to return information to clients about the names and states of the route under each router. If a client chooses a duplicate name, for instance, the route status conditions are used to indicate the route cannot be chosen.  To enable HTTP/2 ALPN on a route it requires a custom (non-wildcard) certificate. This prevents connection coalescing by clients, notably web browsers. We do not support HTTP/2 ALPN on routes that use the default certificate because of the risk of connection re-use/coalescing. Routes that do not have their own custom certificate will not be HTTP/2 ALPN-enabled on either the frontend or the backend.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
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
				Description:         "spec is the desired state of the route",
				MarkdownDescription: "spec is the desired state of the route",
				Attributes: map[string]schema.Attribute{
					"alternate_backends": schema.ListNestedAttribute{
						Description:         "alternateBackends allows up to 3 additional backends to be assigned to the route. Only the Service kind is allowed, and it will be defaulted to Service. Use the weight field in RouteTargetReference object to specify relative preference.",
						MarkdownDescription: "alternateBackends allows up to 3 additional backends to be assigned to the route. Only the Service kind is allowed, and it will be defaulted to Service. Use the weight field in RouteTargetReference object to specify relative preference.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"kind": schema.StringAttribute{
									Description:         "The kind of target that the route is referring to. Currently, only 'Service' is allowed",
									MarkdownDescription: "The kind of target that the route is referring to. Currently, only 'Service' is allowed",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("Service", ""),
									},
								},

								"name": schema.StringAttribute{
									Description:         "name of the service/target that is being referred to. e.g. name of the service",
									MarkdownDescription: "name of the service/target that is being referred to. e.g. name of the service",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
									},
								},

								"weight": schema.Int64Attribute{
									Description:         "weight as an integer between 0 and 256, default 100, that specifies the target's relative weight against other target reference objects. 0 suppresses requests to this backend.",
									MarkdownDescription: "weight as an integer between 0 and 256, default 100, that specifies the target's relative weight against other target reference objects. 0 suppresses requests to this backend.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.Int64{
										int64validator.AtLeast(0),
										int64validator.AtMost(256),
									},
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"host": schema.StringAttribute{
						Description:         "host is an alias/DNS that points to the service. Optional. If not specified a route name will typically be automatically chosen. Must follow DNS952 subdomain conventions.",
						MarkdownDescription: "host is an alias/DNS that points to the service. Optional. If not specified a route name will typically be automatically chosen. Must follow DNS952 subdomain conventions.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(253),
							stringvalidator.RegexMatches(regexp.MustCompile(`^([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])(\.([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9]))*$`), ""),
						},
					},

					"http_headers": schema.SingleNestedAttribute{
						Description:         "httpHeaders defines policy for HTTP headers.",
						MarkdownDescription: "httpHeaders defines policy for HTTP headers.",
						Attributes: map[string]schema.Attribute{
							"actions": schema.SingleNestedAttribute{
								Description:         "actions specifies options for modifying headers and their values. Note that this option only applies to cleartext HTTP connections and to secure HTTP connections for which the ingress controller terminates encryption (that is, edge-terminated or reencrypt connections).  Headers cannot be modified for TLS passthrough connections. Setting the HSTS ('Strict-Transport-Security') header is not supported via actions. 'Strict-Transport-Security' may only be configured using the 'haproxy.router.openshift.io/hsts_header' route annotation, and only in accordance with the policy specified in Ingress.Spec.RequiredHSTSPolicies. In case of HTTP request headers, the actions specified in spec.httpHeaders.actions on the Route will be executed after the actions specified in the IngressController's spec.httpHeaders.actions field. In case of HTTP response headers, the actions specified in spec.httpHeaders.actions on the IngressController will be executed after the actions specified in the Route's spec.httpHeaders.actions field. The headers set via this API will not appear in access logs. Any actions defined here are applied after any actions related to the following other fields: cache-control, spec.clientTLS, spec.httpHeaders.forwardedHeaderPolicy, spec.httpHeaders.uniqueId, and spec.httpHeaders.headerNameCaseAdjustments. The following header names are reserved and may not be modified via this API: Strict-Transport-Security, Proxy, Cookie, Set-Cookie. Note that the total size of all net added headers *after* interpolating dynamic values must not exceed the value of spec.tuningOptions.headerBufferMaxRewriteBytes on the IngressController. Please refer to the documentation for that API field for more details.",
								MarkdownDescription: "actions specifies options for modifying headers and their values. Note that this option only applies to cleartext HTTP connections and to secure HTTP connections for which the ingress controller terminates encryption (that is, edge-terminated or reencrypt connections).  Headers cannot be modified for TLS passthrough connections. Setting the HSTS ('Strict-Transport-Security') header is not supported via actions. 'Strict-Transport-Security' may only be configured using the 'haproxy.router.openshift.io/hsts_header' route annotation, and only in accordance with the policy specified in Ingress.Spec.RequiredHSTSPolicies. In case of HTTP request headers, the actions specified in spec.httpHeaders.actions on the Route will be executed after the actions specified in the IngressController's spec.httpHeaders.actions field. In case of HTTP response headers, the actions specified in spec.httpHeaders.actions on the IngressController will be executed after the actions specified in the Route's spec.httpHeaders.actions field. The headers set via this API will not appear in access logs. Any actions defined here are applied after any actions related to the following other fields: cache-control, spec.clientTLS, spec.httpHeaders.forwardedHeaderPolicy, spec.httpHeaders.uniqueId, and spec.httpHeaders.headerNameCaseAdjustments. The following header names are reserved and may not be modified via this API: Strict-Transport-Security, Proxy, Cookie, Set-Cookie. Note that the total size of all net added headers *after* interpolating dynamic values must not exceed the value of spec.tuningOptions.headerBufferMaxRewriteBytes on the IngressController. Please refer to the documentation for that API field for more details.",
								Attributes: map[string]schema.Attribute{
									"request": schema.ListNestedAttribute{
										Description:         "request is a list of HTTP request headers to modify. Currently, actions may define to either 'Set' or 'Delete' headers values. Actions defined here will modify the request headers of all requests made through a route. These actions are applied to a specific Route defined within a cluster i.e. connections made through a route. Currently, actions may define to either 'Set' or 'Delete' headers values. Route actions will be executed after IngressController actions for request headers. Actions are applied in sequence as defined in this list. A maximum of 20 request header actions may be configured. You can use this field to specify HTTP request headers that should be set or deleted when forwarding connections from the client to your application. Sample fetchers allowed are 'req.hdr' and 'ssl_c_der'. Converters allowed are 'lower' and 'base64'. Example header values: '%[req.hdr(X-target),lower]', '%{+Q}[ssl_c_der,base64]'. Any request header configuration applied directly via a Route resource using this API will override header configuration for a header of the same name applied via spec.httpHeaders.actions on the IngressController or route annotation. Note: This field cannot be used if your route uses TLS passthrough.",
										MarkdownDescription: "request is a list of HTTP request headers to modify. Currently, actions may define to either 'Set' or 'Delete' headers values. Actions defined here will modify the request headers of all requests made through a route. These actions are applied to a specific Route defined within a cluster i.e. connections made through a route. Currently, actions may define to either 'Set' or 'Delete' headers values. Route actions will be executed after IngressController actions for request headers. Actions are applied in sequence as defined in this list. A maximum of 20 request header actions may be configured. You can use this field to specify HTTP request headers that should be set or deleted when forwarding connections from the client to your application. Sample fetchers allowed are 'req.hdr' and 'ssl_c_der'. Converters allowed are 'lower' and 'base64'. Example header values: '%[req.hdr(X-target),lower]', '%{+Q}[ssl_c_der,base64]'. Any request header configuration applied directly via a Route resource using this API will override header configuration for a header of the same name applied via spec.httpHeaders.actions on the IngressController or route annotation. Note: This field cannot be used if your route uses TLS passthrough.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"action": schema.SingleNestedAttribute{
													Description:         "action specifies actions to perform on headers, such as setting or deleting headers.",
													MarkdownDescription: "action specifies actions to perform on headers, such as setting or deleting headers.",
													Attributes: map[string]schema.Attribute{
														"set": schema.SingleNestedAttribute{
															Description:         "set defines the HTTP header that should be set: added if it doesn't exist or replaced if it does. This field is required when type is Set and forbidden otherwise.",
															MarkdownDescription: "set defines the HTTP header that should be set: added if it doesn't exist or replaced if it does. This field is required when type is Set and forbidden otherwise.",
															Attributes: map[string]schema.Attribute{
																"value": schema.StringAttribute{
																	Description:         "value specifies a header value. Dynamic values can be added. The value will be interpreted as an HAProxy format string as defined in http://cbonte.github.io/haproxy-dconv/2.6/configuration.html#8.2.6 and may use HAProxy's %[] syntax and otherwise must be a valid HTTP header value as defined in https://datatracker.ietf.org/doc/html/rfc7230#section-3.2. The value of this field must be no more than 16384 characters in length. Note that the total size of all net added headers *after* interpolating dynamic values must not exceed the value of spec.tuningOptions.headerBufferMaxRewriteBytes on the IngressController.",
																	MarkdownDescription: "value specifies a header value. Dynamic values can be added. The value will be interpreted as an HAProxy format string as defined in http://cbonte.github.io/haproxy-dconv/2.6/configuration.html#8.2.6 and may use HAProxy's %[] syntax and otherwise must be a valid HTTP header value as defined in https://datatracker.ietf.org/doc/html/rfc7230#section-3.2. The value of this field must be no more than 16384 characters in length. Note that the total size of all net added headers *after* interpolating dynamic values must not exceed the value of spec.tuningOptions.headerBufferMaxRewriteBytes on the IngressController.",
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
													Description:         "name specifies the name of a header on which to perform an action. Its value must be a valid HTTP header name as defined in RFC 2616 section 4.2. The name must consist only of alphanumeric and the following special characters, '-!#$%&'*+.^_''. The following header names are reserved and may not be modified via this API: Strict-Transport-Security, Proxy, Cookie, Set-Cookie. It must be no more than 255 characters in length. Header name must be unique.",
													MarkdownDescription: "name specifies the name of a header on which to perform an action. Its value must be a valid HTTP header name as defined in RFC 2616 section 4.2. The name must consist only of alphanumeric and the following special characters, '-!#$%&'*+.^_''. The following header names are reserved and may not be modified via this API: Strict-Transport-Security, Proxy, Cookie, Set-Cookie. It must be no more than 255 characters in length. Header name must be unique.",
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
										Description:         "response is a list of HTTP response headers to modify. Currently, actions may define to either 'Set' or 'Delete' headers values. Actions defined here will modify the response headers of all requests made through a route. These actions are applied to a specific Route defined within a cluster i.e. connections made through a route. Route actions will be executed before IngressController actions for response headers. Actions are applied in sequence as defined in this list. A maximum of 20 response header actions may be configured. You can use this field to specify HTTP response headers that should be set or deleted when forwarding responses from your application to the client. Sample fetchers allowed are 'res.hdr' and 'ssl_c_der'. Converters allowed are 'lower' and 'base64'. Example header values: '%[res.hdr(X-target),lower]', '%{+Q}[ssl_c_der,base64]'. Note: This field cannot be used if your route uses TLS passthrough.",
										MarkdownDescription: "response is a list of HTTP response headers to modify. Currently, actions may define to either 'Set' or 'Delete' headers values. Actions defined here will modify the response headers of all requests made through a route. These actions are applied to a specific Route defined within a cluster i.e. connections made through a route. Route actions will be executed before IngressController actions for response headers. Actions are applied in sequence as defined in this list. A maximum of 20 response header actions may be configured. You can use this field to specify HTTP response headers that should be set or deleted when forwarding responses from your application to the client. Sample fetchers allowed are 'res.hdr' and 'ssl_c_der'. Converters allowed are 'lower' and 'base64'. Example header values: '%[res.hdr(X-target),lower]', '%{+Q}[ssl_c_der,base64]'. Note: This field cannot be used if your route uses TLS passthrough.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"action": schema.SingleNestedAttribute{
													Description:         "action specifies actions to perform on headers, such as setting or deleting headers.",
													MarkdownDescription: "action specifies actions to perform on headers, such as setting or deleting headers.",
													Attributes: map[string]schema.Attribute{
														"set": schema.SingleNestedAttribute{
															Description:         "set defines the HTTP header that should be set: added if it doesn't exist or replaced if it does. This field is required when type is Set and forbidden otherwise.",
															MarkdownDescription: "set defines the HTTP header that should be set: added if it doesn't exist or replaced if it does. This field is required when type is Set and forbidden otherwise.",
															Attributes: map[string]schema.Attribute{
																"value": schema.StringAttribute{
																	Description:         "value specifies a header value. Dynamic values can be added. The value will be interpreted as an HAProxy format string as defined in http://cbonte.github.io/haproxy-dconv/2.6/configuration.html#8.2.6 and may use HAProxy's %[] syntax and otherwise must be a valid HTTP header value as defined in https://datatracker.ietf.org/doc/html/rfc7230#section-3.2. The value of this field must be no more than 16384 characters in length. Note that the total size of all net added headers *after* interpolating dynamic values must not exceed the value of spec.tuningOptions.headerBufferMaxRewriteBytes on the IngressController.",
																	MarkdownDescription: "value specifies a header value. Dynamic values can be added. The value will be interpreted as an HAProxy format string as defined in http://cbonte.github.io/haproxy-dconv/2.6/configuration.html#8.2.6 and may use HAProxy's %[] syntax and otherwise must be a valid HTTP header value as defined in https://datatracker.ietf.org/doc/html/rfc7230#section-3.2. The value of this field must be no more than 16384 characters in length. Note that the total size of all net added headers *after* interpolating dynamic values must not exceed the value of spec.tuningOptions.headerBufferMaxRewriteBytes on the IngressController.",
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
													Description:         "name specifies the name of a header on which to perform an action. Its value must be a valid HTTP header name as defined in RFC 2616 section 4.2. The name must consist only of alphanumeric and the following special characters, '-!#$%&'*+.^_''. The following header names are reserved and may not be modified via this API: Strict-Transport-Security, Proxy, Cookie, Set-Cookie. It must be no more than 255 characters in length. Header name must be unique.",
													MarkdownDescription: "name specifies the name of a header on which to perform an action. Its value must be a valid HTTP header name as defined in RFC 2616 section 4.2. The name must consist only of alphanumeric and the following special characters, '-!#$%&'*+.^_''. The following header names are reserved and may not be modified via this API: Strict-Transport-Security, Proxy, Cookie, Set-Cookie. It must be no more than 255 characters in length. Header name must be unique.",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"path": schema.StringAttribute{
						Description:         "path that the router watches for, to route traffic for to the service. Optional",
						MarkdownDescription: "path that the router watches for, to route traffic for to the service. Optional",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^/`), ""),
						},
					},

					"port": schema.SingleNestedAttribute{
						Description:         "If specified, the port to be used by the router. Most routers will use all endpoints exposed by the service by default - set this value to instruct routers which port to use.",
						MarkdownDescription: "If specified, the port to be used by the router. Most routers will use all endpoints exposed by the service by default - set this value to instruct routers which port to use.",
						Attributes: map[string]schema.Attribute{
							"target_port": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"subdomain": schema.StringAttribute{
						Description:         "subdomain is a DNS subdomain that is requested within the ingress controller's domain (as a subdomain). If host is set this field is ignored. An ingress controller may choose to ignore this suggested name, in which case the controller will report the assigned name in the status.ingress array or refuse to admit the route. If this value is set and the server does not support this field host will be populated automatically. Otherwise host is left empty. The field may have multiple parts separated by a dot, but not all ingress controllers may honor the request. This field may not be changed after creation except by a user with the update routes/custom-host permission.  Example: subdomain 'frontend' automatically receives the router subdomain 'apps.mycluster.com' to have a full hostname 'frontend.apps.mycluster.com'.",
						MarkdownDescription: "subdomain is a DNS subdomain that is requested within the ingress controller's domain (as a subdomain). If host is set this field is ignored. An ingress controller may choose to ignore this suggested name, in which case the controller will report the assigned name in the status.ingress array or refuse to admit the route. If this value is set and the server does not support this field host will be populated automatically. Otherwise host is left empty. The field may have multiple parts separated by a dot, but not all ingress controllers may honor the request. This field may not be changed after creation except by a user with the update routes/custom-host permission.  Example: subdomain 'frontend' automatically receives the router subdomain 'apps.mycluster.com' to have a full hostname 'frontend.apps.mycluster.com'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(253),
							stringvalidator.RegexMatches(regexp.MustCompile(`^([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])(\.([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9]))*$`), ""),
						},
					},

					"tls": schema.SingleNestedAttribute{
						Description:         "The tls field provides the ability to configure certificates and termination for the route.",
						MarkdownDescription: "The tls field provides the ability to configure certificates and termination for the route.",
						Attributes: map[string]schema.Attribute{
							"ca_certificate": schema.StringAttribute{
								Description:         "caCertificate provides the cert authority certificate contents",
								MarkdownDescription: "caCertificate provides the cert authority certificate contents",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"certificate": schema.StringAttribute{
								Description:         "certificate provides certificate contents. This should be a single serving certificate, not a certificate chain. Do not include a CA certificate.",
								MarkdownDescription: "certificate provides certificate contents. This should be a single serving certificate, not a certificate chain. Do not include a CA certificate.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"destination_ca_certificate": schema.StringAttribute{
								Description:         "destinationCACertificate provides the contents of the ca certificate of the final destination.  When using reencrypt termination this file should be provided in order to have routers use it for health checks on the secure connection. If this field is not specified, the router may provide its own destination CA and perform hostname validation using the short service name (service.namespace.svc), which allows infrastructure generated certificates to automatically verify.",
								MarkdownDescription: "destinationCACertificate provides the contents of the ca certificate of the final destination.  When using reencrypt termination this file should be provided in order to have routers use it for health checks on the secure connection. If this field is not specified, the router may provide its own destination CA and perform hostname validation using the short service name (service.namespace.svc), which allows infrastructure generated certificates to automatically verify.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"insecure_edge_termination_policy": schema.StringAttribute{
								Description:         "insecureEdgeTerminationPolicy indicates the desired behavior for insecure connections to a route. While each router may make its own decisions on which ports to expose, this is normally port 80.  * Allow - traffic is sent to the server on the insecure port (edge/reencrypt terminations only) (default). * None - no traffic is allowed on the insecure port. * Redirect - clients are redirected to the secure port.",
								MarkdownDescription: "insecureEdgeTerminationPolicy indicates the desired behavior for insecure connections to a route. While each router may make its own decisions on which ports to expose, this is normally port 80.  * Allow - traffic is sent to the server on the insecure port (edge/reencrypt terminations only) (default). * None - no traffic is allowed on the insecure port. * Redirect - clients are redirected to the secure port.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Allow", "None", "Redirect", ""),
								},
							},

							"key": schema.StringAttribute{
								Description:         "key provides key file contents",
								MarkdownDescription: "key provides key file contents",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"termination": schema.StringAttribute{
								Description:         "termination indicates termination type.  * edge - TLS termination is done by the router and http is used to communicate with the backend (default) * passthrough - Traffic is sent straight to the destination without the router providing TLS termination * reencrypt - TLS termination is done by the router and https is used to communicate with the backend  Note: passthrough termination is incompatible with httpHeader actions",
								MarkdownDescription: "termination indicates termination type.  * edge - TLS termination is done by the router and http is used to communicate with the backend (default) * passthrough - Traffic is sent straight to the destination without the router providing TLS termination * reencrypt - TLS termination is done by the router and https is used to communicate with the backend  Note: passthrough termination is incompatible with httpHeader actions",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("edge", "reencrypt", "passthrough"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"to": schema.SingleNestedAttribute{
						Description:         "to is an object the route should use as the primary backend. Only the Service kind is allowed, and it will be defaulted to Service. If the weight field (0-256 default 100) is set to zero, no traffic will be sent to this backend.",
						MarkdownDescription: "to is an object the route should use as the primary backend. Only the Service kind is allowed, and it will be defaulted to Service. If the weight field (0-256 default 100) is set to zero, no traffic will be sent to this backend.",
						Attributes: map[string]schema.Attribute{
							"kind": schema.StringAttribute{
								Description:         "The kind of target that the route is referring to. Currently, only 'Service' is allowed",
								MarkdownDescription: "The kind of target that the route is referring to. Currently, only 'Service' is allowed",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Service", ""),
								},
							},

							"name": schema.StringAttribute{
								Description:         "name of the service/target that is being referred to. e.g. name of the service",
								MarkdownDescription: "name of the service/target that is being referred to. e.g. name of the service",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
							},

							"weight": schema.Int64Attribute{
								Description:         "weight as an integer between 0 and 256, default 100, that specifies the target's relative weight against other target reference objects. 0 suppresses requests to this backend.",
								MarkdownDescription: "weight as an integer between 0 and 256, default 100, that specifies the target's relative weight against other target reference objects. 0 suppresses requests to this backend.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
									int64validator.AtMost(256),
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"wildcard_policy": schema.StringAttribute{
						Description:         "Wildcard policy if any for the route. Currently only 'Subdomain' or 'None' is allowed.",
						MarkdownDescription: "Wildcard policy if any for the route. Currently only 'Subdomain' or 'None' is allowed.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("None", "Subdomain", ""),
						},
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *RouteOpenshiftIoRouteV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_route_openshift_io_route_v1_manifest")

	var model RouteOpenshiftIoRouteV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("route.openshift.io/v1")
	model.Kind = pointer.String("Route")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
