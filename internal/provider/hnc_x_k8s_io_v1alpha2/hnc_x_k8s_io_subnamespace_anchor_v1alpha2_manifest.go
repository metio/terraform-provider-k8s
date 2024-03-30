/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package hnc_x_k8s_io_v1alpha2

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
	_ datasource.DataSource = &HncXK8SIoSubnamespaceAnchorV1Alpha2Manifest{}
)

func NewHncXK8SIoSubnamespaceAnchorV1Alpha2Manifest() datasource.DataSource {
	return &HncXK8SIoSubnamespaceAnchorV1Alpha2Manifest{}
}

type HncXK8SIoSubnamespaceAnchorV1Alpha2Manifest struct{}

type HncXK8SIoSubnamespaceAnchorV1Alpha2ManifestData struct {
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
		Annotations *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"annotations" json:"annotations,omitempty"`
		Labels *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"labels" json:"labels,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *HncXK8SIoSubnamespaceAnchorV1Alpha2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_hnc_x_k8s_io_subnamespace_anchor_v1alpha2_manifest"
}

func (r *HncXK8SIoSubnamespaceAnchorV1Alpha2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "SubnamespaceAnchor is the Schema for the subnamespace API. See details at http://bit.ly/hnc-self-serve-ux.",
		MarkdownDescription: "SubnamespaceAnchor is the Schema for the subnamespace API. See details at http://bit.ly/hnc-self-serve-ux.",
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
					"annotations": schema.ListNestedAttribute{
						Description:         "Annotations is a list of annotations and values to apply to the current subnamespace and all of its descendants. All annotation keys must match a regex specified on the command line by --managed-namespace-annotation. All annotation keys must be managed annotations (see HNC docs) and must match a regex",
						MarkdownDescription: "Annotations is a list of annotations and values to apply to the current subnamespace and all of its descendants. All annotation keys must match a regex specified on the command line by --managed-namespace-annotation. All annotation keys must be managed annotations (see HNC docs) and must match a regex",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"key": schema.StringAttribute{
									Description:         "Key is the name of the label or annotation. It must conform to the normal rules for Kubernetes label/annotation keys.",
									MarkdownDescription: "Key is the name of the label or annotation. It must conform to the normal rules for Kubernetes label/annotation keys.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"value": schema.StringAttribute{
									Description:         "Value is the value of the label or annotation. It must confirm to the normal rules for Kubernetes label or annoation values, which are far more restrictive for labels than for anntations.",
									MarkdownDescription: "Value is the value of the label or annotation. It must confirm to the normal rules for Kubernetes label or annoation values, which are far more restrictive for labels than for anntations.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"labels": schema.ListNestedAttribute{
						Description:         "Labels is a list of labels and values to apply to the current subnamespace and all of its descendants. All label keys must match a regex specified on the command line by --managed-namespace-label. All label keys must be managed labels (see HNC docs) and must match a regex",
						MarkdownDescription: "Labels is a list of labels and values to apply to the current subnamespace and all of its descendants. All label keys must match a regex specified on the command line by --managed-namespace-label. All label keys must be managed labels (see HNC docs) and must match a regex",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"key": schema.StringAttribute{
									Description:         "Key is the name of the label or annotation. It must conform to the normal rules for Kubernetes label/annotation keys.",
									MarkdownDescription: "Key is the name of the label or annotation. It must conform to the normal rules for Kubernetes label/annotation keys.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"value": schema.StringAttribute{
									Description:         "Value is the value of the label or annotation. It must confirm to the normal rules for Kubernetes label or annoation values, which are far more restrictive for labels than for anntations.",
									MarkdownDescription: "Value is the value of the label or annotation. It must confirm to the normal rules for Kubernetes label or annoation values, which are far more restrictive for labels than for anntations.",
									Required:            true,
									Optional:            false,
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

func (r *HncXK8SIoSubnamespaceAnchorV1Alpha2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_hnc_x_k8s_io_subnamespace_anchor_v1alpha2_manifest")

	var model HncXK8SIoSubnamespaceAnchorV1Alpha2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("hnc.x-k8s.io/v1alpha2")
	model.Kind = pointer.String("SubnamespaceAnchor")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
