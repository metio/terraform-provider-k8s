/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package discovery_k8s_io_v1

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
	_ datasource.DataSource = &DiscoveryK8SIoEndpointSliceV1Manifest{}
)

func NewDiscoveryK8SIoEndpointSliceV1Manifest() datasource.DataSource {
	return &DiscoveryK8SIoEndpointSliceV1Manifest{}
}

type DiscoveryK8SIoEndpointSliceV1Manifest struct{}

type DiscoveryK8SIoEndpointSliceV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	AddressType *string `tfsdk:"address_type" json:"addressType,omitempty"`
	Endpoints   *[]struct {
		Addresses  *[]string `tfsdk:"addresses" json:"addresses,omitempty"`
		Conditions *struct {
			Ready       *bool `tfsdk:"ready" json:"ready,omitempty"`
			Serving     *bool `tfsdk:"serving" json:"serving,omitempty"`
			Terminating *bool `tfsdk:"terminating" json:"terminating,omitempty"`
		} `tfsdk:"conditions" json:"conditions,omitempty"`
		DeprecatedTopology *map[string]string `tfsdk:"deprecated_topology" json:"deprecatedTopology,omitempty"`
		Hints              *struct {
			ForZones *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"for_zones" json:"forZones,omitempty"`
		} `tfsdk:"hints" json:"hints,omitempty"`
		Hostname  *string `tfsdk:"hostname" json:"hostname,omitempty"`
		NodeName  *string `tfsdk:"node_name" json:"nodeName,omitempty"`
		TargetRef *struct {
			ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
			Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
			Name            *string `tfsdk:"name" json:"name,omitempty"`
			Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
			ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
			Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
		} `tfsdk:"target_ref" json:"targetRef,omitempty"`
		Zone *string `tfsdk:"zone" json:"zone,omitempty"`
	} `tfsdk:"endpoints" json:"endpoints,omitempty"`
	Ports *[]struct {
		AppProtocol *string `tfsdk:"app_protocol" json:"appProtocol,omitempty"`
		Name        *string `tfsdk:"name" json:"name,omitempty"`
		Port        *int64  `tfsdk:"port" json:"port,omitempty"`
		Protocol    *string `tfsdk:"protocol" json:"protocol,omitempty"`
	} `tfsdk:"ports" json:"ports,omitempty"`
}

func (r *DiscoveryK8SIoEndpointSliceV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_discovery_k8s_io_endpoint_slice_v1_manifest"
}

func (r *DiscoveryK8SIoEndpointSliceV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "EndpointSlice represents a subset of the endpoints that implement a service. For a given service there may be multiple EndpointSlice objects, selected by labels, which must be joined to produce the full set of endpoints.",
		MarkdownDescription: "EndpointSlice represents a subset of the endpoints that implement a service. For a given service there may be multiple EndpointSlice objects, selected by labels, which must be joined to produce the full set of endpoints.",
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

			"address_type": schema.StringAttribute{
				Description:         "addressType specifies the type of address carried by this EndpointSlice. All addresses in this slice must be the same type. This field is immutable after creation. The following address types are currently supported: * IPv4: Represents an IPv4 Address. * IPv6: Represents an IPv6 Address. * FQDN: Represents a Fully Qualified Domain Name.",
				MarkdownDescription: "addressType specifies the type of address carried by this EndpointSlice. All addresses in this slice must be the same type. This field is immutable after creation. The following address types are currently supported: * IPv4: Represents an IPv4 Address. * IPv6: Represents an IPv6 Address. * FQDN: Represents a Fully Qualified Domain Name.",
				Required:            true,
				Optional:            false,
				Computed:            false,
			},

			"endpoints": schema.ListNestedAttribute{
				Description:         "endpoints is a list of unique endpoints in this slice. Each slice may include a maximum of 1000 endpoints.",
				MarkdownDescription: "endpoints is a list of unique endpoints in this slice. Each slice may include a maximum of 1000 endpoints.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"addresses": schema.ListAttribute{
							Description:         "addresses of this endpoint. The contents of this field are interpreted according to the corresponding EndpointSlice addressType field. Consumers must handle different types of addresses in the context of their own capabilities. This must contain at least one address but no more than 100. These are all assumed to be fungible and clients may choose to only use the first element. Refer to: https://issue.k8s.io/106267",
							MarkdownDescription: "addresses of this endpoint. The contents of this field are interpreted according to the corresponding EndpointSlice addressType field. Consumers must handle different types of addresses in the context of their own capabilities. This must contain at least one address but no more than 100. These are all assumed to be fungible and clients may choose to only use the first element. Refer to: https://issue.k8s.io/106267",
							ElementType:         types.StringType,
							Required:            true,
							Optional:            false,
							Computed:            false,
						},

						"conditions": schema.SingleNestedAttribute{
							Description:         "EndpointConditions represents the current condition of an endpoint.",
							MarkdownDescription: "EndpointConditions represents the current condition of an endpoint.",
							Attributes: map[string]schema.Attribute{
								"ready": schema.BoolAttribute{
									Description:         "ready indicates that this endpoint is prepared to receive traffic, according to whatever system is managing the endpoint. A nil value indicates an unknown state. In most cases consumers should interpret this unknown state as ready. For compatibility reasons, ready should never be 'true' for terminating endpoints, except when the normal readiness behavior is being explicitly overridden, for example when the associated Service has set the publishNotReadyAddresses flag.",
									MarkdownDescription: "ready indicates that this endpoint is prepared to receive traffic, according to whatever system is managing the endpoint. A nil value indicates an unknown state. In most cases consumers should interpret this unknown state as ready. For compatibility reasons, ready should never be 'true' for terminating endpoints, except when the normal readiness behavior is being explicitly overridden, for example when the associated Service has set the publishNotReadyAddresses flag.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"serving": schema.BoolAttribute{
									Description:         "serving is identical to ready except that it is set regardless of the terminating state of endpoints. This condition should be set to true for a ready endpoint that is terminating. If nil, consumers should defer to the ready condition.",
									MarkdownDescription: "serving is identical to ready except that it is set regardless of the terminating state of endpoints. This condition should be set to true for a ready endpoint that is terminating. If nil, consumers should defer to the ready condition.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"terminating": schema.BoolAttribute{
									Description:         "terminating indicates that this endpoint is terminating. A nil value indicates an unknown state. Consumers should interpret this unknown state to mean that the endpoint is not terminating.",
									MarkdownDescription: "terminating indicates that this endpoint is terminating. A nil value indicates an unknown state. Consumers should interpret this unknown state to mean that the endpoint is not terminating.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},
							},
							Required: false,
							Optional: true,
							Computed: false,
						},

						"deprecated_topology": schema.MapAttribute{
							Description:         "deprecatedTopology contains topology information part of the v1beta1 API. This field is deprecated, and will be removed when the v1beta1 API is removed (no sooner than kubernetes v1.24). While this field can hold values, it is not writable through the v1 API, and any attempts to write to it will be silently ignored. Topology information can be found in the zone and nodeName fields instead.",
							MarkdownDescription: "deprecatedTopology contains topology information part of the v1beta1 API. This field is deprecated, and will be removed when the v1beta1 API is removed (no sooner than kubernetes v1.24). While this field can hold values, it is not writable through the v1 API, and any attempts to write to it will be silently ignored. Topology information can be found in the zone and nodeName fields instead.",
							ElementType:         types.StringType,
							Required:            false,
							Optional:            true,
							Computed:            false,
						},

						"hints": schema.SingleNestedAttribute{
							Description:         "EndpointHints provides hints describing how an endpoint should be consumed.",
							MarkdownDescription: "EndpointHints provides hints describing how an endpoint should be consumed.",
							Attributes: map[string]schema.Attribute{
								"for_zones": schema.ListNestedAttribute{
									Description:         "forZones indicates the zone(s) this endpoint should be consumed by to enable topology aware routing.",
									MarkdownDescription: "forZones indicates the zone(s) this endpoint should be consumed by to enable topology aware routing.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "name represents the name of the zone.",
												MarkdownDescription: "name represents the name of the zone.",
												Required:            true,
												Optional:            false,
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

						"hostname": schema.StringAttribute{
							Description:         "hostname of this endpoint. This field may be used by consumers of endpoints to distinguish endpoints from each other (e.g. in DNS names). Multiple endpoints which use the same hostname should be considered fungible (e.g. multiple A values in DNS). Must be lowercase and pass DNS Label (RFC 1123) validation.",
							MarkdownDescription: "hostname of this endpoint. This field may be used by consumers of endpoints to distinguish endpoints from each other (e.g. in DNS names). Multiple endpoints which use the same hostname should be considered fungible (e.g. multiple A values in DNS). Must be lowercase and pass DNS Label (RFC 1123) validation.",
							Required:            false,
							Optional:            true,
							Computed:            false,
						},

						"node_name": schema.StringAttribute{
							Description:         "nodeName represents the name of the Node hosting this endpoint. This can be used to determine endpoints local to a Node.",
							MarkdownDescription: "nodeName represents the name of the Node hosting this endpoint. This can be used to determine endpoints local to a Node.",
							Required:            false,
							Optional:            true,
							Computed:            false,
						},

						"target_ref": schema.SingleNestedAttribute{
							Description:         "ObjectReference contains enough information to let you inspect or modify the referred object.",
							MarkdownDescription: "ObjectReference contains enough information to let you inspect or modify the referred object.",
							Attributes: map[string]schema.Attribute{
								"api_version": schema.StringAttribute{
									Description:         "API version of the referent.",
									MarkdownDescription: "API version of the referent.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"field_path": schema.StringAttribute{
									Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.",
									MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"kind": schema.StringAttribute{
									Description:         "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
									MarkdownDescription: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
									MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
									MarkdownDescription: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"resource_version": schema.StringAttribute{
									Description:         "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
									MarkdownDescription: "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"uid": schema.StringAttribute{
									Description:         "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
									MarkdownDescription: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},
							},
							Required: false,
							Optional: true,
							Computed: false,
						},

						"zone": schema.StringAttribute{
							Description:         "zone is the name of the Zone this endpoint exists in.",
							MarkdownDescription: "zone is the name of the Zone this endpoint exists in.",
							Required:            false,
							Optional:            true,
							Computed:            false,
						},
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},

			"ports": schema.ListNestedAttribute{
				Description:         "ports specifies the list of network ports exposed by each endpoint in this slice. Each port must have a unique name. When ports is empty, it indicates that there are no defined ports. When a port is defined with a nil port value, it indicates 'all ports'. Each slice may include a maximum of 100 ports.",
				MarkdownDescription: "ports specifies the list of network ports exposed by each endpoint in this slice. Each port must have a unique name. When ports is empty, it indicates that there are no defined ports. When a port is defined with a nil port value, it indicates 'all ports'. Each slice may include a maximum of 100 ports.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"app_protocol": schema.StringAttribute{
							Description:         "The application protocol for this port. This is used as a hint for implementations to offer richer behavior for protocols that they understand. This field follows standard Kubernetes label syntax. Valid values are either: * Un-prefixed protocol names - reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names). * Kubernetes-defined prefixed names: * 'kubernetes.io/h2c' - HTTP/2 prior knowledge over cleartext as described in https://www.rfc-editor.org/rfc/rfc9113.html#name-starting-http-2-with-prior- * 'kubernetes.io/ws' - WebSocket over cleartext as described in https://www.rfc-editor.org/rfc/rfc6455 * 'kubernetes.io/wss' - WebSocket over TLS as described in https://www.rfc-editor.org/rfc/rfc6455 * Other protocols should use implementation-defined prefixed names such as mycompany.com/my-custom-protocol.",
							MarkdownDescription: "The application protocol for this port. This is used as a hint for implementations to offer richer behavior for protocols that they understand. This field follows standard Kubernetes label syntax. Valid values are either: * Un-prefixed protocol names - reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names). * Kubernetes-defined prefixed names: * 'kubernetes.io/h2c' - HTTP/2 prior knowledge over cleartext as described in https://www.rfc-editor.org/rfc/rfc9113.html#name-starting-http-2-with-prior- * 'kubernetes.io/ws' - WebSocket over cleartext as described in https://www.rfc-editor.org/rfc/rfc6455 * 'kubernetes.io/wss' - WebSocket over TLS as described in https://www.rfc-editor.org/rfc/rfc6455 * Other protocols should use implementation-defined prefixed names such as mycompany.com/my-custom-protocol.",
							Required:            false,
							Optional:            true,
							Computed:            false,
						},

						"name": schema.StringAttribute{
							Description:         "name represents the name of this port. All ports in an EndpointSlice must have a unique name. If the EndpointSlice is derived from a Kubernetes service, this corresponds to the Service.ports[].name. Name must either be an empty string or pass DNS_LABEL validation: * must be no more than 63 characters long. * must consist of lower case alphanumeric characters or '-'. * must start and end with an alphanumeric character. Default is empty string.",
							MarkdownDescription: "name represents the name of this port. All ports in an EndpointSlice must have a unique name. If the EndpointSlice is derived from a Kubernetes service, this corresponds to the Service.ports[].name. Name must either be an empty string or pass DNS_LABEL validation: * must be no more than 63 characters long. * must consist of lower case alphanumeric characters or '-'. * must start and end with an alphanumeric character. Default is empty string.",
							Required:            false,
							Optional:            true,
							Computed:            false,
						},

						"port": schema.Int64Attribute{
							Description:         "port represents the port number of the endpoint. If this is not specified, ports are not restricted and must be interpreted in the context of the specific consumer.",
							MarkdownDescription: "port represents the port number of the endpoint. If this is not specified, ports are not restricted and must be interpreted in the context of the specific consumer.",
							Required:            false,
							Optional:            true,
							Computed:            false,
						},

						"protocol": schema.StringAttribute{
							Description:         "protocol represents the IP protocol for this port. Must be UDP, TCP, or SCTP. Default is TCP.",
							MarkdownDescription: "protocol represents the IP protocol for this port. Must be UDP, TCP, or SCTP. Default is TCP.",
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
	}
}

func (r *DiscoveryK8SIoEndpointSliceV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_discovery_k8s_io_endpoint_slice_v1_manifest")

	var model DiscoveryK8SIoEndpointSliceV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("discovery.k8s.io/v1")
	model.Kind = pointer.String("EndpointSlice")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
