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

type PkgCrossplaneIoLockV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*PkgCrossplaneIoLockV1Alpha1Resource)(nil)
)

type PkgCrossplaneIoLockV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Packages   types.List   `tfsdk:"packages"`
}

type PkgCrossplaneIoLockV1Alpha1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Packages *[]struct {
		Dependencies *[]struct {
			Constraints *string `tfsdk:"constraints" yaml:"constraints,omitempty"`

			Package *string `tfsdk:"package" yaml:"package,omitempty"`

			Type *string `tfsdk:"type" yaml:"type,omitempty"`
		} `tfsdk:"dependencies" yaml:"dependencies,omitempty"`

		Name *string `tfsdk:"name" yaml:"name,omitempty"`

		Source *string `tfsdk:"source" yaml:"source,omitempty"`

		Type *string `tfsdk:"type" yaml:"type,omitempty"`

		Version *string `tfsdk:"version" yaml:"version,omitempty"`
	} `tfsdk:"packages" yaml:"packages,omitempty"`
}

func NewPkgCrossplaneIoLockV1Alpha1Resource() resource.Resource {
	return &PkgCrossplaneIoLockV1Alpha1Resource{}
}

func (r *PkgCrossplaneIoLockV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_pkg_crossplane_io_lock_v1alpha1"
}

func (r *PkgCrossplaneIoLockV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Lock is the CRD type that tracks package dependencies. [DEPRECATED]: Please use the identical v1beta1 API instead. The v1alpha1 API is scheduled to be removed in Crossplane v1.7.",
		MarkdownDescription: "Lock is the CRD type that tracks package dependencies. [DEPRECATED]: Please use the identical v1beta1 API instead. The v1alpha1 API is scheduled to be removed in Crossplane v1.7.",
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

			"packages": {
				Description:         "",
				MarkdownDescription: "",

				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

					"dependencies": {
						Description:         "Dependencies are the list of dependencies of this package. The order of the dependencies will dictate the order in which they are resolved.",
						MarkdownDescription: "Dependencies are the list of dependencies of this package. The order of the dependencies will dictate the order in which they are resolved.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"constraints": {
								Description:         "Constraints is a valid semver range, which will be used to select a valid dependency version.",
								MarkdownDescription: "Constraints is a valid semver range, which will be used to select a valid dependency version.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"package": {
								Description:         "Package is the OCI image name without a tag or digest.",
								MarkdownDescription: "Package is the OCI image name without a tag or digest.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"type": {
								Description:         "Type is the type of package. Can be either Configuration or Provider.",
								MarkdownDescription: "Type is the type of package. Can be either Configuration or Provider.",

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

					"name": {
						Description:         "Name corresponds to the name of the package revision for this package.",
						MarkdownDescription: "Name corresponds to the name of the package revision for this package.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"source": {
						Description:         "Source is the OCI image name without a tag or digest.",
						MarkdownDescription: "Source is the OCI image name without a tag or digest.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"type": {
						Description:         "Type is the type of package. Can be either Configuration or Provider.",
						MarkdownDescription: "Type is the type of package. Can be either Configuration or Provider.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"version": {
						Description:         "Version is the tag or digest of the OCI image.",
						MarkdownDescription: "Version is the tag or digest of the OCI image.",

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
		},
	}, nil
}

func (r *PkgCrossplaneIoLockV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_pkg_crossplane_io_lock_v1alpha1")

	var state PkgCrossplaneIoLockV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel PkgCrossplaneIoLockV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("pkg.crossplane.io/v1alpha1")
	goModel.Kind = utilities.Ptr("Lock")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *PkgCrossplaneIoLockV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_pkg_crossplane_io_lock_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *PkgCrossplaneIoLockV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_pkg_crossplane_io_lock_v1alpha1")

	var state PkgCrossplaneIoLockV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel PkgCrossplaneIoLockV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("pkg.crossplane.io/v1alpha1")
	goModel.Kind = utilities.Ptr("Lock")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *PkgCrossplaneIoLockV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_pkg_crossplane_io_lock_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
