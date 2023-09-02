/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package source_toolkit_fluxcd_io_v1beta2

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
	_ datasource.DataSource = &SourceToolkitFluxcdIoGitRepositoryV1Beta2Manifest{}
)

func NewSourceToolkitFluxcdIoGitRepositoryV1Beta2Manifest() datasource.DataSource {
	return &SourceToolkitFluxcdIoGitRepositoryV1Beta2Manifest{}
}

type SourceToolkitFluxcdIoGitRepositoryV1Beta2Manifest struct{}

type SourceToolkitFluxcdIoGitRepositoryV1Beta2ManifestData struct {
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
		AccessFrom *struct {
			NamespaceSelectors *[]struct {
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"namespace_selectors" json:"namespaceSelectors,omitempty"`
		} `tfsdk:"access_from" json:"accessFrom,omitempty"`
		GitImplementation *string `tfsdk:"git_implementation" json:"gitImplementation,omitempty"`
		Ignore            *string `tfsdk:"ignore" json:"ignore,omitempty"`
		Include           *[]struct {
			FromPath   *string `tfsdk:"from_path" json:"fromPath,omitempty"`
			Repository *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"repository" json:"repository,omitempty"`
			ToPath *string `tfsdk:"to_path" json:"toPath,omitempty"`
		} `tfsdk:"include" json:"include,omitempty"`
		Interval          *string `tfsdk:"interval" json:"interval,omitempty"`
		RecurseSubmodules *bool   `tfsdk:"recurse_submodules" json:"recurseSubmodules,omitempty"`
		Ref               *struct {
			Branch *string `tfsdk:"branch" json:"branch,omitempty"`
			Commit *string `tfsdk:"commit" json:"commit,omitempty"`
			Name   *string `tfsdk:"name" json:"name,omitempty"`
			Semver *string `tfsdk:"semver" json:"semver,omitempty"`
			Tag    *string `tfsdk:"tag" json:"tag,omitempty"`
		} `tfsdk:"ref" json:"ref,omitempty"`
		SecretRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
		Suspend *bool   `tfsdk:"suspend" json:"suspend,omitempty"`
		Timeout *string `tfsdk:"timeout" json:"timeout,omitempty"`
		Url     *string `tfsdk:"url" json:"url,omitempty"`
		Verify  *struct {
			Mode      *string `tfsdk:"mode" json:"mode,omitempty"`
			SecretRef *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
		} `tfsdk:"verify" json:"verify,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SourceToolkitFluxcdIoGitRepositoryV1Beta2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_source_toolkit_fluxcd_io_git_repository_v1beta2_manifest"
}

func (r *SourceToolkitFluxcdIoGitRepositoryV1Beta2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "GitRepository is the Schema for the gitrepositories API.",
		MarkdownDescription: "GitRepository is the Schema for the gitrepositories API.",
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
				Description:         "GitRepositorySpec specifies the required configuration to produce an Artifact for a Git repository.",
				MarkdownDescription: "GitRepositorySpec specifies the required configuration to produce an Artifact for a Git repository.",
				Attributes: map[string]schema.Attribute{
					"access_from": schema.SingleNestedAttribute{
						Description:         "AccessFrom specifies an Access Control List for allowing cross-namespace references to this object. NOTE: Not implemented, provisional as of https://github.com/fluxcd/flux2/pull/2092",
						MarkdownDescription: "AccessFrom specifies an Access Control List for allowing cross-namespace references to this object. NOTE: Not implemented, provisional as of https://github.com/fluxcd/flux2/pull/2092",
						Attributes: map[string]schema.Attribute{
							"namespace_selectors": schema.ListNestedAttribute{
								Description:         "NamespaceSelectors is the list of namespace selectors to which this ACL applies. Items in this list are evaluated using a logical OR operation.",
								MarkdownDescription: "NamespaceSelectors is the list of namespace selectors to which this ACL applies. Items in this list are evaluated using a logical OR operation.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"match_labels": schema.MapAttribute{
											Description:         "MatchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
											MarkdownDescription: "MatchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"git_implementation": schema.StringAttribute{
						Description:         "GitImplementation specifies which Git client library implementation to use. Defaults to 'go-git', valid values are ('go-git', 'libgit2'). Deprecated: gitImplementation is deprecated now that 'go-git' is the only supported implementation.",
						MarkdownDescription: "GitImplementation specifies which Git client library implementation to use. Defaults to 'go-git', valid values are ('go-git', 'libgit2'). Deprecated: gitImplementation is deprecated now that 'go-git' is the only supported implementation.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("go-git", "libgit2"),
						},
					},

					"ignore": schema.StringAttribute{
						Description:         "Ignore overrides the set of excluded patterns in the .sourceignore format (which is the same as .gitignore). If not provided, a default will be used, consult the documentation for your version to find out what those are.",
						MarkdownDescription: "Ignore overrides the set of excluded patterns in the .sourceignore format (which is the same as .gitignore). If not provided, a default will be used, consult the documentation for your version to find out what those are.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"include": schema.ListNestedAttribute{
						Description:         "Include specifies a list of GitRepository resources which Artifacts should be included in the Artifact produced for this GitRepository.",
						MarkdownDescription: "Include specifies a list of GitRepository resources which Artifacts should be included in the Artifact produced for this GitRepository.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"from_path": schema.StringAttribute{
									Description:         "FromPath specifies the path to copy contents from, defaults to the root of the Artifact.",
									MarkdownDescription: "FromPath specifies the path to copy contents from, defaults to the root of the Artifact.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"repository": schema.SingleNestedAttribute{
									Description:         "GitRepositoryRef specifies the GitRepository which Artifact contents must be included.",
									MarkdownDescription: "GitRepositoryRef specifies the GitRepository which Artifact contents must be included.",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name of the referent.",
											MarkdownDescription: "Name of the referent.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},

								"to_path": schema.StringAttribute{
									Description:         "ToPath specifies the path to copy contents to, defaults to the name of the GitRepositoryRef.",
									MarkdownDescription: "ToPath specifies the path to copy contents to, defaults to the name of the GitRepositoryRef.",
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

					"interval": schema.StringAttribute{
						Description:         "Interval at which to check the GitRepository for updates.",
						MarkdownDescription: "Interval at which to check the GitRepository for updates.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ms|s|m|h))+$`), ""),
						},
					},

					"recurse_submodules": schema.BoolAttribute{
						Description:         "RecurseSubmodules enables the initialization of all submodules within the GitRepository as cloned from the URL, using their default settings.",
						MarkdownDescription: "RecurseSubmodules enables the initialization of all submodules within the GitRepository as cloned from the URL, using their default settings.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ref": schema.SingleNestedAttribute{
						Description:         "Reference specifies the Git reference to resolve and monitor for changes, defaults to the 'master' branch.",
						MarkdownDescription: "Reference specifies the Git reference to resolve and monitor for changes, defaults to the 'master' branch.",
						Attributes: map[string]schema.Attribute{
							"branch": schema.StringAttribute{
								Description:         "Branch to check out, defaults to 'master' if no other field is defined.",
								MarkdownDescription: "Branch to check out, defaults to 'master' if no other field is defined.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"commit": schema.StringAttribute{
								Description:         "Commit SHA to check out, takes precedence over all reference fields.  This can be combined with Branch to shallow clone the branch, in which the commit is expected to exist.",
								MarkdownDescription: "Commit SHA to check out, takes precedence over all reference fields.  This can be combined with Branch to shallow clone the branch, in which the commit is expected to exist.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the reference to check out; takes precedence over Branch, Tag and SemVer.  It must be a valid Git reference: https://git-scm.com/docs/git-check-ref-format#_description Examples: 'refs/heads/main', 'refs/tags/v0.1.0', 'refs/pull/420/head', 'refs/merge-requests/1/head'",
								MarkdownDescription: "Name of the reference to check out; takes precedence over Branch, Tag and SemVer.  It must be a valid Git reference: https://git-scm.com/docs/git-check-ref-format#_description Examples: 'refs/heads/main', 'refs/tags/v0.1.0', 'refs/pull/420/head', 'refs/merge-requests/1/head'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"semver": schema.StringAttribute{
								Description:         "SemVer tag expression to check out, takes precedence over Tag.",
								MarkdownDescription: "SemVer tag expression to check out, takes precedence over Tag.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tag": schema.StringAttribute{
								Description:         "Tag to check out, takes precedence over Branch.",
								MarkdownDescription: "Tag to check out, takes precedence over Branch.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"secret_ref": schema.SingleNestedAttribute{
						Description:         "SecretRef specifies the Secret containing authentication credentials for the GitRepository. For HTTPS repositories the Secret must contain 'username' and 'password' fields for basic auth or 'bearerToken' field for token auth. For SSH repositories the Secret must contain 'identity' and 'known_hosts' fields.",
						MarkdownDescription: "SecretRef specifies the Secret containing authentication credentials for the GitRepository. For HTTPS repositories the Secret must contain 'username' and 'password' fields for basic auth or 'bearerToken' field for token auth. For SSH repositories the Secret must contain 'identity' and 'known_hosts' fields.",
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

					"suspend": schema.BoolAttribute{
						Description:         "Suspend tells the controller to suspend the reconciliation of this GitRepository.",
						MarkdownDescription: "Suspend tells the controller to suspend the reconciliation of this GitRepository.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"timeout": schema.StringAttribute{
						Description:         "Timeout for Git operations like cloning, defaults to 60s.",
						MarkdownDescription: "Timeout for Git operations like cloning, defaults to 60s.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ms|s|m))+$`), ""),
						},
					},

					"url": schema.StringAttribute{
						Description:         "URL specifies the Git repository URL, it can be an HTTP/S or SSH address.",
						MarkdownDescription: "URL specifies the Git repository URL, it can be an HTTP/S or SSH address.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(http|https|ssh)://.*$`), ""),
						},
					},

					"verify": schema.SingleNestedAttribute{
						Description:         "Verification specifies the configuration to verify the Git commit signature(s).",
						MarkdownDescription: "Verification specifies the configuration to verify the Git commit signature(s).",
						Attributes: map[string]schema.Attribute{
							"mode": schema.StringAttribute{
								Description:         "Mode specifies what Git object should be verified, currently ('head').",
								MarkdownDescription: "Mode specifies what Git object should be verified, currently ('head').",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("head"),
								},
							},

							"secret_ref": schema.SingleNestedAttribute{
								Description:         "SecretRef specifies the Secret containing the public keys of trusted Git authors.",
								MarkdownDescription: "SecretRef specifies the Secret containing the public keys of trusted Git authors.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of the referent.",
										MarkdownDescription: "Name of the referent.",
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *SourceToolkitFluxcdIoGitRepositoryV1Beta2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_source_toolkit_fluxcd_io_git_repository_v1beta2_manifest")

	var model SourceToolkitFluxcdIoGitRepositoryV1Beta2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("source.toolkit.fluxcd.io/v1beta2")
	model.Kind = pointer.String("GitRepository")

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
