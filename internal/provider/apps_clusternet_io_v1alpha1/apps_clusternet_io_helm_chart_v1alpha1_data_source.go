/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package apps_clusternet_io_v1alpha1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &AppsClusternetIoHelmChartV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &AppsClusternetIoHelmChartV1Alpha1DataSource{}
)

func NewAppsClusternetIoHelmChartV1Alpha1DataSource() datasource.DataSource {
	return &AppsClusternetIoHelmChartV1Alpha1DataSource{}
}

type AppsClusternetIoHelmChartV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type AppsClusternetIoHelmChartV1Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

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

func (r *AppsClusternetIoHelmChartV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_apps_clusternet_io_helm_chart_v1alpha1"
}

func (r *AppsClusternetIoHelmChartV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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

			"api_version": schema.StringAttribute{
				Description:         "The API group of the requested resource.",
				MarkdownDescription: "The API group of the requested resource.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"kind": schema.StringAttribute{
				Description:         "The type of the requested resource.",
				MarkdownDescription: "The type of the requested resource.",
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
						Optional:            false,
						Computed:            true,
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
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
						Optional:            false,
						Computed:            true,
					},

					"chart": schema.StringAttribute{
						Description:         "Chart is the name of a Helm Chart in the Repository.",
						MarkdownDescription: "Chart is the name of a Helm Chart in the Repository.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"chart_pull_secret": schema.SingleNestedAttribute{
						Description:         "ChartPullSecret is the name of the secret that contains the auth information for the chart repository.",
						MarkdownDescription: "ChartPullSecret is the name of the secret that contains the auth information for the chart repository.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"namespace": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"create_namespace": schema.BoolAttribute{
						Description:         "CreateNamespace create namespace when install helm release",
						MarkdownDescription: "CreateNamespace create namespace when install helm release",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"disable_hooks": schema.BoolAttribute{
						Description:         "DisableHooks disables hook processing if set to true.",
						MarkdownDescription: "DisableHooks disables hook processing if set to true.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"force": schema.BoolAttribute{
						Description:         "Force will, if set to 'true', ignore certain warnings and perform the upgrade anyway. This should be used with caution.",
						MarkdownDescription: "Force will, if set to 'true', ignore certain warnings and perform the upgrade anyway. This should be used with caution.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"replace": schema.BoolAttribute{
						Description:         "Replace will re-use the given name, only if that name is a deleted release that remains in the history. This is unsafe in production.",
						MarkdownDescription: "Replace will re-use the given name, only if that name is a deleted release that remains in the history. This is unsafe in production.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"replace_cr_ds": schema.BoolAttribute{
						Description:         "ReplaceCRDs replace all crds in chart and sub charts before upgrade and install, not working when SkipCRDs true",
						MarkdownDescription: "ReplaceCRDs replace all crds in chart and sub charts before upgrade and install, not working when SkipCRDs true",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"repo": schema.StringAttribute{
						Description:         "a Helm Repository to be used. OCI-based registries are also supported. For example, https://charts.bitnami.com/bitnami or oci://localhost:5000/helm-charts",
						MarkdownDescription: "a Helm Repository to be used. OCI-based registries are also supported. For example, https://charts.bitnami.com/bitnami or oci://localhost:5000/helm-charts",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"skip_cr_ds": schema.BoolAttribute{
						Description:         "SkipCRDs skips installing CRDs when install flag is enabled during upgrade",
						MarkdownDescription: "SkipCRDs skips installing CRDs when install flag is enabled during upgrade",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"target_namespace": schema.StringAttribute{
						Description:         "TargetNamespace specifies the namespace to install this HelmChart",
						MarkdownDescription: "TargetNamespace specifies the namespace to install this HelmChart",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"timeout_seconds": schema.Int64Attribute{
						Description:         "TimeoutSeconds is the timeout of the chart to be install/upgrade/rollback/uninstall",
						MarkdownDescription: "TimeoutSeconds is the timeout of the chart to be install/upgrade/rollback/uninstall",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"upgrade_atomic": schema.BoolAttribute{
						Description:         "UpgradeAtomic, for upgrade case, if true, will roll back failed release.",
						MarkdownDescription: "UpgradeAtomic, for upgrade case, if true, will roll back failed release.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"version": schema.StringAttribute{
						Description:         "ChartVersion is the version of the chart to be deployed. It will be defaulted with current latest version if empty.",
						MarkdownDescription: "ChartVersion is the version of the chart to be deployed. It will be defaulted with current latest version if empty.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"wait": schema.BoolAttribute{
						Description:         "Wait determines whether the wait operation should be performed after helm install, upgrade or uninstall is requested.",
						MarkdownDescription: "Wait determines whether the wait operation should be performed after helm install, upgrade or uninstall is requested.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"wait_for_job": schema.BoolAttribute{
						Description:         "WaitForJobs determines whether the wait operation for the Jobs should be performed after the upgrade is requested.",
						MarkdownDescription: "WaitForJobs determines whether the wait operation for the Jobs should be performed after the upgrade is requested.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *AppsClusternetIoHelmChartV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedDataSourceDataError(request.ProviderData))
	}
}

func (r *AppsClusternetIoHelmChartV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_apps_clusternet_io_helm_chart_v1alpha1")

	var data AppsClusternetIoHelmChartV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "apps.clusternet.io", Version: "v1alpha1", Resource: "helmcharts"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetNamespacedResourceError(err, data.Metadata.Name, data.Metadata.Namespace))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse AppsClusternetIoHelmChartV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("apps.clusternet.io/v1alpha1")
	data.Kind = pointer.String("HelmChart")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
