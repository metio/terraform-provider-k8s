/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package image_toolkit_fluxcd_io_v1beta1

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
	_ datasource.DataSource              = &ImageToolkitFluxcdIoImageUpdateAutomationV1Beta1DataSource{}
	_ datasource.DataSourceWithConfigure = &ImageToolkitFluxcdIoImageUpdateAutomationV1Beta1DataSource{}
)

func NewImageToolkitFluxcdIoImageUpdateAutomationV1Beta1DataSource() datasource.DataSource {
	return &ImageToolkitFluxcdIoImageUpdateAutomationV1Beta1DataSource{}
}

type ImageToolkitFluxcdIoImageUpdateAutomationV1Beta1DataSource struct {
	kubernetesClient dynamic.Interface
}

type ImageToolkitFluxcdIoImageUpdateAutomationV1Beta1DataSourceData struct {
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
		Git *struct {
			Checkout *struct {
				Ref *struct {
					Branch *string `tfsdk:"branch" json:"branch,omitempty"`
					Commit *string `tfsdk:"commit" json:"commit,omitempty"`
					Name   *string `tfsdk:"name" json:"name,omitempty"`
					Semver *string `tfsdk:"semver" json:"semver,omitempty"`
					Tag    *string `tfsdk:"tag" json:"tag,omitempty"`
				} `tfsdk:"ref" json:"ref,omitempty"`
			} `tfsdk:"checkout" json:"checkout,omitempty"`
			Commit *struct {
				Author *struct {
					Email *string `tfsdk:"email" json:"email,omitempty"`
					Name  *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"author" json:"author,omitempty"`
				MessageTemplate *string `tfsdk:"message_template" json:"messageTemplate,omitempty"`
				SigningKey      *struct {
					SecretRef *struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
				} `tfsdk:"signing_key" json:"signingKey,omitempty"`
			} `tfsdk:"commit" json:"commit,omitempty"`
			Push *struct {
				Branch  *string            `tfsdk:"branch" json:"branch,omitempty"`
				Options *map[string]string `tfsdk:"options" json:"options,omitempty"`
				Refspec *string            `tfsdk:"refspec" json:"refspec,omitempty"`
			} `tfsdk:"push" json:"push,omitempty"`
		} `tfsdk:"git" json:"git,omitempty"`
		Interval  *string `tfsdk:"interval" json:"interval,omitempty"`
		SourceRef *struct {
			ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
			Name       *string `tfsdk:"name" json:"name,omitempty"`
			Namespace  *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"source_ref" json:"sourceRef,omitempty"`
		Suspend *bool `tfsdk:"suspend" json:"suspend,omitempty"`
		Update  *struct {
			Path     *string `tfsdk:"path" json:"path,omitempty"`
			Strategy *string `tfsdk:"strategy" json:"strategy,omitempty"`
		} `tfsdk:"update" json:"update,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ImageToolkitFluxcdIoImageUpdateAutomationV1Beta1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_image_toolkit_fluxcd_io_image_update_automation_v1beta1"
}

func (r *ImageToolkitFluxcdIoImageUpdateAutomationV1Beta1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ImageUpdateAutomation is the Schema for the imageupdateautomations API",
		MarkdownDescription: "ImageUpdateAutomation is the Schema for the imageupdateautomations API",
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
				Description:         "ImageUpdateAutomationSpec defines the desired state of ImageUpdateAutomation",
				MarkdownDescription: "ImageUpdateAutomationSpec defines the desired state of ImageUpdateAutomation",
				Attributes: map[string]schema.Attribute{
					"git": schema.SingleNestedAttribute{
						Description:         "GitSpec contains all the git-specific definitions. This is technically optional, but in practice mandatory until there are other kinds of source allowed.",
						MarkdownDescription: "GitSpec contains all the git-specific definitions. This is technically optional, but in practice mandatory until there are other kinds of source allowed.",
						Attributes: map[string]schema.Attribute{
							"checkout": schema.SingleNestedAttribute{
								Description:         "Checkout gives the parameters for cloning the git repository, ready to make changes. If not present, the 'spec.ref' field from the referenced 'GitRepository' or its default will be used.",
								MarkdownDescription: "Checkout gives the parameters for cloning the git repository, ready to make changes. If not present, the 'spec.ref' field from the referenced 'GitRepository' or its default will be used.",
								Attributes: map[string]schema.Attribute{
									"ref": schema.SingleNestedAttribute{
										Description:         "Reference gives a branch, tag or commit to clone from the Git repository.",
										MarkdownDescription: "Reference gives a branch, tag or commit to clone from the Git repository.",
										Attributes: map[string]schema.Attribute{
											"branch": schema.StringAttribute{
												Description:         "Branch to check out, defaults to 'master' if no other field is defined.",
												MarkdownDescription: "Branch to check out, defaults to 'master' if no other field is defined.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"commit": schema.StringAttribute{
												Description:         "Commit SHA to check out, takes precedence over all reference fields.  This can be combined with Branch to shallow clone the branch, in which the commit is expected to exist.",
												MarkdownDescription: "Commit SHA to check out, takes precedence over all reference fields.  This can be combined with Branch to shallow clone the branch, in which the commit is expected to exist.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the reference to check out; takes precedence over Branch, Tag and SemVer.  It must be a valid Git reference: https://git-scm.com/docs/git-check-ref-format#_description Examples: 'refs/heads/main', 'refs/tags/v0.1.0', 'refs/pull/420/head', 'refs/merge-requests/1/head'",
												MarkdownDescription: "Name of the reference to check out; takes precedence over Branch, Tag and SemVer.  It must be a valid Git reference: https://git-scm.com/docs/git-check-ref-format#_description Examples: 'refs/heads/main', 'refs/tags/v0.1.0', 'refs/pull/420/head', 'refs/merge-requests/1/head'",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"semver": schema.StringAttribute{
												Description:         "SemVer tag expression to check out, takes precedence over Tag.",
												MarkdownDescription: "SemVer tag expression to check out, takes precedence over Tag.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"tag": schema.StringAttribute{
												Description:         "Tag to check out, takes precedence over Branch.",
												MarkdownDescription: "Tag to check out, takes precedence over Branch.",
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

							"commit": schema.SingleNestedAttribute{
								Description:         "Commit specifies how to commit to the git repository.",
								MarkdownDescription: "Commit specifies how to commit to the git repository.",
								Attributes: map[string]schema.Attribute{
									"author": schema.SingleNestedAttribute{
										Description:         "Author gives the email and optionally the name to use as the author of commits.",
										MarkdownDescription: "Author gives the email and optionally the name to use as the author of commits.",
										Attributes: map[string]schema.Attribute{
											"email": schema.StringAttribute{
												Description:         "Email gives the email to provide when making a commit.",
												MarkdownDescription: "Email gives the email to provide when making a commit.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"name": schema.StringAttribute{
												Description:         "Name gives the name to provide when making a commit.",
												MarkdownDescription: "Name gives the name to provide when making a commit.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"message_template": schema.StringAttribute{
										Description:         "MessageTemplate provides a template for the commit message, into which will be interpolated the details of the change made.",
										MarkdownDescription: "MessageTemplate provides a template for the commit message, into which will be interpolated the details of the change made.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"signing_key": schema.SingleNestedAttribute{
										Description:         "SigningKey provides the option to sign commits with a GPG key",
										MarkdownDescription: "SigningKey provides the option to sign commits with a GPG key",
										Attributes: map[string]schema.Attribute{
											"secret_ref": schema.SingleNestedAttribute{
												Description:         "SecretRef holds the name to a secret that contains a 'git.asc' key corresponding to the ASCII Armored file containing the GPG signing keypair as the value. It must be in the same namespace as the ImageUpdateAutomation.",
												MarkdownDescription: "SecretRef holds the name to a secret that contains a 'git.asc' key corresponding to the ASCII Armored file containing the GPG signing keypair as the value. It must be in the same namespace as the ImageUpdateAutomation.",
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

							"push": schema.SingleNestedAttribute{
								Description:         "Push specifies how and where to push commits made by the automation. If missing, commits are pushed (back) to '.spec.checkout.branch' or its default.",
								MarkdownDescription: "Push specifies how and where to push commits made by the automation. If missing, commits are pushed (back) to '.spec.checkout.branch' or its default.",
								Attributes: map[string]schema.Attribute{
									"branch": schema.StringAttribute{
										Description:         "Branch specifies that commits should be pushed to the branch named. The branch is created using '.spec.checkout.branch' as the starting point, if it doesn't already exist.",
										MarkdownDescription: "Branch specifies that commits should be pushed to the branch named. The branch is created using '.spec.checkout.branch' as the starting point, if it doesn't already exist.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"options": schema.MapAttribute{
										Description:         "Options specifies the push options that are sent to the Git server when performing a push operation. For details, see: https://git-scm.com/docs/git-push#Documentation/git-push.txt---push-optionltoptiongt",
										MarkdownDescription: "Options specifies the push options that are sent to the Git server when performing a push operation. For details, see: https://git-scm.com/docs/git-push#Documentation/git-push.txt---push-optionltoptiongt",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"refspec": schema.StringAttribute{
										Description:         "Refspec specifies the Git Refspec to use for a push operation. If both Branch and Refspec are provided, then the commit is pushed to the branch and also using the specified refspec. For more details about Git Refspecs, see: https://git-scm.com/book/en/v2/Git-Internals-The-Refspec",
										MarkdownDescription: "Refspec specifies the Git Refspec to use for a push operation. If both Branch and Refspec are provided, then the commit is pushed to the branch and also using the specified refspec. For more details about Git Refspecs, see: https://git-scm.com/book/en/v2/Git-Internals-The-Refspec",
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

					"interval": schema.StringAttribute{
						Description:         "Interval gives an lower bound for how often the automation run should be attempted.",
						MarkdownDescription: "Interval gives an lower bound for how often the automation run should be attempted.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"source_ref": schema.SingleNestedAttribute{
						Description:         "SourceRef refers to the resource giving access details to a git repository.",
						MarkdownDescription: "SourceRef refers to the resource giving access details to a git repository.",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "API version of the referent.",
								MarkdownDescription: "API version of the referent.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind of the referent.",
								MarkdownDescription: "Kind of the referent.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referent.",
								MarkdownDescription: "Name of the referent.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace of the referent, defaults to the namespace of the Kubernetes resource object that contains the reference.",
								MarkdownDescription: "Namespace of the referent, defaults to the namespace of the Kubernetes resource object that contains the reference.",
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
						Description:         "Suspend tells the controller to not run this automation, until it is unset (or set to false). Defaults to false.",
						MarkdownDescription: "Suspend tells the controller to not run this automation, until it is unset (or set to false). Defaults to false.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"update": schema.SingleNestedAttribute{
						Description:         "Update gives the specification for how to update the files in the repository. This can be left empty, to use the default value.",
						MarkdownDescription: "Update gives the specification for how to update the files in the repository. This can be left empty, to use the default value.",
						Attributes: map[string]schema.Attribute{
							"path": schema.StringAttribute{
								Description:         "Path to the directory containing the manifests to be updated. Defaults to 'None', which translates to the root path of the GitRepositoryRef.",
								MarkdownDescription: "Path to the directory containing the manifests to be updated. Defaults to 'None', which translates to the root path of the GitRepositoryRef.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"strategy": schema.StringAttribute{
								Description:         "Strategy names the strategy to be used.",
								MarkdownDescription: "Strategy names the strategy to be used.",
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
	}
}

func (r *ImageToolkitFluxcdIoImageUpdateAutomationV1Beta1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *ImageToolkitFluxcdIoImageUpdateAutomationV1Beta1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_image_toolkit_fluxcd_io_image_update_automation_v1beta1")

	var data ImageToolkitFluxcdIoImageUpdateAutomationV1Beta1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "image.toolkit.fluxcd.io", Version: "v1beta1", Resource: "imageupdateautomations"}).
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

	var readResponse ImageToolkitFluxcdIoImageUpdateAutomationV1Beta1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("image.toolkit.fluxcd.io/v1beta1")
	data.Kind = pointer.String("ImageUpdateAutomation")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
