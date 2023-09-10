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
	_ datasource.DataSource              = &ResourcesTeleportDevTeleportSamlconnectorV2DataSource{}
	_ datasource.DataSourceWithConfigure = &ResourcesTeleportDevTeleportSamlconnectorV2DataSource{}
)

func NewResourcesTeleportDevTeleportSamlconnectorV2DataSource() datasource.DataSource {
	return &ResourcesTeleportDevTeleportSamlconnectorV2DataSource{}
}

type ResourcesTeleportDevTeleportSamlconnectorV2DataSource struct {
	kubernetesClient dynamic.Interface
}

type ResourcesTeleportDevTeleportSamlconnectorV2DataSourceData struct {
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

func (r *ResourcesTeleportDevTeleportSamlconnectorV2DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_resources_teleport_dev_teleport_saml_connector_v2"
}

func (r *ResourcesTeleportDevTeleportSamlconnectorV2DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "SAMLConnector resource definition v2 from Teleport",
				MarkdownDescription: "SAMLConnector resource definition v2 from Teleport",
				Attributes: map[string]schema.Attribute{
					"acs": schema.StringAttribute{
						Description:         "AssertionConsumerService is a URL for assertion consumer service on the service provider (Teleport's side).",
						MarkdownDescription: "AssertionConsumerService is a URL for assertion consumer service on the service provider (Teleport's side).",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"allow_idp_initiated": schema.BoolAttribute{
						Description:         "AllowIDPInitiated is a flag that indicates if the connector can be used for IdP-initiated logins.",
						MarkdownDescription: "AllowIDPInitiated is a flag that indicates if the connector can be used for IdP-initiated logins.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"assertion_key_pair": schema.SingleNestedAttribute{
						Description:         "EncryptionKeyPair is a key pair used for decrypting SAML assertions.",
						MarkdownDescription: "EncryptionKeyPair is a key pair used for decrypting SAML assertions.",
						Attributes: map[string]schema.Attribute{
							"cert": schema.StringAttribute{
								Description:         "Cert is a PEM-encoded x509 certificate.",
								MarkdownDescription: "Cert is a PEM-encoded x509 certificate.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"private_key": schema.StringAttribute{
								Description:         "PrivateKey is a PEM encoded x509 private key.",
								MarkdownDescription: "PrivateKey is a PEM encoded x509 private key.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
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
									Optional:            false,
									Computed:            true,
								},

								"roles": schema.ListAttribute{
									Description:         "Roles is a list of static teleport roles to map to.",
									MarkdownDescription: "Roles is a list of static teleport roles to map to.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"value": schema.StringAttribute{
									Description:         "Value is an attribute statement value to match.",
									MarkdownDescription: "Value is an attribute statement value to match.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"audience": schema.StringAttribute{
						Description:         "Audience uniquely identifies our service provider.",
						MarkdownDescription: "Audience uniquely identifies our service provider.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"cert": schema.StringAttribute{
						Description:         "Cert is the identity provider certificate PEM. IDP signs <Response> responses using this certificate.",
						MarkdownDescription: "Cert is the identity provider certificate PEM. IDP signs <Response> responses using this certificate.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"display": schema.StringAttribute{
						Description:         "Display controls how this connector is displayed.",
						MarkdownDescription: "Display controls how this connector is displayed.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"entity_descriptor": schema.StringAttribute{
						Description:         "EntityDescriptor is XML with descriptor. It can be used to supply configuration parameters in one XML file rather than supplying them in the individual elements.",
						MarkdownDescription: "EntityDescriptor is XML with descriptor. It can be used to supply configuration parameters in one XML file rather than supplying them in the individual elements.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"entity_descriptor_url": schema.StringAttribute{
						Description:         "EntityDescriptorURL is a URL that supplies a configuration XML.",
						MarkdownDescription: "EntityDescriptorURL is a URL that supplies a configuration XML.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"issuer": schema.StringAttribute{
						Description:         "Issuer is the identity provider issuer.",
						MarkdownDescription: "Issuer is the identity provider issuer.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"provider": schema.StringAttribute{
						Description:         "Provider is the external identity provider.",
						MarkdownDescription: "Provider is the external identity provider.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"service_provider_issuer": schema.StringAttribute{
						Description:         "ServiceProviderIssuer is the issuer of the service provider (Teleport).",
						MarkdownDescription: "ServiceProviderIssuer is the issuer of the service provider (Teleport).",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"signing_key_pair": schema.SingleNestedAttribute{
						Description:         "SigningKeyPair is an x509 key pair used to sign AuthnRequest.",
						MarkdownDescription: "SigningKeyPair is an x509 key pair used to sign AuthnRequest.",
						Attributes: map[string]schema.Attribute{
							"cert": schema.StringAttribute{
								Description:         "Cert is a PEM-encoded x509 certificate.",
								MarkdownDescription: "Cert is a PEM-encoded x509 certificate.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"private_key": schema.StringAttribute{
								Description:         "PrivateKey is a PEM encoded x509 private key.",
								MarkdownDescription: "PrivateKey is a PEM encoded x509 private key.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"sso": schema.StringAttribute{
						Description:         "SSO is the URL of the identity provider's SSO service.",
						MarkdownDescription: "SSO is the URL of the identity provider's SSO service.",
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

func (r *ResourcesTeleportDevTeleportSamlconnectorV2DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *ResourcesTeleportDevTeleportSamlconnectorV2DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_resources_teleport_dev_teleport_saml_connector_v2")

	var data ResourcesTeleportDevTeleportSamlconnectorV2DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
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

	var readResponse ResourcesTeleportDevTeleportSamlconnectorV2DataSourceData
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

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("resources.teleport.dev/v2")
	data.Kind = pointer.String("TeleportSAMLConnector")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
