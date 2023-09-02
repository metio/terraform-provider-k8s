/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package gateway_networking_k8s_io_v1alpha2

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
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &GatewayNetworkingK8SIoHTTPRouteV1Alpha2Manifest{}
)

func NewGatewayNetworkingK8SIoHTTPRouteV1Alpha2Manifest() datasource.DataSource {
	return &GatewayNetworkingK8SIoHTTPRouteV1Alpha2Manifest{}
}

type GatewayNetworkingK8SIoHTTPRouteV1Alpha2Manifest struct{}

type GatewayNetworkingK8SIoHTTPRouteV1Alpha2ManifestData struct {
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
		Hostnames  *[]string `tfsdk:"hostnames" json:"hostnames,omitempty"`
		ParentRefs *[]struct {
			Group       *string `tfsdk:"group" json:"group,omitempty"`
			Kind        *string `tfsdk:"kind" json:"kind,omitempty"`
			Name        *string `tfsdk:"name" json:"name,omitempty"`
			Namespace   *string `tfsdk:"namespace" json:"namespace,omitempty"`
			SectionName *string `tfsdk:"section_name" json:"sectionName,omitempty"`
		} `tfsdk:"parent_refs" json:"parentRefs,omitempty"`
		Rules *[]struct {
			BackendRefs *[]struct {
				Filters *[]struct {
					ExtensionRef *struct {
						Group *string `tfsdk:"group" json:"group,omitempty"`
						Kind  *string `tfsdk:"kind" json:"kind,omitempty"`
						Name  *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"extension_ref" json:"extensionRef,omitempty"`
					RequestHeaderModifier *struct {
						Add *[]struct {
							Name  *string `tfsdk:"name" json:"name,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"add" json:"add,omitempty"`
						Remove *[]string `tfsdk:"remove" json:"remove,omitempty"`
						Set    *[]struct {
							Name  *string `tfsdk:"name" json:"name,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"set" json:"set,omitempty"`
					} `tfsdk:"request_header_modifier" json:"requestHeaderModifier,omitempty"`
					RequestMirror *struct {
						BackendRef *struct {
							Group     *string `tfsdk:"group" json:"group,omitempty"`
							Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
							Port      *int64  `tfsdk:"port" json:"port,omitempty"`
						} `tfsdk:"backend_ref" json:"backendRef,omitempty"`
					} `tfsdk:"request_mirror" json:"requestMirror,omitempty"`
					RequestRedirect *struct {
						Hostname   *string `tfsdk:"hostname" json:"hostname,omitempty"`
						Port       *int64  `tfsdk:"port" json:"port,omitempty"`
						Scheme     *string `tfsdk:"scheme" json:"scheme,omitempty"`
						StatusCode *int64  `tfsdk:"status_code" json:"statusCode,omitempty"`
					} `tfsdk:"request_redirect" json:"requestRedirect,omitempty"`
					Type *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"filters" json:"filters,omitempty"`
				Group     *string `tfsdk:"group" json:"group,omitempty"`
				Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				Port      *int64  `tfsdk:"port" json:"port,omitempty"`
				Weight    *int64  `tfsdk:"weight" json:"weight,omitempty"`
			} `tfsdk:"backend_refs" json:"backendRefs,omitempty"`
			Filters *[]struct {
				ExtensionRef *struct {
					Group *string `tfsdk:"group" json:"group,omitempty"`
					Kind  *string `tfsdk:"kind" json:"kind,omitempty"`
					Name  *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"extension_ref" json:"extensionRef,omitempty"`
				RequestHeaderModifier *struct {
					Add *[]struct {
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"add" json:"add,omitempty"`
					Remove *[]string `tfsdk:"remove" json:"remove,omitempty"`
					Set    *[]struct {
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"set" json:"set,omitempty"`
				} `tfsdk:"request_header_modifier" json:"requestHeaderModifier,omitempty"`
				RequestMirror *struct {
					BackendRef *struct {
						Group     *string `tfsdk:"group" json:"group,omitempty"`
						Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						Port      *int64  `tfsdk:"port" json:"port,omitempty"`
					} `tfsdk:"backend_ref" json:"backendRef,omitempty"`
				} `tfsdk:"request_mirror" json:"requestMirror,omitempty"`
				RequestRedirect *struct {
					Hostname   *string `tfsdk:"hostname" json:"hostname,omitempty"`
					Port       *int64  `tfsdk:"port" json:"port,omitempty"`
					Scheme     *string `tfsdk:"scheme" json:"scheme,omitempty"`
					StatusCode *int64  `tfsdk:"status_code" json:"statusCode,omitempty"`
				} `tfsdk:"request_redirect" json:"requestRedirect,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"filters" json:"filters,omitempty"`
			Matches *[]struct {
				Headers *[]struct {
					Name  *string `tfsdk:"name" json:"name,omitempty"`
					Type  *string `tfsdk:"type" json:"type,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"headers" json:"headers,omitempty"`
				Method *string `tfsdk:"method" json:"method,omitempty"`
				Path   *struct {
					Type  *string `tfsdk:"type" json:"type,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"path" json:"path,omitempty"`
				QueryParams *[]struct {
					Name  *string `tfsdk:"name" json:"name,omitempty"`
					Type  *string `tfsdk:"type" json:"type,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"query_params" json:"queryParams,omitempty"`
			} `tfsdk:"matches" json:"matches,omitempty"`
		} `tfsdk:"rules" json:"rules,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *GatewayNetworkingK8SIoHTTPRouteV1Alpha2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_gateway_networking_k8s_io_http_route_v1alpha2_manifest"
}

func (r *GatewayNetworkingK8SIoHTTPRouteV1Alpha2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "HTTPRoute provides a way to route HTTP requests. This includes the capability to match requests by hostname, path, header, or query param. Filters can be used to specify additional processing steps. Backends specify where matching requests should be routed.",
		MarkdownDescription: "HTTPRoute provides a way to route HTTP requests. This includes the capability to match requests by hostname, path, header, or query param. Filters can be used to specify additional processing steps. Backends specify where matching requests should be routed.",
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
				Description:         "Spec defines the desired state of HTTPRoute.",
				MarkdownDescription: "Spec defines the desired state of HTTPRoute.",
				Attributes: map[string]schema.Attribute{
					"hostnames": schema.ListAttribute{
						Description:         "Hostnames defines a set of hostname that should match against the HTTP Host header to select a HTTPRoute to process the request. This matches the RFC 1123 definition of a hostname with 2 notable exceptions:  1. IPs are not allowed. 2. A hostname may be prefixed with a wildcard label ('*.'). The wildcard    label must appear by itself as the first label.  If a hostname is specified by both the Listener and HTTPRoute, there must be at least one intersecting hostname for the HTTPRoute to be attached to the Listener. For example:  * A Listener with 'test.example.com' as the hostname matches HTTPRoutes   that have either not specified any hostnames, or have specified at   least one of 'test.example.com' or '*.example.com'. * A Listener with '*.example.com' as the hostname matches HTTPRoutes   that have either not specified any hostnames or have specified at least   one hostname that matches the Listener hostname. For example,   'test.example.com' and '*.example.com' would both match. On the other   hand, 'example.com' and 'test.example.net' would not match.  If both the Listener and HTTPRoute have specified hostnames, any HTTPRoute hostnames that do not match the Listener hostname MUST be ignored. For example, if a Listener specified '*.example.com', and the HTTPRoute specified 'test.example.com' and 'test.example.net', 'test.example.net' must not be considered for a match.  If both the Listener and HTTPRoute have specified hostnames, and none match with the criteria above, then the HTTPRoute is not accepted. The implementation must raise an 'Accepted' Condition with a status of 'False' in the corresponding RouteParentStatus.  Support: Core",
						MarkdownDescription: "Hostnames defines a set of hostname that should match against the HTTP Host header to select a HTTPRoute to process the request. This matches the RFC 1123 definition of a hostname with 2 notable exceptions:  1. IPs are not allowed. 2. A hostname may be prefixed with a wildcard label ('*.'). The wildcard    label must appear by itself as the first label.  If a hostname is specified by both the Listener and HTTPRoute, there must be at least one intersecting hostname for the HTTPRoute to be attached to the Listener. For example:  * A Listener with 'test.example.com' as the hostname matches HTTPRoutes   that have either not specified any hostnames, or have specified at   least one of 'test.example.com' or '*.example.com'. * A Listener with '*.example.com' as the hostname matches HTTPRoutes   that have either not specified any hostnames or have specified at least   one hostname that matches the Listener hostname. For example,   'test.example.com' and '*.example.com' would both match. On the other   hand, 'example.com' and 'test.example.net' would not match.  If both the Listener and HTTPRoute have specified hostnames, any HTTPRoute hostnames that do not match the Listener hostname MUST be ignored. For example, if a Listener specified '*.example.com', and the HTTPRoute specified 'test.example.com' and 'test.example.net', 'test.example.net' must not be considered for a match.  If both the Listener and HTTPRoute have specified hostnames, and none match with the criteria above, then the HTTPRoute is not accepted. The implementation must raise an 'Accepted' Condition with a status of 'False' in the corresponding RouteParentStatus.  Support: Core",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"parent_refs": schema.ListNestedAttribute{
						Description:         "ParentRefs references the resources (usually Gateways) that a Route wants to be attached to. Note that the referenced parent resource needs to allow this for the attachment to be complete. For Gateways, that means the Gateway needs to allow attachment from Routes of this kind and namespace.  The only kind of parent resource with 'Core' support is Gateway. This API may be extended in the future to support additional kinds of parent resources such as one of the route kinds.  It is invalid to reference an identical parent more than once. It is valid to reference multiple distinct sections within the same parent resource, such as 2 Listeners within a Gateway.  It is possible to separately reference multiple distinct objects that may be collapsed by an implementation. For example, some implementations may choose to merge compatible Gateway Listeners together. If that is the case, the list of routes attached to those resources should also be merged.",
						MarkdownDescription: "ParentRefs references the resources (usually Gateways) that a Route wants to be attached to. Note that the referenced parent resource needs to allow this for the attachment to be complete. For Gateways, that means the Gateway needs to allow attachment from Routes of this kind and namespace.  The only kind of parent resource with 'Core' support is Gateway. This API may be extended in the future to support additional kinds of parent resources such as one of the route kinds.  It is invalid to reference an identical parent more than once. It is valid to reference multiple distinct sections within the same parent resource, such as 2 Listeners within a Gateway.  It is possible to separately reference multiple distinct objects that may be collapsed by an implementation. For example, some implementations may choose to merge compatible Gateway Listeners together. If that is the case, the list of routes attached to those resources should also be merged.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"group": schema.StringAttribute{
									Description:         "Group is the group of the referent.  Support: Core",
									MarkdownDescription: "Group is the group of the referent.  Support: Core",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(253),
										stringvalidator.RegexMatches(regexp.MustCompile(`^$|^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
									},
								},

								"kind": schema.StringAttribute{
									Description:         "Kind is kind of the referent.  Support: Core (Gateway) Support: Custom (Other Resources)",
									MarkdownDescription: "Kind is kind of the referent.  Support: Core (Gateway) Support: Custom (Other Resources)",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(63),
										stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z]([-a-zA-Z0-9]*[a-zA-Z0-9])?$`), ""),
									},
								},

								"name": schema.StringAttribute{
									Description:         "Name is the name of the referent.  Support: Core",
									MarkdownDescription: "Name is the name of the referent.  Support: Core",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(253),
									},
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace is the namespace of the referent. When unspecified (or empty string), this refers to the local namespace of the Route.  Support: Core",
									MarkdownDescription: "Namespace is the namespace of the referent. When unspecified (or empty string), this refers to the local namespace of the Route.  Support: Core",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(63),
										stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
									},
								},

								"section_name": schema.StringAttribute{
									Description:         "SectionName is the name of a section within the target resource. In the following resources, SectionName is interpreted as the following:  * Gateway: Listener Name  Implementations MAY choose to support attaching Routes to other resources. If that is the case, they MUST clearly document how SectionName is interpreted.  When unspecified (empty string), this will reference the entire resource. For the purpose of status, an attachment is considered successful if at least one section in the parent resource accepts it. For example, Gateway listeners can restrict which Routes can attach to them by Route kind, namespace, or hostname. If 1 of 2 Gateway listeners accept attachment from the referencing Route, the Route MUST be considered successfully attached. If no Gateway listeners accept attachment from this Route, the Route MUST be considered detached from the Gateway.  Support: Core",
									MarkdownDescription: "SectionName is the name of a section within the target resource. In the following resources, SectionName is interpreted as the following:  * Gateway: Listener Name  Implementations MAY choose to support attaching Routes to other resources. If that is the case, they MUST clearly document how SectionName is interpreted.  When unspecified (empty string), this will reference the entire resource. For the purpose of status, an attachment is considered successful if at least one section in the parent resource accepts it. For example, Gateway listeners can restrict which Routes can attach to them by Route kind, namespace, or hostname. If 1 of 2 Gateway listeners accept attachment from the referencing Route, the Route MUST be considered successfully attached. If no Gateway listeners accept attachment from this Route, the Route MUST be considered detached from the Gateway.  Support: Core",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(253),
										stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
									},
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"rules": schema.ListNestedAttribute{
						Description:         "Rules are a list of HTTP matchers, filters and actions.",
						MarkdownDescription: "Rules are a list of HTTP matchers, filters and actions.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"backend_refs": schema.ListNestedAttribute{
									Description:         "If unspecified or invalid (refers to a non-existent resource or a Service with no endpoints), the rule performs no forwarding. If there are also no filters specified that would result in a response being sent, a HTTP 503 status code is returned. 503 responses must be sent so that the overall weight is respected; if an invalid backend is requested to have 80% of requests, then 80% of requests must get a 503 instead.  Support: Core for Kubernetes Service Support: Custom for any other resource  Support for weight: Core",
									MarkdownDescription: "If unspecified or invalid (refers to a non-existent resource or a Service with no endpoints), the rule performs no forwarding. If there are also no filters specified that would result in a response being sent, a HTTP 503 status code is returned. 503 responses must be sent so that the overall weight is respected; if an invalid backend is requested to have 80% of requests, then 80% of requests must get a 503 instead.  Support: Core for Kubernetes Service Support: Custom for any other resource  Support for weight: Core",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"filters": schema.ListNestedAttribute{
												Description:         "Filters defined at this level should be executed if and only if the request is being forwarded to the backend defined here.  Support: Custom (For broader support of filters, use the Filters field in HTTPRouteRule.)",
												MarkdownDescription: "Filters defined at this level should be executed if and only if the request is being forwarded to the backend defined here.  Support: Custom (For broader support of filters, use the Filters field in HTTPRouteRule.)",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"extension_ref": schema.SingleNestedAttribute{
															Description:         "ExtensionRef is an optional, implementation-specific extension to the 'filter' behavior.  For example, resource 'myroutefilter' in group 'networking.example.net'). ExtensionRef MUST NOT be used for core and extended filters.  Support: Implementation-specific",
															MarkdownDescription: "ExtensionRef is an optional, implementation-specific extension to the 'filter' behavior.  For example, resource 'myroutefilter' in group 'networking.example.net'). ExtensionRef MUST NOT be used for core and extended filters.  Support: Implementation-specific",
															Attributes: map[string]schema.Attribute{
																"group": schema.StringAttribute{
																	Description:         "Group is the group of the referent. For example, 'networking.k8s.io'. When unspecified (empty string), core API group is inferred.",
																	MarkdownDescription: "Group is the group of the referent. For example, 'networking.k8s.io'. When unspecified (empty string), core API group is inferred.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtMost(253),
																		stringvalidator.RegexMatches(regexp.MustCompile(`^$|^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
																	},
																},

																"kind": schema.StringAttribute{
																	Description:         "Kind is kind of the referent. For example 'HTTPRoute' or 'Service'.",
																	MarkdownDescription: "Kind is kind of the referent. For example 'HTTPRoute' or 'Service'.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtLeast(1),
																		stringvalidator.LengthAtMost(63),
																		stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z]([-a-zA-Z0-9]*[a-zA-Z0-9])?$`), ""),
																	},
																},

																"name": schema.StringAttribute{
																	Description:         "Name is the name of the referent.",
																	MarkdownDescription: "Name is the name of the referent.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtLeast(1),
																		stringvalidator.LengthAtMost(253),
																	},
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"request_header_modifier": schema.SingleNestedAttribute{
															Description:         "RequestHeaderModifier defines a schema for a filter that modifies request headers.  Support: Core",
															MarkdownDescription: "RequestHeaderModifier defines a schema for a filter that modifies request headers.  Support: Core",
															Attributes: map[string]schema.Attribute{
																"add": schema.ListNestedAttribute{
																	Description:         "Add adds the given header(s) (name, value) to the request before the action. It appends to any existing values associated with the header name.  Input:   GET /foo HTTP/1.1   my-header: foo  Config:   add:   - name: 'my-header'     value: 'bar'  Output:   GET /foo HTTP/1.1   my-header: foo   my-header: bar",
																	MarkdownDescription: "Add adds the given header(s) (name, value) to the request before the action. It appends to any existing values associated with the header name.  Input:   GET /foo HTTP/1.1   my-header: foo  Config:   add:   - name: 'my-header'     value: 'bar'  Output:   GET /foo HTTP/1.1   my-header: foo   my-header: bar",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "Name is the name of the HTTP Header to be matched. Name matching MUST be case insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).  If multiple entries specify equivalent header names, the first entry with an equivalent name MUST be considered for a match. Subsequent entries with an equivalent header name MUST be ignored. Due to the case-insensitivity of header names, 'foo' and 'Foo' are considered equivalent.",
																				MarkdownDescription: "Name is the name of the HTTP Header to be matched. Name matching MUST be case insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).  If multiple entries specify equivalent header names, the first entry with an equivalent name MUST be considered for a match. Subsequent entries with an equivalent header name MUST be ignored. Due to the case-insensitivity of header names, 'foo' and 'Foo' are considered equivalent.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																				Validators: []validator.String{
																					stringvalidator.LengthAtLeast(1),
																					stringvalidator.LengthAtMost(256),
																					stringvalidator.RegexMatches(regexp.MustCompile(`^[A-Za-z0-9!#$%&'*+\-.^_\x60|~]+$`), ""),
																				},
																			},

																			"value": schema.StringAttribute{
																				Description:         "Value is the value of HTTP Header to be matched.",
																				MarkdownDescription: "Value is the value of HTTP Header to be matched.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																				Validators: []validator.String{
																					stringvalidator.LengthAtLeast(1),
																					stringvalidator.LengthAtMost(4096),
																				},
																			},
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"remove": schema.ListAttribute{
																	Description:         "Remove the given header(s) from the HTTP request before the action. The value of Remove is a list of HTTP header names. Note that the header names are case-insensitive (see https://datatracker.ietf.org/doc/html/rfc2616#section-4.2).  Input:   GET /foo HTTP/1.1   my-header1: foo   my-header2: bar   my-header3: baz  Config:   remove: ['my-header1', 'my-header3']  Output:   GET /foo HTTP/1.1   my-header2: bar",
																	MarkdownDescription: "Remove the given header(s) from the HTTP request before the action. The value of Remove is a list of HTTP header names. Note that the header names are case-insensitive (see https://datatracker.ietf.org/doc/html/rfc2616#section-4.2).  Input:   GET /foo HTTP/1.1   my-header1: foo   my-header2: bar   my-header3: baz  Config:   remove: ['my-header1', 'my-header3']  Output:   GET /foo HTTP/1.1   my-header2: bar",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"set": schema.ListNestedAttribute{
																	Description:         "Set overwrites the request with the given header (name, value) before the action.  Input:   GET /foo HTTP/1.1   my-header: foo  Config:   set:   - name: 'my-header'     value: 'bar'  Output:   GET /foo HTTP/1.1   my-header: bar",
																	MarkdownDescription: "Set overwrites the request with the given header (name, value) before the action.  Input:   GET /foo HTTP/1.1   my-header: foo  Config:   set:   - name: 'my-header'     value: 'bar'  Output:   GET /foo HTTP/1.1   my-header: bar",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "Name is the name of the HTTP Header to be matched. Name matching MUST be case insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).  If multiple entries specify equivalent header names, the first entry with an equivalent name MUST be considered for a match. Subsequent entries with an equivalent header name MUST be ignored. Due to the case-insensitivity of header names, 'foo' and 'Foo' are considered equivalent.",
																				MarkdownDescription: "Name is the name of the HTTP Header to be matched. Name matching MUST be case insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).  If multiple entries specify equivalent header names, the first entry with an equivalent name MUST be considered for a match. Subsequent entries with an equivalent header name MUST be ignored. Due to the case-insensitivity of header names, 'foo' and 'Foo' are considered equivalent.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																				Validators: []validator.String{
																					stringvalidator.LengthAtLeast(1),
																					stringvalidator.LengthAtMost(256),
																					stringvalidator.RegexMatches(regexp.MustCompile(`^[A-Za-z0-9!#$%&'*+\-.^_\x60|~]+$`), ""),
																				},
																			},

																			"value": schema.StringAttribute{
																				Description:         "Value is the value of HTTP Header to be matched.",
																				MarkdownDescription: "Value is the value of HTTP Header to be matched.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																				Validators: []validator.String{
																					stringvalidator.LengthAtLeast(1),
																					stringvalidator.LengthAtMost(4096),
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

														"request_mirror": schema.SingleNestedAttribute{
															Description:         "RequestMirror defines a schema for a filter that mirrors requests. Requests are sent to the specified destination, but responses from that destination are ignored.  Support: Extended",
															MarkdownDescription: "RequestMirror defines a schema for a filter that mirrors requests. Requests are sent to the specified destination, but responses from that destination are ignored.  Support: Extended",
															Attributes: map[string]schema.Attribute{
																"backend_ref": schema.SingleNestedAttribute{
																	Description:         "BackendRef references a resource where mirrored requests are sent.  If the referent cannot be found, this BackendRef is invalid and must be dropped from the Gateway. The controller must ensure the 'ResolvedRefs' condition on the Route status is set to 'status: False' and not configure this backend in the underlying implementation.  If there is a cross-namespace reference to an *existing* object that is not allowed by a ReferencePolicy, the controller must ensure the 'ResolvedRefs'  condition on the Route is set to 'status: False', with the 'RefNotPermitted' reason and not configure this backend in the underlying implementation.  In either error case, the Message of the 'ResolvedRefs' Condition should be used to provide more detail about the problem.  Support: Extended for Kubernetes Service Support: Custom for any other resource",
																	MarkdownDescription: "BackendRef references a resource where mirrored requests are sent.  If the referent cannot be found, this BackendRef is invalid and must be dropped from the Gateway. The controller must ensure the 'ResolvedRefs' condition on the Route status is set to 'status: False' and not configure this backend in the underlying implementation.  If there is a cross-namespace reference to an *existing* object that is not allowed by a ReferencePolicy, the controller must ensure the 'ResolvedRefs'  condition on the Route is set to 'status: False', with the 'RefNotPermitted' reason and not configure this backend in the underlying implementation.  In either error case, the Message of the 'ResolvedRefs' Condition should be used to provide more detail about the problem.  Support: Extended for Kubernetes Service Support: Custom for any other resource",
																	Attributes: map[string]schema.Attribute{
																		"group": schema.StringAttribute{
																			Description:         "Group is the group of the referent. For example, 'networking.k8s.io'. When unspecified (empty string), core API group is inferred.",
																			MarkdownDescription: "Group is the group of the referent. For example, 'networking.k8s.io'. When unspecified (empty string), core API group is inferred.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.LengthAtMost(253),
																				stringvalidator.RegexMatches(regexp.MustCompile(`^$|^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
																			},
																		},

																		"kind": schema.StringAttribute{
																			Description:         "Kind is kind of the referent. For example 'HTTPRoute' or 'Service'.",
																			MarkdownDescription: "Kind is kind of the referent. For example 'HTTPRoute' or 'Service'.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.LengthAtLeast(1),
																				stringvalidator.LengthAtMost(63),
																				stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z]([-a-zA-Z0-9]*[a-zA-Z0-9])?$`), ""),
																			},
																		},

																		"name": schema.StringAttribute{
																			Description:         "Name is the name of the referent.",
																			MarkdownDescription: "Name is the name of the referent.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.LengthAtLeast(1),
																				stringvalidator.LengthAtMost(253),
																			},
																		},

																		"namespace": schema.StringAttribute{
																			Description:         "Namespace is the namespace of the backend. When unspecified, the local namespace is inferred.  Note that when a namespace is specified, a ReferencePolicy object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferencePolicy documentation for details.  Support: Core",
																			MarkdownDescription: "Namespace is the namespace of the backend. When unspecified, the local namespace is inferred.  Note that when a namespace is specified, a ReferencePolicy object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferencePolicy documentation for details.  Support: Core",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.LengthAtLeast(1),
																				stringvalidator.LengthAtMost(63),
																				stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
																			},
																		},

																		"port": schema.Int64Attribute{
																			Description:         "Port specifies the destination port number to use for this resource. Port is required when the referent is a Kubernetes Service. For other resources, destination port might be derived from the referent resource or this field.",
																			MarkdownDescription: "Port specifies the destination port number to use for this resource. Port is required when the referent is a Kubernetes Service. For other resources, destination port might be derived from the referent resource or this field.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.Int64{
																				int64validator.AtLeast(1),
																				int64validator.AtMost(65535),
																			},
																		},
																	},
																	Required: true,
																	Optional: false,
																	Computed: false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"request_redirect": schema.SingleNestedAttribute{
															Description:         "RequestRedirect defines a schema for a filter that responds to the request with an HTTP redirection.  Support: Core",
															MarkdownDescription: "RequestRedirect defines a schema for a filter that responds to the request with an HTTP redirection.  Support: Core",
															Attributes: map[string]schema.Attribute{
																"hostname": schema.StringAttribute{
																	Description:         "Hostname is the hostname to be used in the value of the 'Location' header in the response. When empty, the hostname of the request is used.  Support: Core",
																	MarkdownDescription: "Hostname is the hostname to be used in the value of the 'Location' header in the response. When empty, the hostname of the request is used.  Support: Core",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtLeast(1),
																		stringvalidator.LengthAtMost(253),
																		stringvalidator.RegexMatches(regexp.MustCompile(`^(\*\.)?[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
																	},
																},

																"port": schema.Int64Attribute{
																	Description:         "Port is the port to be used in the value of the 'Location' header in the response. When empty, port (if specified) of the request is used.  Support: Extended",
																	MarkdownDescription: "Port is the port to be used in the value of the 'Location' header in the response. When empty, port (if specified) of the request is used.  Support: Extended",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(1),
																		int64validator.AtMost(65535),
																	},
																},

																"scheme": schema.StringAttribute{
																	Description:         "Scheme is the scheme to be used in the value of the 'Location' header in the response. When empty, the scheme of the request is used.  Support: Extended",
																	MarkdownDescription: "Scheme is the scheme to be used in the value of the 'Location' header in the response. When empty, the scheme of the request is used.  Support: Extended",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("http", "https"),
																	},
																},

																"status_code": schema.Int64Attribute{
																	Description:         "StatusCode is the HTTP status code to be used in response.  Support: Core",
																	MarkdownDescription: "StatusCode is the HTTP status code to be used in response.  Support: Core",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.OneOf(301, 302),
																	},
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"type": schema.StringAttribute{
															Description:         "Type identifies the type of filter to apply. As with other API fields, types are classified into three conformance levels:  - Core: Filter types and their corresponding configuration defined by   'Support: Core' in this package, e.g. 'RequestHeaderModifier'. All   implementations must support core filters.  - Extended: Filter types and their corresponding configuration defined by   'Support: Extended' in this package, e.g. 'RequestMirror'. Implementers   are encouraged to support extended filters.  - Custom: Filters that are defined and supported by specific vendors.   In the future, filters showing convergence in behavior across multiple   implementations will be considered for inclusion in extended or core   conformance levels. Filter-specific configuration for such filters   is specified using the ExtensionRef field. 'Type' should be set to   'ExtensionRef' for custom filters.  Implementers are encouraged to define custom implementation types to extend the core API with implementation-specific behavior.  If a reference to a custom filter type cannot be resolved, the filter MUST NOT be skipped. Instead, requests that would have been processed by that filter MUST receive a HTTP error response.",
															MarkdownDescription: "Type identifies the type of filter to apply. As with other API fields, types are classified into three conformance levels:  - Core: Filter types and their corresponding configuration defined by   'Support: Core' in this package, e.g. 'RequestHeaderModifier'. All   implementations must support core filters.  - Extended: Filter types and their corresponding configuration defined by   'Support: Extended' in this package, e.g. 'RequestMirror'. Implementers   are encouraged to support extended filters.  - Custom: Filters that are defined and supported by specific vendors.   In the future, filters showing convergence in behavior across multiple   implementations will be considered for inclusion in extended or core   conformance levels. Filter-specific configuration for such filters   is specified using the ExtensionRef field. 'Type' should be set to   'ExtensionRef' for custom filters.  Implementers are encouraged to define custom implementation types to extend the core API with implementation-specific behavior.  If a reference to a custom filter type cannot be resolved, the filter MUST NOT be skipped. Instead, requests that would have been processed by that filter MUST receive a HTTP error response.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("RequestHeaderModifier", "RequestMirror", "RequestRedirect", "ExtensionRef"),
															},
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"group": schema.StringAttribute{
												Description:         "Group is the group of the referent. For example, 'networking.k8s.io'. When unspecified (empty string), core API group is inferred.",
												MarkdownDescription: "Group is the group of the referent. For example, 'networking.k8s.io'. When unspecified (empty string), core API group is inferred.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtMost(253),
													stringvalidator.RegexMatches(regexp.MustCompile(`^$|^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
												},
											},

											"kind": schema.StringAttribute{
												Description:         "Kind is kind of the referent. For example 'HTTPRoute' or 'Service'.",
												MarkdownDescription: "Kind is kind of the referent. For example 'HTTPRoute' or 'Service'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
													stringvalidator.LengthAtMost(63),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z]([-a-zA-Z0-9]*[a-zA-Z0-9])?$`), ""),
												},
											},

											"name": schema.StringAttribute{
												Description:         "Name is the name of the referent.",
												MarkdownDescription: "Name is the name of the referent.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
													stringvalidator.LengthAtMost(253),
												},
											},

											"namespace": schema.StringAttribute{
												Description:         "Namespace is the namespace of the backend. When unspecified, the local namespace is inferred.  Note that when a namespace is specified, a ReferencePolicy object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferencePolicy documentation for details.  Support: Core",
												MarkdownDescription: "Namespace is the namespace of the backend. When unspecified, the local namespace is inferred.  Note that when a namespace is specified, a ReferencePolicy object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferencePolicy documentation for details.  Support: Core",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
													stringvalidator.LengthAtMost(63),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
												},
											},

											"port": schema.Int64Attribute{
												Description:         "Port specifies the destination port number to use for this resource. Port is required when the referent is a Kubernetes Service. For other resources, destination port might be derived from the referent resource or this field.",
												MarkdownDescription: "Port specifies the destination port number to use for this resource. Port is required when the referent is a Kubernetes Service. For other resources, destination port might be derived from the referent resource or this field.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
													int64validator.AtMost(65535),
												},
											},

											"weight": schema.Int64Attribute{
												Description:         "Weight specifies the proportion of requests forwarded to the referenced backend. This is computed as weight/(sum of all weights in this BackendRefs list). For non-zero values, there may be some epsilon from the exact proportion defined here depending on the precision an implementation supports. Weight is not a percentage and the sum of weights does not need to equal 100.  If only one backend is specified and it has a weight greater than 0, 100% of the traffic is forwarded to that backend. If weight is set to 0, no traffic should be forwarded for this entry. If unspecified, weight defaults to 1.  Support for this field varies based on the context where used.",
												MarkdownDescription: "Weight specifies the proportion of requests forwarded to the referenced backend. This is computed as weight/(sum of all weights in this BackendRefs list). For non-zero values, there may be some epsilon from the exact proportion defined here depending on the precision an implementation supports. Weight is not a percentage and the sum of weights does not need to equal 100.  If only one backend is specified and it has a weight greater than 0, 100% of the traffic is forwarded to that backend. If weight is set to 0, no traffic should be forwarded for this entry. If unspecified, weight defaults to 1.  Support for this field varies based on the context where used.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(1e+06),
												},
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"filters": schema.ListNestedAttribute{
									Description:         "Filters define the filters that are applied to requests that match this rule.  The effects of ordering of multiple behaviors are currently unspecified. This can change in the future based on feedback during the alpha stage.  Conformance-levels at this level are defined based on the type of filter:  - ALL core filters MUST be supported by all implementations. - Implementers are encouraged to support extended filters. - Implementation-specific custom filters have no API guarantees across   implementations.  Specifying a core filter multiple times has unspecified or custom conformance.  Support: Core",
									MarkdownDescription: "Filters define the filters that are applied to requests that match this rule.  The effects of ordering of multiple behaviors are currently unspecified. This can change in the future based on feedback during the alpha stage.  Conformance-levels at this level are defined based on the type of filter:  - ALL core filters MUST be supported by all implementations. - Implementers are encouraged to support extended filters. - Implementation-specific custom filters have no API guarantees across   implementations.  Specifying a core filter multiple times has unspecified or custom conformance.  Support: Core",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"extension_ref": schema.SingleNestedAttribute{
												Description:         "ExtensionRef is an optional, implementation-specific extension to the 'filter' behavior.  For example, resource 'myroutefilter' in group 'networking.example.net'). ExtensionRef MUST NOT be used for core and extended filters.  Support: Implementation-specific",
												MarkdownDescription: "ExtensionRef is an optional, implementation-specific extension to the 'filter' behavior.  For example, resource 'myroutefilter' in group 'networking.example.net'). ExtensionRef MUST NOT be used for core and extended filters.  Support: Implementation-specific",
												Attributes: map[string]schema.Attribute{
													"group": schema.StringAttribute{
														Description:         "Group is the group of the referent. For example, 'networking.k8s.io'. When unspecified (empty string), core API group is inferred.",
														MarkdownDescription: "Group is the group of the referent. For example, 'networking.k8s.io'. When unspecified (empty string), core API group is inferred.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtMost(253),
															stringvalidator.RegexMatches(regexp.MustCompile(`^$|^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
														},
													},

													"kind": schema.StringAttribute{
														Description:         "Kind is kind of the referent. For example 'HTTPRoute' or 'Service'.",
														MarkdownDescription: "Kind is kind of the referent. For example 'HTTPRoute' or 'Service'.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.LengthAtMost(63),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z]([-a-zA-Z0-9]*[a-zA-Z0-9])?$`), ""),
														},
													},

													"name": schema.StringAttribute{
														Description:         "Name is the name of the referent.",
														MarkdownDescription: "Name is the name of the referent.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.LengthAtMost(253),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"request_header_modifier": schema.SingleNestedAttribute{
												Description:         "RequestHeaderModifier defines a schema for a filter that modifies request headers.  Support: Core",
												MarkdownDescription: "RequestHeaderModifier defines a schema for a filter that modifies request headers.  Support: Core",
												Attributes: map[string]schema.Attribute{
													"add": schema.ListNestedAttribute{
														Description:         "Add adds the given header(s) (name, value) to the request before the action. It appends to any existing values associated with the header name.  Input:   GET /foo HTTP/1.1   my-header: foo  Config:   add:   - name: 'my-header'     value: 'bar'  Output:   GET /foo HTTP/1.1   my-header: foo   my-header: bar",
														MarkdownDescription: "Add adds the given header(s) (name, value) to the request before the action. It appends to any existing values associated with the header name.  Input:   GET /foo HTTP/1.1   my-header: foo  Config:   add:   - name: 'my-header'     value: 'bar'  Output:   GET /foo HTTP/1.1   my-header: foo   my-header: bar",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is the name of the HTTP Header to be matched. Name matching MUST be case insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).  If multiple entries specify equivalent header names, the first entry with an equivalent name MUST be considered for a match. Subsequent entries with an equivalent header name MUST be ignored. Due to the case-insensitivity of header names, 'foo' and 'Foo' are considered equivalent.",
																	MarkdownDescription: "Name is the name of the HTTP Header to be matched. Name matching MUST be case insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).  If multiple entries specify equivalent header names, the first entry with an equivalent name MUST be considered for a match. Subsequent entries with an equivalent header name MUST be ignored. Due to the case-insensitivity of header names, 'foo' and 'Foo' are considered equivalent.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtLeast(1),
																		stringvalidator.LengthAtMost(256),
																		stringvalidator.RegexMatches(regexp.MustCompile(`^[A-Za-z0-9!#$%&'*+\-.^_\x60|~]+$`), ""),
																	},
																},

																"value": schema.StringAttribute{
																	Description:         "Value is the value of HTTP Header to be matched.",
																	MarkdownDescription: "Value is the value of HTTP Header to be matched.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtLeast(1),
																		stringvalidator.LengthAtMost(4096),
																	},
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"remove": schema.ListAttribute{
														Description:         "Remove the given header(s) from the HTTP request before the action. The value of Remove is a list of HTTP header names. Note that the header names are case-insensitive (see https://datatracker.ietf.org/doc/html/rfc2616#section-4.2).  Input:   GET /foo HTTP/1.1   my-header1: foo   my-header2: bar   my-header3: baz  Config:   remove: ['my-header1', 'my-header3']  Output:   GET /foo HTTP/1.1   my-header2: bar",
														MarkdownDescription: "Remove the given header(s) from the HTTP request before the action. The value of Remove is a list of HTTP header names. Note that the header names are case-insensitive (see https://datatracker.ietf.org/doc/html/rfc2616#section-4.2).  Input:   GET /foo HTTP/1.1   my-header1: foo   my-header2: bar   my-header3: baz  Config:   remove: ['my-header1', 'my-header3']  Output:   GET /foo HTTP/1.1   my-header2: bar",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"set": schema.ListNestedAttribute{
														Description:         "Set overwrites the request with the given header (name, value) before the action.  Input:   GET /foo HTTP/1.1   my-header: foo  Config:   set:   - name: 'my-header'     value: 'bar'  Output:   GET /foo HTTP/1.1   my-header: bar",
														MarkdownDescription: "Set overwrites the request with the given header (name, value) before the action.  Input:   GET /foo HTTP/1.1   my-header: foo  Config:   set:   - name: 'my-header'     value: 'bar'  Output:   GET /foo HTTP/1.1   my-header: bar",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is the name of the HTTP Header to be matched. Name matching MUST be case insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).  If multiple entries specify equivalent header names, the first entry with an equivalent name MUST be considered for a match. Subsequent entries with an equivalent header name MUST be ignored. Due to the case-insensitivity of header names, 'foo' and 'Foo' are considered equivalent.",
																	MarkdownDescription: "Name is the name of the HTTP Header to be matched. Name matching MUST be case insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).  If multiple entries specify equivalent header names, the first entry with an equivalent name MUST be considered for a match. Subsequent entries with an equivalent header name MUST be ignored. Due to the case-insensitivity of header names, 'foo' and 'Foo' are considered equivalent.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtLeast(1),
																		stringvalidator.LengthAtMost(256),
																		stringvalidator.RegexMatches(regexp.MustCompile(`^[A-Za-z0-9!#$%&'*+\-.^_\x60|~]+$`), ""),
																	},
																},

																"value": schema.StringAttribute{
																	Description:         "Value is the value of HTTP Header to be matched.",
																	MarkdownDescription: "Value is the value of HTTP Header to be matched.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtLeast(1),
																		stringvalidator.LengthAtMost(4096),
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

											"request_mirror": schema.SingleNestedAttribute{
												Description:         "RequestMirror defines a schema for a filter that mirrors requests. Requests are sent to the specified destination, but responses from that destination are ignored.  Support: Extended",
												MarkdownDescription: "RequestMirror defines a schema for a filter that mirrors requests. Requests are sent to the specified destination, but responses from that destination are ignored.  Support: Extended",
												Attributes: map[string]schema.Attribute{
													"backend_ref": schema.SingleNestedAttribute{
														Description:         "BackendRef references a resource where mirrored requests are sent.  If the referent cannot be found, this BackendRef is invalid and must be dropped from the Gateway. The controller must ensure the 'ResolvedRefs' condition on the Route status is set to 'status: False' and not configure this backend in the underlying implementation.  If there is a cross-namespace reference to an *existing* object that is not allowed by a ReferencePolicy, the controller must ensure the 'ResolvedRefs'  condition on the Route is set to 'status: False', with the 'RefNotPermitted' reason and not configure this backend in the underlying implementation.  In either error case, the Message of the 'ResolvedRefs' Condition should be used to provide more detail about the problem.  Support: Extended for Kubernetes Service Support: Custom for any other resource",
														MarkdownDescription: "BackendRef references a resource where mirrored requests are sent.  If the referent cannot be found, this BackendRef is invalid and must be dropped from the Gateway. The controller must ensure the 'ResolvedRefs' condition on the Route status is set to 'status: False' and not configure this backend in the underlying implementation.  If there is a cross-namespace reference to an *existing* object that is not allowed by a ReferencePolicy, the controller must ensure the 'ResolvedRefs'  condition on the Route is set to 'status: False', with the 'RefNotPermitted' reason and not configure this backend in the underlying implementation.  In either error case, the Message of the 'ResolvedRefs' Condition should be used to provide more detail about the problem.  Support: Extended for Kubernetes Service Support: Custom for any other resource",
														Attributes: map[string]schema.Attribute{
															"group": schema.StringAttribute{
																Description:         "Group is the group of the referent. For example, 'networking.k8s.io'. When unspecified (empty string), core API group is inferred.",
																MarkdownDescription: "Group is the group of the referent. For example, 'networking.k8s.io'. When unspecified (empty string), core API group is inferred.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtMost(253),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^$|^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
																},
															},

															"kind": schema.StringAttribute{
																Description:         "Kind is kind of the referent. For example 'HTTPRoute' or 'Service'.",
																MarkdownDescription: "Kind is kind of the referent. For example 'HTTPRoute' or 'Service'.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																	stringvalidator.LengthAtMost(63),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z]([-a-zA-Z0-9]*[a-zA-Z0-9])?$`), ""),
																},
															},

															"name": schema.StringAttribute{
																Description:         "Name is the name of the referent.",
																MarkdownDescription: "Name is the name of the referent.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																	stringvalidator.LengthAtMost(253),
																},
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace is the namespace of the backend. When unspecified, the local namespace is inferred.  Note that when a namespace is specified, a ReferencePolicy object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferencePolicy documentation for details.  Support: Core",
																MarkdownDescription: "Namespace is the namespace of the backend. When unspecified, the local namespace is inferred.  Note that when a namespace is specified, a ReferencePolicy object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferencePolicy documentation for details.  Support: Core",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																	stringvalidator.LengthAtMost(63),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
																},
															},

															"port": schema.Int64Attribute{
																Description:         "Port specifies the destination port number to use for this resource. Port is required when the referent is a Kubernetes Service. For other resources, destination port might be derived from the referent resource or this field.",
																MarkdownDescription: "Port specifies the destination port number to use for this resource. Port is required when the referent is a Kubernetes Service. For other resources, destination port might be derived from the referent resource or this field.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(1),
																	int64validator.AtMost(65535),
																},
															},
														},
														Required: true,
														Optional: false,
														Computed: false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"request_redirect": schema.SingleNestedAttribute{
												Description:         "RequestRedirect defines a schema for a filter that responds to the request with an HTTP redirection.  Support: Core",
												MarkdownDescription: "RequestRedirect defines a schema for a filter that responds to the request with an HTTP redirection.  Support: Core",
												Attributes: map[string]schema.Attribute{
													"hostname": schema.StringAttribute{
														Description:         "Hostname is the hostname to be used in the value of the 'Location' header in the response. When empty, the hostname of the request is used.  Support: Core",
														MarkdownDescription: "Hostname is the hostname to be used in the value of the 'Location' header in the response. When empty, the hostname of the request is used.  Support: Core",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.LengthAtMost(253),
															stringvalidator.RegexMatches(regexp.MustCompile(`^(\*\.)?[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
														},
													},

													"port": schema.Int64Attribute{
														Description:         "Port is the port to be used in the value of the 'Location' header in the response. When empty, port (if specified) of the request is used.  Support: Extended",
														MarkdownDescription: "Port is the port to be used in the value of the 'Location' header in the response. When empty, port (if specified) of the request is used.  Support: Extended",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(1),
															int64validator.AtMost(65535),
														},
													},

													"scheme": schema.StringAttribute{
														Description:         "Scheme is the scheme to be used in the value of the 'Location' header in the response. When empty, the scheme of the request is used.  Support: Extended",
														MarkdownDescription: "Scheme is the scheme to be used in the value of the 'Location' header in the response. When empty, the scheme of the request is used.  Support: Extended",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("http", "https"),
														},
													},

													"status_code": schema.Int64Attribute{
														Description:         "StatusCode is the HTTP status code to be used in response.  Support: Core",
														MarkdownDescription: "StatusCode is the HTTP status code to be used in response.  Support: Core",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.OneOf(301, 302),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"type": schema.StringAttribute{
												Description:         "Type identifies the type of filter to apply. As with other API fields, types are classified into three conformance levels:  - Core: Filter types and their corresponding configuration defined by   'Support: Core' in this package, e.g. 'RequestHeaderModifier'. All   implementations must support core filters.  - Extended: Filter types and their corresponding configuration defined by   'Support: Extended' in this package, e.g. 'RequestMirror'. Implementers   are encouraged to support extended filters.  - Custom: Filters that are defined and supported by specific vendors.   In the future, filters showing convergence in behavior across multiple   implementations will be considered for inclusion in extended or core   conformance levels. Filter-specific configuration for such filters   is specified using the ExtensionRef field. 'Type' should be set to   'ExtensionRef' for custom filters.  Implementers are encouraged to define custom implementation types to extend the core API with implementation-specific behavior.  If a reference to a custom filter type cannot be resolved, the filter MUST NOT be skipped. Instead, requests that would have been processed by that filter MUST receive a HTTP error response.",
												MarkdownDescription: "Type identifies the type of filter to apply. As with other API fields, types are classified into three conformance levels:  - Core: Filter types and their corresponding configuration defined by   'Support: Core' in this package, e.g. 'RequestHeaderModifier'. All   implementations must support core filters.  - Extended: Filter types and their corresponding configuration defined by   'Support: Extended' in this package, e.g. 'RequestMirror'. Implementers   are encouraged to support extended filters.  - Custom: Filters that are defined and supported by specific vendors.   In the future, filters showing convergence in behavior across multiple   implementations will be considered for inclusion in extended or core   conformance levels. Filter-specific configuration for such filters   is specified using the ExtensionRef field. 'Type' should be set to   'ExtensionRef' for custom filters.  Implementers are encouraged to define custom implementation types to extend the core API with implementation-specific behavior.  If a reference to a custom filter type cannot be resolved, the filter MUST NOT be skipped. Instead, requests that would have been processed by that filter MUST receive a HTTP error response.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("RequestHeaderModifier", "RequestMirror", "RequestRedirect", "ExtensionRef"),
												},
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"matches": schema.ListNestedAttribute{
									Description:         "Matches define conditions used for matching the rule against incoming HTTP requests. Each match is independent, i.e. this rule will be matched if **any** one of the matches is satisfied.  For example, take the following matches configuration:  ''' matches: - path:     value: '/foo'   headers:   - name: 'version'     value: 'v2' - path:     value: '/v2/foo' '''  For a request to match against this rule, a request must satisfy EITHER of the two conditions:  - path prefixed with '/foo' AND contains the header 'version: v2' - path prefix of '/v2/foo'  See the documentation for HTTPRouteMatch on how to specify multiple match conditions that should be ANDed together.  If no matches are specified, the default is a prefix path match on '/', which has the effect of matching every HTTP request.  Proxy or Load Balancer routing configuration generated from HTTPRoutes MUST prioritize rules based on the following criteria, continuing on ties. Precedence must be given to the the Rule with the largest number of:  * Characters in a matching non-wildcard hostname. * Characters in a matching hostname. * Characters in a matching path. * Header matches. * Query param matches.  If ties still exist across multiple Routes, matching precedence MUST be determined in order of the following criteria, continuing on ties:  * The oldest Route based on creation timestamp. * The Route appearing first in alphabetical order by   '<namespace>/<name>'.  If ties still exist within the Route that has been given precedence, matching precedence MUST be granted to the first matching rule meeting the above criteria.",
									MarkdownDescription: "Matches define conditions used for matching the rule against incoming HTTP requests. Each match is independent, i.e. this rule will be matched if **any** one of the matches is satisfied.  For example, take the following matches configuration:  ''' matches: - path:     value: '/foo'   headers:   - name: 'version'     value: 'v2' - path:     value: '/v2/foo' '''  For a request to match against this rule, a request must satisfy EITHER of the two conditions:  - path prefixed with '/foo' AND contains the header 'version: v2' - path prefix of '/v2/foo'  See the documentation for HTTPRouteMatch on how to specify multiple match conditions that should be ANDed together.  If no matches are specified, the default is a prefix path match on '/', which has the effect of matching every HTTP request.  Proxy or Load Balancer routing configuration generated from HTTPRoutes MUST prioritize rules based on the following criteria, continuing on ties. Precedence must be given to the the Rule with the largest number of:  * Characters in a matching non-wildcard hostname. * Characters in a matching hostname. * Characters in a matching path. * Header matches. * Query param matches.  If ties still exist across multiple Routes, matching precedence MUST be determined in order of the following criteria, continuing on ties:  * The oldest Route based on creation timestamp. * The Route appearing first in alphabetical order by   '<namespace>/<name>'.  If ties still exist within the Route that has been given precedence, matching precedence MUST be granted to the first matching rule meeting the above criteria.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"headers": schema.ListNestedAttribute{
												Description:         "Headers specifies HTTP request header matchers. Multiple match values are ANDed together, meaning, a request must match all the specified headers to select the route.",
												MarkdownDescription: "Headers specifies HTTP request header matchers. Multiple match values are ANDed together, meaning, a request must match all the specified headers to select the route.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name is the name of the HTTP Header to be matched. Name matching MUST be case insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).  If multiple entries specify equivalent header names, only the first entry with an equivalent name MUST be considered for a match. Subsequent entries with an equivalent header name MUST be ignored. Due to the case-insensitivity of header names, 'foo' and 'Foo' are considered equivalent.  When a header is repeated in an HTTP request, it is implementation-specific behavior as to how this is represented. Generally, proxies should follow the guidance from the RFC: https://www.rfc-editor.org/rfc/rfc7230.html#section-3.2.2 regarding processing a repeated header, with special handling for 'Set-Cookie'.",
															MarkdownDescription: "Name is the name of the HTTP Header to be matched. Name matching MUST be case insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).  If multiple entries specify equivalent header names, only the first entry with an equivalent name MUST be considered for a match. Subsequent entries with an equivalent header name MUST be ignored. Due to the case-insensitivity of header names, 'foo' and 'Foo' are considered equivalent.  When a header is repeated in an HTTP request, it is implementation-specific behavior as to how this is represented. Generally, proxies should follow the guidance from the RFC: https://www.rfc-editor.org/rfc/rfc7230.html#section-3.2.2 regarding processing a repeated header, with special handling for 'Set-Cookie'.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtLeast(1),
																stringvalidator.LengthAtMost(256),
																stringvalidator.RegexMatches(regexp.MustCompile(`^[A-Za-z0-9!#$%&'*+\-.^_\x60|~]+$`), ""),
															},
														},

														"type": schema.StringAttribute{
															Description:         "Type specifies how to match against the value of the header.  Support: Core (Exact)  Support: Custom (RegularExpression)  Since RegularExpression HeaderMatchType has custom conformance, implementations can support POSIX, PCRE or any other dialects of regular expressions. Please read the implementation's documentation to determine the supported dialect.",
															MarkdownDescription: "Type specifies how to match against the value of the header.  Support: Core (Exact)  Support: Custom (RegularExpression)  Since RegularExpression HeaderMatchType has custom conformance, implementations can support POSIX, PCRE or any other dialects of regular expressions. Please read the implementation's documentation to determine the supported dialect.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("Exact", "RegularExpression"),
															},
														},

														"value": schema.StringAttribute{
															Description:         "Value is the value of HTTP Header to be matched.",
															MarkdownDescription: "Value is the value of HTTP Header to be matched.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtLeast(1),
																stringvalidator.LengthAtMost(4096),
															},
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"method": schema.StringAttribute{
												Description:         "Method specifies HTTP method matcher. When specified, this route will be matched only if the request has the specified method.  Support: Extended",
												MarkdownDescription: "Method specifies HTTP method matcher. When specified, this route will be matched only if the request has the specified method.  Support: Extended",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH"),
												},
											},

											"path": schema.SingleNestedAttribute{
												Description:         "Path specifies a HTTP request path matcher. If this field is not specified, a default prefix match on the '/' path is provided.",
												MarkdownDescription: "Path specifies a HTTP request path matcher. If this field is not specified, a default prefix match on the '/' path is provided.",
												Attributes: map[string]schema.Attribute{
													"type": schema.StringAttribute{
														Description:         "Type specifies how to match against the path Value.  Support: Core (Exact, PathPrefix)  Support: Custom (RegularExpression)",
														MarkdownDescription: "Type specifies how to match against the path Value.  Support: Core (Exact, PathPrefix)  Support: Custom (RegularExpression)",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("Exact", "PathPrefix", "RegularExpression"),
														},
													},

													"value": schema.StringAttribute{
														Description:         "Value of the HTTP path to match against.",
														MarkdownDescription: "Value of the HTTP path to match against.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtMost(1024),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"query_params": schema.ListNestedAttribute{
												Description:         "QueryParams specifies HTTP query parameter matchers. Multiple match values are ANDed together, meaning, a request must match all the specified query parameters to select the route.",
												MarkdownDescription: "QueryParams specifies HTTP query parameter matchers. Multiple match values are ANDed together, meaning, a request must match all the specified query parameters to select the route.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name is the name of the HTTP query param to be matched. This must be an exact string match. (See https://tools.ietf.org/html/rfc7230#section-2.7.3).",
															MarkdownDescription: "Name is the name of the HTTP query param to be matched. This must be an exact string match. (See https://tools.ietf.org/html/rfc7230#section-2.7.3).",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtLeast(1),
																stringvalidator.LengthAtMost(256),
															},
														},

														"type": schema.StringAttribute{
															Description:         "Type specifies how to match against the value of the query parameter.  Support: Extended (Exact)  Support: Custom (RegularExpression)  Since RegularExpression QueryParamMatchType has custom conformance, implementations can support POSIX, PCRE or any other dialects of regular expressions. Please read the implementation's documentation to determine the supported dialect.",
															MarkdownDescription: "Type specifies how to match against the value of the query parameter.  Support: Extended (Exact)  Support: Custom (RegularExpression)  Since RegularExpression QueryParamMatchType has custom conformance, implementations can support POSIX, PCRE or any other dialects of regular expressions. Please read the implementation's documentation to determine the supported dialect.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("Exact", "RegularExpression"),
															},
														},

														"value": schema.StringAttribute{
															Description:         "Value is the value of HTTP query param to be matched.",
															MarkdownDescription: "Value is the value of HTTP query param to be matched.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtLeast(1),
																stringvalidator.LengthAtMost(1024),
															},
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

func (r *GatewayNetworkingK8SIoHTTPRouteV1Alpha2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_gateway_networking_k8s_io_http_route_v1alpha2_manifest")

	var model GatewayNetworkingK8SIoHTTPRouteV1Alpha2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("gateway.networking.k8s.io/v1alpha2")
	model.Kind = pointer.String("HTTPRoute")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"YAML Error: "+err.Error(),
		)
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
