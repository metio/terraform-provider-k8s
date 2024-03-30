/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package capabilities_3scale_net_v1beta1

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
	_ datasource.DataSource = &Capabilities3ScaleNetCustomPolicyDefinitionV1Beta1Manifest{}
)

func NewCapabilities3ScaleNetCustomPolicyDefinitionV1Beta1Manifest() datasource.DataSource {
	return &Capabilities3ScaleNetCustomPolicyDefinitionV1Beta1Manifest{}
}

type Capabilities3ScaleNetCustomPolicyDefinitionV1Beta1Manifest struct{}

type Capabilities3ScaleNetCustomPolicyDefinitionV1Beta1ManifestData struct {
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
		Name               *string `tfsdk:"name" json:"name,omitempty"`
		ProviderAccountRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"provider_account_ref" json:"providerAccountRef,omitempty"`
		Schema *struct {
			Dollarschema  *string            `tfsdk:"dollarschema" json:"$schema,omitempty"`
			Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
			Description   *[]string          `tfsdk:"description" json:"description,omitempty"`
			Name          *string            `tfsdk:"name" json:"name,omitempty"`
			Summary       *string            `tfsdk:"summary" json:"summary,omitempty"`
			Version       *string            `tfsdk:"version" json:"version,omitempty"`
		} `tfsdk:"schema" json:"schema,omitempty"`
		Version *string `tfsdk:"version" json:"version,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *Capabilities3ScaleNetCustomPolicyDefinitionV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_capabilities_3scale_net_custom_policy_definition_v1beta1_manifest"
}

func (r *Capabilities3ScaleNetCustomPolicyDefinitionV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CustomPolicyDefinition is the Schema for the custompolicydefinitions API",
		MarkdownDescription: "CustomPolicyDefinition is the Schema for the custompolicydefinitions API",
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
				Description:         "CustomPolicyDefinitionSpec defines the desired state of CustomPolicyDefinition",
				MarkdownDescription: "CustomPolicyDefinitionSpec defines the desired state of CustomPolicyDefinition",
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Description:         "Name is the name of the custom policy",
						MarkdownDescription: "Name is the name of the custom policy",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"provider_account_ref": schema.SingleNestedAttribute{
						Description:         "ProviderAccountRef references account provider credentials",
						MarkdownDescription: "ProviderAccountRef references account provider credentials",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
								MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"schema": schema.SingleNestedAttribute{
						Description:         "Schema is the schema of the custom policy",
						MarkdownDescription: "Schema is the schema of the custom policy",
						Attributes: map[string]schema.Attribute{
							"dollarschema": schema.StringAttribute{
								Description:         "Schema the $schema keyword is used to declare that this is a JSON Schema.",
								MarkdownDescription: "Schema the $schema keyword is used to declare that this is a JSON Schema.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"configuration": schema.MapAttribute{
								Description:         "Configuration defines the structural schema for the policy",
								MarkdownDescription: "Configuration defines the structural schema for the policy",
								ElementType:         types.StringType,
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"description": schema.ListAttribute{
								Description:         "Description is an array of description messages for the custom policy schema",
								MarkdownDescription: "Description is an array of description messages for the custom policy schema",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name is the name of the custom policy schema",
								MarkdownDescription: "Name is the name of the custom policy schema",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"summary": schema.StringAttribute{
								Description:         "Summary is the summary of the custom policy schema",
								MarkdownDescription: "Summary is the summary of the custom policy schema",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"version": schema.StringAttribute{
								Description:         "Version is the version of the custom policy schema",
								MarkdownDescription: "Version is the version of the custom policy schema",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"version": schema.StringAttribute{
						Description:         "Version is the version of the custom policy",
						MarkdownDescription: "Version is the version of the custom policy",
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

func (r *Capabilities3ScaleNetCustomPolicyDefinitionV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_capabilities_3scale_net_custom_policy_definition_v1beta1_manifest")

	var model Capabilities3ScaleNetCustomPolicyDefinitionV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("capabilities.3scale.net/v1beta1")
	model.Kind = pointer.String("CustomPolicyDefinition")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
