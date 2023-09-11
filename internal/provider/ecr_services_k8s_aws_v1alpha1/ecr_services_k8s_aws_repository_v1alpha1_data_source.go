/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package ecr_services_k8s_aws_v1alpha1

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
	_ datasource.DataSource              = &EcrServicesK8SAwsRepositoryV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &EcrServicesK8SAwsRepositoryV1Alpha1DataSource{}
)

func NewEcrServicesK8SAwsRepositoryV1Alpha1DataSource() datasource.DataSource {
	return &EcrServicesK8SAwsRepositoryV1Alpha1DataSource{}
}

type EcrServicesK8SAwsRepositoryV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type EcrServicesK8SAwsRepositoryV1Alpha1DataSourceData struct {
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
		EncryptionConfiguration *struct {
			EncryptionType *string `tfsdk:"encryption_type" json:"encryptionType,omitempty"`
			KmsKey         *string `tfsdk:"kms_key" json:"kmsKey,omitempty"`
		} `tfsdk:"encryption_configuration" json:"encryptionConfiguration,omitempty"`
		ImageScanningConfiguration *struct {
			ScanOnPush *bool `tfsdk:"scan_on_push" json:"scanOnPush,omitempty"`
		} `tfsdk:"image_scanning_configuration" json:"imageScanningConfiguration,omitempty"`
		ImageTagMutability *string `tfsdk:"image_tag_mutability" json:"imageTagMutability,omitempty"`
		LifecyclePolicy    *string `tfsdk:"lifecycle_policy" json:"lifecyclePolicy,omitempty"`
		Name               *string `tfsdk:"name" json:"name,omitempty"`
		Policy             *string `tfsdk:"policy" json:"policy,omitempty"`
		RegistryID         *string `tfsdk:"registry_id" json:"registryID,omitempty"`
		Tags               *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *EcrServicesK8SAwsRepositoryV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_ecr_services_k8s_aws_repository_v1alpha1"
}

func (r *EcrServicesK8SAwsRepositoryV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Repository is the Schema for the Repositories API",
		MarkdownDescription: "Repository is the Schema for the Repositories API",
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
				Description:         "RepositorySpec defines the desired state of Repository.  An object representing a repository.",
				MarkdownDescription: "RepositorySpec defines the desired state of Repository.  An object representing a repository.",
				Attributes: map[string]schema.Attribute{
					"encryption_configuration": schema.SingleNestedAttribute{
						Description:         "The encryption configuration for the repository. This determines how the contents of your repository are encrypted at rest.",
						MarkdownDescription: "The encryption configuration for the repository. This determines how the contents of your repository are encrypted at rest.",
						Attributes: map[string]schema.Attribute{
							"encryption_type": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"kms_key": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"image_scanning_configuration": schema.SingleNestedAttribute{
						Description:         "The image scanning configuration for the repository. This determines whether images are scanned for known vulnerabilities after being pushed to the repository.",
						MarkdownDescription: "The image scanning configuration for the repository. This determines whether images are scanned for known vulnerabilities after being pushed to the repository.",
						Attributes: map[string]schema.Attribute{
							"scan_on_push": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"image_tag_mutability": schema.StringAttribute{
						Description:         "The tag mutability setting for the repository. If this parameter is omitted, the default setting of MUTABLE will be used which will allow image tags to be overwritten. If IMMUTABLE is specified, all image tags within the repository will be immutable which will prevent them from being overwritten.",
						MarkdownDescription: "The tag mutability setting for the repository. If this parameter is omitted, the default setting of MUTABLE will be used which will allow image tags to be overwritten. If IMMUTABLE is specified, all image tags within the repository will be immutable which will prevent them from being overwritten.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"lifecycle_policy": schema.StringAttribute{
						Description:         "The JSON repository policy text to apply to the repository.",
						MarkdownDescription: "The JSON repository policy text to apply to the repository.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"name": schema.StringAttribute{
						Description:         "The name to use for the repository. The repository name may be specified on its own (such as nginx-web-app) or it can be prepended with a namespace to group the repository into a category (such as project-a/nginx-web-app).",
						MarkdownDescription: "The name to use for the repository. The repository name may be specified on its own (such as nginx-web-app) or it can be prepended with a namespace to group the repository into a category (such as project-a/nginx-web-app).",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"policy": schema.StringAttribute{
						Description:         "The JSON repository policy text to apply to the repository. For more information, see Amazon ECR repository policies (https://docs.aws.amazon.com/AmazonECR/latest/userguide/repository-policy-examples.html) in the Amazon Elastic Container Registry User Guide.",
						MarkdownDescription: "The JSON repository policy text to apply to the repository. For more information, see Amazon ECR repository policies (https://docs.aws.amazon.com/AmazonECR/latest/userguide/repository-policy-examples.html) in the Amazon Elastic Container Registry User Guide.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"registry_id": schema.StringAttribute{
						Description:         "The Amazon Web Services account ID associated with the registry to create the repository. If you do not specify a registry, the default registry is assumed.",
						MarkdownDescription: "The Amazon Web Services account ID associated with the registry to create the repository. If you do not specify a registry, the default registry is assumed.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"tags": schema.ListNestedAttribute{
						Description:         "The metadata that you apply to the repository to help you categorize and organize them. Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.",
						MarkdownDescription: "The metadata that you apply to the repository to help you categorize and organize them. Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"key": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"value": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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

func (r *EcrServicesK8SAwsRepositoryV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *EcrServicesK8SAwsRepositoryV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_ecr_services_k8s_aws_repository_v1alpha1")

	var data EcrServicesK8SAwsRepositoryV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "ecr.services.k8s.aws", Version: "v1alpha1", Resource: "repositories"}).
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

	var readResponse EcrServicesK8SAwsRepositoryV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("ecr.services.k8s.aws/v1alpha1")
	data.Kind = pointer.String("Repository")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
