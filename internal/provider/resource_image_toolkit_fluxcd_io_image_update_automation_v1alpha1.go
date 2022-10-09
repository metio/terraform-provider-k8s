/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

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

type ImageToolkitFluxcdIoImageUpdateAutomationV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*ImageToolkitFluxcdIoImageUpdateAutomationV1Alpha1Resource)(nil)
)

type ImageToolkitFluxcdIoImageUpdateAutomationV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type ImageToolkitFluxcdIoImageUpdateAutomationV1Alpha1GoModel struct {
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
		Checkout *struct {
			Branch *string `tfsdk:"branch" yaml:"branch,omitempty"`

			GitRepositoryRef *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"git_repository_ref" yaml:"gitRepositoryRef,omitempty"`
		} `tfsdk:"checkout" yaml:"checkout,omitempty"`

		Commit *struct {
			AuthorEmail *string `tfsdk:"author_email" yaml:"authorEmail,omitempty"`

			AuthorName *string `tfsdk:"author_name" yaml:"authorName,omitempty"`

			MessageTemplate *string `tfsdk:"message_template" yaml:"messageTemplate,omitempty"`

			SigningKey *struct {
				SecretRef *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`
			} `tfsdk:"signing_key" yaml:"signingKey,omitempty"`
		} `tfsdk:"commit" yaml:"commit,omitempty"`

		Interval *string `tfsdk:"interval" yaml:"interval,omitempty"`

		Push *struct {
			Branch *string `tfsdk:"branch" yaml:"branch,omitempty"`
		} `tfsdk:"push" yaml:"push,omitempty"`

		Suspend *bool `tfsdk:"suspend" yaml:"suspend,omitempty"`

		Update *struct {
			Path *string `tfsdk:"path" yaml:"path,omitempty"`

			Strategy *string `tfsdk:"strategy" yaml:"strategy,omitempty"`
		} `tfsdk:"update" yaml:"update,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewImageToolkitFluxcdIoImageUpdateAutomationV1Alpha1Resource() resource.Resource {
	return &ImageToolkitFluxcdIoImageUpdateAutomationV1Alpha1Resource{}
}

func (r *ImageToolkitFluxcdIoImageUpdateAutomationV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_image_toolkit_fluxcd_io_image_update_automation_v1alpha1"
}

func (r *ImageToolkitFluxcdIoImageUpdateAutomationV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "ImageUpdateAutomation is the Schema for the imageupdateautomations API",
		MarkdownDescription: "ImageUpdateAutomation is the Schema for the imageupdateautomations API",
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
				Description:         "ImageUpdateAutomationSpec defines the desired state of ImageUpdateAutomation",
				MarkdownDescription: "ImageUpdateAutomationSpec defines the desired state of ImageUpdateAutomation",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"checkout": {
						Description:         "Checkout gives the parameters for cloning the git repository, ready to make changes.",
						MarkdownDescription: "Checkout gives the parameters for cloning the git repository, ready to make changes.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"branch": {
								Description:         "Branch gives the branch to clone from the git repository. If '.spec.push' is not supplied, commits will also be pushed to this branch.",
								MarkdownDescription: "Branch gives the branch to clone from the git repository. If '.spec.push' is not supplied, commits will also be pushed to this branch.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"git_repository_ref": {
								Description:         "GitRepositoryRef refers to the resource giving access details to a git repository to update files in.",
								MarkdownDescription: "GitRepositoryRef refers to the resource giving access details to a git repository to update files in.",

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

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"commit": {
						Description:         "Commit specifies how to commit to the git repository.",
						MarkdownDescription: "Commit specifies how to commit to the git repository.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"author_email": {
								Description:         "AuthorEmail gives the email to provide when making a commit",
								MarkdownDescription: "AuthorEmail gives the email to provide when making a commit",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"author_name": {
								Description:         "AuthorName gives the name to provide when making a commit",
								MarkdownDescription: "AuthorName gives the name to provide when making a commit",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"message_template": {
								Description:         "MessageTemplate provides a template for the commit message, into which will be interpolated the details of the change made.",
								MarkdownDescription: "MessageTemplate provides a template for the commit message, into which will be interpolated the details of the change made.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"signing_key": {
								Description:         "SigningKey provides the option to sign commits with a GPG key",
								MarkdownDescription: "SigningKey provides the option to sign commits with a GPG key",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"secret_ref": {
										Description:         "SecretRef holds the name to a secret that contains a 'git.asc' key corresponding to the ASCII Armored file containing the GPG signing keypair as the value. It must be in the same namespace as the ImageUpdateAutomation.",
										MarkdownDescription: "SecretRef holds the name to a secret that contains a 'git.asc' key corresponding to the ASCII Armored file containing the GPG signing keypair as the value. It must be in the same namespace as the ImageUpdateAutomation.",

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
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"interval": {
						Description:         "Interval gives an lower bound for how often the automation run should be attempted.",
						MarkdownDescription: "Interval gives an lower bound for how often the automation run should be attempted.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"push": {
						Description:         "Push specifies how and where to push commits made by the automation. If missing, commits are pushed (back) to '.spec.checkout.branch'.",
						MarkdownDescription: "Push specifies how and where to push commits made by the automation. If missing, commits are pushed (back) to '.spec.checkout.branch'.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"branch": {
								Description:         "Branch specifies that commits should be pushed to the branch named. The branch is created using '.spec.checkout.branch' as the starting point, if it doesn't already exist.",
								MarkdownDescription: "Branch specifies that commits should be pushed to the branch named. The branch is created using '.spec.checkout.branch' as the starting point, if it doesn't already exist.",

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

					"suspend": {
						Description:         "Suspend tells the controller to not run this automation, until it is unset (or set to false). Defaults to false.",
						MarkdownDescription: "Suspend tells the controller to not run this automation, until it is unset (or set to false). Defaults to false.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"update": {
						Description:         "Update gives the specification for how to update the files in the repository. This can be left empty, to use the default value.",
						MarkdownDescription: "Update gives the specification for how to update the files in the repository. This can be left empty, to use the default value.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"path": {
								Description:         "Path to the directory containing the manifests to be updated. Defaults to 'None', which translates to the root path of the GitRepositoryRef.",
								MarkdownDescription: "Path to the directory containing the manifests to be updated. Defaults to 'None', which translates to the root path of the GitRepositoryRef.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"strategy": {
								Description:         "Strategy names the strategy to be used.",
								MarkdownDescription: "Strategy names the strategy to be used.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("Setters"),
								},
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

func (r *ImageToolkitFluxcdIoImageUpdateAutomationV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_image_toolkit_fluxcd_io_image_update_automation_v1alpha1")

	var state ImageToolkitFluxcdIoImageUpdateAutomationV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ImageToolkitFluxcdIoImageUpdateAutomationV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("image.toolkit.fluxcd.io/v1alpha1")
	goModel.Kind = utilities.Ptr("ImageUpdateAutomation")

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

func (r *ImageToolkitFluxcdIoImageUpdateAutomationV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_image_toolkit_fluxcd_io_image_update_automation_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *ImageToolkitFluxcdIoImageUpdateAutomationV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_image_toolkit_fluxcd_io_image_update_automation_v1alpha1")

	var state ImageToolkitFluxcdIoImageUpdateAutomationV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ImageToolkitFluxcdIoImageUpdateAutomationV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("image.toolkit.fluxcd.io/v1alpha1")
	goModel.Kind = utilities.Ptr("ImageUpdateAutomation")

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

func (r *ImageToolkitFluxcdIoImageUpdateAutomationV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_image_toolkit_fluxcd_io_image_update_automation_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
