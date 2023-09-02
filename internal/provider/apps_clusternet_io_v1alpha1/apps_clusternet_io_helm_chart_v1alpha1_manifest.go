/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package apps_clusternet_io_v1alpha1

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &AppsClusternetIoHelmChartV1Alpha1Manifest{}
)

func NewAppsClusternetIoHelmChartV1Alpha1Manifest() datasource.DataSource {
	return &AppsClusternetIoHelmChartV1Alpha1Manifest{}
}

type AppsClusternetIoHelmChartV1Alpha1Manifest struct{}

type AppsClusternetIoHelmChartV1Alpha1ManifestData struct {
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

	Spec *struct {
		Atomic          *bool   `tfsdk:"atomic" json:"atomic,omitempty"`
		Chart           *string `tfsdk:"chart" json:"chart,omitempty"`
		ChartPullSecret *struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"chart_pull_secret" json:"chartPullSecret,omitempty"`
		CreateNamespace *bool   `tfsdk:"create_namespace" json:"createNamespace,omitempty"`
		DisableHooks    *bool   `tfsdk:"disable_hooks" json:"disableHooks,omitempty"`
		Force           *bool   `tfsdk:"force" json:"force,omitempty"`
		Replace         *bool   `tfsdk:"replace" json:"replace,omitempty"`
		ReplaceCRDs     *bool   `tfsdk:"replace_cr_ds" json:"replaceCRDs,omitempty"`
		Repo            *string `tfsdk:"repo" json:"repo,omitempty"`
		SkipCRDs        *bool   `tfsdk:"skip_cr_ds" json:"skipCRDs,omitempty"`
		TargetNamespace *string `tfsdk:"target_namespace" json:"targetNamespace,omitempty"`
		TimeoutSeconds  *int64  `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
		UpgradeAtomic   *bool   `tfsdk:"upgrade_atomic" json:"upgradeAtomic,omitempty"`
		Version         *string `tfsdk:"version" json:"version,omitempty"`
		Wait            *bool   `tfsdk:"wait" json:"wait,omitempty"`
		WaitForJob      *bool   `tfsdk:"wait_for_job" json:"waitForJob,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AppsClusternetIoHelmChartV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_apps_clusternet_io_helm_chart_v1alpha1_manifest"
}

func (r *AppsClusternetIoHelmChartV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "HelmChart is the Schema for the helm chart",
		MarkdownDescription: "HelmChart is the Schema for the helm chart",
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

			"spec": schema.SingleNestedAttribute{
				Description:         "HelmChartSpec defines the spec of HelmChart",
				MarkdownDescription: "HelmChartSpec defines the spec of HelmChart",
				Attributes: map[string]schema.Attribute{
					"atomic": schema.BoolAttribute{
						Description:         "Atomic, for install case, if true, will uninstall failed release.",
						MarkdownDescription: "Atomic, for install case, if true, will uninstall failed release.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"chart": schema.StringAttribute{
						Description:         "Chart is the name of a Helm Chart in the Repository.",
						MarkdownDescription: "Chart is the name of a Helm Chart in the Repository.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"chart_pull_secret": schema.SingleNestedAttribute{
						Description:         "ChartPullSecret is the name of the secret that contains the auth information for the chart repository.",
						MarkdownDescription: "ChartPullSecret is the name of the secret that contains the auth information for the chart repository.",
						Attributes: map[string]schema.Attribute{
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"create_namespace": schema.BoolAttribute{
						Description:         "CreateNamespace create namespace when install helm release",
						MarkdownDescription: "CreateNamespace create namespace when install helm release",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"disable_hooks": schema.BoolAttribute{
						Description:         "DisableHooks disables hook processing if set to true.",
						MarkdownDescription: "DisableHooks disables hook processing if set to true.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"force": schema.BoolAttribute{
						Description:         "Force will, if set to 'true', ignore certain warnings and perform the upgrade anyway. This should be used with caution.",
						MarkdownDescription: "Force will, if set to 'true', ignore certain warnings and perform the upgrade anyway. This should be used with caution.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"replace": schema.BoolAttribute{
						Description:         "Replace will re-use the given name, only if that name is a deleted release that remains in the history. This is unsafe in production.",
						MarkdownDescription: "Replace will re-use the given name, only if that name is a deleted release that remains in the history. This is unsafe in production.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"replace_cr_ds": schema.BoolAttribute{
						Description:         "ReplaceCRDs replace all crds in chart and sub charts before upgrade and install, not working when SkipCRDs true",
						MarkdownDescription: "ReplaceCRDs replace all crds in chart and sub charts before upgrade and install, not working when SkipCRDs true",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"repo": schema.StringAttribute{
						Description:         "a Helm Repository to be used. OCI-based registries are also supported. For example, https://charts.bitnami.com/bitnami or oci://localhost:5000/helm-charts",
						MarkdownDescription: "a Helm Repository to be used. OCI-based registries are also supported. For example, https://charts.bitnami.com/bitnami or oci://localhost:5000/helm-charts",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(http|https|oci)?://(?:[a-zA-Z]|[0-9]|[$-_@.&+]|[!*\(\),]|(?:%[0-9a-fA-F][0-9a-fA-F]))+$`), ""),
						},
					},

					"skip_cr_ds": schema.BoolAttribute{
						Description:         "SkipCRDs skips installing CRDs when install flag is enabled during upgrade",
						MarkdownDescription: "SkipCRDs skips installing CRDs when install flag is enabled during upgrade",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"target_namespace": schema.StringAttribute{
						Description:         "TargetNamespace specifies the namespace to install this HelmChart",
						MarkdownDescription: "TargetNamespace specifies the namespace to install this HelmChart",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"timeout_seconds": schema.Int64Attribute{
						Description:         "TimeoutSeconds is the timeout of the chart to be install/upgrade/rollback/uninstall",
						MarkdownDescription: "TimeoutSeconds is the timeout of the chart to be install/upgrade/rollback/uninstall",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"upgrade_atomic": schema.BoolAttribute{
						Description:         "UpgradeAtomic, for upgrade case, if true, will roll back failed release.",
						MarkdownDescription: "UpgradeAtomic, for upgrade case, if true, will roll back failed release.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"version": schema.StringAttribute{
						Description:         "ChartVersion is the version of the chart to be deployed. It will be defaulted with current latest version if empty.",
						MarkdownDescription: "ChartVersion is the version of the chart to be deployed. It will be defaulted with current latest version if empty.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"wait": schema.BoolAttribute{
						Description:         "Wait determines whether the wait operation should be performed after helm install, upgrade or uninstall is requested.",
						MarkdownDescription: "Wait determines whether the wait operation should be performed after helm install, upgrade or uninstall is requested.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"wait_for_job": schema.BoolAttribute{
						Description:         "WaitForJobs determines whether the wait operation for the Jobs should be performed after the upgrade is requested.",
						MarkdownDescription: "WaitForJobs determines whether the wait operation for the Jobs should be performed after the upgrade is requested.",
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

func (r *AppsClusternetIoHelmChartV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_apps_clusternet_io_helm_chart_v1alpha1_manifest")

	var model AppsClusternetIoHelmChartV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("apps.clusternet.io/v1alpha1")
	model.Kind = pointer.String("HelmChart")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"YAML Error: "+err.Error(),
		)
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
