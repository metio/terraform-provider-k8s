/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package gitops_hybrid_cloud_patterns_io_v1alpha1

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
	_ datasource.DataSource = &GitopsHybridCloudPatternsIoPatternV1Alpha1Manifest{}
)

func NewGitopsHybridCloudPatternsIoPatternV1Alpha1Manifest() datasource.DataSource {
	return &GitopsHybridCloudPatternsIoPatternV1Alpha1Manifest{}
}

type GitopsHybridCloudPatternsIoPatternV1Alpha1Manifest struct{}

type GitopsHybridCloudPatternsIoPatternV1Alpha1ManifestData struct {
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
		AnalyticsUUID    *string `tfsdk:"analytics_uuid" json:"analyticsUUID,omitempty"`
		ClusterGroupName *string `tfsdk:"cluster_group_name" json:"clusterGroupName,omitempty"`
		ExtraParameters  *[]struct {
			Name  *string `tfsdk:"name" json:"name,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"extra_parameters" json:"extraParameters,omitempty"`
		ExtraValueFiles *[]string `tfsdk:"extra_value_files" json:"extraValueFiles,omitempty"`
		GitOpsSpec      *struct {
			ManualSync *bool `tfsdk:"manual_sync" json:"manualSync,omitempty"`
		} `tfsdk:"git_ops_spec" json:"gitOpsSpec,omitempty"`
		GitSpec *struct {
			Hostname             *string `tfsdk:"hostname" json:"hostname,omitempty"`
			OriginRepo           *string `tfsdk:"origin_repo" json:"originRepo,omitempty"`
			OriginRevision       *string `tfsdk:"origin_revision" json:"originRevision,omitempty"`
			PollInterval         *int64  `tfsdk:"poll_interval" json:"pollInterval,omitempty"`
			TargetRepo           *string `tfsdk:"target_repo" json:"targetRepo,omitempty"`
			TargetRevision       *string `tfsdk:"target_revision" json:"targetRevision,omitempty"`
			TokenSecret          *string `tfsdk:"token_secret" json:"tokenSecret,omitempty"`
			TokenSecretNamespace *string `tfsdk:"token_secret_namespace" json:"tokenSecretNamespace,omitempty"`
		} `tfsdk:"git_spec" json:"gitSpec,omitempty"`
		MultiSourceConfig *struct {
			ClusterGroupChartGitRevision *string `tfsdk:"cluster_group_chart_git_revision" json:"clusterGroupChartGitRevision,omitempty"`
			ClusterGroupChartVersion     *string `tfsdk:"cluster_group_chart_version" json:"clusterGroupChartVersion,omitempty"`
			ClusterGroupGitRepoUrl       *string `tfsdk:"cluster_group_git_repo_url" json:"clusterGroupGitRepoUrl,omitempty"`
			Enabled                      *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
			HelmRepoUrl                  *string `tfsdk:"helm_repo_url" json:"helmRepoUrl,omitempty"`
		} `tfsdk:"multi_source_config" json:"multiSourceConfig,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *GitopsHybridCloudPatternsIoPatternV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_gitops_hybrid_cloud_patterns_io_pattern_v1alpha1_manifest"
}

func (r *GitopsHybridCloudPatternsIoPatternV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Pattern is the Schema for the patterns API",
		MarkdownDescription: "Pattern is the Schema for the patterns API",
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
				Description:         "PatternSpec defines the desired state of Pattern",
				MarkdownDescription: "PatternSpec defines the desired state of Pattern",
				Attributes: map[string]schema.Attribute{
					"analytics_uuid": schema.StringAttribute{
						Description:         "Analytics UUID. Leave empty to autogenerate a random one. Not PII information",
						MarkdownDescription: "Analytics UUID. Leave empty to autogenerate a random one. Not PII information",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"cluster_group_name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"extra_parameters": schema.ListNestedAttribute{
						Description:         ".Name is dot separated per the helm --set syntax, such as: global.something.field",
						MarkdownDescription: ".Name is dot separated per the helm --set syntax, such as: global.something.field",
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

					"extra_value_files": schema.ListAttribute{
						Description:         "URLs to additional Helm parameter files",
						MarkdownDescription: "URLs to additional Helm parameter files",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"git_ops_spec": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"manual_sync": schema.BoolAttribute{
								Description:         "Require manual intervention before Argo will sync new content. Default: False",
								MarkdownDescription: "Require manual intervention before Argo will sync new content. Default: False",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"git_spec": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"hostname": schema.StringAttribute{
								Description:         "Optional. FQDN of the git server if automatic parsing from TargetRepo is broken",
								MarkdownDescription: "Optional. FQDN of the git server if automatic parsing from TargetRepo is broken",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"origin_repo": schema.StringAttribute{
								Description:         "Upstream git repo containing the pattern to deploy. Used when in-cluster fork to point to the upstream pattern repository",
								MarkdownDescription: "Upstream git repo containing the pattern to deploy. Used when in-cluster fork to point to the upstream pattern repository",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"origin_revision": schema.StringAttribute{
								Description:         "Branch, tag or commit in the upstream git repository. Does not support short-sha's. Default to HEAD",
								MarkdownDescription: "Branch, tag or commit in the upstream git repository. Does not support short-sha's. Default to HEAD",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"poll_interval": schema.Int64Attribute{
								Description:         "Interval in seconds to poll for drifts between origin and target repositories. Default: 180 seconds",
								MarkdownDescription: "Interval in seconds to poll for drifts between origin and target repositories. Default: 180 seconds",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"target_repo": schema.StringAttribute{
								Description:         "Git repo containing the pattern to deploy. Must use https/http or, for ssh, git@server:foo/bar.git",
								MarkdownDescription: "Git repo containing the pattern to deploy. Must use https/http or, for ssh, git@server:foo/bar.git",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"target_revision": schema.StringAttribute{
								Description:         "Branch, tag, or commit to deploy.  Does not support short-sha's. Default: HEAD",
								MarkdownDescription: "Branch, tag, or commit to deploy.  Does not support short-sha's. Default: HEAD",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"token_secret": schema.StringAttribute{
								Description:         "Optional. K8s secret name where the info for connecting to git can be found. The supported secrets are modeled after the private repositories in argo (https://argo-cd.readthedocs.io/en/stable/operator-manual/declarative-setup/#repositories) currently ssh and username+password are supported",
								MarkdownDescription: "Optional. K8s secret name where the info for connecting to git can be found. The supported secrets are modeled after the private repositories in argo (https://argo-cd.readthedocs.io/en/stable/operator-manual/declarative-setup/#repositories) currently ssh and username+password are supported",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"token_secret_namespace": schema.StringAttribute{
								Description:         "Optional. K8s secret namespace where the token for connecting to git can be found",
								MarkdownDescription: "Optional. K8s secret namespace where the token for connecting to git can be found",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"multi_source_config": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"cluster_group_chart_git_revision": schema.StringAttribute{
								Description:         "The git reference when deploying the clustergroup helm chart directly from a git repo Defaults to 'main'. (Only used when developing the clustergroup helm chart)",
								MarkdownDescription: "The git reference when deploying the clustergroup helm chart directly from a git repo Defaults to 'main'. (Only used when developing the clustergroup helm chart)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cluster_group_chart_version": schema.StringAttribute{
								Description:         "Which chart version for the clustergroup helm chart. Defaults to '0.8.*'",
								MarkdownDescription: "Which chart version for the clustergroup helm chart. Defaults to '0.8.*'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cluster_group_git_repo_url": schema.StringAttribute{
								Description:         "The url when deploying the clustergroup helm chart directly from a git repo Defaults to '' which means not used (Only used when developing the clustergroup helm chart)",
								MarkdownDescription: "The url when deploying the clustergroup helm chart directly from a git repo Defaults to '' which means not used (Only used when developing the clustergroup helm chart)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enabled": schema.BoolAttribute{
								Description:         "(EXPERIMENTAL) Enable multi-source support when deploying the clustergroup argo application",
								MarkdownDescription: "(EXPERIMENTAL) Enable multi-source support when deploying the clustergroup argo application",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"helm_repo_url": schema.StringAttribute{
								Description:         "The helm chart url to fetch the helm charts from in order to deploy the pattern. Defaults to https://charts.validatedpatterns.io/",
								MarkdownDescription: "The helm chart url to fetch the helm charts from in order to deploy the pattern. Defaults to https://charts.validatedpatterns.io/",
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
	}
}

func (r *GitopsHybridCloudPatternsIoPatternV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_gitops_hybrid_cloud_patterns_io_pattern_v1alpha1_manifest")

	var model GitopsHybridCloudPatternsIoPatternV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("gitops.hybrid-cloud-patterns.io/v1alpha1")
	model.Kind = pointer.String("Pattern")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
