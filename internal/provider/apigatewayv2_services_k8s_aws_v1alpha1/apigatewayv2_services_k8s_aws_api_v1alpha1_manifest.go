/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package apigatewayv2_services_k8s_aws_v1alpha1

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
	_ datasource.DataSource = &Apigatewayv2ServicesK8SAwsApiV1Alpha1Manifest{}
)

func NewApigatewayv2ServicesK8SAwsApiV1Alpha1Manifest() datasource.DataSource {
	return &Apigatewayv2ServicesK8SAwsApiV1Alpha1Manifest{}
}

type Apigatewayv2ServicesK8SAwsApiV1Alpha1Manifest struct{}

type Apigatewayv2ServicesK8SAwsApiV1Alpha1ManifestData struct {
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
		ApiKeySelectionExpression *string `tfsdk:"api_key_selection_expression" json:"apiKeySelectionExpression,omitempty"`
		Basepath                  *string `tfsdk:"basepath" json:"basepath,omitempty"`
		Body                      *string `tfsdk:"body" json:"body,omitempty"`
		CorsConfiguration         *struct {
			AllowCredentials *bool     `tfsdk:"allow_credentials" json:"allowCredentials,omitempty"`
			AllowHeaders     *[]string `tfsdk:"allow_headers" json:"allowHeaders,omitempty"`
			AllowMethods     *[]string `tfsdk:"allow_methods" json:"allowMethods,omitempty"`
			AllowOrigins     *[]string `tfsdk:"allow_origins" json:"allowOrigins,omitempty"`
			ExposeHeaders    *[]string `tfsdk:"expose_headers" json:"exposeHeaders,omitempty"`
			MaxAge           *int64    `tfsdk:"max_age" json:"maxAge,omitempty"`
		} `tfsdk:"cors_configuration" json:"corsConfiguration,omitempty"`
		CredentialsARN            *string            `tfsdk:"credentials_arn" json:"credentialsARN,omitempty"`
		Description               *string            `tfsdk:"description" json:"description,omitempty"`
		DisableExecuteAPIEndpoint *bool              `tfsdk:"disable_execute_api_endpoint" json:"disableExecuteAPIEndpoint,omitempty"`
		DisableSchemaValidation   *bool              `tfsdk:"disable_schema_validation" json:"disableSchemaValidation,omitempty"`
		FailOnWarnings            *bool              `tfsdk:"fail_on_warnings" json:"failOnWarnings,omitempty"`
		Name                      *string            `tfsdk:"name" json:"name,omitempty"`
		ProtocolType              *string            `tfsdk:"protocol_type" json:"protocolType,omitempty"`
		RouteKey                  *string            `tfsdk:"route_key" json:"routeKey,omitempty"`
		RouteSelectionExpression  *string            `tfsdk:"route_selection_expression" json:"routeSelectionExpression,omitempty"`
		Tags                      *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
		Target                    *string            `tfsdk:"target" json:"target,omitempty"`
		Version                   *string            `tfsdk:"version" json:"version,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *Apigatewayv2ServicesK8SAwsApiV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_apigatewayv2_services_k8s_aws_api_v1alpha1_manifest"
}

func (r *Apigatewayv2ServicesK8SAwsApiV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "API is the Schema for the APIS API",
		MarkdownDescription: "API is the Schema for the APIS API",
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
				Description:         "ApiSpec defines the desired state of Api. Represents an API.",
				MarkdownDescription: "ApiSpec defines the desired state of Api. Represents an API.",
				Attributes: map[string]schema.Attribute{
					"api_key_selection_expression": schema.StringAttribute{
						Description:         "An API key selection expression. Supported only for WebSocket APIs. See API Key Selection Expressions (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).",
						MarkdownDescription: "An API key selection expression. Supported only for WebSocket APIs. See API Key Selection Expressions (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"basepath": schema.StringAttribute{
						Description:         "Specifies how to interpret the base path of the API during import. Valid values are ignore, prepend, and split. The default value is ignore. To learn more, see Set the OpenAPI basePath Property (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-import-api-basePath.html). Supported only for HTTP APIs.",
						MarkdownDescription: "Specifies how to interpret the base path of the API during import. Valid values are ignore, prepend, and split. The default value is ignore. To learn more, see Set the OpenAPI basePath Property (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-import-api-basePath.html). Supported only for HTTP APIs.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"body": schema.StringAttribute{
						Description:         "The OpenAPI definition. Supported only for HTTP APIs.",
						MarkdownDescription: "The OpenAPI definition. Supported only for HTTP APIs.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"cors_configuration": schema.SingleNestedAttribute{
						Description:         "A CORS configuration. Supported only for HTTP APIs. See Configuring CORS (https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html) for more information.",
						MarkdownDescription: "A CORS configuration. Supported only for HTTP APIs. See Configuring CORS (https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html) for more information.",
						Attributes: map[string]schema.Attribute{
							"allow_credentials": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"allow_headers": schema.ListAttribute{
								Description:         "Represents a collection of allowed headers. Supported only for HTTP APIs.",
								MarkdownDescription: "Represents a collection of allowed headers. Supported only for HTTP APIs.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"allow_methods": schema.ListAttribute{
								Description:         "Represents a collection of methods. Supported only for HTTP APIs.",
								MarkdownDescription: "Represents a collection of methods. Supported only for HTTP APIs.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"allow_origins": schema.ListAttribute{
								Description:         "Represents a collection of origins. Supported only for HTTP APIs.",
								MarkdownDescription: "Represents a collection of origins. Supported only for HTTP APIs.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"expose_headers": schema.ListAttribute{
								Description:         "Represents a collection of allowed headers. Supported only for HTTP APIs.",
								MarkdownDescription: "Represents a collection of allowed headers. Supported only for HTTP APIs.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_age": schema.Int64Attribute{
								Description:         "An integer with a value between -1 and 86400. Supported only for HTTP APIs.",
								MarkdownDescription: "An integer with a value between -1 and 86400. Supported only for HTTP APIs.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"credentials_arn": schema.StringAttribute{
						Description:         "This property is part of quick create. It specifies the credentials required for the integration, if any. For a Lambda integration, three options are available. To specify an IAM Role for API Gateway to assume, use the role's Amazon Resource Name (ARN). To require that the caller's identity be passed through from the request, specify arn:aws:iam::*:user/*. To use resource-based permissions on supported AWS services, specify null. Currently, this property is not used for HTTP integrations. Supported only for HTTP APIs.",
						MarkdownDescription: "This property is part of quick create. It specifies the credentials required for the integration, if any. For a Lambda integration, three options are available. To specify an IAM Role for API Gateway to assume, use the role's Amazon Resource Name (ARN). To require that the caller's identity be passed through from the request, specify arn:aws:iam::*:user/*. To use resource-based permissions on supported AWS services, specify null. Currently, this property is not used for HTTP integrations. Supported only for HTTP APIs.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"description": schema.StringAttribute{
						Description:         "The description of the API.",
						MarkdownDescription: "The description of the API.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"disable_execute_api_endpoint": schema.BoolAttribute{
						Description:         "Specifies whether clients can invoke your API by using the default execute-api endpoint. By default, clients can invoke your API with the default https://{api_id}.execute-api.{region}.amazonaws.com endpoint. To require that clients use a custom domain name to invoke your API, disable the default endpoint.",
						MarkdownDescription: "Specifies whether clients can invoke your API by using the default execute-api endpoint. By default, clients can invoke your API with the default https://{api_id}.execute-api.{region}.amazonaws.com endpoint. To require that clients use a custom domain name to invoke your API, disable the default endpoint.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"disable_schema_validation": schema.BoolAttribute{
						Description:         "Avoid validating models when creating a deployment. Supported only for WebSocket APIs.",
						MarkdownDescription: "Avoid validating models when creating a deployment. Supported only for WebSocket APIs.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"fail_on_warnings": schema.BoolAttribute{
						Description:         "Specifies whether to rollback the API creation when a warning is encountered. By default, API creation continues if a warning is encountered.",
						MarkdownDescription: "Specifies whether to rollback the API creation when a warning is encountered. By default, API creation continues if a warning is encountered.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"name": schema.StringAttribute{
						Description:         "The name of the API.",
						MarkdownDescription: "The name of the API.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"protocol_type": schema.StringAttribute{
						Description:         "The API protocol.",
						MarkdownDescription: "The API protocol.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"route_key": schema.StringAttribute{
						Description:         "This property is part of quick create. If you don't specify a routeKey, a default route of $default is created. The $default route acts as a catch-all for any request made to your API, for a particular stage. The $default route key can't be modified. You can add routes after creating the API, and you can update the route keys of additional routes. Supported only for HTTP APIs.",
						MarkdownDescription: "This property is part of quick create. If you don't specify a routeKey, a default route of $default is created. The $default route acts as a catch-all for any request made to your API, for a particular stage. The $default route key can't be modified. You can add routes after creating the API, and you can update the route keys of additional routes. Supported only for HTTP APIs.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"route_selection_expression": schema.StringAttribute{
						Description:         "The route selection expression for the API. For HTTP APIs, the routeSelectionExpression must be ${request.method} ${request.path}. If not provided, this will be the default for HTTP APIs. This property is required for WebSocket APIs.",
						MarkdownDescription: "The route selection expression for the API. For HTTP APIs, the routeSelectionExpression must be ${request.method} ${request.path}. If not provided, this will be the default for HTTP APIs. This property is required for WebSocket APIs.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tags": schema.MapAttribute{
						Description:         "The collection of tags. Each tag element is associated with a given resource.",
						MarkdownDescription: "The collection of tags. Each tag element is associated with a given resource.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"target": schema.StringAttribute{
						Description:         "This property is part of quick create. Quick create produces an API with an integration, a default catch-all route, and a default stage which is configured to automatically deploy changes. For HTTP integrations, specify a fully qualified URL. For Lambda integrations, specify a function ARN. The type of the integration will be HTTP_PROXY or AWS_PROXY, respectively. Supported only for HTTP APIs.",
						MarkdownDescription: "This property is part of quick create. Quick create produces an API with an integration, a default catch-all route, and a default stage which is configured to automatically deploy changes. For HTTP integrations, specify a fully qualified URL. For Lambda integrations, specify a function ARN. The type of the integration will be HTTP_PROXY or AWS_PROXY, respectively. Supported only for HTTP APIs.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"version": schema.StringAttribute{
						Description:         "A version identifier for the API.",
						MarkdownDescription: "A version identifier for the API.",
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

func (r *Apigatewayv2ServicesK8SAwsApiV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_apigatewayv2_services_k8s_aws_api_v1alpha1_manifest")

	var model Apigatewayv2ServicesK8SAwsApiV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("apigatewayv2.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("API")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
