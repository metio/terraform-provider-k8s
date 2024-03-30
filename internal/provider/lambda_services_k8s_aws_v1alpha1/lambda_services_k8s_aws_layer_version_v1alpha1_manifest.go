/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package lambda_services_k8s_aws_v1alpha1

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
	_ datasource.DataSource = &LambdaServicesK8SAwsLayerVersionV1Alpha1Manifest{}
)

func NewLambdaServicesK8SAwsLayerVersionV1Alpha1Manifest() datasource.DataSource {
	return &LambdaServicesK8SAwsLayerVersionV1Alpha1Manifest{}
}

type LambdaServicesK8SAwsLayerVersionV1Alpha1Manifest struct{}

type LambdaServicesK8SAwsLayerVersionV1Alpha1ManifestData struct {
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
		CompatibleArchitectures *[]string `tfsdk:"compatible_architectures" json:"compatibleArchitectures,omitempty"`
		CompatibleRuntimes      *[]string `tfsdk:"compatible_runtimes" json:"compatibleRuntimes,omitempty"`
		Content                 *struct {
			S3Bucket        *string `tfsdk:"s3_bucket" json:"s3Bucket,omitempty"`
			S3Key           *string `tfsdk:"s3_key" json:"s3Key,omitempty"`
			S3ObjectVersion *string `tfsdk:"s3_object_version" json:"s3ObjectVersion,omitempty"`
			ZipFile         *string `tfsdk:"zip_file" json:"zipFile,omitempty"`
		} `tfsdk:"content" json:"content,omitempty"`
		Description *string `tfsdk:"description" json:"description,omitempty"`
		LayerName   *string `tfsdk:"layer_name" json:"layerName,omitempty"`
		LicenseInfo *string `tfsdk:"license_info" json:"licenseInfo,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *LambdaServicesK8SAwsLayerVersionV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_lambda_services_k8s_aws_layer_version_v1alpha1_manifest"
}

func (r *LambdaServicesK8SAwsLayerVersionV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "LayerVersion is the Schema for the LayerVersions API",
		MarkdownDescription: "LayerVersion is the Schema for the LayerVersions API",
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
				Description:         "LayerVersionSpec defines the desired state of LayerVersion.",
				MarkdownDescription: "LayerVersionSpec defines the desired state of LayerVersion.",
				Attributes: map[string]schema.Attribute{
					"compatible_architectures": schema.ListAttribute{
						Description:         "A list of compatible instruction set architectures (https://docs.aws.amazon.com/lambda/latest/dg/foundation-arch.html).",
						MarkdownDescription: "A list of compatible instruction set architectures (https://docs.aws.amazon.com/lambda/latest/dg/foundation-arch.html).",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"compatible_runtimes": schema.ListAttribute{
						Description:         "A list of compatible function runtimes (https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html).Used for filtering with ListLayers and ListLayerVersions.",
						MarkdownDescription: "A list of compatible function runtimes (https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html).Used for filtering with ListLayers and ListLayerVersions.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"content": schema.SingleNestedAttribute{
						Description:         "The function layer archive.",
						MarkdownDescription: "The function layer archive.",
						Attributes: map[string]schema.Attribute{
							"s3_bucket": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"s3_key": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"s3_object_version": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"zip_file": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									validators.Base64Validator(),
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"description": schema.StringAttribute{
						Description:         "The description of the version.",
						MarkdownDescription: "The description of the version.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"layer_name": schema.StringAttribute{
						Description:         "The name or Amazon Resource Name (ARN) of the layer.",
						MarkdownDescription: "The name or Amazon Resource Name (ARN) of the layer.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"license_info": schema.StringAttribute{
						Description:         "The layer's software license. It can be any of the following:   * An SPDX license identifier (https://spdx.org/licenses/). For example,   MIT.   * The URL of a license hosted on the internet. For example, https://opensource.org/licenses/MIT.   * The full text of the license.",
						MarkdownDescription: "The layer's software license. It can be any of the following:   * An SPDX license identifier (https://spdx.org/licenses/). For example,   MIT.   * The URL of a license hosted on the internet. For example, https://opensource.org/licenses/MIT.   * The full text of the license.",
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

func (r *LambdaServicesK8SAwsLayerVersionV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_lambda_services_k8s_aws_layer_version_v1alpha1_manifest")

	var model LambdaServicesK8SAwsLayerVersionV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("lambda.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("LayerVersion")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
