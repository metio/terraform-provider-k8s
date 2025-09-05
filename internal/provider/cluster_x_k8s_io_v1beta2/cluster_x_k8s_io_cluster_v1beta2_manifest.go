/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package cluster_x_k8s_io_v1beta2

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/boolvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
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
	_ datasource.DataSource = &ClusterXK8SIoClusterV1Beta2Manifest{}
)

func NewClusterXK8SIoClusterV1Beta2Manifest() datasource.DataSource {
	return &ClusterXK8SIoClusterV1Beta2Manifest{}
}

type ClusterXK8SIoClusterV1Beta2Manifest struct{}

type ClusterXK8SIoClusterV1Beta2ManifestData struct {
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
		AvailabilityGates *[]struct {
			ConditionType *string `tfsdk:"condition_type" json:"conditionType,omitempty"`
			Polarity      *string `tfsdk:"polarity" json:"polarity,omitempty"`
		} `tfsdk:"availability_gates" json:"availabilityGates,omitempty"`
		ClusterNetwork *struct {
			ApiServerPort *int64 `tfsdk:"api_server_port" json:"apiServerPort,omitempty"`
			Pods          *struct {
				CidrBlocks *[]string `tfsdk:"cidr_blocks" json:"cidrBlocks,omitempty"`
			} `tfsdk:"pods" json:"pods,omitempty"`
			ServiceDomain *string `tfsdk:"service_domain" json:"serviceDomain,omitempty"`
			Services      *struct {
				CidrBlocks *[]string `tfsdk:"cidr_blocks" json:"cidrBlocks,omitempty"`
			} `tfsdk:"services" json:"services,omitempty"`
		} `tfsdk:"cluster_network" json:"clusterNetwork,omitempty"`
		ControlPlaneEndpoint *struct {
			Host *string `tfsdk:"host" json:"host,omitempty"`
			Port *int64  `tfsdk:"port" json:"port,omitempty"`
		} `tfsdk:"control_plane_endpoint" json:"controlPlaneEndpoint,omitempty"`
		ControlPlaneRef *struct {
			ApiGroup *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
			Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
			Name     *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"control_plane_ref" json:"controlPlaneRef,omitempty"`
		InfrastructureRef *struct {
			ApiGroup *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
			Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
			Name     *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"infrastructure_ref" json:"infrastructureRef,omitempty"`
		Paused   *bool `tfsdk:"paused" json:"paused,omitempty"`
		Topology *struct {
			ClassRef *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"class_ref" json:"classRef,omitempty"`
			ControlPlane *struct {
				Deletion *struct {
					NodeDeletionTimeoutSeconds     *int64 `tfsdk:"node_deletion_timeout_seconds" json:"nodeDeletionTimeoutSeconds,omitempty"`
					NodeDrainTimeoutSeconds        *int64 `tfsdk:"node_drain_timeout_seconds" json:"nodeDrainTimeoutSeconds,omitempty"`
					NodeVolumeDetachTimeoutSeconds *int64 `tfsdk:"node_volume_detach_timeout_seconds" json:"nodeVolumeDetachTimeoutSeconds,omitempty"`
				} `tfsdk:"deletion" json:"deletion,omitempty"`
				HealthCheck *struct {
					Checks *struct {
						NodeStartupTimeoutSeconds *int64 `tfsdk:"node_startup_timeout_seconds" json:"nodeStartupTimeoutSeconds,omitempty"`
						UnhealthyNodeConditions   *[]struct {
							Status         *string `tfsdk:"status" json:"status,omitempty"`
							TimeoutSeconds *int64  `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
							Type           *string `tfsdk:"type" json:"type,omitempty"`
						} `tfsdk:"unhealthy_node_conditions" json:"unhealthyNodeConditions,omitempty"`
					} `tfsdk:"checks" json:"checks,omitempty"`
					Enabled     *bool `tfsdk:"enabled" json:"enabled,omitempty"`
					Remediation *struct {
						TemplateRef *struct {
							ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
							Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
							Name       *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"template_ref" json:"templateRef,omitempty"`
						TriggerIf *struct {
							UnhealthyInRange           *string `tfsdk:"unhealthy_in_range" json:"unhealthyInRange,omitempty"`
							UnhealthyLessThanOrEqualTo *string `tfsdk:"unhealthy_less_than_or_equal_to" json:"unhealthyLessThanOrEqualTo,omitempty"`
						} `tfsdk:"trigger_if" json:"triggerIf,omitempty"`
					} `tfsdk:"remediation" json:"remediation,omitempty"`
				} `tfsdk:"health_check" json:"healthCheck,omitempty"`
				Metadata *struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				ReadinessGates *[]struct {
					ConditionType *string `tfsdk:"condition_type" json:"conditionType,omitempty"`
					Polarity      *string `tfsdk:"polarity" json:"polarity,omitempty"`
				} `tfsdk:"readiness_gates" json:"readinessGates,omitempty"`
				Replicas  *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
				Variables *struct {
					Overrides *[]struct {
						Name  *string            `tfsdk:"name" json:"name,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"overrides" json:"overrides,omitempty"`
				} `tfsdk:"variables" json:"variables,omitempty"`
			} `tfsdk:"control_plane" json:"controlPlane,omitempty"`
			Variables *[]struct {
				Name  *string            `tfsdk:"name" json:"name,omitempty"`
				Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"variables" json:"variables,omitempty"`
			Version *string `tfsdk:"version" json:"version,omitempty"`
			Workers *struct {
				MachineDeployments *[]struct {
					Class    *string `tfsdk:"class" json:"class,omitempty"`
					Deletion *struct {
						NodeDeletionTimeoutSeconds     *int64  `tfsdk:"node_deletion_timeout_seconds" json:"nodeDeletionTimeoutSeconds,omitempty"`
						NodeDrainTimeoutSeconds        *int64  `tfsdk:"node_drain_timeout_seconds" json:"nodeDrainTimeoutSeconds,omitempty"`
						NodeVolumeDetachTimeoutSeconds *int64  `tfsdk:"node_volume_detach_timeout_seconds" json:"nodeVolumeDetachTimeoutSeconds,omitempty"`
						Order                          *string `tfsdk:"order" json:"order,omitempty"`
					} `tfsdk:"deletion" json:"deletion,omitempty"`
					FailureDomain *string `tfsdk:"failure_domain" json:"failureDomain,omitempty"`
					HealthCheck   *struct {
						Checks *struct {
							NodeStartupTimeoutSeconds *int64 `tfsdk:"node_startup_timeout_seconds" json:"nodeStartupTimeoutSeconds,omitempty"`
							UnhealthyNodeConditions   *[]struct {
								Status         *string `tfsdk:"status" json:"status,omitempty"`
								TimeoutSeconds *int64  `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
								Type           *string `tfsdk:"type" json:"type,omitempty"`
							} `tfsdk:"unhealthy_node_conditions" json:"unhealthyNodeConditions,omitempty"`
						} `tfsdk:"checks" json:"checks,omitempty"`
						Enabled     *bool `tfsdk:"enabled" json:"enabled,omitempty"`
						Remediation *struct {
							MaxInFlight *string `tfsdk:"max_in_flight" json:"maxInFlight,omitempty"`
							TemplateRef *struct {
								ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
								Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
								Name       *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"template_ref" json:"templateRef,omitempty"`
							TriggerIf *struct {
								UnhealthyInRange           *string `tfsdk:"unhealthy_in_range" json:"unhealthyInRange,omitempty"`
								UnhealthyLessThanOrEqualTo *string `tfsdk:"unhealthy_less_than_or_equal_to" json:"unhealthyLessThanOrEqualTo,omitempty"`
							} `tfsdk:"trigger_if" json:"triggerIf,omitempty"`
						} `tfsdk:"remediation" json:"remediation,omitempty"`
					} `tfsdk:"health_check" json:"healthCheck,omitempty"`
					Metadata *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
					MinReadySeconds *int64  `tfsdk:"min_ready_seconds" json:"minReadySeconds,omitempty"`
					Name            *string `tfsdk:"name" json:"name,omitempty"`
					ReadinessGates  *[]struct {
						ConditionType *string `tfsdk:"condition_type" json:"conditionType,omitempty"`
						Polarity      *string `tfsdk:"polarity" json:"polarity,omitempty"`
					} `tfsdk:"readiness_gates" json:"readinessGates,omitempty"`
					Replicas *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
					Rollout  *struct {
						Strategy *struct {
							RollingUpdate *struct {
								MaxSurge       *string `tfsdk:"max_surge" json:"maxSurge,omitempty"`
								MaxUnavailable *string `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
							} `tfsdk:"rolling_update" json:"rollingUpdate,omitempty"`
							Type *string `tfsdk:"type" json:"type,omitempty"`
						} `tfsdk:"strategy" json:"strategy,omitempty"`
					} `tfsdk:"rollout" json:"rollout,omitempty"`
					Variables *struct {
						Overrides *[]struct {
							Name  *string            `tfsdk:"name" json:"name,omitempty"`
							Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"overrides" json:"overrides,omitempty"`
					} `tfsdk:"variables" json:"variables,omitempty"`
				} `tfsdk:"machine_deployments" json:"machineDeployments,omitempty"`
				MachinePools *[]struct {
					Class    *string `tfsdk:"class" json:"class,omitempty"`
					Deletion *struct {
						NodeDeletionTimeoutSeconds     *int64 `tfsdk:"node_deletion_timeout_seconds" json:"nodeDeletionTimeoutSeconds,omitempty"`
						NodeDrainTimeoutSeconds        *int64 `tfsdk:"node_drain_timeout_seconds" json:"nodeDrainTimeoutSeconds,omitempty"`
						NodeVolumeDetachTimeoutSeconds *int64 `tfsdk:"node_volume_detach_timeout_seconds" json:"nodeVolumeDetachTimeoutSeconds,omitempty"`
					} `tfsdk:"deletion" json:"deletion,omitempty"`
					FailureDomains *[]string `tfsdk:"failure_domains" json:"failureDomains,omitempty"`
					Metadata       *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
					MinReadySeconds *int64  `tfsdk:"min_ready_seconds" json:"minReadySeconds,omitempty"`
					Name            *string `tfsdk:"name" json:"name,omitempty"`
					Replicas        *int64  `tfsdk:"replicas" json:"replicas,omitempty"`
					Variables       *struct {
						Overrides *[]struct {
							Name  *string            `tfsdk:"name" json:"name,omitempty"`
							Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"overrides" json:"overrides,omitempty"`
					} `tfsdk:"variables" json:"variables,omitempty"`
				} `tfsdk:"machine_pools" json:"machinePools,omitempty"`
			} `tfsdk:"workers" json:"workers,omitempty"`
		} `tfsdk:"topology" json:"topology,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ClusterXK8SIoClusterV1Beta2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_cluster_x_k8s_io_cluster_v1beta2_manifest"
}

func (r *ClusterXK8SIoClusterV1Beta2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Cluster is the Schema for the clusters API.",
		MarkdownDescription: "Cluster is the Schema for the clusters API.",
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
				Description:         "spec is the desired state of Cluster.",
				MarkdownDescription: "spec is the desired state of Cluster.",
				Attributes: map[string]schema.Attribute{
					"availability_gates": schema.ListNestedAttribute{
						Description:         "availabilityGates specifies additional conditions to include when evaluating Cluster Available condition. If this field is not defined and the Cluster implements a managed topology, availabilityGates from the corresponding ClusterClass will be used, if any.",
						MarkdownDescription: "availabilityGates specifies additional conditions to include when evaluating Cluster Available condition. If this field is not defined and the Cluster implements a managed topology, availabilityGates from the corresponding ClusterClass will be used, if any.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"condition_type": schema.StringAttribute{
									Description:         "conditionType refers to a condition with matching type in the Cluster's condition list. If the conditions doesn't exist, it will be treated as unknown. Note: Both Cluster API conditions or conditions added by 3rd party controllers can be used as availability gates.",
									MarkdownDescription: "conditionType refers to a condition with matching type in the Cluster's condition list. If the conditions doesn't exist, it will be treated as unknown. Note: Both Cluster API conditions or conditions added by 3rd party controllers can be used as availability gates.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(316),
										stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*/)?(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])$`), ""),
									},
								},

								"polarity": schema.StringAttribute{
									Description:         "polarity of the conditionType specified in this availabilityGate. Valid values are Positive, Negative and omitted. When omitted, the default behaviour will be Positive. A positive polarity means that the condition should report a true status under normal conditions. A negative polarity means that the condition should report a false status under normal conditions.",
									MarkdownDescription: "polarity of the conditionType specified in this availabilityGate. Valid values are Positive, Negative and omitted. When omitted, the default behaviour will be Positive. A positive polarity means that the condition should report a true status under normal conditions. A negative polarity means that the condition should report a false status under normal conditions.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("Positive", "Negative"),
									},
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
						Validators: []validator.List{
							listvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("cluster_network"), path.MatchRelative().AtParent().AtName("control_plane_endpoint"), path.MatchRelative().AtParent().AtName("control_plane_ref"), path.MatchRelative().AtParent().AtName("infrastructure_ref"), path.MatchRelative().AtParent().AtName("paused"), path.MatchRelative().AtParent().AtName("topology")),
						},
					},

					"cluster_network": schema.SingleNestedAttribute{
						Description:         "clusterNetwork represents the cluster network configuration.",
						MarkdownDescription: "clusterNetwork represents the cluster network configuration.",
						Attributes: map[string]schema.Attribute{
							"api_server_port": schema.Int64Attribute{
								Description:         "apiServerPort specifies the port the API Server should bind to. Defaults to 6443.",
								MarkdownDescription: "apiServerPort specifies the port the API Server should bind to. Defaults to 6443.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(65535),
									int64validator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("pods"), path.MatchRelative().AtParent().AtName("service_domain"), path.MatchRelative().AtParent().AtName("services")),
								},
							},

							"pods": schema.SingleNestedAttribute{
								Description:         "pods is the network ranges from which Pod networks are allocated.",
								MarkdownDescription: "pods is the network ranges from which Pod networks are allocated.",
								Attributes: map[string]schema.Attribute{
									"cidr_blocks": schema.ListAttribute{
										Description:         "cidrBlocks is a list of CIDR blocks.",
										MarkdownDescription: "cidrBlocks is a list of CIDR blocks.",
										ElementType:         types.StringType,
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
								Validators: []validator.Object{
									objectvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("api_server_port"), path.MatchRelative().AtParent().AtName("service_domain"), path.MatchRelative().AtParent().AtName("services")),
								},
							},

							"service_domain": schema.StringAttribute{
								Description:         "serviceDomain is the domain name for services.",
								MarkdownDescription: "serviceDomain is the domain name for services.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(253),
									stringvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("api_server_port"), path.MatchRelative().AtParent().AtName("pods"), path.MatchRelative().AtParent().AtName("services")),
								},
							},

							"services": schema.SingleNestedAttribute{
								Description:         "services is the network ranges from which service VIPs are allocated.",
								MarkdownDescription: "services is the network ranges from which service VIPs are allocated.",
								Attributes: map[string]schema.Attribute{
									"cidr_blocks": schema.ListAttribute{
										Description:         "cidrBlocks is a list of CIDR blocks.",
										MarkdownDescription: "cidrBlocks is a list of CIDR blocks.",
										ElementType:         types.StringType,
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
								Validators: []validator.Object{
									objectvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("api_server_port"), path.MatchRelative().AtParent().AtName("pods"), path.MatchRelative().AtParent().AtName("service_domain")),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
						Validators: []validator.Object{
							objectvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("availability_gates"), path.MatchRelative().AtParent().AtName("control_plane_endpoint"), path.MatchRelative().AtParent().AtName("control_plane_ref"), path.MatchRelative().AtParent().AtName("infrastructure_ref"), path.MatchRelative().AtParent().AtName("paused"), path.MatchRelative().AtParent().AtName("topology")),
						},
					},

					"control_plane_endpoint": schema.SingleNestedAttribute{
						Description:         "controlPlaneEndpoint represents the endpoint used to communicate with the control plane.",
						MarkdownDescription: "controlPlaneEndpoint represents the endpoint used to communicate with the control plane.",
						Attributes: map[string]schema.Attribute{
							"host": schema.StringAttribute{
								Description:         "host is the hostname on which the API server is serving.",
								MarkdownDescription: "host is the hostname on which the API server is serving.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(512),
									stringvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("port")),
								},
							},

							"port": schema.Int64Attribute{
								Description:         "port is the port on which the API server is serving.",
								MarkdownDescription: "port is the port on which the API server is serving.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(65535),
									int64validator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("host")),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
						Validators: []validator.Object{
							objectvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("availability_gates"), path.MatchRelative().AtParent().AtName("cluster_network"), path.MatchRelative().AtParent().AtName("control_plane_ref"), path.MatchRelative().AtParent().AtName("infrastructure_ref"), path.MatchRelative().AtParent().AtName("paused"), path.MatchRelative().AtParent().AtName("topology")),
						},
					},

					"control_plane_ref": schema.SingleNestedAttribute{
						Description:         "controlPlaneRef is an optional reference to a provider-specific resource that holds the details for provisioning the Control Plane for a Cluster.",
						MarkdownDescription: "controlPlaneRef is an optional reference to a provider-specific resource that holds the details for provisioning the Control Plane for a Cluster.",
						Attributes: map[string]schema.Attribute{
							"api_group": schema.StringAttribute{
								Description:         "apiGroup is the group of the resource being referenced. apiGroup must be fully qualified domain name. The corresponding version for this reference will be looked up from the contract labels of the corresponding CRD of the resource being referenced.",
								MarkdownDescription: "apiGroup is the group of the resource being referenced. apiGroup must be fully qualified domain name. The corresponding version for this reference will be looked up from the contract labels of the corresponding CRD of the resource being referenced.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(253),
									stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
								},
							},

							"kind": schema.StringAttribute{
								Description:         "kind of the resource being referenced. kind must consist of alphanumeric characters or '-', start with an alphabetic character, and end with an alphanumeric character.",
								MarkdownDescription: "kind of the resource being referenced. kind must consist of alphanumeric characters or '-', start with an alphabetic character, and end with an alphanumeric character.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(63),
									stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z]([-a-zA-Z0-9]*[a-zA-Z0-9])?$`), ""),
								},
							},

							"name": schema.StringAttribute{
								Description:         "name of the resource being referenced. name must consist of lower case alphanumeric characters, '-' or '.', and must start and end with an alphanumeric character.",
								MarkdownDescription: "name of the resource being referenced. name must consist of lower case alphanumeric characters, '-' or '.', and must start and end with an alphanumeric character.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(253),
									stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
						Validators: []validator.Object{
							objectvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("availability_gates"), path.MatchRelative().AtParent().AtName("cluster_network"), path.MatchRelative().AtParent().AtName("control_plane_endpoint"), path.MatchRelative().AtParent().AtName("infrastructure_ref"), path.MatchRelative().AtParent().AtName("paused"), path.MatchRelative().AtParent().AtName("topology")),
						},
					},

					"infrastructure_ref": schema.SingleNestedAttribute{
						Description:         "infrastructureRef is a reference to a provider-specific resource that holds the details for provisioning infrastructure for a cluster in said provider.",
						MarkdownDescription: "infrastructureRef is a reference to a provider-specific resource that holds the details for provisioning infrastructure for a cluster in said provider.",
						Attributes: map[string]schema.Attribute{
							"api_group": schema.StringAttribute{
								Description:         "apiGroup is the group of the resource being referenced. apiGroup must be fully qualified domain name. The corresponding version for this reference will be looked up from the contract labels of the corresponding CRD of the resource being referenced.",
								MarkdownDescription: "apiGroup is the group of the resource being referenced. apiGroup must be fully qualified domain name. The corresponding version for this reference will be looked up from the contract labels of the corresponding CRD of the resource being referenced.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(253),
									stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
								},
							},

							"kind": schema.StringAttribute{
								Description:         "kind of the resource being referenced. kind must consist of alphanumeric characters or '-', start with an alphabetic character, and end with an alphanumeric character.",
								MarkdownDescription: "kind of the resource being referenced. kind must consist of alphanumeric characters or '-', start with an alphabetic character, and end with an alphanumeric character.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(63),
									stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z]([-a-zA-Z0-9]*[a-zA-Z0-9])?$`), ""),
								},
							},

							"name": schema.StringAttribute{
								Description:         "name of the resource being referenced. name must consist of lower case alphanumeric characters, '-' or '.', and must start and end with an alphanumeric character.",
								MarkdownDescription: "name of the resource being referenced. name must consist of lower case alphanumeric characters, '-' or '.', and must start and end with an alphanumeric character.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(253),
									stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
						Validators: []validator.Object{
							objectvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("availability_gates"), path.MatchRelative().AtParent().AtName("cluster_network"), path.MatchRelative().AtParent().AtName("control_plane_endpoint"), path.MatchRelative().AtParent().AtName("control_plane_ref"), path.MatchRelative().AtParent().AtName("paused"), path.MatchRelative().AtParent().AtName("topology")),
						},
					},

					"paused": schema.BoolAttribute{
						Description:         "paused can be used to prevent controllers from processing the Cluster and all its associated objects.",
						MarkdownDescription: "paused can be used to prevent controllers from processing the Cluster and all its associated objects.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Bool{
							boolvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("availability_gates"), path.MatchRelative().AtParent().AtName("cluster_network"), path.MatchRelative().AtParent().AtName("control_plane_endpoint"), path.MatchRelative().AtParent().AtName("control_plane_ref"), path.MatchRelative().AtParent().AtName("infrastructure_ref"), path.MatchRelative().AtParent().AtName("topology")),
						},
					},

					"topology": schema.SingleNestedAttribute{
						Description:         "topology encapsulates the topology for the cluster. NOTE: It is required to enable the ClusterTopology feature gate flag to activate managed topologies support; this feature is highly experimental, and parts of it might still be not implemented.",
						MarkdownDescription: "topology encapsulates the topology for the cluster. NOTE: It is required to enable the ClusterTopology feature gate flag to activate managed topologies support; this feature is highly experimental, and parts of it might still be not implemented.",
						Attributes: map[string]schema.Attribute{
							"class_ref": schema.SingleNestedAttribute{
								Description:         "classRef is the ref to the ClusterClass that should be used for the topology.",
								MarkdownDescription: "classRef is the ref to the ClusterClass that should be used for the topology.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "name is the name of the ClusterClass that should be used for the topology. name must be a valid ClusterClass name and because of that be at most 253 characters in length and it must consist only of lower case alphanumeric characters, hyphens (-) and periods (.), and must start and end with an alphanumeric character.",
										MarkdownDescription: "name is the name of the ClusterClass that should be used for the topology. name must be a valid ClusterClass name and because of that be at most 253 characters in length and it must consist only of lower case alphanumeric characters, hyphens (-) and periods (.), and must start and end with an alphanumeric character.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(1),
											stringvalidator.LengthAtMost(253),
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
										},
									},

									"namespace": schema.StringAttribute{
										Description:         "namespace is the namespace of the ClusterClass that should be used for the topology. If namespace is empty or not set, it is defaulted to the namespace of the Cluster object. namespace must be a valid namespace name and because of that be at most 63 characters in length and it must consist only of lower case alphanumeric characters or hyphens (-), and must start and end with an alphanumeric character.",
										MarkdownDescription: "namespace is the namespace of the ClusterClass that should be used for the topology. If namespace is empty or not set, it is defaulted to the namespace of the Cluster object. namespace must be a valid namespace name and because of that be at most 63 characters in length and it must consist only of lower case alphanumeric characters or hyphens (-), and must start and end with an alphanumeric character.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(1),
											stringvalidator.LengthAtMost(63),
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
										},
									},
								},
								Required: true,
								Optional: false,
								Computed: false,
							},

							"control_plane": schema.SingleNestedAttribute{
								Description:         "controlPlane describes the cluster control plane.",
								MarkdownDescription: "controlPlane describes the cluster control plane.",
								Attributes: map[string]schema.Attribute{
									"deletion": schema.SingleNestedAttribute{
										Description:         "deletion contains configuration options for Machine deletion.",
										MarkdownDescription: "deletion contains configuration options for Machine deletion.",
										Attributes: map[string]schema.Attribute{
											"node_deletion_timeout_seconds": schema.Int64Attribute{
												Description:         "nodeDeletionTimeoutSeconds defines how long the controller will attempt to delete the Node that the Machine hosts after the Machine is marked for deletion. A duration of 0 will retry deletion indefinitely. Defaults to 10 seconds.",
												MarkdownDescription: "nodeDeletionTimeoutSeconds defines how long the controller will attempt to delete the Node that the Machine hosts after the Machine is marked for deletion. A duration of 0 will retry deletion indefinitely. Defaults to 10 seconds.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("node_drain_timeout_seconds"), path.MatchRelative().AtParent().AtName("node_volume_detach_timeout_seconds")),
												},
											},

											"node_drain_timeout_seconds": schema.Int64Attribute{
												Description:         "nodeDrainTimeoutSeconds is the total amount of time that the controller will spend on draining a node. The default value is 0, meaning that the node can be drained without any time limitations. NOTE: nodeDrainTimeoutSeconds is different from 'kubectl drain --timeout'",
												MarkdownDescription: "nodeDrainTimeoutSeconds is the total amount of time that the controller will spend on draining a node. The default value is 0, meaning that the node can be drained without any time limitations. NOTE: nodeDrainTimeoutSeconds is different from 'kubectl drain --timeout'",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("node_deletion_timeout_seconds"), path.MatchRelative().AtParent().AtName("node_volume_detach_timeout_seconds")),
												},
											},

											"node_volume_detach_timeout_seconds": schema.Int64Attribute{
												Description:         "nodeVolumeDetachTimeoutSeconds is the total amount of time that the controller will spend on waiting for all volumes to be detached. The default value is 0, meaning that the volumes can be detached without any time limitations.",
												MarkdownDescription: "nodeVolumeDetachTimeoutSeconds is the total amount of time that the controller will spend on waiting for all volumes to be detached. The default value is 0, meaning that the volumes can be detached without any time limitations.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("node_deletion_timeout_seconds"), path.MatchRelative().AtParent().AtName("node_drain_timeout_seconds")),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
										Validators: []validator.Object{
											objectvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("health_check"), path.MatchRelative().AtParent().AtName("metadata"), path.MatchRelative().AtParent().AtName("readiness_gates"), path.MatchRelative().AtParent().AtName("replicas"), path.MatchRelative().AtParent().AtName("variables")),
										},
									},

									"health_check": schema.SingleNestedAttribute{
										Description:         "healthCheck allows to enable, disable and override control plane health check configuration from the ClusterClass for this control plane.",
										MarkdownDescription: "healthCheck allows to enable, disable and override control plane health check configuration from the ClusterClass for this control plane.",
										Attributes: map[string]schema.Attribute{
											"checks": schema.SingleNestedAttribute{
												Description:         "checks are the checks that are used to evaluate if a Machine is healthy. If one of checks and remediation fields are set, the system assumes that an healthCheck override is defined, and as a consequence the checks and remediation fields from Cluster will be used instead of the corresponding fields in ClusterClass. Independent of this configuration the MachineHealthCheck controller will always flag Machines with 'cluster.x-k8s.io/remediate-machine' annotation and Machines with deleted Nodes as unhealthy. Furthermore, if checks.nodeStartupTimeoutSeconds is not set it is defaulted to 10 minutes and evaluated accordingly.",
												MarkdownDescription: "checks are the checks that are used to evaluate if a Machine is healthy. If one of checks and remediation fields are set, the system assumes that an healthCheck override is defined, and as a consequence the checks and remediation fields from Cluster will be used instead of the corresponding fields in ClusterClass. Independent of this configuration the MachineHealthCheck controller will always flag Machines with 'cluster.x-k8s.io/remediate-machine' annotation and Machines with deleted Nodes as unhealthy. Furthermore, if checks.nodeStartupTimeoutSeconds is not set it is defaulted to 10 minutes and evaluated accordingly.",
												Attributes: map[string]schema.Attribute{
													"node_startup_timeout_seconds": schema.Int64Attribute{
														Description:         "nodeStartupTimeoutSeconds allows to set the maximum time for MachineHealthCheck to consider a Machine unhealthy if a corresponding Node isn't associated through a 'Spec.ProviderID' field. The duration set in this field is compared to the greatest of: - Cluster's infrastructure ready condition timestamp (if and when available) - Control Plane's initialized condition timestamp (if and when available) - Machine's infrastructure ready condition timestamp (if and when available) - Machine's metadata creation timestamp Defaults to 10 minutes. If you wish to disable this feature, set the value explicitly to 0.",
														MarkdownDescription: "nodeStartupTimeoutSeconds allows to set the maximum time for MachineHealthCheck to consider a Machine unhealthy if a corresponding Node isn't associated through a 'Spec.ProviderID' field. The duration set in this field is compared to the greatest of: - Cluster's infrastructure ready condition timestamp (if and when available) - Control Plane's initialized condition timestamp (if and when available) - Machine's infrastructure ready condition timestamp (if and when available) - Machine's metadata creation timestamp Defaults to 10 minutes. If you wish to disable this feature, set the value explicitly to 0.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(0),
															int64validator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("unhealthy_node_conditions")),
														},
													},

													"unhealthy_node_conditions": schema.ListNestedAttribute{
														Description:         "unhealthyNodeConditions contains a list of conditions that determine whether a node is considered unhealthy. The conditions are combined in a logical OR, i.e. if any of the conditions is met, the node is unhealthy.",
														MarkdownDescription: "unhealthyNodeConditions contains a list of conditions that determine whether a node is considered unhealthy. The conditions are combined in a logical OR, i.e. if any of the conditions is met, the node is unhealthy.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"status": schema.StringAttribute{
																	Description:         "status of the condition, one of True, False, Unknown.",
																	MarkdownDescription: "status of the condition, one of True, False, Unknown.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtLeast(1),
																	},
																},

																"timeout_seconds": schema.Int64Attribute{
																	Description:         "timeoutSeconds is the duration that a node must be in a given status for, after which the node is considered unhealthy. For example, with a value of '1h', the node must match the status for at least 1 hour before being considered unhealthy.",
																	MarkdownDescription: "timeoutSeconds is the duration that a node must be in a given status for, after which the node is considered unhealthy. For example, with a value of '1h', the node must match the status for at least 1 hour before being considered unhealthy.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(0),
																	},
																},

																"type": schema.StringAttribute{
																	Description:         "type of Node condition",
																	MarkdownDescription: "type of Node condition",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtLeast(1),
																	},
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
														Validators: []validator.List{
															listvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("node_startup_timeout_seconds")),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
												Validators: []validator.Object{
													objectvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("enabled"), path.MatchRelative().AtParent().AtName("remediation")),
												},
											},

											"enabled": schema.BoolAttribute{
												Description:         "enabled controls if a MachineHealthCheck should be created for the target machines. If false: No MachineHealthCheck will be created. If not set(default): A MachineHealthCheck will be created if it is defined here or in the associated ClusterClass. If no MachineHealthCheck is defined then none will be created. If true: A MachineHealthCheck is guaranteed to be created. Cluster validation will block if 'enable' is true and no MachineHealthCheck definition is available.",
												MarkdownDescription: "enabled controls if a MachineHealthCheck should be created for the target machines. If false: No MachineHealthCheck will be created. If not set(default): A MachineHealthCheck will be created if it is defined here or in the associated ClusterClass. If no MachineHealthCheck is defined then none will be created. If true: A MachineHealthCheck is guaranteed to be created. Cluster validation will block if 'enable' is true and no MachineHealthCheck definition is available.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Bool{
													boolvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("checks"), path.MatchRelative().AtParent().AtName("remediation")),
												},
											},

											"remediation": schema.SingleNestedAttribute{
												Description:         "remediation configures if and how remediations are triggered if a Machine is unhealthy. If one of checks and remediation fields are set, the system assumes that an healthCheck override is defined, and as a consequence the checks and remediation fields from cluster will be used instead of the corresponding fields in ClusterClass. If an health check override is defined and remediation or remediation.triggerIf is not set, remediation will always be triggered for unhealthy Machines. If an health check override is defined and remediation or remediation.templateRef is not set, the OwnerRemediated condition will be set on unhealthy Machines to trigger remediation via the owner of the Machines, for example a MachineSet or a KubeadmControlPlane.",
												MarkdownDescription: "remediation configures if and how remediations are triggered if a Machine is unhealthy. If one of checks and remediation fields are set, the system assumes that an healthCheck override is defined, and as a consequence the checks and remediation fields from cluster will be used instead of the corresponding fields in ClusterClass. If an health check override is defined and remediation or remediation.triggerIf is not set, remediation will always be triggered for unhealthy Machines. If an health check override is defined and remediation or remediation.templateRef is not set, the OwnerRemediated condition will be set on unhealthy Machines to trigger remediation via the owner of the Machines, for example a MachineSet or a KubeadmControlPlane.",
												Attributes: map[string]schema.Attribute{
													"template_ref": schema.SingleNestedAttribute{
														Description:         "templateRef is a reference to a remediation template provided by an infrastructure provider. This field is completely optional, when filled, the MachineHealthCheck controller creates a new object from the template referenced and hands off remediation of the machine to a controller that lives outside of Cluster API.",
														MarkdownDescription: "templateRef is a reference to a remediation template provided by an infrastructure provider. This field is completely optional, when filled, the MachineHealthCheck controller creates a new object from the template referenced and hands off remediation of the machine to a controller that lives outside of Cluster API.",
														Attributes: map[string]schema.Attribute{
															"api_version": schema.StringAttribute{
																Description:         "apiVersion of the remediation template. apiVersion must be fully qualified domain name followed by / and a version. NOTE: This field must be kept in sync with the APIVersion of the remediation template.",
																MarkdownDescription: "apiVersion of the remediation template. apiVersion must be fully qualified domain name followed by / and a version. NOTE: This field must be kept in sync with the APIVersion of the remediation template.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																	stringvalidator.LengthAtMost(317),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*\/[a-z]([-a-z0-9]*[a-z0-9])?$`), ""),
																},
															},

															"kind": schema.StringAttribute{
																Description:         "kind of the remediation template. kind must consist of alphanumeric characters or '-', start with an alphabetic character, and end with an alphanumeric character.",
																MarkdownDescription: "kind of the remediation template. kind must consist of alphanumeric characters or '-', start with an alphabetic character, and end with an alphanumeric character.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																	stringvalidator.LengthAtMost(63),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z]([-a-zA-Z0-9]*[a-zA-Z0-9])?$`), ""),
																},
															},

															"name": schema.StringAttribute{
																Description:         "name of the remediation template. name must consist of lower case alphanumeric characters, '-' or '.', and must start and end with an alphanumeric character.",
																MarkdownDescription: "name of the remediation template. name must consist of lower case alphanumeric characters, '-' or '.', and must start and end with an alphanumeric character.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																	stringvalidator.LengthAtMost(253),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
														Validators: []validator.Object{
															objectvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("trigger_if")),
														},
													},

													"trigger_if": schema.SingleNestedAttribute{
														Description:         "triggerIf configures if remediations are triggered. If this field is not set, remediations are always triggered.",
														MarkdownDescription: "triggerIf configures if remediations are triggered. If this field is not set, remediations are always triggered.",
														Attributes: map[string]schema.Attribute{
															"unhealthy_in_range": schema.StringAttribute{
																Description:         "unhealthyInRange specifies that remediations are only triggered if the number of unhealthy Machines is in the configured range. Takes precedence over unhealthyLessThanOrEqualTo. Eg. '[3-5]' - This means that remediation will be allowed only when: (a) there are at least 3 unhealthy Machines (and) (b) there are at most 5 unhealthy Machines",
																MarkdownDescription: "unhealthyInRange specifies that remediations are only triggered if the number of unhealthy Machines is in the configured range. Takes precedence over unhealthyLessThanOrEqualTo. Eg. '[3-5]' - This means that remediation will be allowed only when: (a) there are at least 3 unhealthy Machines (and) (b) there are at most 5 unhealthy Machines",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																	stringvalidator.LengthAtMost(32),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^\[[0-9]+-[0-9]+\]$`), ""),
																	stringvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("unhealthy_less_than_or_equal_to")),
																},
															},

															"unhealthy_less_than_or_equal_to": schema.StringAttribute{
																Description:         "unhealthyLessThanOrEqualTo specifies that remediations are only triggered if the number of unhealthy Machines is less than or equal to the configured value. unhealthyInRange takes precedence if set.",
																MarkdownDescription: "unhealthyLessThanOrEqualTo specifies that remediations are only triggered if the number of unhealthy Machines is less than or equal to the configured value. unhealthyInRange takes precedence if set.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("unhealthy_in_range")),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
														Validators: []validator.Object{
															objectvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("template_ref")),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
												Validators: []validator.Object{
													objectvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("checks"), path.MatchRelative().AtParent().AtName("enabled")),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
										Validators: []validator.Object{
											objectvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("deletion"), path.MatchRelative().AtParent().AtName("metadata"), path.MatchRelative().AtParent().AtName("readiness_gates"), path.MatchRelative().AtParent().AtName("replicas"), path.MatchRelative().AtParent().AtName("variables")),
										},
									},

									"metadata": schema.SingleNestedAttribute{
										Description:         "metadata is the metadata applied to the ControlPlane and the Machines of the ControlPlane if the ControlPlaneTemplate referenced by the ClusterClass is machine based. If not, it is applied only to the ControlPlane. At runtime this metadata is merged with the corresponding metadata from the ClusterClass.",
										MarkdownDescription: "metadata is the metadata applied to the ControlPlane and the Machines of the ControlPlane if the ControlPlaneTemplate referenced by the ClusterClass is machine based. If not, it is applied only to the ControlPlane. At runtime this metadata is merged with the corresponding metadata from the ClusterClass.",
										Attributes: map[string]schema.Attribute{
											"annotations": schema.MapAttribute{
												Description:         "annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
												MarkdownDescription: "annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Map{
													mapvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("labels")),
												},
											},

											"labels": schema.MapAttribute{
												Description:         "labels is a map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
												MarkdownDescription: "labels is a map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Map{
													mapvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("annotations")),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
										Validators: []validator.Object{
											objectvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("deletion"), path.MatchRelative().AtParent().AtName("health_check"), path.MatchRelative().AtParent().AtName("readiness_gates"), path.MatchRelative().AtParent().AtName("replicas"), path.MatchRelative().AtParent().AtName("variables")),
										},
									},

									"readiness_gates": schema.ListNestedAttribute{
										Description:         "readinessGates specifies additional conditions to include when evaluating Machine Ready condition. This field can be used e.g. to instruct the machine controller to include in the computation for Machine's ready computation a condition, managed by an external controllers, reporting the status of special software/hardware installed on the Machine. If this field is not defined, readinessGates from the corresponding ControlPlaneClass will be used, if any. NOTE: Specific control plane provider implementations might automatically extend the list of readinessGates; e.g. the kubeadm control provider adds ReadinessGates for the APIServerPodHealthy, SchedulerPodHealthy conditions, etc.",
										MarkdownDescription: "readinessGates specifies additional conditions to include when evaluating Machine Ready condition. This field can be used e.g. to instruct the machine controller to include in the computation for Machine's ready computation a condition, managed by an external controllers, reporting the status of special software/hardware installed on the Machine. If this field is not defined, readinessGates from the corresponding ControlPlaneClass will be used, if any. NOTE: Specific control plane provider implementations might automatically extend the list of readinessGates; e.g. the kubeadm control provider adds ReadinessGates for the APIServerPodHealthy, SchedulerPodHealthy conditions, etc.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"condition_type": schema.StringAttribute{
													Description:         "conditionType refers to a condition with matching type in the Machine's condition list. If the conditions doesn't exist, it will be treated as unknown. Note: Both Cluster API conditions or conditions added by 3rd party controllers can be used as readiness gates.",
													MarkdownDescription: "conditionType refers to a condition with matching type in the Machine's condition list. If the conditions doesn't exist, it will be treated as unknown. Note: Both Cluster API conditions or conditions added by 3rd party controllers can be used as readiness gates.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
														stringvalidator.LengthAtMost(316),
														stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*/)?(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])$`), ""),
													},
												},

												"polarity": schema.StringAttribute{
													Description:         "polarity of the conditionType specified in this readinessGate. Valid values are Positive, Negative and omitted. When omitted, the default behaviour will be Positive. A positive polarity means that the condition should report a true status under normal conditions. A negative polarity means that the condition should report a false status under normal conditions.",
													MarkdownDescription: "polarity of the conditionType specified in this readinessGate. Valid values are Positive, Negative and omitted. When omitted, the default behaviour will be Positive. A positive polarity means that the condition should report a true status under normal conditions. A negative polarity means that the condition should report a false status under normal conditions.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("Positive", "Negative"),
													},
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
										Validators: []validator.List{
											listvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("deletion"), path.MatchRelative().AtParent().AtName("health_check"), path.MatchRelative().AtParent().AtName("metadata"), path.MatchRelative().AtParent().AtName("replicas"), path.MatchRelative().AtParent().AtName("variables")),
										},
									},

									"replicas": schema.Int64Attribute{
										Description:         "replicas is the number of control plane nodes. If the value is not set, the ControlPlane object is created without the number of Replicas and it's assumed that the control plane controller does not implement support for this field. When specified against a control plane provider that lacks support for this field, this value will be ignored.",
										MarkdownDescription: "replicas is the number of control plane nodes. If the value is not set, the ControlPlane object is created without the number of Replicas and it's assumed that the control plane controller does not implement support for this field. When specified against a control plane provider that lacks support for this field, this value will be ignored.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("deletion"), path.MatchRelative().AtParent().AtName("health_check"), path.MatchRelative().AtParent().AtName("metadata"), path.MatchRelative().AtParent().AtName("readiness_gates"), path.MatchRelative().AtParent().AtName("variables")),
										},
									},

									"variables": schema.SingleNestedAttribute{
										Description:         "variables can be used to customize the ControlPlane through patches.",
										MarkdownDescription: "variables can be used to customize the ControlPlane through patches.",
										Attributes: map[string]schema.Attribute{
											"overrides": schema.ListNestedAttribute{
												Description:         "overrides can be used to override Cluster level variables.",
												MarkdownDescription: "overrides can be used to override Cluster level variables.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "name of the variable.",
															MarkdownDescription: "name of the variable.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtLeast(1),
																stringvalidator.LengthAtMost(256),
															},
														},

														"value": schema.MapAttribute{
															Description:         "value of the variable. Note: the value will be validated against the schema of the corresponding ClusterClassVariable from the ClusterClass. Note: We have to use apiextensionsv1.JSON instead of a custom JSON type, because controller-tools has a hard-coded schema for apiextensionsv1.JSON which cannot be produced by another type via controller-tools, i.e. it is not possible to have no type field. Ref: https://github.com/kubernetes-sigs/controller-tools/blob/d0e03a142d0ecdd5491593e941ee1d6b5d91dba6/pkg/crd/known_types.go#L106-L111",
															MarkdownDescription: "value of the variable. Note: the value will be validated against the schema of the corresponding ClusterClassVariable from the ClusterClass. Note: We have to use apiextensionsv1.JSON instead of a custom JSON type, because controller-tools has a hard-coded schema for apiextensionsv1.JSON which cannot be produced by another type via controller-tools, i.e. it is not possible to have no type field. Ref: https://github.com/kubernetes-sigs/controller-tools/blob/d0e03a142d0ecdd5491593e941ee1d6b5d91dba6/pkg/crd/known_types.go#L106-L111",
															ElementType:         types.StringType,
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
										Validators: []validator.Object{
											objectvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("deletion"), path.MatchRelative().AtParent().AtName("health_check"), path.MatchRelative().AtParent().AtName("metadata"), path.MatchRelative().AtParent().AtName("readiness_gates"), path.MatchRelative().AtParent().AtName("replicas")),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"variables": schema.ListNestedAttribute{
								Description:         "variables can be used to customize the Cluster through patches. They must comply to the corresponding VariableClasses defined in the ClusterClass.",
								MarkdownDescription: "variables can be used to customize the Cluster through patches. They must comply to the corresponding VariableClasses defined in the ClusterClass.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "name of the variable.",
											MarkdownDescription: "name of the variable.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
												stringvalidator.LengthAtMost(256),
											},
										},

										"value": schema.MapAttribute{
											Description:         "value of the variable. Note: the value will be validated against the schema of the corresponding ClusterClassVariable from the ClusterClass. Note: We have to use apiextensionsv1.JSON instead of a custom JSON type, because controller-tools has a hard-coded schema for apiextensionsv1.JSON which cannot be produced by another type via controller-tools, i.e. it is not possible to have no type field. Ref: https://github.com/kubernetes-sigs/controller-tools/blob/d0e03a142d0ecdd5491593e941ee1d6b5d91dba6/pkg/crd/known_types.go#L106-L111",
											MarkdownDescription: "value of the variable. Note: the value will be validated against the schema of the corresponding ClusterClassVariable from the ClusterClass. Note: We have to use apiextensionsv1.JSON instead of a custom JSON type, because controller-tools has a hard-coded schema for apiextensionsv1.JSON which cannot be produced by another type via controller-tools, i.e. it is not possible to have no type field. Ref: https://github.com/kubernetes-sigs/controller-tools/blob/d0e03a142d0ecdd5491593e941ee1d6b5d91dba6/pkg/crd/known_types.go#L106-L111",
											ElementType:         types.StringType,
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

							"version": schema.StringAttribute{
								Description:         "version is the Kubernetes version of the cluster.",
								MarkdownDescription: "version is the Kubernetes version of the cluster.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(256),
								},
							},

							"workers": schema.SingleNestedAttribute{
								Description:         "workers encapsulates the different constructs that form the worker nodes for the cluster.",
								MarkdownDescription: "workers encapsulates the different constructs that form the worker nodes for the cluster.",
								Attributes: map[string]schema.Attribute{
									"machine_deployments": schema.ListNestedAttribute{
										Description:         "machineDeployments is a list of machine deployments in the cluster.",
										MarkdownDescription: "machineDeployments is a list of machine deployments in the cluster.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"class": schema.StringAttribute{
													Description:         "class is the name of the MachineDeploymentClass used to create the set of worker nodes. This should match one of the deployment classes defined in the ClusterClass object mentioned in the 'Cluster.Spec.Class' field.",
													MarkdownDescription: "class is the name of the MachineDeploymentClass used to create the set of worker nodes. This should match one of the deployment classes defined in the ClusterClass object mentioned in the 'Cluster.Spec.Class' field.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
														stringvalidator.LengthAtMost(256),
													},
												},

												"deletion": schema.SingleNestedAttribute{
													Description:         "deletion contains configuration options for Machine deletion.",
													MarkdownDescription: "deletion contains configuration options for Machine deletion.",
													Attributes: map[string]schema.Attribute{
														"node_deletion_timeout_seconds": schema.Int64Attribute{
															Description:         "nodeDeletionTimeoutSeconds defines how long the controller will attempt to delete the Node that the Machine hosts after the Machine is marked for deletion. A duration of 0 will retry deletion indefinitely. Defaults to 10 seconds.",
															MarkdownDescription: "nodeDeletionTimeoutSeconds defines how long the controller will attempt to delete the Node that the Machine hosts after the Machine is marked for deletion. A duration of 0 will retry deletion indefinitely. Defaults to 10 seconds.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(0),
																int64validator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("node_drain_timeout_seconds"), path.MatchRelative().AtParent().AtName("node_volume_detach_timeout_seconds"), path.MatchRelative().AtParent().AtName("order")),
															},
														},

														"node_drain_timeout_seconds": schema.Int64Attribute{
															Description:         "nodeDrainTimeoutSeconds is the total amount of time that the controller will spend on draining a node. The default value is 0, meaning that the node can be drained without any time limitations. NOTE: nodeDrainTimeoutSeconds is different from 'kubectl drain --timeout'",
															MarkdownDescription: "nodeDrainTimeoutSeconds is the total amount of time that the controller will spend on draining a node. The default value is 0, meaning that the node can be drained without any time limitations. NOTE: nodeDrainTimeoutSeconds is different from 'kubectl drain --timeout'",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(0),
																int64validator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("node_deletion_timeout_seconds"), path.MatchRelative().AtParent().AtName("node_volume_detach_timeout_seconds"), path.MatchRelative().AtParent().AtName("order")),
															},
														},

														"node_volume_detach_timeout_seconds": schema.Int64Attribute{
															Description:         "nodeVolumeDetachTimeoutSeconds is the total amount of time that the controller will spend on waiting for all volumes to be detached. The default value is 0, meaning that the volumes can be detached without any time limitations.",
															MarkdownDescription: "nodeVolumeDetachTimeoutSeconds is the total amount of time that the controller will spend on waiting for all volumes to be detached. The default value is 0, meaning that the volumes can be detached without any time limitations.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(0),
																int64validator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("node_deletion_timeout_seconds"), path.MatchRelative().AtParent().AtName("node_drain_timeout_seconds"), path.MatchRelative().AtParent().AtName("order")),
															},
														},

														"order": schema.StringAttribute{
															Description:         "order defines the order in which Machines are deleted when downscaling. Defaults to 'Random'. Valid values are 'Random, 'Newest', 'Oldest'",
															MarkdownDescription: "order defines the order in which Machines are deleted when downscaling. Defaults to 'Random'. Valid values are 'Random, 'Newest', 'Oldest'",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("Random", "Newest", "Oldest"),
																stringvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("node_deletion_timeout_seconds"), path.MatchRelative().AtParent().AtName("node_drain_timeout_seconds"), path.MatchRelative().AtParent().AtName("node_volume_detach_timeout_seconds")),
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"failure_domain": schema.StringAttribute{
													Description:         "failureDomain is the failure domain the machines will be created in. Must match a key in the FailureDomains map stored on the cluster object.",
													MarkdownDescription: "failureDomain is the failure domain the machines will be created in. Must match a key in the FailureDomains map stored on the cluster object.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
														stringvalidator.LengthAtMost(256),
													},
												},

												"health_check": schema.SingleNestedAttribute{
													Description:         "healthCheck allows to enable, disable and override MachineDeployment health check configuration from the ClusterClass for this MachineDeployment.",
													MarkdownDescription: "healthCheck allows to enable, disable and override MachineDeployment health check configuration from the ClusterClass for this MachineDeployment.",
													Attributes: map[string]schema.Attribute{
														"checks": schema.SingleNestedAttribute{
															Description:         "checks are the checks that are used to evaluate if a Machine is healthy. If one of checks and remediation fields are set, the system assumes that an healthCheck override is defined, and as a consequence the checks and remediation fields from Cluster will be used instead of the corresponding fields in ClusterClass. Independent of this configuration the MachineHealthCheck controller will always flag Machines with 'cluster.x-k8s.io/remediate-machine' annotation and Machines with deleted Nodes as unhealthy. Furthermore, if checks.nodeStartupTimeoutSeconds is not set it is defaulted to 10 minutes and evaluated accordingly.",
															MarkdownDescription: "checks are the checks that are used to evaluate if a Machine is healthy. If one of checks and remediation fields are set, the system assumes that an healthCheck override is defined, and as a consequence the checks and remediation fields from Cluster will be used instead of the corresponding fields in ClusterClass. Independent of this configuration the MachineHealthCheck controller will always flag Machines with 'cluster.x-k8s.io/remediate-machine' annotation and Machines with deleted Nodes as unhealthy. Furthermore, if checks.nodeStartupTimeoutSeconds is not set it is defaulted to 10 minutes and evaluated accordingly.",
															Attributes: map[string]schema.Attribute{
																"node_startup_timeout_seconds": schema.Int64Attribute{
																	Description:         "nodeStartupTimeoutSeconds allows to set the maximum time for MachineHealthCheck to consider a Machine unhealthy if a corresponding Node isn't associated through a 'Spec.ProviderID' field. The duration set in this field is compared to the greatest of: - Cluster's infrastructure ready condition timestamp (if and when available) - Control Plane's initialized condition timestamp (if and when available) - Machine's infrastructure ready condition timestamp (if and when available) - Machine's metadata creation timestamp Defaults to 10 minutes. If you wish to disable this feature, set the value explicitly to 0.",
																	MarkdownDescription: "nodeStartupTimeoutSeconds allows to set the maximum time for MachineHealthCheck to consider a Machine unhealthy if a corresponding Node isn't associated through a 'Spec.ProviderID' field. The duration set in this field is compared to the greatest of: - Cluster's infrastructure ready condition timestamp (if and when available) - Control Plane's initialized condition timestamp (if and when available) - Machine's infrastructure ready condition timestamp (if and when available) - Machine's metadata creation timestamp Defaults to 10 minutes. If you wish to disable this feature, set the value explicitly to 0.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(0),
																		int64validator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("unhealthy_node_conditions")),
																	},
																},

																"unhealthy_node_conditions": schema.ListNestedAttribute{
																	Description:         "unhealthyNodeConditions contains a list of conditions that determine whether a node is considered unhealthy. The conditions are combined in a logical OR, i.e. if any of the conditions is met, the node is unhealthy.",
																	MarkdownDescription: "unhealthyNodeConditions contains a list of conditions that determine whether a node is considered unhealthy. The conditions are combined in a logical OR, i.e. if any of the conditions is met, the node is unhealthy.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"status": schema.StringAttribute{
																				Description:         "status of the condition, one of True, False, Unknown.",
																				MarkdownDescription: "status of the condition, one of True, False, Unknown.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																				Validators: []validator.String{
																					stringvalidator.LengthAtLeast(1),
																				},
																			},

																			"timeout_seconds": schema.Int64Attribute{
																				Description:         "timeoutSeconds is the duration that a node must be in a given status for, after which the node is considered unhealthy. For example, with a value of '1h', the node must match the status for at least 1 hour before being considered unhealthy.",
																				MarkdownDescription: "timeoutSeconds is the duration that a node must be in a given status for, after which the node is considered unhealthy. For example, with a value of '1h', the node must match the status for at least 1 hour before being considered unhealthy.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																				Validators: []validator.Int64{
																					int64validator.AtLeast(0),
																				},
																			},

																			"type": schema.StringAttribute{
																				Description:         "type of Node condition",
																				MarkdownDescription: "type of Node condition",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																				Validators: []validator.String{
																					stringvalidator.LengthAtLeast(1),
																				},
																			},
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																	Validators: []validator.List{
																		listvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("node_startup_timeout_seconds")),
																	},
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
															Validators: []validator.Object{
																objectvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("enabled"), path.MatchRelative().AtParent().AtName("remediation")),
															},
														},

														"enabled": schema.BoolAttribute{
															Description:         "enabled controls if a MachineHealthCheck should be created for the target machines. If false: No MachineHealthCheck will be created. If not set(default): A MachineHealthCheck will be created if it is defined here or in the associated ClusterClass. If no MachineHealthCheck is defined then none will be created. If true: A MachineHealthCheck is guaranteed to be created. Cluster validation will block if 'enable' is true and no MachineHealthCheck definition is available.",
															MarkdownDescription: "enabled controls if a MachineHealthCheck should be created for the target machines. If false: No MachineHealthCheck will be created. If not set(default): A MachineHealthCheck will be created if it is defined here or in the associated ClusterClass. If no MachineHealthCheck is defined then none will be created. If true: A MachineHealthCheck is guaranteed to be created. Cluster validation will block if 'enable' is true and no MachineHealthCheck definition is available.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.Bool{
																boolvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("checks"), path.MatchRelative().AtParent().AtName("remediation")),
															},
														},

														"remediation": schema.SingleNestedAttribute{
															Description:         "remediation configures if and how remediations are triggered if a Machine is unhealthy. If one of checks and remediation fields are set, the system assumes that an healthCheck override is defined, and as a consequence the checks and remediation fields from cluster will be used instead of the corresponding fields in ClusterClass. If an health check override is defined and remediation or remediation.triggerIf is not set, remediation will always be triggered for unhealthy Machines. If an health check override is defined and remediation or remediation.templateRef is not set, the OwnerRemediated condition will be set on unhealthy Machines to trigger remediation via the owner of the Machines, for example a MachineSet or a KubeadmControlPlane.",
															MarkdownDescription: "remediation configures if and how remediations are triggered if a Machine is unhealthy. If one of checks and remediation fields are set, the system assumes that an healthCheck override is defined, and as a consequence the checks and remediation fields from cluster will be used instead of the corresponding fields in ClusterClass. If an health check override is defined and remediation or remediation.triggerIf is not set, remediation will always be triggered for unhealthy Machines. If an health check override is defined and remediation or remediation.templateRef is not set, the OwnerRemediated condition will be set on unhealthy Machines to trigger remediation via the owner of the Machines, for example a MachineSet or a KubeadmControlPlane.",
															Attributes: map[string]schema.Attribute{
																"max_in_flight": schema.StringAttribute{
																	Description:         "maxInFlight determines how many in flight remediations should happen at the same time. Remediation only happens on the MachineSet with the most current revision, while older MachineSets (usually present during rollout operations) aren't allowed to remediate. Note: In general (independent of remediations), unhealthy machines are always prioritized during scale down operations over healthy ones. MaxInFlight can be set to a fixed number or a percentage. Example: when this is set to 20%, the MachineSet controller deletes at most 20% of the desired replicas. If not set, remediation is limited to all machines (bounded by replicas) under the active MachineSet's management.",
																	MarkdownDescription: "maxInFlight determines how many in flight remediations should happen at the same time. Remediation only happens on the MachineSet with the most current revision, while older MachineSets (usually present during rollout operations) aren't allowed to remediate. Note: In general (independent of remediations), unhealthy machines are always prioritized during scale down operations over healthy ones. MaxInFlight can be set to a fixed number or a percentage. Example: when this is set to 20%, the MachineSet controller deletes at most 20% of the desired replicas. If not set, remediation is limited to all machines (bounded by replicas) under the active MachineSet's management.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("template_ref"), path.MatchRelative().AtParent().AtName("trigger_if")),
																	},
																},

																"template_ref": schema.SingleNestedAttribute{
																	Description:         "templateRef is a reference to a remediation template provided by an infrastructure provider. This field is completely optional, when filled, the MachineHealthCheck controller creates a new object from the template referenced and hands off remediation of the machine to a controller that lives outside of Cluster API.",
																	MarkdownDescription: "templateRef is a reference to a remediation template provided by an infrastructure provider. This field is completely optional, when filled, the MachineHealthCheck controller creates a new object from the template referenced and hands off remediation of the machine to a controller that lives outside of Cluster API.",
																	Attributes: map[string]schema.Attribute{
																		"api_version": schema.StringAttribute{
																			Description:         "apiVersion of the remediation template. apiVersion must be fully qualified domain name followed by / and a version. NOTE: This field must be kept in sync with the APIVersion of the remediation template.",
																			MarkdownDescription: "apiVersion of the remediation template. apiVersion must be fully qualified domain name followed by / and a version. NOTE: This field must be kept in sync with the APIVersion of the remediation template.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.LengthAtLeast(1),
																				stringvalidator.LengthAtMost(317),
																				stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*\/[a-z]([-a-z0-9]*[a-z0-9])?$`), ""),
																			},
																		},

																		"kind": schema.StringAttribute{
																			Description:         "kind of the remediation template. kind must consist of alphanumeric characters or '-', start with an alphabetic character, and end with an alphanumeric character.",
																			MarkdownDescription: "kind of the remediation template. kind must consist of alphanumeric characters or '-', start with an alphabetic character, and end with an alphanumeric character.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.LengthAtLeast(1),
																				stringvalidator.LengthAtMost(63),
																				stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z]([-a-zA-Z0-9]*[a-zA-Z0-9])?$`), ""),
																			},
																		},

																		"name": schema.StringAttribute{
																			Description:         "name of the remediation template. name must consist of lower case alphanumeric characters, '-' or '.', and must start and end with an alphanumeric character.",
																			MarkdownDescription: "name of the remediation template. name must consist of lower case alphanumeric characters, '-' or '.', and must start and end with an alphanumeric character.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.LengthAtLeast(1),
																				stringvalidator.LengthAtMost(253),
																				stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
																			},
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																	Validators: []validator.Object{
																		objectvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("max_in_flight"), path.MatchRelative().AtParent().AtName("trigger_if")),
																	},
																},

																"trigger_if": schema.SingleNestedAttribute{
																	Description:         "triggerIf configures if remediations are triggered. If this field is not set, remediations are always triggered.",
																	MarkdownDescription: "triggerIf configures if remediations are triggered. If this field is not set, remediations are always triggered.",
																	Attributes: map[string]schema.Attribute{
																		"unhealthy_in_range": schema.StringAttribute{
																			Description:         "unhealthyInRange specifies that remediations are only triggered if the number of unhealthy Machines is in the configured range. Takes precedence over unhealthyLessThanOrEqualTo. Eg. '[3-5]' - This means that remediation will be allowed only when: (a) there are at least 3 unhealthy Machines (and) (b) there are at most 5 unhealthy Machines",
																			MarkdownDescription: "unhealthyInRange specifies that remediations are only triggered if the number of unhealthy Machines is in the configured range. Takes precedence over unhealthyLessThanOrEqualTo. Eg. '[3-5]' - This means that remediation will be allowed only when: (a) there are at least 3 unhealthy Machines (and) (b) there are at most 5 unhealthy Machines",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.LengthAtLeast(1),
																				stringvalidator.LengthAtMost(32),
																				stringvalidator.RegexMatches(regexp.MustCompile(`^\[[0-9]+-[0-9]+\]$`), ""),
																				stringvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("unhealthy_less_than_or_equal_to")),
																			},
																		},

																		"unhealthy_less_than_or_equal_to": schema.StringAttribute{
																			Description:         "unhealthyLessThanOrEqualTo specifies that remediations are only triggered if the number of unhealthy Machines is less than or equal to the configured value. unhealthyInRange takes precedence if set.",
																			MarkdownDescription: "unhealthyLessThanOrEqualTo specifies that remediations are only triggered if the number of unhealthy Machines is less than or equal to the configured value. unhealthyInRange takes precedence if set.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("unhealthy_in_range")),
																			},
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																	Validators: []validator.Object{
																		objectvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("max_in_flight"), path.MatchRelative().AtParent().AtName("template_ref")),
																	},
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
															Validators: []validator.Object{
																objectvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("checks"), path.MatchRelative().AtParent().AtName("enabled")),
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"metadata": schema.SingleNestedAttribute{
													Description:         "metadata is the metadata applied to the MachineDeployment and the machines of the MachineDeployment. At runtime this metadata is merged with the corresponding metadata from the ClusterClass.",
													MarkdownDescription: "metadata is the metadata applied to the MachineDeployment and the machines of the MachineDeployment. At runtime this metadata is merged with the corresponding metadata from the ClusterClass.",
													Attributes: map[string]schema.Attribute{
														"annotations": schema.MapAttribute{
															Description:         "annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
															MarkdownDescription: "annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.Map{
																mapvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("labels")),
															},
														},

														"labels": schema.MapAttribute{
															Description:         "labels is a map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
															MarkdownDescription: "labels is a map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.Map{
																mapvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("annotations")),
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"min_ready_seconds": schema.Int64Attribute{
													Description:         "minReadySeconds is the minimum number of seconds for which a newly created machine should be ready. Defaults to 0 (machine will be considered available as soon as it is ready)",
													MarkdownDescription: "minReadySeconds is the minimum number of seconds for which a newly created machine should be ready. Defaults to 0 (machine will be considered available as soon as it is ready)",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(0),
													},
												},

												"name": schema.StringAttribute{
													Description:         "name is the unique identifier for this MachineDeploymentTopology. The value is used with other unique identifiers to create a MachineDeployment's Name (e.g. cluster's name, etc). In case the name is greater than the allowed maximum length, the values are hashed together.",
													MarkdownDescription: "name is the unique identifier for this MachineDeploymentTopology. The value is used with other unique identifiers to create a MachineDeployment's Name (e.g. cluster's name, etc). In case the name is greater than the allowed maximum length, the values are hashed together.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
														stringvalidator.LengthAtMost(63),
														stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
													},
												},

												"readiness_gates": schema.ListNestedAttribute{
													Description:         "readinessGates specifies additional conditions to include when evaluating Machine Ready condition. This field can be used e.g. to instruct the machine controller to include in the computation for Machine's ready computation a condition, managed by an external controllers, reporting the status of special software/hardware installed on the Machine. If this field is not defined, readinessGates from the corresponding MachineDeploymentClass will be used, if any.",
													MarkdownDescription: "readinessGates specifies additional conditions to include when evaluating Machine Ready condition. This field can be used e.g. to instruct the machine controller to include in the computation for Machine's ready computation a condition, managed by an external controllers, reporting the status of special software/hardware installed on the Machine. If this field is not defined, readinessGates from the corresponding MachineDeploymentClass will be used, if any.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"condition_type": schema.StringAttribute{
																Description:         "conditionType refers to a condition with matching type in the Machine's condition list. If the conditions doesn't exist, it will be treated as unknown. Note: Both Cluster API conditions or conditions added by 3rd party controllers can be used as readiness gates.",
																MarkdownDescription: "conditionType refers to a condition with matching type in the Machine's condition list. If the conditions doesn't exist, it will be treated as unknown. Note: Both Cluster API conditions or conditions added by 3rd party controllers can be used as readiness gates.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																	stringvalidator.LengthAtMost(316),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*/)?(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])$`), ""),
																},
															},

															"polarity": schema.StringAttribute{
																Description:         "polarity of the conditionType specified in this readinessGate. Valid values are Positive, Negative and omitted. When omitted, the default behaviour will be Positive. A positive polarity means that the condition should report a true status under normal conditions. A negative polarity means that the condition should report a false status under normal conditions.",
																MarkdownDescription: "polarity of the conditionType specified in this readinessGate. Valid values are Positive, Negative and omitted. When omitted, the default behaviour will be Positive. A positive polarity means that the condition should report a true status under normal conditions. A negative polarity means that the condition should report a false status under normal conditions.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("Positive", "Negative"),
																},
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"replicas": schema.Int64Attribute{
													Description:         "replicas is the number of worker nodes belonging to this set. If the value is nil, the MachineDeployment is created without the number of Replicas (defaulting to 1) and it's assumed that an external entity (like cluster autoscaler) is responsible for the management of this value.",
													MarkdownDescription: "replicas is the number of worker nodes belonging to this set. If the value is nil, the MachineDeployment is created without the number of Replicas (defaulting to 1) and it's assumed that an external entity (like cluster autoscaler) is responsible for the management of this value.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"rollout": schema.SingleNestedAttribute{
													Description:         "rollout allows you to configure the behaviour of rolling updates to the MachineDeployment Machines. It allows you to define the strategy used during rolling replacements.",
													MarkdownDescription: "rollout allows you to configure the behaviour of rolling updates to the MachineDeployment Machines. It allows you to define the strategy used during rolling replacements.",
													Attributes: map[string]schema.Attribute{
														"strategy": schema.SingleNestedAttribute{
															Description:         "strategy specifies how to roll out control plane Machines.",
															MarkdownDescription: "strategy specifies how to roll out control plane Machines.",
															Attributes: map[string]schema.Attribute{
																"rolling_update": schema.SingleNestedAttribute{
																	Description:         "rollingUpdate is the rolling update config params. Present only if type = RollingUpdate.",
																	MarkdownDescription: "rollingUpdate is the rolling update config params. Present only if type = RollingUpdate.",
																	Attributes: map[string]schema.Attribute{
																		"max_surge": schema.StringAttribute{
																			Description:         "maxSurge is the maximum number of machines that can be scheduled above the desired number of machines. Value can be an absolute number (ex: 5) or a percentage of desired machines (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up. Defaults to 1. Example: when this is set to 30%, the new MachineSet can be scaled up immediately when the rolling update starts, such that the total number of old and new machines do not exceed 130% of desired machines. Once old machines have been killed, new MachineSet can be scaled up further, ensuring that total number of machines running at any time during the update is at most 130% of desired machines.",
																			MarkdownDescription: "maxSurge is the maximum number of machines that can be scheduled above the desired number of machines. Value can be an absolute number (ex: 5) or a percentage of desired machines (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up. Defaults to 1. Example: when this is set to 30%, the new MachineSet can be scaled up immediately when the rolling update starts, such that the total number of old and new machines do not exceed 130% of desired machines. Once old machines have been killed, new MachineSet can be scaled up further, ensuring that total number of machines running at any time during the update is at most 130% of desired machines.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("max_unavailable")),
																			},
																		},

																		"max_unavailable": schema.StringAttribute{
																			Description:         "maxUnavailable is the maximum number of machines that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired machines (ex: 10%). Absolute number is calculated from percentage by rounding down. This can not be 0 if MaxSurge is 0. Defaults to 0. Example: when this is set to 30%, the old MachineSet can be scaled down to 70% of desired machines immediately when the rolling update starts. Once new machines are ready, old MachineSet can be scaled down further, followed by scaling up the new MachineSet, ensuring that the total number of machines available at all times during the update is at least 70% of desired machines.",
																			MarkdownDescription: "maxUnavailable is the maximum number of machines that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired machines (ex: 10%). Absolute number is calculated from percentage by rounding down. This can not be 0 if MaxSurge is 0. Defaults to 0. Example: when this is set to 30%, the old MachineSet can be scaled down to 70% of desired machines immediately when the rolling update starts. Once new machines are ready, old MachineSet can be scaled down further, followed by scaling up the new MachineSet, ensuring that the total number of machines available at all times during the update is at least 70% of desired machines.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("max_surge")),
																			},
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																	Validators: []validator.Object{
																		objectvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("type")),
																	},
																},

																"type": schema.StringAttribute{
																	Description:         "type of rollout. Allowed values are RollingUpdate and OnDelete. Default is RollingUpdate.",
																	MarkdownDescription: "type of rollout. Allowed values are RollingUpdate and OnDelete. Default is RollingUpdate.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("RollingUpdate", "OnDelete"),
																		stringvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("rolling_update")),
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

												"variables": schema.SingleNestedAttribute{
													Description:         "variables can be used to customize the MachineDeployment through patches.",
													MarkdownDescription: "variables can be used to customize the MachineDeployment through patches.",
													Attributes: map[string]schema.Attribute{
														"overrides": schema.ListNestedAttribute{
															Description:         "overrides can be used to override Cluster level variables.",
															MarkdownDescription: "overrides can be used to override Cluster level variables.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "name of the variable.",
																		MarkdownDescription: "name of the variable.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.LengthAtLeast(1),
																			stringvalidator.LengthAtMost(256),
																		},
																	},

																	"value": schema.MapAttribute{
																		Description:         "value of the variable. Note: the value will be validated against the schema of the corresponding ClusterClassVariable from the ClusterClass. Note: We have to use apiextensionsv1.JSON instead of a custom JSON type, because controller-tools has a hard-coded schema for apiextensionsv1.JSON which cannot be produced by another type via controller-tools, i.e. it is not possible to have no type field. Ref: https://github.com/kubernetes-sigs/controller-tools/blob/d0e03a142d0ecdd5491593e941ee1d6b5d91dba6/pkg/crd/known_types.go#L106-L111",
																		MarkdownDescription: "value of the variable. Note: the value will be validated against the schema of the corresponding ClusterClassVariable from the ClusterClass. Note: We have to use apiextensionsv1.JSON instead of a custom JSON type, because controller-tools has a hard-coded schema for apiextensionsv1.JSON which cannot be produced by another type via controller-tools, i.e. it is not possible to have no type field. Ref: https://github.com/kubernetes-sigs/controller-tools/blob/d0e03a142d0ecdd5491593e941ee1d6b5d91dba6/pkg/crd/known_types.go#L106-L111",
																		ElementType:         types.StringType,
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
										},
										Required: false,
										Optional: true,
										Computed: false,
										Validators: []validator.List{
											listvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("machine_pools")),
										},
									},

									"machine_pools": schema.ListNestedAttribute{
										Description:         "machinePools is a list of machine pools in the cluster.",
										MarkdownDescription: "machinePools is a list of machine pools in the cluster.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"class": schema.StringAttribute{
													Description:         "class is the name of the MachinePoolClass used to create the pool of worker nodes. This should match one of the deployment classes defined in the ClusterClass object mentioned in the 'Cluster.Spec.Class' field.",
													MarkdownDescription: "class is the name of the MachinePoolClass used to create the pool of worker nodes. This should match one of the deployment classes defined in the ClusterClass object mentioned in the 'Cluster.Spec.Class' field.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
														stringvalidator.LengthAtMost(256),
													},
												},

												"deletion": schema.SingleNestedAttribute{
													Description:         "deletion contains configuration options for Machine deletion.",
													MarkdownDescription: "deletion contains configuration options for Machine deletion.",
													Attributes: map[string]schema.Attribute{
														"node_deletion_timeout_seconds": schema.Int64Attribute{
															Description:         "nodeDeletionTimeoutSeconds defines how long the controller will attempt to delete the Node that the MachinePool hosts after the MachinePool is marked for deletion. A duration of 0 will retry deletion indefinitely. Defaults to 10 seconds.",
															MarkdownDescription: "nodeDeletionTimeoutSeconds defines how long the controller will attempt to delete the Node that the MachinePool hosts after the MachinePool is marked for deletion. A duration of 0 will retry deletion indefinitely. Defaults to 10 seconds.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(0),
																int64validator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("node_drain_timeout_seconds"), path.MatchRelative().AtParent().AtName("node_volume_detach_timeout_seconds")),
															},
														},

														"node_drain_timeout_seconds": schema.Int64Attribute{
															Description:         "nodeDrainTimeoutSeconds is the total amount of time that the controller will spend on draining a node. The default value is 0, meaning that the node can be drained without any time limitations. NOTE: nodeDrainTimeoutSeconds is different from 'kubectl drain --timeout'",
															MarkdownDescription: "nodeDrainTimeoutSeconds is the total amount of time that the controller will spend on draining a node. The default value is 0, meaning that the node can be drained without any time limitations. NOTE: nodeDrainTimeoutSeconds is different from 'kubectl drain --timeout'",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(0),
																int64validator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("node_deletion_timeout_seconds"), path.MatchRelative().AtParent().AtName("node_volume_detach_timeout_seconds")),
															},
														},

														"node_volume_detach_timeout_seconds": schema.Int64Attribute{
															Description:         "nodeVolumeDetachTimeoutSeconds is the total amount of time that the controller will spend on waiting for all volumes to be detached. The default value is 0, meaning that the volumes can be detached without any time limitations.",
															MarkdownDescription: "nodeVolumeDetachTimeoutSeconds is the total amount of time that the controller will spend on waiting for all volumes to be detached. The default value is 0, meaning that the volumes can be detached without any time limitations.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(0),
																int64validator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("node_deletion_timeout_seconds"), path.MatchRelative().AtParent().AtName("node_drain_timeout_seconds")),
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"failure_domains": schema.ListAttribute{
													Description:         "failureDomains is the list of failure domains the machine pool will be created in. Must match a key in the FailureDomains map stored on the cluster object.",
													MarkdownDescription: "failureDomains is the list of failure domains the machine pool will be created in. Must match a key in the FailureDomains map stored on the cluster object.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"metadata": schema.SingleNestedAttribute{
													Description:         "metadata is the metadata applied to the MachinePool. At runtime this metadata is merged with the corresponding metadata from the ClusterClass.",
													MarkdownDescription: "metadata is the metadata applied to the MachinePool. At runtime this metadata is merged with the corresponding metadata from the ClusterClass.",
													Attributes: map[string]schema.Attribute{
														"annotations": schema.MapAttribute{
															Description:         "annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
															MarkdownDescription: "annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.Map{
																mapvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("labels")),
															},
														},

														"labels": schema.MapAttribute{
															Description:         "labels is a map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
															MarkdownDescription: "labels is a map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.Map{
																mapvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("annotations")),
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"min_ready_seconds": schema.Int64Attribute{
													Description:         "minReadySeconds is the minimum number of seconds for which a newly created machine pool should be ready. Defaults to 0 (machine will be considered available as soon as it is ready)",
													MarkdownDescription: "minReadySeconds is the minimum number of seconds for which a newly created machine pool should be ready. Defaults to 0 (machine will be considered available as soon as it is ready)",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(0),
													},
												},

												"name": schema.StringAttribute{
													Description:         "name is the unique identifier for this MachinePoolTopology. The value is used with other unique identifiers to create a MachinePool's Name (e.g. cluster's name, etc). In case the name is greater than the allowed maximum length, the values are hashed together.",
													MarkdownDescription: "name is the unique identifier for this MachinePoolTopology. The value is used with other unique identifiers to create a MachinePool's Name (e.g. cluster's name, etc). In case the name is greater than the allowed maximum length, the values are hashed together.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
														stringvalidator.LengthAtMost(63),
														stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
													},
												},

												"replicas": schema.Int64Attribute{
													Description:         "replicas is the number of nodes belonging to this pool. If the value is nil, the MachinePool is created without the number of Replicas (defaulting to 1) and it's assumed that an external entity (like cluster autoscaler) is responsible for the management of this value.",
													MarkdownDescription: "replicas is the number of nodes belonging to this pool. If the value is nil, the MachinePool is created without the number of Replicas (defaulting to 1) and it's assumed that an external entity (like cluster autoscaler) is responsible for the management of this value.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"variables": schema.SingleNestedAttribute{
													Description:         "variables can be used to customize the MachinePool through patches.",
													MarkdownDescription: "variables can be used to customize the MachinePool through patches.",
													Attributes: map[string]schema.Attribute{
														"overrides": schema.ListNestedAttribute{
															Description:         "overrides can be used to override Cluster level variables.",
															MarkdownDescription: "overrides can be used to override Cluster level variables.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "name of the variable.",
																		MarkdownDescription: "name of the variable.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.LengthAtLeast(1),
																			stringvalidator.LengthAtMost(256),
																		},
																	},

																	"value": schema.MapAttribute{
																		Description:         "value of the variable. Note: the value will be validated against the schema of the corresponding ClusterClassVariable from the ClusterClass. Note: We have to use apiextensionsv1.JSON instead of a custom JSON type, because controller-tools has a hard-coded schema for apiextensionsv1.JSON which cannot be produced by another type via controller-tools, i.e. it is not possible to have no type field. Ref: https://github.com/kubernetes-sigs/controller-tools/blob/d0e03a142d0ecdd5491593e941ee1d6b5d91dba6/pkg/crd/known_types.go#L106-L111",
																		MarkdownDescription: "value of the variable. Note: the value will be validated against the schema of the corresponding ClusterClassVariable from the ClusterClass. Note: We have to use apiextensionsv1.JSON instead of a custom JSON type, because controller-tools has a hard-coded schema for apiextensionsv1.JSON which cannot be produced by another type via controller-tools, i.e. it is not possible to have no type field. Ref: https://github.com/kubernetes-sigs/controller-tools/blob/d0e03a142d0ecdd5491593e941ee1d6b5d91dba6/pkg/crd/known_types.go#L106-L111",
																		ElementType:         types.StringType,
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
										},
										Required: false,
										Optional: true,
										Computed: false,
										Validators: []validator.List{
											listvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("machine_deployments")),
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
						Validators: []validator.Object{
							objectvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("availability_gates"), path.MatchRelative().AtParent().AtName("cluster_network"), path.MatchRelative().AtParent().AtName("control_plane_endpoint"), path.MatchRelative().AtParent().AtName("control_plane_ref"), path.MatchRelative().AtParent().AtName("infrastructure_ref"), path.MatchRelative().AtParent().AtName("paused")),
						},
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *ClusterXK8SIoClusterV1Beta2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_cluster_x_k8s_io_cluster_v1beta2_manifest")

	var model ClusterXK8SIoClusterV1Beta2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("cluster.x-k8s.io/v1beta2")
	model.Kind = pointer.String("Cluster")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
