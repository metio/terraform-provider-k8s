/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package apigatewayv2_services_k8s_aws_v1alpha1

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
	_ datasource.DataSource = &Apigatewayv2ServicesK8SAwsApiV1Alpha1Manifest{}
)

func NewApigatewayv2ServicesK8SAwsApiV1Alpha1Manifest() datasource.DataSource {
	return &Apigatewayv2ServicesK8SAwsApiV1Alpha1Manifest{}
}

type Apigatewayv2ServicesK8SAwsApiV1Alpha1Manifest struct{}

type Apigatewayv2ServicesK8SAwsApiV1Alpha1ManifestData struct {
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
				Description:         "ApiSpec defines the desired state of Api.Represents an API.",
				MarkdownDescription: "ApiSpec defines the desired state of Api.Represents an API.",
				Attributes: map[string]schema.Attribute{
					"api_key_selection_expression": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"basepath": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"body": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"cors_configuration": schema.SingleNestedAttribute{
						Description:         "Represents a CORS configuration. Supported only for HTTP APIs. See ConfiguringCORS (https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html)for more information.",
						MarkdownDescription: "Represents a CORS configuration. Supported only for HTTP APIs. See ConfiguringCORS (https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html)for more information.",
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
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"description": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"disable_execute_api_endpoint": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"disable_schema_validation": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"fail_on_warnings": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"protocol_type": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"route_key": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"route_selection_expression": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tags": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"target": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"version": schema.StringAttribute{
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
	}
}

func (r *Apigatewayv2ServicesK8SAwsApiV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_apigatewayv2_services_k8s_aws_api_v1alpha1_manifest")

	var model Apigatewayv2ServicesK8SAwsApiV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
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
