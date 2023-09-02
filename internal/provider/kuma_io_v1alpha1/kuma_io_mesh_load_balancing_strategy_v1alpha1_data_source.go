/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kuma_io_v1alpha1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &KumaIoMeshLoadBalancingStrategyV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &KumaIoMeshLoadBalancingStrategyV1Alpha1DataSource{}
)

func NewKumaIoMeshLoadBalancingStrategyV1Alpha1DataSource() datasource.DataSource {
	return &KumaIoMeshLoadBalancingStrategyV1Alpha1DataSource{}
}

type KumaIoMeshLoadBalancingStrategyV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type KumaIoMeshLoadBalancingStrategyV1Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

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
				LoadBalancer *struct {
					LeastRequest *struct {
						ChoiceCount *int64 `tfsdk:"choice_count" json:"choiceCount,omitempty"`
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
					Disabled *bool `tfsdk:"disabled" json:"disabled,omitempty"`
				} `tfsdk:"locality_awareness" json:"localityAwareness,omitempty"`
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

func (r *KumaIoMeshLoadBalancingStrategyV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kuma_io_mesh_load_balancing_strategy_v1alpha1"
}

func (r *KumaIoMeshLoadBalancingStrategyV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
						Optional:            false,
						Computed:            true,
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "Spec is the specification of the Kuma MeshLoadBalancingStrategy resource.",
				MarkdownDescription: "Spec is the specification of the Kuma MeshLoadBalancingStrategy resource.",
				Attributes: map[string]schema.Attribute{
					"target_ref": schema.SingleNestedAttribute{
						Description:         "TargetRef is a reference to the resource the policy takes an effect on. The resource could be either a real store object or virtual resource defined inplace.",
						MarkdownDescription: "TargetRef is a reference to the resource the policy takes an effect on. The resource could be either a real store object or virtual resource defined inplace.",
						Attributes: map[string]schema.Attribute{
							"kind": schema.StringAttribute{
								Description:         "Kind of the referenced resource",
								MarkdownDescription: "Kind of the referenced resource",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"mesh": schema.StringAttribute{
								Description:         "Mesh is reserved for future use to identify cross mesh resources.",
								MarkdownDescription: "Mesh is reserved for future use to identify cross mesh resources.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referenced resource. Can only be used with kinds: 'MeshService', 'MeshServiceSubset' and 'MeshGatewayRoute'",
								MarkdownDescription: "Name of the referenced resource. Can only be used with kinds: 'MeshService', 'MeshServiceSubset' and 'MeshGatewayRoute'",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"tags": schema.MapAttribute{
								Description:         "Tags used to select a subset of proxies by tags. Can only be used with kinds 'MeshSubset' and 'MeshServiceSubset'",
								MarkdownDescription: "Tags used to select a subset of proxies by tags. Can only be used with kinds 'MeshSubset' and 'MeshServiceSubset'",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
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
										"load_balancer": schema.SingleNestedAttribute{
											Description:         "LoadBalancer allows to specify load balancing algorithm.",
											MarkdownDescription: "LoadBalancer allows to specify load balancing algorithm.",
											Attributes: map[string]schema.Attribute{
												"least_request": schema.SingleNestedAttribute{
													Description:         "LeastRequest selects N random available hosts as specified in 'choiceCount' (2 by default) and picks the host which has the fewest active requests",
													MarkdownDescription: "LeastRequest selects N random available hosts as specified in 'choiceCount' (2 by default) and picks the host which has the fewest active requests",
													Attributes: map[string]schema.Attribute{
														"choice_count": schema.Int64Attribute{
															Description:         "ChoiceCount is the number of random healthy hosts from which the host with the fewest active requests will be chosen. Defaults to 2 so that Envoy performs two-choice selection if the field is not set.",
															MarkdownDescription: "ChoiceCount is the number of random healthy hosts from which the host with the fewest active requests will be chosen. Defaults to 2 so that Envoy performs two-choice selection if the field is not set.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"maglev": schema.SingleNestedAttribute{
													Description:         "Maglev implements consistent hashing to upstream hosts. Maglev can be used as a drop in replacement for the ring hash load balancer any place in which consistent hashing is desired.",
													MarkdownDescription: "Maglev implements consistent hashing to upstream hosts. Maglev can be used as a drop in replacement for the ring hash load balancer any place in which consistent hashing is desired.",
													Attributes: map[string]schema.Attribute{
														"hash_policies": schema.ListNestedAttribute{
															Description:         "HashPolicies specify a list of request/connection properties that are used to calculate a hash. These hash policies are executed in the specified order. If a hash policy has the “terminal” attribute set to true, and there is already a hash generated, the hash is returned immediately, ignoring the rest of the hash policy list.",
															MarkdownDescription: "HashPolicies specify a list of request/connection properties that are used to calculate a hash. These hash policies are executed in the specified order. If a hash policy has the “terminal” attribute set to true, and there is already a hash generated, the hash is returned immediately, ignoring the rest of the hash policy list.",
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
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"cookie": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "The name of the cookie that will be used to obtain the hash key.",
																				MarkdownDescription: "The name of the cookie that will be used to obtain the hash key.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"path": schema.StringAttribute{
																				Description:         "The name of the path for the cookie.",
																				MarkdownDescription: "The name of the path for the cookie.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"ttl": schema.StringAttribute{
																				Description:         "If specified, a cookie with the TTL will be generated if the cookie is not present.",
																				MarkdownDescription: "If specified, a cookie with the TTL will be generated if the cookie is not present.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"filter_state": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The name of the Object in the per-request filterState, which is an Envoy::Hashable object. If there is no data associated with the key, or the stored object is not Envoy::Hashable, no hash will be produced.",
																				MarkdownDescription: "The name of the Object in the per-request filterState, which is an Envoy::Hashable object. If there is no data associated with the key, or the stored object is not Envoy::Hashable, no hash will be produced.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"header": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "The name of the request header that will be used to obtain the hash key.",
																				MarkdownDescription: "The name of the request header that will be used to obtain the hash key.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"query_parameter": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "The name of the URL query parameter that will be used to obtain the hash key. If the parameter is not present, no hash will be produced. Query parameter names are case-sensitive.",
																				MarkdownDescription: "The name of the URL query parameter that will be used to obtain the hash key. If the parameter is not present, no hash will be produced. Query parameter names are case-sensitive.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"terminal": schema.BoolAttribute{
																		Description:         "Terminal is a flag that short-circuits the hash computing. This field provides a ‘fallback’ style of configuration: “if a terminal policy doesn’t work, fallback to rest of the policy list”, it saves time when the terminal policy works. If true, and there is already a hash computed, ignore rest of the list of hash polices.",
																		MarkdownDescription: "Terminal is a flag that short-circuits the hash computing. This field provides a ‘fallback’ style of configuration: “if a terminal policy doesn’t work, fallback to rest of the policy list”, it saves time when the terminal policy works. If true, and there is already a hash computed, ignore rest of the list of hash polices.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"type": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"table_size": schema.Int64Attribute{
															Description:         "The table size for Maglev hashing. Maglev aims for “minimal disruption” rather than an absolute guarantee. Minimal disruption means that when the set of upstream hosts change, a connection will likely be sent to the same upstream as it was before. Increasing the table size reduces the amount of disruption. The table size must be prime number limited to 5000011. If it is not specified, the default is 65537.",
															MarkdownDescription: "The table size for Maglev hashing. Maglev aims for “minimal disruption” rather than an absolute guarantee. Minimal disruption means that when the set of upstream hosts change, a connection will likely be sent to the same upstream as it was before. Increasing the table size reduces the amount of disruption. The table size must be prime number limited to 5000011. If it is not specified, the default is 65537.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"random": schema.MapAttribute{
													Description:         "Random selects a random available host. The random load balancer generally performs better than round-robin if no health checking policy is configured. Random selection avoids bias towards the host in the set that comes after a failed host.",
													MarkdownDescription: "Random selects a random available host. The random load balancer generally performs better than round-robin if no health checking policy is configured. Random selection avoids bias towards the host in the set that comes after a failed host.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"ring_hash": schema.SingleNestedAttribute{
													Description:         "RingHash  implements consistent hashing to upstream hosts. Each host is mapped onto a circle (the “ring”) by hashing its address; each request is then routed to a host by hashing some property of the request, and finding the nearest corresponding host clockwise around the ring.",
													MarkdownDescription: "RingHash  implements consistent hashing to upstream hosts. Each host is mapped onto a circle (the “ring”) by hashing its address; each request is then routed to a host by hashing some property of the request, and finding the nearest corresponding host clockwise around the ring.",
													Attributes: map[string]schema.Attribute{
														"hash_function": schema.StringAttribute{
															Description:         "HashFunction is a function used to hash hosts onto the ketama ring. The value defaults to XX_HASH. Available values – XX_HASH, MURMUR_HASH_2.",
															MarkdownDescription: "HashFunction is a function used to hash hosts onto the ketama ring. The value defaults to XX_HASH. Available values – XX_HASH, MURMUR_HASH_2.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"hash_policies": schema.ListNestedAttribute{
															Description:         "HashPolicies specify a list of request/connection properties that are used to calculate a hash. These hash policies are executed in the specified order. If a hash policy has the “terminal” attribute set to true, and there is already a hash generated, the hash is returned immediately, ignoring the rest of the hash policy list.",
															MarkdownDescription: "HashPolicies specify a list of request/connection properties that are used to calculate a hash. These hash policies are executed in the specified order. If a hash policy has the “terminal” attribute set to true, and there is already a hash generated, the hash is returned immediately, ignoring the rest of the hash policy list.",
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
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"cookie": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "The name of the cookie that will be used to obtain the hash key.",
																				MarkdownDescription: "The name of the cookie that will be used to obtain the hash key.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"path": schema.StringAttribute{
																				Description:         "The name of the path for the cookie.",
																				MarkdownDescription: "The name of the path for the cookie.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"ttl": schema.StringAttribute{
																				Description:         "If specified, a cookie with the TTL will be generated if the cookie is not present.",
																				MarkdownDescription: "If specified, a cookie with the TTL will be generated if the cookie is not present.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"filter_state": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The name of the Object in the per-request filterState, which is an Envoy::Hashable object. If there is no data associated with the key, or the stored object is not Envoy::Hashable, no hash will be produced.",
																				MarkdownDescription: "The name of the Object in the per-request filterState, which is an Envoy::Hashable object. If there is no data associated with the key, or the stored object is not Envoy::Hashable, no hash will be produced.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"header": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "The name of the request header that will be used to obtain the hash key.",
																				MarkdownDescription: "The name of the request header that will be used to obtain the hash key.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"query_parameter": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "The name of the URL query parameter that will be used to obtain the hash key. If the parameter is not present, no hash will be produced. Query parameter names are case-sensitive.",
																				MarkdownDescription: "The name of the URL query parameter that will be used to obtain the hash key. If the parameter is not present, no hash will be produced. Query parameter names are case-sensitive.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"terminal": schema.BoolAttribute{
																		Description:         "Terminal is a flag that short-circuits the hash computing. This field provides a ‘fallback’ style of configuration: “if a terminal policy doesn’t work, fallback to rest of the policy list”, it saves time when the terminal policy works. If true, and there is already a hash computed, ignore rest of the list of hash polices.",
																		MarkdownDescription: "Terminal is a flag that short-circuits the hash computing. This field provides a ‘fallback’ style of configuration: “if a terminal policy doesn’t work, fallback to rest of the policy list”, it saves time when the terminal policy works. If true, and there is already a hash computed, ignore rest of the list of hash polices.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"type": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"max_ring_size": schema.Int64Attribute{
															Description:         "Maximum hash ring size. Defaults to 8M entries, and limited to 8M entries, but can be lowered to further constrain resource use.",
															MarkdownDescription: "Maximum hash ring size. Defaults to 8M entries, and limited to 8M entries, but can be lowered to further constrain resource use.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"min_ring_size": schema.Int64Attribute{
															Description:         "Minimum hash ring size. The larger the ring is (that is, the more hashes there are for each provided host) the better the request distribution will reflect the desired weights. Defaults to 1024 entries, and limited to 8M entries.",
															MarkdownDescription: "Minimum hash ring size. The larger the ring is (that is, the more hashes there are for each provided host) the better the request distribution will reflect the desired weights. Defaults to 1024 entries, and limited to 8M entries.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"round_robin": schema.MapAttribute{
													Description:         "RoundRobin is a load balancing algorithm that distributes requests across available upstream hosts in round-robin order.",
													MarkdownDescription: "RoundRobin is a load balancing algorithm that distributes requests across available upstream hosts in round-robin order.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"type": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"locality_awareness": schema.SingleNestedAttribute{
											Description:         "LocalityAwareness contains configuration for locality aware load balancing.",
											MarkdownDescription: "LocalityAwareness contains configuration for locality aware load balancing.",
											Attributes: map[string]schema.Attribute{
												"disabled": schema.BoolAttribute{
													Description:         "Disabled allows to disable locality-aware load balancing. When disabled requests are distributed across all endpoints regardless of locality.",
													MarkdownDescription: "Disabled allows to disable locality-aware load balancing. When disabled requests are distributed across all endpoints regardless of locality.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"target_ref": schema.SingleNestedAttribute{
									Description:         "TargetRef is a reference to the resource that represents a group of destinations.",
									MarkdownDescription: "TargetRef is a reference to the resource that represents a group of destinations.",
									Attributes: map[string]schema.Attribute{
										"kind": schema.StringAttribute{
											Description:         "Kind of the referenced resource",
											MarkdownDescription: "Kind of the referenced resource",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"mesh": schema.StringAttribute{
											Description:         "Mesh is reserved for future use to identify cross mesh resources.",
											MarkdownDescription: "Mesh is reserved for future use to identify cross mesh resources.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"name": schema.StringAttribute{
											Description:         "Name of the referenced resource. Can only be used with kinds: 'MeshService', 'MeshServiceSubset' and 'MeshGatewayRoute'",
											MarkdownDescription: "Name of the referenced resource. Can only be used with kinds: 'MeshService', 'MeshServiceSubset' and 'MeshGatewayRoute'",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"tags": schema.MapAttribute{
											Description:         "Tags used to select a subset of proxies by tags. Can only be used with kinds 'MeshSubset' and 'MeshServiceSubset'",
											MarkdownDescription: "Tags used to select a subset of proxies by tags. Can only be used with kinds 'MeshSubset' and 'MeshServiceSubset'",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *KumaIoMeshLoadBalancingStrategyV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *KumaIoMeshLoadBalancingStrategyV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_kuma_io_mesh_load_balancing_strategy_v1alpha1")

	var data KumaIoMeshLoadBalancingStrategyV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "kuma.io", Version: "v1alpha1", Resource: "MeshLoadBalancingStrategy"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse KumaIoMeshLoadBalancingStrategyV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("kuma.io/v1alpha1")
	data.Kind = pointer.String("MeshLoadBalancingStrategy")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
