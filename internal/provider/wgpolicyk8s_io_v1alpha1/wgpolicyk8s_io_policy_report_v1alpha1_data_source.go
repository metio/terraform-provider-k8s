/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package wgpolicyk8s_io_v1alpha1

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
	_ datasource.DataSource              = &Wgpolicyk8SIoPolicyReportV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &Wgpolicyk8SIoPolicyReportV1Alpha1DataSource{}
)

func NewWgpolicyk8SIoPolicyReportV1Alpha1DataSource() datasource.DataSource {
	return &Wgpolicyk8SIoPolicyReportV1Alpha1DataSource{}
}

type Wgpolicyk8SIoPolicyReportV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type Wgpolicyk8SIoPolicyReportV1Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Results *[]struct {
		Category         *string            `tfsdk:"category" json:"category,omitempty"`
		Data             *map[string]string `tfsdk:"data" json:"data,omitempty"`
		Message          *string            `tfsdk:"message" json:"message,omitempty"`
		Policy           *string            `tfsdk:"policy" json:"policy,omitempty"`
		ResourceSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"resource_selector" json:"resourceSelector,omitempty"`
		Resources *[]struct {
			ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
			Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
			Name            *string `tfsdk:"name" json:"name,omitempty"`
			Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
			ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
			Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
		} `tfsdk:"resources" json:"resources,omitempty"`
		Rule     *string `tfsdk:"rule" json:"rule,omitempty"`
		Scored   *bool   `tfsdk:"scored" json:"scored,omitempty"`
		Severity *string `tfsdk:"severity" json:"severity,omitempty"`
		Status   *string `tfsdk:"status" json:"status,omitempty"`
	} `tfsdk:"results" json:"results,omitempty"`
	Scope *struct {
		ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
		FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
		Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
		Name            *string `tfsdk:"name" json:"name,omitempty"`
		Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
		ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
		Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
	} `tfsdk:"scope" json:"scope,omitempty"`
	ScopeSelector *struct {
		MatchExpressions *[]struct {
			Key      *string   `tfsdk:"key" json:"key,omitempty"`
			Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
			Values   *[]string `tfsdk:"values" json:"values,omitempty"`
		} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
		MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
	} `tfsdk:"scope_selector" json:"scopeSelector,omitempty"`
	Summary *struct {
		Error *int64 `tfsdk:"error" json:"error,omitempty"`
		Fail  *int64 `tfsdk:"fail" json:"fail,omitempty"`
		Pass  *int64 `tfsdk:"pass" json:"pass,omitempty"`
		Skip  *int64 `tfsdk:"skip" json:"skip,omitempty"`
		Warn  *int64 `tfsdk:"warn" json:"warn,omitempty"`
	} `tfsdk:"summary" json:"summary,omitempty"`
}

func (r *Wgpolicyk8SIoPolicyReportV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_wgpolicyk8s_io_policy_report_v1alpha1"
}

func (r *Wgpolicyk8SIoPolicyReportV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "PolicyReport is the Schema for the policyreports API",
		MarkdownDescription: "PolicyReport is the Schema for the policyreports API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"api_version": schema.StringAttribute{
				Description:         "The API group of the requested resource.",
				MarkdownDescription: "The API group of the requested resource.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"kind": schema.StringAttribute{
				Description:         "The type of the requested resource.",
				MarkdownDescription: "The type of the requested resource.",
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

			"results": schema.ListNestedAttribute{
				Description:         "PolicyReportResult provides result details",
				MarkdownDescription: "PolicyReportResult provides result details",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"category": schema.StringAttribute{
							Description:         "Category indicates policy category",
							MarkdownDescription: "Category indicates policy category",
							Required:            false,
							Optional:            false,
							Computed:            true,
						},

						"data": schema.MapAttribute{
							Description:         "Data provides additional information for the policy rule",
							MarkdownDescription: "Data provides additional information for the policy rule",
							ElementType:         types.StringType,
							Required:            false,
							Optional:            false,
							Computed:            true,
						},

						"message": schema.StringAttribute{
							Description:         "Message is a short user friendly description of the policy rule",
							MarkdownDescription: "Message is a short user friendly description of the policy rule",
							Required:            false,
							Optional:            false,
							Computed:            true,
						},

						"policy": schema.StringAttribute{
							Description:         "Policy is the name of the policy",
							MarkdownDescription: "Policy is the name of the policy",
							Required:            false,
							Optional:            false,
							Computed:            true,
						},

						"resource_selector": schema.SingleNestedAttribute{
							Description:         "ResourceSelector is an optional selector for policy results that apply to multiple resources. For example, a policy result may apply to all pods that match a label. Either a Resource or a ResourceSelector can be specified. If neither are provided, the result is assumed to be for the policy report scope.",
							MarkdownDescription: "ResourceSelector is an optional selector for policy results that apply to multiple resources. For example, a policy result may apply to all pods that match a label. Either a Resource or a ResourceSelector can be specified. If neither are provided, the result is assumed to be for the policy report scope.",
							Attributes: map[string]schema.Attribute{
								"match_expressions": schema.ListNestedAttribute{
									Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
									MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "key is the label key that the selector applies to.",
												MarkdownDescription: "key is the label key that the selector applies to.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"operator": schema.StringAttribute{
												Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
												MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"values": schema.ListAttribute{
												Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
												MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
												ElementType:         types.StringType,
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

								"match_labels": schema.MapAttribute{
									Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
									MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

						"resources": schema.ListNestedAttribute{
							Description:         "Resources is an optional reference to the resource checked by the policy and rule",
							MarkdownDescription: "Resources is an optional reference to the resource checked by the policy and rule",
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"api_version": schema.StringAttribute{
										Description:         "API version of the referent.",
										MarkdownDescription: "API version of the referent.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"field_path": schema.StringAttribute{
										Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
										MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"kind": schema.StringAttribute{
										Description:         "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
										MarkdownDescription: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"name": schema.StringAttribute{
										Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"namespace": schema.StringAttribute{
										Description:         "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
										MarkdownDescription: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"resource_version": schema.StringAttribute{
										Description:         "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
										MarkdownDescription: "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"uid": schema.StringAttribute{
										Description:         "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
										MarkdownDescription: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
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

						"rule": schema.StringAttribute{
							Description:         "Rule is the name of the policy rule",
							MarkdownDescription: "Rule is the name of the policy rule",
							Required:            false,
							Optional:            false,
							Computed:            true,
						},

						"scored": schema.BoolAttribute{
							Description:         "Scored indicates if this policy rule is scored",
							MarkdownDescription: "Scored indicates if this policy rule is scored",
							Required:            false,
							Optional:            false,
							Computed:            true,
						},

						"severity": schema.StringAttribute{
							Description:         "Severity indicates policy severity",
							MarkdownDescription: "Severity indicates policy severity",
							Required:            false,
							Optional:            false,
							Computed:            true,
						},

						"status": schema.StringAttribute{
							Description:         "Status indicates the result of the policy rule check",
							MarkdownDescription: "Status indicates the result of the policy rule check",
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

			"scope": schema.SingleNestedAttribute{
				Description:         "Scope is an optional reference to the report scope (e.g. a Deployment, Namespace, or Node)",
				MarkdownDescription: "Scope is an optional reference to the report scope (e.g. a Deployment, Namespace, or Node)",
				Attributes: map[string]schema.Attribute{
					"api_version": schema.StringAttribute{
						Description:         "API version of the referent.",
						MarkdownDescription: "API version of the referent.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"field_path": schema.StringAttribute{
						Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
						MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"kind": schema.StringAttribute{
						Description:         "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
						MarkdownDescription: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"name": schema.StringAttribute{
						Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
						MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"namespace": schema.StringAttribute{
						Description:         "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
						MarkdownDescription: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"resource_version": schema.StringAttribute{
						Description:         "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
						MarkdownDescription: "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"uid": schema.StringAttribute{
						Description:         "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
						MarkdownDescription: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},

			"scope_selector": schema.SingleNestedAttribute{
				Description:         "ScopeSelector is an optional selector for multiple scopes (e.g. Pods). Either one of, or none of, but not both of, Scope or ScopeSelector should be specified.",
				MarkdownDescription: "ScopeSelector is an optional selector for multiple scopes (e.g. Pods). Either one of, or none of, but not both of, Scope or ScopeSelector should be specified.",
				Attributes: map[string]schema.Attribute{
					"match_expressions": schema.ListNestedAttribute{
						Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
						MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"key": schema.StringAttribute{
									Description:         "key is the label key that the selector applies to.",
									MarkdownDescription: "key is the label key that the selector applies to.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"operator": schema.StringAttribute{
									Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
									MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"values": schema.ListAttribute{
									Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
									MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
									ElementType:         types.StringType,
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

					"match_labels": schema.MapAttribute{
						Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
						MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

			"summary": schema.SingleNestedAttribute{
				Description:         "PolicyReportSummary provides a summary of results",
				MarkdownDescription: "PolicyReportSummary provides a summary of results",
				Attributes: map[string]schema.Attribute{
					"error": schema.Int64Attribute{
						Description:         "Error provides the count of policies that could not be evaluated",
						MarkdownDescription: "Error provides the count of policies that could not be evaluated",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"fail": schema.Int64Attribute{
						Description:         "Fail provides the count of policies whose requirements were not met",
						MarkdownDescription: "Fail provides the count of policies whose requirements were not met",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"pass": schema.Int64Attribute{
						Description:         "Pass provides the count of policies whose requirements were met",
						MarkdownDescription: "Pass provides the count of policies whose requirements were met",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"skip": schema.Int64Attribute{
						Description:         "Skip indicates the count of policies that were not selected for evaluation",
						MarkdownDescription: "Skip indicates the count of policies that were not selected for evaluation",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"warn": schema.Int64Attribute{
						Description:         "Warn provides the count of unscored policies whose requirements were not met",
						MarkdownDescription: "Warn provides the count of unscored policies whose requirements were not met",
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

func (r *Wgpolicyk8SIoPolicyReportV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedDataSourceDataError(request.ProviderData))
	}
}

func (r *Wgpolicyk8SIoPolicyReportV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_wgpolicyk8s_io_policy_report_v1alpha1")

	var data Wgpolicyk8SIoPolicyReportV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "wgpolicyk8s.io", Version: "v1alpha1", Resource: "policyreports"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetNamespacedResourceError(err, data.Metadata.Name, data.Metadata.Namespace))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse Wgpolicyk8SIoPolicyReportV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("wgpolicyk8s.io/v1alpha1")
	data.Kind = pointer.String("PolicyReport")
	data.Metadata = readResponse.Metadata
	data.Results = readResponse.Results
	data.Scope = readResponse.Scope
	data.ScopeSelector = readResponse.ScopeSelector
	data.Summary = readResponse.Summary

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
