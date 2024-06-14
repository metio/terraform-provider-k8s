/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package flagger_app_v1beta1

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
	_ datasource.DataSource = &FlaggerAppMetricTemplateV1Beta1Manifest{}
)

func NewFlaggerAppMetricTemplateV1Beta1Manifest() datasource.DataSource {
	return &FlaggerAppMetricTemplateV1Beta1Manifest{}
}

type FlaggerAppMetricTemplateV1Beta1Manifest struct{}

type FlaggerAppMetricTemplateV1Beta1ManifestData struct {
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
		Provider *struct {
			Address            *string `tfsdk:"address" json:"address,omitempty"`
			InsecureSkipVerify *bool   `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
			Region             *string `tfsdk:"region" json:"region,omitempty"`
			SecretRef          *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"provider" json:"provider,omitempty"`
		Query *string `tfsdk:"query" json:"query,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *FlaggerAppMetricTemplateV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_flagger_app_metric_template_v1beta1_manifest"
}

func (r *FlaggerAppMetricTemplateV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "MetricTemplate is the Schema for the MetricTemplates API.",
		MarkdownDescription: "MetricTemplate is the Schema for the MetricTemplates API.",
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
				Description:         "MetricTemplateSpec defines the desired state of a MetricTemplate.",
				MarkdownDescription: "MetricTemplateSpec defines the desired state of a MetricTemplate.",
				Attributes: map[string]schema.Attribute{
					"provider": schema.SingleNestedAttribute{
						Description:         "Provider of this metric template",
						MarkdownDescription: "Provider of this metric template",
						Attributes: map[string]schema.Attribute{
							"address": schema.StringAttribute{
								Description:         "API address of this provider",
								MarkdownDescription: "API address of this provider",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"insecure_skip_verify": schema.BoolAttribute{
								Description:         "Disable SSL certificate validation for the provider address",
								MarkdownDescription: "Disable SSL certificate validation for the provider address",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"region": schema.StringAttribute{
								Description:         "Region of the provider",
								MarkdownDescription: "Region of the provider",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secret_ref": schema.SingleNestedAttribute{
								Description:         "Kubernetes secret reference containing the provider credentials",
								MarkdownDescription: "Kubernetes secret reference containing the provider credentials",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of the Kubernetes secret",
										MarkdownDescription: "Name of the Kubernetes secret",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"type": schema.StringAttribute{
								Description:         "Type of this provider",
								MarkdownDescription: "Type of this provider",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("prometheus", "influxdb", "datadog", "stackdriver", "cloudwatch", "newrelic", "graphite", "dynatrace", "keptn"),
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"query": schema.StringAttribute{
						Description:         "Query of this metric template",
						MarkdownDescription: "Query of this metric template",
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

func (r *FlaggerAppMetricTemplateV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_flagger_app_metric_template_v1beta1_manifest")

	var model FlaggerAppMetricTemplateV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("flagger.app/v1beta1")
	model.Kind = pointer.String("MetricTemplate")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
