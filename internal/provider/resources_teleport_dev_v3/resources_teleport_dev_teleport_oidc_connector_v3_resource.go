/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package resources_teleport_dev_v3

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
	_ resource.Resource                = &ResourcesTeleportDevTeleportOidcconnectorV3Resource{}
	_ resource.ResourceWithConfigure   = &ResourcesTeleportDevTeleportOidcconnectorV3Resource{}
	_ resource.ResourceWithImportState = &ResourcesTeleportDevTeleportOidcconnectorV3Resource{}
)

func NewResourcesTeleportDevTeleportOidcconnectorV3Resource() resource.Resource {
	return &ResourcesTeleportDevTeleportOidcconnectorV3Resource{}
}

type ResourcesTeleportDevTeleportOidcconnectorV3Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type ResourcesTeleportDevTeleportOidcconnectorV3ResourceData struct {
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
		Acr_values             *string `tfsdk:"acr_values" json:"acr_values,omitempty"`
		Allow_unverified_email *bool   `tfsdk:"allow_unverified_email" json:"allow_unverified_email,omitempty"`
		Claims_to_roles        *[]struct {
			Claim *string   `tfsdk:"claim" json:"claim,omitempty"`
			Roles *[]string `tfsdk:"roles" json:"roles,omitempty"`
			Value *string   `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"claims_to_roles" json:"claims_to_roles,omitempty"`
		Client_id                  *string   `tfsdk:"client_id" json:"client_id,omitempty"`
		Client_secret              *string   `tfsdk:"client_secret" json:"client_secret,omitempty"`
		Display                    *string   `tfsdk:"display" json:"display,omitempty"`
		Google_admin_email         *string   `tfsdk:"google_admin_email" json:"google_admin_email,omitempty"`
		Google_service_account     *string   `tfsdk:"google_service_account" json:"google_service_account,omitempty"`
		Google_service_account_uri *string   `tfsdk:"google_service_account_uri" json:"google_service_account_uri,omitempty"`
		Issuer_url                 *string   `tfsdk:"issuer_url" json:"issuer_url,omitempty"`
		Max_age                    *string   `tfsdk:"max_age" json:"max_age,omitempty"`
		Prompt                     *string   `tfsdk:"prompt" json:"prompt,omitempty"`
		Provider                   *string   `tfsdk:"provider" json:"provider,omitempty"`
		Redirect_url               *[]string `tfsdk:"redirect_url" json:"redirect_url,omitempty"`
		Scope                      *[]string `tfsdk:"scope" json:"scope,omitempty"`
		Username_claim             *string   `tfsdk:"username_claim" json:"username_claim,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ResourcesTeleportDevTeleportOidcconnectorV3Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_resources_teleport_dev_teleport_oidc_connector_v3"
}

func (r *ResourcesTeleportDevTeleportOidcconnectorV3Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "OIDCConnector is the Schema for the oidcconnectors API",
		MarkdownDescription: "OIDCConnector is the Schema for the oidcconnectors API",
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
				Description:         "OIDCConnector resource definition v3 from Teleport",
				MarkdownDescription: "OIDCConnector resource definition v3 from Teleport",
				Attributes: map[string]schema.Attribute{
					"acr_values": schema.StringAttribute{
						Description:         "ACR is an Authentication Context Class Reference value. The meaning of the ACR value is context-specific and varies for identity providers.",
						MarkdownDescription: "ACR is an Authentication Context Class Reference value. The meaning of the ACR value is context-specific and varies for identity providers.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"allow_unverified_email": schema.BoolAttribute{
						Description:         "AllowUnverifiedEmail tells the connector to accept OIDC users with unverified emails.",
						MarkdownDescription: "AllowUnverifiedEmail tells the connector to accept OIDC users with unverified emails.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"claims_to_roles": schema.ListNestedAttribute{
						Description:         "ClaimsToRoles specifies a dynamic mapping from claims to roles.",
						MarkdownDescription: "ClaimsToRoles specifies a dynamic mapping from claims to roles.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"claim": schema.StringAttribute{
									Description:         "Claim is a claim name.",
									MarkdownDescription: "Claim is a claim name.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"roles": schema.ListAttribute{
									Description:         "Roles is a list of static teleport roles to match.",
									MarkdownDescription: "Roles is a list of static teleport roles to match.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"value": schema.StringAttribute{
									Description:         "Value is a claim value to match.",
									MarkdownDescription: "Value is a claim value to match.",
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

					"client_id": schema.StringAttribute{
						Description:         "ClientID is the id of the authentication client (Teleport Auth server).",
						MarkdownDescription: "ClientID is the id of the authentication client (Teleport Auth server).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"client_secret": schema.StringAttribute{
						Description:         "ClientSecret is used to authenticate the client.",
						MarkdownDescription: "ClientSecret is used to authenticate the client.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"display": schema.StringAttribute{
						Description:         "Display is the friendly name for this provider.",
						MarkdownDescription: "Display is the friendly name for this provider.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"google_admin_email": schema.StringAttribute{
						Description:         "GoogleAdminEmail is the email of a google admin to impersonate.",
						MarkdownDescription: "GoogleAdminEmail is the email of a google admin to impersonate.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"google_service_account": schema.StringAttribute{
						Description:         "GoogleServiceAccount is a string containing google service account credentials.",
						MarkdownDescription: "GoogleServiceAccount is a string containing google service account credentials.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"google_service_account_uri": schema.StringAttribute{
						Description:         "GoogleServiceAccountURI is a path to a google service account uri.",
						MarkdownDescription: "GoogleServiceAccountURI is a path to a google service account uri.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"issuer_url": schema.StringAttribute{
						Description:         "IssuerURL is the endpoint of the provider, e.g. https://accounts.google.com.",
						MarkdownDescription: "IssuerURL is the endpoint of the provider, e.g. https://accounts.google.com.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"max_age": schema.StringAttribute{
						Description:         "MaxAge is the amount of time that user logins are valid for. If a user logs in, but then does not login again within this time period, they will be forced to re-authenticate.",
						MarkdownDescription: "MaxAge is the amount of time that user logins are valid for. If a user logs in, but then does not login again within this time period, they will be forced to re-authenticate.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"prompt": schema.StringAttribute{
						Description:         "Prompt is an optional OIDC prompt. An empty string omits prompt. If not specified, it defaults to select_account for backwards compatibility.",
						MarkdownDescription: "Prompt is an optional OIDC prompt. An empty string omits prompt. If not specified, it defaults to select_account for backwards compatibility.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"provider": schema.StringAttribute{
						Description:         "Provider is the external identity provider.",
						MarkdownDescription: "Provider is the external identity provider.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"redirect_url": schema.ListAttribute{
						Description:         "RedirectURLs is a list of callback URLs which the identity provider can use to redirect the client back to the Teleport Proxy to complete authentication. This list should match the URLs on the provider's side. The URL used for a given auth request will be chosen to match the requesting Proxy's public address. If there is no match, the first url in the list will be used.",
						MarkdownDescription: "RedirectURLs is a list of callback URLs which the identity provider can use to redirect the client back to the Teleport Proxy to complete authentication. This list should match the URLs on the provider's side. The URL used for a given auth request will be chosen to match the requesting Proxy's public address. If there is no match, the first url in the list will be used.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"scope": schema.ListAttribute{
						Description:         "Scope specifies additional scopes set by provider.",
						MarkdownDescription: "Scope specifies additional scopes set by provider.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"username_claim": schema.StringAttribute{
						Description:         "UsernameClaim specifies the name of the claim from the OIDC connector to be used as the user's username.",
						MarkdownDescription: "UsernameClaim specifies the name of the claim from the OIDC connector to be used as the user's username.",
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

func (r *ResourcesTeleportDevTeleportOidcconnectorV3Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
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

func (r *ResourcesTeleportDevTeleportOidcconnectorV3Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_resources_teleport_dev_teleport_oidc_connector_v3")

	var model ResourcesTeleportDevTeleportOidcconnectorV3ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("resources.teleport.dev/v3")
	model.Kind = pointer.String("TeleportOIDCConnector")

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
		Resource(k8sSchema.GroupVersionResource{Group: "resources.teleport.dev", Version: "v3", Resource: "teleportoidcconnectors"}).
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

	var readResponse ResourcesTeleportDevTeleportOidcconnectorV3ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *ResourcesTeleportDevTeleportOidcconnectorV3Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_resources_teleport_dev_teleport_oidc_connector_v3")

	var data ResourcesTeleportDevTeleportOidcconnectorV3ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "resources.teleport.dev", Version: "v3", Resource: "teleportoidcconnectors"}).
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

	var readResponse ResourcesTeleportDevTeleportOidcconnectorV3ResourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *ResourcesTeleportDevTeleportOidcconnectorV3Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_resources_teleport_dev_teleport_oidc_connector_v3")

	var model ResourcesTeleportDevTeleportOidcconnectorV3ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("resources.teleport.dev/v3")
	model.Kind = pointer.String("TeleportOIDCConnector")

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
		Resource(k8sSchema.GroupVersionResource{Group: "resources.teleport.dev", Version: "v3", Resource: "teleportoidcconnectors"}).
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

	var readResponse ResourcesTeleportDevTeleportOidcconnectorV3ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *ResourcesTeleportDevTeleportOidcconnectorV3Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_resources_teleport_dev_teleport_oidc_connector_v3")

	var data ResourcesTeleportDevTeleportOidcconnectorV3ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "resources.teleport.dev", Version: "v3", Resource: "teleportoidcconnectors"}).
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
				Resource(k8sSchema.GroupVersionResource{Group: "resources.teleport.dev", Version: "v3", Resource: "teleportoidcconnectors"}).
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

func (r *ResourcesTeleportDevTeleportOidcconnectorV3Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
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
