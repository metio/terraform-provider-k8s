/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package chaos_mesh_org_v1alpha1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	k8sTypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"strings"
)

var (
	_ resource.Resource                = &ChaosMeshOrgWorkflowV1Alpha1Resource{}
	_ resource.ResourceWithConfigure   = &ChaosMeshOrgWorkflowV1Alpha1Resource{}
	_ resource.ResourceWithImportState = &ChaosMeshOrgWorkflowV1Alpha1Resource{}
)

func NewChaosMeshOrgWorkflowV1Alpha1Resource() resource.Resource {
	return &ChaosMeshOrgWorkflowV1Alpha1Resource{}
}

type ChaosMeshOrgWorkflowV1Alpha1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type ChaosMeshOrgWorkflowV1Alpha1ResourceData struct {
	ID             types.String `tfsdk:"id" json:"-"`
	ForceConflicts types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager   types.String `tfsdk:"field_manager" json:"-"`
	WaitFor        types.List   `tfsdk:"wait_for" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Entry     *string `tfsdk:"entry" json:"entry,omitempty"`
		Templates *[]struct {
			AbortWithStatusCheck *bool `tfsdk:"abort_with_status_check" json:"abortWithStatusCheck,omitempty"`
			AwsChaos             *struct {
				Action        *string `tfsdk:"action" json:"action,omitempty"`
				AwsRegion     *string `tfsdk:"aws_region" json:"awsRegion,omitempty"`
				DeviceName    *string `tfsdk:"device_name" json:"deviceName,omitempty"`
				Duration      *string `tfsdk:"duration" json:"duration,omitempty"`
				Ec2Instance   *string `tfsdk:"ec2_instance" json:"ec2Instance,omitempty"`
				Endpoint      *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
				RemoteCluster *string `tfsdk:"remote_cluster" json:"remoteCluster,omitempty"`
				SecretName    *string `tfsdk:"secret_name" json:"secretName,omitempty"`
				VolumeID      *string `tfsdk:"volume_id" json:"volumeID,omitempty"`
			} `tfsdk:"aws_chaos" json:"awsChaos,omitempty"`
			AzureChaos *struct {
				Action            *string `tfsdk:"action" json:"action,omitempty"`
				DiskName          *string `tfsdk:"disk_name" json:"diskName,omitempty"`
				Duration          *string `tfsdk:"duration" json:"duration,omitempty"`
				Lun               *int64  `tfsdk:"lun" json:"lun,omitempty"`
				RemoteCluster     *string `tfsdk:"remote_cluster" json:"remoteCluster,omitempty"`
				ResourceGroupName *string `tfsdk:"resource_group_name" json:"resourceGroupName,omitempty"`
				SecretName        *string `tfsdk:"secret_name" json:"secretName,omitempty"`
				SubscriptionID    *string `tfsdk:"subscription_id" json:"subscriptionID,omitempty"`
				VmName            *string `tfsdk:"vm_name" json:"vmName,omitempty"`
			} `tfsdk:"azure_chaos" json:"azureChaos,omitempty"`
			BlockChaos *struct {
				Action         *string   `tfsdk:"action" json:"action,omitempty"`
				ContainerNames *[]string `tfsdk:"container_names" json:"containerNames,omitempty"`
				Delay          *struct {
					Correlation *string `tfsdk:"correlation" json:"correlation,omitempty"`
					Jitter      *string `tfsdk:"jitter" json:"jitter,omitempty"`
					Latency     *string `tfsdk:"latency" json:"latency,omitempty"`
				} `tfsdk:"delay" json:"delay,omitempty"`
				Duration      *string `tfsdk:"duration" json:"duration,omitempty"`
				Mode          *string `tfsdk:"mode" json:"mode,omitempty"`
				RemoteCluster *string `tfsdk:"remote_cluster" json:"remoteCluster,omitempty"`
				Selector      *struct {
					AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" json:"annotationSelectors,omitempty"`
					ExpressionSelectors *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"expression_selectors" json:"expressionSelectors,omitempty"`
					FieldSelectors    *map[string]string   `tfsdk:"field_selectors" json:"fieldSelectors,omitempty"`
					LabelSelectors    *map[string]string   `tfsdk:"label_selectors" json:"labelSelectors,omitempty"`
					Namespaces        *[]string            `tfsdk:"namespaces" json:"namespaces,omitempty"`
					NodeSelectors     *map[string]string   `tfsdk:"node_selectors" json:"nodeSelectors,omitempty"`
					Nodes             *[]string            `tfsdk:"nodes" json:"nodes,omitempty"`
					PodPhaseSelectors *[]string            `tfsdk:"pod_phase_selectors" json:"podPhaseSelectors,omitempty"`
					Pods              *map[string][]string `tfsdk:"pods" json:"pods,omitempty"`
				} `tfsdk:"selector" json:"selector,omitempty"`
				Value      *string `tfsdk:"value" json:"value,omitempty"`
				VolumeName *string `tfsdk:"volume_name" json:"volumeName,omitempty"`
			} `tfsdk:"block_chaos" json:"blockChaos,omitempty"`
			Children            *[]string `tfsdk:"children" json:"children,omitempty"`
			ConditionalBranches *[]struct {
				Expression *string `tfsdk:"expression" json:"expression,omitempty"`
				Target     *string `tfsdk:"target" json:"target,omitempty"`
			} `tfsdk:"conditional_branches" json:"conditionalBranches,omitempty"`
			Deadline *string `tfsdk:"deadline" json:"deadline,omitempty"`
			DnsChaos *struct {
				Action         *string   `tfsdk:"action" json:"action,omitempty"`
				ContainerNames *[]string `tfsdk:"container_names" json:"containerNames,omitempty"`
				Duration       *string   `tfsdk:"duration" json:"duration,omitempty"`
				Mode           *string   `tfsdk:"mode" json:"mode,omitempty"`
				Patterns       *[]string `tfsdk:"patterns" json:"patterns,omitempty"`
				RemoteCluster  *string   `tfsdk:"remote_cluster" json:"remoteCluster,omitempty"`
				Selector       *struct {
					AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" json:"annotationSelectors,omitempty"`
					ExpressionSelectors *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"expression_selectors" json:"expressionSelectors,omitempty"`
					FieldSelectors    *map[string]string   `tfsdk:"field_selectors" json:"fieldSelectors,omitempty"`
					LabelSelectors    *map[string]string   `tfsdk:"label_selectors" json:"labelSelectors,omitempty"`
					Namespaces        *[]string            `tfsdk:"namespaces" json:"namespaces,omitempty"`
					NodeSelectors     *map[string]string   `tfsdk:"node_selectors" json:"nodeSelectors,omitempty"`
					Nodes             *[]string            `tfsdk:"nodes" json:"nodes,omitempty"`
					PodPhaseSelectors *[]string            `tfsdk:"pod_phase_selectors" json:"podPhaseSelectors,omitempty"`
					Pods              *map[string][]string `tfsdk:"pods" json:"pods,omitempty"`
				} `tfsdk:"selector" json:"selector,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"dns_chaos" json:"dnsChaos,omitempty"`
			GcpChaos *struct {
				Action        *string   `tfsdk:"action" json:"action,omitempty"`
				DeviceNames   *[]string `tfsdk:"device_names" json:"deviceNames,omitempty"`
				Duration      *string   `tfsdk:"duration" json:"duration,omitempty"`
				Instance      *string   `tfsdk:"instance" json:"instance,omitempty"`
				Project       *string   `tfsdk:"project" json:"project,omitempty"`
				RemoteCluster *string   `tfsdk:"remote_cluster" json:"remoteCluster,omitempty"`
				SecretName    *string   `tfsdk:"secret_name" json:"secretName,omitempty"`
				Zone          *string   `tfsdk:"zone" json:"zone,omitempty"`
			} `tfsdk:"gcp_chaos" json:"gcpChaos,omitempty"`
			HttpChaos *struct {
				Abort    *bool   `tfsdk:"abort" json:"abort,omitempty"`
				Code     *int64  `tfsdk:"code" json:"code,omitempty"`
				Delay    *string `tfsdk:"delay" json:"delay,omitempty"`
				Duration *string `tfsdk:"duration" json:"duration,omitempty"`
				Method   *string `tfsdk:"method" json:"method,omitempty"`
				Mode     *string `tfsdk:"mode" json:"mode,omitempty"`
				Patch    *struct {
					Body *struct {
						Type  *string `tfsdk:"type" json:"type,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"body" json:"body,omitempty"`
					Headers *[]string `tfsdk:"headers" json:"headers,omitempty"`
					Queries *[]string `tfsdk:"queries" json:"queries,omitempty"`
				} `tfsdk:"patch" json:"patch,omitempty"`
				Path          *string `tfsdk:"path" json:"path,omitempty"`
				Port          *int64  `tfsdk:"port" json:"port,omitempty"`
				RemoteCluster *string `tfsdk:"remote_cluster" json:"remoteCluster,omitempty"`
				Replace       *struct {
					Body    *string            `tfsdk:"body" json:"body,omitempty"`
					Code    *int64             `tfsdk:"code" json:"code,omitempty"`
					Headers *map[string]string `tfsdk:"headers" json:"headers,omitempty"`
					Method  *string            `tfsdk:"method" json:"method,omitempty"`
					Path    *string            `tfsdk:"path" json:"path,omitempty"`
					Queries *map[string]string `tfsdk:"queries" json:"queries,omitempty"`
				} `tfsdk:"replace" json:"replace,omitempty"`
				Request_headers  *map[string]string `tfsdk:"request_headers" json:"request_headers,omitempty"`
				Response_headers *map[string]string `tfsdk:"response_headers" json:"response_headers,omitempty"`
				Selector         *struct {
					AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" json:"annotationSelectors,omitempty"`
					ExpressionSelectors *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"expression_selectors" json:"expressionSelectors,omitempty"`
					FieldSelectors    *map[string]string   `tfsdk:"field_selectors" json:"fieldSelectors,omitempty"`
					LabelSelectors    *map[string]string   `tfsdk:"label_selectors" json:"labelSelectors,omitempty"`
					Namespaces        *[]string            `tfsdk:"namespaces" json:"namespaces,omitempty"`
					NodeSelectors     *map[string]string   `tfsdk:"node_selectors" json:"nodeSelectors,omitempty"`
					Nodes             *[]string            `tfsdk:"nodes" json:"nodes,omitempty"`
					PodPhaseSelectors *[]string            `tfsdk:"pod_phase_selectors" json:"podPhaseSelectors,omitempty"`
					Pods              *map[string][]string `tfsdk:"pods" json:"pods,omitempty"`
				} `tfsdk:"selector" json:"selector,omitempty"`
				Target *string `tfsdk:"target" json:"target,omitempty"`
				Tls    *struct {
					CaName          *string `tfsdk:"ca_name" json:"caName,omitempty"`
					CertName        *string `tfsdk:"cert_name" json:"certName,omitempty"`
					KeyName         *string `tfsdk:"key_name" json:"keyName,omitempty"`
					SecretName      *string `tfsdk:"secret_name" json:"secretName,omitempty"`
					SecretNamespace *string `tfsdk:"secret_namespace" json:"secretNamespace,omitempty"`
				} `tfsdk:"tls" json:"tls,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"http_chaos" json:"httpChaos,omitempty"`
			IoChaos *struct {
				Action *string `tfsdk:"action" json:"action,omitempty"`
				Attr   *struct {
					Atime *struct {
						Nsec *int64 `tfsdk:"nsec" json:"nsec,omitempty"`
						Sec  *int64 `tfsdk:"sec" json:"sec,omitempty"`
					} `tfsdk:"atime" json:"atime,omitempty"`
					Blocks *int64 `tfsdk:"blocks" json:"blocks,omitempty"`
					Ctime  *struct {
						Nsec *int64 `tfsdk:"nsec" json:"nsec,omitempty"`
						Sec  *int64 `tfsdk:"sec" json:"sec,omitempty"`
					} `tfsdk:"ctime" json:"ctime,omitempty"`
					Gid   *int64  `tfsdk:"gid" json:"gid,omitempty"`
					Ino   *int64  `tfsdk:"ino" json:"ino,omitempty"`
					Kind  *string `tfsdk:"kind" json:"kind,omitempty"`
					Mtime *struct {
						Nsec *int64 `tfsdk:"nsec" json:"nsec,omitempty"`
						Sec  *int64 `tfsdk:"sec" json:"sec,omitempty"`
					} `tfsdk:"mtime" json:"mtime,omitempty"`
					Nlink *int64 `tfsdk:"nlink" json:"nlink,omitempty"`
					Perm  *int64 `tfsdk:"perm" json:"perm,omitempty"`
					Rdev  *int64 `tfsdk:"rdev" json:"rdev,omitempty"`
					Size  *int64 `tfsdk:"size" json:"size,omitempty"`
					Uid   *int64 `tfsdk:"uid" json:"uid,omitempty"`
				} `tfsdk:"attr" json:"attr,omitempty"`
				ContainerNames *[]string `tfsdk:"container_names" json:"containerNames,omitempty"`
				Delay          *string   `tfsdk:"delay" json:"delay,omitempty"`
				Duration       *string   `tfsdk:"duration" json:"duration,omitempty"`
				Errno          *int64    `tfsdk:"errno" json:"errno,omitempty"`
				Methods        *[]string `tfsdk:"methods" json:"methods,omitempty"`
				Mistake        *struct {
					Filling        *string `tfsdk:"filling" json:"filling,omitempty"`
					MaxLength      *int64  `tfsdk:"max_length" json:"maxLength,omitempty"`
					MaxOccurrences *int64  `tfsdk:"max_occurrences" json:"maxOccurrences,omitempty"`
				} `tfsdk:"mistake" json:"mistake,omitempty"`
				Mode          *string `tfsdk:"mode" json:"mode,omitempty"`
				Path          *string `tfsdk:"path" json:"path,omitempty"`
				Percent       *int64  `tfsdk:"percent" json:"percent,omitempty"`
				RemoteCluster *string `tfsdk:"remote_cluster" json:"remoteCluster,omitempty"`
				Selector      *struct {
					AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" json:"annotationSelectors,omitempty"`
					ExpressionSelectors *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"expression_selectors" json:"expressionSelectors,omitempty"`
					FieldSelectors    *map[string]string   `tfsdk:"field_selectors" json:"fieldSelectors,omitempty"`
					LabelSelectors    *map[string]string   `tfsdk:"label_selectors" json:"labelSelectors,omitempty"`
					Namespaces        *[]string            `tfsdk:"namespaces" json:"namespaces,omitempty"`
					NodeSelectors     *map[string]string   `tfsdk:"node_selectors" json:"nodeSelectors,omitempty"`
					Nodes             *[]string            `tfsdk:"nodes" json:"nodes,omitempty"`
					PodPhaseSelectors *[]string            `tfsdk:"pod_phase_selectors" json:"podPhaseSelectors,omitempty"`
					Pods              *map[string][]string `tfsdk:"pods" json:"pods,omitempty"`
				} `tfsdk:"selector" json:"selector,omitempty"`
				Value      *string `tfsdk:"value" json:"value,omitempty"`
				VolumePath *string `tfsdk:"volume_path" json:"volumePath,omitempty"`
			} `tfsdk:"io_chaos" json:"ioChaos,omitempty"`
			JvmChaos *struct {
				Action                *string   `tfsdk:"action" json:"action,omitempty"`
				Class                 *string   `tfsdk:"class" json:"class,omitempty"`
				ContainerNames        *[]string `tfsdk:"container_names" json:"containerNames,omitempty"`
				CpuCount              *int64    `tfsdk:"cpu_count" json:"cpuCount,omitempty"`
				Database              *string   `tfsdk:"database" json:"database,omitempty"`
				Duration              *string   `tfsdk:"duration" json:"duration,omitempty"`
				Exception             *string   `tfsdk:"exception" json:"exception,omitempty"`
				Latency               *int64    `tfsdk:"latency" json:"latency,omitempty"`
				MemType               *string   `tfsdk:"mem_type" json:"memType,omitempty"`
				Method                *string   `tfsdk:"method" json:"method,omitempty"`
				Mode                  *string   `tfsdk:"mode" json:"mode,omitempty"`
				MysqlConnectorVersion *string   `tfsdk:"mysql_connector_version" json:"mysqlConnectorVersion,omitempty"`
				Name                  *string   `tfsdk:"name" json:"name,omitempty"`
				Pid                   *int64    `tfsdk:"pid" json:"pid,omitempty"`
				Port                  *int64    `tfsdk:"port" json:"port,omitempty"`
				RemoteCluster         *string   `tfsdk:"remote_cluster" json:"remoteCluster,omitempty"`
				RuleData              *string   `tfsdk:"rule_data" json:"ruleData,omitempty"`
				Selector              *struct {
					AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" json:"annotationSelectors,omitempty"`
					ExpressionSelectors *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"expression_selectors" json:"expressionSelectors,omitempty"`
					FieldSelectors    *map[string]string   `tfsdk:"field_selectors" json:"fieldSelectors,omitempty"`
					LabelSelectors    *map[string]string   `tfsdk:"label_selectors" json:"labelSelectors,omitempty"`
					Namespaces        *[]string            `tfsdk:"namespaces" json:"namespaces,omitempty"`
					NodeSelectors     *map[string]string   `tfsdk:"node_selectors" json:"nodeSelectors,omitempty"`
					Nodes             *[]string            `tfsdk:"nodes" json:"nodes,omitempty"`
					PodPhaseSelectors *[]string            `tfsdk:"pod_phase_selectors" json:"podPhaseSelectors,omitempty"`
					Pods              *map[string][]string `tfsdk:"pods" json:"pods,omitempty"`
				} `tfsdk:"selector" json:"selector,omitempty"`
				SqlType *string `tfsdk:"sql_type" json:"sqlType,omitempty"`
				Table   *string `tfsdk:"table" json:"table,omitempty"`
				Value   *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"jvm_chaos" json:"jvmChaos,omitempty"`
			KernelChaos *struct {
				ContainerNames  *[]string `tfsdk:"container_names" json:"containerNames,omitempty"`
				Duration        *string   `tfsdk:"duration" json:"duration,omitempty"`
				FailKernRequest *struct {
					Callchain *[]struct {
						Funcname   *string `tfsdk:"funcname" json:"funcname,omitempty"`
						Parameters *string `tfsdk:"parameters" json:"parameters,omitempty"`
						Predicate  *string `tfsdk:"predicate" json:"predicate,omitempty"`
					} `tfsdk:"callchain" json:"callchain,omitempty"`
					Failtype    *int64    `tfsdk:"failtype" json:"failtype,omitempty"`
					Headers     *[]string `tfsdk:"headers" json:"headers,omitempty"`
					Probability *int64    `tfsdk:"probability" json:"probability,omitempty"`
					Times       *int64    `tfsdk:"times" json:"times,omitempty"`
				} `tfsdk:"fail_kern_request" json:"failKernRequest,omitempty"`
				Mode          *string `tfsdk:"mode" json:"mode,omitempty"`
				RemoteCluster *string `tfsdk:"remote_cluster" json:"remoteCluster,omitempty"`
				Selector      *struct {
					AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" json:"annotationSelectors,omitempty"`
					ExpressionSelectors *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"expression_selectors" json:"expressionSelectors,omitempty"`
					FieldSelectors    *map[string]string   `tfsdk:"field_selectors" json:"fieldSelectors,omitempty"`
					LabelSelectors    *map[string]string   `tfsdk:"label_selectors" json:"labelSelectors,omitempty"`
					Namespaces        *[]string            `tfsdk:"namespaces" json:"namespaces,omitempty"`
					NodeSelectors     *map[string]string   `tfsdk:"node_selectors" json:"nodeSelectors,omitempty"`
					Nodes             *[]string            `tfsdk:"nodes" json:"nodes,omitempty"`
					PodPhaseSelectors *[]string            `tfsdk:"pod_phase_selectors" json:"podPhaseSelectors,omitempty"`
					Pods              *map[string][]string `tfsdk:"pods" json:"pods,omitempty"`
				} `tfsdk:"selector" json:"selector,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"kernel_chaos" json:"kernelChaos,omitempty"`
			Name         *string `tfsdk:"name" json:"name,omitempty"`
			NetworkChaos *struct {
				Action    *string `tfsdk:"action" json:"action,omitempty"`
				Bandwidth *struct {
					Buffer   *int64  `tfsdk:"buffer" json:"buffer,omitempty"`
					Limit    *int64  `tfsdk:"limit" json:"limit,omitempty"`
					Minburst *int64  `tfsdk:"minburst" json:"minburst,omitempty"`
					Peakrate *int64  `tfsdk:"peakrate" json:"peakrate,omitempty"`
					Rate     *string `tfsdk:"rate" json:"rate,omitempty"`
				} `tfsdk:"bandwidth" json:"bandwidth,omitempty"`
				Corrupt *struct {
					Correlation *string `tfsdk:"correlation" json:"correlation,omitempty"`
					Corrupt     *string `tfsdk:"corrupt" json:"corrupt,omitempty"`
				} `tfsdk:"corrupt" json:"corrupt,omitempty"`
				Delay *struct {
					Correlation *string `tfsdk:"correlation" json:"correlation,omitempty"`
					Jitter      *string `tfsdk:"jitter" json:"jitter,omitempty"`
					Latency     *string `tfsdk:"latency" json:"latency,omitempty"`
					Reorder     *struct {
						Correlation *string `tfsdk:"correlation" json:"correlation,omitempty"`
						Gap         *int64  `tfsdk:"gap" json:"gap,omitempty"`
						Reorder     *string `tfsdk:"reorder" json:"reorder,omitempty"`
					} `tfsdk:"reorder" json:"reorder,omitempty"`
				} `tfsdk:"delay" json:"delay,omitempty"`
				Device    *string `tfsdk:"device" json:"device,omitempty"`
				Direction *string `tfsdk:"direction" json:"direction,omitempty"`
				Duplicate *struct {
					Correlation *string `tfsdk:"correlation" json:"correlation,omitempty"`
					Duplicate   *string `tfsdk:"duplicate" json:"duplicate,omitempty"`
				} `tfsdk:"duplicate" json:"duplicate,omitempty"`
				Duration        *string   `tfsdk:"duration" json:"duration,omitempty"`
				ExternalTargets *[]string `tfsdk:"external_targets" json:"externalTargets,omitempty"`
				Loss            *struct {
					Correlation *string `tfsdk:"correlation" json:"correlation,omitempty"`
					Loss        *string `tfsdk:"loss" json:"loss,omitempty"`
				} `tfsdk:"loss" json:"loss,omitempty"`
				Mode          *string `tfsdk:"mode" json:"mode,omitempty"`
				RemoteCluster *string `tfsdk:"remote_cluster" json:"remoteCluster,omitempty"`
				Selector      *struct {
					AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" json:"annotationSelectors,omitempty"`
					ExpressionSelectors *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"expression_selectors" json:"expressionSelectors,omitempty"`
					FieldSelectors    *map[string]string   `tfsdk:"field_selectors" json:"fieldSelectors,omitempty"`
					LabelSelectors    *map[string]string   `tfsdk:"label_selectors" json:"labelSelectors,omitempty"`
					Namespaces        *[]string            `tfsdk:"namespaces" json:"namespaces,omitempty"`
					NodeSelectors     *map[string]string   `tfsdk:"node_selectors" json:"nodeSelectors,omitempty"`
					Nodes             *[]string            `tfsdk:"nodes" json:"nodes,omitempty"`
					PodPhaseSelectors *[]string            `tfsdk:"pod_phase_selectors" json:"podPhaseSelectors,omitempty"`
					Pods              *map[string][]string `tfsdk:"pods" json:"pods,omitempty"`
				} `tfsdk:"selector" json:"selector,omitempty"`
				Target *struct {
					Mode     *string `tfsdk:"mode" json:"mode,omitempty"`
					Selector *struct {
						AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" json:"annotationSelectors,omitempty"`
						ExpressionSelectors *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"expression_selectors" json:"expressionSelectors,omitempty"`
						FieldSelectors    *map[string]string   `tfsdk:"field_selectors" json:"fieldSelectors,omitempty"`
						LabelSelectors    *map[string]string   `tfsdk:"label_selectors" json:"labelSelectors,omitempty"`
						Namespaces        *[]string            `tfsdk:"namespaces" json:"namespaces,omitempty"`
						NodeSelectors     *map[string]string   `tfsdk:"node_selectors" json:"nodeSelectors,omitempty"`
						Nodes             *[]string            `tfsdk:"nodes" json:"nodes,omitempty"`
						PodPhaseSelectors *[]string            `tfsdk:"pod_phase_selectors" json:"podPhaseSelectors,omitempty"`
						Pods              *map[string][]string `tfsdk:"pods" json:"pods,omitempty"`
					} `tfsdk:"selector" json:"selector,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"target" json:"target,omitempty"`
				TargetDevice *string `tfsdk:"target_device" json:"targetDevice,omitempty"`
				Value        *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"network_chaos" json:"networkChaos,omitempty"`
			PhysicalmachineChaos *struct {
				Action  *string   `tfsdk:"action" json:"action,omitempty"`
				Address *[]string `tfsdk:"address" json:"address,omitempty"`
				Clock   *struct {
					Clock_ids_slice *string `tfsdk:"clock_ids_slice" json:"clock-ids-slice,omitempty"`
					Pid             *int64  `tfsdk:"pid" json:"pid,omitempty"`
					Time_offset     *string `tfsdk:"time_offset" json:"time-offset,omitempty"`
				} `tfsdk:"clock" json:"clock,omitempty"`
				Disk_fill *struct {
					Fill_by_fallocate *bool   `tfsdk:"fill_by_fallocate" json:"fill-by-fallocate,omitempty"`
					Path              *string `tfsdk:"path" json:"path,omitempty"`
					Size              *string `tfsdk:"size" json:"size,omitempty"`
				} `tfsdk:"disk_fill" json:"disk-fill,omitempty"`
				Disk_read_payload *struct {
					Path                *string `tfsdk:"path" json:"path,omitempty"`
					Payload_process_num *int64  `tfsdk:"payload_process_num" json:"payload-process-num,omitempty"`
					Size                *string `tfsdk:"size" json:"size,omitempty"`
				} `tfsdk:"disk_read_payload" json:"disk-read-payload,omitempty"`
				Disk_write_payload *struct {
					Path                *string `tfsdk:"path" json:"path,omitempty"`
					Payload_process_num *int64  `tfsdk:"payload_process_num" json:"payload-process-num,omitempty"`
					Size                *string `tfsdk:"size" json:"size,omitempty"`
				} `tfsdk:"disk_write_payload" json:"disk-write-payload,omitempty"`
				Duration    *string `tfsdk:"duration" json:"duration,omitempty"`
				File_append *struct {
					Count     *int64  `tfsdk:"count" json:"count,omitempty"`
					Data      *string `tfsdk:"data" json:"data,omitempty"`
					File_name *string `tfsdk:"file_name" json:"file-name,omitempty"`
				} `tfsdk:"file_append" json:"file-append,omitempty"`
				File_create *struct {
					Dir_name  *string `tfsdk:"dir_name" json:"dir-name,omitempty"`
					File_name *string `tfsdk:"file_name" json:"file-name,omitempty"`
				} `tfsdk:"file_create" json:"file-create,omitempty"`
				File_delete *struct {
					Dir_name  *string `tfsdk:"dir_name" json:"dir-name,omitempty"`
					File_name *string `tfsdk:"file_name" json:"file-name,omitempty"`
				} `tfsdk:"file_delete" json:"file-delete,omitempty"`
				File_modify *struct {
					File_name *string `tfsdk:"file_name" json:"file-name,omitempty"`
					Privilege *int64  `tfsdk:"privilege" json:"privilege,omitempty"`
				} `tfsdk:"file_modify" json:"file-modify,omitempty"`
				File_rename *struct {
					Dest_file   *string `tfsdk:"dest_file" json:"dest-file,omitempty"`
					Source_file *string `tfsdk:"source_file" json:"source-file,omitempty"`
				} `tfsdk:"file_rename" json:"file-rename,omitempty"`
				File_replace *struct {
					Dest_string   *string `tfsdk:"dest_string" json:"dest-string,omitempty"`
					File_name     *string `tfsdk:"file_name" json:"file-name,omitempty"`
					Line          *int64  `tfsdk:"line" json:"line,omitempty"`
					Origin_string *string `tfsdk:"origin_string" json:"origin-string,omitempty"`
				} `tfsdk:"file_replace" json:"file-replace,omitempty"`
				Http_abort *struct {
					Code        *string   `tfsdk:"code" json:"code,omitempty"`
					Method      *string   `tfsdk:"method" json:"method,omitempty"`
					Path        *string   `tfsdk:"path" json:"path,omitempty"`
					Port        *int64    `tfsdk:"port" json:"port,omitempty"`
					Proxy_ports *[]string `tfsdk:"proxy_ports" json:"proxy_ports,omitempty"`
					Target      *string   `tfsdk:"target" json:"target,omitempty"`
				} `tfsdk:"http_abort" json:"http-abort,omitempty"`
				Http_config *struct {
					File_path *string `tfsdk:"file_path" json:"file_path,omitempty"`
				} `tfsdk:"http_config" json:"http-config,omitempty"`
				Http_delay *struct {
					Code        *string   `tfsdk:"code" json:"code,omitempty"`
					Delay       *string   `tfsdk:"delay" json:"delay,omitempty"`
					Method      *string   `tfsdk:"method" json:"method,omitempty"`
					Path        *string   `tfsdk:"path" json:"path,omitempty"`
					Port        *int64    `tfsdk:"port" json:"port,omitempty"`
					Proxy_ports *[]string `tfsdk:"proxy_ports" json:"proxy_ports,omitempty"`
					Target      *string   `tfsdk:"target" json:"target,omitempty"`
				} `tfsdk:"http_delay" json:"http-delay,omitempty"`
				Http_request *struct {
					Count            *int64  `tfsdk:"count" json:"count,omitempty"`
					Enable_conn_pool *bool   `tfsdk:"enable_conn_pool" json:"enable-conn-pool,omitempty"`
					Url              *string `tfsdk:"url" json:"url,omitempty"`
				} `tfsdk:"http_request" json:"http-request,omitempty"`
				Jvm_exception *struct {
					Class     *string `tfsdk:"class" json:"class,omitempty"`
					Exception *string `tfsdk:"exception" json:"exception,omitempty"`
					Method    *string `tfsdk:"method" json:"method,omitempty"`
					Pid       *int64  `tfsdk:"pid" json:"pid,omitempty"`
					Port      *int64  `tfsdk:"port" json:"port,omitempty"`
				} `tfsdk:"jvm_exception" json:"jvm-exception,omitempty"`
				Jvm_gc *struct {
					Pid  *int64 `tfsdk:"pid" json:"pid,omitempty"`
					Port *int64 `tfsdk:"port" json:"port,omitempty"`
				} `tfsdk:"jvm_gc" json:"jvm-gc,omitempty"`
				Jvm_latency *struct {
					Class   *string `tfsdk:"class" json:"class,omitempty"`
					Latency *int64  `tfsdk:"latency" json:"latency,omitempty"`
					Method  *string `tfsdk:"method" json:"method,omitempty"`
					Pid     *int64  `tfsdk:"pid" json:"pid,omitempty"`
					Port    *int64  `tfsdk:"port" json:"port,omitempty"`
				} `tfsdk:"jvm_latency" json:"jvm-latency,omitempty"`
				Jvm_mysql *struct {
					Database              *string `tfsdk:"database" json:"database,omitempty"`
					Exception             *string `tfsdk:"exception" json:"exception,omitempty"`
					Latency               *int64  `tfsdk:"latency" json:"latency,omitempty"`
					MysqlConnectorVersion *string `tfsdk:"mysql_connector_version" json:"mysqlConnectorVersion,omitempty"`
					Pid                   *int64  `tfsdk:"pid" json:"pid,omitempty"`
					Port                  *int64  `tfsdk:"port" json:"port,omitempty"`
					SqlType               *string `tfsdk:"sql_type" json:"sqlType,omitempty"`
					Table                 *string `tfsdk:"table" json:"table,omitempty"`
				} `tfsdk:"jvm_mysql" json:"jvm-mysql,omitempty"`
				Jvm_return *struct {
					Class  *string `tfsdk:"class" json:"class,omitempty"`
					Method *string `tfsdk:"method" json:"method,omitempty"`
					Pid    *int64  `tfsdk:"pid" json:"pid,omitempty"`
					Port   *int64  `tfsdk:"port" json:"port,omitempty"`
					Value  *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"jvm_return" json:"jvm-return,omitempty"`
				Jvm_rule_data *struct {
					Pid       *int64  `tfsdk:"pid" json:"pid,omitempty"`
					Port      *int64  `tfsdk:"port" json:"port,omitempty"`
					Rule_data *string `tfsdk:"rule_data" json:"rule-data,omitempty"`
				} `tfsdk:"jvm_rule_data" json:"jvm-rule-data,omitempty"`
				Jvm_stress *struct {
					Cpu_count *int64  `tfsdk:"cpu_count" json:"cpu-count,omitempty"`
					Mem_type  *string `tfsdk:"mem_type" json:"mem-type,omitempty"`
					Pid       *int64  `tfsdk:"pid" json:"pid,omitempty"`
					Port      *int64  `tfsdk:"port" json:"port,omitempty"`
				} `tfsdk:"jvm_stress" json:"jvm-stress,omitempty"`
				Kafka_fill *struct {
					Host          *string `tfsdk:"host" json:"host,omitempty"`
					MaxBytes      *int64  `tfsdk:"max_bytes" json:"maxBytes,omitempty"`
					MessageSize   *int64  `tfsdk:"message_size" json:"messageSize,omitempty"`
					Password      *string `tfsdk:"password" json:"password,omitempty"`
					Port          *int64  `tfsdk:"port" json:"port,omitempty"`
					ReloadCommand *string `tfsdk:"reload_command" json:"reloadCommand,omitempty"`
					Topic         *string `tfsdk:"topic" json:"topic,omitempty"`
					Username      *string `tfsdk:"username" json:"username,omitempty"`
				} `tfsdk:"kafka_fill" json:"kafka-fill,omitempty"`
				Kafka_flood *struct {
					Host        *string `tfsdk:"host" json:"host,omitempty"`
					MessageSize *int64  `tfsdk:"message_size" json:"messageSize,omitempty"`
					Password    *string `tfsdk:"password" json:"password,omitempty"`
					Port        *int64  `tfsdk:"port" json:"port,omitempty"`
					Threads     *int64  `tfsdk:"threads" json:"threads,omitempty"`
					Topic       *string `tfsdk:"topic" json:"topic,omitempty"`
					Username    *string `tfsdk:"username" json:"username,omitempty"`
				} `tfsdk:"kafka_flood" json:"kafka-flood,omitempty"`
				Kafka_io *struct {
					ConfigFile  *string `tfsdk:"config_file" json:"configFile,omitempty"`
					NonReadable *bool   `tfsdk:"non_readable" json:"nonReadable,omitempty"`
					NonWritable *bool   `tfsdk:"non_writable" json:"nonWritable,omitempty"`
					Topic       *string `tfsdk:"topic" json:"topic,omitempty"`
				} `tfsdk:"kafka_io" json:"kafka-io,omitempty"`
				Mode              *string `tfsdk:"mode" json:"mode,omitempty"`
				Network_bandwidth *struct {
					Buffer     *int64  `tfsdk:"buffer" json:"buffer,omitempty"`
					Device     *string `tfsdk:"device" json:"device,omitempty"`
					Hostname   *string `tfsdk:"hostname" json:"hostname,omitempty"`
					Ip_address *string `tfsdk:"ip_address" json:"ip-address,omitempty"`
					Limit      *int64  `tfsdk:"limit" json:"limit,omitempty"`
					Minburst   *int64  `tfsdk:"minburst" json:"minburst,omitempty"`
					Peakrate   *int64  `tfsdk:"peakrate" json:"peakrate,omitempty"`
					Rate       *string `tfsdk:"rate" json:"rate,omitempty"`
				} `tfsdk:"network_bandwidth" json:"network-bandwidth,omitempty"`
				Network_corrupt *struct {
					Correlation *string `tfsdk:"correlation" json:"correlation,omitempty"`
					Device      *string `tfsdk:"device" json:"device,omitempty"`
					Egress_port *string `tfsdk:"egress_port" json:"egress-port,omitempty"`
					Hostname    *string `tfsdk:"hostname" json:"hostname,omitempty"`
					Ip_address  *string `tfsdk:"ip_address" json:"ip-address,omitempty"`
					Ip_protocol *string `tfsdk:"ip_protocol" json:"ip-protocol,omitempty"`
					Percent     *string `tfsdk:"percent" json:"percent,omitempty"`
					Source_port *string `tfsdk:"source_port" json:"source-port,omitempty"`
				} `tfsdk:"network_corrupt" json:"network-corrupt,omitempty"`
				Network_delay *struct {
					Accept_tcp_flags *string `tfsdk:"accept_tcp_flags" json:"accept-tcp-flags,omitempty"`
					Correlation      *string `tfsdk:"correlation" json:"correlation,omitempty"`
					Device           *string `tfsdk:"device" json:"device,omitempty"`
					Egress_port      *string `tfsdk:"egress_port" json:"egress-port,omitempty"`
					Hostname         *string `tfsdk:"hostname" json:"hostname,omitempty"`
					Ip_address       *string `tfsdk:"ip_address" json:"ip-address,omitempty"`
					Ip_protocol      *string `tfsdk:"ip_protocol" json:"ip-protocol,omitempty"`
					Jitter           *string `tfsdk:"jitter" json:"jitter,omitempty"`
					Latency          *string `tfsdk:"latency" json:"latency,omitempty"`
					Source_port      *string `tfsdk:"source_port" json:"source-port,omitempty"`
				} `tfsdk:"network_delay" json:"network-delay,omitempty"`
				Network_dns *struct {
					Dns_domain_name *string `tfsdk:"dns_domain_name" json:"dns-domain-name,omitempty"`
					Dns_ip          *string `tfsdk:"dns_ip" json:"dns-ip,omitempty"`
					Dns_server      *string `tfsdk:"dns_server" json:"dns-server,omitempty"`
				} `tfsdk:"network_dns" json:"network-dns,omitempty"`
				Network_down *struct {
					Device   *string `tfsdk:"device" json:"device,omitempty"`
					Duration *string `tfsdk:"duration" json:"duration,omitempty"`
				} `tfsdk:"network_down" json:"network-down,omitempty"`
				Network_duplicate *struct {
					Correlation *string `tfsdk:"correlation" json:"correlation,omitempty"`
					Device      *string `tfsdk:"device" json:"device,omitempty"`
					Egress_port *string `tfsdk:"egress_port" json:"egress-port,omitempty"`
					Hostname    *string `tfsdk:"hostname" json:"hostname,omitempty"`
					Ip_address  *string `tfsdk:"ip_address" json:"ip-address,omitempty"`
					Ip_protocol *string `tfsdk:"ip_protocol" json:"ip-protocol,omitempty"`
					Percent     *string `tfsdk:"percent" json:"percent,omitempty"`
					Source_port *string `tfsdk:"source_port" json:"source-port,omitempty"`
				} `tfsdk:"network_duplicate" json:"network-duplicate,omitempty"`
				Network_flood *struct {
					Duration   *string `tfsdk:"duration" json:"duration,omitempty"`
					Ip_address *string `tfsdk:"ip_address" json:"ip-address,omitempty"`
					Parallel   *int64  `tfsdk:"parallel" json:"parallel,omitempty"`
					Port       *string `tfsdk:"port" json:"port,omitempty"`
					Rate       *string `tfsdk:"rate" json:"rate,omitempty"`
				} `tfsdk:"network_flood" json:"network-flood,omitempty"`
				Network_loss *struct {
					Correlation *string `tfsdk:"correlation" json:"correlation,omitempty"`
					Device      *string `tfsdk:"device" json:"device,omitempty"`
					Egress_port *string `tfsdk:"egress_port" json:"egress-port,omitempty"`
					Hostname    *string `tfsdk:"hostname" json:"hostname,omitempty"`
					Ip_address  *string `tfsdk:"ip_address" json:"ip-address,omitempty"`
					Ip_protocol *string `tfsdk:"ip_protocol" json:"ip-protocol,omitempty"`
					Percent     *string `tfsdk:"percent" json:"percent,omitempty"`
					Source_port *string `tfsdk:"source_port" json:"source-port,omitempty"`
				} `tfsdk:"network_loss" json:"network-loss,omitempty"`
				Network_partition *struct {
					Accept_tcp_flags *string `tfsdk:"accept_tcp_flags" json:"accept-tcp-flags,omitempty"`
					Device           *string `tfsdk:"device" json:"device,omitempty"`
					Direction        *string `tfsdk:"direction" json:"direction,omitempty"`
					Hostname         *string `tfsdk:"hostname" json:"hostname,omitempty"`
					Ip_address       *string `tfsdk:"ip_address" json:"ip-address,omitempty"`
					Ip_protocol      *string `tfsdk:"ip_protocol" json:"ip-protocol,omitempty"`
				} `tfsdk:"network_partition" json:"network-partition,omitempty"`
				Process *struct {
					Process    *string `tfsdk:"process" json:"process,omitempty"`
					RecoverCmd *string `tfsdk:"recover_cmd" json:"recoverCmd,omitempty"`
					Signal     *int64  `tfsdk:"signal" json:"signal,omitempty"`
				} `tfsdk:"process" json:"process,omitempty"`
				Redis_cacheLimit *struct {
					Addr      *string `tfsdk:"addr" json:"addr,omitempty"`
					CacheSize *string `tfsdk:"cache_size" json:"cacheSize,omitempty"`
					Password  *string `tfsdk:"password" json:"password,omitempty"`
					Percent   *string `tfsdk:"percent" json:"percent,omitempty"`
				} `tfsdk:"redis_cache_limit" json:"redis-cacheLimit,omitempty"`
				Redis_expiration *struct {
					Addr       *string `tfsdk:"addr" json:"addr,omitempty"`
					Expiration *string `tfsdk:"expiration" json:"expiration,omitempty"`
					Key        *string `tfsdk:"key" json:"key,omitempty"`
					Option     *string `tfsdk:"option" json:"option,omitempty"`
					Password   *string `tfsdk:"password" json:"password,omitempty"`
				} `tfsdk:"redis_expiration" json:"redis-expiration,omitempty"`
				Redis_penetration *struct {
					Addr       *string `tfsdk:"addr" json:"addr,omitempty"`
					Password   *string `tfsdk:"password" json:"password,omitempty"`
					RequestNum *int64  `tfsdk:"request_num" json:"requestNum,omitempty"`
				} `tfsdk:"redis_penetration" json:"redis-penetration,omitempty"`
				Redis_restart *struct {
					Addr        *string `tfsdk:"addr" json:"addr,omitempty"`
					Conf        *string `tfsdk:"conf" json:"conf,omitempty"`
					FlushConfig *bool   `tfsdk:"flush_config" json:"flushConfig,omitempty"`
					Password    *string `tfsdk:"password" json:"password,omitempty"`
					RedisPath   *bool   `tfsdk:"redis_path" json:"redisPath,omitempty"`
				} `tfsdk:"redis_restart" json:"redis-restart,omitempty"`
				Redis_stop *struct {
					Addr        *string `tfsdk:"addr" json:"addr,omitempty"`
					Conf        *string `tfsdk:"conf" json:"conf,omitempty"`
					FlushConfig *bool   `tfsdk:"flush_config" json:"flushConfig,omitempty"`
					Password    *string `tfsdk:"password" json:"password,omitempty"`
					RedisPath   *bool   `tfsdk:"redis_path" json:"redisPath,omitempty"`
				} `tfsdk:"redis_stop" json:"redis-stop,omitempty"`
				RemoteCluster *string `tfsdk:"remote_cluster" json:"remoteCluster,omitempty"`
				Selector      *struct {
					AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" json:"annotationSelectors,omitempty"`
					ExpressionSelectors *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"expression_selectors" json:"expressionSelectors,omitempty"`
					FieldSelectors   *map[string]string   `tfsdk:"field_selectors" json:"fieldSelectors,omitempty"`
					LabelSelectors   *map[string]string   `tfsdk:"label_selectors" json:"labelSelectors,omitempty"`
					Namespaces       *[]string            `tfsdk:"namespaces" json:"namespaces,omitempty"`
					PhysicalMachines *map[string][]string `tfsdk:"physical_machines" json:"physicalMachines,omitempty"`
				} `tfsdk:"selector" json:"selector,omitempty"`
				Stress_cpu *struct {
					Load    *int64    `tfsdk:"load" json:"load,omitempty"`
					Options *[]string `tfsdk:"options" json:"options,omitempty"`
					Workers *int64    `tfsdk:"workers" json:"workers,omitempty"`
				} `tfsdk:"stress_cpu" json:"stress-cpu,omitempty"`
				Stress_mem *struct {
					Options *[]string `tfsdk:"options" json:"options,omitempty"`
					Size    *string   `tfsdk:"size" json:"size,omitempty"`
				} `tfsdk:"stress_mem" json:"stress-mem,omitempty"`
				Uid          *string `tfsdk:"uid" json:"uid,omitempty"`
				User_defined *struct {
					AttackCmd  *string `tfsdk:"attack_cmd" json:"attackCmd,omitempty"`
					RecoverCmd *string `tfsdk:"recover_cmd" json:"recoverCmd,omitempty"`
				} `tfsdk:"user_defined" json:"user_defined,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
				Vm    *struct {
					Vm_name *string `tfsdk:"vm_name" json:"vm-name,omitempty"`
				} `tfsdk:"vm" json:"vm,omitempty"`
			} `tfsdk:"physicalmachine_chaos" json:"physicalmachineChaos,omitempty"`
			PodChaos *struct {
				Action         *string   `tfsdk:"action" json:"action,omitempty"`
				ContainerNames *[]string `tfsdk:"container_names" json:"containerNames,omitempty"`
				Duration       *string   `tfsdk:"duration" json:"duration,omitempty"`
				GracePeriod    *int64    `tfsdk:"grace_period" json:"gracePeriod,omitempty"`
				Mode           *string   `tfsdk:"mode" json:"mode,omitempty"`
				RemoteCluster  *string   `tfsdk:"remote_cluster" json:"remoteCluster,omitempty"`
				Selector       *struct {
					AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" json:"annotationSelectors,omitempty"`
					ExpressionSelectors *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"expression_selectors" json:"expressionSelectors,omitempty"`
					FieldSelectors    *map[string]string   `tfsdk:"field_selectors" json:"fieldSelectors,omitempty"`
					LabelSelectors    *map[string]string   `tfsdk:"label_selectors" json:"labelSelectors,omitempty"`
					Namespaces        *[]string            `tfsdk:"namespaces" json:"namespaces,omitempty"`
					NodeSelectors     *map[string]string   `tfsdk:"node_selectors" json:"nodeSelectors,omitempty"`
					Nodes             *[]string            `tfsdk:"nodes" json:"nodes,omitempty"`
					PodPhaseSelectors *[]string            `tfsdk:"pod_phase_selectors" json:"podPhaseSelectors,omitempty"`
					Pods              *map[string][]string `tfsdk:"pods" json:"pods,omitempty"`
				} `tfsdk:"selector" json:"selector,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"pod_chaos" json:"podChaos,omitempty"`
			Schedule *struct {
				AwsChaos *struct {
					Action        *string `tfsdk:"action" json:"action,omitempty"`
					AwsRegion     *string `tfsdk:"aws_region" json:"awsRegion,omitempty"`
					DeviceName    *string `tfsdk:"device_name" json:"deviceName,omitempty"`
					Duration      *string `tfsdk:"duration" json:"duration,omitempty"`
					Ec2Instance   *string `tfsdk:"ec2_instance" json:"ec2Instance,omitempty"`
					Endpoint      *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
					RemoteCluster *string `tfsdk:"remote_cluster" json:"remoteCluster,omitempty"`
					SecretName    *string `tfsdk:"secret_name" json:"secretName,omitempty"`
					VolumeID      *string `tfsdk:"volume_id" json:"volumeID,omitempty"`
				} `tfsdk:"aws_chaos" json:"awsChaos,omitempty"`
				AzureChaos *struct {
					Action            *string `tfsdk:"action" json:"action,omitempty"`
					DiskName          *string `tfsdk:"disk_name" json:"diskName,omitempty"`
					Duration          *string `tfsdk:"duration" json:"duration,omitempty"`
					Lun               *int64  `tfsdk:"lun" json:"lun,omitempty"`
					RemoteCluster     *string `tfsdk:"remote_cluster" json:"remoteCluster,omitempty"`
					ResourceGroupName *string `tfsdk:"resource_group_name" json:"resourceGroupName,omitempty"`
					SecretName        *string `tfsdk:"secret_name" json:"secretName,omitempty"`
					SubscriptionID    *string `tfsdk:"subscription_id" json:"subscriptionID,omitempty"`
					VmName            *string `tfsdk:"vm_name" json:"vmName,omitempty"`
				} `tfsdk:"azure_chaos" json:"azureChaos,omitempty"`
				BlockChaos *struct {
					Action         *string   `tfsdk:"action" json:"action,omitempty"`
					ContainerNames *[]string `tfsdk:"container_names" json:"containerNames,omitempty"`
					Delay          *struct {
						Correlation *string `tfsdk:"correlation" json:"correlation,omitempty"`
						Jitter      *string `tfsdk:"jitter" json:"jitter,omitempty"`
						Latency     *string `tfsdk:"latency" json:"latency,omitempty"`
					} `tfsdk:"delay" json:"delay,omitempty"`
					Duration      *string `tfsdk:"duration" json:"duration,omitempty"`
					Mode          *string `tfsdk:"mode" json:"mode,omitempty"`
					RemoteCluster *string `tfsdk:"remote_cluster" json:"remoteCluster,omitempty"`
					Selector      *struct {
						AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" json:"annotationSelectors,omitempty"`
						ExpressionSelectors *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"expression_selectors" json:"expressionSelectors,omitempty"`
						FieldSelectors    *map[string]string   `tfsdk:"field_selectors" json:"fieldSelectors,omitempty"`
						LabelSelectors    *map[string]string   `tfsdk:"label_selectors" json:"labelSelectors,omitempty"`
						Namespaces        *[]string            `tfsdk:"namespaces" json:"namespaces,omitempty"`
						NodeSelectors     *map[string]string   `tfsdk:"node_selectors" json:"nodeSelectors,omitempty"`
						Nodes             *[]string            `tfsdk:"nodes" json:"nodes,omitempty"`
						PodPhaseSelectors *[]string            `tfsdk:"pod_phase_selectors" json:"podPhaseSelectors,omitempty"`
						Pods              *map[string][]string `tfsdk:"pods" json:"pods,omitempty"`
					} `tfsdk:"selector" json:"selector,omitempty"`
					Value      *string `tfsdk:"value" json:"value,omitempty"`
					VolumeName *string `tfsdk:"volume_name" json:"volumeName,omitempty"`
				} `tfsdk:"block_chaos" json:"blockChaos,omitempty"`
				ConcurrencyPolicy *string `tfsdk:"concurrency_policy" json:"concurrencyPolicy,omitempty"`
				DnsChaos          *struct {
					Action         *string   `tfsdk:"action" json:"action,omitempty"`
					ContainerNames *[]string `tfsdk:"container_names" json:"containerNames,omitempty"`
					Duration       *string   `tfsdk:"duration" json:"duration,omitempty"`
					Mode           *string   `tfsdk:"mode" json:"mode,omitempty"`
					Patterns       *[]string `tfsdk:"patterns" json:"patterns,omitempty"`
					RemoteCluster  *string   `tfsdk:"remote_cluster" json:"remoteCluster,omitempty"`
					Selector       *struct {
						AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" json:"annotationSelectors,omitempty"`
						ExpressionSelectors *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"expression_selectors" json:"expressionSelectors,omitempty"`
						FieldSelectors    *map[string]string   `tfsdk:"field_selectors" json:"fieldSelectors,omitempty"`
						LabelSelectors    *map[string]string   `tfsdk:"label_selectors" json:"labelSelectors,omitempty"`
						Namespaces        *[]string            `tfsdk:"namespaces" json:"namespaces,omitempty"`
						NodeSelectors     *map[string]string   `tfsdk:"node_selectors" json:"nodeSelectors,omitempty"`
						Nodes             *[]string            `tfsdk:"nodes" json:"nodes,omitempty"`
						PodPhaseSelectors *[]string            `tfsdk:"pod_phase_selectors" json:"podPhaseSelectors,omitempty"`
						Pods              *map[string][]string `tfsdk:"pods" json:"pods,omitempty"`
					} `tfsdk:"selector" json:"selector,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"dns_chaos" json:"dnsChaos,omitempty"`
				GcpChaos *struct {
					Action        *string   `tfsdk:"action" json:"action,omitempty"`
					DeviceNames   *[]string `tfsdk:"device_names" json:"deviceNames,omitempty"`
					Duration      *string   `tfsdk:"duration" json:"duration,omitempty"`
					Instance      *string   `tfsdk:"instance" json:"instance,omitempty"`
					Project       *string   `tfsdk:"project" json:"project,omitempty"`
					RemoteCluster *string   `tfsdk:"remote_cluster" json:"remoteCluster,omitempty"`
					SecretName    *string   `tfsdk:"secret_name" json:"secretName,omitempty"`
					Zone          *string   `tfsdk:"zone" json:"zone,omitempty"`
				} `tfsdk:"gcp_chaos" json:"gcpChaos,omitempty"`
				HistoryLimit *int64 `tfsdk:"history_limit" json:"historyLimit,omitempty"`
				HttpChaos    *struct {
					Abort    *bool   `tfsdk:"abort" json:"abort,omitempty"`
					Code     *int64  `tfsdk:"code" json:"code,omitempty"`
					Delay    *string `tfsdk:"delay" json:"delay,omitempty"`
					Duration *string `tfsdk:"duration" json:"duration,omitempty"`
					Method   *string `tfsdk:"method" json:"method,omitempty"`
					Mode     *string `tfsdk:"mode" json:"mode,omitempty"`
					Patch    *struct {
						Body *struct {
							Type  *string `tfsdk:"type" json:"type,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"body" json:"body,omitempty"`
						Headers *[]string `tfsdk:"headers" json:"headers,omitempty"`
						Queries *[]string `tfsdk:"queries" json:"queries,omitempty"`
					} `tfsdk:"patch" json:"patch,omitempty"`
					Path          *string `tfsdk:"path" json:"path,omitempty"`
					Port          *int64  `tfsdk:"port" json:"port,omitempty"`
					RemoteCluster *string `tfsdk:"remote_cluster" json:"remoteCluster,omitempty"`
					Replace       *struct {
						Body    *string            `tfsdk:"body" json:"body,omitempty"`
						Code    *int64             `tfsdk:"code" json:"code,omitempty"`
						Headers *map[string]string `tfsdk:"headers" json:"headers,omitempty"`
						Method  *string            `tfsdk:"method" json:"method,omitempty"`
						Path    *string            `tfsdk:"path" json:"path,omitempty"`
						Queries *map[string]string `tfsdk:"queries" json:"queries,omitempty"`
					} `tfsdk:"replace" json:"replace,omitempty"`
					Request_headers  *map[string]string `tfsdk:"request_headers" json:"request_headers,omitempty"`
					Response_headers *map[string]string `tfsdk:"response_headers" json:"response_headers,omitempty"`
					Selector         *struct {
						AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" json:"annotationSelectors,omitempty"`
						ExpressionSelectors *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"expression_selectors" json:"expressionSelectors,omitempty"`
						FieldSelectors    *map[string]string   `tfsdk:"field_selectors" json:"fieldSelectors,omitempty"`
						LabelSelectors    *map[string]string   `tfsdk:"label_selectors" json:"labelSelectors,omitempty"`
						Namespaces        *[]string            `tfsdk:"namespaces" json:"namespaces,omitempty"`
						NodeSelectors     *map[string]string   `tfsdk:"node_selectors" json:"nodeSelectors,omitempty"`
						Nodes             *[]string            `tfsdk:"nodes" json:"nodes,omitempty"`
						PodPhaseSelectors *[]string            `tfsdk:"pod_phase_selectors" json:"podPhaseSelectors,omitempty"`
						Pods              *map[string][]string `tfsdk:"pods" json:"pods,omitempty"`
					} `tfsdk:"selector" json:"selector,omitempty"`
					Target *string `tfsdk:"target" json:"target,omitempty"`
					Tls    *struct {
						CaName          *string `tfsdk:"ca_name" json:"caName,omitempty"`
						CertName        *string `tfsdk:"cert_name" json:"certName,omitempty"`
						KeyName         *string `tfsdk:"key_name" json:"keyName,omitempty"`
						SecretName      *string `tfsdk:"secret_name" json:"secretName,omitempty"`
						SecretNamespace *string `tfsdk:"secret_namespace" json:"secretNamespace,omitempty"`
					} `tfsdk:"tls" json:"tls,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"http_chaos" json:"httpChaos,omitempty"`
				IoChaos *struct {
					Action *string `tfsdk:"action" json:"action,omitempty"`
					Attr   *struct {
						Atime *struct {
							Nsec *int64 `tfsdk:"nsec" json:"nsec,omitempty"`
							Sec  *int64 `tfsdk:"sec" json:"sec,omitempty"`
						} `tfsdk:"atime" json:"atime,omitempty"`
						Blocks *int64 `tfsdk:"blocks" json:"blocks,omitempty"`
						Ctime  *struct {
							Nsec *int64 `tfsdk:"nsec" json:"nsec,omitempty"`
							Sec  *int64 `tfsdk:"sec" json:"sec,omitempty"`
						} `tfsdk:"ctime" json:"ctime,omitempty"`
						Gid   *int64  `tfsdk:"gid" json:"gid,omitempty"`
						Ino   *int64  `tfsdk:"ino" json:"ino,omitempty"`
						Kind  *string `tfsdk:"kind" json:"kind,omitempty"`
						Mtime *struct {
							Nsec *int64 `tfsdk:"nsec" json:"nsec,omitempty"`
							Sec  *int64 `tfsdk:"sec" json:"sec,omitempty"`
						} `tfsdk:"mtime" json:"mtime,omitempty"`
						Nlink *int64 `tfsdk:"nlink" json:"nlink,omitempty"`
						Perm  *int64 `tfsdk:"perm" json:"perm,omitempty"`
						Rdev  *int64 `tfsdk:"rdev" json:"rdev,omitempty"`
						Size  *int64 `tfsdk:"size" json:"size,omitempty"`
						Uid   *int64 `tfsdk:"uid" json:"uid,omitempty"`
					} `tfsdk:"attr" json:"attr,omitempty"`
					ContainerNames *[]string `tfsdk:"container_names" json:"containerNames,omitempty"`
					Delay          *string   `tfsdk:"delay" json:"delay,omitempty"`
					Duration       *string   `tfsdk:"duration" json:"duration,omitempty"`
					Errno          *int64    `tfsdk:"errno" json:"errno,omitempty"`
					Methods        *[]string `tfsdk:"methods" json:"methods,omitempty"`
					Mistake        *struct {
						Filling        *string `tfsdk:"filling" json:"filling,omitempty"`
						MaxLength      *int64  `tfsdk:"max_length" json:"maxLength,omitempty"`
						MaxOccurrences *int64  `tfsdk:"max_occurrences" json:"maxOccurrences,omitempty"`
					} `tfsdk:"mistake" json:"mistake,omitempty"`
					Mode          *string `tfsdk:"mode" json:"mode,omitempty"`
					Path          *string `tfsdk:"path" json:"path,omitempty"`
					Percent       *int64  `tfsdk:"percent" json:"percent,omitempty"`
					RemoteCluster *string `tfsdk:"remote_cluster" json:"remoteCluster,omitempty"`
					Selector      *struct {
						AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" json:"annotationSelectors,omitempty"`
						ExpressionSelectors *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"expression_selectors" json:"expressionSelectors,omitempty"`
						FieldSelectors    *map[string]string   `tfsdk:"field_selectors" json:"fieldSelectors,omitempty"`
						LabelSelectors    *map[string]string   `tfsdk:"label_selectors" json:"labelSelectors,omitempty"`
						Namespaces        *[]string            `tfsdk:"namespaces" json:"namespaces,omitempty"`
						NodeSelectors     *map[string]string   `tfsdk:"node_selectors" json:"nodeSelectors,omitempty"`
						Nodes             *[]string            `tfsdk:"nodes" json:"nodes,omitempty"`
						PodPhaseSelectors *[]string            `tfsdk:"pod_phase_selectors" json:"podPhaseSelectors,omitempty"`
						Pods              *map[string][]string `tfsdk:"pods" json:"pods,omitempty"`
					} `tfsdk:"selector" json:"selector,omitempty"`
					Value      *string `tfsdk:"value" json:"value,omitempty"`
					VolumePath *string `tfsdk:"volume_path" json:"volumePath,omitempty"`
				} `tfsdk:"io_chaos" json:"ioChaos,omitempty"`
				JvmChaos *struct {
					Action                *string   `tfsdk:"action" json:"action,omitempty"`
					Class                 *string   `tfsdk:"class" json:"class,omitempty"`
					ContainerNames        *[]string `tfsdk:"container_names" json:"containerNames,omitempty"`
					CpuCount              *int64    `tfsdk:"cpu_count" json:"cpuCount,omitempty"`
					Database              *string   `tfsdk:"database" json:"database,omitempty"`
					Duration              *string   `tfsdk:"duration" json:"duration,omitempty"`
					Exception             *string   `tfsdk:"exception" json:"exception,omitempty"`
					Latency               *int64    `tfsdk:"latency" json:"latency,omitempty"`
					MemType               *string   `tfsdk:"mem_type" json:"memType,omitempty"`
					Method                *string   `tfsdk:"method" json:"method,omitempty"`
					Mode                  *string   `tfsdk:"mode" json:"mode,omitempty"`
					MysqlConnectorVersion *string   `tfsdk:"mysql_connector_version" json:"mysqlConnectorVersion,omitempty"`
					Name                  *string   `tfsdk:"name" json:"name,omitempty"`
					Pid                   *int64    `tfsdk:"pid" json:"pid,omitempty"`
					Port                  *int64    `tfsdk:"port" json:"port,omitempty"`
					RemoteCluster         *string   `tfsdk:"remote_cluster" json:"remoteCluster,omitempty"`
					RuleData              *string   `tfsdk:"rule_data" json:"ruleData,omitempty"`
					Selector              *struct {
						AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" json:"annotationSelectors,omitempty"`
						ExpressionSelectors *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"expression_selectors" json:"expressionSelectors,omitempty"`
						FieldSelectors    *map[string]string   `tfsdk:"field_selectors" json:"fieldSelectors,omitempty"`
						LabelSelectors    *map[string]string   `tfsdk:"label_selectors" json:"labelSelectors,omitempty"`
						Namespaces        *[]string            `tfsdk:"namespaces" json:"namespaces,omitempty"`
						NodeSelectors     *map[string]string   `tfsdk:"node_selectors" json:"nodeSelectors,omitempty"`
						Nodes             *[]string            `tfsdk:"nodes" json:"nodes,omitempty"`
						PodPhaseSelectors *[]string            `tfsdk:"pod_phase_selectors" json:"podPhaseSelectors,omitempty"`
						Pods              *map[string][]string `tfsdk:"pods" json:"pods,omitempty"`
					} `tfsdk:"selector" json:"selector,omitempty"`
					SqlType *string `tfsdk:"sql_type" json:"sqlType,omitempty"`
					Table   *string `tfsdk:"table" json:"table,omitempty"`
					Value   *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"jvm_chaos" json:"jvmChaos,omitempty"`
				KernelChaos *struct {
					ContainerNames  *[]string `tfsdk:"container_names" json:"containerNames,omitempty"`
					Duration        *string   `tfsdk:"duration" json:"duration,omitempty"`
					FailKernRequest *struct {
						Callchain *[]struct {
							Funcname   *string `tfsdk:"funcname" json:"funcname,omitempty"`
							Parameters *string `tfsdk:"parameters" json:"parameters,omitempty"`
							Predicate  *string `tfsdk:"predicate" json:"predicate,omitempty"`
						} `tfsdk:"callchain" json:"callchain,omitempty"`
						Failtype    *int64    `tfsdk:"failtype" json:"failtype,omitempty"`
						Headers     *[]string `tfsdk:"headers" json:"headers,omitempty"`
						Probability *int64    `tfsdk:"probability" json:"probability,omitempty"`
						Times       *int64    `tfsdk:"times" json:"times,omitempty"`
					} `tfsdk:"fail_kern_request" json:"failKernRequest,omitempty"`
					Mode          *string `tfsdk:"mode" json:"mode,omitempty"`
					RemoteCluster *string `tfsdk:"remote_cluster" json:"remoteCluster,omitempty"`
					Selector      *struct {
						AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" json:"annotationSelectors,omitempty"`
						ExpressionSelectors *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"expression_selectors" json:"expressionSelectors,omitempty"`
						FieldSelectors    *map[string]string   `tfsdk:"field_selectors" json:"fieldSelectors,omitempty"`
						LabelSelectors    *map[string]string   `tfsdk:"label_selectors" json:"labelSelectors,omitempty"`
						Namespaces        *[]string            `tfsdk:"namespaces" json:"namespaces,omitempty"`
						NodeSelectors     *map[string]string   `tfsdk:"node_selectors" json:"nodeSelectors,omitempty"`
						Nodes             *[]string            `tfsdk:"nodes" json:"nodes,omitempty"`
						PodPhaseSelectors *[]string            `tfsdk:"pod_phase_selectors" json:"podPhaseSelectors,omitempty"`
						Pods              *map[string][]string `tfsdk:"pods" json:"pods,omitempty"`
					} `tfsdk:"selector" json:"selector,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"kernel_chaos" json:"kernelChaos,omitempty"`
				NetworkChaos *struct {
					Action    *string `tfsdk:"action" json:"action,omitempty"`
					Bandwidth *struct {
						Buffer   *int64  `tfsdk:"buffer" json:"buffer,omitempty"`
						Limit    *int64  `tfsdk:"limit" json:"limit,omitempty"`
						Minburst *int64  `tfsdk:"minburst" json:"minburst,omitempty"`
						Peakrate *int64  `tfsdk:"peakrate" json:"peakrate,omitempty"`
						Rate     *string `tfsdk:"rate" json:"rate,omitempty"`
					} `tfsdk:"bandwidth" json:"bandwidth,omitempty"`
					Corrupt *struct {
						Correlation *string `tfsdk:"correlation" json:"correlation,omitempty"`
						Corrupt     *string `tfsdk:"corrupt" json:"corrupt,omitempty"`
					} `tfsdk:"corrupt" json:"corrupt,omitempty"`
					Delay *struct {
						Correlation *string `tfsdk:"correlation" json:"correlation,omitempty"`
						Jitter      *string `tfsdk:"jitter" json:"jitter,omitempty"`
						Latency     *string `tfsdk:"latency" json:"latency,omitempty"`
						Reorder     *struct {
							Correlation *string `tfsdk:"correlation" json:"correlation,omitempty"`
							Gap         *int64  `tfsdk:"gap" json:"gap,omitempty"`
							Reorder     *string `tfsdk:"reorder" json:"reorder,omitempty"`
						} `tfsdk:"reorder" json:"reorder,omitempty"`
					} `tfsdk:"delay" json:"delay,omitempty"`
					Device    *string `tfsdk:"device" json:"device,omitempty"`
					Direction *string `tfsdk:"direction" json:"direction,omitempty"`
					Duplicate *struct {
						Correlation *string `tfsdk:"correlation" json:"correlation,omitempty"`
						Duplicate   *string `tfsdk:"duplicate" json:"duplicate,omitempty"`
					} `tfsdk:"duplicate" json:"duplicate,omitempty"`
					Duration        *string   `tfsdk:"duration" json:"duration,omitempty"`
					ExternalTargets *[]string `tfsdk:"external_targets" json:"externalTargets,omitempty"`
					Loss            *struct {
						Correlation *string `tfsdk:"correlation" json:"correlation,omitempty"`
						Loss        *string `tfsdk:"loss" json:"loss,omitempty"`
					} `tfsdk:"loss" json:"loss,omitempty"`
					Mode          *string `tfsdk:"mode" json:"mode,omitempty"`
					RemoteCluster *string `tfsdk:"remote_cluster" json:"remoteCluster,omitempty"`
					Selector      *struct {
						AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" json:"annotationSelectors,omitempty"`
						ExpressionSelectors *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"expression_selectors" json:"expressionSelectors,omitempty"`
						FieldSelectors    *map[string]string   `tfsdk:"field_selectors" json:"fieldSelectors,omitempty"`
						LabelSelectors    *map[string]string   `tfsdk:"label_selectors" json:"labelSelectors,omitempty"`
						Namespaces        *[]string            `tfsdk:"namespaces" json:"namespaces,omitempty"`
						NodeSelectors     *map[string]string   `tfsdk:"node_selectors" json:"nodeSelectors,omitempty"`
						Nodes             *[]string            `tfsdk:"nodes" json:"nodes,omitempty"`
						PodPhaseSelectors *[]string            `tfsdk:"pod_phase_selectors" json:"podPhaseSelectors,omitempty"`
						Pods              *map[string][]string `tfsdk:"pods" json:"pods,omitempty"`
					} `tfsdk:"selector" json:"selector,omitempty"`
					Target *struct {
						Mode     *string `tfsdk:"mode" json:"mode,omitempty"`
						Selector *struct {
							AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" json:"annotationSelectors,omitempty"`
							ExpressionSelectors *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"expression_selectors" json:"expressionSelectors,omitempty"`
							FieldSelectors    *map[string]string   `tfsdk:"field_selectors" json:"fieldSelectors,omitempty"`
							LabelSelectors    *map[string]string   `tfsdk:"label_selectors" json:"labelSelectors,omitempty"`
							Namespaces        *[]string            `tfsdk:"namespaces" json:"namespaces,omitempty"`
							NodeSelectors     *map[string]string   `tfsdk:"node_selectors" json:"nodeSelectors,omitempty"`
							Nodes             *[]string            `tfsdk:"nodes" json:"nodes,omitempty"`
							PodPhaseSelectors *[]string            `tfsdk:"pod_phase_selectors" json:"podPhaseSelectors,omitempty"`
							Pods              *map[string][]string `tfsdk:"pods" json:"pods,omitempty"`
						} `tfsdk:"selector" json:"selector,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"target" json:"target,omitempty"`
					TargetDevice *string `tfsdk:"target_device" json:"targetDevice,omitempty"`
					Value        *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"network_chaos" json:"networkChaos,omitempty"`
				PhysicalmachineChaos *struct {
					Action  *string   `tfsdk:"action" json:"action,omitempty"`
					Address *[]string `tfsdk:"address" json:"address,omitempty"`
					Clock   *struct {
						Clock_ids_slice *string `tfsdk:"clock_ids_slice" json:"clock-ids-slice,omitempty"`
						Pid             *int64  `tfsdk:"pid" json:"pid,omitempty"`
						Time_offset     *string `tfsdk:"time_offset" json:"time-offset,omitempty"`
					} `tfsdk:"clock" json:"clock,omitempty"`
					Disk_fill *struct {
						Fill_by_fallocate *bool   `tfsdk:"fill_by_fallocate" json:"fill-by-fallocate,omitempty"`
						Path              *string `tfsdk:"path" json:"path,omitempty"`
						Size              *string `tfsdk:"size" json:"size,omitempty"`
					} `tfsdk:"disk_fill" json:"disk-fill,omitempty"`
					Disk_read_payload *struct {
						Path                *string `tfsdk:"path" json:"path,omitempty"`
						Payload_process_num *int64  `tfsdk:"payload_process_num" json:"payload-process-num,omitempty"`
						Size                *string `tfsdk:"size" json:"size,omitempty"`
					} `tfsdk:"disk_read_payload" json:"disk-read-payload,omitempty"`
					Disk_write_payload *struct {
						Path                *string `tfsdk:"path" json:"path,omitempty"`
						Payload_process_num *int64  `tfsdk:"payload_process_num" json:"payload-process-num,omitempty"`
						Size                *string `tfsdk:"size" json:"size,omitempty"`
					} `tfsdk:"disk_write_payload" json:"disk-write-payload,omitempty"`
					Duration    *string `tfsdk:"duration" json:"duration,omitempty"`
					File_append *struct {
						Count     *int64  `tfsdk:"count" json:"count,omitempty"`
						Data      *string `tfsdk:"data" json:"data,omitempty"`
						File_name *string `tfsdk:"file_name" json:"file-name,omitempty"`
					} `tfsdk:"file_append" json:"file-append,omitempty"`
					File_create *struct {
						Dir_name  *string `tfsdk:"dir_name" json:"dir-name,omitempty"`
						File_name *string `tfsdk:"file_name" json:"file-name,omitempty"`
					} `tfsdk:"file_create" json:"file-create,omitempty"`
					File_delete *struct {
						Dir_name  *string `tfsdk:"dir_name" json:"dir-name,omitempty"`
						File_name *string `tfsdk:"file_name" json:"file-name,omitempty"`
					} `tfsdk:"file_delete" json:"file-delete,omitempty"`
					File_modify *struct {
						File_name *string `tfsdk:"file_name" json:"file-name,omitempty"`
						Privilege *int64  `tfsdk:"privilege" json:"privilege,omitempty"`
					} `tfsdk:"file_modify" json:"file-modify,omitempty"`
					File_rename *struct {
						Dest_file   *string `tfsdk:"dest_file" json:"dest-file,omitempty"`
						Source_file *string `tfsdk:"source_file" json:"source-file,omitempty"`
					} `tfsdk:"file_rename" json:"file-rename,omitempty"`
					File_replace *struct {
						Dest_string   *string `tfsdk:"dest_string" json:"dest-string,omitempty"`
						File_name     *string `tfsdk:"file_name" json:"file-name,omitempty"`
						Line          *int64  `tfsdk:"line" json:"line,omitempty"`
						Origin_string *string `tfsdk:"origin_string" json:"origin-string,omitempty"`
					} `tfsdk:"file_replace" json:"file-replace,omitempty"`
					Http_abort *struct {
						Code        *string   `tfsdk:"code" json:"code,omitempty"`
						Method      *string   `tfsdk:"method" json:"method,omitempty"`
						Path        *string   `tfsdk:"path" json:"path,omitempty"`
						Port        *int64    `tfsdk:"port" json:"port,omitempty"`
						Proxy_ports *[]string `tfsdk:"proxy_ports" json:"proxy_ports,omitempty"`
						Target      *string   `tfsdk:"target" json:"target,omitempty"`
					} `tfsdk:"http_abort" json:"http-abort,omitempty"`
					Http_config *struct {
						File_path *string `tfsdk:"file_path" json:"file_path,omitempty"`
					} `tfsdk:"http_config" json:"http-config,omitempty"`
					Http_delay *struct {
						Code        *string   `tfsdk:"code" json:"code,omitempty"`
						Delay       *string   `tfsdk:"delay" json:"delay,omitempty"`
						Method      *string   `tfsdk:"method" json:"method,omitempty"`
						Path        *string   `tfsdk:"path" json:"path,omitempty"`
						Port        *int64    `tfsdk:"port" json:"port,omitempty"`
						Proxy_ports *[]string `tfsdk:"proxy_ports" json:"proxy_ports,omitempty"`
						Target      *string   `tfsdk:"target" json:"target,omitempty"`
					} `tfsdk:"http_delay" json:"http-delay,omitempty"`
					Http_request *struct {
						Count            *int64  `tfsdk:"count" json:"count,omitempty"`
						Enable_conn_pool *bool   `tfsdk:"enable_conn_pool" json:"enable-conn-pool,omitempty"`
						Url              *string `tfsdk:"url" json:"url,omitempty"`
					} `tfsdk:"http_request" json:"http-request,omitempty"`
					Jvm_exception *struct {
						Class     *string `tfsdk:"class" json:"class,omitempty"`
						Exception *string `tfsdk:"exception" json:"exception,omitempty"`
						Method    *string `tfsdk:"method" json:"method,omitempty"`
						Pid       *int64  `tfsdk:"pid" json:"pid,omitempty"`
						Port      *int64  `tfsdk:"port" json:"port,omitempty"`
					} `tfsdk:"jvm_exception" json:"jvm-exception,omitempty"`
					Jvm_gc *struct {
						Pid  *int64 `tfsdk:"pid" json:"pid,omitempty"`
						Port *int64 `tfsdk:"port" json:"port,omitempty"`
					} `tfsdk:"jvm_gc" json:"jvm-gc,omitempty"`
					Jvm_latency *struct {
						Class   *string `tfsdk:"class" json:"class,omitempty"`
						Latency *int64  `tfsdk:"latency" json:"latency,omitempty"`
						Method  *string `tfsdk:"method" json:"method,omitempty"`
						Pid     *int64  `tfsdk:"pid" json:"pid,omitempty"`
						Port    *int64  `tfsdk:"port" json:"port,omitempty"`
					} `tfsdk:"jvm_latency" json:"jvm-latency,omitempty"`
					Jvm_mysql *struct {
						Database              *string `tfsdk:"database" json:"database,omitempty"`
						Exception             *string `tfsdk:"exception" json:"exception,omitempty"`
						Latency               *int64  `tfsdk:"latency" json:"latency,omitempty"`
						MysqlConnectorVersion *string `tfsdk:"mysql_connector_version" json:"mysqlConnectorVersion,omitempty"`
						Pid                   *int64  `tfsdk:"pid" json:"pid,omitempty"`
						Port                  *int64  `tfsdk:"port" json:"port,omitempty"`
						SqlType               *string `tfsdk:"sql_type" json:"sqlType,omitempty"`
						Table                 *string `tfsdk:"table" json:"table,omitempty"`
					} `tfsdk:"jvm_mysql" json:"jvm-mysql,omitempty"`
					Jvm_return *struct {
						Class  *string `tfsdk:"class" json:"class,omitempty"`
						Method *string `tfsdk:"method" json:"method,omitempty"`
						Pid    *int64  `tfsdk:"pid" json:"pid,omitempty"`
						Port   *int64  `tfsdk:"port" json:"port,omitempty"`
						Value  *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"jvm_return" json:"jvm-return,omitempty"`
					Jvm_rule_data *struct {
						Pid       *int64  `tfsdk:"pid" json:"pid,omitempty"`
						Port      *int64  `tfsdk:"port" json:"port,omitempty"`
						Rule_data *string `tfsdk:"rule_data" json:"rule-data,omitempty"`
					} `tfsdk:"jvm_rule_data" json:"jvm-rule-data,omitempty"`
					Jvm_stress *struct {
						Cpu_count *int64  `tfsdk:"cpu_count" json:"cpu-count,omitempty"`
						Mem_type  *string `tfsdk:"mem_type" json:"mem-type,omitempty"`
						Pid       *int64  `tfsdk:"pid" json:"pid,omitempty"`
						Port      *int64  `tfsdk:"port" json:"port,omitempty"`
					} `tfsdk:"jvm_stress" json:"jvm-stress,omitempty"`
					Kafka_fill *struct {
						Host          *string `tfsdk:"host" json:"host,omitempty"`
						MaxBytes      *int64  `tfsdk:"max_bytes" json:"maxBytes,omitempty"`
						MessageSize   *int64  `tfsdk:"message_size" json:"messageSize,omitempty"`
						Password      *string `tfsdk:"password" json:"password,omitempty"`
						Port          *int64  `tfsdk:"port" json:"port,omitempty"`
						ReloadCommand *string `tfsdk:"reload_command" json:"reloadCommand,omitempty"`
						Topic         *string `tfsdk:"topic" json:"topic,omitempty"`
						Username      *string `tfsdk:"username" json:"username,omitempty"`
					} `tfsdk:"kafka_fill" json:"kafka-fill,omitempty"`
					Kafka_flood *struct {
						Host        *string `tfsdk:"host" json:"host,omitempty"`
						MessageSize *int64  `tfsdk:"message_size" json:"messageSize,omitempty"`
						Password    *string `tfsdk:"password" json:"password,omitempty"`
						Port        *int64  `tfsdk:"port" json:"port,omitempty"`
						Threads     *int64  `tfsdk:"threads" json:"threads,omitempty"`
						Topic       *string `tfsdk:"topic" json:"topic,omitempty"`
						Username    *string `tfsdk:"username" json:"username,omitempty"`
					} `tfsdk:"kafka_flood" json:"kafka-flood,omitempty"`
					Kafka_io *struct {
						ConfigFile  *string `tfsdk:"config_file" json:"configFile,omitempty"`
						NonReadable *bool   `tfsdk:"non_readable" json:"nonReadable,omitempty"`
						NonWritable *bool   `tfsdk:"non_writable" json:"nonWritable,omitempty"`
						Topic       *string `tfsdk:"topic" json:"topic,omitempty"`
					} `tfsdk:"kafka_io" json:"kafka-io,omitempty"`
					Mode              *string `tfsdk:"mode" json:"mode,omitempty"`
					Network_bandwidth *struct {
						Buffer     *int64  `tfsdk:"buffer" json:"buffer,omitempty"`
						Device     *string `tfsdk:"device" json:"device,omitempty"`
						Hostname   *string `tfsdk:"hostname" json:"hostname,omitempty"`
						Ip_address *string `tfsdk:"ip_address" json:"ip-address,omitempty"`
						Limit      *int64  `tfsdk:"limit" json:"limit,omitempty"`
						Minburst   *int64  `tfsdk:"minburst" json:"minburst,omitempty"`
						Peakrate   *int64  `tfsdk:"peakrate" json:"peakrate,omitempty"`
						Rate       *string `tfsdk:"rate" json:"rate,omitempty"`
					} `tfsdk:"network_bandwidth" json:"network-bandwidth,omitempty"`
					Network_corrupt *struct {
						Correlation *string `tfsdk:"correlation" json:"correlation,omitempty"`
						Device      *string `tfsdk:"device" json:"device,omitempty"`
						Egress_port *string `tfsdk:"egress_port" json:"egress-port,omitempty"`
						Hostname    *string `tfsdk:"hostname" json:"hostname,omitempty"`
						Ip_address  *string `tfsdk:"ip_address" json:"ip-address,omitempty"`
						Ip_protocol *string `tfsdk:"ip_protocol" json:"ip-protocol,omitempty"`
						Percent     *string `tfsdk:"percent" json:"percent,omitempty"`
						Source_port *string `tfsdk:"source_port" json:"source-port,omitempty"`
					} `tfsdk:"network_corrupt" json:"network-corrupt,omitempty"`
					Network_delay *struct {
						Accept_tcp_flags *string `tfsdk:"accept_tcp_flags" json:"accept-tcp-flags,omitempty"`
						Correlation      *string `tfsdk:"correlation" json:"correlation,omitempty"`
						Device           *string `tfsdk:"device" json:"device,omitempty"`
						Egress_port      *string `tfsdk:"egress_port" json:"egress-port,omitempty"`
						Hostname         *string `tfsdk:"hostname" json:"hostname,omitempty"`
						Ip_address       *string `tfsdk:"ip_address" json:"ip-address,omitempty"`
						Ip_protocol      *string `tfsdk:"ip_protocol" json:"ip-protocol,omitempty"`
						Jitter           *string `tfsdk:"jitter" json:"jitter,omitempty"`
						Latency          *string `tfsdk:"latency" json:"latency,omitempty"`
						Source_port      *string `tfsdk:"source_port" json:"source-port,omitempty"`
					} `tfsdk:"network_delay" json:"network-delay,omitempty"`
					Network_dns *struct {
						Dns_domain_name *string `tfsdk:"dns_domain_name" json:"dns-domain-name,omitempty"`
						Dns_ip          *string `tfsdk:"dns_ip" json:"dns-ip,omitempty"`
						Dns_server      *string `tfsdk:"dns_server" json:"dns-server,omitempty"`
					} `tfsdk:"network_dns" json:"network-dns,omitempty"`
					Network_down *struct {
						Device   *string `tfsdk:"device" json:"device,omitempty"`
						Duration *string `tfsdk:"duration" json:"duration,omitempty"`
					} `tfsdk:"network_down" json:"network-down,omitempty"`
					Network_duplicate *struct {
						Correlation *string `tfsdk:"correlation" json:"correlation,omitempty"`
						Device      *string `tfsdk:"device" json:"device,omitempty"`
						Egress_port *string `tfsdk:"egress_port" json:"egress-port,omitempty"`
						Hostname    *string `tfsdk:"hostname" json:"hostname,omitempty"`
						Ip_address  *string `tfsdk:"ip_address" json:"ip-address,omitempty"`
						Ip_protocol *string `tfsdk:"ip_protocol" json:"ip-protocol,omitempty"`
						Percent     *string `tfsdk:"percent" json:"percent,omitempty"`
						Source_port *string `tfsdk:"source_port" json:"source-port,omitempty"`
					} `tfsdk:"network_duplicate" json:"network-duplicate,omitempty"`
					Network_flood *struct {
						Duration   *string `tfsdk:"duration" json:"duration,omitempty"`
						Ip_address *string `tfsdk:"ip_address" json:"ip-address,omitempty"`
						Parallel   *int64  `tfsdk:"parallel" json:"parallel,omitempty"`
						Port       *string `tfsdk:"port" json:"port,omitempty"`
						Rate       *string `tfsdk:"rate" json:"rate,omitempty"`
					} `tfsdk:"network_flood" json:"network-flood,omitempty"`
					Network_loss *struct {
						Correlation *string `tfsdk:"correlation" json:"correlation,omitempty"`
						Device      *string `tfsdk:"device" json:"device,omitempty"`
						Egress_port *string `tfsdk:"egress_port" json:"egress-port,omitempty"`
						Hostname    *string `tfsdk:"hostname" json:"hostname,omitempty"`
						Ip_address  *string `tfsdk:"ip_address" json:"ip-address,omitempty"`
						Ip_protocol *string `tfsdk:"ip_protocol" json:"ip-protocol,omitempty"`
						Percent     *string `tfsdk:"percent" json:"percent,omitempty"`
						Source_port *string `tfsdk:"source_port" json:"source-port,omitempty"`
					} `tfsdk:"network_loss" json:"network-loss,omitempty"`
					Network_partition *struct {
						Accept_tcp_flags *string `tfsdk:"accept_tcp_flags" json:"accept-tcp-flags,omitempty"`
						Device           *string `tfsdk:"device" json:"device,omitempty"`
						Direction        *string `tfsdk:"direction" json:"direction,omitempty"`
						Hostname         *string `tfsdk:"hostname" json:"hostname,omitempty"`
						Ip_address       *string `tfsdk:"ip_address" json:"ip-address,omitempty"`
						Ip_protocol      *string `tfsdk:"ip_protocol" json:"ip-protocol,omitempty"`
					} `tfsdk:"network_partition" json:"network-partition,omitempty"`
					Process *struct {
						Process    *string `tfsdk:"process" json:"process,omitempty"`
						RecoverCmd *string `tfsdk:"recover_cmd" json:"recoverCmd,omitempty"`
						Signal     *int64  `tfsdk:"signal" json:"signal,omitempty"`
					} `tfsdk:"process" json:"process,omitempty"`
					Redis_cacheLimit *struct {
						Addr      *string `tfsdk:"addr" json:"addr,omitempty"`
						CacheSize *string `tfsdk:"cache_size" json:"cacheSize,omitempty"`
						Password  *string `tfsdk:"password" json:"password,omitempty"`
						Percent   *string `tfsdk:"percent" json:"percent,omitempty"`
					} `tfsdk:"redis_cache_limit" json:"redis-cacheLimit,omitempty"`
					Redis_expiration *struct {
						Addr       *string `tfsdk:"addr" json:"addr,omitempty"`
						Expiration *string `tfsdk:"expiration" json:"expiration,omitempty"`
						Key        *string `tfsdk:"key" json:"key,omitempty"`
						Option     *string `tfsdk:"option" json:"option,omitempty"`
						Password   *string `tfsdk:"password" json:"password,omitempty"`
					} `tfsdk:"redis_expiration" json:"redis-expiration,omitempty"`
					Redis_penetration *struct {
						Addr       *string `tfsdk:"addr" json:"addr,omitempty"`
						Password   *string `tfsdk:"password" json:"password,omitempty"`
						RequestNum *int64  `tfsdk:"request_num" json:"requestNum,omitempty"`
					} `tfsdk:"redis_penetration" json:"redis-penetration,omitempty"`
					Redis_restart *struct {
						Addr        *string `tfsdk:"addr" json:"addr,omitempty"`
						Conf        *string `tfsdk:"conf" json:"conf,omitempty"`
						FlushConfig *bool   `tfsdk:"flush_config" json:"flushConfig,omitempty"`
						Password    *string `tfsdk:"password" json:"password,omitempty"`
						RedisPath   *bool   `tfsdk:"redis_path" json:"redisPath,omitempty"`
					} `tfsdk:"redis_restart" json:"redis-restart,omitempty"`
					Redis_stop *struct {
						Addr        *string `tfsdk:"addr" json:"addr,omitempty"`
						Conf        *string `tfsdk:"conf" json:"conf,omitempty"`
						FlushConfig *bool   `tfsdk:"flush_config" json:"flushConfig,omitempty"`
						Password    *string `tfsdk:"password" json:"password,omitempty"`
						RedisPath   *bool   `tfsdk:"redis_path" json:"redisPath,omitempty"`
					} `tfsdk:"redis_stop" json:"redis-stop,omitempty"`
					RemoteCluster *string `tfsdk:"remote_cluster" json:"remoteCluster,omitempty"`
					Selector      *struct {
						AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" json:"annotationSelectors,omitempty"`
						ExpressionSelectors *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"expression_selectors" json:"expressionSelectors,omitempty"`
						FieldSelectors   *map[string]string   `tfsdk:"field_selectors" json:"fieldSelectors,omitempty"`
						LabelSelectors   *map[string]string   `tfsdk:"label_selectors" json:"labelSelectors,omitempty"`
						Namespaces       *[]string            `tfsdk:"namespaces" json:"namespaces,omitempty"`
						PhysicalMachines *map[string][]string `tfsdk:"physical_machines" json:"physicalMachines,omitempty"`
					} `tfsdk:"selector" json:"selector,omitempty"`
					Stress_cpu *struct {
						Load    *int64    `tfsdk:"load" json:"load,omitempty"`
						Options *[]string `tfsdk:"options" json:"options,omitempty"`
						Workers *int64    `tfsdk:"workers" json:"workers,omitempty"`
					} `tfsdk:"stress_cpu" json:"stress-cpu,omitempty"`
					Stress_mem *struct {
						Options *[]string `tfsdk:"options" json:"options,omitempty"`
						Size    *string   `tfsdk:"size" json:"size,omitempty"`
					} `tfsdk:"stress_mem" json:"stress-mem,omitempty"`
					Uid          *string `tfsdk:"uid" json:"uid,omitempty"`
					User_defined *struct {
						AttackCmd  *string `tfsdk:"attack_cmd" json:"attackCmd,omitempty"`
						RecoverCmd *string `tfsdk:"recover_cmd" json:"recoverCmd,omitempty"`
					} `tfsdk:"user_defined" json:"user_defined,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
					Vm    *struct {
						Vm_name *string `tfsdk:"vm_name" json:"vm-name,omitempty"`
					} `tfsdk:"vm" json:"vm,omitempty"`
				} `tfsdk:"physicalmachine_chaos" json:"physicalmachineChaos,omitempty"`
				PodChaos *struct {
					Action         *string   `tfsdk:"action" json:"action,omitempty"`
					ContainerNames *[]string `tfsdk:"container_names" json:"containerNames,omitempty"`
					Duration       *string   `tfsdk:"duration" json:"duration,omitempty"`
					GracePeriod    *int64    `tfsdk:"grace_period" json:"gracePeriod,omitempty"`
					Mode           *string   `tfsdk:"mode" json:"mode,omitempty"`
					RemoteCluster  *string   `tfsdk:"remote_cluster" json:"remoteCluster,omitempty"`
					Selector       *struct {
						AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" json:"annotationSelectors,omitempty"`
						ExpressionSelectors *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"expression_selectors" json:"expressionSelectors,omitempty"`
						FieldSelectors    *map[string]string   `tfsdk:"field_selectors" json:"fieldSelectors,omitempty"`
						LabelSelectors    *map[string]string   `tfsdk:"label_selectors" json:"labelSelectors,omitempty"`
						Namespaces        *[]string            `tfsdk:"namespaces" json:"namespaces,omitempty"`
						NodeSelectors     *map[string]string   `tfsdk:"node_selectors" json:"nodeSelectors,omitempty"`
						Nodes             *[]string            `tfsdk:"nodes" json:"nodes,omitempty"`
						PodPhaseSelectors *[]string            `tfsdk:"pod_phase_selectors" json:"podPhaseSelectors,omitempty"`
						Pods              *map[string][]string `tfsdk:"pods" json:"pods,omitempty"`
					} `tfsdk:"selector" json:"selector,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"pod_chaos" json:"podChaos,omitempty"`
				Schedule                *string `tfsdk:"schedule" json:"schedule,omitempty"`
				StartingDeadlineSeconds *int64  `tfsdk:"starting_deadline_seconds" json:"startingDeadlineSeconds,omitempty"`
				StressChaos             *struct {
					ContainerNames *[]string `tfsdk:"container_names" json:"containerNames,omitempty"`
					Duration       *string   `tfsdk:"duration" json:"duration,omitempty"`
					Mode           *string   `tfsdk:"mode" json:"mode,omitempty"`
					RemoteCluster  *string   `tfsdk:"remote_cluster" json:"remoteCluster,omitempty"`
					Selector       *struct {
						AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" json:"annotationSelectors,omitempty"`
						ExpressionSelectors *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"expression_selectors" json:"expressionSelectors,omitempty"`
						FieldSelectors    *map[string]string   `tfsdk:"field_selectors" json:"fieldSelectors,omitempty"`
						LabelSelectors    *map[string]string   `tfsdk:"label_selectors" json:"labelSelectors,omitempty"`
						Namespaces        *[]string            `tfsdk:"namespaces" json:"namespaces,omitempty"`
						NodeSelectors     *map[string]string   `tfsdk:"node_selectors" json:"nodeSelectors,omitempty"`
						Nodes             *[]string            `tfsdk:"nodes" json:"nodes,omitempty"`
						PodPhaseSelectors *[]string            `tfsdk:"pod_phase_selectors" json:"podPhaseSelectors,omitempty"`
						Pods              *map[string][]string `tfsdk:"pods" json:"pods,omitempty"`
					} `tfsdk:"selector" json:"selector,omitempty"`
					StressngStressors *string `tfsdk:"stressng_stressors" json:"stressngStressors,omitempty"`
					Stressors         *struct {
						Cpu *struct {
							Load    *int64    `tfsdk:"load" json:"load,omitempty"`
							Options *[]string `tfsdk:"options" json:"options,omitempty"`
							Workers *int64    `tfsdk:"workers" json:"workers,omitempty"`
						} `tfsdk:"cpu" json:"cpu,omitempty"`
						Memory *struct {
							OomScoreAdj *int64    `tfsdk:"oom_score_adj" json:"oomScoreAdj,omitempty"`
							Options     *[]string `tfsdk:"options" json:"options,omitempty"`
							Size        *string   `tfsdk:"size" json:"size,omitempty"`
							Workers     *int64    `tfsdk:"workers" json:"workers,omitempty"`
						} `tfsdk:"memory" json:"memory,omitempty"`
					} `tfsdk:"stressors" json:"stressors,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"stress_chaos" json:"stressChaos,omitempty"`
				TimeChaos *struct {
					ClockIds       *[]string `tfsdk:"clock_ids" json:"clockIds,omitempty"`
					ContainerNames *[]string `tfsdk:"container_names" json:"containerNames,omitempty"`
					Duration       *string   `tfsdk:"duration" json:"duration,omitempty"`
					Mode           *string   `tfsdk:"mode" json:"mode,omitempty"`
					RemoteCluster  *string   `tfsdk:"remote_cluster" json:"remoteCluster,omitempty"`
					Selector       *struct {
						AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" json:"annotationSelectors,omitempty"`
						ExpressionSelectors *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"expression_selectors" json:"expressionSelectors,omitempty"`
						FieldSelectors    *map[string]string   `tfsdk:"field_selectors" json:"fieldSelectors,omitempty"`
						LabelSelectors    *map[string]string   `tfsdk:"label_selectors" json:"labelSelectors,omitempty"`
						Namespaces        *[]string            `tfsdk:"namespaces" json:"namespaces,omitempty"`
						NodeSelectors     *map[string]string   `tfsdk:"node_selectors" json:"nodeSelectors,omitempty"`
						Nodes             *[]string            `tfsdk:"nodes" json:"nodes,omitempty"`
						PodPhaseSelectors *[]string            `tfsdk:"pod_phase_selectors" json:"podPhaseSelectors,omitempty"`
						Pods              *map[string][]string `tfsdk:"pods" json:"pods,omitempty"`
					} `tfsdk:"selector" json:"selector,omitempty"`
					TimeOffset *string `tfsdk:"time_offset" json:"timeOffset,omitempty"`
					Value      *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"time_chaos" json:"timeChaos,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"schedule" json:"schedule,omitempty"`
			StatusCheck *struct {
				Duration         *string `tfsdk:"duration" json:"duration,omitempty"`
				FailureThreshold *int64  `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
				Http             *struct {
					Body     *string `tfsdk:"body" json:"body,omitempty"`
					Criteria *struct {
						StatusCode *string `tfsdk:"status_code" json:"statusCode,omitempty"`
					} `tfsdk:"criteria" json:"criteria,omitempty"`
					Headers *map[string][]string `tfsdk:"headers" json:"headers,omitempty"`
					Method  *string              `tfsdk:"method" json:"method,omitempty"`
					Url     *string              `tfsdk:"url" json:"url,omitempty"`
				} `tfsdk:"http" json:"http,omitempty"`
				IntervalSeconds     *int64  `tfsdk:"interval_seconds" json:"intervalSeconds,omitempty"`
				Mode                *string `tfsdk:"mode" json:"mode,omitempty"`
				RecordsHistoryLimit *int64  `tfsdk:"records_history_limit" json:"recordsHistoryLimit,omitempty"`
				SuccessThreshold    *int64  `tfsdk:"success_threshold" json:"successThreshold,omitempty"`
				TimeoutSeconds      *int64  `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
				Type                *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"status_check" json:"statusCheck,omitempty"`
			StressChaos *struct {
				ContainerNames *[]string `tfsdk:"container_names" json:"containerNames,omitempty"`
				Duration       *string   `tfsdk:"duration" json:"duration,omitempty"`
				Mode           *string   `tfsdk:"mode" json:"mode,omitempty"`
				RemoteCluster  *string   `tfsdk:"remote_cluster" json:"remoteCluster,omitempty"`
				Selector       *struct {
					AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" json:"annotationSelectors,omitempty"`
					ExpressionSelectors *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"expression_selectors" json:"expressionSelectors,omitempty"`
					FieldSelectors    *map[string]string   `tfsdk:"field_selectors" json:"fieldSelectors,omitempty"`
					LabelSelectors    *map[string]string   `tfsdk:"label_selectors" json:"labelSelectors,omitempty"`
					Namespaces        *[]string            `tfsdk:"namespaces" json:"namespaces,omitempty"`
					NodeSelectors     *map[string]string   `tfsdk:"node_selectors" json:"nodeSelectors,omitempty"`
					Nodes             *[]string            `tfsdk:"nodes" json:"nodes,omitempty"`
					PodPhaseSelectors *[]string            `tfsdk:"pod_phase_selectors" json:"podPhaseSelectors,omitempty"`
					Pods              *map[string][]string `tfsdk:"pods" json:"pods,omitempty"`
				} `tfsdk:"selector" json:"selector,omitempty"`
				StressngStressors *string `tfsdk:"stressng_stressors" json:"stressngStressors,omitempty"`
				Stressors         *struct {
					Cpu *struct {
						Load    *int64    `tfsdk:"load" json:"load,omitempty"`
						Options *[]string `tfsdk:"options" json:"options,omitempty"`
						Workers *int64    `tfsdk:"workers" json:"workers,omitempty"`
					} `tfsdk:"cpu" json:"cpu,omitempty"`
					Memory *struct {
						OomScoreAdj *int64    `tfsdk:"oom_score_adj" json:"oomScoreAdj,omitempty"`
						Options     *[]string `tfsdk:"options" json:"options,omitempty"`
						Size        *string   `tfsdk:"size" json:"size,omitempty"`
						Workers     *int64    `tfsdk:"workers" json:"workers,omitempty"`
					} `tfsdk:"memory" json:"memory,omitempty"`
				} `tfsdk:"stressors" json:"stressors,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"stress_chaos" json:"stressChaos,omitempty"`
			Task *struct {
				Container *struct {
					Args    *[]string `tfsdk:"args" json:"args,omitempty"`
					Command *[]string `tfsdk:"command" json:"command,omitempty"`
					Env     *[]struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Value     *string `tfsdk:"value" json:"value,omitempty"`
						ValueFrom *struct {
							ConfigMapKeyRef *struct {
								Key      *string `tfsdk:"key" json:"key,omitempty"`
								Name     *string `tfsdk:"name" json:"name,omitempty"`
								Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
							} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
							FieldRef *struct {
								ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
								FieldPath  *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
							} `tfsdk:"field_ref" json:"fieldRef,omitempty"`
							ResourceFieldRef *struct {
								ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
								Divisor       *string `tfsdk:"divisor" json:"divisor,omitempty"`
								Resource      *string `tfsdk:"resource" json:"resource,omitempty"`
							} `tfsdk:"resource_field_ref" json:"resourceFieldRef,omitempty"`
							SecretKeyRef *struct {
								Key      *string `tfsdk:"key" json:"key,omitempty"`
								Name     *string `tfsdk:"name" json:"name,omitempty"`
								Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
							} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						} `tfsdk:"value_from" json:"valueFrom,omitempty"`
					} `tfsdk:"env" json:"env,omitempty"`
					EnvFrom *[]struct {
						ConfigMapRef *struct {
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"config_map_ref" json:"configMapRef,omitempty"`
						Prefix    *string `tfsdk:"prefix" json:"prefix,omitempty"`
						SecretRef *struct {
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
					} `tfsdk:"env_from" json:"envFrom,omitempty"`
					Image           *string `tfsdk:"image" json:"image,omitempty"`
					ImagePullPolicy *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
					Lifecycle       *struct {
						PostStart *struct {
							Exec *struct {
								Command *[]string `tfsdk:"command" json:"command,omitempty"`
							} `tfsdk:"exec" json:"exec,omitempty"`
							HttpGet *struct {
								Host        *string `tfsdk:"host" json:"host,omitempty"`
								HttpHeaders *[]struct {
									Name  *string `tfsdk:"name" json:"name,omitempty"`
									Value *string `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"http_headers" json:"httpHeaders,omitempty"`
								Path   *string `tfsdk:"path" json:"path,omitempty"`
								Port   *string `tfsdk:"port" json:"port,omitempty"`
								Scheme *string `tfsdk:"scheme" json:"scheme,omitempty"`
							} `tfsdk:"http_get" json:"httpGet,omitempty"`
							TcpSocket *struct {
								Host *string `tfsdk:"host" json:"host,omitempty"`
								Port *string `tfsdk:"port" json:"port,omitempty"`
							} `tfsdk:"tcp_socket" json:"tcpSocket,omitempty"`
						} `tfsdk:"post_start" json:"postStart,omitempty"`
						PreStop *struct {
							Exec *struct {
								Command *[]string `tfsdk:"command" json:"command,omitempty"`
							} `tfsdk:"exec" json:"exec,omitempty"`
							HttpGet *struct {
								Host        *string `tfsdk:"host" json:"host,omitempty"`
								HttpHeaders *[]struct {
									Name  *string `tfsdk:"name" json:"name,omitempty"`
									Value *string `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"http_headers" json:"httpHeaders,omitempty"`
								Path   *string `tfsdk:"path" json:"path,omitempty"`
								Port   *string `tfsdk:"port" json:"port,omitempty"`
								Scheme *string `tfsdk:"scheme" json:"scheme,omitempty"`
							} `tfsdk:"http_get" json:"httpGet,omitempty"`
							TcpSocket *struct {
								Host *string `tfsdk:"host" json:"host,omitempty"`
								Port *string `tfsdk:"port" json:"port,omitempty"`
							} `tfsdk:"tcp_socket" json:"tcpSocket,omitempty"`
						} `tfsdk:"pre_stop" json:"preStop,omitempty"`
					} `tfsdk:"lifecycle" json:"lifecycle,omitempty"`
					LivenessProbe *struct {
						Exec *struct {
							Command *[]string `tfsdk:"command" json:"command,omitempty"`
						} `tfsdk:"exec" json:"exec,omitempty"`
						FailureThreshold *int64 `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
						Grpc             *struct {
							Port    *int64  `tfsdk:"port" json:"port,omitempty"`
							Service *string `tfsdk:"service" json:"service,omitempty"`
						} `tfsdk:"grpc" json:"grpc,omitempty"`
						HttpGet *struct {
							Host        *string `tfsdk:"host" json:"host,omitempty"`
							HttpHeaders *[]struct {
								Name  *string `tfsdk:"name" json:"name,omitempty"`
								Value *string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"http_headers" json:"httpHeaders,omitempty"`
							Path   *string `tfsdk:"path" json:"path,omitempty"`
							Port   *string `tfsdk:"port" json:"port,omitempty"`
							Scheme *string `tfsdk:"scheme" json:"scheme,omitempty"`
						} `tfsdk:"http_get" json:"httpGet,omitempty"`
						InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" json:"initialDelaySeconds,omitempty"`
						PeriodSeconds       *int64 `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
						SuccessThreshold    *int64 `tfsdk:"success_threshold" json:"successThreshold,omitempty"`
						TcpSocket           *struct {
							Host *string `tfsdk:"host" json:"host,omitempty"`
							Port *string `tfsdk:"port" json:"port,omitempty"`
						} `tfsdk:"tcp_socket" json:"tcpSocket,omitempty"`
						TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" json:"terminationGracePeriodSeconds,omitempty"`
						TimeoutSeconds                *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
					} `tfsdk:"liveness_probe" json:"livenessProbe,omitempty"`
					Name  *string `tfsdk:"name" json:"name,omitempty"`
					Ports *[]struct {
						ContainerPort *int64  `tfsdk:"container_port" json:"containerPort,omitempty"`
						HostIP        *string `tfsdk:"host_ip" json:"hostIP,omitempty"`
						HostPort      *int64  `tfsdk:"host_port" json:"hostPort,omitempty"`
						Name          *string `tfsdk:"name" json:"name,omitempty"`
						Protocol      *string `tfsdk:"protocol" json:"protocol,omitempty"`
					} `tfsdk:"ports" json:"ports,omitempty"`
					ReadinessProbe *struct {
						Exec *struct {
							Command *[]string `tfsdk:"command" json:"command,omitempty"`
						} `tfsdk:"exec" json:"exec,omitempty"`
						FailureThreshold *int64 `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
						Grpc             *struct {
							Port    *int64  `tfsdk:"port" json:"port,omitempty"`
							Service *string `tfsdk:"service" json:"service,omitempty"`
						} `tfsdk:"grpc" json:"grpc,omitempty"`
						HttpGet *struct {
							Host        *string `tfsdk:"host" json:"host,omitempty"`
							HttpHeaders *[]struct {
								Name  *string `tfsdk:"name" json:"name,omitempty"`
								Value *string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"http_headers" json:"httpHeaders,omitempty"`
							Path   *string `tfsdk:"path" json:"path,omitempty"`
							Port   *string `tfsdk:"port" json:"port,omitempty"`
							Scheme *string `tfsdk:"scheme" json:"scheme,omitempty"`
						} `tfsdk:"http_get" json:"httpGet,omitempty"`
						InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" json:"initialDelaySeconds,omitempty"`
						PeriodSeconds       *int64 `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
						SuccessThreshold    *int64 `tfsdk:"success_threshold" json:"successThreshold,omitempty"`
						TcpSocket           *struct {
							Host *string `tfsdk:"host" json:"host,omitempty"`
							Port *string `tfsdk:"port" json:"port,omitempty"`
						} `tfsdk:"tcp_socket" json:"tcpSocket,omitempty"`
						TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" json:"terminationGracePeriodSeconds,omitempty"`
						TimeoutSeconds                *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
					} `tfsdk:"readiness_probe" json:"readinessProbe,omitempty"`
					Resources *struct {
						Claims *[]struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"claims" json:"claims,omitempty"`
						Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
						Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
					} `tfsdk:"resources" json:"resources,omitempty"`
					SecurityContext *struct {
						AllowPrivilegeEscalation *bool `tfsdk:"allow_privilege_escalation" json:"allowPrivilegeEscalation,omitempty"`
						Capabilities             *struct {
							Add  *[]string `tfsdk:"add" json:"add,omitempty"`
							Drop *[]string `tfsdk:"drop" json:"drop,omitempty"`
						} `tfsdk:"capabilities" json:"capabilities,omitempty"`
						Privileged             *bool   `tfsdk:"privileged" json:"privileged,omitempty"`
						ProcMount              *string `tfsdk:"proc_mount" json:"procMount,omitempty"`
						ReadOnlyRootFilesystem *bool   `tfsdk:"read_only_root_filesystem" json:"readOnlyRootFilesystem,omitempty"`
						RunAsGroup             *int64  `tfsdk:"run_as_group" json:"runAsGroup,omitempty"`
						RunAsNonRoot           *bool   `tfsdk:"run_as_non_root" json:"runAsNonRoot,omitempty"`
						RunAsUser              *int64  `tfsdk:"run_as_user" json:"runAsUser,omitempty"`
						SeLinuxOptions         *struct {
							Level *string `tfsdk:"level" json:"level,omitempty"`
							Role  *string `tfsdk:"role" json:"role,omitempty"`
							Type  *string `tfsdk:"type" json:"type,omitempty"`
							User  *string `tfsdk:"user" json:"user,omitempty"`
						} `tfsdk:"se_linux_options" json:"seLinuxOptions,omitempty"`
						SeccompProfile *struct {
							LocalhostProfile *string `tfsdk:"localhost_profile" json:"localhostProfile,omitempty"`
							Type             *string `tfsdk:"type" json:"type,omitempty"`
						} `tfsdk:"seccomp_profile" json:"seccompProfile,omitempty"`
						WindowsOptions *struct {
							GmsaCredentialSpec     *string `tfsdk:"gmsa_credential_spec" json:"gmsaCredentialSpec,omitempty"`
							GmsaCredentialSpecName *string `tfsdk:"gmsa_credential_spec_name" json:"gmsaCredentialSpecName,omitempty"`
							HostProcess            *bool   `tfsdk:"host_process" json:"hostProcess,omitempty"`
							RunAsUserName          *string `tfsdk:"run_as_user_name" json:"runAsUserName,omitempty"`
						} `tfsdk:"windows_options" json:"windowsOptions,omitempty"`
					} `tfsdk:"security_context" json:"securityContext,omitempty"`
					StartupProbe *struct {
						Exec *struct {
							Command *[]string `tfsdk:"command" json:"command,omitempty"`
						} `tfsdk:"exec" json:"exec,omitempty"`
						FailureThreshold *int64 `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
						Grpc             *struct {
							Port    *int64  `tfsdk:"port" json:"port,omitempty"`
							Service *string `tfsdk:"service" json:"service,omitempty"`
						} `tfsdk:"grpc" json:"grpc,omitempty"`
						HttpGet *struct {
							Host        *string `tfsdk:"host" json:"host,omitempty"`
							HttpHeaders *[]struct {
								Name  *string `tfsdk:"name" json:"name,omitempty"`
								Value *string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"http_headers" json:"httpHeaders,omitempty"`
							Path   *string `tfsdk:"path" json:"path,omitempty"`
							Port   *string `tfsdk:"port" json:"port,omitempty"`
							Scheme *string `tfsdk:"scheme" json:"scheme,omitempty"`
						} `tfsdk:"http_get" json:"httpGet,omitempty"`
						InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" json:"initialDelaySeconds,omitempty"`
						PeriodSeconds       *int64 `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
						SuccessThreshold    *int64 `tfsdk:"success_threshold" json:"successThreshold,omitempty"`
						TcpSocket           *struct {
							Host *string `tfsdk:"host" json:"host,omitempty"`
							Port *string `tfsdk:"port" json:"port,omitempty"`
						} `tfsdk:"tcp_socket" json:"tcpSocket,omitempty"`
						TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" json:"terminationGracePeriodSeconds,omitempty"`
						TimeoutSeconds                *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
					} `tfsdk:"startup_probe" json:"startupProbe,omitempty"`
					Stdin                    *bool   `tfsdk:"stdin" json:"stdin,omitempty"`
					StdinOnce                *bool   `tfsdk:"stdin_once" json:"stdinOnce,omitempty"`
					TerminationMessagePath   *string `tfsdk:"termination_message_path" json:"terminationMessagePath,omitempty"`
					TerminationMessagePolicy *string `tfsdk:"termination_message_policy" json:"terminationMessagePolicy,omitempty"`
					Tty                      *bool   `tfsdk:"tty" json:"tty,omitempty"`
					VolumeDevices            *[]struct {
						DevicePath *string `tfsdk:"device_path" json:"devicePath,omitempty"`
						Name       *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"volume_devices" json:"volumeDevices,omitempty"`
					VolumeMounts *[]struct {
						MountPath        *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
						MountPropagation *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
						Name             *string `tfsdk:"name" json:"name,omitempty"`
						ReadOnly         *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
						SubPath          *string `tfsdk:"sub_path" json:"subPath,omitempty"`
						SubPathExpr      *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
					} `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
					WorkingDir *string `tfsdk:"working_dir" json:"workingDir,omitempty"`
				} `tfsdk:"container" json:"container,omitempty"`
				Volumes *[]struct {
					AwsElasticBlockStore *struct {
						FsType    *string `tfsdk:"fs_type" json:"fsType,omitempty"`
						Partition *int64  `tfsdk:"partition" json:"partition,omitempty"`
						ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
						VolumeID  *string `tfsdk:"volume_id" json:"volumeID,omitempty"`
					} `tfsdk:"aws_elastic_block_store" json:"awsElasticBlockStore,omitempty"`
					AzureDisk *struct {
						CachingMode *string `tfsdk:"caching_mode" json:"cachingMode,omitempty"`
						DiskName    *string `tfsdk:"disk_name" json:"diskName,omitempty"`
						DiskURI     *string `tfsdk:"disk_uri" json:"diskURI,omitempty"`
						FsType      *string `tfsdk:"fs_type" json:"fsType,omitempty"`
						Kind        *string `tfsdk:"kind" json:"kind,omitempty"`
						ReadOnly    *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
					} `tfsdk:"azure_disk" json:"azureDisk,omitempty"`
					AzureFile *struct {
						ReadOnly   *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
						SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
						ShareName  *string `tfsdk:"share_name" json:"shareName,omitempty"`
					} `tfsdk:"azure_file" json:"azureFile,omitempty"`
					Cephfs *struct {
						Monitors   *[]string `tfsdk:"monitors" json:"monitors,omitempty"`
						Path       *string   `tfsdk:"path" json:"path,omitempty"`
						ReadOnly   *bool     `tfsdk:"read_only" json:"readOnly,omitempty"`
						SecretFile *string   `tfsdk:"secret_file" json:"secretFile,omitempty"`
						SecretRef  *struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
						User *string `tfsdk:"user" json:"user,omitempty"`
					} `tfsdk:"cephfs" json:"cephfs,omitempty"`
					Cinder *struct {
						FsType    *string `tfsdk:"fs_type" json:"fsType,omitempty"`
						ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
						SecretRef *struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
						VolumeID *string `tfsdk:"volume_id" json:"volumeID,omitempty"`
					} `tfsdk:"cinder" json:"cinder,omitempty"`
					ConfigMap *struct {
						DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
						Items       *[]struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
							Path *string `tfsdk:"path" json:"path,omitempty"`
						} `tfsdk:"items" json:"items,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"config_map" json:"configMap,omitempty"`
					Csi *struct {
						Driver               *string `tfsdk:"driver" json:"driver,omitempty"`
						FsType               *string `tfsdk:"fs_type" json:"fsType,omitempty"`
						NodePublishSecretRef *struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"node_publish_secret_ref" json:"nodePublishSecretRef,omitempty"`
						ReadOnly         *bool              `tfsdk:"read_only" json:"readOnly,omitempty"`
						VolumeAttributes *map[string]string `tfsdk:"volume_attributes" json:"volumeAttributes,omitempty"`
					} `tfsdk:"csi" json:"csi,omitempty"`
					DownwardAPI *struct {
						DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
						Items       *[]struct {
							FieldRef *struct {
								ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
								FieldPath  *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
							} `tfsdk:"field_ref" json:"fieldRef,omitempty"`
							Mode             *int64  `tfsdk:"mode" json:"mode,omitempty"`
							Path             *string `tfsdk:"path" json:"path,omitempty"`
							ResourceFieldRef *struct {
								ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
								Divisor       *string `tfsdk:"divisor" json:"divisor,omitempty"`
								Resource      *string `tfsdk:"resource" json:"resource,omitempty"`
							} `tfsdk:"resource_field_ref" json:"resourceFieldRef,omitempty"`
						} `tfsdk:"items" json:"items,omitempty"`
					} `tfsdk:"downward_api" json:"downwardAPI,omitempty"`
					EmptyDir *struct {
						Medium    *string `tfsdk:"medium" json:"medium,omitempty"`
						SizeLimit *string `tfsdk:"size_limit" json:"sizeLimit,omitempty"`
					} `tfsdk:"empty_dir" json:"emptyDir,omitempty"`
					Ephemeral *struct {
						VolumeClaimTemplate *struct {
							Metadata *map[string]string `tfsdk:"metadata" json:"metadata,omitempty"`
							Spec     *struct {
								AccessModes *[]string `tfsdk:"access_modes" json:"accessModes,omitempty"`
								DataSource  *struct {
									ApiGroup *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
									Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
									Name     *string `tfsdk:"name" json:"name,omitempty"`
								} `tfsdk:"data_source" json:"dataSource,omitempty"`
								DataSourceRef *struct {
									ApiGroup  *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
									Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
									Name      *string `tfsdk:"name" json:"name,omitempty"`
									Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
								} `tfsdk:"data_source_ref" json:"dataSourceRef,omitempty"`
								Resources *struct {
									Claims *[]struct {
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"claims" json:"claims,omitempty"`
									Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
									Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
								} `tfsdk:"resources" json:"resources,omitempty"`
								Selector *struct {
									MatchExpressions *[]struct {
										Key      *string   `tfsdk:"key" json:"key,omitempty"`
										Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
										Values   *[]string `tfsdk:"values" json:"values,omitempty"`
									} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
									MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
								} `tfsdk:"selector" json:"selector,omitempty"`
								StorageClassName *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
								VolumeMode       *string `tfsdk:"volume_mode" json:"volumeMode,omitempty"`
								VolumeName       *string `tfsdk:"volume_name" json:"volumeName,omitempty"`
							} `tfsdk:"spec" json:"spec,omitempty"`
						} `tfsdk:"volume_claim_template" json:"volumeClaimTemplate,omitempty"`
					} `tfsdk:"ephemeral" json:"ephemeral,omitempty"`
					Fc *struct {
						FsType     *string   `tfsdk:"fs_type" json:"fsType,omitempty"`
						Lun        *int64    `tfsdk:"lun" json:"lun,omitempty"`
						ReadOnly   *bool     `tfsdk:"read_only" json:"readOnly,omitempty"`
						TargetWWNs *[]string `tfsdk:"target_ww_ns" json:"targetWWNs,omitempty"`
						Wwids      *[]string `tfsdk:"wwids" json:"wwids,omitempty"`
					} `tfsdk:"fc" json:"fc,omitempty"`
					FlexVolume *struct {
						Driver    *string            `tfsdk:"driver" json:"driver,omitempty"`
						FsType    *string            `tfsdk:"fs_type" json:"fsType,omitempty"`
						Options   *map[string]string `tfsdk:"options" json:"options,omitempty"`
						ReadOnly  *bool              `tfsdk:"read_only" json:"readOnly,omitempty"`
						SecretRef *struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
					} `tfsdk:"flex_volume" json:"flexVolume,omitempty"`
					Flocker *struct {
						DatasetName *string `tfsdk:"dataset_name" json:"datasetName,omitempty"`
						DatasetUUID *string `tfsdk:"dataset_uuid" json:"datasetUUID,omitempty"`
					} `tfsdk:"flocker" json:"flocker,omitempty"`
					GcePersistentDisk *struct {
						FsType    *string `tfsdk:"fs_type" json:"fsType,omitempty"`
						Partition *int64  `tfsdk:"partition" json:"partition,omitempty"`
						PdName    *string `tfsdk:"pd_name" json:"pdName,omitempty"`
						ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
					} `tfsdk:"gce_persistent_disk" json:"gcePersistentDisk,omitempty"`
					GitRepo *struct {
						Directory  *string `tfsdk:"directory" json:"directory,omitempty"`
						Repository *string `tfsdk:"repository" json:"repository,omitempty"`
						Revision   *string `tfsdk:"revision" json:"revision,omitempty"`
					} `tfsdk:"git_repo" json:"gitRepo,omitempty"`
					Glusterfs *struct {
						Endpoints *string `tfsdk:"endpoints" json:"endpoints,omitempty"`
						Path      *string `tfsdk:"path" json:"path,omitempty"`
						ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
					} `tfsdk:"glusterfs" json:"glusterfs,omitempty"`
					HostPath *struct {
						Path *string `tfsdk:"path" json:"path,omitempty"`
						Type *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"host_path" json:"hostPath,omitempty"`
					Iscsi *struct {
						ChapAuthDiscovery *bool     `tfsdk:"chap_auth_discovery" json:"chapAuthDiscovery,omitempty"`
						ChapAuthSession   *bool     `tfsdk:"chap_auth_session" json:"chapAuthSession,omitempty"`
						FsType            *string   `tfsdk:"fs_type" json:"fsType,omitempty"`
						InitiatorName     *string   `tfsdk:"initiator_name" json:"initiatorName,omitempty"`
						Iqn               *string   `tfsdk:"iqn" json:"iqn,omitempty"`
						IscsiInterface    *string   `tfsdk:"iscsi_interface" json:"iscsiInterface,omitempty"`
						Lun               *int64    `tfsdk:"lun" json:"lun,omitempty"`
						Portals           *[]string `tfsdk:"portals" json:"portals,omitempty"`
						ReadOnly          *bool     `tfsdk:"read_only" json:"readOnly,omitempty"`
						SecretRef         *struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
						TargetPortal *string `tfsdk:"target_portal" json:"targetPortal,omitempty"`
					} `tfsdk:"iscsi" json:"iscsi,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
					Nfs  *struct {
						Path     *string `tfsdk:"path" json:"path,omitempty"`
						ReadOnly *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
						Server   *string `tfsdk:"server" json:"server,omitempty"`
					} `tfsdk:"nfs" json:"nfs,omitempty"`
					PersistentVolumeClaim *struct {
						ClaimName *string `tfsdk:"claim_name" json:"claimName,omitempty"`
						ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
					} `tfsdk:"persistent_volume_claim" json:"persistentVolumeClaim,omitempty"`
					PhotonPersistentDisk *struct {
						FsType *string `tfsdk:"fs_type" json:"fsType,omitempty"`
						PdID   *string `tfsdk:"pd_id" json:"pdID,omitempty"`
					} `tfsdk:"photon_persistent_disk" json:"photonPersistentDisk,omitempty"`
					PortworxVolume *struct {
						FsType   *string `tfsdk:"fs_type" json:"fsType,omitempty"`
						ReadOnly *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
						VolumeID *string `tfsdk:"volume_id" json:"volumeID,omitempty"`
					} `tfsdk:"portworx_volume" json:"portworxVolume,omitempty"`
					Projected *struct {
						DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
						Sources     *[]struct {
							ConfigMap *struct {
								Items *[]struct {
									Key  *string `tfsdk:"key" json:"key,omitempty"`
									Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
									Path *string `tfsdk:"path" json:"path,omitempty"`
								} `tfsdk:"items" json:"items,omitempty"`
								Name     *string `tfsdk:"name" json:"name,omitempty"`
								Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
							} `tfsdk:"config_map" json:"configMap,omitempty"`
							DownwardAPI *struct {
								Items *[]struct {
									FieldRef *struct {
										ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
										FieldPath  *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
									} `tfsdk:"field_ref" json:"fieldRef,omitempty"`
									Mode             *int64  `tfsdk:"mode" json:"mode,omitempty"`
									Path             *string `tfsdk:"path" json:"path,omitempty"`
									ResourceFieldRef *struct {
										ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
										Divisor       *string `tfsdk:"divisor" json:"divisor,omitempty"`
										Resource      *string `tfsdk:"resource" json:"resource,omitempty"`
									} `tfsdk:"resource_field_ref" json:"resourceFieldRef,omitempty"`
								} `tfsdk:"items" json:"items,omitempty"`
							} `tfsdk:"downward_api" json:"downwardAPI,omitempty"`
							Secret *struct {
								Items *[]struct {
									Key  *string `tfsdk:"key" json:"key,omitempty"`
									Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
									Path *string `tfsdk:"path" json:"path,omitempty"`
								} `tfsdk:"items" json:"items,omitempty"`
								Name     *string `tfsdk:"name" json:"name,omitempty"`
								Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
							} `tfsdk:"secret" json:"secret,omitempty"`
							ServiceAccountToken *struct {
								Audience          *string `tfsdk:"audience" json:"audience,omitempty"`
								ExpirationSeconds *int64  `tfsdk:"expiration_seconds" json:"expirationSeconds,omitempty"`
								Path              *string `tfsdk:"path" json:"path,omitempty"`
							} `tfsdk:"service_account_token" json:"serviceAccountToken,omitempty"`
						} `tfsdk:"sources" json:"sources,omitempty"`
					} `tfsdk:"projected" json:"projected,omitempty"`
					Quobyte *struct {
						Group    *string `tfsdk:"group" json:"group,omitempty"`
						ReadOnly *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
						Registry *string `tfsdk:"registry" json:"registry,omitempty"`
						Tenant   *string `tfsdk:"tenant" json:"tenant,omitempty"`
						User     *string `tfsdk:"user" json:"user,omitempty"`
						Volume   *string `tfsdk:"volume" json:"volume,omitempty"`
					} `tfsdk:"quobyte" json:"quobyte,omitempty"`
					Rbd *struct {
						FsType    *string   `tfsdk:"fs_type" json:"fsType,omitempty"`
						Image     *string   `tfsdk:"image" json:"image,omitempty"`
						Keyring   *string   `tfsdk:"keyring" json:"keyring,omitempty"`
						Monitors  *[]string `tfsdk:"monitors" json:"monitors,omitempty"`
						Pool      *string   `tfsdk:"pool" json:"pool,omitempty"`
						ReadOnly  *bool     `tfsdk:"read_only" json:"readOnly,omitempty"`
						SecretRef *struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
						User *string `tfsdk:"user" json:"user,omitempty"`
					} `tfsdk:"rbd" json:"rbd,omitempty"`
					ScaleIO *struct {
						FsType           *string `tfsdk:"fs_type" json:"fsType,omitempty"`
						Gateway          *string `tfsdk:"gateway" json:"gateway,omitempty"`
						ProtectionDomain *string `tfsdk:"protection_domain" json:"protectionDomain,omitempty"`
						ReadOnly         *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
						SecretRef        *struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
						SslEnabled  *bool   `tfsdk:"ssl_enabled" json:"sslEnabled,omitempty"`
						StorageMode *string `tfsdk:"storage_mode" json:"storageMode,omitempty"`
						StoragePool *string `tfsdk:"storage_pool" json:"storagePool,omitempty"`
						System      *string `tfsdk:"system" json:"system,omitempty"`
						VolumeName  *string `tfsdk:"volume_name" json:"volumeName,omitempty"`
					} `tfsdk:"scale_io" json:"scaleIO,omitempty"`
					Secret *struct {
						DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
						Items       *[]struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
							Path *string `tfsdk:"path" json:"path,omitempty"`
						} `tfsdk:"items" json:"items,omitempty"`
						Optional   *bool   `tfsdk:"optional" json:"optional,omitempty"`
						SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
					} `tfsdk:"secret" json:"secret,omitempty"`
					Storageos *struct {
						FsType    *string `tfsdk:"fs_type" json:"fsType,omitempty"`
						ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
						SecretRef *struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
						VolumeName      *string `tfsdk:"volume_name" json:"volumeName,omitempty"`
						VolumeNamespace *string `tfsdk:"volume_namespace" json:"volumeNamespace,omitempty"`
					} `tfsdk:"storageos" json:"storageos,omitempty"`
					VsphereVolume *struct {
						FsType            *string `tfsdk:"fs_type" json:"fsType,omitempty"`
						StoragePolicyID   *string `tfsdk:"storage_policy_id" json:"storagePolicyID,omitempty"`
						StoragePolicyName *string `tfsdk:"storage_policy_name" json:"storagePolicyName,omitempty"`
						VolumePath        *string `tfsdk:"volume_path" json:"volumePath,omitempty"`
					} `tfsdk:"vsphere_volume" json:"vsphereVolume,omitempty"`
				} `tfsdk:"volumes" json:"volumes,omitempty"`
			} `tfsdk:"task" json:"task,omitempty"`
			TemplateType *string `tfsdk:"template_type" json:"templateType,omitempty"`
			TimeChaos    *struct {
				ClockIds       *[]string `tfsdk:"clock_ids" json:"clockIds,omitempty"`
				ContainerNames *[]string `tfsdk:"container_names" json:"containerNames,omitempty"`
				Duration       *string   `tfsdk:"duration" json:"duration,omitempty"`
				Mode           *string   `tfsdk:"mode" json:"mode,omitempty"`
				RemoteCluster  *string   `tfsdk:"remote_cluster" json:"remoteCluster,omitempty"`
				Selector       *struct {
					AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" json:"annotationSelectors,omitempty"`
					ExpressionSelectors *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"expression_selectors" json:"expressionSelectors,omitempty"`
					FieldSelectors    *map[string]string   `tfsdk:"field_selectors" json:"fieldSelectors,omitempty"`
					LabelSelectors    *map[string]string   `tfsdk:"label_selectors" json:"labelSelectors,omitempty"`
					Namespaces        *[]string            `tfsdk:"namespaces" json:"namespaces,omitempty"`
					NodeSelectors     *map[string]string   `tfsdk:"node_selectors" json:"nodeSelectors,omitempty"`
					Nodes             *[]string            `tfsdk:"nodes" json:"nodes,omitempty"`
					PodPhaseSelectors *[]string            `tfsdk:"pod_phase_selectors" json:"podPhaseSelectors,omitempty"`
					Pods              *map[string][]string `tfsdk:"pods" json:"pods,omitempty"`
				} `tfsdk:"selector" json:"selector,omitempty"`
				TimeOffset *string `tfsdk:"time_offset" json:"timeOffset,omitempty"`
				Value      *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"time_chaos" json:"timeChaos,omitempty"`
		} `tfsdk:"templates" json:"templates,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ChaosMeshOrgWorkflowV1Alpha1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_chaos_mesh_org_workflow_v1alpha1"
}

func (r *ChaosMeshOrgWorkflowV1Alpha1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"force_conflicts": schema.BoolAttribute{
				Description:         "If 'true', server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "If `true`, server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"field_manager": schema.BoolAttribute{
				Description:         "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"wait_for": schema.ListNestedAttribute{
				Description:         "Wait for specific conditions after create/update of resources.",
				MarkdownDescription: "Wait for specific conditions after create/update of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"jsonpath": schema.StringAttribute{
							Description:         "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							MarkdownDescription: "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							Required:            true,
							Optional:            false,
							Computed:            false,
						},
						"value": schema.StringAttribute{
							Description:         "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							MarkdownDescription: "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							Required:            false,
							Optional:            true,
							Computed:            true,
						},
						"timeout": schema.StringAttribute{
							Description:         "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							MarkdownDescription: "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("30s"),
						},
					},
				},
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},

					"labels": schema.MapAttribute{
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            true,
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
						Computed:            true,
						Validators: []validator.Map{
							validators.AnnotationValidator(),
						},
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "Spec defines the behavior of a workflow",
				MarkdownDescription: "Spec defines the behavior of a workflow",
				Attributes: map[string]schema.Attribute{
					"entry": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"templates": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"abort_with_status_check": schema.BoolAttribute{
									Description:         "AbortWithStatusCheck describe whether to abort the workflow when the failure threshold of StatusCheck is exceeded. Only used when Type is TypeStatusCheck.",
									MarkdownDescription: "AbortWithStatusCheck describe whether to abort the workflow when the failure threshold of StatusCheck is exceeded. Only used when Type is TypeStatusCheck.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"aws_chaos": schema.SingleNestedAttribute{
									Description:         "AWSChaosSpec is the content of the specification for an AWSChaos",
									MarkdownDescription: "AWSChaosSpec is the content of the specification for an AWSChaos",
									Attributes: map[string]schema.Attribute{
										"action": schema.StringAttribute{
											Description:         "Action defines the specific aws chaos action. Supported action: ec2-stop / ec2-restart / detach-volume Default action: ec2-stop",
											MarkdownDescription: "Action defines the specific aws chaos action. Supported action: ec2-stop / ec2-restart / detach-volume Default action: ec2-stop",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("ec2-stop", "ec2-restart", "detach-volume"),
											},
										},

										"aws_region": schema.StringAttribute{
											Description:         "AWSRegion defines the region of aws.",
											MarkdownDescription: "AWSRegion defines the region of aws.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"device_name": schema.StringAttribute{
											Description:         "DeviceName indicates the name of the device. Needed in detach-volume.",
											MarkdownDescription: "DeviceName indicates the name of the device. Needed in detach-volume.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"duration": schema.StringAttribute{
											Description:         "Duration represents the duration of the chaos action.",
											MarkdownDescription: "Duration represents the duration of the chaos action.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ec2_instance": schema.StringAttribute{
											Description:         "Ec2Instance indicates the ID of the ec2 instance.",
											MarkdownDescription: "Ec2Instance indicates the ID of the ec2 instance.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"endpoint": schema.StringAttribute{
											Description:         "Endpoint indicates the endpoint of the aws server. Just used it in test now.",
											MarkdownDescription: "Endpoint indicates the endpoint of the aws server. Just used it in test now.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"remote_cluster": schema.StringAttribute{
											Description:         "RemoteCluster represents the remote cluster where the chaos will be deployed",
											MarkdownDescription: "RemoteCluster represents the remote cluster where the chaos will be deployed",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secret_name": schema.StringAttribute{
											Description:         "SecretName defines the name of kubernetes secret.",
											MarkdownDescription: "SecretName defines the name of kubernetes secret.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"volume_id": schema.StringAttribute{
											Description:         "EbsVolume indicates the ID of the EBS volume. Needed in detach-volume.",
											MarkdownDescription: "EbsVolume indicates the ID of the EBS volume. Needed in detach-volume.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"azure_chaos": schema.SingleNestedAttribute{
									Description:         "AzureChaosSpec is the content of the specification for an AzureChaos",
									MarkdownDescription: "AzureChaosSpec is the content of the specification for an AzureChaos",
									Attributes: map[string]schema.Attribute{
										"action": schema.StringAttribute{
											Description:         "Action defines the specific azure chaos action. Supported action: vm-stop / vm-restart / disk-detach Default action: vm-stop",
											MarkdownDescription: "Action defines the specific azure chaos action. Supported action: vm-stop / vm-restart / disk-detach Default action: vm-stop",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("vm-stop", "vm-restart", "disk-detach"),
											},
										},

										"disk_name": schema.StringAttribute{
											Description:         "DiskName indicates the name of the disk. Needed in disk-detach.",
											MarkdownDescription: "DiskName indicates the name of the disk. Needed in disk-detach.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"duration": schema.StringAttribute{
											Description:         "Duration represents the duration of the chaos action.",
											MarkdownDescription: "Duration represents the duration of the chaos action.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"lun": schema.Int64Attribute{
											Description:         "LUN indicates the Logical Unit Number of the data disk. Needed in disk-detach.",
											MarkdownDescription: "LUN indicates the Logical Unit Number of the data disk. Needed in disk-detach.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"remote_cluster": schema.StringAttribute{
											Description:         "RemoteCluster represents the remote cluster where the chaos will be deployed",
											MarkdownDescription: "RemoteCluster represents the remote cluster where the chaos will be deployed",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"resource_group_name": schema.StringAttribute{
											Description:         "ResourceGroupName defines the name of ResourceGroup",
											MarkdownDescription: "ResourceGroupName defines the name of ResourceGroup",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"secret_name": schema.StringAttribute{
											Description:         "SecretName defines the name of kubernetes secret. It is used for Azure credentials.",
											MarkdownDescription: "SecretName defines the name of kubernetes secret. It is used for Azure credentials.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"subscription_id": schema.StringAttribute{
											Description:         "SubscriptionID defines the id of Azure subscription.",
											MarkdownDescription: "SubscriptionID defines the id of Azure subscription.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"vm_name": schema.StringAttribute{
											Description:         "VMName defines the name of Virtual Machine",
											MarkdownDescription: "VMName defines the name of Virtual Machine",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"block_chaos": schema.SingleNestedAttribute{
									Description:         "BlockChaosSpec is the content of the specification for a BlockChaos",
									MarkdownDescription: "BlockChaosSpec is the content of the specification for a BlockChaos",
									Attributes: map[string]schema.Attribute{
										"action": schema.StringAttribute{
											Description:         "Action defines the specific block chaos action. Supported action: delay",
											MarkdownDescription: "Action defines the specific block chaos action. Supported action: delay",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("delay"),
											},
										},

										"container_names": schema.ListAttribute{
											Description:         "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
											MarkdownDescription: "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"delay": schema.SingleNestedAttribute{
											Description:         "Delay defines the delay distribution.",
											MarkdownDescription: "Delay defines the delay distribution.",
											Attributes: map[string]schema.Attribute{
												"correlation": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"jitter": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"latency": schema.StringAttribute{
													Description:         "Latency defines the latency of every io request.",
													MarkdownDescription: "Latency defines the latency of every io request.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"duration": schema.StringAttribute{
											Description:         "Duration represents the duration of the chaos action.",
											MarkdownDescription: "Duration represents the duration of the chaos action.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"mode": schema.StringAttribute{
											Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
											MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
											},
										},

										"remote_cluster": schema.StringAttribute{
											Description:         "RemoteCluster represents the remote cluster where the chaos will be deployed",
											MarkdownDescription: "RemoteCluster represents the remote cluster where the chaos will be deployed",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"selector": schema.SingleNestedAttribute{
											Description:         "Selector is used to select pods that are used to inject chaos action.",
											MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",
											Attributes: map[string]schema.Attribute{
												"annotation_selectors": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
													MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"expression_selectors": schema.ListNestedAttribute{
													Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
													MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
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

												"field_selectors": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
													MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"label_selectors": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
													MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespaces": schema.ListAttribute{
													Description:         "Namespaces is a set of namespace to which objects belong.",
													MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"node_selectors": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
													MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"nodes": schema.ListAttribute{
													Description:         "Nodes is a set of node name and objects must belong to these nodes.",
													MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"pod_phase_selectors": schema.ListAttribute{
													Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
													MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"pods": schema.MapAttribute{
													Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
													MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
													ElementType:         types.ListType{ElemType: types.StringType},
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"value": schema.StringAttribute{
											Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
											MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"volume_name": schema.StringAttribute{
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

								"children": schema.ListAttribute{
									Description:         "Children describes the children steps of serial or parallel node. Only used when Type is TypeSerial or TypeParallel.",
									MarkdownDescription: "Children describes the children steps of serial or parallel node. Only used when Type is TypeSerial or TypeParallel.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"conditional_branches": schema.ListNestedAttribute{
									Description:         "ConditionalBranches describes the conditional branches of custom tasks. Only used when Type is TypeTask.",
									MarkdownDescription: "ConditionalBranches describes the conditional branches of custom tasks. Only used when Type is TypeTask.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"expression": schema.StringAttribute{
												Description:         "Expression is the expression for this conditional branch, expected type of result is boolean. If expression is empty, this branch will always be selected/the template will be spawned.",
												MarkdownDescription: "Expression is the expression for this conditional branch, expected type of result is boolean. If expression is empty, this branch will always be selected/the template will be spawned.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"target": schema.StringAttribute{
												Description:         "Target is the name of other template, if expression is evaluated as true, this template will be spawned.",
												MarkdownDescription: "Target is the name of other template, if expression is evaluated as true, this template will be spawned.",
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

								"deadline": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"dns_chaos": schema.SingleNestedAttribute{
									Description:         "DNSChaosSpec defines the desired state of DNSChaos",
									MarkdownDescription: "DNSChaosSpec defines the desired state of DNSChaos",
									Attributes: map[string]schema.Attribute{
										"action": schema.StringAttribute{
											Description:         "Action defines the specific DNS chaos action. Supported action: error, random Default action: error",
											MarkdownDescription: "Action defines the specific DNS chaos action. Supported action: error, random Default action: error",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("error", "random"),
											},
										},

										"container_names": schema.ListAttribute{
											Description:         "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
											MarkdownDescription: "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"duration": schema.StringAttribute{
											Description:         "Duration represents the duration of the chaos action",
											MarkdownDescription: "Duration represents the duration of the chaos action",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"mode": schema.StringAttribute{
											Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
											MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
											},
										},

										"patterns": schema.ListAttribute{
											Description:         "Choose which domain names to take effect, support the placeholder ? and wildcard *, or the Specified domain name. Note:      1. The wildcard * must be at the end of the string. For example, chaos-*.org is invalid.      2. if the patterns is empty, will take effect on all the domain names. For example: 		The value is ['google.com', 'github.*', 'chaos-mes?.org'], 		will take effect on 'google.com', 'github.com' and 'chaos-mesh.org'",
											MarkdownDescription: "Choose which domain names to take effect, support the placeholder ? and wildcard *, or the Specified domain name. Note:      1. The wildcard * must be at the end of the string. For example, chaos-*.org is invalid.      2. if the patterns is empty, will take effect on all the domain names. For example: 		The value is ['google.com', 'github.*', 'chaos-mes?.org'], 		will take effect on 'google.com', 'github.com' and 'chaos-mesh.org'",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"remote_cluster": schema.StringAttribute{
											Description:         "RemoteCluster represents the remote cluster where the chaos will be deployed",
											MarkdownDescription: "RemoteCluster represents the remote cluster where the chaos will be deployed",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"selector": schema.SingleNestedAttribute{
											Description:         "Selector is used to select pods that are used to inject chaos action.",
											MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",
											Attributes: map[string]schema.Attribute{
												"annotation_selectors": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
													MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"expression_selectors": schema.ListNestedAttribute{
													Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
													MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
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

												"field_selectors": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
													MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"label_selectors": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
													MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespaces": schema.ListAttribute{
													Description:         "Namespaces is a set of namespace to which objects belong.",
													MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"node_selectors": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
													MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"nodes": schema.ListAttribute{
													Description:         "Nodes is a set of node name and objects must belong to these nodes.",
													MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"pod_phase_selectors": schema.ListAttribute{
													Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
													MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"pods": schema.MapAttribute{
													Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
													MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
													ElementType:         types.ListType{ElemType: types.StringType},
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"value": schema.StringAttribute{
											Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
											MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"gcp_chaos": schema.SingleNestedAttribute{
									Description:         "GCPChaosSpec is the content of the specification for a GCPChaos",
									MarkdownDescription: "GCPChaosSpec is the content of the specification for a GCPChaos",
									Attributes: map[string]schema.Attribute{
										"action": schema.StringAttribute{
											Description:         "Action defines the specific gcp chaos action. Supported action: node-stop / node-reset / disk-loss Default action: node-stop",
											MarkdownDescription: "Action defines the specific gcp chaos action. Supported action: node-stop / node-reset / disk-loss Default action: node-stop",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("node-stop", "node-reset", "disk-loss"),
											},
										},

										"device_names": schema.ListAttribute{
											Description:         "The device name of disks to detach. Needed in disk-loss.",
											MarkdownDescription: "The device name of disks to detach. Needed in disk-loss.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"duration": schema.StringAttribute{
											Description:         "Duration represents the duration of the chaos action.",
											MarkdownDescription: "Duration represents the duration of the chaos action.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"instance": schema.StringAttribute{
											Description:         "Instance defines the name of the instance",
											MarkdownDescription: "Instance defines the name of the instance",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"project": schema.StringAttribute{
											Description:         "Project defines the ID of gcp project.",
											MarkdownDescription: "Project defines the ID of gcp project.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"remote_cluster": schema.StringAttribute{
											Description:         "RemoteCluster represents the remote cluster where the chaos will be deployed",
											MarkdownDescription: "RemoteCluster represents the remote cluster where the chaos will be deployed",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secret_name": schema.StringAttribute{
											Description:         "SecretName defines the name of kubernetes secret. It is used for GCP credentials.",
											MarkdownDescription: "SecretName defines the name of kubernetes secret. It is used for GCP credentials.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"zone": schema.StringAttribute{
											Description:         "Zone defines the zone of gcp project.",
											MarkdownDescription: "Zone defines the zone of gcp project.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"http_chaos": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"abort": schema.BoolAttribute{
											Description:         "Abort is a rule to abort a http session.",
											MarkdownDescription: "Abort is a rule to abort a http session.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"code": schema.Int64Attribute{
											Description:         "Code is a rule to select target by http status code in response.",
											MarkdownDescription: "Code is a rule to select target by http status code in response.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"delay": schema.StringAttribute{
											Description:         "Delay represents the delay of the target request/response. A duration string is a possibly unsigned sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
											MarkdownDescription: "Delay represents the delay of the target request/response. A duration string is a possibly unsigned sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"duration": schema.StringAttribute{
											Description:         "Duration represents the duration of the chaos action.",
											MarkdownDescription: "Duration represents the duration of the chaos action.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"method": schema.StringAttribute{
											Description:         "Method is a rule to select target by http method in request.",
											MarkdownDescription: "Method is a rule to select target by http method in request.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"mode": schema.StringAttribute{
											Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
											MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
											},
										},

										"patch": schema.SingleNestedAttribute{
											Description:         "Patch is a rule to patch some contents in target.",
											MarkdownDescription: "Patch is a rule to patch some contents in target.",
											Attributes: map[string]schema.Attribute{
												"body": schema.SingleNestedAttribute{
													Description:         "Body is a rule to patch message body of target.",
													MarkdownDescription: "Body is a rule to patch message body of target.",
													Attributes: map[string]schema.Attribute{
														"type": schema.StringAttribute{
															Description:         "Type represents the patch type, only support 'JSON' as [merge patch json](https://tools.ietf.org/html/rfc7396) currently.",
															MarkdownDescription: "Type represents the patch type, only support 'JSON' as [merge patch json](https://tools.ietf.org/html/rfc7396) currently.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"value": schema.StringAttribute{
															Description:         "Value is the patch contents.",
															MarkdownDescription: "Value is the patch contents.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"headers": schema.ListAttribute{
													Description:         "Headers is a rule to append http headers of target. For example: '[['Set-Cookie', '<one cookie>'], ['Set-Cookie', '<another cookie>']]'.",
													MarkdownDescription: "Headers is a rule to append http headers of target. For example: '[['Set-Cookie', '<one cookie>'], ['Set-Cookie', '<another cookie>']]'.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"queries": schema.ListAttribute{
													Description:         "Queries is a rule to append uri queries of target(Request only). For example: '[['foo', 'bar'], ['foo', 'unknown']]'.",
													MarkdownDescription: "Queries is a rule to append uri queries of target(Request only). For example: '[['foo', 'bar'], ['foo', 'unknown']]'.",
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

										"path": schema.StringAttribute{
											Description:         "Path is a rule to select target by uri path in http request.",
											MarkdownDescription: "Path is a rule to select target by uri path in http request.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"port": schema.Int64Attribute{
											Description:         "Port represents the target port to be proxy of.",
											MarkdownDescription: "Port represents the target port to be proxy of.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"remote_cluster": schema.StringAttribute{
											Description:         "RemoteCluster represents the remote cluster where the chaos will be deployed",
											MarkdownDescription: "RemoteCluster represents the remote cluster where the chaos will be deployed",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"replace": schema.SingleNestedAttribute{
											Description:         "Replace is a rule to replace some contents in target.",
											MarkdownDescription: "Replace is a rule to replace some contents in target.",
											Attributes: map[string]schema.Attribute{
												"body": schema.StringAttribute{
													Description:         "Body is a rule to replace http message body in target.",
													MarkdownDescription: "Body is a rule to replace http message body in target.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														validators.Base64Validator(),
													},
												},

												"code": schema.Int64Attribute{
													Description:         "Code is a rule to replace http status code in response.",
													MarkdownDescription: "Code is a rule to replace http status code in response.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"headers": schema.MapAttribute{
													Description:         "Headers is a rule to replace http headers of target. The key-value pairs represent header name and header value pairs.",
													MarkdownDescription: "Headers is a rule to replace http headers of target. The key-value pairs represent header name and header value pairs.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"method": schema.StringAttribute{
													Description:         "Method is a rule to replace http method in request.",
													MarkdownDescription: "Method is a rule to replace http method in request.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"path": schema.StringAttribute{
													Description:         "Path is rule to to replace uri path in http request.",
													MarkdownDescription: "Path is rule to to replace uri path in http request.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"queries": schema.MapAttribute{
													Description:         "Queries is a rule to replace uri queries in http request. For example, with value '{ 'foo': 'unknown' }', the '/?foo=bar' will be altered to '/?foo=unknown',",
													MarkdownDescription: "Queries is a rule to replace uri queries in http request. For example, with value '{ 'foo': 'unknown' }', the '/?foo=bar' will be altered to '/?foo=unknown',",
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

										"request_headers": schema.MapAttribute{
											Description:         "RequestHeaders is a rule to select target by http headers in request. The key-value pairs represent header name and header value pairs.",
											MarkdownDescription: "RequestHeaders is a rule to select target by http headers in request. The key-value pairs represent header name and header value pairs.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"response_headers": schema.MapAttribute{
											Description:         "ResponseHeaders is a rule to select target by http headers in response. The key-value pairs represent header name and header value pairs.",
											MarkdownDescription: "ResponseHeaders is a rule to select target by http headers in response. The key-value pairs represent header name and header value pairs.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"selector": schema.SingleNestedAttribute{
											Description:         "Selector is used to select pods that are used to inject chaos action.",
											MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",
											Attributes: map[string]schema.Attribute{
												"annotation_selectors": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
													MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"expression_selectors": schema.ListNestedAttribute{
													Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
													MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
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

												"field_selectors": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
													MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"label_selectors": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
													MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespaces": schema.ListAttribute{
													Description:         "Namespaces is a set of namespace to which objects belong.",
													MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"node_selectors": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
													MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"nodes": schema.ListAttribute{
													Description:         "Nodes is a set of node name and objects must belong to these nodes.",
													MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"pod_phase_selectors": schema.ListAttribute{
													Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
													MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"pods": schema.MapAttribute{
													Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
													MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
													ElementType:         types.ListType{ElemType: types.StringType},
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"target": schema.StringAttribute{
											Description:         "Target is the object to be selected and injected.",
											MarkdownDescription: "Target is the object to be selected and injected.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Request", "Response"),
											},
										},

										"tls": schema.SingleNestedAttribute{
											Description:         "TLS is the tls config, will override PodHttpChaos if there are multiple HTTPChaos experiments are applied",
											MarkdownDescription: "TLS is the tls config, will override PodHttpChaos if there are multiple HTTPChaos experiments are applied",
											Attributes: map[string]schema.Attribute{
												"ca_name": schema.StringAttribute{
													Description:         "CAName represents the data name of ca file in secret, 'ca.crt' for example",
													MarkdownDescription: "CAName represents the data name of ca file in secret, 'ca.crt' for example",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"cert_name": schema.StringAttribute{
													Description:         "CertName represents the data name of cert file in secret, 'tls.crt' for example",
													MarkdownDescription: "CertName represents the data name of cert file in secret, 'tls.crt' for example",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"key_name": schema.StringAttribute{
													Description:         "KeyName represents the data name of key file in secret, 'tls.key' for example",
													MarkdownDescription: "KeyName represents the data name of key file in secret, 'tls.key' for example",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"secret_name": schema.StringAttribute{
													Description:         "SecretName represents the name of required secret resource",
													MarkdownDescription: "SecretName represents the name of required secret resource",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"secret_namespace": schema.StringAttribute{
													Description:         "SecretNamespace represents the namespace of required secret resource",
													MarkdownDescription: "SecretNamespace represents the namespace of required secret resource",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"value": schema.StringAttribute{
											Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
											MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"io_chaos": schema.SingleNestedAttribute{
									Description:         "IOChaosSpec defines the desired state of IOChaos",
									MarkdownDescription: "IOChaosSpec defines the desired state of IOChaos",
									Attributes: map[string]schema.Attribute{
										"action": schema.StringAttribute{
											Description:         "Action defines the specific pod chaos action. Supported action: latency / fault / attrOverride / mistake",
											MarkdownDescription: "Action defines the specific pod chaos action. Supported action: latency / fault / attrOverride / mistake",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("latency", "fault", "attrOverride", "mistake"),
											},
										},

										"attr": schema.SingleNestedAttribute{
											Description:         "Attr defines the overrided attribution",
											MarkdownDescription: "Attr defines the overrided attribution",
											Attributes: map[string]schema.Attribute{
												"atime": schema.SingleNestedAttribute{
													Description:         "Timespec represents a time",
													MarkdownDescription: "Timespec represents a time",
													Attributes: map[string]schema.Attribute{
														"nsec": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"sec": schema.Int64Attribute{
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

												"blocks": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"ctime": schema.SingleNestedAttribute{
													Description:         "Timespec represents a time",
													MarkdownDescription: "Timespec represents a time",
													Attributes: map[string]schema.Attribute{
														"nsec": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"sec": schema.Int64Attribute{
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

												"gid": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"ino": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"kind": schema.StringAttribute{
													Description:         "FileType represents type of file",
													MarkdownDescription: "FileType represents type of file",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"mtime": schema.SingleNestedAttribute{
													Description:         "Timespec represents a time",
													MarkdownDescription: "Timespec represents a time",
													Attributes: map[string]schema.Attribute{
														"nsec": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"sec": schema.Int64Attribute{
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

												"nlink": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"perm": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"rdev": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"size": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uid": schema.Int64Attribute{
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

										"container_names": schema.ListAttribute{
											Description:         "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
											MarkdownDescription: "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"delay": schema.StringAttribute{
											Description:         "Delay defines the value of I/O chaos action delay. A delay string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
											MarkdownDescription: "Delay defines the value of I/O chaos action delay. A delay string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"duration": schema.StringAttribute{
											Description:         "Duration represents the duration of the chaos action. It is required when the action is 'PodFailureAction'. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '-1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
											MarkdownDescription: "Duration represents the duration of the chaos action. It is required when the action is 'PodFailureAction'. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '-1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"errno": schema.Int64Attribute{
											Description:         "Errno defines the error code that returned by I/O action. refer to: https://www-numi.fnal.gov/offline_software/srt_public_context/WebDocs/Errors/unix_system_errors.html",
											MarkdownDescription: "Errno defines the error code that returned by I/O action. refer to: https://www-numi.fnal.gov/offline_software/srt_public_context/WebDocs/Errors/unix_system_errors.html",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"methods": schema.ListAttribute{
											Description:         "Methods defines the I/O methods for injecting I/O chaos action. default: all I/O methods.",
											MarkdownDescription: "Methods defines the I/O methods for injecting I/O chaos action. default: all I/O methods.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"mistake": schema.SingleNestedAttribute{
											Description:         "Mistake defines what types of incorrectness are injected to IO operations",
											MarkdownDescription: "Mistake defines what types of incorrectness are injected to IO operations",
											Attributes: map[string]schema.Attribute{
												"filling": schema.StringAttribute{
													Description:         "Filling determines what is filled in the mistake data.",
													MarkdownDescription: "Filling determines what is filled in the mistake data.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("zero", "random"),
													},
												},

												"max_length": schema.Int64Attribute{
													Description:         "Max length of each wrong data segment in bytes",
													MarkdownDescription: "Max length of each wrong data segment in bytes",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(1),
													},
												},

												"max_occurrences": schema.Int64Attribute{
													Description:         "There will be [1, MaxOccurrences] segments of wrong data.",
													MarkdownDescription: "There will be [1, MaxOccurrences] segments of wrong data.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(1),
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"mode": schema.StringAttribute{
											Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
											MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
											},
										},

										"path": schema.StringAttribute{
											Description:         "Path defines the path of files for injecting I/O chaos action.",
											MarkdownDescription: "Path defines the path of files for injecting I/O chaos action.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"percent": schema.Int64Attribute{
											Description:         "Percent defines the percentage of injection errors and provides a number from 0-100. default: 100.",
											MarkdownDescription: "Percent defines the percentage of injection errors and provides a number from 0-100. default: 100.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"remote_cluster": schema.StringAttribute{
											Description:         "RemoteCluster represents the remote cluster where the chaos will be deployed",
											MarkdownDescription: "RemoteCluster represents the remote cluster where the chaos will be deployed",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"selector": schema.SingleNestedAttribute{
											Description:         "Selector is used to select pods that are used to inject chaos action.",
											MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",
											Attributes: map[string]schema.Attribute{
												"annotation_selectors": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
													MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"expression_selectors": schema.ListNestedAttribute{
													Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
													MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
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

												"field_selectors": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
													MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"label_selectors": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
													MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespaces": schema.ListAttribute{
													Description:         "Namespaces is a set of namespace to which objects belong.",
													MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"node_selectors": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
													MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"nodes": schema.ListAttribute{
													Description:         "Nodes is a set of node name and objects must belong to these nodes.",
													MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"pod_phase_selectors": schema.ListAttribute{
													Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
													MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"pods": schema.MapAttribute{
													Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
													MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
													ElementType:         types.ListType{ElemType: types.StringType},
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"value": schema.StringAttribute{
											Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
											MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"volume_path": schema.StringAttribute{
											Description:         "VolumePath represents the mount path of injected volume",
											MarkdownDescription: "VolumePath represents the mount path of injected volume",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"jvm_chaos": schema.SingleNestedAttribute{
									Description:         "JVMChaosSpec defines the desired state of JVMChaos",
									MarkdownDescription: "JVMChaosSpec defines the desired state of JVMChaos",
									Attributes: map[string]schema.Attribute{
										"action": schema.StringAttribute{
											Description:         "Action defines the specific jvm chaos action. Supported action: latency;return;exception;stress;gc;ruleData",
											MarkdownDescription: "Action defines the specific jvm chaos action. Supported action: latency;return;exception;stress;gc;ruleData",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("latency", "return", "exception", "stress", "gc", "ruleData", "mysql"),
											},
										},

										"class": schema.StringAttribute{
											Description:         "Java class",
											MarkdownDescription: "Java class",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"container_names": schema.ListAttribute{
											Description:         "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
											MarkdownDescription: "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"cpu_count": schema.Int64Attribute{
											Description:         "the CPU core number needs to use, only set it when action is stress",
											MarkdownDescription: "the CPU core number needs to use, only set it when action is stress",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"database": schema.StringAttribute{
											Description:         "the match database default value is '', means match all database",
											MarkdownDescription: "the match database default value is '', means match all database",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"duration": schema.StringAttribute{
											Description:         "Duration represents the duration of the chaos action",
											MarkdownDescription: "Duration represents the duration of the chaos action",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exception": schema.StringAttribute{
											Description:         "the exception which needs to throw for action 'exception' or the exception message needs to throw in action 'mysql'",
											MarkdownDescription: "the exception which needs to throw for action 'exception' or the exception message needs to throw in action 'mysql'",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"latency": schema.Int64Attribute{
											Description:         "the latency duration for action 'latency', unit ms or the latency duration in action 'mysql'",
											MarkdownDescription: "the latency duration for action 'latency', unit ms or the latency duration in action 'mysql'",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"mem_type": schema.StringAttribute{
											Description:         "the memory type needs to locate, only set it when action is stress, the value can be 'stack' or 'heap'",
											MarkdownDescription: "the memory type needs to locate, only set it when action is stress, the value can be 'stack' or 'heap'",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"method": schema.StringAttribute{
											Description:         "the method in Java class",
											MarkdownDescription: "the method in Java class",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"mode": schema.StringAttribute{
											Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
											MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
											},
										},

										"mysql_connector_version": schema.StringAttribute{
											Description:         "the version of mysql-connector-java, only support 5.X.X(set to '5') and 8.X.X(set to '8') now",
											MarkdownDescription: "the version of mysql-connector-java, only support 5.X.X(set to '5') and 8.X.X(set to '8') now",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "byteman rule name, should be unique, and will generate one if not set",
											MarkdownDescription: "byteman rule name, should be unique, and will generate one if not set",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"pid": schema.Int64Attribute{
											Description:         "the pid of Java process which needs to attach",
											MarkdownDescription: "the pid of Java process which needs to attach",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"port": schema.Int64Attribute{
											Description:         "the port of agent server, default 9277",
											MarkdownDescription: "the port of agent server, default 9277",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"remote_cluster": schema.StringAttribute{
											Description:         "RemoteCluster represents the remote cluster where the chaos will be deployed",
											MarkdownDescription: "RemoteCluster represents the remote cluster where the chaos will be deployed",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"rule_data": schema.StringAttribute{
											Description:         "the byteman rule's data for action 'ruleData'",
											MarkdownDescription: "the byteman rule's data for action 'ruleData'",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"selector": schema.SingleNestedAttribute{
											Description:         "Selector is used to select pods that are used to inject chaos action.",
											MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",
											Attributes: map[string]schema.Attribute{
												"annotation_selectors": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
													MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"expression_selectors": schema.ListNestedAttribute{
													Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
													MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
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

												"field_selectors": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
													MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"label_selectors": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
													MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespaces": schema.ListAttribute{
													Description:         "Namespaces is a set of namespace to which objects belong.",
													MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"node_selectors": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
													MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"nodes": schema.ListAttribute{
													Description:         "Nodes is a set of node name and objects must belong to these nodes.",
													MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"pod_phase_selectors": schema.ListAttribute{
													Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
													MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"pods": schema.MapAttribute{
													Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
													MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
													ElementType:         types.ListType{ElemType: types.StringType},
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"sql_type": schema.StringAttribute{
											Description:         "the match sql type default value is '', means match all SQL type. The value can be 'select', 'insert', 'update', 'delete', 'replace'.",
											MarkdownDescription: "the match sql type default value is '', means match all SQL type. The value can be 'select', 'insert', 'update', 'delete', 'replace'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"table": schema.StringAttribute{
											Description:         "the match table default value is '', means match all table",
											MarkdownDescription: "the match table default value is '', means match all table",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
											MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"kernel_chaos": schema.SingleNestedAttribute{
									Description:         "KernelChaosSpec defines the desired state of KernelChaos",
									MarkdownDescription: "KernelChaosSpec defines the desired state of KernelChaos",
									Attributes: map[string]schema.Attribute{
										"container_names": schema.ListAttribute{
											Description:         "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
											MarkdownDescription: "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"duration": schema.StringAttribute{
											Description:         "Duration represents the duration of the chaos action",
											MarkdownDescription: "Duration represents the duration of the chaos action",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"fail_kern_request": schema.SingleNestedAttribute{
											Description:         "FailKernRequest defines the request of kernel injection",
											MarkdownDescription: "FailKernRequest defines the request of kernel injection",
											Attributes: map[string]schema.Attribute{
												"callchain": schema.ListNestedAttribute{
													Description:         "Callchain indicate a special call chain, such as:     ext4_mount       -> mount_subtree          -> ...             -> should_failslab With an optional set of predicates and an optional set of parameters, which used with predicates. You can read call chan and predicate examples from https://github.com/chaos-mesh/bpfki/tree/develop/examples to learn more. If no special call chain, just keep Callchain empty, which means it will fail at any call chain with slab alloc (eg: kmalloc).",
													MarkdownDescription: "Callchain indicate a special call chain, such as:     ext4_mount       -> mount_subtree          -> ...             -> should_failslab With an optional set of predicates and an optional set of parameters, which used with predicates. You can read call chan and predicate examples from https://github.com/chaos-mesh/bpfki/tree/develop/examples to learn more. If no special call chain, just keep Callchain empty, which means it will fail at any call chain with slab alloc (eg: kmalloc).",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"funcname": schema.StringAttribute{
																Description:         "Funcname can be find from kernel source or '/proc/kallsyms', such as 'ext4_mount'",
																MarkdownDescription: "Funcname can be find from kernel source or '/proc/kallsyms', such as 'ext4_mount'",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"parameters": schema.StringAttribute{
																Description:         "Parameters is used with predicate, for example, if you want to inject slab error in 'd_alloc_parallel(struct dentry *parent, const struct qstr *name)' with a special name 'bananas', you need to set it to 'struct dentry *parent, const struct qstr *name' otherwise omit it.",
																MarkdownDescription: "Parameters is used with predicate, for example, if you want to inject slab error in 'd_alloc_parallel(struct dentry *parent, const struct qstr *name)' with a special name 'bananas', you need to set it to 'struct dentry *parent, const struct qstr *name' otherwise omit it.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"predicate": schema.StringAttribute{
																Description:         "Predicate will access the arguments of this Frame, example with Parameters's, you can set it to 'STRNCMP(name->name, 'bananas', 8)' to make inject only with it, or omit it to inject for all d_alloc_parallel call chain.",
																MarkdownDescription: "Predicate will access the arguments of this Frame, example with Parameters's, you can set it to 'STRNCMP(name->name, 'bananas', 8)' to make inject only with it, or omit it to inject for all d_alloc_parallel call chain.",
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

												"failtype": schema.Int64Attribute{
													Description:         "FailType indicates what to fail, can be set to '0' / '1' / '2' If '0', indicates slab to fail (should_failslab) If '1', indicates alloc_page to fail (should_fail_alloc_page) If '2', indicates bio to fail (should_fail_bio) You can read:   1. https://www.kernel.org/doc/html/latest/fault-injection/fault-injection.html   2. http://github.com/iovisor/bcc/blob/master/tools/inject_example.txt to learn more",
													MarkdownDescription: "FailType indicates what to fail, can be set to '0' / '1' / '2' If '0', indicates slab to fail (should_failslab) If '1', indicates alloc_page to fail (should_fail_alloc_page) If '2', indicates bio to fail (should_fail_bio) You can read:   1. https://www.kernel.org/doc/html/latest/fault-injection/fault-injection.html   2. http://github.com/iovisor/bcc/blob/master/tools/inject_example.txt to learn more",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(0),
														int64validator.AtMost(2),
													},
												},

												"headers": schema.ListAttribute{
													Description:         "Headers indicates the appropriate kernel headers you need. Eg: 'linux/mmzone.h', 'linux/blkdev.h' and so on",
													MarkdownDescription: "Headers indicates the appropriate kernel headers you need. Eg: 'linux/mmzone.h', 'linux/blkdev.h' and so on",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"probability": schema.Int64Attribute{
													Description:         "Probability indicates the fails with probability. If you want 1%, please set this field with 1.",
													MarkdownDescription: "Probability indicates the fails with probability. If you want 1%, please set this field with 1.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(0),
														int64validator.AtMost(100),
													},
												},

												"times": schema.Int64Attribute{
													Description:         "Times indicates the max times of fails.",
													MarkdownDescription: "Times indicates the max times of fails.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(0),
													},
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"mode": schema.StringAttribute{
											Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
											MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
											},
										},

										"remote_cluster": schema.StringAttribute{
											Description:         "RemoteCluster represents the remote cluster where the chaos will be deployed",
											MarkdownDescription: "RemoteCluster represents the remote cluster where the chaos will be deployed",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"selector": schema.SingleNestedAttribute{
											Description:         "Selector is used to select pods that are used to inject chaos action.",
											MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",
											Attributes: map[string]schema.Attribute{
												"annotation_selectors": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
													MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"expression_selectors": schema.ListNestedAttribute{
													Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
													MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
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

												"field_selectors": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
													MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"label_selectors": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
													MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespaces": schema.ListAttribute{
													Description:         "Namespaces is a set of namespace to which objects belong.",
													MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"node_selectors": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
													MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"nodes": schema.ListAttribute{
													Description:         "Nodes is a set of node name and objects must belong to these nodes.",
													MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"pod_phase_selectors": schema.ListAttribute{
													Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
													MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"pods": schema.MapAttribute{
													Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
													MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
													ElementType:         types.ListType{ElemType: types.StringType},
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"value": schema.StringAttribute{
											Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
											MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"network_chaos": schema.SingleNestedAttribute{
									Description:         "NetworkChaosSpec defines the desired state of NetworkChaos",
									MarkdownDescription: "NetworkChaosSpec defines the desired state of NetworkChaos",
									Attributes: map[string]schema.Attribute{
										"action": schema.StringAttribute{
											Description:         "Action defines the specific network chaos action. Supported action: partition, netem, delay, loss, duplicate, corrupt Default action: delay",
											MarkdownDescription: "Action defines the specific network chaos action. Supported action: partition, netem, delay, loss, duplicate, corrupt Default action: delay",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("netem", "delay", "loss", "duplicate", "corrupt", "partition", "bandwidth"),
											},
										},

										"bandwidth": schema.SingleNestedAttribute{
											Description:         "Bandwidth represents the detail about bandwidth control action",
											MarkdownDescription: "Bandwidth represents the detail about bandwidth control action",
											Attributes: map[string]schema.Attribute{
												"buffer": schema.Int64Attribute{
													Description:         "Buffer is the maximum amount of bytes that tokens can be available for instantaneously.",
													MarkdownDescription: "Buffer is the maximum amount of bytes that tokens can be available for instantaneously.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(1),
													},
												},

												"limit": schema.Int64Attribute{
													Description:         "Limit is the number of bytes that can be queued waiting for tokens to become available.",
													MarkdownDescription: "Limit is the number of bytes that can be queued waiting for tokens to become available.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(1),
													},
												},

												"minburst": schema.Int64Attribute{
													Description:         "Minburst specifies the size of the peakrate bucket. For perfect accuracy, should be set to the MTU of the interface.  If a peakrate is needed, but some burstiness is acceptable, this size can be raised. A 3000 byte minburst allows around 3mbit/s of peakrate, given 1000 byte packets.",
													MarkdownDescription: "Minburst specifies the size of the peakrate bucket. For perfect accuracy, should be set to the MTU of the interface.  If a peakrate is needed, but some burstiness is acceptable, this size can be raised. A 3000 byte minburst allows around 3mbit/s of peakrate, given 1000 byte packets.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(0),
													},
												},

												"peakrate": schema.Int64Attribute{
													Description:         "Peakrate is the maximum depletion rate of the bucket. The peakrate does not need to be set, it is only necessary if perfect millisecond timescale shaping is required.",
													MarkdownDescription: "Peakrate is the maximum depletion rate of the bucket. The peakrate does not need to be set, it is only necessary if perfect millisecond timescale shaping is required.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(0),
													},
												},

												"rate": schema.StringAttribute{
													Description:         "Rate is the speed knob. Allows bit, kbit, mbit, gbit, tbit, bps, kbps, mbps, gbps, tbps unit. bps means bytes per second.",
													MarkdownDescription: "Rate is the speed knob. Allows bit, kbit, mbit, gbit, tbit, bps, kbps, mbps, gbps, tbps unit. bps means bytes per second.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"corrupt": schema.SingleNestedAttribute{
											Description:         "Corrupt represents the detail about corrupt action",
											MarkdownDescription: "Corrupt represents the detail about corrupt action",
											Attributes: map[string]schema.Attribute{
												"correlation": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"corrupt": schema.StringAttribute{
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

										"delay": schema.SingleNestedAttribute{
											Description:         "Delay represents the detail about delay action",
											MarkdownDescription: "Delay represents the detail about delay action",
											Attributes: map[string]schema.Attribute{
												"correlation": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"jitter": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"latency": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"reorder": schema.SingleNestedAttribute{
													Description:         "ReorderSpec defines details of packet reorder.",
													MarkdownDescription: "ReorderSpec defines details of packet reorder.",
													Attributes: map[string]schema.Attribute{
														"correlation": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"gap": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"reorder": schema.StringAttribute{
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"device": schema.StringAttribute{
											Description:         "Device represents the network device to be affected.",
											MarkdownDescription: "Device represents the network device to be affected.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"direction": schema.StringAttribute{
											Description:         "Direction represents the direction, this applies on netem and network partition action",
											MarkdownDescription: "Direction represents the direction, this applies on netem and network partition action",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("to", "from", "both"),
											},
										},

										"duplicate": schema.SingleNestedAttribute{
											Description:         "DuplicateSpec represents the detail about loss action",
											MarkdownDescription: "DuplicateSpec represents the detail about loss action",
											Attributes: map[string]schema.Attribute{
												"correlation": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"duplicate": schema.StringAttribute{
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

										"duration": schema.StringAttribute{
											Description:         "Duration represents the duration of the chaos action",
											MarkdownDescription: "Duration represents the duration of the chaos action",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"external_targets": schema.ListAttribute{
											Description:         "ExternalTargets represents network targets outside k8s",
											MarkdownDescription: "ExternalTargets represents network targets outside k8s",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"loss": schema.SingleNestedAttribute{
											Description:         "Loss represents the detail about loss action",
											MarkdownDescription: "Loss represents the detail about loss action",
											Attributes: map[string]schema.Attribute{
												"correlation": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"loss": schema.StringAttribute{
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

										"mode": schema.StringAttribute{
											Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
											MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
											},
										},

										"remote_cluster": schema.StringAttribute{
											Description:         "RemoteCluster represents the remote cluster where the chaos will be deployed",
											MarkdownDescription: "RemoteCluster represents the remote cluster where the chaos will be deployed",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"selector": schema.SingleNestedAttribute{
											Description:         "Selector is used to select pods that are used to inject chaos action.",
											MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",
											Attributes: map[string]schema.Attribute{
												"annotation_selectors": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
													MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"expression_selectors": schema.ListNestedAttribute{
													Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
													MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
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

												"field_selectors": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
													MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"label_selectors": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
													MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespaces": schema.ListAttribute{
													Description:         "Namespaces is a set of namespace to which objects belong.",
													MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"node_selectors": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
													MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"nodes": schema.ListAttribute{
													Description:         "Nodes is a set of node name and objects must belong to these nodes.",
													MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"pod_phase_selectors": schema.ListAttribute{
													Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
													MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"pods": schema.MapAttribute{
													Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
													MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
													ElementType:         types.ListType{ElemType: types.StringType},
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"target": schema.SingleNestedAttribute{
											Description:         "Target represents network target, this applies on netem and network partition action",
											MarkdownDescription: "Target represents network target, this applies on netem and network partition action",
											Attributes: map[string]schema.Attribute{
												"mode": schema.StringAttribute{
													Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
													MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
													},
												},

												"selector": schema.SingleNestedAttribute{
													Description:         "Selector is used to select pods that are used to inject chaos action.",
													MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",
													Attributes: map[string]schema.Attribute{
														"annotation_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
															MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"expression_selectors": schema.ListNestedAttribute{
															Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
															MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
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

														"field_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
															MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"label_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
															MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespaces": schema.ListAttribute{
															Description:         "Namespaces is a set of namespace to which objects belong.",
															MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"node_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
															MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"nodes": schema.ListAttribute{
															Description:         "Nodes is a set of node name and objects must belong to these nodes.",
															MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"pod_phase_selectors": schema.ListAttribute{
															Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
															MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"pods": schema.MapAttribute{
															Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
															MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
															ElementType:         types.ListType{ElemType: types.StringType},
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},

												"value": schema.StringAttribute{
													Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
													MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"target_device": schema.StringAttribute{
											Description:         "TargetDevice represents the network device to be affected in target scope.",
											MarkdownDescription: "TargetDevice represents the network device to be affected in target scope.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
											MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"physicalmachine_chaos": schema.SingleNestedAttribute{
									Description:         "PhysicalMachineChaosSpec defines the desired state of PhysicalMachineChaos",
									MarkdownDescription: "PhysicalMachineChaosSpec defines the desired state of PhysicalMachineChaos",
									Attributes: map[string]schema.Attribute{
										"action": schema.StringAttribute{
											Description:         "the subAction, generate automatically",
											MarkdownDescription: "the subAction, generate automatically",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("stress-cpu", "stress-mem", "disk-read-payload", "disk-write-payload", "disk-fill", "network-corrupt", "network-duplicate", "network-loss", "network-delay", "network-partition", "network-dns", "network-bandwidth", "network-flood", "network-down", "process", "jvm-exception", "jvm-gc", "jvm-latency", "jvm-return", "jvm-stress", "jvm-rule-data", "jvm-mysql", "clock", "redis-expiration", "redis-penetration", "redis-cacheLimit", "redis-restart", "redis-stop", "kafka-fill", "kafka-flood", "kafka-io", "file-create", "file-modify", "file-delete", "file-rename", "file-append", "file-replace", "vm", "user_defined"),
											},
										},

										"address": schema.ListAttribute{
											Description:         "DEPRECATED: Use Selector instead. Only one of Address and Selector could be specified.",
											MarkdownDescription: "DEPRECATED: Use Selector instead. Only one of Address and Selector could be specified.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"clock": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"clock_ids_slice": schema.StringAttribute{
													Description:         "the identifier of the particular clock on which to act. More clock description in linux kernel can be found in man page of clock_getres, clock_gettime, clock_settime. Muti clock ids should be split with ','",
													MarkdownDescription: "the identifier of the particular clock on which to act. More clock description in linux kernel can be found in man page of clock_getres, clock_gettime, clock_settime. Muti clock ids should be split with ','",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"pid": schema.Int64Attribute{
													Description:         "the pid of target program.",
													MarkdownDescription: "the pid of target program.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"time_offset": schema.StringAttribute{
													Description:         "specifies the length of time offset.",
													MarkdownDescription: "specifies the length of time offset.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"disk_fill": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"fill_by_fallocate": schema.BoolAttribute{
													Description:         "fill disk by fallocate",
													MarkdownDescription: "fill disk by fallocate",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"path": schema.StringAttribute{
													Description:         "specifies the location to fill data in. if path not provided, payload will read/write from/into a temp file, temp file will be deleted after writing",
													MarkdownDescription: "specifies the location to fill data in. if path not provided, payload will read/write from/into a temp file, temp file will be deleted after writing",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"size": schema.StringAttribute{
													Description:         "specifies how many units of data will write into the file path. support unit: c=1, w=2, b=512, kB=1000, K=1024, MB=1000*1000, M=1024*1024, GB=1000*1000*1000, G=1024*1024*1024 BYTES. example : 1M | 512kB",
													MarkdownDescription: "specifies how many units of data will write into the file path. support unit: c=1, w=2, b=512, kB=1000, K=1024, MB=1000*1000, M=1024*1024, GB=1000*1000*1000, G=1024*1024*1024 BYTES. example : 1M | 512kB",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"disk_read_payload": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"path": schema.StringAttribute{
													Description:         "specifies the location to fill data in. if path not provided, payload will read/write from/into a temp file, temp file will be deleted after writing",
													MarkdownDescription: "specifies the location to fill data in. if path not provided, payload will read/write from/into a temp file, temp file will be deleted after writing",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"payload_process_num": schema.Int64Attribute{
													Description:         "specifies the number of process work on writing, default 1, only 1-255 is valid value",
													MarkdownDescription: "specifies the number of process work on writing, default 1, only 1-255 is valid value",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"size": schema.StringAttribute{
													Description:         "specifies how many units of data will write into the file path. support unit: c=1, w=2, b=512, kB=1000, K=1024, MB=1000*1000, M=1024*1024, GB=1000*1000*1000, G=1024*1024*1024 BYTES. example : 1M | 512kB",
													MarkdownDescription: "specifies how many units of data will write into the file path. support unit: c=1, w=2, b=512, kB=1000, K=1024, MB=1000*1000, M=1024*1024, GB=1000*1000*1000, G=1024*1024*1024 BYTES. example : 1M | 512kB",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"disk_write_payload": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"path": schema.StringAttribute{
													Description:         "specifies the location to fill data in. if path not provided, payload will read/write from/into a temp file, temp file will be deleted after writing",
													MarkdownDescription: "specifies the location to fill data in. if path not provided, payload will read/write from/into a temp file, temp file will be deleted after writing",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"payload_process_num": schema.Int64Attribute{
													Description:         "specifies the number of process work on writing, default 1, only 1-255 is valid value",
													MarkdownDescription: "specifies the number of process work on writing, default 1, only 1-255 is valid value",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"size": schema.StringAttribute{
													Description:         "specifies how many units of data will write into the file path. support unit: c=1, w=2, b=512, kB=1000, K=1024, MB=1000*1000, M=1024*1024, GB=1000*1000*1000, G=1024*1024*1024 BYTES. example : 1M | 512kB",
													MarkdownDescription: "specifies how many units of data will write into the file path. support unit: c=1, w=2, b=512, kB=1000, K=1024, MB=1000*1000, M=1024*1024, GB=1000*1000*1000, G=1024*1024*1024 BYTES. example : 1M | 512kB",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"duration": schema.StringAttribute{
											Description:         "Duration represents the duration of the chaos action",
											MarkdownDescription: "Duration represents the duration of the chaos action",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"file_append": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"count": schema.Int64Attribute{
													Description:         "Count is the number of times to append the data.",
													MarkdownDescription: "Count is the number of times to append the data.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"data": schema.StringAttribute{
													Description:         "Data is the data for append.",
													MarkdownDescription: "Data is the data for append.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"file_name": schema.StringAttribute{
													Description:         "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
													MarkdownDescription: "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"file_create": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"dir_name": schema.StringAttribute{
													Description:         "DirName is the directory name to create or delete.",
													MarkdownDescription: "DirName is the directory name to create or delete.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"file_name": schema.StringAttribute{
													Description:         "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
													MarkdownDescription: "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"file_delete": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"dir_name": schema.StringAttribute{
													Description:         "DirName is the directory name to create or delete.",
													MarkdownDescription: "DirName is the directory name to create or delete.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"file_name": schema.StringAttribute{
													Description:         "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
													MarkdownDescription: "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"file_modify": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"file_name": schema.StringAttribute{
													Description:         "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
													MarkdownDescription: "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"privilege": schema.Int64Attribute{
													Description:         "Privilege is the file privilege to be set.",
													MarkdownDescription: "Privilege is the file privilege to be set.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"file_rename": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"dest_file": schema.StringAttribute{
													Description:         "DestFile is the name to be renamed.",
													MarkdownDescription: "DestFile is the name to be renamed.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"source_file": schema.StringAttribute{
													Description:         "SourceFile is the name need to be renamed.",
													MarkdownDescription: "SourceFile is the name need to be renamed.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"file_replace": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"dest_string": schema.StringAttribute{
													Description:         "DestStr is the destination string of the file.",
													MarkdownDescription: "DestStr is the destination string of the file.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"file_name": schema.StringAttribute{
													Description:         "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
													MarkdownDescription: "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"line": schema.Int64Attribute{
													Description:         "Line is the line number of the file to be replaced.",
													MarkdownDescription: "Line is the line number of the file to be replaced.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"origin_string": schema.StringAttribute{
													Description:         "OriginStr is the origin string of the file.",
													MarkdownDescription: "OriginStr is the origin string of the file.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"http_abort": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"code": schema.StringAttribute{
													Description:         "Code is a rule to select target by http status code in response",
													MarkdownDescription: "Code is a rule to select target by http status code in response",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"method": schema.StringAttribute{
													Description:         "HTTP method",
													MarkdownDescription: "HTTP method",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"path": schema.StringAttribute{
													Description:         "Match path of Uri with wildcard matches",
													MarkdownDescription: "Match path of Uri with wildcard matches",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"port": schema.Int64Attribute{
													Description:         "The TCP port that the target service listens on",
													MarkdownDescription: "The TCP port that the target service listens on",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"proxy_ports": schema.ListAttribute{
													Description:         "Composed with one of the port of HTTP connection, we will only attack HTTP connection with port inside proxy_ports",
													MarkdownDescription: "Composed with one of the port of HTTP connection, we will only attack HTTP connection with port inside proxy_ports",
													ElementType:         types.StringType,
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"target": schema.StringAttribute{
													Description:         "HTTP target: Request or Response",
													MarkdownDescription: "HTTP target: Request or Response",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"http_config": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"file_path": schema.StringAttribute{
													Description:         "The config file path",
													MarkdownDescription: "The config file path",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"http_delay": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"code": schema.StringAttribute{
													Description:         "Code is a rule to select target by http status code in response",
													MarkdownDescription: "Code is a rule to select target by http status code in response",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"delay": schema.StringAttribute{
													Description:         "Delay represents the delay of the target request/response",
													MarkdownDescription: "Delay represents the delay of the target request/response",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"method": schema.StringAttribute{
													Description:         "HTTP method",
													MarkdownDescription: "HTTP method",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"path": schema.StringAttribute{
													Description:         "Match path of Uri with wildcard matches",
													MarkdownDescription: "Match path of Uri with wildcard matches",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"port": schema.Int64Attribute{
													Description:         "The TCP port that the target service listens on",
													MarkdownDescription: "The TCP port that the target service listens on",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"proxy_ports": schema.ListAttribute{
													Description:         "Composed with one of the port of HTTP connection, we will only attack HTTP connection with port inside proxy_ports",
													MarkdownDescription: "Composed with one of the port of HTTP connection, we will only attack HTTP connection with port inside proxy_ports",
													ElementType:         types.StringType,
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"target": schema.StringAttribute{
													Description:         "HTTP target: Request or Response",
													MarkdownDescription: "HTTP target: Request or Response",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"http_request": schema.SingleNestedAttribute{
											Description:         "used for HTTP request, now only support GET",
											MarkdownDescription: "used for HTTP request, now only support GET",
											Attributes: map[string]schema.Attribute{
												"count": schema.Int64Attribute{
													Description:         "The number of requests to send",
													MarkdownDescription: "The number of requests to send",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"enable_conn_pool": schema.BoolAttribute{
													Description:         "Enable connection pool",
													MarkdownDescription: "Enable connection pool",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"url": schema.StringAttribute{
													Description:         "Request to send'",
													MarkdownDescription: "Request to send'",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"jvm_exception": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"class": schema.StringAttribute{
													Description:         "Java class",
													MarkdownDescription: "Java class",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"exception": schema.StringAttribute{
													Description:         "the exception which needs to throw for action 'exception'",
													MarkdownDescription: "the exception which needs to throw for action 'exception'",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"method": schema.StringAttribute{
													Description:         "the method in Java class",
													MarkdownDescription: "the method in Java class",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"pid": schema.Int64Attribute{
													Description:         "the pid of Java process which needs to attach",
													MarkdownDescription: "the pid of Java process which needs to attach",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"port": schema.Int64Attribute{
													Description:         "the port of agent server, default 9277",
													MarkdownDescription: "the port of agent server, default 9277",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"jvm_gc": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"pid": schema.Int64Attribute{
													Description:         "the pid of Java process which needs to attach",
													MarkdownDescription: "the pid of Java process which needs to attach",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"port": schema.Int64Attribute{
													Description:         "the port of agent server, default 9277",
													MarkdownDescription: "the port of agent server, default 9277",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"jvm_latency": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"class": schema.StringAttribute{
													Description:         "Java class",
													MarkdownDescription: "Java class",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"latency": schema.Int64Attribute{
													Description:         "the latency duration for action 'latency', unit ms",
													MarkdownDescription: "the latency duration for action 'latency', unit ms",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"method": schema.StringAttribute{
													Description:         "the method in Java class",
													MarkdownDescription: "the method in Java class",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"pid": schema.Int64Attribute{
													Description:         "the pid of Java process which needs to attach",
													MarkdownDescription: "the pid of Java process which needs to attach",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"port": schema.Int64Attribute{
													Description:         "the port of agent server, default 9277",
													MarkdownDescription: "the port of agent server, default 9277",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"jvm_mysql": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"database": schema.StringAttribute{
													Description:         "the match database default value is '', means match all database",
													MarkdownDescription: "the match database default value is '', means match all database",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"exception": schema.StringAttribute{
													Description:         "The exception which needs to throw for action 'exception' or the exception message needs to throw in action 'mysql'",
													MarkdownDescription: "The exception which needs to throw for action 'exception' or the exception message needs to throw in action 'mysql'",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"latency": schema.Int64Attribute{
													Description:         "The latency duration for action 'latency' or the latency duration in action 'mysql'",
													MarkdownDescription: "The latency duration for action 'latency' or the latency duration in action 'mysql'",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"mysql_connector_version": schema.StringAttribute{
													Description:         "the version of mysql-connector-java, only support 5.X.X(set to '5') and 8.X.X(set to '8') now",
													MarkdownDescription: "the version of mysql-connector-java, only support 5.X.X(set to '5') and 8.X.X(set to '8') now",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"pid": schema.Int64Attribute{
													Description:         "the pid of Java process which needs to attach",
													MarkdownDescription: "the pid of Java process which needs to attach",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"port": schema.Int64Attribute{
													Description:         "the port of agent server, default 9277",
													MarkdownDescription: "the port of agent server, default 9277",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"sql_type": schema.StringAttribute{
													Description:         "the match sql type default value is '', means match all SQL type. The value can be 'select', 'insert', 'update', 'delete', 'replace'.",
													MarkdownDescription: "the match sql type default value is '', means match all SQL type. The value can be 'select', 'insert', 'update', 'delete', 'replace'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"table": schema.StringAttribute{
													Description:         "the match table default value is '', means match all table",
													MarkdownDescription: "the match table default value is '', means match all table",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"jvm_return": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"class": schema.StringAttribute{
													Description:         "Java class",
													MarkdownDescription: "Java class",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"method": schema.StringAttribute{
													Description:         "the method in Java class",
													MarkdownDescription: "the method in Java class",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"pid": schema.Int64Attribute{
													Description:         "the pid of Java process which needs to attach",
													MarkdownDescription: "the pid of Java process which needs to attach",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"port": schema.Int64Attribute{
													Description:         "the port of agent server, default 9277",
													MarkdownDescription: "the port of agent server, default 9277",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "the return value for action 'return'",
													MarkdownDescription: "the return value for action 'return'",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"jvm_rule_data": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"pid": schema.Int64Attribute{
													Description:         "the pid of Java process which needs to attach",
													MarkdownDescription: "the pid of Java process which needs to attach",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"port": schema.Int64Attribute{
													Description:         "the port of agent server, default 9277",
													MarkdownDescription: "the port of agent server, default 9277",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"rule_data": schema.StringAttribute{
													Description:         "RuleData used to save the rule file's data, will use it when recover",
													MarkdownDescription: "RuleData used to save the rule file's data, will use it when recover",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"jvm_stress": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"cpu_count": schema.Int64Attribute{
													Description:         "the CPU core number need to use, only set it when action is stress",
													MarkdownDescription: "the CPU core number need to use, only set it when action is stress",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"mem_type": schema.StringAttribute{
													Description:         "the memory type need to locate, only set it when action is stress, the value can be 'stack' or 'heap'",
													MarkdownDescription: "the memory type need to locate, only set it when action is stress, the value can be 'stack' or 'heap'",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"pid": schema.Int64Attribute{
													Description:         "the pid of Java process which needs to attach",
													MarkdownDescription: "the pid of Java process which needs to attach",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"port": schema.Int64Attribute{
													Description:         "the port of agent server, default 9277",
													MarkdownDescription: "the port of agent server, default 9277",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"kafka_fill": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"host": schema.StringAttribute{
													Description:         "The host of kafka server",
													MarkdownDescription: "The host of kafka server",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"max_bytes": schema.Int64Attribute{
													Description:         "The max bytes to fill",
													MarkdownDescription: "The max bytes to fill",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"message_size": schema.Int64Attribute{
													Description:         "The size of each message",
													MarkdownDescription: "The size of each message",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"password": schema.StringAttribute{
													Description:         "The password of kafka client",
													MarkdownDescription: "The password of kafka client",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"port": schema.Int64Attribute{
													Description:         "The port of kafka server",
													MarkdownDescription: "The port of kafka server",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"reload_command": schema.StringAttribute{
													Description:         "The command to reload kafka config",
													MarkdownDescription: "The command to reload kafka config",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"topic": schema.StringAttribute{
													Description:         "The topic to attack",
													MarkdownDescription: "The topic to attack",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"username": schema.StringAttribute{
													Description:         "The username of kafka client",
													MarkdownDescription: "The username of kafka client",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"kafka_flood": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"host": schema.StringAttribute{
													Description:         "The host of kafka server",
													MarkdownDescription: "The host of kafka server",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"message_size": schema.Int64Attribute{
													Description:         "The size of each message",
													MarkdownDescription: "The size of each message",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"password": schema.StringAttribute{
													Description:         "The password of kafka client",
													MarkdownDescription: "The password of kafka client",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"port": schema.Int64Attribute{
													Description:         "The port of kafka server",
													MarkdownDescription: "The port of kafka server",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"threads": schema.Int64Attribute{
													Description:         "The number of worker threads",
													MarkdownDescription: "The number of worker threads",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"topic": schema.StringAttribute{
													Description:         "The topic to attack",
													MarkdownDescription: "The topic to attack",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"username": schema.StringAttribute{
													Description:         "The username of kafka client",
													MarkdownDescription: "The username of kafka client",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"kafka_io": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_file": schema.StringAttribute{
													Description:         "The path of server config",
													MarkdownDescription: "The path of server config",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"non_readable": schema.BoolAttribute{
													Description:         "Make kafka cluster non-readable",
													MarkdownDescription: "Make kafka cluster non-readable",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"non_writable": schema.BoolAttribute{
													Description:         "Make kafka cluster non-writable",
													MarkdownDescription: "Make kafka cluster non-writable",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"topic": schema.StringAttribute{
													Description:         "The topic to attack",
													MarkdownDescription: "The topic to attack",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"mode": schema.StringAttribute{
											Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
											MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
											},
										},

										"network_bandwidth": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"buffer": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(1),
													},
												},

												"device": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"hostname": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"ip_address": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"limit": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(1),
													},
												},

												"minburst": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"peakrate": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"rate": schema.StringAttribute{
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

										"network_corrupt": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"correlation": schema.StringAttribute{
													Description:         "correlation is percentage (10 is 10%)",
													MarkdownDescription: "correlation is percentage (10 is 10%)",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"device": schema.StringAttribute{
													Description:         "the network interface to impact",
													MarkdownDescription: "the network interface to impact",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"egress_port": schema.StringAttribute{
													Description:         "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
													MarkdownDescription: "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"hostname": schema.StringAttribute{
													Description:         "only impact traffic to these hostnames",
													MarkdownDescription: "only impact traffic to these hostnames",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"ip_address": schema.StringAttribute{
													Description:         "only impact egress traffic to these IP addresses",
													MarkdownDescription: "only impact egress traffic to these IP addresses",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"ip_protocol": schema.StringAttribute{
													Description:         "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",
													MarkdownDescription: "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"percent": schema.StringAttribute{
													Description:         "percentage of packets to corrupt (10 is 10%)",
													MarkdownDescription: "percentage of packets to corrupt (10 is 10%)",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"source_port": schema.StringAttribute{
													Description:         "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
													MarkdownDescription: "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"network_delay": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"accept_tcp_flags": schema.StringAttribute{
													Description:         "only the packet which match the tcp flag can be accepted, others will be dropped. only set when the IPProtocol is tcp, used for partition.",
													MarkdownDescription: "only the packet which match the tcp flag can be accepted, others will be dropped. only set when the IPProtocol is tcp, used for partition.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"correlation": schema.StringAttribute{
													Description:         "correlation is percentage (10 is 10%)",
													MarkdownDescription: "correlation is percentage (10 is 10%)",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"device": schema.StringAttribute{
													Description:         "the network interface to impact",
													MarkdownDescription: "the network interface to impact",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"egress_port": schema.StringAttribute{
													Description:         "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
													MarkdownDescription: "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"hostname": schema.StringAttribute{
													Description:         "only impact traffic to these hostnames",
													MarkdownDescription: "only impact traffic to these hostnames",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"ip_address": schema.StringAttribute{
													Description:         "only impact egress traffic to these IP addresses",
													MarkdownDescription: "only impact egress traffic to these IP addresses",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"ip_protocol": schema.StringAttribute{
													Description:         "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",
													MarkdownDescription: "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"jitter": schema.StringAttribute{
													Description:         "jitter time, time units: ns, us (or µs), ms, s, m, h.",
													MarkdownDescription: "jitter time, time units: ns, us (or µs), ms, s, m, h.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"latency": schema.StringAttribute{
													Description:         "delay egress time, time units: ns, us (or µs), ms, s, m, h.",
													MarkdownDescription: "delay egress time, time units: ns, us (or µs), ms, s, m, h.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"source_port": schema.StringAttribute{
													Description:         "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
													MarkdownDescription: "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"network_dns": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"dns_domain_name": schema.StringAttribute{
													Description:         "map this host to specified IP",
													MarkdownDescription: "map this host to specified IP",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"dns_ip": schema.StringAttribute{
													Description:         "map specified host to this IP address",
													MarkdownDescription: "map specified host to this IP address",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"dns_server": schema.StringAttribute{
													Description:         "update the DNS server in /etc/resolv.conf with this value",
													MarkdownDescription: "update the DNS server in /etc/resolv.conf with this value",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"network_down": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"device": schema.StringAttribute{
													Description:         "The network interface to impact",
													MarkdownDescription: "The network interface to impact",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"duration": schema.StringAttribute{
													Description:         "NIC down time, time units: ns, us (or µs), ms, s, m, h.",
													MarkdownDescription: "NIC down time, time units: ns, us (or µs), ms, s, m, h.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"network_duplicate": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"correlation": schema.StringAttribute{
													Description:         "correlation is percentage (10 is 10%)",
													MarkdownDescription: "correlation is percentage (10 is 10%)",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"device": schema.StringAttribute{
													Description:         "the network interface to impact",
													MarkdownDescription: "the network interface to impact",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"egress_port": schema.StringAttribute{
													Description:         "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
													MarkdownDescription: "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"hostname": schema.StringAttribute{
													Description:         "only impact traffic to these hostnames",
													MarkdownDescription: "only impact traffic to these hostnames",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"ip_address": schema.StringAttribute{
													Description:         "only impact egress traffic to these IP addresses",
													MarkdownDescription: "only impact egress traffic to these IP addresses",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"ip_protocol": schema.StringAttribute{
													Description:         "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",
													MarkdownDescription: "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"percent": schema.StringAttribute{
													Description:         "percentage of packets to duplicate (10 is 10%)",
													MarkdownDescription: "percentage of packets to duplicate (10 is 10%)",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"source_port": schema.StringAttribute{
													Description:         "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
													MarkdownDescription: "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"network_flood": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"duration": schema.StringAttribute{
													Description:         "The number of seconds to run the iperf test",
													MarkdownDescription: "The number of seconds to run the iperf test",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"ip_address": schema.StringAttribute{
													Description:         "Generate traffic to this IP address",
													MarkdownDescription: "Generate traffic to this IP address",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"parallel": schema.Int64Attribute{
													Description:         "The number of iperf parallel client threads to run",
													MarkdownDescription: "The number of iperf parallel client threads to run",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"port": schema.StringAttribute{
													Description:         "Generate traffic to this port on the IP address",
													MarkdownDescription: "Generate traffic to this port on the IP address",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"rate": schema.StringAttribute{
													Description:         "The speed of network traffic, allows bps, kbps, mbps, gbps, tbps unit. bps means bytes per second",
													MarkdownDescription: "The speed of network traffic, allows bps, kbps, mbps, gbps, tbps unit. bps means bytes per second",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"network_loss": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"correlation": schema.StringAttribute{
													Description:         "correlation is percentage (10 is 10%)",
													MarkdownDescription: "correlation is percentage (10 is 10%)",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"device": schema.StringAttribute{
													Description:         "the network interface to impact",
													MarkdownDescription: "the network interface to impact",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"egress_port": schema.StringAttribute{
													Description:         "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
													MarkdownDescription: "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"hostname": schema.StringAttribute{
													Description:         "only impact traffic to these hostnames",
													MarkdownDescription: "only impact traffic to these hostnames",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"ip_address": schema.StringAttribute{
													Description:         "only impact egress traffic to these IP addresses",
													MarkdownDescription: "only impact egress traffic to these IP addresses",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"ip_protocol": schema.StringAttribute{
													Description:         "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",
													MarkdownDescription: "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"percent": schema.StringAttribute{
													Description:         "percentage of packets to loss (10 is 10%)",
													MarkdownDescription: "percentage of packets to loss (10 is 10%)",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"source_port": schema.StringAttribute{
													Description:         "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
													MarkdownDescription: "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"network_partition": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"accept_tcp_flags": schema.StringAttribute{
													Description:         "only the packet which match the tcp flag can be accepted, others will be dropped. only set when the IPProtocol is tcp, used for partition.",
													MarkdownDescription: "only the packet which match the tcp flag can be accepted, others will be dropped. only set when the IPProtocol is tcp, used for partition.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"device": schema.StringAttribute{
													Description:         "the network interface to impact",
													MarkdownDescription: "the network interface to impact",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"direction": schema.StringAttribute{
													Description:         "specifies the partition direction, values can be 'from', 'to'. 'from' means packets coming from the 'IPAddress' or 'Hostname' and going to your server, 'to' means packets originating from your server and going to the 'IPAddress' or 'Hostname'.",
													MarkdownDescription: "specifies the partition direction, values can be 'from', 'to'. 'from' means packets coming from the 'IPAddress' or 'Hostname' and going to your server, 'to' means packets originating from your server and going to the 'IPAddress' or 'Hostname'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"hostname": schema.StringAttribute{
													Description:         "only impact traffic to these hostnames",
													MarkdownDescription: "only impact traffic to these hostnames",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"ip_address": schema.StringAttribute{
													Description:         "only impact egress traffic to these IP addresses",
													MarkdownDescription: "only impact egress traffic to these IP addresses",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"ip_protocol": schema.StringAttribute{
													Description:         "only impact egress traffic to these IP addresses",
													MarkdownDescription: "only impact egress traffic to these IP addresses",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"process": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"process": schema.StringAttribute{
													Description:         "the process name or the process ID",
													MarkdownDescription: "the process name or the process ID",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"recover_cmd": schema.StringAttribute{
													Description:         "the command to be run when recovering experiment",
													MarkdownDescription: "the command to be run when recovering experiment",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"signal": schema.Int64Attribute{
													Description:         "the signal number to send",
													MarkdownDescription: "the signal number to send",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"redis_cache_limit": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"addr": schema.StringAttribute{
													Description:         "The adress of Redis server",
													MarkdownDescription: "The adress of Redis server",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"cache_size": schema.StringAttribute{
													Description:         "The size of 'maxmemory'",
													MarkdownDescription: "The size of 'maxmemory'",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"password": schema.StringAttribute{
													Description:         "The password of Redis server",
													MarkdownDescription: "The password of Redis server",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"percent": schema.StringAttribute{
													Description:         "Specifies maxmemory as a percentage of the original value",
													MarkdownDescription: "Specifies maxmemory as a percentage of the original value",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"redis_expiration": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"addr": schema.StringAttribute{
													Description:         "The adress of Redis server",
													MarkdownDescription: "The adress of Redis server",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"expiration": schema.StringAttribute{
													Description:         "The expiration of the keys",
													MarkdownDescription: "The expiration of the keys",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"key": schema.StringAttribute{
													Description:         "The keys to be expired",
													MarkdownDescription: "The keys to be expired",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"option": schema.StringAttribute{
													Description:         "Additional options for 'expiration'",
													MarkdownDescription: "Additional options for 'expiration'",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"password": schema.StringAttribute{
													Description:         "The password of Redis server",
													MarkdownDescription: "The password of Redis server",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"redis_penetration": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"addr": schema.StringAttribute{
													Description:         "The adress of Redis server",
													MarkdownDescription: "The adress of Redis server",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"password": schema.StringAttribute{
													Description:         "The password of Redis server",
													MarkdownDescription: "The password of Redis server",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"request_num": schema.Int64Attribute{
													Description:         "The number of requests to be sent",
													MarkdownDescription: "The number of requests to be sent",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"redis_restart": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"addr": schema.StringAttribute{
													Description:         "The adress of Redis server",
													MarkdownDescription: "The adress of Redis server",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"conf": schema.StringAttribute{
													Description:         "The path of Sentinel conf",
													MarkdownDescription: "The path of Sentinel conf",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"flush_config": schema.BoolAttribute{
													Description:         "The control flag determines whether to flush config",
													MarkdownDescription: "The control flag determines whether to flush config",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"password": schema.StringAttribute{
													Description:         "The password of Redis server",
													MarkdownDescription: "The password of Redis server",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"redis_path": schema.BoolAttribute{
													Description:         "The path of 'redis-server' command-line tool",
													MarkdownDescription: "The path of 'redis-server' command-line tool",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"redis_stop": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"addr": schema.StringAttribute{
													Description:         "The adress of Redis server",
													MarkdownDescription: "The adress of Redis server",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"conf": schema.StringAttribute{
													Description:         "The path of Sentinel conf",
													MarkdownDescription: "The path of Sentinel conf",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"flush_config": schema.BoolAttribute{
													Description:         "The control flag determines whether to flush config",
													MarkdownDescription: "The control flag determines whether to flush config",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"password": schema.StringAttribute{
													Description:         "The password of Redis server",
													MarkdownDescription: "The password of Redis server",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"redis_path": schema.BoolAttribute{
													Description:         "The path of 'redis-server' command-line tool",
													MarkdownDescription: "The path of 'redis-server' command-line tool",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"remote_cluster": schema.StringAttribute{
											Description:         "RemoteCluster represents the remote cluster where the chaos will be deployed",
											MarkdownDescription: "RemoteCluster represents the remote cluster where the chaos will be deployed",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"selector": schema.SingleNestedAttribute{
											Description:         "Selector is used to select physical machines that are used to inject chaos action.",
											MarkdownDescription: "Selector is used to select physical machines that are used to inject chaos action.",
											Attributes: map[string]schema.Attribute{
												"annotation_selectors": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
													MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"expression_selectors": schema.ListNestedAttribute{
													Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
													MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
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

												"field_selectors": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
													MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"label_selectors": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
													MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespaces": schema.ListAttribute{
													Description:         "Namespaces is a set of namespace to which objects belong.",
													MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"physical_machines": schema.MapAttribute{
													Description:         "PhysicalMachines is a map of string keys and a set values that used to select physical machines. The key defines the namespace which physical machine belong, and each value is a set of physical machine names.",
													MarkdownDescription: "PhysicalMachines is a map of string keys and a set values that used to select physical machines. The key defines the namespace which physical machine belong, and each value is a set of physical machine names.",
													ElementType:         types.ListType{ElemType: types.StringType},
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"stress_cpu": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"load": schema.Int64Attribute{
													Description:         "specifies P percent loading per CPU worker. 0 is effectively a sleep (no load) and 100 is full loading.",
													MarkdownDescription: "specifies P percent loading per CPU worker. 0 is effectively a sleep (no load) and 100 is full loading.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"options": schema.ListAttribute{
													Description:         "extend stress-ng options",
													MarkdownDescription: "extend stress-ng options",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"workers": schema.Int64Attribute{
													Description:         "specifies N workers to apply the stressor.",
													MarkdownDescription: "specifies N workers to apply the stressor.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"stress_mem": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"options": schema.ListAttribute{
													Description:         "extend stress-ng options",
													MarkdownDescription: "extend stress-ng options",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"size": schema.StringAttribute{
													Description:         "specifies N bytes consumed per vm worker, default is the total available memory. One can specify the size as % of total available memory or in units of B, KB/KiB, MB/MiB, GB/GiB, TB/TiB..",
													MarkdownDescription: "specifies N bytes consumed per vm worker, default is the total available memory. One can specify the size as % of total available memory or in units of B, KB/KiB, MB/MiB, GB/GiB, TB/TiB..",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"uid": schema.StringAttribute{
											Description:         "the experiment ID",
											MarkdownDescription: "the experiment ID",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"user_defined": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"attack_cmd": schema.StringAttribute{
													Description:         "The command to be executed when attack",
													MarkdownDescription: "The command to be executed when attack",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"recover_cmd": schema.StringAttribute{
													Description:         "The command to be executed when recover",
													MarkdownDescription: "The command to be executed when recover",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"value": schema.StringAttribute{
											Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of physical machines to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of physical machines the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
											MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of physical machines to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of physical machines the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"vm": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"vm_name": schema.StringAttribute{
													Description:         "The name of the VM to be injected",
													MarkdownDescription: "The name of the VM to be injected",
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

								"pod_chaos": schema.SingleNestedAttribute{
									Description:         "PodChaosSpec defines the attributes that a user creates on a chaos experiment about pods.",
									MarkdownDescription: "PodChaosSpec defines the attributes that a user creates on a chaos experiment about pods.",
									Attributes: map[string]schema.Attribute{
										"action": schema.StringAttribute{
											Description:         "Action defines the specific pod chaos action. Supported action: pod-kill / pod-failure / container-kill Default action: pod-kill",
											MarkdownDescription: "Action defines the specific pod chaos action. Supported action: pod-kill / pod-failure / container-kill Default action: pod-kill",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("pod-kill", "pod-failure", "container-kill"),
											},
										},

										"container_names": schema.ListAttribute{
											Description:         "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
											MarkdownDescription: "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"duration": schema.StringAttribute{
											Description:         "Duration represents the duration of the chaos action. It is required when the action is 'PodFailureAction'. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '-1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
											MarkdownDescription: "Duration represents the duration of the chaos action. It is required when the action is 'PodFailureAction'. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '-1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"grace_period": schema.Int64Attribute{
											Description:         "GracePeriod is used in pod-kill action. It represents the duration in seconds before the pod should be deleted. Value must be non-negative integer. The default value is zero that indicates delete immediately.",
											MarkdownDescription: "GracePeriod is used in pod-kill action. It represents the duration in seconds before the pod should be deleted. Value must be non-negative integer. The default value is zero that indicates delete immediately.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
											},
										},

										"mode": schema.StringAttribute{
											Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
											MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
											},
										},

										"remote_cluster": schema.StringAttribute{
											Description:         "RemoteCluster represents the remote cluster where the chaos will be deployed",
											MarkdownDescription: "RemoteCluster represents the remote cluster where the chaos will be deployed",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"selector": schema.SingleNestedAttribute{
											Description:         "Selector is used to select pods that are used to inject chaos action.",
											MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",
											Attributes: map[string]schema.Attribute{
												"annotation_selectors": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
													MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"expression_selectors": schema.ListNestedAttribute{
													Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
													MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
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

												"field_selectors": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
													MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"label_selectors": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
													MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespaces": schema.ListAttribute{
													Description:         "Namespaces is a set of namespace to which objects belong.",
													MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"node_selectors": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
													MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"nodes": schema.ListAttribute{
													Description:         "Nodes is a set of node name and objects must belong to these nodes.",
													MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"pod_phase_selectors": schema.ListAttribute{
													Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
													MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"pods": schema.MapAttribute{
													Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
													MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
													ElementType:         types.ListType{ElemType: types.StringType},
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"value": schema.StringAttribute{
											Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
											MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"schedule": schema.SingleNestedAttribute{
									Description:         "Schedule describe the Schedule(describing scheduled chaos) to be injected with chaos nodes. Only used when Type is TypeSchedule.",
									MarkdownDescription: "Schedule describe the Schedule(describing scheduled chaos) to be injected with chaos nodes. Only used when Type is TypeSchedule.",
									Attributes: map[string]schema.Attribute{
										"aws_chaos": schema.SingleNestedAttribute{
											Description:         "AWSChaosSpec is the content of the specification for an AWSChaos",
											MarkdownDescription: "AWSChaosSpec is the content of the specification for an AWSChaos",
											Attributes: map[string]schema.Attribute{
												"action": schema.StringAttribute{
													Description:         "Action defines the specific aws chaos action. Supported action: ec2-stop / ec2-restart / detach-volume Default action: ec2-stop",
													MarkdownDescription: "Action defines the specific aws chaos action. Supported action: ec2-stop / ec2-restart / detach-volume Default action: ec2-stop",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("ec2-stop", "ec2-restart", "detach-volume"),
													},
												},

												"aws_region": schema.StringAttribute{
													Description:         "AWSRegion defines the region of aws.",
													MarkdownDescription: "AWSRegion defines the region of aws.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"device_name": schema.StringAttribute{
													Description:         "DeviceName indicates the name of the device. Needed in detach-volume.",
													MarkdownDescription: "DeviceName indicates the name of the device. Needed in detach-volume.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"duration": schema.StringAttribute{
													Description:         "Duration represents the duration of the chaos action.",
													MarkdownDescription: "Duration represents the duration of the chaos action.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"ec2_instance": schema.StringAttribute{
													Description:         "Ec2Instance indicates the ID of the ec2 instance.",
													MarkdownDescription: "Ec2Instance indicates the ID of the ec2 instance.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"endpoint": schema.StringAttribute{
													Description:         "Endpoint indicates the endpoint of the aws server. Just used it in test now.",
													MarkdownDescription: "Endpoint indicates the endpoint of the aws server. Just used it in test now.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"remote_cluster": schema.StringAttribute{
													Description:         "RemoteCluster represents the remote cluster where the chaos will be deployed",
													MarkdownDescription: "RemoteCluster represents the remote cluster where the chaos will be deployed",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"secret_name": schema.StringAttribute{
													Description:         "SecretName defines the name of kubernetes secret.",
													MarkdownDescription: "SecretName defines the name of kubernetes secret.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"volume_id": schema.StringAttribute{
													Description:         "EbsVolume indicates the ID of the EBS volume. Needed in detach-volume.",
													MarkdownDescription: "EbsVolume indicates the ID of the EBS volume. Needed in detach-volume.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"azure_chaos": schema.SingleNestedAttribute{
											Description:         "AzureChaosSpec is the content of the specification for an AzureChaos",
											MarkdownDescription: "AzureChaosSpec is the content of the specification for an AzureChaos",
											Attributes: map[string]schema.Attribute{
												"action": schema.StringAttribute{
													Description:         "Action defines the specific azure chaos action. Supported action: vm-stop / vm-restart / disk-detach Default action: vm-stop",
													MarkdownDescription: "Action defines the specific azure chaos action. Supported action: vm-stop / vm-restart / disk-detach Default action: vm-stop",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("vm-stop", "vm-restart", "disk-detach"),
													},
												},

												"disk_name": schema.StringAttribute{
													Description:         "DiskName indicates the name of the disk. Needed in disk-detach.",
													MarkdownDescription: "DiskName indicates the name of the disk. Needed in disk-detach.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"duration": schema.StringAttribute{
													Description:         "Duration represents the duration of the chaos action.",
													MarkdownDescription: "Duration represents the duration of the chaos action.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"lun": schema.Int64Attribute{
													Description:         "LUN indicates the Logical Unit Number of the data disk. Needed in disk-detach.",
													MarkdownDescription: "LUN indicates the Logical Unit Number of the data disk. Needed in disk-detach.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"remote_cluster": schema.StringAttribute{
													Description:         "RemoteCluster represents the remote cluster where the chaos will be deployed",
													MarkdownDescription: "RemoteCluster represents the remote cluster where the chaos will be deployed",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"resource_group_name": schema.StringAttribute{
													Description:         "ResourceGroupName defines the name of ResourceGroup",
													MarkdownDescription: "ResourceGroupName defines the name of ResourceGroup",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"secret_name": schema.StringAttribute{
													Description:         "SecretName defines the name of kubernetes secret. It is used for Azure credentials.",
													MarkdownDescription: "SecretName defines the name of kubernetes secret. It is used for Azure credentials.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"subscription_id": schema.StringAttribute{
													Description:         "SubscriptionID defines the id of Azure subscription.",
													MarkdownDescription: "SubscriptionID defines the id of Azure subscription.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"vm_name": schema.StringAttribute{
													Description:         "VMName defines the name of Virtual Machine",
													MarkdownDescription: "VMName defines the name of Virtual Machine",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"block_chaos": schema.SingleNestedAttribute{
											Description:         "BlockChaosSpec is the content of the specification for a BlockChaos",
											MarkdownDescription: "BlockChaosSpec is the content of the specification for a BlockChaos",
											Attributes: map[string]schema.Attribute{
												"action": schema.StringAttribute{
													Description:         "Action defines the specific block chaos action. Supported action: delay",
													MarkdownDescription: "Action defines the specific block chaos action. Supported action: delay",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("delay"),
													},
												},

												"container_names": schema.ListAttribute{
													Description:         "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
													MarkdownDescription: "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"delay": schema.SingleNestedAttribute{
													Description:         "Delay defines the delay distribution.",
													MarkdownDescription: "Delay defines the delay distribution.",
													Attributes: map[string]schema.Attribute{
														"correlation": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"jitter": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"latency": schema.StringAttribute{
															Description:         "Latency defines the latency of every io request.",
															MarkdownDescription: "Latency defines the latency of every io request.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"duration": schema.StringAttribute{
													Description:         "Duration represents the duration of the chaos action.",
													MarkdownDescription: "Duration represents the duration of the chaos action.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"mode": schema.StringAttribute{
													Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
													MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
													},
												},

												"remote_cluster": schema.StringAttribute{
													Description:         "RemoteCluster represents the remote cluster where the chaos will be deployed",
													MarkdownDescription: "RemoteCluster represents the remote cluster where the chaos will be deployed",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"selector": schema.SingleNestedAttribute{
													Description:         "Selector is used to select pods that are used to inject chaos action.",
													MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",
													Attributes: map[string]schema.Attribute{
														"annotation_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
															MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"expression_selectors": schema.ListNestedAttribute{
															Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
															MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
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

														"field_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
															MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"label_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
															MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespaces": schema.ListAttribute{
															Description:         "Namespaces is a set of namespace to which objects belong.",
															MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"node_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
															MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"nodes": schema.ListAttribute{
															Description:         "Nodes is a set of node name and objects must belong to these nodes.",
															MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"pod_phase_selectors": schema.ListAttribute{
															Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
															MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"pods": schema.MapAttribute{
															Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
															MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
															ElementType:         types.ListType{ElemType: types.StringType},
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},

												"value": schema.StringAttribute{
													Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
													MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"volume_name": schema.StringAttribute{
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

										"concurrency_policy": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Forbid", "Allow"),
											},
										},

										"dns_chaos": schema.SingleNestedAttribute{
											Description:         "DNSChaosSpec defines the desired state of DNSChaos",
											MarkdownDescription: "DNSChaosSpec defines the desired state of DNSChaos",
											Attributes: map[string]schema.Attribute{
												"action": schema.StringAttribute{
													Description:         "Action defines the specific DNS chaos action. Supported action: error, random Default action: error",
													MarkdownDescription: "Action defines the specific DNS chaos action. Supported action: error, random Default action: error",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("error", "random"),
													},
												},

												"container_names": schema.ListAttribute{
													Description:         "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
													MarkdownDescription: "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"duration": schema.StringAttribute{
													Description:         "Duration represents the duration of the chaos action",
													MarkdownDescription: "Duration represents the duration of the chaos action",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"mode": schema.StringAttribute{
													Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
													MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
													},
												},

												"patterns": schema.ListAttribute{
													Description:         "Choose which domain names to take effect, support the placeholder ? and wildcard *, or the Specified domain name. Note:      1. The wildcard * must be at the end of the string. For example, chaos-*.org is invalid.      2. if the patterns is empty, will take effect on all the domain names. For example: 		The value is ['google.com', 'github.*', 'chaos-mes?.org'], 		will take effect on 'google.com', 'github.com' and 'chaos-mesh.org'",
													MarkdownDescription: "Choose which domain names to take effect, support the placeholder ? and wildcard *, or the Specified domain name. Note:      1. The wildcard * must be at the end of the string. For example, chaos-*.org is invalid.      2. if the patterns is empty, will take effect on all the domain names. For example: 		The value is ['google.com', 'github.*', 'chaos-mes?.org'], 		will take effect on 'google.com', 'github.com' and 'chaos-mesh.org'",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"remote_cluster": schema.StringAttribute{
													Description:         "RemoteCluster represents the remote cluster where the chaos will be deployed",
													MarkdownDescription: "RemoteCluster represents the remote cluster where the chaos will be deployed",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"selector": schema.SingleNestedAttribute{
													Description:         "Selector is used to select pods that are used to inject chaos action.",
													MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",
													Attributes: map[string]schema.Attribute{
														"annotation_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
															MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"expression_selectors": schema.ListNestedAttribute{
															Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
															MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
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

														"field_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
															MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"label_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
															MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespaces": schema.ListAttribute{
															Description:         "Namespaces is a set of namespace to which objects belong.",
															MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"node_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
															MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"nodes": schema.ListAttribute{
															Description:         "Nodes is a set of node name and objects must belong to these nodes.",
															MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"pod_phase_selectors": schema.ListAttribute{
															Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
															MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"pods": schema.MapAttribute{
															Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
															MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
															ElementType:         types.ListType{ElemType: types.StringType},
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},

												"value": schema.StringAttribute{
													Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
													MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"gcp_chaos": schema.SingleNestedAttribute{
											Description:         "GCPChaosSpec is the content of the specification for a GCPChaos",
											MarkdownDescription: "GCPChaosSpec is the content of the specification for a GCPChaos",
											Attributes: map[string]schema.Attribute{
												"action": schema.StringAttribute{
													Description:         "Action defines the specific gcp chaos action. Supported action: node-stop / node-reset / disk-loss Default action: node-stop",
													MarkdownDescription: "Action defines the specific gcp chaos action. Supported action: node-stop / node-reset / disk-loss Default action: node-stop",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("node-stop", "node-reset", "disk-loss"),
													},
												},

												"device_names": schema.ListAttribute{
													Description:         "The device name of disks to detach. Needed in disk-loss.",
													MarkdownDescription: "The device name of disks to detach. Needed in disk-loss.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"duration": schema.StringAttribute{
													Description:         "Duration represents the duration of the chaos action.",
													MarkdownDescription: "Duration represents the duration of the chaos action.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"instance": schema.StringAttribute{
													Description:         "Instance defines the name of the instance",
													MarkdownDescription: "Instance defines the name of the instance",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"project": schema.StringAttribute{
													Description:         "Project defines the ID of gcp project.",
													MarkdownDescription: "Project defines the ID of gcp project.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"remote_cluster": schema.StringAttribute{
													Description:         "RemoteCluster represents the remote cluster where the chaos will be deployed",
													MarkdownDescription: "RemoteCluster represents the remote cluster where the chaos will be deployed",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"secret_name": schema.StringAttribute{
													Description:         "SecretName defines the name of kubernetes secret. It is used for GCP credentials.",
													MarkdownDescription: "SecretName defines the name of kubernetes secret. It is used for GCP credentials.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"zone": schema.StringAttribute{
													Description:         "Zone defines the zone of gcp project.",
													MarkdownDescription: "Zone defines the zone of gcp project.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"history_limit": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(1),
											},
										},

										"http_chaos": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"abort": schema.BoolAttribute{
													Description:         "Abort is a rule to abort a http session.",
													MarkdownDescription: "Abort is a rule to abort a http session.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"code": schema.Int64Attribute{
													Description:         "Code is a rule to select target by http status code in response.",
													MarkdownDescription: "Code is a rule to select target by http status code in response.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"delay": schema.StringAttribute{
													Description:         "Delay represents the delay of the target request/response. A duration string is a possibly unsigned sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
													MarkdownDescription: "Delay represents the delay of the target request/response. A duration string is a possibly unsigned sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"duration": schema.StringAttribute{
													Description:         "Duration represents the duration of the chaos action.",
													MarkdownDescription: "Duration represents the duration of the chaos action.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"method": schema.StringAttribute{
													Description:         "Method is a rule to select target by http method in request.",
													MarkdownDescription: "Method is a rule to select target by http method in request.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"mode": schema.StringAttribute{
													Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
													MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
													},
												},

												"patch": schema.SingleNestedAttribute{
													Description:         "Patch is a rule to patch some contents in target.",
													MarkdownDescription: "Patch is a rule to patch some contents in target.",
													Attributes: map[string]schema.Attribute{
														"body": schema.SingleNestedAttribute{
															Description:         "Body is a rule to patch message body of target.",
															MarkdownDescription: "Body is a rule to patch message body of target.",
															Attributes: map[string]schema.Attribute{
																"type": schema.StringAttribute{
																	Description:         "Type represents the patch type, only support 'JSON' as [merge patch json](https://tools.ietf.org/html/rfc7396) currently.",
																	MarkdownDescription: "Type represents the patch type, only support 'JSON' as [merge patch json](https://tools.ietf.org/html/rfc7396) currently.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"value": schema.StringAttribute{
																	Description:         "Value is the patch contents.",
																	MarkdownDescription: "Value is the patch contents.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"headers": schema.ListAttribute{
															Description:         "Headers is a rule to append http headers of target. For example: '[['Set-Cookie', '<one cookie>'], ['Set-Cookie', '<another cookie>']]'.",
															MarkdownDescription: "Headers is a rule to append http headers of target. For example: '[['Set-Cookie', '<one cookie>'], ['Set-Cookie', '<another cookie>']]'.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"queries": schema.ListAttribute{
															Description:         "Queries is a rule to append uri queries of target(Request only). For example: '[['foo', 'bar'], ['foo', 'unknown']]'.",
															MarkdownDescription: "Queries is a rule to append uri queries of target(Request only). For example: '[['foo', 'bar'], ['foo', 'unknown']]'.",
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

												"path": schema.StringAttribute{
													Description:         "Path is a rule to select target by uri path in http request.",
													MarkdownDescription: "Path is a rule to select target by uri path in http request.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"port": schema.Int64Attribute{
													Description:         "Port represents the target port to be proxy of.",
													MarkdownDescription: "Port represents the target port to be proxy of.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"remote_cluster": schema.StringAttribute{
													Description:         "RemoteCluster represents the remote cluster where the chaos will be deployed",
													MarkdownDescription: "RemoteCluster represents the remote cluster where the chaos will be deployed",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"replace": schema.SingleNestedAttribute{
													Description:         "Replace is a rule to replace some contents in target.",
													MarkdownDescription: "Replace is a rule to replace some contents in target.",
													Attributes: map[string]schema.Attribute{
														"body": schema.StringAttribute{
															Description:         "Body is a rule to replace http message body in target.",
															MarkdownDescription: "Body is a rule to replace http message body in target.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																validators.Base64Validator(),
															},
														},

														"code": schema.Int64Attribute{
															Description:         "Code is a rule to replace http status code in response.",
															MarkdownDescription: "Code is a rule to replace http status code in response.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"headers": schema.MapAttribute{
															Description:         "Headers is a rule to replace http headers of target. The key-value pairs represent header name and header value pairs.",
															MarkdownDescription: "Headers is a rule to replace http headers of target. The key-value pairs represent header name and header value pairs.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"method": schema.StringAttribute{
															Description:         "Method is a rule to replace http method in request.",
															MarkdownDescription: "Method is a rule to replace http method in request.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"path": schema.StringAttribute{
															Description:         "Path is rule to to replace uri path in http request.",
															MarkdownDescription: "Path is rule to to replace uri path in http request.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"queries": schema.MapAttribute{
															Description:         "Queries is a rule to replace uri queries in http request. For example, with value '{ 'foo': 'unknown' }', the '/?foo=bar' will be altered to '/?foo=unknown',",
															MarkdownDescription: "Queries is a rule to replace uri queries in http request. For example, with value '{ 'foo': 'unknown' }', the '/?foo=bar' will be altered to '/?foo=unknown',",
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

												"request_headers": schema.MapAttribute{
													Description:         "RequestHeaders is a rule to select target by http headers in request. The key-value pairs represent header name and header value pairs.",
													MarkdownDescription: "RequestHeaders is a rule to select target by http headers in request. The key-value pairs represent header name and header value pairs.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"response_headers": schema.MapAttribute{
													Description:         "ResponseHeaders is a rule to select target by http headers in response. The key-value pairs represent header name and header value pairs.",
													MarkdownDescription: "ResponseHeaders is a rule to select target by http headers in response. The key-value pairs represent header name and header value pairs.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"selector": schema.SingleNestedAttribute{
													Description:         "Selector is used to select pods that are used to inject chaos action.",
													MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",
													Attributes: map[string]schema.Attribute{
														"annotation_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
															MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"expression_selectors": schema.ListNestedAttribute{
															Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
															MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
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

														"field_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
															MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"label_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
															MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespaces": schema.ListAttribute{
															Description:         "Namespaces is a set of namespace to which objects belong.",
															MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"node_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
															MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"nodes": schema.ListAttribute{
															Description:         "Nodes is a set of node name and objects must belong to these nodes.",
															MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"pod_phase_selectors": schema.ListAttribute{
															Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
															MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"pods": schema.MapAttribute{
															Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
															MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
															ElementType:         types.ListType{ElemType: types.StringType},
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},

												"target": schema.StringAttribute{
													Description:         "Target is the object to be selected and injected.",
													MarkdownDescription: "Target is the object to be selected and injected.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("Request", "Response"),
													},
												},

												"tls": schema.SingleNestedAttribute{
													Description:         "TLS is the tls config, will override PodHttpChaos if there are multiple HTTPChaos experiments are applied",
													MarkdownDescription: "TLS is the tls config, will override PodHttpChaos if there are multiple HTTPChaos experiments are applied",
													Attributes: map[string]schema.Attribute{
														"ca_name": schema.StringAttribute{
															Description:         "CAName represents the data name of ca file in secret, 'ca.crt' for example",
															MarkdownDescription: "CAName represents the data name of ca file in secret, 'ca.crt' for example",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"cert_name": schema.StringAttribute{
															Description:         "CertName represents the data name of cert file in secret, 'tls.crt' for example",
															MarkdownDescription: "CertName represents the data name of cert file in secret, 'tls.crt' for example",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"key_name": schema.StringAttribute{
															Description:         "KeyName represents the data name of key file in secret, 'tls.key' for example",
															MarkdownDescription: "KeyName represents the data name of key file in secret, 'tls.key' for example",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"secret_name": schema.StringAttribute{
															Description:         "SecretName represents the name of required secret resource",
															MarkdownDescription: "SecretName represents the name of required secret resource",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"secret_namespace": schema.StringAttribute{
															Description:         "SecretNamespace represents the namespace of required secret resource",
															MarkdownDescription: "SecretNamespace represents the namespace of required secret resource",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"value": schema.StringAttribute{
													Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
													MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"io_chaos": schema.SingleNestedAttribute{
											Description:         "IOChaosSpec defines the desired state of IOChaos",
											MarkdownDescription: "IOChaosSpec defines the desired state of IOChaos",
											Attributes: map[string]schema.Attribute{
												"action": schema.StringAttribute{
													Description:         "Action defines the specific pod chaos action. Supported action: latency / fault / attrOverride / mistake",
													MarkdownDescription: "Action defines the specific pod chaos action. Supported action: latency / fault / attrOverride / mistake",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("latency", "fault", "attrOverride", "mistake"),
													},
												},

												"attr": schema.SingleNestedAttribute{
													Description:         "Attr defines the overrided attribution",
													MarkdownDescription: "Attr defines the overrided attribution",
													Attributes: map[string]schema.Attribute{
														"atime": schema.SingleNestedAttribute{
															Description:         "Timespec represents a time",
															MarkdownDescription: "Timespec represents a time",
															Attributes: map[string]schema.Attribute{
																"nsec": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"sec": schema.Int64Attribute{
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

														"blocks": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"ctime": schema.SingleNestedAttribute{
															Description:         "Timespec represents a time",
															MarkdownDescription: "Timespec represents a time",
															Attributes: map[string]schema.Attribute{
																"nsec": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"sec": schema.Int64Attribute{
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

														"gid": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"ino": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"kind": schema.StringAttribute{
															Description:         "FileType represents type of file",
															MarkdownDescription: "FileType represents type of file",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"mtime": schema.SingleNestedAttribute{
															Description:         "Timespec represents a time",
															MarkdownDescription: "Timespec represents a time",
															Attributes: map[string]schema.Attribute{
																"nsec": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"sec": schema.Int64Attribute{
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

														"nlink": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"perm": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"rdev": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"size": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"uid": schema.Int64Attribute{
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

												"container_names": schema.ListAttribute{
													Description:         "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
													MarkdownDescription: "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"delay": schema.StringAttribute{
													Description:         "Delay defines the value of I/O chaos action delay. A delay string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
													MarkdownDescription: "Delay defines the value of I/O chaos action delay. A delay string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"duration": schema.StringAttribute{
													Description:         "Duration represents the duration of the chaos action. It is required when the action is 'PodFailureAction'. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '-1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
													MarkdownDescription: "Duration represents the duration of the chaos action. It is required when the action is 'PodFailureAction'. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '-1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"errno": schema.Int64Attribute{
													Description:         "Errno defines the error code that returned by I/O action. refer to: https://www-numi.fnal.gov/offline_software/srt_public_context/WebDocs/Errors/unix_system_errors.html",
													MarkdownDescription: "Errno defines the error code that returned by I/O action. refer to: https://www-numi.fnal.gov/offline_software/srt_public_context/WebDocs/Errors/unix_system_errors.html",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"methods": schema.ListAttribute{
													Description:         "Methods defines the I/O methods for injecting I/O chaos action. default: all I/O methods.",
													MarkdownDescription: "Methods defines the I/O methods for injecting I/O chaos action. default: all I/O methods.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"mistake": schema.SingleNestedAttribute{
													Description:         "Mistake defines what types of incorrectness are injected to IO operations",
													MarkdownDescription: "Mistake defines what types of incorrectness are injected to IO operations",
													Attributes: map[string]schema.Attribute{
														"filling": schema.StringAttribute{
															Description:         "Filling determines what is filled in the mistake data.",
															MarkdownDescription: "Filling determines what is filled in the mistake data.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("zero", "random"),
															},
														},

														"max_length": schema.Int64Attribute{
															Description:         "Max length of each wrong data segment in bytes",
															MarkdownDescription: "Max length of each wrong data segment in bytes",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(1),
															},
														},

														"max_occurrences": schema.Int64Attribute{
															Description:         "There will be [1, MaxOccurrences] segments of wrong data.",
															MarkdownDescription: "There will be [1, MaxOccurrences] segments of wrong data.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(1),
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"mode": schema.StringAttribute{
													Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
													MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
													},
												},

												"path": schema.StringAttribute{
													Description:         "Path defines the path of files for injecting I/O chaos action.",
													MarkdownDescription: "Path defines the path of files for injecting I/O chaos action.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"percent": schema.Int64Attribute{
													Description:         "Percent defines the percentage of injection errors and provides a number from 0-100. default: 100.",
													MarkdownDescription: "Percent defines the percentage of injection errors and provides a number from 0-100. default: 100.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"remote_cluster": schema.StringAttribute{
													Description:         "RemoteCluster represents the remote cluster where the chaos will be deployed",
													MarkdownDescription: "RemoteCluster represents the remote cluster where the chaos will be deployed",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"selector": schema.SingleNestedAttribute{
													Description:         "Selector is used to select pods that are used to inject chaos action.",
													MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",
													Attributes: map[string]schema.Attribute{
														"annotation_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
															MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"expression_selectors": schema.ListNestedAttribute{
															Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
															MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
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

														"field_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
															MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"label_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
															MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespaces": schema.ListAttribute{
															Description:         "Namespaces is a set of namespace to which objects belong.",
															MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"node_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
															MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"nodes": schema.ListAttribute{
															Description:         "Nodes is a set of node name and objects must belong to these nodes.",
															MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"pod_phase_selectors": schema.ListAttribute{
															Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
															MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"pods": schema.MapAttribute{
															Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
															MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
															ElementType:         types.ListType{ElemType: types.StringType},
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},

												"value": schema.StringAttribute{
													Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
													MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"volume_path": schema.StringAttribute{
													Description:         "VolumePath represents the mount path of injected volume",
													MarkdownDescription: "VolumePath represents the mount path of injected volume",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"jvm_chaos": schema.SingleNestedAttribute{
											Description:         "JVMChaosSpec defines the desired state of JVMChaos",
											MarkdownDescription: "JVMChaosSpec defines the desired state of JVMChaos",
											Attributes: map[string]schema.Attribute{
												"action": schema.StringAttribute{
													Description:         "Action defines the specific jvm chaos action. Supported action: latency;return;exception;stress;gc;ruleData",
													MarkdownDescription: "Action defines the specific jvm chaos action. Supported action: latency;return;exception;stress;gc;ruleData",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("latency", "return", "exception", "stress", "gc", "ruleData", "mysql"),
													},
												},

												"class": schema.StringAttribute{
													Description:         "Java class",
													MarkdownDescription: "Java class",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"container_names": schema.ListAttribute{
													Description:         "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
													MarkdownDescription: "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"cpu_count": schema.Int64Attribute{
													Description:         "the CPU core number needs to use, only set it when action is stress",
													MarkdownDescription: "the CPU core number needs to use, only set it when action is stress",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"database": schema.StringAttribute{
													Description:         "the match database default value is '', means match all database",
													MarkdownDescription: "the match database default value is '', means match all database",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"duration": schema.StringAttribute{
													Description:         "Duration represents the duration of the chaos action",
													MarkdownDescription: "Duration represents the duration of the chaos action",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"exception": schema.StringAttribute{
													Description:         "the exception which needs to throw for action 'exception' or the exception message needs to throw in action 'mysql'",
													MarkdownDescription: "the exception which needs to throw for action 'exception' or the exception message needs to throw in action 'mysql'",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"latency": schema.Int64Attribute{
													Description:         "the latency duration for action 'latency', unit ms or the latency duration in action 'mysql'",
													MarkdownDescription: "the latency duration for action 'latency', unit ms or the latency duration in action 'mysql'",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"mem_type": schema.StringAttribute{
													Description:         "the memory type needs to locate, only set it when action is stress, the value can be 'stack' or 'heap'",
													MarkdownDescription: "the memory type needs to locate, only set it when action is stress, the value can be 'stack' or 'heap'",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"method": schema.StringAttribute{
													Description:         "the method in Java class",
													MarkdownDescription: "the method in Java class",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"mode": schema.StringAttribute{
													Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
													MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
													},
												},

												"mysql_connector_version": schema.StringAttribute{
													Description:         "the version of mysql-connector-java, only support 5.X.X(set to '5') and 8.X.X(set to '8') now",
													MarkdownDescription: "the version of mysql-connector-java, only support 5.X.X(set to '5') and 8.X.X(set to '8') now",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "byteman rule name, should be unique, and will generate one if not set",
													MarkdownDescription: "byteman rule name, should be unique, and will generate one if not set",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"pid": schema.Int64Attribute{
													Description:         "the pid of Java process which needs to attach",
													MarkdownDescription: "the pid of Java process which needs to attach",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"port": schema.Int64Attribute{
													Description:         "the port of agent server, default 9277",
													MarkdownDescription: "the port of agent server, default 9277",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"remote_cluster": schema.StringAttribute{
													Description:         "RemoteCluster represents the remote cluster where the chaos will be deployed",
													MarkdownDescription: "RemoteCluster represents the remote cluster where the chaos will be deployed",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"rule_data": schema.StringAttribute{
													Description:         "the byteman rule's data for action 'ruleData'",
													MarkdownDescription: "the byteman rule's data for action 'ruleData'",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"selector": schema.SingleNestedAttribute{
													Description:         "Selector is used to select pods that are used to inject chaos action.",
													MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",
													Attributes: map[string]schema.Attribute{
														"annotation_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
															MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"expression_selectors": schema.ListNestedAttribute{
															Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
															MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
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

														"field_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
															MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"label_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
															MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespaces": schema.ListAttribute{
															Description:         "Namespaces is a set of namespace to which objects belong.",
															MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"node_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
															MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"nodes": schema.ListAttribute{
															Description:         "Nodes is a set of node name and objects must belong to these nodes.",
															MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"pod_phase_selectors": schema.ListAttribute{
															Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
															MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"pods": schema.MapAttribute{
															Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
															MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
															ElementType:         types.ListType{ElemType: types.StringType},
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},

												"sql_type": schema.StringAttribute{
													Description:         "the match sql type default value is '', means match all SQL type. The value can be 'select', 'insert', 'update', 'delete', 'replace'.",
													MarkdownDescription: "the match sql type default value is '', means match all SQL type. The value can be 'select', 'insert', 'update', 'delete', 'replace'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"table": schema.StringAttribute{
													Description:         "the match table default value is '', means match all table",
													MarkdownDescription: "the match table default value is '', means match all table",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
													MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"kernel_chaos": schema.SingleNestedAttribute{
											Description:         "KernelChaosSpec defines the desired state of KernelChaos",
											MarkdownDescription: "KernelChaosSpec defines the desired state of KernelChaos",
											Attributes: map[string]schema.Attribute{
												"container_names": schema.ListAttribute{
													Description:         "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
													MarkdownDescription: "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"duration": schema.StringAttribute{
													Description:         "Duration represents the duration of the chaos action",
													MarkdownDescription: "Duration represents the duration of the chaos action",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"fail_kern_request": schema.SingleNestedAttribute{
													Description:         "FailKernRequest defines the request of kernel injection",
													MarkdownDescription: "FailKernRequest defines the request of kernel injection",
													Attributes: map[string]schema.Attribute{
														"callchain": schema.ListNestedAttribute{
															Description:         "Callchain indicate a special call chain, such as:     ext4_mount       -> mount_subtree          -> ...             -> should_failslab With an optional set of predicates and an optional set of parameters, which used with predicates. You can read call chan and predicate examples from https://github.com/chaos-mesh/bpfki/tree/develop/examples to learn more. If no special call chain, just keep Callchain empty, which means it will fail at any call chain with slab alloc (eg: kmalloc).",
															MarkdownDescription: "Callchain indicate a special call chain, such as:     ext4_mount       -> mount_subtree          -> ...             -> should_failslab With an optional set of predicates and an optional set of parameters, which used with predicates. You can read call chan and predicate examples from https://github.com/chaos-mesh/bpfki/tree/develop/examples to learn more. If no special call chain, just keep Callchain empty, which means it will fail at any call chain with slab alloc (eg: kmalloc).",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"funcname": schema.StringAttribute{
																		Description:         "Funcname can be find from kernel source or '/proc/kallsyms', such as 'ext4_mount'",
																		MarkdownDescription: "Funcname can be find from kernel source or '/proc/kallsyms', such as 'ext4_mount'",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"parameters": schema.StringAttribute{
																		Description:         "Parameters is used with predicate, for example, if you want to inject slab error in 'd_alloc_parallel(struct dentry *parent, const struct qstr *name)' with a special name 'bananas', you need to set it to 'struct dentry *parent, const struct qstr *name' otherwise omit it.",
																		MarkdownDescription: "Parameters is used with predicate, for example, if you want to inject slab error in 'd_alloc_parallel(struct dentry *parent, const struct qstr *name)' with a special name 'bananas', you need to set it to 'struct dentry *parent, const struct qstr *name' otherwise omit it.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"predicate": schema.StringAttribute{
																		Description:         "Predicate will access the arguments of this Frame, example with Parameters's, you can set it to 'STRNCMP(name->name, 'bananas', 8)' to make inject only with it, or omit it to inject for all d_alloc_parallel call chain.",
																		MarkdownDescription: "Predicate will access the arguments of this Frame, example with Parameters's, you can set it to 'STRNCMP(name->name, 'bananas', 8)' to make inject only with it, or omit it to inject for all d_alloc_parallel call chain.",
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

														"failtype": schema.Int64Attribute{
															Description:         "FailType indicates what to fail, can be set to '0' / '1' / '2' If '0', indicates slab to fail (should_failslab) If '1', indicates alloc_page to fail (should_fail_alloc_page) If '2', indicates bio to fail (should_fail_bio) You can read:   1. https://www.kernel.org/doc/html/latest/fault-injection/fault-injection.html   2. http://github.com/iovisor/bcc/blob/master/tools/inject_example.txt to learn more",
															MarkdownDescription: "FailType indicates what to fail, can be set to '0' / '1' / '2' If '0', indicates slab to fail (should_failslab) If '1', indicates alloc_page to fail (should_fail_alloc_page) If '2', indicates bio to fail (should_fail_bio) You can read:   1. https://www.kernel.org/doc/html/latest/fault-injection/fault-injection.html   2. http://github.com/iovisor/bcc/blob/master/tools/inject_example.txt to learn more",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(0),
																int64validator.AtMost(2),
															},
														},

														"headers": schema.ListAttribute{
															Description:         "Headers indicates the appropriate kernel headers you need. Eg: 'linux/mmzone.h', 'linux/blkdev.h' and so on",
															MarkdownDescription: "Headers indicates the appropriate kernel headers you need. Eg: 'linux/mmzone.h', 'linux/blkdev.h' and so on",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"probability": schema.Int64Attribute{
															Description:         "Probability indicates the fails with probability. If you want 1%, please set this field with 1.",
															MarkdownDescription: "Probability indicates the fails with probability. If you want 1%, please set this field with 1.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(0),
																int64validator.AtMost(100),
															},
														},

														"times": schema.Int64Attribute{
															Description:         "Times indicates the max times of fails.",
															MarkdownDescription: "Times indicates the max times of fails.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(0),
															},
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},

												"mode": schema.StringAttribute{
													Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
													MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
													},
												},

												"remote_cluster": schema.StringAttribute{
													Description:         "RemoteCluster represents the remote cluster where the chaos will be deployed",
													MarkdownDescription: "RemoteCluster represents the remote cluster where the chaos will be deployed",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"selector": schema.SingleNestedAttribute{
													Description:         "Selector is used to select pods that are used to inject chaos action.",
													MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",
													Attributes: map[string]schema.Attribute{
														"annotation_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
															MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"expression_selectors": schema.ListNestedAttribute{
															Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
															MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
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

														"field_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
															MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"label_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
															MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespaces": schema.ListAttribute{
															Description:         "Namespaces is a set of namespace to which objects belong.",
															MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"node_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
															MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"nodes": schema.ListAttribute{
															Description:         "Nodes is a set of node name and objects must belong to these nodes.",
															MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"pod_phase_selectors": schema.ListAttribute{
															Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
															MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"pods": schema.MapAttribute{
															Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
															MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
															ElementType:         types.ListType{ElemType: types.StringType},
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},

												"value": schema.StringAttribute{
													Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
													MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"network_chaos": schema.SingleNestedAttribute{
											Description:         "NetworkChaosSpec defines the desired state of NetworkChaos",
											MarkdownDescription: "NetworkChaosSpec defines the desired state of NetworkChaos",
											Attributes: map[string]schema.Attribute{
												"action": schema.StringAttribute{
													Description:         "Action defines the specific network chaos action. Supported action: partition, netem, delay, loss, duplicate, corrupt Default action: delay",
													MarkdownDescription: "Action defines the specific network chaos action. Supported action: partition, netem, delay, loss, duplicate, corrupt Default action: delay",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("netem", "delay", "loss", "duplicate", "corrupt", "partition", "bandwidth"),
													},
												},

												"bandwidth": schema.SingleNestedAttribute{
													Description:         "Bandwidth represents the detail about bandwidth control action",
													MarkdownDescription: "Bandwidth represents the detail about bandwidth control action",
													Attributes: map[string]schema.Attribute{
														"buffer": schema.Int64Attribute{
															Description:         "Buffer is the maximum amount of bytes that tokens can be available for instantaneously.",
															MarkdownDescription: "Buffer is the maximum amount of bytes that tokens can be available for instantaneously.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(1),
															},
														},

														"limit": schema.Int64Attribute{
															Description:         "Limit is the number of bytes that can be queued waiting for tokens to become available.",
															MarkdownDescription: "Limit is the number of bytes that can be queued waiting for tokens to become available.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(1),
															},
														},

														"minburst": schema.Int64Attribute{
															Description:         "Minburst specifies the size of the peakrate bucket. For perfect accuracy, should be set to the MTU of the interface.  If a peakrate is needed, but some burstiness is acceptable, this size can be raised. A 3000 byte minburst allows around 3mbit/s of peakrate, given 1000 byte packets.",
															MarkdownDescription: "Minburst specifies the size of the peakrate bucket. For perfect accuracy, should be set to the MTU of the interface.  If a peakrate is needed, but some burstiness is acceptable, this size can be raised. A 3000 byte minburst allows around 3mbit/s of peakrate, given 1000 byte packets.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(0),
															},
														},

														"peakrate": schema.Int64Attribute{
															Description:         "Peakrate is the maximum depletion rate of the bucket. The peakrate does not need to be set, it is only necessary if perfect millisecond timescale shaping is required.",
															MarkdownDescription: "Peakrate is the maximum depletion rate of the bucket. The peakrate does not need to be set, it is only necessary if perfect millisecond timescale shaping is required.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(0),
															},
														},

														"rate": schema.StringAttribute{
															Description:         "Rate is the speed knob. Allows bit, kbit, mbit, gbit, tbit, bps, kbps, mbps, gbps, tbps unit. bps means bytes per second.",
															MarkdownDescription: "Rate is the speed knob. Allows bit, kbit, mbit, gbit, tbit, bps, kbps, mbps, gbps, tbps unit. bps means bytes per second.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"corrupt": schema.SingleNestedAttribute{
													Description:         "Corrupt represents the detail about corrupt action",
													MarkdownDescription: "Corrupt represents the detail about corrupt action",
													Attributes: map[string]schema.Attribute{
														"correlation": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"corrupt": schema.StringAttribute{
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

												"delay": schema.SingleNestedAttribute{
													Description:         "Delay represents the detail about delay action",
													MarkdownDescription: "Delay represents the detail about delay action",
													Attributes: map[string]schema.Attribute{
														"correlation": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"jitter": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"latency": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"reorder": schema.SingleNestedAttribute{
															Description:         "ReorderSpec defines details of packet reorder.",
															MarkdownDescription: "ReorderSpec defines details of packet reorder.",
															Attributes: map[string]schema.Attribute{
																"correlation": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"gap": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"reorder": schema.StringAttribute{
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
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"device": schema.StringAttribute{
													Description:         "Device represents the network device to be affected.",
													MarkdownDescription: "Device represents the network device to be affected.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"direction": schema.StringAttribute{
													Description:         "Direction represents the direction, this applies on netem and network partition action",
													MarkdownDescription: "Direction represents the direction, this applies on netem and network partition action",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("to", "from", "both"),
													},
												},

												"duplicate": schema.SingleNestedAttribute{
													Description:         "DuplicateSpec represents the detail about loss action",
													MarkdownDescription: "DuplicateSpec represents the detail about loss action",
													Attributes: map[string]schema.Attribute{
														"correlation": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"duplicate": schema.StringAttribute{
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

												"duration": schema.StringAttribute{
													Description:         "Duration represents the duration of the chaos action",
													MarkdownDescription: "Duration represents the duration of the chaos action",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"external_targets": schema.ListAttribute{
													Description:         "ExternalTargets represents network targets outside k8s",
													MarkdownDescription: "ExternalTargets represents network targets outside k8s",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"loss": schema.SingleNestedAttribute{
													Description:         "Loss represents the detail about loss action",
													MarkdownDescription: "Loss represents the detail about loss action",
													Attributes: map[string]schema.Attribute{
														"correlation": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"loss": schema.StringAttribute{
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

												"mode": schema.StringAttribute{
													Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
													MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
													},
												},

												"remote_cluster": schema.StringAttribute{
													Description:         "RemoteCluster represents the remote cluster where the chaos will be deployed",
													MarkdownDescription: "RemoteCluster represents the remote cluster where the chaos will be deployed",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"selector": schema.SingleNestedAttribute{
													Description:         "Selector is used to select pods that are used to inject chaos action.",
													MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",
													Attributes: map[string]schema.Attribute{
														"annotation_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
															MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"expression_selectors": schema.ListNestedAttribute{
															Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
															MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
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

														"field_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
															MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"label_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
															MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespaces": schema.ListAttribute{
															Description:         "Namespaces is a set of namespace to which objects belong.",
															MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"node_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
															MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"nodes": schema.ListAttribute{
															Description:         "Nodes is a set of node name and objects must belong to these nodes.",
															MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"pod_phase_selectors": schema.ListAttribute{
															Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
															MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"pods": schema.MapAttribute{
															Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
															MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
															ElementType:         types.ListType{ElemType: types.StringType},
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},

												"target": schema.SingleNestedAttribute{
													Description:         "Target represents network target, this applies on netem and network partition action",
													MarkdownDescription: "Target represents network target, this applies on netem and network partition action",
													Attributes: map[string]schema.Attribute{
														"mode": schema.StringAttribute{
															Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
															MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
															},
														},

														"selector": schema.SingleNestedAttribute{
															Description:         "Selector is used to select pods that are used to inject chaos action.",
															MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",
															Attributes: map[string]schema.Attribute{
																"annotation_selectors": schema.MapAttribute{
																	Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
																	MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"expression_selectors": schema.ListNestedAttribute{
																	Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
																	MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
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

																"field_selectors": schema.MapAttribute{
																	Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
																	MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"label_selectors": schema.MapAttribute{
																	Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
																	MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"namespaces": schema.ListAttribute{
																	Description:         "Namespaces is a set of namespace to which objects belong.",
																	MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"node_selectors": schema.MapAttribute{
																	Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
																	MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"nodes": schema.ListAttribute{
																	Description:         "Nodes is a set of node name and objects must belong to these nodes.",
																	MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"pod_phase_selectors": schema.ListAttribute{
																	Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
																	MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"pods": schema.MapAttribute{
																	Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
																	MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
																	ElementType:         types.ListType{ElemType: types.StringType},
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: true,
															Optional: false,
															Computed: false,
														},

														"value": schema.StringAttribute{
															Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
															MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"target_device": schema.StringAttribute{
													Description:         "TargetDevice represents the network device to be affected in target scope.",
													MarkdownDescription: "TargetDevice represents the network device to be affected in target scope.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
													MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"physicalmachine_chaos": schema.SingleNestedAttribute{
											Description:         "PhysicalMachineChaosSpec defines the desired state of PhysicalMachineChaos",
											MarkdownDescription: "PhysicalMachineChaosSpec defines the desired state of PhysicalMachineChaos",
											Attributes: map[string]schema.Attribute{
												"action": schema.StringAttribute{
													Description:         "the subAction, generate automatically",
													MarkdownDescription: "the subAction, generate automatically",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("stress-cpu", "stress-mem", "disk-read-payload", "disk-write-payload", "disk-fill", "network-corrupt", "network-duplicate", "network-loss", "network-delay", "network-partition", "network-dns", "network-bandwidth", "network-flood", "network-down", "process", "jvm-exception", "jvm-gc", "jvm-latency", "jvm-return", "jvm-stress", "jvm-rule-data", "jvm-mysql", "clock", "redis-expiration", "redis-penetration", "redis-cacheLimit", "redis-restart", "redis-stop", "kafka-fill", "kafka-flood", "kafka-io", "file-create", "file-modify", "file-delete", "file-rename", "file-append", "file-replace", "vm", "user_defined"),
													},
												},

												"address": schema.ListAttribute{
													Description:         "DEPRECATED: Use Selector instead. Only one of Address and Selector could be specified.",
													MarkdownDescription: "DEPRECATED: Use Selector instead. Only one of Address and Selector could be specified.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"clock": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"clock_ids_slice": schema.StringAttribute{
															Description:         "the identifier of the particular clock on which to act. More clock description in linux kernel can be found in man page of clock_getres, clock_gettime, clock_settime. Muti clock ids should be split with ','",
															MarkdownDescription: "the identifier of the particular clock on which to act. More clock description in linux kernel can be found in man page of clock_getres, clock_gettime, clock_settime. Muti clock ids should be split with ','",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"pid": schema.Int64Attribute{
															Description:         "the pid of target program.",
															MarkdownDescription: "the pid of target program.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"time_offset": schema.StringAttribute{
															Description:         "specifies the length of time offset.",
															MarkdownDescription: "specifies the length of time offset.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"disk_fill": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"fill_by_fallocate": schema.BoolAttribute{
															Description:         "fill disk by fallocate",
															MarkdownDescription: "fill disk by fallocate",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"path": schema.StringAttribute{
															Description:         "specifies the location to fill data in. if path not provided, payload will read/write from/into a temp file, temp file will be deleted after writing",
															MarkdownDescription: "specifies the location to fill data in. if path not provided, payload will read/write from/into a temp file, temp file will be deleted after writing",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"size": schema.StringAttribute{
															Description:         "specifies how many units of data will write into the file path. support unit: c=1, w=2, b=512, kB=1000, K=1024, MB=1000*1000, M=1024*1024, GB=1000*1000*1000, G=1024*1024*1024 BYTES. example : 1M | 512kB",
															MarkdownDescription: "specifies how many units of data will write into the file path. support unit: c=1, w=2, b=512, kB=1000, K=1024, MB=1000*1000, M=1024*1024, GB=1000*1000*1000, G=1024*1024*1024 BYTES. example : 1M | 512kB",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"disk_read_payload": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"path": schema.StringAttribute{
															Description:         "specifies the location to fill data in. if path not provided, payload will read/write from/into a temp file, temp file will be deleted after writing",
															MarkdownDescription: "specifies the location to fill data in. if path not provided, payload will read/write from/into a temp file, temp file will be deleted after writing",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"payload_process_num": schema.Int64Attribute{
															Description:         "specifies the number of process work on writing, default 1, only 1-255 is valid value",
															MarkdownDescription: "specifies the number of process work on writing, default 1, only 1-255 is valid value",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"size": schema.StringAttribute{
															Description:         "specifies how many units of data will write into the file path. support unit: c=1, w=2, b=512, kB=1000, K=1024, MB=1000*1000, M=1024*1024, GB=1000*1000*1000, G=1024*1024*1024 BYTES. example : 1M | 512kB",
															MarkdownDescription: "specifies how many units of data will write into the file path. support unit: c=1, w=2, b=512, kB=1000, K=1024, MB=1000*1000, M=1024*1024, GB=1000*1000*1000, G=1024*1024*1024 BYTES. example : 1M | 512kB",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"disk_write_payload": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"path": schema.StringAttribute{
															Description:         "specifies the location to fill data in. if path not provided, payload will read/write from/into a temp file, temp file will be deleted after writing",
															MarkdownDescription: "specifies the location to fill data in. if path not provided, payload will read/write from/into a temp file, temp file will be deleted after writing",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"payload_process_num": schema.Int64Attribute{
															Description:         "specifies the number of process work on writing, default 1, only 1-255 is valid value",
															MarkdownDescription: "specifies the number of process work on writing, default 1, only 1-255 is valid value",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"size": schema.StringAttribute{
															Description:         "specifies how many units of data will write into the file path. support unit: c=1, w=2, b=512, kB=1000, K=1024, MB=1000*1000, M=1024*1024, GB=1000*1000*1000, G=1024*1024*1024 BYTES. example : 1M | 512kB",
															MarkdownDescription: "specifies how many units of data will write into the file path. support unit: c=1, w=2, b=512, kB=1000, K=1024, MB=1000*1000, M=1024*1024, GB=1000*1000*1000, G=1024*1024*1024 BYTES. example : 1M | 512kB",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"duration": schema.StringAttribute{
													Description:         "Duration represents the duration of the chaos action",
													MarkdownDescription: "Duration represents the duration of the chaos action",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"file_append": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"count": schema.Int64Attribute{
															Description:         "Count is the number of times to append the data.",
															MarkdownDescription: "Count is the number of times to append the data.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"data": schema.StringAttribute{
															Description:         "Data is the data for append.",
															MarkdownDescription: "Data is the data for append.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"file_name": schema.StringAttribute{
															Description:         "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
															MarkdownDescription: "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"file_create": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"dir_name": schema.StringAttribute{
															Description:         "DirName is the directory name to create or delete.",
															MarkdownDescription: "DirName is the directory name to create or delete.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"file_name": schema.StringAttribute{
															Description:         "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
															MarkdownDescription: "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"file_delete": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"dir_name": schema.StringAttribute{
															Description:         "DirName is the directory name to create or delete.",
															MarkdownDescription: "DirName is the directory name to create or delete.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"file_name": schema.StringAttribute{
															Description:         "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
															MarkdownDescription: "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"file_modify": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"file_name": schema.StringAttribute{
															Description:         "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
															MarkdownDescription: "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"privilege": schema.Int64Attribute{
															Description:         "Privilege is the file privilege to be set.",
															MarkdownDescription: "Privilege is the file privilege to be set.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"file_rename": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"dest_file": schema.StringAttribute{
															Description:         "DestFile is the name to be renamed.",
															MarkdownDescription: "DestFile is the name to be renamed.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"source_file": schema.StringAttribute{
															Description:         "SourceFile is the name need to be renamed.",
															MarkdownDescription: "SourceFile is the name need to be renamed.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"file_replace": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"dest_string": schema.StringAttribute{
															Description:         "DestStr is the destination string of the file.",
															MarkdownDescription: "DestStr is the destination string of the file.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"file_name": schema.StringAttribute{
															Description:         "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
															MarkdownDescription: "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"line": schema.Int64Attribute{
															Description:         "Line is the line number of the file to be replaced.",
															MarkdownDescription: "Line is the line number of the file to be replaced.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"origin_string": schema.StringAttribute{
															Description:         "OriginStr is the origin string of the file.",
															MarkdownDescription: "OriginStr is the origin string of the file.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"http_abort": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"code": schema.StringAttribute{
															Description:         "Code is a rule to select target by http status code in response",
															MarkdownDescription: "Code is a rule to select target by http status code in response",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"method": schema.StringAttribute{
															Description:         "HTTP method",
															MarkdownDescription: "HTTP method",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"path": schema.StringAttribute{
															Description:         "Match path of Uri with wildcard matches",
															MarkdownDescription: "Match path of Uri with wildcard matches",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"port": schema.Int64Attribute{
															Description:         "The TCP port that the target service listens on",
															MarkdownDescription: "The TCP port that the target service listens on",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"proxy_ports": schema.ListAttribute{
															Description:         "Composed with one of the port of HTTP connection, we will only attack HTTP connection with port inside proxy_ports",
															MarkdownDescription: "Composed with one of the port of HTTP connection, we will only attack HTTP connection with port inside proxy_ports",
															ElementType:         types.StringType,
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"target": schema.StringAttribute{
															Description:         "HTTP target: Request or Response",
															MarkdownDescription: "HTTP target: Request or Response",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"http_config": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"file_path": schema.StringAttribute{
															Description:         "The config file path",
															MarkdownDescription: "The config file path",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"http_delay": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"code": schema.StringAttribute{
															Description:         "Code is a rule to select target by http status code in response",
															MarkdownDescription: "Code is a rule to select target by http status code in response",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"delay": schema.StringAttribute{
															Description:         "Delay represents the delay of the target request/response",
															MarkdownDescription: "Delay represents the delay of the target request/response",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"method": schema.StringAttribute{
															Description:         "HTTP method",
															MarkdownDescription: "HTTP method",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"path": schema.StringAttribute{
															Description:         "Match path of Uri with wildcard matches",
															MarkdownDescription: "Match path of Uri with wildcard matches",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"port": schema.Int64Attribute{
															Description:         "The TCP port that the target service listens on",
															MarkdownDescription: "The TCP port that the target service listens on",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"proxy_ports": schema.ListAttribute{
															Description:         "Composed with one of the port of HTTP connection, we will only attack HTTP connection with port inside proxy_ports",
															MarkdownDescription: "Composed with one of the port of HTTP connection, we will only attack HTTP connection with port inside proxy_ports",
															ElementType:         types.StringType,
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"target": schema.StringAttribute{
															Description:         "HTTP target: Request or Response",
															MarkdownDescription: "HTTP target: Request or Response",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"http_request": schema.SingleNestedAttribute{
													Description:         "used for HTTP request, now only support GET",
													MarkdownDescription: "used for HTTP request, now only support GET",
													Attributes: map[string]schema.Attribute{
														"count": schema.Int64Attribute{
															Description:         "The number of requests to send",
															MarkdownDescription: "The number of requests to send",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"enable_conn_pool": schema.BoolAttribute{
															Description:         "Enable connection pool",
															MarkdownDescription: "Enable connection pool",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"url": schema.StringAttribute{
															Description:         "Request to send'",
															MarkdownDescription: "Request to send'",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"jvm_exception": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"class": schema.StringAttribute{
															Description:         "Java class",
															MarkdownDescription: "Java class",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"exception": schema.StringAttribute{
															Description:         "the exception which needs to throw for action 'exception'",
															MarkdownDescription: "the exception which needs to throw for action 'exception'",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"method": schema.StringAttribute{
															Description:         "the method in Java class",
															MarkdownDescription: "the method in Java class",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"pid": schema.Int64Attribute{
															Description:         "the pid of Java process which needs to attach",
															MarkdownDescription: "the pid of Java process which needs to attach",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"port": schema.Int64Attribute{
															Description:         "the port of agent server, default 9277",
															MarkdownDescription: "the port of agent server, default 9277",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"jvm_gc": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"pid": schema.Int64Attribute{
															Description:         "the pid of Java process which needs to attach",
															MarkdownDescription: "the pid of Java process which needs to attach",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"port": schema.Int64Attribute{
															Description:         "the port of agent server, default 9277",
															MarkdownDescription: "the port of agent server, default 9277",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"jvm_latency": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"class": schema.StringAttribute{
															Description:         "Java class",
															MarkdownDescription: "Java class",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"latency": schema.Int64Attribute{
															Description:         "the latency duration for action 'latency', unit ms",
															MarkdownDescription: "the latency duration for action 'latency', unit ms",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"method": schema.StringAttribute{
															Description:         "the method in Java class",
															MarkdownDescription: "the method in Java class",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"pid": schema.Int64Attribute{
															Description:         "the pid of Java process which needs to attach",
															MarkdownDescription: "the pid of Java process which needs to attach",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"port": schema.Int64Attribute{
															Description:         "the port of agent server, default 9277",
															MarkdownDescription: "the port of agent server, default 9277",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"jvm_mysql": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"database": schema.StringAttribute{
															Description:         "the match database default value is '', means match all database",
															MarkdownDescription: "the match database default value is '', means match all database",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"exception": schema.StringAttribute{
															Description:         "The exception which needs to throw for action 'exception' or the exception message needs to throw in action 'mysql'",
															MarkdownDescription: "The exception which needs to throw for action 'exception' or the exception message needs to throw in action 'mysql'",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"latency": schema.Int64Attribute{
															Description:         "The latency duration for action 'latency' or the latency duration in action 'mysql'",
															MarkdownDescription: "The latency duration for action 'latency' or the latency duration in action 'mysql'",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"mysql_connector_version": schema.StringAttribute{
															Description:         "the version of mysql-connector-java, only support 5.X.X(set to '5') and 8.X.X(set to '8') now",
															MarkdownDescription: "the version of mysql-connector-java, only support 5.X.X(set to '5') and 8.X.X(set to '8') now",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"pid": schema.Int64Attribute{
															Description:         "the pid of Java process which needs to attach",
															MarkdownDescription: "the pid of Java process which needs to attach",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"port": schema.Int64Attribute{
															Description:         "the port of agent server, default 9277",
															MarkdownDescription: "the port of agent server, default 9277",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"sql_type": schema.StringAttribute{
															Description:         "the match sql type default value is '', means match all SQL type. The value can be 'select', 'insert', 'update', 'delete', 'replace'.",
															MarkdownDescription: "the match sql type default value is '', means match all SQL type. The value can be 'select', 'insert', 'update', 'delete', 'replace'.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"table": schema.StringAttribute{
															Description:         "the match table default value is '', means match all table",
															MarkdownDescription: "the match table default value is '', means match all table",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"jvm_return": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"class": schema.StringAttribute{
															Description:         "Java class",
															MarkdownDescription: "Java class",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"method": schema.StringAttribute{
															Description:         "the method in Java class",
															MarkdownDescription: "the method in Java class",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"pid": schema.Int64Attribute{
															Description:         "the pid of Java process which needs to attach",
															MarkdownDescription: "the pid of Java process which needs to attach",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"port": schema.Int64Attribute{
															Description:         "the port of agent server, default 9277",
															MarkdownDescription: "the port of agent server, default 9277",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"value": schema.StringAttribute{
															Description:         "the return value for action 'return'",
															MarkdownDescription: "the return value for action 'return'",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"jvm_rule_data": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"pid": schema.Int64Attribute{
															Description:         "the pid of Java process which needs to attach",
															MarkdownDescription: "the pid of Java process which needs to attach",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"port": schema.Int64Attribute{
															Description:         "the port of agent server, default 9277",
															MarkdownDescription: "the port of agent server, default 9277",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"rule_data": schema.StringAttribute{
															Description:         "RuleData used to save the rule file's data, will use it when recover",
															MarkdownDescription: "RuleData used to save the rule file's data, will use it when recover",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"jvm_stress": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"cpu_count": schema.Int64Attribute{
															Description:         "the CPU core number need to use, only set it when action is stress",
															MarkdownDescription: "the CPU core number need to use, only set it when action is stress",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"mem_type": schema.StringAttribute{
															Description:         "the memory type need to locate, only set it when action is stress, the value can be 'stack' or 'heap'",
															MarkdownDescription: "the memory type need to locate, only set it when action is stress, the value can be 'stack' or 'heap'",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"pid": schema.Int64Attribute{
															Description:         "the pid of Java process which needs to attach",
															MarkdownDescription: "the pid of Java process which needs to attach",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"port": schema.Int64Attribute{
															Description:         "the port of agent server, default 9277",
															MarkdownDescription: "the port of agent server, default 9277",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"kafka_fill": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"host": schema.StringAttribute{
															Description:         "The host of kafka server",
															MarkdownDescription: "The host of kafka server",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"max_bytes": schema.Int64Attribute{
															Description:         "The max bytes to fill",
															MarkdownDescription: "The max bytes to fill",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"message_size": schema.Int64Attribute{
															Description:         "The size of each message",
															MarkdownDescription: "The size of each message",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"password": schema.StringAttribute{
															Description:         "The password of kafka client",
															MarkdownDescription: "The password of kafka client",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"port": schema.Int64Attribute{
															Description:         "The port of kafka server",
															MarkdownDescription: "The port of kafka server",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"reload_command": schema.StringAttribute{
															Description:         "The command to reload kafka config",
															MarkdownDescription: "The command to reload kafka config",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"topic": schema.StringAttribute{
															Description:         "The topic to attack",
															MarkdownDescription: "The topic to attack",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"username": schema.StringAttribute{
															Description:         "The username of kafka client",
															MarkdownDescription: "The username of kafka client",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"kafka_flood": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"host": schema.StringAttribute{
															Description:         "The host of kafka server",
															MarkdownDescription: "The host of kafka server",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"message_size": schema.Int64Attribute{
															Description:         "The size of each message",
															MarkdownDescription: "The size of each message",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"password": schema.StringAttribute{
															Description:         "The password of kafka client",
															MarkdownDescription: "The password of kafka client",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"port": schema.Int64Attribute{
															Description:         "The port of kafka server",
															MarkdownDescription: "The port of kafka server",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"threads": schema.Int64Attribute{
															Description:         "The number of worker threads",
															MarkdownDescription: "The number of worker threads",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"topic": schema.StringAttribute{
															Description:         "The topic to attack",
															MarkdownDescription: "The topic to attack",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"username": schema.StringAttribute{
															Description:         "The username of kafka client",
															MarkdownDescription: "The username of kafka client",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"kafka_io": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"config_file": schema.StringAttribute{
															Description:         "The path of server config",
															MarkdownDescription: "The path of server config",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"non_readable": schema.BoolAttribute{
															Description:         "Make kafka cluster non-readable",
															MarkdownDescription: "Make kafka cluster non-readable",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"non_writable": schema.BoolAttribute{
															Description:         "Make kafka cluster non-writable",
															MarkdownDescription: "Make kafka cluster non-writable",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"topic": schema.StringAttribute{
															Description:         "The topic to attack",
															MarkdownDescription: "The topic to attack",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"mode": schema.StringAttribute{
													Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
													MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
													},
												},

												"network_bandwidth": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"buffer": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(1),
															},
														},

														"device": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"hostname": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"ip_address": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"limit": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(1),
															},
														},

														"minburst": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"peakrate": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"rate": schema.StringAttribute{
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

												"network_corrupt": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"correlation": schema.StringAttribute{
															Description:         "correlation is percentage (10 is 10%)",
															MarkdownDescription: "correlation is percentage (10 is 10%)",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"device": schema.StringAttribute{
															Description:         "the network interface to impact",
															MarkdownDescription: "the network interface to impact",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"egress_port": schema.StringAttribute{
															Description:         "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
															MarkdownDescription: "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"hostname": schema.StringAttribute{
															Description:         "only impact traffic to these hostnames",
															MarkdownDescription: "only impact traffic to these hostnames",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"ip_address": schema.StringAttribute{
															Description:         "only impact egress traffic to these IP addresses",
															MarkdownDescription: "only impact egress traffic to these IP addresses",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"ip_protocol": schema.StringAttribute{
															Description:         "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",
															MarkdownDescription: "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"percent": schema.StringAttribute{
															Description:         "percentage of packets to corrupt (10 is 10%)",
															MarkdownDescription: "percentage of packets to corrupt (10 is 10%)",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"source_port": schema.StringAttribute{
															Description:         "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
															MarkdownDescription: "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"network_delay": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"accept_tcp_flags": schema.StringAttribute{
															Description:         "only the packet which match the tcp flag can be accepted, others will be dropped. only set when the IPProtocol is tcp, used for partition.",
															MarkdownDescription: "only the packet which match the tcp flag can be accepted, others will be dropped. only set when the IPProtocol is tcp, used for partition.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"correlation": schema.StringAttribute{
															Description:         "correlation is percentage (10 is 10%)",
															MarkdownDescription: "correlation is percentage (10 is 10%)",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"device": schema.StringAttribute{
															Description:         "the network interface to impact",
															MarkdownDescription: "the network interface to impact",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"egress_port": schema.StringAttribute{
															Description:         "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
															MarkdownDescription: "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"hostname": schema.StringAttribute{
															Description:         "only impact traffic to these hostnames",
															MarkdownDescription: "only impact traffic to these hostnames",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"ip_address": schema.StringAttribute{
															Description:         "only impact egress traffic to these IP addresses",
															MarkdownDescription: "only impact egress traffic to these IP addresses",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"ip_protocol": schema.StringAttribute{
															Description:         "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",
															MarkdownDescription: "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"jitter": schema.StringAttribute{
															Description:         "jitter time, time units: ns, us (or µs), ms, s, m, h.",
															MarkdownDescription: "jitter time, time units: ns, us (or µs), ms, s, m, h.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"latency": schema.StringAttribute{
															Description:         "delay egress time, time units: ns, us (or µs), ms, s, m, h.",
															MarkdownDescription: "delay egress time, time units: ns, us (or µs), ms, s, m, h.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"source_port": schema.StringAttribute{
															Description:         "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
															MarkdownDescription: "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"network_dns": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"dns_domain_name": schema.StringAttribute{
															Description:         "map this host to specified IP",
															MarkdownDescription: "map this host to specified IP",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"dns_ip": schema.StringAttribute{
															Description:         "map specified host to this IP address",
															MarkdownDescription: "map specified host to this IP address",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"dns_server": schema.StringAttribute{
															Description:         "update the DNS server in /etc/resolv.conf with this value",
															MarkdownDescription: "update the DNS server in /etc/resolv.conf with this value",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"network_down": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"device": schema.StringAttribute{
															Description:         "The network interface to impact",
															MarkdownDescription: "The network interface to impact",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"duration": schema.StringAttribute{
															Description:         "NIC down time, time units: ns, us (or µs), ms, s, m, h.",
															MarkdownDescription: "NIC down time, time units: ns, us (or µs), ms, s, m, h.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"network_duplicate": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"correlation": schema.StringAttribute{
															Description:         "correlation is percentage (10 is 10%)",
															MarkdownDescription: "correlation is percentage (10 is 10%)",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"device": schema.StringAttribute{
															Description:         "the network interface to impact",
															MarkdownDescription: "the network interface to impact",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"egress_port": schema.StringAttribute{
															Description:         "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
															MarkdownDescription: "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"hostname": schema.StringAttribute{
															Description:         "only impact traffic to these hostnames",
															MarkdownDescription: "only impact traffic to these hostnames",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"ip_address": schema.StringAttribute{
															Description:         "only impact egress traffic to these IP addresses",
															MarkdownDescription: "only impact egress traffic to these IP addresses",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"ip_protocol": schema.StringAttribute{
															Description:         "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",
															MarkdownDescription: "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"percent": schema.StringAttribute{
															Description:         "percentage of packets to duplicate (10 is 10%)",
															MarkdownDescription: "percentage of packets to duplicate (10 is 10%)",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"source_port": schema.StringAttribute{
															Description:         "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
															MarkdownDescription: "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"network_flood": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"duration": schema.StringAttribute{
															Description:         "The number of seconds to run the iperf test",
															MarkdownDescription: "The number of seconds to run the iperf test",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"ip_address": schema.StringAttribute{
															Description:         "Generate traffic to this IP address",
															MarkdownDescription: "Generate traffic to this IP address",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"parallel": schema.Int64Attribute{
															Description:         "The number of iperf parallel client threads to run",
															MarkdownDescription: "The number of iperf parallel client threads to run",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"port": schema.StringAttribute{
															Description:         "Generate traffic to this port on the IP address",
															MarkdownDescription: "Generate traffic to this port on the IP address",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"rate": schema.StringAttribute{
															Description:         "The speed of network traffic, allows bps, kbps, mbps, gbps, tbps unit. bps means bytes per second",
															MarkdownDescription: "The speed of network traffic, allows bps, kbps, mbps, gbps, tbps unit. bps means bytes per second",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"network_loss": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"correlation": schema.StringAttribute{
															Description:         "correlation is percentage (10 is 10%)",
															MarkdownDescription: "correlation is percentage (10 is 10%)",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"device": schema.StringAttribute{
															Description:         "the network interface to impact",
															MarkdownDescription: "the network interface to impact",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"egress_port": schema.StringAttribute{
															Description:         "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
															MarkdownDescription: "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"hostname": schema.StringAttribute{
															Description:         "only impact traffic to these hostnames",
															MarkdownDescription: "only impact traffic to these hostnames",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"ip_address": schema.StringAttribute{
															Description:         "only impact egress traffic to these IP addresses",
															MarkdownDescription: "only impact egress traffic to these IP addresses",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"ip_protocol": schema.StringAttribute{
															Description:         "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",
															MarkdownDescription: "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"percent": schema.StringAttribute{
															Description:         "percentage of packets to loss (10 is 10%)",
															MarkdownDescription: "percentage of packets to loss (10 is 10%)",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"source_port": schema.StringAttribute{
															Description:         "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
															MarkdownDescription: "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"network_partition": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"accept_tcp_flags": schema.StringAttribute{
															Description:         "only the packet which match the tcp flag can be accepted, others will be dropped. only set when the IPProtocol is tcp, used for partition.",
															MarkdownDescription: "only the packet which match the tcp flag can be accepted, others will be dropped. only set when the IPProtocol is tcp, used for partition.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"device": schema.StringAttribute{
															Description:         "the network interface to impact",
															MarkdownDescription: "the network interface to impact",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"direction": schema.StringAttribute{
															Description:         "specifies the partition direction, values can be 'from', 'to'. 'from' means packets coming from the 'IPAddress' or 'Hostname' and going to your server, 'to' means packets originating from your server and going to the 'IPAddress' or 'Hostname'.",
															MarkdownDescription: "specifies the partition direction, values can be 'from', 'to'. 'from' means packets coming from the 'IPAddress' or 'Hostname' and going to your server, 'to' means packets originating from your server and going to the 'IPAddress' or 'Hostname'.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"hostname": schema.StringAttribute{
															Description:         "only impact traffic to these hostnames",
															MarkdownDescription: "only impact traffic to these hostnames",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"ip_address": schema.StringAttribute{
															Description:         "only impact egress traffic to these IP addresses",
															MarkdownDescription: "only impact egress traffic to these IP addresses",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"ip_protocol": schema.StringAttribute{
															Description:         "only impact egress traffic to these IP addresses",
															MarkdownDescription: "only impact egress traffic to these IP addresses",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"process": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"process": schema.StringAttribute{
															Description:         "the process name or the process ID",
															MarkdownDescription: "the process name or the process ID",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"recover_cmd": schema.StringAttribute{
															Description:         "the command to be run when recovering experiment",
															MarkdownDescription: "the command to be run when recovering experiment",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"signal": schema.Int64Attribute{
															Description:         "the signal number to send",
															MarkdownDescription: "the signal number to send",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"redis_cache_limit": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"addr": schema.StringAttribute{
															Description:         "The adress of Redis server",
															MarkdownDescription: "The adress of Redis server",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"cache_size": schema.StringAttribute{
															Description:         "The size of 'maxmemory'",
															MarkdownDescription: "The size of 'maxmemory'",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"password": schema.StringAttribute{
															Description:         "The password of Redis server",
															MarkdownDescription: "The password of Redis server",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"percent": schema.StringAttribute{
															Description:         "Specifies maxmemory as a percentage of the original value",
															MarkdownDescription: "Specifies maxmemory as a percentage of the original value",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"redis_expiration": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"addr": schema.StringAttribute{
															Description:         "The adress of Redis server",
															MarkdownDescription: "The adress of Redis server",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"expiration": schema.StringAttribute{
															Description:         "The expiration of the keys",
															MarkdownDescription: "The expiration of the keys",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"key": schema.StringAttribute{
															Description:         "The keys to be expired",
															MarkdownDescription: "The keys to be expired",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"option": schema.StringAttribute{
															Description:         "Additional options for 'expiration'",
															MarkdownDescription: "Additional options for 'expiration'",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"password": schema.StringAttribute{
															Description:         "The password of Redis server",
															MarkdownDescription: "The password of Redis server",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"redis_penetration": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"addr": schema.StringAttribute{
															Description:         "The adress of Redis server",
															MarkdownDescription: "The adress of Redis server",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"password": schema.StringAttribute{
															Description:         "The password of Redis server",
															MarkdownDescription: "The password of Redis server",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"request_num": schema.Int64Attribute{
															Description:         "The number of requests to be sent",
															MarkdownDescription: "The number of requests to be sent",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"redis_restart": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"addr": schema.StringAttribute{
															Description:         "The adress of Redis server",
															MarkdownDescription: "The adress of Redis server",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"conf": schema.StringAttribute{
															Description:         "The path of Sentinel conf",
															MarkdownDescription: "The path of Sentinel conf",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"flush_config": schema.BoolAttribute{
															Description:         "The control flag determines whether to flush config",
															MarkdownDescription: "The control flag determines whether to flush config",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"password": schema.StringAttribute{
															Description:         "The password of Redis server",
															MarkdownDescription: "The password of Redis server",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"redis_path": schema.BoolAttribute{
															Description:         "The path of 'redis-server' command-line tool",
															MarkdownDescription: "The path of 'redis-server' command-line tool",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"redis_stop": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"addr": schema.StringAttribute{
															Description:         "The adress of Redis server",
															MarkdownDescription: "The adress of Redis server",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"conf": schema.StringAttribute{
															Description:         "The path of Sentinel conf",
															MarkdownDescription: "The path of Sentinel conf",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"flush_config": schema.BoolAttribute{
															Description:         "The control flag determines whether to flush config",
															MarkdownDescription: "The control flag determines whether to flush config",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"password": schema.StringAttribute{
															Description:         "The password of Redis server",
															MarkdownDescription: "The password of Redis server",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"redis_path": schema.BoolAttribute{
															Description:         "The path of 'redis-server' command-line tool",
															MarkdownDescription: "The path of 'redis-server' command-line tool",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"remote_cluster": schema.StringAttribute{
													Description:         "RemoteCluster represents the remote cluster where the chaos will be deployed",
													MarkdownDescription: "RemoteCluster represents the remote cluster where the chaos will be deployed",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"selector": schema.SingleNestedAttribute{
													Description:         "Selector is used to select physical machines that are used to inject chaos action.",
													MarkdownDescription: "Selector is used to select physical machines that are used to inject chaos action.",
													Attributes: map[string]schema.Attribute{
														"annotation_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
															MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"expression_selectors": schema.ListNestedAttribute{
															Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
															MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
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

														"field_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
															MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"label_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
															MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespaces": schema.ListAttribute{
															Description:         "Namespaces is a set of namespace to which objects belong.",
															MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"physical_machines": schema.MapAttribute{
															Description:         "PhysicalMachines is a map of string keys and a set values that used to select physical machines. The key defines the namespace which physical machine belong, and each value is a set of physical machine names.",
															MarkdownDescription: "PhysicalMachines is a map of string keys and a set values that used to select physical machines. The key defines the namespace which physical machine belong, and each value is a set of physical machine names.",
															ElementType:         types.ListType{ElemType: types.StringType},
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"stress_cpu": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"load": schema.Int64Attribute{
															Description:         "specifies P percent loading per CPU worker. 0 is effectively a sleep (no load) and 100 is full loading.",
															MarkdownDescription: "specifies P percent loading per CPU worker. 0 is effectively a sleep (no load) and 100 is full loading.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"options": schema.ListAttribute{
															Description:         "extend stress-ng options",
															MarkdownDescription: "extend stress-ng options",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"workers": schema.Int64Attribute{
															Description:         "specifies N workers to apply the stressor.",
															MarkdownDescription: "specifies N workers to apply the stressor.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"stress_mem": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"options": schema.ListAttribute{
															Description:         "extend stress-ng options",
															MarkdownDescription: "extend stress-ng options",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"size": schema.StringAttribute{
															Description:         "specifies N bytes consumed per vm worker, default is the total available memory. One can specify the size as % of total available memory or in units of B, KB/KiB, MB/MiB, GB/GiB, TB/TiB..",
															MarkdownDescription: "specifies N bytes consumed per vm worker, default is the total available memory. One can specify the size as % of total available memory or in units of B, KB/KiB, MB/MiB, GB/GiB, TB/TiB..",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"uid": schema.StringAttribute{
													Description:         "the experiment ID",
													MarkdownDescription: "the experiment ID",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"user_defined": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"attack_cmd": schema.StringAttribute{
															Description:         "The command to be executed when attack",
															MarkdownDescription: "The command to be executed when attack",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"recover_cmd": schema.StringAttribute{
															Description:         "The command to be executed when recover",
															MarkdownDescription: "The command to be executed when recover",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"value": schema.StringAttribute{
													Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of physical machines to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of physical machines the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
													MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of physical machines to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of physical machines the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"vm": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"vm_name": schema.StringAttribute{
															Description:         "The name of the VM to be injected",
															MarkdownDescription: "The name of the VM to be injected",
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

										"pod_chaos": schema.SingleNestedAttribute{
											Description:         "PodChaosSpec defines the attributes that a user creates on a chaos experiment about pods.",
											MarkdownDescription: "PodChaosSpec defines the attributes that a user creates on a chaos experiment about pods.",
											Attributes: map[string]schema.Attribute{
												"action": schema.StringAttribute{
													Description:         "Action defines the specific pod chaos action. Supported action: pod-kill / pod-failure / container-kill Default action: pod-kill",
													MarkdownDescription: "Action defines the specific pod chaos action. Supported action: pod-kill / pod-failure / container-kill Default action: pod-kill",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("pod-kill", "pod-failure", "container-kill"),
													},
												},

												"container_names": schema.ListAttribute{
													Description:         "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
													MarkdownDescription: "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"duration": schema.StringAttribute{
													Description:         "Duration represents the duration of the chaos action. It is required when the action is 'PodFailureAction'. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '-1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
													MarkdownDescription: "Duration represents the duration of the chaos action. It is required when the action is 'PodFailureAction'. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '-1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"grace_period": schema.Int64Attribute{
													Description:         "GracePeriod is used in pod-kill action. It represents the duration in seconds before the pod should be deleted. Value must be non-negative integer. The default value is zero that indicates delete immediately.",
													MarkdownDescription: "GracePeriod is used in pod-kill action. It represents the duration in seconds before the pod should be deleted. Value must be non-negative integer. The default value is zero that indicates delete immediately.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(0),
													},
												},

												"mode": schema.StringAttribute{
													Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
													MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
													},
												},

												"remote_cluster": schema.StringAttribute{
													Description:         "RemoteCluster represents the remote cluster where the chaos will be deployed",
													MarkdownDescription: "RemoteCluster represents the remote cluster where the chaos will be deployed",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"selector": schema.SingleNestedAttribute{
													Description:         "Selector is used to select pods that are used to inject chaos action.",
													MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",
													Attributes: map[string]schema.Attribute{
														"annotation_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
															MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"expression_selectors": schema.ListNestedAttribute{
															Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
															MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
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

														"field_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
															MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"label_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
															MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespaces": schema.ListAttribute{
															Description:         "Namespaces is a set of namespace to which objects belong.",
															MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"node_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
															MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"nodes": schema.ListAttribute{
															Description:         "Nodes is a set of node name and objects must belong to these nodes.",
															MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"pod_phase_selectors": schema.ListAttribute{
															Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
															MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"pods": schema.MapAttribute{
															Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
															MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
															ElementType:         types.ListType{ElemType: types.StringType},
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},

												"value": schema.StringAttribute{
													Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
													MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"schedule": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"starting_deadline_seconds": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
											},
										},

										"stress_chaos": schema.SingleNestedAttribute{
											Description:         "StressChaosSpec defines the desired state of StressChaos",
											MarkdownDescription: "StressChaosSpec defines the desired state of StressChaos",
											Attributes: map[string]schema.Attribute{
												"container_names": schema.ListAttribute{
													Description:         "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
													MarkdownDescription: "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"duration": schema.StringAttribute{
													Description:         "Duration represents the duration of the chaos action",
													MarkdownDescription: "Duration represents the duration of the chaos action",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"mode": schema.StringAttribute{
													Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
													MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
													},
												},

												"remote_cluster": schema.StringAttribute{
													Description:         "RemoteCluster represents the remote cluster where the chaos will be deployed",
													MarkdownDescription: "RemoteCluster represents the remote cluster where the chaos will be deployed",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"selector": schema.SingleNestedAttribute{
													Description:         "Selector is used to select pods that are used to inject chaos action.",
													MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",
													Attributes: map[string]schema.Attribute{
														"annotation_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
															MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"expression_selectors": schema.ListNestedAttribute{
															Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
															MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
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

														"field_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
															MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"label_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
															MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespaces": schema.ListAttribute{
															Description:         "Namespaces is a set of namespace to which objects belong.",
															MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"node_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
															MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"nodes": schema.ListAttribute{
															Description:         "Nodes is a set of node name and objects must belong to these nodes.",
															MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"pod_phase_selectors": schema.ListAttribute{
															Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
															MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"pods": schema.MapAttribute{
															Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
															MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
															ElementType:         types.ListType{ElemType: types.StringType},
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},

												"stressng_stressors": schema.StringAttribute{
													Description:         "StressngStressors defines plenty of stressors just like 'Stressors' except that it's an experimental feature and more powerful. You can define stressors in 'stress-ng' (see also 'man stress-ng') dialect, however not all of the supported stressors are well tested. It maybe retired in later releases. You should always use 'Stressors' to define the stressors and use this only when you want more stressors unsupported by 'Stressors'. When both 'StressngStressors' and 'Stressors' are defined, 'StressngStressors' wins.",
													MarkdownDescription: "StressngStressors defines plenty of stressors just like 'Stressors' except that it's an experimental feature and more powerful. You can define stressors in 'stress-ng' (see also 'man stress-ng') dialect, however not all of the supported stressors are well tested. It maybe retired in later releases. You should always use 'Stressors' to define the stressors and use this only when you want more stressors unsupported by 'Stressors'. When both 'StressngStressors' and 'Stressors' are defined, 'StressngStressors' wins.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"stressors": schema.SingleNestedAttribute{
													Description:         "Stressors defines plenty of stressors supported to stress system components out. You can use one or more of them to make up various kinds of stresses. At least one of the stressors should be specified.",
													MarkdownDescription: "Stressors defines plenty of stressors supported to stress system components out. You can use one or more of them to make up various kinds of stresses. At least one of the stressors should be specified.",
													Attributes: map[string]schema.Attribute{
														"cpu": schema.SingleNestedAttribute{
															Description:         "CPUStressor stresses CPU out",
															MarkdownDescription: "CPUStressor stresses CPU out",
															Attributes: map[string]schema.Attribute{
																"load": schema.Int64Attribute{
																	Description:         "Load specifies P percent loading per CPU worker. 0 is effectively a sleep (no load) and 100 is full loading.",
																	MarkdownDescription: "Load specifies P percent loading per CPU worker. 0 is effectively a sleep (no load) and 100 is full loading.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(0),
																		int64validator.AtMost(100),
																	},
																},

																"options": schema.ListAttribute{
																	Description:         "extend stress-ng options",
																	MarkdownDescription: "extend stress-ng options",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"workers": schema.Int64Attribute{
																	Description:         "Workers specifies N workers to apply the stressor. Maximum 8192 workers can run by stress-ng",
																	MarkdownDescription: "Workers specifies N workers to apply the stressor. Maximum 8192 workers can run by stress-ng",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtMost(8192),
																	},
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"memory": schema.SingleNestedAttribute{
															Description:         "MemoryStressor stresses virtual memory out",
															MarkdownDescription: "MemoryStressor stresses virtual memory out",
															Attributes: map[string]schema.Attribute{
																"oom_score_adj": schema.Int64Attribute{
																	Description:         "OOMScoreAdj sets the oom_score_adj of the stress process. See 'man 5 proc' to know more about this option.",
																	MarkdownDescription: "OOMScoreAdj sets the oom_score_adj of the stress process. See 'man 5 proc' to know more about this option.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(-1000),
																		int64validator.AtMost(1000),
																	},
																},

																"options": schema.ListAttribute{
																	Description:         "extend stress-ng options",
																	MarkdownDescription: "extend stress-ng options",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"size": schema.StringAttribute{
																	Description:         "Size specifies N bytes consumed per vm worker, default is the total available memory. One can specify the size as % of total available memory or in units of B, KB/KiB, MB/MiB, GB/GiB, TB/TiB.",
																	MarkdownDescription: "Size specifies N bytes consumed per vm worker, default is the total available memory. One can specify the size as % of total available memory or in units of B, KB/KiB, MB/MiB, GB/GiB, TB/TiB.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"workers": schema.Int64Attribute{
																	Description:         "Workers specifies N workers to apply the stressor. Maximum 8192 workers can run by stress-ng",
																	MarkdownDescription: "Workers specifies N workers to apply the stressor. Maximum 8192 workers can run by stress-ng",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtMost(8192),
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

												"value": schema.StringAttribute{
													Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
													MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"time_chaos": schema.SingleNestedAttribute{
											Description:         "TimeChaosSpec defines the desired state of TimeChaos",
											MarkdownDescription: "TimeChaosSpec defines the desired state of TimeChaos",
											Attributes: map[string]schema.Attribute{
												"clock_ids": schema.ListAttribute{
													Description:         "ClockIds defines all affected clock id All available options are ['CLOCK_REALTIME','CLOCK_MONOTONIC','CLOCK_PROCESS_CPUTIME_ID','CLOCK_THREAD_CPUTIME_ID', 'CLOCK_MONOTONIC_RAW','CLOCK_REALTIME_COARSE','CLOCK_MONOTONIC_COARSE','CLOCK_BOOTTIME','CLOCK_REALTIME_ALARM', 'CLOCK_BOOTTIME_ALARM'] Default value is ['CLOCK_REALTIME']",
													MarkdownDescription: "ClockIds defines all affected clock id All available options are ['CLOCK_REALTIME','CLOCK_MONOTONIC','CLOCK_PROCESS_CPUTIME_ID','CLOCK_THREAD_CPUTIME_ID', 'CLOCK_MONOTONIC_RAW','CLOCK_REALTIME_COARSE','CLOCK_MONOTONIC_COARSE','CLOCK_BOOTTIME','CLOCK_REALTIME_ALARM', 'CLOCK_BOOTTIME_ALARM'] Default value is ['CLOCK_REALTIME']",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"container_names": schema.ListAttribute{
													Description:         "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
													MarkdownDescription: "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"duration": schema.StringAttribute{
													Description:         "Duration represents the duration of the chaos action",
													MarkdownDescription: "Duration represents the duration of the chaos action",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"mode": schema.StringAttribute{
													Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
													MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
													},
												},

												"remote_cluster": schema.StringAttribute{
													Description:         "RemoteCluster represents the remote cluster where the chaos will be deployed",
													MarkdownDescription: "RemoteCluster represents the remote cluster where the chaos will be deployed",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"selector": schema.SingleNestedAttribute{
													Description:         "Selector is used to select pods that are used to inject chaos action.",
													MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",
													Attributes: map[string]schema.Attribute{
														"annotation_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
															MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"expression_selectors": schema.ListNestedAttribute{
															Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
															MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
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

														"field_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
															MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"label_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
															MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespaces": schema.ListAttribute{
															Description:         "Namespaces is a set of namespace to which objects belong.",
															MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"node_selectors": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
															MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"nodes": schema.ListAttribute{
															Description:         "Nodes is a set of node name and objects must belong to these nodes.",
															MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"pod_phase_selectors": schema.ListAttribute{
															Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
															MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"pods": schema.MapAttribute{
															Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
															MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
															ElementType:         types.ListType{ElemType: types.StringType},
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},

												"time_offset": schema.StringAttribute{
													Description:         "TimeOffset defines the delta time of injected program. It's a possibly signed sequence of decimal numbers, such as '300ms', '-1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
													MarkdownDescription: "TimeOffset defines the delta time of injected program. It's a possibly signed sequence of decimal numbers, such as '300ms', '-1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
													MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
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
											Description:         "TODO: use a custom type, as 'TemplateType' contains other possible values",
											MarkdownDescription: "TODO: use a custom type, as 'TemplateType' contains other possible values",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"status_check": schema.SingleNestedAttribute{
									Description:         "StatusCheck describe the behavior of StatusCheck. Only used when Type is TypeStatusCheck.",
									MarkdownDescription: "StatusCheck describe the behavior of StatusCheck. Only used when Type is TypeStatusCheck.",
									Attributes: map[string]schema.Attribute{
										"duration": schema.StringAttribute{
											Description:         "Duration defines the duration of the whole status check if the number of failed execution does not exceed the failure threshold. Duration is available to both 'Synchronous' and 'Continuous' mode. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '-1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
											MarkdownDescription: "Duration defines the duration of the whole status check if the number of failed execution does not exceed the failure threshold. Duration is available to both 'Synchronous' and 'Continuous' mode. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '-1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"failure_threshold": schema.Int64Attribute{
											Description:         "FailureThreshold defines the minimum consecutive failure for the status check to be considered failed.",
											MarkdownDescription: "FailureThreshold defines the minimum consecutive failure for the status check to be considered failed.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(1),
											},
										},

										"http": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"body": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"criteria": schema.SingleNestedAttribute{
													Description:         "Criteria defines how to determine the result of the status check.",
													MarkdownDescription: "Criteria defines how to determine the result of the status check.",
													Attributes: map[string]schema.Attribute{
														"status_code": schema.StringAttribute{
															Description:         "StatusCode defines the expected http status code for the request. A statusCode string could be a single code (e.g. 200), or an inclusive range (e.g. 200-400, both '200' and '400' are included).",
															MarkdownDescription: "StatusCode defines the expected http status code for the request. A statusCode string could be a single code (e.g. 200), or an inclusive range (e.g. 200-400, both '200' and '400' are included).",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},

												"headers": schema.MapAttribute{
													Description:         "A Header represents the key-value pairs in an HTTP header.  The keys should be in canonical form, as returned by CanonicalHeaderKey.",
													MarkdownDescription: "A Header represents the key-value pairs in an HTTP header.  The keys should be in canonical form, as returned by CanonicalHeaderKey.",
													ElementType:         types.ListType{ElemType: types.StringType},
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"method": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("GET", "POST"),
													},
												},

												"url": schema.StringAttribute{
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

										"interval_seconds": schema.Int64Attribute{
											Description:         "IntervalSeconds defines how often (in seconds) to perform an execution of status check.",
											MarkdownDescription: "IntervalSeconds defines how often (in seconds) to perform an execution of status check.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(1),
											},
										},

										"mode": schema.StringAttribute{
											Description:         "Mode defines the execution mode of the status check. Support type: Synchronous / Continuous",
											MarkdownDescription: "Mode defines the execution mode of the status check. Support type: Synchronous / Continuous",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Synchronous", "Continuous"),
											},
										},

										"records_history_limit": schema.Int64Attribute{
											Description:         "RecordsHistoryLimit defines the number of record to retain.",
											MarkdownDescription: "RecordsHistoryLimit defines the number of record to retain.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(1),
												int64validator.AtMost(1000),
											},
										},

										"success_threshold": schema.Int64Attribute{
											Description:         "SuccessThreshold defines the minimum consecutive successes for the status check to be considered successful. SuccessThreshold only works for 'Synchronous' mode.",
											MarkdownDescription: "SuccessThreshold defines the minimum consecutive successes for the status check to be considered successful. SuccessThreshold only works for 'Synchronous' mode.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(1),
											},
										},

										"timeout_seconds": schema.Int64Attribute{
											Description:         "TimeoutSeconds defines the number of seconds after which an execution of status check times out.",
											MarkdownDescription: "TimeoutSeconds defines the number of seconds after which an execution of status check times out.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(1),
											},
										},

										"type": schema.StringAttribute{
											Description:         "Type defines the specific status check type. Support type: HTTP",
											MarkdownDescription: "Type defines the specific status check type. Support type: HTTP",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("HTTP"),
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"stress_chaos": schema.SingleNestedAttribute{
									Description:         "StressChaosSpec defines the desired state of StressChaos",
									MarkdownDescription: "StressChaosSpec defines the desired state of StressChaos",
									Attributes: map[string]schema.Attribute{
										"container_names": schema.ListAttribute{
											Description:         "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
											MarkdownDescription: "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"duration": schema.StringAttribute{
											Description:         "Duration represents the duration of the chaos action",
											MarkdownDescription: "Duration represents the duration of the chaos action",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"mode": schema.StringAttribute{
											Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
											MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
											},
										},

										"remote_cluster": schema.StringAttribute{
											Description:         "RemoteCluster represents the remote cluster where the chaos will be deployed",
											MarkdownDescription: "RemoteCluster represents the remote cluster where the chaos will be deployed",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"selector": schema.SingleNestedAttribute{
											Description:         "Selector is used to select pods that are used to inject chaos action.",
											MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",
											Attributes: map[string]schema.Attribute{
												"annotation_selectors": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
													MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"expression_selectors": schema.ListNestedAttribute{
													Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
													MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
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

												"field_selectors": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
													MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"label_selectors": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
													MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespaces": schema.ListAttribute{
													Description:         "Namespaces is a set of namespace to which objects belong.",
													MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"node_selectors": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
													MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"nodes": schema.ListAttribute{
													Description:         "Nodes is a set of node name and objects must belong to these nodes.",
													MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"pod_phase_selectors": schema.ListAttribute{
													Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
													MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"pods": schema.MapAttribute{
													Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
													MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
													ElementType:         types.ListType{ElemType: types.StringType},
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"stressng_stressors": schema.StringAttribute{
											Description:         "StressngStressors defines plenty of stressors just like 'Stressors' except that it's an experimental feature and more powerful. You can define stressors in 'stress-ng' (see also 'man stress-ng') dialect, however not all of the supported stressors are well tested. It maybe retired in later releases. You should always use 'Stressors' to define the stressors and use this only when you want more stressors unsupported by 'Stressors'. When both 'StressngStressors' and 'Stressors' are defined, 'StressngStressors' wins.",
											MarkdownDescription: "StressngStressors defines plenty of stressors just like 'Stressors' except that it's an experimental feature and more powerful. You can define stressors in 'stress-ng' (see also 'man stress-ng') dialect, however not all of the supported stressors are well tested. It maybe retired in later releases. You should always use 'Stressors' to define the stressors and use this only when you want more stressors unsupported by 'Stressors'. When both 'StressngStressors' and 'Stressors' are defined, 'StressngStressors' wins.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"stressors": schema.SingleNestedAttribute{
											Description:         "Stressors defines plenty of stressors supported to stress system components out. You can use one or more of them to make up various kinds of stresses. At least one of the stressors should be specified.",
											MarkdownDescription: "Stressors defines plenty of stressors supported to stress system components out. You can use one or more of them to make up various kinds of stresses. At least one of the stressors should be specified.",
											Attributes: map[string]schema.Attribute{
												"cpu": schema.SingleNestedAttribute{
													Description:         "CPUStressor stresses CPU out",
													MarkdownDescription: "CPUStressor stresses CPU out",
													Attributes: map[string]schema.Attribute{
														"load": schema.Int64Attribute{
															Description:         "Load specifies P percent loading per CPU worker. 0 is effectively a sleep (no load) and 100 is full loading.",
															MarkdownDescription: "Load specifies P percent loading per CPU worker. 0 is effectively a sleep (no load) and 100 is full loading.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(0),
																int64validator.AtMost(100),
															},
														},

														"options": schema.ListAttribute{
															Description:         "extend stress-ng options",
															MarkdownDescription: "extend stress-ng options",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"workers": schema.Int64Attribute{
															Description:         "Workers specifies N workers to apply the stressor. Maximum 8192 workers can run by stress-ng",
															MarkdownDescription: "Workers specifies N workers to apply the stressor. Maximum 8192 workers can run by stress-ng",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtMost(8192),
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"memory": schema.SingleNestedAttribute{
													Description:         "MemoryStressor stresses virtual memory out",
													MarkdownDescription: "MemoryStressor stresses virtual memory out",
													Attributes: map[string]schema.Attribute{
														"oom_score_adj": schema.Int64Attribute{
															Description:         "OOMScoreAdj sets the oom_score_adj of the stress process. See 'man 5 proc' to know more about this option.",
															MarkdownDescription: "OOMScoreAdj sets the oom_score_adj of the stress process. See 'man 5 proc' to know more about this option.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(-1000),
																int64validator.AtMost(1000),
															},
														},

														"options": schema.ListAttribute{
															Description:         "extend stress-ng options",
															MarkdownDescription: "extend stress-ng options",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"size": schema.StringAttribute{
															Description:         "Size specifies N bytes consumed per vm worker, default is the total available memory. One can specify the size as % of total available memory or in units of B, KB/KiB, MB/MiB, GB/GiB, TB/TiB.",
															MarkdownDescription: "Size specifies N bytes consumed per vm worker, default is the total available memory. One can specify the size as % of total available memory or in units of B, KB/KiB, MB/MiB, GB/GiB, TB/TiB.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"workers": schema.Int64Attribute{
															Description:         "Workers specifies N workers to apply the stressor. Maximum 8192 workers can run by stress-ng",
															MarkdownDescription: "Workers specifies N workers to apply the stressor. Maximum 8192 workers can run by stress-ng",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtMost(8192),
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

										"value": schema.StringAttribute{
											Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
											MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"task": schema.SingleNestedAttribute{
									Description:         "Task describes the behavior of the custom task. Only used when Type is TypeTask.",
									MarkdownDescription: "Task describes the behavior of the custom task. Only used when Type is TypeTask.",
									Attributes: map[string]schema.Attribute{
										"container": schema.SingleNestedAttribute{
											Description:         "Container is the main container image to run in the pod",
											MarkdownDescription: "Container is the main container image to run in the pod",
											Attributes: map[string]schema.Attribute{
												"args": schema.ListAttribute{
													Description:         "Arguments to the entrypoint. The container image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
													MarkdownDescription: "Arguments to the entrypoint. The container image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"command": schema.ListAttribute{
													Description:         "Entrypoint array. Not executed within a shell. The container image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
													MarkdownDescription: "Entrypoint array. Not executed within a shell. The container image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"env": schema.ListNestedAttribute{
													Description:         "List of environment variables to set in the container. Cannot be updated.",
													MarkdownDescription: "List of environment variables to set in the container. Cannot be updated.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
																MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"value": schema.StringAttribute{
																Description:         "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
																MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"value_from": schema.SingleNestedAttribute{
																Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
																MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",
																Attributes: map[string]schema.Attribute{
																	"config_map_key_ref": schema.SingleNestedAttribute{
																		Description:         "Selects a key of a ConfigMap.",
																		MarkdownDescription: "Selects a key of a ConfigMap.",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The key to select.",
																				MarkdownDescription: "The key to select.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"optional": schema.BoolAttribute{
																				Description:         "Specify whether the ConfigMap or its key must be defined",
																				MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"field_ref": schema.SingleNestedAttribute{
																		Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																		MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																		Attributes: map[string]schema.Attribute{
																			"api_version": schema.StringAttribute{
																				Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																				MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"field_path": schema.StringAttribute{
																				Description:         "Path of the field to select in the specified API version.",
																				MarkdownDescription: "Path of the field to select in the specified API version.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"resource_field_ref": schema.SingleNestedAttribute{
																		Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																		MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																		Attributes: map[string]schema.Attribute{
																			"container_name": schema.StringAttribute{
																				Description:         "Container name: required for volumes, optional for env vars",
																				MarkdownDescription: "Container name: required for volumes, optional for env vars",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"divisor": schema.StringAttribute{
																				Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																				MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"resource": schema.StringAttribute{
																				Description:         "Required: resource to select",
																				MarkdownDescription: "Required: resource to select",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"secret_key_ref": schema.SingleNestedAttribute{
																		Description:         "Selects a key of a secret in the pod's namespace",
																		MarkdownDescription: "Selects a key of a secret in the pod's namespace",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The key of the secret to select from.  Must be a valid secret key.",
																				MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"optional": schema.BoolAttribute{
																				Description:         "Specify whether the Secret or its key must be defined",
																				MarkdownDescription: "Specify whether the Secret or its key must be defined",
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
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"env_from": schema.ListNestedAttribute{
													Description:         "List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER. All invalid keys will be reported as an event when the container is starting. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated.",
													MarkdownDescription: "List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER. All invalid keys will be reported as an event when the container is starting. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"config_map_ref": schema.SingleNestedAttribute{
																Description:         "The ConfigMap to select from",
																MarkdownDescription: "The ConfigMap to select from",
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"optional": schema.BoolAttribute{
																		Description:         "Specify whether the ConfigMap must be defined",
																		MarkdownDescription: "Specify whether the ConfigMap must be defined",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"prefix": schema.StringAttribute{
																Description:         "An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.",
																MarkdownDescription: "An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_ref": schema.SingleNestedAttribute{
																Description:         "The Secret to select from",
																MarkdownDescription: "The Secret to select from",
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"optional": schema.BoolAttribute{
																		Description:         "Specify whether the Secret must be defined",
																		MarkdownDescription: "Specify whether the Secret must be defined",
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

												"image": schema.StringAttribute{
													Description:         "Container image name. More info: https://kubernetes.io/docs/concepts/containers/images This field is optional to allow higher level config management to default or override container images in workload controllers like Deployments and StatefulSets.",
													MarkdownDescription: "Container image name. More info: https://kubernetes.io/docs/concepts/containers/images This field is optional to allow higher level config management to default or override container images in workload controllers like Deployments and StatefulSets.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_pull_policy": schema.StringAttribute{
													Description:         "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",
													MarkdownDescription: "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"lifecycle": schema.SingleNestedAttribute{
													Description:         "Actions that the management system should take in response to container lifecycle events. Cannot be updated.",
													MarkdownDescription: "Actions that the management system should take in response to container lifecycle events. Cannot be updated.",
													Attributes: map[string]schema.Attribute{
														"post_start": schema.SingleNestedAttribute{
															Description:         "PostStart is called immediately after a container is created. If the handler fails, the container is terminated and restarted according to its restart policy. Other management of the container blocks until the hook completes. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",
															MarkdownDescription: "PostStart is called immediately after a container is created. If the handler fails, the container is terminated and restarted according to its restart policy. Other management of the container blocks until the hook completes. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",
															Attributes: map[string]schema.Attribute{
																"exec": schema.SingleNestedAttribute{
																	Description:         "Exec specifies the action to take.",
																	MarkdownDescription: "Exec specifies the action to take.",
																	Attributes: map[string]schema.Attribute{
																		"command": schema.ListAttribute{
																			Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																			MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
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

																"http_get": schema.SingleNestedAttribute{
																	Description:         "HTTPGet specifies the http request to perform.",
																	MarkdownDescription: "HTTPGet specifies the http request to perform.",
																	Attributes: map[string]schema.Attribute{
																		"host": schema.StringAttribute{
																			Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																			MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"http_headers": schema.ListNestedAttribute{
																			Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																			MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																						MarkdownDescription: "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"value": schema.StringAttribute{
																						Description:         "The header field value",
																						MarkdownDescription: "The header field value",
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

																		"path": schema.StringAttribute{
																			Description:         "Path to access on the HTTP server.",
																			MarkdownDescription: "Path to access on the HTTP server.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"port": schema.StringAttribute{
																			Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																			MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"scheme": schema.StringAttribute{
																			Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																			MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"tcp_socket": schema.SingleNestedAttribute{
																	Description:         "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",
																	MarkdownDescription: "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",
																	Attributes: map[string]schema.Attribute{
																		"host": schema.StringAttribute{
																			Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																			MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"port": schema.StringAttribute{
																			Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																			MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
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

														"pre_stop": schema.SingleNestedAttribute{
															Description:         "PreStop is called immediately before a container is terminated due to an API request or management event such as liveness/startup probe failure, preemption, resource contention, etc. The handler is not called if the container crashes or exits. The Pod's termination grace period countdown begins before the PreStop hook is executed. Regardless of the outcome of the handler, the container will eventually terminate within the Pod's termination grace period (unless delayed by finalizers). Other management of the container blocks until the hook completes or until the termination grace period is reached. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",
															MarkdownDescription: "PreStop is called immediately before a container is terminated due to an API request or management event such as liveness/startup probe failure, preemption, resource contention, etc. The handler is not called if the container crashes or exits. The Pod's termination grace period countdown begins before the PreStop hook is executed. Regardless of the outcome of the handler, the container will eventually terminate within the Pod's termination grace period (unless delayed by finalizers). Other management of the container blocks until the hook completes or until the termination grace period is reached. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",
															Attributes: map[string]schema.Attribute{
																"exec": schema.SingleNestedAttribute{
																	Description:         "Exec specifies the action to take.",
																	MarkdownDescription: "Exec specifies the action to take.",
																	Attributes: map[string]schema.Attribute{
																		"command": schema.ListAttribute{
																			Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																			MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
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

																"http_get": schema.SingleNestedAttribute{
																	Description:         "HTTPGet specifies the http request to perform.",
																	MarkdownDescription: "HTTPGet specifies the http request to perform.",
																	Attributes: map[string]schema.Attribute{
																		"host": schema.StringAttribute{
																			Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																			MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"http_headers": schema.ListNestedAttribute{
																			Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																			MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																						MarkdownDescription: "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"value": schema.StringAttribute{
																						Description:         "The header field value",
																						MarkdownDescription: "The header field value",
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

																		"path": schema.StringAttribute{
																			Description:         "Path to access on the HTTP server.",
																			MarkdownDescription: "Path to access on the HTTP server.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"port": schema.StringAttribute{
																			Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																			MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"scheme": schema.StringAttribute{
																			Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																			MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"tcp_socket": schema.SingleNestedAttribute{
																	Description:         "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",
																	MarkdownDescription: "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",
																	Attributes: map[string]schema.Attribute{
																		"host": schema.StringAttribute{
																			Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																			MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"port": schema.StringAttribute{
																			Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																			MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
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
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"liveness_probe": schema.SingleNestedAttribute{
													Description:         "Periodic probe of container liveness. Container will be restarted if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
													MarkdownDescription: "Periodic probe of container liveness. Container will be restarted if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
													Attributes: map[string]schema.Attribute{
														"exec": schema.SingleNestedAttribute{
															Description:         "Exec specifies the action to take.",
															MarkdownDescription: "Exec specifies the action to take.",
															Attributes: map[string]schema.Attribute{
																"command": schema.ListAttribute{
																	Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																	MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
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

														"failure_threshold": schema.Int64Attribute{
															Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
															MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"grpc": schema.SingleNestedAttribute{
															Description:         "GRPC specifies an action involving a GRPC port. This is a beta field and requires enabling GRPCContainerProbe feature gate.",
															MarkdownDescription: "GRPC specifies an action involving a GRPC port. This is a beta field and requires enabling GRPCContainerProbe feature gate.",
															Attributes: map[string]schema.Attribute{
																"port": schema.Int64Attribute{
																	Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																	MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"service": schema.StringAttribute{
																	Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",
																	MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"http_get": schema.SingleNestedAttribute{
															Description:         "HTTPGet specifies the http request to perform.",
															MarkdownDescription: "HTTPGet specifies the http request to perform.",
															Attributes: map[string]schema.Attribute{
																"host": schema.StringAttribute{
																	Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																	MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"http_headers": schema.ListNestedAttribute{
																	Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																	MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																				MarkdownDescription: "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"value": schema.StringAttribute{
																				Description:         "The header field value",
																				MarkdownDescription: "The header field value",
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

																"path": schema.StringAttribute{
																	Description:         "Path to access on the HTTP server.",
																	MarkdownDescription: "Path to access on the HTTP server.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"port": schema.StringAttribute{
																	Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																	MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"scheme": schema.StringAttribute{
																	Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																	MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"initial_delay_seconds": schema.Int64Attribute{
															Description:         "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
															MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"period_seconds": schema.Int64Attribute{
															Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
															MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"success_threshold": schema.Int64Attribute{
															Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
															MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"tcp_socket": schema.SingleNestedAttribute{
															Description:         "TCPSocket specifies an action involving a TCP port.",
															MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",
															Attributes: map[string]schema.Attribute{
																"host": schema.StringAttribute{
																	Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																	MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"port": schema.StringAttribute{
																	Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																	MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"termination_grace_period_seconds": schema.Int64Attribute{
															Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
															MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"timeout_seconds": schema.Int64Attribute{
															Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
															MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.",
													MarkdownDescription: "Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"ports": schema.ListNestedAttribute{
													Description:         "List of ports to expose from the container. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default '0.0.0.0' address inside a container will be accessible from the network. Modifying this array with strategic merge patch may corrupt the data. For more information See https://github.com/kubernetes/kubernetes/issues/108255. Cannot be updated.",
													MarkdownDescription: "List of ports to expose from the container. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default '0.0.0.0' address inside a container will be accessible from the network. Modifying this array with strategic merge patch may corrupt the data. For more information See https://github.com/kubernetes/kubernetes/issues/108255. Cannot be updated.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"container_port": schema.Int64Attribute{
																Description:         "Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.",
																MarkdownDescription: "Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"host_ip": schema.StringAttribute{
																Description:         "What host IP to bind the external port to.",
																MarkdownDescription: "What host IP to bind the external port to.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"host_port": schema.Int64Attribute{
																Description:         "Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.",
																MarkdownDescription: "Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services.",
																MarkdownDescription: "If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"protocol": schema.StringAttribute{
																Description:         "Protocol for port. Must be UDP, TCP, or SCTP. Defaults to 'TCP'.",
																MarkdownDescription: "Protocol for port. Must be UDP, TCP, or SCTP. Defaults to 'TCP'.",
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

												"readiness_probe": schema.SingleNestedAttribute{
													Description:         "Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
													MarkdownDescription: "Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
													Attributes: map[string]schema.Attribute{
														"exec": schema.SingleNestedAttribute{
															Description:         "Exec specifies the action to take.",
															MarkdownDescription: "Exec specifies the action to take.",
															Attributes: map[string]schema.Attribute{
																"command": schema.ListAttribute{
																	Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																	MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
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

														"failure_threshold": schema.Int64Attribute{
															Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
															MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"grpc": schema.SingleNestedAttribute{
															Description:         "GRPC specifies an action involving a GRPC port. This is a beta field and requires enabling GRPCContainerProbe feature gate.",
															MarkdownDescription: "GRPC specifies an action involving a GRPC port. This is a beta field and requires enabling GRPCContainerProbe feature gate.",
															Attributes: map[string]schema.Attribute{
																"port": schema.Int64Attribute{
																	Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																	MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"service": schema.StringAttribute{
																	Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",
																	MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"http_get": schema.SingleNestedAttribute{
															Description:         "HTTPGet specifies the http request to perform.",
															MarkdownDescription: "HTTPGet specifies the http request to perform.",
															Attributes: map[string]schema.Attribute{
																"host": schema.StringAttribute{
																	Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																	MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"http_headers": schema.ListNestedAttribute{
																	Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																	MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																				MarkdownDescription: "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"value": schema.StringAttribute{
																				Description:         "The header field value",
																				MarkdownDescription: "The header field value",
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

																"path": schema.StringAttribute{
																	Description:         "Path to access on the HTTP server.",
																	MarkdownDescription: "Path to access on the HTTP server.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"port": schema.StringAttribute{
																	Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																	MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"scheme": schema.StringAttribute{
																	Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																	MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"initial_delay_seconds": schema.Int64Attribute{
															Description:         "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
															MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"period_seconds": schema.Int64Attribute{
															Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
															MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"success_threshold": schema.Int64Attribute{
															Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
															MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"tcp_socket": schema.SingleNestedAttribute{
															Description:         "TCPSocket specifies an action involving a TCP port.",
															MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",
															Attributes: map[string]schema.Attribute{
																"host": schema.StringAttribute{
																	Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																	MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"port": schema.StringAttribute{
																	Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																	MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"termination_grace_period_seconds": schema.Int64Attribute{
															Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
															MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"timeout_seconds": schema.Int64Attribute{
															Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
															MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"resources": schema.SingleNestedAttribute{
													Description:         "Compute Resources required by this container. Cannot be updated. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
													MarkdownDescription: "Compute Resources required by this container. Cannot be updated. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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
															Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
															MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

												"security_context": schema.SingleNestedAttribute{
													Description:         "SecurityContext defines the security options the container should be run with. If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext. More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/",
													MarkdownDescription: "SecurityContext defines the security options the container should be run with. If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext. More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/",
													Attributes: map[string]schema.Attribute{
														"allow_privilege_escalation": schema.BoolAttribute{
															Description:         "AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN Note that this field cannot be set when spec.os.name is windows.",
															MarkdownDescription: "AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN Note that this field cannot be set when spec.os.name is windows.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"capabilities": schema.SingleNestedAttribute{
															Description:         "The capabilities to add/drop when running containers. Defaults to the default set of capabilities granted by the container runtime. Note that this field cannot be set when spec.os.name is windows.",
															MarkdownDescription: "The capabilities to add/drop when running containers. Defaults to the default set of capabilities granted by the container runtime. Note that this field cannot be set when spec.os.name is windows.",
															Attributes: map[string]schema.Attribute{
																"add": schema.ListAttribute{
																	Description:         "Added capabilities",
																	MarkdownDescription: "Added capabilities",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"drop": schema.ListAttribute{
																	Description:         "Removed capabilities",
																	MarkdownDescription: "Removed capabilities",
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

														"privileged": schema.BoolAttribute{
															Description:         "Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false. Note that this field cannot be set when spec.os.name is windows.",
															MarkdownDescription: "Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false. Note that this field cannot be set when spec.os.name is windows.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"proc_mount": schema.StringAttribute{
															Description:         "procMount denotes the type of proc mount to use for the containers. The default is DefaultProcMount which uses the container runtime defaults for readonly paths and masked paths. This requires the ProcMountType feature flag to be enabled. Note that this field cannot be set when spec.os.name is windows.",
															MarkdownDescription: "procMount denotes the type of proc mount to use for the containers. The default is DefaultProcMount which uses the container runtime defaults for readonly paths and masked paths. This requires the ProcMountType feature flag to be enabled. Note that this field cannot be set when spec.os.name is windows.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"read_only_root_filesystem": schema.BoolAttribute{
															Description:         "Whether this container has a read-only root filesystem. Default is false. Note that this field cannot be set when spec.os.name is windows.",
															MarkdownDescription: "Whether this container has a read-only root filesystem. Default is false. Note that this field cannot be set when spec.os.name is windows.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"run_as_group": schema.Int64Attribute{
															Description:         "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",
															MarkdownDescription: "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"run_as_non_root": schema.BoolAttribute{
															Description:         "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
															MarkdownDescription: "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"run_as_user": schema.Int64Attribute{
															Description:         "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",
															MarkdownDescription: "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"se_linux_options": schema.SingleNestedAttribute{
															Description:         "The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",
															MarkdownDescription: "The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",
															Attributes: map[string]schema.Attribute{
																"level": schema.StringAttribute{
																	Description:         "Level is SELinux level label that applies to the container.",
																	MarkdownDescription: "Level is SELinux level label that applies to the container.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"role": schema.StringAttribute{
																	Description:         "Role is a SELinux role label that applies to the container.",
																	MarkdownDescription: "Role is a SELinux role label that applies to the container.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"type": schema.StringAttribute{
																	Description:         "Type is a SELinux type label that applies to the container.",
																	MarkdownDescription: "Type is a SELinux type label that applies to the container.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"user": schema.StringAttribute{
																	Description:         "User is a SELinux user label that applies to the container.",
																	MarkdownDescription: "User is a SELinux user label that applies to the container.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"seccomp_profile": schema.SingleNestedAttribute{
															Description:         "The seccomp options to use by this container. If seccomp options are provided at both the pod & container level, the container options override the pod options. Note that this field cannot be set when spec.os.name is windows.",
															MarkdownDescription: "The seccomp options to use by this container. If seccomp options are provided at both the pod & container level, the container options override the pod options. Note that this field cannot be set when spec.os.name is windows.",
															Attributes: map[string]schema.Attribute{
																"localhost_profile": schema.StringAttribute{
																	Description:         "localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work. Must be a descending path, relative to the kubelet's configured seccomp profile location. Must only be set if type is 'Localhost'.",
																	MarkdownDescription: "localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work. Must be a descending path, relative to the kubelet's configured seccomp profile location. Must only be set if type is 'Localhost'.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"type": schema.StringAttribute{
																	Description:         "type indicates which kind of seccomp profile will be applied. Valid options are:  Localhost - a profile defined in a file on the node should be used. RuntimeDefault - the container runtime default profile should be used. Unconfined - no profile should be applied.",
																	MarkdownDescription: "type indicates which kind of seccomp profile will be applied. Valid options are:  Localhost - a profile defined in a file on the node should be used. RuntimeDefault - the container runtime default profile should be used. Unconfined - no profile should be applied.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"windows_options": schema.SingleNestedAttribute{
															Description:         "The Windows specific settings applied to all containers. If unspecified, the options from the PodSecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is linux.",
															MarkdownDescription: "The Windows specific settings applied to all containers. If unspecified, the options from the PodSecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is linux.",
															Attributes: map[string]schema.Attribute{
																"gmsa_credential_spec": schema.StringAttribute{
																	Description:         "GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.",
																	MarkdownDescription: "GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"gmsa_credential_spec_name": schema.StringAttribute{
																	Description:         "GMSACredentialSpecName is the name of the GMSA credential spec to use.",
																	MarkdownDescription: "GMSACredentialSpecName is the name of the GMSA credential spec to use.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"host_process": schema.BoolAttribute{
																	Description:         "HostProcess determines if a container should be run as a 'Host Process' container. This field is alpha-level and will only be honored by components that enable the WindowsHostProcessContainers feature flag. Setting this field without the feature flag will result in errors when validating the Pod. All of a Pod's containers must have the same effective HostProcess value (it is not allowed to have a mix of HostProcess containers and non-HostProcess containers).  In addition, if HostProcess is true then HostNetwork must also be set to true.",
																	MarkdownDescription: "HostProcess determines if a container should be run as a 'Host Process' container. This field is alpha-level and will only be honored by components that enable the WindowsHostProcessContainers feature flag. Setting this field without the feature flag will result in errors when validating the Pod. All of a Pod's containers must have the same effective HostProcess value (it is not allowed to have a mix of HostProcess containers and non-HostProcess containers).  In addition, if HostProcess is true then HostNetwork must also be set to true.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"run_as_user_name": schema.StringAttribute{
																	Description:         "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
																	MarkdownDescription: "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
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

												"startup_probe": schema.SingleNestedAttribute{
													Description:         "StartupProbe indicates that the Pod has successfully initialized. If specified, no other probes are executed until this completes successfully. If this probe fails, the Pod will be restarted, just as if the livenessProbe failed. This can be used to provide different probe parameters at the beginning of a Pod's lifecycle, when it might take a long time to load data or warm a cache, than during steady-state operation. This cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
													MarkdownDescription: "StartupProbe indicates that the Pod has successfully initialized. If specified, no other probes are executed until this completes successfully. If this probe fails, the Pod will be restarted, just as if the livenessProbe failed. This can be used to provide different probe parameters at the beginning of a Pod's lifecycle, when it might take a long time to load data or warm a cache, than during steady-state operation. This cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
													Attributes: map[string]schema.Attribute{
														"exec": schema.SingleNestedAttribute{
															Description:         "Exec specifies the action to take.",
															MarkdownDescription: "Exec specifies the action to take.",
															Attributes: map[string]schema.Attribute{
																"command": schema.ListAttribute{
																	Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																	MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
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

														"failure_threshold": schema.Int64Attribute{
															Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
															MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"grpc": schema.SingleNestedAttribute{
															Description:         "GRPC specifies an action involving a GRPC port. This is a beta field and requires enabling GRPCContainerProbe feature gate.",
															MarkdownDescription: "GRPC specifies an action involving a GRPC port. This is a beta field and requires enabling GRPCContainerProbe feature gate.",
															Attributes: map[string]schema.Attribute{
																"port": schema.Int64Attribute{
																	Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																	MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"service": schema.StringAttribute{
																	Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",
																	MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"http_get": schema.SingleNestedAttribute{
															Description:         "HTTPGet specifies the http request to perform.",
															MarkdownDescription: "HTTPGet specifies the http request to perform.",
															Attributes: map[string]schema.Attribute{
																"host": schema.StringAttribute{
																	Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																	MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"http_headers": schema.ListNestedAttribute{
																	Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																	MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																				MarkdownDescription: "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"value": schema.StringAttribute{
																				Description:         "The header field value",
																				MarkdownDescription: "The header field value",
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

																"path": schema.StringAttribute{
																	Description:         "Path to access on the HTTP server.",
																	MarkdownDescription: "Path to access on the HTTP server.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"port": schema.StringAttribute{
																	Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																	MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"scheme": schema.StringAttribute{
																	Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																	MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"initial_delay_seconds": schema.Int64Attribute{
															Description:         "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
															MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"period_seconds": schema.Int64Attribute{
															Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
															MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"success_threshold": schema.Int64Attribute{
															Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
															MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"tcp_socket": schema.SingleNestedAttribute{
															Description:         "TCPSocket specifies an action involving a TCP port.",
															MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",
															Attributes: map[string]schema.Attribute{
																"host": schema.StringAttribute{
																	Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																	MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"port": schema.StringAttribute{
																	Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																	MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"termination_grace_period_seconds": schema.Int64Attribute{
															Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
															MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"timeout_seconds": schema.Int64Attribute{
															Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
															MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"stdin": schema.BoolAttribute{
													Description:         "Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false.",
													MarkdownDescription: "Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"stdin_once": schema.BoolAttribute{
													Description:         "Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false",
													MarkdownDescription: "Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"termination_message_path": schema.StringAttribute{
													Description:         "Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to 12kb. Defaults to /dev/termination-log. Cannot be updated.",
													MarkdownDescription: "Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to 12kb. Defaults to /dev/termination-log. Cannot be updated.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"termination_message_policy": schema.StringAttribute{
													Description:         "Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.",
													MarkdownDescription: "Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"tty": schema.BoolAttribute{
													Description:         "Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false.",
													MarkdownDescription: "Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"volume_devices": schema.ListNestedAttribute{
													Description:         "volumeDevices is the list of block devices to be used by the container.",
													MarkdownDescription: "volumeDevices is the list of block devices to be used by the container.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"device_path": schema.StringAttribute{
																Description:         "devicePath is the path inside of the container that the device will be mapped to.",
																MarkdownDescription: "devicePath is the path inside of the container that the device will be mapped to.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "name must match the name of a persistentVolumeClaim in the pod",
																MarkdownDescription: "name must match the name of a persistentVolumeClaim in the pod",
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

												"volume_mounts": schema.ListNestedAttribute{
													Description:         "Pod volumes to mount into the container's filesystem. Cannot be updated.",
													MarkdownDescription: "Pod volumes to mount into the container's filesystem. Cannot be updated.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"mount_path": schema.StringAttribute{
																Description:         "Path within the container at which the volume should be mounted.  Must not contain ':'.",
																MarkdownDescription: "Path within the container at which the volume should be mounted.  Must not contain ':'.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"mount_propagation": schema.StringAttribute{
																Description:         "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.",
																MarkdownDescription: "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "This must match the Name of a Volume.",
																MarkdownDescription: "This must match the Name of a Volume.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"read_only": schema.BoolAttribute{
																Description:         "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",
																MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"sub_path": schema.StringAttribute{
																Description:         "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",
																MarkdownDescription: "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"sub_path_expr": schema.StringAttribute{
																Description:         "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",
																MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",
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

												"working_dir": schema.StringAttribute{
													Description:         "Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.",
													MarkdownDescription: "Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"volumes": schema.ListNestedAttribute{
											Description:         "Volumes is a list of volumes that can be mounted by containers in a template.",
											MarkdownDescription: "Volumes is a list of volumes that can be mounted by containers in a template.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"aws_elastic_block_store": schema.SingleNestedAttribute{
														Description:         "awsElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
														MarkdownDescription: "awsElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
														Attributes: map[string]schema.Attribute{
															"fs_type": schema.StringAttribute{
																Description:         "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore TODO: how do we prevent errors in the filesystem from compromising the machine",
																MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore TODO: how do we prevent errors in the filesystem from compromising the machine",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"partition": schema.Int64Attribute{
																Description:         "partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty).",
																MarkdownDescription: "partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty).",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"read_only": schema.BoolAttribute{
																Description:         "readOnly value true will force the readOnly setting in VolumeMounts. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
																MarkdownDescription: "readOnly value true will force the readOnly setting in VolumeMounts. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"volume_id": schema.StringAttribute{
																Description:         "volumeID is unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
																MarkdownDescription: "volumeID is unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"azure_disk": schema.SingleNestedAttribute{
														Description:         "azureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.",
														MarkdownDescription: "azureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.",
														Attributes: map[string]schema.Attribute{
															"caching_mode": schema.StringAttribute{
																Description:         "cachingMode is the Host Caching mode: None, Read Only, Read Write.",
																MarkdownDescription: "cachingMode is the Host Caching mode: None, Read Only, Read Write.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"disk_name": schema.StringAttribute{
																Description:         "diskName is the Name of the data disk in the blob storage",
																MarkdownDescription: "diskName is the Name of the data disk in the blob storage",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"disk_uri": schema.StringAttribute{
																Description:         "diskURI is the URI of data disk in the blob storage",
																MarkdownDescription: "diskURI is the URI of data disk in the blob storage",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"fs_type": schema.StringAttribute{
																Description:         "fsType is Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																MarkdownDescription: "fsType is Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kind": schema.StringAttribute{
																Description:         "kind expected values are Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared",
																MarkdownDescription: "kind expected values are Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"read_only": schema.BoolAttribute{
																Description:         "readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																MarkdownDescription: "readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"azure_file": schema.SingleNestedAttribute{
														Description:         "azureFile represents an Azure File Service mount on the host and bind mount to the pod.",
														MarkdownDescription: "azureFile represents an Azure File Service mount on the host and bind mount to the pod.",
														Attributes: map[string]schema.Attribute{
															"read_only": schema.BoolAttribute{
																Description:         "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_name": schema.StringAttribute{
																Description:         "secretName is the  name of secret that contains Azure Storage Account Name and Key",
																MarkdownDescription: "secretName is the  name of secret that contains Azure Storage Account Name and Key",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"share_name": schema.StringAttribute{
																Description:         "shareName is the azure share Name",
																MarkdownDescription: "shareName is the azure share Name",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"cephfs": schema.SingleNestedAttribute{
														Description:         "cephFS represents a Ceph FS mount on the host that shares a pod's lifetime",
														MarkdownDescription: "cephFS represents a Ceph FS mount on the host that shares a pod's lifetime",
														Attributes: map[string]schema.Attribute{
															"monitors": schema.ListAttribute{
																Description:         "monitors is Required: Monitors is a collection of Ceph monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																MarkdownDescription: "monitors is Required: Monitors is a collection of Ceph monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																ElementType:         types.StringType,
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"path": schema.StringAttribute{
																Description:         "path is Optional: Used as the mounted root, rather than the full Ceph tree, default is /",
																MarkdownDescription: "path is Optional: Used as the mounted root, rather than the full Ceph tree, default is /",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"read_only": schema.BoolAttribute{
																Description:         "readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																MarkdownDescription: "readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_file": schema.StringAttribute{
																Description:         "secretFile is Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																MarkdownDescription: "secretFile is Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_ref": schema.SingleNestedAttribute{
																Description:         "secretRef is Optional: SecretRef is reference to the authentication secret for User, default is empty. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																MarkdownDescription: "secretRef is Optional: SecretRef is reference to the authentication secret for User, default is empty. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"user": schema.StringAttribute{
																Description:         "user is optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																MarkdownDescription: "user is optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"cinder": schema.SingleNestedAttribute{
														Description:         "cinder represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
														MarkdownDescription: "cinder represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
														Attributes: map[string]schema.Attribute{
															"fs_type": schema.StringAttribute{
																Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"read_only": schema.BoolAttribute{
																Description:         "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_ref": schema.SingleNestedAttribute{
																Description:         "secretRef is optional: points to a secret object containing parameters used to connect to OpenStack.",
																MarkdownDescription: "secretRef is optional: points to a secret object containing parameters used to connect to OpenStack.",
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"volume_id": schema.StringAttribute{
																Description:         "volumeID used to identify the volume in cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																MarkdownDescription: "volumeID used to identify the volume in cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"config_map": schema.SingleNestedAttribute{
														Description:         "configMap represents a configMap that should populate this volume",
														MarkdownDescription: "configMap represents a configMap that should populate this volume",
														Attributes: map[string]schema.Attribute{
															"default_mode": schema.Int64Attribute{
																Description:         "defaultMode is optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																MarkdownDescription: "defaultMode is optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"items": schema.ListNestedAttribute{
																Description:         "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"key": schema.StringAttribute{
																			Description:         "key is the key to project.",
																			MarkdownDescription: "key is the key to project.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"mode": schema.Int64Attribute{
																			Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																			MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"path": schema.StringAttribute{
																			Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																			MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
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

															"name": schema.StringAttribute{
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"optional": schema.BoolAttribute{
																Description:         "optional specify whether the ConfigMap or its keys must be defined",
																MarkdownDescription: "optional specify whether the ConfigMap or its keys must be defined",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"csi": schema.SingleNestedAttribute{
														Description:         "csi (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta feature).",
														MarkdownDescription: "csi (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta feature).",
														Attributes: map[string]schema.Attribute{
															"driver": schema.StringAttribute{
																Description:         "driver is the name of the CSI driver that handles this volume. Consult with your admin for the correct name as registered in the cluster.",
																MarkdownDescription: "driver is the name of the CSI driver that handles this volume. Consult with your admin for the correct name as registered in the cluster.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"fs_type": schema.StringAttribute{
																Description:         "fsType to mount. Ex. 'ext4', 'xfs', 'ntfs'. If not provided, the empty value is passed to the associated CSI driver which will determine the default filesystem to apply.",
																MarkdownDescription: "fsType to mount. Ex. 'ext4', 'xfs', 'ntfs'. If not provided, the empty value is passed to the associated CSI driver which will determine the default filesystem to apply.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"node_publish_secret_ref": schema.SingleNestedAttribute{
																Description:         "nodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls. This field is optional, and  may be empty if no secret is required. If the secret object contains more than one secret, all secret references are passed.",
																MarkdownDescription: "nodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls. This field is optional, and  may be empty if no secret is required. If the secret object contains more than one secret, all secret references are passed.",
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"read_only": schema.BoolAttribute{
																Description:         "readOnly specifies a read-only configuration for the volume. Defaults to false (read/write).",
																MarkdownDescription: "readOnly specifies a read-only configuration for the volume. Defaults to false (read/write).",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"volume_attributes": schema.MapAttribute{
																Description:         "volumeAttributes stores driver-specific properties that are passed to the CSI driver. Consult your driver's documentation for supported values.",
																MarkdownDescription: "volumeAttributes stores driver-specific properties that are passed to the CSI driver. Consult your driver's documentation for supported values.",
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

													"downward_api": schema.SingleNestedAttribute{
														Description:         "downwardAPI represents downward API about the pod that should populate this volume",
														MarkdownDescription: "downwardAPI represents downward API about the pod that should populate this volume",
														Attributes: map[string]schema.Attribute{
															"default_mode": schema.Int64Attribute{
																Description:         "Optional: mode bits to use on created files by default. Must be a Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																MarkdownDescription: "Optional: mode bits to use on created files by default. Must be a Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"items": schema.ListNestedAttribute{
																Description:         "Items is a list of downward API volume file",
																MarkdownDescription: "Items is a list of downward API volume file",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"field_ref": schema.SingleNestedAttribute{
																			Description:         "Required: Selects a field of the pod: only annotations, labels, name and namespace are supported.",
																			MarkdownDescription: "Required: Selects a field of the pod: only annotations, labels, name and namespace are supported.",
																			Attributes: map[string]schema.Attribute{
																				"api_version": schema.StringAttribute{
																					Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																					MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"field_path": schema.StringAttribute{
																					Description:         "Path of the field to select in the specified API version.",
																					MarkdownDescription: "Path of the field to select in the specified API version.",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"mode": schema.Int64Attribute{
																			Description:         "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																			MarkdownDescription: "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"path": schema.StringAttribute{
																			Description:         "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
																			MarkdownDescription: "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"resource_field_ref": schema.SingleNestedAttribute{
																			Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
																			MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
																			Attributes: map[string]schema.Attribute{
																				"container_name": schema.StringAttribute{
																					Description:         "Container name: required for volumes, optional for env vars",
																					MarkdownDescription: "Container name: required for volumes, optional for env vars",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"divisor": schema.StringAttribute{
																					Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																					MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"resource": schema.StringAttribute{
																					Description:         "Required: resource to select",
																					MarkdownDescription: "Required: resource to select",
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

													"empty_dir": schema.SingleNestedAttribute{
														Description:         "emptyDir represents a temporary directory that shares a pod's lifetime. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
														MarkdownDescription: "emptyDir represents a temporary directory that shares a pod's lifetime. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
														Attributes: map[string]schema.Attribute{
															"medium": schema.StringAttribute{
																Description:         "medium represents what type of storage medium should back this directory. The default is '' which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
																MarkdownDescription: "medium represents what type of storage medium should back this directory. The default is '' which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"size_limit": schema.StringAttribute{
																Description:         "sizeLimit is the total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info: http://kubernetes.io/docs/user-guide/volumes#emptydir",
																MarkdownDescription: "sizeLimit is the total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info: http://kubernetes.io/docs/user-guide/volumes#emptydir",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"ephemeral": schema.SingleNestedAttribute{
														Description:         "ephemeral represents a volume that is handled by a cluster storage driver. The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts, and deleted when the pod is removed.  Use this if: a) the volume is only needed while the pod runs, b) features of normal volumes like restoring from snapshot or capacity    tracking are needed, c) the storage driver is specified through a storage class, and d) the storage driver supports dynamic volume provisioning through    a PersistentVolumeClaim (see EphemeralVolumeSource for more    information on the connection between this volume type    and PersistentVolumeClaim).  Use PersistentVolumeClaim or one of the vendor-specific APIs for volumes that persist for longer than the lifecycle of an individual pod.  Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to be used that way - see the documentation of the driver for more information.  A pod can use both types of ephemeral volumes and persistent volumes at the same time.",
														MarkdownDescription: "ephemeral represents a volume that is handled by a cluster storage driver. The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts, and deleted when the pod is removed.  Use this if: a) the volume is only needed while the pod runs, b) features of normal volumes like restoring from snapshot or capacity    tracking are needed, c) the storage driver is specified through a storage class, and d) the storage driver supports dynamic volume provisioning through    a PersistentVolumeClaim (see EphemeralVolumeSource for more    information on the connection between this volume type    and PersistentVolumeClaim).  Use PersistentVolumeClaim or one of the vendor-specific APIs for volumes that persist for longer than the lifecycle of an individual pod.  Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to be used that way - see the documentation of the driver for more information.  A pod can use both types of ephemeral volumes and persistent volumes at the same time.",
														Attributes: map[string]schema.Attribute{
															"volume_claim_template": schema.SingleNestedAttribute{
																Description:         "Will be used to create a stand-alone PVC to provision the volume. The pod in which this EphemeralVolumeSource is embedded will be the owner of the PVC, i.e. the PVC will be deleted together with the pod.  The name of the PVC will be '<pod name>-<volume name>' where '<volume name>' is the name from the 'PodSpec.Volumes' array entry. Pod validation will reject the pod if the concatenated name is not valid for a PVC (for example, too long).  An existing PVC with that name that is not owned by the pod will *not* be used for the pod to avoid using an unrelated volume by mistake. Starting the pod is then blocked until the unrelated PVC is removed. If such a pre-created PVC is meant to be used by the pod, the PVC has to updated with an owner reference to the pod once the pod exists. Normally this should not be necessary, but it may be useful when manually reconstructing a broken cluster.  This field is read-only and no changes will be made by Kubernetes to the PVC after it has been created.  Required, must not be nil.",
																MarkdownDescription: "Will be used to create a stand-alone PVC to provision the volume. The pod in which this EphemeralVolumeSource is embedded will be the owner of the PVC, i.e. the PVC will be deleted together with the pod.  The name of the PVC will be '<pod name>-<volume name>' where '<volume name>' is the name from the 'PodSpec.Volumes' array entry. Pod validation will reject the pod if the concatenated name is not valid for a PVC (for example, too long).  An existing PVC with that name that is not owned by the pod will *not* be used for the pod to avoid using an unrelated volume by mistake. Starting the pod is then blocked until the unrelated PVC is removed. If such a pre-created PVC is meant to be used by the pod, the PVC has to updated with an owner reference to the pod once the pod exists. Normally this should not be necessary, but it may be useful when manually reconstructing a broken cluster.  This field is read-only and no changes will be made by Kubernetes to the PVC after it has been created.  Required, must not be nil.",
																Attributes: map[string]schema.Attribute{
																	"metadata": schema.MapAttribute{
																		Description:         "May contain labels and annotations that will be copied into the PVC when creating it. No other fields are allowed and will be rejected during validation.",
																		MarkdownDescription: "May contain labels and annotations that will be copied into the PVC when creating it. No other fields are allowed and will be rejected during validation.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"spec": schema.SingleNestedAttribute{
																		Description:         "The specification for the PersistentVolumeClaim. The entire content is copied unchanged into the PVC that gets created from this template. The same fields as in a PersistentVolumeClaim are also valid here.",
																		MarkdownDescription: "The specification for the PersistentVolumeClaim. The entire content is copied unchanged into the PVC that gets created from this template. The same fields as in a PersistentVolumeClaim are also valid here.",
																		Attributes: map[string]schema.Attribute{
																			"access_modes": schema.ListAttribute{
																				Description:         "accessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
																				MarkdownDescription: "accessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"data_source": schema.SingleNestedAttribute{
																				Description:         "dataSource field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef, and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified. If the namespace is specified, then dataSourceRef will not be copied to dataSource.",
																				MarkdownDescription: "dataSource field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef, and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified. If the namespace is specified, then dataSourceRef will not be copied to dataSource.",
																				Attributes: map[string]schema.Attribute{
																					"api_group": schema.StringAttribute{
																						Description:         "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
																						MarkdownDescription: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"kind": schema.StringAttribute{
																						Description:         "Kind is the type of resource being referenced",
																						MarkdownDescription: "Kind is the type of resource being referenced",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"name": schema.StringAttribute{
																						Description:         "Name is the name of resource being referenced",
																						MarkdownDescription: "Name is the name of resource being referenced",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},
																				},
																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"data_source_ref": schema.SingleNestedAttribute{
																				Description:         "dataSourceRef specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the dataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, when namespace isn't specified in dataSourceRef, both fields (dataSource and dataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. When namespace is specified in dataSourceRef, dataSource isn't set to the same value and must be empty. There are three important differences between dataSource and dataSourceRef: * While dataSource only allows two specific types of objects, dataSourceRef   allows any non-core object, as well as PersistentVolumeClaim objects. * While dataSource ignores disallowed values (dropping them), dataSourceRef   preserves all values, and generates an error if a disallowed value is   specified. * While dataSource only allows local objects, dataSourceRef allows objects   in any namespaces. (Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled. (Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
																				MarkdownDescription: "dataSourceRef specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the dataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, when namespace isn't specified in dataSourceRef, both fields (dataSource and dataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. When namespace is specified in dataSourceRef, dataSource isn't set to the same value and must be empty. There are three important differences between dataSource and dataSourceRef: * While dataSource only allows two specific types of objects, dataSourceRef   allows any non-core object, as well as PersistentVolumeClaim objects. * While dataSource ignores disallowed values (dropping them), dataSourceRef   preserves all values, and generates an error if a disallowed value is   specified. * While dataSource only allows local objects, dataSourceRef allows objects   in any namespaces. (Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled. (Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
																				Attributes: map[string]schema.Attribute{
																					"api_group": schema.StringAttribute{
																						Description:         "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
																						MarkdownDescription: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"kind": schema.StringAttribute{
																						Description:         "Kind is the type of resource being referenced",
																						MarkdownDescription: "Kind is the type of resource being referenced",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"name": schema.StringAttribute{
																						Description:         "Name is the name of resource being referenced",
																						MarkdownDescription: "Name is the name of resource being referenced",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"namespace": schema.StringAttribute{
																						Description:         "Namespace is the namespace of resource being referenced Note that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details. (Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
																						MarkdownDescription: "Namespace is the namespace of resource being referenced Note that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details. (Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},
																				},
																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"resources": schema.SingleNestedAttribute{
																				Description:         "resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
																				MarkdownDescription: "resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
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
																						Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																						MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

																			"selector": schema.SingleNestedAttribute{
																				Description:         "selector is a label query over volumes to consider for binding.",
																				MarkdownDescription: "selector is a label query over volumes to consider for binding.",
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

																			"storage_class_name": schema.StringAttribute{
																				Description:         "storageClassName is the name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
																				MarkdownDescription: "storageClassName is the name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"volume_mode": schema.StringAttribute{
																				Description:         "volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.",
																				MarkdownDescription: "volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"volume_name": schema.StringAttribute{
																				Description:         "volumeName is the binding reference to the PersistentVolume backing this claim.",
																				MarkdownDescription: "volumeName is the binding reference to the PersistentVolume backing this claim.",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"fc": schema.SingleNestedAttribute{
														Description:         "fc represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.",
														MarkdownDescription: "fc represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.",
														Attributes: map[string]schema.Attribute{
															"fs_type": schema.StringAttribute{
																Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. TODO: how do we prevent errors in the filesystem from compromising the machine",
																MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. TODO: how do we prevent errors in the filesystem from compromising the machine",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"lun": schema.Int64Attribute{
																Description:         "lun is Optional: FC target lun number",
																MarkdownDescription: "lun is Optional: FC target lun number",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"read_only": schema.BoolAttribute{
																Description:         "readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																MarkdownDescription: "readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"target_ww_ns": schema.ListAttribute{
																Description:         "targetWWNs is Optional: FC target worldwide names (WWNs)",
																MarkdownDescription: "targetWWNs is Optional: FC target worldwide names (WWNs)",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"wwids": schema.ListAttribute{
																Description:         "wwids Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.",
																MarkdownDescription: "wwids Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.",
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

													"flex_volume": schema.SingleNestedAttribute{
														Description:         "flexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.",
														MarkdownDescription: "flexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.",
														Attributes: map[string]schema.Attribute{
															"driver": schema.StringAttribute{
																Description:         "driver is the name of the driver to use for this volume.",
																MarkdownDescription: "driver is the name of the driver to use for this volume.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"fs_type": schema.StringAttribute{
																Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. The default filesystem depends on FlexVolume script.",
																MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. The default filesystem depends on FlexVolume script.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"options": schema.MapAttribute{
																Description:         "options is Optional: this field holds extra command options if any.",
																MarkdownDescription: "options is Optional: this field holds extra command options if any.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"read_only": schema.BoolAttribute{
																Description:         "readOnly is Optional: defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																MarkdownDescription: "readOnly is Optional: defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_ref": schema.SingleNestedAttribute{
																Description:         "secretRef is Optional: secretRef is reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts.",
																MarkdownDescription: "secretRef is Optional: secretRef is reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts.",
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

													"flocker": schema.SingleNestedAttribute{
														Description:         "flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running",
														MarkdownDescription: "flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running",
														Attributes: map[string]schema.Attribute{
															"dataset_name": schema.StringAttribute{
																Description:         "datasetName is Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated",
																MarkdownDescription: "datasetName is Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"dataset_uuid": schema.StringAttribute{
																Description:         "datasetUUID is the UUID of the dataset. This is unique identifier of a Flocker dataset",
																MarkdownDescription: "datasetUUID is the UUID of the dataset. This is unique identifier of a Flocker dataset",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"gce_persistent_disk": schema.SingleNestedAttribute{
														Description:         "gcePersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
														MarkdownDescription: "gcePersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
														Attributes: map[string]schema.Attribute{
															"fs_type": schema.StringAttribute{
																Description:         "fsType is filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk TODO: how do we prevent errors in the filesystem from compromising the machine",
																MarkdownDescription: "fsType is filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk TODO: how do we prevent errors in the filesystem from compromising the machine",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"partition": schema.Int64Attribute{
																Description:         "partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																MarkdownDescription: "partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"pd_name": schema.StringAttribute{
																Description:         "pdName is unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																MarkdownDescription: "pdName is unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"read_only": schema.BoolAttribute{
																Description:         "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																MarkdownDescription: "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"git_repo": schema.SingleNestedAttribute{
														Description:         "gitRepo represents a git repository at a particular revision. DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir into the Pod's container.",
														MarkdownDescription: "gitRepo represents a git repository at a particular revision. DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir into the Pod's container.",
														Attributes: map[string]schema.Attribute{
															"directory": schema.StringAttribute{
																Description:         "directory is the target directory name. Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the git repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.",
																MarkdownDescription: "directory is the target directory name. Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the git repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"repository": schema.StringAttribute{
																Description:         "repository is the URL",
																MarkdownDescription: "repository is the URL",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"revision": schema.StringAttribute{
																Description:         "revision is the commit hash for the specified revision.",
																MarkdownDescription: "revision is the commit hash for the specified revision.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"glusterfs": schema.SingleNestedAttribute{
														Description:         "glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/glusterfs/README.md",
														MarkdownDescription: "glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/glusterfs/README.md",
														Attributes: map[string]schema.Attribute{
															"endpoints": schema.StringAttribute{
																Description:         "endpoints is the endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
																MarkdownDescription: "endpoints is the endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"path": schema.StringAttribute{
																Description:         "path is the Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
																MarkdownDescription: "path is the Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"read_only": schema.BoolAttribute{
																Description:         "readOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
																MarkdownDescription: "readOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"host_path": schema.SingleNestedAttribute{
														Description:         "hostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath --- TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can not mount host directories as read/write.",
														MarkdownDescription: "hostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath --- TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can not mount host directories as read/write.",
														Attributes: map[string]schema.Attribute{
															"path": schema.StringAttribute{
																Description:         "path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
																MarkdownDescription: "path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"type": schema.StringAttribute{
																Description:         "type for HostPath Volume Defaults to '' More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
																MarkdownDescription: "type for HostPath Volume Defaults to '' More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"iscsi": schema.SingleNestedAttribute{
														Description:         "iscsi represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://examples.k8s.io/volumes/iscsi/README.md",
														MarkdownDescription: "iscsi represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://examples.k8s.io/volumes/iscsi/README.md",
														Attributes: map[string]schema.Attribute{
															"chap_auth_discovery": schema.BoolAttribute{
																Description:         "chapAuthDiscovery defines whether support iSCSI Discovery CHAP authentication",
																MarkdownDescription: "chapAuthDiscovery defines whether support iSCSI Discovery CHAP authentication",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"chap_auth_session": schema.BoolAttribute{
																Description:         "chapAuthSession defines whether support iSCSI Session CHAP authentication",
																MarkdownDescription: "chapAuthSession defines whether support iSCSI Session CHAP authentication",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"fs_type": schema.StringAttribute{
																Description:         "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi TODO: how do we prevent errors in the filesystem from compromising the machine",
																MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi TODO: how do we prevent errors in the filesystem from compromising the machine",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"initiator_name": schema.StringAttribute{
																Description:         "initiatorName is the custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.",
																MarkdownDescription: "initiatorName is the custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"iqn": schema.StringAttribute{
																Description:         "iqn is the target iSCSI Qualified Name.",
																MarkdownDescription: "iqn is the target iSCSI Qualified Name.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"iscsi_interface": schema.StringAttribute{
																Description:         "iscsiInterface is the interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).",
																MarkdownDescription: "iscsiInterface is the interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"lun": schema.Int64Attribute{
																Description:         "lun represents iSCSI Target Lun number.",
																MarkdownDescription: "lun represents iSCSI Target Lun number.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"portals": schema.ListAttribute{
																Description:         "portals is the iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
																MarkdownDescription: "portals is the iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"read_only": schema.BoolAttribute{
																Description:         "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.",
																MarkdownDescription: "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_ref": schema.SingleNestedAttribute{
																Description:         "secretRef is the CHAP Secret for iSCSI target and initiator authentication",
																MarkdownDescription: "secretRef is the CHAP Secret for iSCSI target and initiator authentication",
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"target_portal": schema.StringAttribute{
																Description:         "targetPortal is iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
																MarkdownDescription: "targetPortal is iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": schema.StringAttribute{
														Description:         "name of the volume. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														MarkdownDescription: "name of the volume. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"nfs": schema.SingleNestedAttribute{
														Description:         "nfs represents an NFS mount on the host that shares a pod's lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
														MarkdownDescription: "nfs represents an NFS mount on the host that shares a pod's lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
														Attributes: map[string]schema.Attribute{
															"path": schema.StringAttribute{
																Description:         "path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																MarkdownDescription: "path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"read_only": schema.BoolAttribute{
																Description:         "readOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																MarkdownDescription: "readOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"server": schema.StringAttribute{
																Description:         "server is the hostname or IP address of the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																MarkdownDescription: "server is the hostname or IP address of the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"persistent_volume_claim": schema.SingleNestedAttribute{
														Description:         "persistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
														MarkdownDescription: "persistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
														Attributes: map[string]schema.Attribute{
															"claim_name": schema.StringAttribute{
																Description:         "claimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
																MarkdownDescription: "claimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"read_only": schema.BoolAttribute{
																Description:         "readOnly Will force the ReadOnly setting in VolumeMounts. Default false.",
																MarkdownDescription: "readOnly Will force the ReadOnly setting in VolumeMounts. Default false.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"photon_persistent_disk": schema.SingleNestedAttribute{
														Description:         "photonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine",
														MarkdownDescription: "photonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine",
														Attributes: map[string]schema.Attribute{
															"fs_type": schema.StringAttribute{
																Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"pd_id": schema.StringAttribute{
																Description:         "pdID is the ID that identifies Photon Controller persistent disk",
																MarkdownDescription: "pdID is the ID that identifies Photon Controller persistent disk",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"portworx_volume": schema.SingleNestedAttribute{
														Description:         "portworxVolume represents a portworx volume attached and mounted on kubelets host machine",
														MarkdownDescription: "portworxVolume represents a portworx volume attached and mounted on kubelets host machine",
														Attributes: map[string]schema.Attribute{
															"fs_type": schema.StringAttribute{
																Description:         "fSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs'. Implicitly inferred to be 'ext4' if unspecified.",
																MarkdownDescription: "fSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs'. Implicitly inferred to be 'ext4' if unspecified.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"read_only": schema.BoolAttribute{
																Description:         "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"volume_id": schema.StringAttribute{
																Description:         "volumeID uniquely identifies a Portworx volume",
																MarkdownDescription: "volumeID uniquely identifies a Portworx volume",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"projected": schema.SingleNestedAttribute{
														Description:         "projected items for all in one resources secrets, configmaps, and downward API",
														MarkdownDescription: "projected items for all in one resources secrets, configmaps, and downward API",
														Attributes: map[string]schema.Attribute{
															"default_mode": schema.Int64Attribute{
																Description:         "defaultMode are the mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																MarkdownDescription: "defaultMode are the mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"sources": schema.ListNestedAttribute{
																Description:         "sources is the list of volume projections",
																MarkdownDescription: "sources is the list of volume projections",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"config_map": schema.SingleNestedAttribute{
																			Description:         "configMap information about the configMap data to project",
																			MarkdownDescription: "configMap information about the configMap data to project",
																			Attributes: map[string]schema.Attribute{
																				"items": schema.ListNestedAttribute{
																					Description:         "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																					MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"key": schema.StringAttribute{
																								Description:         "key is the key to project.",
																								MarkdownDescription: "key is the key to project.",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"mode": schema.Int64Attribute{
																								Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																								MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"path": schema.StringAttribute{
																								Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																								MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
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

																				"name": schema.StringAttribute{
																					Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																					MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"optional": schema.BoolAttribute{
																					Description:         "optional specify whether the ConfigMap or its keys must be defined",
																					MarkdownDescription: "optional specify whether the ConfigMap or its keys must be defined",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"downward_api": schema.SingleNestedAttribute{
																			Description:         "downwardAPI information about the downwardAPI data to project",
																			MarkdownDescription: "downwardAPI information about the downwardAPI data to project",
																			Attributes: map[string]schema.Attribute{
																				"items": schema.ListNestedAttribute{
																					Description:         "Items is a list of DownwardAPIVolume file",
																					MarkdownDescription: "Items is a list of DownwardAPIVolume file",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"field_ref": schema.SingleNestedAttribute{
																								Description:         "Required: Selects a field of the pod: only annotations, labels, name and namespace are supported.",
																								MarkdownDescription: "Required: Selects a field of the pod: only annotations, labels, name and namespace are supported.",
																								Attributes: map[string]schema.Attribute{
																									"api_version": schema.StringAttribute{
																										Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																										MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},

																									"field_path": schema.StringAttribute{
																										Description:         "Path of the field to select in the specified API version.",
																										MarkdownDescription: "Path of the field to select in the specified API version.",
																										Required:            true,
																										Optional:            false,
																										Computed:            false,
																									},
																								},
																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"mode": schema.Int64Attribute{
																								Description:         "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																								MarkdownDescription: "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"path": schema.StringAttribute{
																								Description:         "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
																								MarkdownDescription: "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"resource_field_ref": schema.SingleNestedAttribute{
																								Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
																								MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
																								Attributes: map[string]schema.Attribute{
																									"container_name": schema.StringAttribute{
																										Description:         "Container name: required for volumes, optional for env vars",
																										MarkdownDescription: "Container name: required for volumes, optional for env vars",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},

																									"divisor": schema.StringAttribute{
																										Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																										MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},

																									"resource": schema.StringAttribute{
																										Description:         "Required: resource to select",
																										MarkdownDescription: "Required: resource to select",
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

																		"secret": schema.SingleNestedAttribute{
																			Description:         "secret information about the secret data to project",
																			MarkdownDescription: "secret information about the secret data to project",
																			Attributes: map[string]schema.Attribute{
																				"items": schema.ListNestedAttribute{
																					Description:         "items if unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																					MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"key": schema.StringAttribute{
																								Description:         "key is the key to project.",
																								MarkdownDescription: "key is the key to project.",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"mode": schema.Int64Attribute{
																								Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																								MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"path": schema.StringAttribute{
																								Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																								MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
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

																				"name": schema.StringAttribute{
																					Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																					MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"optional": schema.BoolAttribute{
																					Description:         "optional field specify whether the Secret or its key must be defined",
																					MarkdownDescription: "optional field specify whether the Secret or its key must be defined",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"service_account_token": schema.SingleNestedAttribute{
																			Description:         "serviceAccountToken is information about the serviceAccountToken data to project",
																			MarkdownDescription: "serviceAccountToken is information about the serviceAccountToken data to project",
																			Attributes: map[string]schema.Attribute{
																				"audience": schema.StringAttribute{
																					Description:         "audience is the intended audience of the token. A recipient of a token must identify itself with an identifier specified in the audience of the token, and otherwise should reject the token. The audience defaults to the identifier of the apiserver.",
																					MarkdownDescription: "audience is the intended audience of the token. A recipient of a token must identify itself with an identifier specified in the audience of the token, and otherwise should reject the token. The audience defaults to the identifier of the apiserver.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"expiration_seconds": schema.Int64Attribute{
																					Description:         "expirationSeconds is the requested duration of validity of the service account token. As the token approaches expiration, the kubelet volume plugin will proactively rotate the service account token. The kubelet will start trying to rotate the token if the token is older than 80 percent of its time to live or if the token is older than 24 hours.Defaults to 1 hour and must be at least 10 minutes.",
																					MarkdownDescription: "expirationSeconds is the requested duration of validity of the service account token. As the token approaches expiration, the kubelet volume plugin will proactively rotate the service account token. The kubelet will start trying to rotate the token if the token is older than 80 percent of its time to live or if the token is older than 24 hours.Defaults to 1 hour and must be at least 10 minutes.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"path": schema.StringAttribute{
																					Description:         "path is the path relative to the mount point of the file to project the token into.",
																					MarkdownDescription: "path is the path relative to the mount point of the file to project the token into.",
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

													"quobyte": schema.SingleNestedAttribute{
														Description:         "quobyte represents a Quobyte mount on the host that shares a pod's lifetime",
														MarkdownDescription: "quobyte represents a Quobyte mount on the host that shares a pod's lifetime",
														Attributes: map[string]schema.Attribute{
															"group": schema.StringAttribute{
																Description:         "group to map volume access to Default is no group",
																MarkdownDescription: "group to map volume access to Default is no group",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"read_only": schema.BoolAttribute{
																Description:         "readOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false.",
																MarkdownDescription: "readOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"registry": schema.StringAttribute{
																Description:         "registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes",
																MarkdownDescription: "registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"tenant": schema.StringAttribute{
																Description:         "tenant owning the given Quobyte volume in the Backend Used with dynamically provisioned Quobyte volumes, value is set by the plugin",
																MarkdownDescription: "tenant owning the given Quobyte volume in the Backend Used with dynamically provisioned Quobyte volumes, value is set by the plugin",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"user": schema.StringAttribute{
																Description:         "user to map volume access to Defaults to serivceaccount user",
																MarkdownDescription: "user to map volume access to Defaults to serivceaccount user",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"volume": schema.StringAttribute{
																Description:         "volume is a string that references an already created Quobyte volume by name.",
																MarkdownDescription: "volume is a string that references an already created Quobyte volume by name.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"rbd": schema.SingleNestedAttribute{
														Description:         "rbd represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md",
														MarkdownDescription: "rbd represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md",
														Attributes: map[string]schema.Attribute{
															"fs_type": schema.StringAttribute{
																Description:         "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd TODO: how do we prevent errors in the filesystem from compromising the machine",
																MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd TODO: how do we prevent errors in the filesystem from compromising the machine",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"image": schema.StringAttribute{
																Description:         "image is the rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																MarkdownDescription: "image is the rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"keyring": schema.StringAttribute{
																Description:         "keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																MarkdownDescription: "keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"monitors": schema.ListAttribute{
																Description:         "monitors is a collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																MarkdownDescription: "monitors is a collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																ElementType:         types.StringType,
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"pool": schema.StringAttribute{
																Description:         "pool is the rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																MarkdownDescription: "pool is the rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"read_only": schema.BoolAttribute{
																Description:         "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																MarkdownDescription: "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_ref": schema.SingleNestedAttribute{
																Description:         "secretRef is name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																MarkdownDescription: "secretRef is name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"user": schema.StringAttribute{
																Description:         "user is the rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																MarkdownDescription: "user is the rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"scale_io": schema.SingleNestedAttribute{
														Description:         "scaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.",
														MarkdownDescription: "scaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.",
														Attributes: map[string]schema.Attribute{
															"fs_type": schema.StringAttribute{
																Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Default is 'xfs'.",
																MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Default is 'xfs'.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"gateway": schema.StringAttribute{
																Description:         "gateway is the host address of the ScaleIO API Gateway.",
																MarkdownDescription: "gateway is the host address of the ScaleIO API Gateway.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"protection_domain": schema.StringAttribute{
																Description:         "protectionDomain is the name of the ScaleIO Protection Domain for the configured storage.",
																MarkdownDescription: "protectionDomain is the name of the ScaleIO Protection Domain for the configured storage.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"read_only": schema.BoolAttribute{
																Description:         "readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																MarkdownDescription: "readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_ref": schema.SingleNestedAttribute{
																Description:         "secretRef references to the secret for ScaleIO user and other sensitive information. If this is not provided, Login operation will fail.",
																MarkdownDescription: "secretRef references to the secret for ScaleIO user and other sensitive information. If this is not provided, Login operation will fail.",
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: true,
																Optional: false,
																Computed: false,
															},

															"ssl_enabled": schema.BoolAttribute{
																Description:         "sslEnabled Flag enable/disable SSL communication with Gateway, default false",
																MarkdownDescription: "sslEnabled Flag enable/disable SSL communication with Gateway, default false",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"storage_mode": schema.StringAttribute{
																Description:         "storageMode indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned.",
																MarkdownDescription: "storageMode indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"storage_pool": schema.StringAttribute{
																Description:         "storagePool is the ScaleIO Storage Pool associated with the protection domain.",
																MarkdownDescription: "storagePool is the ScaleIO Storage Pool associated with the protection domain.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"system": schema.StringAttribute{
																Description:         "system is the name of the storage system as configured in ScaleIO.",
																MarkdownDescription: "system is the name of the storage system as configured in ScaleIO.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"volume_name": schema.StringAttribute{
																Description:         "volumeName is the name of a volume already created in the ScaleIO system that is associated with this volume source.",
																MarkdownDescription: "volumeName is the name of a volume already created in the ScaleIO system that is associated with this volume source.",
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
														Description:         "secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
														MarkdownDescription: "secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
														Attributes: map[string]schema.Attribute{
															"default_mode": schema.Int64Attribute{
																Description:         "defaultMode is Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																MarkdownDescription: "defaultMode is Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"items": schema.ListNestedAttribute{
																Description:         "items If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																MarkdownDescription: "items If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"key": schema.StringAttribute{
																			Description:         "key is the key to project.",
																			MarkdownDescription: "key is the key to project.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"mode": schema.Int64Attribute{
																			Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																			MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"path": schema.StringAttribute{
																			Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																			MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
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

															"optional": schema.BoolAttribute{
																Description:         "optional field specify whether the Secret or its keys must be defined",
																MarkdownDescription: "optional field specify whether the Secret or its keys must be defined",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_name": schema.StringAttribute{
																Description:         "secretName is the name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
																MarkdownDescription: "secretName is the name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"storageos": schema.SingleNestedAttribute{
														Description:         "storageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.",
														MarkdownDescription: "storageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.",
														Attributes: map[string]schema.Attribute{
															"fs_type": schema.StringAttribute{
																Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"read_only": schema.BoolAttribute{
																Description:         "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_ref": schema.SingleNestedAttribute{
																Description:         "secretRef specifies the secret to use for obtaining the StorageOS API credentials.  If not specified, default values will be attempted.",
																MarkdownDescription: "secretRef specifies the secret to use for obtaining the StorageOS API credentials.  If not specified, default values will be attempted.",
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"volume_name": schema.StringAttribute{
																Description:         "volumeName is the human-readable name of the StorageOS volume.  Volume names are only unique within a namespace.",
																MarkdownDescription: "volumeName is the human-readable name of the StorageOS volume.  Volume names are only unique within a namespace.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"volume_namespace": schema.StringAttribute{
																Description:         "volumeNamespace specifies the scope of the volume within StorageOS.  If no namespace is specified then the Pod's namespace will be used.  This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to 'default' if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.",
																MarkdownDescription: "volumeNamespace specifies the scope of the volume within StorageOS.  If no namespace is specified then the Pod's namespace will be used.  This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to 'default' if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"vsphere_volume": schema.SingleNestedAttribute{
														Description:         "vsphereVolume represents a vSphere volume attached and mounted on kubelets host machine",
														MarkdownDescription: "vsphereVolume represents a vSphere volume attached and mounted on kubelets host machine",
														Attributes: map[string]schema.Attribute{
															"fs_type": schema.StringAttribute{
																Description:         "fsType is filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																MarkdownDescription: "fsType is filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"storage_policy_id": schema.StringAttribute{
																Description:         "storagePolicyID is the storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.",
																MarkdownDescription: "storagePolicyID is the storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"storage_policy_name": schema.StringAttribute{
																Description:         "storagePolicyName is the storage Policy Based Management (SPBM) profile name.",
																MarkdownDescription: "storagePolicyName is the storage Policy Based Management (SPBM) profile name.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"volume_path": schema.StringAttribute{
																Description:         "volumePath is the path that identifies vSphere volume vmdk",
																MarkdownDescription: "volumePath is the path that identifies vSphere volume vmdk",
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

								"template_type": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"time_chaos": schema.SingleNestedAttribute{
									Description:         "TimeChaosSpec defines the desired state of TimeChaos",
									MarkdownDescription: "TimeChaosSpec defines the desired state of TimeChaos",
									Attributes: map[string]schema.Attribute{
										"clock_ids": schema.ListAttribute{
											Description:         "ClockIds defines all affected clock id All available options are ['CLOCK_REALTIME','CLOCK_MONOTONIC','CLOCK_PROCESS_CPUTIME_ID','CLOCK_THREAD_CPUTIME_ID', 'CLOCK_MONOTONIC_RAW','CLOCK_REALTIME_COARSE','CLOCK_MONOTONIC_COARSE','CLOCK_BOOTTIME','CLOCK_REALTIME_ALARM', 'CLOCK_BOOTTIME_ALARM'] Default value is ['CLOCK_REALTIME']",
											MarkdownDescription: "ClockIds defines all affected clock id All available options are ['CLOCK_REALTIME','CLOCK_MONOTONIC','CLOCK_PROCESS_CPUTIME_ID','CLOCK_THREAD_CPUTIME_ID', 'CLOCK_MONOTONIC_RAW','CLOCK_REALTIME_COARSE','CLOCK_MONOTONIC_COARSE','CLOCK_BOOTTIME','CLOCK_REALTIME_ALARM', 'CLOCK_BOOTTIME_ALARM'] Default value is ['CLOCK_REALTIME']",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"container_names": schema.ListAttribute{
											Description:         "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
											MarkdownDescription: "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"duration": schema.StringAttribute{
											Description:         "Duration represents the duration of the chaos action",
											MarkdownDescription: "Duration represents the duration of the chaos action",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"mode": schema.StringAttribute{
											Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
											MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
											},
										},

										"remote_cluster": schema.StringAttribute{
											Description:         "RemoteCluster represents the remote cluster where the chaos will be deployed",
											MarkdownDescription: "RemoteCluster represents the remote cluster where the chaos will be deployed",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"selector": schema.SingleNestedAttribute{
											Description:         "Selector is used to select pods that are used to inject chaos action.",
											MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",
											Attributes: map[string]schema.Attribute{
												"annotation_selectors": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
													MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"expression_selectors": schema.ListNestedAttribute{
													Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
													MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
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

												"field_selectors": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
													MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"label_selectors": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
													MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespaces": schema.ListAttribute{
													Description:         "Namespaces is a set of namespace to which objects belong.",
													MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"node_selectors": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
													MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"nodes": schema.ListAttribute{
													Description:         "Nodes is a set of node name and objects must belong to these nodes.",
													MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"pod_phase_selectors": schema.ListAttribute{
													Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
													MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"pods": schema.MapAttribute{
													Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
													MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
													ElementType:         types.ListType{ElemType: types.StringType},
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"time_offset": schema.StringAttribute{
											Description:         "TimeOffset defines the delta time of injected program. It's a possibly signed sequence of decimal numbers, such as '300ms', '-1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
											MarkdownDescription: "TimeOffset defines the delta time of injected program. It's a possibly signed sequence of decimal numbers, such as '300ms', '-1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
											MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
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
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *ChaosMeshOrgWorkflowV1Alpha1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *dynamic.DynamicClient, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *ChaosMeshOrgWorkflowV1Alpha1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_chaos_mesh_org_workflow_v1alpha1")

	var model ChaosMeshOrgWorkflowV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("chaos-mesh.org/v1alpha1")
	model.Kind = pointer.String("Workflow")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.Resource(k8sSchema.GroupVersionResource{Group: "chaos-mesh.org", Version: "v1alpha1", Resource: "Workflow"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while creating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse ChaosMeshOrgWorkflowV1Alpha1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *ChaosMeshOrgWorkflowV1Alpha1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_chaos_mesh_org_workflow_v1alpha1")

	var data ChaosMeshOrgWorkflowV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "chaos-mesh.org", Version: "v1alpha1", Resource: "Workflow"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse ChaosMeshOrgWorkflowV1Alpha1ResourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *ChaosMeshOrgWorkflowV1Alpha1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_chaos_mesh_org_workflow_v1alpha1")

	var model ChaosMeshOrgWorkflowV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("chaos-mesh.org/v1alpha1")
	model.Kind = pointer.String("Workflow")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.Resource(k8sSchema.GroupVersionResource{Group: "chaos-mesh.org", Version: "v1alpha1", Resource: "Workflow"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while updating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse ChaosMeshOrgWorkflowV1Alpha1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *ChaosMeshOrgWorkflowV1Alpha1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_chaos_mesh_org_workflow_v1alpha1")

	var data ChaosMeshOrgWorkflowV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "chaos-mesh.org", Version: "v1alpha1", Resource: "Workflow"}).
		Namespace(data.Metadata.Namespace).
		Delete(ctx, data.Metadata.Name, meta.DeleteOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to DELETE resource",
			"An unexpected error occurred while deleting the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"DELETE Error: "+err.Error(),
		)
		return
	}
}

func (r *ChaosMeshOrgWorkflowV1Alpha1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	idParts := strings.Split(request.ID, "/")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'namespace/name' Got: '%q'", request.ID),
		)
		return
	}

	namespace := idParts[0]
	name := idParts[1]
	tflog.Trace(ctx, "parsed import ID", map[string]interface{}{
		"namespace": namespace,
		"name":      name,
	})
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("namespace"), namespace)...)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("name"), name)...)
}
