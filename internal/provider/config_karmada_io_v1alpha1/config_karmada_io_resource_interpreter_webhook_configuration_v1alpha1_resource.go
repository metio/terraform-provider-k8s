/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package config_karmada_io_v1alpha1

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
)

var (
	_ resource.Resource                = &ConfigKarmadaIoResourceInterpreterWebhookConfigurationV1Alpha1Resource{}
	_ resource.ResourceWithConfigure   = &ConfigKarmadaIoResourceInterpreterWebhookConfigurationV1Alpha1Resource{}
	_ resource.ResourceWithImportState = &ConfigKarmadaIoResourceInterpreterWebhookConfigurationV1Alpha1Resource{}
)

func NewConfigKarmadaIoResourceInterpreterWebhookConfigurationV1Alpha1Resource() resource.Resource {
	return &ConfigKarmadaIoResourceInterpreterWebhookConfigurationV1Alpha1Resource{}
}

type ConfigKarmadaIoResourceInterpreterWebhookConfigurationV1Alpha1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type ConfigKarmadaIoResourceInterpreterWebhookConfigurationV1Alpha1ResourceData struct {
	ID             types.String `tfsdk:"id" json:"-"`
	ForceConflicts types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager   types.String `tfsdk:"field_manager" json:"-"`
	WaitFor        types.List   `tfsdk:"wait_for" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Webhooks *[]struct {
		ClientConfig *struct {
			CaBundle *string `tfsdk:"ca_bundle" json:"caBundle,omitempty"`
			Service  *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				Path      *string `tfsdk:"path" json:"path,omitempty"`
				Port      *int64  `tfsdk:"port" json:"port,omitempty"`
			} `tfsdk:"service" json:"service,omitempty"`
			Url *string `tfsdk:"url" json:"url,omitempty"`
		} `tfsdk:"client_config" json:"clientConfig,omitempty"`
		InterpreterContextVersions *[]string `tfsdk:"interpreter_context_versions" json:"interpreterContextVersions,omitempty"`
		Name                       *string   `tfsdk:"name" json:"name,omitempty"`
		Rules                      *[]struct {
			ApiGroups   *[]string `tfsdk:"api_groups" json:"apiGroups,omitempty"`
			ApiVersions *[]string `tfsdk:"api_versions" json:"apiVersions,omitempty"`
			Kinds       *[]string `tfsdk:"kinds" json:"kinds,omitempty"`
			Operations  *[]string `tfsdk:"operations" json:"operations,omitempty"`
		} `tfsdk:"rules" json:"rules,omitempty"`
		TimeoutSeconds *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
	} `tfsdk:"webhooks" json:"webhooks,omitempty"`
}

func (r *ConfigKarmadaIoResourceInterpreterWebhookConfigurationV1Alpha1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_config_karmada_io_resource_interpreter_webhook_configuration_v1alpha1"
}

func (r *ConfigKarmadaIoResourceInterpreterWebhookConfigurationV1Alpha1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ResourceInterpreterWebhookConfiguration describes the configuration of webhooks which take the responsibility to tell karmada the details of the resource object, especially for custom resources.",
		MarkdownDescription: "ResourceInterpreterWebhookConfiguration describes the configuration of webhooks which take the responsibility to tell karmada the details of the resource object, especially for custom resources.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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

			"wait_for": schema.ListNestedAttribute{
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

			"webhooks": schema.ListNestedAttribute{
				Description:         "Webhooks is a list of webhooks and the affected resources and operations.",
				MarkdownDescription: "Webhooks is a list of webhooks and the affected resources and operations.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"client_config": schema.SingleNestedAttribute{
							Description:         "ClientConfig defines how to communicate with the hook.",
							MarkdownDescription: "ClientConfig defines how to communicate with the hook.",
							Attributes: map[string]schema.Attribute{
								"ca_bundle": schema.StringAttribute{
									Description:         "'caBundle' is a PEM encoded CA bundle which will be used to validate the webhook's server certificate. If unspecified, system trust roots on the apiserver are used.",
									MarkdownDescription: "'caBundle' is a PEM encoded CA bundle which will be used to validate the webhook's server certificate. If unspecified, system trust roots on the apiserver are used.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										validators.Base64Validator(),
									},
								},

								"service": schema.SingleNestedAttribute{
									Description:         "'service' is a reference to the service for this webhook. Either 'service' or 'url' must be specified.  If the webhook is running within the cluster, then you should use 'service'.",
									MarkdownDescription: "'service' is a reference to the service for this webhook. Either 'service' or 'url' must be specified.  If the webhook is running within the cluster, then you should use 'service'.",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "'name' is the name of the service. Required",
											MarkdownDescription: "'name' is the name of the service. Required",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"namespace": schema.StringAttribute{
											Description:         "'namespace' is the namespace of the service. Required",
											MarkdownDescription: "'namespace' is the namespace of the service. Required",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"path": schema.StringAttribute{
											Description:         "'path' is an optional URL path which will be sent in any request to this service.",
											MarkdownDescription: "'path' is an optional URL path which will be sent in any request to this service.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"port": schema.Int64Attribute{
											Description:         "If specified, the port on the service that hosting webhook. Default to 443 for backward compatibility. 'port' should be a valid port number (1-65535, inclusive).",
											MarkdownDescription: "If specified, the port on the service that hosting webhook. Default to 443 for backward compatibility. 'port' should be a valid port number (1-65535, inclusive).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"url": schema.StringAttribute{
									Description:         "'url' gives the location of the webhook, in standard URL form ('scheme://host:port/path'). Exactly one of 'url' or 'service' must be specified.  The 'host' should not refer to a service running in the cluster; use the 'service' field instead. The host might be resolved via external DNS in some apiservers (e.g., 'kube-apiserver' cannot resolve in-cluster DNS as that would be a layering violation). 'host' may also be an IP address.  Please note that using 'localhost' or '127.0.0.1' as a 'host' is risky unless you take great care to run this webhook on all hosts which run an apiserver which might need to make calls to this webhook. Such installs are likely to be non-portable, i.e., not easy to turn up in a new cluster.  The scheme must be 'https'; the URL must begin with 'https://'.  A path is optional, and if present may be any string permissible in a URL. You may use the path to pass an arbitrary string to the webhook, for example, a cluster identifier.  Attempting to use a user or basic auth e.g. 'user:password@' is not allowed. Fragments ('#...') and query parameters ('?...') are not allowed, either.",
									MarkdownDescription: "'url' gives the location of the webhook, in standard URL form ('scheme://host:port/path'). Exactly one of 'url' or 'service' must be specified.  The 'host' should not refer to a service running in the cluster; use the 'service' field instead. The host might be resolved via external DNS in some apiservers (e.g., 'kube-apiserver' cannot resolve in-cluster DNS as that would be a layering violation). 'host' may also be an IP address.  Please note that using 'localhost' or '127.0.0.1' as a 'host' is risky unless you take great care to run this webhook on all hosts which run an apiserver which might need to make calls to this webhook. Such installs are likely to be non-portable, i.e., not easy to turn up in a new cluster.  The scheme must be 'https'; the URL must begin with 'https://'.  A path is optional, and if present may be any string permissible in a URL. You may use the path to pass an arbitrary string to the webhook, for example, a cluster identifier.  Attempting to use a user or basic auth e.g. 'user:password@' is not allowed. Fragments ('#...') and query parameters ('?...') are not allowed, either.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},
							},
							Required: true,
							Optional: false,
							Computed: false,
						},

						"interpreter_context_versions": schema.ListAttribute{
							Description:         "InterpreterContextVersions is an ordered list of preferred 'ResourceInterpreterContext' versions the Webhook expects. Karmada will try to use first version in the list which it supports. If none of the versions specified in this list supported by Karmada, validation will fail for this object. If a persisted webhook configuration specifies allowed versions and does not include any versions known to the Karmada, calls to the webhook will fail and be subject to the failure policy.",
							MarkdownDescription: "InterpreterContextVersions is an ordered list of preferred 'ResourceInterpreterContext' versions the Webhook expects. Karmada will try to use first version in the list which it supports. If none of the versions specified in this list supported by Karmada, validation will fail for this object. If a persisted webhook configuration specifies allowed versions and does not include any versions known to the Karmada, calls to the webhook will fail and be subject to the failure policy.",
							ElementType:         types.StringType,
							Required:            true,
							Optional:            false,
							Computed:            false,
						},

						"name": schema.StringAttribute{
							Description:         "Name is the full-qualified name of the webhook.",
							MarkdownDescription: "Name is the full-qualified name of the webhook.",
							Required:            true,
							Optional:            false,
							Computed:            false,
						},

						"rules": schema.ListNestedAttribute{
							Description:         "Rules describes what operations on what resources the webhook cares about. The webhook cares about an operation if it matches any Rule.",
							MarkdownDescription: "Rules describes what operations on what resources the webhook cares about. The webhook cares about an operation if it matches any Rule.",
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"api_groups": schema.ListAttribute{
										Description:         "APIGroups is the API groups the resources belong to. '*' is all groups. If '*' is present, the length of the slice must be one. For example: ['apps', 'batch', 'example.io'] means matches 3 groups. ['*'] means matches all group  Note: The group cloud be empty, e.g the 'core' group of kubernetes, in that case use [''].",
										MarkdownDescription: "APIGroups is the API groups the resources belong to. '*' is all groups. If '*' is present, the length of the slice must be one. For example: ['apps', 'batch', 'example.io'] means matches 3 groups. ['*'] means matches all group  Note: The group cloud be empty, e.g the 'core' group of kubernetes, in that case use [''].",
										ElementType:         types.StringType,
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"api_versions": schema.ListAttribute{
										Description:         "APIVersions is the API versions the resources belong to. '*' is all versions. If '*' is present, the length of the slice must be one. For example: ['v1alpha1', 'v1beta1'] means matches 2 versions. ['*'] means matches all versions.",
										MarkdownDescription: "APIVersions is the API versions the resources belong to. '*' is all versions. If '*' is present, the length of the slice must be one. For example: ['v1alpha1', 'v1beta1'] means matches 2 versions. ['*'] means matches all versions.",
										ElementType:         types.StringType,
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"kinds": schema.ListAttribute{
										Description:         "Kinds is a list of resources this rule applies to. If '*' is present, the length of the slice must be one. For example: ['Deployment', 'Pod'] means matches Deployment and Pod. ['*'] means apply to all resources.",
										MarkdownDescription: "Kinds is a list of resources this rule applies to. If '*' is present, the length of the slice must be one. For example: ['Deployment', 'Pod'] means matches Deployment and Pod. ['*'] means apply to all resources.",
										ElementType:         types.StringType,
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"operations": schema.ListAttribute{
										Description:         "Operations is the operations the hook cares about. If '*' is present, the length of the slice must be one.",
										MarkdownDescription: "Operations is the operations the hook cares about. If '*' is present, the length of the slice must be one.",
										ElementType:         types.StringType,
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
							},
							Required: false,
							Optional: true,
							Computed: false,
						},

						"timeout_seconds": schema.Int64Attribute{
							Description:         "TimeoutSeconds specifies the timeout for this webhook. After the timeout passes, the webhook call will be ignored or the API call will fail based on the failure policy. The timeout value must be between 1 and 30 seconds. Default to 10 seconds.",
							MarkdownDescription: "TimeoutSeconds specifies the timeout for this webhook. After the timeout passes, the webhook call will be ignored or the API call will fail based on the failure policy. The timeout value must be between 1 and 30 seconds. Default to 10 seconds.",
							Required:            false,
							Optional:            true,
							Computed:            false,
						},
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *ConfigKarmadaIoResourceInterpreterWebhookConfigurationV1Alpha1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *dynamic.DynamicClient, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *ConfigKarmadaIoResourceInterpreterWebhookConfigurationV1Alpha1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_config_karmada_io_resource_interpreter_webhook_configuration_v1alpha1")

	var model ConfigKarmadaIoResourceInterpreterWebhookConfigurationV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("config.karmada.io/v1alpha1")
	model.Kind = pointer.String("ResourceInterpreterWebhookConfiguration")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
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
		Resource(k8sSchema.GroupVersionResource{Group: "config.karmada.io", Version: "v1alpha1", Resource: "resourceinterpreterwebhookconfigurations"}).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while creating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse ConfigKarmadaIoResourceInterpreterWebhookConfigurationV1Alpha1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Webhooks = readResponse.Webhooks

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *ConfigKarmadaIoResourceInterpreterWebhookConfigurationV1Alpha1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_config_karmada_io_resource_interpreter_webhook_configuration_v1alpha1")

	var data ConfigKarmadaIoResourceInterpreterWebhookConfigurationV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "config.karmada.io", Version: "v1alpha1", Resource: "resourceinterpreterwebhookconfigurations"}).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse ConfigKarmadaIoResourceInterpreterWebhookConfigurationV1Alpha1ResourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.Metadata = readResponse.Metadata
	data.Webhooks = readResponse.Webhooks

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *ConfigKarmadaIoResourceInterpreterWebhookConfigurationV1Alpha1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_config_karmada_io_resource_interpreter_webhook_configuration_v1alpha1")

	var model ConfigKarmadaIoResourceInterpreterWebhookConfigurationV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("config.karmada.io/v1alpha1")
	model.Kind = pointer.String("ResourceInterpreterWebhookConfiguration")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
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
		Resource(k8sSchema.GroupVersionResource{Group: "config.karmada.io", Version: "v1alpha1", Resource: "resourceinterpreterwebhookconfigurations"}).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while updating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse ConfigKarmadaIoResourceInterpreterWebhookConfigurationV1Alpha1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Webhooks = readResponse.Webhooks

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *ConfigKarmadaIoResourceInterpreterWebhookConfigurationV1Alpha1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_config_karmada_io_resource_interpreter_webhook_configuration_v1alpha1")

	var data ConfigKarmadaIoResourceInterpreterWebhookConfigurationV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "config.karmada.io", Version: "v1alpha1", Resource: "resourceinterpreterwebhookconfigurations"}).
		Delete(ctx, data.Metadata.Name, meta.DeleteOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to DELETE resource",
			"An unexpected error occurred while deleting the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"DELETE Error: "+err.Error(),
		)
		return
	}
}

func (r *ConfigKarmadaIoResourceInterpreterWebhookConfigurationV1Alpha1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	if request.ID == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'name' Got: '%q'", request.ID),
		)
		return
	}
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	resource.ImportStatePassthroughID(ctx, path.Root("metadata").AtName("name"), request, response)
}
