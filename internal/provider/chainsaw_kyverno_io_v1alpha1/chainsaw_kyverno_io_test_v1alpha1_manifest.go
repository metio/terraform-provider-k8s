/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package chainsaw_kyverno_io_v1alpha1

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
	_ datasource.DataSource = &ChainsawKyvernoIoTestV1Alpha1Manifest{}
)

func NewChainsawKyvernoIoTestV1Alpha1Manifest() datasource.DataSource {
	return &ChainsawKyvernoIoTestV1Alpha1Manifest{}
}

type ChainsawKyvernoIoTestV1Alpha1Manifest struct{}

type ChainsawKyvernoIoTestV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Bindings *[]struct {
			Name  *string            `tfsdk:"name" json:"name,omitempty"`
			Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"bindings" json:"bindings,omitempty"`
		Catch *[]struct {
			Command *struct {
				Args     *[]string `tfsdk:"args" json:"args,omitempty"`
				Bindings *[]struct {
					Name  *string            `tfsdk:"name" json:"name,omitempty"`
					Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"bindings" json:"bindings,omitempty"`
				Check    *map[string]string `tfsdk:"check" json:"check,omitempty"`
				Cluster  *string            `tfsdk:"cluster" json:"cluster,omitempty"`
				Clusters *struct {
					Context    *string `tfsdk:"context" json:"context,omitempty"`
					Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
				} `tfsdk:"clusters" json:"clusters,omitempty"`
				Entrypoint *string `tfsdk:"entrypoint" json:"entrypoint,omitempty"`
				Env        *[]struct {
					Name  *string            `tfsdk:"name" json:"name,omitempty"`
					Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"env" json:"env,omitempty"`
				Outputs *[]struct {
					Match *map[string]string `tfsdk:"match" json:"match,omitempty"`
					Name  *string            `tfsdk:"name" json:"name,omitempty"`
					Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"outputs" json:"outputs,omitempty"`
				SkipLogOutput *bool   `tfsdk:"skip_log_output" json:"skipLogOutput,omitempty"`
				Timeout       *string `tfsdk:"timeout" json:"timeout,omitempty"`
			} `tfsdk:"command" json:"command,omitempty"`
			Delete *struct {
				Bindings *[]struct {
					Name  *string            `tfsdk:"name" json:"name,omitempty"`
					Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"bindings" json:"bindings,omitempty"`
				Cluster  *string `tfsdk:"cluster" json:"cluster,omitempty"`
				Clusters *struct {
					Context    *string `tfsdk:"context" json:"context,omitempty"`
					Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
				} `tfsdk:"clusters" json:"clusters,omitempty"`
				DeletionPropagationPolicy *string `tfsdk:"deletion_propagation_policy" json:"deletionPropagationPolicy,omitempty"`
				Expect                    *[]struct {
					Check *map[string]string `tfsdk:"check" json:"check,omitempty"`
					Match *map[string]string `tfsdk:"match" json:"match,omitempty"`
				} `tfsdk:"expect" json:"expect,omitempty"`
				File *string `tfsdk:"file" json:"file,omitempty"`
				Ref  *struct {
					ApiVersion *string            `tfsdk:"api_version" json:"apiVersion,omitempty"`
					Kind       *string            `tfsdk:"kind" json:"kind,omitempty"`
					Labels     *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					Name       *string            `tfsdk:"name" json:"name,omitempty"`
					Namespace  *string            `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"ref" json:"ref,omitempty"`
				Template *bool   `tfsdk:"template" json:"template,omitempty"`
				Timeout  *string `tfsdk:"timeout" json:"timeout,omitempty"`
			} `tfsdk:"delete" json:"delete,omitempty"`
			Describe *struct {
				ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
				Cluster    *string `tfsdk:"cluster" json:"cluster,omitempty"`
				Clusters   *struct {
					Context    *string `tfsdk:"context" json:"context,omitempty"`
					Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
				} `tfsdk:"clusters" json:"clusters,omitempty"`
				Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
				Name       *string `tfsdk:"name" json:"name,omitempty"`
				Namespace  *string `tfsdk:"namespace" json:"namespace,omitempty"`
				Selector   *string `tfsdk:"selector" json:"selector,omitempty"`
				ShowEvents *bool   `tfsdk:"show_events" json:"showEvents,omitempty"`
				Timeout    *string `tfsdk:"timeout" json:"timeout,omitempty"`
			} `tfsdk:"describe" json:"describe,omitempty"`
			Description *string `tfsdk:"description" json:"description,omitempty"`
			Events      *struct {
				Cluster  *string `tfsdk:"cluster" json:"cluster,omitempty"`
				Clusters *struct {
					Context    *string `tfsdk:"context" json:"context,omitempty"`
					Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
				} `tfsdk:"clusters" json:"clusters,omitempty"`
				Format    *string `tfsdk:"format" json:"format,omitempty"`
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				Selector  *string `tfsdk:"selector" json:"selector,omitempty"`
				Timeout   *string `tfsdk:"timeout" json:"timeout,omitempty"`
			} `tfsdk:"events" json:"events,omitempty"`
			Get *struct {
				ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
				Cluster    *string `tfsdk:"cluster" json:"cluster,omitempty"`
				Clusters   *struct {
					Context    *string `tfsdk:"context" json:"context,omitempty"`
					Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
				} `tfsdk:"clusters" json:"clusters,omitempty"`
				Format    *string `tfsdk:"format" json:"format,omitempty"`
				Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				Selector  *string `tfsdk:"selector" json:"selector,omitempty"`
				Timeout   *string `tfsdk:"timeout" json:"timeout,omitempty"`
			} `tfsdk:"get" json:"get,omitempty"`
			PodLogs *struct {
				Cluster  *string `tfsdk:"cluster" json:"cluster,omitempty"`
				Clusters *struct {
					Context    *string `tfsdk:"context" json:"context,omitempty"`
					Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
				} `tfsdk:"clusters" json:"clusters,omitempty"`
				Container *string `tfsdk:"container" json:"container,omitempty"`
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				Selector  *string `tfsdk:"selector" json:"selector,omitempty"`
				Tail      *int64  `tfsdk:"tail" json:"tail,omitempty"`
				Timeout   *string `tfsdk:"timeout" json:"timeout,omitempty"`
			} `tfsdk:"pod_logs" json:"podLogs,omitempty"`
			Script *struct {
				Bindings *[]struct {
					Name  *string            `tfsdk:"name" json:"name,omitempty"`
					Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"bindings" json:"bindings,omitempty"`
				Check    *map[string]string `tfsdk:"check" json:"check,omitempty"`
				Cluster  *string            `tfsdk:"cluster" json:"cluster,omitempty"`
				Clusters *struct {
					Context    *string `tfsdk:"context" json:"context,omitempty"`
					Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
				} `tfsdk:"clusters" json:"clusters,omitempty"`
				Content *string `tfsdk:"content" json:"content,omitempty"`
				Env     *[]struct {
					Name  *string            `tfsdk:"name" json:"name,omitempty"`
					Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"env" json:"env,omitempty"`
				Outputs *[]struct {
					Match *map[string]string `tfsdk:"match" json:"match,omitempty"`
					Name  *string            `tfsdk:"name" json:"name,omitempty"`
					Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"outputs" json:"outputs,omitempty"`
				SkipLogOutput *bool   `tfsdk:"skip_log_output" json:"skipLogOutput,omitempty"`
				Timeout       *string `tfsdk:"timeout" json:"timeout,omitempty"`
			} `tfsdk:"script" json:"script,omitempty"`
			Sleep *struct {
				Duration *string `tfsdk:"duration" json:"duration,omitempty"`
			} `tfsdk:"sleep" json:"sleep,omitempty"`
			Wait *struct {
				ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
				Cluster    *string `tfsdk:"cluster" json:"cluster,omitempty"`
				Clusters   *struct {
					Context    *string `tfsdk:"context" json:"context,omitempty"`
					Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
				} `tfsdk:"clusters" json:"clusters,omitempty"`
				For *struct {
					Condition *struct {
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"condition" json:"condition,omitempty"`
					Deletion *map[string]string `tfsdk:"deletion" json:"deletion,omitempty"`
					JsonPath *struct {
						Path  *string `tfsdk:"path" json:"path,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"json_path" json:"jsonPath,omitempty"`
				} `tfsdk:"for" json:"for,omitempty"`
				Format    *string `tfsdk:"format" json:"format,omitempty"`
				Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				Selector  *string `tfsdk:"selector" json:"selector,omitempty"`
				Timeout   *string `tfsdk:"timeout" json:"timeout,omitempty"`
			} `tfsdk:"wait" json:"wait,omitempty"`
		} `tfsdk:"catch" json:"catch,omitempty"`
		Cluster  *string `tfsdk:"cluster" json:"cluster,omitempty"`
		Clusters *struct {
			Context    *string `tfsdk:"context" json:"context,omitempty"`
			Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
		} `tfsdk:"clusters" json:"clusters,omitempty"`
		Concurrent                  *bool              `tfsdk:"concurrent" json:"concurrent,omitempty"`
		DelayBeforeCleanup          *string            `tfsdk:"delay_before_cleanup" json:"delayBeforeCleanup,omitempty"`
		DeletionPropagationPolicy   *string            `tfsdk:"deletion_propagation_policy" json:"deletionPropagationPolicy,omitempty"`
		Description                 *string            `tfsdk:"description" json:"description,omitempty"`
		ForceTerminationGracePeriod *string            `tfsdk:"force_termination_grace_period" json:"forceTerminationGracePeriod,omitempty"`
		Namespace                   *string            `tfsdk:"namespace" json:"namespace,omitempty"`
		NamespaceTemplate           *map[string]string `tfsdk:"namespace_template" json:"namespaceTemplate,omitempty"`
		Skip                        *bool              `tfsdk:"skip" json:"skip,omitempty"`
		SkipDelete                  *bool              `tfsdk:"skip_delete" json:"skipDelete,omitempty"`
		Steps                       *[]struct {
			Bindings *[]struct {
				Name  *string            `tfsdk:"name" json:"name,omitempty"`
				Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"bindings" json:"bindings,omitempty"`
			Catch *[]struct {
				Command *struct {
					Args     *[]string `tfsdk:"args" json:"args,omitempty"`
					Bindings *[]struct {
						Name  *string            `tfsdk:"name" json:"name,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"bindings" json:"bindings,omitempty"`
					Check    *map[string]string `tfsdk:"check" json:"check,omitempty"`
					Cluster  *string            `tfsdk:"cluster" json:"cluster,omitempty"`
					Clusters *struct {
						Context    *string `tfsdk:"context" json:"context,omitempty"`
						Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
					} `tfsdk:"clusters" json:"clusters,omitempty"`
					Entrypoint *string `tfsdk:"entrypoint" json:"entrypoint,omitempty"`
					Env        *[]struct {
						Name  *string            `tfsdk:"name" json:"name,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"env" json:"env,omitempty"`
					Outputs *[]struct {
						Match *map[string]string `tfsdk:"match" json:"match,omitempty"`
						Name  *string            `tfsdk:"name" json:"name,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"outputs" json:"outputs,omitempty"`
					SkipLogOutput *bool   `tfsdk:"skip_log_output" json:"skipLogOutput,omitempty"`
					Timeout       *string `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"command" json:"command,omitempty"`
				Delete *struct {
					Bindings *[]struct {
						Name  *string            `tfsdk:"name" json:"name,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"bindings" json:"bindings,omitempty"`
					Cluster  *string `tfsdk:"cluster" json:"cluster,omitempty"`
					Clusters *struct {
						Context    *string `tfsdk:"context" json:"context,omitempty"`
						Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
					} `tfsdk:"clusters" json:"clusters,omitempty"`
					DeletionPropagationPolicy *string `tfsdk:"deletion_propagation_policy" json:"deletionPropagationPolicy,omitempty"`
					Expect                    *[]struct {
						Check *map[string]string `tfsdk:"check" json:"check,omitempty"`
						Match *map[string]string `tfsdk:"match" json:"match,omitempty"`
					} `tfsdk:"expect" json:"expect,omitempty"`
					File *string `tfsdk:"file" json:"file,omitempty"`
					Ref  *struct {
						ApiVersion *string            `tfsdk:"api_version" json:"apiVersion,omitempty"`
						Kind       *string            `tfsdk:"kind" json:"kind,omitempty"`
						Labels     *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
						Name       *string            `tfsdk:"name" json:"name,omitempty"`
						Namespace  *string            `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"ref" json:"ref,omitempty"`
					Template *bool   `tfsdk:"template" json:"template,omitempty"`
					Timeout  *string `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"delete" json:"delete,omitempty"`
				Describe *struct {
					ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
					Cluster    *string `tfsdk:"cluster" json:"cluster,omitempty"`
					Clusters   *struct {
						Context    *string `tfsdk:"context" json:"context,omitempty"`
						Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
					} `tfsdk:"clusters" json:"clusters,omitempty"`
					Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
					Name       *string `tfsdk:"name" json:"name,omitempty"`
					Namespace  *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Selector   *string `tfsdk:"selector" json:"selector,omitempty"`
					ShowEvents *bool   `tfsdk:"show_events" json:"showEvents,omitempty"`
					Timeout    *string `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"describe" json:"describe,omitempty"`
				Description *string `tfsdk:"description" json:"description,omitempty"`
				Events      *struct {
					Cluster  *string `tfsdk:"cluster" json:"cluster,omitempty"`
					Clusters *struct {
						Context    *string `tfsdk:"context" json:"context,omitempty"`
						Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
					} `tfsdk:"clusters" json:"clusters,omitempty"`
					Format    *string `tfsdk:"format" json:"format,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Selector  *string `tfsdk:"selector" json:"selector,omitempty"`
					Timeout   *string `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"events" json:"events,omitempty"`
				Get *struct {
					ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
					Cluster    *string `tfsdk:"cluster" json:"cluster,omitempty"`
					Clusters   *struct {
						Context    *string `tfsdk:"context" json:"context,omitempty"`
						Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
					} `tfsdk:"clusters" json:"clusters,omitempty"`
					Format    *string `tfsdk:"format" json:"format,omitempty"`
					Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Selector  *string `tfsdk:"selector" json:"selector,omitempty"`
					Timeout   *string `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"get" json:"get,omitempty"`
				PodLogs *struct {
					Cluster  *string `tfsdk:"cluster" json:"cluster,omitempty"`
					Clusters *struct {
						Context    *string `tfsdk:"context" json:"context,omitempty"`
						Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
					} `tfsdk:"clusters" json:"clusters,omitempty"`
					Container *string `tfsdk:"container" json:"container,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Selector  *string `tfsdk:"selector" json:"selector,omitempty"`
					Tail      *int64  `tfsdk:"tail" json:"tail,omitempty"`
					Timeout   *string `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"pod_logs" json:"podLogs,omitempty"`
				Script *struct {
					Bindings *[]struct {
						Name  *string            `tfsdk:"name" json:"name,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"bindings" json:"bindings,omitempty"`
					Check    *map[string]string `tfsdk:"check" json:"check,omitempty"`
					Cluster  *string            `tfsdk:"cluster" json:"cluster,omitempty"`
					Clusters *struct {
						Context    *string `tfsdk:"context" json:"context,omitempty"`
						Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
					} `tfsdk:"clusters" json:"clusters,omitempty"`
					Content *string `tfsdk:"content" json:"content,omitempty"`
					Env     *[]struct {
						Name  *string            `tfsdk:"name" json:"name,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"env" json:"env,omitempty"`
					Outputs *[]struct {
						Match *map[string]string `tfsdk:"match" json:"match,omitempty"`
						Name  *string            `tfsdk:"name" json:"name,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"outputs" json:"outputs,omitempty"`
					SkipLogOutput *bool   `tfsdk:"skip_log_output" json:"skipLogOutput,omitempty"`
					Timeout       *string `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"script" json:"script,omitempty"`
				Sleep *struct {
					Duration *string `tfsdk:"duration" json:"duration,omitempty"`
				} `tfsdk:"sleep" json:"sleep,omitempty"`
				Wait *struct {
					ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
					Cluster    *string `tfsdk:"cluster" json:"cluster,omitempty"`
					Clusters   *struct {
						Context    *string `tfsdk:"context" json:"context,omitempty"`
						Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
					} `tfsdk:"clusters" json:"clusters,omitempty"`
					For *struct {
						Condition *struct {
							Name  *string `tfsdk:"name" json:"name,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"condition" json:"condition,omitempty"`
						Deletion *map[string]string `tfsdk:"deletion" json:"deletion,omitempty"`
						JsonPath *struct {
							Path  *string `tfsdk:"path" json:"path,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"json_path" json:"jsonPath,omitempty"`
					} `tfsdk:"for" json:"for,omitempty"`
					Format    *string `tfsdk:"format" json:"format,omitempty"`
					Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Selector  *string `tfsdk:"selector" json:"selector,omitempty"`
					Timeout   *string `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"wait" json:"wait,omitempty"`
			} `tfsdk:"catch" json:"catch,omitempty"`
			Cleanup *[]struct {
				Command *struct {
					Args     *[]string `tfsdk:"args" json:"args,omitempty"`
					Bindings *[]struct {
						Name  *string            `tfsdk:"name" json:"name,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"bindings" json:"bindings,omitempty"`
					Check    *map[string]string `tfsdk:"check" json:"check,omitempty"`
					Cluster  *string            `tfsdk:"cluster" json:"cluster,omitempty"`
					Clusters *struct {
						Context    *string `tfsdk:"context" json:"context,omitempty"`
						Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
					} `tfsdk:"clusters" json:"clusters,omitempty"`
					Entrypoint *string `tfsdk:"entrypoint" json:"entrypoint,omitempty"`
					Env        *[]struct {
						Name  *string            `tfsdk:"name" json:"name,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"env" json:"env,omitempty"`
					Outputs *[]struct {
						Match *map[string]string `tfsdk:"match" json:"match,omitempty"`
						Name  *string            `tfsdk:"name" json:"name,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"outputs" json:"outputs,omitempty"`
					SkipLogOutput *bool   `tfsdk:"skip_log_output" json:"skipLogOutput,omitempty"`
					Timeout       *string `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"command" json:"command,omitempty"`
				Delete *struct {
					Bindings *[]struct {
						Name  *string            `tfsdk:"name" json:"name,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"bindings" json:"bindings,omitempty"`
					Cluster  *string `tfsdk:"cluster" json:"cluster,omitempty"`
					Clusters *struct {
						Context    *string `tfsdk:"context" json:"context,omitempty"`
						Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
					} `tfsdk:"clusters" json:"clusters,omitempty"`
					DeletionPropagationPolicy *string `tfsdk:"deletion_propagation_policy" json:"deletionPropagationPolicy,omitempty"`
					Expect                    *[]struct {
						Check *map[string]string `tfsdk:"check" json:"check,omitempty"`
						Match *map[string]string `tfsdk:"match" json:"match,omitempty"`
					} `tfsdk:"expect" json:"expect,omitempty"`
					File *string `tfsdk:"file" json:"file,omitempty"`
					Ref  *struct {
						ApiVersion *string            `tfsdk:"api_version" json:"apiVersion,omitempty"`
						Kind       *string            `tfsdk:"kind" json:"kind,omitempty"`
						Labels     *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
						Name       *string            `tfsdk:"name" json:"name,omitempty"`
						Namespace  *string            `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"ref" json:"ref,omitempty"`
					Template *bool   `tfsdk:"template" json:"template,omitempty"`
					Timeout  *string `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"delete" json:"delete,omitempty"`
				Describe *struct {
					ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
					Cluster    *string `tfsdk:"cluster" json:"cluster,omitempty"`
					Clusters   *struct {
						Context    *string `tfsdk:"context" json:"context,omitempty"`
						Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
					} `tfsdk:"clusters" json:"clusters,omitempty"`
					Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
					Name       *string `tfsdk:"name" json:"name,omitempty"`
					Namespace  *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Selector   *string `tfsdk:"selector" json:"selector,omitempty"`
					ShowEvents *bool   `tfsdk:"show_events" json:"showEvents,omitempty"`
					Timeout    *string `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"describe" json:"describe,omitempty"`
				Description *string `tfsdk:"description" json:"description,omitempty"`
				Events      *struct {
					Cluster  *string `tfsdk:"cluster" json:"cluster,omitempty"`
					Clusters *struct {
						Context    *string `tfsdk:"context" json:"context,omitempty"`
						Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
					} `tfsdk:"clusters" json:"clusters,omitempty"`
					Format    *string `tfsdk:"format" json:"format,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Selector  *string `tfsdk:"selector" json:"selector,omitempty"`
					Timeout   *string `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"events" json:"events,omitempty"`
				Get *struct {
					ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
					Cluster    *string `tfsdk:"cluster" json:"cluster,omitempty"`
					Clusters   *struct {
						Context    *string `tfsdk:"context" json:"context,omitempty"`
						Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
					} `tfsdk:"clusters" json:"clusters,omitempty"`
					Format    *string `tfsdk:"format" json:"format,omitempty"`
					Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Selector  *string `tfsdk:"selector" json:"selector,omitempty"`
					Timeout   *string `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"get" json:"get,omitempty"`
				PodLogs *struct {
					Cluster  *string `tfsdk:"cluster" json:"cluster,omitempty"`
					Clusters *struct {
						Context    *string `tfsdk:"context" json:"context,omitempty"`
						Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
					} `tfsdk:"clusters" json:"clusters,omitempty"`
					Container *string `tfsdk:"container" json:"container,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Selector  *string `tfsdk:"selector" json:"selector,omitempty"`
					Tail      *int64  `tfsdk:"tail" json:"tail,omitempty"`
					Timeout   *string `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"pod_logs" json:"podLogs,omitempty"`
				Script *struct {
					Bindings *[]struct {
						Name  *string            `tfsdk:"name" json:"name,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"bindings" json:"bindings,omitempty"`
					Check    *map[string]string `tfsdk:"check" json:"check,omitempty"`
					Cluster  *string            `tfsdk:"cluster" json:"cluster,omitempty"`
					Clusters *struct {
						Context    *string `tfsdk:"context" json:"context,omitempty"`
						Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
					} `tfsdk:"clusters" json:"clusters,omitempty"`
					Content *string `tfsdk:"content" json:"content,omitempty"`
					Env     *[]struct {
						Name  *string            `tfsdk:"name" json:"name,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"env" json:"env,omitempty"`
					Outputs *[]struct {
						Match *map[string]string `tfsdk:"match" json:"match,omitempty"`
						Name  *string            `tfsdk:"name" json:"name,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"outputs" json:"outputs,omitempty"`
					SkipLogOutput *bool   `tfsdk:"skip_log_output" json:"skipLogOutput,omitempty"`
					Timeout       *string `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"script" json:"script,omitempty"`
				Sleep *struct {
					Duration *string `tfsdk:"duration" json:"duration,omitempty"`
				} `tfsdk:"sleep" json:"sleep,omitempty"`
				Wait *struct {
					ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
					Cluster    *string `tfsdk:"cluster" json:"cluster,omitempty"`
					Clusters   *struct {
						Context    *string `tfsdk:"context" json:"context,omitempty"`
						Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
					} `tfsdk:"clusters" json:"clusters,omitempty"`
					For *struct {
						Condition *struct {
							Name  *string `tfsdk:"name" json:"name,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"condition" json:"condition,omitempty"`
						Deletion *map[string]string `tfsdk:"deletion" json:"deletion,omitempty"`
						JsonPath *struct {
							Path  *string `tfsdk:"path" json:"path,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"json_path" json:"jsonPath,omitempty"`
					} `tfsdk:"for" json:"for,omitempty"`
					Format    *string `tfsdk:"format" json:"format,omitempty"`
					Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Selector  *string `tfsdk:"selector" json:"selector,omitempty"`
					Timeout   *string `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"wait" json:"wait,omitempty"`
			} `tfsdk:"cleanup" json:"cleanup,omitempty"`
			Cluster  *string `tfsdk:"cluster" json:"cluster,omitempty"`
			Clusters *struct {
				Context    *string `tfsdk:"context" json:"context,omitempty"`
				Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
			} `tfsdk:"clusters" json:"clusters,omitempty"`
			DeletionPropagationPolicy *string `tfsdk:"deletion_propagation_policy" json:"deletionPropagationPolicy,omitempty"`
			Description               *string `tfsdk:"description" json:"description,omitempty"`
			Finally                   *[]struct {
				Command *struct {
					Args     *[]string `tfsdk:"args" json:"args,omitempty"`
					Bindings *[]struct {
						Name  *string            `tfsdk:"name" json:"name,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"bindings" json:"bindings,omitempty"`
					Check    *map[string]string `tfsdk:"check" json:"check,omitempty"`
					Cluster  *string            `tfsdk:"cluster" json:"cluster,omitempty"`
					Clusters *struct {
						Context    *string `tfsdk:"context" json:"context,omitempty"`
						Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
					} `tfsdk:"clusters" json:"clusters,omitempty"`
					Entrypoint *string `tfsdk:"entrypoint" json:"entrypoint,omitempty"`
					Env        *[]struct {
						Name  *string            `tfsdk:"name" json:"name,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"env" json:"env,omitempty"`
					Outputs *[]struct {
						Match *map[string]string `tfsdk:"match" json:"match,omitempty"`
						Name  *string            `tfsdk:"name" json:"name,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"outputs" json:"outputs,omitempty"`
					SkipLogOutput *bool   `tfsdk:"skip_log_output" json:"skipLogOutput,omitempty"`
					Timeout       *string `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"command" json:"command,omitempty"`
				Delete *struct {
					Bindings *[]struct {
						Name  *string            `tfsdk:"name" json:"name,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"bindings" json:"bindings,omitempty"`
					Cluster  *string `tfsdk:"cluster" json:"cluster,omitempty"`
					Clusters *struct {
						Context    *string `tfsdk:"context" json:"context,omitempty"`
						Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
					} `tfsdk:"clusters" json:"clusters,omitempty"`
					DeletionPropagationPolicy *string `tfsdk:"deletion_propagation_policy" json:"deletionPropagationPolicy,omitempty"`
					Expect                    *[]struct {
						Check *map[string]string `tfsdk:"check" json:"check,omitempty"`
						Match *map[string]string `tfsdk:"match" json:"match,omitempty"`
					} `tfsdk:"expect" json:"expect,omitempty"`
					File *string `tfsdk:"file" json:"file,omitempty"`
					Ref  *struct {
						ApiVersion *string            `tfsdk:"api_version" json:"apiVersion,omitempty"`
						Kind       *string            `tfsdk:"kind" json:"kind,omitempty"`
						Labels     *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
						Name       *string            `tfsdk:"name" json:"name,omitempty"`
						Namespace  *string            `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"ref" json:"ref,omitempty"`
					Template *bool   `tfsdk:"template" json:"template,omitempty"`
					Timeout  *string `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"delete" json:"delete,omitempty"`
				Describe *struct {
					ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
					Cluster    *string `tfsdk:"cluster" json:"cluster,omitempty"`
					Clusters   *struct {
						Context    *string `tfsdk:"context" json:"context,omitempty"`
						Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
					} `tfsdk:"clusters" json:"clusters,omitempty"`
					Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
					Name       *string `tfsdk:"name" json:"name,omitempty"`
					Namespace  *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Selector   *string `tfsdk:"selector" json:"selector,omitempty"`
					ShowEvents *bool   `tfsdk:"show_events" json:"showEvents,omitempty"`
					Timeout    *string `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"describe" json:"describe,omitempty"`
				Description *string `tfsdk:"description" json:"description,omitempty"`
				Events      *struct {
					Cluster  *string `tfsdk:"cluster" json:"cluster,omitempty"`
					Clusters *struct {
						Context    *string `tfsdk:"context" json:"context,omitempty"`
						Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
					} `tfsdk:"clusters" json:"clusters,omitempty"`
					Format    *string `tfsdk:"format" json:"format,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Selector  *string `tfsdk:"selector" json:"selector,omitempty"`
					Timeout   *string `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"events" json:"events,omitempty"`
				Get *struct {
					ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
					Cluster    *string `tfsdk:"cluster" json:"cluster,omitempty"`
					Clusters   *struct {
						Context    *string `tfsdk:"context" json:"context,omitempty"`
						Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
					} `tfsdk:"clusters" json:"clusters,omitempty"`
					Format    *string `tfsdk:"format" json:"format,omitempty"`
					Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Selector  *string `tfsdk:"selector" json:"selector,omitempty"`
					Timeout   *string `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"get" json:"get,omitempty"`
				PodLogs *struct {
					Cluster  *string `tfsdk:"cluster" json:"cluster,omitempty"`
					Clusters *struct {
						Context    *string `tfsdk:"context" json:"context,omitempty"`
						Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
					} `tfsdk:"clusters" json:"clusters,omitempty"`
					Container *string `tfsdk:"container" json:"container,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Selector  *string `tfsdk:"selector" json:"selector,omitempty"`
					Tail      *int64  `tfsdk:"tail" json:"tail,omitempty"`
					Timeout   *string `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"pod_logs" json:"podLogs,omitempty"`
				Script *struct {
					Bindings *[]struct {
						Name  *string            `tfsdk:"name" json:"name,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"bindings" json:"bindings,omitempty"`
					Check    *map[string]string `tfsdk:"check" json:"check,omitempty"`
					Cluster  *string            `tfsdk:"cluster" json:"cluster,omitempty"`
					Clusters *struct {
						Context    *string `tfsdk:"context" json:"context,omitempty"`
						Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
					} `tfsdk:"clusters" json:"clusters,omitempty"`
					Content *string `tfsdk:"content" json:"content,omitempty"`
					Env     *[]struct {
						Name  *string            `tfsdk:"name" json:"name,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"env" json:"env,omitempty"`
					Outputs *[]struct {
						Match *map[string]string `tfsdk:"match" json:"match,omitempty"`
						Name  *string            `tfsdk:"name" json:"name,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"outputs" json:"outputs,omitempty"`
					SkipLogOutput *bool   `tfsdk:"skip_log_output" json:"skipLogOutput,omitempty"`
					Timeout       *string `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"script" json:"script,omitempty"`
				Sleep *struct {
					Duration *string `tfsdk:"duration" json:"duration,omitempty"`
				} `tfsdk:"sleep" json:"sleep,omitempty"`
				Wait *struct {
					ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
					Cluster    *string `tfsdk:"cluster" json:"cluster,omitempty"`
					Clusters   *struct {
						Context    *string `tfsdk:"context" json:"context,omitempty"`
						Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
					} `tfsdk:"clusters" json:"clusters,omitempty"`
					For *struct {
						Condition *struct {
							Name  *string `tfsdk:"name" json:"name,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"condition" json:"condition,omitempty"`
						Deletion *map[string]string `tfsdk:"deletion" json:"deletion,omitempty"`
						JsonPath *struct {
							Path  *string `tfsdk:"path" json:"path,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"json_path" json:"jsonPath,omitempty"`
					} `tfsdk:"for" json:"for,omitempty"`
					Format    *string `tfsdk:"format" json:"format,omitempty"`
					Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Selector  *string `tfsdk:"selector" json:"selector,omitempty"`
					Timeout   *string `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"wait" json:"wait,omitempty"`
			} `tfsdk:"finally" json:"finally,omitempty"`
			Name       *string `tfsdk:"name" json:"name,omitempty"`
			SkipDelete *bool   `tfsdk:"skip_delete" json:"skipDelete,omitempty"`
			Template   *bool   `tfsdk:"template" json:"template,omitempty"`
			Timeouts   *struct {
				Apply   *string `tfsdk:"apply" json:"apply,omitempty"`
				Assert  *string `tfsdk:"assert" json:"assert,omitempty"`
				Cleanup *string `tfsdk:"cleanup" json:"cleanup,omitempty"`
				Delete  *string `tfsdk:"delete" json:"delete,omitempty"`
				Error   *string `tfsdk:"error" json:"error,omitempty"`
				Exec    *string `tfsdk:"exec" json:"exec,omitempty"`
			} `tfsdk:"timeouts" json:"timeouts,omitempty"`
			Try *[]struct {
				Apply *struct {
					Bindings *[]struct {
						Name  *string            `tfsdk:"name" json:"name,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"bindings" json:"bindings,omitempty"`
					Cluster  *string `tfsdk:"cluster" json:"cluster,omitempty"`
					Clusters *struct {
						Context    *string `tfsdk:"context" json:"context,omitempty"`
						Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
					} `tfsdk:"clusters" json:"clusters,omitempty"`
					DryRun *bool `tfsdk:"dry_run" json:"dryRun,omitempty"`
					Expect *[]struct {
						Check *map[string]string `tfsdk:"check" json:"check,omitempty"`
						Match *map[string]string `tfsdk:"match" json:"match,omitempty"`
					} `tfsdk:"expect" json:"expect,omitempty"`
					File    *string `tfsdk:"file" json:"file,omitempty"`
					Outputs *[]struct {
						Match *map[string]string `tfsdk:"match" json:"match,omitempty"`
						Name  *string            `tfsdk:"name" json:"name,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"outputs" json:"outputs,omitempty"`
					Resource *map[string]string `tfsdk:"resource" json:"resource,omitempty"`
					Template *bool              `tfsdk:"template" json:"template,omitempty"`
					Timeout  *string            `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"apply" json:"apply,omitempty"`
				Assert *struct {
					Bindings *[]struct {
						Name  *string            `tfsdk:"name" json:"name,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"bindings" json:"bindings,omitempty"`
					Cluster  *string `tfsdk:"cluster" json:"cluster,omitempty"`
					Clusters *struct {
						Context    *string `tfsdk:"context" json:"context,omitempty"`
						Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
					} `tfsdk:"clusters" json:"clusters,omitempty"`
					File     *string            `tfsdk:"file" json:"file,omitempty"`
					Resource *map[string]string `tfsdk:"resource" json:"resource,omitempty"`
					Template *bool              `tfsdk:"template" json:"template,omitempty"`
					Timeout  *string            `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"assert" json:"assert,omitempty"`
				Command *struct {
					Args     *[]string `tfsdk:"args" json:"args,omitempty"`
					Bindings *[]struct {
						Name  *string            `tfsdk:"name" json:"name,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"bindings" json:"bindings,omitempty"`
					Check    *map[string]string `tfsdk:"check" json:"check,omitempty"`
					Cluster  *string            `tfsdk:"cluster" json:"cluster,omitempty"`
					Clusters *struct {
						Context    *string `tfsdk:"context" json:"context,omitempty"`
						Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
					} `tfsdk:"clusters" json:"clusters,omitempty"`
					Entrypoint *string `tfsdk:"entrypoint" json:"entrypoint,omitempty"`
					Env        *[]struct {
						Name  *string            `tfsdk:"name" json:"name,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"env" json:"env,omitempty"`
					Outputs *[]struct {
						Match *map[string]string `tfsdk:"match" json:"match,omitempty"`
						Name  *string            `tfsdk:"name" json:"name,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"outputs" json:"outputs,omitempty"`
					SkipLogOutput *bool   `tfsdk:"skip_log_output" json:"skipLogOutput,omitempty"`
					Timeout       *string `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"command" json:"command,omitempty"`
				ContinueOnError *bool `tfsdk:"continue_on_error" json:"continueOnError,omitempty"`
				Create          *struct {
					Bindings *[]struct {
						Name  *string            `tfsdk:"name" json:"name,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"bindings" json:"bindings,omitempty"`
					Cluster  *string `tfsdk:"cluster" json:"cluster,omitempty"`
					Clusters *struct {
						Context    *string `tfsdk:"context" json:"context,omitempty"`
						Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
					} `tfsdk:"clusters" json:"clusters,omitempty"`
					DryRun *bool `tfsdk:"dry_run" json:"dryRun,omitempty"`
					Expect *[]struct {
						Check *map[string]string `tfsdk:"check" json:"check,omitempty"`
						Match *map[string]string `tfsdk:"match" json:"match,omitempty"`
					} `tfsdk:"expect" json:"expect,omitempty"`
					File    *string `tfsdk:"file" json:"file,omitempty"`
					Outputs *[]struct {
						Match *map[string]string `tfsdk:"match" json:"match,omitempty"`
						Name  *string            `tfsdk:"name" json:"name,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"outputs" json:"outputs,omitempty"`
					Resource *map[string]string `tfsdk:"resource" json:"resource,omitempty"`
					Template *bool              `tfsdk:"template" json:"template,omitempty"`
					Timeout  *string            `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"create" json:"create,omitempty"`
				Delete *struct {
					Bindings *[]struct {
						Name  *string            `tfsdk:"name" json:"name,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"bindings" json:"bindings,omitempty"`
					Cluster  *string `tfsdk:"cluster" json:"cluster,omitempty"`
					Clusters *struct {
						Context    *string `tfsdk:"context" json:"context,omitempty"`
						Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
					} `tfsdk:"clusters" json:"clusters,omitempty"`
					DeletionPropagationPolicy *string `tfsdk:"deletion_propagation_policy" json:"deletionPropagationPolicy,omitempty"`
					Expect                    *[]struct {
						Check *map[string]string `tfsdk:"check" json:"check,omitempty"`
						Match *map[string]string `tfsdk:"match" json:"match,omitempty"`
					} `tfsdk:"expect" json:"expect,omitempty"`
					File *string `tfsdk:"file" json:"file,omitempty"`
					Ref  *struct {
						ApiVersion *string            `tfsdk:"api_version" json:"apiVersion,omitempty"`
						Kind       *string            `tfsdk:"kind" json:"kind,omitempty"`
						Labels     *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
						Name       *string            `tfsdk:"name" json:"name,omitempty"`
						Namespace  *string            `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"ref" json:"ref,omitempty"`
					Template *bool   `tfsdk:"template" json:"template,omitempty"`
					Timeout  *string `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"delete" json:"delete,omitempty"`
				Describe *struct {
					ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
					Cluster    *string `tfsdk:"cluster" json:"cluster,omitempty"`
					Clusters   *struct {
						Context    *string `tfsdk:"context" json:"context,omitempty"`
						Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
					} `tfsdk:"clusters" json:"clusters,omitempty"`
					Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
					Name       *string `tfsdk:"name" json:"name,omitempty"`
					Namespace  *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Selector   *string `tfsdk:"selector" json:"selector,omitempty"`
					ShowEvents *bool   `tfsdk:"show_events" json:"showEvents,omitempty"`
					Timeout    *string `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"describe" json:"describe,omitempty"`
				Description *string `tfsdk:"description" json:"description,omitempty"`
				Error       *struct {
					Bindings *[]struct {
						Name  *string            `tfsdk:"name" json:"name,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"bindings" json:"bindings,omitempty"`
					Cluster  *string `tfsdk:"cluster" json:"cluster,omitempty"`
					Clusters *struct {
						Context    *string `tfsdk:"context" json:"context,omitempty"`
						Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
					} `tfsdk:"clusters" json:"clusters,omitempty"`
					File     *string            `tfsdk:"file" json:"file,omitempty"`
					Resource *map[string]string `tfsdk:"resource" json:"resource,omitempty"`
					Template *bool              `tfsdk:"template" json:"template,omitempty"`
					Timeout  *string            `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"error" json:"error,omitempty"`
				Events *struct {
					Cluster  *string `tfsdk:"cluster" json:"cluster,omitempty"`
					Clusters *struct {
						Context    *string `tfsdk:"context" json:"context,omitempty"`
						Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
					} `tfsdk:"clusters" json:"clusters,omitempty"`
					Format    *string `tfsdk:"format" json:"format,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Selector  *string `tfsdk:"selector" json:"selector,omitempty"`
					Timeout   *string `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"events" json:"events,omitempty"`
				Get *struct {
					ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
					Cluster    *string `tfsdk:"cluster" json:"cluster,omitempty"`
					Clusters   *struct {
						Context    *string `tfsdk:"context" json:"context,omitempty"`
						Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
					} `tfsdk:"clusters" json:"clusters,omitempty"`
					Format    *string `tfsdk:"format" json:"format,omitempty"`
					Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Selector  *string `tfsdk:"selector" json:"selector,omitempty"`
					Timeout   *string `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"get" json:"get,omitempty"`
				Patch *struct {
					Bindings *[]struct {
						Name  *string            `tfsdk:"name" json:"name,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"bindings" json:"bindings,omitempty"`
					Cluster  *string `tfsdk:"cluster" json:"cluster,omitempty"`
					Clusters *struct {
						Context    *string `tfsdk:"context" json:"context,omitempty"`
						Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
					} `tfsdk:"clusters" json:"clusters,omitempty"`
					DryRun *bool `tfsdk:"dry_run" json:"dryRun,omitempty"`
					Expect *[]struct {
						Check *map[string]string `tfsdk:"check" json:"check,omitempty"`
						Match *map[string]string `tfsdk:"match" json:"match,omitempty"`
					} `tfsdk:"expect" json:"expect,omitempty"`
					File    *string `tfsdk:"file" json:"file,omitempty"`
					Outputs *[]struct {
						Match *map[string]string `tfsdk:"match" json:"match,omitempty"`
						Name  *string            `tfsdk:"name" json:"name,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"outputs" json:"outputs,omitempty"`
					Resource *map[string]string `tfsdk:"resource" json:"resource,omitempty"`
					Template *bool              `tfsdk:"template" json:"template,omitempty"`
					Timeout  *string            `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"patch" json:"patch,omitempty"`
				PodLogs *struct {
					Cluster  *string `tfsdk:"cluster" json:"cluster,omitempty"`
					Clusters *struct {
						Context    *string `tfsdk:"context" json:"context,omitempty"`
						Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
					} `tfsdk:"clusters" json:"clusters,omitempty"`
					Container *string `tfsdk:"container" json:"container,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Selector  *string `tfsdk:"selector" json:"selector,omitempty"`
					Tail      *int64  `tfsdk:"tail" json:"tail,omitempty"`
					Timeout   *string `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"pod_logs" json:"podLogs,omitempty"`
				Script *struct {
					Bindings *[]struct {
						Name  *string            `tfsdk:"name" json:"name,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"bindings" json:"bindings,omitempty"`
					Check    *map[string]string `tfsdk:"check" json:"check,omitempty"`
					Cluster  *string            `tfsdk:"cluster" json:"cluster,omitempty"`
					Clusters *struct {
						Context    *string `tfsdk:"context" json:"context,omitempty"`
						Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
					} `tfsdk:"clusters" json:"clusters,omitempty"`
					Content *string `tfsdk:"content" json:"content,omitempty"`
					Env     *[]struct {
						Name  *string            `tfsdk:"name" json:"name,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"env" json:"env,omitempty"`
					Outputs *[]struct {
						Match *map[string]string `tfsdk:"match" json:"match,omitempty"`
						Name  *string            `tfsdk:"name" json:"name,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"outputs" json:"outputs,omitempty"`
					SkipLogOutput *bool   `tfsdk:"skip_log_output" json:"skipLogOutput,omitempty"`
					Timeout       *string `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"script" json:"script,omitempty"`
				Sleep *struct {
					Duration *string `tfsdk:"duration" json:"duration,omitempty"`
				} `tfsdk:"sleep" json:"sleep,omitempty"`
				Update *struct {
					Bindings *[]struct {
						Name  *string            `tfsdk:"name" json:"name,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"bindings" json:"bindings,omitempty"`
					Cluster  *string `tfsdk:"cluster" json:"cluster,omitempty"`
					Clusters *struct {
						Context    *string `tfsdk:"context" json:"context,omitempty"`
						Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
					} `tfsdk:"clusters" json:"clusters,omitempty"`
					DryRun *bool `tfsdk:"dry_run" json:"dryRun,omitempty"`
					Expect *[]struct {
						Check *map[string]string `tfsdk:"check" json:"check,omitempty"`
						Match *map[string]string `tfsdk:"match" json:"match,omitempty"`
					} `tfsdk:"expect" json:"expect,omitempty"`
					File    *string `tfsdk:"file" json:"file,omitempty"`
					Outputs *[]struct {
						Match *map[string]string `tfsdk:"match" json:"match,omitempty"`
						Name  *string            `tfsdk:"name" json:"name,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"outputs" json:"outputs,omitempty"`
					Resource *map[string]string `tfsdk:"resource" json:"resource,omitempty"`
					Template *bool              `tfsdk:"template" json:"template,omitempty"`
					Timeout  *string            `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"update" json:"update,omitempty"`
				Wait *struct {
					ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
					Cluster    *string `tfsdk:"cluster" json:"cluster,omitempty"`
					Clusters   *struct {
						Context    *string `tfsdk:"context" json:"context,omitempty"`
						Kubeconfig *string `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
					} `tfsdk:"clusters" json:"clusters,omitempty"`
					For *struct {
						Condition *struct {
							Name  *string `tfsdk:"name" json:"name,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"condition" json:"condition,omitempty"`
						Deletion *map[string]string `tfsdk:"deletion" json:"deletion,omitempty"`
						JsonPath *struct {
							Path  *string `tfsdk:"path" json:"path,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"json_path" json:"jsonPath,omitempty"`
					} `tfsdk:"for" json:"for,omitempty"`
					Format    *string `tfsdk:"format" json:"format,omitempty"`
					Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Selector  *string `tfsdk:"selector" json:"selector,omitempty"`
					Timeout   *string `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"wait" json:"wait,omitempty"`
			} `tfsdk:"try" json:"try,omitempty"`
		} `tfsdk:"steps" json:"steps,omitempty"`
		Template *bool `tfsdk:"template" json:"template,omitempty"`
		Timeouts *struct {
			Apply   *string `tfsdk:"apply" json:"apply,omitempty"`
			Assert  *string `tfsdk:"assert" json:"assert,omitempty"`
			Cleanup *string `tfsdk:"cleanup" json:"cleanup,omitempty"`
			Delete  *string `tfsdk:"delete" json:"delete,omitempty"`
			Error   *string `tfsdk:"error" json:"error,omitempty"`
			Exec    *string `tfsdk:"exec" json:"exec,omitempty"`
		} `tfsdk:"timeouts" json:"timeouts,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ChainsawKyvernoIoTestV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_chainsaw_kyverno_io_test_v1alpha1_manifest"
}

func (r *ChainsawKyvernoIoTestV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Test is the resource that contains a test definition.",
		MarkdownDescription: "Test is the resource that contains a test definition.",
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
				Description:         "Test spec.",
				MarkdownDescription: "Test spec.",
				Attributes: map[string]schema.Attribute{
					"bindings": schema.ListNestedAttribute{
						Description:         "Bindings defines additional binding key/values.",
						MarkdownDescription: "Bindings defines additional binding key/values.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name the name of the binding.",
									MarkdownDescription: "Name the name of the binding.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
									},
								},

								"value": schema.MapAttribute{
									Description:         "Value value of the binding.",
									MarkdownDescription: "Value value of the binding.",
									ElementType:         types.StringType,
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

					"catch": schema.ListNestedAttribute{
						Description:         "Catch defines what the steps will execute when an error happens.This will be combined with catch handlers defined at the step level.",
						MarkdownDescription: "Catch defines what the steps will execute when an error happens.This will be combined with catch handlers defined at the step level.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"command": schema.SingleNestedAttribute{
									Description:         "Command defines a command to run.",
									MarkdownDescription: "Command defines a command to run.",
									Attributes: map[string]schema.Attribute{
										"args": schema.ListAttribute{
											Description:         "Args is the command arguments.",
											MarkdownDescription: "Args is the command arguments.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"bindings": schema.ListNestedAttribute{
											Description:         "Bindings defines additional binding key/values.",
											MarkdownDescription: "Bindings defines additional binding key/values.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "Name the name of the binding.",
														MarkdownDescription: "Name the name of the binding.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
														},
													},

													"value": schema.MapAttribute{
														Description:         "Value value of the binding.",
														MarkdownDescription: "Value value of the binding.",
														ElementType:         types.StringType,
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

										"check": schema.MapAttribute{
											Description:         "Check is an assertion tree to validate the operation outcome.",
											MarkdownDescription: "Check is an assertion tree to validate the operation outcome.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"cluster": schema.StringAttribute{
											Description:         "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
											MarkdownDescription: "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"clusters": schema.SingleNestedAttribute{
											Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
											MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
											Attributes: map[string]schema.Attribute{
												"context": schema.StringAttribute{
													Description:         "Context is the name of the context to use.",
													MarkdownDescription: "Context is the name of the context to use.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"kubeconfig": schema.StringAttribute{
													Description:         "Kubeconfig is the path to the referenced file.",
													MarkdownDescription: "Kubeconfig is the path to the referenced file.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"entrypoint": schema.StringAttribute{
											Description:         "Entrypoint is the command entry point to run.",
											MarkdownDescription: "Entrypoint is the command entry point to run.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"env": schema.ListNestedAttribute{
											Description:         "Env defines additional environment variables.",
											MarkdownDescription: "Env defines additional environment variables.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "Name the name of the binding.",
														MarkdownDescription: "Name the name of the binding.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
														},
													},

													"value": schema.MapAttribute{
														Description:         "Value value of the binding.",
														MarkdownDescription: "Value value of the binding.",
														ElementType:         types.StringType,
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

										"outputs": schema.ListNestedAttribute{
											Description:         "Outputs defines output bindings.",
											MarkdownDescription: "Outputs defines output bindings.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"match": schema.MapAttribute{
														Description:         "Match defines the matching statement.",
														MarkdownDescription: "Match defines the matching statement.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name the name of the binding.",
														MarkdownDescription: "Name the name of the binding.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
														},
													},

													"value": schema.MapAttribute{
														Description:         "Value value of the binding.",
														MarkdownDescription: "Value value of the binding.",
														ElementType:         types.StringType,
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

										"skip_log_output": schema.BoolAttribute{
											Description:         "SkipLogOutput removes the output from the command. Useful for sensitive logs or to reduce noise.",
											MarkdownDescription: "SkipLogOutput removes the output from the command. Useful for sensitive logs or to reduce noise.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"timeout": schema.StringAttribute{
											Description:         "Timeout for the operation. Overrides the global timeout set in the Configuration.",
											MarkdownDescription: "Timeout for the operation. Overrides the global timeout set in the Configuration.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"delete": schema.SingleNestedAttribute{
									Description:         "Delete represents a deletion operation.",
									MarkdownDescription: "Delete represents a deletion operation.",
									Attributes: map[string]schema.Attribute{
										"bindings": schema.ListNestedAttribute{
											Description:         "Bindings defines additional binding key/values.",
											MarkdownDescription: "Bindings defines additional binding key/values.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "Name the name of the binding.",
														MarkdownDescription: "Name the name of the binding.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
														},
													},

													"value": schema.MapAttribute{
														Description:         "Value value of the binding.",
														MarkdownDescription: "Value value of the binding.",
														ElementType:         types.StringType,
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

										"cluster": schema.StringAttribute{
											Description:         "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
											MarkdownDescription: "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"clusters": schema.SingleNestedAttribute{
											Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
											MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
											Attributes: map[string]schema.Attribute{
												"context": schema.StringAttribute{
													Description:         "Context is the name of the context to use.",
													MarkdownDescription: "Context is the name of the context to use.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"kubeconfig": schema.StringAttribute{
													Description:         "Kubeconfig is the path to the referenced file.",
													MarkdownDescription: "Kubeconfig is the path to the referenced file.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"deletion_propagation_policy": schema.StringAttribute{
											Description:         "DeletionPropagationPolicy decides if a deletion will propagate to the dependents ofthe object, and how the garbage collector will handle the propagation.Overrides the deletion propagation policy set in the Configuration, the Test and the TestStep.",
											MarkdownDescription: "DeletionPropagationPolicy decides if a deletion will propagate to the dependents ofthe object, and how the garbage collector will handle the propagation.Overrides the deletion propagation policy set in the Configuration, the Test and the TestStep.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Orphan", "Background", "Foreground"),
											},
										},

										"expect": schema.ListNestedAttribute{
											Description:         "Expect defines a list of matched checks to validate the operation outcome.",
											MarkdownDescription: "Expect defines a list of matched checks to validate the operation outcome.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"check": schema.MapAttribute{
														Description:         "Check defines the verification statement.",
														MarkdownDescription: "Check defines the verification statement.",
														ElementType:         types.StringType,
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"match": schema.MapAttribute{
														Description:         "Match defines the matching statement.",
														MarkdownDescription: "Match defines the matching statement.",
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

										"file": schema.StringAttribute{
											Description:         "File is the path to the referenced file. This can be a direct path to a fileor an expression that matches multiple files, such as 'manifest/*.yaml' for all YAMLfiles within the 'manifest' directory.",
											MarkdownDescription: "File is the path to the referenced file. This can be a direct path to a fileor an expression that matches multiple files, such as 'manifest/*.yaml' for all YAMLfiles within the 'manifest' directory.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ref": schema.SingleNestedAttribute{
											Description:         "Ref determines objects to be deleted.",
											MarkdownDescription: "Ref determines objects to be deleted.",
											Attributes: map[string]schema.Attribute{
												"api_version": schema.StringAttribute{
													Description:         "API version of the referent.",
													MarkdownDescription: "API version of the referent.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"kind": schema.StringAttribute{
													Description:         "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
													MarkdownDescription: "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"labels": schema.MapAttribute{
													Description:         "Label selector to match objects to delete",
													MarkdownDescription: "Label selector to match objects to delete",
													ElementType:         types.StringType,
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"template": schema.BoolAttribute{
											Description:         "Template determines whether resources should be considered for templating.",
											MarkdownDescription: "Template determines whether resources should be considered for templating.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"timeout": schema.StringAttribute{
											Description:         "Timeout for the operation. Overrides the global timeout set in the Configuration.",
											MarkdownDescription: "Timeout for the operation. Overrides the global timeout set in the Configuration.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"describe": schema.SingleNestedAttribute{
									Description:         "Describe determines the resource describe collector to execute.",
									MarkdownDescription: "Describe determines the resource describe collector to execute.",
									Attributes: map[string]schema.Attribute{
										"api_version": schema.StringAttribute{
											Description:         "API version of the referent.",
											MarkdownDescription: "API version of the referent.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"cluster": schema.StringAttribute{
											Description:         "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
											MarkdownDescription: "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"clusters": schema.SingleNestedAttribute{
											Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
											MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
											Attributes: map[string]schema.Attribute{
												"context": schema.StringAttribute{
													Description:         "Context is the name of the context to use.",
													MarkdownDescription: "Context is the name of the context to use.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"kubeconfig": schema.StringAttribute{
													Description:         "Kubeconfig is the path to the referenced file.",
													MarkdownDescription: "Kubeconfig is the path to the referenced file.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"kind": schema.StringAttribute{
											Description:         "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
											MarkdownDescription: "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
											Required:            true,
											Optional:            false,
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

										"selector": schema.StringAttribute{
											Description:         "Selector defines labels selector.",
											MarkdownDescription: "Selector defines labels selector.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"show_events": schema.BoolAttribute{
											Description:         "Show Events indicates whether to include related events.",
											MarkdownDescription: "Show Events indicates whether to include related events.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"timeout": schema.StringAttribute{
											Description:         "Timeout for the operation. Overrides the global timeout set in the Configuration.",
											MarkdownDescription: "Timeout for the operation. Overrides the global timeout set in the Configuration.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"description": schema.StringAttribute{
									Description:         "Description contains a description of the operation.",
									MarkdownDescription: "Description contains a description of the operation.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"events": schema.SingleNestedAttribute{
									Description:         "Events determines the events collector to execute.",
									MarkdownDescription: "Events determines the events collector to execute.",
									Attributes: map[string]schema.Attribute{
										"cluster": schema.StringAttribute{
											Description:         "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
											MarkdownDescription: "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"clusters": schema.SingleNestedAttribute{
											Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
											MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
											Attributes: map[string]schema.Attribute{
												"context": schema.StringAttribute{
													Description:         "Context is the name of the context to use.",
													MarkdownDescription: "Context is the name of the context to use.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"kubeconfig": schema.StringAttribute{
													Description:         "Kubeconfig is the path to the referenced file.",
													MarkdownDescription: "Kubeconfig is the path to the referenced file.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"format": schema.StringAttribute{
											Description:         "Format determines the output format (json or yaml).",
											MarkdownDescription: "Format determines the output format (json or yaml).",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^(?:json|yaml|\(.+\))$`), ""),
											},
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

										"selector": schema.StringAttribute{
											Description:         "Selector defines labels selector.",
											MarkdownDescription: "Selector defines labels selector.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"timeout": schema.StringAttribute{
											Description:         "Timeout for the operation. Overrides the global timeout set in the Configuration.",
											MarkdownDescription: "Timeout for the operation. Overrides the global timeout set in the Configuration.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"get": schema.SingleNestedAttribute{
									Description:         "Get determines the resource get collector to execute.",
									MarkdownDescription: "Get determines the resource get collector to execute.",
									Attributes: map[string]schema.Attribute{
										"api_version": schema.StringAttribute{
											Description:         "API version of the referent.",
											MarkdownDescription: "API version of the referent.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"cluster": schema.StringAttribute{
											Description:         "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
											MarkdownDescription: "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"clusters": schema.SingleNestedAttribute{
											Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
											MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
											Attributes: map[string]schema.Attribute{
												"context": schema.StringAttribute{
													Description:         "Context is the name of the context to use.",
													MarkdownDescription: "Context is the name of the context to use.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"kubeconfig": schema.StringAttribute{
													Description:         "Kubeconfig is the path to the referenced file.",
													MarkdownDescription: "Kubeconfig is the path to the referenced file.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"format": schema.StringAttribute{
											Description:         "Format determines the output format (json or yaml).",
											MarkdownDescription: "Format determines the output format (json or yaml).",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^(?:json|yaml|\(.+\))$`), ""),
											},
										},

										"kind": schema.StringAttribute{
											Description:         "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
											MarkdownDescription: "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
											Required:            true,
											Optional:            false,
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

										"selector": schema.StringAttribute{
											Description:         "Selector defines labels selector.",
											MarkdownDescription: "Selector defines labels selector.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"timeout": schema.StringAttribute{
											Description:         "Timeout for the operation. Overrides the global timeout set in the Configuration.",
											MarkdownDescription: "Timeout for the operation. Overrides the global timeout set in the Configuration.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"pod_logs": schema.SingleNestedAttribute{
									Description:         "PodLogs determines the pod logs collector to execute.",
									MarkdownDescription: "PodLogs determines the pod logs collector to execute.",
									Attributes: map[string]schema.Attribute{
										"cluster": schema.StringAttribute{
											Description:         "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
											MarkdownDescription: "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"clusters": schema.SingleNestedAttribute{
											Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
											MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
											Attributes: map[string]schema.Attribute{
												"context": schema.StringAttribute{
													Description:         "Context is the name of the context to use.",
													MarkdownDescription: "Context is the name of the context to use.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"kubeconfig": schema.StringAttribute{
													Description:         "Kubeconfig is the path to the referenced file.",
													MarkdownDescription: "Kubeconfig is the path to the referenced file.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"container": schema.StringAttribute{
											Description:         "Container in pod to get logs from else --all-containers is used.",
											MarkdownDescription: "Container in pod to get logs from else --all-containers is used.",
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

										"selector": schema.StringAttribute{
											Description:         "Selector defines labels selector.",
											MarkdownDescription: "Selector defines labels selector.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tail": schema.Int64Attribute{
											Description:         "Tail is the number of last lines to collect from pods. If omitted or zero,then the default is 10 if you use a selector, or -1 (all) if you use a pod name.This matches default behavior of 'kubectl logs'.",
											MarkdownDescription: "Tail is the number of last lines to collect from pods. If omitted or zero,then the default is 10 if you use a selector, or -1 (all) if you use a pod name.This matches default behavior of 'kubectl logs'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"timeout": schema.StringAttribute{
											Description:         "Timeout for the operation. Overrides the global timeout set in the Configuration.",
											MarkdownDescription: "Timeout for the operation. Overrides the global timeout set in the Configuration.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"script": schema.SingleNestedAttribute{
									Description:         "Script defines a script to run.",
									MarkdownDescription: "Script defines a script to run.",
									Attributes: map[string]schema.Attribute{
										"bindings": schema.ListNestedAttribute{
											Description:         "Bindings defines additional binding key/values.",
											MarkdownDescription: "Bindings defines additional binding key/values.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "Name the name of the binding.",
														MarkdownDescription: "Name the name of the binding.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
														},
													},

													"value": schema.MapAttribute{
														Description:         "Value value of the binding.",
														MarkdownDescription: "Value value of the binding.",
														ElementType:         types.StringType,
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

										"check": schema.MapAttribute{
											Description:         "Check is an assertion tree to validate the operation outcome.",
											MarkdownDescription: "Check is an assertion tree to validate the operation outcome.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"cluster": schema.StringAttribute{
											Description:         "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
											MarkdownDescription: "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"clusters": schema.SingleNestedAttribute{
											Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
											MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
											Attributes: map[string]schema.Attribute{
												"context": schema.StringAttribute{
													Description:         "Context is the name of the context to use.",
													MarkdownDescription: "Context is the name of the context to use.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"kubeconfig": schema.StringAttribute{
													Description:         "Kubeconfig is the path to the referenced file.",
													MarkdownDescription: "Kubeconfig is the path to the referenced file.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"content": schema.StringAttribute{
											Description:         "Content defines a shell script (run with 'sh -c ...').",
											MarkdownDescription: "Content defines a shell script (run with 'sh -c ...').",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"env": schema.ListNestedAttribute{
											Description:         "Env defines additional environment variables.",
											MarkdownDescription: "Env defines additional environment variables.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "Name the name of the binding.",
														MarkdownDescription: "Name the name of the binding.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
														},
													},

													"value": schema.MapAttribute{
														Description:         "Value value of the binding.",
														MarkdownDescription: "Value value of the binding.",
														ElementType:         types.StringType,
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

										"outputs": schema.ListNestedAttribute{
											Description:         "Outputs defines output bindings.",
											MarkdownDescription: "Outputs defines output bindings.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"match": schema.MapAttribute{
														Description:         "Match defines the matching statement.",
														MarkdownDescription: "Match defines the matching statement.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name the name of the binding.",
														MarkdownDescription: "Name the name of the binding.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
														},
													},

													"value": schema.MapAttribute{
														Description:         "Value value of the binding.",
														MarkdownDescription: "Value value of the binding.",
														ElementType:         types.StringType,
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

										"skip_log_output": schema.BoolAttribute{
											Description:         "SkipLogOutput removes the output from the command. Useful for sensitive logs or to reduce noise.",
											MarkdownDescription: "SkipLogOutput removes the output from the command. Useful for sensitive logs or to reduce noise.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"timeout": schema.StringAttribute{
											Description:         "Timeout for the operation. Overrides the global timeout set in the Configuration.",
											MarkdownDescription: "Timeout for the operation. Overrides the global timeout set in the Configuration.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"sleep": schema.SingleNestedAttribute{
									Description:         "Sleep defines zzzz.",
									MarkdownDescription: "Sleep defines zzzz.",
									Attributes: map[string]schema.Attribute{
										"duration": schema.StringAttribute{
											Description:         "Duration is the delay used for sleeping.",
											MarkdownDescription: "Duration is the delay used for sleeping.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"wait": schema.SingleNestedAttribute{
									Description:         "Wait determines the resource wait collector to execute.",
									MarkdownDescription: "Wait determines the resource wait collector to execute.",
									Attributes: map[string]schema.Attribute{
										"api_version": schema.StringAttribute{
											Description:         "API version of the referent.",
											MarkdownDescription: "API version of the referent.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"cluster": schema.StringAttribute{
											Description:         "Cluster defines the target cluster where the wait operation will be performed (default cluster will be used if not specified).",
											MarkdownDescription: "Cluster defines the target cluster where the wait operation will be performed (default cluster will be used if not specified).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"clusters": schema.SingleNestedAttribute{
											Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
											MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
											Attributes: map[string]schema.Attribute{
												"context": schema.StringAttribute{
													Description:         "Context is the name of the context to use.",
													MarkdownDescription: "Context is the name of the context to use.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"kubeconfig": schema.StringAttribute{
													Description:         "Kubeconfig is the path to the referenced file.",
													MarkdownDescription: "Kubeconfig is the path to the referenced file.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"for": schema.SingleNestedAttribute{
											Description:         "For specifies the condition to wait for.",
											MarkdownDescription: "For specifies the condition to wait for.",
											Attributes: map[string]schema.Attribute{
												"condition": schema.SingleNestedAttribute{
													Description:         "Condition specifies the condition to wait for.",
													MarkdownDescription: "Condition specifies the condition to wait for.",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name defines the specific condition to wait for, e.g., 'Available', 'Ready'.",
															MarkdownDescription: "Name defines the specific condition to wait for, e.g., 'Available', 'Ready'.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"value": schema.StringAttribute{
															Description:         "Value defines the specific condition status to wait for, e.g., 'True', 'False'.",
															MarkdownDescription: "Value defines the specific condition status to wait for, e.g., 'True', 'False'.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"deletion": schema.MapAttribute{
													Description:         "Deletion specifies parameters for waiting on a resource's deletion.",
													MarkdownDescription: "Deletion specifies parameters for waiting on a resource's deletion.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"json_path": schema.SingleNestedAttribute{
													Description:         "JsonPath specifies the json path condition to wait for.",
													MarkdownDescription: "JsonPath specifies the json path condition to wait for.",
													Attributes: map[string]schema.Attribute{
														"path": schema.StringAttribute{
															Description:         "Path defines the json path to wait for, e.g. '{.status.phase}'.",
															MarkdownDescription: "Path defines the json path to wait for, e.g. '{.status.phase}'.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"value": schema.StringAttribute{
															Description:         "Value defines the expected value to wait for, e.g., 'Running'.",
															MarkdownDescription: "Value defines the expected value to wait for, e.g., 'Running'.",
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
											Required: true,
											Optional: false,
											Computed: false,
										},

										"format": schema.StringAttribute{
											Description:         "Format determines the output format (json or yaml).",
											MarkdownDescription: "Format determines the output format (json or yaml).",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^(?:json|yaml|\(.+\))$`), ""),
											},
										},

										"kind": schema.StringAttribute{
											Description:         "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
											MarkdownDescription: "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
											Required:            true,
											Optional:            false,
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

										"selector": schema.StringAttribute{
											Description:         "Selector defines labels selector.",
											MarkdownDescription: "Selector defines labels selector.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"timeout": schema.StringAttribute{
											Description:         "Timeout for the operation. Specifies how long to wait for the condition to be met before timing out.",
											MarkdownDescription: "Timeout for the operation. Specifies how long to wait for the condition to be met before timing out.",
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

					"cluster": schema.StringAttribute{
						Description:         "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
						MarkdownDescription: "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"clusters": schema.SingleNestedAttribute{
						Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
						MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
						Attributes: map[string]schema.Attribute{
							"context": schema.StringAttribute{
								Description:         "Context is the name of the context to use.",
								MarkdownDescription: "Context is the name of the context to use.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kubeconfig": schema.StringAttribute{
								Description:         "Kubeconfig is the path to the referenced file.",
								MarkdownDescription: "Kubeconfig is the path to the referenced file.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"concurrent": schema.BoolAttribute{
						Description:         "Concurrent determines whether the test should run concurrently with other tests.",
						MarkdownDescription: "Concurrent determines whether the test should run concurrently with other tests.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"delay_before_cleanup": schema.StringAttribute{
						Description:         "DelayBeforeCleanup adds a delay between the time a test ends and the time cleanup starts.",
						MarkdownDescription: "DelayBeforeCleanup adds a delay between the time a test ends and the time cleanup starts.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"deletion_propagation_policy": schema.StringAttribute{
						Description:         "DeletionPropagationPolicy decides if a deletion will propagate to the dependents ofthe object, and how the garbage collector will handle the propagation.Overrides the deletion propagation policy set in the Configuration.",
						MarkdownDescription: "DeletionPropagationPolicy decides if a deletion will propagate to the dependents ofthe object, and how the garbage collector will handle the propagation.Overrides the deletion propagation policy set in the Configuration.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Orphan", "Background", "Foreground"),
						},
					},

					"description": schema.StringAttribute{
						Description:         "Description contains a description of the test.",
						MarkdownDescription: "Description contains a description of the test.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"force_termination_grace_period": schema.StringAttribute{
						Description:         "ForceTerminationGracePeriod forces the termination grace period on pods, statefulsets, daemonsets and deployments.",
						MarkdownDescription: "ForceTerminationGracePeriod forces the termination grace period on pods, statefulsets, daemonsets and deployments.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"namespace": schema.StringAttribute{
						Description:         "Namespace determines whether the test should run in a random ephemeral namespace or not.",
						MarkdownDescription: "Namespace determines whether the test should run in a random ephemeral namespace or not.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"namespace_template": schema.MapAttribute{
						Description:         "NamespaceTemplate defines a template to create the test namespace.",
						MarkdownDescription: "NamespaceTemplate defines a template to create the test namespace.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"skip": schema.BoolAttribute{
						Description:         "Skip determines whether the test should skipped.",
						MarkdownDescription: "Skip determines whether the test should skipped.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"skip_delete": schema.BoolAttribute{
						Description:         "SkipDelete determines whether the resources created by the test should be deleted after the test is executed.",
						MarkdownDescription: "SkipDelete determines whether the resources created by the test should be deleted after the test is executed.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"steps": schema.ListNestedAttribute{
						Description:         "Steps defining the test.",
						MarkdownDescription: "Steps defining the test.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"bindings": schema.ListNestedAttribute{
									Description:         "Bindings defines additional binding key/values.",
									MarkdownDescription: "Bindings defines additional binding key/values.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name the name of the binding.",
												MarkdownDescription: "Name the name of the binding.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
												},
											},

											"value": schema.MapAttribute{
												Description:         "Value value of the binding.",
												MarkdownDescription: "Value value of the binding.",
												ElementType:         types.StringType,
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

								"catch": schema.ListNestedAttribute{
									Description:         "Catch defines what the step will execute when an error happens.",
									MarkdownDescription: "Catch defines what the step will execute when an error happens.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"command": schema.SingleNestedAttribute{
												Description:         "Command defines a command to run.",
												MarkdownDescription: "Command defines a command to run.",
												Attributes: map[string]schema.Attribute{
													"args": schema.ListAttribute{
														Description:         "Args is the command arguments.",
														MarkdownDescription: "Args is the command arguments.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"bindings": schema.ListNestedAttribute{
														Description:         "Bindings defines additional binding key/values.",
														MarkdownDescription: "Bindings defines additional binding key/values.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name the name of the binding.",
																	MarkdownDescription: "Name the name of the binding.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
																	},
																},

																"value": schema.MapAttribute{
																	Description:         "Value value of the binding.",
																	MarkdownDescription: "Value value of the binding.",
																	ElementType:         types.StringType,
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

													"check": schema.MapAttribute{
														Description:         "Check is an assertion tree to validate the operation outcome.",
														MarkdownDescription: "Check is an assertion tree to validate the operation outcome.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"cluster": schema.StringAttribute{
														Description:         "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														MarkdownDescription: "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"clusters": schema.SingleNestedAttribute{
														Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
														MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Description:         "Context is the name of the context to use.",
																MarkdownDescription: "Context is the name of the context to use.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubeconfig": schema.StringAttribute{
																Description:         "Kubeconfig is the path to the referenced file.",
																MarkdownDescription: "Kubeconfig is the path to the referenced file.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"entrypoint": schema.StringAttribute{
														Description:         "Entrypoint is the command entry point to run.",
														MarkdownDescription: "Entrypoint is the command entry point to run.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"env": schema.ListNestedAttribute{
														Description:         "Env defines additional environment variables.",
														MarkdownDescription: "Env defines additional environment variables.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name the name of the binding.",
																	MarkdownDescription: "Name the name of the binding.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
																	},
																},

																"value": schema.MapAttribute{
																	Description:         "Value value of the binding.",
																	MarkdownDescription: "Value value of the binding.",
																	ElementType:         types.StringType,
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

													"outputs": schema.ListNestedAttribute{
														Description:         "Outputs defines output bindings.",
														MarkdownDescription: "Outputs defines output bindings.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"match": schema.MapAttribute{
																	Description:         "Match defines the matching statement.",
																	MarkdownDescription: "Match defines the matching statement.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name the name of the binding.",
																	MarkdownDescription: "Name the name of the binding.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
																	},
																},

																"value": schema.MapAttribute{
																	Description:         "Value value of the binding.",
																	MarkdownDescription: "Value value of the binding.",
																	ElementType:         types.StringType,
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

													"skip_log_output": schema.BoolAttribute{
														Description:         "SkipLogOutput removes the output from the command. Useful for sensitive logs or to reduce noise.",
														MarkdownDescription: "SkipLogOutput removes the output from the command. Useful for sensitive logs or to reduce noise.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.StringAttribute{
														Description:         "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														MarkdownDescription: "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"delete": schema.SingleNestedAttribute{
												Description:         "Delete represents a deletion operation.",
												MarkdownDescription: "Delete represents a deletion operation.",
												Attributes: map[string]schema.Attribute{
													"bindings": schema.ListNestedAttribute{
														Description:         "Bindings defines additional binding key/values.",
														MarkdownDescription: "Bindings defines additional binding key/values.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name the name of the binding.",
																	MarkdownDescription: "Name the name of the binding.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
																	},
																},

																"value": schema.MapAttribute{
																	Description:         "Value value of the binding.",
																	MarkdownDescription: "Value value of the binding.",
																	ElementType:         types.StringType,
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

													"cluster": schema.StringAttribute{
														Description:         "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														MarkdownDescription: "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"clusters": schema.SingleNestedAttribute{
														Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
														MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Description:         "Context is the name of the context to use.",
																MarkdownDescription: "Context is the name of the context to use.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubeconfig": schema.StringAttribute{
																Description:         "Kubeconfig is the path to the referenced file.",
																MarkdownDescription: "Kubeconfig is the path to the referenced file.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"deletion_propagation_policy": schema.StringAttribute{
														Description:         "DeletionPropagationPolicy decides if a deletion will propagate to the dependents ofthe object, and how the garbage collector will handle the propagation.Overrides the deletion propagation policy set in the Configuration, the Test and the TestStep.",
														MarkdownDescription: "DeletionPropagationPolicy decides if a deletion will propagate to the dependents ofthe object, and how the garbage collector will handle the propagation.Overrides the deletion propagation policy set in the Configuration, the Test and the TestStep.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("Orphan", "Background", "Foreground"),
														},
													},

													"expect": schema.ListNestedAttribute{
														Description:         "Expect defines a list of matched checks to validate the operation outcome.",
														MarkdownDescription: "Expect defines a list of matched checks to validate the operation outcome.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"check": schema.MapAttribute{
																	Description:         "Check defines the verification statement.",
																	MarkdownDescription: "Check defines the verification statement.",
																	ElementType:         types.StringType,
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"match": schema.MapAttribute{
																	Description:         "Match defines the matching statement.",
																	MarkdownDescription: "Match defines the matching statement.",
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

													"file": schema.StringAttribute{
														Description:         "File is the path to the referenced file. This can be a direct path to a fileor an expression that matches multiple files, such as 'manifest/*.yaml' for all YAMLfiles within the 'manifest' directory.",
														MarkdownDescription: "File is the path to the referenced file. This can be a direct path to a fileor an expression that matches multiple files, such as 'manifest/*.yaml' for all YAMLfiles within the 'manifest' directory.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ref": schema.SingleNestedAttribute{
														Description:         "Ref determines objects to be deleted.",
														MarkdownDescription: "Ref determines objects to be deleted.",
														Attributes: map[string]schema.Attribute{
															"api_version": schema.StringAttribute{
																Description:         "API version of the referent.",
																MarkdownDescription: "API version of the referent.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"kind": schema.StringAttribute{
																Description:         "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
																MarkdownDescription: "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"labels": schema.MapAttribute{
																Description:         "Label selector to match objects to delete",
																MarkdownDescription: "Label selector to match objects to delete",
																ElementType:         types.StringType,
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"template": schema.BoolAttribute{
														Description:         "Template determines whether resources should be considered for templating.",
														MarkdownDescription: "Template determines whether resources should be considered for templating.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.StringAttribute{
														Description:         "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														MarkdownDescription: "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"describe": schema.SingleNestedAttribute{
												Description:         "Describe determines the resource describe collector to execute.",
												MarkdownDescription: "Describe determines the resource describe collector to execute.",
												Attributes: map[string]schema.Attribute{
													"api_version": schema.StringAttribute{
														Description:         "API version of the referent.",
														MarkdownDescription: "API version of the referent.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"cluster": schema.StringAttribute{
														Description:         "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														MarkdownDescription: "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"clusters": schema.SingleNestedAttribute{
														Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
														MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Description:         "Context is the name of the context to use.",
																MarkdownDescription: "Context is the name of the context to use.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubeconfig": schema.StringAttribute{
																Description:         "Kubeconfig is the path to the referenced file.",
																MarkdownDescription: "Kubeconfig is the path to the referenced file.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"kind": schema.StringAttribute{
														Description:         "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
														MarkdownDescription: "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
														Required:            true,
														Optional:            false,
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

													"selector": schema.StringAttribute{
														Description:         "Selector defines labels selector.",
														MarkdownDescription: "Selector defines labels selector.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"show_events": schema.BoolAttribute{
														Description:         "Show Events indicates whether to include related events.",
														MarkdownDescription: "Show Events indicates whether to include related events.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.StringAttribute{
														Description:         "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														MarkdownDescription: "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"description": schema.StringAttribute{
												Description:         "Description contains a description of the operation.",
												MarkdownDescription: "Description contains a description of the operation.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"events": schema.SingleNestedAttribute{
												Description:         "Events determines the events collector to execute.",
												MarkdownDescription: "Events determines the events collector to execute.",
												Attributes: map[string]schema.Attribute{
													"cluster": schema.StringAttribute{
														Description:         "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														MarkdownDescription: "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"clusters": schema.SingleNestedAttribute{
														Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
														MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Description:         "Context is the name of the context to use.",
																MarkdownDescription: "Context is the name of the context to use.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubeconfig": schema.StringAttribute{
																Description:         "Kubeconfig is the path to the referenced file.",
																MarkdownDescription: "Kubeconfig is the path to the referenced file.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"format": schema.StringAttribute{
														Description:         "Format determines the output format (json or yaml).",
														MarkdownDescription: "Format determines the output format (json or yaml).",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.RegexMatches(regexp.MustCompile(`^(?:json|yaml|\(.+\))$`), ""),
														},
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

													"selector": schema.StringAttribute{
														Description:         "Selector defines labels selector.",
														MarkdownDescription: "Selector defines labels selector.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.StringAttribute{
														Description:         "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														MarkdownDescription: "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"get": schema.SingleNestedAttribute{
												Description:         "Get determines the resource get collector to execute.",
												MarkdownDescription: "Get determines the resource get collector to execute.",
												Attributes: map[string]schema.Attribute{
													"api_version": schema.StringAttribute{
														Description:         "API version of the referent.",
														MarkdownDescription: "API version of the referent.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"cluster": schema.StringAttribute{
														Description:         "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														MarkdownDescription: "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"clusters": schema.SingleNestedAttribute{
														Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
														MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Description:         "Context is the name of the context to use.",
																MarkdownDescription: "Context is the name of the context to use.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubeconfig": schema.StringAttribute{
																Description:         "Kubeconfig is the path to the referenced file.",
																MarkdownDescription: "Kubeconfig is the path to the referenced file.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"format": schema.StringAttribute{
														Description:         "Format determines the output format (json or yaml).",
														MarkdownDescription: "Format determines the output format (json or yaml).",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.RegexMatches(regexp.MustCompile(`^(?:json|yaml|\(.+\))$`), ""),
														},
													},

													"kind": schema.StringAttribute{
														Description:         "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
														MarkdownDescription: "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
														Required:            true,
														Optional:            false,
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

													"selector": schema.StringAttribute{
														Description:         "Selector defines labels selector.",
														MarkdownDescription: "Selector defines labels selector.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.StringAttribute{
														Description:         "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														MarkdownDescription: "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"pod_logs": schema.SingleNestedAttribute{
												Description:         "PodLogs determines the pod logs collector to execute.",
												MarkdownDescription: "PodLogs determines the pod logs collector to execute.",
												Attributes: map[string]schema.Attribute{
													"cluster": schema.StringAttribute{
														Description:         "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														MarkdownDescription: "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"clusters": schema.SingleNestedAttribute{
														Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
														MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Description:         "Context is the name of the context to use.",
																MarkdownDescription: "Context is the name of the context to use.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubeconfig": schema.StringAttribute{
																Description:         "Kubeconfig is the path to the referenced file.",
																MarkdownDescription: "Kubeconfig is the path to the referenced file.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"container": schema.StringAttribute{
														Description:         "Container in pod to get logs from else --all-containers is used.",
														MarkdownDescription: "Container in pod to get logs from else --all-containers is used.",
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

													"selector": schema.StringAttribute{
														Description:         "Selector defines labels selector.",
														MarkdownDescription: "Selector defines labels selector.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"tail": schema.Int64Attribute{
														Description:         "Tail is the number of last lines to collect from pods. If omitted or zero,then the default is 10 if you use a selector, or -1 (all) if you use a pod name.This matches default behavior of 'kubectl logs'.",
														MarkdownDescription: "Tail is the number of last lines to collect from pods. If omitted or zero,then the default is 10 if you use a selector, or -1 (all) if you use a pod name.This matches default behavior of 'kubectl logs'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.StringAttribute{
														Description:         "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														MarkdownDescription: "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"script": schema.SingleNestedAttribute{
												Description:         "Script defines a script to run.",
												MarkdownDescription: "Script defines a script to run.",
												Attributes: map[string]schema.Attribute{
													"bindings": schema.ListNestedAttribute{
														Description:         "Bindings defines additional binding key/values.",
														MarkdownDescription: "Bindings defines additional binding key/values.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name the name of the binding.",
																	MarkdownDescription: "Name the name of the binding.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
																	},
																},

																"value": schema.MapAttribute{
																	Description:         "Value value of the binding.",
																	MarkdownDescription: "Value value of the binding.",
																	ElementType:         types.StringType,
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

													"check": schema.MapAttribute{
														Description:         "Check is an assertion tree to validate the operation outcome.",
														MarkdownDescription: "Check is an assertion tree to validate the operation outcome.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"cluster": schema.StringAttribute{
														Description:         "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														MarkdownDescription: "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"clusters": schema.SingleNestedAttribute{
														Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
														MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Description:         "Context is the name of the context to use.",
																MarkdownDescription: "Context is the name of the context to use.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubeconfig": schema.StringAttribute{
																Description:         "Kubeconfig is the path to the referenced file.",
																MarkdownDescription: "Kubeconfig is the path to the referenced file.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"content": schema.StringAttribute{
														Description:         "Content defines a shell script (run with 'sh -c ...').",
														MarkdownDescription: "Content defines a shell script (run with 'sh -c ...').",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"env": schema.ListNestedAttribute{
														Description:         "Env defines additional environment variables.",
														MarkdownDescription: "Env defines additional environment variables.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name the name of the binding.",
																	MarkdownDescription: "Name the name of the binding.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
																	},
																},

																"value": schema.MapAttribute{
																	Description:         "Value value of the binding.",
																	MarkdownDescription: "Value value of the binding.",
																	ElementType:         types.StringType,
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

													"outputs": schema.ListNestedAttribute{
														Description:         "Outputs defines output bindings.",
														MarkdownDescription: "Outputs defines output bindings.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"match": schema.MapAttribute{
																	Description:         "Match defines the matching statement.",
																	MarkdownDescription: "Match defines the matching statement.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name the name of the binding.",
																	MarkdownDescription: "Name the name of the binding.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
																	},
																},

																"value": schema.MapAttribute{
																	Description:         "Value value of the binding.",
																	MarkdownDescription: "Value value of the binding.",
																	ElementType:         types.StringType,
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

													"skip_log_output": schema.BoolAttribute{
														Description:         "SkipLogOutput removes the output from the command. Useful for sensitive logs or to reduce noise.",
														MarkdownDescription: "SkipLogOutput removes the output from the command. Useful for sensitive logs or to reduce noise.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.StringAttribute{
														Description:         "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														MarkdownDescription: "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"sleep": schema.SingleNestedAttribute{
												Description:         "Sleep defines zzzz.",
												MarkdownDescription: "Sleep defines zzzz.",
												Attributes: map[string]schema.Attribute{
													"duration": schema.StringAttribute{
														Description:         "Duration is the delay used for sleeping.",
														MarkdownDescription: "Duration is the delay used for sleeping.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"wait": schema.SingleNestedAttribute{
												Description:         "Wait determines the resource wait collector to execute.",
												MarkdownDescription: "Wait determines the resource wait collector to execute.",
												Attributes: map[string]schema.Attribute{
													"api_version": schema.StringAttribute{
														Description:         "API version of the referent.",
														MarkdownDescription: "API version of the referent.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"cluster": schema.StringAttribute{
														Description:         "Cluster defines the target cluster where the wait operation will be performed (default cluster will be used if not specified).",
														MarkdownDescription: "Cluster defines the target cluster where the wait operation will be performed (default cluster will be used if not specified).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"clusters": schema.SingleNestedAttribute{
														Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
														MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Description:         "Context is the name of the context to use.",
																MarkdownDescription: "Context is the name of the context to use.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubeconfig": schema.StringAttribute{
																Description:         "Kubeconfig is the path to the referenced file.",
																MarkdownDescription: "Kubeconfig is the path to the referenced file.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"for": schema.SingleNestedAttribute{
														Description:         "For specifies the condition to wait for.",
														MarkdownDescription: "For specifies the condition to wait for.",
														Attributes: map[string]schema.Attribute{
															"condition": schema.SingleNestedAttribute{
																Description:         "Condition specifies the condition to wait for.",
																MarkdownDescription: "Condition specifies the condition to wait for.",
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "Name defines the specific condition to wait for, e.g., 'Available', 'Ready'.",
																		MarkdownDescription: "Name defines the specific condition to wait for, e.g., 'Available', 'Ready'.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
																		Description:         "Value defines the specific condition status to wait for, e.g., 'True', 'False'.",
																		MarkdownDescription: "Value defines the specific condition status to wait for, e.g., 'True', 'False'.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"deletion": schema.MapAttribute{
																Description:         "Deletion specifies parameters for waiting on a resource's deletion.",
																MarkdownDescription: "Deletion specifies parameters for waiting on a resource's deletion.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"json_path": schema.SingleNestedAttribute{
																Description:         "JsonPath specifies the json path condition to wait for.",
																MarkdownDescription: "JsonPath specifies the json path condition to wait for.",
																Attributes: map[string]schema.Attribute{
																	"path": schema.StringAttribute{
																		Description:         "Path defines the json path to wait for, e.g. '{.status.phase}'.",
																		MarkdownDescription: "Path defines the json path to wait for, e.g. '{.status.phase}'.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
																		Description:         "Value defines the expected value to wait for, e.g., 'Running'.",
																		MarkdownDescription: "Value defines the expected value to wait for, e.g., 'Running'.",
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
														Required: true,
														Optional: false,
														Computed: false,
													},

													"format": schema.StringAttribute{
														Description:         "Format determines the output format (json or yaml).",
														MarkdownDescription: "Format determines the output format (json or yaml).",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.RegexMatches(regexp.MustCompile(`^(?:json|yaml|\(.+\))$`), ""),
														},
													},

													"kind": schema.StringAttribute{
														Description:         "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
														MarkdownDescription: "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
														Required:            true,
														Optional:            false,
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

													"selector": schema.StringAttribute{
														Description:         "Selector defines labels selector.",
														MarkdownDescription: "Selector defines labels selector.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.StringAttribute{
														Description:         "Timeout for the operation. Specifies how long to wait for the condition to be met before timing out.",
														MarkdownDescription: "Timeout for the operation. Specifies how long to wait for the condition to be met before timing out.",
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

								"cleanup": schema.ListNestedAttribute{
									Description:         "Cleanup defines what will be executed after the test is terminated.",
									MarkdownDescription: "Cleanup defines what will be executed after the test is terminated.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"command": schema.SingleNestedAttribute{
												Description:         "Command defines a command to run.",
												MarkdownDescription: "Command defines a command to run.",
												Attributes: map[string]schema.Attribute{
													"args": schema.ListAttribute{
														Description:         "Args is the command arguments.",
														MarkdownDescription: "Args is the command arguments.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"bindings": schema.ListNestedAttribute{
														Description:         "Bindings defines additional binding key/values.",
														MarkdownDescription: "Bindings defines additional binding key/values.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name the name of the binding.",
																	MarkdownDescription: "Name the name of the binding.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
																	},
																},

																"value": schema.MapAttribute{
																	Description:         "Value value of the binding.",
																	MarkdownDescription: "Value value of the binding.",
																	ElementType:         types.StringType,
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

													"check": schema.MapAttribute{
														Description:         "Check is an assertion tree to validate the operation outcome.",
														MarkdownDescription: "Check is an assertion tree to validate the operation outcome.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"cluster": schema.StringAttribute{
														Description:         "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														MarkdownDescription: "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"clusters": schema.SingleNestedAttribute{
														Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
														MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Description:         "Context is the name of the context to use.",
																MarkdownDescription: "Context is the name of the context to use.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubeconfig": schema.StringAttribute{
																Description:         "Kubeconfig is the path to the referenced file.",
																MarkdownDescription: "Kubeconfig is the path to the referenced file.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"entrypoint": schema.StringAttribute{
														Description:         "Entrypoint is the command entry point to run.",
														MarkdownDescription: "Entrypoint is the command entry point to run.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"env": schema.ListNestedAttribute{
														Description:         "Env defines additional environment variables.",
														MarkdownDescription: "Env defines additional environment variables.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name the name of the binding.",
																	MarkdownDescription: "Name the name of the binding.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
																	},
																},

																"value": schema.MapAttribute{
																	Description:         "Value value of the binding.",
																	MarkdownDescription: "Value value of the binding.",
																	ElementType:         types.StringType,
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

													"outputs": schema.ListNestedAttribute{
														Description:         "Outputs defines output bindings.",
														MarkdownDescription: "Outputs defines output bindings.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"match": schema.MapAttribute{
																	Description:         "Match defines the matching statement.",
																	MarkdownDescription: "Match defines the matching statement.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name the name of the binding.",
																	MarkdownDescription: "Name the name of the binding.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
																	},
																},

																"value": schema.MapAttribute{
																	Description:         "Value value of the binding.",
																	MarkdownDescription: "Value value of the binding.",
																	ElementType:         types.StringType,
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

													"skip_log_output": schema.BoolAttribute{
														Description:         "SkipLogOutput removes the output from the command. Useful for sensitive logs or to reduce noise.",
														MarkdownDescription: "SkipLogOutput removes the output from the command. Useful for sensitive logs or to reduce noise.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.StringAttribute{
														Description:         "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														MarkdownDescription: "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"delete": schema.SingleNestedAttribute{
												Description:         "Delete represents a deletion operation.",
												MarkdownDescription: "Delete represents a deletion operation.",
												Attributes: map[string]schema.Attribute{
													"bindings": schema.ListNestedAttribute{
														Description:         "Bindings defines additional binding key/values.",
														MarkdownDescription: "Bindings defines additional binding key/values.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name the name of the binding.",
																	MarkdownDescription: "Name the name of the binding.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
																	},
																},

																"value": schema.MapAttribute{
																	Description:         "Value value of the binding.",
																	MarkdownDescription: "Value value of the binding.",
																	ElementType:         types.StringType,
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

													"cluster": schema.StringAttribute{
														Description:         "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														MarkdownDescription: "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"clusters": schema.SingleNestedAttribute{
														Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
														MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Description:         "Context is the name of the context to use.",
																MarkdownDescription: "Context is the name of the context to use.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubeconfig": schema.StringAttribute{
																Description:         "Kubeconfig is the path to the referenced file.",
																MarkdownDescription: "Kubeconfig is the path to the referenced file.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"deletion_propagation_policy": schema.StringAttribute{
														Description:         "DeletionPropagationPolicy decides if a deletion will propagate to the dependents ofthe object, and how the garbage collector will handle the propagation.Overrides the deletion propagation policy set in the Configuration, the Test and the TestStep.",
														MarkdownDescription: "DeletionPropagationPolicy decides if a deletion will propagate to the dependents ofthe object, and how the garbage collector will handle the propagation.Overrides the deletion propagation policy set in the Configuration, the Test and the TestStep.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("Orphan", "Background", "Foreground"),
														},
													},

													"expect": schema.ListNestedAttribute{
														Description:         "Expect defines a list of matched checks to validate the operation outcome.",
														MarkdownDescription: "Expect defines a list of matched checks to validate the operation outcome.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"check": schema.MapAttribute{
																	Description:         "Check defines the verification statement.",
																	MarkdownDescription: "Check defines the verification statement.",
																	ElementType:         types.StringType,
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"match": schema.MapAttribute{
																	Description:         "Match defines the matching statement.",
																	MarkdownDescription: "Match defines the matching statement.",
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

													"file": schema.StringAttribute{
														Description:         "File is the path to the referenced file. This can be a direct path to a fileor an expression that matches multiple files, such as 'manifest/*.yaml' for all YAMLfiles within the 'manifest' directory.",
														MarkdownDescription: "File is the path to the referenced file. This can be a direct path to a fileor an expression that matches multiple files, such as 'manifest/*.yaml' for all YAMLfiles within the 'manifest' directory.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ref": schema.SingleNestedAttribute{
														Description:         "Ref determines objects to be deleted.",
														MarkdownDescription: "Ref determines objects to be deleted.",
														Attributes: map[string]schema.Attribute{
															"api_version": schema.StringAttribute{
																Description:         "API version of the referent.",
																MarkdownDescription: "API version of the referent.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"kind": schema.StringAttribute{
																Description:         "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
																MarkdownDescription: "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"labels": schema.MapAttribute{
																Description:         "Label selector to match objects to delete",
																MarkdownDescription: "Label selector to match objects to delete",
																ElementType:         types.StringType,
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"template": schema.BoolAttribute{
														Description:         "Template determines whether resources should be considered for templating.",
														MarkdownDescription: "Template determines whether resources should be considered for templating.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.StringAttribute{
														Description:         "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														MarkdownDescription: "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"describe": schema.SingleNestedAttribute{
												Description:         "Describe determines the resource describe collector to execute.",
												MarkdownDescription: "Describe determines the resource describe collector to execute.",
												Attributes: map[string]schema.Attribute{
													"api_version": schema.StringAttribute{
														Description:         "API version of the referent.",
														MarkdownDescription: "API version of the referent.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"cluster": schema.StringAttribute{
														Description:         "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														MarkdownDescription: "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"clusters": schema.SingleNestedAttribute{
														Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
														MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Description:         "Context is the name of the context to use.",
																MarkdownDescription: "Context is the name of the context to use.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubeconfig": schema.StringAttribute{
																Description:         "Kubeconfig is the path to the referenced file.",
																MarkdownDescription: "Kubeconfig is the path to the referenced file.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"kind": schema.StringAttribute{
														Description:         "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
														MarkdownDescription: "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
														Required:            true,
														Optional:            false,
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

													"selector": schema.StringAttribute{
														Description:         "Selector defines labels selector.",
														MarkdownDescription: "Selector defines labels selector.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"show_events": schema.BoolAttribute{
														Description:         "Show Events indicates whether to include related events.",
														MarkdownDescription: "Show Events indicates whether to include related events.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.StringAttribute{
														Description:         "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														MarkdownDescription: "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"description": schema.StringAttribute{
												Description:         "Description contains a description of the operation.",
												MarkdownDescription: "Description contains a description of the operation.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"events": schema.SingleNestedAttribute{
												Description:         "Events determines the events collector to execute.",
												MarkdownDescription: "Events determines the events collector to execute.",
												Attributes: map[string]schema.Attribute{
													"cluster": schema.StringAttribute{
														Description:         "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														MarkdownDescription: "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"clusters": schema.SingleNestedAttribute{
														Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
														MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Description:         "Context is the name of the context to use.",
																MarkdownDescription: "Context is the name of the context to use.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubeconfig": schema.StringAttribute{
																Description:         "Kubeconfig is the path to the referenced file.",
																MarkdownDescription: "Kubeconfig is the path to the referenced file.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"format": schema.StringAttribute{
														Description:         "Format determines the output format (json or yaml).",
														MarkdownDescription: "Format determines the output format (json or yaml).",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.RegexMatches(regexp.MustCompile(`^(?:json|yaml|\(.+\))$`), ""),
														},
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

													"selector": schema.StringAttribute{
														Description:         "Selector defines labels selector.",
														MarkdownDescription: "Selector defines labels selector.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.StringAttribute{
														Description:         "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														MarkdownDescription: "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"get": schema.SingleNestedAttribute{
												Description:         "Get determines the resource get collector to execute.",
												MarkdownDescription: "Get determines the resource get collector to execute.",
												Attributes: map[string]schema.Attribute{
													"api_version": schema.StringAttribute{
														Description:         "API version of the referent.",
														MarkdownDescription: "API version of the referent.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"cluster": schema.StringAttribute{
														Description:         "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														MarkdownDescription: "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"clusters": schema.SingleNestedAttribute{
														Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
														MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Description:         "Context is the name of the context to use.",
																MarkdownDescription: "Context is the name of the context to use.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubeconfig": schema.StringAttribute{
																Description:         "Kubeconfig is the path to the referenced file.",
																MarkdownDescription: "Kubeconfig is the path to the referenced file.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"format": schema.StringAttribute{
														Description:         "Format determines the output format (json or yaml).",
														MarkdownDescription: "Format determines the output format (json or yaml).",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.RegexMatches(regexp.MustCompile(`^(?:json|yaml|\(.+\))$`), ""),
														},
													},

													"kind": schema.StringAttribute{
														Description:         "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
														MarkdownDescription: "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
														Required:            true,
														Optional:            false,
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

													"selector": schema.StringAttribute{
														Description:         "Selector defines labels selector.",
														MarkdownDescription: "Selector defines labels selector.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.StringAttribute{
														Description:         "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														MarkdownDescription: "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"pod_logs": schema.SingleNestedAttribute{
												Description:         "PodLogs determines the pod logs collector to execute.",
												MarkdownDescription: "PodLogs determines the pod logs collector to execute.",
												Attributes: map[string]schema.Attribute{
													"cluster": schema.StringAttribute{
														Description:         "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														MarkdownDescription: "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"clusters": schema.SingleNestedAttribute{
														Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
														MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Description:         "Context is the name of the context to use.",
																MarkdownDescription: "Context is the name of the context to use.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubeconfig": schema.StringAttribute{
																Description:         "Kubeconfig is the path to the referenced file.",
																MarkdownDescription: "Kubeconfig is the path to the referenced file.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"container": schema.StringAttribute{
														Description:         "Container in pod to get logs from else --all-containers is used.",
														MarkdownDescription: "Container in pod to get logs from else --all-containers is used.",
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

													"selector": schema.StringAttribute{
														Description:         "Selector defines labels selector.",
														MarkdownDescription: "Selector defines labels selector.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"tail": schema.Int64Attribute{
														Description:         "Tail is the number of last lines to collect from pods. If omitted or zero,then the default is 10 if you use a selector, or -1 (all) if you use a pod name.This matches default behavior of 'kubectl logs'.",
														MarkdownDescription: "Tail is the number of last lines to collect from pods. If omitted or zero,then the default is 10 if you use a selector, or -1 (all) if you use a pod name.This matches default behavior of 'kubectl logs'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.StringAttribute{
														Description:         "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														MarkdownDescription: "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"script": schema.SingleNestedAttribute{
												Description:         "Script defines a script to run.",
												MarkdownDescription: "Script defines a script to run.",
												Attributes: map[string]schema.Attribute{
													"bindings": schema.ListNestedAttribute{
														Description:         "Bindings defines additional binding key/values.",
														MarkdownDescription: "Bindings defines additional binding key/values.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name the name of the binding.",
																	MarkdownDescription: "Name the name of the binding.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
																	},
																},

																"value": schema.MapAttribute{
																	Description:         "Value value of the binding.",
																	MarkdownDescription: "Value value of the binding.",
																	ElementType:         types.StringType,
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

													"check": schema.MapAttribute{
														Description:         "Check is an assertion tree to validate the operation outcome.",
														MarkdownDescription: "Check is an assertion tree to validate the operation outcome.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"cluster": schema.StringAttribute{
														Description:         "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														MarkdownDescription: "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"clusters": schema.SingleNestedAttribute{
														Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
														MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Description:         "Context is the name of the context to use.",
																MarkdownDescription: "Context is the name of the context to use.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubeconfig": schema.StringAttribute{
																Description:         "Kubeconfig is the path to the referenced file.",
																MarkdownDescription: "Kubeconfig is the path to the referenced file.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"content": schema.StringAttribute{
														Description:         "Content defines a shell script (run with 'sh -c ...').",
														MarkdownDescription: "Content defines a shell script (run with 'sh -c ...').",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"env": schema.ListNestedAttribute{
														Description:         "Env defines additional environment variables.",
														MarkdownDescription: "Env defines additional environment variables.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name the name of the binding.",
																	MarkdownDescription: "Name the name of the binding.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
																	},
																},

																"value": schema.MapAttribute{
																	Description:         "Value value of the binding.",
																	MarkdownDescription: "Value value of the binding.",
																	ElementType:         types.StringType,
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

													"outputs": schema.ListNestedAttribute{
														Description:         "Outputs defines output bindings.",
														MarkdownDescription: "Outputs defines output bindings.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"match": schema.MapAttribute{
																	Description:         "Match defines the matching statement.",
																	MarkdownDescription: "Match defines the matching statement.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name the name of the binding.",
																	MarkdownDescription: "Name the name of the binding.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
																	},
																},

																"value": schema.MapAttribute{
																	Description:         "Value value of the binding.",
																	MarkdownDescription: "Value value of the binding.",
																	ElementType:         types.StringType,
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

													"skip_log_output": schema.BoolAttribute{
														Description:         "SkipLogOutput removes the output from the command. Useful for sensitive logs or to reduce noise.",
														MarkdownDescription: "SkipLogOutput removes the output from the command. Useful for sensitive logs or to reduce noise.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.StringAttribute{
														Description:         "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														MarkdownDescription: "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"sleep": schema.SingleNestedAttribute{
												Description:         "Sleep defines zzzz.",
												MarkdownDescription: "Sleep defines zzzz.",
												Attributes: map[string]schema.Attribute{
													"duration": schema.StringAttribute{
														Description:         "Duration is the delay used for sleeping.",
														MarkdownDescription: "Duration is the delay used for sleeping.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"wait": schema.SingleNestedAttribute{
												Description:         "Wait determines the resource wait collector to execute.",
												MarkdownDescription: "Wait determines the resource wait collector to execute.",
												Attributes: map[string]schema.Attribute{
													"api_version": schema.StringAttribute{
														Description:         "API version of the referent.",
														MarkdownDescription: "API version of the referent.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"cluster": schema.StringAttribute{
														Description:         "Cluster defines the target cluster where the wait operation will be performed (default cluster will be used if not specified).",
														MarkdownDescription: "Cluster defines the target cluster where the wait operation will be performed (default cluster will be used if not specified).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"clusters": schema.SingleNestedAttribute{
														Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
														MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Description:         "Context is the name of the context to use.",
																MarkdownDescription: "Context is the name of the context to use.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubeconfig": schema.StringAttribute{
																Description:         "Kubeconfig is the path to the referenced file.",
																MarkdownDescription: "Kubeconfig is the path to the referenced file.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"for": schema.SingleNestedAttribute{
														Description:         "For specifies the condition to wait for.",
														MarkdownDescription: "For specifies the condition to wait for.",
														Attributes: map[string]schema.Attribute{
															"condition": schema.SingleNestedAttribute{
																Description:         "Condition specifies the condition to wait for.",
																MarkdownDescription: "Condition specifies the condition to wait for.",
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "Name defines the specific condition to wait for, e.g., 'Available', 'Ready'.",
																		MarkdownDescription: "Name defines the specific condition to wait for, e.g., 'Available', 'Ready'.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
																		Description:         "Value defines the specific condition status to wait for, e.g., 'True', 'False'.",
																		MarkdownDescription: "Value defines the specific condition status to wait for, e.g., 'True', 'False'.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"deletion": schema.MapAttribute{
																Description:         "Deletion specifies parameters for waiting on a resource's deletion.",
																MarkdownDescription: "Deletion specifies parameters for waiting on a resource's deletion.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"json_path": schema.SingleNestedAttribute{
																Description:         "JsonPath specifies the json path condition to wait for.",
																MarkdownDescription: "JsonPath specifies the json path condition to wait for.",
																Attributes: map[string]schema.Attribute{
																	"path": schema.StringAttribute{
																		Description:         "Path defines the json path to wait for, e.g. '{.status.phase}'.",
																		MarkdownDescription: "Path defines the json path to wait for, e.g. '{.status.phase}'.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
																		Description:         "Value defines the expected value to wait for, e.g., 'Running'.",
																		MarkdownDescription: "Value defines the expected value to wait for, e.g., 'Running'.",
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
														Required: true,
														Optional: false,
														Computed: false,
													},

													"format": schema.StringAttribute{
														Description:         "Format determines the output format (json or yaml).",
														MarkdownDescription: "Format determines the output format (json or yaml).",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.RegexMatches(regexp.MustCompile(`^(?:json|yaml|\(.+\))$`), ""),
														},
													},

													"kind": schema.StringAttribute{
														Description:         "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
														MarkdownDescription: "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
														Required:            true,
														Optional:            false,
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

													"selector": schema.StringAttribute{
														Description:         "Selector defines labels selector.",
														MarkdownDescription: "Selector defines labels selector.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.StringAttribute{
														Description:         "Timeout for the operation. Specifies how long to wait for the condition to be met before timing out.",
														MarkdownDescription: "Timeout for the operation. Specifies how long to wait for the condition to be met before timing out.",
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

								"cluster": schema.StringAttribute{
									Description:         "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
									MarkdownDescription: "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"clusters": schema.SingleNestedAttribute{
									Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
									MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
									Attributes: map[string]schema.Attribute{
										"context": schema.StringAttribute{
											Description:         "Context is the name of the context to use.",
											MarkdownDescription: "Context is the name of the context to use.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"kubeconfig": schema.StringAttribute{
											Description:         "Kubeconfig is the path to the referenced file.",
											MarkdownDescription: "Kubeconfig is the path to the referenced file.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"deletion_propagation_policy": schema.StringAttribute{
									Description:         "DeletionPropagationPolicy decides if a deletion will propagate to the dependents ofthe object, and how the garbage collector will handle the propagation.Overrides the deletion propagation policy set in both the Configuration and the Test.",
									MarkdownDescription: "DeletionPropagationPolicy decides if a deletion will propagate to the dependents ofthe object, and how the garbage collector will handle the propagation.Overrides the deletion propagation policy set in both the Configuration and the Test.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("Orphan", "Background", "Foreground"),
									},
								},

								"description": schema.StringAttribute{
									Description:         "Description contains a description of the test step.",
									MarkdownDescription: "Description contains a description of the test step.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"finally": schema.ListNestedAttribute{
									Description:         "Finally defines what the step will execute after the step is terminated.",
									MarkdownDescription: "Finally defines what the step will execute after the step is terminated.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"command": schema.SingleNestedAttribute{
												Description:         "Command defines a command to run.",
												MarkdownDescription: "Command defines a command to run.",
												Attributes: map[string]schema.Attribute{
													"args": schema.ListAttribute{
														Description:         "Args is the command arguments.",
														MarkdownDescription: "Args is the command arguments.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"bindings": schema.ListNestedAttribute{
														Description:         "Bindings defines additional binding key/values.",
														MarkdownDescription: "Bindings defines additional binding key/values.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name the name of the binding.",
																	MarkdownDescription: "Name the name of the binding.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
																	},
																},

																"value": schema.MapAttribute{
																	Description:         "Value value of the binding.",
																	MarkdownDescription: "Value value of the binding.",
																	ElementType:         types.StringType,
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

													"check": schema.MapAttribute{
														Description:         "Check is an assertion tree to validate the operation outcome.",
														MarkdownDescription: "Check is an assertion tree to validate the operation outcome.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"cluster": schema.StringAttribute{
														Description:         "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														MarkdownDescription: "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"clusters": schema.SingleNestedAttribute{
														Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
														MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Description:         "Context is the name of the context to use.",
																MarkdownDescription: "Context is the name of the context to use.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubeconfig": schema.StringAttribute{
																Description:         "Kubeconfig is the path to the referenced file.",
																MarkdownDescription: "Kubeconfig is the path to the referenced file.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"entrypoint": schema.StringAttribute{
														Description:         "Entrypoint is the command entry point to run.",
														MarkdownDescription: "Entrypoint is the command entry point to run.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"env": schema.ListNestedAttribute{
														Description:         "Env defines additional environment variables.",
														MarkdownDescription: "Env defines additional environment variables.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name the name of the binding.",
																	MarkdownDescription: "Name the name of the binding.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
																	},
																},

																"value": schema.MapAttribute{
																	Description:         "Value value of the binding.",
																	MarkdownDescription: "Value value of the binding.",
																	ElementType:         types.StringType,
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

													"outputs": schema.ListNestedAttribute{
														Description:         "Outputs defines output bindings.",
														MarkdownDescription: "Outputs defines output bindings.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"match": schema.MapAttribute{
																	Description:         "Match defines the matching statement.",
																	MarkdownDescription: "Match defines the matching statement.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name the name of the binding.",
																	MarkdownDescription: "Name the name of the binding.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
																	},
																},

																"value": schema.MapAttribute{
																	Description:         "Value value of the binding.",
																	MarkdownDescription: "Value value of the binding.",
																	ElementType:         types.StringType,
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

													"skip_log_output": schema.BoolAttribute{
														Description:         "SkipLogOutput removes the output from the command. Useful for sensitive logs or to reduce noise.",
														MarkdownDescription: "SkipLogOutput removes the output from the command. Useful for sensitive logs or to reduce noise.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.StringAttribute{
														Description:         "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														MarkdownDescription: "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"delete": schema.SingleNestedAttribute{
												Description:         "Delete represents a deletion operation.",
												MarkdownDescription: "Delete represents a deletion operation.",
												Attributes: map[string]schema.Attribute{
													"bindings": schema.ListNestedAttribute{
														Description:         "Bindings defines additional binding key/values.",
														MarkdownDescription: "Bindings defines additional binding key/values.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name the name of the binding.",
																	MarkdownDescription: "Name the name of the binding.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
																	},
																},

																"value": schema.MapAttribute{
																	Description:         "Value value of the binding.",
																	MarkdownDescription: "Value value of the binding.",
																	ElementType:         types.StringType,
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

													"cluster": schema.StringAttribute{
														Description:         "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														MarkdownDescription: "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"clusters": schema.SingleNestedAttribute{
														Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
														MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Description:         "Context is the name of the context to use.",
																MarkdownDescription: "Context is the name of the context to use.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubeconfig": schema.StringAttribute{
																Description:         "Kubeconfig is the path to the referenced file.",
																MarkdownDescription: "Kubeconfig is the path to the referenced file.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"deletion_propagation_policy": schema.StringAttribute{
														Description:         "DeletionPropagationPolicy decides if a deletion will propagate to the dependents ofthe object, and how the garbage collector will handle the propagation.Overrides the deletion propagation policy set in the Configuration, the Test and the TestStep.",
														MarkdownDescription: "DeletionPropagationPolicy decides if a deletion will propagate to the dependents ofthe object, and how the garbage collector will handle the propagation.Overrides the deletion propagation policy set in the Configuration, the Test and the TestStep.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("Orphan", "Background", "Foreground"),
														},
													},

													"expect": schema.ListNestedAttribute{
														Description:         "Expect defines a list of matched checks to validate the operation outcome.",
														MarkdownDescription: "Expect defines a list of matched checks to validate the operation outcome.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"check": schema.MapAttribute{
																	Description:         "Check defines the verification statement.",
																	MarkdownDescription: "Check defines the verification statement.",
																	ElementType:         types.StringType,
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"match": schema.MapAttribute{
																	Description:         "Match defines the matching statement.",
																	MarkdownDescription: "Match defines the matching statement.",
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

													"file": schema.StringAttribute{
														Description:         "File is the path to the referenced file. This can be a direct path to a fileor an expression that matches multiple files, such as 'manifest/*.yaml' for all YAMLfiles within the 'manifest' directory.",
														MarkdownDescription: "File is the path to the referenced file. This can be a direct path to a fileor an expression that matches multiple files, such as 'manifest/*.yaml' for all YAMLfiles within the 'manifest' directory.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ref": schema.SingleNestedAttribute{
														Description:         "Ref determines objects to be deleted.",
														MarkdownDescription: "Ref determines objects to be deleted.",
														Attributes: map[string]schema.Attribute{
															"api_version": schema.StringAttribute{
																Description:         "API version of the referent.",
																MarkdownDescription: "API version of the referent.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"kind": schema.StringAttribute{
																Description:         "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
																MarkdownDescription: "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"labels": schema.MapAttribute{
																Description:         "Label selector to match objects to delete",
																MarkdownDescription: "Label selector to match objects to delete",
																ElementType:         types.StringType,
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"template": schema.BoolAttribute{
														Description:         "Template determines whether resources should be considered for templating.",
														MarkdownDescription: "Template determines whether resources should be considered for templating.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.StringAttribute{
														Description:         "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														MarkdownDescription: "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"describe": schema.SingleNestedAttribute{
												Description:         "Describe determines the resource describe collector to execute.",
												MarkdownDescription: "Describe determines the resource describe collector to execute.",
												Attributes: map[string]schema.Attribute{
													"api_version": schema.StringAttribute{
														Description:         "API version of the referent.",
														MarkdownDescription: "API version of the referent.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"cluster": schema.StringAttribute{
														Description:         "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														MarkdownDescription: "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"clusters": schema.SingleNestedAttribute{
														Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
														MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Description:         "Context is the name of the context to use.",
																MarkdownDescription: "Context is the name of the context to use.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubeconfig": schema.StringAttribute{
																Description:         "Kubeconfig is the path to the referenced file.",
																MarkdownDescription: "Kubeconfig is the path to the referenced file.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"kind": schema.StringAttribute{
														Description:         "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
														MarkdownDescription: "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
														Required:            true,
														Optional:            false,
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

													"selector": schema.StringAttribute{
														Description:         "Selector defines labels selector.",
														MarkdownDescription: "Selector defines labels selector.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"show_events": schema.BoolAttribute{
														Description:         "Show Events indicates whether to include related events.",
														MarkdownDescription: "Show Events indicates whether to include related events.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.StringAttribute{
														Description:         "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														MarkdownDescription: "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"description": schema.StringAttribute{
												Description:         "Description contains a description of the operation.",
												MarkdownDescription: "Description contains a description of the operation.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"events": schema.SingleNestedAttribute{
												Description:         "Events determines the events collector to execute.",
												MarkdownDescription: "Events determines the events collector to execute.",
												Attributes: map[string]schema.Attribute{
													"cluster": schema.StringAttribute{
														Description:         "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														MarkdownDescription: "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"clusters": schema.SingleNestedAttribute{
														Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
														MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Description:         "Context is the name of the context to use.",
																MarkdownDescription: "Context is the name of the context to use.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubeconfig": schema.StringAttribute{
																Description:         "Kubeconfig is the path to the referenced file.",
																MarkdownDescription: "Kubeconfig is the path to the referenced file.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"format": schema.StringAttribute{
														Description:         "Format determines the output format (json or yaml).",
														MarkdownDescription: "Format determines the output format (json or yaml).",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.RegexMatches(regexp.MustCompile(`^(?:json|yaml|\(.+\))$`), ""),
														},
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

													"selector": schema.StringAttribute{
														Description:         "Selector defines labels selector.",
														MarkdownDescription: "Selector defines labels selector.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.StringAttribute{
														Description:         "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														MarkdownDescription: "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"get": schema.SingleNestedAttribute{
												Description:         "Get determines the resource get collector to execute.",
												MarkdownDescription: "Get determines the resource get collector to execute.",
												Attributes: map[string]schema.Attribute{
													"api_version": schema.StringAttribute{
														Description:         "API version of the referent.",
														MarkdownDescription: "API version of the referent.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"cluster": schema.StringAttribute{
														Description:         "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														MarkdownDescription: "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"clusters": schema.SingleNestedAttribute{
														Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
														MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Description:         "Context is the name of the context to use.",
																MarkdownDescription: "Context is the name of the context to use.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubeconfig": schema.StringAttribute{
																Description:         "Kubeconfig is the path to the referenced file.",
																MarkdownDescription: "Kubeconfig is the path to the referenced file.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"format": schema.StringAttribute{
														Description:         "Format determines the output format (json or yaml).",
														MarkdownDescription: "Format determines the output format (json or yaml).",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.RegexMatches(regexp.MustCompile(`^(?:json|yaml|\(.+\))$`), ""),
														},
													},

													"kind": schema.StringAttribute{
														Description:         "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
														MarkdownDescription: "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
														Required:            true,
														Optional:            false,
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

													"selector": schema.StringAttribute{
														Description:         "Selector defines labels selector.",
														MarkdownDescription: "Selector defines labels selector.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.StringAttribute{
														Description:         "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														MarkdownDescription: "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"pod_logs": schema.SingleNestedAttribute{
												Description:         "PodLogs determines the pod logs collector to execute.",
												MarkdownDescription: "PodLogs determines the pod logs collector to execute.",
												Attributes: map[string]schema.Attribute{
													"cluster": schema.StringAttribute{
														Description:         "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														MarkdownDescription: "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"clusters": schema.SingleNestedAttribute{
														Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
														MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Description:         "Context is the name of the context to use.",
																MarkdownDescription: "Context is the name of the context to use.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubeconfig": schema.StringAttribute{
																Description:         "Kubeconfig is the path to the referenced file.",
																MarkdownDescription: "Kubeconfig is the path to the referenced file.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"container": schema.StringAttribute{
														Description:         "Container in pod to get logs from else --all-containers is used.",
														MarkdownDescription: "Container in pod to get logs from else --all-containers is used.",
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

													"selector": schema.StringAttribute{
														Description:         "Selector defines labels selector.",
														MarkdownDescription: "Selector defines labels selector.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"tail": schema.Int64Attribute{
														Description:         "Tail is the number of last lines to collect from pods. If omitted or zero,then the default is 10 if you use a selector, or -1 (all) if you use a pod name.This matches default behavior of 'kubectl logs'.",
														MarkdownDescription: "Tail is the number of last lines to collect from pods. If omitted or zero,then the default is 10 if you use a selector, or -1 (all) if you use a pod name.This matches default behavior of 'kubectl logs'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.StringAttribute{
														Description:         "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														MarkdownDescription: "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"script": schema.SingleNestedAttribute{
												Description:         "Script defines a script to run.",
												MarkdownDescription: "Script defines a script to run.",
												Attributes: map[string]schema.Attribute{
													"bindings": schema.ListNestedAttribute{
														Description:         "Bindings defines additional binding key/values.",
														MarkdownDescription: "Bindings defines additional binding key/values.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name the name of the binding.",
																	MarkdownDescription: "Name the name of the binding.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
																	},
																},

																"value": schema.MapAttribute{
																	Description:         "Value value of the binding.",
																	MarkdownDescription: "Value value of the binding.",
																	ElementType:         types.StringType,
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

													"check": schema.MapAttribute{
														Description:         "Check is an assertion tree to validate the operation outcome.",
														MarkdownDescription: "Check is an assertion tree to validate the operation outcome.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"cluster": schema.StringAttribute{
														Description:         "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														MarkdownDescription: "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"clusters": schema.SingleNestedAttribute{
														Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
														MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Description:         "Context is the name of the context to use.",
																MarkdownDescription: "Context is the name of the context to use.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubeconfig": schema.StringAttribute{
																Description:         "Kubeconfig is the path to the referenced file.",
																MarkdownDescription: "Kubeconfig is the path to the referenced file.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"content": schema.StringAttribute{
														Description:         "Content defines a shell script (run with 'sh -c ...').",
														MarkdownDescription: "Content defines a shell script (run with 'sh -c ...').",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"env": schema.ListNestedAttribute{
														Description:         "Env defines additional environment variables.",
														MarkdownDescription: "Env defines additional environment variables.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name the name of the binding.",
																	MarkdownDescription: "Name the name of the binding.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
																	},
																},

																"value": schema.MapAttribute{
																	Description:         "Value value of the binding.",
																	MarkdownDescription: "Value value of the binding.",
																	ElementType:         types.StringType,
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

													"outputs": schema.ListNestedAttribute{
														Description:         "Outputs defines output bindings.",
														MarkdownDescription: "Outputs defines output bindings.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"match": schema.MapAttribute{
																	Description:         "Match defines the matching statement.",
																	MarkdownDescription: "Match defines the matching statement.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name the name of the binding.",
																	MarkdownDescription: "Name the name of the binding.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
																	},
																},

																"value": schema.MapAttribute{
																	Description:         "Value value of the binding.",
																	MarkdownDescription: "Value value of the binding.",
																	ElementType:         types.StringType,
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

													"skip_log_output": schema.BoolAttribute{
														Description:         "SkipLogOutput removes the output from the command. Useful for sensitive logs or to reduce noise.",
														MarkdownDescription: "SkipLogOutput removes the output from the command. Useful for sensitive logs or to reduce noise.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.StringAttribute{
														Description:         "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														MarkdownDescription: "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"sleep": schema.SingleNestedAttribute{
												Description:         "Sleep defines zzzz.",
												MarkdownDescription: "Sleep defines zzzz.",
												Attributes: map[string]schema.Attribute{
													"duration": schema.StringAttribute{
														Description:         "Duration is the delay used for sleeping.",
														MarkdownDescription: "Duration is the delay used for sleeping.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"wait": schema.SingleNestedAttribute{
												Description:         "Wait determines the resource wait collector to execute.",
												MarkdownDescription: "Wait determines the resource wait collector to execute.",
												Attributes: map[string]schema.Attribute{
													"api_version": schema.StringAttribute{
														Description:         "API version of the referent.",
														MarkdownDescription: "API version of the referent.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"cluster": schema.StringAttribute{
														Description:         "Cluster defines the target cluster where the wait operation will be performed (default cluster will be used if not specified).",
														MarkdownDescription: "Cluster defines the target cluster where the wait operation will be performed (default cluster will be used if not specified).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"clusters": schema.SingleNestedAttribute{
														Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
														MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Description:         "Context is the name of the context to use.",
																MarkdownDescription: "Context is the name of the context to use.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubeconfig": schema.StringAttribute{
																Description:         "Kubeconfig is the path to the referenced file.",
																MarkdownDescription: "Kubeconfig is the path to the referenced file.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"for": schema.SingleNestedAttribute{
														Description:         "For specifies the condition to wait for.",
														MarkdownDescription: "For specifies the condition to wait for.",
														Attributes: map[string]schema.Attribute{
															"condition": schema.SingleNestedAttribute{
																Description:         "Condition specifies the condition to wait for.",
																MarkdownDescription: "Condition specifies the condition to wait for.",
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "Name defines the specific condition to wait for, e.g., 'Available', 'Ready'.",
																		MarkdownDescription: "Name defines the specific condition to wait for, e.g., 'Available', 'Ready'.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
																		Description:         "Value defines the specific condition status to wait for, e.g., 'True', 'False'.",
																		MarkdownDescription: "Value defines the specific condition status to wait for, e.g., 'True', 'False'.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"deletion": schema.MapAttribute{
																Description:         "Deletion specifies parameters for waiting on a resource's deletion.",
																MarkdownDescription: "Deletion specifies parameters for waiting on a resource's deletion.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"json_path": schema.SingleNestedAttribute{
																Description:         "JsonPath specifies the json path condition to wait for.",
																MarkdownDescription: "JsonPath specifies the json path condition to wait for.",
																Attributes: map[string]schema.Attribute{
																	"path": schema.StringAttribute{
																		Description:         "Path defines the json path to wait for, e.g. '{.status.phase}'.",
																		MarkdownDescription: "Path defines the json path to wait for, e.g. '{.status.phase}'.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
																		Description:         "Value defines the expected value to wait for, e.g., 'Running'.",
																		MarkdownDescription: "Value defines the expected value to wait for, e.g., 'Running'.",
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
														Required: true,
														Optional: false,
														Computed: false,
													},

													"format": schema.StringAttribute{
														Description:         "Format determines the output format (json or yaml).",
														MarkdownDescription: "Format determines the output format (json or yaml).",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.RegexMatches(regexp.MustCompile(`^(?:json|yaml|\(.+\))$`), ""),
														},
													},

													"kind": schema.StringAttribute{
														Description:         "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
														MarkdownDescription: "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
														Required:            true,
														Optional:            false,
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

													"selector": schema.StringAttribute{
														Description:         "Selector defines labels selector.",
														MarkdownDescription: "Selector defines labels selector.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.StringAttribute{
														Description:         "Timeout for the operation. Specifies how long to wait for the condition to be met before timing out.",
														MarkdownDescription: "Timeout for the operation. Specifies how long to wait for the condition to be met before timing out.",
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

								"name": schema.StringAttribute{
									Description:         "Name of the step.",
									MarkdownDescription: "Name of the step.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"skip_delete": schema.BoolAttribute{
									Description:         "SkipDelete determines whether the resources created by the step should be deleted after the test step is executed.",
									MarkdownDescription: "SkipDelete determines whether the resources created by the step should be deleted after the test step is executed.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"template": schema.BoolAttribute{
									Description:         "Template determines whether resources should be considered for templating.",
									MarkdownDescription: "Template determines whether resources should be considered for templating.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"timeouts": schema.SingleNestedAttribute{
									Description:         "Timeouts for the test step. Overrides the global timeouts set in the Configuration and the timeouts eventually set in the Test.",
									MarkdownDescription: "Timeouts for the test step. Overrides the global timeouts set in the Configuration and the timeouts eventually set in the Test.",
									Attributes: map[string]schema.Attribute{
										"apply": schema.StringAttribute{
											Description:         "Apply defines the timeout for the apply operation",
											MarkdownDescription: "Apply defines the timeout for the apply operation",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"assert": schema.StringAttribute{
											Description:         "Assert defines the timeout for the assert operation",
											MarkdownDescription: "Assert defines the timeout for the assert operation",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"cleanup": schema.StringAttribute{
											Description:         "Cleanup defines the timeout for the cleanup operation",
											MarkdownDescription: "Cleanup defines the timeout for the cleanup operation",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"delete": schema.StringAttribute{
											Description:         "Delete defines the timeout for the delete operation",
											MarkdownDescription: "Delete defines the timeout for the delete operation",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"error": schema.StringAttribute{
											Description:         "Error defines the timeout for the error operation",
											MarkdownDescription: "Error defines the timeout for the error operation",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exec": schema.StringAttribute{
											Description:         "Exec defines the timeout for exec operations",
											MarkdownDescription: "Exec defines the timeout for exec operations",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"try": schema.ListNestedAttribute{
									Description:         "Try defines what the step will try to execute.",
									MarkdownDescription: "Try defines what the step will try to execute.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"apply": schema.SingleNestedAttribute{
												Description:         "Apply represents resources that should be applied for this test step. This can include thingslike configuration settings or any other resources that need to be available during the test.",
												MarkdownDescription: "Apply represents resources that should be applied for this test step. This can include thingslike configuration settings or any other resources that need to be available during the test.",
												Attributes: map[string]schema.Attribute{
													"bindings": schema.ListNestedAttribute{
														Description:         "Bindings defines additional binding key/values.",
														MarkdownDescription: "Bindings defines additional binding key/values.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name the name of the binding.",
																	MarkdownDescription: "Name the name of the binding.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
																	},
																},

																"value": schema.MapAttribute{
																	Description:         "Value value of the binding.",
																	MarkdownDescription: "Value value of the binding.",
																	ElementType:         types.StringType,
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

													"cluster": schema.StringAttribute{
														Description:         "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														MarkdownDescription: "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"clusters": schema.SingleNestedAttribute{
														Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
														MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Description:         "Context is the name of the context to use.",
																MarkdownDescription: "Context is the name of the context to use.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubeconfig": schema.StringAttribute{
																Description:         "Kubeconfig is the path to the referenced file.",
																MarkdownDescription: "Kubeconfig is the path to the referenced file.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"dry_run": schema.BoolAttribute{
														Description:         "DryRun determines whether the file should be applied in dry run mode.",
														MarkdownDescription: "DryRun determines whether the file should be applied in dry run mode.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"expect": schema.ListNestedAttribute{
														Description:         "Expect defines a list of matched checks to validate the operation outcome.",
														MarkdownDescription: "Expect defines a list of matched checks to validate the operation outcome.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"check": schema.MapAttribute{
																	Description:         "Check defines the verification statement.",
																	MarkdownDescription: "Check defines the verification statement.",
																	ElementType:         types.StringType,
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"match": schema.MapAttribute{
																	Description:         "Match defines the matching statement.",
																	MarkdownDescription: "Match defines the matching statement.",
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

													"file": schema.StringAttribute{
														Description:         "File is the path to the referenced file. This can be a direct path to a fileor an expression that matches multiple files, such as 'manifest/*.yaml' for all YAMLfiles within the 'manifest' directory.",
														MarkdownDescription: "File is the path to the referenced file. This can be a direct path to a fileor an expression that matches multiple files, such as 'manifest/*.yaml' for all YAMLfiles within the 'manifest' directory.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"outputs": schema.ListNestedAttribute{
														Description:         "Outputs defines output bindings.",
														MarkdownDescription: "Outputs defines output bindings.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"match": schema.MapAttribute{
																	Description:         "Match defines the matching statement.",
																	MarkdownDescription: "Match defines the matching statement.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name the name of the binding.",
																	MarkdownDescription: "Name the name of the binding.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
																	},
																},

																"value": schema.MapAttribute{
																	Description:         "Value value of the binding.",
																	MarkdownDescription: "Value value of the binding.",
																	ElementType:         types.StringType,
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

													"resource": schema.MapAttribute{
														Description:         "Resource provides a resource to be applied.",
														MarkdownDescription: "Resource provides a resource to be applied.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"template": schema.BoolAttribute{
														Description:         "Template determines whether resources should be considered for templating.",
														MarkdownDescription: "Template determines whether resources should be considered for templating.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.StringAttribute{
														Description:         "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														MarkdownDescription: "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"assert": schema.SingleNestedAttribute{
												Description:         "Assert represents an assertion to be made. It checks whether the conditions specified in the assertion hold true.",
												MarkdownDescription: "Assert represents an assertion to be made. It checks whether the conditions specified in the assertion hold true.",
												Attributes: map[string]schema.Attribute{
													"bindings": schema.ListNestedAttribute{
														Description:         "Bindings defines additional binding key/values.",
														MarkdownDescription: "Bindings defines additional binding key/values.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name the name of the binding.",
																	MarkdownDescription: "Name the name of the binding.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
																	},
																},

																"value": schema.MapAttribute{
																	Description:         "Value value of the binding.",
																	MarkdownDescription: "Value value of the binding.",
																	ElementType:         types.StringType,
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

													"cluster": schema.StringAttribute{
														Description:         "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														MarkdownDescription: "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"clusters": schema.SingleNestedAttribute{
														Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
														MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Description:         "Context is the name of the context to use.",
																MarkdownDescription: "Context is the name of the context to use.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubeconfig": schema.StringAttribute{
																Description:         "Kubeconfig is the path to the referenced file.",
																MarkdownDescription: "Kubeconfig is the path to the referenced file.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"file": schema.StringAttribute{
														Description:         "File is the path to the referenced file. This can be a direct path to a fileor an expression that matches multiple files, such as 'manifest/*.yaml' for all YAMLfiles within the 'manifest' directory.",
														MarkdownDescription: "File is the path to the referenced file. This can be a direct path to a fileor an expression that matches multiple files, such as 'manifest/*.yaml' for all YAMLfiles within the 'manifest' directory.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"resource": schema.MapAttribute{
														Description:         "Check provides a check used in assertions.",
														MarkdownDescription: "Check provides a check used in assertions.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"template": schema.BoolAttribute{
														Description:         "Template determines whether resources should be considered for templating.",
														MarkdownDescription: "Template determines whether resources should be considered for templating.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.StringAttribute{
														Description:         "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														MarkdownDescription: "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"command": schema.SingleNestedAttribute{
												Description:         "Command defines a command to run.",
												MarkdownDescription: "Command defines a command to run.",
												Attributes: map[string]schema.Attribute{
													"args": schema.ListAttribute{
														Description:         "Args is the command arguments.",
														MarkdownDescription: "Args is the command arguments.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"bindings": schema.ListNestedAttribute{
														Description:         "Bindings defines additional binding key/values.",
														MarkdownDescription: "Bindings defines additional binding key/values.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name the name of the binding.",
																	MarkdownDescription: "Name the name of the binding.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
																	},
																},

																"value": schema.MapAttribute{
																	Description:         "Value value of the binding.",
																	MarkdownDescription: "Value value of the binding.",
																	ElementType:         types.StringType,
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

													"check": schema.MapAttribute{
														Description:         "Check is an assertion tree to validate the operation outcome.",
														MarkdownDescription: "Check is an assertion tree to validate the operation outcome.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"cluster": schema.StringAttribute{
														Description:         "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														MarkdownDescription: "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"clusters": schema.SingleNestedAttribute{
														Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
														MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Description:         "Context is the name of the context to use.",
																MarkdownDescription: "Context is the name of the context to use.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubeconfig": schema.StringAttribute{
																Description:         "Kubeconfig is the path to the referenced file.",
																MarkdownDescription: "Kubeconfig is the path to the referenced file.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"entrypoint": schema.StringAttribute{
														Description:         "Entrypoint is the command entry point to run.",
														MarkdownDescription: "Entrypoint is the command entry point to run.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"env": schema.ListNestedAttribute{
														Description:         "Env defines additional environment variables.",
														MarkdownDescription: "Env defines additional environment variables.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name the name of the binding.",
																	MarkdownDescription: "Name the name of the binding.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
																	},
																},

																"value": schema.MapAttribute{
																	Description:         "Value value of the binding.",
																	MarkdownDescription: "Value value of the binding.",
																	ElementType:         types.StringType,
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

													"outputs": schema.ListNestedAttribute{
														Description:         "Outputs defines output bindings.",
														MarkdownDescription: "Outputs defines output bindings.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"match": schema.MapAttribute{
																	Description:         "Match defines the matching statement.",
																	MarkdownDescription: "Match defines the matching statement.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name the name of the binding.",
																	MarkdownDescription: "Name the name of the binding.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
																	},
																},

																"value": schema.MapAttribute{
																	Description:         "Value value of the binding.",
																	MarkdownDescription: "Value value of the binding.",
																	ElementType:         types.StringType,
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

													"skip_log_output": schema.BoolAttribute{
														Description:         "SkipLogOutput removes the output from the command. Useful for sensitive logs or to reduce noise.",
														MarkdownDescription: "SkipLogOutput removes the output from the command. Useful for sensitive logs or to reduce noise.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.StringAttribute{
														Description:         "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														MarkdownDescription: "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"continue_on_error": schema.BoolAttribute{
												Description:         "ContinueOnError determines whether a test should continue or not in case the operation was not successful.Even if the test continues executing, it will still be reported as failed.",
												MarkdownDescription: "ContinueOnError determines whether a test should continue or not in case the operation was not successful.Even if the test continues executing, it will still be reported as failed.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"create": schema.SingleNestedAttribute{
												Description:         "Create represents a creation operation.",
												MarkdownDescription: "Create represents a creation operation.",
												Attributes: map[string]schema.Attribute{
													"bindings": schema.ListNestedAttribute{
														Description:         "Bindings defines additional binding key/values.",
														MarkdownDescription: "Bindings defines additional binding key/values.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name the name of the binding.",
																	MarkdownDescription: "Name the name of the binding.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
																	},
																},

																"value": schema.MapAttribute{
																	Description:         "Value value of the binding.",
																	MarkdownDescription: "Value value of the binding.",
																	ElementType:         types.StringType,
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

													"cluster": schema.StringAttribute{
														Description:         "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														MarkdownDescription: "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"clusters": schema.SingleNestedAttribute{
														Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
														MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Description:         "Context is the name of the context to use.",
																MarkdownDescription: "Context is the name of the context to use.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubeconfig": schema.StringAttribute{
																Description:         "Kubeconfig is the path to the referenced file.",
																MarkdownDescription: "Kubeconfig is the path to the referenced file.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"dry_run": schema.BoolAttribute{
														Description:         "DryRun determines whether the file should be applied in dry run mode.",
														MarkdownDescription: "DryRun determines whether the file should be applied in dry run mode.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"expect": schema.ListNestedAttribute{
														Description:         "Expect defines a list of matched checks to validate the operation outcome.",
														MarkdownDescription: "Expect defines a list of matched checks to validate the operation outcome.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"check": schema.MapAttribute{
																	Description:         "Check defines the verification statement.",
																	MarkdownDescription: "Check defines the verification statement.",
																	ElementType:         types.StringType,
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"match": schema.MapAttribute{
																	Description:         "Match defines the matching statement.",
																	MarkdownDescription: "Match defines the matching statement.",
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

													"file": schema.StringAttribute{
														Description:         "File is the path to the referenced file. This can be a direct path to a fileor an expression that matches multiple files, such as 'manifest/*.yaml' for all YAMLfiles within the 'manifest' directory.",
														MarkdownDescription: "File is the path to the referenced file. This can be a direct path to a fileor an expression that matches multiple files, such as 'manifest/*.yaml' for all YAMLfiles within the 'manifest' directory.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"outputs": schema.ListNestedAttribute{
														Description:         "Outputs defines output bindings.",
														MarkdownDescription: "Outputs defines output bindings.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"match": schema.MapAttribute{
																	Description:         "Match defines the matching statement.",
																	MarkdownDescription: "Match defines the matching statement.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name the name of the binding.",
																	MarkdownDescription: "Name the name of the binding.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
																	},
																},

																"value": schema.MapAttribute{
																	Description:         "Value value of the binding.",
																	MarkdownDescription: "Value value of the binding.",
																	ElementType:         types.StringType,
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

													"resource": schema.MapAttribute{
														Description:         "Resource provides a resource to be applied.",
														MarkdownDescription: "Resource provides a resource to be applied.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"template": schema.BoolAttribute{
														Description:         "Template determines whether resources should be considered for templating.",
														MarkdownDescription: "Template determines whether resources should be considered for templating.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.StringAttribute{
														Description:         "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														MarkdownDescription: "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"delete": schema.SingleNestedAttribute{
												Description:         "Delete represents a deletion operation.",
												MarkdownDescription: "Delete represents a deletion operation.",
												Attributes: map[string]schema.Attribute{
													"bindings": schema.ListNestedAttribute{
														Description:         "Bindings defines additional binding key/values.",
														MarkdownDescription: "Bindings defines additional binding key/values.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name the name of the binding.",
																	MarkdownDescription: "Name the name of the binding.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
																	},
																},

																"value": schema.MapAttribute{
																	Description:         "Value value of the binding.",
																	MarkdownDescription: "Value value of the binding.",
																	ElementType:         types.StringType,
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

													"cluster": schema.StringAttribute{
														Description:         "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														MarkdownDescription: "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"clusters": schema.SingleNestedAttribute{
														Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
														MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Description:         "Context is the name of the context to use.",
																MarkdownDescription: "Context is the name of the context to use.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubeconfig": schema.StringAttribute{
																Description:         "Kubeconfig is the path to the referenced file.",
																MarkdownDescription: "Kubeconfig is the path to the referenced file.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"deletion_propagation_policy": schema.StringAttribute{
														Description:         "DeletionPropagationPolicy decides if a deletion will propagate to the dependents ofthe object, and how the garbage collector will handle the propagation.Overrides the deletion propagation policy set in the Configuration, the Test and the TestStep.",
														MarkdownDescription: "DeletionPropagationPolicy decides if a deletion will propagate to the dependents ofthe object, and how the garbage collector will handle the propagation.Overrides the deletion propagation policy set in the Configuration, the Test and the TestStep.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("Orphan", "Background", "Foreground"),
														},
													},

													"expect": schema.ListNestedAttribute{
														Description:         "Expect defines a list of matched checks to validate the operation outcome.",
														MarkdownDescription: "Expect defines a list of matched checks to validate the operation outcome.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"check": schema.MapAttribute{
																	Description:         "Check defines the verification statement.",
																	MarkdownDescription: "Check defines the verification statement.",
																	ElementType:         types.StringType,
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"match": schema.MapAttribute{
																	Description:         "Match defines the matching statement.",
																	MarkdownDescription: "Match defines the matching statement.",
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

													"file": schema.StringAttribute{
														Description:         "File is the path to the referenced file. This can be a direct path to a fileor an expression that matches multiple files, such as 'manifest/*.yaml' for all YAMLfiles within the 'manifest' directory.",
														MarkdownDescription: "File is the path to the referenced file. This can be a direct path to a fileor an expression that matches multiple files, such as 'manifest/*.yaml' for all YAMLfiles within the 'manifest' directory.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ref": schema.SingleNestedAttribute{
														Description:         "Ref determines objects to be deleted.",
														MarkdownDescription: "Ref determines objects to be deleted.",
														Attributes: map[string]schema.Attribute{
															"api_version": schema.StringAttribute{
																Description:         "API version of the referent.",
																MarkdownDescription: "API version of the referent.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"kind": schema.StringAttribute{
																Description:         "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
																MarkdownDescription: "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"labels": schema.MapAttribute{
																Description:         "Label selector to match objects to delete",
																MarkdownDescription: "Label selector to match objects to delete",
																ElementType:         types.StringType,
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"template": schema.BoolAttribute{
														Description:         "Template determines whether resources should be considered for templating.",
														MarkdownDescription: "Template determines whether resources should be considered for templating.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.StringAttribute{
														Description:         "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														MarkdownDescription: "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"describe": schema.SingleNestedAttribute{
												Description:         "Describe determines the resource describe collector to execute.",
												MarkdownDescription: "Describe determines the resource describe collector to execute.",
												Attributes: map[string]schema.Attribute{
													"api_version": schema.StringAttribute{
														Description:         "API version of the referent.",
														MarkdownDescription: "API version of the referent.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"cluster": schema.StringAttribute{
														Description:         "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														MarkdownDescription: "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"clusters": schema.SingleNestedAttribute{
														Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
														MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Description:         "Context is the name of the context to use.",
																MarkdownDescription: "Context is the name of the context to use.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubeconfig": schema.StringAttribute{
																Description:         "Kubeconfig is the path to the referenced file.",
																MarkdownDescription: "Kubeconfig is the path to the referenced file.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"kind": schema.StringAttribute{
														Description:         "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
														MarkdownDescription: "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
														Required:            true,
														Optional:            false,
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

													"selector": schema.StringAttribute{
														Description:         "Selector defines labels selector.",
														MarkdownDescription: "Selector defines labels selector.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"show_events": schema.BoolAttribute{
														Description:         "Show Events indicates whether to include related events.",
														MarkdownDescription: "Show Events indicates whether to include related events.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.StringAttribute{
														Description:         "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														MarkdownDescription: "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"description": schema.StringAttribute{
												Description:         "Description contains a description of the operation.",
												MarkdownDescription: "Description contains a description of the operation.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"error": schema.SingleNestedAttribute{
												Description:         "Error represents the expected errors for this test step. If any of these errors occur, the testwill consider them as expected; otherwise, they will be treated as test failures.",
												MarkdownDescription: "Error represents the expected errors for this test step. If any of these errors occur, the testwill consider them as expected; otherwise, they will be treated as test failures.",
												Attributes: map[string]schema.Attribute{
													"bindings": schema.ListNestedAttribute{
														Description:         "Bindings defines additional binding key/values.",
														MarkdownDescription: "Bindings defines additional binding key/values.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name the name of the binding.",
																	MarkdownDescription: "Name the name of the binding.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
																	},
																},

																"value": schema.MapAttribute{
																	Description:         "Value value of the binding.",
																	MarkdownDescription: "Value value of the binding.",
																	ElementType:         types.StringType,
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

													"cluster": schema.StringAttribute{
														Description:         "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														MarkdownDescription: "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"clusters": schema.SingleNestedAttribute{
														Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
														MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Description:         "Context is the name of the context to use.",
																MarkdownDescription: "Context is the name of the context to use.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubeconfig": schema.StringAttribute{
																Description:         "Kubeconfig is the path to the referenced file.",
																MarkdownDescription: "Kubeconfig is the path to the referenced file.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"file": schema.StringAttribute{
														Description:         "File is the path to the referenced file. This can be a direct path to a fileor an expression that matches multiple files, such as 'manifest/*.yaml' for all YAMLfiles within the 'manifest' directory.",
														MarkdownDescription: "File is the path to the referenced file. This can be a direct path to a fileor an expression that matches multiple files, such as 'manifest/*.yaml' for all YAMLfiles within the 'manifest' directory.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"resource": schema.MapAttribute{
														Description:         "Check provides a check used in assertions.",
														MarkdownDescription: "Check provides a check used in assertions.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"template": schema.BoolAttribute{
														Description:         "Template determines whether resources should be considered for templating.",
														MarkdownDescription: "Template determines whether resources should be considered for templating.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.StringAttribute{
														Description:         "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														MarkdownDescription: "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"events": schema.SingleNestedAttribute{
												Description:         "Events determines the events collector to execute.",
												MarkdownDescription: "Events determines the events collector to execute.",
												Attributes: map[string]schema.Attribute{
													"cluster": schema.StringAttribute{
														Description:         "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														MarkdownDescription: "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"clusters": schema.SingleNestedAttribute{
														Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
														MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Description:         "Context is the name of the context to use.",
																MarkdownDescription: "Context is the name of the context to use.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubeconfig": schema.StringAttribute{
																Description:         "Kubeconfig is the path to the referenced file.",
																MarkdownDescription: "Kubeconfig is the path to the referenced file.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"format": schema.StringAttribute{
														Description:         "Format determines the output format (json or yaml).",
														MarkdownDescription: "Format determines the output format (json or yaml).",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.RegexMatches(regexp.MustCompile(`^(?:json|yaml|\(.+\))$`), ""),
														},
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

													"selector": schema.StringAttribute{
														Description:         "Selector defines labels selector.",
														MarkdownDescription: "Selector defines labels selector.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.StringAttribute{
														Description:         "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														MarkdownDescription: "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"get": schema.SingleNestedAttribute{
												Description:         "Get determines the resource get collector to execute.",
												MarkdownDescription: "Get determines the resource get collector to execute.",
												Attributes: map[string]schema.Attribute{
													"api_version": schema.StringAttribute{
														Description:         "API version of the referent.",
														MarkdownDescription: "API version of the referent.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"cluster": schema.StringAttribute{
														Description:         "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														MarkdownDescription: "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"clusters": schema.SingleNestedAttribute{
														Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
														MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Description:         "Context is the name of the context to use.",
																MarkdownDescription: "Context is the name of the context to use.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubeconfig": schema.StringAttribute{
																Description:         "Kubeconfig is the path to the referenced file.",
																MarkdownDescription: "Kubeconfig is the path to the referenced file.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"format": schema.StringAttribute{
														Description:         "Format determines the output format (json or yaml).",
														MarkdownDescription: "Format determines the output format (json or yaml).",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.RegexMatches(regexp.MustCompile(`^(?:json|yaml|\(.+\))$`), ""),
														},
													},

													"kind": schema.StringAttribute{
														Description:         "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
														MarkdownDescription: "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
														Required:            true,
														Optional:            false,
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

													"selector": schema.StringAttribute{
														Description:         "Selector defines labels selector.",
														MarkdownDescription: "Selector defines labels selector.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.StringAttribute{
														Description:         "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														MarkdownDescription: "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"patch": schema.SingleNestedAttribute{
												Description:         "Patch represents a patch operation.",
												MarkdownDescription: "Patch represents a patch operation.",
												Attributes: map[string]schema.Attribute{
													"bindings": schema.ListNestedAttribute{
														Description:         "Bindings defines additional binding key/values.",
														MarkdownDescription: "Bindings defines additional binding key/values.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name the name of the binding.",
																	MarkdownDescription: "Name the name of the binding.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
																	},
																},

																"value": schema.MapAttribute{
																	Description:         "Value value of the binding.",
																	MarkdownDescription: "Value value of the binding.",
																	ElementType:         types.StringType,
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

													"cluster": schema.StringAttribute{
														Description:         "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														MarkdownDescription: "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"clusters": schema.SingleNestedAttribute{
														Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
														MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Description:         "Context is the name of the context to use.",
																MarkdownDescription: "Context is the name of the context to use.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubeconfig": schema.StringAttribute{
																Description:         "Kubeconfig is the path to the referenced file.",
																MarkdownDescription: "Kubeconfig is the path to the referenced file.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"dry_run": schema.BoolAttribute{
														Description:         "DryRun determines whether the file should be applied in dry run mode.",
														MarkdownDescription: "DryRun determines whether the file should be applied in dry run mode.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"expect": schema.ListNestedAttribute{
														Description:         "Expect defines a list of matched checks to validate the operation outcome.",
														MarkdownDescription: "Expect defines a list of matched checks to validate the operation outcome.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"check": schema.MapAttribute{
																	Description:         "Check defines the verification statement.",
																	MarkdownDescription: "Check defines the verification statement.",
																	ElementType:         types.StringType,
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"match": schema.MapAttribute{
																	Description:         "Match defines the matching statement.",
																	MarkdownDescription: "Match defines the matching statement.",
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

													"file": schema.StringAttribute{
														Description:         "File is the path to the referenced file. This can be a direct path to a fileor an expression that matches multiple files, such as 'manifest/*.yaml' for all YAMLfiles within the 'manifest' directory.",
														MarkdownDescription: "File is the path to the referenced file. This can be a direct path to a fileor an expression that matches multiple files, such as 'manifest/*.yaml' for all YAMLfiles within the 'manifest' directory.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"outputs": schema.ListNestedAttribute{
														Description:         "Outputs defines output bindings.",
														MarkdownDescription: "Outputs defines output bindings.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"match": schema.MapAttribute{
																	Description:         "Match defines the matching statement.",
																	MarkdownDescription: "Match defines the matching statement.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name the name of the binding.",
																	MarkdownDescription: "Name the name of the binding.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
																	},
																},

																"value": schema.MapAttribute{
																	Description:         "Value value of the binding.",
																	MarkdownDescription: "Value value of the binding.",
																	ElementType:         types.StringType,
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

													"resource": schema.MapAttribute{
														Description:         "Resource provides a resource to be applied.",
														MarkdownDescription: "Resource provides a resource to be applied.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"template": schema.BoolAttribute{
														Description:         "Template determines whether resources should be considered for templating.",
														MarkdownDescription: "Template determines whether resources should be considered for templating.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.StringAttribute{
														Description:         "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														MarkdownDescription: "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"pod_logs": schema.SingleNestedAttribute{
												Description:         "PodLogs determines the pod logs collector to execute.",
												MarkdownDescription: "PodLogs determines the pod logs collector to execute.",
												Attributes: map[string]schema.Attribute{
													"cluster": schema.StringAttribute{
														Description:         "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														MarkdownDescription: "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"clusters": schema.SingleNestedAttribute{
														Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
														MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Description:         "Context is the name of the context to use.",
																MarkdownDescription: "Context is the name of the context to use.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubeconfig": schema.StringAttribute{
																Description:         "Kubeconfig is the path to the referenced file.",
																MarkdownDescription: "Kubeconfig is the path to the referenced file.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"container": schema.StringAttribute{
														Description:         "Container in pod to get logs from else --all-containers is used.",
														MarkdownDescription: "Container in pod to get logs from else --all-containers is used.",
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

													"selector": schema.StringAttribute{
														Description:         "Selector defines labels selector.",
														MarkdownDescription: "Selector defines labels selector.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"tail": schema.Int64Attribute{
														Description:         "Tail is the number of last lines to collect from pods. If omitted or zero,then the default is 10 if you use a selector, or -1 (all) if you use a pod name.This matches default behavior of 'kubectl logs'.",
														MarkdownDescription: "Tail is the number of last lines to collect from pods. If omitted or zero,then the default is 10 if you use a selector, or -1 (all) if you use a pod name.This matches default behavior of 'kubectl logs'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.StringAttribute{
														Description:         "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														MarkdownDescription: "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"script": schema.SingleNestedAttribute{
												Description:         "Script defines a script to run.",
												MarkdownDescription: "Script defines a script to run.",
												Attributes: map[string]schema.Attribute{
													"bindings": schema.ListNestedAttribute{
														Description:         "Bindings defines additional binding key/values.",
														MarkdownDescription: "Bindings defines additional binding key/values.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name the name of the binding.",
																	MarkdownDescription: "Name the name of the binding.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
																	},
																},

																"value": schema.MapAttribute{
																	Description:         "Value value of the binding.",
																	MarkdownDescription: "Value value of the binding.",
																	ElementType:         types.StringType,
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

													"check": schema.MapAttribute{
														Description:         "Check is an assertion tree to validate the operation outcome.",
														MarkdownDescription: "Check is an assertion tree to validate the operation outcome.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"cluster": schema.StringAttribute{
														Description:         "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														MarkdownDescription: "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"clusters": schema.SingleNestedAttribute{
														Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
														MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Description:         "Context is the name of the context to use.",
																MarkdownDescription: "Context is the name of the context to use.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubeconfig": schema.StringAttribute{
																Description:         "Kubeconfig is the path to the referenced file.",
																MarkdownDescription: "Kubeconfig is the path to the referenced file.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"content": schema.StringAttribute{
														Description:         "Content defines a shell script (run with 'sh -c ...').",
														MarkdownDescription: "Content defines a shell script (run with 'sh -c ...').",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"env": schema.ListNestedAttribute{
														Description:         "Env defines additional environment variables.",
														MarkdownDescription: "Env defines additional environment variables.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name the name of the binding.",
																	MarkdownDescription: "Name the name of the binding.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
																	},
																},

																"value": schema.MapAttribute{
																	Description:         "Value value of the binding.",
																	MarkdownDescription: "Value value of the binding.",
																	ElementType:         types.StringType,
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

													"outputs": schema.ListNestedAttribute{
														Description:         "Outputs defines output bindings.",
														MarkdownDescription: "Outputs defines output bindings.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"match": schema.MapAttribute{
																	Description:         "Match defines the matching statement.",
																	MarkdownDescription: "Match defines the matching statement.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name the name of the binding.",
																	MarkdownDescription: "Name the name of the binding.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
																	},
																},

																"value": schema.MapAttribute{
																	Description:         "Value value of the binding.",
																	MarkdownDescription: "Value value of the binding.",
																	ElementType:         types.StringType,
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

													"skip_log_output": schema.BoolAttribute{
														Description:         "SkipLogOutput removes the output from the command. Useful for sensitive logs or to reduce noise.",
														MarkdownDescription: "SkipLogOutput removes the output from the command. Useful for sensitive logs or to reduce noise.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.StringAttribute{
														Description:         "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														MarkdownDescription: "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"sleep": schema.SingleNestedAttribute{
												Description:         "Sleep defines zzzz.",
												MarkdownDescription: "Sleep defines zzzz.",
												Attributes: map[string]schema.Attribute{
													"duration": schema.StringAttribute{
														Description:         "Duration is the delay used for sleeping.",
														MarkdownDescription: "Duration is the delay used for sleeping.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"update": schema.SingleNestedAttribute{
												Description:         "Update represents an update operation.",
												MarkdownDescription: "Update represents an update operation.",
												Attributes: map[string]schema.Attribute{
													"bindings": schema.ListNestedAttribute{
														Description:         "Bindings defines additional binding key/values.",
														MarkdownDescription: "Bindings defines additional binding key/values.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name the name of the binding.",
																	MarkdownDescription: "Name the name of the binding.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
																	},
																},

																"value": schema.MapAttribute{
																	Description:         "Value value of the binding.",
																	MarkdownDescription: "Value value of the binding.",
																	ElementType:         types.StringType,
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

													"cluster": schema.StringAttribute{
														Description:         "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														MarkdownDescription: "Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"clusters": schema.SingleNestedAttribute{
														Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
														MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Description:         "Context is the name of the context to use.",
																MarkdownDescription: "Context is the name of the context to use.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubeconfig": schema.StringAttribute{
																Description:         "Kubeconfig is the path to the referenced file.",
																MarkdownDescription: "Kubeconfig is the path to the referenced file.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"dry_run": schema.BoolAttribute{
														Description:         "DryRun determines whether the file should be applied in dry run mode.",
														MarkdownDescription: "DryRun determines whether the file should be applied in dry run mode.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"expect": schema.ListNestedAttribute{
														Description:         "Expect defines a list of matched checks to validate the operation outcome.",
														MarkdownDescription: "Expect defines a list of matched checks to validate the operation outcome.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"check": schema.MapAttribute{
																	Description:         "Check defines the verification statement.",
																	MarkdownDescription: "Check defines the verification statement.",
																	ElementType:         types.StringType,
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"match": schema.MapAttribute{
																	Description:         "Match defines the matching statement.",
																	MarkdownDescription: "Match defines the matching statement.",
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

													"file": schema.StringAttribute{
														Description:         "File is the path to the referenced file. This can be a direct path to a fileor an expression that matches multiple files, such as 'manifest/*.yaml' for all YAMLfiles within the 'manifest' directory.",
														MarkdownDescription: "File is the path to the referenced file. This can be a direct path to a fileor an expression that matches multiple files, such as 'manifest/*.yaml' for all YAMLfiles within the 'manifest' directory.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"outputs": schema.ListNestedAttribute{
														Description:         "Outputs defines output bindings.",
														MarkdownDescription: "Outputs defines output bindings.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"match": schema.MapAttribute{
																	Description:         "Match defines the matching statement.",
																	MarkdownDescription: "Match defines the matching statement.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name the name of the binding.",
																	MarkdownDescription: "Name the name of the binding.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\w+|\(.+\))$`), ""),
																	},
																},

																"value": schema.MapAttribute{
																	Description:         "Value value of the binding.",
																	MarkdownDescription: "Value value of the binding.",
																	ElementType:         types.StringType,
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

													"resource": schema.MapAttribute{
														Description:         "Resource provides a resource to be applied.",
														MarkdownDescription: "Resource provides a resource to be applied.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"template": schema.BoolAttribute{
														Description:         "Template determines whether resources should be considered for templating.",
														MarkdownDescription: "Template determines whether resources should be considered for templating.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.StringAttribute{
														Description:         "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														MarkdownDescription: "Timeout for the operation. Overrides the global timeout set in the Configuration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"wait": schema.SingleNestedAttribute{
												Description:         "Wait determines the resource wait collector to execute.",
												MarkdownDescription: "Wait determines the resource wait collector to execute.",
												Attributes: map[string]schema.Attribute{
													"api_version": schema.StringAttribute{
														Description:         "API version of the referent.",
														MarkdownDescription: "API version of the referent.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"cluster": schema.StringAttribute{
														Description:         "Cluster defines the target cluster where the wait operation will be performed (default cluster will be used if not specified).",
														MarkdownDescription: "Cluster defines the target cluster where the wait operation will be performed (default cluster will be used if not specified).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"clusters": schema.SingleNestedAttribute{
														Description:         "Clusters holds a registry to clusters to support multi-cluster tests.",
														MarkdownDescription: "Clusters holds a registry to clusters to support multi-cluster tests.",
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Description:         "Context is the name of the context to use.",
																MarkdownDescription: "Context is the name of the context to use.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubeconfig": schema.StringAttribute{
																Description:         "Kubeconfig is the path to the referenced file.",
																MarkdownDescription: "Kubeconfig is the path to the referenced file.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"for": schema.SingleNestedAttribute{
														Description:         "For specifies the condition to wait for.",
														MarkdownDescription: "For specifies the condition to wait for.",
														Attributes: map[string]schema.Attribute{
															"condition": schema.SingleNestedAttribute{
																Description:         "Condition specifies the condition to wait for.",
																MarkdownDescription: "Condition specifies the condition to wait for.",
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "Name defines the specific condition to wait for, e.g., 'Available', 'Ready'.",
																		MarkdownDescription: "Name defines the specific condition to wait for, e.g., 'Available', 'Ready'.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
																		Description:         "Value defines the specific condition status to wait for, e.g., 'True', 'False'.",
																		MarkdownDescription: "Value defines the specific condition status to wait for, e.g., 'True', 'False'.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"deletion": schema.MapAttribute{
																Description:         "Deletion specifies parameters for waiting on a resource's deletion.",
																MarkdownDescription: "Deletion specifies parameters for waiting on a resource's deletion.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"json_path": schema.SingleNestedAttribute{
																Description:         "JsonPath specifies the json path condition to wait for.",
																MarkdownDescription: "JsonPath specifies the json path condition to wait for.",
																Attributes: map[string]schema.Attribute{
																	"path": schema.StringAttribute{
																		Description:         "Path defines the json path to wait for, e.g. '{.status.phase}'.",
																		MarkdownDescription: "Path defines the json path to wait for, e.g. '{.status.phase}'.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
																		Description:         "Value defines the expected value to wait for, e.g., 'Running'.",
																		MarkdownDescription: "Value defines the expected value to wait for, e.g., 'Running'.",
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
														Required: true,
														Optional: false,
														Computed: false,
													},

													"format": schema.StringAttribute{
														Description:         "Format determines the output format (json or yaml).",
														MarkdownDescription: "Format determines the output format (json or yaml).",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.RegexMatches(regexp.MustCompile(`^(?:json|yaml|\(.+\))$`), ""),
														},
													},

													"kind": schema.StringAttribute{
														Description:         "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
														MarkdownDescription: "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
														Required:            true,
														Optional:            false,
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

													"selector": schema.StringAttribute{
														Description:         "Selector defines labels selector.",
														MarkdownDescription: "Selector defines labels selector.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.StringAttribute{
														Description:         "Timeout for the operation. Specifies how long to wait for the condition to be met before timing out.",
														MarkdownDescription: "Timeout for the operation. Specifies how long to wait for the condition to be met before timing out.",
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
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"template": schema.BoolAttribute{
						Description:         "Template determines whether resources should be considered for templating.",
						MarkdownDescription: "Template determines whether resources should be considered for templating.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"timeouts": schema.SingleNestedAttribute{
						Description:         "Timeouts for the test. Overrides the global timeouts set in the Configuration on a per operation basis.",
						MarkdownDescription: "Timeouts for the test. Overrides the global timeouts set in the Configuration on a per operation basis.",
						Attributes: map[string]schema.Attribute{
							"apply": schema.StringAttribute{
								Description:         "Apply defines the timeout for the apply operation",
								MarkdownDescription: "Apply defines the timeout for the apply operation",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"assert": schema.StringAttribute{
								Description:         "Assert defines the timeout for the assert operation",
								MarkdownDescription: "Assert defines the timeout for the assert operation",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cleanup": schema.StringAttribute{
								Description:         "Cleanup defines the timeout for the cleanup operation",
								MarkdownDescription: "Cleanup defines the timeout for the cleanup operation",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"delete": schema.StringAttribute{
								Description:         "Delete defines the timeout for the delete operation",
								MarkdownDescription: "Delete defines the timeout for the delete operation",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"error": schema.StringAttribute{
								Description:         "Error defines the timeout for the error operation",
								MarkdownDescription: "Error defines the timeout for the error operation",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"exec": schema.StringAttribute{
								Description:         "Exec defines the timeout for exec operations",
								MarkdownDescription: "Exec defines the timeout for exec operations",
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
	}
}

func (r *ChainsawKyvernoIoTestV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_chainsaw_kyverno_io_test_v1alpha1_manifest")

	var model ChainsawKyvernoIoTestV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("chainsaw.kyverno.io/v1alpha1")
	model.Kind = pointer.String("Test")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
