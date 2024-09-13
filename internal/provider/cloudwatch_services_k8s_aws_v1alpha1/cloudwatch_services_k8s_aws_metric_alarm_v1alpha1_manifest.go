/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package cloudwatch_services_k8s_aws_v1alpha1

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
	_ datasource.DataSource = &CloudwatchServicesK8SAwsMetricAlarmV1Alpha1Manifest{}
)

func NewCloudwatchServicesK8SAwsMetricAlarmV1Alpha1Manifest() datasource.DataSource {
	return &CloudwatchServicesK8SAwsMetricAlarmV1Alpha1Manifest{}
}

type CloudwatchServicesK8SAwsMetricAlarmV1Alpha1Manifest struct{}

type CloudwatchServicesK8SAwsMetricAlarmV1Alpha1ManifestData struct {
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
		ActionsEnabled     *bool     `tfsdk:"actions_enabled" json:"actionsEnabled,omitempty"`
		AlarmActions       *[]string `tfsdk:"alarm_actions" json:"alarmActions,omitempty"`
		AlarmDescription   *string   `tfsdk:"alarm_description" json:"alarmDescription,omitempty"`
		ComparisonOperator *string   `tfsdk:"comparison_operator" json:"comparisonOperator,omitempty"`
		DatapointsToAlarm  *int64    `tfsdk:"datapoints_to_alarm" json:"datapointsToAlarm,omitempty"`
		Dimensions         *[]struct {
			Name  *string `tfsdk:"name" json:"name,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"dimensions" json:"dimensions,omitempty"`
		EvaluateLowSampleCountPercentile *string   `tfsdk:"evaluate_low_sample_count_percentile" json:"evaluateLowSampleCountPercentile,omitempty"`
		EvaluationPeriods                *int64    `tfsdk:"evaluation_periods" json:"evaluationPeriods,omitempty"`
		ExtendedStatistic                *string   `tfsdk:"extended_statistic" json:"extendedStatistic,omitempty"`
		InsufficientDataActions          *[]string `tfsdk:"insufficient_data_actions" json:"insufficientDataActions,omitempty"`
		MetricName                       *string   `tfsdk:"metric_name" json:"metricName,omitempty"`
		Metrics                          *[]struct {
			AccountID  *string `tfsdk:"account_id" json:"accountID,omitempty"`
			Expression *string `tfsdk:"expression" json:"expression,omitempty"`
			Id         *string `tfsdk:"id" json:"id,omitempty"`
			Label      *string `tfsdk:"label" json:"label,omitempty"`
			MetricStat *struct {
				Metric *struct {
					Dimensions *[]struct {
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"dimensions" json:"dimensions,omitempty"`
					MetricName *string `tfsdk:"metric_name" json:"metricName,omitempty"`
					Namespace  *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"metric" json:"metric,omitempty"`
				Period *int64  `tfsdk:"period" json:"period,omitempty"`
				Stat   *string `tfsdk:"stat" json:"stat,omitempty"`
				Unit   *string `tfsdk:"unit" json:"unit,omitempty"`
			} `tfsdk:"metric_stat" json:"metricStat,omitempty"`
			Period     *int64 `tfsdk:"period" json:"period,omitempty"`
			ReturnData *bool  `tfsdk:"return_data" json:"returnData,omitempty"`
		} `tfsdk:"metrics" json:"metrics,omitempty"`
		Name      *string   `tfsdk:"name" json:"name,omitempty"`
		Namespace *string   `tfsdk:"namespace" json:"namespace,omitempty"`
		OKActions *[]string `tfsdk:"o_k_actions" json:"oKActions,omitempty"`
		Period    *int64    `tfsdk:"period" json:"period,omitempty"`
		Statistic *string   `tfsdk:"statistic" json:"statistic,omitempty"`
		Tags      *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
		Threshold         *float64 `tfsdk:"threshold" json:"threshold,omitempty"`
		ThresholdMetricID *string  `tfsdk:"threshold_metric_id" json:"thresholdMetricID,omitempty"`
		TreatMissingData  *string  `tfsdk:"treat_missing_data" json:"treatMissingData,omitempty"`
		Unit              *string  `tfsdk:"unit" json:"unit,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CloudwatchServicesK8SAwsMetricAlarmV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_cloudwatch_services_k8s_aws_metric_alarm_v1alpha1_manifest"
}

func (r *CloudwatchServicesK8SAwsMetricAlarmV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "MetricAlarm is the Schema for the MetricAlarms API",
		MarkdownDescription: "MetricAlarm is the Schema for the MetricAlarms API",
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
				Description:         "MetricAlarmSpec defines the desired state of MetricAlarm. The details about a metric alarm.",
				MarkdownDescription: "MetricAlarmSpec defines the desired state of MetricAlarm. The details about a metric alarm.",
				Attributes: map[string]schema.Attribute{
					"actions_enabled": schema.BoolAttribute{
						Description:         "Indicates whether actions should be executed during any changes to the alarm state. The default is TRUE.",
						MarkdownDescription: "Indicates whether actions should be executed during any changes to the alarm state. The default is TRUE.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"alarm_actions": schema.ListAttribute{
						Description:         "The actions to execute when this alarm transitions to the ALARM state from any other state. Each action is specified as an Amazon Resource Name (ARN). Valid values: EC2 actions: * arn:aws:automate:region:ec2:stop * arn:aws:automate:region:ec2:terminate * arn:aws:automate:region:ec2:reboot * arn:aws:automate:region:ec2:recover * arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Stop/1.0 * arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Terminate/1.0 * arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Reboot/1.0 * arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Recover/1.0 Autoscaling action: * arn:aws:autoscaling:region:account-id:scalingPolicy:policy-id:autoScalingGroupName/group-friendly-name:policyName/policy-friendly-name SNS notification action: * arn:aws:sns:region:account-id:sns-topic-name:autoScalingGroupName/group-friendly-name:policyName/policy-friendly-name SSM integration actions: * arn:aws:ssm:region:account-id:opsitem:severity#CATEGORY=category-name * arn:aws:ssm-incidents::account-id:responseplan/response-plan-name",
						MarkdownDescription: "The actions to execute when this alarm transitions to the ALARM state from any other state. Each action is specified as an Amazon Resource Name (ARN). Valid values: EC2 actions: * arn:aws:automate:region:ec2:stop * arn:aws:automate:region:ec2:terminate * arn:aws:automate:region:ec2:reboot * arn:aws:automate:region:ec2:recover * arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Stop/1.0 * arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Terminate/1.0 * arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Reboot/1.0 * arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Recover/1.0 Autoscaling action: * arn:aws:autoscaling:region:account-id:scalingPolicy:policy-id:autoScalingGroupName/group-friendly-name:policyName/policy-friendly-name SNS notification action: * arn:aws:sns:region:account-id:sns-topic-name:autoScalingGroupName/group-friendly-name:policyName/policy-friendly-name SSM integration actions: * arn:aws:ssm:region:account-id:opsitem:severity#CATEGORY=category-name * arn:aws:ssm-incidents::account-id:responseplan/response-plan-name",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"alarm_description": schema.StringAttribute{
						Description:         "The description for the alarm.",
						MarkdownDescription: "The description for the alarm.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"comparison_operator": schema.StringAttribute{
						Description:         "The arithmetic operation to use when comparing the specified statistic and threshold. The specified statistic value is used as the first operand. The values LessThanLowerOrGreaterThanUpperThreshold, LessThanLowerThreshold, and GreaterThanUpperThreshold are used only for alarms based on anomaly detection models.",
						MarkdownDescription: "The arithmetic operation to use when comparing the specified statistic and threshold. The specified statistic value is used as the first operand. The values LessThanLowerOrGreaterThanUpperThreshold, LessThanLowerThreshold, and GreaterThanUpperThreshold are used only for alarms based on anomaly detection models.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"datapoints_to_alarm": schema.Int64Attribute{
						Description:         "The number of data points that must be breaching to trigger the alarm. This is used only if you are setting an 'M out of N' alarm. In that case, this value is the M. For more information, see Evaluating an Alarm (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarm-evaluation) in the Amazon CloudWatch User Guide.",
						MarkdownDescription: "The number of data points that must be breaching to trigger the alarm. This is used only if you are setting an 'M out of N' alarm. In that case, this value is the M. For more information, see Evaluating an Alarm (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarm-evaluation) in the Amazon CloudWatch User Guide.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"dimensions": schema.ListNestedAttribute{
						Description:         "The dimensions for the metric specified in MetricName.",
						MarkdownDescription: "The dimensions for the metric specified in MetricName.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
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

					"evaluate_low_sample_count_percentile": schema.StringAttribute{
						Description:         "Used only for alarms based on percentiles. If you specify ignore, the alarm state does not change during periods with too few data points to be statistically significant. If you specify evaluate or omit this parameter, the alarm is always evaluated and possibly changes state no matter how many data points are available. For more information, see Percentile-Based CloudWatch Alarms and Low Data Samples (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#percentiles-with-low-samples). Valid Values: evaluate | ignore",
						MarkdownDescription: "Used only for alarms based on percentiles. If you specify ignore, the alarm state does not change during periods with too few data points to be statistically significant. If you specify evaluate or omit this parameter, the alarm is always evaluated and possibly changes state no matter how many data points are available. For more information, see Percentile-Based CloudWatch Alarms and Low Data Samples (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#percentiles-with-low-samples). Valid Values: evaluate | ignore",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"evaluation_periods": schema.Int64Attribute{
						Description:         "The number of periods over which data is compared to the specified threshold. If you are setting an alarm that requires that a number of consecutive data points be breaching to trigger the alarm, this value specifies that number. If you are setting an 'M out of N' alarm, this value is the N. An alarm's total current evaluation period can be no longer than one day, so this number multiplied by Period cannot be more than 86,400 seconds.",
						MarkdownDescription: "The number of periods over which data is compared to the specified threshold. If you are setting an alarm that requires that a number of consecutive data points be breaching to trigger the alarm, this value specifies that number. If you are setting an 'M out of N' alarm, this value is the N. An alarm's total current evaluation period can be no longer than one day, so this number multiplied by Period cannot be more than 86,400 seconds.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"extended_statistic": schema.StringAttribute{
						Description:         "The extended statistic for the metric specified in MetricName. When you call PutMetricAlarm and specify a MetricName, you must specify either Statistic or ExtendedStatistic but not both. If you specify ExtendedStatistic, the following are valid values: * p90 * tm90 * tc90 * ts90 * wm90 * IQM * PR(n:m) where n and m are values of the metric * TC(X%:X%) where X is between 10 and 90 inclusive. * TM(X%:X%) where X is between 10 and 90 inclusive. * TS(X%:X%) where X is between 10 and 90 inclusive. * WM(X%:X%) where X is between 10 and 90 inclusive. For more information about these extended statistics, see CloudWatch statistics definitions (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html).",
						MarkdownDescription: "The extended statistic for the metric specified in MetricName. When you call PutMetricAlarm and specify a MetricName, you must specify either Statistic or ExtendedStatistic but not both. If you specify ExtendedStatistic, the following are valid values: * p90 * tm90 * tc90 * ts90 * wm90 * IQM * PR(n:m) where n and m are values of the metric * TC(X%:X%) where X is between 10 and 90 inclusive. * TM(X%:X%) where X is between 10 and 90 inclusive. * TS(X%:X%) where X is between 10 and 90 inclusive. * WM(X%:X%) where X is between 10 and 90 inclusive. For more information about these extended statistics, see CloudWatch statistics definitions (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"insufficient_data_actions": schema.ListAttribute{
						Description:         "The actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state. Each action is specified as an Amazon Resource Name (ARN). Valid values: EC2 actions: * arn:aws:automate:region:ec2:stop * arn:aws:automate:region:ec2:terminate * arn:aws:automate:region:ec2:reboot * arn:aws:automate:region:ec2:recover * arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Stop/1.0 * arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Terminate/1.0 * arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Reboot/1.0 * arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Recover/1.0 Autoscaling action: * arn:aws:autoscaling:region:account-id:scalingPolicy:policy-id:autoScalingGroupName/group-friendly-name:policyName/policy-friendly-name SNS notification action: * arn:aws:sns:region:account-id:sns-topic-name:autoScalingGroupName/group-friendly-name:policyName/policy-friendly-name SSM integration actions: * arn:aws:ssm:region:account-id:opsitem:severity#CATEGORY=category-name * arn:aws:ssm-incidents::account-id:responseplan/response-plan-name",
						MarkdownDescription: "The actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state. Each action is specified as an Amazon Resource Name (ARN). Valid values: EC2 actions: * arn:aws:automate:region:ec2:stop * arn:aws:automate:region:ec2:terminate * arn:aws:automate:region:ec2:reboot * arn:aws:automate:region:ec2:recover * arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Stop/1.0 * arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Terminate/1.0 * arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Reboot/1.0 * arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Recover/1.0 Autoscaling action: * arn:aws:autoscaling:region:account-id:scalingPolicy:policy-id:autoScalingGroupName/group-friendly-name:policyName/policy-friendly-name SNS notification action: * arn:aws:sns:region:account-id:sns-topic-name:autoScalingGroupName/group-friendly-name:policyName/policy-friendly-name SSM integration actions: * arn:aws:ssm:region:account-id:opsitem:severity#CATEGORY=category-name * arn:aws:ssm-incidents::account-id:responseplan/response-plan-name",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"metric_name": schema.StringAttribute{
						Description:         "The name for the metric associated with the alarm. For each PutMetricAlarm operation, you must specify either MetricName or a Metrics array. If you are creating an alarm based on a math expression, you cannot specify this parameter, or any of the Namespace, Dimensions, Period, Unit, Statistic, or ExtendedStatistic parameters. Instead, you specify all this information in the Metrics array.",
						MarkdownDescription: "The name for the metric associated with the alarm. For each PutMetricAlarm operation, you must specify either MetricName or a Metrics array. If you are creating an alarm based on a math expression, you cannot specify this parameter, or any of the Namespace, Dimensions, Period, Unit, Statistic, or ExtendedStatistic parameters. Instead, you specify all this information in the Metrics array.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"metrics": schema.ListNestedAttribute{
						Description:         "An array of MetricDataQuery structures that enable you to create an alarm based on the result of a metric math expression. For each PutMetricAlarm operation, you must specify either MetricName or a Metrics array. Each item in the Metrics array either retrieves a metric or performs a math expression. One item in the Metrics array is the expression that the alarm watches. You designate this expression by setting ReturnData to true for this object in the array. For more information, see MetricDataQuery (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_MetricDataQuery.html). If you use the Metrics parameter, you cannot include the Namespace, MetricName, Dimensions, Period, Unit, Statistic, or ExtendedStatistic parameters of PutMetricAlarm in the same operation. Instead, you retrieve the metrics you are using in your math expression as part of the Metrics array.",
						MarkdownDescription: "An array of MetricDataQuery structures that enable you to create an alarm based on the result of a metric math expression. For each PutMetricAlarm operation, you must specify either MetricName or a Metrics array. Each item in the Metrics array either retrieves a metric or performs a math expression. One item in the Metrics array is the expression that the alarm watches. You designate this expression by setting ReturnData to true for this object in the array. For more information, see MetricDataQuery (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_MetricDataQuery.html). If you use the Metrics parameter, you cannot include the Namespace, MetricName, Dimensions, Period, Unit, Statistic, or ExtendedStatistic parameters of PutMetricAlarm in the same operation. Instead, you retrieve the metrics you are using in your math expression as part of the Metrics array.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"account_id": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"expression": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"id": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"label": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metric_stat": schema.SingleNestedAttribute{
									Description:         "This structure defines the metric to be returned, along with the statistics, period, and units.",
									MarkdownDescription: "This structure defines the metric to be returned, along with the statistics, period, and units.",
									Attributes: map[string]schema.Attribute{
										"metric": schema.SingleNestedAttribute{
											Description:         "Represents a specific metric.",
											MarkdownDescription: "Represents a specific metric.",
											Attributes: map[string]schema.Attribute{
												"dimensions": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
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

												"metric_name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespace": schema.StringAttribute{
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

										"period": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"stat": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"unit": schema.StringAttribute{
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

								"period": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"return_data": schema.BoolAttribute{
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

					"name": schema.StringAttribute{
						Description:         "The name for the alarm. This name must be unique within the Region. The name must contain only UTF-8 characters, and can't contain ASCII control characters",
						MarkdownDescription: "The name for the alarm. This name must be unique within the Region. The name must contain only UTF-8 characters, and can't contain ASCII control characters",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"namespace": schema.StringAttribute{
						Description:         "The namespace for the metric associated specified in MetricName.",
						MarkdownDescription: "The namespace for the metric associated specified in MetricName.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"o_k_actions": schema.ListAttribute{
						Description:         "The actions to execute when this alarm transitions to an OK state from any other state. Each action is specified as an Amazon Resource Name (ARN). Valid values: EC2 actions: * arn:aws:automate:region:ec2:stop * arn:aws:automate:region:ec2:terminate * arn:aws:automate:region:ec2:reboot * arn:aws:automate:region:ec2:recover * arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Stop/1.0 * arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Terminate/1.0 * arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Reboot/1.0 * arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Recover/1.0 Autoscaling action: * arn:aws:autoscaling:region:account-id:scalingPolicy:policy-id:autoScalingGroupName/group-friendly-name:policyName/policy-friendly-name SNS notification action: * arn:aws:sns:region:account-id:sns-topic-name:autoScalingGroupName/group-friendly-name:policyName/policy-friendly-name SSM integration actions: * arn:aws:ssm:region:account-id:opsitem:severity#CATEGORY=category-name * arn:aws:ssm-incidents::account-id:responseplan/response-plan-name",
						MarkdownDescription: "The actions to execute when this alarm transitions to an OK state from any other state. Each action is specified as an Amazon Resource Name (ARN). Valid values: EC2 actions: * arn:aws:automate:region:ec2:stop * arn:aws:automate:region:ec2:terminate * arn:aws:automate:region:ec2:reboot * arn:aws:automate:region:ec2:recover * arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Stop/1.0 * arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Terminate/1.0 * arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Reboot/1.0 * arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Recover/1.0 Autoscaling action: * arn:aws:autoscaling:region:account-id:scalingPolicy:policy-id:autoScalingGroupName/group-friendly-name:policyName/policy-friendly-name SNS notification action: * arn:aws:sns:region:account-id:sns-topic-name:autoScalingGroupName/group-friendly-name:policyName/policy-friendly-name SSM integration actions: * arn:aws:ssm:region:account-id:opsitem:severity#CATEGORY=category-name * arn:aws:ssm-incidents::account-id:responseplan/response-plan-name",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"period": schema.Int64Attribute{
						Description:         "The length, in seconds, used each time the metric specified in MetricName is evaluated. Valid values are 10, 30, and any multiple of 60. Period is required for alarms based on static thresholds. If you are creating an alarm based on a metric math expression, you specify the period for each metric within the objects in the Metrics array. Be sure to specify 10 or 30 only for metrics that are stored by a PutMetricData call with a StorageResolution of 1. If you specify a period of 10 or 30 for a metric that does not have sub-minute resolution, the alarm still attempts to gather data at the period rate that you specify. In this case, it does not receive data for the attempts that do not correspond to a one-minute data resolution, and the alarm might often lapse into INSUFFICENT_DATA status. Specifying 10 or 30 also sets this alarm as a high-resolution alarm, which has a higher charge than other alarms. For more information about pricing, see Amazon CloudWatch Pricing (https://aws.amazon.com/cloudwatch/pricing/). An alarm's total current evaluation period can be no longer than one day, so Period multiplied by EvaluationPeriods cannot be more than 86,400 seconds.",
						MarkdownDescription: "The length, in seconds, used each time the metric specified in MetricName is evaluated. Valid values are 10, 30, and any multiple of 60. Period is required for alarms based on static thresholds. If you are creating an alarm based on a metric math expression, you specify the period for each metric within the objects in the Metrics array. Be sure to specify 10 or 30 only for metrics that are stored by a PutMetricData call with a StorageResolution of 1. If you specify a period of 10 or 30 for a metric that does not have sub-minute resolution, the alarm still attempts to gather data at the period rate that you specify. In this case, it does not receive data for the attempts that do not correspond to a one-minute data resolution, and the alarm might often lapse into INSUFFICENT_DATA status. Specifying 10 or 30 also sets this alarm as a high-resolution alarm, which has a higher charge than other alarms. For more information about pricing, see Amazon CloudWatch Pricing (https://aws.amazon.com/cloudwatch/pricing/). An alarm's total current evaluation period can be no longer than one day, so Period multiplied by EvaluationPeriods cannot be more than 86,400 seconds.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"statistic": schema.StringAttribute{
						Description:         "The statistic for the metric specified in MetricName, other than percentile. For percentile statistics, use ExtendedStatistic. When you call PutMetricAlarm and specify a MetricName, you must specify either Statistic or ExtendedStatistic, but not both.",
						MarkdownDescription: "The statistic for the metric specified in MetricName, other than percentile. For percentile statistics, use ExtendedStatistic. When you call PutMetricAlarm and specify a MetricName, you must specify either Statistic or ExtendedStatistic, but not both.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tags": schema.ListNestedAttribute{
						Description:         "A list of key-value pairs to associate with the alarm. You can associate as many as 50 tags with an alarm. To be able to associate tags with the alarm when you create the alarm, you must have the cloudwatch:TagResource permission. Tags can help you organize and categorize your resources. You can also use them to scope user permissions by granting a user permission to access or change only resources with certain tag values. If you are using this operation to update an existing alarm, any tags you specify in this parameter are ignored. To change the tags of an existing alarm, use TagResource (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_TagResource.html) or UntagResource (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_UntagResource.html).",
						MarkdownDescription: "A list of key-value pairs to associate with the alarm. You can associate as many as 50 tags with an alarm. To be able to associate tags with the alarm when you create the alarm, you must have the cloudwatch:TagResource permission. Tags can help you organize and categorize your resources. You can also use them to scope user permissions by granting a user permission to access or change only resources with certain tag values. If you are using this operation to update an existing alarm, any tags you specify in this parameter are ignored. To change the tags of an existing alarm, use TagResource (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_TagResource.html) or UntagResource (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_UntagResource.html).",
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

					"threshold": schema.Float64Attribute{
						Description:         "The value against which the specified statistic is compared. This parameter is required for alarms based on static thresholds, but should not be used for alarms based on anomaly detection models.",
						MarkdownDescription: "The value against which the specified statistic is compared. This parameter is required for alarms based on static thresholds, but should not be used for alarms based on anomaly detection models.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"threshold_metric_id": schema.StringAttribute{
						Description:         "If this is an alarm based on an anomaly detection model, make this value match the ID of the ANOMALY_DETECTION_BAND function. For an example of how to use this parameter, see the Anomaly Detection Model Alarm example on this page. If your alarm uses this parameter, it cannot have Auto Scaling actions.",
						MarkdownDescription: "If this is an alarm based on an anomaly detection model, make this value match the ID of the ANOMALY_DETECTION_BAND function. For an example of how to use this parameter, see the Anomaly Detection Model Alarm example on this page. If your alarm uses this parameter, it cannot have Auto Scaling actions.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"treat_missing_data": schema.StringAttribute{
						Description:         "Sets how this alarm is to handle missing data points. If TreatMissingData is omitted, the default behavior of missing is used. For more information, see Configuring How CloudWatch Alarms Treats Missing Data (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarms-and-missing-data). Valid Values: breaching | notBreaching | ignore | missing Alarms that evaluate metrics in the AWS/DynamoDB namespace always ignore missing data even if you choose a different option for TreatMissingData. When an AWS/DynamoDB metric has missing data, alarms that evaluate that metric remain in their current state.",
						MarkdownDescription: "Sets how this alarm is to handle missing data points. If TreatMissingData is omitted, the default behavior of missing is used. For more information, see Configuring How CloudWatch Alarms Treats Missing Data (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarms-and-missing-data). Valid Values: breaching | notBreaching | ignore | missing Alarms that evaluate metrics in the AWS/DynamoDB namespace always ignore missing data even if you choose a different option for TreatMissingData. When an AWS/DynamoDB metric has missing data, alarms that evaluate that metric remain in their current state.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"unit": schema.StringAttribute{
						Description:         "The unit of measure for the statistic. For example, the units for the Amazon EC2 NetworkIn metric are Bytes because NetworkIn tracks the number of bytes that an instance receives on all network interfaces. You can also specify a unit when you create a custom metric. Units help provide conceptual meaning to your data. Metric data points that specify a unit of measure, such as Percent, are aggregated separately. If you are creating an alarm based on a metric math expression, you can specify the unit for each metric (if needed) within the objects in the Metrics array. If you don't specify Unit, CloudWatch retrieves all unit types that have been published for the metric and attempts to evaluate the alarm. Usually, metrics are published with only one unit, so the alarm works as intended. However, if the metric is published with multiple types of units and you don't specify a unit, the alarm's behavior is not defined and it behaves unpredictably. We recommend omitting Unit so that you don't inadvertently specify an incorrect unit that is not published for this metric. Doing so causes the alarm to be stuck in the INSUFFICIENT DATA state.",
						MarkdownDescription: "The unit of measure for the statistic. For example, the units for the Amazon EC2 NetworkIn metric are Bytes because NetworkIn tracks the number of bytes that an instance receives on all network interfaces. You can also specify a unit when you create a custom metric. Units help provide conceptual meaning to your data. Metric data points that specify a unit of measure, such as Percent, are aggregated separately. If you are creating an alarm based on a metric math expression, you can specify the unit for each metric (if needed) within the objects in the Metrics array. If you don't specify Unit, CloudWatch retrieves all unit types that have been published for the metric and attempts to evaluate the alarm. Usually, metrics are published with only one unit, so the alarm works as intended. However, if the metric is published with multiple types of units and you don't specify a unit, the alarm's behavior is not defined and it behaves unpredictably. We recommend omitting Unit so that you don't inadvertently specify an incorrect unit that is not published for this metric. Doing so causes the alarm to be stuck in the INSUFFICIENT DATA state.",
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

func (r *CloudwatchServicesK8SAwsMetricAlarmV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_cloudwatch_services_k8s_aws_metric_alarm_v1alpha1_manifest")

	var model CloudwatchServicesK8SAwsMetricAlarmV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("cloudwatch.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("MetricAlarm")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
