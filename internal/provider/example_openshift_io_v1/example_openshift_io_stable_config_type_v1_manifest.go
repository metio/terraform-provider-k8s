/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package example_openshift_io_v1

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
	_ datasource.DataSource = &ExampleOpenshiftIoStableConfigTypeV1Manifest{}
)

func NewExampleOpenshiftIoStableConfigTypeV1Manifest() datasource.DataSource {
	return &ExampleOpenshiftIoStableConfigTypeV1Manifest{}
}

type ExampleOpenshiftIoStableConfigTypeV1Manifest struct{}

type ExampleOpenshiftIoStableConfigTypeV1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		CelUnion *struct {
			OptionalMember *string `tfsdk:"optional_member" json:"optionalMember,omitempty"`
			RequiredMember *string `tfsdk:"required_member" json:"requiredMember,omitempty"`
			Type           *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"cel_union" json:"celUnion,omitempty"`
		EvolvingUnion *struct {
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"evolving_union" json:"evolvingUnion,omitempty"`
		ImmutableField         *string `tfsdk:"immutable_field" json:"immutableField,omitempty"`
		OptionalImmutableField *string `tfsdk:"optional_immutable_field" json:"optionalImmutableField,omitempty"`
		StableField            *string `tfsdk:"stable_field" json:"stableField,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ExampleOpenshiftIoStableConfigTypeV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_example_openshift_io_stable_config_type_v1_manifest"
}

func (r *ExampleOpenshiftIoStableConfigTypeV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "StableConfigType is a stable config type that may include TechPreviewNoUpgrade fields.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		MarkdownDescription: "StableConfigType is a stable config type that may include TechPreviewNoUpgrade fields.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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
				Description:         "spec is the specification of the desired behavior of the StableConfigType.",
				MarkdownDescription: "spec is the specification of the desired behavior of the StableConfigType.",
				Attributes: map[string]schema.Attribute{
					"cel_union": schema.SingleNestedAttribute{
						Description:         "celUnion demonstrates how to validate a discrminated union using CEL",
						MarkdownDescription: "celUnion demonstrates how to validate a discrminated union using CEL",
						Attributes: map[string]schema.Attribute{
							"optional_member": schema.StringAttribute{
								Description:         "optionalMember is a union member that is optional.",
								MarkdownDescription: "optionalMember is a union member that is optional.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"required_member": schema.StringAttribute{
								Description:         "requiredMember is a union member that is required.",
								MarkdownDescription: "requiredMember is a union member that is required.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "type determines which of the union members should be populated.",
								MarkdownDescription: "type determines which of the union members should be populated.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("RequiredMember", "OptionalMember", "EmptyMember"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"evolving_union": schema.SingleNestedAttribute{
						Description:         "evolvingUnion demonstrates how to phase in new values into discriminated union",
						MarkdownDescription: "evolvingUnion demonstrates how to phase in new values into discriminated union",
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Description:         "type is the discriminator. It has different values for Default and for TechPreviewNoUpgrade",
								MarkdownDescription: "type is the discriminator. It has different values for Default and for TechPreviewNoUpgrade",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("", "StableValue"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"immutable_field": schema.StringAttribute{
						Description:         "immutableField is a field that is immutable once the object has been created. It is required at all times.",
						MarkdownDescription: "immutableField is a field that is immutable once the object has been created. It is required at all times.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"optional_immutable_field": schema.StringAttribute{
						Description:         "optionalImmutableField is a field that is immutable once set. It is optional but may not be changed once set.",
						MarkdownDescription: "optionalImmutableField is a field that is immutable once set. It is optional but may not be changed once set.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"stable_field": schema.StringAttribute{
						Description:         "stableField is a field that is present on default clusters and on tech preview clusters  If empty, the platform will choose a good default, which may change over time without notice.",
						MarkdownDescription: "stableField is a field that is present on default clusters and on tech preview clusters  If empty, the platform will choose a good default, which may change over time without notice.",
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

func (r *ExampleOpenshiftIoStableConfigTypeV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_example_openshift_io_stable_config_type_v1_manifest")

	var model ExampleOpenshiftIoStableConfigTypeV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("example.openshift.io/v1")
	model.Kind = pointer.String("StableConfigType")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
