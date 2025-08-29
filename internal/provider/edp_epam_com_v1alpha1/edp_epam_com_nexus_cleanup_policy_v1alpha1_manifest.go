/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package edp_epam_com_v1alpha1

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
	_ datasource.DataSource = &EdpEpamComNexusCleanupPolicyV1Alpha1Manifest{}
)

func NewEdpEpamComNexusCleanupPolicyV1Alpha1Manifest() datasource.DataSource {
	return &EdpEpamComNexusCleanupPolicyV1Alpha1Manifest{}
}

type EdpEpamComNexusCleanupPolicyV1Alpha1Manifest struct{}

type EdpEpamComNexusCleanupPolicyV1Alpha1ManifestData struct {
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
		Criteria *struct {
			AssetRegex      *string `tfsdk:"asset_regex" json:"assetRegex,omitempty"`
			LastBlobUpdated *int64  `tfsdk:"last_blob_updated" json:"lastBlobUpdated,omitempty"`
			LastDownloaded  *int64  `tfsdk:"last_downloaded" json:"lastDownloaded,omitempty"`
			ReleaseType     *string `tfsdk:"release_type" json:"releaseType,omitempty"`
		} `tfsdk:"criteria" json:"criteria,omitempty"`
		Description *string `tfsdk:"description" json:"description,omitempty"`
		Format      *string `tfsdk:"format" json:"format,omitempty"`
		Name        *string `tfsdk:"name" json:"name,omitempty"`
		NexusRef    *struct {
			Kind *string `tfsdk:"kind" json:"kind,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"nexus_ref" json:"nexusRef,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *EdpEpamComNexusCleanupPolicyV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_edp_epam_com_nexus_cleanup_policy_v1alpha1_manifest"
}

func (r *EdpEpamComNexusCleanupPolicyV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "NexusCleanupPolicy is the Schema for the cleanuppolicies API.",
		MarkdownDescription: "NexusCleanupPolicy is the Schema for the cleanuppolicies API.",
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
				Description:         "NexusCleanupPolicySpec defines the desired state of NexusCleanupPolicy.",
				MarkdownDescription: "NexusCleanupPolicySpec defines the desired state of NexusCleanupPolicy.",
				Attributes: map[string]schema.Attribute{
					"criteria": schema.SingleNestedAttribute{
						Description:         "Criteria for the cleanup policy.",
						MarkdownDescription: "Criteria for the cleanup policy.",
						Attributes: map[string]schema.Attribute{
							"asset_regex": schema.StringAttribute{
								Description:         "AssetRegex removes components that match the given regex.",
								MarkdownDescription: "AssetRegex removes components that match the given regex.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"last_blob_updated": schema.Int64Attribute{
								Description:         "LastBlobUpdated removes components published over “x” days ago.",
								MarkdownDescription: "LastBlobUpdated removes components published over “x” days ago.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(24855),
								},
							},

							"last_downloaded": schema.Int64Attribute{
								Description:         "LastDownloaded removes components downloaded over “x” days.",
								MarkdownDescription: "LastDownloaded removes components downloaded over “x” days.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(24855),
								},
							},

							"release_type": schema.StringAttribute{
								Description:         "ReleaseType removes components that are of the following release type.",
								MarkdownDescription: "ReleaseType removes components that are of the following release type.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("RELEASES", "PRERELEASES", ""),
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"description": schema.StringAttribute{
						Description:         "Description of the cleanup policy.",
						MarkdownDescription: "Description of the cleanup policy.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"format": schema.StringAttribute{
						Description:         "Format that this cleanup policy can be applied to.",
						MarkdownDescription: "Format that this cleanup policy can be applied to.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("apt", "bower", "cocoapods", "conan", "conda", "docker", "gitlfs", "go", "helm", "maven2", "npm", "nuget", "p2", "pypi", "r", "raw", "rubygems", "yum"),
						},
					},

					"name": schema.StringAttribute{
						Description:         "Name is a unique name for the cleanup policy.",
						MarkdownDescription: "Name is a unique name for the cleanup policy.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(512),
						},
					},

					"nexus_ref": schema.SingleNestedAttribute{
						Description:         "NexusRef is a reference to Nexus custom resource.",
						MarkdownDescription: "NexusRef is a reference to Nexus custom resource.",
						Attributes: map[string]schema.Attribute{
							"kind": schema.StringAttribute{
								Description:         "Kind specifies the kind of the Nexus resource.",
								MarkdownDescription: "Kind specifies the kind of the Nexus resource.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name specifies the name of the Nexus resource.",
								MarkdownDescription: "Name specifies the name of the Nexus resource.",
								Required:            true,
								Optional:            false,
								Computed:            false,
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

func (r *EdpEpamComNexusCleanupPolicyV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_edp_epam_com_nexus_cleanup_policy_v1alpha1_manifest")

	var model EdpEpamComNexusCleanupPolicyV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("edp.epam.com/v1alpha1")
	model.Kind = pointer.String("NexusCleanupPolicy")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
