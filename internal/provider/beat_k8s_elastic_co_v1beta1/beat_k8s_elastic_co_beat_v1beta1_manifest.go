/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package beat_k8s_elastic_co_v1beta1

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
	_ datasource.DataSource = &BeatK8SElasticCoBeatV1Beta1Manifest{}
)

func NewBeatK8SElasticCoBeatV1Beta1Manifest() datasource.DataSource {
	return &BeatK8SElasticCoBeatV1Beta1Manifest{}
}

type BeatK8SElasticCoBeatV1Beta1Manifest struct{}

type BeatK8SElasticCoBeatV1Beta1ManifestData struct {
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
		Config    *map[string]string `tfsdk:"config" json:"config,omitempty"`
		ConfigRef *struct {
			SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
		} `tfsdk:"config_ref" json:"configRef,omitempty"`
		DaemonSet *struct {
			PodTemplate    *map[string]string `tfsdk:"pod_template" json:"podTemplate,omitempty"`
			UpdateStrategy *struct {
				RollingUpdate *struct {
					MaxSurge       *string `tfsdk:"max_surge" json:"maxSurge,omitempty"`
					MaxUnavailable *string `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
				} `tfsdk:"rolling_update" json:"rollingUpdate,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"update_strategy" json:"updateStrategy,omitempty"`
		} `tfsdk:"daemon_set" json:"daemonSet,omitempty"`
		Deployment *struct {
			PodTemplate *map[string]string `tfsdk:"pod_template" json:"podTemplate,omitempty"`
			Replicas    *int64             `tfsdk:"replicas" json:"replicas,omitempty"`
			Strategy    *struct {
				RollingUpdate *struct {
					MaxSurge       *string `tfsdk:"max_surge" json:"maxSurge,omitempty"`
					MaxUnavailable *string `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
				} `tfsdk:"rolling_update" json:"rollingUpdate,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"strategy" json:"strategy,omitempty"`
		} `tfsdk:"deployment" json:"deployment,omitempty"`
		ElasticsearchRef *struct {
			Name        *string `tfsdk:"name" json:"name,omitempty"`
			Namespace   *string `tfsdk:"namespace" json:"namespace,omitempty"`
			SecretName  *string `tfsdk:"secret_name" json:"secretName,omitempty"`
			ServiceName *string `tfsdk:"service_name" json:"serviceName,omitempty"`
		} `tfsdk:"elasticsearch_ref" json:"elasticsearchRef,omitempty"`
		Image     *string `tfsdk:"image" json:"image,omitempty"`
		KibanaRef *struct {
			Name        *string `tfsdk:"name" json:"name,omitempty"`
			Namespace   *string `tfsdk:"namespace" json:"namespace,omitempty"`
			SecretName  *string `tfsdk:"secret_name" json:"secretName,omitempty"`
			ServiceName *string `tfsdk:"service_name" json:"serviceName,omitempty"`
		} `tfsdk:"kibana_ref" json:"kibanaRef,omitempty"`
		Monitoring *struct {
			Logs *struct {
				ElasticsearchRefs *[]struct {
					Name        *string `tfsdk:"name" json:"name,omitempty"`
					Namespace   *string `tfsdk:"namespace" json:"namespace,omitempty"`
					SecretName  *string `tfsdk:"secret_name" json:"secretName,omitempty"`
					ServiceName *string `tfsdk:"service_name" json:"serviceName,omitempty"`
				} `tfsdk:"elasticsearch_refs" json:"elasticsearchRefs,omitempty"`
			} `tfsdk:"logs" json:"logs,omitempty"`
			Metrics *struct {
				ElasticsearchRefs *[]struct {
					Name        *string `tfsdk:"name" json:"name,omitempty"`
					Namespace   *string `tfsdk:"namespace" json:"namespace,omitempty"`
					SecretName  *string `tfsdk:"secret_name" json:"secretName,omitempty"`
					ServiceName *string `tfsdk:"service_name" json:"serviceName,omitempty"`
				} `tfsdk:"elasticsearch_refs" json:"elasticsearchRefs,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
		} `tfsdk:"monitoring" json:"monitoring,omitempty"`
		RevisionHistoryLimit *int64 `tfsdk:"revision_history_limit" json:"revisionHistoryLimit,omitempty"`
		SecureSettings       *[]struct {
			Entries *[]struct {
				Key  *string `tfsdk:"key" json:"key,omitempty"`
				Path *string `tfsdk:"path" json:"path,omitempty"`
			} `tfsdk:"entries" json:"entries,omitempty"`
			SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
		} `tfsdk:"secure_settings" json:"secureSettings,omitempty"`
		ServiceAccountName *string `tfsdk:"service_account_name" json:"serviceAccountName,omitempty"`
		Type               *string `tfsdk:"type" json:"type,omitempty"`
		Version            *string `tfsdk:"version" json:"version,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *BeatK8SElasticCoBeatV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_beat_k8s_elastic_co_beat_v1beta1_manifest"
}

func (r *BeatK8SElasticCoBeatV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Beat is the Schema for the Beats API.",
		MarkdownDescription: "Beat is the Schema for the Beats API.",
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
				Description:         "BeatSpec defines the desired state of a Beat.",
				MarkdownDescription: "BeatSpec defines the desired state of a Beat.",
				Attributes: map[string]schema.Attribute{
					"config": schema.MapAttribute{
						Description:         "Config holds the Beat configuration. At most one of ['Config', 'ConfigRef'] can be specified.",
						MarkdownDescription: "Config holds the Beat configuration. At most one of ['Config', 'ConfigRef'] can be specified.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"config_ref": schema.SingleNestedAttribute{
						Description:         "ConfigRef contains a reference to an existing Kubernetes Secret holding the Beat configuration. Beat settings must be specified as yaml, under a single 'beat.yml' entry. At most one of ['Config', 'ConfigRef'] can be specified.",
						MarkdownDescription: "ConfigRef contains a reference to an existing Kubernetes Secret holding the Beat configuration. Beat settings must be specified as yaml, under a single 'beat.yml' entry. At most one of ['Config', 'ConfigRef'] can be specified.",
						Attributes: map[string]schema.Attribute{
							"secret_name": schema.StringAttribute{
								Description:         "SecretName is the name of the secret.",
								MarkdownDescription: "SecretName is the name of the secret.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"daemon_set": schema.SingleNestedAttribute{
						Description:         "DaemonSet specifies the Beat should be deployed as a DaemonSet, and allows providing its spec. Cannot be used along with 'deployment'. If both are absent a default for the Type is used.",
						MarkdownDescription: "DaemonSet specifies the Beat should be deployed as a DaemonSet, and allows providing its spec. Cannot be used along with 'deployment'. If both are absent a default for the Type is used.",
						Attributes: map[string]schema.Attribute{
							"pod_template": schema.MapAttribute{
								Description:         "PodTemplateSpec describes the data a pod should have when created from a template",
								MarkdownDescription: "PodTemplateSpec describes the data a pod should have when created from a template",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"update_strategy": schema.SingleNestedAttribute{
								Description:         "DaemonSetUpdateStrategy is a struct used to control the update strategy for a DaemonSet.",
								MarkdownDescription: "DaemonSetUpdateStrategy is a struct used to control the update strategy for a DaemonSet.",
								Attributes: map[string]schema.Attribute{
									"rolling_update": schema.SingleNestedAttribute{
										Description:         "Rolling update config params. Present only if type = 'RollingUpdate'.",
										MarkdownDescription: "Rolling update config params. Present only if type = 'RollingUpdate'.",
										Attributes: map[string]schema.Attribute{
											"max_surge": schema.StringAttribute{
												Description:         "The maximum number of nodes with an existing available DaemonSet pod that can have an updated DaemonSet pod during during an update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up to a minimum of 1. Default value is 0. Example: when this is set to 30%, at most 30% of the total number of nodes that should be running the daemon pod (i.e. status.desiredNumberScheduled) can have their a new pod created before the old pod is marked as deleted. The update starts by launching new pods on 30% of nodes. Once an updated pod is available (Ready for at least minReadySeconds) the old DaemonSet pod on that node is marked deleted. If the old pod becomes unavailable for any reason (Ready transitions to false, is evicted, or is drained) an updated pod is immediatedly created on that node without considering surge limits. Allowing surge implies the possibility that the resources consumed by the daemonset on any given node can double if the readiness check fails, and so resource intensive daemonsets should take into account that they may cause evictions during disruption.",
												MarkdownDescription: "The maximum number of nodes with an existing available DaemonSet pod that can have an updated DaemonSet pod during during an update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up to a minimum of 1. Default value is 0. Example: when this is set to 30%, at most 30% of the total number of nodes that should be running the daemon pod (i.e. status.desiredNumberScheduled) can have their a new pod created before the old pod is marked as deleted. The update starts by launching new pods on 30% of nodes. Once an updated pod is available (Ready for at least minReadySeconds) the old DaemonSet pod on that node is marked deleted. If the old pod becomes unavailable for any reason (Ready transitions to false, is evicted, or is drained) an updated pod is immediatedly created on that node without considering surge limits. Allowing surge implies the possibility that the resources consumed by the daemonset on any given node can double if the readiness check fails, and so resource intensive daemonsets should take into account that they may cause evictions during disruption.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_unavailable": schema.StringAttribute{
												Description:         "The maximum number of DaemonSet pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of total number of DaemonSet pods at the start of the update (ex: 10%). Absolute number is calculated from percentage by rounding up. This cannot be 0 if MaxSurge is 0 Default value is 1. Example: when this is set to 30%, at most 30% of the total number of nodes that should be running the daemon pod (i.e. status.desiredNumberScheduled) can have their pods stopped for an update at any given time. The update starts by stopping at most 30% of those DaemonSet pods and then brings up new DaemonSet pods in their place. Once the new pods are available, it then proceeds onto other DaemonSet pods, thus ensuring that at least 70% of original number of DaemonSet pods are available at all times during the update.",
												MarkdownDescription: "The maximum number of DaemonSet pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of total number of DaemonSet pods at the start of the update (ex: 10%). Absolute number is calculated from percentage by rounding up. This cannot be 0 if MaxSurge is 0 Default value is 1. Example: when this is set to 30%, at most 30% of the total number of nodes that should be running the daemon pod (i.e. status.desiredNumberScheduled) can have their pods stopped for an update at any given time. The update starts by stopping at most 30% of those DaemonSet pods and then brings up new DaemonSet pods in their place. Once the new pods are available, it then proceeds onto other DaemonSet pods, thus ensuring that at least 70% of original number of DaemonSet pods are available at all times during the update.",
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
										Description:         "Type of daemon set update. Can be 'RollingUpdate' or 'OnDelete'. Default is RollingUpdate.",
										MarkdownDescription: "Type of daemon set update. Can be 'RollingUpdate' or 'OnDelete'. Default is RollingUpdate.",
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

					"deployment": schema.SingleNestedAttribute{
						Description:         "Deployment specifies the Beat should be deployed as a Deployment, and allows providing its spec. Cannot be used along with 'daemonSet'. If both are absent a default for the Type is used.",
						MarkdownDescription: "Deployment specifies the Beat should be deployed as a Deployment, and allows providing its spec. Cannot be used along with 'daemonSet'. If both are absent a default for the Type is used.",
						Attributes: map[string]schema.Attribute{
							"pod_template": schema.MapAttribute{
								Description:         "PodTemplateSpec describes the data a pod should have when created from a template",
								MarkdownDescription: "PodTemplateSpec describes the data a pod should have when created from a template",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"replicas": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"strategy": schema.SingleNestedAttribute{
								Description:         "DeploymentStrategy describes how to replace existing pods with new ones.",
								MarkdownDescription: "DeploymentStrategy describes how to replace existing pods with new ones.",
								Attributes: map[string]schema.Attribute{
									"rolling_update": schema.SingleNestedAttribute{
										Description:         "Rolling update config params. Present only if DeploymentStrategyType = RollingUpdate.",
										MarkdownDescription: "Rolling update config params. Present only if DeploymentStrategyType = RollingUpdate.",
										Attributes: map[string]schema.Attribute{
											"max_surge": schema.StringAttribute{
												Description:         "The maximum number of pods that can be scheduled above the desired number of pods. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up. Defaults to 25%. Example: when this is set to 30%, the new ReplicaSet can be scaled up immediately when the rolling update starts, such that the total number of old and new pods do not exceed 130% of desired pods. Once old pods have been killed, new ReplicaSet can be scaled up further, ensuring that total number of pods running at any time during the update is at most 130% of desired pods.",
												MarkdownDescription: "The maximum number of pods that can be scheduled above the desired number of pods. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up. Defaults to 25%. Example: when this is set to 30%, the new ReplicaSet can be scaled up immediately when the rolling update starts, such that the total number of old and new pods do not exceed 130% of desired pods. Once old pods have been killed, new ReplicaSet can be scaled up further, ensuring that total number of pods running at any time during the update is at most 130% of desired pods.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_unavailable": schema.StringAttribute{
												Description:         "The maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). Absolute number is calculated from percentage by rounding down. This can not be 0 if MaxSurge is 0. Defaults to 25%. Example: when this is set to 30%, the old ReplicaSet can be scaled down to 70% of desired pods immediately when the rolling update starts. Once new pods are ready, old ReplicaSet can be scaled down further, followed by scaling up the new ReplicaSet, ensuring that the total number of pods available at all times during the update is at least 70% of desired pods.",
												MarkdownDescription: "The maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). Absolute number is calculated from percentage by rounding down. This can not be 0 if MaxSurge is 0. Defaults to 25%. Example: when this is set to 30%, the old ReplicaSet can be scaled down to 70% of desired pods immediately when the rolling update starts. Once new pods are ready, old ReplicaSet can be scaled down further, followed by scaling up the new ReplicaSet, ensuring that the total number of pods available at all times during the update is at least 70% of desired pods.",
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
										Description:         "Type of deployment. Can be 'Recreate' or 'RollingUpdate'. Default is RollingUpdate.",
										MarkdownDescription: "Type of deployment. Can be 'Recreate' or 'RollingUpdate'. Default is RollingUpdate.",
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

					"elasticsearch_ref": schema.SingleNestedAttribute{
						Description:         "ElasticsearchRef is a reference to an Elasticsearch cluster running in the same Kubernetes cluster.",
						MarkdownDescription: "ElasticsearchRef is a reference to an Elasticsearch cluster running in the same Kubernetes cluster.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of an existing Kubernetes object corresponding to an Elastic resource managed by ECK.",
								MarkdownDescription: "Name of an existing Kubernetes object corresponding to an Elastic resource managed by ECK.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace of the Kubernetes object. If empty, defaults to the current namespace.",
								MarkdownDescription: "Namespace of the Kubernetes object. If empty, defaults to the current namespace.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secret_name": schema.StringAttribute{
								Description:         "SecretName is the name of an existing Kubernetes secret that contains connection information for associating an Elastic resource not managed by the operator. The referenced secret must contain the following: - 'url': the URL to reach the Elastic resource - 'username': the username of the user to be authenticated to the Elastic resource - 'password': the password of the user to be authenticated to the Elastic resource - 'ca.crt': the CA certificate in PEM format (optional) - 'api-key': the key to authenticate against the Elastic resource instead of a username and password (supported only for 'elasticsearchRefs' in AgentSpec and in BeatSpec) This field cannot be used in combination with the other fields name, namespace or serviceName.",
								MarkdownDescription: "SecretName is the name of an existing Kubernetes secret that contains connection information for associating an Elastic resource not managed by the operator. The referenced secret must contain the following: - 'url': the URL to reach the Elastic resource - 'username': the username of the user to be authenticated to the Elastic resource - 'password': the password of the user to be authenticated to the Elastic resource - 'ca.crt': the CA certificate in PEM format (optional) - 'api-key': the key to authenticate against the Elastic resource instead of a username and password (supported only for 'elasticsearchRefs' in AgentSpec and in BeatSpec) This field cannot be used in combination with the other fields name, namespace or serviceName.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_name": schema.StringAttribute{
								Description:         "ServiceName is the name of an existing Kubernetes service which is used to make requests to the referenced object. It has to be in the same namespace as the referenced resource. If left empty, the default HTTP service of the referenced resource is used.",
								MarkdownDescription: "ServiceName is the name of an existing Kubernetes service which is used to make requests to the referenced object. It has to be in the same namespace as the referenced resource. If left empty, the default HTTP service of the referenced resource is used.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"image": schema.StringAttribute{
						Description:         "Image is the Beat Docker image to deploy. Version and Type have to match the Beat in the image.",
						MarkdownDescription: "Image is the Beat Docker image to deploy. Version and Type have to match the Beat in the image.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"kibana_ref": schema.SingleNestedAttribute{
						Description:         "KibanaRef is a reference to a Kibana instance running in the same Kubernetes cluster. It allows automatic setup of dashboards and visualizations.",
						MarkdownDescription: "KibanaRef is a reference to a Kibana instance running in the same Kubernetes cluster. It allows automatic setup of dashboards and visualizations.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of an existing Kubernetes object corresponding to an Elastic resource managed by ECK.",
								MarkdownDescription: "Name of an existing Kubernetes object corresponding to an Elastic resource managed by ECK.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace of the Kubernetes object. If empty, defaults to the current namespace.",
								MarkdownDescription: "Namespace of the Kubernetes object. If empty, defaults to the current namespace.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secret_name": schema.StringAttribute{
								Description:         "SecretName is the name of an existing Kubernetes secret that contains connection information for associating an Elastic resource not managed by the operator. The referenced secret must contain the following: - 'url': the URL to reach the Elastic resource - 'username': the username of the user to be authenticated to the Elastic resource - 'password': the password of the user to be authenticated to the Elastic resource - 'ca.crt': the CA certificate in PEM format (optional) - 'api-key': the key to authenticate against the Elastic resource instead of a username and password (supported only for 'elasticsearchRefs' in AgentSpec and in BeatSpec) This field cannot be used in combination with the other fields name, namespace or serviceName.",
								MarkdownDescription: "SecretName is the name of an existing Kubernetes secret that contains connection information for associating an Elastic resource not managed by the operator. The referenced secret must contain the following: - 'url': the URL to reach the Elastic resource - 'username': the username of the user to be authenticated to the Elastic resource - 'password': the password of the user to be authenticated to the Elastic resource - 'ca.crt': the CA certificate in PEM format (optional) - 'api-key': the key to authenticate against the Elastic resource instead of a username and password (supported only for 'elasticsearchRefs' in AgentSpec and in BeatSpec) This field cannot be used in combination with the other fields name, namespace or serviceName.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_name": schema.StringAttribute{
								Description:         "ServiceName is the name of an existing Kubernetes service which is used to make requests to the referenced object. It has to be in the same namespace as the referenced resource. If left empty, the default HTTP service of the referenced resource is used.",
								MarkdownDescription: "ServiceName is the name of an existing Kubernetes service which is used to make requests to the referenced object. It has to be in the same namespace as the referenced resource. If left empty, the default HTTP service of the referenced resource is used.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"monitoring": schema.SingleNestedAttribute{
						Description:         "Monitoring enables you to collect and ship logs and metrics for this Beat. Metricbeat and/or Filebeat sidecars are configured and send monitoring data to an Elasticsearch monitoring cluster running in the same Kubernetes cluster.",
						MarkdownDescription: "Monitoring enables you to collect and ship logs and metrics for this Beat. Metricbeat and/or Filebeat sidecars are configured and send monitoring data to an Elasticsearch monitoring cluster running in the same Kubernetes cluster.",
						Attributes: map[string]schema.Attribute{
							"logs": schema.SingleNestedAttribute{
								Description:         "Logs holds references to Elasticsearch clusters which receive log data from an associated resource.",
								MarkdownDescription: "Logs holds references to Elasticsearch clusters which receive log data from an associated resource.",
								Attributes: map[string]schema.Attribute{
									"elasticsearch_refs": schema.ListNestedAttribute{
										Description:         "ElasticsearchRefs is a reference to a list of monitoring Elasticsearch clusters running in the same Kubernetes cluster. Due to existing limitations, only a single Elasticsearch cluster is currently supported.",
										MarkdownDescription: "ElasticsearchRefs is a reference to a list of monitoring Elasticsearch clusters running in the same Kubernetes cluster. Due to existing limitations, only a single Elasticsearch cluster is currently supported.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of an existing Kubernetes object corresponding to an Elastic resource managed by ECK.",
													MarkdownDescription: "Name of an existing Kubernetes object corresponding to an Elastic resource managed by ECK.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespace": schema.StringAttribute{
													Description:         "Namespace of the Kubernetes object. If empty, defaults to the current namespace.",
													MarkdownDescription: "Namespace of the Kubernetes object. If empty, defaults to the current namespace.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"secret_name": schema.StringAttribute{
													Description:         "SecretName is the name of an existing Kubernetes secret that contains connection information for associating an Elastic resource not managed by the operator. The referenced secret must contain the following: - 'url': the URL to reach the Elastic resource - 'username': the username of the user to be authenticated to the Elastic resource - 'password': the password of the user to be authenticated to the Elastic resource - 'ca.crt': the CA certificate in PEM format (optional) - 'api-key': the key to authenticate against the Elastic resource instead of a username and password (supported only for 'elasticsearchRefs' in AgentSpec and in BeatSpec) This field cannot be used in combination with the other fields name, namespace or serviceName.",
													MarkdownDescription: "SecretName is the name of an existing Kubernetes secret that contains connection information for associating an Elastic resource not managed by the operator. The referenced secret must contain the following: - 'url': the URL to reach the Elastic resource - 'username': the username of the user to be authenticated to the Elastic resource - 'password': the password of the user to be authenticated to the Elastic resource - 'ca.crt': the CA certificate in PEM format (optional) - 'api-key': the key to authenticate against the Elastic resource instead of a username and password (supported only for 'elasticsearchRefs' in AgentSpec and in BeatSpec) This field cannot be used in combination with the other fields name, namespace or serviceName.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"service_name": schema.StringAttribute{
													Description:         "ServiceName is the name of an existing Kubernetes service which is used to make requests to the referenced object. It has to be in the same namespace as the referenced resource. If left empty, the default HTTP service of the referenced resource is used.",
													MarkdownDescription: "ServiceName is the name of an existing Kubernetes service which is used to make requests to the referenced object. It has to be in the same namespace as the referenced resource. If left empty, the default HTTP service of the referenced resource is used.",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"metrics": schema.SingleNestedAttribute{
								Description:         "Metrics holds references to Elasticsearch clusters which receive monitoring data from this resource.",
								MarkdownDescription: "Metrics holds references to Elasticsearch clusters which receive monitoring data from this resource.",
								Attributes: map[string]schema.Attribute{
									"elasticsearch_refs": schema.ListNestedAttribute{
										Description:         "ElasticsearchRefs is a reference to a list of monitoring Elasticsearch clusters running in the same Kubernetes cluster. Due to existing limitations, only a single Elasticsearch cluster is currently supported.",
										MarkdownDescription: "ElasticsearchRefs is a reference to a list of monitoring Elasticsearch clusters running in the same Kubernetes cluster. Due to existing limitations, only a single Elasticsearch cluster is currently supported.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of an existing Kubernetes object corresponding to an Elastic resource managed by ECK.",
													MarkdownDescription: "Name of an existing Kubernetes object corresponding to an Elastic resource managed by ECK.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespace": schema.StringAttribute{
													Description:         "Namespace of the Kubernetes object. If empty, defaults to the current namespace.",
													MarkdownDescription: "Namespace of the Kubernetes object. If empty, defaults to the current namespace.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"secret_name": schema.StringAttribute{
													Description:         "SecretName is the name of an existing Kubernetes secret that contains connection information for associating an Elastic resource not managed by the operator. The referenced secret must contain the following: - 'url': the URL to reach the Elastic resource - 'username': the username of the user to be authenticated to the Elastic resource - 'password': the password of the user to be authenticated to the Elastic resource - 'ca.crt': the CA certificate in PEM format (optional) - 'api-key': the key to authenticate against the Elastic resource instead of a username and password (supported only for 'elasticsearchRefs' in AgentSpec and in BeatSpec) This field cannot be used in combination with the other fields name, namespace or serviceName.",
													MarkdownDescription: "SecretName is the name of an existing Kubernetes secret that contains connection information for associating an Elastic resource not managed by the operator. The referenced secret must contain the following: - 'url': the URL to reach the Elastic resource - 'username': the username of the user to be authenticated to the Elastic resource - 'password': the password of the user to be authenticated to the Elastic resource - 'ca.crt': the CA certificate in PEM format (optional) - 'api-key': the key to authenticate against the Elastic resource instead of a username and password (supported only for 'elasticsearchRefs' in AgentSpec and in BeatSpec) This field cannot be used in combination with the other fields name, namespace or serviceName.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"service_name": schema.StringAttribute{
													Description:         "ServiceName is the name of an existing Kubernetes service which is used to make requests to the referenced object. It has to be in the same namespace as the referenced resource. If left empty, the default HTTP service of the referenced resource is used.",
													MarkdownDescription: "ServiceName is the name of an existing Kubernetes service which is used to make requests to the referenced object. It has to be in the same namespace as the referenced resource. If left empty, the default HTTP service of the referenced resource is used.",
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

					"revision_history_limit": schema.Int64Attribute{
						Description:         "RevisionHistoryLimit is the number of revisions to retain to allow rollback in the underlying DaemonSet or Deployment.",
						MarkdownDescription: "RevisionHistoryLimit is the number of revisions to retain to allow rollback in the underlying DaemonSet or Deployment.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"secure_settings": schema.ListNestedAttribute{
						Description:         "SecureSettings is a list of references to Kubernetes Secrets containing sensitive configuration options for the Beat. Secrets data can be then referenced in the Beat config using the Secret's keys or as specified in 'Entries' field of each SecureSetting.",
						MarkdownDescription: "SecureSettings is a list of references to Kubernetes Secrets containing sensitive configuration options for the Beat. Secrets data can be then referenced in the Beat config using the Secret's keys or as specified in 'Entries' field of each SecureSetting.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"entries": schema.ListNestedAttribute{
									Description:         "Entries define how to project each key-value pair in the secret to filesystem paths. If not defined, all keys will be projected to similarly named paths in the filesystem. If defined, only the specified keys will be projected to the corresponding paths.",
									MarkdownDescription: "Entries define how to project each key-value pair in the secret to filesystem paths. If not defined, all keys will be projected to similarly named paths in the filesystem. If defined, only the specified keys will be projected to the corresponding paths.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "Key is the key contained in the secret.",
												MarkdownDescription: "Key is the key contained in the secret.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"path": schema.StringAttribute{
												Description:         "Path is the relative file path to map the key to. Path must not be an absolute file path and must not contain any '..' components.",
												MarkdownDescription: "Path is the relative file path to map the key to. Path must not be an absolute file path and must not contain any '..' components.",
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

								"secret_name": schema.StringAttribute{
									Description:         "SecretName is the name of the secret.",
									MarkdownDescription: "SecretName is the name of the secret.",
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

					"service_account_name": schema.StringAttribute{
						Description:         "ServiceAccountName is used to check access from the current resource to Elasticsearch resource in a different namespace. Can only be used if ECK is enforcing RBAC on references.",
						MarkdownDescription: "ServiceAccountName is used to check access from the current resource to Elasticsearch resource in a different namespace. Can only be used if ECK is enforcing RBAC on references.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"type": schema.StringAttribute{
						Description:         "Type is the type of the Beat to deploy (filebeat, metricbeat, heartbeat, auditbeat, journalbeat, packetbeat, and so on). Any string can be used, but well-known types will have the image field defaulted and have the appropriate Elasticsearch roles created automatically. It also allows for dashboard setup when combined with a 'KibanaRef'.",
						MarkdownDescription: "Type is the type of the Beat to deploy (filebeat, metricbeat, heartbeat, auditbeat, journalbeat, packetbeat, and so on). Any string can be used, but well-known types will have the image field defaulted and have the appropriate Elasticsearch roles created automatically. It also allows for dashboard setup when combined with a 'KibanaRef'.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(20),
							stringvalidator.RegexMatches(regexp.MustCompile(`[a-zA-Z0-9-]+`), ""),
						},
					},

					"version": schema.StringAttribute{
						Description:         "Version of the Beat.",
						MarkdownDescription: "Version of the Beat.",
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
	}
}

func (r *BeatK8SElasticCoBeatV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_beat_k8s_elastic_co_beat_v1beta1_manifest")

	var model BeatK8SElasticCoBeatV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("beat.k8s.elastic.co/v1beta1")
	model.Kind = pointer.String("Beat")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
