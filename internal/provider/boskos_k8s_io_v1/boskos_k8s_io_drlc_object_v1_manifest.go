/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package boskos_k8s_io_v1

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
	_ datasource.DataSource = &BoskosK8SIoDrlcobjectV1Manifest{}
)

func NewBoskosK8SIoDrlcobjectV1Manifest() datasource.DataSource {
	return &BoskosK8SIoDrlcobjectV1Manifest{}
}

type BoskosK8SIoDrlcobjectV1Manifest struct{}

type BoskosK8SIoDrlcobjectV1ManifestData struct {
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
		Config *struct {
			Content *string `tfsdk:"content" json:"content,omitempty"`
			Type    *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"config" json:"config,omitempty"`
		Lifespan  *int64             `tfsdk:"lifespan" json:"lifespan,omitempty"`
		Max_count *int64             `tfsdk:"max_count" json:"max-count,omitempty"`
		Min_count *int64             `tfsdk:"min_count" json:"min-count,omitempty"`
		Needs     *map[string]string `tfsdk:"needs" json:"needs,omitempty"`
		State     *string            `tfsdk:"state" json:"state,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *BoskosK8SIoDrlcobjectV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_boskos_k8s_io_drlc_object_v1_manifest"
}

func (r *BoskosK8SIoDrlcobjectV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Defines the lifecycle of a dynamic resource. All Resource of a given type will be constructed using the same configuration",
		MarkdownDescription: "Defines the lifecycle of a dynamic resource. All Resource of a given type will be constructed using the same configuration",
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
					"config": schema.SingleNestedAttribute{
						Description:         "Config information about how to create the object",
						MarkdownDescription: "Config information about how to create the object",
						Attributes: map[string]schema.Attribute{
							"content": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "The dynamic resource type",
								MarkdownDescription: "The dynamic resource type",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"lifespan": schema.Int64Attribute{
						Description:         "Lifespan of a resource, time after which the resource should be reset",
						MarkdownDescription: "Lifespan of a resource, time after which the resource should be reset",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"max_count": schema.Int64Attribute{
						Description:         "Maxiumum number of resources expected. This maximum may be temporarily exceeded while resources are in the process of being deleted, though this is only expected when MaxCount is lowered.",
						MarkdownDescription: "Maxiumum number of resources expected. This maximum may be temporarily exceeded while resources are in the process of being deleted, though this is only expected when MaxCount is lowered.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"min_count": schema.Int64Attribute{
						Description:         "Minimum number of resources to be used as a buffer. Resources in the process of being deleted and cleaned up are included in this count.",
						MarkdownDescription: "Minimum number of resources to be used as a buffer. Resources in the process of being deleted and cleaned up are included in this count.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"needs": schema.MapAttribute{
						Description:         "Define the resource needs to create the object",
						MarkdownDescription: "Define the resource needs to create the object",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"state": schema.StringAttribute{
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
	}
}

func (r *BoskosK8SIoDrlcobjectV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_boskos_k8s_io_drlc_object_v1_manifest")

	var model BoskosK8SIoDrlcobjectV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("boskos.k8s.io/v1")
	model.Kind = pointer.String("DRLCObject")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
