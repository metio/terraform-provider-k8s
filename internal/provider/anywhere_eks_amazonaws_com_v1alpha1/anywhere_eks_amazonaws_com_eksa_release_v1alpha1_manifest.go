/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package anywhere_eks_amazonaws_com_v1alpha1

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
	_ datasource.DataSource = &AnywhereEksAmazonawsComEksareleaseV1Alpha1Manifest{}
)

func NewAnywhereEksAmazonawsComEksareleaseV1Alpha1Manifest() datasource.DataSource {
	return &AnywhereEksAmazonawsComEksareleaseV1Alpha1Manifest{}
}

type AnywhereEksAmazonawsComEksareleaseV1Alpha1Manifest struct{}

type AnywhereEksAmazonawsComEksareleaseV1Alpha1ManifestData struct {
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
		BundleManifestUrl *string `tfsdk:"bundle_manifest_url" json:"bundleManifestUrl,omitempty"`
		BundlesRef        *struct {
			ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Name       *string `tfsdk:"name" json:"name,omitempty"`
			Namespace  *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"bundles_ref" json:"bundlesRef,omitempty"`
		GitCommit   *string `tfsdk:"git_commit" json:"gitCommit,omitempty"`
		ReleaseDate *string `tfsdk:"release_date" json:"releaseDate,omitempty"`
		Version     *string `tfsdk:"version" json:"version,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AnywhereEksAmazonawsComEksareleaseV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_anywhere_eks_amazonaws_com_eksa_release_v1alpha1_manifest"
}

func (r *AnywhereEksAmazonawsComEksareleaseV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "EKSARelease is the mapping between release semver of EKS-A and a Bundles resource on the cluster.",
		MarkdownDescription: "EKSARelease is the mapping between release semver of EKS-A and a Bundles resource on the cluster.",
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
				Description:         "EKSAReleaseSpec defines the desired state of EKSARelease.",
				MarkdownDescription: "EKSAReleaseSpec defines the desired state of EKSARelease.",
				Attributes: map[string]schema.Attribute{
					"bundle_manifest_url": schema.StringAttribute{
						Description:         "Manifest url to parse bundle information from for this EKS-A release",
						MarkdownDescription: "Manifest url to parse bundle information from for this EKS-A release",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"bundles_ref": schema.SingleNestedAttribute{
						Description:         "Reference to a Bundles resource in the cluster",
						MarkdownDescription: "Reference to a Bundles resource in the cluster",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "APIVersion refers to the Bundles APIVersion",
								MarkdownDescription: "APIVersion refers to the Bundles APIVersion",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name refers to the name of the Bundles object in the cluster",
								MarkdownDescription: "Name refers to the name of the Bundles object in the cluster",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace refers to the Bundles's namespace",
								MarkdownDescription: "Namespace refers to the Bundles's namespace",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"git_commit": schema.StringAttribute{
						Description:         "Git commit the component is built from, before any patches",
						MarkdownDescription: "Git commit the component is built from, before any patches",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"release_date": schema.StringAttribute{
						Description:         "Date of EKS-A Release",
						MarkdownDescription: "Date of EKS-A Release",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"version": schema.StringAttribute{
						Description:         "EKS-A release semantic version",
						MarkdownDescription: "EKS-A release semantic version",
						Required:            true,
						Optional:            false,
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

func (r *AnywhereEksAmazonawsComEksareleaseV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_anywhere_eks_amazonaws_com_eksa_release_v1alpha1_manifest")

	var model AnywhereEksAmazonawsComEksareleaseV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("anywhere.eks.amazonaws.com/v1alpha1")
	model.Kind = pointer.String("EKSARelease")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
