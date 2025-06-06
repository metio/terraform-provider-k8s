/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package jetstream_nats_io_v1beta2

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
	_ datasource.DataSource = &JetstreamNatsIoConsumerV1Beta2Manifest{}
)

func NewJetstreamNatsIoConsumerV1Beta2Manifest() datasource.DataSource {
	return &JetstreamNatsIoConsumerV1Beta2Manifest{}
}

type JetstreamNatsIoConsumerV1Beta2Manifest struct{}

type JetstreamNatsIoConsumerV1Beta2ManifestData struct {
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
		Account            *string            `tfsdk:"account" json:"account,omitempty"`
		AckPolicy          *string            `tfsdk:"ack_policy" json:"ackPolicy,omitempty"`
		AckWait            *string            `tfsdk:"ack_wait" json:"ackWait,omitempty"`
		Backoff            *[]string          `tfsdk:"backoff" json:"backoff,omitempty"`
		Creds              *string            `tfsdk:"creds" json:"creds,omitempty"`
		DeliverGroup       *string            `tfsdk:"deliver_group" json:"deliverGroup,omitempty"`
		DeliverPolicy      *string            `tfsdk:"deliver_policy" json:"deliverPolicy,omitempty"`
		DeliverSubject     *string            `tfsdk:"deliver_subject" json:"deliverSubject,omitempty"`
		Description        *string            `tfsdk:"description" json:"description,omitempty"`
		DurableName        *string            `tfsdk:"durable_name" json:"durableName,omitempty"`
		FilterSubject      *string            `tfsdk:"filter_subject" json:"filterSubject,omitempty"`
		FilterSubjects     *[]string          `tfsdk:"filter_subjects" json:"filterSubjects,omitempty"`
		FlowControl        *bool              `tfsdk:"flow_control" json:"flowControl,omitempty"`
		HeadersOnly        *bool              `tfsdk:"headers_only" json:"headersOnly,omitempty"`
		HeartbeatInterval  *string            `tfsdk:"heartbeat_interval" json:"heartbeatInterval,omitempty"`
		InactiveThreshold  *string            `tfsdk:"inactive_threshold" json:"inactiveThreshold,omitempty"`
		JsDomain           *string            `tfsdk:"js_domain" json:"jsDomain,omitempty"`
		MaxAckPending      *int64             `tfsdk:"max_ack_pending" json:"maxAckPending,omitempty"`
		MaxDeliver         *int64             `tfsdk:"max_deliver" json:"maxDeliver,omitempty"`
		MaxRequestBatch    *int64             `tfsdk:"max_request_batch" json:"maxRequestBatch,omitempty"`
		MaxRequestExpires  *string            `tfsdk:"max_request_expires" json:"maxRequestExpires,omitempty"`
		MaxRequestMaxBytes *int64             `tfsdk:"max_request_max_bytes" json:"maxRequestMaxBytes,omitempty"`
		MaxWaiting         *int64             `tfsdk:"max_waiting" json:"maxWaiting,omitempty"`
		MemStorage         *bool              `tfsdk:"mem_storage" json:"memStorage,omitempty"`
		Metadata           *map[string]string `tfsdk:"metadata" json:"metadata,omitempty"`
		Nkey               *string            `tfsdk:"nkey" json:"nkey,omitempty"`
		OptStartSeq        *int64             `tfsdk:"opt_start_seq" json:"optStartSeq,omitempty"`
		OptStartTime       *string            `tfsdk:"opt_start_time" json:"optStartTime,omitempty"`
		PreventDelete      *bool              `tfsdk:"prevent_delete" json:"preventDelete,omitempty"`
		PreventUpdate      *bool              `tfsdk:"prevent_update" json:"preventUpdate,omitempty"`
		RateLimitBps       *int64             `tfsdk:"rate_limit_bps" json:"rateLimitBps,omitempty"`
		ReplayPolicy       *string            `tfsdk:"replay_policy" json:"replayPolicy,omitempty"`
		Replicas           *int64             `tfsdk:"replicas" json:"replicas,omitempty"`
		SampleFreq         *string            `tfsdk:"sample_freq" json:"sampleFreq,omitempty"`
		Servers            *[]string          `tfsdk:"servers" json:"servers,omitempty"`
		StreamName         *string            `tfsdk:"stream_name" json:"streamName,omitempty"`
		Tls                *struct {
			ClientCert *string   `tfsdk:"client_cert" json:"clientCert,omitempty"`
			ClientKey  *string   `tfsdk:"client_key" json:"clientKey,omitempty"`
			RootCas    *[]string `tfsdk:"root_cas" json:"rootCas,omitempty"`
		} `tfsdk:"tls" json:"tls,omitempty"`
		TlsFirst *bool `tfsdk:"tls_first" json:"tlsFirst,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *JetstreamNatsIoConsumerV1Beta2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_jetstream_nats_io_consumer_v1beta2_manifest"
}

func (r *JetstreamNatsIoConsumerV1Beta2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"account": schema.StringAttribute{
						Description:         "Name of the account to which the Consumer belongs.",
						MarkdownDescription: "Name of the account to which the Consumer belongs.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^[^.*>]*$`), ""),
						},
					},

					"ack_policy": schema.StringAttribute{
						Description:         "How messages should be acknowledged.",
						MarkdownDescription: "How messages should be acknowledged.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("none", "all", "explicit"),
						},
					},

					"ack_wait": schema.StringAttribute{
						Description:         "How long to allow messages to remain un-acknowledged before attempting redelivery.",
						MarkdownDescription: "How long to allow messages to remain un-acknowledged before attempting redelivery.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"backoff": schema.ListAttribute{
						Description:         "List of durations representing a retry time scale for NaK'd or retried messages.",
						MarkdownDescription: "List of durations representing a retry time scale for NaK'd or retried messages.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"creds": schema.StringAttribute{
						Description:         "NATS user credentials for connecting to servers. Please make sure your controller has mounted the creds on its path.",
						MarkdownDescription: "NATS user credentials for connecting to servers. Please make sure your controller has mounted the creds on its path.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"deliver_group": schema.StringAttribute{
						Description:         "The name of a queue group.",
						MarkdownDescription: "The name of a queue group.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"deliver_policy": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("all", "last", "new", "byStartSequence", "byStartTime"),
						},
					},

					"deliver_subject": schema.StringAttribute{
						Description:         "The subject to deliver observed messages, when not set, a pull-based Consumer is created.",
						MarkdownDescription: "The subject to deliver observed messages, when not set, a pull-based Consumer is created.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"description": schema.StringAttribute{
						Description:         "The description of the consumer.",
						MarkdownDescription: "The description of the consumer.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"durable_name": schema.StringAttribute{
						Description:         "The name of the Consumer.",
						MarkdownDescription: "The name of the Consumer.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
							stringvalidator.RegexMatches(regexp.MustCompile(`^[^.*>]+$`), ""),
						},
					},

					"filter_subject": schema.StringAttribute{
						Description:         "Select only a specific incoming subjects, supports wildcards.",
						MarkdownDescription: "Select only a specific incoming subjects, supports wildcards.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"filter_subjects": schema.ListAttribute{
						Description:         "List of incoming subjects, supports wildcards. Available since 2.10.",
						MarkdownDescription: "List of incoming subjects, supports wildcards. Available since 2.10.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"flow_control": schema.BoolAttribute{
						Description:         "Enables flow control.",
						MarkdownDescription: "Enables flow control.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"headers_only": schema.BoolAttribute{
						Description:         "When set, only the headers of messages in the stream are delivered, and not the bodies. Additionally, Nats-Msg-Size header is added to indicate the size of the removed payload.",
						MarkdownDescription: "When set, only the headers of messages in the stream are delivered, and not the bodies. Additionally, Nats-Msg-Size header is added to indicate the size of the removed payload.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"heartbeat_interval": schema.StringAttribute{
						Description:         "The interval used to deliver idle heartbeats for push-based consumers, in Go's time.Duration format.",
						MarkdownDescription: "The interval used to deliver idle heartbeats for push-based consumers, in Go's time.Duration format.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"inactive_threshold": schema.StringAttribute{
						Description:         "The idle time an Ephemeral Consumer allows before it is removed.",
						MarkdownDescription: "The idle time an Ephemeral Consumer allows before it is removed.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"js_domain": schema.StringAttribute{
						Description:         "The JetStream domain to use for the consumer.",
						MarkdownDescription: "The JetStream domain to use for the consumer.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"max_ack_pending": schema.Int64Attribute{
						Description:         "Maximum pending Acks before consumers are paused.",
						MarkdownDescription: "Maximum pending Acks before consumers are paused.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"max_deliver": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(-1),
						},
					},

					"max_request_batch": schema.Int64Attribute{
						Description:         "The largest batch property that may be specified when doing a pull on a Pull Consumer.",
						MarkdownDescription: "The largest batch property that may be specified when doing a pull on a Pull Consumer.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"max_request_expires": schema.StringAttribute{
						Description:         "The maximum expires duration that may be set when doing a pull on a Pull Consumer.",
						MarkdownDescription: "The maximum expires duration that may be set when doing a pull on a Pull Consumer.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"max_request_max_bytes": schema.Int64Attribute{
						Description:         "The maximum max_bytes value that maybe set when dong a pull on a Pull Consumer.",
						MarkdownDescription: "The maximum max_bytes value that maybe set when dong a pull on a Pull Consumer.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"max_waiting": schema.Int64Attribute{
						Description:         "The number of pulls that can be outstanding on a pull consumer, pulls received after this is reached are ignored.",
						MarkdownDescription: "The number of pulls that can be outstanding on a pull consumer, pulls received after this is reached are ignored.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"mem_storage": schema.BoolAttribute{
						Description:         "Force the consumer state to be kept in memory rather than inherit the setting from the stream.",
						MarkdownDescription: "Force the consumer state to be kept in memory rather than inherit the setting from the stream.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"metadata": schema.MapAttribute{
						Description:         "Additional Consumer metadata.",
						MarkdownDescription: "Additional Consumer metadata.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"nkey": schema.StringAttribute{
						Description:         "NATS user NKey for connecting to servers.",
						MarkdownDescription: "NATS user NKey for connecting to servers.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"opt_start_seq": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},

					"opt_start_time": schema.StringAttribute{
						Description:         "Time format must be RFC3339.",
						MarkdownDescription: "Time format must be RFC3339.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"prevent_delete": schema.BoolAttribute{
						Description:         "When true, the managed Consumer will not be deleted when the resource is deleted.",
						MarkdownDescription: "When true, the managed Consumer will not be deleted when the resource is deleted.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"prevent_update": schema.BoolAttribute{
						Description:         "When true, the managed Consumer will not be updated when the resource is updated.",
						MarkdownDescription: "When true, the managed Consumer will not be updated when the resource is updated.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"rate_limit_bps": schema.Int64Attribute{
						Description:         "Rate at which messages will be delivered to clients, expressed in bit per second.",
						MarkdownDescription: "Rate at which messages will be delivered to clients, expressed in bit per second.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"replay_policy": schema.StringAttribute{
						Description:         "How messages are sent.",
						MarkdownDescription: "How messages are sent.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("instant", "original"),
						},
					},

					"replicas": schema.Int64Attribute{
						Description:         "When set do not inherit the replica count from the stream but specifically set it to this amount.",
						MarkdownDescription: "When set do not inherit the replica count from the stream but specifically set it to this amount.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"sample_freq": schema.StringAttribute{
						Description:         "What percentage of acknowledgements should be samples for observability.",
						MarkdownDescription: "What percentage of acknowledgements should be samples for observability.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"servers": schema.ListAttribute{
						Description:         "A list of servers for creating consumer.",
						MarkdownDescription: "A list of servers for creating consumer.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"stream_name": schema.StringAttribute{
						Description:         "The name of the Stream to create the Consumer in.",
						MarkdownDescription: "The name of the Stream to create the Consumer in.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tls": schema.SingleNestedAttribute{
						Description:         "A client's TLS certs and keys.",
						MarkdownDescription: "A client's TLS certs and keys.",
						Attributes: map[string]schema.Attribute{
							"client_cert": schema.StringAttribute{
								Description:         "A client's cert filepath. Should be mounted.",
								MarkdownDescription: "A client's cert filepath. Should be mounted.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"client_key": schema.StringAttribute{
								Description:         "A client's key filepath. Should be mounted.",
								MarkdownDescription: "A client's key filepath. Should be mounted.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"root_cas": schema.ListAttribute{
								Description:         "A list of filepaths to CAs. Should be mounted.",
								MarkdownDescription: "A list of filepaths to CAs. Should be mounted.",
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

					"tls_first": schema.BoolAttribute{
						Description:         "When true, the KV Store will initiate TLS before server INFO.",
						MarkdownDescription: "When true, the KV Store will initiate TLS before server INFO.",
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
	}
}

func (r *JetstreamNatsIoConsumerV1Beta2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_jetstream_nats_io_consumer_v1beta2_manifest")

	var model JetstreamNatsIoConsumerV1Beta2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("jetstream.nats.io/v1beta2")
	model.Kind = pointer.String("Consumer")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
