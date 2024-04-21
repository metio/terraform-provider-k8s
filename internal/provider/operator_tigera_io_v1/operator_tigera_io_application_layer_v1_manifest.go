/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package operator_tigera_io_v1

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
	_ datasource.DataSource = &OperatorTigeraIoApplicationLayerV1Manifest{}
)

func NewOperatorTigeraIoApplicationLayerV1Manifest() datasource.DataSource {
	return &OperatorTigeraIoApplicationLayerV1Manifest{}
}

type OperatorTigeraIoApplicationLayerV1Manifest struct{}

type OperatorTigeraIoApplicationLayerV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		ApplicationLayerPolicy *string `tfsdk:"application_layer_policy" json:"applicationLayerPolicy,omitempty"`
		Envoy                  *struct {
			UseRemoteAddress  *bool  `tfsdk:"use_remote_address" json:"useRemoteAddress,omitempty"`
			XffNumTrustedHops *int64 `tfsdk:"xff_num_trusted_hops" json:"xffNumTrustedHops,omitempty"`
		} `tfsdk:"envoy" json:"envoy,omitempty"`
		L7LogCollectorDaemonSet *struct {
			Spec *struct {
				Template *struct {
					Spec *struct {
						Containers *[]struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Resources *struct {
								Claims *[]struct {
									Name *string `tfsdk:"name" json:"name,omitempty"`
								} `tfsdk:"claims" json:"claims,omitempty"`
								Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
								Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
							} `tfsdk:"resources" json:"resources,omitempty"`
						} `tfsdk:"containers" json:"containers,omitempty"`
						InitContainers *[]struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Resources *struct {
								Claims *[]struct {
									Name *string `tfsdk:"name" json:"name,omitempty"`
								} `tfsdk:"claims" json:"claims,omitempty"`
								Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
								Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
							} `tfsdk:"resources" json:"resources,omitempty"`
						} `tfsdk:"init_containers" json:"initContainers,omitempty"`
					} `tfsdk:"spec" json:"spec,omitempty"`
				} `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"spec" json:"spec,omitempty"`
		} `tfsdk:"l7_log_collector_daemon_set" json:"l7LogCollectorDaemonSet,omitempty"`
		LogCollection *struct {
			CollectLogs            *string `tfsdk:"collect_logs" json:"collectLogs,omitempty"`
			LogIntervalSeconds     *int64  `tfsdk:"log_interval_seconds" json:"logIntervalSeconds,omitempty"`
			LogRequestsPerInterval *int64  `tfsdk:"log_requests_per_interval" json:"logRequestsPerInterval,omitempty"`
		} `tfsdk:"log_collection" json:"logCollection,omitempty"`
		WebApplicationFirewall *string `tfsdk:"web_application_firewall" json:"webApplicationFirewall,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OperatorTigeraIoApplicationLayerV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_operator_tigera_io_application_layer_v1_manifest"
}

func (r *OperatorTigeraIoApplicationLayerV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ApplicationLayer is the Schema for the applicationlayers API",
		MarkdownDescription: "ApplicationLayer is the Schema for the applicationlayers API",
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
				Description:         "ApplicationLayerSpec defines the desired state of ApplicationLayer",
				MarkdownDescription: "ApplicationLayerSpec defines the desired state of ApplicationLayer",
				Attributes: map[string]schema.Attribute{
					"application_layer_policy": schema.StringAttribute{
						Description:         "Application Layer Policy controls whether or not ALP enforcement is enabled for the cluster.When enabled, NetworkPolicies with HTTP Match rules may be defined to opt-in workloads for traffic enforcement on the application layer.",
						MarkdownDescription: "Application Layer Policy controls whether or not ALP enforcement is enabled for the cluster.When enabled, NetworkPolicies with HTTP Match rules may be defined to opt-in workloads for traffic enforcement on the application layer.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"envoy": schema.SingleNestedAttribute{
						Description:         "User-configurable settings for the Envoy proxy.",
						MarkdownDescription: "User-configurable settings for the Envoy proxy.",
						Attributes: map[string]schema.Attribute{
							"use_remote_address": schema.BoolAttribute{
								Description:         "If set to true, the Envoy connection manager will use the real remote addressof the client connection when determining internal versus external origin andmanipulating various headers.",
								MarkdownDescription: "If set to true, the Envoy connection manager will use the real remote addressof the client connection when determining internal versus external origin andmanipulating various headers.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"xff_num_trusted_hops": schema.Int64Attribute{
								Description:         "The number of additional ingress proxy hops from the right side of thex-forwarded-for HTTP header to trust when determining the origin client’sIP address. 0 is permitted, but >=1 is the typical setting.",
								MarkdownDescription: "The number of additional ingress proxy hops from the right side of thex-forwarded-for HTTP header to trust when determining the origin client’sIP address. 0 is permitted, but >=1 is the typical setting.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
									int64validator.AtMost(2.147483647e+09),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"l7_log_collector_daemon_set": schema.SingleNestedAttribute{
						Description:         "L7LogCollectorDaemonSet configures the L7LogCollector DaemonSet.",
						MarkdownDescription: "L7LogCollectorDaemonSet configures the L7LogCollector DaemonSet.",
						Attributes: map[string]schema.Attribute{
							"spec": schema.SingleNestedAttribute{
								Description:         "Spec is the specification of the L7LogCollector DaemonSet.",
								MarkdownDescription: "Spec is the specification of the L7LogCollector DaemonSet.",
								Attributes: map[string]schema.Attribute{
									"template": schema.SingleNestedAttribute{
										Description:         "Template describes the L7LogCollector DaemonSet pod that will be created.",
										MarkdownDescription: "Template describes the L7LogCollector DaemonSet pod that will be created.",
										Attributes: map[string]schema.Attribute{
											"spec": schema.SingleNestedAttribute{
												Description:         "Spec is the L7LogCollector DaemonSet's PodSpec.",
												MarkdownDescription: "Spec is the L7LogCollector DaemonSet's PodSpec.",
												Attributes: map[string]schema.Attribute{
													"containers": schema.ListNestedAttribute{
														Description:         "Containers is a list of L7LogCollector DaemonSet containers.If specified, this overrides the specified L7LogCollector DaemonSet containers.If omitted, the L7LogCollector DaemonSet will use its default values for its containers.",
														MarkdownDescription: "Containers is a list of L7LogCollector DaemonSet containers.If specified, this overrides the specified L7LogCollector DaemonSet containers.If omitted, the L7LogCollector DaemonSet will use its default values for its containers.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is an enum which identifies the L7LogCollector DaemonSet container by name.Supported values are: l7-collector, envoy-proxy, dikastes",
																	MarkdownDescription: "Name is an enum which identifies the L7LogCollector DaemonSet container by name.Supported values are: l7-collector, envoy-proxy, dikastes",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("l7-collector", "envoy-proxy", "dikastes"),
																	},
																},

																"resources": schema.SingleNestedAttribute{
																	Description:         "Resources allows customization of limits and requests for compute resources such as cpu and memory.If specified, this overrides the named L7LogCollector DaemonSet container's resources.If omitted, the L7LogCollector DaemonSet will use its default value for this container's resources.",
																	MarkdownDescription: "Resources allows customization of limits and requests for compute resources such as cpu and memory.If specified, this overrides the named L7LogCollector DaemonSet container's resources.If omitted, the L7LogCollector DaemonSet will use its default value for this container's resources.",
																	Attributes: map[string]schema.Attribute{
																		"claims": schema.ListNestedAttribute{
																			Description:         "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
																			MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
																						MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
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

																		"limits": schema.MapAttribute{
																			Description:         "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"requests": schema.MapAttribute{
																			Description:         "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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
														Required: false,
														Optional: true,
														Computed: false,
													},

													"init_containers": schema.ListNestedAttribute{
														Description:         "InitContainers is a list of L7LogCollector DaemonSet init containers.If specified, this overrides the specified L7LogCollector DaemonSet init containers.If omitted, the L7LogCollector DaemonSet will use its default values for its init containers.",
														MarkdownDescription: "InitContainers is a list of L7LogCollector DaemonSet init containers.If specified, this overrides the specified L7LogCollector DaemonSet init containers.If omitted, the L7LogCollector DaemonSet will use its default values for its init containers.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is an enum which identifies the L7LogCollector DaemonSet init container by name.",
																	MarkdownDescription: "Name is an enum which identifies the L7LogCollector DaemonSet init container by name.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"resources": schema.SingleNestedAttribute{
																	Description:         "Resources allows customization of limits and requests for compute resources such as cpu and memory.If specified, this overrides the named L7LogCollector DaemonSet init container's resources.If omitted, the L7LogCollector DaemonSet will use its default value for this init container's resources.",
																	MarkdownDescription: "Resources allows customization of limits and requests for compute resources such as cpu and memory.If specified, this overrides the named L7LogCollector DaemonSet init container's resources.If omitted, the L7LogCollector DaemonSet will use its default value for this init container's resources.",
																	Attributes: map[string]schema.Attribute{
																		"claims": schema.ListNestedAttribute{
																			Description:         "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
																			MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
																						MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
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

																		"limits": schema.MapAttribute{
																			Description:         "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"requests": schema.MapAttribute{
																			Description:         "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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
								Required: false,
								Optional: true,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"log_collection": schema.SingleNestedAttribute{
						Description:         "Specification for application layer (L7) log collection.",
						MarkdownDescription: "Specification for application layer (L7) log collection.",
						Attributes: map[string]schema.Attribute{
							"collect_logs": schema.StringAttribute{
								Description:         "This setting enables or disable log collection.Allowed values are Enabled or Disabled.",
								MarkdownDescription: "This setting enables or disable log collection.Allowed values are Enabled or Disabled.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"log_interval_seconds": schema.Int64Attribute{
								Description:         "Interval in seconds for sending L7 log information for processing.Default: 5 sec",
								MarkdownDescription: "Interval in seconds for sending L7 log information for processing.Default: 5 sec",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"log_requests_per_interval": schema.Int64Attribute{
								Description:         "Maximum number of unique L7 logs that are sent LogIntervalSeconds.Adjust this to limit the number of L7 logs sent per LogIntervalSecondsto felix for further processing, use negative number to ignore limits.Default: -1",
								MarkdownDescription: "Maximum number of unique L7 logs that are sent LogIntervalSeconds.Adjust this to limit the number of L7 logs sent per LogIntervalSecondsto felix for further processing, use negative number to ignore limits.Default: -1",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"web_application_firewall": schema.StringAttribute{
						Description:         "WebApplicationFirewall controls whether or not ModSecurity enforcement is enabled for the cluster.When enabled, Services may opt-in to having ingress traffic examed by ModSecurity.",
						MarkdownDescription: "WebApplicationFirewall controls whether or not ModSecurity enforcement is enabled for the cluster.When enabled, Services may opt-in to having ingress traffic examed by ModSecurity.",
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
	}
}

func (r *OperatorTigeraIoApplicationLayerV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_operator_tigera_io_application_layer_v1_manifest")

	var model OperatorTigeraIoApplicationLayerV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("operator.tigera.io/v1")
	model.Kind = pointer.String("ApplicationLayer")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
