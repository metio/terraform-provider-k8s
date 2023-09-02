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
	_ datasource.DataSource              = &ResourcesTeleportDevTeleportGithubConnectorV3DataSource{}
	_ datasource.DataSourceWithConfigure = &ResourcesTeleportDevTeleportGithubConnectorV3DataSource{}
)

func NewResourcesTeleportDevTeleportGithubConnectorV3DataSource() datasource.DataSource {
	return &ResourcesTeleportDevTeleportGithubConnectorV3DataSource{}
}

type ResourcesTeleportDevTeleportGithubConnectorV3DataSource struct {
	kubernetesClient dynamic.Interface
}

type ResourcesTeleportDevTeleportGithubConnectorV3DataSourceData struct {
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
		Api_endpoint_url *string `tfsdk:"api_endpoint_url" json:"api_endpoint_url,omitempty"`
		Client_id        *string `tfsdk:"client_id" json:"client_id,omitempty"`
		Client_secret    *string `tfsdk:"client_secret" json:"client_secret,omitempty"`
		Display          *string `tfsdk:"display" json:"display,omitempty"`
		Endpoint_url     *string `tfsdk:"endpoint_url" json:"endpoint_url,omitempty"`
		Redirect_url     *string `tfsdk:"redirect_url" json:"redirect_url,omitempty"`
		Teams_to_roles   *[]struct {
			Organization *string   `tfsdk:"organization" json:"organization,omitempty"`
			Roles        *[]string `tfsdk:"roles" json:"roles,omitempty"`
			Team         *string   `tfsdk:"team" json:"team,omitempty"`
		} `tfsdk:"teams_to_roles" json:"teams_to_roles,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ResourcesTeleportDevTeleportGithubConnectorV3DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_resources_teleport_dev_teleport_github_connector_v3"
}

func (r *ResourcesTeleportDevTeleportGithubConnectorV3DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "GithubConnector is the Schema for the githubconnectors API",
		MarkdownDescription: "GithubConnector is the Schema for the githubconnectors API",
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
				Description:         "GithubConnector resource definition v3 from Teleport",
				MarkdownDescription: "GithubConnector resource definition v3 from Teleport",
				Attributes: map[string]schema.Attribute{
					"api_endpoint_url": schema.StringAttribute{
						Description:         "APIEndpointURL is the URL of the API endpoint of the Github instance this connector is for.",
						MarkdownDescription: "APIEndpointURL is the URL of the API endpoint of the Github instance this connector is for.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"client_id": schema.StringAttribute{
						Description:         "ClientID is the Github OAuth app client ID.",
						MarkdownDescription: "ClientID is the Github OAuth app client ID.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"client_secret": schema.StringAttribute{
						Description:         "ClientSecret is the Github OAuth app client secret.",
						MarkdownDescription: "ClientSecret is the Github OAuth app client secret.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"display": schema.StringAttribute{
						Description:         "Display is the connector display name.",
						MarkdownDescription: "Display is the connector display name.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"endpoint_url": schema.StringAttribute{
						Description:         "EndpointURL is the URL of the GitHub instance this connector is for.",
						MarkdownDescription: "EndpointURL is the URL of the GitHub instance this connector is for.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"redirect_url": schema.StringAttribute{
						Description:         "RedirectURL is the authorization callback URL.",
						MarkdownDescription: "RedirectURL is the authorization callback URL.",
						Required:            false,
						Optional:            false,
						Computed:            true,
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
									Optional:            false,
									Computed:            true,
								},

								"roles": schema.ListAttribute{
									Description:         "Roles is a list of allowed logins for this org/team.",
									MarkdownDescription: "Roles is a list of allowed logins for this org/team.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"team": schema.StringAttribute{
									Description:         "Team is a team within the organization a user belongs to.",
									MarkdownDescription: "Team is a team within the organization a user belongs to.",
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
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *ResourcesTeleportDevTeleportGithubConnectorV3DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *ResourcesTeleportDevTeleportGithubConnectorV3DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_resources_teleport_dev_teleport_github_connector_v3")

	var data ResourcesTeleportDevTeleportGithubConnectorV3DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "resources.teleport.dev", Version: "v3", Resource: "TeleportGithubConnector"}).
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

	var readResponse ResourcesTeleportDevTeleportGithubConnectorV3DataSourceData
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
	data.Kind = pointer.String("TeleportGithubConnector")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
