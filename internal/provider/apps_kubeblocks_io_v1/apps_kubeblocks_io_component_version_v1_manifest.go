/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package apps_kubeblocks_io_v1

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
	_ datasource.DataSource = &AppsKubeblocksIoComponentVersionV1Manifest{}
)

func NewAppsKubeblocksIoComponentVersionV1Manifest() datasource.DataSource {
	return &AppsKubeblocksIoComponentVersionV1Manifest{}
}

type AppsKubeblocksIoComponentVersionV1Manifest struct{}

type AppsKubeblocksIoComponentVersionV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		CompatibilityRules *[]struct {
			CompDefs *[]string `tfsdk:"comp_defs" json:"compDefs,omitempty"`
			Releases *[]string `tfsdk:"releases" json:"releases,omitempty"`
		} `tfsdk:"compatibility_rules" json:"compatibilityRules,omitempty"`
		Releases *[]struct {
			Changes        *string            `tfsdk:"changes" json:"changes,omitempty"`
			Images         *map[string]string `tfsdk:"images" json:"images,omitempty"`
			Name           *string            `tfsdk:"name" json:"name,omitempty"`
			ServiceVersion *string            `tfsdk:"service_version" json:"serviceVersion,omitempty"`
		} `tfsdk:"releases" json:"releases,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AppsKubeblocksIoComponentVersionV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_apps_kubeblocks_io_component_version_v1_manifest"
}

func (r *AppsKubeblocksIoComponentVersionV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ComponentVersion is the Schema for the componentversions API",
		MarkdownDescription: "ComponentVersion is the Schema for the componentversions API",
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
				Description:         "ComponentVersionSpec defines the desired state of ComponentVersion",
				MarkdownDescription: "ComponentVersionSpec defines the desired state of ComponentVersion",
				Attributes: map[string]schema.Attribute{
					"compatibility_rules": schema.ListNestedAttribute{
						Description:         "CompatibilityRules defines compatibility rules between sets of component definitions and releases.",
						MarkdownDescription: "CompatibilityRules defines compatibility rules between sets of component definitions and releases.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"comp_defs": schema.ListAttribute{
									Description:         "CompDefs specifies names for the component definitions associated with this ComponentVersion. Each name in the list can represent an exact name, a name prefix, or a regular expression pattern. For example: - 'mysql-8.0.30-v1alpha1': Matches the exact name 'mysql-8.0.30-v1alpha1' - 'mysql-8.0.30': Matches all names starting with 'mysql-8.0.30' - '^mysql-8.0.d{1,2}$': Matches all names starting with 'mysql-8.0.' followed by one or two digits.",
									MarkdownDescription: "CompDefs specifies names for the component definitions associated with this ComponentVersion. Each name in the list can represent an exact name, a name prefix, or a regular expression pattern. For example: - 'mysql-8.0.30-v1alpha1': Matches the exact name 'mysql-8.0.30-v1alpha1' - 'mysql-8.0.30': Matches all names starting with 'mysql-8.0.30' - '^mysql-8.0.d{1,2}$': Matches all names starting with 'mysql-8.0.' followed by one or two digits.",
									ElementType:         types.StringType,
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"releases": schema.ListAttribute{
									Description:         "Releases is a list of identifiers for the releases.",
									MarkdownDescription: "Releases is a list of identifiers for the releases.",
									ElementType:         types.StringType,
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

					"releases": schema.ListNestedAttribute{
						Description:         "Releases represents different releases of component instances within this ComponentVersion.",
						MarkdownDescription: "Releases represents different releases of component instances within this ComponentVersion.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"changes": schema.StringAttribute{
									Description:         "Changes provides information about the changes made in this release.",
									MarkdownDescription: "Changes provides information about the changes made in this release.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(256),
									},
								},

								"images": schema.MapAttribute{
									Description:         "Images define the new images for containers, actions or external applications within the release. If an image is specified for a lifecycle action, the key should be the field name (case-insensitive) of the action in the LifecycleActions struct.",
									MarkdownDescription: "Images define the new images for containers, actions or external applications within the release. If an image is specified for a lifecycle action, the key should be the field name (case-insensitive) of the action in the LifecycleActions struct.",
									ElementType:         types.StringType,
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name is a unique identifier for this release. Cannot be updated.",
									MarkdownDescription: "Name is a unique identifier for this release. Cannot be updated.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(32),
									},
								},

								"service_version": schema.StringAttribute{
									Description:         "ServiceVersion defines the version of the well-known service that the component provides. The version should follow the syntax and semantics of the 'Semantic Versioning' specification (http://semver.org/). If the release is used, it will serve as the service version for component instances, overriding the one defined in the component definition. Cannot be updated.",
									MarkdownDescription: "ServiceVersion defines the version of the well-known service that the component provides. The version should follow the syntax and semantics of the 'Semantic Versioning' specification (http://semver.org/). If the release is used, it will serve as the service version for component instances, overriding the one defined in the component definition. Cannot be updated.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(32),
									},
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *AppsKubeblocksIoComponentVersionV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_apps_kubeblocks_io_component_version_v1_manifest")

	var model AppsKubeblocksIoComponentVersionV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("apps.kubeblocks.io/v1")
	model.Kind = pointer.String("ComponentVersion")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
