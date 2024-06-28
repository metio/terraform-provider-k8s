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
	_ datasource.DataSource = &ResourcesTeleportDevTeleportUserV2Manifest{}
)

func NewResourcesTeleportDevTeleportUserV2Manifest() datasource.DataSource {
	return &ResourcesTeleportDevTeleportUserV2Manifest{}
}

type ResourcesTeleportDevTeleportUserV2Manifest struct{}

type ResourcesTeleportDevTeleportUserV2ManifestData struct {
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
		Github_identities *[]struct {
			Connector_id        *string `tfsdk:"connector_id" json:"connector_id,omitempty"`
			SamlSingleLogoutUrl *string `tfsdk:"saml_single_logout_url" json:"samlSingleLogoutUrl,omitempty"`
			Username            *string `tfsdk:"username" json:"username,omitempty"`
		} `tfsdk:"github_identities" json:"github_identities,omitempty"`
		Oidc_identities *[]struct {
			Connector_id        *string `tfsdk:"connector_id" json:"connector_id,omitempty"`
			SamlSingleLogoutUrl *string `tfsdk:"saml_single_logout_url" json:"samlSingleLogoutUrl,omitempty"`
			Username            *string `tfsdk:"username" json:"username,omitempty"`
		} `tfsdk:"oidc_identities" json:"oidc_identities,omitempty"`
		Roles           *[]string `tfsdk:"roles" json:"roles,omitempty"`
		Saml_identities *[]struct {
			Connector_id        *string `tfsdk:"connector_id" json:"connector_id,omitempty"`
			SamlSingleLogoutUrl *string `tfsdk:"saml_single_logout_url" json:"samlSingleLogoutUrl,omitempty"`
			Username            *string `tfsdk:"username" json:"username,omitempty"`
		} `tfsdk:"saml_identities" json:"saml_identities,omitempty"`
		Traits             *map[string][]string `tfsdk:"traits" json:"traits,omitempty"`
		Trusted_device_ids *[]string            `tfsdk:"trusted_device_ids" json:"trusted_device_ids,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ResourcesTeleportDevTeleportUserV2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_resources_teleport_dev_teleport_user_v2_manifest"
}

func (r *ResourcesTeleportDevTeleportUserV2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "User is the Schema for the users API",
		MarkdownDescription: "User is the Schema for the users API",
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

								"saml_single_logout_url": schema.StringAttribute{
									Description:         "SAMLSingleLogoutURL is the SAML Single log-out URL to initiate SAML SLO (single log-out), if applicable.",
									MarkdownDescription: "SAMLSingleLogoutURL is the SAML Single log-out URL to initiate SAML SLO (single log-out), if applicable.",
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

								"saml_single_logout_url": schema.StringAttribute{
									Description:         "SAMLSingleLogoutURL is the SAML Single log-out URL to initiate SAML SLO (single log-out), if applicable.",
									MarkdownDescription: "SAMLSingleLogoutURL is the SAML Single log-out URL to initiate SAML SLO (single log-out), if applicable.",
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

								"saml_single_logout_url": schema.StringAttribute{
									Description:         "SAMLSingleLogoutURL is the SAML Single log-out URL to initiate SAML SLO (single log-out), if applicable.",
									MarkdownDescription: "SAMLSingleLogoutURL is the SAML Single log-out URL to initiate SAML SLO (single log-out), if applicable.",
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

func (r *ResourcesTeleportDevTeleportUserV2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_resources_teleport_dev_teleport_user_v2_manifest")

	var model ResourcesTeleportDevTeleportUserV2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("resources.teleport.dev/v2")
	model.Kind = pointer.String("TeleportUser")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
