/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

	"regexp"

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

type PolicyLinkerdIoHTTPRouteV1Beta1Resource struct{}

var (
	_ resource.Resource = (*PolicyLinkerdIoHTTPRouteV1Beta1Resource)(nil)
)

type PolicyLinkerdIoHTTPRouteV1Beta1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type PolicyLinkerdIoHTTPRouteV1Beta1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		Hostnames *[]string `tfsdk:"hostnames" yaml:"hostnames,omitempty"`

		ParentRefs *[]struct {
			Group *string `tfsdk:"group" yaml:"group,omitempty"`

			Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

			SectionName *string `tfsdk:"section_name" yaml:"sectionName,omitempty"`
		} `tfsdk:"parent_refs" yaml:"parentRefs,omitempty"`

		Rules *[]struct {
			Filters *[]struct {
				RequestHeaderModifier *struct {
					Add *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"add" yaml:"add,omitempty"`

					Remove *[]string `tfsdk:"remove" yaml:"remove,omitempty"`

					Set *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"set" yaml:"set,omitempty"`
				} `tfsdk:"request_header_modifier" yaml:"requestHeaderModifier,omitempty"`

				RequestRedirect *struct {
					Hostname *string `tfsdk:"hostname" yaml:"hostname,omitempty"`

					Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

					Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`

					StatusCode *int64 `tfsdk:"status_code" yaml:"statusCode,omitempty"`
				} `tfsdk:"request_redirect" yaml:"requestRedirect,omitempty"`

				Type *string `tfsdk:"type" yaml:"type,omitempty"`
			} `tfsdk:"filters" yaml:"filters,omitempty"`

			Matches *[]struct {
				Headers *[]struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"headers" yaml:"headers,omitempty"`

				Method *string `tfsdk:"method" yaml:"method,omitempty"`

				Path *struct {
					Type *string `tfsdk:"type" yaml:"type,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"path" yaml:"path,omitempty"`

				QueryParams *[]struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"query_params" yaml:"queryParams,omitempty"`
			} `tfsdk:"matches" yaml:"matches,omitempty"`
		} `tfsdk:"rules" yaml:"rules,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewPolicyLinkerdIoHTTPRouteV1Beta1Resource() resource.Resource {
	return &PolicyLinkerdIoHTTPRouteV1Beta1Resource{}
}

func (r *PolicyLinkerdIoHTTPRouteV1Beta1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policy_linkerd_io_http_route_v1beta1"
}

func (r *PolicyLinkerdIoHTTPRouteV1Beta1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "HTTPRoute provides a way to route HTTP requests. This includes the capability to match requests by hostname, path, header, or query param. Filters can be used to specify additional processing steps. Backends specify where matching requests should be routed.",
		MarkdownDescription: "HTTPRoute provides a way to route HTTP requests. This includes the capability to match requests by hostname, path, header, or query param. Filters can be used to specify additional processing steps. Backends specify where matching requests should be routed.",
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
				Description:         "Spec defines the desired state of HTTPRoute.",
				MarkdownDescription: "Spec defines the desired state of HTTPRoute.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"hostnames": {
						Description:         "Hostnames defines a set of hostname that should match against the HTTP Host header to select a HTTPRoute to process the request. This matches the RFC 1123 definition of a hostname with 2 notable exceptions:  1. IPs are not allowed. 2. A hostname may be prefixed with a wildcard label ('*.'). The wildcard    label must appear by itself as the first label.  If a hostname is specified by both the Listener and HTTPRoute, there must be at least one intersecting hostname for the HTTPRoute to be attached to the Listener. For example:  * A Listener with 'test.example.com' as the hostname matches HTTPRoutes   that have either not specified any hostnames, or have specified at   least one of 'test.example.com' or '*.example.com'. * A Listener with '*.example.com' as the hostname matches HTTPRoutes   that have either not specified any hostnames or have specified at least   one hostname that matches the Listener hostname. For example,   '*.example.com', 'test.example.com', and 'foo.test.example.com' would   all match. On the other hand, 'example.com' and 'test.example.net' would   not match.  Hostnames that are prefixed with a wildcard label ('*.') are interpreted as a suffix match. That means that a match for '*.example.com' would match both 'test.example.com', and 'foo.test.example.com', but not 'example.com'.  If both the Listener and HTTPRoute have specified hostnames, any HTTPRoute hostnames that do not match the Listener hostname MUST be ignored. For example, if a Listener specified '*.example.com', and the HTTPRoute specified 'test.example.com' and 'test.example.net', 'test.example.net' must not be considered for a match.  If both the Listener and HTTPRoute have specified hostnames, and none match with the criteria above, then the HTTPRoute is not accepted. The implementation must raise an 'Accepted' Condition with a status of 'False' in the corresponding RouteParentStatus.  Support: Core",
						MarkdownDescription: "Hostnames defines a set of hostname that should match against the HTTP Host header to select a HTTPRoute to process the request. This matches the RFC 1123 definition of a hostname with 2 notable exceptions:  1. IPs are not allowed. 2. A hostname may be prefixed with a wildcard label ('*.'). The wildcard    label must appear by itself as the first label.  If a hostname is specified by both the Listener and HTTPRoute, there must be at least one intersecting hostname for the HTTPRoute to be attached to the Listener. For example:  * A Listener with 'test.example.com' as the hostname matches HTTPRoutes   that have either not specified any hostnames, or have specified at   least one of 'test.example.com' or '*.example.com'. * A Listener with '*.example.com' as the hostname matches HTTPRoutes   that have either not specified any hostnames or have specified at least   one hostname that matches the Listener hostname. For example,   '*.example.com', 'test.example.com', and 'foo.test.example.com' would   all match. On the other hand, 'example.com' and 'test.example.net' would   not match.  Hostnames that are prefixed with a wildcard label ('*.') are interpreted as a suffix match. That means that a match for '*.example.com' would match both 'test.example.com', and 'foo.test.example.com', but not 'example.com'.  If both the Listener and HTTPRoute have specified hostnames, any HTTPRoute hostnames that do not match the Listener hostname MUST be ignored. For example, if a Listener specified '*.example.com', and the HTTPRoute specified 'test.example.com' and 'test.example.net', 'test.example.net' must not be considered for a match.  If both the Listener and HTTPRoute have specified hostnames, and none match with the criteria above, then the HTTPRoute is not accepted. The implementation must raise an 'Accepted' Condition with a status of 'False' in the corresponding RouteParentStatus.  Support: Core",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"parent_refs": {
						Description:         "ParentRefs references the resources (usually Gateways) that a Route wants to be attached to. Note that the referenced parent resource needs to allow this for the attachment to be complete. For Gateways, that means the Gateway needs to allow attachment from Routes of this kind and namespace.  The only kind of parent resource with 'Core' support is Gateway. This API may be extended in the future to support additional kinds of parent resources such as one of the route kinds.  It is invalid to reference an identical parent more than once. It is valid to reference multiple distinct sections within the same parent resource, such as 2 Listeners within a Gateway.  It is possible to separately reference multiple distinct objects that may be collapsed by an implementation. For example, some implementations may choose to merge compatible Gateway Listeners together. If that is the case, the list of routes attached to those resources should also be merged.",
						MarkdownDescription: "ParentRefs references the resources (usually Gateways) that a Route wants to be attached to. Note that the referenced parent resource needs to allow this for the attachment to be complete. For Gateways, that means the Gateway needs to allow attachment from Routes of this kind and namespace.  The only kind of parent resource with 'Core' support is Gateway. This API may be extended in the future to support additional kinds of parent resources such as one of the route kinds.  It is invalid to reference an identical parent more than once. It is valid to reference multiple distinct sections within the same parent resource, such as 2 Listeners within a Gateway.  It is possible to separately reference multiple distinct objects that may be collapsed by an implementation. For example, some implementations may choose to merge compatible Gateway Listeners together. If that is the case, the list of routes attached to those resources should also be merged.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"group": {
								Description:         "Group is the group of the referent.  Support: Core",
								MarkdownDescription: "Group is the group of the referent.  Support: Core",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.LengthAtMost(253),

									stringvalidator.RegexMatches(regexp.MustCompile(`^$|^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
								},
							},

							"kind": {
								Description:         "Kind is kind of the referent.  Support: Core (Gateway) Support: Custom (Other Resources)",
								MarkdownDescription: "Kind is kind of the referent.  Support: Core (Gateway) Support: Custom (Other Resources)",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.LengthAtLeast(1),

									stringvalidator.LengthAtMost(63),

									stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z]([-a-zA-Z0-9]*[a-zA-Z0-9])?$`), ""),
								},
							},

							"name": {
								Description:         "Name is the name of the referent.  Support: Core",
								MarkdownDescription: "Name is the name of the referent.  Support: Core",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.LengthAtLeast(1),

									stringvalidator.LengthAtMost(253),
								},
							},

							"namespace": {
								Description:         "Namespace is the namespace of the referent. When unspecified (or empty string), this refers to the local namespace of the Route.  Support: Core",
								MarkdownDescription: "Namespace is the namespace of the referent. When unspecified (or empty string), this refers to the local namespace of the Route.  Support: Core",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.LengthAtLeast(1),

									stringvalidator.LengthAtMost(63),

									stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
								},
							},

							"section_name": {
								Description:         "SectionName is the name of a section within the target resource. In the following resources, SectionName is interpreted as the following:  * Gateway: Listener Name. When both Port (experimental) and SectionName are specified, the name and port of the selected listener must match both specified values.  Implementations MAY choose to support attaching Routes to other resources. If that is the case, they MUST clearly document how SectionName is interpreted.  When unspecified (empty string), this will reference the entire resource. For the purpose of status, an attachment is considered successful if at least one section in the parent resource accepts it. For example, Gateway listeners can restrict which Routes can attach to them by Route kind, namespace, or hostname. If 1 of 2 Gateway listeners accept attachment from the referencing Route, the Route MUST be considered successfully attached. If no Gateway listeners accept attachment from this Route, the Route MUST be considered detached from the Gateway.  Support: Core",
								MarkdownDescription: "SectionName is the name of a section within the target resource. In the following resources, SectionName is interpreted as the following:  * Gateway: Listener Name. When both Port (experimental) and SectionName are specified, the name and port of the selected listener must match both specified values.  Implementations MAY choose to support attaching Routes to other resources. If that is the case, they MUST clearly document how SectionName is interpreted.  When unspecified (empty string), this will reference the entire resource. For the purpose of status, an attachment is considered successful if at least one section in the parent resource accepts it. For example, Gateway listeners can restrict which Routes can attach to them by Route kind, namespace, or hostname. If 1 of 2 Gateway listeners accept attachment from the referencing Route, the Route MUST be considered successfully attached. If no Gateway listeners accept attachment from this Route, the Route MUST be considered detached from the Gateway.  Support: Core",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.LengthAtLeast(1),

									stringvalidator.LengthAtMost(253),

									stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
								},
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"rules": {
						Description:         "Rules are a list of HTTP matchers, filters and actions.",
						MarkdownDescription: "Rules are a list of HTTP matchers, filters and actions.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"filters": {
								Description:         "Filters define the filters that are applied to requests that match this rule.  The effects of ordering of multiple behaviors are currently unspecified. This can change in the future based on feedback during the alpha stage.  Conformance-levels at this level are defined based on the type of filter:  - ALL core filters MUST be supported by all implementations. - Implementers are encouraged to support extended filters. - Implementation-specific custom filters have no API guarantees across   implementations.  Specifying a core filter multiple times has unspecified or custom conformance.  All filters are expected to be compatible with each other except for the URLRewrite and RequestRedirect filters, which may not be combined. If an implementation can not support other combinations of filters, they must clearly document that limitation. In all cases where incompatible or unsupported filters are specified, implementations MUST add a warning condition to status.  Support: Core",
								MarkdownDescription: "Filters define the filters that are applied to requests that match this rule.  The effects of ordering of multiple behaviors are currently unspecified. This can change in the future based on feedback during the alpha stage.  Conformance-levels at this level are defined based on the type of filter:  - ALL core filters MUST be supported by all implementations. - Implementers are encouraged to support extended filters. - Implementation-specific custom filters have no API guarantees across   implementations.  Specifying a core filter multiple times has unspecified or custom conformance.  All filters are expected to be compatible with each other except for the URLRewrite and RequestRedirect filters, which may not be combined. If an implementation can not support other combinations of filters, they must clearly document that limitation. In all cases where incompatible or unsupported filters are specified, implementations MUST add a warning condition to status.  Support: Core",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"request_header_modifier": {
										Description:         "RequestHeaderModifier defines a schema for a filter that modifies request headers.  Support: Core",
										MarkdownDescription: "RequestHeaderModifier defines a schema for a filter that modifies request headers.  Support: Core",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"add": {
												Description:         "Add adds the given header(s) (name, value) to the request before the action. It appends to any existing values associated with the header name.  Input:   GET /foo HTTP/1.1   my-header: foo  Config:   add:   - name: 'my-header'     value: 'bar'  Output:   GET /foo HTTP/1.1   my-header: foo   my-header: bar",
												MarkdownDescription: "Add adds the given header(s) (name, value) to the request before the action. It appends to any existing values associated with the header name.  Input:   GET /foo HTTP/1.1   my-header: foo  Config:   add:   - name: 'my-header'     value: 'bar'  Output:   GET /foo HTTP/1.1   my-header: foo   my-header: bar",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "Name is the name of the HTTP Header to be matched. Name matching MUST be case insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).  If multiple entries specify equivalent header names, the first entry with an equivalent name MUST be considered for a match. Subsequent entries with an equivalent header name MUST be ignored. Due to the case-insensitivity of header names, 'foo' and 'Foo' are considered equivalent.",
														MarkdownDescription: "Name is the name of the HTTP Header to be matched. Name matching MUST be case insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).  If multiple entries specify equivalent header names, the first entry with an equivalent name MUST be considered for a match. Subsequent entries with an equivalent header name MUST be ignored. Due to the case-insensitivity of header names, 'foo' and 'Foo' are considered equivalent.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.LengthAtLeast(1),

															stringvalidator.LengthAtMost(256),

															stringvalidator.RegexMatches(regexp.MustCompile(`^[A-Za-z0-9!#$%&'*+\-.^_\x60|~]+$`), ""),
														},
													},

													"value": {
														Description:         "Value is the value of HTTP Header to be matched.",
														MarkdownDescription: "Value is the value of HTTP Header to be matched.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.LengthAtLeast(1),

															stringvalidator.LengthAtMost(4096),
														},
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"remove": {
												Description:         "Remove the given header(s) from the HTTP request before the action. The value of Remove is a list of HTTP header names. Note that the header names are case-insensitive (see https://datatracker.ietf.org/doc/html/rfc2616#section-4.2).  Input:   GET /foo HTTP/1.1   my-header1: foo   my-header2: bar   my-header3: baz  Config:   remove: ['my-header1', 'my-header3']  Output:   GET /foo HTTP/1.1   my-header2: bar",
												MarkdownDescription: "Remove the given header(s) from the HTTP request before the action. The value of Remove is a list of HTTP header names. Note that the header names are case-insensitive (see https://datatracker.ietf.org/doc/html/rfc2616#section-4.2).  Input:   GET /foo HTTP/1.1   my-header1: foo   my-header2: bar   my-header3: baz  Config:   remove: ['my-header1', 'my-header3']  Output:   GET /foo HTTP/1.1   my-header2: bar",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"set": {
												Description:         "Set overwrites the request with the given header (name, value) before the action.  Input:   GET /foo HTTP/1.1   my-header: foo  Config:   set:   - name: 'my-header'     value: 'bar'  Output:   GET /foo HTTP/1.1   my-header: bar",
												MarkdownDescription: "Set overwrites the request with the given header (name, value) before the action.  Input:   GET /foo HTTP/1.1   my-header: foo  Config:   set:   - name: 'my-header'     value: 'bar'  Output:   GET /foo HTTP/1.1   my-header: bar",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "Name is the name of the HTTP Header to be matched. Name matching MUST be case insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).  If multiple entries specify equivalent header names, the first entry with an equivalent name MUST be considered for a match. Subsequent entries with an equivalent header name MUST be ignored. Due to the case-insensitivity of header names, 'foo' and 'Foo' are considered equivalent.",
														MarkdownDescription: "Name is the name of the HTTP Header to be matched. Name matching MUST be case insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).  If multiple entries specify equivalent header names, the first entry with an equivalent name MUST be considered for a match. Subsequent entries with an equivalent header name MUST be ignored. Due to the case-insensitivity of header names, 'foo' and 'Foo' are considered equivalent.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.LengthAtLeast(1),

															stringvalidator.LengthAtMost(256),

															stringvalidator.RegexMatches(regexp.MustCompile(`^[A-Za-z0-9!#$%&'*+\-.^_\x60|~]+$`), ""),
														},
													},

													"value": {
														Description:         "Value is the value of HTTP Header to be matched.",
														MarkdownDescription: "Value is the value of HTTP Header to be matched.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.LengthAtLeast(1),

															stringvalidator.LengthAtMost(4096),
														},
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

									"request_redirect": {
										Description:         "RequestRedirect defines a schema for a filter that responds to the request with an HTTP redirection.  Support: Core",
										MarkdownDescription: "RequestRedirect defines a schema for a filter that responds to the request with an HTTP redirection.  Support: Core",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"hostname": {
												Description:         "Hostname is the hostname to be used in the value of the 'Location' header in the response. When empty, the hostname of the request is used.  Support: Core",
												MarkdownDescription: "Hostname is the hostname to be used in the value of the 'Location' header in the response. When empty, the hostname of the request is used.  Support: Core",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.LengthAtLeast(1),

													stringvalidator.LengthAtMost(253),

													stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
												},
											},

											"port": {
												Description:         "Port is the port to be used in the value of the 'Location' header in the response. When empty, port (if specified) of the request is used.  Support: Extended",
												MarkdownDescription: "Port is the port to be used in the value of the 'Location' header in the response. When empty, port (if specified) of the request is used.  Support: Extended",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),

													int64validator.AtMost(65535),
												},
											},

											"scheme": {
												Description:         "Scheme is the scheme to be used in the value of the 'Location' header in the response. When empty, the scheme of the request is used.  Support: Extended",
												MarkdownDescription: "Scheme is the scheme to be used in the value of the 'Location' header in the response. When empty, the scheme of the request is used.  Support: Extended",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("http", "https"),
												},
											},

											"status_code": {
												Description:         "StatusCode is the HTTP status code to be used in response.  Support: Core",
												MarkdownDescription: "StatusCode is the HTTP status code to be used in response.  Support: Core",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.OneOf(301, 302),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": {
										Description:         "Type identifies the type of filter to apply. As with other API fields, types are classified into three conformance levels:  - Core: Filter types and their corresponding configuration defined by   'Support: Core' in this package, e.g. 'RequestHeaderModifier'.",
										MarkdownDescription: "Type identifies the type of filter to apply. As with other API fields, types are classified into three conformance levels:  - Core: Filter types and their corresponding configuration defined by   'Support: Core' in this package, e.g. 'RequestHeaderModifier'.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("RequestHeaderModifier", "RequestRedirect"),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"matches": {
								Description:         "Matches define conditions used for matching the rule against incoming HTTP requests. Each match is independent, i.e. this rule will be matched if **any** one of the matches is satisfied.  For example, take the following matches configuration:  ''' matches: - path:     value: '/foo'   headers:   - name: 'version'     value: 'v2' - path:     value: '/v2/foo' '''  For a request to match against this rule, a request must satisfy EITHER of the two conditions:  - path prefixed with '/foo' AND contains the header 'version: v2' - path prefix of '/v2/foo'  See the documentation for HTTPRouteMatch on how to specify multiple match conditions that should be ANDed together.  If no matches are specified, the default is a prefix path match on '/', which has the effect of matching every HTTP request.  Proxy or Load Balancer routing configuration generated from HTTPRoutes MUST prioritize rules based on the following criteria, continuing on ties. Precedence must be given to the the Rule with the largest number of:  * Characters in a matching non-wildcard hostname. * Characters in a matching hostname. * Characters in a matching path. * Header matches. * Query param matches.  If ties still exist across multiple Routes, matching precedence MUST be determined in order of the following criteria, continuing on ties:  * The oldest Route based on creation timestamp. * The Route appearing first in alphabetical order by   '{namespace}/{name}'.  If ties still exist within the Route that has been given precedence, matching precedence MUST be granted to the first matching rule meeting the above criteria.  When no rules matching a request have been successfully attached to the parent a request is coming from, a HTTP 404 status code MUST be returned.",
								MarkdownDescription: "Matches define conditions used for matching the rule against incoming HTTP requests. Each match is independent, i.e. this rule will be matched if **any** one of the matches is satisfied.  For example, take the following matches configuration:  ''' matches: - path:     value: '/foo'   headers:   - name: 'version'     value: 'v2' - path:     value: '/v2/foo' '''  For a request to match against this rule, a request must satisfy EITHER of the two conditions:  - path prefixed with '/foo' AND contains the header 'version: v2' - path prefix of '/v2/foo'  See the documentation for HTTPRouteMatch on how to specify multiple match conditions that should be ANDed together.  If no matches are specified, the default is a prefix path match on '/', which has the effect of matching every HTTP request.  Proxy or Load Balancer routing configuration generated from HTTPRoutes MUST prioritize rules based on the following criteria, continuing on ties. Precedence must be given to the the Rule with the largest number of:  * Characters in a matching non-wildcard hostname. * Characters in a matching hostname. * Characters in a matching path. * Header matches. * Query param matches.  If ties still exist across multiple Routes, matching precedence MUST be determined in order of the following criteria, continuing on ties:  * The oldest Route based on creation timestamp. * The Route appearing first in alphabetical order by   '{namespace}/{name}'.  If ties still exist within the Route that has been given precedence, matching precedence MUST be granted to the first matching rule meeting the above criteria.  When no rules matching a request have been successfully attached to the parent a request is coming from, a HTTP 404 status code MUST be returned.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"headers": {
										Description:         "Headers specifies HTTP request header matchers. Multiple match values are ANDed together, meaning, a request must match all the specified headers to select the route.",
										MarkdownDescription: "Headers specifies HTTP request header matchers. Multiple match values are ANDed together, meaning, a request must match all the specified headers to select the route.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "Name is the name of the HTTP Header to be matched. Name matching MUST be case insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).  If multiple entries specify equivalent header names, only the first entry with an equivalent name MUST be considered for a match. Subsequent entries with an equivalent header name MUST be ignored. Due to the case-insensitivity of header names, 'foo' and 'Foo' are considered equivalent.  When a header is repeated in an HTTP request, it is implementation-specific behavior as to how this is represented. Generally, proxies should follow the guidance from the RFC: https://www.rfc-editor.org/rfc/rfc7230.html#section-3.2.2 regarding processing a repeated header, with special handling for 'Set-Cookie'.",
												MarkdownDescription: "Name is the name of the HTTP Header to be matched. Name matching MUST be case insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).  If multiple entries specify equivalent header names, only the first entry with an equivalent name MUST be considered for a match. Subsequent entries with an equivalent header name MUST be ignored. Due to the case-insensitivity of header names, 'foo' and 'Foo' are considered equivalent.  When a header is repeated in an HTTP request, it is implementation-specific behavior as to how this is represented. Generally, proxies should follow the guidance from the RFC: https://www.rfc-editor.org/rfc/rfc7230.html#section-3.2.2 regarding processing a repeated header, with special handling for 'Set-Cookie'.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.LengthAtLeast(1),

													stringvalidator.LengthAtMost(256),

													stringvalidator.RegexMatches(regexp.MustCompile(`^[A-Za-z0-9!#$%&'*+\-.^_\x60|~]+$`), ""),
												},
											},

											"type": {
												Description:         "Type specifies how to match against the value of the header.  Support: Core (Exact)  Support: Custom (RegularExpression)  Since RegularExpression HeaderMatchType has custom conformance, implementations can support POSIX, PCRE or any other dialects of regular expressions. Please read the implementation's documentation to determine the supported dialect.",
												MarkdownDescription: "Type specifies how to match against the value of the header.  Support: Core (Exact)  Support: Custom (RegularExpression)  Since RegularExpression HeaderMatchType has custom conformance, implementations can support POSIX, PCRE or any other dialects of regular expressions. Please read the implementation's documentation to determine the supported dialect.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("Exact", "RegularExpression"),
												},
											},

											"value": {
												Description:         "Value is the value of HTTP Header to be matched.",
												MarkdownDescription: "Value is the value of HTTP Header to be matched.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.LengthAtLeast(1),

													stringvalidator.LengthAtMost(4096),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"method": {
										Description:         "Method specifies HTTP method matcher. When specified, this route will be matched only if the request has the specified method.  Support: Extended",
										MarkdownDescription: "Method specifies HTTP method matcher. When specified, this route will be matched only if the request has the specified method.  Support: Extended",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH"),
										},
									},

									"path": {
										Description:         "Path specifies a HTTP request path matcher. If this field is not specified, a default prefix match on the '/' path is provided.",
										MarkdownDescription: "Path specifies a HTTP request path matcher. If this field is not specified, a default prefix match on the '/' path is provided.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"type": {
												Description:         "Type specifies how to match against the path Value.  Support: Core (Exact, PathPrefix)  Support: Custom (RegularExpression)",
												MarkdownDescription: "Type specifies how to match against the path Value.  Support: Core (Exact, PathPrefix)  Support: Custom (RegularExpression)",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("Exact", "PathPrefix", "RegularExpression"),
												},
											},

											"value": {
												Description:         "Value of the HTTP path to match against.",
												MarkdownDescription: "Value of the HTTP path to match against.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.LengthAtMost(1024),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"query_params": {
										Description:         "QueryParams specifies HTTP query parameter matchers. Multiple match values are ANDed together, meaning, a request must match all the specified query parameters to select the route.",
										MarkdownDescription: "QueryParams specifies HTTP query parameter matchers. Multiple match values are ANDed together, meaning, a request must match all the specified query parameters to select the route.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "Name is the name of the HTTP query param to be matched. This must be an exact string match. (See https://tools.ietf.org/html/rfc7230#section-2.7.3).",
												MarkdownDescription: "Name is the name of the HTTP query param to be matched. This must be an exact string match. (See https://tools.ietf.org/html/rfc7230#section-2.7.3).",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.LengthAtLeast(1),

													stringvalidator.LengthAtMost(256),
												},
											},

											"type": {
												Description:         "Type specifies how to match against the value of the query parameter.  Support: Extended (Exact)  Support: Custom (RegularExpression)  Since RegularExpression QueryParamMatchType has custom conformance, implementations can support POSIX, PCRE or any other dialects of regular expressions. Please read the implementation's documentation to determine the supported dialect.",
												MarkdownDescription: "Type specifies how to match against the value of the query parameter.  Support: Extended (Exact)  Support: Custom (RegularExpression)  Since RegularExpression QueryParamMatchType has custom conformance, implementations can support POSIX, PCRE or any other dialects of regular expressions. Please read the implementation's documentation to determine the supported dialect.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("Exact", "RegularExpression"),
												},
											},

											"value": {
												Description:         "Value is the value of HTTP query param to be matched.",
												MarkdownDescription: "Value is the value of HTTP query param to be matched.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.LengthAtLeast(1),

													stringvalidator.LengthAtMost(1024),
												},
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

func (r *PolicyLinkerdIoHTTPRouteV1Beta1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_policy_linkerd_io_http_route_v1beta1")

	var state PolicyLinkerdIoHTTPRouteV1Beta1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel PolicyLinkerdIoHTTPRouteV1Beta1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("policy.linkerd.io/v1beta1")
	goModel.Kind = utilities.Ptr("HTTPRoute")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *PolicyLinkerdIoHTTPRouteV1Beta1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_policy_linkerd_io_http_route_v1beta1")
	// NO-OP: All data is already in Terraform state
}

func (r *PolicyLinkerdIoHTTPRouteV1Beta1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_policy_linkerd_io_http_route_v1beta1")

	var state PolicyLinkerdIoHTTPRouteV1Beta1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel PolicyLinkerdIoHTTPRouteV1Beta1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("policy.linkerd.io/v1beta1")
	goModel.Kind = utilities.Ptr("HTTPRoute")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *PolicyLinkerdIoHTTPRouteV1Beta1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_policy_linkerd_io_http_route_v1beta1")
	// NO-OP: Terraform removes the state automatically for us
}
