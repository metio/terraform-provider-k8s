/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package secrets_hashicorp_com_v1beta1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
	_ datasource.DataSource = &SecretsHashicorpComVaultStaticSecretV1Beta1Manifest{}
)

func NewSecretsHashicorpComVaultStaticSecretV1Beta1Manifest() datasource.DataSource {
	return &SecretsHashicorpComVaultStaticSecretV1Beta1Manifest{}
}

type SecretsHashicorpComVaultStaticSecretV1Beta1Manifest struct{}

type SecretsHashicorpComVaultStaticSecretV1Beta1ManifestData struct {
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
		HmacSecretData        *bool   `tfsdk:"hmac_secret_data" json:"hmacSecretData,omitempty"`
		Mount                 *string `tfsdk:"mount" json:"mount,omitempty"`
		Namespace             *string `tfsdk:"namespace" json:"namespace,omitempty"`
		Path                  *string `tfsdk:"path" json:"path,omitempty"`
		RefreshAfter          *string `tfsdk:"refresh_after" json:"refreshAfter,omitempty"`
		RolloutRestartTargets *[]struct {
			Kind *string `tfsdk:"kind" json:"kind,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"rollout_restart_targets" json:"rolloutRestartTargets,omitempty"`
		SyncConfig *struct {
			InstantUpdates *bool `tfsdk:"instant_updates" json:"instantUpdates,omitempty"`
		} `tfsdk:"sync_config" json:"syncConfig,omitempty"`
		Type         *string `tfsdk:"type" json:"type,omitempty"`
		VaultAuthRef *string `tfsdk:"vault_auth_ref" json:"vaultAuthRef,omitempty"`
		Version      *int64  `tfsdk:"version" json:"version,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SecretsHashicorpComVaultStaticSecretV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_secrets_hashicorp_com_vault_static_secret_v1beta1_manifest"
}

func (r *SecretsHashicorpComVaultStaticSecretV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "VaultStaticSecret is the Schema for the vaultstaticsecrets API",
		MarkdownDescription: "VaultStaticSecret is the Schema for the vaultstaticsecrets API",
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
				Description:         "VaultStaticSecretSpec defines the desired state of VaultStaticSecret",
				MarkdownDescription: "VaultStaticSecretSpec defines the desired state of VaultStaticSecret",
				Attributes: map[string]schema.Attribute{
					"destination": schema.SingleNestedAttribute{
						Description:         "Destination provides configuration necessary for syncing the Vault secret to Kubernetes.",
						MarkdownDescription: "Destination provides configuration necessary for syncing the Vault secret to Kubernetes.",
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

					"hmac_secret_data": schema.BoolAttribute{
						Description:         "HMACSecretData determines whether the Operator computes the HMAC of the Secret's data. The MAC value will be stored in the resource's Status.SecretMac field, and will be used for drift detection and during incoming Vault secret comparison. Enabling this feature is recommended to ensure that Secret's data stays consistent with Vault.",
						MarkdownDescription: "HMACSecretData determines whether the Operator computes the HMAC of the Secret's data. The MAC value will be stored in the resource's Status.SecretMac field, and will be used for drift detection and during incoming Vault secret comparison. Enabling this feature is recommended to ensure that Secret's data stays consistent with Vault.",
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
						Description:         "Namespace of the secrets engine mount in Vault. If not set, the namespace that's part of VaultAuth resource will be inferred.",
						MarkdownDescription: "Namespace of the secrets engine mount in Vault. If not set, the namespace that's part of VaultAuth resource will be inferred.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"path": schema.StringAttribute{
						Description:         "Path of the secret in Vault, corresponds to the 'path' parameter for, kv-v1: https://developer.hashicorp.com/vault/api-docs/secret/kv/kv-v1#read-secret kv-v2: https://developer.hashicorp.com/vault/api-docs/secret/kv/kv-v2#read-secret-version",
						MarkdownDescription: "Path of the secret in Vault, corresponds to the 'path' parameter for, kv-v1: https://developer.hashicorp.com/vault/api-docs/secret/kv/kv-v1#read-secret kv-v2: https://developer.hashicorp.com/vault/api-docs/secret/kv/kv-v2#read-secret-version",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"refresh_after": schema.StringAttribute{
						Description:         "RefreshAfter a period of time, in duration notation e.g. 30s, 1m, 24h",
						MarkdownDescription: "RefreshAfter a period of time, in duration notation e.g. 30s, 1m, 24h",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(s|m|h))$`), ""),
						},
					},

					"rollout_restart_targets": schema.ListNestedAttribute{
						Description:         "RolloutRestartTargets should be configured whenever the application(s) consuming the Vault secret does not support dynamically reloading a rotated secret. In that case one, or more RolloutRestartTarget(s) can be configured here. The Operator will trigger a 'rollout-restart' for each target whenever the Vault secret changes between reconciliation events. All configured targets wil be ignored if HMACSecretData is set to false. See RolloutRestartTarget for more details.",
						MarkdownDescription: "RolloutRestartTargets should be configured whenever the application(s) consuming the Vault secret does not support dynamically reloading a rotated secret. In that case one, or more RolloutRestartTarget(s) can be configured here. The Operator will trigger a 'rollout-restart' for each target whenever the Vault secret changes between reconciliation events. All configured targets wil be ignored if HMACSecretData is set to false. See RolloutRestartTarget for more details.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"kind": schema.StringAttribute{
									Description:         "Kind of the resource",
									MarkdownDescription: "Kind of the resource",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("Deployment", "DaemonSet", "StatefulSet", "argo.Rollout"),
									},
								},

								"name": schema.StringAttribute{
									Description:         "Name of the resource",
									MarkdownDescription: "Name of the resource",
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

					"sync_config": schema.SingleNestedAttribute{
						Description:         "SyncConfig configures sync behavior from Vault to VSO",
						MarkdownDescription: "SyncConfig configures sync behavior from Vault to VSO",
						Attributes: map[string]schema.Attribute{
							"instant_updates": schema.BoolAttribute{
								Description:         "InstantUpdates is a flag to indicate that event-driven updates are enabled for this VaultStaticSecret",
								MarkdownDescription: "InstantUpdates is a flag to indicate that event-driven updates are enabled for this VaultStaticSecret",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"type": schema.StringAttribute{
						Description:         "Type of the Vault static secret",
						MarkdownDescription: "Type of the Vault static secret",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("kv-v1", "kv-v2"),
						},
					},

					"vault_auth_ref": schema.StringAttribute{
						Description:         "VaultAuthRef to the VaultAuth resource, can be prefixed with a namespace, eg: 'namespaceA/vaultAuthRefB'. If no namespace prefix is provided it will default to the namespace of the VaultAuth CR. If no value is specified for VaultAuthRef the Operator will default to the 'default' VaultAuth, configured in the operator's namespace.",
						MarkdownDescription: "VaultAuthRef to the VaultAuth resource, can be prefixed with a namespace, eg: 'namespaceA/vaultAuthRefB'. If no namespace prefix is provided it will default to the namespace of the VaultAuth CR. If no value is specified for VaultAuthRef the Operator will default to the 'default' VaultAuth, configured in the operator's namespace.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"version": schema.Int64Attribute{
						Description:         "Version of the secret to fetch. Only valid for type kv-v2. Corresponds to version query parameter: https://developer.hashicorp.com/vault/api-docs/secret/kv/kv-v2#version",
						MarkdownDescription: "Version of the secret to fetch. Only valid for type kv-v2. Corresponds to version query parameter: https://developer.hashicorp.com/vault/api-docs/secret/kv/kv-v2#version",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *SecretsHashicorpComVaultStaticSecretV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_secrets_hashicorp_com_vault_static_secret_v1beta1_manifest")

	var model SecretsHashicorpComVaultStaticSecretV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("secrets.hashicorp.com/v1beta1")
	model.Kind = pointer.String("VaultStaticSecret")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
