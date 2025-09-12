/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package v2_edp_epam_com_v1

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
	_ datasource.DataSource = &V2EdpEpamComStageV1Manifest{}
)

func NewV2EdpEpamComStageV1Manifest() datasource.DataSource {
	return &V2EdpEpamComStageV1Manifest{}
}

type V2EdpEpamComStageV1Manifest struct{}

type V2EdpEpamComStageV1ManifestData struct {
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
		CdPipeline    *string `tfsdk:"cd_pipeline" json:"cdPipeline,omitempty"`
		CleanTemplate *string `tfsdk:"clean_template" json:"cleanTemplate,omitempty"`
		ClusterName   *string `tfsdk:"cluster_name" json:"clusterName,omitempty"`
		Description   *string `tfsdk:"description" json:"description,omitempty"`
		Name          *string `tfsdk:"name" json:"name,omitempty"`
		Namespace     *string `tfsdk:"namespace" json:"namespace,omitempty"`
		Order         *int64  `tfsdk:"order" json:"order,omitempty"`
		QualityGates  *[]struct {
			AutotestName    *string `tfsdk:"autotest_name" json:"autotestName,omitempty"`
			BranchName      *string `tfsdk:"branch_name" json:"branchName,omitempty"`
			QualityGateType *string `tfsdk:"quality_gate_type" json:"qualityGateType,omitempty"`
			StepName        *string `tfsdk:"step_name" json:"stepName,omitempty"`
		} `tfsdk:"quality_gates" json:"qualityGates,omitempty"`
		Source *struct {
			Library *struct {
				Branch *string `tfsdk:"branch" json:"branch,omitempty"`
				Name   *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"library" json:"library,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"source" json:"source,omitempty"`
		TriggerTemplate *string `tfsdk:"trigger_template" json:"triggerTemplate,omitempty"`
		TriggerType     *string `tfsdk:"trigger_type" json:"triggerType,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *V2EdpEpamComStageV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_v2_edp_epam_com_stage_v1_manifest"
}

func (r *V2EdpEpamComStageV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Stage is the Schema for the stages API.",
		MarkdownDescription: "Stage is the Schema for the stages API.",
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
				Description:         "StageSpec defines the desired state of Stage. NOTE: for deleting the stage use stages order - delete only the latest stage.",
				MarkdownDescription: "StageSpec defines the desired state of Stage. NOTE: for deleting the stage use stages order - delete only the latest stage.",
				Attributes: map[string]schema.Attribute{
					"cd_pipeline": schema.StringAttribute{
						Description:         "Name of CD pipeline which this Stage will be linked to.",
						MarkdownDescription: "Name of CD pipeline which this Stage will be linked to.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(2),
						},
					},

					"clean_template": schema.StringAttribute{
						Description:         "CleanTemplate specifies the name of Tekton TriggerTemplate used for cleanup environment pipeline.",
						MarkdownDescription: "CleanTemplate specifies the name of Tekton TriggerTemplate used for cleanup environment pipeline.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"cluster_name": schema.StringAttribute{
						Description:         "Specifies a name of cluster where the application will be deployed. Default value is 'in-cluster' which means that application will be deployed in the same cluster where CD Pipeline is running.",
						MarkdownDescription: "Specifies a name of cluster where the application will be deployed. Default value is 'in-cluster' which means that application will be deployed in the same cluster where CD Pipeline is running.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"description": schema.StringAttribute{
						Description:         "A description of a stage.",
						MarkdownDescription: "A description of a stage.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(0),
						},
					},

					"name": schema.StringAttribute{
						Description:         "Name of a stage.",
						MarkdownDescription: "Name of a stage.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(2),
						},
					},

					"namespace": schema.StringAttribute{
						Description:         "Namespace where the application will be deployed.",
						MarkdownDescription: "Namespace where the application will be deployed.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"order": schema.Int64Attribute{
						Description:         "The order to lay out Stages. The order should start from 0, and the next stages should use +1 for the order.",
						MarkdownDescription: "The order to lay out Stages. The order should start from 0, and the next stages should use +1 for the order.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"quality_gates": schema.ListNestedAttribute{
						Description:         "A list of quality gates to be processed",
						MarkdownDescription: "A list of quality gates to be processed",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"autotest_name": schema.StringAttribute{
									Description:         "A name of autotests to run with quality gate",
									MarkdownDescription: "A name of autotests to run with quality gate",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"branch_name": schema.StringAttribute{
									Description:         "A branch name to use from autotests repository",
									MarkdownDescription: "A branch name to use from autotests repository",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"quality_gate_type": schema.StringAttribute{
									Description:         "A type of quality gate, e.g. 'Manual', 'Autotests'",
									MarkdownDescription: "A type of quality gate, e.g. 'Manual', 'Autotests'",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"step_name": schema.StringAttribute{
									Description:         "Specifies a name of particular",
									MarkdownDescription: "Specifies a name of particular",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(2),
									},
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"source": schema.SingleNestedAttribute{
						Description:         "Specifies a source of a pipeline library which will run release",
						MarkdownDescription: "Specifies a source of a pipeline library which will run release",
						Attributes: map[string]schema.Attribute{
							"library": schema.SingleNestedAttribute{
								Description:         "A reference to a non default source library",
								MarkdownDescription: "A reference to a non default source library",
								Attributes: map[string]schema.Attribute{
									"branch": schema.StringAttribute{
										Description:         "Branch which should be used for a library",
										MarkdownDescription: "Branch which should be used for a library",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "A name of a library",
										MarkdownDescription: "A name of a library",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"type": schema.StringAttribute{
								Description:         "Type of pipeline library, e.g. default, library",
								MarkdownDescription: "Type of pipeline library, e.g. default, library",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"trigger_template": schema.StringAttribute{
						Description:         "Specifies a name of Tekton TriggerTemplate used as a blueprint for deployment pipeline. Default value is 'deploy' which means that default TriggerTemplate will be used. The default TriggerTemplate is delivered using edp-tekton helm chart.",
						MarkdownDescription: "Specifies a name of Tekton TriggerTemplate used as a blueprint for deployment pipeline. Default value is 'deploy' which means that default TriggerTemplate will be used. The default TriggerTemplate is delivered using edp-tekton helm chart.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"trigger_type": schema.StringAttribute{
						Description:         "Stage deployment trigger type.",
						MarkdownDescription: "Stage deployment trigger type.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Auto", "Manual", "Auto-stable"),
						},
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *V2EdpEpamComStageV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_v2_edp_epam_com_stage_v1_manifest")

	var model V2EdpEpamComStageV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("v2.edp.epam.com/v1")
	model.Kind = pointer.String("Stage")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
