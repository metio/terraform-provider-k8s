/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package secrets_hashicorp_com_v1beta1

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
	_ datasource.DataSource = &SecretsHashicorpComVaultPkisecretV1Beta1Manifest{}
)

func NewSecretsHashicorpComVaultPkisecretV1Beta1Manifest() datasource.DataSource {
	return &SecretsHashicorpComVaultPkisecretV1Beta1Manifest{}
}

type SecretsHashicorpComVaultPkisecretV1Beta1Manifest struct{}

type SecretsHashicorpComVaultPkisecretV1Beta1ManifestData struct {
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
		AltNames    *[]string `tfsdk:"alt_names" json:"altNames,omitempty"`
		Clear       *bool     `tfsdk:"clear" json:"clear,omitempty"`
		CommonName  *string   `tfsdk:"common_name" json:"commonName,omitempty"`
		Destination *struct {
			Annotations    *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			Create         *bool              `tfsdk:"create" json:"create,omitempty"`
			Labels         *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Name           *string            `tfsdk:"name" json:"name,omitempty"`
			Overwrite      *bool              `tfsdk:"overwrite" json:"overwrite,omitempty"`
			Transformation *struct {
				ExcludeRaw *bool     `tfsdk:"exclude_raw" json:"excludeRaw,omitempty"`
				Excludes   *[]string `tfsdk:"excludes" json:"excludes,omitempty"`
				Includes   *[]string `tfsdk:"includes" json:"includes,omitempty"`
				Templates  *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
					Text *string `tfsdk:"text" json:"text,omitempty"`
				} `tfsdk:"templates" json:"templates,omitempty"`
				TransformationRefs *[]struct {
					IgnoreExcludes *bool   `tfsdk:"ignore_excludes" json:"ignoreExcludes,omitempty"`
					IgnoreIncludes *bool   `tfsdk:"ignore_includes" json:"ignoreIncludes,omitempty"`
					Name           *string `tfsdk:"name" json:"name,omitempty"`
					Namespace      *string `tfsdk:"namespace" json:"namespace,omitempty"`
					TemplateRefs   *[]struct {
						KeyOverride *string `tfsdk:"key_override" json:"keyOverride,omitempty"`
						Name        *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"template_refs" json:"templateRefs,omitempty"`
				} `tfsdk:"transformation_refs" json:"transformationRefs,omitempty"`
			} `tfsdk:"transformation" json:"transformation,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"destination" json:"destination,omitempty"`
		ExcludeCNFromSans     *bool     `tfsdk:"exclude_cn_from_sans" json:"excludeCNFromSans,omitempty"`
		ExpiryOffset          *string   `tfsdk:"expiry_offset" json:"expiryOffset,omitempty"`
		Format                *string   `tfsdk:"format" json:"format,omitempty"`
		IpSans                *[]string `tfsdk:"ip_sans" json:"ipSans,omitempty"`
		IssuerRef             *string   `tfsdk:"issuer_ref" json:"issuerRef,omitempty"`
		Mount                 *string   `tfsdk:"mount" json:"mount,omitempty"`
		Namespace             *string   `tfsdk:"namespace" json:"namespace,omitempty"`
		NotAfter              *string   `tfsdk:"not_after" json:"notAfter,omitempty"`
		OtherSans             *[]string `tfsdk:"other_sans" json:"otherSans,omitempty"`
		PrivateKeyFormat      *string   `tfsdk:"private_key_format" json:"privateKeyFormat,omitempty"`
		Revoke                *bool     `tfsdk:"revoke" json:"revoke,omitempty"`
		Role                  *string   `tfsdk:"role" json:"role,omitempty"`
		RolloutRestartTargets *[]struct {
			Kind *string `tfsdk:"kind" json:"kind,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"rollout_restart_targets" json:"rolloutRestartTargets,omitempty"`
		Ttl          *string   `tfsdk:"ttl" json:"ttl,omitempty"`
		UriSans      *[]string `tfsdk:"uri_sans" json:"uriSans,omitempty"`
		UserIDs      *[]string `tfsdk:"user_i_ds" json:"userIDs,omitempty"`
		VaultAuthRef *string   `tfsdk:"vault_auth_ref" json:"vaultAuthRef,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SecretsHashicorpComVaultPkisecretV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_secrets_hashicorp_com_vault_pki_secret_v1beta1_manifest"
}

func (r *SecretsHashicorpComVaultPkisecretV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "VaultPKISecret is the Schema for the vaultpkisecrets API",
		MarkdownDescription: "VaultPKISecret is the Schema for the vaultpkisecrets API",
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
				Description:         "VaultPKISecretSpec defines the desired state of VaultPKISecret",
				MarkdownDescription: "VaultPKISecretSpec defines the desired state of VaultPKISecret",
				Attributes: map[string]schema.Attribute{
					"alt_names": schema.ListAttribute{
						Description:         "AltNames to include in the request May contain both DNS names and email addresses.",
						MarkdownDescription: "AltNames to include in the request May contain both DNS names and email addresses.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"clear": schema.BoolAttribute{
						Description:         "Clear the Kubernetes secret when the resource is deleted.",
						MarkdownDescription: "Clear the Kubernetes secret when the resource is deleted.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"common_name": schema.StringAttribute{
						Description:         "CommonName to include in the request.",
						MarkdownDescription: "CommonName to include in the request.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"destination": schema.SingleNestedAttribute{
						Description:         "Destination provides configuration necessary for syncing the Vault secret to Kubernetes. If the type is set to 'kubernetes.io/tls', 'tls.key' will be set to the 'private_key' response from Vault, and 'tls.crt' will be set to 'certificate' + 'ca_chain' from the Vault response ('issuing_ca' is used when 'ca_chain' is empty). The 'remove_roots_from_chain=true' option is used with Vault to exclude the root CA from the Vault response.",
						MarkdownDescription: "Destination provides configuration necessary for syncing the Vault secret to Kubernetes. If the type is set to 'kubernetes.io/tls', 'tls.key' will be set to the 'private_key' response from Vault, and 'tls.crt' will be set to 'certificate' + 'ca_chain' from the Vault response ('issuing_ca' is used when 'ca_chain' is empty). The 'remove_roots_from_chain=true' option is used with Vault to exclude the root CA from the Vault response.",
						Attributes: map[string]schema.Attribute{
							"annotations": schema.MapAttribute{
								Description:         "Annotations to apply to the Secret. Requires Create to be set to true.",
								MarkdownDescription: "Annotations to apply to the Secret. Requires Create to be set to true.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"create": schema.BoolAttribute{
								Description:         "Create the destination Secret. If the Secret already exists this should be set to false.",
								MarkdownDescription: "Create the destination Secret. If the Secret already exists this should be set to false.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"labels": schema.MapAttribute{
								Description:         "Labels to apply to the Secret. Requires Create to be set to true.",
								MarkdownDescription: "Labels to apply to the Secret. Requires Create to be set to true.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the Secret",
								MarkdownDescription: "Name of the Secret",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"overwrite": schema.BoolAttribute{
								Description:         "Overwrite the destination Secret if it exists and Create is true. This is useful when migrating to VSO from a previous secret deployment strategy.",
								MarkdownDescription: "Overwrite the destination Secret if it exists and Create is true. This is useful when migrating to VSO from a previous secret deployment strategy.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"transformation": schema.SingleNestedAttribute{
								Description:         "Transformation provides configuration for transforming the secret data before it is stored in the Destination.",
								MarkdownDescription: "Transformation provides configuration for transforming the secret data before it is stored in the Destination.",
								Attributes: map[string]schema.Attribute{
									"exclude_raw": schema.BoolAttribute{
										Description:         "ExcludeRaw data from the destination Secret. Exclusion policy can be set globally by including 'exclude-raw' in the '--global-transformation-options' command line flag. If set, the command line flag always takes precedence over this configuration.",
										MarkdownDescription: "ExcludeRaw data from the destination Secret. Exclusion policy can be set globally by including 'exclude-raw' in the '--global-transformation-options' command line flag. If set, the command line flag always takes precedence over this configuration.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"excludes": schema.ListAttribute{
										Description:         "Excludes contains regex patterns used to filter top-level source secret data fields for exclusion from the final K8s Secret data. These pattern filters are never applied to templated fields as defined in Templates. They are always applied before any inclusion patterns. To exclude all source secret data fields, you can configure the single pattern '.*'.",
										MarkdownDescription: "Excludes contains regex patterns used to filter top-level source secret data fields for exclusion from the final K8s Secret data. These pattern filters are never applied to templated fields as defined in Templates. They are always applied before any inclusion patterns. To exclude all source secret data fields, you can configure the single pattern '.*'.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"includes": schema.ListAttribute{
										Description:         "Includes contains regex patterns used to filter top-level source secret data fields for inclusion in the final K8s Secret data. These pattern filters are never applied to templated fields as defined in Templates. They are always applied last.",
										MarkdownDescription: "Includes contains regex patterns used to filter top-level source secret data fields for inclusion in the final K8s Secret data. These pattern filters are never applied to templated fields as defined in Templates. They are always applied last.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"templates": schema.SingleNestedAttribute{
										Description:         "Templates maps a template name to its Template. Templates are always included in the rendered K8s Secret, and take precedence over templates defined in a SecretTransformation.",
										MarkdownDescription: "Templates maps a template name to its Template. Templates are always included in the rendered K8s Secret, and take precedence over templates defined in a SecretTransformation.",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the Template",
												MarkdownDescription: "Name of the Template",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"text": schema.StringAttribute{
												Description:         "Text contains the Go text template format. The template references attributes from the data structure of the source secret. Refer to https://pkg.go.dev/text/template for more information.",
												MarkdownDescription: "Text contains the Go text template format. The template references attributes from the data structure of the source secret. Refer to https://pkg.go.dev/text/template for more information.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"transformation_refs": schema.ListNestedAttribute{
										Description:         "TransformationRefs contain references to template configuration from SecretTransformation.",
										MarkdownDescription: "TransformationRefs contain references to template configuration from SecretTransformation.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"ignore_excludes": schema.BoolAttribute{
													Description:         "IgnoreExcludes controls whether to use the SecretTransformation's Excludes data key filters.",
													MarkdownDescription: "IgnoreExcludes controls whether to use the SecretTransformation's Excludes data key filters.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"ignore_includes": schema.BoolAttribute{
													Description:         "IgnoreIncludes controls whether to use the SecretTransformation's Includes data key filters.",
													MarkdownDescription: "IgnoreIncludes controls whether to use the SecretTransformation's Includes data key filters.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the SecretTransformation resource.",
													MarkdownDescription: "Name of the SecretTransformation resource.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"namespace": schema.StringAttribute{
													Description:         "Namespace of the SecretTransformation resource.",
													MarkdownDescription: "Namespace of the SecretTransformation resource.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"template_refs": schema.ListNestedAttribute{
													Description:         "TemplateRefs map to a Template found in this TransformationRef. If empty, then all templates from the SecretTransformation will be rendered to the K8s Secret.",
													MarkdownDescription: "TemplateRefs map to a Template found in this TransformationRef. If empty, then all templates from the SecretTransformation will be rendered to the K8s Secret.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"key_override": schema.StringAttribute{
																Description:         "KeyOverride to the rendered template in the Destination secret. If Key is empty, then the Key from reference spec will be used. Set this to override the Key set from the reference spec.",
																MarkdownDescription: "KeyOverride to the rendered template in the Destination secret. If Key is empty, then the Key from reference spec will be used. Set this to override the Key set from the reference spec.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "Name of the Template in SecretTransformationSpec.Templates. the rendered secret data.",
																MarkdownDescription: "Name of the Template in SecretTransformationSpec.Templates. the rendered secret data.",
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

							"type": schema.StringAttribute{
								Description:         "Type of Kubernetes Secret. Requires Create to be set to true. Defaults to Opaque.",
								MarkdownDescription: "Type of Kubernetes Secret. Requires Create to be set to true. Defaults to Opaque.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"exclude_cn_from_sans": schema.BoolAttribute{
						Description:         "ExcludeCNFromSans from DNS or Email Subject Alternate Names. Default: false",
						MarkdownDescription: "ExcludeCNFromSans from DNS or Email Subject Alternate Names. Default: false",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"expiry_offset": schema.StringAttribute{
						Description:         "ExpiryOffset to use for computing when the certificate should be renewed. The rotation time will be difference between the expiration and the offset. Should be in duration notation e.g. 30s, 120s, etc.",
						MarkdownDescription: "ExpiryOffset to use for computing when the certificate should be renewed. The rotation time will be difference between the expiration and the offset. Should be in duration notation e.g. 30s, 120s, etc.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(s|m|h))$`), ""),
						},
					},

					"format": schema.StringAttribute{
						Description:         "Format for the certificate. Choices: 'pem', 'der', 'pem_bundle'. If 'pem_bundle', any private key and issuing cert will be appended to the certificate pem. If 'der', the value will be base64 encoded. Default: pem",
						MarkdownDescription: "Format for the certificate. Choices: 'pem', 'der', 'pem_bundle'. If 'pem_bundle', any private key and issuing cert will be appended to the certificate pem. If 'der', the value will be base64 encoded. Default: pem",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ip_sans": schema.ListAttribute{
						Description:         "IPSans to include in the request.",
						MarkdownDescription: "IPSans to include in the request.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"issuer_ref": schema.StringAttribute{
						Description:         "IssuerRef reference to an existing PKI issuer, either by Vault-generated identifier, the literal string default to refer to the currently configured default issuer, or the name assigned to an issuer. This parameter is part of the request URL.",
						MarkdownDescription: "IssuerRef reference to an existing PKI issuer, either by Vault-generated identifier, the literal string default to refer to the currently configured default issuer, or the name assigned to an issuer. This parameter is part of the request URL.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"mount": schema.StringAttribute{
						Description:         "Mount for the secret in Vault",
						MarkdownDescription: "Mount for the secret in Vault",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"namespace": schema.StringAttribute{
						Description:         "Namespace to get the secret from in Vault",
						MarkdownDescription: "Namespace to get the secret from in Vault",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"not_after": schema.StringAttribute{
						Description:         "NotAfter field of the certificate with specified date value. The value format should be given in UTC format YYYY-MM-ddTHH:MM:SSZ",
						MarkdownDescription: "NotAfter field of the certificate with specified date value. The value format should be given in UTC format YYYY-MM-ddTHH:MM:SSZ",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"other_sans": schema.ListAttribute{
						Description:         "Requested other SANs, in an array with the format oid;type:value for each entry.",
						MarkdownDescription: "Requested other SANs, in an array with the format oid;type:value for each entry.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"private_key_format": schema.StringAttribute{
						Description:         "PrivateKeyFormat, generally the default will be controlled by the Format parameter as either base64-encoded DER or PEM-encoded DER. However, this can be set to 'pkcs8' to have the returned private key contain base64-encoded pkcs8 or PEM-encoded pkcs8 instead. Default: der",
						MarkdownDescription: "PrivateKeyFormat, generally the default will be controlled by the Format parameter as either base64-encoded DER or PEM-encoded DER. However, this can be set to 'pkcs8' to have the returned private key contain base64-encoded pkcs8 or PEM-encoded pkcs8 instead. Default: der",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"revoke": schema.BoolAttribute{
						Description:         "Revoke the certificate when the resource is deleted.",
						MarkdownDescription: "Revoke the certificate when the resource is deleted.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"role": schema.StringAttribute{
						Description:         "Role in Vault to use when issuing TLS certificates.",
						MarkdownDescription: "Role in Vault to use when issuing TLS certificates.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"rollout_restart_targets": schema.ListNestedAttribute{
						Description:         "RolloutRestartTargets should be configured whenever the application(s) consuming the Vault secret does not support dynamically reloading a rotated secret. In that case one, or more RolloutRestartTarget(s) can be configured here. The Operator will trigger a 'rollout-restart' for each target whenever the Vault secret changes between reconciliation events. See RolloutRestartTarget for more details.",
						MarkdownDescription: "RolloutRestartTargets should be configured whenever the application(s) consuming the Vault secret does not support dynamically reloading a rotated secret. In that case one, or more RolloutRestartTarget(s) can be configured here. The Operator will trigger a 'rollout-restart' for each target whenever the Vault secret changes between reconciliation events. See RolloutRestartTarget for more details.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"kind": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("Deployment", "DaemonSet", "StatefulSet"),
									},
								},

								"name": schema.StringAttribute{
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

					"ttl": schema.StringAttribute{
						Description:         "TTL for the certificate; sets the expiration date. If not specified the Vault role's default, backend default, or system default TTL is used, in that order. Cannot be larger than the mount's max TTL. Note: this only has an effect when generating a CA cert or signing a CA cert, not when generating a CSR for an intermediate CA. Should be in duration notation e.g. 120s, 2h, etc.",
						MarkdownDescription: "TTL for the certificate; sets the expiration date. If not specified the Vault role's default, backend default, or system default TTL is used, in that order. Cannot be larger than the mount's max TTL. Note: this only has an effect when generating a CA cert or signing a CA cert, not when generating a CSR for an intermediate CA. Should be in duration notation e.g. 120s, 2h, etc.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(s|m|h))$`), ""),
						},
					},

					"uri_sans": schema.ListAttribute{
						Description:         "The requested URI SANs.",
						MarkdownDescription: "The requested URI SANs.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"user_i_ds": schema.ListAttribute{
						Description:         "User ID (OID 0.9.2342.19200300.100.1.1) Subject values to be placed on the signed certificate.",
						MarkdownDescription: "User ID (OID 0.9.2342.19200300.100.1.1) Subject values to be placed on the signed certificate.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"vault_auth_ref": schema.StringAttribute{
						Description:         "VaultAuthRef to the VaultAuth resource, can be prefixed with a namespace, eg: 'namespaceA/vaultAuthRefB'. If no namespace prefix is provided it will default to namespace of the VaultAuth CR. If no value is specified for VaultAuthRef the Operator will default to the 'default' VaultAuth, configured in the operator's namespace.",
						MarkdownDescription: "VaultAuthRef to the VaultAuth resource, can be prefixed with a namespace, eg: 'namespaceA/vaultAuthRefB'. If no namespace prefix is provided it will default to namespace of the VaultAuth CR. If no value is specified for VaultAuthRef the Operator will default to the 'default' VaultAuth, configured in the operator's namespace.",
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

func (r *SecretsHashicorpComVaultPkisecretV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_secrets_hashicorp_com_vault_pki_secret_v1beta1_manifest")

	var model SecretsHashicorpComVaultPkisecretV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("secrets.hashicorp.com/v1beta1")
	model.Kind = pointer.String("VaultPKISecret")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
