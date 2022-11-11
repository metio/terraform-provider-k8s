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

type DiscoveryK8SIoEndpointSliceV1Resource struct{}

var (
	_ resource.Resource = (*DiscoveryK8SIoEndpointSliceV1Resource)(nil)
)

type DiscoveryK8SIoEndpointSliceV1TerraformModel struct {
	Id          types.Int64  `tfsdk:"id"`
	YAML        types.String `tfsdk:"yaml"`
	ApiVersion  types.String `tfsdk:"api_version"`
	Kind        types.String `tfsdk:"kind"`
	Metadata    types.Object `tfsdk:"metadata"`
	AddressType types.String `tfsdk:"address_type"`
	Endpoints   types.List   `tfsdk:"endpoints"`
	Ports       types.List   `tfsdk:"ports"`
}

type DiscoveryK8SIoEndpointSliceV1GoModel struct {
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

	AddressType *string `tfsdk:"address_type" yaml:"addressType,omitempty"`

	Endpoints *[]struct {
		Addresses *[]string `tfsdk:"addresses" yaml:"addresses,omitempty"`

		Conditions *struct {
			Ready *bool `tfsdk:"ready" yaml:"ready,omitempty"`

			Serving *bool `tfsdk:"serving" yaml:"serving,omitempty"`

			Terminating *bool `tfsdk:"terminating" yaml:"terminating,omitempty"`
		} `tfsdk:"conditions" yaml:"conditions,omitempty"`

		DeprecatedTopology *map[string]string `tfsdk:"deprecated_topology" yaml:"deprecatedTopology,omitempty"`

		Hints *struct {
			ForZones *[]struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"for_zones" yaml:"forZones,omitempty"`
		} `tfsdk:"hints" yaml:"hints,omitempty"`

		Hostname *string `tfsdk:"hostname" yaml:"hostname,omitempty"`

		NodeName *string `tfsdk:"node_name" yaml:"nodeName,omitempty"`

		TargetRef *struct {
			ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

			FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`

			Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

			ResourceVersion *string `tfsdk:"resource_version" yaml:"resourceVersion,omitempty"`

			Uid *string `tfsdk:"uid" yaml:"uid,omitempty"`
		} `tfsdk:"target_ref" yaml:"targetRef,omitempty"`

		Zone *string `tfsdk:"zone" yaml:"zone,omitempty"`
	} `tfsdk:"endpoints" yaml:"endpoints,omitempty"`

	Ports *[]struct {
		AppProtocol *string `tfsdk:"app_protocol" yaml:"appProtocol,omitempty"`

		Name *string `tfsdk:"name" yaml:"name,omitempty"`

		Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

		Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`
	} `tfsdk:"ports" yaml:"ports,omitempty"`
}

func NewDiscoveryK8SIoEndpointSliceV1Resource() resource.Resource {
	return &DiscoveryK8SIoEndpointSliceV1Resource{}
}

func (r *DiscoveryK8SIoEndpointSliceV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_discovery_k8s_io_endpoint_slice_v1"
}

func (r *DiscoveryK8SIoEndpointSliceV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "EndpointSlice represents a subset of the endpoints that implement a service. For a given service there may be multiple EndpointSlice objects, selected by labels, which must be joined to produce the full set of endpoints.",
		MarkdownDescription: "EndpointSlice represents a subset of the endpoints that implement a service. For a given service there may be multiple EndpointSlice objects, selected by labels, which must be joined to produce the full set of endpoints.",
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

			"address_type": {
				Description:         "addressType specifies the type of address carried by this EndpointSlice. All addresses in this slice must be the same type. This field is immutable after creation. The following address types are currently supported: * IPv4: Represents an IPv4 Address. * IPv6: Represents an IPv6 Address. * FQDN: Represents a Fully Qualified Domain Name.",
				MarkdownDescription: "addressType specifies the type of address carried by this EndpointSlice. All addresses in this slice must be the same type. This field is immutable after creation. The following address types are currently supported: * IPv4: Represents an IPv4 Address. * IPv6: Represents an IPv6 Address. * FQDN: Represents a Fully Qualified Domain Name.",

				Type: types.StringType,

				Required: true,
				Optional: false,
				Computed: false,
			},

			"endpoints": {
				Description:         "endpoints is a list of unique endpoints in this slice. Each slice may include a maximum of 1000 endpoints.",
				MarkdownDescription: "endpoints is a list of unique endpoints in this slice. Each slice may include a maximum of 1000 endpoints.",

				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

					"addresses": {
						Description:         "addresses of this endpoint. The contents of this field are interpreted according to the corresponding EndpointSlice addressType field. Consumers must handle different types of addresses in the context of their own capabilities. This must contain at least one address but no more than 100. These are all assumed to be fungible and clients may choose to only use the first element. Refer to: https://issue.k8s.io/106267",
						MarkdownDescription: "addresses of this endpoint. The contents of this field are interpreted according to the corresponding EndpointSlice addressType field. Consumers must handle different types of addresses in the context of their own capabilities. This must contain at least one address but no more than 100. These are all assumed to be fungible and clients may choose to only use the first element. Refer to: https://issue.k8s.io/106267",

						Type: types.ListType{ElemType: types.StringType},

						Required: true,
						Optional: false,
						Computed: false,
					},

					"conditions": {
						Description:         "EndpointConditions represents the current condition of an endpoint.",
						MarkdownDescription: "EndpointConditions represents the current condition of an endpoint.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"ready": {
								Description:         "ready indicates that this endpoint is prepared to receive traffic, according to whatever system is managing the endpoint. A nil value indicates an unknown state. In most cases consumers should interpret this unknown state as ready. For compatibility reasons, ready should never be 'true' for terminating endpoints.",
								MarkdownDescription: "ready indicates that this endpoint is prepared to receive traffic, according to whatever system is managing the endpoint. A nil value indicates an unknown state. In most cases consumers should interpret this unknown state as ready. For compatibility reasons, ready should never be 'true' for terminating endpoints.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"serving": {
								Description:         "serving is identical to ready except that it is set regardless of the terminating state of endpoints. This condition should be set to true for a ready endpoint that is terminating. If nil, consumers should defer to the ready condition.",
								MarkdownDescription: "serving is identical to ready except that it is set regardless of the terminating state of endpoints. This condition should be set to true for a ready endpoint that is terminating. If nil, consumers should defer to the ready condition.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"terminating": {
								Description:         "terminating indicates that this endpoint is terminating. A nil value indicates an unknown state. Consumers should interpret this unknown state to mean that the endpoint is not terminating.",
								MarkdownDescription: "terminating indicates that this endpoint is terminating. A nil value indicates an unknown state. Consumers should interpret this unknown state to mean that the endpoint is not terminating.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"deprecated_topology": {
						Description:         "deprecatedTopology contains topology information part of the v1beta1 API. This field is deprecated, and will be removed when the v1beta1 API is removed (no sooner than kubernetes v1.24).  While this field can hold values, it is not writable through the v1 API, and any attempts to write to it will be silently ignored. Topology information can be found in the zone and nodeName fields instead.",
						MarkdownDescription: "deprecatedTopology contains topology information part of the v1beta1 API. This field is deprecated, and will be removed when the v1beta1 API is removed (no sooner than kubernetes v1.24).  While this field can hold values, it is not writable through the v1 API, and any attempts to write to it will be silently ignored. Topology information can be found in the zone and nodeName fields instead.",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"hints": {
						Description:         "EndpointHints provides hints describing how an endpoint should be consumed.",
						MarkdownDescription: "EndpointHints provides hints describing how an endpoint should be consumed.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"for_zones": {
								Description:         "forZones indicates the zone(s) this endpoint should be consumed by to enable topology aware routing.",
								MarkdownDescription: "forZones indicates the zone(s) this endpoint should be consumed by to enable topology aware routing.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "name represents the name of the zone.",
										MarkdownDescription: "name represents the name of the zone.",

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
						Description:         "hostname of this endpoint. This field may be used by consumers of endpoints to distinguish endpoints from each other (e.g. in DNS names). Multiple endpoints which use the same hostname should be considered fungible (e.g. multiple A values in DNS). Must be lowercase and pass DNS Label (RFC 1123) validation.",
						MarkdownDescription: "hostname of this endpoint. This field may be used by consumers of endpoints to distinguish endpoints from each other (e.g. in DNS names). Multiple endpoints which use the same hostname should be considered fungible (e.g. multiple A values in DNS). Must be lowercase and pass DNS Label (RFC 1123) validation.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"node_name": {
						Description:         "nodeName represents the name of the Node hosting this endpoint. This can be used to determine endpoints local to a Node.",
						MarkdownDescription: "nodeName represents the name of the Node hosting this endpoint. This can be used to determine endpoints local to a Node.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"target_ref": {
						Description:         "ObjectReference contains enough information to let you inspect or modify the referred object.",
						MarkdownDescription: "ObjectReference contains enough information to let you inspect or modify the referred object.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"api_version": {
								Description:         "API version of the referent.",
								MarkdownDescription: "API version of the referent.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"field_path": {
								Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.",
								MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"kind": {
								Description:         "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								MarkdownDescription: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"namespace": {
								Description:         "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
								MarkdownDescription: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"resource_version": {
								Description:         "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
								MarkdownDescription: "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"uid": {
								Description:         "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
								MarkdownDescription: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",

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

					"zone": {
						Description:         "zone is the name of the Zone this endpoint exists in.",
						MarkdownDescription: "zone is the name of the Zone this endpoint exists in.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: true,
				Optional: false,
				Computed: false,
			},

			"ports": {
				Description:         "ports specifies the list of network ports exposed by each endpoint in this slice. Each port must have a unique name. When ports is empty, it indicates that there are no defined ports. When a port is defined with a nil port value, it indicates 'all ports'. Each slice may include a maximum of 100 ports.",
				MarkdownDescription: "ports specifies the list of network ports exposed by each endpoint in this slice. Each port must have a unique name. When ports is empty, it indicates that there are no defined ports. When a port is defined with a nil port value, it indicates 'all ports'. Each slice may include a maximum of 100 ports.",

				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

					"app_protocol": {
						Description:         "The application protocol for this port. This field follows standard Kubernetes label syntax. Un-prefixed names are reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names). Non-standard protocols should use prefixed names such as mycompany.com/my-custom-protocol.",
						MarkdownDescription: "The application protocol for this port. This field follows standard Kubernetes label syntax. Un-prefixed names are reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names). Non-standard protocols should use prefixed names such as mycompany.com/my-custom-protocol.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"name": {
						Description:         "The name of this port. All ports in an EndpointSlice must have a unique name. If the EndpointSlice is dervied from a Kubernetes service, this corresponds to the Service.ports[].name. Name must either be an empty string or pass DNS_LABEL validation: * must be no more than 63 characters long. * must consist of lower case alphanumeric characters or '-'. * must start and end with an alphanumeric character. Default is empty string.",
						MarkdownDescription: "The name of this port. All ports in an EndpointSlice must have a unique name. If the EndpointSlice is dervied from a Kubernetes service, this corresponds to the Service.ports[].name. Name must either be an empty string or pass DNS_LABEL validation: * must be no more than 63 characters long. * must consist of lower case alphanumeric characters or '-'. * must start and end with an alphanumeric character. Default is empty string.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"port": {
						Description:         "The port number of the endpoint. If this is not specified, ports are not restricted and must be interpreted in the context of the specific consumer.",
						MarkdownDescription: "The port number of the endpoint. If this is not specified, ports are not restricted and must be interpreted in the context of the specific consumer.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"protocol": {
						Description:         "The IP protocol for this port. Must be UDP, TCP, or SCTP. Default is TCP.",
						MarkdownDescription: "The IP protocol for this port. Must be UDP, TCP, or SCTP. Default is TCP.",

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
		},
	}, nil
}

func (r *DiscoveryK8SIoEndpointSliceV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_discovery_k8s_io_endpoint_slice_v1")

	var state DiscoveryK8SIoEndpointSliceV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel DiscoveryK8SIoEndpointSliceV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("discovery.k8s.io/v1")
	goModel.Kind = utilities.Ptr("EndpointSlice")

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

func (r *DiscoveryK8SIoEndpointSliceV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_discovery_k8s_io_endpoint_slice_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *DiscoveryK8SIoEndpointSliceV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_discovery_k8s_io_endpoint_slice_v1")

	var state DiscoveryK8SIoEndpointSliceV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel DiscoveryK8SIoEndpointSliceV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("discovery.k8s.io/v1")
	goModel.Kind = utilities.Ptr("EndpointSlice")

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

func (r *DiscoveryK8SIoEndpointSliceV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_discovery_k8s_io_endpoint_slice_v1")
	// NO-OP: Terraform removes the state automatically for us
}
