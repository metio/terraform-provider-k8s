/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package trace_kubeblocks_io_v1

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
	_ datasource.DataSource = &TraceKubeblocksIoReconciliationTraceV1Manifest{}
)

func NewTraceKubeblocksIoReconciliationTraceV1Manifest() datasource.DataSource {
	return &TraceKubeblocksIoReconciliationTraceV1Manifest{}
}

type TraceKubeblocksIoReconciliationTraceV1Manifest struct{}

type TraceKubeblocksIoReconciliationTraceV1ManifestData struct {
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
		DryRun *struct {
			DesiredSpec *string `tfsdk:"desired_spec" json:"desiredSpec,omitempty"`
		} `tfsdk:"dry_run" json:"dryRun,omitempty"`
		Locale                    *string `tfsdk:"locale" json:"locale,omitempty"`
		StateEvaluationExpression *struct {
			CelExpression *struct {
				Expression *string `tfsdk:"expression" json:"expression,omitempty"`
			} `tfsdk:"cel_expression" json:"celExpression,omitempty"`
		} `tfsdk:"state_evaluation_expression" json:"stateEvaluationExpression,omitempty"`
		TargetObject *struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"target_object" json:"targetObject,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *TraceKubeblocksIoReconciliationTraceV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_trace_kubeblocks_io_reconciliation_trace_v1_manifest"
}

func (r *TraceKubeblocksIoReconciliationTraceV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ReconciliationTrace is the Schema for the reconciliationtraces API",
		MarkdownDescription: "ReconciliationTrace is the Schema for the reconciliationtraces API",
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
				Description:         "ReconciliationTraceSpec defines the desired state of ReconciliationTrace",
				MarkdownDescription: "ReconciliationTraceSpec defines the desired state of ReconciliationTrace",
				Attributes: map[string]schema.Attribute{
					"dry_run": schema.SingleNestedAttribute{
						Description:         "DryRun tells the Controller to simulate the reconciliation process with a new desired spec of the TargetObject. And a reconciliation plan will be generated and described in the ReconciliationTraceStatus. The plan generation process will not impact the state of the TargetObject.",
						MarkdownDescription: "DryRun tells the Controller to simulate the reconciliation process with a new desired spec of the TargetObject. And a reconciliation plan will be generated and described in the ReconciliationTraceStatus. The plan generation process will not impact the state of the TargetObject.",
						Attributes: map[string]schema.Attribute{
							"desired_spec": schema.StringAttribute{
								Description:         "DesiredSpec specifies the desired spec of the TargetObject. The desired spec will be merged into the current spec by a strategic merge patch way to build the final spec, and the reconciliation plan will be calculated by comparing the current spec to the final spec. DesiredSpec should be a valid YAML string.",
								MarkdownDescription: "DesiredSpec specifies the desired spec of the TargetObject. The desired spec will be merged into the current spec by a strategic merge patch way to build the final spec, and the reconciliation plan will be calculated by comparing the current spec to the final spec. DesiredSpec should be a valid YAML string.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"locale": schema.StringAttribute{
						Description:         "Locale specifies the locale to use when localizing the reconciliation trace.",
						MarkdownDescription: "Locale specifies the locale to use when localizing the reconciliation trace.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"state_evaluation_expression": schema.SingleNestedAttribute{
						Description:         "StateEvaluationExpression specifies the state evaluation expression used during reconciliation progress observation. The whole reconciliation process from the creation of the TargetObject to the deletion of it is separated into several reconciliation cycles. The StateEvaluationExpression is applied to the TargetObject, and an evaluation result of true indicates the end of a reconciliation cycle. StateEvaluationExpression overrides the builtin default value.",
						MarkdownDescription: "StateEvaluationExpression specifies the state evaluation expression used during reconciliation progress observation. The whole reconciliation process from the creation of the TargetObject to the deletion of it is separated into several reconciliation cycles. The StateEvaluationExpression is applied to the TargetObject, and an evaluation result of true indicates the end of a reconciliation cycle. StateEvaluationExpression overrides the builtin default value.",
						Attributes: map[string]schema.Attribute{
							"cel_expression": schema.SingleNestedAttribute{
								Description:         "CELExpression specifies to use CEL to evaluation the object state. The root object used in the expression is the primary object.",
								MarkdownDescription: "CELExpression specifies to use CEL to evaluation the object state. The root object used in the expression is the primary object.",
								Attributes: map[string]schema.Attribute{
									"expression": schema.StringAttribute{
										Description:         "Expression specifies the CEL expression.",
										MarkdownDescription: "Expression specifies the CEL expression.",
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

					"target_object": schema.SingleNestedAttribute{
						Description:         "TargetObject specifies the target Cluster object. Default is the Cluster object with same namespace and name as this ReconciliationTrace object.",
						MarkdownDescription: "TargetObject specifies the target Cluster object. Default is the Cluster object with same namespace and name as this ReconciliationTrace object.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of the referent. Default is same as the ReconciliationTrace object.",
								MarkdownDescription: "Name of the referent. Default is same as the ReconciliationTrace object.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace of the referent. Default is same as the ReconciliationTrace object.",
								MarkdownDescription: "Namespace of the referent. Default is same as the ReconciliationTrace object.",
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

func (r *TraceKubeblocksIoReconciliationTraceV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_trace_kubeblocks_io_reconciliation_trace_v1_manifest")

	var model TraceKubeblocksIoReconciliationTraceV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("trace.kubeblocks.io/v1")
	model.Kind = pointer.String("ReconciliationTrace")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
