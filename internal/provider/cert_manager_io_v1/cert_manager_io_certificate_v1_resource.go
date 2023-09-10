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
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	k8sTypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"strings"
)

var (
	_ resource.Resource                = &CertManagerIoCertificateV1Resource{}
	_ resource.ResourceWithConfigure   = &CertManagerIoCertificateV1Resource{}
	_ resource.ResourceWithImportState = &CertManagerIoCertificateV1Resource{}
)

func NewCertManagerIoCertificateV1Resource() resource.Resource {
	return &CertManagerIoCertificateV1Resource{}
}

type CertManagerIoCertificateV1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type CertManagerIoCertificateV1ResourceData struct {
	ID             types.String `tfsdk:"id" json:"-"`
	ForceConflicts types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager   types.String `tfsdk:"field_manager" json:"-"`
	WaitFor        types.List   `tfsdk:"wait_for" json:"-"`

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

func (r *CertManagerIoCertificateV1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_cert_manager_io_certificate_v1"
}

func (r *CertManagerIoCertificateV1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "A Certificate resource should be created to ensure an up to date and signed X.509 certificate is stored in the Kubernetes Secret resource named in 'spec.secretName'.  The stored certificate will be renewed before it expires (as configured by 'spec.renewBefore').",
		MarkdownDescription: "A Certificate resource should be created to ensure an up to date and signed X.509 certificate is stored in the Kubernetes Secret resource named in 'spec.secretName'.  The stored certificate will be renewed before it expires (as configured by 'spec.renewBefore').",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"force_conflicts": schema.BoolAttribute{
				Description:         "If 'true', server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "If `true`, server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"field_manager": schema.BoolAttribute{
				Description:         "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"wait_for": schema.ListNestedAttribute{
				Description:         "Wait for specific conditions after create/update of resources.",
				MarkdownDescription: "Wait for specific conditions after create/update of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"jsonpath": schema.StringAttribute{
							Description:         "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							MarkdownDescription: "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							Required:            true,
							Optional:            false,
							Computed:            false,
						},
						"value": schema.StringAttribute{
							Description:         "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							MarkdownDescription: "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							Required:            false,
							Optional:            true,
							Computed:            true,
						},
						"timeout": schema.StringAttribute{
							Description:         "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							MarkdownDescription: "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("30s"),
						},
					},
				},
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},

					"labels": schema.MapAttribute{
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            true,
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
						Computed:            true,
						Validators: []validator.Map{
							validators.AnnotationValidator(),
						},
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "Specification of the desired state of the Certificate resource. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
				MarkdownDescription: "Specification of the desired state of the Certificate resource. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
				Attributes: map[string]schema.Attribute{
					"additional_output_formats": schema.ListNestedAttribute{
						Description:         "Defines extra output formats of the private key and signed certificate chain to be written to this Certificate's target Secret.  This is an Alpha Feature and is only enabled with the '--feature-gates=AdditionalCertificateOutputFormats=true' option set on both the controller and webhook components.",
						MarkdownDescription: "Defines extra output formats of the private key and signed certificate chain to be written to this Certificate's target Secret.  This is an Alpha Feature and is only enabled with the '--feature-gates=AdditionalCertificateOutputFormats=true' option set on both the controller and webhook components.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"type": schema.StringAttribute{
									Description:         "Type is the name of the format type that should be written to the Certificate's target Secret.",
									MarkdownDescription: "Type is the name of the format type that should be written to the Certificate's target Secret.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("DER", "CombinedPEM"),
									},
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"common_name": schema.StringAttribute{
						Description:         "Requested common name X509 certificate subject attribute. More info: https://datatracker.ietf.org/doc/html/rfc5280#section-4.1.2.6 NOTE: TLS clients will ignore this value when any subject alternative name is set (see https://tools.ietf.org/html/rfc6125#section-6.4.4).  Should have a length of 64 characters or fewer to avoid generating invalid CSRs. Cannot be set if the 'literalSubject' field is set.",
						MarkdownDescription: "Requested common name X509 certificate subject attribute. More info: https://datatracker.ietf.org/doc/html/rfc5280#section-4.1.2.6 NOTE: TLS clients will ignore this value when any subject alternative name is set (see https://tools.ietf.org/html/rfc6125#section-6.4.4).  Should have a length of 64 characters or fewer to avoid generating invalid CSRs. Cannot be set if the 'literalSubject' field is set.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"dns_names": schema.ListAttribute{
						Description:         "Requested DNS subject alternative names.",
						MarkdownDescription: "Requested DNS subject alternative names.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"duration": schema.StringAttribute{
						Description:         "Requested 'duration' (i.e. lifetime) of the Certificate. Note that the issuer may choose to ignore the requested duration, just like any other requested attribute.  If unset, this defaults to 90 days. Minimum accepted duration is 1 hour. Value must be in units accepted by Go time.ParseDuration https://golang.org/pkg/time/#ParseDuration.",
						MarkdownDescription: "Requested 'duration' (i.e. lifetime) of the Certificate. Note that the issuer may choose to ignore the requested duration, just like any other requested attribute.  If unset, this defaults to 90 days. Minimum accepted duration is 1 hour. Value must be in units accepted by Go time.ParseDuration https://golang.org/pkg/time/#ParseDuration.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"email_addresses": schema.ListAttribute{
						Description:         "Requested email subject alternative names.",
						MarkdownDescription: "Requested email subject alternative names.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"encode_usages_in_request": schema.BoolAttribute{
						Description:         "Whether the KeyUsage and ExtKeyUsage extensions should be set in the encoded CSR.  This option defaults to true, and should only be disabled if the target issuer does not support CSRs with these X509 KeyUsage/ ExtKeyUsage extensions.",
						MarkdownDescription: "Whether the KeyUsage and ExtKeyUsage extensions should be set in the encoded CSR.  This option defaults to true, and should only be disabled if the target issuer does not support CSRs with these X509 KeyUsage/ ExtKeyUsage extensions.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ip_addresses": schema.ListAttribute{
						Description:         "Requested IP address subject alternative names.",
						MarkdownDescription: "Requested IP address subject alternative names.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"is_ca": schema.BoolAttribute{
						Description:         "Requested basic constraints isCA value. The isCA value is used to set the 'isCA' field on the created CertificateRequest resources. Note that the issuer may choose to ignore the requested isCA value, just like any other requested attribute.  If true, this will automatically add the 'cert sign' usage to the list of requested 'usages'.",
						MarkdownDescription: "Requested basic constraints isCA value. The isCA value is used to set the 'isCA' field on the created CertificateRequest resources. Note that the issuer may choose to ignore the requested isCA value, just like any other requested attribute.  If true, this will automatically add the 'cert sign' usage to the list of requested 'usages'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"issuer_ref": schema.SingleNestedAttribute{
						Description:         "Reference to the issuer responsible for issuing the certificate. If the issuer is namespace-scoped, it must be in the same namespace as the Certificate. If the issuer is cluster-scoped, it can be used from any namespace.  The 'name' field of the reference must always be specified.",
						MarkdownDescription: "Reference to the issuer responsible for issuing the certificate. If the issuer is namespace-scoped, it must be in the same namespace as the Certificate. If the issuer is cluster-scoped, it can be used from any namespace.  The 'name' field of the reference must always be specified.",
						Attributes: map[string]schema.Attribute{
							"group": schema.StringAttribute{
								Description:         "Group of the resource being referred to.",
								MarkdownDescription: "Group of the resource being referred to.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind of the resource being referred to.",
								MarkdownDescription: "Kind of the resource being referred to.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the resource being referred to.",
								MarkdownDescription: "Name of the resource being referred to.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"keystores": schema.SingleNestedAttribute{
						Description:         "Additional keystore output formats to be stored in the Certificate's Secret.",
						MarkdownDescription: "Additional keystore output formats to be stored in the Certificate's Secret.",
						Attributes: map[string]schema.Attribute{
							"jks": schema.SingleNestedAttribute{
								Description:         "JKS configures options for storing a JKS keystore in the 'spec.secretName' Secret resource.",
								MarkdownDescription: "JKS configures options for storing a JKS keystore in the 'spec.secretName' Secret resource.",
								Attributes: map[string]schema.Attribute{
									"create": schema.BoolAttribute{
										Description:         "Create enables JKS keystore creation for the Certificate. If true, a file named 'keystore.jks' will be created in the target Secret resource, encrypted using the password stored in 'passwordSecretRef'. The keystore file will be updated immediately. If the issuer provided a CA certificate, a file named 'truststore.jks' will also be created in the target Secret resource, encrypted using the password stored in 'passwordSecretRef' containing the issuing Certificate Authority",
										MarkdownDescription: "Create enables JKS keystore creation for the Certificate. If true, a file named 'keystore.jks' will be created in the target Secret resource, encrypted using the password stored in 'passwordSecretRef'. The keystore file will be updated immediately. If the issuer provided a CA certificate, a file named 'truststore.jks' will also be created in the target Secret resource, encrypted using the password stored in 'passwordSecretRef' containing the issuing Certificate Authority",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"password_secret_ref": schema.SingleNestedAttribute{
										Description:         "PasswordSecretRef is a reference to a key in a Secret resource containing the password used to encrypt the JKS keystore.",
										MarkdownDescription: "PasswordSecretRef is a reference to a key in a Secret resource containing the password used to encrypt the JKS keystore.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
												MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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
								Required: false,
								Optional: true,
								Computed: false,
							},

							"pkcs12": schema.SingleNestedAttribute{
								Description:         "PKCS12 configures options for storing a PKCS12 keystore in the 'spec.secretName' Secret resource.",
								MarkdownDescription: "PKCS12 configures options for storing a PKCS12 keystore in the 'spec.secretName' Secret resource.",
								Attributes: map[string]schema.Attribute{
									"create": schema.BoolAttribute{
										Description:         "Create enables PKCS12 keystore creation for the Certificate. If true, a file named 'keystore.p12' will be created in the target Secret resource, encrypted using the password stored in 'passwordSecretRef'. The keystore file will be updated immediately. If the issuer provided a CA certificate, a file named 'truststore.p12' will also be created in the target Secret resource, encrypted using the password stored in 'passwordSecretRef' containing the issuing Certificate Authority",
										MarkdownDescription: "Create enables PKCS12 keystore creation for the Certificate. If true, a file named 'keystore.p12' will be created in the target Secret resource, encrypted using the password stored in 'passwordSecretRef'. The keystore file will be updated immediately. If the issuer provided a CA certificate, a file named 'truststore.p12' will also be created in the target Secret resource, encrypted using the password stored in 'passwordSecretRef' containing the issuing Certificate Authority",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"password_secret_ref": schema.SingleNestedAttribute{
										Description:         "PasswordSecretRef is a reference to a key in a Secret resource containing the password used to encrypt the PKCS12 keystore.",
										MarkdownDescription: "PasswordSecretRef is a reference to a key in a Secret resource containing the password used to encrypt the PKCS12 keystore.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
												MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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
								Required: false,
								Optional: true,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"literal_subject": schema.StringAttribute{
						Description:         "Requested X.509 certificate subject, represented using the LDAP 'String Representation of a Distinguished Name' [1]. Important: the LDAP string format also specifies the order of the attributes in the subject, this is important when issuing certs for LDAP authentication. Example: 'CN=foo,DC=corp,DC=example,DC=com' More info [1]: https://datatracker.ietf.org/doc/html/rfc4514 More info: https://github.com/cert-manager/cert-manager/issues/3203 More info: https://github.com/cert-manager/cert-manager/issues/4424  Cannot be set if the 'subject' or 'commonName' field is set. This is an Alpha Feature and is only enabled with the '--feature-gates=LiteralCertificateSubject=true' option set on both the controller and webhook components.",
						MarkdownDescription: "Requested X.509 certificate subject, represented using the LDAP 'String Representation of a Distinguished Name' [1]. Important: the LDAP string format also specifies the order of the attributes in the subject, this is important when issuing certs for LDAP authentication. Example: 'CN=foo,DC=corp,DC=example,DC=com' More info [1]: https://datatracker.ietf.org/doc/html/rfc4514 More info: https://github.com/cert-manager/cert-manager/issues/3203 More info: https://github.com/cert-manager/cert-manager/issues/4424  Cannot be set if the 'subject' or 'commonName' field is set. This is an Alpha Feature and is only enabled with the '--feature-gates=LiteralCertificateSubject=true' option set on both the controller and webhook components.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"private_key": schema.SingleNestedAttribute{
						Description:         "Private key options. These include the key algorithm and size, the used encoding and the rotation policy.",
						MarkdownDescription: "Private key options. These include the key algorithm and size, the used encoding and the rotation policy.",
						Attributes: map[string]schema.Attribute{
							"algorithm": schema.StringAttribute{
								Description:         "Algorithm is the private key algorithm of the corresponding private key for this certificate.  If provided, allowed values are either 'RSA', 'ECDSA' or 'Ed25519'. If 'algorithm' is specified and 'size' is not provided, key size of 2048 will be used for 'RSA' key algorithm and key size of 256 will be used for 'ECDSA' key algorithm. key size is ignored when using the 'Ed25519' key algorithm.",
								MarkdownDescription: "Algorithm is the private key algorithm of the corresponding private key for this certificate.  If provided, allowed values are either 'RSA', 'ECDSA' or 'Ed25519'. If 'algorithm' is specified and 'size' is not provided, key size of 2048 will be used for 'RSA' key algorithm and key size of 256 will be used for 'ECDSA' key algorithm. key size is ignored when using the 'Ed25519' key algorithm.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("RSA", "ECDSA", "Ed25519"),
								},
							},

							"encoding": schema.StringAttribute{
								Description:         "The private key cryptography standards (PKCS) encoding for this certificate's private key to be encoded in.  If provided, allowed values are 'PKCS1' and 'PKCS8' standing for PKCS#1 and PKCS#8, respectively. Defaults to 'PKCS1' if not specified.",
								MarkdownDescription: "The private key cryptography standards (PKCS) encoding for this certificate's private key to be encoded in.  If provided, allowed values are 'PKCS1' and 'PKCS8' standing for PKCS#1 and PKCS#8, respectively. Defaults to 'PKCS1' if not specified.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("PKCS1", "PKCS8"),
								},
							},

							"rotation_policy": schema.StringAttribute{
								Description:         "RotationPolicy controls how private keys should be regenerated when a re-issuance is being processed.  If set to 'Never', a private key will only be generated if one does not already exist in the target 'spec.secretName'. If one does exists but it does not have the correct algorithm or size, a warning will be raised to await user intervention. If set to 'Always', a private key matching the specified requirements will be generated whenever a re-issuance occurs. Default is 'Never' for backward compatibility.",
								MarkdownDescription: "RotationPolicy controls how private keys should be regenerated when a re-issuance is being processed.  If set to 'Never', a private key will only be generated if one does not already exist in the target 'spec.secretName'. If one does exists but it does not have the correct algorithm or size, a warning will be raised to await user intervention. If set to 'Always', a private key matching the specified requirements will be generated whenever a re-issuance occurs. Default is 'Never' for backward compatibility.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Never", "Always"),
								},
							},

							"size": schema.Int64Attribute{
								Description:         "Size is the key bit size of the corresponding private key for this certificate.  If 'algorithm' is set to 'RSA', valid values are '2048', '4096' or '8192', and will default to '2048' if not specified. If 'algorithm' is set to 'ECDSA', valid values are '256', '384' or '521', and will default to '256' if not specified. If 'algorithm' is set to 'Ed25519', Size is ignored. No other values are allowed.",
								MarkdownDescription: "Size is the key bit size of the corresponding private key for this certificate.  If 'algorithm' is set to 'RSA', valid values are '2048', '4096' or '8192', and will default to '2048' if not specified. If 'algorithm' is set to 'ECDSA', valid values are '256', '384' or '521', and will default to '256' if not specified. If 'algorithm' is set to 'Ed25519', Size is ignored. No other values are allowed.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"renew_before": schema.StringAttribute{
						Description:         "How long before the currently issued certificate's expiry cert-manager should renew the certificate. For example, if a certificate is valid for 60 minutes, and 'renewBefore=10m', cert-manager will begin to attempt to renew the certificate 50 minutes after it was issued (i.e. when there are 10 minutes remaining until the certificate is no longer valid).  NOTE: The actual lifetime of the issued certificate is used to determine the renewal time. If an issuer returns a certificate with a different lifetime than the one requested, cert-manager will use the lifetime of the issued certificate.  If unset, this defaults to 1/3 of the issued certificate's lifetime. Minimum accepted value is 5 minutes. Value must be in units accepted by Go time.ParseDuration https://golang.org/pkg/time/#ParseDuration.",
						MarkdownDescription: "How long before the currently issued certificate's expiry cert-manager should renew the certificate. For example, if a certificate is valid for 60 minutes, and 'renewBefore=10m', cert-manager will begin to attempt to renew the certificate 50 minutes after it was issued (i.e. when there are 10 minutes remaining until the certificate is no longer valid).  NOTE: The actual lifetime of the issued certificate is used to determine the renewal time. If an issuer returns a certificate with a different lifetime than the one requested, cert-manager will use the lifetime of the issued certificate.  If unset, this defaults to 1/3 of the issued certificate's lifetime. Minimum accepted value is 5 minutes. Value must be in units accepted by Go time.ParseDuration https://golang.org/pkg/time/#ParseDuration.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"revision_history_limit": schema.Int64Attribute{
						Description:         "The maximum number of CertificateRequest revisions that are maintained in the Certificate's history. Each revision represents a single 'CertificateRequest' created by this Certificate, either when it was created, renewed, or Spec was changed. Revisions will be removed by oldest first if the number of revisions exceeds this number.  If set, revisionHistoryLimit must be a value of '1' or greater. If unset ('nil'), revisions will not be garbage collected. Default value is 'nil'.",
						MarkdownDescription: "The maximum number of CertificateRequest revisions that are maintained in the Certificate's history. Each revision represents a single 'CertificateRequest' created by this Certificate, either when it was created, renewed, or Spec was changed. Revisions will be removed by oldest first if the number of revisions exceeds this number.  If set, revisionHistoryLimit must be a value of '1' or greater. If unset ('nil'), revisions will not be garbage collected. Default value is 'nil'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"secret_name": schema.StringAttribute{
						Description:         "Name of the Secret resource that will be automatically created and managed by this Certificate resource. It will be populated with a private key and certificate, signed by the denoted issuer. The Secret resource lives in the same namespace as the Certificate resource.",
						MarkdownDescription: "Name of the Secret resource that will be automatically created and managed by this Certificate resource. It will be populated with a private key and certificate, signed by the denoted issuer. The Secret resource lives in the same namespace as the Certificate resource.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"secret_template": schema.SingleNestedAttribute{
						Description:         "Defines annotations and labels to be copied to the Certificate's Secret. Labels and annotations on the Secret will be changed as they appear on the SecretTemplate when added or removed. SecretTemplate annotations are added in conjunction with, and cannot overwrite, the base set of annotations cert-manager sets on the Certificate's Secret.",
						MarkdownDescription: "Defines annotations and labels to be copied to the Certificate's Secret. Labels and annotations on the Secret will be changed as they appear on the SecretTemplate when added or removed. SecretTemplate annotations are added in conjunction with, and cannot overwrite, the base set of annotations cert-manager sets on the Certificate's Secret.",
						Attributes: map[string]schema.Attribute{
							"annotations": schema.MapAttribute{
								Description:         "Annotations is a key value map to be copied to the target Kubernetes Secret.",
								MarkdownDescription: "Annotations is a key value map to be copied to the target Kubernetes Secret.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"labels": schema.MapAttribute{
								Description:         "Labels is a key value map to be copied to the target Kubernetes Secret.",
								MarkdownDescription: "Labels is a key value map to be copied to the target Kubernetes Secret.",
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

					"subject": schema.SingleNestedAttribute{
						Description:         "Requested set of X509 certificate subject attributes. More info: https://datatracker.ietf.org/doc/html/rfc5280#section-4.1.2.6  The common name attribute is specified separately in the 'commonName' field. Cannot be set if the 'literalSubject' field is set.",
						MarkdownDescription: "Requested set of X509 certificate subject attributes. More info: https://datatracker.ietf.org/doc/html/rfc5280#section-4.1.2.6  The common name attribute is specified separately in the 'commonName' field. Cannot be set if the 'literalSubject' field is set.",
						Attributes: map[string]schema.Attribute{
							"countries": schema.ListAttribute{
								Description:         "Countries to be used on the Certificate.",
								MarkdownDescription: "Countries to be used on the Certificate.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"localities": schema.ListAttribute{
								Description:         "Cities to be used on the Certificate.",
								MarkdownDescription: "Cities to be used on the Certificate.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"organizational_units": schema.ListAttribute{
								Description:         "Organizational Units to be used on the Certificate.",
								MarkdownDescription: "Organizational Units to be used on the Certificate.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"organizations": schema.ListAttribute{
								Description:         "Organizations to be used on the Certificate.",
								MarkdownDescription: "Organizations to be used on the Certificate.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"postal_codes": schema.ListAttribute{
								Description:         "Postal codes to be used on the Certificate.",
								MarkdownDescription: "Postal codes to be used on the Certificate.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"provinces": schema.ListAttribute{
								Description:         "State/Provinces to be used on the Certificate.",
								MarkdownDescription: "State/Provinces to be used on the Certificate.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"serial_number": schema.StringAttribute{
								Description:         "Serial number to be used on the Certificate.",
								MarkdownDescription: "Serial number to be used on the Certificate.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"street_addresses": schema.ListAttribute{
								Description:         "Street addresses to be used on the Certificate.",
								MarkdownDescription: "Street addresses to be used on the Certificate.",
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

					"uris": schema.ListAttribute{
						Description:         "Requested URI subject alternative names.",
						MarkdownDescription: "Requested URI subject alternative names.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"usages": schema.ListAttribute{
						Description:         "Requested key usages and extended key usages. These usages are used to set the 'usages' field on the created CertificateRequest resources. If 'encodeUsagesInRequest' is unset or set to 'true', the usages will additionally be encoded in the 'request' field which contains the CSR blob.  If unset, defaults to 'digital signature' and 'key encipherment'.",
						MarkdownDescription: "Requested key usages and extended key usages. These usages are used to set the 'usages' field on the created CertificateRequest resources. If 'encodeUsagesInRequest' is unset or set to 'true', the usages will additionally be encoded in the 'request' field which contains the CSR blob.  If unset, defaults to 'digital signature' and 'key encipherment'.",
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
	}
}

func (r *CertManagerIoCertificateV1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *dynamic.DynamicClient, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *CertManagerIoCertificateV1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_cert_manager_io_certificate_v1")

	var model CertManagerIoCertificateV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("cert-manager.io/v1")
	model.Kind = pointer.String("Certificate")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.Resource(k8sSchema.GroupVersionResource{Group: "cert-manager.io", Version: "v1", Resource: "Certificate"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while creating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse CertManagerIoCertificateV1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *CertManagerIoCertificateV1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_cert_manager_io_certificate_v1")

	var data CertManagerIoCertificateV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
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

	var readResponse CertManagerIoCertificateV1ResourceData
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

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *CertManagerIoCertificateV1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_cert_manager_io_certificate_v1")

	var model CertManagerIoCertificateV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("cert-manager.io/v1")
	model.Kind = pointer.String("Certificate")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.Resource(k8sSchema.GroupVersionResource{Group: "cert-manager.io", Version: "v1", Resource: "Certificate"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while updating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse CertManagerIoCertificateV1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *CertManagerIoCertificateV1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_cert_manager_io_certificate_v1")

	var data CertManagerIoCertificateV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "cert-manager.io", Version: "v1", Resource: "Certificate"}).
		Namespace(data.Metadata.Namespace).
		Delete(ctx, data.Metadata.Name, meta.DeleteOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to DELETE resource",
			"An unexpected error occurred while deleting the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"DELETE Error: "+err.Error(),
		)
		return
	}
}

func (r *CertManagerIoCertificateV1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	idParts := strings.Split(request.ID, "/")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'namespace/name' Got: '%q'", request.ID),
		)
		return
	}

	namespace := idParts[0]
	name := idParts[1]
	tflog.Trace(ctx, "parsed import ID", map[string]interface{}{
		"namespace": namespace,
		"name":      name,
	})
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("namespace"), namespace)...)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("name"), name)...)
}
