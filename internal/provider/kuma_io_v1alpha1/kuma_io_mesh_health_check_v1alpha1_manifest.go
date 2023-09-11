/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kuma_io_v1alpha1

import (
	"context"
	"fmt"
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
	_ datasource.DataSource = &KumaIoMeshHealthCheckV1Alpha1Manifest{}
)

func NewKumaIoMeshHealthCheckV1Alpha1Manifest() datasource.DataSource {
	return &KumaIoMeshHealthCheckV1Alpha1Manifest{}
}

type KumaIoMeshHealthCheckV1Alpha1Manifest struct{}

type KumaIoMeshHealthCheckV1Alpha1ManifestData struct {
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
		TargetRef *struct {
			Kind *string            `tfsdk:"kind" json:"kind,omitempty"`
			Mesh *string            `tfsdk:"mesh" json:"mesh,omitempty"`
			Name *string            `tfsdk:"name" json:"name,omitempty"`
			Tags *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
		} `tfsdk:"target_ref" json:"targetRef,omitempty"`
		To *[]struct {
			Default *struct {
				AlwaysLogHealthCheckFailures *bool   `tfsdk:"always_log_health_check_failures" json:"alwaysLogHealthCheckFailures,omitempty"`
				EventLogPath                 *string `tfsdk:"event_log_path" json:"eventLogPath,omitempty"`
				FailTrafficOnPanic           *bool   `tfsdk:"fail_traffic_on_panic" json:"failTrafficOnPanic,omitempty"`
				Grpc                         *struct {
					Authority   *string `tfsdk:"authority" json:"authority,omitempty"`
					Disabled    *bool   `tfsdk:"disabled" json:"disabled,omitempty"`
					ServiceName *string `tfsdk:"service_name" json:"serviceName,omitempty"`
				} `tfsdk:"grpc" json:"grpc,omitempty"`
				HealthyPanicThreshold *string `tfsdk:"healthy_panic_threshold" json:"healthyPanicThreshold,omitempty"`
				HealthyThreshold      *int64  `tfsdk:"healthy_threshold" json:"healthyThreshold,omitempty"`
				Http                  *struct {
					Disabled            *bool     `tfsdk:"disabled" json:"disabled,omitempty"`
					ExpectedStatuses    *[]string `tfsdk:"expected_statuses" json:"expectedStatuses,omitempty"`
					Path                *string   `tfsdk:"path" json:"path,omitempty"`
					RequestHeadersToAdd *struct {
						Add *[]struct {
							Name  *string `tfsdk:"name" json:"name,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"add" json:"add,omitempty"`
						Set *[]struct {
							Name  *string `tfsdk:"name" json:"name,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"set" json:"set,omitempty"`
					} `tfsdk:"request_headers_to_add" json:"requestHeadersToAdd,omitempty"`
				} `tfsdk:"http" json:"http,omitempty"`
				InitialJitter         *string `tfsdk:"initial_jitter" json:"initialJitter,omitempty"`
				Interval              *string `tfsdk:"interval" json:"interval,omitempty"`
				IntervalJitter        *string `tfsdk:"interval_jitter" json:"intervalJitter,omitempty"`
				IntervalJitterPercent *int64  `tfsdk:"interval_jitter_percent" json:"intervalJitterPercent,omitempty"`
				NoTrafficInterval     *string `tfsdk:"no_traffic_interval" json:"noTrafficInterval,omitempty"`
				ReuseConnection       *bool   `tfsdk:"reuse_connection" json:"reuseConnection,omitempty"`
				Tcp                   *struct {
					Disabled *bool     `tfsdk:"disabled" json:"disabled,omitempty"`
					Receive  *[]string `tfsdk:"receive" json:"receive,omitempty"`
					Send     *string   `tfsdk:"send" json:"send,omitempty"`
				} `tfsdk:"tcp" json:"tcp,omitempty"`
				Timeout            *string `tfsdk:"timeout" json:"timeout,omitempty"`
				UnhealthyThreshold *int64  `tfsdk:"unhealthy_threshold" json:"unhealthyThreshold,omitempty"`
			} `tfsdk:"default" json:"default,omitempty"`
			TargetRef *struct {
				Kind *string            `tfsdk:"kind" json:"kind,omitempty"`
				Mesh *string            `tfsdk:"mesh" json:"mesh,omitempty"`
				Name *string            `tfsdk:"name" json:"name,omitempty"`
				Tags *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
			} `tfsdk:"target_ref" json:"targetRef,omitempty"`
		} `tfsdk:"to" json:"to,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KumaIoMeshHealthCheckV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kuma_io_mesh_health_check_v1alpha1_manifest"
}

func (r *KumaIoMeshHealthCheckV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
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
				Description:         "Spec is the specification of the Kuma MeshHealthCheck resource.",
				MarkdownDescription: "Spec is the specification of the Kuma MeshHealthCheck resource.",
				Attributes: map[string]schema.Attribute{
					"target_ref": schema.SingleNestedAttribute{
						Description:         "TargetRef is a reference to the resource the policy takes an effect on. The resource could be either a real store object or virtual resource defined inplace.",
						MarkdownDescription: "TargetRef is a reference to the resource the policy takes an effect on. The resource could be either a real store object or virtual resource defined inplace.",
						Attributes: map[string]schema.Attribute{
							"kind": schema.StringAttribute{
								Description:         "Kind of the referenced resource",
								MarkdownDescription: "Kind of the referenced resource",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Mesh", "MeshSubset", "MeshGateway", "MeshService", "MeshServiceSubset", "MeshHTTPRoute"),
								},
							},

							"mesh": schema.StringAttribute{
								Description:         "Mesh is reserved for future use to identify cross mesh resources.",
								MarkdownDescription: "Mesh is reserved for future use to identify cross mesh resources.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referenced resource. Can only be used with kinds: 'MeshService', 'MeshServiceSubset' and 'MeshGatewayRoute'",
								MarkdownDescription: "Name of the referenced resource. Can only be used with kinds: 'MeshService', 'MeshServiceSubset' and 'MeshGatewayRoute'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tags": schema.MapAttribute{
								Description:         "Tags used to select a subset of proxies by tags. Can only be used with kinds 'MeshSubset' and 'MeshServiceSubset'",
								MarkdownDescription: "Tags used to select a subset of proxies by tags. Can only be used with kinds 'MeshSubset' and 'MeshServiceSubset'",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"to": schema.ListNestedAttribute{
						Description:         "To list makes a match between the consumed services and corresponding configurations",
						MarkdownDescription: "To list makes a match between the consumed services and corresponding configurations",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"default": schema.SingleNestedAttribute{
									Description:         "Default is a configuration specific to the group of destinations referenced in 'targetRef'",
									MarkdownDescription: "Default is a configuration specific to the group of destinations referenced in 'targetRef'",
									Attributes: map[string]schema.Attribute{
										"always_log_health_check_failures": schema.BoolAttribute{
											Description:         "If set to true, health check failure events will always be logged. If set to false, only the initial health check failure event will be logged. The default value is false.",
											MarkdownDescription: "If set to true, health check failure events will always be logged. If set to false, only the initial health check failure event will be logged. The default value is false.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"event_log_path": schema.StringAttribute{
											Description:         "Specifies the path to the file where Envoy can log health check events. If empty, no event log will be written.",
											MarkdownDescription: "Specifies the path to the file where Envoy can log health check events. If empty, no event log will be written.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"fail_traffic_on_panic": schema.BoolAttribute{
											Description:         "If set to true, Envoy will not consider any hosts when the cluster is in 'panic mode'. Instead, the cluster will fail all requests as if all hosts are unhealthy. This can help avoid potentially overwhelming a failing service.",
											MarkdownDescription: "If set to true, Envoy will not consider any hosts when the cluster is in 'panic mode'. Instead, the cluster will fail all requests as if all hosts are unhealthy. This can help avoid potentially overwhelming a failing service.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"grpc": schema.SingleNestedAttribute{
											Description:         "GrpcHealthCheck defines gRPC configuration which will instruct the service the health check will be made for is a gRPC service.",
											MarkdownDescription: "GrpcHealthCheck defines gRPC configuration which will instruct the service the health check will be made for is a gRPC service.",
											Attributes: map[string]schema.Attribute{
												"authority": schema.StringAttribute{
													Description:         "The value of the :authority header in the gRPC health check request, by default name of the cluster this health check is associated with",
													MarkdownDescription: "The value of the :authority header in the gRPC health check request, by default name of the cluster this health check is associated with",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"disabled": schema.BoolAttribute{
													Description:         "If true the GrpcHealthCheck is disabled",
													MarkdownDescription: "If true the GrpcHealthCheck is disabled",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"service_name": schema.StringAttribute{
													Description:         "Service name parameter which will be sent to gRPC service",
													MarkdownDescription: "Service name parameter which will be sent to gRPC service",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"healthy_panic_threshold": schema.StringAttribute{
											Description:         "Allows to configure panic threshold for Envoy cluster. If not specified, the default is 50%. To disable panic mode, set to 0%. Either int or decimal represented as string.",
											MarkdownDescription: "Allows to configure panic threshold for Envoy cluster. If not specified, the default is 50%. To disable panic mode, set to 0%. Either int or decimal represented as string.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"healthy_threshold": schema.Int64Attribute{
											Description:         "Number of consecutive healthy checks before considering a host healthy.",
											MarkdownDescription: "Number of consecutive healthy checks before considering a host healthy.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"http": schema.SingleNestedAttribute{
											Description:         "HttpHealthCheck defines HTTP configuration which will instruct the service the health check will be made for is an HTTP service.",
											MarkdownDescription: "HttpHealthCheck defines HTTP configuration which will instruct the service the health check will be made for is an HTTP service.",
											Attributes: map[string]schema.Attribute{
												"disabled": schema.BoolAttribute{
													Description:         "If true the HttpHealthCheck is disabled",
													MarkdownDescription: "If true the HttpHealthCheck is disabled",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"expected_statuses": schema.ListAttribute{
													Description:         "List of HTTP response statuses which are considered healthy",
													MarkdownDescription: "List of HTTP response statuses which are considered healthy",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"path": schema.StringAttribute{
													Description:         "The HTTP path which will be requested during the health check (ie. /health)",
													MarkdownDescription: "The HTTP path which will be requested during the health check (ie. /health)",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"request_headers_to_add": schema.SingleNestedAttribute{
													Description:         "The list of HTTP headers which should be added to each health check request",
													MarkdownDescription: "The list of HTTP headers which should be added to each health check request",
													Attributes: map[string]schema.Attribute{
														"add": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.LengthAtLeast(1),
																			stringvalidator.LengthAtMost(256),
																			stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9!#$%&'*+\-.^_\x60|~]+$`), ""),
																		},
																	},

																	"value": schema.StringAttribute{
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

														"set": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.LengthAtLeast(1),
																			stringvalidator.LengthAtMost(256),
																			stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9!#$%&'*+\-.^_\x60|~]+$`), ""),
																		},
																	},

																	"value": schema.StringAttribute{
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
													Required: false,
													Optional: true,
													Computed: false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"initial_jitter": schema.StringAttribute{
											Description:         "If specified, Envoy will start health checking after a random time in ms between 0 and initialJitter. This only applies to the first health check.",
											MarkdownDescription: "If specified, Envoy will start health checking after a random time in ms between 0 and initialJitter. This only applies to the first health check.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"interval": schema.StringAttribute{
											Description:         "Interval between consecutive health checks.",
											MarkdownDescription: "Interval between consecutive health checks.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"interval_jitter": schema.StringAttribute{
											Description:         "If specified, during every interval Envoy will add IntervalJitter to the wait time.",
											MarkdownDescription: "If specified, during every interval Envoy will add IntervalJitter to the wait time.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"interval_jitter_percent": schema.Int64Attribute{
											Description:         "If specified, during every interval Envoy will add IntervalJitter * IntervalJitterPercent / 100 to the wait time. If IntervalJitter and IntervalJitterPercent are both set, both of them will be used to increase the wait time.",
											MarkdownDescription: "If specified, during every interval Envoy will add IntervalJitter * IntervalJitterPercent / 100 to the wait time. If IntervalJitter and IntervalJitterPercent are both set, both of them will be used to increase the wait time.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"no_traffic_interval": schema.StringAttribute{
											Description:         "The 'no traffic interval' is a special health check interval that is used when a cluster has never had traffic routed to it. This lower interval allows cluster information to be kept up to date, without sending a potentially large amount of active health checking traffic for no reason. Once a cluster has been used for traffic routing, Envoy will shift back to using the standard health check interval that is defined. Note that this interval takes precedence over any other. The default value for 'no traffic interval' is 60 seconds.",
											MarkdownDescription: "The 'no traffic interval' is a special health check interval that is used when a cluster has never had traffic routed to it. This lower interval allows cluster information to be kept up to date, without sending a potentially large amount of active health checking traffic for no reason. Once a cluster has been used for traffic routing, Envoy will shift back to using the standard health check interval that is defined. Note that this interval takes precedence over any other. The default value for 'no traffic interval' is 60 seconds.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"reuse_connection": schema.BoolAttribute{
											Description:         "Reuse health check connection between health checks. Default is true.",
											MarkdownDescription: "Reuse health check connection between health checks. Default is true.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tcp": schema.SingleNestedAttribute{
											Description:         "TcpHealthCheck defines configuration for specifying bytes to send and expected response during the health check",
											MarkdownDescription: "TcpHealthCheck defines configuration for specifying bytes to send and expected response during the health check",
											Attributes: map[string]schema.Attribute{
												"disabled": schema.BoolAttribute{
													Description:         "If true the TcpHealthCheck is disabled",
													MarkdownDescription: "If true the TcpHealthCheck is disabled",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"receive": schema.ListAttribute{
													Description:         "List of Base64 encoded blocks of strings expected as a response. When checking the response, 'fuzzy' matching is performed such that each block must be found, and in the order specified, but not necessarily contiguous. If not provided or empty, checks will be performed as 'connect only' and be marked as successful when TCP connection is successfully established.",
													MarkdownDescription: "List of Base64 encoded blocks of strings expected as a response. When checking the response, 'fuzzy' matching is performed such that each block must be found, and in the order specified, but not necessarily contiguous. If not provided or empty, checks will be performed as 'connect only' and be marked as successful when TCP connection is successfully established.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"send": schema.StringAttribute{
													Description:         "Base64 encoded content of the message which will be sent during the health check to the target",
													MarkdownDescription: "Base64 encoded content of the message which will be sent during the health check to the target",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"timeout": schema.StringAttribute{
											Description:         "Maximum time to wait for a health check response.",
											MarkdownDescription: "Maximum time to wait for a health check response.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"unhealthy_threshold": schema.Int64Attribute{
											Description:         "Number of consecutive unhealthy checks before considering a host unhealthy.",
											MarkdownDescription: "Number of consecutive unhealthy checks before considering a host unhealthy.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"target_ref": schema.SingleNestedAttribute{
									Description:         "TargetRef is a reference to the resource that represents a group of destinations.",
									MarkdownDescription: "TargetRef is a reference to the resource that represents a group of destinations.",
									Attributes: map[string]schema.Attribute{
										"kind": schema.StringAttribute{
											Description:         "Kind of the referenced resource",
											MarkdownDescription: "Kind of the referenced resource",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Mesh", "MeshSubset", "MeshGateway", "MeshService", "MeshServiceSubset", "MeshHTTPRoute"),
											},
										},

										"mesh": schema.StringAttribute{
											Description:         "Mesh is reserved for future use to identify cross mesh resources.",
											MarkdownDescription: "Mesh is reserved for future use to identify cross mesh resources.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name of the referenced resource. Can only be used with kinds: 'MeshService', 'MeshServiceSubset' and 'MeshGatewayRoute'",
											MarkdownDescription: "Name of the referenced resource. Can only be used with kinds: 'MeshService', 'MeshServiceSubset' and 'MeshGatewayRoute'",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tags": schema.MapAttribute{
											Description:         "Tags used to select a subset of proxies by tags. Can only be used with kinds 'MeshSubset' and 'MeshServiceSubset'",
											MarkdownDescription: "Tags used to select a subset of proxies by tags. Can only be used with kinds 'MeshSubset' and 'MeshServiceSubset'",
											ElementType:         types.StringType,
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

func (r *KumaIoMeshHealthCheckV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kuma_io_mesh_health_check_v1alpha1_manifest")

	var model KumaIoMeshHealthCheckV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("kuma.io/v1alpha1")
	model.Kind = pointer.String("MeshHealthCheck")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
