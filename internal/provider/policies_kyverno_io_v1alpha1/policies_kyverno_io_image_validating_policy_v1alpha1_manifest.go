/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package policies_kyverno_io_v1alpha1

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
	_ datasource.DataSource = &PoliciesKyvernoIoImageValidatingPolicyV1Alpha1Manifest{}
)

func NewPoliciesKyvernoIoImageValidatingPolicyV1Alpha1Manifest() datasource.DataSource {
	return &PoliciesKyvernoIoImageValidatingPolicyV1Alpha1Manifest{}
}

type PoliciesKyvernoIoImageValidatingPolicyV1Alpha1Manifest struct{}

type PoliciesKyvernoIoImageValidatingPolicyV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Attestations *[]struct {
			Intoto *struct {
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"intoto" json:"intoto,omitempty"`
			Name     *string `tfsdk:"name" json:"name,omitempty"`
			Referrer *struct {
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"referrer" json:"referrer,omitempty"`
		} `tfsdk:"attestations" json:"attestations,omitempty"`
		Attestors *[]struct {
			Cosign *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Certificate *struct {
					Cert *struct {
						Expression *string `tfsdk:"expression" json:"expression,omitempty"`
						Value      *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"cert" json:"cert,omitempty"`
					CertChain *struct {
						Expression *string `tfsdk:"expression" json:"expression,omitempty"`
						Value      *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"cert_chain" json:"certChain,omitempty"`
				} `tfsdk:"certificate" json:"certificate,omitempty"`
				Ctlog *struct {
					CtLogPubKey        *string `tfsdk:"ct_log_pub_key" json:"ctLogPubKey,omitempty"`
					InsecureIgnoreSCT  *bool   `tfsdk:"insecure_ignore_sct" json:"insecureIgnoreSCT,omitempty"`
					InsecureIgnoreTlog *bool   `tfsdk:"insecure_ignore_tlog" json:"insecureIgnoreTlog,omitempty"`
					RekorPubKey        *string `tfsdk:"rekor_pub_key" json:"rekorPubKey,omitempty"`
					TsaCertChain       *string `tfsdk:"tsa_cert_chain" json:"tsaCertChain,omitempty"`
					Url                *string `tfsdk:"url" json:"url,omitempty"`
				} `tfsdk:"ctlog" json:"ctlog,omitempty"`
				Key *struct {
					Data          *string `tfsdk:"data" json:"data,omitempty"`
					Expression    *string `tfsdk:"expression" json:"expression,omitempty"`
					HashAlgorithm *string `tfsdk:"hash_algorithm" json:"hashAlgorithm,omitempty"`
					Kms           *string `tfsdk:"kms" json:"kms,omitempty"`
				} `tfsdk:"key" json:"key,omitempty"`
				Keyless *struct {
					Identities *[]struct {
						Issuer        *string `tfsdk:"issuer" json:"issuer,omitempty"`
						IssuerRegExp  *string `tfsdk:"issuer_reg_exp" json:"issuerRegExp,omitempty"`
						Subject       *string `tfsdk:"subject" json:"subject,omitempty"`
						SubjectRegExp *string `tfsdk:"subject_reg_exp" json:"subjectRegExp,omitempty"`
					} `tfsdk:"identities" json:"identities,omitempty"`
					Roots *string `tfsdk:"roots" json:"roots,omitempty"`
				} `tfsdk:"keyless" json:"keyless,omitempty"`
				Source *struct {
					PullSecrets *[]struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"pull_secrets" json:"PullSecrets,omitempty"`
					Repository *string `tfsdk:"repository" json:"repository,omitempty"`
					TagPrefix  *string `tfsdk:"tag_prefix" json:"tagPrefix,omitempty"`
				} `tfsdk:"source" json:"source,omitempty"`
				Tuf *struct {
					Mirror *string `tfsdk:"mirror" json:"mirror,omitempty"`
					Root   *struct {
						Data *string `tfsdk:"data" json:"data,omitempty"`
						Path *string `tfsdk:"path" json:"path,omitempty"`
					} `tfsdk:"root" json:"root,omitempty"`
				} `tfsdk:"tuf" json:"tuf,omitempty"`
			} `tfsdk:"cosign" json:"cosign,omitempty"`
			Name   *string `tfsdk:"name" json:"name,omitempty"`
			Notary *struct {
				Certs *struct {
					Expression *string `tfsdk:"expression" json:"expression,omitempty"`
					Value      *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"certs" json:"certs,omitempty"`
				TsaCerts *struct {
					Expression *string `tfsdk:"expression" json:"expression,omitempty"`
					Value      *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"tsa_certs" json:"tsaCerts,omitempty"`
			} `tfsdk:"notary" json:"notary,omitempty"`
		} `tfsdk:"attestors" json:"attestors,omitempty"`
		AuditAnnotations *[]struct {
			Key             *string `tfsdk:"key" json:"key,omitempty"`
			ValueExpression *string `tfsdk:"value_expression" json:"valueExpression,omitempty"`
		} `tfsdk:"audit_annotations" json:"auditAnnotations,omitempty"`
		Autogen *struct {
			PodControllers *struct {
				Controllers *[]string `tfsdk:"controllers" json:"controllers,omitempty"`
			} `tfsdk:"pod_controllers" json:"podControllers,omitempty"`
		} `tfsdk:"autogen" json:"autogen,omitempty"`
		Credentials *struct {
			AllowInsecureRegistry *bool     `tfsdk:"allow_insecure_registry" json:"allowInsecureRegistry,omitempty"`
			Providers             *[]string `tfsdk:"providers" json:"providers,omitempty"`
			Secrets               *[]string `tfsdk:"secrets" json:"secrets,omitempty"`
		} `tfsdk:"credentials" json:"credentials,omitempty"`
		Evaluation *struct {
			Admission *struct {
				Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
			} `tfsdk:"admission" json:"admission,omitempty"`
			Background *struct {
				Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
			} `tfsdk:"background" json:"background,omitempty"`
			Mode *string `tfsdk:"mode" json:"mode,omitempty"`
		} `tfsdk:"evaluation" json:"evaluation,omitempty"`
		FailurePolicy *string `tfsdk:"failure_policy" json:"failurePolicy,omitempty"`
		Images        *[]struct {
			Expression *string `tfsdk:"expression" json:"expression,omitempty"`
			Name       *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"images" json:"images,omitempty"`
		MatchConditions *[]struct {
			Expression *string `tfsdk:"expression" json:"expression,omitempty"`
			Name       *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"match_conditions" json:"matchConditions,omitempty"`
		MatchConstraints *struct {
			ExcludeResourceRules *[]struct {
				ApiGroups     *[]string `tfsdk:"api_groups" json:"apiGroups,omitempty"`
				ApiVersions   *[]string `tfsdk:"api_versions" json:"apiVersions,omitempty"`
				Operations    *[]string `tfsdk:"operations" json:"operations,omitempty"`
				ResourceNames *[]string `tfsdk:"resource_names" json:"resourceNames,omitempty"`
				Resources     *[]string `tfsdk:"resources" json:"resources,omitempty"`
				Scope         *string   `tfsdk:"scope" json:"scope,omitempty"`
			} `tfsdk:"exclude_resource_rules" json:"excludeResourceRules,omitempty"`
			MatchPolicy       *string `tfsdk:"match_policy" json:"matchPolicy,omitempty"`
			NamespaceSelector *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
			ObjectSelector *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"object_selector" json:"objectSelector,omitempty"`
			ResourceRules *[]struct {
				ApiGroups     *[]string `tfsdk:"api_groups" json:"apiGroups,omitempty"`
				ApiVersions   *[]string `tfsdk:"api_versions" json:"apiVersions,omitempty"`
				Operations    *[]string `tfsdk:"operations" json:"operations,omitempty"`
				ResourceNames *[]string `tfsdk:"resource_names" json:"resourceNames,omitempty"`
				Resources     *[]string `tfsdk:"resources" json:"resources,omitempty"`
				Scope         *string   `tfsdk:"scope" json:"scope,omitempty"`
			} `tfsdk:"resource_rules" json:"resourceRules,omitempty"`
		} `tfsdk:"match_constraints" json:"matchConstraints,omitempty"`
		MatchImageReferences *[]struct {
			Expression *string `tfsdk:"expression" json:"expression,omitempty"`
			Glob       *string `tfsdk:"glob" json:"glob,omitempty"`
		} `tfsdk:"match_image_references" json:"matchImageReferences,omitempty"`
		ValidationActions        *[]string `tfsdk:"validation_actions" json:"validationActions,omitempty"`
		ValidationConfigurations *struct {
			MutateDigest *bool `tfsdk:"mutate_digest" json:"mutateDigest,omitempty"`
			Required     *bool `tfsdk:"required" json:"required,omitempty"`
			VerifyDigest *bool `tfsdk:"verify_digest" json:"verifyDigest,omitempty"`
		} `tfsdk:"validation_configurations" json:"validationConfigurations,omitempty"`
		Validations *[]struct {
			Expression        *string `tfsdk:"expression" json:"expression,omitempty"`
			Message           *string `tfsdk:"message" json:"message,omitempty"`
			MessageExpression *string `tfsdk:"message_expression" json:"messageExpression,omitempty"`
			Reason            *string `tfsdk:"reason" json:"reason,omitempty"`
		} `tfsdk:"validations" json:"validations,omitempty"`
		Variables *[]struct {
			Expression *string `tfsdk:"expression" json:"expression,omitempty"`
			Name       *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"variables" json:"variables,omitempty"`
		WebhookConfiguration *struct {
			TimeoutSeconds *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
		} `tfsdk:"webhook_configuration" json:"webhookConfiguration,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *PoliciesKyvernoIoImageValidatingPolicyV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_policies_kyverno_io_image_validating_policy_v1alpha1_manifest"
}

func (r *PoliciesKyvernoIoImageValidatingPolicyV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
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
				Description:         "ImageValidatingPolicySpec is the specification of the desired behavior of the ImageValidatingPolicy.",
				MarkdownDescription: "ImageValidatingPolicySpec is the specification of the desired behavior of the ImageValidatingPolicy.",
				Attributes: map[string]schema.Attribute{
					"attestations": schema.ListNestedAttribute{
						Description:         "Attestations provides a list of image metadata to verify",
						MarkdownDescription: "Attestations provides a list of image metadata to verify",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"intoto": schema.SingleNestedAttribute{
									Description:         "InToto defines the details of attestation attached using intoto format",
									MarkdownDescription: "InToto defines the details of attestation attached using intoto format",
									Attributes: map[string]schema.Attribute{
										"type": schema.StringAttribute{
											Description:         "Type defines the type of attestation contained within the statement.",
											MarkdownDescription: "Type defines the type of attestation contained within the statement.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"name": schema.StringAttribute{
									Description:         "Name is the name for this attestation. It is used to refer to the attestation in verification",
									MarkdownDescription: "Name is the name for this attestation. It is used to refer to the attestation in verification",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"referrer": schema.SingleNestedAttribute{
									Description:         "Referrer defines the details of attestation attached using OCI 1.1 format",
									MarkdownDescription: "Referrer defines the details of attestation attached using OCI 1.1 format",
									Attributes: map[string]schema.Attribute{
										"type": schema.StringAttribute{
											Description:         "Type defines the type of attestation attached to the image.",
											MarkdownDescription: "Type defines the type of attestation attached to the image.",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"attestors": schema.ListNestedAttribute{
						Description:         "Attestors provides a list of trusted authorities.",
						MarkdownDescription: "Attestors provides a list of trusted authorities.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"cosign": schema.SingleNestedAttribute{
									Description:         "Cosign defines attestor configuration for Cosign based signatures",
									MarkdownDescription: "Cosign defines attestor configuration for Cosign based signatures",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "Annotations are used for image verification. Every specified key-value pair must exist and match in the verified payload. The payload may contain other key-value pairs.",
											MarkdownDescription: "Annotations are used for image verification. Every specified key-value pair must exist and match in the verified payload. The payload may contain other key-value pairs.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"certificate": schema.SingleNestedAttribute{
											Description:         "Certificate defines the configuration for local signature verification",
											MarkdownDescription: "Certificate defines the configuration for local signature verification",
											Attributes: map[string]schema.Attribute{
												"cert": schema.SingleNestedAttribute{
													Description:         "Certificate is the to the public certificate for local signature verification.",
													MarkdownDescription: "Certificate is the to the public certificate for local signature verification.",
													Attributes: map[string]schema.Attribute{
														"expression": schema.StringAttribute{
															Description:         "Expression defines the a CEL expression input.",
															MarkdownDescription: "Expression defines the a CEL expression input.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"value": schema.StringAttribute{
															Description:         "Value defines the raw string input.",
															MarkdownDescription: "Value defines the raw string input.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"cert_chain": schema.SingleNestedAttribute{
													Description:         "CertificateChain is the list of CA certificates in PEM format which will be needed when building the certificate chain for the signing certificate. Must start with the parent intermediate CA certificate of the signing certificate and end with the root certificate",
													MarkdownDescription: "CertificateChain is the list of CA certificates in PEM format which will be needed when building the certificate chain for the signing certificate. Must start with the parent intermediate CA certificate of the signing certificate and end with the root certificate",
													Attributes: map[string]schema.Attribute{
														"expression": schema.StringAttribute{
															Description:         "Expression defines the a CEL expression input.",
															MarkdownDescription: "Expression defines the a CEL expression input.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"value": schema.StringAttribute{
															Description:         "Value defines the raw string input.",
															MarkdownDescription: "Value defines the raw string input.",
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

										"ctlog": schema.SingleNestedAttribute{
											Description:         "CTLog sets the configuration to verify the authority against a Rekor instance.",
											MarkdownDescription: "CTLog sets the configuration to verify the authority against a Rekor instance.",
											Attributes: map[string]schema.Attribute{
												"ct_log_pub_key": schema.StringAttribute{
													Description:         "CTLogPubKey, if set, is used to validate SCTs against a custom source.",
													MarkdownDescription: "CTLogPubKey, if set, is used to validate SCTs against a custom source.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"insecure_ignore_sct": schema.BoolAttribute{
													Description:         "IgnoreSCT defines whether to use the Signed Certificate Timestamp (SCT) log to check for a certificate timestamp. Default is false. Set to true if this was opted out during signing.",
													MarkdownDescription: "IgnoreSCT defines whether to use the Signed Certificate Timestamp (SCT) log to check for a certificate timestamp. Default is false. Set to true if this was opted out during signing.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"insecure_ignore_tlog": schema.BoolAttribute{
													Description:         "InsecureIgnoreTlog skips transparency log verification.",
													MarkdownDescription: "InsecureIgnoreTlog skips transparency log verification.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"rekor_pub_key": schema.StringAttribute{
													Description:         "RekorPubKey is an optional PEM-encoded public key to use for a custom Rekor. If set, this will be used to validate transparency log signatures from a custom Rekor.",
													MarkdownDescription: "RekorPubKey is an optional PEM-encoded public key to use for a custom Rekor. If set, this will be used to validate transparency log signatures from a custom Rekor.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"tsa_cert_chain": schema.StringAttribute{
													Description:         "TSACertChain, if set, is the PEM-encoded certificate chain file for the RFC3161 timestamp authority. Must contain the root CA certificate. Optionally may contain intermediate CA certificates, and may contain the leaf TSA certificate if not present in the timestamurce.",
													MarkdownDescription: "TSACertChain, if set, is the PEM-encoded certificate chain file for the RFC3161 timestamp authority. Must contain the root CA certificate. Optionally may contain intermediate CA certificates, and may contain the leaf TSA certificate if not present in the timestamurce.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"url": schema.StringAttribute{
													Description:         "URL sets the url to the rekor instance (by default the public rekor.sigstore.dev)",
													MarkdownDescription: "URL sets the url to the rekor instance (by default the public rekor.sigstore.dev)",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"key": schema.SingleNestedAttribute{
											Description:         "Key defines the type of key to validate the image.",
											MarkdownDescription: "Key defines the type of key to validate the image.",
											Attributes: map[string]schema.Attribute{
												"data": schema.StringAttribute{
													Description:         "Data contains the inline public key",
													MarkdownDescription: "Data contains the inline public key",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"expression": schema.StringAttribute{
													Description:         "Expression is a Expression expression that returns the public key.",
													MarkdownDescription: "Expression is a Expression expression that returns the public key.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"hash_algorithm": schema.StringAttribute{
													Description:         "HashAlgorithm specifues signature algorithm for public keys. Supported values are sha224, sha256, sha384 and sha512. Defaults to sha256.",
													MarkdownDescription: "HashAlgorithm specifues signature algorithm for public keys. Supported values are sha224, sha256, sha384 and sha512. Defaults to sha256.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"kms": schema.StringAttribute{
													Description:         "KMS contains the KMS url of the public key Supported formats differ based on the KMS system used.",
													MarkdownDescription: "KMS contains the KMS url of the public key Supported formats differ based on the KMS system used.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"keyless": schema.SingleNestedAttribute{
											Description:         "Keyless sets the configuration to verify the authority against a Fulcio instance.",
											MarkdownDescription: "Keyless sets the configuration to verify the authority against a Fulcio instance.",
											Attributes: map[string]schema.Attribute{
												"identities": schema.ListNestedAttribute{
													Description:         "Identities sets a list of identities.",
													MarkdownDescription: "Identities sets a list of identities.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"issuer": schema.StringAttribute{
																Description:         "Issuer defines the issuer for this identity.",
																MarkdownDescription: "Issuer defines the issuer for this identity.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"issuer_reg_exp": schema.StringAttribute{
																Description:         "IssuerRegExp specifies a regular expression to match the issuer for this identity.",
																MarkdownDescription: "IssuerRegExp specifies a regular expression to match the issuer for this identity.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"subject": schema.StringAttribute{
																Description:         "Subject defines the subject for this identity.",
																MarkdownDescription: "Subject defines the subject for this identity.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"subject_reg_exp": schema.StringAttribute{
																Description:         "SubjectRegExp specifies a regular expression to match the subject for this identity.",
																MarkdownDescription: "SubjectRegExp specifies a regular expression to match the subject for this identity.",
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

												"roots": schema.StringAttribute{
													Description:         "Roots is an optional set of PEM encoded trusted root certificates. If not provided, the system roots are used.",
													MarkdownDescription: "Roots is an optional set of PEM encoded trusted root certificates. If not provided, the system roots are used.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"source": schema.SingleNestedAttribute{
											Description:         "Sources sets the configuration to specify the sources from where to consume the signature and attestations.",
											MarkdownDescription: "Sources sets the configuration to specify the sources from where to consume the signature and attestations.",
											Attributes: map[string]schema.Attribute{
												"pull_secrets": schema.ListNestedAttribute{
													Description:         "SignaturePullSecrets is an optional list of references to secrets in the same namespace as the deploying resource for pulling any of the signatures used by this Source.",
													MarkdownDescription: "SignaturePullSecrets is an optional list of references to secrets in the same namespace as the deploying resource for pulling any of the signatures used by this Source.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"repository": schema.StringAttribute{
													Description:         "Repository defines the location from where to pull the signature / attestations.",
													MarkdownDescription: "Repository defines the location from where to pull the signature / attestations.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"tag_prefix": schema.StringAttribute{
													Description:         "TagPrefix is an optional prefix that signature and attestations have. This is the 'tag based discovery' and in the future once references are fully supported that should likely be the preferred way to handle these.",
													MarkdownDescription: "TagPrefix is an optional prefix that signature and attestations have. This is the 'tag based discovery' and in the future once references are fully supported that should likely be the preferred way to handle these.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"tuf": schema.SingleNestedAttribute{
											Description:         "TUF defines the configuration to fetch sigstore root",
											MarkdownDescription: "TUF defines the configuration to fetch sigstore root",
											Attributes: map[string]schema.Attribute{
												"mirror": schema.StringAttribute{
													Description:         "Mirror is the base URL of Sigstore TUF repository",
													MarkdownDescription: "Mirror is the base URL of Sigstore TUF repository",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"root": schema.SingleNestedAttribute{
													Description:         "Root defines the path or data of the trusted root",
													MarkdownDescription: "Root defines the path or data of the trusted root",
													Attributes: map[string]schema.Attribute{
														"data": schema.StringAttribute{
															Description:         "Data is the base64 encoded TUF root",
															MarkdownDescription: "Data is the base64 encoded TUF root",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"path": schema.StringAttribute{
															Description:         "Path is the URL or File location of the TUF root",
															MarkdownDescription: "Path is the URL or File location of the TUF root",
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

								"name": schema.StringAttribute{
									Description:         "Name is the name for this attestor. It is used to refer to the attestor in verification",
									MarkdownDescription: "Name is the name for this attestor. It is used to refer to the attestor in verification",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"notary": schema.SingleNestedAttribute{
									Description:         "Notary defines attestor configuration for Notary based signatures",
									MarkdownDescription: "Notary defines attestor configuration for Notary based signatures",
									Attributes: map[string]schema.Attribute{
										"certs": schema.SingleNestedAttribute{
											Description:         "Certs define the cert chain for Notary signature verification",
											MarkdownDescription: "Certs define the cert chain for Notary signature verification",
											Attributes: map[string]schema.Attribute{
												"expression": schema.StringAttribute{
													Description:         "Expression defines the a CEL expression input.",
													MarkdownDescription: "Expression defines the a CEL expression input.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "Value defines the raw string input.",
													MarkdownDescription: "Value defines the raw string input.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"tsa_certs": schema.SingleNestedAttribute{
											Description:         "TSACerts define the cert chain for verifying timestamps of notary signature",
											MarkdownDescription: "TSACerts define the cert chain for verifying timestamps of notary signature",
											Attributes: map[string]schema.Attribute{
												"expression": schema.StringAttribute{
													Description:         "Expression defines the a CEL expression input.",
													MarkdownDescription: "Expression defines the a CEL expression input.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "Value defines the raw string input.",
													MarkdownDescription: "Value defines the raw string input.",
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
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"audit_annotations": schema.ListNestedAttribute{
						Description:         "auditAnnotations contains CEL expressions which are used to produce audit annotations for the audit event of the API request. validations and auditAnnotations may not both be empty; a least one of validations or auditAnnotations is required.",
						MarkdownDescription: "auditAnnotations contains CEL expressions which are used to produce audit annotations for the audit event of the API request. validations and auditAnnotations may not both be empty; a least one of validations or auditAnnotations is required.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"key": schema.StringAttribute{
									Description:         "key specifies the audit annotation key. The audit annotation keys of a ValidatingAdmissionPolicy must be unique. The key must be a qualified name ([A-Za-z0-9][-A-Za-z0-9_.]*) no more than 63 bytes in length. The key is combined with the resource name of the ValidatingAdmissionPolicy to construct an audit annotation key: '{ValidatingAdmissionPolicy name}/{key}'. If an admission webhook uses the same resource name as this ValidatingAdmissionPolicy and the same audit annotation key, the annotation key will be identical. In this case, the first annotation written with the key will be included in the audit event and all subsequent annotations with the same key will be discarded. Required.",
									MarkdownDescription: "key specifies the audit annotation key. The audit annotation keys of a ValidatingAdmissionPolicy must be unique. The key must be a qualified name ([A-Za-z0-9][-A-Za-z0-9_.]*) no more than 63 bytes in length. The key is combined with the resource name of the ValidatingAdmissionPolicy to construct an audit annotation key: '{ValidatingAdmissionPolicy name}/{key}'. If an admission webhook uses the same resource name as this ValidatingAdmissionPolicy and the same audit annotation key, the annotation key will be identical. In this case, the first annotation written with the key will be included in the audit event and all subsequent annotations with the same key will be discarded. Required.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"value_expression": schema.StringAttribute{
									Description:         "valueExpression represents the expression which is evaluated by CEL to produce an audit annotation value. The expression must evaluate to either a string or null value. If the expression evaluates to a string, the audit annotation is included with the string value. If the expression evaluates to null or empty string the audit annotation will be omitted. The valueExpression may be no longer than 5kb in length. If the result of the valueExpression is more than 10kb in length, it will be truncated to 10kb. If multiple ValidatingAdmissionPolicyBinding resources match an API request, then the valueExpression will be evaluated for each binding. All unique values produced by the valueExpressions will be joined together in a comma-separated list. Required.",
									MarkdownDescription: "valueExpression represents the expression which is evaluated by CEL to produce an audit annotation value. The expression must evaluate to either a string or null value. If the expression evaluates to a string, the audit annotation is included with the string value. If the expression evaluates to null or empty string the audit annotation will be omitted. The valueExpression may be no longer than 5kb in length. If the result of the valueExpression is more than 10kb in length, it will be truncated to 10kb. If multiple ValidatingAdmissionPolicyBinding resources match an API request, then the valueExpression will be evaluated for each binding. All unique values produced by the valueExpressions will be joined together in a comma-separated list. Required.",
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

					"autogen": schema.SingleNestedAttribute{
						Description:         "AutogenConfiguration defines the configuration for the generation controller.",
						MarkdownDescription: "AutogenConfiguration defines the configuration for the generation controller.",
						Attributes: map[string]schema.Attribute{
							"pod_controllers": schema.SingleNestedAttribute{
								Description:         "PodControllers specifies whether to generate a pod controllers rules.",
								MarkdownDescription: "PodControllers specifies whether to generate a pod controllers rules.",
								Attributes: map[string]schema.Attribute{
									"controllers": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
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

					"credentials": schema.SingleNestedAttribute{
						Description:         "Credentials provides credentials that will be used for authentication with registry.",
						MarkdownDescription: "Credentials provides credentials that will be used for authentication with registry.",
						Attributes: map[string]schema.Attribute{
							"allow_insecure_registry": schema.BoolAttribute{
								Description:         "AllowInsecureRegistry allows insecure access to a registry.",
								MarkdownDescription: "AllowInsecureRegistry allows insecure access to a registry.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"providers": schema.ListAttribute{
								Description:         "Providers specifies a list of OCI Registry names, whose authentication providers are provided. It can be of one of these values: default,google,azure,amazon,github.",
								MarkdownDescription: "Providers specifies a list of OCI Registry names, whose authentication providers are provided. It can be of one of these values: default,google,azure,amazon,github.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secrets": schema.ListAttribute{
								Description:         "Secrets specifies a list of secrets that are provided for credentials. Secrets must live in the Kyverno namespace.",
								MarkdownDescription: "Secrets specifies a list of secrets that are provided for credentials. Secrets must live in the Kyverno namespace.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"evaluation": schema.SingleNestedAttribute{
						Description:         "EvaluationConfiguration defines the configuration for the policy evaluation.",
						MarkdownDescription: "EvaluationConfiguration defines the configuration for the policy evaluation.",
						Attributes: map[string]schema.Attribute{
							"admission": schema.SingleNestedAttribute{
								Description:         "Admission controls policy evaluation during admission.",
								MarkdownDescription: "Admission controls policy evaluation during admission.",
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Description:         "Enabled controls if rules are applied during admission. Optional. Default value is 'true'.",
										MarkdownDescription: "Enabled controls if rules are applied during admission. Optional. Default value is 'true'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"background": schema.SingleNestedAttribute{
								Description:         "Background controls policy evaluation during background scan.",
								MarkdownDescription: "Background controls policy evaluation during background scan.",
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Description:         "Enabled controls if rules are applied to existing resources during a background scan. Optional. Default value is 'true'. The value must be set to 'false' if the policy rule uses variables that are only available in the admission review request (e.g. user name).",
										MarkdownDescription: "Enabled controls if rules are applied to existing resources during a background scan. Optional. Default value is 'true'. The value must be set to 'false' if the policy rule uses variables that are only available in the admission review request (e.g. user name).",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"mode": schema.StringAttribute{
								Description:         "Mode is the mode of policy evaluation. Allowed values are 'Kubernetes' or 'JSON'. Optional. Default value is 'Kubernetes'.",
								MarkdownDescription: "Mode is the mode of policy evaluation. Allowed values are 'Kubernetes' or 'JSON'. Optional. Default value is 'Kubernetes'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"failure_policy": schema.StringAttribute{
						Description:         "FailurePolicy defines how to handle failures for the admission policy. Failures can occur from CEL expression parse errors, type check errors, runtime errors and invalid or mis-configured policy definitions or bindings.",
						MarkdownDescription: "FailurePolicy defines how to handle failures for the admission policy. Failures can occur from CEL expression parse errors, type check errors, runtime errors and invalid or mis-configured policy definitions or bindings.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Ignore", "Fail"),
						},
					},

					"images": schema.ListNestedAttribute{
						Description:         "ImageExtractors is a list of CEL expression to extract images from the resource",
						MarkdownDescription: "ImageExtractors is a list of CEL expression to extract images from the resource",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"expression": schema.StringAttribute{
									Description:         "Expression defines CEL expression to extract images from the resource.",
									MarkdownDescription: "Expression defines CEL expression to extract images from the resource.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name is the name for this imageList. It is used to refer to the images in verification block as images.<name>",
									MarkdownDescription: "Name is the name for this imageList. It is used to refer to the images in verification block as images.<name>",
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

					"match_conditions": schema.ListNestedAttribute{
						Description:         "MatchConditions is a list of conditions that must be met for a request to be validated. Match conditions filter requests that have already been matched by the rules, namespaceSelector, and objectSelector. An empty list of matchConditions matches all requests. There are a maximum of 64 match conditions allowed.",
						MarkdownDescription: "MatchConditions is a list of conditions that must be met for a request to be validated. Match conditions filter requests that have already been matched by the rules, namespaceSelector, and objectSelector. An empty list of matchConditions matches all requests. There are a maximum of 64 match conditions allowed.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"expression": schema.StringAttribute{
									Description:         "Expression represents the expression which will be evaluated by CEL. Must evaluate to bool. CEL expressions have access to the contents of the AdmissionRequest and Authorizer, organized into CEL variables: 'object' - The object from the incoming request. The value is null for DELETE requests. 'oldObject' - The existing object. The value is null for CREATE requests. 'request' - Attributes of the admission request(/pkg/apis/admission/types.go#AdmissionRequest). 'authorizer' - A CEL Authorizer. May be used to perform authorization checks for the principal (user or service account) of the request. See https://pkg.go.dev/k8s.io/apiserver/pkg/cel/library#Authz 'authorizer.requestResource' - A CEL ResourceCheck constructed from the 'authorizer' and configured with the request resource. Documentation on CEL: https://kubernetes.io/docs/reference/using-api/cel/ Required.",
									MarkdownDescription: "Expression represents the expression which will be evaluated by CEL. Must evaluate to bool. CEL expressions have access to the contents of the AdmissionRequest and Authorizer, organized into CEL variables: 'object' - The object from the incoming request. The value is null for DELETE requests. 'oldObject' - The existing object. The value is null for CREATE requests. 'request' - Attributes of the admission request(/pkg/apis/admission/types.go#AdmissionRequest). 'authorizer' - A CEL Authorizer. May be used to perform authorization checks for the principal (user or service account) of the request. See https://pkg.go.dev/k8s.io/apiserver/pkg/cel/library#Authz 'authorizer.requestResource' - A CEL ResourceCheck constructed from the 'authorizer' and configured with the request resource. Documentation on CEL: https://kubernetes.io/docs/reference/using-api/cel/ Required.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name is an identifier for this match condition, used for strategic merging of MatchConditions, as well as providing an identifier for logging purposes. A good name should be descriptive of the associated expression. Name must be a qualified name consisting of alphanumeric characters, '-', '_' or '.', and must start and end with an alphanumeric character (e.g. 'MyName', or 'my.name', or '123-abc', regex used for validation is '([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9]') with an optional DNS subdomain prefix and '/' (e.g. 'example.com/MyName') Required.",
									MarkdownDescription: "Name is an identifier for this match condition, used for strategic merging of MatchConditions, as well as providing an identifier for logging purposes. A good name should be descriptive of the associated expression. Name must be a qualified name consisting of alphanumeric characters, '-', '_' or '.', and must start and end with an alphanumeric character (e.g. 'MyName', or 'my.name', or '123-abc', regex used for validation is '([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9]') with an optional DNS subdomain prefix and '/' (e.g. 'example.com/MyName') Required.",
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

					"match_constraints": schema.SingleNestedAttribute{
						Description:         "MatchConstraints specifies what resources this policy is designed to validate.",
						MarkdownDescription: "MatchConstraints specifies what resources this policy is designed to validate.",
						Attributes: map[string]schema.Attribute{
							"exclude_resource_rules": schema.ListNestedAttribute{
								Description:         "ExcludeResourceRules describes what operations on what resources/subresources the ValidatingAdmissionPolicy should not care about. The exclude rules take precedence over include rules (if a resource matches both, it is excluded)",
								MarkdownDescription: "ExcludeResourceRules describes what operations on what resources/subresources the ValidatingAdmissionPolicy should not care about. The exclude rules take precedence over include rules (if a resource matches both, it is excluded)",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"api_groups": schema.ListAttribute{
											Description:         "APIGroups is the API groups the resources belong to. '*' is all groups. If '*' is present, the length of the slice must be one. Required.",
											MarkdownDescription: "APIGroups is the API groups the resources belong to. '*' is all groups. If '*' is present, the length of the slice must be one. Required.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"api_versions": schema.ListAttribute{
											Description:         "APIVersions is the API versions the resources belong to. '*' is all versions. If '*' is present, the length of the slice must be one. Required.",
											MarkdownDescription: "APIVersions is the API versions the resources belong to. '*' is all versions. If '*' is present, the length of the slice must be one. Required.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"operations": schema.ListAttribute{
											Description:         "Operations is the operations the admission hook cares about - CREATE, UPDATE, DELETE, CONNECT or * for all of those operations and any future admission operations that are added. If '*' is present, the length of the slice must be one. Required.",
											MarkdownDescription: "Operations is the operations the admission hook cares about - CREATE, UPDATE, DELETE, CONNECT or * for all of those operations and any future admission operations that are added. If '*' is present, the length of the slice must be one. Required.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"resource_names": schema.ListAttribute{
											Description:         "ResourceNames is an optional white list of names that the rule applies to. An empty set means that everything is allowed.",
											MarkdownDescription: "ResourceNames is an optional white list of names that the rule applies to. An empty set means that everything is allowed.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"resources": schema.ListAttribute{
											Description:         "Resources is a list of resources this rule applies to. For example: 'pods' means pods. 'pods/log' means the log subresource of pods. '*' means all resources, but not subresources. 'pods/*' means all subresources of pods. '*/scale' means all scale subresources. '*/*' means all resources and their subresources. If wildcard is present, the validation rule will ensure resources do not overlap with each other. Depending on the enclosing object, subresources might not be allowed. Required.",
											MarkdownDescription: "Resources is a list of resources this rule applies to. For example: 'pods' means pods. 'pods/log' means the log subresource of pods. '*' means all resources, but not subresources. 'pods/*' means all subresources of pods. '*/scale' means all scale subresources. '*/*' means all resources and their subresources. If wildcard is present, the validation rule will ensure resources do not overlap with each other. Depending on the enclosing object, subresources might not be allowed. Required.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"scope": schema.StringAttribute{
											Description:         "scope specifies the scope of this rule. Valid values are 'Cluster', 'Namespaced', and '*' 'Cluster' means that only cluster-scoped resources will match this rule. Namespace API objects are cluster-scoped. 'Namespaced' means that only namespaced resources will match this rule. '*' means that there are no scope restrictions. Subresources match the scope of their parent resource. Default is '*'.",
											MarkdownDescription: "scope specifies the scope of this rule. Valid values are 'Cluster', 'Namespaced', and '*' 'Cluster' means that only cluster-scoped resources will match this rule. Namespace API objects are cluster-scoped. 'Namespaced' means that only namespaced resources will match this rule. '*' means that there are no scope restrictions. Subresources match the scope of their parent resource. Default is '*'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"match_policy": schema.StringAttribute{
								Description:         "matchPolicy defines how the 'MatchResources' list is used to match incoming requests. Allowed values are 'Exact' or 'Equivalent'. - Exact: match a request only if it exactly matches a specified rule. For example, if deployments can be modified via apps/v1, apps/v1beta1, and extensions/v1beta1, but 'rules' only included 'apiGroups:['apps'], apiVersions:['v1'], resources: ['deployments']', a request to apps/v1beta1 or extensions/v1beta1 would not be sent to the ValidatingAdmissionPolicy. - Equivalent: match a request if modifies a resource listed in rules, even via another API group or version. For example, if deployments can be modified via apps/v1, apps/v1beta1, and extensions/v1beta1, and 'rules' only included 'apiGroups:['apps'], apiVersions:['v1'], resources: ['deployments']', a request to apps/v1beta1 or extensions/v1beta1 would be converted to apps/v1 and sent to the ValidatingAdmissionPolicy. Defaults to 'Equivalent'",
								MarkdownDescription: "matchPolicy defines how the 'MatchResources' list is used to match incoming requests. Allowed values are 'Exact' or 'Equivalent'. - Exact: match a request only if it exactly matches a specified rule. For example, if deployments can be modified via apps/v1, apps/v1beta1, and extensions/v1beta1, but 'rules' only included 'apiGroups:['apps'], apiVersions:['v1'], resources: ['deployments']', a request to apps/v1beta1 or extensions/v1beta1 would not be sent to the ValidatingAdmissionPolicy. - Equivalent: match a request if modifies a resource listed in rules, even via another API group or version. For example, if deployments can be modified via apps/v1, apps/v1beta1, and extensions/v1beta1, and 'rules' only included 'apiGroups:['apps'], apiVersions:['v1'], resources: ['deployments']', a request to apps/v1beta1 or extensions/v1beta1 would be converted to apps/v1 and sent to the ValidatingAdmissionPolicy. Defaults to 'Equivalent'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace_selector": schema.SingleNestedAttribute{
								Description:         "NamespaceSelector decides whether to run the admission control policy on an object based on whether the namespace for that object matches the selector. If the object itself is a namespace, the matching is performed on object.metadata.labels. If the object is another cluster scoped resource, it never skips the policy. For example, to run the webhook on any objects whose namespace is not associated with 'runlevel' of '0' or '1'; you will set the selector as follows: 'namespaceSelector': { 'matchExpressions': [ { 'key': 'runlevel', 'operator': 'NotIn', 'values': [ '0', '1' ] } ] } If instead you want to only run the policy on any objects whose namespace is associated with the 'environment' of 'prod' or 'staging'; you will set the selector as follows: 'namespaceSelector': { 'matchExpressions': [ { 'key': 'environment', 'operator': 'In', 'values': [ 'prod', 'staging' ] } ] } See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more examples of label selectors. Default to the empty LabelSelector, which matches everything.",
								MarkdownDescription: "NamespaceSelector decides whether to run the admission control policy on an object based on whether the namespace for that object matches the selector. If the object itself is a namespace, the matching is performed on object.metadata.labels. If the object is another cluster scoped resource, it never skips the policy. For example, to run the webhook on any objects whose namespace is not associated with 'runlevel' of '0' or '1'; you will set the selector as follows: 'namespaceSelector': { 'matchExpressions': [ { 'key': 'runlevel', 'operator': 'NotIn', 'values': [ '0', '1' ] } ] } If instead you want to only run the policy on any objects whose namespace is associated with the 'environment' of 'prod' or 'staging'; you will set the selector as follows: 'namespaceSelector': { 'matchExpressions': [ { 'key': 'environment', 'operator': 'In', 'values': [ 'prod', 'staging' ] } ] } See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more examples of label selectors. Default to the empty LabelSelector, which matches everything.",
								Attributes: map[string]schema.Attribute{
									"match_expressions": schema.ListNestedAttribute{
										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "key is the label key that the selector applies to.",
													MarkdownDescription: "key is the label key that the selector applies to.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"operator": schema.StringAttribute{
													Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
													MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"values": schema.ListAttribute{
													Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
													MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"match_labels": schema.MapAttribute{
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"object_selector": schema.SingleNestedAttribute{
								Description:         "ObjectSelector decides whether to run the validation based on if the object has matching labels. objectSelector is evaluated against both the oldObject and newObject that would be sent to the cel validation, and is considered to match if either object matches the selector. A null object (oldObject in the case of create, or newObject in the case of delete) or an object that cannot have labels (like a DeploymentRollback or a PodProxyOptions object) is not considered to match. Use the object selector only if the webhook is opt-in, because end users may skip the admission webhook by setting the labels. Default to the empty LabelSelector, which matches everything.",
								MarkdownDescription: "ObjectSelector decides whether to run the validation based on if the object has matching labels. objectSelector is evaluated against both the oldObject and newObject that would be sent to the cel validation, and is considered to match if either object matches the selector. A null object (oldObject in the case of create, or newObject in the case of delete) or an object that cannot have labels (like a DeploymentRollback or a PodProxyOptions object) is not considered to match. Use the object selector only if the webhook is opt-in, because end users may skip the admission webhook by setting the labels. Default to the empty LabelSelector, which matches everything.",
								Attributes: map[string]schema.Attribute{
									"match_expressions": schema.ListNestedAttribute{
										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "key is the label key that the selector applies to.",
													MarkdownDescription: "key is the label key that the selector applies to.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"operator": schema.StringAttribute{
													Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
													MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"values": schema.ListAttribute{
													Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
													MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"match_labels": schema.MapAttribute{
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"resource_rules": schema.ListNestedAttribute{
								Description:         "ResourceRules describes what operations on what resources/subresources the ValidatingAdmissionPolicy matches. The policy cares about an operation if it matches _any_ Rule.",
								MarkdownDescription: "ResourceRules describes what operations on what resources/subresources the ValidatingAdmissionPolicy matches. The policy cares about an operation if it matches _any_ Rule.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"api_groups": schema.ListAttribute{
											Description:         "APIGroups is the API groups the resources belong to. '*' is all groups. If '*' is present, the length of the slice must be one. Required.",
											MarkdownDescription: "APIGroups is the API groups the resources belong to. '*' is all groups. If '*' is present, the length of the slice must be one. Required.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"api_versions": schema.ListAttribute{
											Description:         "APIVersions is the API versions the resources belong to. '*' is all versions. If '*' is present, the length of the slice must be one. Required.",
											MarkdownDescription: "APIVersions is the API versions the resources belong to. '*' is all versions. If '*' is present, the length of the slice must be one. Required.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"operations": schema.ListAttribute{
											Description:         "Operations is the operations the admission hook cares about - CREATE, UPDATE, DELETE, CONNECT or * for all of those operations and any future admission operations that are added. If '*' is present, the length of the slice must be one. Required.",
											MarkdownDescription: "Operations is the operations the admission hook cares about - CREATE, UPDATE, DELETE, CONNECT or * for all of those operations and any future admission operations that are added. If '*' is present, the length of the slice must be one. Required.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"resource_names": schema.ListAttribute{
											Description:         "ResourceNames is an optional white list of names that the rule applies to. An empty set means that everything is allowed.",
											MarkdownDescription: "ResourceNames is an optional white list of names that the rule applies to. An empty set means that everything is allowed.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"resources": schema.ListAttribute{
											Description:         "Resources is a list of resources this rule applies to. For example: 'pods' means pods. 'pods/log' means the log subresource of pods. '*' means all resources, but not subresources. 'pods/*' means all subresources of pods. '*/scale' means all scale subresources. '*/*' means all resources and their subresources. If wildcard is present, the validation rule will ensure resources do not overlap with each other. Depending on the enclosing object, subresources might not be allowed. Required.",
											MarkdownDescription: "Resources is a list of resources this rule applies to. For example: 'pods' means pods. 'pods/log' means the log subresource of pods. '*' means all resources, but not subresources. 'pods/*' means all subresources of pods. '*/scale' means all scale subresources. '*/*' means all resources and their subresources. If wildcard is present, the validation rule will ensure resources do not overlap with each other. Depending on the enclosing object, subresources might not be allowed. Required.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"scope": schema.StringAttribute{
											Description:         "scope specifies the scope of this rule. Valid values are 'Cluster', 'Namespaced', and '*' 'Cluster' means that only cluster-scoped resources will match this rule. Namespace API objects are cluster-scoped. 'Namespaced' means that only namespaced resources will match this rule. '*' means that there are no scope restrictions. Subresources match the scope of their parent resource. Default is '*'.",
											MarkdownDescription: "scope specifies the scope of this rule. Valid values are 'Cluster', 'Namespaced', and '*' 'Cluster' means that only cluster-scoped resources will match this rule. Namespace API objects are cluster-scoped. 'Namespaced' means that only namespaced resources will match this rule. '*' means that there are no scope restrictions. Subresources match the scope of their parent resource. Default is '*'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
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

					"match_image_references": schema.ListNestedAttribute{
						Description:         "MatchImageReferences is a list of Glob and CELExpressions to match images. Any image that matches one of the rules is considered for validation Any image that does not match a rule is skipped, even when they are passed as arguments to image verification functions",
						MarkdownDescription: "MatchImageReferences is a list of Glob and CELExpressions to match images. Any image that matches one of the rules is considered for validation Any image that does not match a rule is skipped, even when they are passed as arguments to image verification functions",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"expression": schema.StringAttribute{
									Description:         "Expression defines CEL Expressions for matching images",
									MarkdownDescription: "Expression defines CEL Expressions for matching images",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"glob": schema.StringAttribute{
									Description:         "Glob defines a globbing pattern for matching images",
									MarkdownDescription: "Glob defines a globbing pattern for matching images",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"validation_actions": schema.ListAttribute{
						Description:         "ValidationAction specifies the action to be taken when the matched resource violates the policy. Required.",
						MarkdownDescription: "ValidationAction specifies the action to be taken when the matched resource violates the policy. Required.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"validation_configurations": schema.SingleNestedAttribute{
						Description:         "ValidationConfigurations defines settings for mutating and verifying image digests, and enforcing image verification through signatures.",
						MarkdownDescription: "ValidationConfigurations defines settings for mutating and verifying image digests, and enforcing image verification through signatures.",
						Attributes: map[string]schema.Attribute{
							"mutate_digest": schema.BoolAttribute{
								Description:         "MutateDigest enables replacement of image tags with digests. Defaults to true.",
								MarkdownDescription: "MutateDigest enables replacement of image tags with digests. Defaults to true.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"required": schema.BoolAttribute{
								Description:         "Required validates that images are verified, i.e., have passed a signature or attestation check.",
								MarkdownDescription: "Required validates that images are verified, i.e., have passed a signature or attestation check.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"verify_digest": schema.BoolAttribute{
								Description:         "VerifyDigest validates that images have a digest.",
								MarkdownDescription: "VerifyDigest validates that images have a digest.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"validations": schema.ListNestedAttribute{
						Description:         "Validations contain CEL expressions which is used to apply the image validation checks.",
						MarkdownDescription: "Validations contain CEL expressions which is used to apply the image validation checks.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"expression": schema.StringAttribute{
									Description:         "Expression represents the expression which will be evaluated by CEL. ref: https://github.com/google/cel-spec CEL expressions have access to the contents of the API request/response, organized into CEL variables as well as some other useful variables: - 'object' - The object from the incoming request. The value is null for DELETE requests. - 'oldObject' - The existing object. The value is null for CREATE requests. - 'request' - Attributes of the API request([ref](/pkg/apis/admission/types.go#AdmissionRequest)). - 'params' - Parameter resource referred to by the policy binding being evaluated. Only populated if the policy has a ParamKind. - 'namespaceObject' - The namespace object that the incoming object belongs to. The value is null for cluster-scoped resources. - 'variables' - Map of composited variables, from its name to its lazily evaluated value. For example, a variable named 'foo' can be accessed as 'variables.foo'. - 'authorizer' - A CEL Authorizer. May be used to perform authorization checks for the principal (user or service account) of the request. See https://pkg.go.dev/k8s.io/apiserver/pkg/cel/library#Authz - 'authorizer.requestResource' - A CEL ResourceCheck constructed from the 'authorizer' and configured with the request resource. The 'apiVersion', 'kind', 'metadata.name' and 'metadata.generateName' are always accessible from the root of the object. No other metadata properties are accessible. Only property names of the form '[a-zA-Z_.-/][a-zA-Z0-9_.-/]*' are accessible. Accessible property names are escaped according to the following rules when accessed in the expression: - '__' escapes to '__underscores__' - '.' escapes to '__dot__' - '-' escapes to '__dash__' - '/' escapes to '__slash__' - Property names that exactly match a CEL RESERVED keyword escape to '__{keyword}__'. The keywords are: 'true', 'false', 'null', 'in', 'as', 'break', 'const', 'continue', 'else', 'for', 'function', 'if', 'import', 'let', 'loop', 'package', 'namespace', 'return'. Examples: - Expression accessing a property named 'namespace': {'Expression': 'object.__namespace__ > 0'} - Expression accessing a property named 'x-prop': {'Expression': 'object.x__dash__prop > 0'} - Expression accessing a property named 'redact__d': {'Expression': 'object.redact__underscores__d > 0'} Equality on arrays with list type of 'set' or 'map' ignores element order, i.e. [1, 2] == [2, 1]. Concatenation on arrays with x-kubernetes-list-type use the semantics of the list type: - 'set': 'X + Y' performs a union where the array positions of all elements in 'X' are preserved and non-intersecting elements in 'Y' are appended, retaining their partial order. - 'map': 'X + Y' performs a merge where the array positions of all keys in 'X' are preserved but the values are overwritten by values in 'Y' when the key sets of 'X' and 'Y' intersect. Elements in 'Y' with non-intersecting keys are appended, retaining their partial order. Required.",
									MarkdownDescription: "Expression represents the expression which will be evaluated by CEL. ref: https://github.com/google/cel-spec CEL expressions have access to the contents of the API request/response, organized into CEL variables as well as some other useful variables: - 'object' - The object from the incoming request. The value is null for DELETE requests. - 'oldObject' - The existing object. The value is null for CREATE requests. - 'request' - Attributes of the API request([ref](/pkg/apis/admission/types.go#AdmissionRequest)). - 'params' - Parameter resource referred to by the policy binding being evaluated. Only populated if the policy has a ParamKind. - 'namespaceObject' - The namespace object that the incoming object belongs to. The value is null for cluster-scoped resources. - 'variables' - Map of composited variables, from its name to its lazily evaluated value. For example, a variable named 'foo' can be accessed as 'variables.foo'. - 'authorizer' - A CEL Authorizer. May be used to perform authorization checks for the principal (user or service account) of the request. See https://pkg.go.dev/k8s.io/apiserver/pkg/cel/library#Authz - 'authorizer.requestResource' - A CEL ResourceCheck constructed from the 'authorizer' and configured with the request resource. The 'apiVersion', 'kind', 'metadata.name' and 'metadata.generateName' are always accessible from the root of the object. No other metadata properties are accessible. Only property names of the form '[a-zA-Z_.-/][a-zA-Z0-9_.-/]*' are accessible. Accessible property names are escaped according to the following rules when accessed in the expression: - '__' escapes to '__underscores__' - '.' escapes to '__dot__' - '-' escapes to '__dash__' - '/' escapes to '__slash__' - Property names that exactly match a CEL RESERVED keyword escape to '__{keyword}__'. The keywords are: 'true', 'false', 'null', 'in', 'as', 'break', 'const', 'continue', 'else', 'for', 'function', 'if', 'import', 'let', 'loop', 'package', 'namespace', 'return'. Examples: - Expression accessing a property named 'namespace': {'Expression': 'object.__namespace__ > 0'} - Expression accessing a property named 'x-prop': {'Expression': 'object.x__dash__prop > 0'} - Expression accessing a property named 'redact__d': {'Expression': 'object.redact__underscores__d > 0'} Equality on arrays with list type of 'set' or 'map' ignores element order, i.e. [1, 2] == [2, 1]. Concatenation on arrays with x-kubernetes-list-type use the semantics of the list type: - 'set': 'X + Y' performs a union where the array positions of all elements in 'X' are preserved and non-intersecting elements in 'Y' are appended, retaining their partial order. - 'map': 'X + Y' performs a merge where the array positions of all keys in 'X' are preserved but the values are overwritten by values in 'Y' when the key sets of 'X' and 'Y' intersect. Elements in 'Y' with non-intersecting keys are appended, retaining their partial order. Required.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"message": schema.StringAttribute{
									Description:         "Message represents the message displayed when validation fails. The message is required if the Expression contains line breaks. The message must not contain line breaks. If unset, the message is 'failed rule: {Rule}'. e.g. 'must be a URL with the host matching spec.host' If the Expression contains line breaks. Message is required. The message must not contain line breaks. If unset, the message is 'failed Expression: {Expression}'.",
									MarkdownDescription: "Message represents the message displayed when validation fails. The message is required if the Expression contains line breaks. The message must not contain line breaks. If unset, the message is 'failed rule: {Rule}'. e.g. 'must be a URL with the host matching spec.host' If the Expression contains line breaks. Message is required. The message must not contain line breaks. If unset, the message is 'failed Expression: {Expression}'.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"message_expression": schema.StringAttribute{
									Description:         "messageExpression declares a CEL expression that evaluates to the validation failure message that is returned when this rule fails. Since messageExpression is used as a failure message, it must evaluate to a string. If both message and messageExpression are present on a validation, then messageExpression will be used if validation fails. If messageExpression results in a runtime error, the runtime error is logged, and the validation failure message is produced as if the messageExpression field were unset. If messageExpression evaluates to an empty string, a string with only spaces, or a string that contains line breaks, then the validation failure message will also be produced as if the messageExpression field were unset, and the fact that messageExpression produced an empty string/string with only spaces/string with line breaks will be logged. messageExpression has access to all the same variables as the 'expression' except for 'authorizer' and 'authorizer.requestResource'. Example: 'object.x must be less than max ('+string(params.max)+')'",
									MarkdownDescription: "messageExpression declares a CEL expression that evaluates to the validation failure message that is returned when this rule fails. Since messageExpression is used as a failure message, it must evaluate to a string. If both message and messageExpression are present on a validation, then messageExpression will be used if validation fails. If messageExpression results in a runtime error, the runtime error is logged, and the validation failure message is produced as if the messageExpression field were unset. If messageExpression evaluates to an empty string, a string with only spaces, or a string that contains line breaks, then the validation failure message will also be produced as if the messageExpression field were unset, and the fact that messageExpression produced an empty string/string with only spaces/string with line breaks will be logged. messageExpression has access to all the same variables as the 'expression' except for 'authorizer' and 'authorizer.requestResource'. Example: 'object.x must be less than max ('+string(params.max)+')'",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"reason": schema.StringAttribute{
									Description:         "Reason represents a machine-readable description of why this validation failed. If this is the first validation in the list to fail, this reason, as well as the corresponding HTTP response code, are used in the HTTP response to the client. The currently supported reasons are: 'Unauthorized', 'Forbidden', 'Invalid', 'RequestEntityTooLarge'. If not set, StatusReasonInvalid is used in the response to the client.",
									MarkdownDescription: "Reason represents a machine-readable description of why this validation failed. If this is the first validation in the list to fail, this reason, as well as the corresponding HTTP response code, are used in the HTTP response to the client. The currently supported reasons are: 'Unauthorized', 'Forbidden', 'Invalid', 'RequestEntityTooLarge'. If not set, StatusReasonInvalid is used in the response to the client.",
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

					"variables": schema.ListNestedAttribute{
						Description:         "Variables contain definitions of variables that can be used in composition of other expressions. Each variable is defined as a named CEL expression.",
						MarkdownDescription: "Variables contain definitions of variables that can be used in composition of other expressions. Each variable is defined as a named CEL expression.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"expression": schema.StringAttribute{
									Description:         "Expression is the expression that will be evaluated as the value of the variable. The CEL expression has access to the same identifiers as the CEL expressions in Validation.",
									MarkdownDescription: "Expression is the expression that will be evaluated as the value of the variable. The CEL expression has access to the same identifiers as the CEL expressions in Validation.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name is the name of the variable. The name must be a valid CEL identifier and unique among all variables. The variable can be accessed in other expressions through 'variables' For example, if name is 'foo', the variable will be available as 'variables.foo'",
									MarkdownDescription: "Name is the name of the variable. The name must be a valid CEL identifier and unique among all variables. The variable can be accessed in other expressions through 'variables' For example, if name is 'foo', the variable will be available as 'variables.foo'",
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

					"webhook_configuration": schema.SingleNestedAttribute{
						Description:         "WebhookConfiguration defines the configuration for the webhook.",
						MarkdownDescription: "WebhookConfiguration defines the configuration for the webhook.",
						Attributes: map[string]schema.Attribute{
							"timeout_seconds": schema.Int64Attribute{
								Description:         "TimeoutSeconds specifies the maximum time in seconds allowed to apply this policy. After the configured time expires, the admission request may fail, or may simply ignore the policy results, based on the failure policy. The default timeout is 10s, the value must be between 1 and 30 seconds.",
								MarkdownDescription: "TimeoutSeconds specifies the maximum time in seconds allowed to apply this policy. After the configured time expires, the admission request may fail, or may simply ignore the policy results, based on the failure policy. The default timeout is 10s, the value must be between 1 and 30 seconds.",
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
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *PoliciesKyvernoIoImageValidatingPolicyV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_policies_kyverno_io_image_validating_policy_v1alpha1_manifest")

	var model PoliciesKyvernoIoImageValidatingPolicyV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("policies.kyverno.io/v1alpha1")
	model.Kind = pointer.String("ImageValidatingPolicy")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
