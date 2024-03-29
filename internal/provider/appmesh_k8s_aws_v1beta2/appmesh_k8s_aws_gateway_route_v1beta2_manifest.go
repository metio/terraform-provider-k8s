/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package appmesh_k8s_aws_v1beta2

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
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &AppmeshK8SAwsGatewayRouteV1Beta2Manifest{}
)

func NewAppmeshK8SAwsGatewayRouteV1Beta2Manifest() datasource.DataSource {
	return &AppmeshK8SAwsGatewayRouteV1Beta2Manifest{}
}

type AppmeshK8SAwsGatewayRouteV1Beta2Manifest struct{}

type AppmeshK8SAwsGatewayRouteV1Beta2ManifestData struct {
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
		AwsName   *string `tfsdk:"aws_name" json:"awsName,omitempty"`
		GrpcRoute *struct {
			Action *struct {
				Rewrite *struct {
					Hostname *struct {
						DefaultTargetHostname *string `tfsdk:"default_target_hostname" json:"defaultTargetHostname,omitempty"`
					} `tfsdk:"hostname" json:"hostname,omitempty"`
				} `tfsdk:"rewrite" json:"rewrite,omitempty"`
				Target *struct {
					Port           *int64 `tfsdk:"port" json:"port,omitempty"`
					VirtualService *struct {
						VirtualServiceARN *string `tfsdk:"virtual_service_arn" json:"virtualServiceARN,omitempty"`
						VirtualServiceRef *struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"virtual_service_ref" json:"virtualServiceRef,omitempty"`
					} `tfsdk:"virtual_service" json:"virtualService,omitempty"`
				} `tfsdk:"target" json:"target,omitempty"`
			} `tfsdk:"action" json:"action,omitempty"`
			Match *struct {
				Hostname *struct {
					Exact  *string `tfsdk:"exact" json:"exact,omitempty"`
					Suffix *string `tfsdk:"suffix" json:"suffix,omitempty"`
				} `tfsdk:"hostname" json:"hostname,omitempty"`
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
				Port        *int64  `tfsdk:"port" json:"port,omitempty"`
				ServiceName *string `tfsdk:"service_name" json:"serviceName,omitempty"`
			} `tfsdk:"match" json:"match,omitempty"`
		} `tfsdk:"grpc_route" json:"grpcRoute,omitempty"`
		Http2Route *struct {
			Action *struct {
				Rewrite *struct {
					Hostname *struct {
						DefaultTargetHostname *string `tfsdk:"default_target_hostname" json:"defaultTargetHostname,omitempty"`
					} `tfsdk:"hostname" json:"hostname,omitempty"`
					Path *struct {
						Exact *string `tfsdk:"exact" json:"exact,omitempty"`
					} `tfsdk:"path" json:"path,omitempty"`
					Prefix *struct {
						DefaultPrefix *string `tfsdk:"default_prefix" json:"defaultPrefix,omitempty"`
						Value         *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"prefix" json:"prefix,omitempty"`
				} `tfsdk:"rewrite" json:"rewrite,omitempty"`
				Target *struct {
					Port           *int64 `tfsdk:"port" json:"port,omitempty"`
					VirtualService *struct {
						VirtualServiceARN *string `tfsdk:"virtual_service_arn" json:"virtualServiceARN,omitempty"`
						VirtualServiceRef *struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"virtual_service_ref" json:"virtualServiceRef,omitempty"`
					} `tfsdk:"virtual_service" json:"virtualService,omitempty"`
				} `tfsdk:"target" json:"target,omitempty"`
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
				Hostname *struct {
					Exact  *string `tfsdk:"exact" json:"exact,omitempty"`
					Suffix *string `tfsdk:"suffix" json:"suffix,omitempty"`
				} `tfsdk:"hostname" json:"hostname,omitempty"`
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
			} `tfsdk:"match" json:"match,omitempty"`
		} `tfsdk:"http2_route" json:"http2Route,omitempty"`
		HttpRoute *struct {
			Action *struct {
				Rewrite *struct {
					Hostname *struct {
						DefaultTargetHostname *string `tfsdk:"default_target_hostname" json:"defaultTargetHostname,omitempty"`
					} `tfsdk:"hostname" json:"hostname,omitempty"`
					Path *struct {
						Exact *string `tfsdk:"exact" json:"exact,omitempty"`
					} `tfsdk:"path" json:"path,omitempty"`
					Prefix *struct {
						DefaultPrefix *string `tfsdk:"default_prefix" json:"defaultPrefix,omitempty"`
						Value         *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"prefix" json:"prefix,omitempty"`
				} `tfsdk:"rewrite" json:"rewrite,omitempty"`
				Target *struct {
					Port           *int64 `tfsdk:"port" json:"port,omitempty"`
					VirtualService *struct {
						VirtualServiceARN *string `tfsdk:"virtual_service_arn" json:"virtualServiceARN,omitempty"`
						VirtualServiceRef *struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"virtual_service_ref" json:"virtualServiceRef,omitempty"`
					} `tfsdk:"virtual_service" json:"virtualService,omitempty"`
				} `tfsdk:"target" json:"target,omitempty"`
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
				Hostname *struct {
					Exact  *string `tfsdk:"exact" json:"exact,omitempty"`
					Suffix *string `tfsdk:"suffix" json:"suffix,omitempty"`
				} `tfsdk:"hostname" json:"hostname,omitempty"`
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
			} `tfsdk:"match" json:"match,omitempty"`
		} `tfsdk:"http_route" json:"httpRoute,omitempty"`
		MeshRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
			Uid  *string `tfsdk:"uid" json:"uid,omitempty"`
		} `tfsdk:"mesh_ref" json:"meshRef,omitempty"`
		Priority          *int64 `tfsdk:"priority" json:"priority,omitempty"`
		VirtualGatewayRef *struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Uid       *string `tfsdk:"uid" json:"uid,omitempty"`
		} `tfsdk:"virtual_gateway_ref" json:"virtualGatewayRef,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AppmeshK8SAwsGatewayRouteV1Beta2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_appmesh_k8s_aws_gateway_route_v1beta2_manifest"
}

func (r *AppmeshK8SAwsGatewayRouteV1Beta2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "GatewayRoute is the Schema for the gatewayroutes API",
		MarkdownDescription: "GatewayRoute is the Schema for the gatewayroutes API",
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
				Description:         "GatewayRouteSpec defines the desired state of GatewayRoute refers to https://docs.aws.amazon.com/app-mesh/latest/userguide/virtual_gateways.html",
				MarkdownDescription: "GatewayRouteSpec defines the desired state of GatewayRoute refers to https://docs.aws.amazon.com/app-mesh/latest/userguide/virtual_gateways.html",
				Attributes: map[string]schema.Attribute{
					"aws_name": schema.StringAttribute{
						Description:         "AWSName is the AppMesh GatewayRoute object's name. If unspecified or empty, it defaults to be '${name}_${namespace}' of k8s GatewayRoute",
						MarkdownDescription: "AWSName is the AppMesh GatewayRoute object's name. If unspecified or empty, it defaults to be '${name}_${namespace}' of k8s GatewayRoute",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"grpc_route": schema.SingleNestedAttribute{
						Description:         "An object that represents the specification of a gRPC gatewayRoute.",
						MarkdownDescription: "An object that represents the specification of a gRPC gatewayRoute.",
						Attributes: map[string]schema.Attribute{
							"action": schema.SingleNestedAttribute{
								Description:         "An object that represents the action to take if a match is determined.",
								MarkdownDescription: "An object that represents the action to take if a match is determined.",
								Attributes: map[string]schema.Attribute{
									"rewrite": schema.SingleNestedAttribute{
										Description:         "GrpcGatewayRouteRewrite refers to https://docs.aws.amazon.com/app-mesh/latest/APIReference/API_GrpcGatewayRouteRewrite.html",
										MarkdownDescription: "GrpcGatewayRouteRewrite refers to https://docs.aws.amazon.com/app-mesh/latest/APIReference/API_GrpcGatewayRouteRewrite.html",
										Attributes: map[string]schema.Attribute{
											"hostname": schema.SingleNestedAttribute{
												Description:         "GatewayRouteHostnameRewrite refers to https://docs.aws.amazon.com/app-mesh/latest/APIReference/API_GatewayRouteHostnameRewrite.html ENABLE or DISABLE default behavior for Hostname rewrite",
												MarkdownDescription: "GatewayRouteHostnameRewrite refers to https://docs.aws.amazon.com/app-mesh/latest/APIReference/API_GatewayRouteHostnameRewrite.html ENABLE or DISABLE default behavior for Hostname rewrite",
												Attributes: map[string]schema.Attribute{
													"default_target_hostname": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("ENABLED", "DISABLED"),
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

									"target": schema.SingleNestedAttribute{
										Description:         "An object that represents the target that traffic is routed to when a request matches the route.",
										MarkdownDescription: "An object that represents the target that traffic is routed to when a request matches the route.",
										Attributes: map[string]schema.Attribute{
											"port": schema.Int64Attribute{
												Description:         "Specifies the port of the gateway route target",
												MarkdownDescription: "Specifies the port of the gateway route target",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"virtual_service": schema.SingleNestedAttribute{
												Description:         "The virtual service to associate with the gateway route target.",
												MarkdownDescription: "The virtual service to associate with the gateway route target.",
												Attributes: map[string]schema.Attribute{
													"virtual_service_arn": schema.StringAttribute{
														Description:         "Amazon Resource Name to AppMesh VirtualService object to associate with the gateway route virtual service target. Exactly one of 'virtualServiceRef' or 'virtualServiceARN' must be specified.",
														MarkdownDescription: "Amazon Resource Name to AppMesh VirtualService object to associate with the gateway route virtual service target. Exactly one of 'virtualServiceRef' or 'virtualServiceARN' must be specified.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"virtual_service_ref": schema.SingleNestedAttribute{
														Description:         "Reference to Kubernetes VirtualService CR in cluster to associate with the gateway route virtual service target. Exactly one of 'virtualServiceRef' or 'virtualServiceARN' must be specified.",
														MarkdownDescription: "Reference to Kubernetes VirtualService CR in cluster to associate with the gateway route virtual service target. Exactly one of 'virtualServiceRef' or 'virtualServiceARN' must be specified.",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name is the name of VirtualService CR",
																MarkdownDescription: "Name is the name of VirtualService CR",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace is the namespace of VirtualService CR. If unspecified, defaults to the referencing object's namespace",
																MarkdownDescription: "Namespace is the namespace of VirtualService CR. If unspecified, defaults to the referencing object's namespace",
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
								Required: true,
								Optional: false,
								Computed: false,
							},

							"match": schema.SingleNestedAttribute{
								Description:         "An object that represents the criteria for determining a request match.",
								MarkdownDescription: "An object that represents the criteria for determining a request match.",
								Attributes: map[string]schema.Attribute{
									"hostname": schema.SingleNestedAttribute{
										Description:         "The client specified Hostname to match on.",
										MarkdownDescription: "The client specified Hostname to match on.",
										Attributes: map[string]schema.Attribute{
											"exact": schema.StringAttribute{
												Description:         "The value sent by the client must match the specified value exactly.",
												MarkdownDescription: "The value sent by the client must match the specified value exactly.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
													stringvalidator.LengthAtMost(253),
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
													stringvalidator.LengthAtMost(253),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

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

									"port": schema.Int64Attribute{
										Description:         "Specifies the port the request to be matched on",
										MarkdownDescription: "Specifies the port the request to be matched on",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},

									"service_name": schema.StringAttribute{
										Description:         "Either ServiceName or Hostname must be specified. Both are allowed as well The fully qualified domain name for the service to match from the request.",
										MarkdownDescription: "Either ServiceName or Hostname must be specified. Both are allowed as well The fully qualified domain name for the service to match from the request.",
										Required:            false,
										Optional:            true,
										Computed:            false,
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

					"http2_route": schema.SingleNestedAttribute{
						Description:         "An object that represents the specification of an HTTP/2 gatewayRoute.",
						MarkdownDescription: "An object that represents the specification of an HTTP/2 gatewayRoute.",
						Attributes: map[string]schema.Attribute{
							"action": schema.SingleNestedAttribute{
								Description:         "An object that represents the action to take if a match is determined.",
								MarkdownDescription: "An object that represents the action to take if a match is determined.",
								Attributes: map[string]schema.Attribute{
									"rewrite": schema.SingleNestedAttribute{
										Description:         "HTTPGatewayRouteRewrite refers to https://docs.aws.amazon.com/app-mesh/latest/APIReference/API_HttpGatewayRouteRewrite.html",
										MarkdownDescription: "HTTPGatewayRouteRewrite refers to https://docs.aws.amazon.com/app-mesh/latest/APIReference/API_HttpGatewayRouteRewrite.html",
										Attributes: map[string]schema.Attribute{
											"hostname": schema.SingleNestedAttribute{
												Description:         "GatewayRouteHostnameRewrite refers to https://docs.aws.amazon.com/app-mesh/latest/APIReference/API_GatewayRouteHostnameRewrite.html ENABLE or DISABLE default behavior for Hostname rewrite",
												MarkdownDescription: "GatewayRouteHostnameRewrite refers to https://docs.aws.amazon.com/app-mesh/latest/APIReference/API_GatewayRouteHostnameRewrite.html ENABLE or DISABLE default behavior for Hostname rewrite",
												Attributes: map[string]schema.Attribute{
													"default_target_hostname": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("ENABLED", "DISABLED"),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"path": schema.SingleNestedAttribute{
												Description:         "GatewayRoutePathRewrite refers to https://docs.aws.amazon.com/app-mesh/latest/APIReference/API_HttpGatewayRoutePathRewrite.html",
												MarkdownDescription: "GatewayRoutePathRewrite refers to https://docs.aws.amazon.com/app-mesh/latest/APIReference/API_HttpGatewayRoutePathRewrite.html",
												Attributes: map[string]schema.Attribute{
													"exact": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
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

											"prefix": schema.SingleNestedAttribute{
												Description:         "GatewayRoutePrefixRewrite refers to https://docs.aws.amazon.com/app-mesh/latest/APIReference/API_HttpGatewayRoutePrefixRewrite.html",
												MarkdownDescription: "GatewayRoutePrefixRewrite refers to https://docs.aws.amazon.com/app-mesh/latest/APIReference/API_HttpGatewayRoutePrefixRewrite.html",
												Attributes: map[string]schema.Attribute{
													"default_prefix": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("ENABLED", "DISABLED"),
														},
													},

													"value": schema.StringAttribute{
														Description:         "When DefaultPrefix is specified, Value cannot be set",
														MarkdownDescription: "When DefaultPrefix is specified, Value cannot be set",
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"target": schema.SingleNestedAttribute{
										Description:         "An object that represents the target that traffic is routed to when a request matches the route.",
										MarkdownDescription: "An object that represents the target that traffic is routed to when a request matches the route.",
										Attributes: map[string]schema.Attribute{
											"port": schema.Int64Attribute{
												Description:         "Specifies the port of the gateway route target",
												MarkdownDescription: "Specifies the port of the gateway route target",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"virtual_service": schema.SingleNestedAttribute{
												Description:         "The virtual service to associate with the gateway route target.",
												MarkdownDescription: "The virtual service to associate with the gateway route target.",
												Attributes: map[string]schema.Attribute{
													"virtual_service_arn": schema.StringAttribute{
														Description:         "Amazon Resource Name to AppMesh VirtualService object to associate with the gateway route virtual service target. Exactly one of 'virtualServiceRef' or 'virtualServiceARN' must be specified.",
														MarkdownDescription: "Amazon Resource Name to AppMesh VirtualService object to associate with the gateway route virtual service target. Exactly one of 'virtualServiceRef' or 'virtualServiceARN' must be specified.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"virtual_service_ref": schema.SingleNestedAttribute{
														Description:         "Reference to Kubernetes VirtualService CR in cluster to associate with the gateway route virtual service target. Exactly one of 'virtualServiceRef' or 'virtualServiceARN' must be specified.",
														MarkdownDescription: "Reference to Kubernetes VirtualService CR in cluster to associate with the gateway route virtual service target. Exactly one of 'virtualServiceRef' or 'virtualServiceARN' must be specified.",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name is the name of VirtualService CR",
																MarkdownDescription: "Name is the name of VirtualService CR",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace is the namespace of VirtualService CR. If unspecified, defaults to the referencing object's namespace",
																MarkdownDescription: "Namespace is the namespace of VirtualService CR. If unspecified, defaults to the referencing object's namespace",
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

									"hostname": schema.SingleNestedAttribute{
										Description:         "The client specified Hostname to match on.",
										MarkdownDescription: "The client specified Hostname to match on.",
										Attributes: map[string]schema.Attribute{
											"exact": schema.StringAttribute{
												Description:         "The value sent by the client must match the specified value exactly.",
												MarkdownDescription: "The value sent by the client must match the specified value exactly.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
													stringvalidator.LengthAtMost(253),
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
													stringvalidator.LengthAtMost(253),
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
										Description:         "Specified path of the request to be matched on",
										MarkdownDescription: "Specified path of the request to be matched on",
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
										Description:         "Specifies the port the request to be matched on",
										MarkdownDescription: "Specifies the port the request to be matched on",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},

									"prefix": schema.StringAttribute{
										Description:         "Either Prefix or Hostname must be specified. Both are allowed as well. Specifies the prefix to match requests with",
										MarkdownDescription: "Either Prefix or Hostname must be specified. Both are allowed as well. Specifies the prefix to match requests with",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"query_parameters": schema.ListNestedAttribute{
										Description:         "Client specified query parameters to match on",
										MarkdownDescription: "Client specified query parameters to match on",
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

					"http_route": schema.SingleNestedAttribute{
						Description:         "An object that represents the specification of an HTTP gatewayRoute.",
						MarkdownDescription: "An object that represents the specification of an HTTP gatewayRoute.",
						Attributes: map[string]schema.Attribute{
							"action": schema.SingleNestedAttribute{
								Description:         "An object that represents the action to take if a match is determined.",
								MarkdownDescription: "An object that represents the action to take if a match is determined.",
								Attributes: map[string]schema.Attribute{
									"rewrite": schema.SingleNestedAttribute{
										Description:         "HTTPGatewayRouteRewrite refers to https://docs.aws.amazon.com/app-mesh/latest/APIReference/API_HttpGatewayRouteRewrite.html",
										MarkdownDescription: "HTTPGatewayRouteRewrite refers to https://docs.aws.amazon.com/app-mesh/latest/APIReference/API_HttpGatewayRouteRewrite.html",
										Attributes: map[string]schema.Attribute{
											"hostname": schema.SingleNestedAttribute{
												Description:         "GatewayRouteHostnameRewrite refers to https://docs.aws.amazon.com/app-mesh/latest/APIReference/API_GatewayRouteHostnameRewrite.html ENABLE or DISABLE default behavior for Hostname rewrite",
												MarkdownDescription: "GatewayRouteHostnameRewrite refers to https://docs.aws.amazon.com/app-mesh/latest/APIReference/API_GatewayRouteHostnameRewrite.html ENABLE or DISABLE default behavior for Hostname rewrite",
												Attributes: map[string]schema.Attribute{
													"default_target_hostname": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("ENABLED", "DISABLED"),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"path": schema.SingleNestedAttribute{
												Description:         "GatewayRoutePathRewrite refers to https://docs.aws.amazon.com/app-mesh/latest/APIReference/API_HttpGatewayRoutePathRewrite.html",
												MarkdownDescription: "GatewayRoutePathRewrite refers to https://docs.aws.amazon.com/app-mesh/latest/APIReference/API_HttpGatewayRoutePathRewrite.html",
												Attributes: map[string]schema.Attribute{
													"exact": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
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

											"prefix": schema.SingleNestedAttribute{
												Description:         "GatewayRoutePrefixRewrite refers to https://docs.aws.amazon.com/app-mesh/latest/APIReference/API_HttpGatewayRoutePrefixRewrite.html",
												MarkdownDescription: "GatewayRoutePrefixRewrite refers to https://docs.aws.amazon.com/app-mesh/latest/APIReference/API_HttpGatewayRoutePrefixRewrite.html",
												Attributes: map[string]schema.Attribute{
													"default_prefix": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("ENABLED", "DISABLED"),
														},
													},

													"value": schema.StringAttribute{
														Description:         "When DefaultPrefix is specified, Value cannot be set",
														MarkdownDescription: "When DefaultPrefix is specified, Value cannot be set",
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"target": schema.SingleNestedAttribute{
										Description:         "An object that represents the target that traffic is routed to when a request matches the route.",
										MarkdownDescription: "An object that represents the target that traffic is routed to when a request matches the route.",
										Attributes: map[string]schema.Attribute{
											"port": schema.Int64Attribute{
												Description:         "Specifies the port of the gateway route target",
												MarkdownDescription: "Specifies the port of the gateway route target",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"virtual_service": schema.SingleNestedAttribute{
												Description:         "The virtual service to associate with the gateway route target.",
												MarkdownDescription: "The virtual service to associate with the gateway route target.",
												Attributes: map[string]schema.Attribute{
													"virtual_service_arn": schema.StringAttribute{
														Description:         "Amazon Resource Name to AppMesh VirtualService object to associate with the gateway route virtual service target. Exactly one of 'virtualServiceRef' or 'virtualServiceARN' must be specified.",
														MarkdownDescription: "Amazon Resource Name to AppMesh VirtualService object to associate with the gateway route virtual service target. Exactly one of 'virtualServiceRef' or 'virtualServiceARN' must be specified.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"virtual_service_ref": schema.SingleNestedAttribute{
														Description:         "Reference to Kubernetes VirtualService CR in cluster to associate with the gateway route virtual service target. Exactly one of 'virtualServiceRef' or 'virtualServiceARN' must be specified.",
														MarkdownDescription: "Reference to Kubernetes VirtualService CR in cluster to associate with the gateway route virtual service target. Exactly one of 'virtualServiceRef' or 'virtualServiceARN' must be specified.",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name is the name of VirtualService CR",
																MarkdownDescription: "Name is the name of VirtualService CR",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace is the namespace of VirtualService CR. If unspecified, defaults to the referencing object's namespace",
																MarkdownDescription: "Namespace is the namespace of VirtualService CR. If unspecified, defaults to the referencing object's namespace",
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

									"hostname": schema.SingleNestedAttribute{
										Description:         "The client specified Hostname to match on.",
										MarkdownDescription: "The client specified Hostname to match on.",
										Attributes: map[string]schema.Attribute{
											"exact": schema.StringAttribute{
												Description:         "The value sent by the client must match the specified value exactly.",
												MarkdownDescription: "The value sent by the client must match the specified value exactly.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
													stringvalidator.LengthAtMost(253),
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
													stringvalidator.LengthAtMost(253),
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
										Description:         "Specified path of the request to be matched on",
										MarkdownDescription: "Specified path of the request to be matched on",
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
										Description:         "Specifies the port the request to be matched on",
										MarkdownDescription: "Specifies the port the request to be matched on",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},

									"prefix": schema.StringAttribute{
										Description:         "Either Prefix or Hostname must be specified. Both are allowed as well. Specifies the prefix to match requests with",
										MarkdownDescription: "Either Prefix or Hostname must be specified. Both are allowed as well. Specifies the prefix to match requests with",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"query_parameters": schema.ListNestedAttribute{
										Description:         "Client specified query parameters to match on",
										MarkdownDescription: "Client specified query parameters to match on",
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

					"mesh_ref": schema.SingleNestedAttribute{
						Description:         "A reference to k8s Mesh CR that this GatewayRoute belongs to. The admission controller populates it using Meshes's selector, and prevents users from setting this field.  Populated by the system. Read-only.",
						MarkdownDescription: "A reference to k8s Mesh CR that this GatewayRoute belongs to. The admission controller populates it using Meshes's selector, and prevents users from setting this field.  Populated by the system. Read-only.",
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

					"priority": schema.Int64Attribute{
						Description:         "Priority for the gatewayroute. Default Priority is 1000 which is lowest priority",
						MarkdownDescription: "Priority for the gatewayroute. Default Priority is 1000 which is lowest priority",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
							int64validator.AtMost(1000),
						},
					},

					"virtual_gateway_ref": schema.SingleNestedAttribute{
						Description:         "A reference to k8s VirtualGateway CR that this GatewayRoute belongs to. The admission controller populates it using VirtualGateway's selector, and prevents users from setting this field.  Populated by the system. Read-only.",
						MarkdownDescription: "A reference to k8s VirtualGateway CR that this GatewayRoute belongs to. The admission controller populates it using VirtualGateway's selector, and prevents users from setting this field.  Populated by the system. Read-only.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name is the name of VirtualGateway CR",
								MarkdownDescription: "Name is the name of VirtualGateway CR",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace is the namespace of VirtualGateway CR. If unspecified, defaults to the referencing object's namespace",
								MarkdownDescription: "Namespace is the namespace of VirtualGateway CR. If unspecified, defaults to the referencing object's namespace",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"uid": schema.StringAttribute{
								Description:         "UID is the UID of VirtualGateway CR",
								MarkdownDescription: "UID is the UID of VirtualGateway CR",
								Required:            true,
								Optional:            false,
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
	}
}

func (r *AppmeshK8SAwsGatewayRouteV1Beta2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_appmesh_k8s_aws_gateway_route_v1beta2_manifest")

	var model AppmeshK8SAwsGatewayRouteV1Beta2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("appmesh.k8s.aws/v1beta2")
	model.Kind = pointer.String("GatewayRoute")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
