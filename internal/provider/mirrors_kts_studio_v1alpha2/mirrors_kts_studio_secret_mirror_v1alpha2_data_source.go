/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package mirrors_kts_studio_v1alpha2

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
	_ datasource.DataSource              = &MirrorsKtsStudioSecretMirrorV1Alpha2DataSource{}
	_ datasource.DataSourceWithConfigure = &MirrorsKtsStudioSecretMirrorV1Alpha2DataSource{}
)

func NewMirrorsKtsStudioSecretMirrorV1Alpha2DataSource() datasource.DataSource {
	return &MirrorsKtsStudioSecretMirrorV1Alpha2DataSource{}
}

type MirrorsKtsStudioSecretMirrorV1Alpha2DataSource struct {
	kubernetesClient dynamic.Interface
}

type MirrorsKtsStudioSecretMirrorV1Alpha2DataSourceData struct {
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

func (r *MirrorsKtsStudioSecretMirrorV1Alpha2DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_mirrors_kts_studio_secret_mirror_v1alpha2"
}

func (r *MirrorsKtsStudioSecretMirrorV1Alpha2DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "SecretMirrorSpec defines the desired behaviour of Secret mirroring",
				MarkdownDescription: "SecretMirrorSpec defines the desired behaviour of Secret mirroring",
				Attributes: map[string]schema.Attribute{
					"delete_policy": schema.StringAttribute{
						Description:         "What to do with Secret objects created by a SecretMirror. Two policies exist – delete (deletes all created secrets) and retain (leaves them in the cluster). Default: delete",
						MarkdownDescription: "What to do with Secret objects created by a SecretMirror. Two policies exist – delete (deletes all created secrets) and retain (leaves them in the cluster). Default: delete",
						Required:            false,
						Optional:            false,
						Computed:            true,
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
								Optional:            false,
								Computed:            true,
							},

							"type": schema.StringAttribute{
								Description:         "Destination type. Possible values — namespaces, vault. Default: namespaces",
								MarkdownDescription: "Destination type. Possible values — namespaces, vault. Default: namespaces",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"vault": schema.SingleNestedAttribute{
								Description:         "VaultSpec contains information of secret location",
								MarkdownDescription: "VaultSpec contains information of secret location",
								Attributes: map[string]schema.Attribute{
									"addr": schema.StringAttribute{
										Description:         "Addr specifies a Vault endpoint URL (e.g. https://vault.example.com)",
										MarkdownDescription: "Addr specifies a Vault endpoint URL (e.g. https://vault.example.com)",
										Required:            false,
										Optional:            false,
										Computed:            true,
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
														Optional:            false,
														Computed:            true,
													},

													"role_id_key": schema.StringAttribute{
														Description:         "A key in the SecretRef which contains role-id value. Default: role-id",
														MarkdownDescription: "A key in the SecretRef which contains role-id value. Default: role-id",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"secret_id_key": schema.StringAttribute{
														Description:         "A key in the SecretRef which contains secret-id value. Default: secret-id",
														MarkdownDescription: "A key in the SecretRef which contains secret-id value. Default: secret-id",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"secret_ref": schema.SingleNestedAttribute{
														Description:         "Reference to a Secret containing role-id and secret-id",
														MarkdownDescription: "Reference to a Secret containing role-id and secret-id",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name is unique within a namespace to reference a secret resource.",
																MarkdownDescription: "Name is unique within a namespace to reference a secret resource.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace defines the space within which the secret name must be unique.",
																MarkdownDescription: "Namespace defines the space within which the secret name must be unique.",
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
												Required: false,
												Optional: false,
												Computed: true,
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
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace defines the space within which the secret name must be unique.",
																MarkdownDescription: "Namespace defines the space within which the secret name must be unique.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"token_key": schema.StringAttribute{
														Description:         "A key in the SecretRef which contains token value. Default: token",
														MarkdownDescription: "A key in the SecretRef which contains token value. Default: token",
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
										Required: false,
										Optional: false,
										Computed: true,
									},

									"path": schema.StringAttribute{
										Description:         "Path specifies a vault secret path (e.g. secret/data/some-secret or mongodb/creds/mymongo)",
										MarkdownDescription: "Path specifies a vault secret path (e.g. secret/data/some-secret or mongodb/creds/mymongo)",
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
						Required: false,
						Optional: false,
						Computed: true,
					},

					"poll_period_seconds": schema.Int64Attribute{
						Description:         "How often to check for secret changes. Default: 180 seconds",
						MarkdownDescription: "How often to check for secret changes. Default: 180 seconds",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"source": schema.SingleNestedAttribute{
						Description:         "SecretMirrorSource defines where to extract a secret data from",
						MarkdownDescription: "SecretMirrorSource defines where to extract a secret data from",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"type": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"vault": schema.SingleNestedAttribute{
								Description:         "VaultSpec contains information of secret location",
								MarkdownDescription: "VaultSpec contains information of secret location",
								Attributes: map[string]schema.Attribute{
									"addr": schema.StringAttribute{
										Description:         "Addr specifies a Vault endpoint URL (e.g. https://vault.example.com)",
										MarkdownDescription: "Addr specifies a Vault endpoint URL (e.g. https://vault.example.com)",
										Required:            false,
										Optional:            false,
										Computed:            true,
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
														Optional:            false,
														Computed:            true,
													},

													"role_id_key": schema.StringAttribute{
														Description:         "A key in the SecretRef which contains role-id value. Default: role-id",
														MarkdownDescription: "A key in the SecretRef which contains role-id value. Default: role-id",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"secret_id_key": schema.StringAttribute{
														Description:         "A key in the SecretRef which contains secret-id value. Default: secret-id",
														MarkdownDescription: "A key in the SecretRef which contains secret-id value. Default: secret-id",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"secret_ref": schema.SingleNestedAttribute{
														Description:         "Reference to a Secret containing role-id and secret-id",
														MarkdownDescription: "Reference to a Secret containing role-id and secret-id",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name is unique within a namespace to reference a secret resource.",
																MarkdownDescription: "Name is unique within a namespace to reference a secret resource.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace defines the space within which the secret name must be unique.",
																MarkdownDescription: "Namespace defines the space within which the secret name must be unique.",
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
												Required: false,
												Optional: false,
												Computed: true,
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
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace defines the space within which the secret name must be unique.",
																MarkdownDescription: "Namespace defines the space within which the secret name must be unique.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"token_key": schema.StringAttribute{
														Description:         "A key in the SecretRef which contains token value. Default: token",
														MarkdownDescription: "A key in the SecretRef which contains token value. Default: token",
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
										Required: false,
										Optional: false,
										Computed: true,
									},

									"path": schema.StringAttribute{
										Description:         "Path specifies a vault secret path (e.g. secret/data/some-secret or mongodb/creds/mymongo)",
										MarkdownDescription: "Path specifies a vault secret path (e.g. secret/data/some-secret or mongodb/creds/mymongo)",
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

func (r *MirrorsKtsStudioSecretMirrorV1Alpha2DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *MirrorsKtsStudioSecretMirrorV1Alpha2DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_mirrors_kts_studio_secret_mirror_v1alpha2")

	var data MirrorsKtsStudioSecretMirrorV1Alpha2DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "mirrors.kts.studio", Version: "v1alpha2", Resource: "secretmirrors"}).
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

	var readResponse MirrorsKtsStudioSecretMirrorV1Alpha2DataSourceData
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
	data.ApiVersion = pointer.String("mirrors.kts.studio/v1alpha2")
	data.Kind = pointer.String("SecretMirror")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
