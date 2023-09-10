/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package lambda_services_k8s_aws_v1alpha1

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
	_ datasource.DataSource = &LambdaServicesK8SAwsFunctionUrlconfigV1Alpha1Manifest{}
)

func NewLambdaServicesK8SAwsFunctionUrlconfigV1Alpha1Manifest() datasource.DataSource {
	return &LambdaServicesK8SAwsFunctionUrlconfigV1Alpha1Manifest{}
}

type LambdaServicesK8SAwsFunctionUrlconfigV1Alpha1Manifest struct{}

type LambdaServicesK8SAwsFunctionUrlconfigV1Alpha1ManifestData struct {
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
		AuthType *string `tfsdk:"auth_type" json:"authType,omitempty"`
		Cors     *struct {
			AllowCredentials *bool     `tfsdk:"allow_credentials" json:"allowCredentials,omitempty"`
			AllowHeaders     *[]string `tfsdk:"allow_headers" json:"allowHeaders,omitempty"`
			AllowMethods     *[]string `tfsdk:"allow_methods" json:"allowMethods,omitempty"`
			AllowOrigins     *[]string `tfsdk:"allow_origins" json:"allowOrigins,omitempty"`
			ExposeHeaders    *[]string `tfsdk:"expose_headers" json:"exposeHeaders,omitempty"`
			MaxAge           *int64    `tfsdk:"max_age" json:"maxAge,omitempty"`
		} `tfsdk:"cors" json:"cors,omitempty"`
		FunctionName *string `tfsdk:"function_name" json:"functionName,omitempty"`
		FunctionRef  *struct {
			From *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"function_ref" json:"functionRef,omitempty"`
		Qualifier *string `tfsdk:"qualifier" json:"qualifier,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *LambdaServicesK8SAwsFunctionUrlconfigV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_lambda_services_k8s_aws_function_url_config_v1alpha1_manifest"
}

func (r *LambdaServicesK8SAwsFunctionUrlconfigV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "FunctionURLConfig is the Schema for the FunctionURLConfigs API",
		MarkdownDescription: "FunctionURLConfig is the Schema for the FunctionURLConfigs API",
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
				Description:         "FunctionUrlConfigSpec defines the desired state of FunctionUrlConfig.  Details about a Lambda function URL.",
				MarkdownDescription: "FunctionUrlConfigSpec defines the desired state of FunctionUrlConfig.  Details about a Lambda function URL.",
				Attributes: map[string]schema.Attribute{
					"auth_type": schema.StringAttribute{
						Description:         "The type of authentication that your function URL uses. Set to AWS_IAM if you want to restrict access to authenticated IAM users only. Set to NONE if you want to bypass IAM authentication to create a public endpoint. For more information, see Security and auth model for Lambda function URLs (https://docs.aws.amazon.com/lambda/latest/dg/urls-auth.html).",
						MarkdownDescription: "The type of authentication that your function URL uses. Set to AWS_IAM if you want to restrict access to authenticated IAM users only. Set to NONE if you want to bypass IAM authentication to create a public endpoint. For more information, see Security and auth model for Lambda function URLs (https://docs.aws.amazon.com/lambda/latest/dg/urls-auth.html).",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"cors": schema.SingleNestedAttribute{
						Description:         "The cross-origin resource sharing (CORS) (https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) settings for your function URL.",
						MarkdownDescription: "The cross-origin resource sharing (CORS) (https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) settings for your function URL.",
						Attributes: map[string]schema.Attribute{
							"allow_credentials": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"allow_headers": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"allow_methods": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"allow_origins": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"expose_headers": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_age": schema.Int64Attribute{
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
						Description:         "The name of the Lambda function.  Name formats  * Function name – my-function.  * Function ARN – arn:aws:lambda:us-west-2:123456789012:function:my-function.  * Partial ARN – 123456789012:function:my-function.  The length constraint applies only to the full ARN. If you specify only the function name, it is limited to 64 characters in length.",
						MarkdownDescription: "The name of the Lambda function.  Name formats  * Function name – my-function.  * Function ARN – arn:aws:lambda:us-west-2:123456789012:function:my-function.  * Partial ARN – 123456789012:function:my-function.  The length constraint applies only to the full ARN. If you specify only the function name, it is limited to 64 characters in length.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"function_ref": schema.SingleNestedAttribute{
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef:  from: name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef:  from: name: my-api",
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

					"qualifier": schema.StringAttribute{
						Description:         "The alias name.",
						MarkdownDescription: "The alias name.",
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

func (r *LambdaServicesK8SAwsFunctionUrlconfigV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_lambda_services_k8s_aws_function_url_config_v1alpha1_manifest")

	var model LambdaServicesK8SAwsFunctionUrlconfigV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("lambda.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("FunctionURLConfig")

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
