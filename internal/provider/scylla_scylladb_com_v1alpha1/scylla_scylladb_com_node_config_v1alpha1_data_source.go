/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package scylla_scylladb_com_v1alpha1

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
	_ datasource.DataSource              = &ScyllaScylladbComNodeConfigV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &ScyllaScylladbComNodeConfigV1Alpha1DataSource{}
)

func NewScyllaScylladbComNodeConfigV1Alpha1DataSource() datasource.DataSource {
	return &ScyllaScylladbComNodeConfigV1Alpha1DataSource{}
}

type ScyllaScylladbComNodeConfigV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type ScyllaScylladbComNodeConfigV1Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		DisableOptimizations *bool `tfsdk:"disable_optimizations" json:"disableOptimizations,omitempty"`
		LocalDiskSetup       *struct {
			Filesystems *[]struct {
				Device *string `tfsdk:"device" json:"device,omitempty"`
				Type   *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"filesystems" json:"filesystems,omitempty"`
			Mounts *[]struct {
				Device             *string   `tfsdk:"device" json:"device,omitempty"`
				FsType             *string   `tfsdk:"fs_type" json:"fsType,omitempty"`
				MountPoint         *string   `tfsdk:"mount_point" json:"mountPoint,omitempty"`
				UnsupportedOptions *[]string `tfsdk:"unsupported_options" json:"unsupportedOptions,omitempty"`
			} `tfsdk:"mounts" json:"mounts,omitempty"`
			Raids *[]struct {
				RAID0 *struct {
					Devices *struct {
						ModelRegex *string `tfsdk:"model_regex" json:"modelRegex,omitempty"`
						NameRegex  *string `tfsdk:"name_regex" json:"nameRegex,omitempty"`
					} `tfsdk:"devices" json:"devices,omitempty"`
				} `tfsdk:"raid0" json:"RAID0,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"raids" json:"raids,omitempty"`
		} `tfsdk:"local_disk_setup" json:"localDiskSetup,omitempty"`
		Placement *struct {
			Affinity *struct {
				NodeAffinity *struct {
					PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
						Preference *struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchFields *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_fields" json:"matchFields,omitempty"`
						} `tfsdk:"preference" json:"preference,omitempty"`
						Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
					} `tfsdk:"preferred_during_scheduling_ignored_during_execution" json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
					RequiredDuringSchedulingIgnoredDuringExecution *struct {
						NodeSelectorTerms *[]struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchFields *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_fields" json:"matchFields,omitempty"`
						} `tfsdk:"node_selector_terms" json:"nodeSelectorTerms,omitempty"`
					} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
				} `tfsdk:"node_affinity" json:"nodeAffinity,omitempty"`
				PodAffinity *struct {
					PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
						PodAffinityTerm *struct {
							LabelSelector *struct {
								MatchExpressions *[]struct {
									Key      *string   `tfsdk:"key" json:"key,omitempty"`
									Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
									Values   *[]string `tfsdk:"values" json:"values,omitempty"`
								} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
								MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
							} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
							NamespaceSelector *struct {
								MatchExpressions *[]struct {
									Key      *string   `tfsdk:"key" json:"key,omitempty"`
									Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
									Values   *[]string `tfsdk:"values" json:"values,omitempty"`
								} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
								MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
							} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
							Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
							TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
						} `tfsdk:"pod_affinity_term" json:"podAffinityTerm,omitempty"`
						Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
					} `tfsdk:"preferred_during_scheduling_ignored_during_execution" json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
					RequiredDuringSchedulingIgnoredDuringExecution *[]struct {
						LabelSelector *struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
						} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
						NamespaceSelector *struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
						} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
						Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
						TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
					} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
				} `tfsdk:"pod_affinity" json:"podAffinity,omitempty"`
				PodAntiAffinity *struct {
					PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
						PodAffinityTerm *struct {
							LabelSelector *struct {
								MatchExpressions *[]struct {
									Key      *string   `tfsdk:"key" json:"key,omitempty"`
									Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
									Values   *[]string `tfsdk:"values" json:"values,omitempty"`
								} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
								MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
							} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
							NamespaceSelector *struct {
								MatchExpressions *[]struct {
									Key      *string   `tfsdk:"key" json:"key,omitempty"`
									Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
									Values   *[]string `tfsdk:"values" json:"values,omitempty"`
								} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
								MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
							} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
							Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
							TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
						} `tfsdk:"pod_affinity_term" json:"podAffinityTerm,omitempty"`
						Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
					} `tfsdk:"preferred_during_scheduling_ignored_during_execution" json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
					RequiredDuringSchedulingIgnoredDuringExecution *[]struct {
						LabelSelector *struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
						} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
						NamespaceSelector *struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
						} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
						Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
						TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
					} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
				} `tfsdk:"pod_anti_affinity" json:"podAntiAffinity,omitempty"`
			} `tfsdk:"affinity" json:"affinity,omitempty"`
			NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
			Tolerations  *[]struct {
				Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
				Key               *string `tfsdk:"key" json:"key,omitempty"`
				Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
				TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
				Value             *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"tolerations" json:"tolerations,omitempty"`
		} `tfsdk:"placement" json:"placement,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ScyllaScylladbComNodeConfigV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_scylla_scylladb_com_node_config_v1alpha1"
}

func (r *ScyllaScylladbComNodeConfigV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"disable_optimizations": schema.BoolAttribute{
						Description:         "disableOptimizations controls if nodes matching placement requirements are going to be optimized. Turning off optimizations on already optimized Nodes does not revert changes.",
						MarkdownDescription: "disableOptimizations controls if nodes matching placement requirements are going to be optimized. Turning off optimizations on already optimized Nodes does not revert changes.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"local_disk_setup": schema.SingleNestedAttribute{
						Description:         "localDiskSetup contains options of automatic local disk setup.",
						MarkdownDescription: "localDiskSetup contains options of automatic local disk setup.",
						Attributes: map[string]schema.Attribute{
							"filesystems": schema.ListNestedAttribute{
								Description:         "filesystems is a list of filesystem configurations.",
								MarkdownDescription: "filesystems is a list of filesystem configurations.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"device": schema.StringAttribute{
											Description:         "device is a path to the device where the desired filesystem should be created.",
											MarkdownDescription: "device is a path to the device where the desired filesystem should be created.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"type": schema.StringAttribute{
											Description:         "type is a desired filesystem type.",
											MarkdownDescription: "type is a desired filesystem type.",
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

							"mounts": schema.ListNestedAttribute{
								Description:         "mounts is a list of mount configuration.",
								MarkdownDescription: "mounts is a list of mount configuration.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"device": schema.StringAttribute{
											Description:         "device is path to a device that should be mounted.",
											MarkdownDescription: "device is path to a device that should be mounted.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"fs_type": schema.StringAttribute{
											Description:         "fsType specifies the filesystem on the device.",
											MarkdownDescription: "fsType specifies the filesystem on the device.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"mount_point": schema.StringAttribute{
											Description:         "mountPoint is a path where the device should be mounted at.",
											MarkdownDescription: "mountPoint is a path where the device should be mounted at.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"unsupported_options": schema.ListAttribute{
											Description:         "unsupportedOptions is a list of mount options used during device mounting. unsupported in this field name means that we won't support all the available options passed down using this field.",
											MarkdownDescription: "unsupportedOptions is a list of mount options used during device mounting. unsupported in this field name means that we won't support all the available options passed down using this field.",
											ElementType:         types.StringType,
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

							"raids": schema.ListNestedAttribute{
								Description:         "raids is a list of raid configurations.",
								MarkdownDescription: "raids is a list of raid configurations.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"raid0": schema.SingleNestedAttribute{
											Description:         "RAID0 specifies RAID0 options.",
											MarkdownDescription: "RAID0 specifies RAID0 options.",
											Attributes: map[string]schema.Attribute{
												"devices": schema.SingleNestedAttribute{
													Description:         "devices defines which devices constitute the raid array.",
													MarkdownDescription: "devices defines which devices constitute the raid array.",
													Attributes: map[string]schema.Attribute{
														"model_regex": schema.StringAttribute{
															Description:         "modelRegex is a regular expression filtering devices by their model name.",
															MarkdownDescription: "modelRegex is a regular expression filtering devices by their model name.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"name_regex": schema.StringAttribute{
															Description:         "nameRegex is a regular expression filtering devices by their name.",
															MarkdownDescription: "nameRegex is a regular expression filtering devices by their name.",
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

										"name": schema.StringAttribute{
											Description:         "name specifies the name of the raid device to be created under in '/dev/md/'.",
											MarkdownDescription: "name specifies the name of the raid device to be created under in '/dev/md/'.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"type": schema.StringAttribute{
											Description:         "type is a type of raid array.",
											MarkdownDescription: "type is a type of raid array.",
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
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"placement": schema.SingleNestedAttribute{
						Description:         "placement contains scheduling rules for NodeConfig Pods.",
						MarkdownDescription: "placement contains scheduling rules for NodeConfig Pods.",
						Attributes: map[string]schema.Attribute{
							"affinity": schema.SingleNestedAttribute{
								Description:         "affinity is a group of affinity scheduling rules for NodeConfig Pods.",
								MarkdownDescription: "affinity is a group of affinity scheduling rules for NodeConfig Pods.",
								Attributes: map[string]schema.Attribute{
									"node_affinity": schema.SingleNestedAttribute{
										Description:         "Describes node affinity scheduling rules for the pod.",
										MarkdownDescription: "Describes node affinity scheduling rules for the pod.",
										Attributes: map[string]schema.Attribute{
											"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
												Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.",
												MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"preference": schema.SingleNestedAttribute{
															Description:         "A node selector term, associated with the corresponding weight.",
															MarkdownDescription: "A node selector term, associated with the corresponding weight.",
															Attributes: map[string]schema.Attribute{
																"match_expressions": schema.ListNestedAttribute{
																	Description:         "A list of node selector requirements by node's labels.",
																	MarkdownDescription: "A list of node selector requirements by node's labels.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The label key that the selector applies to.",
																				MarkdownDescription: "The label key that the selector applies to.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"operator": schema.StringAttribute{
																				Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																				MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"values": schema.ListAttribute{
																				Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																				ElementType:         types.StringType,
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

																"match_fields": schema.ListNestedAttribute{
																	Description:         "A list of node selector requirements by node's fields.",
																	MarkdownDescription: "A list of node selector requirements by node's fields.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The label key that the selector applies to.",
																				MarkdownDescription: "The label key that the selector applies to.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"operator": schema.StringAttribute{
																				Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																				MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"values": schema.ListAttribute{
																				Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																				ElementType:         types.StringType,
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
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"weight": schema.Int64Attribute{
															Description:         "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
															MarkdownDescription: "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
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

											"required_during_scheduling_ignored_during_execution": schema.SingleNestedAttribute{
												Description:         "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.",
												MarkdownDescription: "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.",
												Attributes: map[string]schema.Attribute{
													"node_selector_terms": schema.ListNestedAttribute{
														Description:         "Required. A list of node selector terms. The terms are ORed.",
														MarkdownDescription: "Required. A list of node selector terms. The terms are ORed.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"match_expressions": schema.ListNestedAttribute{
																	Description:         "A list of node selector requirements by node's labels.",
																	MarkdownDescription: "A list of node selector requirements by node's labels.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The label key that the selector applies to.",
																				MarkdownDescription: "The label key that the selector applies to.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"operator": schema.StringAttribute{
																				Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																				MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"values": schema.ListAttribute{
																				Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																				ElementType:         types.StringType,
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

																"match_fields": schema.ListNestedAttribute{
																	Description:         "A list of node selector requirements by node's fields.",
																	MarkdownDescription: "A list of node selector requirements by node's fields.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The label key that the selector applies to.",
																				MarkdownDescription: "The label key that the selector applies to.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"operator": schema.StringAttribute{
																				Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																				MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"values": schema.ListAttribute{
																				Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																				ElementType:         types.StringType,
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
										Required: false,
										Optional: false,
										Computed: true,
									},

									"pod_affinity": schema.SingleNestedAttribute{
										Description:         "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",
										MarkdownDescription: "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",
										Attributes: map[string]schema.Attribute{
											"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
												Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
												MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"pod_affinity_term": schema.SingleNestedAttribute{
															Description:         "Required. A pod affinity term, associated with the corresponding weight.",
															MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",
															Attributes: map[string]schema.Attribute{
																"label_selector": schema.SingleNestedAttribute{
																	Description:         "A label query over a set of resources, in this case pods.",
																	MarkdownDescription: "A label query over a set of resources, in this case pods.",
																	Attributes: map[string]schema.Attribute{
																		"match_expressions": schema.ListNestedAttribute{
																			Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																			MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "key is the label key that the selector applies to.",
																						MarkdownDescription: "key is the label key that the selector applies to.",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"operator": schema.StringAttribute{
																						Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"values": schema.ListAttribute{
																						Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																						MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																						ElementType:         types.StringType,
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

																		"match_labels": schema.MapAttribute{
																			Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																			MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

																"namespace_selector": schema.SingleNestedAttribute{
																	Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																	MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																	Attributes: map[string]schema.Attribute{
																		"match_expressions": schema.ListNestedAttribute{
																			Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																			MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "key is the label key that the selector applies to.",
																						MarkdownDescription: "key is the label key that the selector applies to.",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"operator": schema.StringAttribute{
																						Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"values": schema.ListAttribute{
																						Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																						MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																						ElementType:         types.StringType,
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

																		"match_labels": schema.MapAttribute{
																			Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																			MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

																"namespaces": schema.ListAttribute{
																	Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																	MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"topology_key": schema.StringAttribute{
																	Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																	MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"weight": schema.Int64Attribute{
															Description:         "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
															MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
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

											"required_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
												Description:         "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
												MarkdownDescription: "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"label_selector": schema.SingleNestedAttribute{
															Description:         "A label query over a set of resources, in this case pods.",
															MarkdownDescription: "A label query over a set of resources, in this case pods.",
															Attributes: map[string]schema.Attribute{
																"match_expressions": schema.ListNestedAttribute{
																	Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "key is the label key that the selector applies to.",
																				MarkdownDescription: "key is the label key that the selector applies to.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"operator": schema.StringAttribute{
																				Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"values": schema.ListAttribute{
																				Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																				ElementType:         types.StringType,
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

																"match_labels": schema.MapAttribute{
																	Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

														"namespace_selector": schema.SingleNestedAttribute{
															Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
															MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
															Attributes: map[string]schema.Attribute{
																"match_expressions": schema.ListNestedAttribute{
																	Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "key is the label key that the selector applies to.",
																				MarkdownDescription: "key is the label key that the selector applies to.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"operator": schema.StringAttribute{
																				Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"values": schema.ListAttribute{
																				Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																				ElementType:         types.StringType,
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

																"match_labels": schema.MapAttribute{
																	Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

														"namespaces": schema.ListAttribute{
															Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
															MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"topology_key": schema.StringAttribute{
															Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
															MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
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
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"pod_anti_affinity": schema.SingleNestedAttribute{
										Description:         "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",
										MarkdownDescription: "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",
										Attributes: map[string]schema.Attribute{
											"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
												Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
												MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"pod_affinity_term": schema.SingleNestedAttribute{
															Description:         "Required. A pod affinity term, associated with the corresponding weight.",
															MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",
															Attributes: map[string]schema.Attribute{
																"label_selector": schema.SingleNestedAttribute{
																	Description:         "A label query over a set of resources, in this case pods.",
																	MarkdownDescription: "A label query over a set of resources, in this case pods.",
																	Attributes: map[string]schema.Attribute{
																		"match_expressions": schema.ListNestedAttribute{
																			Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																			MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "key is the label key that the selector applies to.",
																						MarkdownDescription: "key is the label key that the selector applies to.",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"operator": schema.StringAttribute{
																						Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"values": schema.ListAttribute{
																						Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																						MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																						ElementType:         types.StringType,
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

																		"match_labels": schema.MapAttribute{
																			Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																			MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

																"namespace_selector": schema.SingleNestedAttribute{
																	Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																	MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																	Attributes: map[string]schema.Attribute{
																		"match_expressions": schema.ListNestedAttribute{
																			Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																			MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "key is the label key that the selector applies to.",
																						MarkdownDescription: "key is the label key that the selector applies to.",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"operator": schema.StringAttribute{
																						Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"values": schema.ListAttribute{
																						Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																						MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																						ElementType:         types.StringType,
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

																		"match_labels": schema.MapAttribute{
																			Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																			MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

																"namespaces": schema.ListAttribute{
																	Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																	MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"topology_key": schema.StringAttribute{
																	Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																	MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"weight": schema.Int64Attribute{
															Description:         "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
															MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
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

											"required_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
												Description:         "If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
												MarkdownDescription: "If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"label_selector": schema.SingleNestedAttribute{
															Description:         "A label query over a set of resources, in this case pods.",
															MarkdownDescription: "A label query over a set of resources, in this case pods.",
															Attributes: map[string]schema.Attribute{
																"match_expressions": schema.ListNestedAttribute{
																	Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "key is the label key that the selector applies to.",
																				MarkdownDescription: "key is the label key that the selector applies to.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"operator": schema.StringAttribute{
																				Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"values": schema.ListAttribute{
																				Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																				ElementType:         types.StringType,
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

																"match_labels": schema.MapAttribute{
																	Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

														"namespace_selector": schema.SingleNestedAttribute{
															Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
															MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
															Attributes: map[string]schema.Attribute{
																"match_expressions": schema.ListNestedAttribute{
																	Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "key is the label key that the selector applies to.",
																				MarkdownDescription: "key is the label key that the selector applies to.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"operator": schema.StringAttribute{
																				Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"values": schema.ListAttribute{
																				Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																				ElementType:         types.StringType,
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

																"match_labels": schema.MapAttribute{
																	Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

														"namespaces": schema.ListAttribute{
															Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
															MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"topology_key": schema.StringAttribute{
															Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
															MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
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

							"node_selector": schema.MapAttribute{
								Description:         "nodeSelector is a selector which must be true for the NodeConfig Pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node.",
								MarkdownDescription: "nodeSelector is a selector which must be true for the NodeConfig Pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"tolerations": schema.ListNestedAttribute{
								Description:         "tolerations is a group of tolerations NodeConfig Pods are going to have.",
								MarkdownDescription: "tolerations is a group of tolerations NodeConfig Pods are going to have.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"effect": schema.StringAttribute{
											Description:         "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
											MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"key": schema.StringAttribute{
											Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
											MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"operator": schema.StringAttribute{
											Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
											MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"toleration_seconds": schema.Int64Attribute{
											Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
											MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"value": schema.StringAttribute{
											Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
											MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
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

func (r *ScyllaScylladbComNodeConfigV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *ScyllaScylladbComNodeConfigV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_scylla_scylladb_com_node_config_v1alpha1")

	var data ScyllaScylladbComNodeConfigV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "scylla.scylladb.com", Version: "v1alpha1", Resource: "NodeConfig"}).
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

	var readResponse ScyllaScylladbComNodeConfigV1Alpha1DataSourceData
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

	data.ID = types.StringValue(data.Metadata.Name)
	data.ApiVersion = pointer.String("scylla.scylladb.com/v1alpha1")
	data.Kind = pointer.String("NodeConfig")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}