/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package lambda_services_k8s_aws_v1alpha1

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
	_ datasource.DataSource = &LambdaServicesK8SAwsAliasV1Alpha1Manifest{}
)

func NewLambdaServicesK8SAwsAliasV1Alpha1Manifest() datasource.DataSource {
	return &LambdaServicesK8SAwsAliasV1Alpha1Manifest{}
}

type LambdaServicesK8SAwsAliasV1Alpha1Manifest struct{}

type LambdaServicesK8SAwsAliasV1Alpha1ManifestData struct {
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
		Description               *string `tfsdk:"description" json:"description,omitempty"`
		FunctionEventInvokeConfig *struct {
			DestinationConfig *struct {
				OnFailure *struct {
					Destination *string `tfsdk:"destination" json:"destination,omitempty"`
				} `tfsdk:"on_failure" json:"onFailure,omitempty"`
				OnSuccess *struct {
					Destination *string `tfsdk:"destination" json:"destination,omitempty"`
				} `tfsdk:"on_success" json:"onSuccess,omitempty"`
			} `tfsdk:"destination_config" json:"destinationConfig,omitempty"`
			FunctionName             *string `tfsdk:"function_name" json:"functionName,omitempty"`
			MaximumEventAgeInSeconds *int64  `tfsdk:"maximum_event_age_in_seconds" json:"maximumEventAgeInSeconds,omitempty"`
			MaximumRetryAttempts     *int64  `tfsdk:"maximum_retry_attempts" json:"maximumRetryAttempts,omitempty"`
			Qualifier                *string `tfsdk:"qualifier" json:"qualifier,omitempty"`
		} `tfsdk:"function_event_invoke_config" json:"functionEventInvokeConfig,omitempty"`
		FunctionName *string `tfsdk:"function_name" json:"functionName,omitempty"`
		FunctionRef  *struct {
			From *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"function_ref" json:"functionRef,omitempty"`
		FunctionVersion *string `tfsdk:"function_version" json:"functionVersion,omitempty"`
		Name            *string `tfsdk:"name" json:"name,omitempty"`
		Permissions     *[]struct {
			Action              *string `tfsdk:"action" json:"action,omitempty"`
			EventSourceToken    *string `tfsdk:"event_source_token" json:"eventSourceToken,omitempty"`
			FunctionURLAuthType *string `tfsdk:"function_url_auth_type" json:"functionURLAuthType,omitempty"`
			Principal           *string `tfsdk:"principal" json:"principal,omitempty"`
			PrincipalOrgID      *string `tfsdk:"principal_org_id" json:"principalOrgID,omitempty"`
			RevisionID          *string `tfsdk:"revision_id" json:"revisionID,omitempty"`
			SourceARN           *string `tfsdk:"source_arn" json:"sourceARN,omitempty"`
			SourceAccount       *string `tfsdk:"source_account" json:"sourceAccount,omitempty"`
			StatementID         *string `tfsdk:"statement_id" json:"statementID,omitempty"`
		} `tfsdk:"permissions" json:"permissions,omitempty"`
		ProvisionedConcurrencyConfig *struct {
			FunctionName                    *string `tfsdk:"function_name" json:"functionName,omitempty"`
			ProvisionedConcurrentExecutions *int64  `tfsdk:"provisioned_concurrent_executions" json:"provisionedConcurrentExecutions,omitempty"`
			Qualifier                       *string `tfsdk:"qualifier" json:"qualifier,omitempty"`
		} `tfsdk:"provisioned_concurrency_config" json:"provisionedConcurrencyConfig,omitempty"`
		RoutingConfig *struct {
			AdditionalVersionWeights *map[string]string `tfsdk:"additional_version_weights" json:"additionalVersionWeights,omitempty"`
		} `tfsdk:"routing_config" json:"routingConfig,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *LambdaServicesK8SAwsAliasV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_lambda_services_k8s_aws_alias_v1alpha1_manifest"
}

func (r *LambdaServicesK8SAwsAliasV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Alias is the Schema for the Aliases API",
		MarkdownDescription: "Alias is the Schema for the Aliases API",
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"description": schema.StringAttribute{
						Description:         "A description of the alias.",
						MarkdownDescription: "A description of the alias.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"function_event_invoke_config": schema.SingleNestedAttribute{
						Description:         "Configures options for asynchronous invocation on an alias. - DestinationConfig A destination for events after they have been sent to a function for processing. Types of Destinations: Function - The Amazon Resource Name (ARN) of a Lambda function. Queue - The ARN of a standard SQS queue. Topic - The ARN of a standard SNS topic. Event Bus - The ARN of an Amazon EventBridge event bus. - MaximumEventAgeInSeconds The maximum age of a request that Lambda sends to a function for processing. - MaximumRetryAttempts The maximum number of times to retry when the function returns an error.",
						MarkdownDescription: "Configures options for asynchronous invocation on an alias. - DestinationConfig A destination for events after they have been sent to a function for processing. Types of Destinations: Function - The Amazon Resource Name (ARN) of a Lambda function. Queue - The ARN of a standard SQS queue. Topic - The ARN of a standard SNS topic. Event Bus - The ARN of an Amazon EventBridge event bus. - MaximumEventAgeInSeconds The maximum age of a request that Lambda sends to a function for processing. - MaximumRetryAttempts The maximum number of times to retry when the function returns an error.",
						Attributes: map[string]schema.Attribute{
							"destination_config": schema.SingleNestedAttribute{
								Description:         "A configuration object that specifies the destination of an event after Lambda processes it. For more information, see Adding a destination (https://docs.aws.amazon.com/lambda/latest/dg/invocation-async-retain-records.html#invocation-async-destinations).",
								MarkdownDescription: "A configuration object that specifies the destination of an event after Lambda processes it. For more information, see Adding a destination (https://docs.aws.amazon.com/lambda/latest/dg/invocation-async-retain-records.html#invocation-async-destinations).",
								Attributes: map[string]schema.Attribute{
									"on_failure": schema.SingleNestedAttribute{
										Description:         "A destination for events that failed processing. For more information, see Adding a destination (https://docs.aws.amazon.com/lambda/latest/dg/invocation-async-retain-records.html#invocation-async-destinations).",
										MarkdownDescription: "A destination for events that failed processing. For more information, see Adding a destination (https://docs.aws.amazon.com/lambda/latest/dg/invocation-async-retain-records.html#invocation-async-destinations).",
										Attributes: map[string]schema.Attribute{
											"destination": schema.StringAttribute{
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

									"on_success": schema.SingleNestedAttribute{
										Description:         "A destination for events that were processed successfully. To retain records of successful asynchronous invocations (https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#invocation-async-destinations), you can configure an Amazon SNS topic, Amazon SQS queue, Lambda function, or Amazon EventBridge event bus as the destination. OnSuccess is not supported in CreateEventSourceMapping or UpdateEventSourceMapping requests.",
										MarkdownDescription: "A destination for events that were processed successfully. To retain records of successful asynchronous invocations (https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#invocation-async-destinations), you can configure an Amazon SNS topic, Amazon SQS queue, Lambda function, or Amazon EventBridge event bus as the destination. OnSuccess is not supported in CreateEventSourceMapping or UpdateEventSourceMapping requests.",
										Attributes: map[string]schema.Attribute{
											"destination": schema.StringAttribute{
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

							"function_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"maximum_event_age_in_seconds": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"maximum_retry_attempts": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"qualifier": schema.StringAttribute{
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

					"function_name": schema.StringAttribute{
						Description:         "The name or ARN of the Lambda function. Name formats * Function name - MyFunction. * Function ARN - arn:aws:lambda:us-west-2:123456789012:function:MyFunction. * Partial ARN - 123456789012:function:MyFunction. The length constraint applies only to the full ARN. If you specify only the function name, it is limited to 64 characters in length. Regex Pattern: '^(arn:(aws[a-zA-Z-]*)?:lambda:)?([a-z]{2}(-gov)?-[a-z]+-d{1}:)?(d{12}:)?(function:)?([a-zA-Z0-9-_]+)(:($LATEST|[a-zA-Z0-9-_]+))?$'",
						MarkdownDescription: "The name or ARN of the Lambda function. Name formats * Function name - MyFunction. * Function ARN - arn:aws:lambda:us-west-2:123456789012:function:MyFunction. * Partial ARN - 123456789012:function:MyFunction. The length constraint applies only to the full ARN. If you specify only the function name, it is limited to 64 characters in length. Regex Pattern: '^(arn:(aws[a-zA-Z-]*)?:lambda:)?([a-z]{2}(-gov)?-[a-z]+-d{1}:)?(d{12}:)?(function:)?([a-zA-Z0-9-_]+)(:($LATEST|[a-zA-Z0-9-_]+))?$'",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"function_ref": schema.SingleNestedAttribute{
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef: from: name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef: from: name: my-api",
						Attributes: map[string]schema.Attribute{
							"from": schema.SingleNestedAttribute{
								Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"function_version": schema.StringAttribute{
						Description:         "The function version that the alias invokes. Regex Pattern: '^($LATEST|[0-9]+)$'",
						MarkdownDescription: "The function version that the alias invokes. Regex Pattern: '^($LATEST|[0-9]+)$'",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"name": schema.StringAttribute{
						Description:         "The name of the alias. Regex Pattern: '^(?!^[0-9]+$)([a-zA-Z0-9-_]+)$'",
						MarkdownDescription: "The name of the alias. Regex Pattern: '^(?!^[0-9]+$)([a-zA-Z0-9-_]+)$'",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"permissions": schema.ListNestedAttribute{
						Description:         "Permissions configures a set of Lambda permissions to grant to an alias.",
						MarkdownDescription: "Permissions configures a set of Lambda permissions to grant to an alias.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"action": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"event_source_token": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"function_url_auth_type": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"principal": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"principal_org_id": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"revision_id": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"source_arn": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"source_account": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"statement_id": schema.StringAttribute{
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

					"provisioned_concurrency_config": schema.SingleNestedAttribute{
						Description:         "Configures provisioned concurrency to a function's alias - ProvisionedConcurrentExecutions The amount of provisioned concurrency to allocate for the version or alias. Minimum value of 1 is required",
						MarkdownDescription: "Configures provisioned concurrency to a function's alias - ProvisionedConcurrentExecutions The amount of provisioned concurrency to allocate for the version or alias. Minimum value of 1 is required",
						Attributes: map[string]schema.Attribute{
							"function_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"provisioned_concurrent_executions": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"qualifier": schema.StringAttribute{
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

					"routing_config": schema.SingleNestedAttribute{
						Description:         "The routing configuration (https://docs.aws.amazon.com/lambda/latest/dg/configuration-aliases.html#configuring-alias-routing) of the alias.",
						MarkdownDescription: "The routing configuration (https://docs.aws.amazon.com/lambda/latest/dg/configuration-aliases.html#configuring-alias-routing) of the alias.",
						Attributes: map[string]schema.Attribute{
							"additional_version_weights": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
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
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *LambdaServicesK8SAwsAliasV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_lambda_services_k8s_aws_alias_v1alpha1_manifest")

	var model LambdaServicesK8SAwsAliasV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("lambda.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("Alias")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
