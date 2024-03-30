/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package cluster_x_k8s_io_v1alpha4

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
	_ datasource.DataSource = &ClusterXK8SIoMachineSetV1Alpha4Manifest{}
)

func NewClusterXK8SIoMachineSetV1Alpha4Manifest() datasource.DataSource {
	return &ClusterXK8SIoMachineSetV1Alpha4Manifest{}
}

type ClusterXK8SIoMachineSetV1Alpha4Manifest struct{}

type ClusterXK8SIoMachineSetV1Alpha4ManifestData struct {
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
		ClusterName     *string `tfsdk:"cluster_name" json:"clusterName,omitempty"`
		DeletePolicy    *string `tfsdk:"delete_policy" json:"deletePolicy,omitempty"`
		MinReadySeconds *int64  `tfsdk:"min_ready_seconds" json:"minReadySeconds,omitempty"`
		Replicas        *int64  `tfsdk:"replicas" json:"replicas,omitempty"`
		Selector        *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"selector" json:"selector,omitempty"`
		Template *struct {
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
				NodeDrainTimeout *string `tfsdk:"node_drain_timeout" json:"nodeDrainTimeout,omitempty"`
				ProviderID       *string `tfsdk:"provider_id" json:"providerID,omitempty"`
				Version          *string `tfsdk:"version" json:"version,omitempty"`
			} `tfsdk:"spec" json:"spec,omitempty"`
		} `tfsdk:"template" json:"template,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ClusterXK8SIoMachineSetV1Alpha4Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_cluster_x_k8s_io_machine_set_v1alpha4_manifest"
}

func (r *ClusterXK8SIoMachineSetV1Alpha4Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "MachineSet is the Schema for the machinesets API.Deprecated: This type will be removed in one of the next releases.",
		MarkdownDescription: "MachineSet is the Schema for the machinesets API.Deprecated: This type will be removed in one of the next releases.",
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
				Description:         "MachineSetSpec defines the desired state of MachineSet.",
				MarkdownDescription: "MachineSetSpec defines the desired state of MachineSet.",
				Attributes: map[string]schema.Attribute{
					"cluster_name": schema.StringAttribute{
						Description:         "ClusterName is the name of the Cluster this object belongs to.",
						MarkdownDescription: "ClusterName is the name of the Cluster this object belongs to.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
						},
					},

					"delete_policy": schema.StringAttribute{
						Description:         "DeletePolicy defines the policy used to identify nodes to delete when downscaling.Defaults to 'Random'.  Valid values are 'Random, 'Newest', 'Oldest'",
						MarkdownDescription: "DeletePolicy defines the policy used to identify nodes to delete when downscaling.Defaults to 'Random'.  Valid values are 'Random, 'Newest', 'Oldest'",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Random", "Newest", "Oldest"),
						},
					},

					"min_ready_seconds": schema.Int64Attribute{
						Description:         "MinReadySeconds is the minimum number of seconds for which a newly created machine should be ready.Defaults to 0 (machine will be considered available as soon as it is ready)",
						MarkdownDescription: "MinReadySeconds is the minimum number of seconds for which a newly created machine should be ready.Defaults to 0 (machine will be considered available as soon as it is ready)",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"replicas": schema.Int64Attribute{
						Description:         "Replicas is the number of desired replicas.This is a pointer to distinguish between explicit zero and unspecified.Defaults to 1.",
						MarkdownDescription: "Replicas is the number of desired replicas.This is a pointer to distinguish between explicit zero and unspecified.Defaults to 1.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"selector": schema.SingleNestedAttribute{
						Description:         "Selector is a label query over machines that should match the replica count.Label keys and values that must match in order to be controlled by this MachineSet.It must match the machine template's labels.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors",
						MarkdownDescription: "Selector is a label query over machines that should match the replica count.Label keys and values that must match in order to be controlled by this MachineSet.It must match the machine template's labels.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors",
						Attributes: map[string]schema.Attribute{
							"match_expressions": schema.ListNestedAttribute{
								Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "key is the label key that the selector applies to.",
											MarkdownDescription: "key is the label key that the selector applies to.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"operator": schema.StringAttribute{
											Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
											MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"values": schema.ListAttribute{
											Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
											ElementType:         types.StringType,
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

							"match_labels": schema.MapAttribute{
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

					"template": schema.SingleNestedAttribute{
						Description:         "Template is the object that describes the machine that will be created ifinsufficient replicas are detected.Object references to custom resources are treated as templates.",
						MarkdownDescription: "Template is the object that describes the machine that will be created ifinsufficient replicas are detected.Object references to custom resources are treated as templates.",
						Attributes: map[string]schema.Attribute{
							"metadata": schema.SingleNestedAttribute{
								Description:         "Standard object's metadata.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
								MarkdownDescription: "Standard object's metadata.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Annotations is an unstructured key value map stored with a resource that may beset by external tools to store and retrieve arbitrary metadata. They are notqueryable and should be preserved when modifying objects.More info: http://kubernetes.io/docs/user-guide/annotations",
										MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may beset by external tools to store and retrieve arbitrary metadata. They are notqueryable and should be preserved when modifying objects.More info: http://kubernetes.io/docs/user-guide/annotations",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.MapAttribute{
										Description:         "Map of string keys and values that can be used to organize and categorize(scope and select) objects. May match selectors of replication controllersand services.More info: http://kubernetes.io/docs/user-guide/labels",
										MarkdownDescription: "Map of string keys and values that can be used to organize and categorize(scope and select) objects. May match selectors of replication controllersand services.More info: http://kubernetes.io/docs/user-guide/labels",
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
								Description:         "Specification of the desired behavior of the machine.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
								MarkdownDescription: "Specification of the desired behavior of the machine.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
								Attributes: map[string]schema.Attribute{
									"bootstrap": schema.SingleNestedAttribute{
										Description:         "Bootstrap is a reference to a local struct which encapsulatesfields to configure the Machine’s bootstrapping mechanism.",
										MarkdownDescription: "Bootstrap is a reference to a local struct which encapsulatesfields to configure the Machine’s bootstrapping mechanism.",
										Attributes: map[string]schema.Attribute{
											"config_ref": schema.SingleNestedAttribute{
												Description:         "ConfigRef is a reference to a bootstrap provider-specific resourcethat holds configuration details. The reference is optional toallow users/operators to specify Bootstrap.DataSecretName withoutthe need of a controller.",
												MarkdownDescription: "ConfigRef is a reference to a bootstrap provider-specific resourcethat holds configuration details. The reference is optional toallow users/operators to specify Bootstrap.DataSecretName withoutthe need of a controller.",
												Attributes: map[string]schema.Attribute{
													"api_version": schema.StringAttribute{
														Description:         "API version of the referent.",
														MarkdownDescription: "API version of the referent.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"field_path": schema.StringAttribute{
														Description:         "If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.",
														MarkdownDescription: "If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"kind": schema.StringAttribute{
														Description:         "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
														MarkdownDescription: "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
														MarkdownDescription: "Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"resource_version": schema.StringAttribute{
														Description:         "Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
														MarkdownDescription: "Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"uid": schema.StringAttribute{
														Description:         "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
														MarkdownDescription: "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
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
												Description:         "DataSecretName is the name of the secret that stores the bootstrap data script.If nil, the Machine should remain in the Pending state.",
												MarkdownDescription: "DataSecretName is the name of the secret that stores the bootstrap data script.If nil, the Machine should remain in the Pending state.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"cluster_name": schema.StringAttribute{
										Description:         "ClusterName is the name of the Cluster this object belongs to.",
										MarkdownDescription: "ClusterName is the name of the Cluster this object belongs to.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(1),
										},
									},

									"failure_domain": schema.StringAttribute{
										Description:         "FailureDomain is the failure domain the machine will be created in.Must match a key in the FailureDomains map stored on the cluster object.",
										MarkdownDescription: "FailureDomain is the failure domain the machine will be created in.Must match a key in the FailureDomains map stored on the cluster object.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"infrastructure_ref": schema.SingleNestedAttribute{
										Description:         "InfrastructureRef is a required reference to a custom resourceoffered by an infrastructure provider.",
										MarkdownDescription: "InfrastructureRef is a required reference to a custom resourceoffered by an infrastructure provider.",
										Attributes: map[string]schema.Attribute{
											"api_version": schema.StringAttribute{
												Description:         "API version of the referent.",
												MarkdownDescription: "API version of the referent.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"field_path": schema.StringAttribute{
												Description:         "If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.",
												MarkdownDescription: "If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"kind": schema.StringAttribute{
												Description:         "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
												MarkdownDescription: "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
												MarkdownDescription: "Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"resource_version": schema.StringAttribute{
												Description:         "Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
												MarkdownDescription: "Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"uid": schema.StringAttribute{
												Description:         "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
												MarkdownDescription: "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"node_drain_timeout": schema.StringAttribute{
										Description:         "NodeDrainTimeout is the total amount of time that the controller will spend on draining a node.The default value is 0, meaning that the node can be drained without any time limitations.NOTE: NodeDrainTimeout is different from 'kubectl drain --timeout'",
										MarkdownDescription: "NodeDrainTimeout is the total amount of time that the controller will spend on draining a node.The default value is 0, meaning that the node can be drained without any time limitations.NOTE: NodeDrainTimeout is different from 'kubectl drain --timeout'",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"provider_id": schema.StringAttribute{
										Description:         "ProviderID is the identification ID of the machine provided by the provider.This field must match the provider ID as seen on the node object corresponding to this machine.This field is required by higher level consumers of cluster-api. Example use case is cluster autoscalerwith cluster-api as provider. Clean-up logic in the autoscaler compares machines to nodes to find outmachines at provider which could not get registered as Kubernetes nodes. With cluster-api as ageneric out-of-tree provider for autoscaler, this field is required by autoscaler to beable to have a provider view of the list of machines. Another list of nodes is queried from the k8s apiserverand then a comparison is done to find out unregistered machines and are marked for delete.This field will be set by the actuators and consumed by higher level entities like autoscaler that willbe interfacing with cluster-api as generic provider.",
										MarkdownDescription: "ProviderID is the identification ID of the machine provided by the provider.This field must match the provider ID as seen on the node object corresponding to this machine.This field is required by higher level consumers of cluster-api. Example use case is cluster autoscalerwith cluster-api as provider. Clean-up logic in the autoscaler compares machines to nodes to find outmachines at provider which could not get registered as Kubernetes nodes. With cluster-api as ageneric out-of-tree provider for autoscaler, this field is required by autoscaler to beable to have a provider view of the list of machines. Another list of nodes is queried from the k8s apiserverand then a comparison is done to find out unregistered machines and are marked for delete.This field will be set by the actuators and consumed by higher level entities like autoscaler that willbe interfacing with cluster-api as generic provider.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"version": schema.StringAttribute{
										Description:         "Version defines the desired Kubernetes version.This field is meant to be optionally used by bootstrap providers.",
										MarkdownDescription: "Version defines the desired Kubernetes version.This field is meant to be optionally used by bootstrap providers.",
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *ClusterXK8SIoMachineSetV1Alpha4Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_cluster_x_k8s_io_machine_set_v1alpha4_manifest")

	var model ClusterXK8SIoMachineSetV1Alpha4ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("cluster.x-k8s.io/v1alpha4")
	model.Kind = pointer.String("MachineSet")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
