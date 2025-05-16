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
	_ datasource.DataSource = &GeneratorsExternalSecretsIoStssessionTokenV1Alpha1Manifest{}
)

func NewGeneratorsExternalSecretsIoStssessionTokenV1Alpha1Manifest() datasource.DataSource {
	return &GeneratorsExternalSecretsIoStssessionTokenV1Alpha1Manifest{}
}

type GeneratorsExternalSecretsIoStssessionTokenV1Alpha1Manifest struct{}

type GeneratorsExternalSecretsIoStssessionTokenV1Alpha1ManifestData struct {
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
			Jwt *struct {
				ServiceAccountRef *struct {
					Audiences *[]string `tfsdk:"audiences" json:"audiences,omitempty"`
					Name      *string   `tfsdk:"name" json:"name,omitempty"`
					Namespace *string   `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"service_account_ref" json:"serviceAccountRef,omitempty"`
			} `tfsdk:"jwt" json:"jwt,omitempty"`
			SecretRef *struct {
				AccessKeyIDSecretRef *struct {
					Key       *string `tfsdk:"key" json:"key,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"access_key_id_secret_ref" json:"accessKeyIDSecretRef,omitempty"`
				SecretAccessKeySecretRef *struct {
					Key       *string `tfsdk:"key" json:"key,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"secret_access_key_secret_ref" json:"secretAccessKeySecretRef,omitempty"`
				SessionTokenSecretRef *struct {
					Key       *string `tfsdk:"key" json:"key,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"session_token_secret_ref" json:"sessionTokenSecretRef,omitempty"`
			} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
		} `tfsdk:"auth" json:"auth,omitempty"`
		Region            *string `tfsdk:"region" json:"region,omitempty"`
		RequestParameters *struct {
			SerialNumber    *string `tfsdk:"serial_number" json:"serialNumber,omitempty"`
			SessionDuration *int64  `tfsdk:"session_duration" json:"sessionDuration,omitempty"`
			TokenCode       *string `tfsdk:"token_code" json:"tokenCode,omitempty"`
		} `tfsdk:"request_parameters" json:"requestParameters,omitempty"`
		Role *string `tfsdk:"role" json:"role,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *GeneratorsExternalSecretsIoStssessionTokenV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_generators_external_secrets_io_sts_session_token_v1alpha1_manifest"
}

func (r *GeneratorsExternalSecretsIoStssessionTokenV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "STSSessionToken uses the GetSessionToken API to retrieve an authorization token. The authorization token is valid for 12 hours. The authorizationToken returned is a base64 encoded string that can be decoded. For more information, see GetSessionToken (https://docs.aws.amazon.com/STS/latest/APIReference/API_GetSessionToken.html).",
		MarkdownDescription: "STSSessionToken uses the GetSessionToken API to retrieve an authorization token. The authorization token is valid for 12 hours. The authorizationToken returned is a base64 encoded string that can be decoded. For more information, see GetSessionToken (https://docs.aws.amazon.com/STS/latest/APIReference/API_GetSessionToken.html).",
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
						Description:         "Auth defines how to authenticate with AWS",
						MarkdownDescription: "Auth defines how to authenticate with AWS",
						Attributes: map[string]schema.Attribute{
							"jwt": schema.SingleNestedAttribute{
								Description:         "Authenticate against AWS using service account tokens.",
								MarkdownDescription: "Authenticate against AWS using service account tokens.",
								Attributes: map[string]schema.Attribute{
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
										Required: false,
										Optional: true,
										Computed: false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"secret_ref": schema.SingleNestedAttribute{
								Description:         "AWSAuthSecretRef holds secret references for AWS credentials both AccessKeyID and SecretAccessKey must be defined in order to properly authenticate.",
								MarkdownDescription: "AWSAuthSecretRef holds secret references for AWS credentials both AccessKeyID and SecretAccessKey must be defined in order to properly authenticate.",
								Attributes: map[string]schema.Attribute{
									"access_key_id_secret_ref": schema.SingleNestedAttribute{
										Description:         "The AccessKeyID is used for authentication",
										MarkdownDescription: "The AccessKeyID is used for authentication",
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

									"session_token_secret_ref": schema.SingleNestedAttribute{
										Description:         "The SessionToken used for authentication This must be defined if AccessKeyID and SecretAccessKey are temporary credentials see: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_use-resources.html",
										MarkdownDescription: "The SessionToken used for authentication This must be defined if AccessKeyID and SecretAccessKey are temporary credentials see: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_use-resources.html",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"region": schema.StringAttribute{
						Description:         "Region specifies the region to operate in.",
						MarkdownDescription: "Region specifies the region to operate in.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"request_parameters": schema.SingleNestedAttribute{
						Description:         "RequestParameters contains parameters that can be passed to the STS service.",
						MarkdownDescription: "RequestParameters contains parameters that can be passed to the STS service.",
						Attributes: map[string]schema.Attribute{
							"serial_number": schema.StringAttribute{
								Description:         "SerialNumber is the identification number of the MFA device that is associated with the IAM user who is making the GetSessionToken call. Possible values: hardware device (such as GAHT12345678) or an Amazon Resource Name (ARN) for a virtual device (such as arn:aws:iam::123456789012:mfa/user)",
								MarkdownDescription: "SerialNumber is the identification number of the MFA device that is associated with the IAM user who is making the GetSessionToken call. Possible values: hardware device (such as GAHT12345678) or an Amazon Resource Name (ARN) for a virtual device (such as arn:aws:iam::123456789012:mfa/user)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"session_duration": schema.Int64Attribute{
								Description:         "SessionDuration The duration, in seconds, that the credentials should remain valid. Acceptable durations for IAM user sessions range from 900 seconds (15 minutes) to 129,600 seconds (36 hours), with 43,200 seconds (12 hours) as the default.",
								MarkdownDescription: "SessionDuration The duration, in seconds, that the credentials should remain valid. Acceptable durations for IAM user sessions range from 900 seconds (15 minutes) to 129,600 seconds (36 hours), with 43,200 seconds (12 hours) as the default.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"token_code": schema.StringAttribute{
								Description:         "TokenCode is the value provided by the MFA device, if MFA is required.",
								MarkdownDescription: "TokenCode is the value provided by the MFA device, if MFA is required.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"role": schema.StringAttribute{
						Description:         "You can assume a role before making calls to the desired AWS service.",
						MarkdownDescription: "You can assume a role before making calls to the desired AWS service.",
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

func (r *GeneratorsExternalSecretsIoStssessionTokenV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_generators_external_secrets_io_sts_session_token_v1alpha1_manifest")

	var model GeneratorsExternalSecretsIoStssessionTokenV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("generators.external-secrets.io/v1alpha1")
	model.Kind = pointer.String("STSSessionToken")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
