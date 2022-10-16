/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

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

type FlowcontrolApiserverK8SIoPriorityLevelConfigurationV1Beta3Resource struct{}

var (
	_ resource.Resource = (*FlowcontrolApiserverK8SIoPriorityLevelConfigurationV1Beta3Resource)(nil)
)

type FlowcontrolApiserverK8SIoPriorityLevelConfigurationV1Beta3TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type FlowcontrolApiserverK8SIoPriorityLevelConfigurationV1Beta3GoModel struct {
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
		Limited *struct {
			LimitResponse *struct {
				Queuing *struct {
					HandSize *int64 `tfsdk:"hand_size" yaml:"handSize,omitempty"`

					QueueLengthLimit *int64 `tfsdk:"queue_length_limit" yaml:"queueLengthLimit,omitempty"`

					Queues *int64 `tfsdk:"queues" yaml:"queues,omitempty"`
				} `tfsdk:"queuing" yaml:"queuing,omitempty"`

				Type *string `tfsdk:"type" yaml:"type,omitempty"`
			} `tfsdk:"limit_response" yaml:"limitResponse,omitempty"`

			NominalConcurrencyShares *int64 `tfsdk:"nominal_concurrency_shares" yaml:"nominalConcurrencyShares,omitempty"`
		} `tfsdk:"limited" yaml:"limited,omitempty"`

		Type *string `tfsdk:"type" yaml:"type,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewFlowcontrolApiserverK8SIoPriorityLevelConfigurationV1Beta3Resource() resource.Resource {
	return &FlowcontrolApiserverK8SIoPriorityLevelConfigurationV1Beta3Resource{}
}

func (r *FlowcontrolApiserverK8SIoPriorityLevelConfigurationV1Beta3Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_flowcontrol_apiserver_k8s_io_priority_level_configuration_v1beta3"
}

func (r *FlowcontrolApiserverK8SIoPriorityLevelConfigurationV1Beta3Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "PriorityLevelConfiguration represents the configuration of a priority level.",
		MarkdownDescription: "PriorityLevelConfiguration represents the configuration of a priority level.",
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
				Description:         "PriorityLevelConfigurationSpec specifies the configuration of a priority level.",
				MarkdownDescription: "PriorityLevelConfigurationSpec specifies the configuration of a priority level.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"limited": {
						Description:         "LimitedPriorityLevelConfiguration specifies how to handle requests that are subject to limits. It addresses two issues:  - How are requests for this priority level limited?  - What should be done with requests that exceed the limit?",
						MarkdownDescription: "LimitedPriorityLevelConfiguration specifies how to handle requests that are subject to limits. It addresses two issues:  - How are requests for this priority level limited?  - What should be done with requests that exceed the limit?",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"limit_response": {
								Description:         "LimitResponse defines how to handle requests that can not be executed right now.",
								MarkdownDescription: "LimitResponse defines how to handle requests that can not be executed right now.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"queuing": {
										Description:         "QueuingConfiguration holds the configuration parameters for queuing",
										MarkdownDescription: "QueuingConfiguration holds the configuration parameters for queuing",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"hand_size": {
												Description:         "'handSize' is a small positive number that configures the shuffle sharding of requests into queues.  When enqueuing a request at this priority level the request's flow identifier (a string pair) is hashed and the hash value is used to shuffle the list of queues and deal a hand of the size specified here.  The request is put into one of the shortest queues in that hand. 'handSize' must be no larger than 'queues', and should be significantly smaller (so that a few heavy flows do not saturate most of the queues).  See the user-facing documentation for more extensive guidance on setting this field.  This field has a default value of 8.",
												MarkdownDescription: "'handSize' is a small positive number that configures the shuffle sharding of requests into queues.  When enqueuing a request at this priority level the request's flow identifier (a string pair) is hashed and the hash value is used to shuffle the list of queues and deal a hand of the size specified here.  The request is put into one of the shortest queues in that hand. 'handSize' must be no larger than 'queues', and should be significantly smaller (so that a few heavy flows do not saturate most of the queues).  See the user-facing documentation for more extensive guidance on setting this field.  This field has a default value of 8.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"queue_length_limit": {
												Description:         "'queueLengthLimit' is the maximum number of requests allowed to be waiting in a given queue of this priority level at a time; excess requests are rejected.  This value must be positive.  If not specified, it will be defaulted to 50.",
												MarkdownDescription: "'queueLengthLimit' is the maximum number of requests allowed to be waiting in a given queue of this priority level at a time; excess requests are rejected.  This value must be positive.  If not specified, it will be defaulted to 50.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"queues": {
												Description:         "'queues' is the number of queues for this priority level. The queues exist independently at each apiserver. The value must be positive.  Setting it to 1 effectively precludes shufflesharding and thus makes the distinguisher method of associated flow schemas irrelevant.  This field has a default value of 64.",
												MarkdownDescription: "'queues' is the number of queues for this priority level. The queues exist independently at each apiserver. The value must be positive.  Setting it to 1 effectively precludes shufflesharding and thus makes the distinguisher method of associated flow schemas irrelevant.  This field has a default value of 64.",

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

									"type": {
										Description:         "'type' is 'Queue' or 'Reject'. 'Queue' means that requests that can not be executed upon arrival are held in a queue until they can be executed or a queuing limit is reached. 'Reject' means that requests that can not be executed upon arrival are rejected. Required.",
										MarkdownDescription: "'type' is 'Queue' or 'Reject'. 'Queue' means that requests that can not be executed upon arrival are held in a queue until they can be executed or a queuing limit is reached. 'Reject' means that requests that can not be executed upon arrival are rejected. Required.",

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

							"nominal_concurrency_shares": {
								Description:         "'nominalConcurrencyShares' (NCS) contributes to the computation of the NominalConcurrencyLimit (NominalCL) of this level. This is the number of execution seats available at this priority level. This is used both for requests dispatched from this priority level as well as requests dispatched from other priority levels borrowing seats from this level. The server's concurrency limit (ServerCL) is divided among the Limited priority levels in proportion to their NCS values:NominalCL(i)  = ceil( ServerCL * NCS(i) / sum_ncs ) sum_ncs = sum[limited priority level k] NCS(k)Bigger numbers mean a larger nominal concurrency limit, at the expense of every other Limited priority level. This field has a default value of 30.",
								MarkdownDescription: "'nominalConcurrencyShares' (NCS) contributes to the computation of the NominalConcurrencyLimit (NominalCL) of this level. This is the number of execution seats available at this priority level. This is used both for requests dispatched from this priority level as well as requests dispatched from other priority levels borrowing seats from this level. The server's concurrency limit (ServerCL) is divided among the Limited priority levels in proportion to their NCS values:NominalCL(i)  = ceil( ServerCL * NCS(i) / sum_ncs ) sum_ncs = sum[limited priority level k] NCS(k)Bigger numbers mean a larger nominal concurrency limit, at the expense of every other Limited priority level. This field has a default value of 30.",

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

					"type": {
						Description:         "'type' indicates whether this priority level is subject to limitation on request execution.  A value of ''Exempt'' means that requests of this priority level are not subject to a limit (and thus are never queued) and do not detract from the capacity made available to other priority levels.  A value of ''Limited'' means that (a) requests of this priority level _are_ subject to limits and (b) some of the server's limited capacity is made available exclusively to this priority level. Required.",
						MarkdownDescription: "'type' indicates whether this priority level is subject to limitation on request execution.  A value of ''Exempt'' means that requests of this priority level are not subject to a limit (and thus are never queued) and do not detract from the capacity made available to other priority levels.  A value of ''Limited'' means that (a) requests of this priority level _are_ subject to limits and (b) some of the server's limited capacity is made available exclusively to this priority level. Required.",

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
		},
	}, nil
}

func (r *FlowcontrolApiserverK8SIoPriorityLevelConfigurationV1Beta3Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_flowcontrol_apiserver_k8s_io_priority_level_configuration_v1beta3")

	var state FlowcontrolApiserverK8SIoPriorityLevelConfigurationV1Beta3TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel FlowcontrolApiserverK8SIoPriorityLevelConfigurationV1Beta3GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("flowcontrol.apiserver.k8s.io/v1beta3")
	goModel.Kind = utilities.Ptr("PriorityLevelConfiguration")

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

func (r *FlowcontrolApiserverK8SIoPriorityLevelConfigurationV1Beta3Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_flowcontrol_apiserver_k8s_io_priority_level_configuration_v1beta3")
	// NO-OP: All data is already in Terraform state
}

func (r *FlowcontrolApiserverK8SIoPriorityLevelConfigurationV1Beta3Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_flowcontrol_apiserver_k8s_io_priority_level_configuration_v1beta3")

	var state FlowcontrolApiserverK8SIoPriorityLevelConfigurationV1Beta3TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel FlowcontrolApiserverK8SIoPriorityLevelConfigurationV1Beta3GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("flowcontrol.apiserver.k8s.io/v1beta3")
	goModel.Kind = utilities.Ptr("PriorityLevelConfiguration")

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

func (r *FlowcontrolApiserverK8SIoPriorityLevelConfigurationV1Beta3Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_flowcontrol_apiserver_k8s_io_priority_level_configuration_v1beta3")
	// NO-OP: Terraform removes the state automatically for us
}
