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

type GatewayNetworkingK8SIoGatewayV1Alpha2Resource struct{}

var (
	_ resource.Resource = (*GatewayNetworkingK8SIoGatewayV1Alpha2Resource)(nil)
)

type GatewayNetworkingK8SIoGatewayV1Alpha2TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type GatewayNetworkingK8SIoGatewayV1Alpha2GoModel struct {
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
		Addresses *[]struct {
			Type *string `tfsdk:"type" yaml:"type,omitempty"`

			Value *string `tfsdk:"value" yaml:"value,omitempty"`
		} `tfsdk:"addresses" yaml:"addresses,omitempty"`

		GatewayClassName *string `tfsdk:"gateway_class_name" yaml:"gatewayClassName,omitempty"`

		Listeners *[]struct {
			AllowedRoutes *struct {
				Namespaces *struct {
					From *string `tfsdk:"from" yaml:"from,omitempty"`

					Selector *struct {
						MatchExpressions *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
						} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

						MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
					} `tfsdk:"selector" yaml:"selector,omitempty"`
				} `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

				Kinds *[]struct {
					Group *string `tfsdk:"group" yaml:"group,omitempty"`

					Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`
				} `tfsdk:"kinds" yaml:"kinds,omitempty"`
			} `tfsdk:"allowed_routes" yaml:"allowedRoutes,omitempty"`

			Hostname *string `tfsdk:"hostname" yaml:"hostname,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

			Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`

			Tls *struct {
				Options *map[string]string `tfsdk:"options" yaml:"options,omitempty"`

				CertificateRefs *[]struct {
					Group *string `tfsdk:"group" yaml:"group,omitempty"`

					Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
				} `tfsdk:"certificate_refs" yaml:"certificateRefs,omitempty"`

				Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`
			} `tfsdk:"tls" yaml:"tls,omitempty"`
		} `tfsdk:"listeners" yaml:"listeners,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewGatewayNetworkingK8SIoGatewayV1Alpha2Resource() resource.Resource {
	return &GatewayNetworkingK8SIoGatewayV1Alpha2Resource{}
}

func (r *GatewayNetworkingK8SIoGatewayV1Alpha2Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_networking_k8s_io_gateway_v1alpha2"
}

func (r *GatewayNetworkingK8SIoGatewayV1Alpha2Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Gateway represents an instance of a service-traffic handling infrastructure by binding Listeners to a set of IP addresses.",
		MarkdownDescription: "Gateway represents an instance of a service-traffic handling infrastructure by binding Listeners to a set of IP addresses.",
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
				Description:         "Spec defines the desired state of Gateway.",
				MarkdownDescription: "Spec defines the desired state of Gateway.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"addresses": {
						Description:         "Addresses requested for this Gateway. This is optional and behavior can depend on the implementation. If a value is set in the spec and the requested address is invalid or unavailable, the implementation MUST indicate this in the associated entry in GatewayStatus.Addresses.  The Addresses field represents a request for the address(es) on the 'outside of the Gateway', that traffic bound for this Gateway will use. This could be the IP address or hostname of an external load balancer or other networking infrastructure, or some other address that traffic will be sent to.  The .listener.hostname field is used to route traffic that has already arrived at the Gateway to the correct in-cluster destination.  If no Addresses are specified, the implementation MAY schedule the Gateway in an implementation-specific manner, assigning an appropriate set of Addresses.  The implementation MUST bind all Listeners to every GatewayAddress that it assigns to the Gateway and add a corresponding entry in GatewayStatus.Addresses.  Support: Core",
						MarkdownDescription: "Addresses requested for this Gateway. This is optional and behavior can depend on the implementation. If a value is set in the spec and the requested address is invalid or unavailable, the implementation MUST indicate this in the associated entry in GatewayStatus.Addresses.  The Addresses field represents a request for the address(es) on the 'outside of the Gateway', that traffic bound for this Gateway will use. This could be the IP address or hostname of an external load balancer or other networking infrastructure, or some other address that traffic will be sent to.  The .listener.hostname field is used to route traffic that has already arrived at the Gateway to the correct in-cluster destination.  If no Addresses are specified, the implementation MAY schedule the Gateway in an implementation-specific manner, assigning an appropriate set of Addresses.  The implementation MUST bind all Listeners to every GatewayAddress that it assigns to the Gateway and add a corresponding entry in GatewayStatus.Addresses.  Support: Core",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"type": {
								Description:         "Type of the address.",
								MarkdownDescription: "Type of the address.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"value": {
								Description:         "Value of the address. The validity of the values will depend on the type and support by the controller.  Examples: '1.2.3.4', '128::1', 'my-ip-address'.",
								MarkdownDescription: "Value of the address. The validity of the values will depend on the type and support by the controller.  Examples: '1.2.3.4', '128::1', 'my-ip-address'.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"gateway_class_name": {
						Description:         "GatewayClassName used for this Gateway. This is the name of a GatewayClass resource.",
						MarkdownDescription: "GatewayClassName used for this Gateway. This is the name of a GatewayClass resource.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"listeners": {
						Description:         "Listeners associated with this Gateway. Listeners define logical endpoints that are bound on this Gateway's addresses. At least one Listener MUST be specified.  Each listener in a Gateway must have a unique combination of Hostname, Port, and Protocol.  An implementation MAY group Listeners by Port and then collapse each group of Listeners into a single Listener if the implementation determines that the Listeners in the group are 'compatible'. An implementation MAY also group together and collapse compatible Listeners belonging to different Gateways.  For example, an implementation might consider Listeners to be compatible with each other if all of the following conditions are met:  1. Either each Listener within the group specifies the 'HTTP'    Protocol or each Listener within the group specifies either    the 'HTTPS' or 'TLS' Protocol.  2. Each Listener within the group specifies a Hostname that is unique    within the group.  3. As a special case, one Listener within a group may omit Hostname,    in which case this Listener matches when no other Listener    matches.  If the implementation does collapse compatible Listeners, the hostname provided in the incoming client request MUST be matched to a Listener to find the correct set of Routes. The incoming hostname MUST be matched using the Hostname field for each Listener in order of most to least specific. That is, exact matches must be processed before wildcard matches.  If this field specifies multiple Listeners that have the same Port value but are not compatible, the implementation must raise a 'Conflicted' condition in the Listener status.  Support: Core",
						MarkdownDescription: "Listeners associated with this Gateway. Listeners define logical endpoints that are bound on this Gateway's addresses. At least one Listener MUST be specified.  Each listener in a Gateway must have a unique combination of Hostname, Port, and Protocol.  An implementation MAY group Listeners by Port and then collapse each group of Listeners into a single Listener if the implementation determines that the Listeners in the group are 'compatible'. An implementation MAY also group together and collapse compatible Listeners belonging to different Gateways.  For example, an implementation might consider Listeners to be compatible with each other if all of the following conditions are met:  1. Either each Listener within the group specifies the 'HTTP'    Protocol or each Listener within the group specifies either    the 'HTTPS' or 'TLS' Protocol.  2. Each Listener within the group specifies a Hostname that is unique    within the group.  3. As a special case, one Listener within a group may omit Hostname,    in which case this Listener matches when no other Listener    matches.  If the implementation does collapse compatible Listeners, the hostname provided in the incoming client request MUST be matched to a Listener to find the correct set of Routes. The incoming hostname MUST be matched using the Hostname field for each Listener in order of most to least specific. That is, exact matches must be processed before wildcard matches.  If this field specifies multiple Listeners that have the same Port value but are not compatible, the implementation must raise a 'Conflicted' condition in the Listener status.  Support: Core",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"allowed_routes": {
								Description:         "AllowedRoutes defines the types of routes that MAY be attached to a Listener and the trusted namespaces where those Route resources MAY be present.  Although a client request may match multiple route rules, only one rule may ultimately receive the request. Matching precedence MUST be determined in order of the following criteria:  * The most specific match as defined by the Route type. * The oldest Route based on creation timestamp. For example, a Route with   a creation timestamp of '2020-09-08 01:02:03' is given precedence over   a Route with a creation timestamp of '2020-09-08 01:02:04'. * If everything else is equivalent, the Route appearing first in   alphabetical order (namespace/name) should be given precedence. For   example, foo/bar is given precedence over foo/baz.  All valid rules within a Route attached to this Listener should be implemented. Invalid Route rules can be ignored (sometimes that will mean the full Route). If a Route rule transitions from valid to invalid, support for that Route rule should be dropped to ensure consistency. For example, even if a filter specified by a Route rule is invalid, the rest of the rules within that Route should still be supported.  Support: Core",
								MarkdownDescription: "AllowedRoutes defines the types of routes that MAY be attached to a Listener and the trusted namespaces where those Route resources MAY be present.  Although a client request may match multiple route rules, only one rule may ultimately receive the request. Matching precedence MUST be determined in order of the following criteria:  * The most specific match as defined by the Route type. * The oldest Route based on creation timestamp. For example, a Route with   a creation timestamp of '2020-09-08 01:02:03' is given precedence over   a Route with a creation timestamp of '2020-09-08 01:02:04'. * If everything else is equivalent, the Route appearing first in   alphabetical order (namespace/name) should be given precedence. For   example, foo/bar is given precedence over foo/baz.  All valid rules within a Route attached to this Listener should be implemented. Invalid Route rules can be ignored (sometimes that will mean the full Route). If a Route rule transitions from valid to invalid, support for that Route rule should be dropped to ensure consistency. For example, even if a filter specified by a Route rule is invalid, the rest of the rules within that Route should still be supported.  Support: Core",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"namespaces": {
										Description:         "Namespaces indicates namespaces from which Routes may be attached to this Listener. This is restricted to the namespace of this Gateway by default.  Support: Core",
										MarkdownDescription: "Namespaces indicates namespaces from which Routes may be attached to this Listener. This is restricted to the namespace of this Gateway by default.  Support: Core",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"from": {
												Description:         "From indicates where Routes will be selected for this Gateway. Possible values are: * All: Routes in all namespaces may be used by this Gateway. * Selector: Routes in namespaces selected by the selector may be used by   this Gateway. * Same: Only Routes in the same namespace may be used by this Gateway.  Support: Core",
												MarkdownDescription: "From indicates where Routes will be selected for this Gateway. Possible values are: * All: Routes in all namespaces may be used by this Gateway. * Selector: Routes in namespaces selected by the selector may be used by   this Gateway. * Same: Only Routes in the same namespace may be used by this Gateway.  Support: Core",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"selector": {
												Description:         "Selector must be specified when From is set to 'Selector'. In that case, only Routes in Namespaces matching this Selector will be selected by this Gateway. This field is ignored for other values of 'From'.  Support: Core",
												MarkdownDescription: "Selector must be specified when From is set to 'Selector'. In that case, only Routes in Namespaces matching this Selector will be selected by this Gateway. This field is ignored for other values of 'From'.  Support: Core",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"match_expressions": {
														Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
														MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "key is the label key that the selector applies to.",
																MarkdownDescription: "key is the label key that the selector applies to.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"operator": {
																Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"values": {
																Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"match_labels": {
														Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
														MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

														Type: types.MapType{ElemType: types.StringType},

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

									"kinds": {
										Description:         "Kinds specifies the groups and kinds of Routes that are allowed to bind to this Gateway Listener. When unspecified or empty, the kinds of Routes selected are determined using the Listener protocol.  A RouteGroupKind MUST correspond to kinds of Routes that are compatible with the application protocol specified in the Listener's Protocol field. If an implementation does not support or recognize this resource type, it MUST set the 'ResolvedRefs' condition to False for this Listener with the 'InvalidRoutesRef' reason.  Support: Core",
										MarkdownDescription: "Kinds specifies the groups and kinds of Routes that are allowed to bind to this Gateway Listener. When unspecified or empty, the kinds of Routes selected are determined using the Listener protocol.  A RouteGroupKind MUST correspond to kinds of Routes that are compatible with the application protocol specified in the Listener's Protocol field. If an implementation does not support or recognize this resource type, it MUST set the 'ResolvedRefs' condition to False for this Listener with the 'InvalidRoutesRef' reason.  Support: Core",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"group": {
												Description:         "Group is the group of the Route.",
												MarkdownDescription: "Group is the group of the Route.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"kind": {
												Description:         "Kind is the kind of the Route.",
												MarkdownDescription: "Kind is the kind of the Route.",

												Type: types.StringType,

												Required: true,
												Optional: false,
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

							"hostname": {
								Description:         "Hostname specifies the virtual hostname to match for protocol types that define this concept. When unspecified, all hostnames are matched. This field is ignored for protocols that don't require hostname based matching.  Implementations MUST apply Hostname matching appropriately for each of the following protocols:  * TLS: The Listener Hostname MUST match the SNI. * HTTP: The Listener Hostname MUST match the Host header of the request. * HTTPS: The Listener Hostname SHOULD match at both the TLS and HTTP   protocol layers as described above. If an implementation does not   ensure that both the SNI and Host header match the Listener hostname,   it MUST clearly document that.  For HTTPRoute and TLSRoute resources, there is an interaction with the 'spec.hostnames' array. When both listener and route specify hostnames, there MUST be an intersection between the values for a Route to be accepted. For more information, refer to the Route specific Hostnames documentation.  Support: Core",
								MarkdownDescription: "Hostname specifies the virtual hostname to match for protocol types that define this concept. When unspecified, all hostnames are matched. This field is ignored for protocols that don't require hostname based matching.  Implementations MUST apply Hostname matching appropriately for each of the following protocols:  * TLS: The Listener Hostname MUST match the SNI. * HTTP: The Listener Hostname MUST match the Host header of the request. * HTTPS: The Listener Hostname SHOULD match at both the TLS and HTTP   protocol layers as described above. If an implementation does not   ensure that both the SNI and Host header match the Listener hostname,   it MUST clearly document that.  For HTTPRoute and TLSRoute resources, there is an interaction with the 'spec.hostnames' array. When both listener and route specify hostnames, there MUST be an intersection between the values for a Route to be accepted. For more information, refer to the Route specific Hostnames documentation.  Support: Core",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "Name is the name of the Listener.  Support: Core",
								MarkdownDescription: "Name is the name of the Listener.  Support: Core",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"port": {
								Description:         "Port is the network port. Multiple listeners may use the same port, subject to the Listener compatibility rules.  Support: Core",
								MarkdownDescription: "Port is the network port. Multiple listeners may use the same port, subject to the Listener compatibility rules.  Support: Core",

								Type: types.Int64Type,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"protocol": {
								Description:         "Protocol specifies the network protocol this listener expects to receive.  Support: Core",
								MarkdownDescription: "Protocol specifies the network protocol this listener expects to receive.  Support: Core",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"tls": {
								Description:         "TLS is the TLS configuration for the Listener. This field is required if the Protocol field is 'HTTPS' or 'TLS'. It is invalid to set this field if the Protocol field is 'HTTP', 'TCP', or 'UDP'.  The association of SNIs to Certificate defined in GatewayTLSConfig is defined based on the Hostname field for this listener.  The GatewayClass MUST use the longest matching SNI out of all available certificates for any TLS handshake.  Support: Core",
								MarkdownDescription: "TLS is the TLS configuration for the Listener. This field is required if the Protocol field is 'HTTPS' or 'TLS'. It is invalid to set this field if the Protocol field is 'HTTP', 'TCP', or 'UDP'.  The association of SNIs to Certificate defined in GatewayTLSConfig is defined based on the Hostname field for this listener.  The GatewayClass MUST use the longest matching SNI out of all available certificates for any TLS handshake.  Support: Core",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"options": {
										Description:         "Options are a list of key/value pairs to enable extended TLS configuration for each implementation. For example, configuring the minimum TLS version or supported cipher suites.  A set of common keys MAY be defined by the API in the future. To avoid any ambiguity, implementation-specific definitions MUST use domain-prefixed names, such as 'example.com/my-custom-option'. Un-prefixed names are reserved for key names defined by Gateway API.  Support: Implementation-specific",
										MarkdownDescription: "Options are a list of key/value pairs to enable extended TLS configuration for each implementation. For example, configuring the minimum TLS version or supported cipher suites.  A set of common keys MAY be defined by the API in the future. To avoid any ambiguity, implementation-specific definitions MUST use domain-prefixed names, such as 'example.com/my-custom-option'. Un-prefixed names are reserved for key names defined by Gateway API.  Support: Implementation-specific",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"certificate_refs": {
										Description:         "CertificateRefs contains a series of references to Kubernetes objects that contains TLS certificates and private keys. These certificates are used to establish a TLS handshake for requests that match the hostname of the associated listener.  A single CertificateRef to a Kubernetes Secret has 'Core' support. Implementations MAY choose to support attaching multiple certificates to a Listener, but this behavior is implementation-specific.  References to a resource in different namespace are invalid UNLESS there is a ReferencePolicy in the target namespace that allows the certificate to be attached. If a ReferencePolicy does not allow this reference, the 'ResolvedRefs' condition MUST be set to False for this listener with the 'InvalidCertificateRef' reason.  This field is required to have at least one element when the mode is set to 'Terminate' (default) and is optional otherwise.  CertificateRefs can reference to standard Kubernetes resources, i.e. Secret, or implementation-specific custom resources.  Support: Core - A single reference to a Kubernetes Secret  Support: Implementation-specific (More than one reference or other resource types)",
										MarkdownDescription: "CertificateRefs contains a series of references to Kubernetes objects that contains TLS certificates and private keys. These certificates are used to establish a TLS handshake for requests that match the hostname of the associated listener.  A single CertificateRef to a Kubernetes Secret has 'Core' support. Implementations MAY choose to support attaching multiple certificates to a Listener, but this behavior is implementation-specific.  References to a resource in different namespace are invalid UNLESS there is a ReferencePolicy in the target namespace that allows the certificate to be attached. If a ReferencePolicy does not allow this reference, the 'ResolvedRefs' condition MUST be set to False for this listener with the 'InvalidCertificateRef' reason.  This field is required to have at least one element when the mode is set to 'Terminate' (default) and is optional otherwise.  CertificateRefs can reference to standard Kubernetes resources, i.e. Secret, or implementation-specific custom resources.  Support: Core - A single reference to a Kubernetes Secret  Support: Implementation-specific (More than one reference or other resource types)",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"group": {
												Description:         "Group is the group of the referent. For example, 'networking.k8s.io'. When unspecified (empty string), core API group is inferred.",
												MarkdownDescription: "Group is the group of the referent. For example, 'networking.k8s.io'. When unspecified (empty string), core API group is inferred.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"kind": {
												Description:         "Kind is kind of the referent. For example 'HTTPRoute' or 'Service'.",
												MarkdownDescription: "Kind is kind of the referent. For example 'HTTPRoute' or 'Service'.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "Name is the name of the referent.",
												MarkdownDescription: "Name is the name of the referent.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"namespace": {
												Description:         "Namespace is the namespace of the backend. When unspecified, the local namespace is inferred.  Note that when a namespace is specified, a ReferencePolicy object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferencePolicy documentation for details.  Support: Core",
												MarkdownDescription: "Namespace is the namespace of the backend. When unspecified, the local namespace is inferred.  Note that when a namespace is specified, a ReferencePolicy object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferencePolicy documentation for details.  Support: Core",

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

									"mode": {
										Description:         "Mode defines the TLS behavior for the TLS session initiated by the client. There are two possible modes:  - Terminate: The TLS session between the downstream client   and the Gateway is terminated at the Gateway. This mode requires   certificateRefs to be set and contain at least one element. - Passthrough: The TLS session is NOT terminated by the Gateway. This   implies that the Gateway can't decipher the TLS stream except for   the ClientHello message of the TLS protocol.   CertificateRefs field is ignored in this mode.  Support: Core",
										MarkdownDescription: "Mode defines the TLS behavior for the TLS session initiated by the client. There are two possible modes:  - Terminate: The TLS session between the downstream client   and the Gateway is terminated at the Gateway. This mode requires   certificateRefs to be set and contain at least one element. - Passthrough: The TLS session is NOT terminated by the Gateway. This   implies that the Gateway can't decipher the TLS stream except for   the ClientHello message of the TLS protocol.   CertificateRefs field is ignored in this mode.  Support: Core",

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

						Required: true,
						Optional: false,
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

func (r *GatewayNetworkingK8SIoGatewayV1Alpha2Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_gateway_networking_k8s_io_gateway_v1alpha2")

	var state GatewayNetworkingK8SIoGatewayV1Alpha2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel GatewayNetworkingK8SIoGatewayV1Alpha2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("gateway.networking.k8s.io/v1alpha2")
	goModel.Kind = utilities.Ptr("Gateway")

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

func (r *GatewayNetworkingK8SIoGatewayV1Alpha2Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_gateway_networking_k8s_io_gateway_v1alpha2")
	// NO-OP: All data is already in Terraform state
}

func (r *GatewayNetworkingK8SIoGatewayV1Alpha2Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_gateway_networking_k8s_io_gateway_v1alpha2")

	var state GatewayNetworkingK8SIoGatewayV1Alpha2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel GatewayNetworkingK8SIoGatewayV1Alpha2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("gateway.networking.k8s.io/v1alpha2")
	goModel.Kind = utilities.Ptr("Gateway")

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

func (r *GatewayNetworkingK8SIoGatewayV1Alpha2Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_gateway_networking_k8s_io_gateway_v1alpha2")
	// NO-OP: Terraform removes the state automatically for us
}
