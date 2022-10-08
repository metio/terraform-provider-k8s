/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"gopkg.in/yaml.v3"
	"time"
)

type CertManagerIoCertificateV1Resource struct{}

var (
	_ resource.Resource = (*CertManagerIoCertificateV1Resource)(nil)
)

type CertManagerIoCertificateV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type CertManagerIoCertificateV1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		AdditionalOutputFormats *[]struct {
			Type *string `tfsdk:"type" yaml:"type,omitempty"`
		} `tfsdk:"additional_output_formats" yaml:"additionalOutputFormats,omitempty"`

		CommonName *string `tfsdk:"common_name" yaml:"commonName,omitempty"`

		DnsNames *[]string `tfsdk:"dns_names" yaml:"dnsNames,omitempty"`

		Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

		EmailAddresses *[]string `tfsdk:"email_addresses" yaml:"emailAddresses,omitempty"`

		EncodeUsagesInRequest *bool `tfsdk:"encode_usages_in_request" yaml:"encodeUsagesInRequest,omitempty"`

		IpAddresses *[]string `tfsdk:"ip_addresses" yaml:"ipAddresses,omitempty"`

		IsCA *bool `tfsdk:"is_ca" yaml:"isCA,omitempty"`

		IssuerRef *struct {
			Group *string `tfsdk:"group" yaml:"group,omitempty"`

			Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"issuer_ref" yaml:"issuerRef,omitempty"`

		Keystores *struct {
			Jks *struct {
				Create *bool `tfsdk:"create" yaml:"create,omitempty"`

				PasswordSecretRef *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"password_secret_ref" yaml:"passwordSecretRef,omitempty"`
			} `tfsdk:"jks" yaml:"jks,omitempty"`

			Pkcs12 *struct {
				Create *bool `tfsdk:"create" yaml:"create,omitempty"`

				PasswordSecretRef *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"password_secret_ref" yaml:"passwordSecretRef,omitempty"`
			} `tfsdk:"pkcs12" yaml:"pkcs12,omitempty"`
		} `tfsdk:"keystores" yaml:"keystores,omitempty"`

		LiteralSubject *string `tfsdk:"literal_subject" yaml:"literalSubject,omitempty"`

		PrivateKey *struct {
			Algorithm *string `tfsdk:"algorithm" yaml:"algorithm,omitempty"`

			Encoding *string `tfsdk:"encoding" yaml:"encoding,omitempty"`

			RotationPolicy *string `tfsdk:"rotation_policy" yaml:"rotationPolicy,omitempty"`

			Size *int64 `tfsdk:"size" yaml:"size,omitempty"`
		} `tfsdk:"private_key" yaml:"privateKey,omitempty"`

		RenewBefore *string `tfsdk:"renew_before" yaml:"renewBefore,omitempty"`

		RevisionHistoryLimit *int64 `tfsdk:"revision_history_limit" yaml:"revisionHistoryLimit,omitempty"`

		SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`

		SecretTemplate *struct {
			Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

			Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`
		} `tfsdk:"secret_template" yaml:"secretTemplate,omitempty"`

		Subject *struct {
			Countries *[]string `tfsdk:"countries" yaml:"countries,omitempty"`

			Localities *[]string `tfsdk:"localities" yaml:"localities,omitempty"`

			OrganizationalUnits *[]string `tfsdk:"organizational_units" yaml:"organizationalUnits,omitempty"`

			Organizations *[]string `tfsdk:"organizations" yaml:"organizations,omitempty"`

			PostalCodes *[]string `tfsdk:"postal_codes" yaml:"postalCodes,omitempty"`

			Provinces *[]string `tfsdk:"provinces" yaml:"provinces,omitempty"`

			SerialNumber *string `tfsdk:"serial_number" yaml:"serialNumber,omitempty"`

			StreetAddresses *[]string `tfsdk:"street_addresses" yaml:"streetAddresses,omitempty"`
		} `tfsdk:"subject" yaml:"subject,omitempty"`

		Uris *[]string `tfsdk:"uris" yaml:"uris,omitempty"`

		Usages *[]string `tfsdk:"usages" yaml:"usages,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewCertManagerIoCertificateV1Resource() resource.Resource {
	return &CertManagerIoCertificateV1Resource{}
}

func (r *CertManagerIoCertificateV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cert_manager_io_certificate_v1"
}

func (r *CertManagerIoCertificateV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "A Certificate resource should be created to ensure an up to date and signed x509 certificate is stored in the Kubernetes Secret resource named in 'spec.secretName'.  The stored certificate will be renewed before it expires (as configured by 'spec.renewBefore').",
		MarkdownDescription: "A Certificate resource should be created to ensure an up to date and signed x509 certificate is stored in the Kubernetes Secret resource named in 'spec.secretName'.  The stored certificate will be renewed before it expires (as configured by 'spec.renewBefore').",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description:         "The timestamp of the last change to this resource.",
				MarkdownDescription: "The timestamp of the last change to this resource.",
				Type:                types.Int64Type,
				Computed:            true,
				Optional:            false,
			},

			"yaml": {
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"metadata": {
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Type:                types.StringType,
						Required:            true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.RequiresReplace(),
						},
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
					},

					"namespace": {
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Type:                types.StringType,
						Optional:            true,
					},

					"labels": {
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.LabelValidator(),
						},
					},
					"annotations": {
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.AnnotationValidator(),
						},
					},
				}),
			},

			"api_version": {
				Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"kind": {
				Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"spec": {
				Description:         "Desired state of the Certificate resource.",
				MarkdownDescription: "Desired state of the Certificate resource.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"additional_output_formats": {
						Description:         "AdditionalOutputFormats defines extra output formats of the private key and signed certificate chain to be written to this Certificate's target Secret. This is an Alpha Feature and is only enabled with the '--feature-gates=AdditionalCertificateOutputFormats=true' option on both the controller and webhook components.",
						MarkdownDescription: "AdditionalOutputFormats defines extra output formats of the private key and signed certificate chain to be written to this Certificate's target Secret. This is an Alpha Feature and is only enabled with the '--feature-gates=AdditionalCertificateOutputFormats=true' option on both the controller and webhook components.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"type": {
								Description:         "Type is the name of the format type that should be written to the Certificate's target Secret.",
								MarkdownDescription: "Type is the name of the format type that should be written to the Certificate's target Secret.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"common_name": {
						Description:         "CommonName is a common name to be used on the Certificate. The CommonName should have a length of 64 characters or fewer to avoid generating invalid CSRs. This value is ignored by TLS clients when any subject alt name is set. This is x509 behaviour: https://tools.ietf.org/html/rfc6125#section-6.4.4",
						MarkdownDescription: "CommonName is a common name to be used on the Certificate. The CommonName should have a length of 64 characters or fewer to avoid generating invalid CSRs. This value is ignored by TLS clients when any subject alt name is set. This is x509 behaviour: https://tools.ietf.org/html/rfc6125#section-6.4.4",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"dns_names": {
						Description:         "DNSNames is a list of DNS subjectAltNames to be set on the Certificate.",
						MarkdownDescription: "DNSNames is a list of DNS subjectAltNames to be set on the Certificate.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"duration": {
						Description:         "The requested 'duration' (i.e. lifetime) of the Certificate. This option may be ignored/overridden by some issuer types. If unset this defaults to 90 days. Certificate will be renewed either 2/3 through its duration or 'renewBefore' period before its expiry, whichever is later. Minimum accepted duration is 1 hour. Value must be in units accepted by Go time.ParseDuration https://golang.org/pkg/time/#ParseDuration",
						MarkdownDescription: "The requested 'duration' (i.e. lifetime) of the Certificate. This option may be ignored/overridden by some issuer types. If unset this defaults to 90 days. Certificate will be renewed either 2/3 through its duration or 'renewBefore' period before its expiry, whichever is later. Minimum accepted duration is 1 hour. Value must be in units accepted by Go time.ParseDuration https://golang.org/pkg/time/#ParseDuration",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"email_addresses": {
						Description:         "EmailAddresses is a list of email subjectAltNames to be set on the Certificate.",
						MarkdownDescription: "EmailAddresses is a list of email subjectAltNames to be set on the Certificate.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"encode_usages_in_request": {
						Description:         "EncodeUsagesInRequest controls whether key usages should be present in the CertificateRequest",
						MarkdownDescription: "EncodeUsagesInRequest controls whether key usages should be present in the CertificateRequest",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ip_addresses": {
						Description:         "IPAddresses is a list of IP address subjectAltNames to be set on the Certificate.",
						MarkdownDescription: "IPAddresses is a list of IP address subjectAltNames to be set on the Certificate.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"is_ca": {
						Description:         "IsCA will mark this Certificate as valid for certificate signing. This will automatically add the 'cert sign' usage to the list of 'usages'.",
						MarkdownDescription: "IsCA will mark this Certificate as valid for certificate signing. This will automatically add the 'cert sign' usage to the list of 'usages'.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"issuer_ref": {
						Description:         "IssuerRef is a reference to the issuer for this certificate. If the 'kind' field is not set, or set to 'Issuer', an Issuer resource with the given name in the same namespace as the Certificate will be used. If the 'kind' field is set to 'ClusterIssuer', a ClusterIssuer with the provided name will be used. The 'name' field in this stanza is required at all times.",
						MarkdownDescription: "IssuerRef is a reference to the issuer for this certificate. If the 'kind' field is not set, or set to 'Issuer', an Issuer resource with the given name in the same namespace as the Certificate will be used. If the 'kind' field is set to 'ClusterIssuer', a ClusterIssuer with the provided name will be used. The 'name' field in this stanza is required at all times.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"group": {
								Description:         "Group of the resource being referred to.",
								MarkdownDescription: "Group of the resource being referred to.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"kind": {
								Description:         "Kind of the resource being referred to.",
								MarkdownDescription: "Kind of the resource being referred to.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "Name of the resource being referred to.",
								MarkdownDescription: "Name of the resource being referred to.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"keystores": {
						Description:         "Keystores configures additional keystore output formats stored in the 'secretName' Secret resource.",
						MarkdownDescription: "Keystores configures additional keystore output formats stored in the 'secretName' Secret resource.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"jks": {
								Description:         "JKS configures options for storing a JKS keystore in the 'spec.secretName' Secret resource.",
								MarkdownDescription: "JKS configures options for storing a JKS keystore in the 'spec.secretName' Secret resource.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"create": {
										Description:         "Create enables JKS keystore creation for the Certificate. If true, a file named 'keystore.jks' will be created in the target Secret resource, encrypted using the password stored in 'passwordSecretRef'. The keystore file will only be updated upon re-issuance. A file named 'truststore.jks' will also be created in the target Secret resource, encrypted using the password stored in 'passwordSecretRef' containing the issuing Certificate Authority",
										MarkdownDescription: "Create enables JKS keystore creation for the Certificate. If true, a file named 'keystore.jks' will be created in the target Secret resource, encrypted using the password stored in 'passwordSecretRef'. The keystore file will only be updated upon re-issuance. A file named 'truststore.jks' will also be created in the target Secret resource, encrypted using the password stored in 'passwordSecretRef' containing the issuing Certificate Authority",

										Type: types.BoolType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"password_secret_ref": {
										Description:         "PasswordSecretRef is a reference to a key in a Secret resource containing the password used to encrypt the JKS keystore.",
										MarkdownDescription: "PasswordSecretRef is a reference to a key in a Secret resource containing the password used to encrypt the JKS keystore.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
												MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pkcs12": {
								Description:         "PKCS12 configures options for storing a PKCS12 keystore in the 'spec.secretName' Secret resource.",
								MarkdownDescription: "PKCS12 configures options for storing a PKCS12 keystore in the 'spec.secretName' Secret resource.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"create": {
										Description:         "Create enables PKCS12 keystore creation for the Certificate. If true, a file named 'keystore.p12' will be created in the target Secret resource, encrypted using the password stored in 'passwordSecretRef'. The keystore file will only be updated upon re-issuance. A file named 'truststore.p12' will also be created in the target Secret resource, encrypted using the password stored in 'passwordSecretRef' containing the issuing Certificate Authority",
										MarkdownDescription: "Create enables PKCS12 keystore creation for the Certificate. If true, a file named 'keystore.p12' will be created in the target Secret resource, encrypted using the password stored in 'passwordSecretRef'. The keystore file will only be updated upon re-issuance. A file named 'truststore.p12' will also be created in the target Secret resource, encrypted using the password stored in 'passwordSecretRef' containing the issuing Certificate Authority",

										Type: types.BoolType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"password_secret_ref": {
										Description:         "PasswordSecretRef is a reference to a key in a Secret resource containing the password used to encrypt the PKCS12 keystore.",
										MarkdownDescription: "PasswordSecretRef is a reference to a key in a Secret resource containing the password used to encrypt the PKCS12 keystore.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
												MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"literal_subject": {
						Description:         "LiteralSubject is an LDAP formatted string that represents the [X.509 Subject field](https://datatracker.ietf.org/doc/html/rfc5280#section-4.1.2.6). Use this *instead* of the Subject field if you need to ensure the correct ordering of the RDN sequence, such as when issuing certs for LDAP authentication. See https://github.com/cert-manager/cert-manager/issues/3203, https://github.com/cert-manager/cert-manager/issues/4424. This field is alpha level and is only supported by cert-manager installations where LiteralCertificateSubject feature gate is enabled on both cert-manager controller and webhook.",
						MarkdownDescription: "LiteralSubject is an LDAP formatted string that represents the [X.509 Subject field](https://datatracker.ietf.org/doc/html/rfc5280#section-4.1.2.6). Use this *instead* of the Subject field if you need to ensure the correct ordering of the RDN sequence, such as when issuing certs for LDAP authentication. See https://github.com/cert-manager/cert-manager/issues/3203, https://github.com/cert-manager/cert-manager/issues/4424. This field is alpha level and is only supported by cert-manager installations where LiteralCertificateSubject feature gate is enabled on both cert-manager controller and webhook.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"private_key": {
						Description:         "Options to control private keys used for the Certificate.",
						MarkdownDescription: "Options to control private keys used for the Certificate.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"algorithm": {
								Description:         "Algorithm is the private key algorithm of the corresponding private key for this certificate. If provided, allowed values are either 'RSA','Ed25519' or 'ECDSA' If 'algorithm' is specified and 'size' is not provided, key size of 256 will be used for 'ECDSA' key algorithm and key size of 2048 will be used for 'RSA' key algorithm. key size is ignored when using the 'Ed25519' key algorithm.",
								MarkdownDescription: "Algorithm is the private key algorithm of the corresponding private key for this certificate. If provided, allowed values are either 'RSA','Ed25519' or 'ECDSA' If 'algorithm' is specified and 'size' is not provided, key size of 256 will be used for 'ECDSA' key algorithm and key size of 2048 will be used for 'RSA' key algorithm. key size is ignored when using the 'Ed25519' key algorithm.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"encoding": {
								Description:         "The private key cryptography standards (PKCS) encoding for this certificate's private key to be encoded in. If provided, allowed values are 'PKCS1' and 'PKCS8' standing for PKCS#1 and PKCS#8, respectively. Defaults to 'PKCS1' if not specified.",
								MarkdownDescription: "The private key cryptography standards (PKCS) encoding for this certificate's private key to be encoded in. If provided, allowed values are 'PKCS1' and 'PKCS8' standing for PKCS#1 and PKCS#8, respectively. Defaults to 'PKCS1' if not specified.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"rotation_policy": {
								Description:         "RotationPolicy controls how private keys should be regenerated when a re-issuance is being processed. If set to Never, a private key will only be generated if one does not already exist in the target 'spec.secretName'. If one does exists but it does not have the correct algorithm or size, a warning will be raised to await user intervention. If set to Always, a private key matching the specified requirements will be generated whenever a re-issuance occurs. Default is 'Never' for backward compatibility.",
								MarkdownDescription: "RotationPolicy controls how private keys should be regenerated when a re-issuance is being processed. If set to Never, a private key will only be generated if one does not already exist in the target 'spec.secretName'. If one does exists but it does not have the correct algorithm or size, a warning will be raised to await user intervention. If set to Always, a private key matching the specified requirements will be generated whenever a re-issuance occurs. Default is 'Never' for backward compatibility.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"size": {
								Description:         "Size is the key bit size of the corresponding private key for this certificate. If 'algorithm' is set to 'RSA', valid values are '2048', '4096' or '8192', and will default to '2048' if not specified. If 'algorithm' is set to 'ECDSA', valid values are '256', '384' or '521', and will default to '256' if not specified. If 'algorithm' is set to 'Ed25519', Size is ignored. No other values are allowed.",
								MarkdownDescription: "Size is the key bit size of the corresponding private key for this certificate. If 'algorithm' is set to 'RSA', valid values are '2048', '4096' or '8192', and will default to '2048' if not specified. If 'algorithm' is set to 'ECDSA', valid values are '256', '384' or '521', and will default to '256' if not specified. If 'algorithm' is set to 'Ed25519', Size is ignored. No other values are allowed.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"renew_before": {
						Description:         "How long before the currently issued certificate's expiry cert-manager should renew the certificate. The default is 2/3 of the issued certificate's duration. Minimum accepted value is 5 minutes. Value must be in units accepted by Go time.ParseDuration https://golang.org/pkg/time/#ParseDuration",
						MarkdownDescription: "How long before the currently issued certificate's expiry cert-manager should renew the certificate. The default is 2/3 of the issued certificate's duration. Minimum accepted value is 5 minutes. Value must be in units accepted by Go time.ParseDuration https://golang.org/pkg/time/#ParseDuration",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"revision_history_limit": {
						Description:         "revisionHistoryLimit is the maximum number of CertificateRequest revisions that are maintained in the Certificate's history. Each revision represents a single 'CertificateRequest' created by this Certificate, either when it was created, renewed, or Spec was changed. Revisions will be removed by oldest first if the number of revisions exceeds this number. If set, revisionHistoryLimit must be a value of '1' or greater. If unset ('nil'), revisions will not be garbage collected. Default value is 'nil'.",
						MarkdownDescription: "revisionHistoryLimit is the maximum number of CertificateRequest revisions that are maintained in the Certificate's history. Each revision represents a single 'CertificateRequest' created by this Certificate, either when it was created, renewed, or Spec was changed. Revisions will be removed by oldest first if the number of revisions exceeds this number. If set, revisionHistoryLimit must be a value of '1' or greater. If unset ('nil'), revisions will not be garbage collected. Default value is 'nil'.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"secret_name": {
						Description:         "SecretName is the name of the secret resource that will be automatically created and managed by this Certificate resource. It will be populated with a private key and certificate, signed by the denoted issuer.",
						MarkdownDescription: "SecretName is the name of the secret resource that will be automatically created and managed by this Certificate resource. It will be populated with a private key and certificate, signed by the denoted issuer.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"secret_template": {
						Description:         "SecretTemplate defines annotations and labels to be copied to the Certificate's Secret. Labels and annotations on the Secret will be changed as they appear on the SecretTemplate when added or removed. SecretTemplate annotations are added in conjunction with, and cannot overwrite, the base set of annotations cert-manager sets on the Certificate's Secret.",
						MarkdownDescription: "SecretTemplate defines annotations and labels to be copied to the Certificate's Secret. Labels and annotations on the Secret will be changed as they appear on the SecretTemplate when added or removed. SecretTemplate annotations are added in conjunction with, and cannot overwrite, the base set of annotations cert-manager sets on the Certificate's Secret.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"annotations": {
								Description:         "Annotations is a key value map to be copied to the target Kubernetes Secret.",
								MarkdownDescription: "Annotations is a key value map to be copied to the target Kubernetes Secret.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"labels": {
								Description:         "Labels is a key value map to be copied to the target Kubernetes Secret.",
								MarkdownDescription: "Labels is a key value map to be copied to the target Kubernetes Secret.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"subject": {
						Description:         "Full X509 name specification (https://golang.org/pkg/crypto/x509/pkix/#Name).",
						MarkdownDescription: "Full X509 name specification (https://golang.org/pkg/crypto/x509/pkix/#Name).",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"countries": {
								Description:         "Countries to be used on the Certificate.",
								MarkdownDescription: "Countries to be used on the Certificate.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"localities": {
								Description:         "Cities to be used on the Certificate.",
								MarkdownDescription: "Cities to be used on the Certificate.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"organizational_units": {
								Description:         "Organizational Units to be used on the Certificate.",
								MarkdownDescription: "Organizational Units to be used on the Certificate.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"organizations": {
								Description:         "Organizations to be used on the Certificate.",
								MarkdownDescription: "Organizations to be used on the Certificate.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"postal_codes": {
								Description:         "Postal codes to be used on the Certificate.",
								MarkdownDescription: "Postal codes to be used on the Certificate.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"provinces": {
								Description:         "State/Provinces to be used on the Certificate.",
								MarkdownDescription: "State/Provinces to be used on the Certificate.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"serial_number": {
								Description:         "Serial number to be used on the Certificate.",
								MarkdownDescription: "Serial number to be used on the Certificate.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"street_addresses": {
								Description:         "Street addresses to be used on the Certificate.",
								MarkdownDescription: "Street addresses to be used on the Certificate.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"uris": {
						Description:         "URIs is a list of URI subjectAltNames to be set on the Certificate.",
						MarkdownDescription: "URIs is a list of URI subjectAltNames to be set on the Certificate.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"usages": {
						Description:         "Usages is the set of x509 usages that are requested for the certificate. Defaults to 'digital signature' and 'key encipherment' if not specified.",
						MarkdownDescription: "Usages is the set of x509 usages that are requested for the certificate. Defaults to 'digital signature' and 'key encipherment' if not specified.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}, nil
}

func (r *CertManagerIoCertificateV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_cert_manager_io_certificate_v1")

	var state CertManagerIoCertificateV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CertManagerIoCertificateV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("cert-manager.io/v1")
	goModel.Kind = utilities.Ptr("Certificate")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *CertManagerIoCertificateV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_cert_manager_io_certificate_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *CertManagerIoCertificateV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_cert_manager_io_certificate_v1")

	var state CertManagerIoCertificateV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CertManagerIoCertificateV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("cert-manager.io/v1")
	goModel.Kind = utilities.Ptr("Certificate")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *CertManagerIoCertificateV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_cert_manager_io_certificate_v1")
	// NO-OP: Terraform removes the state automatically for us
}
