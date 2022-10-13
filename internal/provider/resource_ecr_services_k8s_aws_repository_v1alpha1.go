/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"gopkg.in/yaml.v3"
	"time"
)

type EcrServicesK8SAwsRepositoryV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*EcrServicesK8SAwsRepositoryV1Alpha1Resource)(nil)
)

type EcrServicesK8SAwsRepositoryV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type EcrServicesK8SAwsRepositoryV1Alpha1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		EncryptionConfiguration *struct {
			EncryptionType *string `tfsdk:"encryption_type" yaml:"encryptionType,omitempty"`

			KmsKey *string `tfsdk:"kms_key" yaml:"kmsKey,omitempty"`
		} `tfsdk:"encryption_configuration" yaml:"encryptionConfiguration,omitempty"`

		ImageScanningConfiguration *struct {
			ScanOnPush *bool `tfsdk:"scan_on_push" yaml:"scanOnPush,omitempty"`
		} `tfsdk:"image_scanning_configuration" yaml:"imageScanningConfiguration,omitempty"`

		ImageTagMutability *string `tfsdk:"image_tag_mutability" yaml:"imageTagMutability,omitempty"`

		LifecyclePolicy *string `tfsdk:"lifecycle_policy" yaml:"lifecyclePolicy,omitempty"`

		Name *string `tfsdk:"name" yaml:"name,omitempty"`

		Policy *string `tfsdk:"policy" yaml:"policy,omitempty"`

		RegistryID *string `tfsdk:"registry_id" yaml:"registryID,omitempty"`

		Tags *[]struct {
			Key *string `tfsdk:"key" yaml:"key,omitempty"`

			Value *string `tfsdk:"value" yaml:"value,omitempty"`
		} `tfsdk:"tags" yaml:"tags,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewEcrServicesK8SAwsRepositoryV1Alpha1Resource() resource.Resource {
	return &EcrServicesK8SAwsRepositoryV1Alpha1Resource{}
}

func (r *EcrServicesK8SAwsRepositoryV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ecr_services_k8s_aws_repository_v1alpha1"
}

func (r *EcrServicesK8SAwsRepositoryV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Repository is the Schema for the Repositories API",
		MarkdownDescription: "Repository is the Schema for the Repositories API",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description:         "The timestamp of the last change to this resource.",
				MarkdownDescription: "The timestamp of the last change to this resource.",
				Type:                types.Int64Type,
				Computed:            true,
				Optional:            false,
			},

			"yaml": {
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"metadata": {
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Type:                types.StringType,
						Required:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
					},

					"namespace": {
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Type:                types.StringType,
						Optional:            true,
					},

					"labels": {
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.LabelValidator(),
						},
					},
					"annotations": {
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.AnnotationValidator(),
						},
					},
				}),
			},

			"api_version": {
				Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"kind": {
				Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"spec": {
				Description:         "RepositorySpec defines the desired state of Repository.  An object representing a repository.",
				MarkdownDescription: "RepositorySpec defines the desired state of Repository.  An object representing a repository.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"encryption_configuration": {
						Description:         "The encryption configuration for the repository. This determines how the contents of your repository are encrypted at rest.",
						MarkdownDescription: "The encryption configuration for the repository. This determines how the contents of your repository are encrypted at rest.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"encryption_type": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"kms_key": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"image_scanning_configuration": {
						Description:         "The image scanning configuration for the repository. This determines whether images are scanned for known vulnerabilities after being pushed to the repository.",
						MarkdownDescription: "The image scanning configuration for the repository. This determines whether images are scanned for known vulnerabilities after being pushed to the repository.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"scan_on_push": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"image_tag_mutability": {
						Description:         "The tag mutability setting for the repository. If this parameter is omitted, the default setting of MUTABLE will be used which will allow image tags to be overwritten. If IMMUTABLE is specified, all image tags within the repository will be immutable which will prevent them from being overwritten.",
						MarkdownDescription: "The tag mutability setting for the repository. If this parameter is omitted, the default setting of MUTABLE will be used which will allow image tags to be overwritten. If IMMUTABLE is specified, all image tags within the repository will be immutable which will prevent them from being overwritten.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"lifecycle_policy": {
						Description:         "The JSON repository policy text to apply to the repository.",
						MarkdownDescription: "The JSON repository policy text to apply to the repository.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"name": {
						Description:         "The name to use for the repository. The repository name may be specified on its own (such as nginx-web-app) or it can be prepended with a namespace to group the repository into a category (such as project-a/nginx-web-app).",
						MarkdownDescription: "The name to use for the repository. The repository name may be specified on its own (such as nginx-web-app) or it can be prepended with a namespace to group the repository into a category (such as project-a/nginx-web-app).",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"policy": {
						Description:         "The JSON repository policy text to apply to the repository. For more information, see Amazon ECR repository policies (https://docs.aws.amazon.com/AmazonECR/latest/userguide/repository-policy-examples.html) in the Amazon Elastic Container Registry User Guide.",
						MarkdownDescription: "The JSON repository policy text to apply to the repository. For more information, see Amazon ECR repository policies (https://docs.aws.amazon.com/AmazonECR/latest/userguide/repository-policy-examples.html) in the Amazon Elastic Container Registry User Guide.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"registry_id": {
						Description:         "The Amazon Web Services account ID associated with the registry to create the repository. If you do not specify a registry, the default registry is assumed.",
						MarkdownDescription: "The Amazon Web Services account ID associated with the registry to create the repository. If you do not specify a registry, the default registry is assumed.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"tags": {
						Description:         "The metadata that you apply to the repository to help you categorize and organize them. Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.",
						MarkdownDescription: "The metadata that you apply to the repository to help you categorize and organize them. Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"key": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"value": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *EcrServicesK8SAwsRepositoryV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_ecr_services_k8s_aws_repository_v1alpha1")

	var state EcrServicesK8SAwsRepositoryV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel EcrServicesK8SAwsRepositoryV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("ecr.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("Repository")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *EcrServicesK8SAwsRepositoryV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_ecr_services_k8s_aws_repository_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *EcrServicesK8SAwsRepositoryV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_ecr_services_k8s_aws_repository_v1alpha1")

	var state EcrServicesK8SAwsRepositoryV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel EcrServicesK8SAwsRepositoryV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("ecr.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("Repository")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *EcrServicesK8SAwsRepositoryV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_ecr_services_k8s_aws_repository_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
