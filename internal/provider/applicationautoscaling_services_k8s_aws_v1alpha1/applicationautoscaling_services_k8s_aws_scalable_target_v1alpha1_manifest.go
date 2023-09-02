/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package applicationautoscaling_services_k8s_aws_v1alpha1

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &ApplicationautoscalingServicesK8SAwsScalableTargetV1Alpha1Manifest{}
)

func NewApplicationautoscalingServicesK8SAwsScalableTargetV1Alpha1Manifest() datasource.DataSource {
	return &ApplicationautoscalingServicesK8SAwsScalableTargetV1Alpha1Manifest{}
}

type ApplicationautoscalingServicesK8SAwsScalableTargetV1Alpha1Manifest struct{}

type ApplicationautoscalingServicesK8SAwsScalableTargetV1Alpha1ManifestData struct {
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
		MaxCapacity       *int64  `tfsdk:"max_capacity" json:"maxCapacity,omitempty"`
		MinCapacity       *int64  `tfsdk:"min_capacity" json:"minCapacity,omitempty"`
		ResourceID        *string `tfsdk:"resource_id" json:"resourceID,omitempty"`
		RoleARN           *string `tfsdk:"role_arn" json:"roleARN,omitempty"`
		ScalableDimension *string `tfsdk:"scalable_dimension" json:"scalableDimension,omitempty"`
		ServiceNamespace  *string `tfsdk:"service_namespace" json:"serviceNamespace,omitempty"`
		SuspendedState    *struct {
			DynamicScalingInSuspended  *bool `tfsdk:"dynamic_scaling_in_suspended" json:"dynamicScalingInSuspended,omitempty"`
			DynamicScalingOutSuspended *bool `tfsdk:"dynamic_scaling_out_suspended" json:"dynamicScalingOutSuspended,omitempty"`
			ScheduledScalingSuspended  *bool `tfsdk:"scheduled_scaling_suspended" json:"scheduledScalingSuspended,omitempty"`
		} `tfsdk:"suspended_state" json:"suspendedState,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ApplicationautoscalingServicesK8SAwsScalableTargetV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_applicationautoscaling_services_k8s_aws_scalable_target_v1alpha1_manifest"
}

func (r *ApplicationautoscalingServicesK8SAwsScalableTargetV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ScalableTarget is the Schema for the ScalableTargets API",
		MarkdownDescription: "ScalableTarget is the Schema for the ScalableTargets API",
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
				Description:         "ScalableTargetSpec defines the desired state of ScalableTarget.  Represents a scalable target.",
				MarkdownDescription: "ScalableTargetSpec defines the desired state of ScalableTarget.  Represents a scalable target.",
				Attributes: map[string]schema.Attribute{
					"max_capacity": schema.Int64Attribute{
						Description:         "The maximum value that you plan to scale out to. When a scaling policy is in effect, Application Auto Scaling can scale out (expand) as needed to the maximum capacity limit in response to changing demand. This property is required when registering a new scalable target.  Although you can specify a large maximum capacity, note that service quotas may impose lower limits. Each service has its own default quotas for the maximum capacity of the resource. If you want to specify a higher limit, you can request an increase. For more information, consult the documentation for that service. For information about the default quotas for each service, see Service Endpoints and Quotas (https://docs.aws.amazon.com/general/latest/gr/aws-service-information.html) in the Amazon Web Services General Reference.",
						MarkdownDescription: "The maximum value that you plan to scale out to. When a scaling policy is in effect, Application Auto Scaling can scale out (expand) as needed to the maximum capacity limit in response to changing demand. This property is required when registering a new scalable target.  Although you can specify a large maximum capacity, note that service quotas may impose lower limits. Each service has its own default quotas for the maximum capacity of the resource. If you want to specify a higher limit, you can request an increase. For more information, consult the documentation for that service. For information about the default quotas for each service, see Service Endpoints and Quotas (https://docs.aws.amazon.com/general/latest/gr/aws-service-information.html) in the Amazon Web Services General Reference.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"min_capacity": schema.Int64Attribute{
						Description:         "The minimum value that you plan to scale in to. When a scaling policy is in effect, Application Auto Scaling can scale in (contract) as needed to the minimum capacity limit in response to changing demand. This property is required when registering a new scalable target.  For certain resources, the minimum value allowed is 0. This includes Lambda provisioned concurrency, Spot Fleet, ECS services, Aurora DB clusters, EMR clusters, and custom resources. For all other resources, the minimum value allowed is 1.",
						MarkdownDescription: "The minimum value that you plan to scale in to. When a scaling policy is in effect, Application Auto Scaling can scale in (contract) as needed to the minimum capacity limit in response to changing demand. This property is required when registering a new scalable target.  For certain resources, the minimum value allowed is 0. This includes Lambda provisioned concurrency, Spot Fleet, ECS services, Aurora DB clusters, EMR clusters, and custom resources. For all other resources, the minimum value allowed is 1.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"resource_id": schema.StringAttribute{
						Description:         "The identifier of the resource that is associated with the scalable target. This string consists of the resource type and unique identifier.  * ECS service - The resource type is service and the unique identifier is the cluster name and service name. Example: service/default/sample-webapp.  * Spot Fleet - The resource type is spot-fleet-request and the unique identifier is the Spot Fleet request ID. Example: spot-fleet-request/sfr-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE.  * EMR cluster - The resource type is instancegroup and the unique identifier is the cluster ID and instance group ID. Example: instancegroup/j-2EEZNYKUA1NTV/ig-1791Y4E1L8YI0.  * AppStream 2.0 fleet - The resource type is fleet and the unique identifier is the fleet name. Example: fleet/sample-fleet.  * DynamoDB table - The resource type is table and the unique identifier is the table name. Example: table/my-table.  * DynamoDB global secondary index - The resource type is index and the unique identifier is the index name. Example: table/my-table/index/my-table-index.  * Aurora DB cluster - The resource type is cluster and the unique identifier is the cluster name. Example: cluster:my-db-cluster.  * SageMaker endpoint variant - The resource type is variant and the unique identifier is the resource ID. Example: endpoint/my-end-point/variant/KMeansClustering.  * Custom resources are not supported with a resource type. This parameter must specify the OutputValue from the CloudFormation template stack used to access the resources. The unique identifier is defined by the service provider. More information is available in our GitHub repository (https://github.com/aws/aws-auto-scaling-custom-resource).  * Amazon Comprehend document classification endpoint - The resource type and unique identifier are specified using the endpoint ARN. Example: arn:aws:comprehend:us-west-2:123456789012:document-classifier-endpoint/EXAMPLE.  * Amazon Comprehend entity recognizer endpoint - The resource type and unique identifier are specified using the endpoint ARN. Example: arn:aws:comprehend:us-west-2:123456789012:entity-recognizer-endpoint/EXAMPLE.  * Lambda provisioned concurrency - The resource type is function and the unique identifier is the function name with a function version or alias name suffix that is not $LATEST. Example: function:my-function:prod or function:my-function:1.  * Amazon Keyspaces table - The resource type is table and the unique identifier is the table name. Example: keyspace/mykeyspace/table/mytable.  * Amazon MSK cluster - The resource type and unique identifier are specified using the cluster ARN. Example: arn:aws:kafka:us-east-1:123456789012:cluster/demo-cluster-1/6357e0b2-0e6a-4b86-a0b4-70df934c2e31-5.  * Amazon ElastiCache replication group - The resource type is replication-group and the unique identifier is the replication group name. Example: replication-group/mycluster.  * Neptune cluster - The resource type is cluster and the unique identifier is the cluster name. Example: cluster:mycluster.",
						MarkdownDescription: "The identifier of the resource that is associated with the scalable target. This string consists of the resource type and unique identifier.  * ECS service - The resource type is service and the unique identifier is the cluster name and service name. Example: service/default/sample-webapp.  * Spot Fleet - The resource type is spot-fleet-request and the unique identifier is the Spot Fleet request ID. Example: spot-fleet-request/sfr-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE.  * EMR cluster - The resource type is instancegroup and the unique identifier is the cluster ID and instance group ID. Example: instancegroup/j-2EEZNYKUA1NTV/ig-1791Y4E1L8YI0.  * AppStream 2.0 fleet - The resource type is fleet and the unique identifier is the fleet name. Example: fleet/sample-fleet.  * DynamoDB table - The resource type is table and the unique identifier is the table name. Example: table/my-table.  * DynamoDB global secondary index - The resource type is index and the unique identifier is the index name. Example: table/my-table/index/my-table-index.  * Aurora DB cluster - The resource type is cluster and the unique identifier is the cluster name. Example: cluster:my-db-cluster.  * SageMaker endpoint variant - The resource type is variant and the unique identifier is the resource ID. Example: endpoint/my-end-point/variant/KMeansClustering.  * Custom resources are not supported with a resource type. This parameter must specify the OutputValue from the CloudFormation template stack used to access the resources. The unique identifier is defined by the service provider. More information is available in our GitHub repository (https://github.com/aws/aws-auto-scaling-custom-resource).  * Amazon Comprehend document classification endpoint - The resource type and unique identifier are specified using the endpoint ARN. Example: arn:aws:comprehend:us-west-2:123456789012:document-classifier-endpoint/EXAMPLE.  * Amazon Comprehend entity recognizer endpoint - The resource type and unique identifier are specified using the endpoint ARN. Example: arn:aws:comprehend:us-west-2:123456789012:entity-recognizer-endpoint/EXAMPLE.  * Lambda provisioned concurrency - The resource type is function and the unique identifier is the function name with a function version or alias name suffix that is not $LATEST. Example: function:my-function:prod or function:my-function:1.  * Amazon Keyspaces table - The resource type is table and the unique identifier is the table name. Example: keyspace/mykeyspace/table/mytable.  * Amazon MSK cluster - The resource type and unique identifier are specified using the cluster ARN. Example: arn:aws:kafka:us-east-1:123456789012:cluster/demo-cluster-1/6357e0b2-0e6a-4b86-a0b4-70df934c2e31-5.  * Amazon ElastiCache replication group - The resource type is replication-group and the unique identifier is the replication group name. Example: replication-group/mycluster.  * Neptune cluster - The resource type is cluster and the unique identifier is the cluster name. Example: cluster:mycluster.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"role_arn": schema.StringAttribute{
						Description:         "This parameter is required for services that do not support service-linked roles (such as Amazon EMR), and it must specify the ARN of an IAM role that allows Application Auto Scaling to modify the scalable target on your behalf.  If the service supports service-linked roles, Application Auto Scaling uses a service-linked role, which it creates if it does not yet exist. For more information, see Application Auto Scaling IAM roles (https://docs.aws.amazon.com/autoscaling/application/userguide/security_iam_service-with-iam.html#security_iam_service-with-iam-roles).",
						MarkdownDescription: "This parameter is required for services that do not support service-linked roles (such as Amazon EMR), and it must specify the ARN of an IAM role that allows Application Auto Scaling to modify the scalable target on your behalf.  If the service supports service-linked roles, Application Auto Scaling uses a service-linked role, which it creates if it does not yet exist. For more information, see Application Auto Scaling IAM roles (https://docs.aws.amazon.com/autoscaling/application/userguide/security_iam_service-with-iam.html#security_iam_service-with-iam-roles).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"scalable_dimension": schema.StringAttribute{
						Description:         "The scalable dimension associated with the scalable target. This string consists of the service namespace, resource type, and scaling property.  * ecs:service:DesiredCount - The desired task count of an ECS service.  * elasticmapreduce:instancegroup:InstanceCount - The instance count of an EMR Instance Group.  * ec2:spot-fleet-request:TargetCapacity - The target capacity of a Spot Fleet.  * appstream:fleet:DesiredCapacity - The desired capacity of an AppStream 2.0 fleet.  * dynamodb:table:ReadCapacityUnits - The provisioned read capacity for a DynamoDB table.  * dynamodb:table:WriteCapacityUnits - The provisioned write capacity for a DynamoDB table.  * dynamodb:index:ReadCapacityUnits - The provisioned read capacity for a DynamoDB global secondary index.  * dynamodb:index:WriteCapacityUnits - The provisioned write capacity for a DynamoDB global secondary index.  * rds:cluster:ReadReplicaCount - The count of Aurora Replicas in an Aurora DB cluster. Available for Aurora MySQL-compatible edition and Aurora PostgreSQL-compatible edition.  * sagemaker:variant:DesiredInstanceCount - The number of EC2 instances for an SageMaker model endpoint variant.  * custom-resource:ResourceType:Property - The scalable dimension for a custom resource provided by your own application or service.  * comprehend:document-classifier-endpoint:DesiredInferenceUnits - The number of inference units for an Amazon Comprehend document classification endpoint.  * comprehend:entity-recognizer-endpoint:DesiredInferenceUnits - The number of inference units for an Amazon Comprehend entity recognizer endpoint.  * lambda:function:ProvisionedConcurrency - The provisioned concurrency for a Lambda function.  * cassandra:table:ReadCapacityUnits - The provisioned read capacity for an Amazon Keyspaces table.  * cassandra:table:WriteCapacityUnits - The provisioned write capacity for an Amazon Keyspaces table.  * kafka:broker-storage:VolumeSize - The provisioned volume size (in GiB) for brokers in an Amazon MSK cluster.  * elasticache:replication-group:NodeGroups - The number of node groups for an Amazon ElastiCache replication group.  * elasticache:replication-group:Replicas - The number of replicas per node group for an Amazon ElastiCache replication group.  * neptune:cluster:ReadReplicaCount - The count of read replicas in an Amazon Neptune DB cluster.",
						MarkdownDescription: "The scalable dimension associated with the scalable target. This string consists of the service namespace, resource type, and scaling property.  * ecs:service:DesiredCount - The desired task count of an ECS service.  * elasticmapreduce:instancegroup:InstanceCount - The instance count of an EMR Instance Group.  * ec2:spot-fleet-request:TargetCapacity - The target capacity of a Spot Fleet.  * appstream:fleet:DesiredCapacity - The desired capacity of an AppStream 2.0 fleet.  * dynamodb:table:ReadCapacityUnits - The provisioned read capacity for a DynamoDB table.  * dynamodb:table:WriteCapacityUnits - The provisioned write capacity for a DynamoDB table.  * dynamodb:index:ReadCapacityUnits - The provisioned read capacity for a DynamoDB global secondary index.  * dynamodb:index:WriteCapacityUnits - The provisioned write capacity for a DynamoDB global secondary index.  * rds:cluster:ReadReplicaCount - The count of Aurora Replicas in an Aurora DB cluster. Available for Aurora MySQL-compatible edition and Aurora PostgreSQL-compatible edition.  * sagemaker:variant:DesiredInstanceCount - The number of EC2 instances for an SageMaker model endpoint variant.  * custom-resource:ResourceType:Property - The scalable dimension for a custom resource provided by your own application or service.  * comprehend:document-classifier-endpoint:DesiredInferenceUnits - The number of inference units for an Amazon Comprehend document classification endpoint.  * comprehend:entity-recognizer-endpoint:DesiredInferenceUnits - The number of inference units for an Amazon Comprehend entity recognizer endpoint.  * lambda:function:ProvisionedConcurrency - The provisioned concurrency for a Lambda function.  * cassandra:table:ReadCapacityUnits - The provisioned read capacity for an Amazon Keyspaces table.  * cassandra:table:WriteCapacityUnits - The provisioned write capacity for an Amazon Keyspaces table.  * kafka:broker-storage:VolumeSize - The provisioned volume size (in GiB) for brokers in an Amazon MSK cluster.  * elasticache:replication-group:NodeGroups - The number of node groups for an Amazon ElastiCache replication group.  * elasticache:replication-group:Replicas - The number of replicas per node group for an Amazon ElastiCache replication group.  * neptune:cluster:ReadReplicaCount - The count of read replicas in an Amazon Neptune DB cluster.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"service_namespace": schema.StringAttribute{
						Description:         "The namespace of the Amazon Web Services service that provides the resource. For a resource provided by your own application or service, use custom-resource instead.",
						MarkdownDescription: "The namespace of the Amazon Web Services service that provides the resource. For a resource provided by your own application or service, use custom-resource instead.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"suspended_state": schema.SingleNestedAttribute{
						Description:         "An embedded object that contains attributes and attribute values that are used to suspend and resume automatic scaling. Setting the value of an attribute to true suspends the specified scaling activities. Setting it to false (default) resumes the specified scaling activities.  Suspension Outcomes  * For DynamicScalingInSuspended, while a suspension is in effect, all scale-in activities that are triggered by a scaling policy are suspended.  * For DynamicScalingOutSuspended, while a suspension is in effect, all scale-out activities that are triggered by a scaling policy are suspended.  * For ScheduledScalingSuspended, while a suspension is in effect, all scaling activities that involve scheduled actions are suspended.  For more information, see Suspending and resuming scaling (https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-suspend-resume-scaling.html) in the Application Auto Scaling User Guide.",
						MarkdownDescription: "An embedded object that contains attributes and attribute values that are used to suspend and resume automatic scaling. Setting the value of an attribute to true suspends the specified scaling activities. Setting it to false (default) resumes the specified scaling activities.  Suspension Outcomes  * For DynamicScalingInSuspended, while a suspension is in effect, all scale-in activities that are triggered by a scaling policy are suspended.  * For DynamicScalingOutSuspended, while a suspension is in effect, all scale-out activities that are triggered by a scaling policy are suspended.  * For ScheduledScalingSuspended, while a suspension is in effect, all scaling activities that involve scheduled actions are suspended.  For more information, see Suspending and resuming scaling (https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-suspend-resume-scaling.html) in the Application Auto Scaling User Guide.",
						Attributes: map[string]schema.Attribute{
							"dynamic_scaling_in_suspended": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"dynamic_scaling_out_suspended": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"scheduled_scaling_suspended": schema.BoolAttribute{
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
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *ApplicationautoscalingServicesK8SAwsScalableTargetV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_applicationautoscaling_services_k8s_aws_scalable_target_v1alpha1_manifest")

	var model ApplicationautoscalingServicesK8SAwsScalableTargetV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("applicationautoscaling.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("ScalableTarget")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"YAML Error: "+err.Error(),
		)
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
