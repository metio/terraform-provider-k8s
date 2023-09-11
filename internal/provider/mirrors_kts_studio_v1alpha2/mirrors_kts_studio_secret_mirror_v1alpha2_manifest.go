/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package mirrors_kts_studio_v1alpha2

import (
	"context"
	"fmt"
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
	_ datasource.DataSource = &MirrorsKtsStudioSecretMirrorV1Alpha2Manifest{}
)

func NewMirrorsKtsStudioSecretMirrorV1Alpha2Manifest() datasource.DataSource {
	return &MirrorsKtsStudioSecretMirrorV1Alpha2Manifest{}
}

type MirrorsKtsStudioSecretMirrorV1Alpha2Manifest struct{}

type MirrorsKtsStudioSecretMirrorV1Alpha2ManifestData struct {
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

func (r *MirrorsKtsStudioSecretMirrorV1Alpha2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_mirrors_kts_studio_secret_mirror_v1alpha2_manifest"
}

func (r *MirrorsKtsStudioSecretMirrorV1Alpha2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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

func (r *MirrorsKtsStudioSecretMirrorV1Alpha2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_mirrors_kts_studio_secret_mirror_v1alpha2_manifest")

	var model MirrorsKtsStudioSecretMirrorV1Alpha2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("mirrors.kts.studio/v1alpha2")
	model.Kind = pointer.String("SecretMirror")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
