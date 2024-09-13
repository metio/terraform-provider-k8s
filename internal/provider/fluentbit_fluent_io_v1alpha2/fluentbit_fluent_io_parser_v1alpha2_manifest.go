/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package fluentbit_fluent_io_v1alpha2

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
	_ datasource.DataSource = &FluentbitFluentIoParserV1Alpha2Manifest{}
)

func NewFluentbitFluentIoParserV1Alpha2Manifest() datasource.DataSource {
	return &FluentbitFluentIoParserV1Alpha2Manifest{}
}

type FluentbitFluentIoParserV1Alpha2Manifest struct{}

type FluentbitFluentIoParserV1Alpha2ManifestData struct {
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
		Decoders *[]struct {
			DecodeField   *string `tfsdk:"decode_field" json:"decodeField,omitempty"`
			DecodeFieldAs *string `tfsdk:"decode_field_as" json:"decodeFieldAs,omitempty"`
		} `tfsdk:"decoders" json:"decoders,omitempty"`
		Json *struct {
			TimeFormat *string `tfsdk:"time_format" json:"timeFormat,omitempty"`
			TimeKeep   *bool   `tfsdk:"time_keep" json:"timeKeep,omitempty"`
			TimeKey    *string `tfsdk:"time_key" json:"timeKey,omitempty"`
		} `tfsdk:"json" json:"json,omitempty"`
		Logfmt *map[string]string `tfsdk:"logfmt" json:"logfmt,omitempty"`
		Ltsv   *struct {
			TimeFormat *string `tfsdk:"time_format" json:"timeFormat,omitempty"`
			TimeKeep   *bool   `tfsdk:"time_keep" json:"timeKeep,omitempty"`
			TimeKey    *string `tfsdk:"time_key" json:"timeKey,omitempty"`
			Types      *string `tfsdk:"types" json:"types,omitempty"`
		} `tfsdk:"ltsv" json:"ltsv,omitempty"`
		Regex *struct {
			Regex      *string `tfsdk:"regex" json:"regex,omitempty"`
			TimeFormat *string `tfsdk:"time_format" json:"timeFormat,omitempty"`
			TimeKeep   *bool   `tfsdk:"time_keep" json:"timeKeep,omitempty"`
			TimeKey    *string `tfsdk:"time_key" json:"timeKey,omitempty"`
			TimeOffset *string `tfsdk:"time_offset" json:"timeOffset,omitempty"`
			Types      *string `tfsdk:"types" json:"types,omitempty"`
		} `tfsdk:"regex" json:"regex,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *FluentbitFluentIoParserV1Alpha2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_fluentbit_fluent_io_parser_v1alpha2_manifest"
}

func (r *FluentbitFluentIoParserV1Alpha2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Parser is the Schema for namespace level parser API",
		MarkdownDescription: "Parser is the Schema for namespace level parser API",
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
				Description:         "ParserSpec defines the desired state of ClusterParser",
				MarkdownDescription: "ParserSpec defines the desired state of ClusterParser",
				Attributes: map[string]schema.Attribute{
					"decoders": schema.ListNestedAttribute{
						Description:         "Decoders are a built-in feature available through the Parsers file, each Parser definition can optionally set one or multiple decoders. There are two type of decoders type: Decode_Field and Decode_Field_As.",
						MarkdownDescription: "Decoders are a built-in feature available through the Parsers file, each Parser definition can optionally set one or multiple decoders. There are two type of decoders type: Decode_Field and Decode_Field_As.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"decode_field": schema.StringAttribute{
									Description:         "If the content can be decoded in a structured message, append that structure message (keys and values) to the original log message.",
									MarkdownDescription: "If the content can be decoded in a structured message, append that structure message (keys and values) to the original log message.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"decode_field_as": schema.StringAttribute{
									Description:         "Any content decoded (unstructured or structured) will be replaced in the same key/value, no extra keys are added.",
									MarkdownDescription: "Any content decoded (unstructured or structured) will be replaced in the same key/value, no extra keys are added.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"json": schema.SingleNestedAttribute{
						Description:         "JSON defines json parser configuration.",
						MarkdownDescription: "JSON defines json parser configuration.",
						Attributes: map[string]schema.Attribute{
							"time_format": schema.StringAttribute{
								Description:         "Time_Format, eg. %Y-%m-%dT%H:%M:%S %z",
								MarkdownDescription: "Time_Format, eg. %Y-%m-%dT%H:%M:%S %z",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"time_keep": schema.BoolAttribute{
								Description:         "Time_Keep",
								MarkdownDescription: "Time_Keep",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"time_key": schema.StringAttribute{
								Description:         "Time_Key",
								MarkdownDescription: "Time_Key",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"logfmt": schema.MapAttribute{
						Description:         "Logfmt defines logfmt parser configuration.",
						MarkdownDescription: "Logfmt defines logfmt parser configuration.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ltsv": schema.SingleNestedAttribute{
						Description:         "LTSV defines ltsv parser configuration.",
						MarkdownDescription: "LTSV defines ltsv parser configuration.",
						Attributes: map[string]schema.Attribute{
							"time_format": schema.StringAttribute{
								Description:         "Time_Format, eg. %Y-%m-%dT%H:%M:%S %z",
								MarkdownDescription: "Time_Format, eg. %Y-%m-%dT%H:%M:%S %z",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"time_keep": schema.BoolAttribute{
								Description:         "Time_Keep",
								MarkdownDescription: "Time_Keep",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"time_key": schema.StringAttribute{
								Description:         "Time_Key",
								MarkdownDescription: "Time_Key",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"types": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"regex": schema.SingleNestedAttribute{
						Description:         "Regex defines regex parser configuration.",
						MarkdownDescription: "Regex defines regex parser configuration.",
						Attributes: map[string]schema.Attribute{
							"regex": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"time_format": schema.StringAttribute{
								Description:         "Time_Format, eg. %Y-%m-%dT%H:%M:%S %z",
								MarkdownDescription: "Time_Format, eg. %Y-%m-%dT%H:%M:%S %z",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"time_keep": schema.BoolAttribute{
								Description:         "Time_Keep",
								MarkdownDescription: "Time_Keep",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"time_key": schema.StringAttribute{
								Description:         "Time_Key",
								MarkdownDescription: "Time_Key",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"time_offset": schema.StringAttribute{
								Description:         "Time_Offset, eg. +0200",
								MarkdownDescription: "Time_Offset, eg. +0200",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"types": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
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
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *FluentbitFluentIoParserV1Alpha2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_fluentbit_fluent_io_parser_v1alpha2_manifest")

	var model FluentbitFluentIoParserV1Alpha2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("fluentbit.fluent.io/v1alpha2")
	model.Kind = pointer.String("Parser")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
