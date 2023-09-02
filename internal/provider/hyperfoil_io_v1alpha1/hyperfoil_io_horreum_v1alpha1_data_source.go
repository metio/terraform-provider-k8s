/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package hyperfoil_io_v1alpha1

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
	_ datasource.DataSource              = &HyperfoilIoHorreumV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &HyperfoilIoHorreumV1Alpha1DataSource{}
)

func NewHyperfoilIoHorreumV1Alpha1DataSource() datasource.DataSource {
	return &HyperfoilIoHorreumV1Alpha1DataSource{}
}

type HyperfoilIoHorreumV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type HyperfoilIoHorreumV1Alpha1DataSourceData struct {
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
		AdminSecret *string `tfsdk:"admin_secret" json:"adminSecret,omitempty"`
		Database    *struct {
			Host   *string `tfsdk:"host" json:"host,omitempty"`
			Name   *string `tfsdk:"name" json:"name,omitempty"`
			Port   *int64  `tfsdk:"port" json:"port,omitempty"`
			Secret *string `tfsdk:"secret" json:"secret,omitempty"`
		} `tfsdk:"database" json:"database,omitempty"`
		Grafana *struct {
			AdminSecret *string `tfsdk:"admin_secret" json:"adminSecret,omitempty"`
			External    *struct {
				InternalUri *string `tfsdk:"internal_uri" json:"internalUri,omitempty"`
				PublicUri   *string `tfsdk:"public_uri" json:"publicUri,omitempty"`
			} `tfsdk:"external" json:"external,omitempty"`
			Image *string `tfsdk:"image" json:"image,omitempty"`
			Route *struct {
				Host *string `tfsdk:"host" json:"host,omitempty"`
				Tls  *string `tfsdk:"tls" json:"tls,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"route" json:"route,omitempty"`
			ServiceType *string `tfsdk:"service_type" json:"serviceType,omitempty"`
			Theme       *string `tfsdk:"theme" json:"theme,omitempty"`
		} `tfsdk:"grafana" json:"grafana,omitempty"`
		Image    *string `tfsdk:"image" json:"image,omitempty"`
		Keycloak *struct {
			AdminSecret *string `tfsdk:"admin_secret" json:"adminSecret,omitempty"`
			Database    *struct {
				Host   *string `tfsdk:"host" json:"host,omitempty"`
				Name   *string `tfsdk:"name" json:"name,omitempty"`
				Port   *int64  `tfsdk:"port" json:"port,omitempty"`
				Secret *string `tfsdk:"secret" json:"secret,omitempty"`
			} `tfsdk:"database" json:"database,omitempty"`
			External *struct {
				InternalUri *string `tfsdk:"internal_uri" json:"internalUri,omitempty"`
				PublicUri   *string `tfsdk:"public_uri" json:"publicUri,omitempty"`
			} `tfsdk:"external" json:"external,omitempty"`
			Image *string `tfsdk:"image" json:"image,omitempty"`
			Route *struct {
				Host *string `tfsdk:"host" json:"host,omitempty"`
				Tls  *string `tfsdk:"tls" json:"tls,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"route" json:"route,omitempty"`
			ServiceType *string `tfsdk:"service_type" json:"serviceType,omitempty"`
		} `tfsdk:"keycloak" json:"keycloak,omitempty"`
		NodeHost *string `tfsdk:"node_host" json:"nodeHost,omitempty"`
		Postgres *struct {
			AdminSecret           *string `tfsdk:"admin_secret" json:"adminSecret,omitempty"`
			Enabled               *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
			Image                 *string `tfsdk:"image" json:"image,omitempty"`
			PersistentVolumeClaim *string `tfsdk:"persistent_volume_claim" json:"persistentVolumeClaim,omitempty"`
			User                  *int64  `tfsdk:"user" json:"user,omitempty"`
		} `tfsdk:"postgres" json:"postgres,omitempty"`
		Route *struct {
			Host *string `tfsdk:"host" json:"host,omitempty"`
			Tls  *string `tfsdk:"tls" json:"tls,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"route" json:"route,omitempty"`
		ServiceType *string `tfsdk:"service_type" json:"serviceType,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *HyperfoilIoHorreumV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_hyperfoil_io_horreum_v1alpha1"
}

func (r *HyperfoilIoHorreumV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Horreum is the object configuring Horreum performance results repository",
		MarkdownDescription: "Horreum is the object configuring Horreum performance results repository",
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
				Description:         "HorreumSpec defines the desired state of Horreum",
				MarkdownDescription: "HorreumSpec defines the desired state of Horreum",
				Attributes: map[string]schema.Attribute{
					"admin_secret": schema.StringAttribute{
						Description:         "Name of secret resource with data 'username' and 'password'. This will be the first user that get's created in Horreum with the 'admin' role, therefore it can create other users and teams. Created automatically if it does not exist.",
						MarkdownDescription: "Name of secret resource with data 'username' and 'password'. This will be the first user that get's created in Horreum with the 'admin' role, therefore it can create other users and teams. Created automatically if it does not exist.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"database": schema.SingleNestedAttribute{
						Description:         "Database coordinates for Horreum data. Besides 'username' and 'password' the secret must also contain key 'dbsecret' that will be used to sign access to the database.",
						MarkdownDescription: "Database coordinates for Horreum data. Besides 'username' and 'password' the secret must also contain key 'dbsecret' that will be used to sign access to the database.",
						Attributes: map[string]schema.Attribute{
							"host": schema.StringAttribute{
								Description:         "Hostname for the database",
								MarkdownDescription: "Hostname for the database",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the database",
								MarkdownDescription: "Name of the database",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"port": schema.Int64Attribute{
								Description:         "Database port; defaults to 5432",
								MarkdownDescription: "Database port; defaults to 5432",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"secret": schema.StringAttribute{
								Description:         "Name of secret resource with data 'username' and 'password'. Created if does not exist.",
								MarkdownDescription: "Name of secret resource with data 'username' and 'password'. Created if does not exist.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"grafana": schema.SingleNestedAttribute{
						Description:         "Grafana specification",
						MarkdownDescription: "Grafana specification",
						Attributes: map[string]schema.Attribute{
							"admin_secret": schema.StringAttribute{
								Description:         "Secret used for admin access to Grafana. Created if it does not exist. Must contain keys 'username' and 'password'.",
								MarkdownDescription: "Secret used for admin access to Grafana. Created if it does not exist. Must contain keys 'username' and 'password'.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"external": schema.SingleNestedAttribute{
								Description:         "When this is set Grafana instance will not be deployed and Horreum will use this external instance.",
								MarkdownDescription: "When this is set Grafana instance will not be deployed and Horreum will use this external instance.",
								Attributes: map[string]schema.Attribute{
									"internal_uri": schema.StringAttribute{
										Description:         "Internal URI - Horreum will use this for communication but won't disclose that.",
										MarkdownDescription: "Internal URI - Horreum will use this for communication but won't disclose that.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"public_uri": schema.StringAttribute{
										Description:         "Public facing URI - Horreum will send this URI to the clients.",
										MarkdownDescription: "Public facing URI - Horreum will send this URI to the clients.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"image": schema.StringAttribute{
								Description:         "Custom Grafana image. Defaults to registry.redhat.io/openshift4/ose-grafana:latest",
								MarkdownDescription: "Custom Grafana image. Defaults to registry.redhat.io/openshift4/ose-grafana:latest",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"route": schema.SingleNestedAttribute{
								Description:         "Route for external access.",
								MarkdownDescription: "Route for external access.",
								Attributes: map[string]schema.Attribute{
									"host": schema.StringAttribute{
										Description:         "Host for the route leading to Controller REST endpoint. Example: horreum.apps.mycloud.example.com",
										MarkdownDescription: "Host for the route leading to Controller REST endpoint. Example: horreum.apps.mycloud.example.com",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"tls": schema.StringAttribute{
										Description:         "Optional for edge and reencrypt routes, required for passthrough; Name of the secret hosting 'tls.crt', 'tls.key' and optionally 'ca.crt'",
										MarkdownDescription: "Optional for edge and reencrypt routes, required for passthrough; Name of the secret hosting 'tls.crt', 'tls.key' and optionally 'ca.crt'",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"type": schema.StringAttribute{
										Description:         "Either 'http' (for plain-text routes - not recommended), 'edge', 'reencrypt' or 'passthrough'",
										MarkdownDescription: "Either 'http' (for plain-text routes - not recommended), 'edge', 'reencrypt' or 'passthrough'",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"service_type": schema.StringAttribute{
								Description:         "Alternative service type when routes are not available (e.g. on vanilla K8s)",
								MarkdownDescription: "Alternative service type when routes are not available (e.g. on vanilla K8s)",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"theme": schema.StringAttribute{
								Description:         "Default theme that should be used - one of 'dark' or 'light'. Defaults to 'light'.",
								MarkdownDescription: "Default theme that should be used - one of 'dark' or 'light'. Defaults to 'light'.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"image": schema.StringAttribute{
						Description:         "Horreum image. Defaults to quay.io/hyperfoil/horreum:latest",
						MarkdownDescription: "Horreum image. Defaults to quay.io/hyperfoil/horreum:latest",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"keycloak": schema.SingleNestedAttribute{
						Description:         "Keycloak specification",
						MarkdownDescription: "Keycloak specification",
						Attributes: map[string]schema.Attribute{
							"admin_secret": schema.StringAttribute{
								Description:         "Secret used for admin access to the deployed Keycloak instance. Created if does not exist. Must contain keys 'username' and 'password'.",
								MarkdownDescription: "Secret used for admin access to the deployed Keycloak instance. Created if does not exist. Must contain keys 'username' and 'password'.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"database": schema.SingleNestedAttribute{
								Description:         "Database coordinates Keycloak should use",
								MarkdownDescription: "Database coordinates Keycloak should use",
								Attributes: map[string]schema.Attribute{
									"host": schema.StringAttribute{
										Description:         "Hostname for the database",
										MarkdownDescription: "Hostname for the database",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"name": schema.StringAttribute{
										Description:         "Name of the database",
										MarkdownDescription: "Name of the database",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"port": schema.Int64Attribute{
										Description:         "Database port; defaults to 5432",
										MarkdownDescription: "Database port; defaults to 5432",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"secret": schema.StringAttribute{
										Description:         "Name of secret resource with data 'username' and 'password'. Created if does not exist.",
										MarkdownDescription: "Name of secret resource with data 'username' and 'password'. Created if does not exist.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"external": schema.SingleNestedAttribute{
								Description:         "When this is set Keycloak instance will not be deployed and Horreum will use this external instance.",
								MarkdownDescription: "When this is set Keycloak instance will not be deployed and Horreum will use this external instance.",
								Attributes: map[string]schema.Attribute{
									"internal_uri": schema.StringAttribute{
										Description:         "Internal URI - Horreum will use this for communication but won't disclose that.",
										MarkdownDescription: "Internal URI - Horreum will use this for communication but won't disclose that.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"public_uri": schema.StringAttribute{
										Description:         "Public facing URI - Horreum will send this URI to the clients.",
										MarkdownDescription: "Public facing URI - Horreum will send this URI to the clients.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"image": schema.StringAttribute{
								Description:         "Image that should be used for Keycloak deployment. Defaults to quay.io/keycloak/keycloak:latest",
								MarkdownDescription: "Image that should be used for Keycloak deployment. Defaults to quay.io/keycloak/keycloak:latest",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"route": schema.SingleNestedAttribute{
								Description:         "Route for external access to the Keycloak instance.",
								MarkdownDescription: "Route for external access to the Keycloak instance.",
								Attributes: map[string]schema.Attribute{
									"host": schema.StringAttribute{
										Description:         "Host for the route leading to Controller REST endpoint. Example: horreum.apps.mycloud.example.com",
										MarkdownDescription: "Host for the route leading to Controller REST endpoint. Example: horreum.apps.mycloud.example.com",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"tls": schema.StringAttribute{
										Description:         "Optional for edge and reencrypt routes, required for passthrough; Name of the secret hosting 'tls.crt', 'tls.key' and optionally 'ca.crt'",
										MarkdownDescription: "Optional for edge and reencrypt routes, required for passthrough; Name of the secret hosting 'tls.crt', 'tls.key' and optionally 'ca.crt'",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"type": schema.StringAttribute{
										Description:         "Either 'http' (for plain-text routes - not recommended), 'edge', 'reencrypt' or 'passthrough'",
										MarkdownDescription: "Either 'http' (for plain-text routes - not recommended), 'edge', 'reencrypt' or 'passthrough'",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"service_type": schema.StringAttribute{
								Description:         "Alternative service type when routes are not available (e.g. on vanilla K8s)",
								MarkdownDescription: "Alternative service type when routes are not available (e.g. on vanilla K8s)",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"node_host": schema.StringAttribute{
						Description:         "Host used for NodePort services",
						MarkdownDescription: "Host used for NodePort services",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"postgres": schema.SingleNestedAttribute{
						Description:         "PostgreSQL specification",
						MarkdownDescription: "PostgreSQL specification",
						Attributes: map[string]schema.Attribute{
							"admin_secret": schema.StringAttribute{
								Description:         "Secret used for unrestricted access to the database. Created if does not exist. Must contain keys 'username' and 'password'.",
								MarkdownDescription: "Secret used for unrestricted access to the database. Created if does not exist. Must contain keys 'username' and 'password'.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"enabled": schema.BoolAttribute{
								Description:         "True (or omitted) to deploy PostgreSQL database",
								MarkdownDescription: "True (or omitted) to deploy PostgreSQL database",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"image": schema.StringAttribute{
								Description:         "Image used for PostgreSQL deployment. Defaults to registry.redhat.io/rhel8/postgresql-12:latest",
								MarkdownDescription: "Image used for PostgreSQL deployment. Defaults to registry.redhat.io/rhel8/postgresql-12:latest",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"persistent_volume_claim": schema.StringAttribute{
								Description:         "Name of PVC where the database will store the data. If empty, ephemeral storage will be used.",
								MarkdownDescription: "Name of PVC where the database will store the data. If empty, ephemeral storage will be used.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"user": schema.Int64Attribute{
								Description:         "Id of the user the container should run as",
								MarkdownDescription: "Id of the user the container should run as",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"route": schema.SingleNestedAttribute{
						Description:         "Route for external access",
						MarkdownDescription: "Route for external access",
						Attributes: map[string]schema.Attribute{
							"host": schema.StringAttribute{
								Description:         "Host for the route leading to Controller REST endpoint. Example: horreum.apps.mycloud.example.com",
								MarkdownDescription: "Host for the route leading to Controller REST endpoint. Example: horreum.apps.mycloud.example.com",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"tls": schema.StringAttribute{
								Description:         "Optional for edge and reencrypt routes, required for passthrough; Name of the secret hosting 'tls.crt', 'tls.key' and optionally 'ca.crt'",
								MarkdownDescription: "Optional for edge and reencrypt routes, required for passthrough; Name of the secret hosting 'tls.crt', 'tls.key' and optionally 'ca.crt'",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"type": schema.StringAttribute{
								Description:         "Either 'http' (for plain-text routes - not recommended), 'edge', 'reencrypt' or 'passthrough'",
								MarkdownDescription: "Either 'http' (for plain-text routes - not recommended), 'edge', 'reencrypt' or 'passthrough'",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"service_type": schema.StringAttribute{
						Description:         "Alternative service type when routes are not available (e.g. on vanilla K8s)",
						MarkdownDescription: "Alternative service type when routes are not available (e.g. on vanilla K8s)",
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

func (r *HyperfoilIoHorreumV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *HyperfoilIoHorreumV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_hyperfoil_io_horreum_v1alpha1")

	var data HyperfoilIoHorreumV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "hyperfoil.io", Version: "v1alpha1", Resource: "Horreum"}).
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

	var readResponse HyperfoilIoHorreumV1Alpha1DataSourceData
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
	data.ApiVersion = pointer.String("hyperfoil.io/v1alpha1")
	data.Kind = pointer.String("Horreum")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
