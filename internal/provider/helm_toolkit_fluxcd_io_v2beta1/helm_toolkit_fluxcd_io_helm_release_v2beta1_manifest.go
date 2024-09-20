/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package helm_toolkit_fluxcd_io_v2beta1

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
	_ datasource.DataSource = &HelmToolkitFluxcdIoHelmReleaseV2Beta1Manifest{}
)

func NewHelmToolkitFluxcdIoHelmReleaseV2Beta1Manifest() datasource.DataSource {
	return &HelmToolkitFluxcdIoHelmReleaseV2Beta1Manifest{}
}

type HelmToolkitFluxcdIoHelmReleaseV2Beta1Manifest struct{}

type HelmToolkitFluxcdIoHelmReleaseV2Beta1ManifestData struct {
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
		Chart *struct {
			Metadata *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"metadata" json:"metadata,omitempty"`
			Spec *struct {
				Chart             *string `tfsdk:"chart" json:"chart,omitempty"`
				Interval          *string `tfsdk:"interval" json:"interval,omitempty"`
				ReconcileStrategy *string `tfsdk:"reconcile_strategy" json:"reconcileStrategy,omitempty"`
				SourceRef         *struct {
					ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
					Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
					Name       *string `tfsdk:"name" json:"name,omitempty"`
					Namespace  *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"source_ref" json:"sourceRef,omitempty"`
				ValuesFile  *string   `tfsdk:"values_file" json:"valuesFile,omitempty"`
				ValuesFiles *[]string `tfsdk:"values_files" json:"valuesFiles,omitempty"`
				Verify      *struct {
					Provider  *string `tfsdk:"provider" json:"provider,omitempty"`
					SecretRef *struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
				} `tfsdk:"verify" json:"verify,omitempty"`
				Version *string `tfsdk:"version" json:"version,omitempty"`
			} `tfsdk:"spec" json:"spec,omitempty"`
		} `tfsdk:"chart" json:"chart,omitempty"`
		ChartRef *struct {
			ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
			Name       *string `tfsdk:"name" json:"name,omitempty"`
			Namespace  *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"chart_ref" json:"chartRef,omitempty"`
		DependsOn *[]struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"depends_on" json:"dependsOn,omitempty"`
		DriftDetection *struct {
			Ignore *[]struct {
				Paths  *[]string `tfsdk:"paths" json:"paths,omitempty"`
				Target *struct {
					AnnotationSelector *string `tfsdk:"annotation_selector" json:"annotationSelector,omitempty"`
					Group              *string `tfsdk:"group" json:"group,omitempty"`
					Kind               *string `tfsdk:"kind" json:"kind,omitempty"`
					LabelSelector      *string `tfsdk:"label_selector" json:"labelSelector,omitempty"`
					Name               *string `tfsdk:"name" json:"name,omitempty"`
					Namespace          *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Version            *string `tfsdk:"version" json:"version,omitempty"`
				} `tfsdk:"target" json:"target,omitempty"`
			} `tfsdk:"ignore" json:"ignore,omitempty"`
			Mode *string `tfsdk:"mode" json:"mode,omitempty"`
		} `tfsdk:"drift_detection" json:"driftDetection,omitempty"`
		Install *struct {
			Crds                     *string `tfsdk:"crds" json:"crds,omitempty"`
			CreateNamespace          *bool   `tfsdk:"create_namespace" json:"createNamespace,omitempty"`
			DisableHooks             *bool   `tfsdk:"disable_hooks" json:"disableHooks,omitempty"`
			DisableOpenAPIValidation *bool   `tfsdk:"disable_open_api_validation" json:"disableOpenAPIValidation,omitempty"`
			DisableWait              *bool   `tfsdk:"disable_wait" json:"disableWait,omitempty"`
			DisableWaitForJobs       *bool   `tfsdk:"disable_wait_for_jobs" json:"disableWaitForJobs,omitempty"`
			Remediation              *struct {
				IgnoreTestFailures   *bool  `tfsdk:"ignore_test_failures" json:"ignoreTestFailures,omitempty"`
				RemediateLastFailure *bool  `tfsdk:"remediate_last_failure" json:"remediateLastFailure,omitempty"`
				Retries              *int64 `tfsdk:"retries" json:"retries,omitempty"`
			} `tfsdk:"remediation" json:"remediation,omitempty"`
			Replace  *bool   `tfsdk:"replace" json:"replace,omitempty"`
			SkipCRDs *bool   `tfsdk:"skip_cr_ds" json:"skipCRDs,omitempty"`
			Timeout  *string `tfsdk:"timeout" json:"timeout,omitempty"`
		} `tfsdk:"install" json:"install,omitempty"`
		Interval   *string `tfsdk:"interval" json:"interval,omitempty"`
		KubeConfig *struct {
			SecretRef *struct {
				Key  *string `tfsdk:"key" json:"key,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
		} `tfsdk:"kube_config" json:"kubeConfig,omitempty"`
		MaxHistory       *int64 `tfsdk:"max_history" json:"maxHistory,omitempty"`
		PersistentClient *bool  `tfsdk:"persistent_client" json:"persistentClient,omitempty"`
		PostRenderers    *[]struct {
			Kustomize *struct {
				Images *[]struct {
					Digest  *string `tfsdk:"digest" json:"digest,omitempty"`
					Name    *string `tfsdk:"name" json:"name,omitempty"`
					NewName *string `tfsdk:"new_name" json:"newName,omitempty"`
					NewTag  *string `tfsdk:"new_tag" json:"newTag,omitempty"`
				} `tfsdk:"images" json:"images,omitempty"`
				Patches *[]struct {
					Patch  *string `tfsdk:"patch" json:"patch,omitempty"`
					Target *struct {
						AnnotationSelector *string `tfsdk:"annotation_selector" json:"annotationSelector,omitempty"`
						Group              *string `tfsdk:"group" json:"group,omitempty"`
						Kind               *string `tfsdk:"kind" json:"kind,omitempty"`
						LabelSelector      *string `tfsdk:"label_selector" json:"labelSelector,omitempty"`
						Name               *string `tfsdk:"name" json:"name,omitempty"`
						Namespace          *string `tfsdk:"namespace" json:"namespace,omitempty"`
						Version            *string `tfsdk:"version" json:"version,omitempty"`
					} `tfsdk:"target" json:"target,omitempty"`
				} `tfsdk:"patches" json:"patches,omitempty"`
				PatchesJson6902 *[]struct {
					Patch *[]struct {
						From  *string            `tfsdk:"from" json:"from,omitempty"`
						Op    *string            `tfsdk:"op" json:"op,omitempty"`
						Path  *string            `tfsdk:"path" json:"path,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"patch" json:"patch,omitempty"`
					Target *struct {
						AnnotationSelector *string `tfsdk:"annotation_selector" json:"annotationSelector,omitempty"`
						Group              *string `tfsdk:"group" json:"group,omitempty"`
						Kind               *string `tfsdk:"kind" json:"kind,omitempty"`
						LabelSelector      *string `tfsdk:"label_selector" json:"labelSelector,omitempty"`
						Name               *string `tfsdk:"name" json:"name,omitempty"`
						Namespace          *string `tfsdk:"namespace" json:"namespace,omitempty"`
						Version            *string `tfsdk:"version" json:"version,omitempty"`
					} `tfsdk:"target" json:"target,omitempty"`
				} `tfsdk:"patches_json6902" json:"patchesJson6902,omitempty"`
				PatchesStrategicMerge *[]string `tfsdk:"patches_strategic_merge" json:"patchesStrategicMerge,omitempty"`
			} `tfsdk:"kustomize" json:"kustomize,omitempty"`
		} `tfsdk:"post_renderers" json:"postRenderers,omitempty"`
		ReleaseName *string `tfsdk:"release_name" json:"releaseName,omitempty"`
		Rollback    *struct {
			CleanupOnFail      *bool   `tfsdk:"cleanup_on_fail" json:"cleanupOnFail,omitempty"`
			DisableHooks       *bool   `tfsdk:"disable_hooks" json:"disableHooks,omitempty"`
			DisableWait        *bool   `tfsdk:"disable_wait" json:"disableWait,omitempty"`
			DisableWaitForJobs *bool   `tfsdk:"disable_wait_for_jobs" json:"disableWaitForJobs,omitempty"`
			Force              *bool   `tfsdk:"force" json:"force,omitempty"`
			Recreate           *bool   `tfsdk:"recreate" json:"recreate,omitempty"`
			Timeout            *string `tfsdk:"timeout" json:"timeout,omitempty"`
		} `tfsdk:"rollback" json:"rollback,omitempty"`
		ServiceAccountName *string `tfsdk:"service_account_name" json:"serviceAccountName,omitempty"`
		StorageNamespace   *string `tfsdk:"storage_namespace" json:"storageNamespace,omitempty"`
		Suspend            *bool   `tfsdk:"suspend" json:"suspend,omitempty"`
		TargetNamespace    *string `tfsdk:"target_namespace" json:"targetNamespace,omitempty"`
		Test               *struct {
			Enable         *bool   `tfsdk:"enable" json:"enable,omitempty"`
			IgnoreFailures *bool   `tfsdk:"ignore_failures" json:"ignoreFailures,omitempty"`
			Timeout        *string `tfsdk:"timeout" json:"timeout,omitempty"`
		} `tfsdk:"test" json:"test,omitempty"`
		Timeout   *string `tfsdk:"timeout" json:"timeout,omitempty"`
		Uninstall *struct {
			DeletionPropagation *string `tfsdk:"deletion_propagation" json:"deletionPropagation,omitempty"`
			DisableHooks        *bool   `tfsdk:"disable_hooks" json:"disableHooks,omitempty"`
			DisableWait         *bool   `tfsdk:"disable_wait" json:"disableWait,omitempty"`
			KeepHistory         *bool   `tfsdk:"keep_history" json:"keepHistory,omitempty"`
			Timeout             *string `tfsdk:"timeout" json:"timeout,omitempty"`
		} `tfsdk:"uninstall" json:"uninstall,omitempty"`
		Upgrade *struct {
			CleanupOnFail            *bool   `tfsdk:"cleanup_on_fail" json:"cleanupOnFail,omitempty"`
			Crds                     *string `tfsdk:"crds" json:"crds,omitempty"`
			DisableHooks             *bool   `tfsdk:"disable_hooks" json:"disableHooks,omitempty"`
			DisableOpenAPIValidation *bool   `tfsdk:"disable_open_api_validation" json:"disableOpenAPIValidation,omitempty"`
			DisableWait              *bool   `tfsdk:"disable_wait" json:"disableWait,omitempty"`
			DisableWaitForJobs       *bool   `tfsdk:"disable_wait_for_jobs" json:"disableWaitForJobs,omitempty"`
			Force                    *bool   `tfsdk:"force" json:"force,omitempty"`
			PreserveValues           *bool   `tfsdk:"preserve_values" json:"preserveValues,omitempty"`
			Remediation              *struct {
				IgnoreTestFailures   *bool   `tfsdk:"ignore_test_failures" json:"ignoreTestFailures,omitempty"`
				RemediateLastFailure *bool   `tfsdk:"remediate_last_failure" json:"remediateLastFailure,omitempty"`
				Retries              *int64  `tfsdk:"retries" json:"retries,omitempty"`
				Strategy             *string `tfsdk:"strategy" json:"strategy,omitempty"`
			} `tfsdk:"remediation" json:"remediation,omitempty"`
			Timeout *string `tfsdk:"timeout" json:"timeout,omitempty"`
		} `tfsdk:"upgrade" json:"upgrade,omitempty"`
		Values     *map[string]string `tfsdk:"values" json:"values,omitempty"`
		ValuesFrom *[]struct {
			Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
			Name       *string `tfsdk:"name" json:"name,omitempty"`
			Optional   *bool   `tfsdk:"optional" json:"optional,omitempty"`
			TargetPath *string `tfsdk:"target_path" json:"targetPath,omitempty"`
			ValuesKey  *string `tfsdk:"values_key" json:"valuesKey,omitempty"`
		} `tfsdk:"values_from" json:"valuesFrom,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *HelmToolkitFluxcdIoHelmReleaseV2Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_helm_toolkit_fluxcd_io_helm_release_v2beta1_manifest"
}

func (r *HelmToolkitFluxcdIoHelmReleaseV2Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "HelmRelease is the Schema for the helmreleases API",
		MarkdownDescription: "HelmRelease is the Schema for the helmreleases API",
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
				Description:         "HelmReleaseSpec defines the desired state of a Helm release.",
				MarkdownDescription: "HelmReleaseSpec defines the desired state of a Helm release.",
				Attributes: map[string]schema.Attribute{
					"chart": schema.SingleNestedAttribute{
						Description:         "Chart defines the template of the v1beta2.HelmChart that should be created for this HelmRelease.",
						MarkdownDescription: "Chart defines the template of the v1beta2.HelmChart that should be created for this HelmRelease.",
						Attributes: map[string]schema.Attribute{
							"metadata": schema.SingleNestedAttribute{
								Description:         "ObjectMeta holds the template for metadata like labels and annotations.",
								MarkdownDescription: "ObjectMeta holds the template for metadata like labels and annotations.",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/",
										MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.MapAttribute{
										Description:         "Map of string keys and values that can be used to organize and categorize (scope and select) objects. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/",
										MarkdownDescription: "Map of string keys and values that can be used to organize and categorize (scope and select) objects. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/",
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

							"spec": schema.SingleNestedAttribute{
								Description:         "Spec holds the template for the v1beta2.HelmChartSpec for this HelmRelease.",
								MarkdownDescription: "Spec holds the template for the v1beta2.HelmChartSpec for this HelmRelease.",
								Attributes: map[string]schema.Attribute{
									"chart": schema.StringAttribute{
										Description:         "The name or path the Helm chart is available at in the SourceRef.",
										MarkdownDescription: "The name or path the Helm chart is available at in the SourceRef.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"interval": schema.StringAttribute{
										Description:         "Interval at which to check the v1beta2.Source for updates. Defaults to 'HelmReleaseSpec.Interval'.",
										MarkdownDescription: "Interval at which to check the v1beta2.Source for updates. Defaults to 'HelmReleaseSpec.Interval'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ms|s|m|h))+$`), ""),
										},
									},

									"reconcile_strategy": schema.StringAttribute{
										Description:         "Determines what enables the creation of a new artifact. Valid values are ('ChartVersion', 'Revision'). See the documentation of the values for an explanation on their behavior. Defaults to ChartVersion when omitted.",
										MarkdownDescription: "Determines what enables the creation of a new artifact. Valid values are ('ChartVersion', 'Revision'). See the documentation of the values for an explanation on their behavior. Defaults to ChartVersion when omitted.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("ChartVersion", "Revision"),
										},
									},

									"source_ref": schema.SingleNestedAttribute{
										Description:         "The name and namespace of the v1beta2.Source the chart is available at.",
										MarkdownDescription: "The name and namespace of the v1beta2.Source the chart is available at.",
										Attributes: map[string]schema.Attribute{
											"api_version": schema.StringAttribute{
												Description:         "APIVersion of the referent.",
												MarkdownDescription: "APIVersion of the referent.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"kind": schema.StringAttribute{
												Description:         "Kind of the referent.",
												MarkdownDescription: "Kind of the referent.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("HelmRepository", "GitRepository", "Bucket"),
												},
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent.",
												MarkdownDescription: "Name of the referent.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
													stringvalidator.LengthAtMost(253),
												},
											},

											"namespace": schema.StringAttribute{
												Description:         "Namespace of the referent.",
												MarkdownDescription: "Namespace of the referent.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
													stringvalidator.LengthAtMost(63),
												},
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"values_file": schema.StringAttribute{
										Description:         "Alternative values file to use as the default chart values, expected to be a relative path in the SourceRef. Deprecated in favor of ValuesFiles, for backwards compatibility the file defined here is merged before the ValuesFiles items. Ignored when omitted.",
										MarkdownDescription: "Alternative values file to use as the default chart values, expected to be a relative path in the SourceRef. Deprecated in favor of ValuesFiles, for backwards compatibility the file defined here is merged before the ValuesFiles items. Ignored when omitted.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"values_files": schema.ListAttribute{
										Description:         "Alternative list of values files to use as the chart values (values.yaml is not included by default), expected to be a relative path in the SourceRef. Values files are merged in the order of this list with the last file overriding the first. Ignored when omitted.",
										MarkdownDescription: "Alternative list of values files to use as the chart values (values.yaml is not included by default), expected to be a relative path in the SourceRef. Values files are merged in the order of this list with the last file overriding the first. Ignored when omitted.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"verify": schema.SingleNestedAttribute{
										Description:         "Verify contains the secret name containing the trusted public keys used to verify the signature and specifies which provider to use to check whether OCI image is authentic. This field is only supported for OCI sources. Chart dependencies, which are not bundled in the umbrella chart artifact, are not verified.",
										MarkdownDescription: "Verify contains the secret name containing the trusted public keys used to verify the signature and specifies which provider to use to check whether OCI image is authentic. This field is only supported for OCI sources. Chart dependencies, which are not bundled in the umbrella chart artifact, are not verified.",
										Attributes: map[string]schema.Attribute{
											"provider": schema.StringAttribute{
												Description:         "Provider specifies the technology used to sign the OCI Helm chart.",
												MarkdownDescription: "Provider specifies the technology used to sign the OCI Helm chart.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("cosign"),
												},
											},

											"secret_ref": schema.SingleNestedAttribute{
												Description:         "SecretRef specifies the Kubernetes Secret containing the trusted public keys.",
												MarkdownDescription: "SecretRef specifies the Kubernetes Secret containing the trusted public keys.",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "Name of the referent.",
														MarkdownDescription: "Name of the referent.",
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

									"version": schema.StringAttribute{
										Description:         "Version semver expression, ignored for charts from v1beta2.GitRepository and v1beta2.Bucket sources. Defaults to latest when omitted.",
										MarkdownDescription: "Version semver expression, ignored for charts from v1beta2.GitRepository and v1beta2.Bucket sources. Defaults to latest when omitted.",
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
						Required: true,
						Optional: false,
						Computed: false,
					},

					"chart_ref": schema.SingleNestedAttribute{
						Description:         "ChartRef holds a reference to a source controller resource containing the Helm chart artifact. Note: this field is provisional to the v2 API, and not actively used by v2beta1 HelmReleases.",
						MarkdownDescription: "ChartRef holds a reference to a source controller resource containing the Helm chart artifact. Note: this field is provisional to the v2 API, and not actively used by v2beta1 HelmReleases.",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "APIVersion of the referent.",
								MarkdownDescription: "APIVersion of the referent.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind of the referent.",
								MarkdownDescription: "Kind of the referent.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("OCIRepository", "HelmChart"),
								},
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referent.",
								MarkdownDescription: "Name of the referent.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(253),
								},
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace of the referent, defaults to the namespace of the Kubernetes resource object that contains the reference.",
								MarkdownDescription: "Namespace of the referent, defaults to the namespace of the Kubernetes resource object that contains the reference.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(63),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"depends_on": schema.ListNestedAttribute{
						Description:         "DependsOn may contain a meta.NamespacedObjectReference slice with references to HelmRelease resources that must be ready before this HelmRelease can be reconciled.",
						MarkdownDescription: "DependsOn may contain a meta.NamespacedObjectReference slice with references to HelmRelease resources that must be ready before this HelmRelease can be reconciled.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name of the referent.",
									MarkdownDescription: "Name of the referent.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace of the referent, when not specified it acts as LocalObjectReference.",
									MarkdownDescription: "Namespace of the referent, when not specified it acts as LocalObjectReference.",
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

					"drift_detection": schema.SingleNestedAttribute{
						Description:         "DriftDetection holds the configuration for detecting and handling differences between the manifest in the Helm storage and the resources currently existing in the cluster. Note: this field is provisional to the v2beta2 API, and not actively used by v2beta1 HelmReleases.",
						MarkdownDescription: "DriftDetection holds the configuration for detecting and handling differences between the manifest in the Helm storage and the resources currently existing in the cluster. Note: this field is provisional to the v2beta2 API, and not actively used by v2beta1 HelmReleases.",
						Attributes: map[string]schema.Attribute{
							"ignore": schema.ListNestedAttribute{
								Description:         "Ignore contains a list of rules for specifying which changes to ignore during diffing.",
								MarkdownDescription: "Ignore contains a list of rules for specifying which changes to ignore during diffing.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"paths": schema.ListAttribute{
											Description:         "Paths is a list of JSON Pointer (RFC 6901) paths to be excluded from consideration in a Kubernetes object.",
											MarkdownDescription: "Paths is a list of JSON Pointer (RFC 6901) paths to be excluded from consideration in a Kubernetes object.",
											ElementType:         types.StringType,
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"target": schema.SingleNestedAttribute{
											Description:         "Target is a selector for specifying Kubernetes objects to which this rule applies. If Target is not set, the Paths will be ignored for all Kubernetes objects within the manifest of the Helm release.",
											MarkdownDescription: "Target is a selector for specifying Kubernetes objects to which this rule applies. If Target is not set, the Paths will be ignored for all Kubernetes objects within the manifest of the Helm release.",
											Attributes: map[string]schema.Attribute{
												"annotation_selector": schema.StringAttribute{
													Description:         "AnnotationSelector is a string that follows the label selection expression https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#api It matches with the resource annotations.",
													MarkdownDescription: "AnnotationSelector is a string that follows the label selection expression https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#api It matches with the resource annotations.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"group": schema.StringAttribute{
													Description:         "Group is the API group to select resources from. Together with Version and Kind it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",
													MarkdownDescription: "Group is the API group to select resources from. Together with Version and Kind it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"kind": schema.StringAttribute{
													Description:         "Kind of the API Group to select resources from. Together with Group and Version it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",
													MarkdownDescription: "Kind of the API Group to select resources from. Together with Group and Version it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"label_selector": schema.StringAttribute{
													Description:         "LabelSelector is a string that follows the label selection expression https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#api It matches with the resource labels.",
													MarkdownDescription: "LabelSelector is a string that follows the label selection expression https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#api It matches with the resource labels.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name to match resources with.",
													MarkdownDescription: "Name to match resources with.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespace": schema.StringAttribute{
													Description:         "Namespace to select resources from.",
													MarkdownDescription: "Namespace to select resources from.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"version": schema.StringAttribute{
													Description:         "Version of the API Group to select resources from. Together with Group and Kind it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",
													MarkdownDescription: "Version of the API Group to select resources from. Together with Group and Kind it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",
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

							"mode": schema.StringAttribute{
								Description:         "Mode defines how differences should be handled between the Helm manifest and the manifest currently applied to the cluster. If not explicitly set, it defaults to DiffModeDisabled.",
								MarkdownDescription: "Mode defines how differences should be handled between the Helm manifest and the manifest currently applied to the cluster. If not explicitly set, it defaults to DiffModeDisabled.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("enabled", "warn", "disabled"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"install": schema.SingleNestedAttribute{
						Description:         "Install holds the configuration for Helm install actions for this HelmRelease.",
						MarkdownDescription: "Install holds the configuration for Helm install actions for this HelmRelease.",
						Attributes: map[string]schema.Attribute{
							"crds": schema.StringAttribute{
								Description:         "CRDs upgrade CRDs from the Helm Chart's crds directory according to the CRD upgrade policy provided here. Valid values are 'Skip', 'Create' or 'CreateReplace'. Default is 'Create' and if omitted CRDs are installed but not updated. Skip: do neither install nor replace (update) any CRDs. Create: new CRDs are created, existing CRDs are neither updated nor deleted. CreateReplace: new CRDs are created, existing CRDs are updated (replaced) but not deleted. By default, CRDs are applied (installed) during Helm install action. With this option users can opt-in to CRD replace existing CRDs on Helm install actions, which is not (yet) natively supported by Helm. https://helm.sh/docs/chart_best_practices/custom_resource_definitions.",
								MarkdownDescription: "CRDs upgrade CRDs from the Helm Chart's crds directory according to the CRD upgrade policy provided here. Valid values are 'Skip', 'Create' or 'CreateReplace'. Default is 'Create' and if omitted CRDs are installed but not updated. Skip: do neither install nor replace (update) any CRDs. Create: new CRDs are created, existing CRDs are neither updated nor deleted. CreateReplace: new CRDs are created, existing CRDs are updated (replaced) but not deleted. By default, CRDs are applied (installed) during Helm install action. With this option users can opt-in to CRD replace existing CRDs on Helm install actions, which is not (yet) natively supported by Helm. https://helm.sh/docs/chart_best_practices/custom_resource_definitions.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Skip", "Create", "CreateReplace"),
								},
							},

							"create_namespace": schema.BoolAttribute{
								Description:         "CreateNamespace tells the Helm install action to create the HelmReleaseSpec.TargetNamespace if it does not exist yet. On uninstall, the namespace will not be garbage collected.",
								MarkdownDescription: "CreateNamespace tells the Helm install action to create the HelmReleaseSpec.TargetNamespace if it does not exist yet. On uninstall, the namespace will not be garbage collected.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"disable_hooks": schema.BoolAttribute{
								Description:         "DisableHooks prevents hooks from running during the Helm install action.",
								MarkdownDescription: "DisableHooks prevents hooks from running during the Helm install action.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"disable_open_api_validation": schema.BoolAttribute{
								Description:         "DisableOpenAPIValidation prevents the Helm install action from validating rendered templates against the Kubernetes OpenAPI Schema.",
								MarkdownDescription: "DisableOpenAPIValidation prevents the Helm install action from validating rendered templates against the Kubernetes OpenAPI Schema.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"disable_wait": schema.BoolAttribute{
								Description:         "DisableWait disables the waiting for resources to be ready after a Helm install has been performed.",
								MarkdownDescription: "DisableWait disables the waiting for resources to be ready after a Helm install has been performed.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"disable_wait_for_jobs": schema.BoolAttribute{
								Description:         "DisableWaitForJobs disables waiting for jobs to complete after a Helm install has been performed.",
								MarkdownDescription: "DisableWaitForJobs disables waiting for jobs to complete after a Helm install has been performed.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"remediation": schema.SingleNestedAttribute{
								Description:         "Remediation holds the remediation configuration for when the Helm install action for the HelmRelease fails. The default is to not perform any action.",
								MarkdownDescription: "Remediation holds the remediation configuration for when the Helm install action for the HelmRelease fails. The default is to not perform any action.",
								Attributes: map[string]schema.Attribute{
									"ignore_test_failures": schema.BoolAttribute{
										Description:         "IgnoreTestFailures tells the controller to skip remediation when the Helm tests are run after an install action but fail. Defaults to 'Test.IgnoreFailures'.",
										MarkdownDescription: "IgnoreTestFailures tells the controller to skip remediation when the Helm tests are run after an install action but fail. Defaults to 'Test.IgnoreFailures'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"remediate_last_failure": schema.BoolAttribute{
										Description:         "RemediateLastFailure tells the controller to remediate the last failure, when no retries remain. Defaults to 'false'.",
										MarkdownDescription: "RemediateLastFailure tells the controller to remediate the last failure, when no retries remain. Defaults to 'false'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"retries": schema.Int64Attribute{
										Description:         "Retries is the number of retries that should be attempted on failures before bailing. Remediation, using an uninstall, is performed between each attempt. Defaults to '0', a negative integer equals to unlimited retries.",
										MarkdownDescription: "Retries is the number of retries that should be attempted on failures before bailing. Remediation, using an uninstall, is performed between each attempt. Defaults to '0', a negative integer equals to unlimited retries.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"replace": schema.BoolAttribute{
								Description:         "Replace tells the Helm install action to re-use the 'ReleaseName', but only if that name is a deleted release which remains in the history.",
								MarkdownDescription: "Replace tells the Helm install action to re-use the 'ReleaseName', but only if that name is a deleted release which remains in the history.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"skip_cr_ds": schema.BoolAttribute{
								Description:         "SkipCRDs tells the Helm install action to not install any CRDs. By default, CRDs are installed if not already present. Deprecated use CRD policy ('crds') attribute with value 'Skip' instead.",
								MarkdownDescription: "SkipCRDs tells the Helm install action to not install any CRDs. By default, CRDs are installed if not already present. Deprecated use CRD policy ('crds') attribute with value 'Skip' instead.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"timeout": schema.StringAttribute{
								Description:         "Timeout is the time to wait for any individual Kubernetes operation (like Jobs for hooks) during the performance of a Helm install action. Defaults to 'HelmReleaseSpec.Timeout'.",
								MarkdownDescription: "Timeout is the time to wait for any individual Kubernetes operation (like Jobs for hooks) during the performance of a Helm install action. Defaults to 'HelmReleaseSpec.Timeout'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ms|s|m|h))+$`), ""),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"interval": schema.StringAttribute{
						Description:         "Interval at which to reconcile the Helm release. This interval is approximate and may be subject to jitter to ensure efficient use of resources.",
						MarkdownDescription: "Interval at which to reconcile the Helm release. This interval is approximate and may be subject to jitter to ensure efficient use of resources.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ms|s|m|h))+$`), ""),
						},
					},

					"kube_config": schema.SingleNestedAttribute{
						Description:         "KubeConfig for reconciling the HelmRelease on a remote cluster. When used in combination with HelmReleaseSpec.ServiceAccountName, forces the controller to act on behalf of that Service Account at the target cluster. If the --default-service-account flag is set, its value will be used as a controller level fallback for when HelmReleaseSpec.ServiceAccountName is empty.",
						MarkdownDescription: "KubeConfig for reconciling the HelmRelease on a remote cluster. When used in combination with HelmReleaseSpec.ServiceAccountName, forces the controller to act on behalf of that Service Account at the target cluster. If the --default-service-account flag is set, its value will be used as a controller level fallback for when HelmReleaseSpec.ServiceAccountName is empty.",
						Attributes: map[string]schema.Attribute{
							"secret_ref": schema.SingleNestedAttribute{
								Description:         "SecretRef holds the name of a secret that contains a key with the kubeconfig file as the value. If no key is set, the key will default to 'value'. It is recommended that the kubeconfig is self-contained, and the secret is regularly updated if credentials such as a cloud-access-token expire. Cloud specific 'cmd-path' auth helpers will not function without adding binaries and credentials to the Pod that is responsible for reconciling Kubernetes resources.",
								MarkdownDescription: "SecretRef holds the name of a secret that contains a key with the kubeconfig file as the value. If no key is set, the key will default to 'value'. It is recommended that the kubeconfig is self-contained, and the secret is regularly updated if credentials such as a cloud-access-token expire. Cloud specific 'cmd-path' auth helpers will not function without adding binaries and credentials to the Pod that is responsible for reconciling Kubernetes resources.",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "Key in the Secret, when not specified an implementation-specific default key is used.",
										MarkdownDescription: "Key in the Secret, when not specified an implementation-specific default key is used.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name of the Secret.",
										MarkdownDescription: "Name of the Secret.",
										Required:            true,
										Optional:            false,
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

					"max_history": schema.Int64Attribute{
						Description:         "MaxHistory is the number of revisions saved by Helm for this HelmRelease. Use '0' for an unlimited number of revisions; defaults to '10'.",
						MarkdownDescription: "MaxHistory is the number of revisions saved by Helm for this HelmRelease. Use '0' for an unlimited number of revisions; defaults to '10'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"persistent_client": schema.BoolAttribute{
						Description:         "PersistentClient tells the controller to use a persistent Kubernetes client for this release. When enabled, the client will be reused for the duration of the reconciliation, instead of being created and destroyed for each (step of a) Helm action. This can improve performance, but may cause issues with some Helm charts that for example do create Custom Resource Definitions during installation outside Helm's CRD lifecycle hooks, which are then not observed to be available by e.g. post-install hooks. If not set, it defaults to true.",
						MarkdownDescription: "PersistentClient tells the controller to use a persistent Kubernetes client for this release. When enabled, the client will be reused for the duration of the reconciliation, instead of being created and destroyed for each (step of a) Helm action. This can improve performance, but may cause issues with some Helm charts that for example do create Custom Resource Definitions during installation outside Helm's CRD lifecycle hooks, which are then not observed to be available by e.g. post-install hooks. If not set, it defaults to true.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"post_renderers": schema.ListNestedAttribute{
						Description:         "PostRenderers holds an array of Helm PostRenderers, which will be applied in order of their definition.",
						MarkdownDescription: "PostRenderers holds an array of Helm PostRenderers, which will be applied in order of their definition.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"kustomize": schema.SingleNestedAttribute{
									Description:         "Kustomization to apply as PostRenderer.",
									MarkdownDescription: "Kustomization to apply as PostRenderer.",
									Attributes: map[string]schema.Attribute{
										"images": schema.ListNestedAttribute{
											Description:         "Images is a list of (image name, new name, new tag or digest) for changing image names, tags or digests. This can also be achieved with a patch, but this operator is simpler to specify.",
											MarkdownDescription: "Images is a list of (image name, new name, new tag or digest) for changing image names, tags or digests. This can also be achieved with a patch, but this operator is simpler to specify.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"digest": schema.StringAttribute{
														Description:         "Digest is the value used to replace the original image tag. If digest is present NewTag value is ignored.",
														MarkdownDescription: "Digest is the value used to replace the original image tag. If digest is present NewTag value is ignored.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name is a tag-less image name.",
														MarkdownDescription: "Name is a tag-less image name.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"new_name": schema.StringAttribute{
														Description:         "NewName is the value used to replace the original name.",
														MarkdownDescription: "NewName is the value used to replace the original name.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"new_tag": schema.StringAttribute{
														Description:         "NewTag is the value used to replace the original tag.",
														MarkdownDescription: "NewTag is the value used to replace the original tag.",
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

										"patches": schema.ListNestedAttribute{
											Description:         "Strategic merge and JSON patches, defined as inline YAML objects, capable of targeting objects based on kind, label and annotation selectors.",
											MarkdownDescription: "Strategic merge and JSON patches, defined as inline YAML objects, capable of targeting objects based on kind, label and annotation selectors.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"patch": schema.StringAttribute{
														Description:         "Patch contains an inline StrategicMerge patch or an inline JSON6902 patch with an array of operation objects.",
														MarkdownDescription: "Patch contains an inline StrategicMerge patch or an inline JSON6902 patch with an array of operation objects.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"target": schema.SingleNestedAttribute{
														Description:         "Target points to the resources that the patch document should be applied to.",
														MarkdownDescription: "Target points to the resources that the patch document should be applied to.",
														Attributes: map[string]schema.Attribute{
															"annotation_selector": schema.StringAttribute{
																Description:         "AnnotationSelector is a string that follows the label selection expression https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#api It matches with the resource annotations.",
																MarkdownDescription: "AnnotationSelector is a string that follows the label selection expression https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#api It matches with the resource annotations.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"group": schema.StringAttribute{
																Description:         "Group is the API group to select resources from. Together with Version and Kind it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",
																MarkdownDescription: "Group is the API group to select resources from. Together with Version and Kind it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kind": schema.StringAttribute{
																Description:         "Kind of the API Group to select resources from. Together with Group and Version it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",
																MarkdownDescription: "Kind of the API Group to select resources from. Together with Group and Version it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"label_selector": schema.StringAttribute{
																Description:         "LabelSelector is a string that follows the label selection expression https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#api It matches with the resource labels.",
																MarkdownDescription: "LabelSelector is a string that follows the label selection expression https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#api It matches with the resource labels.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "Name to match resources with.",
																MarkdownDescription: "Name to match resources with.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace to select resources from.",
																MarkdownDescription: "Namespace to select resources from.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"version": schema.StringAttribute{
																Description:         "Version of the API Group to select resources from. Together with Group and Kind it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",
																MarkdownDescription: "Version of the API Group to select resources from. Together with Group and Kind it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",
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

										"patches_json6902": schema.ListNestedAttribute{
											Description:         "JSON 6902 patches, defined as inline YAML objects.",
											MarkdownDescription: "JSON 6902 patches, defined as inline YAML objects.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"patch": schema.ListNestedAttribute{
														Description:         "Patch contains the JSON6902 patch document with an array of operation objects.",
														MarkdownDescription: "Patch contains the JSON6902 patch document with an array of operation objects.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"from": schema.StringAttribute{
																	Description:         "From contains a JSON-pointer value that references a location within the target document where the operation is performed. The meaning of the value depends on the value of Op, and is NOT taken into account by all operations.",
																	MarkdownDescription: "From contains a JSON-pointer value that references a location within the target document where the operation is performed. The meaning of the value depends on the value of Op, and is NOT taken into account by all operations.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"op": schema.StringAttribute{
																	Description:         "Op indicates the operation to perform. Its value MUST be one of 'add', 'remove', 'replace', 'move', 'copy', or 'test'. https://datatracker.ietf.org/doc/html/rfc6902#section-4",
																	MarkdownDescription: "Op indicates the operation to perform. Its value MUST be one of 'add', 'remove', 'replace', 'move', 'copy', or 'test'. https://datatracker.ietf.org/doc/html/rfc6902#section-4",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("test", "remove", "add", "replace", "move", "copy"),
																	},
																},

																"path": schema.StringAttribute{
																	Description:         "Path contains the JSON-pointer value that references a location within the target document where the operation is performed. The meaning of the value depends on the value of Op.",
																	MarkdownDescription: "Path contains the JSON-pointer value that references a location within the target document where the operation is performed. The meaning of the value depends on the value of Op.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"value": schema.MapAttribute{
																	Description:         "Value contains a valid JSON structure. The meaning of the value depends on the value of Op, and is NOT taken into account by all operations.",
																	MarkdownDescription: "Value contains a valid JSON structure. The meaning of the value depends on the value of Op, and is NOT taken into account by all operations.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
														},
														Required: true,
														Optional: false,
														Computed: false,
													},

													"target": schema.SingleNestedAttribute{
														Description:         "Target points to the resources that the patch document should be applied to.",
														MarkdownDescription: "Target points to the resources that the patch document should be applied to.",
														Attributes: map[string]schema.Attribute{
															"annotation_selector": schema.StringAttribute{
																Description:         "AnnotationSelector is a string that follows the label selection expression https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#api It matches with the resource annotations.",
																MarkdownDescription: "AnnotationSelector is a string that follows the label selection expression https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#api It matches with the resource annotations.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"group": schema.StringAttribute{
																Description:         "Group is the API group to select resources from. Together with Version and Kind it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",
																MarkdownDescription: "Group is the API group to select resources from. Together with Version and Kind it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kind": schema.StringAttribute{
																Description:         "Kind of the API Group to select resources from. Together with Group and Version it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",
																MarkdownDescription: "Kind of the API Group to select resources from. Together with Group and Version it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"label_selector": schema.StringAttribute{
																Description:         "LabelSelector is a string that follows the label selection expression https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#api It matches with the resource labels.",
																MarkdownDescription: "LabelSelector is a string that follows the label selection expression https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#api It matches with the resource labels.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "Name to match resources with.",
																MarkdownDescription: "Name to match resources with.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace to select resources from.",
																MarkdownDescription: "Namespace to select resources from.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"version": schema.StringAttribute{
																Description:         "Version of the API Group to select resources from. Together with Group and Kind it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",
																MarkdownDescription: "Version of the API Group to select resources from. Together with Group and Kind it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"patches_strategic_merge": schema.ListAttribute{
											Description:         "Strategic merge patches, defined as inline YAML objects.",
											MarkdownDescription: "Strategic merge patches, defined as inline YAML objects.",
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

					"release_name": schema.StringAttribute{
						Description:         "ReleaseName used for the Helm release. Defaults to a composition of '[TargetNamespace-]Name'.",
						MarkdownDescription: "ReleaseName used for the Helm release. Defaults to a composition of '[TargetNamespace-]Name'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
							stringvalidator.LengthAtMost(53),
						},
					},

					"rollback": schema.SingleNestedAttribute{
						Description:         "Rollback holds the configuration for Helm rollback actions for this HelmRelease.",
						MarkdownDescription: "Rollback holds the configuration for Helm rollback actions for this HelmRelease.",
						Attributes: map[string]schema.Attribute{
							"cleanup_on_fail": schema.BoolAttribute{
								Description:         "CleanupOnFail allows deletion of new resources created during the Helm rollback action when it fails.",
								MarkdownDescription: "CleanupOnFail allows deletion of new resources created during the Helm rollback action when it fails.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"disable_hooks": schema.BoolAttribute{
								Description:         "DisableHooks prevents hooks from running during the Helm rollback action.",
								MarkdownDescription: "DisableHooks prevents hooks from running during the Helm rollback action.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"disable_wait": schema.BoolAttribute{
								Description:         "DisableWait disables the waiting for resources to be ready after a Helm rollback has been performed.",
								MarkdownDescription: "DisableWait disables the waiting for resources to be ready after a Helm rollback has been performed.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"disable_wait_for_jobs": schema.BoolAttribute{
								Description:         "DisableWaitForJobs disables waiting for jobs to complete after a Helm rollback has been performed.",
								MarkdownDescription: "DisableWaitForJobs disables waiting for jobs to complete after a Helm rollback has been performed.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"force": schema.BoolAttribute{
								Description:         "Force forces resource updates through a replacement strategy.",
								MarkdownDescription: "Force forces resource updates through a replacement strategy.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"recreate": schema.BoolAttribute{
								Description:         "Recreate performs pod restarts for the resource if applicable.",
								MarkdownDescription: "Recreate performs pod restarts for the resource if applicable.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"timeout": schema.StringAttribute{
								Description:         "Timeout is the time to wait for any individual Kubernetes operation (like Jobs for hooks) during the performance of a Helm rollback action. Defaults to 'HelmReleaseSpec.Timeout'.",
								MarkdownDescription: "Timeout is the time to wait for any individual Kubernetes operation (like Jobs for hooks) during the performance of a Helm rollback action. Defaults to 'HelmReleaseSpec.Timeout'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ms|s|m|h))+$`), ""),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"service_account_name": schema.StringAttribute{
						Description:         "The name of the Kubernetes service account to impersonate when reconciling this HelmRelease.",
						MarkdownDescription: "The name of the Kubernetes service account to impersonate when reconciling this HelmRelease.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"storage_namespace": schema.StringAttribute{
						Description:         "StorageNamespace used for the Helm storage. Defaults to the namespace of the HelmRelease.",
						MarkdownDescription: "StorageNamespace used for the Helm storage. Defaults to the namespace of the HelmRelease.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
							stringvalidator.LengthAtMost(63),
						},
					},

					"suspend": schema.BoolAttribute{
						Description:         "Suspend tells the controller to suspend reconciliation for this HelmRelease, it does not apply to already started reconciliations. Defaults to false.",
						MarkdownDescription: "Suspend tells the controller to suspend reconciliation for this HelmRelease, it does not apply to already started reconciliations. Defaults to false.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"target_namespace": schema.StringAttribute{
						Description:         "TargetNamespace to target when performing operations for the HelmRelease. Defaults to the namespace of the HelmRelease.",
						MarkdownDescription: "TargetNamespace to target when performing operations for the HelmRelease. Defaults to the namespace of the HelmRelease.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
							stringvalidator.LengthAtMost(63),
						},
					},

					"test": schema.SingleNestedAttribute{
						Description:         "Test holds the configuration for Helm test actions for this HelmRelease.",
						MarkdownDescription: "Test holds the configuration for Helm test actions for this HelmRelease.",
						Attributes: map[string]schema.Attribute{
							"enable": schema.BoolAttribute{
								Description:         "Enable enables Helm test actions for this HelmRelease after an Helm install or upgrade action has been performed.",
								MarkdownDescription: "Enable enables Helm test actions for this HelmRelease after an Helm install or upgrade action has been performed.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ignore_failures": schema.BoolAttribute{
								Description:         "IgnoreFailures tells the controller to skip remediation when the Helm tests are run but fail. Can be overwritten for tests run after install or upgrade actions in 'Install.IgnoreTestFailures' and 'Upgrade.IgnoreTestFailures'.",
								MarkdownDescription: "IgnoreFailures tells the controller to skip remediation when the Helm tests are run but fail. Can be overwritten for tests run after install or upgrade actions in 'Install.IgnoreTestFailures' and 'Upgrade.IgnoreTestFailures'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"timeout": schema.StringAttribute{
								Description:         "Timeout is the time to wait for any individual Kubernetes operation during the performance of a Helm test action. Defaults to 'HelmReleaseSpec.Timeout'.",
								MarkdownDescription: "Timeout is the time to wait for any individual Kubernetes operation during the performance of a Helm test action. Defaults to 'HelmReleaseSpec.Timeout'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ms|s|m|h))+$`), ""),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"timeout": schema.StringAttribute{
						Description:         "Timeout is the time to wait for any individual Kubernetes operation (like Jobs for hooks) during the performance of a Helm action. Defaults to '5m0s'.",
						MarkdownDescription: "Timeout is the time to wait for any individual Kubernetes operation (like Jobs for hooks) during the performance of a Helm action. Defaults to '5m0s'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ms|s|m|h))+$`), ""),
						},
					},

					"uninstall": schema.SingleNestedAttribute{
						Description:         "Uninstall holds the configuration for Helm uninstall actions for this HelmRelease.",
						MarkdownDescription: "Uninstall holds the configuration for Helm uninstall actions for this HelmRelease.",
						Attributes: map[string]schema.Attribute{
							"deletion_propagation": schema.StringAttribute{
								Description:         "DeletionPropagation specifies the deletion propagation policy when a Helm uninstall is performed.",
								MarkdownDescription: "DeletionPropagation specifies the deletion propagation policy when a Helm uninstall is performed.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("background", "foreground", "orphan"),
								},
							},

							"disable_hooks": schema.BoolAttribute{
								Description:         "DisableHooks prevents hooks from running during the Helm rollback action.",
								MarkdownDescription: "DisableHooks prevents hooks from running during the Helm rollback action.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"disable_wait": schema.BoolAttribute{
								Description:         "DisableWait disables waiting for all the resources to be deleted after a Helm uninstall is performed.",
								MarkdownDescription: "DisableWait disables waiting for all the resources to be deleted after a Helm uninstall is performed.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"keep_history": schema.BoolAttribute{
								Description:         "KeepHistory tells Helm to remove all associated resources and mark the release as deleted, but retain the release history.",
								MarkdownDescription: "KeepHistory tells Helm to remove all associated resources and mark the release as deleted, but retain the release history.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"timeout": schema.StringAttribute{
								Description:         "Timeout is the time to wait for any individual Kubernetes operation (like Jobs for hooks) during the performance of a Helm uninstall action. Defaults to 'HelmReleaseSpec.Timeout'.",
								MarkdownDescription: "Timeout is the time to wait for any individual Kubernetes operation (like Jobs for hooks) during the performance of a Helm uninstall action. Defaults to 'HelmReleaseSpec.Timeout'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ms|s|m|h))+$`), ""),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"upgrade": schema.SingleNestedAttribute{
						Description:         "Upgrade holds the configuration for Helm upgrade actions for this HelmRelease.",
						MarkdownDescription: "Upgrade holds the configuration for Helm upgrade actions for this HelmRelease.",
						Attributes: map[string]schema.Attribute{
							"cleanup_on_fail": schema.BoolAttribute{
								Description:         "CleanupOnFail allows deletion of new resources created during the Helm upgrade action when it fails.",
								MarkdownDescription: "CleanupOnFail allows deletion of new resources created during the Helm upgrade action when it fails.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"crds": schema.StringAttribute{
								Description:         "CRDs upgrade CRDs from the Helm Chart's crds directory according to the CRD upgrade policy provided here. Valid values are 'Skip', 'Create' or 'CreateReplace'. Default is 'Skip' and if omitted CRDs are neither installed nor upgraded. Skip: do neither install nor replace (update) any CRDs. Create: new CRDs are created, existing CRDs are neither updated nor deleted. CreateReplace: new CRDs are created, existing CRDs are updated (replaced) but not deleted. By default, CRDs are not applied during Helm upgrade action. With this option users can opt-in to CRD upgrade, which is not (yet) natively supported by Helm. https://helm.sh/docs/chart_best_practices/custom_resource_definitions.",
								MarkdownDescription: "CRDs upgrade CRDs from the Helm Chart's crds directory according to the CRD upgrade policy provided here. Valid values are 'Skip', 'Create' or 'CreateReplace'. Default is 'Skip' and if omitted CRDs are neither installed nor upgraded. Skip: do neither install nor replace (update) any CRDs. Create: new CRDs are created, existing CRDs are neither updated nor deleted. CreateReplace: new CRDs are created, existing CRDs are updated (replaced) but not deleted. By default, CRDs are not applied during Helm upgrade action. With this option users can opt-in to CRD upgrade, which is not (yet) natively supported by Helm. https://helm.sh/docs/chart_best_practices/custom_resource_definitions.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Skip", "Create", "CreateReplace"),
								},
							},

							"disable_hooks": schema.BoolAttribute{
								Description:         "DisableHooks prevents hooks from running during the Helm upgrade action.",
								MarkdownDescription: "DisableHooks prevents hooks from running during the Helm upgrade action.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"disable_open_api_validation": schema.BoolAttribute{
								Description:         "DisableOpenAPIValidation prevents the Helm upgrade action from validating rendered templates against the Kubernetes OpenAPI Schema.",
								MarkdownDescription: "DisableOpenAPIValidation prevents the Helm upgrade action from validating rendered templates against the Kubernetes OpenAPI Schema.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"disable_wait": schema.BoolAttribute{
								Description:         "DisableWait disables the waiting for resources to be ready after a Helm upgrade has been performed.",
								MarkdownDescription: "DisableWait disables the waiting for resources to be ready after a Helm upgrade has been performed.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"disable_wait_for_jobs": schema.BoolAttribute{
								Description:         "DisableWaitForJobs disables waiting for jobs to complete after a Helm upgrade has been performed.",
								MarkdownDescription: "DisableWaitForJobs disables waiting for jobs to complete after a Helm upgrade has been performed.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"force": schema.BoolAttribute{
								Description:         "Force forces resource updates through a replacement strategy.",
								MarkdownDescription: "Force forces resource updates through a replacement strategy.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"preserve_values": schema.BoolAttribute{
								Description:         "PreserveValues will make Helm reuse the last release's values and merge in overrides from 'Values'. Setting this flag makes the HelmRelease non-declarative.",
								MarkdownDescription: "PreserveValues will make Helm reuse the last release's values and merge in overrides from 'Values'. Setting this flag makes the HelmRelease non-declarative.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"remediation": schema.SingleNestedAttribute{
								Description:         "Remediation holds the remediation configuration for when the Helm upgrade action for the HelmRelease fails. The default is to not perform any action.",
								MarkdownDescription: "Remediation holds the remediation configuration for when the Helm upgrade action for the HelmRelease fails. The default is to not perform any action.",
								Attributes: map[string]schema.Attribute{
									"ignore_test_failures": schema.BoolAttribute{
										Description:         "IgnoreTestFailures tells the controller to skip remediation when the Helm tests are run after an upgrade action but fail. Defaults to 'Test.IgnoreFailures'.",
										MarkdownDescription: "IgnoreTestFailures tells the controller to skip remediation when the Helm tests are run after an upgrade action but fail. Defaults to 'Test.IgnoreFailures'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"remediate_last_failure": schema.BoolAttribute{
										Description:         "RemediateLastFailure tells the controller to remediate the last failure, when no retries remain. Defaults to 'false' unless 'Retries' is greater than 0.",
										MarkdownDescription: "RemediateLastFailure tells the controller to remediate the last failure, when no retries remain. Defaults to 'false' unless 'Retries' is greater than 0.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"retries": schema.Int64Attribute{
										Description:         "Retries is the number of retries that should be attempted on failures before bailing. Remediation, using 'Strategy', is performed between each attempt. Defaults to '0', a negative integer equals to unlimited retries.",
										MarkdownDescription: "Retries is the number of retries that should be attempted on failures before bailing. Remediation, using 'Strategy', is performed between each attempt. Defaults to '0', a negative integer equals to unlimited retries.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"strategy": schema.StringAttribute{
										Description:         "Strategy to use for failure remediation. Defaults to 'rollback'.",
										MarkdownDescription: "Strategy to use for failure remediation. Defaults to 'rollback'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("rollback", "uninstall"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"timeout": schema.StringAttribute{
								Description:         "Timeout is the time to wait for any individual Kubernetes operation (like Jobs for hooks) during the performance of a Helm upgrade action. Defaults to 'HelmReleaseSpec.Timeout'.",
								MarkdownDescription: "Timeout is the time to wait for any individual Kubernetes operation (like Jobs for hooks) during the performance of a Helm upgrade action. Defaults to 'HelmReleaseSpec.Timeout'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ms|s|m|h))+$`), ""),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"values": schema.MapAttribute{
						Description:         "Values holds the values for this Helm release.",
						MarkdownDescription: "Values holds the values for this Helm release.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"values_from": schema.ListNestedAttribute{
						Description:         "ValuesFrom holds references to resources containing Helm values for this HelmRelease, and information about how they should be merged.",
						MarkdownDescription: "ValuesFrom holds references to resources containing Helm values for this HelmRelease, and information about how they should be merged.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"kind": schema.StringAttribute{
									Description:         "Kind of the values referent, valid values are ('Secret', 'ConfigMap').",
									MarkdownDescription: "Kind of the values referent, valid values are ('Secret', 'ConfigMap').",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("Secret", "ConfigMap"),
									},
								},

								"name": schema.StringAttribute{
									Description:         "Name of the values referent. Should reside in the same namespace as the referring resource.",
									MarkdownDescription: "Name of the values referent. Should reside in the same namespace as the referring resource.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(253),
									},
								},

								"optional": schema.BoolAttribute{
									Description:         "Optional marks this ValuesReference as optional. When set, a not found error for the values reference is ignored, but any ValuesKey, TargetPath or transient error will still result in a reconciliation failure.",
									MarkdownDescription: "Optional marks this ValuesReference as optional. When set, a not found error for the values reference is ignored, but any ValuesKey, TargetPath or transient error will still result in a reconciliation failure.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"target_path": schema.StringAttribute{
									Description:         "TargetPath is the YAML dot notation path the value should be merged at. When set, the ValuesKey is expected to be a single flat value. Defaults to 'None', which results in the values getting merged at the root.",
									MarkdownDescription: "TargetPath is the YAML dot notation path the value should be merged at. When set, the ValuesKey is expected to be a single flat value. Defaults to 'None', which results in the values getting merged at the root.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(250),
										stringvalidator.RegexMatches(regexp.MustCompile(`^([a-zA-Z0-9_\-.\\\/]|\[[0-9]{1,5}\])+$`), ""),
									},
								},

								"values_key": schema.StringAttribute{
									Description:         "ValuesKey is the data key where the values.yaml or a specific value can be found at. Defaults to 'values.yaml'. When set, must be a valid Data Key, consisting of alphanumeric characters, '-', '_' or '.'.",
									MarkdownDescription: "ValuesKey is the data key where the values.yaml or a specific value can be found at. Defaults to 'values.yaml'. When set, must be a valid Data Key, consisting of alphanumeric characters, '-', '_' or '.'.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(253),
										stringvalidator.RegexMatches(regexp.MustCompile(`^[\-._a-zA-Z0-9]+$`), ""),
									},
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
	}
}

func (r *HelmToolkitFluxcdIoHelmReleaseV2Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_helm_toolkit_fluxcd_io_helm_release_v2beta1_manifest")

	var model HelmToolkitFluxcdIoHelmReleaseV2Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("helm.toolkit.fluxcd.io/v2beta1")
	model.Kind = pointer.String("HelmRelease")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
