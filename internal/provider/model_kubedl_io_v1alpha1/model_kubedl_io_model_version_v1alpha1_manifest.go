/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package model_kubedl_io_v1alpha1

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
	_ datasource.DataSource = &ModelKubedlIoModelVersionV1Alpha1Manifest{}
)

func NewModelKubedlIoModelVersionV1Alpha1Manifest() datasource.DataSource {
	return &ModelKubedlIoModelVersionV1Alpha1Manifest{}
}

type ModelKubedlIoModelVersionV1Alpha1Manifest struct{}

type ModelKubedlIoModelVersionV1Alpha1ManifestData struct {
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
		CreatedBy *string `tfsdk:"created_by" json:"createdBy,omitempty"`
		ImageRepo *string `tfsdk:"image_repo" json:"imageRepo,omitempty"`
		ImageTag  *string `tfsdk:"image_tag" json:"imageTag,omitempty"`
		ModelName *string `tfsdk:"model_name" json:"modelName,omitempty"`
		Storage   *struct {
			AWSEfs *struct {
				Attributes   *map[string]string `tfsdk:"attributes" json:"attributes,omitempty"`
				VolumeHandle *string            `tfsdk:"volume_handle" json:"volumeHandle,omitempty"`
			} `tfsdk:"aws_efs" json:"AWSEfs,omitempty"`
			LocalStorage *struct {
				MountPath *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
				NodeName  *string `tfsdk:"node_name" json:"nodeName,omitempty"`
				Path      *string `tfsdk:"path" json:"path,omitempty"`
			} `tfsdk:"local_storage" json:"localStorage,omitempty"`
			Nfs *struct {
				MountPath *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
				Path      *string `tfsdk:"path" json:"path,omitempty"`
				Server    *string `tfsdk:"server" json:"server,omitempty"`
			} `tfsdk:"nfs" json:"nfs,omitempty"`
		} `tfsdk:"storage" json:"storage,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ModelKubedlIoModelVersionV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_model_kubedl_io_model_version_v1alpha1_manifest"
}

func (r *ModelKubedlIoModelVersionV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"created_by": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image_repo": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image_tag": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"model_name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"storage": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"aws_efs": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"attributes": schema.MapAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"volume_handle": schema.StringAttribute{
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

							"local_storage": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"mount_path": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"node_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"path": schema.StringAttribute{
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

							"nfs": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"mount_path": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"path": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"server": schema.StringAttribute{
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

func (r *ModelKubedlIoModelVersionV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_model_kubedl_io_model_version_v1alpha1_manifest")

	var model ModelKubedlIoModelVersionV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("model.kubedl.io/v1alpha1")
	model.Kind = pointer.String("ModelVersion")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
