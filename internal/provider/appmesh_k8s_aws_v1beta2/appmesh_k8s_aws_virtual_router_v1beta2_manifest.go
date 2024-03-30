/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package appmesh_k8s_aws_v1beta2

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
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &AppmeshK8SAwsVirtualRouterV1Beta2Manifest{}
)

func NewAppmeshK8SAwsVirtualRouterV1Beta2Manifest() datasource.DataSource {
	return &AppmeshK8SAwsVirtualRouterV1Beta2Manifest{}
}

type AppmeshK8SAwsVirtualRouterV1Beta2Manifest struct{}

type AppmeshK8SAwsVirtualRouterV1Beta2ManifestData struct {
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
		AwsName   *string `tfsdk:"aws_name" json:"awsName,omitempty"`
		Listeners *[]struct {
			PortMapping *struct {
				Port     *int64  `tfsdk:"port" json:"port,omitempty"`
				Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
			} `tfsdk:"port_mapping" json:"portMapping,omitempty"`
		} `tfsdk:"listeners" json:"listeners,omitempty"`
		MeshRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
			Uid  *string `tfsdk:"uid" json:"uid,omitempty"`
		} `tfsdk:"mesh_ref" json:"meshRef,omitempty"`
		Routes *[]struct {
			GrpcRoute *struct {
				Action *struct {
					WeightedTargets *[]struct {
						Port           *int64  `tfsdk:"port" json:"port,omitempty"`
						VirtualNodeARN *string `tfsdk:"virtual_node_arn" json:"virtualNodeARN,omitempty"`
						VirtualNodeRef *struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"virtual_node_ref" json:"virtualNodeRef,omitempty"`
						Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
					} `tfsdk:"weighted_targets" json:"weightedTargets,omitempty"`
				} `tfsdk:"action" json:"action,omitempty"`
				Match *struct {
					Metadata *[]struct {
						Invert *bool `tfsdk:"invert" json:"invert,omitempty"`
						Match  *struct {
							Exact  *string `tfsdk:"exact" json:"exact,omitempty"`
							Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
							Range  *struct {
								End   *int64 `tfsdk:"end" json:"end,omitempty"`
								Start *int64 `tfsdk:"start" json:"start,omitempty"`
							} `tfsdk:"range" json:"range,omitempty"`
							Regex  *string `tfsdk:"regex" json:"regex,omitempty"`
							Suffix *string `tfsdk:"suffix" json:"suffix,omitempty"`
						} `tfsdk:"match" json:"match,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
					MethodName  *string `tfsdk:"method_name" json:"methodName,omitempty"`
					Port        *int64  `tfsdk:"port" json:"port,omitempty"`
					ServiceName *string `tfsdk:"service_name" json:"serviceName,omitempty"`
				} `tfsdk:"match" json:"match,omitempty"`
				RetryPolicy *struct {
					GrpcRetryEvents *[]string `tfsdk:"grpc_retry_events" json:"grpcRetryEvents,omitempty"`
					HttpRetryEvents *[]string `tfsdk:"http_retry_events" json:"httpRetryEvents,omitempty"`
					MaxRetries      *int64    `tfsdk:"max_retries" json:"maxRetries,omitempty"`
					PerRetryTimeout *struct {
						Unit  *string `tfsdk:"unit" json:"unit,omitempty"`
						Value *int64  `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"per_retry_timeout" json:"perRetryTimeout,omitempty"`
					TcpRetryEvents *[]string `tfsdk:"tcp_retry_events" json:"tcpRetryEvents,omitempty"`
				} `tfsdk:"retry_policy" json:"retryPolicy,omitempty"`
				Timeout *struct {
					Idle *struct {
						Unit  *string `tfsdk:"unit" json:"unit,omitempty"`
						Value *int64  `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"idle" json:"idle,omitempty"`
					PerRequest *struct {
						Unit  *string `tfsdk:"unit" json:"unit,omitempty"`
						Value *int64  `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"per_request" json:"perRequest,omitempty"`
				} `tfsdk:"timeout" json:"timeout,omitempty"`
			} `tfsdk:"grpc_route" json:"grpcRoute,omitempty"`
			Http2Route *struct {
				Action *struct {
					WeightedTargets *[]struct {
						Port           *int64  `tfsdk:"port" json:"port,omitempty"`
						VirtualNodeARN *string `tfsdk:"virtual_node_arn" json:"virtualNodeARN,omitempty"`
						VirtualNodeRef *struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"virtual_node_ref" json:"virtualNodeRef,omitempty"`
						Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
					} `tfsdk:"weighted_targets" json:"weightedTargets,omitempty"`
				} `tfsdk:"action" json:"action,omitempty"`
				Match *struct {
					Headers *[]struct {
						Invert *bool `tfsdk:"invert" json:"invert,omitempty"`
						Match  *struct {
							Exact  *string `tfsdk:"exact" json:"exact,omitempty"`
							Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
							Range  *struct {
								End   *int64 `tfsdk:"end" json:"end,omitempty"`
								Start *int64 `tfsdk:"start" json:"start,omitempty"`
							} `tfsdk:"range" json:"range,omitempty"`
							Regex  *string `tfsdk:"regex" json:"regex,omitempty"`
							Suffix *string `tfsdk:"suffix" json:"suffix,omitempty"`
						} `tfsdk:"match" json:"match,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"headers" json:"headers,omitempty"`
					Method *string `tfsdk:"method" json:"method,omitempty"`
					Path   *struct {
						Exact *string `tfsdk:"exact" json:"exact,omitempty"`
						Regex *string `tfsdk:"regex" json:"regex,omitempty"`
					} `tfsdk:"path" json:"path,omitempty"`
					Port            *int64  `tfsdk:"port" json:"port,omitempty"`
					Prefix          *string `tfsdk:"prefix" json:"prefix,omitempty"`
					QueryParameters *[]struct {
						Match *struct {
							Exact *string `tfsdk:"exact" json:"exact,omitempty"`
						} `tfsdk:"match" json:"match,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"query_parameters" json:"queryParameters,omitempty"`
					Scheme *string `tfsdk:"scheme" json:"scheme,omitempty"`
				} `tfsdk:"match" json:"match,omitempty"`
				RetryPolicy *struct {
					HttpRetryEvents *[]string `tfsdk:"http_retry_events" json:"httpRetryEvents,omitempty"`
					MaxRetries      *int64    `tfsdk:"max_retries" json:"maxRetries,omitempty"`
					PerRetryTimeout *struct {
						Unit  *string `tfsdk:"unit" json:"unit,omitempty"`
						Value *int64  `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"per_retry_timeout" json:"perRetryTimeout,omitempty"`
					TcpRetryEvents *[]string `tfsdk:"tcp_retry_events" json:"tcpRetryEvents,omitempty"`
				} `tfsdk:"retry_policy" json:"retryPolicy,omitempty"`
				Timeout *struct {
					Idle *struct {
						Unit  *string `tfsdk:"unit" json:"unit,omitempty"`
						Value *int64  `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"idle" json:"idle,omitempty"`
					PerRequest *struct {
						Unit  *string `tfsdk:"unit" json:"unit,omitempty"`
						Value *int64  `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"per_request" json:"perRequest,omitempty"`
				} `tfsdk:"timeout" json:"timeout,omitempty"`
			} `tfsdk:"http2_route" json:"http2Route,omitempty"`
			HttpRoute *struct {
				Action *struct {
					WeightedTargets *[]struct {
						Port           *int64  `tfsdk:"port" json:"port,omitempty"`
						VirtualNodeARN *string `tfsdk:"virtual_node_arn" json:"virtualNodeARN,omitempty"`
						VirtualNodeRef *struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"virtual_node_ref" json:"virtualNodeRef,omitempty"`
						Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
					} `tfsdk:"weighted_targets" json:"weightedTargets,omitempty"`
				} `tfsdk:"action" json:"action,omitempty"`
				Match *struct {
					Headers *[]struct {
						Invert *bool `tfsdk:"invert" json:"invert,omitempty"`
						Match  *struct {
							Exact  *string `tfsdk:"exact" json:"exact,omitempty"`
							Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
							Range  *struct {
								End   *int64 `tfsdk:"end" json:"end,omitempty"`
								Start *int64 `tfsdk:"start" json:"start,omitempty"`
							} `tfsdk:"range" json:"range,omitempty"`
							Regex  *string `tfsdk:"regex" json:"regex,omitempty"`
							Suffix *string `tfsdk:"suffix" json:"suffix,omitempty"`
						} `tfsdk:"match" json:"match,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"headers" json:"headers,omitempty"`
					Method *string `tfsdk:"method" json:"method,omitempty"`
					Path   *struct {
						Exact *string `tfsdk:"exact" json:"exact,omitempty"`
						Regex *string `tfsdk:"regex" json:"regex,omitempty"`
					} `tfsdk:"path" json:"path,omitempty"`
					Port            *int64  `tfsdk:"port" json:"port,omitempty"`
					Prefix          *string `tfsdk:"prefix" json:"prefix,omitempty"`
					QueryParameters *[]struct {
						Match *struct {
							Exact *string `tfsdk:"exact" json:"exact,omitempty"`
						} `tfsdk:"match" json:"match,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"query_parameters" json:"queryParameters,omitempty"`
					Scheme *string `tfsdk:"scheme" json:"scheme,omitempty"`
				} `tfsdk:"match" json:"match,omitempty"`
				RetryPolicy *struct {
					HttpRetryEvents *[]string `tfsdk:"http_retry_events" json:"httpRetryEvents,omitempty"`
					MaxRetries      *int64    `tfsdk:"max_retries" json:"maxRetries,omitempty"`
					PerRetryTimeout *struct {
						Unit  *string `tfsdk:"unit" json:"unit,omitempty"`
						Value *int64  `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"per_retry_timeout" json:"perRetryTimeout,omitempty"`
					TcpRetryEvents *[]string `tfsdk:"tcp_retry_events" json:"tcpRetryEvents,omitempty"`
				} `tfsdk:"retry_policy" json:"retryPolicy,omitempty"`
				Timeout *struct {
					Idle *struct {
						Unit  *string `tfsdk:"unit" json:"unit,omitempty"`
						Value *int64  `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"idle" json:"idle,omitempty"`
					PerRequest *struct {
						Unit  *string `tfsdk:"unit" json:"unit,omitempty"`
						Value *int64  `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"per_request" json:"perRequest,omitempty"`
				} `tfsdk:"timeout" json:"timeout,omitempty"`
			} `tfsdk:"http_route" json:"httpRoute,omitempty"`
			Name     *string `tfsdk:"name" json:"name,omitempty"`
			Priority *int64  `tfsdk:"priority" json:"priority,omitempty"`
			TcpRoute *struct {
				Action *struct {
					WeightedTargets *[]struct {
						Port           *int64  `tfsdk:"port" json:"port,omitempty"`
						VirtualNodeARN *string `tfsdk:"virtual_node_arn" json:"virtualNodeARN,omitempty"`
						VirtualNodeRef *struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"virtual_node_ref" json:"virtualNodeRef,omitempty"`
						Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
					} `tfsdk:"weighted_targets" json:"weightedTargets,omitempty"`
				} `tfsdk:"action" json:"action,omitempty"`
				Match *struct {
					Port *int64 `tfsdk:"port" json:"port,omitempty"`
				} `tfsdk:"match" json:"match,omitempty"`
				Timeout *struct {
					Idle *struct {
						Unit  *string `tfsdk:"unit" json:"unit,omitempty"`
						Value *int64  `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"idle" json:"idle,omitempty"`
				} `tfsdk:"timeout" json:"timeout,omitempty"`
			} `tfsdk:"tcp_route" json:"tcpRoute,omitempty"`
		} `tfsdk:"routes" json:"routes,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AppmeshK8SAwsVirtualRouterV1Beta2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_appmesh_k8s_aws_virtual_router_v1beta2_manifest"
}

func (r *AppmeshK8SAwsVirtualRouterV1Beta2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "VirtualRouter is the Schema for the virtualrouters API",
		MarkdownDescription: "VirtualRouter is the Schema for the virtualrouters API",
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
				Description:         "VirtualRouterSpec defines the desired state of VirtualRouter refers to https://docs.aws.amazon.com/app-mesh/latest/APIReference/API_VirtualRouterSpec.html",
				MarkdownDescription: "VirtualRouterSpec defines the desired state of VirtualRouter refers to https://docs.aws.amazon.com/app-mesh/latest/APIReference/API_VirtualRouterSpec.html",
				Attributes: map[string]schema.Attribute{
					"aws_name": schema.StringAttribute{
						Description:         "AWSName is the AppMesh VirtualRouter object's name. If unspecified or empty, it defaults to be '${name}_${namespace}' of k8s VirtualRouter",
						MarkdownDescription: "AWSName is the AppMesh VirtualRouter object's name. If unspecified or empty, it defaults to be '${name}_${namespace}' of k8s VirtualRouter",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"listeners": schema.ListNestedAttribute{
						Description:         "The listeners that the virtual router is expected to receive inbound traffic from",
						MarkdownDescription: "The listeners that the virtual router is expected to receive inbound traffic from",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"port_mapping": schema.SingleNestedAttribute{
									Description:         "The port mapping information for the listener.",
									MarkdownDescription: "The port mapping information for the listener.",
									Attributes: map[string]schema.Attribute{
										"port": schema.Int64Attribute{
											Description:         "The port used for the port mapping.",
											MarkdownDescription: "The port used for the port mapping.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(1),
												int64validator.AtMost(65535),
											},
										},

										"protocol": schema.StringAttribute{
											Description:         "The protocol used for the port mapping.",
											MarkdownDescription: "The protocol used for the port mapping.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("grpc", "http", "http2", "tcp"),
											},
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"mesh_ref": schema.SingleNestedAttribute{
						Description:         "A reference to k8s Mesh CR that this VirtualRouter belongs to. The admission controller populates it using Meshes's selector, and prevents users from setting this field.  Populated by the system. Read-only.",
						MarkdownDescription: "A reference to k8s Mesh CR that this VirtualRouter belongs to. The admission controller populates it using Meshes's selector, and prevents users from setting this field.  Populated by the system. Read-only.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name is the name of Mesh CR",
								MarkdownDescription: "Name is the name of Mesh CR",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"uid": schema.StringAttribute{
								Description:         "UID is the UID of Mesh CR",
								MarkdownDescription: "UID is the UID of Mesh CR",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"routes": schema.ListNestedAttribute{
						Description:         "The routes associated with VirtualRouter",
						MarkdownDescription: "The routes associated with VirtualRouter",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"grpc_route": schema.SingleNestedAttribute{
									Description:         "An object that represents the specification of a gRPC route.",
									MarkdownDescription: "An object that represents the specification of a gRPC route.",
									Attributes: map[string]schema.Attribute{
										"action": schema.SingleNestedAttribute{
											Description:         "An object that represents the action to take if a match is determined.",
											MarkdownDescription: "An object that represents the action to take if a match is determined.",
											Attributes: map[string]schema.Attribute{
												"weighted_targets": schema.ListNestedAttribute{
													Description:         "An object that represents the targets that traffic is routed to when a request matches the route.",
													MarkdownDescription: "An object that represents the targets that traffic is routed to when a request matches the route.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"port": schema.Int64Attribute{
																Description:         "Specifies the targeted port of the weighted object",
																MarkdownDescription: "Specifies the targeted port of the weighted object",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(0),
																},
															},

															"virtual_node_arn": schema.StringAttribute{
																Description:         "Amazon Resource Name to AppMesh VirtualNode object to associate with the weighted target. Exactly one of 'virtualNodeRef' or 'virtualNodeARN' must be specified.",
																MarkdownDescription: "Amazon Resource Name to AppMesh VirtualNode object to associate with the weighted target. Exactly one of 'virtualNodeRef' or 'virtualNodeARN' must be specified.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"virtual_node_ref": schema.SingleNestedAttribute{
																Description:         "Reference to Kubernetes VirtualNode CR in cluster to associate with the weighted target. Exactly one of 'virtualNodeRef' or 'virtualNodeARN' must be specified.",
																MarkdownDescription: "Reference to Kubernetes VirtualNode CR in cluster to associate with the weighted target. Exactly one of 'virtualNodeRef' or 'virtualNodeARN' must be specified.",
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "Name is the name of VirtualNode CR",
																		MarkdownDescription: "Name is the name of VirtualNode CR",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"namespace": schema.StringAttribute{
																		Description:         "Namespace is the namespace of VirtualNode CR. If unspecified, defaults to the referencing object's namespace",
																		MarkdownDescription: "Namespace is the namespace of VirtualNode CR. If unspecified, defaults to the referencing object's namespace",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"weight": schema.Int64Attribute{
																Description:         "The relative weight of the weighted target.",
																MarkdownDescription: "The relative weight of the weighted target.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(0),
																	int64validator.AtMost(100),
																},
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

										"match": schema.SingleNestedAttribute{
											Description:         "An object that represents the criteria for determining a request match.",
											MarkdownDescription: "An object that represents the criteria for determining a request match.",
											Attributes: map[string]schema.Attribute{
												"metadata": schema.ListNestedAttribute{
													Description:         "An object that represents the data to match from the request.",
													MarkdownDescription: "An object that represents the data to match from the request.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"invert": schema.BoolAttribute{
																Description:         "Specify True to match anything except the match criteria. The default value is False.",
																MarkdownDescription: "Specify True to match anything except the match criteria. The default value is False.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"match": schema.SingleNestedAttribute{
																Description:         "An object that represents the data to match from the request.",
																MarkdownDescription: "An object that represents the data to match from the request.",
																Attributes: map[string]schema.Attribute{
																	"exact": schema.StringAttribute{
																		Description:         "The value sent by the client must match the specified value exactly.",
																		MarkdownDescription: "The value sent by the client must match the specified value exactly.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.LengthAtLeast(1),
																			stringvalidator.LengthAtMost(255),
																		},
																	},

																	"prefix": schema.StringAttribute{
																		Description:         "The value sent by the client must begin with the specified characters.",
																		MarkdownDescription: "The value sent by the client must begin with the specified characters.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.LengthAtLeast(1),
																			stringvalidator.LengthAtMost(255),
																		},
																	},

																	"range": schema.SingleNestedAttribute{
																		Description:         "An object that represents the range of values to match on",
																		MarkdownDescription: "An object that represents the range of values to match on",
																		Attributes: map[string]schema.Attribute{
																			"end": schema.Int64Attribute{
																				Description:         "The end of the range.",
																				MarkdownDescription: "The end of the range.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"start": schema.Int64Attribute{
																				Description:         "The start of the range.",
																				MarkdownDescription: "The start of the range.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"regex": schema.StringAttribute{
																		Description:         "The value sent by the client must include the specified characters.",
																		MarkdownDescription: "The value sent by the client must include the specified characters.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.LengthAtLeast(1),
																			stringvalidator.LengthAtMost(255),
																		},
																	},

																	"suffix": schema.StringAttribute{
																		Description:         "The value sent by the client must end with the specified characters.",
																		MarkdownDescription: "The value sent by the client must end with the specified characters.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.LengthAtLeast(1),
																			stringvalidator.LengthAtMost(255),
																		},
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the route.",
																MarkdownDescription: "The name of the route.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																	stringvalidator.LengthAtMost(50),
																},
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"method_name": schema.StringAttribute{
													Description:         "The method name to match from the request. If you specify a name, you must also specify a serviceName.",
													MarkdownDescription: "The method name to match from the request. If you specify a name, you must also specify a serviceName.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
														stringvalidator.LengthAtMost(50),
													},
												},

												"port": schema.Int64Attribute{
													Description:         "Specifies the port to match requests with",
													MarkdownDescription: "Specifies the port to match requests with",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(0),
													},
												},

												"service_name": schema.StringAttribute{
													Description:         "The fully qualified domain name for the service to match from the request.",
													MarkdownDescription: "The fully qualified domain name for the service to match from the request.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"retry_policy": schema.SingleNestedAttribute{
											Description:         "An object that represents a retry policy.",
											MarkdownDescription: "An object that represents a retry policy.",
											Attributes: map[string]schema.Attribute{
												"grpc_retry_events": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"http_retry_events": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"max_retries": schema.Int64Attribute{
													Description:         "The maximum number of retry attempts.",
													MarkdownDescription: "The maximum number of retry attempts.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(0),
													},
												},

												"per_retry_timeout": schema.SingleNestedAttribute{
													Description:         "An object that represents a duration of time.",
													MarkdownDescription: "An object that represents a duration of time.",
													Attributes: map[string]schema.Attribute{
														"unit": schema.StringAttribute{
															Description:         "A unit of time.",
															MarkdownDescription: "A unit of time.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("s", "ms"),
															},
														},

														"value": schema.Int64Attribute{
															Description:         "A number of time units.",
															MarkdownDescription: "A number of time units.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(0),
															},
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},

												"tcp_retry_events": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
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

										"timeout": schema.SingleNestedAttribute{
											Description:         "An object that represents a grpc timeout.",
											MarkdownDescription: "An object that represents a grpc timeout.",
											Attributes: map[string]schema.Attribute{
												"idle": schema.SingleNestedAttribute{
													Description:         "An object that represents idle timeout duration.",
													MarkdownDescription: "An object that represents idle timeout duration.",
													Attributes: map[string]schema.Attribute{
														"unit": schema.StringAttribute{
															Description:         "A unit of time.",
															MarkdownDescription: "A unit of time.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("s", "ms"),
															},
														},

														"value": schema.Int64Attribute{
															Description:         "A number of time units.",
															MarkdownDescription: "A number of time units.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(0),
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"per_request": schema.SingleNestedAttribute{
													Description:         "An object that represents per request timeout duration.",
													MarkdownDescription: "An object that represents per request timeout duration.",
													Attributes: map[string]schema.Attribute{
														"unit": schema.StringAttribute{
															Description:         "A unit of time.",
															MarkdownDescription: "A unit of time.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("s", "ms"),
															},
														},

														"value": schema.Int64Attribute{
															Description:         "A number of time units.",
															MarkdownDescription: "A number of time units.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(0),
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
									Required: false,
									Optional: true,
									Computed: false,
								},

								"http2_route": schema.SingleNestedAttribute{
									Description:         "An object that represents the specification of an HTTP/2 route.",
									MarkdownDescription: "An object that represents the specification of an HTTP/2 route.",
									Attributes: map[string]schema.Attribute{
										"action": schema.SingleNestedAttribute{
											Description:         "An object that represents the action to take if a match is determined.",
											MarkdownDescription: "An object that represents the action to take if a match is determined.",
											Attributes: map[string]schema.Attribute{
												"weighted_targets": schema.ListNestedAttribute{
													Description:         "An object that represents the targets that traffic is routed to when a request matches the route.",
													MarkdownDescription: "An object that represents the targets that traffic is routed to when a request matches the route.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"port": schema.Int64Attribute{
																Description:         "Specifies the targeted port of the weighted object",
																MarkdownDescription: "Specifies the targeted port of the weighted object",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(0),
																},
															},

															"virtual_node_arn": schema.StringAttribute{
																Description:         "Amazon Resource Name to AppMesh VirtualNode object to associate with the weighted target. Exactly one of 'virtualNodeRef' or 'virtualNodeARN' must be specified.",
																MarkdownDescription: "Amazon Resource Name to AppMesh VirtualNode object to associate with the weighted target. Exactly one of 'virtualNodeRef' or 'virtualNodeARN' must be specified.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"virtual_node_ref": schema.SingleNestedAttribute{
																Description:         "Reference to Kubernetes VirtualNode CR in cluster to associate with the weighted target. Exactly one of 'virtualNodeRef' or 'virtualNodeARN' must be specified.",
																MarkdownDescription: "Reference to Kubernetes VirtualNode CR in cluster to associate with the weighted target. Exactly one of 'virtualNodeRef' or 'virtualNodeARN' must be specified.",
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "Name is the name of VirtualNode CR",
																		MarkdownDescription: "Name is the name of VirtualNode CR",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"namespace": schema.StringAttribute{
																		Description:         "Namespace is the namespace of VirtualNode CR. If unspecified, defaults to the referencing object's namespace",
																		MarkdownDescription: "Namespace is the namespace of VirtualNode CR. If unspecified, defaults to the referencing object's namespace",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"weight": schema.Int64Attribute{
																Description:         "The relative weight of the weighted target.",
																MarkdownDescription: "The relative weight of the weighted target.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(0),
																	int64validator.AtMost(100),
																},
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

										"match": schema.SingleNestedAttribute{
											Description:         "An object that represents the criteria for determining a request match.",
											MarkdownDescription: "An object that represents the criteria for determining a request match.",
											Attributes: map[string]schema.Attribute{
												"headers": schema.ListNestedAttribute{
													Description:         "An object that represents the client request headers to match on.",
													MarkdownDescription: "An object that represents the client request headers to match on.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"invert": schema.BoolAttribute{
																Description:         "Specify True to match anything except the match criteria. The default value is False.",
																MarkdownDescription: "Specify True to match anything except the match criteria. The default value is False.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"match": schema.SingleNestedAttribute{
																Description:         "The HeaderMatchMethod object.",
																MarkdownDescription: "The HeaderMatchMethod object.",
																Attributes: map[string]schema.Attribute{
																	"exact": schema.StringAttribute{
																		Description:         "The value sent by the client must match the specified value exactly.",
																		MarkdownDescription: "The value sent by the client must match the specified value exactly.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.LengthAtLeast(1),
																			stringvalidator.LengthAtMost(255),
																		},
																	},

																	"prefix": schema.StringAttribute{
																		Description:         "The value sent by the client must begin with the specified characters.",
																		MarkdownDescription: "The value sent by the client must begin with the specified characters.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.LengthAtLeast(1),
																			stringvalidator.LengthAtMost(255),
																		},
																	},

																	"range": schema.SingleNestedAttribute{
																		Description:         "An object that represents the range of values to match on.",
																		MarkdownDescription: "An object that represents the range of values to match on.",
																		Attributes: map[string]schema.Attribute{
																			"end": schema.Int64Attribute{
																				Description:         "The end of the range.",
																				MarkdownDescription: "The end of the range.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"start": schema.Int64Attribute{
																				Description:         "The start of the range.",
																				MarkdownDescription: "The start of the range.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"regex": schema.StringAttribute{
																		Description:         "The value sent by the client must include the specified characters.",
																		MarkdownDescription: "The value sent by the client must include the specified characters.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.LengthAtLeast(1),
																			stringvalidator.LengthAtMost(255),
																		},
																	},

																	"suffix": schema.StringAttribute{
																		Description:         "The value sent by the client must end with the specified characters.",
																		MarkdownDescription: "The value sent by the client must end with the specified characters.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.LengthAtLeast(1),
																			stringvalidator.LengthAtMost(255),
																		},
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": schema.StringAttribute{
																Description:         "A name for the HTTP header in the client request that will be matched on.",
																MarkdownDescription: "A name for the HTTP header in the client request that will be matched on.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																	stringvalidator.LengthAtMost(50),
																},
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"method": schema.StringAttribute{
													Description:         "The client request method to match on.",
													MarkdownDescription: "The client request method to match on.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("CONNECT", "DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT", "TRACE"),
													},
												},

												"path": schema.SingleNestedAttribute{
													Description:         "The client specified Path to match on.",
													MarkdownDescription: "The client specified Path to match on.",
													Attributes: map[string]schema.Attribute{
														"exact": schema.StringAttribute{
															Description:         "The value sent by the client must match the specified value exactly.",
															MarkdownDescription: "The value sent by the client must match the specified value exactly.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtLeast(1),
																stringvalidator.LengthAtMost(255),
															},
														},

														"regex": schema.StringAttribute{
															Description:         "The value sent by the client must end with the specified characters.",
															MarkdownDescription: "The value sent by the client must end with the specified characters.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtLeast(1),
																stringvalidator.LengthAtMost(255),
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"port": schema.Int64Attribute{
													Description:         "Specifies the port to match requests with",
													MarkdownDescription: "Specifies the port to match requests with",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(0),
													},
												},

												"prefix": schema.StringAttribute{
													Description:         "Specifies the prefix to match requests with",
													MarkdownDescription: "Specifies the prefix to match requests with",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"query_parameters": schema.ListNestedAttribute{
													Description:         "The client specified queryParameters to match on",
													MarkdownDescription: "The client specified queryParameters to match on",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"match": schema.SingleNestedAttribute{
																Description:         "The QueryMatchMethod object.",
																MarkdownDescription: "The QueryMatchMethod object.",
																Attributes: map[string]schema.Attribute{
																	"exact": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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

												"scheme": schema.StringAttribute{
													Description:         "The client request scheme to match on",
													MarkdownDescription: "The client request scheme to match on",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("http", "https"),
													},
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"retry_policy": schema.SingleNestedAttribute{
											Description:         "An object that represents a retry policy.",
											MarkdownDescription: "An object that represents a retry policy.",
											Attributes: map[string]schema.Attribute{
												"http_retry_events": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"max_retries": schema.Int64Attribute{
													Description:         "The maximum number of retry attempts.",
													MarkdownDescription: "The maximum number of retry attempts.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(0),
													},
												},

												"per_retry_timeout": schema.SingleNestedAttribute{
													Description:         "An object that represents a duration of time",
													MarkdownDescription: "An object that represents a duration of time",
													Attributes: map[string]schema.Attribute{
														"unit": schema.StringAttribute{
															Description:         "A unit of time.",
															MarkdownDescription: "A unit of time.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("s", "ms"),
															},
														},

														"value": schema.Int64Attribute{
															Description:         "A number of time units.",
															MarkdownDescription: "A number of time units.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(0),
															},
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},

												"tcp_retry_events": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
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

										"timeout": schema.SingleNestedAttribute{
											Description:         "An object that represents a http timeout.",
											MarkdownDescription: "An object that represents a http timeout.",
											Attributes: map[string]schema.Attribute{
												"idle": schema.SingleNestedAttribute{
													Description:         "An object that represents idle timeout duration.",
													MarkdownDescription: "An object that represents idle timeout duration.",
													Attributes: map[string]schema.Attribute{
														"unit": schema.StringAttribute{
															Description:         "A unit of time.",
															MarkdownDescription: "A unit of time.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("s", "ms"),
															},
														},

														"value": schema.Int64Attribute{
															Description:         "A number of time units.",
															MarkdownDescription: "A number of time units.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(0),
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"per_request": schema.SingleNestedAttribute{
													Description:         "An object that represents per request timeout duration.",
													MarkdownDescription: "An object that represents per request timeout duration.",
													Attributes: map[string]schema.Attribute{
														"unit": schema.StringAttribute{
															Description:         "A unit of time.",
															MarkdownDescription: "A unit of time.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("s", "ms"),
															},
														},

														"value": schema.Int64Attribute{
															Description:         "A number of time units.",
															MarkdownDescription: "A number of time units.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(0),
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
									Required: false,
									Optional: true,
									Computed: false,
								},

								"http_route": schema.SingleNestedAttribute{
									Description:         "An object that represents the specification of an HTTP route.",
									MarkdownDescription: "An object that represents the specification of an HTTP route.",
									Attributes: map[string]schema.Attribute{
										"action": schema.SingleNestedAttribute{
											Description:         "An object that represents the action to take if a match is determined.",
											MarkdownDescription: "An object that represents the action to take if a match is determined.",
											Attributes: map[string]schema.Attribute{
												"weighted_targets": schema.ListNestedAttribute{
													Description:         "An object that represents the targets that traffic is routed to when a request matches the route.",
													MarkdownDescription: "An object that represents the targets that traffic is routed to when a request matches the route.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"port": schema.Int64Attribute{
																Description:         "Specifies the targeted port of the weighted object",
																MarkdownDescription: "Specifies the targeted port of the weighted object",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(0),
																},
															},

															"virtual_node_arn": schema.StringAttribute{
																Description:         "Amazon Resource Name to AppMesh VirtualNode object to associate with the weighted target. Exactly one of 'virtualNodeRef' or 'virtualNodeARN' must be specified.",
																MarkdownDescription: "Amazon Resource Name to AppMesh VirtualNode object to associate with the weighted target. Exactly one of 'virtualNodeRef' or 'virtualNodeARN' must be specified.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"virtual_node_ref": schema.SingleNestedAttribute{
																Description:         "Reference to Kubernetes VirtualNode CR in cluster to associate with the weighted target. Exactly one of 'virtualNodeRef' or 'virtualNodeARN' must be specified.",
																MarkdownDescription: "Reference to Kubernetes VirtualNode CR in cluster to associate with the weighted target. Exactly one of 'virtualNodeRef' or 'virtualNodeARN' must be specified.",
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "Name is the name of VirtualNode CR",
																		MarkdownDescription: "Name is the name of VirtualNode CR",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"namespace": schema.StringAttribute{
																		Description:         "Namespace is the namespace of VirtualNode CR. If unspecified, defaults to the referencing object's namespace",
																		MarkdownDescription: "Namespace is the namespace of VirtualNode CR. If unspecified, defaults to the referencing object's namespace",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"weight": schema.Int64Attribute{
																Description:         "The relative weight of the weighted target.",
																MarkdownDescription: "The relative weight of the weighted target.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(0),
																	int64validator.AtMost(100),
																},
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

										"match": schema.SingleNestedAttribute{
											Description:         "An object that represents the criteria for determining a request match.",
											MarkdownDescription: "An object that represents the criteria for determining a request match.",
											Attributes: map[string]schema.Attribute{
												"headers": schema.ListNestedAttribute{
													Description:         "An object that represents the client request headers to match on.",
													MarkdownDescription: "An object that represents the client request headers to match on.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"invert": schema.BoolAttribute{
																Description:         "Specify True to match anything except the match criteria. The default value is False.",
																MarkdownDescription: "Specify True to match anything except the match criteria. The default value is False.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"match": schema.SingleNestedAttribute{
																Description:         "The HeaderMatchMethod object.",
																MarkdownDescription: "The HeaderMatchMethod object.",
																Attributes: map[string]schema.Attribute{
																	"exact": schema.StringAttribute{
																		Description:         "The value sent by the client must match the specified value exactly.",
																		MarkdownDescription: "The value sent by the client must match the specified value exactly.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.LengthAtLeast(1),
																			stringvalidator.LengthAtMost(255),
																		},
																	},

																	"prefix": schema.StringAttribute{
																		Description:         "The value sent by the client must begin with the specified characters.",
																		MarkdownDescription: "The value sent by the client must begin with the specified characters.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.LengthAtLeast(1),
																			stringvalidator.LengthAtMost(255),
																		},
																	},

																	"range": schema.SingleNestedAttribute{
																		Description:         "An object that represents the range of values to match on.",
																		MarkdownDescription: "An object that represents the range of values to match on.",
																		Attributes: map[string]schema.Attribute{
																			"end": schema.Int64Attribute{
																				Description:         "The end of the range.",
																				MarkdownDescription: "The end of the range.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"start": schema.Int64Attribute{
																				Description:         "The start of the range.",
																				MarkdownDescription: "The start of the range.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"regex": schema.StringAttribute{
																		Description:         "The value sent by the client must include the specified characters.",
																		MarkdownDescription: "The value sent by the client must include the specified characters.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.LengthAtLeast(1),
																			stringvalidator.LengthAtMost(255),
																		},
																	},

																	"suffix": schema.StringAttribute{
																		Description:         "The value sent by the client must end with the specified characters.",
																		MarkdownDescription: "The value sent by the client must end with the specified characters.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.LengthAtLeast(1),
																			stringvalidator.LengthAtMost(255),
																		},
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": schema.StringAttribute{
																Description:         "A name for the HTTP header in the client request that will be matched on.",
																MarkdownDescription: "A name for the HTTP header in the client request that will be matched on.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																	stringvalidator.LengthAtMost(50),
																},
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"method": schema.StringAttribute{
													Description:         "The client request method to match on.",
													MarkdownDescription: "The client request method to match on.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("CONNECT", "DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT", "TRACE"),
													},
												},

												"path": schema.SingleNestedAttribute{
													Description:         "The client specified Path to match on.",
													MarkdownDescription: "The client specified Path to match on.",
													Attributes: map[string]schema.Attribute{
														"exact": schema.StringAttribute{
															Description:         "The value sent by the client must match the specified value exactly.",
															MarkdownDescription: "The value sent by the client must match the specified value exactly.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtLeast(1),
																stringvalidator.LengthAtMost(255),
															},
														},

														"regex": schema.StringAttribute{
															Description:         "The value sent by the client must end with the specified characters.",
															MarkdownDescription: "The value sent by the client must end with the specified characters.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtLeast(1),
																stringvalidator.LengthAtMost(255),
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"port": schema.Int64Attribute{
													Description:         "Specifies the port to match requests with",
													MarkdownDescription: "Specifies the port to match requests with",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(0),
													},
												},

												"prefix": schema.StringAttribute{
													Description:         "Specifies the prefix to match requests with",
													MarkdownDescription: "Specifies the prefix to match requests with",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"query_parameters": schema.ListNestedAttribute{
													Description:         "The client specified queryParameters to match on",
													MarkdownDescription: "The client specified queryParameters to match on",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"match": schema.SingleNestedAttribute{
																Description:         "The QueryMatchMethod object.",
																MarkdownDescription: "The QueryMatchMethod object.",
																Attributes: map[string]schema.Attribute{
																	"exact": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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

												"scheme": schema.StringAttribute{
													Description:         "The client request scheme to match on",
													MarkdownDescription: "The client request scheme to match on",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("http", "https"),
													},
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"retry_policy": schema.SingleNestedAttribute{
											Description:         "An object that represents a retry policy.",
											MarkdownDescription: "An object that represents a retry policy.",
											Attributes: map[string]schema.Attribute{
												"http_retry_events": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"max_retries": schema.Int64Attribute{
													Description:         "The maximum number of retry attempts.",
													MarkdownDescription: "The maximum number of retry attempts.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(0),
													},
												},

												"per_retry_timeout": schema.SingleNestedAttribute{
													Description:         "An object that represents a duration of time",
													MarkdownDescription: "An object that represents a duration of time",
													Attributes: map[string]schema.Attribute{
														"unit": schema.StringAttribute{
															Description:         "A unit of time.",
															MarkdownDescription: "A unit of time.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("s", "ms"),
															},
														},

														"value": schema.Int64Attribute{
															Description:         "A number of time units.",
															MarkdownDescription: "A number of time units.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(0),
															},
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},

												"tcp_retry_events": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
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

										"timeout": schema.SingleNestedAttribute{
											Description:         "An object that represents a http timeout.",
											MarkdownDescription: "An object that represents a http timeout.",
											Attributes: map[string]schema.Attribute{
												"idle": schema.SingleNestedAttribute{
													Description:         "An object that represents idle timeout duration.",
													MarkdownDescription: "An object that represents idle timeout duration.",
													Attributes: map[string]schema.Attribute{
														"unit": schema.StringAttribute{
															Description:         "A unit of time.",
															MarkdownDescription: "A unit of time.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("s", "ms"),
															},
														},

														"value": schema.Int64Attribute{
															Description:         "A number of time units.",
															MarkdownDescription: "A number of time units.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(0),
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"per_request": schema.SingleNestedAttribute{
													Description:         "An object that represents per request timeout duration.",
													MarkdownDescription: "An object that represents per request timeout duration.",
													Attributes: map[string]schema.Attribute{
														"unit": schema.StringAttribute{
															Description:         "A unit of time.",
															MarkdownDescription: "A unit of time.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("s", "ms"),
															},
														},

														"value": schema.Int64Attribute{
															Description:         "A number of time units.",
															MarkdownDescription: "A number of time units.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(0),
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
									Required: false,
									Optional: true,
									Computed: false,
								},

								"name": schema.StringAttribute{
									Description:         "Route's name",
									MarkdownDescription: "Route's name",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"priority": schema.Int64Attribute{
									Description:         "The priority for the route.",
									MarkdownDescription: "The priority for the route.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.Int64{
										int64validator.AtLeast(0),
										int64validator.AtMost(1000),
									},
								},

								"tcp_route": schema.SingleNestedAttribute{
									Description:         "An object that represents the specification of a TCP route.",
									MarkdownDescription: "An object that represents the specification of a TCP route.",
									Attributes: map[string]schema.Attribute{
										"action": schema.SingleNestedAttribute{
											Description:         "The action to take if a match is determined.",
											MarkdownDescription: "The action to take if a match is determined.",
											Attributes: map[string]schema.Attribute{
												"weighted_targets": schema.ListNestedAttribute{
													Description:         "An object that represents the targets that traffic is routed to when a request matches the route.",
													MarkdownDescription: "An object that represents the targets that traffic is routed to when a request matches the route.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"port": schema.Int64Attribute{
																Description:         "Specifies the targeted port of the weighted object",
																MarkdownDescription: "Specifies the targeted port of the weighted object",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(0),
																},
															},

															"virtual_node_arn": schema.StringAttribute{
																Description:         "Amazon Resource Name to AppMesh VirtualNode object to associate with the weighted target. Exactly one of 'virtualNodeRef' or 'virtualNodeARN' must be specified.",
																MarkdownDescription: "Amazon Resource Name to AppMesh VirtualNode object to associate with the weighted target. Exactly one of 'virtualNodeRef' or 'virtualNodeARN' must be specified.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"virtual_node_ref": schema.SingleNestedAttribute{
																Description:         "Reference to Kubernetes VirtualNode CR in cluster to associate with the weighted target. Exactly one of 'virtualNodeRef' or 'virtualNodeARN' must be specified.",
																MarkdownDescription: "Reference to Kubernetes VirtualNode CR in cluster to associate with the weighted target. Exactly one of 'virtualNodeRef' or 'virtualNodeARN' must be specified.",
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "Name is the name of VirtualNode CR",
																		MarkdownDescription: "Name is the name of VirtualNode CR",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"namespace": schema.StringAttribute{
																		Description:         "Namespace is the namespace of VirtualNode CR. If unspecified, defaults to the referencing object's namespace",
																		MarkdownDescription: "Namespace is the namespace of VirtualNode CR. If unspecified, defaults to the referencing object's namespace",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"weight": schema.Int64Attribute{
																Description:         "The relative weight of the weighted target.",
																MarkdownDescription: "The relative weight of the weighted target.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(0),
																	int64validator.AtMost(100),
																},
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

										"match": schema.SingleNestedAttribute{
											Description:         "An object that represents the criteria for determining a request match.",
											MarkdownDescription: "An object that represents the criteria for determining a request match.",
											Attributes: map[string]schema.Attribute{
												"port": schema.Int64Attribute{
													Description:         "Specifies the port to match requests with",
													MarkdownDescription: "Specifies the port to match requests with",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(0),
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"timeout": schema.SingleNestedAttribute{
											Description:         "An object that represents a tcp timeout.",
											MarkdownDescription: "An object that represents a tcp timeout.",
											Attributes: map[string]schema.Attribute{
												"idle": schema.SingleNestedAttribute{
													Description:         "An object that represents idle timeout duration.",
													MarkdownDescription: "An object that represents idle timeout duration.",
													Attributes: map[string]schema.Attribute{
														"unit": schema.StringAttribute{
															Description:         "A unit of time.",
															MarkdownDescription: "A unit of time.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("s", "ms"),
															},
														},

														"value": schema.Int64Attribute{
															Description:         "A number of time units.",
															MarkdownDescription: "A number of time units.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(0),
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
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *AppmeshK8SAwsVirtualRouterV1Beta2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_appmesh_k8s_aws_virtual_router_v1beta2_manifest")

	var model AppmeshK8SAwsVirtualRouterV1Beta2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("appmesh.k8s.aws/v1beta2")
	model.Kind = pointer.String("VirtualRouter")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
