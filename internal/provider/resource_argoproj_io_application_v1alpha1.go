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

type ArgoprojIoApplicationV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*ArgoprojIoApplicationV1Alpha1Resource)(nil)
)

type ArgoprojIoApplicationV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Operation  types.Object `tfsdk:"operation"`
	Spec       types.Object `tfsdk:"spec"`
}

type ArgoprojIoApplicationV1Alpha1GoModel struct {
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

	Operation *struct {
		Info *[]struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Value *string `tfsdk:"value" yaml:"value,omitempty"`
		} `tfsdk:"info" yaml:"info,omitempty"`

		InitiatedBy *struct {
			Automated *bool `tfsdk:"automated" yaml:"automated,omitempty"`

			Username *string `tfsdk:"username" yaml:"username,omitempty"`
		} `tfsdk:"initiated_by" yaml:"initiatedBy,omitempty"`

		Retry *struct {
			Backoff *struct {
				MaxDuration *string `tfsdk:"max_duration" yaml:"maxDuration,omitempty"`

				Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

				Factor *int64 `tfsdk:"factor" yaml:"factor,omitempty"`
			} `tfsdk:"backoff" yaml:"backoff,omitempty"`

			Limit *int64 `tfsdk:"limit" yaml:"limit,omitempty"`
		} `tfsdk:"retry" yaml:"retry,omitempty"`

		Sync *struct {
			Revision *string `tfsdk:"revision" yaml:"revision,omitempty"`

			Source *struct {
				Helm *struct {
					IgnoreMissingValueFiles *bool `tfsdk:"ignore_missing_value_files" yaml:"ignoreMissingValueFiles,omitempty"`

					Parameters *[]struct {
						ForceString *bool `tfsdk:"force_string" yaml:"forceString,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"parameters" yaml:"parameters,omitempty"`

					PassCredentials *bool `tfsdk:"pass_credentials" yaml:"passCredentials,omitempty"`

					Values *string `tfsdk:"values" yaml:"values,omitempty"`

					FileParameters *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Path *string `tfsdk:"path" yaml:"path,omitempty"`
					} `tfsdk:"file_parameters" yaml:"fileParameters,omitempty"`

					ReleaseName *string `tfsdk:"release_name" yaml:"releaseName,omitempty"`

					SkipCrds *bool `tfsdk:"skip_crds" yaml:"skipCrds,omitempty"`

					ValueFiles *[]string `tfsdk:"value_files" yaml:"valueFiles,omitempty"`

					Version *string `tfsdk:"version" yaml:"version,omitempty"`
				} `tfsdk:"helm" yaml:"helm,omitempty"`

				Kustomize *struct {
					NameSuffix *string `tfsdk:"name_suffix" yaml:"nameSuffix,omitempty"`

					Version *string `tfsdk:"version" yaml:"version,omitempty"`

					CommonAnnotations *map[string]string `tfsdk:"common_annotations" yaml:"commonAnnotations,omitempty"`

					CommonLabels *map[string]string `tfsdk:"common_labels" yaml:"commonLabels,omitempty"`

					ForceCommonAnnotations *bool `tfsdk:"force_common_annotations" yaml:"forceCommonAnnotations,omitempty"`

					ForceCommonLabels *bool `tfsdk:"force_common_labels" yaml:"forceCommonLabels,omitempty"`

					Images *[]string `tfsdk:"images" yaml:"images,omitempty"`

					NamePrefix *string `tfsdk:"name_prefix" yaml:"namePrefix,omitempty"`
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

				Chart *string `tfsdk:"chart" yaml:"chart,omitempty"`

				Directory *struct {
					Exclude *string `tfsdk:"exclude" yaml:"exclude,omitempty"`

					Include *string `tfsdk:"include" yaml:"include,omitempty"`

					Jsonnet *struct {
						Tlas *[]struct {
							Code *bool `tfsdk:"code" yaml:"code,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"tlas" yaml:"tlas,omitempty"`

						ExtVars *[]struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`

							Code *bool `tfsdk:"code" yaml:"code,omitempty"`
						} `tfsdk:"ext_vars" yaml:"extVars,omitempty"`

						Libs *[]string `tfsdk:"libs" yaml:"libs,omitempty"`
					} `tfsdk:"jsonnet" yaml:"jsonnet,omitempty"`

					Recurse *bool `tfsdk:"recurse" yaml:"recurse,omitempty"`
				} `tfsdk:"directory" yaml:"directory,omitempty"`
			} `tfsdk:"source" yaml:"source,omitempty"`

			SyncOptions *[]string `tfsdk:"sync_options" yaml:"syncOptions,omitempty"`

			SyncStrategy *struct {
				Hook *struct {
					Force *bool `tfsdk:"force" yaml:"force,omitempty"`
				} `tfsdk:"hook" yaml:"hook,omitempty"`

				Apply *struct {
					Force *bool `tfsdk:"force" yaml:"force,omitempty"`
				} `tfsdk:"apply" yaml:"apply,omitempty"`
			} `tfsdk:"sync_strategy" yaml:"syncStrategy,omitempty"`

			DryRun *bool `tfsdk:"dry_run" yaml:"dryRun,omitempty"`

			Manifests *[]string `tfsdk:"manifests" yaml:"manifests,omitempty"`

			Prune *bool `tfsdk:"prune" yaml:"prune,omitempty"`

			Resources *[]struct {
				Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

				Group *string `tfsdk:"group" yaml:"group,omitempty"`
			} `tfsdk:"resources" yaml:"resources,omitempty"`
		} `tfsdk:"sync" yaml:"sync,omitempty"`
	} `tfsdk:"operation" yaml:"operation,omitempty"`

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
			TargetRevision *string `tfsdk:"target_revision" yaml:"targetRevision,omitempty"`

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
				Values *string `tfsdk:"values" yaml:"values,omitempty"`

				Version *string `tfsdk:"version" yaml:"version,omitempty"`

				Parameters *[]struct {
					ForceString *bool `tfsdk:"force_string" yaml:"forceString,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"parameters" yaml:"parameters,omitempty"`

				PassCredentials *bool `tfsdk:"pass_credentials" yaml:"passCredentials,omitempty"`

				ReleaseName *string `tfsdk:"release_name" yaml:"releaseName,omitempty"`

				SkipCrds *bool `tfsdk:"skip_crds" yaml:"skipCrds,omitempty"`

				ValueFiles *[]string `tfsdk:"value_files" yaml:"valueFiles,omitempty"`

				FileParameters *[]struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Path *string `tfsdk:"path" yaml:"path,omitempty"`
				} `tfsdk:"file_parameters" yaml:"fileParameters,omitempty"`

				IgnoreMissingValueFiles *bool `tfsdk:"ignore_missing_value_files" yaml:"ignoreMissingValueFiles,omitempty"`
			} `tfsdk:"helm" yaml:"helm,omitempty"`

			Kustomize *struct {
				NamePrefix *string `tfsdk:"name_prefix" yaml:"namePrefix,omitempty"`

				NameSuffix *string `tfsdk:"name_suffix" yaml:"nameSuffix,omitempty"`

				Version *string `tfsdk:"version" yaml:"version,omitempty"`

				CommonAnnotations *map[string]string `tfsdk:"common_annotations" yaml:"commonAnnotations,omitempty"`

				CommonLabels *map[string]string `tfsdk:"common_labels" yaml:"commonLabels,omitempty"`

				ForceCommonAnnotations *bool `tfsdk:"force_common_annotations" yaml:"forceCommonAnnotations,omitempty"`

				ForceCommonLabels *bool `tfsdk:"force_common_labels" yaml:"forceCommonLabels,omitempty"`

				Images *[]string `tfsdk:"images" yaml:"images,omitempty"`
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
		} `tfsdk:"source" yaml:"source,omitempty"`

		SyncPolicy *struct {
			Automated *struct {
				Prune *bool `tfsdk:"prune" yaml:"prune,omitempty"`

				SelfHeal *bool `tfsdk:"self_heal" yaml:"selfHeal,omitempty"`

				AllowEmpty *bool `tfsdk:"allow_empty" yaml:"allowEmpty,omitempty"`
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
}

func NewArgoprojIoApplicationV1Alpha1Resource() resource.Resource {
	return &ArgoprojIoApplicationV1Alpha1Resource{}
}

func (r *ArgoprojIoApplicationV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_argoproj_io_application_v1alpha1"
}

func (r *ArgoprojIoApplicationV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Application is a definition of Application resource.",
		MarkdownDescription: "Application is a definition of Application resource.",
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
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.RequiresReplace(),
						},
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

			"operation": {
				Description:         "Operation contains information about a requested or running operation",
				MarkdownDescription: "Operation contains information about a requested or running operation",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"info": {
						Description:         "Info is a list of informational items for this operation",
						MarkdownDescription: "Info is a list of informational items for this operation",

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

					"initiated_by": {
						Description:         "InitiatedBy contains information about who initiated the operations",
						MarkdownDescription: "InitiatedBy contains information about who initiated the operations",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"automated": {
								Description:         "Automated is set to true if operation was initiated automatically by the application controller.",
								MarkdownDescription: "Automated is set to true if operation was initiated automatically by the application controller.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"username": {
								Description:         "Username contains the name of a user who started operation",
								MarkdownDescription: "Username contains the name of a user who started operation",

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

					"retry": {
						Description:         "Retry controls the strategy to apply if a sync fails",
						MarkdownDescription: "Retry controls the strategy to apply if a sync fails",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"backoff": {
								Description:         "Backoff controls how to backoff on subsequent retries of failed syncs",
								MarkdownDescription: "Backoff controls how to backoff on subsequent retries of failed syncs",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"max_duration": {
										Description:         "MaxDuration is the maximum amount of time allowed for the backoff strategy",
										MarkdownDescription: "MaxDuration is the maximum amount of time allowed for the backoff strategy",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"duration": {
										Description:         "Duration is the amount to back off. Default unit is seconds, but could also be a duration (e.g. '2m', '1h')",
										MarkdownDescription: "Duration is the amount to back off. Default unit is seconds, but could also be a duration (e.g. '2m', '1h')",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"factor": {
										Description:         "Factor is a factor to multiply the base duration after each failed retry",
										MarkdownDescription: "Factor is a factor to multiply the base duration after each failed retry",

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

							"limit": {
								Description:         "Limit is the maximum number of attempts for retrying a failed sync. If set to 0, no retries will be performed.",
								MarkdownDescription: "Limit is the maximum number of attempts for retrying a failed sync. If set to 0, no retries will be performed.",

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

					"sync": {
						Description:         "Sync contains parameters for the operation",
						MarkdownDescription: "Sync contains parameters for the operation",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"revision": {
								Description:         "Revision is the revision (Git) or chart version (Helm) which to sync the application to If omitted, will use the revision specified in app spec.",
								MarkdownDescription: "Revision is the revision (Git) or chart version (Helm) which to sync the application to If omitted, will use the revision specified in app spec.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"source": {
								Description:         "Source overrides the source definition set in the application. This is typically set in a Rollback operation and is nil during a Sync operation",
								MarkdownDescription: "Source overrides the source definition set in the application. This is typically set in a Rollback operation and is nil during a Sync operation",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"helm": {
										Description:         "Helm holds helm specific options",
										MarkdownDescription: "Helm holds helm specific options",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"ignore_missing_value_files": {
												Description:         "IgnoreMissingValueFiles prevents helm template from failing when valueFiles do not exist locally by not appending them to helm template --values",
												MarkdownDescription: "IgnoreMissingValueFiles prevents helm template from failing when valueFiles do not exist locally by not appending them to helm template --values",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"parameters": {
												Description:         "Parameters is a list of Helm parameters which are passed to the helm template command upon manifest generation",
												MarkdownDescription: "Parameters is a list of Helm parameters which are passed to the helm template command upon manifest generation",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"force_string": {
														Description:         "ForceString determines whether to tell Helm to interpret booleans and numbers as strings",
														MarkdownDescription: "ForceString determines whether to tell Helm to interpret booleans and numbers as strings",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": {
														Description:         "Name is the name of the Helm parameter",
														MarkdownDescription: "Name is the name of the Helm parameter",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"value": {
														Description:         "Value is the value for the Helm parameter",
														MarkdownDescription: "Value is the value for the Helm parameter",

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
												Description:         "PassCredentials pass credentials to all domains (Helm's --pass-credentials)",
												MarkdownDescription: "PassCredentials pass credentials to all domains (Helm's --pass-credentials)",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"values": {
												Description:         "Values specifies Helm values to be passed to helm template, typically defined as a block",
												MarkdownDescription: "Values specifies Helm values to be passed to helm template, typically defined as a block",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"file_parameters": {
												Description:         "FileParameters are file parameters to the helm template",
												MarkdownDescription: "FileParameters are file parameters to the helm template",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "Name is the name of the Helm parameter",
														MarkdownDescription: "Name is the name of the Helm parameter",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"path": {
														Description:         "Path is the path to the file containing the values for the Helm parameter",
														MarkdownDescription: "Path is the path to the file containing the values for the Helm parameter",

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

											"release_name": {
												Description:         "ReleaseName is the Helm release name to use. If omitted it will use the application name",
												MarkdownDescription: "ReleaseName is the Helm release name to use. If omitted it will use the application name",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"skip_crds": {
												Description:         "SkipCrds skips custom resource definition installation step (Helm's --skip-crds)",
												MarkdownDescription: "SkipCrds skips custom resource definition installation step (Helm's --skip-crds)",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value_files": {
												Description:         "ValuesFiles is a list of Helm value files to use when generating a template",
												MarkdownDescription: "ValuesFiles is a list of Helm value files to use when generating a template",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"version": {
												Description:         "Version is the Helm version to use for templating ('3')",
												MarkdownDescription: "Version is the Helm version to use for templating ('3')",

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
										Description:         "Kustomize holds kustomize specific options",
										MarkdownDescription: "Kustomize holds kustomize specific options",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name_suffix": {
												Description:         "NameSuffix is a suffix appended to resources for Kustomize apps",
												MarkdownDescription: "NameSuffix is a suffix appended to resources for Kustomize apps",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"version": {
												Description:         "Version controls which version of Kustomize to use for rendering manifests",
												MarkdownDescription: "Version controls which version of Kustomize to use for rendering manifests",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"common_annotations": {
												Description:         "CommonAnnotations is a list of additional annotations to add to rendered manifests",
												MarkdownDescription: "CommonAnnotations is a list of additional annotations to add to rendered manifests",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"common_labels": {
												Description:         "CommonLabels is a list of additional labels to add to rendered manifests",
												MarkdownDescription: "CommonLabels is a list of additional labels to add to rendered manifests",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"force_common_annotations": {
												Description:         "ForceCommonAnnotations specifies whether to force applying common annotations to resources for Kustomize apps",
												MarkdownDescription: "ForceCommonAnnotations specifies whether to force applying common annotations to resources for Kustomize apps",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"force_common_labels": {
												Description:         "ForceCommonLabels specifies whether to force applying common labels to resources for Kustomize apps",
												MarkdownDescription: "ForceCommonLabels specifies whether to force applying common labels to resources for Kustomize apps",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"images": {
												Description:         "Images is a list of Kustomize image override specifications",
												MarkdownDescription: "Images is a list of Kustomize image override specifications",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name_prefix": {
												Description:         "NamePrefix is a prefix appended to resources for Kustomize apps",
												MarkdownDescription: "NamePrefix is a prefix appended to resources for Kustomize apps",

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
										Description:         "Path is a directory path within the Git repository, and is only valid for applications sourced from Git.",
										MarkdownDescription: "Path is a directory path within the Git repository, and is only valid for applications sourced from Git.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"plugin": {
										Description:         "ConfigManagementPlugin holds config management plugin specific options",
										MarkdownDescription: "ConfigManagementPlugin holds config management plugin specific options",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"env": {
												Description:         "Env is a list of environment variable entries",
												MarkdownDescription: "Env is a list of environment variable entries",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "Name is the name of the variable, usually expressed in uppercase",
														MarkdownDescription: "Name is the name of the variable, usually expressed in uppercase",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"value": {
														Description:         "Value is the value of the variable",
														MarkdownDescription: "Value is the value of the variable",

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
										Description:         "RepoURL is the URL to the repository (Git or Helm) that contains the application manifests",
										MarkdownDescription: "RepoURL is the URL to the repository (Git or Helm) that contains the application manifests",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"target_revision": {
										Description:         "TargetRevision defines the revision of the source to sync the application to. In case of Git, this can be commit, tag, or branch. If omitted, will equal to HEAD. In case of Helm, this is a semver tag for the Chart's version.",
										MarkdownDescription: "TargetRevision defines the revision of the source to sync the application to. In case of Git, this can be commit, tag, or branch. If omitted, will equal to HEAD. In case of Helm, this is a semver tag for the Chart's version.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"chart": {
										Description:         "Chart is a Helm chart name, and must be specified for applications sourced from a Helm repo.",
										MarkdownDescription: "Chart is a Helm chart name, and must be specified for applications sourced from a Helm repo.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"directory": {
										Description:         "Directory holds path/directory specific options",
										MarkdownDescription: "Directory holds path/directory specific options",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"exclude": {
												Description:         "Exclude contains a glob pattern to match paths against that should be explicitly excluded from being used during manifest generation",
												MarkdownDescription: "Exclude contains a glob pattern to match paths against that should be explicitly excluded from being used during manifest generation",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"include": {
												Description:         "Include contains a glob pattern to match paths against that should be explicitly included during manifest generation",
												MarkdownDescription: "Include contains a glob pattern to match paths against that should be explicitly included during manifest generation",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"jsonnet": {
												Description:         "Jsonnet holds options specific to Jsonnet",
												MarkdownDescription: "Jsonnet holds options specific to Jsonnet",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"tlas": {
														Description:         "TLAS is a list of Jsonnet Top-level Arguments",
														MarkdownDescription: "TLAS is a list of Jsonnet Top-level Arguments",

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

													"ext_vars": {
														Description:         "ExtVars is a list of Jsonnet External Variables",
														MarkdownDescription: "ExtVars is a list of Jsonnet External Variables",

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

															"code": {
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

													"libs": {
														Description:         "Additional library search dirs",
														MarkdownDescription: "Additional library search dirs",

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

											"recurse": {
												Description:         "Recurse specifies whether to scan a directory recursively for manifests",
												MarkdownDescription: "Recurse specifies whether to scan a directory recursively for manifests",

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

							"sync_options": {
								Description:         "SyncOptions provide per-sync sync-options, e.g. Validate=false",
								MarkdownDescription: "SyncOptions provide per-sync sync-options, e.g. Validate=false",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"sync_strategy": {
								Description:         "SyncStrategy describes how to perform the sync",
								MarkdownDescription: "SyncStrategy describes how to perform the sync",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"hook": {
										Description:         "Hook will submit any referenced resources to perform the sync. This is the default strategy",
										MarkdownDescription: "Hook will submit any referenced resources to perform the sync. This is the default strategy",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"force": {
												Description:         "Force indicates whether or not to supply the --force flag to 'kubectl apply'. The --force flag deletes and re-create the resource, when PATCH encounters conflict and has retried for 5 times.",
												MarkdownDescription: "Force indicates whether or not to supply the --force flag to 'kubectl apply'. The --force flag deletes and re-create the resource, when PATCH encounters conflict and has retried for 5 times.",

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

									"apply": {
										Description:         "Apply will perform a 'kubectl apply' to perform the sync.",
										MarkdownDescription: "Apply will perform a 'kubectl apply' to perform the sync.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"force": {
												Description:         "Force indicates whether or not to supply the --force flag to 'kubectl apply'. The --force flag deletes and re-create the resource, when PATCH encounters conflict and has retried for 5 times.",
												MarkdownDescription: "Force indicates whether or not to supply the --force flag to 'kubectl apply'. The --force flag deletes and re-create the resource, when PATCH encounters conflict and has retried for 5 times.",

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

							"dry_run": {
								Description:         "DryRun specifies to perform a 'kubectl apply --dry-run' without actually performing the sync",
								MarkdownDescription: "DryRun specifies to perform a 'kubectl apply --dry-run' without actually performing the sync",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"manifests": {
								Description:         "Manifests is an optional field that overrides sync source with a local directory for development",
								MarkdownDescription: "Manifests is an optional field that overrides sync source with a local directory for development",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"prune": {
								Description:         "Prune specifies to delete resources from the cluster that are no longer tracked in git",
								MarkdownDescription: "Prune specifies to delete resources from the cluster that are no longer tracked in git",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"resources": {
								Description:         "Resources describes which resources shall be part of the sync",
								MarkdownDescription: "Resources describes which resources shall be part of the sync",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"kind": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: true,
										Optional: false,
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

									"namespace": {
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

			"spec": {
				Description:         "ApplicationSpec represents desired application state. Contains link to repository with application definition and additional parameters link definition revision.",
				MarkdownDescription: "ApplicationSpec represents desired application state. Contains link to repository with application definition and additional parameters link definition revision.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"destination": {
						Description:         "Destination is a reference to the target Kubernetes server and namespace",
						MarkdownDescription: "Destination is a reference to the target Kubernetes server and namespace",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "Name is an alternate way of specifying the target cluster by its symbolic name",
								MarkdownDescription: "Name is an alternate way of specifying the target cluster by its symbolic name",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"namespace": {
								Description:         "Namespace specifies the target namespace for the application's resources. The namespace will only be set for namespace-scoped resources that have not set a value for .metadata.namespace",
								MarkdownDescription: "Namespace specifies the target namespace for the application's resources. The namespace will only be set for namespace-scoped resources that have not set a value for .metadata.namespace",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"server": {
								Description:         "Server specifies the URL of the target cluster and must be set to the Kubernetes control plane API",
								MarkdownDescription: "Server specifies the URL of the target cluster and must be set to the Kubernetes control plane API",

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
						Description:         "IgnoreDifferences is a list of resources and their fields which should be ignored during comparison",
						MarkdownDescription: "IgnoreDifferences is a list of resources and their fields which should be ignored during comparison",

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
								Description:         "ManagedFieldsManagers is a list of trusted managers. Fields mutated by those managers will take precedence over the desired state defined in the SCM and won't be displayed in diffs",
								MarkdownDescription: "ManagedFieldsManagers is a list of trusted managers. Fields mutated by those managers will take precedence over the desired state defined in the SCM and won't be displayed in diffs",

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
						Description:         "Info contains a list of information (URLs, email addresses, and plain text) that relates to the application",
						MarkdownDescription: "Info contains a list of information (URLs, email addresses, and plain text) that relates to the application",

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
						Description:         "Project is a reference to the project this application belongs to. The empty string means that application belongs to the 'default' project.",
						MarkdownDescription: "Project is a reference to the project this application belongs to. The empty string means that application belongs to the 'default' project.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"revision_history_limit": {
						Description:         "RevisionHistoryLimit limits the number of items kept in the application's revision history, which is used for informational purposes as well as for rollbacks to previous versions. This should only be changed in exceptional circumstances. Setting to zero will store no history. This will reduce storage used. Increasing will increase the space used to store the history, so we do not recommend increasing it. Default is 10.",
						MarkdownDescription: "RevisionHistoryLimit limits the number of items kept in the application's revision history, which is used for informational purposes as well as for rollbacks to previous versions. This should only be changed in exceptional circumstances. Setting to zero will store no history. This will reduce storage used. Increasing will increase the space used to store the history, so we do not recommend increasing it. Default is 10.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"source": {
						Description:         "Source is a reference to the location of the application's manifests or chart",
						MarkdownDescription: "Source is a reference to the location of the application's manifests or chart",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"target_revision": {
								Description:         "TargetRevision defines the revision of the source to sync the application to. In case of Git, this can be commit, tag, or branch. If omitted, will equal to HEAD. In case of Helm, this is a semver tag for the Chart's version.",
								MarkdownDescription: "TargetRevision defines the revision of the source to sync the application to. In case of Git, this can be commit, tag, or branch. If omitted, will equal to HEAD. In case of Helm, this is a semver tag for the Chart's version.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"chart": {
								Description:         "Chart is a Helm chart name, and must be specified for applications sourced from a Helm repo.",
								MarkdownDescription: "Chart is a Helm chart name, and must be specified for applications sourced from a Helm repo.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"directory": {
								Description:         "Directory holds path/directory specific options",
								MarkdownDescription: "Directory holds path/directory specific options",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"exclude": {
										Description:         "Exclude contains a glob pattern to match paths against that should be explicitly excluded from being used during manifest generation",
										MarkdownDescription: "Exclude contains a glob pattern to match paths against that should be explicitly excluded from being used during manifest generation",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"include": {
										Description:         "Include contains a glob pattern to match paths against that should be explicitly included during manifest generation",
										MarkdownDescription: "Include contains a glob pattern to match paths against that should be explicitly included during manifest generation",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"jsonnet": {
										Description:         "Jsonnet holds options specific to Jsonnet",
										MarkdownDescription: "Jsonnet holds options specific to Jsonnet",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"ext_vars": {
												Description:         "ExtVars is a list of Jsonnet External Variables",
												MarkdownDescription: "ExtVars is a list of Jsonnet External Variables",

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
												Description:         "Additional library search dirs",
												MarkdownDescription: "Additional library search dirs",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"tlas": {
												Description:         "TLAS is a list of Jsonnet Top-level Arguments",
												MarkdownDescription: "TLAS is a list of Jsonnet Top-level Arguments",

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
										Description:         "Recurse specifies whether to scan a directory recursively for manifests",
										MarkdownDescription: "Recurse specifies whether to scan a directory recursively for manifests",

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
								Description:         "Helm holds helm specific options",
								MarkdownDescription: "Helm holds helm specific options",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"values": {
										Description:         "Values specifies Helm values to be passed to helm template, typically defined as a block",
										MarkdownDescription: "Values specifies Helm values to be passed to helm template, typically defined as a block",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"version": {
										Description:         "Version is the Helm version to use for templating ('3')",
										MarkdownDescription: "Version is the Helm version to use for templating ('3')",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"parameters": {
										Description:         "Parameters is a list of Helm parameters which are passed to the helm template command upon manifest generation",
										MarkdownDescription: "Parameters is a list of Helm parameters which are passed to the helm template command upon manifest generation",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"force_string": {
												Description:         "ForceString determines whether to tell Helm to interpret booleans and numbers as strings",
												MarkdownDescription: "ForceString determines whether to tell Helm to interpret booleans and numbers as strings",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "Name is the name of the Helm parameter",
												MarkdownDescription: "Name is the name of the Helm parameter",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value": {
												Description:         "Value is the value for the Helm parameter",
												MarkdownDescription: "Value is the value for the Helm parameter",

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
										Description:         "PassCredentials pass credentials to all domains (Helm's --pass-credentials)",
										MarkdownDescription: "PassCredentials pass credentials to all domains (Helm's --pass-credentials)",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"release_name": {
										Description:         "ReleaseName is the Helm release name to use. If omitted it will use the application name",
										MarkdownDescription: "ReleaseName is the Helm release name to use. If omitted it will use the application name",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"skip_crds": {
										Description:         "SkipCrds skips custom resource definition installation step (Helm's --skip-crds)",
										MarkdownDescription: "SkipCrds skips custom resource definition installation step (Helm's --skip-crds)",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"value_files": {
										Description:         "ValuesFiles is a list of Helm value files to use when generating a template",
										MarkdownDescription: "ValuesFiles is a list of Helm value files to use when generating a template",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"file_parameters": {
										Description:         "FileParameters are file parameters to the helm template",
										MarkdownDescription: "FileParameters are file parameters to the helm template",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "Name is the name of the Helm parameter",
												MarkdownDescription: "Name is the name of the Helm parameter",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"path": {
												Description:         "Path is the path to the file containing the values for the Helm parameter",
												MarkdownDescription: "Path is the path to the file containing the values for the Helm parameter",

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
										Description:         "IgnoreMissingValueFiles prevents helm template from failing when valueFiles do not exist locally by not appending them to helm template --values",
										MarkdownDescription: "IgnoreMissingValueFiles prevents helm template from failing when valueFiles do not exist locally by not appending them to helm template --values",

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

							"kustomize": {
								Description:         "Kustomize holds kustomize specific options",
								MarkdownDescription: "Kustomize holds kustomize specific options",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name_prefix": {
										Description:         "NamePrefix is a prefix appended to resources for Kustomize apps",
										MarkdownDescription: "NamePrefix is a prefix appended to resources for Kustomize apps",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name_suffix": {
										Description:         "NameSuffix is a suffix appended to resources for Kustomize apps",
										MarkdownDescription: "NameSuffix is a suffix appended to resources for Kustomize apps",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"version": {
										Description:         "Version controls which version of Kustomize to use for rendering manifests",
										MarkdownDescription: "Version controls which version of Kustomize to use for rendering manifests",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"common_annotations": {
										Description:         "CommonAnnotations is a list of additional annotations to add to rendered manifests",
										MarkdownDescription: "CommonAnnotations is a list of additional annotations to add to rendered manifests",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"common_labels": {
										Description:         "CommonLabels is a list of additional labels to add to rendered manifests",
										MarkdownDescription: "CommonLabels is a list of additional labels to add to rendered manifests",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"force_common_annotations": {
										Description:         "ForceCommonAnnotations specifies whether to force applying common annotations to resources for Kustomize apps",
										MarkdownDescription: "ForceCommonAnnotations specifies whether to force applying common annotations to resources for Kustomize apps",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"force_common_labels": {
										Description:         "ForceCommonLabels specifies whether to force applying common labels to resources for Kustomize apps",
										MarkdownDescription: "ForceCommonLabels specifies whether to force applying common labels to resources for Kustomize apps",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"images": {
										Description:         "Images is a list of Kustomize image override specifications",
										MarkdownDescription: "Images is a list of Kustomize image override specifications",

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

							"path": {
								Description:         "Path is a directory path within the Git repository, and is only valid for applications sourced from Git.",
								MarkdownDescription: "Path is a directory path within the Git repository, and is only valid for applications sourced from Git.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"plugin": {
								Description:         "ConfigManagementPlugin holds config management plugin specific options",
								MarkdownDescription: "ConfigManagementPlugin holds config management plugin specific options",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"env": {
										Description:         "Env is a list of environment variable entries",
										MarkdownDescription: "Env is a list of environment variable entries",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "Name is the name of the variable, usually expressed in uppercase",
												MarkdownDescription: "Name is the name of the variable, usually expressed in uppercase",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"value": {
												Description:         "Value is the value of the variable",
												MarkdownDescription: "Value is the value of the variable",

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
								Description:         "RepoURL is the URL to the repository (Git or Helm) that contains the application manifests",
								MarkdownDescription: "RepoURL is the URL to the repository (Git or Helm) that contains the application manifests",

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

					"sync_policy": {
						Description:         "SyncPolicy controls when and how a sync will be performed",
						MarkdownDescription: "SyncPolicy controls when and how a sync will be performed",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"automated": {
								Description:         "Automated will keep an application synced to the target revision",
								MarkdownDescription: "Automated will keep an application synced to the target revision",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"prune": {
										Description:         "Prune specifies whether to delete resources from the cluster that are not found in the sources anymore as part of automated sync (default: false)",
										MarkdownDescription: "Prune specifies whether to delete resources from the cluster that are not found in the sources anymore as part of automated sync (default: false)",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"self_heal": {
										Description:         "SelfHeal specifes whether to revert resources back to their desired state upon modification in the cluster (default: false)",
										MarkdownDescription: "SelfHeal specifes whether to revert resources back to their desired state upon modification in the cluster (default: false)",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"allow_empty": {
										Description:         "AllowEmpty allows apps have zero live resources (default: false)",
										MarkdownDescription: "AllowEmpty allows apps have zero live resources (default: false)",

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
								Description:         "Retry controls failed sync retry behavior",
								MarkdownDescription: "Retry controls failed sync retry behavior",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"backoff": {
										Description:         "Backoff controls how to backoff on subsequent retries of failed syncs",
										MarkdownDescription: "Backoff controls how to backoff on subsequent retries of failed syncs",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"duration": {
												Description:         "Duration is the amount to back off. Default unit is seconds, but could also be a duration (e.g. '2m', '1h')",
												MarkdownDescription: "Duration is the amount to back off. Default unit is seconds, but could also be a duration (e.g. '2m', '1h')",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"factor": {
												Description:         "Factor is a factor to multiply the base duration after each failed retry",
												MarkdownDescription: "Factor is a factor to multiply the base duration after each failed retry",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"max_duration": {
												Description:         "MaxDuration is the maximum amount of time allowed for the backoff strategy",
												MarkdownDescription: "MaxDuration is the maximum amount of time allowed for the backoff strategy",

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
										Description:         "Limit is the maximum number of attempts for retrying a failed sync. If set to 0, no retries will be performed.",
										MarkdownDescription: "Limit is the maximum number of attempts for retrying a failed sync. If set to 0, no retries will be performed.",

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
								Description:         "Options allow you to specify whole app sync-options",
								MarkdownDescription: "Options allow you to specify whole app sync-options",

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
		},
	}, nil
}

func (r *ArgoprojIoApplicationV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_argoproj_io_application_v1alpha1")

	var state ArgoprojIoApplicationV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ArgoprojIoApplicationV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("argoproj.io/v1alpha1")
	goModel.Kind = utilities.Ptr("Application")

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

func (r *ArgoprojIoApplicationV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_argoproj_io_application_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *ArgoprojIoApplicationV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_argoproj_io_application_v1alpha1")

	var state ArgoprojIoApplicationV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ArgoprojIoApplicationV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("argoproj.io/v1alpha1")
	goModel.Kind = utilities.Ptr("Application")

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

func (r *ArgoprojIoApplicationV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_argoproj_io_application_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
