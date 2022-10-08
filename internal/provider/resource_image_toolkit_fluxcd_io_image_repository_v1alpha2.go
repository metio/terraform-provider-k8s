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

type ImageToolkitFluxcdIoImageRepositoryV1Alpha2Resource struct{}

var (
	_ resource.Resource = (*ImageToolkitFluxcdIoImageRepositoryV1Alpha2Resource)(nil)
)

type ImageToolkitFluxcdIoImageRepositoryV1Alpha2TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type ImageToolkitFluxcdIoImageRepositoryV1Alpha2GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		Suspend *bool `tfsdk:"suspend" yaml:"suspend,omitempty"`

		Timeout *string `tfsdk:"timeout" yaml:"timeout,omitempty"`

		CertSecretRef *struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"cert_secret_ref" yaml:"certSecretRef,omitempty"`

		Image *string `tfsdk:"image" yaml:"image,omitempty"`

		Interval *string `tfsdk:"interval" yaml:"interval,omitempty"`

		SecretRef *struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewImageToolkitFluxcdIoImageRepositoryV1Alpha2Resource() resource.Resource {
	return &ImageToolkitFluxcdIoImageRepositoryV1Alpha2Resource{}
}

func (r *ImageToolkitFluxcdIoImageRepositoryV1Alpha2Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_image_toolkit_fluxcd_io_image_repository_v1alpha2"
}

func (r *ImageToolkitFluxcdIoImageRepositoryV1Alpha2Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "ImageRepository is the Schema for the imagerepositories API",
		MarkdownDescription: "ImageRepository is the Schema for the imagerepositories API",
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
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.RequiresReplace(),
						},
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
				Description:         "ImageRepositorySpec defines the parameters for scanning an image repository, e.g., 'fluxcd/flux'.",
				MarkdownDescription: "ImageRepositorySpec defines the parameters for scanning an image repository, e.g., 'fluxcd/flux'.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"suspend": {
						Description:         "This flag tells the controller to suspend subsequent image scans. It does not apply to already started scans. Defaults to false.",
						MarkdownDescription: "This flag tells the controller to suspend subsequent image scans. It does not apply to already started scans. Defaults to false.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"timeout": {
						Description:         "Timeout for image scanning. Defaults to 'Interval' duration.",
						MarkdownDescription: "Timeout for image scanning. Defaults to 'Interval' duration.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"cert_secret_ref": {
						Description:         "CertSecretRef can be given the name of a secret containing either or both of  - a PEM-encoded client certificate ('certFile') and private key ('keyFile'); - a PEM-encoded CA certificate ('caFile')  and whichever are supplied, will be used for connecting to the registry. The client cert and key are useful if you are authenticating with a certificate; the CA cert is useful if you are using a self-signed server certificate.",
						MarkdownDescription: "CertSecretRef can be given the name of a secret containing either or both of  - a PEM-encoded client certificate ('certFile') and private key ('keyFile'); - a PEM-encoded CA certificate ('caFile')  and whichever are supplied, will be used for connecting to the registry. The client cert and key are useful if you are authenticating with a certificate; the CA cert is useful if you are using a self-signed server certificate.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "Name of the referent.",
								MarkdownDescription: "Name of the referent.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"image": {
						Description:         "Image is the name of the image repository",
						MarkdownDescription: "Image is the name of the image repository",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"interval": {
						Description:         "Interval is the length of time to wait between scans of the image repository.",
						MarkdownDescription: "Interval is the length of time to wait between scans of the image repository.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"secret_ref": {
						Description:         "SecretRef can be given the name of a secret containing credentials to use for the image registry. The secret should be created with 'kubectl create secret docker-registry', or the equivalent.",
						MarkdownDescription: "SecretRef can be given the name of a secret containing credentials to use for the image registry. The secret should be created with 'kubectl create secret docker-registry', or the equivalent.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "Name of the referent.",
								MarkdownDescription: "Name of the referent.",

								Type: types.StringType,

								Required: true,
								Optional: false,
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

func (r *ImageToolkitFluxcdIoImageRepositoryV1Alpha2Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_image_toolkit_fluxcd_io_image_repository_v1alpha2")

	var state ImageToolkitFluxcdIoImageRepositoryV1Alpha2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ImageToolkitFluxcdIoImageRepositoryV1Alpha2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("image.toolkit.fluxcd.io/v1alpha2")
	goModel.Kind = utilities.Ptr("ImageRepository")

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

func (r *ImageToolkitFluxcdIoImageRepositoryV1Alpha2Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_image_toolkit_fluxcd_io_image_repository_v1alpha2")
	// NO-OP: All data is already in Terraform state
}

func (r *ImageToolkitFluxcdIoImageRepositoryV1Alpha2Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_image_toolkit_fluxcd_io_image_repository_v1alpha2")

	var state ImageToolkitFluxcdIoImageRepositoryV1Alpha2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ImageToolkitFluxcdIoImageRepositoryV1Alpha2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("image.toolkit.fluxcd.io/v1alpha2")
	goModel.Kind = utilities.Ptr("ImageRepository")

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

func (r *ImageToolkitFluxcdIoImageRepositoryV1Alpha2Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_image_toolkit_fluxcd_io_image_repository_v1alpha2")
	// NO-OP: Terraform removes the state automatically for us
}
