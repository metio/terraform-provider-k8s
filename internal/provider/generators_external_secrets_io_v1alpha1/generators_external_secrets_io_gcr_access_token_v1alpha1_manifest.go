/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package generators_external_secrets_io_v1alpha1

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &GeneratorsExternalSecretsIoGcraccessTokenV1Alpha1Manifest{}
)

func NewGeneratorsExternalSecretsIoGcraccessTokenV1Alpha1Manifest() datasource.DataSource {
	return &GeneratorsExternalSecretsIoGcraccessTokenV1Alpha1Manifest{}
}

type GeneratorsExternalSecretsIoGcraccessTokenV1Alpha1Manifest struct{}

type GeneratorsExternalSecretsIoGcraccessTokenV1Alpha1ManifestData struct {
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
		Auth *struct {
			SecretRef *struct {
				SecretAccessKeySecretRef *struct {
					Key       *string `tfsdk:"key" json:"key,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"secret_access_key_secret_ref" json:"secretAccessKeySecretRef,omitempty"`
			} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
			WorkloadIdentity *struct {
				ClusterLocation   *string `tfsdk:"cluster_location" json:"clusterLocation,omitempty"`
				ClusterName       *string `tfsdk:"cluster_name" json:"clusterName,omitempty"`
				ClusterProjectID  *string `tfsdk:"cluster_project_id" json:"clusterProjectID,omitempty"`
				ServiceAccountRef *struct {
					Audiences *[]string `tfsdk:"audiences" json:"audiences,omitempty"`
					Name      *string   `tfsdk:"name" json:"name,omitempty"`
					Namespace *string   `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"service_account_ref" json:"serviceAccountRef,omitempty"`
			} `tfsdk:"workload_identity" json:"workloadIdentity,omitempty"`
			WorkloadIdentityFederation *struct {
				Audience               *string `tfsdk:"audience" json:"audience,omitempty"`
				AwsSecurityCredentials *struct {
					AwsCredentialsSecretRef *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"aws_credentials_secret_ref" json:"awsCredentialsSecretRef,omitempty"`
					Region *string `tfsdk:"region" json:"region,omitempty"`
				} `tfsdk:"aws_security_credentials" json:"awsSecurityCredentials,omitempty"`
				CredConfig *struct {
					Key       *string `tfsdk:"key" json:"key,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"cred_config" json:"credConfig,omitempty"`
				ExternalTokenEndpoint *string `tfsdk:"external_token_endpoint" json:"externalTokenEndpoint,omitempty"`
				ServiceAccountRef     *struct {
					Audiences *[]string `tfsdk:"audiences" json:"audiences,omitempty"`
					Name      *string   `tfsdk:"name" json:"name,omitempty"`
					Namespace *string   `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"service_account_ref" json:"serviceAccountRef,omitempty"`
			} `tfsdk:"workload_identity_federation" json:"workloadIdentityFederation,omitempty"`
		} `tfsdk:"auth" json:"auth,omitempty"`
		ProjectID *string `tfsdk:"project_id" json:"projectID,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *GeneratorsExternalSecretsIoGcraccessTokenV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_generators_external_secrets_io_gcr_access_token_v1alpha1_manifest"
}

func (r *GeneratorsExternalSecretsIoGcraccessTokenV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "GCRAccessToken generates an GCP access token that can be used to authenticate with GCR.",
		MarkdownDescription: "GCRAccessToken generates an GCP access token that can be used to authenticate with GCR.",
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"auth": schema.SingleNestedAttribute{
						Description:         "Auth defines the means for authenticating with GCP",
						MarkdownDescription: "Auth defines the means for authenticating with GCP",
						Attributes: map[string]schema.Attribute{
							"secret_ref": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"secret_access_key_secret_ref": schema.SingleNestedAttribute{
										Description:         "The SecretAccessKey is used for authentication",
										MarkdownDescription: "The SecretAccessKey is used for authentication",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "A key in the referenced Secret. Some instances of this field may be defaulted, in others it may be required.",
												MarkdownDescription: "A key in the referenced Secret. Some instances of this field may be defaulted, in others it may be required.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
													stringvalidator.LengthAtMost(253),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[-._a-zA-Z0-9]+$`), ""),
												},
											},

											"name": schema.StringAttribute{
												Description:         "The name of the Secret resource being referred to.",
												MarkdownDescription: "The name of the Secret resource being referred to.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
													stringvalidator.LengthAtMost(253),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
												},
											},

											"namespace": schema.StringAttribute{
												Description:         "The namespace of the Secret resource being referred to. Ignored if referent is not cluster-scoped, otherwise defaults to the namespace of the referent.",
												MarkdownDescription: "The namespace of the Secret resource being referred to. Ignored if referent is not cluster-scoped, otherwise defaults to the namespace of the referent.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
													stringvalidator.LengthAtMost(63),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
												},
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

							"workload_identity": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cluster_location": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"cluster_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"cluster_project_id": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"service_account_ref": schema.SingleNestedAttribute{
										Description:         "A reference to a ServiceAccount resource.",
										MarkdownDescription: "A reference to a ServiceAccount resource.",
										Attributes: map[string]schema.Attribute{
											"audiences": schema.ListAttribute{
												Description:         "Audience specifies the 'aud' claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
												MarkdownDescription: "Audience specifies the 'aud' claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "The name of the ServiceAccount resource being referred to.",
												MarkdownDescription: "The name of the ServiceAccount resource being referred to.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
													stringvalidator.LengthAtMost(253),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
												},
											},

											"namespace": schema.StringAttribute{
												Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped, otherwise defaults to the namespace of the referent.",
												MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped, otherwise defaults to the namespace of the referent.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
													stringvalidator.LengthAtMost(63),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
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

							"workload_identity_federation": schema.SingleNestedAttribute{
								Description:         "GCPWorkloadIdentityFederation holds the configurations required for generating federated access tokens.",
								MarkdownDescription: "GCPWorkloadIdentityFederation holds the configurations required for generating federated access tokens.",
								Attributes: map[string]schema.Attribute{
									"audience": schema.StringAttribute{
										Description:         "audience is the Secure Token Service (STS) audience which contains the resource name for the workload identity pool and the provider identifier in that pool. If specified, Audience found in the external account credential config will be overridden with the configured value. audience must be provided when serviceAccountRef or awsSecurityCredentials is configured.",
										MarkdownDescription: "audience is the Secure Token Service (STS) audience which contains the resource name for the workload identity pool and the provider identifier in that pool. If specified, Audience found in the external account credential config will be overridden with the configured value. audience must be provided when serviceAccountRef or awsSecurityCredentials is configured.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"aws_security_credentials": schema.SingleNestedAttribute{
										Description:         "awsSecurityCredentials is for configuring AWS region and credentials to use for obtaining the access token, when using the AWS metadata server is not an option.",
										MarkdownDescription: "awsSecurityCredentials is for configuring AWS region and credentials to use for obtaining the access token, when using the AWS metadata server is not an option.",
										Attributes: map[string]schema.Attribute{
											"aws_credentials_secret_ref": schema.SingleNestedAttribute{
												Description:         "awsCredentialsSecretRef is the reference to the secret which holds the AWS credentials. Secret should be created with below names for keys - aws_access_key_id: Access Key ID, which is the unique identifier for the AWS account or the IAM user. - aws_secret_access_key: Secret Access Key, which is used to authenticate requests made to AWS services. - aws_session_token: Session Token, is the short-lived token to authenticate requests made to AWS services.",
												MarkdownDescription: "awsCredentialsSecretRef is the reference to the secret which holds the AWS credentials. Secret should be created with below names for keys - aws_access_key_id: Access Key ID, which is the unique identifier for the AWS account or the IAM user. - aws_secret_access_key: Secret Access Key, which is used to authenticate requests made to AWS services. - aws_session_token: Session Token, is the short-lived token to authenticate requests made to AWS services.",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "name of the secret.",
														MarkdownDescription: "name of the secret.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.LengthAtMost(253),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
														},
													},

													"namespace": schema.StringAttribute{
														Description:         "namespace in which the secret exists. If empty, secret will looked up in local namespace.",
														MarkdownDescription: "namespace in which the secret exists. If empty, secret will looked up in local namespace.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.LengthAtMost(63),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
														},
													},
												},
												Required: true,
												Optional: false,
												Computed: false,
											},

											"region": schema.StringAttribute{
												Description:         "region is for configuring the AWS region to be used.",
												MarkdownDescription: "region is for configuring the AWS region to be used.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
													stringvalidator.LengthAtMost(50),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9-]+$`), ""),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"cred_config": schema.SingleNestedAttribute{
										Description:         "credConfig holds the configmap reference containing the GCP external account credential configuration in JSON format and the key name containing the json data. For using Kubernetes cluster as the identity provider, use serviceAccountRef instead. Operators mounted serviceaccount token cannot be used as the token source, instead serviceAccountRef must be used by providing operators service account details.",
										MarkdownDescription: "credConfig holds the configmap reference containing the GCP external account credential configuration in JSON format and the key name containing the json data. For using Kubernetes cluster as the identity provider, use serviceAccountRef instead. Operators mounted serviceaccount token cannot be used as the token source, instead serviceAccountRef must be used by providing operators service account details.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "key name holding the external account credential config.",
												MarkdownDescription: "key name holding the external account credential config.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
													stringvalidator.LengthAtMost(253),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[-._a-zA-Z0-9]+$`), ""),
												},
											},

											"name": schema.StringAttribute{
												Description:         "name of the configmap.",
												MarkdownDescription: "name of the configmap.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
													stringvalidator.LengthAtMost(253),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
												},
											},

											"namespace": schema.StringAttribute{
												Description:         "namespace in which the configmap exists. If empty, configmap will looked up in local namespace.",
												MarkdownDescription: "namespace in which the configmap exists. If empty, configmap will looked up in local namespace.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
													stringvalidator.LengthAtMost(63),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"external_token_endpoint": schema.StringAttribute{
										Description:         "externalTokenEndpoint is the endpoint explicitly set up to provide tokens, which will be matched against the credential_source.url in the provided credConfig. This field is merely to double-check the external token source URL is having the expected value.",
										MarkdownDescription: "externalTokenEndpoint is the endpoint explicitly set up to provide tokens, which will be matched against the credential_source.url in the provided credConfig. This field is merely to double-check the external token source URL is having the expected value.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"service_account_ref": schema.SingleNestedAttribute{
										Description:         "serviceAccountRef is the reference to the kubernetes ServiceAccount to be used for obtaining the tokens, when Kubernetes is configured as provider in workload identity pool.",
										MarkdownDescription: "serviceAccountRef is the reference to the kubernetes ServiceAccount to be used for obtaining the tokens, when Kubernetes is configured as provider in workload identity pool.",
										Attributes: map[string]schema.Attribute{
											"audiences": schema.ListAttribute{
												Description:         "Audience specifies the 'aud' claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
												MarkdownDescription: "Audience specifies the 'aud' claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "The name of the ServiceAccount resource being referred to.",
												MarkdownDescription: "The name of the ServiceAccount resource being referred to.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
													stringvalidator.LengthAtMost(253),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
												},
											},

											"namespace": schema.StringAttribute{
												Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped, otherwise defaults to the namespace of the referent.",
												MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped, otherwise defaults to the namespace of the referent.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
													stringvalidator.LengthAtMost(63),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
												},
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
						Required: true,
						Optional: false,
						Computed: false,
					},

					"project_id": schema.StringAttribute{
						Description:         "ProjectID defines which project to use to authenticate with",
						MarkdownDescription: "ProjectID defines which project to use to authenticate with",
						Required:            true,
						Optional:            false,
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

func (r *GeneratorsExternalSecretsIoGcraccessTokenV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_generators_external_secrets_io_gcr_access_token_v1alpha1_manifest")

	var model GeneratorsExternalSecretsIoGcraccessTokenV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("generators.external-secrets.io/v1alpha1")
	model.Kind = pointer.String("GCRAccessToken")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
