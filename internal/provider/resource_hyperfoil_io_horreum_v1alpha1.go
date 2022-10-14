/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"gopkg.in/yaml.v3"
	"time"
)

type HyperfoilIoHorreumV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*HyperfoilIoHorreumV1Alpha1Resource)(nil)
)

type HyperfoilIoHorreumV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type HyperfoilIoHorreumV1Alpha1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		AdminSecret *string `tfsdk:"admin_secret" yaml:"adminSecret,omitempty"`

		Database *struct {
			Host *string `tfsdk:"host" yaml:"host,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

			Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`
		} `tfsdk:"database" yaml:"database,omitempty"`

		Grafana *struct {
			AdminSecret *string `tfsdk:"admin_secret" yaml:"adminSecret,omitempty"`

			External *struct {
				InternalUri *string `tfsdk:"internal_uri" yaml:"internalUri,omitempty"`

				PublicUri *string `tfsdk:"public_uri" yaml:"publicUri,omitempty"`
			} `tfsdk:"external" yaml:"external,omitempty"`

			Image *string `tfsdk:"image" yaml:"image,omitempty"`

			Route *struct {
				Host *string `tfsdk:"host" yaml:"host,omitempty"`

				Tls *string `tfsdk:"tls" yaml:"tls,omitempty"`

				Type *string `tfsdk:"type" yaml:"type,omitempty"`
			} `tfsdk:"route" yaml:"route,omitempty"`

			ServiceType *string `tfsdk:"service_type" yaml:"serviceType,omitempty"`

			Theme *string `tfsdk:"theme" yaml:"theme,omitempty"`
		} `tfsdk:"grafana" yaml:"grafana,omitempty"`

		Image *string `tfsdk:"image" yaml:"image,omitempty"`

		Keycloak *struct {
			AdminSecret *string `tfsdk:"admin_secret" yaml:"adminSecret,omitempty"`

			Database *struct {
				Host *string `tfsdk:"host" yaml:"host,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

				Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`
			} `tfsdk:"database" yaml:"database,omitempty"`

			External *struct {
				InternalUri *string `tfsdk:"internal_uri" yaml:"internalUri,omitempty"`

				PublicUri *string `tfsdk:"public_uri" yaml:"publicUri,omitempty"`
			} `tfsdk:"external" yaml:"external,omitempty"`

			Image *string `tfsdk:"image" yaml:"image,omitempty"`

			Route *struct {
				Host *string `tfsdk:"host" yaml:"host,omitempty"`

				Tls *string `tfsdk:"tls" yaml:"tls,omitempty"`

				Type *string `tfsdk:"type" yaml:"type,omitempty"`
			} `tfsdk:"route" yaml:"route,omitempty"`

			ServiceType *string `tfsdk:"service_type" yaml:"serviceType,omitempty"`
		} `tfsdk:"keycloak" yaml:"keycloak,omitempty"`

		NodeHost *string `tfsdk:"node_host" yaml:"nodeHost,omitempty"`

		Postgres *struct {
			AdminSecret *string `tfsdk:"admin_secret" yaml:"adminSecret,omitempty"`

			Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

			Image *string `tfsdk:"image" yaml:"image,omitempty"`

			PersistentVolumeClaim *string `tfsdk:"persistent_volume_claim" yaml:"persistentVolumeClaim,omitempty"`

			User *int64 `tfsdk:"user" yaml:"user,omitempty"`
		} `tfsdk:"postgres" yaml:"postgres,omitempty"`

		Route *struct {
			Host *string `tfsdk:"host" yaml:"host,omitempty"`

			Tls *string `tfsdk:"tls" yaml:"tls,omitempty"`

			Type *string `tfsdk:"type" yaml:"type,omitempty"`
		} `tfsdk:"route" yaml:"route,omitempty"`

		ServiceType *string `tfsdk:"service_type" yaml:"serviceType,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewHyperfoilIoHorreumV1Alpha1Resource() resource.Resource {
	return &HyperfoilIoHorreumV1Alpha1Resource{}
}

func (r *HyperfoilIoHorreumV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_hyperfoil_io_horreum_v1alpha1"
}

func (r *HyperfoilIoHorreumV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Horreum is the object configuring Horreum performance results repository",
		MarkdownDescription: "Horreum is the object configuring Horreum performance results repository",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description:         "The timestamp of the last change to this resource.",
				MarkdownDescription: "The timestamp of the last change to this resource.",
				Type:                types.Int64Type,
				Computed:            true,
				Optional:            false,
			},

			"yaml": {
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"metadata": {
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Type:                types.StringType,
						Required:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
					},

					"namespace": {
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Type:                types.StringType,
						Optional:            true,
					},

					"labels": {
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.LabelValidator(),
						},
					},
					"annotations": {
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.AnnotationValidator(),
						},
					},
				}),
			},

			"api_version": {
				Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"kind": {
				Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"spec": {
				Description:         "HorreumSpec defines the desired state of Horreum",
				MarkdownDescription: "HorreumSpec defines the desired state of Horreum",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"admin_secret": {
						Description:         "Name of secret resource with data 'username' and 'password'. This will be the first user that get's created in Horreum with the 'admin' role, therefore it can create other users and teams. Created automatically if it does not exist.",
						MarkdownDescription: "Name of secret resource with data 'username' and 'password'. This will be the first user that get's created in Horreum with the 'admin' role, therefore it can create other users and teams. Created automatically if it does not exist.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"database": {
						Description:         "Database coordinates for Horreum data. Besides 'username' and 'password' the secret must also contain key 'dbsecret' that will be used to sign access to the database.",
						MarkdownDescription: "Database coordinates for Horreum data. Besides 'username' and 'password' the secret must also contain key 'dbsecret' that will be used to sign access to the database.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"host": {
								Description:         "Hostname for the database",
								MarkdownDescription: "Hostname for the database",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "Name of the database",
								MarkdownDescription: "Name of the database",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"port": {
								Description:         "Database port; defaults to 5432",
								MarkdownDescription: "Database port; defaults to 5432",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"secret": {
								Description:         "Name of secret resource with data 'username' and 'password'. Created if does not exist.",
								MarkdownDescription: "Name of secret resource with data 'username' and 'password'. Created if does not exist.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"grafana": {
						Description:         "Grafana specification",
						MarkdownDescription: "Grafana specification",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"admin_secret": {
								Description:         "Secret used for admin access to Grafana. Created if it does not exist. Must contain keys 'username' and 'password'.",
								MarkdownDescription: "Secret used for admin access to Grafana. Created if it does not exist. Must contain keys 'username' and 'password'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"external": {
								Description:         "When this is set Grafana instance will not be deployed and Horreum will use this external instance.",
								MarkdownDescription: "When this is set Grafana instance will not be deployed and Horreum will use this external instance.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"internal_uri": {
										Description:         "Internal URI - Horreum will use this for communication but won't disclose that.",
										MarkdownDescription: "Internal URI - Horreum will use this for communication but won't disclose that.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"public_uri": {
										Description:         "Public facing URI - Horreum will send this URI to the clients.",
										MarkdownDescription: "Public facing URI - Horreum will send this URI to the clients.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"image": {
								Description:         "Custom Grafana image. Defaults to registry.redhat.io/openshift4/ose-grafana:latest",
								MarkdownDescription: "Custom Grafana image. Defaults to registry.redhat.io/openshift4/ose-grafana:latest",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"route": {
								Description:         "Route for external access.",
								MarkdownDescription: "Route for external access.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"host": {
										Description:         "Host for the route leading to Controller REST endpoint. Example: horreum.apps.mycloud.example.com",
										MarkdownDescription: "Host for the route leading to Controller REST endpoint. Example: horreum.apps.mycloud.example.com",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls": {
										Description:         "Optional for edge and reencrypt routes, required for passthrough; Name of the secret hosting 'tls.crt', 'tls.key' and optionally 'ca.crt'",
										MarkdownDescription: "Optional for edge and reencrypt routes, required for passthrough; Name of the secret hosting 'tls.crt', 'tls.key' and optionally 'ca.crt'",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": {
										Description:         "Either 'http' (for plain-text routes - not recommended), 'edge', 'reencrypt' or 'passthrough'",
										MarkdownDescription: "Either 'http' (for plain-text routes - not recommended), 'edge', 'reencrypt' or 'passthrough'",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"service_type": {
								Description:         "Alternative service type when routes are not available (e.g. on vanilla K8s)",
								MarkdownDescription: "Alternative service type when routes are not available (e.g. on vanilla K8s)",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"theme": {
								Description:         "Default theme that should be used - one of 'dark' or 'light'. Defaults to 'light'.",
								MarkdownDescription: "Default theme that should be used - one of 'dark' or 'light'. Defaults to 'light'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"image": {
						Description:         "Horreum image. Defaults to quay.io/hyperfoil/horreum:latest",
						MarkdownDescription: "Horreum image. Defaults to quay.io/hyperfoil/horreum:latest",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"keycloak": {
						Description:         "Keycloak specification",
						MarkdownDescription: "Keycloak specification",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"admin_secret": {
								Description:         "Secret used for admin access to the deployed Keycloak instance. Created if does not exist. Must contain keys 'username' and 'password'.",
								MarkdownDescription: "Secret used for admin access to the deployed Keycloak instance. Created if does not exist. Must contain keys 'username' and 'password'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"database": {
								Description:         "Database coordinates Keycloak should use",
								MarkdownDescription: "Database coordinates Keycloak should use",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"host": {
										Description:         "Hostname for the database",
										MarkdownDescription: "Hostname for the database",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "Name of the database",
										MarkdownDescription: "Name of the database",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"port": {
										Description:         "Database port; defaults to 5432",
										MarkdownDescription: "Database port; defaults to 5432",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret": {
										Description:         "Name of secret resource with data 'username' and 'password'. Created if does not exist.",
										MarkdownDescription: "Name of secret resource with data 'username' and 'password'. Created if does not exist.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"external": {
								Description:         "When this is set Keycloak instance will not be deployed and Horreum will use this external instance.",
								MarkdownDescription: "When this is set Keycloak instance will not be deployed and Horreum will use this external instance.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"internal_uri": {
										Description:         "Internal URI - Horreum will use this for communication but won't disclose that.",
										MarkdownDescription: "Internal URI - Horreum will use this for communication but won't disclose that.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"public_uri": {
										Description:         "Public facing URI - Horreum will send this URI to the clients.",
										MarkdownDescription: "Public facing URI - Horreum will send this URI to the clients.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"image": {
								Description:         "Image that should be used for Keycloak deployment. Defaults to quay.io/keycloak/keycloak:latest",
								MarkdownDescription: "Image that should be used for Keycloak deployment. Defaults to quay.io/keycloak/keycloak:latest",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"route": {
								Description:         "Route for external access to the Keycloak instance.",
								MarkdownDescription: "Route for external access to the Keycloak instance.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"host": {
										Description:         "Host for the route leading to Controller REST endpoint. Example: horreum.apps.mycloud.example.com",
										MarkdownDescription: "Host for the route leading to Controller REST endpoint. Example: horreum.apps.mycloud.example.com",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls": {
										Description:         "Optional for edge and reencrypt routes, required for passthrough; Name of the secret hosting 'tls.crt', 'tls.key' and optionally 'ca.crt'",
										MarkdownDescription: "Optional for edge and reencrypt routes, required for passthrough; Name of the secret hosting 'tls.crt', 'tls.key' and optionally 'ca.crt'",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": {
										Description:         "Either 'http' (for plain-text routes - not recommended), 'edge', 'reencrypt' or 'passthrough'",
										MarkdownDescription: "Either 'http' (for plain-text routes - not recommended), 'edge', 'reencrypt' or 'passthrough'",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"service_type": {
								Description:         "Alternative service type when routes are not available (e.g. on vanilla K8s)",
								MarkdownDescription: "Alternative service type when routes are not available (e.g. on vanilla K8s)",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"node_host": {
						Description:         "Host used for NodePort services",
						MarkdownDescription: "Host used for NodePort services",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"postgres": {
						Description:         "PostgreSQL specification",
						MarkdownDescription: "PostgreSQL specification",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"admin_secret": {
								Description:         "Secret used for unrestricted access to the database. Created if does not exist. Must contain keys 'username' and 'password'.",
								MarkdownDescription: "Secret used for unrestricted access to the database. Created if does not exist. Must contain keys 'username' and 'password'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enabled": {
								Description:         "True (or omitted) to deploy PostgreSQL database",
								MarkdownDescription: "True (or omitted) to deploy PostgreSQL database",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"image": {
								Description:         "Image used for PostgreSQL deployment. Defaults to registry.redhat.io/rhel8/postgresql-12:latest",
								MarkdownDescription: "Image used for PostgreSQL deployment. Defaults to registry.redhat.io/rhel8/postgresql-12:latest",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"persistent_volume_claim": {
								Description:         "Name of PVC where the database will store the data. If empty, ephemeral storage will be used.",
								MarkdownDescription: "Name of PVC where the database will store the data. If empty, ephemeral storage will be used.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"user": {
								Description:         "Id of the user the container should run as",
								MarkdownDescription: "Id of the user the container should run as",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"route": {
						Description:         "Route for external access",
						MarkdownDescription: "Route for external access",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"host": {
								Description:         "Host for the route leading to Controller REST endpoint. Example: horreum.apps.mycloud.example.com",
								MarkdownDescription: "Host for the route leading to Controller REST endpoint. Example: horreum.apps.mycloud.example.com",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"tls": {
								Description:         "Optional for edge and reencrypt routes, required for passthrough; Name of the secret hosting 'tls.crt', 'tls.key' and optionally 'ca.crt'",
								MarkdownDescription: "Optional for edge and reencrypt routes, required for passthrough; Name of the secret hosting 'tls.crt', 'tls.key' and optionally 'ca.crt'",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"type": {
								Description:         "Either 'http' (for plain-text routes - not recommended), 'edge', 'reencrypt' or 'passthrough'",
								MarkdownDescription: "Either 'http' (for plain-text routes - not recommended), 'edge', 'reencrypt' or 'passthrough'",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"service_type": {
						Description:         "Alternative service type when routes are not available (e.g. on vanilla K8s)",
						MarkdownDescription: "Alternative service type when routes are not available (e.g. on vanilla K8s)",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *HyperfoilIoHorreumV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_hyperfoil_io_horreum_v1alpha1")

	var state HyperfoilIoHorreumV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel HyperfoilIoHorreumV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("hyperfoil.io/v1alpha1")
	goModel.Kind = utilities.Ptr("Horreum")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *HyperfoilIoHorreumV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_hyperfoil_io_horreum_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *HyperfoilIoHorreumV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_hyperfoil_io_horreum_v1alpha1")

	var state HyperfoilIoHorreumV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel HyperfoilIoHorreumV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("hyperfoil.io/v1alpha1")
	goModel.Kind = utilities.Ptr("Horreum")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *HyperfoilIoHorreumV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_hyperfoil_io_horreum_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
