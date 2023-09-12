/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package cluster_x_k8s_io_v1beta1

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
	_ datasource.DataSource              = &ClusterXK8SIoClusterV1Beta1DataSource{}
	_ datasource.DataSourceWithConfigure = &ClusterXK8SIoClusterV1Beta1DataSource{}
)

func NewClusterXK8SIoClusterV1Beta1DataSource() datasource.DataSource {
	return &ClusterXK8SIoClusterV1Beta1DataSource{}
}

type ClusterXK8SIoClusterV1Beta1DataSource struct {
	kubernetesClient dynamic.Interface
}

type ClusterXK8SIoClusterV1Beta1DataSourceData struct {
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
			Class        *string `tfsdk:"class" json:"class,omitempty"`
			ControlPlane *struct {
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
				Replicas                *int64  `tfsdk:"replicas" json:"replicas,omitempty"`
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
					Replicas                *int64  `tfsdk:"replicas" json:"replicas,omitempty"`
					Strategy                *struct {
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

func (r *ClusterXK8SIoClusterV1Beta1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_cluster_x_k8s_io_cluster_v1beta1"
}

func (r *ClusterXK8SIoClusterV1Beta1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Cluster is the Schema for the clusters API.",
		MarkdownDescription: "Cluster is the Schema for the clusters API.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"api_version": schema.StringAttribute{
				Description:         "The API group of the requested resource.",
				MarkdownDescription: "The API group of the requested resource.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"kind": schema.StringAttribute{
				Description:         "The type of the requested resource.",
				MarkdownDescription: "The type of the requested resource.",
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
				Description:         "ClusterSpec defines the desired state of Cluster.",
				MarkdownDescription: "ClusterSpec defines the desired state of Cluster.",
				Attributes: map[string]schema.Attribute{
					"cluster_network": schema.SingleNestedAttribute{
						Description:         "Cluster network configuration.",
						MarkdownDescription: "Cluster network configuration.",
						Attributes: map[string]schema.Attribute{
							"api_server_port": schema.Int64Attribute{
								Description:         "APIServerPort specifies the port the API Server should bind to. Defaults to 6443.",
								MarkdownDescription: "APIServerPort specifies the port the API Server should bind to. Defaults to 6443.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"pods": schema.SingleNestedAttribute{
								Description:         "The network ranges from which Pod networks are allocated.",
								MarkdownDescription: "The network ranges from which Pod networks are allocated.",
								Attributes: map[string]schema.Attribute{
									"cidr_blocks": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
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

							"service_domain": schema.StringAttribute{
								Description:         "Domain name for services.",
								MarkdownDescription: "Domain name for services.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"services": schema.SingleNestedAttribute{
								Description:         "The network ranges from which service VIPs are allocated.",
								MarkdownDescription: "The network ranges from which service VIPs are allocated.",
								Attributes: map[string]schema.Attribute{
									"cidr_blocks": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
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
						Required: false,
						Optional: false,
						Computed: true,
					},

					"control_plane_endpoint": schema.SingleNestedAttribute{
						Description:         "ControlPlaneEndpoint represents the endpoint used to communicate with the control plane.",
						MarkdownDescription: "ControlPlaneEndpoint represents the endpoint used to communicate with the control plane.",
						Attributes: map[string]schema.Attribute{
							"host": schema.StringAttribute{
								Description:         "The hostname on which the API server is serving.",
								MarkdownDescription: "The hostname on which the API server is serving.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"port": schema.Int64Attribute{
								Description:         "The port on which the API server is serving.",
								MarkdownDescription: "The port on which the API server is serving.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"control_plane_ref": schema.SingleNestedAttribute{
						Description:         "ControlPlaneRef is an optional reference to a provider-specific resource that holds the details for provisioning the Control Plane for a Cluster.",
						MarkdownDescription: "ControlPlaneRef is an optional reference to a provider-specific resource that holds the details for provisioning the Control Plane for a Cluster.",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "API version of the referent.",
								MarkdownDescription: "API version of the referent.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"field_path": schema.StringAttribute{
								Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
								MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								MarkdownDescription: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
								MarkdownDescription: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"resource_version": schema.StringAttribute{
								Description:         "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
								MarkdownDescription: "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"uid": schema.StringAttribute{
								Description:         "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
								MarkdownDescription: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"infrastructure_ref": schema.SingleNestedAttribute{
						Description:         "InfrastructureRef is a reference to a provider-specific resource that holds the details for provisioning infrastructure for a cluster in said provider.",
						MarkdownDescription: "InfrastructureRef is a reference to a provider-specific resource that holds the details for provisioning infrastructure for a cluster in said provider.",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "API version of the referent.",
								MarkdownDescription: "API version of the referent.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"field_path": schema.StringAttribute{
								Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
								MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								MarkdownDescription: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
								MarkdownDescription: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"resource_version": schema.StringAttribute{
								Description:         "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
								MarkdownDescription: "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"uid": schema.StringAttribute{
								Description:         "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
								MarkdownDescription: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"paused": schema.BoolAttribute{
						Description:         "Paused can be used to prevent controllers from processing the Cluster and all its associated objects.",
						MarkdownDescription: "Paused can be used to prevent controllers from processing the Cluster and all its associated objects.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"topology": schema.SingleNestedAttribute{
						Description:         "This encapsulates the topology for the cluster. NOTE: It is required to enable the ClusterTopology feature gate flag to activate managed topologies support; this feature is highly experimental, and parts of it might still be not implemented.",
						MarkdownDescription: "This encapsulates the topology for the cluster. NOTE: It is required to enable the ClusterTopology feature gate flag to activate managed topologies support; this feature is highly experimental, and parts of it might still be not implemented.",
						Attributes: map[string]schema.Attribute{
							"class": schema.StringAttribute{
								Description:         "The name of the ClusterClass object to create the topology.",
								MarkdownDescription: "The name of the ClusterClass object to create the topology.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"control_plane": schema.SingleNestedAttribute{
								Description:         "ControlPlane describes the cluster control plane.",
								MarkdownDescription: "ControlPlane describes the cluster control plane.",
								Attributes: map[string]schema.Attribute{
									"machine_health_check": schema.SingleNestedAttribute{
										Description:         "MachineHealthCheck allows to enable, disable and override the MachineHealthCheck configuration in the ClusterClass for this control plane.",
										MarkdownDescription: "MachineHealthCheck allows to enable, disable and override the MachineHealthCheck configuration in the ClusterClass for this control plane.",
										Attributes: map[string]schema.Attribute{
											"enable": schema.BoolAttribute{
												Description:         "Enable controls if a MachineHealthCheck should be created for the target machines.  If false: No MachineHealthCheck will be created.  If not set(default): A MachineHealthCheck will be created if it is defined here or in the associated ClusterClass. If no MachineHealthCheck is defined then none will be created.  If true: A MachineHealthCheck is guaranteed to be created. Cluster validation will block if 'enable' is true and no MachineHealthCheck definition is available.",
												MarkdownDescription: "Enable controls if a MachineHealthCheck should be created for the target machines.  If false: No MachineHealthCheck will be created.  If not set(default): A MachineHealthCheck will be created if it is defined here or in the associated ClusterClass. If no MachineHealthCheck is defined then none will be created.  If true: A MachineHealthCheck is guaranteed to be created. Cluster validation will block if 'enable' is true and no MachineHealthCheck definition is available.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"max_unhealthy": schema.StringAttribute{
												Description:         "Any further remediation is only allowed if at most 'MaxUnhealthy' machines selected by 'selector' are not healthy.",
												MarkdownDescription: "Any further remediation is only allowed if at most 'MaxUnhealthy' machines selected by 'selector' are not healthy.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"node_startup_timeout": schema.StringAttribute{
												Description:         "Machines older than this duration without a node will be considered to have failed and will be remediated. If you wish to disable this feature, set the value explicitly to 0.",
												MarkdownDescription: "Machines older than this duration without a node will be considered to have failed and will be remediated. If you wish to disable this feature, set the value explicitly to 0.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"remediation_template": schema.SingleNestedAttribute{
												Description:         "RemediationTemplate is a reference to a remediation template provided by an infrastructure provider.  This field is completely optional, when filled, the MachineHealthCheck controller creates a new object from the template referenced and hands off remediation of the machine to a controller that lives outside of Cluster API.",
												MarkdownDescription: "RemediationTemplate is a reference to a remediation template provided by an infrastructure provider.  This field is completely optional, when filled, the MachineHealthCheck controller creates a new object from the template referenced and hands off remediation of the machine to a controller that lives outside of Cluster API.",
												Attributes: map[string]schema.Attribute{
													"api_version": schema.StringAttribute{
														Description:         "API version of the referent.",
														MarkdownDescription: "API version of the referent.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"field_path": schema.StringAttribute{
														Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
														MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"kind": schema.StringAttribute{
														Description:         "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
														MarkdownDescription: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
														MarkdownDescription: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"resource_version": schema.StringAttribute{
														Description:         "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
														MarkdownDescription: "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"uid": schema.StringAttribute{
														Description:         "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
														MarkdownDescription: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"unhealthy_conditions": schema.ListNestedAttribute{
												Description:         "UnhealthyConditions contains a list of the conditions that determine whether a node is considered unhealthy. The conditions are combined in a logical OR, i.e. if any of the conditions is met, the node is unhealthy.",
												MarkdownDescription: "UnhealthyConditions contains a list of the conditions that determine whether a node is considered unhealthy. The conditions are combined in a logical OR, i.e. if any of the conditions is met, the node is unhealthy.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"status": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"timeout": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
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

											"unhealthy_range": schema.StringAttribute{
												Description:         "Any further remediation is only allowed if the number of machines selected by 'selector' as not healthy is within the range of 'UnhealthyRange'. Takes precedence over MaxUnhealthy. Eg. '[3-5]' - This means that remediation will be allowed only when: (a) there are at least 3 unhealthy machines (and) (b) there are at most 5 unhealthy machines",
												MarkdownDescription: "Any further remediation is only allowed if the number of machines selected by 'selector' as not healthy is within the range of 'UnhealthyRange'. Takes precedence over MaxUnhealthy. Eg. '[3-5]' - This means that remediation will be allowed only when: (a) there are at least 3 unhealthy machines (and) (b) there are at most 5 unhealthy machines",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"metadata": schema.SingleNestedAttribute{
										Description:         "Metadata is the metadata applied to the ControlPlane and the Machines of the ControlPlane if the ControlPlaneTemplate referenced by the ClusterClass is machine based. If not, it is applied only to the ControlPlane. At runtime this metadata is merged with the corresponding metadata from the ClusterClass.",
										MarkdownDescription: "Metadata is the metadata applied to the ControlPlane and the Machines of the ControlPlane if the ControlPlaneTemplate referenced by the ClusterClass is machine based. If not, it is applied only to the ControlPlane. At runtime this metadata is merged with the corresponding metadata from the ClusterClass.",
										Attributes: map[string]schema.Attribute{
											"annotations": schema.MapAttribute{
												Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
												MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"labels": schema.MapAttribute{
												Description:         "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
												MarkdownDescription: "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
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

									"node_deletion_timeout": schema.StringAttribute{
										Description:         "NodeDeletionTimeout defines how long the controller will attempt to delete the Node that the Machine hosts after the Machine is marked for deletion. A duration of 0 will retry deletion indefinitely. Defaults to 10 seconds.",
										MarkdownDescription: "NodeDeletionTimeout defines how long the controller will attempt to delete the Node that the Machine hosts after the Machine is marked for deletion. A duration of 0 will retry deletion indefinitely. Defaults to 10 seconds.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"node_drain_timeout": schema.StringAttribute{
										Description:         "NodeDrainTimeout is the total amount of time that the controller will spend on draining a node. The default value is 0, meaning that the node can be drained without any time limitations. NOTE: NodeDrainTimeout is different from 'kubectl drain --timeout'",
										MarkdownDescription: "NodeDrainTimeout is the total amount of time that the controller will spend on draining a node. The default value is 0, meaning that the node can be drained without any time limitations. NOTE: NodeDrainTimeout is different from 'kubectl drain --timeout'",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"node_volume_detach_timeout": schema.StringAttribute{
										Description:         "NodeVolumeDetachTimeout is the total amount of time that the controller will spend on waiting for all volumes to be detached. The default value is 0, meaning that the volumes can be detached without any time limitations.",
										MarkdownDescription: "NodeVolumeDetachTimeout is the total amount of time that the controller will spend on waiting for all volumes to be detached. The default value is 0, meaning that the volumes can be detached without any time limitations.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"replicas": schema.Int64Attribute{
										Description:         "Replicas is the number of control plane nodes. If the value is nil, the ControlPlane object is created without the number of Replicas and it's assumed that the control plane controller does not implement support for this field. When specified against a control plane provider that lacks support for this field, this value will be ignored.",
										MarkdownDescription: "Replicas is the number of control plane nodes. If the value is nil, the ControlPlane object is created without the number of Replicas and it's assumed that the control plane controller does not implement support for this field. When specified against a control plane provider that lacks support for this field, this value will be ignored.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"rollout_after": schema.StringAttribute{
								Description:         "RolloutAfter performs a rollout of the entire cluster one component at a time, control plane first and then machine deployments.  Deprecated: This field has no function and is going to be removed in the next apiVersion.",
								MarkdownDescription: "RolloutAfter performs a rollout of the entire cluster one component at a time, control plane first and then machine deployments.  Deprecated: This field has no function and is going to be removed in the next apiVersion.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"variables": schema.ListNestedAttribute{
								Description:         "Variables can be used to customize the Cluster through patches. They must comply to the corresponding VariableClasses defined in the ClusterClass.",
								MarkdownDescription: "Variables can be used to customize the Cluster through patches. They must comply to the corresponding VariableClasses defined in the ClusterClass.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"definition_from": schema.StringAttribute{
											Description:         "DefinitionFrom specifies where the definition of this Variable is from. DefinitionFrom is 'inline' when the definition is from the ClusterClass '.spec.variables' or the name of a patch defined in the ClusterClass '.spec.patches' where the patch is external and provides external variables. This field is mandatory if the variable has 'DefinitionsConflict: true' in ClusterClass 'status.variables[]'",
											MarkdownDescription: "DefinitionFrom specifies where the definition of this Variable is from. DefinitionFrom is 'inline' when the definition is from the ClusterClass '.spec.variables' or the name of a patch defined in the ClusterClass '.spec.patches' where the patch is external and provides external variables. This field is mandatory if the variable has 'DefinitionsConflict: true' in ClusterClass 'status.variables[]'",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"name": schema.StringAttribute{
											Description:         "Name of the variable.",
											MarkdownDescription: "Name of the variable.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"value": schema.MapAttribute{
											Description:         "Value of the variable. Note: the value will be validated against the schema of the corresponding ClusterClassVariable from the ClusterClass. Note: We have to use apiextensionsv1.JSON instead of a custom JSON type, because controller-tools has a hard-coded schema for apiextensionsv1.JSON which cannot be produced by another type via controller-tools, i.e. it is not possible to have no type field. Ref: https://github.com/kubernetes-sigs/controller-tools/blob/d0e03a142d0ecdd5491593e941ee1d6b5d91dba6/pkg/crd/known_types.go#L106-L111",
											MarkdownDescription: "Value of the variable. Note: the value will be validated against the schema of the corresponding ClusterClassVariable from the ClusterClass. Note: We have to use apiextensionsv1.JSON instead of a custom JSON type, because controller-tools has a hard-coded schema for apiextensionsv1.JSON which cannot be produced by another type via controller-tools, i.e. it is not possible to have no type field. Ref: https://github.com/kubernetes-sigs/controller-tools/blob/d0e03a142d0ecdd5491593e941ee1d6b5d91dba6/pkg/crd/known_types.go#L106-L111",
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

							"version": schema.StringAttribute{
								Description:         "The Kubernetes version of the cluster.",
								MarkdownDescription: "The Kubernetes version of the cluster.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"workers": schema.SingleNestedAttribute{
								Description:         "Workers encapsulates the different constructs that form the worker nodes for the cluster.",
								MarkdownDescription: "Workers encapsulates the different constructs that form the worker nodes for the cluster.",
								Attributes: map[string]schema.Attribute{
									"machine_deployments": schema.ListNestedAttribute{
										Description:         "MachineDeployments is a list of machine deployments in the cluster.",
										MarkdownDescription: "MachineDeployments is a list of machine deployments in the cluster.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"class": schema.StringAttribute{
													Description:         "Class is the name of the MachineDeploymentClass used to create the set of worker nodes. This should match one of the deployment classes defined in the ClusterClass object mentioned in the 'Cluster.Spec.Class' field.",
													MarkdownDescription: "Class is the name of the MachineDeploymentClass used to create the set of worker nodes. This should match one of the deployment classes defined in the ClusterClass object mentioned in the 'Cluster.Spec.Class' field.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"failure_domain": schema.StringAttribute{
													Description:         "FailureDomain is the failure domain the machines will be created in. Must match a key in the FailureDomains map stored on the cluster object.",
													MarkdownDescription: "FailureDomain is the failure domain the machines will be created in. Must match a key in the FailureDomains map stored on the cluster object.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"machine_health_check": schema.SingleNestedAttribute{
													Description:         "MachineHealthCheck allows to enable, disable and override the MachineHealthCheck configuration in the ClusterClass for this MachineDeployment.",
													MarkdownDescription: "MachineHealthCheck allows to enable, disable and override the MachineHealthCheck configuration in the ClusterClass for this MachineDeployment.",
													Attributes: map[string]schema.Attribute{
														"enable": schema.BoolAttribute{
															Description:         "Enable controls if a MachineHealthCheck should be created for the target machines.  If false: No MachineHealthCheck will be created.  If not set(default): A MachineHealthCheck will be created if it is defined here or in the associated ClusterClass. If no MachineHealthCheck is defined then none will be created.  If true: A MachineHealthCheck is guaranteed to be created. Cluster validation will block if 'enable' is true and no MachineHealthCheck definition is available.",
															MarkdownDescription: "Enable controls if a MachineHealthCheck should be created for the target machines.  If false: No MachineHealthCheck will be created.  If not set(default): A MachineHealthCheck will be created if it is defined here or in the associated ClusterClass. If no MachineHealthCheck is defined then none will be created.  If true: A MachineHealthCheck is guaranteed to be created. Cluster validation will block if 'enable' is true and no MachineHealthCheck definition is available.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"max_unhealthy": schema.StringAttribute{
															Description:         "Any further remediation is only allowed if at most 'MaxUnhealthy' machines selected by 'selector' are not healthy.",
															MarkdownDescription: "Any further remediation is only allowed if at most 'MaxUnhealthy' machines selected by 'selector' are not healthy.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"node_startup_timeout": schema.StringAttribute{
															Description:         "Machines older than this duration without a node will be considered to have failed and will be remediated. If you wish to disable this feature, set the value explicitly to 0.",
															MarkdownDescription: "Machines older than this duration without a node will be considered to have failed and will be remediated. If you wish to disable this feature, set the value explicitly to 0.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"remediation_template": schema.SingleNestedAttribute{
															Description:         "RemediationTemplate is a reference to a remediation template provided by an infrastructure provider.  This field is completely optional, when filled, the MachineHealthCheck controller creates a new object from the template referenced and hands off remediation of the machine to a controller that lives outside of Cluster API.",
															MarkdownDescription: "RemediationTemplate is a reference to a remediation template provided by an infrastructure provider.  This field is completely optional, when filled, the MachineHealthCheck controller creates a new object from the template referenced and hands off remediation of the machine to a controller that lives outside of Cluster API.",
															Attributes: map[string]schema.Attribute{
																"api_version": schema.StringAttribute{
																	Description:         "API version of the referent.",
																	MarkdownDescription: "API version of the referent.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"field_path": schema.StringAttribute{
																	Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
																	MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"kind": schema.StringAttribute{
																	Description:         "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
																	MarkdownDescription: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"name": schema.StringAttribute{
																	Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"namespace": schema.StringAttribute{
																	Description:         "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
																	MarkdownDescription: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"resource_version": schema.StringAttribute{
																	Description:         "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
																	MarkdownDescription: "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"uid": schema.StringAttribute{
																	Description:         "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
																	MarkdownDescription: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"unhealthy_conditions": schema.ListNestedAttribute{
															Description:         "UnhealthyConditions contains a list of the conditions that determine whether a node is considered unhealthy. The conditions are combined in a logical OR, i.e. if any of the conditions is met, the node is unhealthy.",
															MarkdownDescription: "UnhealthyConditions contains a list of the conditions that determine whether a node is considered unhealthy. The conditions are combined in a logical OR, i.e. if any of the conditions is met, the node is unhealthy.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"status": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"timeout": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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

														"unhealthy_range": schema.StringAttribute{
															Description:         "Any further remediation is only allowed if the number of machines selected by 'selector' as not healthy is within the range of 'UnhealthyRange'. Takes precedence over MaxUnhealthy. Eg. '[3-5]' - This means that remediation will be allowed only when: (a) there are at least 3 unhealthy machines (and) (b) there are at most 5 unhealthy machines",
															MarkdownDescription: "Any further remediation is only allowed if the number of machines selected by 'selector' as not healthy is within the range of 'UnhealthyRange'. Takes precedence over MaxUnhealthy. Eg. '[3-5]' - This means that remediation will be allowed only when: (a) there are at least 3 unhealthy machines (and) (b) there are at most 5 unhealthy machines",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"metadata": schema.SingleNestedAttribute{
													Description:         "Metadata is the metadata applied to the MachineDeployment and the machines of the MachineDeployment. At runtime this metadata is merged with the corresponding metadata from the ClusterClass.",
													MarkdownDescription: "Metadata is the metadata applied to the MachineDeployment and the machines of the MachineDeployment. At runtime this metadata is merged with the corresponding metadata from the ClusterClass.",
													Attributes: map[string]schema.Attribute{
														"annotations": schema.MapAttribute{
															Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
															MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"labels": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
															MarkdownDescription: "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
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

												"min_ready_seconds": schema.Int64Attribute{
													Description:         "Minimum number of seconds for which a newly created machine should be ready. Defaults to 0 (machine will be considered available as soon as it is ready)",
													MarkdownDescription: "Minimum number of seconds for which a newly created machine should be ready. Defaults to 0 (machine will be considered available as soon as it is ready)",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"name": schema.StringAttribute{
													Description:         "Name is the unique identifier for this MachineDeploymentTopology. The value is used with other unique identifiers to create a MachineDeployment's Name (e.g. cluster's name, etc). In case the name is greater than the allowed maximum length, the values are hashed together.",
													MarkdownDescription: "Name is the unique identifier for this MachineDeploymentTopology. The value is used with other unique identifiers to create a MachineDeployment's Name (e.g. cluster's name, etc). In case the name is greater than the allowed maximum length, the values are hashed together.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"node_deletion_timeout": schema.StringAttribute{
													Description:         "NodeDeletionTimeout defines how long the controller will attempt to delete the Node that the Machine hosts after the Machine is marked for deletion. A duration of 0 will retry deletion indefinitely. Defaults to 10 seconds.",
													MarkdownDescription: "NodeDeletionTimeout defines how long the controller will attempt to delete the Node that the Machine hosts after the Machine is marked for deletion. A duration of 0 will retry deletion indefinitely. Defaults to 10 seconds.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"node_drain_timeout": schema.StringAttribute{
													Description:         "NodeDrainTimeout is the total amount of time that the controller will spend on draining a node. The default value is 0, meaning that the node can be drained without any time limitations. NOTE: NodeDrainTimeout is different from 'kubectl drain --timeout'",
													MarkdownDescription: "NodeDrainTimeout is the total amount of time that the controller will spend on draining a node. The default value is 0, meaning that the node can be drained without any time limitations. NOTE: NodeDrainTimeout is different from 'kubectl drain --timeout'",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"node_volume_detach_timeout": schema.StringAttribute{
													Description:         "NodeVolumeDetachTimeout is the total amount of time that the controller will spend on waiting for all volumes to be detached. The default value is 0, meaning that the volumes can be detached without any time limitations.",
													MarkdownDescription: "NodeVolumeDetachTimeout is the total amount of time that the controller will spend on waiting for all volumes to be detached. The default value is 0, meaning that the volumes can be detached without any time limitations.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"replicas": schema.Int64Attribute{
													Description:         "Replicas is the number of worker nodes belonging to this set. If the value is nil, the MachineDeployment is created without the number of Replicas (defaulting to 1) and it's assumed that an external entity (like cluster autoscaler) is responsible for the management of this value.",
													MarkdownDescription: "Replicas is the number of worker nodes belonging to this set. If the value is nil, the MachineDeployment is created without the number of Replicas (defaulting to 1) and it's assumed that an external entity (like cluster autoscaler) is responsible for the management of this value.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"strategy": schema.SingleNestedAttribute{
													Description:         "The deployment strategy to use to replace existing machines with new ones.",
													MarkdownDescription: "The deployment strategy to use to replace existing machines with new ones.",
													Attributes: map[string]schema.Attribute{
														"rolling_update": schema.SingleNestedAttribute{
															Description:         "Rolling update config params. Present only if MachineDeploymentStrategyType = RollingUpdate.",
															MarkdownDescription: "Rolling update config params. Present only if MachineDeploymentStrategyType = RollingUpdate.",
															Attributes: map[string]schema.Attribute{
																"delete_policy": schema.StringAttribute{
																	Description:         "DeletePolicy defines the policy used by the MachineDeployment to identify nodes to delete when downscaling. Valid values are 'Random, 'Newest', 'Oldest' When no value is supplied, the default DeletePolicy of MachineSet is used",
																	MarkdownDescription: "DeletePolicy defines the policy used by the MachineDeployment to identify nodes to delete when downscaling. Valid values are 'Random, 'Newest', 'Oldest' When no value is supplied, the default DeletePolicy of MachineSet is used",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"max_surge": schema.StringAttribute{
																	Description:         "The maximum number of machines that can be scheduled above the desired number of machines. Value can be an absolute number (ex: 5) or a percentage of desired machines (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up. Defaults to 1. Example: when this is set to 30%, the new MachineSet can be scaled up immediately when the rolling update starts, such that the total number of old and new machines do not exceed 130% of desired machines. Once old machines have been killed, new MachineSet can be scaled up further, ensuring that total number of machines running at any time during the update is at most 130% of desired machines.",
																	MarkdownDescription: "The maximum number of machines that can be scheduled above the desired number of machines. Value can be an absolute number (ex: 5) or a percentage of desired machines (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up. Defaults to 1. Example: when this is set to 30%, the new MachineSet can be scaled up immediately when the rolling update starts, such that the total number of old and new machines do not exceed 130% of desired machines. Once old machines have been killed, new MachineSet can be scaled up further, ensuring that total number of machines running at any time during the update is at most 130% of desired machines.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"max_unavailable": schema.StringAttribute{
																	Description:         "The maximum number of machines that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired machines (ex: 10%). Absolute number is calculated from percentage by rounding down. This can not be 0 if MaxSurge is 0. Defaults to 0. Example: when this is set to 30%, the old MachineSet can be scaled down to 70% of desired machines immediately when the rolling update starts. Once new machines are ready, old MachineSet can be scaled down further, followed by scaling up the new MachineSet, ensuring that the total number of machines available at all times during the update is at least 70% of desired machines.",
																	MarkdownDescription: "The maximum number of machines that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired machines (ex: 10%). Absolute number is calculated from percentage by rounding down. This can not be 0 if MaxSurge is 0. Defaults to 0. Example: when this is set to 30%, the old MachineSet can be scaled down to 70% of desired machines immediately when the rolling update starts. Once new machines are ready, old MachineSet can be scaled down further, followed by scaling up the new MachineSet, ensuring that the total number of machines available at all times during the update is at least 70% of desired machines.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"type": schema.StringAttribute{
															Description:         "Type of deployment. Allowed values are RollingUpdate and OnDelete. The default is RollingUpdate.",
															MarkdownDescription: "Type of deployment. Allowed values are RollingUpdate and OnDelete. The default is RollingUpdate.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"variables": schema.SingleNestedAttribute{
													Description:         "Variables can be used to customize the MachineDeployment through patches.",
													MarkdownDescription: "Variables can be used to customize the MachineDeployment through patches.",
													Attributes: map[string]schema.Attribute{
														"overrides": schema.ListNestedAttribute{
															Description:         "Overrides can be used to override Cluster level variables.",
															MarkdownDescription: "Overrides can be used to override Cluster level variables.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"definition_from": schema.StringAttribute{
																		Description:         "DefinitionFrom specifies where the definition of this Variable is from. DefinitionFrom is 'inline' when the definition is from the ClusterClass '.spec.variables' or the name of a patch defined in the ClusterClass '.spec.patches' where the patch is external and provides external variables. This field is mandatory if the variable has 'DefinitionsConflict: true' in ClusterClass 'status.variables[]'",
																		MarkdownDescription: "DefinitionFrom specifies where the definition of this Variable is from. DefinitionFrom is 'inline' when the definition is from the ClusterClass '.spec.variables' or the name of a patch defined in the ClusterClass '.spec.patches' where the patch is external and provides external variables. This field is mandatory if the variable has 'DefinitionsConflict: true' in ClusterClass 'status.variables[]'",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the variable.",
																		MarkdownDescription: "Name of the variable.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"value": schema.MapAttribute{
																		Description:         "Value of the variable. Note: the value will be validated against the schema of the corresponding ClusterClassVariable from the ClusterClass. Note: We have to use apiextensionsv1.JSON instead of a custom JSON type, because controller-tools has a hard-coded schema for apiextensionsv1.JSON which cannot be produced by another type via controller-tools, i.e. it is not possible to have no type field. Ref: https://github.com/kubernetes-sigs/controller-tools/blob/d0e03a142d0ecdd5491593e941ee1d6b5d91dba6/pkg/crd/known_types.go#L106-L111",
																		MarkdownDescription: "Value of the variable. Note: the value will be validated against the schema of the corresponding ClusterClassVariable from the ClusterClass. Note: We have to use apiextensionsv1.JSON instead of a custom JSON type, because controller-tools has a hard-coded schema for apiextensionsv1.JSON which cannot be produced by another type via controller-tools, i.e. it is not possible to have no type field. Ref: https://github.com/kubernetes-sigs/controller-tools/blob/d0e03a142d0ecdd5491593e941ee1d6b5d91dba6/pkg/crd/known_types.go#L106-L111",
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
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"machine_pools": schema.ListNestedAttribute{
										Description:         "MachinePools is a list of machine pools in the cluster.",
										MarkdownDescription: "MachinePools is a list of machine pools in the cluster.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"class": schema.StringAttribute{
													Description:         "Class is the name of the MachinePoolClass used to create the pool of worker nodes. This should match one of the deployment classes defined in the ClusterClass object mentioned in the 'Cluster.Spec.Class' field.",
													MarkdownDescription: "Class is the name of the MachinePoolClass used to create the pool of worker nodes. This should match one of the deployment classes defined in the ClusterClass object mentioned in the 'Cluster.Spec.Class' field.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"failure_domains": schema.ListAttribute{
													Description:         "FailureDomains is the list of failure domains the machine pool will be created in. Must match a key in the FailureDomains map stored on the cluster object.",
													MarkdownDescription: "FailureDomains is the list of failure domains the machine pool will be created in. Must match a key in the FailureDomains map stored on the cluster object.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"metadata": schema.SingleNestedAttribute{
													Description:         "Metadata is the metadata applied to the MachinePool. At runtime this metadata is merged with the corresponding metadata from the ClusterClass.",
													MarkdownDescription: "Metadata is the metadata applied to the MachinePool. At runtime this metadata is merged with the corresponding metadata from the ClusterClass.",
													Attributes: map[string]schema.Attribute{
														"annotations": schema.MapAttribute{
															Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
															MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"labels": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
															MarkdownDescription: "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
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

												"min_ready_seconds": schema.Int64Attribute{
													Description:         "Minimum number of seconds for which a newly created machine pool should be ready. Defaults to 0 (machine will be considered available as soon as it is ready)",
													MarkdownDescription: "Minimum number of seconds for which a newly created machine pool should be ready. Defaults to 0 (machine will be considered available as soon as it is ready)",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"name": schema.StringAttribute{
													Description:         "Name is the unique identifier for this MachinePoolTopology. The value is used with other unique identifiers to create a MachinePool's Name (e.g. cluster's name, etc). In case the name is greater than the allowed maximum length, the values are hashed together.",
													MarkdownDescription: "Name is the unique identifier for this MachinePoolTopology. The value is used with other unique identifiers to create a MachinePool's Name (e.g. cluster's name, etc). In case the name is greater than the allowed maximum length, the values are hashed together.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"node_deletion_timeout": schema.StringAttribute{
													Description:         "NodeDeletionTimeout defines how long the controller will attempt to delete the Node that the MachinePool hosts after the MachinePool is marked for deletion. A duration of 0 will retry deletion indefinitely. Defaults to 10 seconds.",
													MarkdownDescription: "NodeDeletionTimeout defines how long the controller will attempt to delete the Node that the MachinePool hosts after the MachinePool is marked for deletion. A duration of 0 will retry deletion indefinitely. Defaults to 10 seconds.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"node_drain_timeout": schema.StringAttribute{
													Description:         "NodeDrainTimeout is the total amount of time that the controller will spend on draining a node. The default value is 0, meaning that the node can be drained without any time limitations. NOTE: NodeDrainTimeout is different from 'kubectl drain --timeout'",
													MarkdownDescription: "NodeDrainTimeout is the total amount of time that the controller will spend on draining a node. The default value is 0, meaning that the node can be drained without any time limitations. NOTE: NodeDrainTimeout is different from 'kubectl drain --timeout'",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"node_volume_detach_timeout": schema.StringAttribute{
													Description:         "NodeVolumeDetachTimeout is the total amount of time that the controller will spend on waiting for all volumes to be detached. The default value is 0, meaning that the volumes can be detached without any time limitations.",
													MarkdownDescription: "NodeVolumeDetachTimeout is the total amount of time that the controller will spend on waiting for all volumes to be detached. The default value is 0, meaning that the volumes can be detached without any time limitations.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"replicas": schema.Int64Attribute{
													Description:         "Replicas is the number of nodes belonging to this pool. If the value is nil, the MachinePool is created without the number of Replicas (defaulting to 1) and it's assumed that an external entity (like cluster autoscaler) is responsible for the management of this value.",
													MarkdownDescription: "Replicas is the number of nodes belonging to this pool. If the value is nil, the MachinePool is created without the number of Replicas (defaulting to 1) and it's assumed that an external entity (like cluster autoscaler) is responsible for the management of this value.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"variables": schema.SingleNestedAttribute{
													Description:         "Variables can be used to customize the MachinePool through patches.",
													MarkdownDescription: "Variables can be used to customize the MachinePool through patches.",
													Attributes: map[string]schema.Attribute{
														"overrides": schema.ListNestedAttribute{
															Description:         "Overrides can be used to override Cluster level variables.",
															MarkdownDescription: "Overrides can be used to override Cluster level variables.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"definition_from": schema.StringAttribute{
																		Description:         "DefinitionFrom specifies where the definition of this Variable is from. DefinitionFrom is 'inline' when the definition is from the ClusterClass '.spec.variables' or the name of a patch defined in the ClusterClass '.spec.patches' where the patch is external and provides external variables. This field is mandatory if the variable has 'DefinitionsConflict: true' in ClusterClass 'status.variables[]'",
																		MarkdownDescription: "DefinitionFrom specifies where the definition of this Variable is from. DefinitionFrom is 'inline' when the definition is from the ClusterClass '.spec.variables' or the name of a patch defined in the ClusterClass '.spec.patches' where the patch is external and provides external variables. This field is mandatory if the variable has 'DefinitionsConflict: true' in ClusterClass 'status.variables[]'",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the variable.",
																		MarkdownDescription: "Name of the variable.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"value": schema.MapAttribute{
																		Description:         "Value of the variable. Note: the value will be validated against the schema of the corresponding ClusterClassVariable from the ClusterClass. Note: We have to use apiextensionsv1.JSON instead of a custom JSON type, because controller-tools has a hard-coded schema for apiextensionsv1.JSON which cannot be produced by another type via controller-tools, i.e. it is not possible to have no type field. Ref: https://github.com/kubernetes-sigs/controller-tools/blob/d0e03a142d0ecdd5491593e941ee1d6b5d91dba6/pkg/crd/known_types.go#L106-L111",
																		MarkdownDescription: "Value of the variable. Note: the value will be validated against the schema of the corresponding ClusterClassVariable from the ClusterClass. Note: We have to use apiextensionsv1.JSON instead of a custom JSON type, because controller-tools has a hard-coded schema for apiextensionsv1.JSON which cannot be produced by another type via controller-tools, i.e. it is not possible to have no type field. Ref: https://github.com/kubernetes-sigs/controller-tools/blob/d0e03a142d0ecdd5491593e941ee1d6b5d91dba6/pkg/crd/known_types.go#L106-L111",
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
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *ClusterXK8SIoClusterV1Beta1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedDataSourceDataError(request.ProviderData))
	}
}

func (r *ClusterXK8SIoClusterV1Beta1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_cluster_x_k8s_io_cluster_v1beta1")

	var data ClusterXK8SIoClusterV1Beta1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "cluster.x-k8s.io", Version: "v1beta1", Resource: "clusters"}).
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

	var readResponse ClusterXK8SIoClusterV1Beta1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("cluster.x-k8s.io/v1beta1")
	data.Kind = pointer.String("Cluster")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
