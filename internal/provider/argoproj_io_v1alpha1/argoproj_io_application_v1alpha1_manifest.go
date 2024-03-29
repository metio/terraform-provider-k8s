/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package argoproj_io_v1alpha1

import (
	"context"
	"fmt"
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
	_ datasource.DataSource = &ArgoprojIoApplicationV1Alpha1Manifest{}
)

func NewArgoprojIoApplicationV1Alpha1Manifest() datasource.DataSource {
	return &ArgoprojIoApplicationV1Alpha1Manifest{}
}

type ArgoprojIoApplicationV1Alpha1Manifest struct{}

type ArgoprojIoApplicationV1Alpha1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Operation *struct {
		Info *[]struct {
			Name  *string `tfsdk:"name" json:"name,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"info" json:"info,omitempty"`
		InitiatedBy *struct {
			Automated *bool   `tfsdk:"automated" json:"automated,omitempty"`
			Username  *string `tfsdk:"username" json:"username,omitempty"`
		} `tfsdk:"initiated_by" json:"initiatedBy,omitempty"`
		Retry *struct {
			Backoff *struct {
				Duration    *string `tfsdk:"duration" json:"duration,omitempty"`
				Factor      *int64  `tfsdk:"factor" json:"factor,omitempty"`
				MaxDuration *string `tfsdk:"max_duration" json:"maxDuration,omitempty"`
			} `tfsdk:"backoff" json:"backoff,omitempty"`
			Limit *int64 `tfsdk:"limit" json:"limit,omitempty"`
		} `tfsdk:"retry" json:"retry,omitempty"`
		Sync *struct {
			DryRun    *bool     `tfsdk:"dry_run" json:"dryRun,omitempty"`
			Manifests *[]string `tfsdk:"manifests" json:"manifests,omitempty"`
			Prune     *bool     `tfsdk:"prune" json:"prune,omitempty"`
			Resources *[]struct {
				Group     *string `tfsdk:"group" json:"group,omitempty"`
				Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
			Revision  *string   `tfsdk:"revision" json:"revision,omitempty"`
			Revisions *[]string `tfsdk:"revisions" json:"revisions,omitempty"`
			Source    *struct {
				Chart     *string `tfsdk:"chart" json:"chart,omitempty"`
				Directory *struct {
					Exclude *string `tfsdk:"exclude" json:"exclude,omitempty"`
					Include *string `tfsdk:"include" json:"include,omitempty"`
					Jsonnet *struct {
						ExtVars *[]struct {
							Code  *bool   `tfsdk:"code" json:"code,omitempty"`
							Name  *string `tfsdk:"name" json:"name,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"ext_vars" json:"extVars,omitempty"`
						Libs *[]string `tfsdk:"libs" json:"libs,omitempty"`
						Tlas *[]struct {
							Code  *bool   `tfsdk:"code" json:"code,omitempty"`
							Name  *string `tfsdk:"name" json:"name,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"tlas" json:"tlas,omitempty"`
					} `tfsdk:"jsonnet" json:"jsonnet,omitempty"`
					Recurse *bool `tfsdk:"recurse" json:"recurse,omitempty"`
				} `tfsdk:"directory" json:"directory,omitempty"`
				Helm *struct {
					FileParameters *[]struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
						Path *string `tfsdk:"path" json:"path,omitempty"`
					} `tfsdk:"file_parameters" json:"fileParameters,omitempty"`
					IgnoreMissingValueFiles *bool `tfsdk:"ignore_missing_value_files" json:"ignoreMissingValueFiles,omitempty"`
					Parameters              *[]struct {
						ForceString *bool   `tfsdk:"force_string" json:"forceString,omitempty"`
						Name        *string `tfsdk:"name" json:"name,omitempty"`
						Value       *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"parameters" json:"parameters,omitempty"`
					PassCredentials *bool              `tfsdk:"pass_credentials" json:"passCredentials,omitempty"`
					ReleaseName     *string            `tfsdk:"release_name" json:"releaseName,omitempty"`
					SkipCrds        *bool              `tfsdk:"skip_crds" json:"skipCrds,omitempty"`
					ValueFiles      *[]string          `tfsdk:"value_files" json:"valueFiles,omitempty"`
					Values          *string            `tfsdk:"values" json:"values,omitempty"`
					ValuesObject    *map[string]string `tfsdk:"values_object" json:"valuesObject,omitempty"`
					Version         *string            `tfsdk:"version" json:"version,omitempty"`
				} `tfsdk:"helm" json:"helm,omitempty"`
				Kustomize *struct {
					CommonAnnotations         *map[string]string `tfsdk:"common_annotations" json:"commonAnnotations,omitempty"`
					CommonAnnotationsEnvsubst *bool              `tfsdk:"common_annotations_envsubst" json:"commonAnnotationsEnvsubst,omitempty"`
					CommonLabels              *map[string]string `tfsdk:"common_labels" json:"commonLabels,omitempty"`
					Components                *[]string          `tfsdk:"components" json:"components,omitempty"`
					ForceCommonAnnotations    *bool              `tfsdk:"force_common_annotations" json:"forceCommonAnnotations,omitempty"`
					ForceCommonLabels         *bool              `tfsdk:"force_common_labels" json:"forceCommonLabels,omitempty"`
					Images                    *[]string          `tfsdk:"images" json:"images,omitempty"`
					NamePrefix                *string            `tfsdk:"name_prefix" json:"namePrefix,omitempty"`
					NameSuffix                *string            `tfsdk:"name_suffix" json:"nameSuffix,omitempty"`
					Namespace                 *string            `tfsdk:"namespace" json:"namespace,omitempty"`
					Patches                   *[]struct {
						Options *map[string]string `tfsdk:"options" json:"options,omitempty"`
						Patch   *string            `tfsdk:"patch" json:"patch,omitempty"`
						Path    *string            `tfsdk:"path" json:"path,omitempty"`
						Target  *struct {
							AnnotationSelector *string `tfsdk:"annotation_selector" json:"annotationSelector,omitempty"`
							Group              *string `tfsdk:"group" json:"group,omitempty"`
							Kind               *string `tfsdk:"kind" json:"kind,omitempty"`
							LabelSelector      *string `tfsdk:"label_selector" json:"labelSelector,omitempty"`
							Name               *string `tfsdk:"name" json:"name,omitempty"`
							Namespace          *string `tfsdk:"namespace" json:"namespace,omitempty"`
							Version            *string `tfsdk:"version" json:"version,omitempty"`
						} `tfsdk:"target" json:"target,omitempty"`
					} `tfsdk:"patches" json:"patches,omitempty"`
					Replicas *[]struct {
						Count *string `tfsdk:"count" json:"count,omitempty"`
						Name  *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"replicas" json:"replicas,omitempty"`
					Version *string `tfsdk:"version" json:"version,omitempty"`
				} `tfsdk:"kustomize" json:"kustomize,omitempty"`
				Path   *string `tfsdk:"path" json:"path,omitempty"`
				Plugin *struct {
					Env *[]struct {
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"env" json:"env,omitempty"`
					Name       *string `tfsdk:"name" json:"name,omitempty"`
					Parameters *[]struct {
						Array  *[]string          `tfsdk:"array" json:"array,omitempty"`
						Map    *map[string]string `tfsdk:"map" json:"map,omitempty"`
						Name   *string            `tfsdk:"name" json:"name,omitempty"`
						String *string            `tfsdk:"string" json:"string,omitempty"`
					} `tfsdk:"parameters" json:"parameters,omitempty"`
				} `tfsdk:"plugin" json:"plugin,omitempty"`
				Ref            *string `tfsdk:"ref" json:"ref,omitempty"`
				RepoURL        *string `tfsdk:"repo_url" json:"repoURL,omitempty"`
				TargetRevision *string `tfsdk:"target_revision" json:"targetRevision,omitempty"`
			} `tfsdk:"source" json:"source,omitempty"`
			Sources *[]struct {
				Chart     *string `tfsdk:"chart" json:"chart,omitempty"`
				Directory *struct {
					Exclude *string `tfsdk:"exclude" json:"exclude,omitempty"`
					Include *string `tfsdk:"include" json:"include,omitempty"`
					Jsonnet *struct {
						ExtVars *[]struct {
							Code  *bool   `tfsdk:"code" json:"code,omitempty"`
							Name  *string `tfsdk:"name" json:"name,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"ext_vars" json:"extVars,omitempty"`
						Libs *[]string `tfsdk:"libs" json:"libs,omitempty"`
						Tlas *[]struct {
							Code  *bool   `tfsdk:"code" json:"code,omitempty"`
							Name  *string `tfsdk:"name" json:"name,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"tlas" json:"tlas,omitempty"`
					} `tfsdk:"jsonnet" json:"jsonnet,omitempty"`
					Recurse *bool `tfsdk:"recurse" json:"recurse,omitempty"`
				} `tfsdk:"directory" json:"directory,omitempty"`
				Helm *struct {
					FileParameters *[]struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
						Path *string `tfsdk:"path" json:"path,omitempty"`
					} `tfsdk:"file_parameters" json:"fileParameters,omitempty"`
					IgnoreMissingValueFiles *bool `tfsdk:"ignore_missing_value_files" json:"ignoreMissingValueFiles,omitempty"`
					Parameters              *[]struct {
						ForceString *bool   `tfsdk:"force_string" json:"forceString,omitempty"`
						Name        *string `tfsdk:"name" json:"name,omitempty"`
						Value       *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"parameters" json:"parameters,omitempty"`
					PassCredentials *bool              `tfsdk:"pass_credentials" json:"passCredentials,omitempty"`
					ReleaseName     *string            `tfsdk:"release_name" json:"releaseName,omitempty"`
					SkipCrds        *bool              `tfsdk:"skip_crds" json:"skipCrds,omitempty"`
					ValueFiles      *[]string          `tfsdk:"value_files" json:"valueFiles,omitempty"`
					Values          *string            `tfsdk:"values" json:"values,omitempty"`
					ValuesObject    *map[string]string `tfsdk:"values_object" json:"valuesObject,omitempty"`
					Version         *string            `tfsdk:"version" json:"version,omitempty"`
				} `tfsdk:"helm" json:"helm,omitempty"`
				Kustomize *struct {
					CommonAnnotations         *map[string]string `tfsdk:"common_annotations" json:"commonAnnotations,omitempty"`
					CommonAnnotationsEnvsubst *bool              `tfsdk:"common_annotations_envsubst" json:"commonAnnotationsEnvsubst,omitempty"`
					CommonLabels              *map[string]string `tfsdk:"common_labels" json:"commonLabels,omitempty"`
					Components                *[]string          `tfsdk:"components" json:"components,omitempty"`
					ForceCommonAnnotations    *bool              `tfsdk:"force_common_annotations" json:"forceCommonAnnotations,omitempty"`
					ForceCommonLabels         *bool              `tfsdk:"force_common_labels" json:"forceCommonLabels,omitempty"`
					Images                    *[]string          `tfsdk:"images" json:"images,omitempty"`
					NamePrefix                *string            `tfsdk:"name_prefix" json:"namePrefix,omitempty"`
					NameSuffix                *string            `tfsdk:"name_suffix" json:"nameSuffix,omitempty"`
					Namespace                 *string            `tfsdk:"namespace" json:"namespace,omitempty"`
					Patches                   *[]struct {
						Options *map[string]string `tfsdk:"options" json:"options,omitempty"`
						Patch   *string            `tfsdk:"patch" json:"patch,omitempty"`
						Path    *string            `tfsdk:"path" json:"path,omitempty"`
						Target  *struct {
							AnnotationSelector *string `tfsdk:"annotation_selector" json:"annotationSelector,omitempty"`
							Group              *string `tfsdk:"group" json:"group,omitempty"`
							Kind               *string `tfsdk:"kind" json:"kind,omitempty"`
							LabelSelector      *string `tfsdk:"label_selector" json:"labelSelector,omitempty"`
							Name               *string `tfsdk:"name" json:"name,omitempty"`
							Namespace          *string `tfsdk:"namespace" json:"namespace,omitempty"`
							Version            *string `tfsdk:"version" json:"version,omitempty"`
						} `tfsdk:"target" json:"target,omitempty"`
					} `tfsdk:"patches" json:"patches,omitempty"`
					Replicas *[]struct {
						Count *string `tfsdk:"count" json:"count,omitempty"`
						Name  *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"replicas" json:"replicas,omitempty"`
					Version *string `tfsdk:"version" json:"version,omitempty"`
				} `tfsdk:"kustomize" json:"kustomize,omitempty"`
				Path   *string `tfsdk:"path" json:"path,omitempty"`
				Plugin *struct {
					Env *[]struct {
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"env" json:"env,omitempty"`
					Name       *string `tfsdk:"name" json:"name,omitempty"`
					Parameters *[]struct {
						Array  *[]string          `tfsdk:"array" json:"array,omitempty"`
						Map    *map[string]string `tfsdk:"map" json:"map,omitempty"`
						Name   *string            `tfsdk:"name" json:"name,omitempty"`
						String *string            `tfsdk:"string" json:"string,omitempty"`
					} `tfsdk:"parameters" json:"parameters,omitempty"`
				} `tfsdk:"plugin" json:"plugin,omitempty"`
				Ref            *string `tfsdk:"ref" json:"ref,omitempty"`
				RepoURL        *string `tfsdk:"repo_url" json:"repoURL,omitempty"`
				TargetRevision *string `tfsdk:"target_revision" json:"targetRevision,omitempty"`
			} `tfsdk:"sources" json:"sources,omitempty"`
			SyncOptions  *[]string `tfsdk:"sync_options" json:"syncOptions,omitempty"`
			SyncStrategy *struct {
				Apply *struct {
					Force *bool `tfsdk:"force" json:"force,omitempty"`
				} `tfsdk:"apply" json:"apply,omitempty"`
				Hook *struct {
					Force *bool `tfsdk:"force" json:"force,omitempty"`
				} `tfsdk:"hook" json:"hook,omitempty"`
			} `tfsdk:"sync_strategy" json:"syncStrategy,omitempty"`
		} `tfsdk:"sync" json:"sync,omitempty"`
	} `tfsdk:"operation" json:"operation,omitempty"`
	Spec *struct {
		Destination *struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Server    *string `tfsdk:"server" json:"server,omitempty"`
		} `tfsdk:"destination" json:"destination,omitempty"`
		IgnoreDifferences *[]struct {
			Group                 *string   `tfsdk:"group" json:"group,omitempty"`
			JqPathExpressions     *[]string `tfsdk:"jq_path_expressions" json:"jqPathExpressions,omitempty"`
			JsonPointers          *[]string `tfsdk:"json_pointers" json:"jsonPointers,omitempty"`
			Kind                  *string   `tfsdk:"kind" json:"kind,omitempty"`
			ManagedFieldsManagers *[]string `tfsdk:"managed_fields_managers" json:"managedFieldsManagers,omitempty"`
			Name                  *string   `tfsdk:"name" json:"name,omitempty"`
			Namespace             *string   `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"ignore_differences" json:"ignoreDifferences,omitempty"`
		Info *[]struct {
			Name  *string `tfsdk:"name" json:"name,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"info" json:"info,omitempty"`
		Project              *string `tfsdk:"project" json:"project,omitempty"`
		RevisionHistoryLimit *int64  `tfsdk:"revision_history_limit" json:"revisionHistoryLimit,omitempty"`
		Source               *struct {
			Chart     *string `tfsdk:"chart" json:"chart,omitempty"`
			Directory *struct {
				Exclude *string `tfsdk:"exclude" json:"exclude,omitempty"`
				Include *string `tfsdk:"include" json:"include,omitempty"`
				Jsonnet *struct {
					ExtVars *[]struct {
						Code  *bool   `tfsdk:"code" json:"code,omitempty"`
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"ext_vars" json:"extVars,omitempty"`
					Libs *[]string `tfsdk:"libs" json:"libs,omitempty"`
					Tlas *[]struct {
						Code  *bool   `tfsdk:"code" json:"code,omitempty"`
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"tlas" json:"tlas,omitempty"`
				} `tfsdk:"jsonnet" json:"jsonnet,omitempty"`
				Recurse *bool `tfsdk:"recurse" json:"recurse,omitempty"`
			} `tfsdk:"directory" json:"directory,omitempty"`
			Helm *struct {
				FileParameters *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
					Path *string `tfsdk:"path" json:"path,omitempty"`
				} `tfsdk:"file_parameters" json:"fileParameters,omitempty"`
				IgnoreMissingValueFiles *bool `tfsdk:"ignore_missing_value_files" json:"ignoreMissingValueFiles,omitempty"`
				Parameters              *[]struct {
					ForceString *bool   `tfsdk:"force_string" json:"forceString,omitempty"`
					Name        *string `tfsdk:"name" json:"name,omitempty"`
					Value       *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"parameters" json:"parameters,omitempty"`
				PassCredentials *bool              `tfsdk:"pass_credentials" json:"passCredentials,omitempty"`
				ReleaseName     *string            `tfsdk:"release_name" json:"releaseName,omitempty"`
				SkipCrds        *bool              `tfsdk:"skip_crds" json:"skipCrds,omitempty"`
				ValueFiles      *[]string          `tfsdk:"value_files" json:"valueFiles,omitempty"`
				Values          *string            `tfsdk:"values" json:"values,omitempty"`
				ValuesObject    *map[string]string `tfsdk:"values_object" json:"valuesObject,omitempty"`
				Version         *string            `tfsdk:"version" json:"version,omitempty"`
			} `tfsdk:"helm" json:"helm,omitempty"`
			Kustomize *struct {
				CommonAnnotations         *map[string]string `tfsdk:"common_annotations" json:"commonAnnotations,omitempty"`
				CommonAnnotationsEnvsubst *bool              `tfsdk:"common_annotations_envsubst" json:"commonAnnotationsEnvsubst,omitempty"`
				CommonLabels              *map[string]string `tfsdk:"common_labels" json:"commonLabels,omitempty"`
				Components                *[]string          `tfsdk:"components" json:"components,omitempty"`
				ForceCommonAnnotations    *bool              `tfsdk:"force_common_annotations" json:"forceCommonAnnotations,omitempty"`
				ForceCommonLabels         *bool              `tfsdk:"force_common_labels" json:"forceCommonLabels,omitempty"`
				Images                    *[]string          `tfsdk:"images" json:"images,omitempty"`
				NamePrefix                *string            `tfsdk:"name_prefix" json:"namePrefix,omitempty"`
				NameSuffix                *string            `tfsdk:"name_suffix" json:"nameSuffix,omitempty"`
				Namespace                 *string            `tfsdk:"namespace" json:"namespace,omitempty"`
				Patches                   *[]struct {
					Options *map[string]string `tfsdk:"options" json:"options,omitempty"`
					Patch   *string            `tfsdk:"patch" json:"patch,omitempty"`
					Path    *string            `tfsdk:"path" json:"path,omitempty"`
					Target  *struct {
						AnnotationSelector *string `tfsdk:"annotation_selector" json:"annotationSelector,omitempty"`
						Group              *string `tfsdk:"group" json:"group,omitempty"`
						Kind               *string `tfsdk:"kind" json:"kind,omitempty"`
						LabelSelector      *string `tfsdk:"label_selector" json:"labelSelector,omitempty"`
						Name               *string `tfsdk:"name" json:"name,omitempty"`
						Namespace          *string `tfsdk:"namespace" json:"namespace,omitempty"`
						Version            *string `tfsdk:"version" json:"version,omitempty"`
					} `tfsdk:"target" json:"target,omitempty"`
				} `tfsdk:"patches" json:"patches,omitempty"`
				Replicas *[]struct {
					Count *string `tfsdk:"count" json:"count,omitempty"`
					Name  *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"replicas" json:"replicas,omitempty"`
				Version *string `tfsdk:"version" json:"version,omitempty"`
			} `tfsdk:"kustomize" json:"kustomize,omitempty"`
			Path   *string `tfsdk:"path" json:"path,omitempty"`
			Plugin *struct {
				Env *[]struct {
					Name  *string `tfsdk:"name" json:"name,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"env" json:"env,omitempty"`
				Name       *string `tfsdk:"name" json:"name,omitempty"`
				Parameters *[]struct {
					Array  *[]string          `tfsdk:"array" json:"array,omitempty"`
					Map    *map[string]string `tfsdk:"map" json:"map,omitempty"`
					Name   *string            `tfsdk:"name" json:"name,omitempty"`
					String *string            `tfsdk:"string" json:"string,omitempty"`
				} `tfsdk:"parameters" json:"parameters,omitempty"`
			} `tfsdk:"plugin" json:"plugin,omitempty"`
			Ref            *string `tfsdk:"ref" json:"ref,omitempty"`
			RepoURL        *string `tfsdk:"repo_url" json:"repoURL,omitempty"`
			TargetRevision *string `tfsdk:"target_revision" json:"targetRevision,omitempty"`
		} `tfsdk:"source" json:"source,omitempty"`
		Sources *[]struct {
			Chart     *string `tfsdk:"chart" json:"chart,omitempty"`
			Directory *struct {
				Exclude *string `tfsdk:"exclude" json:"exclude,omitempty"`
				Include *string `tfsdk:"include" json:"include,omitempty"`
				Jsonnet *struct {
					ExtVars *[]struct {
						Code  *bool   `tfsdk:"code" json:"code,omitempty"`
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"ext_vars" json:"extVars,omitempty"`
					Libs *[]string `tfsdk:"libs" json:"libs,omitempty"`
					Tlas *[]struct {
						Code  *bool   `tfsdk:"code" json:"code,omitempty"`
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"tlas" json:"tlas,omitempty"`
				} `tfsdk:"jsonnet" json:"jsonnet,omitempty"`
				Recurse *bool `tfsdk:"recurse" json:"recurse,omitempty"`
			} `tfsdk:"directory" json:"directory,omitempty"`
			Helm *struct {
				FileParameters *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
					Path *string `tfsdk:"path" json:"path,omitempty"`
				} `tfsdk:"file_parameters" json:"fileParameters,omitempty"`
				IgnoreMissingValueFiles *bool `tfsdk:"ignore_missing_value_files" json:"ignoreMissingValueFiles,omitempty"`
				Parameters              *[]struct {
					ForceString *bool   `tfsdk:"force_string" json:"forceString,omitempty"`
					Name        *string `tfsdk:"name" json:"name,omitempty"`
					Value       *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"parameters" json:"parameters,omitempty"`
				PassCredentials *bool              `tfsdk:"pass_credentials" json:"passCredentials,omitempty"`
				ReleaseName     *string            `tfsdk:"release_name" json:"releaseName,omitempty"`
				SkipCrds        *bool              `tfsdk:"skip_crds" json:"skipCrds,omitempty"`
				ValueFiles      *[]string          `tfsdk:"value_files" json:"valueFiles,omitempty"`
				Values          *string            `tfsdk:"values" json:"values,omitempty"`
				ValuesObject    *map[string]string `tfsdk:"values_object" json:"valuesObject,omitempty"`
				Version         *string            `tfsdk:"version" json:"version,omitempty"`
			} `tfsdk:"helm" json:"helm,omitempty"`
			Kustomize *struct {
				CommonAnnotations         *map[string]string `tfsdk:"common_annotations" json:"commonAnnotations,omitempty"`
				CommonAnnotationsEnvsubst *bool              `tfsdk:"common_annotations_envsubst" json:"commonAnnotationsEnvsubst,omitempty"`
				CommonLabels              *map[string]string `tfsdk:"common_labels" json:"commonLabels,omitempty"`
				Components                *[]string          `tfsdk:"components" json:"components,omitempty"`
				ForceCommonAnnotations    *bool              `tfsdk:"force_common_annotations" json:"forceCommonAnnotations,omitempty"`
				ForceCommonLabels         *bool              `tfsdk:"force_common_labels" json:"forceCommonLabels,omitempty"`
				Images                    *[]string          `tfsdk:"images" json:"images,omitempty"`
				NamePrefix                *string            `tfsdk:"name_prefix" json:"namePrefix,omitempty"`
				NameSuffix                *string            `tfsdk:"name_suffix" json:"nameSuffix,omitempty"`
				Namespace                 *string            `tfsdk:"namespace" json:"namespace,omitempty"`
				Patches                   *[]struct {
					Options *map[string]string `tfsdk:"options" json:"options,omitempty"`
					Patch   *string            `tfsdk:"patch" json:"patch,omitempty"`
					Path    *string            `tfsdk:"path" json:"path,omitempty"`
					Target  *struct {
						AnnotationSelector *string `tfsdk:"annotation_selector" json:"annotationSelector,omitempty"`
						Group              *string `tfsdk:"group" json:"group,omitempty"`
						Kind               *string `tfsdk:"kind" json:"kind,omitempty"`
						LabelSelector      *string `tfsdk:"label_selector" json:"labelSelector,omitempty"`
						Name               *string `tfsdk:"name" json:"name,omitempty"`
						Namespace          *string `tfsdk:"namespace" json:"namespace,omitempty"`
						Version            *string `tfsdk:"version" json:"version,omitempty"`
					} `tfsdk:"target" json:"target,omitempty"`
				} `tfsdk:"patches" json:"patches,omitempty"`
				Replicas *[]struct {
					Count *string `tfsdk:"count" json:"count,omitempty"`
					Name  *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"replicas" json:"replicas,omitempty"`
				Version *string `tfsdk:"version" json:"version,omitempty"`
			} `tfsdk:"kustomize" json:"kustomize,omitempty"`
			Path   *string `tfsdk:"path" json:"path,omitempty"`
			Plugin *struct {
				Env *[]struct {
					Name  *string `tfsdk:"name" json:"name,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"env" json:"env,omitempty"`
				Name       *string `tfsdk:"name" json:"name,omitempty"`
				Parameters *[]struct {
					Array  *[]string          `tfsdk:"array" json:"array,omitempty"`
					Map    *map[string]string `tfsdk:"map" json:"map,omitempty"`
					Name   *string            `tfsdk:"name" json:"name,omitempty"`
					String *string            `tfsdk:"string" json:"string,omitempty"`
				} `tfsdk:"parameters" json:"parameters,omitempty"`
			} `tfsdk:"plugin" json:"plugin,omitempty"`
			Ref            *string `tfsdk:"ref" json:"ref,omitempty"`
			RepoURL        *string `tfsdk:"repo_url" json:"repoURL,omitempty"`
			TargetRevision *string `tfsdk:"target_revision" json:"targetRevision,omitempty"`
		} `tfsdk:"sources" json:"sources,omitempty"`
		SyncPolicy *struct {
			Automated *struct {
				AllowEmpty *bool `tfsdk:"allow_empty" json:"allowEmpty,omitempty"`
				Prune      *bool `tfsdk:"prune" json:"prune,omitempty"`
				SelfHeal   *bool `tfsdk:"self_heal" json:"selfHeal,omitempty"`
			} `tfsdk:"automated" json:"automated,omitempty"`
			ManagedNamespaceMetadata *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"managed_namespace_metadata" json:"managedNamespaceMetadata,omitempty"`
			Retry *struct {
				Backoff *struct {
					Duration    *string `tfsdk:"duration" json:"duration,omitempty"`
					Factor      *int64  `tfsdk:"factor" json:"factor,omitempty"`
					MaxDuration *string `tfsdk:"max_duration" json:"maxDuration,omitempty"`
				} `tfsdk:"backoff" json:"backoff,omitempty"`
				Limit *int64 `tfsdk:"limit" json:"limit,omitempty"`
			} `tfsdk:"retry" json:"retry,omitempty"`
			SyncOptions *[]string `tfsdk:"sync_options" json:"syncOptions,omitempty"`
		} `tfsdk:"sync_policy" json:"syncPolicy,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ArgoprojIoApplicationV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_argoproj_io_application_v1alpha1_manifest"
}

func (r *ArgoprojIoApplicationV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Application is a definition of Application resource.",
		MarkdownDescription: "Application is a definition of Application resource.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

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

			"operation": schema.SingleNestedAttribute{
				Description:         "Operation contains information about a requested or running operation",
				MarkdownDescription: "Operation contains information about a requested or running operation",
				Attributes: map[string]schema.Attribute{
					"info": schema.ListNestedAttribute{
						Description:         "Info is a list of informational items for this operation",
						MarkdownDescription: "Info is a list of informational items for this operation",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"value": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"initiated_by": schema.SingleNestedAttribute{
						Description:         "InitiatedBy contains information about who initiated the operations",
						MarkdownDescription: "InitiatedBy contains information about who initiated the operations",
						Attributes: map[string]schema.Attribute{
							"automated": schema.BoolAttribute{
								Description:         "Automated is set to true if operation was initiated automatically by the application controller.",
								MarkdownDescription: "Automated is set to true if operation was initiated automatically by the application controller.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"username": schema.StringAttribute{
								Description:         "Username contains the name of a user who started operation",
								MarkdownDescription: "Username contains the name of a user who started operation",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"retry": schema.SingleNestedAttribute{
						Description:         "Retry controls the strategy to apply if a sync fails",
						MarkdownDescription: "Retry controls the strategy to apply if a sync fails",
						Attributes: map[string]schema.Attribute{
							"backoff": schema.SingleNestedAttribute{
								Description:         "Backoff controls how to backoff on subsequent retries of failed syncs",
								MarkdownDescription: "Backoff controls how to backoff on subsequent retries of failed syncs",
								Attributes: map[string]schema.Attribute{
									"duration": schema.StringAttribute{
										Description:         "Duration is the amount to back off. Default unit is seconds, but could also be a duration (e.g. '2m', '1h')",
										MarkdownDescription: "Duration is the amount to back off. Default unit is seconds, but could also be a duration (e.g. '2m', '1h')",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"factor": schema.Int64Attribute{
										Description:         "Factor is a factor to multiply the base duration after each failed retry",
										MarkdownDescription: "Factor is a factor to multiply the base duration after each failed retry",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_duration": schema.StringAttribute{
										Description:         "MaxDuration is the maximum amount of time allowed for the backoff strategy",
										MarkdownDescription: "MaxDuration is the maximum amount of time allowed for the backoff strategy",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"limit": schema.Int64Attribute{
								Description:         "Limit is the maximum number of attempts for retrying a failed sync. If set to 0, no retries will be performed.",
								MarkdownDescription: "Limit is the maximum number of attempts for retrying a failed sync. If set to 0, no retries will be performed.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"sync": schema.SingleNestedAttribute{
						Description:         "Sync contains parameters for the operation",
						MarkdownDescription: "Sync contains parameters for the operation",
						Attributes: map[string]schema.Attribute{
							"dry_run": schema.BoolAttribute{
								Description:         "DryRun specifies to perform a 'kubectl apply --dry-run' without actually performing the sync",
								MarkdownDescription: "DryRun specifies to perform a 'kubectl apply --dry-run' without actually performing the sync",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"manifests": schema.ListAttribute{
								Description:         "Manifests is an optional field that overrides sync source with a local directory for development",
								MarkdownDescription: "Manifests is an optional field that overrides sync source with a local directory for development",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"prune": schema.BoolAttribute{
								Description:         "Prune specifies to delete resources from the cluster that are no longer tracked in git",
								MarkdownDescription: "Prune specifies to delete resources from the cluster that are no longer tracked in git",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resources": schema.ListNestedAttribute{
								Description:         "Resources describes which resources shall be part of the sync",
								MarkdownDescription: "Resources describes which resources shall be part of the sync",
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
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"revision": schema.StringAttribute{
								Description:         "Revision is the revision (Git) or chart version (Helm) which to sync the application to If omitted, will use the revision specified in app spec.",
								MarkdownDescription: "Revision is the revision (Git) or chart version (Helm) which to sync the application to If omitted, will use the revision specified in app spec.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"revisions": schema.ListAttribute{
								Description:         "Revisions is the list of revision (Git) or chart version (Helm) which to sync each source in sources field for the application to If omitted, will use the revision specified in app spec.",
								MarkdownDescription: "Revisions is the list of revision (Git) or chart version (Helm) which to sync each source in sources field for the application to If omitted, will use the revision specified in app spec.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"source": schema.SingleNestedAttribute{
								Description:         "Source overrides the source definition set in the application. This is typically set in a Rollback operation and is nil during a Sync operation",
								MarkdownDescription: "Source overrides the source definition set in the application. This is typically set in a Rollback operation and is nil during a Sync operation",
								Attributes: map[string]schema.Attribute{
									"chart": schema.StringAttribute{
										Description:         "Chart is a Helm chart name, and must be specified for applications sourced from a Helm repo.",
										MarkdownDescription: "Chart is a Helm chart name, and must be specified for applications sourced from a Helm repo.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"directory": schema.SingleNestedAttribute{
										Description:         "Directory holds path/directory specific options",
										MarkdownDescription: "Directory holds path/directory specific options",
										Attributes: map[string]schema.Attribute{
											"exclude": schema.StringAttribute{
												Description:         "Exclude contains a glob pattern to match paths against that should be explicitly excluded from being used during manifest generation",
												MarkdownDescription: "Exclude contains a glob pattern to match paths against that should be explicitly excluded from being used during manifest generation",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"include": schema.StringAttribute{
												Description:         "Include contains a glob pattern to match paths against that should be explicitly included during manifest generation",
												MarkdownDescription: "Include contains a glob pattern to match paths against that should be explicitly included during manifest generation",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"jsonnet": schema.SingleNestedAttribute{
												Description:         "Jsonnet holds options specific to Jsonnet",
												MarkdownDescription: "Jsonnet holds options specific to Jsonnet",
												Attributes: map[string]schema.Attribute{
													"ext_vars": schema.ListNestedAttribute{
														Description:         "ExtVars is a list of Jsonnet External Variables",
														MarkdownDescription: "ExtVars is a list of Jsonnet External Variables",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"code": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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

																"value": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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

													"libs": schema.ListAttribute{
														Description:         "Additional library search dirs",
														MarkdownDescription: "Additional library search dirs",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"tlas": schema.ListNestedAttribute{
														Description:         "TLAS is a list of Jsonnet Top-level Arguments",
														MarkdownDescription: "TLAS is a list of Jsonnet Top-level Arguments",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"code": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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

																"value": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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

											"recurse": schema.BoolAttribute{
												Description:         "Recurse specifies whether to scan a directory recursively for manifests",
												MarkdownDescription: "Recurse specifies whether to scan a directory recursively for manifests",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"helm": schema.SingleNestedAttribute{
										Description:         "Helm holds helm specific options",
										MarkdownDescription: "Helm holds helm specific options",
										Attributes: map[string]schema.Attribute{
											"file_parameters": schema.ListNestedAttribute{
												Description:         "FileParameters are file parameters to the helm template",
												MarkdownDescription: "FileParameters are file parameters to the helm template",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name is the name of the Helm parameter",
															MarkdownDescription: "Name is the name of the Helm parameter",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"path": schema.StringAttribute{
															Description:         "Path is the path to the file containing the values for the Helm parameter",
															MarkdownDescription: "Path is the path to the file containing the values for the Helm parameter",
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

											"ignore_missing_value_files": schema.BoolAttribute{
												Description:         "IgnoreMissingValueFiles prevents helm template from failing when valueFiles do not exist locally by not appending them to helm template --values",
												MarkdownDescription: "IgnoreMissingValueFiles prevents helm template from failing when valueFiles do not exist locally by not appending them to helm template --values",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"parameters": schema.ListNestedAttribute{
												Description:         "Parameters is a list of Helm parameters which are passed to the helm template command upon manifest generation",
												MarkdownDescription: "Parameters is a list of Helm parameters which are passed to the helm template command upon manifest generation",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"force_string": schema.BoolAttribute{
															Description:         "ForceString determines whether to tell Helm to interpret booleans and numbers as strings",
															MarkdownDescription: "ForceString determines whether to tell Helm to interpret booleans and numbers as strings",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name is the name of the Helm parameter",
															MarkdownDescription: "Name is the name of the Helm parameter",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"value": schema.StringAttribute{
															Description:         "Value is the value for the Helm parameter",
															MarkdownDescription: "Value is the value for the Helm parameter",
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

											"pass_credentials": schema.BoolAttribute{
												Description:         "PassCredentials pass credentials to all domains (Helm's --pass-credentials)",
												MarkdownDescription: "PassCredentials pass credentials to all domains (Helm's --pass-credentials)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"release_name": schema.StringAttribute{
												Description:         "ReleaseName is the Helm release name to use. If omitted it will use the application name",
												MarkdownDescription: "ReleaseName is the Helm release name to use. If omitted it will use the application name",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"skip_crds": schema.BoolAttribute{
												Description:         "SkipCrds skips custom resource definition installation step (Helm's --skip-crds)",
												MarkdownDescription: "SkipCrds skips custom resource definition installation step (Helm's --skip-crds)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"value_files": schema.ListAttribute{
												Description:         "ValuesFiles is a list of Helm value files to use when generating a template",
												MarkdownDescription: "ValuesFiles is a list of Helm value files to use when generating a template",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"values": schema.StringAttribute{
												Description:         "Values specifies Helm values to be passed to helm template, typically defined as a block. ValuesObject takes precedence over Values, so use one or the other.",
												MarkdownDescription: "Values specifies Helm values to be passed to helm template, typically defined as a block. ValuesObject takes precedence over Values, so use one or the other.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"values_object": schema.MapAttribute{
												Description:         "ValuesObject specifies Helm values to be passed to helm template, defined as a map. This takes precedence over Values.",
												MarkdownDescription: "ValuesObject specifies Helm values to be passed to helm template, defined as a map. This takes precedence over Values.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"version": schema.StringAttribute{
												Description:         "Version is the Helm version to use for templating ('3')",
												MarkdownDescription: "Version is the Helm version to use for templating ('3')",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"kustomize": schema.SingleNestedAttribute{
										Description:         "Kustomize holds kustomize specific options",
										MarkdownDescription: "Kustomize holds kustomize specific options",
										Attributes: map[string]schema.Attribute{
											"common_annotations": schema.MapAttribute{
												Description:         "CommonAnnotations is a list of additional annotations to add to rendered manifests",
												MarkdownDescription: "CommonAnnotations is a list of additional annotations to add to rendered manifests",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"common_annotations_envsubst": schema.BoolAttribute{
												Description:         "CommonAnnotationsEnvsubst specifies whether to apply env variables substitution for annotation values",
												MarkdownDescription: "CommonAnnotationsEnvsubst specifies whether to apply env variables substitution for annotation values",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"common_labels": schema.MapAttribute{
												Description:         "CommonLabels is a list of additional labels to add to rendered manifests",
												MarkdownDescription: "CommonLabels is a list of additional labels to add to rendered manifests",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"components": schema.ListAttribute{
												Description:         "Components specifies a list of kustomize components to add to the kustomization before building",
												MarkdownDescription: "Components specifies a list of kustomize components to add to the kustomization before building",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"force_common_annotations": schema.BoolAttribute{
												Description:         "ForceCommonAnnotations specifies whether to force applying common annotations to resources for Kustomize apps",
												MarkdownDescription: "ForceCommonAnnotations specifies whether to force applying common annotations to resources for Kustomize apps",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"force_common_labels": schema.BoolAttribute{
												Description:         "ForceCommonLabels specifies whether to force applying common labels to resources for Kustomize apps",
												MarkdownDescription: "ForceCommonLabels specifies whether to force applying common labels to resources for Kustomize apps",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"images": schema.ListAttribute{
												Description:         "Images is a list of Kustomize image override specifications",
												MarkdownDescription: "Images is a list of Kustomize image override specifications",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name_prefix": schema.StringAttribute{
												Description:         "NamePrefix is a prefix appended to resources for Kustomize apps",
												MarkdownDescription: "NamePrefix is a prefix appended to resources for Kustomize apps",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name_suffix": schema.StringAttribute{
												Description:         "NameSuffix is a suffix appended to resources for Kustomize apps",
												MarkdownDescription: "NameSuffix is a suffix appended to resources for Kustomize apps",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "Namespace sets the namespace that Kustomize adds to all resources",
												MarkdownDescription: "Namespace sets the namespace that Kustomize adds to all resources",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"patches": schema.ListNestedAttribute{
												Description:         "Patches is a list of Kustomize patches",
												MarkdownDescription: "Patches is a list of Kustomize patches",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"options": schema.MapAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"patch": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"path": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"target": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"annotation_selector": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

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

																"label_selector": schema.StringAttribute{
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

											"replicas": schema.ListNestedAttribute{
												Description:         "Replicas is a list of Kustomize Replicas override specifications",
												MarkdownDescription: "Replicas is a list of Kustomize Replicas override specifications",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"count": schema.StringAttribute{
															Description:         "Number of replicas",
															MarkdownDescription: "Number of replicas",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name of Deployment or StatefulSet",
															MarkdownDescription: "Name of Deployment or StatefulSet",
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

											"version": schema.StringAttribute{
												Description:         "Version controls which version of Kustomize to use for rendering manifests",
												MarkdownDescription: "Version controls which version of Kustomize to use for rendering manifests",
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
										Description:         "Path is a directory path within the Git repository, and is only valid for applications sourced from Git.",
										MarkdownDescription: "Path is a directory path within the Git repository, and is only valid for applications sourced from Git.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"plugin": schema.SingleNestedAttribute{
										Description:         "Plugin holds config management plugin specific options",
										MarkdownDescription: "Plugin holds config management plugin specific options",
										Attributes: map[string]schema.Attribute{
											"env": schema.ListNestedAttribute{
												Description:         "Env is a list of environment variable entries",
												MarkdownDescription: "Env is a list of environment variable entries",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name is the name of the variable, usually expressed in uppercase",
															MarkdownDescription: "Name is the name of the variable, usually expressed in uppercase",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"value": schema.StringAttribute{
															Description:         "Value is the value of the variable",
															MarkdownDescription: "Value is the value of the variable",
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
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"parameters": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"array": schema.ListAttribute{
															Description:         "Array is the value of an array type parameter.",
															MarkdownDescription: "Array is the value of an array type parameter.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"map": schema.MapAttribute{
															Description:         "Map is the value of a map type parameter.",
															MarkdownDescription: "Map is the value of a map type parameter.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name is the name identifying a parameter.",
															MarkdownDescription: "Name is the name identifying a parameter.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"string": schema.StringAttribute{
															Description:         "String_ is the value of a string type parameter.",
															MarkdownDescription: "String_ is the value of a string type parameter.",
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

									"ref": schema.StringAttribute{
										Description:         "Ref is reference to another source within sources field. This field will not be used if used with a 'source' tag.",
										MarkdownDescription: "Ref is reference to another source within sources field. This field will not be used if used with a 'source' tag.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"repo_url": schema.StringAttribute{
										Description:         "RepoURL is the URL to the repository (Git or Helm) that contains the application manifests",
										MarkdownDescription: "RepoURL is the URL to the repository (Git or Helm) that contains the application manifests",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"target_revision": schema.StringAttribute{
										Description:         "TargetRevision defines the revision of the source to sync the application to. In case of Git, this can be commit, tag, or branch. If omitted, will equal to HEAD. In case of Helm, this is a semver tag for the Chart's version.",
										MarkdownDescription: "TargetRevision defines the revision of the source to sync the application to. In case of Git, this can be commit, tag, or branch. If omitted, will equal to HEAD. In case of Helm, this is a semver tag for the Chart's version.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"sources": schema.ListNestedAttribute{
								Description:         "Sources overrides the source definition set in the application. This is typically set in a Rollback operation and is nil during a Sync operation",
								MarkdownDescription: "Sources overrides the source definition set in the application. This is typically set in a Rollback operation and is nil during a Sync operation",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"chart": schema.StringAttribute{
											Description:         "Chart is a Helm chart name, and must be specified for applications sourced from a Helm repo.",
											MarkdownDescription: "Chart is a Helm chart name, and must be specified for applications sourced from a Helm repo.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"directory": schema.SingleNestedAttribute{
											Description:         "Directory holds path/directory specific options",
											MarkdownDescription: "Directory holds path/directory specific options",
											Attributes: map[string]schema.Attribute{
												"exclude": schema.StringAttribute{
													Description:         "Exclude contains a glob pattern to match paths against that should be explicitly excluded from being used during manifest generation",
													MarkdownDescription: "Exclude contains a glob pattern to match paths against that should be explicitly excluded from being used during manifest generation",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"include": schema.StringAttribute{
													Description:         "Include contains a glob pattern to match paths against that should be explicitly included during manifest generation",
													MarkdownDescription: "Include contains a glob pattern to match paths against that should be explicitly included during manifest generation",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"jsonnet": schema.SingleNestedAttribute{
													Description:         "Jsonnet holds options specific to Jsonnet",
													MarkdownDescription: "Jsonnet holds options specific to Jsonnet",
													Attributes: map[string]schema.Attribute{
														"ext_vars": schema.ListNestedAttribute{
															Description:         "ExtVars is a list of Jsonnet External Variables",
															MarkdownDescription: "ExtVars is a list of Jsonnet External Variables",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"code": schema.BoolAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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

																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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

														"libs": schema.ListAttribute{
															Description:         "Additional library search dirs",
															MarkdownDescription: "Additional library search dirs",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"tlas": schema.ListNestedAttribute{
															Description:         "TLAS is a list of Jsonnet Top-level Arguments",
															MarkdownDescription: "TLAS is a list of Jsonnet Top-level Arguments",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"code": schema.BoolAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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

																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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

												"recurse": schema.BoolAttribute{
													Description:         "Recurse specifies whether to scan a directory recursively for manifests",
													MarkdownDescription: "Recurse specifies whether to scan a directory recursively for manifests",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"helm": schema.SingleNestedAttribute{
											Description:         "Helm holds helm specific options",
											MarkdownDescription: "Helm holds helm specific options",
											Attributes: map[string]schema.Attribute{
												"file_parameters": schema.ListNestedAttribute{
													Description:         "FileParameters are file parameters to the helm template",
													MarkdownDescription: "FileParameters are file parameters to the helm template",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name is the name of the Helm parameter",
																MarkdownDescription: "Name is the name of the Helm parameter",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"path": schema.StringAttribute{
																Description:         "Path is the path to the file containing the values for the Helm parameter",
																MarkdownDescription: "Path is the path to the file containing the values for the Helm parameter",
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

												"ignore_missing_value_files": schema.BoolAttribute{
													Description:         "IgnoreMissingValueFiles prevents helm template from failing when valueFiles do not exist locally by not appending them to helm template --values",
													MarkdownDescription: "IgnoreMissingValueFiles prevents helm template from failing when valueFiles do not exist locally by not appending them to helm template --values",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"parameters": schema.ListNestedAttribute{
													Description:         "Parameters is a list of Helm parameters which are passed to the helm template command upon manifest generation",
													MarkdownDescription: "Parameters is a list of Helm parameters which are passed to the helm template command upon manifest generation",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"force_string": schema.BoolAttribute{
																Description:         "ForceString determines whether to tell Helm to interpret booleans and numbers as strings",
																MarkdownDescription: "ForceString determines whether to tell Helm to interpret booleans and numbers as strings",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "Name is the name of the Helm parameter",
																MarkdownDescription: "Name is the name of the Helm parameter",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"value": schema.StringAttribute{
																Description:         "Value is the value for the Helm parameter",
																MarkdownDescription: "Value is the value for the Helm parameter",
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

												"pass_credentials": schema.BoolAttribute{
													Description:         "PassCredentials pass credentials to all domains (Helm's --pass-credentials)",
													MarkdownDescription: "PassCredentials pass credentials to all domains (Helm's --pass-credentials)",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"release_name": schema.StringAttribute{
													Description:         "ReleaseName is the Helm release name to use. If omitted it will use the application name",
													MarkdownDescription: "ReleaseName is the Helm release name to use. If omitted it will use the application name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"skip_crds": schema.BoolAttribute{
													Description:         "SkipCrds skips custom resource definition installation step (Helm's --skip-crds)",
													MarkdownDescription: "SkipCrds skips custom resource definition installation step (Helm's --skip-crds)",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value_files": schema.ListAttribute{
													Description:         "ValuesFiles is a list of Helm value files to use when generating a template",
													MarkdownDescription: "ValuesFiles is a list of Helm value files to use when generating a template",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"values": schema.StringAttribute{
													Description:         "Values specifies Helm values to be passed to helm template, typically defined as a block. ValuesObject takes precedence over Values, so use one or the other.",
													MarkdownDescription: "Values specifies Helm values to be passed to helm template, typically defined as a block. ValuesObject takes precedence over Values, so use one or the other.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"values_object": schema.MapAttribute{
													Description:         "ValuesObject specifies Helm values to be passed to helm template, defined as a map. This takes precedence over Values.",
													MarkdownDescription: "ValuesObject specifies Helm values to be passed to helm template, defined as a map. This takes precedence over Values.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"version": schema.StringAttribute{
													Description:         "Version is the Helm version to use for templating ('3')",
													MarkdownDescription: "Version is the Helm version to use for templating ('3')",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"kustomize": schema.SingleNestedAttribute{
											Description:         "Kustomize holds kustomize specific options",
											MarkdownDescription: "Kustomize holds kustomize specific options",
											Attributes: map[string]schema.Attribute{
												"common_annotations": schema.MapAttribute{
													Description:         "CommonAnnotations is a list of additional annotations to add to rendered manifests",
													MarkdownDescription: "CommonAnnotations is a list of additional annotations to add to rendered manifests",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"common_annotations_envsubst": schema.BoolAttribute{
													Description:         "CommonAnnotationsEnvsubst specifies whether to apply env variables substitution for annotation values",
													MarkdownDescription: "CommonAnnotationsEnvsubst specifies whether to apply env variables substitution for annotation values",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"common_labels": schema.MapAttribute{
													Description:         "CommonLabels is a list of additional labels to add to rendered manifests",
													MarkdownDescription: "CommonLabels is a list of additional labels to add to rendered manifests",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"components": schema.ListAttribute{
													Description:         "Components specifies a list of kustomize components to add to the kustomization before building",
													MarkdownDescription: "Components specifies a list of kustomize components to add to the kustomization before building",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"force_common_annotations": schema.BoolAttribute{
													Description:         "ForceCommonAnnotations specifies whether to force applying common annotations to resources for Kustomize apps",
													MarkdownDescription: "ForceCommonAnnotations specifies whether to force applying common annotations to resources for Kustomize apps",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"force_common_labels": schema.BoolAttribute{
													Description:         "ForceCommonLabels specifies whether to force applying common labels to resources for Kustomize apps",
													MarkdownDescription: "ForceCommonLabels specifies whether to force applying common labels to resources for Kustomize apps",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"images": schema.ListAttribute{
													Description:         "Images is a list of Kustomize image override specifications",
													MarkdownDescription: "Images is a list of Kustomize image override specifications",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name_prefix": schema.StringAttribute{
													Description:         "NamePrefix is a prefix appended to resources for Kustomize apps",
													MarkdownDescription: "NamePrefix is a prefix appended to resources for Kustomize apps",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name_suffix": schema.StringAttribute{
													Description:         "NameSuffix is a suffix appended to resources for Kustomize apps",
													MarkdownDescription: "NameSuffix is a suffix appended to resources for Kustomize apps",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespace": schema.StringAttribute{
													Description:         "Namespace sets the namespace that Kustomize adds to all resources",
													MarkdownDescription: "Namespace sets the namespace that Kustomize adds to all resources",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"patches": schema.ListNestedAttribute{
													Description:         "Patches is a list of Kustomize patches",
													MarkdownDescription: "Patches is a list of Kustomize patches",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"options": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"patch": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"path": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"target": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"annotation_selector": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

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

																	"label_selector": schema.StringAttribute{
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

												"replicas": schema.ListNestedAttribute{
													Description:         "Replicas is a list of Kustomize Replicas override specifications",
													MarkdownDescription: "Replicas is a list of Kustomize Replicas override specifications",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"count": schema.StringAttribute{
																Description:         "Number of replicas",
																MarkdownDescription: "Number of replicas",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "Name of Deployment or StatefulSet",
																MarkdownDescription: "Name of Deployment or StatefulSet",
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

												"version": schema.StringAttribute{
													Description:         "Version controls which version of Kustomize to use for rendering manifests",
													MarkdownDescription: "Version controls which version of Kustomize to use for rendering manifests",
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
											Description:         "Path is a directory path within the Git repository, and is only valid for applications sourced from Git.",
											MarkdownDescription: "Path is a directory path within the Git repository, and is only valid for applications sourced from Git.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"plugin": schema.SingleNestedAttribute{
											Description:         "Plugin holds config management plugin specific options",
											MarkdownDescription: "Plugin holds config management plugin specific options",
											Attributes: map[string]schema.Attribute{
												"env": schema.ListNestedAttribute{
													Description:         "Env is a list of environment variable entries",
													MarkdownDescription: "Env is a list of environment variable entries",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name is the name of the variable, usually expressed in uppercase",
																MarkdownDescription: "Name is the name of the variable, usually expressed in uppercase",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"value": schema.StringAttribute{
																Description:         "Value is the value of the variable",
																MarkdownDescription: "Value is the value of the variable",
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
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"parameters": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"array": schema.ListAttribute{
																Description:         "Array is the value of an array type parameter.",
																MarkdownDescription: "Array is the value of an array type parameter.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"map": schema.MapAttribute{
																Description:         "Map is the value of a map type parameter.",
																MarkdownDescription: "Map is the value of a map type parameter.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "Name is the name identifying a parameter.",
																MarkdownDescription: "Name is the name identifying a parameter.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"string": schema.StringAttribute{
																Description:         "String_ is the value of a string type parameter.",
																MarkdownDescription: "String_ is the value of a string type parameter.",
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

										"ref": schema.StringAttribute{
											Description:         "Ref is reference to another source within sources field. This field will not be used if used with a 'source' tag.",
											MarkdownDescription: "Ref is reference to another source within sources field. This field will not be used if used with a 'source' tag.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"repo_url": schema.StringAttribute{
											Description:         "RepoURL is the URL to the repository (Git or Helm) that contains the application manifests",
											MarkdownDescription: "RepoURL is the URL to the repository (Git or Helm) that contains the application manifests",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"target_revision": schema.StringAttribute{
											Description:         "TargetRevision defines the revision of the source to sync the application to. In case of Git, this can be commit, tag, or branch. If omitted, will equal to HEAD. In case of Helm, this is a semver tag for the Chart's version.",
											MarkdownDescription: "TargetRevision defines the revision of the source to sync the application to. In case of Git, this can be commit, tag, or branch. If omitted, will equal to HEAD. In case of Helm, this is a semver tag for the Chart's version.",
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

							"sync_options": schema.ListAttribute{
								Description:         "SyncOptions provide per-sync sync-options, e.g. Validate=false",
								MarkdownDescription: "SyncOptions provide per-sync sync-options, e.g. Validate=false",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"sync_strategy": schema.SingleNestedAttribute{
								Description:         "SyncStrategy describes how to perform the sync",
								MarkdownDescription: "SyncStrategy describes how to perform the sync",
								Attributes: map[string]schema.Attribute{
									"apply": schema.SingleNestedAttribute{
										Description:         "Apply will perform a 'kubectl apply' to perform the sync.",
										MarkdownDescription: "Apply will perform a 'kubectl apply' to perform the sync.",
										Attributes: map[string]schema.Attribute{
											"force": schema.BoolAttribute{
												Description:         "Force indicates whether or not to supply the --force flag to 'kubectl apply'. The --force flag deletes and re-create the resource, when PATCH encounters conflict and has retried for 5 times.",
												MarkdownDescription: "Force indicates whether or not to supply the --force flag to 'kubectl apply'. The --force flag deletes and re-create the resource, when PATCH encounters conflict and has retried for 5 times.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"hook": schema.SingleNestedAttribute{
										Description:         "Hook will submit any referenced resources to perform the sync. This is the default strategy",
										MarkdownDescription: "Hook will submit any referenced resources to perform the sync. This is the default strategy",
										Attributes: map[string]schema.Attribute{
											"force": schema.BoolAttribute{
												Description:         "Force indicates whether or not to supply the --force flag to 'kubectl apply'. The --force flag deletes and re-create the resource, when PATCH encounters conflict and has retried for 5 times.",
												MarkdownDescription: "Force indicates whether or not to supply the --force flag to 'kubectl apply'. The --force flag deletes and re-create the resource, when PATCH encounters conflict and has retried for 5 times.",
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
				Required: false,
				Optional: true,
				Computed: false,
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "ApplicationSpec represents desired application state. Contains link to repository with application definition and additional parameters link definition revision.",
				MarkdownDescription: "ApplicationSpec represents desired application state. Contains link to repository with application definition and additional parameters link definition revision.",
				Attributes: map[string]schema.Attribute{
					"destination": schema.SingleNestedAttribute{
						Description:         "Destination is a reference to the target Kubernetes server and namespace",
						MarkdownDescription: "Destination is a reference to the target Kubernetes server and namespace",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name is an alternate way of specifying the target cluster by its symbolic name. This must be set if Server is not set.",
								MarkdownDescription: "Name is an alternate way of specifying the target cluster by its symbolic name. This must be set if Server is not set.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace specifies the target namespace for the application's resources. The namespace will only be set for namespace-scoped resources that have not set a value for .metadata.namespace",
								MarkdownDescription: "Namespace specifies the target namespace for the application's resources. The namespace will only be set for namespace-scoped resources that have not set a value for .metadata.namespace",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"server": schema.StringAttribute{
								Description:         "Server specifies the URL of the target cluster's Kubernetes control plane API. This must be set if Name is not set.",
								MarkdownDescription: "Server specifies the URL of the target cluster's Kubernetes control plane API. This must be set if Name is not set.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"ignore_differences": schema.ListNestedAttribute{
						Description:         "IgnoreDifferences is a list of resources and their fields which should be ignored during comparison",
						MarkdownDescription: "IgnoreDifferences is a list of resources and their fields which should be ignored during comparison",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"group": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"jq_path_expressions": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"json_pointers": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
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

								"managed_fields_managers": schema.ListAttribute{
									Description:         "ManagedFieldsManagers is a list of trusted managers. Fields mutated by those managers will take precedence over the desired state defined in the SCM and won't be displayed in diffs",
									MarkdownDescription: "ManagedFieldsManagers is a list of trusted managers. Fields mutated by those managers will take precedence over the desired state defined in the SCM and won't be displayed in diffs",
									ElementType:         types.StringType,
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
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"info": schema.ListNestedAttribute{
						Description:         "Info contains a list of information (URLs, email addresses, and plain text) that relates to the application",
						MarkdownDescription: "Info contains a list of information (URLs, email addresses, and plain text) that relates to the application",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"value": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"project": schema.StringAttribute{
						Description:         "Project is a reference to the project this application belongs to. The empty string means that application belongs to the 'default' project.",
						MarkdownDescription: "Project is a reference to the project this application belongs to. The empty string means that application belongs to the 'default' project.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"revision_history_limit": schema.Int64Attribute{
						Description:         "RevisionHistoryLimit limits the number of items kept in the application's revision history, which is used for informational purposes as well as for rollbacks to previous versions. This should only be changed in exceptional circumstances. Setting to zero will store no history. This will reduce storage used. Increasing will increase the space used to store the history, so we do not recommend increasing it. Default is 10.",
						MarkdownDescription: "RevisionHistoryLimit limits the number of items kept in the application's revision history, which is used for informational purposes as well as for rollbacks to previous versions. This should only be changed in exceptional circumstances. Setting to zero will store no history. This will reduce storage used. Increasing will increase the space used to store the history, so we do not recommend increasing it. Default is 10.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"source": schema.SingleNestedAttribute{
						Description:         "Source is a reference to the location of the application's manifests or chart",
						MarkdownDescription: "Source is a reference to the location of the application's manifests or chart",
						Attributes: map[string]schema.Attribute{
							"chart": schema.StringAttribute{
								Description:         "Chart is a Helm chart name, and must be specified for applications sourced from a Helm repo.",
								MarkdownDescription: "Chart is a Helm chart name, and must be specified for applications sourced from a Helm repo.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"directory": schema.SingleNestedAttribute{
								Description:         "Directory holds path/directory specific options",
								MarkdownDescription: "Directory holds path/directory specific options",
								Attributes: map[string]schema.Attribute{
									"exclude": schema.StringAttribute{
										Description:         "Exclude contains a glob pattern to match paths against that should be explicitly excluded from being used during manifest generation",
										MarkdownDescription: "Exclude contains a glob pattern to match paths against that should be explicitly excluded from being used during manifest generation",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"include": schema.StringAttribute{
										Description:         "Include contains a glob pattern to match paths against that should be explicitly included during manifest generation",
										MarkdownDescription: "Include contains a glob pattern to match paths against that should be explicitly included during manifest generation",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"jsonnet": schema.SingleNestedAttribute{
										Description:         "Jsonnet holds options specific to Jsonnet",
										MarkdownDescription: "Jsonnet holds options specific to Jsonnet",
										Attributes: map[string]schema.Attribute{
											"ext_vars": schema.ListNestedAttribute{
												Description:         "ExtVars is a list of Jsonnet External Variables",
												MarkdownDescription: "ExtVars is a list of Jsonnet External Variables",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"code": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
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

														"value": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
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

											"libs": schema.ListAttribute{
												Description:         "Additional library search dirs",
												MarkdownDescription: "Additional library search dirs",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"tlas": schema.ListNestedAttribute{
												Description:         "TLAS is a list of Jsonnet Top-level Arguments",
												MarkdownDescription: "TLAS is a list of Jsonnet Top-level Arguments",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"code": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
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

														"value": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
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

									"recurse": schema.BoolAttribute{
										Description:         "Recurse specifies whether to scan a directory recursively for manifests",
										MarkdownDescription: "Recurse specifies whether to scan a directory recursively for manifests",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"helm": schema.SingleNestedAttribute{
								Description:         "Helm holds helm specific options",
								MarkdownDescription: "Helm holds helm specific options",
								Attributes: map[string]schema.Attribute{
									"file_parameters": schema.ListNestedAttribute{
										Description:         "FileParameters are file parameters to the helm template",
										MarkdownDescription: "FileParameters are file parameters to the helm template",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name is the name of the Helm parameter",
													MarkdownDescription: "Name is the name of the Helm parameter",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"path": schema.StringAttribute{
													Description:         "Path is the path to the file containing the values for the Helm parameter",
													MarkdownDescription: "Path is the path to the file containing the values for the Helm parameter",
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

									"ignore_missing_value_files": schema.BoolAttribute{
										Description:         "IgnoreMissingValueFiles prevents helm template from failing when valueFiles do not exist locally by not appending them to helm template --values",
										MarkdownDescription: "IgnoreMissingValueFiles prevents helm template from failing when valueFiles do not exist locally by not appending them to helm template --values",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"parameters": schema.ListNestedAttribute{
										Description:         "Parameters is a list of Helm parameters which are passed to the helm template command upon manifest generation",
										MarkdownDescription: "Parameters is a list of Helm parameters which are passed to the helm template command upon manifest generation",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"force_string": schema.BoolAttribute{
													Description:         "ForceString determines whether to tell Helm to interpret booleans and numbers as strings",
													MarkdownDescription: "ForceString determines whether to tell Helm to interpret booleans and numbers as strings",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name is the name of the Helm parameter",
													MarkdownDescription: "Name is the name of the Helm parameter",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "Value is the value for the Helm parameter",
													MarkdownDescription: "Value is the value for the Helm parameter",
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

									"pass_credentials": schema.BoolAttribute{
										Description:         "PassCredentials pass credentials to all domains (Helm's --pass-credentials)",
										MarkdownDescription: "PassCredentials pass credentials to all domains (Helm's --pass-credentials)",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"release_name": schema.StringAttribute{
										Description:         "ReleaseName is the Helm release name to use. If omitted it will use the application name",
										MarkdownDescription: "ReleaseName is the Helm release name to use. If omitted it will use the application name",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"skip_crds": schema.BoolAttribute{
										Description:         "SkipCrds skips custom resource definition installation step (Helm's --skip-crds)",
										MarkdownDescription: "SkipCrds skips custom resource definition installation step (Helm's --skip-crds)",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"value_files": schema.ListAttribute{
										Description:         "ValuesFiles is a list of Helm value files to use when generating a template",
										MarkdownDescription: "ValuesFiles is a list of Helm value files to use when generating a template",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"values": schema.StringAttribute{
										Description:         "Values specifies Helm values to be passed to helm template, typically defined as a block. ValuesObject takes precedence over Values, so use one or the other.",
										MarkdownDescription: "Values specifies Helm values to be passed to helm template, typically defined as a block. ValuesObject takes precedence over Values, so use one or the other.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"values_object": schema.MapAttribute{
										Description:         "ValuesObject specifies Helm values to be passed to helm template, defined as a map. This takes precedence over Values.",
										MarkdownDescription: "ValuesObject specifies Helm values to be passed to helm template, defined as a map. This takes precedence over Values.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"version": schema.StringAttribute{
										Description:         "Version is the Helm version to use for templating ('3')",
										MarkdownDescription: "Version is the Helm version to use for templating ('3')",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"kustomize": schema.SingleNestedAttribute{
								Description:         "Kustomize holds kustomize specific options",
								MarkdownDescription: "Kustomize holds kustomize specific options",
								Attributes: map[string]schema.Attribute{
									"common_annotations": schema.MapAttribute{
										Description:         "CommonAnnotations is a list of additional annotations to add to rendered manifests",
										MarkdownDescription: "CommonAnnotations is a list of additional annotations to add to rendered manifests",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"common_annotations_envsubst": schema.BoolAttribute{
										Description:         "CommonAnnotationsEnvsubst specifies whether to apply env variables substitution for annotation values",
										MarkdownDescription: "CommonAnnotationsEnvsubst specifies whether to apply env variables substitution for annotation values",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"common_labels": schema.MapAttribute{
										Description:         "CommonLabels is a list of additional labels to add to rendered manifests",
										MarkdownDescription: "CommonLabels is a list of additional labels to add to rendered manifests",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"components": schema.ListAttribute{
										Description:         "Components specifies a list of kustomize components to add to the kustomization before building",
										MarkdownDescription: "Components specifies a list of kustomize components to add to the kustomization before building",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"force_common_annotations": schema.BoolAttribute{
										Description:         "ForceCommonAnnotations specifies whether to force applying common annotations to resources for Kustomize apps",
										MarkdownDescription: "ForceCommonAnnotations specifies whether to force applying common annotations to resources for Kustomize apps",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"force_common_labels": schema.BoolAttribute{
										Description:         "ForceCommonLabels specifies whether to force applying common labels to resources for Kustomize apps",
										MarkdownDescription: "ForceCommonLabels specifies whether to force applying common labels to resources for Kustomize apps",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"images": schema.ListAttribute{
										Description:         "Images is a list of Kustomize image override specifications",
										MarkdownDescription: "Images is a list of Kustomize image override specifications",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"name_prefix": schema.StringAttribute{
										Description:         "NamePrefix is a prefix appended to resources for Kustomize apps",
										MarkdownDescription: "NamePrefix is a prefix appended to resources for Kustomize apps",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"name_suffix": schema.StringAttribute{
										Description:         "NameSuffix is a suffix appended to resources for Kustomize apps",
										MarkdownDescription: "NameSuffix is a suffix appended to resources for Kustomize apps",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
										Description:         "Namespace sets the namespace that Kustomize adds to all resources",
										MarkdownDescription: "Namespace sets the namespace that Kustomize adds to all resources",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"patches": schema.ListNestedAttribute{
										Description:         "Patches is a list of Kustomize patches",
										MarkdownDescription: "Patches is a list of Kustomize patches",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"options": schema.MapAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"patch": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"path": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"target": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"annotation_selector": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

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

														"label_selector": schema.StringAttribute{
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

									"replicas": schema.ListNestedAttribute{
										Description:         "Replicas is a list of Kustomize Replicas override specifications",
										MarkdownDescription: "Replicas is a list of Kustomize Replicas override specifications",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"count": schema.StringAttribute{
													Description:         "Number of replicas",
													MarkdownDescription: "Number of replicas",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name of Deployment or StatefulSet",
													MarkdownDescription: "Name of Deployment or StatefulSet",
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

									"version": schema.StringAttribute{
										Description:         "Version controls which version of Kustomize to use for rendering manifests",
										MarkdownDescription: "Version controls which version of Kustomize to use for rendering manifests",
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
								Description:         "Path is a directory path within the Git repository, and is only valid for applications sourced from Git.",
								MarkdownDescription: "Path is a directory path within the Git repository, and is only valid for applications sourced from Git.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"plugin": schema.SingleNestedAttribute{
								Description:         "Plugin holds config management plugin specific options",
								MarkdownDescription: "Plugin holds config management plugin specific options",
								Attributes: map[string]schema.Attribute{
									"env": schema.ListNestedAttribute{
										Description:         "Env is a list of environment variable entries",
										MarkdownDescription: "Env is a list of environment variable entries",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name is the name of the variable, usually expressed in uppercase",
													MarkdownDescription: "Name is the name of the variable, usually expressed in uppercase",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "Value is the value of the variable",
													MarkdownDescription: "Value is the value of the variable",
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
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"parameters": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"array": schema.ListAttribute{
													Description:         "Array is the value of an array type parameter.",
													MarkdownDescription: "Array is the value of an array type parameter.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"map": schema.MapAttribute{
													Description:         "Map is the value of a map type parameter.",
													MarkdownDescription: "Map is the value of a map type parameter.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name is the name identifying a parameter.",
													MarkdownDescription: "Name is the name identifying a parameter.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"string": schema.StringAttribute{
													Description:         "String_ is the value of a string type parameter.",
													MarkdownDescription: "String_ is the value of a string type parameter.",
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

							"ref": schema.StringAttribute{
								Description:         "Ref is reference to another source within sources field. This field will not be used if used with a 'source' tag.",
								MarkdownDescription: "Ref is reference to another source within sources field. This field will not be used if used with a 'source' tag.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"repo_url": schema.StringAttribute{
								Description:         "RepoURL is the URL to the repository (Git or Helm) that contains the application manifests",
								MarkdownDescription: "RepoURL is the URL to the repository (Git or Helm) that contains the application manifests",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"target_revision": schema.StringAttribute{
								Description:         "TargetRevision defines the revision of the source to sync the application to. In case of Git, this can be commit, tag, or branch. If omitted, will equal to HEAD. In case of Helm, this is a semver tag for the Chart's version.",
								MarkdownDescription: "TargetRevision defines the revision of the source to sync the application to. In case of Git, this can be commit, tag, or branch. If omitted, will equal to HEAD. In case of Helm, this is a semver tag for the Chart's version.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"sources": schema.ListNestedAttribute{
						Description:         "Sources is a reference to the location of the application's manifests or chart",
						MarkdownDescription: "Sources is a reference to the location of the application's manifests or chart",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"chart": schema.StringAttribute{
									Description:         "Chart is a Helm chart name, and must be specified for applications sourced from a Helm repo.",
									MarkdownDescription: "Chart is a Helm chart name, and must be specified for applications sourced from a Helm repo.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"directory": schema.SingleNestedAttribute{
									Description:         "Directory holds path/directory specific options",
									MarkdownDescription: "Directory holds path/directory specific options",
									Attributes: map[string]schema.Attribute{
										"exclude": schema.StringAttribute{
											Description:         "Exclude contains a glob pattern to match paths against that should be explicitly excluded from being used during manifest generation",
											MarkdownDescription: "Exclude contains a glob pattern to match paths against that should be explicitly excluded from being used during manifest generation",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"include": schema.StringAttribute{
											Description:         "Include contains a glob pattern to match paths against that should be explicitly included during manifest generation",
											MarkdownDescription: "Include contains a glob pattern to match paths against that should be explicitly included during manifest generation",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"jsonnet": schema.SingleNestedAttribute{
											Description:         "Jsonnet holds options specific to Jsonnet",
											MarkdownDescription: "Jsonnet holds options specific to Jsonnet",
											Attributes: map[string]schema.Attribute{
												"ext_vars": schema.ListNestedAttribute{
													Description:         "ExtVars is a list of Jsonnet External Variables",
													MarkdownDescription: "ExtVars is a list of Jsonnet External Variables",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"code": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
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

															"value": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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

												"libs": schema.ListAttribute{
													Description:         "Additional library search dirs",
													MarkdownDescription: "Additional library search dirs",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"tlas": schema.ListNestedAttribute{
													Description:         "TLAS is a list of Jsonnet Top-level Arguments",
													MarkdownDescription: "TLAS is a list of Jsonnet Top-level Arguments",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"code": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
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

															"value": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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

										"recurse": schema.BoolAttribute{
											Description:         "Recurse specifies whether to scan a directory recursively for manifests",
											MarkdownDescription: "Recurse specifies whether to scan a directory recursively for manifests",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"helm": schema.SingleNestedAttribute{
									Description:         "Helm holds helm specific options",
									MarkdownDescription: "Helm holds helm specific options",
									Attributes: map[string]schema.Attribute{
										"file_parameters": schema.ListNestedAttribute{
											Description:         "FileParameters are file parameters to the helm template",
											MarkdownDescription: "FileParameters are file parameters to the helm template",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "Name is the name of the Helm parameter",
														MarkdownDescription: "Name is the name of the Helm parameter",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"path": schema.StringAttribute{
														Description:         "Path is the path to the file containing the values for the Helm parameter",
														MarkdownDescription: "Path is the path to the file containing the values for the Helm parameter",
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

										"ignore_missing_value_files": schema.BoolAttribute{
											Description:         "IgnoreMissingValueFiles prevents helm template from failing when valueFiles do not exist locally by not appending them to helm template --values",
											MarkdownDescription: "IgnoreMissingValueFiles prevents helm template from failing when valueFiles do not exist locally by not appending them to helm template --values",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"parameters": schema.ListNestedAttribute{
											Description:         "Parameters is a list of Helm parameters which are passed to the helm template command upon manifest generation",
											MarkdownDescription: "Parameters is a list of Helm parameters which are passed to the helm template command upon manifest generation",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"force_string": schema.BoolAttribute{
														Description:         "ForceString determines whether to tell Helm to interpret booleans and numbers as strings",
														MarkdownDescription: "ForceString determines whether to tell Helm to interpret booleans and numbers as strings",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name is the name of the Helm parameter",
														MarkdownDescription: "Name is the name of the Helm parameter",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"value": schema.StringAttribute{
														Description:         "Value is the value for the Helm parameter",
														MarkdownDescription: "Value is the value for the Helm parameter",
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

										"pass_credentials": schema.BoolAttribute{
											Description:         "PassCredentials pass credentials to all domains (Helm's --pass-credentials)",
											MarkdownDescription: "PassCredentials pass credentials to all domains (Helm's --pass-credentials)",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"release_name": schema.StringAttribute{
											Description:         "ReleaseName is the Helm release name to use. If omitted it will use the application name",
											MarkdownDescription: "ReleaseName is the Helm release name to use. If omitted it will use the application name",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"skip_crds": schema.BoolAttribute{
											Description:         "SkipCrds skips custom resource definition installation step (Helm's --skip-crds)",
											MarkdownDescription: "SkipCrds skips custom resource definition installation step (Helm's --skip-crds)",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"value_files": schema.ListAttribute{
											Description:         "ValuesFiles is a list of Helm value files to use when generating a template",
											MarkdownDescription: "ValuesFiles is a list of Helm value files to use when generating a template",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"values": schema.StringAttribute{
											Description:         "Values specifies Helm values to be passed to helm template, typically defined as a block. ValuesObject takes precedence over Values, so use one or the other.",
											MarkdownDescription: "Values specifies Helm values to be passed to helm template, typically defined as a block. ValuesObject takes precedence over Values, so use one or the other.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"values_object": schema.MapAttribute{
											Description:         "ValuesObject specifies Helm values to be passed to helm template, defined as a map. This takes precedence over Values.",
											MarkdownDescription: "ValuesObject specifies Helm values to be passed to helm template, defined as a map. This takes precedence over Values.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"version": schema.StringAttribute{
											Description:         "Version is the Helm version to use for templating ('3')",
											MarkdownDescription: "Version is the Helm version to use for templating ('3')",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"kustomize": schema.SingleNestedAttribute{
									Description:         "Kustomize holds kustomize specific options",
									MarkdownDescription: "Kustomize holds kustomize specific options",
									Attributes: map[string]schema.Attribute{
										"common_annotations": schema.MapAttribute{
											Description:         "CommonAnnotations is a list of additional annotations to add to rendered manifests",
											MarkdownDescription: "CommonAnnotations is a list of additional annotations to add to rendered manifests",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"common_annotations_envsubst": schema.BoolAttribute{
											Description:         "CommonAnnotationsEnvsubst specifies whether to apply env variables substitution for annotation values",
											MarkdownDescription: "CommonAnnotationsEnvsubst specifies whether to apply env variables substitution for annotation values",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"common_labels": schema.MapAttribute{
											Description:         "CommonLabels is a list of additional labels to add to rendered manifests",
											MarkdownDescription: "CommonLabels is a list of additional labels to add to rendered manifests",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"components": schema.ListAttribute{
											Description:         "Components specifies a list of kustomize components to add to the kustomization before building",
											MarkdownDescription: "Components specifies a list of kustomize components to add to the kustomization before building",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"force_common_annotations": schema.BoolAttribute{
											Description:         "ForceCommonAnnotations specifies whether to force applying common annotations to resources for Kustomize apps",
											MarkdownDescription: "ForceCommonAnnotations specifies whether to force applying common annotations to resources for Kustomize apps",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"force_common_labels": schema.BoolAttribute{
											Description:         "ForceCommonLabels specifies whether to force applying common labels to resources for Kustomize apps",
											MarkdownDescription: "ForceCommonLabels specifies whether to force applying common labels to resources for Kustomize apps",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"images": schema.ListAttribute{
											Description:         "Images is a list of Kustomize image override specifications",
											MarkdownDescription: "Images is a list of Kustomize image override specifications",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name_prefix": schema.StringAttribute{
											Description:         "NamePrefix is a prefix appended to resources for Kustomize apps",
											MarkdownDescription: "NamePrefix is a prefix appended to resources for Kustomize apps",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name_suffix": schema.StringAttribute{
											Description:         "NameSuffix is a suffix appended to resources for Kustomize apps",
											MarkdownDescription: "NameSuffix is a suffix appended to resources for Kustomize apps",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"namespace": schema.StringAttribute{
											Description:         "Namespace sets the namespace that Kustomize adds to all resources",
											MarkdownDescription: "Namespace sets the namespace that Kustomize adds to all resources",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"patches": schema.ListNestedAttribute{
											Description:         "Patches is a list of Kustomize patches",
											MarkdownDescription: "Patches is a list of Kustomize patches",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"options": schema.MapAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"patch": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"path": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"target": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"annotation_selector": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

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

															"label_selector": schema.StringAttribute{
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

										"replicas": schema.ListNestedAttribute{
											Description:         "Replicas is a list of Kustomize Replicas override specifications",
											MarkdownDescription: "Replicas is a list of Kustomize Replicas override specifications",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"count": schema.StringAttribute{
														Description:         "Number of replicas",
														MarkdownDescription: "Number of replicas",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of Deployment or StatefulSet",
														MarkdownDescription: "Name of Deployment or StatefulSet",
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

										"version": schema.StringAttribute{
											Description:         "Version controls which version of Kustomize to use for rendering manifests",
											MarkdownDescription: "Version controls which version of Kustomize to use for rendering manifests",
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
									Description:         "Path is a directory path within the Git repository, and is only valid for applications sourced from Git.",
									MarkdownDescription: "Path is a directory path within the Git repository, and is only valid for applications sourced from Git.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"plugin": schema.SingleNestedAttribute{
									Description:         "Plugin holds config management plugin specific options",
									MarkdownDescription: "Plugin holds config management plugin specific options",
									Attributes: map[string]schema.Attribute{
										"env": schema.ListNestedAttribute{
											Description:         "Env is a list of environment variable entries",
											MarkdownDescription: "Env is a list of environment variable entries",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "Name is the name of the variable, usually expressed in uppercase",
														MarkdownDescription: "Name is the name of the variable, usually expressed in uppercase",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"value": schema.StringAttribute{
														Description:         "Value is the value of the variable",
														MarkdownDescription: "Value is the value of the variable",
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
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"parameters": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"array": schema.ListAttribute{
														Description:         "Array is the value of an array type parameter.",
														MarkdownDescription: "Array is the value of an array type parameter.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"map": schema.MapAttribute{
														Description:         "Map is the value of a map type parameter.",
														MarkdownDescription: "Map is the value of a map type parameter.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name is the name identifying a parameter.",
														MarkdownDescription: "Name is the name identifying a parameter.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"string": schema.StringAttribute{
														Description:         "String_ is the value of a string type parameter.",
														MarkdownDescription: "String_ is the value of a string type parameter.",
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

								"ref": schema.StringAttribute{
									Description:         "Ref is reference to another source within sources field. This field will not be used if used with a 'source' tag.",
									MarkdownDescription: "Ref is reference to another source within sources field. This field will not be used if used with a 'source' tag.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"repo_url": schema.StringAttribute{
									Description:         "RepoURL is the URL to the repository (Git or Helm) that contains the application manifests",
									MarkdownDescription: "RepoURL is the URL to the repository (Git or Helm) that contains the application manifests",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"target_revision": schema.StringAttribute{
									Description:         "TargetRevision defines the revision of the source to sync the application to. In case of Git, this can be commit, tag, or branch. If omitted, will equal to HEAD. In case of Helm, this is a semver tag for the Chart's version.",
									MarkdownDescription: "TargetRevision defines the revision of the source to sync the application to. In case of Git, this can be commit, tag, or branch. If omitted, will equal to HEAD. In case of Helm, this is a semver tag for the Chart's version.",
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

					"sync_policy": schema.SingleNestedAttribute{
						Description:         "SyncPolicy controls when and how a sync will be performed",
						MarkdownDescription: "SyncPolicy controls when and how a sync will be performed",
						Attributes: map[string]schema.Attribute{
							"automated": schema.SingleNestedAttribute{
								Description:         "Automated will keep an application synced to the target revision",
								MarkdownDescription: "Automated will keep an application synced to the target revision",
								Attributes: map[string]schema.Attribute{
									"allow_empty": schema.BoolAttribute{
										Description:         "AllowEmpty allows apps have zero live resources (default: false)",
										MarkdownDescription: "AllowEmpty allows apps have zero live resources (default: false)",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"prune": schema.BoolAttribute{
										Description:         "Prune specifies whether to delete resources from the cluster that are not found in the sources anymore as part of automated sync (default: false)",
										MarkdownDescription: "Prune specifies whether to delete resources from the cluster that are not found in the sources anymore as part of automated sync (default: false)",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"self_heal": schema.BoolAttribute{
										Description:         "SelfHeal specifies whether to revert resources back to their desired state upon modification in the cluster (default: false)",
										MarkdownDescription: "SelfHeal specifies whether to revert resources back to their desired state upon modification in the cluster (default: false)",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"managed_namespace_metadata": schema.SingleNestedAttribute{
								Description:         "ManagedNamespaceMetadata controls metadata in the given namespace (if CreateNamespace=true)",
								MarkdownDescription: "ManagedNamespaceMetadata controls metadata in the given namespace (if CreateNamespace=true)",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.MapAttribute{
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

							"retry": schema.SingleNestedAttribute{
								Description:         "Retry controls failed sync retry behavior",
								MarkdownDescription: "Retry controls failed sync retry behavior",
								Attributes: map[string]schema.Attribute{
									"backoff": schema.SingleNestedAttribute{
										Description:         "Backoff controls how to backoff on subsequent retries of failed syncs",
										MarkdownDescription: "Backoff controls how to backoff on subsequent retries of failed syncs",
										Attributes: map[string]schema.Attribute{
											"duration": schema.StringAttribute{
												Description:         "Duration is the amount to back off. Default unit is seconds, but could also be a duration (e.g. '2m', '1h')",
												MarkdownDescription: "Duration is the amount to back off. Default unit is seconds, but could also be a duration (e.g. '2m', '1h')",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"factor": schema.Int64Attribute{
												Description:         "Factor is a factor to multiply the base duration after each failed retry",
												MarkdownDescription: "Factor is a factor to multiply the base duration after each failed retry",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_duration": schema.StringAttribute{
												Description:         "MaxDuration is the maximum amount of time allowed for the backoff strategy",
												MarkdownDescription: "MaxDuration is the maximum amount of time allowed for the backoff strategy",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"limit": schema.Int64Attribute{
										Description:         "Limit is the maximum number of attempts for retrying a failed sync. If set to 0, no retries will be performed.",
										MarkdownDescription: "Limit is the maximum number of attempts for retrying a failed sync. If set to 0, no retries will be performed.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"sync_options": schema.ListAttribute{
								Description:         "Options allow you to specify whole app sync-options",
								MarkdownDescription: "Options allow you to specify whole app sync-options",
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
		},
	}
}

func (r *ArgoprojIoApplicationV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_argoproj_io_application_v1alpha1_manifest")

	var model ArgoprojIoApplicationV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("argoproj.io/v1alpha1")
	model.Kind = pointer.String("Application")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
