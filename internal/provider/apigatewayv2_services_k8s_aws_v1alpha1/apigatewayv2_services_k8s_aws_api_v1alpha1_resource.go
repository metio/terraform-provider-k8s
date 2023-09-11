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
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	k8sTypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"strings"
	"time"
)

var (
	_ resource.Resource                = &Apigatewayv2ServicesK8SAwsApiV1Alpha1Resource{}
	_ resource.ResourceWithConfigure   = &Apigatewayv2ServicesK8SAwsApiV1Alpha1Resource{}
	_ resource.ResourceWithImportState = &Apigatewayv2ServicesK8SAwsApiV1Alpha1Resource{}
)

func NewApigatewayv2ServicesK8SAwsApiV1Alpha1Resource() resource.Resource {
	return &Apigatewayv2ServicesK8SAwsApiV1Alpha1Resource{}
}

type Apigatewayv2ServicesK8SAwsApiV1Alpha1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type Apigatewayv2ServicesK8SAwsApiV1Alpha1ResourceData struct {
	ID             types.String `tfsdk:"id" json:"-"`
	ForceConflicts types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager   types.String `tfsdk:"field_manager" json:"-"`
	WaitForUpsert  types.List   `tfsdk:"wait_for_upsert" json:"-"`
	WaitForDelete  types.Object `tfsdk:"wait_for_delete" json:"-"`

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

func (r *Apigatewayv2ServicesK8SAwsApiV1Alpha1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_apigatewayv2_services_k8s_aws_api_v1alpha1"
}

func (r *Apigatewayv2ServicesK8SAwsApiV1Alpha1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
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

			"force_conflicts": schema.BoolAttribute{
				Description:         "If 'true', server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "If `true`, server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"field_manager": schema.BoolAttribute{
				Description:         "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"wait_for_upsert": schema.ListNestedAttribute{
				Description:         "Wait for specific conditions after create/update of resources.",
				MarkdownDescription: "Wait for specific conditions after create/update of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"jsonpath": schema.StringAttribute{
							Description:         "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							MarkdownDescription: "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							Required:            true,
							Optional:            false,
							Computed:            false,
						},
						"value": schema.StringAttribute{
							Description:         "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							MarkdownDescription: "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							Required:            false,
							Optional:            true,
							Computed:            true,
						},
						"timeout": schema.StringAttribute{
							Description:         "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							MarkdownDescription: "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("30s"),
						},
						"poll_interval": schema.StringAttribute{
							Description:         "The length of time to wait before checking again.",
							MarkdownDescription: "The length of time to wait before checking again.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("5s"),
						},
					},
				},
			},

			"wait_for_delete": schema.SingleNestedAttribute{
				Description:         "Wait for deletion of resources.",
				MarkdownDescription: "Wait for deletion of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"timeout": schema.StringAttribute{
						Description:         "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
						MarkdownDescription: "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             stringdefault.StaticString("30s"),
					},
					"poll_interval": schema.StringAttribute{
						Description:         "The length of time to wait before checking again.",
						MarkdownDescription: "The length of time to wait before checking again.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             stringdefault.StaticString("5s"),
					},
				},
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},

					"labels": schema.MapAttribute{
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            true,
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
						Computed:            true,
						Validators: []validator.Map{
							validators.AnnotationValidator(),
						},
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "ApiSpec defines the desired state of Api.  Represents an API.",
				MarkdownDescription: "ApiSpec defines the desired state of Api.  Represents an API.",
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
						Description:         "Represents a CORS configuration. Supported only for HTTP APIs. See Configuring CORS (https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html) for more information.",
						MarkdownDescription: "Represents a CORS configuration. Supported only for HTTP APIs. See Configuring CORS (https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html) for more information.",
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

func (r *Apigatewayv2ServicesK8SAwsApiV1Alpha1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedResourceDataError(request.ProviderData))
	}
}

func (r *Apigatewayv2ServicesK8SAwsApiV1Alpha1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_apigatewayv2_services_k8s_aws_api_v1alpha1")

	var model Apigatewayv2ServicesK8SAwsApiV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("apigatewayv2.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("API")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonMarshalError(err))
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "apigatewayv2.services.k8s.aws", Version: "v1alpha1", Resource: "apis"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.Append(utilities.PatchError(err))
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse Apigatewayv2ServicesK8SAwsApiV1Alpha1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *Apigatewayv2ServicesK8SAwsApiV1Alpha1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_apigatewayv2_services_k8s_aws_api_v1alpha1")

	var data Apigatewayv2ServicesK8SAwsApiV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "apigatewayv2.services.k8s.aws", Version: "v1alpha1", Resource: "apis"}).
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

	var readResponse Apigatewayv2ServicesK8SAwsApiV1Alpha1ResourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *Apigatewayv2ServicesK8SAwsApiV1Alpha1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_apigatewayv2_services_k8s_aws_api_v1alpha1")

	var model Apigatewayv2ServicesK8SAwsApiV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("apigatewayv2.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("API")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonMarshalError(err))
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "apigatewayv2.services.k8s.aws", Version: "v1alpha1", Resource: "apis"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.Append(utilities.PatchError(err))
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse Apigatewayv2ServicesK8SAwsApiV1Alpha1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *Apigatewayv2ServicesK8SAwsApiV1Alpha1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_apigatewayv2_services_k8s_aws_api_v1alpha1")

	var data Apigatewayv2ServicesK8SAwsApiV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "apigatewayv2.services.k8s.aws", Version: "v1alpha1", Resource: "apis"}).
		Namespace(data.Metadata.Namespace).
		Delete(ctx, data.Metadata.Name, meta.DeleteOptions{})
	if utilities.IsDeletionError(err) {
		response.Diagnostics.Append(utilities.DeleteError(err))
		return
	}

	if !data.WaitForDelete.IsNull() {
		timeout := utilities.DetermineTimeout(data.WaitForDelete.Attributes())
		pollInterval := utilities.DeterminePollInterval(data.WaitForDelete.Attributes())

		startTime := time.Now()
		for {
			_, err := r.kubernetesClient.
				Resource(k8sSchema.GroupVersionResource{Group: "apigatewayv2.services.k8s.aws", Version: "v1alpha1", Resource: "apis"}).
				Namespace(data.Metadata.Namespace).
				Get(ctx, data.Metadata.Name, meta.GetOptions{})
			if utilities.IsNotFound(err) || timeout == time.Second*0 {
				break
			}
			if time.Now().After(startTime.Add(timeout)) {
				response.Diagnostics.Append(utilities.WaitTimeoutExceeded())
				return
			}
			time.Sleep(pollInterval)
		}
	}
}

func (r *Apigatewayv2ServicesK8SAwsApiV1Alpha1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	idParts := strings.Split(request.ID, "/")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'namespace/name' Got: '%q'", request.ID),
		)
		return
	}

	namespace := idParts[0]
	name := idParts[1]
	tflog.Trace(ctx, "parsed import ID", map[string]interface{}{
		"namespace": namespace,
		"name":      name,
	})
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("namespace"), namespace)...)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("name"), name)...)
}
