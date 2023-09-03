/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package applicationautoscaling_services_k8s_aws_v1alpha1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &ApplicationautoscalingServicesK8SAwsScalingPolicyV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &ApplicationautoscalingServicesK8SAwsScalingPolicyV1Alpha1DataSource{}
)

func NewApplicationautoscalingServicesK8SAwsScalingPolicyV1Alpha1DataSource() datasource.DataSource {
	return &ApplicationautoscalingServicesK8SAwsScalingPolicyV1Alpha1DataSource{}
}

type ApplicationautoscalingServicesK8SAwsScalingPolicyV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type ApplicationautoscalingServicesK8SAwsScalingPolicyV1Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		PolicyName                     *string `tfsdk:"policy_name" json:"policyName,omitempty"`
		PolicyType                     *string `tfsdk:"policy_type" json:"policyType,omitempty"`
		ResourceID                     *string `tfsdk:"resource_id" json:"resourceID,omitempty"`
		ScalableDimension              *string `tfsdk:"scalable_dimension" json:"scalableDimension,omitempty"`
		ServiceNamespace               *string `tfsdk:"service_namespace" json:"serviceNamespace,omitempty"`
		StepScalingPolicyConfiguration *struct {
			AdjustmentType         *string `tfsdk:"adjustment_type" json:"adjustmentType,omitempty"`
			Cooldown               *int64  `tfsdk:"cooldown" json:"cooldown,omitempty"`
			MetricAggregationType  *string `tfsdk:"metric_aggregation_type" json:"metricAggregationType,omitempty"`
			MinAdjustmentMagnitude *int64  `tfsdk:"min_adjustment_magnitude" json:"minAdjustmentMagnitude,omitempty"`
			StepAdjustments        *[]struct {
				MetricIntervalLowerBound *float64 `tfsdk:"metric_interval_lower_bound" json:"metricIntervalLowerBound,omitempty"`
				MetricIntervalUpperBound *float64 `tfsdk:"metric_interval_upper_bound" json:"metricIntervalUpperBound,omitempty"`
				ScalingAdjustment        *int64   `tfsdk:"scaling_adjustment" json:"scalingAdjustment,omitempty"`
			} `tfsdk:"step_adjustments" json:"stepAdjustments,omitempty"`
		} `tfsdk:"step_scaling_policy_configuration" json:"stepScalingPolicyConfiguration,omitempty"`
		TargetTrackingScalingPolicyConfiguration *struct {
			CustomizedMetricSpecification *struct {
				Dimensions *[]struct {
					Name  *string `tfsdk:"name" json:"name,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"dimensions" json:"dimensions,omitempty"`
				MetricName *string `tfsdk:"metric_name" json:"metricName,omitempty"`
				Namespace  *string `tfsdk:"namespace" json:"namespace,omitempty"`
				Statistic  *string `tfsdk:"statistic" json:"statistic,omitempty"`
				Unit       *string `tfsdk:"unit" json:"unit,omitempty"`
			} `tfsdk:"customized_metric_specification" json:"customizedMetricSpecification,omitempty"`
			DisableScaleIn                *bool `tfsdk:"disable_scale_in" json:"disableScaleIn,omitempty"`
			PredefinedMetricSpecification *struct {
				PredefinedMetricType *string `tfsdk:"predefined_metric_type" json:"predefinedMetricType,omitempty"`
				ResourceLabel        *string `tfsdk:"resource_label" json:"resourceLabel,omitempty"`
			} `tfsdk:"predefined_metric_specification" json:"predefinedMetricSpecification,omitempty"`
			ScaleInCooldown  *int64   `tfsdk:"scale_in_cooldown" json:"scaleInCooldown,omitempty"`
			ScaleOutCooldown *int64   `tfsdk:"scale_out_cooldown" json:"scaleOutCooldown,omitempty"`
			TargetValue      *float64 `tfsdk:"target_value" json:"targetValue,omitempty"`
		} `tfsdk:"target_tracking_scaling_policy_configuration" json:"targetTrackingScalingPolicyConfiguration,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ApplicationautoscalingServicesK8SAwsScalingPolicyV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_applicationautoscaling_services_k8s_aws_scaling_policy_v1alpha1"
}

func (r *ApplicationautoscalingServicesK8SAwsScalingPolicyV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ScalingPolicy is the Schema for the ScalingPolicies API",
		MarkdownDescription: "ScalingPolicy is the Schema for the ScalingPolicies API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
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
						Optional:            false,
						Computed:            true,
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "ScalingPolicySpec defines the desired state of ScalingPolicy.  Represents a scaling policy to use with Application Auto Scaling.  For more information about configuring scaling policies for a specific service, see Getting started with Application Auto Scaling (https://docs.aws.amazon.com/autoscaling/application/userguide/getting-started.html) in the Application Auto Scaling User Guide.",
				MarkdownDescription: "ScalingPolicySpec defines the desired state of ScalingPolicy.  Represents a scaling policy to use with Application Auto Scaling.  For more information about configuring scaling policies for a specific service, see Getting started with Application Auto Scaling (https://docs.aws.amazon.com/autoscaling/application/userguide/getting-started.html) in the Application Auto Scaling User Guide.",
				Attributes: map[string]schema.Attribute{
					"policy_name": schema.StringAttribute{
						Description:         "The name of the scaling policy.",
						MarkdownDescription: "The name of the scaling policy.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"policy_type": schema.StringAttribute{
						Description:         "The policy type. This parameter is required if you are creating a scaling policy.  The following policy types are supported:  TargetTrackingScaling—Not supported for Amazon EMR  StepScaling—Not supported for DynamoDB, Amazon Comprehend, Lambda, Amazon Keyspaces, Amazon MSK, Amazon ElastiCache, or Neptune.  For more information, see Target tracking scaling policies (https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-target-tracking.html) and Step scaling policies (https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-step-scaling-policies.html) in the Application Auto Scaling User Guide.",
						MarkdownDescription: "The policy type. This parameter is required if you are creating a scaling policy.  The following policy types are supported:  TargetTrackingScaling—Not supported for Amazon EMR  StepScaling—Not supported for DynamoDB, Amazon Comprehend, Lambda, Amazon Keyspaces, Amazon MSK, Amazon ElastiCache, or Neptune.  For more information, see Target tracking scaling policies (https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-target-tracking.html) and Step scaling policies (https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-step-scaling-policies.html) in the Application Auto Scaling User Guide.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"resource_id": schema.StringAttribute{
						Description:         "The identifier of the resource associated with the scaling policy. This string consists of the resource type and unique identifier.  * ECS service - The resource type is service and the unique identifier is the cluster name and service name. Example: service/default/sample-webapp.  * Spot Fleet - The resource type is spot-fleet-request and the unique identifier is the Spot Fleet request ID. Example: spot-fleet-request/sfr-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE.  * EMR cluster - The resource type is instancegroup and the unique identifier is the cluster ID and instance group ID. Example: instancegroup/j-2EEZNYKUA1NTV/ig-1791Y4E1L8YI0.  * AppStream 2.0 fleet - The resource type is fleet and the unique identifier is the fleet name. Example: fleet/sample-fleet.  * DynamoDB table - The resource type is table and the unique identifier is the table name. Example: table/my-table.  * DynamoDB global secondary index - The resource type is index and the unique identifier is the index name. Example: table/my-table/index/my-table-index.  * Aurora DB cluster - The resource type is cluster and the unique identifier is the cluster name. Example: cluster:my-db-cluster.  * SageMaker endpoint variant - The resource type is variant and the unique identifier is the resource ID. Example: endpoint/my-end-point/variant/KMeansClustering.  * Custom resources are not supported with a resource type. This parameter must specify the OutputValue from the CloudFormation template stack used to access the resources. The unique identifier is defined by the service provider. More information is available in our GitHub repository (https://github.com/aws/aws-auto-scaling-custom-resource).  * Amazon Comprehend document classification endpoint - The resource type and unique identifier are specified using the endpoint ARN. Example: arn:aws:comprehend:us-west-2:123456789012:document-classifier-endpoint/EXAMPLE.  * Amazon Comprehend entity recognizer endpoint - The resource type and unique identifier are specified using the endpoint ARN. Example: arn:aws:comprehend:us-west-2:123456789012:entity-recognizer-endpoint/EXAMPLE.  * Lambda provisioned concurrency - The resource type is function and the unique identifier is the function name with a function version or alias name suffix that is not $LATEST. Example: function:my-function:prod or function:my-function:1.  * Amazon Keyspaces table - The resource type is table and the unique identifier is the table name. Example: keyspace/mykeyspace/table/mytable.  * Amazon MSK cluster - The resource type and unique identifier are specified using the cluster ARN. Example: arn:aws:kafka:us-east-1:123456789012:cluster/demo-cluster-1/6357e0b2-0e6a-4b86-a0b4-70df934c2e31-5.  * Amazon ElastiCache replication group - The resource type is replication-group and the unique identifier is the replication group name. Example: replication-group/mycluster.  * Neptune cluster - The resource type is cluster and the unique identifier is the cluster name. Example: cluster:mycluster.",
						MarkdownDescription: "The identifier of the resource associated with the scaling policy. This string consists of the resource type and unique identifier.  * ECS service - The resource type is service and the unique identifier is the cluster name and service name. Example: service/default/sample-webapp.  * Spot Fleet - The resource type is spot-fleet-request and the unique identifier is the Spot Fleet request ID. Example: spot-fleet-request/sfr-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE.  * EMR cluster - The resource type is instancegroup and the unique identifier is the cluster ID and instance group ID. Example: instancegroup/j-2EEZNYKUA1NTV/ig-1791Y4E1L8YI0.  * AppStream 2.0 fleet - The resource type is fleet and the unique identifier is the fleet name. Example: fleet/sample-fleet.  * DynamoDB table - The resource type is table and the unique identifier is the table name. Example: table/my-table.  * DynamoDB global secondary index - The resource type is index and the unique identifier is the index name. Example: table/my-table/index/my-table-index.  * Aurora DB cluster - The resource type is cluster and the unique identifier is the cluster name. Example: cluster:my-db-cluster.  * SageMaker endpoint variant - The resource type is variant and the unique identifier is the resource ID. Example: endpoint/my-end-point/variant/KMeansClustering.  * Custom resources are not supported with a resource type. This parameter must specify the OutputValue from the CloudFormation template stack used to access the resources. The unique identifier is defined by the service provider. More information is available in our GitHub repository (https://github.com/aws/aws-auto-scaling-custom-resource).  * Amazon Comprehend document classification endpoint - The resource type and unique identifier are specified using the endpoint ARN. Example: arn:aws:comprehend:us-west-2:123456789012:document-classifier-endpoint/EXAMPLE.  * Amazon Comprehend entity recognizer endpoint - The resource type and unique identifier are specified using the endpoint ARN. Example: arn:aws:comprehend:us-west-2:123456789012:entity-recognizer-endpoint/EXAMPLE.  * Lambda provisioned concurrency - The resource type is function and the unique identifier is the function name with a function version or alias name suffix that is not $LATEST. Example: function:my-function:prod or function:my-function:1.  * Amazon Keyspaces table - The resource type is table and the unique identifier is the table name. Example: keyspace/mykeyspace/table/mytable.  * Amazon MSK cluster - The resource type and unique identifier are specified using the cluster ARN. Example: arn:aws:kafka:us-east-1:123456789012:cluster/demo-cluster-1/6357e0b2-0e6a-4b86-a0b4-70df934c2e31-5.  * Amazon ElastiCache replication group - The resource type is replication-group and the unique identifier is the replication group name. Example: replication-group/mycluster.  * Neptune cluster - The resource type is cluster and the unique identifier is the cluster name. Example: cluster:mycluster.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"scalable_dimension": schema.StringAttribute{
						Description:         "The scalable dimension. This string consists of the service namespace, resource type, and scaling property.  * ecs:service:DesiredCount - The desired task count of an ECS service.  * elasticmapreduce:instancegroup:InstanceCount - The instance count of an EMR Instance Group.  * ec2:spot-fleet-request:TargetCapacity - The target capacity of a Spot Fleet.  * appstream:fleet:DesiredCapacity - The desired capacity of an AppStream 2.0 fleet.  * dynamodb:table:ReadCapacityUnits - The provisioned read capacity for a DynamoDB table.  * dynamodb:table:WriteCapacityUnits - The provisioned write capacity for a DynamoDB table.  * dynamodb:index:ReadCapacityUnits - The provisioned read capacity for a DynamoDB global secondary index.  * dynamodb:index:WriteCapacityUnits - The provisioned write capacity for a DynamoDB global secondary index.  * rds:cluster:ReadReplicaCount - The count of Aurora Replicas in an Aurora DB cluster. Available for Aurora MySQL-compatible edition and Aurora PostgreSQL-compatible edition.  * sagemaker:variant:DesiredInstanceCount - The number of EC2 instances for an SageMaker model endpoint variant.  * custom-resource:ResourceType:Property - The scalable dimension for a custom resource provided by your own application or service.  * comprehend:document-classifier-endpoint:DesiredInferenceUnits - The number of inference units for an Amazon Comprehend document classification endpoint.  * comprehend:entity-recognizer-endpoint:DesiredInferenceUnits - The number of inference units for an Amazon Comprehend entity recognizer endpoint.  * lambda:function:ProvisionedConcurrency - The provisioned concurrency for a Lambda function.  * cassandra:table:ReadCapacityUnits - The provisioned read capacity for an Amazon Keyspaces table.  * cassandra:table:WriteCapacityUnits - The provisioned write capacity for an Amazon Keyspaces table.  * kafka:broker-storage:VolumeSize - The provisioned volume size (in GiB) for brokers in an Amazon MSK cluster.  * elasticache:replication-group:NodeGroups - The number of node groups for an Amazon ElastiCache replication group.  * elasticache:replication-group:Replicas - The number of replicas per node group for an Amazon ElastiCache replication group.  * neptune:cluster:ReadReplicaCount - The count of read replicas in an Amazon Neptune DB cluster.",
						MarkdownDescription: "The scalable dimension. This string consists of the service namespace, resource type, and scaling property.  * ecs:service:DesiredCount - The desired task count of an ECS service.  * elasticmapreduce:instancegroup:InstanceCount - The instance count of an EMR Instance Group.  * ec2:spot-fleet-request:TargetCapacity - The target capacity of a Spot Fleet.  * appstream:fleet:DesiredCapacity - The desired capacity of an AppStream 2.0 fleet.  * dynamodb:table:ReadCapacityUnits - The provisioned read capacity for a DynamoDB table.  * dynamodb:table:WriteCapacityUnits - The provisioned write capacity for a DynamoDB table.  * dynamodb:index:ReadCapacityUnits - The provisioned read capacity for a DynamoDB global secondary index.  * dynamodb:index:WriteCapacityUnits - The provisioned write capacity for a DynamoDB global secondary index.  * rds:cluster:ReadReplicaCount - The count of Aurora Replicas in an Aurora DB cluster. Available for Aurora MySQL-compatible edition and Aurora PostgreSQL-compatible edition.  * sagemaker:variant:DesiredInstanceCount - The number of EC2 instances for an SageMaker model endpoint variant.  * custom-resource:ResourceType:Property - The scalable dimension for a custom resource provided by your own application or service.  * comprehend:document-classifier-endpoint:DesiredInferenceUnits - The number of inference units for an Amazon Comprehend document classification endpoint.  * comprehend:entity-recognizer-endpoint:DesiredInferenceUnits - The number of inference units for an Amazon Comprehend entity recognizer endpoint.  * lambda:function:ProvisionedConcurrency - The provisioned concurrency for a Lambda function.  * cassandra:table:ReadCapacityUnits - The provisioned read capacity for an Amazon Keyspaces table.  * cassandra:table:WriteCapacityUnits - The provisioned write capacity for an Amazon Keyspaces table.  * kafka:broker-storage:VolumeSize - The provisioned volume size (in GiB) for brokers in an Amazon MSK cluster.  * elasticache:replication-group:NodeGroups - The number of node groups for an Amazon ElastiCache replication group.  * elasticache:replication-group:Replicas - The number of replicas per node group for an Amazon ElastiCache replication group.  * neptune:cluster:ReadReplicaCount - The count of read replicas in an Amazon Neptune DB cluster.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"service_namespace": schema.StringAttribute{
						Description:         "The namespace of the Amazon Web Services service that provides the resource. For a resource provided by your own application or service, use custom-resource instead.",
						MarkdownDescription: "The namespace of the Amazon Web Services service that provides the resource. For a resource provided by your own application or service, use custom-resource instead.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"step_scaling_policy_configuration": schema.SingleNestedAttribute{
						Description:         "A step scaling policy.  This parameter is required if you are creating a policy and the policy type is StepScaling.",
						MarkdownDescription: "A step scaling policy.  This parameter is required if you are creating a policy and the policy type is StepScaling.",
						Attributes: map[string]schema.Attribute{
							"adjustment_type": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"cooldown": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"metric_aggregation_type": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"min_adjustment_magnitude": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"step_adjustments": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"metric_interval_lower_bound": schema.Float64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"metric_interval_upper_bound": schema.Float64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"scaling_adjustment": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"target_tracking_scaling_policy_configuration": schema.SingleNestedAttribute{
						Description:         "A target tracking scaling policy. Includes support for predefined or customized metrics.  This parameter is required if you are creating a policy and the policy type is TargetTrackingScaling.",
						MarkdownDescription: "A target tracking scaling policy. Includes support for predefined or customized metrics.  This parameter is required if you are creating a policy and the policy type is TargetTrackingScaling.",
						Attributes: map[string]schema.Attribute{
							"customized_metric_specification": schema.SingleNestedAttribute{
								Description:         "Represents a CloudWatch metric of your choosing for a target tracking scaling policy to use with Application Auto Scaling.  For information about the available metrics for a service, see Amazon Web Services Services That Publish CloudWatch Metrics (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/aws-services-cloudwatch-metrics.html) in the Amazon CloudWatch User Guide.  To create your customized metric specification:  * Add values for each required parameter from CloudWatch. You can use an existing metric, or a new metric that you create. To use your own metric, you must first publish the metric to CloudWatch. For more information, see Publish Custom Metrics (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/publishingMetrics.html) in the Amazon CloudWatch User Guide.  * Choose a metric that changes proportionally with capacity. The value of the metric should increase or decrease in inverse proportion to the number of capacity units. That is, the value of the metric should decrease when capacity increases, and increase when capacity decreases.  For more information about CloudWatch, see Amazon CloudWatch Concepts (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_concepts.html).",
								MarkdownDescription: "Represents a CloudWatch metric of your choosing for a target tracking scaling policy to use with Application Auto Scaling.  For information about the available metrics for a service, see Amazon Web Services Services That Publish CloudWatch Metrics (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/aws-services-cloudwatch-metrics.html) in the Amazon CloudWatch User Guide.  To create your customized metric specification:  * Add values for each required parameter from CloudWatch. You can use an existing metric, or a new metric that you create. To use your own metric, you must first publish the metric to CloudWatch. For more information, see Publish Custom Metrics (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/publishingMetrics.html) in the Amazon CloudWatch User Guide.  * Choose a metric that changes proportionally with capacity. The value of the metric should increase or decrease in inverse proportion to the number of capacity units. That is, the value of the metric should decrease when capacity increases, and increase when capacity decreases.  For more information about CloudWatch, see Amazon CloudWatch Concepts (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_concepts.html).",
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
													Optional:            false,
													Computed:            true,
												},

												"value": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"metric_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"namespace": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"statistic": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"unit": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"disable_scale_in": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"predefined_metric_specification": schema.SingleNestedAttribute{
								Description:         "Represents a predefined metric for a target tracking scaling policy to use with Application Auto Scaling.  Only the Amazon Web Services that you're using send metrics to Amazon CloudWatch. To determine whether a desired metric already exists by looking up its namespace and dimension using the CloudWatch metrics dashboard in the console, follow the procedure in Building dashboards with CloudWatch (https://docs.aws.amazon.com/autoscaling/application/userguide/monitoring-cloudwatch.html) in the Application Auto Scaling User Guide.",
								MarkdownDescription: "Represents a predefined metric for a target tracking scaling policy to use with Application Auto Scaling.  Only the Amazon Web Services that you're using send metrics to Amazon CloudWatch. To determine whether a desired metric already exists by looking up its namespace and dimension using the CloudWatch metrics dashboard in the console, follow the procedure in Building dashboards with CloudWatch (https://docs.aws.amazon.com/autoscaling/application/userguide/monitoring-cloudwatch.html) in the Application Auto Scaling User Guide.",
								Attributes: map[string]schema.Attribute{
									"predefined_metric_type": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"resource_label": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"scale_in_cooldown": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"scale_out_cooldown": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"target_value": schema.Float64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *ApplicationautoscalingServicesK8SAwsScalingPolicyV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *ApplicationautoscalingServicesK8SAwsScalingPolicyV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_applicationautoscaling_services_k8s_aws_scaling_policy_v1alpha1")

	var data ApplicationautoscalingServicesK8SAwsScalingPolicyV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "applicationautoscaling.services.k8s.aws", Version: "v1alpha1", Resource: "ScalingPolicy"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse ApplicationautoscalingServicesK8SAwsScalingPolicyV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("applicationautoscaling.services.k8s.aws/v1alpha1")
	data.Kind = pointer.String("ScalingPolicy")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
