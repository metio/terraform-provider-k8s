/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package pkg_crossplane_io_v1beta1

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
	_ datasource.DataSource = &PkgCrossplaneIoLockV1Beta1Manifest{}
)

func NewPkgCrossplaneIoLockV1Beta1Manifest() datasource.DataSource {
	return &PkgCrossplaneIoLockV1Beta1Manifest{}
}

type PkgCrossplaneIoLockV1Beta1Manifest struct{}

type PkgCrossplaneIoLockV1Beta1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Packages *[]struct {
		Dependencies *[]struct {
			Constraints *string `tfsdk:"constraints" json:"constraints,omitempty"`
			Package     *string `tfsdk:"package" json:"package,omitempty"`
			Type        *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"dependencies" json:"dependencies,omitempty"`
		Name    *string `tfsdk:"name" json:"name,omitempty"`
		Source  *string `tfsdk:"source" json:"source,omitempty"`
		Type    *string `tfsdk:"type" json:"type,omitempty"`
		Version *string `tfsdk:"version" json:"version,omitempty"`
	} `tfsdk:"packages" json:"packages,omitempty"`
}

func (r *PkgCrossplaneIoLockV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_pkg_crossplane_io_lock_v1beta1_manifest"
}

func (r *PkgCrossplaneIoLockV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Lock is the CRD type that tracks package dependencies.",
		MarkdownDescription: "Lock is the CRD type that tracks package dependencies.",
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

			"packages": schema.ListNestedAttribute{
				Description:         "",
				MarkdownDescription: "",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"dependencies": schema.ListNestedAttribute{
							Description:         "Dependencies are the list of dependencies of this package. The order of the dependencies will dictate the order in which they are resolved.",
							MarkdownDescription: "Dependencies are the list of dependencies of this package. The order of the dependencies will dictate the order in which they are resolved.",
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"constraints": schema.StringAttribute{
										Description:         "Constraints is a valid semver range or a digest, which will be used to select a valid dependency version.",
										MarkdownDescription: "Constraints is a valid semver range or a digest, which will be used to select a valid dependency version.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"package": schema.StringAttribute{
										Description:         "Package is the OCI image name without a tag or digest.",
										MarkdownDescription: "Package is the OCI image name without a tag or digest.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"type": schema.StringAttribute{
										Description:         "Type is the type of package. Can be either Configuration or Provider.",
										MarkdownDescription: "Type is the type of package. Can be either Configuration or Provider.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
							},
							Required: true,
							Optional: false,
							Computed: false,
						},

						"name": schema.StringAttribute{
							Description:         "Name corresponds to the name of the package revision for this package.",
							MarkdownDescription: "Name corresponds to the name of the package revision for this package.",
							Required:            true,
							Optional:            false,
							Computed:            false,
						},

						"source": schema.StringAttribute{
							Description:         "Source is the OCI image name without a tag or digest.",
							MarkdownDescription: "Source is the OCI image name without a tag or digest.",
							Required:            true,
							Optional:            false,
							Computed:            false,
						},

						"type": schema.StringAttribute{
							Description:         "Type is the type of package. Can be either Configuration or Provider.",
							MarkdownDescription: "Type is the type of package. Can be either Configuration or Provider.",
							Required:            true,
							Optional:            false,
							Computed:            false,
						},

						"version": schema.StringAttribute{
							Description:         "Version is the tag or digest of the OCI image.",
							MarkdownDescription: "Version is the tag or digest of the OCI image.",
							Required:            true,
							Optional:            false,
							Computed:            false,
						},
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *PkgCrossplaneIoLockV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_pkg_crossplane_io_lock_v1beta1_manifest")

	var model PkgCrossplaneIoLockV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("pkg.crossplane.io/v1beta1")
	model.Kind = pointer.String("Lock")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
