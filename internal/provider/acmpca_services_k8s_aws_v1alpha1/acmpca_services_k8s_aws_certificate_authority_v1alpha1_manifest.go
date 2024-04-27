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
	_ datasource.DataSource = &AcmpcaServicesK8SAwsCertificateAuthorityV1Alpha1Manifest{}
)

func NewAcmpcaServicesK8SAwsCertificateAuthorityV1Alpha1Manifest() datasource.DataSource {
	return &AcmpcaServicesK8SAwsCertificateAuthorityV1Alpha1Manifest{}
}

type AcmpcaServicesK8SAwsCertificateAuthorityV1Alpha1Manifest struct{}

type AcmpcaServicesK8SAwsCertificateAuthorityV1Alpha1ManifestData struct {
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
		CertificateAuthorityConfiguration *struct {
			CsrExtensions *struct {
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
				SubjectInformationAccess *[]struct {
					AccessLocation *struct {
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
					} `tfsdk:"access_location" json:"accessLocation,omitempty"`
					AccessMethod *struct {
						AccessMethodType       *string `tfsdk:"access_method_type" json:"accessMethodType,omitempty"`
						CustomObjectIdentifier *string `tfsdk:"custom_object_identifier" json:"customObjectIdentifier,omitempty"`
					} `tfsdk:"access_method" json:"accessMethod,omitempty"`
				} `tfsdk:"subject_information_access" json:"subjectInformationAccess,omitempty"`
			} `tfsdk:"csr_extensions" json:"csrExtensions,omitempty"`
			KeyAlgorithm     *string `tfsdk:"key_algorithm" json:"keyAlgorithm,omitempty"`
			SigningAlgorithm *string `tfsdk:"signing_algorithm" json:"signingAlgorithm,omitempty"`
			Subject          *struct {
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
		} `tfsdk:"certificate_authority_configuration" json:"certificateAuthorityConfiguration,omitempty"`
		KeyStorageSecurityStandard *string `tfsdk:"key_storage_security_standard" json:"keyStorageSecurityStandard,omitempty"`
		RevocationConfiguration    *struct {
			CrlConfiguration *struct {
				CustomCNAME      *string `tfsdk:"custom_cname" json:"customCNAME,omitempty"`
				Enabled          *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
				ExpirationInDays *int64  `tfsdk:"expiration_in_days" json:"expirationInDays,omitempty"`
				S3BucketName     *string `tfsdk:"s3_bucket_name" json:"s3BucketName,omitempty"`
				S3ObjectACL      *string `tfsdk:"s3_object_acl" json:"s3ObjectACL,omitempty"`
			} `tfsdk:"crl_configuration" json:"crlConfiguration,omitempty"`
			OcspConfiguration *struct {
				Enabled         *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
				OcspCustomCNAME *string `tfsdk:"ocsp_custom_cname" json:"ocspCustomCNAME,omitempty"`
			} `tfsdk:"ocsp_configuration" json:"ocspConfiguration,omitempty"`
		} `tfsdk:"revocation_configuration" json:"revocationConfiguration,omitempty"`
		Tags *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
		Type      *string `tfsdk:"type" json:"type,omitempty"`
		UsageMode *string `tfsdk:"usage_mode" json:"usageMode,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AcmpcaServicesK8SAwsCertificateAuthorityV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_acmpca_services_k8s_aws_certificate_authority_v1alpha1_manifest"
}

func (r *AcmpcaServicesK8SAwsCertificateAuthorityV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CertificateAuthority is the Schema for the CertificateAuthorities API",
		MarkdownDescription: "CertificateAuthority is the Schema for the CertificateAuthorities API",
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
				Description:         "CertificateAuthoritySpec defines the desired state of CertificateAuthority.Contains information about your private certificate authority (CA). Yourprivate CA can issue and revoke X.509 digital certificates. Digital certificatesverify that the entity named in the certificate Subject field owns or controlsthe public key contained in the Subject Public Key Info field. Call the CreateCertificateAuthority(https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthority.html)action to create your private CA. You must then call the GetCertificateAuthorityCertificate(https://docs.aws.amazon.com/privateca/latest/APIReference/API_GetCertificateAuthorityCertificate.html)action to retrieve a private CA certificate signing request (CSR). Sign theCSR with your Amazon Web Services Private CA-hosted or on-premises root orsubordinate CA certificate. Call the ImportCertificateAuthorityCertificate(https://docs.aws.amazon.com/privateca/latest/APIReference/API_ImportCertificateAuthorityCertificate.html)action to import the signed certificate into Certificate Manager (ACM).",
				MarkdownDescription: "CertificateAuthoritySpec defines the desired state of CertificateAuthority.Contains information about your private certificate authority (CA). Yourprivate CA can issue and revoke X.509 digital certificates. Digital certificatesverify that the entity named in the certificate Subject field owns or controlsthe public key contained in the Subject Public Key Info field. Call the CreateCertificateAuthority(https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthority.html)action to create your private CA. You must then call the GetCertificateAuthorityCertificate(https://docs.aws.amazon.com/privateca/latest/APIReference/API_GetCertificateAuthorityCertificate.html)action to retrieve a private CA certificate signing request (CSR). Sign theCSR with your Amazon Web Services Private CA-hosted or on-premises root orsubordinate CA certificate. Call the ImportCertificateAuthorityCertificate(https://docs.aws.amazon.com/privateca/latest/APIReference/API_ImportCertificateAuthorityCertificate.html)action to import the signed certificate into Certificate Manager (ACM).",
				Attributes: map[string]schema.Attribute{
					"certificate_authority_configuration": schema.SingleNestedAttribute{
						Description:         "Name and bit size of the private key algorithm, the name of the signing algorithm,and X.500 certificate subject information.",
						MarkdownDescription: "Name and bit size of the private key algorithm, the name of the signing algorithm,and X.500 certificate subject information.",
						Attributes: map[string]schema.Attribute{
							"csr_extensions": schema.SingleNestedAttribute{
								Description:         "Describes the certificate extensions to be added to the certificate signingrequest (CSR).",
								MarkdownDescription: "Describes the certificate extensions to be added to the certificate signingrequest (CSR).",
								Attributes: map[string]schema.Attribute{
									"key_usage": schema.SingleNestedAttribute{
										Description:         "Defines one or more purposes for which the key contained in the certificatecan be used. Default value for each option is false.",
										MarkdownDescription: "Defines one or more purposes for which the key contained in the certificatecan be used. Default value for each option is false.",
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

									"subject_information_access": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"access_location": schema.SingleNestedAttribute{
													Description:         "Describes an ASN.1 X.400 GeneralName as defined in RFC 5280 (https://datatracker.ietf.org/doc/html/rfc5280).Only one of the following naming options should be provided. Providing morethan one option results in an InvalidArgsException error.",
													MarkdownDescription: "Describes an ASN.1 X.400 GeneralName as defined in RFC 5280 (https://datatracker.ietf.org/doc/html/rfc5280).Only one of the following naming options should be provided. Providing morethan one option results in an InvalidArgsException error.",
													Attributes: map[string]schema.Attribute{
														"directory_name": schema.SingleNestedAttribute{
															Description:         "Contains information about the certificate subject. The Subject field inthe certificate identifies the entity that owns or controls the public keyin the certificate. The entity can be a user, computer, device, or service.The Subject must contain an X.500 distinguished name (DN). A DN is a sequenceof relative distinguished names (RDNs). The RDNs are separated by commasin the certificate.",
															MarkdownDescription: "Contains information about the certificate subject. The Subject field inthe certificate identifies the entity that owns or controls the public keyin the certificate. The entity can be a user, computer, device, or service.The Subject must contain an X.500 distinguished name (DN). A DN is a sequenceof relative distinguished names (RDNs). The RDNs are separated by commasin the certificate.",
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
															Description:         "Describes an Electronic Data Interchange (EDI) entity as described in asdefined in Subject Alternative Name (https://datatracker.ietf.org/doc/html/rfc5280)in RFC 5280.",
															MarkdownDescription: "Describes an Electronic Data Interchange (EDI) entity as described in asdefined in Subject Alternative Name (https://datatracker.ietf.org/doc/html/rfc5280)in RFC 5280.",
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
															Description:         "Defines a custom ASN.1 X.400 GeneralName using an object identifier (OID)and value. The OID must satisfy the regular expression shown below. For moreinformation, see NIST's definition of Object Identifier (OID) (https://csrc.nist.gov/glossary/term/Object_Identifier).",
															MarkdownDescription: "Defines a custom ASN.1 X.400 GeneralName using an object identifier (OID)and value. The OID must satisfy the regular expression shown below. For moreinformation, see NIST's definition of Object Identifier (OID) (https://csrc.nist.gov/glossary/term/Object_Identifier).",
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
													Required: false,
													Optional: true,
													Computed: false,
												},

												"access_method": schema.SingleNestedAttribute{
													Description:         "Describes the type and format of extension access. Only one of CustomObjectIdentifieror AccessMethodType may be provided. Providing both results in InvalidArgsException.",
													MarkdownDescription: "Describes the type and format of extension access. Only one of CustomObjectIdentifieror AccessMethodType may be provided. Providing both results in InvalidArgsException.",
													Attributes: map[string]schema.Attribute{
														"access_method_type": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"custom_object_identifier": schema.StringAttribute{
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
								Required: false,
								Optional: true,
								Computed: false,
							},

							"key_algorithm": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"signing_algorithm": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"subject": schema.SingleNestedAttribute{
								Description:         "Contains information about the certificate subject. The Subject field inthe certificate identifies the entity that owns or controls the public keyin the certificate. The entity can be a user, computer, device, or service.The Subject must contain an X.500 distinguished name (DN). A DN is a sequenceof relative distinguished names (RDNs). The RDNs are separated by commasin the certificate.",
								MarkdownDescription: "Contains information about the certificate subject. The Subject field inthe certificate identifies the entity that owns or controls the public keyin the certificate. The entity can be a user, computer, device, or service.The Subject must contain an X.500 distinguished name (DN). A DN is a sequenceof relative distinguished names (RDNs). The RDNs are separated by commasin the certificate.",
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
						Required: true,
						Optional: false,
						Computed: false,
					},

					"key_storage_security_standard": schema.StringAttribute{
						Description:         "Specifies a cryptographic key management compliance standard used for handlingCA keys.Default: FIPS_140_2_LEVEL_3_OR_HIGHERSome Amazon Web Services Regions do not support the default. When creatinga CA in these Regions, you must provide FIPS_140_2_LEVEL_2_OR_HIGHER as theargument for KeyStorageSecurityStandard. Failure to do this results in anInvalidArgsException with the message, 'A certificate authority cannot becreated in this region with the specified security standard.'For information about security standard support in various Regions, see Storageand security compliance of Amazon Web Services Private CA private keys (https://docs.aws.amazon.com/privateca/latest/userguide/data-protection.html#private-keys).",
						MarkdownDescription: "Specifies a cryptographic key management compliance standard used for handlingCA keys.Default: FIPS_140_2_LEVEL_3_OR_HIGHERSome Amazon Web Services Regions do not support the default. When creatinga CA in these Regions, you must provide FIPS_140_2_LEVEL_2_OR_HIGHER as theargument for KeyStorageSecurityStandard. Failure to do this results in anInvalidArgsException with the message, 'A certificate authority cannot becreated in this region with the specified security standard.'For information about security standard support in various Regions, see Storageand security compliance of Amazon Web Services Private CA private keys (https://docs.aws.amazon.com/privateca/latest/userguide/data-protection.html#private-keys).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"revocation_configuration": schema.SingleNestedAttribute{
						Description:         "Contains information to enable Online Certificate Status Protocol (OCSP)support, to enable a certificate revocation list (CRL), to enable both, orto enable neither. The default is for both certificate validation mechanismsto be disabled.The following requirements apply to revocation configurations.   * A configuration disabling CRLs or OCSP must contain only the Enabled=False   parameter, and will fail if other parameters such as CustomCname or ExpirationInDays   are included.   * In a CRL configuration, the S3BucketName parameter must conform to Amazon   S3 bucket naming rules (https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucketnamingrules.html).   * A configuration containing a custom Canonical Name (CNAME) parameter   for CRLs or OCSP must conform to RFC2396 (https://www.ietf.org/rfc/rfc2396.txt)   restrictions on the use of special characters in a CNAME.   * In a CRL or OCSP configuration, the value of a CNAME parameter must   not include a protocol prefix such as 'http://' or 'https://'.For more information, see the OcspConfiguration (https://docs.aws.amazon.com/privateca/latest/APIReference/API_OcspConfiguration.html)and CrlConfiguration (https://docs.aws.amazon.com/privateca/latest/APIReference/API_CrlConfiguration.html)types.",
						MarkdownDescription: "Contains information to enable Online Certificate Status Protocol (OCSP)support, to enable a certificate revocation list (CRL), to enable both, orto enable neither. The default is for both certificate validation mechanismsto be disabled.The following requirements apply to revocation configurations.   * A configuration disabling CRLs or OCSP must contain only the Enabled=False   parameter, and will fail if other parameters such as CustomCname or ExpirationInDays   are included.   * In a CRL configuration, the S3BucketName parameter must conform to Amazon   S3 bucket naming rules (https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucketnamingrules.html).   * A configuration containing a custom Canonical Name (CNAME) parameter   for CRLs or OCSP must conform to RFC2396 (https://www.ietf.org/rfc/rfc2396.txt)   restrictions on the use of special characters in a CNAME.   * In a CRL or OCSP configuration, the value of a CNAME parameter must   not include a protocol prefix such as 'http://' or 'https://'.For more information, see the OcspConfiguration (https://docs.aws.amazon.com/privateca/latest/APIReference/API_OcspConfiguration.html)and CrlConfiguration (https://docs.aws.amazon.com/privateca/latest/APIReference/API_CrlConfiguration.html)types.",
						Attributes: map[string]schema.Attribute{
							"crl_configuration": schema.SingleNestedAttribute{
								Description:         "Contains configuration information for a certificate revocation list (CRL).Your private certificate authority (CA) creates base CRLs. Delta CRLs arenot supported. You can enable CRLs for your new or an existing private CAby setting the Enabled parameter to true. Your private CA writes CRLs toan S3 bucket that you specify in the S3BucketName parameter. You can hidethe name of your bucket by specifying a value for the CustomCname parameter.Your private CA copies the CNAME or the S3 bucket name to the CRL DistributionPoints extension of each certificate it issues. Your S3 bucket policy mustgive write permission to Amazon Web Services Private CA.Amazon Web Services Private CA assets that are stored in Amazon S3 can beprotected with encryption. For more information, see Encrypting Your CRLs(https://docs.aws.amazon.com/privateca/latest/userguide/PcaCreateCa.html#crl-encryption).Your private CA uses the value in the ExpirationInDays parameter to calculatethe nextUpdate field in the CRL. The CRL is refreshed prior to a certificate'sexpiration date or when a certificate is revoked. When a certificate is revoked,it appears in the CRL until the certificate expires, and then in one additionalCRL after expiration, and it always appears in the audit report.A CRL is typically updated approximately 30 minutes after a certificate isrevoked. If for any reason a CRL update fails, Amazon Web Services PrivateCA makes further attempts every 15 minutes.CRLs contain the following fields:   * Version: The current version number defined in RFC 5280 is V2. The integer   value is 0x1.   * Signature Algorithm: The name of the algorithm used to sign the CRL.   * Issuer: The X.500 distinguished name of your private CA that issued   the CRL.   * Last Update: The issue date and time of this CRL.   * Next Update: The day and time by which the next CRL will be issued.   * Revoked Certificates: List of revoked certificates. Each list item contains   the following information. Serial Number: The serial number, in hexadecimal   format, of the revoked certificate. Revocation Date: Date and time the   certificate was revoked. CRL Entry Extensions: Optional extensions for   the CRL entry. X509v3 CRL Reason Code: Reason the certificate was revoked.   * CRL Extensions: Optional extensions for the CRL. X509v3 Authority Key   Identifier: Identifies the public key associated with the private key   used to sign the certificate. X509v3 CRL Number:: Decimal sequence number   for the CRL.   * Signature Algorithm: Algorithm used by your private CA to sign the CRL.   * Signature Value: Signature computed over the CRL.Certificate revocation lists created by Amazon Web Services Private CA areDER-encoded. You can use the following OpenSSL command to list a CRL.openssl crl -inform DER -text -in crl_path -nooutFor more information, see Planning a certificate revocation list (CRL) (https://docs.aws.amazon.com/privateca/latest/userguide/crl-planning.html)in the Amazon Web Services Private Certificate Authority User Guide",
								MarkdownDescription: "Contains configuration information for a certificate revocation list (CRL).Your private certificate authority (CA) creates base CRLs. Delta CRLs arenot supported. You can enable CRLs for your new or an existing private CAby setting the Enabled parameter to true. Your private CA writes CRLs toan S3 bucket that you specify in the S3BucketName parameter. You can hidethe name of your bucket by specifying a value for the CustomCname parameter.Your private CA copies the CNAME or the S3 bucket name to the CRL DistributionPoints extension of each certificate it issues. Your S3 bucket policy mustgive write permission to Amazon Web Services Private CA.Amazon Web Services Private CA assets that are stored in Amazon S3 can beprotected with encryption. For more information, see Encrypting Your CRLs(https://docs.aws.amazon.com/privateca/latest/userguide/PcaCreateCa.html#crl-encryption).Your private CA uses the value in the ExpirationInDays parameter to calculatethe nextUpdate field in the CRL. The CRL is refreshed prior to a certificate'sexpiration date or when a certificate is revoked. When a certificate is revoked,it appears in the CRL until the certificate expires, and then in one additionalCRL after expiration, and it always appears in the audit report.A CRL is typically updated approximately 30 minutes after a certificate isrevoked. If for any reason a CRL update fails, Amazon Web Services PrivateCA makes further attempts every 15 minutes.CRLs contain the following fields:   * Version: The current version number defined in RFC 5280 is V2. The integer   value is 0x1.   * Signature Algorithm: The name of the algorithm used to sign the CRL.   * Issuer: The X.500 distinguished name of your private CA that issued   the CRL.   * Last Update: The issue date and time of this CRL.   * Next Update: The day and time by which the next CRL will be issued.   * Revoked Certificates: List of revoked certificates. Each list item contains   the following information. Serial Number: The serial number, in hexadecimal   format, of the revoked certificate. Revocation Date: Date and time the   certificate was revoked. CRL Entry Extensions: Optional extensions for   the CRL entry. X509v3 CRL Reason Code: Reason the certificate was revoked.   * CRL Extensions: Optional extensions for the CRL. X509v3 Authority Key   Identifier: Identifies the public key associated with the private key   used to sign the certificate. X509v3 CRL Number:: Decimal sequence number   for the CRL.   * Signature Algorithm: Algorithm used by your private CA to sign the CRL.   * Signature Value: Signature computed over the CRL.Certificate revocation lists created by Amazon Web Services Private CA areDER-encoded. You can use the following OpenSSL command to list a CRL.openssl crl -inform DER -text -in crl_path -nooutFor more information, see Planning a certificate revocation list (CRL) (https://docs.aws.amazon.com/privateca/latest/userguide/crl-planning.html)in the Amazon Web Services Private Certificate Authority User Guide",
								Attributes: map[string]schema.Attribute{
									"custom_cname": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"enabled": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"expiration_in_days": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"s3_bucket_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"s3_object_acl": schema.StringAttribute{
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

							"ocsp_configuration": schema.SingleNestedAttribute{
								Description:         "Contains information to enable and configure Online Certificate Status Protocol(OCSP) for validating certificate revocation status.When you revoke a certificate, OCSP responses may take up to 60 minutes toreflect the new status.",
								MarkdownDescription: "Contains information to enable and configure Online Certificate Status Protocol(OCSP) for validating certificate revocation status.When you revoke a certificate, OCSP responses may take up to 60 minutes toreflect the new status.",
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ocsp_custom_cname": schema.StringAttribute{
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

					"tags": schema.ListNestedAttribute{
						Description:         "Key-value pairs that will be attached to the new private CA. You can associateup to 50 tags with a private CA. For information using tags with IAM to managepermissions, see Controlling Access Using IAM Tags (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_iam-tags.html).",
						MarkdownDescription: "Key-value pairs that will be attached to the new private CA. You can associateup to 50 tags with a private CA. For information using tags with IAM to managepermissions, see Controlling Access Using IAM Tags (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_iam-tags.html).",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"key": schema.StringAttribute{
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

					"type": schema.StringAttribute{
						Description:         "The type of the certificate authority.",
						MarkdownDescription: "The type of the certificate authority.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"usage_mode": schema.StringAttribute{
						Description:         "Specifies whether the CA issues general-purpose certificates that typicallyrequire a revocation mechanism, or short-lived certificates that may optionallyomit revocation because they expire quickly. Short-lived certificate validityis limited to seven days.The default value is GENERAL_PURPOSE.",
						MarkdownDescription: "Specifies whether the CA issues general-purpose certificates that typicallyrequire a revocation mechanism, or short-lived certificates that may optionallyomit revocation because they expire quickly. Short-lived certificate validityis limited to seven days.The default value is GENERAL_PURPOSE.",
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

func (r *AcmpcaServicesK8SAwsCertificateAuthorityV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_acmpca_services_k8s_aws_certificate_authority_v1alpha1_manifest")

	var model AcmpcaServicesK8SAwsCertificateAuthorityV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("acmpca.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("CertificateAuthority")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
