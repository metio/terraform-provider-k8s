/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package mirrors_kts_studio_v1alpha2

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
	_ resource.Resource                = &MirrorsKtsStudioSecretMirrorV1Alpha2Resource{}
	_ resource.ResourceWithConfigure   = &MirrorsKtsStudioSecretMirrorV1Alpha2Resource{}
	_ resource.ResourceWithImportState = &MirrorsKtsStudioSecretMirrorV1Alpha2Resource{}
)

func NewMirrorsKtsStudioSecretMirrorV1Alpha2Resource() resource.Resource {
	return &MirrorsKtsStudioSecretMirrorV1Alpha2Resource{}
}

type MirrorsKtsStudioSecretMirrorV1Alpha2Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type MirrorsKtsStudioSecretMirrorV1Alpha2ResourceData struct {
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
		DeletePolicy *string `tfsdk:"delete_policy" json:"deletePolicy,omitempty"`
		Destination  *struct {
			Namespaces *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
			Type       *string   `tfsdk:"type" json:"type,omitempty"`
			Vault      *struct {
				Addr *string `tfsdk:"addr" json:"addr,omitempty"`
				Auth *struct {
					Approle *struct {
						AppRolePath *string `tfsdk:"app_role_path" json:"appRolePath,omitempty"`
						RoleIDKey   *string `tfsdk:"role_id_key" json:"roleIDKey,omitempty"`
						SecretIDKey *string `tfsdk:"secret_id_key" json:"secretIDKey,omitempty"`
						SecretRef   *struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
					} `tfsdk:"approle" json:"approle,omitempty"`
					Token *struct {
						SecretRef *struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
						TokenKey *string `tfsdk:"token_key" json:"tokenKey,omitempty"`
					} `tfsdk:"token" json:"token,omitempty"`
				} `tfsdk:"auth" json:"auth,omitempty"`
				Path *string `tfsdk:"path" json:"path,omitempty"`
			} `tfsdk:"vault" json:"vault,omitempty"`
		} `tfsdk:"destination" json:"destination,omitempty"`
		PollPeriodSeconds *int64 `tfsdk:"poll_period_seconds" json:"pollPeriodSeconds,omitempty"`
		Source            *struct {
			Name  *string `tfsdk:"name" json:"name,omitempty"`
			Type  *string `tfsdk:"type" json:"type,omitempty"`
			Vault *struct {
				Addr *string `tfsdk:"addr" json:"addr,omitempty"`
				Auth *struct {
					Approle *struct {
						AppRolePath *string `tfsdk:"app_role_path" json:"appRolePath,omitempty"`
						RoleIDKey   *string `tfsdk:"role_id_key" json:"roleIDKey,omitempty"`
						SecretIDKey *string `tfsdk:"secret_id_key" json:"secretIDKey,omitempty"`
						SecretRef   *struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
					} `tfsdk:"approle" json:"approle,omitempty"`
					Token *struct {
						SecretRef *struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
						TokenKey *string `tfsdk:"token_key" json:"tokenKey,omitempty"`
					} `tfsdk:"token" json:"token,omitempty"`
				} `tfsdk:"auth" json:"auth,omitempty"`
				Path *string `tfsdk:"path" json:"path,omitempty"`
			} `tfsdk:"vault" json:"vault,omitempty"`
		} `tfsdk:"source" json:"source,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *MirrorsKtsStudioSecretMirrorV1Alpha2Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_mirrors_kts_studio_secret_mirror_v1alpha2"
}

func (r *MirrorsKtsStudioSecretMirrorV1Alpha2Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "SecretMirror is the Schema for the secretmirrors API",
		MarkdownDescription: "SecretMirror is the Schema for the secretmirrors API",
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
				Description:         "SecretMirrorSpec defines the desired behaviour of Secret mirroring",
				MarkdownDescription: "SecretMirrorSpec defines the desired behaviour of Secret mirroring",
				Attributes: map[string]schema.Attribute{
					"delete_policy": schema.StringAttribute{
						Description:         "What to do with Secret objects created by a SecretMirror. Two policies exist – delete (deletes all created secrets) and retain (leaves them in the cluster). Default: delete",
						MarkdownDescription: "What to do with Secret objects created by a SecretMirror. Two policies exist – delete (deletes all created secrets) and retain (leaves them in the cluster). Default: delete",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("delete", "retain"),
						},
					},

					"destination": schema.SingleNestedAttribute{
						Description:         "SecretMirrorDestination defines where to sync a secret data to",
						MarkdownDescription: "SecretMirrorDestination defines where to sync a secret data to",
						Attributes: map[string]schema.Attribute{
							"namespaces": schema.ListAttribute{
								Description:         "An array of regular expressions to match namespaces where to copy a source secret",
								MarkdownDescription: "An array of regular expressions to match namespaces where to copy a source secret",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "Destination type. Possible values — namespaces, vault. Default: namespaces",
								MarkdownDescription: "Destination type. Possible values — namespaces, vault. Default: namespaces",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("namespaces", "vault"),
								},
							},

							"vault": schema.SingleNestedAttribute{
								Description:         "VaultSpec contains information of secret location",
								MarkdownDescription: "VaultSpec contains information of secret location",
								Attributes: map[string]schema.Attribute{
									"addr": schema.StringAttribute{
										Description:         "Addr specifies a Vault endpoint URL (e.g. https://vault.example.com)",
										MarkdownDescription: "Addr specifies a Vault endpoint URL (e.g. https://vault.example.com)",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"auth": schema.SingleNestedAttribute{
										Description:         "VaultAuthSpec describes how to authenticate against a Vault server",
										MarkdownDescription: "VaultAuthSpec describes how to authenticate against a Vault server",
										Attributes: map[string]schema.Attribute{
											"approle": schema.SingleNestedAttribute{
												Description:         "VaultAppRoleAuthSpec specifies approle-specific auth data",
												MarkdownDescription: "VaultAppRoleAuthSpec specifies approle-specific auth data",
												Attributes: map[string]schema.Attribute{
													"app_role_path": schema.StringAttribute{
														Description:         "approle Vault prefix. Default: approle",
														MarkdownDescription: "approle Vault prefix. Default: approle",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"role_id_key": schema.StringAttribute{
														Description:         "A key in the SecretRef which contains role-id value. Default: role-id",
														MarkdownDescription: "A key in the SecretRef which contains role-id value. Default: role-id",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"secret_id_key": schema.StringAttribute{
														Description:         "A key in the SecretRef which contains secret-id value. Default: secret-id",
														MarkdownDescription: "A key in the SecretRef which contains secret-id value. Default: secret-id",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"secret_ref": schema.SingleNestedAttribute{
														Description:         "Reference to a Secret containing role-id and secret-id",
														MarkdownDescription: "Reference to a Secret containing role-id and secret-id",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name is unique within a namespace to reference a secret resource.",
																MarkdownDescription: "Name is unique within a namespace to reference a secret resource.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace defines the space within which the secret name must be unique.",
																MarkdownDescription: "Namespace defines the space within which the secret name must be unique.",
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
												Required: false,
												Optional: true,
												Computed: false,
											},

											"token": schema.SingleNestedAttribute{
												Description:         "VaultTokenAuthSpec specifies token-specific auth data",
												MarkdownDescription: "VaultTokenAuthSpec specifies token-specific auth data",
												Attributes: map[string]schema.Attribute{
													"secret_ref": schema.SingleNestedAttribute{
														Description:         "Reference to a Secret containing token",
														MarkdownDescription: "Reference to a Secret containing token",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name is unique within a namespace to reference a secret resource.",
																MarkdownDescription: "Name is unique within a namespace to reference a secret resource.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace defines the space within which the secret name must be unique.",
																MarkdownDescription: "Namespace defines the space within which the secret name must be unique.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"token_key": schema.StringAttribute{
														Description:         "A key in the SecretRef which contains token value. Default: token",
														MarkdownDescription: "A key in the SecretRef which contains token value. Default: token",
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
										Required: false,
										Optional: true,
										Computed: false,
									},

									"path": schema.StringAttribute{
										Description:         "Path specifies a vault secret path (e.g. secret/data/some-secret or mongodb/creds/mymongo)",
										MarkdownDescription: "Path specifies a vault secret path (e.g. secret/data/some-secret or mongodb/creds/mymongo)",
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"poll_period_seconds": schema.Int64Attribute{
						Description:         "How often to check for secret changes. Default: 180 seconds",
						MarkdownDescription: "How often to check for secret changes. Default: 180 seconds",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"source": schema.SingleNestedAttribute{
						Description:         "SecretMirrorSource defines where to extract a secret data from",
						MarkdownDescription: "SecretMirrorSource defines where to extract a secret data from",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("secret", "vault"),
								},
							},

							"vault": schema.SingleNestedAttribute{
								Description:         "VaultSpec contains information of secret location",
								MarkdownDescription: "VaultSpec contains information of secret location",
								Attributes: map[string]schema.Attribute{
									"addr": schema.StringAttribute{
										Description:         "Addr specifies a Vault endpoint URL (e.g. https://vault.example.com)",
										MarkdownDescription: "Addr specifies a Vault endpoint URL (e.g. https://vault.example.com)",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"auth": schema.SingleNestedAttribute{
										Description:         "VaultAuthSpec describes how to authenticate against a Vault server",
										MarkdownDescription: "VaultAuthSpec describes how to authenticate against a Vault server",
										Attributes: map[string]schema.Attribute{
											"approle": schema.SingleNestedAttribute{
												Description:         "VaultAppRoleAuthSpec specifies approle-specific auth data",
												MarkdownDescription: "VaultAppRoleAuthSpec specifies approle-specific auth data",
												Attributes: map[string]schema.Attribute{
													"app_role_path": schema.StringAttribute{
														Description:         "approle Vault prefix. Default: approle",
														MarkdownDescription: "approle Vault prefix. Default: approle",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"role_id_key": schema.StringAttribute{
														Description:         "A key in the SecretRef which contains role-id value. Default: role-id",
														MarkdownDescription: "A key in the SecretRef which contains role-id value. Default: role-id",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"secret_id_key": schema.StringAttribute{
														Description:         "A key in the SecretRef which contains secret-id value. Default: secret-id",
														MarkdownDescription: "A key in the SecretRef which contains secret-id value. Default: secret-id",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"secret_ref": schema.SingleNestedAttribute{
														Description:         "Reference to a Secret containing role-id and secret-id",
														MarkdownDescription: "Reference to a Secret containing role-id and secret-id",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name is unique within a namespace to reference a secret resource.",
																MarkdownDescription: "Name is unique within a namespace to reference a secret resource.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace defines the space within which the secret name must be unique.",
																MarkdownDescription: "Namespace defines the space within which the secret name must be unique.",
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
												Required: false,
												Optional: true,
												Computed: false,
											},

											"token": schema.SingleNestedAttribute{
												Description:         "VaultTokenAuthSpec specifies token-specific auth data",
												MarkdownDescription: "VaultTokenAuthSpec specifies token-specific auth data",
												Attributes: map[string]schema.Attribute{
													"secret_ref": schema.SingleNestedAttribute{
														Description:         "Reference to a Secret containing token",
														MarkdownDescription: "Reference to a Secret containing token",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name is unique within a namespace to reference a secret resource.",
																MarkdownDescription: "Name is unique within a namespace to reference a secret resource.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace defines the space within which the secret name must be unique.",
																MarkdownDescription: "Namespace defines the space within which the secret name must be unique.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"token_key": schema.StringAttribute{
														Description:         "A key in the SecretRef which contains token value. Default: token",
														MarkdownDescription: "A key in the SecretRef which contains token value. Default: token",
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
										Required: false,
										Optional: true,
										Computed: false,
									},

									"path": schema.StringAttribute{
										Description:         "Path specifies a vault secret path (e.g. secret/data/some-secret or mongodb/creds/mymongo)",
										MarkdownDescription: "Path specifies a vault secret path (e.g. secret/data/some-secret or mongodb/creds/mymongo)",
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

func (r *MirrorsKtsStudioSecretMirrorV1Alpha2Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
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

func (r *MirrorsKtsStudioSecretMirrorV1Alpha2Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_mirrors_kts_studio_secret_mirror_v1alpha2")

	var model MirrorsKtsStudioSecretMirrorV1Alpha2ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("mirrors.kts.studio/v1alpha2")
	model.Kind = pointer.String("SecretMirror")

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
		Resource(k8sSchema.GroupVersionResource{Group: "mirrors.kts.studio", Version: "v1alpha2", Resource: "secretmirrors"}).
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

	var readResponse MirrorsKtsStudioSecretMirrorV1Alpha2ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *MirrorsKtsStudioSecretMirrorV1Alpha2Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_mirrors_kts_studio_secret_mirror_v1alpha2")

	var data MirrorsKtsStudioSecretMirrorV1Alpha2ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "mirrors.kts.studio", Version: "v1alpha2", Resource: "secretmirrors"}).
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

	var readResponse MirrorsKtsStudioSecretMirrorV1Alpha2ResourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *MirrorsKtsStudioSecretMirrorV1Alpha2Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_mirrors_kts_studio_secret_mirror_v1alpha2")

	var model MirrorsKtsStudioSecretMirrorV1Alpha2ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("mirrors.kts.studio/v1alpha2")
	model.Kind = pointer.String("SecretMirror")

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
		Resource(k8sSchema.GroupVersionResource{Group: "mirrors.kts.studio", Version: "v1alpha2", Resource: "secretmirrors"}).
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

	var readResponse MirrorsKtsStudioSecretMirrorV1Alpha2ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *MirrorsKtsStudioSecretMirrorV1Alpha2Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_mirrors_kts_studio_secret_mirror_v1alpha2")

	var data MirrorsKtsStudioSecretMirrorV1Alpha2ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "mirrors.kts.studio", Version: "v1alpha2", Resource: "secretmirrors"}).
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
				Resource(k8sSchema.GroupVersionResource{Group: "mirrors.kts.studio", Version: "v1alpha2", Resource: "secretmirrors"}).
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

func (r *MirrorsKtsStudioSecretMirrorV1Alpha2Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
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
