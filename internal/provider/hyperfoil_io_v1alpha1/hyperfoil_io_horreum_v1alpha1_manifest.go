/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package hyperfoil_io_v1alpha1

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &HyperfoilIoHorreumV1Alpha1Manifest{}
)

func NewHyperfoilIoHorreumV1Alpha1Manifest() datasource.DataSource {
	return &HyperfoilIoHorreumV1Alpha1Manifest{}
}

type HyperfoilIoHorreumV1Alpha1Manifest struct{}

type HyperfoilIoHorreumV1Alpha1ManifestData struct {
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
		AdminSecret *string `tfsdk:"admin_secret" json:"adminSecret,omitempty"`
		Database    *struct {
			Host   *string `tfsdk:"host" json:"host,omitempty"`
			Name   *string `tfsdk:"name" json:"name,omitempty"`
			Port   *int64  `tfsdk:"port" json:"port,omitempty"`
			Secret *string `tfsdk:"secret" json:"secret,omitempty"`
		} `tfsdk:"database" json:"database,omitempty"`
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

func (r *HyperfoilIoHorreumV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_hyperfoil_io_horreum_v1alpha1_manifest"
}

func (r *HyperfoilIoHorreumV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "HorreumSpec defines the desired state of Horreum",
				MarkdownDescription: "HorreumSpec defines the desired state of Horreum",
				Attributes: map[string]schema.Attribute{
					"admin_secret": schema.StringAttribute{
						Description:         "Name of secret resource with data 'username' and 'password'. This will be the first user that get's created in Horreum with the 'admin' role, therefore it can create other users and teams. Created automatically if it does not exist.",
						MarkdownDescription: "Name of secret resource with data 'username' and 'password'. This will be the first user that get's created in Horreum with the 'admin' role, therefore it can create other users and teams. Created automatically if it does not exist.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"database": schema.SingleNestedAttribute{
						Description:         "Database coordinates for Horreum data. Besides 'username' and 'password' the secret must also contain key 'dbsecret' that will be used to sign access to the database.",
						MarkdownDescription: "Database coordinates for Horreum data. Besides 'username' and 'password' the secret must also contain key 'dbsecret' that will be used to sign access to the database.",
						Attributes: map[string]schema.Attribute{
							"host": schema.StringAttribute{
								Description:         "Hostname for the database",
								MarkdownDescription: "Hostname for the database",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the database",
								MarkdownDescription: "Name of the database",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "Database port; defaults to 5432",
								MarkdownDescription: "Database port; defaults to 5432",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secret": schema.StringAttribute{
								Description:         "Name of secret resource with data 'username' and 'password'. Created if does not exist.",
								MarkdownDescription: "Name of secret resource with data 'username' and 'password'. Created if does not exist.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"image": schema.StringAttribute{
						Description:         "Horreum image. Defaults to quay.io/hyperfoil/horreum:latest",
						MarkdownDescription: "Horreum image. Defaults to quay.io/hyperfoil/horreum:latest",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"keycloak": schema.SingleNestedAttribute{
						Description:         "Keycloak specification",
						MarkdownDescription: "Keycloak specification",
						Attributes: map[string]schema.Attribute{
							"admin_secret": schema.StringAttribute{
								Description:         "Secret used for admin access to the deployed Keycloak instance. Created if does not exist. Must contain keys 'username' and 'password'.",
								MarkdownDescription: "Secret used for admin access to the deployed Keycloak instance. Created if does not exist. Must contain keys 'username' and 'password'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"database": schema.SingleNestedAttribute{
								Description:         "Database coordinates Keycloak should use",
								MarkdownDescription: "Database coordinates Keycloak should use",
								Attributes: map[string]schema.Attribute{
									"host": schema.StringAttribute{
										Description:         "Hostname for the database",
										MarkdownDescription: "Hostname for the database",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name of the database",
										MarkdownDescription: "Name of the database",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"port": schema.Int64Attribute{
										Description:         "Database port; defaults to 5432",
										MarkdownDescription: "Database port; defaults to 5432",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"secret": schema.StringAttribute{
										Description:         "Name of secret resource with data 'username' and 'password'. Created if does not exist.",
										MarkdownDescription: "Name of secret resource with data 'username' and 'password'. Created if does not exist.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"external": schema.SingleNestedAttribute{
								Description:         "When this is set Keycloak instance will not be deployed and Horreum will use this external instance.",
								MarkdownDescription: "When this is set Keycloak instance will not be deployed and Horreum will use this external instance.",
								Attributes: map[string]schema.Attribute{
									"internal_uri": schema.StringAttribute{
										Description:         "Internal URI - Horreum will use this for communication but won't disclose that.",
										MarkdownDescription: "Internal URI - Horreum will use this for communication but won't disclose that.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"public_uri": schema.StringAttribute{
										Description:         "Public facing URI - Horreum will send this URI to the clients.",
										MarkdownDescription: "Public facing URI - Horreum will send this URI to the clients.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"image": schema.StringAttribute{
								Description:         "Image that should be used for Keycloak deployment. Defaults to quay.io/keycloak/keycloak:latest",
								MarkdownDescription: "Image that should be used for Keycloak deployment. Defaults to quay.io/keycloak/keycloak:latest",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"route": schema.SingleNestedAttribute{
								Description:         "Route for external access to the Keycloak instance.",
								MarkdownDescription: "Route for external access to the Keycloak instance.",
								Attributes: map[string]schema.Attribute{
									"host": schema.StringAttribute{
										Description:         "Host for the route leading to Controller REST endpoint. Example: horreum.apps.mycloud.example.com",
										MarkdownDescription: "Host for the route leading to Controller REST endpoint. Example: horreum.apps.mycloud.example.com",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tls": schema.StringAttribute{
										Description:         "Optional for edge and reencrypt routes, required for passthrough; Name of the secret hosting 'tls.crt', 'tls.key' and optionally 'ca.crt'",
										MarkdownDescription: "Optional for edge and reencrypt routes, required for passthrough; Name of the secret hosting 'tls.crt', 'tls.key' and optionally 'ca.crt'",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"type": schema.StringAttribute{
										Description:         "Either 'http' (for plain-text routes - not recommended), 'edge', 'reencrypt' or 'passthrough'",
										MarkdownDescription: "Either 'http' (for plain-text routes - not recommended), 'edge', 'reencrypt' or 'passthrough'",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"service_type": schema.StringAttribute{
								Description:         "Alternative service type when routes are not available (e.g. on vanilla K8s)",
								MarkdownDescription: "Alternative service type when routes are not available (e.g. on vanilla K8s)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"node_host": schema.StringAttribute{
						Description:         "Host used for NodePort services",
						MarkdownDescription: "Host used for NodePort services",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"postgres": schema.SingleNestedAttribute{
						Description:         "PostgreSQL specification",
						MarkdownDescription: "PostgreSQL specification",
						Attributes: map[string]schema.Attribute{
							"admin_secret": schema.StringAttribute{
								Description:         "Secret used for unrestricted access to the database. Created if does not exist. Must contain keys 'username' and 'password'.",
								MarkdownDescription: "Secret used for unrestricted access to the database. Created if does not exist. Must contain keys 'username' and 'password'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enabled": schema.BoolAttribute{
								Description:         "True (or omitted) to deploy PostgreSQL database",
								MarkdownDescription: "True (or omitted) to deploy PostgreSQL database",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image": schema.StringAttribute{
								Description:         "Image used for PostgreSQL deployment. Defaults to registry.redhat.io/rhel8/postgresql-12:latest",
								MarkdownDescription: "Image used for PostgreSQL deployment. Defaults to registry.redhat.io/rhel8/postgresql-12:latest",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"persistent_volume_claim": schema.StringAttribute{
								Description:         "Name of PVC where the database will store the data. If empty, ephemeral storage will be used.",
								MarkdownDescription: "Name of PVC where the database will store the data. If empty, ephemeral storage will be used.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"user": schema.Int64Attribute{
								Description:         "Id of the user the container should run as",
								MarkdownDescription: "Id of the user the container should run as",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"route": schema.SingleNestedAttribute{
						Description:         "Route for external access",
						MarkdownDescription: "Route for external access",
						Attributes: map[string]schema.Attribute{
							"host": schema.StringAttribute{
								Description:         "Host for the route leading to Controller REST endpoint. Example: horreum.apps.mycloud.example.com",
								MarkdownDescription: "Host for the route leading to Controller REST endpoint. Example: horreum.apps.mycloud.example.com",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tls": schema.StringAttribute{
								Description:         "Optional for edge and reencrypt routes, required for passthrough; Name of the secret hosting 'tls.crt', 'tls.key' and optionally 'ca.crt'",
								MarkdownDescription: "Optional for edge and reencrypt routes, required for passthrough; Name of the secret hosting 'tls.crt', 'tls.key' and optionally 'ca.crt'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "Either 'http' (for plain-text routes - not recommended), 'edge', 'reencrypt' or 'passthrough'",
								MarkdownDescription: "Either 'http' (for plain-text routes - not recommended), 'edge', 'reencrypt' or 'passthrough'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"service_type": schema.StringAttribute{
						Description:         "Alternative service type when routes are not available (e.g. on vanilla K8s)",
						MarkdownDescription: "Alternative service type when routes are not available (e.g. on vanilla K8s)",
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

func (r *HyperfoilIoHorreumV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_hyperfoil_io_horreum_v1alpha1_manifest")

	var model HyperfoilIoHorreumV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("hyperfoil.io/v1alpha1")
	model.Kind = pointer.String("Horreum")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"YAML Error: "+err.Error(),
		)
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
