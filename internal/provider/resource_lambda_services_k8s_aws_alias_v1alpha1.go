/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"gopkg.in/yaml.v3"
	"time"
)

type LambdaServicesK8SAwsAliasV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*LambdaServicesK8SAwsAliasV1Alpha1Resource)(nil)
)

type LambdaServicesK8SAwsAliasV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type LambdaServicesK8SAwsAliasV1Alpha1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		Description *string `tfsdk:"description" yaml:"description,omitempty"`

		FunctionName *string `tfsdk:"function_name" yaml:"functionName,omitempty"`

		FunctionRef *struct {
			From *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"from" yaml:"from,omitempty"`
		} `tfsdk:"function_ref" yaml:"functionRef,omitempty"`

		FunctionVersion *string `tfsdk:"function_version" yaml:"functionVersion,omitempty"`

		Name *string `tfsdk:"name" yaml:"name,omitempty"`

		RoutingConfig *struct {
			AdditionalVersionWeights *map[string]string `tfsdk:"additional_version_weights" yaml:"additionalVersionWeights,omitempty"`
		} `tfsdk:"routing_config" yaml:"routingConfig,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewLambdaServicesK8SAwsAliasV1Alpha1Resource() resource.Resource {
	return &LambdaServicesK8SAwsAliasV1Alpha1Resource{}
}

func (r *LambdaServicesK8SAwsAliasV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lambda_services_k8s_aws_alias_v1alpha1"
}

func (r *LambdaServicesK8SAwsAliasV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Alias is the Schema for the Aliases API",
		MarkdownDescription: "Alias is the Schema for the Aliases API",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description:         "The timestamp of the last change to this resource.",
				MarkdownDescription: "The timestamp of the last change to this resource.",
				Type:                types.Int64Type,
				Computed:            true,
				Optional:            false,
			},

			"yaml": {
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"metadata": {
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Type:                types.StringType,
						Required:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
					},

					"namespace": {
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Type:                types.StringType,
						Optional:            true,
					},

					"labels": {
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.LabelValidator(),
						},
					},
					"annotations": {
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.AnnotationValidator(),
						},
					},
				}),
			},

			"api_version": {
				Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"kind": {
				Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"spec": {
				Description:         "",
				MarkdownDescription: "",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"description": {
						Description:         "A description of the alias.",
						MarkdownDescription: "A description of the alias.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"function_name": {
						Description:         "The name of the Lambda function.  Name formats  * Function name - MyFunction.  * Function ARN - arn:aws:lambda:us-west-2:123456789012:function:MyFunction.  * Partial ARN - 123456789012:function:MyFunction.  The length constraint applies only to the full ARN. If you specify only the function name, it is limited to 64 characters in length.",
						MarkdownDescription: "The name of the Lambda function.  Name formats  * Function name - MyFunction.  * Function ARN - arn:aws:lambda:us-west-2:123456789012:function:MyFunction.  * Partial ARN - 123456789012:function:MyFunction.  The length constraint applies only to the full ARN. If you specify only the function name, it is limited to 64 characters in length.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"function_ref": {
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef: from: name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef: from: name: my-api",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"from": {
								Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"function_version": {
						Description:         "The function version that the alias invokes.",
						MarkdownDescription: "The function version that the alias invokes.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"name": {
						Description:         "The name of the alias.",
						MarkdownDescription: "The name of the alias.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"routing_config": {
						Description:         "The routing configuration (https://docs.aws.amazon.com/lambda/latest/dg/configuration-aliases.html#configuring-alias-routing) of the alias.",
						MarkdownDescription: "The routing configuration (https://docs.aws.amazon.com/lambda/latest/dg/configuration-aliases.html#configuring-alias-routing) of the alias.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"additional_version_weights": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *LambdaServicesK8SAwsAliasV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_lambda_services_k8s_aws_alias_v1alpha1")

	var state LambdaServicesK8SAwsAliasV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel LambdaServicesK8SAwsAliasV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("lambda.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("Alias")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *LambdaServicesK8SAwsAliasV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_lambda_services_k8s_aws_alias_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *LambdaServicesK8SAwsAliasV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_lambda_services_k8s_aws_alias_v1alpha1")

	var state LambdaServicesK8SAwsAliasV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel LambdaServicesK8SAwsAliasV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("lambda.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("Alias")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *LambdaServicesK8SAwsAliasV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_lambda_services_k8s_aws_alias_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
