/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package apps_clusternet_io_v1alpha1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
	_ datasource.DataSource = &AppsClusternetIoLocalizationV1Alpha1Manifest{}
)

func NewAppsClusternetIoLocalizationV1Alpha1Manifest() datasource.DataSource {
	return &AppsClusternetIoLocalizationV1Alpha1Manifest{}
}

type AppsClusternetIoLocalizationV1Alpha1Manifest struct{}

type AppsClusternetIoLocalizationV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Feed *struct {
			ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
			Name       *string `tfsdk:"name" json:"name,omitempty"`
			Namespace  *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"feed" json:"feed,omitempty"`
		OverridePolicy *string `tfsdk:"override_policy" json:"overridePolicy,omitempty"`
		Overrides      *[]struct {
			Name          *string `tfsdk:"name" json:"name,omitempty"`
			OverrideChart *bool   `tfsdk:"override_chart" json:"overrideChart,omitempty"`
			Type          *string `tfsdk:"type" json:"type,omitempty"`
			Value         *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"overrides" json:"overrides,omitempty"`
		Priority *int64 `tfsdk:"priority" json:"priority,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AppsClusternetIoLocalizationV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_apps_clusternet_io_localization_v1alpha1_manifest"
}

func (r *AppsClusternetIoLocalizationV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Localization represents the override config for a group of resources.",
		MarkdownDescription: "Localization represents the override config for a group of resources.",
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
				Description:         "LocalizationSpec defines the desired state of Localization",
				MarkdownDescription: "LocalizationSpec defines the desired state of Localization",
				Attributes: map[string]schema.Attribute{
					"feed": schema.SingleNestedAttribute{
						Description:         "Feed holds references to the objects the Localization applies to.",
						MarkdownDescription: "Feed holds references to the objects the Localization applies to.",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "APIVersion defines the versioned schema of this representation of an object.",
								MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind is a string value representing the REST resource this object represents.In CamelCase.",
								MarkdownDescription: "Kind is a string value representing the REST resource this object represents.In CamelCase.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the target resource.",
								MarkdownDescription: "Name of the target resource.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace of the target resource.",
								MarkdownDescription: "Namespace of the target resource.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"override_policy": schema.StringAttribute{
						Description:         "OverridePolicy specifies the override policy for this Localization.",
						MarkdownDescription: "OverridePolicy specifies the override policy for this Localization.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("ApplyNow", "ApplyLater"),
						},
					},

					"overrides": schema.ListNestedAttribute{
						Description:         "Overrides holds all the OverrideConfig.",
						MarkdownDescription: "Overrides holds all the OverrideConfig.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name indicate the OverrideConfig name.",
									MarkdownDescription: "Name indicate the OverrideConfig name.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"override_chart": schema.BoolAttribute{
									Description:         "OverrideChart indicates whether the override value for the HelmChart CR.",
									MarkdownDescription: "OverrideChart indicates whether the override value for the HelmChart CR.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"type": schema.StringAttribute{
									Description:         "Type specifies the override type for override value.",
									MarkdownDescription: "Type specifies the override type for override value.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("Helm", "JSONPatch", "MergePatch"),
									},
								},

								"value": schema.StringAttribute{
									Description:         "Value represents override value.",
									MarkdownDescription: "Value represents override value.",
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

					"priority": schema.Int64Attribute{
						Description:         "Priority is an integer defining the relative importance of this Localization compared to others.Lower numbers are considered lower priority.And these Localization(s) will be applied by order from lower priority to higher.That means override values in lower Localization will be overridden by those in higher Localization.",
						MarkdownDescription: "Priority is an integer defining the relative importance of this Localization compared to others.Lower numbers are considered lower priority.And these Localization(s) will be applied by order from lower priority to higher.That means override values in lower Localization will be overridden by those in higher Localization.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
							int64validator.AtMost(1000),
						},
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *AppsClusternetIoLocalizationV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_apps_clusternet_io_localization_v1alpha1_manifest")

	var model AppsClusternetIoLocalizationV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("apps.clusternet.io/v1alpha1")
	model.Kind = pointer.String("Localization")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
