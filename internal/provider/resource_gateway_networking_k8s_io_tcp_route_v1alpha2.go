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

type GatewayNetworkingK8SIoTCPRouteV1Alpha2Resource struct{}

var (
	_ resource.Resource = (*GatewayNetworkingK8SIoTCPRouteV1Alpha2Resource)(nil)
)

type GatewayNetworkingK8SIoTCPRouteV1Alpha2TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type GatewayNetworkingK8SIoTCPRouteV1Alpha2GoModel struct {
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
		ParentRefs *[]struct {
			Group *string `tfsdk:"group" yaml:"group,omitempty"`

			Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

			SectionName *string `tfsdk:"section_name" yaml:"sectionName,omitempty"`
		} `tfsdk:"parent_refs" yaml:"parentRefs,omitempty"`

		Rules *[]struct {
			BackendRefs *[]struct {
				Group *string `tfsdk:"group" yaml:"group,omitempty"`

				Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

				Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

				Weight *int64 `tfsdk:"weight" yaml:"weight,omitempty"`
			} `tfsdk:"backend_refs" yaml:"backendRefs,omitempty"`
		} `tfsdk:"rules" yaml:"rules,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewGatewayNetworkingK8SIoTCPRouteV1Alpha2Resource() resource.Resource {
	return &GatewayNetworkingK8SIoTCPRouteV1Alpha2Resource{}
}

func (r *GatewayNetworkingK8SIoTCPRouteV1Alpha2Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_networking_k8s_io_tcp_route_v1alpha2"
}

func (r *GatewayNetworkingK8SIoTCPRouteV1Alpha2Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "TCPRoute provides a way to route TCP requests. When combined with a Gateway listener, it can be used to forward connections on the port specified by the listener to a set of backends specified by the TCPRoute.",
		MarkdownDescription: "TCPRoute provides a way to route TCP requests. When combined with a Gateway listener, it can be used to forward connections on the port specified by the listener to a set of backends specified by the TCPRoute.",
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
				Description:         "Spec defines the desired state of TCPRoute.",
				MarkdownDescription: "Spec defines the desired state of TCPRoute.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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
								Description:         "SectionName is the name of a section within the target resource. In the following resources, SectionName is interpreted as the following:  * Gateway: Listener Name  Implementations MAY choose to support attaching Routes to other resources. If that is the case, they MUST clearly document how SectionName is interpreted.  When unspecified (empty string), this will reference the entire resource. For the purpose of status, an attachment is considered successful if at least one section in the parent resource accepts it. For example, Gateway listeners can restrict which Routes can attach to them by Route kind, namespace, or hostname. If 1 of 2 Gateway listeners accept attachment from the referencing Route, the Route MUST be considered successfully attached. If no Gateway listeners accept attachment from this Route, the Route MUST be considered detached from the Gateway.  Support: Core",
								MarkdownDescription: "SectionName is the name of a section within the target resource. In the following resources, SectionName is interpreted as the following:  * Gateway: Listener Name  Implementations MAY choose to support attaching Routes to other resources. If that is the case, they MUST clearly document how SectionName is interpreted.  When unspecified (empty string), this will reference the entire resource. For the purpose of status, an attachment is considered successful if at least one section in the parent resource accepts it. For example, Gateway listeners can restrict which Routes can attach to them by Route kind, namespace, or hostname. If 1 of 2 Gateway listeners accept attachment from the referencing Route, the Route MUST be considered successfully attached. If no Gateway listeners accept attachment from this Route, the Route MUST be considered detached from the Gateway.  Support: Core",

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
						Description:         "Rules are a list of TCP matchers and actions.",
						MarkdownDescription: "Rules are a list of TCP matchers and actions.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"backend_refs": {
								Description:         "BackendRefs defines the backend(s) where matching requests should be sent. If unspecified or invalid (refers to a non-existent resource or a Service with no endpoints), the underlying implementation MUST actively reject connection attempts to this backend. Connection rejections must respect weight; if an invalid backend is requested to have 80% of connections, then 80% of connections must be rejected instead.  Support: Core for Kubernetes Service Support: Custom for any other resource  Support for weight: Extended",
								MarkdownDescription: "BackendRefs defines the backend(s) where matching requests should be sent. If unspecified or invalid (refers to a non-existent resource or a Service with no endpoints), the underlying implementation MUST actively reject connection attempts to this backend. Connection rejections must respect weight; if an invalid backend is requested to have 80% of connections, then 80% of connections must be rejected instead.  Support: Core for Kubernetes Service Support: Custom for any other resource  Support for weight: Extended",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"group": {
										Description:         "Group is the group of the referent. For example, 'networking.k8s.io'. When unspecified (empty string), core API group is inferred.",
										MarkdownDescription: "Group is the group of the referent. For example, 'networking.k8s.io'. When unspecified (empty string), core API group is inferred.",

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
										Description:         "Kind is kind of the referent. For example 'HTTPRoute' or 'Service'.",
										MarkdownDescription: "Kind is kind of the referent. For example 'HTTPRoute' or 'Service'.",

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
										Description:         "Name is the name of the referent.",
										MarkdownDescription: "Name is the name of the referent.",

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
										Description:         "Namespace is the namespace of the backend. When unspecified, the local namespace is inferred.  Note that when a namespace is specified, a ReferencePolicy object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferencePolicy documentation for details.  Support: Core",
										MarkdownDescription: "Namespace is the namespace of the backend. When unspecified, the local namespace is inferred.  Note that when a namespace is specified, a ReferencePolicy object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferencePolicy documentation for details.  Support: Core",

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

									"port": {
										Description:         "Port specifies the destination port number to use for this resource. Port is required when the referent is a Kubernetes Service. For other resources, destination port might be derived from the referent resource or this field.",
										MarkdownDescription: "Port specifies the destination port number to use for this resource. Port is required when the referent is a Kubernetes Service. For other resources, destination port might be derived from the referent resource or this field.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),

											int64validator.AtMost(65535),
										},
									},

									"weight": {
										Description:         "Weight specifies the proportion of requests forwarded to the referenced backend. This is computed as weight/(sum of all weights in this BackendRefs list). For non-zero values, there may be some epsilon from the exact proportion defined here depending on the precision an implementation supports. Weight is not a percentage and the sum of weights does not need to equal 100.  If only one backend is specified and it has a weight greater than 0, 100% of the traffic is forwarded to that backend. If weight is set to 0, no traffic should be forwarded for this entry. If unspecified, weight defaults to 1.  Support for this field varies based on the context where used.",
										MarkdownDescription: "Weight specifies the proportion of requests forwarded to the referenced backend. This is computed as weight/(sum of all weights in this BackendRefs list). For non-zero values, there may be some epsilon from the exact proportion defined here depending on the precision an implementation supports. Weight is not a percentage and the sum of weights does not need to equal 100.  If only one backend is specified and it has a weight greater than 0, 100% of the traffic is forwarded to that backend. If weight is set to 0, no traffic should be forwarded for this entry. If unspecified, weight defaults to 1.  Support for this field varies based on the context where used.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(0),

											int64validator.AtMost(1e+06),
										},
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

func (r *GatewayNetworkingK8SIoTCPRouteV1Alpha2Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_gateway_networking_k8s_io_tcp_route_v1alpha2")

	var state GatewayNetworkingK8SIoTCPRouteV1Alpha2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel GatewayNetworkingK8SIoTCPRouteV1Alpha2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("gateway.networking.k8s.io/v1alpha2")
	goModel.Kind = utilities.Ptr("TCPRoute")

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

func (r *GatewayNetworkingK8SIoTCPRouteV1Alpha2Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_gateway_networking_k8s_io_tcp_route_v1alpha2")
	// NO-OP: All data is already in Terraform state
}

func (r *GatewayNetworkingK8SIoTCPRouteV1Alpha2Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_gateway_networking_k8s_io_tcp_route_v1alpha2")

	var state GatewayNetworkingK8SIoTCPRouteV1Alpha2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel GatewayNetworkingK8SIoTCPRouteV1Alpha2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("gateway.networking.k8s.io/v1alpha2")
	goModel.Kind = utilities.Ptr("TCPRoute")

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

func (r *GatewayNetworkingK8SIoTCPRouteV1Alpha2Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_gateway_networking_k8s_io_tcp_route_v1alpha2")
	// NO-OP: Terraform removes the state automatically for us
}
