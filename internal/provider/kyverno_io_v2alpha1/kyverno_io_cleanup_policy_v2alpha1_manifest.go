/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kyverno_io_v2alpha1

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
	_ datasource.DataSource = &KyvernoIoCleanupPolicyV2Alpha1Manifest{}
)

func NewKyvernoIoCleanupPolicyV2Alpha1Manifest() datasource.DataSource {
	return &KyvernoIoCleanupPolicyV2Alpha1Manifest{}
}

type KyvernoIoCleanupPolicyV2Alpha1Manifest struct{}

type KyvernoIoCleanupPolicyV2Alpha1ManifestData struct {
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
		Conditions *struct {
			All *[]struct {
				Key      *map[string]string `tfsdk:"key" json:"key,omitempty"`
				Message  *string            `tfsdk:"message" json:"message,omitempty"`
				Operator *string            `tfsdk:"operator" json:"operator,omitempty"`
				Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"all" json:"all,omitempty"`
			Any *[]struct {
				Key      *map[string]string `tfsdk:"key" json:"key,omitempty"`
				Message  *string            `tfsdk:"message" json:"message,omitempty"`
				Operator *string            `tfsdk:"operator" json:"operator,omitempty"`
				Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"any" json:"any,omitempty"`
		} `tfsdk:"conditions" json:"conditions,omitempty"`
		Context *[]struct {
			ApiCall *struct {
				Data *[]struct {
					Key   *string            `tfsdk:"key" json:"key,omitempty"`
					Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"data" json:"data,omitempty"`
				JmesPath *string `tfsdk:"jmes_path" json:"jmesPath,omitempty"`
				Method   *string `tfsdk:"method" json:"method,omitempty"`
				Service  *struct {
					CaBundle *string `tfsdk:"ca_bundle" json:"caBundle,omitempty"`
					Url      *string `tfsdk:"url" json:"url,omitempty"`
				} `tfsdk:"service" json:"service,omitempty"`
				UrlPath *string `tfsdk:"url_path" json:"urlPath,omitempty"`
			} `tfsdk:"api_call" json:"apiCall,omitempty"`
			ConfigMap *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"config_map" json:"configMap,omitempty"`
			ImageRegistry *struct {
				ImageRegistryCredentials *struct {
					AllowInsecureRegistry *bool     `tfsdk:"allow_insecure_registry" json:"allowInsecureRegistry,omitempty"`
					Providers             *[]string `tfsdk:"providers" json:"providers,omitempty"`
					Secrets               *[]string `tfsdk:"secrets" json:"secrets,omitempty"`
				} `tfsdk:"image_registry_credentials" json:"imageRegistryCredentials,omitempty"`
				JmesPath  *string `tfsdk:"jmes_path" json:"jmesPath,omitempty"`
				Reference *string `tfsdk:"reference" json:"reference,omitempty"`
			} `tfsdk:"image_registry" json:"imageRegistry,omitempty"`
			Name     *string `tfsdk:"name" json:"name,omitempty"`
			Variable *struct {
				Default  *map[string]string `tfsdk:"default" json:"default,omitempty"`
				JmesPath *string            `tfsdk:"jmes_path" json:"jmesPath,omitempty"`
				Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"variable" json:"variable,omitempty"`
		} `tfsdk:"context" json:"context,omitempty"`
		Exclude *struct {
			All *[]struct {
				ClusterRoles *[]string `tfsdk:"cluster_roles" json:"clusterRoles,omitempty"`
				Resources    *struct {
					Annotations       *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Kinds             *[]string          `tfsdk:"kinds" json:"kinds,omitempty"`
					Name              *string            `tfsdk:"name" json:"name,omitempty"`
					Names             *[]string          `tfsdk:"names" json:"names,omitempty"`
					NamespaceSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
					Namespaces *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
					Operations *[]string `tfsdk:"operations" json:"operations,omitempty"`
					Selector   *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"selector" json:"selector,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				Roles    *[]string `tfsdk:"roles" json:"roles,omitempty"`
				Subjects *[]struct {
					ApiGroup  *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
					Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"subjects" json:"subjects,omitempty"`
			} `tfsdk:"all" json:"all,omitempty"`
			Any *[]struct {
				ClusterRoles *[]string `tfsdk:"cluster_roles" json:"clusterRoles,omitempty"`
				Resources    *struct {
					Annotations       *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Kinds             *[]string          `tfsdk:"kinds" json:"kinds,omitempty"`
					Name              *string            `tfsdk:"name" json:"name,omitempty"`
					Names             *[]string          `tfsdk:"names" json:"names,omitempty"`
					NamespaceSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
					Namespaces *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
					Operations *[]string `tfsdk:"operations" json:"operations,omitempty"`
					Selector   *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"selector" json:"selector,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				Roles    *[]string `tfsdk:"roles" json:"roles,omitempty"`
				Subjects *[]struct {
					ApiGroup  *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
					Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"subjects" json:"subjects,omitempty"`
			} `tfsdk:"any" json:"any,omitempty"`
		} `tfsdk:"exclude" json:"exclude,omitempty"`
		Match *struct {
			All *[]struct {
				ClusterRoles *[]string `tfsdk:"cluster_roles" json:"clusterRoles,omitempty"`
				Resources    *struct {
					Annotations       *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Kinds             *[]string          `tfsdk:"kinds" json:"kinds,omitempty"`
					Name              *string            `tfsdk:"name" json:"name,omitempty"`
					Names             *[]string          `tfsdk:"names" json:"names,omitempty"`
					NamespaceSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
					Namespaces *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
					Operations *[]string `tfsdk:"operations" json:"operations,omitempty"`
					Selector   *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"selector" json:"selector,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				Roles    *[]string `tfsdk:"roles" json:"roles,omitempty"`
				Subjects *[]struct {
					ApiGroup  *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
					Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"subjects" json:"subjects,omitempty"`
			} `tfsdk:"all" json:"all,omitempty"`
			Any *[]struct {
				ClusterRoles *[]string `tfsdk:"cluster_roles" json:"clusterRoles,omitempty"`
				Resources    *struct {
					Annotations       *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Kinds             *[]string          `tfsdk:"kinds" json:"kinds,omitempty"`
					Name              *string            `tfsdk:"name" json:"name,omitempty"`
					Names             *[]string          `tfsdk:"names" json:"names,omitempty"`
					NamespaceSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
					Namespaces *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
					Operations *[]string `tfsdk:"operations" json:"operations,omitempty"`
					Selector   *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"selector" json:"selector,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				Roles    *[]string `tfsdk:"roles" json:"roles,omitempty"`
				Subjects *[]struct {
					ApiGroup  *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
					Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"subjects" json:"subjects,omitempty"`
			} `tfsdk:"any" json:"any,omitempty"`
		} `tfsdk:"match" json:"match,omitempty"`
		Schedule *string `tfsdk:"schedule" json:"schedule,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KyvernoIoCleanupPolicyV2Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kyverno_io_cleanup_policy_v2alpha1_manifest"
}

func (r *KyvernoIoCleanupPolicyV2Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CleanupPolicy defines a rule for resource cleanup.",
		MarkdownDescription: "CleanupPolicy defines a rule for resource cleanup.",
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
				Description:         "Spec declares policy behaviors.",
				MarkdownDescription: "Spec declares policy behaviors.",
				Attributes: map[string]schema.Attribute{
					"conditions": schema.SingleNestedAttribute{
						Description:         "Conditions defines the conditions used to select the resources which will be cleaned up.",
						MarkdownDescription: "Conditions defines the conditions used to select the resources which will be cleaned up.",
						Attributes: map[string]schema.Attribute{
							"all": schema.ListNestedAttribute{
								Description:         "AllConditions enable variable-based conditional rule execution. This is useful for finer control of when an rule is applied. A condition can reference object data using JMESPath notation. Here, all of the conditions need to pass.",
								MarkdownDescription: "AllConditions enable variable-based conditional rule execution. This is useful for finer control of when an rule is applied. A condition can reference object data using JMESPath notation. Here, all of the conditions need to pass.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.MapAttribute{
											Description:         "Key is the context entry (using JMESPath) for conditional rule evaluation.",
											MarkdownDescription: "Key is the context entry (using JMESPath) for conditional rule evaluation.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"message": schema.StringAttribute{
											Description:         "Message is an optional display message",
											MarkdownDescription: "Message is an optional display message",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"operator": schema.StringAttribute{
											Description:         "Operator is the conditional operation to perform. Valid operators are: Equals, NotEquals, In, AnyIn, AllIn, NotIn, AnyNotIn, AllNotIn, GreaterThanOrEquals, GreaterThan, LessThanOrEquals, LessThan, DurationGreaterThanOrEquals, DurationGreaterThan, DurationLessThanOrEquals, DurationLessThan",
											MarkdownDescription: "Operator is the conditional operation to perform. Valid operators are: Equals, NotEquals, In, AnyIn, AllIn, NotIn, AnyNotIn, AllNotIn, GreaterThanOrEquals, GreaterThan, LessThanOrEquals, LessThan, DurationGreaterThanOrEquals, DurationGreaterThan, DurationLessThanOrEquals, DurationLessThan",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Equals", "NotEquals", "AnyIn", "AllIn", "AnyNotIn", "AllNotIn", "GreaterThanOrEquals", "GreaterThan", "LessThanOrEquals", "LessThan", "DurationGreaterThanOrEquals", "DurationGreaterThan", "DurationLessThanOrEquals", "DurationLessThan"),
											},
										},

										"value": schema.MapAttribute{
											Description:         "Value is the conditional value, or set of values. The values can be fixed set or can be variables declared using JMESPath.",
											MarkdownDescription: "Value is the conditional value, or set of values. The values can be fixed set or can be variables declared using JMESPath.",
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

							"any": schema.ListNestedAttribute{
								Description:         "AnyConditions enable variable-based conditional rule execution. This is useful for finer control of when an rule is applied. A condition can reference object data using JMESPath notation. Here, at least one of the conditions need to pass.",
								MarkdownDescription: "AnyConditions enable variable-based conditional rule execution. This is useful for finer control of when an rule is applied. A condition can reference object data using JMESPath notation. Here, at least one of the conditions need to pass.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.MapAttribute{
											Description:         "Key is the context entry (using JMESPath) for conditional rule evaluation.",
											MarkdownDescription: "Key is the context entry (using JMESPath) for conditional rule evaluation.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"message": schema.StringAttribute{
											Description:         "Message is an optional display message",
											MarkdownDescription: "Message is an optional display message",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"operator": schema.StringAttribute{
											Description:         "Operator is the conditional operation to perform. Valid operators are: Equals, NotEquals, In, AnyIn, AllIn, NotIn, AnyNotIn, AllNotIn, GreaterThanOrEquals, GreaterThan, LessThanOrEquals, LessThan, DurationGreaterThanOrEquals, DurationGreaterThan, DurationLessThanOrEquals, DurationLessThan",
											MarkdownDescription: "Operator is the conditional operation to perform. Valid operators are: Equals, NotEquals, In, AnyIn, AllIn, NotIn, AnyNotIn, AllNotIn, GreaterThanOrEquals, GreaterThan, LessThanOrEquals, LessThan, DurationGreaterThanOrEquals, DurationGreaterThan, DurationLessThanOrEquals, DurationLessThan",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Equals", "NotEquals", "AnyIn", "AllIn", "AnyNotIn", "AllNotIn", "GreaterThanOrEquals", "GreaterThan", "LessThanOrEquals", "LessThan", "DurationGreaterThanOrEquals", "DurationGreaterThan", "DurationLessThanOrEquals", "DurationLessThan"),
											},
										},

										"value": schema.MapAttribute{
											Description:         "Value is the conditional value, or set of values. The values can be fixed set or can be variables declared using JMESPath.",
											MarkdownDescription: "Value is the conditional value, or set of values. The values can be fixed set or can be variables declared using JMESPath.",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"context": schema.ListNestedAttribute{
						Description:         "Context defines variables and data sources that can be used during rule execution.",
						MarkdownDescription: "Context defines variables and data sources that can be used during rule execution.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"api_call": schema.SingleNestedAttribute{
									Description:         "APICall is an HTTP request to the Kubernetes API server, or other JSON web service. The data returned is stored in the context with the name for the context entry.",
									MarkdownDescription: "APICall is an HTTP request to the Kubernetes API server, or other JSON web service. The data returned is stored in the context with the name for the context entry.",
									Attributes: map[string]schema.Attribute{
										"data": schema.ListNestedAttribute{
											Description:         "Data specifies the POST data sent to the server.",
											MarkdownDescription: "Data specifies the POST data sent to the server.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "Key is a unique identifier for the data value",
														MarkdownDescription: "Key is a unique identifier for the data value",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"value": schema.MapAttribute{
														Description:         "Value is the data value",
														MarkdownDescription: "Value is the data value",
														ElementType:         types.StringType,
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

										"jmes_path": schema.StringAttribute{
											Description:         "JMESPath is an optional JSON Match Expression that can be used to transform the JSON response returned from the server. For example a JMESPath of 'items | length(@)' applied to the API server response for the URLPath '/apis/apps/v1/deployments' will return the total count of deployments across all namespaces.",
											MarkdownDescription: "JMESPath is an optional JSON Match Expression that can be used to transform the JSON response returned from the server. For example a JMESPath of 'items | length(@)' applied to the API server response for the URLPath '/apis/apps/v1/deployments' will return the total count of deployments across all namespaces.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"method": schema.StringAttribute{
											Description:         "Method is the HTTP request type (GET or POST).",
											MarkdownDescription: "Method is the HTTP request type (GET or POST).",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("GET", "POST"),
											},
										},

										"service": schema.SingleNestedAttribute{
											Description:         "Service is an API call to a JSON web service",
											MarkdownDescription: "Service is an API call to a JSON web service",
											Attributes: map[string]schema.Attribute{
												"ca_bundle": schema.StringAttribute{
													Description:         "CABundle is a PEM encoded CA bundle which will be used to validate the server certificate.",
													MarkdownDescription: "CABundle is a PEM encoded CA bundle which will be used to validate the server certificate.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"url": schema.StringAttribute{
													Description:         "URL is the JSON web service URL. A typical form is 'https://{service}.{namespace}:{port}/{path}'.",
													MarkdownDescription: "URL is the JSON web service URL. A typical form is 'https://{service}.{namespace}:{port}/{path}'.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"url_path": schema.StringAttribute{
											Description:         "URLPath is the URL path to be used in the HTTP GET or POST request to the Kubernetes API server (e.g. '/api/v1/namespaces' or  '/apis/apps/v1/deployments'). The format required is the same format used by the 'kubectl get --raw' command. See https://kyverno.io/docs/writing-policies/external-data-sources/#variables-from-kubernetes-api-server-calls for details.",
											MarkdownDescription: "URLPath is the URL path to be used in the HTTP GET or POST request to the Kubernetes API server (e.g. '/api/v1/namespaces' or  '/apis/apps/v1/deployments'). The format required is the same format used by the 'kubectl get --raw' command. See https://kyverno.io/docs/writing-policies/external-data-sources/#variables-from-kubernetes-api-server-calls for details.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"config_map": schema.SingleNestedAttribute{
									Description:         "ConfigMap is the ConfigMap reference.",
									MarkdownDescription: "ConfigMap is the ConfigMap reference.",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name is the ConfigMap name.",
											MarkdownDescription: "Name is the ConfigMap name.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"namespace": schema.StringAttribute{
											Description:         "Namespace is the ConfigMap namespace.",
											MarkdownDescription: "Namespace is the ConfigMap namespace.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"image_registry": schema.SingleNestedAttribute{
									Description:         "ImageRegistry defines requests to an OCI/Docker V2 registry to fetch image details.",
									MarkdownDescription: "ImageRegistry defines requests to an OCI/Docker V2 registry to fetch image details.",
									Attributes: map[string]schema.Attribute{
										"image_registry_credentials": schema.SingleNestedAttribute{
											Description:         "ImageRegistryCredentials provides credentials that will be used for authentication with registry",
											MarkdownDescription: "ImageRegistryCredentials provides credentials that will be used for authentication with registry",
											Attributes: map[string]schema.Attribute{
												"allow_insecure_registry": schema.BoolAttribute{
													Description:         "AllowInsecureRegistry allows insecure access to a registry.",
													MarkdownDescription: "AllowInsecureRegistry allows insecure access to a registry.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"providers": schema.ListAttribute{
													Description:         "Providers specifies a list of OCI Registry names, whose authentication providers are provided. It can be of one of these values: default,google,azure,amazon,github.",
													MarkdownDescription: "Providers specifies a list of OCI Registry names, whose authentication providers are provided. It can be of one of these values: default,google,azure,amazon,github.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"secrets": schema.ListAttribute{
													Description:         "Secrets specifies a list of secrets that are provided for credentials. Secrets must live in the Kyverno namespace.",
													MarkdownDescription: "Secrets specifies a list of secrets that are provided for credentials. Secrets must live in the Kyverno namespace.",
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

										"jmes_path": schema.StringAttribute{
											Description:         "JMESPath is an optional JSON Match Expression that can be used to transform the ImageData struct returned as a result of processing the image reference.",
											MarkdownDescription: "JMESPath is an optional JSON Match Expression that can be used to transform the ImageData struct returned as a result of processing the image reference.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"reference": schema.StringAttribute{
											Description:         "Reference is image reference to a container image in the registry. Example: ghcr.io/kyverno/kyverno:latest",
											MarkdownDescription: "Reference is image reference to a container image in the registry. Example: ghcr.io/kyverno/kyverno:latest",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"name": schema.StringAttribute{
									Description:         "Name is the variable name.",
									MarkdownDescription: "Name is the variable name.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"variable": schema.SingleNestedAttribute{
									Description:         "Variable defines an arbitrary JMESPath context variable that can be defined inline.",
									MarkdownDescription: "Variable defines an arbitrary JMESPath context variable that can be defined inline.",
									Attributes: map[string]schema.Attribute{
										"default": schema.MapAttribute{
											Description:         "Default is an optional arbitrary JSON object that the variable may take if the JMESPath expression evaluates to nil",
											MarkdownDescription: "Default is an optional arbitrary JSON object that the variable may take if the JMESPath expression evaluates to nil",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"jmes_path": schema.StringAttribute{
											Description:         "JMESPath is an optional JMESPath Expression that can be used to transform the variable.",
											MarkdownDescription: "JMESPath is an optional JMESPath Expression that can be used to transform the variable.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"value": schema.MapAttribute{
											Description:         "Value is any arbitrary JSON object representable in YAML or JSON form.",
											MarkdownDescription: "Value is any arbitrary JSON object representable in YAML or JSON form.",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"exclude": schema.SingleNestedAttribute{
						Description:         "ExcludeResources defines when cleanuppolicy should not be applied. The exclude criteria can include resource information (e.g. kind, name, namespace, labels) and admission review request information like the name or role.",
						MarkdownDescription: "ExcludeResources defines when cleanuppolicy should not be applied. The exclude criteria can include resource information (e.g. kind, name, namespace, labels) and admission review request information like the name or role.",
						Attributes: map[string]schema.Attribute{
							"all": schema.ListNestedAttribute{
								Description:         "All allows specifying resources which will be ANDed",
								MarkdownDescription: "All allows specifying resources which will be ANDed",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"cluster_roles": schema.ListAttribute{
											Description:         "ClusterRoles is the list of cluster-wide role names for the user.",
											MarkdownDescription: "ClusterRoles is the list of cluster-wide role names for the user.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"resources": schema.SingleNestedAttribute{
											Description:         "ResourceDescription contains information about the resource being created or modified.",
											MarkdownDescription: "ResourceDescription contains information about the resource being created or modified.",
											Attributes: map[string]schema.Attribute{
												"annotations": schema.MapAttribute{
													Description:         "Annotations is a  map of annotations (key-value pairs of type string). Annotation keys and values support the wildcard characters '*' (matches zero or many characters) and '?' (matches at least one character).",
													MarkdownDescription: "Annotations is a  map of annotations (key-value pairs of type string). Annotation keys and values support the wildcard characters '*' (matches zero or many characters) and '?' (matches at least one character).",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"kinds": schema.ListAttribute{
													Description:         "Kinds is a list of resource kinds.",
													MarkdownDescription: "Kinds is a list of resource kinds.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name is the name of the resource. The name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character). NOTE: 'Name' is being deprecated in favor of 'Names'.",
													MarkdownDescription: "Name is the name of the resource. The name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character). NOTE: 'Name' is being deprecated in favor of 'Names'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"names": schema.ListAttribute{
													Description:         "Names are the names of the resources. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",
													MarkdownDescription: "Names are the names of the resources. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespace_selector": schema.SingleNestedAttribute{
													Description:         "NamespaceSelector is a label selector for the resource namespace. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character).Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",
													MarkdownDescription: "NamespaceSelector is a label selector for the resource namespace. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character).Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",
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

												"namespaces": schema.ListAttribute{
													Description:         "Namespaces is a list of namespaces names. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",
													MarkdownDescription: "Namespaces is a list of namespaces names. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"operations": schema.ListAttribute{
													Description:         "Operations can contain values ['CREATE, 'UPDATE', 'CONNECT', 'DELETE'], which are used to match a specific action.",
													MarkdownDescription: "Operations can contain values ['CREATE, 'UPDATE', 'CONNECT', 'DELETE'], which are used to match a specific action.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"selector": schema.SingleNestedAttribute{
													Description:         "Selector is a label selector. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character). Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",
													MarkdownDescription: "Selector is a label selector. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character). Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",
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
										},

										"roles": schema.ListAttribute{
											Description:         "Roles is the list of namespaced role names for the user.",
											MarkdownDescription: "Roles is the list of namespaced role names for the user.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"subjects": schema.ListNestedAttribute{
											Description:         "Subjects is the list of subject names like users, user groups, and service accounts.",
											MarkdownDescription: "Subjects is the list of subject names like users, user groups, and service accounts.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"api_group": schema.StringAttribute{
														Description:         "APIGroup holds the API group of the referenced subject. Defaults to '' for ServiceAccount subjects. Defaults to 'rbac.authorization.k8s.io' for User and Group subjects.",
														MarkdownDescription: "APIGroup holds the API group of the referenced subject. Defaults to '' for ServiceAccount subjects. Defaults to 'rbac.authorization.k8s.io' for User and Group subjects.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"kind": schema.StringAttribute{
														Description:         "Kind of object being referenced. Values defined by this API group are 'User', 'Group', and 'ServiceAccount'. If the Authorizer does not recognized the kind value, the Authorizer should report an error.",
														MarkdownDescription: "Kind of object being referenced. Values defined by this API group are 'User', 'Group', and 'ServiceAccount'. If the Authorizer does not recognized the kind value, the Authorizer should report an error.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the object being referenced.",
														MarkdownDescription: "Name of the object being referenced.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the referenced object.  If the object kind is non-namespace, such as 'User' or 'Group', and this value is not empty the Authorizer should report an error.",
														MarkdownDescription: "Namespace of the referenced object.  If the object kind is non-namespace, such as 'User' or 'Group', and this value is not empty the Authorizer should report an error.",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"any": schema.ListNestedAttribute{
								Description:         "Any allows specifying resources which will be ORed",
								MarkdownDescription: "Any allows specifying resources which will be ORed",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"cluster_roles": schema.ListAttribute{
											Description:         "ClusterRoles is the list of cluster-wide role names for the user.",
											MarkdownDescription: "ClusterRoles is the list of cluster-wide role names for the user.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"resources": schema.SingleNestedAttribute{
											Description:         "ResourceDescription contains information about the resource being created or modified.",
											MarkdownDescription: "ResourceDescription contains information about the resource being created or modified.",
											Attributes: map[string]schema.Attribute{
												"annotations": schema.MapAttribute{
													Description:         "Annotations is a  map of annotations (key-value pairs of type string). Annotation keys and values support the wildcard characters '*' (matches zero or many characters) and '?' (matches at least one character).",
													MarkdownDescription: "Annotations is a  map of annotations (key-value pairs of type string). Annotation keys and values support the wildcard characters '*' (matches zero or many characters) and '?' (matches at least one character).",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"kinds": schema.ListAttribute{
													Description:         "Kinds is a list of resource kinds.",
													MarkdownDescription: "Kinds is a list of resource kinds.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name is the name of the resource. The name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character). NOTE: 'Name' is being deprecated in favor of 'Names'.",
													MarkdownDescription: "Name is the name of the resource. The name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character). NOTE: 'Name' is being deprecated in favor of 'Names'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"names": schema.ListAttribute{
													Description:         "Names are the names of the resources. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",
													MarkdownDescription: "Names are the names of the resources. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespace_selector": schema.SingleNestedAttribute{
													Description:         "NamespaceSelector is a label selector for the resource namespace. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character).Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",
													MarkdownDescription: "NamespaceSelector is a label selector for the resource namespace. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character).Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",
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

												"namespaces": schema.ListAttribute{
													Description:         "Namespaces is a list of namespaces names. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",
													MarkdownDescription: "Namespaces is a list of namespaces names. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"operations": schema.ListAttribute{
													Description:         "Operations can contain values ['CREATE, 'UPDATE', 'CONNECT', 'DELETE'], which are used to match a specific action.",
													MarkdownDescription: "Operations can contain values ['CREATE, 'UPDATE', 'CONNECT', 'DELETE'], which are used to match a specific action.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"selector": schema.SingleNestedAttribute{
													Description:         "Selector is a label selector. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character). Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",
													MarkdownDescription: "Selector is a label selector. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character). Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",
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
										},

										"roles": schema.ListAttribute{
											Description:         "Roles is the list of namespaced role names for the user.",
											MarkdownDescription: "Roles is the list of namespaced role names for the user.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"subjects": schema.ListNestedAttribute{
											Description:         "Subjects is the list of subject names like users, user groups, and service accounts.",
											MarkdownDescription: "Subjects is the list of subject names like users, user groups, and service accounts.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"api_group": schema.StringAttribute{
														Description:         "APIGroup holds the API group of the referenced subject. Defaults to '' for ServiceAccount subjects. Defaults to 'rbac.authorization.k8s.io' for User and Group subjects.",
														MarkdownDescription: "APIGroup holds the API group of the referenced subject. Defaults to '' for ServiceAccount subjects. Defaults to 'rbac.authorization.k8s.io' for User and Group subjects.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"kind": schema.StringAttribute{
														Description:         "Kind of object being referenced. Values defined by this API group are 'User', 'Group', and 'ServiceAccount'. If the Authorizer does not recognized the kind value, the Authorizer should report an error.",
														MarkdownDescription: "Kind of object being referenced. Values defined by this API group are 'User', 'Group', and 'ServiceAccount'. If the Authorizer does not recognized the kind value, the Authorizer should report an error.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the object being referenced.",
														MarkdownDescription: "Name of the object being referenced.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the referenced object.  If the object kind is non-namespace, such as 'User' or 'Group', and this value is not empty the Authorizer should report an error.",
														MarkdownDescription: "Namespace of the referenced object.  If the object kind is non-namespace, such as 'User' or 'Group', and this value is not empty the Authorizer should report an error.",
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

					"match": schema.SingleNestedAttribute{
						Description:         "MatchResources defines when cleanuppolicy should be applied. The match criteria can include resource information (e.g. kind, name, namespace, labels) and admission review request information like the user name or role. At least one kind is required.",
						MarkdownDescription: "MatchResources defines when cleanuppolicy should be applied. The match criteria can include resource information (e.g. kind, name, namespace, labels) and admission review request information like the user name or role. At least one kind is required.",
						Attributes: map[string]schema.Attribute{
							"all": schema.ListNestedAttribute{
								Description:         "All allows specifying resources which will be ANDed",
								MarkdownDescription: "All allows specifying resources which will be ANDed",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"cluster_roles": schema.ListAttribute{
											Description:         "ClusterRoles is the list of cluster-wide role names for the user.",
											MarkdownDescription: "ClusterRoles is the list of cluster-wide role names for the user.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"resources": schema.SingleNestedAttribute{
											Description:         "ResourceDescription contains information about the resource being created or modified.",
											MarkdownDescription: "ResourceDescription contains information about the resource being created or modified.",
											Attributes: map[string]schema.Attribute{
												"annotations": schema.MapAttribute{
													Description:         "Annotations is a  map of annotations (key-value pairs of type string). Annotation keys and values support the wildcard characters '*' (matches zero or many characters) and '?' (matches at least one character).",
													MarkdownDescription: "Annotations is a  map of annotations (key-value pairs of type string). Annotation keys and values support the wildcard characters '*' (matches zero or many characters) and '?' (matches at least one character).",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"kinds": schema.ListAttribute{
													Description:         "Kinds is a list of resource kinds.",
													MarkdownDescription: "Kinds is a list of resource kinds.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name is the name of the resource. The name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character). NOTE: 'Name' is being deprecated in favor of 'Names'.",
													MarkdownDescription: "Name is the name of the resource. The name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character). NOTE: 'Name' is being deprecated in favor of 'Names'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"names": schema.ListAttribute{
													Description:         "Names are the names of the resources. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",
													MarkdownDescription: "Names are the names of the resources. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespace_selector": schema.SingleNestedAttribute{
													Description:         "NamespaceSelector is a label selector for the resource namespace. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character).Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",
													MarkdownDescription: "NamespaceSelector is a label selector for the resource namespace. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character).Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",
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

												"namespaces": schema.ListAttribute{
													Description:         "Namespaces is a list of namespaces names. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",
													MarkdownDescription: "Namespaces is a list of namespaces names. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"operations": schema.ListAttribute{
													Description:         "Operations can contain values ['CREATE, 'UPDATE', 'CONNECT', 'DELETE'], which are used to match a specific action.",
													MarkdownDescription: "Operations can contain values ['CREATE, 'UPDATE', 'CONNECT', 'DELETE'], which are used to match a specific action.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"selector": schema.SingleNestedAttribute{
													Description:         "Selector is a label selector. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character). Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",
													MarkdownDescription: "Selector is a label selector. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character). Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",
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
										},

										"roles": schema.ListAttribute{
											Description:         "Roles is the list of namespaced role names for the user.",
											MarkdownDescription: "Roles is the list of namespaced role names for the user.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"subjects": schema.ListNestedAttribute{
											Description:         "Subjects is the list of subject names like users, user groups, and service accounts.",
											MarkdownDescription: "Subjects is the list of subject names like users, user groups, and service accounts.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"api_group": schema.StringAttribute{
														Description:         "APIGroup holds the API group of the referenced subject. Defaults to '' for ServiceAccount subjects. Defaults to 'rbac.authorization.k8s.io' for User and Group subjects.",
														MarkdownDescription: "APIGroup holds the API group of the referenced subject. Defaults to '' for ServiceAccount subjects. Defaults to 'rbac.authorization.k8s.io' for User and Group subjects.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"kind": schema.StringAttribute{
														Description:         "Kind of object being referenced. Values defined by this API group are 'User', 'Group', and 'ServiceAccount'. If the Authorizer does not recognized the kind value, the Authorizer should report an error.",
														MarkdownDescription: "Kind of object being referenced. Values defined by this API group are 'User', 'Group', and 'ServiceAccount'. If the Authorizer does not recognized the kind value, the Authorizer should report an error.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the object being referenced.",
														MarkdownDescription: "Name of the object being referenced.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the referenced object.  If the object kind is non-namespace, such as 'User' or 'Group', and this value is not empty the Authorizer should report an error.",
														MarkdownDescription: "Namespace of the referenced object.  If the object kind is non-namespace, such as 'User' or 'Group', and this value is not empty the Authorizer should report an error.",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"any": schema.ListNestedAttribute{
								Description:         "Any allows specifying resources which will be ORed",
								MarkdownDescription: "Any allows specifying resources which will be ORed",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"cluster_roles": schema.ListAttribute{
											Description:         "ClusterRoles is the list of cluster-wide role names for the user.",
											MarkdownDescription: "ClusterRoles is the list of cluster-wide role names for the user.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"resources": schema.SingleNestedAttribute{
											Description:         "ResourceDescription contains information about the resource being created or modified.",
											MarkdownDescription: "ResourceDescription contains information about the resource being created or modified.",
											Attributes: map[string]schema.Attribute{
												"annotations": schema.MapAttribute{
													Description:         "Annotations is a  map of annotations (key-value pairs of type string). Annotation keys and values support the wildcard characters '*' (matches zero or many characters) and '?' (matches at least one character).",
													MarkdownDescription: "Annotations is a  map of annotations (key-value pairs of type string). Annotation keys and values support the wildcard characters '*' (matches zero or many characters) and '?' (matches at least one character).",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"kinds": schema.ListAttribute{
													Description:         "Kinds is a list of resource kinds.",
													MarkdownDescription: "Kinds is a list of resource kinds.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name is the name of the resource. The name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character). NOTE: 'Name' is being deprecated in favor of 'Names'.",
													MarkdownDescription: "Name is the name of the resource. The name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character). NOTE: 'Name' is being deprecated in favor of 'Names'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"names": schema.ListAttribute{
													Description:         "Names are the names of the resources. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",
													MarkdownDescription: "Names are the names of the resources. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespace_selector": schema.SingleNestedAttribute{
													Description:         "NamespaceSelector is a label selector for the resource namespace. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character).Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",
													MarkdownDescription: "NamespaceSelector is a label selector for the resource namespace. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character).Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",
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

												"namespaces": schema.ListAttribute{
													Description:         "Namespaces is a list of namespaces names. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",
													MarkdownDescription: "Namespaces is a list of namespaces names. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"operations": schema.ListAttribute{
													Description:         "Operations can contain values ['CREATE, 'UPDATE', 'CONNECT', 'DELETE'], which are used to match a specific action.",
													MarkdownDescription: "Operations can contain values ['CREATE, 'UPDATE', 'CONNECT', 'DELETE'], which are used to match a specific action.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"selector": schema.SingleNestedAttribute{
													Description:         "Selector is a label selector. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character). Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",
													MarkdownDescription: "Selector is a label selector. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character). Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",
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
										},

										"roles": schema.ListAttribute{
											Description:         "Roles is the list of namespaced role names for the user.",
											MarkdownDescription: "Roles is the list of namespaced role names for the user.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"subjects": schema.ListNestedAttribute{
											Description:         "Subjects is the list of subject names like users, user groups, and service accounts.",
											MarkdownDescription: "Subjects is the list of subject names like users, user groups, and service accounts.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"api_group": schema.StringAttribute{
														Description:         "APIGroup holds the API group of the referenced subject. Defaults to '' for ServiceAccount subjects. Defaults to 'rbac.authorization.k8s.io' for User and Group subjects.",
														MarkdownDescription: "APIGroup holds the API group of the referenced subject. Defaults to '' for ServiceAccount subjects. Defaults to 'rbac.authorization.k8s.io' for User and Group subjects.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"kind": schema.StringAttribute{
														Description:         "Kind of object being referenced. Values defined by this API group are 'User', 'Group', and 'ServiceAccount'. If the Authorizer does not recognized the kind value, the Authorizer should report an error.",
														MarkdownDescription: "Kind of object being referenced. Values defined by this API group are 'User', 'Group', and 'ServiceAccount'. If the Authorizer does not recognized the kind value, the Authorizer should report an error.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the object being referenced.",
														MarkdownDescription: "Name of the object being referenced.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the referenced object.  If the object kind is non-namespace, such as 'User' or 'Group', and this value is not empty the Authorizer should report an error.",
														MarkdownDescription: "Namespace of the referenced object.  If the object kind is non-namespace, such as 'User' or 'Group', and this value is not empty the Authorizer should report an error.",
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

					"schedule": schema.StringAttribute{
						Description:         "The schedule in Cron format",
						MarkdownDescription: "The schedule in Cron format",
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
	}
}

func (r *KyvernoIoCleanupPolicyV2Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kyverno_io_cleanup_policy_v2alpha1_manifest")

	var model KyvernoIoCleanupPolicyV2Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("kyverno.io/v2alpha1")
	model.Kind = pointer.String("CleanupPolicy")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
