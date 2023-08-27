/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package flowcontrol_apiserver_k8s_io_v1beta2

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
	_ datasource.DataSource              = &FlowcontrolApiserverK8SIoPriorityLevelConfigurationV1Beta2DataSource{}
	_ datasource.DataSourceWithConfigure = &FlowcontrolApiserverK8SIoPriorityLevelConfigurationV1Beta2DataSource{}
)

func NewFlowcontrolApiserverK8SIoPriorityLevelConfigurationV1Beta2DataSource() datasource.DataSource {
	return &FlowcontrolApiserverK8SIoPriorityLevelConfigurationV1Beta2DataSource{}
}

type FlowcontrolApiserverK8SIoPriorityLevelConfigurationV1Beta2DataSource struct {
	kubernetesClient dynamic.Interface
}

type FlowcontrolApiserverK8SIoPriorityLevelConfigurationV1Beta2DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Limited *struct {
			AssuredConcurrencyShares *int64 `tfsdk:"assured_concurrency_shares" json:"assuredConcurrencyShares,omitempty"`
			LimitResponse            *struct {
				Queuing *struct {
					HandSize         *int64 `tfsdk:"hand_size" json:"handSize,omitempty"`
					QueueLengthLimit *int64 `tfsdk:"queue_length_limit" json:"queueLengthLimit,omitempty"`
					Queues           *int64 `tfsdk:"queues" json:"queues,omitempty"`
				} `tfsdk:"queuing" json:"queuing,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"limit_response" json:"limitResponse,omitempty"`
		} `tfsdk:"limited" json:"limited,omitempty"`
		Type *string `tfsdk:"type" json:"type,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *FlowcontrolApiserverK8SIoPriorityLevelConfigurationV1Beta2DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_flowcontrol_apiserver_k8s_io_priority_level_configuration_v1beta2"
}

func (r *FlowcontrolApiserverK8SIoPriorityLevelConfigurationV1Beta2DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "PriorityLevelConfiguration represents the configuration of a priority level.",
		MarkdownDescription: "PriorityLevelConfiguration represents the configuration of a priority level.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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

			"spec": schema.SingleNestedAttribute{
				Description:         "PriorityLevelConfigurationSpec specifies the configuration of a priority level.",
				MarkdownDescription: "PriorityLevelConfigurationSpec specifies the configuration of a priority level.",
				Attributes: map[string]schema.Attribute{
					"limited": schema.SingleNestedAttribute{
						Description:         "LimitedPriorityLevelConfiguration specifies how to handle requests that are subject to limits. It addresses two issues:  - How are requests for this priority level limited?  - What should be done with requests that exceed the limit?",
						MarkdownDescription: "LimitedPriorityLevelConfiguration specifies how to handle requests that are subject to limits. It addresses two issues:  - How are requests for this priority level limited?  - What should be done with requests that exceed the limit?",
						Attributes: map[string]schema.Attribute{
							"assured_concurrency_shares": schema.Int64Attribute{
								Description:         "'assuredConcurrencyShares' (ACS) configures the execution limit, which is a limit on the number of requests of this priority level that may be exeucting at a given time.  ACS must be a positive number. The server's concurrency limit (SCL) is divided among the concurrency-controlled priority levels in proportion to their assured concurrency shares. This produces the assured concurrency value (ACV) --- the number of requests that may be executing at a time --- for each such priority level:            ACV(l) = ceil( SCL * ACS(l) / ( sum[priority levels k] ACS(k) ) )bigger numbers of ACS mean more reserved concurrent requests (at the expense of every other PL). This field has a default value of 30.",
								MarkdownDescription: "'assuredConcurrencyShares' (ACS) configures the execution limit, which is a limit on the number of requests of this priority level that may be exeucting at a given time.  ACS must be a positive number. The server's concurrency limit (SCL) is divided among the concurrency-controlled priority levels in proportion to their assured concurrency shares. This produces the assured concurrency value (ACV) --- the number of requests that may be executing at a time --- for each such priority level:            ACV(l) = ceil( SCL * ACS(l) / ( sum[priority levels k] ACS(k) ) )bigger numbers of ACS mean more reserved concurrent requests (at the expense of every other PL). This field has a default value of 30.",
								Required:            false,
								Optional:            false,
								Computed:            true,
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
												Optional:            false,
												Computed:            true,
											},

											"queue_length_limit": schema.Int64Attribute{
												Description:         "'queueLengthLimit' is the maximum number of requests allowed to be waiting in a given queue of this priority level at a time; excess requests are rejected.  This value must be positive.  If not specified, it will be defaulted to 50.",
												MarkdownDescription: "'queueLengthLimit' is the maximum number of requests allowed to be waiting in a given queue of this priority level at a time; excess requests are rejected.  This value must be positive.  If not specified, it will be defaulted to 50.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"queues": schema.Int64Attribute{
												Description:         "'queues' is the number of queues for this priority level. The queues exist independently at each apiserver. The value must be positive.  Setting it to 1 effectively precludes shufflesharding and thus makes the distinguisher method of associated flow schemas irrelevant.  This field has a default value of 64.",
												MarkdownDescription: "'queues' is the number of queues for this priority level. The queues exist independently at each apiserver. The value must be positive.  Setting it to 1 effectively precludes shufflesharding and thus makes the distinguisher method of associated flow schemas irrelevant.  This field has a default value of 64.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"type": schema.StringAttribute{
										Description:         "'type' is 'Queue' or 'Reject'. 'Queue' means that requests that can not be executed upon arrival are held in a queue until they can be executed or a queuing limit is reached. 'Reject' means that requests that can not be executed upon arrival are rejected. Required.",
										MarkdownDescription: "'type' is 'Queue' or 'Reject'. 'Queue' means that requests that can not be executed upon arrival are held in a queue until they can be executed or a queuing limit is reached. 'Reject' means that requests that can not be executed upon arrival are rejected. Required.",
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
						Required: false,
						Optional: false,
						Computed: true,
					},

					"type": schema.StringAttribute{
						Description:         "'type' indicates whether this priority level is subject to limitation on request execution.  A value of ''Exempt'' means that requests of this priority level are not subject to a limit (and thus are never queued) and do not detract from the capacity made available to other priority levels.  A value of ''Limited'' means that (a) requests of this priority level _are_ subject to limits and (b) some of the server's limited capacity is made available exclusively to this priority level. Required.",
						MarkdownDescription: "'type' indicates whether this priority level is subject to limitation on request execution.  A value of ''Exempt'' means that requests of this priority level are not subject to a limit (and thus are never queued) and do not detract from the capacity made available to other priority levels.  A value of ''Limited'' means that (a) requests of this priority level _are_ subject to limits and (b) some of the server's limited capacity is made available exclusively to this priority level. Required.",
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

func (r *FlowcontrolApiserverK8SIoPriorityLevelConfigurationV1Beta2DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *FlowcontrolApiserverK8SIoPriorityLevelConfigurationV1Beta2DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_flowcontrol_apiserver_k8s_io_priority_level_configuration_v1beta2")

	var data FlowcontrolApiserverK8SIoPriorityLevelConfigurationV1Beta2DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "flowcontrol.apiserver.k8s.io", Version: "v1beta2", Resource: "PriorityLevelConfiguration"}).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse FlowcontrolApiserverK8SIoPriorityLevelConfigurationV1Beta2DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.ID = types.StringValue(data.Metadata.Name)
	data.ApiVersion = pointer.String("flowcontrol.apiserver.k8s.io/v1beta2")
	data.Kind = pointer.String("PriorityLevelConfiguration")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
