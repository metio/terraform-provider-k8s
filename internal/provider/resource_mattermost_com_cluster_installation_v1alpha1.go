/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"gopkg.in/yaml.v3"
	"time"
)

type MattermostComClusterInstallationV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*MattermostComClusterInstallationV1Alpha1Resource)(nil)
)

type MattermostComClusterInstallationV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type MattermostComClusterInstallationV1Alpha1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		Affinity *struct {
			NodeAffinity *struct {
				PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
					Preference *struct {
						MatchExpressions *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
						} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

						MatchFields *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
						} `tfsdk:"match_fields" yaml:"matchFields,omitempty"`
					} `tfsdk:"preference" yaml:"preference,omitempty"`

					Weight *int64 `tfsdk:"weight" yaml:"weight,omitempty"`
				} `tfsdk:"preferred_during_scheduling_ignored_during_execution" yaml:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`

				RequiredDuringSchedulingIgnoredDuringExecution *struct {
					NodeSelectorTerms *[]struct {
						MatchExpressions *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
						} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

						MatchFields *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
						} `tfsdk:"match_fields" yaml:"matchFields,omitempty"`
					} `tfsdk:"node_selector_terms" yaml:"nodeSelectorTerms,omitempty"`
				} `tfsdk:"required_during_scheduling_ignored_during_execution" yaml:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
			} `tfsdk:"node_affinity" yaml:"nodeAffinity,omitempty"`

			PodAffinity *struct {
				PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
					PodAffinityTerm *struct {
						LabelSelector *struct {
							MatchExpressions *[]struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

								Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
							} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

							MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
						} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

						NamespaceSelector *struct {
							MatchExpressions *[]struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

								Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
							} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

							MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
						} `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`

						Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

						TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`
					} `tfsdk:"pod_affinity_term" yaml:"podAffinityTerm,omitempty"`

					Weight *int64 `tfsdk:"weight" yaml:"weight,omitempty"`
				} `tfsdk:"preferred_during_scheduling_ignored_during_execution" yaml:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`

				RequiredDuringSchedulingIgnoredDuringExecution *[]struct {
					LabelSelector *struct {
						MatchExpressions *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
						} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

						MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
					} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

					NamespaceSelector *struct {
						MatchExpressions *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
						} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

						MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
					} `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`

					Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

					TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`
				} `tfsdk:"required_during_scheduling_ignored_during_execution" yaml:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
			} `tfsdk:"pod_affinity" yaml:"podAffinity,omitempty"`

			PodAntiAffinity *struct {
				PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
					PodAffinityTerm *struct {
						LabelSelector *struct {
							MatchExpressions *[]struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

								Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
							} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

							MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
						} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

						NamespaceSelector *struct {
							MatchExpressions *[]struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

								Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
							} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

							MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
						} `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`

						Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

						TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`
					} `tfsdk:"pod_affinity_term" yaml:"podAffinityTerm,omitempty"`

					Weight *int64 `tfsdk:"weight" yaml:"weight,omitempty"`
				} `tfsdk:"preferred_during_scheduling_ignored_during_execution" yaml:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`

				RequiredDuringSchedulingIgnoredDuringExecution *[]struct {
					LabelSelector *struct {
						MatchExpressions *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
						} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

						MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
					} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

					NamespaceSelector *struct {
						MatchExpressions *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
						} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

						MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
					} `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`

					Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

					TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`
				} `tfsdk:"required_during_scheduling_ignored_during_execution" yaml:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
			} `tfsdk:"pod_anti_affinity" yaml:"podAntiAffinity,omitempty"`
		} `tfsdk:"affinity" yaml:"affinity,omitempty"`

		BlueGreen *struct {
			Blue *struct {
				Image *string `tfsdk:"image" yaml:"image,omitempty"`

				IngressName *string `tfsdk:"ingress_name" yaml:"ingressName,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				ResourceLabels *map[string]string `tfsdk:"resource_labels" yaml:"resourceLabels,omitempty"`

				Version *string `tfsdk:"version" yaml:"version,omitempty"`
			} `tfsdk:"blue" yaml:"blue,omitempty"`

			Enable *bool `tfsdk:"enable" yaml:"enable,omitempty"`

			Green *struct {
				Image *string `tfsdk:"image" yaml:"image,omitempty"`

				IngressName *string `tfsdk:"ingress_name" yaml:"ingressName,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				ResourceLabels *map[string]string `tfsdk:"resource_labels" yaml:"resourceLabels,omitempty"`

				Version *string `tfsdk:"version" yaml:"version,omitempty"`
			} `tfsdk:"green" yaml:"green,omitempty"`

			ProductionDeployment *string `tfsdk:"production_deployment" yaml:"productionDeployment,omitempty"`
		} `tfsdk:"blue_green" yaml:"blueGreen,omitempty"`

		Canary *struct {
			Deployment *struct {
				Image *string `tfsdk:"image" yaml:"image,omitempty"`

				IngressName *string `tfsdk:"ingress_name" yaml:"ingressName,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				ResourceLabels *map[string]string `tfsdk:"resource_labels" yaml:"resourceLabels,omitempty"`

				Version *string `tfsdk:"version" yaml:"version,omitempty"`
			} `tfsdk:"deployment" yaml:"deployment,omitempty"`

			Enable *bool `tfsdk:"enable" yaml:"enable,omitempty"`
		} `tfsdk:"canary" yaml:"canary,omitempty"`

		Database *struct {
			BackupRemoteDeletePolicy *string `tfsdk:"backup_remote_delete_policy" yaml:"backupRemoteDeletePolicy,omitempty"`

			BackupRestoreSecretName *string `tfsdk:"backup_restore_secret_name" yaml:"backupRestoreSecretName,omitempty"`

			BackupSchedule *string `tfsdk:"backup_schedule" yaml:"backupSchedule,omitempty"`

			BackupSecretName *string `tfsdk:"backup_secret_name" yaml:"backupSecretName,omitempty"`

			BackupURL *string `tfsdk:"backup_url" yaml:"backupURL,omitempty"`

			InitBucketURL *string `tfsdk:"init_bucket_url" yaml:"initBucketURL,omitempty"`

			Replicas *int64 `tfsdk:"replicas" yaml:"replicas,omitempty"`

			Resources *struct {
				Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

				Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
			} `tfsdk:"resources" yaml:"resources,omitempty"`

			Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`

			StorageSize *string `tfsdk:"storage_size" yaml:"storageSize,omitempty"`

			Type *string `tfsdk:"type" yaml:"type,omitempty"`
		} `tfsdk:"database" yaml:"database,omitempty"`

		ElasticSearch *struct {
			Host *string `tfsdk:"host" yaml:"host,omitempty"`

			Password *string `tfsdk:"password" yaml:"password,omitempty"`

			Username *string `tfsdk:"username" yaml:"username,omitempty"`
		} `tfsdk:"elastic_search" yaml:"elasticSearch,omitempty"`

		Image *string `tfsdk:"image" yaml:"image,omitempty"`

		ImagePullPolicy *string `tfsdk:"image_pull_policy" yaml:"imagePullPolicy,omitempty"`

		IngressAnnotations *map[string]string `tfsdk:"ingress_annotations" yaml:"ingressAnnotations,omitempty"`

		IngressName *string `tfsdk:"ingress_name" yaml:"ingressName,omitempty"`

		LivenessProbe *struct {
			Exec *struct {
				Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
			} `tfsdk:"exec" yaml:"exec,omitempty"`

			FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

			Grpc *struct {
				Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

				Service *string `tfsdk:"service" yaml:"service,omitempty"`
			} `tfsdk:"grpc" yaml:"grpc,omitempty"`

			HttpGet *struct {
				Host *string `tfsdk:"host" yaml:"host,omitempty"`

				HttpHeaders *[]struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

				Path *string `tfsdk:"path" yaml:"path,omitempty"`

				Port *string `tfsdk:"port" yaml:"port,omitempty"`

				Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
			} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

			InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

			PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

			SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

			TcpSocket *struct {
				Host *string `tfsdk:"host" yaml:"host,omitempty"`

				Port *string `tfsdk:"port" yaml:"port,omitempty"`
			} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`

			TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" yaml:"terminationGracePeriodSeconds,omitempty"`

			TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
		} `tfsdk:"liveness_probe" yaml:"livenessProbe,omitempty"`

		MattermostEnv *[]struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Value *string `tfsdk:"value" yaml:"value,omitempty"`

			ValueFrom *struct {
				ConfigMapKeyRef *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`

				FieldRef *struct {
					ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

					FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`
				} `tfsdk:"field_ref" yaml:"fieldRef,omitempty"`

				ResourceFieldRef *struct {
					ContainerName *string `tfsdk:"container_name" yaml:"containerName,omitempty"`

					Divisor *string `tfsdk:"divisor" yaml:"divisor,omitempty"`

					Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`
				} `tfsdk:"resource_field_ref" yaml:"resourceFieldRef,omitempty"`

				SecretKeyRef *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
			} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
		} `tfsdk:"mattermost_env" yaml:"mattermostEnv,omitempty"`

		MattermostLicenseSecret *string `tfsdk:"mattermost_license_secret" yaml:"mattermostLicenseSecret,omitempty"`

		Migrate *bool `tfsdk:"migrate" yaml:"migrate,omitempty"`

		Minio *struct {
			ExternalBucket *string `tfsdk:"external_bucket" yaml:"externalBucket,omitempty"`

			ExternalURL *string `tfsdk:"external_url" yaml:"externalURL,omitempty"`

			Replicas *int64 `tfsdk:"replicas" yaml:"replicas,omitempty"`

			Resources *struct {
				Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

				Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
			} `tfsdk:"resources" yaml:"resources,omitempty"`

			Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`

			StorageSize *string `tfsdk:"storage_size" yaml:"storageSize,omitempty"`
		} `tfsdk:"minio" yaml:"minio,omitempty"`

		NodeSelector *map[string]string `tfsdk:"node_selector" yaml:"nodeSelector,omitempty"`

		ReadinessProbe *struct {
			Exec *struct {
				Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
			} `tfsdk:"exec" yaml:"exec,omitempty"`

			FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

			Grpc *struct {
				Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

				Service *string `tfsdk:"service" yaml:"service,omitempty"`
			} `tfsdk:"grpc" yaml:"grpc,omitempty"`

			HttpGet *struct {
				Host *string `tfsdk:"host" yaml:"host,omitempty"`

				HttpHeaders *[]struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

				Path *string `tfsdk:"path" yaml:"path,omitempty"`

				Port *string `tfsdk:"port" yaml:"port,omitempty"`

				Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
			} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

			InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

			PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

			SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

			TcpSocket *struct {
				Host *string `tfsdk:"host" yaml:"host,omitempty"`

				Port *string `tfsdk:"port" yaml:"port,omitempty"`
			} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`

			TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" yaml:"terminationGracePeriodSeconds,omitempty"`

			TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
		} `tfsdk:"readiness_probe" yaml:"readinessProbe,omitempty"`

		Replicas *int64 `tfsdk:"replicas" yaml:"replicas,omitempty"`

		ResourceLabels *map[string]string `tfsdk:"resource_labels" yaml:"resourceLabels,omitempty"`

		Resources *struct {
			Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

			Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
		} `tfsdk:"resources" yaml:"resources,omitempty"`

		ServiceAnnotations *map[string]string `tfsdk:"service_annotations" yaml:"serviceAnnotations,omitempty"`

		Size *string `tfsdk:"size" yaml:"size,omitempty"`

		UseIngressTLS *bool `tfsdk:"use_ingress_tls" yaml:"useIngressTLS,omitempty"`

		UseServiceLoadBalancer *bool `tfsdk:"use_service_load_balancer" yaml:"useServiceLoadBalancer,omitempty"`

		Version *string `tfsdk:"version" yaml:"version,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewMattermostComClusterInstallationV1Alpha1Resource() resource.Resource {
	return &MattermostComClusterInstallationV1Alpha1Resource{}
}

func (r *MattermostComClusterInstallationV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mattermost_com_cluster_installation_v1alpha1"
}

func (r *MattermostComClusterInstallationV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "ClusterInstallation is the Schema for the clusterinstallations API",
		MarkdownDescription: "ClusterInstallation is the Schema for the clusterinstallations API",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description:         "The timestamp of the last change to this resource.",
				MarkdownDescription: "The timestamp of the last change to this resource.",
				Type:                types.Int64Type,
				Computed:            true,
				Optional:            false,
			},

			"yaml": {
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"metadata": {
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Type:                types.StringType,
						Required:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
					},

					"namespace": {
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Type:                types.StringType,
						Optional:            true,
					},

					"labels": {
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.LabelValidator(),
						},
					},
					"annotations": {
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.AnnotationValidator(),
						},
					},
				}),
			},

			"api_version": {
				Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"kind": {
				Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"spec": {
				Description:         "Specification of the desired behavior of the Mattermost cluster. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status",
				MarkdownDescription: "Specification of the desired behavior of the Mattermost cluster. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"affinity": {
						Description:         "If specified, affinity will define the pod's scheduling constraints",
						MarkdownDescription: "If specified, affinity will define the pod's scheduling constraints",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"node_affinity": {
								Description:         "Describes node affinity scheduling rules for the pod.",
								MarkdownDescription: "Describes node affinity scheduling rules for the pod.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"preferred_during_scheduling_ignored_during_execution": {
										Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.",
										MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"preference": {
												Description:         "A node selector term, associated with the corresponding weight.",
												MarkdownDescription: "A node selector term, associated with the corresponding weight.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"match_expressions": {
														Description:         "A list of node selector requirements by node's labels.",
														MarkdownDescription: "A list of node selector requirements by node's labels.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The label key that the selector applies to.",
																MarkdownDescription: "The label key that the selector applies to.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"operator": {
																Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"values": {
																Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"match_fields": {
														Description:         "A list of node selector requirements by node's fields.",
														MarkdownDescription: "A list of node selector requirements by node's fields.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The label key that the selector applies to.",
																MarkdownDescription: "The label key that the selector applies to.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"operator": {
																Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"values": {
																Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},

											"weight": {
												Description:         "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
												MarkdownDescription: "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",

												Type: types.Int64Type,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"required_during_scheduling_ignored_during_execution": {
										Description:         "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.",
										MarkdownDescription: "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"node_selector_terms": {
												Description:         "Required. A list of node selector terms. The terms are ORed.",
												MarkdownDescription: "Required. A list of node selector terms. The terms are ORed.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"match_expressions": {
														Description:         "A list of node selector requirements by node's labels.",
														MarkdownDescription: "A list of node selector requirements by node's labels.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The label key that the selector applies to.",
																MarkdownDescription: "The label key that the selector applies to.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"operator": {
																Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"values": {
																Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"match_fields": {
														Description:         "A list of node selector requirements by node's fields.",
														MarkdownDescription: "A list of node selector requirements by node's fields.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The label key that the selector applies to.",
																MarkdownDescription: "The label key that the selector applies to.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"operator": {
																Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"values": {
																Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pod_affinity": {
								Description:         "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",
								MarkdownDescription: "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"preferred_during_scheduling_ignored_during_execution": {
										Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
										MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"pod_affinity_term": {
												Description:         "Required. A pod affinity term, associated with the corresponding weight.",
												MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"label_selector": {
														Description:         "A label query over a set of resources, in this case pods.",
														MarkdownDescription: "A label query over a set of resources, in this case pods.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"match_expressions": {
																Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "key is the label key that the selector applies to.",
																		MarkdownDescription: "key is the label key that the selector applies to.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"operator": {
																		Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"values": {
																		Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"match_labels": {
																Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespace_selector": {
														Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces. This field is beta-level and is only honored when PodAffinityNamespaceSelector feature is enabled.",
														MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces. This field is beta-level and is only honored when PodAffinityNamespaceSelector feature is enabled.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"match_expressions": {
																Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "key is the label key that the selector applies to.",
																		MarkdownDescription: "key is the label key that the selector applies to.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"operator": {
																		Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"values": {
																		Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"match_labels": {
																Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespaces": {
														Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'",
														MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"topology_key": {
														Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
														MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},

											"weight": {
												Description:         "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
												MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",

												Type: types.Int64Type,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"required_during_scheduling_ignored_during_execution": {
										Description:         "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
										MarkdownDescription: "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"label_selector": {
												Description:         "A label query over a set of resources, in this case pods.",
												MarkdownDescription: "A label query over a set of resources, in this case pods.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"match_expressions": {
														Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
														MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "key is the label key that the selector applies to.",
																MarkdownDescription: "key is the label key that the selector applies to.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"operator": {
																Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"values": {
																Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"match_labels": {
														Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
														MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"namespace_selector": {
												Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces. This field is beta-level and is only honored when PodAffinityNamespaceSelector feature is enabled.",
												MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces. This field is beta-level and is only honored when PodAffinityNamespaceSelector feature is enabled.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"match_expressions": {
														Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
														MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "key is the label key that the selector applies to.",
																MarkdownDescription: "key is the label key that the selector applies to.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"operator": {
																Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"values": {
																Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"match_labels": {
														Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
														MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"namespaces": {
												Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'",
												MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"topology_key": {
												Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
												MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pod_anti_affinity": {
								Description:         "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",
								MarkdownDescription: "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"preferred_during_scheduling_ignored_during_execution": {
										Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
										MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"pod_affinity_term": {
												Description:         "Required. A pod affinity term, associated with the corresponding weight.",
												MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"label_selector": {
														Description:         "A label query over a set of resources, in this case pods.",
														MarkdownDescription: "A label query over a set of resources, in this case pods.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"match_expressions": {
																Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "key is the label key that the selector applies to.",
																		MarkdownDescription: "key is the label key that the selector applies to.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"operator": {
																		Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"values": {
																		Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"match_labels": {
																Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespace_selector": {
														Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces. This field is beta-level and is only honored when PodAffinityNamespaceSelector feature is enabled.",
														MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces. This field is beta-level and is only honored when PodAffinityNamespaceSelector feature is enabled.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"match_expressions": {
																Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "key is the label key that the selector applies to.",
																		MarkdownDescription: "key is the label key that the selector applies to.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"operator": {
																		Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"values": {
																		Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"match_labels": {
																Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespaces": {
														Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'",
														MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"topology_key": {
														Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
														MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},

											"weight": {
												Description:         "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
												MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",

												Type: types.Int64Type,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"required_during_scheduling_ignored_during_execution": {
										Description:         "If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
										MarkdownDescription: "If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"label_selector": {
												Description:         "A label query over a set of resources, in this case pods.",
												MarkdownDescription: "A label query over a set of resources, in this case pods.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"match_expressions": {
														Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
														MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "key is the label key that the selector applies to.",
																MarkdownDescription: "key is the label key that the selector applies to.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"operator": {
																Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"values": {
																Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"match_labels": {
														Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
														MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"namespace_selector": {
												Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces. This field is beta-level and is only honored when PodAffinityNamespaceSelector feature is enabled.",
												MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces. This field is beta-level and is only honored when PodAffinityNamespaceSelector feature is enabled.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"match_expressions": {
														Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
														MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "key is the label key that the selector applies to.",
																MarkdownDescription: "key is the label key that the selector applies to.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"operator": {
																Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"values": {
																Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"match_labels": {
														Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
														MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"namespaces": {
												Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'",
												MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"topology_key": {
												Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
												MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"blue_green": {
						Description:         "BlueGreen defines the configuration of BlueGreen deployment for a ClusterInstallation",
						MarkdownDescription: "BlueGreen defines the configuration of BlueGreen deployment for a ClusterInstallation",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"blue": {
								Description:         "Blue defines the blue deployment.",
								MarkdownDescription: "Blue defines the blue deployment.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"image": {
										Description:         "Image defines the base Docker image that will be used for the deployment. Required when BlueGreen or Canary is enabled.",
										MarkdownDescription: "Image defines the base Docker image that will be used for the deployment. Required when BlueGreen or Canary is enabled.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"ingress_name": {
										Description:         "IngressName defines the ingress name that will be used by the deployment. This option is not used for Canary builds.",
										MarkdownDescription: "IngressName defines the ingress name that will be used by the deployment. This option is not used for Canary builds.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "Name defines the name of the deployment",
										MarkdownDescription: "Name defines the name of the deployment",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"resource_labels": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"version": {
										Description:         "Version defines the Docker image version that will be used for the deployment. Required when BlueGreen or Canary is enabled.",
										MarkdownDescription: "Version defines the Docker image version that will be used for the deployment. Required when BlueGreen or Canary is enabled.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable": {
								Description:         "Enable defines if BlueGreen deployment will be applied.",
								MarkdownDescription: "Enable defines if BlueGreen deployment will be applied.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"green": {
								Description:         "Green defines the green deployment.",
								MarkdownDescription: "Green defines the green deployment.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"image": {
										Description:         "Image defines the base Docker image that will be used for the deployment. Required when BlueGreen or Canary is enabled.",
										MarkdownDescription: "Image defines the base Docker image that will be used for the deployment. Required when BlueGreen or Canary is enabled.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"ingress_name": {
										Description:         "IngressName defines the ingress name that will be used by the deployment. This option is not used for Canary builds.",
										MarkdownDescription: "IngressName defines the ingress name that will be used by the deployment. This option is not used for Canary builds.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "Name defines the name of the deployment",
										MarkdownDescription: "Name defines the name of the deployment",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"resource_labels": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"version": {
										Description:         "Version defines the Docker image version that will be used for the deployment. Required when BlueGreen or Canary is enabled.",
										MarkdownDescription: "Version defines the Docker image version that will be used for the deployment. Required when BlueGreen or Canary is enabled.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"production_deployment": {
								Description:         "ProductionDeployment defines if the current production is blue or green.",
								MarkdownDescription: "ProductionDeployment defines if the current production is blue or green.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"canary": {
						Description:         "Canary defines the configuration of Canary deployment for a ClusterInstallation",
						MarkdownDescription: "Canary defines the configuration of Canary deployment for a ClusterInstallation",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"deployment": {
								Description:         "Deployment defines the canary deployment.",
								MarkdownDescription: "Deployment defines the canary deployment.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"image": {
										Description:         "Image defines the base Docker image that will be used for the deployment. Required when BlueGreen or Canary is enabled.",
										MarkdownDescription: "Image defines the base Docker image that will be used for the deployment. Required when BlueGreen or Canary is enabled.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"ingress_name": {
										Description:         "IngressName defines the ingress name that will be used by the deployment. This option is not used for Canary builds.",
										MarkdownDescription: "IngressName defines the ingress name that will be used by the deployment. This option is not used for Canary builds.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "Name defines the name of the deployment",
										MarkdownDescription: "Name defines the name of the deployment",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"resource_labels": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"version": {
										Description:         "Version defines the Docker image version that will be used for the deployment. Required when BlueGreen or Canary is enabled.",
										MarkdownDescription: "Version defines the Docker image version that will be used for the deployment. Required when BlueGreen or Canary is enabled.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable": {
								Description:         "Enable defines if a canary build will be deployed.",
								MarkdownDescription: "Enable defines if a canary build will be deployed.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"database": {
						Description:         "Database defines the database configuration for a ClusterInstallation.",
						MarkdownDescription: "Database defines the database configuration for a ClusterInstallation.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"backup_remote_delete_policy": {
								Description:         "Defines the backup retention policy.",
								MarkdownDescription: "Defines the backup retention policy.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"backup_restore_secret_name": {
								Description:         "Defines the secret to be used when performing a database restore.",
								MarkdownDescription: "Defines the secret to be used when performing a database restore.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"backup_schedule": {
								Description:         "Defines the interval for backups in cron expression format.",
								MarkdownDescription: "Defines the interval for backups in cron expression format.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"backup_secret_name": {
								Description:         "Defines the secret to be used for uploading/restoring backup.",
								MarkdownDescription: "Defines the secret to be used for uploading/restoring backup.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"backup_url": {
								Description:         "Defines the object storage url for uploading backups.",
								MarkdownDescription: "Defines the object storage url for uploading backups.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"init_bucket_url": {
								Description:         "Defines the AWS S3 bucket where the Database Backup is stored. The operator will download the file to restore the data.",
								MarkdownDescription: "Defines the AWS S3 bucket where the Database Backup is stored. The operator will download the file to restore the data.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"replicas": {
								Description:         "Defines the number of database replicas. For redundancy use at least 2 replicas. Setting this will override the number of replicas set by 'Size'.",
								MarkdownDescription: "Defines the number of database replicas. For redundancy use at least 2 replicas. Setting this will override the number of replicas set by 'Size'.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"resources": {
								Description:         "Defines the resource requests and limits for the database pods.",
								MarkdownDescription: "Defines the resource requests and limits for the database pods.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"limits": {
										Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"requests": {
										Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"secret": {
								Description:         "Optionally enter the name of an already-existing Secret for connecting to the database. This secret should be configured as follows:  User-Managed Database   - Key: DB_CONNECTION_STRING | Value: <FULL_DATABASE_CONNECTION_STRING> Operator-Managed Database   - Key: ROOT_PASSWORD | Value: <ROOT_DATABASE_PASSWORD>   - Key: USER | Value: <USER_NAME>   - Key: PASSWORD | Value: <USER_PASSWORD>   - Key: DATABASE Value: <DATABASE_NAME>  Notes:   If you define all secret values for both User-Managed and   Operator-Managed database types, the User-Managed connection string will   take precedence and the Operator-Managed values will be ignored. If the   secret is left blank, the default behavior is to use an Operator-Managed   database with strong randomly-generated database credentials.",
								MarkdownDescription: "Optionally enter the name of an already-existing Secret for connecting to the database. This secret should be configured as follows:  User-Managed Database   - Key: DB_CONNECTION_STRING | Value: <FULL_DATABASE_CONNECTION_STRING> Operator-Managed Database   - Key: ROOT_PASSWORD | Value: <ROOT_DATABASE_PASSWORD>   - Key: USER | Value: <USER_NAME>   - Key: PASSWORD | Value: <USER_PASSWORD>   - Key: DATABASE Value: <DATABASE_NAME>  Notes:   If you define all secret values for both User-Managed and   Operator-Managed database types, the User-Managed connection string will   take precedence and the Operator-Managed values will be ignored. If the   secret is left blank, the default behavior is to use an Operator-Managed   database with strong randomly-generated database credentials.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"storage_size": {
								Description:         "Defines the storage size for the database. ie 50Gi",
								MarkdownDescription: "Defines the storage size for the database. ie 50Gi",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"type": {
								Description:         "Defines the type of database to use for an Operator-Managed database. This value is ignored when using a User-Managed database.",
								MarkdownDescription: "Defines the type of database to use for an Operator-Managed database. This value is ignored when using a User-Managed database.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"elastic_search": {
						Description:         "ElasticSearch defines the ElasticSearch configuration for a ClusterInstallation.",
						MarkdownDescription: "ElasticSearch defines the ElasticSearch configuration for a ClusterInstallation.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"host": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"password": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"username": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"image": {
						Description:         "Image defines the ClusterInstallation Docker image.",
						MarkdownDescription: "Image defines the ClusterInstallation Docker image.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"image_pull_policy": {
						Description:         "Specify deployment pull policy.",
						MarkdownDescription: "Specify deployment pull policy.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ingress_annotations": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ingress_name": {
						Description:         "IngressName defines the name to be used when creating the ingress rules",
						MarkdownDescription: "IngressName defines the name to be used when creating the ingress rules",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"liveness_probe": {
						Description:         "Defines the probe to check if the application is up and running.",
						MarkdownDescription: "Defines the probe to check if the application is up and running.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"exec": {
								Description:         "Exec specifies the action to take.",
								MarkdownDescription: "Exec specifies the action to take.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"command": {
										Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
										MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"failure_threshold": {
								Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
								MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"grpc": {
								Description:         "GRPC specifies an action involving a GRPC port. This is an alpha field and requires enabling GRPCContainerProbe feature gate.",
								MarkdownDescription: "GRPC specifies an action involving a GRPC port. This is an alpha field and requires enabling GRPCContainerProbe feature gate.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"port": {
										Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
										MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",

										Type: types.Int64Type,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"service": {
										Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",
										MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"http_get": {
								Description:         "HTTPGet specifies the http request to perform.",
								MarkdownDescription: "HTTPGet specifies the http request to perform.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"host": {
										Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
										MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"http_headers": {
										Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
										MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "The header field name",
												MarkdownDescription: "The header field name",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"value": {
												Description:         "The header field value",
												MarkdownDescription: "The header field value",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"path": {
										Description:         "Path to access on the HTTP server.",
										MarkdownDescription: "Path to access on the HTTP server.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"port": {
										Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
										MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"scheme": {
										Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
										MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"initial_delay_seconds": {
								Description:         "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
								MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"period_seconds": {
								Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
								MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"success_threshold": {
								Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
								MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"tcp_socket": {
								Description:         "TCPSocket specifies an action involving a TCP port.",
								MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"host": {
										Description:         "Optional: Host name to connect to, defaults to the pod IP.",
										MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"port": {
										Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
										MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"termination_grace_period_seconds": {
								Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
								MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"timeout_seconds": {
								Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
								MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"mattermost_env": {
						Description:         "Optional environment variables to set in the Mattermost application pods.",
						MarkdownDescription: "Optional environment variables to set in the Mattermost application pods.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
								MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"value": {
								Description:         "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
								MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"value_from": {
								Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
								MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"config_map_key_ref": {
										Description:         "Selects a key of a ConfigMap.",
										MarkdownDescription: "Selects a key of a ConfigMap.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "The key to select.",
												MarkdownDescription: "The key to select.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"name": {
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"optional": {
												Description:         "Specify whether the ConfigMap or its key must be defined",
												MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"field_ref": {
										Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
										MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"api_version": {
												Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
												MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"field_path": {
												Description:         "Path of the field to select in the specified API version.",
												MarkdownDescription: "Path of the field to select in the specified API version.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"resource_field_ref": {
										Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
										MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"container_name": {
												Description:         "Container name: required for volumes, optional for env vars",
												MarkdownDescription: "Container name: required for volumes, optional for env vars",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"divisor": {
												Description:         "Specifies the output format of the exposed resources, defaults to '1'",
												MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"resource": {
												Description:         "Required: resource to select",
												MarkdownDescription: "Required: resource to select",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret_key_ref": {
										Description:         "Selects a key of a secret in the pod's namespace",
										MarkdownDescription: "Selects a key of a secret in the pod's namespace",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "The key of the secret to select from.  Must be a valid secret key.",
												MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"name": {
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"optional": {
												Description:         "Specify whether the Secret or its key must be defined",
												MarkdownDescription: "Specify whether the Secret or its key must be defined",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"mattermost_license_secret": {
						Description:         "Secret that contains the mattermost license",
						MarkdownDescription: "Secret that contains the mattermost license",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"migrate": {
						Description:         "Migrate specifies that the ClusterInstallation CR should be migrated to the Mattermost CR. CAUTION: Some features like BlueGreen or Canary are not supported with a new Custom Resource therefore migration should be performed with extra caution.",
						MarkdownDescription: "Migrate specifies that the ClusterInstallation CR should be migrated to the Mattermost CR. CAUTION: Some features like BlueGreen or Canary are not supported with a new Custom Resource therefore migration should be performed with extra caution.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"minio": {
						Description:         "Minio defines the configuration of Minio for a ClusterInstallation.",
						MarkdownDescription: "Minio defines the configuration of Minio for a ClusterInstallation.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"external_bucket": {
								Description:         "Set to the bucket name of your external MinIO or S3.",
								MarkdownDescription: "Set to the bucket name of your external MinIO or S3.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"external_url": {
								Description:         "Set to use an external MinIO deployment or S3. Must also set 'Secret' and 'ExternalBucket'.",
								MarkdownDescription: "Set to use an external MinIO deployment or S3. Must also set 'Secret' and 'ExternalBucket'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"replicas": {
								Description:         "Defines the number of Minio replicas. Supply 1 to run Minio in standalone mode with no redundancy. Supply 4 or more to run Minio in distributed mode. Note that it is not possible to upgrade Minio from standalone to distributed mode. Setting this will override the number of replicas set by 'Size'. More info: https://docs.min.io/docs/distributed-minio-quickstart-guide.html",
								MarkdownDescription: "Defines the number of Minio replicas. Supply 1 to run Minio in standalone mode with no redundancy. Supply 4 or more to run Minio in distributed mode. Note that it is not possible to upgrade Minio from standalone to distributed mode. Setting this will override the number of replicas set by 'Size'. More info: https://docs.min.io/docs/distributed-minio-quickstart-guide.html",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"resources": {
								Description:         "Defines the resource requests and limits for the Minio pods.",
								MarkdownDescription: "Defines the resource requests and limits for the Minio pods.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"limits": {
										Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"requests": {
										Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"secret": {
								Description:         "Optionally enter the name of already existing secret. Secret should have two values: 'accesskey' and 'secretkey'. Required when 'ExternalURL' is set.",
								MarkdownDescription: "Optionally enter the name of already existing secret. Secret should have two values: 'accesskey' and 'secretkey'. Required when 'ExternalURL' is set.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"storage_size": {
								Description:         "Defines the storage size for Minio. ie 50Gi",
								MarkdownDescription: "Defines the storage size for Minio. ie 50Gi",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"node_selector": {
						Description:         "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/",
						MarkdownDescription: "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"readiness_probe": {
						Description:         "Defines the probe to check if the application is ready to accept traffic.",
						MarkdownDescription: "Defines the probe to check if the application is ready to accept traffic.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"exec": {
								Description:         "Exec specifies the action to take.",
								MarkdownDescription: "Exec specifies the action to take.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"command": {
										Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
										MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"failure_threshold": {
								Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
								MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"grpc": {
								Description:         "GRPC specifies an action involving a GRPC port. This is an alpha field and requires enabling GRPCContainerProbe feature gate.",
								MarkdownDescription: "GRPC specifies an action involving a GRPC port. This is an alpha field and requires enabling GRPCContainerProbe feature gate.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"port": {
										Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
										MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",

										Type: types.Int64Type,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"service": {
										Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",
										MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"http_get": {
								Description:         "HTTPGet specifies the http request to perform.",
								MarkdownDescription: "HTTPGet specifies the http request to perform.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"host": {
										Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
										MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"http_headers": {
										Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
										MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "The header field name",
												MarkdownDescription: "The header field name",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"value": {
												Description:         "The header field value",
												MarkdownDescription: "The header field value",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"path": {
										Description:         "Path to access on the HTTP server.",
										MarkdownDescription: "Path to access on the HTTP server.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"port": {
										Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
										MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"scheme": {
										Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
										MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"initial_delay_seconds": {
								Description:         "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
								MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"period_seconds": {
								Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
								MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"success_threshold": {
								Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
								MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"tcp_socket": {
								Description:         "TCPSocket specifies an action involving a TCP port.",
								MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"host": {
										Description:         "Optional: Host name to connect to, defaults to the pod IP.",
										MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"port": {
										Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
										MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"termination_grace_period_seconds": {
								Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
								MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"timeout_seconds": {
								Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
								MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"replicas": {
						Description:         "Replicas defines the number of replicas to use for the Mattermost app servers. Setting this will override the number of replicas set by 'Size'.",
						MarkdownDescription: "Replicas defines the number of replicas to use for the Mattermost app servers. Setting this will override the number of replicas set by 'Size'.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"resource_labels": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"resources": {
						Description:         "Defines the resource requests and limits for the Mattermost app server pods.",
						MarkdownDescription: "Defines the resource requests and limits for the Mattermost app server pods.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"limits": {
								Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"requests": {
								Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"service_annotations": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"size": {
						Description:         "Size defines the size of the ClusterInstallation. This is typically specified in number of users. This will override replica and resource requests/limits appropriately for the provided number of users. This is a write-only field - its value is erased after setting appropriate values of resources. Accepted values are: 100users, 1000users, 5000users, 10000users, 250000users. If replicas and resource requests/limits are not specified, and Size is not provided the configuration for 5000users will be applied. Setting 'Replicas', 'Resources', 'Minio.Replicas', 'Minio.Resource', 'Database.Replicas', or 'Database.Resources' will override the values set by Size. Setting new Size will override previous values regardless if set by Size or manually.",
						MarkdownDescription: "Size defines the size of the ClusterInstallation. This is typically specified in number of users. This will override replica and resource requests/limits appropriately for the provided number of users. This is a write-only field - its value is erased after setting appropriate values of resources. Accepted values are: 100users, 1000users, 5000users, 10000users, 250000users. If replicas and resource requests/limits are not specified, and Size is not provided the configuration for 5000users will be applied. Setting 'Replicas', 'Resources', 'Minio.Replicas', 'Minio.Resource', 'Database.Replicas', or 'Database.Resources' will override the values set by Size. Setting new Size will override previous values regardless if set by Size or manually.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"use_ingress_tls": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"use_service_load_balancer": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"version": {
						Description:         "Version defines the ClusterInstallation Docker image version.",
						MarkdownDescription: "Version defines the ClusterInstallation Docker image version.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}, nil
}

func (r *MattermostComClusterInstallationV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_mattermost_com_cluster_installation_v1alpha1")

	var state MattermostComClusterInstallationV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel MattermostComClusterInstallationV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("mattermost.com/v1alpha1")
	goModel.Kind = utilities.Ptr("ClusterInstallation")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *MattermostComClusterInstallationV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_mattermost_com_cluster_installation_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *MattermostComClusterInstallationV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_mattermost_com_cluster_installation_v1alpha1")

	var state MattermostComClusterInstallationV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel MattermostComClusterInstallationV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("mattermost.com/v1alpha1")
	goModel.Kind = utilities.Ptr("ClusterInstallation")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *MattermostComClusterInstallationV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_mattermost_com_cluster_installation_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
