/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package resources_teleport_dev_v2

import (
	"context"
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
	_ datasource.DataSource = &ResourcesTeleportDevTeleportSamlconnectorV2Manifest{}
)

func NewResourcesTeleportDevTeleportSamlconnectorV2Manifest() datasource.DataSource {
	return &ResourcesTeleportDevTeleportSamlconnectorV2Manifest{}
}

type ResourcesTeleportDevTeleportSamlconnectorV2Manifest struct{}

type ResourcesTeleportDevTeleportSamlconnectorV2ManifestData struct {
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
		Audience                 *string `tfsdk:"audience" json:"audience,omitempty"`
		Cert                     *string `tfsdk:"cert" json:"cert,omitempty"`
		Client_redirect_settings *struct {
			Allowed_https_hostnames *[]string `tfsdk:"allowed_https_hostnames" json:"allowed_https_hostnames,omitempty"`
		} `tfsdk:"client_redirect_settings" json:"client_redirect_settings,omitempty"`
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
		Single_logout_url *string `tfsdk:"single_logout_url" json:"single_logout_url,omitempty"`
		Sso               *string `tfsdk:"sso" json:"sso,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ResourcesTeleportDevTeleportSamlconnectorV2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_resources_teleport_dev_teleport_saml_connector_v2_manifest"
}

func (r *ResourcesTeleportDevTeleportSamlconnectorV2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "SAMLConnector is the Schema for the samlconnectors API",
		MarkdownDescription: "SAMLConnector is the Schema for the samlconnectors API",
		Attributes: map[string]schema.Attribute{
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

					"client_redirect_settings": schema.SingleNestedAttribute{
						Description:         "ClientRedirectSettings defines which client redirect URLs are allowed for non-browser SSO logins other than the standard localhost ones.",
						MarkdownDescription: "ClientRedirectSettings defines which client redirect URLs are allowed for non-browser SSO logins other than the standard localhost ones.",
						Attributes: map[string]schema.Attribute{
							"allowed_https_hostnames": schema.ListAttribute{
								Description:         "a list of hostnames allowed for https client redirect URLs",
								MarkdownDescription: "a list of hostnames allowed for https client redirect URLs",
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

					"single_logout_url": schema.StringAttribute{
						Description:         "SingleLogoutURL is the SAML Single log-out URL to initiate SAML SLO (single log-out). If this is not provided, SLO is disabled.",
						MarkdownDescription: "SingleLogoutURL is the SAML Single log-out URL to initiate SAML SLO (single log-out). If this is not provided, SLO is disabled.",
						Required:            false,
						Optional:            true,
						Computed:            false,
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

func (r *ResourcesTeleportDevTeleportSamlconnectorV2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_resources_teleport_dev_teleport_saml_connector_v2_manifest")

	var model ResourcesTeleportDevTeleportSamlconnectorV2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("resources.teleport.dev/v2")
	model.Kind = pointer.String("TeleportSAMLConnector")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
