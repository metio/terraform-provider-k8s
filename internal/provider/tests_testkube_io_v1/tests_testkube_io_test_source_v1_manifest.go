/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package tests_testkube_io_v1

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
	_ datasource.DataSource = &TestsTestkubeIoTestSourceV1Manifest{}
)

func NewTestsTestkubeIoTestSourceV1Manifest() datasource.DataSource {
	return &TestsTestkubeIoTestSourceV1Manifest{}
}

type TestsTestkubeIoTestSourceV1Manifest struct{}

type TestsTestkubeIoTestSourceV1ManifestData struct {
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
		Data       *string `tfsdk:"data" json:"data,omitempty"`
		Repository *struct {
			AuthType          *string `tfsdk:"auth_type" json:"authType,omitempty"`
			Branch            *string `tfsdk:"branch" json:"branch,omitempty"`
			CertificateSecret *string `tfsdk:"certificate_secret" json:"certificateSecret,omitempty"`
			Commit            *string `tfsdk:"commit" json:"commit,omitempty"`
			Path              *string `tfsdk:"path" json:"path,omitempty"`
			TokenSecret       *struct {
				Key  *string `tfsdk:"key" json:"key,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"token_secret" json:"tokenSecret,omitempty"`
			Type           *string `tfsdk:"type" json:"type,omitempty"`
			Uri            *string `tfsdk:"uri" json:"uri,omitempty"`
			UsernameSecret *struct {
				Key  *string `tfsdk:"key" json:"key,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"username_secret" json:"usernameSecret,omitempty"`
			WorkingDir *string `tfsdk:"working_dir" json:"workingDir,omitempty"`
		} `tfsdk:"repository" json:"repository,omitempty"`
		Type *string `tfsdk:"type" json:"type,omitempty"`
		Uri  *string `tfsdk:"uri" json:"uri,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *TestsTestkubeIoTestSourceV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_tests_testkube_io_test_source_v1_manifest"
}

func (r *TestsTestkubeIoTestSourceV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "TestSource is the Schema for the testsources API",
		MarkdownDescription: "TestSource is the Schema for the testsources API",
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
				Description:         "TestSourceSpec defines the desired state of TestSource",
				MarkdownDescription: "TestSourceSpec defines the desired state of TestSource",
				Attributes: map[string]schema.Attribute{
					"data": schema.StringAttribute{
						Description:         "test content body",
						MarkdownDescription: "test content body",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"repository": schema.SingleNestedAttribute{
						Description:         "repository of test content",
						MarkdownDescription: "repository of test content",
						Attributes: map[string]schema.Attribute{
							"auth_type": schema.StringAttribute{
								Description:         "auth type for git requests",
								MarkdownDescription: "auth type for git requests",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("basic", "header"),
								},
							},

							"branch": schema.StringAttribute{
								Description:         "branch/tag name for checkout",
								MarkdownDescription: "branch/tag name for checkout",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"certificate_secret": schema.StringAttribute{
								Description:         "git auth certificate secret for private repositories",
								MarkdownDescription: "git auth certificate secret for private repositories",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"commit": schema.StringAttribute{
								Description:         "commit id (sha) for checkout",
								MarkdownDescription: "commit id (sha) for checkout",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"path": schema.StringAttribute{
								Description:         "If specified, does a sparse checkout of the repository at the given path",
								MarkdownDescription: "If specified, does a sparse checkout of the repository at the given path",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"token_secret": schema.SingleNestedAttribute{
								Description:         "SecretRef is the Testkube internal reference for secret storage in Kubernetes secrets",
								MarkdownDescription: "SecretRef is the Testkube internal reference for secret storage in Kubernetes secrets",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "object key",
										MarkdownDescription: "object key",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "object name",
										MarkdownDescription: "object name",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"type": schema.StringAttribute{
								Description:         "VCS repository type",
								MarkdownDescription: "VCS repository type",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"uri": schema.StringAttribute{
								Description:         "uri of content file or git directory",
								MarkdownDescription: "uri of content file or git directory",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"username_secret": schema.SingleNestedAttribute{
								Description:         "SecretRef is the Testkube internal reference for secret storage in Kubernetes secrets",
								MarkdownDescription: "SecretRef is the Testkube internal reference for secret storage in Kubernetes secrets",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "object key",
										MarkdownDescription: "object key",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "object name",
										MarkdownDescription: "object name",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"working_dir": schema.StringAttribute{
								Description:         "if provided we checkout the whole repository and run test from this directory",
								MarkdownDescription: "if provided we checkout the whole repository and run test from this directory",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"type": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("string", "file-uri", "git-file", "git-dir", "git"),
						},
					},

					"uri": schema.StringAttribute{
						Description:         "uri of test content",
						MarkdownDescription: "uri of test content",
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
	}
}

func (r *TestsTestkubeIoTestSourceV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_tests_testkube_io_test_source_v1_manifest")

	var model TestsTestkubeIoTestSourceV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("tests.testkube.io/v1")
	model.Kind = pointer.String("TestSource")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
