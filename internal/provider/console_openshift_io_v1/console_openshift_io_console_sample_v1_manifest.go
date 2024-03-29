/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package console_openshift_io_v1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &ConsoleOpenshiftIoConsoleSampleV1Manifest{}
)

func NewConsoleOpenshiftIoConsoleSampleV1Manifest() datasource.DataSource {
	return &ConsoleOpenshiftIoConsoleSampleV1Manifest{}
}

type ConsoleOpenshiftIoConsoleSampleV1Manifest struct{}

type ConsoleOpenshiftIoConsoleSampleV1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Abstract    *string `tfsdk:"abstract" json:"abstract,omitempty"`
		Description *string `tfsdk:"description" json:"description,omitempty"`
		Icon        *string `tfsdk:"icon" json:"icon,omitempty"`
		Provider    *string `tfsdk:"provider" json:"provider,omitempty"`
		Source      *struct {
			ContainerImport *struct {
				Image   *string `tfsdk:"image" json:"image,omitempty"`
				Service *struct {
					TargetPort *int64 `tfsdk:"target_port" json:"targetPort,omitempty"`
				} `tfsdk:"service" json:"service,omitempty"`
			} `tfsdk:"container_import" json:"containerImport,omitempty"`
			GitImport *struct {
				Repository *struct {
					ContextDir *string `tfsdk:"context_dir" json:"contextDir,omitempty"`
					Revision   *string `tfsdk:"revision" json:"revision,omitempty"`
					Url        *string `tfsdk:"url" json:"url,omitempty"`
				} `tfsdk:"repository" json:"repository,omitempty"`
				Service *struct {
					TargetPort *int64 `tfsdk:"target_port" json:"targetPort,omitempty"`
				} `tfsdk:"service" json:"service,omitempty"`
			} `tfsdk:"git_import" json:"gitImport,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"source" json:"source,omitempty"`
		Tags  *[]string `tfsdk:"tags" json:"tags,omitempty"`
		Title *string   `tfsdk:"title" json:"title,omitempty"`
		Type  *string   `tfsdk:"type" json:"type,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ConsoleOpenshiftIoConsoleSampleV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_console_openshift_io_console_sample_v1_manifest"
}

func (r *ConsoleOpenshiftIoConsoleSampleV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ConsoleSample is an extension to customizing OpenShift web console by adding samples.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		MarkdownDescription: "ConsoleSample is an extension to customizing OpenShift web console by adding samples.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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
				Description:         "spec contains configuration for a console sample.",
				MarkdownDescription: "spec contains configuration for a console sample.",
				Attributes: map[string]schema.Attribute{
					"abstract": schema.StringAttribute{
						Description:         "abstract is a short introduction to the sample.  It is required and must be no more than 100 characters in length.  The abstract is shown on the sample card tile below the title and provider and is limited to three lines of content.",
						MarkdownDescription: "abstract is a short introduction to the sample.  It is required and must be no more than 100 characters in length.  The abstract is shown on the sample card tile below the title and provider and is limited to three lines of content.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(100),
						},
					},

					"description": schema.StringAttribute{
						Description:         "description is a long form explanation of the sample.  It is required and can have a maximum length of **4096** characters.  It is a README.md-like content for additional information, links, pre-conditions, and other instructions. It will be rendered as Markdown so that it can contain line breaks, links, and other simple formatting.",
						MarkdownDescription: "description is a long form explanation of the sample.  It is required and can have a maximum length of **4096** characters.  It is a README.md-like content for additional information, links, pre-conditions, and other instructions. It will be rendered as Markdown so that it can contain line breaks, links, and other simple formatting.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(4096),
						},
					},

					"icon": schema.StringAttribute{
						Description:         "icon is an optional base64 encoded image and shown beside the sample title.  The format must follow the data: URL format and can have a maximum size of **10 KB**.  data:[<mediatype>][;base64],<base64 encoded image>  For example:  data:image;base64,             plus the base64 encoded image.  Vector images can also be used. SVG icons must start with:  data:image/svg+xml;base64,     plus the base64 encoded SVG image.  All sample catalog icons will be shown on a white background (also when the dark theme is used). The web console ensures that different aspect ratios work correctly. Currently, the surface of the icon is at most 40x100px.  For more information on the data URL format, please visit https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/Data_URLs.",
						MarkdownDescription: "icon is an optional base64 encoded image and shown beside the sample title.  The format must follow the data: URL format and can have a maximum size of **10 KB**.  data:[<mediatype>][;base64],<base64 encoded image>  For example:  data:image;base64,             plus the base64 encoded image.  Vector images can also be used. SVG icons must start with:  data:image/svg+xml;base64,     plus the base64 encoded SVG image.  All sample catalog icons will be shown on a white background (also when the dark theme is used). The web console ensures that different aspect ratios work correctly. Currently, the surface of the icon is at most 40x100px.  For more information on the data URL format, please visit https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/Data_URLs.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(14000),
							stringvalidator.RegexMatches(regexp.MustCompile(`^data:([a-z/\.+0-9]*;(([-a-zA-Z0-9=])*;)?)?base64,`), ""),
						},
					},

					"provider": schema.StringAttribute{
						Description:         "provider is an optional label to honor who provides the sample.  It is optional and must be no more than 50 characters in length.  A provider can be a company like 'Red Hat' or an organization like 'CNCF' or 'Knative'.  Currently, the provider is only shown on the sample card tile below the title with the prefix 'Provided by '",
						MarkdownDescription: "provider is an optional label to honor who provides the sample.  It is optional and must be no more than 50 characters in length.  A provider can be a company like 'Red Hat' or an organization like 'CNCF' or 'Knative'.  Currently, the provider is only shown on the sample card tile below the title with the prefix 'Provided by '",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(50),
						},
					},

					"source": schema.SingleNestedAttribute{
						Description:         "source defines where to deploy the sample service from. The sample may be sourced from an external git repository or container image.",
						MarkdownDescription: "source defines where to deploy the sample service from. The sample may be sourced from an external git repository or container image.",
						Attributes: map[string]schema.Attribute{
							"container_import": schema.SingleNestedAttribute{
								Description:         "containerImport allows the user import a container image.",
								MarkdownDescription: "containerImport allows the user import a container image.",
								Attributes: map[string]schema.Attribute{
									"image": schema.StringAttribute{
										Description:         "reference to a container image that provides a HTTP service. The service must be exposed on the default port (8080) unless otherwise configured with the port field.  Supported formats: - <repository-name>/<image-name> - docker.io/<repository-name>/<image-name> - quay.io/<repository-name>/<image-name> - quay.io/<repository-name>/<image-name>@sha256:<image hash> - quay.io/<repository-name>/<image-name>:<tag>",
										MarkdownDescription: "reference to a container image that provides a HTTP service. The service must be exposed on the default port (8080) unless otherwise configured with the port field.  Supported formats: - <repository-name>/<image-name> - docker.io/<repository-name>/<image-name> - quay.io/<repository-name>/<image-name> - quay.io/<repository-name>/<image-name>@sha256:<image hash> - quay.io/<repository-name>/<image-name>:<tag>",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(1),
											stringvalidator.LengthAtMost(256),
										},
									},

									"service": schema.SingleNestedAttribute{
										Description:         "service contains configuration for the Service resource created for this sample.",
										MarkdownDescription: "service contains configuration for the Service resource created for this sample.",
										Attributes: map[string]schema.Attribute{
											"target_port": schema.Int64Attribute{
												Description:         "targetPort is the port that the service listens on for HTTP requests. This port will be used for Service and Route created for this sample. Port must be in the range 1 to 65535. Default port is 8080.",
												MarkdownDescription: "targetPort is the port that the service listens on for HTTP requests. This port will be used for Service and Route created for this sample. Port must be in the range 1 to 65535. Default port is 8080.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
													int64validator.AtMost(65535),
												},
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

							"git_import": schema.SingleNestedAttribute{
								Description:         "gitImport allows the user to import code from a git repository.",
								MarkdownDescription: "gitImport allows the user to import code from a git repository.",
								Attributes: map[string]schema.Attribute{
									"repository": schema.SingleNestedAttribute{
										Description:         "repository contains the reference to the actual Git repository.",
										MarkdownDescription: "repository contains the reference to the actual Git repository.",
										Attributes: map[string]schema.Attribute{
											"context_dir": schema.StringAttribute{
												Description:         "contextDir is used to specify a directory within the repository to build the component. Must start with '/' and have a maximum length of 256 characters. When omitted, the default value is to build from the root of the repository.",
												MarkdownDescription: "contextDir is used to specify a directory within the repository to build the component. Must start with '/' and have a maximum length of 256 characters. When omitted, the default value is to build from the root of the repository.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtMost(256),
													stringvalidator.RegexMatches(regexp.MustCompile(`^/`), ""),
												},
											},

											"revision": schema.StringAttribute{
												Description:         "revision is the git revision at which to clone the git repository Can be used to clone a specific branch, tag or commit SHA. Must be at most 256 characters in length. When omitted the repository's default branch is used.",
												MarkdownDescription: "revision is the git revision at which to clone the git repository Can be used to clone a specific branch, tag or commit SHA. Must be at most 256 characters in length. When omitted the repository's default branch is used.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtMost(256),
												},
											},

											"url": schema.StringAttribute{
												Description:         "url of the Git repository that contains a HTTP service. The HTTP service must be exposed on the default port (8080) unless otherwise configured with the port field.  Only public repositories on GitHub, GitLab and Bitbucket are currently supported:  - https://github.com/<org>/<repository> - https://gitlab.com/<org>/<repository> - https://bitbucket.org/<org>/<repository>  The url must have a maximum length of 256 characters.",
												MarkdownDescription: "url of the Git repository that contains a HTTP service. The HTTP service must be exposed on the default port (8080) unless otherwise configured with the port field.  Only public repositories on GitHub, GitLab and Bitbucket are currently supported:  - https://github.com/<org>/<repository> - https://gitlab.com/<org>/<repository> - https://bitbucket.org/<org>/<repository>  The url must have a maximum length of 256 characters.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
													stringvalidator.LengthAtMost(256),
													stringvalidator.RegexMatches(regexp.MustCompile(`^https:\/\/(github.com|gitlab.com|bitbucket.org)\/[a-zA-Z0-9-]+\/[a-zA-Z0-9-]+(.git)?$`), ""),
												},
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"service": schema.SingleNestedAttribute{
										Description:         "service contains configuration for the Service resource created for this sample.",
										MarkdownDescription: "service contains configuration for the Service resource created for this sample.",
										Attributes: map[string]schema.Attribute{
											"target_port": schema.Int64Attribute{
												Description:         "targetPort is the port that the service listens on for HTTP requests. This port will be used for Service created for this sample. Port must be in the range 1 to 65535. Default port is 8080.",
												MarkdownDescription: "targetPort is the port that the service listens on for HTTP requests. This port will be used for Service created for this sample. Port must be in the range 1 to 65535. Default port is 8080.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
													int64validator.AtMost(65535),
												},
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

							"type": schema.StringAttribute{
								Description:         "type of the sample, currently supported: 'GitImport';'ContainerImport'",
								MarkdownDescription: "type of the sample, currently supported: 'GitImport';'ContainerImport'",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"tags": schema.ListAttribute{
						Description:         "tags are optional string values that can be used to find samples in the samples catalog.  Examples of common tags may be 'Java', 'Quarkus', etc.  They will be displayed on the samples details page.",
						MarkdownDescription: "tags are optional string values that can be used to find samples in the samples catalog.  Examples of common tags may be 'Java', 'Quarkus', etc.  They will be displayed on the samples details page.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"title": schema.StringAttribute{
						Description:         "title is the display name of the sample.  It is required and must be no more than 50 characters in length.",
						MarkdownDescription: "title is the display name of the sample.  It is required and must be no more than 50 characters in length.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
							stringvalidator.LengthAtMost(50),
						},
					},

					"type": schema.StringAttribute{
						Description:         "type is an optional label to group multiple samples.  It is optional and must be no more than 20 characters in length.  Recommendation is a singular term like 'Builder Image', 'Devfile' or 'Serverless Function'.  Currently, the type is shown a badge on the sample card tile in the top right corner.",
						MarkdownDescription: "type is an optional label to group multiple samples.  It is optional and must be no more than 20 characters in length.  Recommendation is a singular term like 'Builder Image', 'Devfile' or 'Serverless Function'.  Currently, the type is shown a badge on the sample card tile in the top right corner.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(20),
						},
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *ConsoleOpenshiftIoConsoleSampleV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_console_openshift_io_console_sample_v1_manifest")

	var model ConsoleOpenshiftIoConsoleSampleV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("console.openshift.io/v1")
	model.Kind = pointer.String("ConsoleSample")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
