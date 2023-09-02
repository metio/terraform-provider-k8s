/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package keda_sh_v1alpha1

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
	_ datasource.DataSource              = &KedaShClusterTriggerAuthenticationV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &KedaShClusterTriggerAuthenticationV1Alpha1DataSource{}
)

func NewKedaShClusterTriggerAuthenticationV1Alpha1DataSource() datasource.DataSource {
	return &KedaShClusterTriggerAuthenticationV1Alpha1DataSource{}
}

type KedaShClusterTriggerAuthenticationV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type KedaShClusterTriggerAuthenticationV1Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		AzureKeyVault *struct {
			Cloud *struct {
				ActiveDirectoryEndpoint *string `tfsdk:"active_directory_endpoint" json:"activeDirectoryEndpoint,omitempty"`
				KeyVaultResourceURL     *string `tfsdk:"key_vault_resource_url" json:"keyVaultResourceURL,omitempty"`
				Type                    *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"cloud" json:"cloud,omitempty"`
			Credentials *struct {
				ClientId     *string `tfsdk:"client_id" json:"clientId,omitempty"`
				ClientSecret *struct {
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"client_secret" json:"clientSecret,omitempty"`
				TenantId *string `tfsdk:"tenant_id" json:"tenantId,omitempty"`
			} `tfsdk:"credentials" json:"credentials,omitempty"`
			PodIdentity *struct {
				IdentityId *string `tfsdk:"identity_id" json:"identityId,omitempty"`
				Provider   *string `tfsdk:"provider" json:"provider,omitempty"`
			} `tfsdk:"pod_identity" json:"podIdentity,omitempty"`
			Secrets *[]struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Parameter *string `tfsdk:"parameter" json:"parameter,omitempty"`
				Version   *string `tfsdk:"version" json:"version,omitempty"`
			} `tfsdk:"secrets" json:"secrets,omitempty"`
			VaultUri *string `tfsdk:"vault_uri" json:"vaultUri,omitempty"`
		} `tfsdk:"azure_key_vault" json:"azureKeyVault,omitempty"`
		Env *[]struct {
			ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
			Name          *string `tfsdk:"name" json:"name,omitempty"`
			Parameter     *string `tfsdk:"parameter" json:"parameter,omitempty"`
		} `tfsdk:"env" json:"env,omitempty"`
		HashiCorpVault *struct {
			Address        *string `tfsdk:"address" json:"address,omitempty"`
			Authentication *string `tfsdk:"authentication" json:"authentication,omitempty"`
			Credential     *struct {
				ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				Token          *string `tfsdk:"token" json:"token,omitempty"`
			} `tfsdk:"credential" json:"credential,omitempty"`
			Mount     *string `tfsdk:"mount" json:"mount,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Role      *string `tfsdk:"role" json:"role,omitempty"`
			Secrets   *[]struct {
				Key       *string `tfsdk:"key" json:"key,omitempty"`
				Parameter *string `tfsdk:"parameter" json:"parameter,omitempty"`
				Path      *string `tfsdk:"path" json:"path,omitempty"`
			} `tfsdk:"secrets" json:"secrets,omitempty"`
		} `tfsdk:"hashi_corp_vault" json:"hashiCorpVault,omitempty"`
		PodIdentity *struct {
			IdentityId *string `tfsdk:"identity_id" json:"identityId,omitempty"`
			Provider   *string `tfsdk:"provider" json:"provider,omitempty"`
		} `tfsdk:"pod_identity" json:"podIdentity,omitempty"`
		SecretTargetRef *[]struct {
			Key       *string `tfsdk:"key" json:"key,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Parameter *string `tfsdk:"parameter" json:"parameter,omitempty"`
		} `tfsdk:"secret_target_ref" json:"secretTargetRef,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KedaShClusterTriggerAuthenticationV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_keda_sh_cluster_trigger_authentication_v1alpha1"
}

func (r *KedaShClusterTriggerAuthenticationV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ClusterTriggerAuthentication defines how a trigger can authenticate globally",
		MarkdownDescription: "ClusterTriggerAuthentication defines how a trigger can authenticate globally",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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
				Description:         "TriggerAuthenticationSpec defines the various ways to authenticate",
				MarkdownDescription: "TriggerAuthenticationSpec defines the various ways to authenticate",
				Attributes: map[string]schema.Attribute{
					"azure_key_vault": schema.SingleNestedAttribute{
						Description:         "AzureKeyVault is used to authenticate using Azure Key Vault",
						MarkdownDescription: "AzureKeyVault is used to authenticate using Azure Key Vault",
						Attributes: map[string]schema.Attribute{
							"cloud": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"active_directory_endpoint": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"key_vault_resource_url": schema.StringAttribute{
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
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"credentials": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"client_id": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"client_secret": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value_from": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"secret_key_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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

									"tenant_id": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"pod_identity": schema.SingleNestedAttribute{
								Description:         "AuthPodIdentity allows users to select the platform native identity mechanism",
								MarkdownDescription: "AuthPodIdentity allows users to select the platform native identity mechanism",
								Attributes: map[string]schema.Attribute{
									"identity_id": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"provider": schema.StringAttribute{
										Description:         "PodIdentityProvider contains the list of providers",
										MarkdownDescription: "PodIdentityProvider contains the list of providers",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"secrets": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"parameter": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"version": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
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

							"vault_uri": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"env": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"container_name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"parameter": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"hashi_corp_vault": schema.SingleNestedAttribute{
						Description:         "HashiCorpVault is used to authenticate using Hashicorp Vault",
						MarkdownDescription: "HashiCorpVault is used to authenticate using Hashicorp Vault",
						Attributes: map[string]schema.Attribute{
							"address": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"authentication": schema.StringAttribute{
								Description:         "VaultAuthentication contains the list of Hashicorp Vault authentication methods",
								MarkdownDescription: "VaultAuthentication contains the list of Hashicorp Vault authentication methods",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"credential": schema.SingleNestedAttribute{
								Description:         "Credential defines the Hashicorp Vault credentials depending on the authentication method",
								MarkdownDescription: "Credential defines the Hashicorp Vault credentials depending on the authentication method",
								Attributes: map[string]schema.Attribute{
									"service_account": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"token": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"mount": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"namespace": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"role": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"secrets": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"parameter": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
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

					"pod_identity": schema.SingleNestedAttribute{
						Description:         "AuthPodIdentity allows users to select the platform native identity mechanism",
						MarkdownDescription: "AuthPodIdentity allows users to select the platform native identity mechanism",
						Attributes: map[string]schema.Attribute{
							"identity_id": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"provider": schema.StringAttribute{
								Description:         "PodIdentityProvider contains the list of providers",
								MarkdownDescription: "PodIdentityProvider contains the list of providers",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"secret_target_ref": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"key": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"parameter": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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

func (r *KedaShClusterTriggerAuthenticationV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *KedaShClusterTriggerAuthenticationV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_keda_sh_cluster_trigger_authentication_v1alpha1")

	var data KedaShClusterTriggerAuthenticationV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "keda.sh", Version: "v1alpha1", Resource: "ClusterTriggerAuthentication"}).
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

	var readResponse KedaShClusterTriggerAuthenticationV1Alpha1DataSourceData
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

	data.ID = types.StringValue(data.Metadata.Name)
	data.ApiVersion = pointer.String("keda.sh/v1alpha1")
	data.Kind = pointer.String("ClusterTriggerAuthentication")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
