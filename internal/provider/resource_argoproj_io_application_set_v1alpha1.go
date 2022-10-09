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

type ArgoprojIoApplicationSetV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*ArgoprojIoApplicationSetV1Alpha1Resource)(nil)
)

type ArgoprojIoApplicationSetV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type ArgoprojIoApplicationSetV1Alpha1GoModel struct {
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
		Generators *[]struct {
			ClusterDecisionResource *struct {
				ConfigMapRef *string `tfsdk:"config_map_ref" yaml:"configMapRef,omitempty"`

				LabelSelector *struct {
					MatchExpressions *[]struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

						Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
					} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

					MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
				} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				RequeueAfterSeconds *int64 `tfsdk:"requeue_after_seconds" yaml:"requeueAfterSeconds,omitempty"`

				Template *struct {
					Metadata *struct {
						Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Finalizers *[]string `tfsdk:"finalizers" yaml:"finalizers,omitempty"`

						Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`

					Spec *struct {
						Destination *struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

							Server *string `tfsdk:"server" yaml:"server,omitempty"`
						} `tfsdk:"destination" yaml:"destination,omitempty"`

						IgnoreDifferences *[]struct {
							Group *string `tfsdk:"group" yaml:"group,omitempty"`

							JqPathExpressions *[]string `tfsdk:"jq_path_expressions" yaml:"jqPathExpressions,omitempty"`

							JsonPointers *[]string `tfsdk:"json_pointers" yaml:"jsonPointers,omitempty"`

							Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

							ManagedFieldsManagers *[]string `tfsdk:"managed_fields_managers" yaml:"managedFieldsManagers,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
						} `tfsdk:"ignore_differences" yaml:"ignoreDifferences,omitempty"`

						Info *[]struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"info" yaml:"info,omitempty"`

						Project *string `tfsdk:"project" yaml:"project,omitempty"`

						RevisionHistoryLimit *int64 `tfsdk:"revision_history_limit" yaml:"revisionHistoryLimit,omitempty"`

						Source *struct {
							Chart *string `tfsdk:"chart" yaml:"chart,omitempty"`

							Directory *struct {
								Exclude *string `tfsdk:"exclude" yaml:"exclude,omitempty"`

								Include *string `tfsdk:"include" yaml:"include,omitempty"`

								Jsonnet *struct {
									ExtVars *[]struct {
										Code *bool `tfsdk:"code" yaml:"code,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`

										Value *string `tfsdk:"value" yaml:"value,omitempty"`
									} `tfsdk:"ext_vars" yaml:"extVars,omitempty"`

									Libs *[]string `tfsdk:"libs" yaml:"libs,omitempty"`

									Tlas *[]struct {
										Code *bool `tfsdk:"code" yaml:"code,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`

										Value *string `tfsdk:"value" yaml:"value,omitempty"`
									} `tfsdk:"tlas" yaml:"tlas,omitempty"`
								} `tfsdk:"jsonnet" yaml:"jsonnet,omitempty"`

								Recurse *bool `tfsdk:"recurse" yaml:"recurse,omitempty"`
							} `tfsdk:"directory" yaml:"directory,omitempty"`

							Helm *struct {
								FileParameters *[]struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Path *string `tfsdk:"path" yaml:"path,omitempty"`
								} `tfsdk:"file_parameters" yaml:"fileParameters,omitempty"`

								IgnoreMissingValueFiles *bool `tfsdk:"ignore_missing_value_files" yaml:"ignoreMissingValueFiles,omitempty"`

								Parameters *[]struct {
									ForceString *bool `tfsdk:"force_string" yaml:"forceString,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"parameters" yaml:"parameters,omitempty"`

								PassCredentials *bool `tfsdk:"pass_credentials" yaml:"passCredentials,omitempty"`

								ReleaseName *string `tfsdk:"release_name" yaml:"releaseName,omitempty"`

								SkipCrds *bool `tfsdk:"skip_crds" yaml:"skipCrds,omitempty"`

								ValueFiles *[]string `tfsdk:"value_files" yaml:"valueFiles,omitempty"`

								Values *string `tfsdk:"values" yaml:"values,omitempty"`

								Version *string `tfsdk:"version" yaml:"version,omitempty"`
							} `tfsdk:"helm" yaml:"helm,omitempty"`

							Kustomize *struct {
								CommonAnnotations *map[string]string `tfsdk:"common_annotations" yaml:"commonAnnotations,omitempty"`

								CommonLabels *map[string]string `tfsdk:"common_labels" yaml:"commonLabels,omitempty"`

								ForceCommonAnnotations *bool `tfsdk:"force_common_annotations" yaml:"forceCommonAnnotations,omitempty"`

								ForceCommonLabels *bool `tfsdk:"force_common_labels" yaml:"forceCommonLabels,omitempty"`

								Images *[]string `tfsdk:"images" yaml:"images,omitempty"`

								NamePrefix *string `tfsdk:"name_prefix" yaml:"namePrefix,omitempty"`

								NameSuffix *string `tfsdk:"name_suffix" yaml:"nameSuffix,omitempty"`

								Version *string `tfsdk:"version" yaml:"version,omitempty"`
							} `tfsdk:"kustomize" yaml:"kustomize,omitempty"`

							Path *string `tfsdk:"path" yaml:"path,omitempty"`

							Plugin *struct {
								Env *[]struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"env" yaml:"env,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`
							} `tfsdk:"plugin" yaml:"plugin,omitempty"`

							RepoURL *string `tfsdk:"repo_url" yaml:"repoURL,omitempty"`

							TargetRevision *string `tfsdk:"target_revision" yaml:"targetRevision,omitempty"`
						} `tfsdk:"source" yaml:"source,omitempty"`

						SyncPolicy *struct {
							Automated *struct {
								AllowEmpty *bool `tfsdk:"allow_empty" yaml:"allowEmpty,omitempty"`

								Prune *bool `tfsdk:"prune" yaml:"prune,omitempty"`

								SelfHeal *bool `tfsdk:"self_heal" yaml:"selfHeal,omitempty"`
							} `tfsdk:"automated" yaml:"automated,omitempty"`

							Retry *struct {
								Backoff *struct {
									Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

									Factor *int64 `tfsdk:"factor" yaml:"factor,omitempty"`

									MaxDuration *string `tfsdk:"max_duration" yaml:"maxDuration,omitempty"`
								} `tfsdk:"backoff" yaml:"backoff,omitempty"`

								Limit *int64 `tfsdk:"limit" yaml:"limit,omitempty"`
							} `tfsdk:"retry" yaml:"retry,omitempty"`

							SyncOptions *[]string `tfsdk:"sync_options" yaml:"syncOptions,omitempty"`
						} `tfsdk:"sync_policy" yaml:"syncPolicy,omitempty"`
					} `tfsdk:"spec" yaml:"spec,omitempty"`
				} `tfsdk:"template" yaml:"template,omitempty"`

				Values *map[string]string `tfsdk:"values" yaml:"values,omitempty"`
			} `tfsdk:"cluster_decision_resource" yaml:"clusterDecisionResource,omitempty"`

			Clusters *struct {
				Selector *struct {
					MatchExpressions *[]struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

						Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
					} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

					MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
				} `tfsdk:"selector" yaml:"selector,omitempty"`

				Template *struct {
					Metadata *struct {
						Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Finalizers *[]string `tfsdk:"finalizers" yaml:"finalizers,omitempty"`

						Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`

					Spec *struct {
						Destination *struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

							Server *string `tfsdk:"server" yaml:"server,omitempty"`
						} `tfsdk:"destination" yaml:"destination,omitempty"`

						IgnoreDifferences *[]struct {
							Group *string `tfsdk:"group" yaml:"group,omitempty"`

							JqPathExpressions *[]string `tfsdk:"jq_path_expressions" yaml:"jqPathExpressions,omitempty"`

							JsonPointers *[]string `tfsdk:"json_pointers" yaml:"jsonPointers,omitempty"`

							Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

							ManagedFieldsManagers *[]string `tfsdk:"managed_fields_managers" yaml:"managedFieldsManagers,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
						} `tfsdk:"ignore_differences" yaml:"ignoreDifferences,omitempty"`

						Info *[]struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"info" yaml:"info,omitempty"`

						Project *string `tfsdk:"project" yaml:"project,omitempty"`

						RevisionHistoryLimit *int64 `tfsdk:"revision_history_limit" yaml:"revisionHistoryLimit,omitempty"`

						Source *struct {
							Chart *string `tfsdk:"chart" yaml:"chart,omitempty"`

							Directory *struct {
								Exclude *string `tfsdk:"exclude" yaml:"exclude,omitempty"`

								Include *string `tfsdk:"include" yaml:"include,omitempty"`

								Jsonnet *struct {
									ExtVars *[]struct {
										Code *bool `tfsdk:"code" yaml:"code,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`

										Value *string `tfsdk:"value" yaml:"value,omitempty"`
									} `tfsdk:"ext_vars" yaml:"extVars,omitempty"`

									Libs *[]string `tfsdk:"libs" yaml:"libs,omitempty"`

									Tlas *[]struct {
										Code *bool `tfsdk:"code" yaml:"code,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`

										Value *string `tfsdk:"value" yaml:"value,omitempty"`
									} `tfsdk:"tlas" yaml:"tlas,omitempty"`
								} `tfsdk:"jsonnet" yaml:"jsonnet,omitempty"`

								Recurse *bool `tfsdk:"recurse" yaml:"recurse,omitempty"`
							} `tfsdk:"directory" yaml:"directory,omitempty"`

							Helm *struct {
								FileParameters *[]struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Path *string `tfsdk:"path" yaml:"path,omitempty"`
								} `tfsdk:"file_parameters" yaml:"fileParameters,omitempty"`

								IgnoreMissingValueFiles *bool `tfsdk:"ignore_missing_value_files" yaml:"ignoreMissingValueFiles,omitempty"`

								Parameters *[]struct {
									ForceString *bool `tfsdk:"force_string" yaml:"forceString,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"parameters" yaml:"parameters,omitempty"`

								PassCredentials *bool `tfsdk:"pass_credentials" yaml:"passCredentials,omitempty"`

								ReleaseName *string `tfsdk:"release_name" yaml:"releaseName,omitempty"`

								SkipCrds *bool `tfsdk:"skip_crds" yaml:"skipCrds,omitempty"`

								ValueFiles *[]string `tfsdk:"value_files" yaml:"valueFiles,omitempty"`

								Values *string `tfsdk:"values" yaml:"values,omitempty"`

								Version *string `tfsdk:"version" yaml:"version,omitempty"`
							} `tfsdk:"helm" yaml:"helm,omitempty"`

							Kustomize *struct {
								CommonAnnotations *map[string]string `tfsdk:"common_annotations" yaml:"commonAnnotations,omitempty"`

								CommonLabels *map[string]string `tfsdk:"common_labels" yaml:"commonLabels,omitempty"`

								ForceCommonAnnotations *bool `tfsdk:"force_common_annotations" yaml:"forceCommonAnnotations,omitempty"`

								ForceCommonLabels *bool `tfsdk:"force_common_labels" yaml:"forceCommonLabels,omitempty"`

								Images *[]string `tfsdk:"images" yaml:"images,omitempty"`

								NamePrefix *string `tfsdk:"name_prefix" yaml:"namePrefix,omitempty"`

								NameSuffix *string `tfsdk:"name_suffix" yaml:"nameSuffix,omitempty"`

								Version *string `tfsdk:"version" yaml:"version,omitempty"`
							} `tfsdk:"kustomize" yaml:"kustomize,omitempty"`

							Path *string `tfsdk:"path" yaml:"path,omitempty"`

							Plugin *struct {
								Env *[]struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"env" yaml:"env,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`
							} `tfsdk:"plugin" yaml:"plugin,omitempty"`

							RepoURL *string `tfsdk:"repo_url" yaml:"repoURL,omitempty"`

							TargetRevision *string `tfsdk:"target_revision" yaml:"targetRevision,omitempty"`
						} `tfsdk:"source" yaml:"source,omitempty"`

						SyncPolicy *struct {
							Automated *struct {
								AllowEmpty *bool `tfsdk:"allow_empty" yaml:"allowEmpty,omitempty"`

								Prune *bool `tfsdk:"prune" yaml:"prune,omitempty"`

								SelfHeal *bool `tfsdk:"self_heal" yaml:"selfHeal,omitempty"`
							} `tfsdk:"automated" yaml:"automated,omitempty"`

							Retry *struct {
								Backoff *struct {
									Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

									Factor *int64 `tfsdk:"factor" yaml:"factor,omitempty"`

									MaxDuration *string `tfsdk:"max_duration" yaml:"maxDuration,omitempty"`
								} `tfsdk:"backoff" yaml:"backoff,omitempty"`

								Limit *int64 `tfsdk:"limit" yaml:"limit,omitempty"`
							} `tfsdk:"retry" yaml:"retry,omitempty"`

							SyncOptions *[]string `tfsdk:"sync_options" yaml:"syncOptions,omitempty"`
						} `tfsdk:"sync_policy" yaml:"syncPolicy,omitempty"`
					} `tfsdk:"spec" yaml:"spec,omitempty"`
				} `tfsdk:"template" yaml:"template,omitempty"`

				Values *map[string]string `tfsdk:"values" yaml:"values,omitempty"`
			} `tfsdk:"clusters" yaml:"clusters,omitempty"`

			Git *struct {
				Directories *[]struct {
					Exclude *bool `tfsdk:"exclude" yaml:"exclude,omitempty"`

					Path *string `tfsdk:"path" yaml:"path,omitempty"`
				} `tfsdk:"directories" yaml:"directories,omitempty"`

				Files *[]struct {
					Path *string `tfsdk:"path" yaml:"path,omitempty"`
				} `tfsdk:"files" yaml:"files,omitempty"`

				RepoURL *string `tfsdk:"repo_url" yaml:"repoURL,omitempty"`

				RequeueAfterSeconds *int64 `tfsdk:"requeue_after_seconds" yaml:"requeueAfterSeconds,omitempty"`

				Revision *string `tfsdk:"revision" yaml:"revision,omitempty"`

				Template *struct {
					Metadata *struct {
						Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Finalizers *[]string `tfsdk:"finalizers" yaml:"finalizers,omitempty"`

						Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`

					Spec *struct {
						Destination *struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

							Server *string `tfsdk:"server" yaml:"server,omitempty"`
						} `tfsdk:"destination" yaml:"destination,omitempty"`

						IgnoreDifferences *[]struct {
							Group *string `tfsdk:"group" yaml:"group,omitempty"`

							JqPathExpressions *[]string `tfsdk:"jq_path_expressions" yaml:"jqPathExpressions,omitempty"`

							JsonPointers *[]string `tfsdk:"json_pointers" yaml:"jsonPointers,omitempty"`

							Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

							ManagedFieldsManagers *[]string `tfsdk:"managed_fields_managers" yaml:"managedFieldsManagers,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
						} `tfsdk:"ignore_differences" yaml:"ignoreDifferences,omitempty"`

						Info *[]struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"info" yaml:"info,omitempty"`

						Project *string `tfsdk:"project" yaml:"project,omitempty"`

						RevisionHistoryLimit *int64 `tfsdk:"revision_history_limit" yaml:"revisionHistoryLimit,omitempty"`

						Source *struct {
							Chart *string `tfsdk:"chart" yaml:"chart,omitempty"`

							Directory *struct {
								Exclude *string `tfsdk:"exclude" yaml:"exclude,omitempty"`

								Include *string `tfsdk:"include" yaml:"include,omitempty"`

								Jsonnet *struct {
									ExtVars *[]struct {
										Code *bool `tfsdk:"code" yaml:"code,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`

										Value *string `tfsdk:"value" yaml:"value,omitempty"`
									} `tfsdk:"ext_vars" yaml:"extVars,omitempty"`

									Libs *[]string `tfsdk:"libs" yaml:"libs,omitempty"`

									Tlas *[]struct {
										Code *bool `tfsdk:"code" yaml:"code,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`

										Value *string `tfsdk:"value" yaml:"value,omitempty"`
									} `tfsdk:"tlas" yaml:"tlas,omitempty"`
								} `tfsdk:"jsonnet" yaml:"jsonnet,omitempty"`

								Recurse *bool `tfsdk:"recurse" yaml:"recurse,omitempty"`
							} `tfsdk:"directory" yaml:"directory,omitempty"`

							Helm *struct {
								FileParameters *[]struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Path *string `tfsdk:"path" yaml:"path,omitempty"`
								} `tfsdk:"file_parameters" yaml:"fileParameters,omitempty"`

								IgnoreMissingValueFiles *bool `tfsdk:"ignore_missing_value_files" yaml:"ignoreMissingValueFiles,omitempty"`

								Parameters *[]struct {
									ForceString *bool `tfsdk:"force_string" yaml:"forceString,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"parameters" yaml:"parameters,omitempty"`

								PassCredentials *bool `tfsdk:"pass_credentials" yaml:"passCredentials,omitempty"`

								ReleaseName *string `tfsdk:"release_name" yaml:"releaseName,omitempty"`

								SkipCrds *bool `tfsdk:"skip_crds" yaml:"skipCrds,omitempty"`

								ValueFiles *[]string `tfsdk:"value_files" yaml:"valueFiles,omitempty"`

								Values *string `tfsdk:"values" yaml:"values,omitempty"`

								Version *string `tfsdk:"version" yaml:"version,omitempty"`
							} `tfsdk:"helm" yaml:"helm,omitempty"`

							Kustomize *struct {
								CommonAnnotations *map[string]string `tfsdk:"common_annotations" yaml:"commonAnnotations,omitempty"`

								CommonLabels *map[string]string `tfsdk:"common_labels" yaml:"commonLabels,omitempty"`

								ForceCommonAnnotations *bool `tfsdk:"force_common_annotations" yaml:"forceCommonAnnotations,omitempty"`

								ForceCommonLabels *bool `tfsdk:"force_common_labels" yaml:"forceCommonLabels,omitempty"`

								Images *[]string `tfsdk:"images" yaml:"images,omitempty"`

								NamePrefix *string `tfsdk:"name_prefix" yaml:"namePrefix,omitempty"`

								NameSuffix *string `tfsdk:"name_suffix" yaml:"nameSuffix,omitempty"`

								Version *string `tfsdk:"version" yaml:"version,omitempty"`
							} `tfsdk:"kustomize" yaml:"kustomize,omitempty"`

							Path *string `tfsdk:"path" yaml:"path,omitempty"`

							Plugin *struct {
								Env *[]struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"env" yaml:"env,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`
							} `tfsdk:"plugin" yaml:"plugin,omitempty"`

							RepoURL *string `tfsdk:"repo_url" yaml:"repoURL,omitempty"`

							TargetRevision *string `tfsdk:"target_revision" yaml:"targetRevision,omitempty"`
						} `tfsdk:"source" yaml:"source,omitempty"`

						SyncPolicy *struct {
							Automated *struct {
								AllowEmpty *bool `tfsdk:"allow_empty" yaml:"allowEmpty,omitempty"`

								Prune *bool `tfsdk:"prune" yaml:"prune,omitempty"`

								SelfHeal *bool `tfsdk:"self_heal" yaml:"selfHeal,omitempty"`
							} `tfsdk:"automated" yaml:"automated,omitempty"`

							Retry *struct {
								Backoff *struct {
									Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

									Factor *int64 `tfsdk:"factor" yaml:"factor,omitempty"`

									MaxDuration *string `tfsdk:"max_duration" yaml:"maxDuration,omitempty"`
								} `tfsdk:"backoff" yaml:"backoff,omitempty"`

								Limit *int64 `tfsdk:"limit" yaml:"limit,omitempty"`
							} `tfsdk:"retry" yaml:"retry,omitempty"`

							SyncOptions *[]string `tfsdk:"sync_options" yaml:"syncOptions,omitempty"`
						} `tfsdk:"sync_policy" yaml:"syncPolicy,omitempty"`
					} `tfsdk:"spec" yaml:"spec,omitempty"`
				} `tfsdk:"template" yaml:"template,omitempty"`
			} `tfsdk:"git" yaml:"git,omitempty"`

			List *struct {
				Elements *[]string `tfsdk:"elements" yaml:"elements,omitempty"`

				Template *struct {
					Metadata *struct {
						Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Finalizers *[]string `tfsdk:"finalizers" yaml:"finalizers,omitempty"`

						Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`

					Spec *struct {
						Destination *struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

							Server *string `tfsdk:"server" yaml:"server,omitempty"`
						} `tfsdk:"destination" yaml:"destination,omitempty"`

						IgnoreDifferences *[]struct {
							Group *string `tfsdk:"group" yaml:"group,omitempty"`

							JqPathExpressions *[]string `tfsdk:"jq_path_expressions" yaml:"jqPathExpressions,omitempty"`

							JsonPointers *[]string `tfsdk:"json_pointers" yaml:"jsonPointers,omitempty"`

							Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

							ManagedFieldsManagers *[]string `tfsdk:"managed_fields_managers" yaml:"managedFieldsManagers,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
						} `tfsdk:"ignore_differences" yaml:"ignoreDifferences,omitempty"`

						Info *[]struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"info" yaml:"info,omitempty"`

						Project *string `tfsdk:"project" yaml:"project,omitempty"`

						RevisionHistoryLimit *int64 `tfsdk:"revision_history_limit" yaml:"revisionHistoryLimit,omitempty"`

						Source *struct {
							Chart *string `tfsdk:"chart" yaml:"chart,omitempty"`

							Directory *struct {
								Exclude *string `tfsdk:"exclude" yaml:"exclude,omitempty"`

								Include *string `tfsdk:"include" yaml:"include,omitempty"`

								Jsonnet *struct {
									ExtVars *[]struct {
										Code *bool `tfsdk:"code" yaml:"code,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`

										Value *string `tfsdk:"value" yaml:"value,omitempty"`
									} `tfsdk:"ext_vars" yaml:"extVars,omitempty"`

									Libs *[]string `tfsdk:"libs" yaml:"libs,omitempty"`

									Tlas *[]struct {
										Code *bool `tfsdk:"code" yaml:"code,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`

										Value *string `tfsdk:"value" yaml:"value,omitempty"`
									} `tfsdk:"tlas" yaml:"tlas,omitempty"`
								} `tfsdk:"jsonnet" yaml:"jsonnet,omitempty"`

								Recurse *bool `tfsdk:"recurse" yaml:"recurse,omitempty"`
							} `tfsdk:"directory" yaml:"directory,omitempty"`

							Helm *struct {
								FileParameters *[]struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Path *string `tfsdk:"path" yaml:"path,omitempty"`
								} `tfsdk:"file_parameters" yaml:"fileParameters,omitempty"`

								IgnoreMissingValueFiles *bool `tfsdk:"ignore_missing_value_files" yaml:"ignoreMissingValueFiles,omitempty"`

								Parameters *[]struct {
									ForceString *bool `tfsdk:"force_string" yaml:"forceString,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"parameters" yaml:"parameters,omitempty"`

								PassCredentials *bool `tfsdk:"pass_credentials" yaml:"passCredentials,omitempty"`

								ReleaseName *string `tfsdk:"release_name" yaml:"releaseName,omitempty"`

								SkipCrds *bool `tfsdk:"skip_crds" yaml:"skipCrds,omitempty"`

								ValueFiles *[]string `tfsdk:"value_files" yaml:"valueFiles,omitempty"`

								Values *string `tfsdk:"values" yaml:"values,omitempty"`

								Version *string `tfsdk:"version" yaml:"version,omitempty"`
							} `tfsdk:"helm" yaml:"helm,omitempty"`

							Kustomize *struct {
								CommonAnnotations *map[string]string `tfsdk:"common_annotations" yaml:"commonAnnotations,omitempty"`

								CommonLabels *map[string]string `tfsdk:"common_labels" yaml:"commonLabels,omitempty"`

								ForceCommonAnnotations *bool `tfsdk:"force_common_annotations" yaml:"forceCommonAnnotations,omitempty"`

								ForceCommonLabels *bool `tfsdk:"force_common_labels" yaml:"forceCommonLabels,omitempty"`

								Images *[]string `tfsdk:"images" yaml:"images,omitempty"`

								NamePrefix *string `tfsdk:"name_prefix" yaml:"namePrefix,omitempty"`

								NameSuffix *string `tfsdk:"name_suffix" yaml:"nameSuffix,omitempty"`

								Version *string `tfsdk:"version" yaml:"version,omitempty"`
							} `tfsdk:"kustomize" yaml:"kustomize,omitempty"`

							Path *string `tfsdk:"path" yaml:"path,omitempty"`

							Plugin *struct {
								Env *[]struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"env" yaml:"env,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`
							} `tfsdk:"plugin" yaml:"plugin,omitempty"`

							RepoURL *string `tfsdk:"repo_url" yaml:"repoURL,omitempty"`

							TargetRevision *string `tfsdk:"target_revision" yaml:"targetRevision,omitempty"`
						} `tfsdk:"source" yaml:"source,omitempty"`

						SyncPolicy *struct {
							Automated *struct {
								AllowEmpty *bool `tfsdk:"allow_empty" yaml:"allowEmpty,omitempty"`

								Prune *bool `tfsdk:"prune" yaml:"prune,omitempty"`

								SelfHeal *bool `tfsdk:"self_heal" yaml:"selfHeal,omitempty"`
							} `tfsdk:"automated" yaml:"automated,omitempty"`

							Retry *struct {
								Backoff *struct {
									Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

									Factor *int64 `tfsdk:"factor" yaml:"factor,omitempty"`

									MaxDuration *string `tfsdk:"max_duration" yaml:"maxDuration,omitempty"`
								} `tfsdk:"backoff" yaml:"backoff,omitempty"`

								Limit *int64 `tfsdk:"limit" yaml:"limit,omitempty"`
							} `tfsdk:"retry" yaml:"retry,omitempty"`

							SyncOptions *[]string `tfsdk:"sync_options" yaml:"syncOptions,omitempty"`
						} `tfsdk:"sync_policy" yaml:"syncPolicy,omitempty"`
					} `tfsdk:"spec" yaml:"spec,omitempty"`
				} `tfsdk:"template" yaml:"template,omitempty"`
			} `tfsdk:"list" yaml:"list,omitempty"`

			Matrix *struct {
				Generators *[]struct {
					ClusterDecisionResource *struct {
						ConfigMapRef *string `tfsdk:"config_map_ref" yaml:"configMapRef,omitempty"`

						LabelSelector *struct {
							MatchExpressions *[]struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

								Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
							} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

							MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
						} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						RequeueAfterSeconds *int64 `tfsdk:"requeue_after_seconds" yaml:"requeueAfterSeconds,omitempty"`

						Template *struct {
							Metadata *struct {
								Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

								Finalizers *[]string `tfsdk:"finalizers" yaml:"finalizers,omitempty"`

								Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
							} `tfsdk:"metadata" yaml:"metadata,omitempty"`

							Spec *struct {
								Destination *struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

									Server *string `tfsdk:"server" yaml:"server,omitempty"`
								} `tfsdk:"destination" yaml:"destination,omitempty"`

								IgnoreDifferences *[]struct {
									Group *string `tfsdk:"group" yaml:"group,omitempty"`

									JqPathExpressions *[]string `tfsdk:"jq_path_expressions" yaml:"jqPathExpressions,omitempty"`

									JsonPointers *[]string `tfsdk:"json_pointers" yaml:"jsonPointers,omitempty"`

									Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

									ManagedFieldsManagers *[]string `tfsdk:"managed_fields_managers" yaml:"managedFieldsManagers,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
								} `tfsdk:"ignore_differences" yaml:"ignoreDifferences,omitempty"`

								Info *[]struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"info" yaml:"info,omitempty"`

								Project *string `tfsdk:"project" yaml:"project,omitempty"`

								RevisionHistoryLimit *int64 `tfsdk:"revision_history_limit" yaml:"revisionHistoryLimit,omitempty"`

								Source *struct {
									Chart *string `tfsdk:"chart" yaml:"chart,omitempty"`

									Directory *struct {
										Exclude *string `tfsdk:"exclude" yaml:"exclude,omitempty"`

										Include *string `tfsdk:"include" yaml:"include,omitempty"`

										Jsonnet *struct {
											ExtVars *[]struct {
												Code *bool `tfsdk:"code" yaml:"code,omitempty"`

												Name *string `tfsdk:"name" yaml:"name,omitempty"`

												Value *string `tfsdk:"value" yaml:"value,omitempty"`
											} `tfsdk:"ext_vars" yaml:"extVars,omitempty"`

											Libs *[]string `tfsdk:"libs" yaml:"libs,omitempty"`

											Tlas *[]struct {
												Code *bool `tfsdk:"code" yaml:"code,omitempty"`

												Name *string `tfsdk:"name" yaml:"name,omitempty"`

												Value *string `tfsdk:"value" yaml:"value,omitempty"`
											} `tfsdk:"tlas" yaml:"tlas,omitempty"`
										} `tfsdk:"jsonnet" yaml:"jsonnet,omitempty"`

										Recurse *bool `tfsdk:"recurse" yaml:"recurse,omitempty"`
									} `tfsdk:"directory" yaml:"directory,omitempty"`

									Helm *struct {
										FileParameters *[]struct {
											Name *string `tfsdk:"name" yaml:"name,omitempty"`

											Path *string `tfsdk:"path" yaml:"path,omitempty"`
										} `tfsdk:"file_parameters" yaml:"fileParameters,omitempty"`

										IgnoreMissingValueFiles *bool `tfsdk:"ignore_missing_value_files" yaml:"ignoreMissingValueFiles,omitempty"`

										Parameters *[]struct {
											ForceString *bool `tfsdk:"force_string" yaml:"forceString,omitempty"`

											Name *string `tfsdk:"name" yaml:"name,omitempty"`

											Value *string `tfsdk:"value" yaml:"value,omitempty"`
										} `tfsdk:"parameters" yaml:"parameters,omitempty"`

										PassCredentials *bool `tfsdk:"pass_credentials" yaml:"passCredentials,omitempty"`

										ReleaseName *string `tfsdk:"release_name" yaml:"releaseName,omitempty"`

										SkipCrds *bool `tfsdk:"skip_crds" yaml:"skipCrds,omitempty"`

										ValueFiles *[]string `tfsdk:"value_files" yaml:"valueFiles,omitempty"`

										Values *string `tfsdk:"values" yaml:"values,omitempty"`

										Version *string `tfsdk:"version" yaml:"version,omitempty"`
									} `tfsdk:"helm" yaml:"helm,omitempty"`

									Kustomize *struct {
										CommonAnnotations *map[string]string `tfsdk:"common_annotations" yaml:"commonAnnotations,omitempty"`

										CommonLabels *map[string]string `tfsdk:"common_labels" yaml:"commonLabels,omitempty"`

										ForceCommonAnnotations *bool `tfsdk:"force_common_annotations" yaml:"forceCommonAnnotations,omitempty"`

										ForceCommonLabels *bool `tfsdk:"force_common_labels" yaml:"forceCommonLabels,omitempty"`

										Images *[]string `tfsdk:"images" yaml:"images,omitempty"`

										NamePrefix *string `tfsdk:"name_prefix" yaml:"namePrefix,omitempty"`

										NameSuffix *string `tfsdk:"name_suffix" yaml:"nameSuffix,omitempty"`

										Version *string `tfsdk:"version" yaml:"version,omitempty"`
									} `tfsdk:"kustomize" yaml:"kustomize,omitempty"`

									Path *string `tfsdk:"path" yaml:"path,omitempty"`

									Plugin *struct {
										Env *[]struct {
											Name *string `tfsdk:"name" yaml:"name,omitempty"`

											Value *string `tfsdk:"value" yaml:"value,omitempty"`
										} `tfsdk:"env" yaml:"env,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"plugin" yaml:"plugin,omitempty"`

									RepoURL *string `tfsdk:"repo_url" yaml:"repoURL,omitempty"`

									TargetRevision *string `tfsdk:"target_revision" yaml:"targetRevision,omitempty"`
								} `tfsdk:"source" yaml:"source,omitempty"`

								SyncPolicy *struct {
									Automated *struct {
										AllowEmpty *bool `tfsdk:"allow_empty" yaml:"allowEmpty,omitempty"`

										Prune *bool `tfsdk:"prune" yaml:"prune,omitempty"`

										SelfHeal *bool `tfsdk:"self_heal" yaml:"selfHeal,omitempty"`
									} `tfsdk:"automated" yaml:"automated,omitempty"`

									Retry *struct {
										Backoff *struct {
											Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

											Factor *int64 `tfsdk:"factor" yaml:"factor,omitempty"`

											MaxDuration *string `tfsdk:"max_duration" yaml:"maxDuration,omitempty"`
										} `tfsdk:"backoff" yaml:"backoff,omitempty"`

										Limit *int64 `tfsdk:"limit" yaml:"limit,omitempty"`
									} `tfsdk:"retry" yaml:"retry,omitempty"`

									SyncOptions *[]string `tfsdk:"sync_options" yaml:"syncOptions,omitempty"`
								} `tfsdk:"sync_policy" yaml:"syncPolicy,omitempty"`
							} `tfsdk:"spec" yaml:"spec,omitempty"`
						} `tfsdk:"template" yaml:"template,omitempty"`

						Values *map[string]string `tfsdk:"values" yaml:"values,omitempty"`
					} `tfsdk:"cluster_decision_resource" yaml:"clusterDecisionResource,omitempty"`

					Clusters *struct {
						Selector *struct {
							MatchExpressions *[]struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

								Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
							} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

							MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
						} `tfsdk:"selector" yaml:"selector,omitempty"`

						Template *struct {
							Metadata *struct {
								Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

								Finalizers *[]string `tfsdk:"finalizers" yaml:"finalizers,omitempty"`

								Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
							} `tfsdk:"metadata" yaml:"metadata,omitempty"`

							Spec *struct {
								Destination *struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

									Server *string `tfsdk:"server" yaml:"server,omitempty"`
								} `tfsdk:"destination" yaml:"destination,omitempty"`

								IgnoreDifferences *[]struct {
									Group *string `tfsdk:"group" yaml:"group,omitempty"`

									JqPathExpressions *[]string `tfsdk:"jq_path_expressions" yaml:"jqPathExpressions,omitempty"`

									JsonPointers *[]string `tfsdk:"json_pointers" yaml:"jsonPointers,omitempty"`

									Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

									ManagedFieldsManagers *[]string `tfsdk:"managed_fields_managers" yaml:"managedFieldsManagers,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
								} `tfsdk:"ignore_differences" yaml:"ignoreDifferences,omitempty"`

								Info *[]struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"info" yaml:"info,omitempty"`

								Project *string `tfsdk:"project" yaml:"project,omitempty"`

								RevisionHistoryLimit *int64 `tfsdk:"revision_history_limit" yaml:"revisionHistoryLimit,omitempty"`

								Source *struct {
									Chart *string `tfsdk:"chart" yaml:"chart,omitempty"`

									Directory *struct {
										Exclude *string `tfsdk:"exclude" yaml:"exclude,omitempty"`

										Include *string `tfsdk:"include" yaml:"include,omitempty"`

										Jsonnet *struct {
											ExtVars *[]struct {
												Code *bool `tfsdk:"code" yaml:"code,omitempty"`

												Name *string `tfsdk:"name" yaml:"name,omitempty"`

												Value *string `tfsdk:"value" yaml:"value,omitempty"`
											} `tfsdk:"ext_vars" yaml:"extVars,omitempty"`

											Libs *[]string `tfsdk:"libs" yaml:"libs,omitempty"`

											Tlas *[]struct {
												Code *bool `tfsdk:"code" yaml:"code,omitempty"`

												Name *string `tfsdk:"name" yaml:"name,omitempty"`

												Value *string `tfsdk:"value" yaml:"value,omitempty"`
											} `tfsdk:"tlas" yaml:"tlas,omitempty"`
										} `tfsdk:"jsonnet" yaml:"jsonnet,omitempty"`

										Recurse *bool `tfsdk:"recurse" yaml:"recurse,omitempty"`
									} `tfsdk:"directory" yaml:"directory,omitempty"`

									Helm *struct {
										FileParameters *[]struct {
											Name *string `tfsdk:"name" yaml:"name,omitempty"`

											Path *string `tfsdk:"path" yaml:"path,omitempty"`
										} `tfsdk:"file_parameters" yaml:"fileParameters,omitempty"`

										IgnoreMissingValueFiles *bool `tfsdk:"ignore_missing_value_files" yaml:"ignoreMissingValueFiles,omitempty"`

										Parameters *[]struct {
											ForceString *bool `tfsdk:"force_string" yaml:"forceString,omitempty"`

											Name *string `tfsdk:"name" yaml:"name,omitempty"`

											Value *string `tfsdk:"value" yaml:"value,omitempty"`
										} `tfsdk:"parameters" yaml:"parameters,omitempty"`

										PassCredentials *bool `tfsdk:"pass_credentials" yaml:"passCredentials,omitempty"`

										ReleaseName *string `tfsdk:"release_name" yaml:"releaseName,omitempty"`

										SkipCrds *bool `tfsdk:"skip_crds" yaml:"skipCrds,omitempty"`

										ValueFiles *[]string `tfsdk:"value_files" yaml:"valueFiles,omitempty"`

										Values *string `tfsdk:"values" yaml:"values,omitempty"`

										Version *string `tfsdk:"version" yaml:"version,omitempty"`
									} `tfsdk:"helm" yaml:"helm,omitempty"`

									Kustomize *struct {
										CommonAnnotations *map[string]string `tfsdk:"common_annotations" yaml:"commonAnnotations,omitempty"`

										CommonLabels *map[string]string `tfsdk:"common_labels" yaml:"commonLabels,omitempty"`

										ForceCommonAnnotations *bool `tfsdk:"force_common_annotations" yaml:"forceCommonAnnotations,omitempty"`

										ForceCommonLabels *bool `tfsdk:"force_common_labels" yaml:"forceCommonLabels,omitempty"`

										Images *[]string `tfsdk:"images" yaml:"images,omitempty"`

										NamePrefix *string `tfsdk:"name_prefix" yaml:"namePrefix,omitempty"`

										NameSuffix *string `tfsdk:"name_suffix" yaml:"nameSuffix,omitempty"`

										Version *string `tfsdk:"version" yaml:"version,omitempty"`
									} `tfsdk:"kustomize" yaml:"kustomize,omitempty"`

									Path *string `tfsdk:"path" yaml:"path,omitempty"`

									Plugin *struct {
										Env *[]struct {
											Name *string `tfsdk:"name" yaml:"name,omitempty"`

											Value *string `tfsdk:"value" yaml:"value,omitempty"`
										} `tfsdk:"env" yaml:"env,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"plugin" yaml:"plugin,omitempty"`

									RepoURL *string `tfsdk:"repo_url" yaml:"repoURL,omitempty"`

									TargetRevision *string `tfsdk:"target_revision" yaml:"targetRevision,omitempty"`
								} `tfsdk:"source" yaml:"source,omitempty"`

								SyncPolicy *struct {
									Automated *struct {
										AllowEmpty *bool `tfsdk:"allow_empty" yaml:"allowEmpty,omitempty"`

										Prune *bool `tfsdk:"prune" yaml:"prune,omitempty"`

										SelfHeal *bool `tfsdk:"self_heal" yaml:"selfHeal,omitempty"`
									} `tfsdk:"automated" yaml:"automated,omitempty"`

									Retry *struct {
										Backoff *struct {
											Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

											Factor *int64 `tfsdk:"factor" yaml:"factor,omitempty"`

											MaxDuration *string `tfsdk:"max_duration" yaml:"maxDuration,omitempty"`
										} `tfsdk:"backoff" yaml:"backoff,omitempty"`

										Limit *int64 `tfsdk:"limit" yaml:"limit,omitempty"`
									} `tfsdk:"retry" yaml:"retry,omitempty"`

									SyncOptions *[]string `tfsdk:"sync_options" yaml:"syncOptions,omitempty"`
								} `tfsdk:"sync_policy" yaml:"syncPolicy,omitempty"`
							} `tfsdk:"spec" yaml:"spec,omitempty"`
						} `tfsdk:"template" yaml:"template,omitempty"`

						Values *map[string]string `tfsdk:"values" yaml:"values,omitempty"`
					} `tfsdk:"clusters" yaml:"clusters,omitempty"`

					Git *struct {
						Directories *[]struct {
							Exclude *bool `tfsdk:"exclude" yaml:"exclude,omitempty"`

							Path *string `tfsdk:"path" yaml:"path,omitempty"`
						} `tfsdk:"directories" yaml:"directories,omitempty"`

						Files *[]struct {
							Path *string `tfsdk:"path" yaml:"path,omitempty"`
						} `tfsdk:"files" yaml:"files,omitempty"`

						RepoURL *string `tfsdk:"repo_url" yaml:"repoURL,omitempty"`

						RequeueAfterSeconds *int64 `tfsdk:"requeue_after_seconds" yaml:"requeueAfterSeconds,omitempty"`

						Revision *string `tfsdk:"revision" yaml:"revision,omitempty"`

						Template *struct {
							Metadata *struct {
								Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

								Finalizers *[]string `tfsdk:"finalizers" yaml:"finalizers,omitempty"`

								Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
							} `tfsdk:"metadata" yaml:"metadata,omitempty"`

							Spec *struct {
								Destination *struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

									Server *string `tfsdk:"server" yaml:"server,omitempty"`
								} `tfsdk:"destination" yaml:"destination,omitempty"`

								IgnoreDifferences *[]struct {
									Group *string `tfsdk:"group" yaml:"group,omitempty"`

									JqPathExpressions *[]string `tfsdk:"jq_path_expressions" yaml:"jqPathExpressions,omitempty"`

									JsonPointers *[]string `tfsdk:"json_pointers" yaml:"jsonPointers,omitempty"`

									Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

									ManagedFieldsManagers *[]string `tfsdk:"managed_fields_managers" yaml:"managedFieldsManagers,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
								} `tfsdk:"ignore_differences" yaml:"ignoreDifferences,omitempty"`

								Info *[]struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"info" yaml:"info,omitempty"`

								Project *string `tfsdk:"project" yaml:"project,omitempty"`

								RevisionHistoryLimit *int64 `tfsdk:"revision_history_limit" yaml:"revisionHistoryLimit,omitempty"`

								Source *struct {
									Chart *string `tfsdk:"chart" yaml:"chart,omitempty"`

									Directory *struct {
										Exclude *string `tfsdk:"exclude" yaml:"exclude,omitempty"`

										Include *string `tfsdk:"include" yaml:"include,omitempty"`

										Jsonnet *struct {
											ExtVars *[]struct {
												Code *bool `tfsdk:"code" yaml:"code,omitempty"`

												Name *string `tfsdk:"name" yaml:"name,omitempty"`

												Value *string `tfsdk:"value" yaml:"value,omitempty"`
											} `tfsdk:"ext_vars" yaml:"extVars,omitempty"`

											Libs *[]string `tfsdk:"libs" yaml:"libs,omitempty"`

											Tlas *[]struct {
												Code *bool `tfsdk:"code" yaml:"code,omitempty"`

												Name *string `tfsdk:"name" yaml:"name,omitempty"`

												Value *string `tfsdk:"value" yaml:"value,omitempty"`
											} `tfsdk:"tlas" yaml:"tlas,omitempty"`
										} `tfsdk:"jsonnet" yaml:"jsonnet,omitempty"`

										Recurse *bool `tfsdk:"recurse" yaml:"recurse,omitempty"`
									} `tfsdk:"directory" yaml:"directory,omitempty"`

									Helm *struct {
										FileParameters *[]struct {
											Name *string `tfsdk:"name" yaml:"name,omitempty"`

											Path *string `tfsdk:"path" yaml:"path,omitempty"`
										} `tfsdk:"file_parameters" yaml:"fileParameters,omitempty"`

										IgnoreMissingValueFiles *bool `tfsdk:"ignore_missing_value_files" yaml:"ignoreMissingValueFiles,omitempty"`

										Parameters *[]struct {
											ForceString *bool `tfsdk:"force_string" yaml:"forceString,omitempty"`

											Name *string `tfsdk:"name" yaml:"name,omitempty"`

											Value *string `tfsdk:"value" yaml:"value,omitempty"`
										} `tfsdk:"parameters" yaml:"parameters,omitempty"`

										PassCredentials *bool `tfsdk:"pass_credentials" yaml:"passCredentials,omitempty"`

										ReleaseName *string `tfsdk:"release_name" yaml:"releaseName,omitempty"`

										SkipCrds *bool `tfsdk:"skip_crds" yaml:"skipCrds,omitempty"`

										ValueFiles *[]string `tfsdk:"value_files" yaml:"valueFiles,omitempty"`

										Values *string `tfsdk:"values" yaml:"values,omitempty"`

										Version *string `tfsdk:"version" yaml:"version,omitempty"`
									} `tfsdk:"helm" yaml:"helm,omitempty"`

									Kustomize *struct {
										CommonAnnotations *map[string]string `tfsdk:"common_annotations" yaml:"commonAnnotations,omitempty"`

										CommonLabels *map[string]string `tfsdk:"common_labels" yaml:"commonLabels,omitempty"`

										ForceCommonAnnotations *bool `tfsdk:"force_common_annotations" yaml:"forceCommonAnnotations,omitempty"`

										ForceCommonLabels *bool `tfsdk:"force_common_labels" yaml:"forceCommonLabels,omitempty"`

										Images *[]string `tfsdk:"images" yaml:"images,omitempty"`

										NamePrefix *string `tfsdk:"name_prefix" yaml:"namePrefix,omitempty"`

										NameSuffix *string `tfsdk:"name_suffix" yaml:"nameSuffix,omitempty"`

										Version *string `tfsdk:"version" yaml:"version,omitempty"`
									} `tfsdk:"kustomize" yaml:"kustomize,omitempty"`

									Path *string `tfsdk:"path" yaml:"path,omitempty"`

									Plugin *struct {
										Env *[]struct {
											Name *string `tfsdk:"name" yaml:"name,omitempty"`

											Value *string `tfsdk:"value" yaml:"value,omitempty"`
										} `tfsdk:"env" yaml:"env,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"plugin" yaml:"plugin,omitempty"`

									RepoURL *string `tfsdk:"repo_url" yaml:"repoURL,omitempty"`

									TargetRevision *string `tfsdk:"target_revision" yaml:"targetRevision,omitempty"`
								} `tfsdk:"source" yaml:"source,omitempty"`

								SyncPolicy *struct {
									Automated *struct {
										AllowEmpty *bool `tfsdk:"allow_empty" yaml:"allowEmpty,omitempty"`

										Prune *bool `tfsdk:"prune" yaml:"prune,omitempty"`

										SelfHeal *bool `tfsdk:"self_heal" yaml:"selfHeal,omitempty"`
									} `tfsdk:"automated" yaml:"automated,omitempty"`

									Retry *struct {
										Backoff *struct {
											Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

											Factor *int64 `tfsdk:"factor" yaml:"factor,omitempty"`

											MaxDuration *string `tfsdk:"max_duration" yaml:"maxDuration,omitempty"`
										} `tfsdk:"backoff" yaml:"backoff,omitempty"`

										Limit *int64 `tfsdk:"limit" yaml:"limit,omitempty"`
									} `tfsdk:"retry" yaml:"retry,omitempty"`

									SyncOptions *[]string `tfsdk:"sync_options" yaml:"syncOptions,omitempty"`
								} `tfsdk:"sync_policy" yaml:"syncPolicy,omitempty"`
							} `tfsdk:"spec" yaml:"spec,omitempty"`
						} `tfsdk:"template" yaml:"template,omitempty"`
					} `tfsdk:"git" yaml:"git,omitempty"`

					List *struct {
						Elements *[]string `tfsdk:"elements" yaml:"elements,omitempty"`

						Template *struct {
							Metadata *struct {
								Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

								Finalizers *[]string `tfsdk:"finalizers" yaml:"finalizers,omitempty"`

								Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
							} `tfsdk:"metadata" yaml:"metadata,omitempty"`

							Spec *struct {
								Destination *struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

									Server *string `tfsdk:"server" yaml:"server,omitempty"`
								} `tfsdk:"destination" yaml:"destination,omitempty"`

								IgnoreDifferences *[]struct {
									Group *string `tfsdk:"group" yaml:"group,omitempty"`

									JqPathExpressions *[]string `tfsdk:"jq_path_expressions" yaml:"jqPathExpressions,omitempty"`

									JsonPointers *[]string `tfsdk:"json_pointers" yaml:"jsonPointers,omitempty"`

									Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

									ManagedFieldsManagers *[]string `tfsdk:"managed_fields_managers" yaml:"managedFieldsManagers,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
								} `tfsdk:"ignore_differences" yaml:"ignoreDifferences,omitempty"`

								Info *[]struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"info" yaml:"info,omitempty"`

								Project *string `tfsdk:"project" yaml:"project,omitempty"`

								RevisionHistoryLimit *int64 `tfsdk:"revision_history_limit" yaml:"revisionHistoryLimit,omitempty"`

								Source *struct {
									Chart *string `tfsdk:"chart" yaml:"chart,omitempty"`

									Directory *struct {
										Exclude *string `tfsdk:"exclude" yaml:"exclude,omitempty"`

										Include *string `tfsdk:"include" yaml:"include,omitempty"`

										Jsonnet *struct {
											ExtVars *[]struct {
												Code *bool `tfsdk:"code" yaml:"code,omitempty"`

												Name *string `tfsdk:"name" yaml:"name,omitempty"`

												Value *string `tfsdk:"value" yaml:"value,omitempty"`
											} `tfsdk:"ext_vars" yaml:"extVars,omitempty"`

											Libs *[]string `tfsdk:"libs" yaml:"libs,omitempty"`

											Tlas *[]struct {
												Code *bool `tfsdk:"code" yaml:"code,omitempty"`

												Name *string `tfsdk:"name" yaml:"name,omitempty"`

												Value *string `tfsdk:"value" yaml:"value,omitempty"`
											} `tfsdk:"tlas" yaml:"tlas,omitempty"`
										} `tfsdk:"jsonnet" yaml:"jsonnet,omitempty"`

										Recurse *bool `tfsdk:"recurse" yaml:"recurse,omitempty"`
									} `tfsdk:"directory" yaml:"directory,omitempty"`

									Helm *struct {
										FileParameters *[]struct {
											Name *string `tfsdk:"name" yaml:"name,omitempty"`

											Path *string `tfsdk:"path" yaml:"path,omitempty"`
										} `tfsdk:"file_parameters" yaml:"fileParameters,omitempty"`

										IgnoreMissingValueFiles *bool `tfsdk:"ignore_missing_value_files" yaml:"ignoreMissingValueFiles,omitempty"`

										Parameters *[]struct {
											ForceString *bool `tfsdk:"force_string" yaml:"forceString,omitempty"`

											Name *string `tfsdk:"name" yaml:"name,omitempty"`

											Value *string `tfsdk:"value" yaml:"value,omitempty"`
										} `tfsdk:"parameters" yaml:"parameters,omitempty"`

										PassCredentials *bool `tfsdk:"pass_credentials" yaml:"passCredentials,omitempty"`

										ReleaseName *string `tfsdk:"release_name" yaml:"releaseName,omitempty"`

										SkipCrds *bool `tfsdk:"skip_crds" yaml:"skipCrds,omitempty"`

										ValueFiles *[]string `tfsdk:"value_files" yaml:"valueFiles,omitempty"`

										Values *string `tfsdk:"values" yaml:"values,omitempty"`

										Version *string `tfsdk:"version" yaml:"version,omitempty"`
									} `tfsdk:"helm" yaml:"helm,omitempty"`

									Kustomize *struct {
										CommonAnnotations *map[string]string `tfsdk:"common_annotations" yaml:"commonAnnotations,omitempty"`

										CommonLabels *map[string]string `tfsdk:"common_labels" yaml:"commonLabels,omitempty"`

										ForceCommonAnnotations *bool `tfsdk:"force_common_annotations" yaml:"forceCommonAnnotations,omitempty"`

										ForceCommonLabels *bool `tfsdk:"force_common_labels" yaml:"forceCommonLabels,omitempty"`

										Images *[]string `tfsdk:"images" yaml:"images,omitempty"`

										NamePrefix *string `tfsdk:"name_prefix" yaml:"namePrefix,omitempty"`

										NameSuffix *string `tfsdk:"name_suffix" yaml:"nameSuffix,omitempty"`

										Version *string `tfsdk:"version" yaml:"version,omitempty"`
									} `tfsdk:"kustomize" yaml:"kustomize,omitempty"`

									Path *string `tfsdk:"path" yaml:"path,omitempty"`

									Plugin *struct {
										Env *[]struct {
											Name *string `tfsdk:"name" yaml:"name,omitempty"`

											Value *string `tfsdk:"value" yaml:"value,omitempty"`
										} `tfsdk:"env" yaml:"env,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"plugin" yaml:"plugin,omitempty"`

									RepoURL *string `tfsdk:"repo_url" yaml:"repoURL,omitempty"`

									TargetRevision *string `tfsdk:"target_revision" yaml:"targetRevision,omitempty"`
								} `tfsdk:"source" yaml:"source,omitempty"`

								SyncPolicy *struct {
									Automated *struct {
										AllowEmpty *bool `tfsdk:"allow_empty" yaml:"allowEmpty,omitempty"`

										Prune *bool `tfsdk:"prune" yaml:"prune,omitempty"`

										SelfHeal *bool `tfsdk:"self_heal" yaml:"selfHeal,omitempty"`
									} `tfsdk:"automated" yaml:"automated,omitempty"`

									Retry *struct {
										Backoff *struct {
											Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

											Factor *int64 `tfsdk:"factor" yaml:"factor,omitempty"`

											MaxDuration *string `tfsdk:"max_duration" yaml:"maxDuration,omitempty"`
										} `tfsdk:"backoff" yaml:"backoff,omitempty"`

										Limit *int64 `tfsdk:"limit" yaml:"limit,omitempty"`
									} `tfsdk:"retry" yaml:"retry,omitempty"`

									SyncOptions *[]string `tfsdk:"sync_options" yaml:"syncOptions,omitempty"`
								} `tfsdk:"sync_policy" yaml:"syncPolicy,omitempty"`
							} `tfsdk:"spec" yaml:"spec,omitempty"`
						} `tfsdk:"template" yaml:"template,omitempty"`
					} `tfsdk:"list" yaml:"list,omitempty"`

					Matrix *map[string]string `tfsdk:"matrix" yaml:"matrix,omitempty"`

					Merge *map[string]string `tfsdk:"merge" yaml:"merge,omitempty"`

					PullRequest *struct {
						BitbucketServer *struct {
							Api *string `tfsdk:"api" yaml:"api,omitempty"`

							BasicAuth *struct {
								PasswordRef *struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
								} `tfsdk:"password_ref" yaml:"passwordRef,omitempty"`

								Username *string `tfsdk:"username" yaml:"username,omitempty"`
							} `tfsdk:"basic_auth" yaml:"basicAuth,omitempty"`

							Project *string `tfsdk:"project" yaml:"project,omitempty"`

							Repo *string `tfsdk:"repo" yaml:"repo,omitempty"`
						} `tfsdk:"bitbucket_server" yaml:"bitbucketServer,omitempty"`

						Filters *[]struct {
							BranchMatch *string `tfsdk:"branch_match" yaml:"branchMatch,omitempty"`
						} `tfsdk:"filters" yaml:"filters,omitempty"`

						Gitea *struct {
							Api *string `tfsdk:"api" yaml:"api,omitempty"`

							Insecure *bool `tfsdk:"insecure" yaml:"insecure,omitempty"`

							Owner *string `tfsdk:"owner" yaml:"owner,omitempty"`

							Repo *string `tfsdk:"repo" yaml:"repo,omitempty"`

							TokenRef *struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
							} `tfsdk:"token_ref" yaml:"tokenRef,omitempty"`
						} `tfsdk:"gitea" yaml:"gitea,omitempty"`

						Github *struct {
							Api *string `tfsdk:"api" yaml:"api,omitempty"`

							AppSecretName *string `tfsdk:"app_secret_name" yaml:"appSecretName,omitempty"`

							Labels *[]string `tfsdk:"labels" yaml:"labels,omitempty"`

							Owner *string `tfsdk:"owner" yaml:"owner,omitempty"`

							Repo *string `tfsdk:"repo" yaml:"repo,omitempty"`

							TokenRef *struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
							} `tfsdk:"token_ref" yaml:"tokenRef,omitempty"`
						} `tfsdk:"github" yaml:"github,omitempty"`

						Gitlab *struct {
							Api *string `tfsdk:"api" yaml:"api,omitempty"`

							Labels *[]string `tfsdk:"labels" yaml:"labels,omitempty"`

							Project *string `tfsdk:"project" yaml:"project,omitempty"`

							PullRequestState *string `tfsdk:"pull_request_state" yaml:"pullRequestState,omitempty"`

							TokenRef *struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
							} `tfsdk:"token_ref" yaml:"tokenRef,omitempty"`
						} `tfsdk:"gitlab" yaml:"gitlab,omitempty"`

						RequeueAfterSeconds *int64 `tfsdk:"requeue_after_seconds" yaml:"requeueAfterSeconds,omitempty"`

						Template *struct {
							Metadata *struct {
								Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

								Finalizers *[]string `tfsdk:"finalizers" yaml:"finalizers,omitempty"`

								Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
							} `tfsdk:"metadata" yaml:"metadata,omitempty"`

							Spec *struct {
								Destination *struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

									Server *string `tfsdk:"server" yaml:"server,omitempty"`
								} `tfsdk:"destination" yaml:"destination,omitempty"`

								IgnoreDifferences *[]struct {
									Group *string `tfsdk:"group" yaml:"group,omitempty"`

									JqPathExpressions *[]string `tfsdk:"jq_path_expressions" yaml:"jqPathExpressions,omitempty"`

									JsonPointers *[]string `tfsdk:"json_pointers" yaml:"jsonPointers,omitempty"`

									Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

									ManagedFieldsManagers *[]string `tfsdk:"managed_fields_managers" yaml:"managedFieldsManagers,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
								} `tfsdk:"ignore_differences" yaml:"ignoreDifferences,omitempty"`

								Info *[]struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"info" yaml:"info,omitempty"`

								Project *string `tfsdk:"project" yaml:"project,omitempty"`

								RevisionHistoryLimit *int64 `tfsdk:"revision_history_limit" yaml:"revisionHistoryLimit,omitempty"`

								Source *struct {
									Chart *string `tfsdk:"chart" yaml:"chart,omitempty"`

									Directory *struct {
										Exclude *string `tfsdk:"exclude" yaml:"exclude,omitempty"`

										Include *string `tfsdk:"include" yaml:"include,omitempty"`

										Jsonnet *struct {
											ExtVars *[]struct {
												Code *bool `tfsdk:"code" yaml:"code,omitempty"`

												Name *string `tfsdk:"name" yaml:"name,omitempty"`

												Value *string `tfsdk:"value" yaml:"value,omitempty"`
											} `tfsdk:"ext_vars" yaml:"extVars,omitempty"`

											Libs *[]string `tfsdk:"libs" yaml:"libs,omitempty"`

											Tlas *[]struct {
												Code *bool `tfsdk:"code" yaml:"code,omitempty"`

												Name *string `tfsdk:"name" yaml:"name,omitempty"`

												Value *string `tfsdk:"value" yaml:"value,omitempty"`
											} `tfsdk:"tlas" yaml:"tlas,omitempty"`
										} `tfsdk:"jsonnet" yaml:"jsonnet,omitempty"`

										Recurse *bool `tfsdk:"recurse" yaml:"recurse,omitempty"`
									} `tfsdk:"directory" yaml:"directory,omitempty"`

									Helm *struct {
										FileParameters *[]struct {
											Name *string `tfsdk:"name" yaml:"name,omitempty"`

											Path *string `tfsdk:"path" yaml:"path,omitempty"`
										} `tfsdk:"file_parameters" yaml:"fileParameters,omitempty"`

										IgnoreMissingValueFiles *bool `tfsdk:"ignore_missing_value_files" yaml:"ignoreMissingValueFiles,omitempty"`

										Parameters *[]struct {
											ForceString *bool `tfsdk:"force_string" yaml:"forceString,omitempty"`

											Name *string `tfsdk:"name" yaml:"name,omitempty"`

											Value *string `tfsdk:"value" yaml:"value,omitempty"`
										} `tfsdk:"parameters" yaml:"parameters,omitempty"`

										PassCredentials *bool `tfsdk:"pass_credentials" yaml:"passCredentials,omitempty"`

										ReleaseName *string `tfsdk:"release_name" yaml:"releaseName,omitempty"`

										SkipCrds *bool `tfsdk:"skip_crds" yaml:"skipCrds,omitempty"`

										ValueFiles *[]string `tfsdk:"value_files" yaml:"valueFiles,omitempty"`

										Values *string `tfsdk:"values" yaml:"values,omitempty"`

										Version *string `tfsdk:"version" yaml:"version,omitempty"`
									} `tfsdk:"helm" yaml:"helm,omitempty"`

									Kustomize *struct {
										CommonAnnotations *map[string]string `tfsdk:"common_annotations" yaml:"commonAnnotations,omitempty"`

										CommonLabels *map[string]string `tfsdk:"common_labels" yaml:"commonLabels,omitempty"`

										ForceCommonAnnotations *bool `tfsdk:"force_common_annotations" yaml:"forceCommonAnnotations,omitempty"`

										ForceCommonLabels *bool `tfsdk:"force_common_labels" yaml:"forceCommonLabels,omitempty"`

										Images *[]string `tfsdk:"images" yaml:"images,omitempty"`

										NamePrefix *string `tfsdk:"name_prefix" yaml:"namePrefix,omitempty"`

										NameSuffix *string `tfsdk:"name_suffix" yaml:"nameSuffix,omitempty"`

										Version *string `tfsdk:"version" yaml:"version,omitempty"`
									} `tfsdk:"kustomize" yaml:"kustomize,omitempty"`

									Path *string `tfsdk:"path" yaml:"path,omitempty"`

									Plugin *struct {
										Env *[]struct {
											Name *string `tfsdk:"name" yaml:"name,omitempty"`

											Value *string `tfsdk:"value" yaml:"value,omitempty"`
										} `tfsdk:"env" yaml:"env,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"plugin" yaml:"plugin,omitempty"`

									RepoURL *string `tfsdk:"repo_url" yaml:"repoURL,omitempty"`

									TargetRevision *string `tfsdk:"target_revision" yaml:"targetRevision,omitempty"`
								} `tfsdk:"source" yaml:"source,omitempty"`

								SyncPolicy *struct {
									Automated *struct {
										AllowEmpty *bool `tfsdk:"allow_empty" yaml:"allowEmpty,omitempty"`

										Prune *bool `tfsdk:"prune" yaml:"prune,omitempty"`

										SelfHeal *bool `tfsdk:"self_heal" yaml:"selfHeal,omitempty"`
									} `tfsdk:"automated" yaml:"automated,omitempty"`

									Retry *struct {
										Backoff *struct {
											Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

											Factor *int64 `tfsdk:"factor" yaml:"factor,omitempty"`

											MaxDuration *string `tfsdk:"max_duration" yaml:"maxDuration,omitempty"`
										} `tfsdk:"backoff" yaml:"backoff,omitempty"`

										Limit *int64 `tfsdk:"limit" yaml:"limit,omitempty"`
									} `tfsdk:"retry" yaml:"retry,omitempty"`

									SyncOptions *[]string `tfsdk:"sync_options" yaml:"syncOptions,omitempty"`
								} `tfsdk:"sync_policy" yaml:"syncPolicy,omitempty"`
							} `tfsdk:"spec" yaml:"spec,omitempty"`
						} `tfsdk:"template" yaml:"template,omitempty"`
					} `tfsdk:"pull_request" yaml:"pullRequest,omitempty"`

					ScmProvider *struct {
						AzureDevOps *struct {
							AccessTokenRef *struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
							} `tfsdk:"access_token_ref" yaml:"accessTokenRef,omitempty"`

							AllBranches *bool `tfsdk:"all_branches" yaml:"allBranches,omitempty"`

							Api *string `tfsdk:"api" yaml:"api,omitempty"`

							Organization *string `tfsdk:"organization" yaml:"organization,omitempty"`

							TeamProject *string `tfsdk:"team_project" yaml:"teamProject,omitempty"`
						} `tfsdk:"azure_dev_ops" yaml:"azureDevOps,omitempty"`

						Bitbucket *struct {
							AllBranches *bool `tfsdk:"all_branches" yaml:"allBranches,omitempty"`

							AppPasswordRef *struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
							} `tfsdk:"app_password_ref" yaml:"appPasswordRef,omitempty"`

							Owner *string `tfsdk:"owner" yaml:"owner,omitempty"`

							User *string `tfsdk:"user" yaml:"user,omitempty"`
						} `tfsdk:"bitbucket" yaml:"bitbucket,omitempty"`

						BitbucketServer *struct {
							AllBranches *bool `tfsdk:"all_branches" yaml:"allBranches,omitempty"`

							Api *string `tfsdk:"api" yaml:"api,omitempty"`

							BasicAuth *struct {
								PasswordRef *struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
								} `tfsdk:"password_ref" yaml:"passwordRef,omitempty"`

								Username *string `tfsdk:"username" yaml:"username,omitempty"`
							} `tfsdk:"basic_auth" yaml:"basicAuth,omitempty"`

							Project *string `tfsdk:"project" yaml:"project,omitempty"`
						} `tfsdk:"bitbucket_server" yaml:"bitbucketServer,omitempty"`

						CloneProtocol *string `tfsdk:"clone_protocol" yaml:"cloneProtocol,omitempty"`

						Filters *[]struct {
							BranchMatch *string `tfsdk:"branch_match" yaml:"branchMatch,omitempty"`

							LabelMatch *string `tfsdk:"label_match" yaml:"labelMatch,omitempty"`

							PathsDoNotExist *[]string `tfsdk:"paths_do_not_exist" yaml:"pathsDoNotExist,omitempty"`

							PathsExist *[]string `tfsdk:"paths_exist" yaml:"pathsExist,omitempty"`

							RepositoryMatch *string `tfsdk:"repository_match" yaml:"repositoryMatch,omitempty"`
						} `tfsdk:"filters" yaml:"filters,omitempty"`

						Gitea *struct {
							AllBranches *bool `tfsdk:"all_branches" yaml:"allBranches,omitempty"`

							Api *string `tfsdk:"api" yaml:"api,omitempty"`

							Insecure *bool `tfsdk:"insecure" yaml:"insecure,omitempty"`

							Owner *string `tfsdk:"owner" yaml:"owner,omitempty"`

							TokenRef *struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
							} `tfsdk:"token_ref" yaml:"tokenRef,omitempty"`
						} `tfsdk:"gitea" yaml:"gitea,omitempty"`

						Github *struct {
							AllBranches *bool `tfsdk:"all_branches" yaml:"allBranches,omitempty"`

							Api *string `tfsdk:"api" yaml:"api,omitempty"`

							AppSecretName *string `tfsdk:"app_secret_name" yaml:"appSecretName,omitempty"`

							Organization *string `tfsdk:"organization" yaml:"organization,omitempty"`

							TokenRef *struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
							} `tfsdk:"token_ref" yaml:"tokenRef,omitempty"`
						} `tfsdk:"github" yaml:"github,omitempty"`

						Gitlab *struct {
							AllBranches *bool `tfsdk:"all_branches" yaml:"allBranches,omitempty"`

							Api *string `tfsdk:"api" yaml:"api,omitempty"`

							Group *string `tfsdk:"group" yaml:"group,omitempty"`

							IncludeSubgroups *bool `tfsdk:"include_subgroups" yaml:"includeSubgroups,omitempty"`

							TokenRef *struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
							} `tfsdk:"token_ref" yaml:"tokenRef,omitempty"`
						} `tfsdk:"gitlab" yaml:"gitlab,omitempty"`

						RequeueAfterSeconds *int64 `tfsdk:"requeue_after_seconds" yaml:"requeueAfterSeconds,omitempty"`

						Template *struct {
							Metadata *struct {
								Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

								Finalizers *[]string `tfsdk:"finalizers" yaml:"finalizers,omitempty"`

								Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
							} `tfsdk:"metadata" yaml:"metadata,omitempty"`

							Spec *struct {
								Destination *struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

									Server *string `tfsdk:"server" yaml:"server,omitempty"`
								} `tfsdk:"destination" yaml:"destination,omitempty"`

								IgnoreDifferences *[]struct {
									Group *string `tfsdk:"group" yaml:"group,omitempty"`

									JqPathExpressions *[]string `tfsdk:"jq_path_expressions" yaml:"jqPathExpressions,omitempty"`

									JsonPointers *[]string `tfsdk:"json_pointers" yaml:"jsonPointers,omitempty"`

									Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

									ManagedFieldsManagers *[]string `tfsdk:"managed_fields_managers" yaml:"managedFieldsManagers,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
								} `tfsdk:"ignore_differences" yaml:"ignoreDifferences,omitempty"`

								Info *[]struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"info" yaml:"info,omitempty"`

								Project *string `tfsdk:"project" yaml:"project,omitempty"`

								RevisionHistoryLimit *int64 `tfsdk:"revision_history_limit" yaml:"revisionHistoryLimit,omitempty"`

								Source *struct {
									Chart *string `tfsdk:"chart" yaml:"chart,omitempty"`

									Directory *struct {
										Exclude *string `tfsdk:"exclude" yaml:"exclude,omitempty"`

										Include *string `tfsdk:"include" yaml:"include,omitempty"`

										Jsonnet *struct {
											ExtVars *[]struct {
												Code *bool `tfsdk:"code" yaml:"code,omitempty"`

												Name *string `tfsdk:"name" yaml:"name,omitempty"`

												Value *string `tfsdk:"value" yaml:"value,omitempty"`
											} `tfsdk:"ext_vars" yaml:"extVars,omitempty"`

											Libs *[]string `tfsdk:"libs" yaml:"libs,omitempty"`

											Tlas *[]struct {
												Code *bool `tfsdk:"code" yaml:"code,omitempty"`

												Name *string `tfsdk:"name" yaml:"name,omitempty"`

												Value *string `tfsdk:"value" yaml:"value,omitempty"`
											} `tfsdk:"tlas" yaml:"tlas,omitempty"`
										} `tfsdk:"jsonnet" yaml:"jsonnet,omitempty"`

										Recurse *bool `tfsdk:"recurse" yaml:"recurse,omitempty"`
									} `tfsdk:"directory" yaml:"directory,omitempty"`

									Helm *struct {
										FileParameters *[]struct {
											Name *string `tfsdk:"name" yaml:"name,omitempty"`

											Path *string `tfsdk:"path" yaml:"path,omitempty"`
										} `tfsdk:"file_parameters" yaml:"fileParameters,omitempty"`

										IgnoreMissingValueFiles *bool `tfsdk:"ignore_missing_value_files" yaml:"ignoreMissingValueFiles,omitempty"`

										Parameters *[]struct {
											ForceString *bool `tfsdk:"force_string" yaml:"forceString,omitempty"`

											Name *string `tfsdk:"name" yaml:"name,omitempty"`

											Value *string `tfsdk:"value" yaml:"value,omitempty"`
										} `tfsdk:"parameters" yaml:"parameters,omitempty"`

										PassCredentials *bool `tfsdk:"pass_credentials" yaml:"passCredentials,omitempty"`

										ReleaseName *string `tfsdk:"release_name" yaml:"releaseName,omitempty"`

										SkipCrds *bool `tfsdk:"skip_crds" yaml:"skipCrds,omitempty"`

										ValueFiles *[]string `tfsdk:"value_files" yaml:"valueFiles,omitempty"`

										Values *string `tfsdk:"values" yaml:"values,omitempty"`

										Version *string `tfsdk:"version" yaml:"version,omitempty"`
									} `tfsdk:"helm" yaml:"helm,omitempty"`

									Kustomize *struct {
										CommonAnnotations *map[string]string `tfsdk:"common_annotations" yaml:"commonAnnotations,omitempty"`

										CommonLabels *map[string]string `tfsdk:"common_labels" yaml:"commonLabels,omitempty"`

										ForceCommonAnnotations *bool `tfsdk:"force_common_annotations" yaml:"forceCommonAnnotations,omitempty"`

										ForceCommonLabels *bool `tfsdk:"force_common_labels" yaml:"forceCommonLabels,omitempty"`

										Images *[]string `tfsdk:"images" yaml:"images,omitempty"`

										NamePrefix *string `tfsdk:"name_prefix" yaml:"namePrefix,omitempty"`

										NameSuffix *string `tfsdk:"name_suffix" yaml:"nameSuffix,omitempty"`

										Version *string `tfsdk:"version" yaml:"version,omitempty"`
									} `tfsdk:"kustomize" yaml:"kustomize,omitempty"`

									Path *string `tfsdk:"path" yaml:"path,omitempty"`

									Plugin *struct {
										Env *[]struct {
											Name *string `tfsdk:"name" yaml:"name,omitempty"`

											Value *string `tfsdk:"value" yaml:"value,omitempty"`
										} `tfsdk:"env" yaml:"env,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"plugin" yaml:"plugin,omitempty"`

									RepoURL *string `tfsdk:"repo_url" yaml:"repoURL,omitempty"`

									TargetRevision *string `tfsdk:"target_revision" yaml:"targetRevision,omitempty"`
								} `tfsdk:"source" yaml:"source,omitempty"`

								SyncPolicy *struct {
									Automated *struct {
										AllowEmpty *bool `tfsdk:"allow_empty" yaml:"allowEmpty,omitempty"`

										Prune *bool `tfsdk:"prune" yaml:"prune,omitempty"`

										SelfHeal *bool `tfsdk:"self_heal" yaml:"selfHeal,omitempty"`
									} `tfsdk:"automated" yaml:"automated,omitempty"`

									Retry *struct {
										Backoff *struct {
											Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

											Factor *int64 `tfsdk:"factor" yaml:"factor,omitempty"`

											MaxDuration *string `tfsdk:"max_duration" yaml:"maxDuration,omitempty"`
										} `tfsdk:"backoff" yaml:"backoff,omitempty"`

										Limit *int64 `tfsdk:"limit" yaml:"limit,omitempty"`
									} `tfsdk:"retry" yaml:"retry,omitempty"`

									SyncOptions *[]string `tfsdk:"sync_options" yaml:"syncOptions,omitempty"`
								} `tfsdk:"sync_policy" yaml:"syncPolicy,omitempty"`
							} `tfsdk:"spec" yaml:"spec,omitempty"`
						} `tfsdk:"template" yaml:"template,omitempty"`
					} `tfsdk:"scm_provider" yaml:"scmProvider,omitempty"`

					Selector *struct {
						MatchExpressions *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
						} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

						MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
					} `tfsdk:"selector" yaml:"selector,omitempty"`
				} `tfsdk:"generators" yaml:"generators,omitempty"`

				Template *struct {
					Metadata *struct {
						Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Finalizers *[]string `tfsdk:"finalizers" yaml:"finalizers,omitempty"`

						Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`

					Spec *struct {
						Destination *struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

							Server *string `tfsdk:"server" yaml:"server,omitempty"`
						} `tfsdk:"destination" yaml:"destination,omitempty"`

						IgnoreDifferences *[]struct {
							Group *string `tfsdk:"group" yaml:"group,omitempty"`

							JqPathExpressions *[]string `tfsdk:"jq_path_expressions" yaml:"jqPathExpressions,omitempty"`

							JsonPointers *[]string `tfsdk:"json_pointers" yaml:"jsonPointers,omitempty"`

							Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

							ManagedFieldsManagers *[]string `tfsdk:"managed_fields_managers" yaml:"managedFieldsManagers,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
						} `tfsdk:"ignore_differences" yaml:"ignoreDifferences,omitempty"`

						Info *[]struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"info" yaml:"info,omitempty"`

						Project *string `tfsdk:"project" yaml:"project,omitempty"`

						RevisionHistoryLimit *int64 `tfsdk:"revision_history_limit" yaml:"revisionHistoryLimit,omitempty"`

						Source *struct {
							Chart *string `tfsdk:"chart" yaml:"chart,omitempty"`

							Directory *struct {
								Exclude *string `tfsdk:"exclude" yaml:"exclude,omitempty"`

								Include *string `tfsdk:"include" yaml:"include,omitempty"`

								Jsonnet *struct {
									ExtVars *[]struct {
										Code *bool `tfsdk:"code" yaml:"code,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`

										Value *string `tfsdk:"value" yaml:"value,omitempty"`
									} `tfsdk:"ext_vars" yaml:"extVars,omitempty"`

									Libs *[]string `tfsdk:"libs" yaml:"libs,omitempty"`

									Tlas *[]struct {
										Code *bool `tfsdk:"code" yaml:"code,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`

										Value *string `tfsdk:"value" yaml:"value,omitempty"`
									} `tfsdk:"tlas" yaml:"tlas,omitempty"`
								} `tfsdk:"jsonnet" yaml:"jsonnet,omitempty"`

								Recurse *bool `tfsdk:"recurse" yaml:"recurse,omitempty"`
							} `tfsdk:"directory" yaml:"directory,omitempty"`

							Helm *struct {
								FileParameters *[]struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Path *string `tfsdk:"path" yaml:"path,omitempty"`
								} `tfsdk:"file_parameters" yaml:"fileParameters,omitempty"`

								IgnoreMissingValueFiles *bool `tfsdk:"ignore_missing_value_files" yaml:"ignoreMissingValueFiles,omitempty"`

								Parameters *[]struct {
									ForceString *bool `tfsdk:"force_string" yaml:"forceString,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"parameters" yaml:"parameters,omitempty"`

								PassCredentials *bool `tfsdk:"pass_credentials" yaml:"passCredentials,omitempty"`

								ReleaseName *string `tfsdk:"release_name" yaml:"releaseName,omitempty"`

								SkipCrds *bool `tfsdk:"skip_crds" yaml:"skipCrds,omitempty"`

								ValueFiles *[]string `tfsdk:"value_files" yaml:"valueFiles,omitempty"`

								Values *string `tfsdk:"values" yaml:"values,omitempty"`

								Version *string `tfsdk:"version" yaml:"version,omitempty"`
							} `tfsdk:"helm" yaml:"helm,omitempty"`

							Kustomize *struct {
								CommonAnnotations *map[string]string `tfsdk:"common_annotations" yaml:"commonAnnotations,omitempty"`

								CommonLabels *map[string]string `tfsdk:"common_labels" yaml:"commonLabels,omitempty"`

								ForceCommonAnnotations *bool `tfsdk:"force_common_annotations" yaml:"forceCommonAnnotations,omitempty"`

								ForceCommonLabels *bool `tfsdk:"force_common_labels" yaml:"forceCommonLabels,omitempty"`

								Images *[]string `tfsdk:"images" yaml:"images,omitempty"`

								NamePrefix *string `tfsdk:"name_prefix" yaml:"namePrefix,omitempty"`

								NameSuffix *string `tfsdk:"name_suffix" yaml:"nameSuffix,omitempty"`

								Version *string `tfsdk:"version" yaml:"version,omitempty"`
							} `tfsdk:"kustomize" yaml:"kustomize,omitempty"`

							Path *string `tfsdk:"path" yaml:"path,omitempty"`

							Plugin *struct {
								Env *[]struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"env" yaml:"env,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`
							} `tfsdk:"plugin" yaml:"plugin,omitempty"`

							RepoURL *string `tfsdk:"repo_url" yaml:"repoURL,omitempty"`

							TargetRevision *string `tfsdk:"target_revision" yaml:"targetRevision,omitempty"`
						} `tfsdk:"source" yaml:"source,omitempty"`

						SyncPolicy *struct {
							Automated *struct {
								AllowEmpty *bool `tfsdk:"allow_empty" yaml:"allowEmpty,omitempty"`

								Prune *bool `tfsdk:"prune" yaml:"prune,omitempty"`

								SelfHeal *bool `tfsdk:"self_heal" yaml:"selfHeal,omitempty"`
							} `tfsdk:"automated" yaml:"automated,omitempty"`

							Retry *struct {
								Backoff *struct {
									Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

									Factor *int64 `tfsdk:"factor" yaml:"factor,omitempty"`

									MaxDuration *string `tfsdk:"max_duration" yaml:"maxDuration,omitempty"`
								} `tfsdk:"backoff" yaml:"backoff,omitempty"`

								Limit *int64 `tfsdk:"limit" yaml:"limit,omitempty"`
							} `tfsdk:"retry" yaml:"retry,omitempty"`

							SyncOptions *[]string `tfsdk:"sync_options" yaml:"syncOptions,omitempty"`
						} `tfsdk:"sync_policy" yaml:"syncPolicy,omitempty"`
					} `tfsdk:"spec" yaml:"spec,omitempty"`
				} `tfsdk:"template" yaml:"template,omitempty"`
			} `tfsdk:"matrix" yaml:"matrix,omitempty"`

			Merge *struct {
				Generators *[]struct {
					ClusterDecisionResource *struct {
						ConfigMapRef *string `tfsdk:"config_map_ref" yaml:"configMapRef,omitempty"`

						LabelSelector *struct {
							MatchExpressions *[]struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

								Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
							} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

							MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
						} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						RequeueAfterSeconds *int64 `tfsdk:"requeue_after_seconds" yaml:"requeueAfterSeconds,omitempty"`

						Template *struct {
							Metadata *struct {
								Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

								Finalizers *[]string `tfsdk:"finalizers" yaml:"finalizers,omitempty"`

								Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
							} `tfsdk:"metadata" yaml:"metadata,omitempty"`

							Spec *struct {
								Destination *struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

									Server *string `tfsdk:"server" yaml:"server,omitempty"`
								} `tfsdk:"destination" yaml:"destination,omitempty"`

								IgnoreDifferences *[]struct {
									Group *string `tfsdk:"group" yaml:"group,omitempty"`

									JqPathExpressions *[]string `tfsdk:"jq_path_expressions" yaml:"jqPathExpressions,omitempty"`

									JsonPointers *[]string `tfsdk:"json_pointers" yaml:"jsonPointers,omitempty"`

									Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

									ManagedFieldsManagers *[]string `tfsdk:"managed_fields_managers" yaml:"managedFieldsManagers,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
								} `tfsdk:"ignore_differences" yaml:"ignoreDifferences,omitempty"`

								Info *[]struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"info" yaml:"info,omitempty"`

								Project *string `tfsdk:"project" yaml:"project,omitempty"`

								RevisionHistoryLimit *int64 `tfsdk:"revision_history_limit" yaml:"revisionHistoryLimit,omitempty"`

								Source *struct {
									Chart *string `tfsdk:"chart" yaml:"chart,omitempty"`

									Directory *struct {
										Exclude *string `tfsdk:"exclude" yaml:"exclude,omitempty"`

										Include *string `tfsdk:"include" yaml:"include,omitempty"`

										Jsonnet *struct {
											ExtVars *[]struct {
												Code *bool `tfsdk:"code" yaml:"code,omitempty"`

												Name *string `tfsdk:"name" yaml:"name,omitempty"`

												Value *string `tfsdk:"value" yaml:"value,omitempty"`
											} `tfsdk:"ext_vars" yaml:"extVars,omitempty"`

											Libs *[]string `tfsdk:"libs" yaml:"libs,omitempty"`

											Tlas *[]struct {
												Code *bool `tfsdk:"code" yaml:"code,omitempty"`

												Name *string `tfsdk:"name" yaml:"name,omitempty"`

												Value *string `tfsdk:"value" yaml:"value,omitempty"`
											} `tfsdk:"tlas" yaml:"tlas,omitempty"`
										} `tfsdk:"jsonnet" yaml:"jsonnet,omitempty"`

										Recurse *bool `tfsdk:"recurse" yaml:"recurse,omitempty"`
									} `tfsdk:"directory" yaml:"directory,omitempty"`

									Helm *struct {
										FileParameters *[]struct {
											Name *string `tfsdk:"name" yaml:"name,omitempty"`

											Path *string `tfsdk:"path" yaml:"path,omitempty"`
										} `tfsdk:"file_parameters" yaml:"fileParameters,omitempty"`

										IgnoreMissingValueFiles *bool `tfsdk:"ignore_missing_value_files" yaml:"ignoreMissingValueFiles,omitempty"`

										Parameters *[]struct {
											ForceString *bool `tfsdk:"force_string" yaml:"forceString,omitempty"`

											Name *string `tfsdk:"name" yaml:"name,omitempty"`

											Value *string `tfsdk:"value" yaml:"value,omitempty"`
										} `tfsdk:"parameters" yaml:"parameters,omitempty"`

										PassCredentials *bool `tfsdk:"pass_credentials" yaml:"passCredentials,omitempty"`

										ReleaseName *string `tfsdk:"release_name" yaml:"releaseName,omitempty"`

										SkipCrds *bool `tfsdk:"skip_crds" yaml:"skipCrds,omitempty"`

										ValueFiles *[]string `tfsdk:"value_files" yaml:"valueFiles,omitempty"`

										Values *string `tfsdk:"values" yaml:"values,omitempty"`

										Version *string `tfsdk:"version" yaml:"version,omitempty"`
									} `tfsdk:"helm" yaml:"helm,omitempty"`

									Kustomize *struct {
										CommonAnnotations *map[string]string `tfsdk:"common_annotations" yaml:"commonAnnotations,omitempty"`

										CommonLabels *map[string]string `tfsdk:"common_labels" yaml:"commonLabels,omitempty"`

										ForceCommonAnnotations *bool `tfsdk:"force_common_annotations" yaml:"forceCommonAnnotations,omitempty"`

										ForceCommonLabels *bool `tfsdk:"force_common_labels" yaml:"forceCommonLabels,omitempty"`

										Images *[]string `tfsdk:"images" yaml:"images,omitempty"`

										NamePrefix *string `tfsdk:"name_prefix" yaml:"namePrefix,omitempty"`

										NameSuffix *string `tfsdk:"name_suffix" yaml:"nameSuffix,omitempty"`

										Version *string `tfsdk:"version" yaml:"version,omitempty"`
									} `tfsdk:"kustomize" yaml:"kustomize,omitempty"`

									Path *string `tfsdk:"path" yaml:"path,omitempty"`

									Plugin *struct {
										Env *[]struct {
											Name *string `tfsdk:"name" yaml:"name,omitempty"`

											Value *string `tfsdk:"value" yaml:"value,omitempty"`
										} `tfsdk:"env" yaml:"env,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"plugin" yaml:"plugin,omitempty"`

									RepoURL *string `tfsdk:"repo_url" yaml:"repoURL,omitempty"`

									TargetRevision *string `tfsdk:"target_revision" yaml:"targetRevision,omitempty"`
								} `tfsdk:"source" yaml:"source,omitempty"`

								SyncPolicy *struct {
									Automated *struct {
										AllowEmpty *bool `tfsdk:"allow_empty" yaml:"allowEmpty,omitempty"`

										Prune *bool `tfsdk:"prune" yaml:"prune,omitempty"`

										SelfHeal *bool `tfsdk:"self_heal" yaml:"selfHeal,omitempty"`
									} `tfsdk:"automated" yaml:"automated,omitempty"`

									Retry *struct {
										Backoff *struct {
											Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

											Factor *int64 `tfsdk:"factor" yaml:"factor,omitempty"`

											MaxDuration *string `tfsdk:"max_duration" yaml:"maxDuration,omitempty"`
										} `tfsdk:"backoff" yaml:"backoff,omitempty"`

										Limit *int64 `tfsdk:"limit" yaml:"limit,omitempty"`
									} `tfsdk:"retry" yaml:"retry,omitempty"`

									SyncOptions *[]string `tfsdk:"sync_options" yaml:"syncOptions,omitempty"`
								} `tfsdk:"sync_policy" yaml:"syncPolicy,omitempty"`
							} `tfsdk:"spec" yaml:"spec,omitempty"`
						} `tfsdk:"template" yaml:"template,omitempty"`

						Values *map[string]string `tfsdk:"values" yaml:"values,omitempty"`
					} `tfsdk:"cluster_decision_resource" yaml:"clusterDecisionResource,omitempty"`

					Clusters *struct {
						Selector *struct {
							MatchExpressions *[]struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

								Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
							} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

							MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
						} `tfsdk:"selector" yaml:"selector,omitempty"`

						Template *struct {
							Metadata *struct {
								Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

								Finalizers *[]string `tfsdk:"finalizers" yaml:"finalizers,omitempty"`

								Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
							} `tfsdk:"metadata" yaml:"metadata,omitempty"`

							Spec *struct {
								Destination *struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

									Server *string `tfsdk:"server" yaml:"server,omitempty"`
								} `tfsdk:"destination" yaml:"destination,omitempty"`

								IgnoreDifferences *[]struct {
									Group *string `tfsdk:"group" yaml:"group,omitempty"`

									JqPathExpressions *[]string `tfsdk:"jq_path_expressions" yaml:"jqPathExpressions,omitempty"`

									JsonPointers *[]string `tfsdk:"json_pointers" yaml:"jsonPointers,omitempty"`

									Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

									ManagedFieldsManagers *[]string `tfsdk:"managed_fields_managers" yaml:"managedFieldsManagers,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
								} `tfsdk:"ignore_differences" yaml:"ignoreDifferences,omitempty"`

								Info *[]struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"info" yaml:"info,omitempty"`

								Project *string `tfsdk:"project" yaml:"project,omitempty"`

								RevisionHistoryLimit *int64 `tfsdk:"revision_history_limit" yaml:"revisionHistoryLimit,omitempty"`

								Source *struct {
									Chart *string `tfsdk:"chart" yaml:"chart,omitempty"`

									Directory *struct {
										Exclude *string `tfsdk:"exclude" yaml:"exclude,omitempty"`

										Include *string `tfsdk:"include" yaml:"include,omitempty"`

										Jsonnet *struct {
											ExtVars *[]struct {
												Code *bool `tfsdk:"code" yaml:"code,omitempty"`

												Name *string `tfsdk:"name" yaml:"name,omitempty"`

												Value *string `tfsdk:"value" yaml:"value,omitempty"`
											} `tfsdk:"ext_vars" yaml:"extVars,omitempty"`

											Libs *[]string `tfsdk:"libs" yaml:"libs,omitempty"`

											Tlas *[]struct {
												Code *bool `tfsdk:"code" yaml:"code,omitempty"`

												Name *string `tfsdk:"name" yaml:"name,omitempty"`

												Value *string `tfsdk:"value" yaml:"value,omitempty"`
											} `tfsdk:"tlas" yaml:"tlas,omitempty"`
										} `tfsdk:"jsonnet" yaml:"jsonnet,omitempty"`

										Recurse *bool `tfsdk:"recurse" yaml:"recurse,omitempty"`
									} `tfsdk:"directory" yaml:"directory,omitempty"`

									Helm *struct {
										FileParameters *[]struct {
											Name *string `tfsdk:"name" yaml:"name,omitempty"`

											Path *string `tfsdk:"path" yaml:"path,omitempty"`
										} `tfsdk:"file_parameters" yaml:"fileParameters,omitempty"`

										IgnoreMissingValueFiles *bool `tfsdk:"ignore_missing_value_files" yaml:"ignoreMissingValueFiles,omitempty"`

										Parameters *[]struct {
											ForceString *bool `tfsdk:"force_string" yaml:"forceString,omitempty"`

											Name *string `tfsdk:"name" yaml:"name,omitempty"`

											Value *string `tfsdk:"value" yaml:"value,omitempty"`
										} `tfsdk:"parameters" yaml:"parameters,omitempty"`

										PassCredentials *bool `tfsdk:"pass_credentials" yaml:"passCredentials,omitempty"`

										ReleaseName *string `tfsdk:"release_name" yaml:"releaseName,omitempty"`

										SkipCrds *bool `tfsdk:"skip_crds" yaml:"skipCrds,omitempty"`

										ValueFiles *[]string `tfsdk:"value_files" yaml:"valueFiles,omitempty"`

										Values *string `tfsdk:"values" yaml:"values,omitempty"`

										Version *string `tfsdk:"version" yaml:"version,omitempty"`
									} `tfsdk:"helm" yaml:"helm,omitempty"`

									Kustomize *struct {
										CommonAnnotations *map[string]string `tfsdk:"common_annotations" yaml:"commonAnnotations,omitempty"`

										CommonLabels *map[string]string `tfsdk:"common_labels" yaml:"commonLabels,omitempty"`

										ForceCommonAnnotations *bool `tfsdk:"force_common_annotations" yaml:"forceCommonAnnotations,omitempty"`

										ForceCommonLabels *bool `tfsdk:"force_common_labels" yaml:"forceCommonLabels,omitempty"`

										Images *[]string `tfsdk:"images" yaml:"images,omitempty"`

										NamePrefix *string `tfsdk:"name_prefix" yaml:"namePrefix,omitempty"`

										NameSuffix *string `tfsdk:"name_suffix" yaml:"nameSuffix,omitempty"`

										Version *string `tfsdk:"version" yaml:"version,omitempty"`
									} `tfsdk:"kustomize" yaml:"kustomize,omitempty"`

									Path *string `tfsdk:"path" yaml:"path,omitempty"`

									Plugin *struct {
										Env *[]struct {
											Name *string `tfsdk:"name" yaml:"name,omitempty"`

											Value *string `tfsdk:"value" yaml:"value,omitempty"`
										} `tfsdk:"env" yaml:"env,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"plugin" yaml:"plugin,omitempty"`

									RepoURL *string `tfsdk:"repo_url" yaml:"repoURL,omitempty"`

									TargetRevision *string `tfsdk:"target_revision" yaml:"targetRevision,omitempty"`
								} `tfsdk:"source" yaml:"source,omitempty"`

								SyncPolicy *struct {
									Automated *struct {
										AllowEmpty *bool `tfsdk:"allow_empty" yaml:"allowEmpty,omitempty"`

										Prune *bool `tfsdk:"prune" yaml:"prune,omitempty"`

										SelfHeal *bool `tfsdk:"self_heal" yaml:"selfHeal,omitempty"`
									} `tfsdk:"automated" yaml:"automated,omitempty"`

									Retry *struct {
										Backoff *struct {
											Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

											Factor *int64 `tfsdk:"factor" yaml:"factor,omitempty"`

											MaxDuration *string `tfsdk:"max_duration" yaml:"maxDuration,omitempty"`
										} `tfsdk:"backoff" yaml:"backoff,omitempty"`

										Limit *int64 `tfsdk:"limit" yaml:"limit,omitempty"`
									} `tfsdk:"retry" yaml:"retry,omitempty"`

									SyncOptions *[]string `tfsdk:"sync_options" yaml:"syncOptions,omitempty"`
								} `tfsdk:"sync_policy" yaml:"syncPolicy,omitempty"`
							} `tfsdk:"spec" yaml:"spec,omitempty"`
						} `tfsdk:"template" yaml:"template,omitempty"`

						Values *map[string]string `tfsdk:"values" yaml:"values,omitempty"`
					} `tfsdk:"clusters" yaml:"clusters,omitempty"`

					Git *struct {
						Directories *[]struct {
							Exclude *bool `tfsdk:"exclude" yaml:"exclude,omitempty"`

							Path *string `tfsdk:"path" yaml:"path,omitempty"`
						} `tfsdk:"directories" yaml:"directories,omitempty"`

						Files *[]struct {
							Path *string `tfsdk:"path" yaml:"path,omitempty"`
						} `tfsdk:"files" yaml:"files,omitempty"`

						RepoURL *string `tfsdk:"repo_url" yaml:"repoURL,omitempty"`

						RequeueAfterSeconds *int64 `tfsdk:"requeue_after_seconds" yaml:"requeueAfterSeconds,omitempty"`

						Revision *string `tfsdk:"revision" yaml:"revision,omitempty"`

						Template *struct {
							Metadata *struct {
								Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

								Finalizers *[]string `tfsdk:"finalizers" yaml:"finalizers,omitempty"`

								Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
							} `tfsdk:"metadata" yaml:"metadata,omitempty"`

							Spec *struct {
								Destination *struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

									Server *string `tfsdk:"server" yaml:"server,omitempty"`
								} `tfsdk:"destination" yaml:"destination,omitempty"`

								IgnoreDifferences *[]struct {
									Group *string `tfsdk:"group" yaml:"group,omitempty"`

									JqPathExpressions *[]string `tfsdk:"jq_path_expressions" yaml:"jqPathExpressions,omitempty"`

									JsonPointers *[]string `tfsdk:"json_pointers" yaml:"jsonPointers,omitempty"`

									Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

									ManagedFieldsManagers *[]string `tfsdk:"managed_fields_managers" yaml:"managedFieldsManagers,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
								} `tfsdk:"ignore_differences" yaml:"ignoreDifferences,omitempty"`

								Info *[]struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"info" yaml:"info,omitempty"`

								Project *string `tfsdk:"project" yaml:"project,omitempty"`

								RevisionHistoryLimit *int64 `tfsdk:"revision_history_limit" yaml:"revisionHistoryLimit,omitempty"`

								Source *struct {
									Chart *string `tfsdk:"chart" yaml:"chart,omitempty"`

									Directory *struct {
										Exclude *string `tfsdk:"exclude" yaml:"exclude,omitempty"`

										Include *string `tfsdk:"include" yaml:"include,omitempty"`

										Jsonnet *struct {
											ExtVars *[]struct {
												Code *bool `tfsdk:"code" yaml:"code,omitempty"`

												Name *string `tfsdk:"name" yaml:"name,omitempty"`

												Value *string `tfsdk:"value" yaml:"value,omitempty"`
											} `tfsdk:"ext_vars" yaml:"extVars,omitempty"`

											Libs *[]string `tfsdk:"libs" yaml:"libs,omitempty"`

											Tlas *[]struct {
												Code *bool `tfsdk:"code" yaml:"code,omitempty"`

												Name *string `tfsdk:"name" yaml:"name,omitempty"`

												Value *string `tfsdk:"value" yaml:"value,omitempty"`
											} `tfsdk:"tlas" yaml:"tlas,omitempty"`
										} `tfsdk:"jsonnet" yaml:"jsonnet,omitempty"`

										Recurse *bool `tfsdk:"recurse" yaml:"recurse,omitempty"`
									} `tfsdk:"directory" yaml:"directory,omitempty"`

									Helm *struct {
										FileParameters *[]struct {
											Name *string `tfsdk:"name" yaml:"name,omitempty"`

											Path *string `tfsdk:"path" yaml:"path,omitempty"`
										} `tfsdk:"file_parameters" yaml:"fileParameters,omitempty"`

										IgnoreMissingValueFiles *bool `tfsdk:"ignore_missing_value_files" yaml:"ignoreMissingValueFiles,omitempty"`

										Parameters *[]struct {
											ForceString *bool `tfsdk:"force_string" yaml:"forceString,omitempty"`

											Name *string `tfsdk:"name" yaml:"name,omitempty"`

											Value *string `tfsdk:"value" yaml:"value,omitempty"`
										} `tfsdk:"parameters" yaml:"parameters,omitempty"`

										PassCredentials *bool `tfsdk:"pass_credentials" yaml:"passCredentials,omitempty"`

										ReleaseName *string `tfsdk:"release_name" yaml:"releaseName,omitempty"`

										SkipCrds *bool `tfsdk:"skip_crds" yaml:"skipCrds,omitempty"`

										ValueFiles *[]string `tfsdk:"value_files" yaml:"valueFiles,omitempty"`

										Values *string `tfsdk:"values" yaml:"values,omitempty"`

										Version *string `tfsdk:"version" yaml:"version,omitempty"`
									} `tfsdk:"helm" yaml:"helm,omitempty"`

									Kustomize *struct {
										CommonAnnotations *map[string]string `tfsdk:"common_annotations" yaml:"commonAnnotations,omitempty"`

										CommonLabels *map[string]string `tfsdk:"common_labels" yaml:"commonLabels,omitempty"`

										ForceCommonAnnotations *bool `tfsdk:"force_common_annotations" yaml:"forceCommonAnnotations,omitempty"`

										ForceCommonLabels *bool `tfsdk:"force_common_labels" yaml:"forceCommonLabels,omitempty"`

										Images *[]string `tfsdk:"images" yaml:"images,omitempty"`

										NamePrefix *string `tfsdk:"name_prefix" yaml:"namePrefix,omitempty"`

										NameSuffix *string `tfsdk:"name_suffix" yaml:"nameSuffix,omitempty"`

										Version *string `tfsdk:"version" yaml:"version,omitempty"`
									} `tfsdk:"kustomize" yaml:"kustomize,omitempty"`

									Path *string `tfsdk:"path" yaml:"path,omitempty"`

									Plugin *struct {
										Env *[]struct {
											Name *string `tfsdk:"name" yaml:"name,omitempty"`

											Value *string `tfsdk:"value" yaml:"value,omitempty"`
										} `tfsdk:"env" yaml:"env,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"plugin" yaml:"plugin,omitempty"`

									RepoURL *string `tfsdk:"repo_url" yaml:"repoURL,omitempty"`

									TargetRevision *string `tfsdk:"target_revision" yaml:"targetRevision,omitempty"`
								} `tfsdk:"source" yaml:"source,omitempty"`

								SyncPolicy *struct {
									Automated *struct {
										AllowEmpty *bool `tfsdk:"allow_empty" yaml:"allowEmpty,omitempty"`

										Prune *bool `tfsdk:"prune" yaml:"prune,omitempty"`

										SelfHeal *bool `tfsdk:"self_heal" yaml:"selfHeal,omitempty"`
									} `tfsdk:"automated" yaml:"automated,omitempty"`

									Retry *struct {
										Backoff *struct {
											Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

											Factor *int64 `tfsdk:"factor" yaml:"factor,omitempty"`

											MaxDuration *string `tfsdk:"max_duration" yaml:"maxDuration,omitempty"`
										} `tfsdk:"backoff" yaml:"backoff,omitempty"`

										Limit *int64 `tfsdk:"limit" yaml:"limit,omitempty"`
									} `tfsdk:"retry" yaml:"retry,omitempty"`

									SyncOptions *[]string `tfsdk:"sync_options" yaml:"syncOptions,omitempty"`
								} `tfsdk:"sync_policy" yaml:"syncPolicy,omitempty"`
							} `tfsdk:"spec" yaml:"spec,omitempty"`
						} `tfsdk:"template" yaml:"template,omitempty"`
					} `tfsdk:"git" yaml:"git,omitempty"`

					List *struct {
						Elements *[]string `tfsdk:"elements" yaml:"elements,omitempty"`

						Template *struct {
							Metadata *struct {
								Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

								Finalizers *[]string `tfsdk:"finalizers" yaml:"finalizers,omitempty"`

								Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
							} `tfsdk:"metadata" yaml:"metadata,omitempty"`

							Spec *struct {
								Destination *struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

									Server *string `tfsdk:"server" yaml:"server,omitempty"`
								} `tfsdk:"destination" yaml:"destination,omitempty"`

								IgnoreDifferences *[]struct {
									Group *string `tfsdk:"group" yaml:"group,omitempty"`

									JqPathExpressions *[]string `tfsdk:"jq_path_expressions" yaml:"jqPathExpressions,omitempty"`

									JsonPointers *[]string `tfsdk:"json_pointers" yaml:"jsonPointers,omitempty"`

									Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

									ManagedFieldsManagers *[]string `tfsdk:"managed_fields_managers" yaml:"managedFieldsManagers,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
								} `tfsdk:"ignore_differences" yaml:"ignoreDifferences,omitempty"`

								Info *[]struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"info" yaml:"info,omitempty"`

								Project *string `tfsdk:"project" yaml:"project,omitempty"`

								RevisionHistoryLimit *int64 `tfsdk:"revision_history_limit" yaml:"revisionHistoryLimit,omitempty"`

								Source *struct {
									Chart *string `tfsdk:"chart" yaml:"chart,omitempty"`

									Directory *struct {
										Exclude *string `tfsdk:"exclude" yaml:"exclude,omitempty"`

										Include *string `tfsdk:"include" yaml:"include,omitempty"`

										Jsonnet *struct {
											ExtVars *[]struct {
												Code *bool `tfsdk:"code" yaml:"code,omitempty"`

												Name *string `tfsdk:"name" yaml:"name,omitempty"`

												Value *string `tfsdk:"value" yaml:"value,omitempty"`
											} `tfsdk:"ext_vars" yaml:"extVars,omitempty"`

											Libs *[]string `tfsdk:"libs" yaml:"libs,omitempty"`

											Tlas *[]struct {
												Code *bool `tfsdk:"code" yaml:"code,omitempty"`

												Name *string `tfsdk:"name" yaml:"name,omitempty"`

												Value *string `tfsdk:"value" yaml:"value,omitempty"`
											} `tfsdk:"tlas" yaml:"tlas,omitempty"`
										} `tfsdk:"jsonnet" yaml:"jsonnet,omitempty"`

										Recurse *bool `tfsdk:"recurse" yaml:"recurse,omitempty"`
									} `tfsdk:"directory" yaml:"directory,omitempty"`

									Helm *struct {
										FileParameters *[]struct {
											Name *string `tfsdk:"name" yaml:"name,omitempty"`

											Path *string `tfsdk:"path" yaml:"path,omitempty"`
										} `tfsdk:"file_parameters" yaml:"fileParameters,omitempty"`

										IgnoreMissingValueFiles *bool `tfsdk:"ignore_missing_value_files" yaml:"ignoreMissingValueFiles,omitempty"`

										Parameters *[]struct {
											ForceString *bool `tfsdk:"force_string" yaml:"forceString,omitempty"`

											Name *string `tfsdk:"name" yaml:"name,omitempty"`

											Value *string `tfsdk:"value" yaml:"value,omitempty"`
										} `tfsdk:"parameters" yaml:"parameters,omitempty"`

										PassCredentials *bool `tfsdk:"pass_credentials" yaml:"passCredentials,omitempty"`

										ReleaseName *string `tfsdk:"release_name" yaml:"releaseName,omitempty"`

										SkipCrds *bool `tfsdk:"skip_crds" yaml:"skipCrds,omitempty"`

										ValueFiles *[]string `tfsdk:"value_files" yaml:"valueFiles,omitempty"`

										Values *string `tfsdk:"values" yaml:"values,omitempty"`

										Version *string `tfsdk:"version" yaml:"version,omitempty"`
									} `tfsdk:"helm" yaml:"helm,omitempty"`

									Kustomize *struct {
										CommonAnnotations *map[string]string `tfsdk:"common_annotations" yaml:"commonAnnotations,omitempty"`

										CommonLabels *map[string]string `tfsdk:"common_labels" yaml:"commonLabels,omitempty"`

										ForceCommonAnnotations *bool `tfsdk:"force_common_annotations" yaml:"forceCommonAnnotations,omitempty"`

										ForceCommonLabels *bool `tfsdk:"force_common_labels" yaml:"forceCommonLabels,omitempty"`

										Images *[]string `tfsdk:"images" yaml:"images,omitempty"`

										NamePrefix *string `tfsdk:"name_prefix" yaml:"namePrefix,omitempty"`

										NameSuffix *string `tfsdk:"name_suffix" yaml:"nameSuffix,omitempty"`

										Version *string `tfsdk:"version" yaml:"version,omitempty"`
									} `tfsdk:"kustomize" yaml:"kustomize,omitempty"`

									Path *string `tfsdk:"path" yaml:"path,omitempty"`

									Plugin *struct {
										Env *[]struct {
											Name *string `tfsdk:"name" yaml:"name,omitempty"`

											Value *string `tfsdk:"value" yaml:"value,omitempty"`
										} `tfsdk:"env" yaml:"env,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"plugin" yaml:"plugin,omitempty"`

									RepoURL *string `tfsdk:"repo_url" yaml:"repoURL,omitempty"`

									TargetRevision *string `tfsdk:"target_revision" yaml:"targetRevision,omitempty"`
								} `tfsdk:"source" yaml:"source,omitempty"`

								SyncPolicy *struct {
									Automated *struct {
										AllowEmpty *bool `tfsdk:"allow_empty" yaml:"allowEmpty,omitempty"`

										Prune *bool `tfsdk:"prune" yaml:"prune,omitempty"`

										SelfHeal *bool `tfsdk:"self_heal" yaml:"selfHeal,omitempty"`
									} `tfsdk:"automated" yaml:"automated,omitempty"`

									Retry *struct {
										Backoff *struct {
											Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

											Factor *int64 `tfsdk:"factor" yaml:"factor,omitempty"`

											MaxDuration *string `tfsdk:"max_duration" yaml:"maxDuration,omitempty"`
										} `tfsdk:"backoff" yaml:"backoff,omitempty"`

										Limit *int64 `tfsdk:"limit" yaml:"limit,omitempty"`
									} `tfsdk:"retry" yaml:"retry,omitempty"`

									SyncOptions *[]string `tfsdk:"sync_options" yaml:"syncOptions,omitempty"`
								} `tfsdk:"sync_policy" yaml:"syncPolicy,omitempty"`
							} `tfsdk:"spec" yaml:"spec,omitempty"`
						} `tfsdk:"template" yaml:"template,omitempty"`
					} `tfsdk:"list" yaml:"list,omitempty"`

					Matrix *map[string]string `tfsdk:"matrix" yaml:"matrix,omitempty"`

					Merge *map[string]string `tfsdk:"merge" yaml:"merge,omitempty"`

					PullRequest *struct {
						BitbucketServer *struct {
							Api *string `tfsdk:"api" yaml:"api,omitempty"`

							BasicAuth *struct {
								PasswordRef *struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
								} `tfsdk:"password_ref" yaml:"passwordRef,omitempty"`

								Username *string `tfsdk:"username" yaml:"username,omitempty"`
							} `tfsdk:"basic_auth" yaml:"basicAuth,omitempty"`

							Project *string `tfsdk:"project" yaml:"project,omitempty"`

							Repo *string `tfsdk:"repo" yaml:"repo,omitempty"`
						} `tfsdk:"bitbucket_server" yaml:"bitbucketServer,omitempty"`

						Filters *[]struct {
							BranchMatch *string `tfsdk:"branch_match" yaml:"branchMatch,omitempty"`
						} `tfsdk:"filters" yaml:"filters,omitempty"`

						Gitea *struct {
							Api *string `tfsdk:"api" yaml:"api,omitempty"`

							Insecure *bool `tfsdk:"insecure" yaml:"insecure,omitempty"`

							Owner *string `tfsdk:"owner" yaml:"owner,omitempty"`

							Repo *string `tfsdk:"repo" yaml:"repo,omitempty"`

							TokenRef *struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
							} `tfsdk:"token_ref" yaml:"tokenRef,omitempty"`
						} `tfsdk:"gitea" yaml:"gitea,omitempty"`

						Github *struct {
							Api *string `tfsdk:"api" yaml:"api,omitempty"`

							AppSecretName *string `tfsdk:"app_secret_name" yaml:"appSecretName,omitempty"`

							Labels *[]string `tfsdk:"labels" yaml:"labels,omitempty"`

							Owner *string `tfsdk:"owner" yaml:"owner,omitempty"`

							Repo *string `tfsdk:"repo" yaml:"repo,omitempty"`

							TokenRef *struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
							} `tfsdk:"token_ref" yaml:"tokenRef,omitempty"`
						} `tfsdk:"github" yaml:"github,omitempty"`

						Gitlab *struct {
							Api *string `tfsdk:"api" yaml:"api,omitempty"`

							Labels *[]string `tfsdk:"labels" yaml:"labels,omitempty"`

							Project *string `tfsdk:"project" yaml:"project,omitempty"`

							PullRequestState *string `tfsdk:"pull_request_state" yaml:"pullRequestState,omitempty"`

							TokenRef *struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
							} `tfsdk:"token_ref" yaml:"tokenRef,omitempty"`
						} `tfsdk:"gitlab" yaml:"gitlab,omitempty"`

						RequeueAfterSeconds *int64 `tfsdk:"requeue_after_seconds" yaml:"requeueAfterSeconds,omitempty"`

						Template *struct {
							Metadata *struct {
								Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

								Finalizers *[]string `tfsdk:"finalizers" yaml:"finalizers,omitempty"`

								Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
							} `tfsdk:"metadata" yaml:"metadata,omitempty"`

							Spec *struct {
								Destination *struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

									Server *string `tfsdk:"server" yaml:"server,omitempty"`
								} `tfsdk:"destination" yaml:"destination,omitempty"`

								IgnoreDifferences *[]struct {
									Group *string `tfsdk:"group" yaml:"group,omitempty"`

									JqPathExpressions *[]string `tfsdk:"jq_path_expressions" yaml:"jqPathExpressions,omitempty"`

									JsonPointers *[]string `tfsdk:"json_pointers" yaml:"jsonPointers,omitempty"`

									Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

									ManagedFieldsManagers *[]string `tfsdk:"managed_fields_managers" yaml:"managedFieldsManagers,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
								} `tfsdk:"ignore_differences" yaml:"ignoreDifferences,omitempty"`

								Info *[]struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"info" yaml:"info,omitempty"`

								Project *string `tfsdk:"project" yaml:"project,omitempty"`

								RevisionHistoryLimit *int64 `tfsdk:"revision_history_limit" yaml:"revisionHistoryLimit,omitempty"`

								Source *struct {
									Chart *string `tfsdk:"chart" yaml:"chart,omitempty"`

									Directory *struct {
										Exclude *string `tfsdk:"exclude" yaml:"exclude,omitempty"`

										Include *string `tfsdk:"include" yaml:"include,omitempty"`

										Jsonnet *struct {
											ExtVars *[]struct {
												Code *bool `tfsdk:"code" yaml:"code,omitempty"`

												Name *string `tfsdk:"name" yaml:"name,omitempty"`

												Value *string `tfsdk:"value" yaml:"value,omitempty"`
											} `tfsdk:"ext_vars" yaml:"extVars,omitempty"`

											Libs *[]string `tfsdk:"libs" yaml:"libs,omitempty"`

											Tlas *[]struct {
												Code *bool `tfsdk:"code" yaml:"code,omitempty"`

												Name *string `tfsdk:"name" yaml:"name,omitempty"`

												Value *string `tfsdk:"value" yaml:"value,omitempty"`
											} `tfsdk:"tlas" yaml:"tlas,omitempty"`
										} `tfsdk:"jsonnet" yaml:"jsonnet,omitempty"`

										Recurse *bool `tfsdk:"recurse" yaml:"recurse,omitempty"`
									} `tfsdk:"directory" yaml:"directory,omitempty"`

									Helm *struct {
										FileParameters *[]struct {
											Name *string `tfsdk:"name" yaml:"name,omitempty"`

											Path *string `tfsdk:"path" yaml:"path,omitempty"`
										} `tfsdk:"file_parameters" yaml:"fileParameters,omitempty"`

										IgnoreMissingValueFiles *bool `tfsdk:"ignore_missing_value_files" yaml:"ignoreMissingValueFiles,omitempty"`

										Parameters *[]struct {
											ForceString *bool `tfsdk:"force_string" yaml:"forceString,omitempty"`

											Name *string `tfsdk:"name" yaml:"name,omitempty"`

											Value *string `tfsdk:"value" yaml:"value,omitempty"`
										} `tfsdk:"parameters" yaml:"parameters,omitempty"`

										PassCredentials *bool `tfsdk:"pass_credentials" yaml:"passCredentials,omitempty"`

										ReleaseName *string `tfsdk:"release_name" yaml:"releaseName,omitempty"`

										SkipCrds *bool `tfsdk:"skip_crds" yaml:"skipCrds,omitempty"`

										ValueFiles *[]string `tfsdk:"value_files" yaml:"valueFiles,omitempty"`

										Values *string `tfsdk:"values" yaml:"values,omitempty"`

										Version *string `tfsdk:"version" yaml:"version,omitempty"`
									} `tfsdk:"helm" yaml:"helm,omitempty"`

									Kustomize *struct {
										CommonAnnotations *map[string]string `tfsdk:"common_annotations" yaml:"commonAnnotations,omitempty"`

										CommonLabels *map[string]string `tfsdk:"common_labels" yaml:"commonLabels,omitempty"`

										ForceCommonAnnotations *bool `tfsdk:"force_common_annotations" yaml:"forceCommonAnnotations,omitempty"`

										ForceCommonLabels *bool `tfsdk:"force_common_labels" yaml:"forceCommonLabels,omitempty"`

										Images *[]string `tfsdk:"images" yaml:"images,omitempty"`

										NamePrefix *string `tfsdk:"name_prefix" yaml:"namePrefix,omitempty"`

										NameSuffix *string `tfsdk:"name_suffix" yaml:"nameSuffix,omitempty"`

										Version *string `tfsdk:"version" yaml:"version,omitempty"`
									} `tfsdk:"kustomize" yaml:"kustomize,omitempty"`

									Path *string `tfsdk:"path" yaml:"path,omitempty"`

									Plugin *struct {
										Env *[]struct {
											Name *string `tfsdk:"name" yaml:"name,omitempty"`

											Value *string `tfsdk:"value" yaml:"value,omitempty"`
										} `tfsdk:"env" yaml:"env,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"plugin" yaml:"plugin,omitempty"`

									RepoURL *string `tfsdk:"repo_url" yaml:"repoURL,omitempty"`

									TargetRevision *string `tfsdk:"target_revision" yaml:"targetRevision,omitempty"`
								} `tfsdk:"source" yaml:"source,omitempty"`

								SyncPolicy *struct {
									Automated *struct {
										AllowEmpty *bool `tfsdk:"allow_empty" yaml:"allowEmpty,omitempty"`

										Prune *bool `tfsdk:"prune" yaml:"prune,omitempty"`

										SelfHeal *bool `tfsdk:"self_heal" yaml:"selfHeal,omitempty"`
									} `tfsdk:"automated" yaml:"automated,omitempty"`

									Retry *struct {
										Backoff *struct {
											Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

											Factor *int64 `tfsdk:"factor" yaml:"factor,omitempty"`

											MaxDuration *string `tfsdk:"max_duration" yaml:"maxDuration,omitempty"`
										} `tfsdk:"backoff" yaml:"backoff,omitempty"`

										Limit *int64 `tfsdk:"limit" yaml:"limit,omitempty"`
									} `tfsdk:"retry" yaml:"retry,omitempty"`

									SyncOptions *[]string `tfsdk:"sync_options" yaml:"syncOptions,omitempty"`
								} `tfsdk:"sync_policy" yaml:"syncPolicy,omitempty"`
							} `tfsdk:"spec" yaml:"spec,omitempty"`
						} `tfsdk:"template" yaml:"template,omitempty"`
					} `tfsdk:"pull_request" yaml:"pullRequest,omitempty"`

					ScmProvider *struct {
						AzureDevOps *struct {
							AccessTokenRef *struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
							} `tfsdk:"access_token_ref" yaml:"accessTokenRef,omitempty"`

							AllBranches *bool `tfsdk:"all_branches" yaml:"allBranches,omitempty"`

							Api *string `tfsdk:"api" yaml:"api,omitempty"`

							Organization *string `tfsdk:"organization" yaml:"organization,omitempty"`

							TeamProject *string `tfsdk:"team_project" yaml:"teamProject,omitempty"`
						} `tfsdk:"azure_dev_ops" yaml:"azureDevOps,omitempty"`

						Bitbucket *struct {
							AllBranches *bool `tfsdk:"all_branches" yaml:"allBranches,omitempty"`

							AppPasswordRef *struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
							} `tfsdk:"app_password_ref" yaml:"appPasswordRef,omitempty"`

							Owner *string `tfsdk:"owner" yaml:"owner,omitempty"`

							User *string `tfsdk:"user" yaml:"user,omitempty"`
						} `tfsdk:"bitbucket" yaml:"bitbucket,omitempty"`

						BitbucketServer *struct {
							AllBranches *bool `tfsdk:"all_branches" yaml:"allBranches,omitempty"`

							Api *string `tfsdk:"api" yaml:"api,omitempty"`

							BasicAuth *struct {
								PasswordRef *struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
								} `tfsdk:"password_ref" yaml:"passwordRef,omitempty"`

								Username *string `tfsdk:"username" yaml:"username,omitempty"`
							} `tfsdk:"basic_auth" yaml:"basicAuth,omitempty"`

							Project *string `tfsdk:"project" yaml:"project,omitempty"`
						} `tfsdk:"bitbucket_server" yaml:"bitbucketServer,omitempty"`

						CloneProtocol *string `tfsdk:"clone_protocol" yaml:"cloneProtocol,omitempty"`

						Filters *[]struct {
							BranchMatch *string `tfsdk:"branch_match" yaml:"branchMatch,omitempty"`

							LabelMatch *string `tfsdk:"label_match" yaml:"labelMatch,omitempty"`

							PathsDoNotExist *[]string `tfsdk:"paths_do_not_exist" yaml:"pathsDoNotExist,omitempty"`

							PathsExist *[]string `tfsdk:"paths_exist" yaml:"pathsExist,omitempty"`

							RepositoryMatch *string `tfsdk:"repository_match" yaml:"repositoryMatch,omitempty"`
						} `tfsdk:"filters" yaml:"filters,omitempty"`

						Gitea *struct {
							AllBranches *bool `tfsdk:"all_branches" yaml:"allBranches,omitempty"`

							Api *string `tfsdk:"api" yaml:"api,omitempty"`

							Insecure *bool `tfsdk:"insecure" yaml:"insecure,omitempty"`

							Owner *string `tfsdk:"owner" yaml:"owner,omitempty"`

							TokenRef *struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
							} `tfsdk:"token_ref" yaml:"tokenRef,omitempty"`
						} `tfsdk:"gitea" yaml:"gitea,omitempty"`

						Github *struct {
							AllBranches *bool `tfsdk:"all_branches" yaml:"allBranches,omitempty"`

							Api *string `tfsdk:"api" yaml:"api,omitempty"`

							AppSecretName *string `tfsdk:"app_secret_name" yaml:"appSecretName,omitempty"`

							Organization *string `tfsdk:"organization" yaml:"organization,omitempty"`

							TokenRef *struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
							} `tfsdk:"token_ref" yaml:"tokenRef,omitempty"`
						} `tfsdk:"github" yaml:"github,omitempty"`

						Gitlab *struct {
							AllBranches *bool `tfsdk:"all_branches" yaml:"allBranches,omitempty"`

							Api *string `tfsdk:"api" yaml:"api,omitempty"`

							Group *string `tfsdk:"group" yaml:"group,omitempty"`

							IncludeSubgroups *bool `tfsdk:"include_subgroups" yaml:"includeSubgroups,omitempty"`

							TokenRef *struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
							} `tfsdk:"token_ref" yaml:"tokenRef,omitempty"`
						} `tfsdk:"gitlab" yaml:"gitlab,omitempty"`

						RequeueAfterSeconds *int64 `tfsdk:"requeue_after_seconds" yaml:"requeueAfterSeconds,omitempty"`

						Template *struct {
							Metadata *struct {
								Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

								Finalizers *[]string `tfsdk:"finalizers" yaml:"finalizers,omitempty"`

								Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
							} `tfsdk:"metadata" yaml:"metadata,omitempty"`

							Spec *struct {
								Destination *struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

									Server *string `tfsdk:"server" yaml:"server,omitempty"`
								} `tfsdk:"destination" yaml:"destination,omitempty"`

								IgnoreDifferences *[]struct {
									Group *string `tfsdk:"group" yaml:"group,omitempty"`

									JqPathExpressions *[]string `tfsdk:"jq_path_expressions" yaml:"jqPathExpressions,omitempty"`

									JsonPointers *[]string `tfsdk:"json_pointers" yaml:"jsonPointers,omitempty"`

									Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

									ManagedFieldsManagers *[]string `tfsdk:"managed_fields_managers" yaml:"managedFieldsManagers,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
								} `tfsdk:"ignore_differences" yaml:"ignoreDifferences,omitempty"`

								Info *[]struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"info" yaml:"info,omitempty"`

								Project *string `tfsdk:"project" yaml:"project,omitempty"`

								RevisionHistoryLimit *int64 `tfsdk:"revision_history_limit" yaml:"revisionHistoryLimit,omitempty"`

								Source *struct {
									Chart *string `tfsdk:"chart" yaml:"chart,omitempty"`

									Directory *struct {
										Exclude *string `tfsdk:"exclude" yaml:"exclude,omitempty"`

										Include *string `tfsdk:"include" yaml:"include,omitempty"`

										Jsonnet *struct {
											ExtVars *[]struct {
												Code *bool `tfsdk:"code" yaml:"code,omitempty"`

												Name *string `tfsdk:"name" yaml:"name,omitempty"`

												Value *string `tfsdk:"value" yaml:"value,omitempty"`
											} `tfsdk:"ext_vars" yaml:"extVars,omitempty"`

											Libs *[]string `tfsdk:"libs" yaml:"libs,omitempty"`

											Tlas *[]struct {
												Code *bool `tfsdk:"code" yaml:"code,omitempty"`

												Name *string `tfsdk:"name" yaml:"name,omitempty"`

												Value *string `tfsdk:"value" yaml:"value,omitempty"`
											} `tfsdk:"tlas" yaml:"tlas,omitempty"`
										} `tfsdk:"jsonnet" yaml:"jsonnet,omitempty"`

										Recurse *bool `tfsdk:"recurse" yaml:"recurse,omitempty"`
									} `tfsdk:"directory" yaml:"directory,omitempty"`

									Helm *struct {
										FileParameters *[]struct {
											Name *string `tfsdk:"name" yaml:"name,omitempty"`

											Path *string `tfsdk:"path" yaml:"path,omitempty"`
										} `tfsdk:"file_parameters" yaml:"fileParameters,omitempty"`

										IgnoreMissingValueFiles *bool `tfsdk:"ignore_missing_value_files" yaml:"ignoreMissingValueFiles,omitempty"`

										Parameters *[]struct {
											ForceString *bool `tfsdk:"force_string" yaml:"forceString,omitempty"`

											Name *string `tfsdk:"name" yaml:"name,omitempty"`

											Value *string `tfsdk:"value" yaml:"value,omitempty"`
										} `tfsdk:"parameters" yaml:"parameters,omitempty"`

										PassCredentials *bool `tfsdk:"pass_credentials" yaml:"passCredentials,omitempty"`

										ReleaseName *string `tfsdk:"release_name" yaml:"releaseName,omitempty"`

										SkipCrds *bool `tfsdk:"skip_crds" yaml:"skipCrds,omitempty"`

										ValueFiles *[]string `tfsdk:"value_files" yaml:"valueFiles,omitempty"`

										Values *string `tfsdk:"values" yaml:"values,omitempty"`

										Version *string `tfsdk:"version" yaml:"version,omitempty"`
									} `tfsdk:"helm" yaml:"helm,omitempty"`

									Kustomize *struct {
										CommonAnnotations *map[string]string `tfsdk:"common_annotations" yaml:"commonAnnotations,omitempty"`

										CommonLabels *map[string]string `tfsdk:"common_labels" yaml:"commonLabels,omitempty"`

										ForceCommonAnnotations *bool `tfsdk:"force_common_annotations" yaml:"forceCommonAnnotations,omitempty"`

										ForceCommonLabels *bool `tfsdk:"force_common_labels" yaml:"forceCommonLabels,omitempty"`

										Images *[]string `tfsdk:"images" yaml:"images,omitempty"`

										NamePrefix *string `tfsdk:"name_prefix" yaml:"namePrefix,omitempty"`

										NameSuffix *string `tfsdk:"name_suffix" yaml:"nameSuffix,omitempty"`

										Version *string `tfsdk:"version" yaml:"version,omitempty"`
									} `tfsdk:"kustomize" yaml:"kustomize,omitempty"`

									Path *string `tfsdk:"path" yaml:"path,omitempty"`

									Plugin *struct {
										Env *[]struct {
											Name *string `tfsdk:"name" yaml:"name,omitempty"`

											Value *string `tfsdk:"value" yaml:"value,omitempty"`
										} `tfsdk:"env" yaml:"env,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"plugin" yaml:"plugin,omitempty"`

									RepoURL *string `tfsdk:"repo_url" yaml:"repoURL,omitempty"`

									TargetRevision *string `tfsdk:"target_revision" yaml:"targetRevision,omitempty"`
								} `tfsdk:"source" yaml:"source,omitempty"`

								SyncPolicy *struct {
									Automated *struct {
										AllowEmpty *bool `tfsdk:"allow_empty" yaml:"allowEmpty,omitempty"`

										Prune *bool `tfsdk:"prune" yaml:"prune,omitempty"`

										SelfHeal *bool `tfsdk:"self_heal" yaml:"selfHeal,omitempty"`
									} `tfsdk:"automated" yaml:"automated,omitempty"`

									Retry *struct {
										Backoff *struct {
											Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

											Factor *int64 `tfsdk:"factor" yaml:"factor,omitempty"`

											MaxDuration *string `tfsdk:"max_duration" yaml:"maxDuration,omitempty"`
										} `tfsdk:"backoff" yaml:"backoff,omitempty"`

										Limit *int64 `tfsdk:"limit" yaml:"limit,omitempty"`
									} `tfsdk:"retry" yaml:"retry,omitempty"`

									SyncOptions *[]string `tfsdk:"sync_options" yaml:"syncOptions,omitempty"`
								} `tfsdk:"sync_policy" yaml:"syncPolicy,omitempty"`
							} `tfsdk:"spec" yaml:"spec,omitempty"`
						} `tfsdk:"template" yaml:"template,omitempty"`
					} `tfsdk:"scm_provider" yaml:"scmProvider,omitempty"`

					Selector *struct {
						MatchExpressions *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
						} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

						MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
					} `tfsdk:"selector" yaml:"selector,omitempty"`
				} `tfsdk:"generators" yaml:"generators,omitempty"`

				MergeKeys *[]string `tfsdk:"merge_keys" yaml:"mergeKeys,omitempty"`

				Template *struct {
					Metadata *struct {
						Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Finalizers *[]string `tfsdk:"finalizers" yaml:"finalizers,omitempty"`

						Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`

					Spec *struct {
						Destination *struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

							Server *string `tfsdk:"server" yaml:"server,omitempty"`
						} `tfsdk:"destination" yaml:"destination,omitempty"`

						IgnoreDifferences *[]struct {
							Group *string `tfsdk:"group" yaml:"group,omitempty"`

							JqPathExpressions *[]string `tfsdk:"jq_path_expressions" yaml:"jqPathExpressions,omitempty"`

							JsonPointers *[]string `tfsdk:"json_pointers" yaml:"jsonPointers,omitempty"`

							Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

							ManagedFieldsManagers *[]string `tfsdk:"managed_fields_managers" yaml:"managedFieldsManagers,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
						} `tfsdk:"ignore_differences" yaml:"ignoreDifferences,omitempty"`

						Info *[]struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"info" yaml:"info,omitempty"`

						Project *string `tfsdk:"project" yaml:"project,omitempty"`

						RevisionHistoryLimit *int64 `tfsdk:"revision_history_limit" yaml:"revisionHistoryLimit,omitempty"`

						Source *struct {
							Chart *string `tfsdk:"chart" yaml:"chart,omitempty"`

							Directory *struct {
								Exclude *string `tfsdk:"exclude" yaml:"exclude,omitempty"`

								Include *string `tfsdk:"include" yaml:"include,omitempty"`

								Jsonnet *struct {
									ExtVars *[]struct {
										Code *bool `tfsdk:"code" yaml:"code,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`

										Value *string `tfsdk:"value" yaml:"value,omitempty"`
									} `tfsdk:"ext_vars" yaml:"extVars,omitempty"`

									Libs *[]string `tfsdk:"libs" yaml:"libs,omitempty"`

									Tlas *[]struct {
										Code *bool `tfsdk:"code" yaml:"code,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`

										Value *string `tfsdk:"value" yaml:"value,omitempty"`
									} `tfsdk:"tlas" yaml:"tlas,omitempty"`
								} `tfsdk:"jsonnet" yaml:"jsonnet,omitempty"`

								Recurse *bool `tfsdk:"recurse" yaml:"recurse,omitempty"`
							} `tfsdk:"directory" yaml:"directory,omitempty"`

							Helm *struct {
								FileParameters *[]struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Path *string `tfsdk:"path" yaml:"path,omitempty"`
								} `tfsdk:"file_parameters" yaml:"fileParameters,omitempty"`

								IgnoreMissingValueFiles *bool `tfsdk:"ignore_missing_value_files" yaml:"ignoreMissingValueFiles,omitempty"`

								Parameters *[]struct {
									ForceString *bool `tfsdk:"force_string" yaml:"forceString,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"parameters" yaml:"parameters,omitempty"`

								PassCredentials *bool `tfsdk:"pass_credentials" yaml:"passCredentials,omitempty"`

								ReleaseName *string `tfsdk:"release_name" yaml:"releaseName,omitempty"`

								SkipCrds *bool `tfsdk:"skip_crds" yaml:"skipCrds,omitempty"`

								ValueFiles *[]string `tfsdk:"value_files" yaml:"valueFiles,omitempty"`

								Values *string `tfsdk:"values" yaml:"values,omitempty"`

								Version *string `tfsdk:"version" yaml:"version,omitempty"`
							} `tfsdk:"helm" yaml:"helm,omitempty"`

							Kustomize *struct {
								CommonAnnotations *map[string]string `tfsdk:"common_annotations" yaml:"commonAnnotations,omitempty"`

								CommonLabels *map[string]string `tfsdk:"common_labels" yaml:"commonLabels,omitempty"`

								ForceCommonAnnotations *bool `tfsdk:"force_common_annotations" yaml:"forceCommonAnnotations,omitempty"`

								ForceCommonLabels *bool `tfsdk:"force_common_labels" yaml:"forceCommonLabels,omitempty"`

								Images *[]string `tfsdk:"images" yaml:"images,omitempty"`

								NamePrefix *string `tfsdk:"name_prefix" yaml:"namePrefix,omitempty"`

								NameSuffix *string `tfsdk:"name_suffix" yaml:"nameSuffix,omitempty"`

								Version *string `tfsdk:"version" yaml:"version,omitempty"`
							} `tfsdk:"kustomize" yaml:"kustomize,omitempty"`

							Path *string `tfsdk:"path" yaml:"path,omitempty"`

							Plugin *struct {
								Env *[]struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"env" yaml:"env,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`
							} `tfsdk:"plugin" yaml:"plugin,omitempty"`

							RepoURL *string `tfsdk:"repo_url" yaml:"repoURL,omitempty"`

							TargetRevision *string `tfsdk:"target_revision" yaml:"targetRevision,omitempty"`
						} `tfsdk:"source" yaml:"source,omitempty"`

						SyncPolicy *struct {
							Automated *struct {
								AllowEmpty *bool `tfsdk:"allow_empty" yaml:"allowEmpty,omitempty"`

								Prune *bool `tfsdk:"prune" yaml:"prune,omitempty"`

								SelfHeal *bool `tfsdk:"self_heal" yaml:"selfHeal,omitempty"`
							} `tfsdk:"automated" yaml:"automated,omitempty"`

							Retry *struct {
								Backoff *struct {
									Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

									Factor *int64 `tfsdk:"factor" yaml:"factor,omitempty"`

									MaxDuration *string `tfsdk:"max_duration" yaml:"maxDuration,omitempty"`
								} `tfsdk:"backoff" yaml:"backoff,omitempty"`

								Limit *int64 `tfsdk:"limit" yaml:"limit,omitempty"`
							} `tfsdk:"retry" yaml:"retry,omitempty"`

							SyncOptions *[]string `tfsdk:"sync_options" yaml:"syncOptions,omitempty"`
						} `tfsdk:"sync_policy" yaml:"syncPolicy,omitempty"`
					} `tfsdk:"spec" yaml:"spec,omitempty"`
				} `tfsdk:"template" yaml:"template,omitempty"`
			} `tfsdk:"merge" yaml:"merge,omitempty"`

			PullRequest *struct {
				BitbucketServer *struct {
					Api *string `tfsdk:"api" yaml:"api,omitempty"`

					BasicAuth *struct {
						PasswordRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
						} `tfsdk:"password_ref" yaml:"passwordRef,omitempty"`

						Username *string `tfsdk:"username" yaml:"username,omitempty"`
					} `tfsdk:"basic_auth" yaml:"basicAuth,omitempty"`

					Project *string `tfsdk:"project" yaml:"project,omitempty"`

					Repo *string `tfsdk:"repo" yaml:"repo,omitempty"`
				} `tfsdk:"bitbucket_server" yaml:"bitbucketServer,omitempty"`

				Filters *[]struct {
					BranchMatch *string `tfsdk:"branch_match" yaml:"branchMatch,omitempty"`
				} `tfsdk:"filters" yaml:"filters,omitempty"`

				Gitea *struct {
					Api *string `tfsdk:"api" yaml:"api,omitempty"`

					Insecure *bool `tfsdk:"insecure" yaml:"insecure,omitempty"`

					Owner *string `tfsdk:"owner" yaml:"owner,omitempty"`

					Repo *string `tfsdk:"repo" yaml:"repo,omitempty"`

					TokenRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
					} `tfsdk:"token_ref" yaml:"tokenRef,omitempty"`
				} `tfsdk:"gitea" yaml:"gitea,omitempty"`

				Github *struct {
					Api *string `tfsdk:"api" yaml:"api,omitempty"`

					AppSecretName *string `tfsdk:"app_secret_name" yaml:"appSecretName,omitempty"`

					Labels *[]string `tfsdk:"labels" yaml:"labels,omitempty"`

					Owner *string `tfsdk:"owner" yaml:"owner,omitempty"`

					Repo *string `tfsdk:"repo" yaml:"repo,omitempty"`

					TokenRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
					} `tfsdk:"token_ref" yaml:"tokenRef,omitempty"`
				} `tfsdk:"github" yaml:"github,omitempty"`

				Gitlab *struct {
					Api *string `tfsdk:"api" yaml:"api,omitempty"`

					Labels *[]string `tfsdk:"labels" yaml:"labels,omitempty"`

					Project *string `tfsdk:"project" yaml:"project,omitempty"`

					PullRequestState *string `tfsdk:"pull_request_state" yaml:"pullRequestState,omitempty"`

					TokenRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
					} `tfsdk:"token_ref" yaml:"tokenRef,omitempty"`
				} `tfsdk:"gitlab" yaml:"gitlab,omitempty"`

				RequeueAfterSeconds *int64 `tfsdk:"requeue_after_seconds" yaml:"requeueAfterSeconds,omitempty"`

				Template *struct {
					Metadata *struct {
						Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Finalizers *[]string `tfsdk:"finalizers" yaml:"finalizers,omitempty"`

						Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`

					Spec *struct {
						Destination *struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

							Server *string `tfsdk:"server" yaml:"server,omitempty"`
						} `tfsdk:"destination" yaml:"destination,omitempty"`

						IgnoreDifferences *[]struct {
							Group *string `tfsdk:"group" yaml:"group,omitempty"`

							JqPathExpressions *[]string `tfsdk:"jq_path_expressions" yaml:"jqPathExpressions,omitempty"`

							JsonPointers *[]string `tfsdk:"json_pointers" yaml:"jsonPointers,omitempty"`

							Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

							ManagedFieldsManagers *[]string `tfsdk:"managed_fields_managers" yaml:"managedFieldsManagers,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
						} `tfsdk:"ignore_differences" yaml:"ignoreDifferences,omitempty"`

						Info *[]struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"info" yaml:"info,omitempty"`

						Project *string `tfsdk:"project" yaml:"project,omitempty"`

						RevisionHistoryLimit *int64 `tfsdk:"revision_history_limit" yaml:"revisionHistoryLimit,omitempty"`

						Source *struct {
							Chart *string `tfsdk:"chart" yaml:"chart,omitempty"`

							Directory *struct {
								Exclude *string `tfsdk:"exclude" yaml:"exclude,omitempty"`

								Include *string `tfsdk:"include" yaml:"include,omitempty"`

								Jsonnet *struct {
									ExtVars *[]struct {
										Code *bool `tfsdk:"code" yaml:"code,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`

										Value *string `tfsdk:"value" yaml:"value,omitempty"`
									} `tfsdk:"ext_vars" yaml:"extVars,omitempty"`

									Libs *[]string `tfsdk:"libs" yaml:"libs,omitempty"`

									Tlas *[]struct {
										Code *bool `tfsdk:"code" yaml:"code,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`

										Value *string `tfsdk:"value" yaml:"value,omitempty"`
									} `tfsdk:"tlas" yaml:"tlas,omitempty"`
								} `tfsdk:"jsonnet" yaml:"jsonnet,omitempty"`

								Recurse *bool `tfsdk:"recurse" yaml:"recurse,omitempty"`
							} `tfsdk:"directory" yaml:"directory,omitempty"`

							Helm *struct {
								FileParameters *[]struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Path *string `tfsdk:"path" yaml:"path,omitempty"`
								} `tfsdk:"file_parameters" yaml:"fileParameters,omitempty"`

								IgnoreMissingValueFiles *bool `tfsdk:"ignore_missing_value_files" yaml:"ignoreMissingValueFiles,omitempty"`

								Parameters *[]struct {
									ForceString *bool `tfsdk:"force_string" yaml:"forceString,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"parameters" yaml:"parameters,omitempty"`

								PassCredentials *bool `tfsdk:"pass_credentials" yaml:"passCredentials,omitempty"`

								ReleaseName *string `tfsdk:"release_name" yaml:"releaseName,omitempty"`

								SkipCrds *bool `tfsdk:"skip_crds" yaml:"skipCrds,omitempty"`

								ValueFiles *[]string `tfsdk:"value_files" yaml:"valueFiles,omitempty"`

								Values *string `tfsdk:"values" yaml:"values,omitempty"`

								Version *string `tfsdk:"version" yaml:"version,omitempty"`
							} `tfsdk:"helm" yaml:"helm,omitempty"`

							Kustomize *struct {
								CommonAnnotations *map[string]string `tfsdk:"common_annotations" yaml:"commonAnnotations,omitempty"`

								CommonLabels *map[string]string `tfsdk:"common_labels" yaml:"commonLabels,omitempty"`

								ForceCommonAnnotations *bool `tfsdk:"force_common_annotations" yaml:"forceCommonAnnotations,omitempty"`

								ForceCommonLabels *bool `tfsdk:"force_common_labels" yaml:"forceCommonLabels,omitempty"`

								Images *[]string `tfsdk:"images" yaml:"images,omitempty"`

								NamePrefix *string `tfsdk:"name_prefix" yaml:"namePrefix,omitempty"`

								NameSuffix *string `tfsdk:"name_suffix" yaml:"nameSuffix,omitempty"`

								Version *string `tfsdk:"version" yaml:"version,omitempty"`
							} `tfsdk:"kustomize" yaml:"kustomize,omitempty"`

							Path *string `tfsdk:"path" yaml:"path,omitempty"`

							Plugin *struct {
								Env *[]struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"env" yaml:"env,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`
							} `tfsdk:"plugin" yaml:"plugin,omitempty"`

							RepoURL *string `tfsdk:"repo_url" yaml:"repoURL,omitempty"`

							TargetRevision *string `tfsdk:"target_revision" yaml:"targetRevision,omitempty"`
						} `tfsdk:"source" yaml:"source,omitempty"`

						SyncPolicy *struct {
							Automated *struct {
								AllowEmpty *bool `tfsdk:"allow_empty" yaml:"allowEmpty,omitempty"`

								Prune *bool `tfsdk:"prune" yaml:"prune,omitempty"`

								SelfHeal *bool `tfsdk:"self_heal" yaml:"selfHeal,omitempty"`
							} `tfsdk:"automated" yaml:"automated,omitempty"`

							Retry *struct {
								Backoff *struct {
									Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

									Factor *int64 `tfsdk:"factor" yaml:"factor,omitempty"`

									MaxDuration *string `tfsdk:"max_duration" yaml:"maxDuration,omitempty"`
								} `tfsdk:"backoff" yaml:"backoff,omitempty"`

								Limit *int64 `tfsdk:"limit" yaml:"limit,omitempty"`
							} `tfsdk:"retry" yaml:"retry,omitempty"`

							SyncOptions *[]string `tfsdk:"sync_options" yaml:"syncOptions,omitempty"`
						} `tfsdk:"sync_policy" yaml:"syncPolicy,omitempty"`
					} `tfsdk:"spec" yaml:"spec,omitempty"`
				} `tfsdk:"template" yaml:"template,omitempty"`
			} `tfsdk:"pull_request" yaml:"pullRequest,omitempty"`

			ScmProvider *struct {
				AzureDevOps *struct {
					AccessTokenRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
					} `tfsdk:"access_token_ref" yaml:"accessTokenRef,omitempty"`

					AllBranches *bool `tfsdk:"all_branches" yaml:"allBranches,omitempty"`

					Api *string `tfsdk:"api" yaml:"api,omitempty"`

					Organization *string `tfsdk:"organization" yaml:"organization,omitempty"`

					TeamProject *string `tfsdk:"team_project" yaml:"teamProject,omitempty"`
				} `tfsdk:"azure_dev_ops" yaml:"azureDevOps,omitempty"`

				Bitbucket *struct {
					AllBranches *bool `tfsdk:"all_branches" yaml:"allBranches,omitempty"`

					AppPasswordRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
					} `tfsdk:"app_password_ref" yaml:"appPasswordRef,omitempty"`

					Owner *string `tfsdk:"owner" yaml:"owner,omitempty"`

					User *string `tfsdk:"user" yaml:"user,omitempty"`
				} `tfsdk:"bitbucket" yaml:"bitbucket,omitempty"`

				BitbucketServer *struct {
					AllBranches *bool `tfsdk:"all_branches" yaml:"allBranches,omitempty"`

					Api *string `tfsdk:"api" yaml:"api,omitempty"`

					BasicAuth *struct {
						PasswordRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
						} `tfsdk:"password_ref" yaml:"passwordRef,omitempty"`

						Username *string `tfsdk:"username" yaml:"username,omitempty"`
					} `tfsdk:"basic_auth" yaml:"basicAuth,omitempty"`

					Project *string `tfsdk:"project" yaml:"project,omitempty"`
				} `tfsdk:"bitbucket_server" yaml:"bitbucketServer,omitempty"`

				CloneProtocol *string `tfsdk:"clone_protocol" yaml:"cloneProtocol,omitempty"`

				Filters *[]struct {
					BranchMatch *string `tfsdk:"branch_match" yaml:"branchMatch,omitempty"`

					LabelMatch *string `tfsdk:"label_match" yaml:"labelMatch,omitempty"`

					PathsDoNotExist *[]string `tfsdk:"paths_do_not_exist" yaml:"pathsDoNotExist,omitempty"`

					PathsExist *[]string `tfsdk:"paths_exist" yaml:"pathsExist,omitempty"`

					RepositoryMatch *string `tfsdk:"repository_match" yaml:"repositoryMatch,omitempty"`
				} `tfsdk:"filters" yaml:"filters,omitempty"`

				Gitea *struct {
					AllBranches *bool `tfsdk:"all_branches" yaml:"allBranches,omitempty"`

					Api *string `tfsdk:"api" yaml:"api,omitempty"`

					Insecure *bool `tfsdk:"insecure" yaml:"insecure,omitempty"`

					Owner *string `tfsdk:"owner" yaml:"owner,omitempty"`

					TokenRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
					} `tfsdk:"token_ref" yaml:"tokenRef,omitempty"`
				} `tfsdk:"gitea" yaml:"gitea,omitempty"`

				Github *struct {
					AllBranches *bool `tfsdk:"all_branches" yaml:"allBranches,omitempty"`

					Api *string `tfsdk:"api" yaml:"api,omitempty"`

					AppSecretName *string `tfsdk:"app_secret_name" yaml:"appSecretName,omitempty"`

					Organization *string `tfsdk:"organization" yaml:"organization,omitempty"`

					TokenRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
					} `tfsdk:"token_ref" yaml:"tokenRef,omitempty"`
				} `tfsdk:"github" yaml:"github,omitempty"`

				Gitlab *struct {
					AllBranches *bool `tfsdk:"all_branches" yaml:"allBranches,omitempty"`

					Api *string `tfsdk:"api" yaml:"api,omitempty"`

					Group *string `tfsdk:"group" yaml:"group,omitempty"`

					IncludeSubgroups *bool `tfsdk:"include_subgroups" yaml:"includeSubgroups,omitempty"`

					TokenRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
					} `tfsdk:"token_ref" yaml:"tokenRef,omitempty"`
				} `tfsdk:"gitlab" yaml:"gitlab,omitempty"`

				RequeueAfterSeconds *int64 `tfsdk:"requeue_after_seconds" yaml:"requeueAfterSeconds,omitempty"`

				Template *struct {
					Metadata *struct {
						Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Finalizers *[]string `tfsdk:"finalizers" yaml:"finalizers,omitempty"`

						Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`

					Spec *struct {
						Destination *struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

							Server *string `tfsdk:"server" yaml:"server,omitempty"`
						} `tfsdk:"destination" yaml:"destination,omitempty"`

						IgnoreDifferences *[]struct {
							Group *string `tfsdk:"group" yaml:"group,omitempty"`

							JqPathExpressions *[]string `tfsdk:"jq_path_expressions" yaml:"jqPathExpressions,omitempty"`

							JsonPointers *[]string `tfsdk:"json_pointers" yaml:"jsonPointers,omitempty"`

							Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

							ManagedFieldsManagers *[]string `tfsdk:"managed_fields_managers" yaml:"managedFieldsManagers,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
						} `tfsdk:"ignore_differences" yaml:"ignoreDifferences,omitempty"`

						Info *[]struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"info" yaml:"info,omitempty"`

						Project *string `tfsdk:"project" yaml:"project,omitempty"`

						RevisionHistoryLimit *int64 `tfsdk:"revision_history_limit" yaml:"revisionHistoryLimit,omitempty"`

						Source *struct {
							Chart *string `tfsdk:"chart" yaml:"chart,omitempty"`

							Directory *struct {
								Exclude *string `tfsdk:"exclude" yaml:"exclude,omitempty"`

								Include *string `tfsdk:"include" yaml:"include,omitempty"`

								Jsonnet *struct {
									ExtVars *[]struct {
										Code *bool `tfsdk:"code" yaml:"code,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`

										Value *string `tfsdk:"value" yaml:"value,omitempty"`
									} `tfsdk:"ext_vars" yaml:"extVars,omitempty"`

									Libs *[]string `tfsdk:"libs" yaml:"libs,omitempty"`

									Tlas *[]struct {
										Code *bool `tfsdk:"code" yaml:"code,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`

										Value *string `tfsdk:"value" yaml:"value,omitempty"`
									} `tfsdk:"tlas" yaml:"tlas,omitempty"`
								} `tfsdk:"jsonnet" yaml:"jsonnet,omitempty"`

								Recurse *bool `tfsdk:"recurse" yaml:"recurse,omitempty"`
							} `tfsdk:"directory" yaml:"directory,omitempty"`

							Helm *struct {
								FileParameters *[]struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Path *string `tfsdk:"path" yaml:"path,omitempty"`
								} `tfsdk:"file_parameters" yaml:"fileParameters,omitempty"`

								IgnoreMissingValueFiles *bool `tfsdk:"ignore_missing_value_files" yaml:"ignoreMissingValueFiles,omitempty"`

								Parameters *[]struct {
									ForceString *bool `tfsdk:"force_string" yaml:"forceString,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"parameters" yaml:"parameters,omitempty"`

								PassCredentials *bool `tfsdk:"pass_credentials" yaml:"passCredentials,omitempty"`

								ReleaseName *string `tfsdk:"release_name" yaml:"releaseName,omitempty"`

								SkipCrds *bool `tfsdk:"skip_crds" yaml:"skipCrds,omitempty"`

								ValueFiles *[]string `tfsdk:"value_files" yaml:"valueFiles,omitempty"`

								Values *string `tfsdk:"values" yaml:"values,omitempty"`

								Version *string `tfsdk:"version" yaml:"version,omitempty"`
							} `tfsdk:"helm" yaml:"helm,omitempty"`

							Kustomize *struct {
								CommonAnnotations *map[string]string `tfsdk:"common_annotations" yaml:"commonAnnotations,omitempty"`

								CommonLabels *map[string]string `tfsdk:"common_labels" yaml:"commonLabels,omitempty"`

								ForceCommonAnnotations *bool `tfsdk:"force_common_annotations" yaml:"forceCommonAnnotations,omitempty"`

								ForceCommonLabels *bool `tfsdk:"force_common_labels" yaml:"forceCommonLabels,omitempty"`

								Images *[]string `tfsdk:"images" yaml:"images,omitempty"`

								NamePrefix *string `tfsdk:"name_prefix" yaml:"namePrefix,omitempty"`

								NameSuffix *string `tfsdk:"name_suffix" yaml:"nameSuffix,omitempty"`

								Version *string `tfsdk:"version" yaml:"version,omitempty"`
							} `tfsdk:"kustomize" yaml:"kustomize,omitempty"`

							Path *string `tfsdk:"path" yaml:"path,omitempty"`

							Plugin *struct {
								Env *[]struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"env" yaml:"env,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`
							} `tfsdk:"plugin" yaml:"plugin,omitempty"`

							RepoURL *string `tfsdk:"repo_url" yaml:"repoURL,omitempty"`

							TargetRevision *string `tfsdk:"target_revision" yaml:"targetRevision,omitempty"`
						} `tfsdk:"source" yaml:"source,omitempty"`

						SyncPolicy *struct {
							Automated *struct {
								AllowEmpty *bool `tfsdk:"allow_empty" yaml:"allowEmpty,omitempty"`

								Prune *bool `tfsdk:"prune" yaml:"prune,omitempty"`

								SelfHeal *bool `tfsdk:"self_heal" yaml:"selfHeal,omitempty"`
							} `tfsdk:"automated" yaml:"automated,omitempty"`

							Retry *struct {
								Backoff *struct {
									Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

									Factor *int64 `tfsdk:"factor" yaml:"factor,omitempty"`

									MaxDuration *string `tfsdk:"max_duration" yaml:"maxDuration,omitempty"`
								} `tfsdk:"backoff" yaml:"backoff,omitempty"`

								Limit *int64 `tfsdk:"limit" yaml:"limit,omitempty"`
							} `tfsdk:"retry" yaml:"retry,omitempty"`

							SyncOptions *[]string `tfsdk:"sync_options" yaml:"syncOptions,omitempty"`
						} `tfsdk:"sync_policy" yaml:"syncPolicy,omitempty"`
					} `tfsdk:"spec" yaml:"spec,omitempty"`
				} `tfsdk:"template" yaml:"template,omitempty"`
			} `tfsdk:"scm_provider" yaml:"scmProvider,omitempty"`

			Selector *struct {
				MatchExpressions *[]struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

					Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
				} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

				MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
			} `tfsdk:"selector" yaml:"selector,omitempty"`
		} `tfsdk:"generators" yaml:"generators,omitempty"`

		GoTemplate *bool `tfsdk:"go_template" yaml:"goTemplate,omitempty"`

		SyncPolicy *struct {
			PreserveResourcesOnDeletion *bool `tfsdk:"preserve_resources_on_deletion" yaml:"preserveResourcesOnDeletion,omitempty"`
		} `tfsdk:"sync_policy" yaml:"syncPolicy,omitempty"`

		Template *struct {
			Metadata *struct {
				Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

				Finalizers *[]string `tfsdk:"finalizers" yaml:"finalizers,omitempty"`

				Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
			} `tfsdk:"metadata" yaml:"metadata,omitempty"`

			Spec *struct {
				Destination *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

					Server *string `tfsdk:"server" yaml:"server,omitempty"`
				} `tfsdk:"destination" yaml:"destination,omitempty"`

				IgnoreDifferences *[]struct {
					Group *string `tfsdk:"group" yaml:"group,omitempty"`

					JqPathExpressions *[]string `tfsdk:"jq_path_expressions" yaml:"jqPathExpressions,omitempty"`

					JsonPointers *[]string `tfsdk:"json_pointers" yaml:"jsonPointers,omitempty"`

					Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

					ManagedFieldsManagers *[]string `tfsdk:"managed_fields_managers" yaml:"managedFieldsManagers,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
				} `tfsdk:"ignore_differences" yaml:"ignoreDifferences,omitempty"`

				Info *[]struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"info" yaml:"info,omitempty"`

				Project *string `tfsdk:"project" yaml:"project,omitempty"`

				RevisionHistoryLimit *int64 `tfsdk:"revision_history_limit" yaml:"revisionHistoryLimit,omitempty"`

				Source *struct {
					Chart *string `tfsdk:"chart" yaml:"chart,omitempty"`

					Directory *struct {
						Exclude *string `tfsdk:"exclude" yaml:"exclude,omitempty"`

						Include *string `tfsdk:"include" yaml:"include,omitempty"`

						Jsonnet *struct {
							ExtVars *[]struct {
								Code *bool `tfsdk:"code" yaml:"code,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Value *string `tfsdk:"value" yaml:"value,omitempty"`
							} `tfsdk:"ext_vars" yaml:"extVars,omitempty"`

							Libs *[]string `tfsdk:"libs" yaml:"libs,omitempty"`

							Tlas *[]struct {
								Code *bool `tfsdk:"code" yaml:"code,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Value *string `tfsdk:"value" yaml:"value,omitempty"`
							} `tfsdk:"tlas" yaml:"tlas,omitempty"`
						} `tfsdk:"jsonnet" yaml:"jsonnet,omitempty"`

						Recurse *bool `tfsdk:"recurse" yaml:"recurse,omitempty"`
					} `tfsdk:"directory" yaml:"directory,omitempty"`

					Helm *struct {
						FileParameters *[]struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Path *string `tfsdk:"path" yaml:"path,omitempty"`
						} `tfsdk:"file_parameters" yaml:"fileParameters,omitempty"`

						IgnoreMissingValueFiles *bool `tfsdk:"ignore_missing_value_files" yaml:"ignoreMissingValueFiles,omitempty"`

						Parameters *[]struct {
							ForceString *bool `tfsdk:"force_string" yaml:"forceString,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"parameters" yaml:"parameters,omitempty"`

						PassCredentials *bool `tfsdk:"pass_credentials" yaml:"passCredentials,omitempty"`

						ReleaseName *string `tfsdk:"release_name" yaml:"releaseName,omitempty"`

						SkipCrds *bool `tfsdk:"skip_crds" yaml:"skipCrds,omitempty"`

						ValueFiles *[]string `tfsdk:"value_files" yaml:"valueFiles,omitempty"`

						Values *string `tfsdk:"values" yaml:"values,omitempty"`

						Version *string `tfsdk:"version" yaml:"version,omitempty"`
					} `tfsdk:"helm" yaml:"helm,omitempty"`

					Kustomize *struct {
						CommonAnnotations *map[string]string `tfsdk:"common_annotations" yaml:"commonAnnotations,omitempty"`

						CommonLabels *map[string]string `tfsdk:"common_labels" yaml:"commonLabels,omitempty"`

						ForceCommonAnnotations *bool `tfsdk:"force_common_annotations" yaml:"forceCommonAnnotations,omitempty"`

						ForceCommonLabels *bool `tfsdk:"force_common_labels" yaml:"forceCommonLabels,omitempty"`

						Images *[]string `tfsdk:"images" yaml:"images,omitempty"`

						NamePrefix *string `tfsdk:"name_prefix" yaml:"namePrefix,omitempty"`

						NameSuffix *string `tfsdk:"name_suffix" yaml:"nameSuffix,omitempty"`

						Version *string `tfsdk:"version" yaml:"version,omitempty"`
					} `tfsdk:"kustomize" yaml:"kustomize,omitempty"`

					Path *string `tfsdk:"path" yaml:"path,omitempty"`

					Plugin *struct {
						Env *[]struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"env" yaml:"env,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`
					} `tfsdk:"plugin" yaml:"plugin,omitempty"`

					RepoURL *string `tfsdk:"repo_url" yaml:"repoURL,omitempty"`

					TargetRevision *string `tfsdk:"target_revision" yaml:"targetRevision,omitempty"`
				} `tfsdk:"source" yaml:"source,omitempty"`

				SyncPolicy *struct {
					Automated *struct {
						AllowEmpty *bool `tfsdk:"allow_empty" yaml:"allowEmpty,omitempty"`

						Prune *bool `tfsdk:"prune" yaml:"prune,omitempty"`

						SelfHeal *bool `tfsdk:"self_heal" yaml:"selfHeal,omitempty"`
					} `tfsdk:"automated" yaml:"automated,omitempty"`

					Retry *struct {
						Backoff *struct {
							Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

							Factor *int64 `tfsdk:"factor" yaml:"factor,omitempty"`

							MaxDuration *string `tfsdk:"max_duration" yaml:"maxDuration,omitempty"`
						} `tfsdk:"backoff" yaml:"backoff,omitempty"`

						Limit *int64 `tfsdk:"limit" yaml:"limit,omitempty"`
					} `tfsdk:"retry" yaml:"retry,omitempty"`

					SyncOptions *[]string `tfsdk:"sync_options" yaml:"syncOptions,omitempty"`
				} `tfsdk:"sync_policy" yaml:"syncPolicy,omitempty"`
			} `tfsdk:"spec" yaml:"spec,omitempty"`
		} `tfsdk:"template" yaml:"template,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewArgoprojIoApplicationSetV1Alpha1Resource() resource.Resource {
	return &ArgoprojIoApplicationSetV1Alpha1Resource{}
}

func (r *ArgoprojIoApplicationSetV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_argoproj_io_application_set_v1alpha1"
}

func (r *ArgoprojIoApplicationSetV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "",
		MarkdownDescription: "",
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
				Description:         "",
				MarkdownDescription: "",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"generators": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"cluster_decision_resource": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"config_map_ref": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"label_selector": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"match_expressions": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"operator": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"values": {
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

											"match_labels": {
												Description:         "",
												MarkdownDescription: "",

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

									"name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"requeue_after_seconds": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"template": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"metadata": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"finalizers": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

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
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},

											"spec": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"destination": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

															"server": {
																Description:         "",
																MarkdownDescription: "",

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

													"ignore_differences": {
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

															"jq_path_expressions": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"json_pointers": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"kind": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"managed_fields_managers": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

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
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"info": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"value": {
																Description:         "",
																MarkdownDescription: "",

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

													"project": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"revision_history_limit": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"source": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"chart": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"directory": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"exclude": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"include": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"jsonnet": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"ext_vars": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"code": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"name": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"value": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"libs": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"tlas": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"code": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"name": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"value": {
																						Description:         "",
																						MarkdownDescription: "",

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

																	"recurse": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"helm": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"file_parameters": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"path": {
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

																	"ignore_missing_value_files": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"parameters": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"force_string": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

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

																			"value": {
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

																	"pass_credentials": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"release_name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"skip_crds": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"value_files": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"values": {
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

															"kustomize": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"common_annotations": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"common_labels": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"force_common_annotations": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"force_common_labels": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"images": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"name_prefix": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"name_suffix": {
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

															"path": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"plugin": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"env": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"value": {
																				Description:         "",
																				MarkdownDescription: "",

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

															"repo_url": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"target_revision": {
																Description:         "",
																MarkdownDescription: "",

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

													"sync_policy": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"automated": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"allow_empty": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"prune": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"self_heal": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"retry": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"backoff": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"duration": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"factor": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.Int64Type,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"max_duration": {
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

																	"limit": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"sync_options": {
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

									"values": {
										Description:         "",
										MarkdownDescription: "",

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

							"clusters": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"selector": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"match_expressions": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"operator": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"values": {
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

											"match_labels": {
												Description:         "",
												MarkdownDescription: "",

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

									"template": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"metadata": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"finalizers": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

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
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},

											"spec": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"destination": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

															"server": {
																Description:         "",
																MarkdownDescription: "",

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

													"ignore_differences": {
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

															"jq_path_expressions": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"json_pointers": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"kind": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"managed_fields_managers": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

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
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"info": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"value": {
																Description:         "",
																MarkdownDescription: "",

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

													"project": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"revision_history_limit": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"source": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"chart": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"directory": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"exclude": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"include": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"jsonnet": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"ext_vars": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"code": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"name": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"value": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"libs": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"tlas": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"code": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"name": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"value": {
																						Description:         "",
																						MarkdownDescription: "",

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

																	"recurse": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"helm": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"file_parameters": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"path": {
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

																	"ignore_missing_value_files": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"parameters": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"force_string": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

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

																			"value": {
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

																	"pass_credentials": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"release_name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"skip_crds": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"value_files": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"values": {
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

															"kustomize": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"common_annotations": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"common_labels": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"force_common_annotations": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"force_common_labels": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"images": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"name_prefix": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"name_suffix": {
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

															"path": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"plugin": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"env": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"value": {
																				Description:         "",
																				MarkdownDescription: "",

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

															"repo_url": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"target_revision": {
																Description:         "",
																MarkdownDescription: "",

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

													"sync_policy": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"automated": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"allow_empty": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"prune": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"self_heal": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"retry": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"backoff": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"duration": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"factor": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.Int64Type,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"max_duration": {
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

																	"limit": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"sync_options": {
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

									"values": {
										Description:         "",
										MarkdownDescription: "",

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

							"git": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"directories": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"exclude": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"path": {
												Description:         "",
												MarkdownDescription: "",

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

									"files": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"path": {
												Description:         "",
												MarkdownDescription: "",

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

									"repo_url": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"requeue_after_seconds": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"revision": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"template": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"metadata": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"finalizers": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

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
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},

											"spec": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"destination": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

															"server": {
																Description:         "",
																MarkdownDescription: "",

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

													"ignore_differences": {
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

															"jq_path_expressions": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"json_pointers": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"kind": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"managed_fields_managers": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

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
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"info": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"value": {
																Description:         "",
																MarkdownDescription: "",

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

													"project": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"revision_history_limit": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"source": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"chart": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"directory": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"exclude": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"include": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"jsonnet": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"ext_vars": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"code": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"name": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"value": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"libs": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"tlas": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"code": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"name": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"value": {
																						Description:         "",
																						MarkdownDescription: "",

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

																	"recurse": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"helm": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"file_parameters": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"path": {
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

																	"ignore_missing_value_files": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"parameters": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"force_string": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

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

																			"value": {
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

																	"pass_credentials": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"release_name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"skip_crds": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"value_files": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"values": {
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

															"kustomize": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"common_annotations": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"common_labels": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"force_common_annotations": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"force_common_labels": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"images": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"name_prefix": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"name_suffix": {
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

															"path": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"plugin": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"env": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"value": {
																				Description:         "",
																				MarkdownDescription: "",

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

															"repo_url": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"target_revision": {
																Description:         "",
																MarkdownDescription: "",

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

													"sync_policy": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"automated": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"allow_empty": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"prune": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"self_heal": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"retry": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"backoff": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"duration": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"factor": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.Int64Type,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"max_duration": {
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

																	"limit": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"sync_options": {
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

							"list": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"elements": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: true,
										Optional: false,
										Computed: false,
									},

									"template": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"metadata": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"finalizers": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

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
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},

											"spec": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"destination": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

															"server": {
																Description:         "",
																MarkdownDescription: "",

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

													"ignore_differences": {
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

															"jq_path_expressions": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"json_pointers": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"kind": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"managed_fields_managers": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

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
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"info": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"value": {
																Description:         "",
																MarkdownDescription: "",

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

													"project": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"revision_history_limit": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"source": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"chart": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"directory": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"exclude": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"include": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"jsonnet": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"ext_vars": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"code": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"name": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"value": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"libs": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"tlas": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"code": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"name": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"value": {
																						Description:         "",
																						MarkdownDescription: "",

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

																	"recurse": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"helm": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"file_parameters": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"path": {
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

																	"ignore_missing_value_files": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"parameters": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"force_string": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

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

																			"value": {
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

																	"pass_credentials": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"release_name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"skip_crds": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"value_files": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"values": {
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

															"kustomize": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"common_annotations": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"common_labels": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"force_common_annotations": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"force_common_labels": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"images": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"name_prefix": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"name_suffix": {
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

															"path": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"plugin": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"env": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"value": {
																				Description:         "",
																				MarkdownDescription: "",

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

															"repo_url": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"target_revision": {
																Description:         "",
																MarkdownDescription: "",

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

													"sync_policy": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"automated": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"allow_empty": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"prune": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"self_heal": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"retry": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"backoff": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"duration": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"factor": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.Int64Type,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"max_duration": {
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

																	"limit": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"sync_options": {
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

							"matrix": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"generators": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"cluster_decision_resource": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"config_map_ref": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"label_selector": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"match_expressions": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"operator": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"values": {
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

															"match_labels": {
																Description:         "",
																MarkdownDescription: "",

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

													"name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"requeue_after_seconds": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"template": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"metadata": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"annotations": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"finalizers": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"labels": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

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
																}),

																Required: true,
																Optional: false,
																Computed: false,
															},

															"spec": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"destination": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

																			"server": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"ignore_differences": {
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

																			"jq_path_expressions": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"json_pointers": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"kind": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"managed_fields_managers": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

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
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"info": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"value": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"project": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"revision_history_limit": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.Int64Type,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"source": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"chart": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"directory": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"exclude": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"include": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"jsonnet": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"ext_vars": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																									"code": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"name": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.StringType,

																										Required: true,
																										Optional: false,
																										Computed: false,
																									},

																									"value": {
																										Description:         "",
																										MarkdownDescription: "",

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

																							"libs": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.ListType{ElemType: types.StringType},

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"tlas": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																									"code": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"name": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.StringType,

																										Required: true,
																										Optional: false,
																										Computed: false,
																									},

																									"value": {
																										Description:         "",
																										MarkdownDescription: "",

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

																					"recurse": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"helm": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"file_parameters": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"name": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"path": {
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

																					"ignore_missing_value_files": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"parameters": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"force_string": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.BoolType,

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

																							"value": {
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

																					"pass_credentials": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"release_name": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"skip_crds": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"value_files": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.ListType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"values": {
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

																			"kustomize": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"common_annotations": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"common_labels": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"force_common_annotations": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"force_common_labels": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"images": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.ListType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"name_prefix": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"name_suffix": {
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

																			"path": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"plugin": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"env": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"name": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: true,
																								Optional: false,
																								Computed: false,
																							},

																							"value": {
																								Description:         "",
																								MarkdownDescription: "",

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

																			"repo_url": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"target_revision": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"sync_policy": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"automated": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"allow_empty": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"prune": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"self_heal": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"retry": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"backoff": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"duration": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"factor": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.Int64Type,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"max_duration": {
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

																					"limit": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"sync_options": {
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

													"values": {
														Description:         "",
														MarkdownDescription: "",

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

											"clusters": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"selector": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"match_expressions": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"operator": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"values": {
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

															"match_labels": {
																Description:         "",
																MarkdownDescription: "",

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

													"template": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"metadata": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"annotations": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"finalizers": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"labels": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

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
																}),

																Required: true,
																Optional: false,
																Computed: false,
															},

															"spec": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"destination": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

																			"server": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"ignore_differences": {
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

																			"jq_path_expressions": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"json_pointers": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"kind": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"managed_fields_managers": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

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
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"info": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"value": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"project": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"revision_history_limit": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.Int64Type,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"source": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"chart": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"directory": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"exclude": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"include": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"jsonnet": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"ext_vars": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																									"code": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"name": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.StringType,

																										Required: true,
																										Optional: false,
																										Computed: false,
																									},

																									"value": {
																										Description:         "",
																										MarkdownDescription: "",

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

																							"libs": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.ListType{ElemType: types.StringType},

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"tlas": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																									"code": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"name": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.StringType,

																										Required: true,
																										Optional: false,
																										Computed: false,
																									},

																									"value": {
																										Description:         "",
																										MarkdownDescription: "",

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

																					"recurse": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"helm": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"file_parameters": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"name": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"path": {
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

																					"ignore_missing_value_files": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"parameters": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"force_string": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.BoolType,

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

																							"value": {
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

																					"pass_credentials": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"release_name": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"skip_crds": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"value_files": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.ListType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"values": {
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

																			"kustomize": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"common_annotations": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"common_labels": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"force_common_annotations": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"force_common_labels": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"images": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.ListType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"name_prefix": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"name_suffix": {
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

																			"path": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"plugin": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"env": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"name": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: true,
																								Optional: false,
																								Computed: false,
																							},

																							"value": {
																								Description:         "",
																								MarkdownDescription: "",

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

																			"repo_url": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"target_revision": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"sync_policy": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"automated": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"allow_empty": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"prune": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"self_heal": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"retry": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"backoff": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"duration": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"factor": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.Int64Type,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"max_duration": {
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

																					"limit": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"sync_options": {
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

													"values": {
														Description:         "",
														MarkdownDescription: "",

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

											"git": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"directories": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"exclude": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"path": {
																Description:         "",
																MarkdownDescription: "",

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

													"files": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"path": {
																Description:         "",
																MarkdownDescription: "",

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

													"repo_url": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"requeue_after_seconds": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"revision": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"template": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"metadata": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"annotations": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"finalizers": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"labels": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

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
																}),

																Required: true,
																Optional: false,
																Computed: false,
															},

															"spec": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"destination": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

																			"server": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"ignore_differences": {
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

																			"jq_path_expressions": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"json_pointers": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"kind": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"managed_fields_managers": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

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
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"info": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"value": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"project": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"revision_history_limit": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.Int64Type,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"source": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"chart": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"directory": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"exclude": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"include": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"jsonnet": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"ext_vars": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																									"code": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"name": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.StringType,

																										Required: true,
																										Optional: false,
																										Computed: false,
																									},

																									"value": {
																										Description:         "",
																										MarkdownDescription: "",

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

																							"libs": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.ListType{ElemType: types.StringType},

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"tlas": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																									"code": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"name": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.StringType,

																										Required: true,
																										Optional: false,
																										Computed: false,
																									},

																									"value": {
																										Description:         "",
																										MarkdownDescription: "",

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

																					"recurse": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"helm": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"file_parameters": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"name": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"path": {
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

																					"ignore_missing_value_files": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"parameters": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"force_string": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.BoolType,

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

																							"value": {
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

																					"pass_credentials": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"release_name": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"skip_crds": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"value_files": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.ListType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"values": {
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

																			"kustomize": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"common_annotations": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"common_labels": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"force_common_annotations": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"force_common_labels": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"images": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.ListType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"name_prefix": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"name_suffix": {
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

																			"path": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"plugin": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"env": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"name": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: true,
																								Optional: false,
																								Computed: false,
																							},

																							"value": {
																								Description:         "",
																								MarkdownDescription: "",

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

																			"repo_url": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"target_revision": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"sync_policy": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"automated": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"allow_empty": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"prune": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"self_heal": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"retry": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"backoff": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"duration": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"factor": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.Int64Type,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"max_duration": {
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

																					"limit": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"sync_options": {
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

											"list": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"elements": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: true,
														Optional: false,
														Computed: false,
													},

													"template": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"metadata": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"annotations": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"finalizers": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"labels": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

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
																}),

																Required: true,
																Optional: false,
																Computed: false,
															},

															"spec": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"destination": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

																			"server": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"ignore_differences": {
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

																			"jq_path_expressions": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"json_pointers": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"kind": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"managed_fields_managers": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

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
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"info": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"value": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"project": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"revision_history_limit": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.Int64Type,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"source": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"chart": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"directory": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"exclude": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"include": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"jsonnet": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"ext_vars": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																									"code": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"name": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.StringType,

																										Required: true,
																										Optional: false,
																										Computed: false,
																									},

																									"value": {
																										Description:         "",
																										MarkdownDescription: "",

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

																							"libs": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.ListType{ElemType: types.StringType},

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"tlas": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																									"code": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"name": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.StringType,

																										Required: true,
																										Optional: false,
																										Computed: false,
																									},

																									"value": {
																										Description:         "",
																										MarkdownDescription: "",

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

																					"recurse": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"helm": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"file_parameters": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"name": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"path": {
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

																					"ignore_missing_value_files": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"parameters": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"force_string": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.BoolType,

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

																							"value": {
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

																					"pass_credentials": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"release_name": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"skip_crds": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"value_files": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.ListType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"values": {
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

																			"kustomize": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"common_annotations": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"common_labels": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"force_common_annotations": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"force_common_labels": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"images": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.ListType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"name_prefix": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"name_suffix": {
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

																			"path": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"plugin": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"env": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"name": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: true,
																								Optional: false,
																								Computed: false,
																							},

																							"value": {
																								Description:         "",
																								MarkdownDescription: "",

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

																			"repo_url": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"target_revision": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"sync_policy": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"automated": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"allow_empty": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"prune": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"self_heal": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"retry": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"backoff": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"duration": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"factor": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.Int64Type,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"max_duration": {
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

																					"limit": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"sync_options": {
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

											"matrix": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"merge": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"pull_request": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"bitbucket_server": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"api": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"basic_auth": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"password_ref": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"secret_name": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"username": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"project": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"repo": {
																Description:         "",
																MarkdownDescription: "",

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

													"filters": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"branch_match": {
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

													"gitea": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"api": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"insecure": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"owner": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"repo": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"token_ref": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"secret_name": {
																		Description:         "",
																		MarkdownDescription: "",

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

													"github": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"api": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"app_secret_name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"labels": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"owner": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"repo": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"token_ref": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"secret_name": {
																		Description:         "",
																		MarkdownDescription: "",

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

													"gitlab": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"api": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"labels": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"project": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"pull_request_state": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"token_ref": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"secret_name": {
																		Description:         "",
																		MarkdownDescription: "",

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

													"requeue_after_seconds": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"template": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"metadata": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"annotations": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"finalizers": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"labels": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

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
																}),

																Required: true,
																Optional: false,
																Computed: false,
															},

															"spec": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"destination": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

																			"server": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"ignore_differences": {
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

																			"jq_path_expressions": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"json_pointers": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"kind": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"managed_fields_managers": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

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
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"info": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"value": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"project": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"revision_history_limit": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.Int64Type,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"source": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"chart": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"directory": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"exclude": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"include": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"jsonnet": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"ext_vars": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																									"code": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"name": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.StringType,

																										Required: true,
																										Optional: false,
																										Computed: false,
																									},

																									"value": {
																										Description:         "",
																										MarkdownDescription: "",

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

																							"libs": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.ListType{ElemType: types.StringType},

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"tlas": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																									"code": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"name": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.StringType,

																										Required: true,
																										Optional: false,
																										Computed: false,
																									},

																									"value": {
																										Description:         "",
																										MarkdownDescription: "",

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

																					"recurse": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"helm": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"file_parameters": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"name": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"path": {
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

																					"ignore_missing_value_files": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"parameters": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"force_string": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.BoolType,

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

																							"value": {
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

																					"pass_credentials": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"release_name": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"skip_crds": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"value_files": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.ListType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"values": {
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

																			"kustomize": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"common_annotations": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"common_labels": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"force_common_annotations": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"force_common_labels": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"images": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.ListType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"name_prefix": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"name_suffix": {
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

																			"path": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"plugin": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"env": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"name": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: true,
																								Optional: false,
																								Computed: false,
																							},

																							"value": {
																								Description:         "",
																								MarkdownDescription: "",

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

																			"repo_url": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"target_revision": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"sync_policy": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"automated": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"allow_empty": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"prune": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"self_heal": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"retry": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"backoff": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"duration": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"factor": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.Int64Type,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"max_duration": {
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

																					"limit": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"sync_options": {
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

											"scm_provider": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"azure_dev_ops": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"access_token_ref": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"secret_name": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"all_branches": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"api": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"organization": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"team_project": {
																Description:         "",
																MarkdownDescription: "",

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

													"bitbucket": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"all_branches": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"app_password_ref": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"secret_name": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"owner": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"user": {
																Description:         "",
																MarkdownDescription: "",

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

													"bitbucket_server": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"all_branches": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"api": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"basic_auth": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"password_ref": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"secret_name": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"username": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"project": {
																Description:         "",
																MarkdownDescription: "",

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

													"clone_protocol": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"filters": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"branch_match": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"label_match": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"paths_do_not_exist": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"paths_exist": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"repository_match": {
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

													"gitea": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"all_branches": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"api": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"insecure": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"owner": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"token_ref": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"secret_name": {
																		Description:         "",
																		MarkdownDescription: "",

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

													"github": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"all_branches": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"api": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"app_secret_name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"organization": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"token_ref": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"secret_name": {
																		Description:         "",
																		MarkdownDescription: "",

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

													"gitlab": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"all_branches": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"api": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"group": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"include_subgroups": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"token_ref": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"secret_name": {
																		Description:         "",
																		MarkdownDescription: "",

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

													"requeue_after_seconds": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"template": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"metadata": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"annotations": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"finalizers": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"labels": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

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
																}),

																Required: true,
																Optional: false,
																Computed: false,
															},

															"spec": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"destination": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

																			"server": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"ignore_differences": {
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

																			"jq_path_expressions": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"json_pointers": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"kind": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"managed_fields_managers": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

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
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"info": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"value": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"project": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"revision_history_limit": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.Int64Type,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"source": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"chart": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"directory": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"exclude": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"include": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"jsonnet": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"ext_vars": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																									"code": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"name": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.StringType,

																										Required: true,
																										Optional: false,
																										Computed: false,
																									},

																									"value": {
																										Description:         "",
																										MarkdownDescription: "",

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

																							"libs": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.ListType{ElemType: types.StringType},

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"tlas": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																									"code": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"name": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.StringType,

																										Required: true,
																										Optional: false,
																										Computed: false,
																									},

																									"value": {
																										Description:         "",
																										MarkdownDescription: "",

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

																					"recurse": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"helm": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"file_parameters": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"name": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"path": {
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

																					"ignore_missing_value_files": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"parameters": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"force_string": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.BoolType,

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

																							"value": {
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

																					"pass_credentials": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"release_name": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"skip_crds": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"value_files": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.ListType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"values": {
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

																			"kustomize": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"common_annotations": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"common_labels": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"force_common_annotations": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"force_common_labels": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"images": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.ListType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"name_prefix": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"name_suffix": {
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

																			"path": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"plugin": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"env": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"name": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: true,
																								Optional: false,
																								Computed: false,
																							},

																							"value": {
																								Description:         "",
																								MarkdownDescription: "",

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

																			"repo_url": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"target_revision": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"sync_policy": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"automated": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"allow_empty": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"prune": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"self_heal": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"retry": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"backoff": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"duration": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"factor": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.Int64Type,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"max_duration": {
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

																					"limit": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"sync_options": {
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

											"selector": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"match_expressions": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"operator": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"values": {
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

													"match_labels": {
														Description:         "",
														MarkdownDescription: "",

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

										Required: true,
										Optional: false,
										Computed: false,
									},

									"template": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"metadata": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"finalizers": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

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
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},

											"spec": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"destination": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

															"server": {
																Description:         "",
																MarkdownDescription: "",

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

													"ignore_differences": {
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

															"jq_path_expressions": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"json_pointers": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"kind": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"managed_fields_managers": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

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
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"info": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"value": {
																Description:         "",
																MarkdownDescription: "",

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

													"project": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"revision_history_limit": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"source": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"chart": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"directory": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"exclude": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"include": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"jsonnet": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"ext_vars": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"code": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"name": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"value": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"libs": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"tlas": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"code": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"name": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"value": {
																						Description:         "",
																						MarkdownDescription: "",

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

																	"recurse": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"helm": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"file_parameters": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"path": {
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

																	"ignore_missing_value_files": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"parameters": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"force_string": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

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

																			"value": {
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

																	"pass_credentials": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"release_name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"skip_crds": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"value_files": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"values": {
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

															"kustomize": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"common_annotations": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"common_labels": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"force_common_annotations": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"force_common_labels": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"images": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"name_prefix": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"name_suffix": {
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

															"path": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"plugin": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"env": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"value": {
																				Description:         "",
																				MarkdownDescription: "",

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

															"repo_url": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"target_revision": {
																Description:         "",
																MarkdownDescription: "",

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

													"sync_policy": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"automated": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"allow_empty": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"prune": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"self_heal": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"retry": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"backoff": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"duration": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"factor": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.Int64Type,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"max_duration": {
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

																	"limit": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"sync_options": {
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

							"merge": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"generators": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"cluster_decision_resource": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"config_map_ref": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"label_selector": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"match_expressions": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"operator": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"values": {
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

															"match_labels": {
																Description:         "",
																MarkdownDescription: "",

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

													"name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"requeue_after_seconds": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"template": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"metadata": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"annotations": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"finalizers": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"labels": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

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
																}),

																Required: true,
																Optional: false,
																Computed: false,
															},

															"spec": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"destination": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

																			"server": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"ignore_differences": {
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

																			"jq_path_expressions": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"json_pointers": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"kind": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"managed_fields_managers": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

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
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"info": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"value": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"project": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"revision_history_limit": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.Int64Type,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"source": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"chart": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"directory": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"exclude": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"include": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"jsonnet": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"ext_vars": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																									"code": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"name": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.StringType,

																										Required: true,
																										Optional: false,
																										Computed: false,
																									},

																									"value": {
																										Description:         "",
																										MarkdownDescription: "",

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

																							"libs": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.ListType{ElemType: types.StringType},

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"tlas": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																									"code": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"name": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.StringType,

																										Required: true,
																										Optional: false,
																										Computed: false,
																									},

																									"value": {
																										Description:         "",
																										MarkdownDescription: "",

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

																					"recurse": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"helm": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"file_parameters": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"name": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"path": {
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

																					"ignore_missing_value_files": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"parameters": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"force_string": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.BoolType,

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

																							"value": {
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

																					"pass_credentials": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"release_name": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"skip_crds": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"value_files": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.ListType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"values": {
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

																			"kustomize": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"common_annotations": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"common_labels": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"force_common_annotations": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"force_common_labels": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"images": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.ListType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"name_prefix": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"name_suffix": {
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

																			"path": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"plugin": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"env": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"name": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: true,
																								Optional: false,
																								Computed: false,
																							},

																							"value": {
																								Description:         "",
																								MarkdownDescription: "",

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

																			"repo_url": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"target_revision": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"sync_policy": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"automated": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"allow_empty": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"prune": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"self_heal": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"retry": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"backoff": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"duration": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"factor": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.Int64Type,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"max_duration": {
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

																					"limit": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"sync_options": {
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

													"values": {
														Description:         "",
														MarkdownDescription: "",

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

											"clusters": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"selector": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"match_expressions": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"operator": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"values": {
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

															"match_labels": {
																Description:         "",
																MarkdownDescription: "",

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

													"template": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"metadata": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"annotations": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"finalizers": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"labels": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

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
																}),

																Required: true,
																Optional: false,
																Computed: false,
															},

															"spec": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"destination": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

																			"server": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"ignore_differences": {
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

																			"jq_path_expressions": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"json_pointers": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"kind": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"managed_fields_managers": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

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
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"info": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"value": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"project": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"revision_history_limit": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.Int64Type,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"source": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"chart": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"directory": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"exclude": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"include": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"jsonnet": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"ext_vars": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																									"code": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"name": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.StringType,

																										Required: true,
																										Optional: false,
																										Computed: false,
																									},

																									"value": {
																										Description:         "",
																										MarkdownDescription: "",

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

																							"libs": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.ListType{ElemType: types.StringType},

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"tlas": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																									"code": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"name": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.StringType,

																										Required: true,
																										Optional: false,
																										Computed: false,
																									},

																									"value": {
																										Description:         "",
																										MarkdownDescription: "",

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

																					"recurse": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"helm": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"file_parameters": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"name": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"path": {
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

																					"ignore_missing_value_files": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"parameters": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"force_string": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.BoolType,

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

																							"value": {
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

																					"pass_credentials": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"release_name": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"skip_crds": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"value_files": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.ListType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"values": {
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

																			"kustomize": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"common_annotations": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"common_labels": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"force_common_annotations": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"force_common_labels": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"images": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.ListType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"name_prefix": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"name_suffix": {
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

																			"path": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"plugin": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"env": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"name": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: true,
																								Optional: false,
																								Computed: false,
																							},

																							"value": {
																								Description:         "",
																								MarkdownDescription: "",

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

																			"repo_url": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"target_revision": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"sync_policy": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"automated": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"allow_empty": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"prune": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"self_heal": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"retry": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"backoff": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"duration": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"factor": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.Int64Type,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"max_duration": {
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

																					"limit": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"sync_options": {
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

													"values": {
														Description:         "",
														MarkdownDescription: "",

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

											"git": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"directories": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"exclude": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"path": {
																Description:         "",
																MarkdownDescription: "",

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

													"files": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"path": {
																Description:         "",
																MarkdownDescription: "",

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

													"repo_url": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"requeue_after_seconds": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"revision": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"template": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"metadata": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"annotations": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"finalizers": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"labels": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

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
																}),

																Required: true,
																Optional: false,
																Computed: false,
															},

															"spec": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"destination": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

																			"server": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"ignore_differences": {
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

																			"jq_path_expressions": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"json_pointers": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"kind": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"managed_fields_managers": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

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
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"info": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"value": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"project": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"revision_history_limit": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.Int64Type,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"source": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"chart": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"directory": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"exclude": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"include": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"jsonnet": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"ext_vars": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																									"code": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"name": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.StringType,

																										Required: true,
																										Optional: false,
																										Computed: false,
																									},

																									"value": {
																										Description:         "",
																										MarkdownDescription: "",

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

																							"libs": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.ListType{ElemType: types.StringType},

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"tlas": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																									"code": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"name": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.StringType,

																										Required: true,
																										Optional: false,
																										Computed: false,
																									},

																									"value": {
																										Description:         "",
																										MarkdownDescription: "",

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

																					"recurse": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"helm": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"file_parameters": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"name": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"path": {
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

																					"ignore_missing_value_files": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"parameters": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"force_string": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.BoolType,

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

																							"value": {
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

																					"pass_credentials": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"release_name": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"skip_crds": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"value_files": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.ListType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"values": {
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

																			"kustomize": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"common_annotations": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"common_labels": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"force_common_annotations": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"force_common_labels": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"images": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.ListType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"name_prefix": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"name_suffix": {
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

																			"path": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"plugin": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"env": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"name": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: true,
																								Optional: false,
																								Computed: false,
																							},

																							"value": {
																								Description:         "",
																								MarkdownDescription: "",

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

																			"repo_url": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"target_revision": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"sync_policy": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"automated": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"allow_empty": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"prune": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"self_heal": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"retry": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"backoff": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"duration": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"factor": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.Int64Type,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"max_duration": {
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

																					"limit": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"sync_options": {
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

											"list": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"elements": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: true,
														Optional: false,
														Computed: false,
													},

													"template": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"metadata": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"annotations": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"finalizers": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"labels": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

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
																}),

																Required: true,
																Optional: false,
																Computed: false,
															},

															"spec": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"destination": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

																			"server": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"ignore_differences": {
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

																			"jq_path_expressions": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"json_pointers": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"kind": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"managed_fields_managers": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

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
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"info": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"value": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"project": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"revision_history_limit": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.Int64Type,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"source": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"chart": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"directory": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"exclude": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"include": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"jsonnet": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"ext_vars": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																									"code": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"name": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.StringType,

																										Required: true,
																										Optional: false,
																										Computed: false,
																									},

																									"value": {
																										Description:         "",
																										MarkdownDescription: "",

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

																							"libs": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.ListType{ElemType: types.StringType},

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"tlas": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																									"code": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"name": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.StringType,

																										Required: true,
																										Optional: false,
																										Computed: false,
																									},

																									"value": {
																										Description:         "",
																										MarkdownDescription: "",

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

																					"recurse": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"helm": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"file_parameters": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"name": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"path": {
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

																					"ignore_missing_value_files": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"parameters": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"force_string": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.BoolType,

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

																							"value": {
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

																					"pass_credentials": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"release_name": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"skip_crds": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"value_files": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.ListType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"values": {
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

																			"kustomize": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"common_annotations": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"common_labels": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"force_common_annotations": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"force_common_labels": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"images": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.ListType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"name_prefix": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"name_suffix": {
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

																			"path": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"plugin": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"env": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"name": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: true,
																								Optional: false,
																								Computed: false,
																							},

																							"value": {
																								Description:         "",
																								MarkdownDescription: "",

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

																			"repo_url": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"target_revision": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"sync_policy": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"automated": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"allow_empty": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"prune": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"self_heal": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"retry": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"backoff": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"duration": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"factor": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.Int64Type,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"max_duration": {
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

																					"limit": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"sync_options": {
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

											"matrix": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"merge": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"pull_request": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"bitbucket_server": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"api": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"basic_auth": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"password_ref": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"secret_name": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"username": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"project": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"repo": {
																Description:         "",
																MarkdownDescription: "",

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

													"filters": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"branch_match": {
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

													"gitea": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"api": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"insecure": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"owner": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"repo": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"token_ref": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"secret_name": {
																		Description:         "",
																		MarkdownDescription: "",

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

													"github": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"api": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"app_secret_name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"labels": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"owner": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"repo": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"token_ref": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"secret_name": {
																		Description:         "",
																		MarkdownDescription: "",

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

													"gitlab": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"api": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"labels": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"project": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"pull_request_state": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"token_ref": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"secret_name": {
																		Description:         "",
																		MarkdownDescription: "",

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

													"requeue_after_seconds": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"template": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"metadata": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"annotations": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"finalizers": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"labels": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

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
																}),

																Required: true,
																Optional: false,
																Computed: false,
															},

															"spec": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"destination": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

																			"server": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"ignore_differences": {
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

																			"jq_path_expressions": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"json_pointers": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"kind": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"managed_fields_managers": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

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
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"info": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"value": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"project": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"revision_history_limit": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.Int64Type,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"source": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"chart": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"directory": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"exclude": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"include": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"jsonnet": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"ext_vars": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																									"code": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"name": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.StringType,

																										Required: true,
																										Optional: false,
																										Computed: false,
																									},

																									"value": {
																										Description:         "",
																										MarkdownDescription: "",

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

																							"libs": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.ListType{ElemType: types.StringType},

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"tlas": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																									"code": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"name": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.StringType,

																										Required: true,
																										Optional: false,
																										Computed: false,
																									},

																									"value": {
																										Description:         "",
																										MarkdownDescription: "",

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

																					"recurse": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"helm": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"file_parameters": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"name": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"path": {
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

																					"ignore_missing_value_files": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"parameters": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"force_string": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.BoolType,

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

																							"value": {
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

																					"pass_credentials": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"release_name": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"skip_crds": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"value_files": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.ListType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"values": {
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

																			"kustomize": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"common_annotations": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"common_labels": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"force_common_annotations": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"force_common_labels": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"images": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.ListType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"name_prefix": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"name_suffix": {
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

																			"path": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"plugin": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"env": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"name": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: true,
																								Optional: false,
																								Computed: false,
																							},

																							"value": {
																								Description:         "",
																								MarkdownDescription: "",

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

																			"repo_url": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"target_revision": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"sync_policy": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"automated": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"allow_empty": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"prune": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"self_heal": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"retry": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"backoff": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"duration": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"factor": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.Int64Type,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"max_duration": {
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

																					"limit": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"sync_options": {
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

											"scm_provider": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"azure_dev_ops": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"access_token_ref": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"secret_name": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"all_branches": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"api": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"organization": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"team_project": {
																Description:         "",
																MarkdownDescription: "",

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

													"bitbucket": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"all_branches": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"app_password_ref": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"secret_name": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"owner": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"user": {
																Description:         "",
																MarkdownDescription: "",

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

													"bitbucket_server": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"all_branches": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"api": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"basic_auth": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"password_ref": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"secret_name": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"username": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"project": {
																Description:         "",
																MarkdownDescription: "",

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

													"clone_protocol": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"filters": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"branch_match": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"label_match": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"paths_do_not_exist": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"paths_exist": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"repository_match": {
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

													"gitea": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"all_branches": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"api": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"insecure": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"owner": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"token_ref": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"secret_name": {
																		Description:         "",
																		MarkdownDescription: "",

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

													"github": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"all_branches": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"api": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"app_secret_name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"organization": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"token_ref": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"secret_name": {
																		Description:         "",
																		MarkdownDescription: "",

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

													"gitlab": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"all_branches": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"api": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"group": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"include_subgroups": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"token_ref": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"secret_name": {
																		Description:         "",
																		MarkdownDescription: "",

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

													"requeue_after_seconds": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"template": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"metadata": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"annotations": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"finalizers": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"labels": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

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
																}),

																Required: true,
																Optional: false,
																Computed: false,
															},

															"spec": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"destination": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

																			"server": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"ignore_differences": {
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

																			"jq_path_expressions": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"json_pointers": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"kind": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"managed_fields_managers": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

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
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"info": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"value": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"project": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"revision_history_limit": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.Int64Type,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"source": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"chart": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"directory": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"exclude": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"include": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"jsonnet": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"ext_vars": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																									"code": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"name": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.StringType,

																										Required: true,
																										Optional: false,
																										Computed: false,
																									},

																									"value": {
																										Description:         "",
																										MarkdownDescription: "",

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

																							"libs": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.ListType{ElemType: types.StringType},

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"tlas": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																									"code": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"name": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.StringType,

																										Required: true,
																										Optional: false,
																										Computed: false,
																									},

																									"value": {
																										Description:         "",
																										MarkdownDescription: "",

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

																					"recurse": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"helm": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"file_parameters": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"name": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"path": {
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

																					"ignore_missing_value_files": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"parameters": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"force_string": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.BoolType,

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

																							"value": {
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

																					"pass_credentials": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"release_name": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"skip_crds": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"value_files": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.ListType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"values": {
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

																			"kustomize": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"common_annotations": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"common_labels": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"force_common_annotations": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"force_common_labels": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"images": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.ListType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"name_prefix": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"name_suffix": {
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

																			"path": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"plugin": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"env": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"name": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: true,
																								Optional: false,
																								Computed: false,
																							},

																							"value": {
																								Description:         "",
																								MarkdownDescription: "",

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

																			"repo_url": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"target_revision": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"sync_policy": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"automated": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"allow_empty": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"prune": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"self_heal": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"retry": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"backoff": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"duration": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"factor": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.Int64Type,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"max_duration": {
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

																					"limit": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"sync_options": {
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

											"selector": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"match_expressions": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"operator": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"values": {
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

													"match_labels": {
														Description:         "",
														MarkdownDescription: "",

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

										Required: true,
										Optional: false,
										Computed: false,
									},

									"merge_keys": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: true,
										Optional: false,
										Computed: false,
									},

									"template": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"metadata": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"finalizers": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

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
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},

											"spec": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"destination": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

															"server": {
																Description:         "",
																MarkdownDescription: "",

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

													"ignore_differences": {
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

															"jq_path_expressions": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"json_pointers": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"kind": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"managed_fields_managers": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

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
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"info": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"value": {
																Description:         "",
																MarkdownDescription: "",

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

													"project": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"revision_history_limit": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"source": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"chart": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"directory": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"exclude": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"include": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"jsonnet": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"ext_vars": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"code": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"name": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"value": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"libs": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"tlas": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"code": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"name": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"value": {
																						Description:         "",
																						MarkdownDescription: "",

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

																	"recurse": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"helm": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"file_parameters": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"path": {
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

																	"ignore_missing_value_files": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"parameters": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"force_string": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

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

																			"value": {
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

																	"pass_credentials": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"release_name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"skip_crds": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"value_files": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"values": {
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

															"kustomize": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"common_annotations": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"common_labels": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"force_common_annotations": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"force_common_labels": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"images": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"name_prefix": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"name_suffix": {
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

															"path": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"plugin": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"env": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"value": {
																				Description:         "",
																				MarkdownDescription: "",

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

															"repo_url": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"target_revision": {
																Description:         "",
																MarkdownDescription: "",

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

													"sync_policy": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"automated": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"allow_empty": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"prune": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"self_heal": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"retry": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"backoff": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"duration": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"factor": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.Int64Type,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"max_duration": {
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

																	"limit": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"sync_options": {
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

							"pull_request": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"bitbucket_server": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"api": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"basic_auth": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"password_ref": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"secret_name": {
																Description:         "",
																MarkdownDescription: "",

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

													"username": {
														Description:         "",
														MarkdownDescription: "",

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

											"project": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"repo": {
												Description:         "",
												MarkdownDescription: "",

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

									"filters": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"branch_match": {
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

									"gitea": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"api": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"insecure": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"owner": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"repo": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"token_ref": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"secret_name": {
														Description:         "",
														MarkdownDescription: "",

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

									"github": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"api": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"app_secret_name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"labels": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"owner": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"repo": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"token_ref": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"secret_name": {
														Description:         "",
														MarkdownDescription: "",

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

									"gitlab": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"api": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"labels": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"project": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"pull_request_state": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"token_ref": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"secret_name": {
														Description:         "",
														MarkdownDescription: "",

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

									"requeue_after_seconds": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"template": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"metadata": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"finalizers": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

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
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},

											"spec": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"destination": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

															"server": {
																Description:         "",
																MarkdownDescription: "",

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

													"ignore_differences": {
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

															"jq_path_expressions": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"json_pointers": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"kind": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"managed_fields_managers": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

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
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"info": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"value": {
																Description:         "",
																MarkdownDescription: "",

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

													"project": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"revision_history_limit": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"source": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"chart": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"directory": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"exclude": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"include": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"jsonnet": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"ext_vars": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"code": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"name": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"value": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"libs": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"tlas": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"code": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"name": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"value": {
																						Description:         "",
																						MarkdownDescription: "",

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

																	"recurse": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"helm": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"file_parameters": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"path": {
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

																	"ignore_missing_value_files": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"parameters": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"force_string": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

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

																			"value": {
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

																	"pass_credentials": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"release_name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"skip_crds": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"value_files": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"values": {
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

															"kustomize": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"common_annotations": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"common_labels": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"force_common_annotations": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"force_common_labels": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"images": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"name_prefix": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"name_suffix": {
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

															"path": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"plugin": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"env": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"value": {
																				Description:         "",
																				MarkdownDescription: "",

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

															"repo_url": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"target_revision": {
																Description:         "",
																MarkdownDescription: "",

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

													"sync_policy": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"automated": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"allow_empty": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"prune": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"self_heal": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"retry": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"backoff": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"duration": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"factor": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.Int64Type,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"max_duration": {
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

																	"limit": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"sync_options": {
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

							"scm_provider": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"azure_dev_ops": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"access_token_ref": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"secret_name": {
														Description:         "",
														MarkdownDescription: "",

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

											"all_branches": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"api": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"organization": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"team_project": {
												Description:         "",
												MarkdownDescription: "",

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

									"bitbucket": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"all_branches": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"app_password_ref": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"secret_name": {
														Description:         "",
														MarkdownDescription: "",

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

											"owner": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"user": {
												Description:         "",
												MarkdownDescription: "",

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

									"bitbucket_server": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"all_branches": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"api": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"basic_auth": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"password_ref": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"secret_name": {
																Description:         "",
																MarkdownDescription: "",

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

													"username": {
														Description:         "",
														MarkdownDescription: "",

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

											"project": {
												Description:         "",
												MarkdownDescription: "",

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

									"clone_protocol": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"filters": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"branch_match": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"label_match": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"paths_do_not_exist": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"paths_exist": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"repository_match": {
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

									"gitea": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"all_branches": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"api": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"insecure": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"owner": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"token_ref": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"secret_name": {
														Description:         "",
														MarkdownDescription: "",

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

									"github": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"all_branches": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"api": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"app_secret_name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"organization": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"token_ref": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"secret_name": {
														Description:         "",
														MarkdownDescription: "",

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

									"gitlab": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"all_branches": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"api": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"group": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"include_subgroups": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"token_ref": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"secret_name": {
														Description:         "",
														MarkdownDescription: "",

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

									"requeue_after_seconds": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"template": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"metadata": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"finalizers": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

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
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},

											"spec": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"destination": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

															"server": {
																Description:         "",
																MarkdownDescription: "",

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

													"ignore_differences": {
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

															"jq_path_expressions": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"json_pointers": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"kind": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"managed_fields_managers": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

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
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"info": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"value": {
																Description:         "",
																MarkdownDescription: "",

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

													"project": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"revision_history_limit": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"source": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"chart": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"directory": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"exclude": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"include": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"jsonnet": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"ext_vars": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"code": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"name": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"value": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"libs": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"tlas": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"code": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"name": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"value": {
																						Description:         "",
																						MarkdownDescription: "",

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

																	"recurse": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"helm": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"file_parameters": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"path": {
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

																	"ignore_missing_value_files": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"parameters": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"force_string": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

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

																			"value": {
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

																	"pass_credentials": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"release_name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"skip_crds": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"value_files": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"values": {
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

															"kustomize": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"common_annotations": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"common_labels": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"force_common_annotations": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"force_common_labels": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"images": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"name_prefix": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"name_suffix": {
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

															"path": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"plugin": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"env": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"value": {
																				Description:         "",
																				MarkdownDescription: "",

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

															"repo_url": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"target_revision": {
																Description:         "",
																MarkdownDescription: "",

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

													"sync_policy": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"automated": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"allow_empty": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"prune": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"self_heal": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"retry": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"backoff": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"duration": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"factor": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.Int64Type,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"max_duration": {
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

																	"limit": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"sync_options": {
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

							"selector": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"match_expressions": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"operator": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"values": {
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

									"match_labels": {
										Description:         "",
										MarkdownDescription: "",

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

						Required: true,
						Optional: false,
						Computed: false,
					},

					"go_template": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"sync_policy": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"preserve_resources_on_deletion": {
								Description:         "",
								MarkdownDescription: "",

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

					"template": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"metadata": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"annotations": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"finalizers": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"labels": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.MapType{ElemType: types.StringType},

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
								}),

								Required: true,
								Optional: false,
								Computed: false,
							},

							"spec": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"destination": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

											"server": {
												Description:         "",
												MarkdownDescription: "",

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

									"ignore_differences": {
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

											"jq_path_expressions": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"json_pointers": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"kind": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"managed_fields_managers": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"info": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"value": {
												Description:         "",
												MarkdownDescription: "",

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

									"project": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"revision_history_limit": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"source": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"chart": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"directory": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"exclude": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"include": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"jsonnet": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"ext_vars": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"code": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"libs": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"tlas": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"code": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

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

													"recurse": {
														Description:         "",
														MarkdownDescription: "",

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

											"helm": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"file_parameters": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"path": {
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

													"ignore_missing_value_files": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"parameters": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"force_string": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

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

															"value": {
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

													"pass_credentials": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"release_name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"skip_crds": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"value_files": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"values": {
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

											"kustomize": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"common_annotations": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"common_labels": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"force_common_annotations": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"force_common_labels": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"images": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name_prefix": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name_suffix": {
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

											"path": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"plugin": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"env": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"value": {
																Description:         "",
																MarkdownDescription: "",

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

											"repo_url": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"target_revision": {
												Description:         "",
												MarkdownDescription: "",

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

									"sync_policy": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"automated": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"allow_empty": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"prune": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"self_heal": {
														Description:         "",
														MarkdownDescription: "",

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

											"retry": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"backoff": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"duration": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"factor": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"max_duration": {
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

													"limit": {
														Description:         "",
														MarkdownDescription: "",

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

											"sync_options": {
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
								}),

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
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

func (r *ArgoprojIoApplicationSetV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_argoproj_io_application_set_v1alpha1")

	var state ArgoprojIoApplicationSetV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ArgoprojIoApplicationSetV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("argoproj.io/v1alpha1")
	goModel.Kind = utilities.Ptr("ApplicationSet")

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

func (r *ArgoprojIoApplicationSetV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_argoproj_io_application_set_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *ArgoprojIoApplicationSetV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_argoproj_io_application_set_v1alpha1")

	var state ArgoprojIoApplicationSetV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ArgoprojIoApplicationSetV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("argoproj.io/v1alpha1")
	goModel.Kind = utilities.Ptr("ApplicationSet")

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

func (r *ArgoprojIoApplicationSetV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_argoproj_io_application_set_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
