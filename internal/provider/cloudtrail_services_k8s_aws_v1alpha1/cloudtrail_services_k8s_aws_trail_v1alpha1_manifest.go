/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package cloudtrail_services_k8s_aws_v1alpha1

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
	_ datasource.DataSource = &CloudtrailServicesK8SAwsTrailV1Alpha1Manifest{}
)

func NewCloudtrailServicesK8SAwsTrailV1Alpha1Manifest() datasource.DataSource {
	return &CloudtrailServicesK8SAwsTrailV1Alpha1Manifest{}
}

type CloudtrailServicesK8SAwsTrailV1Alpha1Manifest struct{}

type CloudtrailServicesK8SAwsTrailV1Alpha1ManifestData struct {
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
		CloudWatchLogsLogGroupARN  *string `tfsdk:"cloud_watch_logs_log_group_arn" json:"cloudWatchLogsLogGroupARN,omitempty"`
		CloudWatchLogsRoleARN      *string `tfsdk:"cloud_watch_logs_role_arn" json:"cloudWatchLogsRoleARN,omitempty"`
		EnableLogFileValidation    *bool   `tfsdk:"enable_log_file_validation" json:"enableLogFileValidation,omitempty"`
		IncludeGlobalServiceEvents *bool   `tfsdk:"include_global_service_events" json:"includeGlobalServiceEvents,omitempty"`
		IsMultiRegionTrail         *bool   `tfsdk:"is_multi_region_trail" json:"isMultiRegionTrail,omitempty"`
		IsOrganizationTrail        *bool   `tfsdk:"is_organization_trail" json:"isOrganizationTrail,omitempty"`
		KmsKeyID                   *string `tfsdk:"kms_key_id" json:"kmsKeyID,omitempty"`
		Name                       *string `tfsdk:"name" json:"name,omitempty"`
		S3BucketName               *string `tfsdk:"s3_bucket_name" json:"s3BucketName,omitempty"`
		S3KeyPrefix                *string `tfsdk:"s3_key_prefix" json:"s3KeyPrefix,omitempty"`
		SnsTopicName               *string `tfsdk:"sns_topic_name" json:"snsTopicName,omitempty"`
		Tags                       *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CloudtrailServicesK8SAwsTrailV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_cloudtrail_services_k8s_aws_trail_v1alpha1_manifest"
}

func (r *CloudtrailServicesK8SAwsTrailV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Trail is the Schema for the Trails API",
		MarkdownDescription: "Trail is the Schema for the Trails API",
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
				Description:         "TrailSpec defines the desired state of Trail.The settings for a trail.",
				MarkdownDescription: "TrailSpec defines the desired state of Trail.The settings for a trail.",
				Attributes: map[string]schema.Attribute{
					"cloud_watch_logs_log_group_arn": schema.StringAttribute{
						Description:         "Specifies a log group name using an Amazon Resource Name (ARN), a uniqueidentifier that represents the log group to which CloudTrail logs will bedelivered. Not required unless you specify CloudWatchLogsRoleArn.",
						MarkdownDescription: "Specifies a log group name using an Amazon Resource Name (ARN), a uniqueidentifier that represents the log group to which CloudTrail logs will bedelivered. Not required unless you specify CloudWatchLogsRoleArn.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"cloud_watch_logs_role_arn": schema.StringAttribute{
						Description:         "Specifies the role for the CloudWatch Logs endpoint to assume to write toa user's log group.",
						MarkdownDescription: "Specifies the role for the CloudWatch Logs endpoint to assume to write toa user's log group.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_log_file_validation": schema.BoolAttribute{
						Description:         "Specifies whether log file integrity validation is enabled. The default isfalse.When you disable log file integrity validation, the chain of digest filesis broken after one hour. CloudTrail does not create digest files for logfiles that were delivered during a period in which log file integrity validationwas disabled. For example, if you enable log file integrity validation atnoon on January 1, disable it at noon on January 2, and re-enable it at noonon January 10, digest files will not be created for the log files deliveredfrom noon on January 2 to noon on January 10. The same applies whenever youstop CloudTrail logging or delete a trail.",
						MarkdownDescription: "Specifies whether log file integrity validation is enabled. The default isfalse.When you disable log file integrity validation, the chain of digest filesis broken after one hour. CloudTrail does not create digest files for logfiles that were delivered during a period in which log file integrity validationwas disabled. For example, if you enable log file integrity validation atnoon on January 1, disable it at noon on January 2, and re-enable it at noonon January 10, digest files will not be created for the log files deliveredfrom noon on January 2 to noon on January 10. The same applies whenever youstop CloudTrail logging or delete a trail.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"include_global_service_events": schema.BoolAttribute{
						Description:         "Specifies whether the trail is publishing events from global services suchas IAM to the log files.",
						MarkdownDescription: "Specifies whether the trail is publishing events from global services suchas IAM to the log files.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"is_multi_region_trail": schema.BoolAttribute{
						Description:         "Specifies whether the trail is created in the current region or in all regions.The default is false, which creates a trail only in the region where youare signed in. As a best practice, consider creating trails that log eventsin all regions.",
						MarkdownDescription: "Specifies whether the trail is created in the current region or in all regions.The default is false, which creates a trail only in the region where youare signed in. As a best practice, consider creating trails that log eventsin all regions.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"is_organization_trail": schema.BoolAttribute{
						Description:         "Specifies whether the trail is created for all accounts in an organizationin Organizations, or only for the current Amazon Web Services account. Thedefault is false, and cannot be true unless the call is made on behalf ofan Amazon Web Services account that is the management account for an organizationin Organizations.",
						MarkdownDescription: "Specifies whether the trail is created for all accounts in an organizationin Organizations, or only for the current Amazon Web Services account. Thedefault is false, and cannot be true unless the call is made on behalf ofan Amazon Web Services account that is the management account for an organizationin Organizations.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"kms_key_id": schema.StringAttribute{
						Description:         "Specifies the KMS key ID to use to encrypt the logs delivered by CloudTrail.The value can be an alias name prefixed by 'alias/', a fully specified ARNto an alias, a fully specified ARN to a key, or a globally unique identifier.CloudTrail also supports KMS multi-Region keys. For more information aboutmulti-Region keys, see Using multi-Region keys (https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html)in the Key Management Service Developer Guide.Examples:   * alias/MyAliasName   * arn:aws:kms:us-east-2:123456789012:alias/MyAliasName   * arn:aws:kms:us-east-2:123456789012:key/12345678-1234-1234-1234-123456789012   * 12345678-1234-1234-1234-123456789012",
						MarkdownDescription: "Specifies the KMS key ID to use to encrypt the logs delivered by CloudTrail.The value can be an alias name prefixed by 'alias/', a fully specified ARNto an alias, a fully specified ARN to a key, or a globally unique identifier.CloudTrail also supports KMS multi-Region keys. For more information aboutmulti-Region keys, see Using multi-Region keys (https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html)in the Key Management Service Developer Guide.Examples:   * alias/MyAliasName   * arn:aws:kms:us-east-2:123456789012:alias/MyAliasName   * arn:aws:kms:us-east-2:123456789012:key/12345678-1234-1234-1234-123456789012   * 12345678-1234-1234-1234-123456789012",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"name": schema.StringAttribute{
						Description:         "Specifies the name of the trail. The name must meet the following requirements:   * Contain only ASCII letters (a-z, A-Z), numbers (0-9), periods (.), underscores   (_), or dashes (-)   * Start with a letter or number, and end with a letter or number   * Be between 3 and 128 characters   * Have no adjacent periods, underscores or dashes. Names like my-_namespace   and my--namespace are not valid.   * Not be in IP address format (for example, 192.168.5.4)",
						MarkdownDescription: "Specifies the name of the trail. The name must meet the following requirements:   * Contain only ASCII letters (a-z, A-Z), numbers (0-9), periods (.), underscores   (_), or dashes (-)   * Start with a letter or number, and end with a letter or number   * Be between 3 and 128 characters   * Have no adjacent periods, underscores or dashes. Names like my-_namespace   and my--namespace are not valid.   * Not be in IP address format (for example, 192.168.5.4)",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"s3_bucket_name": schema.StringAttribute{
						Description:         "Specifies the name of the Amazon S3 bucket designated for publishing logfiles. See Amazon S3 Bucket Naming Requirements (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/create_trail_naming_policy.html).",
						MarkdownDescription: "Specifies the name of the Amazon S3 bucket designated for publishing logfiles. See Amazon S3 Bucket Naming Requirements (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/create_trail_naming_policy.html).",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"s3_key_prefix": schema.StringAttribute{
						Description:         "Specifies the Amazon S3 key prefix that comes after the name of the bucketyou have designated for log file delivery. For more information, see FindingYour CloudTrail Log Files (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-find-log-files.html).The maximum length is 200 characters.",
						MarkdownDescription: "Specifies the Amazon S3 key prefix that comes after the name of the bucketyou have designated for log file delivery. For more information, see FindingYour CloudTrail Log Files (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-find-log-files.html).The maximum length is 200 characters.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"sns_topic_name": schema.StringAttribute{
						Description:         "Specifies the name of the Amazon SNS topic defined for notification of logfile delivery. The maximum length is 256 characters.",
						MarkdownDescription: "Specifies the name of the Amazon SNS topic defined for notification of logfile delivery. The maximum length is 256 characters.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tags": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *CloudtrailServicesK8SAwsTrailV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_cloudtrail_services_k8s_aws_trail_v1alpha1_manifest")

	var model CloudtrailServicesK8SAwsTrailV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("cloudtrail.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("Trail")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
