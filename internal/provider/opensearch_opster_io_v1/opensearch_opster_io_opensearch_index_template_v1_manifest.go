/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package opensearch_opster_io_v1

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
	_ datasource.DataSource = &OpensearchOpsterIoOpensearchIndexTemplateV1Manifest{}
)

func NewOpensearchOpsterIoOpensearchIndexTemplateV1Manifest() datasource.DataSource {
	return &OpensearchOpsterIoOpensearchIndexTemplateV1Manifest{}
}

type OpensearchOpsterIoOpensearchIndexTemplateV1Manifest struct{}

type OpensearchOpsterIoOpensearchIndexTemplateV1ManifestData struct {
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
		_meta      *map[string]string `tfsdk:"_meta" json:"_meta,omitempty"`
		ComposedOf *[]string          `tfsdk:"composed_of" json:"composedOf,omitempty"`
		DataStream *struct {
			Timestamp_field *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"timestamp_field" json:"timestamp_field,omitempty"`
		} `tfsdk:"data_stream" json:"dataStream,omitempty"`
		IndexPatterns     *[]string `tfsdk:"index_patterns" json:"indexPatterns,omitempty"`
		Name              *string   `tfsdk:"name" json:"name,omitempty"`
		OpensearchCluster *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"opensearch_cluster" json:"opensearchCluster,omitempty"`
		Priority *int64 `tfsdk:"priority" json:"priority,omitempty"`
		Template *struct {
			Aliases *struct {
				Alias        *string            `tfsdk:"alias" json:"alias,omitempty"`
				Filter       *map[string]string `tfsdk:"filter" json:"filter,omitempty"`
				Index        *string            `tfsdk:"index" json:"index,omitempty"`
				IsWriteIndex *bool              `tfsdk:"is_write_index" json:"isWriteIndex,omitempty"`
				Routing      *string            `tfsdk:"routing" json:"routing,omitempty"`
			} `tfsdk:"aliases" json:"aliases,omitempty"`
			Mappings *map[string]string `tfsdk:"mappings" json:"mappings,omitempty"`
			Settings *map[string]string `tfsdk:"settings" json:"settings,omitempty"`
		} `tfsdk:"template" json:"template,omitempty"`
		Version *int64 `tfsdk:"version" json:"version,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OpensearchOpsterIoOpensearchIndexTemplateV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_opensearch_opster_io_opensearch_index_template_v1_manifest"
}

func (r *OpensearchOpsterIoOpensearchIndexTemplateV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "OpensearchIndexTemplate is the schema for the OpenSearch index templates API",
		MarkdownDescription: "OpensearchIndexTemplate is the schema for the OpenSearch index templates API",
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"_meta": schema.MapAttribute{
						Description:         "Optional user metadata about the index template",
						MarkdownDescription: "Optional user metadata about the index template",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"composed_of": schema.ListAttribute{
						Description:         "An ordered list of component template names. Component templates are merged in the order specified, meaning that the last component template specified has the highest precedence",
						MarkdownDescription: "An ordered list of component template names. Component templates are merged in the order specified, meaning that the last component template specified has the highest precedence",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"data_stream": schema.SingleNestedAttribute{
						Description:         "The dataStream config that should be applied",
						MarkdownDescription: "The dataStream config that should be applied",
						Attributes: map[string]schema.Attribute{
							"timestamp_field": schema.SingleNestedAttribute{
								Description:         "TimestampField for dataStream",
								MarkdownDescription: "TimestampField for dataStream",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of the field that are used for the DataStream",
										MarkdownDescription: "Name of the field that are used for the DataStream",
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"index_patterns": schema.ListAttribute{
						Description:         "Array of wildcard expressions used to match the names of indices during creation",
						MarkdownDescription: "Array of wildcard expressions used to match the names of indices during creation",
						ElementType:         types.StringType,
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"name": schema.StringAttribute{
						Description:         "The name of the index template. Defaults to metadata.name",
						MarkdownDescription: "The name of the index template. Defaults to metadata.name",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"opensearch_cluster": schema.SingleNestedAttribute{
						Description:         "LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.",
						MarkdownDescription: "LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
								MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"priority": schema.Int64Attribute{
						Description:         "Priority to determine index template precedence when a new data stream or index is created. The index template with the highest priority is chosen",
						MarkdownDescription: "Priority to determine index template precedence when a new data stream or index is created. The index template with the highest priority is chosen",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"template": schema.SingleNestedAttribute{
						Description:         "The template that should be applied",
						MarkdownDescription: "The template that should be applied",
						Attributes: map[string]schema.Attribute{
							"aliases": schema.SingleNestedAttribute{
								Description:         "Aliases to add",
								MarkdownDescription: "Aliases to add",
								Attributes: map[string]schema.Attribute{
									"alias": schema.StringAttribute{
										Description:         "The name of the alias.",
										MarkdownDescription: "The name of the alias.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"filter": schema.MapAttribute{
										Description:         "Query used to limit documents the alias can access.",
										MarkdownDescription: "Query used to limit documents the alias can access.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"index": schema.StringAttribute{
										Description:         "The name of the index that the alias points to.",
										MarkdownDescription: "The name of the index that the alias points to.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"is_write_index": schema.BoolAttribute{
										Description:         "If true, the index is the write index for the alias",
										MarkdownDescription: "If true, the index is the write index for the alias",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"routing": schema.StringAttribute{
										Description:         "Value used to route indexing and search operations to a specific shard.",
										MarkdownDescription: "Value used to route indexing and search operations to a specific shard.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"mappings": schema.MapAttribute{
								Description:         "Mapping for fields in the index",
								MarkdownDescription: "Mapping for fields in the index",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"settings": schema.MapAttribute{
								Description:         "Configuration options for the index",
								MarkdownDescription: "Configuration options for the index",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"version": schema.Int64Attribute{
						Description:         "Version number used to manage the component template externally",
						MarkdownDescription: "Version number used to manage the component template externally",
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

func (r *OpensearchOpsterIoOpensearchIndexTemplateV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_opensearch_opster_io_opensearch_index_template_v1_manifest")

	var model OpensearchOpsterIoOpensearchIndexTemplateV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("opensearch.opster.io/v1")
	model.Kind = pointer.String("OpensearchIndexTemplate")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
