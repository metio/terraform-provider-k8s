/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package cluster_x_k8s_io_v1beta2

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &ClusterXK8SIoMachinePoolV1Beta2Manifest{}
)

func NewClusterXK8SIoMachinePoolV1Beta2Manifest() datasource.DataSource {
	return &ClusterXK8SIoMachinePoolV1Beta2Manifest{}
}

type ClusterXK8SIoMachinePoolV1Beta2Manifest struct{}

type ClusterXK8SIoMachinePoolV1Beta2ManifestData struct {
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
		ClusterName    *string   `tfsdk:"cluster_name" json:"clusterName,omitempty"`
		FailureDomains *[]string `tfsdk:"failure_domains" json:"failureDomains,omitempty"`
		ProviderIDList *[]string `tfsdk:"provider_id_list" json:"providerIDList,omitempty"`
		Replicas       *int64    `tfsdk:"replicas" json:"replicas,omitempty"`
		Template       *struct {
			Metadata *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"metadata" json:"metadata,omitempty"`
			Spec *struct {
				Bootstrap *struct {
					ConfigRef *struct {
						ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
						FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
						Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
						Name            *string `tfsdk:"name" json:"name,omitempty"`
						Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
						ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
						Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
					} `tfsdk:"config_ref" json:"configRef,omitempty"`
					DataSecretName *string `tfsdk:"data_secret_name" json:"dataSecretName,omitempty"`
				} `tfsdk:"bootstrap" json:"bootstrap,omitempty"`
				ClusterName       *string `tfsdk:"cluster_name" json:"clusterName,omitempty"`
				FailureDomain     *string `tfsdk:"failure_domain" json:"failureDomain,omitempty"`
				InfrastructureRef *struct {
					ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
					FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
					Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
					Name            *string `tfsdk:"name" json:"name,omitempty"`
					Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
					ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
					Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
				} `tfsdk:"infrastructure_ref" json:"infrastructureRef,omitempty"`
				MinReadySeconds                *int64  `tfsdk:"min_ready_seconds" json:"minReadySeconds,omitempty"`
				NodeDeletionTimeoutSeconds     *int64  `tfsdk:"node_deletion_timeout_seconds" json:"nodeDeletionTimeoutSeconds,omitempty"`
				NodeDrainTimeoutSeconds        *int64  `tfsdk:"node_drain_timeout_seconds" json:"nodeDrainTimeoutSeconds,omitempty"`
				NodeVolumeDetachTimeoutSeconds *int64  `tfsdk:"node_volume_detach_timeout_seconds" json:"nodeVolumeDetachTimeoutSeconds,omitempty"`
				ProviderID                     *string `tfsdk:"provider_id" json:"providerID,omitempty"`
				ReadinessGates                 *[]struct {
					ConditionType *string `tfsdk:"condition_type" json:"conditionType,omitempty"`
					Polarity      *string `tfsdk:"polarity" json:"polarity,omitempty"`
				} `tfsdk:"readiness_gates" json:"readinessGates,omitempty"`
				Version *string `tfsdk:"version" json:"version,omitempty"`
			} `tfsdk:"spec" json:"spec,omitempty"`
		} `tfsdk:"template" json:"template,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ClusterXK8SIoMachinePoolV1Beta2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_cluster_x_k8s_io_machine_pool_v1beta2_manifest"
}

func (r *ClusterXK8SIoMachinePoolV1Beta2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "MachinePool is the Schema for the machinepools API. NOTE: This CRD can only be used if the MachinePool feature gate is enabled.",
		MarkdownDescription: "MachinePool is the Schema for the machinepools API. NOTE: This CRD can only be used if the MachinePool feature gate is enabled.",
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
				Description:         "spec is the desired state of MachinePool.",
				MarkdownDescription: "spec is the desired state of MachinePool.",
				Attributes: map[string]schema.Attribute{
					"cluster_name": schema.StringAttribute{
						Description:         "clusterName is the name of the Cluster this object belongs to.",
						MarkdownDescription: "clusterName is the name of the Cluster this object belongs to.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
							stringvalidator.LengthAtMost(63),
						},
					},

					"failure_domains": schema.ListAttribute{
						Description:         "failureDomains is the list of failure domains this MachinePool should be attached to.",
						MarkdownDescription: "failureDomains is the list of failure domains this MachinePool should be attached to.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"provider_id_list": schema.ListAttribute{
						Description:         "providerIDList are the identification IDs of machine instances provided by the provider. This field must match the provider IDs as seen on the node objects corresponding to a machine pool's machine instances.",
						MarkdownDescription: "providerIDList are the identification IDs of machine instances provided by the provider. This field must match the provider IDs as seen on the node objects corresponding to a machine pool's machine instances.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"replicas": schema.Int64Attribute{
						Description:         "replicas is the number of desired machines. Defaults to 1. This is a pointer to distinguish between explicit zero and not specified.",
						MarkdownDescription: "replicas is the number of desired machines. Defaults to 1. This is a pointer to distinguish between explicit zero and not specified.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"template": schema.SingleNestedAttribute{
						Description:         "template describes the machines that will be created.",
						MarkdownDescription: "template describes the machines that will be created.",
						Attributes: map[string]schema.Attribute{
							"metadata": schema.SingleNestedAttribute{
								Description:         "metadata is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
								MarkdownDescription: "metadata is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
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

							"spec": schema.SingleNestedAttribute{
								Description:         "spec is the specification of the desired behavior of the machine. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
								MarkdownDescription: "spec is the specification of the desired behavior of the machine. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
								Attributes: map[string]schema.Attribute{
									"bootstrap": schema.SingleNestedAttribute{
										Description:         "bootstrap is a reference to a local struct which encapsulates fields to configure the Machine’s bootstrapping mechanism.",
										MarkdownDescription: "bootstrap is a reference to a local struct which encapsulates fields to configure the Machine’s bootstrapping mechanism.",
										Attributes: map[string]schema.Attribute{
											"config_ref": schema.SingleNestedAttribute{
												Description:         "configRef is a reference to a bootstrap provider-specific resource that holds configuration details. The reference is optional to allow users/operators to specify Bootstrap.DataSecretName without the need of a controller.",
												MarkdownDescription: "configRef is a reference to a bootstrap provider-specific resource that holds configuration details. The reference is optional to allow users/operators to specify Bootstrap.DataSecretName without the need of a controller.",
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

											"data_secret_name": schema.StringAttribute{
												Description:         "dataSecretName is the name of the secret that stores the bootstrap data script. If nil, the Machine should remain in the Pending state.",
												MarkdownDescription: "dataSecretName is the name of the secret that stores the bootstrap data script. If nil, the Machine should remain in the Pending state.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(0),
													stringvalidator.LengthAtMost(253),
												},
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"cluster_name": schema.StringAttribute{
										Description:         "clusterName is the name of the Cluster this object belongs to.",
										MarkdownDescription: "clusterName is the name of the Cluster this object belongs to.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(1),
											stringvalidator.LengthAtMost(63),
										},
									},

									"failure_domain": schema.StringAttribute{
										Description:         "failureDomain is the failure domain the machine will be created in. Must match the name of a FailureDomain from the Cluster status.",
										MarkdownDescription: "failureDomain is the failure domain the machine will be created in. Must match the name of a FailureDomain from the Cluster status.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(1),
											stringvalidator.LengthAtMost(256),
										},
									},

									"infrastructure_ref": schema.SingleNestedAttribute{
										Description:         "infrastructureRef is a required reference to a custom resource offered by an infrastructure provider.",
										MarkdownDescription: "infrastructureRef is a required reference to a custom resource offered by an infrastructure provider.",
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
										Required: true,
										Optional: false,
										Computed: false,
									},

									"min_ready_seconds": schema.Int64Attribute{
										Description:         "minReadySeconds is the minimum number of seconds for which a Machine should be ready before considering it available. Defaults to 0 (Machine will be considered available as soon as the Machine is ready)",
										MarkdownDescription: "minReadySeconds is the minimum number of seconds for which a Machine should be ready before considering it available. Defaults to 0 (Machine will be considered available as soon as the Machine is ready)",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"node_deletion_timeout_seconds": schema.Int64Attribute{
										Description:         "nodeDeletionTimeoutSeconds defines how long the controller will attempt to delete the Node that the Machine hosts after the Machine is marked for deletion. A duration of 0 will retry deletion indefinitely. Defaults to 10 seconds.",
										MarkdownDescription: "nodeDeletionTimeoutSeconds defines how long the controller will attempt to delete the Node that the Machine hosts after the Machine is marked for deletion. A duration of 0 will retry deletion indefinitely. Defaults to 10 seconds.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
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
										},
									},

									"provider_id": schema.StringAttribute{
										Description:         "providerID is the identification ID of the machine provided by the provider. This field must match the provider ID as seen on the node object corresponding to this machine. This field is required by higher level consumers of cluster-api. Example use case is cluster autoscaler with cluster-api as provider. Clean-up logic in the autoscaler compares machines to nodes to find out machines at provider which could not get registered as Kubernetes nodes. With cluster-api as a generic out-of-tree provider for autoscaler, this field is required by autoscaler to be able to have a provider view of the list of machines. Another list of nodes is queried from the k8s apiserver and then a comparison is done to find out unregistered machines and are marked for delete. This field will be set by the actuators and consumed by higher level entities like autoscaler that will be interfacing with cluster-api as generic provider.",
										MarkdownDescription: "providerID is the identification ID of the machine provided by the provider. This field must match the provider ID as seen on the node object corresponding to this machine. This field is required by higher level consumers of cluster-api. Example use case is cluster autoscaler with cluster-api as provider. Clean-up logic in the autoscaler compares machines to nodes to find out machines at provider which could not get registered as Kubernetes nodes. With cluster-api as a generic out-of-tree provider for autoscaler, this field is required by autoscaler to be able to have a provider view of the list of machines. Another list of nodes is queried from the k8s apiserver and then a comparison is done to find out unregistered machines and are marked for delete. This field will be set by the actuators and consumed by higher level entities like autoscaler that will be interfacing with cluster-api as generic provider.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(1),
											stringvalidator.LengthAtMost(512),
										},
									},

									"readiness_gates": schema.ListNestedAttribute{
										Description:         "readinessGates specifies additional conditions to include when evaluating Machine Ready condition. This field can be used e.g. by Cluster API control plane providers to extend the semantic of the Ready condition for the Machine they control, like the kubeadm control provider adding ReadinessGates for the APIServerPodHealthy, SchedulerPodHealthy conditions, etc. Another example are external controllers, e.g. responsible to install special software/hardware on the Machines; they can include the status of those components with a new condition and add this condition to ReadinessGates. NOTE: This field is considered only for computing v1beta2 conditions. NOTE: In case readinessGates conditions start with the APIServer, ControllerManager, Scheduler prefix, and all those readiness gates condition are reporting the same message, when computing the Machine's Ready condition those readinessGates will be replaced by a single entry reporting 'Control plane components: ' + message. This helps to improve readability of conditions bubbling up to the Machine's owner resource / to the Cluster).",
										MarkdownDescription: "readinessGates specifies additional conditions to include when evaluating Machine Ready condition. This field can be used e.g. by Cluster API control plane providers to extend the semantic of the Ready condition for the Machine they control, like the kubeadm control provider adding ReadinessGates for the APIServerPodHealthy, SchedulerPodHealthy conditions, etc. Another example are external controllers, e.g. responsible to install special software/hardware on the Machines; they can include the status of those components with a new condition and add this condition to ReadinessGates. NOTE: This field is considered only for computing v1beta2 conditions. NOTE: In case readinessGates conditions start with the APIServer, ControllerManager, Scheduler prefix, and all those readiness gates condition are reporting the same message, when computing the Machine's Ready condition those readinessGates will be replaced by a single entry reporting 'Control plane components: ' + message. This helps to improve readability of conditions bubbling up to the Machine's owner resource / to the Cluster).",
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

									"version": schema.StringAttribute{
										Description:         "version defines the desired Kubernetes version. This field is meant to be optionally used by bootstrap providers.",
										MarkdownDescription: "version defines the desired Kubernetes version. This field is meant to be optionally used by bootstrap providers.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(1),
											stringvalidator.LengthAtMost(256),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
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

func (r *ClusterXK8SIoMachinePoolV1Beta2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_cluster_x_k8s_io_machine_pool_v1beta2_manifest")

	var model ClusterXK8SIoMachinePoolV1Beta2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("cluster.x-k8s.io/v1beta2")
	model.Kind = pointer.String("MachinePool")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
