/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package gateway_networking_k8s_io_v1alpha2

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &GatewayNetworkingK8SIoGrpcrouteV1Alpha2DataSource{}
	_ datasource.DataSourceWithConfigure = &GatewayNetworkingK8SIoGrpcrouteV1Alpha2DataSource{}
)

func NewGatewayNetworkingK8SIoGrpcrouteV1Alpha2DataSource() datasource.DataSource {
	return &GatewayNetworkingK8SIoGrpcrouteV1Alpha2DataSource{}
}

type GatewayNetworkingK8SIoGrpcrouteV1Alpha2DataSource struct {
	kubernetesClient dynamic.Interface
}

type GatewayNetworkingK8SIoGrpcrouteV1Alpha2DataSourceData struct {
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
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"filters" json:"filters,omitempty"`
			Matches *[]struct {
				Headers *[]struct {
					Name  *string `tfsdk:"name" json:"name,omitempty"`
					Type  *string `tfsdk:"type" json:"type,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"headers" json:"headers,omitempty"`
				Method *struct {
					Method  *string `tfsdk:"method" json:"method,omitempty"`
					Service *string `tfsdk:"service" json:"service,omitempty"`
					Type    *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"method" json:"method,omitempty"`
			} `tfsdk:"matches" json:"matches,omitempty"`
		} `tfsdk:"rules" json:"rules,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *GatewayNetworkingK8SIoGrpcrouteV1Alpha2DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_gateway_networking_k8s_io_grpc_route_v1alpha2"
}

func (r *GatewayNetworkingK8SIoGrpcrouteV1Alpha2DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "GRPCRoute provides a way to route gRPC requests. This includes the capability to match requests by hostname, gRPC service, gRPC method, or HTTP/2 header. Filters can be used to specify additional processing steps. Backends specify where matching requests will be routed.  GRPCRoute falls under extended support within the Gateway API. Within the following specification, the word 'MUST' indicates that an implementation supporting GRPCRoute must conform to the indicated requirement, but an implementation not supporting this route type need not follow the requirement unless explicitly indicated.  Implementations supporting 'GRPCRoute' with the 'HTTPS' 'ProtocolType' MUST accept HTTP/2 connections without an initial upgrade from HTTP/1.1, i.e. via ALPN. If the implementation does not support this, then it MUST set the 'Accepted' condition to 'False' for the affected listener with a reason of 'UnsupportedProtocol'.  Implementations MAY also accept HTTP/2 connections with an upgrade from HTTP/1.  Implementations supporting 'GRPCRoute' with the 'HTTP' 'ProtocolType' MUST support HTTP/2 over cleartext TCP (h2c, https://www.rfc-editor.org/rfc/rfc7540#section-3.1) without an initial upgrade from HTTP/1.1, i.e. with prior knowledge (https://www.rfc-editor.org/rfc/rfc7540#section-3.4). If the implementation does not support this, then it MUST set the 'Accepted' condition to 'False' for the affected listener with a reason of 'UnsupportedProtocol'. Implementations MAY also accept HTTP/2 connections with an upgrade from HTTP/1, i.e. without prior knowledge.",
		MarkdownDescription: "GRPCRoute provides a way to route gRPC requests. This includes the capability to match requests by hostname, gRPC service, gRPC method, or HTTP/2 header. Filters can be used to specify additional processing steps. Backends specify where matching requests will be routed.  GRPCRoute falls under extended support within the Gateway API. Within the following specification, the word 'MUST' indicates that an implementation supporting GRPCRoute must conform to the indicated requirement, but an implementation not supporting this route type need not follow the requirement unless explicitly indicated.  Implementations supporting 'GRPCRoute' with the 'HTTPS' 'ProtocolType' MUST accept HTTP/2 connections without an initial upgrade from HTTP/1.1, i.e. via ALPN. If the implementation does not support this, then it MUST set the 'Accepted' condition to 'False' for the affected listener with a reason of 'UnsupportedProtocol'.  Implementations MAY also accept HTTP/2 connections with an upgrade from HTTP/1.  Implementations supporting 'GRPCRoute' with the 'HTTP' 'ProtocolType' MUST support HTTP/2 over cleartext TCP (h2c, https://www.rfc-editor.org/rfc/rfc7540#section-3.1) without an initial upgrade from HTTP/1.1, i.e. with prior knowledge (https://www.rfc-editor.org/rfc/rfc7540#section-3.4). If the implementation does not support this, then it MUST set the 'Accepted' condition to 'False' for the affected listener with a reason of 'UnsupportedProtocol'. Implementations MAY also accept HTTP/2 connections with an upgrade from HTTP/1, i.e. without prior knowledge.",
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
				Description:         "Spec defines the desired state of GRPCRoute.",
				MarkdownDescription: "Spec defines the desired state of GRPCRoute.",
				Attributes: map[string]schema.Attribute{
					"hostnames": schema.ListAttribute{
						Description:         "Hostnames defines a set of hostnames to match against the GRPC Host header to select a GRPCRoute to process the request. This matches the RFC 1123 definition of a hostname with 2 notable exceptions:  1. IPs are not allowed. 2. A hostname may be prefixed with a wildcard label ('*.'). The wildcard label MUST appear by itself as the first label.  If a hostname is specified by both the Listener and GRPCRoute, there MUST be at least one intersecting hostname for the GRPCRoute to be attached to the Listener. For example:  * A Listener with 'test.example.com' as the hostname matches GRPCRoutes that have either not specified any hostnames, or have specified at least one of 'test.example.com' or '*.example.com'. * A Listener with '*.example.com' as the hostname matches GRPCRoutes that have either not specified any hostnames or have specified at least one hostname that matches the Listener hostname. For example, 'test.example.com' and '*.example.com' would both match. On the other hand, 'example.com' and 'test.example.net' would not match.  Hostnames that are prefixed with a wildcard label ('*.') are interpreted as a suffix match. That means that a match for '*.example.com' would match both 'test.example.com', and 'foo.test.example.com', but not 'example.com'.  If both the Listener and GRPCRoute have specified hostnames, any GRPCRoute hostnames that do not match the Listener hostname MUST be ignored. For example, if a Listener specified '*.example.com', and the GRPCRoute specified 'test.example.com' and 'test.example.net', 'test.example.net' MUST NOT be considered for a match.  If both the Listener and GRPCRoute have specified hostnames, and none match with the criteria above, then the GRPCRoute MUST NOT be accepted by the implementation. The implementation MUST raise an 'Accepted' Condition with a status of 'False' in the corresponding RouteParentStatus.  If a Route (A) of type HTTPRoute or GRPCRoute is attached to a Listener and that listener already has another Route (B) of the other type attached and the intersection of the hostnames of A and B is non-empty, then the implementation MUST accept exactly one of these two routes, determined by the following criteria, in order:  * The oldest Route based on creation timestamp. * The Route appearing first in alphabetical order by '{namespace}/{name}'.  The rejected Route MUST raise an 'Accepted' condition with a status of 'False' in the corresponding RouteParentStatus.  Support: Core",
						MarkdownDescription: "Hostnames defines a set of hostnames to match against the GRPC Host header to select a GRPCRoute to process the request. This matches the RFC 1123 definition of a hostname with 2 notable exceptions:  1. IPs are not allowed. 2. A hostname may be prefixed with a wildcard label ('*.'). The wildcard label MUST appear by itself as the first label.  If a hostname is specified by both the Listener and GRPCRoute, there MUST be at least one intersecting hostname for the GRPCRoute to be attached to the Listener. For example:  * A Listener with 'test.example.com' as the hostname matches GRPCRoutes that have either not specified any hostnames, or have specified at least one of 'test.example.com' or '*.example.com'. * A Listener with '*.example.com' as the hostname matches GRPCRoutes that have either not specified any hostnames or have specified at least one hostname that matches the Listener hostname. For example, 'test.example.com' and '*.example.com' would both match. On the other hand, 'example.com' and 'test.example.net' would not match.  Hostnames that are prefixed with a wildcard label ('*.') are interpreted as a suffix match. That means that a match for '*.example.com' would match both 'test.example.com', and 'foo.test.example.com', but not 'example.com'.  If both the Listener and GRPCRoute have specified hostnames, any GRPCRoute hostnames that do not match the Listener hostname MUST be ignored. For example, if a Listener specified '*.example.com', and the GRPCRoute specified 'test.example.com' and 'test.example.net', 'test.example.net' MUST NOT be considered for a match.  If both the Listener and GRPCRoute have specified hostnames, and none match with the criteria above, then the GRPCRoute MUST NOT be accepted by the implementation. The implementation MUST raise an 'Accepted' Condition with a status of 'False' in the corresponding RouteParentStatus.  If a Route (A) of type HTTPRoute or GRPCRoute is attached to a Listener and that listener already has another Route (B) of the other type attached and the intersection of the hostnames of A and B is non-empty, then the implementation MUST accept exactly one of these two routes, determined by the following criteria, in order:  * The oldest Route based on creation timestamp. * The Route appearing first in alphabetical order by '{namespace}/{name}'.  The rejected Route MUST raise an 'Accepted' condition with a status of 'False' in the corresponding RouteParentStatus.  Support: Core",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"parent_refs": schema.ListNestedAttribute{
						Description:         "ParentRefs references the resources (usually Gateways) that a Route wants to be attached to. Note that the referenced parent resource needs to allow this for the attachment to be complete. For Gateways, that means the Gateway needs to allow attachment from Routes of this kind and namespace. For Services, that means the Service must either be in the same namespace for a 'producer' route, or the mesh implementation must support and allow 'consumer' routes for the referenced Service. ReferenceGrant is not applicable for governing ParentRefs to Services - it is not possible to create a 'producer' route for a Service in a different namespace from the Route.  There are two kinds of parent resources with 'Core' support:  * Gateway (Gateway conformance profile) * Service (Mesh conformance profile, experimental, ClusterIP Services only)  This API may be extended in the future to support additional kinds of parent resources.  It is invalid to reference an identical parent more than once. It is valid to reference multiple distinct sections within the same parent resource, such as two separate Listeners on the same Gateway or two separate ports on the same Service.  It is possible to separately reference multiple distinct objects that may be collapsed by an implementation. For example, some implementations may choose to merge compatible Gateway Listeners together. If that is the case, the list of routes attached to those resources should also be merged.  Note that for ParentRefs that cross namespace boundaries, there are specific rules. Cross-namespace references are only valid if they are explicitly allowed by something in the namespace they are referring to. For example, Gateway has the AllowedRoutes field, and ReferenceGrant provides a generic way to enable other kinds of cross-namespace reference.  ParentRefs from a Route to a Service in the same namespace are 'producer' routes, which apply default routing rules to inbound connections from any namespace to the Service.  ParentRefs from a Route to a Service in a different namespace are 'consumer' routes, and these routing rules are only applied to outbound connections originating from the same namespace as the Route, for which the intended destination of the connections are a Service targeted as a ParentRef of the Route.  ",
						MarkdownDescription: "ParentRefs references the resources (usually Gateways) that a Route wants to be attached to. Note that the referenced parent resource needs to allow this for the attachment to be complete. For Gateways, that means the Gateway needs to allow attachment from Routes of this kind and namespace. For Services, that means the Service must either be in the same namespace for a 'producer' route, or the mesh implementation must support and allow 'consumer' routes for the referenced Service. ReferenceGrant is not applicable for governing ParentRefs to Services - it is not possible to create a 'producer' route for a Service in a different namespace from the Route.  There are two kinds of parent resources with 'Core' support:  * Gateway (Gateway conformance profile) * Service (Mesh conformance profile, experimental, ClusterIP Services only)  This API may be extended in the future to support additional kinds of parent resources.  It is invalid to reference an identical parent more than once. It is valid to reference multiple distinct sections within the same parent resource, such as two separate Listeners on the same Gateway or two separate ports on the same Service.  It is possible to separately reference multiple distinct objects that may be collapsed by an implementation. For example, some implementations may choose to merge compatible Gateway Listeners together. If that is the case, the list of routes attached to those resources should also be merged.  Note that for ParentRefs that cross namespace boundaries, there are specific rules. Cross-namespace references are only valid if they are explicitly allowed by something in the namespace they are referring to. For example, Gateway has the AllowedRoutes field, and ReferenceGrant provides a generic way to enable other kinds of cross-namespace reference.  ParentRefs from a Route to a Service in the same namespace are 'producer' routes, which apply default routing rules to inbound connections from any namespace to the Service.  ParentRefs from a Route to a Service in a different namespace are 'consumer' routes, and these routing rules are only applied to outbound connections originating from the same namespace as the Route, for which the intended destination of the connections are a Service targeted as a ParentRef of the Route.  ",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"group": schema.StringAttribute{
									Description:         "Group is the group of the referent. When unspecified, 'gateway.networking.k8s.io' is inferred. To set the core API group (such as for a 'Service' kind referent), Group must be explicitly set to '' (empty string).  Support: Core",
									MarkdownDescription: "Group is the group of the referent. When unspecified, 'gateway.networking.k8s.io' is inferred. To set the core API group (such as for a 'Service' kind referent), Group must be explicitly set to '' (empty string).  Support: Core",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"kind": schema.StringAttribute{
									Description:         "Kind is kind of the referent.  There are two kinds of parent resources with 'Core' support:  * Gateway (Gateway conformance profile) * Service (Mesh conformance profile, experimental, ClusterIP Services only)  Support for other resources is Implementation-Specific.",
									MarkdownDescription: "Kind is kind of the referent.  There are two kinds of parent resources with 'Core' support:  * Gateway (Gateway conformance profile) * Service (Mesh conformance profile, experimental, ClusterIP Services only)  Support for other resources is Implementation-Specific.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"name": schema.StringAttribute{
									Description:         "Name is the name of the referent.  Support: Core",
									MarkdownDescription: "Name is the name of the referent.  Support: Core",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace is the namespace of the referent. When unspecified, this refers to the local namespace of the Route.  Note that there are specific rules for ParentRefs which cross namespace boundaries. Cross-namespace references are only valid if they are explicitly allowed by something in the namespace they are referring to. For example: Gateway has the AllowedRoutes field, and ReferenceGrant provides a generic way to enable any other kind of cross-namespace reference.  ParentRefs from a Route to a Service in the same namespace are 'producer' routes, which apply default routing rules to inbound connections from any namespace to the Service.  ParentRefs from a Route to a Service in a different namespace are 'consumer' routes, and these routing rules are only applied to outbound connections originating from the same namespace as the Route, for which the intended destination of the connections are a Service targeted as a ParentRef of the Route.  Support: Core",
									MarkdownDescription: "Namespace is the namespace of the referent. When unspecified, this refers to the local namespace of the Route.  Note that there are specific rules for ParentRefs which cross namespace boundaries. Cross-namespace references are only valid if they are explicitly allowed by something in the namespace they are referring to. For example: Gateway has the AllowedRoutes field, and ReferenceGrant provides a generic way to enable any other kind of cross-namespace reference.  ParentRefs from a Route to a Service in the same namespace are 'producer' routes, which apply default routing rules to inbound connections from any namespace to the Service.  ParentRefs from a Route to a Service in a different namespace are 'consumer' routes, and these routing rules are only applied to outbound connections originating from the same namespace as the Route, for which the intended destination of the connections are a Service targeted as a ParentRef of the Route.  Support: Core",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"port": schema.Int64Attribute{
									Description:         "Port is the network port this Route targets. It can be interpreted differently based on the type of parent resource.  When the parent resource is a Gateway, this targets all listeners listening on the specified port that also support this kind of Route(and select this Route). It's not recommended to set 'Port' unless the networking behaviors specified in a Route must apply to a specific port as opposed to a listener(s) whose port(s) may be changed. When both Port and SectionName are specified, the name and port of the selected listener must match both specified values.  When the parent resource is a Service, this targets a specific port in the Service spec. When both Port (experimental) and SectionName are specified, the name and port of the selected port must match both specified values.  Implementations MAY choose to support other parent resources. Implementations supporting other types of parent resources MUST clearly document how/if Port is interpreted.  For the purpose of status, an attachment is considered successful as long as the parent resource accepts it partially. For example, Gateway listeners can restrict which Routes can attach to them by Route kind, namespace, or hostname. If 1 of 2 Gateway listeners accept attachment from the referencing Route, the Route MUST be considered successfully attached. If no Gateway listeners accept attachment from this Route, the Route MUST be considered detached from the Gateway.  Support: Extended  ",
									MarkdownDescription: "Port is the network port this Route targets. It can be interpreted differently based on the type of parent resource.  When the parent resource is a Gateway, this targets all listeners listening on the specified port that also support this kind of Route(and select this Route). It's not recommended to set 'Port' unless the networking behaviors specified in a Route must apply to a specific port as opposed to a listener(s) whose port(s) may be changed. When both Port and SectionName are specified, the name and port of the selected listener must match both specified values.  When the parent resource is a Service, this targets a specific port in the Service spec. When both Port (experimental) and SectionName are specified, the name and port of the selected port must match both specified values.  Implementations MAY choose to support other parent resources. Implementations supporting other types of parent resources MUST clearly document how/if Port is interpreted.  For the purpose of status, an attachment is considered successful as long as the parent resource accepts it partially. For example, Gateway listeners can restrict which Routes can attach to them by Route kind, namespace, or hostname. If 1 of 2 Gateway listeners accept attachment from the referencing Route, the Route MUST be considered successfully attached. If no Gateway listeners accept attachment from this Route, the Route MUST be considered detached from the Gateway.  Support: Extended  ",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"section_name": schema.StringAttribute{
									Description:         "SectionName is the name of a section within the target resource. In the following resources, SectionName is interpreted as the following:  * Gateway: Listener Name. When both Port (experimental) and SectionName are specified, the name and port of the selected listener must match both specified values. * Service: Port Name. When both Port (experimental) and SectionName are specified, the name and port of the selected listener must match both specified values. Note that attaching Routes to Services as Parents is part of experimental Mesh support and is not supported for any other purpose.  Implementations MAY choose to support attaching Routes to other resources. If that is the case, they MUST clearly document how SectionName is interpreted.  When unspecified (empty string), this will reference the entire resource. For the purpose of status, an attachment is considered successful if at least one section in the parent resource accepts it. For example, Gateway listeners can restrict which Routes can attach to them by Route kind, namespace, or hostname. If 1 of 2 Gateway listeners accept attachment from the referencing Route, the Route MUST be considered successfully attached. If no Gateway listeners accept attachment from this Route, the Route MUST be considered detached from the Gateway.  Support: Core",
									MarkdownDescription: "SectionName is the name of a section within the target resource. In the following resources, SectionName is interpreted as the following:  * Gateway: Listener Name. When both Port (experimental) and SectionName are specified, the name and port of the selected listener must match both specified values. * Service: Port Name. When both Port (experimental) and SectionName are specified, the name and port of the selected listener must match both specified values. Note that attaching Routes to Services as Parents is part of experimental Mesh support and is not supported for any other purpose.  Implementations MAY choose to support attaching Routes to other resources. If that is the case, they MUST clearly document how SectionName is interpreted.  When unspecified (empty string), this will reference the entire resource. For the purpose of status, an attachment is considered successful if at least one section in the parent resource accepts it. For example, Gateway listeners can restrict which Routes can attach to them by Route kind, namespace, or hostname. If 1 of 2 Gateway listeners accept attachment from the referencing Route, the Route MUST be considered successfully attached. If no Gateway listeners accept attachment from this Route, the Route MUST be considered detached from the Gateway.  Support: Core",
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

					"rules": schema.ListNestedAttribute{
						Description:         "Rules are a list of GRPC matchers, filters and actions.",
						MarkdownDescription: "Rules are a list of GRPC matchers, filters and actions.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"backend_refs": schema.ListNestedAttribute{
									Description:         "BackendRefs defines the backend(s) where matching requests should be sent.  Failure behavior here depends on how many BackendRefs are specified and how many are invalid.  If *all* entries in BackendRefs are invalid, and there are also no filters specified in this route rule, *all* traffic which matches this rule MUST receive an 'UNAVAILABLE' status.  See the GRPCBackendRef definition for the rules about what makes a single GRPCBackendRef invalid.  When a GRPCBackendRef is invalid, 'UNAVAILABLE' statuses MUST be returned for requests that would have otherwise been routed to an invalid backend. If multiple backends are specified, and some are invalid, the proportion of requests that would otherwise have been routed to an invalid backend MUST receive an 'UNAVAILABLE' status.  For example, if two backends are specified with equal weights, and one is invalid, 50 percent of traffic MUST receive an 'UNAVAILABLE' status. Implementations may choose how that 50 percent is determined.  Support: Core for Kubernetes Service  Support: Implementation-specific for any other resource  Support for weight: Core",
									MarkdownDescription: "BackendRefs defines the backend(s) where matching requests should be sent.  Failure behavior here depends on how many BackendRefs are specified and how many are invalid.  If *all* entries in BackendRefs are invalid, and there are also no filters specified in this route rule, *all* traffic which matches this rule MUST receive an 'UNAVAILABLE' status.  See the GRPCBackendRef definition for the rules about what makes a single GRPCBackendRef invalid.  When a GRPCBackendRef is invalid, 'UNAVAILABLE' statuses MUST be returned for requests that would have otherwise been routed to an invalid backend. If multiple backends are specified, and some are invalid, the proportion of requests that would otherwise have been routed to an invalid backend MUST receive an 'UNAVAILABLE' status.  For example, if two backends are specified with equal weights, and one is invalid, 50 percent of traffic MUST receive an 'UNAVAILABLE' status. Implementations may choose how that 50 percent is determined.  Support: Core for Kubernetes Service  Support: Implementation-specific for any other resource  Support for weight: Core",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"filters": schema.ListNestedAttribute{
												Description:         "Filters defined at this level MUST be executed if and only if the request is being forwarded to the backend defined here.  Support: Implementation-specific (For broader support of filters, use the Filters field in GRPCRouteRule.)",
												MarkdownDescription: "Filters defined at this level MUST be executed if and only if the request is being forwarded to the backend defined here.  Support: Implementation-specific (For broader support of filters, use the Filters field in GRPCRouteRule.)",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"extension_ref": schema.SingleNestedAttribute{
															Description:         "ExtensionRef is an optional, implementation-specific extension to the 'filter' behavior.  For example, resource 'myroutefilter' in group 'networking.example.net'). ExtensionRef MUST NOT be used for core and extended filters.  Support: Implementation-specific  This filter can be used multiple times within the same rule.",
															MarkdownDescription: "ExtensionRef is an optional, implementation-specific extension to the 'filter' behavior.  For example, resource 'myroutefilter' in group 'networking.example.net'). ExtensionRef MUST NOT be used for core and extended filters.  Support: Implementation-specific  This filter can be used multiple times within the same rule.",
															Attributes: map[string]schema.Attribute{
																"group": schema.StringAttribute{
																	Description:         "Group is the group of the referent. For example, 'gateway.networking.k8s.io'. When unspecified or empty string, core API group is inferred.",
																	MarkdownDescription: "Group is the group of the referent. For example, 'gateway.networking.k8s.io'. When unspecified or empty string, core API group is inferred.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"kind": schema.StringAttribute{
																	Description:         "Kind is kind of the referent. For example 'HTTPRoute' or 'Service'.",
																	MarkdownDescription: "Kind is kind of the referent. For example 'HTTPRoute' or 'Service'.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"name": schema.StringAttribute{
																	Description:         "Name is the name of the referent.",
																	MarkdownDescription: "Name is the name of the referent.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"request_header_modifier": schema.SingleNestedAttribute{
															Description:         "RequestHeaderModifier defines a schema for a filter that modifies request headers.  Support: Core",
															MarkdownDescription: "RequestHeaderModifier defines a schema for a filter that modifies request headers.  Support: Core",
															Attributes: map[string]schema.Attribute{
																"add": schema.ListNestedAttribute{
																	Description:         "Add adds the given header(s) (name, value) to the request before the action. It appends to any existing values associated with the header name.  Input: GET /foo HTTP/1.1 my-header: foo  Config: add: - name: 'my-header' value: 'bar,baz'  Output: GET /foo HTTP/1.1 my-header: foo,bar,baz",
																	MarkdownDescription: "Add adds the given header(s) (name, value) to the request before the action. It appends to any existing values associated with the header name.  Input: GET /foo HTTP/1.1 my-header: foo  Config: add: - name: 'my-header' value: 'bar,baz'  Output: GET /foo HTTP/1.1 my-header: foo,bar,baz",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "Name is the name of the HTTP Header to be matched. Name matching MUST be case insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).  If multiple entries specify equivalent header names, the first entry with an equivalent name MUST be considered for a match. Subsequent entries with an equivalent header name MUST be ignored. Due to the case-insensitivity of header names, 'foo' and 'Foo' are considered equivalent.",
																				MarkdownDescription: "Name is the name of the HTTP Header to be matched. Name matching MUST be case insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).  If multiple entries specify equivalent header names, the first entry with an equivalent name MUST be considered for a match. Subsequent entries with an equivalent header name MUST be ignored. Due to the case-insensitivity of header names, 'foo' and 'Foo' are considered equivalent.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"value": schema.StringAttribute{
																				Description:         "Value is the value of HTTP Header to be matched.",
																				MarkdownDescription: "Value is the value of HTTP Header to be matched.",
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

																"remove": schema.ListAttribute{
																	Description:         "Remove the given header(s) from the HTTP request before the action. The value of Remove is a list of HTTP header names. Note that the header names are case-insensitive (see https://datatracker.ietf.org/doc/html/rfc2616#section-4.2).  Input: GET /foo HTTP/1.1 my-header1: foo my-header2: bar my-header3: baz  Config: remove: ['my-header1', 'my-header3']  Output: GET /foo HTTP/1.1 my-header2: bar",
																	MarkdownDescription: "Remove the given header(s) from the HTTP request before the action. The value of Remove is a list of HTTP header names. Note that the header names are case-insensitive (see https://datatracker.ietf.org/doc/html/rfc2616#section-4.2).  Input: GET /foo HTTP/1.1 my-header1: foo my-header2: bar my-header3: baz  Config: remove: ['my-header1', 'my-header3']  Output: GET /foo HTTP/1.1 my-header2: bar",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"set": schema.ListNestedAttribute{
																	Description:         "Set overwrites the request with the given header (name, value) before the action.  Input: GET /foo HTTP/1.1 my-header: foo  Config: set: - name: 'my-header' value: 'bar'  Output: GET /foo HTTP/1.1 my-header: bar",
																	MarkdownDescription: "Set overwrites the request with the given header (name, value) before the action.  Input: GET /foo HTTP/1.1 my-header: foo  Config: set: - name: 'my-header' value: 'bar'  Output: GET /foo HTTP/1.1 my-header: bar",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "Name is the name of the HTTP Header to be matched. Name matching MUST be case insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).  If multiple entries specify equivalent header names, the first entry with an equivalent name MUST be considered for a match. Subsequent entries with an equivalent header name MUST be ignored. Due to the case-insensitivity of header names, 'foo' and 'Foo' are considered equivalent.",
																				MarkdownDescription: "Name is the name of the HTTP Header to be matched. Name matching MUST be case insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).  If multiple entries specify equivalent header names, the first entry with an equivalent name MUST be considered for a match. Subsequent entries with an equivalent header name MUST be ignored. Due to the case-insensitivity of header names, 'foo' and 'Foo' are considered equivalent.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"value": schema.StringAttribute{
																				Description:         "Value is the value of HTTP Header to be matched.",
																				MarkdownDescription: "Value is the value of HTTP Header to be matched.",
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
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"request_mirror": schema.SingleNestedAttribute{
															Description:         "RequestMirror defines a schema for a filter that mirrors requests. Requests are sent to the specified destination, but responses from that destination are ignored.  This filter can be used multiple times within the same rule. Note that not all implementations will be able to support mirroring to multiple backends.  Support: Extended",
															MarkdownDescription: "RequestMirror defines a schema for a filter that mirrors requests. Requests are sent to the specified destination, but responses from that destination are ignored.  This filter can be used multiple times within the same rule. Note that not all implementations will be able to support mirroring to multiple backends.  Support: Extended",
															Attributes: map[string]schema.Attribute{
																"backend_ref": schema.SingleNestedAttribute{
																	Description:         "BackendRef references a resource where mirrored requests are sent.  Mirrored requests must be sent only to a single destination endpoint within this BackendRef, irrespective of how many endpoints are present within this BackendRef.  If the referent cannot be found, this BackendRef is invalid and must be dropped from the Gateway. The controller must ensure the 'ResolvedRefs' condition on the Route status is set to 'status: False' and not configure this backend in the underlying implementation.  If there is a cross-namespace reference to an *existing* object that is not allowed by a ReferenceGrant, the controller must ensure the 'ResolvedRefs'  condition on the Route is set to 'status: False', with the 'RefNotPermitted' reason and not configure this backend in the underlying implementation.  In either error case, the Message of the 'ResolvedRefs' Condition should be used to provide more detail about the problem.  Support: Extended for Kubernetes Service  Support: Implementation-specific for any other resource",
																	MarkdownDescription: "BackendRef references a resource where mirrored requests are sent.  Mirrored requests must be sent only to a single destination endpoint within this BackendRef, irrespective of how many endpoints are present within this BackendRef.  If the referent cannot be found, this BackendRef is invalid and must be dropped from the Gateway. The controller must ensure the 'ResolvedRefs' condition on the Route status is set to 'status: False' and not configure this backend in the underlying implementation.  If there is a cross-namespace reference to an *existing* object that is not allowed by a ReferenceGrant, the controller must ensure the 'ResolvedRefs'  condition on the Route is set to 'status: False', with the 'RefNotPermitted' reason and not configure this backend in the underlying implementation.  In either error case, the Message of the 'ResolvedRefs' Condition should be used to provide more detail about the problem.  Support: Extended for Kubernetes Service  Support: Implementation-specific for any other resource",
																	Attributes: map[string]schema.Attribute{
																		"group": schema.StringAttribute{
																			Description:         "Group is the group of the referent. For example, 'gateway.networking.k8s.io'. When unspecified or empty string, core API group is inferred.",
																			MarkdownDescription: "Group is the group of the referent. For example, 'gateway.networking.k8s.io'. When unspecified or empty string, core API group is inferred.",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"kind": schema.StringAttribute{
																			Description:         "Kind is the Kubernetes resource kind of the referent. For example 'Service'.  Defaults to 'Service' when not specified.  ExternalName services can refer to CNAME DNS records that may live outside of the cluster and as such are difficult to reason about in terms of conformance. They also may not be safe to forward to (see CVE-2021-25740 for more information). Implementations SHOULD NOT support ExternalName Services.  Support: Core (Services with a type other than ExternalName)  Support: Implementation-specific (Services with type ExternalName)",
																			MarkdownDescription: "Kind is the Kubernetes resource kind of the referent. For example 'Service'.  Defaults to 'Service' when not specified.  ExternalName services can refer to CNAME DNS records that may live outside of the cluster and as such are difficult to reason about in terms of conformance. They also may not be safe to forward to (see CVE-2021-25740 for more information). Implementations SHOULD NOT support ExternalName Services.  Support: Core (Services with a type other than ExternalName)  Support: Implementation-specific (Services with type ExternalName)",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"name": schema.StringAttribute{
																			Description:         "Name is the name of the referent.",
																			MarkdownDescription: "Name is the name of the referent.",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"namespace": schema.StringAttribute{
																			Description:         "Namespace is the namespace of the backend. When unspecified, the local namespace is inferred.  Note that when a namespace different than the local namespace is specified, a ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details.  Support: Core",
																			MarkdownDescription: "Namespace is the namespace of the backend. When unspecified, the local namespace is inferred.  Note that when a namespace different than the local namespace is specified, a ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details.  Support: Core",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"port": schema.Int64Attribute{
																			Description:         "Port specifies the destination port number to use for this resource. Port is required when the referent is a Kubernetes Service. In this case, the port number is the service port number, not the target port. For other resources, destination port might be derived from the referent resource or this field.",
																			MarkdownDescription: "Port specifies the destination port number to use for this resource. Port is required when the referent is a Kubernetes Service. In this case, the port number is the service port number, not the target port. For other resources, destination port might be derived from the referent resource or this field.",
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

														"response_header_modifier": schema.SingleNestedAttribute{
															Description:         "ResponseHeaderModifier defines a schema for a filter that modifies response headers.  Support: Extended",
															MarkdownDescription: "ResponseHeaderModifier defines a schema for a filter that modifies response headers.  Support: Extended",
															Attributes: map[string]schema.Attribute{
																"add": schema.ListNestedAttribute{
																	Description:         "Add adds the given header(s) (name, value) to the request before the action. It appends to any existing values associated with the header name.  Input: GET /foo HTTP/1.1 my-header: foo  Config: add: - name: 'my-header' value: 'bar,baz'  Output: GET /foo HTTP/1.1 my-header: foo,bar,baz",
																	MarkdownDescription: "Add adds the given header(s) (name, value) to the request before the action. It appends to any existing values associated with the header name.  Input: GET /foo HTTP/1.1 my-header: foo  Config: add: - name: 'my-header' value: 'bar,baz'  Output: GET /foo HTTP/1.1 my-header: foo,bar,baz",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "Name is the name of the HTTP Header to be matched. Name matching MUST be case insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).  If multiple entries specify equivalent header names, the first entry with an equivalent name MUST be considered for a match. Subsequent entries with an equivalent header name MUST be ignored. Due to the case-insensitivity of header names, 'foo' and 'Foo' are considered equivalent.",
																				MarkdownDescription: "Name is the name of the HTTP Header to be matched. Name matching MUST be case insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).  If multiple entries specify equivalent header names, the first entry with an equivalent name MUST be considered for a match. Subsequent entries with an equivalent header name MUST be ignored. Due to the case-insensitivity of header names, 'foo' and 'Foo' are considered equivalent.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"value": schema.StringAttribute{
																				Description:         "Value is the value of HTTP Header to be matched.",
																				MarkdownDescription: "Value is the value of HTTP Header to be matched.",
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

																"remove": schema.ListAttribute{
																	Description:         "Remove the given header(s) from the HTTP request before the action. The value of Remove is a list of HTTP header names. Note that the header names are case-insensitive (see https://datatracker.ietf.org/doc/html/rfc2616#section-4.2).  Input: GET /foo HTTP/1.1 my-header1: foo my-header2: bar my-header3: baz  Config: remove: ['my-header1', 'my-header3']  Output: GET /foo HTTP/1.1 my-header2: bar",
																	MarkdownDescription: "Remove the given header(s) from the HTTP request before the action. The value of Remove is a list of HTTP header names. Note that the header names are case-insensitive (see https://datatracker.ietf.org/doc/html/rfc2616#section-4.2).  Input: GET /foo HTTP/1.1 my-header1: foo my-header2: bar my-header3: baz  Config: remove: ['my-header1', 'my-header3']  Output: GET /foo HTTP/1.1 my-header2: bar",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"set": schema.ListNestedAttribute{
																	Description:         "Set overwrites the request with the given header (name, value) before the action.  Input: GET /foo HTTP/1.1 my-header: foo  Config: set: - name: 'my-header' value: 'bar'  Output: GET /foo HTTP/1.1 my-header: bar",
																	MarkdownDescription: "Set overwrites the request with the given header (name, value) before the action.  Input: GET /foo HTTP/1.1 my-header: foo  Config: set: - name: 'my-header' value: 'bar'  Output: GET /foo HTTP/1.1 my-header: bar",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "Name is the name of the HTTP Header to be matched. Name matching MUST be case insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).  If multiple entries specify equivalent header names, the first entry with an equivalent name MUST be considered for a match. Subsequent entries with an equivalent header name MUST be ignored. Due to the case-insensitivity of header names, 'foo' and 'Foo' are considered equivalent.",
																				MarkdownDescription: "Name is the name of the HTTP Header to be matched. Name matching MUST be case insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).  If multiple entries specify equivalent header names, the first entry with an equivalent name MUST be considered for a match. Subsequent entries with an equivalent header name MUST be ignored. Due to the case-insensitivity of header names, 'foo' and 'Foo' are considered equivalent.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"value": schema.StringAttribute{
																				Description:         "Value is the value of HTTP Header to be matched.",
																				MarkdownDescription: "Value is the value of HTTP Header to be matched.",
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
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"type": schema.StringAttribute{
															Description:         "Type identifies the type of filter to apply. As with other API fields, types are classified into three conformance levels:  - Core: Filter types and their corresponding configuration defined by 'Support: Core' in this package, e.g. 'RequestHeaderModifier'. All implementations supporting GRPCRoute MUST support core filters.  - Extended: Filter types and their corresponding configuration defined by 'Support: Extended' in this package, e.g. 'RequestMirror'. Implementers are encouraged to support extended filters.  - Implementation-specific: Filters that are defined and supported by specific vendors. In the future, filters showing convergence in behavior across multiple implementations will be considered for inclusion in extended or core conformance levels. Filter-specific configuration for such filters is specified using the ExtensionRef field. 'Type' MUST be set to 'ExtensionRef' for custom filters.  Implementers are encouraged to define custom implementation types to extend the core API with implementation-specific behavior.  If a reference to a custom filter type cannot be resolved, the filter MUST NOT be skipped. Instead, requests that would have been processed by that filter MUST receive a HTTP error response.  ",
															MarkdownDescription: "Type identifies the type of filter to apply. As with other API fields, types are classified into three conformance levels:  - Core: Filter types and their corresponding configuration defined by 'Support: Core' in this package, e.g. 'RequestHeaderModifier'. All implementations supporting GRPCRoute MUST support core filters.  - Extended: Filter types and their corresponding configuration defined by 'Support: Extended' in this package, e.g. 'RequestMirror'. Implementers are encouraged to support extended filters.  - Implementation-specific: Filters that are defined and supported by specific vendors. In the future, filters showing convergence in behavior across multiple implementations will be considered for inclusion in extended or core conformance levels. Filter-specific configuration for such filters is specified using the ExtensionRef field. 'Type' MUST be set to 'ExtensionRef' for custom filters.  Implementers are encouraged to define custom implementation types to extend the core API with implementation-specific behavior.  If a reference to a custom filter type cannot be resolved, the filter MUST NOT be skipped. Instead, requests that would have been processed by that filter MUST receive a HTTP error response.  ",
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

											"group": schema.StringAttribute{
												Description:         "Group is the group of the referent. For example, 'gateway.networking.k8s.io'. When unspecified or empty string, core API group is inferred.",
												MarkdownDescription: "Group is the group of the referent. For example, 'gateway.networking.k8s.io'. When unspecified or empty string, core API group is inferred.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"kind": schema.StringAttribute{
												Description:         "Kind is the Kubernetes resource kind of the referent. For example 'Service'.  Defaults to 'Service' when not specified.  ExternalName services can refer to CNAME DNS records that may live outside of the cluster and as such are difficult to reason about in terms of conformance. They also may not be safe to forward to (see CVE-2021-25740 for more information). Implementations SHOULD NOT support ExternalName Services.  Support: Core (Services with a type other than ExternalName)  Support: Implementation-specific (Services with type ExternalName)",
												MarkdownDescription: "Kind is the Kubernetes resource kind of the referent. For example 'Service'.  Defaults to 'Service' when not specified.  ExternalName services can refer to CNAME DNS records that may live outside of the cluster and as such are difficult to reason about in terms of conformance. They also may not be safe to forward to (see CVE-2021-25740 for more information). Implementations SHOULD NOT support ExternalName Services.  Support: Core (Services with a type other than ExternalName)  Support: Implementation-specific (Services with type ExternalName)",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"name": schema.StringAttribute{
												Description:         "Name is the name of the referent.",
												MarkdownDescription: "Name is the name of the referent.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"namespace": schema.StringAttribute{
												Description:         "Namespace is the namespace of the backend. When unspecified, the local namespace is inferred.  Note that when a namespace different than the local namespace is specified, a ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details.  Support: Core",
												MarkdownDescription: "Namespace is the namespace of the backend. When unspecified, the local namespace is inferred.  Note that when a namespace different than the local namespace is specified, a ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details.  Support: Core",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"port": schema.Int64Attribute{
												Description:         "Port specifies the destination port number to use for this resource. Port is required when the referent is a Kubernetes Service. In this case, the port number is the service port number, not the target port. For other resources, destination port might be derived from the referent resource or this field.",
												MarkdownDescription: "Port specifies the destination port number to use for this resource. Port is required when the referent is a Kubernetes Service. In this case, the port number is the service port number, not the target port. For other resources, destination port might be derived from the referent resource or this field.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"weight": schema.Int64Attribute{
												Description:         "Weight specifies the proportion of requests forwarded to the referenced backend. This is computed as weight/(sum of all weights in this BackendRefs list). For non-zero values, there may be some epsilon from the exact proportion defined here depending on the precision an implementation supports. Weight is not a percentage and the sum of weights does not need to equal 100.  If only one backend is specified and it has a weight greater than 0, 100% of the traffic is forwarded to that backend. If weight is set to 0, no traffic should be forwarded for this entry. If unspecified, weight defaults to 1.  Support for this field varies based on the context where used.",
												MarkdownDescription: "Weight specifies the proportion of requests forwarded to the referenced backend. This is computed as weight/(sum of all weights in this BackendRefs list). For non-zero values, there may be some epsilon from the exact proportion defined here depending on the precision an implementation supports. Weight is not a percentage and the sum of weights does not need to equal 100.  If only one backend is specified and it has a weight greater than 0, 100% of the traffic is forwarded to that backend. If weight is set to 0, no traffic should be forwarded for this entry. If unspecified, weight defaults to 1.  Support for this field varies based on the context where used.",
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

								"filters": schema.ListNestedAttribute{
									Description:         "Filters define the filters that are applied to requests that match this rule.  The effects of ordering of multiple behaviors are currently unspecified. This can change in the future based on feedback during the alpha stage.  Conformance-levels at this level are defined based on the type of filter:  - ALL core filters MUST be supported by all implementations that support GRPCRoute. - Implementers are encouraged to support extended filters. - Implementation-specific custom filters have no API guarantees across implementations.  Specifying the same filter multiple times is not supported unless explicitly indicated in the filter.  If an implementation can not support a combination of filters, it must clearly document that limitation. In cases where incompatible or unsupported filters are specified and cause the 'Accepted' condition to be set to status 'False', implementations may use the 'IncompatibleFilters' reason to specify this configuration error.  Support: Core",
									MarkdownDescription: "Filters define the filters that are applied to requests that match this rule.  The effects of ordering of multiple behaviors are currently unspecified. This can change in the future based on feedback during the alpha stage.  Conformance-levels at this level are defined based on the type of filter:  - ALL core filters MUST be supported by all implementations that support GRPCRoute. - Implementers are encouraged to support extended filters. - Implementation-specific custom filters have no API guarantees across implementations.  Specifying the same filter multiple times is not supported unless explicitly indicated in the filter.  If an implementation can not support a combination of filters, it must clearly document that limitation. In cases where incompatible or unsupported filters are specified and cause the 'Accepted' condition to be set to status 'False', implementations may use the 'IncompatibleFilters' reason to specify this configuration error.  Support: Core",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"extension_ref": schema.SingleNestedAttribute{
												Description:         "ExtensionRef is an optional, implementation-specific extension to the 'filter' behavior.  For example, resource 'myroutefilter' in group 'networking.example.net'). ExtensionRef MUST NOT be used for core and extended filters.  Support: Implementation-specific  This filter can be used multiple times within the same rule.",
												MarkdownDescription: "ExtensionRef is an optional, implementation-specific extension to the 'filter' behavior.  For example, resource 'myroutefilter' in group 'networking.example.net'). ExtensionRef MUST NOT be used for core and extended filters.  Support: Implementation-specific  This filter can be used multiple times within the same rule.",
												Attributes: map[string]schema.Attribute{
													"group": schema.StringAttribute{
														Description:         "Group is the group of the referent. For example, 'gateway.networking.k8s.io'. When unspecified or empty string, core API group is inferred.",
														MarkdownDescription: "Group is the group of the referent. For example, 'gateway.networking.k8s.io'. When unspecified or empty string, core API group is inferred.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"kind": schema.StringAttribute{
														Description:         "Kind is kind of the referent. For example 'HTTPRoute' or 'Service'.",
														MarkdownDescription: "Kind is kind of the referent. For example 'HTTPRoute' or 'Service'.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "Name is the name of the referent.",
														MarkdownDescription: "Name is the name of the referent.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"request_header_modifier": schema.SingleNestedAttribute{
												Description:         "RequestHeaderModifier defines a schema for a filter that modifies request headers.  Support: Core",
												MarkdownDescription: "RequestHeaderModifier defines a schema for a filter that modifies request headers.  Support: Core",
												Attributes: map[string]schema.Attribute{
													"add": schema.ListNestedAttribute{
														Description:         "Add adds the given header(s) (name, value) to the request before the action. It appends to any existing values associated with the header name.  Input: GET /foo HTTP/1.1 my-header: foo  Config: add: - name: 'my-header' value: 'bar,baz'  Output: GET /foo HTTP/1.1 my-header: foo,bar,baz",
														MarkdownDescription: "Add adds the given header(s) (name, value) to the request before the action. It appends to any existing values associated with the header name.  Input: GET /foo HTTP/1.1 my-header: foo  Config: add: - name: 'my-header' value: 'bar,baz'  Output: GET /foo HTTP/1.1 my-header: foo,bar,baz",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is the name of the HTTP Header to be matched. Name matching MUST be case insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).  If multiple entries specify equivalent header names, the first entry with an equivalent name MUST be considered for a match. Subsequent entries with an equivalent header name MUST be ignored. Due to the case-insensitivity of header names, 'foo' and 'Foo' are considered equivalent.",
																	MarkdownDescription: "Name is the name of the HTTP Header to be matched. Name matching MUST be case insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).  If multiple entries specify equivalent header names, the first entry with an equivalent name MUST be considered for a match. Subsequent entries with an equivalent header name MUST be ignored. Due to the case-insensitivity of header names, 'foo' and 'Foo' are considered equivalent.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"value": schema.StringAttribute{
																	Description:         "Value is the value of HTTP Header to be matched.",
																	MarkdownDescription: "Value is the value of HTTP Header to be matched.",
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

													"remove": schema.ListAttribute{
														Description:         "Remove the given header(s) from the HTTP request before the action. The value of Remove is a list of HTTP header names. Note that the header names are case-insensitive (see https://datatracker.ietf.org/doc/html/rfc2616#section-4.2).  Input: GET /foo HTTP/1.1 my-header1: foo my-header2: bar my-header3: baz  Config: remove: ['my-header1', 'my-header3']  Output: GET /foo HTTP/1.1 my-header2: bar",
														MarkdownDescription: "Remove the given header(s) from the HTTP request before the action. The value of Remove is a list of HTTP header names. Note that the header names are case-insensitive (see https://datatracker.ietf.org/doc/html/rfc2616#section-4.2).  Input: GET /foo HTTP/1.1 my-header1: foo my-header2: bar my-header3: baz  Config: remove: ['my-header1', 'my-header3']  Output: GET /foo HTTP/1.1 my-header2: bar",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"set": schema.ListNestedAttribute{
														Description:         "Set overwrites the request with the given header (name, value) before the action.  Input: GET /foo HTTP/1.1 my-header: foo  Config: set: - name: 'my-header' value: 'bar'  Output: GET /foo HTTP/1.1 my-header: bar",
														MarkdownDescription: "Set overwrites the request with the given header (name, value) before the action.  Input: GET /foo HTTP/1.1 my-header: foo  Config: set: - name: 'my-header' value: 'bar'  Output: GET /foo HTTP/1.1 my-header: bar",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is the name of the HTTP Header to be matched. Name matching MUST be case insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).  If multiple entries specify equivalent header names, the first entry with an equivalent name MUST be considered for a match. Subsequent entries with an equivalent header name MUST be ignored. Due to the case-insensitivity of header names, 'foo' and 'Foo' are considered equivalent.",
																	MarkdownDescription: "Name is the name of the HTTP Header to be matched. Name matching MUST be case insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).  If multiple entries specify equivalent header names, the first entry with an equivalent name MUST be considered for a match. Subsequent entries with an equivalent header name MUST be ignored. Due to the case-insensitivity of header names, 'foo' and 'Foo' are considered equivalent.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"value": schema.StringAttribute{
																	Description:         "Value is the value of HTTP Header to be matched.",
																	MarkdownDescription: "Value is the value of HTTP Header to be matched.",
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
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"request_mirror": schema.SingleNestedAttribute{
												Description:         "RequestMirror defines a schema for a filter that mirrors requests. Requests are sent to the specified destination, but responses from that destination are ignored.  This filter can be used multiple times within the same rule. Note that not all implementations will be able to support mirroring to multiple backends.  Support: Extended",
												MarkdownDescription: "RequestMirror defines a schema for a filter that mirrors requests. Requests are sent to the specified destination, but responses from that destination are ignored.  This filter can be used multiple times within the same rule. Note that not all implementations will be able to support mirroring to multiple backends.  Support: Extended",
												Attributes: map[string]schema.Attribute{
													"backend_ref": schema.SingleNestedAttribute{
														Description:         "BackendRef references a resource where mirrored requests are sent.  Mirrored requests must be sent only to a single destination endpoint within this BackendRef, irrespective of how many endpoints are present within this BackendRef.  If the referent cannot be found, this BackendRef is invalid and must be dropped from the Gateway. The controller must ensure the 'ResolvedRefs' condition on the Route status is set to 'status: False' and not configure this backend in the underlying implementation.  If there is a cross-namespace reference to an *existing* object that is not allowed by a ReferenceGrant, the controller must ensure the 'ResolvedRefs'  condition on the Route is set to 'status: False', with the 'RefNotPermitted' reason and not configure this backend in the underlying implementation.  In either error case, the Message of the 'ResolvedRefs' Condition should be used to provide more detail about the problem.  Support: Extended for Kubernetes Service  Support: Implementation-specific for any other resource",
														MarkdownDescription: "BackendRef references a resource where mirrored requests are sent.  Mirrored requests must be sent only to a single destination endpoint within this BackendRef, irrespective of how many endpoints are present within this BackendRef.  If the referent cannot be found, this BackendRef is invalid and must be dropped from the Gateway. The controller must ensure the 'ResolvedRefs' condition on the Route status is set to 'status: False' and not configure this backend in the underlying implementation.  If there is a cross-namespace reference to an *existing* object that is not allowed by a ReferenceGrant, the controller must ensure the 'ResolvedRefs'  condition on the Route is set to 'status: False', with the 'RefNotPermitted' reason and not configure this backend in the underlying implementation.  In either error case, the Message of the 'ResolvedRefs' Condition should be used to provide more detail about the problem.  Support: Extended for Kubernetes Service  Support: Implementation-specific for any other resource",
														Attributes: map[string]schema.Attribute{
															"group": schema.StringAttribute{
																Description:         "Group is the group of the referent. For example, 'gateway.networking.k8s.io'. When unspecified or empty string, core API group is inferred.",
																MarkdownDescription: "Group is the group of the referent. For example, 'gateway.networking.k8s.io'. When unspecified or empty string, core API group is inferred.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"kind": schema.StringAttribute{
																Description:         "Kind is the Kubernetes resource kind of the referent. For example 'Service'.  Defaults to 'Service' when not specified.  ExternalName services can refer to CNAME DNS records that may live outside of the cluster and as such are difficult to reason about in terms of conformance. They also may not be safe to forward to (see CVE-2021-25740 for more information). Implementations SHOULD NOT support ExternalName Services.  Support: Core (Services with a type other than ExternalName)  Support: Implementation-specific (Services with type ExternalName)",
																MarkdownDescription: "Kind is the Kubernetes resource kind of the referent. For example 'Service'.  Defaults to 'Service' when not specified.  ExternalName services can refer to CNAME DNS records that may live outside of the cluster and as such are difficult to reason about in terms of conformance. They also may not be safe to forward to (see CVE-2021-25740 for more information). Implementations SHOULD NOT support ExternalName Services.  Support: Core (Services with a type other than ExternalName)  Support: Implementation-specific (Services with type ExternalName)",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "Name is the name of the referent.",
																MarkdownDescription: "Name is the name of the referent.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace is the namespace of the backend. When unspecified, the local namespace is inferred.  Note that when a namespace different than the local namespace is specified, a ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details.  Support: Core",
																MarkdownDescription: "Namespace is the namespace of the backend. When unspecified, the local namespace is inferred.  Note that when a namespace different than the local namespace is specified, a ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details.  Support: Core",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"port": schema.Int64Attribute{
																Description:         "Port specifies the destination port number to use for this resource. Port is required when the referent is a Kubernetes Service. In this case, the port number is the service port number, not the target port. For other resources, destination port might be derived from the referent resource or this field.",
																MarkdownDescription: "Port specifies the destination port number to use for this resource. Port is required when the referent is a Kubernetes Service. In this case, the port number is the service port number, not the target port. For other resources, destination port might be derived from the referent resource or this field.",
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

											"response_header_modifier": schema.SingleNestedAttribute{
												Description:         "ResponseHeaderModifier defines a schema for a filter that modifies response headers.  Support: Extended",
												MarkdownDescription: "ResponseHeaderModifier defines a schema for a filter that modifies response headers.  Support: Extended",
												Attributes: map[string]schema.Attribute{
													"add": schema.ListNestedAttribute{
														Description:         "Add adds the given header(s) (name, value) to the request before the action. It appends to any existing values associated with the header name.  Input: GET /foo HTTP/1.1 my-header: foo  Config: add: - name: 'my-header' value: 'bar,baz'  Output: GET /foo HTTP/1.1 my-header: foo,bar,baz",
														MarkdownDescription: "Add adds the given header(s) (name, value) to the request before the action. It appends to any existing values associated with the header name.  Input: GET /foo HTTP/1.1 my-header: foo  Config: add: - name: 'my-header' value: 'bar,baz'  Output: GET /foo HTTP/1.1 my-header: foo,bar,baz",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is the name of the HTTP Header to be matched. Name matching MUST be case insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).  If multiple entries specify equivalent header names, the first entry with an equivalent name MUST be considered for a match. Subsequent entries with an equivalent header name MUST be ignored. Due to the case-insensitivity of header names, 'foo' and 'Foo' are considered equivalent.",
																	MarkdownDescription: "Name is the name of the HTTP Header to be matched. Name matching MUST be case insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).  If multiple entries specify equivalent header names, the first entry with an equivalent name MUST be considered for a match. Subsequent entries with an equivalent header name MUST be ignored. Due to the case-insensitivity of header names, 'foo' and 'Foo' are considered equivalent.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"value": schema.StringAttribute{
																	Description:         "Value is the value of HTTP Header to be matched.",
																	MarkdownDescription: "Value is the value of HTTP Header to be matched.",
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

													"remove": schema.ListAttribute{
														Description:         "Remove the given header(s) from the HTTP request before the action. The value of Remove is a list of HTTP header names. Note that the header names are case-insensitive (see https://datatracker.ietf.org/doc/html/rfc2616#section-4.2).  Input: GET /foo HTTP/1.1 my-header1: foo my-header2: bar my-header3: baz  Config: remove: ['my-header1', 'my-header3']  Output: GET /foo HTTP/1.1 my-header2: bar",
														MarkdownDescription: "Remove the given header(s) from the HTTP request before the action. The value of Remove is a list of HTTP header names. Note that the header names are case-insensitive (see https://datatracker.ietf.org/doc/html/rfc2616#section-4.2).  Input: GET /foo HTTP/1.1 my-header1: foo my-header2: bar my-header3: baz  Config: remove: ['my-header1', 'my-header3']  Output: GET /foo HTTP/1.1 my-header2: bar",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"set": schema.ListNestedAttribute{
														Description:         "Set overwrites the request with the given header (name, value) before the action.  Input: GET /foo HTTP/1.1 my-header: foo  Config: set: - name: 'my-header' value: 'bar'  Output: GET /foo HTTP/1.1 my-header: bar",
														MarkdownDescription: "Set overwrites the request with the given header (name, value) before the action.  Input: GET /foo HTTP/1.1 my-header: foo  Config: set: - name: 'my-header' value: 'bar'  Output: GET /foo HTTP/1.1 my-header: bar",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is the name of the HTTP Header to be matched. Name matching MUST be case insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).  If multiple entries specify equivalent header names, the first entry with an equivalent name MUST be considered for a match. Subsequent entries with an equivalent header name MUST be ignored. Due to the case-insensitivity of header names, 'foo' and 'Foo' are considered equivalent.",
																	MarkdownDescription: "Name is the name of the HTTP Header to be matched. Name matching MUST be case insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).  If multiple entries specify equivalent header names, the first entry with an equivalent name MUST be considered for a match. Subsequent entries with an equivalent header name MUST be ignored. Due to the case-insensitivity of header names, 'foo' and 'Foo' are considered equivalent.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"value": schema.StringAttribute{
																	Description:         "Value is the value of HTTP Header to be matched.",
																	MarkdownDescription: "Value is the value of HTTP Header to be matched.",
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
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"type": schema.StringAttribute{
												Description:         "Type identifies the type of filter to apply. As with other API fields, types are classified into three conformance levels:  - Core: Filter types and their corresponding configuration defined by 'Support: Core' in this package, e.g. 'RequestHeaderModifier'. All implementations supporting GRPCRoute MUST support core filters.  - Extended: Filter types and their corresponding configuration defined by 'Support: Extended' in this package, e.g. 'RequestMirror'. Implementers are encouraged to support extended filters.  - Implementation-specific: Filters that are defined and supported by specific vendors. In the future, filters showing convergence in behavior across multiple implementations will be considered for inclusion in extended or core conformance levels. Filter-specific configuration for such filters is specified using the ExtensionRef field. 'Type' MUST be set to 'ExtensionRef' for custom filters.  Implementers are encouraged to define custom implementation types to extend the core API with implementation-specific behavior.  If a reference to a custom filter type cannot be resolved, the filter MUST NOT be skipped. Instead, requests that would have been processed by that filter MUST receive a HTTP error response.  ",
												MarkdownDescription: "Type identifies the type of filter to apply. As with other API fields, types are classified into three conformance levels:  - Core: Filter types and their corresponding configuration defined by 'Support: Core' in this package, e.g. 'RequestHeaderModifier'. All implementations supporting GRPCRoute MUST support core filters.  - Extended: Filter types and their corresponding configuration defined by 'Support: Extended' in this package, e.g. 'RequestMirror'. Implementers are encouraged to support extended filters.  - Implementation-specific: Filters that are defined and supported by specific vendors. In the future, filters showing convergence in behavior across multiple implementations will be considered for inclusion in extended or core conformance levels. Filter-specific configuration for such filters is specified using the ExtensionRef field. 'Type' MUST be set to 'ExtensionRef' for custom filters.  Implementers are encouraged to define custom implementation types to extend the core API with implementation-specific behavior.  If a reference to a custom filter type cannot be resolved, the filter MUST NOT be skipped. Instead, requests that would have been processed by that filter MUST receive a HTTP error response.  ",
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

								"matches": schema.ListNestedAttribute{
									Description:         "Matches define conditions used for matching the rule against incoming gRPC requests. Each match is independent, i.e. this rule will be matched if **any** one of the matches is satisfied.  For example, take the following matches configuration:  ''' matches: - method: service: foo.bar headers: values: version: 2 - method: service: foo.bar.v2 '''  For a request to match against this rule, it MUST satisfy EITHER of the two conditions:  - service of foo.bar AND contains the header 'version: 2' - service of foo.bar.v2  See the documentation for GRPCRouteMatch on how to specify multiple match conditions to be ANDed together.  If no matches are specified, the implementation MUST match every gRPC request.  Proxy or Load Balancer routing configuration generated from GRPCRoutes MUST prioritize rules based on the following criteria, continuing on ties. Merging MUST not be done between GRPCRoutes and HTTPRoutes. Precedence MUST be given to the rule with the largest number of:  * Characters in a matching non-wildcard hostname. * Characters in a matching hostname. * Characters in a matching service. * Characters in a matching method. * Header matches.  If ties still exist across multiple Routes, matching precedence MUST be determined in order of the following criteria, continuing on ties:  * The oldest Route based on creation timestamp. * The Route appearing first in alphabetical order by '{namespace}/{name}'.  If ties still exist within the Route that has been given precedence, matching precedence MUST be granted to the first matching rule meeting the above criteria.",
									MarkdownDescription: "Matches define conditions used for matching the rule against incoming gRPC requests. Each match is independent, i.e. this rule will be matched if **any** one of the matches is satisfied.  For example, take the following matches configuration:  ''' matches: - method: service: foo.bar headers: values: version: 2 - method: service: foo.bar.v2 '''  For a request to match against this rule, it MUST satisfy EITHER of the two conditions:  - service of foo.bar AND contains the header 'version: 2' - service of foo.bar.v2  See the documentation for GRPCRouteMatch on how to specify multiple match conditions to be ANDed together.  If no matches are specified, the implementation MUST match every gRPC request.  Proxy or Load Balancer routing configuration generated from GRPCRoutes MUST prioritize rules based on the following criteria, continuing on ties. Merging MUST not be done between GRPCRoutes and HTTPRoutes. Precedence MUST be given to the rule with the largest number of:  * Characters in a matching non-wildcard hostname. * Characters in a matching hostname. * Characters in a matching service. * Characters in a matching method. * Header matches.  If ties still exist across multiple Routes, matching precedence MUST be determined in order of the following criteria, continuing on ties:  * The oldest Route based on creation timestamp. * The Route appearing first in alphabetical order by '{namespace}/{name}'.  If ties still exist within the Route that has been given precedence, matching precedence MUST be granted to the first matching rule meeting the above criteria.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"headers": schema.ListNestedAttribute{
												Description:         "Headers specifies gRPC request header matchers. Multiple match values are ANDed together, meaning, a request MUST match all the specified headers to select the route.",
												MarkdownDescription: "Headers specifies gRPC request header matchers. Multiple match values are ANDed together, meaning, a request MUST match all the specified headers to select the route.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name is the name of the gRPC Header to be matched.  If multiple entries specify equivalent header names, only the first entry with an equivalent name MUST be considered for a match. Subsequent entries with an equivalent header name MUST be ignored. Due to the case-insensitivity of header names, 'foo' and 'Foo' are considered equivalent.",
															MarkdownDescription: "Name is the name of the gRPC Header to be matched.  If multiple entries specify equivalent header names, only the first entry with an equivalent name MUST be considered for a match. Subsequent entries with an equivalent header name MUST be ignored. Due to the case-insensitivity of header names, 'foo' and 'Foo' are considered equivalent.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"type": schema.StringAttribute{
															Description:         "Type specifies how to match against the value of the header.",
															MarkdownDescription: "Type specifies how to match against the value of the header.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"value": schema.StringAttribute{
															Description:         "Value is the value of the gRPC Header to be matched.",
															MarkdownDescription: "Value is the value of the gRPC Header to be matched.",
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

											"method": schema.SingleNestedAttribute{
												Description:         "Method specifies a gRPC request service/method matcher. If this field is not specified, all services and methods will match.",
												MarkdownDescription: "Method specifies a gRPC request service/method matcher. If this field is not specified, all services and methods will match.",
												Attributes: map[string]schema.Attribute{
													"method": schema.StringAttribute{
														Description:         "Value of the method to match against. If left empty or omitted, will match all services.  At least one of Service and Method MUST be a non-empty string.",
														MarkdownDescription: "Value of the method to match against. If left empty or omitted, will match all services.  At least one of Service and Method MUST be a non-empty string.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"service": schema.StringAttribute{
														Description:         "Value of the service to match against. If left empty or omitted, will match any service.  At least one of Service and Method MUST be a non-empty string.",
														MarkdownDescription: "Value of the service to match against. If left empty or omitted, will match any service.  At least one of Service and Method MUST be a non-empty string.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"type": schema.StringAttribute{
														Description:         "Type specifies how to match against the service and/or method. Support: Core (Exact with service and method specified)  Support: Implementation-specific (Exact with method specified but no service specified)  Support: Implementation-specific (RegularExpression)",
														MarkdownDescription: "Type specifies how to match against the service and/or method. Support: Core (Exact with service and method specified)  Support: Implementation-specific (Exact with method specified but no service specified)  Support: Implementation-specific (RegularExpression)",
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
									},
									Required: false,
									Optional: false,
									Computed: true,
								},
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

func (r *GatewayNetworkingK8SIoGrpcrouteV1Alpha2DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *GatewayNetworkingK8SIoGrpcrouteV1Alpha2DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_gateway_networking_k8s_io_grpc_route_v1alpha2")

	var data GatewayNetworkingK8SIoGrpcrouteV1Alpha2DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "gateway.networking.k8s.io", Version: "v1alpha2", Resource: "GRPCRoute"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
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

	var readResponse GatewayNetworkingK8SIoGrpcrouteV1Alpha2DataSourceData
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
	data.ApiVersion = pointer.String("gateway.networking.k8s.io/v1alpha2")
	data.Kind = pointer.String("GRPCRoute")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
