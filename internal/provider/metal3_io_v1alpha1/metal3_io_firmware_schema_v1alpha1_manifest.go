/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package metal3_io_v1alpha1

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
	_ datasource.DataSource = &Metal3IoFirmwareSchemaV1Alpha1Manifest{}
)

func NewMetal3IoFirmwareSchemaV1Alpha1Manifest() datasource.DataSource {
	return &Metal3IoFirmwareSchemaV1Alpha1Manifest{}
}

type Metal3IoFirmwareSchemaV1Alpha1Manifest struct{}

type Metal3IoFirmwareSchemaV1Alpha1ManifestData struct {
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
		HardwareModel  *string `tfsdk:"hardware_model" json:"hardwareModel,omitempty"`
		HardwareVendor *string `tfsdk:"hardware_vendor" json:"hardwareVendor,omitempty"`
		Schema         *struct {
			Allowable_values *[]string `tfsdk:"allowable_values" json:"allowable_values,omitempty"`
			Attribute_type   *string   `tfsdk:"attribute_type" json:"attribute_type,omitempty"`
			Lower_bound      *int64    `tfsdk:"lower_bound" json:"lower_bound,omitempty"`
			Max_length       *int64    `tfsdk:"max_length" json:"max_length,omitempty"`
			Min_length       *int64    `tfsdk:"min_length" json:"min_length,omitempty"`
			Read_only        *bool     `tfsdk:"read_only" json:"read_only,omitempty"`
			Unique           *bool     `tfsdk:"unique" json:"unique,omitempty"`
			Upper_bound      *int64    `tfsdk:"upper_bound" json:"upper_bound,omitempty"`
		} `tfsdk:"schema" json:"schema,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *Metal3IoFirmwareSchemaV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_metal3_io_firmware_schema_v1alpha1_manifest"
}

func (r *Metal3IoFirmwareSchemaV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "FirmwareSchema is the Schema for the firmwareschemas API.",
		MarkdownDescription: "FirmwareSchema is the Schema for the firmwareschemas API.",
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
				Description:         "FirmwareSchemaSpec defines the desired state of FirmwareSchema.",
				MarkdownDescription: "FirmwareSchemaSpec defines the desired state of FirmwareSchema.",
				Attributes: map[string]schema.Attribute{
					"hardware_model": schema.StringAttribute{
						Description:         "The hardware model associated with this schema",
						MarkdownDescription: "The hardware model associated with this schema",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"hardware_vendor": schema.StringAttribute{
						Description:         "The hardware vendor associated with this schema",
						MarkdownDescription: "The hardware vendor associated with this schema",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"schema": schema.SingleNestedAttribute{
						Description:         "Map of firmware name to schema",
						MarkdownDescription: "Map of firmware name to schema",
						Attributes: map[string]schema.Attribute{
							"allowable_values": schema.ListAttribute{
								Description:         "The allowable value for an Enumeration type setting.",
								MarkdownDescription: "The allowable value for an Enumeration type setting.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"attribute_type": schema.StringAttribute{
								Description:         "The type of setting.",
								MarkdownDescription: "The type of setting.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Enumeration", "String", "Integer", "Boolean", "Password"),
								},
							},

							"lower_bound": schema.Int64Attribute{
								Description:         "The lowest value for an Integer type setting.",
								MarkdownDescription: "The lowest value for an Integer type setting.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_length": schema.Int64Attribute{
								Description:         "Maximum length for a String type setting.",
								MarkdownDescription: "Maximum length for a String type setting.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"min_length": schema.Int64Attribute{
								Description:         "Minimum length for a String type setting.",
								MarkdownDescription: "Minimum length for a String type setting.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"read_only": schema.BoolAttribute{
								Description:         "Whether or not this setting is read only.",
								MarkdownDescription: "Whether or not this setting is read only.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"unique": schema.BoolAttribute{
								Description:         "Whether or not this setting's value is unique to this node, e.g. a serial number.",
								MarkdownDescription: "Whether or not this setting's value is unique to this node, e.g. a serial number.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"upper_bound": schema.Int64Attribute{
								Description:         "The highest value for an Integer type setting.",
								MarkdownDescription: "The highest value for an Integer type setting.",
								Required:            false,
								Optional:            true,
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

func (r *Metal3IoFirmwareSchemaV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_metal3_io_firmware_schema_v1alpha1_manifest")

	var model Metal3IoFirmwareSchemaV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("metal3.io/v1alpha1")
	model.Kind = pointer.String("FirmwareSchema")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
