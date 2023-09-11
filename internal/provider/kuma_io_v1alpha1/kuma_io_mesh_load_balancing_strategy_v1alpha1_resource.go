/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kuma_io_v1alpha1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	k8sTypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"strings"
	"time"
)

var (
	_ resource.Resource                = &KumaIoMeshLoadBalancingStrategyV1Alpha1Resource{}
	_ resource.ResourceWithConfigure   = &KumaIoMeshLoadBalancingStrategyV1Alpha1Resource{}
	_ resource.ResourceWithImportState = &KumaIoMeshLoadBalancingStrategyV1Alpha1Resource{}
)

func NewKumaIoMeshLoadBalancingStrategyV1Alpha1Resource() resource.Resource {
	return &KumaIoMeshLoadBalancingStrategyV1Alpha1Resource{}
}

type KumaIoMeshLoadBalancingStrategyV1Alpha1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type KumaIoMeshLoadBalancingStrategyV1Alpha1ResourceData struct {
	ID             types.String `tfsdk:"id" json:"-"`
	ForceConflicts types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager   types.String `tfsdk:"field_manager" json:"-"`
	WaitForUpsert  types.List   `tfsdk:"wait_for_upsert" json:"-"`
	WaitForDelete  types.Object `tfsdk:"wait_for_delete" json:"-"`

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

func (r *KumaIoMeshLoadBalancingStrategyV1Alpha1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kuma_io_mesh_load_balancing_strategy_v1alpha1"
}

func (r *KumaIoMeshLoadBalancingStrategyV1Alpha1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
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

			"force_conflicts": schema.BoolAttribute{
				Description:         "If 'true', server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "If `true`, server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"field_manager": schema.BoolAttribute{
				Description:         "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"wait_for_upsert": schema.ListNestedAttribute{
				Description:         "Wait for specific conditions after create/update of resources.",
				MarkdownDescription: "Wait for specific conditions after create/update of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"jsonpath": schema.StringAttribute{
							Description:         "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							MarkdownDescription: "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							Required:            true,
							Optional:            false,
							Computed:            false,
						},
						"value": schema.StringAttribute{
							Description:         "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							MarkdownDescription: "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							Required:            false,
							Optional:            true,
							Computed:            true,
						},
						"timeout": schema.StringAttribute{
							Description:         "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							MarkdownDescription: "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("30s"),
						},
						"poll_interval": schema.StringAttribute{
							Description:         "The length of time to wait before checking again.",
							MarkdownDescription: "The length of time to wait before checking again.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("5s"),
						},
					},
				},
			},

			"wait_for_delete": schema.SingleNestedAttribute{
				Description:         "Wait for deletion of resources.",
				MarkdownDescription: "Wait for deletion of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"timeout": schema.StringAttribute{
						Description:         "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
						MarkdownDescription: "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             stringdefault.StaticString("30s"),
					},
					"poll_interval": schema.StringAttribute{
						Description:         "The length of time to wait before checking again.",
						MarkdownDescription: "The length of time to wait before checking again.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             stringdefault.StaticString("5s"),
					},
				},
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},

					"labels": schema.MapAttribute{
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            true,
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
						Computed:            true,
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
																				Description:         "The name of the Object in the per-request filterState, which is an Envoy::Hashable object. If there is no data associated with the key, or the stored object is not Envoy::Hashable, no hash will be produced.",
																				MarkdownDescription: "The name of the Object in the per-request filterState, which is an Envoy::Hashable object. If there is no data associated with the key, or the stored object is not Envoy::Hashable, no hash will be produced.",
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
																				Description:         "The name of the URL query parameter that will be used to obtain the hash key. If the parameter is not present, no hash will be produced. Query parameter names are case-sensitive.",
																				MarkdownDescription: "The name of the URL query parameter that will be used to obtain the hash key. If the parameter is not present, no hash will be produced. Query parameter names are case-sensitive.",
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
																		Description:         "Terminal is a flag that short-circuits the hash computing. This field provides a ‘fallback’ style of configuration: “if a terminal policy doesn’t work, fallback to rest of the policy list”, it saves time when the terminal policy works. If true, and there is already a hash computed, ignore rest of the list of hash polices.",
																		MarkdownDescription: "Terminal is a flag that short-circuits the hash computing. This field provides a ‘fallback’ style of configuration: “if a terminal policy doesn’t work, fallback to rest of the policy list”, it saves time when the terminal policy works. If true, and there is already a hash computed, ignore rest of the list of hash polices.",
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
															Description:         "The table size for Maglev hashing. Maglev aims for “minimal disruption” rather than an absolute guarantee. Minimal disruption means that when the set of upstream hosts change, a connection will likely be sent to the same upstream as it was before. Increasing the table size reduces the amount of disruption. The table size must be prime number limited to 5000011. If it is not specified, the default is 65537.",
															MarkdownDescription: "The table size for Maglev hashing. Maglev aims for “minimal disruption” rather than an absolute guarantee. Minimal disruption means that when the set of upstream hosts change, a connection will likely be sent to the same upstream as it was before. Increasing the table size reduces the amount of disruption. The table size must be prime number limited to 5000011. If it is not specified, the default is 65537.",
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
													Description:         "Random selects a random available host. The random load balancer generally performs better than round-robin if no health checking policy is configured. Random selection avoids bias towards the host in the set that comes after a failed host.",
													MarkdownDescription: "Random selects a random available host. The random load balancer generally performs better than round-robin if no health checking policy is configured. Random selection avoids bias towards the host in the set that comes after a failed host.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"ring_hash": schema.SingleNestedAttribute{
													Description:         "RingHash  implements consistent hashing to upstream hosts. Each host is mapped onto a circle (the “ring”) by hashing its address; each request is then routed to a host by hashing some property of the request, and finding the nearest corresponding host clockwise around the ring.",
													MarkdownDescription: "RingHash  implements consistent hashing to upstream hosts. Each host is mapped onto a circle (the “ring”) by hashing its address; each request is then routed to a host by hashing some property of the request, and finding the nearest corresponding host clockwise around the ring.",
													Attributes: map[string]schema.Attribute{
														"hash_function": schema.StringAttribute{
															Description:         "HashFunction is a function used to hash hosts onto the ketama ring. The value defaults to XX_HASH. Available values – XX_HASH, MURMUR_HASH_2.",
															MarkdownDescription: "HashFunction is a function used to hash hosts onto the ketama ring. The value defaults to XX_HASH. Available values – XX_HASH, MURMUR_HASH_2.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("XXHash", "MurmurHash2"),
															},
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
																				Description:         "The name of the Object in the per-request filterState, which is an Envoy::Hashable object. If there is no data associated with the key, or the stored object is not Envoy::Hashable, no hash will be produced.",
																				MarkdownDescription: "The name of the Object in the per-request filterState, which is an Envoy::Hashable object. If there is no data associated with the key, or the stored object is not Envoy::Hashable, no hash will be produced.",
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
																				Description:         "The name of the URL query parameter that will be used to obtain the hash key. If the parameter is not present, no hash will be produced. Query parameter names are case-sensitive.",
																				MarkdownDescription: "The name of the URL query parameter that will be used to obtain the hash key. If the parameter is not present, no hash will be produced. Query parameter names are case-sensitive.",
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
																		Description:         "Terminal is a flag that short-circuits the hash computing. This field provides a ‘fallback’ style of configuration: “if a terminal policy doesn’t work, fallback to rest of the policy list”, it saves time when the terminal policy works. If true, and there is already a hash computed, ignore rest of the list of hash polices.",
																		MarkdownDescription: "Terminal is a flag that short-circuits the hash computing. This field provides a ‘fallback’ style of configuration: “if a terminal policy doesn’t work, fallback to rest of the policy list”, it saves time when the terminal policy works. If true, and there is already a hash computed, ignore rest of the list of hash polices.",
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
															Description:         "Maximum hash ring size. Defaults to 8M entries, and limited to 8M entries, but can be lowered to further constrain resource use.",
															MarkdownDescription: "Maximum hash ring size. Defaults to 8M entries, and limited to 8M entries, but can be lowered to further constrain resource use.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(1),
																int64validator.AtMost(8e+06),
															},
														},

														"min_ring_size": schema.Int64Attribute{
															Description:         "Minimum hash ring size. The larger the ring is (that is, the more hashes there are for each provided host) the better the request distribution will reflect the desired weights. Defaults to 1024 entries, and limited to 8M entries.",
															MarkdownDescription: "Minimum hash ring size. The larger the ring is (that is, the more hashes there are for each provided host) the better the request distribution will reflect the desired weights. Defaults to 1024 entries, and limited to 8M entries.",
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
													Description:         "RoundRobin is a load balancing algorithm that distributes requests across available upstream hosts in round-robin order.",
													MarkdownDescription: "RoundRobin is a load balancing algorithm that distributes requests across available upstream hosts in round-robin order.",
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
												"disabled": schema.BoolAttribute{
													Description:         "Disabled allows to disable locality-aware load balancing. When disabled requests are distributed across all endpoints regardless of locality.",
													MarkdownDescription: "Disabled allows to disable locality-aware load balancing. When disabled requests are distributed across all endpoints regardless of locality.",
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

func (r *KumaIoMeshLoadBalancingStrategyV1Alpha1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedResourceDataError(request.ProviderData))
	}
}

func (r *KumaIoMeshLoadBalancingStrategyV1Alpha1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_kuma_io_mesh_load_balancing_strategy_v1alpha1")

	var model KumaIoMeshLoadBalancingStrategyV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("kuma.io/v1alpha1")
	model.Kind = pointer.String("MeshLoadBalancingStrategy")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonMarshalError(err))
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "kuma.io", Version: "v1alpha1", Resource: "meshloadbalancingstrategies"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.Append(utilities.PatchError(err))
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse KumaIoMeshLoadBalancingStrategyV1Alpha1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *KumaIoMeshLoadBalancingStrategyV1Alpha1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kuma_io_mesh_load_balancing_strategy_v1alpha1")

	var data KumaIoMeshLoadBalancingStrategyV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "kuma.io", Version: "v1alpha1", Resource: "meshloadbalancingstrategies"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetNamespacedResourceError(err, data.Metadata.Name, data.Metadata.Namespace))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse KumaIoMeshLoadBalancingStrategyV1Alpha1ResourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *KumaIoMeshLoadBalancingStrategyV1Alpha1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_kuma_io_mesh_load_balancing_strategy_v1alpha1")

	var model KumaIoMeshLoadBalancingStrategyV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("kuma.io/v1alpha1")
	model.Kind = pointer.String("MeshLoadBalancingStrategy")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonMarshalError(err))
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "kuma.io", Version: "v1alpha1", Resource: "meshloadbalancingstrategies"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.Append(utilities.PatchError(err))
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse KumaIoMeshLoadBalancingStrategyV1Alpha1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *KumaIoMeshLoadBalancingStrategyV1Alpha1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_kuma_io_mesh_load_balancing_strategy_v1alpha1")

	var data KumaIoMeshLoadBalancingStrategyV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "kuma.io", Version: "v1alpha1", Resource: "meshloadbalancingstrategies"}).
		Namespace(data.Metadata.Namespace).
		Delete(ctx, data.Metadata.Name, meta.DeleteOptions{})
	if utilities.IsDeletionError(err) {
		response.Diagnostics.Append(utilities.DeleteError(err))
		return
	}

	if !data.WaitForDelete.IsNull() {
		timeout := utilities.DetermineTimeout(data.WaitForDelete.Attributes())
		pollInterval := utilities.DeterminePollInterval(data.WaitForDelete.Attributes())

		startTime := time.Now()
		for {
			_, err := r.kubernetesClient.
				Resource(k8sSchema.GroupVersionResource{Group: "kuma.io", Version: "v1alpha1", Resource: "meshloadbalancingstrategies"}).
				Namespace(data.Metadata.Namespace).
				Get(ctx, data.Metadata.Name, meta.GetOptions{})
			if utilities.IsNotFound(err) || timeout == time.Second*0 {
				break
			}
			if time.Now().After(startTime.Add(timeout)) {
				response.Diagnostics.Append(utilities.WaitTimeoutExceeded())
				return
			}
			time.Sleep(pollInterval)
		}
	}
}

func (r *KumaIoMeshLoadBalancingStrategyV1Alpha1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	idParts := strings.Split(request.ID, "/")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'namespace/name' Got: '%q'", request.ID),
		)
		return
	}

	namespace := idParts[0]
	name := idParts[1]
	tflog.Trace(ctx, "parsed import ID", map[string]interface{}{
		"namespace": namespace,
		"name":      name,
	})
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("namespace"), namespace)...)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("name"), name)...)
}
