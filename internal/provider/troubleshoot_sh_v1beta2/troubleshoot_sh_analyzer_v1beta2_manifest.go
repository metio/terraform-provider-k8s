/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package troubleshoot_sh_v1beta2

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
	_ datasource.DataSource = &TroubleshootShAnalyzerV1Beta2Manifest{}
)

func NewTroubleshootShAnalyzerV1Beta2Manifest() datasource.DataSource {
	return &TroubleshootShAnalyzerV1Beta2Manifest{}
}

type TroubleshootShAnalyzerV1Beta2Manifest struct{}

type TroubleshootShAnalyzerV1Beta2ManifestData struct {
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
		Analyzers *[]struct {
			CephStatus *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Namespace     *string            `tfsdk:"namespace" json:"namespace,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"ceph_status" json:"cephStatus,omitempty"`
			Certificates *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName   *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				Exclude     UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes    *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"certificates" json:"certificates,omitempty"`
			ClusterContainerStatuses *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName   *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				Exclude     UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Namespaces  *[]string          `tfsdk:"namespaces" json:"namespaces,omitempty"`
				Outcomes    *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				RestartCount *int64  `tfsdk:"restart_count" json:"restartCount,omitempty"`
				Strict       UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"cluster_container_statuses" json:"clusterContainerStatuses,omitempty"`
			ClusterPodStatuses *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName   *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				Exclude     UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Namespaces  *[]string          `tfsdk:"namespaces" json:"namespaces,omitempty"`
				Outcomes    *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"cluster_pod_statuses" json:"clusterPodStatuses,omitempty"`
			ClusterResource *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				ClusterScoped *bool              `tfsdk:"cluster_scoped" json:"clusterScoped,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				ExpectedValue *string            `tfsdk:"expected_value" json:"expectedValue,omitempty"`
				Kind          *string            `tfsdk:"kind" json:"kind,omitempty"`
				Name          *string            `tfsdk:"name" json:"name,omitempty"`
				Namespace     *string            `tfsdk:"namespace" json:"namespace,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Regex       *string `tfsdk:"regex" json:"regex,omitempty"`
				RegexGroups *string `tfsdk:"regex_groups" json:"regexGroups,omitempty"`
				Strict      UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
				YamlPath    *string `tfsdk:"yaml_path" json:"yamlPath,omitempty"`
			} `tfsdk:"cluster_resource" json:"clusterResource,omitempty"`
			ClusterVersion *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName   *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				Exclude     UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes    *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"cluster_version" json:"clusterVersion,omitempty"`
			ConfigMap *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				ConfigMapName *string            `tfsdk:"config_map_name" json:"configMapName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Key           *string            `tfsdk:"key" json:"key,omitempty"`
				Namespace     *string            `tfsdk:"namespace" json:"namespace,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"config_map" json:"configMap,omitempty"`
			ContainerRuntime *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName   *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				Exclude     UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes    *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"container_runtime" json:"containerRuntime,omitempty"`
			CustomResourceDefinition *struct {
				Annotations                  *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName                    *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CustomResourceDefinitionName *string            `tfsdk:"custom_resource_definition_name" json:"customResourceDefinitionName,omitempty"`
				Exclude                      UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes                     *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"custom_resource_definition" json:"customResourceDefinition,omitempty"`
			DeploymentStatus *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName   *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				Exclude     UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Name        *string            `tfsdk:"name" json:"name,omitempty"`
				Namespace   *string            `tfsdk:"namespace" json:"namespace,omitempty"`
				Namespaces  *[]string          `tfsdk:"namespaces" json:"namespaces,omitempty"`
				Outcomes    *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"deployment_status" json:"deploymentStatus,omitempty"`
			Distribution *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName   *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				Exclude     UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes    *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"distribution" json:"distribution,omitempty"`
			Event *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Kind          *string            `tfsdk:"kind" json:"kind,omitempty"`
				Namespace     *string            `tfsdk:"namespace" json:"namespace,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Reason *string `tfsdk:"reason" json:"reason,omitempty"`
				Regex  *string `tfsdk:"regex" json:"regex,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"event" json:"event,omitempty"`
			Goldpinger *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				FilePath      *string            `tfsdk:"file_path" json:"filePath,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"goldpinger" json:"goldpinger,omitempty"`
			Http *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"http" json:"http,omitempty"`
			ImagePullSecret *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName   *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				Exclude     UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes    *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				RegistryName *string `tfsdk:"registry_name" json:"registryName,omitempty"`
				Strict       UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"image_pull_secret" json:"imagePullSecret,omitempty"`
			Ingress *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName   *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				Exclude     UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				IngressName *string            `tfsdk:"ingress_name" json:"ingressName,omitempty"`
				Namespace   *string            `tfsdk:"namespace" json:"namespace,omitempty"`
				Outcomes    *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"ingress" json:"ingress,omitempty"`
			JobStatus *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName   *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				Exclude     UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Name        *string            `tfsdk:"name" json:"name,omitempty"`
				Namespace   *string            `tfsdk:"namespace" json:"namespace,omitempty"`
				Namespaces  *[]string          `tfsdk:"namespaces" json:"namespaces,omitempty"`
				Outcomes    *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"job_status" json:"jobStatus,omitempty"`
			JsonCompare *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				FileName      *string            `tfsdk:"file_name" json:"fileName,omitempty"`
				JsonPath      *string            `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Path   *string `tfsdk:"path" json:"path,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
				Value  *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"json_compare" json:"jsonCompare,omitempty"`
			Longhorn *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Namespace     *string            `tfsdk:"namespace" json:"namespace,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"longhorn" json:"longhorn,omitempty"`
			Mssql *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				FileName      *string            `tfsdk:"file_name" json:"fileName,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"mssql" json:"mssql,omitempty"`
			Mysql *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				FileName      *string            `tfsdk:"file_name" json:"fileName,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"mysql" json:"mysql,omitempty"`
			NodeMetrics *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Filters       *struct {
					Pvc *struct {
						NameRegex *string `tfsdk:"name_regex" json:"nameRegex,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"pvc" json:"pvc,omitempty"`
				} `tfsdk:"filters" json:"filters,omitempty"`
				Outcomes *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"node_metrics" json:"nodeMetrics,omitempty"`
			NodeResources *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName   *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				Exclude     UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Filters     *struct {
					CpuAllocatable              *string `tfsdk:"cpu_allocatable" json:"cpuAllocatable,omitempty"`
					CpuArchitecture             *string `tfsdk:"cpu_architecture" json:"cpuArchitecture,omitempty"`
					CpuCapacity                 *string `tfsdk:"cpu_capacity" json:"cpuCapacity,omitempty"`
					EphemeralStorageAllocatable *string `tfsdk:"ephemeral_storage_allocatable" json:"ephemeralStorageAllocatable,omitempty"`
					EphemeralStorageCapacity    *string `tfsdk:"ephemeral_storage_capacity" json:"ephemeralStorageCapacity,omitempty"`
					MemoryAllocatable           *string `tfsdk:"memory_allocatable" json:"memoryAllocatable,omitempty"`
					MemoryCapacity              *string `tfsdk:"memory_capacity" json:"memoryCapacity,omitempty"`
					PodAllocatable              *string `tfsdk:"pod_allocatable" json:"podAllocatable,omitempty"`
					PodCapacity                 *string `tfsdk:"pod_capacity" json:"podCapacity,omitempty"`
					ResourceAllocatable         *string `tfsdk:"resource_allocatable" json:"resourceAllocatable,omitempty"`
					ResourceCapacity            *string `tfsdk:"resource_capacity" json:"resourceCapacity,omitempty"`
					ResourceName                *string `tfsdk:"resource_name" json:"resourceName,omitempty"`
					Selector                    *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabel *map[string]string `tfsdk:"match_label" json:"matchLabel,omitempty"`
					} `tfsdk:"selector" json:"selector,omitempty"`
				} `tfsdk:"filters" json:"filters,omitempty"`
				Outcomes *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"node_resources" json:"nodeResources,omitempty"`
			Postgres *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				FileName      *string            `tfsdk:"file_name" json:"fileName,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"postgres" json:"postgres,omitempty"`
			Redis *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				FileName      *string            `tfsdk:"file_name" json:"fileName,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"redis" json:"redis,omitempty"`
			RegistryImages *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"registry_images" json:"registryImages,omitempty"`
			ReplicasetStatus *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName   *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				Exclude     UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Name        *string            `tfsdk:"name" json:"name,omitempty"`
				Namespace   *string            `tfsdk:"namespace" json:"namespace,omitempty"`
				Namespaces  *[]string          `tfsdk:"namespaces" json:"namespaces,omitempty"`
				Outcomes    *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Selector *[]string `tfsdk:"selector" json:"selector,omitempty"`
				Strict   UNKNOWN   `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"replicaset_status" json:"replicasetStatus,omitempty"`
			Secret *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName   *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				Exclude     UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Key         *string            `tfsdk:"key" json:"key,omitempty"`
				Namespace   *string            `tfsdk:"namespace" json:"namespace,omitempty"`
				Outcomes    *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
				Strict     UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"secret" json:"secret,omitempty"`
			StatefulsetStatus *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName   *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				Exclude     UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Name        *string            `tfsdk:"name" json:"name,omitempty"`
				Namespace   *string            `tfsdk:"namespace" json:"namespace,omitempty"`
				Namespaces  *[]string          `tfsdk:"namespaces" json:"namespaces,omitempty"`
				Outcomes    *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"statefulset_status" json:"statefulsetStatus,omitempty"`
			StorageClass *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName   *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				Exclude     UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes    *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				StorageClassName *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
				Strict           UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"storage_class" json:"storageClass,omitempty"`
			Sysctl *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName   *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				Exclude     UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes    *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"sysctl" json:"sysctl,omitempty"`
			TextAnalyze *struct {
				Annotations     *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName       *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName   *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude         UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				ExcludeFiles    *[]string          `tfsdk:"exclude_files" json:"excludeFiles,omitempty"`
				FileName        *string            `tfsdk:"file_name" json:"fileName,omitempty"`
				IgnoreIfNoFiles *bool              `tfsdk:"ignore_if_no_files" json:"ignoreIfNoFiles,omitempty"`
				Outcomes        *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Regex       *string `tfsdk:"regex" json:"regex,omitempty"`
				RegexGroups *string `tfsdk:"regex_groups" json:"regexGroups,omitempty"`
				Strict      UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"text_analyze" json:"textAnalyze,omitempty"`
			Velero *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName   *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				Exclude     UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Strict      UNKNOWN            `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"velero" json:"velero,omitempty"`
			WeaveReport *struct {
				Annotations    *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName      *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				Exclude        UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				ReportFileGlob *string            `tfsdk:"report_file_glob" json:"reportFileGlob,omitempty"`
				Strict         UNKNOWN            `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"weave_report" json:"weaveReport,omitempty"`
			YamlCompare *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				FileName      *string            `tfsdk:"file_name" json:"fileName,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Path   *string `tfsdk:"path" json:"path,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
				Value  *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"yaml_compare" json:"yamlCompare,omitempty"`
		} `tfsdk:"analyzers" json:"analyzers,omitempty"`
		HostAnalyzers *[]struct {
			BlockDevices *struct {
				Annotations                *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName                  *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName              *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude                    UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				IncludeUnmountedPartitions *bool              `tfsdk:"include_unmounted_partitions" json:"includeUnmountedPartitions,omitempty"`
				MinimumAcceptableSize      *int64             `tfsdk:"minimum_acceptable_size" json:"minimumAcceptableSize,omitempty"`
				Outcomes                   *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"block_devices" json:"blockDevices,omitempty"`
			Certificate *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"certificate" json:"certificate,omitempty"`
			CertificatesCollection *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"certificates_collection" json:"certificatesCollection,omitempty"`
			Cpu *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"cpu" json:"cpu,omitempty"`
			DiskUsage *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"disk_usage" json:"diskUsage,omitempty"`
			FilesystemPerformance *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"filesystem_performance" json:"filesystemPerformance,omitempty"`
			HostOS *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"host_os" json:"hostOS,omitempty"`
			HostServices *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"host_services" json:"hostServices,omitempty"`
			Http *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"http" json:"http,omitempty"`
			HttpLoadBalancer *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"http_load_balancer" json:"httpLoadBalancer,omitempty"`
			Ipv4Interfaces *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"ipv4_interfaces" json:"ipv4Interfaces,omitempty"`
			JsonCompare *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				FileName      *string            `tfsdk:"file_name" json:"fileName,omitempty"`
				JsonPath      *string            `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Path   *string `tfsdk:"path" json:"path,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
				Value  *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"json_compare" json:"jsonCompare,omitempty"`
			KernelConfigs *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				SelectedConfigs *[]string `tfsdk:"selected_configs" json:"selectedConfigs,omitempty"`
				Strict          UNKNOWN   `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"kernel_configs" json:"kernelConfigs,omitempty"`
			KernelModules *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"kernel_modules" json:"kernelModules,omitempty"`
			Memory *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"memory" json:"memory,omitempty"`
			NetworkNamespaceConnectivity *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"network_namespace_connectivity" json:"networkNamespaceConnectivity,omitempty"`
			SubnetAvailable *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"subnet_available" json:"subnetAvailable,omitempty"`
			SubnetContainsIP *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				Cidr          *string            `tfsdk:"cidr" json:"cidr,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Ip            *string            `tfsdk:"ip" json:"ip,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"subnet_contains_ip" json:"subnetContainsIP,omitempty"`
			Sysctl *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"sysctl" json:"sysctl,omitempty"`
			SystemPackages *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"system_packages" json:"systemPackages,omitempty"`
			TcpConnect *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"tcp_connect" json:"tcpConnect,omitempty"`
			TcpLoadBalancer *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"tcp_load_balancer" json:"tcpLoadBalancer,omitempty"`
			TcpPortStatus *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"tcp_port_status" json:"tcpPortStatus,omitempty"`
			TextAnalyze *struct {
				Annotations     *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName       *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName   *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude         UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				ExcludeFiles    *[]string          `tfsdk:"exclude_files" json:"excludeFiles,omitempty"`
				FileName        *string            `tfsdk:"file_name" json:"fileName,omitempty"`
				IgnoreIfNoFiles *bool              `tfsdk:"ignore_if_no_files" json:"ignoreIfNoFiles,omitempty"`
				Outcomes        *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Regex       *string `tfsdk:"regex" json:"regex,omitempty"`
				RegexGroups *string `tfsdk:"regex_groups" json:"regexGroups,omitempty"`
				Strict      UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"text_analyze" json:"textAnalyze,omitempty"`
			Time *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"time" json:"time,omitempty"`
			UdpPortStatus *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"udp_port_status" json:"udpPortStatus,omitempty"`
		} `tfsdk:"host_analyzers" json:"hostAnalyzers,omitempty"`
		Uri *string `tfsdk:"uri" json:"uri,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *TroubleshootShAnalyzerV1Beta2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_troubleshoot_sh_analyzer_v1beta2_manifest"
}

func (r *TroubleshootShAnalyzerV1Beta2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Analyzer is the Schema for the analyzers API",
		MarkdownDescription: "Analyzer is the Schema for the analyzers API",
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
				Description:         "AnalyzerSpec defines the desired state of Analyzer",
				MarkdownDescription: "AnalyzerSpec defines the desired state of Analyzer",
				Attributes: map[string]schema.Attribute{
					"analyzers": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"ceph_status": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"namespace": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"certificates": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"cluster_container_statuses": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"namespaces": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"restart_count": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"cluster_pod_statuses": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"namespaces": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"cluster_resource": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"cluster_scoped": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"expected_value": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"kind": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"namespace": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"regex": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"regex_groups": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"yaml_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"cluster_version": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"config_map": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"config_map_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"key": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"namespace": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"container_runtime": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"custom_resource_definition": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"custom_resource_definition_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"deployment_status": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"namespace": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"namespaces": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"distribution": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"event": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"kind": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"namespace": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"reason": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"regex": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"goldpinger": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"file_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"http": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"image_pull_secret": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"registry_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"ingress": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ingress_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"namespace": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"job_status": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"namespace": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"namespaces": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"json_compare": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"file_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"longhorn": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"namespace": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"mssql": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"file_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"mysql": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"file_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"node_metrics": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"filters": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"pvc": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"name_regex": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespace": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
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

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"node_resources": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"filters": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"cpu_allocatable": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"cpu_architecture": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"cpu_capacity": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"ephemeral_storage_allocatable": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"ephemeral_storage_capacity": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"memory_allocatable": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"memory_capacity": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"pod_allocatable": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"pod_capacity": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"resource_allocatable": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"resource_capacity": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"resource_name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"selector": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"match_expressions": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
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

														"match_label": schema.MapAttribute{
															Description:         "",
															MarkdownDescription: "",
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

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"postgres": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"file_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"redis": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"file_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"registry_images": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"replicaset_status": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"namespace": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"namespaces": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"selector": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"secret": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"key": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"namespace": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"secret_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"statefulset_status": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"namespace": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"namespaces": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"storage_class": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"storage_class_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"sysctl": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"text_analyze": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude_files": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"file_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ignore_if_no_files": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"regex": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"regex_groups": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"velero": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"weave_report": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"report_file_glob": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"yaml_compare": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"file_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"host_analyzers": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"block_devices": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"include_unmounted_partitions": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"minimum_acceptable_size": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"certificate": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"certificates_collection": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"cpu": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"disk_usage": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"filesystem_performance": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"host_os": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"host_services": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"http": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"http_load_balancer": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"ipv4_interfaces": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"json_compare": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"file_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"kernel_configs": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"selected_configs": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"kernel_modules": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"memory": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"network_namespace_connectivity": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"subnet_available": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"subnet_contains_ip": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"cidr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ip": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"sysctl": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"system_packages": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"tcp_connect": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"tcp_load_balancer": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"tcp_port_status": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"text_analyze": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude_files": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"file_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ignore_if_no_files": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"regex": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"regex_groups": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"time": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"udp_port_status": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"uri": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
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
	}
}

func (r *TroubleshootShAnalyzerV1Beta2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_troubleshoot_sh_analyzer_v1beta2_manifest")

	var model TroubleshootShAnalyzerV1Beta2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("troubleshoot.sh/v1beta2")
	model.Kind = pointer.String("Analyzer")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
