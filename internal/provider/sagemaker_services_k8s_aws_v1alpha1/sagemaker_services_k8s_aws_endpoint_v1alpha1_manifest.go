/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package sagemaker_services_k8s_aws_v1alpha1

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
	_ datasource.DataSource = &SagemakerServicesK8SAwsEndpointV1Alpha1Manifest{}
)

func NewSagemakerServicesK8SAwsEndpointV1Alpha1Manifest() datasource.DataSource {
	return &SagemakerServicesK8SAwsEndpointV1Alpha1Manifest{}
}

type SagemakerServicesK8SAwsEndpointV1Alpha1Manifest struct{}

type SagemakerServicesK8SAwsEndpointV1Alpha1ManifestData struct {
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
		DeploymentConfig *struct {
			AutoRollbackConfiguration *struct {
				Alarms *[]struct {
					AlarmName *string `tfsdk:"alarm_name" json:"alarmName,omitempty"`
				} `tfsdk:"alarms" json:"alarms,omitempty"`
			} `tfsdk:"auto_rollback_configuration" json:"autoRollbackConfiguration,omitempty"`
			BlueGreenUpdatePolicy *struct {
				MaximumExecutionTimeoutInSeconds *int64 `tfsdk:"maximum_execution_timeout_in_seconds" json:"maximumExecutionTimeoutInSeconds,omitempty"`
				TerminationWaitInSeconds         *int64 `tfsdk:"termination_wait_in_seconds" json:"terminationWaitInSeconds,omitempty"`
				TrafficRoutingConfiguration      *struct {
					CanarySize *struct {
						Type_ *string `tfsdk:"type_" json:"type_,omitempty"`
						Value *int64  `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"canary_size" json:"canarySize,omitempty"`
					LinearStepSize *struct {
						Type_ *string `tfsdk:"type_" json:"type_,omitempty"`
						Value *int64  `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"linear_step_size" json:"linearStepSize,omitempty"`
					Type_                 *string `tfsdk:"type_" json:"type_,omitempty"`
					WaitIntervalInSeconds *int64  `tfsdk:"wait_interval_in_seconds" json:"waitIntervalInSeconds,omitempty"`
				} `tfsdk:"traffic_routing_configuration" json:"trafficRoutingConfiguration,omitempty"`
			} `tfsdk:"blue_green_update_policy" json:"blueGreenUpdatePolicy,omitempty"`
		} `tfsdk:"deployment_config" json:"deploymentConfig,omitempty"`
		EndpointConfigName *string `tfsdk:"endpoint_config_name" json:"endpointConfigName,omitempty"`
		EndpointName       *string `tfsdk:"endpoint_name" json:"endpointName,omitempty"`
		Tags               *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SagemakerServicesK8SAwsEndpointV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_sagemaker_services_k8s_aws_endpoint_v1alpha1_manifest"
}

func (r *SagemakerServicesK8SAwsEndpointV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Endpoint is the Schema for the Endpoints API",
		MarkdownDescription: "Endpoint is the Schema for the Endpoints API",
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
				Description:         "EndpointSpec defines the desired state of Endpoint.  A hosted endpoint for real-time inference.",
				MarkdownDescription: "EndpointSpec defines the desired state of Endpoint.  A hosted endpoint for real-time inference.",
				Attributes: map[string]schema.Attribute{
					"deployment_config": schema.SingleNestedAttribute{
						Description:         "The deployment configuration for an endpoint, which contains the desired deployment strategy and rollback configurations.",
						MarkdownDescription: "The deployment configuration for an endpoint, which contains the desired deployment strategy and rollback configurations.",
						Attributes: map[string]schema.Attribute{
							"auto_rollback_configuration": schema.SingleNestedAttribute{
								Description:         "Automatic rollback configuration for handling endpoint deployment failures and recovery.",
								MarkdownDescription: "Automatic rollback configuration for handling endpoint deployment failures and recovery.",
								Attributes: map[string]schema.Attribute{
									"alarms": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"alarm_name": schema.StringAttribute{
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

							"blue_green_update_policy": schema.SingleNestedAttribute{
								Description:         "Update policy for a blue/green deployment. If this update policy is specified, SageMaker creates a new fleet during the deployment while maintaining the old fleet. SageMaker flips traffic to the new fleet according to the specified traffic routing configuration. Only one update policy should be used in the deployment configuration. If no update policy is specified, SageMaker uses a blue/green deployment strategy with all at once traffic shifting by default.",
								MarkdownDescription: "Update policy for a blue/green deployment. If this update policy is specified, SageMaker creates a new fleet during the deployment while maintaining the old fleet. SageMaker flips traffic to the new fleet according to the specified traffic routing configuration. Only one update policy should be used in the deployment configuration. If no update policy is specified, SageMaker uses a blue/green deployment strategy with all at once traffic shifting by default.",
								Attributes: map[string]schema.Attribute{
									"maximum_execution_timeout_in_seconds": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"termination_wait_in_seconds": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"traffic_routing_configuration": schema.SingleNestedAttribute{
										Description:         "Defines the traffic routing strategy during an endpoint deployment to shift traffic from the old fleet to the new fleet.",
										MarkdownDescription: "Defines the traffic routing strategy during an endpoint deployment to shift traffic from the old fleet to the new fleet.",
										Attributes: map[string]schema.Attribute{
											"canary_size": schema.SingleNestedAttribute{
												Description:         "Specifies the endpoint capacity to activate for production.",
												MarkdownDescription: "Specifies the endpoint capacity to activate for production.",
												Attributes: map[string]schema.Attribute{
													"type_": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"value": schema.Int64Attribute{
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

											"linear_step_size": schema.SingleNestedAttribute{
												Description:         "Specifies the endpoint capacity to activate for production.",
												MarkdownDescription: "Specifies the endpoint capacity to activate for production.",
												Attributes: map[string]schema.Attribute{
													"type_": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"value": schema.Int64Attribute{
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

											"type_": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"wait_interval_in_seconds": schema.Int64Attribute{
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

					"endpoint_config_name": schema.StringAttribute{
						Description:         "The name of an endpoint configuration. For more information, see CreateEndpointConfig.",
						MarkdownDescription: "The name of an endpoint configuration. For more information, see CreateEndpointConfig.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"endpoint_name": schema.StringAttribute{
						Description:         "The name of the endpoint.The name must be unique within an Amazon Web Services Region in your Amazon Web Services account. The name is case-insensitive in CreateEndpoint, but the case is preserved and must be matched in .",
						MarkdownDescription: "The name of the endpoint.The name must be unique within an Amazon Web Services Region in your Amazon Web Services account. The name is case-insensitive in CreateEndpoint, but the case is preserved and must be matched in .",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"tags": schema.ListNestedAttribute{
						Description:         "An array of key-value pairs. You can use tags to categorize your Amazon Web Services resources in different ways, for example, by purpose, owner, or environment. For more information, see Tagging Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html).",
						MarkdownDescription: "An array of key-value pairs. You can use tags to categorize your Amazon Web Services resources in different ways, for example, by purpose, owner, or environment. For more information, see Tagging Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html).",
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

func (r *SagemakerServicesK8SAwsEndpointV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_sagemaker_services_k8s_aws_endpoint_v1alpha1_manifest")

	var model SagemakerServicesK8SAwsEndpointV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("sagemaker.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("Endpoint")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
