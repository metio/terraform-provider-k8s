/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package operator_tigera_io_v1

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
	_ datasource.DataSource              = &OperatorTigeraIoImageSetV1DataSource{}
	_ datasource.DataSourceWithConfigure = &OperatorTigeraIoImageSetV1DataSource{}
)

func NewOperatorTigeraIoImageSetV1DataSource() datasource.DataSource {
	return &OperatorTigeraIoImageSetV1DataSource{}
}

type OperatorTigeraIoImageSetV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type OperatorTigeraIoImageSetV1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

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

func (r *OperatorTigeraIoImageSetV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_operator_tigera_io_image_set_v1"
}

func (r *OperatorTigeraIoImageSetV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ImageSet is used to specify image digests for the images that the operator deploys. The name of the ImageSet is expected to be in the format '<variant>-<release>'. The 'variant' used is 'enterprise' if the InstallationSpec Variant is 'TigeraSecureEnterprise' otherwise it is 'calico'. The 'release' must match the version of the variant that the operator is built to deploy, this version can be obtained by passing the '--version' flag to the operator binary.",
		MarkdownDescription: "ImageSet is used to specify image digests for the images that the operator deploys. The name of the ImageSet is expected to be in the format '<variant>-<release>'. The 'variant' used is 'enterprise' if the InstallationSpec Variant is 'TigeraSecureEnterprise' otherwise it is 'calico'. The 'release' must match the version of the variant that the operator is built to deploy, this version can be obtained by passing the '--version' flag to the operator binary.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"image": schema.StringAttribute{
									Description:         "Image is an image that the operator deploys and instead of using the built in tag the operator will use the Digest for the image identifier. The value should be the image name without registry or tag or digest. For the image 'docker.io/calico/node:v3.17.1' it should be represented as 'calico/node'",
									MarkdownDescription: "Image is an image that the operator deploys and instead of using the built in tag the operator will use the Digest for the image identifier. The value should be the image name without registry or tag or digest. For the image 'docker.io/calico/node:v3.17.1' it should be represented as 'calico/node'",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},
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

func (r *OperatorTigeraIoImageSetV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *OperatorTigeraIoImageSetV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_operator_tigera_io_image_set_v1")

	var data OperatorTigeraIoImageSetV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "operator.tigera.io", Version: "v1", Resource: "imagesets"}).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetResourceError(err, data.Metadata.Name))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse OperatorTigeraIoImageSetV1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(data.Metadata.Name)
	data.ApiVersion = pointer.String("operator.tigera.io/v1")
	data.Kind = pointer.String("ImageSet")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
