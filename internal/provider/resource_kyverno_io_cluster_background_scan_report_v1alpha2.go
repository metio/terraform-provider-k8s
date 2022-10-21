/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

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

type KyvernoIoClusterBackgroundScanReportV1Alpha2Resource struct{}

var (
	_ resource.Resource = (*KyvernoIoClusterBackgroundScanReportV1Alpha2Resource)(nil)
)

type KyvernoIoClusterBackgroundScanReportV1Alpha2TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type KyvernoIoClusterBackgroundScanReportV1Alpha2GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		Results *[]struct {
			Category *string `tfsdk:"category" yaml:"category,omitempty"`

			Message *string `tfsdk:"message" yaml:"message,omitempty"`

			Policy *string `tfsdk:"policy" yaml:"policy,omitempty"`

			Properties *map[string]string `tfsdk:"properties" yaml:"properties,omitempty"`

			ResourceSelector *struct {
				MatchExpressions *[]struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

					Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
				} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

				MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
			} `tfsdk:"resource_selector" yaml:"resourceSelector,omitempty"`

			Resources *[]struct {
				ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

				FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`

				Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

				ResourceVersion *string `tfsdk:"resource_version" yaml:"resourceVersion,omitempty"`

				Uid *string `tfsdk:"uid" yaml:"uid,omitempty"`
			} `tfsdk:"resources" yaml:"resources,omitempty"`

			Result *string `tfsdk:"result" yaml:"result,omitempty"`

			Rule *string `tfsdk:"rule" yaml:"rule,omitempty"`

			Scored *bool `tfsdk:"scored" yaml:"scored,omitempty"`

			Severity *string `tfsdk:"severity" yaml:"severity,omitempty"`

			Source *string `tfsdk:"source" yaml:"source,omitempty"`

			Timestamp *struct {
				Nanos *int64 `tfsdk:"nanos" yaml:"nanos,omitempty"`

				Seconds *int64 `tfsdk:"seconds" yaml:"seconds,omitempty"`
			} `tfsdk:"timestamp" yaml:"timestamp,omitempty"`
		} `tfsdk:"results" yaml:"results,omitempty"`

		Summary *struct {
			Error *int64 `tfsdk:"error" yaml:"error,omitempty"`

			Fail *int64 `tfsdk:"fail" yaml:"fail,omitempty"`

			Pass *int64 `tfsdk:"pass" yaml:"pass,omitempty"`

			Skip *int64 `tfsdk:"skip" yaml:"skip,omitempty"`

			Warn *int64 `tfsdk:"warn" yaml:"warn,omitempty"`
		} `tfsdk:"summary" yaml:"summary,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewKyvernoIoClusterBackgroundScanReportV1Alpha2Resource() resource.Resource {
	return &KyvernoIoClusterBackgroundScanReportV1Alpha2Resource{}
}

func (r *KyvernoIoClusterBackgroundScanReportV1Alpha2Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_kyverno_io_cluster_background_scan_report_v1alpha2"
}

func (r *KyvernoIoClusterBackgroundScanReportV1Alpha2Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "ClusterBackgroundScanReport is the Schema for the ClusterBackgroundScanReports API",
		MarkdownDescription: "ClusterBackgroundScanReport is the Schema for the ClusterBackgroundScanReports API",
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
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
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
				Description:         "",
				MarkdownDescription: "",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"results": {
						Description:         "PolicyReportResult provides result details",
						MarkdownDescription: "PolicyReportResult provides result details",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"category": {
								Description:         "Category indicates policy category",
								MarkdownDescription: "Category indicates policy category",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"message": {
								Description:         "Description is a short user friendly message for the policy rule",
								MarkdownDescription: "Description is a short user friendly message for the policy rule",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"policy": {
								Description:         "Policy is the name or identifier of the policy",
								MarkdownDescription: "Policy is the name or identifier of the policy",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"properties": {
								Description:         "Properties provides additional information for the policy rule",
								MarkdownDescription: "Properties provides additional information for the policy rule",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"resource_selector": {
								Description:         "SubjectSelector is an optional label selector for checked Kubernetes resources. For example, a policy result may apply to all pods that match a label. Either a Subject or a SubjectSelector can be specified. If neither are provided, the result is assumed to be for the policy report scope.",
								MarkdownDescription: "SubjectSelector is an optional label selector for checked Kubernetes resources. For example, a policy result may apply to all pods that match a label. Either a Subject or a SubjectSelector can be specified. If neither are provided, the result is assumed to be for the policy report scope.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"match_expressions": {
										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "key is the label key that the selector applies to.",
												MarkdownDescription: "key is the label key that the selector applies to.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"operator": {
												Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
												MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"values": {
												Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
												MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

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

									"match_labels": {
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

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

							"resources": {
								Description:         "Subjects is an optional reference to the checked Kubernetes resources",
								MarkdownDescription: "Subjects is an optional reference to the checked Kubernetes resources",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"api_version": {
										Description:         "API version of the referent.",
										MarkdownDescription: "API version of the referent.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"field_path": {
										Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
										MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"kind": {
										Description:         "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
										MarkdownDescription: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"namespace": {
										Description:         "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
										MarkdownDescription: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"resource_version": {
										Description:         "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
										MarkdownDescription: "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"uid": {
										Description:         "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
										MarkdownDescription: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"result": {
								Description:         "Result indicates the outcome of the policy rule execution",
								MarkdownDescription: "Result indicates the outcome of the policy rule execution",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("pass", "fail", "warn", "error", "skip"),
								},
							},

							"rule": {
								Description:         "Rule is the name or identifier of the rule within the policy",
								MarkdownDescription: "Rule is the name or identifier of the rule within the policy",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"scored": {
								Description:         "Scored indicates if this result is scored",
								MarkdownDescription: "Scored indicates if this result is scored",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"severity": {
								Description:         "Severity indicates policy check result criticality",
								MarkdownDescription: "Severity indicates policy check result criticality",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("critical", "high", "low", "medium", "info"),
								},
							},

							"source": {
								Description:         "Source is an identifier for the policy engine that manages this report",
								MarkdownDescription: "Source is an identifier for the policy engine that manages this report",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"timestamp": {
								Description:         "Timestamp indicates the time the result was found",
								MarkdownDescription: "Timestamp indicates the time the result was found",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"nanos": {
										Description:         "Non-negative fractions of a second at nanosecond resolution. Negative second values with fractions must still have non-negative nanos values that count forward in time. Must be from 0 to 999,999,999 inclusive. This field may be limited in precision depending on context.",
										MarkdownDescription: "Non-negative fractions of a second at nanosecond resolution. Negative second values with fractions must still have non-negative nanos values that count forward in time. Must be from 0 to 999,999,999 inclusive. This field may be limited in precision depending on context.",

										Type: types.Int64Type,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"seconds": {
										Description:         "Represents seconds of UTC time since Unix epoch 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59Z inclusive.",
										MarkdownDescription: "Represents seconds of UTC time since Unix epoch 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59Z inclusive.",

										Type: types.Int64Type,

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

					"summary": {
						Description:         "PolicyReportSummary provides a summary of results",
						MarkdownDescription: "PolicyReportSummary provides a summary of results",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"error": {
								Description:         "Error provides the count of policies that could not be evaluated",
								MarkdownDescription: "Error provides the count of policies that could not be evaluated",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"fail": {
								Description:         "Fail provides the count of policies whose requirements were not met",
								MarkdownDescription: "Fail provides the count of policies whose requirements were not met",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pass": {
								Description:         "Pass provides the count of policies whose requirements were met",
								MarkdownDescription: "Pass provides the count of policies whose requirements were met",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"skip": {
								Description:         "Skip indicates the count of policies that were not selected for evaluation",
								MarkdownDescription: "Skip indicates the count of policies that were not selected for evaluation",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"warn": {
								Description:         "Warn provides the count of non-scored policies whose requirements were not met",
								MarkdownDescription: "Warn provides the count of non-scored policies whose requirements were not met",

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
				}),

				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}, nil
}

func (r *KyvernoIoClusterBackgroundScanReportV1Alpha2Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_kyverno_io_cluster_background_scan_report_v1alpha2")

	var state KyvernoIoClusterBackgroundScanReportV1Alpha2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel KyvernoIoClusterBackgroundScanReportV1Alpha2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("kyverno.io/v1alpha2")
	goModel.Kind = utilities.Ptr("ClusterBackgroundScanReport")

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

func (r *KyvernoIoClusterBackgroundScanReportV1Alpha2Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kyverno_io_cluster_background_scan_report_v1alpha2")
	// NO-OP: All data is already in Terraform state
}

func (r *KyvernoIoClusterBackgroundScanReportV1Alpha2Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_kyverno_io_cluster_background_scan_report_v1alpha2")

	var state KyvernoIoClusterBackgroundScanReportV1Alpha2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel KyvernoIoClusterBackgroundScanReportV1Alpha2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("kyverno.io/v1alpha2")
	goModel.Kind = utilities.Ptr("ClusterBackgroundScanReport")

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

func (r *KyvernoIoClusterBackgroundScanReportV1Alpha2Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_kyverno_io_cluster_background_scan_report_v1alpha2")
	// NO-OP: Terraform removes the state automatically for us
}
