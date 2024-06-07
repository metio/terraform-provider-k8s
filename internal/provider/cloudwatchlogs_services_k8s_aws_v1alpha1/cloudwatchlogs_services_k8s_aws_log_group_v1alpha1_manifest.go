/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package cloudwatchlogs_services_k8s_aws_v1alpha1

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
	_ datasource.DataSource = &CloudwatchlogsServicesK8SAwsLogGroupV1Alpha1Manifest{}
)

func NewCloudwatchlogsServicesK8SAwsLogGroupV1Alpha1Manifest() datasource.DataSource {
	return &CloudwatchlogsServicesK8SAwsLogGroupV1Alpha1Manifest{}
}

type CloudwatchlogsServicesK8SAwsLogGroupV1Alpha1Manifest struct{}

type CloudwatchlogsServicesK8SAwsLogGroupV1Alpha1ManifestData struct {
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
		KmsKeyID  *string `tfsdk:"kms_key_id" json:"kmsKeyID,omitempty"`
		KmsKeyRef *struct {
			From *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"kms_key_ref" json:"kmsKeyRef,omitempty"`
		Name                *string `tfsdk:"name" json:"name,omitempty"`
		RetentionDays       *int64  `tfsdk:"retention_days" json:"retentionDays,omitempty"`
		SubscriptionFilters *[]struct {
			DestinationARN *string `tfsdk:"destination_arn" json:"destinationARN,omitempty"`
			Distribution   *string `tfsdk:"distribution" json:"distribution,omitempty"`
			FilterName     *string `tfsdk:"filter_name" json:"filterName,omitempty"`
			FilterPattern  *string `tfsdk:"filter_pattern" json:"filterPattern,omitempty"`
			RoleARN        *string `tfsdk:"role_arn" json:"roleARN,omitempty"`
		} `tfsdk:"subscription_filters" json:"subscriptionFilters,omitempty"`
		Tags *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CloudwatchlogsServicesK8SAwsLogGroupV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_cloudwatchlogs_services_k8s_aws_log_group_v1alpha1_manifest"
}

func (r *CloudwatchlogsServicesK8SAwsLogGroupV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "LogGroup is the Schema for the LogGroups API",
		MarkdownDescription: "LogGroup is the Schema for the LogGroups API",
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
				Description:         "LogGroupSpec defines the desired state of LogGroup.Represents a log group.",
				MarkdownDescription: "LogGroupSpec defines the desired state of LogGroup.Represents a log group.",
				Attributes: map[string]schema.Attribute{
					"kms_key_id": schema.StringAttribute{
						Description:         "The Amazon Resource Name (ARN) of the KMS key to use when encrypting logdata. For more information, see Amazon Resource Names (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-kms).",
						MarkdownDescription: "The Amazon Resource Name (ARN) of the KMS key to use when encrypting logdata. For more information, see Amazon Resource Names (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-kms).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"kms_key_ref": schema.SingleNestedAttribute{
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReferencetype to provide more user friendly syntax for references using 'from' fieldEx:APIIDRef:	from:	  name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReferencetype to provide more user friendly syntax for references using 'from' fieldEx:APIIDRef:	from:	  name: my-api",
						Attributes: map[string]schema.Attribute{
							"from": schema.SingleNestedAttribute{
								Description:         "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
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

					"name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"retention_days": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"subscription_filters": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"destination_arn": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"distribution": schema.StringAttribute{
									Description:         "The method used to distribute log data to the destination, which can be eitherrandom or grouped by log stream.",
									MarkdownDescription: "The method used to distribute log data to the destination, which can be eitherrandom or grouped by log stream.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"filter_name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"filter_pattern": schema.StringAttribute{
									Description:         "A symbolic description of how CloudWatch Logs should interpret the data ineach log event. For example, a log event can contain timestamps, IP addresses,strings, and so on. You use the filter pattern to specify what to look forin the log event message.",
									MarkdownDescription: "A symbolic description of how CloudWatch Logs should interpret the data ineach log event. For example, a log event can contain timestamps, IP addresses,strings, and so on. You use the filter pattern to specify what to look forin the log event message.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"role_arn": schema.StringAttribute{
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

					"tags": schema.MapAttribute{
						Description:         "The key-value pairs to use for the tags.You can grant users access to certain log groups while preventing them fromaccessing other log groups. To do so, tag your groups and use IAM policiesthat refer to those tags. To assign tags when you create a log group, youmust have either the logs:TagResource or logs:TagLogGroup permission. Formore information about tagging, see Tagging Amazon Web Services resources(https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html). For moreinformation about using tags to control access, see Controlling access toAmazon Web Services resources using tags (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_tags.html).",
						MarkdownDescription: "The key-value pairs to use for the tags.You can grant users access to certain log groups while preventing them fromaccessing other log groups. To do so, tag your groups and use IAM policiesthat refer to those tags. To assign tags when you create a log group, youmust have either the logs:TagResource or logs:TagLogGroup permission. Formore information about tagging, see Tagging Amazon Web Services resources(https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html). For moreinformation about using tags to control access, see Controlling access toAmazon Web Services resources using tags (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_tags.html).",
						ElementType:         types.StringType,
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

func (r *CloudwatchlogsServicesK8SAwsLogGroupV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_cloudwatchlogs_services_k8s_aws_log_group_v1alpha1_manifest")

	var model CloudwatchlogsServicesK8SAwsLogGroupV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("cloudwatchlogs.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("LogGroup")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
