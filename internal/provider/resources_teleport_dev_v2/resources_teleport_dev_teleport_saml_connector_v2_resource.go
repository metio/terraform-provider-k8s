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
	_ resource.Resource                = &ResourcesTeleportDevTeleportSAMLConnectorV2Resource{}
	_ resource.ResourceWithConfigure   = &ResourcesTeleportDevTeleportSAMLConnectorV2Resource{}
	_ resource.ResourceWithImportState = &ResourcesTeleportDevTeleportSAMLConnectorV2Resource{}
)

func NewResourcesTeleportDevTeleportSAMLConnectorV2Resource() resource.Resource {
	return &ResourcesTeleportDevTeleportSAMLConnectorV2Resource{}
}

type ResourcesTeleportDevTeleportSAMLConnectorV2Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type ResourcesTeleportDevTeleportSAMLConnectorV2ResourceData struct {
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
		Acs                 *string `tfsdk:"acs" json:"acs,omitempty"`
		Allow_idp_initiated *bool   `tfsdk:"allow_idp_initiated" json:"allow_idp_initiated,omitempty"`
		Assertion_key_pair  *struct {
			Cert        *string `tfsdk:"cert" json:"cert,omitempty"`
			Private_key *string `tfsdk:"private_key" json:"private_key,omitempty"`
		} `tfsdk:"assertion_key_pair" json:"assertion_key_pair,omitempty"`
		Attributes_to_roles *[]struct {
			Name  *string   `tfsdk:"name" json:"name,omitempty"`
			Roles *[]string `tfsdk:"roles" json:"roles,omitempty"`
			Value *string   `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"attributes_to_roles" json:"attributes_to_roles,omitempty"`
		Audience                *string `tfsdk:"audience" json:"audience,omitempty"`
		Cert                    *string `tfsdk:"cert" json:"cert,omitempty"`
		Display                 *string `tfsdk:"display" json:"display,omitempty"`
		Entity_descriptor       *string `tfsdk:"entity_descriptor" json:"entity_descriptor,omitempty"`
		Entity_descriptor_url   *string `tfsdk:"entity_descriptor_url" json:"entity_descriptor_url,omitempty"`
		Issuer                  *string `tfsdk:"issuer" json:"issuer,omitempty"`
		Provider                *string `tfsdk:"provider" json:"provider,omitempty"`
		Service_provider_issuer *string `tfsdk:"service_provider_issuer" json:"service_provider_issuer,omitempty"`
		Signing_key_pair        *struct {
			Cert        *string `tfsdk:"cert" json:"cert,omitempty"`
			Private_key *string `tfsdk:"private_key" json:"private_key,omitempty"`
		} `tfsdk:"signing_key_pair" json:"signing_key_pair,omitempty"`
		Sso *string `tfsdk:"sso" json:"sso,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ResourcesTeleportDevTeleportSAMLConnectorV2Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_resources_teleport_dev_teleport_saml_connector_v2"
}

func (r *ResourcesTeleportDevTeleportSAMLConnectorV2Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "SAMLConnector is the Schema for the samlconnectors API",
		MarkdownDescription: "SAMLConnector is the Schema for the samlconnectors API",
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
				Description:         "SAMLConnector resource definition v2 from Teleport",
				MarkdownDescription: "SAMLConnector resource definition v2 from Teleport",
				Attributes: map[string]schema.Attribute{
					"acs": schema.StringAttribute{
						Description:         "AssertionConsumerService is a URL for assertion consumer service on the service provider (Teleport's side).",
						MarkdownDescription: "AssertionConsumerService is a URL for assertion consumer service on the service provider (Teleport's side).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"allow_idp_initiated": schema.BoolAttribute{
						Description:         "AllowIDPInitiated is a flag that indicates if the connector can be used for IdP-initiated logins.",
						MarkdownDescription: "AllowIDPInitiated is a flag that indicates if the connector can be used for IdP-initiated logins.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"assertion_key_pair": schema.SingleNestedAttribute{
						Description:         "EncryptionKeyPair is a key pair used for decrypting SAML assertions.",
						MarkdownDescription: "EncryptionKeyPair is a key pair used for decrypting SAML assertions.",
						Attributes: map[string]schema.Attribute{
							"cert": schema.StringAttribute{
								Description:         "Cert is a PEM-encoded x509 certificate.",
								MarkdownDescription: "Cert is a PEM-encoded x509 certificate.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"private_key": schema.StringAttribute{
								Description:         "PrivateKey is a PEM encoded x509 private key.",
								MarkdownDescription: "PrivateKey is a PEM encoded x509 private key.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"attributes_to_roles": schema.ListNestedAttribute{
						Description:         "AttributesToRoles is a list of mappings of attribute statements to roles.",
						MarkdownDescription: "AttributesToRoles is a list of mappings of attribute statements to roles.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name is an attribute statement name.",
									MarkdownDescription: "Name is an attribute statement name.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"roles": schema.ListAttribute{
									Description:         "Roles is a list of static teleport roles to map to.",
									MarkdownDescription: "Roles is a list of static teleport roles to map to.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"value": schema.StringAttribute{
									Description:         "Value is an attribute statement value to match.",
									MarkdownDescription: "Value is an attribute statement value to match.",
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

					"audience": schema.StringAttribute{
						Description:         "Audience uniquely identifies our service provider.",
						MarkdownDescription: "Audience uniquely identifies our service provider.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"cert": schema.StringAttribute{
						Description:         "Cert is the identity provider certificate PEM. IDP signs <Response> responses using this certificate.",
						MarkdownDescription: "Cert is the identity provider certificate PEM. IDP signs <Response> responses using this certificate.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"display": schema.StringAttribute{
						Description:         "Display controls how this connector is displayed.",
						MarkdownDescription: "Display controls how this connector is displayed.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"entity_descriptor": schema.StringAttribute{
						Description:         "EntityDescriptor is XML with descriptor. It can be used to supply configuration parameters in one XML file rather than supplying them in the individual elements.",
						MarkdownDescription: "EntityDescriptor is XML with descriptor. It can be used to supply configuration parameters in one XML file rather than supplying them in the individual elements.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"entity_descriptor_url": schema.StringAttribute{
						Description:         "EntityDescriptorURL is a URL that supplies a configuration XML.",
						MarkdownDescription: "EntityDescriptorURL is a URL that supplies a configuration XML.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"issuer": schema.StringAttribute{
						Description:         "Issuer is the identity provider issuer.",
						MarkdownDescription: "Issuer is the identity provider issuer.",
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

					"service_provider_issuer": schema.StringAttribute{
						Description:         "ServiceProviderIssuer is the issuer of the service provider (Teleport).",
						MarkdownDescription: "ServiceProviderIssuer is the issuer of the service provider (Teleport).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"signing_key_pair": schema.SingleNestedAttribute{
						Description:         "SigningKeyPair is an x509 key pair used to sign AuthnRequest.",
						MarkdownDescription: "SigningKeyPair is an x509 key pair used to sign AuthnRequest.",
						Attributes: map[string]schema.Attribute{
							"cert": schema.StringAttribute{
								Description:         "Cert is a PEM-encoded x509 certificate.",
								MarkdownDescription: "Cert is a PEM-encoded x509 certificate.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"private_key": schema.StringAttribute{
								Description:         "PrivateKey is a PEM encoded x509 private key.",
								MarkdownDescription: "PrivateKey is a PEM encoded x509 private key.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"sso": schema.StringAttribute{
						Description:         "SSO is the URL of the identity provider's SSO service.",
						MarkdownDescription: "SSO is the URL of the identity provider's SSO service.",
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

func (r *ResourcesTeleportDevTeleportSAMLConnectorV2Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
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

func (r *ResourcesTeleportDevTeleportSAMLConnectorV2Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_resources_teleport_dev_teleport_saml_connector_v2")

	var model ResourcesTeleportDevTeleportSAMLConnectorV2ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("resources.teleport.dev/v2")
	model.Kind = pointer.String("TeleportSAMLConnector")

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

	patchResponse, err := r.kubernetesClient.Resource(k8sSchema.GroupVersionResource{Group: "resources.teleport.dev", Version: "v2", Resource: "TeleportSAMLConnector"}).
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

	var readResponse ResourcesTeleportDevTeleportSAMLConnectorV2ResourceData
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

func (r *ResourcesTeleportDevTeleportSAMLConnectorV2Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_resources_teleport_dev_teleport_saml_connector_v2")

	var data ResourcesTeleportDevTeleportSAMLConnectorV2ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "resources.teleport.dev", Version: "v2", Resource: "TeleportSAMLConnector"}).
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

	var readResponse ResourcesTeleportDevTeleportSAMLConnectorV2ResourceData
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

func (r *ResourcesTeleportDevTeleportSAMLConnectorV2Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_resources_teleport_dev_teleport_saml_connector_v2")

	var model ResourcesTeleportDevTeleportSAMLConnectorV2ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("resources.teleport.dev/v2")
	model.Kind = pointer.String("TeleportSAMLConnector")

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

	patchResponse, err := r.kubernetesClient.Resource(k8sSchema.GroupVersionResource{Group: "resources.teleport.dev", Version: "v2", Resource: "TeleportSAMLConnector"}).
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

	var readResponse ResourcesTeleportDevTeleportSAMLConnectorV2ResourceData
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

func (r *ResourcesTeleportDevTeleportSAMLConnectorV2Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_resources_teleport_dev_teleport_saml_connector_v2")

	var data ResourcesTeleportDevTeleportSAMLConnectorV2ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "resources.teleport.dev", Version: "v2", Resource: "TeleportSAMLConnector"}).
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

func (r *ResourcesTeleportDevTeleportSAMLConnectorV2Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
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