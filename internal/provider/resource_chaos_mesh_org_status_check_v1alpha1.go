/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"

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

type ChaosMeshOrgStatusCheckV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*ChaosMeshOrgStatusCheckV1Alpha1Resource)(nil)
)

type ChaosMeshOrgStatusCheckV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type ChaosMeshOrgStatusCheckV1Alpha1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

		FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

		Http *struct {
			Body *string `tfsdk:"body" yaml:"body,omitempty"`

			Criteria *struct {
				StatusCode *string `tfsdk:"status_code" yaml:"statusCode,omitempty"`
			} `tfsdk:"criteria" yaml:"criteria,omitempty"`

			Headers *map[string][]string `tfsdk:"headers" yaml:"headers,omitempty"`

			Method *string `tfsdk:"method" yaml:"method,omitempty"`

			Url *string `tfsdk:"url" yaml:"url,omitempty"`
		} `tfsdk:"http" yaml:"http,omitempty"`

		IntervalSeconds *int64 `tfsdk:"interval_seconds" yaml:"intervalSeconds,omitempty"`

		Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`

		RecordsHistoryLimit *int64 `tfsdk:"records_history_limit" yaml:"recordsHistoryLimit,omitempty"`

		SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

		TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`

		Type *string `tfsdk:"type" yaml:"type,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewChaosMeshOrgStatusCheckV1Alpha1Resource() resource.Resource {
	return &ChaosMeshOrgStatusCheckV1Alpha1Resource{}
}

func (r *ChaosMeshOrgStatusCheckV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_chaos_mesh_org_status_check_v1alpha1"
}

func (r *ChaosMeshOrgStatusCheckV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "",
		MarkdownDescription: "",
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
				Description:         "Spec defines the behavior of a status check",
				MarkdownDescription: "Spec defines the behavior of a status check",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"duration": {
						Description:         "Duration defines the duration of the whole status check if the number of failed execution does not exceed the failure threshold. Duration is available to both 'Synchronous' and 'Continuous' mode. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '-1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
						MarkdownDescription: "Duration defines the duration of the whole status check if the number of failed execution does not exceed the failure threshold. Duration is available to both 'Synchronous' and 'Continuous' mode. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '-1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"failure_threshold": {
						Description:         "FailureThreshold defines the minimum consecutive failure for the status check to be considered failed.",
						MarkdownDescription: "FailureThreshold defines the minimum consecutive failure for the status check to be considered failed.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							int64validator.AtLeast(1),
						},
					},

					"http": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"body": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"criteria": {
								Description:         "Criteria defines how to determine the result of the status check.",
								MarkdownDescription: "Criteria defines how to determine the result of the status check.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"status_code": {
										Description:         "StatusCode defines the expected http status code for the request. A statusCode string could be a single code (e.g. 200), or an inclusive range (e.g. 200-400, both '200' and '400' are included).",
										MarkdownDescription: "StatusCode defines the expected http status code for the request. A statusCode string could be a single code (e.g. 200), or an inclusive range (e.g. 200-400, both '200' and '400' are included).",

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

							"headers": {
								Description:         "A Header represents the key-value pairs in an HTTP header.  The keys should be in canonical form, as returned by CanonicalHeaderKey.",
								MarkdownDescription: "A Header represents the key-value pairs in an HTTP header.  The keys should be in canonical form, as returned by CanonicalHeaderKey.",

								Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"method": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("GET", "POST"),
								},
							},

							"url": {
								Description:         "",
								MarkdownDescription: "",

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

					"interval_seconds": {
						Description:         "IntervalSeconds defines how often (in seconds) to perform an execution of status check.",
						MarkdownDescription: "IntervalSeconds defines how often (in seconds) to perform an execution of status check.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							int64validator.AtLeast(1),
						},
					},

					"mode": {
						Description:         "Mode defines the execution mode of the status check. Support type: Synchronous / Continuous",
						MarkdownDescription: "Mode defines the execution mode of the status check. Support type: Synchronous / Continuous",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("Synchronous", "Continuous"),
						},
					},

					"records_history_limit": {
						Description:         "RecordsHistoryLimit defines the number of record to retain.",
						MarkdownDescription: "RecordsHistoryLimit defines the number of record to retain.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							int64validator.AtLeast(1),

							int64validator.AtMost(1000),
						},
					},

					"success_threshold": {
						Description:         "SuccessThreshold defines the minimum consecutive successes for the status check to be considered successful. SuccessThreshold only works for 'Synchronous' mode.",
						MarkdownDescription: "SuccessThreshold defines the minimum consecutive successes for the status check to be considered successful. SuccessThreshold only works for 'Synchronous' mode.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							int64validator.AtLeast(1),
						},
					},

					"timeout_seconds": {
						Description:         "TimeoutSeconds defines the number of seconds after which an execution of status check times out.",
						MarkdownDescription: "TimeoutSeconds defines the number of seconds after which an execution of status check times out.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							int64validator.AtLeast(1),
						},
					},

					"type": {
						Description:         "Type defines the specific status check type. Support type: HTTP",
						MarkdownDescription: "Type defines the specific status check type. Support type: HTTP",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("HTTP"),
						},
					},
				}),

				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}, nil
}

func (r *ChaosMeshOrgStatusCheckV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_chaos_mesh_org_status_check_v1alpha1")

	var state ChaosMeshOrgStatusCheckV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ChaosMeshOrgStatusCheckV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("chaos-mesh.org/v1alpha1")
	goModel.Kind = utilities.Ptr("StatusCheck")

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

func (r *ChaosMeshOrgStatusCheckV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_chaos_mesh_org_status_check_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *ChaosMeshOrgStatusCheckV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_chaos_mesh_org_status_check_v1alpha1")

	var state ChaosMeshOrgStatusCheckV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ChaosMeshOrgStatusCheckV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("chaos-mesh.org/v1alpha1")
	goModel.Kind = utilities.Ptr("StatusCheck")

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

func (r *ChaosMeshOrgStatusCheckV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_chaos_mesh_org_status_check_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
