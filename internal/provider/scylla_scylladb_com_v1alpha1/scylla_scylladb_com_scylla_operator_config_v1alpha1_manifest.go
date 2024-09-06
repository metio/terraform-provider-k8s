/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package scylla_scylladb_com_v1alpha1

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
	_ datasource.DataSource = &ScyllaScylladbComScyllaOperatorConfigV1Alpha1Manifest{}
)

func NewScyllaScylladbComScyllaOperatorConfigV1Alpha1Manifest() datasource.DataSource {
	return &ScyllaScylladbComScyllaOperatorConfigV1Alpha1Manifest{}
}

type ScyllaScylladbComScyllaOperatorConfigV1Alpha1Manifest struct{}

type ScyllaScylladbComScyllaOperatorConfigV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		ScyllaUtilsImage                     *string `tfsdk:"scylla_utils_image" json:"scyllaUtilsImage,omitempty"`
		UnsupportedBashToolsImageOverride    *string `tfsdk:"unsupported_bash_tools_image_override" json:"unsupportedBashToolsImageOverride,omitempty"`
		UnsupportedGrafanaImageOverride      *string `tfsdk:"unsupported_grafana_image_override" json:"unsupportedGrafanaImageOverride,omitempty"`
		UnsupportedPrometheusVersionOverride *string `tfsdk:"unsupported_prometheus_version_override" json:"unsupportedPrometheusVersionOverride,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ScyllaScylladbComScyllaOperatorConfigV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_scylla_scylladb_com_scylla_operator_config_v1alpha1_manifest"
}

func (r *ScyllaScylladbComScyllaOperatorConfigV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ScyllaOperatorConfig describes the Scylla Operator configuration.",
		MarkdownDescription: "ScyllaOperatorConfig describes the Scylla Operator configuration.",
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
				Description:         "spec defines the desired state of the operator.",
				MarkdownDescription: "spec defines the desired state of the operator.",
				Attributes: map[string]schema.Attribute{
					"scylla_utils_image": schema.StringAttribute{
						Description:         "scyllaUtilsImage is a ScyllaDB image used for running ScyllaDB utilities.",
						MarkdownDescription: "scyllaUtilsImage is a ScyllaDB image used for running ScyllaDB utilities.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"unsupported_bash_tools_image_override": schema.StringAttribute{
						Description:         "unsupportedBashToolsImageOverride allows to adjust a generic Bash image with extra tools used by the operator for auxiliary purposes. Setting this field renders your cluster unsupported. Use at your own risk.",
						MarkdownDescription: "unsupportedBashToolsImageOverride allows to adjust a generic Bash image with extra tools used by the operator for auxiliary purposes. Setting this field renders your cluster unsupported. Use at your own risk.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"unsupported_grafana_image_override": schema.StringAttribute{
						Description:         "unsupportedGrafanaImageOverride allows to adjust Grafana image used by the operator for testing, dev or emergencies. Setting this field renders your cluster unsupported. Use at your own risk.",
						MarkdownDescription: "unsupportedGrafanaImageOverride allows to adjust Grafana image used by the operator for testing, dev or emergencies. Setting this field renders your cluster unsupported. Use at your own risk.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"unsupported_prometheus_version_override": schema.StringAttribute{
						Description:         "unsupportedPrometheusVersionOverride allows to adjust Prometheus version used by the operator for testing, dev or emergencies. Setting this field renders your cluster unsupported. Use at your own risk.",
						MarkdownDescription: "unsupportedPrometheusVersionOverride allows to adjust Prometheus version used by the operator for testing, dev or emergencies. Setting this field renders your cluster unsupported. Use at your own risk.",
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

func (r *ScyllaScylladbComScyllaOperatorConfigV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_scylla_scylladb_com_scylla_operator_config_v1alpha1_manifest")

	var model ScyllaScylladbComScyllaOperatorConfigV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("scylla.scylladb.com/v1alpha1")
	model.Kind = pointer.String("ScyllaOperatorConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
