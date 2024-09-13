/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package secrets_hashicorp_com_v1beta1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &SecretsHashicorpComVaultAuthV1Beta1Manifest{}
)

func NewSecretsHashicorpComVaultAuthV1Beta1Manifest() datasource.DataSource {
	return &SecretsHashicorpComVaultAuthV1Beta1Manifest{}
}

type SecretsHashicorpComVaultAuthV1Beta1Manifest struct{}

type SecretsHashicorpComVaultAuthV1Beta1ManifestData struct {
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
		AllowedNamespaces *[]string `tfsdk:"allowed_namespaces" json:"allowedNamespaces,omitempty"`
		AppRole           *struct {
			RoleId    *string `tfsdk:"role_id" json:"roleId,omitempty"`
			SecretRef *string `tfsdk:"secret_ref" json:"secretRef,omitempty"`
		} `tfsdk:"app_role" json:"appRole,omitempty"`
		Aws *struct {
			HeaderValue        *string `tfsdk:"header_value" json:"headerValue,omitempty"`
			IamEndpoint        *string `tfsdk:"iam_endpoint" json:"iamEndpoint,omitempty"`
			IrsaServiceAccount *string `tfsdk:"irsa_service_account" json:"irsaServiceAccount,omitempty"`
			Region             *string `tfsdk:"region" json:"region,omitempty"`
			Role               *string `tfsdk:"role" json:"role,omitempty"`
			SecretRef          *string `tfsdk:"secret_ref" json:"secretRef,omitempty"`
			SessionName        *string `tfsdk:"session_name" json:"sessionName,omitempty"`
			StsEndpoint        *string `tfsdk:"sts_endpoint" json:"stsEndpoint,omitempty"`
		} `tfsdk:"aws" json:"aws,omitempty"`
		Gcp *struct {
			ClusterName                    *string `tfsdk:"cluster_name" json:"clusterName,omitempty"`
			ProjectID                      *string `tfsdk:"project_id" json:"projectID,omitempty"`
			Region                         *string `tfsdk:"region" json:"region,omitempty"`
			Role                           *string `tfsdk:"role" json:"role,omitempty"`
			WorkloadIdentityServiceAccount *string `tfsdk:"workload_identity_service_account" json:"workloadIdentityServiceAccount,omitempty"`
		} `tfsdk:"gcp" json:"gcp,omitempty"`
		Headers *map[string]string `tfsdk:"headers" json:"headers,omitempty"`
		Jwt     *struct {
			Audiences              *[]string `tfsdk:"audiences" json:"audiences,omitempty"`
			Role                   *string   `tfsdk:"role" json:"role,omitempty"`
			SecretRef              *string   `tfsdk:"secret_ref" json:"secretRef,omitempty"`
			ServiceAccount         *string   `tfsdk:"service_account" json:"serviceAccount,omitempty"`
			TokenExpirationSeconds *int64    `tfsdk:"token_expiration_seconds" json:"tokenExpirationSeconds,omitempty"`
		} `tfsdk:"jwt" json:"jwt,omitempty"`
		Kubernetes *struct {
			Audiences              *[]string `tfsdk:"audiences" json:"audiences,omitempty"`
			Role                   *string   `tfsdk:"role" json:"role,omitempty"`
			ServiceAccount         *string   `tfsdk:"service_account" json:"serviceAccount,omitempty"`
			TokenExpirationSeconds *int64    `tfsdk:"token_expiration_seconds" json:"tokenExpirationSeconds,omitempty"`
		} `tfsdk:"kubernetes" json:"kubernetes,omitempty"`
		Method            *string            `tfsdk:"method" json:"method,omitempty"`
		Mount             *string            `tfsdk:"mount" json:"mount,omitempty"`
		Namespace         *string            `tfsdk:"namespace" json:"namespace,omitempty"`
		Params            *map[string]string `tfsdk:"params" json:"params,omitempty"`
		StorageEncryption *struct {
			KeyName *string `tfsdk:"key_name" json:"keyName,omitempty"`
			Mount   *string `tfsdk:"mount" json:"mount,omitempty"`
		} `tfsdk:"storage_encryption" json:"storageEncryption,omitempty"`
		VaultAuthGlobalRef *struct {
			AllowDefault  *bool `tfsdk:"allow_default" json:"allowDefault,omitempty"`
			MergeStrategy *struct {
				Headers *string `tfsdk:"headers" json:"headers,omitempty"`
				Params  *string `tfsdk:"params" json:"params,omitempty"`
			} `tfsdk:"merge_strategy" json:"mergeStrategy,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"vault_auth_global_ref" json:"vaultAuthGlobalRef,omitempty"`
		VaultConnectionRef *string `tfsdk:"vault_connection_ref" json:"vaultConnectionRef,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SecretsHashicorpComVaultAuthV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_secrets_hashicorp_com_vault_auth_v1beta1_manifest"
}

func (r *SecretsHashicorpComVaultAuthV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "VaultAuth is the Schema for the vaultauths API",
		MarkdownDescription: "VaultAuth is the Schema for the vaultauths API",
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
				Description:         "VaultAuthSpec defines the desired state of VaultAuth",
				MarkdownDescription: "VaultAuthSpec defines the desired state of VaultAuth",
				Attributes: map[string]schema.Attribute{
					"allowed_namespaces": schema.ListAttribute{
						Description:         "AllowedNamespaces Kubernetes Namespaces which are allow-listed for use with this AuthMethod. This field allows administrators to customize which Kubernetes namespaces are authorized to use with this AuthMethod. While Vault will still enforce its own rules, this has the added configurability of restricting which VaultAuthMethods can be used by which namespaces. Accepted values: []{'*'} - wildcard, all namespaces. []{'a', 'b'} - list of namespaces. unset - disallow all namespaces except the Operator's the VaultAuthMethod's namespace, this is the default behavior.",
						MarkdownDescription: "AllowedNamespaces Kubernetes Namespaces which are allow-listed for use with this AuthMethod. This field allows administrators to customize which Kubernetes namespaces are authorized to use with this AuthMethod. While Vault will still enforce its own rules, this has the added configurability of restricting which VaultAuthMethods can be used by which namespaces. Accepted values: []{'*'} - wildcard, all namespaces. []{'a', 'b'} - list of namespaces. unset - disallow all namespaces except the Operator's the VaultAuthMethod's namespace, this is the default behavior.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"app_role": schema.SingleNestedAttribute{
						Description:         "AppRole specific auth configuration, requires that the Method be set to 'appRole'.",
						MarkdownDescription: "AppRole specific auth configuration, requires that the Method be set to 'appRole'.",
						Attributes: map[string]schema.Attribute{
							"role_id": schema.StringAttribute{
								Description:         "RoleID of the AppRole Role to use for authenticating to Vault.",
								MarkdownDescription: "RoleID of the AppRole Role to use for authenticating to Vault.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secret_ref": schema.StringAttribute{
								Description:         "SecretRef is the name of a Kubernetes secret in the consumer's (VDS/VSS/PKI) namespace which provides the AppRole Role's SecretID. The secret must have a key named 'id' which holds the AppRole Role's secretID.",
								MarkdownDescription: "SecretRef is the name of a Kubernetes secret in the consumer's (VDS/VSS/PKI) namespace which provides the AppRole Role's SecretID. The secret must have a key named 'id' which holds the AppRole Role's secretID.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"aws": schema.SingleNestedAttribute{
						Description:         "AWS specific auth configuration, requires that Method be set to 'aws'.",
						MarkdownDescription: "AWS specific auth configuration, requires that Method be set to 'aws'.",
						Attributes: map[string]schema.Attribute{
							"header_value": schema.StringAttribute{
								Description:         "The Vault header value to include in the STS signing request",
								MarkdownDescription: "The Vault header value to include in the STS signing request",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"iam_endpoint": schema.StringAttribute{
								Description:         "The IAM endpoint to use; if not set will use the default",
								MarkdownDescription: "The IAM endpoint to use; if not set will use the default",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"irsa_service_account": schema.StringAttribute{
								Description:         "IRSAServiceAccount name to use with IAM Roles for Service Accounts (IRSA), and should be annotated with 'eks.amazonaws.com/role-arn'. This ServiceAccount will be checked for other EKS annotations: eks.amazonaws.com/audience and eks.amazonaws.com/token-expiration",
								MarkdownDescription: "IRSAServiceAccount name to use with IAM Roles for Service Accounts (IRSA), and should be annotated with 'eks.amazonaws.com/role-arn'. This ServiceAccount will be checked for other EKS annotations: eks.amazonaws.com/audience and eks.amazonaws.com/token-expiration",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"region": schema.StringAttribute{
								Description:         "AWS Region to use for signing the authentication request",
								MarkdownDescription: "AWS Region to use for signing the authentication request",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"role": schema.StringAttribute{
								Description:         "Vault role to use for authenticating",
								MarkdownDescription: "Vault role to use for authenticating",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secret_ref": schema.StringAttribute{
								Description:         "SecretRef is the name of a Kubernetes Secret in the consumer's (VDS/VSS/PKI) namespace which holds credentials for AWS. Expected keys include 'access_key_id', 'secret_access_key', 'session_token'",
								MarkdownDescription: "SecretRef is the name of a Kubernetes Secret in the consumer's (VDS/VSS/PKI) namespace which holds credentials for AWS. Expected keys include 'access_key_id', 'secret_access_key', 'session_token'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"session_name": schema.StringAttribute{
								Description:         "The role session name to use when creating a webidentity provider",
								MarkdownDescription: "The role session name to use when creating a webidentity provider",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"sts_endpoint": schema.StringAttribute{
								Description:         "The STS endpoint to use; if not set will use the default",
								MarkdownDescription: "The STS endpoint to use; if not set will use the default",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"gcp": schema.SingleNestedAttribute{
						Description:         "GCP specific auth configuration, requires that Method be set to 'gcp'.",
						MarkdownDescription: "GCP specific auth configuration, requires that Method be set to 'gcp'.",
						Attributes: map[string]schema.Attribute{
							"cluster_name": schema.StringAttribute{
								Description:         "GKE cluster name. Defaults to the cluster-name returned from the operator pod's local metadata server.",
								MarkdownDescription: "GKE cluster name. Defaults to the cluster-name returned from the operator pod's local metadata server.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"project_id": schema.StringAttribute{
								Description:         "GCP project ID. Defaults to the project-id returned from the operator pod's local metadata server.",
								MarkdownDescription: "GCP project ID. Defaults to the project-id returned from the operator pod's local metadata server.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"region": schema.StringAttribute{
								Description:         "GCP Region of the GKE cluster's identity provider. Defaults to the region returned from the operator pod's local metadata server.",
								MarkdownDescription: "GCP Region of the GKE cluster's identity provider. Defaults to the region returned from the operator pod's local metadata server.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"role": schema.StringAttribute{
								Description:         "Vault role to use for authenticating",
								MarkdownDescription: "Vault role to use for authenticating",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"workload_identity_service_account": schema.StringAttribute{
								Description:         "WorkloadIdentityServiceAccount is the name of a Kubernetes service account (in the same Kubernetes namespace as the Vault*Secret referencing this resource) which has been configured for workload identity in GKE. Should be annotated with 'iam.gke.io/gcp-service-account'.",
								MarkdownDescription: "WorkloadIdentityServiceAccount is the name of a Kubernetes service account (in the same Kubernetes namespace as the Vault*Secret referencing this resource) which has been configured for workload identity in GKE. Should be annotated with 'iam.gke.io/gcp-service-account'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"headers": schema.MapAttribute{
						Description:         "Headers to be included in all Vault requests.",
						MarkdownDescription: "Headers to be included in all Vault requests.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"jwt": schema.SingleNestedAttribute{
						Description:         "JWT specific auth configuration, requires that the Method be set to 'jwt'.",
						MarkdownDescription: "JWT specific auth configuration, requires that the Method be set to 'jwt'.",
						Attributes: map[string]schema.Attribute{
							"audiences": schema.ListAttribute{
								Description:         "TokenAudiences to include in the ServiceAccount token.",
								MarkdownDescription: "TokenAudiences to include in the ServiceAccount token.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"role": schema.StringAttribute{
								Description:         "Role to use for authenticating to Vault.",
								MarkdownDescription: "Role to use for authenticating to Vault.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secret_ref": schema.StringAttribute{
								Description:         "SecretRef is the name of a Kubernetes secret in the consumer's (VDS/VSS/PKI) namespace which provides the JWT token to authenticate to Vault's JWT authentication backend. The secret must have a key named 'jwt' which holds the JWT token.",
								MarkdownDescription: "SecretRef is the name of a Kubernetes secret in the consumer's (VDS/VSS/PKI) namespace which provides the JWT token to authenticate to Vault's JWT authentication backend. The secret must have a key named 'jwt' which holds the JWT token.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_account": schema.StringAttribute{
								Description:         "ServiceAccount to use when creating a ServiceAccount token to authenticate to Vault's JWT authentication backend.",
								MarkdownDescription: "ServiceAccount to use when creating a ServiceAccount token to authenticate to Vault's JWT authentication backend.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"token_expiration_seconds": schema.Int64Attribute{
								Description:         "TokenExpirationSeconds to set the ServiceAccount token.",
								MarkdownDescription: "TokenExpirationSeconds to set the ServiceAccount token.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(600),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"kubernetes": schema.SingleNestedAttribute{
						Description:         "Kubernetes specific auth configuration, requires that the Method be set to 'kubernetes'.",
						MarkdownDescription: "Kubernetes specific auth configuration, requires that the Method be set to 'kubernetes'.",
						Attributes: map[string]schema.Attribute{
							"audiences": schema.ListAttribute{
								Description:         "TokenAudiences to include in the ServiceAccount token.",
								MarkdownDescription: "TokenAudiences to include in the ServiceAccount token.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"role": schema.StringAttribute{
								Description:         "Role to use for authenticating to Vault.",
								MarkdownDescription: "Role to use for authenticating to Vault.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_account": schema.StringAttribute{
								Description:         "ServiceAccount to use when authenticating to Vault's authentication backend. This must reside in the consuming secret's (VDS/VSS/PKI) namespace.",
								MarkdownDescription: "ServiceAccount to use when authenticating to Vault's authentication backend. This must reside in the consuming secret's (VDS/VSS/PKI) namespace.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"token_expiration_seconds": schema.Int64Attribute{
								Description:         "TokenExpirationSeconds to set the ServiceAccount token.",
								MarkdownDescription: "TokenExpirationSeconds to set the ServiceAccount token.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(600),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"method": schema.StringAttribute{
						Description:         "Method to use when authenticating to Vault.",
						MarkdownDescription: "Method to use when authenticating to Vault.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("kubernetes", "jwt", "appRole", "aws", "gcp"),
						},
					},

					"mount": schema.StringAttribute{
						Description:         "Mount to use when authenticating to auth method.",
						MarkdownDescription: "Mount to use when authenticating to auth method.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"namespace": schema.StringAttribute{
						Description:         "Namespace to auth to in Vault",
						MarkdownDescription: "Namespace to auth to in Vault",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"params": schema.MapAttribute{
						Description:         "Params to use when authenticating to Vault",
						MarkdownDescription: "Params to use when authenticating to Vault",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"storage_encryption": schema.SingleNestedAttribute{
						Description:         "StorageEncryption provides the necessary configuration to encrypt the client storage cache. This should only be configured when client cache persistence with encryption is enabled. This is done by passing setting the manager's commandline argument --client-cache-persistence-model=direct-encrypted. Typically, there should only ever be one VaultAuth configured with StorageEncryption in the Cluster, and it should have the label: cacheStorageEncryption=true",
						MarkdownDescription: "StorageEncryption provides the necessary configuration to encrypt the client storage cache. This should only be configured when client cache persistence with encryption is enabled. This is done by passing setting the manager's commandline argument --client-cache-persistence-model=direct-encrypted. Typically, there should only ever be one VaultAuth configured with StorageEncryption in the Cluster, and it should have the label: cacheStorageEncryption=true",
						Attributes: map[string]schema.Attribute{
							"key_name": schema.StringAttribute{
								Description:         "KeyName to use for encrypt/decrypt operations via Vault Transit.",
								MarkdownDescription: "KeyName to use for encrypt/decrypt operations via Vault Transit.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"mount": schema.StringAttribute{
								Description:         "Mount path of the Transit engine in Vault.",
								MarkdownDescription: "Mount path of the Transit engine in Vault.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"vault_auth_global_ref": schema.SingleNestedAttribute{
						Description:         "VaultAuthGlobalRef.",
						MarkdownDescription: "VaultAuthGlobalRef.",
						Attributes: map[string]schema.Attribute{
							"allow_default": schema.BoolAttribute{
								Description:         "AllowDefault when set to true will use the default VaultAuthGlobal resource as the default if Name is not set. The 'allow-default-globals' option must be set on the operator's '-global-vault-auth-options' flag The default VaultAuthGlobal search is conditional. When a ref Namespace is set, the search for the default VaultAuthGlobal resource is constrained to that namespace. Otherwise, the search order is: 1. The default VaultAuthGlobal resource in the referring VaultAuth resource's namespace. 2. The default VaultAuthGlobal resource in the Operator's namespace.",
								MarkdownDescription: "AllowDefault when set to true will use the default VaultAuthGlobal resource as the default if Name is not set. The 'allow-default-globals' option must be set on the operator's '-global-vault-auth-options' flag The default VaultAuthGlobal search is conditional. When a ref Namespace is set, the search for the default VaultAuthGlobal resource is constrained to that namespace. Otherwise, the search order is: 1. The default VaultAuthGlobal resource in the referring VaultAuth resource's namespace. 2. The default VaultAuthGlobal resource in the Operator's namespace.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"merge_strategy": schema.SingleNestedAttribute{
								Description:         "MergeStrategy configures the merge strategy for HTTP headers and parameters that are included in all Vault authentication requests.",
								MarkdownDescription: "MergeStrategy configures the merge strategy for HTTP headers and parameters that are included in all Vault authentication requests.",
								Attributes: map[string]schema.Attribute{
									"headers": schema.StringAttribute{
										Description:         "Headers configures the merge strategy for HTTP headers that are included in all Vault requests. Choices are 'union', 'replace', or 'none'. If 'union' is set, the headers from the VaultAuthGlobal and VaultAuth resources are merged. The headers from the VaultAuth always take precedence. If 'replace' is set, the first set of non-empty headers taken in order from: VaultAuth, VaultAuthGlobal auth method, VaultGlobal default headers. If 'none' is set, the headers from the VaultAuthGlobal resource are ignored and only the headers from the VaultAuth resource are used. The default is 'none'.",
										MarkdownDescription: "Headers configures the merge strategy for HTTP headers that are included in all Vault requests. Choices are 'union', 'replace', or 'none'. If 'union' is set, the headers from the VaultAuthGlobal and VaultAuth resources are merged. The headers from the VaultAuth always take precedence. If 'replace' is set, the first set of non-empty headers taken in order from: VaultAuth, VaultAuthGlobal auth method, VaultGlobal default headers. If 'none' is set, the headers from the VaultAuthGlobal resource are ignored and only the headers from the VaultAuth resource are used. The default is 'none'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("union", "replace", "none"),
										},
									},

									"params": schema.StringAttribute{
										Description:         "Params configures the merge strategy for HTTP parameters that are included in all Vault requests. Choices are 'union', 'replace', or 'none'. If 'union' is set, the parameters from the VaultAuthGlobal and VaultAuth resources are merged. The parameters from the VaultAuth always take precedence. If 'replace' is set, the first set of non-empty parameters taken in order from: VaultAuth, VaultAuthGlobal auth method, VaultGlobal default parameters. If 'none' is set, the parameters from the VaultAuthGlobal resource are ignored and only the parameters from the VaultAuth resource are used. The default is 'none'.",
										MarkdownDescription: "Params configures the merge strategy for HTTP parameters that are included in all Vault requests. Choices are 'union', 'replace', or 'none'. If 'union' is set, the parameters from the VaultAuthGlobal and VaultAuth resources are merged. The parameters from the VaultAuth always take precedence. If 'replace' is set, the first set of non-empty parameters taken in order from: VaultAuth, VaultAuthGlobal auth method, VaultGlobal default parameters. If 'none' is set, the parameters from the VaultAuthGlobal resource are ignored and only the parameters from the VaultAuth resource are used. The default is 'none'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("union", "replace", "none"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the VaultAuthGlobal resource.",
								MarkdownDescription: "Name of the VaultAuthGlobal resource.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9.-]{1,253})$`), ""),
								},
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace of the VaultAuthGlobal resource. If not provided, the namespace of the referring VaultAuth resource is used.",
								MarkdownDescription: "Namespace of the VaultAuthGlobal resource. If not provided, the namespace of the referring VaultAuth resource is used.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9.-]{1,253})$`), ""),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"vault_connection_ref": schema.StringAttribute{
						Description:         "VaultConnectionRef to the VaultConnection resource, can be prefixed with a namespace, eg: 'namespaceA/vaultConnectionRefB'. If no namespace prefix is provided it will default to namespace of the VaultConnection CR. If no value is specified for VaultConnectionRef the Operator will default to the 'default' VaultConnection, configured in the operator's namespace.",
						MarkdownDescription: "VaultConnectionRef to the VaultConnection resource, can be prefixed with a namespace, eg: 'namespaceA/vaultConnectionRefB'. If no namespace prefix is provided it will default to namespace of the VaultConnection CR. If no value is specified for VaultConnectionRef the Operator will default to the 'default' VaultConnection, configured in the operator's namespace.",
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

func (r *SecretsHashicorpComVaultAuthV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_secrets_hashicorp_com_vault_auth_v1beta1_manifest")

	var model SecretsHashicorpComVaultAuthV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("secrets.hashicorp.com/v1beta1")
	model.Kind = pointer.String("VaultAuth")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
