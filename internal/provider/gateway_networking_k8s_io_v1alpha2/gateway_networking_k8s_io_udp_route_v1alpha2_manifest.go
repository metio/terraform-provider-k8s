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
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &GatewayNetworkingK8SIoUdprouteV1Alpha2Manifest{}
)

func NewGatewayNetworkingK8SIoUdprouteV1Alpha2Manifest() datasource.DataSource {
	return &GatewayNetworkingK8SIoUdprouteV1Alpha2Manifest{}
}

type GatewayNetworkingK8SIoUdprouteV1Alpha2Manifest struct{}

type GatewayNetworkingK8SIoUdprouteV1Alpha2ManifestData struct {
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
				Group     *string `tfsdk:"group" json:"group,omitempty"`
				Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				Port      *int64  `tfsdk:"port" json:"port,omitempty"`
				Weight    *int64  `tfsdk:"weight" json:"weight,omitempty"`
			} `tfsdk:"backend_refs" json:"backendRefs,omitempty"`
		} `tfsdk:"rules" json:"rules,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *GatewayNetworkingK8SIoUdprouteV1Alpha2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_gateway_networking_k8s_io_udp_route_v1alpha2_manifest"
}

func (r *GatewayNetworkingK8SIoUdprouteV1Alpha2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "UDPRoute provides a way to route UDP traffic. When combined with a Gatewaylistener, it can be used to forward traffic on the port specified by thelistener to a set of backends specified by the UDPRoute.",
		MarkdownDescription: "UDPRoute provides a way to route UDP traffic. When combined with a Gatewaylistener, it can be used to forward traffic on the port specified by thelistener to a set of backends specified by the UDPRoute.",
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
				Description:         "Spec defines the desired state of UDPRoute.",
				MarkdownDescription: "Spec defines the desired state of UDPRoute.",
				Attributes: map[string]schema.Attribute{
					"parent_refs": schema.ListNestedAttribute{
						Description:         "ParentRefs references the resources (usually Gateways) that a Route wantsto be attached to. Note that the referenced parent resource needs toallow this for the attachment to be complete. For Gateways, that meansthe Gateway needs to allow attachment from Routes of this kind andnamespace. For Services, that means the Service must either be in the samenamespace for a 'producer' route, or the mesh implementation must supportand allow 'consumer' routes for the referenced Service. ReferenceGrant isnot applicable for governing ParentRefs to Services - it is not possible tocreate a 'producer' route for a Service in a different namespace from theRoute.There are two kinds of parent resources with 'Core' support:* Gateway (Gateway conformance profile)* Service (Mesh conformance profile, experimental, ClusterIP Services only)This API may be extended in the future to support additional kinds of parentresources.ParentRefs must be _distinct_. This means either that:* They select different objects.  If this is the case, then parentRef  entries are distinct. In terms of fields, this means that the  multi-part key defined by 'group', 'kind', 'namespace', and 'name' must  be unique across all parentRef entries in the Route.* They do not select different objects, but for each optional field used,  each ParentRef that selects the same object must set the same set of  optional fields to different values. If one ParentRef sets a  combination of optional fields, all must set the same combination.Some examples:* If one ParentRef sets 'sectionName', all ParentRefs referencing the  same object must also set 'sectionName'.* If one ParentRef sets 'port', all ParentRefs referencing the same  object must also set 'port'.* If one ParentRef sets 'sectionName' and 'port', all ParentRefs  referencing the same object must also set 'sectionName' and 'port'.It is possible to separately reference multiple distinct objects that maybe collapsed by an implementation. For example, some implementations maychoose to merge compatible Gateway Listeners together. If that is thecase, the list of routes attached to those resources should also bemerged.Note that for ParentRefs that cross namespace boundaries, there are specificrules. Cross-namespace references are only valid if they are explicitlyallowed by something in the namespace they are referring to. For example,Gateway has the AllowedRoutes field, and ReferenceGrant provides ageneric way to enable other kinds of cross-namespace reference.ParentRefs from a Route to a Service in the same namespace are 'producer'routes, which apply default routing rules to inbound connections fromany namespace to the Service.ParentRefs from a Route to a Service in a different namespace are'consumer' routes, and these routing rules are only applied to outboundconnections originating from the same namespace as the Route, for whichthe intended destination of the connections are a Service targeted as aParentRef of the Route.",
						MarkdownDescription: "ParentRefs references the resources (usually Gateways) that a Route wantsto be attached to. Note that the referenced parent resource needs toallow this for the attachment to be complete. For Gateways, that meansthe Gateway needs to allow attachment from Routes of this kind andnamespace. For Services, that means the Service must either be in the samenamespace for a 'producer' route, or the mesh implementation must supportand allow 'consumer' routes for the referenced Service. ReferenceGrant isnot applicable for governing ParentRefs to Services - it is not possible tocreate a 'producer' route for a Service in a different namespace from theRoute.There are two kinds of parent resources with 'Core' support:* Gateway (Gateway conformance profile)* Service (Mesh conformance profile, experimental, ClusterIP Services only)This API may be extended in the future to support additional kinds of parentresources.ParentRefs must be _distinct_. This means either that:* They select different objects.  If this is the case, then parentRef  entries are distinct. In terms of fields, this means that the  multi-part key defined by 'group', 'kind', 'namespace', and 'name' must  be unique across all parentRef entries in the Route.* They do not select different objects, but for each optional field used,  each ParentRef that selects the same object must set the same set of  optional fields to different values. If one ParentRef sets a  combination of optional fields, all must set the same combination.Some examples:* If one ParentRef sets 'sectionName', all ParentRefs referencing the  same object must also set 'sectionName'.* If one ParentRef sets 'port', all ParentRefs referencing the same  object must also set 'port'.* If one ParentRef sets 'sectionName' and 'port', all ParentRefs  referencing the same object must also set 'sectionName' and 'port'.It is possible to separately reference multiple distinct objects that maybe collapsed by an implementation. For example, some implementations maychoose to merge compatible Gateway Listeners together. If that is thecase, the list of routes attached to those resources should also bemerged.Note that for ParentRefs that cross namespace boundaries, there are specificrules. Cross-namespace references are only valid if they are explicitlyallowed by something in the namespace they are referring to. For example,Gateway has the AllowedRoutes field, and ReferenceGrant provides ageneric way to enable other kinds of cross-namespace reference.ParentRefs from a Route to a Service in the same namespace are 'producer'routes, which apply default routing rules to inbound connections fromany namespace to the Service.ParentRefs from a Route to a Service in a different namespace are'consumer' routes, and these routing rules are only applied to outboundconnections originating from the same namespace as the Route, for whichthe intended destination of the connections are a Service targeted as aParentRef of the Route.",
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
									Description:         "Kind is kind of the referent.There are two kinds of parent resources with 'Core' support:* Gateway (Gateway conformance profile)* Service (Mesh conformance profile, experimental, ClusterIP Services only)Support for other resources is Implementation-Specific.",
									MarkdownDescription: "Kind is kind of the referent.There are two kinds of parent resources with 'Core' support:* Gateway (Gateway conformance profile)* Service (Mesh conformance profile, experimental, ClusterIP Services only)Support for other resources is Implementation-Specific.",
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
									Description:         "SectionName is the name of a section within the target resource. In thefollowing resources, SectionName is interpreted as the following:* Gateway: Listener Name. When both Port (experimental) and SectionNameare specified, the name and port of the selected listener must matchboth specified values.* Service: Port Name. When both Port (experimental) and SectionNameare specified, the name and port of the selected listener must matchboth specified values. Note that attaching Routes to Services as Parentsis part of experimental Mesh support and is not supported for any otherpurpose.Implementations MAY choose to support attaching Routes to other resources.If that is the case, they MUST clearly document how SectionName isinterpreted.When unspecified (empty string), this will reference the entire resource.For the purpose of status, an attachment is considered successful if atleast one section in the parent resource accepts it. For example, Gatewaylisteners can restrict which Routes can attach to them by Route kind,namespace, or hostname. If 1 of 2 Gateway listeners accept attachment fromthe referencing Route, the Route MUST be considered successfullyattached. If no Gateway listeners accept attachment from this Route, theRoute MUST be considered detached from the Gateway.Support: Core",
									MarkdownDescription: "SectionName is the name of a section within the target resource. In thefollowing resources, SectionName is interpreted as the following:* Gateway: Listener Name. When both Port (experimental) and SectionNameare specified, the name and port of the selected listener must matchboth specified values.* Service: Port Name. When both Port (experimental) and SectionNameare specified, the name and port of the selected listener must matchboth specified values. Note that attaching Routes to Services as Parentsis part of experimental Mesh support and is not supported for any otherpurpose.Implementations MAY choose to support attaching Routes to other resources.If that is the case, they MUST clearly document how SectionName isinterpreted.When unspecified (empty string), this will reference the entire resource.For the purpose of status, an attachment is considered successful if atleast one section in the parent resource accepts it. For example, Gatewaylisteners can restrict which Routes can attach to them by Route kind,namespace, or hostname. If 1 of 2 Gateway listeners accept attachment fromthe referencing Route, the Route MUST be considered successfullyattached. If no Gateway listeners accept attachment from this Route, theRoute MUST be considered detached from the Gateway.Support: Core",
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
						Description:         "Rules are a list of UDP matchers and actions.",
						MarkdownDescription: "Rules are a list of UDP matchers and actions.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"backend_refs": schema.ListNestedAttribute{
									Description:         "BackendRefs defines the backend(s) where matching requests should besent. If unspecified or invalid (refers to a non-existent resource or aService with no endpoints), the underlying implementation MUST activelyreject connection attempts to this backend. Packet drops mustrespect weight; if an invalid backend is requested to have 80% ofthe packets, then 80% of packets must be dropped instead.Support: Core for Kubernetes ServiceSupport: Extended for Kubernetes ServiceImportSupport: Implementation-specific for any other resourceSupport for weight: Extended",
									MarkdownDescription: "BackendRefs defines the backend(s) where matching requests should besent. If unspecified or invalid (refers to a non-existent resource or aService with no endpoints), the underlying implementation MUST activelyreject connection attempts to this backend. Packet drops mustrespect weight; if an invalid backend is requested to have 80% ofthe packets, then 80% of packets must be dropped instead.Support: Core for Kubernetes ServiceSupport: Extended for Kubernetes ServiceImportSupport: Implementation-specific for any other resourceSupport for weight: Extended",
									NestedObject: schema.NestedAttributeObject{
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
							},
						},
						Required: true,
						Optional: false,
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

func (r *GatewayNetworkingK8SIoUdprouteV1Alpha2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_gateway_networking_k8s_io_udp_route_v1alpha2_manifest")

	var model GatewayNetworkingK8SIoUdprouteV1Alpha2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("gateway.networking.k8s.io/v1alpha2")
	model.Kind = pointer.String("UDPRoute")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
