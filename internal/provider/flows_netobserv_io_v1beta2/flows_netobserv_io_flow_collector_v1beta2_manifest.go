/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package flows_netobserv_io_v1beta2

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
	_ datasource.DataSource = &FlowsNetobservIoFlowCollectorV1Beta2Manifest{}
)

func NewFlowsNetobservIoFlowCollectorV1Beta2Manifest() datasource.DataSource {
	return &FlowsNetobservIoFlowCollectorV1Beta2Manifest{}
}

type FlowsNetobservIoFlowCollectorV1Beta2Manifest struct{}

type FlowsNetobservIoFlowCollectorV1Beta2ManifestData struct {
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
				Advanced *struct {
					Env *map[string]string `tfsdk:"env" json:"env,omitempty"`
				} `tfsdk:"advanced" json:"advanced,omitempty"`
				CacheActiveTimeout *string   `tfsdk:"cache_active_timeout" json:"cacheActiveTimeout,omitempty"`
				CacheMaxFlows      *int64    `tfsdk:"cache_max_flows" json:"cacheMaxFlows,omitempty"`
				ExcludeInterfaces  *[]string `tfsdk:"exclude_interfaces" json:"excludeInterfaces,omitempty"`
				Features           *[]string `tfsdk:"features" json:"features,omitempty"`
				ImagePullPolicy    *string   `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
				Interfaces         *[]string `tfsdk:"interfaces" json:"interfaces,omitempty"`
				KafkaBatchSize     *int64    `tfsdk:"kafka_batch_size" json:"kafkaBatchSize,omitempty"`
				LogLevel           *string   `tfsdk:"log_level" json:"logLevel,omitempty"`
				Metrics            *struct {
					Enable *bool `tfsdk:"enable" json:"enable,omitempty"`
					Server *struct {
						Port *int64 `tfsdk:"port" json:"port,omitempty"`
						Tls  *struct {
							InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
							Provided           *struct {
								CertFile  *string `tfsdk:"cert_file" json:"certFile,omitempty"`
								CertKey   *string `tfsdk:"cert_key" json:"certKey,omitempty"`
								Name      *string `tfsdk:"name" json:"name,omitempty"`
								Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
								Type      *string `tfsdk:"type" json:"type,omitempty"`
							} `tfsdk:"provided" json:"provided,omitempty"`
							ProvidedCaFile *struct {
								File      *string `tfsdk:"file" json:"file,omitempty"`
								Name      *string `tfsdk:"name" json:"name,omitempty"`
								Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
								Type      *string `tfsdk:"type" json:"type,omitempty"`
							} `tfsdk:"provided_ca_file" json:"providedCaFile,omitempty"`
							Type *string `tfsdk:"type" json:"type,omitempty"`
						} `tfsdk:"tls" json:"tls,omitempty"`
					} `tfsdk:"server" json:"server,omitempty"`
				} `tfsdk:"metrics" json:"metrics,omitempty"`
				Privileged *bool `tfsdk:"privileged" json:"privileged,omitempty"`
				Resources  *struct {
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
			Advanced *struct {
				Args     *[]string          `tfsdk:"args" json:"args,omitempty"`
				Env      *map[string]string `tfsdk:"env" json:"env,omitempty"`
				Port     *int64             `tfsdk:"port" json:"port,omitempty"`
				Register *bool              `tfsdk:"register" json:"register,omitempty"`
			} `tfsdk:"advanced" json:"advanced,omitempty"`
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
			Enable          *bool   `tfsdk:"enable" json:"enable,omitempty"`
			ImagePullPolicy *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
			LogLevel        *string `tfsdk:"log_level" json:"logLevel,omitempty"`
			PortNaming      *struct {
				Enable    *bool              `tfsdk:"enable" json:"enable,omitempty"`
				PortNames *map[string]string `tfsdk:"port_names" json:"portNames,omitempty"`
			} `tfsdk:"port_naming" json:"portNaming,omitempty"`
			QuickFilters *[]struct {
				Default *bool              `tfsdk:"default" json:"default,omitempty"`
				Filter  *map[string]string `tfsdk:"filter" json:"filter,omitempty"`
				Name    *string            `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"quick_filters" json:"quickFilters,omitempty"`
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
			Advanced *struct {
				StaticLabels    *map[string]string `tfsdk:"static_labels" json:"staticLabels,omitempty"`
				WriteMaxBackoff *string            `tfsdk:"write_max_backoff" json:"writeMaxBackoff,omitempty"`
				WriteMaxRetries *int64             `tfsdk:"write_max_retries" json:"writeMaxRetries,omitempty"`
				WriteMinBackoff *string            `tfsdk:"write_min_backoff" json:"writeMinBackoff,omitempty"`
			} `tfsdk:"advanced" json:"advanced,omitempty"`
			Enable    *bool `tfsdk:"enable" json:"enable,omitempty"`
			LokiStack *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"loki_stack" json:"lokiStack,omitempty"`
			Manual *struct {
				AuthToken   *string `tfsdk:"auth_token" json:"authToken,omitempty"`
				IngesterUrl *string `tfsdk:"ingester_url" json:"ingesterUrl,omitempty"`
				QuerierUrl  *string `tfsdk:"querier_url" json:"querierUrl,omitempty"`
				StatusTls   *struct {
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
				} `tfsdk:"status_tls" json:"statusTls,omitempty"`
				StatusUrl *string `tfsdk:"status_url" json:"statusUrl,omitempty"`
				TenantID  *string `tfsdk:"tenant_id" json:"tenantID,omitempty"`
				Tls       *struct {
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
			} `tfsdk:"manual" json:"manual,omitempty"`
			Microservices *struct {
				IngesterUrl *string `tfsdk:"ingester_url" json:"ingesterUrl,omitempty"`
				QuerierUrl  *string `tfsdk:"querier_url" json:"querierUrl,omitempty"`
				TenantID    *string `tfsdk:"tenant_id" json:"tenantID,omitempty"`
				Tls         *struct {
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
			} `tfsdk:"microservices" json:"microservices,omitempty"`
			Mode       *string `tfsdk:"mode" json:"mode,omitempty"`
			Monolithic *struct {
				TenantID *string `tfsdk:"tenant_id" json:"tenantID,omitempty"`
				Tls      *struct {
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
			} `tfsdk:"monolithic" json:"monolithic,omitempty"`
			ReadTimeout    *string `tfsdk:"read_timeout" json:"readTimeout,omitempty"`
			WriteBatchSize *int64  `tfsdk:"write_batch_size" json:"writeBatchSize,omitempty"`
			WriteBatchWait *string `tfsdk:"write_batch_wait" json:"writeBatchWait,omitempty"`
			WriteTimeout   *string `tfsdk:"write_timeout" json:"writeTimeout,omitempty"`
		} `tfsdk:"loki" json:"loki,omitempty"`
		Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		Processor *struct {
			AddZone  *bool `tfsdk:"add_zone" json:"addZone,omitempty"`
			Advanced *struct {
				ConversationEndTimeout         *string            `tfsdk:"conversation_end_timeout" json:"conversationEndTimeout,omitempty"`
				ConversationHeartbeatInterval  *string            `tfsdk:"conversation_heartbeat_interval" json:"conversationHeartbeatInterval,omitempty"`
				ConversationTerminatingTimeout *string            `tfsdk:"conversation_terminating_timeout" json:"conversationTerminatingTimeout,omitempty"`
				DropUnusedFields               *bool              `tfsdk:"drop_unused_fields" json:"dropUnusedFields,omitempty"`
				EnableKubeProbes               *bool              `tfsdk:"enable_kube_probes" json:"enableKubeProbes,omitempty"`
				Env                            *map[string]string `tfsdk:"env" json:"env,omitempty"`
				HealthPort                     *int64             `tfsdk:"health_port" json:"healthPort,omitempty"`
				Port                           *int64             `tfsdk:"port" json:"port,omitempty"`
				ProfilePort                    *int64             `tfsdk:"profile_port" json:"profilePort,omitempty"`
			} `tfsdk:"advanced" json:"advanced,omitempty"`
			ClusterName             *string `tfsdk:"cluster_name" json:"clusterName,omitempty"`
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
			LogTypes                   *string `tfsdk:"log_types" json:"logTypes,omitempty"`
			Metrics                    *struct {
				DisableAlerts *[]string `tfsdk:"disable_alerts" json:"disableAlerts,omitempty"`
				IncludeList   *[]string `tfsdk:"include_list" json:"includeList,omitempty"`
				Server        *struct {
					Port *int64 `tfsdk:"port" json:"port,omitempty"`
					Tls  *struct {
						InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
						Provided           *struct {
							CertFile  *string `tfsdk:"cert_file" json:"certFile,omitempty"`
							CertKey   *string `tfsdk:"cert_key" json:"certKey,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
							Type      *string `tfsdk:"type" json:"type,omitempty"`
						} `tfsdk:"provided" json:"provided,omitempty"`
						ProvidedCaFile *struct {
							File      *string `tfsdk:"file" json:"file,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
							Type      *string `tfsdk:"type" json:"type,omitempty"`
						} `tfsdk:"provided_ca_file" json:"providedCaFile,omitempty"`
						Type *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"tls" json:"tls,omitempty"`
				} `tfsdk:"server" json:"server,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			MultiClusterDeployment *bool `tfsdk:"multi_cluster_deployment" json:"multiClusterDeployment,omitempty"`
			Resources              *struct {
				Claims *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
		} `tfsdk:"processor" json:"processor,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *FlowsNetobservIoFlowCollectorV1Beta2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_flows_netobserv_io_flow_collector_v1beta2_manifest"
}

func (r *FlowsNetobservIoFlowCollectorV1Beta2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "'FlowCollector' is the schema for the network flows collection API, which pilots and configures the underlying deployments.",
		MarkdownDescription: "'FlowCollector' is the schema for the network flows collection API, which pilots and configures the underlying deployments.",
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
				Description:         "Defines the desired state of the FlowCollector resource. <br><br> *: the mention of 'unsupported', or 'deprecated' for a feature throughout this document means that this feature is not officially supported by Red Hat. It might have been, for example, contributed by the community and accepted without a formal agreement for maintenance. The product maintainers might provide some support for these features as a best effort only.",
				MarkdownDescription: "Defines the desired state of the FlowCollector resource. <br><br> *: the mention of 'unsupported', or 'deprecated' for a feature throughout this document means that this feature is not officially supported by Red Hat. It might have been, for example, contributed by the community and accepted without a formal agreement for maintenance. The product maintainers might provide some support for these features as a best effort only.",
				Attributes: map[string]schema.Attribute{
					"agent": schema.SingleNestedAttribute{
						Description:         "Agent configuration for flows extraction.",
						MarkdownDescription: "Agent configuration for flows extraction.",
						Attributes: map[string]schema.Attribute{
							"ebpf": schema.SingleNestedAttribute{
								Description:         "'ebpf' describes the settings related to the eBPF-based flow reporter when 'spec.agent.type' is set to 'eBPF'.",
								MarkdownDescription: "'ebpf' describes the settings related to the eBPF-based flow reporter when 'spec.agent.type' is set to 'eBPF'.",
								Attributes: map[string]schema.Attribute{
									"advanced": schema.SingleNestedAttribute{
										Description:         "'advanced' allows setting some aspects of the internal configuration of the eBPF agent. This section is aimed mostly for debugging and fine-grained performance optimizations, such as 'GOGC' and 'GOMAXPROCS' env vars. Set these values at your own risk.",
										MarkdownDescription: "'advanced' allows setting some aspects of the internal configuration of the eBPF agent. This section is aimed mostly for debugging and fine-grained performance optimizations, such as 'GOGC' and 'GOMAXPROCS' env vars. Set these values at your own risk.",
										Attributes: map[string]schema.Attribute{
											"env": schema.MapAttribute{
												Description:         "'env' allows passing custom environment variables to underlying components. Useful for passing some very concrete performance-tuning options, such as 'GOGC' and 'GOMAXPROCS', that should not be publicly exposed as part of the FlowCollector descriptor, as they are only useful in edge debug or support scenarios.",
												MarkdownDescription: "'env' allows passing custom environment variables to underlying components. Useful for passing some very concrete performance-tuning options, such as 'GOGC' and 'GOMAXPROCS', that should not be publicly exposed as part of the FlowCollector descriptor, as they are only useful in edge debug or support scenarios.",
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

									"cache_active_timeout": schema.StringAttribute{
										Description:         "'cacheActiveTimeout' is the max period during which the reporter aggregates flows before sending. Increasing 'cacheMaxFlows' and 'cacheActiveTimeout' can decrease the network traffic overhead and the CPU load, however you can expect higher memory consumption and an increased latency in the flow collection.",
										MarkdownDescription: "'cacheActiveTimeout' is the max period during which the reporter aggregates flows before sending. Increasing 'cacheMaxFlows' and 'cacheActiveTimeout' can decrease the network traffic overhead and the CPU load, however you can expect higher memory consumption and an increased latency in the flow collection.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(ns|ms|s|m)?$`), ""),
										},
									},

									"cache_max_flows": schema.Int64Attribute{
										Description:         "'cacheMaxFlows' is the max number of flows in an aggregate; when reached, the reporter sends the flows. Increasing 'cacheMaxFlows' and 'cacheActiveTimeout' can decrease the network traffic overhead and the CPU load, however you can expect higher memory consumption and an increased latency in the flow collection.",
										MarkdownDescription: "'cacheMaxFlows' is the max number of flows in an aggregate; when reached, the reporter sends the flows. Increasing 'cacheMaxFlows' and 'cacheActiveTimeout' can decrease the network traffic overhead and the CPU load, however you can expect higher memory consumption and an increased latency in the flow collection.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"exclude_interfaces": schema.ListAttribute{
										Description:         "'excludeInterfaces' contains the interface names that are excluded from flow tracing. An entry enclosed by slashes, such as '/br-/', is matched as a regular expression. Otherwise it is matched as a case-sensitive string.",
										MarkdownDescription: "'excludeInterfaces' contains the interface names that are excluded from flow tracing. An entry enclosed by slashes, such as '/br-/', is matched as a regular expression. Otherwise it is matched as a case-sensitive string.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"features": schema.ListAttribute{
										Description:         "List of additional features to enable. They are all disabled by default. Enabling additional features might have performance impacts. Possible values are:<br> - 'PacketDrop': enable the packets drop flows logging feature. This feature requires mounting the kernel debug filesystem, so the eBPF pod has to run as privileged. If the 'spec.agent.ebpf.privileged' parameter is not set, an error is reported.<br> - 'DNSTracking': enable the DNS tracking feature.<br> - 'FlowRTT': enable flow latency (RTT) calculations in the eBPF agent during TCP handshakes. This feature better works with 'sampling' set to 1.<br>",
										MarkdownDescription: "List of additional features to enable. They are all disabled by default. Enabling additional features might have performance impacts. Possible values are:<br> - 'PacketDrop': enable the packets drop flows logging feature. This feature requires mounting the kernel debug filesystem, so the eBPF pod has to run as privileged. If the 'spec.agent.ebpf.privileged' parameter is not set, an error is reported.<br> - 'DNSTracking': enable the DNS tracking feature.<br> - 'FlowRTT': enable flow latency (RTT) calculations in the eBPF agent during TCP handshakes. This feature better works with 'sampling' set to 1.<br>",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"image_pull_policy": schema.StringAttribute{
										Description:         "'imagePullPolicy' is the Kubernetes pull policy for the image defined above",
										MarkdownDescription: "'imagePullPolicy' is the Kubernetes pull policy for the image defined above",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("IfNotPresent", "Always", "Never"),
										},
									},

									"interfaces": schema.ListAttribute{
										Description:         "'interfaces' contains the interface names from where flows are collected. If empty, the agent fetches all the interfaces in the system, excepting the ones listed in ExcludeInterfaces. An entry enclosed by slashes, such as '/br-/', is matched as a regular expression. Otherwise it is matched as a case-sensitive string.",
										MarkdownDescription: "'interfaces' contains the interface names from where flows are collected. If empty, the agent fetches all the interfaces in the system, excepting the ones listed in ExcludeInterfaces. An entry enclosed by slashes, such as '/br-/', is matched as a regular expression. Otherwise it is matched as a case-sensitive string.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"kafka_batch_size": schema.Int64Attribute{
										Description:         "'kafkaBatchSize' limits the maximum size of a request in bytes before being sent to a partition. Ignored when not using Kafka. Default: 1MB.",
										MarkdownDescription: "'kafkaBatchSize' limits the maximum size of a request in bytes before being sent to a partition. Ignored when not using Kafka. Default: 1MB.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"log_level": schema.StringAttribute{
										Description:         "'logLevel' defines the log level for the NetObserv eBPF Agent",
										MarkdownDescription: "'logLevel' defines the log level for the NetObserv eBPF Agent",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("trace", "debug", "info", "warn", "error", "fatal", "panic"),
										},
									},

									"metrics": schema.SingleNestedAttribute{
										Description:         "'metrics' defines the eBPF agent configuration regarding metrics",
										MarkdownDescription: "'metrics' defines the eBPF agent configuration regarding metrics",
										Attributes: map[string]schema.Attribute{
											"enable": schema.BoolAttribute{
												Description:         "Set 'enable' to 'true' to enable eBPF agent metrics collection.",
												MarkdownDescription: "Set 'enable' to 'true' to enable eBPF agent metrics collection.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"server": schema.SingleNestedAttribute{
												Description:         "Metrics server endpoint configuration for Prometheus scraper",
												MarkdownDescription: "Metrics server endpoint configuration for Prometheus scraper",
												Attributes: map[string]schema.Attribute{
													"port": schema.Int64Attribute{
														Description:         "The prometheus HTTP port",
														MarkdownDescription: "The prometheus HTTP port",
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
															"insecure_skip_verify": schema.BoolAttribute{
																Description:         "'insecureSkipVerify' allows skipping client-side verification of the provided certificate. If set to 'true', the 'providedCaFile' field is ignored.",
																MarkdownDescription: "'insecureSkipVerify' allows skipping client-side verification of the provided certificate. If set to 'true', the 'providedCaFile' field is ignored.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"provided": schema.SingleNestedAttribute{
																Description:         "TLS configuration when 'type' is set to 'Provided'.",
																MarkdownDescription: "TLS configuration when 'type' is set to 'Provided'.",
																Attributes: map[string]schema.Attribute{
																	"cert_file": schema.StringAttribute{
																		Description:         "'certFile' defines the path to the certificate file name within the config map or secret",
																		MarkdownDescription: "'certFile' defines the path to the certificate file name within the config map or secret",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"cert_key": schema.StringAttribute{
																		Description:         "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
																		MarkdownDescription: "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the config map or secret containing certificates",
																		MarkdownDescription: "Name of the config map or secret containing certificates",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"namespace": schema.StringAttribute{
																		Description:         "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
																		MarkdownDescription: "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"type": schema.StringAttribute{
																		Description:         "Type for the certificate reference: 'configmap' or 'secret'",
																		MarkdownDescription: "Type for the certificate reference: 'configmap' or 'secret'",
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

															"provided_ca_file": schema.SingleNestedAttribute{
																Description:         "Reference to the CA file when 'type' is set to 'Provided'.",
																MarkdownDescription: "Reference to the CA file when 'type' is set to 'Provided'.",
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
																Description:         "Select the type of TLS configuration:<br> - 'Disabled' (default) to not configure TLS for the endpoint. - 'Provided' to manually provide cert file and a key file. - 'Auto' to use OpenShift auto generated certificate using annotations.",
																MarkdownDescription: "Select the type of TLS configuration:<br> - 'Disabled' (default) to not configure TLS for the endpoint. - 'Provided' to manually provide cert file and a key file. - 'Auto' to use OpenShift auto generated certificate using annotations.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("Disabled", "Provided", "Auto"),
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

									"privileged": schema.BoolAttribute{
										Description:         "Privileged mode for the eBPF Agent container. When ignored or set to 'false', the operator sets granular capabilities (BPF, PERFMON, NET_ADMIN, SYS_RESOURCE) to the container. If for some reason these capabilities cannot be set, such as if an old kernel version not knowing CAP_BPF is in use, then you can turn on this mode for more global privileges. Some agent features require the privileged mode, such as packet drops tracking (see 'features') and SR-IOV support.",
										MarkdownDescription: "Privileged mode for the eBPF Agent container. When ignored or set to 'false', the operator sets granular capabilities (BPF, PERFMON, NET_ADMIN, SYS_RESOURCE) to the container. If for some reason these capabilities cannot be set, such as if an old kernel version not knowing CAP_BPF is in use, then you can turn on this mode for more global privileges. Some agent features require the privileged mode, such as packet drops tracking (see 'features') and SR-IOV support.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"resources": schema.SingleNestedAttribute{
										Description:         "'resources' are the compute resources required by this container. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "'resources' are the compute resources required by this container. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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
										Description:         "Sampling rate of the flow reporter. 100 means one flow on 100 is sent. 0 or 1 means all flows are sampled.",
										MarkdownDescription: "Sampling rate of the flow reporter. 100 means one flow on 100 is sent. 0 or 1 means all flows are sampled.",
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
								Description:         "'ipfix' [deprecated (*)] - describes the settings related to the IPFIX-based flow reporter when 'spec.agent.type' is set to 'IPFIX'.",
								MarkdownDescription: "'ipfix' [deprecated (*)] - describes the settings related to the IPFIX-based flow reporter when 'spec.agent.type' is set to 'IPFIX'.",
								Attributes: map[string]schema.Attribute{
									"cache_active_timeout": schema.StringAttribute{
										Description:         "'cacheActiveTimeout' is the max period during which the reporter aggregates flows before sending.",
										MarkdownDescription: "'cacheActiveTimeout' is the max period during which the reporter aggregates flows before sending.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(ns|ms|s|m)?$`), ""),
										},
									},

									"cache_max_flows": schema.Int64Attribute{
										Description:         "'cacheMaxFlows' is the max number of flows in an aggregate; when reached, the reporter sends the flows.",
										MarkdownDescription: "'cacheMaxFlows' is the max number of flows in an aggregate; when reached, the reporter sends the flows.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},

									"cluster_network_operator": schema.SingleNestedAttribute{
										Description:         "'clusterNetworkOperator' defines the settings related to the OpenShift Cluster Network Operator, when available.",
										MarkdownDescription: "'clusterNetworkOperator' defines the settings related to the OpenShift Cluster Network Operator, when available.",
										Attributes: map[string]schema.Attribute{
											"namespace": schema.StringAttribute{
												Description:         "Namespace  where the config map is going to be deployed.",
												MarkdownDescription: "Namespace  where the config map is going to be deployed.",
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
										Description:         "'forceSampleAll' allows disabling sampling in the IPFIX-based flow reporter. It is not recommended to sample all the traffic with IPFIX, as it might generate cluster instability. If you REALLY want to do that, set this flag to 'true'. Use at your own risk. When it is set to 'true', the value of 'sampling' is ignored.",
										MarkdownDescription: "'forceSampleAll' allows disabling sampling in the IPFIX-based flow reporter. It is not recommended to sample all the traffic with IPFIX, as it might generate cluster instability. If you REALLY want to do that, set this flag to 'true'. Use at your own risk. When it is set to 'true', the value of 'sampling' is ignored.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ovn_kubernetes": schema.SingleNestedAttribute{
										Description:         "'ovnKubernetes' defines the settings of the OVN-Kubernetes CNI, when available. This configuration is used when using OVN's IPFIX exports, without OpenShift. When using OpenShift, refer to the 'clusterNetworkOperator' property instead.",
										MarkdownDescription: "'ovnKubernetes' defines the settings of the OVN-Kubernetes CNI, when available. This configuration is used when using OVN's IPFIX exports, without OpenShift. When using OpenShift, refer to the 'clusterNetworkOperator' property instead.",
										Attributes: map[string]schema.Attribute{
											"container_name": schema.StringAttribute{
												Description:         "'containerName' defines the name of the container to configure for IPFIX.",
												MarkdownDescription: "'containerName' defines the name of the container to configure for IPFIX.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"daemon_set_name": schema.StringAttribute{
												Description:         "'daemonSetName' defines the name of the DaemonSet controlling the OVN-Kubernetes pods.",
												MarkdownDescription: "'daemonSetName' defines the name of the DaemonSet controlling the OVN-Kubernetes pods.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "Namespace where OVN-Kubernetes pods are deployed.",
												MarkdownDescription: "Namespace where OVN-Kubernetes pods are deployed.",
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
										Description:         "'sampling' is the sampling rate on the reporter. 100 means one flow on 100 is sent. To ensure cluster stability, it is not possible to set a value below 2. If you really want to sample every packet, which might impact the cluster stability, refer to 'forceSampleAll'. Alternatively, you can use the eBPF Agent instead of IPFIX.",
										MarkdownDescription: "'sampling' is the sampling rate on the reporter. 100 means one flow on 100 is sent. To ensure cluster stability, it is not possible to set a value below 2. If you really want to sample every packet, which might impact the cluster stability, refer to 'forceSampleAll'. Alternatively, you can use the eBPF Agent instead of IPFIX.",
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
								Description:         "'type' [deprecated (*)] selects the flows tracing agent. The only possible value is 'eBPF' (default), to use NetObserv eBPF agent.<br> Previously, using an IPFIX collector was allowed, but was deprecated and it is now removed.<br> Setting 'IPFIX' is ignored and still use the eBPF Agent. Since there is only a single option here, this field will be remove in a future API version.",
								MarkdownDescription: "'type' [deprecated (*)] selects the flows tracing agent. The only possible value is 'eBPF' (default), to use NetObserv eBPF agent.<br> Previously, using an IPFIX collector was allowed, but was deprecated and it is now removed.<br> Setting 'IPFIX' is ignored and still use the eBPF Agent. Since there is only a single option here, this field will be remove in a future API version.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("eBPF", "IPFIX"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"console_plugin": schema.SingleNestedAttribute{
						Description:         "'consolePlugin' defines the settings related to the OpenShift Console plugin, when available.",
						MarkdownDescription: "'consolePlugin' defines the settings related to the OpenShift Console plugin, when available.",
						Attributes: map[string]schema.Attribute{
							"advanced": schema.SingleNestedAttribute{
								Description:         "'advanced' allows setting some aspects of the internal configuration of the console plugin. This section is aimed mostly for debugging and fine-grained performance optimizations, such as 'GOGC' and 'GOMAXPROCS' env vars. Set these values at your own risk.",
								MarkdownDescription: "'advanced' allows setting some aspects of the internal configuration of the console plugin. This section is aimed mostly for debugging and fine-grained performance optimizations, such as 'GOGC' and 'GOMAXPROCS' env vars. Set these values at your own risk.",
								Attributes: map[string]schema.Attribute{
									"args": schema.ListAttribute{
										Description:         "'args' allows passing custom arguments to underlying components. Useful for overriding some parameters, such as an url or a configuration path, that should not be publicly exposed as part of the FlowCollector descriptor, as they are only useful in edge debug or support scenarios.",
										MarkdownDescription: "'args' allows passing custom arguments to underlying components. Useful for overriding some parameters, such as an url or a configuration path, that should not be publicly exposed as part of the FlowCollector descriptor, as they are only useful in edge debug or support scenarios.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"env": schema.MapAttribute{
										Description:         "'env' allows passing custom environment variables to underlying components. Useful for passing some very concrete performance-tuning options, such as 'GOGC' and 'GOMAXPROCS', that should not be publicly exposed as part of the FlowCollector descriptor, as they are only useful in edge debug or support scenarios.",
										MarkdownDescription: "'env' allows passing custom environment variables to underlying components. Useful for passing some very concrete performance-tuning options, such as 'GOGC' and 'GOMAXPROCS', that should not be publicly exposed as part of the FlowCollector descriptor, as they are only useful in edge debug or support scenarios.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"port": schema.Int64Attribute{
										Description:         "'port' is the plugin service port. Do not use 9002, which is reserved for metrics.",
										MarkdownDescription: "'port' is the plugin service port. Do not use 9002, which is reserved for metrics.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
											int64validator.AtMost(65535),
										},
									},

									"register": schema.BoolAttribute{
										Description:         "'register' allows, when set to 'true', to automatically register the provided console plugin with the OpenShift Console operator. When set to 'false', you can still register it manually by editing console.operator.openshift.io/cluster with the following command: 'oc patch console.operator.openshift.io cluster --type='json' -p '[{'op': 'add', 'path': '/spec/plugins/-', 'value': 'netobserv-plugin'}]''",
										MarkdownDescription: "'register' allows, when set to 'true', to automatically register the provided console plugin with the OpenShift Console operator. When set to 'false', you can still register it manually by editing console.operator.openshift.io/cluster with the following command: 'oc patch console.operator.openshift.io cluster --type='json' -p '[{'op': 'add', 'path': '/spec/plugins/-', 'value': 'netobserv-plugin'}]''",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"autoscaler": schema.SingleNestedAttribute{
								Description:         "'autoscaler' spec of a horizontal pod autoscaler to set up for the plugin Deployment.",
								MarkdownDescription: "'autoscaler' spec of a horizontal pod autoscaler to set up for the plugin Deployment.",
								Attributes: map[string]schema.Attribute{
									"max_replicas": schema.Int64Attribute{
										Description:         "'maxReplicas' is the upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas.",
										MarkdownDescription: "'maxReplicas' is the upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"metrics": schema.ListNestedAttribute{
										Description:         "Metrics used by the pod autoscaler",
										MarkdownDescription: "Metrics used by the pod autoscaler",
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
										Description:         "'minReplicas' is the lower limit for the number of replicas to which the autoscaler can scale down. It defaults to 1 pod. minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External metric is configured. Scaling is active as long as at least one metric value is available.",
										MarkdownDescription: "'minReplicas' is the lower limit for the number of replicas to which the autoscaler can scale down. It defaults to 1 pod. minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External metric is configured. Scaling is active as long as at least one metric value is available.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"status": schema.StringAttribute{
										Description:         "'status' describes the desired status regarding deploying an horizontal pod autoscaler.<br> - 'Disabled' does not deploy an horizontal pod autoscaler.<br> - 'Enabled' deploys an horizontal pod autoscaler.<br>",
										MarkdownDescription: "'status' describes the desired status regarding deploying an horizontal pod autoscaler.<br> - 'Disabled' does not deploy an horizontal pod autoscaler.<br> - 'Enabled' deploys an horizontal pod autoscaler.<br>",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("Disabled", "Enabled"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable": schema.BoolAttribute{
								Description:         "Enables the console plugin deployment. 'spec.loki.enable' must also be 'true'",
								MarkdownDescription: "Enables the console plugin deployment. 'spec.loki.enable' must also be 'true'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image_pull_policy": schema.StringAttribute{
								Description:         "'imagePullPolicy' is the Kubernetes pull policy for the image defined above",
								MarkdownDescription: "'imagePullPolicy' is the Kubernetes pull policy for the image defined above",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("IfNotPresent", "Always", "Never"),
								},
							},

							"log_level": schema.StringAttribute{
								Description:         "'logLevel' for the console plugin backend",
								MarkdownDescription: "'logLevel' for the console plugin backend",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("trace", "debug", "info", "warn", "error", "fatal", "panic"),
								},
							},

							"port_naming": schema.SingleNestedAttribute{
								Description:         "'portNaming' defines the configuration of the port-to-service name translation",
								MarkdownDescription: "'portNaming' defines the configuration of the port-to-service name translation",
								Attributes: map[string]schema.Attribute{
									"enable": schema.BoolAttribute{
										Description:         "Enable the console plugin port-to-service name translation",
										MarkdownDescription: "Enable the console plugin port-to-service name translation",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"port_names": schema.MapAttribute{
										Description:         "'portNames' defines additional port names to use in the console, for example, 'portNames: {'3100': 'loki'}'.",
										MarkdownDescription: "'portNames' defines additional port names to use in the console, for example, 'portNames: {'3100': 'loki'}'.",
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
								Description:         "'quickFilters' configures quick filter presets for the Console plugin",
								MarkdownDescription: "'quickFilters' configures quick filter presets for the Console plugin",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"default": schema.BoolAttribute{
											Description:         "'default' defines whether this filter should be active by default or not",
											MarkdownDescription: "'default' defines whether this filter should be active by default or not",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"filter": schema.MapAttribute{
											Description:         "'filter' is a set of keys and values to be set when this filter is selected. Each key can relate to a list of values using a coma-separated string, for example, 'filter: {'src_namespace': 'namespace1,namespace2'}'.",
											MarkdownDescription: "'filter' is a set of keys and values to be set when this filter is selected. Each key can relate to a list of values using a coma-separated string, for example, 'filter: {'src_namespace': 'namespace1,namespace2'}'.",
											ElementType:         types.StringType,
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name of the filter, that is displayed in the Console",
											MarkdownDescription: "Name of the filter, that is displayed in the Console",
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

							"replicas": schema.Int64Attribute{
								Description:         "'replicas' defines the number of replicas (pods) to start.",
								MarkdownDescription: "'replicas' defines the number of replicas (pods) to start.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
								},
							},

							"resources": schema.SingleNestedAttribute{
								Description:         "'resources', in terms of compute resources, required by this container. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "'resources', in terms of compute resources, required by this container. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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
						Description:         "'deploymentModel' defines the desired type of deployment for flow processing. Possible values are:<br> - 'Direct' (default) to make the flow processor listening directly from the agents.<br> - 'Kafka' to make flows sent to a Kafka pipeline before consumption by the processor.<br> Kafka can provide better scalability, resiliency, and high availability (for more details, see https://www.redhat.com/en/topics/integration/what-is-apache-kafka).",
						MarkdownDescription: "'deploymentModel' defines the desired type of deployment for flow processing. Possible values are:<br> - 'Direct' (default) to make the flow processor listening directly from the agents.<br> - 'Kafka' to make flows sent to a Kafka pipeline before consumption by the processor.<br> Kafka can provide better scalability, resiliency, and high availability (for more details, see https://www.redhat.com/en/topics/integration/what-is-apache-kafka).",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Direct", "Kafka"),
						},
					},

					"exporters": schema.ListNestedAttribute{
						Description:         "'exporters' define additional optional exporters for custom consumption or storage.",
						MarkdownDescription: "'exporters' define additional optional exporters for custom consumption or storage.",
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
									Description:         "Kafka configuration, such as the address and topic, to send enriched flows to.",
									MarkdownDescription: "Kafka configuration, such as the address and topic, to send enriched flows to.",
									Attributes: map[string]schema.Attribute{
										"address": schema.StringAttribute{
											Description:         "Address of the Kafka server",
											MarkdownDescription: "Address of the Kafka server",
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
													Description:         "Type of SASL authentication to use, or 'Disabled' if SASL is not used",
													MarkdownDescription: "Type of SASL authentication to use, or 'Disabled' if SASL is not used",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("Disabled", "Plain", "ScramSHA512"),
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"tls": schema.SingleNestedAttribute{
											Description:         "TLS client configuration. When using TLS, verify that the address matches the Kafka port used for TLS, generally 9093.",
											MarkdownDescription: "TLS client configuration. When using TLS, verify that the address matches the Kafka port used for TLS, generally 9093.",
											Attributes: map[string]schema.Attribute{
												"ca_cert": schema.SingleNestedAttribute{
													Description:         "'caCert' defines the reference of the certificate for the Certificate Authority",
													MarkdownDescription: "'caCert' defines the reference of the certificate for the Certificate Authority",
													Attributes: map[string]schema.Attribute{
														"cert_file": schema.StringAttribute{
															Description:         "'certFile' defines the path to the certificate file name within the config map or secret",
															MarkdownDescription: "'certFile' defines the path to the certificate file name within the config map or secret",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"cert_key": schema.StringAttribute{
															Description:         "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
															MarkdownDescription: "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the config map or secret containing certificates",
															MarkdownDescription: "Name of the config map or secret containing certificates",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespace": schema.StringAttribute{
															Description:         "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
															MarkdownDescription: "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"type": schema.StringAttribute{
															Description:         "Type for the certificate reference: 'configmap' or 'secret'",
															MarkdownDescription: "Type for the certificate reference: 'configmap' or 'secret'",
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
													Description:         "Enable TLS",
													MarkdownDescription: "Enable TLS",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"insecure_skip_verify": schema.BoolAttribute{
													Description:         "'insecureSkipVerify' allows skipping client-side verification of the server certificate. If set to 'true', the 'caCert' field is ignored.",
													MarkdownDescription: "'insecureSkipVerify' allows skipping client-side verification of the server certificate. If set to 'true', the 'caCert' field is ignored.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"user_cert": schema.SingleNestedAttribute{
													Description:         "'userCert' defines the user certificate reference and is used for mTLS (you can ignore it when using one-way TLS)",
													MarkdownDescription: "'userCert' defines the user certificate reference and is used for mTLS (you can ignore it when using one-way TLS)",
													Attributes: map[string]schema.Attribute{
														"cert_file": schema.StringAttribute{
															Description:         "'certFile' defines the path to the certificate file name within the config map or secret",
															MarkdownDescription: "'certFile' defines the path to the certificate file name within the config map or secret",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"cert_key": schema.StringAttribute{
															Description:         "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
															MarkdownDescription: "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the config map or secret containing certificates",
															MarkdownDescription: "Name of the config map or secret containing certificates",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespace": schema.StringAttribute{
															Description:         "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
															MarkdownDescription: "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"type": schema.StringAttribute{
															Description:         "Type for the certificate reference: 'configmap' or 'secret'",
															MarkdownDescription: "Type for the certificate reference: 'configmap' or 'secret'",
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
											Description:         "Kafka topic to use. It must exist. NetObserv does not create it.",
											MarkdownDescription: "Kafka topic to use. It must exist. NetObserv does not create it.",
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
									Description:         "'type' selects the type of exporters. The available options are 'Kafka' and 'IPFIX'.",
									MarkdownDescription: "'type' selects the type of exporters. The available options are 'Kafka' and 'IPFIX'.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("Kafka", "IPFIX"),
									},
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"kafka": schema.SingleNestedAttribute{
						Description:         "Kafka configuration, allowing to use Kafka as a broker as part of the flow collection pipeline. Available when the 'spec.deploymentModel' is 'Kafka'.",
						MarkdownDescription: "Kafka configuration, allowing to use Kafka as a broker as part of the flow collection pipeline. Available when the 'spec.deploymentModel' is 'Kafka'.",
						Attributes: map[string]schema.Attribute{
							"address": schema.StringAttribute{
								Description:         "Address of the Kafka server",
								MarkdownDescription: "Address of the Kafka server",
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
										Description:         "Type of SASL authentication to use, or 'Disabled' if SASL is not used",
										MarkdownDescription: "Type of SASL authentication to use, or 'Disabled' if SASL is not used",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("Disabled", "Plain", "ScramSHA512"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"tls": schema.SingleNestedAttribute{
								Description:         "TLS client configuration. When using TLS, verify that the address matches the Kafka port used for TLS, generally 9093.",
								MarkdownDescription: "TLS client configuration. When using TLS, verify that the address matches the Kafka port used for TLS, generally 9093.",
								Attributes: map[string]schema.Attribute{
									"ca_cert": schema.SingleNestedAttribute{
										Description:         "'caCert' defines the reference of the certificate for the Certificate Authority",
										MarkdownDescription: "'caCert' defines the reference of the certificate for the Certificate Authority",
										Attributes: map[string]schema.Attribute{
											"cert_file": schema.StringAttribute{
												Description:         "'certFile' defines the path to the certificate file name within the config map or secret",
												MarkdownDescription: "'certFile' defines the path to the certificate file name within the config map or secret",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"cert_key": schema.StringAttribute{
												Description:         "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
												MarkdownDescription: "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the config map or secret containing certificates",
												MarkdownDescription: "Name of the config map or secret containing certificates",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
												MarkdownDescription: "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
												Description:         "Type for the certificate reference: 'configmap' or 'secret'",
												MarkdownDescription: "Type for the certificate reference: 'configmap' or 'secret'",
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
										Description:         "Enable TLS",
										MarkdownDescription: "Enable TLS",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"insecure_skip_verify": schema.BoolAttribute{
										Description:         "'insecureSkipVerify' allows skipping client-side verification of the server certificate. If set to 'true', the 'caCert' field is ignored.",
										MarkdownDescription: "'insecureSkipVerify' allows skipping client-side verification of the server certificate. If set to 'true', the 'caCert' field is ignored.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"user_cert": schema.SingleNestedAttribute{
										Description:         "'userCert' defines the user certificate reference and is used for mTLS (you can ignore it when using one-way TLS)",
										MarkdownDescription: "'userCert' defines the user certificate reference and is used for mTLS (you can ignore it when using one-way TLS)",
										Attributes: map[string]schema.Attribute{
											"cert_file": schema.StringAttribute{
												Description:         "'certFile' defines the path to the certificate file name within the config map or secret",
												MarkdownDescription: "'certFile' defines the path to the certificate file name within the config map or secret",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"cert_key": schema.StringAttribute{
												Description:         "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
												MarkdownDescription: "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the config map or secret containing certificates",
												MarkdownDescription: "Name of the config map or secret containing certificates",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
												MarkdownDescription: "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
												Description:         "Type for the certificate reference: 'configmap' or 'secret'",
												MarkdownDescription: "Type for the certificate reference: 'configmap' or 'secret'",
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
								Description:         "Kafka topic to use. It must exist. NetObserv does not create it.",
								MarkdownDescription: "Kafka topic to use. It must exist. NetObserv does not create it.",
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
						Description:         "'loki', the flow store, client settings.",
						MarkdownDescription: "'loki', the flow store, client settings.",
						Attributes: map[string]schema.Attribute{
							"advanced": schema.SingleNestedAttribute{
								Description:         "'advanced' allows setting some aspects of the internal configuration of the Loki clients. This section is aimed mostly for debugging and fine-grained performance optimizations.",
								MarkdownDescription: "'advanced' allows setting some aspects of the internal configuration of the Loki clients. This section is aimed mostly for debugging and fine-grained performance optimizations.",
								Attributes: map[string]schema.Attribute{
									"static_labels": schema.MapAttribute{
										Description:         "'staticLabels' is a map of common labels to set on each flow in Loki storage.",
										MarkdownDescription: "'staticLabels' is a map of common labels to set on each flow in Loki storage.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"write_max_backoff": schema.StringAttribute{
										Description:         "'writeMaxBackoff' is the maximum backoff time for Loki client connection between retries.",
										MarkdownDescription: "'writeMaxBackoff' is the maximum backoff time for Loki client connection between retries.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"write_max_retries": schema.Int64Attribute{
										Description:         "'writeMaxRetries' is the maximum number of retries for Loki client connections.",
										MarkdownDescription: "'writeMaxRetries' is the maximum number of retries for Loki client connections.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},

									"write_min_backoff": schema.StringAttribute{
										Description:         "'writeMinBackoff' is the initial backoff time for Loki client connection between retries.",
										MarkdownDescription: "'writeMinBackoff' is the initial backoff time for Loki client connection between retries.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable": schema.BoolAttribute{
								Description:         "Set 'enable' to 'true' to store flows in Loki. It is required for the OpenShift Console plugin installation.",
								MarkdownDescription: "Set 'enable' to 'true' to store flows in Loki. It is required for the OpenShift Console plugin installation.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"loki_stack": schema.SingleNestedAttribute{
								Description:         "Loki configuration for 'LokiStack' mode. This is useful for an easy loki-operator configuration. It is ignored for other modes.",
								MarkdownDescription: "Loki configuration for 'LokiStack' mode. This is useful for an easy loki-operator configuration. It is ignored for other modes.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of an existing LokiStack resource to use.",
										MarkdownDescription: "Name of an existing LokiStack resource to use.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
										Description:         "Namespace where this 'LokiStack' resource is located. If omited, it is assumed to be the same as 'spec.namespace'.",
										MarkdownDescription: "Namespace where this 'LokiStack' resource is located. If omited, it is assumed to be the same as 'spec.namespace'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"manual": schema.SingleNestedAttribute{
								Description:         "Loki configuration for 'Manual' mode. This is the most flexible configuration. It is ignored for other modes.",
								MarkdownDescription: "Loki configuration for 'Manual' mode. This is the most flexible configuration. It is ignored for other modes.",
								Attributes: map[string]schema.Attribute{
									"auth_token": schema.StringAttribute{
										Description:         "'authToken' describes the way to get a token to authenticate to Loki.<br> - 'Disabled' does not send any token with the request.<br> - 'Forward' forwards the user token for authorization.<br> - 'Host' [deprecated (*)] - uses the local pod service account to authenticate to Loki.<br> When using the Loki Operator, this must be set to 'Forward'.",
										MarkdownDescription: "'authToken' describes the way to get a token to authenticate to Loki.<br> - 'Disabled' does not send any token with the request.<br> - 'Forward' forwards the user token for authorization.<br> - 'Host' [deprecated (*)] - uses the local pod service account to authenticate to Loki.<br> When using the Loki Operator, this must be set to 'Forward'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("Disabled", "Host", "Forward"),
										},
									},

									"ingester_url": schema.StringAttribute{
										Description:         "'ingesterUrl' is the address of an existing Loki ingester service to push the flows to. When using the Loki Operator, set it to the Loki gateway service with the 'network' tenant set in path, for example https://loki-gateway-http.netobserv.svc:8080/api/logs/v1/network.",
										MarkdownDescription: "'ingesterUrl' is the address of an existing Loki ingester service to push the flows to. When using the Loki Operator, set it to the Loki gateway service with the 'network' tenant set in path, for example https://loki-gateway-http.netobserv.svc:8080/api/logs/v1/network.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"querier_url": schema.StringAttribute{
										Description:         "'querierUrl' specifies the address of the Loki querier service. When using the Loki Operator, set it to the Loki gateway service with the 'network' tenant set in path, for example https://loki-gateway-http.netobserv.svc:8080/api/logs/v1/network.",
										MarkdownDescription: "'querierUrl' specifies the address of the Loki querier service. When using the Loki Operator, set it to the Loki gateway service with the 'network' tenant set in path, for example https://loki-gateway-http.netobserv.svc:8080/api/logs/v1/network.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"status_tls": schema.SingleNestedAttribute{
										Description:         "TLS client configuration for Loki status URL.",
										MarkdownDescription: "TLS client configuration for Loki status URL.",
										Attributes: map[string]schema.Attribute{
											"ca_cert": schema.SingleNestedAttribute{
												Description:         "'caCert' defines the reference of the certificate for the Certificate Authority",
												MarkdownDescription: "'caCert' defines the reference of the certificate for the Certificate Authority",
												Attributes: map[string]schema.Attribute{
													"cert_file": schema.StringAttribute{
														Description:         "'certFile' defines the path to the certificate file name within the config map or secret",
														MarkdownDescription: "'certFile' defines the path to the certificate file name within the config map or secret",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"cert_key": schema.StringAttribute{
														Description:         "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
														MarkdownDescription: "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the config map or secret containing certificates",
														MarkdownDescription: "Name of the config map or secret containing certificates",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
														MarkdownDescription: "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "Type for the certificate reference: 'configmap' or 'secret'",
														MarkdownDescription: "Type for the certificate reference: 'configmap' or 'secret'",
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
												Description:         "Enable TLS",
												MarkdownDescription: "Enable TLS",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"insecure_skip_verify": schema.BoolAttribute{
												Description:         "'insecureSkipVerify' allows skipping client-side verification of the server certificate. If set to 'true', the 'caCert' field is ignored.",
												MarkdownDescription: "'insecureSkipVerify' allows skipping client-side verification of the server certificate. If set to 'true', the 'caCert' field is ignored.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"user_cert": schema.SingleNestedAttribute{
												Description:         "'userCert' defines the user certificate reference and is used for mTLS (you can ignore it when using one-way TLS)",
												MarkdownDescription: "'userCert' defines the user certificate reference and is used for mTLS (you can ignore it when using one-way TLS)",
												Attributes: map[string]schema.Attribute{
													"cert_file": schema.StringAttribute{
														Description:         "'certFile' defines the path to the certificate file name within the config map or secret",
														MarkdownDescription: "'certFile' defines the path to the certificate file name within the config map or secret",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"cert_key": schema.StringAttribute{
														Description:         "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
														MarkdownDescription: "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the config map or secret containing certificates",
														MarkdownDescription: "Name of the config map or secret containing certificates",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
														MarkdownDescription: "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "Type for the certificate reference: 'configmap' or 'secret'",
														MarkdownDescription: "Type for the certificate reference: 'configmap' or 'secret'",
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

									"status_url": schema.StringAttribute{
										Description:         "'statusUrl' specifies the address of the Loki '/ready', '/metrics' and '/config' endpoints, in case it is different from the Loki querier URL. If empty, the 'querierUrl' value is used. This is useful to show error messages and some context in the frontend. When using the Loki Operator, set it to the Loki HTTP query frontend service, for example https://loki-query-frontend-http.netobserv.svc:3100/. 'statusTLS' configuration is used when 'statusUrl' is set.",
										MarkdownDescription: "'statusUrl' specifies the address of the Loki '/ready', '/metrics' and '/config' endpoints, in case it is different from the Loki querier URL. If empty, the 'querierUrl' value is used. This is useful to show error messages and some context in the frontend. When using the Loki Operator, set it to the Loki HTTP query frontend service, for example https://loki-query-frontend-http.netobserv.svc:3100/. 'statusTLS' configuration is used when 'statusUrl' is set.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tenant_id": schema.StringAttribute{
										Description:         "'tenantID' is the Loki 'X-Scope-OrgID' that identifies the tenant for each request. When using the Loki Operator, set it to 'network', which corresponds to a special tenant mode.",
										MarkdownDescription: "'tenantID' is the Loki 'X-Scope-OrgID' that identifies the tenant for each request. When using the Loki Operator, set it to 'network', which corresponds to a special tenant mode.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tls": schema.SingleNestedAttribute{
										Description:         "TLS client configuration for Loki URL.",
										MarkdownDescription: "TLS client configuration for Loki URL.",
										Attributes: map[string]schema.Attribute{
											"ca_cert": schema.SingleNestedAttribute{
												Description:         "'caCert' defines the reference of the certificate for the Certificate Authority",
												MarkdownDescription: "'caCert' defines the reference of the certificate for the Certificate Authority",
												Attributes: map[string]schema.Attribute{
													"cert_file": schema.StringAttribute{
														Description:         "'certFile' defines the path to the certificate file name within the config map or secret",
														MarkdownDescription: "'certFile' defines the path to the certificate file name within the config map or secret",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"cert_key": schema.StringAttribute{
														Description:         "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
														MarkdownDescription: "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the config map or secret containing certificates",
														MarkdownDescription: "Name of the config map or secret containing certificates",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
														MarkdownDescription: "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "Type for the certificate reference: 'configmap' or 'secret'",
														MarkdownDescription: "Type for the certificate reference: 'configmap' or 'secret'",
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
												Description:         "Enable TLS",
												MarkdownDescription: "Enable TLS",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"insecure_skip_verify": schema.BoolAttribute{
												Description:         "'insecureSkipVerify' allows skipping client-side verification of the server certificate. If set to 'true', the 'caCert' field is ignored.",
												MarkdownDescription: "'insecureSkipVerify' allows skipping client-side verification of the server certificate. If set to 'true', the 'caCert' field is ignored.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"user_cert": schema.SingleNestedAttribute{
												Description:         "'userCert' defines the user certificate reference and is used for mTLS (you can ignore it when using one-way TLS)",
												MarkdownDescription: "'userCert' defines the user certificate reference and is used for mTLS (you can ignore it when using one-way TLS)",
												Attributes: map[string]schema.Attribute{
													"cert_file": schema.StringAttribute{
														Description:         "'certFile' defines the path to the certificate file name within the config map or secret",
														MarkdownDescription: "'certFile' defines the path to the certificate file name within the config map or secret",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"cert_key": schema.StringAttribute{
														Description:         "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
														MarkdownDescription: "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the config map or secret containing certificates",
														MarkdownDescription: "Name of the config map or secret containing certificates",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
														MarkdownDescription: "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "Type for the certificate reference: 'configmap' or 'secret'",
														MarkdownDescription: "Type for the certificate reference: 'configmap' or 'secret'",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"microservices": schema.SingleNestedAttribute{
								Description:         "Loki configuration for 'Microservices' mode. Use this option when Loki is installed using the microservices deployment mode (https://grafana.com/docs/loki/latest/fundamentals/architecture/deployment-modes/#microservices-mode). It is ignored for other modes.",
								MarkdownDescription: "Loki configuration for 'Microservices' mode. Use this option when Loki is installed using the microservices deployment mode (https://grafana.com/docs/loki/latest/fundamentals/architecture/deployment-modes/#microservices-mode). It is ignored for other modes.",
								Attributes: map[string]schema.Attribute{
									"ingester_url": schema.StringAttribute{
										Description:         "'ingesterUrl' is the address of an existing Loki ingester service to push the flows to.",
										MarkdownDescription: "'ingesterUrl' is the address of an existing Loki ingester service to push the flows to.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"querier_url": schema.StringAttribute{
										Description:         "'querierURL' specifies the address of the Loki querier service.",
										MarkdownDescription: "'querierURL' specifies the address of the Loki querier service.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tenant_id": schema.StringAttribute{
										Description:         "'tenantID' is the Loki 'X-Scope-OrgID' header that identifies the tenant for each request.",
										MarkdownDescription: "'tenantID' is the Loki 'X-Scope-OrgID' header that identifies the tenant for each request.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tls": schema.SingleNestedAttribute{
										Description:         "TLS client configuration for Loki URL.",
										MarkdownDescription: "TLS client configuration for Loki URL.",
										Attributes: map[string]schema.Attribute{
											"ca_cert": schema.SingleNestedAttribute{
												Description:         "'caCert' defines the reference of the certificate for the Certificate Authority",
												MarkdownDescription: "'caCert' defines the reference of the certificate for the Certificate Authority",
												Attributes: map[string]schema.Attribute{
													"cert_file": schema.StringAttribute{
														Description:         "'certFile' defines the path to the certificate file name within the config map or secret",
														MarkdownDescription: "'certFile' defines the path to the certificate file name within the config map or secret",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"cert_key": schema.StringAttribute{
														Description:         "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
														MarkdownDescription: "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the config map or secret containing certificates",
														MarkdownDescription: "Name of the config map or secret containing certificates",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
														MarkdownDescription: "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "Type for the certificate reference: 'configmap' or 'secret'",
														MarkdownDescription: "Type for the certificate reference: 'configmap' or 'secret'",
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
												Description:         "Enable TLS",
												MarkdownDescription: "Enable TLS",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"insecure_skip_verify": schema.BoolAttribute{
												Description:         "'insecureSkipVerify' allows skipping client-side verification of the server certificate. If set to 'true', the 'caCert' field is ignored.",
												MarkdownDescription: "'insecureSkipVerify' allows skipping client-side verification of the server certificate. If set to 'true', the 'caCert' field is ignored.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"user_cert": schema.SingleNestedAttribute{
												Description:         "'userCert' defines the user certificate reference and is used for mTLS (you can ignore it when using one-way TLS)",
												MarkdownDescription: "'userCert' defines the user certificate reference and is used for mTLS (you can ignore it when using one-way TLS)",
												Attributes: map[string]schema.Attribute{
													"cert_file": schema.StringAttribute{
														Description:         "'certFile' defines the path to the certificate file name within the config map or secret",
														MarkdownDescription: "'certFile' defines the path to the certificate file name within the config map or secret",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"cert_key": schema.StringAttribute{
														Description:         "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
														MarkdownDescription: "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the config map or secret containing certificates",
														MarkdownDescription: "Name of the config map or secret containing certificates",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
														MarkdownDescription: "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "Type for the certificate reference: 'configmap' or 'secret'",
														MarkdownDescription: "Type for the certificate reference: 'configmap' or 'secret'",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"mode": schema.StringAttribute{
								Description:         "'mode' must be set according to the installation mode of Loki:<br> - Use 'LokiStack' when Loki is managed using the Loki Operator<br> - Use 'Monolithic' when Loki is installed as a monolithic workload<br> - Use 'Microservices' when Loki is installed as microservices, but without Loki Operator<br> - Use 'Manual' if none of the options above match your setup<br>",
								MarkdownDescription: "'mode' must be set according to the installation mode of Loki:<br> - Use 'LokiStack' when Loki is managed using the Loki Operator<br> - Use 'Monolithic' when Loki is installed as a monolithic workload<br> - Use 'Microservices' when Loki is installed as microservices, but without Loki Operator<br> - Use 'Manual' if none of the options above match your setup<br>",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Manual", "LokiStack", "Monolithic", "Microservices"),
								},
							},

							"monolithic": schema.SingleNestedAttribute{
								Description:         "Loki configuration for 'Monolithic' mode. Use this option when Loki is installed using the monolithic deployment mode (https://grafana.com/docs/loki/latest/fundamentals/architecture/deployment-modes/#monolithic-mode). It is ignored for other modes.",
								MarkdownDescription: "Loki configuration for 'Monolithic' mode. Use this option when Loki is installed using the monolithic deployment mode (https://grafana.com/docs/loki/latest/fundamentals/architecture/deployment-modes/#monolithic-mode). It is ignored for other modes.",
								Attributes: map[string]schema.Attribute{
									"tenant_id": schema.StringAttribute{
										Description:         "'tenantID' is the Loki 'X-Scope-OrgID' header that identifies the tenant for each request.",
										MarkdownDescription: "'tenantID' is the Loki 'X-Scope-OrgID' header that identifies the tenant for each request.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tls": schema.SingleNestedAttribute{
										Description:         "TLS client configuration for Loki URL.",
										MarkdownDescription: "TLS client configuration for Loki URL.",
										Attributes: map[string]schema.Attribute{
											"ca_cert": schema.SingleNestedAttribute{
												Description:         "'caCert' defines the reference of the certificate for the Certificate Authority",
												MarkdownDescription: "'caCert' defines the reference of the certificate for the Certificate Authority",
												Attributes: map[string]schema.Attribute{
													"cert_file": schema.StringAttribute{
														Description:         "'certFile' defines the path to the certificate file name within the config map or secret",
														MarkdownDescription: "'certFile' defines the path to the certificate file name within the config map or secret",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"cert_key": schema.StringAttribute{
														Description:         "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
														MarkdownDescription: "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the config map or secret containing certificates",
														MarkdownDescription: "Name of the config map or secret containing certificates",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
														MarkdownDescription: "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "Type for the certificate reference: 'configmap' or 'secret'",
														MarkdownDescription: "Type for the certificate reference: 'configmap' or 'secret'",
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
												Description:         "Enable TLS",
												MarkdownDescription: "Enable TLS",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"insecure_skip_verify": schema.BoolAttribute{
												Description:         "'insecureSkipVerify' allows skipping client-side verification of the server certificate. If set to 'true', the 'caCert' field is ignored.",
												MarkdownDescription: "'insecureSkipVerify' allows skipping client-side verification of the server certificate. If set to 'true', the 'caCert' field is ignored.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"user_cert": schema.SingleNestedAttribute{
												Description:         "'userCert' defines the user certificate reference and is used for mTLS (you can ignore it when using one-way TLS)",
												MarkdownDescription: "'userCert' defines the user certificate reference and is used for mTLS (you can ignore it when using one-way TLS)",
												Attributes: map[string]schema.Attribute{
													"cert_file": schema.StringAttribute{
														Description:         "'certFile' defines the path to the certificate file name within the config map or secret",
														MarkdownDescription: "'certFile' defines the path to the certificate file name within the config map or secret",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"cert_key": schema.StringAttribute{
														Description:         "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
														MarkdownDescription: "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the config map or secret containing certificates",
														MarkdownDescription: "Name of the config map or secret containing certificates",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
														MarkdownDescription: "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "Type for the certificate reference: 'configmap' or 'secret'",
														MarkdownDescription: "Type for the certificate reference: 'configmap' or 'secret'",
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
										Description:         "'url' is the unique address of an existing Loki service that points to both the ingester and the querier.",
										MarkdownDescription: "'url' is the unique address of an existing Loki service that points to both the ingester and the querier.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"read_timeout": schema.StringAttribute{
								Description:         "'readTimeout' is the maximum console plugin loki query total time limit. A timeout of zero means no timeout.",
								MarkdownDescription: "'readTimeout' is the maximum console plugin loki query total time limit. A timeout of zero means no timeout.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"write_batch_size": schema.Int64Attribute{
								Description:         "'writeBatchSize' is the maximum batch size (in bytes) of Loki logs to accumulate before sending.",
								MarkdownDescription: "'writeBatchSize' is the maximum batch size (in bytes) of Loki logs to accumulate before sending.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
								},
							},

							"write_batch_wait": schema.StringAttribute{
								Description:         "'writeBatchWait' is the maximum time to wait before sending a Loki batch.",
								MarkdownDescription: "'writeBatchWait' is the maximum time to wait before sending a Loki batch.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"write_timeout": schema.StringAttribute{
								Description:         "'writeTimeout' is the maximum Loki time connection / request limit. A timeout of zero means no timeout.",
								MarkdownDescription: "'writeTimeout' is the maximum Loki time connection / request limit. A timeout of zero means no timeout.",
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
						Description:         "Namespace where NetObserv pods are deployed.",
						MarkdownDescription: "Namespace where NetObserv pods are deployed.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"processor": schema.SingleNestedAttribute{
						Description:         "'processor' defines the settings of the component that receives the flows from the agent, enriches them, generates metrics, and forwards them to the Loki persistence layer and/or any available exporter.",
						MarkdownDescription: "'processor' defines the settings of the component that receives the flows from the agent, enriches them, generates metrics, and forwards them to the Loki persistence layer and/or any available exporter.",
						Attributes: map[string]schema.Attribute{
							"add_zone": schema.BoolAttribute{
								Description:         "'addZone' allows availability zone awareness by labelling flows with their source and destination zones. This feature requires the 'topology.kubernetes.io/zone' label to be set on nodes.",
								MarkdownDescription: "'addZone' allows availability zone awareness by labelling flows with their source and destination zones. This feature requires the 'topology.kubernetes.io/zone' label to be set on nodes.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"advanced": schema.SingleNestedAttribute{
								Description:         "'advanced' allows setting some aspects of the internal configuration of the flow processor. This section is aimed mostly for debugging and fine-grained performance optimizations, such as 'GOGC' and 'GOMAXPROCS' env vars. Set these values at your own risk.",
								MarkdownDescription: "'advanced' allows setting some aspects of the internal configuration of the flow processor. This section is aimed mostly for debugging and fine-grained performance optimizations, such as 'GOGC' and 'GOMAXPROCS' env vars. Set these values at your own risk.",
								Attributes: map[string]schema.Attribute{
									"conversation_end_timeout": schema.StringAttribute{
										Description:         "'conversationEndTimeout' is the time to wait after a network flow is received, to consider the conversation ended. This delay is ignored when a FIN packet is collected for TCP flows (see 'conversationTerminatingTimeout' instead).",
										MarkdownDescription: "'conversationEndTimeout' is the time to wait after a network flow is received, to consider the conversation ended. This delay is ignored when a FIN packet is collected for TCP flows (see 'conversationTerminatingTimeout' instead).",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"conversation_heartbeat_interval": schema.StringAttribute{
										Description:         "'conversationHeartbeatInterval' is the time to wait between 'tick' events of a conversation",
										MarkdownDescription: "'conversationHeartbeatInterval' is the time to wait between 'tick' events of a conversation",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"conversation_terminating_timeout": schema.StringAttribute{
										Description:         "'conversationTerminatingTimeout' is the time to wait from detected FIN flag to end a conversation. Only relevant for TCP flows.",
										MarkdownDescription: "'conversationTerminatingTimeout' is the time to wait from detected FIN flag to end a conversation. Only relevant for TCP flows.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"drop_unused_fields": schema.BoolAttribute{
										Description:         "'dropUnusedFields' [deprecated (*)] this setting is not used anymore.",
										MarkdownDescription: "'dropUnusedFields' [deprecated (*)] this setting is not used anymore.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"enable_kube_probes": schema.BoolAttribute{
										Description:         "'enableKubeProbes' is a flag to enable or disable Kubernetes liveness and readiness probes",
										MarkdownDescription: "'enableKubeProbes' is a flag to enable or disable Kubernetes liveness and readiness probes",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"env": schema.MapAttribute{
										Description:         "'env' allows passing custom environment variables to underlying components. Useful for passing some very concrete performance-tuning options, such as 'GOGC' and 'GOMAXPROCS', that should not be publicly exposed as part of the FlowCollector descriptor, as they are only useful in edge debug or support scenarios.",
										MarkdownDescription: "'env' allows passing custom environment variables to underlying components. Useful for passing some very concrete performance-tuning options, such as 'GOGC' and 'GOMAXPROCS', that should not be publicly exposed as part of the FlowCollector descriptor, as they are only useful in edge debug or support scenarios.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"health_port": schema.Int64Attribute{
										Description:         "'healthPort' is a collector HTTP port in the Pod that exposes the health check API",
										MarkdownDescription: "'healthPort' is a collector HTTP port in the Pod that exposes the health check API",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
											int64validator.AtMost(65535),
										},
									},

									"port": schema.Int64Attribute{
										Description:         "Port of the flow collector (host port). By convention, some values are forbidden. It must be greater than 1024 and different from 4500, 4789 and 6081.",
										MarkdownDescription: "Port of the flow collector (host port). By convention, some values are forbidden. It must be greater than 1024 and different from 4500, 4789 and 6081.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1025),
											int64validator.AtMost(65535),
										},
									},

									"profile_port": schema.Int64Attribute{
										Description:         "'profilePort' allows setting up a Go pprof profiler listening to this port",
										MarkdownDescription: "'profilePort' allows setting up a Go pprof profiler listening to this port",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
											int64validator.AtMost(65535),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"cluster_name": schema.StringAttribute{
								Description:         "'clusterName' is the name of the cluster to appear in the flows data. This is useful in a multi-cluster context. When using OpenShift, leave empty to make it automatically determined.",
								MarkdownDescription: "'clusterName' is the name of the cluster to appear in the flows data. This is useful in a multi-cluster context. When using OpenShift, leave empty to make it automatically determined.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image_pull_policy": schema.StringAttribute{
								Description:         "'imagePullPolicy' is the Kubernetes pull policy for the image defined above",
								MarkdownDescription: "'imagePullPolicy' is the Kubernetes pull policy for the image defined above",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("IfNotPresent", "Always", "Never"),
								},
							},

							"kafka_consumer_autoscaler": schema.SingleNestedAttribute{
								Description:         "'kafkaConsumerAutoscaler' is the spec of a horizontal pod autoscaler to set up for 'flowlogs-pipeline-transformer', which consumes Kafka messages. This setting is ignored when Kafka is disabled.",
								MarkdownDescription: "'kafkaConsumerAutoscaler' is the spec of a horizontal pod autoscaler to set up for 'flowlogs-pipeline-transformer', which consumes Kafka messages. This setting is ignored when Kafka is disabled.",
								Attributes: map[string]schema.Attribute{
									"max_replicas": schema.Int64Attribute{
										Description:         "'maxReplicas' is the upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas.",
										MarkdownDescription: "'maxReplicas' is the upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"metrics": schema.ListNestedAttribute{
										Description:         "Metrics used by the pod autoscaler",
										MarkdownDescription: "Metrics used by the pod autoscaler",
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
										Description:         "'minReplicas' is the lower limit for the number of replicas to which the autoscaler can scale down. It defaults to 1 pod. minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External metric is configured. Scaling is active as long as at least one metric value is available.",
										MarkdownDescription: "'minReplicas' is the lower limit for the number of replicas to which the autoscaler can scale down. It defaults to 1 pod. minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External metric is configured. Scaling is active as long as at least one metric value is available.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"status": schema.StringAttribute{
										Description:         "'status' describes the desired status regarding deploying an horizontal pod autoscaler.<br> - 'Disabled' does not deploy an horizontal pod autoscaler.<br> - 'Enabled' deploys an horizontal pod autoscaler.<br>",
										MarkdownDescription: "'status' describes the desired status regarding deploying an horizontal pod autoscaler.<br> - 'Disabled' does not deploy an horizontal pod autoscaler.<br> - 'Enabled' deploys an horizontal pod autoscaler.<br>",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("Disabled", "Enabled"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"kafka_consumer_batch_size": schema.Int64Attribute{
								Description:         "'kafkaConsumerBatchSize' indicates to the broker the maximum batch size, in bytes, that the consumer accepts. Ignored when not using Kafka. Default: 10MB.",
								MarkdownDescription: "'kafkaConsumerBatchSize' indicates to the broker the maximum batch size, in bytes, that the consumer accepts. Ignored when not using Kafka. Default: 10MB.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kafka_consumer_queue_capacity": schema.Int64Attribute{
								Description:         "'kafkaConsumerQueueCapacity' defines the capacity of the internal message queue used in the Kafka consumer client. Ignored when not using Kafka.",
								MarkdownDescription: "'kafkaConsumerQueueCapacity' defines the capacity of the internal message queue used in the Kafka consumer client. Ignored when not using Kafka.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kafka_consumer_replicas": schema.Int64Attribute{
								Description:         "'kafkaConsumerReplicas' defines the number of replicas (pods) to start for 'flowlogs-pipeline-transformer', which consumes Kafka messages. This setting is ignored when Kafka is disabled.",
								MarkdownDescription: "'kafkaConsumerReplicas' defines the number of replicas (pods) to start for 'flowlogs-pipeline-transformer', which consumes Kafka messages. This setting is ignored when Kafka is disabled.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
								},
							},

							"log_level": schema.StringAttribute{
								Description:         "'logLevel' of the processor runtime",
								MarkdownDescription: "'logLevel' of the processor runtime",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("trace", "debug", "info", "warn", "error", "fatal", "panic"),
								},
							},

							"log_types": schema.StringAttribute{
								Description:         "'logTypes' defines the desired record types to generate. Possible values are:<br> - 'Flows' (default) to export regular network flows<br> - 'Conversations' to generate events for started conversations, ended conversations as well as periodic 'tick' updates<br> - 'EndedConversations' to generate only ended conversations events<br> - 'All' to generate both network flows and all conversations events<br>",
								MarkdownDescription: "'logTypes' defines the desired record types to generate. Possible values are:<br> - 'Flows' (default) to export regular network flows<br> - 'Conversations' to generate events for started conversations, ended conversations as well as periodic 'tick' updates<br> - 'EndedConversations' to generate only ended conversations events<br> - 'All' to generate both network flows and all conversations events<br>",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Flows", "Conversations", "EndedConversations", "All"),
								},
							},

							"metrics": schema.SingleNestedAttribute{
								Description:         "'Metrics' define the processor configuration regarding metrics",
								MarkdownDescription: "'Metrics' define the processor configuration regarding metrics",
								Attributes: map[string]schema.Attribute{
									"disable_alerts": schema.ListAttribute{
										Description:         "'disableAlerts' is a list of alerts that should be disabled. Possible values are:<br> 'NetObservNoFlows', which is triggered when no flows are being observed for a certain period.<br> 'NetObservLokiError', which is triggered when flows are being dropped due to Loki errors.<br>",
										MarkdownDescription: "'disableAlerts' is a list of alerts that should be disabled. Possible values are:<br> 'NetObservNoFlows', which is triggered when no flows are being observed for a certain period.<br> 'NetObservLokiError', which is triggered when flows are being dropped due to Loki errors.<br>",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"include_list": schema.ListAttribute{
										Description:         "'includeList' is a list of metric names to specify which ones to generate. The names correspond to the names in Prometheus without the prefix. For example, 'namespace_egress_packets_total' shows up as 'netobserv_namespace_egress_packets_total' in Prometheus. Note that the more metrics you add, the bigger is the impact on Prometheus workload resources. Metrics enabled by default are: 'namespace_flows_total', 'node_ingress_bytes_total', 'workload_ingress_bytes_total', 'namespace_drop_packets_total' (when 'PacketDrop' feature is enabled), 'namespace_rtt_seconds' (when 'FlowRTT' feature is enabled), 'namespace_dns_latency_seconds' (when 'DNSTracking' feature is enabled). More information, with full list of available metrics: https://github.com/netobserv/network-observability-operator/blob/main/docs/Metrics.md",
										MarkdownDescription: "'includeList' is a list of metric names to specify which ones to generate. The names correspond to the names in Prometheus without the prefix. For example, 'namespace_egress_packets_total' shows up as 'netobserv_namespace_egress_packets_total' in Prometheus. Note that the more metrics you add, the bigger is the impact on Prometheus workload resources. Metrics enabled by default are: 'namespace_flows_total', 'node_ingress_bytes_total', 'workload_ingress_bytes_total', 'namespace_drop_packets_total' (when 'PacketDrop' feature is enabled), 'namespace_rtt_seconds' (when 'FlowRTT' feature is enabled), 'namespace_dns_latency_seconds' (when 'DNSTracking' feature is enabled). More information, with full list of available metrics: https://github.com/netobserv/network-observability-operator/blob/main/docs/Metrics.md",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"server": schema.SingleNestedAttribute{
										Description:         "Metrics server endpoint configuration for Prometheus scraper",
										MarkdownDescription: "Metrics server endpoint configuration for Prometheus scraper",
										Attributes: map[string]schema.Attribute{
											"port": schema.Int64Attribute{
												Description:         "The prometheus HTTP port",
												MarkdownDescription: "The prometheus HTTP port",
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
													"insecure_skip_verify": schema.BoolAttribute{
														Description:         "'insecureSkipVerify' allows skipping client-side verification of the provided certificate. If set to 'true', the 'providedCaFile' field is ignored.",
														MarkdownDescription: "'insecureSkipVerify' allows skipping client-side verification of the provided certificate. If set to 'true', the 'providedCaFile' field is ignored.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"provided": schema.SingleNestedAttribute{
														Description:         "TLS configuration when 'type' is set to 'Provided'.",
														MarkdownDescription: "TLS configuration when 'type' is set to 'Provided'.",
														Attributes: map[string]schema.Attribute{
															"cert_file": schema.StringAttribute{
																Description:         "'certFile' defines the path to the certificate file name within the config map or secret",
																MarkdownDescription: "'certFile' defines the path to the certificate file name within the config map or secret",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"cert_key": schema.StringAttribute{
																Description:         "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
																MarkdownDescription: "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "Name of the config map or secret containing certificates",
																MarkdownDescription: "Name of the config map or secret containing certificates",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
																MarkdownDescription: "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"type": schema.StringAttribute{
																Description:         "Type for the certificate reference: 'configmap' or 'secret'",
																MarkdownDescription: "Type for the certificate reference: 'configmap' or 'secret'",
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

													"provided_ca_file": schema.SingleNestedAttribute{
														Description:         "Reference to the CA file when 'type' is set to 'Provided'.",
														MarkdownDescription: "Reference to the CA file when 'type' is set to 'Provided'.",
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
														Description:         "Select the type of TLS configuration:<br> - 'Disabled' (default) to not configure TLS for the endpoint. - 'Provided' to manually provide cert file and a key file. - 'Auto' to use OpenShift auto generated certificate using annotations.",
														MarkdownDescription: "Select the type of TLS configuration:<br> - 'Disabled' (default) to not configure TLS for the endpoint. - 'Provided' to manually provide cert file and a key file. - 'Auto' to use OpenShift auto generated certificate using annotations.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("Disabled", "Provided", "Auto"),
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

							"multi_cluster_deployment": schema.BoolAttribute{
								Description:         "Set 'multiClusterDeployment' to 'true' to enable multi clusters feature. This adds 'clusterName' label to flows data",
								MarkdownDescription: "Set 'multiClusterDeployment' to 'true' to enable multi clusters feature. This adds 'clusterName' label to flows data",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resources": schema.SingleNestedAttribute{
								Description:         "'resources' are the compute resources required by this container. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "'resources' are the compute resources required by this container. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

func (r *FlowsNetobservIoFlowCollectorV1Beta2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_flows_netobserv_io_flow_collector_v1beta2_manifest")

	var model FlowsNetobservIoFlowCollectorV1Beta2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("flows.netobserv.io/v1beta2")
	model.Kind = pointer.String("FlowCollector")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
