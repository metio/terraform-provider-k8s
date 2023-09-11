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
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	k8sTypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"strings"
	"time"
)

var (
	_ resource.Resource                = &HyperfoilIoHorreumV1Alpha1Resource{}
	_ resource.ResourceWithConfigure   = &HyperfoilIoHorreumV1Alpha1Resource{}
	_ resource.ResourceWithImportState = &HyperfoilIoHorreumV1Alpha1Resource{}
)

func NewHyperfoilIoHorreumV1Alpha1Resource() resource.Resource {
	return &HyperfoilIoHorreumV1Alpha1Resource{}
}

type HyperfoilIoHorreumV1Alpha1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type HyperfoilIoHorreumV1Alpha1ResourceData struct {
	ID             types.String `tfsdk:"id" json:"-"`
	ForceConflicts types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager   types.String `tfsdk:"field_manager" json:"-"`
	WaitForUpsert  types.List   `tfsdk:"wait_for_upsert" json:"-"`
	WaitForDelete  types.Object `tfsdk:"wait_for_delete" json:"-"`

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

func (r *HyperfoilIoHorreumV1Alpha1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_hyperfoil_io_horreum_v1alpha1"
}

func (r *HyperfoilIoHorreumV1Alpha1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
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

			"force_conflicts": schema.BoolAttribute{
				Description:         "If 'true', server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "If `true`, server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"field_manager": schema.BoolAttribute{
				Description:         "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"wait_for_upsert": schema.ListNestedAttribute{
				Description:         "Wait for specific conditions after create/update of resources.",
				MarkdownDescription: "Wait for specific conditions after create/update of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"jsonpath": schema.StringAttribute{
							Description:         "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							MarkdownDescription: "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							Required:            true,
							Optional:            false,
							Computed:            false,
						},
						"value": schema.StringAttribute{
							Description:         "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							MarkdownDescription: "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							Required:            false,
							Optional:            true,
							Computed:            true,
						},
						"timeout": schema.StringAttribute{
							Description:         "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							MarkdownDescription: "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("30s"),
						},
						"poll_interval": schema.StringAttribute{
							Description:         "The length of time to wait before checking again.",
							MarkdownDescription: "The length of time to wait before checking again.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("5s"),
						},
					},
				},
			},

			"wait_for_delete": schema.SingleNestedAttribute{
				Description:         "Wait for deletion of resources.",
				MarkdownDescription: "Wait for deletion of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"timeout": schema.StringAttribute{
						Description:         "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
						MarkdownDescription: "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             stringdefault.StaticString("30s"),
					},
					"poll_interval": schema.StringAttribute{
						Description:         "The length of time to wait before checking again.",
						MarkdownDescription: "The length of time to wait before checking again.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             stringdefault.StaticString("5s"),
					},
				},
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},

					"labels": schema.MapAttribute{
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            true,
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
						Computed:            true,
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

func (r *HyperfoilIoHorreumV1Alpha1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedResourceDataError(request.ProviderData))
	}
}

func (r *HyperfoilIoHorreumV1Alpha1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_hyperfoil_io_horreum_v1alpha1")

	var model HyperfoilIoHorreumV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("hyperfoil.io/v1alpha1")
	model.Kind = pointer.String("Horreum")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonMarshalError(err))
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "hyperfoil.io", Version: "v1alpha1", Resource: "horreums"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.Append(utilities.PatchError(err))
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse HyperfoilIoHorreumV1Alpha1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *HyperfoilIoHorreumV1Alpha1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_hyperfoil_io_horreum_v1alpha1")

	var data HyperfoilIoHorreumV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "hyperfoil.io", Version: "v1alpha1", Resource: "horreums"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetNamespacedResourceError(err, data.Metadata.Name, data.Metadata.Namespace))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse HyperfoilIoHorreumV1Alpha1ResourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *HyperfoilIoHorreumV1Alpha1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_hyperfoil_io_horreum_v1alpha1")

	var model HyperfoilIoHorreumV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("hyperfoil.io/v1alpha1")
	model.Kind = pointer.String("Horreum")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonMarshalError(err))
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "hyperfoil.io", Version: "v1alpha1", Resource: "horreums"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.Append(utilities.PatchError(err))
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse HyperfoilIoHorreumV1Alpha1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *HyperfoilIoHorreumV1Alpha1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_hyperfoil_io_horreum_v1alpha1")

	var data HyperfoilIoHorreumV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "hyperfoil.io", Version: "v1alpha1", Resource: "horreums"}).
		Namespace(data.Metadata.Namespace).
		Delete(ctx, data.Metadata.Name, meta.DeleteOptions{})
	if utilities.IsDeletionError(err) {
		response.Diagnostics.Append(utilities.DeleteError(err))
		return
	}

	if !data.WaitForDelete.IsNull() {
		timeout := utilities.DetermineTimeout(data.WaitForDelete.Attributes())
		pollInterval := utilities.DeterminePollInterval(data.WaitForDelete.Attributes())

		startTime := time.Now()
		for {
			_, err := r.kubernetesClient.
				Resource(k8sSchema.GroupVersionResource{Group: "hyperfoil.io", Version: "v1alpha1", Resource: "horreums"}).
				Namespace(data.Metadata.Namespace).
				Get(ctx, data.Metadata.Name, meta.GetOptions{})
			if utilities.IsNotFound(err) || timeout == time.Second*0 {
				break
			}
			if time.Now().After(startTime.Add(timeout)) {
				response.Diagnostics.Append(utilities.WaitTimeoutExceeded())
				return
			}
			time.Sleep(pollInterval)
		}
	}
}

func (r *HyperfoilIoHorreumV1Alpha1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	idParts := strings.Split(request.ID, "/")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'namespace/name' Got: '%q'", request.ID),
		)
		return
	}

	namespace := idParts[0]
	name := idParts[1]
	tflog.Trace(ctx, "parsed import ID", map[string]interface{}{
		"namespace": namespace,
		"name":      name,
	})
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("namespace"), namespace)...)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("name"), name)...)
}
