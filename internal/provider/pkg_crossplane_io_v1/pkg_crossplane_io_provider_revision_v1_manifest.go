/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package pkg_crossplane_io_v1

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
	_ datasource.DataSource = &PkgCrossplaneIoProviderRevisionV1Manifest{}
)

func NewPkgCrossplaneIoProviderRevisionV1Manifest() datasource.DataSource {
	return &PkgCrossplaneIoProviderRevisionV1Manifest{}
}

type PkgCrossplaneIoProviderRevisionV1Manifest struct{}

type PkgCrossplaneIoProviderRevisionV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		CommonLabels        *map[string]string `tfsdk:"common_labels" json:"commonLabels,omitempty"`
		ControllerConfigRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"controller_config_ref" json:"controllerConfigRef,omitempty"`
		DesiredState                *string `tfsdk:"desired_state" json:"desiredState,omitempty"`
		IgnoreCrossplaneConstraints *bool   `tfsdk:"ignore_crossplane_constraints" json:"ignoreCrossplaneConstraints,omitempty"`
		Image                       *string `tfsdk:"image" json:"image,omitempty"`
		PackagePullPolicy           *string `tfsdk:"package_pull_policy" json:"packagePullPolicy,omitempty"`
		PackagePullSecrets          *[]struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"package_pull_secrets" json:"packagePullSecrets,omitempty"`
		Revision         *int64 `tfsdk:"revision" json:"revision,omitempty"`
		RuntimeConfigRef *struct {
			ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
			Name       *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"runtime_config_ref" json:"runtimeConfigRef,omitempty"`
		SkipDependencyResolution *bool   `tfsdk:"skip_dependency_resolution" json:"skipDependencyResolution,omitempty"`
		TlsClientSecretName      *string `tfsdk:"tls_client_secret_name" json:"tlsClientSecretName,omitempty"`
		TlsServerSecretName      *string `tfsdk:"tls_server_secret_name" json:"tlsServerSecretName,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *PkgCrossplaneIoProviderRevisionV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_pkg_crossplane_io_provider_revision_v1_manifest"
}

func (r *PkgCrossplaneIoProviderRevisionV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "A ProviderRevision that has been added to Crossplane.",
		MarkdownDescription: "A ProviderRevision that has been added to Crossplane.",
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
				Description:         "ProviderRevisionSpec specifies configuration for a ProviderRevision.",
				MarkdownDescription: "ProviderRevisionSpec specifies configuration for a ProviderRevision.",
				Attributes: map[string]schema.Attribute{
					"common_labels": schema.MapAttribute{
						Description:         "Map of string keys and values that can be used to organize and categorize(scope and select) objects. May match selectors of replication controllersand services.More info: http://kubernetes.io/docs/user-guide/labels",
						MarkdownDescription: "Map of string keys and values that can be used to organize and categorize(scope and select) objects. May match selectors of replication controllersand services.More info: http://kubernetes.io/docs/user-guide/labels",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"controller_config_ref": schema.SingleNestedAttribute{
						Description:         "ControllerConfigRef references a ControllerConfig resource that will beused to configure the packaged controller Deployment.Deprecated: Use RuntimeConfigReference instead.",
						MarkdownDescription: "ControllerConfigRef references a ControllerConfig resource that will beused to configure the packaged controller Deployment.Deprecated: Use RuntimeConfigReference instead.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of the ControllerConfig.",
								MarkdownDescription: "Name of the ControllerConfig.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"desired_state": schema.StringAttribute{
						Description:         "DesiredState of the PackageRevision. Can be either Active or Inactive.",
						MarkdownDescription: "DesiredState of the PackageRevision. Can be either Active or Inactive.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"ignore_crossplane_constraints": schema.BoolAttribute{
						Description:         "IgnoreCrossplaneConstraints indicates to the package manager whether tohonor Crossplane version constrains specified by the package.Default is false.",
						MarkdownDescription: "IgnoreCrossplaneConstraints indicates to the package manager whether tohonor Crossplane version constrains specified by the package.Default is false.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image": schema.StringAttribute{
						Description:         "Package image used by install Pod to extract package contents.",
						MarkdownDescription: "Package image used by install Pod to extract package contents.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"package_pull_policy": schema.StringAttribute{
						Description:         "PackagePullPolicy defines the pull policy for the package. It is alsoapplied to any images pulled for the package, such as a provider'scontroller image.Default is IfNotPresent.",
						MarkdownDescription: "PackagePullPolicy defines the pull policy for the package. It is alsoapplied to any images pulled for the package, such as a provider'scontroller image.Default is IfNotPresent.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"package_pull_secrets": schema.ListNestedAttribute{
						Description:         "PackagePullSecrets are named secrets in the same namespace that can beused to fetch packages from private registries. They are also applied toany images pulled for the package, such as a provider's controller image.",
						MarkdownDescription: "PackagePullSecrets are named secrets in the same namespace that can beused to fetch packages from private registries. They are also applied toany images pulled for the package, such as a provider's controller image.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
									MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"revision": schema.Int64Attribute{
						Description:         "Revision number. Indicates when the revision will be garbage collectedbased on the parent's RevisionHistoryLimit.",
						MarkdownDescription: "Revision number. Indicates when the revision will be garbage collectedbased on the parent's RevisionHistoryLimit.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"runtime_config_ref": schema.SingleNestedAttribute{
						Description:         "RuntimeConfigRef references a RuntimeConfig resource that will be usedto configure the package runtime.",
						MarkdownDescription: "RuntimeConfigRef references a RuntimeConfig resource that will be usedto configure the package runtime.",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "API version of the referent.",
								MarkdownDescription: "API version of the referent.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind of the referent.",
								MarkdownDescription: "Kind of the referent.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the RuntimeConfig.",
								MarkdownDescription: "Name of the RuntimeConfig.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"skip_dependency_resolution": schema.BoolAttribute{
						Description:         "SkipDependencyResolution indicates to the package manager whether to skipresolving dependencies for a package. Setting this value to true may haveunintended consequences.Default is false.",
						MarkdownDescription: "SkipDependencyResolution indicates to the package manager whether to skipresolving dependencies for a package. Setting this value to true may haveunintended consequences.Default is false.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tls_client_secret_name": schema.StringAttribute{
						Description:         "TLSClientSecretName is the name of the TLS Secret that stores clientcertificates of the Provider.",
						MarkdownDescription: "TLSClientSecretName is the name of the TLS Secret that stores clientcertificates of the Provider.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tls_server_secret_name": schema.StringAttribute{
						Description:         "TLSServerSecretName is the name of the TLS Secret that stores servercertificates of the Provider.",
						MarkdownDescription: "TLSServerSecretName is the name of the TLS Secret that stores servercertificates of the Provider.",
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

func (r *PkgCrossplaneIoProviderRevisionV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_pkg_crossplane_io_provider_revision_v1_manifest")

	var model PkgCrossplaneIoProviderRevisionV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("pkg.crossplane.io/v1")
	model.Kind = pointer.String("ProviderRevision")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
