/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kuma_io_v1alpha1

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
	_ datasource.DataSource = &KumaIoMeshLoadBalancingStrategyV1Alpha1Manifest{}
)

func NewKumaIoMeshLoadBalancingStrategyV1Alpha1Manifest() datasource.DataSource {
	return &KumaIoMeshLoadBalancingStrategyV1Alpha1Manifest{}
}

type KumaIoMeshLoadBalancingStrategyV1Alpha1Manifest struct{}

type KumaIoMeshLoadBalancingStrategyV1Alpha1ManifestData struct {
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
				LoadBalancer *struct {
					LeastRequest *struct {
						ActiveRequestBias *string `tfsdk:"active_request_bias" json:"activeRequestBias,omitempty"`
						ChoiceCount       *int64  `tfsdk:"choice_count" json:"choiceCount,omitempty"`
					} `tfsdk:"least_request" json:"leastRequest,omitempty"`
					Maglev *struct {
						HashPolicies *[]struct {
							Connection *struct {
								SourceIP *bool `tfsdk:"source_ip" json:"sourceIP,omitempty"`
							} `tfsdk:"connection" json:"connection,omitempty"`
							Cookie *struct {
								Name *string `tfsdk:"name" json:"name,omitempty"`
								Path *string `tfsdk:"path" json:"path,omitempty"`
								Ttl  *string `tfsdk:"ttl" json:"ttl,omitempty"`
							} `tfsdk:"cookie" json:"cookie,omitempty"`
							FilterState *struct {
								Key *string `tfsdk:"key" json:"key,omitempty"`
							} `tfsdk:"filter_state" json:"filterState,omitempty"`
							Header *struct {
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"header" json:"header,omitempty"`
							QueryParameter *struct {
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"query_parameter" json:"queryParameter,omitempty"`
							Terminal *bool   `tfsdk:"terminal" json:"terminal,omitempty"`
							Type     *string `tfsdk:"type" json:"type,omitempty"`
						} `tfsdk:"hash_policies" json:"hashPolicies,omitempty"`
						TableSize *int64 `tfsdk:"table_size" json:"tableSize,omitempty"`
					} `tfsdk:"maglev" json:"maglev,omitempty"`
					Random   *map[string]string `tfsdk:"random" json:"random,omitempty"`
					RingHash *struct {
						HashFunction *string `tfsdk:"hash_function" json:"hashFunction,omitempty"`
						HashPolicies *[]struct {
							Connection *struct {
								SourceIP *bool `tfsdk:"source_ip" json:"sourceIP,omitempty"`
							} `tfsdk:"connection" json:"connection,omitempty"`
							Cookie *struct {
								Name *string `tfsdk:"name" json:"name,omitempty"`
								Path *string `tfsdk:"path" json:"path,omitempty"`
								Ttl  *string `tfsdk:"ttl" json:"ttl,omitempty"`
							} `tfsdk:"cookie" json:"cookie,omitempty"`
							FilterState *struct {
								Key *string `tfsdk:"key" json:"key,omitempty"`
							} `tfsdk:"filter_state" json:"filterState,omitempty"`
							Header *struct {
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"header" json:"header,omitempty"`
							QueryParameter *struct {
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"query_parameter" json:"queryParameter,omitempty"`
							Terminal *bool   `tfsdk:"terminal" json:"terminal,omitempty"`
							Type     *string `tfsdk:"type" json:"type,omitempty"`
						} `tfsdk:"hash_policies" json:"hashPolicies,omitempty"`
						MaxRingSize *int64 `tfsdk:"max_ring_size" json:"maxRingSize,omitempty"`
						MinRingSize *int64 `tfsdk:"min_ring_size" json:"minRingSize,omitempty"`
					} `tfsdk:"ring_hash" json:"ringHash,omitempty"`
					RoundRobin *map[string]string `tfsdk:"round_robin" json:"roundRobin,omitempty"`
					Type       *string            `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"load_balancer" json:"loadBalancer,omitempty"`
				LocalityAwareness *struct {
					CrossZone *struct {
						Failover *[]struct {
							From *struct {
								Zones *[]string `tfsdk:"zones" json:"zones,omitempty"`
							} `tfsdk:"from" json:"from,omitempty"`
							To *struct {
								Type  *string   `tfsdk:"type" json:"type,omitempty"`
								Zones *[]string `tfsdk:"zones" json:"zones,omitempty"`
							} `tfsdk:"to" json:"to,omitempty"`
						} `tfsdk:"failover" json:"failover,omitempty"`
						FailoverThreshold *struct {
							Percentage *string `tfsdk:"percentage" json:"percentage,omitempty"`
						} `tfsdk:"failover_threshold" json:"failoverThreshold,omitempty"`
					} `tfsdk:"cross_zone" json:"crossZone,omitempty"`
					Disabled  *bool `tfsdk:"disabled" json:"disabled,omitempty"`
					LocalZone *struct {
						AffinityTags *[]struct {
							Key    *string `tfsdk:"key" json:"key,omitempty"`
							Weight *int64  `tfsdk:"weight" json:"weight,omitempty"`
						} `tfsdk:"affinity_tags" json:"affinityTags,omitempty"`
					} `tfsdk:"local_zone" json:"localZone,omitempty"`
				} `tfsdk:"locality_awareness" json:"localityAwareness,omitempty"`
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

func (r *KumaIoMeshLoadBalancingStrategyV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kuma_io_mesh_load_balancing_strategy_v1alpha1_manifest"
}

func (r *KumaIoMeshLoadBalancingStrategyV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "Spec is the specification of the Kuma MeshLoadBalancingStrategy resource.",
				MarkdownDescription: "Spec is the specification of the Kuma MeshLoadBalancingStrategy resource.",
				Attributes: map[string]schema.Attribute{
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
									stringvalidator.OneOf("Mesh", "MeshSubset", "MeshGateway", "MeshService", "MeshExternalService", "MeshServiceSubset", "MeshHTTPRoute"),
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

					"to": schema.ListNestedAttribute{
						Description:         "To list makes a match between the consumed services and corresponding configurations",
						MarkdownDescription: "To list makes a match between the consumed services and corresponding configurations",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"default": schema.SingleNestedAttribute{
									Description:         "Default is a configuration specific to the group of destinations referenced in'targetRef'",
									MarkdownDescription: "Default is a configuration specific to the group of destinations referenced in'targetRef'",
									Attributes: map[string]schema.Attribute{
										"load_balancer": schema.SingleNestedAttribute{
											Description:         "LoadBalancer allows to specify load balancing algorithm.",
											MarkdownDescription: "LoadBalancer allows to specify load balancing algorithm.",
											Attributes: map[string]schema.Attribute{
												"least_request": schema.SingleNestedAttribute{
													Description:         "LeastRequest selects N random available hosts as specified in 'choiceCount' (2 by default)and picks the host which has the fewest active requests",
													MarkdownDescription: "LeastRequest selects N random available hosts as specified in 'choiceCount' (2 by default)and picks the host which has the fewest active requests",
													Attributes: map[string]schema.Attribute{
														"active_request_bias": schema.StringAttribute{
															Description:         "ActiveRequestBias refers to dynamic weights applied when hosts have varying loadbalancing weights. A higher value here aggressively reduces the weight of endpointsthat are currently handling active requests. In essence, the higher the ActiveRequestBiasvalue, the more forcefully it reduces the load balancing weight of endpoints that areactively serving requests.",
															MarkdownDescription: "ActiveRequestBias refers to dynamic weights applied when hosts have varying loadbalancing weights. A higher value here aggressively reduces the weight of endpointsthat are currently handling active requests. In essence, the higher the ActiveRequestBiasvalue, the more forcefully it reduces the load balancing weight of endpoints that areactively serving requests.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"choice_count": schema.Int64Attribute{
															Description:         "ChoiceCount is the number of random healthy hosts from which the host withthe fewest active requests will be chosen. Defaults to 2 so that Envoy performstwo-choice selection if the field is not set.",
															MarkdownDescription: "ChoiceCount is the number of random healthy hosts from which the host withthe fewest active requests will be chosen. Defaults to 2 so that Envoy performstwo-choice selection if the field is not set.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(2),
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"maglev": schema.SingleNestedAttribute{
													Description:         "Maglev implements consistent hashing to upstream hosts. Maglev can be used asa drop in replacement for the ring hash load balancer any place in whichconsistent hashing is desired.",
													MarkdownDescription: "Maglev implements consistent hashing to upstream hosts. Maglev can be used asa drop in replacement for the ring hash load balancer any place in whichconsistent hashing is desired.",
													Attributes: map[string]schema.Attribute{
														"hash_policies": schema.ListNestedAttribute{
															Description:         "HashPolicies specify a list of request/connection properties that are used to calculate a hash.These hash policies are executed in the specified order. If a hash policy has the “terminal” attributeset to true, and there is already a hash generated, the hash is returned immediately,ignoring the rest of the hash policy list.",
															MarkdownDescription: "HashPolicies specify a list of request/connection properties that are used to calculate a hash.These hash policies are executed in the specified order. If a hash policy has the “terminal” attributeset to true, and there is already a hash generated, the hash is returned immediately,ignoring the rest of the hash policy list.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"connection": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"source_ip": schema.BoolAttribute{
																				Description:         "Hash on source IP address.",
																				MarkdownDescription: "Hash on source IP address.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"cookie": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "The name of the cookie that will be used to obtain the hash key.",
																				MarkdownDescription: "The name of the cookie that will be used to obtain the hash key.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																				Validators: []validator.String{
																					stringvalidator.LengthAtLeast(1),
																				},
																			},

																			"path": schema.StringAttribute{
																				Description:         "The name of the path for the cookie.",
																				MarkdownDescription: "The name of the path for the cookie.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"ttl": schema.StringAttribute{
																				Description:         "If specified, a cookie with the TTL will be generated if the cookie is not present.",
																				MarkdownDescription: "If specified, a cookie with the TTL will be generated if the cookie is not present.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"filter_state": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The name of the Object in the per-request filterState, which isan Envoy::Hashable object. If there is no data associated with the key,or the stored object is not Envoy::Hashable, no hash will be produced.",
																				MarkdownDescription: "The name of the Object in the per-request filterState, which isan Envoy::Hashable object. If there is no data associated with the key,or the stored object is not Envoy::Hashable, no hash will be produced.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																				Validators: []validator.String{
																					stringvalidator.LengthAtLeast(1),
																				},
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"header": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "The name of the request header that will be used to obtain the hash key.",
																				MarkdownDescription: "The name of the request header that will be used to obtain the hash key.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																				Validators: []validator.String{
																					stringvalidator.LengthAtLeast(1),
																				},
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"query_parameter": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "The name of the URL query parameter that will be used to obtain the hash key.If the parameter is not present, no hash will be produced. Query parameter namesare case-sensitive.",
																				MarkdownDescription: "The name of the URL query parameter that will be used to obtain the hash key.If the parameter is not present, no hash will be produced. Query parameter namesare case-sensitive.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																				Validators: []validator.String{
																					stringvalidator.LengthAtLeast(1),
																				},
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"terminal": schema.BoolAttribute{
																		Description:         "Terminal is a flag that short-circuits the hash computing. This field providesa ‘fallback’ style of configuration: “if a terminal policy doesn’t work, fallbackto rest of the policy list”, it saves time when the terminal policy works.If true, and there is already a hash computed, ignore rest of the list of hash polices.",
																		MarkdownDescription: "Terminal is a flag that short-circuits the hash computing. This field providesa ‘fallback’ style of configuration: “if a terminal policy doesn’t work, fallbackto rest of the policy list”, it saves time when the terminal policy works.If true, and there is already a hash computed, ignore rest of the list of hash polices.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"type": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("Header", "Cookie", "SourceIP", "QueryParameter", "FilterState"),
																		},
																	},
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"table_size": schema.Int64Attribute{
															Description:         "The table size for Maglev hashing. Maglev aims for “minimal disruption”rather than an absolute guarantee. Minimal disruption means that whenthe set of upstream hosts change, a connection will likely be sentto the same upstream as it was before. Increasing the table size reducesthe amount of disruption. The table size must be prime number limited to 5000011.If it is not specified, the default is 65537.",
															MarkdownDescription: "The table size for Maglev hashing. Maglev aims for “minimal disruption”rather than an absolute guarantee. Minimal disruption means that whenthe set of upstream hosts change, a connection will likely be sentto the same upstream as it was before. Increasing the table size reducesthe amount of disruption. The table size must be prime number limited to 5000011.If it is not specified, the default is 65537.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(1),
																int64validator.AtMost(5.000011e+06),
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"random": schema.MapAttribute{
													Description:         "Random selects a random available host. The random load balancer generallyperforms better than round-robin if no health checking policy is configured.Random selection avoids bias towards the host in the set that comes after a failed host.",
													MarkdownDescription: "Random selects a random available host. The random load balancer generallyperforms better than round-robin if no health checking policy is configured.Random selection avoids bias towards the host in the set that comes after a failed host.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"ring_hash": schema.SingleNestedAttribute{
													Description:         "RingHash  implements consistent hashing to upstream hosts. Each host is mappedonto a circle (the “ring”) by hashing its address; each request is then routedto a host by hashing some property of the request, and finding the nearestcorresponding host clockwise around the ring.",
													MarkdownDescription: "RingHash  implements consistent hashing to upstream hosts. Each host is mappedonto a circle (the “ring”) by hashing its address; each request is then routedto a host by hashing some property of the request, and finding the nearestcorresponding host clockwise around the ring.",
													Attributes: map[string]schema.Attribute{
														"hash_function": schema.StringAttribute{
															Description:         "HashFunction is a function used to hash hosts onto the ketama ring.The value defaults to XX_HASH. Available values – XX_HASH, MURMUR_HASH_2.",
															MarkdownDescription: "HashFunction is a function used to hash hosts onto the ketama ring.The value defaults to XX_HASH. Available values – XX_HASH, MURMUR_HASH_2.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("XXHash", "MurmurHash2"),
															},
														},

														"hash_policies": schema.ListNestedAttribute{
															Description:         "HashPolicies specify a list of request/connection properties that are used to calculate a hash.These hash policies are executed in the specified order. If a hash policy has the “terminal” attributeset to true, and there is already a hash generated, the hash is returned immediately,ignoring the rest of the hash policy list.",
															MarkdownDescription: "HashPolicies specify a list of request/connection properties that are used to calculate a hash.These hash policies are executed in the specified order. If a hash policy has the “terminal” attributeset to true, and there is already a hash generated, the hash is returned immediately,ignoring the rest of the hash policy list.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"connection": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"source_ip": schema.BoolAttribute{
																				Description:         "Hash on source IP address.",
																				MarkdownDescription: "Hash on source IP address.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"cookie": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "The name of the cookie that will be used to obtain the hash key.",
																				MarkdownDescription: "The name of the cookie that will be used to obtain the hash key.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																				Validators: []validator.String{
																					stringvalidator.LengthAtLeast(1),
																				},
																			},

																			"path": schema.StringAttribute{
																				Description:         "The name of the path for the cookie.",
																				MarkdownDescription: "The name of the path for the cookie.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"ttl": schema.StringAttribute{
																				Description:         "If specified, a cookie with the TTL will be generated if the cookie is not present.",
																				MarkdownDescription: "If specified, a cookie with the TTL will be generated if the cookie is not present.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"filter_state": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The name of the Object in the per-request filterState, which isan Envoy::Hashable object. If there is no data associated with the key,or the stored object is not Envoy::Hashable, no hash will be produced.",
																				MarkdownDescription: "The name of the Object in the per-request filterState, which isan Envoy::Hashable object. If there is no data associated with the key,or the stored object is not Envoy::Hashable, no hash will be produced.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																				Validators: []validator.String{
																					stringvalidator.LengthAtLeast(1),
																				},
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"header": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "The name of the request header that will be used to obtain the hash key.",
																				MarkdownDescription: "The name of the request header that will be used to obtain the hash key.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																				Validators: []validator.String{
																					stringvalidator.LengthAtLeast(1),
																				},
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"query_parameter": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "The name of the URL query parameter that will be used to obtain the hash key.If the parameter is not present, no hash will be produced. Query parameter namesare case-sensitive.",
																				MarkdownDescription: "The name of the URL query parameter that will be used to obtain the hash key.If the parameter is not present, no hash will be produced. Query parameter namesare case-sensitive.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																				Validators: []validator.String{
																					stringvalidator.LengthAtLeast(1),
																				},
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"terminal": schema.BoolAttribute{
																		Description:         "Terminal is a flag that short-circuits the hash computing. This field providesa ‘fallback’ style of configuration: “if a terminal policy doesn’t work, fallbackto rest of the policy list”, it saves time when the terminal policy works.If true, and there is already a hash computed, ignore rest of the list of hash polices.",
																		MarkdownDescription: "Terminal is a flag that short-circuits the hash computing. This field providesa ‘fallback’ style of configuration: “if a terminal policy doesn’t work, fallbackto rest of the policy list”, it saves time when the terminal policy works.If true, and there is already a hash computed, ignore rest of the list of hash polices.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"type": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("Header", "Cookie", "SourceIP", "QueryParameter", "FilterState"),
																		},
																	},
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"max_ring_size": schema.Int64Attribute{
															Description:         "Maximum hash ring size. Defaults to 8M entries, and limited to 8M entries,but can be lowered to further constrain resource use.",
															MarkdownDescription: "Maximum hash ring size. Defaults to 8M entries, and limited to 8M entries,but can be lowered to further constrain resource use.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(1),
																int64validator.AtMost(8e+06),
															},
														},

														"min_ring_size": schema.Int64Attribute{
															Description:         "Minimum hash ring size. The larger the ring is (that is,the more hashes there are for each provided host) the better the request distributionwill reflect the desired weights. Defaults to 1024 entries, and limited to 8M entries.",
															MarkdownDescription: "Minimum hash ring size. The larger the ring is (that is,the more hashes there are for each provided host) the better the request distributionwill reflect the desired weights. Defaults to 1024 entries, and limited to 8M entries.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(1),
																int64validator.AtMost(8e+06),
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"round_robin": schema.MapAttribute{
													Description:         "RoundRobin is a load balancing algorithm that distributes requestsacross available upstream hosts in round-robin order.",
													MarkdownDescription: "RoundRobin is a load balancing algorithm that distributes requestsacross available upstream hosts in round-robin order.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"type": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("RoundRobin", "LeastRequest", "RingHash", "Random", "Maglev"),
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"locality_awareness": schema.SingleNestedAttribute{
											Description:         "LocalityAwareness contains configuration for locality aware load balancing.",
											MarkdownDescription: "LocalityAwareness contains configuration for locality aware load balancing.",
											Attributes: map[string]schema.Attribute{
												"cross_zone": schema.SingleNestedAttribute{
													Description:         "CrossZone defines locality aware load balancing priorities when dataplane proxies inside local zoneare unavailable",
													MarkdownDescription: "CrossZone defines locality aware load balancing priorities when dataplane proxies inside local zoneare unavailable",
													Attributes: map[string]schema.Attribute{
														"failover": schema.ListNestedAttribute{
															Description:         "Failover defines list of load balancing rules in order of priority",
															MarkdownDescription: "Failover defines list of load balancing rules in order of priority",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"from": schema.SingleNestedAttribute{
																		Description:         "From defines the list of zones to which the rule applies",
																		MarkdownDescription: "From defines the list of zones to which the rule applies",
																		Attributes: map[string]schema.Attribute{
																			"zones": schema.ListAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				ElementType:         types.StringType,
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"to": schema.SingleNestedAttribute{
																		Description:         "To defines to which zones the traffic should be load balanced",
																		MarkdownDescription: "To defines to which zones the traffic should be load balanced",
																		Attributes: map[string]schema.Attribute{
																			"type": schema.StringAttribute{
																				Description:         "Type defines how target zones will be picked from available zones",
																				MarkdownDescription: "Type defines how target zones will be picked from available zones",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																				Validators: []validator.String{
																					stringvalidator.OneOf("None", "Only", "Any", "AnyExcept"),
																				},
																			},

																			"zones": schema.ListAttribute{
																				Description:         "",
																				MarkdownDescription: "",
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

														"failover_threshold": schema.SingleNestedAttribute{
															Description:         "FailoverThreshold defines the percentage of live destination dataplane proxies below which load balancing to thenext priority starts.Example: If you configure failoverThreshold to 70, and you have deployed 10 destination dataplane proxies.Load balancing to next priority will start when number of live destination dataplane proxies drops below 7.Default 50",
															MarkdownDescription: "FailoverThreshold defines the percentage of live destination dataplane proxies below which load balancing to thenext priority starts.Example: If you configure failoverThreshold to 70, and you have deployed 10 destination dataplane proxies.Load balancing to next priority will start when number of live destination dataplane proxies drops below 7.Default 50",
															Attributes: map[string]schema.Attribute{
																"percentage": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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

												"disabled": schema.BoolAttribute{
													Description:         "Disabled allows to disable locality-aware load balancing.When disabled requests are distributed across all endpoints regardless of locality.",
													MarkdownDescription: "Disabled allows to disable locality-aware load balancing.When disabled requests are distributed across all endpoints regardless of locality.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"local_zone": schema.SingleNestedAttribute{
													Description:         "LocalZone defines locality aware load balancing priorities between dataplane proxies inside a zone",
													MarkdownDescription: "LocalZone defines locality aware load balancing priorities between dataplane proxies inside a zone",
													Attributes: map[string]schema.Attribute{
														"affinity_tags": schema.ListNestedAttribute{
															Description:         "AffinityTags list of tags for local zone load balancing.",
															MarkdownDescription: "AffinityTags list of tags for local zone load balancing.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "Key defines tag for which affinity is configured",
																		MarkdownDescription: "Key defines tag for which affinity is configured",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"weight": schema.Int64Attribute{
																		Description:         "Weight of the tag used for load balancing. The bigger the weight the bigger the priority.Percentage of local traffic load balanced to tag is computed by dividing weight by sum of weights from all tags.For example with two affinity tags first with weight 80 and second with weight 20,then 80% of traffic will be redirected to the first tag, and 20% of traffic will be redirected to second one.Setting weights is not mandatory. When weights are not set control plane will compute default weight based on list order.Default: If you do not specify weight we will adjust them so that 90% traffic goes to first tag, 9% to next, and 1% to third and so on.",
																		MarkdownDescription: "Weight of the tag used for load balancing. The bigger the weight the bigger the priority.Percentage of local traffic load balanced to tag is computed by dividing weight by sum of weights from all tags.For example with two affinity tags first with weight 80 and second with weight 20,then 80% of traffic will be redirected to the first tag, and 20% of traffic will be redirected to second one.Setting weights is not mandatory. When weights are not set control plane will compute default weight based on list order.Default: If you do not specify weight we will adjust them so that 90% traffic goes to first tag, 9% to next, and 1% to third and so on.",
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
									Description:         "TargetRef is a reference to the resource that represents a group ofdestinations.",
									MarkdownDescription: "TargetRef is a reference to the resource that represents a group ofdestinations.",
									Attributes: map[string]schema.Attribute{
										"kind": schema.StringAttribute{
											Description:         "Kind of the referenced resource",
											MarkdownDescription: "Kind of the referenced resource",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Mesh", "MeshSubset", "MeshGateway", "MeshService", "MeshExternalService", "MeshServiceSubset", "MeshHTTPRoute"),
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

func (r *KumaIoMeshLoadBalancingStrategyV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kuma_io_mesh_load_balancing_strategy_v1alpha1_manifest")

	var model KumaIoMeshLoadBalancingStrategyV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("kuma.io/v1alpha1")
	model.Kind = pointer.String("MeshLoadBalancingStrategy")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
