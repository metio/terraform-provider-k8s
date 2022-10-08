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

type SourceToolkitFluxcdIoGitRepositoryV1Beta2Resource struct{}

var (
	_ resource.Resource = (*SourceToolkitFluxcdIoGitRepositoryV1Beta2Resource)(nil)
)

type SourceToolkitFluxcdIoGitRepositoryV1Beta2TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type SourceToolkitFluxcdIoGitRepositoryV1Beta2GoModel struct {
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
		Ignore *string `tfsdk:"ignore" yaml:"ignore,omitempty"`

		Interval *string `tfsdk:"interval" yaml:"interval,omitempty"`

		Ref *struct {
			Branch *string `tfsdk:"branch" yaml:"branch,omitempty"`

			Commit *string `tfsdk:"commit" yaml:"commit,omitempty"`

			Semver *string `tfsdk:"semver" yaml:"semver,omitempty"`

			Tag *string `tfsdk:"tag" yaml:"tag,omitempty"`
		} `tfsdk:"ref" yaml:"ref,omitempty"`

		Timeout *string `tfsdk:"timeout" yaml:"timeout,omitempty"`

		SecretRef *struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

		Suspend *bool `tfsdk:"suspend" yaml:"suspend,omitempty"`

		Url *string `tfsdk:"url" yaml:"url,omitempty"`

		Verify *struct {
			SecretRef *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

			Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`
		} `tfsdk:"verify" yaml:"verify,omitempty"`

		AccessFrom *struct {
			NamespaceSelectors *[]struct {
				MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
			} `tfsdk:"namespace_selectors" yaml:"namespaceSelectors,omitempty"`
		} `tfsdk:"access_from" yaml:"accessFrom,omitempty"`

		GitImplementation *string `tfsdk:"git_implementation" yaml:"gitImplementation,omitempty"`

		Include *[]struct {
			FromPath *string `tfsdk:"from_path" yaml:"fromPath,omitempty"`

			Repository *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"repository" yaml:"repository,omitempty"`

			ToPath *string `tfsdk:"to_path" yaml:"toPath,omitempty"`
		} `tfsdk:"include" yaml:"include,omitempty"`

		RecurseSubmodules *bool `tfsdk:"recurse_submodules" yaml:"recurseSubmodules,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewSourceToolkitFluxcdIoGitRepositoryV1Beta2Resource() resource.Resource {
	return &SourceToolkitFluxcdIoGitRepositoryV1Beta2Resource{}
}

func (r *SourceToolkitFluxcdIoGitRepositoryV1Beta2Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_toolkit_fluxcd_io_git_repository_v1beta2"
}

func (r *SourceToolkitFluxcdIoGitRepositoryV1Beta2Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "GitRepository is the Schema for the gitrepositories API.",
		MarkdownDescription: "GitRepository is the Schema for the gitrepositories API.",
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

			"spec": {
				Description:         "GitRepositorySpec specifies the required configuration to produce an Artifact for a Git repository.",
				MarkdownDescription: "GitRepositorySpec specifies the required configuration to produce an Artifact for a Git repository.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"ignore": {
						Description:         "Ignore overrides the set of excluded patterns in the .sourceignore format (which is the same as .gitignore). If not provided, a default will be used, consult the documentation for your version to find out what those are.",
						MarkdownDescription: "Ignore overrides the set of excluded patterns in the .sourceignore format (which is the same as .gitignore). If not provided, a default will be used, consult the documentation for your version to find out what those are.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"interval": {
						Description:         "Interval at which to check the GitRepository for updates.",
						MarkdownDescription: "Interval at which to check the GitRepository for updates.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"ref": {
						Description:         "Reference specifies the Git reference to resolve and monitor for changes, defaults to the 'master' branch.",
						MarkdownDescription: "Reference specifies the Git reference to resolve and monitor for changes, defaults to the 'master' branch.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"branch": {
								Description:         "Branch to check out, defaults to 'master' if no other field is defined.  When GitRepositorySpec.GitImplementation is set to 'go-git', a shallow clone of the specified branch is performed.",
								MarkdownDescription: "Branch to check out, defaults to 'master' if no other field is defined.  When GitRepositorySpec.GitImplementation is set to 'go-git', a shallow clone of the specified branch is performed.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"commit": {
								Description:         "Commit SHA to check out, takes precedence over all reference fields.  When GitRepositorySpec.GitImplementation is set to 'go-git', this can be combined with Branch to shallow clone the branch, in which the commit is expected to exist.",
								MarkdownDescription: "Commit SHA to check out, takes precedence over all reference fields.  When GitRepositorySpec.GitImplementation is set to 'go-git', this can be combined with Branch to shallow clone the branch, in which the commit is expected to exist.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"semver": {
								Description:         "SemVer tag expression to check out, takes precedence over Tag.",
								MarkdownDescription: "SemVer tag expression to check out, takes precedence over Tag.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"tag": {
								Description:         "Tag to check out, takes precedence over Branch.",
								MarkdownDescription: "Tag to check out, takes precedence over Branch.",

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

					"timeout": {
						Description:         "Timeout for Git operations like cloning, defaults to 60s.",
						MarkdownDescription: "Timeout for Git operations like cloning, defaults to 60s.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"secret_ref": {
						Description:         "SecretRef specifies the Secret containing authentication credentials for the GitRepository. For HTTPS repositories the Secret must contain 'username' and 'password' fields. For SSH repositories the Secret must contain 'identity' and 'known_hosts' fields.",
						MarkdownDescription: "SecretRef specifies the Secret containing authentication credentials for the GitRepository. For HTTPS repositories the Secret must contain 'username' and 'password' fields. For SSH repositories the Secret must contain 'identity' and 'known_hosts' fields.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "Name of the referent.",
								MarkdownDescription: "Name of the referent.",

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

					"suspend": {
						Description:         "Suspend tells the controller to suspend the reconciliation of this GitRepository.",
						MarkdownDescription: "Suspend tells the controller to suspend the reconciliation of this GitRepository.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"url": {
						Description:         "URL specifies the Git repository URL, it can be an HTTP/S or SSH address.",
						MarkdownDescription: "URL specifies the Git repository URL, it can be an HTTP/S or SSH address.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"verify": {
						Description:         "Verification specifies the configuration to verify the Git commit signature(s).",
						MarkdownDescription: "Verification specifies the configuration to verify the Git commit signature(s).",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"secret_ref": {
								Description:         "SecretRef specifies the Secret containing the public keys of trusted Git authors.",
								MarkdownDescription: "SecretRef specifies the Secret containing the public keys of trusted Git authors.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Name of the referent.",
										MarkdownDescription: "Name of the referent.",

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

							"mode": {
								Description:         "Mode specifies what Git object should be verified, currently ('head').",
								MarkdownDescription: "Mode specifies what Git object should be verified, currently ('head').",

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

					"access_from": {
						Description:         "AccessFrom specifies an Access Control List for allowing cross-namespace references to this object. NOTE: Not implemented, provisional as of https://github.com/fluxcd/flux2/pull/2092",
						MarkdownDescription: "AccessFrom specifies an Access Control List for allowing cross-namespace references to this object. NOTE: Not implemented, provisional as of https://github.com/fluxcd/flux2/pull/2092",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"namespace_selectors": {
								Description:         "NamespaceSelectors is the list of namespace selectors to which this ACL applies. Items in this list are evaluated using a logical OR operation.",
								MarkdownDescription: "NamespaceSelectors is the list of namespace selectors to which this ACL applies. Items in this list are evaluated using a logical OR operation.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"match_labels": {
										Description:         "MatchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "MatchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

										Type: types.MapType{ElemType: types.StringType},

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

					"git_implementation": {
						Description:         "GitImplementation specifies which Git client library implementation to use. Defaults to 'go-git', valid values are ('go-git', 'libgit2').",
						MarkdownDescription: "GitImplementation specifies which Git client library implementation to use. Defaults to 'go-git', valid values are ('go-git', 'libgit2').",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"include": {
						Description:         "Include specifies a list of GitRepository resources which Artifacts should be included in the Artifact produced for this GitRepository.",
						MarkdownDescription: "Include specifies a list of GitRepository resources which Artifacts should be included in the Artifact produced for this GitRepository.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"from_path": {
								Description:         "FromPath specifies the path to copy contents from, defaults to the root of the Artifact.",
								MarkdownDescription: "FromPath specifies the path to copy contents from, defaults to the root of the Artifact.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"repository": {
								Description:         "GitRepositoryRef specifies the GitRepository which Artifact contents must be included.",
								MarkdownDescription: "GitRepositoryRef specifies the GitRepository which Artifact contents must be included.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Name of the referent.",
										MarkdownDescription: "Name of the referent.",

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

							"to_path": {
								Description:         "ToPath specifies the path to copy contents to, defaults to the name of the GitRepositoryRef.",
								MarkdownDescription: "ToPath specifies the path to copy contents to, defaults to the name of the GitRepositoryRef.",

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

					"recurse_submodules": {
						Description:         "RecurseSubmodules enables the initialization of all submodules within the GitRepository as cloned from the URL, using their default settings. This option is available only when using the 'go-git' GitImplementation.",
						MarkdownDescription: "RecurseSubmodules enables the initialization of all submodules within the GitRepository as cloned from the URL, using their default settings. This option is available only when using the 'go-git' GitImplementation.",

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
		},
	}, nil
}

func (r *SourceToolkitFluxcdIoGitRepositoryV1Beta2Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_source_toolkit_fluxcd_io_git_repository_v1beta2")

	var state SourceToolkitFluxcdIoGitRepositoryV1Beta2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel SourceToolkitFluxcdIoGitRepositoryV1Beta2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("source.toolkit.fluxcd.io/v1beta2")
	goModel.Kind = utilities.Ptr("GitRepository")

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

func (r *SourceToolkitFluxcdIoGitRepositoryV1Beta2Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_source_toolkit_fluxcd_io_git_repository_v1beta2")
	// NO-OP: All data is already in Terraform state
}

func (r *SourceToolkitFluxcdIoGitRepositoryV1Beta2Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_source_toolkit_fluxcd_io_git_repository_v1beta2")

	var state SourceToolkitFluxcdIoGitRepositoryV1Beta2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel SourceToolkitFluxcdIoGitRepositoryV1Beta2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("source.toolkit.fluxcd.io/v1beta2")
	goModel.Kind = utilities.Ptr("GitRepository")

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

func (r *SourceToolkitFluxcdIoGitRepositoryV1Beta2Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_source_toolkit_fluxcd_io_git_repository_v1beta2")
	// NO-OP: Terraform removes the state automatically for us
}
