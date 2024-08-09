/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package gateway_networking_k8s_io_v1beta1

import (
	"context"
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
	_ datasource.DataSource = &GatewayNetworkingK8SIoHttprouteV1Beta1Manifest{}
)

func NewGatewayNetworkingK8SIoHttprouteV1Beta1Manifest() datasource.DataSource {
	return &GatewayNetworkingK8SIoHttprouteV1Beta1Manifest{}
}

type GatewayNetworkingK8SIoHttprouteV1Beta1Manifest struct{}

type GatewayNetworkingK8SIoHttprouteV1Beta1ManifestData struct {
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
			Port        *int64  `tfsdk:"port" json:"port,omitempty"`
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
						Hostname *string `tfsdk:"hostname" json:"hostname,omitempty"`
						Path     *struct {
							ReplaceFullPath    *string `tfsdk:"replace_full_path" json:"replaceFullPath,omitempty"`
							ReplacePrefixMatch *string `tfsdk:"replace_prefix_match" json:"replacePrefixMatch,omitempty"`
							Type               *string `tfsdk:"type" json:"type,omitempty"`
						} `tfsdk:"path" json:"path,omitempty"`
						Port       *int64  `tfsdk:"port" json:"port,omitempty"`
						Scheme     *string `tfsdk:"scheme" json:"scheme,omitempty"`
						StatusCode *int64  `tfsdk:"status_code" json:"statusCode,omitempty"`
					} `tfsdk:"request_redirect" json:"requestRedirect,omitempty"`
					ResponseHeaderModifier *struct {
						Add *[]struct {
							Name  *string `tfsdk:"name" json:"name,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"add" json:"add,omitempty"`
						Remove *[]string `tfsdk:"remove" json:"remove,omitempty"`
						Set    *[]struct {
							Name  *string `tfsdk:"name" json:"name,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"set" json:"set,omitempty"`
					} `tfsdk:"response_header_modifier" json:"responseHeaderModifier,omitempty"`
					Type       *string `tfsdk:"type" json:"type,omitempty"`
					UrlRewrite *struct {
						Hostname *string `tfsdk:"hostname" json:"hostname,omitempty"`
						Path     *struct {
							ReplaceFullPath    *string `tfsdk:"replace_full_path" json:"replaceFullPath,omitempty"`
							ReplacePrefixMatch *string `tfsdk:"replace_prefix_match" json:"replacePrefixMatch,omitempty"`
							Type               *string `tfsdk:"type" json:"type,omitempty"`
						} `tfsdk:"path" json:"path,omitempty"`
					} `tfsdk:"url_rewrite" json:"urlRewrite,omitempty"`
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
					Hostname *string `tfsdk:"hostname" json:"hostname,omitempty"`
					Path     *struct {
						ReplaceFullPath    *string `tfsdk:"replace_full_path" json:"replaceFullPath,omitempty"`
						ReplacePrefixMatch *string `tfsdk:"replace_prefix_match" json:"replacePrefixMatch,omitempty"`
						Type               *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"path" json:"path,omitempty"`
					Port       *int64  `tfsdk:"port" json:"port,omitempty"`
					Scheme     *string `tfsdk:"scheme" json:"scheme,omitempty"`
					StatusCode *int64  `tfsdk:"status_code" json:"statusCode,omitempty"`
				} `tfsdk:"request_redirect" json:"requestRedirect,omitempty"`
				ResponseHeaderModifier *struct {
					Add *[]struct {
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"add" json:"add,omitempty"`
					Remove *[]string `tfsdk:"remove" json:"remove,omitempty"`
					Set    *[]struct {
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"set" json:"set,omitempty"`
				} `tfsdk:"response_header_modifier" json:"responseHeaderModifier,omitempty"`
				Type       *string `tfsdk:"type" json:"type,omitempty"`
				UrlRewrite *struct {
					Hostname *string `tfsdk:"hostname" json:"hostname,omitempty"`
					Path     *struct {
						ReplaceFullPath    *string `tfsdk:"replace_full_path" json:"replaceFullPath,omitempty"`
						ReplacePrefixMatch *string `tfsdk:"replace_prefix_match" json:"replacePrefixMatch,omitempty"`
						Type               *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"path" json:"path,omitempty"`
				} `tfsdk:"url_rewrite" json:"urlRewrite,omitempty"`
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
			SessionPersistence *struct {
				AbsoluteTimeout *string `tfsdk:"absolute_timeout" json:"absoluteTimeout,omitempty"`
				CookieConfig    *struct {
					LifetimeType *string `tfsdk:"lifetime_type" json:"lifetimeType,omitempty"`
				} `tfsdk:"cookie_config" json:"cookieConfig,omitempty"`
				IdleTimeout *string `tfsdk:"idle_timeout" json:"idleTimeout,omitempty"`
				SessionName *string `tfsdk:"session_name" json:"sessionName,omitempty"`
				Type        *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"session_persistence" json:"sessionPersistence,omitempty"`
			Timeouts *struct {
				BackendRequest *string `tfsdk:"backend_request" json:"backendRequest,omitempty"`
				Request        *string `tfsdk:"request" json:"request,omitempty"`
			} `tfsdk:"timeouts" json:"timeouts,omitempty"`
		} `tfsdk:"rules" json:"rules,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *GatewayNetworkingK8SIoHttprouteV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_gateway_networking_k8s_io_http_route_v1beta1_manifest"
}

func (r *GatewayNetworkingK8SIoHttprouteV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "HTTPRoute provides a way to route HTTP requests. This includes the capabilityto match requests by hostname, path, header, or query param. Filters can beused to specify additional processing steps. Backends specify where matchingrequests should be routed.",
		MarkdownDescription: "HTTPRoute provides a way to route HTTP requests. This includes the capabilityto match requests by hostname, path, header, or query param. Filters can beused to specify additional processing steps. Backends specify where matchingrequests should be routed.",
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
				Description:         "Spec defines the desired state of HTTPRoute.",
				MarkdownDescription: "Spec defines the desired state of HTTPRoute.",
				Attributes: map[string]schema.Attribute{
					"hostnames": schema.ListAttribute{
						Description:         "Hostnames defines a set of hostnames that should match against the HTTP Hostheader to select a HTTPRoute used to process the request. ImplementationsMUST ignore any port value specified in the HTTP Host header whileperforming a match and (absent of any applicable header modificationconfiguration) MUST forward this header unmodified to the backend.Valid values for Hostnames are determined by RFC 1123 definition of ahostname with 2 notable exceptions:1. IPs are not allowed.2. A hostname may be prefixed with a wildcard label ('*.'). The wildcard   label must appear by itself as the first label.If a hostname is specified by both the Listener and HTTPRoute, theremust be at least one intersecting hostname for the HTTPRoute to beattached to the Listener. For example:* A Listener with 'test.example.com' as the hostname matches HTTPRoutes  that have either not specified any hostnames, or have specified at  least one of 'test.example.com' or '*.example.com'.* A Listener with '*.example.com' as the hostname matches HTTPRoutes  that have either not specified any hostnames or have specified at least  one hostname that matches the Listener hostname. For example,  '*.example.com', 'test.example.com', and 'foo.test.example.com' would  all match. On the other hand, 'example.com' and 'test.example.net' would  not match.Hostnames that are prefixed with a wildcard label ('*.') are interpretedas a suffix match. That means that a match for '*.example.com' would matchboth 'test.example.com', and 'foo.test.example.com', but not 'example.com'.If both the Listener and HTTPRoute have specified hostnames, anyHTTPRoute hostnames that do not match the Listener hostname MUST beignored. For example, if a Listener specified '*.example.com', and theHTTPRoute specified 'test.example.com' and 'test.example.net','test.example.net' must not be considered for a match.If both the Listener and HTTPRoute have specified hostnames, and nonematch with the criteria above, then the HTTPRoute is not accepted. Theimplementation must raise an 'Accepted' Condition with a status of'False' in the corresponding RouteParentStatus.In the event that multiple HTTPRoutes specify intersecting hostnames (e.g.overlapping wildcard matching and exact matching hostnames), precedence mustbe given to rules from the HTTPRoute with the largest number of:* Characters in a matching non-wildcard hostname.* Characters in a matching hostname.If ties exist across multiple Routes, the matching precedence rules forHTTPRouteMatches takes over.Support: Core",
						MarkdownDescription: "Hostnames defines a set of hostnames that should match against the HTTP Hostheader to select a HTTPRoute used to process the request. ImplementationsMUST ignore any port value specified in the HTTP Host header whileperforming a match and (absent of any applicable header modificationconfiguration) MUST forward this header unmodified to the backend.Valid values for Hostnames are determined by RFC 1123 definition of ahostname with 2 notable exceptions:1. IPs are not allowed.2. A hostname may be prefixed with a wildcard label ('*.'). The wildcard   label must appear by itself as the first label.If a hostname is specified by both the Listener and HTTPRoute, theremust be at least one intersecting hostname for the HTTPRoute to beattached to the Listener. For example:* A Listener with 'test.example.com' as the hostname matches HTTPRoutes  that have either not specified any hostnames, or have specified at  least one of 'test.example.com' or '*.example.com'.* A Listener with '*.example.com' as the hostname matches HTTPRoutes  that have either not specified any hostnames or have specified at least  one hostname that matches the Listener hostname. For example,  '*.example.com', 'test.example.com', and 'foo.test.example.com' would  all match. On the other hand, 'example.com' and 'test.example.net' would  not match.Hostnames that are prefixed with a wildcard label ('*.') are interpretedas a suffix match. That means that a match for '*.example.com' would matchboth 'test.example.com', and 'foo.test.example.com', but not 'example.com'.If both the Listener and HTTPRoute have specified hostnames, anyHTTPRoute hostnames that do not match the Listener hostname MUST beignored. For example, if a Listener specified '*.example.com', and theHTTPRoute specified 'test.example.com' and 'test.example.net','test.example.net' must not be considered for a match.If both the Listener and HTTPRoute have specified hostnames, and nonematch with the criteria above, then the HTTPRoute is not accepted. Theimplementation must raise an 'Accepted' Condition with a status of'False' in the corresponding RouteParentStatus.In the event that multiple HTTPRoutes specify intersecting hostnames (e.g.overlapping wildcard matching and exact matching hostnames), precedence mustbe given to rules from the HTTPRoute with the largest number of:* Characters in a matching non-wildcard hostname.* Characters in a matching hostname.If ties exist across multiple Routes, the matching precedence rules forHTTPRouteMatches takes over.Support: Core",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"parent_refs": schema.ListNestedAttribute{
						Description:         "ParentRefs references the resources (usually Gateways) that a Route wantsto be attached to. Note that the referenced parent resource needs toallow this for the attachment to be complete. For Gateways, that meansthe Gateway needs to allow attachment from Routes of this kind andnamespace. For Services, that means the Service must either be in the samenamespace for a 'producer' route, or the mesh implementation must supportand allow 'consumer' routes for the referenced Service. ReferenceGrant isnot applicable for governing ParentRefs to Services - it is not possible tocreate a 'producer' route for a Service in a different namespace from theRoute.There are two kinds of parent resources with 'Core' support:* Gateway (Gateway conformance profile)* Service (Mesh conformance profile, ClusterIP Services only)This API may be extended in the future to support additional kinds of parentresources.ParentRefs must be _distinct_. This means either that:* They select different objects.  If this is the case, then parentRef  entries are distinct. In terms of fields, this means that the  multi-part key defined by 'group', 'kind', 'namespace', and 'name' must  be unique across all parentRef entries in the Route.* They do not select different objects, but for each optional field used,  each ParentRef that selects the same object must set the same set of  optional fields to different values. If one ParentRef sets a  combination of optional fields, all must set the same combination.Some examples:* If one ParentRef sets 'sectionName', all ParentRefs referencing the  same object must also set 'sectionName'.* If one ParentRef sets 'port', all ParentRefs referencing the same  object must also set 'port'.* If one ParentRef sets 'sectionName' and 'port', all ParentRefs  referencing the same object must also set 'sectionName' and 'port'.It is possible to separately reference multiple distinct objects that maybe collapsed by an implementation. For example, some implementations maychoose to merge compatible Gateway Listeners together. If that is thecase, the list of routes attached to those resources should also bemerged.Note that for ParentRefs that cross namespace boundaries, there are specificrules. Cross-namespace references are only valid if they are explicitlyallowed by something in the namespace they are referring to. For example,Gateway has the AllowedRoutes field, and ReferenceGrant provides ageneric way to enable other kinds of cross-namespace reference.ParentRefs from a Route to a Service in the same namespace are 'producer'routes, which apply default routing rules to inbound connections fromany namespace to the Service.ParentRefs from a Route to a Service in a different namespace are'consumer' routes, and these routing rules are only applied to outboundconnections originating from the same namespace as the Route, for whichthe intended destination of the connections are a Service targeted as aParentRef of the Route.",
						MarkdownDescription: "ParentRefs references the resources (usually Gateways) that a Route wantsto be attached to. Note that the referenced parent resource needs toallow this for the attachment to be complete. For Gateways, that meansthe Gateway needs to allow attachment from Routes of this kind andnamespace. For Services, that means the Service must either be in the samenamespace for a 'producer' route, or the mesh implementation must supportand allow 'consumer' routes for the referenced Service. ReferenceGrant isnot applicable for governing ParentRefs to Services - it is not possible tocreate a 'producer' route for a Service in a different namespace from theRoute.There are two kinds of parent resources with 'Core' support:* Gateway (Gateway conformance profile)* Service (Mesh conformance profile, ClusterIP Services only)This API may be extended in the future to support additional kinds of parentresources.ParentRefs must be _distinct_. This means either that:* They select different objects.  If this is the case, then parentRef  entries are distinct. In terms of fields, this means that the  multi-part key defined by 'group', 'kind', 'namespace', and 'name' must  be unique across all parentRef entries in the Route.* They do not select different objects, but for each optional field used,  each ParentRef that selects the same object must set the same set of  optional fields to different values. If one ParentRef sets a  combination of optional fields, all must set the same combination.Some examples:* If one ParentRef sets 'sectionName', all ParentRefs referencing the  same object must also set 'sectionName'.* If one ParentRef sets 'port', all ParentRefs referencing the same  object must also set 'port'.* If one ParentRef sets 'sectionName' and 'port', all ParentRefs  referencing the same object must also set 'sectionName' and 'port'.It is possible to separately reference multiple distinct objects that maybe collapsed by an implementation. For example, some implementations maychoose to merge compatible Gateway Listeners together. If that is thecase, the list of routes attached to those resources should also bemerged.Note that for ParentRefs that cross namespace boundaries, there are specificrules. Cross-namespace references are only valid if they are explicitlyallowed by something in the namespace they are referring to. For example,Gateway has the AllowedRoutes field, and ReferenceGrant provides ageneric way to enable other kinds of cross-namespace reference.ParentRefs from a Route to a Service in the same namespace are 'producer'routes, which apply default routing rules to inbound connections fromany namespace to the Service.ParentRefs from a Route to a Service in a different namespace are'consumer' routes, and these routing rules are only applied to outboundconnections originating from the same namespace as the Route, for whichthe intended destination of the connections are a Service targeted as aParentRef of the Route.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"group": schema.StringAttribute{
									Description:         "Group is the group of the referent.When unspecified, 'gateway.networking.k8s.io' is inferred.To set the core API group (such as for a 'Service' kind referent),Group must be explicitly set to '' (empty string).Support: Core",
									MarkdownDescription: "Group is the group of the referent.When unspecified, 'gateway.networking.k8s.io' is inferred.To set the core API group (such as for a 'Service' kind referent),Group must be explicitly set to '' (empty string).Support: Core",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(253),
										stringvalidator.RegexMatches(regexp.MustCompile(`^$|^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
									},
								},

								"kind": schema.StringAttribute{
									Description:         "Kind is kind of the referent.There are two kinds of parent resources with 'Core' support:* Gateway (Gateway conformance profile)* Service (Mesh conformance profile, ClusterIP Services only)Support for other resources is Implementation-Specific.",
									MarkdownDescription: "Kind is kind of the referent.There are two kinds of parent resources with 'Core' support:* Gateway (Gateway conformance profile)* Service (Mesh conformance profile, ClusterIP Services only)Support for other resources is Implementation-Specific.",
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
									Description:         "Name is the name of the referent.Support: Core",
									MarkdownDescription: "Name is the name of the referent.Support: Core",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(253),
									},
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace is the namespace of the referent. When unspecified, this refersto the local namespace of the Route.Note that there are specific rules for ParentRefs which cross namespaceboundaries. Cross-namespace references are only valid if they are explicitlyallowed by something in the namespace they are referring to. For example:Gateway has the AllowedRoutes field, and ReferenceGrant provides ageneric way to enable any other kind of cross-namespace reference.ParentRefs from a Route to a Service in the same namespace are 'producer'routes, which apply default routing rules to inbound connections fromany namespace to the Service.ParentRefs from a Route to a Service in a different namespace are'consumer' routes, and these routing rules are only applied to outboundconnections originating from the same namespace as the Route, for whichthe intended destination of the connections are a Service targeted as aParentRef of the Route.Support: Core",
									MarkdownDescription: "Namespace is the namespace of the referent. When unspecified, this refersto the local namespace of the Route.Note that there are specific rules for ParentRefs which cross namespaceboundaries. Cross-namespace references are only valid if they are explicitlyallowed by something in the namespace they are referring to. For example:Gateway has the AllowedRoutes field, and ReferenceGrant provides ageneric way to enable any other kind of cross-namespace reference.ParentRefs from a Route to a Service in the same namespace are 'producer'routes, which apply default routing rules to inbound connections fromany namespace to the Service.ParentRefs from a Route to a Service in a different namespace are'consumer' routes, and these routing rules are only applied to outboundconnections originating from the same namespace as the Route, for whichthe intended destination of the connections are a Service targeted as aParentRef of the Route.Support: Core",
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
									Description:         "Port is the network port this Route targets. It can be interpreteddifferently based on the type of parent resource.When the parent resource is a Gateway, this targets all listenerslistening on the specified port that also support this kind of Route(andselect this Route). It's not recommended to set 'Port' unless thenetworking behaviors specified in a Route must apply to a specific portas opposed to a listener(s) whose port(s) may be changed. When both Portand SectionName are specified, the name and port of the selected listenermust match both specified values.When the parent resource is a Service, this targets a specific port in theService spec. When both Port (experimental) and SectionName are specified,the name and port of the selected port must match both specified values.Implementations MAY choose to support other parent resources.Implementations supporting other types of parent resources MUST clearlydocument how/if Port is interpreted.For the purpose of status, an attachment is considered successful aslong as the parent resource accepts it partially. For example, Gatewaylisteners can restrict which Routes can attach to them by Route kind,namespace, or hostname. If 1 of 2 Gateway listeners accept attachmentfrom the referencing Route, the Route MUST be considered successfullyattached. If no Gateway listeners accept attachment from this Route,the Route MUST be considered detached from the Gateway.Support: Extended",
									MarkdownDescription: "Port is the network port this Route targets. It can be interpreteddifferently based on the type of parent resource.When the parent resource is a Gateway, this targets all listenerslistening on the specified port that also support this kind of Route(andselect this Route). It's not recommended to set 'Port' unless thenetworking behaviors specified in a Route must apply to a specific portas opposed to a listener(s) whose port(s) may be changed. When both Portand SectionName are specified, the name and port of the selected listenermust match both specified values.When the parent resource is a Service, this targets a specific port in theService spec. When both Port (experimental) and SectionName are specified,the name and port of the selected port must match both specified values.Implementations MAY choose to support other parent resources.Implementations supporting other types of parent resources MUST clearlydocument how/if Port is interpreted.For the purpose of status, an attachment is considered successful aslong as the parent resource accepts it partially. For example, Gatewaylisteners can restrict which Routes can attach to them by Route kind,namespace, or hostname. If 1 of 2 Gateway listeners accept attachmentfrom the referencing Route, the Route MUST be considered successfullyattached. If no Gateway listeners accept attachment from this Route,the Route MUST be considered detached from the Gateway.Support: Extended",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.Int64{
										int64validator.AtLeast(1),
										int64validator.AtMost(65535),
									},
								},

								"section_name": schema.StringAttribute{
									Description:         "SectionName is the name of a section within the target resource. In thefollowing resources, SectionName is interpreted as the following:* Gateway: Listener name. When both Port (experimental) and SectionNameare specified, the name and port of the selected listener must matchboth specified values.* Service: Port name. When both Port (experimental) and SectionNameare specified, the name and port of the selected listener must matchboth specified values.Implementations MAY choose to support attaching Routes to other resources.If that is the case, they MUST clearly document how SectionName isinterpreted.When unspecified (empty string), this will reference the entire resource.For the purpose of status, an attachment is considered successful if atleast one section in the parent resource accepts it. For example, Gatewaylisteners can restrict which Routes can attach to them by Route kind,namespace, or hostname. If 1 of 2 Gateway listeners accept attachment fromthe referencing Route, the Route MUST be considered successfullyattached. If no Gateway listeners accept attachment from this Route, theRoute MUST be considered detached from the Gateway.Support: Core",
									MarkdownDescription: "SectionName is the name of a section within the target resource. In thefollowing resources, SectionName is interpreted as the following:* Gateway: Listener name. When both Port (experimental) and SectionNameare specified, the name and port of the selected listener must matchboth specified values.* Service: Port name. When both Port (experimental) and SectionNameare specified, the name and port of the selected listener must matchboth specified values.Implementations MAY choose to support attaching Routes to other resources.If that is the case, they MUST clearly document how SectionName isinterpreted.When unspecified (empty string), this will reference the entire resource.For the purpose of status, an attachment is considered successful if atleast one section in the parent resource accepts it. For example, Gatewaylisteners can restrict which Routes can attach to them by Route kind,namespace, or hostname. If 1 of 2 Gateway listeners accept attachment fromthe referencing Route, the Route MUST be considered successfullyattached. If no Gateway listeners accept attachment from this Route, theRoute MUST be considered detached from the Gateway.Support: Core",
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
									Description:         "BackendRefs defines the backend(s) where matching requests should besent.Failure behavior here depends on how many BackendRefs are specified andhow many are invalid.If *all* entries in BackendRefs are invalid, and there are also no filtersspecified in this route rule, *all* traffic which matches this rule MUSTreceive a 500 status code.See the HTTPBackendRef definition for the rules about what makes a singleHTTPBackendRef invalid.When a HTTPBackendRef is invalid, 500 status codes MUST be returned forrequests that would have otherwise been routed to an invalid backend. Ifmultiple backends are specified, and some are invalid, the proportion ofrequests that would otherwise have been routed to an invalid backendMUST receive a 500 status code.For example, if two backends are specified with equal weights, and one isinvalid, 50 percent of traffic must receive a 500. Implementations maychoose how that 50 percent is determined.When a HTTPBackendRef refers to a Service that has no ready endpoints,implementations SHOULD return a 503 for requests to that backend instead.If an implementation chooses to do this, all of the above rules for 500 responsesMUST also apply for responses that return a 503.Support: Core for Kubernetes ServiceSupport: Extended for Kubernetes ServiceImportSupport: Implementation-specific for any other resourceSupport for weight: Core",
									MarkdownDescription: "BackendRefs defines the backend(s) where matching requests should besent.Failure behavior here depends on how many BackendRefs are specified andhow many are invalid.If *all* entries in BackendRefs are invalid, and there are also no filtersspecified in this route rule, *all* traffic which matches this rule MUSTreceive a 500 status code.See the HTTPBackendRef definition for the rules about what makes a singleHTTPBackendRef invalid.When a HTTPBackendRef is invalid, 500 status codes MUST be returned forrequests that would have otherwise been routed to an invalid backend. Ifmultiple backends are specified, and some are invalid, the proportion ofrequests that would otherwise have been routed to an invalid backendMUST receive a 500 status code.For example, if two backends are specified with equal weights, and one isinvalid, 50 percent of traffic must receive a 500. Implementations maychoose how that 50 percent is determined.When a HTTPBackendRef refers to a Service that has no ready endpoints,implementations SHOULD return a 503 for requests to that backend instead.If an implementation chooses to do this, all of the above rules for 500 responsesMUST also apply for responses that return a 503.Support: Core for Kubernetes ServiceSupport: Extended for Kubernetes ServiceImportSupport: Implementation-specific for any other resourceSupport for weight: Core",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"filters": schema.ListNestedAttribute{
												Description:         "Filters defined at this level should be executed if and only if therequest is being forwarded to the backend defined here.Support: Implementation-specific (For broader support of filters, use theFilters field in HTTPRouteRule.)",
												MarkdownDescription: "Filters defined at this level should be executed if and only if therequest is being forwarded to the backend defined here.Support: Implementation-specific (For broader support of filters, use theFilters field in HTTPRouteRule.)",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"extension_ref": schema.SingleNestedAttribute{
															Description:         "ExtensionRef is an optional, implementation-specific extension to the'filter' behavior.  For example, resource 'myroutefilter' in group'networking.example.net'). ExtensionRef MUST NOT be used for core andextended filters.This filter can be used multiple times within the same rule.Support: Implementation-specific",
															MarkdownDescription: "ExtensionRef is an optional, implementation-specific extension to the'filter' behavior.  For example, resource 'myroutefilter' in group'networking.example.net'). ExtensionRef MUST NOT be used for core andextended filters.This filter can be used multiple times within the same rule.Support: Implementation-specific",
															Attributes: map[string]schema.Attribute{
																"group": schema.StringAttribute{
																	Description:         "Group is the group of the referent. For example, 'gateway.networking.k8s.io'.When unspecified or empty string, core API group is inferred.",
																	MarkdownDescription: "Group is the group of the referent. For example, 'gateway.networking.k8s.io'.When unspecified or empty string, core API group is inferred.",
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
															Description:         "RequestHeaderModifier defines a schema for a filter that modifies requestheaders.Support: Core",
															MarkdownDescription: "RequestHeaderModifier defines a schema for a filter that modifies requestheaders.Support: Core",
															Attributes: map[string]schema.Attribute{
																"add": schema.ListNestedAttribute{
																	Description:         "Add adds the given header(s) (name, value) to the requestbefore the action. It appends to any existing values associatedwith the header name.Input:  GET /foo HTTP/1.1  my-header: fooConfig:  add:  - name: 'my-header'    value: 'bar,baz'Output:  GET /foo HTTP/1.1  my-header: foo,bar,baz",
																	MarkdownDescription: "Add adds the given header(s) (name, value) to the requestbefore the action. It appends to any existing values associatedwith the header name.Input:  GET /foo HTTP/1.1  my-header: fooConfig:  add:  - name: 'my-header'    value: 'bar,baz'Output:  GET /foo HTTP/1.1  my-header: foo,bar,baz",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "Name is the name of the HTTP Header to be matched. Name matching MUST becase insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).If multiple entries specify equivalent header names, the first entry withan equivalent name MUST be considered for a match. Subsequent entrieswith an equivalent header name MUST be ignored. Due to thecase-insensitivity of header names, 'foo' and 'Foo' are consideredequivalent.",
																				MarkdownDescription: "Name is the name of the HTTP Header to be matched. Name matching MUST becase insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).If multiple entries specify equivalent header names, the first entry withan equivalent name MUST be considered for a match. Subsequent entrieswith an equivalent header name MUST be ignored. Due to thecase-insensitivity of header names, 'foo' and 'Foo' are consideredequivalent.",
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
																	Description:         "Remove the given header(s) from the HTTP request before the action. Thevalue of Remove is a list of HTTP header names. Note that the headernames are case-insensitive (seehttps://datatracker.ietf.org/doc/html/rfc2616#section-4.2).Input:  GET /foo HTTP/1.1  my-header1: foo  my-header2: bar  my-header3: bazConfig:  remove: ['my-header1', 'my-header3']Output:  GET /foo HTTP/1.1  my-header2: bar",
																	MarkdownDescription: "Remove the given header(s) from the HTTP request before the action. Thevalue of Remove is a list of HTTP header names. Note that the headernames are case-insensitive (seehttps://datatracker.ietf.org/doc/html/rfc2616#section-4.2).Input:  GET /foo HTTP/1.1  my-header1: foo  my-header2: bar  my-header3: bazConfig:  remove: ['my-header1', 'my-header3']Output:  GET /foo HTTP/1.1  my-header2: bar",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"set": schema.ListNestedAttribute{
																	Description:         "Set overwrites the request with the given header (name, value)before the action.Input:  GET /foo HTTP/1.1  my-header: fooConfig:  set:  - name: 'my-header'    value: 'bar'Output:  GET /foo HTTP/1.1  my-header: bar",
																	MarkdownDescription: "Set overwrites the request with the given header (name, value)before the action.Input:  GET /foo HTTP/1.1  my-header: fooConfig:  set:  - name: 'my-header'    value: 'bar'Output:  GET /foo HTTP/1.1  my-header: bar",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "Name is the name of the HTTP Header to be matched. Name matching MUST becase insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).If multiple entries specify equivalent header names, the first entry withan equivalent name MUST be considered for a match. Subsequent entrieswith an equivalent header name MUST be ignored. Due to thecase-insensitivity of header names, 'foo' and 'Foo' are consideredequivalent.",
																				MarkdownDescription: "Name is the name of the HTTP Header to be matched. Name matching MUST becase insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).If multiple entries specify equivalent header names, the first entry withan equivalent name MUST be considered for a match. Subsequent entrieswith an equivalent header name MUST be ignored. Due to thecase-insensitivity of header names, 'foo' and 'Foo' are consideredequivalent.",
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
															Description:         "RequestMirror defines a schema for a filter that mirrors requests.Requests are sent to the specified destination, but responses fromthat destination are ignored.This filter can be used multiple times within the same rule. Note thatnot all implementations will be able to support mirroring to multiplebackends.Support: Extended",
															MarkdownDescription: "RequestMirror defines a schema for a filter that mirrors requests.Requests are sent to the specified destination, but responses fromthat destination are ignored.This filter can be used multiple times within the same rule. Note thatnot all implementations will be able to support mirroring to multiplebackends.Support: Extended",
															Attributes: map[string]schema.Attribute{
																"backend_ref": schema.SingleNestedAttribute{
																	Description:         "BackendRef references a resource where mirrored requests are sent.Mirrored requests must be sent only to a single destination endpointwithin this BackendRef, irrespective of how many endpoints are presentwithin this BackendRef.If the referent cannot be found, this BackendRef is invalid and must bedropped from the Gateway. The controller must ensure the 'ResolvedRefs'condition on the Route status is set to 'status: False' and not configurethis backend in the underlying implementation.If there is a cross-namespace reference to an *existing* objectthat is not allowed by a ReferenceGrant, the controller must ensure the'ResolvedRefs'  condition on the Route is set to 'status: False',with the 'RefNotPermitted' reason and not configure this backend in theunderlying implementation.In either error case, the Message of the 'ResolvedRefs' Conditionshould be used to provide more detail about the problem.Support: Extended for Kubernetes ServiceSupport: Implementation-specific for any other resource",
																	MarkdownDescription: "BackendRef references a resource where mirrored requests are sent.Mirrored requests must be sent only to a single destination endpointwithin this BackendRef, irrespective of how many endpoints are presentwithin this BackendRef.If the referent cannot be found, this BackendRef is invalid and must bedropped from the Gateway. The controller must ensure the 'ResolvedRefs'condition on the Route status is set to 'status: False' and not configurethis backend in the underlying implementation.If there is a cross-namespace reference to an *existing* objectthat is not allowed by a ReferenceGrant, the controller must ensure the'ResolvedRefs'  condition on the Route is set to 'status: False',with the 'RefNotPermitted' reason and not configure this backend in theunderlying implementation.In either error case, the Message of the 'ResolvedRefs' Conditionshould be used to provide more detail about the problem.Support: Extended for Kubernetes ServiceSupport: Implementation-specific for any other resource",
																	Attributes: map[string]schema.Attribute{
																		"group": schema.StringAttribute{
																			Description:         "Group is the group of the referent. For example, 'gateway.networking.k8s.io'.When unspecified or empty string, core API group is inferred.",
																			MarkdownDescription: "Group is the group of the referent. For example, 'gateway.networking.k8s.io'.When unspecified or empty string, core API group is inferred.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.LengthAtMost(253),
																				stringvalidator.RegexMatches(regexp.MustCompile(`^$|^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
																			},
																		},

																		"kind": schema.StringAttribute{
																			Description:         "Kind is the Kubernetes resource kind of the referent. For example'Service'.Defaults to 'Service' when not specified.ExternalName services can refer to CNAME DNS records that may liveoutside of the cluster and as such are difficult to reason about interms of conformance. They also may not be safe to forward to (seeCVE-2021-25740 for more information). Implementations SHOULD NOTsupport ExternalName Services.Support: Core (Services with a type other than ExternalName)Support: Implementation-specific (Services with type ExternalName)",
																			MarkdownDescription: "Kind is the Kubernetes resource kind of the referent. For example'Service'.Defaults to 'Service' when not specified.ExternalName services can refer to CNAME DNS records that may liveoutside of the cluster and as such are difficult to reason about interms of conformance. They also may not be safe to forward to (seeCVE-2021-25740 for more information). Implementations SHOULD NOTsupport ExternalName Services.Support: Core (Services with a type other than ExternalName)Support: Implementation-specific (Services with type ExternalName)",
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
																			Description:         "Namespace is the namespace of the backend. When unspecified, the localnamespace is inferred.Note that when a namespace different than the local namespace is specified,a ReferenceGrant object is required in the referent namespace to allow thatnamespace's owner to accept the reference. See the ReferenceGrantdocumentation for details.Support: Core",
																			MarkdownDescription: "Namespace is the namespace of the backend. When unspecified, the localnamespace is inferred.Note that when a namespace different than the local namespace is specified,a ReferenceGrant object is required in the referent namespace to allow thatnamespace's owner to accept the reference. See the ReferenceGrantdocumentation for details.Support: Core",
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
																			Description:         "Port specifies the destination port number to use for this resource.Port is required when the referent is a Kubernetes Service. In thiscase, the port number is the service port number, not the target port.For other resources, destination port might be derived from the referentresource or this field.",
																			MarkdownDescription: "Port specifies the destination port number to use for this resource.Port is required when the referent is a Kubernetes Service. In thiscase, the port number is the service port number, not the target port.For other resources, destination port might be derived from the referentresource or this field.",
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
															Description:         "RequestRedirect defines a schema for a filter that responds to therequest with an HTTP redirection.Support: Core",
															MarkdownDescription: "RequestRedirect defines a schema for a filter that responds to therequest with an HTTP redirection.Support: Core",
															Attributes: map[string]schema.Attribute{
																"hostname": schema.StringAttribute{
																	Description:         "Hostname is the hostname to be used in the value of the 'Location'header in the response.When empty, the hostname in the 'Host' header of the request is used.Support: Core",
																	MarkdownDescription: "Hostname is the hostname to be used in the value of the 'Location'header in the response.When empty, the hostname in the 'Host' header of the request is used.Support: Core",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtLeast(1),
																		stringvalidator.LengthAtMost(253),
																		stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
																	},
																},

																"path": schema.SingleNestedAttribute{
																	Description:         "Path defines parameters used to modify the path of the incoming request.The modified path is then used to construct the 'Location' header. Whenempty, the request path is used as-is.Support: Extended",
																	MarkdownDescription: "Path defines parameters used to modify the path of the incoming request.The modified path is then used to construct the 'Location' header. Whenempty, the request path is used as-is.Support: Extended",
																	Attributes: map[string]schema.Attribute{
																		"replace_full_path": schema.StringAttribute{
																			Description:         "ReplaceFullPath specifies the value with which to replace the full pathof a request during a rewrite or redirect.",
																			MarkdownDescription: "ReplaceFullPath specifies the value with which to replace the full pathof a request during a rewrite or redirect.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.LengthAtMost(1024),
																			},
																		},

																		"replace_prefix_match": schema.StringAttribute{
																			Description:         "ReplacePrefixMatch specifies the value with which to replace the prefixmatch of a request during a rewrite or redirect. For example, a requestto '/foo/bar' with a prefix match of '/foo' and a ReplacePrefixMatchof '/xyz' would be modified to '/xyz/bar'.Note that this matches the behavior of the PathPrefix match type. Thismatches full path elements. A path element refers to the list of labelsin the path split by the '/' separator. When specified, a trailing '/' isignored. For example, the paths '/abc', '/abc/', and '/abc/def' would allmatch the prefix '/abc', but the path '/abcd' would not.ReplacePrefixMatch is only compatible with a 'PathPrefix' HTTPRouteMatch.Using any other HTTPRouteMatch type on the same HTTPRouteRule will result inthe implementation setting the Accepted Condition for the Route to 'status: False'.Request Path | Prefix Match | Replace Prefix | Modified Path-------------|--------------|----------------|----------/foo/bar     | /foo         | /xyz           | /xyz/bar/foo/bar     | /foo         | /xyz/          | /xyz/bar/foo/bar     | /foo/        | /xyz           | /xyz/bar/foo/bar     | /foo/        | /xyz/          | /xyz/bar/foo         | /foo         | /xyz           | /xyz/foo/        | /foo         | /xyz           | /xyz//foo/bar     | /foo         | <empty string> | /bar/foo/        | /foo         | <empty string> | //foo         | /foo         | <empty string> | //foo/        | /foo         | /              | //foo         | /foo         | /              | /",
																			MarkdownDescription: "ReplacePrefixMatch specifies the value with which to replace the prefixmatch of a request during a rewrite or redirect. For example, a requestto '/foo/bar' with a prefix match of '/foo' and a ReplacePrefixMatchof '/xyz' would be modified to '/xyz/bar'.Note that this matches the behavior of the PathPrefix match type. Thismatches full path elements. A path element refers to the list of labelsin the path split by the '/' separator. When specified, a trailing '/' isignored. For example, the paths '/abc', '/abc/', and '/abc/def' would allmatch the prefix '/abc', but the path '/abcd' would not.ReplacePrefixMatch is only compatible with a 'PathPrefix' HTTPRouteMatch.Using any other HTTPRouteMatch type on the same HTTPRouteRule will result inthe implementation setting the Accepted Condition for the Route to 'status: False'.Request Path | Prefix Match | Replace Prefix | Modified Path-------------|--------------|----------------|----------/foo/bar     | /foo         | /xyz           | /xyz/bar/foo/bar     | /foo         | /xyz/          | /xyz/bar/foo/bar     | /foo/        | /xyz           | /xyz/bar/foo/bar     | /foo/        | /xyz/          | /xyz/bar/foo         | /foo         | /xyz           | /xyz/foo/        | /foo         | /xyz           | /xyz//foo/bar     | /foo         | <empty string> | /bar/foo/        | /foo         | <empty string> | //foo         | /foo         | <empty string> | //foo/        | /foo         | /              | //foo         | /foo         | /              | /",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.LengthAtMost(1024),
																			},
																		},

																		"type": schema.StringAttribute{
																			Description:         "Type defines the type of path modifier. Additional types may beadded in a future release of the API.Note that values may be added to this enum, implementationsmust ensure that unknown values will not cause a crash.Unknown values here must result in the implementation setting theAccepted Condition for the Route to 'status: False', with aReason of 'UnsupportedValue'.",
																			MarkdownDescription: "Type defines the type of path modifier. Additional types may beadded in a future release of the API.Note that values may be added to this enum, implementationsmust ensure that unknown values will not cause a crash.Unknown values here must result in the implementation setting theAccepted Condition for the Route to 'status: False', with aReason of 'UnsupportedValue'.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.OneOf("ReplaceFullPath", "ReplacePrefixMatch"),
																			},
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"port": schema.Int64Attribute{
																	Description:         "Port is the port to be used in the value of the 'Location'header in the response.If no port is specified, the redirect port MUST be derived using thefollowing rules:* If redirect scheme is not-empty, the redirect port MUST be the well-known  port associated with the redirect scheme. Specifically 'http' to port 80  and 'https' to port 443. If the redirect scheme does not have a  well-known port, the listener port of the Gateway SHOULD be used.* If redirect scheme is empty, the redirect port MUST be the Gateway  Listener port.Implementations SHOULD NOT add the port number in the 'Location'header in the following cases:* A Location header that will use HTTP (whether that is determined via  the Listener protocol or the Scheme field) _and_ use port 80.* A Location header that will use HTTPS (whether that is determined via  the Listener protocol or the Scheme field) _and_ use port 443.Support: Extended",
																	MarkdownDescription: "Port is the port to be used in the value of the 'Location'header in the response.If no port is specified, the redirect port MUST be derived using thefollowing rules:* If redirect scheme is not-empty, the redirect port MUST be the well-known  port associated with the redirect scheme. Specifically 'http' to port 80  and 'https' to port 443. If the redirect scheme does not have a  well-known port, the listener port of the Gateway SHOULD be used.* If redirect scheme is empty, the redirect port MUST be the Gateway  Listener port.Implementations SHOULD NOT add the port number in the 'Location'header in the following cases:* A Location header that will use HTTP (whether that is determined via  the Listener protocol or the Scheme field) _and_ use port 80.* A Location header that will use HTTPS (whether that is determined via  the Listener protocol or the Scheme field) _and_ use port 443.Support: Extended",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(1),
																		int64validator.AtMost(65535),
																	},
																},

																"scheme": schema.StringAttribute{
																	Description:         "Scheme is the scheme to be used in the value of the 'Location' header inthe response. When empty, the scheme of the request is used.Scheme redirects can affect the port of the redirect, for more information,refer to the documentation for the port field of this filter.Note that values may be added to this enum, implementationsmust ensure that unknown values will not cause a crash.Unknown values here must result in the implementation setting theAccepted Condition for the Route to 'status: False', with aReason of 'UnsupportedValue'.Support: Extended",
																	MarkdownDescription: "Scheme is the scheme to be used in the value of the 'Location' header inthe response. When empty, the scheme of the request is used.Scheme redirects can affect the port of the redirect, for more information,refer to the documentation for the port field of this filter.Note that values may be added to this enum, implementationsmust ensure that unknown values will not cause a crash.Unknown values here must result in the implementation setting theAccepted Condition for the Route to 'status: False', with aReason of 'UnsupportedValue'.Support: Extended",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("http", "https"),
																	},
																},

																"status_code": schema.Int64Attribute{
																	Description:         "StatusCode is the HTTP status code to be used in response.Note that values may be added to this enum, implementationsmust ensure that unknown values will not cause a crash.Unknown values here must result in the implementation setting theAccepted Condition for the Route to 'status: False', with aReason of 'UnsupportedValue'.Support: Core",
																	MarkdownDescription: "StatusCode is the HTTP status code to be used in response.Note that values may be added to this enum, implementationsmust ensure that unknown values will not cause a crash.Unknown values here must result in the implementation setting theAccepted Condition for the Route to 'status: False', with aReason of 'UnsupportedValue'.Support: Core",
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

														"response_header_modifier": schema.SingleNestedAttribute{
															Description:         "ResponseHeaderModifier defines a schema for a filter that modifies responseheaders.Support: Extended",
															MarkdownDescription: "ResponseHeaderModifier defines a schema for a filter that modifies responseheaders.Support: Extended",
															Attributes: map[string]schema.Attribute{
																"add": schema.ListNestedAttribute{
																	Description:         "Add adds the given header(s) (name, value) to the requestbefore the action. It appends to any existing values associatedwith the header name.Input:  GET /foo HTTP/1.1  my-header: fooConfig:  add:  - name: 'my-header'    value: 'bar,baz'Output:  GET /foo HTTP/1.1  my-header: foo,bar,baz",
																	MarkdownDescription: "Add adds the given header(s) (name, value) to the requestbefore the action. It appends to any existing values associatedwith the header name.Input:  GET /foo HTTP/1.1  my-header: fooConfig:  add:  - name: 'my-header'    value: 'bar,baz'Output:  GET /foo HTTP/1.1  my-header: foo,bar,baz",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "Name is the name of the HTTP Header to be matched. Name matching MUST becase insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).If multiple entries specify equivalent header names, the first entry withan equivalent name MUST be considered for a match. Subsequent entrieswith an equivalent header name MUST be ignored. Due to thecase-insensitivity of header names, 'foo' and 'Foo' are consideredequivalent.",
																				MarkdownDescription: "Name is the name of the HTTP Header to be matched. Name matching MUST becase insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).If multiple entries specify equivalent header names, the first entry withan equivalent name MUST be considered for a match. Subsequent entrieswith an equivalent header name MUST be ignored. Due to thecase-insensitivity of header names, 'foo' and 'Foo' are consideredequivalent.",
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
																	Description:         "Remove the given header(s) from the HTTP request before the action. Thevalue of Remove is a list of HTTP header names. Note that the headernames are case-insensitive (seehttps://datatracker.ietf.org/doc/html/rfc2616#section-4.2).Input:  GET /foo HTTP/1.1  my-header1: foo  my-header2: bar  my-header3: bazConfig:  remove: ['my-header1', 'my-header3']Output:  GET /foo HTTP/1.1  my-header2: bar",
																	MarkdownDescription: "Remove the given header(s) from the HTTP request before the action. Thevalue of Remove is a list of HTTP header names. Note that the headernames are case-insensitive (seehttps://datatracker.ietf.org/doc/html/rfc2616#section-4.2).Input:  GET /foo HTTP/1.1  my-header1: foo  my-header2: bar  my-header3: bazConfig:  remove: ['my-header1', 'my-header3']Output:  GET /foo HTTP/1.1  my-header2: bar",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"set": schema.ListNestedAttribute{
																	Description:         "Set overwrites the request with the given header (name, value)before the action.Input:  GET /foo HTTP/1.1  my-header: fooConfig:  set:  - name: 'my-header'    value: 'bar'Output:  GET /foo HTTP/1.1  my-header: bar",
																	MarkdownDescription: "Set overwrites the request with the given header (name, value)before the action.Input:  GET /foo HTTP/1.1  my-header: fooConfig:  set:  - name: 'my-header'    value: 'bar'Output:  GET /foo HTTP/1.1  my-header: bar",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "Name is the name of the HTTP Header to be matched. Name matching MUST becase insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).If multiple entries specify equivalent header names, the first entry withan equivalent name MUST be considered for a match. Subsequent entrieswith an equivalent header name MUST be ignored. Due to thecase-insensitivity of header names, 'foo' and 'Foo' are consideredequivalent.",
																				MarkdownDescription: "Name is the name of the HTTP Header to be matched. Name matching MUST becase insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).If multiple entries specify equivalent header names, the first entry withan equivalent name MUST be considered for a match. Subsequent entrieswith an equivalent header name MUST be ignored. Due to thecase-insensitivity of header names, 'foo' and 'Foo' are consideredequivalent.",
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

														"type": schema.StringAttribute{
															Description:         "Type identifies the type of filter to apply. As with other API fields,types are classified into three conformance levels:- Core: Filter types and their corresponding configuration defined by  'Support: Core' in this package, e.g. 'RequestHeaderModifier'. All  implementations must support core filters.- Extended: Filter types and their corresponding configuration defined by  'Support: Extended' in this package, e.g. 'RequestMirror'. Implementers  are encouraged to support extended filters.- Implementation-specific: Filters that are defined and supported by  specific vendors.  In the future, filters showing convergence in behavior across multiple  implementations will be considered for inclusion in extended or core  conformance levels. Filter-specific configuration for such filters  is specified using the ExtensionRef field. 'Type' should be set to  'ExtensionRef' for custom filters.Implementers are encouraged to define custom implementation types toextend the core API with implementation-specific behavior.If a reference to a custom filter type cannot be resolved, the filterMUST NOT be skipped. Instead, requests that would have been processed bythat filter MUST receive a HTTP error response.Note that values may be added to this enum, implementationsmust ensure that unknown values will not cause a crash.Unknown values here must result in the implementation setting theAccepted Condition for the Route to 'status: False', with aReason of 'UnsupportedValue'.",
															MarkdownDescription: "Type identifies the type of filter to apply. As with other API fields,types are classified into three conformance levels:- Core: Filter types and their corresponding configuration defined by  'Support: Core' in this package, e.g. 'RequestHeaderModifier'. All  implementations must support core filters.- Extended: Filter types and their corresponding configuration defined by  'Support: Extended' in this package, e.g. 'RequestMirror'. Implementers  are encouraged to support extended filters.- Implementation-specific: Filters that are defined and supported by  specific vendors.  In the future, filters showing convergence in behavior across multiple  implementations will be considered for inclusion in extended or core  conformance levels. Filter-specific configuration for such filters  is specified using the ExtensionRef field. 'Type' should be set to  'ExtensionRef' for custom filters.Implementers are encouraged to define custom implementation types toextend the core API with implementation-specific behavior.If a reference to a custom filter type cannot be resolved, the filterMUST NOT be skipped. Instead, requests that would have been processed bythat filter MUST receive a HTTP error response.Note that values may be added to this enum, implementationsmust ensure that unknown values will not cause a crash.Unknown values here must result in the implementation setting theAccepted Condition for the Route to 'status: False', with aReason of 'UnsupportedValue'.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("RequestHeaderModifier", "ResponseHeaderModifier", "RequestMirror", "RequestRedirect", "URLRewrite", "ExtensionRef"),
															},
														},

														"url_rewrite": schema.SingleNestedAttribute{
															Description:         "URLRewrite defines a schema for a filter that modifies a request during forwarding.Support: Extended",
															MarkdownDescription: "URLRewrite defines a schema for a filter that modifies a request during forwarding.Support: Extended",
															Attributes: map[string]schema.Attribute{
																"hostname": schema.StringAttribute{
																	Description:         "Hostname is the value to be used to replace the Host header value duringforwarding.Support: Extended",
																	MarkdownDescription: "Hostname is the value to be used to replace the Host header value duringforwarding.Support: Extended",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtLeast(1),
																		stringvalidator.LengthAtMost(253),
																		stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
																	},
																},

																"path": schema.SingleNestedAttribute{
																	Description:         "Path defines a path rewrite.Support: Extended",
																	MarkdownDescription: "Path defines a path rewrite.Support: Extended",
																	Attributes: map[string]schema.Attribute{
																		"replace_full_path": schema.StringAttribute{
																			Description:         "ReplaceFullPath specifies the value with which to replace the full pathof a request during a rewrite or redirect.",
																			MarkdownDescription: "ReplaceFullPath specifies the value with which to replace the full pathof a request during a rewrite or redirect.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.LengthAtMost(1024),
																			},
																		},

																		"replace_prefix_match": schema.StringAttribute{
																			Description:         "ReplacePrefixMatch specifies the value with which to replace the prefixmatch of a request during a rewrite or redirect. For example, a requestto '/foo/bar' with a prefix match of '/foo' and a ReplacePrefixMatchof '/xyz' would be modified to '/xyz/bar'.Note that this matches the behavior of the PathPrefix match type. Thismatches full path elements. A path element refers to the list of labelsin the path split by the '/' separator. When specified, a trailing '/' isignored. For example, the paths '/abc', '/abc/', and '/abc/def' would allmatch the prefix '/abc', but the path '/abcd' would not.ReplacePrefixMatch is only compatible with a 'PathPrefix' HTTPRouteMatch.Using any other HTTPRouteMatch type on the same HTTPRouteRule will result inthe implementation setting the Accepted Condition for the Route to 'status: False'.Request Path | Prefix Match | Replace Prefix | Modified Path-------------|--------------|----------------|----------/foo/bar     | /foo         | /xyz           | /xyz/bar/foo/bar     | /foo         | /xyz/          | /xyz/bar/foo/bar     | /foo/        | /xyz           | /xyz/bar/foo/bar     | /foo/        | /xyz/          | /xyz/bar/foo         | /foo         | /xyz           | /xyz/foo/        | /foo         | /xyz           | /xyz//foo/bar     | /foo         | <empty string> | /bar/foo/        | /foo         | <empty string> | //foo         | /foo         | <empty string> | //foo/        | /foo         | /              | //foo         | /foo         | /              | /",
																			MarkdownDescription: "ReplacePrefixMatch specifies the value with which to replace the prefixmatch of a request during a rewrite or redirect. For example, a requestto '/foo/bar' with a prefix match of '/foo' and a ReplacePrefixMatchof '/xyz' would be modified to '/xyz/bar'.Note that this matches the behavior of the PathPrefix match type. Thismatches full path elements. A path element refers to the list of labelsin the path split by the '/' separator. When specified, a trailing '/' isignored. For example, the paths '/abc', '/abc/', and '/abc/def' would allmatch the prefix '/abc', but the path '/abcd' would not.ReplacePrefixMatch is only compatible with a 'PathPrefix' HTTPRouteMatch.Using any other HTTPRouteMatch type on the same HTTPRouteRule will result inthe implementation setting the Accepted Condition for the Route to 'status: False'.Request Path | Prefix Match | Replace Prefix | Modified Path-------------|--------------|----------------|----------/foo/bar     | /foo         | /xyz           | /xyz/bar/foo/bar     | /foo         | /xyz/          | /xyz/bar/foo/bar     | /foo/        | /xyz           | /xyz/bar/foo/bar     | /foo/        | /xyz/          | /xyz/bar/foo         | /foo         | /xyz           | /xyz/foo/        | /foo         | /xyz           | /xyz//foo/bar     | /foo         | <empty string> | /bar/foo/        | /foo         | <empty string> | //foo         | /foo         | <empty string> | //foo/        | /foo         | /              | //foo         | /foo         | /              | /",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.LengthAtMost(1024),
																			},
																		},

																		"type": schema.StringAttribute{
																			Description:         "Type defines the type of path modifier. Additional types may beadded in a future release of the API.Note that values may be added to this enum, implementationsmust ensure that unknown values will not cause a crash.Unknown values here must result in the implementation setting theAccepted Condition for the Route to 'status: False', with aReason of 'UnsupportedValue'.",
																			MarkdownDescription: "Type defines the type of path modifier. Additional types may beadded in a future release of the API.Note that values may be added to this enum, implementationsmust ensure that unknown values will not cause a crash.Unknown values here must result in the implementation setting theAccepted Condition for the Route to 'status: False', with aReason of 'UnsupportedValue'.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.OneOf("ReplaceFullPath", "ReplacePrefixMatch"),
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
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"group": schema.StringAttribute{
												Description:         "Group is the group of the referent. For example, 'gateway.networking.k8s.io'.When unspecified or empty string, core API group is inferred.",
												MarkdownDescription: "Group is the group of the referent. For example, 'gateway.networking.k8s.io'.When unspecified or empty string, core API group is inferred.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtMost(253),
													stringvalidator.RegexMatches(regexp.MustCompile(`^$|^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
												},
											},

											"kind": schema.StringAttribute{
												Description:         "Kind is the Kubernetes resource kind of the referent. For example'Service'.Defaults to 'Service' when not specified.ExternalName services can refer to CNAME DNS records that may liveoutside of the cluster and as such are difficult to reason about interms of conformance. They also may not be safe to forward to (seeCVE-2021-25740 for more information). Implementations SHOULD NOTsupport ExternalName Services.Support: Core (Services with a type other than ExternalName)Support: Implementation-specific (Services with type ExternalName)",
												MarkdownDescription: "Kind is the Kubernetes resource kind of the referent. For example'Service'.Defaults to 'Service' when not specified.ExternalName services can refer to CNAME DNS records that may liveoutside of the cluster and as such are difficult to reason about interms of conformance. They also may not be safe to forward to (seeCVE-2021-25740 for more information). Implementations SHOULD NOTsupport ExternalName Services.Support: Core (Services with a type other than ExternalName)Support: Implementation-specific (Services with type ExternalName)",
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
												Description:         "Namespace is the namespace of the backend. When unspecified, the localnamespace is inferred.Note that when a namespace different than the local namespace is specified,a ReferenceGrant object is required in the referent namespace to allow thatnamespace's owner to accept the reference. See the ReferenceGrantdocumentation for details.Support: Core",
												MarkdownDescription: "Namespace is the namespace of the backend. When unspecified, the localnamespace is inferred.Note that when a namespace different than the local namespace is specified,a ReferenceGrant object is required in the referent namespace to allow thatnamespace's owner to accept the reference. See the ReferenceGrantdocumentation for details.Support: Core",
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
												Description:         "Port specifies the destination port number to use for this resource.Port is required when the referent is a Kubernetes Service. In thiscase, the port number is the service port number, not the target port.For other resources, destination port might be derived from the referentresource or this field.",
												MarkdownDescription: "Port specifies the destination port number to use for this resource.Port is required when the referent is a Kubernetes Service. In thiscase, the port number is the service port number, not the target port.For other resources, destination port might be derived from the referentresource or this field.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
													int64validator.AtMost(65535),
												},
											},

											"weight": schema.Int64Attribute{
												Description:         "Weight specifies the proportion of requests forwarded to the referencedbackend. This is computed as weight/(sum of all weights in thisBackendRefs list). For non-zero values, there may be some epsilon fromthe exact proportion defined here depending on the precision animplementation supports. Weight is not a percentage and the sum ofweights does not need to equal 100.If only one backend is specified and it has a weight greater than 0, 100%of the traffic is forwarded to that backend. If weight is set to 0, notraffic should be forwarded for this entry. If unspecified, weightdefaults to 1.Support for this field varies based on the context where used.",
												MarkdownDescription: "Weight specifies the proportion of requests forwarded to the referencedbackend. This is computed as weight/(sum of all weights in thisBackendRefs list). For non-zero values, there may be some epsilon fromthe exact proportion defined here depending on the precision animplementation supports. Weight is not a percentage and the sum ofweights does not need to equal 100.If only one backend is specified and it has a weight greater than 0, 100%of the traffic is forwarded to that backend. If weight is set to 0, notraffic should be forwarded for this entry. If unspecified, weightdefaults to 1.Support for this field varies based on the context where used.",
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
									Description:         "Filters define the filters that are applied to requests that matchthis rule.Wherever possible, implementations SHOULD implement filters in the orderthey are specified.Implementations MAY choose to implement this ordering strictly, rejectingany combination or order of filters that can not be supported. If implementationschoose a strict interpretation of filter ordering, they MUST clearly documentthat behavior.To reject an invalid combination or order of filters, implementations SHOULDconsider the Route Rules with this configuration invalid. If all Route Rulesin a Route are invalid, the entire Route would be considered invalid. If onlya portion of Route Rules are invalid, implementations MUST set the'PartiallyInvalid' condition for the Route.Conformance-levels at this level are defined based on the type of filter:- ALL core filters MUST be supported by all implementations.- Implementers are encouraged to support extended filters.- Implementation-specific custom filters have no API guarantees across  implementations.Specifying the same filter multiple times is not supported unless explicitlyindicated in the filter.All filters are expected to be compatible with each other except for theURLRewrite and RequestRedirect filters, which may not be combined. If animplementation can not support other combinations of filters, they must clearlydocument that limitation. In cases where incompatible or unsupportedfilters are specified and cause the 'Accepted' condition to be set to status'False', implementations may use the 'IncompatibleFilters' reason to specifythis configuration error.Support: Core",
									MarkdownDescription: "Filters define the filters that are applied to requests that matchthis rule.Wherever possible, implementations SHOULD implement filters in the orderthey are specified.Implementations MAY choose to implement this ordering strictly, rejectingany combination or order of filters that can not be supported. If implementationschoose a strict interpretation of filter ordering, they MUST clearly documentthat behavior.To reject an invalid combination or order of filters, implementations SHOULDconsider the Route Rules with this configuration invalid. If all Route Rulesin a Route are invalid, the entire Route would be considered invalid. If onlya portion of Route Rules are invalid, implementations MUST set the'PartiallyInvalid' condition for the Route.Conformance-levels at this level are defined based on the type of filter:- ALL core filters MUST be supported by all implementations.- Implementers are encouraged to support extended filters.- Implementation-specific custom filters have no API guarantees across  implementations.Specifying the same filter multiple times is not supported unless explicitlyindicated in the filter.All filters are expected to be compatible with each other except for theURLRewrite and RequestRedirect filters, which may not be combined. If animplementation can not support other combinations of filters, they must clearlydocument that limitation. In cases where incompatible or unsupportedfilters are specified and cause the 'Accepted' condition to be set to status'False', implementations may use the 'IncompatibleFilters' reason to specifythis configuration error.Support: Core",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"extension_ref": schema.SingleNestedAttribute{
												Description:         "ExtensionRef is an optional, implementation-specific extension to the'filter' behavior.  For example, resource 'myroutefilter' in group'networking.example.net'). ExtensionRef MUST NOT be used for core andextended filters.This filter can be used multiple times within the same rule.Support: Implementation-specific",
												MarkdownDescription: "ExtensionRef is an optional, implementation-specific extension to the'filter' behavior.  For example, resource 'myroutefilter' in group'networking.example.net'). ExtensionRef MUST NOT be used for core andextended filters.This filter can be used multiple times within the same rule.Support: Implementation-specific",
												Attributes: map[string]schema.Attribute{
													"group": schema.StringAttribute{
														Description:         "Group is the group of the referent. For example, 'gateway.networking.k8s.io'.When unspecified or empty string, core API group is inferred.",
														MarkdownDescription: "Group is the group of the referent. For example, 'gateway.networking.k8s.io'.When unspecified or empty string, core API group is inferred.",
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
												Description:         "RequestHeaderModifier defines a schema for a filter that modifies requestheaders.Support: Core",
												MarkdownDescription: "RequestHeaderModifier defines a schema for a filter that modifies requestheaders.Support: Core",
												Attributes: map[string]schema.Attribute{
													"add": schema.ListNestedAttribute{
														Description:         "Add adds the given header(s) (name, value) to the requestbefore the action. It appends to any existing values associatedwith the header name.Input:  GET /foo HTTP/1.1  my-header: fooConfig:  add:  - name: 'my-header'    value: 'bar,baz'Output:  GET /foo HTTP/1.1  my-header: foo,bar,baz",
														MarkdownDescription: "Add adds the given header(s) (name, value) to the requestbefore the action. It appends to any existing values associatedwith the header name.Input:  GET /foo HTTP/1.1  my-header: fooConfig:  add:  - name: 'my-header'    value: 'bar,baz'Output:  GET /foo HTTP/1.1  my-header: foo,bar,baz",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is the name of the HTTP Header to be matched. Name matching MUST becase insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).If multiple entries specify equivalent header names, the first entry withan equivalent name MUST be considered for a match. Subsequent entrieswith an equivalent header name MUST be ignored. Due to thecase-insensitivity of header names, 'foo' and 'Foo' are consideredequivalent.",
																	MarkdownDescription: "Name is the name of the HTTP Header to be matched. Name matching MUST becase insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).If multiple entries specify equivalent header names, the first entry withan equivalent name MUST be considered for a match. Subsequent entrieswith an equivalent header name MUST be ignored. Due to thecase-insensitivity of header names, 'foo' and 'Foo' are consideredequivalent.",
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
														Description:         "Remove the given header(s) from the HTTP request before the action. Thevalue of Remove is a list of HTTP header names. Note that the headernames are case-insensitive (seehttps://datatracker.ietf.org/doc/html/rfc2616#section-4.2).Input:  GET /foo HTTP/1.1  my-header1: foo  my-header2: bar  my-header3: bazConfig:  remove: ['my-header1', 'my-header3']Output:  GET /foo HTTP/1.1  my-header2: bar",
														MarkdownDescription: "Remove the given header(s) from the HTTP request before the action. Thevalue of Remove is a list of HTTP header names. Note that the headernames are case-insensitive (seehttps://datatracker.ietf.org/doc/html/rfc2616#section-4.2).Input:  GET /foo HTTP/1.1  my-header1: foo  my-header2: bar  my-header3: bazConfig:  remove: ['my-header1', 'my-header3']Output:  GET /foo HTTP/1.1  my-header2: bar",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"set": schema.ListNestedAttribute{
														Description:         "Set overwrites the request with the given header (name, value)before the action.Input:  GET /foo HTTP/1.1  my-header: fooConfig:  set:  - name: 'my-header'    value: 'bar'Output:  GET /foo HTTP/1.1  my-header: bar",
														MarkdownDescription: "Set overwrites the request with the given header (name, value)before the action.Input:  GET /foo HTTP/1.1  my-header: fooConfig:  set:  - name: 'my-header'    value: 'bar'Output:  GET /foo HTTP/1.1  my-header: bar",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is the name of the HTTP Header to be matched. Name matching MUST becase insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).If multiple entries specify equivalent header names, the first entry withan equivalent name MUST be considered for a match. Subsequent entrieswith an equivalent header name MUST be ignored. Due to thecase-insensitivity of header names, 'foo' and 'Foo' are consideredequivalent.",
																	MarkdownDescription: "Name is the name of the HTTP Header to be matched. Name matching MUST becase insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).If multiple entries specify equivalent header names, the first entry withan equivalent name MUST be considered for a match. Subsequent entrieswith an equivalent header name MUST be ignored. Due to thecase-insensitivity of header names, 'foo' and 'Foo' are consideredequivalent.",
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
												Description:         "RequestMirror defines a schema for a filter that mirrors requests.Requests are sent to the specified destination, but responses fromthat destination are ignored.This filter can be used multiple times within the same rule. Note thatnot all implementations will be able to support mirroring to multiplebackends.Support: Extended",
												MarkdownDescription: "RequestMirror defines a schema for a filter that mirrors requests.Requests are sent to the specified destination, but responses fromthat destination are ignored.This filter can be used multiple times within the same rule. Note thatnot all implementations will be able to support mirroring to multiplebackends.Support: Extended",
												Attributes: map[string]schema.Attribute{
													"backend_ref": schema.SingleNestedAttribute{
														Description:         "BackendRef references a resource where mirrored requests are sent.Mirrored requests must be sent only to a single destination endpointwithin this BackendRef, irrespective of how many endpoints are presentwithin this BackendRef.If the referent cannot be found, this BackendRef is invalid and must bedropped from the Gateway. The controller must ensure the 'ResolvedRefs'condition on the Route status is set to 'status: False' and not configurethis backend in the underlying implementation.If there is a cross-namespace reference to an *existing* objectthat is not allowed by a ReferenceGrant, the controller must ensure the'ResolvedRefs'  condition on the Route is set to 'status: False',with the 'RefNotPermitted' reason and not configure this backend in theunderlying implementation.In either error case, the Message of the 'ResolvedRefs' Conditionshould be used to provide more detail about the problem.Support: Extended for Kubernetes ServiceSupport: Implementation-specific for any other resource",
														MarkdownDescription: "BackendRef references a resource where mirrored requests are sent.Mirrored requests must be sent only to a single destination endpointwithin this BackendRef, irrespective of how many endpoints are presentwithin this BackendRef.If the referent cannot be found, this BackendRef is invalid and must bedropped from the Gateway. The controller must ensure the 'ResolvedRefs'condition on the Route status is set to 'status: False' and not configurethis backend in the underlying implementation.If there is a cross-namespace reference to an *existing* objectthat is not allowed by a ReferenceGrant, the controller must ensure the'ResolvedRefs'  condition on the Route is set to 'status: False',with the 'RefNotPermitted' reason and not configure this backend in theunderlying implementation.In either error case, the Message of the 'ResolvedRefs' Conditionshould be used to provide more detail about the problem.Support: Extended for Kubernetes ServiceSupport: Implementation-specific for any other resource",
														Attributes: map[string]schema.Attribute{
															"group": schema.StringAttribute{
																Description:         "Group is the group of the referent. For example, 'gateway.networking.k8s.io'.When unspecified or empty string, core API group is inferred.",
																MarkdownDescription: "Group is the group of the referent. For example, 'gateway.networking.k8s.io'.When unspecified or empty string, core API group is inferred.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtMost(253),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^$|^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
																},
															},

															"kind": schema.StringAttribute{
																Description:         "Kind is the Kubernetes resource kind of the referent. For example'Service'.Defaults to 'Service' when not specified.ExternalName services can refer to CNAME DNS records that may liveoutside of the cluster and as such are difficult to reason about interms of conformance. They also may not be safe to forward to (seeCVE-2021-25740 for more information). Implementations SHOULD NOTsupport ExternalName Services.Support: Core (Services with a type other than ExternalName)Support: Implementation-specific (Services with type ExternalName)",
																MarkdownDescription: "Kind is the Kubernetes resource kind of the referent. For example'Service'.Defaults to 'Service' when not specified.ExternalName services can refer to CNAME DNS records that may liveoutside of the cluster and as such are difficult to reason about interms of conformance. They also may not be safe to forward to (seeCVE-2021-25740 for more information). Implementations SHOULD NOTsupport ExternalName Services.Support: Core (Services with a type other than ExternalName)Support: Implementation-specific (Services with type ExternalName)",
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
																Description:         "Namespace is the namespace of the backend. When unspecified, the localnamespace is inferred.Note that when a namespace different than the local namespace is specified,a ReferenceGrant object is required in the referent namespace to allow thatnamespace's owner to accept the reference. See the ReferenceGrantdocumentation for details.Support: Core",
																MarkdownDescription: "Namespace is the namespace of the backend. When unspecified, the localnamespace is inferred.Note that when a namespace different than the local namespace is specified,a ReferenceGrant object is required in the referent namespace to allow thatnamespace's owner to accept the reference. See the ReferenceGrantdocumentation for details.Support: Core",
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
																Description:         "Port specifies the destination port number to use for this resource.Port is required when the referent is a Kubernetes Service. In thiscase, the port number is the service port number, not the target port.For other resources, destination port might be derived from the referentresource or this field.",
																MarkdownDescription: "Port specifies the destination port number to use for this resource.Port is required when the referent is a Kubernetes Service. In thiscase, the port number is the service port number, not the target port.For other resources, destination port might be derived from the referentresource or this field.",
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
												Description:         "RequestRedirect defines a schema for a filter that responds to therequest with an HTTP redirection.Support: Core",
												MarkdownDescription: "RequestRedirect defines a schema for a filter that responds to therequest with an HTTP redirection.Support: Core",
												Attributes: map[string]schema.Attribute{
													"hostname": schema.StringAttribute{
														Description:         "Hostname is the hostname to be used in the value of the 'Location'header in the response.When empty, the hostname in the 'Host' header of the request is used.Support: Core",
														MarkdownDescription: "Hostname is the hostname to be used in the value of the 'Location'header in the response.When empty, the hostname in the 'Host' header of the request is used.Support: Core",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.LengthAtMost(253),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
														},
													},

													"path": schema.SingleNestedAttribute{
														Description:         "Path defines parameters used to modify the path of the incoming request.The modified path is then used to construct the 'Location' header. Whenempty, the request path is used as-is.Support: Extended",
														MarkdownDescription: "Path defines parameters used to modify the path of the incoming request.The modified path is then used to construct the 'Location' header. Whenempty, the request path is used as-is.Support: Extended",
														Attributes: map[string]schema.Attribute{
															"replace_full_path": schema.StringAttribute{
																Description:         "ReplaceFullPath specifies the value with which to replace the full pathof a request during a rewrite or redirect.",
																MarkdownDescription: "ReplaceFullPath specifies the value with which to replace the full pathof a request during a rewrite or redirect.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtMost(1024),
																},
															},

															"replace_prefix_match": schema.StringAttribute{
																Description:         "ReplacePrefixMatch specifies the value with which to replace the prefixmatch of a request during a rewrite or redirect. For example, a requestto '/foo/bar' with a prefix match of '/foo' and a ReplacePrefixMatchof '/xyz' would be modified to '/xyz/bar'.Note that this matches the behavior of the PathPrefix match type. Thismatches full path elements. A path element refers to the list of labelsin the path split by the '/' separator. When specified, a trailing '/' isignored. For example, the paths '/abc', '/abc/', and '/abc/def' would allmatch the prefix '/abc', but the path '/abcd' would not.ReplacePrefixMatch is only compatible with a 'PathPrefix' HTTPRouteMatch.Using any other HTTPRouteMatch type on the same HTTPRouteRule will result inthe implementation setting the Accepted Condition for the Route to 'status: False'.Request Path | Prefix Match | Replace Prefix | Modified Path-------------|--------------|----------------|----------/foo/bar     | /foo         | /xyz           | /xyz/bar/foo/bar     | /foo         | /xyz/          | /xyz/bar/foo/bar     | /foo/        | /xyz           | /xyz/bar/foo/bar     | /foo/        | /xyz/          | /xyz/bar/foo         | /foo         | /xyz           | /xyz/foo/        | /foo         | /xyz           | /xyz//foo/bar     | /foo         | <empty string> | /bar/foo/        | /foo         | <empty string> | //foo         | /foo         | <empty string> | //foo/        | /foo         | /              | //foo         | /foo         | /              | /",
																MarkdownDescription: "ReplacePrefixMatch specifies the value with which to replace the prefixmatch of a request during a rewrite or redirect. For example, a requestto '/foo/bar' with a prefix match of '/foo' and a ReplacePrefixMatchof '/xyz' would be modified to '/xyz/bar'.Note that this matches the behavior of the PathPrefix match type. Thismatches full path elements. A path element refers to the list of labelsin the path split by the '/' separator. When specified, a trailing '/' isignored. For example, the paths '/abc', '/abc/', and '/abc/def' would allmatch the prefix '/abc', but the path '/abcd' would not.ReplacePrefixMatch is only compatible with a 'PathPrefix' HTTPRouteMatch.Using any other HTTPRouteMatch type on the same HTTPRouteRule will result inthe implementation setting the Accepted Condition for the Route to 'status: False'.Request Path | Prefix Match | Replace Prefix | Modified Path-------------|--------------|----------------|----------/foo/bar     | /foo         | /xyz           | /xyz/bar/foo/bar     | /foo         | /xyz/          | /xyz/bar/foo/bar     | /foo/        | /xyz           | /xyz/bar/foo/bar     | /foo/        | /xyz/          | /xyz/bar/foo         | /foo         | /xyz           | /xyz/foo/        | /foo         | /xyz           | /xyz//foo/bar     | /foo         | <empty string> | /bar/foo/        | /foo         | <empty string> | //foo         | /foo         | <empty string> | //foo/        | /foo         | /              | //foo         | /foo         | /              | /",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtMost(1024),
																},
															},

															"type": schema.StringAttribute{
																Description:         "Type defines the type of path modifier. Additional types may beadded in a future release of the API.Note that values may be added to this enum, implementationsmust ensure that unknown values will not cause a crash.Unknown values here must result in the implementation setting theAccepted Condition for the Route to 'status: False', with aReason of 'UnsupportedValue'.",
																MarkdownDescription: "Type defines the type of path modifier. Additional types may beadded in a future release of the API.Note that values may be added to this enum, implementationsmust ensure that unknown values will not cause a crash.Unknown values here must result in the implementation setting theAccepted Condition for the Route to 'status: False', with aReason of 'UnsupportedValue'.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("ReplaceFullPath", "ReplacePrefixMatch"),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"port": schema.Int64Attribute{
														Description:         "Port is the port to be used in the value of the 'Location'header in the response.If no port is specified, the redirect port MUST be derived using thefollowing rules:* If redirect scheme is not-empty, the redirect port MUST be the well-known  port associated with the redirect scheme. Specifically 'http' to port 80  and 'https' to port 443. If the redirect scheme does not have a  well-known port, the listener port of the Gateway SHOULD be used.* If redirect scheme is empty, the redirect port MUST be the Gateway  Listener port.Implementations SHOULD NOT add the port number in the 'Location'header in the following cases:* A Location header that will use HTTP (whether that is determined via  the Listener protocol or the Scheme field) _and_ use port 80.* A Location header that will use HTTPS (whether that is determined via  the Listener protocol or the Scheme field) _and_ use port 443.Support: Extended",
														MarkdownDescription: "Port is the port to be used in the value of the 'Location'header in the response.If no port is specified, the redirect port MUST be derived using thefollowing rules:* If redirect scheme is not-empty, the redirect port MUST be the well-known  port associated with the redirect scheme. Specifically 'http' to port 80  and 'https' to port 443. If the redirect scheme does not have a  well-known port, the listener port of the Gateway SHOULD be used.* If redirect scheme is empty, the redirect port MUST be the Gateway  Listener port.Implementations SHOULD NOT add the port number in the 'Location'header in the following cases:* A Location header that will use HTTP (whether that is determined via  the Listener protocol or the Scheme field) _and_ use port 80.* A Location header that will use HTTPS (whether that is determined via  the Listener protocol or the Scheme field) _and_ use port 443.Support: Extended",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(1),
															int64validator.AtMost(65535),
														},
													},

													"scheme": schema.StringAttribute{
														Description:         "Scheme is the scheme to be used in the value of the 'Location' header inthe response. When empty, the scheme of the request is used.Scheme redirects can affect the port of the redirect, for more information,refer to the documentation for the port field of this filter.Note that values may be added to this enum, implementationsmust ensure that unknown values will not cause a crash.Unknown values here must result in the implementation setting theAccepted Condition for the Route to 'status: False', with aReason of 'UnsupportedValue'.Support: Extended",
														MarkdownDescription: "Scheme is the scheme to be used in the value of the 'Location' header inthe response. When empty, the scheme of the request is used.Scheme redirects can affect the port of the redirect, for more information,refer to the documentation for the port field of this filter.Note that values may be added to this enum, implementationsmust ensure that unknown values will not cause a crash.Unknown values here must result in the implementation setting theAccepted Condition for the Route to 'status: False', with aReason of 'UnsupportedValue'.Support: Extended",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("http", "https"),
														},
													},

													"status_code": schema.Int64Attribute{
														Description:         "StatusCode is the HTTP status code to be used in response.Note that values may be added to this enum, implementationsmust ensure that unknown values will not cause a crash.Unknown values here must result in the implementation setting theAccepted Condition for the Route to 'status: False', with aReason of 'UnsupportedValue'.Support: Core",
														MarkdownDescription: "StatusCode is the HTTP status code to be used in response.Note that values may be added to this enum, implementationsmust ensure that unknown values will not cause a crash.Unknown values here must result in the implementation setting theAccepted Condition for the Route to 'status: False', with aReason of 'UnsupportedValue'.Support: Core",
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

											"response_header_modifier": schema.SingleNestedAttribute{
												Description:         "ResponseHeaderModifier defines a schema for a filter that modifies responseheaders.Support: Extended",
												MarkdownDescription: "ResponseHeaderModifier defines a schema for a filter that modifies responseheaders.Support: Extended",
												Attributes: map[string]schema.Attribute{
													"add": schema.ListNestedAttribute{
														Description:         "Add adds the given header(s) (name, value) to the requestbefore the action. It appends to any existing values associatedwith the header name.Input:  GET /foo HTTP/1.1  my-header: fooConfig:  add:  - name: 'my-header'    value: 'bar,baz'Output:  GET /foo HTTP/1.1  my-header: foo,bar,baz",
														MarkdownDescription: "Add adds the given header(s) (name, value) to the requestbefore the action. It appends to any existing values associatedwith the header name.Input:  GET /foo HTTP/1.1  my-header: fooConfig:  add:  - name: 'my-header'    value: 'bar,baz'Output:  GET /foo HTTP/1.1  my-header: foo,bar,baz",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is the name of the HTTP Header to be matched. Name matching MUST becase insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).If multiple entries specify equivalent header names, the first entry withan equivalent name MUST be considered for a match. Subsequent entrieswith an equivalent header name MUST be ignored. Due to thecase-insensitivity of header names, 'foo' and 'Foo' are consideredequivalent.",
																	MarkdownDescription: "Name is the name of the HTTP Header to be matched. Name matching MUST becase insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).If multiple entries specify equivalent header names, the first entry withan equivalent name MUST be considered for a match. Subsequent entrieswith an equivalent header name MUST be ignored. Due to thecase-insensitivity of header names, 'foo' and 'Foo' are consideredequivalent.",
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
														Description:         "Remove the given header(s) from the HTTP request before the action. Thevalue of Remove is a list of HTTP header names. Note that the headernames are case-insensitive (seehttps://datatracker.ietf.org/doc/html/rfc2616#section-4.2).Input:  GET /foo HTTP/1.1  my-header1: foo  my-header2: bar  my-header3: bazConfig:  remove: ['my-header1', 'my-header3']Output:  GET /foo HTTP/1.1  my-header2: bar",
														MarkdownDescription: "Remove the given header(s) from the HTTP request before the action. Thevalue of Remove is a list of HTTP header names. Note that the headernames are case-insensitive (seehttps://datatracker.ietf.org/doc/html/rfc2616#section-4.2).Input:  GET /foo HTTP/1.1  my-header1: foo  my-header2: bar  my-header3: bazConfig:  remove: ['my-header1', 'my-header3']Output:  GET /foo HTTP/1.1  my-header2: bar",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"set": schema.ListNestedAttribute{
														Description:         "Set overwrites the request with the given header (name, value)before the action.Input:  GET /foo HTTP/1.1  my-header: fooConfig:  set:  - name: 'my-header'    value: 'bar'Output:  GET /foo HTTP/1.1  my-header: bar",
														MarkdownDescription: "Set overwrites the request with the given header (name, value)before the action.Input:  GET /foo HTTP/1.1  my-header: fooConfig:  set:  - name: 'my-header'    value: 'bar'Output:  GET /foo HTTP/1.1  my-header: bar",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is the name of the HTTP Header to be matched. Name matching MUST becase insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).If multiple entries specify equivalent header names, the first entry withan equivalent name MUST be considered for a match. Subsequent entrieswith an equivalent header name MUST be ignored. Due to thecase-insensitivity of header names, 'foo' and 'Foo' are consideredequivalent.",
																	MarkdownDescription: "Name is the name of the HTTP Header to be matched. Name matching MUST becase insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).If multiple entries specify equivalent header names, the first entry withan equivalent name MUST be considered for a match. Subsequent entrieswith an equivalent header name MUST be ignored. Due to thecase-insensitivity of header names, 'foo' and 'Foo' are consideredequivalent.",
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

											"type": schema.StringAttribute{
												Description:         "Type identifies the type of filter to apply. As with other API fields,types are classified into three conformance levels:- Core: Filter types and their corresponding configuration defined by  'Support: Core' in this package, e.g. 'RequestHeaderModifier'. All  implementations must support core filters.- Extended: Filter types and their corresponding configuration defined by  'Support: Extended' in this package, e.g. 'RequestMirror'. Implementers  are encouraged to support extended filters.- Implementation-specific: Filters that are defined and supported by  specific vendors.  In the future, filters showing convergence in behavior across multiple  implementations will be considered for inclusion in extended or core  conformance levels. Filter-specific configuration for such filters  is specified using the ExtensionRef field. 'Type' should be set to  'ExtensionRef' for custom filters.Implementers are encouraged to define custom implementation types toextend the core API with implementation-specific behavior.If a reference to a custom filter type cannot be resolved, the filterMUST NOT be skipped. Instead, requests that would have been processed bythat filter MUST receive a HTTP error response.Note that values may be added to this enum, implementationsmust ensure that unknown values will not cause a crash.Unknown values here must result in the implementation setting theAccepted Condition for the Route to 'status: False', with aReason of 'UnsupportedValue'.",
												MarkdownDescription: "Type identifies the type of filter to apply. As with other API fields,types are classified into three conformance levels:- Core: Filter types and their corresponding configuration defined by  'Support: Core' in this package, e.g. 'RequestHeaderModifier'. All  implementations must support core filters.- Extended: Filter types and their corresponding configuration defined by  'Support: Extended' in this package, e.g. 'RequestMirror'. Implementers  are encouraged to support extended filters.- Implementation-specific: Filters that are defined and supported by  specific vendors.  In the future, filters showing convergence in behavior across multiple  implementations will be considered for inclusion in extended or core  conformance levels. Filter-specific configuration for such filters  is specified using the ExtensionRef field. 'Type' should be set to  'ExtensionRef' for custom filters.Implementers are encouraged to define custom implementation types toextend the core API with implementation-specific behavior.If a reference to a custom filter type cannot be resolved, the filterMUST NOT be skipped. Instead, requests that would have been processed bythat filter MUST receive a HTTP error response.Note that values may be added to this enum, implementationsmust ensure that unknown values will not cause a crash.Unknown values here must result in the implementation setting theAccepted Condition for the Route to 'status: False', with aReason of 'UnsupportedValue'.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("RequestHeaderModifier", "ResponseHeaderModifier", "RequestMirror", "RequestRedirect", "URLRewrite", "ExtensionRef"),
												},
											},

											"url_rewrite": schema.SingleNestedAttribute{
												Description:         "URLRewrite defines a schema for a filter that modifies a request during forwarding.Support: Extended",
												MarkdownDescription: "URLRewrite defines a schema for a filter that modifies a request during forwarding.Support: Extended",
												Attributes: map[string]schema.Attribute{
													"hostname": schema.StringAttribute{
														Description:         "Hostname is the value to be used to replace the Host header value duringforwarding.Support: Extended",
														MarkdownDescription: "Hostname is the value to be used to replace the Host header value duringforwarding.Support: Extended",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.LengthAtMost(253),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
														},
													},

													"path": schema.SingleNestedAttribute{
														Description:         "Path defines a path rewrite.Support: Extended",
														MarkdownDescription: "Path defines a path rewrite.Support: Extended",
														Attributes: map[string]schema.Attribute{
															"replace_full_path": schema.StringAttribute{
																Description:         "ReplaceFullPath specifies the value with which to replace the full pathof a request during a rewrite or redirect.",
																MarkdownDescription: "ReplaceFullPath specifies the value with which to replace the full pathof a request during a rewrite or redirect.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtMost(1024),
																},
															},

															"replace_prefix_match": schema.StringAttribute{
																Description:         "ReplacePrefixMatch specifies the value with which to replace the prefixmatch of a request during a rewrite or redirect. For example, a requestto '/foo/bar' with a prefix match of '/foo' and a ReplacePrefixMatchof '/xyz' would be modified to '/xyz/bar'.Note that this matches the behavior of the PathPrefix match type. Thismatches full path elements. A path element refers to the list of labelsin the path split by the '/' separator. When specified, a trailing '/' isignored. For example, the paths '/abc', '/abc/', and '/abc/def' would allmatch the prefix '/abc', but the path '/abcd' would not.ReplacePrefixMatch is only compatible with a 'PathPrefix' HTTPRouteMatch.Using any other HTTPRouteMatch type on the same HTTPRouteRule will result inthe implementation setting the Accepted Condition for the Route to 'status: False'.Request Path | Prefix Match | Replace Prefix | Modified Path-------------|--------------|----------------|----------/foo/bar     | /foo         | /xyz           | /xyz/bar/foo/bar     | /foo         | /xyz/          | /xyz/bar/foo/bar     | /foo/        | /xyz           | /xyz/bar/foo/bar     | /foo/        | /xyz/          | /xyz/bar/foo         | /foo         | /xyz           | /xyz/foo/        | /foo         | /xyz           | /xyz//foo/bar     | /foo         | <empty string> | /bar/foo/        | /foo         | <empty string> | //foo         | /foo         | <empty string> | //foo/        | /foo         | /              | //foo         | /foo         | /              | /",
																MarkdownDescription: "ReplacePrefixMatch specifies the value with which to replace the prefixmatch of a request during a rewrite or redirect. For example, a requestto '/foo/bar' with a prefix match of '/foo' and a ReplacePrefixMatchof '/xyz' would be modified to '/xyz/bar'.Note that this matches the behavior of the PathPrefix match type. Thismatches full path elements. A path element refers to the list of labelsin the path split by the '/' separator. When specified, a trailing '/' isignored. For example, the paths '/abc', '/abc/', and '/abc/def' would allmatch the prefix '/abc', but the path '/abcd' would not.ReplacePrefixMatch is only compatible with a 'PathPrefix' HTTPRouteMatch.Using any other HTTPRouteMatch type on the same HTTPRouteRule will result inthe implementation setting the Accepted Condition for the Route to 'status: False'.Request Path | Prefix Match | Replace Prefix | Modified Path-------------|--------------|----------------|----------/foo/bar     | /foo         | /xyz           | /xyz/bar/foo/bar     | /foo         | /xyz/          | /xyz/bar/foo/bar     | /foo/        | /xyz           | /xyz/bar/foo/bar     | /foo/        | /xyz/          | /xyz/bar/foo         | /foo         | /xyz           | /xyz/foo/        | /foo         | /xyz           | /xyz//foo/bar     | /foo         | <empty string> | /bar/foo/        | /foo         | <empty string> | //foo         | /foo         | <empty string> | //foo/        | /foo         | /              | //foo         | /foo         | /              | /",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtMost(1024),
																},
															},

															"type": schema.StringAttribute{
																Description:         "Type defines the type of path modifier. Additional types may beadded in a future release of the API.Note that values may be added to this enum, implementationsmust ensure that unknown values will not cause a crash.Unknown values here must result in the implementation setting theAccepted Condition for the Route to 'status: False', with aReason of 'UnsupportedValue'.",
																MarkdownDescription: "Type defines the type of path modifier. Additional types may beadded in a future release of the API.Note that values may be added to this enum, implementationsmust ensure that unknown values will not cause a crash.Unknown values here must result in the implementation setting theAccepted Condition for the Route to 'status: False', with aReason of 'UnsupportedValue'.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("ReplaceFullPath", "ReplacePrefixMatch"),
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
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"matches": schema.ListNestedAttribute{
									Description:         "Matches define conditions used for matching the rule against incomingHTTP requests. Each match is independent, i.e. this rule will be matchedif **any** one of the matches is satisfied.For example, take the following matches configuration:'''matches:- path:    value: '/foo'  headers:  - name: 'version'    value: 'v2'- path:    value: '/v2/foo''''For a request to match against this rule, a request must satisfyEITHER of the two conditions:- path prefixed with '/foo' AND contains the header 'version: v2'- path prefix of '/v2/foo'See the documentation for HTTPRouteMatch on how to specify multiplematch conditions that should be ANDed together.If no matches are specified, the default is a prefixpath match on '/', which has the effect of matching everyHTTP request.Proxy or Load Balancer routing configuration generated from HTTPRoutesMUST prioritize matches based on the following criteria, continuing onties. Across all rules specified on applicable Routes, precedence must begiven to the match having:* 'Exact' path match.* 'Prefix' path match with largest number of characters.* Method match.* Largest number of header matches.* Largest number of query param matches.Note: The precedence of RegularExpression path matches are implementation-specific.If ties still exist across multiple Routes, matching precedence MUST bedetermined in order of the following criteria, continuing on ties:* The oldest Route based on creation timestamp.* The Route appearing first in alphabetical order by  '{namespace}/{name}'.If ties still exist within an HTTPRoute, matching precedence MUST be grantedto the FIRST matching rule (in list order) with a match meeting the abovecriteria.When no rules matching a request have been successfully attached to theparent a request is coming from, a HTTP 404 status code MUST be returned.",
									MarkdownDescription: "Matches define conditions used for matching the rule against incomingHTTP requests. Each match is independent, i.e. this rule will be matchedif **any** one of the matches is satisfied.For example, take the following matches configuration:'''matches:- path:    value: '/foo'  headers:  - name: 'version'    value: 'v2'- path:    value: '/v2/foo''''For a request to match against this rule, a request must satisfyEITHER of the two conditions:- path prefixed with '/foo' AND contains the header 'version: v2'- path prefix of '/v2/foo'See the documentation for HTTPRouteMatch on how to specify multiplematch conditions that should be ANDed together.If no matches are specified, the default is a prefixpath match on '/', which has the effect of matching everyHTTP request.Proxy or Load Balancer routing configuration generated from HTTPRoutesMUST prioritize matches based on the following criteria, continuing onties. Across all rules specified on applicable Routes, precedence must begiven to the match having:* 'Exact' path match.* 'Prefix' path match with largest number of characters.* Method match.* Largest number of header matches.* Largest number of query param matches.Note: The precedence of RegularExpression path matches are implementation-specific.If ties still exist across multiple Routes, matching precedence MUST bedetermined in order of the following criteria, continuing on ties:* The oldest Route based on creation timestamp.* The Route appearing first in alphabetical order by  '{namespace}/{name}'.If ties still exist within an HTTPRoute, matching precedence MUST be grantedto the FIRST matching rule (in list order) with a match meeting the abovecriteria.When no rules matching a request have been successfully attached to theparent a request is coming from, a HTTP 404 status code MUST be returned.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"headers": schema.ListNestedAttribute{
												Description:         "Headers specifies HTTP request header matchers. Multiple match values areANDed together, meaning, a request must match all the specified headersto select the route.",
												MarkdownDescription: "Headers specifies HTTP request header matchers. Multiple match values areANDed together, meaning, a request must match all the specified headersto select the route.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name is the name of the HTTP Header to be matched. Name matching MUST becase insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).If multiple entries specify equivalent header names, only the firstentry with an equivalent name MUST be considered for a match. Subsequententries with an equivalent header name MUST be ignored. Due to thecase-insensitivity of header names, 'foo' and 'Foo' are consideredequivalent.When a header is repeated in an HTTP request, it isimplementation-specific behavior as to how this is represented.Generally, proxies should follow the guidance from the RFC:https://www.rfc-editor.org/rfc/rfc7230.html#section-3.2.2 regardingprocessing a repeated header, with special handling for 'Set-Cookie'.",
															MarkdownDescription: "Name is the name of the HTTP Header to be matched. Name matching MUST becase insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).If multiple entries specify equivalent header names, only the firstentry with an equivalent name MUST be considered for a match. Subsequententries with an equivalent header name MUST be ignored. Due to thecase-insensitivity of header names, 'foo' and 'Foo' are consideredequivalent.When a header is repeated in an HTTP request, it isimplementation-specific behavior as to how this is represented.Generally, proxies should follow the guidance from the RFC:https://www.rfc-editor.org/rfc/rfc7230.html#section-3.2.2 regardingprocessing a repeated header, with special handling for 'Set-Cookie'.",
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
															Description:         "Type specifies how to match against the value of the header.Support: Core (Exact)Support: Implementation-specific (RegularExpression)Since RegularExpression HeaderMatchType has implementation-specificconformance, implementations can support POSIX, PCRE or any other dialectsof regular expressions. Please read the implementation's documentation todetermine the supported dialect.",
															MarkdownDescription: "Type specifies how to match against the value of the header.Support: Core (Exact)Support: Implementation-specific (RegularExpression)Since RegularExpression HeaderMatchType has implementation-specificconformance, implementations can support POSIX, PCRE or any other dialectsof regular expressions. Please read the implementation's documentation todetermine the supported dialect.",
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
												Description:         "Method specifies HTTP method matcher.When specified, this route will be matched only if the request has thespecified method.Support: Extended",
												MarkdownDescription: "Method specifies HTTP method matcher.When specified, this route will be matched only if the request has thespecified method.Support: Extended",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH"),
												},
											},

											"path": schema.SingleNestedAttribute{
												Description:         "Path specifies a HTTP request path matcher. If this field is notspecified, a default prefix match on the '/' path is provided.",
												MarkdownDescription: "Path specifies a HTTP request path matcher. If this field is notspecified, a default prefix match on the '/' path is provided.",
												Attributes: map[string]schema.Attribute{
													"type": schema.StringAttribute{
														Description:         "Type specifies how to match against the path Value.Support: Core (Exact, PathPrefix)Support: Implementation-specific (RegularExpression)",
														MarkdownDescription: "Type specifies how to match against the path Value.Support: Core (Exact, PathPrefix)Support: Implementation-specific (RegularExpression)",
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
												Description:         "QueryParams specifies HTTP query parameter matchers. Multiple matchvalues are ANDed together, meaning, a request must match all thespecified query parameters to select the route.Support: Extended",
												MarkdownDescription: "QueryParams specifies HTTP query parameter matchers. Multiple matchvalues are ANDed together, meaning, a request must match all thespecified query parameters to select the route.Support: Extended",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name is the name of the HTTP query param to be matched. This must be anexact string match. (Seehttps://tools.ietf.org/html/rfc7230#section-2.7.3).If multiple entries specify equivalent query param names, only the firstentry with an equivalent name MUST be considered for a match. Subsequententries with an equivalent query param name MUST be ignored.If a query param is repeated in an HTTP request, the behavior ispurposely left undefined, since different data planes have differentcapabilities. However, it is *recommended* that implementations shouldmatch against the first value of the param if the data plane supports it,as this behavior is expected in other load balancing contexts outside ofthe Gateway API.Users SHOULD NOT route traffic based on repeated query params to guardthemselves against potential differences in the implementations.",
															MarkdownDescription: "Name is the name of the HTTP query param to be matched. This must be anexact string match. (Seehttps://tools.ietf.org/html/rfc7230#section-2.7.3).If multiple entries specify equivalent query param names, only the firstentry with an equivalent name MUST be considered for a match. Subsequententries with an equivalent query param name MUST be ignored.If a query param is repeated in an HTTP request, the behavior ispurposely left undefined, since different data planes have differentcapabilities. However, it is *recommended* that implementations shouldmatch against the first value of the param if the data plane supports it,as this behavior is expected in other load balancing contexts outside ofthe Gateway API.Users SHOULD NOT route traffic based on repeated query params to guardthemselves against potential differences in the implementations.",
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
															Description:         "Type specifies how to match against the value of the query parameter.Support: Extended (Exact)Support: Implementation-specific (RegularExpression)Since RegularExpression QueryParamMatchType has Implementation-specificconformance, implementations can support POSIX, PCRE or any otherdialects of regular expressions. Please read the implementation'sdocumentation to determine the supported dialect.",
															MarkdownDescription: "Type specifies how to match against the value of the query parameter.Support: Extended (Exact)Support: Implementation-specific (RegularExpression)Since RegularExpression QueryParamMatchType has Implementation-specificconformance, implementations can support POSIX, PCRE or any otherdialects of regular expressions. Please read the implementation'sdocumentation to determine the supported dialect.",
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

								"session_persistence": schema.SingleNestedAttribute{
									Description:         "SessionPersistence defines and configures session persistencefor the route rule.Support: Extended",
									MarkdownDescription: "SessionPersistence defines and configures session persistencefor the route rule.Support: Extended",
									Attributes: map[string]schema.Attribute{
										"absolute_timeout": schema.StringAttribute{
											Description:         "AbsoluteTimeout defines the absolute timeout of the persistentsession. Once the AbsoluteTimeout duration has elapsed, thesession becomes invalid.Support: Extended",
											MarkdownDescription: "AbsoluteTimeout defines the absolute timeout of the persistentsession. Once the AbsoluteTimeout duration has elapsed, thesession becomes invalid.Support: Extended",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]{1,5}(h|m|s|ms)){1,4}$`), ""),
											},
										},

										"cookie_config": schema.SingleNestedAttribute{
											Description:         "CookieConfig provides configuration settings that are specificto cookie-based session persistence.Support: Core",
											MarkdownDescription: "CookieConfig provides configuration settings that are specificto cookie-based session persistence.Support: Core",
											Attributes: map[string]schema.Attribute{
												"lifetime_type": schema.StringAttribute{
													Description:         "LifetimeType specifies whether the cookie has a permanent orsession-based lifetime. A permanent cookie persists until itsspecified expiry time, defined by the Expires or Max-Age cookieattributes, while a session cookie is deleted when the currentsession ends.When set to 'Permanent', AbsoluteTimeout indicates thecookie's lifetime via the Expires or Max-Age cookie attributesand is required.When set to 'Session', AbsoluteTimeout indicates theabsolute lifetime of the cookie tracked by the gateway andis optional.Support: Core for 'Session' typeSupport: Extended for 'Permanent' type",
													MarkdownDescription: "LifetimeType specifies whether the cookie has a permanent orsession-based lifetime. A permanent cookie persists until itsspecified expiry time, defined by the Expires or Max-Age cookieattributes, while a session cookie is deleted when the currentsession ends.When set to 'Permanent', AbsoluteTimeout indicates thecookie's lifetime via the Expires or Max-Age cookie attributesand is required.When set to 'Session', AbsoluteTimeout indicates theabsolute lifetime of the cookie tracked by the gateway andis optional.Support: Core for 'Session' typeSupport: Extended for 'Permanent' type",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("Permanent", "Session"),
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"idle_timeout": schema.StringAttribute{
											Description:         "IdleTimeout defines the idle timeout of the persistent session.Once the session has been idle for more than the specifiedIdleTimeout duration, the session becomes invalid.Support: Extended",
											MarkdownDescription: "IdleTimeout defines the idle timeout of the persistent session.Once the session has been idle for more than the specifiedIdleTimeout duration, the session becomes invalid.Support: Extended",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]{1,5}(h|m|s|ms)){1,4}$`), ""),
											},
										},

										"session_name": schema.StringAttribute{
											Description:         "SessionName defines the name of the persistent session tokenwhich may be reflected in the cookie or the header. Usersshould avoid reusing session names to prevent unintendedconsequences, such as rejection or unpredictable behavior.Support: Implementation-specific",
											MarkdownDescription: "SessionName defines the name of the persistent session tokenwhich may be reflected in the cookie or the header. Usersshould avoid reusing session names to prevent unintendedconsequences, such as rejection or unpredictable behavior.Support: Implementation-specific",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtMost(128),
											},
										},

										"type": schema.StringAttribute{
											Description:         "Type defines the type of session persistence such as throughthe use a header or cookie. Defaults to cookie based sessionpersistence.Support: Core for 'Cookie' typeSupport: Extended for 'Header' type",
											MarkdownDescription: "Type defines the type of session persistence such as throughthe use a header or cookie. Defaults to cookie based sessionpersistence.Support: Core for 'Cookie' typeSupport: Extended for 'Header' type",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Cookie", "Header"),
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"timeouts": schema.SingleNestedAttribute{
									Description:         "Timeouts defines the timeouts that can be configured for an HTTP request.Support: Extended",
									MarkdownDescription: "Timeouts defines the timeouts that can be configured for an HTTP request.Support: Extended",
									Attributes: map[string]schema.Attribute{
										"backend_request": schema.StringAttribute{
											Description:         "BackendRequest specifies a timeout for an individual request from the gatewayto a backend. This covers the time from when the request first starts beingsent from the gateway to when the full response has been received from the backend.Setting a timeout to the zero duration (e.g. '0s') SHOULD disable the timeoutcompletely. Implementations that cannot completely disable the timeout MUSTinstead interpret the zero duration as the longest possible value to whichthe timeout can be set.An entire client HTTP transaction with a gateway, covered by the Request timeout,may result in more than one call from the gateway to the destination backend,for example, if automatic retries are supported.The value of BackendRequest must be a Gateway API Duration string as defined byGEP-2257.  When this field is unspecified, its behavior is implementation-specific;when specified, the value of BackendRequest must be no more than the value of theRequest timeout (since the Request timeout encompasses the BackendRequest timeout).Support: Extended",
											MarkdownDescription: "BackendRequest specifies a timeout for an individual request from the gatewayto a backend. This covers the time from when the request first starts beingsent from the gateway to when the full response has been received from the backend.Setting a timeout to the zero duration (e.g. '0s') SHOULD disable the timeoutcompletely. Implementations that cannot completely disable the timeout MUSTinstead interpret the zero duration as the longest possible value to whichthe timeout can be set.An entire client HTTP transaction with a gateway, covered by the Request timeout,may result in more than one call from the gateway to the destination backend,for example, if automatic retries are supported.The value of BackendRequest must be a Gateway API Duration string as defined byGEP-2257.  When this field is unspecified, its behavior is implementation-specific;when specified, the value of BackendRequest must be no more than the value of theRequest timeout (since the Request timeout encompasses the BackendRequest timeout).Support: Extended",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]{1,5}(h|m|s|ms)){1,4}$`), ""),
											},
										},

										"request": schema.StringAttribute{
											Description:         "Request specifies the maximum duration for a gateway to respond to an HTTP request.If the gateway has not been able to respond before this deadline is met, the gatewayMUST return a timeout error.For example, setting the 'rules.timeouts.request' field to the value '10s' in an'HTTPRoute' will cause a timeout if a client request is taking longer than 10 secondsto complete.Setting a timeout to the zero duration (e.g. '0s') SHOULD disable the timeoutcompletely. Implementations that cannot completely disable the timeout MUSTinstead interpret the zero duration as the longest possible value to whichthe timeout can be set.This timeout is intended to cover as close to the whole request-response transactionas possible although an implementation MAY choose to start the timeout after the entirerequest stream has been received instead of immediately after the transaction isinitiated by the client.The value of Request is a Gateway API Duration string as defined by GEP-2257. When thisfield is unspecified, request timeout behavior is implementation-specific.Support: Extended",
											MarkdownDescription: "Request specifies the maximum duration for a gateway to respond to an HTTP request.If the gateway has not been able to respond before this deadline is met, the gatewayMUST return a timeout error.For example, setting the 'rules.timeouts.request' field to the value '10s' in an'HTTPRoute' will cause a timeout if a client request is taking longer than 10 secondsto complete.Setting a timeout to the zero duration (e.g. '0s') SHOULD disable the timeoutcompletely. Implementations that cannot completely disable the timeout MUSTinstead interpret the zero duration as the longest possible value to whichthe timeout can be set.This timeout is intended to cover as close to the whole request-response transactionas possible although an implementation MAY choose to start the timeout after the entirerequest stream has been received instead of immediately after the transaction isinitiated by the client.The value of Request is a Gateway API Duration string as defined by GEP-2257. When thisfield is unspecified, request timeout behavior is implementation-specific.Support: Extended",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]{1,5}(h|m|s|ms)){1,4}$`), ""),
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

func (r *GatewayNetworkingK8SIoHttprouteV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_gateway_networking_k8s_io_http_route_v1beta1_manifest")

	var model GatewayNetworkingK8SIoHttprouteV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("gateway.networking.k8s.io/v1beta1")
	model.Kind = pointer.String("HTTPRoute")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
