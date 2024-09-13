/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package acmpca_services_k8s_aws_v1alpha1

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
	_ datasource.DataSource = &AcmpcaServicesK8SAwsCertificateV1Alpha1Manifest{}
)

func NewAcmpcaServicesK8SAwsCertificateV1Alpha1Manifest() datasource.DataSource {
	return &AcmpcaServicesK8SAwsCertificateV1Alpha1Manifest{}
}

type AcmpcaServicesK8SAwsCertificateV1Alpha1Manifest struct{}

type AcmpcaServicesK8SAwsCertificateV1Alpha1ManifestData struct {
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
		ApiPassthrough *struct {
			Extensions *struct {
				CertificatePolicies *[]struct {
					CertPolicyID     *string `tfsdk:"cert_policy_id" json:"certPolicyID,omitempty"`
					PolicyQualifiers *[]struct {
						PolicyQualifierID *string `tfsdk:"policy_qualifier_id" json:"policyQualifierID,omitempty"`
						Qualifier         *struct {
							CpsURI *string `tfsdk:"cps_uri" json:"cpsURI,omitempty"`
						} `tfsdk:"qualifier" json:"qualifier,omitempty"`
					} `tfsdk:"policy_qualifiers" json:"policyQualifiers,omitempty"`
				} `tfsdk:"certificate_policies" json:"certificatePolicies,omitempty"`
				CustomExtensions *[]struct {
					Critical         *bool   `tfsdk:"critical" json:"critical,omitempty"`
					ObjectIdentifier *string `tfsdk:"object_identifier" json:"objectIdentifier,omitempty"`
					Value            *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"custom_extensions" json:"customExtensions,omitempty"`
				ExtendedKeyUsage *[]struct {
					ExtendedKeyUsageObjectIdentifier *string `tfsdk:"extended_key_usage_object_identifier" json:"extendedKeyUsageObjectIdentifier,omitempty"`
					ExtendedKeyUsageType             *string `tfsdk:"extended_key_usage_type" json:"extendedKeyUsageType,omitempty"`
				} `tfsdk:"extended_key_usage" json:"extendedKeyUsage,omitempty"`
				KeyUsage *struct {
					CrlSign          *bool `tfsdk:"crl_sign" json:"crlSign,omitempty"`
					DataEncipherment *bool `tfsdk:"data_encipherment" json:"dataEncipherment,omitempty"`
					DecipherOnly     *bool `tfsdk:"decipher_only" json:"decipherOnly,omitempty"`
					DigitalSignature *bool `tfsdk:"digital_signature" json:"digitalSignature,omitempty"`
					EncipherOnly     *bool `tfsdk:"encipher_only" json:"encipherOnly,omitempty"`
					KeyAgreement     *bool `tfsdk:"key_agreement" json:"keyAgreement,omitempty"`
					KeyCertSign      *bool `tfsdk:"key_cert_sign" json:"keyCertSign,omitempty"`
					KeyEncipherment  *bool `tfsdk:"key_encipherment" json:"keyEncipherment,omitempty"`
					NonRepudiation   *bool `tfsdk:"non_repudiation" json:"nonRepudiation,omitempty"`
				} `tfsdk:"key_usage" json:"keyUsage,omitempty"`
				SubjectAlternativeNames *[]struct {
					DirectoryName *struct {
						CommonName       *string `tfsdk:"common_name" json:"commonName,omitempty"`
						Country          *string `tfsdk:"country" json:"country,omitempty"`
						CustomAttributes *[]struct {
							ObjectIdentifier *string `tfsdk:"object_identifier" json:"objectIdentifier,omitempty"`
							Value            *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"custom_attributes" json:"customAttributes,omitempty"`
						DistinguishedNameQualifier *string `tfsdk:"distinguished_name_qualifier" json:"distinguishedNameQualifier,omitempty"`
						GenerationQualifier        *string `tfsdk:"generation_qualifier" json:"generationQualifier,omitempty"`
						GivenName                  *string `tfsdk:"given_name" json:"givenName,omitempty"`
						Initials                   *string `tfsdk:"initials" json:"initials,omitempty"`
						Locality                   *string `tfsdk:"locality" json:"locality,omitempty"`
						Organization               *string `tfsdk:"organization" json:"organization,omitempty"`
						OrganizationalUnit         *string `tfsdk:"organizational_unit" json:"organizationalUnit,omitempty"`
						Pseudonym                  *string `tfsdk:"pseudonym" json:"pseudonym,omitempty"`
						SerialNumber               *string `tfsdk:"serial_number" json:"serialNumber,omitempty"`
						State                      *string `tfsdk:"state" json:"state,omitempty"`
						Surname                    *string `tfsdk:"surname" json:"surname,omitempty"`
						Title                      *string `tfsdk:"title" json:"title,omitempty"`
					} `tfsdk:"directory_name" json:"directoryName,omitempty"`
					DnsName      *string `tfsdk:"dns_name" json:"dnsName,omitempty"`
					EdiPartyName *struct {
						NameAssigner *string `tfsdk:"name_assigner" json:"nameAssigner,omitempty"`
						PartyName    *string `tfsdk:"party_name" json:"partyName,omitempty"`
					} `tfsdk:"edi_party_name" json:"ediPartyName,omitempty"`
					IpAddress *string `tfsdk:"ip_address" json:"ipAddress,omitempty"`
					OtherName *struct {
						TypeID *string `tfsdk:"type_id" json:"typeID,omitempty"`
						Value  *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"other_name" json:"otherName,omitempty"`
					RegisteredID              *string `tfsdk:"registered_id" json:"registeredID,omitempty"`
					Rfc822Name                *string `tfsdk:"rfc822_name" json:"rfc822Name,omitempty"`
					UniformResourceIdentifier *string `tfsdk:"uniform_resource_identifier" json:"uniformResourceIdentifier,omitempty"`
				} `tfsdk:"subject_alternative_names" json:"subjectAlternativeNames,omitempty"`
			} `tfsdk:"extensions" json:"extensions,omitempty"`
			Subject *struct {
				CommonName       *string `tfsdk:"common_name" json:"commonName,omitempty"`
				Country          *string `tfsdk:"country" json:"country,omitempty"`
				CustomAttributes *[]struct {
					ObjectIdentifier *string `tfsdk:"object_identifier" json:"objectIdentifier,omitempty"`
					Value            *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"custom_attributes" json:"customAttributes,omitempty"`
				DistinguishedNameQualifier *string `tfsdk:"distinguished_name_qualifier" json:"distinguishedNameQualifier,omitempty"`
				GenerationQualifier        *string `tfsdk:"generation_qualifier" json:"generationQualifier,omitempty"`
				GivenName                  *string `tfsdk:"given_name" json:"givenName,omitempty"`
				Initials                   *string `tfsdk:"initials" json:"initials,omitempty"`
				Locality                   *string `tfsdk:"locality" json:"locality,omitempty"`
				Organization               *string `tfsdk:"organization" json:"organization,omitempty"`
				OrganizationalUnit         *string `tfsdk:"organizational_unit" json:"organizationalUnit,omitempty"`
				Pseudonym                  *string `tfsdk:"pseudonym" json:"pseudonym,omitempty"`
				SerialNumber               *string `tfsdk:"serial_number" json:"serialNumber,omitempty"`
				State                      *string `tfsdk:"state" json:"state,omitempty"`
				Surname                    *string `tfsdk:"surname" json:"surname,omitempty"`
				Title                      *string `tfsdk:"title" json:"title,omitempty"`
			} `tfsdk:"subject" json:"subject,omitempty"`
		} `tfsdk:"api_passthrough" json:"apiPassthrough,omitempty"`
		CertificateAuthorityARN *string `tfsdk:"certificate_authority_arn" json:"certificateAuthorityARN,omitempty"`
		CertificateAuthorityRef *struct {
			From *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"certificate_authority_ref" json:"certificateAuthorityRef,omitempty"`
		CertificateOutput *struct {
			Key       *string `tfsdk:"key" json:"key,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"certificate_output" json:"certificateOutput,omitempty"`
		CertificateSigningRequest    *string `tfsdk:"certificate_signing_request" json:"certificateSigningRequest,omitempty"`
		CertificateSigningRequestRef *struct {
			From *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"certificate_signing_request_ref" json:"certificateSigningRequestRef,omitempty"`
		SigningAlgorithm *string `tfsdk:"signing_algorithm" json:"signingAlgorithm,omitempty"`
		TemplateARN      *string `tfsdk:"template_arn" json:"templateARN,omitempty"`
		Validity         *struct {
			Type  *string `tfsdk:"type" json:"type,omitempty"`
			Value *int64  `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"validity" json:"validity,omitempty"`
		ValidityNotBefore *struct {
			Type  *string `tfsdk:"type" json:"type,omitempty"`
			Value *int64  `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"validity_not_before" json:"validityNotBefore,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AcmpcaServicesK8SAwsCertificateV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_acmpca_services_k8s_aws_certificate_v1alpha1_manifest"
}

func (r *AcmpcaServicesK8SAwsCertificateV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Certificate is the Schema for the Certificates API",
		MarkdownDescription: "Certificate is the Schema for the Certificates API",
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
				Description:         "CertificateSpec defines the desired state of Certificate.",
				MarkdownDescription: "CertificateSpec defines the desired state of Certificate.",
				Attributes: map[string]schema.Attribute{
					"api_passthrough": schema.SingleNestedAttribute{
						Description:         "Specifies X.509 certificate information to be included in the issued certificate. An APIPassthrough or APICSRPassthrough template variant must be selected, or else this parameter is ignored. For more information about using these templates, see Understanding Certificate Templates (https://docs.aws.amazon.com/privateca/latest/userguide/UsingTemplates.html). If conflicting or duplicate certificate information is supplied during certificate issuance, Amazon Web Services Private CA applies order of operation rules (https://docs.aws.amazon.com/privateca/latest/userguide/UsingTemplates.html#template-order-of-operations) to determine what information is used.",
						MarkdownDescription: "Specifies X.509 certificate information to be included in the issued certificate. An APIPassthrough or APICSRPassthrough template variant must be selected, or else this parameter is ignored. For more information about using these templates, see Understanding Certificate Templates (https://docs.aws.amazon.com/privateca/latest/userguide/UsingTemplates.html). If conflicting or duplicate certificate information is supplied during certificate issuance, Amazon Web Services Private CA applies order of operation rules (https://docs.aws.amazon.com/privateca/latest/userguide/UsingTemplates.html#template-order-of-operations) to determine what information is used.",
						Attributes: map[string]schema.Attribute{
							"extensions": schema.SingleNestedAttribute{
								Description:         "Contains X.509 extension information for a certificate.",
								MarkdownDescription: "Contains X.509 extension information for a certificate.",
								Attributes: map[string]schema.Attribute{
									"certificate_policies": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"cert_policy_id": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"policy_qualifiers": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"policy_qualifier_id": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"qualifier": schema.SingleNestedAttribute{
																Description:         "Defines a PolicyInformation qualifier. Amazon Web Services Private CA supports the certification practice statement (CPS) qualifier (https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.4) defined in RFC 5280.",
																MarkdownDescription: "Defines a PolicyInformation qualifier. Amazon Web Services Private CA supports the certification practice statement (CPS) qualifier (https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.4) defined in RFC 5280.",
																Attributes: map[string]schema.Attribute{
																	"cps_uri": schema.StringAttribute{
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

									"custom_extensions": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"critical": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"object_identifier": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
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

									"extended_key_usage": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"extended_key_usage_object_identifier": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"extended_key_usage_type": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
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

									"key_usage": schema.SingleNestedAttribute{
										Description:         "Defines one or more purposes for which the key contained in the certificate can be used. Default value for each option is false.",
										MarkdownDescription: "Defines one or more purposes for which the key contained in the certificate can be used. Default value for each option is false.",
										Attributes: map[string]schema.Attribute{
											"crl_sign": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"data_encipherment": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"decipher_only": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"digital_signature": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"encipher_only": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"key_agreement": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"key_cert_sign": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"key_encipherment": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"non_repudiation": schema.BoolAttribute{
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

									"subject_alternative_names": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"directory_name": schema.SingleNestedAttribute{
													Description:         "Contains information about the certificate subject. The Subject field in the certificate identifies the entity that owns or controls the public key in the certificate. The entity can be a user, computer, device, or service. The Subject must contain an X.500 distinguished name (DN). A DN is a sequence of relative distinguished names (RDNs). The RDNs are separated by commas in the certificate.",
													MarkdownDescription: "Contains information about the certificate subject. The Subject field in the certificate identifies the entity that owns or controls the public key in the certificate. The entity can be a user, computer, device, or service. The Subject must contain an X.500 distinguished name (DN). A DN is a sequence of relative distinguished names (RDNs). The RDNs are separated by commas in the certificate.",
													Attributes: map[string]schema.Attribute{
														"common_name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"country": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"custom_attributes": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"object_identifier": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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

														"distinguished_name_qualifier": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"generation_qualifier": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"given_name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"initials": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"locality": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"organization": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"organizational_unit": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"pseudonym": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"serial_number": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"state": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"surname": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"title": schema.StringAttribute{
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

												"dns_name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"edi_party_name": schema.SingleNestedAttribute{
													Description:         "Describes an Electronic Data Interchange (EDI) entity as described in as defined in Subject Alternative Name (https://datatracker.ietf.org/doc/html/rfc5280) in RFC 5280.",
													MarkdownDescription: "Describes an Electronic Data Interchange (EDI) entity as described in as defined in Subject Alternative Name (https://datatracker.ietf.org/doc/html/rfc5280) in RFC 5280.",
													Attributes: map[string]schema.Attribute{
														"name_assigner": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"party_name": schema.StringAttribute{
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

												"ip_address": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"other_name": schema.SingleNestedAttribute{
													Description:         "Defines a custom ASN.1 X.400 GeneralName using an object identifier (OID) and value. The OID must satisfy the regular expression shown below. For more information, see NIST's definition of Object Identifier (OID) (https://csrc.nist.gov/glossary/term/Object_Identifier).",
													MarkdownDescription: "Defines a custom ASN.1 X.400 GeneralName using an object identifier (OID) and value. The OID must satisfy the regular expression shown below. For more information, see NIST's definition of Object Identifier (OID) (https://csrc.nist.gov/glossary/term/Object_Identifier).",
													Attributes: map[string]schema.Attribute{
														"type_id": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"value": schema.StringAttribute{
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

												"registered_id": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"rfc822_name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uniform_resource_identifier": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
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

							"subject": schema.SingleNestedAttribute{
								Description:         "Contains information about the certificate subject. The Subject field in the certificate identifies the entity that owns or controls the public key in the certificate. The entity can be a user, computer, device, or service. The Subject must contain an X.500 distinguished name (DN). A DN is a sequence of relative distinguished names (RDNs). The RDNs are separated by commas in the certificate.",
								MarkdownDescription: "Contains information about the certificate subject. The Subject field in the certificate identifies the entity that owns or controls the public key in the certificate. The entity can be a user, computer, device, or service. The Subject must contain an X.500 distinguished name (DN). A DN is a sequence of relative distinguished names (RDNs). The RDNs are separated by commas in the certificate.",
								Attributes: map[string]schema.Attribute{
									"common_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"country": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"custom_attributes": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"object_identifier": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
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

									"distinguished_name_qualifier": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"generation_qualifier": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"given_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"initials": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"locality": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"organization": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"organizational_unit": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"pseudonym": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"serial_number": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"state": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"surname": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"title": schema.StringAttribute{
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"certificate_authority_arn": schema.StringAttribute{
						Description:         "The Amazon Resource Name (ARN) that was returned when you called CreateCertificateAuthority (https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthority.html). This must be of the form: arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012",
						MarkdownDescription: "The Amazon Resource Name (ARN) that was returned when you called CreateCertificateAuthority (https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthority.html). This must be of the form: arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"certificate_authority_ref": schema.SingleNestedAttribute{
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef: from: name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef: from: name: my-api",
						Attributes: map[string]schema.Attribute{
							"from": schema.SingleNestedAttribute{
								Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
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

					"certificate_output": schema.SingleNestedAttribute{
						Description:         "SecretKeyReference combines a k8s corev1.SecretReference with a specific key within the referred-to Secret",
						MarkdownDescription: "SecretKeyReference combines a k8s corev1.SecretReference with a specific key within the referred-to Secret",
						Attributes: map[string]schema.Attribute{
							"key": schema.StringAttribute{
								Description:         "Key is the key within the secret",
								MarkdownDescription: "Key is the key within the secret",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "name is unique within a namespace to reference a secret resource.",
								MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "namespace defines the space within which the secret name must be unique.",
								MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"certificate_signing_request": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"certificate_signing_request_ref": schema.SingleNestedAttribute{
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef: from: name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef: from: name: my-api",
						Attributes: map[string]schema.Attribute{
							"from": schema.SingleNestedAttribute{
								Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
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

					"signing_algorithm": schema.StringAttribute{
						Description:         "The name of the algorithm that will be used to sign the certificate to be issued. This parameter should not be confused with the SigningAlgorithm parameter used to sign a CSR in the CreateCertificateAuthority action. The specified signing algorithm family (RSA or ECDSA) must match the algorithm family of the CA's secret key.",
						MarkdownDescription: "The name of the algorithm that will be used to sign the certificate to be issued. This parameter should not be confused with the SigningAlgorithm parameter used to sign a CSR in the CreateCertificateAuthority action. The specified signing algorithm family (RSA or ECDSA) must match the algorithm family of the CA's secret key.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"template_arn": schema.StringAttribute{
						Description:         "Specifies a custom configuration template to use when issuing a certificate. If this parameter is not provided, Amazon Web Services Private CA defaults to the EndEntityCertificate/V1 template. For CA certificates, you should choose the shortest path length that meets your needs. The path length is indicated by the PathLenN portion of the ARN, where N is the CA depth (https://docs.aws.amazon.com/privateca/latest/userguide/PcaTerms.html#terms-cadepth). Note: The CA depth configured on a subordinate CA certificate must not exceed the limit set by its parents in the CA hierarchy. For a list of TemplateArn values supported by Amazon Web Services Private CA, see Understanding Certificate Templates (https://docs.aws.amazon.com/privateca/latest/userguide/UsingTemplates.html).",
						MarkdownDescription: "Specifies a custom configuration template to use when issuing a certificate. If this parameter is not provided, Amazon Web Services Private CA defaults to the EndEntityCertificate/V1 template. For CA certificates, you should choose the shortest path length that meets your needs. The path length is indicated by the PathLenN portion of the ARN, where N is the CA depth (https://docs.aws.amazon.com/privateca/latest/userguide/PcaTerms.html#terms-cadepth). Note: The CA depth configured on a subordinate CA certificate must not exceed the limit set by its parents in the CA hierarchy. For a list of TemplateArn values supported by Amazon Web Services Private CA, see Understanding Certificate Templates (https://docs.aws.amazon.com/privateca/latest/userguide/UsingTemplates.html).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"validity": schema.SingleNestedAttribute{
						Description:         "Information describing the end of the validity period of the certificate. This parameter sets the “Not After” date for the certificate. Certificate validity is the period of time during which a certificate is valid. Validity can be expressed as an explicit date and time when the certificate expires, or as a span of time after issuance, stated in days, months, or years. For more information, see Validity (https://datatracker.ietf.org/doc/html/rfc5280#section-4.1.2.5) in RFC 5280. This value is unaffected when ValidityNotBefore is also specified. For example, if Validity is set to 20 days in the future, the certificate will expire 20 days from issuance time regardless of the ValidityNotBefore value. The end of the validity period configured on a certificate must not exceed the limit set on its parents in the CA hierarchy.",
						MarkdownDescription: "Information describing the end of the validity period of the certificate. This parameter sets the “Not After” date for the certificate. Certificate validity is the period of time during which a certificate is valid. Validity can be expressed as an explicit date and time when the certificate expires, or as a span of time after issuance, stated in days, months, or years. For more information, see Validity (https://datatracker.ietf.org/doc/html/rfc5280#section-4.1.2.5) in RFC 5280. This value is unaffected when ValidityNotBefore is also specified. For example, if Validity is set to 20 days in the future, the certificate will expire 20 days from issuance time regardless of the ValidityNotBefore value. The end of the validity period configured on a certificate must not exceed the limit set on its parents in the CA hierarchy.",
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"value": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"validity_not_before": schema.SingleNestedAttribute{
						Description:         "Information describing the start of the validity period of the certificate. This parameter sets the “Not Before' date for the certificate. By default, when issuing a certificate, Amazon Web Services Private CA sets the 'Not Before' date to the issuance time minus 60 minutes. This compensates for clock inconsistencies across computer systems. The ValidityNotBefore parameter can be used to customize the “Not Before” value. Unlike the Validity parameter, the ValidityNotBefore parameter is optional. The ValidityNotBefore value is expressed as an explicit date and time, using the Validity type value ABSOLUTE. For more information, see Validity (https://docs.aws.amazon.com/privateca/latest/APIReference/API_Validity.html) in this API reference and Validity (https://datatracker.ietf.org/doc/html/rfc5280#section-4.1.2.5) in RFC 5280.",
						MarkdownDescription: "Information describing the start of the validity period of the certificate. This parameter sets the “Not Before' date for the certificate. By default, when issuing a certificate, Amazon Web Services Private CA sets the 'Not Before' date to the issuance time minus 60 minutes. This compensates for clock inconsistencies across computer systems. The ValidityNotBefore parameter can be used to customize the “Not Before” value. Unlike the Validity parameter, the ValidityNotBefore parameter is optional. The ValidityNotBefore value is expressed as an explicit date and time, using the Validity type value ABSOLUTE. For more information, see Validity (https://docs.aws.amazon.com/privateca/latest/APIReference/API_Validity.html) in this API reference and Validity (https://datatracker.ietf.org/doc/html/rfc5280#section-4.1.2.5) in RFC 5280.",
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"value": schema.Int64Attribute{
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *AcmpcaServicesK8SAwsCertificateV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_acmpca_services_k8s_aws_certificate_v1alpha1_manifest")

	var model AcmpcaServicesK8SAwsCertificateV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("acmpca.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("Certificate")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
