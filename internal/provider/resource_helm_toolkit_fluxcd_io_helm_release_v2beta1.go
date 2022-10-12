/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

	"regexp"

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

type HelmToolkitFluxcdIoHelmReleaseV2Beta1Resource struct{}

var (
	_ resource.Resource = (*HelmToolkitFluxcdIoHelmReleaseV2Beta1Resource)(nil)
)

type HelmToolkitFluxcdIoHelmReleaseV2Beta1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type HelmToolkitFluxcdIoHelmReleaseV2Beta1GoModel struct {
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
		Chart *struct {
			Spec *struct {
				Chart *string `tfsdk:"chart" yaml:"chart,omitempty"`

				Interval *string `tfsdk:"interval" yaml:"interval,omitempty"`

				ReconcileStrategy *string `tfsdk:"reconcile_strategy" yaml:"reconcileStrategy,omitempty"`

				SourceRef *struct {
					ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

					Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
				} `tfsdk:"source_ref" yaml:"sourceRef,omitempty"`

				ValuesFile *string `tfsdk:"values_file" yaml:"valuesFile,omitempty"`

				ValuesFiles *[]string `tfsdk:"values_files" yaml:"valuesFiles,omitempty"`

				Version *string `tfsdk:"version" yaml:"version,omitempty"`
			} `tfsdk:"spec" yaml:"spec,omitempty"`
		} `tfsdk:"chart" yaml:"chart,omitempty"`

		DependsOn *[]struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
		} `tfsdk:"depends_on" yaml:"dependsOn,omitempty"`

		Install *struct {
			Crds *string `tfsdk:"crds" yaml:"crds,omitempty"`

			CreateNamespace *bool `tfsdk:"create_namespace" yaml:"createNamespace,omitempty"`

			DisableHooks *bool `tfsdk:"disable_hooks" yaml:"disableHooks,omitempty"`

			DisableOpenAPIValidation *bool `tfsdk:"disable_open_api_validation" yaml:"disableOpenAPIValidation,omitempty"`

			DisableWait *bool `tfsdk:"disable_wait" yaml:"disableWait,omitempty"`

			DisableWaitForJobs *bool `tfsdk:"disable_wait_for_jobs" yaml:"disableWaitForJobs,omitempty"`

			Remediation *struct {
				IgnoreTestFailures *bool `tfsdk:"ignore_test_failures" yaml:"ignoreTestFailures,omitempty"`

				RemediateLastFailure *bool `tfsdk:"remediate_last_failure" yaml:"remediateLastFailure,omitempty"`

				Retries *int64 `tfsdk:"retries" yaml:"retries,omitempty"`
			} `tfsdk:"remediation" yaml:"remediation,omitempty"`

			Replace *bool `tfsdk:"replace" yaml:"replace,omitempty"`

			SkipCRDs *bool `tfsdk:"skip_cr_ds" yaml:"skipCRDs,omitempty"`

			Timeout *string `tfsdk:"timeout" yaml:"timeout,omitempty"`
		} `tfsdk:"install" yaml:"install,omitempty"`

		Interval *string `tfsdk:"interval" yaml:"interval,omitempty"`

		KubeConfig *struct {
			SecretRef *struct {
				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`
		} `tfsdk:"kube_config" yaml:"kubeConfig,omitempty"`

		MaxHistory *int64 `tfsdk:"max_history" yaml:"maxHistory,omitempty"`

		PostRenderers *[]struct {
			Kustomize *struct {
				Images *[]struct {
					Digest *string `tfsdk:"digest" yaml:"digest,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					NewName *string `tfsdk:"new_name" yaml:"newName,omitempty"`

					NewTag *string `tfsdk:"new_tag" yaml:"newTag,omitempty"`
				} `tfsdk:"images" yaml:"images,omitempty"`

				Patches *[]struct {
					Patch *string `tfsdk:"patch" yaml:"patch,omitempty"`

					Target *struct {
						AnnotationSelector *string `tfsdk:"annotation_selector" yaml:"annotationSelector,omitempty"`

						Group *string `tfsdk:"group" yaml:"group,omitempty"`

						Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

						LabelSelector *string `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

						Version *string `tfsdk:"version" yaml:"version,omitempty"`
					} `tfsdk:"target" yaml:"target,omitempty"`
				} `tfsdk:"patches" yaml:"patches,omitempty"`

				PatchesJson6902 *[]struct {
					Patch *[]struct {
						From *string `tfsdk:"from" yaml:"from,omitempty"`

						Op *string `tfsdk:"op" yaml:"op,omitempty"`

						Path *string `tfsdk:"path" yaml:"path,omitempty"`

						Value utilities.Dynamic `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"patch" yaml:"patch,omitempty"`

					Target *struct {
						AnnotationSelector *string `tfsdk:"annotation_selector" yaml:"annotationSelector,omitempty"`

						Group *string `tfsdk:"group" yaml:"group,omitempty"`

						Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

						LabelSelector *string `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

						Version *string `tfsdk:"version" yaml:"version,omitempty"`
					} `tfsdk:"target" yaml:"target,omitempty"`
				} `tfsdk:"patches_json6902" yaml:"patchesJson6902,omitempty"`

				PatchesStrategicMerge *[]string `tfsdk:"patches_strategic_merge" yaml:"patchesStrategicMerge,omitempty"`
			} `tfsdk:"kustomize" yaml:"kustomize,omitempty"`
		} `tfsdk:"post_renderers" yaml:"postRenderers,omitempty"`

		ReleaseName *string `tfsdk:"release_name" yaml:"releaseName,omitempty"`

		Rollback *struct {
			CleanupOnFail *bool `tfsdk:"cleanup_on_fail" yaml:"cleanupOnFail,omitempty"`

			DisableHooks *bool `tfsdk:"disable_hooks" yaml:"disableHooks,omitempty"`

			DisableWait *bool `tfsdk:"disable_wait" yaml:"disableWait,omitempty"`

			DisableWaitForJobs *bool `tfsdk:"disable_wait_for_jobs" yaml:"disableWaitForJobs,omitempty"`

			Force *bool `tfsdk:"force" yaml:"force,omitempty"`

			Recreate *bool `tfsdk:"recreate" yaml:"recreate,omitempty"`

			Timeout *string `tfsdk:"timeout" yaml:"timeout,omitempty"`
		} `tfsdk:"rollback" yaml:"rollback,omitempty"`

		ServiceAccountName *string `tfsdk:"service_account_name" yaml:"serviceAccountName,omitempty"`

		StorageNamespace *string `tfsdk:"storage_namespace" yaml:"storageNamespace,omitempty"`

		Suspend *bool `tfsdk:"suspend" yaml:"suspend,omitempty"`

		TargetNamespace *string `tfsdk:"target_namespace" yaml:"targetNamespace,omitempty"`

		Test *struct {
			Enable *bool `tfsdk:"enable" yaml:"enable,omitempty"`

			IgnoreFailures *bool `tfsdk:"ignore_failures" yaml:"ignoreFailures,omitempty"`

			Timeout *string `tfsdk:"timeout" yaml:"timeout,omitempty"`
		} `tfsdk:"test" yaml:"test,omitempty"`

		Timeout *string `tfsdk:"timeout" yaml:"timeout,omitempty"`

		Uninstall *struct {
			DisableHooks *bool `tfsdk:"disable_hooks" yaml:"disableHooks,omitempty"`

			DisableWait *bool `tfsdk:"disable_wait" yaml:"disableWait,omitempty"`

			KeepHistory *bool `tfsdk:"keep_history" yaml:"keepHistory,omitempty"`

			Timeout *string `tfsdk:"timeout" yaml:"timeout,omitempty"`
		} `tfsdk:"uninstall" yaml:"uninstall,omitempty"`

		Upgrade *struct {
			CleanupOnFail *bool `tfsdk:"cleanup_on_fail" yaml:"cleanupOnFail,omitempty"`

			Crds *string `tfsdk:"crds" yaml:"crds,omitempty"`

			DisableHooks *bool `tfsdk:"disable_hooks" yaml:"disableHooks,omitempty"`

			DisableOpenAPIValidation *bool `tfsdk:"disable_open_api_validation" yaml:"disableOpenAPIValidation,omitempty"`

			DisableWait *bool `tfsdk:"disable_wait" yaml:"disableWait,omitempty"`

			DisableWaitForJobs *bool `tfsdk:"disable_wait_for_jobs" yaml:"disableWaitForJobs,omitempty"`

			Force *bool `tfsdk:"force" yaml:"force,omitempty"`

			PreserveValues *bool `tfsdk:"preserve_values" yaml:"preserveValues,omitempty"`

			Remediation *struct {
				IgnoreTestFailures *bool `tfsdk:"ignore_test_failures" yaml:"ignoreTestFailures,omitempty"`

				RemediateLastFailure *bool `tfsdk:"remediate_last_failure" yaml:"remediateLastFailure,omitempty"`

				Retries *int64 `tfsdk:"retries" yaml:"retries,omitempty"`

				Strategy *string `tfsdk:"strategy" yaml:"strategy,omitempty"`
			} `tfsdk:"remediation" yaml:"remediation,omitempty"`

			Timeout *string `tfsdk:"timeout" yaml:"timeout,omitempty"`
		} `tfsdk:"upgrade" yaml:"upgrade,omitempty"`

		Values utilities.Dynamic `tfsdk:"values" yaml:"values,omitempty"`

		ValuesFrom *[]struct {
			Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`

			TargetPath *string `tfsdk:"target_path" yaml:"targetPath,omitempty"`

			ValuesKey *string `tfsdk:"values_key" yaml:"valuesKey,omitempty"`
		} `tfsdk:"values_from" yaml:"valuesFrom,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewHelmToolkitFluxcdIoHelmReleaseV2Beta1Resource() resource.Resource {
	return &HelmToolkitFluxcdIoHelmReleaseV2Beta1Resource{}
}

func (r *HelmToolkitFluxcdIoHelmReleaseV2Beta1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_helm_toolkit_fluxcd_io_helm_release_v2beta1"
}

func (r *HelmToolkitFluxcdIoHelmReleaseV2Beta1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "HelmRelease is the Schema for the helmreleases API",
		MarkdownDescription: "HelmRelease is the Schema for the helmreleases API",
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
				Description:         "HelmReleaseSpec defines the desired state of a Helm release.",
				MarkdownDescription: "HelmReleaseSpec defines the desired state of a Helm release.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"chart": {
						Description:         "Chart defines the template of the v1beta2.HelmChart that should be created for this HelmRelease.",
						MarkdownDescription: "Chart defines the template of the v1beta2.HelmChart that should be created for this HelmRelease.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"spec": {
								Description:         "Spec holds the template for the v1beta2.HelmChartSpec for this HelmRelease.",
								MarkdownDescription: "Spec holds the template for the v1beta2.HelmChartSpec for this HelmRelease.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"chart": {
										Description:         "The name or path the Helm chart is available at in the SourceRef.",
										MarkdownDescription: "The name or path the Helm chart is available at in the SourceRef.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"interval": {
										Description:         "Interval at which to check the v1beta2.Source for updates. Defaults to 'HelmReleaseSpec.Interval'.",
										MarkdownDescription: "Interval at which to check the v1beta2.Source for updates. Defaults to 'HelmReleaseSpec.Interval'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ms|s|m|h))+$`), ""),
										},
									},

									"reconcile_strategy": {
										Description:         "Determines what enables the creation of a new artifact. Valid values are ('ChartVersion', 'Revision'). See the documentation of the values for an explanation on their behavior. Defaults to ChartVersion when omitted.",
										MarkdownDescription: "Determines what enables the creation of a new artifact. Valid values are ('ChartVersion', 'Revision'). See the documentation of the values for an explanation on their behavior. Defaults to ChartVersion when omitted.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("ChartVersion", "Revision"),
										},
									},

									"source_ref": {
										Description:         "The name and namespace of the v1beta2.Source the chart is available at.",
										MarkdownDescription: "The name and namespace of the v1beta2.Source the chart is available at.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"api_version": {
												Description:         "APIVersion of the referent.",
												MarkdownDescription: "APIVersion of the referent.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"kind": {
												Description:         "Kind of the referent.",
												MarkdownDescription: "Kind of the referent.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("HelmRepository", "GitRepository", "Bucket"),
												},
											},

											"name": {
												Description:         "Name of the referent.",
												MarkdownDescription: "Name of the referent.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.LengthAtLeast(1),

													stringvalidator.LengthAtMost(253),
												},
											},

											"namespace": {
												Description:         "Namespace of the referent.",
												MarkdownDescription: "Namespace of the referent.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.LengthAtLeast(1),

													stringvalidator.LengthAtMost(63),
												},
											},
										}),

										Required: true,
										Optional: false,
										Computed: false,
									},

									"values_file": {
										Description:         "Alternative values file to use as the default chart values, expected to be a relative path in the SourceRef. Deprecated in favor of ValuesFiles, for backwards compatibility the file defined here is merged before the ValuesFiles items. Ignored when omitted.",
										MarkdownDescription: "Alternative values file to use as the default chart values, expected to be a relative path in the SourceRef. Deprecated in favor of ValuesFiles, for backwards compatibility the file defined here is merged before the ValuesFiles items. Ignored when omitted.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"values_files": {
										Description:         "Alternative list of values files to use as the chart values (values.yaml is not included by default), expected to be a relative path in the SourceRef. Values files are merged in the order of this list with the last file overriding the first. Ignored when omitted.",
										MarkdownDescription: "Alternative list of values files to use as the chart values (values.yaml is not included by default), expected to be a relative path in the SourceRef. Values files are merged in the order of this list with the last file overriding the first. Ignored when omitted.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"version": {
										Description:         "Version semver expression, ignored for charts from v1beta2.GitRepository and v1beta2.Bucket sources. Defaults to latest when omitted.",
										MarkdownDescription: "Version semver expression, ignored for charts from v1beta2.GitRepository and v1beta2.Bucket sources. Defaults to latest when omitted.",

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
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"depends_on": {
						Description:         "DependsOn may contain a meta.NamespacedObjectReference slice with references to HelmRelease resources that must be ready before this HelmRelease can be reconciled.",
						MarkdownDescription: "DependsOn may contain a meta.NamespacedObjectReference slice with references to HelmRelease resources that must be ready before this HelmRelease can be reconciled.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "Name of the referent.",
								MarkdownDescription: "Name of the referent.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"namespace": {
								Description:         "Namespace of the referent, when not specified it acts as LocalObjectReference.",
								MarkdownDescription: "Namespace of the referent, when not specified it acts as LocalObjectReference.",

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

					"install": {
						Description:         "Install holds the configuration for Helm install actions for this HelmRelease.",
						MarkdownDescription: "Install holds the configuration for Helm install actions for this HelmRelease.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"crds": {
								Description:         "CRDs upgrade CRDs from the Helm Chart's crds directory according to the CRD upgrade policy provided here. Valid values are 'Skip', 'Create' or 'CreateReplace'. Default is 'Create' and if omitted CRDs are installed but not updated.  Skip: do neither install nor replace (update) any CRDs.  Create: new CRDs are created, existing CRDs are neither updated nor deleted.  CreateReplace: new CRDs are created, existing CRDs are updated (replaced) but not deleted.  By default, CRDs are applied (installed) during Helm install action. With this option users can opt-in to CRD replace existing CRDs on Helm install actions, which is not (yet) natively supported by Helm. https://helm.sh/docs/chart_best_practices/custom_resource_definitions.",
								MarkdownDescription: "CRDs upgrade CRDs from the Helm Chart's crds directory according to the CRD upgrade policy provided here. Valid values are 'Skip', 'Create' or 'CreateReplace'. Default is 'Create' and if omitted CRDs are installed but not updated.  Skip: do neither install nor replace (update) any CRDs.  Create: new CRDs are created, existing CRDs are neither updated nor deleted.  CreateReplace: new CRDs are created, existing CRDs are updated (replaced) but not deleted.  By default, CRDs are applied (installed) during Helm install action. With this option users can opt-in to CRD replace existing CRDs on Helm install actions, which is not (yet) natively supported by Helm. https://helm.sh/docs/chart_best_practices/custom_resource_definitions.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("Skip", "Create", "CreateReplace"),
								},
							},

							"create_namespace": {
								Description:         "CreateNamespace tells the Helm install action to create the HelmReleaseSpec.TargetNamespace if it does not exist yet. On uninstall, the namespace will not be garbage collected.",
								MarkdownDescription: "CreateNamespace tells the Helm install action to create the HelmReleaseSpec.TargetNamespace if it does not exist yet. On uninstall, the namespace will not be garbage collected.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"disable_hooks": {
								Description:         "DisableHooks prevents hooks from running during the Helm install action.",
								MarkdownDescription: "DisableHooks prevents hooks from running during the Helm install action.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"disable_open_api_validation": {
								Description:         "DisableOpenAPIValidation prevents the Helm install action from validating rendered templates against the Kubernetes OpenAPI Schema.",
								MarkdownDescription: "DisableOpenAPIValidation prevents the Helm install action from validating rendered templates against the Kubernetes OpenAPI Schema.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"disable_wait": {
								Description:         "DisableWait disables the waiting for resources to be ready after a Helm install has been performed.",
								MarkdownDescription: "DisableWait disables the waiting for resources to be ready after a Helm install has been performed.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"disable_wait_for_jobs": {
								Description:         "DisableWaitForJobs disables waiting for jobs to complete after a Helm install has been performed.",
								MarkdownDescription: "DisableWaitForJobs disables waiting for jobs to complete after a Helm install has been performed.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"remediation": {
								Description:         "Remediation holds the remediation configuration for when the Helm install action for the HelmRelease fails. The default is to not perform any action.",
								MarkdownDescription: "Remediation holds the remediation configuration for when the Helm install action for the HelmRelease fails. The default is to not perform any action.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"ignore_test_failures": {
										Description:         "IgnoreTestFailures tells the controller to skip remediation when the Helm tests are run after an install action but fail. Defaults to 'Test.IgnoreFailures'.",
										MarkdownDescription: "IgnoreTestFailures tells the controller to skip remediation when the Helm tests are run after an install action but fail. Defaults to 'Test.IgnoreFailures'.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"remediate_last_failure": {
										Description:         "RemediateLastFailure tells the controller to remediate the last failure, when no retries remain. Defaults to 'false'.",
										MarkdownDescription: "RemediateLastFailure tells the controller to remediate the last failure, when no retries remain. Defaults to 'false'.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"retries": {
										Description:         "Retries is the number of retries that should be attempted on failures before bailing. Remediation, using an uninstall, is performed between each attempt. Defaults to '0', a negative integer equals to unlimited retries.",
										MarkdownDescription: "Retries is the number of retries that should be attempted on failures before bailing. Remediation, using an uninstall, is performed between each attempt. Defaults to '0', a negative integer equals to unlimited retries.",

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

							"replace": {
								Description:         "Replace tells the Helm install action to re-use the 'ReleaseName', but only if that name is a deleted release which remains in the history.",
								MarkdownDescription: "Replace tells the Helm install action to re-use the 'ReleaseName', but only if that name is a deleted release which remains in the history.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"skip_cr_ds": {
								Description:         "SkipCRDs tells the Helm install action to not install any CRDs. By default, CRDs are installed if not already present.  Deprecated use CRD policy ('crds') attribute with value 'Skip' instead.",
								MarkdownDescription: "SkipCRDs tells the Helm install action to not install any CRDs. By default, CRDs are installed if not already present.  Deprecated use CRD policy ('crds') attribute with value 'Skip' instead.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"timeout": {
								Description:         "Timeout is the time to wait for any individual Kubernetes operation (like Jobs for hooks) during the performance of a Helm install action. Defaults to 'HelmReleaseSpec.Timeout'.",
								MarkdownDescription: "Timeout is the time to wait for any individual Kubernetes operation (like Jobs for hooks) during the performance of a Helm install action. Defaults to 'HelmReleaseSpec.Timeout'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ms|s|m))+$`), ""),
								},
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"interval": {
						Description:         "Interval at which to reconcile the Helm release.",
						MarkdownDescription: "Interval at which to reconcile the Helm release.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ms|s|m|h))+$`), ""),
						},
					},

					"kube_config": {
						Description:         "KubeConfig for reconciling the HelmRelease on a remote cluster. When used in combination with HelmReleaseSpec.ServiceAccountName, forces the controller to act on behalf of that Service Account at the target cluster. If the --default-service-account flag is set, its value will be used as a controller level fallback for when HelmReleaseSpec.ServiceAccountName is empty.",
						MarkdownDescription: "KubeConfig for reconciling the HelmRelease on a remote cluster. When used in combination with HelmReleaseSpec.ServiceAccountName, forces the controller to act on behalf of that Service Account at the target cluster. If the --default-service-account flag is set, its value will be used as a controller level fallback for when HelmReleaseSpec.ServiceAccountName is empty.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"secret_ref": {
								Description:         "SecretRef holds the name to a secret that contains a key with the kubeconfig file as the value. If no key is specified the key will default to 'value'. The secret must be in the same namespace as the HelmRelease. It is recommended that the kubeconfig is self-contained, and the secret is regularly updated if credentials such as a cloud-access-token expire. Cloud specific 'cmd-path' auth helpers will not function without adding binaries and credentials to the Pod that is responsible for reconciling the HelmRelease.",
								MarkdownDescription: "SecretRef holds the name to a secret that contains a key with the kubeconfig file as the value. If no key is specified the key will default to 'value'. The secret must be in the same namespace as the HelmRelease. It is recommended that the kubeconfig is self-contained, and the secret is regularly updated if credentials such as a cloud-access-token expire. Cloud specific 'cmd-path' auth helpers will not function without adding binaries and credentials to the Pod that is responsible for reconciling the HelmRelease.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"key": {
										Description:         "Key in the Secret, when not specified an implementation-specific default key is used.",
										MarkdownDescription: "Key in the Secret, when not specified an implementation-specific default key is used.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "Name of the Secret.",
										MarkdownDescription: "Name of the Secret.",

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

					"max_history": {
						Description:         "MaxHistory is the number of revisions saved by Helm for this HelmRelease. Use '0' for an unlimited number of revisions; defaults to '10'.",
						MarkdownDescription: "MaxHistory is the number of revisions saved by Helm for this HelmRelease. Use '0' for an unlimited number of revisions; defaults to '10'.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"post_renderers": {
						Description:         "PostRenderers holds an array of Helm PostRenderers, which will be applied in order of their definition.",
						MarkdownDescription: "PostRenderers holds an array of Helm PostRenderers, which will be applied in order of their definition.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"kustomize": {
								Description:         "Kustomization to apply as PostRenderer.",
								MarkdownDescription: "Kustomization to apply as PostRenderer.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"images": {
										Description:         "Images is a list of (image name, new name, new tag or digest) for changing image names, tags or digests. This can also be achieved with a patch, but this operator is simpler to specify.",
										MarkdownDescription: "Images is a list of (image name, new name, new tag or digest) for changing image names, tags or digests. This can also be achieved with a patch, but this operator is simpler to specify.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"digest": {
												Description:         "Digest is the value used to replace the original image tag. If digest is present NewTag value is ignored.",
												MarkdownDescription: "Digest is the value used to replace the original image tag. If digest is present NewTag value is ignored.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "Name is a tag-less image name.",
												MarkdownDescription: "Name is a tag-less image name.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"new_name": {
												Description:         "NewName is the value used to replace the original name.",
												MarkdownDescription: "NewName is the value used to replace the original name.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"new_tag": {
												Description:         "NewTag is the value used to replace the original tag.",
												MarkdownDescription: "NewTag is the value used to replace the original tag.",

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

									"patches": {
										Description:         "Strategic merge and JSON patches, defined as inline YAML objects, capable of targeting objects based on kind, label and annotation selectors.",
										MarkdownDescription: "Strategic merge and JSON patches, defined as inline YAML objects, capable of targeting objects based on kind, label and annotation selectors.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"patch": {
												Description:         "Patch contains an inline StrategicMerge patch or an inline JSON6902 patch with an array of operation objects.",
												MarkdownDescription: "Patch contains an inline StrategicMerge patch or an inline JSON6902 patch with an array of operation objects.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"target": {
												Description:         "Target points to the resources that the patch document should be applied to.",
												MarkdownDescription: "Target points to the resources that the patch document should be applied to.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotation_selector": {
														Description:         "AnnotationSelector is a string that follows the label selection expression https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#api It matches with the resource annotations.",
														MarkdownDescription: "AnnotationSelector is a string that follows the label selection expression https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#api It matches with the resource annotations.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"group": {
														Description:         "Group is the API group to select resources from. Together with Version and Kind it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",
														MarkdownDescription: "Group is the API group to select resources from. Together with Version and Kind it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"kind": {
														Description:         "Kind of the API Group to select resources from. Together with Group and Version it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",
														MarkdownDescription: "Kind of the API Group to select resources from. Together with Group and Version it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"label_selector": {
														Description:         "LabelSelector is a string that follows the label selection expression https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#api It matches with the resource labels.",
														MarkdownDescription: "LabelSelector is a string that follows the label selection expression https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#api It matches with the resource labels.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": {
														Description:         "Name to match resources with.",
														MarkdownDescription: "Name to match resources with.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespace": {
														Description:         "Namespace to select resources from.",
														MarkdownDescription: "Namespace to select resources from.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"version": {
														Description:         "Version of the API Group to select resources from. Together with Group and Kind it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",
														MarkdownDescription: "Version of the API Group to select resources from. Together with Group and Kind it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",

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

									"patches_json6902": {
										Description:         "JSON 6902 patches, defined as inline YAML objects.",
										MarkdownDescription: "JSON 6902 patches, defined as inline YAML objects.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"patch": {
												Description:         "Patch contains the JSON6902 patch document with an array of operation objects.",
												MarkdownDescription: "Patch contains the JSON6902 patch document with an array of operation objects.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"from": {
														Description:         "From contains a JSON-pointer value that references a location within the target document where the operation is performed. The meaning of the value depends on the value of Op, and is NOT taken into account by all operations.",
														MarkdownDescription: "From contains a JSON-pointer value that references a location within the target document where the operation is performed. The meaning of the value depends on the value of Op, and is NOT taken into account by all operations.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"op": {
														Description:         "Op indicates the operation to perform. Its value MUST be one of 'add', 'remove', 'replace', 'move', 'copy', or 'test'. https://datatracker.ietf.org/doc/html/rfc6902#section-4",
														MarkdownDescription: "Op indicates the operation to perform. Its value MUST be one of 'add', 'remove', 'replace', 'move', 'copy', or 'test'. https://datatracker.ietf.org/doc/html/rfc6902#section-4",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("test", "remove", "add", "replace", "move", "copy"),
														},
													},

													"path": {
														Description:         "Path contains the JSON-pointer value that references a location within the target document where the operation is performed. The meaning of the value depends on the value of Op.",
														MarkdownDescription: "Path contains the JSON-pointer value that references a location within the target document where the operation is performed. The meaning of the value depends on the value of Op.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"value": {
														Description:         "Value contains a valid JSON structure. The meaning of the value depends on the value of Op, and is NOT taken into account by all operations.",
														MarkdownDescription: "Value contains a valid JSON structure. The meaning of the value depends on the value of Op, and is NOT taken into account by all operations.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},

											"target": {
												Description:         "Target points to the resources that the patch document should be applied to.",
												MarkdownDescription: "Target points to the resources that the patch document should be applied to.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotation_selector": {
														Description:         "AnnotationSelector is a string that follows the label selection expression https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#api It matches with the resource annotations.",
														MarkdownDescription: "AnnotationSelector is a string that follows the label selection expression https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#api It matches with the resource annotations.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"group": {
														Description:         "Group is the API group to select resources from. Together with Version and Kind it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",
														MarkdownDescription: "Group is the API group to select resources from. Together with Version and Kind it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"kind": {
														Description:         "Kind of the API Group to select resources from. Together with Group and Version it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",
														MarkdownDescription: "Kind of the API Group to select resources from. Together with Group and Version it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"label_selector": {
														Description:         "LabelSelector is a string that follows the label selection expression https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#api It matches with the resource labels.",
														MarkdownDescription: "LabelSelector is a string that follows the label selection expression https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#api It matches with the resource labels.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": {
														Description:         "Name to match resources with.",
														MarkdownDescription: "Name to match resources with.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespace": {
														Description:         "Namespace to select resources from.",
														MarkdownDescription: "Namespace to select resources from.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"version": {
														Description:         "Version of the API Group to select resources from. Together with Group and Kind it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",
														MarkdownDescription: "Version of the API Group to select resources from. Together with Group and Kind it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",

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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"patches_strategic_merge": {
										Description:         "Strategic merge patches, defined as inline YAML objects.",
										MarkdownDescription: "Strategic merge patches, defined as inline YAML objects.",

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

						Required: false,
						Optional: true,
						Computed: false,
					},

					"release_name": {
						Description:         "ReleaseName used for the Helm release. Defaults to a composition of '[TargetNamespace-]Name'.",
						MarkdownDescription: "ReleaseName used for the Helm release. Defaults to a composition of '[TargetNamespace-]Name'.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.LengthAtLeast(1),

							stringvalidator.LengthAtMost(53),
						},
					},

					"rollback": {
						Description:         "Rollback holds the configuration for Helm rollback actions for this HelmRelease.",
						MarkdownDescription: "Rollback holds the configuration for Helm rollback actions for this HelmRelease.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"cleanup_on_fail": {
								Description:         "CleanupOnFail allows deletion of new resources created during the Helm rollback action when it fails.",
								MarkdownDescription: "CleanupOnFail allows deletion of new resources created during the Helm rollback action when it fails.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"disable_hooks": {
								Description:         "DisableHooks prevents hooks from running during the Helm rollback action.",
								MarkdownDescription: "DisableHooks prevents hooks from running during the Helm rollback action.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"disable_wait": {
								Description:         "DisableWait disables the waiting for resources to be ready after a Helm rollback has been performed.",
								MarkdownDescription: "DisableWait disables the waiting for resources to be ready after a Helm rollback has been performed.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"disable_wait_for_jobs": {
								Description:         "DisableWaitForJobs disables waiting for jobs to complete after a Helm rollback has been performed.",
								MarkdownDescription: "DisableWaitForJobs disables waiting for jobs to complete after a Helm rollback has been performed.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"force": {
								Description:         "Force forces resource updates through a replacement strategy.",
								MarkdownDescription: "Force forces resource updates through a replacement strategy.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"recreate": {
								Description:         "Recreate performs pod restarts for the resource if applicable.",
								MarkdownDescription: "Recreate performs pod restarts for the resource if applicable.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"timeout": {
								Description:         "Timeout is the time to wait for any individual Kubernetes operation (like Jobs for hooks) during the performance of a Helm rollback action. Defaults to 'HelmReleaseSpec.Timeout'.",
								MarkdownDescription: "Timeout is the time to wait for any individual Kubernetes operation (like Jobs for hooks) during the performance of a Helm rollback action. Defaults to 'HelmReleaseSpec.Timeout'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ms|s|m))+$`), ""),
								},
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"service_account_name": {
						Description:         "The name of the Kubernetes service account to impersonate when reconciling this HelmRelease.",
						MarkdownDescription: "The name of the Kubernetes service account to impersonate when reconciling this HelmRelease.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"storage_namespace": {
						Description:         "StorageNamespace used for the Helm storage. Defaults to the namespace of the HelmRelease.",
						MarkdownDescription: "StorageNamespace used for the Helm storage. Defaults to the namespace of the HelmRelease.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.LengthAtLeast(1),

							stringvalidator.LengthAtMost(63),
						},
					},

					"suspend": {
						Description:         "Suspend tells the controller to suspend reconciliation for this HelmRelease, it does not apply to already started reconciliations. Defaults to false.",
						MarkdownDescription: "Suspend tells the controller to suspend reconciliation for this HelmRelease, it does not apply to already started reconciliations. Defaults to false.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"target_namespace": {
						Description:         "TargetNamespace to target when performing operations for the HelmRelease. Defaults to the namespace of the HelmRelease.",
						MarkdownDescription: "TargetNamespace to target when performing operations for the HelmRelease. Defaults to the namespace of the HelmRelease.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.LengthAtLeast(1),

							stringvalidator.LengthAtMost(63),
						},
					},

					"test": {
						Description:         "Test holds the configuration for Helm test actions for this HelmRelease.",
						MarkdownDescription: "Test holds the configuration for Helm test actions for this HelmRelease.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"enable": {
								Description:         "Enable enables Helm test actions for this HelmRelease after an Helm install or upgrade action has been performed.",
								MarkdownDescription: "Enable enables Helm test actions for this HelmRelease after an Helm install or upgrade action has been performed.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ignore_failures": {
								Description:         "IgnoreFailures tells the controller to skip remediation when the Helm tests are run but fail. Can be overwritten for tests run after install or upgrade actions in 'Install.IgnoreTestFailures' and 'Upgrade.IgnoreTestFailures'.",
								MarkdownDescription: "IgnoreFailures tells the controller to skip remediation when the Helm tests are run but fail. Can be overwritten for tests run after install or upgrade actions in 'Install.IgnoreTestFailures' and 'Upgrade.IgnoreTestFailures'.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"timeout": {
								Description:         "Timeout is the time to wait for any individual Kubernetes operation during the performance of a Helm test action. Defaults to 'HelmReleaseSpec.Timeout'.",
								MarkdownDescription: "Timeout is the time to wait for any individual Kubernetes operation during the performance of a Helm test action. Defaults to 'HelmReleaseSpec.Timeout'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ms|s|m))+$`), ""),
								},
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"timeout": {
						Description:         "Timeout is the time to wait for any individual Kubernetes operation (like Jobs for hooks) during the performance of a Helm action. Defaults to '5m0s'.",
						MarkdownDescription: "Timeout is the time to wait for any individual Kubernetes operation (like Jobs for hooks) during the performance of a Helm action. Defaults to '5m0s'.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ms|s|m))+$`), ""),
						},
					},

					"uninstall": {
						Description:         "Uninstall holds the configuration for Helm uninstall actions for this HelmRelease.",
						MarkdownDescription: "Uninstall holds the configuration for Helm uninstall actions for this HelmRelease.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"disable_hooks": {
								Description:         "DisableHooks prevents hooks from running during the Helm rollback action.",
								MarkdownDescription: "DisableHooks prevents hooks from running during the Helm rollback action.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"disable_wait": {
								Description:         "DisableWait disables waiting for all the resources to be deleted after a Helm uninstall is performed.",
								MarkdownDescription: "DisableWait disables waiting for all the resources to be deleted after a Helm uninstall is performed.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"keep_history": {
								Description:         "KeepHistory tells Helm to remove all associated resources and mark the release as deleted, but retain the release history.",
								MarkdownDescription: "KeepHistory tells Helm to remove all associated resources and mark the release as deleted, but retain the release history.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"timeout": {
								Description:         "Timeout is the time to wait for any individual Kubernetes operation (like Jobs for hooks) during the performance of a Helm uninstall action. Defaults to 'HelmReleaseSpec.Timeout'.",
								MarkdownDescription: "Timeout is the time to wait for any individual Kubernetes operation (like Jobs for hooks) during the performance of a Helm uninstall action. Defaults to 'HelmReleaseSpec.Timeout'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ms|s|m))+$`), ""),
								},
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"upgrade": {
						Description:         "Upgrade holds the configuration for Helm upgrade actions for this HelmRelease.",
						MarkdownDescription: "Upgrade holds the configuration for Helm upgrade actions for this HelmRelease.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"cleanup_on_fail": {
								Description:         "CleanupOnFail allows deletion of new resources created during the Helm upgrade action when it fails.",
								MarkdownDescription: "CleanupOnFail allows deletion of new resources created during the Helm upgrade action when it fails.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"crds": {
								Description:         "CRDs upgrade CRDs from the Helm Chart's crds directory according to the CRD upgrade policy provided here. Valid values are 'Skip', 'Create' or 'CreateReplace'. Default is 'Skip' and if omitted CRDs are neither installed nor upgraded.  Skip: do neither install nor replace (update) any CRDs.  Create: new CRDs are created, existing CRDs are neither updated nor deleted.  CreateReplace: new CRDs are created, existing CRDs are updated (replaced) but not deleted.  By default, CRDs are not applied during Helm upgrade action. With this option users can opt-in to CRD upgrade, which is not (yet) natively supported by Helm. https://helm.sh/docs/chart_best_practices/custom_resource_definitions.",
								MarkdownDescription: "CRDs upgrade CRDs from the Helm Chart's crds directory according to the CRD upgrade policy provided here. Valid values are 'Skip', 'Create' or 'CreateReplace'. Default is 'Skip' and if omitted CRDs are neither installed nor upgraded.  Skip: do neither install nor replace (update) any CRDs.  Create: new CRDs are created, existing CRDs are neither updated nor deleted.  CreateReplace: new CRDs are created, existing CRDs are updated (replaced) but not deleted.  By default, CRDs are not applied during Helm upgrade action. With this option users can opt-in to CRD upgrade, which is not (yet) natively supported by Helm. https://helm.sh/docs/chart_best_practices/custom_resource_definitions.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("Skip", "Create", "CreateReplace"),
								},
							},

							"disable_hooks": {
								Description:         "DisableHooks prevents hooks from running during the Helm upgrade action.",
								MarkdownDescription: "DisableHooks prevents hooks from running during the Helm upgrade action.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"disable_open_api_validation": {
								Description:         "DisableOpenAPIValidation prevents the Helm upgrade action from validating rendered templates against the Kubernetes OpenAPI Schema.",
								MarkdownDescription: "DisableOpenAPIValidation prevents the Helm upgrade action from validating rendered templates against the Kubernetes OpenAPI Schema.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"disable_wait": {
								Description:         "DisableWait disables the waiting for resources to be ready after a Helm upgrade has been performed.",
								MarkdownDescription: "DisableWait disables the waiting for resources to be ready after a Helm upgrade has been performed.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"disable_wait_for_jobs": {
								Description:         "DisableWaitForJobs disables waiting for jobs to complete after a Helm upgrade has been performed.",
								MarkdownDescription: "DisableWaitForJobs disables waiting for jobs to complete after a Helm upgrade has been performed.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"force": {
								Description:         "Force forces resource updates through a replacement strategy.",
								MarkdownDescription: "Force forces resource updates through a replacement strategy.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"preserve_values": {
								Description:         "PreserveValues will make Helm reuse the last release's values and merge in overrides from 'Values'. Setting this flag makes the HelmRelease non-declarative.",
								MarkdownDescription: "PreserveValues will make Helm reuse the last release's values and merge in overrides from 'Values'. Setting this flag makes the HelmRelease non-declarative.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"remediation": {
								Description:         "Remediation holds the remediation configuration for when the Helm upgrade action for the HelmRelease fails. The default is to not perform any action.",
								MarkdownDescription: "Remediation holds the remediation configuration for when the Helm upgrade action for the HelmRelease fails. The default is to not perform any action.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"ignore_test_failures": {
										Description:         "IgnoreTestFailures tells the controller to skip remediation when the Helm tests are run after an upgrade action but fail. Defaults to 'Test.IgnoreFailures'.",
										MarkdownDescription: "IgnoreTestFailures tells the controller to skip remediation when the Helm tests are run after an upgrade action but fail. Defaults to 'Test.IgnoreFailures'.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"remediate_last_failure": {
										Description:         "RemediateLastFailure tells the controller to remediate the last failure, when no retries remain. Defaults to 'false' unless 'Retries' is greater than 0.",
										MarkdownDescription: "RemediateLastFailure tells the controller to remediate the last failure, when no retries remain. Defaults to 'false' unless 'Retries' is greater than 0.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"retries": {
										Description:         "Retries is the number of retries that should be attempted on failures before bailing. Remediation, using 'Strategy', is performed between each attempt. Defaults to '0', a negative integer equals to unlimited retries.",
										MarkdownDescription: "Retries is the number of retries that should be attempted on failures before bailing. Remediation, using 'Strategy', is performed between each attempt. Defaults to '0', a negative integer equals to unlimited retries.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"strategy": {
										Description:         "Strategy to use for failure remediation. Defaults to 'rollback'.",
										MarkdownDescription: "Strategy to use for failure remediation. Defaults to 'rollback'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("rollback", "uninstall"),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"timeout": {
								Description:         "Timeout is the time to wait for any individual Kubernetes operation (like Jobs for hooks) during the performance of a Helm upgrade action. Defaults to 'HelmReleaseSpec.Timeout'.",
								MarkdownDescription: "Timeout is the time to wait for any individual Kubernetes operation (like Jobs for hooks) during the performance of a Helm upgrade action. Defaults to 'HelmReleaseSpec.Timeout'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ms|s|m))+$`), ""),
								},
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"values": {
						Description:         "Values holds the values for this Helm release.",
						MarkdownDescription: "Values holds the values for this Helm release.",

						Type: utilities.DynamicType{},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"values_from": {
						Description:         "ValuesFrom holds references to resources containing Helm values for this HelmRelease, and information about how they should be merged.",
						MarkdownDescription: "ValuesFrom holds references to resources containing Helm values for this HelmRelease, and information about how they should be merged.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"kind": {
								Description:         "Kind of the values referent, valid values are ('Secret', 'ConfigMap').",
								MarkdownDescription: "Kind of the values referent, valid values are ('Secret', 'ConfigMap').",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("Secret", "ConfigMap"),
								},
							},

							"name": {
								Description:         "Name of the values referent. Should reside in the same namespace as the referring resource.",
								MarkdownDescription: "Name of the values referent. Should reside in the same namespace as the referring resource.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.LengthAtLeast(1),

									stringvalidator.LengthAtMost(253),
								},
							},

							"optional": {
								Description:         "Optional marks this ValuesReference as optional. When set, a not found error for the values reference is ignored, but any ValuesKey, TargetPath or transient error will still result in a reconciliation failure.",
								MarkdownDescription: "Optional marks this ValuesReference as optional. When set, a not found error for the values reference is ignored, but any ValuesKey, TargetPath or transient error will still result in a reconciliation failure.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"target_path": {
								Description:         "TargetPath is the YAML dot notation path the value should be merged at. When set, the ValuesKey is expected to be a single flat value. Defaults to 'None', which results in the values getting merged at the root.",
								MarkdownDescription: "TargetPath is the YAML dot notation path the value should be merged at. When set, the ValuesKey is expected to be a single flat value. Defaults to 'None', which results in the values getting merged at the root.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.LengthAtMost(250),

									stringvalidator.RegexMatches(regexp.MustCompile(`^([a-zA-Z0-9_\-.\\\/]|\[[0-9]{1,5}\])+$`), ""),
								},
							},

							"values_key": {
								Description:         "ValuesKey is the data key where the values.yaml or a specific value can be found at. Defaults to 'values.yaml'. When set, must be a valid Data Key, consisting of alphanumeric characters, '-', '_' or '.'.",
								MarkdownDescription: "ValuesKey is the data key where the values.yaml or a specific value can be found at. Defaults to 'values.yaml'. When set, must be a valid Data Key, consisting of alphanumeric characters, '-', '_' or '.'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.LengthAtMost(253),

									stringvalidator.RegexMatches(regexp.MustCompile(`^[\-._a-zA-Z0-9]+$`), ""),
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
		},
	}, nil
}

func (r *HelmToolkitFluxcdIoHelmReleaseV2Beta1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_helm_toolkit_fluxcd_io_helm_release_v2beta1")

	var state HelmToolkitFluxcdIoHelmReleaseV2Beta1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel HelmToolkitFluxcdIoHelmReleaseV2Beta1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("helm.toolkit.fluxcd.io/v2beta1")
	goModel.Kind = utilities.Ptr("HelmRelease")

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

func (r *HelmToolkitFluxcdIoHelmReleaseV2Beta1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_helm_toolkit_fluxcd_io_helm_release_v2beta1")
	// NO-OP: All data is already in Terraform state
}

func (r *HelmToolkitFluxcdIoHelmReleaseV2Beta1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_helm_toolkit_fluxcd_io_helm_release_v2beta1")

	var state HelmToolkitFluxcdIoHelmReleaseV2Beta1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel HelmToolkitFluxcdIoHelmReleaseV2Beta1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("helm.toolkit.fluxcd.io/v2beta1")
	goModel.Kind = utilities.Ptr("HelmRelease")

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

func (r *HelmToolkitFluxcdIoHelmReleaseV2Beta1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_helm_toolkit_fluxcd_io_helm_release_v2beta1")
	// NO-OP: Terraform removes the state automatically for us
}
