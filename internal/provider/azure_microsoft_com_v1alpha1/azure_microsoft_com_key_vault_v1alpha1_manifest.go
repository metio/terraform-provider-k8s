/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package azure_microsoft_com_v1alpha1

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &AzureMicrosoftComKeyVaultV1Alpha1Manifest{}
)

func NewAzureMicrosoftComKeyVaultV1Alpha1Manifest() datasource.DataSource {
	return &AzureMicrosoftComKeyVaultV1Alpha1Manifest{}
}

type AzureMicrosoftComKeyVaultV1Alpha1Manifest struct{}

type AzureMicrosoftComKeyVaultV1Alpha1ManifestData struct {
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
		AccessPolicies *[]struct {
			ApplicationID *string `tfsdk:"application_id" json:"applicationID,omitempty"`
			ClientID      *string `tfsdk:"client_id" json:"clientID,omitempty"`
			ObjectID      *string `tfsdk:"object_id" json:"objectID,omitempty"`
			Permissions   *struct {
				Certificates *[]string `tfsdk:"certificates" json:"certificates,omitempty"`
				Keys         *[]string `tfsdk:"keys" json:"keys,omitempty"`
				Secrets      *[]string `tfsdk:"secrets" json:"secrets,omitempty"`
				Storage      *[]string `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"permissions" json:"permissions,omitempty"`
			TenantID *string `tfsdk:"tenant_id" json:"tenantID,omitempty"`
		} `tfsdk:"access_policies" json:"accessPolicies,omitempty"`
		EnableSoftDelete *bool   `tfsdk:"enable_soft_delete" json:"enableSoftDelete,omitempty"`
		Location         *string `tfsdk:"location" json:"location,omitempty"`
		NetworkPolicies  *struct {
			Bypass              *string   `tfsdk:"bypass" json:"bypass,omitempty"`
			DefaultAction       *string   `tfsdk:"default_action" json:"defaultAction,omitempty"`
			IpRules             *[]string `tfsdk:"ip_rules" json:"ipRules,omitempty"`
			VirtualNetworkRules *[]string `tfsdk:"virtual_network_rules" json:"virtualNetworkRules,omitempty"`
		} `tfsdk:"network_policies" json:"networkPolicies,omitempty"`
		ResourceGroup *string `tfsdk:"resource_group" json:"resourceGroup,omitempty"`
		Sku           *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"sku" json:"sku,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AzureMicrosoftComKeyVaultV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_azure_microsoft_com_key_vault_v1alpha1_manifest"
}

func (r *AzureMicrosoftComKeyVaultV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "KeyVault is the Schema for the keyvaults API",
		MarkdownDescription: "KeyVault is the Schema for the keyvaults API",
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
				Description:         "KeyVaultSpec defines the desired state of KeyVault",
				MarkdownDescription: "KeyVaultSpec defines the desired state of KeyVault",
				Attributes: map[string]schema.Attribute{
					"access_policies": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"application_id": schema.StringAttribute{
									Description:         "ApplicationID -  Application ID of the client making request on behalf of a principal",
									MarkdownDescription: "ApplicationID -  Application ID of the client making request on behalf of a principal",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"client_id": schema.StringAttribute{
									Description:         "ClientID - The client ID of a user, service principal or security group in the Azure Active Directory tenant for the vault. The client ID must be unique for the list of access policies. TODO: Remove this in a future API version, see: https://github.com/Azure/azure-service-operator/issues/1351",
									MarkdownDescription: "ClientID - The client ID of a user, service principal or security group in the Azure Active Directory tenant for the vault. The client ID must be unique for the list of access policies. TODO: Remove this in a future API version, see: https://github.com/Azure/azure-service-operator/issues/1351",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"object_id": schema.StringAttribute{
									Description:         "ObjectID is the AAD object id of the entity to provide access to.",
									MarkdownDescription: "ObjectID is the AAD object id of the entity to provide access to.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"permissions": schema.SingleNestedAttribute{
									Description:         "Permissions - Permissions the identity has for keys, secrets, and certificates.",
									MarkdownDescription: "Permissions - Permissions the identity has for keys, secrets, and certificates.",
									Attributes: map[string]schema.Attribute{
										"certificates": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"keys": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secrets": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"storage": schema.ListAttribute{
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

								"tenant_id": schema.StringAttribute{
									Description:         "TenantID - The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.",
									MarkdownDescription: "TenantID - The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.",
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

					"enable_soft_delete": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"location": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"network_policies": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"bypass": schema.StringAttribute{
								Description:         "Bypass - Tells what traffic can bypass network rules. This can be 'AzureServices' or 'None'.  If not specified the default is 'AzureServices'. Possible values include: 'AzureServices', 'None'",
								MarkdownDescription: "Bypass - Tells what traffic can bypass network rules. This can be 'AzureServices' or 'None'.  If not specified the default is 'AzureServices'. Possible values include: 'AzureServices', 'None'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"default_action": schema.StringAttribute{
								Description:         "DefaultAction - The default action when no rule from ipRules and from virtualNetworkRules match. This is only used after the bypass property has been evaluated. Possible values include: 'Allow', 'Deny'",
								MarkdownDescription: "DefaultAction - The default action when no rule from ipRules and from virtualNetworkRules match. This is only used after the bypass property has been evaluated. Possible values include: 'Allow', 'Deny'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ip_rules": schema.ListAttribute{
								Description:         "IPRules - The list of IP address rules.",
								MarkdownDescription: "IPRules - The list of IP address rules.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"virtual_network_rules": schema.ListAttribute{
								Description:         "VirtualNetworkRules - The list of virtual network rules.",
								MarkdownDescription: "VirtualNetworkRules - The list of virtual network rules.",
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

					"resource_group": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
							stringvalidator.RegexMatches(regexp.MustCompile(`^[-\w\._\(\)]+$`), ""),
						},
					},

					"sku": schema.SingleNestedAttribute{
						Description:         "KeyVaultSku the SKU of the Key Vault",
						MarkdownDescription: "KeyVaultSku the SKU of the Key Vault",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name - The SKU name. Required for account creation; optional for update. Possible values include: 'Premium', 'Standard'",
								MarkdownDescription: "Name - The SKU name. Required for account creation; optional for update. Possible values include: 'Premium', 'Standard'",
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
	}
}

func (r *AzureMicrosoftComKeyVaultV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_azure_microsoft_com_key_vault_v1alpha1_manifest")

	var model AzureMicrosoftComKeyVaultV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("azure.microsoft.com/v1alpha1")
	model.Kind = pointer.String("KeyVault")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
