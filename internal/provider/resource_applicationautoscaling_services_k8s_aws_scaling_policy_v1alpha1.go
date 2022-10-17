/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"gopkg.in/yaml.v3"
	"time"
)

type ApplicationautoscalingServicesK8SAwsScalingPolicyV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*ApplicationautoscalingServicesK8SAwsScalingPolicyV1Alpha1Resource)(nil)
)

type ApplicationautoscalingServicesK8SAwsScalingPolicyV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type ApplicationautoscalingServicesK8SAwsScalingPolicyV1Alpha1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		PolicyName *string `tfsdk:"policy_name" yaml:"policyName,omitempty"`

		PolicyType *string `tfsdk:"policy_type" yaml:"policyType,omitempty"`

		ResourceID *string `tfsdk:"resource_id" yaml:"resourceID,omitempty"`

		ScalableDimension *string `tfsdk:"scalable_dimension" yaml:"scalableDimension,omitempty"`

		ServiceNamespace *string `tfsdk:"service_namespace" yaml:"serviceNamespace,omitempty"`

		StepScalingPolicyConfiguration *struct {
			AdjustmentType *string `tfsdk:"adjustment_type" yaml:"adjustmentType,omitempty"`

			Cooldown *int64 `tfsdk:"cooldown" yaml:"cooldown,omitempty"`

			MetricAggregationType *string `tfsdk:"metric_aggregation_type" yaml:"metricAggregationType,omitempty"`

			MinAdjustmentMagnitude *int64 `tfsdk:"min_adjustment_magnitude" yaml:"minAdjustmentMagnitude,omitempty"`

			StepAdjustments *[]struct {
				MetricIntervalLowerBound utilities.DynamicNumber `tfsdk:"metric_interval_lower_bound" yaml:"metricIntervalLowerBound,omitempty"`

				MetricIntervalUpperBound utilities.DynamicNumber `tfsdk:"metric_interval_upper_bound" yaml:"metricIntervalUpperBound,omitempty"`

				ScalingAdjustment *int64 `tfsdk:"scaling_adjustment" yaml:"scalingAdjustment,omitempty"`
			} `tfsdk:"step_adjustments" yaml:"stepAdjustments,omitempty"`
		} `tfsdk:"step_scaling_policy_configuration" yaml:"stepScalingPolicyConfiguration,omitempty"`

		TargetTrackingScalingPolicyConfiguration *struct {
			CustomizedMetricSpecification *struct {
				Dimensions *[]struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"dimensions" yaml:"dimensions,omitempty"`

				MetricName *string `tfsdk:"metric_name" yaml:"metricName,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

				Statistic *string `tfsdk:"statistic" yaml:"statistic,omitempty"`

				Unit *string `tfsdk:"unit" yaml:"unit,omitempty"`
			} `tfsdk:"customized_metric_specification" yaml:"customizedMetricSpecification,omitempty"`

			DisableScaleIn *bool `tfsdk:"disable_scale_in" yaml:"disableScaleIn,omitempty"`

			PredefinedMetricSpecification *struct {
				PredefinedMetricType *string `tfsdk:"predefined_metric_type" yaml:"predefinedMetricType,omitempty"`

				ResourceLabel *string `tfsdk:"resource_label" yaml:"resourceLabel,omitempty"`
			} `tfsdk:"predefined_metric_specification" yaml:"predefinedMetricSpecification,omitempty"`

			ScaleInCooldown *int64 `tfsdk:"scale_in_cooldown" yaml:"scaleInCooldown,omitempty"`

			ScaleOutCooldown *int64 `tfsdk:"scale_out_cooldown" yaml:"scaleOutCooldown,omitempty"`

			TargetValue utilities.DynamicNumber `tfsdk:"target_value" yaml:"targetValue,omitempty"`
		} `tfsdk:"target_tracking_scaling_policy_configuration" yaml:"targetTrackingScalingPolicyConfiguration,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewApplicationautoscalingServicesK8SAwsScalingPolicyV1Alpha1Resource() resource.Resource {
	return &ApplicationautoscalingServicesK8SAwsScalingPolicyV1Alpha1Resource{}
}

func (r *ApplicationautoscalingServicesK8SAwsScalingPolicyV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_applicationautoscaling_services_k8s_aws_scaling_policy_v1alpha1"
}

func (r *ApplicationautoscalingServicesK8SAwsScalingPolicyV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "ScalingPolicy is the Schema for the ScalingPolicies API",
		MarkdownDescription: "ScalingPolicy is the Schema for the ScalingPolicies API",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description:         "The timestamp of the last change to this resource.",
				MarkdownDescription: "The timestamp of the last change to this resource.",
				Type:                types.Int64Type,
				Computed:            true,
				Optional:            false,
			},

			"yaml": {
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"metadata": {
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Type:                types.StringType,
						Required:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
					},

					"namespace": {
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Type:                types.StringType,
						Optional:            true,
					},

					"labels": {
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.LabelValidator(),
						},
					},
					"annotations": {
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.AnnotationValidator(),
						},
					},
				}),
			},

			"api_version": {
				Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"kind": {
				Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"spec": {
				Description:         "ScalingPolicySpec defines the desired state of ScalingPolicy.  Represents a scaling policy to use with Application Auto Scaling.  For more information about configuring scaling policies for a specific service, see Getting started with Application Auto Scaling (https://docs.aws.amazon.com/autoscaling/application/userguide/getting-started.html) in the Application Auto Scaling User Guide.",
				MarkdownDescription: "ScalingPolicySpec defines the desired state of ScalingPolicy.  Represents a scaling policy to use with Application Auto Scaling.  For more information about configuring scaling policies for a specific service, see Getting started with Application Auto Scaling (https://docs.aws.amazon.com/autoscaling/application/userguide/getting-started.html) in the Application Auto Scaling User Guide.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"policy_name": {
						Description:         "The name of the scaling policy.",
						MarkdownDescription: "The name of the scaling policy.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"policy_type": {
						Description:         "The policy type. This parameter is required if you are creating a scaling policy.  The following policy types are supported:  TargetTrackingScaling—Not supported for Amazon EMR  StepScaling—Not supported for DynamoDB, Amazon Comprehend, Lambda, Amazon Keyspaces, Amazon MSK, Amazon ElastiCache, or Neptune.  For more information, see Target tracking scaling policies (https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-target-tracking.html) and Step scaling policies (https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-step-scaling-policies.html) in the Application Auto Scaling User Guide.",
						MarkdownDescription: "The policy type. This parameter is required if you are creating a scaling policy.  The following policy types are supported:  TargetTrackingScaling—Not supported for Amazon EMR  StepScaling—Not supported for DynamoDB, Amazon Comprehend, Lambda, Amazon Keyspaces, Amazon MSK, Amazon ElastiCache, or Neptune.  For more information, see Target tracking scaling policies (https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-target-tracking.html) and Step scaling policies (https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-step-scaling-policies.html) in the Application Auto Scaling User Guide.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"resource_id": {
						Description:         "The identifier of the resource associated with the scaling policy. This string consists of the resource type and unique identifier.  * ECS service - The resource type is service and the unique identifier is the cluster name and service name. Example: service/default/sample-webapp.  * Spot Fleet - The resource type is spot-fleet-request and the unique identifier is the Spot Fleet request ID. Example: spot-fleet-request/sfr-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE.  * EMR cluster - The resource type is instancegroup and the unique identifier is the cluster ID and instance group ID. Example: instancegroup/j-2EEZNYKUA1NTV/ig-1791Y4E1L8YI0.  * AppStream 2.0 fleet - The resource type is fleet and the unique identifier is the fleet name. Example: fleet/sample-fleet.  * DynamoDB table - The resource type is table and the unique identifier is the table name. Example: table/my-table.  * DynamoDB global secondary index - The resource type is index and the unique identifier is the index name. Example: table/my-table/index/my-table-index.  * Aurora DB cluster - The resource type is cluster and the unique identifier is the cluster name. Example: cluster:my-db-cluster.  * SageMaker endpoint variant - The resource type is variant and the unique identifier is the resource ID. Example: endpoint/my-end-point/variant/KMeansClustering.  * Custom resources are not supported with a resource type. This parameter must specify the OutputValue from the CloudFormation template stack used to access the resources. The unique identifier is defined by the service provider. More information is available in our GitHub repository (https://github.com/aws/aws-auto-scaling-custom-resource).  * Amazon Comprehend document classification endpoint - The resource type and unique identifier are specified using the endpoint ARN. Example: arn:aws:comprehend:us-west-2:123456789012:document-classifier-endpoint/EXAMPLE.  * Amazon Comprehend entity recognizer endpoint - The resource type and unique identifier are specified using the endpoint ARN. Example: arn:aws:comprehend:us-west-2:123456789012:entity-recognizer-endpoint/EXAMPLE.  * Lambda provisioned concurrency - The resource type is function and the unique identifier is the function name with a function version or alias name suffix that is not $LATEST. Example: function:my-function:prod or function:my-function:1.  * Amazon Keyspaces table - The resource type is table and the unique identifier is the table name. Example: keyspace/mykeyspace/table/mytable.  * Amazon MSK cluster - The resource type and unique identifier are specified using the cluster ARN. Example: arn:aws:kafka:us-east-1:123456789012:cluster/demo-cluster-1/6357e0b2-0e6a-4b86-a0b4-70df934c2e31-5.  * Amazon ElastiCache replication group - The resource type is replication-group and the unique identifier is the replication group name. Example: replication-group/mycluster.  * Neptune cluster - The resource type is cluster and the unique identifier is the cluster name. Example: cluster:mycluster.",
						MarkdownDescription: "The identifier of the resource associated with the scaling policy. This string consists of the resource type and unique identifier.  * ECS service - The resource type is service and the unique identifier is the cluster name and service name. Example: service/default/sample-webapp.  * Spot Fleet - The resource type is spot-fleet-request and the unique identifier is the Spot Fleet request ID. Example: spot-fleet-request/sfr-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE.  * EMR cluster - The resource type is instancegroup and the unique identifier is the cluster ID and instance group ID. Example: instancegroup/j-2EEZNYKUA1NTV/ig-1791Y4E1L8YI0.  * AppStream 2.0 fleet - The resource type is fleet and the unique identifier is the fleet name. Example: fleet/sample-fleet.  * DynamoDB table - The resource type is table and the unique identifier is the table name. Example: table/my-table.  * DynamoDB global secondary index - The resource type is index and the unique identifier is the index name. Example: table/my-table/index/my-table-index.  * Aurora DB cluster - The resource type is cluster and the unique identifier is the cluster name. Example: cluster:my-db-cluster.  * SageMaker endpoint variant - The resource type is variant and the unique identifier is the resource ID. Example: endpoint/my-end-point/variant/KMeansClustering.  * Custom resources are not supported with a resource type. This parameter must specify the OutputValue from the CloudFormation template stack used to access the resources. The unique identifier is defined by the service provider. More information is available in our GitHub repository (https://github.com/aws/aws-auto-scaling-custom-resource).  * Amazon Comprehend document classification endpoint - The resource type and unique identifier are specified using the endpoint ARN. Example: arn:aws:comprehend:us-west-2:123456789012:document-classifier-endpoint/EXAMPLE.  * Amazon Comprehend entity recognizer endpoint - The resource type and unique identifier are specified using the endpoint ARN. Example: arn:aws:comprehend:us-west-2:123456789012:entity-recognizer-endpoint/EXAMPLE.  * Lambda provisioned concurrency - The resource type is function and the unique identifier is the function name with a function version or alias name suffix that is not $LATEST. Example: function:my-function:prod or function:my-function:1.  * Amazon Keyspaces table - The resource type is table and the unique identifier is the table name. Example: keyspace/mykeyspace/table/mytable.  * Amazon MSK cluster - The resource type and unique identifier are specified using the cluster ARN. Example: arn:aws:kafka:us-east-1:123456789012:cluster/demo-cluster-1/6357e0b2-0e6a-4b86-a0b4-70df934c2e31-5.  * Amazon ElastiCache replication group - The resource type is replication-group and the unique identifier is the replication group name. Example: replication-group/mycluster.  * Neptune cluster - The resource type is cluster and the unique identifier is the cluster name. Example: cluster:mycluster.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"scalable_dimension": {
						Description:         "The scalable dimension. This string consists of the service namespace, resource type, and scaling property.  * ecs:service:DesiredCount - The desired task count of an ECS service.  * elasticmapreduce:instancegroup:InstanceCount - The instance count of an EMR Instance Group.  * ec2:spot-fleet-request:TargetCapacity - The target capacity of a Spot Fleet.  * appstream:fleet:DesiredCapacity - The desired capacity of an AppStream 2.0 fleet.  * dynamodb:table:ReadCapacityUnits - The provisioned read capacity for a DynamoDB table.  * dynamodb:table:WriteCapacityUnits - The provisioned write capacity for a DynamoDB table.  * dynamodb:index:ReadCapacityUnits - The provisioned read capacity for a DynamoDB global secondary index.  * dynamodb:index:WriteCapacityUnits - The provisioned write capacity for a DynamoDB global secondary index.  * rds:cluster:ReadReplicaCount - The count of Aurora Replicas in an Aurora DB cluster. Available for Aurora MySQL-compatible edition and Aurora PostgreSQL-compatible edition.  * sagemaker:variant:DesiredInstanceCount - The number of EC2 instances for an SageMaker model endpoint variant.  * custom-resource:ResourceType:Property - The scalable dimension for a custom resource provided by your own application or service.  * comprehend:document-classifier-endpoint:DesiredInferenceUnits - The number of inference units for an Amazon Comprehend document classification endpoint.  * comprehend:entity-recognizer-endpoint:DesiredInferenceUnits - The number of inference units for an Amazon Comprehend entity recognizer endpoint.  * lambda:function:ProvisionedConcurrency - The provisioned concurrency for a Lambda function.  * cassandra:table:ReadCapacityUnits - The provisioned read capacity for an Amazon Keyspaces table.  * cassandra:table:WriteCapacityUnits - The provisioned write capacity for an Amazon Keyspaces table.  * kafka:broker-storage:VolumeSize - The provisioned volume size (in GiB) for brokers in an Amazon MSK cluster.  * elasticache:replication-group:NodeGroups - The number of node groups for an Amazon ElastiCache replication group.  * elasticache:replication-group:Replicas - The number of replicas per node group for an Amazon ElastiCache replication group.  * neptune:cluster:ReadReplicaCount - The count of read replicas in an Amazon Neptune DB cluster.",
						MarkdownDescription: "The scalable dimension. This string consists of the service namespace, resource type, and scaling property.  * ecs:service:DesiredCount - The desired task count of an ECS service.  * elasticmapreduce:instancegroup:InstanceCount - The instance count of an EMR Instance Group.  * ec2:spot-fleet-request:TargetCapacity - The target capacity of a Spot Fleet.  * appstream:fleet:DesiredCapacity - The desired capacity of an AppStream 2.0 fleet.  * dynamodb:table:ReadCapacityUnits - The provisioned read capacity for a DynamoDB table.  * dynamodb:table:WriteCapacityUnits - The provisioned write capacity for a DynamoDB table.  * dynamodb:index:ReadCapacityUnits - The provisioned read capacity for a DynamoDB global secondary index.  * dynamodb:index:WriteCapacityUnits - The provisioned write capacity for a DynamoDB global secondary index.  * rds:cluster:ReadReplicaCount - The count of Aurora Replicas in an Aurora DB cluster. Available for Aurora MySQL-compatible edition and Aurora PostgreSQL-compatible edition.  * sagemaker:variant:DesiredInstanceCount - The number of EC2 instances for an SageMaker model endpoint variant.  * custom-resource:ResourceType:Property - The scalable dimension for a custom resource provided by your own application or service.  * comprehend:document-classifier-endpoint:DesiredInferenceUnits - The number of inference units for an Amazon Comprehend document classification endpoint.  * comprehend:entity-recognizer-endpoint:DesiredInferenceUnits - The number of inference units for an Amazon Comprehend entity recognizer endpoint.  * lambda:function:ProvisionedConcurrency - The provisioned concurrency for a Lambda function.  * cassandra:table:ReadCapacityUnits - The provisioned read capacity for an Amazon Keyspaces table.  * cassandra:table:WriteCapacityUnits - The provisioned write capacity for an Amazon Keyspaces table.  * kafka:broker-storage:VolumeSize - The provisioned volume size (in GiB) for brokers in an Amazon MSK cluster.  * elasticache:replication-group:NodeGroups - The number of node groups for an Amazon ElastiCache replication group.  * elasticache:replication-group:Replicas - The number of replicas per node group for an Amazon ElastiCache replication group.  * neptune:cluster:ReadReplicaCount - The count of read replicas in an Amazon Neptune DB cluster.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"service_namespace": {
						Description:         "The namespace of the Amazon Web Services service that provides the resource. For a resource provided by your own application or service, use custom-resource instead.",
						MarkdownDescription: "The namespace of the Amazon Web Services service that provides the resource. For a resource provided by your own application or service, use custom-resource instead.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"step_scaling_policy_configuration": {
						Description:         "A step scaling policy.  This parameter is required if you are creating a policy and the policy type is StepScaling.",
						MarkdownDescription: "A step scaling policy.  This parameter is required if you are creating a policy and the policy type is StepScaling.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"adjustment_type": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"cooldown": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"metric_aggregation_type": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"min_adjustment_magnitude": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"step_adjustments": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"metric_interval_lower_bound": {
										Description:         "",
										MarkdownDescription: "",

										Type: utilities.DynamicNumberType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"metric_interval_upper_bound": {
										Description:         "",
										MarkdownDescription: "",

										Type: utilities.DynamicNumberType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"scaling_adjustment": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"target_tracking_scaling_policy_configuration": {
						Description:         "A target tracking scaling policy. Includes support for predefined or customized metrics.  This parameter is required if you are creating a policy and the policy type is TargetTrackingScaling.",
						MarkdownDescription: "A target tracking scaling policy. Includes support for predefined or customized metrics.  This parameter is required if you are creating a policy and the policy type is TargetTrackingScaling.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"customized_metric_specification": {
								Description:         "Represents a CloudWatch metric of your choosing for a target tracking scaling policy to use with Application Auto Scaling.  For information about the available metrics for a service, see Amazon Web Services Services That Publish CloudWatch Metrics (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/aws-services-cloudwatch-metrics.html) in the Amazon CloudWatch User Guide.  To create your customized metric specification:  * Add values for each required parameter from CloudWatch. You can use an existing metric, or a new metric that you create. To use your own metric, you must first publish the metric to CloudWatch. For more information, see Publish Custom Metrics (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/publishingMetrics.html) in the Amazon CloudWatch User Guide.  * Choose a metric that changes proportionally with capacity. The value of the metric should increase or decrease in inverse proportion to the number of capacity units. That is, the value of the metric should decrease when capacity increases, and increase when capacity decreases.  For more information about CloudWatch, see Amazon CloudWatch Concepts (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_concepts.html).",
								MarkdownDescription: "Represents a CloudWatch metric of your choosing for a target tracking scaling policy to use with Application Auto Scaling.  For information about the available metrics for a service, see Amazon Web Services Services That Publish CloudWatch Metrics (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/aws-services-cloudwatch-metrics.html) in the Amazon CloudWatch User Guide.  To create your customized metric specification:  * Add values for each required parameter from CloudWatch. You can use an existing metric, or a new metric that you create. To use your own metric, you must first publish the metric to CloudWatch. For more information, see Publish Custom Metrics (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/publishingMetrics.html) in the Amazon CloudWatch User Guide.  * Choose a metric that changes proportionally with capacity. The value of the metric should increase or decrease in inverse proportion to the number of capacity units. That is, the value of the metric should decrease when capacity increases, and increase when capacity decreases.  For more information about CloudWatch, see Amazon CloudWatch Concepts (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_concepts.html).",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"dimensions": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"metric_name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"namespace": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"statistic": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"unit": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"disable_scale_in": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"predefined_metric_specification": {
								Description:         "Represents a predefined metric for a target tracking scaling policy to use with Application Auto Scaling.  Only the Amazon Web Services that you're using send metrics to Amazon CloudWatch. To determine whether a desired metric already exists by looking up its namespace and dimension using the CloudWatch metrics dashboard in the console, follow the procedure in Building dashboards with CloudWatch (https://docs.aws.amazon.com/autoscaling/application/userguide/monitoring-cloudwatch.html) in the Application Auto Scaling User Guide.",
								MarkdownDescription: "Represents a predefined metric for a target tracking scaling policy to use with Application Auto Scaling.  Only the Amazon Web Services that you're using send metrics to Amazon CloudWatch. To determine whether a desired metric already exists by looking up its namespace and dimension using the CloudWatch metrics dashboard in the console, follow the procedure in Building dashboards with CloudWatch (https://docs.aws.amazon.com/autoscaling/application/userguide/monitoring-cloudwatch.html) in the Application Auto Scaling User Guide.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"predefined_metric_type": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"resource_label": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"scale_in_cooldown": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"scale_out_cooldown": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"target_value": {
								Description:         "",
								MarkdownDescription: "",

								Type: utilities.DynamicNumberType{},

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *ApplicationautoscalingServicesK8SAwsScalingPolicyV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_applicationautoscaling_services_k8s_aws_scaling_policy_v1alpha1")

	var state ApplicationautoscalingServicesK8SAwsScalingPolicyV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ApplicationautoscalingServicesK8SAwsScalingPolicyV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("applicationautoscaling.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("ScalingPolicy")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *ApplicationautoscalingServicesK8SAwsScalingPolicyV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_applicationautoscaling_services_k8s_aws_scaling_policy_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *ApplicationautoscalingServicesK8SAwsScalingPolicyV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_applicationautoscaling_services_k8s_aws_scaling_policy_v1alpha1")

	var state ApplicationautoscalingServicesK8SAwsScalingPolicyV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ApplicationautoscalingServicesK8SAwsScalingPolicyV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("applicationautoscaling.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("ScalingPolicy")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *ApplicationautoscalingServicesK8SAwsScalingPolicyV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_applicationautoscaling_services_k8s_aws_scaling_policy_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
