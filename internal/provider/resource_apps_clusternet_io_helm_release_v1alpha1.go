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

type AppsClusternetIoHelmReleaseV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*AppsClusternetIoHelmReleaseV1Alpha1Resource)(nil)
)

type AppsClusternetIoHelmReleaseV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type AppsClusternetIoHelmReleaseV1Alpha1GoModel struct {
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
		Atomic *bool `tfsdk:"atomic" yaml:"atomic,omitempty"`

		Chart *string `tfsdk:"chart" yaml:"chart,omitempty"`

		ChartPullSecret *struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
		} `tfsdk:"chart_pull_secret" yaml:"chartPullSecret,omitempty"`

		CreateNamespace *bool `tfsdk:"create_namespace" yaml:"createNamespace,omitempty"`

		DisableHooks *bool `tfsdk:"disable_hooks" yaml:"disableHooks,omitempty"`

		Force *bool `tfsdk:"force" yaml:"force,omitempty"`

		Overrides *string `tfsdk:"overrides" yaml:"overrides,omitempty"`

		ReleaseName *string `tfsdk:"release_name" yaml:"releaseName,omitempty"`

		Replace *bool `tfsdk:"replace" yaml:"replace,omitempty"`

		Repo *string `tfsdk:"repo" yaml:"repo,omitempty"`

		SkipCRDs *bool `tfsdk:"skip_cr_ds" yaml:"skipCRDs,omitempty"`

		TargetNamespace *string `tfsdk:"target_namespace" yaml:"targetNamespace,omitempty"`

		TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`

		Version *string `tfsdk:"version" yaml:"version,omitempty"`

		Wait *bool `tfsdk:"wait" yaml:"wait,omitempty"`

		WaitForJob *bool `tfsdk:"wait_for_job" yaml:"waitForJob,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewAppsClusternetIoHelmReleaseV1Alpha1Resource() resource.Resource {
	return &AppsClusternetIoHelmReleaseV1Alpha1Resource{}
}

func (r *AppsClusternetIoHelmReleaseV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_apps_clusternet_io_helm_release_v1alpha1"
}

func (r *AppsClusternetIoHelmReleaseV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "HelmRelease is the Schema for the helm release",
		MarkdownDescription: "HelmRelease is the Schema for the helm release",
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
				Description:         "HelmReleaseSpec defines the spec of HelmRelease",
				MarkdownDescription: "HelmReleaseSpec defines the spec of HelmRelease",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"atomic": {
						Description:         "Atomic, if true, for install case, will uninstall failed release, for upgrade case, will roll back on failure.",
						MarkdownDescription: "Atomic, if true, for install case, will uninstall failed release, for upgrade case, will roll back on failure.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"chart": {
						Description:         "Chart is the name of a Helm Chart in the Repository.",
						MarkdownDescription: "Chart is the name of a Helm Chart in the Repository.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"chart_pull_secret": {
						Description:         "ChartPullSecret is the name of the secret that contains the auth information for the chart repository.",
						MarkdownDescription: "ChartPullSecret is the name of the secret that contains the auth information for the chart repository.",

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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"create_namespace": {
						Description:         "CreateNamespace create namespace when install helm release",
						MarkdownDescription: "CreateNamespace create namespace when install helm release",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"disable_hooks": {
						Description:         "DisableHooks disables hook processing if set to true.",
						MarkdownDescription: "DisableHooks disables hook processing if set to true.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"force": {
						Description:         "Force will, if set to 'true', ignore certain warnings and perform the upgrade anyway. This should be used with caution.",
						MarkdownDescription: "Force will, if set to 'true', ignore certain warnings and perform the upgrade anyway. This should be used with caution.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"overrides": {
						Description:         "Overrides specifies the override values for this release.",
						MarkdownDescription: "Overrides specifies the override values for this release.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							validators.Base64Validator(),
						},
					},

					"release_name": {
						Description:         "ReleaseName specifies the desired release name in child cluster. If nil, the default release name will be in the format of '{Description Name}-{HelmChart Namespace}-{HelmChart Name}'",
						MarkdownDescription: "ReleaseName specifies the desired release name in child cluster. If nil, the default release name will be in the format of '{Description Name}-{HelmChart Namespace}-{HelmChart Name}'",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"replace": {
						Description:         "Replace will re-use the given name, only if that name is a deleted release that remains in the history. This is unsafe in production.",
						MarkdownDescription: "Replace will re-use the given name, only if that name is a deleted release that remains in the history. This is unsafe in production.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"repo": {
						Description:         "a Helm Repository to be used. OCI-based registries are also supported. For example, https://charts.bitnami.com/bitnami or oci://localhost:5000/helm-charts",
						MarkdownDescription: "a Helm Repository to be used. OCI-based registries are also supported. For example, https://charts.bitnami.com/bitnami or oci://localhost:5000/helm-charts",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.RegexMatches(regexp.MustCompile(`^(http|https|oci)?://(?:[a-zA-Z]|[0-9]|[$-_@.&+]|[!*\(\),]|(?:%[0-9a-fA-F][0-9a-fA-F]))+$`), ""),
						},
					},

					"skip_cr_ds": {
						Description:         "SkipCRDs skips installing CRDs when install flag is enabled during upgrade",
						MarkdownDescription: "SkipCRDs skips installing CRDs when install flag is enabled during upgrade",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"target_namespace": {
						Description:         "TargetNamespace specifies the namespace to install the chart",
						MarkdownDescription: "TargetNamespace specifies the namespace to install the chart",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"timeout_seconds": {
						Description:         "TimeoutSeconds is the timeout of the chart to be install/upgrade/rollback/uninstall",
						MarkdownDescription: "TimeoutSeconds is the timeout of the chart to be install/upgrade/rollback/uninstall",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"version": {
						Description:         "ChartVersion is the version of the chart to be deployed. It will be defaulted with current latest version if empty.",
						MarkdownDescription: "ChartVersion is the version of the chart to be deployed. It will be defaulted with current latest version if empty.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"wait": {
						Description:         "Wait determines whether the wait operation should be performed after the upgrade is requested.",
						MarkdownDescription: "Wait determines whether the wait operation should be performed after the upgrade is requested.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"wait_for_job": {
						Description:         "WaitForJobs determines whether the wait operation for the Jobs should be performed after the upgrade is requested.",
						MarkdownDescription: "WaitForJobs determines whether the wait operation for the Jobs should be performed after the upgrade is requested.",

						Type: types.BoolType,

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

func (r *AppsClusternetIoHelmReleaseV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_apps_clusternet_io_helm_release_v1alpha1")

	var state AppsClusternetIoHelmReleaseV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel AppsClusternetIoHelmReleaseV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("apps.clusternet.io/v1alpha1")
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

func (r *AppsClusternetIoHelmReleaseV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_apps_clusternet_io_helm_release_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *AppsClusternetIoHelmReleaseV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_apps_clusternet_io_helm_release_v1alpha1")

	var state AppsClusternetIoHelmReleaseV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel AppsClusternetIoHelmReleaseV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("apps.clusternet.io/v1alpha1")
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

func (r *AppsClusternetIoHelmReleaseV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_apps_clusternet_io_helm_release_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
