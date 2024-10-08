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
	_ datasource.DataSource = &PkgCrossplaneIoConfigurationRevisionV1Manifest{}
)

func NewPkgCrossplaneIoConfigurationRevisionV1Manifest() datasource.DataSource {
	return &PkgCrossplaneIoConfigurationRevisionV1Manifest{}
}

type PkgCrossplaneIoConfigurationRevisionV1Manifest struct{}

type PkgCrossplaneIoConfigurationRevisionV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		CommonLabels                *map[string]string `tfsdk:"common_labels" json:"commonLabels,omitempty"`
		DesiredState                *string            `tfsdk:"desired_state" json:"desiredState,omitempty"`
		IgnoreCrossplaneConstraints *bool              `tfsdk:"ignore_crossplane_constraints" json:"ignoreCrossplaneConstraints,omitempty"`
		Image                       *string            `tfsdk:"image" json:"image,omitempty"`
		PackagePullPolicy           *string            `tfsdk:"package_pull_policy" json:"packagePullPolicy,omitempty"`
		PackagePullSecrets          *[]struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"package_pull_secrets" json:"packagePullSecrets,omitempty"`
		Revision                 *int64 `tfsdk:"revision" json:"revision,omitempty"`
		SkipDependencyResolution *bool  `tfsdk:"skip_dependency_resolution" json:"skipDependencyResolution,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *PkgCrossplaneIoConfigurationRevisionV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_pkg_crossplane_io_configuration_revision_v1_manifest"
}

func (r *PkgCrossplaneIoConfigurationRevisionV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "A ConfigurationRevision represents a revision of a Configuration. Crossplane creates new revisions when there are changes to a Configuration. Crossplane creates and manages ConfigurationRevision. Don't directly edit ConfigurationRevisions.",
		MarkdownDescription: "A ConfigurationRevision represents a revision of a Configuration. Crossplane creates new revisions when there are changes to a Configuration. Crossplane creates and manages ConfigurationRevision. Don't directly edit ConfigurationRevisions.",
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
				Description:         "PackageRevisionSpec specifies the desired state of a PackageRevision.",
				MarkdownDescription: "PackageRevisionSpec specifies the desired state of a PackageRevision.",
				Attributes: map[string]schema.Attribute{
					"common_labels": schema.MapAttribute{
						Description:         "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/",
						MarkdownDescription: "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"desired_state": schema.StringAttribute{
						Description:         "DesiredState of the PackageRevision. Can be either Active or Inactive.",
						MarkdownDescription: "DesiredState of the PackageRevision. Can be either Active or Inactive.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"ignore_crossplane_constraints": schema.BoolAttribute{
						Description:         "IgnoreCrossplaneConstraints indicates to the package manager whether to honor Crossplane version constrains specified by the package. Default is false.",
						MarkdownDescription: "IgnoreCrossplaneConstraints indicates to the package manager whether to honor Crossplane version constrains specified by the package. Default is false.",
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
						Description:         "PackagePullPolicy defines the pull policy for the package. It is also applied to any images pulled for the package, such as a provider's controller image. Default is IfNotPresent.",
						MarkdownDescription: "PackagePullPolicy defines the pull policy for the package. It is also applied to any images pulled for the package, such as a provider's controller image. Default is IfNotPresent.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"package_pull_secrets": schema.ListNestedAttribute{
						Description:         "PackagePullSecrets are named secrets in the same namespace that can be used to fetch packages from private registries. They are also applied to any images pulled for the package, such as a provider's controller image.",
						MarkdownDescription: "PackagePullSecrets are named secrets in the same namespace that can be used to fetch packages from private registries. They are also applied to any images pulled for the package, such as a provider's controller image.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
									MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
						Description:         "Revision number. Indicates when the revision will be garbage collected based on the parent's RevisionHistoryLimit.",
						MarkdownDescription: "Revision number. Indicates when the revision will be garbage collected based on the parent's RevisionHistoryLimit.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"skip_dependency_resolution": schema.BoolAttribute{
						Description:         "SkipDependencyResolution indicates to the package manager whether to skip resolving dependencies for a package. Setting this value to true may have unintended consequences. Default is false.",
						MarkdownDescription: "SkipDependencyResolution indicates to the package manager whether to skip resolving dependencies for a package. Setting this value to true may have unintended consequences. Default is false.",
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

func (r *PkgCrossplaneIoConfigurationRevisionV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_pkg_crossplane_io_configuration_revision_v1_manifest")

	var model PkgCrossplaneIoConfigurationRevisionV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("pkg.crossplane.io/v1")
	model.Kind = pointer.String("ConfigurationRevision")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
