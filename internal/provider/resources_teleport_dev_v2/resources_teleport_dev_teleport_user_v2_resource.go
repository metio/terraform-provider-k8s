/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package resources_teleport_dev_v2

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
)

var (
	_ resource.Resource                = &ResourcesTeleportDevTeleportUserV2Resource{}
	_ resource.ResourceWithConfigure   = &ResourcesTeleportDevTeleportUserV2Resource{}
	_ resource.ResourceWithImportState = &ResourcesTeleportDevTeleportUserV2Resource{}
)

func NewResourcesTeleportDevTeleportUserV2Resource() resource.Resource {
	return &ResourcesTeleportDevTeleportUserV2Resource{}
}

type ResourcesTeleportDevTeleportUserV2Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type ResourcesTeleportDevTeleportUserV2ResourceData struct {
	ID             types.String `tfsdk:"id" json:"-"`
	ForceConflicts types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager   types.String `tfsdk:"field_manager" json:"-"`
	WaitFor        types.List   `tfsdk:"wait_for" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Github_identities *[]struct {
			Connector_id *string `tfsdk:"connector_id" json:"connector_id,omitempty"`
			Username     *string `tfsdk:"username" json:"username,omitempty"`
		} `tfsdk:"github_identities" json:"github_identities,omitempty"`
		Oidc_identities *[]struct {
			Connector_id *string `tfsdk:"connector_id" json:"connector_id,omitempty"`
			Username     *string `tfsdk:"username" json:"username,omitempty"`
		} `tfsdk:"oidc_identities" json:"oidc_identities,omitempty"`
		Roles           *[]string `tfsdk:"roles" json:"roles,omitempty"`
		Saml_identities *[]struct {
			Connector_id *string `tfsdk:"connector_id" json:"connector_id,omitempty"`
			Username     *string `tfsdk:"username" json:"username,omitempty"`
		} `tfsdk:"saml_identities" json:"saml_identities,omitempty"`
		Traits             *map[string][]string `tfsdk:"traits" json:"traits,omitempty"`
		Trusted_device_ids *[]string            `tfsdk:"trusted_device_ids" json:"trusted_device_ids,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ResourcesTeleportDevTeleportUserV2Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_resources_teleport_dev_teleport_user_v2"
}

func (r *ResourcesTeleportDevTeleportUserV2Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "User is the Schema for the users API",
		MarkdownDescription: "User is the Schema for the users API",
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
				Description:         "User resource definition v2 from Teleport",
				MarkdownDescription: "User resource definition v2 from Teleport",
				Attributes: map[string]schema.Attribute{
					"github_identities": schema.ListNestedAttribute{
						Description:         "GithubIdentities list associated Github OAuth2 identities that let user log in using externally verified identity",
						MarkdownDescription: "GithubIdentities list associated Github OAuth2 identities that let user log in using externally verified identity",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"connector_id": schema.StringAttribute{
									Description:         "ConnectorID is id of registered OIDC connector, e.g. 'google-example.com'",
									MarkdownDescription: "ConnectorID is id of registered OIDC connector, e.g. 'google-example.com'",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"username": schema.StringAttribute{
									Description:         "Username is username supplied by external identity provider",
									MarkdownDescription: "Username is username supplied by external identity provider",
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

					"oidc_identities": schema.ListNestedAttribute{
						Description:         "OIDCIdentities lists associated OpenID Connect identities that let user log in using externally verified identity",
						MarkdownDescription: "OIDCIdentities lists associated OpenID Connect identities that let user log in using externally verified identity",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"connector_id": schema.StringAttribute{
									Description:         "ConnectorID is id of registered OIDC connector, e.g. 'google-example.com'",
									MarkdownDescription: "ConnectorID is id of registered OIDC connector, e.g. 'google-example.com'",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"username": schema.StringAttribute{
									Description:         "Username is username supplied by external identity provider",
									MarkdownDescription: "Username is username supplied by external identity provider",
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

					"roles": schema.ListAttribute{
						Description:         "Roles is a list of roles assigned to user",
						MarkdownDescription: "Roles is a list of roles assigned to user",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"saml_identities": schema.ListNestedAttribute{
						Description:         "SAMLIdentities lists associated SAML identities that let user log in using externally verified identity",
						MarkdownDescription: "SAMLIdentities lists associated SAML identities that let user log in using externally verified identity",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"connector_id": schema.StringAttribute{
									Description:         "ConnectorID is id of registered OIDC connector, e.g. 'google-example.com'",
									MarkdownDescription: "ConnectorID is id of registered OIDC connector, e.g. 'google-example.com'",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"username": schema.StringAttribute{
									Description:         "Username is username supplied by external identity provider",
									MarkdownDescription: "Username is username supplied by external identity provider",
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

					"traits": schema.MapAttribute{
						Description:         "Traits are key/value pairs received from an identity provider (through OIDC claims or SAML assertions) or from a system administrator for local accounts. Traits are used to populate role variables.",
						MarkdownDescription: "Traits are key/value pairs received from an identity provider (through OIDC claims or SAML assertions) or from a system administrator for local accounts. Traits are used to populate role variables.",
						ElementType:         types.ListType{ElemType: types.StringType},
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"trusted_device_ids": schema.ListAttribute{
						Description:         "TrustedDeviceIDs contains the IDs of trusted devices enrolled by the user. Managed by the Device Trust subsystem, avoid manual edits.",
						MarkdownDescription: "TrustedDeviceIDs contains the IDs of trusted devices enrolled by the user. Managed by the Device Trust subsystem, avoid manual edits.",
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

func (r *ResourcesTeleportDevTeleportUserV2Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
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

func (r *ResourcesTeleportDevTeleportUserV2Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_resources_teleport_dev_teleport_user_v2")

	var model ResourcesTeleportDevTeleportUserV2ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("resources.teleport.dev/v2")
	model.Kind = pointer.String("TeleportUser")

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

	patchResponse, err := r.kubernetesClient.Resource(k8sSchema.GroupVersionResource{Group: "resources.teleport.dev", Version: "v2", Resource: "TeleportUser"}).
		Namespace(model.Metadata.Namespace).
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

	var readResponse ResourcesTeleportDevTeleportUserV2ResourceData
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
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *ResourcesTeleportDevTeleportUserV2Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_resources_teleport_dev_teleport_user_v2")

	var data ResourcesTeleportDevTeleportUserV2ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "resources.teleport.dev", Version: "v2", Resource: "TeleportUser"}).
		Namespace(data.Metadata.Namespace).
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

	var readResponse ResourcesTeleportDevTeleportUserV2ResourceData
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
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *ResourcesTeleportDevTeleportUserV2Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_resources_teleport_dev_teleport_user_v2")

	var model ResourcesTeleportDevTeleportUserV2ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("resources.teleport.dev/v2")
	model.Kind = pointer.String("TeleportUser")

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

	patchResponse, err := r.kubernetesClient.Resource(k8sSchema.GroupVersionResource{Group: "resources.teleport.dev", Version: "v2", Resource: "TeleportUser"}).
		Namespace(model.Metadata.Namespace).
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

	var readResponse ResourcesTeleportDevTeleportUserV2ResourceData
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
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *ResourcesTeleportDevTeleportUserV2Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_resources_teleport_dev_teleport_user_v2")

	var data ResourcesTeleportDevTeleportUserV2ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "resources.teleport.dev", Version: "v2", Resource: "TeleportUser"}).
		Namespace(data.Metadata.Namespace).
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

func (r *ResourcesTeleportDevTeleportUserV2Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
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
