/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package resources_teleport_dev_v3

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
	_ datasource.DataSource = &ResourcesTeleportDevTeleportGithubConnectorV3Manifest{}
)

func NewResourcesTeleportDevTeleportGithubConnectorV3Manifest() datasource.DataSource {
	return &ResourcesTeleportDevTeleportGithubConnectorV3Manifest{}
}

type ResourcesTeleportDevTeleportGithubConnectorV3Manifest struct{}

type ResourcesTeleportDevTeleportGithubConnectorV3ManifestData struct {
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
		Api_endpoint_url         *string `tfsdk:"api_endpoint_url" json:"api_endpoint_url,omitempty"`
		Client_id                *string `tfsdk:"client_id" json:"client_id,omitempty"`
		Client_redirect_settings *struct {
			Allowed_https_hostnames      *[]string `tfsdk:"allowed_https_hostnames" json:"allowed_https_hostnames,omitempty"`
			Insecure_allowed_cidr_ranges *[]string `tfsdk:"insecure_allowed_cidr_ranges" json:"insecure_allowed_cidr_ranges,omitempty"`
		} `tfsdk:"client_redirect_settings" json:"client_redirect_settings,omitempty"`
		Client_secret  *string `tfsdk:"client_secret" json:"client_secret,omitempty"`
		Display        *string `tfsdk:"display" json:"display,omitempty"`
		Endpoint_url   *string `tfsdk:"endpoint_url" json:"endpoint_url,omitempty"`
		Redirect_url   *string `tfsdk:"redirect_url" json:"redirect_url,omitempty"`
		Teams_to_roles *[]struct {
			Organization *string   `tfsdk:"organization" json:"organization,omitempty"`
			Roles        *[]string `tfsdk:"roles" json:"roles,omitempty"`
			Team         *string   `tfsdk:"team" json:"team,omitempty"`
		} `tfsdk:"teams_to_roles" json:"teams_to_roles,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ResourcesTeleportDevTeleportGithubConnectorV3Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_resources_teleport_dev_teleport_github_connector_v3_manifest"
}

func (r *ResourcesTeleportDevTeleportGithubConnectorV3Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "GithubConnector is the Schema for the githubconnectors API",
		MarkdownDescription: "GithubConnector is the Schema for the githubconnectors API",
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
				Description:         "GithubConnector resource definition v3 from Teleport",
				MarkdownDescription: "GithubConnector resource definition v3 from Teleport",
				Attributes: map[string]schema.Attribute{
					"api_endpoint_url": schema.StringAttribute{
						Description:         "APIEndpointURL is the URL of the API endpoint of the Github instance this connector is for.",
						MarkdownDescription: "APIEndpointURL is the URL of the API endpoint of the Github instance this connector is for.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"client_id": schema.StringAttribute{
						Description:         "ClientID is the Github OAuth app client ID.",
						MarkdownDescription: "ClientID is the Github OAuth app client ID.",
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

							"insecure_allowed_cidr_ranges": schema.ListAttribute{
								Description:         "a list of CIDRs allowed for HTTP or HTTPS client redirect URLs",
								MarkdownDescription: "a list of CIDRs allowed for HTTP or HTTPS client redirect URLs",
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

					"client_secret": schema.StringAttribute{
						Description:         "ClientSecret is the Github OAuth app client secret. This field supports secret lookup. See the operator documentation for more details.",
						MarkdownDescription: "ClientSecret is the Github OAuth app client secret. This field supports secret lookup. See the operator documentation for more details.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"display": schema.StringAttribute{
						Description:         "Display is the connector display name.",
						MarkdownDescription: "Display is the connector display name.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"endpoint_url": schema.StringAttribute{
						Description:         "EndpointURL is the URL of the GitHub instance this connector is for.",
						MarkdownDescription: "EndpointURL is the URL of the GitHub instance this connector is for.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"redirect_url": schema.StringAttribute{
						Description:         "RedirectURL is the authorization callback URL.",
						MarkdownDescription: "RedirectURL is the authorization callback URL.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"teams_to_roles": schema.ListNestedAttribute{
						Description:         "TeamsToRoles maps Github team memberships onto allowed roles.",
						MarkdownDescription: "TeamsToRoles maps Github team memberships onto allowed roles.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"organization": schema.StringAttribute{
									Description:         "Organization is a Github organization a user belongs to.",
									MarkdownDescription: "Organization is a Github organization a user belongs to.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"roles": schema.ListAttribute{
									Description:         "Roles is a list of allowed logins for this org/team.",
									MarkdownDescription: "Roles is a list of allowed logins for this org/team.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"team": schema.StringAttribute{
									Description:         "Team is a team within the organization a user belongs to.",
									MarkdownDescription: "Team is a team within the organization a user belongs to.",
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *ResourcesTeleportDevTeleportGithubConnectorV3Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_resources_teleport_dev_teleport_github_connector_v3_manifest")

	var model ResourcesTeleportDevTeleportGithubConnectorV3ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("resources.teleport.dev/v3")
	model.Kind = pointer.String("TeleportGithubConnector")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
