/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package sns_services_k8s_aws_v1alpha1

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
	_ datasource.DataSource = &SnsServicesK8SAwsSubscriptionV1Alpha1Manifest{}
)

func NewSnsServicesK8SAwsSubscriptionV1Alpha1Manifest() datasource.DataSource {
	return &SnsServicesK8SAwsSubscriptionV1Alpha1Manifest{}
}

type SnsServicesK8SAwsSubscriptionV1Alpha1Manifest struct{}

type SnsServicesK8SAwsSubscriptionV1Alpha1ManifestData struct {
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
		DeliveryPolicy      *string `tfsdk:"delivery_policy" json:"deliveryPolicy,omitempty"`
		Endpoint            *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
		FilterPolicy        *string `tfsdk:"filter_policy" json:"filterPolicy,omitempty"`
		FilterPolicyScope   *string `tfsdk:"filter_policy_scope" json:"filterPolicyScope,omitempty"`
		Protocol            *string `tfsdk:"protocol" json:"protocol,omitempty"`
		RawMessageDelivery  *string `tfsdk:"raw_message_delivery" json:"rawMessageDelivery,omitempty"`
		RedrivePolicy       *string `tfsdk:"redrive_policy" json:"redrivePolicy,omitempty"`
		SubscriptionRoleARN *string `tfsdk:"subscription_role_arn" json:"subscriptionRoleARN,omitempty"`
		TopicARN            *string `tfsdk:"topic_arn" json:"topicARN,omitempty"`
		TopicRef            *struct {
			From *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"topic_ref" json:"topicRef,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SnsServicesK8SAwsSubscriptionV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_sns_services_k8s_aws_subscription_v1alpha1_manifest"
}

func (r *SnsServicesK8SAwsSubscriptionV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Subscription is the Schema for the Subscriptions API",
		MarkdownDescription: "Subscription is the Schema for the Subscriptions API",
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
				Description:         "SubscriptionSpec defines the desired state of Subscription.A wrapper type for the attributes of an Amazon SNS subscription.",
				MarkdownDescription: "SubscriptionSpec defines the desired state of Subscription.A wrapper type for the attributes of an Amazon SNS subscription.",
				Attributes: map[string]schema.Attribute{
					"delivery_policy": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"endpoint": schema.StringAttribute{
						Description:         "The endpoint that you want to receive notifications. Endpoints vary by protocol:   * For the http protocol, the (public) endpoint is a URL beginning with   http://.   * For the https protocol, the (public) endpoint is a URL beginning with   https://.   * For the email protocol, the endpoint is an email address.   * For the email-json protocol, the endpoint is an email address.   * For the sms protocol, the endpoint is a phone number of an SMS-enabled   device.   * For the sqs protocol, the endpoint is the ARN of an Amazon SQS queue.   * For the application protocol, the endpoint is the EndpointArn of a mobile   app and device.   * For the lambda protocol, the endpoint is the ARN of an Lambda function.   * For the firehose protocol, the endpoint is the ARN of an Amazon Kinesis   Data Firehose delivery stream.",
						MarkdownDescription: "The endpoint that you want to receive notifications. Endpoints vary by protocol:   * For the http protocol, the (public) endpoint is a URL beginning with   http://.   * For the https protocol, the (public) endpoint is a URL beginning with   https://.   * For the email protocol, the endpoint is an email address.   * For the email-json protocol, the endpoint is an email address.   * For the sms protocol, the endpoint is a phone number of an SMS-enabled   device.   * For the sqs protocol, the endpoint is the ARN of an Amazon SQS queue.   * For the application protocol, the endpoint is the EndpointArn of a mobile   app and device.   * For the lambda protocol, the endpoint is the ARN of an Lambda function.   * For the firehose protocol, the endpoint is the ARN of an Amazon Kinesis   Data Firehose delivery stream.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"filter_policy": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"filter_policy_scope": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"protocol": schema.StringAttribute{
						Description:         "The protocol that you want to use. Supported protocols include:   * http – delivery of JSON-encoded message via HTTP POST   * https – delivery of JSON-encoded message via HTTPS POST   * email – delivery of message via SMTP   * email-json – delivery of JSON-encoded message via SMTP   * sms – delivery of message via SMS   * sqs – delivery of JSON-encoded message to an Amazon SQS queue   * application – delivery of JSON-encoded message to an EndpointArn for   a mobile app and device   * lambda – delivery of JSON-encoded message to an Lambda function   * firehose – delivery of JSON-encoded message to an Amazon Kinesis Data   Firehose delivery stream.",
						MarkdownDescription: "The protocol that you want to use. Supported protocols include:   * http – delivery of JSON-encoded message via HTTP POST   * https – delivery of JSON-encoded message via HTTPS POST   * email – delivery of message via SMTP   * email-json – delivery of JSON-encoded message via SMTP   * sms – delivery of message via SMS   * sqs – delivery of JSON-encoded message to an Amazon SQS queue   * application – delivery of JSON-encoded message to an EndpointArn for   a mobile app and device   * lambda – delivery of JSON-encoded message to an Lambda function   * firehose – delivery of JSON-encoded message to an Amazon Kinesis Data   Firehose delivery stream.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"raw_message_delivery": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"redrive_policy": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"subscription_role_arn": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"topic_arn": schema.StringAttribute{
						Description:         "The ARN of the topic you want to subscribe to.",
						MarkdownDescription: "The ARN of the topic you want to subscribe to.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"topic_ref": schema.SingleNestedAttribute{
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *SnsServicesK8SAwsSubscriptionV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_sns_services_k8s_aws_subscription_v1alpha1_manifest")

	var model SnsServicesK8SAwsSubscriptionV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("sns.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("Subscription")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
