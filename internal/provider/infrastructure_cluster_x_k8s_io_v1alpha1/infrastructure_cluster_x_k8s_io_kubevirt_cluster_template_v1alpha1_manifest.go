/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package infrastructure_cluster_x_k8s_io_v1alpha1

import (
	"context"
	"fmt"
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
	_ datasource.DataSource = &InfrastructureClusterXK8SIoKubevirtClusterTemplateV1Alpha1Manifest{}
)

func NewInfrastructureClusterXK8SIoKubevirtClusterTemplateV1Alpha1Manifest() datasource.DataSource {
	return &InfrastructureClusterXK8SIoKubevirtClusterTemplateV1Alpha1Manifest{}
}

type InfrastructureClusterXK8SIoKubevirtClusterTemplateV1Alpha1Manifest struct{}

type InfrastructureClusterXK8SIoKubevirtClusterTemplateV1Alpha1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
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
		Template *struct {
			Metadata *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"metadata" json:"metadata,omitempty"`
			Spec *struct {
				ControlPlaneEndpoint *struct {
					Host *string `tfsdk:"host" json:"host,omitempty"`
					Port *int64  `tfsdk:"port" json:"port,omitempty"`
				} `tfsdk:"control_plane_endpoint" json:"controlPlaneEndpoint,omitempty"`
				ControlPlaneServiceTemplate *struct {
					Metadata *map[string]string `tfsdk:"metadata" json:"metadata,omitempty"`
					Spec     *struct {
						Type *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"spec" json:"spec,omitempty"`
				} `tfsdk:"control_plane_service_template" json:"controlPlaneServiceTemplate,omitempty"`
				InfraClusterSecretRef *struct {
					ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
					FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
					Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
					Name            *string `tfsdk:"name" json:"name,omitempty"`
					Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
					ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
					Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
				} `tfsdk:"infra_cluster_secret_ref" json:"infraClusterSecretRef,omitempty"`
				SshKeys *struct {
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
				} `tfsdk:"ssh_keys" json:"sshKeys,omitempty"`
			} `tfsdk:"spec" json:"spec,omitempty"`
		} `tfsdk:"template" json:"template,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *InfrastructureClusterXK8SIoKubevirtClusterTemplateV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_infrastructure_cluster_x_k8s_io_kubevirt_cluster_template_v1alpha1_manifest"
}

func (r *InfrastructureClusterXK8SIoKubevirtClusterTemplateV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "KubevirtClusterTemplate is the Schema for the kubevirtclustertemplates API.",
		MarkdownDescription: "KubevirtClusterTemplate is the Schema for the kubevirtclustertemplates API.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

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
				Description:         "KubevirtClusterTemplateSpec defines the desired state of KubevirtClusterTemplate.",
				MarkdownDescription: "KubevirtClusterTemplateSpec defines the desired state of KubevirtClusterTemplate.",
				Attributes: map[string]schema.Attribute{
					"template": schema.SingleNestedAttribute{
						Description:         "KubevirtClusterTemplateResource describes the data needed to create a KubevirtCluster from a template.",
						MarkdownDescription: "KubevirtClusterTemplateResource describes the data needed to create a KubevirtCluster from a template.",
						Attributes: map[string]schema.Attribute{
							"metadata": schema.SingleNestedAttribute{
								Description:         "ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create. This is a copy of customizable fields from metav1.ObjectMeta.  ObjectMeta is embedded in 'Machine.Spec', 'MachineDeployment.Template' and 'MachineSet.Template', which are not top-level Kubernetes objects. Given that metav1.ObjectMeta has lots of special cases and read-only fields which end up in the generated CRD validation, having it as a subset simplifies the API and some issues that can impact user experience.  During the [upgrade to controller-tools@v2](https://github.com/kubernetes-sigs/cluster-api/pull/1054) for v1alpha2, we noticed a failure would occur running Cluster API test suite against the new CRDs, specifically 'spec.metadata.creationTimestamp in body must be of type string: 'null''. The investigation showed that 'controller-tools@v2' behaves differently than its previous version when handling types from [metav1](k8s.io/apimachinery/pkg/apis/meta/v1) package.  In more details, we found that embedded (non-top level) types that embedded 'metav1.ObjectMeta' had validation properties, including for 'creationTimestamp' (metav1.Time). The 'metav1.Time' type specifies a custom json marshaller that, when IsZero() is true, returns 'null' which breaks validation because the field isn't marked as nullable.  In future versions, controller-tools@v2 might allow overriding the type and validation for embedded types. When that happens, this hack should be revisited.",
								MarkdownDescription: "ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create. This is a copy of customizable fields from metav1.ObjectMeta.  ObjectMeta is embedded in 'Machine.Spec', 'MachineDeployment.Template' and 'MachineSet.Template', which are not top-level Kubernetes objects. Given that metav1.ObjectMeta has lots of special cases and read-only fields which end up in the generated CRD validation, having it as a subset simplifies the API and some issues that can impact user experience.  During the [upgrade to controller-tools@v2](https://github.com/kubernetes-sigs/cluster-api/pull/1054) for v1alpha2, we noticed a failure would occur running Cluster API test suite against the new CRDs, specifically 'spec.metadata.creationTimestamp in body must be of type string: 'null''. The investigation showed that 'controller-tools@v2' behaves differently than its previous version when handling types from [metav1](k8s.io/apimachinery/pkg/apis/meta/v1) package.  In more details, we found that embedded (non-top level) types that embedded 'metav1.ObjectMeta' had validation properties, including for 'creationTimestamp' (metav1.Time). The 'metav1.Time' type specifies a custom json marshaller that, when IsZero() is true, returns 'null' which breaks validation because the field isn't marked as nullable.  In future versions, controller-tools@v2 might allow overriding the type and validation for embedded types. When that happens, this hack should be revisited.",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
										MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.MapAttribute{
										Description:         "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
										MarkdownDescription: "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
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
								Description:         "KubevirtClusterSpec defines the desired state of KubevirtCluster.",
								MarkdownDescription: "KubevirtClusterSpec defines the desired state of KubevirtCluster.",
								Attributes: map[string]schema.Attribute{
									"control_plane_endpoint": schema.SingleNestedAttribute{
										Description:         "ControlPlaneEndpoint represents the endpoint used to communicate with the control plane.",
										MarkdownDescription: "ControlPlaneEndpoint represents the endpoint used to communicate with the control plane.",
										Attributes: map[string]schema.Attribute{
											"host": schema.StringAttribute{
												Description:         "Host is the hostname on which the API server is serving.",
												MarkdownDescription: "Host is the hostname on which the API server is serving.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"port": schema.Int64Attribute{
												Description:         "Port is the port on which the API server is serving.",
												MarkdownDescription: "Port is the port on which the API server is serving.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"control_plane_service_template": schema.SingleNestedAttribute{
										Description:         "ControlPlaneServiceTemplate can be used to modify service that fronts the control plane nodes to handle the api-server traffic (port 6443). This field is optional, by default control plane nodes will use a service of type ClusterIP, which will make workload cluster only accessible within the same cluster. Note, this does not aim to expose the entire Service spec to users, but only provides capability to modify the service metadata and the service type.",
										MarkdownDescription: "ControlPlaneServiceTemplate can be used to modify service that fronts the control plane nodes to handle the api-server traffic (port 6443). This field is optional, by default control plane nodes will use a service of type ClusterIP, which will make workload cluster only accessible within the same cluster. Note, this does not aim to expose the entire Service spec to users, but only provides capability to modify the service metadata and the service type.",
										Attributes: map[string]schema.Attribute{
											"metadata": schema.MapAttribute{
												Description:         "Service metadata allows to set labels, annotations and namespace for the service. When infraClusterSecretRef is used, ControlPlaneService take the kubeconfig namespace by default if metadata.namespace is not specified. This field is optional.",
												MarkdownDescription: "Service metadata allows to set labels, annotations and namespace for the service. When infraClusterSecretRef is used, ControlPlaneService take the kubeconfig namespace by default if metadata.namespace is not specified. This field is optional.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"spec": schema.SingleNestedAttribute{
												Description:         "Service specification allows to override some fields in the service spec. Note, it does not aim cover all fields of the service spec.",
												MarkdownDescription: "Service specification allows to override some fields in the service spec. Note, it does not aim cover all fields of the service spec.",
												Attributes: map[string]schema.Attribute{
													"type": schema.StringAttribute{
														Description:         "Type determines how the Service is exposed. Defaults to ClusterIP. Valid options are ExternalName, ClusterIP, NodePort, and LoadBalancer. More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types",
														MarkdownDescription: "Type determines how the Service is exposed. Defaults to ClusterIP. Valid options are ExternalName, ClusterIP, NodePort, and LoadBalancer. More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types",
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

									"infra_cluster_secret_ref": schema.SingleNestedAttribute{
										Description:         "InfraClusterSecretRef is a reference to a secret with a kubeconfig for external cluster used for infra.",
										MarkdownDescription: "InfraClusterSecretRef is a reference to a secret with a kubeconfig for external cluster used for infra.",
										Attributes: map[string]schema.Attribute{
											"api_version": schema.StringAttribute{
												Description:         "API version of the referent.",
												MarkdownDescription: "API version of the referent.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"field_path": schema.StringAttribute{
												Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
												MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
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

									"ssh_keys": schema.SingleNestedAttribute{
										Description:         "SSHKeys is a reference to a local struct for SSH keys persistence.",
										MarkdownDescription: "SSHKeys is a reference to a local struct for SSH keys persistence.",
										Attributes: map[string]schema.Attribute{
											"config_ref": schema.SingleNestedAttribute{
												Description:         "ConfigRef is a reference to a resource containing the keys. The reference is optional to allow users/operators to specify Bootstrap.DataSecretName without the need of a controller.",
												MarkdownDescription: "ConfigRef is a reference to a resource containing the keys. The reference is optional to allow users/operators to specify Bootstrap.DataSecretName without the need of a controller.",
												Attributes: map[string]schema.Attribute{
													"api_version": schema.StringAttribute{
														Description:         "API version of the referent.",
														MarkdownDescription: "API version of the referent.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"field_path": schema.StringAttribute{
														Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
														MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
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
												Description:         "DataSecretName is the name of the secret that stores ssh keys.",
												MarkdownDescription: "DataSecretName is the name of the secret that stores ssh keys.",
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
								Required: true,
								Optional: false,
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

func (r *InfrastructureClusterXK8SIoKubevirtClusterTemplateV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_infrastructure_cluster_x_k8s_io_kubevirt_cluster_template_v1alpha1_manifest")

	var model InfrastructureClusterXK8SIoKubevirtClusterTemplateV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("infrastructure.cluster.x-k8s.io/v1alpha1")
	model.Kind = pointer.String("KubevirtClusterTemplate")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
