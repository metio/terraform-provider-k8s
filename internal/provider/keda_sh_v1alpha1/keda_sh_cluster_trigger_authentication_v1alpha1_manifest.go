/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package keda_sh_v1alpha1

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
	_ datasource.DataSource = &KedaShClusterTriggerAuthenticationV1Alpha1Manifest{}
)

func NewKedaShClusterTriggerAuthenticationV1Alpha1Manifest() datasource.DataSource {
	return &KedaShClusterTriggerAuthenticationV1Alpha1Manifest{}
}

type KedaShClusterTriggerAuthenticationV1Alpha1Manifest struct{}

type KedaShClusterTriggerAuthenticationV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		AwsSecretManager *struct {
			Credentials *struct {
				AccessKey *struct {
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"access_key" json:"accessKey,omitempty"`
				AccessSecretKey *struct {
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"access_secret_key" json:"accessSecretKey,omitempty"`
				AccessToken *struct {
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"access_token" json:"accessToken,omitempty"`
			} `tfsdk:"credentials" json:"credentials,omitempty"`
			PodIdentity *struct {
				IdentityAuthorityHost *string `tfsdk:"identity_authority_host" json:"identityAuthorityHost,omitempty"`
				IdentityId            *string `tfsdk:"identity_id" json:"identityId,omitempty"`
				IdentityOwner         *string `tfsdk:"identity_owner" json:"identityOwner,omitempty"`
				IdentityTenantId      *string `tfsdk:"identity_tenant_id" json:"identityTenantId,omitempty"`
				Provider              *string `tfsdk:"provider" json:"provider,omitempty"`
				RoleArn               *string `tfsdk:"role_arn" json:"roleArn,omitempty"`
			} `tfsdk:"pod_identity" json:"podIdentity,omitempty"`
			Region  *string `tfsdk:"region" json:"region,omitempty"`
			Secrets *[]struct {
				Name         *string `tfsdk:"name" json:"name,omitempty"`
				Parameter    *string `tfsdk:"parameter" json:"parameter,omitempty"`
				VersionId    *string `tfsdk:"version_id" json:"versionId,omitempty"`
				VersionStage *string `tfsdk:"version_stage" json:"versionStage,omitempty"`
			} `tfsdk:"secrets" json:"secrets,omitempty"`
		} `tfsdk:"aws_secret_manager" json:"awsSecretManager,omitempty"`
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
				IdentityAuthorityHost *string `tfsdk:"identity_authority_host" json:"identityAuthorityHost,omitempty"`
				IdentityId            *string `tfsdk:"identity_id" json:"identityId,omitempty"`
				IdentityOwner         *string `tfsdk:"identity_owner" json:"identityOwner,omitempty"`
				IdentityTenantId      *string `tfsdk:"identity_tenant_id" json:"identityTenantId,omitempty"`
				Provider              *string `tfsdk:"provider" json:"provider,omitempty"`
				RoleArn               *string `tfsdk:"role_arn" json:"roleArn,omitempty"`
			} `tfsdk:"pod_identity" json:"podIdentity,omitempty"`
			Secrets *[]struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Parameter *string `tfsdk:"parameter" json:"parameter,omitempty"`
				Version   *string `tfsdk:"version" json:"version,omitempty"`
			} `tfsdk:"secrets" json:"secrets,omitempty"`
			VaultUri *string `tfsdk:"vault_uri" json:"vaultUri,omitempty"`
		} `tfsdk:"azure_key_vault" json:"azureKeyVault,omitempty"`
		ConfigMapTargetRef *[]struct {
			Key       *string `tfsdk:"key" json:"key,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Parameter *string `tfsdk:"parameter" json:"parameter,omitempty"`
		} `tfsdk:"config_map_target_ref" json:"configMapTargetRef,omitempty"`
		Env *[]struct {
			ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
			Name          *string `tfsdk:"name" json:"name,omitempty"`
			Parameter     *string `tfsdk:"parameter" json:"parameter,omitempty"`
		} `tfsdk:"env" json:"env,omitempty"`
		GcpSecretManager *struct {
			Credentials *struct {
				ClientSecret *struct {
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"client_secret" json:"clientSecret,omitempty"`
			} `tfsdk:"credentials" json:"credentials,omitempty"`
			PodIdentity *struct {
				IdentityAuthorityHost *string `tfsdk:"identity_authority_host" json:"identityAuthorityHost,omitempty"`
				IdentityId            *string `tfsdk:"identity_id" json:"identityId,omitempty"`
				IdentityOwner         *string `tfsdk:"identity_owner" json:"identityOwner,omitempty"`
				IdentityTenantId      *string `tfsdk:"identity_tenant_id" json:"identityTenantId,omitempty"`
				Provider              *string `tfsdk:"provider" json:"provider,omitempty"`
				RoleArn               *string `tfsdk:"role_arn" json:"roleArn,omitempty"`
			} `tfsdk:"pod_identity" json:"podIdentity,omitempty"`
			Secrets *[]struct {
				Id        *string `tfsdk:"id" json:"id,omitempty"`
				Parameter *string `tfsdk:"parameter" json:"parameter,omitempty"`
				Version   *string `tfsdk:"version" json:"version,omitempty"`
			} `tfsdk:"secrets" json:"secrets,omitempty"`
		} `tfsdk:"gcp_secret_manager" json:"gcpSecretManager,omitempty"`
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
				PkiData   *struct {
					AltNames   *string `tfsdk:"alt_names" json:"altNames,omitempty"`
					CommonName *string `tfsdk:"common_name" json:"commonName,omitempty"`
					Format     *string `tfsdk:"format" json:"format,omitempty"`
					IpSans     *string `tfsdk:"ip_sans" json:"ipSans,omitempty"`
					OtherSans  *string `tfsdk:"other_sans" json:"otherSans,omitempty"`
					Ttl        *string `tfsdk:"ttl" json:"ttl,omitempty"`
					UriSans    *string `tfsdk:"uri_sans" json:"uriSans,omitempty"`
				} `tfsdk:"pki_data" json:"pkiData,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"secrets" json:"secrets,omitempty"`
		} `tfsdk:"hashi_corp_vault" json:"hashiCorpVault,omitempty"`
		PodIdentity *struct {
			IdentityAuthorityHost *string `tfsdk:"identity_authority_host" json:"identityAuthorityHost,omitempty"`
			IdentityId            *string `tfsdk:"identity_id" json:"identityId,omitempty"`
			IdentityOwner         *string `tfsdk:"identity_owner" json:"identityOwner,omitempty"`
			IdentityTenantId      *string `tfsdk:"identity_tenant_id" json:"identityTenantId,omitempty"`
			Provider              *string `tfsdk:"provider" json:"provider,omitempty"`
			RoleArn               *string `tfsdk:"role_arn" json:"roleArn,omitempty"`
		} `tfsdk:"pod_identity" json:"podIdentity,omitempty"`
		SecretTargetRef *[]struct {
			Key       *string `tfsdk:"key" json:"key,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Parameter *string `tfsdk:"parameter" json:"parameter,omitempty"`
		} `tfsdk:"secret_target_ref" json:"secretTargetRef,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KedaShClusterTriggerAuthenticationV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_keda_sh_cluster_trigger_authentication_v1alpha1_manifest"
}

func (r *KedaShClusterTriggerAuthenticationV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ClusterTriggerAuthentication defines how a trigger can authenticate globally",
		MarkdownDescription: "ClusterTriggerAuthentication defines how a trigger can authenticate globally",
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
				Description:         "TriggerAuthenticationSpec defines the various ways to authenticate",
				MarkdownDescription: "TriggerAuthenticationSpec defines the various ways to authenticate",
				Attributes: map[string]schema.Attribute{
					"aws_secret_manager": schema.SingleNestedAttribute{
						Description:         "AwsSecretManager is used to authenticate using AwsSecretManager",
						MarkdownDescription: "AwsSecretManager is used to authenticate using AwsSecretManager",
						Attributes: map[string]schema.Attribute{
							"credentials": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"access_key": schema.SingleNestedAttribute{
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
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: true,
														Optional: false,
														Computed: false,
													},
												},
												Required: true,
												Optional: false,
												Computed: false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"access_secret_key": schema.SingleNestedAttribute{
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
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: true,
														Optional: false,
														Computed: false,
													},
												},
												Required: true,
												Optional: false,
												Computed: false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"access_token": schema.SingleNestedAttribute{
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
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: true,
														Optional: false,
														Computed: false,
													},
												},
												Required: true,
												Optional: false,
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

							"pod_identity": schema.SingleNestedAttribute{
								Description:         "AuthPodIdentity allows users to select the platform native identity mechanism",
								MarkdownDescription: "AuthPodIdentity allows users to select the platform native identity mechanism",
								Attributes: map[string]schema.Attribute{
									"identity_authority_host": schema.StringAttribute{
										Description:         "Set identityAuthorityHost to override the default Azure authority host. If this is set, then the IdentityTenantID must also be set",
										MarkdownDescription: "Set identityAuthorityHost to override the default Azure authority host. If this is set, then the IdentityTenantID must also be set",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"identity_id": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"identity_owner": schema.StringAttribute{
										Description:         "IdentityOwner configures which identity has to be used during auto discovery, keda or the scaled workload. Mutually exclusive with roleArn",
										MarkdownDescription: "IdentityOwner configures which identity has to be used during auto discovery, keda or the scaled workload. Mutually exclusive with roleArn",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("keda", "workload"),
										},
									},

									"identity_tenant_id": schema.StringAttribute{
										Description:         "Set identityTenantId to override the default Azure tenant id. If this is set, then the IdentityID must also be set",
										MarkdownDescription: "Set identityTenantId to override the default Azure tenant id. If this is set, then the IdentityID must also be set",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"provider": schema.StringAttribute{
										Description:         "PodIdentityProvider contains the list of providers",
										MarkdownDescription: "PodIdentityProvider contains the list of providers",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("azure-workload", "gcp", "aws", "aws-eks", "none"),
										},
									},

									"role_arn": schema.StringAttribute{
										Description:         "RoleArn sets the AWS RoleArn to be used. Mutually exclusive with IdentityOwner",
										MarkdownDescription: "RoleArn sets the AWS RoleArn to be used. Mutually exclusive with IdentityOwner",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"region": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secrets": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"parameter": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"version_id": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"version_stage": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
								},
								Required: true,
								Optional: false,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

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
										Optional:            true,
										Computed:            false,
									},

									"key_vault_resource_url": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"type": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"credentials": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"client_id": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
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
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: true,
														Optional: false,
														Computed: false,
													},
												},
												Required: true,
												Optional: false,
												Computed: false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"tenant_id": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"pod_identity": schema.SingleNestedAttribute{
								Description:         "AuthPodIdentity allows users to select the platform native identity mechanism",
								MarkdownDescription: "AuthPodIdentity allows users to select the platform native identity mechanism",
								Attributes: map[string]schema.Attribute{
									"identity_authority_host": schema.StringAttribute{
										Description:         "Set identityAuthorityHost to override the default Azure authority host. If this is set, then the IdentityTenantID must also be set",
										MarkdownDescription: "Set identityAuthorityHost to override the default Azure authority host. If this is set, then the IdentityTenantID must also be set",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"identity_id": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"identity_owner": schema.StringAttribute{
										Description:         "IdentityOwner configures which identity has to be used during auto discovery, keda or the scaled workload. Mutually exclusive with roleArn",
										MarkdownDescription: "IdentityOwner configures which identity has to be used during auto discovery, keda or the scaled workload. Mutually exclusive with roleArn",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("keda", "workload"),
										},
									},

									"identity_tenant_id": schema.StringAttribute{
										Description:         "Set identityTenantId to override the default Azure tenant id. If this is set, then the IdentityID must also be set",
										MarkdownDescription: "Set identityTenantId to override the default Azure tenant id. If this is set, then the IdentityID must also be set",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"provider": schema.StringAttribute{
										Description:         "PodIdentityProvider contains the list of providers",
										MarkdownDescription: "PodIdentityProvider contains the list of providers",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("azure-workload", "gcp", "aws", "aws-eks", "none"),
										},
									},

									"role_arn": schema.StringAttribute{
										Description:         "RoleArn sets the AWS RoleArn to be used. Mutually exclusive with IdentityOwner",
										MarkdownDescription: "RoleArn sets the AWS RoleArn to be used. Mutually exclusive with IdentityOwner",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"secrets": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"parameter": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"version": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
								},
								Required: true,
								Optional: false,
								Computed: false,
							},

							"vault_uri": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"config_map_target_ref": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"key": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"parameter": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
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
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"parameter": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"gcp_secret_manager": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"credentials": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
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
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: true,
														Optional: false,
														Computed: false,
													},
												},
												Required: true,
												Optional: false,
												Computed: false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"pod_identity": schema.SingleNestedAttribute{
								Description:         "AuthPodIdentity allows users to select the platform native identity mechanism",
								MarkdownDescription: "AuthPodIdentity allows users to select the platform native identity mechanism",
								Attributes: map[string]schema.Attribute{
									"identity_authority_host": schema.StringAttribute{
										Description:         "Set identityAuthorityHost to override the default Azure authority host. If this is set, then the IdentityTenantID must also be set",
										MarkdownDescription: "Set identityAuthorityHost to override the default Azure authority host. If this is set, then the IdentityTenantID must also be set",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"identity_id": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"identity_owner": schema.StringAttribute{
										Description:         "IdentityOwner configures which identity has to be used during auto discovery, keda or the scaled workload. Mutually exclusive with roleArn",
										MarkdownDescription: "IdentityOwner configures which identity has to be used during auto discovery, keda or the scaled workload. Mutually exclusive with roleArn",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("keda", "workload"),
										},
									},

									"identity_tenant_id": schema.StringAttribute{
										Description:         "Set identityTenantId to override the default Azure tenant id. If this is set, then the IdentityID must also be set",
										MarkdownDescription: "Set identityTenantId to override the default Azure tenant id. If this is set, then the IdentityID must also be set",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"provider": schema.StringAttribute{
										Description:         "PodIdentityProvider contains the list of providers",
										MarkdownDescription: "PodIdentityProvider contains the list of providers",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("azure-workload", "gcp", "aws", "aws-eks", "none"),
										},
									},

									"role_arn": schema.StringAttribute{
										Description:         "RoleArn sets the AWS RoleArn to be used. Mutually exclusive with IdentityOwner",
										MarkdownDescription: "RoleArn sets the AWS RoleArn to be used. Mutually exclusive with IdentityOwner",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"secrets": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"parameter": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"version": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
								},
								Required: true,
								Optional: false,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"hashi_corp_vault": schema.SingleNestedAttribute{
						Description:         "HashiCorpVault is used to authenticate using Hashicorp Vault",
						MarkdownDescription: "HashiCorpVault is used to authenticate using Hashicorp Vault",
						Attributes: map[string]schema.Attribute{
							"address": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"authentication": schema.StringAttribute{
								Description:         "VaultAuthentication contains the list of Hashicorp Vault authentication methods",
								MarkdownDescription: "VaultAuthentication contains the list of Hashicorp Vault authentication methods",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"credential": schema.SingleNestedAttribute{
								Description:         "Credential defines the Hashicorp Vault credentials depending on the authentication method",
								MarkdownDescription: "Credential defines the Hashicorp Vault credentials depending on the authentication method",
								Attributes: map[string]schema.Attribute{
									"service_account": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"token": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"mount": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"role": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secrets": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"parameter": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"pki_data": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"alt_names": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"common_name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"format": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"ip_sans": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"other_sans": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"ttl": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri_sans": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"type": schema.StringAttribute{
											Description:         "VaultSecretType defines the type of vault secret",
											MarkdownDescription: "VaultSecretType defines the type of vault secret",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
								},
								Required: true,
								Optional: false,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"pod_identity": schema.SingleNestedAttribute{
						Description:         "AuthPodIdentity allows users to select the platform native identity mechanism",
						MarkdownDescription: "AuthPodIdentity allows users to select the platform native identity mechanism",
						Attributes: map[string]schema.Attribute{
							"identity_authority_host": schema.StringAttribute{
								Description:         "Set identityAuthorityHost to override the default Azure authority host. If this is set, then the IdentityTenantID must also be set",
								MarkdownDescription: "Set identityAuthorityHost to override the default Azure authority host. If this is set, then the IdentityTenantID must also be set",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"identity_id": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"identity_owner": schema.StringAttribute{
								Description:         "IdentityOwner configures which identity has to be used during auto discovery, keda or the scaled workload. Mutually exclusive with roleArn",
								MarkdownDescription: "IdentityOwner configures which identity has to be used during auto discovery, keda or the scaled workload. Mutually exclusive with roleArn",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("keda", "workload"),
								},
							},

							"identity_tenant_id": schema.StringAttribute{
								Description:         "Set identityTenantId to override the default Azure tenant id. If this is set, then the IdentityID must also be set",
								MarkdownDescription: "Set identityTenantId to override the default Azure tenant id. If this is set, then the IdentityID must also be set",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"provider": schema.StringAttribute{
								Description:         "PodIdentityProvider contains the list of providers",
								MarkdownDescription: "PodIdentityProvider contains the list of providers",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("azure-workload", "gcp", "aws", "aws-eks", "none"),
								},
							},

							"role_arn": schema.StringAttribute{
								Description:         "RoleArn sets the AWS RoleArn to be used. Mutually exclusive with IdentityOwner",
								MarkdownDescription: "RoleArn sets the AWS RoleArn to be used. Mutually exclusive with IdentityOwner",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"secret_target_ref": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"key": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"parameter": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *KedaShClusterTriggerAuthenticationV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_keda_sh_cluster_trigger_authentication_v1alpha1_manifest")

	var model KedaShClusterTriggerAuthenticationV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("keda.sh/v1alpha1")
	model.Kind = pointer.String("ClusterTriggerAuthentication")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
