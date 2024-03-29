/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package trust_cert_manager_io_v1alpha1

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
	_ datasource.DataSource = &TrustCertManagerIoBundleV1Alpha1Manifest{}
)

func NewTrustCertManagerIoBundleV1Alpha1Manifest() datasource.DataSource {
	return &TrustCertManagerIoBundleV1Alpha1Manifest{}
}

type TrustCertManagerIoBundleV1Alpha1Manifest struct{}

type TrustCertManagerIoBundleV1Alpha1ManifestData struct {
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
		Sources *[]struct {
			ConfigMap *struct {
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Selector *struct {
					MatchExpressions *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
					MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
				} `tfsdk:"selector" json:"selector,omitempty"`
			} `tfsdk:"config_map" json:"configMap,omitempty"`
			InLine *string `tfsdk:"in_line" json:"inLine,omitempty"`
			Secret *struct {
				Key      *string `tfsdk:"key" json:"key,omitempty"`
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
			UseDefaultCAs *bool `tfsdk:"use_default_c_as" json:"useDefaultCAs,omitempty"`
		} `tfsdk:"sources" json:"sources,omitempty"`
		Target *struct {
			AdditionalFormats *struct {
				Jks *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Password *string `tfsdk:"password" json:"password,omitempty"`
				} `tfsdk:"jks" json:"jks,omitempty"`
				Pkcs12 *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Password *string `tfsdk:"password" json:"password,omitempty"`
				} `tfsdk:"pkcs12" json:"pkcs12,omitempty"`
			} `tfsdk:"additional_formats" json:"additionalFormats,omitempty"`
			ConfigMap *struct {
				Key *string `tfsdk:"key" json:"key,omitempty"`
			} `tfsdk:"config_map" json:"configMap,omitempty"`
			NamespaceSelector *struct {
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
			Secret *struct {
				Key *string `tfsdk:"key" json:"key,omitempty"`
			} `tfsdk:"secret" json:"secret,omitempty"`
		} `tfsdk:"target" json:"target,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *TrustCertManagerIoBundleV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_trust_cert_manager_io_bundle_v1alpha1_manifest"
}

func (r *TrustCertManagerIoBundleV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
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
				Description:         "Desired state of the Bundle resource.",
				MarkdownDescription: "Desired state of the Bundle resource.",
				Attributes: map[string]schema.Attribute{
					"sources": schema.ListNestedAttribute{
						Description:         "Sources is a set of references to data whose data will sync to the target.",
						MarkdownDescription: "Sources is a set of references to data whose data will sync to the target.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"config_map": schema.SingleNestedAttribute{
									Description:         "ConfigMap is a reference (by name) to a ConfigMap's 'data' key, or to alist of ConfigMap's 'data' key using label selector, in the trust Namespace.",
									MarkdownDescription: "ConfigMap is a reference (by name) to a ConfigMap's 'data' key, or to alist of ConfigMap's 'data' key using label selector, in the trust Namespace.",
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "Key is the key of the entry in the object's 'data' field to be used.",
											MarkdownDescription: "Key is the key of the entry in the object's 'data' field to be used.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name is the name of the source object in the trust Namespace.This field must be left empty when 'selector' is set",
											MarkdownDescription: "Name is the name of the source object in the trust Namespace.This field must be left empty when 'selector' is set",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"selector": schema.SingleNestedAttribute{
											Description:         "Selector is the label selector to use to fetch a list of objects. Must not be setwhen 'Name' is set.",
											MarkdownDescription: "Selector is the label selector to use to fetch a list of objects. Must not be setwhen 'Name' is set.",
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
																Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"values": schema.ListAttribute{
																Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
													Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
													MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

								"in_line": schema.StringAttribute{
									Description:         "InLine is a simple string to append as the source data.",
									MarkdownDescription: "InLine is a simple string to append as the source data.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"secret": schema.SingleNestedAttribute{
									Description:         "Secret is a reference (by name) to a Secret's 'data' key, or to alist of Secret's 'data' key using label selector, in the trust Namespace.",
									MarkdownDescription: "Secret is a reference (by name) to a Secret's 'data' key, or to alist of Secret's 'data' key using label selector, in the trust Namespace.",
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "Key is the key of the entry in the object's 'data' field to be used.",
											MarkdownDescription: "Key is the key of the entry in the object's 'data' field to be used.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name is the name of the source object in the trust Namespace.This field must be left empty when 'selector' is set",
											MarkdownDescription: "Name is the name of the source object in the trust Namespace.This field must be left empty when 'selector' is set",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"selector": schema.SingleNestedAttribute{
											Description:         "Selector is the label selector to use to fetch a list of objects. Must not be setwhen 'Name' is set.",
											MarkdownDescription: "Selector is the label selector to use to fetch a list of objects. Must not be setwhen 'Name' is set.",
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
																Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"values": schema.ListAttribute{
																Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
													Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
													MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

								"use_default_c_as": schema.BoolAttribute{
									Description:         "UseDefaultCAs, when true, requests the default CA bundle to be used as a source.Default CAs are available if trust-manager was installed via Helmor was otherwise set up to include a package-injecting init container by using the'--default-package-location' flag when starting the trust-manager controller.If default CAs were not configured at start-up, any request to use the defaultCAs will fail.The version of the default CA package which is used for a Bundle is stored in thedefaultCAPackageVersion field of the Bundle's status field.",
									MarkdownDescription: "UseDefaultCAs, when true, requests the default CA bundle to be used as a source.Default CAs are available if trust-manager was installed via Helmor was otherwise set up to include a package-injecting init container by using the'--default-package-location' flag when starting the trust-manager controller.If default CAs were not configured at start-up, any request to use the defaultCAs will fail.The version of the default CA package which is used for a Bundle is stored in thedefaultCAPackageVersion field of the Bundle's status field.",
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

					"target": schema.SingleNestedAttribute{
						Description:         "Target is the target location in all namespaces to sync source data to.",
						MarkdownDescription: "Target is the target location in all namespaces to sync source data to.",
						Attributes: map[string]schema.Attribute{
							"additional_formats": schema.SingleNestedAttribute{
								Description:         "AdditionalFormats specifies any additional formats to write to the target",
								MarkdownDescription: "AdditionalFormats specifies any additional formats to write to the target",
								Attributes: map[string]schema.Attribute{
									"jks": schema.SingleNestedAttribute{
										Description:         "JKS requests a JKS-formatted binary trust bundle to be written to the target.The bundle has 'changeit' as the default password.For more information refer to this link https://cert-manager.io/docs/faq/#keystore-passwords",
										MarkdownDescription: "JKS requests a JKS-formatted binary trust bundle to be written to the target.The bundle has 'changeit' as the default password.For more information refer to this link https://cert-manager.io/docs/faq/#keystore-passwords",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "Key is the key of the entry in the object's 'data' field to be used.",
												MarkdownDescription: "Key is the key of the entry in the object's 'data' field to be used.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"password": schema.StringAttribute{
												Description:         "Password for JKS trust store",
												MarkdownDescription: "Password for JKS trust store",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
													stringvalidator.LengthAtMost(128),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"pkcs12": schema.SingleNestedAttribute{
										Description:         "PKCS12 requests a PKCS12-formatted binary trust bundle to be written to the target.The bundle is by default created without a password.",
										MarkdownDescription: "PKCS12 requests a PKCS12-formatted binary trust bundle to be written to the target.The bundle is by default created without a password.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "Key is the key of the entry in the object's 'data' field to be used.",
												MarkdownDescription: "Key is the key of the entry in the object's 'data' field to be used.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"password": schema.StringAttribute{
												Description:         "Password for PKCS12 trust store",
												MarkdownDescription: "Password for PKCS12 trust store",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtMost(128),
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

							"config_map": schema.SingleNestedAttribute{
								Description:         "ConfigMap is the target ConfigMap in Namespaces that all Bundle sourcedata will be synced to.",
								MarkdownDescription: "ConfigMap is the target ConfigMap in Namespaces that all Bundle sourcedata will be synced to.",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "Key is the key of the entry in the object's 'data' field to be used.",
										MarkdownDescription: "Key is the key of the entry in the object's 'data' field to be used.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"namespace_selector": schema.SingleNestedAttribute{
								Description:         "NamespaceSelector will, if set, only sync the target resource inNamespaces which match the selector.",
								MarkdownDescription: "NamespaceSelector will, if set, only sync the target resource inNamespaces which match the selector.",
								Attributes: map[string]schema.Attribute{
									"match_labels": schema.MapAttribute{
										Description:         "MatchLabels matches on the set of labels that must be present on aNamespace for the Bundle target to be synced there.",
										MarkdownDescription: "MatchLabels matches on the set of labels that must be present on aNamespace for the Bundle target to be synced there.",
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

							"secret": schema.SingleNestedAttribute{
								Description:         "Secret is the target Secret that all Bundle source data will be synced to.Using Secrets as targets is only supported if enabled at trust-manager startup.By default, trust-manager has no permissions for writing to secrets and can only read secrets in the trust namespace.",
								MarkdownDescription: "Secret is the target Secret that all Bundle source data will be synced to.Using Secrets as targets is only supported if enabled at trust-manager startup.By default, trust-manager has no permissions for writing to secrets and can only read secrets in the trust namespace.",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "Key is the key of the entry in the object's 'data' field to be used.",
										MarkdownDescription: "Key is the key of the entry in the object's 'data' field to be used.",
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
						Required: true,
						Optional: false,
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

func (r *TrustCertManagerIoBundleV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_trust_cert_manager_io_bundle_v1alpha1_manifest")

	var model TrustCertManagerIoBundleV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("trust.cert-manager.io/v1alpha1")
	model.Kind = pointer.String("Bundle")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
