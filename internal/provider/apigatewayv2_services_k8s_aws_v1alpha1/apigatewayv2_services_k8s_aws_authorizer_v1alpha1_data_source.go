/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package apigatewayv2_services_k8s_aws_v1alpha1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &Apigatewayv2ServicesK8SAwsAuthorizerV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &Apigatewayv2ServicesK8SAwsAuthorizerV1Alpha1DataSource{}
)

func NewApigatewayv2ServicesK8SAwsAuthorizerV1Alpha1DataSource() datasource.DataSource {
	return &Apigatewayv2ServicesK8SAwsAuthorizerV1Alpha1DataSource{}
}

type Apigatewayv2ServicesK8SAwsAuthorizerV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type Apigatewayv2ServicesK8SAwsAuthorizerV1Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		ApiID  *string `tfsdk:"api_id" json:"apiID,omitempty"`
		ApiRef *struct {
			From *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"api_ref" json:"apiRef,omitempty"`
		AuthorizerCredentialsARN       *string   `tfsdk:"authorizer_credentials_arn" json:"authorizerCredentialsARN,omitempty"`
		AuthorizerPayloadFormatVersion *string   `tfsdk:"authorizer_payload_format_version" json:"authorizerPayloadFormatVersion,omitempty"`
		AuthorizerResultTtlInSeconds   *int64    `tfsdk:"authorizer_result_ttl_in_seconds" json:"authorizerResultTtlInSeconds,omitempty"`
		AuthorizerType                 *string   `tfsdk:"authorizer_type" json:"authorizerType,omitempty"`
		AuthorizerURI                  *string   `tfsdk:"authorizer_uri" json:"authorizerURI,omitempty"`
		EnableSimpleResponses          *bool     `tfsdk:"enable_simple_responses" json:"enableSimpleResponses,omitempty"`
		IdentitySource                 *[]string `tfsdk:"identity_source" json:"identitySource,omitempty"`
		IdentityValidationExpression   *string   `tfsdk:"identity_validation_expression" json:"identityValidationExpression,omitempty"`
		JwtConfiguration               *struct {
			Audience *[]string `tfsdk:"audience" json:"audience,omitempty"`
			Issuer   *string   `tfsdk:"issuer" json:"issuer,omitempty"`
		} `tfsdk:"jwt_configuration" json:"jwtConfiguration,omitempty"`
		Name *string `tfsdk:"name" json:"name,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *Apigatewayv2ServicesK8SAwsAuthorizerV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_apigatewayv2_services_k8s_aws_authorizer_v1alpha1"
}

func (r *Apigatewayv2ServicesK8SAwsAuthorizerV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Authorizer is the Schema for the Authorizers API",
		MarkdownDescription: "Authorizer is the Schema for the Authorizers API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"api_version": schema.StringAttribute{
				Description:         "The API group of the requested resource.",
				MarkdownDescription: "The API group of the requested resource.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"kind": schema.StringAttribute{
				Description:         "The type of the requested resource.",
				MarkdownDescription: "The type of the requested resource.",
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
						Optional:            false,
						Computed:            true,
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "AuthorizerSpec defines the desired state of Authorizer.  Represents an authorizer.",
				MarkdownDescription: "AuthorizerSpec defines the desired state of Authorizer.  Represents an authorizer.",
				Attributes: map[string]schema.Attribute{
					"api_id": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
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
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"authorizer_credentials_arn": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"authorizer_payload_format_version": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"authorizer_result_ttl_in_seconds": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"authorizer_type": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"authorizer_uri": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"enable_simple_responses": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"identity_source": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"identity_validation_expression": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"jwt_configuration": schema.SingleNestedAttribute{
						Description:         "Represents the configuration of a JWT authorizer. Required for the JWT authorizer type. Supported only for HTTP APIs.",
						MarkdownDescription: "Represents the configuration of a JWT authorizer. Required for the JWT authorizer type. Supported only for HTTP APIs.",
						Attributes: map[string]schema.Attribute{
							"audience": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"issuer": schema.StringAttribute{
								Description:         "A string representation of a URI with a length between [1-2048].",
								MarkdownDescription: "A string representation of a URI with a length between [1-2048].",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *Apigatewayv2ServicesK8SAwsAuthorizerV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedDataSourceDataError(request.ProviderData))
	}
}

func (r *Apigatewayv2ServicesK8SAwsAuthorizerV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_apigatewayv2_services_k8s_aws_authorizer_v1alpha1")

	var data Apigatewayv2ServicesK8SAwsAuthorizerV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "apigatewayv2.services.k8s.aws", Version: "v1alpha1", Resource: "authorizers"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetNamespacedResourceError(err, data.Metadata.Name, data.Metadata.Namespace))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse Apigatewayv2ServicesK8SAwsAuthorizerV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("apigatewayv2.services.k8s.aws/v1alpha1")
	data.Kind = pointer.String("Authorizer")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
