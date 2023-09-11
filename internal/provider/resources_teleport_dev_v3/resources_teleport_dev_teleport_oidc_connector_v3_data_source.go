/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package resources_teleport_dev_v3

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"net/http"
)

var (
	_ datasource.DataSource              = &ResourcesTeleportDevTeleportOidcconnectorV3DataSource{}
	_ datasource.DataSourceWithConfigure = &ResourcesTeleportDevTeleportOidcconnectorV3DataSource{}
)

func NewResourcesTeleportDevTeleportOidcconnectorV3DataSource() datasource.DataSource {
	return &ResourcesTeleportDevTeleportOidcconnectorV3DataSource{}
}

type ResourcesTeleportDevTeleportOidcconnectorV3DataSource struct {
	kubernetesClient dynamic.Interface
}

type ResourcesTeleportDevTeleportOidcconnectorV3DataSourceData struct {
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

func (r *ResourcesTeleportDevTeleportOidcconnectorV3DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_resources_teleport_dev_teleport_oidc_connector_v3"
}

func (r *ResourcesTeleportDevTeleportOidcconnectorV3DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "OIDCConnector resource definition v3 from Teleport",
				MarkdownDescription: "OIDCConnector resource definition v3 from Teleport",
				Attributes: map[string]schema.Attribute{
					"acr_values": schema.StringAttribute{
						Description:         "ACR is an Authentication Context Class Reference value. The meaning of the ACR value is context-specific and varies for identity providers.",
						MarkdownDescription: "ACR is an Authentication Context Class Reference value. The meaning of the ACR value is context-specific and varies for identity providers.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"allow_unverified_email": schema.BoolAttribute{
						Description:         "AllowUnverifiedEmail tells the connector to accept OIDC users with unverified emails.",
						MarkdownDescription: "AllowUnverifiedEmail tells the connector to accept OIDC users with unverified emails.",
						Required:            false,
						Optional:            false,
						Computed:            true,
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
									Optional:            false,
									Computed:            true,
								},

								"roles": schema.ListAttribute{
									Description:         "Roles is a list of static teleport roles to match.",
									MarkdownDescription: "Roles is a list of static teleport roles to match.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"value": schema.StringAttribute{
									Description:         "Value is a claim value to match.",
									MarkdownDescription: "Value is a claim value to match.",
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

					"client_id": schema.StringAttribute{
						Description:         "ClientID is the id of the authentication client (Teleport Auth server).",
						MarkdownDescription: "ClientID is the id of the authentication client (Teleport Auth server).",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"client_secret": schema.StringAttribute{
						Description:         "ClientSecret is used to authenticate the client.",
						MarkdownDescription: "ClientSecret is used to authenticate the client.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"display": schema.StringAttribute{
						Description:         "Display is the friendly name for this provider.",
						MarkdownDescription: "Display is the friendly name for this provider.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"google_admin_email": schema.StringAttribute{
						Description:         "GoogleAdminEmail is the email of a google admin to impersonate.",
						MarkdownDescription: "GoogleAdminEmail is the email of a google admin to impersonate.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"google_service_account": schema.StringAttribute{
						Description:         "GoogleServiceAccount is a string containing google service account credentials.",
						MarkdownDescription: "GoogleServiceAccount is a string containing google service account credentials.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"google_service_account_uri": schema.StringAttribute{
						Description:         "GoogleServiceAccountURI is a path to a google service account uri.",
						MarkdownDescription: "GoogleServiceAccountURI is a path to a google service account uri.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"issuer_url": schema.StringAttribute{
						Description:         "IssuerURL is the endpoint of the provider, e.g. https://accounts.google.com.",
						MarkdownDescription: "IssuerURL is the endpoint of the provider, e.g. https://accounts.google.com.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"max_age": schema.StringAttribute{
						Description:         "MaxAge is the amount of time that user logins are valid for. If a user logs in, but then does not login again within this time period, they will be forced to re-authenticate.",
						MarkdownDescription: "MaxAge is the amount of time that user logins are valid for. If a user logs in, but then does not login again within this time period, they will be forced to re-authenticate.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"prompt": schema.StringAttribute{
						Description:         "Prompt is an optional OIDC prompt. An empty string omits prompt. If not specified, it defaults to select_account for backwards compatibility.",
						MarkdownDescription: "Prompt is an optional OIDC prompt. An empty string omits prompt. If not specified, it defaults to select_account for backwards compatibility.",
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

					"redirect_url": schema.ListAttribute{
						Description:         "RedirectURLs is a list of callback URLs which the identity provider can use to redirect the client back to the Teleport Proxy to complete authentication. This list should match the URLs on the provider's side. The URL used for a given auth request will be chosen to match the requesting Proxy's public address. If there is no match, the first url in the list will be used.",
						MarkdownDescription: "RedirectURLs is a list of callback URLs which the identity provider can use to redirect the client back to the Teleport Proxy to complete authentication. This list should match the URLs on the provider's side. The URL used for a given auth request will be chosen to match the requesting Proxy's public address. If there is no match, the first url in the list will be used.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"scope": schema.ListAttribute{
						Description:         "Scope specifies additional scopes set by provider.",
						MarkdownDescription: "Scope specifies additional scopes set by provider.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"username_claim": schema.StringAttribute{
						Description:         "UsernameClaim specifies the name of the claim from the OIDC connector to be used as the user's username.",
						MarkdownDescription: "UsernameClaim specifies the name of the claim from the OIDC connector to be used as the user's username.",
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

func (r *ResourcesTeleportDevTeleportOidcconnectorV3DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *ResourcesTeleportDevTeleportOidcconnectorV3DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_resources_teleport_dev_teleport_oidc_connector_v3")

	var data ResourcesTeleportDevTeleportOidcconnectorV3DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "resources.teleport.dev", Version: "v3", Resource: "teleportoidcconnectors"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		var statusError *k8sErrors.StatusError
		if errors.As(err, &statusError) {
			if statusError.Status().Code == http.StatusNotFound {
				response.Diagnostics.AddError(
					"Unable to find resource",
					fmt.Sprintf("The requested resource cannot be found. "+
						"Make sure that it does exist in your cluster and you have set the correct name and namespace configured.\n\n"+
						"Namespace: %s\n"+
						"Name: %s", data.Metadata.Namespace, data.Metadata.Name),
				)
				return
			}
		} else {
			response.Diagnostics.AddError(
				"Unable to GET resource",
				fmt.Sprintf("An unexpected error occurred while reading the resource. "+
					"Please report this issue to the provider developers.\n\n"+
					"GET Error (%T): %s", err, err.Error()),
			)
		}
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

	var readResponse ResourcesTeleportDevTeleportOidcconnectorV3DataSourceData
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
	data.ApiVersion = pointer.String("resources.teleport.dev/v3")
	data.Kind = pointer.String("TeleportOIDCConnector")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
