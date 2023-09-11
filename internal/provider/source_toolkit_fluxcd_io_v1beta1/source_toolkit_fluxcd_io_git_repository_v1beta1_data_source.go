/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package source_toolkit_fluxcd_io_v1beta1

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
	_ datasource.DataSource              = &SourceToolkitFluxcdIoGitRepositoryV1Beta1DataSource{}
	_ datasource.DataSourceWithConfigure = &SourceToolkitFluxcdIoGitRepositoryV1Beta1DataSource{}
)

func NewSourceToolkitFluxcdIoGitRepositoryV1Beta1DataSource() datasource.DataSource {
	return &SourceToolkitFluxcdIoGitRepositoryV1Beta1DataSource{}
}

type SourceToolkitFluxcdIoGitRepositoryV1Beta1DataSource struct {
	kubernetesClient dynamic.Interface
}

type SourceToolkitFluxcdIoGitRepositoryV1Beta1DataSourceData struct {
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

func (r *SourceToolkitFluxcdIoGitRepositoryV1Beta1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_source_toolkit_fluxcd_io_git_repository_v1beta1"
}

func (r *SourceToolkitFluxcdIoGitRepositoryV1Beta1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "GitRepository is the Schema for the gitrepositories API",
		MarkdownDescription: "GitRepository is the Schema for the gitrepositories API",
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
				Description:         "GitRepositorySpec defines the desired state of a Git repository.",
				MarkdownDescription: "GitRepositorySpec defines the desired state of a Git repository.",
				Attributes: map[string]schema.Attribute{
					"access_from": schema.SingleNestedAttribute{
						Description:         "AccessFrom defines an Access Control List for allowing cross-namespace references to this object.",
						MarkdownDescription: "AccessFrom defines an Access Control List for allowing cross-namespace references to this object.",
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
											Optional:            false,
											Computed:            true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"git_implementation": schema.StringAttribute{
						Description:         "Determines which git client library to use. Defaults to go-git, valid values are ('go-git', 'libgit2').",
						MarkdownDescription: "Determines which git client library to use. Defaults to go-git, valid values are ('go-git', 'libgit2').",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"ignore": schema.StringAttribute{
						Description:         "Ignore overrides the set of excluded patterns in the .sourceignore format (which is the same as .gitignore). If not provided, a default will be used, consult the documentation for your version to find out what those are.",
						MarkdownDescription: "Ignore overrides the set of excluded patterns in the .sourceignore format (which is the same as .gitignore). If not provided, a default will be used, consult the documentation for your version to find out what those are.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"include": schema.ListNestedAttribute{
						Description:         "Extra git repositories to map into the repository",
						MarkdownDescription: "Extra git repositories to map into the repository",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"from_path": schema.StringAttribute{
									Description:         "The path to copy contents from, defaults to the root directory.",
									MarkdownDescription: "The path to copy contents from, defaults to the root directory.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"repository": schema.SingleNestedAttribute{
									Description:         "Reference to a GitRepository to include.",
									MarkdownDescription: "Reference to a GitRepository to include.",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name of the referent.",
											MarkdownDescription: "Name of the referent.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"to_path": schema.StringAttribute{
									Description:         "The path to copy contents to, defaults to the name of the source ref.",
									MarkdownDescription: "The path to copy contents to, defaults to the name of the source ref.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"interval": schema.StringAttribute{
						Description:         "The interval at which to check for repository updates.",
						MarkdownDescription: "The interval at which to check for repository updates.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"recurse_submodules": schema.BoolAttribute{
						Description:         "When enabled, after the clone is created, initializes all submodules within, using their default settings. This option is available only when using the 'go-git' GitImplementation.",
						MarkdownDescription: "When enabled, after the clone is created, initializes all submodules within, using their default settings. This option is available only when using the 'go-git' GitImplementation.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"ref": schema.SingleNestedAttribute{
						Description:         "The Git reference to checkout and monitor for changes, defaults to master branch.",
						MarkdownDescription: "The Git reference to checkout and monitor for changes, defaults to master branch.",
						Attributes: map[string]schema.Attribute{
							"branch": schema.StringAttribute{
								Description:         "The Git branch to checkout, defaults to master.",
								MarkdownDescription: "The Git branch to checkout, defaults to master.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"commit": schema.StringAttribute{
								Description:         "The Git commit SHA to checkout, if specified Tag filters will be ignored.",
								MarkdownDescription: "The Git commit SHA to checkout, if specified Tag filters will be ignored.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"semver": schema.StringAttribute{
								Description:         "The Git tag semver expression, takes precedence over Tag.",
								MarkdownDescription: "The Git tag semver expression, takes precedence over Tag.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"tag": schema.StringAttribute{
								Description:         "The Git tag to checkout, takes precedence over Branch.",
								MarkdownDescription: "The Git tag to checkout, takes precedence over Branch.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"secret_ref": schema.SingleNestedAttribute{
						Description:         "The secret name containing the Git credentials. For HTTPS repositories the secret must contain username and password fields. For SSH repositories the secret must contain identity and known_hosts fields.",
						MarkdownDescription: "The secret name containing the Git credentials. For HTTPS repositories the secret must contain username and password fields. For SSH repositories the secret must contain identity and known_hosts fields.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of the referent.",
								MarkdownDescription: "Name of the referent.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"suspend": schema.BoolAttribute{
						Description:         "This flag tells the controller to suspend the reconciliation of this source.",
						MarkdownDescription: "This flag tells the controller to suspend the reconciliation of this source.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"timeout": schema.StringAttribute{
						Description:         "The timeout for remote Git operations like cloning, defaults to 60s.",
						MarkdownDescription: "The timeout for remote Git operations like cloning, defaults to 60s.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"url": schema.StringAttribute{
						Description:         "The repository URL, can be a HTTP/S or SSH address.",
						MarkdownDescription: "The repository URL, can be a HTTP/S or SSH address.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"verify": schema.SingleNestedAttribute{
						Description:         "Verify OpenPGP signature for the Git commit HEAD points to.",
						MarkdownDescription: "Verify OpenPGP signature for the Git commit HEAD points to.",
						Attributes: map[string]schema.Attribute{
							"mode": schema.StringAttribute{
								Description:         "Mode describes what git object should be verified, currently ('head').",
								MarkdownDescription: "Mode describes what git object should be verified, currently ('head').",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"secret_ref": schema.SingleNestedAttribute{
								Description:         "The secret name containing the public keys of all trusted Git authors.",
								MarkdownDescription: "The secret name containing the public keys of all trusted Git authors.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of the referent.",
										MarkdownDescription: "Name of the referent.",
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
						Required: false,
						Optional: false,
						Computed: true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *SourceToolkitFluxcdIoGitRepositoryV1Beta1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *SourceToolkitFluxcdIoGitRepositoryV1Beta1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_source_toolkit_fluxcd_io_git_repository_v1beta1")

	var data SourceToolkitFluxcdIoGitRepositoryV1Beta1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "source.toolkit.fluxcd.io", Version: "v1beta1", Resource: "gitrepositories"}).
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

	var readResponse SourceToolkitFluxcdIoGitRepositoryV1Beta1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("source.toolkit.fluxcd.io/v1beta1")
	data.Kind = pointer.String("GitRepository")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
