/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package opensearch_opster_io_v1

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
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &OpensearchOpsterIoOpenSearchIsmpolicyV1Manifest{}
)

func NewOpensearchOpsterIoOpenSearchIsmpolicyV1Manifest() datasource.DataSource {
	return &OpensearchOpsterIoOpenSearchIsmpolicyV1Manifest{}
}

type OpensearchOpsterIoOpenSearchIsmpolicyV1Manifest struct{}

type OpensearchOpsterIoOpenSearchIsmpolicyV1ManifestData struct {
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
		ApplyToExistingIndices *bool   `tfsdk:"apply_to_existing_indices" json:"applyToExistingIndices,omitempty"`
		DefaultState           *string `tfsdk:"default_state" json:"defaultState,omitempty"`
		Description            *string `tfsdk:"description" json:"description,omitempty"`
		ErrorNotification      *struct {
			Channel     *string `tfsdk:"channel" json:"channel,omitempty"`
			Destination *struct {
				Amazon *struct {
					Url *string `tfsdk:"url" json:"url,omitempty"`
				} `tfsdk:"amazon" json:"amazon,omitempty"`
				Chime *struct {
					Url *string `tfsdk:"url" json:"url,omitempty"`
				} `tfsdk:"chime" json:"chime,omitempty"`
				CustomWebhook *struct {
					Url *string `tfsdk:"url" json:"url,omitempty"`
				} `tfsdk:"custom_webhook" json:"customWebhook,omitempty"`
				Slack *struct {
					Url *string `tfsdk:"url" json:"url,omitempty"`
				} `tfsdk:"slack" json:"slack,omitempty"`
			} `tfsdk:"destination" json:"destination,omitempty"`
			MessageTemplate *struct {
				Source *string `tfsdk:"source" json:"source,omitempty"`
			} `tfsdk:"message_template" json:"messageTemplate,omitempty"`
		} `tfsdk:"error_notification" json:"errorNotification,omitempty"`
		IsmTemplate *struct {
			IndexPatterns *[]string `tfsdk:"index_patterns" json:"indexPatterns,omitempty"`
			Priority      *int64    `tfsdk:"priority" json:"priority,omitempty"`
		} `tfsdk:"ism_template" json:"ismTemplate,omitempty"`
		OpensearchCluster *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"opensearch_cluster" json:"opensearchCluster,omitempty"`
		PolicyId *string `tfsdk:"policy_id" json:"policyId,omitempty"`
		States   *[]struct {
			Actions *[]struct {
				Alias *struct {
					Actions *[]struct {
						Add *struct {
							Aliases      *[]string `tfsdk:"aliases" json:"aliases,omitempty"`
							Index        *string   `tfsdk:"index" json:"index,omitempty"`
							IsWriteIndex *bool     `tfsdk:"is_write_index" json:"isWriteIndex,omitempty"`
							Routing      *string   `tfsdk:"routing" json:"routing,omitempty"`
						} `tfsdk:"add" json:"add,omitempty"`
						Remove *struct {
							Aliases      *[]string `tfsdk:"aliases" json:"aliases,omitempty"`
							Index        *string   `tfsdk:"index" json:"index,omitempty"`
							IsWriteIndex *bool     `tfsdk:"is_write_index" json:"isWriteIndex,omitempty"`
							Routing      *string   `tfsdk:"routing" json:"routing,omitempty"`
						} `tfsdk:"remove" json:"remove,omitempty"`
					} `tfsdk:"actions" json:"actions,omitempty"`
				} `tfsdk:"alias" json:"alias,omitempty"`
				Allocation *struct {
					Exclude *string `tfsdk:"exclude" json:"exclude,omitempty"`
					Include *string `tfsdk:"include" json:"include,omitempty"`
					Require *string `tfsdk:"require" json:"require,omitempty"`
					WaitFor *string `tfsdk:"wait_for" json:"waitFor,omitempty"`
				} `tfsdk:"allocation" json:"allocation,omitempty"`
				Close      *map[string]string `tfsdk:"close" json:"close,omitempty"`
				Delete     *map[string]string `tfsdk:"delete" json:"delete,omitempty"`
				ForceMerge *struct {
					MaxNumSegments *int64 `tfsdk:"max_num_segments" json:"maxNumSegments,omitempty"`
				} `tfsdk:"force_merge" json:"forceMerge,omitempty"`
				IndexPriority *struct {
					Priority *int64 `tfsdk:"priority" json:"priority,omitempty"`
				} `tfsdk:"index_priority" json:"indexPriority,omitempty"`
				Notification *struct {
					Destination     *string `tfsdk:"destination" json:"destination,omitempty"`
					MessageTemplate *struct {
						Source *string `tfsdk:"source" json:"source,omitempty"`
					} `tfsdk:"message_template" json:"messageTemplate,omitempty"`
				} `tfsdk:"notification" json:"notification,omitempty"`
				Open         *map[string]string `tfsdk:"open" json:"open,omitempty"`
				ReadOnly     *map[string]string `tfsdk:"read_only" json:"readOnly,omitempty"`
				ReadWrite    *map[string]string `tfsdk:"read_write" json:"readWrite,omitempty"`
				ReplicaCount *struct {
					NumberOfReplicas *int64 `tfsdk:"number_of_replicas" json:"numberOfReplicas,omitempty"`
				} `tfsdk:"replica_count" json:"replicaCount,omitempty"`
				Retry *struct {
					Backoff *string `tfsdk:"backoff" json:"backoff,omitempty"`
					Count   *int64  `tfsdk:"count" json:"count,omitempty"`
					Delay   *string `tfsdk:"delay" json:"delay,omitempty"`
				} `tfsdk:"retry" json:"retry,omitempty"`
				Rollover *struct {
					MinDocCount         *int64  `tfsdk:"min_doc_count" json:"minDocCount,omitempty"`
					MinIndexAge         *string `tfsdk:"min_index_age" json:"minIndexAge,omitempty"`
					MinPrimaryShardSize *string `tfsdk:"min_primary_shard_size" json:"minPrimaryShardSize,omitempty"`
					MinSize             *string `tfsdk:"min_size" json:"minSize,omitempty"`
				} `tfsdk:"rollover" json:"rollover,omitempty"`
				Rollup *map[string]string `tfsdk:"rollup" json:"rollup,omitempty"`
				Shrink *struct {
					ForceUnsafe              *bool   `tfsdk:"force_unsafe" json:"forceUnsafe,omitempty"`
					MaxShardSize             *string `tfsdk:"max_shard_size" json:"maxShardSize,omitempty"`
					NumNewShards             *int64  `tfsdk:"num_new_shards" json:"numNewShards,omitempty"`
					PercentageOfSourceShards *int64  `tfsdk:"percentage_of_source_shards" json:"percentageOfSourceShards,omitempty"`
					TargetIndexNameTemplate  *string `tfsdk:"target_index_name_template" json:"targetIndexNameTemplate,omitempty"`
				} `tfsdk:"shrink" json:"shrink,omitempty"`
				Snapshot *struct {
					Repository *string `tfsdk:"repository" json:"repository,omitempty"`
					Snapshot   *string `tfsdk:"snapshot" json:"snapshot,omitempty"`
				} `tfsdk:"snapshot" json:"snapshot,omitempty"`
				Timeout *string `tfsdk:"timeout" json:"timeout,omitempty"`
			} `tfsdk:"actions" json:"actions,omitempty"`
			Name        *string `tfsdk:"name" json:"name,omitempty"`
			Transitions *[]struct {
				Conditions *struct {
					Cron *struct {
						Cron *struct {
							Expression *string `tfsdk:"expression" json:"expression,omitempty"`
							Timezone   *string `tfsdk:"timezone" json:"timezone,omitempty"`
						} `tfsdk:"cron" json:"cron,omitempty"`
					} `tfsdk:"cron" json:"cron,omitempty"`
					MinDocCount    *int64  `tfsdk:"min_doc_count" json:"minDocCount,omitempty"`
					MinIndexAge    *string `tfsdk:"min_index_age" json:"minIndexAge,omitempty"`
					MinRolloverAge *string `tfsdk:"min_rollover_age" json:"minRolloverAge,omitempty"`
					MinSize        *string `tfsdk:"min_size" json:"minSize,omitempty"`
				} `tfsdk:"conditions" json:"conditions,omitempty"`
				StateName *string `tfsdk:"state_name" json:"stateName,omitempty"`
			} `tfsdk:"transitions" json:"transitions,omitempty"`
		} `tfsdk:"states" json:"states,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OpensearchOpsterIoOpenSearchIsmpolicyV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_opensearch_opster_io_open_search_ism_policy_v1_manifest"
}

func (r *OpensearchOpsterIoOpenSearchIsmpolicyV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "ISMPolicySpec is the specification for the ISM policy for OS.",
				MarkdownDescription: "ISMPolicySpec is the specification for the ISM policy for OS.",
				Attributes: map[string]schema.Attribute{
					"apply_to_existing_indices": schema.BoolAttribute{
						Description:         "If true, apply the policy to existing indices that match the index patterns in the ISM template.",
						MarkdownDescription: "If true, apply the policy to existing indices that match the index patterns in the ISM template.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"default_state": schema.StringAttribute{
						Description:         "The default starting state for each index that uses this policy.",
						MarkdownDescription: "The default starting state for each index that uses this policy.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"description": schema.StringAttribute{
						Description:         "A human-readable description of the policy.",
						MarkdownDescription: "A human-readable description of the policy.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"error_notification": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"channel": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"destination": schema.SingleNestedAttribute{
								Description:         "The destination URL.",
								MarkdownDescription: "The destination URL.",
								Attributes: map[string]schema.Attribute{
									"amazon": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"url": schema.StringAttribute{
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

									"chime": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"url": schema.StringAttribute{
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

									"custom_webhook": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"url": schema.StringAttribute{
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

									"slack": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"url": schema.StringAttribute{
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"message_template": schema.SingleNestedAttribute{
								Description:         "The text of the message",
								MarkdownDescription: "The text of the message",
								Attributes: map[string]schema.Attribute{
									"source": schema.StringAttribute{
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"ism_template": schema.SingleNestedAttribute{
						Description:         "Specify an ISM template pattern that matches the index to apply the policy.",
						MarkdownDescription: "Specify an ISM template pattern that matches the index to apply the policy.",
						Attributes: map[string]schema.Attribute{
							"index_patterns": schema.ListAttribute{
								Description:         "Index patterns on which this policy has to be applied",
								MarkdownDescription: "Index patterns on which this policy has to be applied",
								ElementType:         types.StringType,
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"priority": schema.Int64Attribute{
								Description:         "Priority of the template, defaults to 0",
								MarkdownDescription: "Priority of the template, defaults to 0",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"opensearch_cluster": schema.SingleNestedAttribute{
						Description:         "LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.",
						MarkdownDescription: "LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"policy_id": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"states": schema.ListNestedAttribute{
						Description:         "The states that you define in the policy.",
						MarkdownDescription: "The states that you define in the policy.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"actions": schema.ListNestedAttribute{
									Description:         "The actions to execute after entering a state.",
									MarkdownDescription: "The actions to execute after entering a state.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"alias": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"actions": schema.ListNestedAttribute{
														Description:         "Allocate the index to a node with a specified attribute.",
														MarkdownDescription: "Allocate the index to a node with a specified attribute.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"add": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"aliases": schema.ListAttribute{
																			Description:         "The name of the alias.",
																			MarkdownDescription: "The name of the alias.",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"index": schema.StringAttribute{
																			Description:         "The name of the index that the alias points to.",
																			MarkdownDescription: "The name of the index that the alias points to.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"is_write_index": schema.BoolAttribute{
																			Description:         "Specify the index that accepts any write operations to the alias.",
																			MarkdownDescription: "Specify the index that accepts any write operations to the alias.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"routing": schema.StringAttribute{
																			Description:         "Limit search to an associated shard value",
																			MarkdownDescription: "Limit search to an associated shard value",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"remove": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"aliases": schema.ListAttribute{
																			Description:         "The name of the alias.",
																			MarkdownDescription: "The name of the alias.",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"index": schema.StringAttribute{
																			Description:         "The name of the index that the alias points to.",
																			MarkdownDescription: "The name of the index that the alias points to.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"is_write_index": schema.BoolAttribute{
																			Description:         "Specify the index that accepts any write operations to the alias.",
																			MarkdownDescription: "Specify the index that accepts any write operations to the alias.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"routing": schema.StringAttribute{
																			Description:         "Limit search to an associated shard value",
																			MarkdownDescription: "Limit search to an associated shard value",
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
														Required: true,
														Optional: false,
														Computed: false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"allocation": schema.SingleNestedAttribute{
												Description:         "Allocate the index to a node with a specific attribute set",
												MarkdownDescription: "Allocate the index to a node with a specific attribute set",
												Attributes: map[string]schema.Attribute{
													"exclude": schema.StringAttribute{
														Description:         "Allocate the index to a node with a specified attribute.",
														MarkdownDescription: "Allocate the index to a node with a specified attribute.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"include": schema.StringAttribute{
														Description:         "Allocate the index to a node with any of the specified attributes.",
														MarkdownDescription: "Allocate the index to a node with any of the specified attributes.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"require": schema.StringAttribute{
														Description:         "Don’t allocate the index to a node with any of the specified attributes.",
														MarkdownDescription: "Don’t allocate the index to a node with any of the specified attributes.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"wait_for": schema.StringAttribute{
														Description:         "Wait for the policy to execute before allocating the index to a node with a specified attribute.",
														MarkdownDescription: "Wait for the policy to execute before allocating the index to a node with a specified attribute.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"close": schema.MapAttribute{
												Description:         "Closes the managed index.",
												MarkdownDescription: "Closes the managed index.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"delete": schema.MapAttribute{
												Description:         "Deletes a managed index.",
												MarkdownDescription: "Deletes a managed index.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"force_merge": schema.SingleNestedAttribute{
												Description:         "Reduces the number of Lucene segments by merging the segments of individual shards.",
												MarkdownDescription: "Reduces the number of Lucene segments by merging the segments of individual shards.",
												Attributes: map[string]schema.Attribute{
													"max_num_segments": schema.Int64Attribute{
														Description:         "The number of segments to reduce the shard to.",
														MarkdownDescription: "The number of segments to reduce the shard to.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"index_priority": schema.SingleNestedAttribute{
												Description:         "Set the priority for the index in a specific state.",
												MarkdownDescription: "Set the priority for the index in a specific state.",
												Attributes: map[string]schema.Attribute{
													"priority": schema.Int64Attribute{
														Description:         "The priority for the index as soon as it enters a state.",
														MarkdownDescription: "The priority for the index as soon as it enters a state.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"notification": schema.SingleNestedAttribute{
												Description:         "Name string 'json:'name,omitempty''",
												MarkdownDescription: "Name string 'json:'name,omitempty''",
												Attributes: map[string]schema.Attribute{
													"destination": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"message_template": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"source": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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

											"open": schema.MapAttribute{
												Description:         "Opens a managed index.",
												MarkdownDescription: "Opens a managed index.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"read_only": schema.MapAttribute{
												Description:         "Sets a managed index to be read only.",
												MarkdownDescription: "Sets a managed index to be read only.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"read_write": schema.MapAttribute{
												Description:         "Sets a managed index to be writeable.",
												MarkdownDescription: "Sets a managed index to be writeable.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"replica_count": schema.SingleNestedAttribute{
												Description:         "Sets the number of replicas to assign to an index.",
												MarkdownDescription: "Sets the number of replicas to assign to an index.",
												Attributes: map[string]schema.Attribute{
													"number_of_replicas": schema.Int64Attribute{
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

											"retry": schema.SingleNestedAttribute{
												Description:         "The retry configuration for the action.",
												MarkdownDescription: "The retry configuration for the action.",
												Attributes: map[string]schema.Attribute{
													"backoff": schema.StringAttribute{
														Description:         "The backoff policy type to use when retrying.",
														MarkdownDescription: "The backoff policy type to use when retrying.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"count": schema.Int64Attribute{
														Description:         "The number of retry counts.",
														MarkdownDescription: "The number of retry counts.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"delay": schema.StringAttribute{
														Description:         "The time to wait between retries.",
														MarkdownDescription: "The time to wait between retries.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"rollover": schema.SingleNestedAttribute{
												Description:         "Rolls an alias over to a new index when the managed index meets one of the rollover conditions.",
												MarkdownDescription: "Rolls an alias over to a new index when the managed index meets one of the rollover conditions.",
												Attributes: map[string]schema.Attribute{
													"min_doc_count": schema.Int64Attribute{
														Description:         "The minimum number of documents required to roll over the index.",
														MarkdownDescription: "The minimum number of documents required to roll over the index.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"min_index_age": schema.StringAttribute{
														Description:         "The minimum age required to roll over the index.",
														MarkdownDescription: "The minimum age required to roll over the index.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"min_primary_shard_size": schema.StringAttribute{
														Description:         "The minimum storage size of a single primary shard required to roll over the index.",
														MarkdownDescription: "The minimum storage size of a single primary shard required to roll over the index.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"min_size": schema.StringAttribute{
														Description:         "The minimum size of the total primary shard storage (not counting replicas) required to roll over the index.",
														MarkdownDescription: "The minimum size of the total primary shard storage (not counting replicas) required to roll over the index.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"rollup": schema.MapAttribute{
												Description:         "Periodically reduce data granularity by rolling up old data into summarized indexes.",
												MarkdownDescription: "Periodically reduce data granularity by rolling up old data into summarized indexes.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"shrink": schema.SingleNestedAttribute{
												Description:         "Allows you to reduce the number of primary shards in your indexes",
												MarkdownDescription: "Allows you to reduce the number of primary shards in your indexes",
												Attributes: map[string]schema.Attribute{
													"force_unsafe": schema.BoolAttribute{
														Description:         "If true, executes the shrink action even if there are no replicas.",
														MarkdownDescription: "If true, executes the shrink action even if there are no replicas.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"max_shard_size": schema.StringAttribute{
														Description:         "The maximum size in bytes of a shard for the target index.",
														MarkdownDescription: "The maximum size in bytes of a shard for the target index.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"num_new_shards": schema.Int64Attribute{
														Description:         "The maximum number of primary shards in the shrunken index.",
														MarkdownDescription: "The maximum number of primary shards in the shrunken index.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"percentage_of_source_shards": schema.Int64Attribute{
														Description:         "Percentage of the number of original primary shards to shrink.",
														MarkdownDescription: "Percentage of the number of original primary shards to shrink.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"target_index_name_template": schema.StringAttribute{
														Description:         "The name of the shrunken index.",
														MarkdownDescription: "The name of the shrunken index.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"snapshot": schema.SingleNestedAttribute{
												Description:         "Back up your cluster’s indexes and state",
												MarkdownDescription: "Back up your cluster’s indexes and state",
												Attributes: map[string]schema.Attribute{
													"repository": schema.StringAttribute{
														Description:         "The repository name that you register through the native snapshot API operations.",
														MarkdownDescription: "The repository name that you register through the native snapshot API operations.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"snapshot": schema.StringAttribute{
														Description:         "The name of the snapshot.",
														MarkdownDescription: "The name of the snapshot.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"timeout": schema.StringAttribute{
												Description:         "The timeout period for the action. Accepts time units for minutes, hours, and days.",
												MarkdownDescription: "The timeout period for the action. Accepts time units for minutes, hours, and days.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},

								"name": schema.StringAttribute{
									Description:         "The name of the state.",
									MarkdownDescription: "The name of the state.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"transitions": schema.ListNestedAttribute{
									Description:         "The next states and the conditions required to transition to those states. If no transitions exist, the policy assumes that it’s complete and can now stop managing the index",
									MarkdownDescription: "The next states and the conditions required to transition to those states. If no transitions exist, the policy assumes that it’s complete and can now stop managing the index",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"conditions": schema.SingleNestedAttribute{
												Description:         "conditions for the transition.",
												MarkdownDescription: "conditions for the transition.",
												Attributes: map[string]schema.Attribute{
													"cron": schema.SingleNestedAttribute{
														Description:         "The cron job that triggers the transition if no other transition happens first.",
														MarkdownDescription: "The cron job that triggers the transition if no other transition happens first.",
														Attributes: map[string]schema.Attribute{
															"cron": schema.SingleNestedAttribute{
																Description:         "A wrapper for the cron job that triggers the transition if no other transition happens first. This wrapper is here to adhere to the OpenSearch API.",
																MarkdownDescription: "A wrapper for the cron job that triggers the transition if no other transition happens first. This wrapper is here to adhere to the OpenSearch API.",
																Attributes: map[string]schema.Attribute{
																	"expression": schema.StringAttribute{
																		Description:         "The cron expression that triggers the transition.",
																		MarkdownDescription: "The cron expression that triggers the transition.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"timezone": schema.StringAttribute{
																		Description:         "The timezone that triggers the transition.",
																		MarkdownDescription: "The timezone that triggers the transition.",
																		Required:            true,
																		Optional:            false,
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

													"min_doc_count": schema.Int64Attribute{
														Description:         "The minimum document count of the index required to transition.",
														MarkdownDescription: "The minimum document count of the index required to transition.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"min_index_age": schema.StringAttribute{
														Description:         "The minimum age of the index required to transition.",
														MarkdownDescription: "The minimum age of the index required to transition.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"min_rollover_age": schema.StringAttribute{
														Description:         "The minimum age required after a rollover has occurred to transition to the next state.",
														MarkdownDescription: "The minimum age required after a rollover has occurred to transition to the next state.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"min_size": schema.StringAttribute{
														Description:         "The minimum size of the total primary shard storage (not counting replicas) required to transition.",
														MarkdownDescription: "The minimum size of the total primary shard storage (not counting replicas) required to transition.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: true,
												Optional: false,
												Computed: false,
											},

											"state_name": schema.StringAttribute{
												Description:         "The name of the state to transition to if the conditions are met.",
												MarkdownDescription: "The name of the state to transition to if the conditions are met.",
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
		},
	}
}

func (r *OpensearchOpsterIoOpenSearchIsmpolicyV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_opensearch_opster_io_open_search_ism_policy_v1_manifest")

	var model OpensearchOpsterIoOpenSearchIsmpolicyV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("opensearch.opster.io/v1")
	model.Kind = pointer.String("OpenSearchISMPolicy")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
