/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package cluster_x_k8s_io_v1beta1

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
	_ datasource.DataSource = &ClusterXK8SIoClusterV1Beta1Manifest{}
)

func NewClusterXK8SIoClusterV1Beta1Manifest() datasource.DataSource {
	return &ClusterXK8SIoClusterV1Beta1Manifest{}
}

type ClusterXK8SIoClusterV1Beta1Manifest struct{}

type ClusterXK8SIoClusterV1Beta1ManifestData struct {
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
			ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
			Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
			Name            *string `tfsdk:"name" json:"name,omitempty"`
			Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
			ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
			Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
		} `tfsdk:"control_plane_ref" json:"controlPlaneRef,omitempty"`
		InfrastructureRef *struct {
			ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
			Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
			Name            *string `tfsdk:"name" json:"name,omitempty"`
			Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
			ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
			Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
		} `tfsdk:"infrastructure_ref" json:"infrastructureRef,omitempty"`
		Paused   *bool `tfsdk:"paused" json:"paused,omitempty"`
		Topology *struct {
			Class          *string `tfsdk:"class" json:"class,omitempty"`
			ClassNamespace *string `tfsdk:"class_namespace" json:"classNamespace,omitempty"`
			ControlPlane   *struct {
				MachineHealthCheck *struct {
					Enable              *bool   `tfsdk:"enable" json:"enable,omitempty"`
					MaxUnhealthy        *string `tfsdk:"max_unhealthy" json:"maxUnhealthy,omitempty"`
					NodeStartupTimeout  *string `tfsdk:"node_startup_timeout" json:"nodeStartupTimeout,omitempty"`
					RemediationTemplate *struct {
						ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
						FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
						Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
						Name            *string `tfsdk:"name" json:"name,omitempty"`
						Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
						ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
						Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
					} `tfsdk:"remediation_template" json:"remediationTemplate,omitempty"`
					UnhealthyConditions *[]struct {
						Status  *string `tfsdk:"status" json:"status,omitempty"`
						Timeout *string `tfsdk:"timeout" json:"timeout,omitempty"`
						Type    *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"unhealthy_conditions" json:"unhealthyConditions,omitempty"`
					UnhealthyRange *string `tfsdk:"unhealthy_range" json:"unhealthyRange,omitempty"`
				} `tfsdk:"machine_health_check" json:"machineHealthCheck,omitempty"`
				Metadata *struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				NodeDeletionTimeout     *string `tfsdk:"node_deletion_timeout" json:"nodeDeletionTimeout,omitempty"`
				NodeDrainTimeout        *string `tfsdk:"node_drain_timeout" json:"nodeDrainTimeout,omitempty"`
				NodeVolumeDetachTimeout *string `tfsdk:"node_volume_detach_timeout" json:"nodeVolumeDetachTimeout,omitempty"`
				ReadinessGates          *[]struct {
					ConditionType *string `tfsdk:"condition_type" json:"conditionType,omitempty"`
					Polarity      *string `tfsdk:"polarity" json:"polarity,omitempty"`
				} `tfsdk:"readiness_gates" json:"readinessGates,omitempty"`
				Replicas  *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
				Variables *struct {
					Overrides *[]struct {
						DefinitionFrom *string            `tfsdk:"definition_from" json:"definitionFrom,omitempty"`
						Name           *string            `tfsdk:"name" json:"name,omitempty"`
						Value          *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"overrides" json:"overrides,omitempty"`
				} `tfsdk:"variables" json:"variables,omitempty"`
			} `tfsdk:"control_plane" json:"controlPlane,omitempty"`
			RolloutAfter *string `tfsdk:"rollout_after" json:"rolloutAfter,omitempty"`
			Variables    *[]struct {
				DefinitionFrom *string            `tfsdk:"definition_from" json:"definitionFrom,omitempty"`
				Name           *string            `tfsdk:"name" json:"name,omitempty"`
				Value          *map[string]string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"variables" json:"variables,omitempty"`
			Version *string `tfsdk:"version" json:"version,omitempty"`
			Workers *struct {
				MachineDeployments *[]struct {
					Class              *string `tfsdk:"class" json:"class,omitempty"`
					FailureDomain      *string `tfsdk:"failure_domain" json:"failureDomain,omitempty"`
					MachineHealthCheck *struct {
						Enable              *bool   `tfsdk:"enable" json:"enable,omitempty"`
						MaxUnhealthy        *string `tfsdk:"max_unhealthy" json:"maxUnhealthy,omitempty"`
						NodeStartupTimeout  *string `tfsdk:"node_startup_timeout" json:"nodeStartupTimeout,omitempty"`
						RemediationTemplate *struct {
							ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
							FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
							Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
							ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
							Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
						} `tfsdk:"remediation_template" json:"remediationTemplate,omitempty"`
						UnhealthyConditions *[]struct {
							Status  *string `tfsdk:"status" json:"status,omitempty"`
							Timeout *string `tfsdk:"timeout" json:"timeout,omitempty"`
							Type    *string `tfsdk:"type" json:"type,omitempty"`
						} `tfsdk:"unhealthy_conditions" json:"unhealthyConditions,omitempty"`
						UnhealthyRange *string `tfsdk:"unhealthy_range" json:"unhealthyRange,omitempty"`
					} `tfsdk:"machine_health_check" json:"machineHealthCheck,omitempty"`
					Metadata *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
					MinReadySeconds         *int64  `tfsdk:"min_ready_seconds" json:"minReadySeconds,omitempty"`
					Name                    *string `tfsdk:"name" json:"name,omitempty"`
					NodeDeletionTimeout     *string `tfsdk:"node_deletion_timeout" json:"nodeDeletionTimeout,omitempty"`
					NodeDrainTimeout        *string `tfsdk:"node_drain_timeout" json:"nodeDrainTimeout,omitempty"`
					NodeVolumeDetachTimeout *string `tfsdk:"node_volume_detach_timeout" json:"nodeVolumeDetachTimeout,omitempty"`
					ReadinessGates          *[]struct {
						ConditionType *string `tfsdk:"condition_type" json:"conditionType,omitempty"`
						Polarity      *string `tfsdk:"polarity" json:"polarity,omitempty"`
					} `tfsdk:"readiness_gates" json:"readinessGates,omitempty"`
					Replicas *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
					Strategy *struct {
						Remediation *struct {
							MaxInFlight *string `tfsdk:"max_in_flight" json:"maxInFlight,omitempty"`
						} `tfsdk:"remediation" json:"remediation,omitempty"`
						RollingUpdate *struct {
							DeletePolicy   *string `tfsdk:"delete_policy" json:"deletePolicy,omitempty"`
							MaxSurge       *string `tfsdk:"max_surge" json:"maxSurge,omitempty"`
							MaxUnavailable *string `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
						} `tfsdk:"rolling_update" json:"rollingUpdate,omitempty"`
						Type *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"strategy" json:"strategy,omitempty"`
					Variables *struct {
						Overrides *[]struct {
							DefinitionFrom *string            `tfsdk:"definition_from" json:"definitionFrom,omitempty"`
							Name           *string            `tfsdk:"name" json:"name,omitempty"`
							Value          *map[string]string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"overrides" json:"overrides,omitempty"`
					} `tfsdk:"variables" json:"variables,omitempty"`
				} `tfsdk:"machine_deployments" json:"machineDeployments,omitempty"`
				MachinePools *[]struct {
					Class          *string   `tfsdk:"class" json:"class,omitempty"`
					FailureDomains *[]string `tfsdk:"failure_domains" json:"failureDomains,omitempty"`
					Metadata       *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
					MinReadySeconds         *int64  `tfsdk:"min_ready_seconds" json:"minReadySeconds,omitempty"`
					Name                    *string `tfsdk:"name" json:"name,omitempty"`
					NodeDeletionTimeout     *string `tfsdk:"node_deletion_timeout" json:"nodeDeletionTimeout,omitempty"`
					NodeDrainTimeout        *string `tfsdk:"node_drain_timeout" json:"nodeDrainTimeout,omitempty"`
					NodeVolumeDetachTimeout *string `tfsdk:"node_volume_detach_timeout" json:"nodeVolumeDetachTimeout,omitempty"`
					Replicas                *int64  `tfsdk:"replicas" json:"replicas,omitempty"`
					Variables               *struct {
						Overrides *[]struct {
							DefinitionFrom *string            `tfsdk:"definition_from" json:"definitionFrom,omitempty"`
							Name           *string            `tfsdk:"name" json:"name,omitempty"`
							Value          *map[string]string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"overrides" json:"overrides,omitempty"`
					} `tfsdk:"variables" json:"variables,omitempty"`
				} `tfsdk:"machine_pools" json:"machinePools,omitempty"`
			} `tfsdk:"workers" json:"workers,omitempty"`
		} `tfsdk:"topology" json:"topology,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ClusterXK8SIoClusterV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_cluster_x_k8s_io_cluster_v1beta1_manifest"
}

func (r *ClusterXK8SIoClusterV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
						Description:         "availabilityGates specifies additional conditions to include when evaluating Cluster Available condition. If this field is not defined and the Cluster implements a managed topology, availabilityGates from the corresponding ClusterClass will be used, if any. NOTE: this field is considered only for computing v1beta2 conditions.",
						MarkdownDescription: "availabilityGates specifies additional conditions to include when evaluating Cluster Available condition. If this field is not defined and the Cluster implements a managed topology, availabilityGates from the corresponding ClusterClass will be used, if any. NOTE: this field is considered only for computing v1beta2 conditions.",
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
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
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
									stringvalidator.LengthAtMost(512),
								},
							},

							"port": schema.Int64Attribute{
								Description:         "port is the port on which the API server is serving.",
								MarkdownDescription: "port is the port on which the API server is serving.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"control_plane_ref": schema.SingleNestedAttribute{
						Description:         "controlPlaneRef is an optional reference to a provider-specific resource that holds the details for provisioning the Control Plane for a Cluster.",
						MarkdownDescription: "controlPlaneRef is an optional reference to a provider-specific resource that holds the details for provisioning the Control Plane for a Cluster.",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "API version of the referent.",
								MarkdownDescription: "API version of the referent.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"field_path": schema.StringAttribute{
								Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.",
								MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								MarkdownDescription: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
								MarkdownDescription: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resource_version": schema.StringAttribute{
								Description:         "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
								MarkdownDescription: "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"uid": schema.StringAttribute{
								Description:         "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
								MarkdownDescription: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"infrastructure_ref": schema.SingleNestedAttribute{
						Description:         "infrastructureRef is a reference to a provider-specific resource that holds the details for provisioning infrastructure for a cluster in said provider.",
						MarkdownDescription: "infrastructureRef is a reference to a provider-specific resource that holds the details for provisioning infrastructure for a cluster in said provider.",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "API version of the referent.",
								MarkdownDescription: "API version of the referent.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"field_path": schema.StringAttribute{
								Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.",
								MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								MarkdownDescription: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
								MarkdownDescription: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resource_version": schema.StringAttribute{
								Description:         "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
								MarkdownDescription: "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"uid": schema.StringAttribute{
								Description:         "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
								MarkdownDescription: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"paused": schema.BoolAttribute{
						Description:         "paused can be used to prevent controllers from processing the Cluster and all its associated objects.",
						MarkdownDescription: "paused can be used to prevent controllers from processing the Cluster and all its associated objects.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"topology": schema.SingleNestedAttribute{
						Description:         "topology encapsulates the topology for the cluster. NOTE: It is required to enable the ClusterTopology feature gate flag to activate managed topologies support; this feature is highly experimental, and parts of it might still be not implemented.",
						MarkdownDescription: "topology encapsulates the topology for the cluster. NOTE: It is required to enable the ClusterTopology feature gate flag to activate managed topologies support; this feature is highly experimental, and parts of it might still be not implemented.",
						Attributes: map[string]schema.Attribute{
							"class": schema.StringAttribute{
								Description:         "class is the name of the ClusterClass object to create the topology.",
								MarkdownDescription: "class is the name of the ClusterClass object to create the topology.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(253),
								},
							},

							"class_namespace": schema.StringAttribute{
								Description:         "classNamespace is the namespace of the ClusterClass that should be used for the topology. If classNamespace is empty or not set, it is defaulted to the namespace of the Cluster object. classNamespace must be a valid namespace name and because of that be at most 63 characters in length and it must consist only of lower case alphanumeric characters or hyphens (-), and must start and end with an alphanumeric character.",
								MarkdownDescription: "classNamespace is the namespace of the ClusterClass that should be used for the topology. If classNamespace is empty or not set, it is defaulted to the namespace of the Cluster object. classNamespace must be a valid namespace name and because of that be at most 63 characters in length and it must consist only of lower case alphanumeric characters or hyphens (-), and must start and end with an alphanumeric character.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(63),
									stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
								},
							},

							"control_plane": schema.SingleNestedAttribute{
								Description:         "controlPlane describes the cluster control plane.",
								MarkdownDescription: "controlPlane describes the cluster control plane.",
								Attributes: map[string]schema.Attribute{
									"machine_health_check": schema.SingleNestedAttribute{
										Description:         "machineHealthCheck allows to enable, disable and override the MachineHealthCheck configuration in the ClusterClass for this control plane.",
										MarkdownDescription: "machineHealthCheck allows to enable, disable and override the MachineHealthCheck configuration in the ClusterClass for this control plane.",
										Attributes: map[string]schema.Attribute{
											"enable": schema.BoolAttribute{
												Description:         "enable controls if a MachineHealthCheck should be created for the target machines. If false: No MachineHealthCheck will be created. If not set(default): A MachineHealthCheck will be created if it is defined here or in the associated ClusterClass. If no MachineHealthCheck is defined then none will be created. If true: A MachineHealthCheck is guaranteed to be created. Cluster validation will block if 'enable' is true and no MachineHealthCheck definition is available.",
												MarkdownDescription: "enable controls if a MachineHealthCheck should be created for the target machines. If false: No MachineHealthCheck will be created. If not set(default): A MachineHealthCheck will be created if it is defined here or in the associated ClusterClass. If no MachineHealthCheck is defined then none will be created. If true: A MachineHealthCheck is guaranteed to be created. Cluster validation will block if 'enable' is true and no MachineHealthCheck definition is available.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_unhealthy": schema.StringAttribute{
												Description:         "maxUnhealthy specifies the maximum number of unhealthy machines allowed. Any further remediation is only allowed if at most 'maxUnhealthy' machines selected by 'selector' are not healthy.",
												MarkdownDescription: "maxUnhealthy specifies the maximum number of unhealthy machines allowed. Any further remediation is only allowed if at most 'maxUnhealthy' machines selected by 'selector' are not healthy.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"node_startup_timeout": schema.StringAttribute{
												Description:         "nodeStartupTimeout allows to set the maximum time for MachineHealthCheck to consider a Machine unhealthy if a corresponding Node isn't associated through a 'Spec.ProviderID' field. The duration set in this field is compared to the greatest of: - Cluster's infrastructure ready condition timestamp (if and when available) - Control Plane's initialized condition timestamp (if and when available) - Machine's infrastructure ready condition timestamp (if and when available) - Machine's metadata creation timestamp Defaults to 10 minutes. If you wish to disable this feature, set the value explicitly to 0.",
												MarkdownDescription: "nodeStartupTimeout allows to set the maximum time for MachineHealthCheck to consider a Machine unhealthy if a corresponding Node isn't associated through a 'Spec.ProviderID' field. The duration set in this field is compared to the greatest of: - Cluster's infrastructure ready condition timestamp (if and when available) - Control Plane's initialized condition timestamp (if and when available) - Machine's infrastructure ready condition timestamp (if and when available) - Machine's metadata creation timestamp Defaults to 10 minutes. If you wish to disable this feature, set the value explicitly to 0.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"remediation_template": schema.SingleNestedAttribute{
												Description:         "remediationTemplate is a reference to a remediation template provided by an infrastructure provider. This field is completely optional, when filled, the MachineHealthCheck controller creates a new object from the template referenced and hands off remediation of the machine to a controller that lives outside of Cluster API.",
												MarkdownDescription: "remediationTemplate is a reference to a remediation template provided by an infrastructure provider. This field is completely optional, when filled, the MachineHealthCheck controller creates a new object from the template referenced and hands off remediation of the machine to a controller that lives outside of Cluster API.",
												Attributes: map[string]schema.Attribute{
													"api_version": schema.StringAttribute{
														Description:         "API version of the referent.",
														MarkdownDescription: "API version of the referent.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"field_path": schema.StringAttribute{
														Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.",
														MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"kind": schema.StringAttribute{
														Description:         "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
														MarkdownDescription: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
														MarkdownDescription: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"resource_version": schema.StringAttribute{
														Description:         "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
														MarkdownDescription: "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"uid": schema.StringAttribute{
														Description:         "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
														MarkdownDescription: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"unhealthy_conditions": schema.ListNestedAttribute{
												Description:         "unhealthyConditions contains a list of the conditions that determine whether a node is considered unhealthy. The conditions are combined in a logical OR, i.e. if any of the conditions is met, the node is unhealthy.",
												MarkdownDescription: "unhealthyConditions contains a list of the conditions that determine whether a node is considered unhealthy. The conditions are combined in a logical OR, i.e. if any of the conditions is met, the node is unhealthy.",
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

														"timeout": schema.StringAttribute{
															Description:         "timeout is the duration that a node must be in a given status for, after which the node is considered unhealthy. For example, with a value of '1h', the node must match the status for at least 1 hour before being considered unhealthy.",
															MarkdownDescription: "timeout is the duration that a node must be in a given status for, after which the node is considered unhealthy. For example, with a value of '1h', the node must match the status for at least 1 hour before being considered unhealthy.",
															Required:            true,
															Optional:            false,
															Computed:            false,
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
											},

											"unhealthy_range": schema.StringAttribute{
												Description:         "unhealthyRange specifies the range of unhealthy machines allowed. Any further remediation is only allowed if the number of machines selected by 'selector' as not healthy is within the range of 'unhealthyRange'. Takes precedence over maxUnhealthy. Eg. '[3-5]' - This means that remediation will be allowed only when: (a) there are at least 3 unhealthy machines (and) (b) there are at most 5 unhealthy machines",
												MarkdownDescription: "unhealthyRange specifies the range of unhealthy machines allowed. Any further remediation is only allowed if the number of machines selected by 'selector' as not healthy is within the range of 'unhealthyRange'. Takes precedence over maxUnhealthy. Eg. '[3-5]' - This means that remediation will be allowed only when: (a) there are at least 3 unhealthy machines (and) (b) there are at most 5 unhealthy machines",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
													stringvalidator.LengthAtMost(32),
													stringvalidator.RegexMatches(regexp.MustCompile(`^\[[0-9]+-[0-9]+\]$`), ""),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
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
											},

											"labels": schema.MapAttribute{
												Description:         "labels is a map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
												MarkdownDescription: "labels is a map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
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

									"node_deletion_timeout": schema.StringAttribute{
										Description:         "nodeDeletionTimeout defines how long the controller will attempt to delete the Node that the Machine hosts after the Machine is marked for deletion. A duration of 0 will retry deletion indefinitely. Defaults to 10 seconds.",
										MarkdownDescription: "nodeDeletionTimeout defines how long the controller will attempt to delete the Node that the Machine hosts after the Machine is marked for deletion. A duration of 0 will retry deletion indefinitely. Defaults to 10 seconds.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"node_drain_timeout": schema.StringAttribute{
										Description:         "nodeDrainTimeout is the total amount of time that the controller will spend on draining a node. The default value is 0, meaning that the node can be drained without any time limitations. NOTE: NodeDrainTimeout is different from 'kubectl drain --timeout'",
										MarkdownDescription: "nodeDrainTimeout is the total amount of time that the controller will spend on draining a node. The default value is 0, meaning that the node can be drained without any time limitations. NOTE: NodeDrainTimeout is different from 'kubectl drain --timeout'",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"node_volume_detach_timeout": schema.StringAttribute{
										Description:         "nodeVolumeDetachTimeout is the total amount of time that the controller will spend on waiting for all volumes to be detached. The default value is 0, meaning that the volumes can be detached without any time limitations.",
										MarkdownDescription: "nodeVolumeDetachTimeout is the total amount of time that the controller will spend on waiting for all volumes to be detached. The default value is 0, meaning that the volumes can be detached without any time limitations.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"readiness_gates": schema.ListNestedAttribute{
										Description:         "readinessGates specifies additional conditions to include when evaluating Machine Ready condition. This field can be used e.g. to instruct the machine controller to include in the computation for Machine's ready computation a condition, managed by an external controllers, reporting the status of special software/hardware installed on the Machine. If this field is not defined, readinessGates from the corresponding ControlPlaneClass will be used, if any. NOTE: This field is considered only for computing v1beta2 conditions. NOTE: Specific control plane provider implementations might automatically extend the list of readinessGates; e.g. the kubeadm control provider adds ReadinessGates for the APIServerPodHealthy, SchedulerPodHealthy conditions, etc.",
										MarkdownDescription: "readinessGates specifies additional conditions to include when evaluating Machine Ready condition. This field can be used e.g. to instruct the machine controller to include in the computation for Machine's ready computation a condition, managed by an external controllers, reporting the status of special software/hardware installed on the Machine. If this field is not defined, readinessGates from the corresponding ControlPlaneClass will be used, if any. NOTE: This field is considered only for computing v1beta2 conditions. NOTE: Specific control plane provider implementations might automatically extend the list of readinessGates; e.g. the kubeadm control provider adds ReadinessGates for the APIServerPodHealthy, SchedulerPodHealthy conditions, etc.",
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
										Description:         "replicas is the number of control plane nodes. If the value is nil, the ControlPlane object is created without the number of Replicas and it's assumed that the control plane controller does not implement support for this field. When specified against a control plane provider that lacks support for this field, this value will be ignored.",
										MarkdownDescription: "replicas is the number of control plane nodes. If the value is nil, the ControlPlane object is created without the number of Replicas and it's assumed that the control plane controller does not implement support for this field. When specified against a control plane provider that lacks support for this field, this value will be ignored.",
										Required:            false,
										Optional:            true,
										Computed:            false,
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
														"definition_from": schema.StringAttribute{
															Description:         "definitionFrom specifies where the definition of this Variable is from. Deprecated: This field is deprecated, must not be set anymore and is going to be removed in the next apiVersion.",
															MarkdownDescription: "definitionFrom specifies where the definition of this Variable is from. Deprecated: This field is deprecated, must not be set anymore and is going to be removed in the next apiVersion.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtMost(256),
															},
														},

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
								Required: false,
								Optional: true,
								Computed: false,
							},

							"rollout_after": schema.StringAttribute{
								Description:         "rolloutAfter performs a rollout of the entire cluster one component at a time, control plane first and then machine deployments. Deprecated: This field has no function and is going to be removed in the next apiVersion.",
								MarkdownDescription: "rolloutAfter performs a rollout of the entire cluster one component at a time, control plane first and then machine deployments. Deprecated: This field has no function and is going to be removed in the next apiVersion.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									validators.DateTime64Validator(),
								},
							},

							"variables": schema.ListNestedAttribute{
								Description:         "variables can be used to customize the Cluster through patches. They must comply to the corresponding VariableClasses defined in the ClusterClass.",
								MarkdownDescription: "variables can be used to customize the Cluster through patches. They must comply to the corresponding VariableClasses defined in the ClusterClass.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"definition_from": schema.StringAttribute{
											Description:         "definitionFrom specifies where the definition of this Variable is from. Deprecated: This field is deprecated, must not be set anymore and is going to be removed in the next apiVersion.",
											MarkdownDescription: "definitionFrom specifies where the definition of this Variable is from. Deprecated: This field is deprecated, must not be set anymore and is going to be removed in the next apiVersion.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtMost(256),
											},
										},

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

												"machine_health_check": schema.SingleNestedAttribute{
													Description:         "machineHealthCheck allows to enable, disable and override the MachineHealthCheck configuration in the ClusterClass for this MachineDeployment.",
													MarkdownDescription: "machineHealthCheck allows to enable, disable and override the MachineHealthCheck configuration in the ClusterClass for this MachineDeployment.",
													Attributes: map[string]schema.Attribute{
														"enable": schema.BoolAttribute{
															Description:         "enable controls if a MachineHealthCheck should be created for the target machines. If false: No MachineHealthCheck will be created. If not set(default): A MachineHealthCheck will be created if it is defined here or in the associated ClusterClass. If no MachineHealthCheck is defined then none will be created. If true: A MachineHealthCheck is guaranteed to be created. Cluster validation will block if 'enable' is true and no MachineHealthCheck definition is available.",
															MarkdownDescription: "enable controls if a MachineHealthCheck should be created for the target machines. If false: No MachineHealthCheck will be created. If not set(default): A MachineHealthCheck will be created if it is defined here or in the associated ClusterClass. If no MachineHealthCheck is defined then none will be created. If true: A MachineHealthCheck is guaranteed to be created. Cluster validation will block if 'enable' is true and no MachineHealthCheck definition is available.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"max_unhealthy": schema.StringAttribute{
															Description:         "maxUnhealthy specifies the maximum number of unhealthy machines allowed. Any further remediation is only allowed if at most 'maxUnhealthy' machines selected by 'selector' are not healthy.",
															MarkdownDescription: "maxUnhealthy specifies the maximum number of unhealthy machines allowed. Any further remediation is only allowed if at most 'maxUnhealthy' machines selected by 'selector' are not healthy.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"node_startup_timeout": schema.StringAttribute{
															Description:         "nodeStartupTimeout allows to set the maximum time for MachineHealthCheck to consider a Machine unhealthy if a corresponding Node isn't associated through a 'Spec.ProviderID' field. The duration set in this field is compared to the greatest of: - Cluster's infrastructure ready condition timestamp (if and when available) - Control Plane's initialized condition timestamp (if and when available) - Machine's infrastructure ready condition timestamp (if and when available) - Machine's metadata creation timestamp Defaults to 10 minutes. If you wish to disable this feature, set the value explicitly to 0.",
															MarkdownDescription: "nodeStartupTimeout allows to set the maximum time for MachineHealthCheck to consider a Machine unhealthy if a corresponding Node isn't associated through a 'Spec.ProviderID' field. The duration set in this field is compared to the greatest of: - Cluster's infrastructure ready condition timestamp (if and when available) - Control Plane's initialized condition timestamp (if and when available) - Machine's infrastructure ready condition timestamp (if and when available) - Machine's metadata creation timestamp Defaults to 10 minutes. If you wish to disable this feature, set the value explicitly to 0.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"remediation_template": schema.SingleNestedAttribute{
															Description:         "remediationTemplate is a reference to a remediation template provided by an infrastructure provider. This field is completely optional, when filled, the MachineHealthCheck controller creates a new object from the template referenced and hands off remediation of the machine to a controller that lives outside of Cluster API.",
															MarkdownDescription: "remediationTemplate is a reference to a remediation template provided by an infrastructure provider. This field is completely optional, when filled, the MachineHealthCheck controller creates a new object from the template referenced and hands off remediation of the machine to a controller that lives outside of Cluster API.",
															Attributes: map[string]schema.Attribute{
																"api_version": schema.StringAttribute{
																	Description:         "API version of the referent.",
																	MarkdownDescription: "API version of the referent.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"field_path": schema.StringAttribute{
																	Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.",
																	MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"kind": schema.StringAttribute{
																	Description:         "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
																	MarkdownDescription: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"namespace": schema.StringAttribute{
																	Description:         "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
																	MarkdownDescription: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"resource_version": schema.StringAttribute{
																	Description:         "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
																	MarkdownDescription: "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"uid": schema.StringAttribute{
																	Description:         "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
																	MarkdownDescription: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"unhealthy_conditions": schema.ListNestedAttribute{
															Description:         "unhealthyConditions contains a list of the conditions that determine whether a node is considered unhealthy. The conditions are combined in a logical OR, i.e. if any of the conditions is met, the node is unhealthy.",
															MarkdownDescription: "unhealthyConditions contains a list of the conditions that determine whether a node is considered unhealthy. The conditions are combined in a logical OR, i.e. if any of the conditions is met, the node is unhealthy.",
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

																	"timeout": schema.StringAttribute{
																		Description:         "timeout is the duration that a node must be in a given status for, after which the node is considered unhealthy. For example, with a value of '1h', the node must match the status for at least 1 hour before being considered unhealthy.",
																		MarkdownDescription: "timeout is the duration that a node must be in a given status for, after which the node is considered unhealthy. For example, with a value of '1h', the node must match the status for at least 1 hour before being considered unhealthy.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
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
														},

														"unhealthy_range": schema.StringAttribute{
															Description:         "unhealthyRange specifies the range of unhealthy machines allowed. Any further remediation is only allowed if the number of machines selected by 'selector' as not healthy is within the range of 'unhealthyRange'. Takes precedence over maxUnhealthy. Eg. '[3-5]' - This means that remediation will be allowed only when: (a) there are at least 3 unhealthy machines (and) (b) there are at most 5 unhealthy machines",
															MarkdownDescription: "unhealthyRange specifies the range of unhealthy machines allowed. Any further remediation is only allowed if the number of machines selected by 'selector' as not healthy is within the range of 'unhealthyRange'. Takes precedence over maxUnhealthy. Eg. '[3-5]' - This means that remediation will be allowed only when: (a) there are at least 3 unhealthy machines (and) (b) there are at most 5 unhealthy machines",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtLeast(1),
																stringvalidator.LengthAtMost(32),
																stringvalidator.RegexMatches(regexp.MustCompile(`^\[[0-9]+-[0-9]+\]$`), ""),
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
														},

														"labels": schema.MapAttribute{
															Description:         "labels is a map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
															MarkdownDescription: "labels is a map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
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

												"min_ready_seconds": schema.Int64Attribute{
													Description:         "minReadySeconds is the minimum number of seconds for which a newly created machine should be ready. Defaults to 0 (machine will be considered available as soon as it is ready)",
													MarkdownDescription: "minReadySeconds is the minimum number of seconds for which a newly created machine should be ready. Defaults to 0 (machine will be considered available as soon as it is ready)",
													Required:            false,
													Optional:            true,
													Computed:            false,
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
													},
												},

												"node_deletion_timeout": schema.StringAttribute{
													Description:         "nodeDeletionTimeout defines how long the controller will attempt to delete the Node that the Machine hosts after the Machine is marked for deletion. A duration of 0 will retry deletion indefinitely. Defaults to 10 seconds.",
													MarkdownDescription: "nodeDeletionTimeout defines how long the controller will attempt to delete the Node that the Machine hosts after the Machine is marked for deletion. A duration of 0 will retry deletion indefinitely. Defaults to 10 seconds.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"node_drain_timeout": schema.StringAttribute{
													Description:         "nodeDrainTimeout is the total amount of time that the controller will spend on draining a node. The default value is 0, meaning that the node can be drained without any time limitations. NOTE: NodeDrainTimeout is different from 'kubectl drain --timeout'",
													MarkdownDescription: "nodeDrainTimeout is the total amount of time that the controller will spend on draining a node. The default value is 0, meaning that the node can be drained without any time limitations. NOTE: NodeDrainTimeout is different from 'kubectl drain --timeout'",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"node_volume_detach_timeout": schema.StringAttribute{
													Description:         "nodeVolumeDetachTimeout is the total amount of time that the controller will spend on waiting for all volumes to be detached. The default value is 0, meaning that the volumes can be detached without any time limitations.",
													MarkdownDescription: "nodeVolumeDetachTimeout is the total amount of time that the controller will spend on waiting for all volumes to be detached. The default value is 0, meaning that the volumes can be detached without any time limitations.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"readiness_gates": schema.ListNestedAttribute{
													Description:         "readinessGates specifies additional conditions to include when evaluating Machine Ready condition. This field can be used e.g. to instruct the machine controller to include in the computation for Machine's ready computation a condition, managed by an external controllers, reporting the status of special software/hardware installed on the Machine. If this field is not defined, readinessGates from the corresponding MachineDeploymentClass will be used, if any. NOTE: This field is considered only for computing v1beta2 conditions.",
													MarkdownDescription: "readinessGates specifies additional conditions to include when evaluating Machine Ready condition. This field can be used e.g. to instruct the machine controller to include in the computation for Machine's ready computation a condition, managed by an external controllers, reporting the status of special software/hardware installed on the Machine. If this field is not defined, readinessGates from the corresponding MachineDeploymentClass will be used, if any. NOTE: This field is considered only for computing v1beta2 conditions.",
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

												"strategy": schema.SingleNestedAttribute{
													Description:         "strategy is the deployment strategy to use to replace existing machines with new ones.",
													MarkdownDescription: "strategy is the deployment strategy to use to replace existing machines with new ones.",
													Attributes: map[string]schema.Attribute{
														"remediation": schema.SingleNestedAttribute{
															Description:         "remediation controls the strategy of remediating unhealthy machines and how remediating operations should occur during the lifecycle of the dependant MachineSets.",
															MarkdownDescription: "remediation controls the strategy of remediating unhealthy machines and how remediating operations should occur during the lifecycle of the dependant MachineSets.",
															Attributes: map[string]schema.Attribute{
																"max_in_flight": schema.StringAttribute{
																	Description:         "maxInFlight determines how many in flight remediations should happen at the same time. Remediation only happens on the MachineSet with the most current revision, while older MachineSets (usually present during rollout operations) aren't allowed to remediate. Note: In general (independent of remediations), unhealthy machines are always prioritized during scale down operations over healthy ones. MaxInFlight can be set to a fixed number or a percentage. Example: when this is set to 20%, the MachineSet controller deletes at most 20% of the desired replicas. If not set, remediation is limited to all machines (bounded by replicas) under the active MachineSet's management.",
																	MarkdownDescription: "maxInFlight determines how many in flight remediations should happen at the same time. Remediation only happens on the MachineSet with the most current revision, while older MachineSets (usually present during rollout operations) aren't allowed to remediate. Note: In general (independent of remediations), unhealthy machines are always prioritized during scale down operations over healthy ones. MaxInFlight can be set to a fixed number or a percentage. Example: when this is set to 20%, the MachineSet controller deletes at most 20% of the desired replicas. If not set, remediation is limited to all machines (bounded by replicas) under the active MachineSet's management.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"rolling_update": schema.SingleNestedAttribute{
															Description:         "rollingUpdate is the rolling update config params. Present only if MachineDeploymentStrategyType = RollingUpdate.",
															MarkdownDescription: "rollingUpdate is the rolling update config params. Present only if MachineDeploymentStrategyType = RollingUpdate.",
															Attributes: map[string]schema.Attribute{
																"delete_policy": schema.StringAttribute{
																	Description:         "deletePolicy defines the policy used by the MachineDeployment to identify nodes to delete when downscaling. Valid values are 'Random, 'Newest', 'Oldest' When no value is supplied, the default DeletePolicy of MachineSet is used",
																	MarkdownDescription: "deletePolicy defines the policy used by the MachineDeployment to identify nodes to delete when downscaling. Valid values are 'Random, 'Newest', 'Oldest' When no value is supplied, the default DeletePolicy of MachineSet is used",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("Random", "Newest", "Oldest"),
																	},
																},

																"max_surge": schema.StringAttribute{
																	Description:         "maxSurge is the maximum number of machines that can be scheduled above the desired number of machines. Value can be an absolute number (ex: 5) or a percentage of desired machines (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up. Defaults to 1. Example: when this is set to 30%, the new MachineSet can be scaled up immediately when the rolling update starts, such that the total number of old and new machines do not exceed 130% of desired machines. Once old machines have been killed, new MachineSet can be scaled up further, ensuring that total number of machines running at any time during the update is at most 130% of desired machines.",
																	MarkdownDescription: "maxSurge is the maximum number of machines that can be scheduled above the desired number of machines. Value can be an absolute number (ex: 5) or a percentage of desired machines (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up. Defaults to 1. Example: when this is set to 30%, the new MachineSet can be scaled up immediately when the rolling update starts, such that the total number of old and new machines do not exceed 130% of desired machines. Once old machines have been killed, new MachineSet can be scaled up further, ensuring that total number of machines running at any time during the update is at most 130% of desired machines.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"max_unavailable": schema.StringAttribute{
																	Description:         "maxUnavailable is the maximum number of machines that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired machines (ex: 10%). Absolute number is calculated from percentage by rounding down. This can not be 0 if MaxSurge is 0. Defaults to 0. Example: when this is set to 30%, the old MachineSet can be scaled down to 70% of desired machines immediately when the rolling update starts. Once new machines are ready, old MachineSet can be scaled down further, followed by scaling up the new MachineSet, ensuring that the total number of machines available at all times during the update is at least 70% of desired machines.",
																	MarkdownDescription: "maxUnavailable is the maximum number of machines that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired machines (ex: 10%). Absolute number is calculated from percentage by rounding down. This can not be 0 if MaxSurge is 0. Defaults to 0. Example: when this is set to 30%, the old MachineSet can be scaled down to 70% of desired machines immediately when the rolling update starts. Once new machines are ready, old MachineSet can be scaled down further, followed by scaling up the new MachineSet, ensuring that the total number of machines available at all times during the update is at least 70% of desired machines.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"type": schema.StringAttribute{
															Description:         "type of deployment. Allowed values are RollingUpdate and OnDelete. The default is RollingUpdate.",
															MarkdownDescription: "type of deployment. Allowed values are RollingUpdate and OnDelete. The default is RollingUpdate.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("RollingUpdate", "OnDelete"),
															},
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
																	"definition_from": schema.StringAttribute{
																		Description:         "definitionFrom specifies where the definition of this Variable is from. Deprecated: This field is deprecated, must not be set anymore and is going to be removed in the next apiVersion.",
																		MarkdownDescription: "definitionFrom specifies where the definition of this Variable is from. Deprecated: This field is deprecated, must not be set anymore and is going to be removed in the next apiVersion.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.LengthAtMost(256),
																		},
																	},

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
														},

														"labels": schema.MapAttribute{
															Description:         "labels is a map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
															MarkdownDescription: "labels is a map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
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

												"min_ready_seconds": schema.Int64Attribute{
													Description:         "minReadySeconds is the minimum number of seconds for which a newly created machine pool should be ready. Defaults to 0 (machine will be considered available as soon as it is ready)",
													MarkdownDescription: "minReadySeconds is the minimum number of seconds for which a newly created machine pool should be ready. Defaults to 0 (machine will be considered available as soon as it is ready)",
													Required:            false,
													Optional:            true,
													Computed:            false,
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
													},
												},

												"node_deletion_timeout": schema.StringAttribute{
													Description:         "nodeDeletionTimeout defines how long the controller will attempt to delete the Node that the MachinePool hosts after the MachinePool is marked for deletion. A duration of 0 will retry deletion indefinitely. Defaults to 10 seconds.",
													MarkdownDescription: "nodeDeletionTimeout defines how long the controller will attempt to delete the Node that the MachinePool hosts after the MachinePool is marked for deletion. A duration of 0 will retry deletion indefinitely. Defaults to 10 seconds.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"node_drain_timeout": schema.StringAttribute{
													Description:         "nodeDrainTimeout is the total amount of time that the controller will spend on draining a node. The default value is 0, meaning that the node can be drained without any time limitations. NOTE: NodeDrainTimeout is different from 'kubectl drain --timeout'",
													MarkdownDescription: "nodeDrainTimeout is the total amount of time that the controller will spend on draining a node. The default value is 0, meaning that the node can be drained without any time limitations. NOTE: NodeDrainTimeout is different from 'kubectl drain --timeout'",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"node_volume_detach_timeout": schema.StringAttribute{
													Description:         "nodeVolumeDetachTimeout is the total amount of time that the controller will spend on waiting for all volumes to be detached. The default value is 0, meaning that the volumes can be detached without any time limitations.",
													MarkdownDescription: "nodeVolumeDetachTimeout is the total amount of time that the controller will spend on waiting for all volumes to be detached. The default value is 0, meaning that the volumes can be detached without any time limitations.",
													Required:            false,
													Optional:            true,
													Computed:            false,
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
																	"definition_from": schema.StringAttribute{
																		Description:         "definitionFrom specifies where the definition of this Variable is from. Deprecated: This field is deprecated, must not be set anymore and is going to be removed in the next apiVersion.",
																		MarkdownDescription: "definitionFrom specifies where the definition of this Variable is from. Deprecated: This field is deprecated, must not be set anymore and is going to be removed in the next apiVersion.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.LengthAtMost(256),
																		},
																	},

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
	}
}

func (r *ClusterXK8SIoClusterV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_cluster_x_k8s_io_cluster_v1beta1_manifest")

	var model ClusterXK8SIoClusterV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("cluster.x-k8s.io/v1beta1")
	model.Kind = pointer.String("Cluster")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
