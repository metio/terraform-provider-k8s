/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kuma_io_v1alpha1

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
	_ datasource.DataSource = &KumaIoContainerPatchV1Alpha1Manifest{}
)

func NewKumaIoContainerPatchV1Alpha1Manifest() datasource.DataSource {
	return &KumaIoContainerPatchV1Alpha1Manifest{}
}

type KumaIoContainerPatchV1Alpha1Manifest struct{}

type KumaIoContainerPatchV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Mesh *string `tfsdk:"mesh" json:"mesh,omitempty"`
	Spec *struct {
		InitPatch *[]struct {
			From  *string `tfsdk:"from" json:"from,omitempty"`
			Op    *string `tfsdk:"op" json:"op,omitempty"`
			Path  *string `tfsdk:"path" json:"path,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"init_patch" json:"initPatch,omitempty"`
		SidecarPatch *[]struct {
			From  *string `tfsdk:"from" json:"from,omitempty"`
			Op    *string `tfsdk:"op" json:"op,omitempty"`
			Path  *string `tfsdk:"path" json:"path,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"sidecar_patch" json:"sidecarPatch,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KumaIoContainerPatchV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kuma_io_container_patch_v1alpha1_manifest"
}

func (r *KumaIoContainerPatchV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ContainerPatch stores a list of patches to apply to init and sidecar containers.",
		MarkdownDescription: "ContainerPatch stores a list of patches to apply to init and sidecar containers.",
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

			"mesh": schema.StringAttribute{
				Description:         "",
				MarkdownDescription: "",
				Required:            false,
				Optional:            true,
				Computed:            false,
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "ContainerPatchSpec specifies the options available for a ContainerPatch",
				MarkdownDescription: "ContainerPatchSpec specifies the options available for a ContainerPatch",
				Attributes: map[string]schema.Attribute{
					"init_patch": schema.ListNestedAttribute{
						Description:         "InitPatch specifies jsonpatch to apply to an init container.",
						MarkdownDescription: "InitPatch specifies jsonpatch to apply to an init container.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"from": schema.StringAttribute{
									Description:         "From is a jsonpatch from string, used by move and copy operations.",
									MarkdownDescription: "From is a jsonpatch from string, used by move and copy operations.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"op": schema.StringAttribute{
									Description:         "Op is a jsonpatch operation string.",
									MarkdownDescription: "Op is a jsonpatch operation string.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("add", "remove", "replace", "move", "copy"),
									},
								},

								"path": schema.StringAttribute{
									Description:         "Path is a jsonpatch path string.",
									MarkdownDescription: "Path is a jsonpatch path string.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"value": schema.StringAttribute{
									Description:         "Value must be a string representing a valid json object usedby replace and add operations. String has to be escaped with ' to be valid a json object.",
									MarkdownDescription: "Value must be a string representing a valid json object usedby replace and add operations. String has to be escaped with ' to be valid a json object.",
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

					"sidecar_patch": schema.ListNestedAttribute{
						Description:         "SidecarPatch specifies jsonpatch to apply to a sidecar container.",
						MarkdownDescription: "SidecarPatch specifies jsonpatch to apply to a sidecar container.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"from": schema.StringAttribute{
									Description:         "From is a jsonpatch from string, used by move and copy operations.",
									MarkdownDescription: "From is a jsonpatch from string, used by move and copy operations.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"op": schema.StringAttribute{
									Description:         "Op is a jsonpatch operation string.",
									MarkdownDescription: "Op is a jsonpatch operation string.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("add", "remove", "replace", "move", "copy"),
									},
								},

								"path": schema.StringAttribute{
									Description:         "Path is a jsonpatch path string.",
									MarkdownDescription: "Path is a jsonpatch path string.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"value": schema.StringAttribute{
									Description:         "Value must be a string representing a valid json object usedby replace and add operations. String has to be escaped with ' to be valid a json object.",
									MarkdownDescription: "Value must be a string representing a valid json object usedby replace and add operations. String has to be escaped with ' to be valid a json object.",
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *KumaIoContainerPatchV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kuma_io_container_patch_v1alpha1_manifest")

	var model KumaIoContainerPatchV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("kuma.io/v1alpha1")
	model.Kind = pointer.String("ContainerPatch")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
