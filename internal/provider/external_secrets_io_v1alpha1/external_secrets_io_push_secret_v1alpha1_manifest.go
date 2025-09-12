/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package external_secrets_io_v1alpha1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
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
	_ datasource.DataSource = &ExternalSecretsIoPushSecretV1Alpha1Manifest{}
)

func NewExternalSecretsIoPushSecretV1Alpha1Manifest() datasource.DataSource {
	return &ExternalSecretsIoPushSecretV1Alpha1Manifest{}
}

type ExternalSecretsIoPushSecretV1Alpha1Manifest struct{}

type ExternalSecretsIoPushSecretV1Alpha1ManifestData struct {
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
		Data *[]struct {
			ConversionStrategy *string `tfsdk:"conversion_strategy" json:"conversionStrategy,omitempty"`
			Match              *struct {
				RemoteRef *struct {
					Property  *string `tfsdk:"property" json:"property,omitempty"`
					RemoteKey *string `tfsdk:"remote_key" json:"remoteKey,omitempty"`
				} `tfsdk:"remote_ref" json:"remoteRef,omitempty"`
				SecretKey *string `tfsdk:"secret_key" json:"secretKey,omitempty"`
			} `tfsdk:"match" json:"match,omitempty"`
			Metadata *map[string]string `tfsdk:"metadata" json:"metadata,omitempty"`
		} `tfsdk:"data" json:"data,omitempty"`
		DeletionPolicy  *string `tfsdk:"deletion_policy" json:"deletionPolicy,omitempty"`
		RefreshInterval *string `tfsdk:"refresh_interval" json:"refreshInterval,omitempty"`
		SecretStoreRefs *[]struct {
			Kind          *string `tfsdk:"kind" json:"kind,omitempty"`
			LabelSelector *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"secret_store_refs" json:"secretStoreRefs,omitempty"`
		Selector *struct {
			GeneratorRef *struct {
				ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
				Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
				Name       *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"generator_ref" json:"generatorRef,omitempty"`
			Secret *struct {
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Selector *struct {
					MatchExpressions *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
					MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
				} `tfsdk:"selector" json:"selector,omitempty"`
			} `tfsdk:"secret" json:"secret,omitempty"`
		} `tfsdk:"selector" json:"selector,omitempty"`
		Template *struct {
			Data          *map[string]string `tfsdk:"data" json:"data,omitempty"`
			EngineVersion *string            `tfsdk:"engine_version" json:"engineVersion,omitempty"`
			MergePolicy   *string            `tfsdk:"merge_policy" json:"mergePolicy,omitempty"`
			Metadata      *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Finalizers  *[]string          `tfsdk:"finalizers" json:"finalizers,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"metadata" json:"metadata,omitempty"`
			TemplateFrom *[]struct {
				ConfigMap *struct {
					Items *[]struct {
						Key        *string `tfsdk:"key" json:"key,omitempty"`
						TemplateAs *string `tfsdk:"template_as" json:"templateAs,omitempty"`
					} `tfsdk:"items" json:"items,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"config_map" json:"configMap,omitempty"`
				Literal *string `tfsdk:"literal" json:"literal,omitempty"`
				Secret  *struct {
					Items *[]struct {
						Key        *string `tfsdk:"key" json:"key,omitempty"`
						TemplateAs *string `tfsdk:"template_as" json:"templateAs,omitempty"`
					} `tfsdk:"items" json:"items,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"secret" json:"secret,omitempty"`
				Target *string `tfsdk:"target" json:"target,omitempty"`
			} `tfsdk:"template_from" json:"templateFrom,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"template" json:"template,omitempty"`
		UpdatePolicy *string `tfsdk:"update_policy" json:"updatePolicy,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ExternalSecretsIoPushSecretV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_external_secrets_io_push_secret_v1alpha1_manifest"
}

func (r *ExternalSecretsIoPushSecretV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "PushSecretSpec configures the behavior of the PushSecret.",
				MarkdownDescription: "PushSecretSpec configures the behavior of the PushSecret.",
				Attributes: map[string]schema.Attribute{
					"data": schema.ListNestedAttribute{
						Description:         "Secret Data that should be pushed to providers",
						MarkdownDescription: "Secret Data that should be pushed to providers",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"conversion_strategy": schema.StringAttribute{
									Description:         "Used to define a conversion Strategy for the secret keys",
									MarkdownDescription: "Used to define a conversion Strategy for the secret keys",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("None", "ReverseUnicode"),
									},
								},

								"match": schema.SingleNestedAttribute{
									Description:         "Match a given Secret Key to be pushed to the provider.",
									MarkdownDescription: "Match a given Secret Key to be pushed to the provider.",
									Attributes: map[string]schema.Attribute{
										"remote_ref": schema.SingleNestedAttribute{
											Description:         "Remote Refs to push to providers.",
											MarkdownDescription: "Remote Refs to push to providers.",
											Attributes: map[string]schema.Attribute{
												"property": schema.StringAttribute{
													Description:         "Name of the property in the resulting secret",
													MarkdownDescription: "Name of the property in the resulting secret",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"remote_key": schema.StringAttribute{
													Description:         "Name of the resulting provider secret.",
													MarkdownDescription: "Name of the resulting provider secret.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"secret_key": schema.StringAttribute{
											Description:         "Secret Key to be pushed",
											MarkdownDescription: "Secret Key to be pushed",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},

								"metadata": schema.MapAttribute{
									Description:         "Metadata is metadata attached to the secret. The structure of metadata is provider specific, please look it up in the provider documentation.",
									MarkdownDescription: "Metadata is metadata attached to the secret. The structure of metadata is provider specific, please look it up in the provider documentation.",
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

					"deletion_policy": schema.StringAttribute{
						Description:         "Deletion Policy to handle Secrets in the provider.",
						MarkdownDescription: "Deletion Policy to handle Secrets in the provider.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Delete", "None"),
						},
					},

					"refresh_interval": schema.StringAttribute{
						Description:         "The Interval to which External Secrets will try to push a secret definition",
						MarkdownDescription: "The Interval to which External Secrets will try to push a secret definition",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"secret_store_refs": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"kind": schema.StringAttribute{
									Description:         "Kind of the SecretStore resource (SecretStore or ClusterSecretStore)",
									MarkdownDescription: "Kind of the SecretStore resource (SecretStore or ClusterSecretStore)",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("SecretStore", "ClusterSecretStore"),
									},
								},

								"label_selector": schema.SingleNestedAttribute{
									Description:         "Optionally, sync to secret stores with label selector",
									MarkdownDescription: "Optionally, sync to secret stores with label selector",
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

								"name": schema.StringAttribute{
									Description:         "Optionally, sync to the SecretStore of the given name",
									MarkdownDescription: "Optionally, sync to the SecretStore of the given name",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(253),
										stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
									},
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"selector": schema.SingleNestedAttribute{
						Description:         "The Secret Selector (k8s source) for the Push Secret",
						MarkdownDescription: "The Secret Selector (k8s source) for the Push Secret",
						Attributes: map[string]schema.Attribute{
							"generator_ref": schema.SingleNestedAttribute{
								Description:         "Point to a generator to create a Secret.",
								MarkdownDescription: "Point to a generator to create a Secret.",
								Attributes: map[string]schema.Attribute{
									"api_version": schema.StringAttribute{
										Description:         "Specify the apiVersion of the generator resource",
										MarkdownDescription: "Specify the apiVersion of the generator resource",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"kind": schema.StringAttribute{
										Description:         "Specify the Kind of the generator resource",
										MarkdownDescription: "Specify the Kind of the generator resource",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("ACRAccessToken", "ClusterGenerator", "ECRAuthorizationToken", "Fake", "GCRAccessToken", "GithubAccessToken", "QuayAccessToken", "Password", "SSHKey", "STSSessionToken", "UUID", "VaultDynamicSecret", "Webhook", "Grafana", "MFA"),
										},
									},

									"name": schema.StringAttribute{
										Description:         "Specify the name of the generator resource",
										MarkdownDescription: "Specify the name of the generator resource",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(1),
											stringvalidator.LengthAtMost(253),
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
								Validators: []validator.Object{
									objectvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("secret")),
								},
							},

							"secret": schema.SingleNestedAttribute{
								Description:         "Select a Secret to Push.",
								MarkdownDescription: "Select a Secret to Push.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of the Secret. The Secret must exist in the same namespace as the PushSecret manifest.",
										MarkdownDescription: "Name of the Secret. The Secret must exist in the same namespace as the PushSecret manifest.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(1),
											stringvalidator.LengthAtMost(253),
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
										},
									},

									"selector": schema.SingleNestedAttribute{
										Description:         "Selector chooses secrets using a labelSelector.",
										MarkdownDescription: "Selector chooses secrets using a labelSelector.",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
								Validators: []validator.Object{
									objectvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("generator_ref")),
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"template": schema.SingleNestedAttribute{
						Description:         "Template defines a blueprint for the created Secret resource.",
						MarkdownDescription: "Template defines a blueprint for the created Secret resource.",
						Attributes: map[string]schema.Attribute{
							"data": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"engine_version": schema.StringAttribute{
								Description:         "EngineVersion specifies the template engine version that should be used to compile/execute the template specified in .data and .templateFrom[].",
								MarkdownDescription: "EngineVersion specifies the template engine version that should be used to compile/execute the template specified in .data and .templateFrom[].",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("v2"),
								},
							},

							"merge_policy": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Replace", "Merge"),
								},
							},

							"metadata": schema.SingleNestedAttribute{
								Description:         "ExternalSecretTemplateMetadata defines metadata fields for the Secret blueprint.",
								MarkdownDescription: "ExternalSecretTemplateMetadata defines metadata fields for the Secret blueprint.",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"finalizers": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.MapAttribute{
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

							"template_from": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"config_map": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"items": schema.ListNestedAttribute{
													Description:         "A list of keys in the ConfigMap/Secret to use as templates for Secret data",
													MarkdownDescription: "A list of keys in the ConfigMap/Secret to use as templates for Secret data",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "A key in the ConfigMap/Secret",
																MarkdownDescription: "A key in the ConfigMap/Secret",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																	stringvalidator.LengthAtMost(253),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[-._a-zA-Z0-9]+$`), ""),
																},
															},

															"template_as": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("Values", "KeysAndValues"),
																},
															},
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},

												"name": schema.StringAttribute{
													Description:         "The name of the ConfigMap/Secret resource",
													MarkdownDescription: "The name of the ConfigMap/Secret resource",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
														stringvalidator.LengthAtMost(253),
														stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"literal": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secret": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"items": schema.ListNestedAttribute{
													Description:         "A list of keys in the ConfigMap/Secret to use as templates for Secret data",
													MarkdownDescription: "A list of keys in the ConfigMap/Secret to use as templates for Secret data",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "A key in the ConfigMap/Secret",
																MarkdownDescription: "A key in the ConfigMap/Secret",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																	stringvalidator.LengthAtMost(253),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[-._a-zA-Z0-9]+$`), ""),
																},
															},

															"template_as": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("Values", "KeysAndValues"),
																},
															},
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},

												"name": schema.StringAttribute{
													Description:         "The name of the ConfigMap/Secret resource",
													MarkdownDescription: "The name of the ConfigMap/Secret resource",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
														stringvalidator.LengthAtMost(253),
														stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"target": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Data", "Annotations", "Labels"),
											},
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"type": schema.StringAttribute{
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

					"update_policy": schema.StringAttribute{
						Description:         "UpdatePolicy to handle Secrets in the provider.",
						MarkdownDescription: "UpdatePolicy to handle Secrets in the provider.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Replace", "IfNotExists"),
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

func (r *ExternalSecretsIoPushSecretV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_external_secrets_io_push_secret_v1alpha1_manifest")

	var model ExternalSecretsIoPushSecretV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("external-secrets.io/v1alpha1")
	model.Kind = pointer.String("PushSecret")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
