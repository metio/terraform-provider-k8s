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
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &Apigatewayv2ServicesK8SAwsStageV1Alpha1Manifest{}
)

func NewApigatewayv2ServicesK8SAwsStageV1Alpha1Manifest() datasource.DataSource {
	return &Apigatewayv2ServicesK8SAwsStageV1Alpha1Manifest{}
}

type Apigatewayv2ServicesK8SAwsStageV1Alpha1Manifest struct{}

type Apigatewayv2ServicesK8SAwsStageV1Alpha1ManifestData struct {
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
		AccessLogSettings *struct {
			DestinationARN *string `tfsdk:"destination_arn" json:"destinationARN,omitempty"`
			Format         *string `tfsdk:"format" json:"format,omitempty"`
		} `tfsdk:"access_log_settings" json:"accessLogSettings,omitempty"`
		ApiID  *string `tfsdk:"api_id" json:"apiID,omitempty"`
		ApiRef *struct {
			From *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"api_ref" json:"apiRef,omitempty"`
		AutoDeploy           *bool   `tfsdk:"auto_deploy" json:"autoDeploy,omitempty"`
		ClientCertificateID  *string `tfsdk:"client_certificate_id" json:"clientCertificateID,omitempty"`
		DefaultRouteSettings *struct {
			DataTraceEnabled       *bool    `tfsdk:"data_trace_enabled" json:"dataTraceEnabled,omitempty"`
			DetailedMetricsEnabled *bool    `tfsdk:"detailed_metrics_enabled" json:"detailedMetricsEnabled,omitempty"`
			LoggingLevel           *string  `tfsdk:"logging_level" json:"loggingLevel,omitempty"`
			ThrottlingBurstLimit   *int64   `tfsdk:"throttling_burst_limit" json:"throttlingBurstLimit,omitempty"`
			ThrottlingRateLimit    *float64 `tfsdk:"throttling_rate_limit" json:"throttlingRateLimit,omitempty"`
		} `tfsdk:"default_route_settings" json:"defaultRouteSettings,omitempty"`
		DeploymentID  *string `tfsdk:"deployment_id" json:"deploymentID,omitempty"`
		DeploymentRef *struct {
			From *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"deployment_ref" json:"deploymentRef,omitempty"`
		Description   *string `tfsdk:"description" json:"description,omitempty"`
		RouteSettings *struct {
			DataTraceEnabled       *bool    `tfsdk:"data_trace_enabled" json:"dataTraceEnabled,omitempty"`
			DetailedMetricsEnabled *bool    `tfsdk:"detailed_metrics_enabled" json:"detailedMetricsEnabled,omitempty"`
			LoggingLevel           *string  `tfsdk:"logging_level" json:"loggingLevel,omitempty"`
			ThrottlingBurstLimit   *int64   `tfsdk:"throttling_burst_limit" json:"throttlingBurstLimit,omitempty"`
			ThrottlingRateLimit    *float64 `tfsdk:"throttling_rate_limit" json:"throttlingRateLimit,omitempty"`
		} `tfsdk:"route_settings" json:"routeSettings,omitempty"`
		StageName      *string            `tfsdk:"stage_name" json:"stageName,omitempty"`
		StageVariables *map[string]string `tfsdk:"stage_variables" json:"stageVariables,omitempty"`
		Tags           *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *Apigatewayv2ServicesK8SAwsStageV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_apigatewayv2_services_k8s_aws_stage_v1alpha1_manifest"
}

func (r *Apigatewayv2ServicesK8SAwsStageV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Stage is the Schema for the Stages API",
		MarkdownDescription: "Stage is the Schema for the Stages API",
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
				Description:         "StageSpec defines the desired state of Stage.  Represents an API stage.",
				MarkdownDescription: "StageSpec defines the desired state of Stage.  Represents an API stage.",
				Attributes: map[string]schema.Attribute{
					"access_log_settings": schema.SingleNestedAttribute{
						Description:         "Settings for logging access in a stage.",
						MarkdownDescription: "Settings for logging access in a stage.",
						Attributes: map[string]schema.Attribute{
							"destination_arn": schema.StringAttribute{
								Description:         "Represents an Amazon Resource Name (ARN).",
								MarkdownDescription: "Represents an Amazon Resource Name (ARN).",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"format": schema.StringAttribute{
								Description:         "A string with a length between [1-1024].",
								MarkdownDescription: "A string with a length between [1-1024].",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"api_id": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"api_ref": schema.SingleNestedAttribute{
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

					"auto_deploy": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"client_certificate_id": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"default_route_settings": schema.SingleNestedAttribute{
						Description:         "Represents a collection of route settings.",
						MarkdownDescription: "Represents a collection of route settings.",
						Attributes: map[string]schema.Attribute{
							"data_trace_enabled": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"detailed_metrics_enabled": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"logging_level": schema.StringAttribute{
								Description:         "The logging level.",
								MarkdownDescription: "The logging level.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"throttling_burst_limit": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"throttling_rate_limit": schema.Float64Attribute{
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

					"deployment_id": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"deployment_ref": schema.SingleNestedAttribute{
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

					"description": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"route_settings": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"data_trace_enabled": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"detailed_metrics_enabled": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"logging_level": schema.StringAttribute{
								Description:         "The logging level.",
								MarkdownDescription: "The logging level.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"throttling_burst_limit": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"throttling_rate_limit": schema.Float64Attribute{
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

					"stage_name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"stage_variables": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *Apigatewayv2ServicesK8SAwsStageV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_apigatewayv2_services_k8s_aws_stage_v1alpha1_manifest")

	var model Apigatewayv2ServicesK8SAwsStageV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("apigatewayv2.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("Stage")

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
