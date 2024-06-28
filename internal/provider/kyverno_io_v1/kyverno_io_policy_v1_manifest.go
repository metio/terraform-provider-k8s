/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kyverno_io_v1

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
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &KyvernoIoPolicyV1Manifest{}
)

func NewKyvernoIoPolicyV1Manifest() datasource.DataSource {
	return &KyvernoIoPolicyV1Manifest{}
}

type KyvernoIoPolicyV1Manifest struct{}

type KyvernoIoPolicyV1ManifestData struct {
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
		Admission                      *bool   `tfsdk:"admission" json:"admission,omitempty"`
		ApplyRules                     *string `tfsdk:"apply_rules" json:"applyRules,omitempty"`
		Background                     *bool   `tfsdk:"background" json:"background,omitempty"`
		FailurePolicy                  *string `tfsdk:"failure_policy" json:"failurePolicy,omitempty"`
		GenerateExisting               *bool   `tfsdk:"generate_existing" json:"generateExisting,omitempty"`
		GenerateExistingOnPolicyUpdate *bool   `tfsdk:"generate_existing_on_policy_update" json:"generateExistingOnPolicyUpdate,omitempty"`
		MutateExistingOnPolicyUpdate   *bool   `tfsdk:"mutate_existing_on_policy_update" json:"mutateExistingOnPolicyUpdate,omitempty"`
		Rules                          *[]struct {
			CelPreconditions *[]struct {
				Expression *string `tfsdk:"expression" json:"expression,omitempty"`
				Name       *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"cel_preconditions" json:"celPreconditions,omitempty"`
			Context *[]struct {
				ApiCall *struct {
					Data *[]struct {
						Key   *string            `tfsdk:"key" json:"key,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"data" json:"data,omitempty"`
					JmesPath *string `tfsdk:"jmes_path" json:"jmesPath,omitempty"`
					Method   *string `tfsdk:"method" json:"method,omitempty"`
					Service  *struct {
						CaBundle *string `tfsdk:"ca_bundle" json:"caBundle,omitempty"`
						Url      *string `tfsdk:"url" json:"url,omitempty"`
					} `tfsdk:"service" json:"service,omitempty"`
					UrlPath *string `tfsdk:"url_path" json:"urlPath,omitempty"`
				} `tfsdk:"api_call" json:"apiCall,omitempty"`
				ConfigMap *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"config_map" json:"configMap,omitempty"`
				GlobalReference *struct {
					JmesPath *string `tfsdk:"jmes_path" json:"jmesPath,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"global_reference" json:"globalReference,omitempty"`
				ImageRegistry *struct {
					ImageRegistryCredentials *struct {
						AllowInsecureRegistry *bool     `tfsdk:"allow_insecure_registry" json:"allowInsecureRegistry,omitempty"`
						Providers             *[]string `tfsdk:"providers" json:"providers,omitempty"`
						Secrets               *[]string `tfsdk:"secrets" json:"secrets,omitempty"`
					} `tfsdk:"image_registry_credentials" json:"imageRegistryCredentials,omitempty"`
					JmesPath  *string `tfsdk:"jmes_path" json:"jmesPath,omitempty"`
					Reference *string `tfsdk:"reference" json:"reference,omitempty"`
				} `tfsdk:"image_registry" json:"imageRegistry,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Variable *struct {
					Default  *map[string]string `tfsdk:"default" json:"default,omitempty"`
					JmesPath *string            `tfsdk:"jmes_path" json:"jmesPath,omitempty"`
					Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"variable" json:"variable,omitempty"`
			} `tfsdk:"context" json:"context,omitempty"`
			Exclude *struct {
				All *[]struct {
					ClusterRoles *[]string `tfsdk:"cluster_roles" json:"clusterRoles,omitempty"`
					Resources    *struct {
						Annotations       *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Kinds             *[]string          `tfsdk:"kinds" json:"kinds,omitempty"`
						Name              *string            `tfsdk:"name" json:"name,omitempty"`
						Names             *[]string          `tfsdk:"names" json:"names,omitempty"`
						NamespaceSelector *struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
						} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
						Namespaces *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
						Operations *[]string `tfsdk:"operations" json:"operations,omitempty"`
						Selector   *struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
						} `tfsdk:"selector" json:"selector,omitempty"`
					} `tfsdk:"resources" json:"resources,omitempty"`
					Roles    *[]string `tfsdk:"roles" json:"roles,omitempty"`
					Subjects *[]struct {
						ApiGroup  *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
						Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"subjects" json:"subjects,omitempty"`
				} `tfsdk:"all" json:"all,omitempty"`
				Any *[]struct {
					ClusterRoles *[]string `tfsdk:"cluster_roles" json:"clusterRoles,omitempty"`
					Resources    *struct {
						Annotations       *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Kinds             *[]string          `tfsdk:"kinds" json:"kinds,omitempty"`
						Name              *string            `tfsdk:"name" json:"name,omitempty"`
						Names             *[]string          `tfsdk:"names" json:"names,omitempty"`
						NamespaceSelector *struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
						} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
						Namespaces *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
						Operations *[]string `tfsdk:"operations" json:"operations,omitempty"`
						Selector   *struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
						} `tfsdk:"selector" json:"selector,omitempty"`
					} `tfsdk:"resources" json:"resources,omitempty"`
					Roles    *[]string `tfsdk:"roles" json:"roles,omitempty"`
					Subjects *[]struct {
						ApiGroup  *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
						Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"subjects" json:"subjects,omitempty"`
				} `tfsdk:"any" json:"any,omitempty"`
				ClusterRoles *[]string `tfsdk:"cluster_roles" json:"clusterRoles,omitempty"`
				Resources    *struct {
					Annotations       *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Kinds             *[]string          `tfsdk:"kinds" json:"kinds,omitempty"`
					Name              *string            `tfsdk:"name" json:"name,omitempty"`
					Names             *[]string          `tfsdk:"names" json:"names,omitempty"`
					NamespaceSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
					Namespaces *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
					Operations *[]string `tfsdk:"operations" json:"operations,omitempty"`
					Selector   *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"selector" json:"selector,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				Roles    *[]string `tfsdk:"roles" json:"roles,omitempty"`
				Subjects *[]struct {
					ApiGroup  *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
					Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"subjects" json:"subjects,omitempty"`
			} `tfsdk:"exclude" json:"exclude,omitempty"`
			Generate *struct {
				ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
				Clone      *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"clone" json:"clone,omitempty"`
				CloneList *struct {
					Kinds     *[]string `tfsdk:"kinds" json:"kinds,omitempty"`
					Namespace *string   `tfsdk:"namespace" json:"namespace,omitempty"`
					Selector  *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"selector" json:"selector,omitempty"`
				} `tfsdk:"clone_list" json:"cloneList,omitempty"`
				Data                           *map[string]string `tfsdk:"data" json:"data,omitempty"`
				GenerateExisting               *bool              `tfsdk:"generate_existing" json:"generateExisting,omitempty"`
				Kind                           *string            `tfsdk:"kind" json:"kind,omitempty"`
				Name                           *string            `tfsdk:"name" json:"name,omitempty"`
				Namespace                      *string            `tfsdk:"namespace" json:"namespace,omitempty"`
				OrphanDownstreamOnPolicyDelete *bool              `tfsdk:"orphan_downstream_on_policy_delete" json:"orphanDownstreamOnPolicyDelete,omitempty"`
				Synchronize                    *bool              `tfsdk:"synchronize" json:"synchronize,omitempty"`
				Uid                            *string            `tfsdk:"uid" json:"uid,omitempty"`
			} `tfsdk:"generate" json:"generate,omitempty"`
			ImageExtractors *map[string]string `tfsdk:"image_extractors" json:"imageExtractors,omitempty"`
			Match           *struct {
				All *[]struct {
					ClusterRoles *[]string `tfsdk:"cluster_roles" json:"clusterRoles,omitempty"`
					Resources    *struct {
						Annotations       *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Kinds             *[]string          `tfsdk:"kinds" json:"kinds,omitempty"`
						Name              *string            `tfsdk:"name" json:"name,omitempty"`
						Names             *[]string          `tfsdk:"names" json:"names,omitempty"`
						NamespaceSelector *struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
						} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
						Namespaces *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
						Operations *[]string `tfsdk:"operations" json:"operations,omitempty"`
						Selector   *struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
						} `tfsdk:"selector" json:"selector,omitempty"`
					} `tfsdk:"resources" json:"resources,omitempty"`
					Roles    *[]string `tfsdk:"roles" json:"roles,omitempty"`
					Subjects *[]struct {
						ApiGroup  *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
						Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"subjects" json:"subjects,omitempty"`
				} `tfsdk:"all" json:"all,omitempty"`
				Any *[]struct {
					ClusterRoles *[]string `tfsdk:"cluster_roles" json:"clusterRoles,omitempty"`
					Resources    *struct {
						Annotations       *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Kinds             *[]string          `tfsdk:"kinds" json:"kinds,omitempty"`
						Name              *string            `tfsdk:"name" json:"name,omitempty"`
						Names             *[]string          `tfsdk:"names" json:"names,omitempty"`
						NamespaceSelector *struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
						} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
						Namespaces *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
						Operations *[]string `tfsdk:"operations" json:"operations,omitempty"`
						Selector   *struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
						} `tfsdk:"selector" json:"selector,omitempty"`
					} `tfsdk:"resources" json:"resources,omitempty"`
					Roles    *[]string `tfsdk:"roles" json:"roles,omitempty"`
					Subjects *[]struct {
						ApiGroup  *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
						Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"subjects" json:"subjects,omitempty"`
				} `tfsdk:"any" json:"any,omitempty"`
				ClusterRoles *[]string `tfsdk:"cluster_roles" json:"clusterRoles,omitempty"`
				Resources    *struct {
					Annotations       *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Kinds             *[]string          `tfsdk:"kinds" json:"kinds,omitempty"`
					Name              *string            `tfsdk:"name" json:"name,omitempty"`
					Names             *[]string          `tfsdk:"names" json:"names,omitempty"`
					NamespaceSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
					Namespaces *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
					Operations *[]string `tfsdk:"operations" json:"operations,omitempty"`
					Selector   *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"selector" json:"selector,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				Roles    *[]string `tfsdk:"roles" json:"roles,omitempty"`
				Subjects *[]struct {
					ApiGroup  *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
					Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"subjects" json:"subjects,omitempty"`
			} `tfsdk:"match" json:"match,omitempty"`
			Mutate *struct {
				Foreach *[]struct {
					Context *[]struct {
						ApiCall *struct {
							Data *[]struct {
								Key   *string            `tfsdk:"key" json:"key,omitempty"`
								Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"data" json:"data,omitempty"`
							JmesPath *string `tfsdk:"jmes_path" json:"jmesPath,omitempty"`
							Method   *string `tfsdk:"method" json:"method,omitempty"`
							Service  *struct {
								CaBundle *string `tfsdk:"ca_bundle" json:"caBundle,omitempty"`
								Url      *string `tfsdk:"url" json:"url,omitempty"`
							} `tfsdk:"service" json:"service,omitempty"`
							UrlPath *string `tfsdk:"url_path" json:"urlPath,omitempty"`
						} `tfsdk:"api_call" json:"apiCall,omitempty"`
						ConfigMap *struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"config_map" json:"configMap,omitempty"`
						GlobalReference *struct {
							JmesPath *string `tfsdk:"jmes_path" json:"jmesPath,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"global_reference" json:"globalReference,omitempty"`
						ImageRegistry *struct {
							ImageRegistryCredentials *struct {
								AllowInsecureRegistry *bool     `tfsdk:"allow_insecure_registry" json:"allowInsecureRegistry,omitempty"`
								Providers             *[]string `tfsdk:"providers" json:"providers,omitempty"`
								Secrets               *[]string `tfsdk:"secrets" json:"secrets,omitempty"`
							} `tfsdk:"image_registry_credentials" json:"imageRegistryCredentials,omitempty"`
							JmesPath  *string `tfsdk:"jmes_path" json:"jmesPath,omitempty"`
							Reference *string `tfsdk:"reference" json:"reference,omitempty"`
						} `tfsdk:"image_registry" json:"imageRegistry,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Variable *struct {
							Default  *map[string]string `tfsdk:"default" json:"default,omitempty"`
							JmesPath *string            `tfsdk:"jmes_path" json:"jmesPath,omitempty"`
							Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"variable" json:"variable,omitempty"`
					} `tfsdk:"context" json:"context,omitempty"`
					Foreach             *map[string]string `tfsdk:"foreach" json:"foreach,omitempty"`
					List                *string            `tfsdk:"list" json:"list,omitempty"`
					Order               *string            `tfsdk:"order" json:"order,omitempty"`
					PatchStrategicMerge *map[string]string `tfsdk:"patch_strategic_merge" json:"patchStrategicMerge,omitempty"`
					PatchesJson6902     *string            `tfsdk:"patches_json6902" json:"patchesJson6902,omitempty"`
					Preconditions       *struct {
						All *[]struct {
							Key      *map[string]string `tfsdk:"key" json:"key,omitempty"`
							Message  *string            `tfsdk:"message" json:"message,omitempty"`
							Operator *string            `tfsdk:"operator" json:"operator,omitempty"`
							Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"all" json:"all,omitempty"`
						Any *[]struct {
							Key      *map[string]string `tfsdk:"key" json:"key,omitempty"`
							Message  *string            `tfsdk:"message" json:"message,omitempty"`
							Operator *string            `tfsdk:"operator" json:"operator,omitempty"`
							Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"any" json:"any,omitempty"`
					} `tfsdk:"preconditions" json:"preconditions,omitempty"`
				} `tfsdk:"foreach" json:"foreach,omitempty"`
				MutateExistingOnPolicyUpdate *bool              `tfsdk:"mutate_existing_on_policy_update" json:"mutateExistingOnPolicyUpdate,omitempty"`
				PatchStrategicMerge          *map[string]string `tfsdk:"patch_strategic_merge" json:"patchStrategicMerge,omitempty"`
				PatchesJson6902              *string            `tfsdk:"patches_json6902" json:"patchesJson6902,omitempty"`
				Targets                      *[]struct {
					ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
					Context    *[]struct {
						ApiCall *struct {
							Data *[]struct {
								Key   *string            `tfsdk:"key" json:"key,omitempty"`
								Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"data" json:"data,omitempty"`
							JmesPath *string `tfsdk:"jmes_path" json:"jmesPath,omitempty"`
							Method   *string `tfsdk:"method" json:"method,omitempty"`
							Service  *struct {
								CaBundle *string `tfsdk:"ca_bundle" json:"caBundle,omitempty"`
								Url      *string `tfsdk:"url" json:"url,omitempty"`
							} `tfsdk:"service" json:"service,omitempty"`
							UrlPath *string `tfsdk:"url_path" json:"urlPath,omitempty"`
						} `tfsdk:"api_call" json:"apiCall,omitempty"`
						ConfigMap *struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"config_map" json:"configMap,omitempty"`
						GlobalReference *struct {
							JmesPath *string `tfsdk:"jmes_path" json:"jmesPath,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"global_reference" json:"globalReference,omitempty"`
						ImageRegistry *struct {
							ImageRegistryCredentials *struct {
								AllowInsecureRegistry *bool     `tfsdk:"allow_insecure_registry" json:"allowInsecureRegistry,omitempty"`
								Providers             *[]string `tfsdk:"providers" json:"providers,omitempty"`
								Secrets               *[]string `tfsdk:"secrets" json:"secrets,omitempty"`
							} `tfsdk:"image_registry_credentials" json:"imageRegistryCredentials,omitempty"`
							JmesPath  *string `tfsdk:"jmes_path" json:"jmesPath,omitempty"`
							Reference *string `tfsdk:"reference" json:"reference,omitempty"`
						} `tfsdk:"image_registry" json:"imageRegistry,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Variable *struct {
							Default  *map[string]string `tfsdk:"default" json:"default,omitempty"`
							JmesPath *string            `tfsdk:"jmes_path" json:"jmesPath,omitempty"`
							Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"variable" json:"variable,omitempty"`
					} `tfsdk:"context" json:"context,omitempty"`
					Kind          *string            `tfsdk:"kind" json:"kind,omitempty"`
					Name          *string            `tfsdk:"name" json:"name,omitempty"`
					Namespace     *string            `tfsdk:"namespace" json:"namespace,omitempty"`
					Preconditions *map[string]string `tfsdk:"preconditions" json:"preconditions,omitempty"`
					Uid           *string            `tfsdk:"uid" json:"uid,omitempty"`
				} `tfsdk:"targets" json:"targets,omitempty"`
			} `tfsdk:"mutate" json:"mutate,omitempty"`
			Name                   *string            `tfsdk:"name" json:"name,omitempty"`
			Preconditions          *map[string]string `tfsdk:"preconditions" json:"preconditions,omitempty"`
			SkipBackgroundRequests *bool              `tfsdk:"skip_background_requests" json:"skipBackgroundRequests,omitempty"`
			Validate               *struct {
				AnyPattern *map[string]string `tfsdk:"any_pattern" json:"anyPattern,omitempty"`
				Cel        *struct {
					AuditAnnotations *[]struct {
						Key             *string `tfsdk:"key" json:"key,omitempty"`
						ValueExpression *string `tfsdk:"value_expression" json:"valueExpression,omitempty"`
					} `tfsdk:"audit_annotations" json:"auditAnnotations,omitempty"`
					Expressions *[]struct {
						Expression        *string `tfsdk:"expression" json:"expression,omitempty"`
						Message           *string `tfsdk:"message" json:"message,omitempty"`
						MessageExpression *string `tfsdk:"message_expression" json:"messageExpression,omitempty"`
						Reason            *string `tfsdk:"reason" json:"reason,omitempty"`
					} `tfsdk:"expressions" json:"expressions,omitempty"`
					ParamKind *struct {
						ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
						Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
					} `tfsdk:"param_kind" json:"paramKind,omitempty"`
					ParamRef *struct {
						Name                    *string `tfsdk:"name" json:"name,omitempty"`
						Namespace               *string `tfsdk:"namespace" json:"namespace,omitempty"`
						ParameterNotFoundAction *string `tfsdk:"parameter_not_found_action" json:"parameterNotFoundAction,omitempty"`
						Selector                *struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
						} `tfsdk:"selector" json:"selector,omitempty"`
					} `tfsdk:"param_ref" json:"paramRef,omitempty"`
					Variables *[]struct {
						Expression *string `tfsdk:"expression" json:"expression,omitempty"`
						Name       *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"variables" json:"variables,omitempty"`
				} `tfsdk:"cel" json:"cel,omitempty"`
				Deny *struct {
					Conditions *map[string]string `tfsdk:"conditions" json:"conditions,omitempty"`
				} `tfsdk:"deny" json:"deny,omitempty"`
				Foreach *[]struct {
					AnyPattern *map[string]string `tfsdk:"any_pattern" json:"anyPattern,omitempty"`
					Context    *[]struct {
						ApiCall *struct {
							Data *[]struct {
								Key   *string            `tfsdk:"key" json:"key,omitempty"`
								Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"data" json:"data,omitempty"`
							JmesPath *string `tfsdk:"jmes_path" json:"jmesPath,omitempty"`
							Method   *string `tfsdk:"method" json:"method,omitempty"`
							Service  *struct {
								CaBundle *string `tfsdk:"ca_bundle" json:"caBundle,omitempty"`
								Url      *string `tfsdk:"url" json:"url,omitempty"`
							} `tfsdk:"service" json:"service,omitempty"`
							UrlPath *string `tfsdk:"url_path" json:"urlPath,omitempty"`
						} `tfsdk:"api_call" json:"apiCall,omitempty"`
						ConfigMap *struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"config_map" json:"configMap,omitempty"`
						GlobalReference *struct {
							JmesPath *string `tfsdk:"jmes_path" json:"jmesPath,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"global_reference" json:"globalReference,omitempty"`
						ImageRegistry *struct {
							ImageRegistryCredentials *struct {
								AllowInsecureRegistry *bool     `tfsdk:"allow_insecure_registry" json:"allowInsecureRegistry,omitempty"`
								Providers             *[]string `tfsdk:"providers" json:"providers,omitempty"`
								Secrets               *[]string `tfsdk:"secrets" json:"secrets,omitempty"`
							} `tfsdk:"image_registry_credentials" json:"imageRegistryCredentials,omitempty"`
							JmesPath  *string `tfsdk:"jmes_path" json:"jmesPath,omitempty"`
							Reference *string `tfsdk:"reference" json:"reference,omitempty"`
						} `tfsdk:"image_registry" json:"imageRegistry,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Variable *struct {
							Default  *map[string]string `tfsdk:"default" json:"default,omitempty"`
							JmesPath *string            `tfsdk:"jmes_path" json:"jmesPath,omitempty"`
							Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"variable" json:"variable,omitempty"`
					} `tfsdk:"context" json:"context,omitempty"`
					Deny *struct {
						Conditions *map[string]string `tfsdk:"conditions" json:"conditions,omitempty"`
					} `tfsdk:"deny" json:"deny,omitempty"`
					ElementScope  *bool              `tfsdk:"element_scope" json:"elementScope,omitempty"`
					Foreach       *map[string]string `tfsdk:"foreach" json:"foreach,omitempty"`
					List          *string            `tfsdk:"list" json:"list,omitempty"`
					Pattern       *map[string]string `tfsdk:"pattern" json:"pattern,omitempty"`
					Preconditions *struct {
						All *[]struct {
							Key      *map[string]string `tfsdk:"key" json:"key,omitempty"`
							Message  *string            `tfsdk:"message" json:"message,omitempty"`
							Operator *string            `tfsdk:"operator" json:"operator,omitempty"`
							Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"all" json:"all,omitempty"`
						Any *[]struct {
							Key      *map[string]string `tfsdk:"key" json:"key,omitempty"`
							Message  *string            `tfsdk:"message" json:"message,omitempty"`
							Operator *string            `tfsdk:"operator" json:"operator,omitempty"`
							Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"any" json:"any,omitempty"`
					} `tfsdk:"preconditions" json:"preconditions,omitempty"`
				} `tfsdk:"foreach" json:"foreach,omitempty"`
				Manifests *struct {
					AnnotationDomain *string `tfsdk:"annotation_domain" json:"annotationDomain,omitempty"`
					Attestors        *[]struct {
						Count   *int64 `tfsdk:"count" json:"count,omitempty"`
						Entries *[]struct {
							Annotations  *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
							Attestor     *map[string]string `tfsdk:"attestor" json:"attestor,omitempty"`
							Certificates *struct {
								Cert      *string `tfsdk:"cert" json:"cert,omitempty"`
								CertChain *string `tfsdk:"cert_chain" json:"certChain,omitempty"`
								Ctlog     *struct {
									IgnoreSCT    *bool   `tfsdk:"ignore_sct" json:"ignoreSCT,omitempty"`
									Pubkey       *string `tfsdk:"pubkey" json:"pubkey,omitempty"`
									TsaCertChain *string `tfsdk:"tsa_cert_chain" json:"tsaCertChain,omitempty"`
								} `tfsdk:"ctlog" json:"ctlog,omitempty"`
								Rekor *struct {
									IgnoreTlog *bool   `tfsdk:"ignore_tlog" json:"ignoreTlog,omitempty"`
									Pubkey     *string `tfsdk:"pubkey" json:"pubkey,omitempty"`
									Url        *string `tfsdk:"url" json:"url,omitempty"`
								} `tfsdk:"rekor" json:"rekor,omitempty"`
							} `tfsdk:"certificates" json:"certificates,omitempty"`
							Keyless *struct {
								AdditionalExtensions *map[string]string `tfsdk:"additional_extensions" json:"additionalExtensions,omitempty"`
								Ctlog                *struct {
									IgnoreSCT    *bool   `tfsdk:"ignore_sct" json:"ignoreSCT,omitempty"`
									Pubkey       *string `tfsdk:"pubkey" json:"pubkey,omitempty"`
									TsaCertChain *string `tfsdk:"tsa_cert_chain" json:"tsaCertChain,omitempty"`
								} `tfsdk:"ctlog" json:"ctlog,omitempty"`
								Issuer *string `tfsdk:"issuer" json:"issuer,omitempty"`
								Rekor  *struct {
									IgnoreTlog *bool   `tfsdk:"ignore_tlog" json:"ignoreTlog,omitempty"`
									Pubkey     *string `tfsdk:"pubkey" json:"pubkey,omitempty"`
									Url        *string `tfsdk:"url" json:"url,omitempty"`
								} `tfsdk:"rekor" json:"rekor,omitempty"`
								Roots   *string `tfsdk:"roots" json:"roots,omitempty"`
								Subject *string `tfsdk:"subject" json:"subject,omitempty"`
							} `tfsdk:"keyless" json:"keyless,omitempty"`
							Keys *struct {
								Ctlog *struct {
									IgnoreSCT    *bool   `tfsdk:"ignore_sct" json:"ignoreSCT,omitempty"`
									Pubkey       *string `tfsdk:"pubkey" json:"pubkey,omitempty"`
									TsaCertChain *string `tfsdk:"tsa_cert_chain" json:"tsaCertChain,omitempty"`
								} `tfsdk:"ctlog" json:"ctlog,omitempty"`
								Kms        *string `tfsdk:"kms" json:"kms,omitempty"`
								PublicKeys *string `tfsdk:"public_keys" json:"publicKeys,omitempty"`
								Rekor      *struct {
									IgnoreTlog *bool   `tfsdk:"ignore_tlog" json:"ignoreTlog,omitempty"`
									Pubkey     *string `tfsdk:"pubkey" json:"pubkey,omitempty"`
									Url        *string `tfsdk:"url" json:"url,omitempty"`
								} `tfsdk:"rekor" json:"rekor,omitempty"`
								Secret *struct {
									Name      *string `tfsdk:"name" json:"name,omitempty"`
									Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
								} `tfsdk:"secret" json:"secret,omitempty"`
								SignatureAlgorithm *string `tfsdk:"signature_algorithm" json:"signatureAlgorithm,omitempty"`
							} `tfsdk:"keys" json:"keys,omitempty"`
							Repository *string `tfsdk:"repository" json:"repository,omitempty"`
						} `tfsdk:"entries" json:"entries,omitempty"`
					} `tfsdk:"attestors" json:"attestors,omitempty"`
					DryRun *struct {
						Enable    *bool   `tfsdk:"enable" json:"enable,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"dry_run" json:"dryRun,omitempty"`
					IgnoreFields *[]struct {
						Fields  *[]string `tfsdk:"fields" json:"fields,omitempty"`
						Objects *[]struct {
							Group     *string `tfsdk:"group" json:"group,omitempty"`
							Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
							Version   *string `tfsdk:"version" json:"version,omitempty"`
						} `tfsdk:"objects" json:"objects,omitempty"`
					} `tfsdk:"ignore_fields" json:"ignoreFields,omitempty"`
					Repository *string `tfsdk:"repository" json:"repository,omitempty"`
				} `tfsdk:"manifests" json:"manifests,omitempty"`
				Message     *string            `tfsdk:"message" json:"message,omitempty"`
				Pattern     *map[string]string `tfsdk:"pattern" json:"pattern,omitempty"`
				PodSecurity *struct {
					Exclude *[]struct {
						ControlName     *string   `tfsdk:"control_name" json:"controlName,omitempty"`
						Images          *[]string `tfsdk:"images" json:"images,omitempty"`
						RestrictedField *string   `tfsdk:"restricted_field" json:"restrictedField,omitempty"`
						Values          *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"exclude" json:"exclude,omitempty"`
					Level   *string `tfsdk:"level" json:"level,omitempty"`
					Version *string `tfsdk:"version" json:"version,omitempty"`
				} `tfsdk:"pod_security" json:"podSecurity,omitempty"`
				ValidationFailureAction          *string `tfsdk:"validation_failure_action" json:"validationFailureAction,omitempty"`
				ValidationFailureActionOverrides *[]struct {
					Action            *string `tfsdk:"action" json:"action,omitempty"`
					NamespaceSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
					Namespaces *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
				} `tfsdk:"validation_failure_action_overrides" json:"validationFailureActionOverrides,omitempty"`
			} `tfsdk:"validate" json:"validate,omitempty"`
			VerifyImages *[]struct {
				AdditionalExtensions *map[string]string `tfsdk:"additional_extensions" json:"additionalExtensions,omitempty"`
				Annotations          *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Attestations         *[]struct {
					Attestors *[]struct {
						Count   *int64 `tfsdk:"count" json:"count,omitempty"`
						Entries *[]struct {
							Annotations  *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
							Attestor     *map[string]string `tfsdk:"attestor" json:"attestor,omitempty"`
							Certificates *struct {
								Cert      *string `tfsdk:"cert" json:"cert,omitempty"`
								CertChain *string `tfsdk:"cert_chain" json:"certChain,omitempty"`
								Ctlog     *struct {
									IgnoreSCT    *bool   `tfsdk:"ignore_sct" json:"ignoreSCT,omitempty"`
									Pubkey       *string `tfsdk:"pubkey" json:"pubkey,omitempty"`
									TsaCertChain *string `tfsdk:"tsa_cert_chain" json:"tsaCertChain,omitempty"`
								} `tfsdk:"ctlog" json:"ctlog,omitempty"`
								Rekor *struct {
									IgnoreTlog *bool   `tfsdk:"ignore_tlog" json:"ignoreTlog,omitempty"`
									Pubkey     *string `tfsdk:"pubkey" json:"pubkey,omitempty"`
									Url        *string `tfsdk:"url" json:"url,omitempty"`
								} `tfsdk:"rekor" json:"rekor,omitempty"`
							} `tfsdk:"certificates" json:"certificates,omitempty"`
							Keyless *struct {
								AdditionalExtensions *map[string]string `tfsdk:"additional_extensions" json:"additionalExtensions,omitempty"`
								Ctlog                *struct {
									IgnoreSCT    *bool   `tfsdk:"ignore_sct" json:"ignoreSCT,omitempty"`
									Pubkey       *string `tfsdk:"pubkey" json:"pubkey,omitempty"`
									TsaCertChain *string `tfsdk:"tsa_cert_chain" json:"tsaCertChain,omitempty"`
								} `tfsdk:"ctlog" json:"ctlog,omitempty"`
								Issuer *string `tfsdk:"issuer" json:"issuer,omitempty"`
								Rekor  *struct {
									IgnoreTlog *bool   `tfsdk:"ignore_tlog" json:"ignoreTlog,omitempty"`
									Pubkey     *string `tfsdk:"pubkey" json:"pubkey,omitempty"`
									Url        *string `tfsdk:"url" json:"url,omitempty"`
								} `tfsdk:"rekor" json:"rekor,omitempty"`
								Roots   *string `tfsdk:"roots" json:"roots,omitempty"`
								Subject *string `tfsdk:"subject" json:"subject,omitempty"`
							} `tfsdk:"keyless" json:"keyless,omitempty"`
							Keys *struct {
								Ctlog *struct {
									IgnoreSCT    *bool   `tfsdk:"ignore_sct" json:"ignoreSCT,omitempty"`
									Pubkey       *string `tfsdk:"pubkey" json:"pubkey,omitempty"`
									TsaCertChain *string `tfsdk:"tsa_cert_chain" json:"tsaCertChain,omitempty"`
								} `tfsdk:"ctlog" json:"ctlog,omitempty"`
								Kms        *string `tfsdk:"kms" json:"kms,omitempty"`
								PublicKeys *string `tfsdk:"public_keys" json:"publicKeys,omitempty"`
								Rekor      *struct {
									IgnoreTlog *bool   `tfsdk:"ignore_tlog" json:"ignoreTlog,omitempty"`
									Pubkey     *string `tfsdk:"pubkey" json:"pubkey,omitempty"`
									Url        *string `tfsdk:"url" json:"url,omitempty"`
								} `tfsdk:"rekor" json:"rekor,omitempty"`
								Secret *struct {
									Name      *string `tfsdk:"name" json:"name,omitempty"`
									Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
								} `tfsdk:"secret" json:"secret,omitempty"`
								SignatureAlgorithm *string `tfsdk:"signature_algorithm" json:"signatureAlgorithm,omitempty"`
							} `tfsdk:"keys" json:"keys,omitempty"`
							Repository *string `tfsdk:"repository" json:"repository,omitempty"`
						} `tfsdk:"entries" json:"entries,omitempty"`
					} `tfsdk:"attestors" json:"attestors,omitempty"`
					Conditions *[]struct {
						All *[]struct {
							Key      *map[string]string `tfsdk:"key" json:"key,omitempty"`
							Message  *string            `tfsdk:"message" json:"message,omitempty"`
							Operator *string            `tfsdk:"operator" json:"operator,omitempty"`
							Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"all" json:"all,omitempty"`
						Any *[]struct {
							Key      *map[string]string `tfsdk:"key" json:"key,omitempty"`
							Message  *string            `tfsdk:"message" json:"message,omitempty"`
							Operator *string            `tfsdk:"operator" json:"operator,omitempty"`
							Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"any" json:"any,omitempty"`
					} `tfsdk:"conditions" json:"conditions,omitempty"`
					PredicateType *string `tfsdk:"predicate_type" json:"predicateType,omitempty"`
					Type          *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"attestations" json:"attestations,omitempty"`
				Attestors *[]struct {
					Count   *int64 `tfsdk:"count" json:"count,omitempty"`
					Entries *[]struct {
						Annotations  *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Attestor     *map[string]string `tfsdk:"attestor" json:"attestor,omitempty"`
						Certificates *struct {
							Cert      *string `tfsdk:"cert" json:"cert,omitempty"`
							CertChain *string `tfsdk:"cert_chain" json:"certChain,omitempty"`
							Ctlog     *struct {
								IgnoreSCT    *bool   `tfsdk:"ignore_sct" json:"ignoreSCT,omitempty"`
								Pubkey       *string `tfsdk:"pubkey" json:"pubkey,omitempty"`
								TsaCertChain *string `tfsdk:"tsa_cert_chain" json:"tsaCertChain,omitempty"`
							} `tfsdk:"ctlog" json:"ctlog,omitempty"`
							Rekor *struct {
								IgnoreTlog *bool   `tfsdk:"ignore_tlog" json:"ignoreTlog,omitempty"`
								Pubkey     *string `tfsdk:"pubkey" json:"pubkey,omitempty"`
								Url        *string `tfsdk:"url" json:"url,omitempty"`
							} `tfsdk:"rekor" json:"rekor,omitempty"`
						} `tfsdk:"certificates" json:"certificates,omitempty"`
						Keyless *struct {
							AdditionalExtensions *map[string]string `tfsdk:"additional_extensions" json:"additionalExtensions,omitempty"`
							Ctlog                *struct {
								IgnoreSCT    *bool   `tfsdk:"ignore_sct" json:"ignoreSCT,omitempty"`
								Pubkey       *string `tfsdk:"pubkey" json:"pubkey,omitempty"`
								TsaCertChain *string `tfsdk:"tsa_cert_chain" json:"tsaCertChain,omitempty"`
							} `tfsdk:"ctlog" json:"ctlog,omitempty"`
							Issuer *string `tfsdk:"issuer" json:"issuer,omitempty"`
							Rekor  *struct {
								IgnoreTlog *bool   `tfsdk:"ignore_tlog" json:"ignoreTlog,omitempty"`
								Pubkey     *string `tfsdk:"pubkey" json:"pubkey,omitempty"`
								Url        *string `tfsdk:"url" json:"url,omitempty"`
							} `tfsdk:"rekor" json:"rekor,omitempty"`
							Roots   *string `tfsdk:"roots" json:"roots,omitempty"`
							Subject *string `tfsdk:"subject" json:"subject,omitempty"`
						} `tfsdk:"keyless" json:"keyless,omitempty"`
						Keys *struct {
							Ctlog *struct {
								IgnoreSCT    *bool   `tfsdk:"ignore_sct" json:"ignoreSCT,omitempty"`
								Pubkey       *string `tfsdk:"pubkey" json:"pubkey,omitempty"`
								TsaCertChain *string `tfsdk:"tsa_cert_chain" json:"tsaCertChain,omitempty"`
							} `tfsdk:"ctlog" json:"ctlog,omitempty"`
							Kms        *string `tfsdk:"kms" json:"kms,omitempty"`
							PublicKeys *string `tfsdk:"public_keys" json:"publicKeys,omitempty"`
							Rekor      *struct {
								IgnoreTlog *bool   `tfsdk:"ignore_tlog" json:"ignoreTlog,omitempty"`
								Pubkey     *string `tfsdk:"pubkey" json:"pubkey,omitempty"`
								Url        *string `tfsdk:"url" json:"url,omitempty"`
							} `tfsdk:"rekor" json:"rekor,omitempty"`
							Secret *struct {
								Name      *string `tfsdk:"name" json:"name,omitempty"`
								Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
							} `tfsdk:"secret" json:"secret,omitempty"`
							SignatureAlgorithm *string `tfsdk:"signature_algorithm" json:"signatureAlgorithm,omitempty"`
						} `tfsdk:"keys" json:"keys,omitempty"`
						Repository *string `tfsdk:"repository" json:"repository,omitempty"`
					} `tfsdk:"entries" json:"entries,omitempty"`
				} `tfsdk:"attestors" json:"attestors,omitempty"`
				CosignOCI11              *bool     `tfsdk:"cosign_oci11" json:"cosignOCI11,omitempty"`
				Image                    *string   `tfsdk:"image" json:"image,omitempty"`
				ImageReferences          *[]string `tfsdk:"image_references" json:"imageReferences,omitempty"`
				ImageRegistryCredentials *struct {
					AllowInsecureRegistry *bool     `tfsdk:"allow_insecure_registry" json:"allowInsecureRegistry,omitempty"`
					Providers             *[]string `tfsdk:"providers" json:"providers,omitempty"`
					Secrets               *[]string `tfsdk:"secrets" json:"secrets,omitempty"`
				} `tfsdk:"image_registry_credentials" json:"imageRegistryCredentials,omitempty"`
				Issuer              *string   `tfsdk:"issuer" json:"issuer,omitempty"`
				Key                 *string   `tfsdk:"key" json:"key,omitempty"`
				MutateDigest        *bool     `tfsdk:"mutate_digest" json:"mutateDigest,omitempty"`
				Repository          *string   `tfsdk:"repository" json:"repository,omitempty"`
				Required            *bool     `tfsdk:"required" json:"required,omitempty"`
				Roots               *string   `tfsdk:"roots" json:"roots,omitempty"`
				SkipImageReferences *[]string `tfsdk:"skip_image_references" json:"skipImageReferences,omitempty"`
				Subject             *string   `tfsdk:"subject" json:"subject,omitempty"`
				Type                *string   `tfsdk:"type" json:"type,omitempty"`
				UseCache            *bool     `tfsdk:"use_cache" json:"useCache,omitempty"`
				VerifyDigest        *bool     `tfsdk:"verify_digest" json:"verifyDigest,omitempty"`
			} `tfsdk:"verify_images" json:"verifyImages,omitempty"`
		} `tfsdk:"rules" json:"rules,omitempty"`
		SchemaValidation                 *bool   `tfsdk:"schema_validation" json:"schemaValidation,omitempty"`
		UseServerSideApply               *bool   `tfsdk:"use_server_side_apply" json:"useServerSideApply,omitempty"`
		ValidationFailureAction          *string `tfsdk:"validation_failure_action" json:"validationFailureAction,omitempty"`
		ValidationFailureActionOverrides *[]struct {
			Action            *string `tfsdk:"action" json:"action,omitempty"`
			NamespaceSelector *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
			Namespaces *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
		} `tfsdk:"validation_failure_action_overrides" json:"validationFailureActionOverrides,omitempty"`
		WebhookConfiguration *struct {
			FailurePolicy   *string `tfsdk:"failure_policy" json:"failurePolicy,omitempty"`
			MatchConditions *[]struct {
				Expression *string `tfsdk:"expression" json:"expression,omitempty"`
				Name       *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"match_conditions" json:"matchConditions,omitempty"`
			TimeoutSeconds *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
		} `tfsdk:"webhook_configuration" json:"webhookConfiguration,omitempty"`
		WebhookTimeoutSeconds *int64 `tfsdk:"webhook_timeout_seconds" json:"webhookTimeoutSeconds,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KyvernoIoPolicyV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kyverno_io_policy_v1_manifest"
}

func (r *KyvernoIoPolicyV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Policy declares validation, mutation, and generation behaviors for matching resources.See: https://kyverno.io/docs/writing-policies/ for more information.",
		MarkdownDescription: "Policy declares validation, mutation, and generation behaviors for matching resources.See: https://kyverno.io/docs/writing-policies/ for more information.",
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
				Description:         "Spec defines policy behaviors and contains one or more rules.",
				MarkdownDescription: "Spec defines policy behaviors and contains one or more rules.",
				Attributes: map[string]schema.Attribute{
					"admission": schema.BoolAttribute{
						Description:         "Admission controls if rules are applied during admission.Optional. Default value is 'true'.",
						MarkdownDescription: "Admission controls if rules are applied during admission.Optional. Default value is 'true'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"apply_rules": schema.StringAttribute{
						Description:         "ApplyRules controls how rules in a policy are applied. Rule are processed inthe order of declaration. When set to 'One' processing stops after a rule hasbeen applied i.e. the rule matches and results in a pass, fail, or error. Whenset to 'All' all rules in the policy are processed. The default is 'All'.",
						MarkdownDescription: "ApplyRules controls how rules in a policy are applied. Rule are processed inthe order of declaration. When set to 'One' processing stops after a rule hasbeen applied i.e. the rule matches and results in a pass, fail, or error. Whenset to 'All' all rules in the policy are processed. The default is 'All'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("All", "One"),
						},
					},

					"background": schema.BoolAttribute{
						Description:         "Background controls if rules are applied to existing resources during a background scan.Optional. Default value is 'true'. The value must be set to 'false' if the policy ruleuses variables that are only available in the admission review request (e.g. user name).",
						MarkdownDescription: "Background controls if rules are applied to existing resources during a background scan.Optional. Default value is 'true'. The value must be set to 'false' if the policy ruleuses variables that are only available in the admission review request (e.g. user name).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"failure_policy": schema.StringAttribute{
						Description:         "Deprecated, use failurePolicy under the webhookConfiguration instead.",
						MarkdownDescription: "Deprecated, use failurePolicy under the webhookConfiguration instead.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Ignore", "Fail"),
						},
					},

					"generate_existing": schema.BoolAttribute{
						Description:         "Deprecated, use generateExisting under the generate rule instead",
						MarkdownDescription: "Deprecated, use generateExisting under the generate rule instead",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"generate_existing_on_policy_update": schema.BoolAttribute{
						Description:         "Deprecated, use generateExisting instead",
						MarkdownDescription: "Deprecated, use generateExisting instead",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"mutate_existing_on_policy_update": schema.BoolAttribute{
						Description:         "Deprecated, use mutateExistingOnPolicyUpdate under the mutate rule instead",
						MarkdownDescription: "Deprecated, use mutateExistingOnPolicyUpdate under the mutate rule instead",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"rules": schema.ListNestedAttribute{
						Description:         "Rules is a list of Rule instances. A Policy contains multiple rules andeach rule can validate, mutate, or generate resources.",
						MarkdownDescription: "Rules is a list of Rule instances. A Policy contains multiple rules andeach rule can validate, mutate, or generate resources.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"cel_preconditions": schema.ListNestedAttribute{
									Description:         "CELPreconditions are used to determine if a policy rule should be applied by evaluating aset of CEL conditions. It can only be used with the validate.cel subrule",
									MarkdownDescription: "CELPreconditions are used to determine if a policy rule should be applied by evaluating aset of CEL conditions. It can only be used with the validate.cel subrule",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"expression": schema.StringAttribute{
												Description:         "Expression represents the expression which will be evaluated by CEL. Must evaluate to bool.CEL expressions have access to the contents of the AdmissionRequest and Authorizer, organized into CEL variables:'object' - The object from the incoming request. The value is null for DELETE requests.'oldObject' - The existing object. The value is null for CREATE requests.'request' - Attributes of the admission request(/pkg/apis/admission/types.go#AdmissionRequest).'authorizer' - A CEL Authorizer. May be used to perform authorization checks for the principal (user or service account) of the request.  See https://pkg.go.dev/k8s.io/apiserver/pkg/cel/library#Authz'authorizer.requestResource' - A CEL ResourceCheck constructed from the 'authorizer' and configured with the  request resource.Documentation on CEL: https://kubernetes.io/docs/reference/using-api/cel/Required.",
												MarkdownDescription: "Expression represents the expression which will be evaluated by CEL. Must evaluate to bool.CEL expressions have access to the contents of the AdmissionRequest and Authorizer, organized into CEL variables:'object' - The object from the incoming request. The value is null for DELETE requests.'oldObject' - The existing object. The value is null for CREATE requests.'request' - Attributes of the admission request(/pkg/apis/admission/types.go#AdmissionRequest).'authorizer' - A CEL Authorizer. May be used to perform authorization checks for the principal (user or service account) of the request.  See https://pkg.go.dev/k8s.io/apiserver/pkg/cel/library#Authz'authorizer.requestResource' - A CEL ResourceCheck constructed from the 'authorizer' and configured with the  request resource.Documentation on CEL: https://kubernetes.io/docs/reference/using-api/cel/Required.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name is an identifier for this match condition, used for strategic merging of MatchConditions,as well as providing an identifier for logging purposes. A good name should be descriptive ofthe associated expression.Name must be a qualified name consisting of alphanumeric characters, '-', '_' or '.', andmust start and end with an alphanumeric character (e.g. 'MyName',  or 'my.name',  or'123-abc', regex used for validation is '([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9]') with anoptional DNS subdomain prefix and '/' (e.g. 'example.com/MyName')Required.",
												MarkdownDescription: "Name is an identifier for this match condition, used for strategic merging of MatchConditions,as well as providing an identifier for logging purposes. A good name should be descriptive ofthe associated expression.Name must be a qualified name consisting of alphanumeric characters, '-', '_' or '.', andmust start and end with an alphanumeric character (e.g. 'MyName',  or 'my.name',  or'123-abc', regex used for validation is '([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9]') with anoptional DNS subdomain prefix and '/' (e.g. 'example.com/MyName')Required.",
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

								"context": schema.ListNestedAttribute{
									Description:         "Context defines variables and data sources that can be used during rule execution.",
									MarkdownDescription: "Context defines variables and data sources that can be used during rule execution.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"api_call": schema.SingleNestedAttribute{
												Description:         "APICall is an HTTP request to the Kubernetes API server, or other JSON web service.The data returned is stored in the context with the name for the context entry.",
												MarkdownDescription: "APICall is an HTTP request to the Kubernetes API server, or other JSON web service.The data returned is stored in the context with the name for the context entry.",
												Attributes: map[string]schema.Attribute{
													"data": schema.ListNestedAttribute{
														Description:         "The data object specifies the POST data sent to the server.Only applicable when the method field is set to POST.",
														MarkdownDescription: "The data object specifies the POST data sent to the server.Only applicable when the method field is set to POST.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "Key is a unique identifier for the data value",
																	MarkdownDescription: "Key is a unique identifier for the data value",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"value": schema.MapAttribute{
																	Description:         "Value is the data value",
																	MarkdownDescription: "Value is the data value",
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

													"jmes_path": schema.StringAttribute{
														Description:         "JMESPath is an optional JSON Match Expression that can be used totransform the JSON response returned from the server. For examplea JMESPath of 'items | length(@)' applied to the API server responsefor the URLPath '/apis/apps/v1/deployments' will return the total countof deployments across all namespaces.",
														MarkdownDescription: "JMESPath is an optional JSON Match Expression that can be used totransform the JSON response returned from the server. For examplea JMESPath of 'items | length(@)' applied to the API server responsefor the URLPath '/apis/apps/v1/deployments' will return the total countof deployments across all namespaces.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"method": schema.StringAttribute{
														Description:         "Method is the HTTP request type (GET or POST). Defaults to GET.",
														MarkdownDescription: "Method is the HTTP request type (GET or POST). Defaults to GET.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("GET", "POST"),
														},
													},

													"service": schema.SingleNestedAttribute{
														Description:         "Service is an API call to a JSON web service.This is used for non-Kubernetes API server calls.It's mutually exclusive with the URLPath field.",
														MarkdownDescription: "Service is an API call to a JSON web service.This is used for non-Kubernetes API server calls.It's mutually exclusive with the URLPath field.",
														Attributes: map[string]schema.Attribute{
															"ca_bundle": schema.StringAttribute{
																Description:         "CABundle is a PEM encoded CA bundle which will be used to validatethe server certificate.",
																MarkdownDescription: "CABundle is a PEM encoded CA bundle which will be used to validatethe server certificate.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"url": schema.StringAttribute{
																Description:         "URL is the JSON web service URL. A typical form is'https://{service}.{namespace}:{port}/{path}'.",
																MarkdownDescription: "URL is the JSON web service URL. A typical form is'https://{service}.{namespace}:{port}/{path}'.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"url_path": schema.StringAttribute{
														Description:         "URLPath is the URL path to be used in the HTTP GET or POST request to theKubernetes API server (e.g. '/api/v1/namespaces' or  '/apis/apps/v1/deployments').The format required is the same format used by the 'kubectl get --raw' command.See https://kyverno.io/docs/writing-policies/external-data-sources/#variables-from-kubernetes-api-server-callsfor details.It's mutually exclusive with the Service field.",
														MarkdownDescription: "URLPath is the URL path to be used in the HTTP GET or POST request to theKubernetes API server (e.g. '/api/v1/namespaces' or  '/apis/apps/v1/deployments').The format required is the same format used by the 'kubectl get --raw' command.See https://kyverno.io/docs/writing-policies/external-data-sources/#variables-from-kubernetes-api-server-callsfor details.It's mutually exclusive with the Service field.",
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
												Description:         "ConfigMap is the ConfigMap reference.",
												MarkdownDescription: "ConfigMap is the ConfigMap reference.",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "Name is the ConfigMap name.",
														MarkdownDescription: "Name is the ConfigMap name.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace is the ConfigMap namespace.",
														MarkdownDescription: "Namespace is the ConfigMap namespace.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"global_reference": schema.SingleNestedAttribute{
												Description:         "GlobalContextEntryReference is a reference to a cached global context entry.",
												MarkdownDescription: "GlobalContextEntryReference is a reference to a cached global context entry.",
												Attributes: map[string]schema.Attribute{
													"jmes_path": schema.StringAttribute{
														Description:         "JMESPath is an optional JSON Match Expression that can be used totransform the JSON response returned from the server. For examplea JMESPath of 'items | length(@)' applied to the API server responsefor the URLPath '/apis/apps/v1/deployments' will return the total countof deployments across all namespaces.",
														MarkdownDescription: "JMESPath is an optional JSON Match Expression that can be used totransform the JSON response returned from the server. For examplea JMESPath of 'items | length(@)' applied to the API server responsefor the URLPath '/apis/apps/v1/deployments' will return the total countof deployments across all namespaces.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the global context entry",
														MarkdownDescription: "Name of the global context entry",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"image_registry": schema.SingleNestedAttribute{
												Description:         "ImageRegistry defines requests to an OCI/Docker V2 registry to fetch imagedetails.",
												MarkdownDescription: "ImageRegistry defines requests to an OCI/Docker V2 registry to fetch imagedetails.",
												Attributes: map[string]schema.Attribute{
													"image_registry_credentials": schema.SingleNestedAttribute{
														Description:         "ImageRegistryCredentials provides credentials that will be used for authentication with registry",
														MarkdownDescription: "ImageRegistryCredentials provides credentials that will be used for authentication with registry",
														Attributes: map[string]schema.Attribute{
															"allow_insecure_registry": schema.BoolAttribute{
																Description:         "AllowInsecureRegistry allows insecure access to a registry.",
																MarkdownDescription: "AllowInsecureRegistry allows insecure access to a registry.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"providers": schema.ListAttribute{
																Description:         "Providers specifies a list of OCI Registry names, whose authentication providers are provided.It can be of one of these values: default,google,azure,amazon,github.",
																MarkdownDescription: "Providers specifies a list of OCI Registry names, whose authentication providers are provided.It can be of one of these values: default,google,azure,amazon,github.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secrets": schema.ListAttribute{
																Description:         "Secrets specifies a list of secrets that are provided for credentials.Secrets must live in the Kyverno namespace.",
																MarkdownDescription: "Secrets specifies a list of secrets that are provided for credentials.Secrets must live in the Kyverno namespace.",
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

													"jmes_path": schema.StringAttribute{
														Description:         "JMESPath is an optional JSON Match Expression that can be used totransform the ImageData struct returned as a result of processingthe image reference.",
														MarkdownDescription: "JMESPath is an optional JSON Match Expression that can be used totransform the ImageData struct returned as a result of processingthe image reference.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"reference": schema.StringAttribute{
														Description:         "Reference is image reference to a container image in the registry.Example: ghcr.io/kyverno/kyverno:latest",
														MarkdownDescription: "Reference is image reference to a container image in the registry.Example: ghcr.io/kyverno/kyverno:latest",
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
												Description:         "Name is the variable name.",
												MarkdownDescription: "Name is the variable name.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"variable": schema.SingleNestedAttribute{
												Description:         "Variable defines an arbitrary JMESPath context variable that can be defined inline.",
												MarkdownDescription: "Variable defines an arbitrary JMESPath context variable that can be defined inline.",
												Attributes: map[string]schema.Attribute{
													"default": schema.MapAttribute{
														Description:         "Default is an optional arbitrary JSON object that the variable may take if the JMESPathexpression evaluates to nil",
														MarkdownDescription: "Default is an optional arbitrary JSON object that the variable may take if the JMESPathexpression evaluates to nil",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"jmes_path": schema.StringAttribute{
														Description:         "JMESPath is an optional JMESPath Expression that can be used totransform the variable.",
														MarkdownDescription: "JMESPath is an optional JMESPath Expression that can be used totransform the variable.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"value": schema.MapAttribute{
														Description:         "Value is any arbitrary JSON object representable in YAML or JSON form.",
														MarkdownDescription: "Value is any arbitrary JSON object representable in YAML or JSON form.",
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
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"exclude": schema.SingleNestedAttribute{
									Description:         "ExcludeResources defines when this policy rule should not be applied. The excludecriteria can include resource information (e.g. kind, name, namespace, labels)and admission review request information like the name or role.",
									MarkdownDescription: "ExcludeResources defines when this policy rule should not be applied. The excludecriteria can include resource information (e.g. kind, name, namespace, labels)and admission review request information like the name or role.",
									Attributes: map[string]schema.Attribute{
										"all": schema.ListNestedAttribute{
											Description:         "All allows specifying resources which will be ANDed",
											MarkdownDescription: "All allows specifying resources which will be ANDed",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"cluster_roles": schema.ListAttribute{
														Description:         "ClusterRoles is the list of cluster-wide role names for the user.",
														MarkdownDescription: "ClusterRoles is the list of cluster-wide role names for the user.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"resources": schema.SingleNestedAttribute{
														Description:         "ResourceDescription contains information about the resource being created or modified.",
														MarkdownDescription: "ResourceDescription contains information about the resource being created or modified.",
														Attributes: map[string]schema.Attribute{
															"annotations": schema.MapAttribute{
																Description:         "Annotations is a  map of annotations (key-value pairs of type string). Annotation keysand values support the wildcard characters '*' (matches zero or many characters) and'?' (matches at least one character).",
																MarkdownDescription: "Annotations is a  map of annotations (key-value pairs of type string). Annotation keysand values support the wildcard characters '*' (matches zero or many characters) and'?' (matches at least one character).",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kinds": schema.ListAttribute{
																Description:         "Kinds is a list of resource kinds.",
																MarkdownDescription: "Kinds is a list of resource kinds.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "Name is the name of the resource. The name supports wildcard characters'*' (matches zero or many characters) and '?' (at least one character).NOTE: 'Name' is being deprecated in favor of 'Names'.",
																MarkdownDescription: "Name is the name of the resource. The name supports wildcard characters'*' (matches zero or many characters) and '?' (at least one character).NOTE: 'Name' is being deprecated in favor of 'Names'.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"names": schema.ListAttribute{
																Description:         "Names are the names of the resources. Each name supports wildcard characters'*' (matches zero or many characters) and '?' (at least one character).",
																MarkdownDescription: "Names are the names of the resources. Each name supports wildcard characters'*' (matches zero or many characters) and '?' (at least one character).",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace_selector": schema.SingleNestedAttribute{
																Description:         "NamespaceSelector is a label selector for the resource namespace. Label keys and valuesin 'matchLabels' support the wildcard characters '*' (matches zero or many characters)and '?' (matches one character).Wildcards allows writing label selectors like['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value butdoes not match an empty label set.",
																MarkdownDescription: "NamespaceSelector is a label selector for the resource namespace. Label keys and valuesin 'matchLabels' support the wildcard characters '*' (matches zero or many characters)and '?' (matches one character).Wildcards allows writing label selectors like['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value butdoes not match an empty label set.",
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
																Required: false,
																Optional: true,
																Computed: false,
															},

															"namespaces": schema.ListAttribute{
																Description:         "Namespaces is a list of namespaces names. Each name supports wildcard characters'*' (matches zero or many characters) and '?' (at least one character).",
																MarkdownDescription: "Namespaces is a list of namespaces names. Each name supports wildcard characters'*' (matches zero or many characters) and '?' (at least one character).",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"operations": schema.ListAttribute{
																Description:         "Operations can contain values ['CREATE, 'UPDATE', 'CONNECT', 'DELETE'], which are used to match a specific action.",
																MarkdownDescription: "Operations can contain values ['CREATE, 'UPDATE', 'CONNECT', 'DELETE'], which are used to match a specific action.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"selector": schema.SingleNestedAttribute{
																Description:         "Selector is a label selector. Label keys and values in 'matchLabels' support the wildcardcharacters '*' (matches zero or many characters) and '?' (matches one character).Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note thatusing ['*' : '*'] matches any key and value but does not match an empty label set.",
																MarkdownDescription: "Selector is a label selector. Label keys and values in 'matchLabels' support the wildcardcharacters '*' (matches zero or many characters) and '?' (matches one character).Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note thatusing ['*' : '*'] matches any key and value but does not match an empty label set.",
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
																Required: false,
																Optional: true,
																Computed: false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"roles": schema.ListAttribute{
														Description:         "Roles is the list of namespaced role names for the user.",
														MarkdownDescription: "Roles is the list of namespaced role names for the user.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"subjects": schema.ListNestedAttribute{
														Description:         "Subjects is the list of subject names like users, user groups, and service accounts.",
														MarkdownDescription: "Subjects is the list of subject names like users, user groups, and service accounts.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"api_group": schema.StringAttribute{
																	Description:         "APIGroup holds the API group of the referenced subject.Defaults to '' for ServiceAccount subjects.Defaults to 'rbac.authorization.k8s.io' for User and Group subjects.",
																	MarkdownDescription: "APIGroup holds the API group of the referenced subject.Defaults to '' for ServiceAccount subjects.Defaults to 'rbac.authorization.k8s.io' for User and Group subjects.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"kind": schema.StringAttribute{
																	Description:         "Kind of object being referenced. Values defined by this API group are 'User', 'Group', and 'ServiceAccount'.If the Authorizer does not recognized the kind value, the Authorizer should report an error.",
																	MarkdownDescription: "Kind of object being referenced. Values defined by this API group are 'User', 'Group', and 'ServiceAccount'.If the Authorizer does not recognized the kind value, the Authorizer should report an error.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name of the object being referenced.",
																	MarkdownDescription: "Name of the object being referenced.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"namespace": schema.StringAttribute{
																	Description:         "Namespace of the referenced object.  If the object kind is non-namespace, such as 'User' or 'Group', and this value is not emptythe Authorizer should report an error.",
																	MarkdownDescription: "Namespace of the referenced object.  If the object kind is non-namespace, such as 'User' or 'Group', and this value is not emptythe Authorizer should report an error.",
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"any": schema.ListNestedAttribute{
											Description:         "Any allows specifying resources which will be ORed",
											MarkdownDescription: "Any allows specifying resources which will be ORed",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"cluster_roles": schema.ListAttribute{
														Description:         "ClusterRoles is the list of cluster-wide role names for the user.",
														MarkdownDescription: "ClusterRoles is the list of cluster-wide role names for the user.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"resources": schema.SingleNestedAttribute{
														Description:         "ResourceDescription contains information about the resource being created or modified.",
														MarkdownDescription: "ResourceDescription contains information about the resource being created or modified.",
														Attributes: map[string]schema.Attribute{
															"annotations": schema.MapAttribute{
																Description:         "Annotations is a  map of annotations (key-value pairs of type string). Annotation keysand values support the wildcard characters '*' (matches zero or many characters) and'?' (matches at least one character).",
																MarkdownDescription: "Annotations is a  map of annotations (key-value pairs of type string). Annotation keysand values support the wildcard characters '*' (matches zero or many characters) and'?' (matches at least one character).",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kinds": schema.ListAttribute{
																Description:         "Kinds is a list of resource kinds.",
																MarkdownDescription: "Kinds is a list of resource kinds.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "Name is the name of the resource. The name supports wildcard characters'*' (matches zero or many characters) and '?' (at least one character).NOTE: 'Name' is being deprecated in favor of 'Names'.",
																MarkdownDescription: "Name is the name of the resource. The name supports wildcard characters'*' (matches zero or many characters) and '?' (at least one character).NOTE: 'Name' is being deprecated in favor of 'Names'.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"names": schema.ListAttribute{
																Description:         "Names are the names of the resources. Each name supports wildcard characters'*' (matches zero or many characters) and '?' (at least one character).",
																MarkdownDescription: "Names are the names of the resources. Each name supports wildcard characters'*' (matches zero or many characters) and '?' (at least one character).",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace_selector": schema.SingleNestedAttribute{
																Description:         "NamespaceSelector is a label selector for the resource namespace. Label keys and valuesin 'matchLabels' support the wildcard characters '*' (matches zero or many characters)and '?' (matches one character).Wildcards allows writing label selectors like['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value butdoes not match an empty label set.",
																MarkdownDescription: "NamespaceSelector is a label selector for the resource namespace. Label keys and valuesin 'matchLabels' support the wildcard characters '*' (matches zero or many characters)and '?' (matches one character).Wildcards allows writing label selectors like['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value butdoes not match an empty label set.",
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
																Required: false,
																Optional: true,
																Computed: false,
															},

															"namespaces": schema.ListAttribute{
																Description:         "Namespaces is a list of namespaces names. Each name supports wildcard characters'*' (matches zero or many characters) and '?' (at least one character).",
																MarkdownDescription: "Namespaces is a list of namespaces names. Each name supports wildcard characters'*' (matches zero or many characters) and '?' (at least one character).",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"operations": schema.ListAttribute{
																Description:         "Operations can contain values ['CREATE, 'UPDATE', 'CONNECT', 'DELETE'], which are used to match a specific action.",
																MarkdownDescription: "Operations can contain values ['CREATE, 'UPDATE', 'CONNECT', 'DELETE'], which are used to match a specific action.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"selector": schema.SingleNestedAttribute{
																Description:         "Selector is a label selector. Label keys and values in 'matchLabels' support the wildcardcharacters '*' (matches zero or many characters) and '?' (matches one character).Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note thatusing ['*' : '*'] matches any key and value but does not match an empty label set.",
																MarkdownDescription: "Selector is a label selector. Label keys and values in 'matchLabels' support the wildcardcharacters '*' (matches zero or many characters) and '?' (matches one character).Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note thatusing ['*' : '*'] matches any key and value but does not match an empty label set.",
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
																Required: false,
																Optional: true,
																Computed: false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"roles": schema.ListAttribute{
														Description:         "Roles is the list of namespaced role names for the user.",
														MarkdownDescription: "Roles is the list of namespaced role names for the user.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"subjects": schema.ListNestedAttribute{
														Description:         "Subjects is the list of subject names like users, user groups, and service accounts.",
														MarkdownDescription: "Subjects is the list of subject names like users, user groups, and service accounts.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"api_group": schema.StringAttribute{
																	Description:         "APIGroup holds the API group of the referenced subject.Defaults to '' for ServiceAccount subjects.Defaults to 'rbac.authorization.k8s.io' for User and Group subjects.",
																	MarkdownDescription: "APIGroup holds the API group of the referenced subject.Defaults to '' for ServiceAccount subjects.Defaults to 'rbac.authorization.k8s.io' for User and Group subjects.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"kind": schema.StringAttribute{
																	Description:         "Kind of object being referenced. Values defined by this API group are 'User', 'Group', and 'ServiceAccount'.If the Authorizer does not recognized the kind value, the Authorizer should report an error.",
																	MarkdownDescription: "Kind of object being referenced. Values defined by this API group are 'User', 'Group', and 'ServiceAccount'.If the Authorizer does not recognized the kind value, the Authorizer should report an error.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name of the object being referenced.",
																	MarkdownDescription: "Name of the object being referenced.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"namespace": schema.StringAttribute{
																	Description:         "Namespace of the referenced object.  If the object kind is non-namespace, such as 'User' or 'Group', and this value is not emptythe Authorizer should report an error.",
																	MarkdownDescription: "Namespace of the referenced object.  If the object kind is non-namespace, such as 'User' or 'Group', and this value is not emptythe Authorizer should report an error.",
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"cluster_roles": schema.ListAttribute{
											Description:         "ClusterRoles is the list of cluster-wide role names for the user.",
											MarkdownDescription: "ClusterRoles is the list of cluster-wide role names for the user.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"resources": schema.SingleNestedAttribute{
											Description:         "ResourceDescription contains information about the resource being created or modified.Requires at least one tag to be specified when under MatchResources.Specifying ResourceDescription directly under match is being deprecated.Please specify under 'any' or 'all' instead.",
											MarkdownDescription: "ResourceDescription contains information about the resource being created or modified.Requires at least one tag to be specified when under MatchResources.Specifying ResourceDescription directly under match is being deprecated.Please specify under 'any' or 'all' instead.",
											Attributes: map[string]schema.Attribute{
												"annotations": schema.MapAttribute{
													Description:         "Annotations is a  map of annotations (key-value pairs of type string). Annotation keysand values support the wildcard characters '*' (matches zero or many characters) and'?' (matches at least one character).",
													MarkdownDescription: "Annotations is a  map of annotations (key-value pairs of type string). Annotation keysand values support the wildcard characters '*' (matches zero or many characters) and'?' (matches at least one character).",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"kinds": schema.ListAttribute{
													Description:         "Kinds is a list of resource kinds.",
													MarkdownDescription: "Kinds is a list of resource kinds.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name is the name of the resource. The name supports wildcard characters'*' (matches zero or many characters) and '?' (at least one character).NOTE: 'Name' is being deprecated in favor of 'Names'.",
													MarkdownDescription: "Name is the name of the resource. The name supports wildcard characters'*' (matches zero or many characters) and '?' (at least one character).NOTE: 'Name' is being deprecated in favor of 'Names'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"names": schema.ListAttribute{
													Description:         "Names are the names of the resources. Each name supports wildcard characters'*' (matches zero or many characters) and '?' (at least one character).",
													MarkdownDescription: "Names are the names of the resources. Each name supports wildcard characters'*' (matches zero or many characters) and '?' (at least one character).",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespace_selector": schema.SingleNestedAttribute{
													Description:         "NamespaceSelector is a label selector for the resource namespace. Label keys and valuesin 'matchLabels' support the wildcard characters '*' (matches zero or many characters)and '?' (matches one character).Wildcards allows writing label selectors like['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value butdoes not match an empty label set.",
													MarkdownDescription: "NamespaceSelector is a label selector for the resource namespace. Label keys and valuesin 'matchLabels' support the wildcard characters '*' (matches zero or many characters)and '?' (matches one character).Wildcards allows writing label selectors like['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value butdoes not match an empty label set.",
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
													Required: false,
													Optional: true,
													Computed: false,
												},

												"namespaces": schema.ListAttribute{
													Description:         "Namespaces is a list of namespaces names. Each name supports wildcard characters'*' (matches zero or many characters) and '?' (at least one character).",
													MarkdownDescription: "Namespaces is a list of namespaces names. Each name supports wildcard characters'*' (matches zero or many characters) and '?' (at least one character).",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"operations": schema.ListAttribute{
													Description:         "Operations can contain values ['CREATE, 'UPDATE', 'CONNECT', 'DELETE'], which are used to match a specific action.",
													MarkdownDescription: "Operations can contain values ['CREATE, 'UPDATE', 'CONNECT', 'DELETE'], which are used to match a specific action.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"selector": schema.SingleNestedAttribute{
													Description:         "Selector is a label selector. Label keys and values in 'matchLabels' support the wildcardcharacters '*' (matches zero or many characters) and '?' (matches one character).Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note thatusing ['*' : '*'] matches any key and value but does not match an empty label set.",
													MarkdownDescription: "Selector is a label selector. Label keys and values in 'matchLabels' support the wildcardcharacters '*' (matches zero or many characters) and '?' (matches one character).Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note thatusing ['*' : '*'] matches any key and value but does not match an empty label set.",
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
													Required: false,
													Optional: true,
													Computed: false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"roles": schema.ListAttribute{
											Description:         "Roles is the list of namespaced role names for the user.",
											MarkdownDescription: "Roles is the list of namespaced role names for the user.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"subjects": schema.ListNestedAttribute{
											Description:         "Subjects is the list of subject names like users, user groups, and service accounts.",
											MarkdownDescription: "Subjects is the list of subject names like users, user groups, and service accounts.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"api_group": schema.StringAttribute{
														Description:         "APIGroup holds the API group of the referenced subject.Defaults to '' for ServiceAccount subjects.Defaults to 'rbac.authorization.k8s.io' for User and Group subjects.",
														MarkdownDescription: "APIGroup holds the API group of the referenced subject.Defaults to '' for ServiceAccount subjects.Defaults to 'rbac.authorization.k8s.io' for User and Group subjects.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"kind": schema.StringAttribute{
														Description:         "Kind of object being referenced. Values defined by this API group are 'User', 'Group', and 'ServiceAccount'.If the Authorizer does not recognized the kind value, the Authorizer should report an error.",
														MarkdownDescription: "Kind of object being referenced. Values defined by this API group are 'User', 'Group', and 'ServiceAccount'.If the Authorizer does not recognized the kind value, the Authorizer should report an error.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the object being referenced.",
														MarkdownDescription: "Name of the object being referenced.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the referenced object.  If the object kind is non-namespace, such as 'User' or 'Group', and this value is not emptythe Authorizer should report an error.",
														MarkdownDescription: "Namespace of the referenced object.  If the object kind is non-namespace, such as 'User' or 'Group', and this value is not emptythe Authorizer should report an error.",
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

								"generate": schema.SingleNestedAttribute{
									Description:         "Generation is used to create new resources.",
									MarkdownDescription: "Generation is used to create new resources.",
									Attributes: map[string]schema.Attribute{
										"api_version": schema.StringAttribute{
											Description:         "APIVersion specifies resource apiVersion.",
											MarkdownDescription: "APIVersion specifies resource apiVersion.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"clone": schema.SingleNestedAttribute{
											Description:         "Clone specifies the source resource used to populate each generated resource.At most one of Data or Clone can be specified. If neither are provided, the generatedresource will be created with default data only.",
											MarkdownDescription: "Clone specifies the source resource used to populate each generated resource.At most one of Data or Clone can be specified. If neither are provided, the generatedresource will be created with default data only.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name specifies name of the resource.",
													MarkdownDescription: "Name specifies name of the resource.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespace": schema.StringAttribute{
													Description:         "Namespace specifies source resource namespace.",
													MarkdownDescription: "Namespace specifies source resource namespace.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"clone_list": schema.SingleNestedAttribute{
											Description:         "CloneList specifies the list of source resource used to populate each generated resource.",
											MarkdownDescription: "CloneList specifies the list of source resource used to populate each generated resource.",
											Attributes: map[string]schema.Attribute{
												"kinds": schema.ListAttribute{
													Description:         "Kinds is a list of resource kinds.",
													MarkdownDescription: "Kinds is a list of resource kinds.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespace": schema.StringAttribute{
													Description:         "Namespace specifies source resource namespace.",
													MarkdownDescription: "Namespace specifies source resource namespace.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"selector": schema.SingleNestedAttribute{
													Description:         "Selector is a label selector. Label keys and values in 'matchLabels'.wildcard characters are not supported.",
													MarkdownDescription: "Selector is a label selector. Label keys and values in 'matchLabels'.wildcard characters are not supported.",
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
													Required: false,
													Optional: true,
													Computed: false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"data": schema.MapAttribute{
											Description:         "Data provides the resource declaration used to populate each generated resource.At most one of Data or Clone must be specified. If neither are provided, the generatedresource will be created with default data only.",
											MarkdownDescription: "Data provides the resource declaration used to populate each generated resource.At most one of Data or Clone must be specified. If neither are provided, the generatedresource will be created with default data only.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"generate_existing": schema.BoolAttribute{
											Description:         "GenerateExisting controls whether to trigger the rule in existing resourcesIf is set to 'true' the rule will be triggered and applied to existing matched resources.",
											MarkdownDescription: "GenerateExisting controls whether to trigger the rule in existing resourcesIf is set to 'true' the rule will be triggered and applied to existing matched resources.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"kind": schema.StringAttribute{
											Description:         "Kind specifies resource kind.",
											MarkdownDescription: "Kind specifies resource kind.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name specifies the resource name.",
											MarkdownDescription: "Name specifies the resource name.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"namespace": schema.StringAttribute{
											Description:         "Namespace specifies resource namespace.",
											MarkdownDescription: "Namespace specifies resource namespace.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"orphan_downstream_on_policy_delete": schema.BoolAttribute{
											Description:         "OrphanDownstreamOnPolicyDelete controls whether generated resources should be deleted when the rule that generatedthem is deleted with synchronization enabled. This option is only applicable to generate rules of the data type.See https://kyverno.io/docs/writing-policies/generate/#data-examples.Defaults to 'false' if not specified.",
											MarkdownDescription: "OrphanDownstreamOnPolicyDelete controls whether generated resources should be deleted when the rule that generatedthem is deleted with synchronization enabled. This option is only applicable to generate rules of the data type.See https://kyverno.io/docs/writing-policies/generate/#data-examples.Defaults to 'false' if not specified.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"synchronize": schema.BoolAttribute{
											Description:         "Synchronize controls if generated resources should be kept in-sync with their source resource.If Synchronize is set to 'true' changes to generated resources will be overwritten with resourcedata from Data or the resource specified in the Clone declaration.Optional. Defaults to 'false' if not specified.",
											MarkdownDescription: "Synchronize controls if generated resources should be kept in-sync with their source resource.If Synchronize is set to 'true' changes to generated resources will be overwritten with resourcedata from Data or the resource specified in the Clone declaration.Optional. Defaults to 'false' if not specified.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"uid": schema.StringAttribute{
											Description:         "UID specifies the resource uid.",
											MarkdownDescription: "UID specifies the resource uid.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"image_extractors": schema.MapAttribute{
									Description:         "ImageExtractors defines a mapping from kinds to ImageExtractorConfigs.This config is only valid for verifyImages rules.",
									MarkdownDescription: "ImageExtractors defines a mapping from kinds to ImageExtractorConfigs.This config is only valid for verifyImages rules.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"match": schema.SingleNestedAttribute{
									Description:         "MatchResources defines when this policy rule should be applied. The matchcriteria can include resource information (e.g. kind, name, namespace, labels)and admission review request information like the user name or role.At least one kind is required.",
									MarkdownDescription: "MatchResources defines when this policy rule should be applied. The matchcriteria can include resource information (e.g. kind, name, namespace, labels)and admission review request information like the user name or role.At least one kind is required.",
									Attributes: map[string]schema.Attribute{
										"all": schema.ListNestedAttribute{
											Description:         "All allows specifying resources which will be ANDed",
											MarkdownDescription: "All allows specifying resources which will be ANDed",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"cluster_roles": schema.ListAttribute{
														Description:         "ClusterRoles is the list of cluster-wide role names for the user.",
														MarkdownDescription: "ClusterRoles is the list of cluster-wide role names for the user.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"resources": schema.SingleNestedAttribute{
														Description:         "ResourceDescription contains information about the resource being created or modified.",
														MarkdownDescription: "ResourceDescription contains information about the resource being created or modified.",
														Attributes: map[string]schema.Attribute{
															"annotations": schema.MapAttribute{
																Description:         "Annotations is a  map of annotations (key-value pairs of type string). Annotation keysand values support the wildcard characters '*' (matches zero or many characters) and'?' (matches at least one character).",
																MarkdownDescription: "Annotations is a  map of annotations (key-value pairs of type string). Annotation keysand values support the wildcard characters '*' (matches zero or many characters) and'?' (matches at least one character).",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kinds": schema.ListAttribute{
																Description:         "Kinds is a list of resource kinds.",
																MarkdownDescription: "Kinds is a list of resource kinds.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "Name is the name of the resource. The name supports wildcard characters'*' (matches zero or many characters) and '?' (at least one character).NOTE: 'Name' is being deprecated in favor of 'Names'.",
																MarkdownDescription: "Name is the name of the resource. The name supports wildcard characters'*' (matches zero or many characters) and '?' (at least one character).NOTE: 'Name' is being deprecated in favor of 'Names'.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"names": schema.ListAttribute{
																Description:         "Names are the names of the resources. Each name supports wildcard characters'*' (matches zero or many characters) and '?' (at least one character).",
																MarkdownDescription: "Names are the names of the resources. Each name supports wildcard characters'*' (matches zero or many characters) and '?' (at least one character).",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace_selector": schema.SingleNestedAttribute{
																Description:         "NamespaceSelector is a label selector for the resource namespace. Label keys and valuesin 'matchLabels' support the wildcard characters '*' (matches zero or many characters)and '?' (matches one character).Wildcards allows writing label selectors like['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value butdoes not match an empty label set.",
																MarkdownDescription: "NamespaceSelector is a label selector for the resource namespace. Label keys and valuesin 'matchLabels' support the wildcard characters '*' (matches zero or many characters)and '?' (matches one character).Wildcards allows writing label selectors like['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value butdoes not match an empty label set.",
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
																Required: false,
																Optional: true,
																Computed: false,
															},

															"namespaces": schema.ListAttribute{
																Description:         "Namespaces is a list of namespaces names. Each name supports wildcard characters'*' (matches zero or many characters) and '?' (at least one character).",
																MarkdownDescription: "Namespaces is a list of namespaces names. Each name supports wildcard characters'*' (matches zero or many characters) and '?' (at least one character).",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"operations": schema.ListAttribute{
																Description:         "Operations can contain values ['CREATE, 'UPDATE', 'CONNECT', 'DELETE'], which are used to match a specific action.",
																MarkdownDescription: "Operations can contain values ['CREATE, 'UPDATE', 'CONNECT', 'DELETE'], which are used to match a specific action.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"selector": schema.SingleNestedAttribute{
																Description:         "Selector is a label selector. Label keys and values in 'matchLabels' support the wildcardcharacters '*' (matches zero or many characters) and '?' (matches one character).Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note thatusing ['*' : '*'] matches any key and value but does not match an empty label set.",
																MarkdownDescription: "Selector is a label selector. Label keys and values in 'matchLabels' support the wildcardcharacters '*' (matches zero or many characters) and '?' (matches one character).Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note thatusing ['*' : '*'] matches any key and value but does not match an empty label set.",
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
																Required: false,
																Optional: true,
																Computed: false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"roles": schema.ListAttribute{
														Description:         "Roles is the list of namespaced role names for the user.",
														MarkdownDescription: "Roles is the list of namespaced role names for the user.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"subjects": schema.ListNestedAttribute{
														Description:         "Subjects is the list of subject names like users, user groups, and service accounts.",
														MarkdownDescription: "Subjects is the list of subject names like users, user groups, and service accounts.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"api_group": schema.StringAttribute{
																	Description:         "APIGroup holds the API group of the referenced subject.Defaults to '' for ServiceAccount subjects.Defaults to 'rbac.authorization.k8s.io' for User and Group subjects.",
																	MarkdownDescription: "APIGroup holds the API group of the referenced subject.Defaults to '' for ServiceAccount subjects.Defaults to 'rbac.authorization.k8s.io' for User and Group subjects.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"kind": schema.StringAttribute{
																	Description:         "Kind of object being referenced. Values defined by this API group are 'User', 'Group', and 'ServiceAccount'.If the Authorizer does not recognized the kind value, the Authorizer should report an error.",
																	MarkdownDescription: "Kind of object being referenced. Values defined by this API group are 'User', 'Group', and 'ServiceAccount'.If the Authorizer does not recognized the kind value, the Authorizer should report an error.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name of the object being referenced.",
																	MarkdownDescription: "Name of the object being referenced.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"namespace": schema.StringAttribute{
																	Description:         "Namespace of the referenced object.  If the object kind is non-namespace, such as 'User' or 'Group', and this value is not emptythe Authorizer should report an error.",
																	MarkdownDescription: "Namespace of the referenced object.  If the object kind is non-namespace, such as 'User' or 'Group', and this value is not emptythe Authorizer should report an error.",
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"any": schema.ListNestedAttribute{
											Description:         "Any allows specifying resources which will be ORed",
											MarkdownDescription: "Any allows specifying resources which will be ORed",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"cluster_roles": schema.ListAttribute{
														Description:         "ClusterRoles is the list of cluster-wide role names for the user.",
														MarkdownDescription: "ClusterRoles is the list of cluster-wide role names for the user.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"resources": schema.SingleNestedAttribute{
														Description:         "ResourceDescription contains information about the resource being created or modified.",
														MarkdownDescription: "ResourceDescription contains information about the resource being created or modified.",
														Attributes: map[string]schema.Attribute{
															"annotations": schema.MapAttribute{
																Description:         "Annotations is a  map of annotations (key-value pairs of type string). Annotation keysand values support the wildcard characters '*' (matches zero or many characters) and'?' (matches at least one character).",
																MarkdownDescription: "Annotations is a  map of annotations (key-value pairs of type string). Annotation keysand values support the wildcard characters '*' (matches zero or many characters) and'?' (matches at least one character).",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kinds": schema.ListAttribute{
																Description:         "Kinds is a list of resource kinds.",
																MarkdownDescription: "Kinds is a list of resource kinds.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "Name is the name of the resource. The name supports wildcard characters'*' (matches zero or many characters) and '?' (at least one character).NOTE: 'Name' is being deprecated in favor of 'Names'.",
																MarkdownDescription: "Name is the name of the resource. The name supports wildcard characters'*' (matches zero or many characters) and '?' (at least one character).NOTE: 'Name' is being deprecated in favor of 'Names'.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"names": schema.ListAttribute{
																Description:         "Names are the names of the resources. Each name supports wildcard characters'*' (matches zero or many characters) and '?' (at least one character).",
																MarkdownDescription: "Names are the names of the resources. Each name supports wildcard characters'*' (matches zero or many characters) and '?' (at least one character).",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace_selector": schema.SingleNestedAttribute{
																Description:         "NamespaceSelector is a label selector for the resource namespace. Label keys and valuesin 'matchLabels' support the wildcard characters '*' (matches zero or many characters)and '?' (matches one character).Wildcards allows writing label selectors like['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value butdoes not match an empty label set.",
																MarkdownDescription: "NamespaceSelector is a label selector for the resource namespace. Label keys and valuesin 'matchLabels' support the wildcard characters '*' (matches zero or many characters)and '?' (matches one character).Wildcards allows writing label selectors like['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value butdoes not match an empty label set.",
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
																Required: false,
																Optional: true,
																Computed: false,
															},

															"namespaces": schema.ListAttribute{
																Description:         "Namespaces is a list of namespaces names. Each name supports wildcard characters'*' (matches zero or many characters) and '?' (at least one character).",
																MarkdownDescription: "Namespaces is a list of namespaces names. Each name supports wildcard characters'*' (matches zero or many characters) and '?' (at least one character).",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"operations": schema.ListAttribute{
																Description:         "Operations can contain values ['CREATE, 'UPDATE', 'CONNECT', 'DELETE'], which are used to match a specific action.",
																MarkdownDescription: "Operations can contain values ['CREATE, 'UPDATE', 'CONNECT', 'DELETE'], which are used to match a specific action.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"selector": schema.SingleNestedAttribute{
																Description:         "Selector is a label selector. Label keys and values in 'matchLabels' support the wildcardcharacters '*' (matches zero or many characters) and '?' (matches one character).Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note thatusing ['*' : '*'] matches any key and value but does not match an empty label set.",
																MarkdownDescription: "Selector is a label selector. Label keys and values in 'matchLabels' support the wildcardcharacters '*' (matches zero or many characters) and '?' (matches one character).Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note thatusing ['*' : '*'] matches any key and value but does not match an empty label set.",
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
																Required: false,
																Optional: true,
																Computed: false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"roles": schema.ListAttribute{
														Description:         "Roles is the list of namespaced role names for the user.",
														MarkdownDescription: "Roles is the list of namespaced role names for the user.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"subjects": schema.ListNestedAttribute{
														Description:         "Subjects is the list of subject names like users, user groups, and service accounts.",
														MarkdownDescription: "Subjects is the list of subject names like users, user groups, and service accounts.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"api_group": schema.StringAttribute{
																	Description:         "APIGroup holds the API group of the referenced subject.Defaults to '' for ServiceAccount subjects.Defaults to 'rbac.authorization.k8s.io' for User and Group subjects.",
																	MarkdownDescription: "APIGroup holds the API group of the referenced subject.Defaults to '' for ServiceAccount subjects.Defaults to 'rbac.authorization.k8s.io' for User and Group subjects.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"kind": schema.StringAttribute{
																	Description:         "Kind of object being referenced. Values defined by this API group are 'User', 'Group', and 'ServiceAccount'.If the Authorizer does not recognized the kind value, the Authorizer should report an error.",
																	MarkdownDescription: "Kind of object being referenced. Values defined by this API group are 'User', 'Group', and 'ServiceAccount'.If the Authorizer does not recognized the kind value, the Authorizer should report an error.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name of the object being referenced.",
																	MarkdownDescription: "Name of the object being referenced.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"namespace": schema.StringAttribute{
																	Description:         "Namespace of the referenced object.  If the object kind is non-namespace, such as 'User' or 'Group', and this value is not emptythe Authorizer should report an error.",
																	MarkdownDescription: "Namespace of the referenced object.  If the object kind is non-namespace, such as 'User' or 'Group', and this value is not emptythe Authorizer should report an error.",
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"cluster_roles": schema.ListAttribute{
											Description:         "ClusterRoles is the list of cluster-wide role names for the user.",
											MarkdownDescription: "ClusterRoles is the list of cluster-wide role names for the user.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"resources": schema.SingleNestedAttribute{
											Description:         "ResourceDescription contains information about the resource being created or modified.Requires at least one tag to be specified when under MatchResources.Specifying ResourceDescription directly under match is being deprecated.Please specify under 'any' or 'all' instead.",
											MarkdownDescription: "ResourceDescription contains information about the resource being created or modified.Requires at least one tag to be specified when under MatchResources.Specifying ResourceDescription directly under match is being deprecated.Please specify under 'any' or 'all' instead.",
											Attributes: map[string]schema.Attribute{
												"annotations": schema.MapAttribute{
													Description:         "Annotations is a  map of annotations (key-value pairs of type string). Annotation keysand values support the wildcard characters '*' (matches zero or many characters) and'?' (matches at least one character).",
													MarkdownDescription: "Annotations is a  map of annotations (key-value pairs of type string). Annotation keysand values support the wildcard characters '*' (matches zero or many characters) and'?' (matches at least one character).",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"kinds": schema.ListAttribute{
													Description:         "Kinds is a list of resource kinds.",
													MarkdownDescription: "Kinds is a list of resource kinds.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name is the name of the resource. The name supports wildcard characters'*' (matches zero or many characters) and '?' (at least one character).NOTE: 'Name' is being deprecated in favor of 'Names'.",
													MarkdownDescription: "Name is the name of the resource. The name supports wildcard characters'*' (matches zero or many characters) and '?' (at least one character).NOTE: 'Name' is being deprecated in favor of 'Names'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"names": schema.ListAttribute{
													Description:         "Names are the names of the resources. Each name supports wildcard characters'*' (matches zero or many characters) and '?' (at least one character).",
													MarkdownDescription: "Names are the names of the resources. Each name supports wildcard characters'*' (matches zero or many characters) and '?' (at least one character).",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespace_selector": schema.SingleNestedAttribute{
													Description:         "NamespaceSelector is a label selector for the resource namespace. Label keys and valuesin 'matchLabels' support the wildcard characters '*' (matches zero or many characters)and '?' (matches one character).Wildcards allows writing label selectors like['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value butdoes not match an empty label set.",
													MarkdownDescription: "NamespaceSelector is a label selector for the resource namespace. Label keys and valuesin 'matchLabels' support the wildcard characters '*' (matches zero or many characters)and '?' (matches one character).Wildcards allows writing label selectors like['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value butdoes not match an empty label set.",
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
													Required: false,
													Optional: true,
													Computed: false,
												},

												"namespaces": schema.ListAttribute{
													Description:         "Namespaces is a list of namespaces names. Each name supports wildcard characters'*' (matches zero or many characters) and '?' (at least one character).",
													MarkdownDescription: "Namespaces is a list of namespaces names. Each name supports wildcard characters'*' (matches zero or many characters) and '?' (at least one character).",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"operations": schema.ListAttribute{
													Description:         "Operations can contain values ['CREATE, 'UPDATE', 'CONNECT', 'DELETE'], which are used to match a specific action.",
													MarkdownDescription: "Operations can contain values ['CREATE, 'UPDATE', 'CONNECT', 'DELETE'], which are used to match a specific action.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"selector": schema.SingleNestedAttribute{
													Description:         "Selector is a label selector. Label keys and values in 'matchLabels' support the wildcardcharacters '*' (matches zero or many characters) and '?' (matches one character).Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note thatusing ['*' : '*'] matches any key and value but does not match an empty label set.",
													MarkdownDescription: "Selector is a label selector. Label keys and values in 'matchLabels' support the wildcardcharacters '*' (matches zero or many characters) and '?' (matches one character).Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note thatusing ['*' : '*'] matches any key and value but does not match an empty label set.",
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
													Required: false,
													Optional: true,
													Computed: false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"roles": schema.ListAttribute{
											Description:         "Roles is the list of namespaced role names for the user.",
											MarkdownDescription: "Roles is the list of namespaced role names for the user.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"subjects": schema.ListNestedAttribute{
											Description:         "Subjects is the list of subject names like users, user groups, and service accounts.",
											MarkdownDescription: "Subjects is the list of subject names like users, user groups, and service accounts.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"api_group": schema.StringAttribute{
														Description:         "APIGroup holds the API group of the referenced subject.Defaults to '' for ServiceAccount subjects.Defaults to 'rbac.authorization.k8s.io' for User and Group subjects.",
														MarkdownDescription: "APIGroup holds the API group of the referenced subject.Defaults to '' for ServiceAccount subjects.Defaults to 'rbac.authorization.k8s.io' for User and Group subjects.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"kind": schema.StringAttribute{
														Description:         "Kind of object being referenced. Values defined by this API group are 'User', 'Group', and 'ServiceAccount'.If the Authorizer does not recognized the kind value, the Authorizer should report an error.",
														MarkdownDescription: "Kind of object being referenced. Values defined by this API group are 'User', 'Group', and 'ServiceAccount'.If the Authorizer does not recognized the kind value, the Authorizer should report an error.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the object being referenced.",
														MarkdownDescription: "Name of the object being referenced.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the referenced object.  If the object kind is non-namespace, such as 'User' or 'Group', and this value is not emptythe Authorizer should report an error.",
														MarkdownDescription: "Namespace of the referenced object.  If the object kind is non-namespace, such as 'User' or 'Group', and this value is not emptythe Authorizer should report an error.",
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

								"mutate": schema.SingleNestedAttribute{
									Description:         "Mutation is used to modify matching resources.",
									MarkdownDescription: "Mutation is used to modify matching resources.",
									Attributes: map[string]schema.Attribute{
										"foreach": schema.ListNestedAttribute{
											Description:         "ForEach applies mutation rules to a list of sub-elements by creating a context for each entry in the list and looping over it to apply the specified logic.",
											MarkdownDescription: "ForEach applies mutation rules to a list of sub-elements by creating a context for each entry in the list and looping over it to apply the specified logic.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"context": schema.ListNestedAttribute{
														Description:         "Context defines variables and data sources that can be used during rule execution.",
														MarkdownDescription: "Context defines variables and data sources that can be used during rule execution.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"api_call": schema.SingleNestedAttribute{
																	Description:         "APICall is an HTTP request to the Kubernetes API server, or other JSON web service.The data returned is stored in the context with the name for the context entry.",
																	MarkdownDescription: "APICall is an HTTP request to the Kubernetes API server, or other JSON web service.The data returned is stored in the context with the name for the context entry.",
																	Attributes: map[string]schema.Attribute{
																		"data": schema.ListNestedAttribute{
																			Description:         "The data object specifies the POST data sent to the server.Only applicable when the method field is set to POST.",
																			MarkdownDescription: "The data object specifies the POST data sent to the server.Only applicable when the method field is set to POST.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "Key is a unique identifier for the data value",
																						MarkdownDescription: "Key is a unique identifier for the data value",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"value": schema.MapAttribute{
																						Description:         "Value is the data value",
																						MarkdownDescription: "Value is the data value",
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

																		"jmes_path": schema.StringAttribute{
																			Description:         "JMESPath is an optional JSON Match Expression that can be used totransform the JSON response returned from the server. For examplea JMESPath of 'items | length(@)' applied to the API server responsefor the URLPath '/apis/apps/v1/deployments' will return the total countof deployments across all namespaces.",
																			MarkdownDescription: "JMESPath is an optional JSON Match Expression that can be used totransform the JSON response returned from the server. For examplea JMESPath of 'items | length(@)' applied to the API server responsefor the URLPath '/apis/apps/v1/deployments' will return the total countof deployments across all namespaces.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"method": schema.StringAttribute{
																			Description:         "Method is the HTTP request type (GET or POST). Defaults to GET.",
																			MarkdownDescription: "Method is the HTTP request type (GET or POST). Defaults to GET.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.OneOf("GET", "POST"),
																			},
																		},

																		"service": schema.SingleNestedAttribute{
																			Description:         "Service is an API call to a JSON web service.This is used for non-Kubernetes API server calls.It's mutually exclusive with the URLPath field.",
																			MarkdownDescription: "Service is an API call to a JSON web service.This is used for non-Kubernetes API server calls.It's mutually exclusive with the URLPath field.",
																			Attributes: map[string]schema.Attribute{
																				"ca_bundle": schema.StringAttribute{
																					Description:         "CABundle is a PEM encoded CA bundle which will be used to validatethe server certificate.",
																					MarkdownDescription: "CABundle is a PEM encoded CA bundle which will be used to validatethe server certificate.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"url": schema.StringAttribute{
																					Description:         "URL is the JSON web service URL. A typical form is'https://{service}.{namespace}:{port}/{path}'.",
																					MarkdownDescription: "URL is the JSON web service URL. A typical form is'https://{service}.{namespace}:{port}/{path}'.",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"url_path": schema.StringAttribute{
																			Description:         "URLPath is the URL path to be used in the HTTP GET or POST request to theKubernetes API server (e.g. '/api/v1/namespaces' or  '/apis/apps/v1/deployments').The format required is the same format used by the 'kubectl get --raw' command.See https://kyverno.io/docs/writing-policies/external-data-sources/#variables-from-kubernetes-api-server-callsfor details.It's mutually exclusive with the Service field.",
																			MarkdownDescription: "URLPath is the URL path to be used in the HTTP GET or POST request to theKubernetes API server (e.g. '/api/v1/namespaces' or  '/apis/apps/v1/deployments').The format required is the same format used by the 'kubectl get --raw' command.See https://kyverno.io/docs/writing-policies/external-data-sources/#variables-from-kubernetes-api-server-callsfor details.It's mutually exclusive with the Service field.",
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
																	Description:         "ConfigMap is the ConfigMap reference.",
																	MarkdownDescription: "ConfigMap is the ConfigMap reference.",
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Description:         "Name is the ConfigMap name.",
																			MarkdownDescription: "Name is the ConfigMap name.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"namespace": schema.StringAttribute{
																			Description:         "Namespace is the ConfigMap namespace.",
																			MarkdownDescription: "Namespace is the ConfigMap namespace.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"global_reference": schema.SingleNestedAttribute{
																	Description:         "GlobalContextEntryReference is a reference to a cached global context entry.",
																	MarkdownDescription: "GlobalContextEntryReference is a reference to a cached global context entry.",
																	Attributes: map[string]schema.Attribute{
																		"jmes_path": schema.StringAttribute{
																			Description:         "JMESPath is an optional JSON Match Expression that can be used totransform the JSON response returned from the server. For examplea JMESPath of 'items | length(@)' applied to the API server responsefor the URLPath '/apis/apps/v1/deployments' will return the total countof deployments across all namespaces.",
																			MarkdownDescription: "JMESPath is an optional JSON Match Expression that can be used totransform the JSON response returned from the server. For examplea JMESPath of 'items | length(@)' applied to the API server responsefor the URLPath '/apis/apps/v1/deployments' will return the total countof deployments across all namespaces.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"name": schema.StringAttribute{
																			Description:         "Name of the global context entry",
																			MarkdownDescription: "Name of the global context entry",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"image_registry": schema.SingleNestedAttribute{
																	Description:         "ImageRegistry defines requests to an OCI/Docker V2 registry to fetch imagedetails.",
																	MarkdownDescription: "ImageRegistry defines requests to an OCI/Docker V2 registry to fetch imagedetails.",
																	Attributes: map[string]schema.Attribute{
																		"image_registry_credentials": schema.SingleNestedAttribute{
																			Description:         "ImageRegistryCredentials provides credentials that will be used for authentication with registry",
																			MarkdownDescription: "ImageRegistryCredentials provides credentials that will be used for authentication with registry",
																			Attributes: map[string]schema.Attribute{
																				"allow_insecure_registry": schema.BoolAttribute{
																					Description:         "AllowInsecureRegistry allows insecure access to a registry.",
																					MarkdownDescription: "AllowInsecureRegistry allows insecure access to a registry.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"providers": schema.ListAttribute{
																					Description:         "Providers specifies a list of OCI Registry names, whose authentication providers are provided.It can be of one of these values: default,google,azure,amazon,github.",
																					MarkdownDescription: "Providers specifies a list of OCI Registry names, whose authentication providers are provided.It can be of one of these values: default,google,azure,amazon,github.",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"secrets": schema.ListAttribute{
																					Description:         "Secrets specifies a list of secrets that are provided for credentials.Secrets must live in the Kyverno namespace.",
																					MarkdownDescription: "Secrets specifies a list of secrets that are provided for credentials.Secrets must live in the Kyverno namespace.",
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

																		"jmes_path": schema.StringAttribute{
																			Description:         "JMESPath is an optional JSON Match Expression that can be used totransform the ImageData struct returned as a result of processingthe image reference.",
																			MarkdownDescription: "JMESPath is an optional JSON Match Expression that can be used totransform the ImageData struct returned as a result of processingthe image reference.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"reference": schema.StringAttribute{
																			Description:         "Reference is image reference to a container image in the registry.Example: ghcr.io/kyverno/kyverno:latest",
																			MarkdownDescription: "Reference is image reference to a container image in the registry.Example: ghcr.io/kyverno/kyverno:latest",
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
																	Description:         "Name is the variable name.",
																	MarkdownDescription: "Name is the variable name.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"variable": schema.SingleNestedAttribute{
																	Description:         "Variable defines an arbitrary JMESPath context variable that can be defined inline.",
																	MarkdownDescription: "Variable defines an arbitrary JMESPath context variable that can be defined inline.",
																	Attributes: map[string]schema.Attribute{
																		"default": schema.MapAttribute{
																			Description:         "Default is an optional arbitrary JSON object that the variable may take if the JMESPathexpression evaluates to nil",
																			MarkdownDescription: "Default is an optional arbitrary JSON object that the variable may take if the JMESPathexpression evaluates to nil",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"jmes_path": schema.StringAttribute{
																			Description:         "JMESPath is an optional JMESPath Expression that can be used totransform the variable.",
																			MarkdownDescription: "JMESPath is an optional JMESPath Expression that can be used totransform the variable.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"value": schema.MapAttribute{
																			Description:         "Value is any arbitrary JSON object representable in YAML or JSON form.",
																			MarkdownDescription: "Value is any arbitrary JSON object representable in YAML or JSON form.",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"foreach": schema.MapAttribute{
														Description:         "Foreach declares a nested foreach iterator",
														MarkdownDescription: "Foreach declares a nested foreach iterator",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"list": schema.StringAttribute{
														Description:         "List specifies a JMESPath expression that results in one or more elementsto which the validation logic is applied.",
														MarkdownDescription: "List specifies a JMESPath expression that results in one or more elementsto which the validation logic is applied.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"order": schema.StringAttribute{
														Description:         "Order defines the iteration order on the list.Can be Ascending to iterate from first to last element or Descending to iterate in from last to first element.",
														MarkdownDescription: "Order defines the iteration order on the list.Can be Ascending to iterate from first to last element or Descending to iterate in from last to first element.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("Ascending", "Descending"),
														},
													},

													"patch_strategic_merge": schema.MapAttribute{
														Description:         "PatchStrategicMerge is a strategic merge patch used to modify resources.See https://kubernetes.io/docs/tasks/manage-kubernetes-objects/update-api-object-kubectl-patch/and https://kubectl.docs.kubernetes.io/references/kustomize/patchesstrategicmerge/.",
														MarkdownDescription: "PatchStrategicMerge is a strategic merge patch used to modify resources.See https://kubernetes.io/docs/tasks/manage-kubernetes-objects/update-api-object-kubectl-patch/and https://kubectl.docs.kubernetes.io/references/kustomize/patchesstrategicmerge/.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"patches_json6902": schema.StringAttribute{
														Description:         "PatchesJSON6902 is a list of RFC 6902 JSON Patch declarations used to modify resources.See https://tools.ietf.org/html/rfc6902 and https://kubectl.docs.kubernetes.io/references/kustomize/patchesjson6902/.",
														MarkdownDescription: "PatchesJSON6902 is a list of RFC 6902 JSON Patch declarations used to modify resources.See https://tools.ietf.org/html/rfc6902 and https://kubectl.docs.kubernetes.io/references/kustomize/patchesjson6902/.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"preconditions": schema.SingleNestedAttribute{
														Description:         "AnyAllConditions are used to determine if a policy rule should be applied by evaluating aset of conditions. The declaration can contain nested 'any' or 'all' statements.See: https://kyverno.io/docs/writing-policies/preconditions/",
														MarkdownDescription: "AnyAllConditions are used to determine if a policy rule should be applied by evaluating aset of conditions. The declaration can contain nested 'any' or 'all' statements.See: https://kyverno.io/docs/writing-policies/preconditions/",
														Attributes: map[string]schema.Attribute{
															"all": schema.ListNestedAttribute{
																Description:         "AllConditions enable variable-based conditional rule execution. This is useful forfiner control of when an rule is applied. A condition can reference object datausing JMESPath notation.Here, all of the conditions need to pass",
																MarkdownDescription: "AllConditions enable variable-based conditional rule execution. This is useful forfiner control of when an rule is applied. A condition can reference object datausing JMESPath notation.Here, all of the conditions need to pass",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"key": schema.MapAttribute{
																			Description:         "Key is the context entry (using JMESPath) for conditional rule evaluation.",
																			MarkdownDescription: "Key is the context entry (using JMESPath) for conditional rule evaluation.",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"message": schema.StringAttribute{
																			Description:         "Message is an optional display message",
																			MarkdownDescription: "Message is an optional display message",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"operator": schema.StringAttribute{
																			Description:         "Operator is the conditional operation to perform. Valid operators are:Equals, NotEquals, In, AnyIn, AllIn, NotIn, AnyNotIn, AllNotIn, GreaterThanOrEquals,GreaterThan, LessThanOrEquals, LessThan, DurationGreaterThanOrEquals, DurationGreaterThan,DurationLessThanOrEquals, DurationLessThan",
																			MarkdownDescription: "Operator is the conditional operation to perform. Valid operators are:Equals, NotEquals, In, AnyIn, AllIn, NotIn, AnyNotIn, AllNotIn, GreaterThanOrEquals,GreaterThan, LessThanOrEquals, LessThan, DurationGreaterThanOrEquals, DurationGreaterThan,DurationLessThanOrEquals, DurationLessThan",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.OneOf("Equals", "NotEquals", "In", "AnyIn", "AllIn", "NotIn", "AnyNotIn", "AllNotIn", "GreaterThanOrEquals", "GreaterThan", "LessThanOrEquals", "LessThan", "DurationGreaterThanOrEquals", "DurationGreaterThan", "DurationLessThanOrEquals", "DurationLessThan"),
																			},
																		},

																		"value": schema.MapAttribute{
																			Description:         "Value is the conditional value, or set of values. The values can be fixed setor can be variables declared using JMESPath.",
																			MarkdownDescription: "Value is the conditional value, or set of values. The values can be fixed setor can be variables declared using JMESPath.",
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

															"any": schema.ListNestedAttribute{
																Description:         "AnyConditions enable variable-based conditional rule execution. This is useful forfiner control of when an rule is applied. A condition can reference object datausing JMESPath notation.Here, at least one of the conditions need to pass",
																MarkdownDescription: "AnyConditions enable variable-based conditional rule execution. This is useful forfiner control of when an rule is applied. A condition can reference object datausing JMESPath notation.Here, at least one of the conditions need to pass",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"key": schema.MapAttribute{
																			Description:         "Key is the context entry (using JMESPath) for conditional rule evaluation.",
																			MarkdownDescription: "Key is the context entry (using JMESPath) for conditional rule evaluation.",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"message": schema.StringAttribute{
																			Description:         "Message is an optional display message",
																			MarkdownDescription: "Message is an optional display message",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"operator": schema.StringAttribute{
																			Description:         "Operator is the conditional operation to perform. Valid operators are:Equals, NotEquals, In, AnyIn, AllIn, NotIn, AnyNotIn, AllNotIn, GreaterThanOrEquals,GreaterThan, LessThanOrEquals, LessThan, DurationGreaterThanOrEquals, DurationGreaterThan,DurationLessThanOrEquals, DurationLessThan",
																			MarkdownDescription: "Operator is the conditional operation to perform. Valid operators are:Equals, NotEquals, In, AnyIn, AllIn, NotIn, AnyNotIn, AllNotIn, GreaterThanOrEquals,GreaterThan, LessThanOrEquals, LessThan, DurationGreaterThanOrEquals, DurationGreaterThan,DurationLessThanOrEquals, DurationLessThan",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.OneOf("Equals", "NotEquals", "In", "AnyIn", "AllIn", "NotIn", "AnyNotIn", "AllNotIn", "GreaterThanOrEquals", "GreaterThan", "LessThanOrEquals", "LessThan", "DurationGreaterThanOrEquals", "DurationGreaterThan", "DurationLessThanOrEquals", "DurationLessThan"),
																			},
																		},

																		"value": schema.MapAttribute{
																			Description:         "Value is the conditional value, or set of values. The values can be fixed setor can be variables declared using JMESPath.",
																			MarkdownDescription: "Value is the conditional value, or set of values. The values can be fixed setor can be variables declared using JMESPath.",
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

										"mutate_existing_on_policy_update": schema.BoolAttribute{
											Description:         "MutateExistingOnPolicyUpdate controls if the mutateExisting rule will be applied on policy events.",
											MarkdownDescription: "MutateExistingOnPolicyUpdate controls if the mutateExisting rule will be applied on policy events.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"patch_strategic_merge": schema.MapAttribute{
											Description:         "PatchStrategicMerge is a strategic merge patch used to modify resources.See https://kubernetes.io/docs/tasks/manage-kubernetes-objects/update-api-object-kubectl-patch/and https://kubectl.docs.kubernetes.io/references/kustomize/patchesstrategicmerge/.",
											MarkdownDescription: "PatchStrategicMerge is a strategic merge patch used to modify resources.See https://kubernetes.io/docs/tasks/manage-kubernetes-objects/update-api-object-kubectl-patch/and https://kubectl.docs.kubernetes.io/references/kustomize/patchesstrategicmerge/.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"patches_json6902": schema.StringAttribute{
											Description:         "PatchesJSON6902 is a list of RFC 6902 JSON Patch declarations used to modify resources.See https://tools.ietf.org/html/rfc6902 and https://kubectl.docs.kubernetes.io/references/kustomize/patchesjson6902/.",
											MarkdownDescription: "PatchesJSON6902 is a list of RFC 6902 JSON Patch declarations used to modify resources.See https://tools.ietf.org/html/rfc6902 and https://kubectl.docs.kubernetes.io/references/kustomize/patchesjson6902/.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"targets": schema.ListNestedAttribute{
											Description:         "Targets defines the target resources to be mutated.",
											MarkdownDescription: "Targets defines the target resources to be mutated.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"api_version": schema.StringAttribute{
														Description:         "APIVersion specifies resource apiVersion.",
														MarkdownDescription: "APIVersion specifies resource apiVersion.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"context": schema.ListNestedAttribute{
														Description:         "Context defines variables and data sources that can be used during rule execution.",
														MarkdownDescription: "Context defines variables and data sources that can be used during rule execution.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"api_call": schema.SingleNestedAttribute{
																	Description:         "APICall is an HTTP request to the Kubernetes API server, or other JSON web service.The data returned is stored in the context with the name for the context entry.",
																	MarkdownDescription: "APICall is an HTTP request to the Kubernetes API server, or other JSON web service.The data returned is stored in the context with the name for the context entry.",
																	Attributes: map[string]schema.Attribute{
																		"data": schema.ListNestedAttribute{
																			Description:         "The data object specifies the POST data sent to the server.Only applicable when the method field is set to POST.",
																			MarkdownDescription: "The data object specifies the POST data sent to the server.Only applicable when the method field is set to POST.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "Key is a unique identifier for the data value",
																						MarkdownDescription: "Key is a unique identifier for the data value",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"value": schema.MapAttribute{
																						Description:         "Value is the data value",
																						MarkdownDescription: "Value is the data value",
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

																		"jmes_path": schema.StringAttribute{
																			Description:         "JMESPath is an optional JSON Match Expression that can be used totransform the JSON response returned from the server. For examplea JMESPath of 'items | length(@)' applied to the API server responsefor the URLPath '/apis/apps/v1/deployments' will return the total countof deployments across all namespaces.",
																			MarkdownDescription: "JMESPath is an optional JSON Match Expression that can be used totransform the JSON response returned from the server. For examplea JMESPath of 'items | length(@)' applied to the API server responsefor the URLPath '/apis/apps/v1/deployments' will return the total countof deployments across all namespaces.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"method": schema.StringAttribute{
																			Description:         "Method is the HTTP request type (GET or POST). Defaults to GET.",
																			MarkdownDescription: "Method is the HTTP request type (GET or POST). Defaults to GET.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.OneOf("GET", "POST"),
																			},
																		},

																		"service": schema.SingleNestedAttribute{
																			Description:         "Service is an API call to a JSON web service.This is used for non-Kubernetes API server calls.It's mutually exclusive with the URLPath field.",
																			MarkdownDescription: "Service is an API call to a JSON web service.This is used for non-Kubernetes API server calls.It's mutually exclusive with the URLPath field.",
																			Attributes: map[string]schema.Attribute{
																				"ca_bundle": schema.StringAttribute{
																					Description:         "CABundle is a PEM encoded CA bundle which will be used to validatethe server certificate.",
																					MarkdownDescription: "CABundle is a PEM encoded CA bundle which will be used to validatethe server certificate.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"url": schema.StringAttribute{
																					Description:         "URL is the JSON web service URL. A typical form is'https://{service}.{namespace}:{port}/{path}'.",
																					MarkdownDescription: "URL is the JSON web service URL. A typical form is'https://{service}.{namespace}:{port}/{path}'.",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"url_path": schema.StringAttribute{
																			Description:         "URLPath is the URL path to be used in the HTTP GET or POST request to theKubernetes API server (e.g. '/api/v1/namespaces' or  '/apis/apps/v1/deployments').The format required is the same format used by the 'kubectl get --raw' command.See https://kyverno.io/docs/writing-policies/external-data-sources/#variables-from-kubernetes-api-server-callsfor details.It's mutually exclusive with the Service field.",
																			MarkdownDescription: "URLPath is the URL path to be used in the HTTP GET or POST request to theKubernetes API server (e.g. '/api/v1/namespaces' or  '/apis/apps/v1/deployments').The format required is the same format used by the 'kubectl get --raw' command.See https://kyverno.io/docs/writing-policies/external-data-sources/#variables-from-kubernetes-api-server-callsfor details.It's mutually exclusive with the Service field.",
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
																	Description:         "ConfigMap is the ConfigMap reference.",
																	MarkdownDescription: "ConfigMap is the ConfigMap reference.",
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Description:         "Name is the ConfigMap name.",
																			MarkdownDescription: "Name is the ConfigMap name.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"namespace": schema.StringAttribute{
																			Description:         "Namespace is the ConfigMap namespace.",
																			MarkdownDescription: "Namespace is the ConfigMap namespace.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"global_reference": schema.SingleNestedAttribute{
																	Description:         "GlobalContextEntryReference is a reference to a cached global context entry.",
																	MarkdownDescription: "GlobalContextEntryReference is a reference to a cached global context entry.",
																	Attributes: map[string]schema.Attribute{
																		"jmes_path": schema.StringAttribute{
																			Description:         "JMESPath is an optional JSON Match Expression that can be used totransform the JSON response returned from the server. For examplea JMESPath of 'items | length(@)' applied to the API server responsefor the URLPath '/apis/apps/v1/deployments' will return the total countof deployments across all namespaces.",
																			MarkdownDescription: "JMESPath is an optional JSON Match Expression that can be used totransform the JSON response returned from the server. For examplea JMESPath of 'items | length(@)' applied to the API server responsefor the URLPath '/apis/apps/v1/deployments' will return the total countof deployments across all namespaces.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"name": schema.StringAttribute{
																			Description:         "Name of the global context entry",
																			MarkdownDescription: "Name of the global context entry",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"image_registry": schema.SingleNestedAttribute{
																	Description:         "ImageRegistry defines requests to an OCI/Docker V2 registry to fetch imagedetails.",
																	MarkdownDescription: "ImageRegistry defines requests to an OCI/Docker V2 registry to fetch imagedetails.",
																	Attributes: map[string]schema.Attribute{
																		"image_registry_credentials": schema.SingleNestedAttribute{
																			Description:         "ImageRegistryCredentials provides credentials that will be used for authentication with registry",
																			MarkdownDescription: "ImageRegistryCredentials provides credentials that will be used for authentication with registry",
																			Attributes: map[string]schema.Attribute{
																				"allow_insecure_registry": schema.BoolAttribute{
																					Description:         "AllowInsecureRegistry allows insecure access to a registry.",
																					MarkdownDescription: "AllowInsecureRegistry allows insecure access to a registry.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"providers": schema.ListAttribute{
																					Description:         "Providers specifies a list of OCI Registry names, whose authentication providers are provided.It can be of one of these values: default,google,azure,amazon,github.",
																					MarkdownDescription: "Providers specifies a list of OCI Registry names, whose authentication providers are provided.It can be of one of these values: default,google,azure,amazon,github.",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"secrets": schema.ListAttribute{
																					Description:         "Secrets specifies a list of secrets that are provided for credentials.Secrets must live in the Kyverno namespace.",
																					MarkdownDescription: "Secrets specifies a list of secrets that are provided for credentials.Secrets must live in the Kyverno namespace.",
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

																		"jmes_path": schema.StringAttribute{
																			Description:         "JMESPath is an optional JSON Match Expression that can be used totransform the ImageData struct returned as a result of processingthe image reference.",
																			MarkdownDescription: "JMESPath is an optional JSON Match Expression that can be used totransform the ImageData struct returned as a result of processingthe image reference.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"reference": schema.StringAttribute{
																			Description:         "Reference is image reference to a container image in the registry.Example: ghcr.io/kyverno/kyverno:latest",
																			MarkdownDescription: "Reference is image reference to a container image in the registry.Example: ghcr.io/kyverno/kyverno:latest",
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
																	Description:         "Name is the variable name.",
																	MarkdownDescription: "Name is the variable name.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"variable": schema.SingleNestedAttribute{
																	Description:         "Variable defines an arbitrary JMESPath context variable that can be defined inline.",
																	MarkdownDescription: "Variable defines an arbitrary JMESPath context variable that can be defined inline.",
																	Attributes: map[string]schema.Attribute{
																		"default": schema.MapAttribute{
																			Description:         "Default is an optional arbitrary JSON object that the variable may take if the JMESPathexpression evaluates to nil",
																			MarkdownDescription: "Default is an optional arbitrary JSON object that the variable may take if the JMESPathexpression evaluates to nil",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"jmes_path": schema.StringAttribute{
																			Description:         "JMESPath is an optional JMESPath Expression that can be used totransform the variable.",
																			MarkdownDescription: "JMESPath is an optional JMESPath Expression that can be used totransform the variable.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"value": schema.MapAttribute{
																			Description:         "Value is any arbitrary JSON object representable in YAML or JSON form.",
																			MarkdownDescription: "Value is any arbitrary JSON object representable in YAML or JSON form.",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"kind": schema.StringAttribute{
														Description:         "Kind specifies resource kind.",
														MarkdownDescription: "Kind specifies resource kind.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name specifies the resource name.",
														MarkdownDescription: "Name specifies the resource name.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace specifies resource namespace.",
														MarkdownDescription: "Namespace specifies resource namespace.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"preconditions": schema.MapAttribute{
														Description:         "Preconditions are used to determine if a policy rule should be applied by evaluating aset of conditions. The declaration can contain nested 'any' or 'all' statements. A direct listof conditions (without 'any' or 'all' statements is supported for backwards compatibility butwill be deprecated in the next major release.See: https://kyverno.io/docs/writing-policies/preconditions/",
														MarkdownDescription: "Preconditions are used to determine if a policy rule should be applied by evaluating aset of conditions. The declaration can contain nested 'any' or 'all' statements. A direct listof conditions (without 'any' or 'all' statements is supported for backwards compatibility butwill be deprecated in the next major release.See: https://kyverno.io/docs/writing-policies/preconditions/",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"uid": schema.StringAttribute{
														Description:         "UID specifies the resource uid.",
														MarkdownDescription: "UID specifies the resource uid.",
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

								"name": schema.StringAttribute{
									Description:         "Name is a label to identify the rule, It must be unique within the policy.",
									MarkdownDescription: "Name is a label to identify the rule, It must be unique within the policy.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(63),
									},
								},

								"preconditions": schema.MapAttribute{
									Description:         "Preconditions are used to determine if a policy rule should be applied by evaluating aset of conditions. The declaration can contain nested 'any' or 'all' statements. A direct listof conditions (without 'any' or 'all' statements is supported for backwards compatibility butwill be deprecated in the next major release.See: https://kyverno.io/docs/writing-policies/preconditions/",
									MarkdownDescription: "Preconditions are used to determine if a policy rule should be applied by evaluating aset of conditions. The declaration can contain nested 'any' or 'all' statements. A direct listof conditions (without 'any' or 'all' statements is supported for backwards compatibility butwill be deprecated in the next major release.See: https://kyverno.io/docs/writing-policies/preconditions/",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"skip_background_requests": schema.BoolAttribute{
									Description:         "SkipBackgroundRequests bypasses admission requests that are sent by the background controller.The default value is set to 'true', it must be set to 'false' to applygenerate and mutateExisting rules to those requests.",
									MarkdownDescription: "SkipBackgroundRequests bypasses admission requests that are sent by the background controller.The default value is set to 'true', it must be set to 'false' to applygenerate and mutateExisting rules to those requests.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"validate": schema.SingleNestedAttribute{
									Description:         "Validation is used to validate matching resources.",
									MarkdownDescription: "Validation is used to validate matching resources.",
									Attributes: map[string]schema.Attribute{
										"any_pattern": schema.MapAttribute{
											Description:         "AnyPattern specifies list of validation patterns. At least one of the patternsmust be satisfied for the validation rule to succeed.",
											MarkdownDescription: "AnyPattern specifies list of validation patterns. At least one of the patternsmust be satisfied for the validation rule to succeed.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"cel": schema.SingleNestedAttribute{
											Description:         "CEL allows validation checks using the Common Expression Language (https://kubernetes.io/docs/reference/using-api/cel/).",
											MarkdownDescription: "CEL allows validation checks using the Common Expression Language (https://kubernetes.io/docs/reference/using-api/cel/).",
											Attributes: map[string]schema.Attribute{
												"audit_annotations": schema.ListNestedAttribute{
													Description:         "AuditAnnotations contains CEL expressions which are used to produce audit annotations for the audit event of the API request.",
													MarkdownDescription: "AuditAnnotations contains CEL expressions which are used to produce audit annotations for the audit event of the API request.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "key specifies the audit annotation key. The audit annotation keys ofa ValidatingAdmissionPolicy must be unique. The key must be a qualifiedname ([A-Za-z0-9][-A-Za-z0-9_.]*) no more than 63 bytes in length.The key is combined with the resource name of theValidatingAdmissionPolicy to construct an audit annotation key:'{ValidatingAdmissionPolicy name}/{key}'.If an admission webhook uses the same resource name as this ValidatingAdmissionPolicyand the same audit annotation key, the annotation key will be identical.In this case, the first annotation written with the key will be includedin the audit event and all subsequent annotations with the same keywill be discarded.Required.",
																MarkdownDescription: "key specifies the audit annotation key. The audit annotation keys ofa ValidatingAdmissionPolicy must be unique. The key must be a qualifiedname ([A-Za-z0-9][-A-Za-z0-9_.]*) no more than 63 bytes in length.The key is combined with the resource name of theValidatingAdmissionPolicy to construct an audit annotation key:'{ValidatingAdmissionPolicy name}/{key}'.If an admission webhook uses the same resource name as this ValidatingAdmissionPolicyand the same audit annotation key, the annotation key will be identical.In this case, the first annotation written with the key will be includedin the audit event and all subsequent annotations with the same keywill be discarded.Required.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"value_expression": schema.StringAttribute{
																Description:         "valueExpression represents the expression which is evaluated by CEL toproduce an audit annotation value. The expression must evaluate to eithera string or null value. If the expression evaluates to a string, theaudit annotation is included with the string value. If the expressionevaluates to null or empty string the audit annotation will be omitted.The valueExpression may be no longer than 5kb in length.If the result of the valueExpression is more than 10kb in length, itwill be truncated to 10kb.If multiple ValidatingAdmissionPolicyBinding resources match anAPI request, then the valueExpression will be evaluated foreach binding. All unique values produced by the valueExpressionswill be joined together in a comma-separated list.Required.",
																MarkdownDescription: "valueExpression represents the expression which is evaluated by CEL toproduce an audit annotation value. The expression must evaluate to eithera string or null value. If the expression evaluates to a string, theaudit annotation is included with the string value. If the expressionevaluates to null or empty string the audit annotation will be omitted.The valueExpression may be no longer than 5kb in length.If the result of the valueExpression is more than 10kb in length, itwill be truncated to 10kb.If multiple ValidatingAdmissionPolicyBinding resources match anAPI request, then the valueExpression will be evaluated foreach binding. All unique values produced by the valueExpressionswill be joined together in a comma-separated list.Required.",
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

												"expressions": schema.ListNestedAttribute{
													Description:         "Expressions is a list of CELExpression types.",
													MarkdownDescription: "Expressions is a list of CELExpression types.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"expression": schema.StringAttribute{
																Description:         "Expression represents the expression which will be evaluated by CEL.ref: https://github.com/google/cel-specCEL expressions have access to the contents of the API request/response, organized into CEL variables as well as some other useful variables:- 'object' - The object from the incoming request. The value is null for DELETE requests.- 'oldObject' - The existing object. The value is null for CREATE requests.- 'request' - Attributes of the API request([ref](/pkg/apis/admission/types.go#AdmissionRequest)).- 'params' - Parameter resource referred to by the policy binding being evaluated. Only populated if the policy has a ParamKind.- 'namespaceObject' - The namespace object that the incoming object belongs to. The value is null for cluster-scoped resources.- 'variables' - Map of composited variables, from its name to its lazily evaluated value.  For example, a variable named 'foo' can be accessed as 'variables.foo'.- 'authorizer' - A CEL Authorizer. May be used to perform authorization checks for the principal (user or service account) of the request.  See https://pkg.go.dev/k8s.io/apiserver/pkg/cel/library#Authz- 'authorizer.requestResource' - A CEL ResourceCheck constructed from the 'authorizer' and configured with the  request resource.The 'apiVersion', 'kind', 'metadata.name' and 'metadata.generateName' are always accessible from the root of theobject. No other metadata properties are accessible.Only property names of the form '[a-zA-Z_.-/][a-zA-Z0-9_.-/]*' are accessible.Accessible property names are escaped according to the following rules when accessed in the expression:- '__' escapes to '__underscores__'- '.' escapes to '__dot__'- '-' escapes to '__dash__'- '/' escapes to '__slash__'- Property names that exactly match a CEL RESERVED keyword escape to '__{keyword}__'. The keywords are:	  'true', 'false', 'null', 'in', 'as', 'break', 'const', 'continue', 'else', 'for', 'function', 'if',	  'import', 'let', 'loop', 'package', 'namespace', 'return'.Examples:  - Expression accessing a property named 'namespace': {'Expression': 'object.__namespace__ > 0'}  - Expression accessing a property named 'x-prop': {'Expression': 'object.x__dash__prop > 0'}  - Expression accessing a property named 'redact__d': {'Expression': 'object.redact__underscores__d > 0'}Equality on arrays with list type of 'set' or 'map' ignores element order, i.e. [1, 2] == [2, 1].Concatenation on arrays with x-kubernetes-list-type use the semantics of the list type:  - 'set': 'X + Y' performs a union where the array positions of all elements in 'X' are preserved and    non-intersecting elements in 'Y' are appended, retaining their partial order.  - 'map': 'X + Y' performs a merge where the array positions of all keys in 'X' are preserved but the values    are overwritten by values in 'Y' when the key sets of 'X' and 'Y' intersect. Elements in 'Y' with    non-intersecting keys are appended, retaining their partial order.Required.",
																MarkdownDescription: "Expression represents the expression which will be evaluated by CEL.ref: https://github.com/google/cel-specCEL expressions have access to the contents of the API request/response, organized into CEL variables as well as some other useful variables:- 'object' - The object from the incoming request. The value is null for DELETE requests.- 'oldObject' - The existing object. The value is null for CREATE requests.- 'request' - Attributes of the API request([ref](/pkg/apis/admission/types.go#AdmissionRequest)).- 'params' - Parameter resource referred to by the policy binding being evaluated. Only populated if the policy has a ParamKind.- 'namespaceObject' - The namespace object that the incoming object belongs to. The value is null for cluster-scoped resources.- 'variables' - Map of composited variables, from its name to its lazily evaluated value.  For example, a variable named 'foo' can be accessed as 'variables.foo'.- 'authorizer' - A CEL Authorizer. May be used to perform authorization checks for the principal (user or service account) of the request.  See https://pkg.go.dev/k8s.io/apiserver/pkg/cel/library#Authz- 'authorizer.requestResource' - A CEL ResourceCheck constructed from the 'authorizer' and configured with the  request resource.The 'apiVersion', 'kind', 'metadata.name' and 'metadata.generateName' are always accessible from the root of theobject. No other metadata properties are accessible.Only property names of the form '[a-zA-Z_.-/][a-zA-Z0-9_.-/]*' are accessible.Accessible property names are escaped according to the following rules when accessed in the expression:- '__' escapes to '__underscores__'- '.' escapes to '__dot__'- '-' escapes to '__dash__'- '/' escapes to '__slash__'- Property names that exactly match a CEL RESERVED keyword escape to '__{keyword}__'. The keywords are:	  'true', 'false', 'null', 'in', 'as', 'break', 'const', 'continue', 'else', 'for', 'function', 'if',	  'import', 'let', 'loop', 'package', 'namespace', 'return'.Examples:  - Expression accessing a property named 'namespace': {'Expression': 'object.__namespace__ > 0'}  - Expression accessing a property named 'x-prop': {'Expression': 'object.x__dash__prop > 0'}  - Expression accessing a property named 'redact__d': {'Expression': 'object.redact__underscores__d > 0'}Equality on arrays with list type of 'set' or 'map' ignores element order, i.e. [1, 2] == [2, 1].Concatenation on arrays with x-kubernetes-list-type use the semantics of the list type:  - 'set': 'X + Y' performs a union where the array positions of all elements in 'X' are preserved and    non-intersecting elements in 'Y' are appended, retaining their partial order.  - 'map': 'X + Y' performs a merge where the array positions of all keys in 'X' are preserved but the values    are overwritten by values in 'Y' when the key sets of 'X' and 'Y' intersect. Elements in 'Y' with    non-intersecting keys are appended, retaining their partial order.Required.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"message": schema.StringAttribute{
																Description:         "Message represents the message displayed when validation fails. The message is required if the Expression containsline breaks. The message must not contain line breaks.If unset, the message is 'failed rule: {Rule}'.e.g. 'must be a URL with the host matching spec.host'If the Expression contains line breaks. Message is required.The message must not contain line breaks.If unset, the message is 'failed Expression: {Expression}'.",
																MarkdownDescription: "Message represents the message displayed when validation fails. The message is required if the Expression containsline breaks. The message must not contain line breaks.If unset, the message is 'failed rule: {Rule}'.e.g. 'must be a URL with the host matching spec.host'If the Expression contains line breaks. Message is required.The message must not contain line breaks.If unset, the message is 'failed Expression: {Expression}'.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"message_expression": schema.StringAttribute{
																Description:         "messageExpression declares a CEL expression that evaluates to the validation failure message that is returned when this rule fails.Since messageExpression is used as a failure message, it must evaluate to a string.If both message and messageExpression are present on a validation, then messageExpression will be used if validation fails.If messageExpression results in a runtime error, the runtime error is logged, and the validation failure message is producedas if the messageExpression field were unset. If messageExpression evaluates to an empty string, a string with only spaces, or a stringthat contains line breaks, then the validation failure message will also be produced as if the messageExpression field were unset, andthe fact that messageExpression produced an empty string/string with only spaces/string with line breaks will be logged.messageExpression has access to all the same variables as the 'expression' except for 'authorizer' and 'authorizer.requestResource'.Example:'object.x must be less than max ('+string(params.max)+')'",
																MarkdownDescription: "messageExpression declares a CEL expression that evaluates to the validation failure message that is returned when this rule fails.Since messageExpression is used as a failure message, it must evaluate to a string.If both message and messageExpression are present on a validation, then messageExpression will be used if validation fails.If messageExpression results in a runtime error, the runtime error is logged, and the validation failure message is producedas if the messageExpression field were unset. If messageExpression evaluates to an empty string, a string with only spaces, or a stringthat contains line breaks, then the validation failure message will also be produced as if the messageExpression field were unset, andthe fact that messageExpression produced an empty string/string with only spaces/string with line breaks will be logged.messageExpression has access to all the same variables as the 'expression' except for 'authorizer' and 'authorizer.requestResource'.Example:'object.x must be less than max ('+string(params.max)+')'",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"reason": schema.StringAttribute{
																Description:         "Reason represents a machine-readable description of why this validation failed.If this is the first validation in the list to fail, this reason, as well as thecorresponding HTTP response code, are used in theHTTP response to the client.The currently supported reasons are: 'Unauthorized', 'Forbidden', 'Invalid', 'RequestEntityTooLarge'.If not set, StatusReasonInvalid is used in the response to the client.",
																MarkdownDescription: "Reason represents a machine-readable description of why this validation failed.If this is the first validation in the list to fail, this reason, as well as thecorresponding HTTP response code, are used in theHTTP response to the client.The currently supported reasons are: 'Unauthorized', 'Forbidden', 'Invalid', 'RequestEntityTooLarge'.If not set, StatusReasonInvalid is used in the response to the client.",
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

												"param_kind": schema.SingleNestedAttribute{
													Description:         "ParamKind is a tuple of Group Kind and Version.",
													MarkdownDescription: "ParamKind is a tuple of Group Kind and Version.",
													Attributes: map[string]schema.Attribute{
														"api_version": schema.StringAttribute{
															Description:         "APIVersion is the API group version the resources belong to.In format of 'group/version'.Required.",
															MarkdownDescription: "APIVersion is the API group version the resources belong to.In format of 'group/version'.Required.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"kind": schema.StringAttribute{
															Description:         "Kind is the API kind the resources belong to.Required.",
															MarkdownDescription: "Kind is the API kind the resources belong to.Required.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"param_ref": schema.SingleNestedAttribute{
													Description:         "ParamRef references a parameter resource.",
													MarkdownDescription: "ParamRef references a parameter resource.",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "'name' is the name of the resource being referenced.'name' and 'selector' are mutually exclusive properties. If one is set,the other must be unset.",
															MarkdownDescription: "'name' is the name of the resource being referenced.'name' and 'selector' are mutually exclusive properties. If one is set,the other must be unset.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespace": schema.StringAttribute{
															Description:         "namespace is the namespace of the referenced resource. Allows limitingthe search for params to a specific namespace. Applies to both 'name' and'selector' fields.A per-namespace parameter may be used by specifying a namespace-scoped'paramKind' in the policy and leaving this field empty.- If 'paramKind' is cluster-scoped, this field MUST be unset. Setting thisfield results in a configuration error.- If 'paramKind' is namespace-scoped, the namespace of the object beingevaluated for admission will be used when this field is left unset. Takecare that if this is left empty the binding must not match any cluster-scopedresources, which will result in an error.",
															MarkdownDescription: "namespace is the namespace of the referenced resource. Allows limitingthe search for params to a specific namespace. Applies to both 'name' and'selector' fields.A per-namespace parameter may be used by specifying a namespace-scoped'paramKind' in the policy and leaving this field empty.- If 'paramKind' is cluster-scoped, this field MUST be unset. Setting thisfield results in a configuration error.- If 'paramKind' is namespace-scoped, the namespace of the object beingevaluated for admission will be used when this field is left unset. Takecare that if this is left empty the binding must not match any cluster-scopedresources, which will result in an error.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"parameter_not_found_action": schema.StringAttribute{
															Description:         "'parameterNotFoundAction' controls the behavior of the binding when the resourceexists, and name or selector is valid, but there are no parametersmatched by the binding. If the value is set to 'Allow', then nomatched parameters will be treated as successful validation by the binding.If set to 'Deny', then no matched parameters will be subject to the'failurePolicy' of the policy.Allowed values are 'Allow' or 'Deny'Default to 'Deny'",
															MarkdownDescription: "'parameterNotFoundAction' controls the behavior of the binding when the resourceexists, and name or selector is valid, but there are no parametersmatched by the binding. If the value is set to 'Allow', then nomatched parameters will be treated as successful validation by the binding.If set to 'Deny', then no matched parameters will be subject to the'failurePolicy' of the policy.Allowed values are 'Allow' or 'Deny'Default to 'Deny'",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"selector": schema.SingleNestedAttribute{
															Description:         "selector can be used to match multiple param objects based on their labels.Supply selector: {} to match all resources of the ParamKind.If multiple params are found, they are all evaluated with the policy expressionsand the results are ANDed together.One of 'name' or 'selector' must be set, but 'name' and 'selector' aremutually exclusive properties. If one is set, the other must be unset.",
															MarkdownDescription: "selector can be used to match multiple param objects based on their labels.Supply selector: {} to match all resources of the ParamKind.If multiple params are found, they are all evaluated with the policy expressionsand the results are ANDed together.One of 'name' or 'selector' must be set, but 'name' and 'selector' aremutually exclusive properties. If one is set, the other must be unset.",
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
															Required: false,
															Optional: true,
															Computed: false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"variables": schema.ListNestedAttribute{
													Description:         "Variables contain definitions of variables that can be used in composition of other expressions.Each variable is defined as a named CEL expression.The variables defined here will be available under 'variables' in other expressions of the policy.",
													MarkdownDescription: "Variables contain definitions of variables that can be used in composition of other expressions.Each variable is defined as a named CEL expression.The variables defined here will be available under 'variables' in other expressions of the policy.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"expression": schema.StringAttribute{
																Description:         "Expression is the expression that will be evaluated as the value of the variable.The CEL expression has access to the same identifiers as the CEL expressions in Validation.",
																MarkdownDescription: "Expression is the expression that will be evaluated as the value of the variable.The CEL expression has access to the same identifiers as the CEL expressions in Validation.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "Name is the name of the variable. The name must be a valid CEL identifier and unique among all variables.The variable can be accessed in other expressions through 'variables'For example, if name is 'foo', the variable will be available as 'variables.foo'",
																MarkdownDescription: "Name is the name of the variable. The name must be a valid CEL identifier and unique among all variables.The variable can be accessed in other expressions through 'variables'For example, if name is 'foo', the variable will be available as 'variables.foo'",
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"deny": schema.SingleNestedAttribute{
											Description:         "Deny defines conditions used to pass or fail a validation rule.",
											MarkdownDescription: "Deny defines conditions used to pass or fail a validation rule.",
											Attributes: map[string]schema.Attribute{
												"conditions": schema.MapAttribute{
													Description:         "Multiple conditions can be declared under an 'any' or 'all' statement. A direct listof conditions (without 'any' or 'all' statements) is also supported for backwards compatibilitybut will be deprecated in the next major release.See: https://kyverno.io/docs/writing-policies/validate/#deny-rules",
													MarkdownDescription: "Multiple conditions can be declared under an 'any' or 'all' statement. A direct listof conditions (without 'any' or 'all' statements) is also supported for backwards compatibilitybut will be deprecated in the next major release.See: https://kyverno.io/docs/writing-policies/validate/#deny-rules",
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

										"foreach": schema.ListNestedAttribute{
											Description:         "ForEach applies validate rules to a list of sub-elements by creating a context for each entry in the list and looping over it to apply the specified logic.",
											MarkdownDescription: "ForEach applies validate rules to a list of sub-elements by creating a context for each entry in the list and looping over it to apply the specified logic.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"any_pattern": schema.MapAttribute{
														Description:         "AnyPattern specifies list of validation patterns. At least one of the patternsmust be satisfied for the validation rule to succeed.",
														MarkdownDescription: "AnyPattern specifies list of validation patterns. At least one of the patternsmust be satisfied for the validation rule to succeed.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"context": schema.ListNestedAttribute{
														Description:         "Context defines variables and data sources that can be used during rule execution.",
														MarkdownDescription: "Context defines variables and data sources that can be used during rule execution.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"api_call": schema.SingleNestedAttribute{
																	Description:         "APICall is an HTTP request to the Kubernetes API server, or other JSON web service.The data returned is stored in the context with the name for the context entry.",
																	MarkdownDescription: "APICall is an HTTP request to the Kubernetes API server, or other JSON web service.The data returned is stored in the context with the name for the context entry.",
																	Attributes: map[string]schema.Attribute{
																		"data": schema.ListNestedAttribute{
																			Description:         "The data object specifies the POST data sent to the server.Only applicable when the method field is set to POST.",
																			MarkdownDescription: "The data object specifies the POST data sent to the server.Only applicable when the method field is set to POST.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "Key is a unique identifier for the data value",
																						MarkdownDescription: "Key is a unique identifier for the data value",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"value": schema.MapAttribute{
																						Description:         "Value is the data value",
																						MarkdownDescription: "Value is the data value",
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

																		"jmes_path": schema.StringAttribute{
																			Description:         "JMESPath is an optional JSON Match Expression that can be used totransform the JSON response returned from the server. For examplea JMESPath of 'items | length(@)' applied to the API server responsefor the URLPath '/apis/apps/v1/deployments' will return the total countof deployments across all namespaces.",
																			MarkdownDescription: "JMESPath is an optional JSON Match Expression that can be used totransform the JSON response returned from the server. For examplea JMESPath of 'items | length(@)' applied to the API server responsefor the URLPath '/apis/apps/v1/deployments' will return the total countof deployments across all namespaces.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"method": schema.StringAttribute{
																			Description:         "Method is the HTTP request type (GET or POST). Defaults to GET.",
																			MarkdownDescription: "Method is the HTTP request type (GET or POST). Defaults to GET.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.OneOf("GET", "POST"),
																			},
																		},

																		"service": schema.SingleNestedAttribute{
																			Description:         "Service is an API call to a JSON web service.This is used for non-Kubernetes API server calls.It's mutually exclusive with the URLPath field.",
																			MarkdownDescription: "Service is an API call to a JSON web service.This is used for non-Kubernetes API server calls.It's mutually exclusive with the URLPath field.",
																			Attributes: map[string]schema.Attribute{
																				"ca_bundle": schema.StringAttribute{
																					Description:         "CABundle is a PEM encoded CA bundle which will be used to validatethe server certificate.",
																					MarkdownDescription: "CABundle is a PEM encoded CA bundle which will be used to validatethe server certificate.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"url": schema.StringAttribute{
																					Description:         "URL is the JSON web service URL. A typical form is'https://{service}.{namespace}:{port}/{path}'.",
																					MarkdownDescription: "URL is the JSON web service URL. A typical form is'https://{service}.{namespace}:{port}/{path}'.",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"url_path": schema.StringAttribute{
																			Description:         "URLPath is the URL path to be used in the HTTP GET or POST request to theKubernetes API server (e.g. '/api/v1/namespaces' or  '/apis/apps/v1/deployments').The format required is the same format used by the 'kubectl get --raw' command.See https://kyverno.io/docs/writing-policies/external-data-sources/#variables-from-kubernetes-api-server-callsfor details.It's mutually exclusive with the Service field.",
																			MarkdownDescription: "URLPath is the URL path to be used in the HTTP GET or POST request to theKubernetes API server (e.g. '/api/v1/namespaces' or  '/apis/apps/v1/deployments').The format required is the same format used by the 'kubectl get --raw' command.See https://kyverno.io/docs/writing-policies/external-data-sources/#variables-from-kubernetes-api-server-callsfor details.It's mutually exclusive with the Service field.",
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
																	Description:         "ConfigMap is the ConfigMap reference.",
																	MarkdownDescription: "ConfigMap is the ConfigMap reference.",
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Description:         "Name is the ConfigMap name.",
																			MarkdownDescription: "Name is the ConfigMap name.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"namespace": schema.StringAttribute{
																			Description:         "Namespace is the ConfigMap namespace.",
																			MarkdownDescription: "Namespace is the ConfigMap namespace.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"global_reference": schema.SingleNestedAttribute{
																	Description:         "GlobalContextEntryReference is a reference to a cached global context entry.",
																	MarkdownDescription: "GlobalContextEntryReference is a reference to a cached global context entry.",
																	Attributes: map[string]schema.Attribute{
																		"jmes_path": schema.StringAttribute{
																			Description:         "JMESPath is an optional JSON Match Expression that can be used totransform the JSON response returned from the server. For examplea JMESPath of 'items | length(@)' applied to the API server responsefor the URLPath '/apis/apps/v1/deployments' will return the total countof deployments across all namespaces.",
																			MarkdownDescription: "JMESPath is an optional JSON Match Expression that can be used totransform the JSON response returned from the server. For examplea JMESPath of 'items | length(@)' applied to the API server responsefor the URLPath '/apis/apps/v1/deployments' will return the total countof deployments across all namespaces.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"name": schema.StringAttribute{
																			Description:         "Name of the global context entry",
																			MarkdownDescription: "Name of the global context entry",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"image_registry": schema.SingleNestedAttribute{
																	Description:         "ImageRegistry defines requests to an OCI/Docker V2 registry to fetch imagedetails.",
																	MarkdownDescription: "ImageRegistry defines requests to an OCI/Docker V2 registry to fetch imagedetails.",
																	Attributes: map[string]schema.Attribute{
																		"image_registry_credentials": schema.SingleNestedAttribute{
																			Description:         "ImageRegistryCredentials provides credentials that will be used for authentication with registry",
																			MarkdownDescription: "ImageRegistryCredentials provides credentials that will be used for authentication with registry",
																			Attributes: map[string]schema.Attribute{
																				"allow_insecure_registry": schema.BoolAttribute{
																					Description:         "AllowInsecureRegistry allows insecure access to a registry.",
																					MarkdownDescription: "AllowInsecureRegistry allows insecure access to a registry.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"providers": schema.ListAttribute{
																					Description:         "Providers specifies a list of OCI Registry names, whose authentication providers are provided.It can be of one of these values: default,google,azure,amazon,github.",
																					MarkdownDescription: "Providers specifies a list of OCI Registry names, whose authentication providers are provided.It can be of one of these values: default,google,azure,amazon,github.",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"secrets": schema.ListAttribute{
																					Description:         "Secrets specifies a list of secrets that are provided for credentials.Secrets must live in the Kyverno namespace.",
																					MarkdownDescription: "Secrets specifies a list of secrets that are provided for credentials.Secrets must live in the Kyverno namespace.",
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

																		"jmes_path": schema.StringAttribute{
																			Description:         "JMESPath is an optional JSON Match Expression that can be used totransform the ImageData struct returned as a result of processingthe image reference.",
																			MarkdownDescription: "JMESPath is an optional JSON Match Expression that can be used totransform the ImageData struct returned as a result of processingthe image reference.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"reference": schema.StringAttribute{
																			Description:         "Reference is image reference to a container image in the registry.Example: ghcr.io/kyverno/kyverno:latest",
																			MarkdownDescription: "Reference is image reference to a container image in the registry.Example: ghcr.io/kyverno/kyverno:latest",
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
																	Description:         "Name is the variable name.",
																	MarkdownDescription: "Name is the variable name.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"variable": schema.SingleNestedAttribute{
																	Description:         "Variable defines an arbitrary JMESPath context variable that can be defined inline.",
																	MarkdownDescription: "Variable defines an arbitrary JMESPath context variable that can be defined inline.",
																	Attributes: map[string]schema.Attribute{
																		"default": schema.MapAttribute{
																			Description:         "Default is an optional arbitrary JSON object that the variable may take if the JMESPathexpression evaluates to nil",
																			MarkdownDescription: "Default is an optional arbitrary JSON object that the variable may take if the JMESPathexpression evaluates to nil",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"jmes_path": schema.StringAttribute{
																			Description:         "JMESPath is an optional JMESPath Expression that can be used totransform the variable.",
																			MarkdownDescription: "JMESPath is an optional JMESPath Expression that can be used totransform the variable.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"value": schema.MapAttribute{
																			Description:         "Value is any arbitrary JSON object representable in YAML or JSON form.",
																			MarkdownDescription: "Value is any arbitrary JSON object representable in YAML or JSON form.",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"deny": schema.SingleNestedAttribute{
														Description:         "Deny defines conditions used to pass or fail a validation rule.",
														MarkdownDescription: "Deny defines conditions used to pass or fail a validation rule.",
														Attributes: map[string]schema.Attribute{
															"conditions": schema.MapAttribute{
																Description:         "Multiple conditions can be declared under an 'any' or 'all' statement. A direct listof conditions (without 'any' or 'all' statements) is also supported for backwards compatibilitybut will be deprecated in the next major release.See: https://kyverno.io/docs/writing-policies/validate/#deny-rules",
																MarkdownDescription: "Multiple conditions can be declared under an 'any' or 'all' statement. A direct listof conditions (without 'any' or 'all' statements) is also supported for backwards compatibilitybut will be deprecated in the next major release.See: https://kyverno.io/docs/writing-policies/validate/#deny-rules",
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

													"element_scope": schema.BoolAttribute{
														Description:         "ElementScope specifies whether to use the current list element as the scope for validation. Defaults to 'true' if not specified.When set to 'false', 'request.object' is used as the validation scope within the foreachblock to allow referencing other elements in the subtree.",
														MarkdownDescription: "ElementScope specifies whether to use the current list element as the scope for validation. Defaults to 'true' if not specified.When set to 'false', 'request.object' is used as the validation scope within the foreachblock to allow referencing other elements in the subtree.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"foreach": schema.MapAttribute{
														Description:         "Foreach declares a nested foreach iterator",
														MarkdownDescription: "Foreach declares a nested foreach iterator",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"list": schema.StringAttribute{
														Description:         "List specifies a JMESPath expression that results in one or more elementsto which the validation logic is applied.",
														MarkdownDescription: "List specifies a JMESPath expression that results in one or more elementsto which the validation logic is applied.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"pattern": schema.MapAttribute{
														Description:         "Pattern specifies an overlay-style pattern used to check resources.",
														MarkdownDescription: "Pattern specifies an overlay-style pattern used to check resources.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"preconditions": schema.SingleNestedAttribute{
														Description:         "AnyAllConditions are used to determine if a policy rule should be applied by evaluating aset of conditions. The declaration can contain nested 'any' or 'all' statements.See: https://kyverno.io/docs/writing-policies/preconditions/",
														MarkdownDescription: "AnyAllConditions are used to determine if a policy rule should be applied by evaluating aset of conditions. The declaration can contain nested 'any' or 'all' statements.See: https://kyverno.io/docs/writing-policies/preconditions/",
														Attributes: map[string]schema.Attribute{
															"all": schema.ListNestedAttribute{
																Description:         "AllConditions enable variable-based conditional rule execution. This is useful forfiner control of when an rule is applied. A condition can reference object datausing JMESPath notation.Here, all of the conditions need to pass",
																MarkdownDescription: "AllConditions enable variable-based conditional rule execution. This is useful forfiner control of when an rule is applied. A condition can reference object datausing JMESPath notation.Here, all of the conditions need to pass",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"key": schema.MapAttribute{
																			Description:         "Key is the context entry (using JMESPath) for conditional rule evaluation.",
																			MarkdownDescription: "Key is the context entry (using JMESPath) for conditional rule evaluation.",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"message": schema.StringAttribute{
																			Description:         "Message is an optional display message",
																			MarkdownDescription: "Message is an optional display message",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"operator": schema.StringAttribute{
																			Description:         "Operator is the conditional operation to perform. Valid operators are:Equals, NotEquals, In, AnyIn, AllIn, NotIn, AnyNotIn, AllNotIn, GreaterThanOrEquals,GreaterThan, LessThanOrEquals, LessThan, DurationGreaterThanOrEquals, DurationGreaterThan,DurationLessThanOrEquals, DurationLessThan",
																			MarkdownDescription: "Operator is the conditional operation to perform. Valid operators are:Equals, NotEquals, In, AnyIn, AllIn, NotIn, AnyNotIn, AllNotIn, GreaterThanOrEquals,GreaterThan, LessThanOrEquals, LessThan, DurationGreaterThanOrEquals, DurationGreaterThan,DurationLessThanOrEquals, DurationLessThan",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.OneOf("Equals", "NotEquals", "In", "AnyIn", "AllIn", "NotIn", "AnyNotIn", "AllNotIn", "GreaterThanOrEquals", "GreaterThan", "LessThanOrEquals", "LessThan", "DurationGreaterThanOrEquals", "DurationGreaterThan", "DurationLessThanOrEquals", "DurationLessThan"),
																			},
																		},

																		"value": schema.MapAttribute{
																			Description:         "Value is the conditional value, or set of values. The values can be fixed setor can be variables declared using JMESPath.",
																			MarkdownDescription: "Value is the conditional value, or set of values. The values can be fixed setor can be variables declared using JMESPath.",
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

															"any": schema.ListNestedAttribute{
																Description:         "AnyConditions enable variable-based conditional rule execution. This is useful forfiner control of when an rule is applied. A condition can reference object datausing JMESPath notation.Here, at least one of the conditions need to pass",
																MarkdownDescription: "AnyConditions enable variable-based conditional rule execution. This is useful forfiner control of when an rule is applied. A condition can reference object datausing JMESPath notation.Here, at least one of the conditions need to pass",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"key": schema.MapAttribute{
																			Description:         "Key is the context entry (using JMESPath) for conditional rule evaluation.",
																			MarkdownDescription: "Key is the context entry (using JMESPath) for conditional rule evaluation.",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"message": schema.StringAttribute{
																			Description:         "Message is an optional display message",
																			MarkdownDescription: "Message is an optional display message",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"operator": schema.StringAttribute{
																			Description:         "Operator is the conditional operation to perform. Valid operators are:Equals, NotEquals, In, AnyIn, AllIn, NotIn, AnyNotIn, AllNotIn, GreaterThanOrEquals,GreaterThan, LessThanOrEquals, LessThan, DurationGreaterThanOrEquals, DurationGreaterThan,DurationLessThanOrEquals, DurationLessThan",
																			MarkdownDescription: "Operator is the conditional operation to perform. Valid operators are:Equals, NotEquals, In, AnyIn, AllIn, NotIn, AnyNotIn, AllNotIn, GreaterThanOrEquals,GreaterThan, LessThanOrEquals, LessThan, DurationGreaterThanOrEquals, DurationGreaterThan,DurationLessThanOrEquals, DurationLessThan",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.OneOf("Equals", "NotEquals", "In", "AnyIn", "AllIn", "NotIn", "AnyNotIn", "AllNotIn", "GreaterThanOrEquals", "GreaterThan", "LessThanOrEquals", "LessThan", "DurationGreaterThanOrEquals", "DurationGreaterThan", "DurationLessThanOrEquals", "DurationLessThan"),
																			},
																		},

																		"value": schema.MapAttribute{
																			Description:         "Value is the conditional value, or set of values. The values can be fixed setor can be variables declared using JMESPath.",
																			MarkdownDescription: "Value is the conditional value, or set of values. The values can be fixed setor can be variables declared using JMESPath.",
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

										"manifests": schema.SingleNestedAttribute{
											Description:         "Manifest specifies conditions for manifest verification",
											MarkdownDescription: "Manifest specifies conditions for manifest verification",
											Attributes: map[string]schema.Attribute{
												"annotation_domain": schema.StringAttribute{
													Description:         "AnnotationDomain is custom domain of annotation for message and signature. Default is 'cosign.sigstore.dev'.",
													MarkdownDescription: "AnnotationDomain is custom domain of annotation for message and signature. Default is 'cosign.sigstore.dev'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"attestors": schema.ListNestedAttribute{
													Description:         "Attestors specified the required attestors (i.e. authorities)",
													MarkdownDescription: "Attestors specified the required attestors (i.e. authorities)",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"count": schema.Int64Attribute{
																Description:         "Count specifies the required number of entries that must match. If the count is null, all entries must match(a logical AND). If the count is 1, at least one entry must match (a logical OR). If the count contains avalue N, then N must be less than or equal to the size of entries, and at least N entries must match.",
																MarkdownDescription: "Count specifies the required number of entries that must match. If the count is null, all entries must match(a logical AND). If the count is 1, at least one entry must match (a logical OR). If the count contains avalue N, then N must be less than or equal to the size of entries, and at least N entries must match.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(1),
																},
															},

															"entries": schema.ListNestedAttribute{
																Description:         "Entries contains the available attestors. An attestor can be a static key,attributes for keyless verification, or a nested attestor declaration.",
																MarkdownDescription: "Entries contains the available attestors. An attestor can be a static key,attributes for keyless verification, or a nested attestor declaration.",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"annotations": schema.MapAttribute{
																			Description:         "Annotations are used for image verification.Every specified key-value pair must exist and match in the verified payload.The payload may contain other key-value pairs.",
																			MarkdownDescription: "Annotations are used for image verification.Every specified key-value pair must exist and match in the verified payload.The payload may contain other key-value pairs.",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"attestor": schema.MapAttribute{
																			Description:         "Attestor is a nested set of Attestor used to specify a more complex set of match authorities.",
																			MarkdownDescription: "Attestor is a nested set of Attestor used to specify a more complex set of match authorities.",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"certificates": schema.SingleNestedAttribute{
																			Description:         "Certificates specifies one or more certificates.",
																			MarkdownDescription: "Certificates specifies one or more certificates.",
																			Attributes: map[string]schema.Attribute{
																				"cert": schema.StringAttribute{
																					Description:         "Cert is an optional PEM-encoded public certificate.",
																					MarkdownDescription: "Cert is an optional PEM-encoded public certificate.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"cert_chain": schema.StringAttribute{
																					Description:         "CertChain is an optional PEM encoded set of certificates used to verify.",
																					MarkdownDescription: "CertChain is an optional PEM encoded set of certificates used to verify.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"ctlog": schema.SingleNestedAttribute{
																					Description:         "CTLog (certificate timestamp log) provides a configuration for validation of Signed CertificateTimestamps (SCTs). If the value is unset, the default behavior by Cosign is used.",
																					MarkdownDescription: "CTLog (certificate timestamp log) provides a configuration for validation of Signed CertificateTimestamps (SCTs). If the value is unset, the default behavior by Cosign is used.",
																					Attributes: map[string]schema.Attribute{
																						"ignore_sct": schema.BoolAttribute{
																							Description:         "IgnoreSCT defines whether to use the Signed Certificate Timestamp (SCT) log to check for a certificatetimestamp. Default is false. Set to true if this was opted out during signing.",
																							MarkdownDescription: "IgnoreSCT defines whether to use the Signed Certificate Timestamp (SCT) log to check for a certificatetimestamp. Default is false. Set to true if this was opted out during signing.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"pubkey": schema.StringAttribute{
																							Description:         "PubKey, if set, is used to validate SCTs against a custom source.",
																							MarkdownDescription: "PubKey, if set, is used to validate SCTs against a custom source.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"tsa_cert_chain": schema.StringAttribute{
																							Description:         "TSACertChain, if set, is the PEM-encoded certificate chain file for the RFC3161 timestamp authority. Mustcontain the root CA certificate. Optionally may contain intermediate CA certificates, andmay contain the leaf TSA certificate if not present in the timestamurce.",
																							MarkdownDescription: "TSACertChain, if set, is the PEM-encoded certificate chain file for the RFC3161 timestamp authority. Mustcontain the root CA certificate. Optionally may contain intermediate CA certificates, andmay contain the leaf TSA certificate if not present in the timestamurce.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},
																					},
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"rekor": schema.SingleNestedAttribute{
																					Description:         "Rekor provides configuration for the Rekor transparency log service. If an empty objectis provided the public instance of Rekor (https://rekor.sigstore.dev) is used.",
																					MarkdownDescription: "Rekor provides configuration for the Rekor transparency log service. If an empty objectis provided the public instance of Rekor (https://rekor.sigstore.dev) is used.",
																					Attributes: map[string]schema.Attribute{
																						"ignore_tlog": schema.BoolAttribute{
																							Description:         "IgnoreTlog skips transparency log verification.",
																							MarkdownDescription: "IgnoreTlog skips transparency log verification.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"pubkey": schema.StringAttribute{
																							Description:         "RekorPubKey is an optional PEM-encoded public key to use for a custom Rekor.If set, this will be used to validate transparency log signatures from a custom Rekor.",
																							MarkdownDescription: "RekorPubKey is an optional PEM-encoded public key to use for a custom Rekor.If set, this will be used to validate transparency log signatures from a custom Rekor.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"url": schema.StringAttribute{
																							Description:         "URL is the address of the transparency log. Defaults to the public Rekor log instance https://rekor.sigstore.dev.",
																							MarkdownDescription: "URL is the address of the transparency log. Defaults to the public Rekor log instance https://rekor.sigstore.dev.",
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

																		"keyless": schema.SingleNestedAttribute{
																			Description:         "Keyless is a set of attribute used to verify a Sigstore keyless attestor.See https://github.com/sigstore/cosign/blob/main/KEYLESS.md.",
																			MarkdownDescription: "Keyless is a set of attribute used to verify a Sigstore keyless attestor.See https://github.com/sigstore/cosign/blob/main/KEYLESS.md.",
																			Attributes: map[string]schema.Attribute{
																				"additional_extensions": schema.MapAttribute{
																					Description:         "AdditionalExtensions are certificate-extensions used for keyless signing.",
																					MarkdownDescription: "AdditionalExtensions are certificate-extensions used for keyless signing.",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"ctlog": schema.SingleNestedAttribute{
																					Description:         "CTLog (certificate timestamp log) provides a configuration for validation of Signed CertificateTimestamps (SCTs). If the value is unset, the default behavior by Cosign is used.",
																					MarkdownDescription: "CTLog (certificate timestamp log) provides a configuration for validation of Signed CertificateTimestamps (SCTs). If the value is unset, the default behavior by Cosign is used.",
																					Attributes: map[string]schema.Attribute{
																						"ignore_sct": schema.BoolAttribute{
																							Description:         "IgnoreSCT defines whether to use the Signed Certificate Timestamp (SCT) log to check for a certificatetimestamp. Default is false. Set to true if this was opted out during signing.",
																							MarkdownDescription: "IgnoreSCT defines whether to use the Signed Certificate Timestamp (SCT) log to check for a certificatetimestamp. Default is false. Set to true if this was opted out during signing.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"pubkey": schema.StringAttribute{
																							Description:         "PubKey, if set, is used to validate SCTs against a custom source.",
																							MarkdownDescription: "PubKey, if set, is used to validate SCTs against a custom source.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"tsa_cert_chain": schema.StringAttribute{
																							Description:         "TSACertChain, if set, is the PEM-encoded certificate chain file for the RFC3161 timestamp authority. Mustcontain the root CA certificate. Optionally may contain intermediate CA certificates, andmay contain the leaf TSA certificate if not present in the timestamurce.",
																							MarkdownDescription: "TSACertChain, if set, is the PEM-encoded certificate chain file for the RFC3161 timestamp authority. Mustcontain the root CA certificate. Optionally may contain intermediate CA certificates, andmay contain the leaf TSA certificate if not present in the timestamurce.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},
																					},
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"issuer": schema.StringAttribute{
																					Description:         "Issuer is the certificate issuer used for keyless signing.",
																					MarkdownDescription: "Issuer is the certificate issuer used for keyless signing.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"rekor": schema.SingleNestedAttribute{
																					Description:         "Rekor provides configuration for the Rekor transparency log service. If an empty objectis provided the public instance of Rekor (https://rekor.sigstore.dev) is used.",
																					MarkdownDescription: "Rekor provides configuration for the Rekor transparency log service. If an empty objectis provided the public instance of Rekor (https://rekor.sigstore.dev) is used.",
																					Attributes: map[string]schema.Attribute{
																						"ignore_tlog": schema.BoolAttribute{
																							Description:         "IgnoreTlog skips transparency log verification.",
																							MarkdownDescription: "IgnoreTlog skips transparency log verification.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"pubkey": schema.StringAttribute{
																							Description:         "RekorPubKey is an optional PEM-encoded public key to use for a custom Rekor.If set, this will be used to validate transparency log signatures from a custom Rekor.",
																							MarkdownDescription: "RekorPubKey is an optional PEM-encoded public key to use for a custom Rekor.If set, this will be used to validate transparency log signatures from a custom Rekor.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"url": schema.StringAttribute{
																							Description:         "URL is the address of the transparency log. Defaults to the public Rekor log instance https://rekor.sigstore.dev.",
																							MarkdownDescription: "URL is the address of the transparency log. Defaults to the public Rekor log instance https://rekor.sigstore.dev.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},
																					},
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"roots": schema.StringAttribute{
																					Description:         "Roots is an optional set of PEM encoded trusted root certificates.If not provided, the system roots are used.",
																					MarkdownDescription: "Roots is an optional set of PEM encoded trusted root certificates.If not provided, the system roots are used.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"subject": schema.StringAttribute{
																					Description:         "Subject is the verified identity used for keyless signing, for example the email address.",
																					MarkdownDescription: "Subject is the verified identity used for keyless signing, for example the email address.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"keys": schema.SingleNestedAttribute{
																			Description:         "Keys specifies one or more public keys.",
																			MarkdownDescription: "Keys specifies one or more public keys.",
																			Attributes: map[string]schema.Attribute{
																				"ctlog": schema.SingleNestedAttribute{
																					Description:         "CTLog (certificate timestamp log) provides a configuration for validation of Signed CertificateTimestamps (SCTs). If the value is unset, the default behavior by Cosign is used.",
																					MarkdownDescription: "CTLog (certificate timestamp log) provides a configuration for validation of Signed CertificateTimestamps (SCTs). If the value is unset, the default behavior by Cosign is used.",
																					Attributes: map[string]schema.Attribute{
																						"ignore_sct": schema.BoolAttribute{
																							Description:         "IgnoreSCT defines whether to use the Signed Certificate Timestamp (SCT) log to check for a certificatetimestamp. Default is false. Set to true if this was opted out during signing.",
																							MarkdownDescription: "IgnoreSCT defines whether to use the Signed Certificate Timestamp (SCT) log to check for a certificatetimestamp. Default is false. Set to true if this was opted out during signing.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"pubkey": schema.StringAttribute{
																							Description:         "PubKey, if set, is used to validate SCTs against a custom source.",
																							MarkdownDescription: "PubKey, if set, is used to validate SCTs against a custom source.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"tsa_cert_chain": schema.StringAttribute{
																							Description:         "TSACertChain, if set, is the PEM-encoded certificate chain file for the RFC3161 timestamp authority. Mustcontain the root CA certificate. Optionally may contain intermediate CA certificates, andmay contain the leaf TSA certificate if not present in the timestamurce.",
																							MarkdownDescription: "TSACertChain, if set, is the PEM-encoded certificate chain file for the RFC3161 timestamp authority. Mustcontain the root CA certificate. Optionally may contain intermediate CA certificates, andmay contain the leaf TSA certificate if not present in the timestamurce.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},
																					},
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"kms": schema.StringAttribute{
																					Description:         "KMS provides the URI to the public key stored in a Key Management System. See:https://github.com/sigstore/cosign/blob/main/KMS.md",
																					MarkdownDescription: "KMS provides the URI to the public key stored in a Key Management System. See:https://github.com/sigstore/cosign/blob/main/KMS.md",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"public_keys": schema.StringAttribute{
																					Description:         "Keys is a set of X.509 public keys used to verify image signatures. The keys can be directlyspecified or can be a variable reference to a key specified in a ConfigMap (seehttps://kyverno.io/docs/writing-policies/variables/), or reference a standard Kubernetes Secretelsewhere in the cluster by specifying it in the format 'k8s://<namespace>/<secret_name>'.The named Secret must specify a key 'cosign.pub' containing the public key used forverification, (see https://github.com/sigstore/cosign/blob/main/KMS.md#kubernetes-secret).When multiple keys are specified each key is processed as a separate staticKey entry(.attestors[*].entries.keys) within the set of attestors and the count is applied across the keys.",
																					MarkdownDescription: "Keys is a set of X.509 public keys used to verify image signatures. The keys can be directlyspecified or can be a variable reference to a key specified in a ConfigMap (seehttps://kyverno.io/docs/writing-policies/variables/), or reference a standard Kubernetes Secretelsewhere in the cluster by specifying it in the format 'k8s://<namespace>/<secret_name>'.The named Secret must specify a key 'cosign.pub' containing the public key used forverification, (see https://github.com/sigstore/cosign/blob/main/KMS.md#kubernetes-secret).When multiple keys are specified each key is processed as a separate staticKey entry(.attestors[*].entries.keys) within the set of attestors and the count is applied across the keys.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"rekor": schema.SingleNestedAttribute{
																					Description:         "Rekor provides configuration for the Rekor transparency log service. If an empty objectis provided the public instance of Rekor (https://rekor.sigstore.dev) is used.",
																					MarkdownDescription: "Rekor provides configuration for the Rekor transparency log service. If an empty objectis provided the public instance of Rekor (https://rekor.sigstore.dev) is used.",
																					Attributes: map[string]schema.Attribute{
																						"ignore_tlog": schema.BoolAttribute{
																							Description:         "IgnoreTlog skips transparency log verification.",
																							MarkdownDescription: "IgnoreTlog skips transparency log verification.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"pubkey": schema.StringAttribute{
																							Description:         "RekorPubKey is an optional PEM-encoded public key to use for a custom Rekor.If set, this will be used to validate transparency log signatures from a custom Rekor.",
																							MarkdownDescription: "RekorPubKey is an optional PEM-encoded public key to use for a custom Rekor.If set, this will be used to validate transparency log signatures from a custom Rekor.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"url": schema.StringAttribute{
																							Description:         "URL is the address of the transparency log. Defaults to the public Rekor log instance https://rekor.sigstore.dev.",
																							MarkdownDescription: "URL is the address of the transparency log. Defaults to the public Rekor log instance https://rekor.sigstore.dev.",
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
																					Description:         "Reference to a Secret resource that contains a public key",
																					MarkdownDescription: "Reference to a Secret resource that contains a public key",
																					Attributes: map[string]schema.Attribute{
																						"name": schema.StringAttribute{
																							Description:         "Name of the secret. The provided secret must contain a key named cosign.pub.",
																							MarkdownDescription: "Name of the secret. The provided secret must contain a key named cosign.pub.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"namespace": schema.StringAttribute{
																							Description:         "Namespace name where the Secret exists.",
																							MarkdownDescription: "Namespace name where the Secret exists.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},
																					},
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"signature_algorithm": schema.StringAttribute{
																					Description:         "Specify signature algorithm for public keys. Supported values are sha224, sha256, sha384 and sha512.",
																					MarkdownDescription: "Specify signature algorithm for public keys. Supported values are sha224, sha256, sha384 and sha512.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"repository": schema.StringAttribute{
																			Description:         "Repository is an optional alternate OCI repository to use for signatures and attestations that match this rule.If specified Repository will override other OCI image repository locations for this Attestor.",
																			MarkdownDescription: "Repository is an optional alternate OCI repository to use for signatures and attestations that match this rule.If specified Repository will override other OCI image repository locations for this Attestor.",
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
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"dry_run": schema.SingleNestedAttribute{
													Description:         "DryRun configuration",
													MarkdownDescription: "DryRun configuration",
													Attributes: map[string]schema.Attribute{
														"enable": schema.BoolAttribute{
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

												"ignore_fields": schema.ListNestedAttribute{
													Description:         "Fields which will be ignored while comparing manifests.",
													MarkdownDescription: "Fields which will be ignored while comparing manifests.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"fields": schema.ListAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"objects": schema.ListNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"group": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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

																		"name": schema.StringAttribute{
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

																		"version": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"repository": schema.StringAttribute{
													Description:         "Repository is an optional alternate OCI repository to use for resource bundle reference.The repository can be overridden per Attestor or Attestation.",
													MarkdownDescription: "Repository is an optional alternate OCI repository to use for resource bundle reference.The repository can be overridden per Attestor or Attestation.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"message": schema.StringAttribute{
											Description:         "Message specifies a custom message to be displayed on failure.",
											MarkdownDescription: "Message specifies a custom message to be displayed on failure.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"pattern": schema.MapAttribute{
											Description:         "Pattern specifies an overlay-style pattern used to check resources.",
											MarkdownDescription: "Pattern specifies an overlay-style pattern used to check resources.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"pod_security": schema.SingleNestedAttribute{
											Description:         "PodSecurity applies exemptions for Kubernetes Pod Security admissionby specifying exclusions for Pod Security Standards controls.",
											MarkdownDescription: "PodSecurity applies exemptions for Kubernetes Pod Security admissionby specifying exclusions for Pod Security Standards controls.",
											Attributes: map[string]schema.Attribute{
												"exclude": schema.ListNestedAttribute{
													Description:         "Exclude specifies the Pod Security Standard controls to be excluded.",
													MarkdownDescription: "Exclude specifies the Pod Security Standard controls to be excluded.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"control_name": schema.StringAttribute{
																Description:         "ControlName specifies the name of the Pod Security Standard control.See: https://kubernetes.io/docs/concepts/security/pod-security-standards/",
																MarkdownDescription: "ControlName specifies the name of the Pod Security Standard control.See: https://kubernetes.io/docs/concepts/security/pod-security-standards/",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("HostProcess", "Host Namespaces", "Privileged Containers", "Capabilities", "HostPath Volumes", "Host Ports", "AppArmor", "SELinux", "/proc Mount Type", "Seccomp", "Sysctls", "Volume Types", "Privilege Escalation", "Running as Non-root", "Running as Non-root user"),
																},
															},

															"images": schema.ListAttribute{
																Description:         "Images selects matching containers and applies the container level PSS.Each image is the image name consisting of the registry address, repository, image, and tag.Empty list matches no containers, PSS checks are applied at the pod level only.Wildcards ('*' and '?') are allowed. See: https://kubernetes.io/docs/concepts/containers/images.",
																MarkdownDescription: "Images selects matching containers and applies the container level PSS.Each image is the image name consisting of the registry address, repository, image, and tag.Empty list matches no containers, PSS checks are applied at the pod level only.Wildcards ('*' and '?') are allowed. See: https://kubernetes.io/docs/concepts/containers/images.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"restricted_field": schema.StringAttribute{
																Description:         "RestrictedField selects the field for the given Pod Security Standard control.When not set, all restricted fields for the control are selected.",
																MarkdownDescription: "RestrictedField selects the field for the given Pod Security Standard control.When not set, all restricted fields for the control are selected.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"values": schema.ListAttribute{
																Description:         "Values defines the allowed values that can be excluded.",
																MarkdownDescription: "Values defines the allowed values that can be excluded.",
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

												"level": schema.StringAttribute{
													Description:         "Level defines the Pod Security Standard level to be applied to workloads.Allowed values are privileged, baseline, and restricted.",
													MarkdownDescription: "Level defines the Pod Security Standard level to be applied to workloads.Allowed values are privileged, baseline, and restricted.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("privileged", "baseline", "restricted"),
													},
												},

												"version": schema.StringAttribute{
													Description:         "Version defines the Pod Security Standard versions that Kubernetes supports.Allowed values are v1.19, v1.20, v1.21, v1.22, v1.23, v1.24, v1.25, v1.26, v1.27, v1.28, v1.29, latest. Defaults to latest.",
													MarkdownDescription: "Version defines the Pod Security Standard versions that Kubernetes supports.Allowed values are v1.19, v1.20, v1.21, v1.22, v1.23, v1.24, v1.25, v1.26, v1.27, v1.28, v1.29, latest. Defaults to latest.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("v1.19", "v1.20", "v1.21", "v1.22", "v1.23", "v1.24", "v1.25", "v1.26", "v1.27", "v1.28", "v1.29", "latest"),
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"validation_failure_action": schema.StringAttribute{
											Description:         "ValidationFailureAction defines if a validation policy rule violation should blockthe admission review request (enforce), or allow (audit) the admission review requestand report an error in a policy report. Optional.Allowed values are audit or enforce.",
											MarkdownDescription: "ValidationFailureAction defines if a validation policy rule violation should blockthe admission review request (enforce), or allow (audit) the admission review requestand report an error in a policy report. Optional.Allowed values are audit or enforce.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("audit", "enforce", "Audit", "Enforce"),
											},
										},

										"validation_failure_action_overrides": schema.ListNestedAttribute{
											Description:         "ValidationFailureActionOverrides is a Cluster Policy attribute that specifies ValidationFailureActionnamespace-wise. It overrides ValidationFailureAction for the specified namespaces.",
											MarkdownDescription: "ValidationFailureActionOverrides is a Cluster Policy attribute that specifies ValidationFailureActionnamespace-wise. It overrides ValidationFailureAction for the specified namespaces.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"action": schema.StringAttribute{
														Description:         "ValidationFailureAction defines the policy validation failure action",
														MarkdownDescription: "ValidationFailureAction defines the policy validation failure action",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("audit", "enforce", "Audit", "Enforce"),
														},
													},

													"namespace_selector": schema.SingleNestedAttribute{
														Description:         "A label selector is a label query over a set of resources. The result of matchLabels andmatchExpressions are ANDed. An empty label selector matches all objects. A nulllabel selector matches no objects.",
														MarkdownDescription: "A label selector is a label query over a set of resources. The result of matchLabels andmatchExpressions are ANDed. An empty label selector matches all objects. A nulllabel selector matches no objects.",
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
														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespaces": schema.ListAttribute{
														Description:         "",
														MarkdownDescription: "",
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
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"verify_images": schema.ListNestedAttribute{
									Description:         "VerifyImages is used to verify image signatures and mutate them to add a digest",
									MarkdownDescription: "VerifyImages is used to verify image signatures and mutate them to add a digest",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"additional_extensions": schema.MapAttribute{
												Description:         "Deprecated.",
												MarkdownDescription: "Deprecated.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"annotations": schema.MapAttribute{
												Description:         "Deprecated. Use annotations per Attestor instead.",
												MarkdownDescription: "Deprecated. Use annotations per Attestor instead.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"attestations": schema.ListNestedAttribute{
												Description:         "Attestations are optional checks for signed in-toto Statements used to verify the image.See https://github.com/in-toto/attestation. Kyverno fetches signed attestations from theOCI registry and decodes them into a list of Statement declarations.",
												MarkdownDescription: "Attestations are optional checks for signed in-toto Statements used to verify the image.See https://github.com/in-toto/attestation. Kyverno fetches signed attestations from theOCI registry and decodes them into a list of Statement declarations.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"attestors": schema.ListNestedAttribute{
															Description:         "Attestors specify the required attestors (i.e. authorities).",
															MarkdownDescription: "Attestors specify the required attestors (i.e. authorities).",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"count": schema.Int64Attribute{
																		Description:         "Count specifies the required number of entries that must match. If the count is null, all entries must match(a logical AND). If the count is 1, at least one entry must match (a logical OR). If the count contains avalue N, then N must be less than or equal to the size of entries, and at least N entries must match.",
																		MarkdownDescription: "Count specifies the required number of entries that must match. If the count is null, all entries must match(a logical AND). If the count is 1, at least one entry must match (a logical OR). If the count contains avalue N, then N must be less than or equal to the size of entries, and at least N entries must match.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.Int64{
																			int64validator.AtLeast(1),
																		},
																	},

																	"entries": schema.ListNestedAttribute{
																		Description:         "Entries contains the available attestors. An attestor can be a static key,attributes for keyless verification, or a nested attestor declaration.",
																		MarkdownDescription: "Entries contains the available attestors. An attestor can be a static key,attributes for keyless verification, or a nested attestor declaration.",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"annotations": schema.MapAttribute{
																					Description:         "Annotations are used for image verification.Every specified key-value pair must exist and match in the verified payload.The payload may contain other key-value pairs.",
																					MarkdownDescription: "Annotations are used for image verification.Every specified key-value pair must exist and match in the verified payload.The payload may contain other key-value pairs.",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"attestor": schema.MapAttribute{
																					Description:         "Attestor is a nested set of Attestor used to specify a more complex set of match authorities.",
																					MarkdownDescription: "Attestor is a nested set of Attestor used to specify a more complex set of match authorities.",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"certificates": schema.SingleNestedAttribute{
																					Description:         "Certificates specifies one or more certificates.",
																					MarkdownDescription: "Certificates specifies one or more certificates.",
																					Attributes: map[string]schema.Attribute{
																						"cert": schema.StringAttribute{
																							Description:         "Cert is an optional PEM-encoded public certificate.",
																							MarkdownDescription: "Cert is an optional PEM-encoded public certificate.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"cert_chain": schema.StringAttribute{
																							Description:         "CertChain is an optional PEM encoded set of certificates used to verify.",
																							MarkdownDescription: "CertChain is an optional PEM encoded set of certificates used to verify.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"ctlog": schema.SingleNestedAttribute{
																							Description:         "CTLog (certificate timestamp log) provides a configuration for validation of Signed CertificateTimestamps (SCTs). If the value is unset, the default behavior by Cosign is used.",
																							MarkdownDescription: "CTLog (certificate timestamp log) provides a configuration for validation of Signed CertificateTimestamps (SCTs). If the value is unset, the default behavior by Cosign is used.",
																							Attributes: map[string]schema.Attribute{
																								"ignore_sct": schema.BoolAttribute{
																									Description:         "IgnoreSCT defines whether to use the Signed Certificate Timestamp (SCT) log to check for a certificatetimestamp. Default is false. Set to true if this was opted out during signing.",
																									MarkdownDescription: "IgnoreSCT defines whether to use the Signed Certificate Timestamp (SCT) log to check for a certificatetimestamp. Default is false. Set to true if this was opted out during signing.",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"pubkey": schema.StringAttribute{
																									Description:         "PubKey, if set, is used to validate SCTs against a custom source.",
																									MarkdownDescription: "PubKey, if set, is used to validate SCTs against a custom source.",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"tsa_cert_chain": schema.StringAttribute{
																									Description:         "TSACertChain, if set, is the PEM-encoded certificate chain file for the RFC3161 timestamp authority. Mustcontain the root CA certificate. Optionally may contain intermediate CA certificates, andmay contain the leaf TSA certificate if not present in the timestamurce.",
																									MarkdownDescription: "TSACertChain, if set, is the PEM-encoded certificate chain file for the RFC3161 timestamp authority. Mustcontain the root CA certificate. Optionally may contain intermediate CA certificates, andmay contain the leaf TSA certificate if not present in the timestamurce.",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},
																							},
																							Required: false,
																							Optional: true,
																							Computed: false,
																						},

																						"rekor": schema.SingleNestedAttribute{
																							Description:         "Rekor provides configuration for the Rekor transparency log service. If an empty objectis provided the public instance of Rekor (https://rekor.sigstore.dev) is used.",
																							MarkdownDescription: "Rekor provides configuration for the Rekor transparency log service. If an empty objectis provided the public instance of Rekor (https://rekor.sigstore.dev) is used.",
																							Attributes: map[string]schema.Attribute{
																								"ignore_tlog": schema.BoolAttribute{
																									Description:         "IgnoreTlog skips transparency log verification.",
																									MarkdownDescription: "IgnoreTlog skips transparency log verification.",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"pubkey": schema.StringAttribute{
																									Description:         "RekorPubKey is an optional PEM-encoded public key to use for a custom Rekor.If set, this will be used to validate transparency log signatures from a custom Rekor.",
																									MarkdownDescription: "RekorPubKey is an optional PEM-encoded public key to use for a custom Rekor.If set, this will be used to validate transparency log signatures from a custom Rekor.",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"url": schema.StringAttribute{
																									Description:         "URL is the address of the transparency log. Defaults to the public Rekor log instance https://rekor.sigstore.dev.",
																									MarkdownDescription: "URL is the address of the transparency log. Defaults to the public Rekor log instance https://rekor.sigstore.dev.",
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

																				"keyless": schema.SingleNestedAttribute{
																					Description:         "Keyless is a set of attribute used to verify a Sigstore keyless attestor.See https://github.com/sigstore/cosign/blob/main/KEYLESS.md.",
																					MarkdownDescription: "Keyless is a set of attribute used to verify a Sigstore keyless attestor.See https://github.com/sigstore/cosign/blob/main/KEYLESS.md.",
																					Attributes: map[string]schema.Attribute{
																						"additional_extensions": schema.MapAttribute{
																							Description:         "AdditionalExtensions are certificate-extensions used for keyless signing.",
																							MarkdownDescription: "AdditionalExtensions are certificate-extensions used for keyless signing.",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"ctlog": schema.SingleNestedAttribute{
																							Description:         "CTLog (certificate timestamp log) provides a configuration for validation of Signed CertificateTimestamps (SCTs). If the value is unset, the default behavior by Cosign is used.",
																							MarkdownDescription: "CTLog (certificate timestamp log) provides a configuration for validation of Signed CertificateTimestamps (SCTs). If the value is unset, the default behavior by Cosign is used.",
																							Attributes: map[string]schema.Attribute{
																								"ignore_sct": schema.BoolAttribute{
																									Description:         "IgnoreSCT defines whether to use the Signed Certificate Timestamp (SCT) log to check for a certificatetimestamp. Default is false. Set to true if this was opted out during signing.",
																									MarkdownDescription: "IgnoreSCT defines whether to use the Signed Certificate Timestamp (SCT) log to check for a certificatetimestamp. Default is false. Set to true if this was opted out during signing.",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"pubkey": schema.StringAttribute{
																									Description:         "PubKey, if set, is used to validate SCTs against a custom source.",
																									MarkdownDescription: "PubKey, if set, is used to validate SCTs against a custom source.",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"tsa_cert_chain": schema.StringAttribute{
																									Description:         "TSACertChain, if set, is the PEM-encoded certificate chain file for the RFC3161 timestamp authority. Mustcontain the root CA certificate. Optionally may contain intermediate CA certificates, andmay contain the leaf TSA certificate if not present in the timestamurce.",
																									MarkdownDescription: "TSACertChain, if set, is the PEM-encoded certificate chain file for the RFC3161 timestamp authority. Mustcontain the root CA certificate. Optionally may contain intermediate CA certificates, andmay contain the leaf TSA certificate if not present in the timestamurce.",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},
																							},
																							Required: false,
																							Optional: true,
																							Computed: false,
																						},

																						"issuer": schema.StringAttribute{
																							Description:         "Issuer is the certificate issuer used for keyless signing.",
																							MarkdownDescription: "Issuer is the certificate issuer used for keyless signing.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"rekor": schema.SingleNestedAttribute{
																							Description:         "Rekor provides configuration for the Rekor transparency log service. If an empty objectis provided the public instance of Rekor (https://rekor.sigstore.dev) is used.",
																							MarkdownDescription: "Rekor provides configuration for the Rekor transparency log service. If an empty objectis provided the public instance of Rekor (https://rekor.sigstore.dev) is used.",
																							Attributes: map[string]schema.Attribute{
																								"ignore_tlog": schema.BoolAttribute{
																									Description:         "IgnoreTlog skips transparency log verification.",
																									MarkdownDescription: "IgnoreTlog skips transparency log verification.",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"pubkey": schema.StringAttribute{
																									Description:         "RekorPubKey is an optional PEM-encoded public key to use for a custom Rekor.If set, this will be used to validate transparency log signatures from a custom Rekor.",
																									MarkdownDescription: "RekorPubKey is an optional PEM-encoded public key to use for a custom Rekor.If set, this will be used to validate transparency log signatures from a custom Rekor.",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"url": schema.StringAttribute{
																									Description:         "URL is the address of the transparency log. Defaults to the public Rekor log instance https://rekor.sigstore.dev.",
																									MarkdownDescription: "URL is the address of the transparency log. Defaults to the public Rekor log instance https://rekor.sigstore.dev.",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},
																							},
																							Required: false,
																							Optional: true,
																							Computed: false,
																						},

																						"roots": schema.StringAttribute{
																							Description:         "Roots is an optional set of PEM encoded trusted root certificates.If not provided, the system roots are used.",
																							MarkdownDescription: "Roots is an optional set of PEM encoded trusted root certificates.If not provided, the system roots are used.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"subject": schema.StringAttribute{
																							Description:         "Subject is the verified identity used for keyless signing, for example the email address.",
																							MarkdownDescription: "Subject is the verified identity used for keyless signing, for example the email address.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},
																					},
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"keys": schema.SingleNestedAttribute{
																					Description:         "Keys specifies one or more public keys.",
																					MarkdownDescription: "Keys specifies one or more public keys.",
																					Attributes: map[string]schema.Attribute{
																						"ctlog": schema.SingleNestedAttribute{
																							Description:         "CTLog (certificate timestamp log) provides a configuration for validation of Signed CertificateTimestamps (SCTs). If the value is unset, the default behavior by Cosign is used.",
																							MarkdownDescription: "CTLog (certificate timestamp log) provides a configuration for validation of Signed CertificateTimestamps (SCTs). If the value is unset, the default behavior by Cosign is used.",
																							Attributes: map[string]schema.Attribute{
																								"ignore_sct": schema.BoolAttribute{
																									Description:         "IgnoreSCT defines whether to use the Signed Certificate Timestamp (SCT) log to check for a certificatetimestamp. Default is false. Set to true if this was opted out during signing.",
																									MarkdownDescription: "IgnoreSCT defines whether to use the Signed Certificate Timestamp (SCT) log to check for a certificatetimestamp. Default is false. Set to true if this was opted out during signing.",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"pubkey": schema.StringAttribute{
																									Description:         "PubKey, if set, is used to validate SCTs against a custom source.",
																									MarkdownDescription: "PubKey, if set, is used to validate SCTs against a custom source.",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"tsa_cert_chain": schema.StringAttribute{
																									Description:         "TSACertChain, if set, is the PEM-encoded certificate chain file for the RFC3161 timestamp authority. Mustcontain the root CA certificate. Optionally may contain intermediate CA certificates, andmay contain the leaf TSA certificate if not present in the timestamurce.",
																									MarkdownDescription: "TSACertChain, if set, is the PEM-encoded certificate chain file for the RFC3161 timestamp authority. Mustcontain the root CA certificate. Optionally may contain intermediate CA certificates, andmay contain the leaf TSA certificate if not present in the timestamurce.",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},
																							},
																							Required: false,
																							Optional: true,
																							Computed: false,
																						},

																						"kms": schema.StringAttribute{
																							Description:         "KMS provides the URI to the public key stored in a Key Management System. See:https://github.com/sigstore/cosign/blob/main/KMS.md",
																							MarkdownDescription: "KMS provides the URI to the public key stored in a Key Management System. See:https://github.com/sigstore/cosign/blob/main/KMS.md",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"public_keys": schema.StringAttribute{
																							Description:         "Keys is a set of X.509 public keys used to verify image signatures. The keys can be directlyspecified or can be a variable reference to a key specified in a ConfigMap (seehttps://kyverno.io/docs/writing-policies/variables/), or reference a standard Kubernetes Secretelsewhere in the cluster by specifying it in the format 'k8s://<namespace>/<secret_name>'.The named Secret must specify a key 'cosign.pub' containing the public key used forverification, (see https://github.com/sigstore/cosign/blob/main/KMS.md#kubernetes-secret).When multiple keys are specified each key is processed as a separate staticKey entry(.attestors[*].entries.keys) within the set of attestors and the count is applied across the keys.",
																							MarkdownDescription: "Keys is a set of X.509 public keys used to verify image signatures. The keys can be directlyspecified or can be a variable reference to a key specified in a ConfigMap (seehttps://kyverno.io/docs/writing-policies/variables/), or reference a standard Kubernetes Secretelsewhere in the cluster by specifying it in the format 'k8s://<namespace>/<secret_name>'.The named Secret must specify a key 'cosign.pub' containing the public key used forverification, (see https://github.com/sigstore/cosign/blob/main/KMS.md#kubernetes-secret).When multiple keys are specified each key is processed as a separate staticKey entry(.attestors[*].entries.keys) within the set of attestors and the count is applied across the keys.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"rekor": schema.SingleNestedAttribute{
																							Description:         "Rekor provides configuration for the Rekor transparency log service. If an empty objectis provided the public instance of Rekor (https://rekor.sigstore.dev) is used.",
																							MarkdownDescription: "Rekor provides configuration for the Rekor transparency log service. If an empty objectis provided the public instance of Rekor (https://rekor.sigstore.dev) is used.",
																							Attributes: map[string]schema.Attribute{
																								"ignore_tlog": schema.BoolAttribute{
																									Description:         "IgnoreTlog skips transparency log verification.",
																									MarkdownDescription: "IgnoreTlog skips transparency log verification.",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"pubkey": schema.StringAttribute{
																									Description:         "RekorPubKey is an optional PEM-encoded public key to use for a custom Rekor.If set, this will be used to validate transparency log signatures from a custom Rekor.",
																									MarkdownDescription: "RekorPubKey is an optional PEM-encoded public key to use for a custom Rekor.If set, this will be used to validate transparency log signatures from a custom Rekor.",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"url": schema.StringAttribute{
																									Description:         "URL is the address of the transparency log. Defaults to the public Rekor log instance https://rekor.sigstore.dev.",
																									MarkdownDescription: "URL is the address of the transparency log. Defaults to the public Rekor log instance https://rekor.sigstore.dev.",
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
																							Description:         "Reference to a Secret resource that contains a public key",
																							MarkdownDescription: "Reference to a Secret resource that contains a public key",
																							Attributes: map[string]schema.Attribute{
																								"name": schema.StringAttribute{
																									Description:         "Name of the secret. The provided secret must contain a key named cosign.pub.",
																									MarkdownDescription: "Name of the secret. The provided secret must contain a key named cosign.pub.",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"namespace": schema.StringAttribute{
																									Description:         "Namespace name where the Secret exists.",
																									MarkdownDescription: "Namespace name where the Secret exists.",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},
																							},
																							Required: false,
																							Optional: true,
																							Computed: false,
																						},

																						"signature_algorithm": schema.StringAttribute{
																							Description:         "Specify signature algorithm for public keys. Supported values are sha224, sha256, sha384 and sha512.",
																							MarkdownDescription: "Specify signature algorithm for public keys. Supported values are sha224, sha256, sha384 and sha512.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},
																					},
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"repository": schema.StringAttribute{
																					Description:         "Repository is an optional alternate OCI repository to use for signatures and attestations that match this rule.If specified Repository will override other OCI image repository locations for this Attestor.",
																					MarkdownDescription: "Repository is an optional alternate OCI repository to use for signatures and attestations that match this rule.If specified Repository will override other OCI image repository locations for this Attestor.",
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
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"conditions": schema.ListNestedAttribute{
															Description:         "Conditions are used to verify attributes within a Predicate. If no Conditions are specifiedthe attestation check is satisfied as long there are predicates that match the predicate type.",
															MarkdownDescription: "Conditions are used to verify attributes within a Predicate. If no Conditions are specifiedthe attestation check is satisfied as long there are predicates that match the predicate type.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"all": schema.ListNestedAttribute{
																		Description:         "AllConditions enable variable-based conditional rule execution. This is useful forfiner control of when an rule is applied. A condition can reference object datausing JMESPath notation.Here, all of the conditions need to pass",
																		MarkdownDescription: "AllConditions enable variable-based conditional rule execution. This is useful forfiner control of when an rule is applied. A condition can reference object datausing JMESPath notation.Here, all of the conditions need to pass",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"key": schema.MapAttribute{
																					Description:         "Key is the context entry (using JMESPath) for conditional rule evaluation.",
																					MarkdownDescription: "Key is the context entry (using JMESPath) for conditional rule evaluation.",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"message": schema.StringAttribute{
																					Description:         "Message is an optional display message",
																					MarkdownDescription: "Message is an optional display message",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"operator": schema.StringAttribute{
																					Description:         "Operator is the conditional operation to perform. Valid operators are:Equals, NotEquals, In, AnyIn, AllIn, NotIn, AnyNotIn, AllNotIn, GreaterThanOrEquals,GreaterThan, LessThanOrEquals, LessThan, DurationGreaterThanOrEquals, DurationGreaterThan,DurationLessThanOrEquals, DurationLessThan",
																					MarkdownDescription: "Operator is the conditional operation to perform. Valid operators are:Equals, NotEquals, In, AnyIn, AllIn, NotIn, AnyNotIn, AllNotIn, GreaterThanOrEquals,GreaterThan, LessThanOrEquals, LessThan, DurationGreaterThanOrEquals, DurationGreaterThan,DurationLessThanOrEquals, DurationLessThan",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																					Validators: []validator.String{
																						stringvalidator.OneOf("Equals", "NotEquals", "In", "AnyIn", "AllIn", "NotIn", "AnyNotIn", "AllNotIn", "GreaterThanOrEquals", "GreaterThan", "LessThanOrEquals", "LessThan", "DurationGreaterThanOrEquals", "DurationGreaterThan", "DurationLessThanOrEquals", "DurationLessThan"),
																					},
																				},

																				"value": schema.MapAttribute{
																					Description:         "Value is the conditional value, or set of values. The values can be fixed setor can be variables declared using JMESPath.",
																					MarkdownDescription: "Value is the conditional value, or set of values. The values can be fixed setor can be variables declared using JMESPath.",
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

																	"any": schema.ListNestedAttribute{
																		Description:         "AnyConditions enable variable-based conditional rule execution. This is useful forfiner control of when an rule is applied. A condition can reference object datausing JMESPath notation.Here, at least one of the conditions need to pass",
																		MarkdownDescription: "AnyConditions enable variable-based conditional rule execution. This is useful forfiner control of when an rule is applied. A condition can reference object datausing JMESPath notation.Here, at least one of the conditions need to pass",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"key": schema.MapAttribute{
																					Description:         "Key is the context entry (using JMESPath) for conditional rule evaluation.",
																					MarkdownDescription: "Key is the context entry (using JMESPath) for conditional rule evaluation.",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"message": schema.StringAttribute{
																					Description:         "Message is an optional display message",
																					MarkdownDescription: "Message is an optional display message",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"operator": schema.StringAttribute{
																					Description:         "Operator is the conditional operation to perform. Valid operators are:Equals, NotEquals, In, AnyIn, AllIn, NotIn, AnyNotIn, AllNotIn, GreaterThanOrEquals,GreaterThan, LessThanOrEquals, LessThan, DurationGreaterThanOrEquals, DurationGreaterThan,DurationLessThanOrEquals, DurationLessThan",
																					MarkdownDescription: "Operator is the conditional operation to perform. Valid operators are:Equals, NotEquals, In, AnyIn, AllIn, NotIn, AnyNotIn, AllNotIn, GreaterThanOrEquals,GreaterThan, LessThanOrEquals, LessThan, DurationGreaterThanOrEquals, DurationGreaterThan,DurationLessThanOrEquals, DurationLessThan",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																					Validators: []validator.String{
																						stringvalidator.OneOf("Equals", "NotEquals", "In", "AnyIn", "AllIn", "NotIn", "AnyNotIn", "AllNotIn", "GreaterThanOrEquals", "GreaterThan", "LessThanOrEquals", "LessThan", "DurationGreaterThanOrEquals", "DurationGreaterThan", "DurationLessThanOrEquals", "DurationLessThan"),
																					},
																				},

																				"value": schema.MapAttribute{
																					Description:         "Value is the conditional value, or set of values. The values can be fixed setor can be variables declared using JMESPath.",
																					MarkdownDescription: "Value is the conditional value, or set of values. The values can be fixed setor can be variables declared using JMESPath.",
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
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"predicate_type": schema.StringAttribute{
															Description:         "Deprecated in favour of 'Type', to be removed soon",
															MarkdownDescription: "Deprecated in favour of 'Type', to be removed soon",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"type": schema.StringAttribute{
															Description:         "Type defines the type of attestation contained within the Statement.",
															MarkdownDescription: "Type defines the type of attestation contained within the Statement.",
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

											"attestors": schema.ListNestedAttribute{
												Description:         "Attestors specified the required attestors (i.e. authorities)",
												MarkdownDescription: "Attestors specified the required attestors (i.e. authorities)",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"count": schema.Int64Attribute{
															Description:         "Count specifies the required number of entries that must match. If the count is null, all entries must match(a logical AND). If the count is 1, at least one entry must match (a logical OR). If the count contains avalue N, then N must be less than or equal to the size of entries, and at least N entries must match.",
															MarkdownDescription: "Count specifies the required number of entries that must match. If the count is null, all entries must match(a logical AND). If the count is 1, at least one entry must match (a logical OR). If the count contains avalue N, then N must be less than or equal to the size of entries, and at least N entries must match.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(1),
															},
														},

														"entries": schema.ListNestedAttribute{
															Description:         "Entries contains the available attestors. An attestor can be a static key,attributes for keyless verification, or a nested attestor declaration.",
															MarkdownDescription: "Entries contains the available attestors. An attestor can be a static key,attributes for keyless verification, or a nested attestor declaration.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"annotations": schema.MapAttribute{
																		Description:         "Annotations are used for image verification.Every specified key-value pair must exist and match in the verified payload.The payload may contain other key-value pairs.",
																		MarkdownDescription: "Annotations are used for image verification.Every specified key-value pair must exist and match in the verified payload.The payload may contain other key-value pairs.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"attestor": schema.MapAttribute{
																		Description:         "Attestor is a nested set of Attestor used to specify a more complex set of match authorities.",
																		MarkdownDescription: "Attestor is a nested set of Attestor used to specify a more complex set of match authorities.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"certificates": schema.SingleNestedAttribute{
																		Description:         "Certificates specifies one or more certificates.",
																		MarkdownDescription: "Certificates specifies one or more certificates.",
																		Attributes: map[string]schema.Attribute{
																			"cert": schema.StringAttribute{
																				Description:         "Cert is an optional PEM-encoded public certificate.",
																				MarkdownDescription: "Cert is an optional PEM-encoded public certificate.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"cert_chain": schema.StringAttribute{
																				Description:         "CertChain is an optional PEM encoded set of certificates used to verify.",
																				MarkdownDescription: "CertChain is an optional PEM encoded set of certificates used to verify.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"ctlog": schema.SingleNestedAttribute{
																				Description:         "CTLog (certificate timestamp log) provides a configuration for validation of Signed CertificateTimestamps (SCTs). If the value is unset, the default behavior by Cosign is used.",
																				MarkdownDescription: "CTLog (certificate timestamp log) provides a configuration for validation of Signed CertificateTimestamps (SCTs). If the value is unset, the default behavior by Cosign is used.",
																				Attributes: map[string]schema.Attribute{
																					"ignore_sct": schema.BoolAttribute{
																						Description:         "IgnoreSCT defines whether to use the Signed Certificate Timestamp (SCT) log to check for a certificatetimestamp. Default is false. Set to true if this was opted out during signing.",
																						MarkdownDescription: "IgnoreSCT defines whether to use the Signed Certificate Timestamp (SCT) log to check for a certificatetimestamp. Default is false. Set to true if this was opted out during signing.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"pubkey": schema.StringAttribute{
																						Description:         "PubKey, if set, is used to validate SCTs against a custom source.",
																						MarkdownDescription: "PubKey, if set, is used to validate SCTs against a custom source.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"tsa_cert_chain": schema.StringAttribute{
																						Description:         "TSACertChain, if set, is the PEM-encoded certificate chain file for the RFC3161 timestamp authority. Mustcontain the root CA certificate. Optionally may contain intermediate CA certificates, andmay contain the leaf TSA certificate if not present in the timestamurce.",
																						MarkdownDescription: "TSACertChain, if set, is the PEM-encoded certificate chain file for the RFC3161 timestamp authority. Mustcontain the root CA certificate. Optionally may contain intermediate CA certificates, andmay contain the leaf TSA certificate if not present in the timestamurce.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},
																				},
																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"rekor": schema.SingleNestedAttribute{
																				Description:         "Rekor provides configuration for the Rekor transparency log service. If an empty objectis provided the public instance of Rekor (https://rekor.sigstore.dev) is used.",
																				MarkdownDescription: "Rekor provides configuration for the Rekor transparency log service. If an empty objectis provided the public instance of Rekor (https://rekor.sigstore.dev) is used.",
																				Attributes: map[string]schema.Attribute{
																					"ignore_tlog": schema.BoolAttribute{
																						Description:         "IgnoreTlog skips transparency log verification.",
																						MarkdownDescription: "IgnoreTlog skips transparency log verification.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"pubkey": schema.StringAttribute{
																						Description:         "RekorPubKey is an optional PEM-encoded public key to use for a custom Rekor.If set, this will be used to validate transparency log signatures from a custom Rekor.",
																						MarkdownDescription: "RekorPubKey is an optional PEM-encoded public key to use for a custom Rekor.If set, this will be used to validate transparency log signatures from a custom Rekor.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"url": schema.StringAttribute{
																						Description:         "URL is the address of the transparency log. Defaults to the public Rekor log instance https://rekor.sigstore.dev.",
																						MarkdownDescription: "URL is the address of the transparency log. Defaults to the public Rekor log instance https://rekor.sigstore.dev.",
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

																	"keyless": schema.SingleNestedAttribute{
																		Description:         "Keyless is a set of attribute used to verify a Sigstore keyless attestor.See https://github.com/sigstore/cosign/blob/main/KEYLESS.md.",
																		MarkdownDescription: "Keyless is a set of attribute used to verify a Sigstore keyless attestor.See https://github.com/sigstore/cosign/blob/main/KEYLESS.md.",
																		Attributes: map[string]schema.Attribute{
																			"additional_extensions": schema.MapAttribute{
																				Description:         "AdditionalExtensions are certificate-extensions used for keyless signing.",
																				MarkdownDescription: "AdditionalExtensions are certificate-extensions used for keyless signing.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"ctlog": schema.SingleNestedAttribute{
																				Description:         "CTLog (certificate timestamp log) provides a configuration for validation of Signed CertificateTimestamps (SCTs). If the value is unset, the default behavior by Cosign is used.",
																				MarkdownDescription: "CTLog (certificate timestamp log) provides a configuration for validation of Signed CertificateTimestamps (SCTs). If the value is unset, the default behavior by Cosign is used.",
																				Attributes: map[string]schema.Attribute{
																					"ignore_sct": schema.BoolAttribute{
																						Description:         "IgnoreSCT defines whether to use the Signed Certificate Timestamp (SCT) log to check for a certificatetimestamp. Default is false. Set to true if this was opted out during signing.",
																						MarkdownDescription: "IgnoreSCT defines whether to use the Signed Certificate Timestamp (SCT) log to check for a certificatetimestamp. Default is false. Set to true if this was opted out during signing.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"pubkey": schema.StringAttribute{
																						Description:         "PubKey, if set, is used to validate SCTs against a custom source.",
																						MarkdownDescription: "PubKey, if set, is used to validate SCTs against a custom source.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"tsa_cert_chain": schema.StringAttribute{
																						Description:         "TSACertChain, if set, is the PEM-encoded certificate chain file for the RFC3161 timestamp authority. Mustcontain the root CA certificate. Optionally may contain intermediate CA certificates, andmay contain the leaf TSA certificate if not present in the timestamurce.",
																						MarkdownDescription: "TSACertChain, if set, is the PEM-encoded certificate chain file for the RFC3161 timestamp authority. Mustcontain the root CA certificate. Optionally may contain intermediate CA certificates, andmay contain the leaf TSA certificate if not present in the timestamurce.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},
																				},
																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"issuer": schema.StringAttribute{
																				Description:         "Issuer is the certificate issuer used for keyless signing.",
																				MarkdownDescription: "Issuer is the certificate issuer used for keyless signing.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"rekor": schema.SingleNestedAttribute{
																				Description:         "Rekor provides configuration for the Rekor transparency log service. If an empty objectis provided the public instance of Rekor (https://rekor.sigstore.dev) is used.",
																				MarkdownDescription: "Rekor provides configuration for the Rekor transparency log service. If an empty objectis provided the public instance of Rekor (https://rekor.sigstore.dev) is used.",
																				Attributes: map[string]schema.Attribute{
																					"ignore_tlog": schema.BoolAttribute{
																						Description:         "IgnoreTlog skips transparency log verification.",
																						MarkdownDescription: "IgnoreTlog skips transparency log verification.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"pubkey": schema.StringAttribute{
																						Description:         "RekorPubKey is an optional PEM-encoded public key to use for a custom Rekor.If set, this will be used to validate transparency log signatures from a custom Rekor.",
																						MarkdownDescription: "RekorPubKey is an optional PEM-encoded public key to use for a custom Rekor.If set, this will be used to validate transparency log signatures from a custom Rekor.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"url": schema.StringAttribute{
																						Description:         "URL is the address of the transparency log. Defaults to the public Rekor log instance https://rekor.sigstore.dev.",
																						MarkdownDescription: "URL is the address of the transparency log. Defaults to the public Rekor log instance https://rekor.sigstore.dev.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},
																				},
																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"roots": schema.StringAttribute{
																				Description:         "Roots is an optional set of PEM encoded trusted root certificates.If not provided, the system roots are used.",
																				MarkdownDescription: "Roots is an optional set of PEM encoded trusted root certificates.If not provided, the system roots are used.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"subject": schema.StringAttribute{
																				Description:         "Subject is the verified identity used for keyless signing, for example the email address.",
																				MarkdownDescription: "Subject is the verified identity used for keyless signing, for example the email address.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"keys": schema.SingleNestedAttribute{
																		Description:         "Keys specifies one or more public keys.",
																		MarkdownDescription: "Keys specifies one or more public keys.",
																		Attributes: map[string]schema.Attribute{
																			"ctlog": schema.SingleNestedAttribute{
																				Description:         "CTLog (certificate timestamp log) provides a configuration for validation of Signed CertificateTimestamps (SCTs). If the value is unset, the default behavior by Cosign is used.",
																				MarkdownDescription: "CTLog (certificate timestamp log) provides a configuration for validation of Signed CertificateTimestamps (SCTs). If the value is unset, the default behavior by Cosign is used.",
																				Attributes: map[string]schema.Attribute{
																					"ignore_sct": schema.BoolAttribute{
																						Description:         "IgnoreSCT defines whether to use the Signed Certificate Timestamp (SCT) log to check for a certificatetimestamp. Default is false. Set to true if this was opted out during signing.",
																						MarkdownDescription: "IgnoreSCT defines whether to use the Signed Certificate Timestamp (SCT) log to check for a certificatetimestamp. Default is false. Set to true if this was opted out during signing.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"pubkey": schema.StringAttribute{
																						Description:         "PubKey, if set, is used to validate SCTs against a custom source.",
																						MarkdownDescription: "PubKey, if set, is used to validate SCTs against a custom source.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"tsa_cert_chain": schema.StringAttribute{
																						Description:         "TSACertChain, if set, is the PEM-encoded certificate chain file for the RFC3161 timestamp authority. Mustcontain the root CA certificate. Optionally may contain intermediate CA certificates, andmay contain the leaf TSA certificate if not present in the timestamurce.",
																						MarkdownDescription: "TSACertChain, if set, is the PEM-encoded certificate chain file for the RFC3161 timestamp authority. Mustcontain the root CA certificate. Optionally may contain intermediate CA certificates, andmay contain the leaf TSA certificate if not present in the timestamurce.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},
																				},
																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"kms": schema.StringAttribute{
																				Description:         "KMS provides the URI to the public key stored in a Key Management System. See:https://github.com/sigstore/cosign/blob/main/KMS.md",
																				MarkdownDescription: "KMS provides the URI to the public key stored in a Key Management System. See:https://github.com/sigstore/cosign/blob/main/KMS.md",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"public_keys": schema.StringAttribute{
																				Description:         "Keys is a set of X.509 public keys used to verify image signatures. The keys can be directlyspecified or can be a variable reference to a key specified in a ConfigMap (seehttps://kyverno.io/docs/writing-policies/variables/), or reference a standard Kubernetes Secretelsewhere in the cluster by specifying it in the format 'k8s://<namespace>/<secret_name>'.The named Secret must specify a key 'cosign.pub' containing the public key used forverification, (see https://github.com/sigstore/cosign/blob/main/KMS.md#kubernetes-secret).When multiple keys are specified each key is processed as a separate staticKey entry(.attestors[*].entries.keys) within the set of attestors and the count is applied across the keys.",
																				MarkdownDescription: "Keys is a set of X.509 public keys used to verify image signatures. The keys can be directlyspecified or can be a variable reference to a key specified in a ConfigMap (seehttps://kyverno.io/docs/writing-policies/variables/), or reference a standard Kubernetes Secretelsewhere in the cluster by specifying it in the format 'k8s://<namespace>/<secret_name>'.The named Secret must specify a key 'cosign.pub' containing the public key used forverification, (see https://github.com/sigstore/cosign/blob/main/KMS.md#kubernetes-secret).When multiple keys are specified each key is processed as a separate staticKey entry(.attestors[*].entries.keys) within the set of attestors and the count is applied across the keys.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"rekor": schema.SingleNestedAttribute{
																				Description:         "Rekor provides configuration for the Rekor transparency log service. If an empty objectis provided the public instance of Rekor (https://rekor.sigstore.dev) is used.",
																				MarkdownDescription: "Rekor provides configuration for the Rekor transparency log service. If an empty objectis provided the public instance of Rekor (https://rekor.sigstore.dev) is used.",
																				Attributes: map[string]schema.Attribute{
																					"ignore_tlog": schema.BoolAttribute{
																						Description:         "IgnoreTlog skips transparency log verification.",
																						MarkdownDescription: "IgnoreTlog skips transparency log verification.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"pubkey": schema.StringAttribute{
																						Description:         "RekorPubKey is an optional PEM-encoded public key to use for a custom Rekor.If set, this will be used to validate transparency log signatures from a custom Rekor.",
																						MarkdownDescription: "RekorPubKey is an optional PEM-encoded public key to use for a custom Rekor.If set, this will be used to validate transparency log signatures from a custom Rekor.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"url": schema.StringAttribute{
																						Description:         "URL is the address of the transparency log. Defaults to the public Rekor log instance https://rekor.sigstore.dev.",
																						MarkdownDescription: "URL is the address of the transparency log. Defaults to the public Rekor log instance https://rekor.sigstore.dev.",
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
																				Description:         "Reference to a Secret resource that contains a public key",
																				MarkdownDescription: "Reference to a Secret resource that contains a public key",
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "Name of the secret. The provided secret must contain a key named cosign.pub.",
																						MarkdownDescription: "Name of the secret. The provided secret must contain a key named cosign.pub.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"namespace": schema.StringAttribute{
																						Description:         "Namespace name where the Secret exists.",
																						MarkdownDescription: "Namespace name where the Secret exists.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},
																				},
																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"signature_algorithm": schema.StringAttribute{
																				Description:         "Specify signature algorithm for public keys. Supported values are sha224, sha256, sha384 and sha512.",
																				MarkdownDescription: "Specify signature algorithm for public keys. Supported values are sha224, sha256, sha384 and sha512.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"repository": schema.StringAttribute{
																		Description:         "Repository is an optional alternate OCI repository to use for signatures and attestations that match this rule.If specified Repository will override other OCI image repository locations for this Attestor.",
																		MarkdownDescription: "Repository is an optional alternate OCI repository to use for signatures and attestations that match this rule.If specified Repository will override other OCI image repository locations for this Attestor.",
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
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"cosign_oci11": schema.BoolAttribute{
												Description:         "CosignOCI11 enables the experimental OCI 1.1 behaviour in cosign image verification.Defaults to false.",
												MarkdownDescription: "CosignOCI11 enables the experimental OCI 1.1 behaviour in cosign image verification.Defaults to false.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"image": schema.StringAttribute{
												Description:         "Deprecated. Use ImageReferences instead.",
												MarkdownDescription: "Deprecated. Use ImageReferences instead.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"image_references": schema.ListAttribute{
												Description:         "ImageReferences is a list of matching image reference patterns. At least one pattern in thelist must match the image for the rule to apply. Each image reference consists of a registryaddress (defaults to docker.io), repository, image, and tag (defaults to latest).Wildcards ('*' and '?') are allowed. See: https://kubernetes.io/docs/concepts/containers/images.",
												MarkdownDescription: "ImageReferences is a list of matching image reference patterns. At least one pattern in thelist must match the image for the rule to apply. Each image reference consists of a registryaddress (defaults to docker.io), repository, image, and tag (defaults to latest).Wildcards ('*' and '?') are allowed. See: https://kubernetes.io/docs/concepts/containers/images.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"image_registry_credentials": schema.SingleNestedAttribute{
												Description:         "ImageRegistryCredentials provides credentials that will be used for authentication with registry.",
												MarkdownDescription: "ImageRegistryCredentials provides credentials that will be used for authentication with registry.",
												Attributes: map[string]schema.Attribute{
													"allow_insecure_registry": schema.BoolAttribute{
														Description:         "AllowInsecureRegistry allows insecure access to a registry.",
														MarkdownDescription: "AllowInsecureRegistry allows insecure access to a registry.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"providers": schema.ListAttribute{
														Description:         "Providers specifies a list of OCI Registry names, whose authentication providers are provided.It can be of one of these values: default,google,azure,amazon,github.",
														MarkdownDescription: "Providers specifies a list of OCI Registry names, whose authentication providers are provided.It can be of one of these values: default,google,azure,amazon,github.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"secrets": schema.ListAttribute{
														Description:         "Secrets specifies a list of secrets that are provided for credentials.Secrets must live in the Kyverno namespace.",
														MarkdownDescription: "Secrets specifies a list of secrets that are provided for credentials.Secrets must live in the Kyverno namespace.",
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

											"issuer": schema.StringAttribute{
												Description:         "Deprecated. Use KeylessAttestor instead.",
												MarkdownDescription: "Deprecated. Use KeylessAttestor instead.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"key": schema.StringAttribute{
												Description:         "Deprecated. Use StaticKeyAttestor instead.",
												MarkdownDescription: "Deprecated. Use StaticKeyAttestor instead.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"mutate_digest": schema.BoolAttribute{
												Description:         "MutateDigest enables replacement of image tags with digests.Defaults to true.",
												MarkdownDescription: "MutateDigest enables replacement of image tags with digests.Defaults to true.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"repository": schema.StringAttribute{
												Description:         "Repository is an optional alternate OCI repository to use for image signatures and attestations that match this rule.If specified Repository will override the default OCI image repository configured for the installation.The repository can also be overridden per Attestor or Attestation.",
												MarkdownDescription: "Repository is an optional alternate OCI repository to use for image signatures and attestations that match this rule.If specified Repository will override the default OCI image repository configured for the installation.The repository can also be overridden per Attestor or Attestation.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"required": schema.BoolAttribute{
												Description:         "Required validates that images are verified i.e. have matched passed a signature or attestation check.",
												MarkdownDescription: "Required validates that images are verified i.e. have matched passed a signature or attestation check.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"roots": schema.StringAttribute{
												Description:         "Deprecated. Use KeylessAttestor instead.",
												MarkdownDescription: "Deprecated. Use KeylessAttestor instead.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"skip_image_references": schema.ListAttribute{
												Description:         "SkipImageReferences is a list of matching image reference patterns that should be skipped.At least one pattern in the list must match the image for the rule to be skipped. Each image referenceconsists of a registry address (defaults to docker.io), repository, image, and tag (defaults to latest).Wildcards ('*' and '?') are allowed. See: https://kubernetes.io/docs/concepts/containers/images.",
												MarkdownDescription: "SkipImageReferences is a list of matching image reference patterns that should be skipped.At least one pattern in the list must match the image for the rule to be skipped. Each image referenceconsists of a registry address (defaults to docker.io), repository, image, and tag (defaults to latest).Wildcards ('*' and '?') are allowed. See: https://kubernetes.io/docs/concepts/containers/images.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"subject": schema.StringAttribute{
												Description:         "Deprecated. Use KeylessAttestor instead.",
												MarkdownDescription: "Deprecated. Use KeylessAttestor instead.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
												Description:         "Type specifies the method of signature validation. The allowed optionsare Cosign and Notary. By default Cosign is used if a type is not specified.",
												MarkdownDescription: "Type specifies the method of signature validation. The allowed optionsare Cosign and Notary. By default Cosign is used if a type is not specified.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("Cosign", "Notary"),
												},
											},

											"use_cache": schema.BoolAttribute{
												Description:         "UseCache enables caching of image verify responses for this rule.",
												MarkdownDescription: "UseCache enables caching of image verify responses for this rule.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"verify_digest": schema.BoolAttribute{
												Description:         "VerifyDigest validates that images have a digest.",
												MarkdownDescription: "VerifyDigest validates that images have a digest.",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"schema_validation": schema.BoolAttribute{
						Description:         "Deprecated.",
						MarkdownDescription: "Deprecated.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"use_server_side_apply": schema.BoolAttribute{
						Description:         "UseServerSideApply controls whether to use server-side apply for generate rulesIf is set to 'true' create & update for generate rules will use apply instead of create/update.Defaults to 'false' if not specified.",
						MarkdownDescription: "UseServerSideApply controls whether to use server-side apply for generate rulesIf is set to 'true' create & update for generate rules will use apply instead of create/update.Defaults to 'false' if not specified.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"validation_failure_action": schema.StringAttribute{
						Description:         "Deprecated, use validationFailureAction under the validate rule instead.",
						MarkdownDescription: "Deprecated, use validationFailureAction under the validate rule instead.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("audit", "enforce", "Audit", "Enforce"),
						},
					},

					"validation_failure_action_overrides": schema.ListNestedAttribute{
						Description:         "Deprecated, use validationFailureActionOverrides under the validate rule instead.",
						MarkdownDescription: "Deprecated, use validationFailureActionOverrides under the validate rule instead.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"action": schema.StringAttribute{
									Description:         "ValidationFailureAction defines the policy validation failure action",
									MarkdownDescription: "ValidationFailureAction defines the policy validation failure action",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("audit", "enforce", "Audit", "Enforce"),
									},
								},

								"namespace_selector": schema.SingleNestedAttribute{
									Description:         "A label selector is a label query over a set of resources. The result of matchLabels andmatchExpressions are ANDed. An empty label selector matches all objects. A nulllabel selector matches no objects.",
									MarkdownDescription: "A label selector is a label query over a set of resources. The result of matchLabels andmatchExpressions are ANDed. An empty label selector matches all objects. A nulllabel selector matches no objects.",
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
									Required: false,
									Optional: true,
									Computed: false,
								},

								"namespaces": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"webhook_configuration": schema.SingleNestedAttribute{
						Description:         "WebhookConfiguration specifies the custom configuration for Kubernetes admission webhookconfiguration.",
						MarkdownDescription: "WebhookConfiguration specifies the custom configuration for Kubernetes admission webhookconfiguration.",
						Attributes: map[string]schema.Attribute{
							"failure_policy": schema.StringAttribute{
								Description:         "FailurePolicy defines how unexpected policy errors and webhook response timeout errors are handled.Rules within the same policy share the same failure behavior.This field should not be accessed directly, instead 'GetFailurePolicy()' should be used.Allowed values are Ignore or Fail. Defaults to Fail.",
								MarkdownDescription: "FailurePolicy defines how unexpected policy errors and webhook response timeout errors are handled.Rules within the same policy share the same failure behavior.This field should not be accessed directly, instead 'GetFailurePolicy()' should be used.Allowed values are Ignore or Fail. Defaults to Fail.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Ignore", "Fail"),
								},
							},

							"match_conditions": schema.ListNestedAttribute{
								Description:         "MatchCondition configures admission webhook matchConditions.Requires Kubernetes 1.27 or later.",
								MarkdownDescription: "MatchCondition configures admission webhook matchConditions.Requires Kubernetes 1.27 or later.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"expression": schema.StringAttribute{
											Description:         "Expression represents the expression which will be evaluated by CEL. Must evaluate to bool.CEL expressions have access to the contents of the AdmissionRequest and Authorizer, organized into CEL variables:'object' - The object from the incoming request. The value is null for DELETE requests.'oldObject' - The existing object. The value is null for CREATE requests.'request' - Attributes of the admission request(/pkg/apis/admission/types.go#AdmissionRequest).'authorizer' - A CEL Authorizer. May be used to perform authorization checks for the principal (user or service account) of the request.  See https://pkg.go.dev/k8s.io/apiserver/pkg/cel/library#Authz'authorizer.requestResource' - A CEL ResourceCheck constructed from the 'authorizer' and configured with the  request resource.Documentation on CEL: https://kubernetes.io/docs/reference/using-api/cel/Required.",
											MarkdownDescription: "Expression represents the expression which will be evaluated by CEL. Must evaluate to bool.CEL expressions have access to the contents of the AdmissionRequest and Authorizer, organized into CEL variables:'object' - The object from the incoming request. The value is null for DELETE requests.'oldObject' - The existing object. The value is null for CREATE requests.'request' - Attributes of the admission request(/pkg/apis/admission/types.go#AdmissionRequest).'authorizer' - A CEL Authorizer. May be used to perform authorization checks for the principal (user or service account) of the request.  See https://pkg.go.dev/k8s.io/apiserver/pkg/cel/library#Authz'authorizer.requestResource' - A CEL ResourceCheck constructed from the 'authorizer' and configured with the  request resource.Documentation on CEL: https://kubernetes.io/docs/reference/using-api/cel/Required.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name is an identifier for this match condition, used for strategic merging of MatchConditions,as well as providing an identifier for logging purposes. A good name should be descriptive ofthe associated expression.Name must be a qualified name consisting of alphanumeric characters, '-', '_' or '.', andmust start and end with an alphanumeric character (e.g. 'MyName',  or 'my.name',  or'123-abc', regex used for validation is '([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9]') with anoptional DNS subdomain prefix and '/' (e.g. 'example.com/MyName')Required.",
											MarkdownDescription: "Name is an identifier for this match condition, used for strategic merging of MatchConditions,as well as providing an identifier for logging purposes. A good name should be descriptive ofthe associated expression.Name must be a qualified name consisting of alphanumeric characters, '-', '_' or '.', andmust start and end with an alphanumeric character (e.g. 'MyName',  or 'my.name',  or'123-abc', regex used for validation is '([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9]') with anoptional DNS subdomain prefix and '/' (e.g. 'example.com/MyName')Required.",
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

							"timeout_seconds": schema.Int64Attribute{
								Description:         "TimeoutSeconds specifies the maximum time in seconds allowed to apply this policy.After the configured time expires, the admission request may fail, or may simply ignore the policy results,based on the failure policy. The default timeout is 10s, the value must be between 1 and 30 seconds.",
								MarkdownDescription: "TimeoutSeconds specifies the maximum time in seconds allowed to apply this policy.After the configured time expires, the admission request may fail, or may simply ignore the policy results,based on the failure policy. The default timeout is 10s, the value must be between 1 and 30 seconds.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"webhook_timeout_seconds": schema.Int64Attribute{
						Description:         "Deprecated, use webhookTimeoutSeconds under webhookConfiguration instead.",
						MarkdownDescription: "Deprecated, use webhookTimeoutSeconds under webhookConfiguration instead.",
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
	}
}

func (r *KyvernoIoPolicyV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kyverno_io_policy_v1_manifest")

	var model KyvernoIoPolicyV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("kyverno.io/v1")
	model.Kind = pointer.String("Policy")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
