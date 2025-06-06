/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package gateway_networking_x_k8s_io_v1alpha1

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
	_ datasource.DataSource = &GatewayNetworkingXK8SIoXlistenerSetV1Alpha1Manifest{}
)

func NewGatewayNetworkingXK8SIoXlistenerSetV1Alpha1Manifest() datasource.DataSource {
	return &GatewayNetworkingXK8SIoXlistenerSetV1Alpha1Manifest{}
}

type GatewayNetworkingXK8SIoXlistenerSetV1Alpha1Manifest struct{}

type GatewayNetworkingXK8SIoXlistenerSetV1Alpha1ManifestData struct {
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
		Listeners *[]struct {
			AllowedRoutes *struct {
				Kinds *[]struct {
					Group *string `tfsdk:"group" json:"group,omitempty"`
					Kind  *string `tfsdk:"kind" json:"kind,omitempty"`
				} `tfsdk:"kinds" json:"kinds,omitempty"`
				Namespaces *struct {
					From     *string `tfsdk:"from" json:"from,omitempty"`
					Selector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"selector" json:"selector,omitempty"`
				} `tfsdk:"namespaces" json:"namespaces,omitempty"`
			} `tfsdk:"allowed_routes" json:"allowedRoutes,omitempty"`
			Hostname *string `tfsdk:"hostname" json:"hostname,omitempty"`
			Name     *string `tfsdk:"name" json:"name,omitempty"`
			Port     *int64  `tfsdk:"port" json:"port,omitempty"`
			Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
			Tls      *struct {
				CertificateRefs *[]struct {
					Group     *string `tfsdk:"group" json:"group,omitempty"`
					Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"certificate_refs" json:"certificateRefs,omitempty"`
				FrontendValidation *struct {
					CaCertificateRefs *[]struct {
						Group     *string `tfsdk:"group" json:"group,omitempty"`
						Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"ca_certificate_refs" json:"caCertificateRefs,omitempty"`
				} `tfsdk:"frontend_validation" json:"frontendValidation,omitempty"`
				Mode    *string            `tfsdk:"mode" json:"mode,omitempty"`
				Options *map[string]string `tfsdk:"options" json:"options,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
		} `tfsdk:"listeners" json:"listeners,omitempty"`
		ParentRef *struct {
			Group     *string `tfsdk:"group" json:"group,omitempty"`
			Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"parent_ref" json:"parentRef,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *GatewayNetworkingXK8SIoXlistenerSetV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_gateway_networking_x_k8s_io_x_listener_set_v1alpha1_manifest"
}

func (r *GatewayNetworkingXK8SIoXlistenerSetV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "XListenerSet defines a set of additional listeners to attach to an existing Gateway.",
		MarkdownDescription: "XListenerSet defines a set of additional listeners to attach to an existing Gateway.",
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
				Description:         "Spec defines the desired state of ListenerSet.",
				MarkdownDescription: "Spec defines the desired state of ListenerSet.",
				Attributes: map[string]schema.Attribute{
					"listeners": schema.ListNestedAttribute{
						Description:         "Listeners associated with this ListenerSet. Listeners define logical endpoints that are bound on this referenced parent Gateway's addresses. Listeners in a 'Gateway' and their attached 'ListenerSets' are concatenated as a list when programming the underlying infrastructure. Each listener name does not need to be unique across the Gateway and ListenerSets. See ListenerEntry.Name for more details. Implementations MUST treat the parent Gateway as having the merged list of all listeners from itself and attached ListenerSets using the following precedence: 1. 'parent' Gateway 2. ListenerSet ordered by creation time (oldest first) 3. ListenerSet ordered alphabetically by “{namespace}/{name}”. An implementation MAY reject listeners by setting the ListenerEntryStatus 'Accepted'' condition to False with the Reason 'TooManyListeners' If a listener has a conflict, this will be reported in the Status.ListenerEntryStatus setting the 'Conflicted' condition to True. Implementations SHOULD be cautious about what information from the parent or siblings are reported to avoid accidentally leaking sensitive information that the child would not otherwise have access to. This can include contents of secrets etc.",
						MarkdownDescription: "Listeners associated with this ListenerSet. Listeners define logical endpoints that are bound on this referenced parent Gateway's addresses. Listeners in a 'Gateway' and their attached 'ListenerSets' are concatenated as a list when programming the underlying infrastructure. Each listener name does not need to be unique across the Gateway and ListenerSets. See ListenerEntry.Name for more details. Implementations MUST treat the parent Gateway as having the merged list of all listeners from itself and attached ListenerSets using the following precedence: 1. 'parent' Gateway 2. ListenerSet ordered by creation time (oldest first) 3. ListenerSet ordered alphabetically by “{namespace}/{name}”. An implementation MAY reject listeners by setting the ListenerEntryStatus 'Accepted'' condition to False with the Reason 'TooManyListeners' If a listener has a conflict, this will be reported in the Status.ListenerEntryStatus setting the 'Conflicted' condition to True. Implementations SHOULD be cautious about what information from the parent or siblings are reported to avoid accidentally leaking sensitive information that the child would not otherwise have access to. This can include contents of secrets etc.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"allowed_routes": schema.SingleNestedAttribute{
									Description:         "AllowedRoutes defines the types of routes that MAY be attached to a Listener and the trusted namespaces where those Route resources MAY be present. Although a client request may match multiple route rules, only one rule may ultimately receive the request. Matching precedence MUST be determined in order of the following criteria: * The most specific match as defined by the Route type. * The oldest Route based on creation timestamp. For example, a Route with a creation timestamp of '2020-09-08 01:02:03' is given precedence over a Route with a creation timestamp of '2020-09-08 01:02:04'. * If everything else is equivalent, the Route appearing first in alphabetical order (namespace/name) should be given precedence. For example, foo/bar is given precedence over foo/baz. All valid rules within a Route attached to this Listener should be implemented. Invalid Route rules can be ignored (sometimes that will mean the full Route). If a Route rule transitions from valid to invalid, support for that Route rule should be dropped to ensure consistency. For example, even if a filter specified by a Route rule is invalid, the rest of the rules within that Route should still be supported.",
									MarkdownDescription: "AllowedRoutes defines the types of routes that MAY be attached to a Listener and the trusted namespaces where those Route resources MAY be present. Although a client request may match multiple route rules, only one rule may ultimately receive the request. Matching precedence MUST be determined in order of the following criteria: * The most specific match as defined by the Route type. * The oldest Route based on creation timestamp. For example, a Route with a creation timestamp of '2020-09-08 01:02:03' is given precedence over a Route with a creation timestamp of '2020-09-08 01:02:04'. * If everything else is equivalent, the Route appearing first in alphabetical order (namespace/name) should be given precedence. For example, foo/bar is given precedence over foo/baz. All valid rules within a Route attached to this Listener should be implemented. Invalid Route rules can be ignored (sometimes that will mean the full Route). If a Route rule transitions from valid to invalid, support for that Route rule should be dropped to ensure consistency. For example, even if a filter specified by a Route rule is invalid, the rest of the rules within that Route should still be supported.",
									Attributes: map[string]schema.Attribute{
										"kinds": schema.ListNestedAttribute{
											Description:         "Kinds specifies the groups and kinds of Routes that are allowed to bind to this Gateway Listener. When unspecified or empty, the kinds of Routes selected are determined using the Listener protocol. A RouteGroupKind MUST correspond to kinds of Routes that are compatible with the application protocol specified in the Listener's Protocol field. If an implementation does not support or recognize this resource type, it MUST set the 'ResolvedRefs' condition to False for this Listener with the 'InvalidRouteKinds' reason. Support: Core",
											MarkdownDescription: "Kinds specifies the groups and kinds of Routes that are allowed to bind to this Gateway Listener. When unspecified or empty, the kinds of Routes selected are determined using the Listener protocol. A RouteGroupKind MUST correspond to kinds of Routes that are compatible with the application protocol specified in the Listener's Protocol field. If an implementation does not support or recognize this resource type, it MUST set the 'ResolvedRefs' condition to False for this Listener with the 'InvalidRouteKinds' reason. Support: Core",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"group": schema.StringAttribute{
														Description:         "Group is the group of the Route.",
														MarkdownDescription: "Group is the group of the Route.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtMost(253),
															stringvalidator.RegexMatches(regexp.MustCompile(`^$|^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
														},
													},

													"kind": schema.StringAttribute{
														Description:         "Kind is the kind of the Route.",
														MarkdownDescription: "Kind is the kind of the Route.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.LengthAtMost(63),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z]([-a-zA-Z0-9]*[a-zA-Z0-9])?$`), ""),
														},
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"namespaces": schema.SingleNestedAttribute{
											Description:         "Namespaces indicates namespaces from which Routes may be attached to this Listener. This is restricted to the namespace of this Gateway by default. Support: Core",
											MarkdownDescription: "Namespaces indicates namespaces from which Routes may be attached to this Listener. This is restricted to the namespace of this Gateway by default. Support: Core",
											Attributes: map[string]schema.Attribute{
												"from": schema.StringAttribute{
													Description:         "From indicates where Routes will be selected for this Gateway. Possible values are: * All: Routes in all namespaces may be used by this Gateway. * Selector: Routes in namespaces selected by the selector may be used by this Gateway. * Same: Only Routes in the same namespace may be used by this Gateway. Support: Core",
													MarkdownDescription: "From indicates where Routes will be selected for this Gateway. Possible values are: * All: Routes in all namespaces may be used by this Gateway. * Selector: Routes in namespaces selected by the selector may be used by this Gateway. * Same: Only Routes in the same namespace may be used by this Gateway. Support: Core",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("All", "Selector", "Same"),
													},
												},

												"selector": schema.SingleNestedAttribute{
													Description:         "Selector must be specified when From is set to 'Selector'. In that case, only Routes in Namespaces matching this Selector will be selected by this Gateway. This field is ignored for other values of 'From'. Support: Core",
													MarkdownDescription: "Selector must be specified when From is set to 'Selector'. In that case, only Routes in Namespaces matching this Selector will be selected by this Gateway. This field is ignored for other values of 'From'. Support: Core",
													Attributes: map[string]schema.Attribute{
														"match_expressions": schema.ListNestedAttribute{
															Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "key is the label key that the selector applies to.",
																		MarkdownDescription: "key is the label key that the selector applies to.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"values": schema.ListAttribute{
																		Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		ElementType:         types.StringType,
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

														"match_labels": schema.MapAttribute{
															Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
															MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

								"hostname": schema.StringAttribute{
									Description:         "Hostname specifies the virtual hostname to match for protocol types that define this concept. When unspecified, all hostnames are matched. This field is ignored for protocols that don't require hostname based matching. Implementations MUST apply Hostname matching appropriately for each of the following protocols: * TLS: The Listener Hostname MUST match the SNI. * HTTP: The Listener Hostname MUST match the Host header of the request. * HTTPS: The Listener Hostname SHOULD match at both the TLS and HTTP protocol layers as described above. If an implementation does not ensure that both the SNI and Host header match the Listener hostname, it MUST clearly document that. For HTTPRoute and TLSRoute resources, there is an interaction with the 'spec.hostnames' array. When both listener and route specify hostnames, there MUST be an intersection between the values for a Route to be accepted. For more information, refer to the Route specific Hostnames documentation. Hostnames that are prefixed with a wildcard label ('*.') are interpreted as a suffix match. That means that a match for '*.example.com' would match both 'test.example.com', and 'foo.test.example.com', but not 'example.com'.",
									MarkdownDescription: "Hostname specifies the virtual hostname to match for protocol types that define this concept. When unspecified, all hostnames are matched. This field is ignored for protocols that don't require hostname based matching. Implementations MUST apply Hostname matching appropriately for each of the following protocols: * TLS: The Listener Hostname MUST match the SNI. * HTTP: The Listener Hostname MUST match the Host header of the request. * HTTPS: The Listener Hostname SHOULD match at both the TLS and HTTP protocol layers as described above. If an implementation does not ensure that both the SNI and Host header match the Listener hostname, it MUST clearly document that. For HTTPRoute and TLSRoute resources, there is an interaction with the 'spec.hostnames' array. When both listener and route specify hostnames, there MUST be an intersection between the values for a Route to be accepted. For more information, refer to the Route specific Hostnames documentation. Hostnames that are prefixed with a wildcard label ('*.') are interpreted as a suffix match. That means that a match for '*.example.com' would match both 'test.example.com', and 'foo.test.example.com', but not 'example.com'.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(253),
										stringvalidator.RegexMatches(regexp.MustCompile(`^(\*\.)?[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
									},
								},

								"name": schema.StringAttribute{
									Description:         "Name is the name of the Listener. This name MUST be unique within a ListenerSet. Name is not required to be unique across a Gateway and ListenerSets. Routes can attach to a Listener by having a ListenerSet as a parentRef and setting the SectionName",
									MarkdownDescription: "Name is the name of the Listener. This name MUST be unique within a ListenerSet. Name is not required to be unique across a Gateway and ListenerSets. Routes can attach to a Listener by having a ListenerSet as a parentRef and setting the SectionName",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(253),
										stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
									},
								},

								"port": schema.Int64Attribute{
									Description:         "Port is the network port. Multiple listeners may use the same port, subject to the Listener compatibility rules.",
									MarkdownDescription: "Port is the network port. Multiple listeners may use the same port, subject to the Listener compatibility rules.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.Int64{
										int64validator.AtLeast(1),
										int64validator.AtMost(65535),
									},
								},

								"protocol": schema.StringAttribute{
									Description:         "Protocol specifies the network protocol this listener expects to receive.",
									MarkdownDescription: "Protocol specifies the network protocol this listener expects to receive.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(255),
										stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9]([-a-zA-Z0-9]*[a-zA-Z0-9])?$|[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*\/[A-Za-z0-9]+$`), ""),
									},
								},

								"tls": schema.SingleNestedAttribute{
									Description:         "TLS is the TLS configuration for the Listener. This field is required if the Protocol field is 'HTTPS' or 'TLS'. It is invalid to set this field if the Protocol field is 'HTTP', 'TCP', or 'UDP'. The association of SNIs to Certificate defined in GatewayTLSConfig is defined based on the Hostname field for this listener. The GatewayClass MUST use the longest matching SNI out of all available certificates for any TLS handshake.",
									MarkdownDescription: "TLS is the TLS configuration for the Listener. This field is required if the Protocol field is 'HTTPS' or 'TLS'. It is invalid to set this field if the Protocol field is 'HTTP', 'TCP', or 'UDP'. The association of SNIs to Certificate defined in GatewayTLSConfig is defined based on the Hostname field for this listener. The GatewayClass MUST use the longest matching SNI out of all available certificates for any TLS handshake.",
									Attributes: map[string]schema.Attribute{
										"certificate_refs": schema.ListNestedAttribute{
											Description:         "CertificateRefs contains a series of references to Kubernetes objects that contains TLS certificates and private keys. These certificates are used to establish a TLS handshake for requests that match the hostname of the associated listener. A single CertificateRef to a Kubernetes Secret has 'Core' support. Implementations MAY choose to support attaching multiple certificates to a Listener, but this behavior is implementation-specific. References to a resource in different namespace are invalid UNLESS there is a ReferenceGrant in the target namespace that allows the certificate to be attached. If a ReferenceGrant does not allow this reference, the 'ResolvedRefs' condition MUST be set to False for this listener with the 'RefNotPermitted' reason. This field is required to have at least one element when the mode is set to 'Terminate' (default) and is optional otherwise. CertificateRefs can reference to standard Kubernetes resources, i.e. Secret, or implementation-specific custom resources. Support: Core - A single reference to a Kubernetes Secret of type kubernetes.io/tls Support: Implementation-specific (More than one reference or other resource types)",
											MarkdownDescription: "CertificateRefs contains a series of references to Kubernetes objects that contains TLS certificates and private keys. These certificates are used to establish a TLS handshake for requests that match the hostname of the associated listener. A single CertificateRef to a Kubernetes Secret has 'Core' support. Implementations MAY choose to support attaching multiple certificates to a Listener, but this behavior is implementation-specific. References to a resource in different namespace are invalid UNLESS there is a ReferenceGrant in the target namespace that allows the certificate to be attached. If a ReferenceGrant does not allow this reference, the 'ResolvedRefs' condition MUST be set to False for this listener with the 'RefNotPermitted' reason. This field is required to have at least one element when the mode is set to 'Terminate' (default) and is optional otherwise. CertificateRefs can reference to standard Kubernetes resources, i.e. Secret, or implementation-specific custom resources. Support: Core - A single reference to a Kubernetes Secret of type kubernetes.io/tls Support: Implementation-specific (More than one reference or other resource types)",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"group": schema.StringAttribute{
														Description:         "Group is the group of the referent. For example, 'gateway.networking.k8s.io'. When unspecified or empty string, core API group is inferred.",
														MarkdownDescription: "Group is the group of the referent. For example, 'gateway.networking.k8s.io'. When unspecified or empty string, core API group is inferred.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtMost(253),
															stringvalidator.RegexMatches(regexp.MustCompile(`^$|^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
														},
													},

													"kind": schema.StringAttribute{
														Description:         "Kind is kind of the referent. For example 'Secret'.",
														MarkdownDescription: "Kind is kind of the referent. For example 'Secret'.",
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
														Description:         "Namespace is the namespace of the referenced object. When unspecified, the local namespace is inferred. Note that when a namespace different than the local namespace is specified, a ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details. Support: Core",
														MarkdownDescription: "Namespace is the namespace of the referenced object. When unspecified, the local namespace is inferred. Note that when a namespace different than the local namespace is specified, a ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details. Support: Core",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.LengthAtMost(63),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
														},
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"frontend_validation": schema.SingleNestedAttribute{
											Description:         "FrontendValidation holds configuration information for validating the frontend (client). Setting this field will require clients to send a client certificate required for validation during the TLS handshake. In browsers this may result in a dialog appearing that requests a user to specify the client certificate. The maximum depth of a certificate chain accepted in verification is Implementation specific. Support: Extended",
											MarkdownDescription: "FrontendValidation holds configuration information for validating the frontend (client). Setting this field will require clients to send a client certificate required for validation during the TLS handshake. In browsers this may result in a dialog appearing that requests a user to specify the client certificate. The maximum depth of a certificate chain accepted in verification is Implementation specific. Support: Extended",
											Attributes: map[string]schema.Attribute{
												"ca_certificate_refs": schema.ListNestedAttribute{
													Description:         "CACertificateRefs contains one or more references to Kubernetes objects that contain TLS certificates of the Certificate Authorities that can be used as a trust anchor to validate the certificates presented by the client. A single CA certificate reference to a Kubernetes ConfigMap has 'Core' support. Implementations MAY choose to support attaching multiple CA certificates to a Listener, but this behavior is implementation-specific. Support: Core - A single reference to a Kubernetes ConfigMap with the CA certificate in a key named 'ca.crt'. Support: Implementation-specific (More than one reference, or other kinds of resources). References to a resource in a different namespace are invalid UNLESS there is a ReferenceGrant in the target namespace that allows the certificate to be attached. If a ReferenceGrant does not allow this reference, the 'ResolvedRefs' condition MUST be set to False for this listener with the 'RefNotPermitted' reason.",
													MarkdownDescription: "CACertificateRefs contains one or more references to Kubernetes objects that contain TLS certificates of the Certificate Authorities that can be used as a trust anchor to validate the certificates presented by the client. A single CA certificate reference to a Kubernetes ConfigMap has 'Core' support. Implementations MAY choose to support attaching multiple CA certificates to a Listener, but this behavior is implementation-specific. Support: Core - A single reference to a Kubernetes ConfigMap with the CA certificate in a key named 'ca.crt'. Support: Implementation-specific (More than one reference, or other kinds of resources). References to a resource in a different namespace are invalid UNLESS there is a ReferenceGrant in the target namespace that allows the certificate to be attached. If a ReferenceGrant does not allow this reference, the 'ResolvedRefs' condition MUST be set to False for this listener with the 'RefNotPermitted' reason.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"group": schema.StringAttribute{
																Description:         "Group is the group of the referent. For example, 'gateway.networking.k8s.io'. When set to the empty string, core API group is inferred.",
																MarkdownDescription: "Group is the group of the referent. For example, 'gateway.networking.k8s.io'. When set to the empty string, core API group is inferred.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtMost(253),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^$|^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
																},
															},

															"kind": schema.StringAttribute{
																Description:         "Kind is kind of the referent. For example 'ConfigMap' or 'Service'.",
																MarkdownDescription: "Kind is kind of the referent. For example 'ConfigMap' or 'Service'.",
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

															"namespace": schema.StringAttribute{
																Description:         "Namespace is the namespace of the referenced object. When unspecified, the local namespace is inferred. Note that when a namespace different than the local namespace is specified, a ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details. Support: Core",
																MarkdownDescription: "Namespace is the namespace of the referenced object. When unspecified, the local namespace is inferred. Note that when a namespace different than the local namespace is specified, a ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details. Support: Core",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																	stringvalidator.LengthAtMost(63),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
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

										"mode": schema.StringAttribute{
											Description:         "Mode defines the TLS behavior for the TLS session initiated by the client. There are two possible modes: - Terminate: The TLS session between the downstream client and the Gateway is terminated at the Gateway. This mode requires certificates to be specified in some way, such as populating the certificateRefs field. - Passthrough: The TLS session is NOT terminated by the Gateway. This implies that the Gateway can't decipher the TLS stream except for the ClientHello message of the TLS protocol. The certificateRefs field is ignored in this mode. Support: Core",
											MarkdownDescription: "Mode defines the TLS behavior for the TLS session initiated by the client. There are two possible modes: - Terminate: The TLS session between the downstream client and the Gateway is terminated at the Gateway. This mode requires certificates to be specified in some way, such as populating the certificateRefs field. - Passthrough: The TLS session is NOT terminated by the Gateway. This implies that the Gateway can't decipher the TLS stream except for the ClientHello message of the TLS protocol. The certificateRefs field is ignored in this mode. Support: Core",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Terminate", "Passthrough"),
											},
										},

										"options": schema.MapAttribute{
											Description:         "Options are a list of key/value pairs to enable extended TLS configuration for each implementation. For example, configuring the minimum TLS version or supported cipher suites. A set of common keys MAY be defined by the API in the future. To avoid any ambiguity, implementation-specific definitions MUST use domain-prefixed names, such as 'example.com/my-custom-option'. Un-prefixed names are reserved for key names defined by Gateway API. Support: Implementation-specific",
											MarkdownDescription: "Options are a list of key/value pairs to enable extended TLS configuration for each implementation. For example, configuring the minimum TLS version or supported cipher suites. A set of common keys MAY be defined by the API in the future. To avoid any ambiguity, implementation-specific definitions MUST use domain-prefixed names, such as 'example.com/my-custom-option'. Un-prefixed names are reserved for key names defined by Gateway API. Support: Implementation-specific",
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
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"parent_ref": schema.SingleNestedAttribute{
						Description:         "ParentRef references the Gateway that the listeners are attached to.",
						MarkdownDescription: "ParentRef references the Gateway that the listeners are attached to.",
						Attributes: map[string]schema.Attribute{
							"group": schema.StringAttribute{
								Description:         "Group is the group of the referent.",
								MarkdownDescription: "Group is the group of the referent.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtMost(253),
									stringvalidator.RegexMatches(regexp.MustCompile(`^$|^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
								},
							},

							"kind": schema.StringAttribute{
								Description:         "Kind is kind of the referent. For example 'Gateway'.",
								MarkdownDescription: "Kind is kind of the referent. For example 'Gateway'.",
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
								Description:         "Namespace is the namespace of the referent. If not present, the namespace of the referent is assumed to be the same as the namespace of the referring object.",
								MarkdownDescription: "Namespace is the namespace of the referent. If not present, the namespace of the referent is assumed to be the same as the namespace of the referring object.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(63),
									stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
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

func (r *GatewayNetworkingXK8SIoXlistenerSetV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_gateway_networking_x_k8s_io_x_listener_set_v1alpha1_manifest")

	var model GatewayNetworkingXK8SIoXlistenerSetV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("gateway.networking.x-k8s.io/v1alpha1")
	model.Kind = pointer.String("XListenerSet")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
