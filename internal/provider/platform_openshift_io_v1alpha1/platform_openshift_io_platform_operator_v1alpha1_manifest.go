/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package platform_openshift_io_v1alpha1

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &PlatformOpenshiftIoPlatformOperatorV1Alpha1Manifest{}
)

func NewPlatformOpenshiftIoPlatformOperatorV1Alpha1Manifest() datasource.DataSource {
	return &PlatformOpenshiftIoPlatformOperatorV1Alpha1Manifest{}
}

type PlatformOpenshiftIoPlatformOperatorV1Alpha1Manifest struct{}

type PlatformOpenshiftIoPlatformOperatorV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Package *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"package" json:"package,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *PlatformOpenshiftIoPlatformOperatorV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_platform_openshift_io_platform_operator_v1alpha1_manifest"
}

func (r *PlatformOpenshiftIoPlatformOperatorV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "PlatformOperator is the Schema for the PlatformOperators API.  Compatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.",
		MarkdownDescription: "PlatformOperator is the Schema for the PlatformOperators API.  Compatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.",
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
				Description:         "PlatformOperatorSpec defines the desired state of PlatformOperator.",
				MarkdownDescription: "PlatformOperatorSpec defines the desired state of PlatformOperator.",
				Attributes: map[string]schema.Attribute{
					"package": schema.SingleNestedAttribute{
						Description:         "package contains the desired package and its configuration for this PlatformOperator.",
						MarkdownDescription: "package contains the desired package and its configuration for this PlatformOperator.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "name contains the desired OLM-based Operator package name that is defined in an existing CatalogSource resource in the cluster.  This configured package will be managed with the cluster's lifecycle. In the current implementation, it will be retrieving this name from a list of supported operators out of the catalogs included with OpenShift.  ---",
								MarkdownDescription: "name contains the desired OLM-based Operator package name that is defined in an existing CatalogSource resource in the cluster.  This configured package will be managed with the cluster's lifecycle. In the current implementation, it will be retrieving this name from a list of supported operators out of the catalogs included with OpenShift.  ---",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtMost(56),
									stringvalidator.RegexMatches(regexp.MustCompile(`[a-z0-9]([-a-z0-9]*[a-z0-9])?`), ""),
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *PlatformOpenshiftIoPlatformOperatorV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_platform_openshift_io_platform_operator_v1alpha1_manifest")

	var model PlatformOpenshiftIoPlatformOperatorV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("platform.openshift.io/v1alpha1")
	model.Kind = pointer.String("PlatformOperator")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
