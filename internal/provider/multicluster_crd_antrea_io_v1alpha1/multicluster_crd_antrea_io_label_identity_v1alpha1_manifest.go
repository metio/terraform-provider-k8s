/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package multicluster_crd_antrea_io_v1alpha1

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
	_ datasource.DataSource = &MulticlusterCrdAntreaIoLabelIdentityV1Alpha1Manifest{}
)

func NewMulticlusterCrdAntreaIoLabelIdentityV1Alpha1Manifest() datasource.DataSource {
	return &MulticlusterCrdAntreaIoLabelIdentityV1Alpha1Manifest{}
}

type MulticlusterCrdAntreaIoLabelIdentityV1Alpha1Manifest struct{}

type MulticlusterCrdAntreaIoLabelIdentityV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Id    *int64  `tfsdk:"id" json:"id,omitempty"`
		Label *string `tfsdk:"label" json:"label,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *MulticlusterCrdAntreaIoLabelIdentityV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_multicluster_crd_antrea_io_label_identity_v1alpha1_manifest"
}

func (r *MulticlusterCrdAntreaIoLabelIdentityV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "LabelIdentity is an imported label identity from the ClusterSet. For each unique label identity, a LabelIdentity will be created in the member cluster.",
		MarkdownDescription: "LabelIdentity is an imported label identity from the ClusterSet. For each unique label identity, a LabelIdentity will be created in the member cluster.",
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
					"id": schema.Int64Attribute{
						Description:         "ID is the ID allocated for the label identity by the leader cluster.",
						MarkdownDescription: "ID is the ID allocated for the label identity by the leader cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"label": schema.StringAttribute{
						Description:         "Label is the normalized string of a label identity. The format of normalized label identity is 'ns:(?P<nslabels>(.)*)&pod:(?P<podlabels>(.)*)' E.g., 'ns:kubernetes.io/metadata.name=kube-system&pod:app=db'",
						MarkdownDescription: "Label is the normalized string of a label identity. The format of normalized label identity is 'ns:(?P<nslabels>(.)*)&pod:(?P<podlabels>(.)*)' E.g., 'ns:kubernetes.io/metadata.name=kube-system&pod:app=db'",
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

func (r *MulticlusterCrdAntreaIoLabelIdentityV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_multicluster_crd_antrea_io_label_identity_v1alpha1_manifest")

	var model MulticlusterCrdAntreaIoLabelIdentityV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("multicluster.crd.antrea.io/v1alpha1")
	model.Kind = pointer.String("LabelIdentity")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
