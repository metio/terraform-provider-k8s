/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package operator_tigera_io_v1

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
	_ datasource.DataSource = &OperatorTigeraIoImageSetV1Manifest{}
)

func NewOperatorTigeraIoImageSetV1Manifest() datasource.DataSource {
	return &OperatorTigeraIoImageSetV1Manifest{}
}

type OperatorTigeraIoImageSetV1Manifest struct{}

type OperatorTigeraIoImageSetV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Images *[]struct {
			Digest *string `tfsdk:"digest" json:"digest,omitempty"`
			Image  *string `tfsdk:"image" json:"image,omitempty"`
		} `tfsdk:"images" json:"images,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OperatorTigeraIoImageSetV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_operator_tigera_io_image_set_v1_manifest"
}

func (r *OperatorTigeraIoImageSetV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ImageSet is used to specify image digests for the images that the operator deploys. The name of the ImageSet is expected to be in the format '<variant>-<release>'. The 'variant' used is 'enterprise' if the InstallationSpec Variant is 'TigeraSecureEnterprise' otherwise it is 'calico'. The 'release' must match the version of the variant that the operator is built to deploy, this version can be obtained by passing the '--version' flag to the operator binary.",
		MarkdownDescription: "ImageSet is used to specify image digests for the images that the operator deploys. The name of the ImageSet is expected to be in the format '<variant>-<release>'. The 'variant' used is 'enterprise' if the InstallationSpec Variant is 'TigeraSecureEnterprise' otherwise it is 'calico'. The 'release' must match the version of the variant that the operator is built to deploy, this version can be obtained by passing the '--version' flag to the operator binary.",
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
				Description:         "ImageSetSpec defines the desired state of ImageSet.",
				MarkdownDescription: "ImageSetSpec defines the desired state of ImageSet.",
				Attributes: map[string]schema.Attribute{
					"images": schema.ListNestedAttribute{
						Description:         "Images is the list of images to use digests. All images that the operator will deploy must be specified.",
						MarkdownDescription: "Images is the list of images to use digests. All images that the operator will deploy must be specified.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"digest": schema.StringAttribute{
									Description:         "Digest is the image identifier that will be used for the Image. The field should not include a leading '@' and must be prefixed with 'sha256:'.",
									MarkdownDescription: "Digest is the image identifier that will be used for the Image. The field should not include a leading '@' and must be prefixed with 'sha256:'.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"image": schema.StringAttribute{
									Description:         "Image is an image that the operator deploys and instead of using the built in tag the operator will use the Digest for the image identifier. The value should be the image name without registry or tag or digest. For the image 'docker.io/calico/node:v3.17.1' it should be represented as 'calico/node'",
									MarkdownDescription: "Image is an image that the operator deploys and instead of using the built in tag the operator will use the Digest for the image identifier. The value should be the image name without registry or tag or digest. For the image 'docker.io/calico/node:v3.17.1' it should be represented as 'calico/node'",
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

func (r *OperatorTigeraIoImageSetV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_operator_tigera_io_image_set_v1_manifest")

	var model OperatorTigeraIoImageSetV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("operator.tigera.io/v1")
	model.Kind = pointer.String("ImageSet")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
