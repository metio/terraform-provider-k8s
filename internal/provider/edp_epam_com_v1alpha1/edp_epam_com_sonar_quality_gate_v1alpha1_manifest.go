/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package edp_epam_com_v1alpha1

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
	_ datasource.DataSource = &EdpEpamComSonarQualityGateV1Alpha1Manifest{}
)

func NewEdpEpamComSonarQualityGateV1Alpha1Manifest() datasource.DataSource {
	return &EdpEpamComSonarQualityGateV1Alpha1Manifest{}
}

type EdpEpamComSonarQualityGateV1Alpha1Manifest struct{}

type EdpEpamComSonarQualityGateV1Alpha1ManifestData struct {
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
		Conditions *struct {
			Error *string `tfsdk:"error" json:"error,omitempty"`
			Op    *string `tfsdk:"op" json:"op,omitempty"`
		} `tfsdk:"conditions" json:"conditions,omitempty"`
		Default  *bool   `tfsdk:"default" json:"default,omitempty"`
		Name     *string `tfsdk:"name" json:"name,omitempty"`
		SonarRef *struct {
			Kind *string `tfsdk:"kind" json:"kind,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"sonar_ref" json:"sonarRef,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *EdpEpamComSonarQualityGateV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_edp_epam_com_sonar_quality_gate_v1alpha1_manifest"
}

func (r *EdpEpamComSonarQualityGateV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "SonarQualityGate is the Schema for the sonarqualitygates API",
		MarkdownDescription: "SonarQualityGate is the Schema for the sonarqualitygates API",
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
				Description:         "SonarQualityGateSpec defines the desired state of SonarQualityGate",
				MarkdownDescription: "SonarQualityGateSpec defines the desired state of SonarQualityGate",
				Attributes: map[string]schema.Attribute{
					"conditions": schema.SingleNestedAttribute{
						Description:         "Conditions is a list of conditions for quality gate. Key is a metric name, value is a condition.",
						MarkdownDescription: "Conditions is a list of conditions for quality gate. Key is a metric name, value is a condition.",
						Attributes: map[string]schema.Attribute{
							"error": schema.StringAttribute{
								Description:         "Error is condition error threshold.",
								MarkdownDescription: "Error is condition error threshold.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtMost(64),
								},
							},

							"op": schema.StringAttribute{
								Description:         "Op is condition operator. LT = is lower than GT = is greater than",
								MarkdownDescription: "Op is condition operator. LT = is lower than GT = is greater than",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("LT", "GT"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"default": schema.BoolAttribute{
						Description:         "Default is a flag to set quality gate as default. Only one quality gate can be default. If several quality gates have default flag, the random one will be chosen. Default quality gate can't be deleted. You need to set another quality gate as default before.",
						MarkdownDescription: "Default is a flag to set quality gate as default. Only one quality gate can be default. If several quality gates have default flag, the random one will be chosen. Default quality gate can't be deleted. You need to set another quality gate as default before.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"name": schema.StringAttribute{
						Description:         "Name is a name of quality gate. Name should be unique across all quality gates. Don't change this field after creation otherwise quality gate will be recreated.",
						MarkdownDescription: "Name is a name of quality gate. Name should be unique across all quality gates. Don't change this field after creation otherwise quality gate will be recreated.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(100),
						},
					},

					"sonar_ref": schema.SingleNestedAttribute{
						Description:         "SonarRef is a reference to Sonar custom resource.",
						MarkdownDescription: "SonarRef is a reference to Sonar custom resource.",
						Attributes: map[string]schema.Attribute{
							"kind": schema.StringAttribute{
								Description:         "Kind specifies the kind of the Sonar resource.",
								MarkdownDescription: "Kind specifies the kind of the Sonar resource.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name specifies the name of the Sonar resource.",
								MarkdownDescription: "Name specifies the name of the Sonar resource.",
								Required:            true,
								Optional:            false,
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

func (r *EdpEpamComSonarQualityGateV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_edp_epam_com_sonar_quality_gate_v1alpha1_manifest")

	var model EdpEpamComSonarQualityGateV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("edp.epam.com/v1alpha1")
	model.Kind = pointer.String("SonarQualityGate")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
