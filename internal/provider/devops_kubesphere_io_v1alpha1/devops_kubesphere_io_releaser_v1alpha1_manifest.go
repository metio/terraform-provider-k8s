/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package devops_kubesphere_io_v1alpha1

import (
	"context"
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
	_ datasource.DataSource = &DevopsKubesphereIoReleaserV1Alpha1Manifest{}
)

func NewDevopsKubesphereIoReleaserV1Alpha1Manifest() datasource.DataSource {
	return &DevopsKubesphereIoReleaserV1Alpha1Manifest{}
}

type DevopsKubesphereIoReleaserV1Alpha1Manifest struct{}

type DevopsKubesphereIoReleaserV1Alpha1ManifestData struct {
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
		GitOps *struct {
			Enable     *bool `tfsdk:"enable" json:"enable,omitempty"`
			Repository *struct {
				Action   *string `tfsdk:"action" json:"action,omitempty"`
				Address  *string `tfsdk:"address" json:"address,omitempty"`
				Branch   *string `tfsdk:"branch" json:"branch,omitempty"`
				Message  *string `tfsdk:"message" json:"message,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Provider *string `tfsdk:"provider" json:"provider,omitempty"`
				Version  *string `tfsdk:"version" json:"version,omitempty"`
			} `tfsdk:"repository" json:"repository,omitempty"`
			Secret *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"secret" json:"secret,omitempty"`
		} `tfsdk:"git_ops" json:"gitOps,omitempty"`
		Phase        *string `tfsdk:"phase" json:"phase,omitempty"`
		Repositories *[]struct {
			Action   *string `tfsdk:"action" json:"action,omitempty"`
			Address  *string `tfsdk:"address" json:"address,omitempty"`
			Branch   *string `tfsdk:"branch" json:"branch,omitempty"`
			Message  *string `tfsdk:"message" json:"message,omitempty"`
			Name     *string `tfsdk:"name" json:"name,omitempty"`
			Provider *string `tfsdk:"provider" json:"provider,omitempty"`
			Version  *string `tfsdk:"version" json:"version,omitempty"`
		} `tfsdk:"repositories" json:"repositories,omitempty"`
		Secret *struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"secret" json:"secret,omitempty"`
		Version *string `tfsdk:"version" json:"version,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *DevopsKubesphereIoReleaserV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_devops_kubesphere_io_releaser_v1alpha1_manifest"
}

func (r *DevopsKubesphereIoReleaserV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Releaser is the Schema for the releasers API",
		MarkdownDescription: "Releaser is the Schema for the releasers API",
		Attributes: map[string]schema.Attribute{
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
				Description:         "ReleaserSpec defines the desired state of Releaser",
				MarkdownDescription: "ReleaserSpec defines the desired state of Releaser",
				Attributes: map[string]schema.Attribute{
					"git_ops": schema.SingleNestedAttribute{
						Description:         "GitOps indicates to integrate with GitOps",
						MarkdownDescription: "GitOps indicates to integrate with GitOps",
						Attributes: map[string]schema.Attribute{
							"enable": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"repository": schema.SingleNestedAttribute{
								Description:         "Repository represents a git repository",
								MarkdownDescription: "Repository represents a git repository",
								Attributes: map[string]schema.Attribute{
									"action": schema.StringAttribute{
										Description:         "Action indicates the action once the request phase to be ready",
										MarkdownDescription: "Action indicates the action once the request phase to be ready",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"address": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"branch": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"message": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"provider": schema.StringAttribute{
										Description:         "Provider represents a git provider, such as: GitHub, Gitlab",
										MarkdownDescription: "Provider represents a git provider, such as: GitHub, Gitlab",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"version": schema.StringAttribute{
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

							"secret": schema.SingleNestedAttribute{
								Description:         "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
								MarkdownDescription: "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name is unique within a namespace to reference a secret resource.",
										MarkdownDescription: "Name is unique within a namespace to reference a secret resource.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
										Description:         "Namespace defines the space within which the secret name must be unique.",
										MarkdownDescription: "Namespace defines the space within which the secret name must be unique.",
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

					"phase": schema.StringAttribute{
						Description:         "Phase is the stage of a release request",
						MarkdownDescription: "Phase is the stage of a release request",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"repositories": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"action": schema.StringAttribute{
									Description:         "Action indicates the action once the request phase to be ready",
									MarkdownDescription: "Action indicates the action once the request phase to be ready",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"address": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"branch": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"message": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"provider": schema.StringAttribute{
									Description:         "Provider represents a git provider, such as: GitHub, Gitlab",
									MarkdownDescription: "Provider represents a git provider, such as: GitHub, Gitlab",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"version": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"secret": schema.SingleNestedAttribute{
						Description:         "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
						MarkdownDescription: "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name is unique within a namespace to reference a secret resource.",
								MarkdownDescription: "Name is unique within a namespace to reference a secret resource.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace defines the space within which the secret name must be unique.",
								MarkdownDescription: "Namespace defines the space within which the secret name must be unique.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"version": schema.StringAttribute{
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
		},
	}
}

func (r *DevopsKubesphereIoReleaserV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_devops_kubesphere_io_releaser_v1alpha1_manifest")

	var model DevopsKubesphereIoReleaserV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("devops.kubesphere.io/v1alpha1")
	model.Kind = pointer.String("Releaser")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
