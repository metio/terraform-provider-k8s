/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package atlasmap_io_v1alpha1

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
	_ datasource.DataSource = &AtlasmapIoAtlasMapV1Alpha1Manifest{}
)

func NewAtlasmapIoAtlasMapV1Alpha1Manifest() datasource.DataSource {
	return &AtlasmapIoAtlasMapV1Alpha1Manifest{}
}

type AtlasmapIoAtlasMapV1Alpha1Manifest struct{}

type AtlasmapIoAtlasMapV1Alpha1ManifestData struct {
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
		LimitCPU      *string `tfsdk:"limit_cpu" json:"limitCPU,omitempty"`
		LimitMemory   *string `tfsdk:"limit_memory" json:"limitMemory,omitempty"`
		Replicas      *int64  `tfsdk:"replicas" json:"replicas,omitempty"`
		RequestCPU    *string `tfsdk:"request_cpu" json:"requestCPU,omitempty"`
		RequestMemory *string `tfsdk:"request_memory" json:"requestMemory,omitempty"`
		RouteHostName *string `tfsdk:"route_host_name" json:"routeHostName,omitempty"`
		Version       *string `tfsdk:"version" json:"version,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AtlasmapIoAtlasMapV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_atlasmap_io_atlas_map_v1alpha1_manifest"
}

func (r *AtlasmapIoAtlasMapV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "AtlasMap is the Schema for the atlasmaps API",
		MarkdownDescription: "AtlasMap is the Schema for the atlasmaps API",
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
				Description:         "AtlasMapSpec defines the desired state of AtlasMap",
				MarkdownDescription: "AtlasMapSpec defines the desired state of AtlasMap",
				Attributes: map[string]schema.Attribute{
					"limit_cpu": schema.StringAttribute{
						Description:         "The amount of CPU to limit",
						MarkdownDescription: "The amount of CPU to limit",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`[0-9]+m?$`), ""),
						},
					},

					"limit_memory": schema.StringAttribute{
						Description:         "The amount of memory to request",
						MarkdownDescription: "The amount of memory to request",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`[0-9]+([kKmMgGtTpPeE]i?)?$`), ""),
						},
					},

					"replicas": schema.Int64Attribute{
						Description:         "Replicas determines the desired number of running AtlasMap pods",
						MarkdownDescription: "Replicas determines the desired number of running AtlasMap pods",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"request_cpu": schema.StringAttribute{
						Description:         "The amount of CPU to request",
						MarkdownDescription: "The amount of CPU to request",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`[0-9]+m?$`), ""),
						},
					},

					"request_memory": schema.StringAttribute{
						Description:         "The amount of memory to request",
						MarkdownDescription: "The amount of memory to request",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`[0-9]+([kKmMgGtTpPeE]i?)?$`), ""),
						},
					},

					"route_host_name": schema.StringAttribute{
						Description:         "RouteHostName sets the host name to use on the Ingress or OpenShift Route",
						MarkdownDescription: "RouteHostName sets the host name to use on the Ingress or OpenShift Route",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"version": schema.StringAttribute{
						Description:         "Version sets the version of the container image used for AtlasMap",
						MarkdownDescription: "Version sets the version of the container image used for AtlasMap",
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

func (r *AtlasmapIoAtlasMapV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_atlasmap_io_atlas_map_v1alpha1_manifest")

	var model AtlasmapIoAtlasMapV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("atlasmap.io/v1alpha1")
	model.Kind = pointer.String("AtlasMap")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
