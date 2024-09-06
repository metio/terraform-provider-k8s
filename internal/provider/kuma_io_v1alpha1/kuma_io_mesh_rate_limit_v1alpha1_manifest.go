/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kuma_io_v1alpha1

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &KumaIoMeshRateLimitV1Alpha1Manifest{}
)

func NewKumaIoMeshRateLimitV1Alpha1Manifest() datasource.DataSource {
	return &KumaIoMeshRateLimitV1Alpha1Manifest{}
}

type KumaIoMeshRateLimitV1Alpha1Manifest struct{}

type KumaIoMeshRateLimitV1Alpha1ManifestData struct {
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
		From *[]struct {
			Default *struct {
				Local *struct {
					Http *struct {
						Disabled    *bool `tfsdk:"disabled" json:"disabled,omitempty"`
						OnRateLimit *struct {
							Headers *struct {
								Add *[]struct {
									Name  *string `tfsdk:"name" json:"name,omitempty"`
									Value *string `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"add" json:"add,omitempty"`
								Set *[]struct {
									Name  *string `tfsdk:"name" json:"name,omitempty"`
									Value *string `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"set" json:"set,omitempty"`
							} `tfsdk:"headers" json:"headers,omitempty"`
							Status *int64 `tfsdk:"status" json:"status,omitempty"`
						} `tfsdk:"on_rate_limit" json:"onRateLimit,omitempty"`
						RequestRate *struct {
							Interval *string `tfsdk:"interval" json:"interval,omitempty"`
							Num      *int64  `tfsdk:"num" json:"num,omitempty"`
						} `tfsdk:"request_rate" json:"requestRate,omitempty"`
					} `tfsdk:"http" json:"http,omitempty"`
					Tcp *struct {
						ConnectionRate *struct {
							Interval *string `tfsdk:"interval" json:"interval,omitempty"`
							Num      *int64  `tfsdk:"num" json:"num,omitempty"`
						} `tfsdk:"connection_rate" json:"connectionRate,omitempty"`
						Disabled *bool `tfsdk:"disabled" json:"disabled,omitempty"`
					} `tfsdk:"tcp" json:"tcp,omitempty"`
				} `tfsdk:"local" json:"local,omitempty"`
			} `tfsdk:"default" json:"default,omitempty"`
			TargetRef *struct {
				Kind        *string            `tfsdk:"kind" json:"kind,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				Mesh        *string            `tfsdk:"mesh" json:"mesh,omitempty"`
				Name        *string            `tfsdk:"name" json:"name,omitempty"`
				Namespace   *string            `tfsdk:"namespace" json:"namespace,omitempty"`
				ProxyTypes  *[]string          `tfsdk:"proxy_types" json:"proxyTypes,omitempty"`
				SectionName *string            `tfsdk:"section_name" json:"sectionName,omitempty"`
				Tags        *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
			} `tfsdk:"target_ref" json:"targetRef,omitempty"`
		} `tfsdk:"from" json:"from,omitempty"`
		TargetRef *struct {
			Kind        *string            `tfsdk:"kind" json:"kind,omitempty"`
			Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Mesh        *string            `tfsdk:"mesh" json:"mesh,omitempty"`
			Name        *string            `tfsdk:"name" json:"name,omitempty"`
			Namespace   *string            `tfsdk:"namespace" json:"namespace,omitempty"`
			ProxyTypes  *[]string          `tfsdk:"proxy_types" json:"proxyTypes,omitempty"`
			SectionName *string            `tfsdk:"section_name" json:"sectionName,omitempty"`
			Tags        *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
		} `tfsdk:"target_ref" json:"targetRef,omitempty"`
		To *[]struct {
			Default *struct {
				Local *struct {
					Http *struct {
						Disabled    *bool `tfsdk:"disabled" json:"disabled,omitempty"`
						OnRateLimit *struct {
							Headers *struct {
								Add *[]struct {
									Name  *string `tfsdk:"name" json:"name,omitempty"`
									Value *string `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"add" json:"add,omitempty"`
								Set *[]struct {
									Name  *string `tfsdk:"name" json:"name,omitempty"`
									Value *string `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"set" json:"set,omitempty"`
							} `tfsdk:"headers" json:"headers,omitempty"`
							Status *int64 `tfsdk:"status" json:"status,omitempty"`
						} `tfsdk:"on_rate_limit" json:"onRateLimit,omitempty"`
						RequestRate *struct {
							Interval *string `tfsdk:"interval" json:"interval,omitempty"`
							Num      *int64  `tfsdk:"num" json:"num,omitempty"`
						} `tfsdk:"request_rate" json:"requestRate,omitempty"`
					} `tfsdk:"http" json:"http,omitempty"`
					Tcp *struct {
						ConnectionRate *struct {
							Interval *string `tfsdk:"interval" json:"interval,omitempty"`
							Num      *int64  `tfsdk:"num" json:"num,omitempty"`
						} `tfsdk:"connection_rate" json:"connectionRate,omitempty"`
						Disabled *bool `tfsdk:"disabled" json:"disabled,omitempty"`
					} `tfsdk:"tcp" json:"tcp,omitempty"`
				} `tfsdk:"local" json:"local,omitempty"`
			} `tfsdk:"default" json:"default,omitempty"`
			TargetRef *struct {
				Kind        *string            `tfsdk:"kind" json:"kind,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				Mesh        *string            `tfsdk:"mesh" json:"mesh,omitempty"`
				Name        *string            `tfsdk:"name" json:"name,omitempty"`
				Namespace   *string            `tfsdk:"namespace" json:"namespace,omitempty"`
				ProxyTypes  *[]string          `tfsdk:"proxy_types" json:"proxyTypes,omitempty"`
				SectionName *string            `tfsdk:"section_name" json:"sectionName,omitempty"`
				Tags        *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
			} `tfsdk:"target_ref" json:"targetRef,omitempty"`
		} `tfsdk:"to" json:"to,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KumaIoMeshRateLimitV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kuma_io_mesh_rate_limit_v1alpha1_manifest"
}

func (r *KumaIoMeshRateLimitV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
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
				Description:         "Spec is the specification of the Kuma MeshRateLimit resource.",
				MarkdownDescription: "Spec is the specification of the Kuma MeshRateLimit resource.",
				Attributes: map[string]schema.Attribute{
					"from": schema.ListNestedAttribute{
						Description:         "From list makes a match between clients and corresponding configurations",
						MarkdownDescription: "From list makes a match between clients and corresponding configurations",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"default": schema.SingleNestedAttribute{
									Description:         "Default is a configuration specific to the group of clients referenced in'targetRef'",
									MarkdownDescription: "Default is a configuration specific to the group of clients referenced in'targetRef'",
									Attributes: map[string]schema.Attribute{
										"local": schema.SingleNestedAttribute{
											Description:         "LocalConf defines local http or/and tcp rate limit configuration",
											MarkdownDescription: "LocalConf defines local http or/and tcp rate limit configuration",
											Attributes: map[string]schema.Attribute{
												"http": schema.SingleNestedAttribute{
													Description:         "LocalHTTP defines configuration of local HTTP rate limitinghttps://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/local_rate_limit_filter",
													MarkdownDescription: "LocalHTTP defines configuration of local HTTP rate limitinghttps://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/local_rate_limit_filter",
													Attributes: map[string]schema.Attribute{
														"disabled": schema.BoolAttribute{
															Description:         "Define if rate limiting should be disabled.",
															MarkdownDescription: "Define if rate limiting should be disabled.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"on_rate_limit": schema.SingleNestedAttribute{
															Description:         "Describes the actions to take on a rate limit event",
															MarkdownDescription: "Describes the actions to take on a rate limit event",
															Attributes: map[string]schema.Attribute{
																"headers": schema.SingleNestedAttribute{
																	Description:         "The Headers to be added to the HTTP response on a rate limit event",
																	MarkdownDescription: "The Headers to be added to the HTTP response on a rate limit event",
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

																"status": schema.Int64Attribute{
																	Description:         "The HTTP status code to be set on a rate limit event",
																	MarkdownDescription: "The HTTP status code to be set on a rate limit event",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"request_rate": schema.SingleNestedAttribute{
															Description:         "Defines how many requests are allowed per interval.",
															MarkdownDescription: "Defines how many requests are allowed per interval.",
															Attributes: map[string]schema.Attribute{
																"interval": schema.StringAttribute{
																	Description:         "The interval the number of units is accounted for.",
																	MarkdownDescription: "The interval the number of units is accounted for.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"num": schema.Int64Attribute{
																	Description:         "Number of units per interval (depending on usage it can be a number of requests,or a number of connections).",
																	MarkdownDescription: "Number of units per interval (depending on usage it can be a number of requests,or a number of connections).",
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

												"tcp": schema.SingleNestedAttribute{
													Description:         "LocalTCP defines confguration of local TCP rate limitinghttps://www.envoyproxy.io/docs/envoy/latest/configuration/listeners/network_filters/local_rate_limit_filter",
													MarkdownDescription: "LocalTCP defines confguration of local TCP rate limitinghttps://www.envoyproxy.io/docs/envoy/latest/configuration/listeners/network_filters/local_rate_limit_filter",
													Attributes: map[string]schema.Attribute{
														"connection_rate": schema.SingleNestedAttribute{
															Description:         "Defines how many connections are allowed per interval.",
															MarkdownDescription: "Defines how many connections are allowed per interval.",
															Attributes: map[string]schema.Attribute{
																"interval": schema.StringAttribute{
																	Description:         "The interval the number of units is accounted for.",
																	MarkdownDescription: "The interval the number of units is accounted for.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"num": schema.Int64Attribute{
																	Description:         "Number of units per interval (depending on usage it can be a number of requests,or a number of connections).",
																	MarkdownDescription: "Number of units per interval (depending on usage it can be a number of requests,or a number of connections).",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"disabled": schema.BoolAttribute{
															Description:         "Define if rate limiting should be disabled.Default: false",
															MarkdownDescription: "Define if rate limiting should be disabled.Default: false",
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

								"target_ref": schema.SingleNestedAttribute{
									Description:         "TargetRef is a reference to the resource that represents a group ofclients.",
									MarkdownDescription: "TargetRef is a reference to the resource that represents a group ofclients.",
									Attributes: map[string]schema.Attribute{
										"kind": schema.StringAttribute{
											Description:         "Kind of the referenced resource",
											MarkdownDescription: "Kind of the referenced resource",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Mesh", "MeshSubset", "MeshGateway", "MeshService", "MeshExternalService", "MeshMultiZoneService", "MeshServiceSubset", "MeshHTTPRoute"),
											},
										},

										"labels": schema.MapAttribute{
											Description:         "Labels are used to select group of MeshServices that match labels. Either Labels orName and Namespace can be used.",
											MarkdownDescription: "Labels are used to select group of MeshServices that match labels. Either Labels orName and Namespace can be used.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"mesh": schema.StringAttribute{
											Description:         "Mesh is reserved for future use to identify cross mesh resources.",
											MarkdownDescription: "Mesh is reserved for future use to identify cross mesh resources.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name of the referenced resource. Can only be used with kinds: 'MeshService','MeshServiceSubset' and 'MeshGatewayRoute'",
											MarkdownDescription: "Name of the referenced resource. Can only be used with kinds: 'MeshService','MeshServiceSubset' and 'MeshGatewayRoute'",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"namespace": schema.StringAttribute{
											Description:         "Namespace specifies the namespace of target resource. If empty only resources in policy namespacewill be targeted.",
											MarkdownDescription: "Namespace specifies the namespace of target resource. If empty only resources in policy namespacewill be targeted.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"proxy_types": schema.ListAttribute{
											Description:         "ProxyTypes specifies the data plane types that are subject to the policy. When not specified,all data plane types are targeted by the policy.",
											MarkdownDescription: "ProxyTypes specifies the data plane types that are subject to the policy. When not specified,all data plane types are targeted by the policy.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"section_name": schema.StringAttribute{
											Description:         "SectionName is used to target specific section of resource.For example, you can target port from MeshService.ports[] by its name. Only traffic to this port will be affected.",
											MarkdownDescription: "SectionName is used to target specific section of resource.For example, you can target port from MeshService.ports[] by its name. Only traffic to this port will be affected.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tags": schema.MapAttribute{
											Description:         "Tags used to select a subset of proxies by tags. Can only be used with kinds'MeshSubset' and 'MeshServiceSubset'",
											MarkdownDescription: "Tags used to select a subset of proxies by tags. Can only be used with kinds'MeshSubset' and 'MeshServiceSubset'",
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

					"target_ref": schema.SingleNestedAttribute{
						Description:         "TargetRef is a reference to the resource the policy takes an effect on.The resource could be either a real store object or virtual resourcedefined inplace.",
						MarkdownDescription: "TargetRef is a reference to the resource the policy takes an effect on.The resource could be either a real store object or virtual resourcedefined inplace.",
						Attributes: map[string]schema.Attribute{
							"kind": schema.StringAttribute{
								Description:         "Kind of the referenced resource",
								MarkdownDescription: "Kind of the referenced resource",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Mesh", "MeshSubset", "MeshGateway", "MeshService", "MeshExternalService", "MeshMultiZoneService", "MeshServiceSubset", "MeshHTTPRoute"),
								},
							},

							"labels": schema.MapAttribute{
								Description:         "Labels are used to select group of MeshServices that match labels. Either Labels orName and Namespace can be used.",
								MarkdownDescription: "Labels are used to select group of MeshServices that match labels. Either Labels orName and Namespace can be used.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"mesh": schema.StringAttribute{
								Description:         "Mesh is reserved for future use to identify cross mesh resources.",
								MarkdownDescription: "Mesh is reserved for future use to identify cross mesh resources.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referenced resource. Can only be used with kinds: 'MeshService','MeshServiceSubset' and 'MeshGatewayRoute'",
								MarkdownDescription: "Name of the referenced resource. Can only be used with kinds: 'MeshService','MeshServiceSubset' and 'MeshGatewayRoute'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace specifies the namespace of target resource. If empty only resources in policy namespacewill be targeted.",
								MarkdownDescription: "Namespace specifies the namespace of target resource. If empty only resources in policy namespacewill be targeted.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"proxy_types": schema.ListAttribute{
								Description:         "ProxyTypes specifies the data plane types that are subject to the policy. When not specified,all data plane types are targeted by the policy.",
								MarkdownDescription: "ProxyTypes specifies the data plane types that are subject to the policy. When not specified,all data plane types are targeted by the policy.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"section_name": schema.StringAttribute{
								Description:         "SectionName is used to target specific section of resource.For example, you can target port from MeshService.ports[] by its name. Only traffic to this port will be affected.",
								MarkdownDescription: "SectionName is used to target specific section of resource.For example, you can target port from MeshService.ports[] by its name. Only traffic to this port will be affected.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tags": schema.MapAttribute{
								Description:         "Tags used to select a subset of proxies by tags. Can only be used with kinds'MeshSubset' and 'MeshServiceSubset'",
								MarkdownDescription: "Tags used to select a subset of proxies by tags. Can only be used with kinds'MeshSubset' and 'MeshServiceSubset'",
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

					"to": schema.ListNestedAttribute{
						Description:         "To list makes a match between clients and corresponding configurations",
						MarkdownDescription: "To list makes a match between clients and corresponding configurations",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"default": schema.SingleNestedAttribute{
									Description:         "Default is a configuration specific to the group of clients referenced in'targetRef'",
									MarkdownDescription: "Default is a configuration specific to the group of clients referenced in'targetRef'",
									Attributes: map[string]schema.Attribute{
										"local": schema.SingleNestedAttribute{
											Description:         "LocalConf defines local http or/and tcp rate limit configuration",
											MarkdownDescription: "LocalConf defines local http or/and tcp rate limit configuration",
											Attributes: map[string]schema.Attribute{
												"http": schema.SingleNestedAttribute{
													Description:         "LocalHTTP defines configuration of local HTTP rate limitinghttps://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/local_rate_limit_filter",
													MarkdownDescription: "LocalHTTP defines configuration of local HTTP rate limitinghttps://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/local_rate_limit_filter",
													Attributes: map[string]schema.Attribute{
														"disabled": schema.BoolAttribute{
															Description:         "Define if rate limiting should be disabled.",
															MarkdownDescription: "Define if rate limiting should be disabled.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"on_rate_limit": schema.SingleNestedAttribute{
															Description:         "Describes the actions to take on a rate limit event",
															MarkdownDescription: "Describes the actions to take on a rate limit event",
															Attributes: map[string]schema.Attribute{
																"headers": schema.SingleNestedAttribute{
																	Description:         "The Headers to be added to the HTTP response on a rate limit event",
																	MarkdownDescription: "The Headers to be added to the HTTP response on a rate limit event",
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

																"status": schema.Int64Attribute{
																	Description:         "The HTTP status code to be set on a rate limit event",
																	MarkdownDescription: "The HTTP status code to be set on a rate limit event",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"request_rate": schema.SingleNestedAttribute{
															Description:         "Defines how many requests are allowed per interval.",
															MarkdownDescription: "Defines how many requests are allowed per interval.",
															Attributes: map[string]schema.Attribute{
																"interval": schema.StringAttribute{
																	Description:         "The interval the number of units is accounted for.",
																	MarkdownDescription: "The interval the number of units is accounted for.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"num": schema.Int64Attribute{
																	Description:         "Number of units per interval (depending on usage it can be a number of requests,or a number of connections).",
																	MarkdownDescription: "Number of units per interval (depending on usage it can be a number of requests,or a number of connections).",
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

												"tcp": schema.SingleNestedAttribute{
													Description:         "LocalTCP defines confguration of local TCP rate limitinghttps://www.envoyproxy.io/docs/envoy/latest/configuration/listeners/network_filters/local_rate_limit_filter",
													MarkdownDescription: "LocalTCP defines confguration of local TCP rate limitinghttps://www.envoyproxy.io/docs/envoy/latest/configuration/listeners/network_filters/local_rate_limit_filter",
													Attributes: map[string]schema.Attribute{
														"connection_rate": schema.SingleNestedAttribute{
															Description:         "Defines how many connections are allowed per interval.",
															MarkdownDescription: "Defines how many connections are allowed per interval.",
															Attributes: map[string]schema.Attribute{
																"interval": schema.StringAttribute{
																	Description:         "The interval the number of units is accounted for.",
																	MarkdownDescription: "The interval the number of units is accounted for.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"num": schema.Int64Attribute{
																	Description:         "Number of units per interval (depending on usage it can be a number of requests,or a number of connections).",
																	MarkdownDescription: "Number of units per interval (depending on usage it can be a number of requests,or a number of connections).",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"disabled": schema.BoolAttribute{
															Description:         "Define if rate limiting should be disabled.Default: false",
															MarkdownDescription: "Define if rate limiting should be disabled.Default: false",
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

								"target_ref": schema.SingleNestedAttribute{
									Description:         "TargetRef is a reference to the resource that represents a group ofclients.",
									MarkdownDescription: "TargetRef is a reference to the resource that represents a group ofclients.",
									Attributes: map[string]schema.Attribute{
										"kind": schema.StringAttribute{
											Description:         "Kind of the referenced resource",
											MarkdownDescription: "Kind of the referenced resource",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Mesh", "MeshSubset", "MeshGateway", "MeshService", "MeshExternalService", "MeshMultiZoneService", "MeshServiceSubset", "MeshHTTPRoute"),
											},
										},

										"labels": schema.MapAttribute{
											Description:         "Labels are used to select group of MeshServices that match labels. Either Labels orName and Namespace can be used.",
											MarkdownDescription: "Labels are used to select group of MeshServices that match labels. Either Labels orName and Namespace can be used.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"mesh": schema.StringAttribute{
											Description:         "Mesh is reserved for future use to identify cross mesh resources.",
											MarkdownDescription: "Mesh is reserved for future use to identify cross mesh resources.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name of the referenced resource. Can only be used with kinds: 'MeshService','MeshServiceSubset' and 'MeshGatewayRoute'",
											MarkdownDescription: "Name of the referenced resource. Can only be used with kinds: 'MeshService','MeshServiceSubset' and 'MeshGatewayRoute'",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"namespace": schema.StringAttribute{
											Description:         "Namespace specifies the namespace of target resource. If empty only resources in policy namespacewill be targeted.",
											MarkdownDescription: "Namespace specifies the namespace of target resource. If empty only resources in policy namespacewill be targeted.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"proxy_types": schema.ListAttribute{
											Description:         "ProxyTypes specifies the data plane types that are subject to the policy. When not specified,all data plane types are targeted by the policy.",
											MarkdownDescription: "ProxyTypes specifies the data plane types that are subject to the policy. When not specified,all data plane types are targeted by the policy.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"section_name": schema.StringAttribute{
											Description:         "SectionName is used to target specific section of resource.For example, you can target port from MeshService.ports[] by its name. Only traffic to this port will be affected.",
											MarkdownDescription: "SectionName is used to target specific section of resource.For example, you can target port from MeshService.ports[] by its name. Only traffic to this port will be affected.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tags": schema.MapAttribute{
											Description:         "Tags used to select a subset of proxies by tags. Can only be used with kinds'MeshSubset' and 'MeshServiceSubset'",
											MarkdownDescription: "Tags used to select a subset of proxies by tags. Can only be used with kinds'MeshSubset' and 'MeshServiceSubset'",
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

func (r *KumaIoMeshRateLimitV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kuma_io_mesh_rate_limit_v1alpha1_manifest")

	var model KumaIoMeshRateLimitV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("kuma.io/v1alpha1")
	model.Kind = pointer.String("MeshRateLimit")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
