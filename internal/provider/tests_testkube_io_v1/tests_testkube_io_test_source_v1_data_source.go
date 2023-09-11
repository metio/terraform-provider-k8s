/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package tests_testkube_io_v1

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
	_ datasource.DataSource              = &TestsTestkubeIoTestSourceV1DataSource{}
	_ datasource.DataSourceWithConfigure = &TestsTestkubeIoTestSourceV1DataSource{}
)

func NewTestsTestkubeIoTestSourceV1DataSource() datasource.DataSource {
	return &TestsTestkubeIoTestSourceV1DataSource{}
}

type TestsTestkubeIoTestSourceV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type TestsTestkubeIoTestSourceV1DataSourceData struct {
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

func (r *TestsTestkubeIoTestSourceV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_tests_testkube_io_test_source_v1"
}

func (r *TestsTestkubeIoTestSourceV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "TestSourceSpec defines the desired state of TestSource",
				MarkdownDescription: "TestSourceSpec defines the desired state of TestSource",
				Attributes: map[string]schema.Attribute{
					"data": schema.StringAttribute{
						Description:         "test content body",
						MarkdownDescription: "test content body",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"repository": schema.SingleNestedAttribute{
						Description:         "repository of test content",
						MarkdownDescription: "repository of test content",
						Attributes: map[string]schema.Attribute{
							"auth_type": schema.StringAttribute{
								Description:         "auth type for git requests",
								MarkdownDescription: "auth type for git requests",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"branch": schema.StringAttribute{
								Description:         "branch/tag name for checkout",
								MarkdownDescription: "branch/tag name for checkout",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"certificate_secret": schema.StringAttribute{
								Description:         "git auth certificate secret for private repositories",
								MarkdownDescription: "git auth certificate secret for private repositories",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"commit": schema.StringAttribute{
								Description:         "commit id (sha) for checkout",
								MarkdownDescription: "commit id (sha) for checkout",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"path": schema.StringAttribute{
								Description:         "if needed we can checkout particular path (dir or file) in case of BIG/mono repositories",
								MarkdownDescription: "if needed we can checkout particular path (dir or file) in case of BIG/mono repositories",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"token_secret": schema.SingleNestedAttribute{
								Description:         "Testkube internal reference for secret storage in Kubernetes secrets",
								MarkdownDescription: "Testkube internal reference for secret storage in Kubernetes secrets",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "object key",
										MarkdownDescription: "object key",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"name": schema.StringAttribute{
										Description:         "object name",
										MarkdownDescription: "object name",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"type": schema.StringAttribute{
								Description:         "VCS repository type",
								MarkdownDescription: "VCS repository type",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"uri": schema.StringAttribute{
								Description:         "uri of content file or git directory",
								MarkdownDescription: "uri of content file or git directory",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"username_secret": schema.SingleNestedAttribute{
								Description:         "Testkube internal reference for secret storage in Kubernetes secrets",
								MarkdownDescription: "Testkube internal reference for secret storage in Kubernetes secrets",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "object key",
										MarkdownDescription: "object key",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"name": schema.StringAttribute{
										Description:         "object name",
										MarkdownDescription: "object name",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"working_dir": schema.StringAttribute{
								Description:         "if provided we checkout the whole repository and run test from this directory",
								MarkdownDescription: "if provided we checkout the whole repository and run test from this directory",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"type": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"uri": schema.StringAttribute{
						Description:         "uri of test content",
						MarkdownDescription: "uri of test content",
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

func (r *TestsTestkubeIoTestSourceV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *TestsTestkubeIoTestSourceV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_tests_testkube_io_test_source_v1")

	var data TestsTestkubeIoTestSourceV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "tests.testkube.io", Version: "v1", Resource: "testsources"}).
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

	var readResponse TestsTestkubeIoTestSourceV1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("tests.testkube.io/v1")
	data.Kind = pointer.String("TestSource")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
