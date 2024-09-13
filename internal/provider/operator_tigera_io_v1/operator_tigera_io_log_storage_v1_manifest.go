/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package operator_tigera_io_v1

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
	_ datasource.DataSource = &OperatorTigeraIoLogStorageV1Manifest{}
)

func NewOperatorTigeraIoLogStorageV1Manifest() datasource.DataSource {
	return &OperatorTigeraIoLogStorageV1Manifest{}
}

type OperatorTigeraIoLogStorageV1Manifest struct{}

type OperatorTigeraIoLogStorageV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		ComponentResources *[]struct {
			ComponentName        *string `tfsdk:"component_name" json:"componentName,omitempty"`
			ResourceRequirements *struct {
				Claims *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resource_requirements" json:"resourceRequirements,omitempty"`
		} `tfsdk:"component_resources" json:"componentResources,omitempty"`
		DataNodeSelector       *map[string]string `tfsdk:"data_node_selector" json:"dataNodeSelector,omitempty"`
		EckOperatorStatefulSet *struct {
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
		} `tfsdk:"eck_operator_stateful_set" json:"eckOperatorStatefulSet,omitempty"`
		ElasticsearchMetricsDeployment *struct {
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
		} `tfsdk:"elasticsearch_metrics_deployment" json:"elasticsearchMetricsDeployment,omitempty"`
		Indices *struct {
			Replicas *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
		} `tfsdk:"indices" json:"indices,omitempty"`
		Kibana *struct {
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
		} `tfsdk:"kibana" json:"kibana,omitempty"`
		LinseedDeployment *struct {
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
		} `tfsdk:"linseed_deployment" json:"linseedDeployment,omitempty"`
		Nodes *struct {
			Count    *int64 `tfsdk:"count" json:"count,omitempty"`
			NodeSets *[]struct {
				SelectionAttributes *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					NodeLabel *string `tfsdk:"node_label" json:"nodeLabel,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"selection_attributes" json:"selectionAttributes,omitempty"`
			} `tfsdk:"node_sets" json:"nodeSets,omitempty"`
			ResourceRequirements *struct {
				Claims *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resource_requirements" json:"resourceRequirements,omitempty"`
		} `tfsdk:"nodes" json:"nodes,omitempty"`
		Retention *struct {
			AuditReports      *int64 `tfsdk:"audit_reports" json:"auditReports,omitempty"`
			BgpLogs           *int64 `tfsdk:"bgp_logs" json:"bgpLogs,omitempty"`
			ComplianceReports *int64 `tfsdk:"compliance_reports" json:"complianceReports,omitempty"`
			DnsLogs           *int64 `tfsdk:"dns_logs" json:"dnsLogs,omitempty"`
			Flows             *int64 `tfsdk:"flows" json:"flows,omitempty"`
			Snapshots         *int64 `tfsdk:"snapshots" json:"snapshots,omitempty"`
		} `tfsdk:"retention" json:"retention,omitempty"`
		StorageClassName *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OperatorTigeraIoLogStorageV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_operator_tigera_io_log_storage_v1_manifest"
}

func (r *OperatorTigeraIoLogStorageV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "LogStorage installs the components required for Tigera flow and DNS log storage. At most one instance of this resource is supported. It must be named 'tigera-secure'. When created, this installs an Elasticsearch cluster for use by Calico Enterprise.",
		MarkdownDescription: "LogStorage installs the components required for Tigera flow and DNS log storage. At most one instance of this resource is supported. It must be named 'tigera-secure'. When created, this installs an Elasticsearch cluster for use by Calico Enterprise.",
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
				Description:         "Specification of the desired state for Tigera log storage.",
				MarkdownDescription: "Specification of the desired state for Tigera log storage.",
				Attributes: map[string]schema.Attribute{
					"component_resources": schema.ListNestedAttribute{
						Description:         "ComponentResources can be used to customize the resource requirements for each component. Only ECKOperator is supported for this spec.",
						MarkdownDescription: "ComponentResources can be used to customize the resource requirements for each component. Only ECKOperator is supported for this spec.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"component_name": schema.StringAttribute{
									Description:         "Deprecated. Please use ECKOperatorStatefulSet. ComponentName is an enum which identifies the component",
									MarkdownDescription: "Deprecated. Please use ECKOperatorStatefulSet. ComponentName is an enum which identifies the component",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("ECKOperator"),
									},
								},

								"resource_requirements": schema.SingleNestedAttribute{
									Description:         "ResourceRequirements allows customization of limits and requests for compute resources such as cpu and memory.",
									MarkdownDescription: "ResourceRequirements allows customization of limits and requests for compute resources such as cpu and memory.",
									Attributes: map[string]schema.Attribute{
										"claims": schema.ListNestedAttribute{
											Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
											MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
														MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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
											Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
											MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"requests": schema.MapAttribute{
											Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
											MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

					"data_node_selector": schema.MapAttribute{
						Description:         "DataNodeSelector gives you more control over the node that Elasticsearch will run on. The contents of DataNodeSelector will be added to the PodSpec of the Elasticsearch nodes. For the pod to be eligible to run on a node, the node must have each of the indicated key-value pairs as labels as well as access to the specified StorageClassName.",
						MarkdownDescription: "DataNodeSelector gives you more control over the node that Elasticsearch will run on. The contents of DataNodeSelector will be added to the PodSpec of the Elasticsearch nodes. For the pod to be eligible to run on a node, the node must have each of the indicated key-value pairs as labels as well as access to the specified StorageClassName.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"eck_operator_stateful_set": schema.SingleNestedAttribute{
						Description:         "ECKOperatorStatefulSet configures the ECKOperator StatefulSet. If used in conjunction with the deprecated ComponentResources, then these overrides take precedence.",
						MarkdownDescription: "ECKOperatorStatefulSet configures the ECKOperator StatefulSet. If used in conjunction with the deprecated ComponentResources, then these overrides take precedence.",
						Attributes: map[string]schema.Attribute{
							"spec": schema.SingleNestedAttribute{
								Description:         "Spec is the specification of the ECKOperator StatefulSet.",
								MarkdownDescription: "Spec is the specification of the ECKOperator StatefulSet.",
								Attributes: map[string]schema.Attribute{
									"template": schema.SingleNestedAttribute{
										Description:         "Template describes the ECKOperator StatefulSet pod that will be created.",
										MarkdownDescription: "Template describes the ECKOperator StatefulSet pod that will be created.",
										Attributes: map[string]schema.Attribute{
											"spec": schema.SingleNestedAttribute{
												Description:         "Spec is the ECKOperator StatefulSet's PodSpec.",
												MarkdownDescription: "Spec is the ECKOperator StatefulSet's PodSpec.",
												Attributes: map[string]schema.Attribute{
													"containers": schema.ListNestedAttribute{
														Description:         "Containers is a list of ECKOperator StatefulSet containers. If specified, this overrides the specified ECKOperator StatefulSet containers. If omitted, the ECKOperator StatefulSet will use its default values for its containers.",
														MarkdownDescription: "Containers is a list of ECKOperator StatefulSet containers. If specified, this overrides the specified ECKOperator StatefulSet containers. If omitted, the ECKOperator StatefulSet will use its default values for its containers.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is an enum which identifies the ECKOperator StatefulSet container by name. Supported values are: manager",
																	MarkdownDescription: "Name is an enum which identifies the ECKOperator StatefulSet container by name. Supported values are: manager",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("manager"),
																	},
																},

																"resources": schema.SingleNestedAttribute{
																	Description:         "Resources allows customization of limits and requests for compute resources such as cpu and memory. If specified, this overrides the named ECKOperator StatefulSet container's resources. If omitted, the ECKOperator StatefulSet will use its default value for this container's resources.",
																	MarkdownDescription: "Resources allows customization of limits and requests for compute resources such as cpu and memory. If specified, this overrides the named ECKOperator StatefulSet container's resources. If omitted, the ECKOperator StatefulSet will use its default value for this container's resources.",
																	Attributes: map[string]schema.Attribute{
																		"claims": schema.ListNestedAttribute{
																			Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
																			MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
																						MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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
																			Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"requests": schema.MapAttribute{
																			Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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
														Description:         "InitContainers is a list of ECKOperator StatefulSet init containers. If specified, this overrides the specified ECKOperator StatefulSet init containers. If omitted, the ECKOperator StatefulSet will use its default values for its init containers.",
														MarkdownDescription: "InitContainers is a list of ECKOperator StatefulSet init containers. If specified, this overrides the specified ECKOperator StatefulSet init containers. If omitted, the ECKOperator StatefulSet will use its default values for its init containers.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is an enum which identifies the ECKOperator StatefulSet init container by name.",
																	MarkdownDescription: "Name is an enum which identifies the ECKOperator StatefulSet init container by name.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"resources": schema.SingleNestedAttribute{
																	Description:         "Resources allows customization of limits and requests for compute resources such as cpu and memory. If specified, this overrides the named ECKOperator StatefulSet init container's resources. If omitted, the ECKOperator StatefulSet will use its default value for this init container's resources.",
																	MarkdownDescription: "Resources allows customization of limits and requests for compute resources such as cpu and memory. If specified, this overrides the named ECKOperator StatefulSet init container's resources. If omitted, the ECKOperator StatefulSet will use its default value for this init container's resources.",
																	Attributes: map[string]schema.Attribute{
																		"claims": schema.ListNestedAttribute{
																			Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
																			MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
																						MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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
																			Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"requests": schema.MapAttribute{
																			Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

					"elasticsearch_metrics_deployment": schema.SingleNestedAttribute{
						Description:         "ElasticsearchMetricsDeployment configures the tigera-elasticsearch-metric Deployment.",
						MarkdownDescription: "ElasticsearchMetricsDeployment configures the tigera-elasticsearch-metric Deployment.",
						Attributes: map[string]schema.Attribute{
							"spec": schema.SingleNestedAttribute{
								Description:         "Spec is the specification of the ElasticsearchMetrics Deployment.",
								MarkdownDescription: "Spec is the specification of the ElasticsearchMetrics Deployment.",
								Attributes: map[string]schema.Attribute{
									"template": schema.SingleNestedAttribute{
										Description:         "Template describes the ElasticsearchMetrics Deployment pod that will be created.",
										MarkdownDescription: "Template describes the ElasticsearchMetrics Deployment pod that will be created.",
										Attributes: map[string]schema.Attribute{
											"spec": schema.SingleNestedAttribute{
												Description:         "Spec is the ElasticsearchMetrics Deployment's PodSpec.",
												MarkdownDescription: "Spec is the ElasticsearchMetrics Deployment's PodSpec.",
												Attributes: map[string]schema.Attribute{
													"containers": schema.ListNestedAttribute{
														Description:         "Containers is a list of ElasticsearchMetricsDeployment containers. If specified, this overrides the specified ElasticsearchMetricsDeployment containers. If omitted, the ElasticsearchMetrics Deployment will use its default values for its containers.",
														MarkdownDescription: "Containers is a list of ElasticsearchMetricsDeployment containers. If specified, this overrides the specified ElasticsearchMetricsDeployment containers. If omitted, the ElasticsearchMetrics Deployment will use its default values for its containers.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is an enum which identifies the ElasticsearchMetricsDeployment container by name. Supported values are: tigera-elasticsearch-metrics",
																	MarkdownDescription: "Name is an enum which identifies the ElasticsearchMetricsDeployment container by name. Supported values are: tigera-elasticsearch-metrics",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("tigera-elasticsearch-metrics"),
																	},
																},

																"resources": schema.SingleNestedAttribute{
																	Description:         "Resources allows customization of limits and requests for compute resources such as cpu and memory. If specified, this overrides the named ElasticsearchMetricsDeployment container's resources. If omitted, the ElasticsearchMetrics Deployment will use its default value for this container's resources.",
																	MarkdownDescription: "Resources allows customization of limits and requests for compute resources such as cpu and memory. If specified, this overrides the named ElasticsearchMetricsDeployment container's resources. If omitted, the ElasticsearchMetrics Deployment will use its default value for this container's resources.",
																	Attributes: map[string]schema.Attribute{
																		"claims": schema.ListNestedAttribute{
																			Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
																			MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
																						MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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
																			Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"requests": schema.MapAttribute{
																			Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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
														Description:         "InitContainers is a list of ElasticsearchMetricsDeployment init containers. If specified, this overrides the specified ElasticsearchMetricsDeployment init containers. If omitted, the ElasticsearchMetrics Deployment will use its default values for its init containers.",
														MarkdownDescription: "InitContainers is a list of ElasticsearchMetricsDeployment init containers. If specified, this overrides the specified ElasticsearchMetricsDeployment init containers. If omitted, the ElasticsearchMetrics Deployment will use its default values for its init containers.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is an enum which identifies the ElasticsearchMetricsDeployment init container by name. Supported values are: tigera-ee-elasticsearch-metrics-tls-key-cert-provisioner",
																	MarkdownDescription: "Name is an enum which identifies the ElasticsearchMetricsDeployment init container by name. Supported values are: tigera-ee-elasticsearch-metrics-tls-key-cert-provisioner",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("tigera-ee-elasticsearch-metrics-tls-key-cert-provisioner"),
																	},
																},

																"resources": schema.SingleNestedAttribute{
																	Description:         "Resources allows customization of limits and requests for compute resources such as cpu and memory. If specified, this overrides the named ElasticsearchMetricsDeployment init container's resources. If omitted, the ElasticsearchMetrics Deployment will use its default value for this init container's resources.",
																	MarkdownDescription: "Resources allows customization of limits and requests for compute resources such as cpu and memory. If specified, this overrides the named ElasticsearchMetricsDeployment init container's resources. If omitted, the ElasticsearchMetrics Deployment will use its default value for this init container's resources.",
																	Attributes: map[string]schema.Attribute{
																		"claims": schema.ListNestedAttribute{
																			Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
																			MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
																						MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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
																			Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"requests": schema.MapAttribute{
																			Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

					"indices": schema.SingleNestedAttribute{
						Description:         "Index defines the configuration for the indices in the Elasticsearch cluster.",
						MarkdownDescription: "Index defines the configuration for the indices in the Elasticsearch cluster.",
						Attributes: map[string]schema.Attribute{
							"replicas": schema.Int64Attribute{
								Description:         "Replicas defines how many replicas each index will have. See https://www.elastic.co/guide/en/elasticsearch/reference/current/scalability.html",
								MarkdownDescription: "Replicas defines how many replicas each index will have. See https://www.elastic.co/guide/en/elasticsearch/reference/current/scalability.html",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"kibana": schema.SingleNestedAttribute{
						Description:         "Kibana configures the Kibana Spec.",
						MarkdownDescription: "Kibana configures the Kibana Spec.",
						Attributes: map[string]schema.Attribute{
							"spec": schema.SingleNestedAttribute{
								Description:         "Spec is the specification of the Kibana.",
								MarkdownDescription: "Spec is the specification of the Kibana.",
								Attributes: map[string]schema.Attribute{
									"template": schema.SingleNestedAttribute{
										Description:         "Template describes the Kibana pod that will be created.",
										MarkdownDescription: "Template describes the Kibana pod that will be created.",
										Attributes: map[string]schema.Attribute{
											"spec": schema.SingleNestedAttribute{
												Description:         "Spec is the Kibana's PodSpec.",
												MarkdownDescription: "Spec is the Kibana's PodSpec.",
												Attributes: map[string]schema.Attribute{
													"containers": schema.ListNestedAttribute{
														Description:         "Containers is a list of Kibana containers. If specified, this overrides the specified Kibana Deployment containers. If omitted, the Kibana Deployment will use its default values for its containers.",
														MarkdownDescription: "Containers is a list of Kibana containers. If specified, this overrides the specified Kibana Deployment containers. If omitted, the Kibana Deployment will use its default values for its containers.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is an enum which identifies the Kibana Deployment container by name. Supported values are: kibana",
																	MarkdownDescription: "Name is an enum which identifies the Kibana Deployment container by name. Supported values are: kibana",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("kibana"),
																	},
																},

																"resources": schema.SingleNestedAttribute{
																	Description:         "Resources allows customization of limits and requests for compute resources such as cpu and memory. If specified, this overrides the named Kibana container's resources. If omitted, the Kibana will use its default value for this container's resources.",
																	MarkdownDescription: "Resources allows customization of limits and requests for compute resources such as cpu and memory. If specified, this overrides the named Kibana container's resources. If omitted, the Kibana will use its default value for this container's resources.",
																	Attributes: map[string]schema.Attribute{
																		"claims": schema.ListNestedAttribute{
																			Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
																			MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
																						MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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
																			Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"requests": schema.MapAttribute{
																			Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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
														Description:         "InitContainers is a list of Kibana init containers. If specified, this overrides the specified Kibana Deployment init containers. If omitted, the Kibana Deployment will use its default values for its init containers.",
														MarkdownDescription: "InitContainers is a list of Kibana init containers. If specified, this overrides the specified Kibana Deployment init containers. If omitted, the Kibana Deployment will use its default values for its init containers.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is an enum which identifies the Kibana init container by name. Supported values are: key-cert-provisioner",
																	MarkdownDescription: "Name is an enum which identifies the Kibana init container by name. Supported values are: key-cert-provisioner",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("key-cert-provisioner"),
																	},
																},

																"resources": schema.SingleNestedAttribute{
																	Description:         "Resources allows customization of limits and requests for compute resources such as cpu and memory. If specified, this overrides the named Kibana Deployment init container's resources. If omitted, the Kibana Deployment will use its default value for this init container's resources. If used in conjunction with the deprecated ComponentResources, then this value takes precedence.",
																	MarkdownDescription: "Resources allows customization of limits and requests for compute resources such as cpu and memory. If specified, this overrides the named Kibana Deployment init container's resources. If omitted, the Kibana Deployment will use its default value for this init container's resources. If used in conjunction with the deprecated ComponentResources, then this value takes precedence.",
																	Attributes: map[string]schema.Attribute{
																		"claims": schema.ListNestedAttribute{
																			Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
																			MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
																						MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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
																			Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"requests": schema.MapAttribute{
																			Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

					"linseed_deployment": schema.SingleNestedAttribute{
						Description:         "LinseedDeployment configures the linseed Deployment.",
						MarkdownDescription: "LinseedDeployment configures the linseed Deployment.",
						Attributes: map[string]schema.Attribute{
							"spec": schema.SingleNestedAttribute{
								Description:         "Spec is the specification of the linseed Deployment.",
								MarkdownDescription: "Spec is the specification of the linseed Deployment.",
								Attributes: map[string]schema.Attribute{
									"template": schema.SingleNestedAttribute{
										Description:         "Template describes the linseed Deployment pod that will be created.",
										MarkdownDescription: "Template describes the linseed Deployment pod that will be created.",
										Attributes: map[string]schema.Attribute{
											"spec": schema.SingleNestedAttribute{
												Description:         "Spec is the linseed Deployment's PodSpec.",
												MarkdownDescription: "Spec is the linseed Deployment's PodSpec.",
												Attributes: map[string]schema.Attribute{
													"containers": schema.ListNestedAttribute{
														Description:         "Containers is a list of linseed containers. If specified, this overrides the specified linseed Deployment containers. If omitted, the linseed Deployment will use its default values for its containers.",
														MarkdownDescription: "Containers is a list of linseed containers. If specified, this overrides the specified linseed Deployment containers. If omitted, the linseed Deployment will use its default values for its containers.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is an enum which identifies the linseed Deployment container by name. Supported values are: tigera-linseed",
																	MarkdownDescription: "Name is an enum which identifies the linseed Deployment container by name. Supported values are: tigera-linseed",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("tigera-linseed"),
																	},
																},

																"resources": schema.SingleNestedAttribute{
																	Description:         "Resources allows customization of limits and requests for compute resources such as cpu and memory. If specified, this overrides the named linseed Deployment container's resources. If omitted, the linseed Deployment will use its default value for this container's resources.",
																	MarkdownDescription: "Resources allows customization of limits and requests for compute resources such as cpu and memory. If specified, this overrides the named linseed Deployment container's resources. If omitted, the linseed Deployment will use its default value for this container's resources.",
																	Attributes: map[string]schema.Attribute{
																		"claims": schema.ListNestedAttribute{
																			Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
																			MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
																						MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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
																			Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"requests": schema.MapAttribute{
																			Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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
														Description:         "InitContainers is a list of linseed init containers. If specified, this overrides the specified linseed Deployment init containers. If omitted, the linseed Deployment will use its default values for its init containers.",
														MarkdownDescription: "InitContainers is a list of linseed init containers. If specified, this overrides the specified linseed Deployment init containers. If omitted, the linseed Deployment will use its default values for its init containers.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is an enum which identifies the linseed Deployment init container by name. Supported values are: tigera-secure-linseed-token-tls-key-cert-provisioner,tigera-secure-linseed-cert-key-cert-provisioner",
																	MarkdownDescription: "Name is an enum which identifies the linseed Deployment init container by name. Supported values are: tigera-secure-linseed-token-tls-key-cert-provisioner,tigera-secure-linseed-cert-key-cert-provisioner",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("tigera-secure-linseed-token-tls-key-cert-provisioner", "tigera-secure-linseed-cert-key-cert-provisioner"),
																	},
																},

																"resources": schema.SingleNestedAttribute{
																	Description:         "Resources allows customization of limits and requests for compute resources such as cpu and memory. If specified, this overrides the named linseed Deployment init container's resources. If omitted, the linseed Deployment will use its default value for this init container's resources.",
																	MarkdownDescription: "Resources allows customization of limits and requests for compute resources such as cpu and memory. If specified, this overrides the named linseed Deployment init container's resources. If omitted, the linseed Deployment will use its default value for this init container's resources.",
																	Attributes: map[string]schema.Attribute{
																		"claims": schema.ListNestedAttribute{
																			Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
																			MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
																						MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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
																			Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"requests": schema.MapAttribute{
																			Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

					"nodes": schema.SingleNestedAttribute{
						Description:         "Nodes defines the configuration for a set of identical Elasticsearch cluster nodes, each of type master, data, and ingest.",
						MarkdownDescription: "Nodes defines the configuration for a set of identical Elasticsearch cluster nodes, each of type master, data, and ingest.",
						Attributes: map[string]schema.Attribute{
							"count": schema.Int64Attribute{
								Description:         "Count defines the number of nodes in the Elasticsearch cluster.",
								MarkdownDescription: "Count defines the number of nodes in the Elasticsearch cluster.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"node_sets": schema.ListNestedAttribute{
								Description:         "NodeSets defines configuration specific to each Elasticsearch Node Set",
								MarkdownDescription: "NodeSets defines configuration specific to each Elasticsearch Node Set",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"selection_attributes": schema.ListNestedAttribute{
											Description:         "SelectionAttributes defines K8s node attributes a NodeSet should use when setting the Node Affinity selectors and Elasticsearch cluster awareness attributes for the Elasticsearch nodes. The list of SelectionAttributes are used to define Node Affinities and set the node awareness configuration in the running Elasticsearch instance.",
											MarkdownDescription: "SelectionAttributes defines K8s node attributes a NodeSet should use when setting the Node Affinity selectors and Elasticsearch cluster awareness attributes for the Elasticsearch nodes. The list of SelectionAttributes are used to define Node Affinities and set the node awareness configuration in the running Elasticsearch instance.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"node_label": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"resource_requirements": schema.SingleNestedAttribute{
								Description:         "ResourceRequirements defines the resource limits and requirements for the Elasticsearch cluster.",
								MarkdownDescription: "ResourceRequirements defines the resource limits and requirements for the Elasticsearch cluster.",
								Attributes: map[string]schema.Attribute{
									"claims": schema.ListNestedAttribute{
										Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
										MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
													MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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
										Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"requests": schema.MapAttribute{
										Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"retention": schema.SingleNestedAttribute{
						Description:         "Retention defines how long data is retained in the Elasticsearch cluster before it is cleared.",
						MarkdownDescription: "Retention defines how long data is retained in the Elasticsearch cluster before it is cleared.",
						Attributes: map[string]schema.Attribute{
							"audit_reports": schema.Int64Attribute{
								Description:         "AuditReports configures the retention period for audit logs, in days. Logs written on a day that started at least this long ago are removed. To keep logs for at least x days, use a retention period of x+1. Default: 91",
								MarkdownDescription: "AuditReports configures the retention period for audit logs, in days. Logs written on a day that started at least this long ago are removed. To keep logs for at least x days, use a retention period of x+1. Default: 91",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"bgp_logs": schema.Int64Attribute{
								Description:         "BGPLogs configures the retention period for BGP logs, in days. Logs written on a day that started at least this long ago are removed. To keep logs for at least x days, use a retention period of x+1. Default: 8",
								MarkdownDescription: "BGPLogs configures the retention period for BGP logs, in days. Logs written on a day that started at least this long ago are removed. To keep logs for at least x days, use a retention period of x+1. Default: 8",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"compliance_reports": schema.Int64Attribute{
								Description:         "ComplianceReports configures the retention period for compliance reports, in days. Reports are output from the analysis of the system state and audit events for compliance reporting. Consult the Compliance Reporting documentation for more details on reports. Logs written on a day that started at least this long ago are removed. To keep logs for at least x days, use a retention period of x+1. Default: 91",
								MarkdownDescription: "ComplianceReports configures the retention period for compliance reports, in days. Reports are output from the analysis of the system state and audit events for compliance reporting. Consult the Compliance Reporting documentation for more details on reports. Logs written on a day that started at least this long ago are removed. To keep logs for at least x days, use a retention period of x+1. Default: 91",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"dns_logs": schema.Int64Attribute{
								Description:         "DNSLogs configures the retention period for DNS logs, in days. Logs written on a day that started at least this long ago are removed. To keep logs for at least x days, use a retention period of x+1. Default: 8",
								MarkdownDescription: "DNSLogs configures the retention period for DNS logs, in days. Logs written on a day that started at least this long ago are removed. To keep logs for at least x days, use a retention period of x+1. Default: 8",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"flows": schema.Int64Attribute{
								Description:         "Flows configures the retention period for flow logs, in days. Logs written on a day that started at least this long ago are removed. To keep logs for at least x days, use a retention period of x+1. Default: 8",
								MarkdownDescription: "Flows configures the retention period for flow logs, in days. Logs written on a day that started at least this long ago are removed. To keep logs for at least x days, use a retention period of x+1. Default: 8",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"snapshots": schema.Int64Attribute{
								Description:         "Snapshots configures the retention period for snapshots, in days. Snapshots are periodic captures of resources which along with audit events are used to generate reports. Consult the Compliance Reporting documentation for more details on snapshots. Logs written on a day that started at least this long ago are removed. To keep logs for at least x days, use a retention period of x+1. Default: 91",
								MarkdownDescription: "Snapshots configures the retention period for snapshots, in days. Snapshots are periodic captures of resources which along with audit events are used to generate reports. Consult the Compliance Reporting documentation for more details on snapshots. Logs written on a day that started at least this long ago are removed. To keep logs for at least x days, use a retention period of x+1. Default: 91",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"storage_class_name": schema.StringAttribute{
						Description:         "StorageClassName will populate the PersistentVolumeClaim.StorageClassName that is used to provision disks to the Tigera Elasticsearch cluster. The StorageClassName should only be modified when no LogStorage is currently active. We recommend choosing a storage class dedicated to Tigera LogStorage only. Otherwise, data retention cannot be guaranteed during upgrades. See https://docs.tigera.io/maintenance/upgrading for up-to-date instructions. Default: tigera-elasticsearch",
						MarkdownDescription: "StorageClassName will populate the PersistentVolumeClaim.StorageClassName that is used to provision disks to the Tigera Elasticsearch cluster. The StorageClassName should only be modified when no LogStorage is currently active. We recommend choosing a storage class dedicated to Tigera LogStorage only. Otherwise, data retention cannot be guaranteed during upgrades. See https://docs.tigera.io/maintenance/upgrading for up-to-date instructions. Default: tigera-elasticsearch",
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

func (r *OperatorTigeraIoLogStorageV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_operator_tigera_io_log_storage_v1_manifest")

	var model OperatorTigeraIoLogStorageV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("operator.tigera.io/v1")
	model.Kind = pointer.String("LogStorage")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
