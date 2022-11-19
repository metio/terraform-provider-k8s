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

type PkgCrossplaneIoConfigurationRevisionV1Resource struct{}

var (
	_ resource.Resource = (*PkgCrossplaneIoConfigurationRevisionV1Resource)(nil)
)

type PkgCrossplaneIoConfigurationRevisionV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type PkgCrossplaneIoConfigurationRevisionV1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		ControllerConfigRef *struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"controller_config_ref" yaml:"controllerConfigRef,omitempty"`

		DesiredState *string `tfsdk:"desired_state" yaml:"desiredState,omitempty"`

		IgnoreCrossplaneConstraints *bool `tfsdk:"ignore_crossplane_constraints" yaml:"ignoreCrossplaneConstraints,omitempty"`

		Image *string `tfsdk:"image" yaml:"image,omitempty"`

		PackagePullPolicy *string `tfsdk:"package_pull_policy" yaml:"packagePullPolicy,omitempty"`

		PackagePullSecrets *[]struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"package_pull_secrets" yaml:"packagePullSecrets,omitempty"`

		Revision *int64 `tfsdk:"revision" yaml:"revision,omitempty"`

		SkipDependencyResolution *bool `tfsdk:"skip_dependency_resolution" yaml:"skipDependencyResolution,omitempty"`

		WebhookTLSSecretName *string `tfsdk:"webhook_tls_secret_name" yaml:"webhookTLSSecretName,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewPkgCrossplaneIoConfigurationRevisionV1Resource() resource.Resource {
	return &PkgCrossplaneIoConfigurationRevisionV1Resource{}
}

func (r *PkgCrossplaneIoConfigurationRevisionV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_pkg_crossplane_io_configuration_revision_v1"
}

func (r *PkgCrossplaneIoConfigurationRevisionV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "A ConfigurationRevision that has been added to Crossplane.",
		MarkdownDescription: "A ConfigurationRevision that has been added to Crossplane.",
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

			"spec": {
				Description:         "PackageRevisionSpec specifies the desired state of a PackageRevision.",
				MarkdownDescription: "PackageRevisionSpec specifies the desired state of a PackageRevision.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"controller_config_ref": {
						Description:         "ControllerConfigRef references a ControllerConfig resource that will be used to configure the packaged controller Deployment.",
						MarkdownDescription: "ControllerConfigRef references a ControllerConfig resource that will be used to configure the packaged controller Deployment.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "Name of the ControllerConfig.",
								MarkdownDescription: "Name of the ControllerConfig.",

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

					"desired_state": {
						Description:         "DesiredState of the PackageRevision. Can be either Active or Inactive.",
						MarkdownDescription: "DesiredState of the PackageRevision. Can be either Active or Inactive.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"ignore_crossplane_constraints": {
						Description:         "IgnoreCrossplaneConstraints indicates to the package manager whether to honor Crossplane version constrains specified by the package. Default is false.",
						MarkdownDescription: "IgnoreCrossplaneConstraints indicates to the package manager whether to honor Crossplane version constrains specified by the package. Default is false.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"image": {
						Description:         "Package image used by install Pod to extract package contents.",
						MarkdownDescription: "Package image used by install Pod to extract package contents.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"package_pull_policy": {
						Description:         "PackagePullPolicy defines the pull policy for the package. It is also applied to any images pulled for the package, such as a provider's controller image. Default is IfNotPresent.",
						MarkdownDescription: "PackagePullPolicy defines the pull policy for the package. It is also applied to any images pulled for the package, such as a provider's controller image. Default is IfNotPresent.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"package_pull_secrets": {
						Description:         "PackagePullSecrets are named secrets in the same namespace that can be used to fetch packages from private registries. They are also applied to any images pulled for the package, such as a provider's controller image.",
						MarkdownDescription: "PackagePullSecrets are named secrets in the same namespace that can be used to fetch packages from private registries. They are also applied to any images pulled for the package, such as a provider's controller image.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
								MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

					"revision": {
						Description:         "Revision number. Indicates when the revision will be garbage collected based on the parent's RevisionHistoryLimit.",
						MarkdownDescription: "Revision number. Indicates when the revision will be garbage collected based on the parent's RevisionHistoryLimit.",

						Type: types.Int64Type,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"skip_dependency_resolution": {
						Description:         "SkipDependencyResolution indicates to the package manager whether to skip resolving dependencies for a package. Setting this value to true may have unintended consequences. Default is false.",
						MarkdownDescription: "SkipDependencyResolution indicates to the package manager whether to skip resolving dependencies for a package. Setting this value to true may have unintended consequences. Default is false.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"webhook_tls_secret_name": {
						Description:         "WebhookTLSSecretName is the name of the TLS Secret that will be used by the provider to serve a TLS-enabled webhook server. The certificate will be injected to webhook configurations as well as CRD conversion webhook strategy if needed. If it's not given, provider will not have a certificate mounted to its filesystem, webhook configurations won't be deployed and if there is a CRD with webhook conversion strategy, the installation will fail.",
						MarkdownDescription: "WebhookTLSSecretName is the name of the TLS Secret that will be used by the provider to serve a TLS-enabled webhook server. The certificate will be injected to webhook configurations as well as CRD conversion webhook strategy if needed. If it's not given, provider will not have a certificate mounted to its filesystem, webhook configurations won't be deployed and if there is a CRD with webhook conversion strategy, the installation will fail.",

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
		},
	}, nil
}

func (r *PkgCrossplaneIoConfigurationRevisionV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_pkg_crossplane_io_configuration_revision_v1")

	var state PkgCrossplaneIoConfigurationRevisionV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel PkgCrossplaneIoConfigurationRevisionV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("pkg.crossplane.io/v1")
	goModel.Kind = utilities.Ptr("ConfigurationRevision")

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

func (r *PkgCrossplaneIoConfigurationRevisionV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_pkg_crossplane_io_configuration_revision_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *PkgCrossplaneIoConfigurationRevisionV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_pkg_crossplane_io_configuration_revision_v1")

	var state PkgCrossplaneIoConfigurationRevisionV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel PkgCrossplaneIoConfigurationRevisionV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("pkg.crossplane.io/v1")
	goModel.Kind = utilities.Ptr("ConfigurationRevision")

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

func (r *PkgCrossplaneIoConfigurationRevisionV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_pkg_crossplane_io_configuration_revision_v1")
	// NO-OP: Terraform removes the state automatically for us
}
