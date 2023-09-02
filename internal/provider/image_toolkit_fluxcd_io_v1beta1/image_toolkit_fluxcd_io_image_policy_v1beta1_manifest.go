/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package image_toolkit_fluxcd_io_v1beta1

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &ImageToolkitFluxcdIoImagePolicyV1Beta1Manifest{}
)

func NewImageToolkitFluxcdIoImagePolicyV1Beta1Manifest() datasource.DataSource {
	return &ImageToolkitFluxcdIoImagePolicyV1Beta1Manifest{}
}

type ImageToolkitFluxcdIoImagePolicyV1Beta1Manifest struct{}

type ImageToolkitFluxcdIoImagePolicyV1Beta1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
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
		FilterTags *struct {
			Extract *string `tfsdk:"extract" json:"extract,omitempty"`
			Pattern *string `tfsdk:"pattern" json:"pattern,omitempty"`
		} `tfsdk:"filter_tags" json:"filterTags,omitempty"`
		ImageRepositoryRef *struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"image_repository_ref" json:"imageRepositoryRef,omitempty"`
		Policy *struct {
			Alphabetical *struct {
				Order *string `tfsdk:"order" json:"order,omitempty"`
			} `tfsdk:"alphabetical" json:"alphabetical,omitempty"`
			Numerical *struct {
				Order *string `tfsdk:"order" json:"order,omitempty"`
			} `tfsdk:"numerical" json:"numerical,omitempty"`
			Semver *struct {
				Range *string `tfsdk:"range" json:"range,omitempty"`
			} `tfsdk:"semver" json:"semver,omitempty"`
		} `tfsdk:"policy" json:"policy,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ImageToolkitFluxcdIoImagePolicyV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_image_toolkit_fluxcd_io_image_policy_v1beta1_manifest"
}

func (r *ImageToolkitFluxcdIoImagePolicyV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ImagePolicy is the Schema for the imagepolicies API",
		MarkdownDescription: "ImagePolicy is the Schema for the imagepolicies API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

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
				Description:         "ImagePolicySpec defines the parameters for calculating the ImagePolicy",
				MarkdownDescription: "ImagePolicySpec defines the parameters for calculating the ImagePolicy",
				Attributes: map[string]schema.Attribute{
					"filter_tags": schema.SingleNestedAttribute{
						Description:         "FilterTags enables filtering for only a subset of tags based on a set of rules. If no rules are provided, all the tags from the repository will be ordered and compared.",
						MarkdownDescription: "FilterTags enables filtering for only a subset of tags based on a set of rules. If no rules are provided, all the tags from the repository will be ordered and compared.",
						Attributes: map[string]schema.Attribute{
							"extract": schema.StringAttribute{
								Description:         "Extract allows a capture group to be extracted from the specified regular expression pattern, useful before tag evaluation.",
								MarkdownDescription: "Extract allows a capture group to be extracted from the specified regular expression pattern, useful before tag evaluation.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pattern": schema.StringAttribute{
								Description:         "Pattern specifies a regular expression pattern used to filter for image tags.",
								MarkdownDescription: "Pattern specifies a regular expression pattern used to filter for image tags.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"image_repository_ref": schema.SingleNestedAttribute{
						Description:         "ImageRepositoryRef points at the object specifying the image being scanned",
						MarkdownDescription: "ImageRepositoryRef points at the object specifying the image being scanned",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of the referent.",
								MarkdownDescription: "Name of the referent.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace of the referent, when not specified it acts as LocalObjectReference.",
								MarkdownDescription: "Namespace of the referent, when not specified it acts as LocalObjectReference.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"policy": schema.SingleNestedAttribute{
						Description:         "Policy gives the particulars of the policy to be followed in selecting the most recent image",
						MarkdownDescription: "Policy gives the particulars of the policy to be followed in selecting the most recent image",
						Attributes: map[string]schema.Attribute{
							"alphabetical": schema.SingleNestedAttribute{
								Description:         "Alphabetical set of rules to use for alphabetical ordering of the tags.",
								MarkdownDescription: "Alphabetical set of rules to use for alphabetical ordering of the tags.",
								Attributes: map[string]schema.Attribute{
									"order": schema.StringAttribute{
										Description:         "Order specifies the sorting order of the tags. Given the letters of the alphabet as tags, ascending order would select Z, and descending order would select A.",
										MarkdownDescription: "Order specifies the sorting order of the tags. Given the letters of the alphabet as tags, ascending order would select Z, and descending order would select A.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("asc", "desc"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"numerical": schema.SingleNestedAttribute{
								Description:         "Numerical set of rules to use for numerical ordering of the tags.",
								MarkdownDescription: "Numerical set of rules to use for numerical ordering of the tags.",
								Attributes: map[string]schema.Attribute{
									"order": schema.StringAttribute{
										Description:         "Order specifies the sorting order of the tags. Given the integer values from 0 to 9 as tags, ascending order would select 9, and descending order would select 0.",
										MarkdownDescription: "Order specifies the sorting order of the tags. Given the integer values from 0 to 9 as tags, ascending order would select 9, and descending order would select 0.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("asc", "desc"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"semver": schema.SingleNestedAttribute{
								Description:         "SemVer gives a semantic version range to check against the tags available.",
								MarkdownDescription: "SemVer gives a semantic version range to check against the tags available.",
								Attributes: map[string]schema.Attribute{
									"range": schema.StringAttribute{
										Description:         "Range gives a semver range for the image tag; the highest version within the range that's a tag yields the latest image.",
										MarkdownDescription: "Range gives a semver range for the image tag; the highest version within the range that's a tag yields the latest image.",
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

func (r *ImageToolkitFluxcdIoImagePolicyV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_image_toolkit_fluxcd_io_image_policy_v1beta1_manifest")

	var model ImageToolkitFluxcdIoImagePolicyV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("image.toolkit.fluxcd.io/v1beta1")
	model.Kind = pointer.String("ImagePolicy")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"YAML Error: "+err.Error(),
		)
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
