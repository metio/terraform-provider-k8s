/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package flows_netobserv_io_v1alpha1

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
	_ datasource.DataSource = &FlowsNetobservIoFlowCollectorV1Alpha1Manifest{}
)

func NewFlowsNetobservIoFlowCollectorV1Alpha1Manifest() datasource.DataSource {
	return &FlowsNetobservIoFlowCollectorV1Alpha1Manifest{}
}

type FlowsNetobservIoFlowCollectorV1Alpha1Manifest struct{}

type FlowsNetobservIoFlowCollectorV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Agent *struct {
			Ebpf *struct {
				CacheActiveTimeout *string `tfsdk:"cache_active_timeout" json:"cacheActiveTimeout,omitempty"`
				CacheMaxFlows      *int64  `tfsdk:"cache_max_flows" json:"cacheMaxFlows,omitempty"`
				Debug              *struct {
					Env *map[string]string `tfsdk:"env" json:"env,omitempty"`
				} `tfsdk:"debug" json:"debug,omitempty"`
				ExcludeInterfaces *[]string `tfsdk:"exclude_interfaces" json:"excludeInterfaces,omitempty"`
				ImagePullPolicy   *string   `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
				Interfaces        *[]string `tfsdk:"interfaces" json:"interfaces,omitempty"`
				KafkaBatchSize    *int64    `tfsdk:"kafka_batch_size" json:"kafkaBatchSize,omitempty"`
				LogLevel          *string   `tfsdk:"log_level" json:"logLevel,omitempty"`
				Privileged        *bool     `tfsdk:"privileged" json:"privileged,omitempty"`
				Resources         *struct {
					Claims *[]struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"claims" json:"claims,omitempty"`
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				Sampling *int64 `tfsdk:"sampling" json:"sampling,omitempty"`
			} `tfsdk:"ebpf" json:"ebpf,omitempty"`
			Ipfix *struct {
				CacheActiveTimeout     *string `tfsdk:"cache_active_timeout" json:"cacheActiveTimeout,omitempty"`
				CacheMaxFlows          *int64  `tfsdk:"cache_max_flows" json:"cacheMaxFlows,omitempty"`
				ClusterNetworkOperator *struct {
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"cluster_network_operator" json:"clusterNetworkOperator,omitempty"`
				ForceSampleAll *bool `tfsdk:"force_sample_all" json:"forceSampleAll,omitempty"`
				OvnKubernetes  *struct {
					ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
					DaemonSetName *string `tfsdk:"daemon_set_name" json:"daemonSetName,omitempty"`
					Namespace     *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"ovn_kubernetes" json:"ovnKubernetes,omitempty"`
				Sampling *int64 `tfsdk:"sampling" json:"sampling,omitempty"`
			} `tfsdk:"ipfix" json:"ipfix,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"agent" json:"agent,omitempty"`
		ConsolePlugin *struct {
			Autoscaler *struct {
				MaxReplicas *int64 `tfsdk:"max_replicas" json:"maxReplicas,omitempty"`
				Metrics     *[]struct {
					ContainerResource *struct {
						Container *string `tfsdk:"container" json:"container,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Target    *struct {
							AverageUtilization *int64  `tfsdk:"average_utilization" json:"averageUtilization,omitempty"`
							AverageValue       *string `tfsdk:"average_value" json:"averageValue,omitempty"`
							Type               *string `tfsdk:"type" json:"type,omitempty"`
							Value              *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"target" json:"target,omitempty"`
					} `tfsdk:"container_resource" json:"containerResource,omitempty"`
					External *struct {
						Metric *struct {
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Selector *struct {
								MatchExpressions *[]struct {
									Key      *string   `tfsdk:"key" json:"key,omitempty"`
									Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
									Values   *[]string `tfsdk:"values" json:"values,omitempty"`
								} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
								MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
							} `tfsdk:"selector" json:"selector,omitempty"`
						} `tfsdk:"metric" json:"metric,omitempty"`
						Target *struct {
							AverageUtilization *int64  `tfsdk:"average_utilization" json:"averageUtilization,omitempty"`
							AverageValue       *string `tfsdk:"average_value" json:"averageValue,omitempty"`
							Type               *string `tfsdk:"type" json:"type,omitempty"`
							Value              *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"target" json:"target,omitempty"`
					} `tfsdk:"external" json:"external,omitempty"`
					Object *struct {
						DescribedObject *struct {
							ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
							Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
							Name       *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"described_object" json:"describedObject,omitempty"`
						Metric *struct {
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Selector *struct {
								MatchExpressions *[]struct {
									Key      *string   `tfsdk:"key" json:"key,omitempty"`
									Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
									Values   *[]string `tfsdk:"values" json:"values,omitempty"`
								} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
								MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
							} `tfsdk:"selector" json:"selector,omitempty"`
						} `tfsdk:"metric" json:"metric,omitempty"`
						Target *struct {
							AverageUtilization *int64  `tfsdk:"average_utilization" json:"averageUtilization,omitempty"`
							AverageValue       *string `tfsdk:"average_value" json:"averageValue,omitempty"`
							Type               *string `tfsdk:"type" json:"type,omitempty"`
							Value              *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"target" json:"target,omitempty"`
					} `tfsdk:"object" json:"object,omitempty"`
					Pods *struct {
						Metric *struct {
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Selector *struct {
								MatchExpressions *[]struct {
									Key      *string   `tfsdk:"key" json:"key,omitempty"`
									Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
									Values   *[]string `tfsdk:"values" json:"values,omitempty"`
								} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
								MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
							} `tfsdk:"selector" json:"selector,omitempty"`
						} `tfsdk:"metric" json:"metric,omitempty"`
						Target *struct {
							AverageUtilization *int64  `tfsdk:"average_utilization" json:"averageUtilization,omitempty"`
							AverageValue       *string `tfsdk:"average_value" json:"averageValue,omitempty"`
							Type               *string `tfsdk:"type" json:"type,omitempty"`
							Value              *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"target" json:"target,omitempty"`
					} `tfsdk:"pods" json:"pods,omitempty"`
					Resource *struct {
						Name   *string `tfsdk:"name" json:"name,omitempty"`
						Target *struct {
							AverageUtilization *int64  `tfsdk:"average_utilization" json:"averageUtilization,omitempty"`
							AverageValue       *string `tfsdk:"average_value" json:"averageValue,omitempty"`
							Type               *string `tfsdk:"type" json:"type,omitempty"`
							Value              *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"target" json:"target,omitempty"`
					} `tfsdk:"resource" json:"resource,omitempty"`
					Type *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"metrics" json:"metrics,omitempty"`
				MinReplicas *int64  `tfsdk:"min_replicas" json:"minReplicas,omitempty"`
				Status      *string `tfsdk:"status" json:"status,omitempty"`
			} `tfsdk:"autoscaler" json:"autoscaler,omitempty"`
			ImagePullPolicy *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
			LogLevel        *string `tfsdk:"log_level" json:"logLevel,omitempty"`
			Port            *int64  `tfsdk:"port" json:"port,omitempty"`
			PortNaming      *struct {
				Enable    *bool              `tfsdk:"enable" json:"enable,omitempty"`
				PortNames *map[string]string `tfsdk:"port_names" json:"portNames,omitempty"`
			} `tfsdk:"port_naming" json:"portNaming,omitempty"`
			QuickFilters *[]struct {
				Default *bool              `tfsdk:"default" json:"default,omitempty"`
				Filter  *map[string]string `tfsdk:"filter" json:"filter,omitempty"`
				Name    *string            `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"quick_filters" json:"quickFilters,omitempty"`
			Register  *bool  `tfsdk:"register" json:"register,omitempty"`
			Replicas  *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
			Resources *struct {
				Claims *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
		} `tfsdk:"console_plugin" json:"consolePlugin,omitempty"`
		DeploymentModel *string `tfsdk:"deployment_model" json:"deploymentModel,omitempty"`
		Exporters       *[]struct {
			Ipfix *struct {
				TargetHost *string `tfsdk:"target_host" json:"targetHost,omitempty"`
				TargetPort *int64  `tfsdk:"target_port" json:"targetPort,omitempty"`
				Transport  *string `tfsdk:"transport" json:"transport,omitempty"`
			} `tfsdk:"ipfix" json:"ipfix,omitempty"`
			Kafka *struct {
				Address *string `tfsdk:"address" json:"address,omitempty"`
				Sasl    *struct {
					ClientIDReference *struct {
						File      *string `tfsdk:"file" json:"file,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						Type      *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"client_id_reference" json:"clientIDReference,omitempty"`
					ClientSecretReference *struct {
						File      *string `tfsdk:"file" json:"file,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						Type      *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"client_secret_reference" json:"clientSecretReference,omitempty"`
					Type *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"sasl" json:"sasl,omitempty"`
				Tls *struct {
					CaCert *struct {
						CertFile  *string `tfsdk:"cert_file" json:"certFile,omitempty"`
						CertKey   *string `tfsdk:"cert_key" json:"certKey,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						Type      *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"ca_cert" json:"caCert,omitempty"`
					Enable             *bool `tfsdk:"enable" json:"enable,omitempty"`
					InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
					UserCert           *struct {
						CertFile  *string `tfsdk:"cert_file" json:"certFile,omitempty"`
						CertKey   *string `tfsdk:"cert_key" json:"certKey,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						Type      *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"user_cert" json:"userCert,omitempty"`
				} `tfsdk:"tls" json:"tls,omitempty"`
				Topic *string `tfsdk:"topic" json:"topic,omitempty"`
			} `tfsdk:"kafka" json:"kafka,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"exporters" json:"exporters,omitempty"`
		Kafka *struct {
			Address *string `tfsdk:"address" json:"address,omitempty"`
			Sasl    *struct {
				ClientIDReference *struct {
					File      *string `tfsdk:"file" json:"file,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Type      *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"client_id_reference" json:"clientIDReference,omitempty"`
				ClientSecretReference *struct {
					File      *string `tfsdk:"file" json:"file,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Type      *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"client_secret_reference" json:"clientSecretReference,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"sasl" json:"sasl,omitempty"`
			Tls *struct {
				CaCert *struct {
					CertFile  *string `tfsdk:"cert_file" json:"certFile,omitempty"`
					CertKey   *string `tfsdk:"cert_key" json:"certKey,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Type      *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"ca_cert" json:"caCert,omitempty"`
				Enable             *bool `tfsdk:"enable" json:"enable,omitempty"`
				InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
				UserCert           *struct {
					CertFile  *string `tfsdk:"cert_file" json:"certFile,omitempty"`
					CertKey   *string `tfsdk:"cert_key" json:"certKey,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Type      *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"user_cert" json:"userCert,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
			Topic *string `tfsdk:"topic" json:"topic,omitempty"`
		} `tfsdk:"kafka" json:"kafka,omitempty"`
		Loki *struct {
			AuthToken    *string            `tfsdk:"auth_token" json:"authToken,omitempty"`
			BatchSize    *int64             `tfsdk:"batch_size" json:"batchSize,omitempty"`
			BatchWait    *string            `tfsdk:"batch_wait" json:"batchWait,omitempty"`
			MaxBackoff   *string            `tfsdk:"max_backoff" json:"maxBackoff,omitempty"`
			MaxRetries   *int64             `tfsdk:"max_retries" json:"maxRetries,omitempty"`
			MinBackoff   *string            `tfsdk:"min_backoff" json:"minBackoff,omitempty"`
			QuerierUrl   *string            `tfsdk:"querier_url" json:"querierUrl,omitempty"`
			StaticLabels *map[string]string `tfsdk:"static_labels" json:"staticLabels,omitempty"`
			StatusUrl    *string            `tfsdk:"status_url" json:"statusUrl,omitempty"`
			TenantID     *string            `tfsdk:"tenant_id" json:"tenantID,omitempty"`
			Timeout      *string            `tfsdk:"timeout" json:"timeout,omitempty"`
			Tls          *struct {
				CaCert *struct {
					CertFile  *string `tfsdk:"cert_file" json:"certFile,omitempty"`
					CertKey   *string `tfsdk:"cert_key" json:"certKey,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Type      *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"ca_cert" json:"caCert,omitempty"`
				Enable             *bool `tfsdk:"enable" json:"enable,omitempty"`
				InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
				UserCert           *struct {
					CertFile  *string `tfsdk:"cert_file" json:"certFile,omitempty"`
					CertKey   *string `tfsdk:"cert_key" json:"certKey,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Type      *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"user_cert" json:"userCert,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
			Url *string `tfsdk:"url" json:"url,omitempty"`
		} `tfsdk:"loki" json:"loki,omitempty"`
		Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		Processor *struct {
			Debug *struct {
				Env *map[string]string `tfsdk:"env" json:"env,omitempty"`
			} `tfsdk:"debug" json:"debug,omitempty"`
			DropUnusedFields        *bool   `tfsdk:"drop_unused_fields" json:"dropUnusedFields,omitempty"`
			EnableKubeProbes        *bool   `tfsdk:"enable_kube_probes" json:"enableKubeProbes,omitempty"`
			HealthPort              *int64  `tfsdk:"health_port" json:"healthPort,omitempty"`
			ImagePullPolicy         *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
			KafkaConsumerAutoscaler *struct {
				MaxReplicas *int64 `tfsdk:"max_replicas" json:"maxReplicas,omitempty"`
				Metrics     *[]struct {
					ContainerResource *struct {
						Container *string `tfsdk:"container" json:"container,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Target    *struct {
							AverageUtilization *int64  `tfsdk:"average_utilization" json:"averageUtilization,omitempty"`
							AverageValue       *string `tfsdk:"average_value" json:"averageValue,omitempty"`
							Type               *string `tfsdk:"type" json:"type,omitempty"`
							Value              *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"target" json:"target,omitempty"`
					} `tfsdk:"container_resource" json:"containerResource,omitempty"`
					External *struct {
						Metric *struct {
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Selector *struct {
								MatchExpressions *[]struct {
									Key      *string   `tfsdk:"key" json:"key,omitempty"`
									Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
									Values   *[]string `tfsdk:"values" json:"values,omitempty"`
								} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
								MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
							} `tfsdk:"selector" json:"selector,omitempty"`
						} `tfsdk:"metric" json:"metric,omitempty"`
						Target *struct {
							AverageUtilization *int64  `tfsdk:"average_utilization" json:"averageUtilization,omitempty"`
							AverageValue       *string `tfsdk:"average_value" json:"averageValue,omitempty"`
							Type               *string `tfsdk:"type" json:"type,omitempty"`
							Value              *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"target" json:"target,omitempty"`
					} `tfsdk:"external" json:"external,omitempty"`
					Object *struct {
						DescribedObject *struct {
							ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
							Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
							Name       *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"described_object" json:"describedObject,omitempty"`
						Metric *struct {
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Selector *struct {
								MatchExpressions *[]struct {
									Key      *string   `tfsdk:"key" json:"key,omitempty"`
									Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
									Values   *[]string `tfsdk:"values" json:"values,omitempty"`
								} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
								MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
							} `tfsdk:"selector" json:"selector,omitempty"`
						} `tfsdk:"metric" json:"metric,omitempty"`
						Target *struct {
							AverageUtilization *int64  `tfsdk:"average_utilization" json:"averageUtilization,omitempty"`
							AverageValue       *string `tfsdk:"average_value" json:"averageValue,omitempty"`
							Type               *string `tfsdk:"type" json:"type,omitempty"`
							Value              *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"target" json:"target,omitempty"`
					} `tfsdk:"object" json:"object,omitempty"`
					Pods *struct {
						Metric *struct {
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Selector *struct {
								MatchExpressions *[]struct {
									Key      *string   `tfsdk:"key" json:"key,omitempty"`
									Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
									Values   *[]string `tfsdk:"values" json:"values,omitempty"`
								} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
								MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
							} `tfsdk:"selector" json:"selector,omitempty"`
						} `tfsdk:"metric" json:"metric,omitempty"`
						Target *struct {
							AverageUtilization *int64  `tfsdk:"average_utilization" json:"averageUtilization,omitempty"`
							AverageValue       *string `tfsdk:"average_value" json:"averageValue,omitempty"`
							Type               *string `tfsdk:"type" json:"type,omitempty"`
							Value              *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"target" json:"target,omitempty"`
					} `tfsdk:"pods" json:"pods,omitempty"`
					Resource *struct {
						Name   *string `tfsdk:"name" json:"name,omitempty"`
						Target *struct {
							AverageUtilization *int64  `tfsdk:"average_utilization" json:"averageUtilization,omitempty"`
							AverageValue       *string `tfsdk:"average_value" json:"averageValue,omitempty"`
							Type               *string `tfsdk:"type" json:"type,omitempty"`
							Value              *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"target" json:"target,omitempty"`
					} `tfsdk:"resource" json:"resource,omitempty"`
					Type *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"metrics" json:"metrics,omitempty"`
				MinReplicas *int64  `tfsdk:"min_replicas" json:"minReplicas,omitempty"`
				Status      *string `tfsdk:"status" json:"status,omitempty"`
			} `tfsdk:"kafka_consumer_autoscaler" json:"kafkaConsumerAutoscaler,omitempty"`
			KafkaConsumerBatchSize     *int64  `tfsdk:"kafka_consumer_batch_size" json:"kafkaConsumerBatchSize,omitempty"`
			KafkaConsumerQueueCapacity *int64  `tfsdk:"kafka_consumer_queue_capacity" json:"kafkaConsumerQueueCapacity,omitempty"`
			KafkaConsumerReplicas      *int64  `tfsdk:"kafka_consumer_replicas" json:"kafkaConsumerReplicas,omitempty"`
			LogLevel                   *string `tfsdk:"log_level" json:"logLevel,omitempty"`
			Metrics                    *struct {
				IgnoreTags *[]string `tfsdk:"ignore_tags" json:"ignoreTags,omitempty"`
				Server     *struct {
					Port *int64 `tfsdk:"port" json:"port,omitempty"`
					Tls  *struct {
						Provided *struct {
							CertFile  *string `tfsdk:"cert_file" json:"certFile,omitempty"`
							CertKey   *string `tfsdk:"cert_key" json:"certKey,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
							Type      *string `tfsdk:"type" json:"type,omitempty"`
						} `tfsdk:"provided" json:"provided,omitempty"`
						Type *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"tls" json:"tls,omitempty"`
				} `tfsdk:"server" json:"server,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Port        *int64 `tfsdk:"port" json:"port,omitempty"`
			ProfilePort *int64 `tfsdk:"profile_port" json:"profilePort,omitempty"`
			Resources   *struct {
				Claims *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
		} `tfsdk:"processor" json:"processor,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *FlowsNetobservIoFlowCollectorV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_flows_netobserv_io_flow_collector_v1alpha1_manifest"
}

func (r *FlowsNetobservIoFlowCollectorV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "FlowCollector is the Schema for the flowcollectors API, which pilots and configures netflow collection.  Deprecated: This package will be removed in one of the next releases.",
		MarkdownDescription: "FlowCollector is the Schema for the flowcollectors API, which pilots and configures netflow collection.  Deprecated: This package will be removed in one of the next releases.",
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
				Description:         "FlowCollectorSpec defines the desired state of FlowCollector",
				MarkdownDescription: "FlowCollectorSpec defines the desired state of FlowCollector",
				Attributes: map[string]schema.Attribute{
					"agent": schema.SingleNestedAttribute{
						Description:         "agent for flows extraction.",
						MarkdownDescription: "agent for flows extraction.",
						Attributes: map[string]schema.Attribute{
							"ebpf": schema.SingleNestedAttribute{
								Description:         "ebpf describes the settings related to the eBPF-based flow reporter when the 'agent.type' property is set to 'EBPF'.",
								MarkdownDescription: "ebpf describes the settings related to the eBPF-based flow reporter when the 'agent.type' property is set to 'EBPF'.",
								Attributes: map[string]schema.Attribute{
									"cache_active_timeout": schema.StringAttribute{
										Description:         "cacheActiveTimeout is the max period during which the reporter will aggregate flows before sending. Increasing 'cacheMaxFlows' and 'cacheActiveTimeout' can decrease the network traffic overhead and the CPU load, however you can expect higher memory consumption and an increased latency in the flow collection.",
										MarkdownDescription: "cacheActiveTimeout is the max period during which the reporter will aggregate flows before sending. Increasing 'cacheMaxFlows' and 'cacheActiveTimeout' can decrease the network traffic overhead and the CPU load, however you can expect higher memory consumption and an increased latency in the flow collection.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(ns|ms|s|m)?$`), ""),
										},
									},

									"cache_max_flows": schema.Int64Attribute{
										Description:         "cacheMaxFlows is the max number of flows in an aggregate; when reached, the reporter sends the flows. Increasing 'cacheMaxFlows' and 'cacheActiveTimeout' can decrease the network traffic overhead and the CPU load, however you can expect higher memory consumption and an increased latency in the flow collection.",
										MarkdownDescription: "cacheMaxFlows is the max number of flows in an aggregate; when reached, the reporter sends the flows. Increasing 'cacheMaxFlows' and 'cacheActiveTimeout' can decrease the network traffic overhead and the CPU load, however you can expect higher memory consumption and an increased latency in the flow collection.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"debug": schema.SingleNestedAttribute{
										Description:         "Debug allows setting some aspects of the internal configuration of the eBPF agent. This section is aimed exclusively for debugging and fine-grained performance optimizations (for example GOGC, GOMAXPROCS env vars). Users setting its values do it at their own risk.",
										MarkdownDescription: "Debug allows setting some aspects of the internal configuration of the eBPF agent. This section is aimed exclusively for debugging and fine-grained performance optimizations (for example GOGC, GOMAXPROCS env vars). Users setting its values do it at their own risk.",
										Attributes: map[string]schema.Attribute{
											"env": schema.MapAttribute{
												Description:         "env allows passing custom environment variables to the NetObserv Agent. Useful for passing some very concrete performance-tuning options (such as GOGC, GOMAXPROCS) that shouldn't be publicly exposed as part of the FlowCollector descriptor, as they are only useful in edge debug and support scenarios.",
												MarkdownDescription: "env allows passing custom environment variables to the NetObserv Agent. Useful for passing some very concrete performance-tuning options (such as GOGC, GOMAXPROCS) that shouldn't be publicly exposed as part of the FlowCollector descriptor, as they are only useful in edge debug and support scenarios.",
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

									"exclude_interfaces": schema.ListAttribute{
										Description:         "excludeInterfaces contains the interface names that will be excluded from flow tracing. If an entry is enclosed by slashes (such as '/br-/'), it will match as regular expression, otherwise it will be matched as a case-sensitive string.",
										MarkdownDescription: "excludeInterfaces contains the interface names that will be excluded from flow tracing. If an entry is enclosed by slashes (such as '/br-/'), it will match as regular expression, otherwise it will be matched as a case-sensitive string.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"image_pull_policy": schema.StringAttribute{
										Description:         "imagePullPolicy is the Kubernetes pull policy for the image defined above",
										MarkdownDescription: "imagePullPolicy is the Kubernetes pull policy for the image defined above",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("IfNotPresent", "Always", "Never"),
										},
									},

									"interfaces": schema.ListAttribute{
										Description:         "interfaces contains the interface names from where flows will be collected. If empty, the agent will fetch all the interfaces in the system, excepting the ones listed in ExcludeInterfaces. If an entry is enclosed by slashes (such as '/br-/'), it will match as regular expression, otherwise it will be matched as a case-sensitive string.",
										MarkdownDescription: "interfaces contains the interface names from where flows will be collected. If empty, the agent will fetch all the interfaces in the system, excepting the ones listed in ExcludeInterfaces. If an entry is enclosed by slashes (such as '/br-/'), it will match as regular expression, otherwise it will be matched as a case-sensitive string.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"kafka_batch_size": schema.Int64Attribute{
										Description:         "kafkaBatchSize limits the maximum size of a request in bytes before being sent to a partition. Ignored when not using Kafka. Default: 1MB.",
										MarkdownDescription: "kafkaBatchSize limits the maximum size of a request in bytes before being sent to a partition. Ignored when not using Kafka. Default: 1MB.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"log_level": schema.StringAttribute{
										Description:         "logLevel defines the log level for the NetObserv eBPF Agent",
										MarkdownDescription: "logLevel defines the log level for the NetObserv eBPF Agent",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("trace", "debug", "info", "warn", "error", "fatal", "panic"),
										},
									},

									"privileged": schema.BoolAttribute{
										Description:         "privileged mode for the eBPF Agent container. In general this setting can be ignored or set to false: in that case, the operator will set granular capabilities (BPF, PERFMON, NET_ADMIN, SYS_RESOURCE) to the container, to enable its correct operation. If for some reason these capabilities cannot be set (for example old kernel version not knowing CAP_BPF) then you can turn on this mode for more global privileges.",
										MarkdownDescription: "privileged mode for the eBPF Agent container. In general this setting can be ignored or set to false: in that case, the operator will set granular capabilities (BPF, PERFMON, NET_ADMIN, SYS_RESOURCE) to the container, to enable its correct operation. If for some reason these capabilities cannot be set (for example old kernel version not knowing CAP_BPF) then you can turn on this mode for more global privileges.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"resources": schema.SingleNestedAttribute{
										Description:         "resources are the compute resources required by this container. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "resources are the compute resources required by this container. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										Attributes: map[string]schema.Attribute{
											"claims": schema.ListNestedAttribute{
												Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
												MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
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

									"sampling": schema.Int64Attribute{
										Description:         "sampling rate of the flow reporter. 100 means one flow on 100 is sent. 0 or 1 means all flows are sampled.",
										MarkdownDescription: "sampling rate of the flow reporter. 100 means one flow on 100 is sent. 0 or 1 means all flows are sampled.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"ipfix": schema.SingleNestedAttribute{
								Description:         "ipfix describes the settings related to the IPFIX-based flow reporter when the 'agent.type' property is set to 'IPFIX'.",
								MarkdownDescription: "ipfix describes the settings related to the IPFIX-based flow reporter when the 'agent.type' property is set to 'IPFIX'.",
								Attributes: map[string]schema.Attribute{
									"cache_active_timeout": schema.StringAttribute{
										Description:         "cacheActiveTimeout is the max period during which the reporter will aggregate flows before sending",
										MarkdownDescription: "cacheActiveTimeout is the max period during which the reporter will aggregate flows before sending",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(ns|ms|s|m)?$`), ""),
										},
									},

									"cache_max_flows": schema.Int64Attribute{
										Description:         "cacheMaxFlows is the max number of flows in an aggregate; when reached, the reporter sends the flows",
										MarkdownDescription: "cacheMaxFlows is the max number of flows in an aggregate; when reached, the reporter sends the flows",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},

									"cluster_network_operator": schema.SingleNestedAttribute{
										Description:         "clusterNetworkOperator defines the settings related to the OpenShift Cluster Network Operator, when available.",
										MarkdownDescription: "clusterNetworkOperator defines the settings related to the OpenShift Cluster Network Operator, when available.",
										Attributes: map[string]schema.Attribute{
											"namespace": schema.StringAttribute{
												Description:         "namespace  where the config map is going to be deployed.",
												MarkdownDescription: "namespace  where the config map is going to be deployed.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"force_sample_all": schema.BoolAttribute{
										Description:         "forceSampleAll allows disabling sampling in the IPFIX-based flow reporter. It is not recommended to sample all the traffic with IPFIX, as it might generate cluster instability. If you REALLY want to do that, set this flag to true. Use at your own risk. When it is set to true, the value of 'sampling' is ignored.",
										MarkdownDescription: "forceSampleAll allows disabling sampling in the IPFIX-based flow reporter. It is not recommended to sample all the traffic with IPFIX, as it might generate cluster instability. If you REALLY want to do that, set this flag to true. Use at your own risk. When it is set to true, the value of 'sampling' is ignored.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ovn_kubernetes": schema.SingleNestedAttribute{
										Description:         "ovnKubernetes defines the settings of the OVN-Kubernetes CNI, when available. This configuration is used when using OVN's IPFIX exports, without OpenShift. When using OpenShift, refer to the 'clusterNetworkOperator' property instead.",
										MarkdownDescription: "ovnKubernetes defines the settings of the OVN-Kubernetes CNI, when available. This configuration is used when using OVN's IPFIX exports, without OpenShift. When using OpenShift, refer to the 'clusterNetworkOperator' property instead.",
										Attributes: map[string]schema.Attribute{
											"container_name": schema.StringAttribute{
												Description:         "containerName defines the name of the container to configure for IPFIX.",
												MarkdownDescription: "containerName defines the name of the container to configure for IPFIX.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"daemon_set_name": schema.StringAttribute{
												Description:         "daemonSetName defines the name of the DaemonSet controlling the OVN-Kubernetes pods.",
												MarkdownDescription: "daemonSetName defines the name of the DaemonSet controlling the OVN-Kubernetes pods.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "namespace where OVN-Kubernetes pods are deployed.",
												MarkdownDescription: "namespace where OVN-Kubernetes pods are deployed.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"sampling": schema.Int64Attribute{
										Description:         "sampling is the sampling rate on the reporter. 100 means one flow on 100 is sent. To ensure cluster stability, it is not possible to set a value below 2. If you really want to sample every packet, which might impact the cluster stability, refer to 'forceSampleAll'. Alternatively, you can use the eBPF Agent instead of IPFIX.",
										MarkdownDescription: "sampling is the sampling rate on the reporter. 100 means one flow on 100 is sent. To ensure cluster stability, it is not possible to set a value below 2. If you really want to sample every packet, which might impact the cluster stability, refer to 'forceSampleAll'. Alternatively, you can use the eBPF Agent instead of IPFIX.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(2),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"type": schema.StringAttribute{
								Description:         "type selects the flows tracing agent. Possible values are 'EBPF' (default) to use NetObserv eBPF agent, 'IPFIX' to use the legacy IPFIX collector. 'EBPF' is recommended in most cases as it offers better performances and should work regardless of the CNI installed on the cluster. 'IPFIX' works with OVN-Kubernetes CNI (other CNIs could work if they support exporting IPFIX, but they would require manual configuration).",
								MarkdownDescription: "type selects the flows tracing agent. Possible values are 'EBPF' (default) to use NetObserv eBPF agent, 'IPFIX' to use the legacy IPFIX collector. 'EBPF' is recommended in most cases as it offers better performances and should work regardless of the CNI installed on the cluster. 'IPFIX' works with OVN-Kubernetes CNI (other CNIs could work if they support exporting IPFIX, but they would require manual configuration).",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("EBPF", "IPFIX"),
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"console_plugin": schema.SingleNestedAttribute{
						Description:         "consolePlugin defines the settings related to the OpenShift Console plugin, when available.",
						MarkdownDescription: "consolePlugin defines the settings related to the OpenShift Console plugin, when available.",
						Attributes: map[string]schema.Attribute{
							"autoscaler": schema.SingleNestedAttribute{
								Description:         "autoscaler spec of a horizontal pod autoscaler to set up for the plugin Deployment.",
								MarkdownDescription: "autoscaler spec of a horizontal pod autoscaler to set up for the plugin Deployment.",
								Attributes: map[string]schema.Attribute{
									"max_replicas": schema.Int64Attribute{
										Description:         "maxReplicas is the upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas.",
										MarkdownDescription: "maxReplicas is the upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"metrics": schema.ListNestedAttribute{
										Description:         "metrics used by the pod autoscaler",
										MarkdownDescription: "metrics used by the pod autoscaler",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"container_resource": schema.SingleNestedAttribute{
													Description:         "containerResource refers to a resource metric (such as those specified in requests and limits) known to Kubernetes describing a single container in each pod of the current scale target (e.g. CPU or memory). Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the 'pods' source. This is an alpha feature and can be enabled by the HPAContainerMetrics feature flag.",
													MarkdownDescription: "containerResource refers to a resource metric (such as those specified in requests and limits) known to Kubernetes describing a single container in each pod of the current scale target (e.g. CPU or memory). Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the 'pods' source. This is an alpha feature and can be enabled by the HPAContainerMetrics feature flag.",
													Attributes: map[string]schema.Attribute{
														"container": schema.StringAttribute{
															Description:         "container is the name of the container in the pods of the scaling target",
															MarkdownDescription: "container is the name of the container in the pods of the scaling target",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "name is the name of the resource in question.",
															MarkdownDescription: "name is the name of the resource in question.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"target": schema.SingleNestedAttribute{
															Description:         "target specifies the target value for the given metric",
															MarkdownDescription: "target specifies the target value for the given metric",
															Attributes: map[string]schema.Attribute{
																"average_utilization": schema.Int64Attribute{
																	Description:         "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
																	MarkdownDescription: "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"average_value": schema.StringAttribute{
																	Description:         "averageValue is the target value of the average of the metric across all relevant pods (as a quantity)",
																	MarkdownDescription: "averageValue is the target value of the average of the metric across all relevant pods (as a quantity)",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"type": schema.StringAttribute{
																	Description:         "type represents whether the metric type is Utilization, Value, or AverageValue",
																	MarkdownDescription: "type represents whether the metric type is Utilization, Value, or AverageValue",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"value": schema.StringAttribute{
																	Description:         "value is the target value of the metric (as a quantity).",
																	MarkdownDescription: "value is the target value of the metric (as a quantity).",
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
													Required: false,
													Optional: true,
													Computed: false,
												},

												"external": schema.SingleNestedAttribute{
													Description:         "external refers to a global metric that is not associated with any Kubernetes object. It allows autoscaling based on information coming from components running outside of cluster (for example length of queue in cloud messaging service, or QPS from loadbalancer running outside of cluster).",
													MarkdownDescription: "external refers to a global metric that is not associated with any Kubernetes object. It allows autoscaling based on information coming from components running outside of cluster (for example length of queue in cloud messaging service, or QPS from loadbalancer running outside of cluster).",
													Attributes: map[string]schema.Attribute{
														"metric": schema.SingleNestedAttribute{
															Description:         "metric identifies the target metric by name and selector",
															MarkdownDescription: "metric identifies the target metric by name and selector",
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "name is the name of the given metric",
																	MarkdownDescription: "name is the name of the given metric",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"selector": schema.SingleNestedAttribute{
																	Description:         "selector is the string-encoded form of a standard kubernetes label selector for the given metric When set, it is passed as an additional parameter to the metrics server for more specific metrics scoping. When unset, just the metricName will be used to gather metrics.",
																	MarkdownDescription: "selector is the string-encoded form of a standard kubernetes label selector for the given metric When set, it is passed as an additional parameter to the metrics server for more specific metrics scoping. When unset, just the metricName will be used to gather metrics.",
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
																						Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"values": schema.ListAttribute{
																						Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																						MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
																			Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																			MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
															Required: true,
															Optional: false,
															Computed: false,
														},

														"target": schema.SingleNestedAttribute{
															Description:         "target specifies the target value for the given metric",
															MarkdownDescription: "target specifies the target value for the given metric",
															Attributes: map[string]schema.Attribute{
																"average_utilization": schema.Int64Attribute{
																	Description:         "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
																	MarkdownDescription: "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"average_value": schema.StringAttribute{
																	Description:         "averageValue is the target value of the average of the metric across all relevant pods (as a quantity)",
																	MarkdownDescription: "averageValue is the target value of the average of the metric across all relevant pods (as a quantity)",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"type": schema.StringAttribute{
																	Description:         "type represents whether the metric type is Utilization, Value, or AverageValue",
																	MarkdownDescription: "type represents whether the metric type is Utilization, Value, or AverageValue",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"value": schema.StringAttribute{
																	Description:         "value is the target value of the metric (as a quantity).",
																	MarkdownDescription: "value is the target value of the metric (as a quantity).",
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
													Required: false,
													Optional: true,
													Computed: false,
												},

												"object": schema.SingleNestedAttribute{
													Description:         "object refers to a metric describing a single kubernetes object (for example, hits-per-second on an Ingress object).",
													MarkdownDescription: "object refers to a metric describing a single kubernetes object (for example, hits-per-second on an Ingress object).",
													Attributes: map[string]schema.Attribute{
														"described_object": schema.SingleNestedAttribute{
															Description:         "describedObject specifies the descriptions of a object,such as kind,name apiVersion",
															MarkdownDescription: "describedObject specifies the descriptions of a object,such as kind,name apiVersion",
															Attributes: map[string]schema.Attribute{
																"api_version": schema.StringAttribute{
																	Description:         "apiVersion is the API version of the referent",
																	MarkdownDescription: "apiVersion is the API version of the referent",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"kind": schema.StringAttribute{
																	Description:         "kind is the kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
																	MarkdownDescription: "kind is the kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "name is the name of the referent; More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	MarkdownDescription: "name is the name of the referent; More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},
															},
															Required: true,
															Optional: false,
															Computed: false,
														},

														"metric": schema.SingleNestedAttribute{
															Description:         "metric identifies the target metric by name and selector",
															MarkdownDescription: "metric identifies the target metric by name and selector",
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "name is the name of the given metric",
																	MarkdownDescription: "name is the name of the given metric",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"selector": schema.SingleNestedAttribute{
																	Description:         "selector is the string-encoded form of a standard kubernetes label selector for the given metric When set, it is passed as an additional parameter to the metrics server for more specific metrics scoping. When unset, just the metricName will be used to gather metrics.",
																	MarkdownDescription: "selector is the string-encoded form of a standard kubernetes label selector for the given metric When set, it is passed as an additional parameter to the metrics server for more specific metrics scoping. When unset, just the metricName will be used to gather metrics.",
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
																						Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"values": schema.ListAttribute{
																						Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																						MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
																			Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																			MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
															Required: true,
															Optional: false,
															Computed: false,
														},

														"target": schema.SingleNestedAttribute{
															Description:         "target specifies the target value for the given metric",
															MarkdownDescription: "target specifies the target value for the given metric",
															Attributes: map[string]schema.Attribute{
																"average_utilization": schema.Int64Attribute{
																	Description:         "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
																	MarkdownDescription: "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"average_value": schema.StringAttribute{
																	Description:         "averageValue is the target value of the average of the metric across all relevant pods (as a quantity)",
																	MarkdownDescription: "averageValue is the target value of the average of the metric across all relevant pods (as a quantity)",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"type": schema.StringAttribute{
																	Description:         "type represents whether the metric type is Utilization, Value, or AverageValue",
																	MarkdownDescription: "type represents whether the metric type is Utilization, Value, or AverageValue",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"value": schema.StringAttribute{
																	Description:         "value is the target value of the metric (as a quantity).",
																	MarkdownDescription: "value is the target value of the metric (as a quantity).",
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
													Required: false,
													Optional: true,
													Computed: false,
												},

												"pods": schema.SingleNestedAttribute{
													Description:         "pods refers to a metric describing each pod in the current scale target (for example, transactions-processed-per-second).  The values will be averaged together before being compared to the target value.",
													MarkdownDescription: "pods refers to a metric describing each pod in the current scale target (for example, transactions-processed-per-second).  The values will be averaged together before being compared to the target value.",
													Attributes: map[string]schema.Attribute{
														"metric": schema.SingleNestedAttribute{
															Description:         "metric identifies the target metric by name and selector",
															MarkdownDescription: "metric identifies the target metric by name and selector",
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "name is the name of the given metric",
																	MarkdownDescription: "name is the name of the given metric",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"selector": schema.SingleNestedAttribute{
																	Description:         "selector is the string-encoded form of a standard kubernetes label selector for the given metric When set, it is passed as an additional parameter to the metrics server for more specific metrics scoping. When unset, just the metricName will be used to gather metrics.",
																	MarkdownDescription: "selector is the string-encoded form of a standard kubernetes label selector for the given metric When set, it is passed as an additional parameter to the metrics server for more specific metrics scoping. When unset, just the metricName will be used to gather metrics.",
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
																						Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"values": schema.ListAttribute{
																						Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																						MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
																			Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																			MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
															Required: true,
															Optional: false,
															Computed: false,
														},

														"target": schema.SingleNestedAttribute{
															Description:         "target specifies the target value for the given metric",
															MarkdownDescription: "target specifies the target value for the given metric",
															Attributes: map[string]schema.Attribute{
																"average_utilization": schema.Int64Attribute{
																	Description:         "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
																	MarkdownDescription: "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"average_value": schema.StringAttribute{
																	Description:         "averageValue is the target value of the average of the metric across all relevant pods (as a quantity)",
																	MarkdownDescription: "averageValue is the target value of the average of the metric across all relevant pods (as a quantity)",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"type": schema.StringAttribute{
																	Description:         "type represents whether the metric type is Utilization, Value, or AverageValue",
																	MarkdownDescription: "type represents whether the metric type is Utilization, Value, or AverageValue",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"value": schema.StringAttribute{
																	Description:         "value is the target value of the metric (as a quantity).",
																	MarkdownDescription: "value is the target value of the metric (as a quantity).",
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
													Required: false,
													Optional: true,
													Computed: false,
												},

												"resource": schema.SingleNestedAttribute{
													Description:         "resource refers to a resource metric (such as those specified in requests and limits) known to Kubernetes describing each pod in the current scale target (e.g. CPU or memory). Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the 'pods' source.",
													MarkdownDescription: "resource refers to a resource metric (such as those specified in requests and limits) known to Kubernetes describing each pod in the current scale target (e.g. CPU or memory). Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the 'pods' source.",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "name is the name of the resource in question.",
															MarkdownDescription: "name is the name of the resource in question.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"target": schema.SingleNestedAttribute{
															Description:         "target specifies the target value for the given metric",
															MarkdownDescription: "target specifies the target value for the given metric",
															Attributes: map[string]schema.Attribute{
																"average_utilization": schema.Int64Attribute{
																	Description:         "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
																	MarkdownDescription: "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"average_value": schema.StringAttribute{
																	Description:         "averageValue is the target value of the average of the metric across all relevant pods (as a quantity)",
																	MarkdownDescription: "averageValue is the target value of the average of the metric across all relevant pods (as a quantity)",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"type": schema.StringAttribute{
																	Description:         "type represents whether the metric type is Utilization, Value, or AverageValue",
																	MarkdownDescription: "type represents whether the metric type is Utilization, Value, or AverageValue",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"value": schema.StringAttribute{
																	Description:         "value is the target value of the metric (as a quantity).",
																	MarkdownDescription: "value is the target value of the metric (as a quantity).",
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
													Required: false,
													Optional: true,
													Computed: false,
												},

												"type": schema.StringAttribute{
													Description:         "type is the type of metric source.  It should be one of 'ContainerResource', 'External', 'Object', 'Pods' or 'Resource', each mapping to a matching field in the object. Note: 'ContainerResource' type is available on when the feature-gate HPAContainerMetrics is enabled",
													MarkdownDescription: "type is the type of metric source.  It should be one of 'ContainerResource', 'External', 'Object', 'Pods' or 'Resource', each mapping to a matching field in the object. Note: 'ContainerResource' type is available on when the feature-gate HPAContainerMetrics is enabled",
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

									"min_replicas": schema.Int64Attribute{
										Description:         "minReplicas is the lower limit for the number of replicas to which the autoscaler can scale down.  It defaults to 1 pod.  minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External metric is configured.  Scaling is active as long as at least one metric value is available.",
										MarkdownDescription: "minReplicas is the lower limit for the number of replicas to which the autoscaler can scale down.  It defaults to 1 pod.  minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External metric is configured.  Scaling is active as long as at least one metric value is available.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"status": schema.StringAttribute{
										Description:         "Status describe the desired status regarding deploying an horizontal pod autoscaler DISABLED will not deploy an horizontal pod autoscaler ENABLED will deploy an horizontal pod autoscaler",
										MarkdownDescription: "Status describe the desired status regarding deploying an horizontal pod autoscaler DISABLED will not deploy an horizontal pod autoscaler ENABLED will deploy an horizontal pod autoscaler",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("DISABLED", "ENABLED"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"image_pull_policy": schema.StringAttribute{
								Description:         "imagePullPolicy is the Kubernetes pull policy for the image defined above",
								MarkdownDescription: "imagePullPolicy is the Kubernetes pull policy for the image defined above",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("IfNotPresent", "Always", "Never"),
								},
							},

							"log_level": schema.StringAttribute{
								Description:         "logLevel for the console plugin backend",
								MarkdownDescription: "logLevel for the console plugin backend",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("trace", "debug", "info", "warn", "error", "fatal", "panic"),
								},
							},

							"port": schema.Int64Attribute{
								Description:         "port is the plugin service port",
								MarkdownDescription: "port is the plugin service port",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(65535),
								},
							},

							"port_naming": schema.SingleNestedAttribute{
								Description:         "portNaming defines the configuration of the port-to-service name translation",
								MarkdownDescription: "portNaming defines the configuration of the port-to-service name translation",
								Attributes: map[string]schema.Attribute{
									"enable": schema.BoolAttribute{
										Description:         "enable the console plugin port-to-service name translation",
										MarkdownDescription: "enable the console plugin port-to-service name translation",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"port_names": schema.MapAttribute{
										Description:         "portNames defines additional port names to use in the console. Example: portNames: {'3100': 'loki'}",
										MarkdownDescription: "portNames defines additional port names to use in the console. Example: portNames: {'3100': 'loki'}",
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

							"quick_filters": schema.ListNestedAttribute{
								Description:         "quickFilters configures quick filter presets for the Console plugin",
								MarkdownDescription: "quickFilters configures quick filter presets for the Console plugin",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"default": schema.BoolAttribute{
											Description:         "default defines whether this filter should be active by default or not",
											MarkdownDescription: "default defines whether this filter should be active by default or not",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"filter": schema.MapAttribute{
											Description:         "filter is a set of keys and values to be set when this filter is selected. Each key can relate to a list of values using a coma-separated string. Example: filter: {'src_namespace': 'namespace1,namespace2'}",
											MarkdownDescription: "filter is a set of keys and values to be set when this filter is selected. Each key can relate to a list of values using a coma-separated string. Example: filter: {'src_namespace': 'namespace1,namespace2'}",
											ElementType:         types.StringType,
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "name of the filter, that will be displayed in Console",
											MarkdownDescription: "name of the filter, that will be displayed in Console",
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

							"register": schema.BoolAttribute{
								Description:         "register allows, when set to true, to automatically register the provided console plugin with the OpenShift Console operator. When set to false, you can still register it manually by editing console.operator.openshift.io/cluster. E.g: oc patch console.operator.openshift.io cluster --type='json' -p '[{'op': 'add', 'path': '/spec/plugins/-', 'value': 'netobserv-plugin'}]'",
								MarkdownDescription: "register allows, when set to true, to automatically register the provided console plugin with the OpenShift Console operator. When set to false, you can still register it manually by editing console.operator.openshift.io/cluster. E.g: oc patch console.operator.openshift.io cluster --type='json' -p '[{'op': 'add', 'path': '/spec/plugins/-', 'value': 'netobserv-plugin'}]'",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"replicas": schema.Int64Attribute{
								Description:         "replicas defines the number of replicas (pods) to start.",
								MarkdownDescription: "replicas defines the number of replicas (pods) to start.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
								},
							},

							"resources": schema.SingleNestedAttribute{
								Description:         "resources, in terms of compute resources, required by this container. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "resources, in terms of compute resources, required by this container. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								Attributes: map[string]schema.Attribute{
									"claims": schema.ListNestedAttribute{
										Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
										MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
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

					"deployment_model": schema.StringAttribute{
						Description:         "deploymentModel defines the desired type of deployment for flow processing. Possible values are 'DIRECT' (default) to make the flow processor listening directly from the agents, or 'KAFKA' to make flows sent to a Kafka pipeline before consumption by the processor. Kafka can provide better scalability, resiliency and high availability (for more details, see https://www.redhat.com/en/topics/integration/what-is-apache-kafka).",
						MarkdownDescription: "deploymentModel defines the desired type of deployment for flow processing. Possible values are 'DIRECT' (default) to make the flow processor listening directly from the agents, or 'KAFKA' to make flows sent to a Kafka pipeline before consumption by the processor. Kafka can provide better scalability, resiliency and high availability (for more details, see https://www.redhat.com/en/topics/integration/what-is-apache-kafka).",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("DIRECT", "KAFKA"),
						},
					},

					"exporters": schema.ListNestedAttribute{
						Description:         "exporters defines additional optional exporters for custom consumption or storage. This is an experimental feature. Currently, only KAFKA exporter is available.",
						MarkdownDescription: "exporters defines additional optional exporters for custom consumption or storage. This is an experimental feature. Currently, only KAFKA exporter is available.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"ipfix": schema.SingleNestedAttribute{
									Description:         "IPFIX configuration, such as the IP address and port to send enriched IPFIX flows to.",
									MarkdownDescription: "IPFIX configuration, such as the IP address and port to send enriched IPFIX flows to.",
									Attributes: map[string]schema.Attribute{
										"target_host": schema.StringAttribute{
											Description:         "Address of the IPFIX external receiver",
											MarkdownDescription: "Address of the IPFIX external receiver",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"target_port": schema.Int64Attribute{
											Description:         "Port for the IPFIX external receiver",
											MarkdownDescription: "Port for the IPFIX external receiver",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"transport": schema.StringAttribute{
											Description:         "Transport protocol ('TCP' or 'UDP') to be used for the IPFIX connection, defaults to 'TCP'.",
											MarkdownDescription: "Transport protocol ('TCP' or 'UDP') to be used for the IPFIX connection, defaults to 'TCP'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("TCP", "UDP"),
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"kafka": schema.SingleNestedAttribute{
									Description:         "kafka configuration, such as address or topic, to send enriched flows to.",
									MarkdownDescription: "kafka configuration, such as address or topic, to send enriched flows to.",
									Attributes: map[string]schema.Attribute{
										"address": schema.StringAttribute{
											Description:         "address of the Kafka server",
											MarkdownDescription: "address of the Kafka server",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"sasl": schema.SingleNestedAttribute{
											Description:         "SASL authentication configuration. [Unsupported (*)].",
											MarkdownDescription: "SASL authentication configuration. [Unsupported (*)].",
											Attributes: map[string]schema.Attribute{
												"client_id_reference": schema.SingleNestedAttribute{
													Description:         "Reference to the secret or config map containing the client ID",
													MarkdownDescription: "Reference to the secret or config map containing the client ID",
													Attributes: map[string]schema.Attribute{
														"file": schema.StringAttribute{
															Description:         "File name within the config map or secret",
															MarkdownDescription: "File name within the config map or secret",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the config map or secret containing the file",
															MarkdownDescription: "Name of the config map or secret containing the file",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespace": schema.StringAttribute{
															Description:         "Namespace of the config map or secret containing the file. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
															MarkdownDescription: "Namespace of the config map or secret containing the file. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"type": schema.StringAttribute{
															Description:         "Type for the file reference: 'configmap' or 'secret'",
															MarkdownDescription: "Type for the file reference: 'configmap' or 'secret'",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("configmap", "secret"),
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"client_secret_reference": schema.SingleNestedAttribute{
													Description:         "Reference to the secret or config map containing the client secret",
													MarkdownDescription: "Reference to the secret or config map containing the client secret",
													Attributes: map[string]schema.Attribute{
														"file": schema.StringAttribute{
															Description:         "File name within the config map or secret",
															MarkdownDescription: "File name within the config map or secret",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the config map or secret containing the file",
															MarkdownDescription: "Name of the config map or secret containing the file",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespace": schema.StringAttribute{
															Description:         "Namespace of the config map or secret containing the file. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
															MarkdownDescription: "Namespace of the config map or secret containing the file. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"type": schema.StringAttribute{
															Description:         "Type for the file reference: 'configmap' or 'secret'",
															MarkdownDescription: "Type for the file reference: 'configmap' or 'secret'",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("configmap", "secret"),
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"type": schema.StringAttribute{
													Description:         "Type of SASL authentication to use, or 'DISABLED' if SASL is not used",
													MarkdownDescription: "Type of SASL authentication to use, or 'DISABLED' if SASL is not used",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("DISABLED", "PLAIN", "SCRAM-SHA512"),
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"tls": schema.SingleNestedAttribute{
											Description:         "tls client configuration. When using TLS, verify that the address matches the Kafka port used for TLS, generally 9093. Note that, when eBPF agents are used, Kafka certificate needs to be copied in the agent namespace (by default it's netobserv-privileged).",
											MarkdownDescription: "tls client configuration. When using TLS, verify that the address matches the Kafka port used for TLS, generally 9093. Note that, when eBPF agents are used, Kafka certificate needs to be copied in the agent namespace (by default it's netobserv-privileged).",
											Attributes: map[string]schema.Attribute{
												"ca_cert": schema.SingleNestedAttribute{
													Description:         "caCert defines the reference of the certificate for the Certificate Authority",
													MarkdownDescription: "caCert defines the reference of the certificate for the Certificate Authority",
													Attributes: map[string]schema.Attribute{
														"cert_file": schema.StringAttribute{
															Description:         "certFile defines the path to the certificate file name within the config map or secret",
															MarkdownDescription: "certFile defines the path to the certificate file name within the config map or secret",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"cert_key": schema.StringAttribute{
															Description:         "certKey defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
															MarkdownDescription: "certKey defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "name of the config map or secret containing certificates",
															MarkdownDescription: "name of the config map or secret containing certificates",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespace": schema.StringAttribute{
															Description:         "namespace of the config map or secret containing certificates. If omitted, assumes same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret will be copied so that it can be mounted as required.",
															MarkdownDescription: "namespace of the config map or secret containing certificates. If omitted, assumes same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret will be copied so that it can be mounted as required.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"type": schema.StringAttribute{
															Description:         "type for the certificate reference: 'configmap' or 'secret'",
															MarkdownDescription: "type for the certificate reference: 'configmap' or 'secret'",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("configmap", "secret"),
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"enable": schema.BoolAttribute{
													Description:         "enable TLS",
													MarkdownDescription: "enable TLS",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"insecure_skip_verify": schema.BoolAttribute{
													Description:         "insecureSkipVerify allows skipping client-side verification of the server certificate If set to true, CACert field will be ignored",
													MarkdownDescription: "insecureSkipVerify allows skipping client-side verification of the server certificate If set to true, CACert field will be ignored",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"user_cert": schema.SingleNestedAttribute{
													Description:         "userCert defines the user certificate reference, used for mTLS (you can ignore it when using regular, one-way TLS)",
													MarkdownDescription: "userCert defines the user certificate reference, used for mTLS (you can ignore it when using regular, one-way TLS)",
													Attributes: map[string]schema.Attribute{
														"cert_file": schema.StringAttribute{
															Description:         "certFile defines the path to the certificate file name within the config map or secret",
															MarkdownDescription: "certFile defines the path to the certificate file name within the config map or secret",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"cert_key": schema.StringAttribute{
															Description:         "certKey defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
															MarkdownDescription: "certKey defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "name of the config map or secret containing certificates",
															MarkdownDescription: "name of the config map or secret containing certificates",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespace": schema.StringAttribute{
															Description:         "namespace of the config map or secret containing certificates. If omitted, assumes same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret will be copied so that it can be mounted as required.",
															MarkdownDescription: "namespace of the config map or secret containing certificates. If omitted, assumes same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret will be copied so that it can be mounted as required.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"type": schema.StringAttribute{
															Description:         "type for the certificate reference: 'configmap' or 'secret'",
															MarkdownDescription: "type for the certificate reference: 'configmap' or 'secret'",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("configmap", "secret"),
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

										"topic": schema.StringAttribute{
											Description:         "kafka topic to use. It must exist, NetObserv will not create it.",
											MarkdownDescription: "kafka topic to use. It must exist, NetObserv will not create it.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"type": schema.StringAttribute{
									Description:         "'type' selects the type of exporters. The available options are 'KAFKA' and 'IPFIX'.",
									MarkdownDescription: "'type' selects the type of exporters. The available options are 'KAFKA' and 'IPFIX'.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("KAFKA", "IPFIX"),
									},
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"kafka": schema.SingleNestedAttribute{
						Description:         "kafka configuration, allowing to use Kafka as a broker as part of the flow collection pipeline. Available when the 'spec.deploymentModel' is 'KAFKA'.",
						MarkdownDescription: "kafka configuration, allowing to use Kafka as a broker as part of the flow collection pipeline. Available when the 'spec.deploymentModel' is 'KAFKA'.",
						Attributes: map[string]schema.Attribute{
							"address": schema.StringAttribute{
								Description:         "address of the Kafka server",
								MarkdownDescription: "address of the Kafka server",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"sasl": schema.SingleNestedAttribute{
								Description:         "SASL authentication configuration. [Unsupported (*)].",
								MarkdownDescription: "SASL authentication configuration. [Unsupported (*)].",
								Attributes: map[string]schema.Attribute{
									"client_id_reference": schema.SingleNestedAttribute{
										Description:         "Reference to the secret or config map containing the client ID",
										MarkdownDescription: "Reference to the secret or config map containing the client ID",
										Attributes: map[string]schema.Attribute{
											"file": schema.StringAttribute{
												Description:         "File name within the config map or secret",
												MarkdownDescription: "File name within the config map or secret",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the config map or secret containing the file",
												MarkdownDescription: "Name of the config map or secret containing the file",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "Namespace of the config map or secret containing the file. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
												MarkdownDescription: "Namespace of the config map or secret containing the file. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
												Description:         "Type for the file reference: 'configmap' or 'secret'",
												MarkdownDescription: "Type for the file reference: 'configmap' or 'secret'",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("configmap", "secret"),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"client_secret_reference": schema.SingleNestedAttribute{
										Description:         "Reference to the secret or config map containing the client secret",
										MarkdownDescription: "Reference to the secret or config map containing the client secret",
										Attributes: map[string]schema.Attribute{
											"file": schema.StringAttribute{
												Description:         "File name within the config map or secret",
												MarkdownDescription: "File name within the config map or secret",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the config map or secret containing the file",
												MarkdownDescription: "Name of the config map or secret containing the file",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "Namespace of the config map or secret containing the file. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
												MarkdownDescription: "Namespace of the config map or secret containing the file. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
												Description:         "Type for the file reference: 'configmap' or 'secret'",
												MarkdownDescription: "Type for the file reference: 'configmap' or 'secret'",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("configmap", "secret"),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": schema.StringAttribute{
										Description:         "Type of SASL authentication to use, or 'DISABLED' if SASL is not used",
										MarkdownDescription: "Type of SASL authentication to use, or 'DISABLED' if SASL is not used",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("DISABLED", "PLAIN", "SCRAM-SHA512"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"tls": schema.SingleNestedAttribute{
								Description:         "tls client configuration. When using TLS, verify that the address matches the Kafka port used for TLS, generally 9093. Note that, when eBPF agents are used, Kafka certificate needs to be copied in the agent namespace (by default it's netobserv-privileged).",
								MarkdownDescription: "tls client configuration. When using TLS, verify that the address matches the Kafka port used for TLS, generally 9093. Note that, when eBPF agents are used, Kafka certificate needs to be copied in the agent namespace (by default it's netobserv-privileged).",
								Attributes: map[string]schema.Attribute{
									"ca_cert": schema.SingleNestedAttribute{
										Description:         "caCert defines the reference of the certificate for the Certificate Authority",
										MarkdownDescription: "caCert defines the reference of the certificate for the Certificate Authority",
										Attributes: map[string]schema.Attribute{
											"cert_file": schema.StringAttribute{
												Description:         "certFile defines the path to the certificate file name within the config map or secret",
												MarkdownDescription: "certFile defines the path to the certificate file name within the config map or secret",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"cert_key": schema.StringAttribute{
												Description:         "certKey defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
												MarkdownDescription: "certKey defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "name of the config map or secret containing certificates",
												MarkdownDescription: "name of the config map or secret containing certificates",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "namespace of the config map or secret containing certificates. If omitted, assumes same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret will be copied so that it can be mounted as required.",
												MarkdownDescription: "namespace of the config map or secret containing certificates. If omitted, assumes same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret will be copied so that it can be mounted as required.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
												Description:         "type for the certificate reference: 'configmap' or 'secret'",
												MarkdownDescription: "type for the certificate reference: 'configmap' or 'secret'",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("configmap", "secret"),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"enable": schema.BoolAttribute{
										Description:         "enable TLS",
										MarkdownDescription: "enable TLS",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"insecure_skip_verify": schema.BoolAttribute{
										Description:         "insecureSkipVerify allows skipping client-side verification of the server certificate If set to true, CACert field will be ignored",
										MarkdownDescription: "insecureSkipVerify allows skipping client-side verification of the server certificate If set to true, CACert field will be ignored",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"user_cert": schema.SingleNestedAttribute{
										Description:         "userCert defines the user certificate reference, used for mTLS (you can ignore it when using regular, one-way TLS)",
										MarkdownDescription: "userCert defines the user certificate reference, used for mTLS (you can ignore it when using regular, one-way TLS)",
										Attributes: map[string]schema.Attribute{
											"cert_file": schema.StringAttribute{
												Description:         "certFile defines the path to the certificate file name within the config map or secret",
												MarkdownDescription: "certFile defines the path to the certificate file name within the config map or secret",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"cert_key": schema.StringAttribute{
												Description:         "certKey defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
												MarkdownDescription: "certKey defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "name of the config map or secret containing certificates",
												MarkdownDescription: "name of the config map or secret containing certificates",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "namespace of the config map or secret containing certificates. If omitted, assumes same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret will be copied so that it can be mounted as required.",
												MarkdownDescription: "namespace of the config map or secret containing certificates. If omitted, assumes same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret will be copied so that it can be mounted as required.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
												Description:         "type for the certificate reference: 'configmap' or 'secret'",
												MarkdownDescription: "type for the certificate reference: 'configmap' or 'secret'",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("configmap", "secret"),
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

							"topic": schema.StringAttribute{
								Description:         "kafka topic to use. It must exist, NetObserv will not create it.",
								MarkdownDescription: "kafka topic to use. It must exist, NetObserv will not create it.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"loki": schema.SingleNestedAttribute{
						Description:         "loki, the flow store, client settings.",
						MarkdownDescription: "loki, the flow store, client settings.",
						Attributes: map[string]schema.Attribute{
							"auth_token": schema.StringAttribute{
								Description:         "AuthToken describe the way to get a token to authenticate to Loki. DISABLED will not send any token with the request. HOST will use the local pod service account to authenticate to Loki. FORWARD will forward user token, in this mode, pod that are not receiving user request like the processor will use the local pod service account. Similar to HOST mode. When using the Loki Operator, set it to 'HOST' or 'FORWARD'.",
								MarkdownDescription: "AuthToken describe the way to get a token to authenticate to Loki. DISABLED will not send any token with the request. HOST will use the local pod service account to authenticate to Loki. FORWARD will forward user token, in this mode, pod that are not receiving user request like the processor will use the local pod service account. Similar to HOST mode. When using the Loki Operator, set it to 'HOST' or 'FORWARD'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("DISABLED", "HOST", "FORWARD"),
								},
							},

							"batch_size": schema.Int64Attribute{
								Description:         "batchSize is max batch size (in bytes) of logs to accumulate before sending.",
								MarkdownDescription: "batchSize is max batch size (in bytes) of logs to accumulate before sending.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
								},
							},

							"batch_wait": schema.StringAttribute{
								Description:         "batchWait is max time to wait before sending a batch.",
								MarkdownDescription: "batchWait is max time to wait before sending a batch.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_backoff": schema.StringAttribute{
								Description:         "maxBackoff is the maximum backoff time for client connection between retries.",
								MarkdownDescription: "maxBackoff is the maximum backoff time for client connection between retries.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_retries": schema.Int64Attribute{
								Description:         "maxRetries is the maximum number of retries for client connections.",
								MarkdownDescription: "maxRetries is the maximum number of retries for client connections.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
								},
							},

							"min_backoff": schema.StringAttribute{
								Description:         "minBackoff is the initial backoff time for client connection between retries.",
								MarkdownDescription: "minBackoff is the initial backoff time for client connection between retries.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"querier_url": schema.StringAttribute{
								Description:         "querierURL specifies the address of the Loki querier service, in case it is different from the Loki ingester URL. If empty, the URL value will be used (assuming that the Loki ingester and querier are in the same server). When using the Loki Operator, do not set it, since ingestion and queries use the Loki gateway.",
								MarkdownDescription: "querierURL specifies the address of the Loki querier service, in case it is different from the Loki ingester URL. If empty, the URL value will be used (assuming that the Loki ingester and querier are in the same server). When using the Loki Operator, do not set it, since ingestion and queries use the Loki gateway.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"static_labels": schema.MapAttribute{
								Description:         "staticLabels is a map of common labels to set on each flow.",
								MarkdownDescription: "staticLabels is a map of common labels to set on each flow.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"status_url": schema.StringAttribute{
								Description:         "statusURL specifies the address of the Loki /ready /metrics /config endpoints, in case it is different from the Loki querier URL. If empty, the QuerierURL value will be used. This is useful to show error messages and some context in the frontend. When using the Loki Operator, set it to the Loki HTTP query frontend service, for example https://loki-query-frontend-http.netobserv.svc:3100/.",
								MarkdownDescription: "statusURL specifies the address of the Loki /ready /metrics /config endpoints, in case it is different from the Loki querier URL. If empty, the QuerierURL value will be used. This is useful to show error messages and some context in the frontend. When using the Loki Operator, set it to the Loki HTTP query frontend service, for example https://loki-query-frontend-http.netobserv.svc:3100/.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tenant_id": schema.StringAttribute{
								Description:         "tenantID is the Loki X-Scope-OrgID that identifies the tenant for each request. When using the Loki Operator, set it to 'network', which corresponds to a special tenant mode.",
								MarkdownDescription: "tenantID is the Loki X-Scope-OrgID that identifies the tenant for each request. When using the Loki Operator, set it to 'network', which corresponds to a special tenant mode.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"timeout": schema.StringAttribute{
								Description:         "timeout is the maximum time connection / request limit. A Timeout of zero means no timeout.",
								MarkdownDescription: "timeout is the maximum time connection / request limit. A Timeout of zero means no timeout.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tls": schema.SingleNestedAttribute{
								Description:         "tls client configuration.",
								MarkdownDescription: "tls client configuration.",
								Attributes: map[string]schema.Attribute{
									"ca_cert": schema.SingleNestedAttribute{
										Description:         "caCert defines the reference of the certificate for the Certificate Authority",
										MarkdownDescription: "caCert defines the reference of the certificate for the Certificate Authority",
										Attributes: map[string]schema.Attribute{
											"cert_file": schema.StringAttribute{
												Description:         "certFile defines the path to the certificate file name within the config map or secret",
												MarkdownDescription: "certFile defines the path to the certificate file name within the config map or secret",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"cert_key": schema.StringAttribute{
												Description:         "certKey defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
												MarkdownDescription: "certKey defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "name of the config map or secret containing certificates",
												MarkdownDescription: "name of the config map or secret containing certificates",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "namespace of the config map or secret containing certificates. If omitted, assumes same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret will be copied so that it can be mounted as required.",
												MarkdownDescription: "namespace of the config map or secret containing certificates. If omitted, assumes same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret will be copied so that it can be mounted as required.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
												Description:         "type for the certificate reference: 'configmap' or 'secret'",
												MarkdownDescription: "type for the certificate reference: 'configmap' or 'secret'",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("configmap", "secret"),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"enable": schema.BoolAttribute{
										Description:         "enable TLS",
										MarkdownDescription: "enable TLS",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"insecure_skip_verify": schema.BoolAttribute{
										Description:         "insecureSkipVerify allows skipping client-side verification of the server certificate If set to true, CACert field will be ignored",
										MarkdownDescription: "insecureSkipVerify allows skipping client-side verification of the server certificate If set to true, CACert field will be ignored",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"user_cert": schema.SingleNestedAttribute{
										Description:         "userCert defines the user certificate reference, used for mTLS (you can ignore it when using regular, one-way TLS)",
										MarkdownDescription: "userCert defines the user certificate reference, used for mTLS (you can ignore it when using regular, one-way TLS)",
										Attributes: map[string]schema.Attribute{
											"cert_file": schema.StringAttribute{
												Description:         "certFile defines the path to the certificate file name within the config map or secret",
												MarkdownDescription: "certFile defines the path to the certificate file name within the config map or secret",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"cert_key": schema.StringAttribute{
												Description:         "certKey defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
												MarkdownDescription: "certKey defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "name of the config map or secret containing certificates",
												MarkdownDescription: "name of the config map or secret containing certificates",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "namespace of the config map or secret containing certificates. If omitted, assumes same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret will be copied so that it can be mounted as required.",
												MarkdownDescription: "namespace of the config map or secret containing certificates. If omitted, assumes same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret will be copied so that it can be mounted as required.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
												Description:         "type for the certificate reference: 'configmap' or 'secret'",
												MarkdownDescription: "type for the certificate reference: 'configmap' or 'secret'",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("configmap", "secret"),
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

							"url": schema.StringAttribute{
								Description:         "url is the address of an existing Loki service to push the flows to. When using the Loki Operator, set it to the Loki gateway service with the 'network' tenant set in path, for example https://loki-gateway-http.netobserv.svc:8080/api/logs/v1/network.",
								MarkdownDescription: "url is the address of an existing Loki service to push the flows to. When using the Loki Operator, set it to the Loki gateway service with the 'network' tenant set in path, for example https://loki-gateway-http.netobserv.svc:8080/api/logs/v1/network.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"namespace": schema.StringAttribute{
						Description:         "namespace where NetObserv pods are deployed. If empty, the namespace of the operator is going to be used.",
						MarkdownDescription: "namespace where NetObserv pods are deployed. If empty, the namespace of the operator is going to be used.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"processor": schema.SingleNestedAttribute{
						Description:         "processor defines the settings of the component that receives the flows from the agent, enriches them, and forwards them to the Loki persistence layer.",
						MarkdownDescription: "processor defines the settings of the component that receives the flows from the agent, enriches them, and forwards them to the Loki persistence layer.",
						Attributes: map[string]schema.Attribute{
							"debug": schema.SingleNestedAttribute{
								Description:         "Debug allows setting some aspects of the internal configuration of the flow processor. This section is aimed exclusively for debugging and fine-grained performance optimizations (for example GOGC, GOMAXPROCS env vars). Users setting its values do it at their own risk.",
								MarkdownDescription: "Debug allows setting some aspects of the internal configuration of the flow processor. This section is aimed exclusively for debugging and fine-grained performance optimizations (for example GOGC, GOMAXPROCS env vars). Users setting its values do it at their own risk.",
								Attributes: map[string]schema.Attribute{
									"env": schema.MapAttribute{
										Description:         "env allows passing custom environment variables to the NetObserv Agent. Useful for passing some very concrete performance-tuning options (such as GOGC, GOMAXPROCS) that shouldn't be publicly exposed as part of the FlowCollector descriptor, as they are only useful in edge debug and support scenarios.",
										MarkdownDescription: "env allows passing custom environment variables to the NetObserv Agent. Useful for passing some very concrete performance-tuning options (such as GOGC, GOMAXPROCS) that shouldn't be publicly exposed as part of the FlowCollector descriptor, as they are only useful in edge debug and support scenarios.",
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

							"drop_unused_fields": schema.BoolAttribute{
								Description:         "dropUnusedFields allows, when set to true, to drop fields that are known to be unused by OVS, in order to save storage space.",
								MarkdownDescription: "dropUnusedFields allows, when set to true, to drop fields that are known to be unused by OVS, in order to save storage space.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_kube_probes": schema.BoolAttribute{
								Description:         "enableKubeProbes is a flag to enable or disable Kubernetes liveness and readiness probes",
								MarkdownDescription: "enableKubeProbes is a flag to enable or disable Kubernetes liveness and readiness probes",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"health_port": schema.Int64Attribute{
								Description:         "healthPort is a collector HTTP port in the Pod that exposes the health check API",
								MarkdownDescription: "healthPort is a collector HTTP port in the Pod that exposes the health check API",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(65535),
								},
							},

							"image_pull_policy": schema.StringAttribute{
								Description:         "imagePullPolicy is the Kubernetes pull policy for the image defined above",
								MarkdownDescription: "imagePullPolicy is the Kubernetes pull policy for the image defined above",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("IfNotPresent", "Always", "Never"),
								},
							},

							"kafka_consumer_autoscaler": schema.SingleNestedAttribute{
								Description:         "kafkaConsumerAutoscaler spec of a horizontal pod autoscaler to set up for flowlogs-pipeline-transformer, which consumes Kafka messages. This setting is ignored when Kafka is disabled.",
								MarkdownDescription: "kafkaConsumerAutoscaler spec of a horizontal pod autoscaler to set up for flowlogs-pipeline-transformer, which consumes Kafka messages. This setting is ignored when Kafka is disabled.",
								Attributes: map[string]schema.Attribute{
									"max_replicas": schema.Int64Attribute{
										Description:         "maxReplicas is the upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas.",
										MarkdownDescription: "maxReplicas is the upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"metrics": schema.ListNestedAttribute{
										Description:         "metrics used by the pod autoscaler",
										MarkdownDescription: "metrics used by the pod autoscaler",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"container_resource": schema.SingleNestedAttribute{
													Description:         "containerResource refers to a resource metric (such as those specified in requests and limits) known to Kubernetes describing a single container in each pod of the current scale target (e.g. CPU or memory). Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the 'pods' source. This is an alpha feature and can be enabled by the HPAContainerMetrics feature flag.",
													MarkdownDescription: "containerResource refers to a resource metric (such as those specified in requests and limits) known to Kubernetes describing a single container in each pod of the current scale target (e.g. CPU or memory). Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the 'pods' source. This is an alpha feature and can be enabled by the HPAContainerMetrics feature flag.",
													Attributes: map[string]schema.Attribute{
														"container": schema.StringAttribute{
															Description:         "container is the name of the container in the pods of the scaling target",
															MarkdownDescription: "container is the name of the container in the pods of the scaling target",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "name is the name of the resource in question.",
															MarkdownDescription: "name is the name of the resource in question.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"target": schema.SingleNestedAttribute{
															Description:         "target specifies the target value for the given metric",
															MarkdownDescription: "target specifies the target value for the given metric",
															Attributes: map[string]schema.Attribute{
																"average_utilization": schema.Int64Attribute{
																	Description:         "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
																	MarkdownDescription: "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"average_value": schema.StringAttribute{
																	Description:         "averageValue is the target value of the average of the metric across all relevant pods (as a quantity)",
																	MarkdownDescription: "averageValue is the target value of the average of the metric across all relevant pods (as a quantity)",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"type": schema.StringAttribute{
																	Description:         "type represents whether the metric type is Utilization, Value, or AverageValue",
																	MarkdownDescription: "type represents whether the metric type is Utilization, Value, or AverageValue",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"value": schema.StringAttribute{
																	Description:         "value is the target value of the metric (as a quantity).",
																	MarkdownDescription: "value is the target value of the metric (as a quantity).",
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
													Required: false,
													Optional: true,
													Computed: false,
												},

												"external": schema.SingleNestedAttribute{
													Description:         "external refers to a global metric that is not associated with any Kubernetes object. It allows autoscaling based on information coming from components running outside of cluster (for example length of queue in cloud messaging service, or QPS from loadbalancer running outside of cluster).",
													MarkdownDescription: "external refers to a global metric that is not associated with any Kubernetes object. It allows autoscaling based on information coming from components running outside of cluster (for example length of queue in cloud messaging service, or QPS from loadbalancer running outside of cluster).",
													Attributes: map[string]schema.Attribute{
														"metric": schema.SingleNestedAttribute{
															Description:         "metric identifies the target metric by name and selector",
															MarkdownDescription: "metric identifies the target metric by name and selector",
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "name is the name of the given metric",
																	MarkdownDescription: "name is the name of the given metric",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"selector": schema.SingleNestedAttribute{
																	Description:         "selector is the string-encoded form of a standard kubernetes label selector for the given metric When set, it is passed as an additional parameter to the metrics server for more specific metrics scoping. When unset, just the metricName will be used to gather metrics.",
																	MarkdownDescription: "selector is the string-encoded form of a standard kubernetes label selector for the given metric When set, it is passed as an additional parameter to the metrics server for more specific metrics scoping. When unset, just the metricName will be used to gather metrics.",
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
																						Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"values": schema.ListAttribute{
																						Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																						MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
																			Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																			MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
															Required: true,
															Optional: false,
															Computed: false,
														},

														"target": schema.SingleNestedAttribute{
															Description:         "target specifies the target value for the given metric",
															MarkdownDescription: "target specifies the target value for the given metric",
															Attributes: map[string]schema.Attribute{
																"average_utilization": schema.Int64Attribute{
																	Description:         "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
																	MarkdownDescription: "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"average_value": schema.StringAttribute{
																	Description:         "averageValue is the target value of the average of the metric across all relevant pods (as a quantity)",
																	MarkdownDescription: "averageValue is the target value of the average of the metric across all relevant pods (as a quantity)",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"type": schema.StringAttribute{
																	Description:         "type represents whether the metric type is Utilization, Value, or AverageValue",
																	MarkdownDescription: "type represents whether the metric type is Utilization, Value, or AverageValue",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"value": schema.StringAttribute{
																	Description:         "value is the target value of the metric (as a quantity).",
																	MarkdownDescription: "value is the target value of the metric (as a quantity).",
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
													Required: false,
													Optional: true,
													Computed: false,
												},

												"object": schema.SingleNestedAttribute{
													Description:         "object refers to a metric describing a single kubernetes object (for example, hits-per-second on an Ingress object).",
													MarkdownDescription: "object refers to a metric describing a single kubernetes object (for example, hits-per-second on an Ingress object).",
													Attributes: map[string]schema.Attribute{
														"described_object": schema.SingleNestedAttribute{
															Description:         "describedObject specifies the descriptions of a object,such as kind,name apiVersion",
															MarkdownDescription: "describedObject specifies the descriptions of a object,such as kind,name apiVersion",
															Attributes: map[string]schema.Attribute{
																"api_version": schema.StringAttribute{
																	Description:         "apiVersion is the API version of the referent",
																	MarkdownDescription: "apiVersion is the API version of the referent",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"kind": schema.StringAttribute{
																	Description:         "kind is the kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
																	MarkdownDescription: "kind is the kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "name is the name of the referent; More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	MarkdownDescription: "name is the name of the referent; More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},
															},
															Required: true,
															Optional: false,
															Computed: false,
														},

														"metric": schema.SingleNestedAttribute{
															Description:         "metric identifies the target metric by name and selector",
															MarkdownDescription: "metric identifies the target metric by name and selector",
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "name is the name of the given metric",
																	MarkdownDescription: "name is the name of the given metric",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"selector": schema.SingleNestedAttribute{
																	Description:         "selector is the string-encoded form of a standard kubernetes label selector for the given metric When set, it is passed as an additional parameter to the metrics server for more specific metrics scoping. When unset, just the metricName will be used to gather metrics.",
																	MarkdownDescription: "selector is the string-encoded form of a standard kubernetes label selector for the given metric When set, it is passed as an additional parameter to the metrics server for more specific metrics scoping. When unset, just the metricName will be used to gather metrics.",
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
																						Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"values": schema.ListAttribute{
																						Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																						MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
																			Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																			MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
															Required: true,
															Optional: false,
															Computed: false,
														},

														"target": schema.SingleNestedAttribute{
															Description:         "target specifies the target value for the given metric",
															MarkdownDescription: "target specifies the target value for the given metric",
															Attributes: map[string]schema.Attribute{
																"average_utilization": schema.Int64Attribute{
																	Description:         "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
																	MarkdownDescription: "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"average_value": schema.StringAttribute{
																	Description:         "averageValue is the target value of the average of the metric across all relevant pods (as a quantity)",
																	MarkdownDescription: "averageValue is the target value of the average of the metric across all relevant pods (as a quantity)",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"type": schema.StringAttribute{
																	Description:         "type represents whether the metric type is Utilization, Value, or AverageValue",
																	MarkdownDescription: "type represents whether the metric type is Utilization, Value, or AverageValue",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"value": schema.StringAttribute{
																	Description:         "value is the target value of the metric (as a quantity).",
																	MarkdownDescription: "value is the target value of the metric (as a quantity).",
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
													Required: false,
													Optional: true,
													Computed: false,
												},

												"pods": schema.SingleNestedAttribute{
													Description:         "pods refers to a metric describing each pod in the current scale target (for example, transactions-processed-per-second).  The values will be averaged together before being compared to the target value.",
													MarkdownDescription: "pods refers to a metric describing each pod in the current scale target (for example, transactions-processed-per-second).  The values will be averaged together before being compared to the target value.",
													Attributes: map[string]schema.Attribute{
														"metric": schema.SingleNestedAttribute{
															Description:         "metric identifies the target metric by name and selector",
															MarkdownDescription: "metric identifies the target metric by name and selector",
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "name is the name of the given metric",
																	MarkdownDescription: "name is the name of the given metric",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"selector": schema.SingleNestedAttribute{
																	Description:         "selector is the string-encoded form of a standard kubernetes label selector for the given metric When set, it is passed as an additional parameter to the metrics server for more specific metrics scoping. When unset, just the metricName will be used to gather metrics.",
																	MarkdownDescription: "selector is the string-encoded form of a standard kubernetes label selector for the given metric When set, it is passed as an additional parameter to the metrics server for more specific metrics scoping. When unset, just the metricName will be used to gather metrics.",
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
																						Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"values": schema.ListAttribute{
																						Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																						MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
																			Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																			MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
															Required: true,
															Optional: false,
															Computed: false,
														},

														"target": schema.SingleNestedAttribute{
															Description:         "target specifies the target value for the given metric",
															MarkdownDescription: "target specifies the target value for the given metric",
															Attributes: map[string]schema.Attribute{
																"average_utilization": schema.Int64Attribute{
																	Description:         "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
																	MarkdownDescription: "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"average_value": schema.StringAttribute{
																	Description:         "averageValue is the target value of the average of the metric across all relevant pods (as a quantity)",
																	MarkdownDescription: "averageValue is the target value of the average of the metric across all relevant pods (as a quantity)",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"type": schema.StringAttribute{
																	Description:         "type represents whether the metric type is Utilization, Value, or AverageValue",
																	MarkdownDescription: "type represents whether the metric type is Utilization, Value, or AverageValue",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"value": schema.StringAttribute{
																	Description:         "value is the target value of the metric (as a quantity).",
																	MarkdownDescription: "value is the target value of the metric (as a quantity).",
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
													Required: false,
													Optional: true,
													Computed: false,
												},

												"resource": schema.SingleNestedAttribute{
													Description:         "resource refers to a resource metric (such as those specified in requests and limits) known to Kubernetes describing each pod in the current scale target (e.g. CPU or memory). Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the 'pods' source.",
													MarkdownDescription: "resource refers to a resource metric (such as those specified in requests and limits) known to Kubernetes describing each pod in the current scale target (e.g. CPU or memory). Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the 'pods' source.",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "name is the name of the resource in question.",
															MarkdownDescription: "name is the name of the resource in question.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"target": schema.SingleNestedAttribute{
															Description:         "target specifies the target value for the given metric",
															MarkdownDescription: "target specifies the target value for the given metric",
															Attributes: map[string]schema.Attribute{
																"average_utilization": schema.Int64Attribute{
																	Description:         "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
																	MarkdownDescription: "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"average_value": schema.StringAttribute{
																	Description:         "averageValue is the target value of the average of the metric across all relevant pods (as a quantity)",
																	MarkdownDescription: "averageValue is the target value of the average of the metric across all relevant pods (as a quantity)",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"type": schema.StringAttribute{
																	Description:         "type represents whether the metric type is Utilization, Value, or AverageValue",
																	MarkdownDescription: "type represents whether the metric type is Utilization, Value, or AverageValue",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"value": schema.StringAttribute{
																	Description:         "value is the target value of the metric (as a quantity).",
																	MarkdownDescription: "value is the target value of the metric (as a quantity).",
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
													Required: false,
													Optional: true,
													Computed: false,
												},

												"type": schema.StringAttribute{
													Description:         "type is the type of metric source.  It should be one of 'ContainerResource', 'External', 'Object', 'Pods' or 'Resource', each mapping to a matching field in the object. Note: 'ContainerResource' type is available on when the feature-gate HPAContainerMetrics is enabled",
													MarkdownDescription: "type is the type of metric source.  It should be one of 'ContainerResource', 'External', 'Object', 'Pods' or 'Resource', each mapping to a matching field in the object. Note: 'ContainerResource' type is available on when the feature-gate HPAContainerMetrics is enabled",
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

									"min_replicas": schema.Int64Attribute{
										Description:         "minReplicas is the lower limit for the number of replicas to which the autoscaler can scale down.  It defaults to 1 pod.  minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External metric is configured.  Scaling is active as long as at least one metric value is available.",
										MarkdownDescription: "minReplicas is the lower limit for the number of replicas to which the autoscaler can scale down.  It defaults to 1 pod.  minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External metric is configured.  Scaling is active as long as at least one metric value is available.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"status": schema.StringAttribute{
										Description:         "Status describe the desired status regarding deploying an horizontal pod autoscaler DISABLED will not deploy an horizontal pod autoscaler ENABLED will deploy an horizontal pod autoscaler",
										MarkdownDescription: "Status describe the desired status regarding deploying an horizontal pod autoscaler DISABLED will not deploy an horizontal pod autoscaler ENABLED will deploy an horizontal pod autoscaler",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("DISABLED", "ENABLED"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"kafka_consumer_batch_size": schema.Int64Attribute{
								Description:         "kafkaConsumerBatchSize indicates to the broker the maximum batch size, in bytes, that the consumer will accept. Ignored when not using Kafka. Default: 10MB.",
								MarkdownDescription: "kafkaConsumerBatchSize indicates to the broker the maximum batch size, in bytes, that the consumer will accept. Ignored when not using Kafka. Default: 10MB.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kafka_consumer_queue_capacity": schema.Int64Attribute{
								Description:         "kafkaConsumerQueueCapacity defines the capacity of the internal message queue used in the Kafka consumer client. Ignored when not using Kafka.",
								MarkdownDescription: "kafkaConsumerQueueCapacity defines the capacity of the internal message queue used in the Kafka consumer client. Ignored when not using Kafka.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kafka_consumer_replicas": schema.Int64Attribute{
								Description:         "kafkaConsumerReplicas defines the number of replicas (pods) to start for flowlogs-pipeline-transformer, which consumes Kafka messages. This setting is ignored when Kafka is disabled.",
								MarkdownDescription: "kafkaConsumerReplicas defines the number of replicas (pods) to start for flowlogs-pipeline-transformer, which consumes Kafka messages. This setting is ignored when Kafka is disabled.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
								},
							},

							"log_level": schema.StringAttribute{
								Description:         "logLevel of the collector runtime",
								MarkdownDescription: "logLevel of the collector runtime",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("trace", "debug", "info", "warn", "error", "fatal", "panic"),
								},
							},

							"metrics": schema.SingleNestedAttribute{
								Description:         "Metrics define the processor configuration regarding metrics",
								MarkdownDescription: "Metrics define the processor configuration regarding metrics",
								Attributes: map[string]schema.Attribute{
									"ignore_tags": schema.ListAttribute{
										Description:         "ignoreTags is a list of tags to specify which metrics to ignore. Each metric is associated with a list of tags. More details in https://github.com/netobserv/network-observability-operator/tree/main/controllers/flowlogspipeline/metrics_definitions . Available tags are: egress, ingress, flows, bytes, packets, namespaces, nodes, workloads",
										MarkdownDescription: "ignoreTags is a list of tags to specify which metrics to ignore. Each metric is associated with a list of tags. More details in https://github.com/netobserv/network-observability-operator/tree/main/controllers/flowlogspipeline/metrics_definitions . Available tags are: egress, ingress, flows, bytes, packets, namespaces, nodes, workloads",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"server": schema.SingleNestedAttribute{
										Description:         "metricsServer endpoint configuration for Prometheus scraper",
										MarkdownDescription: "metricsServer endpoint configuration for Prometheus scraper",
										Attributes: map[string]schema.Attribute{
											"port": schema.Int64Attribute{
												Description:         "the prometheus HTTP port",
												MarkdownDescription: "the prometheus HTTP port",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
													int64validator.AtMost(65535),
												},
											},

											"tls": schema.SingleNestedAttribute{
												Description:         "TLS configuration.",
												MarkdownDescription: "TLS configuration.",
												Attributes: map[string]schema.Attribute{
													"provided": schema.SingleNestedAttribute{
														Description:         "TLS configuration.",
														MarkdownDescription: "TLS configuration.",
														Attributes: map[string]schema.Attribute{
															"cert_file": schema.StringAttribute{
																Description:         "certFile defines the path to the certificate file name within the config map or secret",
																MarkdownDescription: "certFile defines the path to the certificate file name within the config map or secret",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"cert_key": schema.StringAttribute{
																Description:         "certKey defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
																MarkdownDescription: "certKey defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "name of the config map or secret containing certificates",
																MarkdownDescription: "name of the config map or secret containing certificates",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "namespace of the config map or secret containing certificates. If omitted, assumes same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret will be copied so that it can be mounted as required.",
																MarkdownDescription: "namespace of the config map or secret containing certificates. If omitted, assumes same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret will be copied so that it can be mounted as required.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"type": schema.StringAttribute{
																Description:         "type for the certificate reference: 'configmap' or 'secret'",
																MarkdownDescription: "type for the certificate reference: 'configmap' or 'secret'",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("configmap", "secret"),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"type": schema.StringAttribute{
														Description:         "Select the type of TLS configuration 'DISABLED' (default) to not configure TLS for the endpoint, 'PROVIDED' to manually provide cert file and a key file, and 'AUTO' to use OpenShift auto generated certificate using annotations",
														MarkdownDescription: "Select the type of TLS configuration 'DISABLED' (default) to not configure TLS for the endpoint, 'PROVIDED' to manually provide cert file and a key file, and 'AUTO' to use OpenShift auto generated certificate using annotations",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("DISABLED", "PROVIDED", "AUTO"),
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

							"port": schema.Int64Attribute{
								Description:         "port of the flow collector (host port) By conventions, some value are not authorized port must not be below 1024 and must not equal this values: 4789,6081,500, and 4500",
								MarkdownDescription: "port of the flow collector (host port) By conventions, some value are not authorized port must not be below 1024 and must not equal this values: 4789,6081,500, and 4500",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1025),
									int64validator.AtMost(65535),
								},
							},

							"profile_port": schema.Int64Attribute{
								Description:         "profilePort allows setting up a Go pprof profiler listening to this port",
								MarkdownDescription: "profilePort allows setting up a Go pprof profiler listening to this port",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
									int64validator.AtMost(65535),
								},
							},

							"resources": schema.SingleNestedAttribute{
								Description:         "resources are the compute resources required by this container. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "resources are the compute resources required by this container. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								Attributes: map[string]schema.Attribute{
									"claims": schema.ListNestedAttribute{
										Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
										MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *FlowsNetobservIoFlowCollectorV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_flows_netobserv_io_flow_collector_v1alpha1_manifest")

	var model FlowsNetobservIoFlowCollectorV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("flows.netobserv.io/v1alpha1")
	model.Kind = pointer.String("FlowCollector")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
