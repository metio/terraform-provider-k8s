/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package flowcontrol_apiserver_k8s_io_v1beta3

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
	_ datasource.DataSource = &FlowcontrolApiserverK8SIoPriorityLevelConfigurationV1Beta3Manifest{}
)

func NewFlowcontrolApiserverK8SIoPriorityLevelConfigurationV1Beta3Manifest() datasource.DataSource {
	return &FlowcontrolApiserverK8SIoPriorityLevelConfigurationV1Beta3Manifest{}
}

type FlowcontrolApiserverK8SIoPriorityLevelConfigurationV1Beta3Manifest struct{}

type FlowcontrolApiserverK8SIoPriorityLevelConfigurationV1Beta3ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Exempt *struct {
			LendablePercent          *int64 `tfsdk:"lendable_percent" json:"lendablePercent,omitempty"`
			NominalConcurrencyShares *int64 `tfsdk:"nominal_concurrency_shares" json:"nominalConcurrencyShares,omitempty"`
		} `tfsdk:"exempt" json:"exempt,omitempty"`
		Limited *struct {
			BorrowingLimitPercent *int64 `tfsdk:"borrowing_limit_percent" json:"borrowingLimitPercent,omitempty"`
			LendablePercent       *int64 `tfsdk:"lendable_percent" json:"lendablePercent,omitempty"`
			LimitResponse         *struct {
				Queuing *struct {
					HandSize         *int64 `tfsdk:"hand_size" json:"handSize,omitempty"`
					QueueLengthLimit *int64 `tfsdk:"queue_length_limit" json:"queueLengthLimit,omitempty"`
					Queues           *int64 `tfsdk:"queues" json:"queues,omitempty"`
				} `tfsdk:"queuing" json:"queuing,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"limit_response" json:"limitResponse,omitempty"`
			NominalConcurrencyShares *int64 `tfsdk:"nominal_concurrency_shares" json:"nominalConcurrencyShares,omitempty"`
		} `tfsdk:"limited" json:"limited,omitempty"`
		Type *string `tfsdk:"type" json:"type,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *FlowcontrolApiserverK8SIoPriorityLevelConfigurationV1Beta3Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_flowcontrol_apiserver_k8s_io_priority_level_configuration_v1beta3_manifest"
}

func (r *FlowcontrolApiserverK8SIoPriorityLevelConfigurationV1Beta3Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "PriorityLevelConfiguration represents the configuration of a priority level.",
		MarkdownDescription: "PriorityLevelConfiguration represents the configuration of a priority level.",
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
				Description:         "PriorityLevelConfigurationSpec specifies the configuration of a priority level.",
				MarkdownDescription: "PriorityLevelConfigurationSpec specifies the configuration of a priority level.",
				Attributes: map[string]schema.Attribute{
					"exempt": schema.SingleNestedAttribute{
						Description:         "ExemptPriorityLevelConfiguration describes the configurable aspects of the handling of exempt requests. In the mandatory exempt configuration object the values in the fields here can be modified by authorized users, unlike the rest of the 'spec'.",
						MarkdownDescription: "ExemptPriorityLevelConfiguration describes the configurable aspects of the handling of exempt requests. In the mandatory exempt configuration object the values in the fields here can be modified by authorized users, unlike the rest of the 'spec'.",
						Attributes: map[string]schema.Attribute{
							"lendable_percent": schema.Int64Attribute{
								Description:         "'lendablePercent' prescribes the fraction of the level's NominalCL that can be borrowed by other priority levels.  This value of this field must be between 0 and 100, inclusive, and it defaults to 0. The number of seats that other levels can borrow from this level, known as this level's LendableConcurrencyLimit (LendableCL), is defined as follows.LendableCL(i) = round( NominalCL(i) * lendablePercent(i)/100.0 )",
								MarkdownDescription: "'lendablePercent' prescribes the fraction of the level's NominalCL that can be borrowed by other priority levels.  This value of this field must be between 0 and 100, inclusive, and it defaults to 0. The number of seats that other levels can borrow from this level, known as this level's LendableConcurrencyLimit (LendableCL), is defined as follows.LendableCL(i) = round( NominalCL(i) * lendablePercent(i)/100.0 )",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"nominal_concurrency_shares": schema.Int64Attribute{
								Description:         "'nominalConcurrencyShares' (NCS) contributes to the computation of the NominalConcurrencyLimit (NominalCL) of this level. This is the number of execution seats nominally reserved for this priority level. This DOES NOT limit the dispatching from this priority level but affects the other priority levels through the borrowing mechanism. The server's concurrency limit (ServerCL) is divided among all the priority levels in proportion to their NCS values:NominalCL(i)  = ceil( ServerCL * NCS(i) / sum_ncs ) sum_ncs = sum[priority level k] NCS(k)Bigger numbers mean a larger nominal concurrency limit, at the expense of every other priority level. This field has a default value of zero.",
								MarkdownDescription: "'nominalConcurrencyShares' (NCS) contributes to the computation of the NominalConcurrencyLimit (NominalCL) of this level. This is the number of execution seats nominally reserved for this priority level. This DOES NOT limit the dispatching from this priority level but affects the other priority levels through the borrowing mechanism. The server's concurrency limit (ServerCL) is divided among all the priority levels in proportion to their NCS values:NominalCL(i)  = ceil( ServerCL * NCS(i) / sum_ncs ) sum_ncs = sum[priority level k] NCS(k)Bigger numbers mean a larger nominal concurrency limit, at the expense of every other priority level. This field has a default value of zero.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"limited": schema.SingleNestedAttribute{
						Description:         "LimitedPriorityLevelConfiguration specifies how to handle requests that are subject to limits. It addresses two issues:  - How are requests for this priority level limited?  - What should be done with requests that exceed the limit?",
						MarkdownDescription: "LimitedPriorityLevelConfiguration specifies how to handle requests that are subject to limits. It addresses two issues:  - How are requests for this priority level limited?  - What should be done with requests that exceed the limit?",
						Attributes: map[string]schema.Attribute{
							"borrowing_limit_percent": schema.Int64Attribute{
								Description:         "'borrowingLimitPercent', if present, configures a limit on how many seats this priority level can borrow from other priority levels. The limit is known as this level's BorrowingConcurrencyLimit (BorrowingCL) and is a limit on the total number of seats that this level may borrow at any one time. This field holds the ratio of that limit to the level's nominal concurrency limit. When this field is non-nil, it must hold a non-negative integer and the limit is calculated as follows.BorrowingCL(i) = round( NominalCL(i) * borrowingLimitPercent(i)/100.0 )The value of this field can be more than 100, implying that this priority level can borrow a number of seats that is greater than its own nominal concurrency limit (NominalCL). When this field is left 'nil', the limit is effectively infinite.",
								MarkdownDescription: "'borrowingLimitPercent', if present, configures a limit on how many seats this priority level can borrow from other priority levels. The limit is known as this level's BorrowingConcurrencyLimit (BorrowingCL) and is a limit on the total number of seats that this level may borrow at any one time. This field holds the ratio of that limit to the level's nominal concurrency limit. When this field is non-nil, it must hold a non-negative integer and the limit is calculated as follows.BorrowingCL(i) = round( NominalCL(i) * borrowingLimitPercent(i)/100.0 )The value of this field can be more than 100, implying that this priority level can borrow a number of seats that is greater than its own nominal concurrency limit (NominalCL). When this field is left 'nil', the limit is effectively infinite.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"lendable_percent": schema.Int64Attribute{
								Description:         "'lendablePercent' prescribes the fraction of the level's NominalCL that can be borrowed by other priority levels. The value of this field must be between 0 and 100, inclusive, and it defaults to 0. The number of seats that other levels can borrow from this level, known as this level's LendableConcurrencyLimit (LendableCL), is defined as follows.LendableCL(i) = round( NominalCL(i) * lendablePercent(i)/100.0 )",
								MarkdownDescription: "'lendablePercent' prescribes the fraction of the level's NominalCL that can be borrowed by other priority levels. The value of this field must be between 0 and 100, inclusive, and it defaults to 0. The number of seats that other levels can borrow from this level, known as this level's LendableConcurrencyLimit (LendableCL), is defined as follows.LendableCL(i) = round( NominalCL(i) * lendablePercent(i)/100.0 )",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"limit_response": schema.SingleNestedAttribute{
								Description:         "LimitResponse defines how to handle requests that can not be executed right now.",
								MarkdownDescription: "LimitResponse defines how to handle requests that can not be executed right now.",
								Attributes: map[string]schema.Attribute{
									"queuing": schema.SingleNestedAttribute{
										Description:         "QueuingConfiguration holds the configuration parameters for queuing",
										MarkdownDescription: "QueuingConfiguration holds the configuration parameters for queuing",
										Attributes: map[string]schema.Attribute{
											"hand_size": schema.Int64Attribute{
												Description:         "'handSize' is a small positive number that configures the shuffle sharding of requests into queues.  When enqueuing a request at this priority level the request's flow identifier (a string pair) is hashed and the hash value is used to shuffle the list of queues and deal a hand of the size specified here.  The request is put into one of the shortest queues in that hand. 'handSize' must be no larger than 'queues', and should be significantly smaller (so that a few heavy flows do not saturate most of the queues).  See the user-facing documentation for more extensive guidance on setting this field.  This field has a default value of 8.",
												MarkdownDescription: "'handSize' is a small positive number that configures the shuffle sharding of requests into queues.  When enqueuing a request at this priority level the request's flow identifier (a string pair) is hashed and the hash value is used to shuffle the list of queues and deal a hand of the size specified here.  The request is put into one of the shortest queues in that hand. 'handSize' must be no larger than 'queues', and should be significantly smaller (so that a few heavy flows do not saturate most of the queues).  See the user-facing documentation for more extensive guidance on setting this field.  This field has a default value of 8.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"queue_length_limit": schema.Int64Attribute{
												Description:         "'queueLengthLimit' is the maximum number of requests allowed to be waiting in a given queue of this priority level at a time; excess requests are rejected.  This value must be positive.  If not specified, it will be defaulted to 50.",
												MarkdownDescription: "'queueLengthLimit' is the maximum number of requests allowed to be waiting in a given queue of this priority level at a time; excess requests are rejected.  This value must be positive.  If not specified, it will be defaulted to 50.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"queues": schema.Int64Attribute{
												Description:         "'queues' is the number of queues for this priority level. The queues exist independently at each apiserver. The value must be positive.  Setting it to 1 effectively precludes shufflesharding and thus makes the distinguisher method of associated flow schemas irrelevant.  This field has a default value of 64.",
												MarkdownDescription: "'queues' is the number of queues for this priority level. The queues exist independently at each apiserver. The value must be positive.  Setting it to 1 effectively precludes shufflesharding and thus makes the distinguisher method of associated flow schemas irrelevant.  This field has a default value of 64.",
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
										Description:         "'type' is 'Queue' or 'Reject'. 'Queue' means that requests that can not be executed upon arrival are held in a queue until they can be executed or a queuing limit is reached. 'Reject' means that requests that can not be executed upon arrival are rejected. Required.",
										MarkdownDescription: "'type' is 'Queue' or 'Reject'. 'Queue' means that requests that can not be executed upon arrival are held in a queue until they can be executed or a queuing limit is reached. 'Reject' means that requests that can not be executed upon arrival are rejected. Required.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"nominal_concurrency_shares": schema.Int64Attribute{
								Description:         "'nominalConcurrencyShares' (NCS) contributes to the computation of the NominalConcurrencyLimit (NominalCL) of this level. This is the number of execution seats available at this priority level. This is used both for requests dispatched from this priority level as well as requests dispatched from other priority levels borrowing seats from this level. The server's concurrency limit (ServerCL) is divided among the Limited priority levels in proportion to their NCS values:NominalCL(i)  = ceil( ServerCL * NCS(i) / sum_ncs ) sum_ncs = sum[priority level k] NCS(k)Bigger numbers mean a larger nominal concurrency limit, at the expense of every other priority level. This field has a default value of 30.",
								MarkdownDescription: "'nominalConcurrencyShares' (NCS) contributes to the computation of the NominalConcurrencyLimit (NominalCL) of this level. This is the number of execution seats available at this priority level. This is used both for requests dispatched from this priority level as well as requests dispatched from other priority levels borrowing seats from this level. The server's concurrency limit (ServerCL) is divided among the Limited priority levels in proportion to their NCS values:NominalCL(i)  = ceil( ServerCL * NCS(i) / sum_ncs ) sum_ncs = sum[priority level k] NCS(k)Bigger numbers mean a larger nominal concurrency limit, at the expense of every other priority level. This field has a default value of 30.",
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
						Description:         "'type' indicates whether this priority level is subject to limitation on request execution.  A value of ''Exempt'' means that requests of this priority level are not subject to a limit (and thus are never queued) and do not detract from the capacity made available to other priority levels.  A value of ''Limited'' means that (a) requests of this priority level _are_ subject to limits and (b) some of the server's limited capacity is made available exclusively to this priority level. Required.",
						MarkdownDescription: "'type' indicates whether this priority level is subject to limitation on request execution.  A value of ''Exempt'' means that requests of this priority level are not subject to a limit (and thus are never queued) and do not detract from the capacity made available to other priority levels.  A value of ''Limited'' means that (a) requests of this priority level _are_ subject to limits and (b) some of the server's limited capacity is made available exclusively to this priority level. Required.",
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
	}
}

func (r *FlowcontrolApiserverK8SIoPriorityLevelConfigurationV1Beta3Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_flowcontrol_apiserver_k8s_io_priority_level_configuration_v1beta3_manifest")

	var model FlowcontrolApiserverK8SIoPriorityLevelConfigurationV1Beta3ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("flowcontrol.apiserver.k8s.io/v1beta3")
	model.Kind = pointer.String("PriorityLevelConfiguration")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
