/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package sfn_services_k8s_aws_v1alpha1

import (
	"context"
	"fmt"
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
	_ datasource.DataSource = &SfnServicesK8SAwsStateMachineV1Alpha1Manifest{}
)

func NewSfnServicesK8SAwsStateMachineV1Alpha1Manifest() datasource.DataSource {
	return &SfnServicesK8SAwsStateMachineV1Alpha1Manifest{}
}

type SfnServicesK8SAwsStateMachineV1Alpha1Manifest struct{}

type SfnServicesK8SAwsStateMachineV1Alpha1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
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
		Definition           *string `tfsdk:"definition" json:"definition,omitempty"`
		LoggingConfiguration *struct {
			Destinations *[]struct {
				CloudWatchLogsLogGroup *struct {
					LogGroupARN *string `tfsdk:"log_group_arn" json:"logGroupARN,omitempty"`
				} `tfsdk:"cloud_watch_logs_log_group" json:"cloudWatchLogsLogGroup,omitempty"`
			} `tfsdk:"destinations" json:"destinations,omitempty"`
			IncludeExecutionData *bool   `tfsdk:"include_execution_data" json:"includeExecutionData,omitempty"`
			Level                *string `tfsdk:"level" json:"level,omitempty"`
		} `tfsdk:"logging_configuration" json:"loggingConfiguration,omitempty"`
		Name    *string `tfsdk:"name" json:"name,omitempty"`
		RoleARN *string `tfsdk:"role_arn" json:"roleARN,omitempty"`
		Tags    *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
		TracingConfiguration *struct {
			Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
		} `tfsdk:"tracing_configuration" json:"tracingConfiguration,omitempty"`
		Type_ *string `tfsdk:"type_" json:"type_,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SfnServicesK8SAwsStateMachineV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_sfn_services_k8s_aws_state_machine_v1alpha1_manifest"
}

func (r *SfnServicesK8SAwsStateMachineV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "StateMachine is the Schema for the StateMachines API",
		MarkdownDescription: "StateMachine is the Schema for the StateMachines API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

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
				Description:         "StateMachineSpec defines the desired state of StateMachine.",
				MarkdownDescription: "StateMachineSpec defines the desired state of StateMachine.",
				Attributes: map[string]schema.Attribute{
					"definition": schema.StringAttribute{
						Description:         "The Amazon States Language definition of the state machine. See Amazon StatesLanguage (https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html).",
						MarkdownDescription: "The Amazon States Language definition of the state machine. See Amazon StatesLanguage (https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html).",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"logging_configuration": schema.SingleNestedAttribute{
						Description:         "Defines what execution history events are logged and where they are logged.By default, the level is set to OFF. For more information see Log Levels(https://docs.aws.amazon.com/step-functions/latest/dg/cloudwatch-log-level.html)in the AWS Step Functions User Guide.",
						MarkdownDescription: "Defines what execution history events are logged and where they are logged.By default, the level is set to OFF. For more information see Log Levels(https://docs.aws.amazon.com/step-functions/latest/dg/cloudwatch-log-level.html)in the AWS Step Functions User Guide.",
						Attributes: map[string]schema.Attribute{
							"destinations": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"cloud_watch_logs_log_group": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"log_group_arn": schema.StringAttribute{
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"include_execution_data": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"level": schema.StringAttribute{
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

					"name": schema.StringAttribute{
						Description:         "The name of the state machine.A name must not contain:   * white space   * brackets < > { } [ ]   * wildcard characters ? *   * special characters ' # %  ^ | ~ ' $ & , ; : /   * control characters (U+0000-001F, U+007F-009F)To enable logging with CloudWatch Logs, the name should only contain 0-9,A-Z, a-z, - and _.",
						MarkdownDescription: "The name of the state machine.A name must not contain:   * white space   * brackets < > { } [ ]   * wildcard characters ? *   * special characters ' # %  ^ | ~ ' $ & , ; : /   * control characters (U+0000-001F, U+007F-009F)To enable logging with CloudWatch Logs, the name should only contain 0-9,A-Z, a-z, - and _.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"role_arn": schema.StringAttribute{
						Description:         "The Amazon Resource Name (ARN) of the IAM role to use for this state machine.",
						MarkdownDescription: "The Amazon Resource Name (ARN) of the IAM role to use for this state machine.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"tags": schema.ListNestedAttribute{
						Description:         "Tags to be added when creating a state machine.An array of key-value pairs. For more information, see Using Cost AllocationTags (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html)in the AWS Billing and Cost Management User Guide, and Controlling AccessUsing IAM Tags (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_iam-tags.html).Tags may only contain Unicode letters, digits, white space, or these symbols:_ . : / = + - @.",
						MarkdownDescription: "Tags to be added when creating a state machine.An array of key-value pairs. For more information, see Using Cost AllocationTags (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html)in the AWS Billing and Cost Management User Guide, and Controlling AccessUsing IAM Tags (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_iam-tags.html).Tags may only contain Unicode letters, digits, white space, or these symbols:_ . : / = + - @.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"key": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"value": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"tracing_configuration": schema.SingleNestedAttribute{
						Description:         "Selects whether AWS X-Ray tracing is enabled.",
						MarkdownDescription: "Selects whether AWS X-Ray tracing is enabled.",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
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

					"type_": schema.StringAttribute{
						Description:         "Determines whether a Standard or Express state machine is created. The defaultis STANDARD. You cannot update the type of a state machine once it has beencreated.",
						MarkdownDescription: "Determines whether a Standard or Express state machine is created. The defaultis STANDARD. You cannot update the type of a state machine once it has beencreated.",
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

func (r *SfnServicesK8SAwsStateMachineV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_sfn_services_k8s_aws_state_machine_v1alpha1_manifest")

	var model SfnServicesK8SAwsStateMachineV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("sfn.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("StateMachine")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
