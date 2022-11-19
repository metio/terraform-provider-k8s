/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

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

type KyvernoIoPolicyV1Resource struct{}

var (
	_ resource.Resource = (*KyvernoIoPolicyV1Resource)(nil)
)

type KyvernoIoPolicyV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type KyvernoIoPolicyV1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		ApplyRules *string `tfsdk:"apply_rules" yaml:"applyRules,omitempty"`

		Background *bool `tfsdk:"background" yaml:"background,omitempty"`

		FailurePolicy *string `tfsdk:"failure_policy" yaml:"failurePolicy,omitempty"`

		GenerateExistingOnPolicyUpdate *bool `tfsdk:"generate_existing_on_policy_update" yaml:"generateExistingOnPolicyUpdate,omitempty"`

		MutateExistingOnPolicyUpdate *bool `tfsdk:"mutate_existing_on_policy_update" yaml:"mutateExistingOnPolicyUpdate,omitempty"`

		Rules *[]struct {
			Context *[]struct {
				ApiCall *struct {
					JmesPath *string `tfsdk:"jmes_path" yaml:"jmesPath,omitempty"`

					UrlPath *string `tfsdk:"url_path" yaml:"urlPath,omitempty"`
				} `tfsdk:"api_call" yaml:"apiCall,omitempty"`

				ConfigMap *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
				} `tfsdk:"config_map" yaml:"configMap,omitempty"`

				ImageRegistry *struct {
					JmesPath *string `tfsdk:"jmes_path" yaml:"jmesPath,omitempty"`

					Reference *string `tfsdk:"reference" yaml:"reference,omitempty"`
				} `tfsdk:"image_registry" yaml:"imageRegistry,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Variable *struct {
					Default utilities.Dynamic `tfsdk:"default" yaml:"default,omitempty"`

					JmesPath *string `tfsdk:"jmes_path" yaml:"jmesPath,omitempty"`

					Value utilities.Dynamic `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"variable" yaml:"variable,omitempty"`
			} `tfsdk:"context" yaml:"context,omitempty"`

			Exclude *struct {
				All *[]struct {
					ClusterRoles *[]string `tfsdk:"cluster_roles" yaml:"clusterRoles,omitempty"`

					Resources *struct {
						Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Kinds *[]string `tfsdk:"kinds" yaml:"kinds,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Names *[]string `tfsdk:"names" yaml:"names,omitempty"`

						NamespaceSelector *struct {
							MatchExpressions *[]struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

								Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
							} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

							MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
						} `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`

						Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

						Selector *struct {
							MatchExpressions *[]struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

								Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
							} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

							MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
						} `tfsdk:"selector" yaml:"selector,omitempty"`
					} `tfsdk:"resources" yaml:"resources,omitempty"`

					Roles *[]string `tfsdk:"roles" yaml:"roles,omitempty"`

					Subjects *[]struct {
						ApiGroup *string `tfsdk:"api_group" yaml:"apiGroup,omitempty"`

						Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"subjects" yaml:"subjects,omitempty"`
				} `tfsdk:"all" yaml:"all,omitempty"`

				Any *[]struct {
					ClusterRoles *[]string `tfsdk:"cluster_roles" yaml:"clusterRoles,omitempty"`

					Resources *struct {
						Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Kinds *[]string `tfsdk:"kinds" yaml:"kinds,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Names *[]string `tfsdk:"names" yaml:"names,omitempty"`

						NamespaceSelector *struct {
							MatchExpressions *[]struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

								Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
							} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

							MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
						} `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`

						Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

						Selector *struct {
							MatchExpressions *[]struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

								Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
							} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

							MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
						} `tfsdk:"selector" yaml:"selector,omitempty"`
					} `tfsdk:"resources" yaml:"resources,omitempty"`

					Roles *[]string `tfsdk:"roles" yaml:"roles,omitempty"`

					Subjects *[]struct {
						ApiGroup *string `tfsdk:"api_group" yaml:"apiGroup,omitempty"`

						Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"subjects" yaml:"subjects,omitempty"`
				} `tfsdk:"any" yaml:"any,omitempty"`

				ClusterRoles *[]string `tfsdk:"cluster_roles" yaml:"clusterRoles,omitempty"`

				Resources *struct {
					Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

					Kinds *[]string `tfsdk:"kinds" yaml:"kinds,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Names *[]string `tfsdk:"names" yaml:"names,omitempty"`

					NamespaceSelector *struct {
						MatchExpressions *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
						} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

						MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
					} `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`

					Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

					Selector *struct {
						MatchExpressions *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
						} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

						MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
					} `tfsdk:"selector" yaml:"selector,omitempty"`
				} `tfsdk:"resources" yaml:"resources,omitempty"`

				Roles *[]string `tfsdk:"roles" yaml:"roles,omitempty"`

				Subjects *[]struct {
					ApiGroup *string `tfsdk:"api_group" yaml:"apiGroup,omitempty"`

					Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
				} `tfsdk:"subjects" yaml:"subjects,omitempty"`
			} `tfsdk:"exclude" yaml:"exclude,omitempty"`

			Generate *struct {
				ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

				Clone *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
				} `tfsdk:"clone" yaml:"clone,omitempty"`

				CloneList *struct {
					Kinds *[]string `tfsdk:"kinds" yaml:"kinds,omitempty"`

					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

					Selector *struct {
						MatchExpressions *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
						} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

						MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
					} `tfsdk:"selector" yaml:"selector,omitempty"`
				} `tfsdk:"clone_list" yaml:"cloneList,omitempty"`

				Data utilities.Dynamic `tfsdk:"data" yaml:"data,omitempty"`

				Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

				Synchronize *bool `tfsdk:"synchronize" yaml:"synchronize,omitempty"`
			} `tfsdk:"generate" yaml:"generate,omitempty"`

			ImageExtractors *map[string]string `tfsdk:"image_extractors" yaml:"imageExtractors,omitempty"`

			Match *struct {
				All *[]struct {
					ClusterRoles *[]string `tfsdk:"cluster_roles" yaml:"clusterRoles,omitempty"`

					Resources *struct {
						Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Kinds *[]string `tfsdk:"kinds" yaml:"kinds,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Names *[]string `tfsdk:"names" yaml:"names,omitempty"`

						NamespaceSelector *struct {
							MatchExpressions *[]struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

								Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
							} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

							MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
						} `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`

						Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

						Selector *struct {
							MatchExpressions *[]struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

								Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
							} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

							MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
						} `tfsdk:"selector" yaml:"selector,omitempty"`
					} `tfsdk:"resources" yaml:"resources,omitempty"`

					Roles *[]string `tfsdk:"roles" yaml:"roles,omitempty"`

					Subjects *[]struct {
						ApiGroup *string `tfsdk:"api_group" yaml:"apiGroup,omitempty"`

						Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"subjects" yaml:"subjects,omitempty"`
				} `tfsdk:"all" yaml:"all,omitempty"`

				Any *[]struct {
					ClusterRoles *[]string `tfsdk:"cluster_roles" yaml:"clusterRoles,omitempty"`

					Resources *struct {
						Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Kinds *[]string `tfsdk:"kinds" yaml:"kinds,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Names *[]string `tfsdk:"names" yaml:"names,omitempty"`

						NamespaceSelector *struct {
							MatchExpressions *[]struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

								Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
							} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

							MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
						} `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`

						Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

						Selector *struct {
							MatchExpressions *[]struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

								Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
							} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

							MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
						} `tfsdk:"selector" yaml:"selector,omitempty"`
					} `tfsdk:"resources" yaml:"resources,omitempty"`

					Roles *[]string `tfsdk:"roles" yaml:"roles,omitempty"`

					Subjects *[]struct {
						ApiGroup *string `tfsdk:"api_group" yaml:"apiGroup,omitempty"`

						Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"subjects" yaml:"subjects,omitempty"`
				} `tfsdk:"any" yaml:"any,omitempty"`

				ClusterRoles *[]string `tfsdk:"cluster_roles" yaml:"clusterRoles,omitempty"`

				Resources *struct {
					Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

					Kinds *[]string `tfsdk:"kinds" yaml:"kinds,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Names *[]string `tfsdk:"names" yaml:"names,omitempty"`

					NamespaceSelector *struct {
						MatchExpressions *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
						} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

						MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
					} `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`

					Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

					Selector *struct {
						MatchExpressions *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
						} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

						MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
					} `tfsdk:"selector" yaml:"selector,omitempty"`
				} `tfsdk:"resources" yaml:"resources,omitempty"`

				Roles *[]string `tfsdk:"roles" yaml:"roles,omitempty"`

				Subjects *[]struct {
					ApiGroup *string `tfsdk:"api_group" yaml:"apiGroup,omitempty"`

					Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
				} `tfsdk:"subjects" yaml:"subjects,omitempty"`
			} `tfsdk:"match" yaml:"match,omitempty"`

			Mutate *struct {
				Foreach *[]struct {
					Context *[]struct {
						ApiCall *struct {
							JmesPath *string `tfsdk:"jmes_path" yaml:"jmesPath,omitempty"`

							UrlPath *string `tfsdk:"url_path" yaml:"urlPath,omitempty"`
						} `tfsdk:"api_call" yaml:"apiCall,omitempty"`

						ConfigMap *struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
						} `tfsdk:"config_map" yaml:"configMap,omitempty"`

						ImageRegistry *struct {
							JmesPath *string `tfsdk:"jmes_path" yaml:"jmesPath,omitempty"`

							Reference *string `tfsdk:"reference" yaml:"reference,omitempty"`
						} `tfsdk:"image_registry" yaml:"imageRegistry,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Variable *struct {
							Default utilities.Dynamic `tfsdk:"default" yaml:"default,omitempty"`

							JmesPath *string `tfsdk:"jmes_path" yaml:"jmesPath,omitempty"`

							Value utilities.Dynamic `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"variable" yaml:"variable,omitempty"`
					} `tfsdk:"context" yaml:"context,omitempty"`

					List *string `tfsdk:"list" yaml:"list,omitempty"`

					PatchStrategicMerge utilities.Dynamic `tfsdk:"patch_strategic_merge" yaml:"patchStrategicMerge,omitempty"`

					PatchesJson6902 *string `tfsdk:"patches_json6902" yaml:"patchesJson6902,omitempty"`

					Preconditions *struct {
						All *[]struct {
							Key utilities.Dynamic `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Value utilities.Dynamic `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"all" yaml:"all,omitempty"`

						Any *[]struct {
							Key utilities.Dynamic `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Value utilities.Dynamic `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"any" yaml:"any,omitempty"`
					} `tfsdk:"preconditions" yaml:"preconditions,omitempty"`
				} `tfsdk:"foreach" yaml:"foreach,omitempty"`

				PatchStrategicMerge utilities.Dynamic `tfsdk:"patch_strategic_merge" yaml:"patchStrategicMerge,omitempty"`

				PatchesJson6902 *string `tfsdk:"patches_json6902" yaml:"patchesJson6902,omitempty"`

				Targets *[]struct {
					ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

					Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
				} `tfsdk:"targets" yaml:"targets,omitempty"`
			} `tfsdk:"mutate" yaml:"mutate,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Preconditions utilities.Dynamic `tfsdk:"preconditions" yaml:"preconditions,omitempty"`

			Validate *struct {
				AnyPattern utilities.Dynamic `tfsdk:"any_pattern" yaml:"anyPattern,omitempty"`

				Deny *struct {
					Conditions utilities.Dynamic `tfsdk:"conditions" yaml:"conditions,omitempty"`
				} `tfsdk:"deny" yaml:"deny,omitempty"`

				Foreach *[]struct {
					AnyPattern utilities.Dynamic `tfsdk:"any_pattern" yaml:"anyPattern,omitempty"`

					Context *[]struct {
						ApiCall *struct {
							JmesPath *string `tfsdk:"jmes_path" yaml:"jmesPath,omitempty"`

							UrlPath *string `tfsdk:"url_path" yaml:"urlPath,omitempty"`
						} `tfsdk:"api_call" yaml:"apiCall,omitempty"`

						ConfigMap *struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
						} `tfsdk:"config_map" yaml:"configMap,omitempty"`

						ImageRegistry *struct {
							JmesPath *string `tfsdk:"jmes_path" yaml:"jmesPath,omitempty"`

							Reference *string `tfsdk:"reference" yaml:"reference,omitempty"`
						} `tfsdk:"image_registry" yaml:"imageRegistry,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Variable *struct {
							Default utilities.Dynamic `tfsdk:"default" yaml:"default,omitempty"`

							JmesPath *string `tfsdk:"jmes_path" yaml:"jmesPath,omitempty"`

							Value utilities.Dynamic `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"variable" yaml:"variable,omitempty"`
					} `tfsdk:"context" yaml:"context,omitempty"`

					Deny *struct {
						Conditions utilities.Dynamic `tfsdk:"conditions" yaml:"conditions,omitempty"`
					} `tfsdk:"deny" yaml:"deny,omitempty"`

					ElementScope *bool `tfsdk:"element_scope" yaml:"elementScope,omitempty"`

					List *string `tfsdk:"list" yaml:"list,omitempty"`

					Pattern utilities.Dynamic `tfsdk:"pattern" yaml:"pattern,omitempty"`

					Preconditions *struct {
						All *[]struct {
							Key utilities.Dynamic `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Value utilities.Dynamic `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"all" yaml:"all,omitempty"`

						Any *[]struct {
							Key utilities.Dynamic `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Value utilities.Dynamic `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"any" yaml:"any,omitempty"`
					} `tfsdk:"preconditions" yaml:"preconditions,omitempty"`
				} `tfsdk:"foreach" yaml:"foreach,omitempty"`

				Manifests *struct {
					AnnotationDomain *string `tfsdk:"annotation_domain" yaml:"annotationDomain,omitempty"`

					Attestors *[]struct {
						Count *int64 `tfsdk:"count" yaml:"count,omitempty"`

						Entries *[]struct {
							Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

							Attestor utilities.Dynamic `tfsdk:"attestor" yaml:"attestor,omitempty"`

							Certificates *struct {
								Cert *string `tfsdk:"cert" yaml:"cert,omitempty"`

								CertChain *string `tfsdk:"cert_chain" yaml:"certChain,omitempty"`

								Rekor *struct {
									Url *string `tfsdk:"url" yaml:"url,omitempty"`
								} `tfsdk:"rekor" yaml:"rekor,omitempty"`
							} `tfsdk:"certificates" yaml:"certificates,omitempty"`

							Keyless *struct {
								AdditionalExtensions *map[string]string `tfsdk:"additional_extensions" yaml:"additionalExtensions,omitempty"`

								Issuer *string `tfsdk:"issuer" yaml:"issuer,omitempty"`

								Rekor *struct {
									Url *string `tfsdk:"url" yaml:"url,omitempty"`
								} `tfsdk:"rekor" yaml:"rekor,omitempty"`

								Roots *string `tfsdk:"roots" yaml:"roots,omitempty"`

								Subject *string `tfsdk:"subject" yaml:"subject,omitempty"`
							} `tfsdk:"keyless" yaml:"keyless,omitempty"`

							Keys *struct {
								Kms *string `tfsdk:"kms" yaml:"kms,omitempty"`

								PublicKeys *string `tfsdk:"public_keys" yaml:"publicKeys,omitempty"`

								Rekor *struct {
									Url *string `tfsdk:"url" yaml:"url,omitempty"`
								} `tfsdk:"rekor" yaml:"rekor,omitempty"`

								Secret *struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
								} `tfsdk:"secret" yaml:"secret,omitempty"`

								SignatureAlgorithm *string `tfsdk:"signature_algorithm" yaml:"signatureAlgorithm,omitempty"`
							} `tfsdk:"keys" yaml:"keys,omitempty"`

							Repository *string `tfsdk:"repository" yaml:"repository,omitempty"`
						} `tfsdk:"entries" yaml:"entries,omitempty"`
					} `tfsdk:"attestors" yaml:"attestors,omitempty"`

					DryRun *struct {
						Enable *bool `tfsdk:"enable" yaml:"enable,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"dry_run" yaml:"dryRun,omitempty"`

					IgnoreFields *[]struct {
						Fields *[]string `tfsdk:"fields" yaml:"fields,omitempty"`

						Objects *[]struct {
							Group *string `tfsdk:"group" yaml:"group,omitempty"`

							Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

							Version *string `tfsdk:"version" yaml:"version,omitempty"`
						} `tfsdk:"objects" yaml:"objects,omitempty"`
					} `tfsdk:"ignore_fields" yaml:"ignoreFields,omitempty"`

					Repository *string `tfsdk:"repository" yaml:"repository,omitempty"`
				} `tfsdk:"manifests" yaml:"manifests,omitempty"`

				Message *string `tfsdk:"message" yaml:"message,omitempty"`

				Pattern utilities.Dynamic `tfsdk:"pattern" yaml:"pattern,omitempty"`

				PodSecurity *struct {
					Exclude *[]struct {
						ControlName *string `tfsdk:"control_name" yaml:"controlName,omitempty"`

						Images *[]string `tfsdk:"images" yaml:"images,omitempty"`
					} `tfsdk:"exclude" yaml:"exclude,omitempty"`

					Level *string `tfsdk:"level" yaml:"level,omitempty"`

					Version *string `tfsdk:"version" yaml:"version,omitempty"`
				} `tfsdk:"pod_security" yaml:"podSecurity,omitempty"`
			} `tfsdk:"validate" yaml:"validate,omitempty"`

			VerifyImages *[]struct {
				AdditionalExtensions *map[string]string `tfsdk:"additional_extensions" yaml:"additionalExtensions,omitempty"`

				Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

				Attestations *[]struct {
					Conditions *[]struct {
						All *[]struct {
							Key utilities.Dynamic `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Value utilities.Dynamic `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"all" yaml:"all,omitempty"`

						Any *[]struct {
							Key utilities.Dynamic `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Value utilities.Dynamic `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"any" yaml:"any,omitempty"`
					} `tfsdk:"conditions" yaml:"conditions,omitempty"`

					PredicateType *string `tfsdk:"predicate_type" yaml:"predicateType,omitempty"`
				} `tfsdk:"attestations" yaml:"attestations,omitempty"`

				Attestors *[]struct {
					Count *int64 `tfsdk:"count" yaml:"count,omitempty"`

					Entries *[]struct {
						Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Attestor utilities.Dynamic `tfsdk:"attestor" yaml:"attestor,omitempty"`

						Certificates *struct {
							Cert *string `tfsdk:"cert" yaml:"cert,omitempty"`

							CertChain *string `tfsdk:"cert_chain" yaml:"certChain,omitempty"`

							Rekor *struct {
								Url *string `tfsdk:"url" yaml:"url,omitempty"`
							} `tfsdk:"rekor" yaml:"rekor,omitempty"`
						} `tfsdk:"certificates" yaml:"certificates,omitempty"`

						Keyless *struct {
							AdditionalExtensions *map[string]string `tfsdk:"additional_extensions" yaml:"additionalExtensions,omitempty"`

							Issuer *string `tfsdk:"issuer" yaml:"issuer,omitempty"`

							Rekor *struct {
								Url *string `tfsdk:"url" yaml:"url,omitempty"`
							} `tfsdk:"rekor" yaml:"rekor,omitempty"`

							Roots *string `tfsdk:"roots" yaml:"roots,omitempty"`

							Subject *string `tfsdk:"subject" yaml:"subject,omitempty"`
						} `tfsdk:"keyless" yaml:"keyless,omitempty"`

						Keys *struct {
							Kms *string `tfsdk:"kms" yaml:"kms,omitempty"`

							PublicKeys *string `tfsdk:"public_keys" yaml:"publicKeys,omitempty"`

							Rekor *struct {
								Url *string `tfsdk:"url" yaml:"url,omitempty"`
							} `tfsdk:"rekor" yaml:"rekor,omitempty"`

							Secret *struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
							} `tfsdk:"secret" yaml:"secret,omitempty"`

							SignatureAlgorithm *string `tfsdk:"signature_algorithm" yaml:"signatureAlgorithm,omitempty"`
						} `tfsdk:"keys" yaml:"keys,omitempty"`

						Repository *string `tfsdk:"repository" yaml:"repository,omitempty"`
					} `tfsdk:"entries" yaml:"entries,omitempty"`
				} `tfsdk:"attestors" yaml:"attestors,omitempty"`

				Image *string `tfsdk:"image" yaml:"image,omitempty"`

				ImageReferences *[]string `tfsdk:"image_references" yaml:"imageReferences,omitempty"`

				Issuer *string `tfsdk:"issuer" yaml:"issuer,omitempty"`

				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				MutateDigest *bool `tfsdk:"mutate_digest" yaml:"mutateDigest,omitempty"`

				Repository *string `tfsdk:"repository" yaml:"repository,omitempty"`

				Required *bool `tfsdk:"required" yaml:"required,omitempty"`

				Roots *string `tfsdk:"roots" yaml:"roots,omitempty"`

				Subject *string `tfsdk:"subject" yaml:"subject,omitempty"`

				VerifyDigest *bool `tfsdk:"verify_digest" yaml:"verifyDigest,omitempty"`
			} `tfsdk:"verify_images" yaml:"verifyImages,omitempty"`
		} `tfsdk:"rules" yaml:"rules,omitempty"`

		SchemaValidation *bool `tfsdk:"schema_validation" yaml:"schemaValidation,omitempty"`

		ValidationFailureAction *string `tfsdk:"validation_failure_action" yaml:"validationFailureAction,omitempty"`

		ValidationFailureActionOverrides *[]struct {
			Action *string `tfsdk:"action" yaml:"action,omitempty"`

			Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`
		} `tfsdk:"validation_failure_action_overrides" yaml:"validationFailureActionOverrides,omitempty"`

		WebhookTimeoutSeconds *int64 `tfsdk:"webhook_timeout_seconds" yaml:"webhookTimeoutSeconds,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewKyvernoIoPolicyV1Resource() resource.Resource {
	return &KyvernoIoPolicyV1Resource{}
}

func (r *KyvernoIoPolicyV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_kyverno_io_policy_v1"
}

func (r *KyvernoIoPolicyV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Policy declares validation, mutation, and generation behaviors for matching resources. See: https://kyverno.io/docs/writing-policies/ for more information.",
		MarkdownDescription: "Policy declares validation, mutation, and generation behaviors for matching resources. See: https://kyverno.io/docs/writing-policies/ for more information.",
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
				Description:         "Spec defines policy behaviors and contains one or more rules.",
				MarkdownDescription: "Spec defines policy behaviors and contains one or more rules.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"apply_rules": {
						Description:         "ApplyRules controls how rules in a policy are applied. Rule are processed in the order of declaration. When set to 'One' processing stops after a rule has been applied i.e. the rule matches and results in a pass, fail, or error. When set to 'All' all rules in the policy are processed. The default is 'All'.",
						MarkdownDescription: "ApplyRules controls how rules in a policy are applied. Rule are processed in the order of declaration. When set to 'One' processing stops after a rule has been applied i.e. the rule matches and results in a pass, fail, or error. When set to 'All' all rules in the policy are processed. The default is 'All'.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("All", "One"),
						},
					},

					"background": {
						Description:         "Background controls if rules are applied to existing resources during a background scan. Optional. Default value is 'true'. The value must be set to 'false' if the policy rule uses variables that are only available in the admission review request (e.g. user name).",
						MarkdownDescription: "Background controls if rules are applied to existing resources during a background scan. Optional. Default value is 'true'. The value must be set to 'false' if the policy rule uses variables that are only available in the admission review request (e.g. user name).",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"failure_policy": {
						Description:         "FailurePolicy defines how unexpected policy errors and webhook response timeout errors are handled. Rules within the same policy share the same failure behavior. This field should not be accessed directly, instead 'GetFailurePolicy()' should be used. Allowed values are Ignore or Fail. Defaults to Fail.",
						MarkdownDescription: "FailurePolicy defines how unexpected policy errors and webhook response timeout errors are handled. Rules within the same policy share the same failure behavior. This field should not be accessed directly, instead 'GetFailurePolicy()' should be used. Allowed values are Ignore or Fail. Defaults to Fail.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("Ignore", "Fail"),
						},
					},

					"generate_existing_on_policy_update": {
						Description:         "GenerateExistingOnPolicyUpdate controls whether to trigger generate rule in existing resources If is set to 'true' generate rule will be triggered and applied to existing matched resources. Defaults to 'false' if not specified.",
						MarkdownDescription: "GenerateExistingOnPolicyUpdate controls whether to trigger generate rule in existing resources If is set to 'true' generate rule will be triggered and applied to existing matched resources. Defaults to 'false' if not specified.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"mutate_existing_on_policy_update": {
						Description:         "MutateExistingOnPolicyUpdate controls if a mutateExisting policy is applied on policy events. Default value is 'false'.",
						MarkdownDescription: "MutateExistingOnPolicyUpdate controls if a mutateExisting policy is applied on policy events. Default value is 'false'.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"rules": {
						Description:         "Rules is a list of Rule instances. A Policy contains multiple rules and each rule can validate, mutate, or generate resources.",
						MarkdownDescription: "Rules is a list of Rule instances. A Policy contains multiple rules and each rule can validate, mutate, or generate resources.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"context": {
								Description:         "Context defines variables and data sources that can be used during rule execution.",
								MarkdownDescription: "Context defines variables and data sources that can be used during rule execution.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"api_call": {
										Description:         "APICall defines an HTTP request to the Kubernetes API server. The JSON data retrieved is stored in the context.",
										MarkdownDescription: "APICall defines an HTTP request to the Kubernetes API server. The JSON data retrieved is stored in the context.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"jmes_path": {
												Description:         "JMESPath is an optional JSON Match Expression that can be used to transform the JSON response returned from the API server. For example a JMESPath of 'items | length(@)' applied to the API server response to the URLPath '/apis/apps/v1/deployments' will return the total count of deployments across all namespaces.",
												MarkdownDescription: "JMESPath is an optional JSON Match Expression that can be used to transform the JSON response returned from the API server. For example a JMESPath of 'items | length(@)' applied to the API server response to the URLPath '/apis/apps/v1/deployments' will return the total count of deployments across all namespaces.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"url_path": {
												Description:         "URLPath is the URL path to be used in the HTTP GET request to the Kubernetes API server (e.g. '/api/v1/namespaces' or  '/apis/apps/v1/deployments'). The format required is the same format used by the 'kubectl get --raw' command.",
												MarkdownDescription: "URLPath is the URL path to be used in the HTTP GET request to the Kubernetes API server (e.g. '/api/v1/namespaces' or  '/apis/apps/v1/deployments'). The format required is the same format used by the 'kubectl get --raw' command.",

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

									"config_map": {
										Description:         "ConfigMap is the ConfigMap reference.",
										MarkdownDescription: "ConfigMap is the ConfigMap reference.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "Name is the ConfigMap name.",
												MarkdownDescription: "Name is the ConfigMap name.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"namespace": {
												Description:         "Namespace is the ConfigMap namespace.",
												MarkdownDescription: "Namespace is the ConfigMap namespace.",

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

									"image_registry": {
										Description:         "ImageRegistry defines requests to an OCI/Docker V2 registry to fetch image details.",
										MarkdownDescription: "ImageRegistry defines requests to an OCI/Docker V2 registry to fetch image details.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"jmes_path": {
												Description:         "JMESPath is an optional JSON Match Expression that can be used to transform the ImageData struct returned as a result of processing the image reference.",
												MarkdownDescription: "JMESPath is an optional JSON Match Expression that can be used to transform the ImageData struct returned as a result of processing the image reference.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"reference": {
												Description:         "Reference is image reference to a container image in the registry. Example: ghcr.io/kyverno/kyverno:latest",
												MarkdownDescription: "Reference is image reference to a container image in the registry. Example: ghcr.io/kyverno/kyverno:latest",

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

									"name": {
										Description:         "Name is the variable name.",
										MarkdownDescription: "Name is the variable name.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"variable": {
										Description:         "Variable defines an arbitrary JMESPath context variable that can be defined inline.",
										MarkdownDescription: "Variable defines an arbitrary JMESPath context variable that can be defined inline.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"default": {
												Description:         "Default is an optional arbitrary JSON object that the variable may take if the JMESPath expression evaluates to nil",
												MarkdownDescription: "Default is an optional arbitrary JSON object that the variable may take if the JMESPath expression evaluates to nil",

												Type: utilities.DynamicType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"jmes_path": {
												Description:         "JMESPath is an optional JMESPath Expression that can be used to transform the variable.",
												MarkdownDescription: "JMESPath is an optional JMESPath Expression that can be used to transform the variable.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value": {
												Description:         "Value is any arbitrary JSON object representable in YAML or JSON form.",
												MarkdownDescription: "Value is any arbitrary JSON object representable in YAML or JSON form.",

												Type: utilities.DynamicType{},

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

							"exclude": {
								Description:         "ExcludeResources defines when this policy rule should not be applied. The exclude criteria can include resource information (e.g. kind, name, namespace, labels) and admission review request information like the name or role.",
								MarkdownDescription: "ExcludeResources defines when this policy rule should not be applied. The exclude criteria can include resource information (e.g. kind, name, namespace, labels) and admission review request information like the name or role.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"all": {
										Description:         "All allows specifying resources which will be ANDed",
										MarkdownDescription: "All allows specifying resources which will be ANDed",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"cluster_roles": {
												Description:         "ClusterRoles is the list of cluster-wide role names for the user.",
												MarkdownDescription: "ClusterRoles is the list of cluster-wide role names for the user.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"resources": {
												Description:         "ResourceDescription contains information about the resource being created or modified.",
												MarkdownDescription: "ResourceDescription contains information about the resource being created or modified.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations is a  map of annotations (key-value pairs of type string). Annotation keys and values support the wildcard characters '*' (matches zero or many characters) and '?' (matches at least one character).",
														MarkdownDescription: "Annotations is a  map of annotations (key-value pairs of type string). Annotation keys and values support the wildcard characters '*' (matches zero or many characters) and '?' (matches at least one character).",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"kinds": {
														Description:         "Kinds is a list of resource kinds.",
														MarkdownDescription: "Kinds is a list of resource kinds.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": {
														Description:         "Name is the name of the resource. The name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character). NOTE: 'Name' is being deprecated in favor of 'Names'.",
														MarkdownDescription: "Name is the name of the resource. The name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character). NOTE: 'Name' is being deprecated in favor of 'Names'.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"names": {
														Description:         "Names are the names of the resources. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",
														MarkdownDescription: "Names are the names of the resources. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespace_selector": {
														Description:         "NamespaceSelector is a label selector for the resource namespace. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character).Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",
														MarkdownDescription: "NamespaceSelector is a label selector for the resource namespace. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character).Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",

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
														Description:         "Namespaces is a list of namespaces names. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",
														MarkdownDescription: "Namespaces is a list of namespaces names. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"selector": {
														Description:         "Selector is a label selector. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character). Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",
														MarkdownDescription: "Selector is a label selector. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character). Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",

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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"roles": {
												Description:         "Roles is the list of namespaced role names for the user.",
												MarkdownDescription: "Roles is the list of namespaced role names for the user.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"subjects": {
												Description:         "Subjects is the list of subject names like users, user groups, and service accounts.",
												MarkdownDescription: "Subjects is the list of subject names like users, user groups, and service accounts.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"api_group": {
														Description:         "APIGroup holds the API group of the referenced subject. Defaults to '' for ServiceAccount subjects. Defaults to 'rbac.authorization.k8s.io' for User and Group subjects.",
														MarkdownDescription: "APIGroup holds the API group of the referenced subject. Defaults to '' for ServiceAccount subjects. Defaults to 'rbac.authorization.k8s.io' for User and Group subjects.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"kind": {
														Description:         "Kind of object being referenced. Values defined by this API group are 'User', 'Group', and 'ServiceAccount'. If the Authorizer does not recognized the kind value, the Authorizer should report an error.",
														MarkdownDescription: "Kind of object being referenced. Values defined by this API group are 'User', 'Group', and 'ServiceAccount'. If the Authorizer does not recognized the kind value, the Authorizer should report an error.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"name": {
														Description:         "Name of the object being referenced.",
														MarkdownDescription: "Name of the object being referenced.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"namespace": {
														Description:         "Namespace of the referenced object.  If the object kind is non-namespace, such as 'User' or 'Group', and this value is not empty the Authorizer should report an error.",
														MarkdownDescription: "Namespace of the referenced object.  If the object kind is non-namespace, such as 'User' or 'Group', and this value is not empty the Authorizer should report an error.",

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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"any": {
										Description:         "Any allows specifying resources which will be ORed",
										MarkdownDescription: "Any allows specifying resources which will be ORed",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"cluster_roles": {
												Description:         "ClusterRoles is the list of cluster-wide role names for the user.",
												MarkdownDescription: "ClusterRoles is the list of cluster-wide role names for the user.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"resources": {
												Description:         "ResourceDescription contains information about the resource being created or modified.",
												MarkdownDescription: "ResourceDescription contains information about the resource being created or modified.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations is a  map of annotations (key-value pairs of type string). Annotation keys and values support the wildcard characters '*' (matches zero or many characters) and '?' (matches at least one character).",
														MarkdownDescription: "Annotations is a  map of annotations (key-value pairs of type string). Annotation keys and values support the wildcard characters '*' (matches zero or many characters) and '?' (matches at least one character).",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"kinds": {
														Description:         "Kinds is a list of resource kinds.",
														MarkdownDescription: "Kinds is a list of resource kinds.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": {
														Description:         "Name is the name of the resource. The name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character). NOTE: 'Name' is being deprecated in favor of 'Names'.",
														MarkdownDescription: "Name is the name of the resource. The name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character). NOTE: 'Name' is being deprecated in favor of 'Names'.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"names": {
														Description:         "Names are the names of the resources. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",
														MarkdownDescription: "Names are the names of the resources. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespace_selector": {
														Description:         "NamespaceSelector is a label selector for the resource namespace. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character).Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",
														MarkdownDescription: "NamespaceSelector is a label selector for the resource namespace. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character).Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",

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
														Description:         "Namespaces is a list of namespaces names. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",
														MarkdownDescription: "Namespaces is a list of namespaces names. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"selector": {
														Description:         "Selector is a label selector. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character). Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",
														MarkdownDescription: "Selector is a label selector. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character). Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",

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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"roles": {
												Description:         "Roles is the list of namespaced role names for the user.",
												MarkdownDescription: "Roles is the list of namespaced role names for the user.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"subjects": {
												Description:         "Subjects is the list of subject names like users, user groups, and service accounts.",
												MarkdownDescription: "Subjects is the list of subject names like users, user groups, and service accounts.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"api_group": {
														Description:         "APIGroup holds the API group of the referenced subject. Defaults to '' for ServiceAccount subjects. Defaults to 'rbac.authorization.k8s.io' for User and Group subjects.",
														MarkdownDescription: "APIGroup holds the API group of the referenced subject. Defaults to '' for ServiceAccount subjects. Defaults to 'rbac.authorization.k8s.io' for User and Group subjects.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"kind": {
														Description:         "Kind of object being referenced. Values defined by this API group are 'User', 'Group', and 'ServiceAccount'. If the Authorizer does not recognized the kind value, the Authorizer should report an error.",
														MarkdownDescription: "Kind of object being referenced. Values defined by this API group are 'User', 'Group', and 'ServiceAccount'. If the Authorizer does not recognized the kind value, the Authorizer should report an error.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"name": {
														Description:         "Name of the object being referenced.",
														MarkdownDescription: "Name of the object being referenced.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"namespace": {
														Description:         "Namespace of the referenced object.  If the object kind is non-namespace, such as 'User' or 'Group', and this value is not empty the Authorizer should report an error.",
														MarkdownDescription: "Namespace of the referenced object.  If the object kind is non-namespace, such as 'User' or 'Group', and this value is not empty the Authorizer should report an error.",

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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"cluster_roles": {
										Description:         "ClusterRoles is the list of cluster-wide role names for the user.",
										MarkdownDescription: "ClusterRoles is the list of cluster-wide role names for the user.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"resources": {
										Description:         "ResourceDescription contains information about the resource being created or modified. Requires at least one tag to be specified when under MatchResources. Specifying ResourceDescription directly under match is being deprecated. Please specify under 'any' or 'all' instead.",
										MarkdownDescription: "ResourceDescription contains information about the resource being created or modified. Requires at least one tag to be specified when under MatchResources. Specifying ResourceDescription directly under match is being deprecated. Please specify under 'any' or 'all' instead.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"annotations": {
												Description:         "Annotations is a  map of annotations (key-value pairs of type string). Annotation keys and values support the wildcard characters '*' (matches zero or many characters) and '?' (matches at least one character).",
												MarkdownDescription: "Annotations is a  map of annotations (key-value pairs of type string). Annotation keys and values support the wildcard characters '*' (matches zero or many characters) and '?' (matches at least one character).",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"kinds": {
												Description:         "Kinds is a list of resource kinds.",
												MarkdownDescription: "Kinds is a list of resource kinds.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "Name is the name of the resource. The name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character). NOTE: 'Name' is being deprecated in favor of 'Names'.",
												MarkdownDescription: "Name is the name of the resource. The name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character). NOTE: 'Name' is being deprecated in favor of 'Names'.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"names": {
												Description:         "Names are the names of the resources. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",
												MarkdownDescription: "Names are the names of the resources. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"namespace_selector": {
												Description:         "NamespaceSelector is a label selector for the resource namespace. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character).Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",
												MarkdownDescription: "NamespaceSelector is a label selector for the resource namespace. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character).Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",

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
												Description:         "Namespaces is a list of namespaces names. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",
												MarkdownDescription: "Namespaces is a list of namespaces names. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"selector": {
												Description:         "Selector is a label selector. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character). Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",
												MarkdownDescription: "Selector is a label selector. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character). Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",

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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"roles": {
										Description:         "Roles is the list of namespaced role names for the user.",
										MarkdownDescription: "Roles is the list of namespaced role names for the user.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"subjects": {
										Description:         "Subjects is the list of subject names like users, user groups, and service accounts.",
										MarkdownDescription: "Subjects is the list of subject names like users, user groups, and service accounts.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"api_group": {
												Description:         "APIGroup holds the API group of the referenced subject. Defaults to '' for ServiceAccount subjects. Defaults to 'rbac.authorization.k8s.io' for User and Group subjects.",
												MarkdownDescription: "APIGroup holds the API group of the referenced subject. Defaults to '' for ServiceAccount subjects. Defaults to 'rbac.authorization.k8s.io' for User and Group subjects.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"kind": {
												Description:         "Kind of object being referenced. Values defined by this API group are 'User', 'Group', and 'ServiceAccount'. If the Authorizer does not recognized the kind value, the Authorizer should report an error.",
												MarkdownDescription: "Kind of object being referenced. Values defined by this API group are 'User', 'Group', and 'ServiceAccount'. If the Authorizer does not recognized the kind value, the Authorizer should report an error.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"name": {
												Description:         "Name of the object being referenced.",
												MarkdownDescription: "Name of the object being referenced.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"namespace": {
												Description:         "Namespace of the referenced object.  If the object kind is non-namespace, such as 'User' or 'Group', and this value is not empty the Authorizer should report an error.",
												MarkdownDescription: "Namespace of the referenced object.  If the object kind is non-namespace, such as 'User' or 'Group', and this value is not empty the Authorizer should report an error.",

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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"generate": {
								Description:         "Generation is used to create new resources.",
								MarkdownDescription: "Generation is used to create new resources.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"api_version": {
										Description:         "APIVersion specifies resource apiVersion.",
										MarkdownDescription: "APIVersion specifies resource apiVersion.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"clone": {
										Description:         "Clone specifies the source resource used to populate each generated resource. At most one of Data or Clone can be specified. If neither are provided, the generated resource will be created with default data only.",
										MarkdownDescription: "Clone specifies the source resource used to populate each generated resource. At most one of Data or Clone can be specified. If neither are provided, the generated resource will be created with default data only.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "Name specifies name of the resource.",
												MarkdownDescription: "Name specifies name of the resource.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"namespace": {
												Description:         "Namespace specifies source resource namespace.",
												MarkdownDescription: "Namespace specifies source resource namespace.",

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

									"clone_list": {
										Description:         "CloneList specifies the list of source resource used to populate each generated resource.",
										MarkdownDescription: "CloneList specifies the list of source resource used to populate each generated resource.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"kinds": {
												Description:         "Kinds is a list of resource kinds.",
												MarkdownDescription: "Kinds is a list of resource kinds.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"namespace": {
												Description:         "Namespace specifies source resource namespace.",
												MarkdownDescription: "Namespace specifies source resource namespace.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"selector": {
												Description:         "Selector is a label selector. Label keys and values in 'matchLabels'. wildcard characters are not supported.",
												MarkdownDescription: "Selector is a label selector. Label keys and values in 'matchLabels'. wildcard characters are not supported.",

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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"data": {
										Description:         "Data provides the resource declaration used to populate each generated resource. At most one of Data or Clone must be specified. If neither are provided, the generated resource will be created with default data only.",
										MarkdownDescription: "Data provides the resource declaration used to populate each generated resource. At most one of Data or Clone must be specified. If neither are provided, the generated resource will be created with default data only.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"kind": {
										Description:         "Kind specifies resource kind.",
										MarkdownDescription: "Kind specifies resource kind.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "Name specifies the resource name.",
										MarkdownDescription: "Name specifies the resource name.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"namespace": {
										Description:         "Namespace specifies resource namespace.",
										MarkdownDescription: "Namespace specifies resource namespace.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"synchronize": {
										Description:         "Synchronize controls if generated resources should be kept in-sync with their source resource. If Synchronize is set to 'true' changes to generated resources will be overwritten with resource data from Data or the resource specified in the Clone declaration. Optional. Defaults to 'false' if not specified.",
										MarkdownDescription: "Synchronize controls if generated resources should be kept in-sync with their source resource. If Synchronize is set to 'true' changes to generated resources will be overwritten with resource data from Data or the resource specified in the Clone declaration. Optional. Defaults to 'false' if not specified.",

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

							"image_extractors": {
								Description:         "ImageExtractors defines a mapping from kinds to ImageExtractorConfigs. This config is only valid for verifyImages rules.",
								MarkdownDescription: "ImageExtractors defines a mapping from kinds to ImageExtractorConfigs. This config is only valid for verifyImages rules.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"match": {
								Description:         "MatchResources defines when this policy rule should be applied. The match criteria can include resource information (e.g. kind, name, namespace, labels) and admission review request information like the user name or role. At least one kind is required.",
								MarkdownDescription: "MatchResources defines when this policy rule should be applied. The match criteria can include resource information (e.g. kind, name, namespace, labels) and admission review request information like the user name or role. At least one kind is required.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"all": {
										Description:         "All allows specifying resources which will be ANDed",
										MarkdownDescription: "All allows specifying resources which will be ANDed",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"cluster_roles": {
												Description:         "ClusterRoles is the list of cluster-wide role names for the user.",
												MarkdownDescription: "ClusterRoles is the list of cluster-wide role names for the user.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"resources": {
												Description:         "ResourceDescription contains information about the resource being created or modified.",
												MarkdownDescription: "ResourceDescription contains information about the resource being created or modified.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations is a  map of annotations (key-value pairs of type string). Annotation keys and values support the wildcard characters '*' (matches zero or many characters) and '?' (matches at least one character).",
														MarkdownDescription: "Annotations is a  map of annotations (key-value pairs of type string). Annotation keys and values support the wildcard characters '*' (matches zero or many characters) and '?' (matches at least one character).",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"kinds": {
														Description:         "Kinds is a list of resource kinds.",
														MarkdownDescription: "Kinds is a list of resource kinds.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": {
														Description:         "Name is the name of the resource. The name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character). NOTE: 'Name' is being deprecated in favor of 'Names'.",
														MarkdownDescription: "Name is the name of the resource. The name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character). NOTE: 'Name' is being deprecated in favor of 'Names'.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"names": {
														Description:         "Names are the names of the resources. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",
														MarkdownDescription: "Names are the names of the resources. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespace_selector": {
														Description:         "NamespaceSelector is a label selector for the resource namespace. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character).Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",
														MarkdownDescription: "NamespaceSelector is a label selector for the resource namespace. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character).Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",

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
														Description:         "Namespaces is a list of namespaces names. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",
														MarkdownDescription: "Namespaces is a list of namespaces names. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"selector": {
														Description:         "Selector is a label selector. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character). Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",
														MarkdownDescription: "Selector is a label selector. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character). Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",

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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"roles": {
												Description:         "Roles is the list of namespaced role names for the user.",
												MarkdownDescription: "Roles is the list of namespaced role names for the user.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"subjects": {
												Description:         "Subjects is the list of subject names like users, user groups, and service accounts.",
												MarkdownDescription: "Subjects is the list of subject names like users, user groups, and service accounts.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"api_group": {
														Description:         "APIGroup holds the API group of the referenced subject. Defaults to '' for ServiceAccount subjects. Defaults to 'rbac.authorization.k8s.io' for User and Group subjects.",
														MarkdownDescription: "APIGroup holds the API group of the referenced subject. Defaults to '' for ServiceAccount subjects. Defaults to 'rbac.authorization.k8s.io' for User and Group subjects.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"kind": {
														Description:         "Kind of object being referenced. Values defined by this API group are 'User', 'Group', and 'ServiceAccount'. If the Authorizer does not recognized the kind value, the Authorizer should report an error.",
														MarkdownDescription: "Kind of object being referenced. Values defined by this API group are 'User', 'Group', and 'ServiceAccount'. If the Authorizer does not recognized the kind value, the Authorizer should report an error.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"name": {
														Description:         "Name of the object being referenced.",
														MarkdownDescription: "Name of the object being referenced.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"namespace": {
														Description:         "Namespace of the referenced object.  If the object kind is non-namespace, such as 'User' or 'Group', and this value is not empty the Authorizer should report an error.",
														MarkdownDescription: "Namespace of the referenced object.  If the object kind is non-namespace, such as 'User' or 'Group', and this value is not empty the Authorizer should report an error.",

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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"any": {
										Description:         "Any allows specifying resources which will be ORed",
										MarkdownDescription: "Any allows specifying resources which will be ORed",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"cluster_roles": {
												Description:         "ClusterRoles is the list of cluster-wide role names for the user.",
												MarkdownDescription: "ClusterRoles is the list of cluster-wide role names for the user.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"resources": {
												Description:         "ResourceDescription contains information about the resource being created or modified.",
												MarkdownDescription: "ResourceDescription contains information about the resource being created or modified.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations is a  map of annotations (key-value pairs of type string). Annotation keys and values support the wildcard characters '*' (matches zero or many characters) and '?' (matches at least one character).",
														MarkdownDescription: "Annotations is a  map of annotations (key-value pairs of type string). Annotation keys and values support the wildcard characters '*' (matches zero or many characters) and '?' (matches at least one character).",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"kinds": {
														Description:         "Kinds is a list of resource kinds.",
														MarkdownDescription: "Kinds is a list of resource kinds.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": {
														Description:         "Name is the name of the resource. The name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character). NOTE: 'Name' is being deprecated in favor of 'Names'.",
														MarkdownDescription: "Name is the name of the resource. The name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character). NOTE: 'Name' is being deprecated in favor of 'Names'.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"names": {
														Description:         "Names are the names of the resources. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",
														MarkdownDescription: "Names are the names of the resources. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespace_selector": {
														Description:         "NamespaceSelector is a label selector for the resource namespace. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character).Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",
														MarkdownDescription: "NamespaceSelector is a label selector for the resource namespace. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character).Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",

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
														Description:         "Namespaces is a list of namespaces names. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",
														MarkdownDescription: "Namespaces is a list of namespaces names. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"selector": {
														Description:         "Selector is a label selector. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character). Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",
														MarkdownDescription: "Selector is a label selector. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character). Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",

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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"roles": {
												Description:         "Roles is the list of namespaced role names for the user.",
												MarkdownDescription: "Roles is the list of namespaced role names for the user.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"subjects": {
												Description:         "Subjects is the list of subject names like users, user groups, and service accounts.",
												MarkdownDescription: "Subjects is the list of subject names like users, user groups, and service accounts.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"api_group": {
														Description:         "APIGroup holds the API group of the referenced subject. Defaults to '' for ServiceAccount subjects. Defaults to 'rbac.authorization.k8s.io' for User and Group subjects.",
														MarkdownDescription: "APIGroup holds the API group of the referenced subject. Defaults to '' for ServiceAccount subjects. Defaults to 'rbac.authorization.k8s.io' for User and Group subjects.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"kind": {
														Description:         "Kind of object being referenced. Values defined by this API group are 'User', 'Group', and 'ServiceAccount'. If the Authorizer does not recognized the kind value, the Authorizer should report an error.",
														MarkdownDescription: "Kind of object being referenced. Values defined by this API group are 'User', 'Group', and 'ServiceAccount'. If the Authorizer does not recognized the kind value, the Authorizer should report an error.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"name": {
														Description:         "Name of the object being referenced.",
														MarkdownDescription: "Name of the object being referenced.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"namespace": {
														Description:         "Namespace of the referenced object.  If the object kind is non-namespace, such as 'User' or 'Group', and this value is not empty the Authorizer should report an error.",
														MarkdownDescription: "Namespace of the referenced object.  If the object kind is non-namespace, such as 'User' or 'Group', and this value is not empty the Authorizer should report an error.",

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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"cluster_roles": {
										Description:         "ClusterRoles is the list of cluster-wide role names for the user.",
										MarkdownDescription: "ClusterRoles is the list of cluster-wide role names for the user.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"resources": {
										Description:         "ResourceDescription contains information about the resource being created or modified. Requires at least one tag to be specified when under MatchResources. Specifying ResourceDescription directly under match is being deprecated. Please specify under 'any' or 'all' instead.",
										MarkdownDescription: "ResourceDescription contains information about the resource being created or modified. Requires at least one tag to be specified when under MatchResources. Specifying ResourceDescription directly under match is being deprecated. Please specify under 'any' or 'all' instead.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"annotations": {
												Description:         "Annotations is a  map of annotations (key-value pairs of type string). Annotation keys and values support the wildcard characters '*' (matches zero or many characters) and '?' (matches at least one character).",
												MarkdownDescription: "Annotations is a  map of annotations (key-value pairs of type string). Annotation keys and values support the wildcard characters '*' (matches zero or many characters) and '?' (matches at least one character).",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"kinds": {
												Description:         "Kinds is a list of resource kinds.",
												MarkdownDescription: "Kinds is a list of resource kinds.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "Name is the name of the resource. The name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character). NOTE: 'Name' is being deprecated in favor of 'Names'.",
												MarkdownDescription: "Name is the name of the resource. The name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character). NOTE: 'Name' is being deprecated in favor of 'Names'.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"names": {
												Description:         "Names are the names of the resources. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",
												MarkdownDescription: "Names are the names of the resources. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"namespace_selector": {
												Description:         "NamespaceSelector is a label selector for the resource namespace. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character).Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",
												MarkdownDescription: "NamespaceSelector is a label selector for the resource namespace. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character).Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",

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
												Description:         "Namespaces is a list of namespaces names. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",
												MarkdownDescription: "Namespaces is a list of namespaces names. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"selector": {
												Description:         "Selector is a label selector. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character). Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",
												MarkdownDescription: "Selector is a label selector. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character). Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",

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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"roles": {
										Description:         "Roles is the list of namespaced role names for the user.",
										MarkdownDescription: "Roles is the list of namespaced role names for the user.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"subjects": {
										Description:         "Subjects is the list of subject names like users, user groups, and service accounts.",
										MarkdownDescription: "Subjects is the list of subject names like users, user groups, and service accounts.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"api_group": {
												Description:         "APIGroup holds the API group of the referenced subject. Defaults to '' for ServiceAccount subjects. Defaults to 'rbac.authorization.k8s.io' for User and Group subjects.",
												MarkdownDescription: "APIGroup holds the API group of the referenced subject. Defaults to '' for ServiceAccount subjects. Defaults to 'rbac.authorization.k8s.io' for User and Group subjects.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"kind": {
												Description:         "Kind of object being referenced. Values defined by this API group are 'User', 'Group', and 'ServiceAccount'. If the Authorizer does not recognized the kind value, the Authorizer should report an error.",
												MarkdownDescription: "Kind of object being referenced. Values defined by this API group are 'User', 'Group', and 'ServiceAccount'. If the Authorizer does not recognized the kind value, the Authorizer should report an error.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"name": {
												Description:         "Name of the object being referenced.",
												MarkdownDescription: "Name of the object being referenced.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"namespace": {
												Description:         "Namespace of the referenced object.  If the object kind is non-namespace, such as 'User' or 'Group', and this value is not empty the Authorizer should report an error.",
												MarkdownDescription: "Namespace of the referenced object.  If the object kind is non-namespace, such as 'User' or 'Group', and this value is not empty the Authorizer should report an error.",

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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"mutate": {
								Description:         "Mutation is used to modify matching resources.",
								MarkdownDescription: "Mutation is used to modify matching resources.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"foreach": {
										Description:         "ForEach applies mutation rules to a list of sub-elements by creating a context for each entry in the list and looping over it to apply the specified logic.",
										MarkdownDescription: "ForEach applies mutation rules to a list of sub-elements by creating a context for each entry in the list and looping over it to apply the specified logic.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"context": {
												Description:         "Context defines variables and data sources that can be used during rule execution.",
												MarkdownDescription: "Context defines variables and data sources that can be used during rule execution.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"api_call": {
														Description:         "APICall defines an HTTP request to the Kubernetes API server. The JSON data retrieved is stored in the context.",
														MarkdownDescription: "APICall defines an HTTP request to the Kubernetes API server. The JSON data retrieved is stored in the context.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"jmes_path": {
																Description:         "JMESPath is an optional JSON Match Expression that can be used to transform the JSON response returned from the API server. For example a JMESPath of 'items | length(@)' applied to the API server response to the URLPath '/apis/apps/v1/deployments' will return the total count of deployments across all namespaces.",
																MarkdownDescription: "JMESPath is an optional JSON Match Expression that can be used to transform the JSON response returned from the API server. For example a JMESPath of 'items | length(@)' applied to the API server response to the URLPath '/apis/apps/v1/deployments' will return the total count of deployments across all namespaces.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"url_path": {
																Description:         "URLPath is the URL path to be used in the HTTP GET request to the Kubernetes API server (e.g. '/api/v1/namespaces' or  '/apis/apps/v1/deployments'). The format required is the same format used by the 'kubectl get --raw' command.",
																MarkdownDescription: "URLPath is the URL path to be used in the HTTP GET request to the Kubernetes API server (e.g. '/api/v1/namespaces' or  '/apis/apps/v1/deployments'). The format required is the same format used by the 'kubectl get --raw' command.",

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

													"config_map": {
														Description:         "ConfigMap is the ConfigMap reference.",
														MarkdownDescription: "ConfigMap is the ConfigMap reference.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "Name is the ConfigMap name.",
																MarkdownDescription: "Name is the ConfigMap name.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"namespace": {
																Description:         "Namespace is the ConfigMap namespace.",
																MarkdownDescription: "Namespace is the ConfigMap namespace.",

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

													"image_registry": {
														Description:         "ImageRegistry defines requests to an OCI/Docker V2 registry to fetch image details.",
														MarkdownDescription: "ImageRegistry defines requests to an OCI/Docker V2 registry to fetch image details.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"jmes_path": {
																Description:         "JMESPath is an optional JSON Match Expression that can be used to transform the ImageData struct returned as a result of processing the image reference.",
																MarkdownDescription: "JMESPath is an optional JSON Match Expression that can be used to transform the ImageData struct returned as a result of processing the image reference.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"reference": {
																Description:         "Reference is image reference to a container image in the registry. Example: ghcr.io/kyverno/kyverno:latest",
																MarkdownDescription: "Reference is image reference to a container image in the registry. Example: ghcr.io/kyverno/kyverno:latest",

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

													"name": {
														Description:         "Name is the variable name.",
														MarkdownDescription: "Name is the variable name.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"variable": {
														Description:         "Variable defines an arbitrary JMESPath context variable that can be defined inline.",
														MarkdownDescription: "Variable defines an arbitrary JMESPath context variable that can be defined inline.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"default": {
																Description:         "Default is an optional arbitrary JSON object that the variable may take if the JMESPath expression evaluates to nil",
																MarkdownDescription: "Default is an optional arbitrary JSON object that the variable may take if the JMESPath expression evaluates to nil",

																Type: utilities.DynamicType{},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"jmes_path": {
																Description:         "JMESPath is an optional JMESPath Expression that can be used to transform the variable.",
																MarkdownDescription: "JMESPath is an optional JMESPath Expression that can be used to transform the variable.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"value": {
																Description:         "Value is any arbitrary JSON object representable in YAML or JSON form.",
																MarkdownDescription: "Value is any arbitrary JSON object representable in YAML or JSON form.",

																Type: utilities.DynamicType{},

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

											"list": {
												Description:         "List specifies a JMESPath expression that results in one or more elements to which the validation logic is applied.",
												MarkdownDescription: "List specifies a JMESPath expression that results in one or more elements to which the validation logic is applied.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"patch_strategic_merge": {
												Description:         "PatchStrategicMerge is a strategic merge patch used to modify resources. See https://kubernetes.io/docs/tasks/manage-kubernetes-objects/update-api-object-kubectl-patch/ and https://kubectl.docs.kubernetes.io/references/kustomize/patchesstrategicmerge/.",
												MarkdownDescription: "PatchStrategicMerge is a strategic merge patch used to modify resources. See https://kubernetes.io/docs/tasks/manage-kubernetes-objects/update-api-object-kubectl-patch/ and https://kubectl.docs.kubernetes.io/references/kustomize/patchesstrategicmerge/.",

												Type: utilities.DynamicType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"patches_json6902": {
												Description:         "PatchesJSON6902 is a list of RFC 6902 JSON Patch declarations used to modify resources. See https://tools.ietf.org/html/rfc6902 and https://kubectl.docs.kubernetes.io/references/kustomize/patchesjson6902/.",
												MarkdownDescription: "PatchesJSON6902 is a list of RFC 6902 JSON Patch declarations used to modify resources. See https://tools.ietf.org/html/rfc6902 and https://kubectl.docs.kubernetes.io/references/kustomize/patchesjson6902/.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"preconditions": {
												Description:         "AnyAllConditions are used to determine if a policy rule should be applied by evaluating a set of conditions. The declaration can contain nested 'any' or 'all' statements. See: https://kyverno.io/docs/writing-policies/preconditions/",
												MarkdownDescription: "AnyAllConditions are used to determine if a policy rule should be applied by evaluating a set of conditions. The declaration can contain nested 'any' or 'all' statements. See: https://kyverno.io/docs/writing-policies/preconditions/",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"all": {
														Description:         "AllConditions enable variable-based conditional rule execution. This is useful for finer control of when an rule is applied. A condition can reference object data using JMESPath notation. Here, all of the conditions need to pass",
														MarkdownDescription: "AllConditions enable variable-based conditional rule execution. This is useful for finer control of when an rule is applied. A condition can reference object data using JMESPath notation. Here, all of the conditions need to pass",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "Key is the context entry (using JMESPath) for conditional rule evaluation.",
																MarkdownDescription: "Key is the context entry (using JMESPath) for conditional rule evaluation.",

																Type: utilities.DynamicType{},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"operator": {
																Description:         "Operator is the conditional operation to perform. Valid operators are: Equals, NotEquals, In, AnyIn, AllIn, NotIn, AnyNotIn, AllNotIn, GreaterThanOrEquals, GreaterThan, LessThanOrEquals, LessThan, DurationGreaterThanOrEquals, DurationGreaterThan, DurationLessThanOrEquals, DurationLessThan",
																MarkdownDescription: "Operator is the conditional operation to perform. Valid operators are: Equals, NotEquals, In, AnyIn, AllIn, NotIn, AnyNotIn, AllNotIn, GreaterThanOrEquals, GreaterThan, LessThanOrEquals, LessThan, DurationGreaterThanOrEquals, DurationGreaterThan, DurationLessThanOrEquals, DurationLessThan",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,

																Validators: []tfsdk.AttributeValidator{

																	stringvalidator.OneOf("Equals", "NotEquals", "In", "AnyIn", "AllIn", "NotIn", "AnyNotIn", "AllNotIn", "GreaterThanOrEquals", "GreaterThan", "LessThanOrEquals", "LessThan", "DurationGreaterThanOrEquals", "DurationGreaterThan", "DurationLessThanOrEquals", "DurationLessThan"),
																},
															},

															"value": {
																Description:         "Value is the conditional value, or set of values. The values can be fixed set or can be variables declared using JMESPath.",
																MarkdownDescription: "Value is the conditional value, or set of values. The values can be fixed set or can be variables declared using JMESPath.",

																Type: utilities.DynamicType{},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"any": {
														Description:         "AnyConditions enable variable-based conditional rule execution. This is useful for finer control of when an rule is applied. A condition can reference object data using JMESPath notation. Here, at least one of the conditions need to pass",
														MarkdownDescription: "AnyConditions enable variable-based conditional rule execution. This is useful for finer control of when an rule is applied. A condition can reference object data using JMESPath notation. Here, at least one of the conditions need to pass",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "Key is the context entry (using JMESPath) for conditional rule evaluation.",
																MarkdownDescription: "Key is the context entry (using JMESPath) for conditional rule evaluation.",

																Type: utilities.DynamicType{},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"operator": {
																Description:         "Operator is the conditional operation to perform. Valid operators are: Equals, NotEquals, In, AnyIn, AllIn, NotIn, AnyNotIn, AllNotIn, GreaterThanOrEquals, GreaterThan, LessThanOrEquals, LessThan, DurationGreaterThanOrEquals, DurationGreaterThan, DurationLessThanOrEquals, DurationLessThan",
																MarkdownDescription: "Operator is the conditional operation to perform. Valid operators are: Equals, NotEquals, In, AnyIn, AllIn, NotIn, AnyNotIn, AllNotIn, GreaterThanOrEquals, GreaterThan, LessThanOrEquals, LessThan, DurationGreaterThanOrEquals, DurationGreaterThan, DurationLessThanOrEquals, DurationLessThan",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,

																Validators: []tfsdk.AttributeValidator{

																	stringvalidator.OneOf("Equals", "NotEquals", "In", "AnyIn", "AllIn", "NotIn", "AnyNotIn", "AllNotIn", "GreaterThanOrEquals", "GreaterThan", "LessThanOrEquals", "LessThan", "DurationGreaterThanOrEquals", "DurationGreaterThan", "DurationLessThanOrEquals", "DurationLessThan"),
																},
															},

															"value": {
																Description:         "Value is the conditional value, or set of values. The values can be fixed set or can be variables declared using JMESPath.",
																MarkdownDescription: "Value is the conditional value, or set of values. The values can be fixed set or can be variables declared using JMESPath.",

																Type: utilities.DynamicType{},

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

									"patch_strategic_merge": {
										Description:         "PatchStrategicMerge is a strategic merge patch used to modify resources. See https://kubernetes.io/docs/tasks/manage-kubernetes-objects/update-api-object-kubectl-patch/ and https://kubectl.docs.kubernetes.io/references/kustomize/patchesstrategicmerge/.",
										MarkdownDescription: "PatchStrategicMerge is a strategic merge patch used to modify resources. See https://kubernetes.io/docs/tasks/manage-kubernetes-objects/update-api-object-kubectl-patch/ and https://kubectl.docs.kubernetes.io/references/kustomize/patchesstrategicmerge/.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"patches_json6902": {
										Description:         "PatchesJSON6902 is a list of RFC 6902 JSON Patch declarations used to modify resources. See https://tools.ietf.org/html/rfc6902 and https://kubectl.docs.kubernetes.io/references/kustomize/patchesjson6902/.",
										MarkdownDescription: "PatchesJSON6902 is a list of RFC 6902 JSON Patch declarations used to modify resources. See https://tools.ietf.org/html/rfc6902 and https://kubectl.docs.kubernetes.io/references/kustomize/patchesjson6902/.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"targets": {
										Description:         "Targets defines the target resources to be mutated.",
										MarkdownDescription: "Targets defines the target resources to be mutated.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"api_version": {
												Description:         "APIVersion specifies resource apiVersion.",
												MarkdownDescription: "APIVersion specifies resource apiVersion.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"kind": {
												Description:         "Kind specifies resource kind.",
												MarkdownDescription: "Kind specifies resource kind.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "Name specifies the resource name.",
												MarkdownDescription: "Name specifies the resource name.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"namespace": {
												Description:         "Namespace specifies resource namespace.",
												MarkdownDescription: "Namespace specifies resource namespace.",

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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "Name is a label to identify the rule, It must be unique within the policy.",
								MarkdownDescription: "Name is a label to identify the rule, It must be unique within the policy.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.LengthAtMost(63),
								},
							},

							"preconditions": {
								Description:         "Preconditions are used to determine if a policy rule should be applied by evaluating a set of conditions. The declaration can contain nested 'any' or 'all' statements. A direct list of conditions (without 'any' or 'all' statements is supported for backwards compatibility but will be deprecated in the next major release. See: https://kyverno.io/docs/writing-policies/preconditions/",
								MarkdownDescription: "Preconditions are used to determine if a policy rule should be applied by evaluating a set of conditions. The declaration can contain nested 'any' or 'all' statements. A direct list of conditions (without 'any' or 'all' statements is supported for backwards compatibility but will be deprecated in the next major release. See: https://kyverno.io/docs/writing-policies/preconditions/",

								Type: utilities.DynamicType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"validate": {
								Description:         "Validation is used to validate matching resources.",
								MarkdownDescription: "Validation is used to validate matching resources.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"any_pattern": {
										Description:         "AnyPattern specifies list of validation patterns. At least one of the patterns must be satisfied for the validation rule to succeed.",
										MarkdownDescription: "AnyPattern specifies list of validation patterns. At least one of the patterns must be satisfied for the validation rule to succeed.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"deny": {
										Description:         "Deny defines conditions used to pass or fail a validation rule.",
										MarkdownDescription: "Deny defines conditions used to pass or fail a validation rule.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"conditions": {
												Description:         "Multiple conditions can be declared under an 'any' or 'all' statement. A direct list of conditions (without 'any' or 'all' statements) is also supported for backwards compatibility but will be deprecated in the next major release. See: https://kyverno.io/docs/writing-policies/validate/#deny-rules",
												MarkdownDescription: "Multiple conditions can be declared under an 'any' or 'all' statement. A direct list of conditions (without 'any' or 'all' statements) is also supported for backwards compatibility but will be deprecated in the next major release. See: https://kyverno.io/docs/writing-policies/validate/#deny-rules",

												Type: utilities.DynamicType{},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"foreach": {
										Description:         "ForEach applies validate rules to a list of sub-elements by creating a context for each entry in the list and looping over it to apply the specified logic.",
										MarkdownDescription: "ForEach applies validate rules to a list of sub-elements by creating a context for each entry in the list and looping over it to apply the specified logic.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"any_pattern": {
												Description:         "AnyPattern specifies list of validation patterns. At least one of the patterns must be satisfied for the validation rule to succeed.",
												MarkdownDescription: "AnyPattern specifies list of validation patterns. At least one of the patterns must be satisfied for the validation rule to succeed.",

												Type: utilities.DynamicType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"context": {
												Description:         "Context defines variables and data sources that can be used during rule execution.",
												MarkdownDescription: "Context defines variables and data sources that can be used during rule execution.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"api_call": {
														Description:         "APICall defines an HTTP request to the Kubernetes API server. The JSON data retrieved is stored in the context.",
														MarkdownDescription: "APICall defines an HTTP request to the Kubernetes API server. The JSON data retrieved is stored in the context.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"jmes_path": {
																Description:         "JMESPath is an optional JSON Match Expression that can be used to transform the JSON response returned from the API server. For example a JMESPath of 'items | length(@)' applied to the API server response to the URLPath '/apis/apps/v1/deployments' will return the total count of deployments across all namespaces.",
																MarkdownDescription: "JMESPath is an optional JSON Match Expression that can be used to transform the JSON response returned from the API server. For example a JMESPath of 'items | length(@)' applied to the API server response to the URLPath '/apis/apps/v1/deployments' will return the total count of deployments across all namespaces.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"url_path": {
																Description:         "URLPath is the URL path to be used in the HTTP GET request to the Kubernetes API server (e.g. '/api/v1/namespaces' or  '/apis/apps/v1/deployments'). The format required is the same format used by the 'kubectl get --raw' command.",
																MarkdownDescription: "URLPath is the URL path to be used in the HTTP GET request to the Kubernetes API server (e.g. '/api/v1/namespaces' or  '/apis/apps/v1/deployments'). The format required is the same format used by the 'kubectl get --raw' command.",

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

													"config_map": {
														Description:         "ConfigMap is the ConfigMap reference.",
														MarkdownDescription: "ConfigMap is the ConfigMap reference.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "Name is the ConfigMap name.",
																MarkdownDescription: "Name is the ConfigMap name.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"namespace": {
																Description:         "Namespace is the ConfigMap namespace.",
																MarkdownDescription: "Namespace is the ConfigMap namespace.",

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

													"image_registry": {
														Description:         "ImageRegistry defines requests to an OCI/Docker V2 registry to fetch image details.",
														MarkdownDescription: "ImageRegistry defines requests to an OCI/Docker V2 registry to fetch image details.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"jmes_path": {
																Description:         "JMESPath is an optional JSON Match Expression that can be used to transform the ImageData struct returned as a result of processing the image reference.",
																MarkdownDescription: "JMESPath is an optional JSON Match Expression that can be used to transform the ImageData struct returned as a result of processing the image reference.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"reference": {
																Description:         "Reference is image reference to a container image in the registry. Example: ghcr.io/kyverno/kyverno:latest",
																MarkdownDescription: "Reference is image reference to a container image in the registry. Example: ghcr.io/kyverno/kyverno:latest",

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

													"name": {
														Description:         "Name is the variable name.",
														MarkdownDescription: "Name is the variable name.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"variable": {
														Description:         "Variable defines an arbitrary JMESPath context variable that can be defined inline.",
														MarkdownDescription: "Variable defines an arbitrary JMESPath context variable that can be defined inline.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"default": {
																Description:         "Default is an optional arbitrary JSON object that the variable may take if the JMESPath expression evaluates to nil",
																MarkdownDescription: "Default is an optional arbitrary JSON object that the variable may take if the JMESPath expression evaluates to nil",

																Type: utilities.DynamicType{},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"jmes_path": {
																Description:         "JMESPath is an optional JMESPath Expression that can be used to transform the variable.",
																MarkdownDescription: "JMESPath is an optional JMESPath Expression that can be used to transform the variable.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"value": {
																Description:         "Value is any arbitrary JSON object representable in YAML or JSON form.",
																MarkdownDescription: "Value is any arbitrary JSON object representable in YAML or JSON form.",

																Type: utilities.DynamicType{},

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

											"deny": {
												Description:         "Deny defines conditions used to pass or fail a validation rule.",
												MarkdownDescription: "Deny defines conditions used to pass or fail a validation rule.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"conditions": {
														Description:         "Multiple conditions can be declared under an 'any' or 'all' statement. A direct list of conditions (without 'any' or 'all' statements) is also supported for backwards compatibility but will be deprecated in the next major release. See: https://kyverno.io/docs/writing-policies/validate/#deny-rules",
														MarkdownDescription: "Multiple conditions can be declared under an 'any' or 'all' statement. A direct list of conditions (without 'any' or 'all' statements) is also supported for backwards compatibility but will be deprecated in the next major release. See: https://kyverno.io/docs/writing-policies/validate/#deny-rules",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"element_scope": {
												Description:         "ElementScope specifies whether to use the current list element as the scope for validation. Defaults to 'true' if not specified. When set to 'false', 'request.object' is used as the validation scope within the foreach block to allow referencing other elements in the subtree.",
												MarkdownDescription: "ElementScope specifies whether to use the current list element as the scope for validation. Defaults to 'true' if not specified. When set to 'false', 'request.object' is used as the validation scope within the foreach block to allow referencing other elements in the subtree.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"list": {
												Description:         "List specifies a JMESPath expression that results in one or more elements to which the validation logic is applied.",
												MarkdownDescription: "List specifies a JMESPath expression that results in one or more elements to which the validation logic is applied.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"pattern": {
												Description:         "Pattern specifies an overlay-style pattern used to check resources.",
												MarkdownDescription: "Pattern specifies an overlay-style pattern used to check resources.",

												Type: utilities.DynamicType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"preconditions": {
												Description:         "AnyAllConditions are used to determine if a policy rule should be applied by evaluating a set of conditions. The declaration can contain nested 'any' or 'all' statements. See: https://kyverno.io/docs/writing-policies/preconditions/",
												MarkdownDescription: "AnyAllConditions are used to determine if a policy rule should be applied by evaluating a set of conditions. The declaration can contain nested 'any' or 'all' statements. See: https://kyverno.io/docs/writing-policies/preconditions/",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"all": {
														Description:         "AllConditions enable variable-based conditional rule execution. This is useful for finer control of when an rule is applied. A condition can reference object data using JMESPath notation. Here, all of the conditions need to pass",
														MarkdownDescription: "AllConditions enable variable-based conditional rule execution. This is useful for finer control of when an rule is applied. A condition can reference object data using JMESPath notation. Here, all of the conditions need to pass",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "Key is the context entry (using JMESPath) for conditional rule evaluation.",
																MarkdownDescription: "Key is the context entry (using JMESPath) for conditional rule evaluation.",

																Type: utilities.DynamicType{},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"operator": {
																Description:         "Operator is the conditional operation to perform. Valid operators are: Equals, NotEquals, In, AnyIn, AllIn, NotIn, AnyNotIn, AllNotIn, GreaterThanOrEquals, GreaterThan, LessThanOrEquals, LessThan, DurationGreaterThanOrEquals, DurationGreaterThan, DurationLessThanOrEquals, DurationLessThan",
																MarkdownDescription: "Operator is the conditional operation to perform. Valid operators are: Equals, NotEquals, In, AnyIn, AllIn, NotIn, AnyNotIn, AllNotIn, GreaterThanOrEquals, GreaterThan, LessThanOrEquals, LessThan, DurationGreaterThanOrEquals, DurationGreaterThan, DurationLessThanOrEquals, DurationLessThan",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,

																Validators: []tfsdk.AttributeValidator{

																	stringvalidator.OneOf("Equals", "NotEquals", "In", "AnyIn", "AllIn", "NotIn", "AnyNotIn", "AllNotIn", "GreaterThanOrEquals", "GreaterThan", "LessThanOrEquals", "LessThan", "DurationGreaterThanOrEquals", "DurationGreaterThan", "DurationLessThanOrEquals", "DurationLessThan"),
																},
															},

															"value": {
																Description:         "Value is the conditional value, or set of values. The values can be fixed set or can be variables declared using JMESPath.",
																MarkdownDescription: "Value is the conditional value, or set of values. The values can be fixed set or can be variables declared using JMESPath.",

																Type: utilities.DynamicType{},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"any": {
														Description:         "AnyConditions enable variable-based conditional rule execution. This is useful for finer control of when an rule is applied. A condition can reference object data using JMESPath notation. Here, at least one of the conditions need to pass",
														MarkdownDescription: "AnyConditions enable variable-based conditional rule execution. This is useful for finer control of when an rule is applied. A condition can reference object data using JMESPath notation. Here, at least one of the conditions need to pass",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "Key is the context entry (using JMESPath) for conditional rule evaluation.",
																MarkdownDescription: "Key is the context entry (using JMESPath) for conditional rule evaluation.",

																Type: utilities.DynamicType{},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"operator": {
																Description:         "Operator is the conditional operation to perform. Valid operators are: Equals, NotEquals, In, AnyIn, AllIn, NotIn, AnyNotIn, AllNotIn, GreaterThanOrEquals, GreaterThan, LessThanOrEquals, LessThan, DurationGreaterThanOrEquals, DurationGreaterThan, DurationLessThanOrEquals, DurationLessThan",
																MarkdownDescription: "Operator is the conditional operation to perform. Valid operators are: Equals, NotEquals, In, AnyIn, AllIn, NotIn, AnyNotIn, AllNotIn, GreaterThanOrEquals, GreaterThan, LessThanOrEquals, LessThan, DurationGreaterThanOrEquals, DurationGreaterThan, DurationLessThanOrEquals, DurationLessThan",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,

																Validators: []tfsdk.AttributeValidator{

																	stringvalidator.OneOf("Equals", "NotEquals", "In", "AnyIn", "AllIn", "NotIn", "AnyNotIn", "AllNotIn", "GreaterThanOrEquals", "GreaterThan", "LessThanOrEquals", "LessThan", "DurationGreaterThanOrEquals", "DurationGreaterThan", "DurationLessThanOrEquals", "DurationLessThan"),
																},
															},

															"value": {
																Description:         "Value is the conditional value, or set of values. The values can be fixed set or can be variables declared using JMESPath.",
																MarkdownDescription: "Value is the conditional value, or set of values. The values can be fixed set or can be variables declared using JMESPath.",

																Type: utilities.DynamicType{},

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

									"manifests": {
										Description:         "Manifest specifies conditions for manifest verification",
										MarkdownDescription: "Manifest specifies conditions for manifest verification",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"annotation_domain": {
												Description:         "AnnotationDomain is custom domain of annotation for message and signature. Default is 'cosign.sigstore.dev'.",
												MarkdownDescription: "AnnotationDomain is custom domain of annotation for message and signature. Default is 'cosign.sigstore.dev'.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"attestors": {
												Description:         "Attestors specified the required attestors (i.e. authorities)",
												MarkdownDescription: "Attestors specified the required attestors (i.e. authorities)",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"count": {
														Description:         "Count specifies the required number of entries that must match. If the count is null, all entries must match (a logical AND). If the count is 1, at least one entry must match (a logical OR). If the count contains a value N, then N must be less than or equal to the size of entries, and at least N entries must match.",
														MarkdownDescription: "Count specifies the required number of entries that must match. If the count is null, all entries must match (a logical AND). If the count is 1, at least one entry must match (a logical OR). If the count contains a value N, then N must be less than or equal to the size of entries, and at least N entries must match.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															int64validator.AtLeast(1),
														},
													},

													"entries": {
														Description:         "Entries contains the available attestors. An attestor can be a static key, attributes for keyless verification, or a nested attestor declaration.",
														MarkdownDescription: "Entries contains the available attestors. An attestor can be a static key, attributes for keyless verification, or a nested attestor declaration.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"annotations": {
																Description:         "Annotations are used for image verification. Every specified key-value pair must exist and match in the verified payload. The payload may contain other key-value pairs.",
																MarkdownDescription: "Annotations are used for image verification. Every specified key-value pair must exist and match in the verified payload. The payload may contain other key-value pairs.",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"attestor": {
																Description:         "Attestor is a nested AttestorSet used to specify a more complex set of match authorities",
																MarkdownDescription: "Attestor is a nested AttestorSet used to specify a more complex set of match authorities",

																Type: utilities.DynamicType{},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"certificates": {
																Description:         "Certificates specifies one or more certificates",
																MarkdownDescription: "Certificates specifies one or more certificates",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"cert": {
																		Description:         "Certificate is an optional PEM encoded public certificate.",
																		MarkdownDescription: "Certificate is an optional PEM encoded public certificate.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"cert_chain": {
																		Description:         "CertificateChain is an optional PEM encoded set of certificates used to verify",
																		MarkdownDescription: "CertificateChain is an optional PEM encoded set of certificates used to verify",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"rekor": {
																		Description:         "Rekor provides configuration for the Rekor transparency log service. If the value is nil, Rekor is not checked. If an empty object is provided the public instance of Rekor (https://rekor.sigstore.dev) is used.",
																		MarkdownDescription: "Rekor provides configuration for the Rekor transparency log service. If the value is nil, Rekor is not checked. If an empty object is provided the public instance of Rekor (https://rekor.sigstore.dev) is used.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"url": {
																				Description:         "URL is the address of the transparency log. Defaults to the public log https://rekor.sigstore.dev.",
																				MarkdownDescription: "URL is the address of the transparency log. Defaults to the public log https://rekor.sigstore.dev.",

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

															"keyless": {
																Description:         "Keyless is a set of attribute used to verify a Sigstore keyless attestor. See https://github.com/sigstore/cosign/blob/main/KEYLESS.md.",
																MarkdownDescription: "Keyless is a set of attribute used to verify a Sigstore keyless attestor. See https://github.com/sigstore/cosign/blob/main/KEYLESS.md.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"additional_extensions": {
																		Description:         "AdditionalExtensions are certificate-extensions used for keyless signing.",
																		MarkdownDescription: "AdditionalExtensions are certificate-extensions used for keyless signing.",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"issuer": {
																		Description:         "Issuer is the certificate issuer used for keyless signing.",
																		MarkdownDescription: "Issuer is the certificate issuer used for keyless signing.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"rekor": {
																		Description:         "Rekor provides configuration for the Rekor transparency log service. If the value is nil, Rekor is not checked and a root certificate chain is expected instead. If an empty object is provided the public instance of Rekor (https://rekor.sigstore.dev) is used.",
																		MarkdownDescription: "Rekor provides configuration for the Rekor transparency log service. If the value is nil, Rekor is not checked and a root certificate chain is expected instead. If an empty object is provided the public instance of Rekor (https://rekor.sigstore.dev) is used.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"url": {
																				Description:         "URL is the address of the transparency log. Defaults to the public log https://rekor.sigstore.dev.",
																				MarkdownDescription: "URL is the address of the transparency log. Defaults to the public log https://rekor.sigstore.dev.",

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

																	"roots": {
																		Description:         "Roots is an optional set of PEM encoded trusted root certificates. If not provided, the system roots are used.",
																		MarkdownDescription: "Roots is an optional set of PEM encoded trusted root certificates. If not provided, the system roots are used.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"subject": {
																		Description:         "Subject is the verified identity used for keyless signing, for example the email address",
																		MarkdownDescription: "Subject is the verified identity used for keyless signing, for example the email address",

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

															"keys": {
																Description:         "Keys specifies one or more public keys",
																MarkdownDescription: "Keys specifies one or more public keys",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"kms": {
																		Description:         "KMS provides the URI to the public key stored in a Key Management System. See: https://github.com/sigstore/cosign/blob/main/KMS.md",
																		MarkdownDescription: "KMS provides the URI to the public key stored in a Key Management System. See: https://github.com/sigstore/cosign/blob/main/KMS.md",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"public_keys": {
																		Description:         "Keys is a set of X.509 public keys used to verify image signatures. The keys can be directly specified or can be a variable reference to a key specified in a ConfigMap (see https://kyverno.io/docs/writing-policies/variables/). When multiple keys are specified each key is processed as a separate staticKey entry (.attestors[*].entries.keys) within the set of attestors and the count is applied across the keys.",
																		MarkdownDescription: "Keys is a set of X.509 public keys used to verify image signatures. The keys can be directly specified or can be a variable reference to a key specified in a ConfigMap (see https://kyverno.io/docs/writing-policies/variables/). When multiple keys are specified each key is processed as a separate staticKey entry (.attestors[*].entries.keys) within the set of attestors and the count is applied across the keys.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"rekor": {
																		Description:         "Rekor provides configuration for the Rekor transparency log service. If the value is nil, Rekor is not checked. If an empty object is provided the public instance of Rekor (https://rekor.sigstore.dev) is used.",
																		MarkdownDescription: "Rekor provides configuration for the Rekor transparency log service. If the value is nil, Rekor is not checked. If an empty object is provided the public instance of Rekor (https://rekor.sigstore.dev) is used.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"url": {
																				Description:         "URL is the address of the transparency log. Defaults to the public log https://rekor.sigstore.dev.",
																				MarkdownDescription: "URL is the address of the transparency log. Defaults to the public log https://rekor.sigstore.dev.",

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

																	"secret": {
																		Description:         "Reference to a Secret resource that contains a public key",
																		MarkdownDescription: "Reference to a Secret resource that contains a public key",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "name of the secret",
																				MarkdownDescription: "name of the secret",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"namespace": {
																				Description:         "namespace name in which secret is created",
																				MarkdownDescription: "namespace name in which secret is created",

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

																	"signature_algorithm": {
																		Description:         "Specify signature algorithm for public keys. Supported values are sha256 and sha512",
																		MarkdownDescription: "Specify signature algorithm for public keys. Supported values are sha256 and sha512",

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

															"repository": {
																Description:         "Repository is an optional alternate OCI repository to use for signatures and attestations that match this rule. If specified Repository will override other OCI image repository locations for this Attestor.",
																MarkdownDescription: "Repository is an optional alternate OCI repository to use for signatures and attestations that match this rule. If specified Repository will override other OCI image repository locations for this Attestor.",

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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"dry_run": {
												Description:         "DryRun configuration",
												MarkdownDescription: "DryRun configuration",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"enable": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespace": {
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

											"ignore_fields": {
												Description:         "Fields which will be ignored while comparing manifests.",
												MarkdownDescription: "Fields which will be ignored while comparing manifests.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"fields": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"objects": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"group": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"kind": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"namespace": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"version": {
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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"repository": {
												Description:         "Repository is an optional alternate OCI repository to use for resource bundle reference. The repository can be overridden per Attestor or Attestation.",
												MarkdownDescription: "Repository is an optional alternate OCI repository to use for resource bundle reference. The repository can be overridden per Attestor or Attestation.",

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

									"message": {
										Description:         "Message specifies a custom message to be displayed on failure.",
										MarkdownDescription: "Message specifies a custom message to be displayed on failure.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"pattern": {
										Description:         "Pattern specifies an overlay-style pattern used to check resources.",
										MarkdownDescription: "Pattern specifies an overlay-style pattern used to check resources.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"pod_security": {
										Description:         "PodSecurity applies exemptions for Kubernetes Pod Security admission by specifying exclusions for Pod Security Standards controls.",
										MarkdownDescription: "PodSecurity applies exemptions for Kubernetes Pod Security admission by specifying exclusions for Pod Security Standards controls.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"exclude": {
												Description:         "Exclude specifies the Pod Security Standard controls to be excluded.",
												MarkdownDescription: "Exclude specifies the Pod Security Standard controls to be excluded.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"control_name": {
														Description:         "ControlName specifies the name of the Pod Security Standard control. See: https://kubernetes.io/docs/concepts/security/pod-security-standards/",
														MarkdownDescription: "ControlName specifies the name of the Pod Security Standard control. See: https://kubernetes.io/docs/concepts/security/pod-security-standards/",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("HostProcess", "Host Namespaces", "Privileged Containers", "Capabilities", "HostPath Volumes", "Host Ports", "AppArmor", "SELinux", "/proc Mount Type", "Seccomp", "Sysctls", "Volume Types", "Privilege Escalation", "Running as Non-root", "Running as Non-root user"),
														},
													},

													"images": {
														Description:         "Images selects matching containers and applies the container level PSS. Each image is the image name consisting of the registry address, repository, image, and tag. Empty list matches no containers, PSS checks are applied at the pod level only. Wildcards ('*' and '?') are allowed. See: https://kubernetes.io/docs/concepts/containers/images.",
														MarkdownDescription: "Images selects matching containers and applies the container level PSS. Each image is the image name consisting of the registry address, repository, image, and tag. Empty list matches no containers, PSS checks are applied at the pod level only. Wildcards ('*' and '?') are allowed. See: https://kubernetes.io/docs/concepts/containers/images.",

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

											"level": {
												Description:         "Level defines the Pod Security Standard level to be applied to workloads. Allowed values are privileged, baseline, and restricted.",
												MarkdownDescription: "Level defines the Pod Security Standard level to be applied to workloads. Allowed values are privileged, baseline, and restricted.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("privileged", "baseline", "restricted"),
												},
											},

											"version": {
												Description:         "Version defines the Pod Security Standard versions that Kubernetes supports. Allowed values are v1.19, v1.20, v1.21, v1.22, v1.23, v1.24, v1.25, latest. Defaults to latest.",
												MarkdownDescription: "Version defines the Pod Security Standard versions that Kubernetes supports. Allowed values are v1.19, v1.20, v1.21, v1.22, v1.23, v1.24, v1.25, latest. Defaults to latest.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("v1.19", "v1.20", "v1.21", "v1.22", "v1.23", "v1.24", "v1.25", "latest"),
												},
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

							"verify_images": {
								Description:         "VerifyImages is used to verify image signatures and mutate them to add a digest",
								MarkdownDescription: "VerifyImages is used to verify image signatures and mutate them to add a digest",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"additional_extensions": {
										Description:         "AdditionalExtensions are certificate-extensions used for keyless signing. Deprecated.",
										MarkdownDescription: "AdditionalExtensions are certificate-extensions used for keyless signing. Deprecated.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"annotations": {
										Description:         "Annotations are used for image verification. Every specified key-value pair must exist and match in the verified payload. The payload may contain other key-value pairs. Deprecated. Use annotations per Attestor instead.",
										MarkdownDescription: "Annotations are used for image verification. Every specified key-value pair must exist and match in the verified payload. The payload may contain other key-value pairs. Deprecated. Use annotations per Attestor instead.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"attestations": {
										Description:         "Attestations are optional checks for signed in-toto Statements used to verify the image. See https://github.com/in-toto/attestation. Kyverno fetches signed attestations from the OCI registry and decodes them into a list of Statement declarations.",
										MarkdownDescription: "Attestations are optional checks for signed in-toto Statements used to verify the image. See https://github.com/in-toto/attestation. Kyverno fetches signed attestations from the OCI registry and decodes them into a list of Statement declarations.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"conditions": {
												Description:         "Conditions are used to verify attributes within a Predicate. If no Conditions are specified the attestation check is satisfied as long there are predicates that match the predicate type.",
												MarkdownDescription: "Conditions are used to verify attributes within a Predicate. If no Conditions are specified the attestation check is satisfied as long there are predicates that match the predicate type.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"all": {
														Description:         "AllConditions enable variable-based conditional rule execution. This is useful for finer control of when an rule is applied. A condition can reference object data using JMESPath notation. Here, all of the conditions need to pass",
														MarkdownDescription: "AllConditions enable variable-based conditional rule execution. This is useful for finer control of when an rule is applied. A condition can reference object data using JMESPath notation. Here, all of the conditions need to pass",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "Key is the context entry (using JMESPath) for conditional rule evaluation.",
																MarkdownDescription: "Key is the context entry (using JMESPath) for conditional rule evaluation.",

																Type: utilities.DynamicType{},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"operator": {
																Description:         "Operator is the conditional operation to perform. Valid operators are: Equals, NotEquals, In, AnyIn, AllIn, NotIn, AnyNotIn, AllNotIn, GreaterThanOrEquals, GreaterThan, LessThanOrEquals, LessThan, DurationGreaterThanOrEquals, DurationGreaterThan, DurationLessThanOrEquals, DurationLessThan",
																MarkdownDescription: "Operator is the conditional operation to perform. Valid operators are: Equals, NotEquals, In, AnyIn, AllIn, NotIn, AnyNotIn, AllNotIn, GreaterThanOrEquals, GreaterThan, LessThanOrEquals, LessThan, DurationGreaterThanOrEquals, DurationGreaterThan, DurationLessThanOrEquals, DurationLessThan",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,

																Validators: []tfsdk.AttributeValidator{

																	stringvalidator.OneOf("Equals", "NotEquals", "In", "AnyIn", "AllIn", "NotIn", "AnyNotIn", "AllNotIn", "GreaterThanOrEquals", "GreaterThan", "LessThanOrEquals", "LessThan", "DurationGreaterThanOrEquals", "DurationGreaterThan", "DurationLessThanOrEquals", "DurationLessThan"),
																},
															},

															"value": {
																Description:         "Value is the conditional value, or set of values. The values can be fixed set or can be variables declared using JMESPath.",
																MarkdownDescription: "Value is the conditional value, or set of values. The values can be fixed set or can be variables declared using JMESPath.",

																Type: utilities.DynamicType{},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"any": {
														Description:         "AnyConditions enable variable-based conditional rule execution. This is useful for finer control of when an rule is applied. A condition can reference object data using JMESPath notation. Here, at least one of the conditions need to pass",
														MarkdownDescription: "AnyConditions enable variable-based conditional rule execution. This is useful for finer control of when an rule is applied. A condition can reference object data using JMESPath notation. Here, at least one of the conditions need to pass",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "Key is the context entry (using JMESPath) for conditional rule evaluation.",
																MarkdownDescription: "Key is the context entry (using JMESPath) for conditional rule evaluation.",

																Type: utilities.DynamicType{},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"operator": {
																Description:         "Operator is the conditional operation to perform. Valid operators are: Equals, NotEquals, In, AnyIn, AllIn, NotIn, AnyNotIn, AllNotIn, GreaterThanOrEquals, GreaterThan, LessThanOrEquals, LessThan, DurationGreaterThanOrEquals, DurationGreaterThan, DurationLessThanOrEquals, DurationLessThan",
																MarkdownDescription: "Operator is the conditional operation to perform. Valid operators are: Equals, NotEquals, In, AnyIn, AllIn, NotIn, AnyNotIn, AllNotIn, GreaterThanOrEquals, GreaterThan, LessThanOrEquals, LessThan, DurationGreaterThanOrEquals, DurationGreaterThan, DurationLessThanOrEquals, DurationLessThan",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,

																Validators: []tfsdk.AttributeValidator{

																	stringvalidator.OneOf("Equals", "NotEquals", "In", "AnyIn", "AllIn", "NotIn", "AnyNotIn", "AllNotIn", "GreaterThanOrEquals", "GreaterThan", "LessThanOrEquals", "LessThan", "DurationGreaterThanOrEquals", "DurationGreaterThan", "DurationLessThanOrEquals", "DurationLessThan"),
																},
															},

															"value": {
																Description:         "Value is the conditional value, or set of values. The values can be fixed set or can be variables declared using JMESPath.",
																MarkdownDescription: "Value is the conditional value, or set of values. The values can be fixed set or can be variables declared using JMESPath.",

																Type: utilities.DynamicType{},

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

											"predicate_type": {
												Description:         "PredicateType defines the type of Predicate contained within the Statement.",
												MarkdownDescription: "PredicateType defines the type of Predicate contained within the Statement.",

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

									"attestors": {
										Description:         "Attestors specified the required attestors (i.e. authorities)",
										MarkdownDescription: "Attestors specified the required attestors (i.e. authorities)",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"count": {
												Description:         "Count specifies the required number of entries that must match. If the count is null, all entries must match (a logical AND). If the count is 1, at least one entry must match (a logical OR). If the count contains a value N, then N must be less than or equal to the size of entries, and at least N entries must match.",
												MarkdownDescription: "Count specifies the required number of entries that must match. If the count is null, all entries must match (a logical AND). If the count is 1, at least one entry must match (a logical OR). If the count contains a value N, then N must be less than or equal to the size of entries, and at least N entries must match.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},

											"entries": {
												Description:         "Entries contains the available attestors. An attestor can be a static key, attributes for keyless verification, or a nested attestor declaration.",
												MarkdownDescription: "Entries contains the available attestors. An attestor can be a static key, attributes for keyless verification, or a nested attestor declaration.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations are used for image verification. Every specified key-value pair must exist and match in the verified payload. The payload may contain other key-value pairs.",
														MarkdownDescription: "Annotations are used for image verification. Every specified key-value pair must exist and match in the verified payload. The payload may contain other key-value pairs.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"attestor": {
														Description:         "Attestor is a nested AttestorSet used to specify a more complex set of match authorities",
														MarkdownDescription: "Attestor is a nested AttestorSet used to specify a more complex set of match authorities",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"certificates": {
														Description:         "Certificates specifies one or more certificates",
														MarkdownDescription: "Certificates specifies one or more certificates",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"cert": {
																Description:         "Certificate is an optional PEM encoded public certificate.",
																MarkdownDescription: "Certificate is an optional PEM encoded public certificate.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"cert_chain": {
																Description:         "CertificateChain is an optional PEM encoded set of certificates used to verify",
																MarkdownDescription: "CertificateChain is an optional PEM encoded set of certificates used to verify",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"rekor": {
																Description:         "Rekor provides configuration for the Rekor transparency log service. If the value is nil, Rekor is not checked. If an empty object is provided the public instance of Rekor (https://rekor.sigstore.dev) is used.",
																MarkdownDescription: "Rekor provides configuration for the Rekor transparency log service. If the value is nil, Rekor is not checked. If an empty object is provided the public instance of Rekor (https://rekor.sigstore.dev) is used.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"url": {
																		Description:         "URL is the address of the transparency log. Defaults to the public log https://rekor.sigstore.dev.",
																		MarkdownDescription: "URL is the address of the transparency log. Defaults to the public log https://rekor.sigstore.dev.",

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

													"keyless": {
														Description:         "Keyless is a set of attribute used to verify a Sigstore keyless attestor. See https://github.com/sigstore/cosign/blob/main/KEYLESS.md.",
														MarkdownDescription: "Keyless is a set of attribute used to verify a Sigstore keyless attestor. See https://github.com/sigstore/cosign/blob/main/KEYLESS.md.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"additional_extensions": {
																Description:         "AdditionalExtensions are certificate-extensions used for keyless signing.",
																MarkdownDescription: "AdditionalExtensions are certificate-extensions used for keyless signing.",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"issuer": {
																Description:         "Issuer is the certificate issuer used for keyless signing.",
																MarkdownDescription: "Issuer is the certificate issuer used for keyless signing.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"rekor": {
																Description:         "Rekor provides configuration for the Rekor transparency log service. If the value is nil, Rekor is not checked and a root certificate chain is expected instead. If an empty object is provided the public instance of Rekor (https://rekor.sigstore.dev) is used.",
																MarkdownDescription: "Rekor provides configuration for the Rekor transparency log service. If the value is nil, Rekor is not checked and a root certificate chain is expected instead. If an empty object is provided the public instance of Rekor (https://rekor.sigstore.dev) is used.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"url": {
																		Description:         "URL is the address of the transparency log. Defaults to the public log https://rekor.sigstore.dev.",
																		MarkdownDescription: "URL is the address of the transparency log. Defaults to the public log https://rekor.sigstore.dev.",

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

															"roots": {
																Description:         "Roots is an optional set of PEM encoded trusted root certificates. If not provided, the system roots are used.",
																MarkdownDescription: "Roots is an optional set of PEM encoded trusted root certificates. If not provided, the system roots are used.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"subject": {
																Description:         "Subject is the verified identity used for keyless signing, for example the email address",
																MarkdownDescription: "Subject is the verified identity used for keyless signing, for example the email address",

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

													"keys": {
														Description:         "Keys specifies one or more public keys",
														MarkdownDescription: "Keys specifies one or more public keys",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"kms": {
																Description:         "KMS provides the URI to the public key stored in a Key Management System. See: https://github.com/sigstore/cosign/blob/main/KMS.md",
																MarkdownDescription: "KMS provides the URI to the public key stored in a Key Management System. See: https://github.com/sigstore/cosign/blob/main/KMS.md",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"public_keys": {
																Description:         "Keys is a set of X.509 public keys used to verify image signatures. The keys can be directly specified or can be a variable reference to a key specified in a ConfigMap (see https://kyverno.io/docs/writing-policies/variables/). When multiple keys are specified each key is processed as a separate staticKey entry (.attestors[*].entries.keys) within the set of attestors and the count is applied across the keys.",
																MarkdownDescription: "Keys is a set of X.509 public keys used to verify image signatures. The keys can be directly specified or can be a variable reference to a key specified in a ConfigMap (see https://kyverno.io/docs/writing-policies/variables/). When multiple keys are specified each key is processed as a separate staticKey entry (.attestors[*].entries.keys) within the set of attestors and the count is applied across the keys.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"rekor": {
																Description:         "Rekor provides configuration for the Rekor transparency log service. If the value is nil, Rekor is not checked. If an empty object is provided the public instance of Rekor (https://rekor.sigstore.dev) is used.",
																MarkdownDescription: "Rekor provides configuration for the Rekor transparency log service. If the value is nil, Rekor is not checked. If an empty object is provided the public instance of Rekor (https://rekor.sigstore.dev) is used.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"url": {
																		Description:         "URL is the address of the transparency log. Defaults to the public log https://rekor.sigstore.dev.",
																		MarkdownDescription: "URL is the address of the transparency log. Defaults to the public log https://rekor.sigstore.dev.",

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

															"secret": {
																Description:         "Reference to a Secret resource that contains a public key",
																MarkdownDescription: "Reference to a Secret resource that contains a public key",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "name of the secret",
																		MarkdownDescription: "name of the secret",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"namespace": {
																		Description:         "namespace name in which secret is created",
																		MarkdownDescription: "namespace name in which secret is created",

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

															"signature_algorithm": {
																Description:         "Specify signature algorithm for public keys. Supported values are sha256 and sha512",
																MarkdownDescription: "Specify signature algorithm for public keys. Supported values are sha256 and sha512",

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

													"repository": {
														Description:         "Repository is an optional alternate OCI repository to use for signatures and attestations that match this rule. If specified Repository will override other OCI image repository locations for this Attestor.",
														MarkdownDescription: "Repository is an optional alternate OCI repository to use for signatures and attestations that match this rule. If specified Repository will override other OCI image repository locations for this Attestor.",

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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"image": {
										Description:         "Image is the image name consisting of the registry address, repository, image, and tag. Wildcards ('*' and '?') are allowed. See: https://kubernetes.io/docs/concepts/containers/images. Deprecated. Use ImageReferences instead.",
										MarkdownDescription: "Image is the image name consisting of the registry address, repository, image, and tag. Wildcards ('*' and '?') are allowed. See: https://kubernetes.io/docs/concepts/containers/images. Deprecated. Use ImageReferences instead.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"image_references": {
										Description:         "ImageReferences is a list of matching image reference patterns. At least one pattern in the list must match the image for the rule to apply. Each image reference consists of a registry address (defaults to docker.io), repository, image, and tag (defaults to latest). Wildcards ('*' and '?') are allowed. See: https://kubernetes.io/docs/concepts/containers/images.",
										MarkdownDescription: "ImageReferences is a list of matching image reference patterns. At least one pattern in the list must match the image for the rule to apply. Each image reference consists of a registry address (defaults to docker.io), repository, image, and tag (defaults to latest). Wildcards ('*' and '?') are allowed. See: https://kubernetes.io/docs/concepts/containers/images.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"issuer": {
										Description:         "Issuer is the certificate issuer used for keyless signing. Deprecated. Use KeylessAttestor instead.",
										MarkdownDescription: "Issuer is the certificate issuer used for keyless signing. Deprecated. Use KeylessAttestor instead.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"key": {
										Description:         "Key is the PEM encoded public key that the image or attestation is signed with. Deprecated. Use StaticKeyAttestor instead.",
										MarkdownDescription: "Key is the PEM encoded public key that the image or attestation is signed with. Deprecated. Use StaticKeyAttestor instead.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"mutate_digest": {
										Description:         "MutateDigest enables replacement of image tags with digests. Defaults to true.",
										MarkdownDescription: "MutateDigest enables replacement of image tags with digests. Defaults to true.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"repository": {
										Description:         "Repository is an optional alternate OCI repository to use for image signatures and attestations that match this rule. If specified Repository will override the default OCI image repository configured for the installation. The repository can also be overridden per Attestor or Attestation.",
										MarkdownDescription: "Repository is an optional alternate OCI repository to use for image signatures and attestations that match this rule. If specified Repository will override the default OCI image repository configured for the installation. The repository can also be overridden per Attestor or Attestation.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"required": {
										Description:         "Required validates that images are verified i.e. have matched passed a signature or attestation check.",
										MarkdownDescription: "Required validates that images are verified i.e. have matched passed a signature or attestation check.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"roots": {
										Description:         "Roots is the PEM encoded Root certificate chain used for keyless signing Deprecated. Use KeylessAttestor instead.",
										MarkdownDescription: "Roots is the PEM encoded Root certificate chain used for keyless signing Deprecated. Use KeylessAttestor instead.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"subject": {
										Description:         "Subject is the identity used for keyless signing, for example an email address Deprecated. Use KeylessAttestor instead.",
										MarkdownDescription: "Subject is the identity used for keyless signing, for example an email address Deprecated. Use KeylessAttestor instead.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"verify_digest": {
										Description:         "VerifyDigest validates that images have a digest.",
										MarkdownDescription: "VerifyDigest validates that images have a digest.",

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

					"schema_validation": {
						Description:         "SchemaValidation skips validation checks for policies as well as patched resources. Optional. The default value is set to 'true', it must be set to 'false' to disable the validation checks.",
						MarkdownDescription: "SchemaValidation skips validation checks for policies as well as patched resources. Optional. The default value is set to 'true', it must be set to 'false' to disable the validation checks.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"validation_failure_action": {
						Description:         "ValidationFailureAction defines if a validation policy rule violation should block the admission review request (enforce), or allow (audit) the admission review request and report an error in a policy report. Optional. Allowed values are audit or enforce. The default value is 'audit'.",
						MarkdownDescription: "ValidationFailureAction defines if a validation policy rule violation should block the admission review request (enforce), or allow (audit) the admission review request and report an error in a policy report. Optional. Allowed values are audit or enforce. The default value is 'audit'.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("audit", "enforce", "Audit", "Enforce"),
						},
					},

					"validation_failure_action_overrides": {
						Description:         "ValidationFailureActionOverrides is a Cluster Policy attribute that specifies ValidationFailureAction namespace-wise. It overrides ValidationFailureAction for the specified namespaces.",
						MarkdownDescription: "ValidationFailureActionOverrides is a Cluster Policy attribute that specifies ValidationFailureAction namespace-wise. It overrides ValidationFailureAction for the specified namespaces.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"action": {
								Description:         "ValidationFailureAction defines the policy validation failure action",
								MarkdownDescription: "ValidationFailureAction defines the policy validation failure action",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("audit", "enforce"),
								},
							},

							"namespaces": {
								Description:         "",
								MarkdownDescription: "",

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

					"webhook_timeout_seconds": {
						Description:         "WebhookTimeoutSeconds specifies the maximum time in seconds allowed to apply this policy. After the configured time expires, the admission request may fail, or may simply ignore the policy results, based on the failure policy. The default timeout is 10s, the value must be between 1 and 30 seconds.",
						MarkdownDescription: "WebhookTimeoutSeconds specifies the maximum time in seconds allowed to apply this policy. After the configured time expires, the admission request may fail, or may simply ignore the policy results, based on the failure policy. The default timeout is 10s, the value must be between 1 and 30 seconds.",

						Type: types.Int64Type,

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

func (r *KyvernoIoPolicyV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_kyverno_io_policy_v1")

	var state KyvernoIoPolicyV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel KyvernoIoPolicyV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("kyverno.io/v1")
	goModel.Kind = utilities.Ptr("Policy")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *KyvernoIoPolicyV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kyverno_io_policy_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *KyvernoIoPolicyV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_kyverno_io_policy_v1")

	var state KyvernoIoPolicyV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel KyvernoIoPolicyV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("kyverno.io/v1")
	goModel.Kind = utilities.Ptr("Policy")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *KyvernoIoPolicyV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_kyverno_io_policy_v1")
	// NO-OP: Terraform removes the state automatically for us
}
