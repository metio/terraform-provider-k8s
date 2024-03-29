/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package config_openshift_io_v1

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
	_ datasource.DataSource = &ConfigOpenshiftIoApiserverV1Manifest{}
)

func NewConfigOpenshiftIoApiserverV1Manifest() datasource.DataSource {
	return &ConfigOpenshiftIoApiserverV1Manifest{}
}

type ConfigOpenshiftIoApiserverV1Manifest struct{}

type ConfigOpenshiftIoApiserverV1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		AdditionalCORSAllowedOrigins *[]string `tfsdk:"additional_cors_allowed_origins" json:"additionalCORSAllowedOrigins,omitempty"`
		Audit                        *struct {
			CustomRules *[]struct {
				Group   *string `tfsdk:"group" json:"group,omitempty"`
				Profile *string `tfsdk:"profile" json:"profile,omitempty"`
			} `tfsdk:"custom_rules" json:"customRules,omitempty"`
			Profile *string `tfsdk:"profile" json:"profile,omitempty"`
		} `tfsdk:"audit" json:"audit,omitempty"`
		ClientCA *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"client_ca" json:"clientCA,omitempty"`
		Encryption *struct {
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"encryption" json:"encryption,omitempty"`
		ServingCerts *struct {
			NamedCertificates *[]struct {
				Names              *[]string `tfsdk:"names" json:"names,omitempty"`
				ServingCertificate *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"serving_certificate" json:"servingCertificate,omitempty"`
			} `tfsdk:"named_certificates" json:"namedCertificates,omitempty"`
		} `tfsdk:"serving_certs" json:"servingCerts,omitempty"`
		TlsSecurityProfile *struct {
			Custom *struct {
				Ciphers       *[]string `tfsdk:"ciphers" json:"ciphers,omitempty"`
				MinTLSVersion *string   `tfsdk:"min_tls_version" json:"minTLSVersion,omitempty"`
			} `tfsdk:"custom" json:"custom,omitempty"`
			Intermediate *map[string]string `tfsdk:"intermediate" json:"intermediate,omitempty"`
			Modern       *map[string]string `tfsdk:"modern" json:"modern,omitempty"`
			Old          *map[string]string `tfsdk:"old" json:"old,omitempty"`
			Type         *string            `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"tls_security_profile" json:"tlsSecurityProfile,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ConfigOpenshiftIoApiserverV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_config_openshift_io_api_server_v1_manifest"
}

func (r *ConfigOpenshiftIoApiserverV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "APIServer holds configuration (like serving certificates, client CA and CORS domains) shared by all API servers in the system, among them especially kube-apiserver and openshift-apiserver. The canonical name of an instance is 'cluster'.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		MarkdownDescription: "APIServer holds configuration (like serving certificates, client CA and CORS domains) shared by all API servers in the system, among them especially kube-apiserver and openshift-apiserver. The canonical name of an instance is 'cluster'.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

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
				Description:         "spec holds user settable values for configuration",
				MarkdownDescription: "spec holds user settable values for configuration",
				Attributes: map[string]schema.Attribute{
					"additional_cors_allowed_origins": schema.ListAttribute{
						Description:         "additionalCORSAllowedOrigins lists additional, user-defined regular expressions describing hosts for which the API server allows access using the CORS headers. This may be needed to access the API and the integrated OAuth server from JavaScript applications. The values are regular expressions that correspond to the Golang regular expression language.",
						MarkdownDescription: "additionalCORSAllowedOrigins lists additional, user-defined regular expressions describing hosts for which the API server allows access using the CORS headers. This may be needed to access the API and the integrated OAuth server from JavaScript applications. The values are regular expressions that correspond to the Golang regular expression language.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"audit": schema.SingleNestedAttribute{
						Description:         "audit specifies the settings for audit configuration to be applied to all OpenShift-provided API servers in the cluster.",
						MarkdownDescription: "audit specifies the settings for audit configuration to be applied to all OpenShift-provided API servers in the cluster.",
						Attributes: map[string]schema.Attribute{
							"custom_rules": schema.ListNestedAttribute{
								Description:         "customRules specify profiles per group. These profile take precedence over the top-level profile field if they apply. They are evaluation from top to bottom and the first one that matches, applies.",
								MarkdownDescription: "customRules specify profiles per group. These profile take precedence over the top-level profile field if they apply. They are evaluation from top to bottom and the first one that matches, applies.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"group": schema.StringAttribute{
											Description:         "group is a name of group a request user must be member of in order to this profile to apply.",
											MarkdownDescription: "group is a name of group a request user must be member of in order to this profile to apply.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
											},
										},

										"profile": schema.StringAttribute{
											Description:         "profile specifies the name of the desired audit policy configuration to be deployed to all OpenShift-provided API servers in the cluster.  The following profiles are provided: - Default: the existing default policy. - WriteRequestBodies: like 'Default', but logs request and response HTTP payloads for write requests (create, update, patch). - AllRequestBodies: like 'WriteRequestBodies', but also logs request and response HTTP payloads for read requests (get, list). - None: no requests are logged at all, not even oauthaccesstokens and oauthauthorizetokens.  If unset, the 'Default' profile is used as the default.",
											MarkdownDescription: "profile specifies the name of the desired audit policy configuration to be deployed to all OpenShift-provided API servers in the cluster.  The following profiles are provided: - Default: the existing default policy. - WriteRequestBodies: like 'Default', but logs request and response HTTP payloads for write requests (create, update, patch). - AllRequestBodies: like 'WriteRequestBodies', but also logs request and response HTTP payloads for read requests (get, list). - None: no requests are logged at all, not even oauthaccesstokens and oauthauthorizetokens.  If unset, the 'Default' profile is used as the default.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Default", "WriteRequestBodies", "AllRequestBodies", "None"),
											},
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"profile": schema.StringAttribute{
								Description:         "profile specifies the name of the desired top-level audit profile to be applied to all requests sent to any of the OpenShift-provided API servers in the cluster (kube-apiserver, openshift-apiserver and oauth-apiserver), with the exception of those requests that match one or more of the customRules.  The following profiles are provided: - Default: default policy which means MetaData level logging with the exception of events (not logged at all), oauthaccesstokens and oauthauthorizetokens (both logged at RequestBody level). - WriteRequestBodies: like 'Default', but logs request and response HTTP payloads for write requests (create, update, patch). - AllRequestBodies: like 'WriteRequestBodies', but also logs request and response HTTP payloads for read requests (get, list). - None: no requests are logged at all, not even oauthaccesstokens and oauthauthorizetokens.  Warning: It is not recommended to disable audit logging by using the 'None' profile unless you are fully aware of the risks of not logging data that can be beneficial when troubleshooting issues. If you disable audit logging and a support situation arises, you might need to enable audit logging and reproduce the issue in order to troubleshoot properly.  If unset, the 'Default' profile is used as the default.",
								MarkdownDescription: "profile specifies the name of the desired top-level audit profile to be applied to all requests sent to any of the OpenShift-provided API servers in the cluster (kube-apiserver, openshift-apiserver and oauth-apiserver), with the exception of those requests that match one or more of the customRules.  The following profiles are provided: - Default: default policy which means MetaData level logging with the exception of events (not logged at all), oauthaccesstokens and oauthauthorizetokens (both logged at RequestBody level). - WriteRequestBodies: like 'Default', but logs request and response HTTP payloads for write requests (create, update, patch). - AllRequestBodies: like 'WriteRequestBodies', but also logs request and response HTTP payloads for read requests (get, list). - None: no requests are logged at all, not even oauthaccesstokens and oauthauthorizetokens.  Warning: It is not recommended to disable audit logging by using the 'None' profile unless you are fully aware of the risks of not logging data that can be beneficial when troubleshooting issues. If you disable audit logging and a support situation arises, you might need to enable audit logging and reproduce the issue in order to troubleshoot properly.  If unset, the 'Default' profile is used as the default.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Default", "WriteRequestBodies", "AllRequestBodies", "None"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"client_ca": schema.SingleNestedAttribute{
						Description:         "clientCA references a ConfigMap containing a certificate bundle for the signers that will be recognized for incoming client certificates in addition to the operator managed signers. If this is empty, then only operator managed signers are valid. You usually only have to set this if you have your own PKI you wish to honor client certificates from. The ConfigMap must exist in the openshift-config namespace and contain the following required fields: - ConfigMap.Data['ca-bundle.crt'] - CA bundle.",
						MarkdownDescription: "clientCA references a ConfigMap containing a certificate bundle for the signers that will be recognized for incoming client certificates in addition to the operator managed signers. If this is empty, then only operator managed signers are valid. You usually only have to set this if you have your own PKI you wish to honor client certificates from. The ConfigMap must exist in the openshift-config namespace and contain the following required fields: - ConfigMap.Data['ca-bundle.crt'] - CA bundle.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "name is the metadata.name of the referenced config map",
								MarkdownDescription: "name is the metadata.name of the referenced config map",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"encryption": schema.SingleNestedAttribute{
						Description:         "encryption allows the configuration of encryption of resources at the datastore layer.",
						MarkdownDescription: "encryption allows the configuration of encryption of resources at the datastore layer.",
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Description:         "type defines what encryption type should be used to encrypt resources at the datastore layer. When this field is unset (i.e. when it is set to the empty string), identity is implied. The behavior of unset can and will change over time.  Even if encryption is enabled by default, the meaning of unset may change to a different encryption type based on changes in best practices.  When encryption is enabled, all sensitive resources shipped with the platform are encrypted. This list of sensitive resources can and will change over time.  The current authoritative list is:  1. secrets 2. configmaps 3. routes.route.openshift.io 4. oauthaccesstokens.oauth.openshift.io 5. oauthauthorizetokens.oauth.openshift.io",
								MarkdownDescription: "type defines what encryption type should be used to encrypt resources at the datastore layer. When this field is unset (i.e. when it is set to the empty string), identity is implied. The behavior of unset can and will change over time.  Even if encryption is enabled by default, the meaning of unset may change to a different encryption type based on changes in best practices.  When encryption is enabled, all sensitive resources shipped with the platform are encrypted. This list of sensitive resources can and will change over time.  The current authoritative list is:  1. secrets 2. configmaps 3. routes.route.openshift.io 4. oauthaccesstokens.oauth.openshift.io 5. oauthauthorizetokens.oauth.openshift.io",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("", "identity", "aescbc", "aesgcm"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"serving_certs": schema.SingleNestedAttribute{
						Description:         "servingCert is the TLS cert info for serving secure traffic. If not specified, operator managed certificates will be used for serving secure traffic.",
						MarkdownDescription: "servingCert is the TLS cert info for serving secure traffic. If not specified, operator managed certificates will be used for serving secure traffic.",
						Attributes: map[string]schema.Attribute{
							"named_certificates": schema.ListNestedAttribute{
								Description:         "namedCertificates references secrets containing the TLS cert info for serving secure traffic to specific hostnames. If no named certificates are provided, or no named certificates match the server name as understood by a client, the defaultServingCertificate will be used.",
								MarkdownDescription: "namedCertificates references secrets containing the TLS cert info for serving secure traffic to specific hostnames. If no named certificates are provided, or no named certificates match the server name as understood by a client, the defaultServingCertificate will be used.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"names": schema.ListAttribute{
											Description:         "names is a optional list of explicit DNS names (leading wildcards allowed) that should use this certificate to serve secure traffic. If no names are provided, the implicit names will be extracted from the certificates. Exact names trump over wildcard names. Explicit names defined here trump over extracted implicit names.",
											MarkdownDescription: "names is a optional list of explicit DNS names (leading wildcards allowed) that should use this certificate to serve secure traffic. If no names are provided, the implicit names will be extracted from the certificates. Exact names trump over wildcard names. Explicit names defined here trump over extracted implicit names.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"serving_certificate": schema.SingleNestedAttribute{
											Description:         "servingCertificate references a kubernetes.io/tls type secret containing the TLS cert info for serving secure traffic. The secret must exist in the openshift-config namespace and contain the following required fields: - Secret.Data['tls.key'] - TLS private key. - Secret.Data['tls.crt'] - TLS certificate.",
											MarkdownDescription: "servingCertificate references a kubernetes.io/tls type secret containing the TLS cert info for serving secure traffic. The secret must exist in the openshift-config namespace and contain the following required fields: - Secret.Data['tls.key'] - TLS private key. - Secret.Data['tls.crt'] - TLS certificate.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "name is the metadata.name of the referenced secret",
													MarkdownDescription: "name is the metadata.name of the referenced secret",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"tls_security_profile": schema.SingleNestedAttribute{
						Description:         "tlsSecurityProfile specifies settings for TLS connections for externally exposed servers.  If unset, a default (which may change between releases) is chosen. Note that only Old, Intermediate and Custom profiles are currently supported, and the maximum available minTLSVersion is VersionTLS12.",
						MarkdownDescription: "tlsSecurityProfile specifies settings for TLS connections for externally exposed servers.  If unset, a default (which may change between releases) is chosen. Note that only Old, Intermediate and Custom profiles are currently supported, and the maximum available minTLSVersion is VersionTLS12.",
						Attributes: map[string]schema.Attribute{
							"custom": schema.SingleNestedAttribute{
								Description:         "custom is a user-defined TLS security profile. Be extremely careful using a custom profile as invalid configurations can be catastrophic. An example custom profile looks like this:  ciphers:  - ECDHE-ECDSA-CHACHA20-POLY1305  - ECDHE-RSA-CHACHA20-POLY1305  - ECDHE-RSA-AES128-GCM-SHA256  - ECDHE-ECDSA-AES128-GCM-SHA256  minTLSVersion: VersionTLS11",
								MarkdownDescription: "custom is a user-defined TLS security profile. Be extremely careful using a custom profile as invalid configurations can be catastrophic. An example custom profile looks like this:  ciphers:  - ECDHE-ECDSA-CHACHA20-POLY1305  - ECDHE-RSA-CHACHA20-POLY1305  - ECDHE-RSA-AES128-GCM-SHA256  - ECDHE-ECDSA-AES128-GCM-SHA256  minTLSVersion: VersionTLS11",
								Attributes: map[string]schema.Attribute{
									"ciphers": schema.ListAttribute{
										Description:         "ciphers is used to specify the cipher algorithms that are negotiated during the TLS handshake.  Operators may remove entries their operands do not support.  For example, to use DES-CBC3-SHA  (yaml):  ciphers: - DES-CBC3-SHA",
										MarkdownDescription: "ciphers is used to specify the cipher algorithms that are negotiated during the TLS handshake.  Operators may remove entries their operands do not support.  For example, to use DES-CBC3-SHA  (yaml):  ciphers: - DES-CBC3-SHA",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"min_tls_version": schema.StringAttribute{
										Description:         "minTLSVersion is used to specify the minimal version of the TLS protocol that is negotiated during the TLS handshake. For example, to use TLS versions 1.1, 1.2 and 1.3 (yaml):  minTLSVersion: VersionTLS11  NOTE: currently the highest minTLSVersion allowed is VersionTLS12",
										MarkdownDescription: "minTLSVersion is used to specify the minimal version of the TLS protocol that is negotiated during the TLS handshake. For example, to use TLS versions 1.1, 1.2 and 1.3 (yaml):  minTLSVersion: VersionTLS11  NOTE: currently the highest minTLSVersion allowed is VersionTLS12",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("VersionTLS10", "VersionTLS11", "VersionTLS12", "VersionTLS13"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"intermediate": schema.MapAttribute{
								Description:         "intermediate is a TLS security profile based on:  https://wiki.mozilla.org/Security/Server_Side_TLS#Intermediate_compatibility_.28recommended.29  and looks like this (yaml):  ciphers:  - TLS_AES_128_GCM_SHA256  - TLS_AES_256_GCM_SHA384  - TLS_CHACHA20_POLY1305_SHA256  - ECDHE-ECDSA-AES128-GCM-SHA256  - ECDHE-RSA-AES128-GCM-SHA256  - ECDHE-ECDSA-AES256-GCM-SHA384  - ECDHE-RSA-AES256-GCM-SHA384  - ECDHE-ECDSA-CHACHA20-POLY1305  - ECDHE-RSA-CHACHA20-POLY1305  - DHE-RSA-AES128-GCM-SHA256  - DHE-RSA-AES256-GCM-SHA384  minTLSVersion: VersionTLS12",
								MarkdownDescription: "intermediate is a TLS security profile based on:  https://wiki.mozilla.org/Security/Server_Side_TLS#Intermediate_compatibility_.28recommended.29  and looks like this (yaml):  ciphers:  - TLS_AES_128_GCM_SHA256  - TLS_AES_256_GCM_SHA384  - TLS_CHACHA20_POLY1305_SHA256  - ECDHE-ECDSA-AES128-GCM-SHA256  - ECDHE-RSA-AES128-GCM-SHA256  - ECDHE-ECDSA-AES256-GCM-SHA384  - ECDHE-RSA-AES256-GCM-SHA384  - ECDHE-ECDSA-CHACHA20-POLY1305  - ECDHE-RSA-CHACHA20-POLY1305  - DHE-RSA-AES128-GCM-SHA256  - DHE-RSA-AES256-GCM-SHA384  minTLSVersion: VersionTLS12",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"modern": schema.MapAttribute{
								Description:         "modern is a TLS security profile based on:  https://wiki.mozilla.org/Security/Server_Side_TLS#Modern_compatibility  and looks like this (yaml):  ciphers:  - TLS_AES_128_GCM_SHA256  - TLS_AES_256_GCM_SHA384  - TLS_CHACHA20_POLY1305_SHA256  minTLSVersion: VersionTLS13",
								MarkdownDescription: "modern is a TLS security profile based on:  https://wiki.mozilla.org/Security/Server_Side_TLS#Modern_compatibility  and looks like this (yaml):  ciphers:  - TLS_AES_128_GCM_SHA256  - TLS_AES_256_GCM_SHA384  - TLS_CHACHA20_POLY1305_SHA256  minTLSVersion: VersionTLS13",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"old": schema.MapAttribute{
								Description:         "old is a TLS security profile based on:  https://wiki.mozilla.org/Security/Server_Side_TLS#Old_backward_compatibility  and looks like this (yaml):  ciphers:  - TLS_AES_128_GCM_SHA256  - TLS_AES_256_GCM_SHA384  - TLS_CHACHA20_POLY1305_SHA256  - ECDHE-ECDSA-AES128-GCM-SHA256  - ECDHE-RSA-AES128-GCM-SHA256  - ECDHE-ECDSA-AES256-GCM-SHA384  - ECDHE-RSA-AES256-GCM-SHA384  - ECDHE-ECDSA-CHACHA20-POLY1305  - ECDHE-RSA-CHACHA20-POLY1305  - DHE-RSA-AES128-GCM-SHA256  - DHE-RSA-AES256-GCM-SHA384  - DHE-RSA-CHACHA20-POLY1305  - ECDHE-ECDSA-AES128-SHA256  - ECDHE-RSA-AES128-SHA256  - ECDHE-ECDSA-AES128-SHA  - ECDHE-RSA-AES128-SHA  - ECDHE-ECDSA-AES256-SHA384  - ECDHE-RSA-AES256-SHA384  - ECDHE-ECDSA-AES256-SHA  - ECDHE-RSA-AES256-SHA  - DHE-RSA-AES128-SHA256  - DHE-RSA-AES256-SHA256  - AES128-GCM-SHA256  - AES256-GCM-SHA384  - AES128-SHA256  - AES256-SHA256  - AES128-SHA  - AES256-SHA  - DES-CBC3-SHA  minTLSVersion: VersionTLS10",
								MarkdownDescription: "old is a TLS security profile based on:  https://wiki.mozilla.org/Security/Server_Side_TLS#Old_backward_compatibility  and looks like this (yaml):  ciphers:  - TLS_AES_128_GCM_SHA256  - TLS_AES_256_GCM_SHA384  - TLS_CHACHA20_POLY1305_SHA256  - ECDHE-ECDSA-AES128-GCM-SHA256  - ECDHE-RSA-AES128-GCM-SHA256  - ECDHE-ECDSA-AES256-GCM-SHA384  - ECDHE-RSA-AES256-GCM-SHA384  - ECDHE-ECDSA-CHACHA20-POLY1305  - ECDHE-RSA-CHACHA20-POLY1305  - DHE-RSA-AES128-GCM-SHA256  - DHE-RSA-AES256-GCM-SHA384  - DHE-RSA-CHACHA20-POLY1305  - ECDHE-ECDSA-AES128-SHA256  - ECDHE-RSA-AES128-SHA256  - ECDHE-ECDSA-AES128-SHA  - ECDHE-RSA-AES128-SHA  - ECDHE-ECDSA-AES256-SHA384  - ECDHE-RSA-AES256-SHA384  - ECDHE-ECDSA-AES256-SHA  - ECDHE-RSA-AES256-SHA  - DHE-RSA-AES128-SHA256  - DHE-RSA-AES256-SHA256  - AES128-GCM-SHA256  - AES256-GCM-SHA384  - AES128-SHA256  - AES256-SHA256  - AES128-SHA  - AES256-SHA  - DES-CBC3-SHA  minTLSVersion: VersionTLS10",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "type is one of Old, Intermediate, Modern or Custom. Custom provides the ability to specify individual TLS security profile parameters. Old, Intermediate and Modern are TLS security profiles based on:  https://wiki.mozilla.org/Security/Server_Side_TLS#Recommended_configurations  The profiles are intent based, so they may change over time as new ciphers are developed and existing ciphers are found to be insecure.  Depending on precisely which ciphers are available to a process, the list may be reduced.  Note that the Modern profile is currently not supported because it is not yet well adopted by common software libraries.",
								MarkdownDescription: "type is one of Old, Intermediate, Modern or Custom. Custom provides the ability to specify individual TLS security profile parameters. Old, Intermediate and Modern are TLS security profiles based on:  https://wiki.mozilla.org/Security/Server_Side_TLS#Recommended_configurations  The profiles are intent based, so they may change over time as new ciphers are developed and existing ciphers are found to be insecure.  Depending on precisely which ciphers are available to a process, the list may be reduced.  Note that the Modern profile is currently not supported because it is not yet well adopted by common software libraries.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Old", "Intermediate", "Modern", "Custom"),
								},
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

func (r *ConfigOpenshiftIoApiserverV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_config_openshift_io_api_server_v1_manifest")

	var model ConfigOpenshiftIoApiserverV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("config.openshift.io/v1")
	model.Kind = pointer.String("APIServer")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
