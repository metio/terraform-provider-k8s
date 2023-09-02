/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package cert_manager_io_v1

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
	_ datasource.DataSource              = &CertManagerIoCertificateV1DataSource{}
	_ datasource.DataSourceWithConfigure = &CertManagerIoCertificateV1DataSource{}
)

func NewCertManagerIoCertificateV1DataSource() datasource.DataSource {
	return &CertManagerIoCertificateV1DataSource{}
}

type CertManagerIoCertificateV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type CertManagerIoCertificateV1DataSourceData struct {
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
		AdditionalOutputFormats *[]struct {
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"additional_output_formats" json:"additionalOutputFormats,omitempty"`
		CommonName            *string   `tfsdk:"common_name" json:"commonName,omitempty"`
		DnsNames              *[]string `tfsdk:"dns_names" json:"dnsNames,omitempty"`
		Duration              *string   `tfsdk:"duration" json:"duration,omitempty"`
		EmailAddresses        *[]string `tfsdk:"email_addresses" json:"emailAddresses,omitempty"`
		EncodeUsagesInRequest *bool     `tfsdk:"encode_usages_in_request" json:"encodeUsagesInRequest,omitempty"`
		IpAddresses           *[]string `tfsdk:"ip_addresses" json:"ipAddresses,omitempty"`
		IsCA                  *bool     `tfsdk:"is_ca" json:"isCA,omitempty"`
		IssuerRef             *struct {
			Group *string `tfsdk:"group" json:"group,omitempty"`
			Kind  *string `tfsdk:"kind" json:"kind,omitempty"`
			Name  *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"issuer_ref" json:"issuerRef,omitempty"`
		Keystores *struct {
			Jks *struct {
				Create            *bool `tfsdk:"create" json:"create,omitempty"`
				PasswordSecretRef *struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"password_secret_ref" json:"passwordSecretRef,omitempty"`
			} `tfsdk:"jks" json:"jks,omitempty"`
			Pkcs12 *struct {
				Create            *bool `tfsdk:"create" json:"create,omitempty"`
				PasswordSecretRef *struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"password_secret_ref" json:"passwordSecretRef,omitempty"`
			} `tfsdk:"pkcs12" json:"pkcs12,omitempty"`
		} `tfsdk:"keystores" json:"keystores,omitempty"`
		LiteralSubject *string `tfsdk:"literal_subject" json:"literalSubject,omitempty"`
		PrivateKey     *struct {
			Algorithm      *string `tfsdk:"algorithm" json:"algorithm,omitempty"`
			Encoding       *string `tfsdk:"encoding" json:"encoding,omitempty"`
			RotationPolicy *string `tfsdk:"rotation_policy" json:"rotationPolicy,omitempty"`
			Size           *int64  `tfsdk:"size" json:"size,omitempty"`
		} `tfsdk:"private_key" json:"privateKey,omitempty"`
		RenewBefore          *string `tfsdk:"renew_before" json:"renewBefore,omitempty"`
		RevisionHistoryLimit *int64  `tfsdk:"revision_history_limit" json:"revisionHistoryLimit,omitempty"`
		SecretName           *string `tfsdk:"secret_name" json:"secretName,omitempty"`
		SecretTemplate       *struct {
			Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		} `tfsdk:"secret_template" json:"secretTemplate,omitempty"`
		Subject *struct {
			Countries           *[]string `tfsdk:"countries" json:"countries,omitempty"`
			Localities          *[]string `tfsdk:"localities" json:"localities,omitempty"`
			OrganizationalUnits *[]string `tfsdk:"organizational_units" json:"organizationalUnits,omitempty"`
			Organizations       *[]string `tfsdk:"organizations" json:"organizations,omitempty"`
			PostalCodes         *[]string `tfsdk:"postal_codes" json:"postalCodes,omitempty"`
			Provinces           *[]string `tfsdk:"provinces" json:"provinces,omitempty"`
			SerialNumber        *string   `tfsdk:"serial_number" json:"serialNumber,omitempty"`
			StreetAddresses     *[]string `tfsdk:"street_addresses" json:"streetAddresses,omitempty"`
		} `tfsdk:"subject" json:"subject,omitempty"`
		Uris   *[]string `tfsdk:"uris" json:"uris,omitempty"`
		Usages *[]string `tfsdk:"usages" json:"usages,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CertManagerIoCertificateV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_cert_manager_io_certificate_v1"
}

func (r *CertManagerIoCertificateV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "A Certificate resource should be created to ensure an up to date and signed x509 certificate is stored in the Kubernetes Secret resource named in 'spec.secretName'.  The stored certificate will be renewed before it expires (as configured by 'spec.renewBefore').",
		MarkdownDescription: "A Certificate resource should be created to ensure an up to date and signed x509 certificate is stored in the Kubernetes Secret resource named in 'spec.secretName'.  The stored certificate will be renewed before it expires (as configured by 'spec.renewBefore').",
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
				Description:         "Desired state of the Certificate resource.",
				MarkdownDescription: "Desired state of the Certificate resource.",
				Attributes: map[string]schema.Attribute{
					"additional_output_formats": schema.ListNestedAttribute{
						Description:         "AdditionalOutputFormats defines extra output formats of the private key and signed certificate chain to be written to this Certificate's target Secret. This is an Alpha Feature and is only enabled with the '--feature-gates=AdditionalCertificateOutputFormats=true' option on both the controller and webhook components.",
						MarkdownDescription: "AdditionalOutputFormats defines extra output formats of the private key and signed certificate chain to be written to this Certificate's target Secret. This is an Alpha Feature and is only enabled with the '--feature-gates=AdditionalCertificateOutputFormats=true' option on both the controller and webhook components.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"type": schema.StringAttribute{
									Description:         "Type is the name of the format type that should be written to the Certificate's target Secret.",
									MarkdownDescription: "Type is the name of the format type that should be written to the Certificate's target Secret.",
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

					"common_name": schema.StringAttribute{
						Description:         "CommonName is a common name to be used on the Certificate. The CommonName should have a length of 64 characters or fewer to avoid generating invalid CSRs. This value is ignored by TLS clients when any subject alt name is set. This is x509 behaviour: https://tools.ietf.org/html/rfc6125#section-6.4.4",
						MarkdownDescription: "CommonName is a common name to be used on the Certificate. The CommonName should have a length of 64 characters or fewer to avoid generating invalid CSRs. This value is ignored by TLS clients when any subject alt name is set. This is x509 behaviour: https://tools.ietf.org/html/rfc6125#section-6.4.4",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"dns_names": schema.ListAttribute{
						Description:         "DNSNames is a list of DNS subjectAltNames to be set on the Certificate.",
						MarkdownDescription: "DNSNames is a list of DNS subjectAltNames to be set on the Certificate.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"duration": schema.StringAttribute{
						Description:         "The requested 'duration' (i.e. lifetime) of the Certificate. This option may be ignored/overridden by some issuer types. If unset this defaults to 90 days. Certificate will be renewed either 2/3 through its duration or 'renewBefore' period before its expiry, whichever is later. Minimum accepted duration is 1 hour. Value must be in units accepted by Go time.ParseDuration https://golang.org/pkg/time/#ParseDuration",
						MarkdownDescription: "The requested 'duration' (i.e. lifetime) of the Certificate. This option may be ignored/overridden by some issuer types. If unset this defaults to 90 days. Certificate will be renewed either 2/3 through its duration or 'renewBefore' period before its expiry, whichever is later. Minimum accepted duration is 1 hour. Value must be in units accepted by Go time.ParseDuration https://golang.org/pkg/time/#ParseDuration",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"email_addresses": schema.ListAttribute{
						Description:         "EmailAddresses is a list of email subjectAltNames to be set on the Certificate.",
						MarkdownDescription: "EmailAddresses is a list of email subjectAltNames to be set on the Certificate.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"encode_usages_in_request": schema.BoolAttribute{
						Description:         "EncodeUsagesInRequest controls whether key usages should be present in the CertificateRequest",
						MarkdownDescription: "EncodeUsagesInRequest controls whether key usages should be present in the CertificateRequest",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"ip_addresses": schema.ListAttribute{
						Description:         "IPAddresses is a list of IP address subjectAltNames to be set on the Certificate.",
						MarkdownDescription: "IPAddresses is a list of IP address subjectAltNames to be set on the Certificate.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"is_ca": schema.BoolAttribute{
						Description:         "IsCA will mark this Certificate as valid for certificate signing. This will automatically add the 'cert sign' usage to the list of 'usages'.",
						MarkdownDescription: "IsCA will mark this Certificate as valid for certificate signing. This will automatically add the 'cert sign' usage to the list of 'usages'.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"issuer_ref": schema.SingleNestedAttribute{
						Description:         "IssuerRef is a reference to the issuer for this certificate. If the 'kind' field is not set, or set to 'Issuer', an Issuer resource with the given name in the same namespace as the Certificate will be used. If the 'kind' field is set to 'ClusterIssuer', a ClusterIssuer with the provided name will be used. The 'name' field in this stanza is required at all times.",
						MarkdownDescription: "IssuerRef is a reference to the issuer for this certificate. If the 'kind' field is not set, or set to 'Issuer', an Issuer resource with the given name in the same namespace as the Certificate will be used. If the 'kind' field is set to 'ClusterIssuer', a ClusterIssuer with the provided name will be used. The 'name' field in this stanza is required at all times.",
						Attributes: map[string]schema.Attribute{
							"group": schema.StringAttribute{
								Description:         "Group of the resource being referred to.",
								MarkdownDescription: "Group of the resource being referred to.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind of the resource being referred to.",
								MarkdownDescription: "Kind of the resource being referred to.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the resource being referred to.",
								MarkdownDescription: "Name of the resource being referred to.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"keystores": schema.SingleNestedAttribute{
						Description:         "Keystores configures additional keystore output formats stored in the 'secretName' Secret resource.",
						MarkdownDescription: "Keystores configures additional keystore output formats stored in the 'secretName' Secret resource.",
						Attributes: map[string]schema.Attribute{
							"jks": schema.SingleNestedAttribute{
								Description:         "JKS configures options for storing a JKS keystore in the 'spec.secretName' Secret resource.",
								MarkdownDescription: "JKS configures options for storing a JKS keystore in the 'spec.secretName' Secret resource.",
								Attributes: map[string]schema.Attribute{
									"create": schema.BoolAttribute{
										Description:         "Create enables JKS keystore creation for the Certificate. If true, a file named 'keystore.jks' will be created in the target Secret resource, encrypted using the password stored in 'passwordSecretRef'. The keystore file will be updated immediately. If the issuer provided a CA certificate, a file named 'truststore.jks' will also be created in the target Secret resource, encrypted using the password stored in 'passwordSecretRef' containing the issuing Certificate Authority",
										MarkdownDescription: "Create enables JKS keystore creation for the Certificate. If true, a file named 'keystore.jks' will be created in the target Secret resource, encrypted using the password stored in 'passwordSecretRef'. The keystore file will be updated immediately. If the issuer provided a CA certificate, a file named 'truststore.jks' will also be created in the target Secret resource, encrypted using the password stored in 'passwordSecretRef' containing the issuing Certificate Authority",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"password_secret_ref": schema.SingleNestedAttribute{
										Description:         "PasswordSecretRef is a reference to a key in a Secret resource containing the password used to encrypt the JKS keystore.",
										MarkdownDescription: "PasswordSecretRef is a reference to a key in a Secret resource containing the password used to encrypt the JKS keystore.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
												MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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

							"pkcs12": schema.SingleNestedAttribute{
								Description:         "PKCS12 configures options for storing a PKCS12 keystore in the 'spec.secretName' Secret resource.",
								MarkdownDescription: "PKCS12 configures options for storing a PKCS12 keystore in the 'spec.secretName' Secret resource.",
								Attributes: map[string]schema.Attribute{
									"create": schema.BoolAttribute{
										Description:         "Create enables PKCS12 keystore creation for the Certificate. If true, a file named 'keystore.p12' will be created in the target Secret resource, encrypted using the password stored in 'passwordSecretRef'. The keystore file will be updated immediately. If the issuer provided a CA certificate, a file named 'truststore.p12' will also be created in the target Secret resource, encrypted using the password stored in 'passwordSecretRef' containing the issuing Certificate Authority",
										MarkdownDescription: "Create enables PKCS12 keystore creation for the Certificate. If true, a file named 'keystore.p12' will be created in the target Secret resource, encrypted using the password stored in 'passwordSecretRef'. The keystore file will be updated immediately. If the issuer provided a CA certificate, a file named 'truststore.p12' will also be created in the target Secret resource, encrypted using the password stored in 'passwordSecretRef' containing the issuing Certificate Authority",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"password_secret_ref": schema.SingleNestedAttribute{
										Description:         "PasswordSecretRef is a reference to a key in a Secret resource containing the password used to encrypt the PKCS12 keystore.",
										MarkdownDescription: "PasswordSecretRef is a reference to a key in a Secret resource containing the password used to encrypt the PKCS12 keystore.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
												MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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

					"literal_subject": schema.StringAttribute{
						Description:         "LiteralSubject is an LDAP formatted string that represents the [X.509 Subject field](https://datatracker.ietf.org/doc/html/rfc5280#section-4.1.2.6). Use this *instead* of the Subject field if you need to ensure the correct ordering of the RDN sequence, such as when issuing certs for LDAP authentication. See https://github.com/cert-manager/cert-manager/issues/3203, https://github.com/cert-manager/cert-manager/issues/4424. This field is alpha level and is only supported by cert-manager installations where LiteralCertificateSubject feature gate is enabled on both cert-manager controller and webhook.",
						MarkdownDescription: "LiteralSubject is an LDAP formatted string that represents the [X.509 Subject field](https://datatracker.ietf.org/doc/html/rfc5280#section-4.1.2.6). Use this *instead* of the Subject field if you need to ensure the correct ordering of the RDN sequence, such as when issuing certs for LDAP authentication. See https://github.com/cert-manager/cert-manager/issues/3203, https://github.com/cert-manager/cert-manager/issues/4424. This field is alpha level and is only supported by cert-manager installations where LiteralCertificateSubject feature gate is enabled on both cert-manager controller and webhook.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"private_key": schema.SingleNestedAttribute{
						Description:         "Options to control private keys used for the Certificate.",
						MarkdownDescription: "Options to control private keys used for the Certificate.",
						Attributes: map[string]schema.Attribute{
							"algorithm": schema.StringAttribute{
								Description:         "Algorithm is the private key algorithm of the corresponding private key for this certificate. If provided, allowed values are either 'RSA','Ed25519' or 'ECDSA' If 'algorithm' is specified and 'size' is not provided, key size of 256 will be used for 'ECDSA' key algorithm and key size of 2048 will be used for 'RSA' key algorithm. key size is ignored when using the 'Ed25519' key algorithm.",
								MarkdownDescription: "Algorithm is the private key algorithm of the corresponding private key for this certificate. If provided, allowed values are either 'RSA','Ed25519' or 'ECDSA' If 'algorithm' is specified and 'size' is not provided, key size of 256 will be used for 'ECDSA' key algorithm and key size of 2048 will be used for 'RSA' key algorithm. key size is ignored when using the 'Ed25519' key algorithm.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"encoding": schema.StringAttribute{
								Description:         "The private key cryptography standards (PKCS) encoding for this certificate's private key to be encoded in. If provided, allowed values are 'PKCS1' and 'PKCS8' standing for PKCS#1 and PKCS#8, respectively. Defaults to 'PKCS1' if not specified.",
								MarkdownDescription: "The private key cryptography standards (PKCS) encoding for this certificate's private key to be encoded in. If provided, allowed values are 'PKCS1' and 'PKCS8' standing for PKCS#1 and PKCS#8, respectively. Defaults to 'PKCS1' if not specified.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"rotation_policy": schema.StringAttribute{
								Description:         "RotationPolicy controls how private keys should be regenerated when a re-issuance is being processed. If set to Never, a private key will only be generated if one does not already exist in the target 'spec.secretName'. If one does exists but it does not have the correct algorithm or size, a warning will be raised to await user intervention. If set to Always, a private key matching the specified requirements will be generated whenever a re-issuance occurs. Default is 'Never' for backward compatibility.",
								MarkdownDescription: "RotationPolicy controls how private keys should be regenerated when a re-issuance is being processed. If set to Never, a private key will only be generated if one does not already exist in the target 'spec.secretName'. If one does exists but it does not have the correct algorithm or size, a warning will be raised to await user intervention. If set to Always, a private key matching the specified requirements will be generated whenever a re-issuance occurs. Default is 'Never' for backward compatibility.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"size": schema.Int64Attribute{
								Description:         "Size is the key bit size of the corresponding private key for this certificate. If 'algorithm' is set to 'RSA', valid values are '2048', '4096' or '8192', and will default to '2048' if not specified. If 'algorithm' is set to 'ECDSA', valid values are '256', '384' or '521', and will default to '256' if not specified. If 'algorithm' is set to 'Ed25519', Size is ignored. No other values are allowed.",
								MarkdownDescription: "Size is the key bit size of the corresponding private key for this certificate. If 'algorithm' is set to 'RSA', valid values are '2048', '4096' or '8192', and will default to '2048' if not specified. If 'algorithm' is set to 'ECDSA', valid values are '256', '384' or '521', and will default to '256' if not specified. If 'algorithm' is set to 'Ed25519', Size is ignored. No other values are allowed.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"renew_before": schema.StringAttribute{
						Description:         "How long before the currently issued certificate's expiry cert-manager should renew the certificate. The default is 2/3 of the issued certificate's duration. Minimum accepted value is 5 minutes. Value must be in units accepted by Go time.ParseDuration https://golang.org/pkg/time/#ParseDuration",
						MarkdownDescription: "How long before the currently issued certificate's expiry cert-manager should renew the certificate. The default is 2/3 of the issued certificate's duration. Minimum accepted value is 5 minutes. Value must be in units accepted by Go time.ParseDuration https://golang.org/pkg/time/#ParseDuration",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"revision_history_limit": schema.Int64Attribute{
						Description:         "revisionHistoryLimit is the maximum number of CertificateRequest revisions that are maintained in the Certificate's history. Each revision represents a single 'CertificateRequest' created by this Certificate, either when it was created, renewed, or Spec was changed. Revisions will be removed by oldest first if the number of revisions exceeds this number. If set, revisionHistoryLimit must be a value of '1' or greater. If unset ('nil'), revisions will not be garbage collected. Default value is 'nil'.",
						MarkdownDescription: "revisionHistoryLimit is the maximum number of CertificateRequest revisions that are maintained in the Certificate's history. Each revision represents a single 'CertificateRequest' created by this Certificate, either when it was created, renewed, or Spec was changed. Revisions will be removed by oldest first if the number of revisions exceeds this number. If set, revisionHistoryLimit must be a value of '1' or greater. If unset ('nil'), revisions will not be garbage collected. Default value is 'nil'.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"secret_name": schema.StringAttribute{
						Description:         "SecretName is the name of the secret resource that will be automatically created and managed by this Certificate resource. It will be populated with a private key and certificate, signed by the denoted issuer.",
						MarkdownDescription: "SecretName is the name of the secret resource that will be automatically created and managed by this Certificate resource. It will be populated with a private key and certificate, signed by the denoted issuer.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"secret_template": schema.SingleNestedAttribute{
						Description:         "SecretTemplate defines annotations and labels to be copied to the Certificate's Secret. Labels and annotations on the Secret will be changed as they appear on the SecretTemplate when added or removed. SecretTemplate annotations are added in conjunction with, and cannot overwrite, the base set of annotations cert-manager sets on the Certificate's Secret.",
						MarkdownDescription: "SecretTemplate defines annotations and labels to be copied to the Certificate's Secret. Labels and annotations on the Secret will be changed as they appear on the SecretTemplate when added or removed. SecretTemplate annotations are added in conjunction with, and cannot overwrite, the base set of annotations cert-manager sets on the Certificate's Secret.",
						Attributes: map[string]schema.Attribute{
							"annotations": schema.MapAttribute{
								Description:         "Annotations is a key value map to be copied to the target Kubernetes Secret.",
								MarkdownDescription: "Annotations is a key value map to be copied to the target Kubernetes Secret.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"labels": schema.MapAttribute{
								Description:         "Labels is a key value map to be copied to the target Kubernetes Secret.",
								MarkdownDescription: "Labels is a key value map to be copied to the target Kubernetes Secret.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"subject": schema.SingleNestedAttribute{
						Description:         "Full X509 name specification (https://golang.org/pkg/crypto/x509/pkix/#Name).",
						MarkdownDescription: "Full X509 name specification (https://golang.org/pkg/crypto/x509/pkix/#Name).",
						Attributes: map[string]schema.Attribute{
							"countries": schema.ListAttribute{
								Description:         "Countries to be used on the Certificate.",
								MarkdownDescription: "Countries to be used on the Certificate.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"localities": schema.ListAttribute{
								Description:         "Cities to be used on the Certificate.",
								MarkdownDescription: "Cities to be used on the Certificate.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"organizational_units": schema.ListAttribute{
								Description:         "Organizational Units to be used on the Certificate.",
								MarkdownDescription: "Organizational Units to be used on the Certificate.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"organizations": schema.ListAttribute{
								Description:         "Organizations to be used on the Certificate.",
								MarkdownDescription: "Organizations to be used on the Certificate.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"postal_codes": schema.ListAttribute{
								Description:         "Postal codes to be used on the Certificate.",
								MarkdownDescription: "Postal codes to be used on the Certificate.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"provinces": schema.ListAttribute{
								Description:         "State/Provinces to be used on the Certificate.",
								MarkdownDescription: "State/Provinces to be used on the Certificate.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"serial_number": schema.StringAttribute{
								Description:         "Serial number to be used on the Certificate.",
								MarkdownDescription: "Serial number to be used on the Certificate.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"street_addresses": schema.ListAttribute{
								Description:         "Street addresses to be used on the Certificate.",
								MarkdownDescription: "Street addresses to be used on the Certificate.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"uris": schema.ListAttribute{
						Description:         "URIs is a list of URI subjectAltNames to be set on the Certificate.",
						MarkdownDescription: "URIs is a list of URI subjectAltNames to be set on the Certificate.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"usages": schema.ListAttribute{
						Description:         "Usages is the set of x509 usages that are requested for the certificate. Defaults to 'digital signature' and 'key encipherment' if not specified.",
						MarkdownDescription: "Usages is the set of x509 usages that are requested for the certificate. Defaults to 'digital signature' and 'key encipherment' if not specified.",
						ElementType:         types.StringType,
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
	}
}

func (r *CertManagerIoCertificateV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *CertManagerIoCertificateV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_cert_manager_io_certificate_v1")

	var data CertManagerIoCertificateV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "cert-manager.io", Version: "v1", Resource: "Certificate"}).
		Namespace(data.Metadata.Namespace).
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

	var readResponse CertManagerIoCertificateV1DataSourceData
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
	data.ApiVersion = pointer.String("cert-manager.io/v1")
	data.Kind = pointer.String("Certificate")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}