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
	_ datasource.DataSource = &JetstreamNatsIoStreamV1Beta2Manifest{}
)

func NewJetstreamNatsIoStreamV1Beta2Manifest() datasource.DataSource {
	return &JetstreamNatsIoStreamV1Beta2Manifest{}
}

type JetstreamNatsIoStreamV1Beta2Manifest struct{}

type JetstreamNatsIoStreamV1Beta2ManifestData struct {
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
		Account        *string `tfsdk:"account" json:"account,omitempty"`
		AllowDirect    *bool   `tfsdk:"allow_direct" json:"allowDirect,omitempty"`
		AllowRollup    *bool   `tfsdk:"allow_rollup" json:"allowRollup,omitempty"`
		Compression    *string `tfsdk:"compression" json:"compression,omitempty"`
		ConsumerLimits *struct {
			InactiveThreshold *string `tfsdk:"inactive_threshold" json:"inactiveThreshold,omitempty"`
			MaxAckPending     *int64  `tfsdk:"max_ack_pending" json:"maxAckPending,omitempty"`
		} `tfsdk:"consumer_limits" json:"consumerLimits,omitempty"`
		Creds             *string            `tfsdk:"creds" json:"creds,omitempty"`
		DenyDelete        *bool              `tfsdk:"deny_delete" json:"denyDelete,omitempty"`
		DenyPurge         *bool              `tfsdk:"deny_purge" json:"denyPurge,omitempty"`
		Description       *string            `tfsdk:"description" json:"description,omitempty"`
		Discard           *string            `tfsdk:"discard" json:"discard,omitempty"`
		DiscardPerSubject *bool              `tfsdk:"discard_per_subject" json:"discardPerSubject,omitempty"`
		DuplicateWindow   *string            `tfsdk:"duplicate_window" json:"duplicateWindow,omitempty"`
		FirstSequence     *float64           `tfsdk:"first_sequence" json:"firstSequence,omitempty"`
		MaxAge            *string            `tfsdk:"max_age" json:"maxAge,omitempty"`
		MaxBytes          *int64             `tfsdk:"max_bytes" json:"maxBytes,omitempty"`
		MaxConsumers      *int64             `tfsdk:"max_consumers" json:"maxConsumers,omitempty"`
		MaxMsgSize        *int64             `tfsdk:"max_msg_size" json:"maxMsgSize,omitempty"`
		MaxMsgs           *int64             `tfsdk:"max_msgs" json:"maxMsgs,omitempty"`
		MaxMsgsPerSubject *int64             `tfsdk:"max_msgs_per_subject" json:"maxMsgsPerSubject,omitempty"`
		Metadata          *map[string]string `tfsdk:"metadata" json:"metadata,omitempty"`
		Mirror            *struct {
			ExternalApiPrefix     *string `tfsdk:"external_api_prefix" json:"externalApiPrefix,omitempty"`
			ExternalDeliverPrefix *string `tfsdk:"external_deliver_prefix" json:"externalDeliverPrefix,omitempty"`
			FilterSubject         *string `tfsdk:"filter_subject" json:"filterSubject,omitempty"`
			Name                  *string `tfsdk:"name" json:"name,omitempty"`
			OptStartSeq           *int64  `tfsdk:"opt_start_seq" json:"optStartSeq,omitempty"`
			OptStartTime          *string `tfsdk:"opt_start_time" json:"optStartTime,omitempty"`
			SubjectTransforms     *[]struct {
				Dest   *string `tfsdk:"dest" json:"dest,omitempty"`
				Source *string `tfsdk:"source" json:"source,omitempty"`
			} `tfsdk:"subject_transforms" json:"subjectTransforms,omitempty"`
		} `tfsdk:"mirror" json:"mirror,omitempty"`
		MirrorDirect *bool   `tfsdk:"mirror_direct" json:"mirrorDirect,omitempty"`
		Name         *string `tfsdk:"name" json:"name,omitempty"`
		Nkey         *string `tfsdk:"nkey" json:"nkey,omitempty"`
		NoAck        *bool   `tfsdk:"no_ack" json:"noAck,omitempty"`
		Placement    *struct {
			Cluster *string   `tfsdk:"cluster" json:"cluster,omitempty"`
			Tags    *[]string `tfsdk:"tags" json:"tags,omitempty"`
		} `tfsdk:"placement" json:"placement,omitempty"`
		PreventDelete *bool  `tfsdk:"prevent_delete" json:"preventDelete,omitempty"`
		PreventUpdate *bool  `tfsdk:"prevent_update" json:"preventUpdate,omitempty"`
		Replicas      *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
		Republish     *struct {
			Destination *string `tfsdk:"destination" json:"destination,omitempty"`
			Source      *string `tfsdk:"source" json:"source,omitempty"`
		} `tfsdk:"republish" json:"republish,omitempty"`
		Retention *string   `tfsdk:"retention" json:"retention,omitempty"`
		Sealed    *bool     `tfsdk:"sealed" json:"sealed,omitempty"`
		Servers   *[]string `tfsdk:"servers" json:"servers,omitempty"`
		Sources   *[]struct {
			ExternalApiPrefix     *string `tfsdk:"external_api_prefix" json:"externalApiPrefix,omitempty"`
			ExternalDeliverPrefix *string `tfsdk:"external_deliver_prefix" json:"externalDeliverPrefix,omitempty"`
			FilterSubject         *string `tfsdk:"filter_subject" json:"filterSubject,omitempty"`
			Name                  *string `tfsdk:"name" json:"name,omitempty"`
			OptStartSeq           *int64  `tfsdk:"opt_start_seq" json:"optStartSeq,omitempty"`
			OptStartTime          *string `tfsdk:"opt_start_time" json:"optStartTime,omitempty"`
			SubjectTransforms     *[]struct {
				Dest   *string `tfsdk:"dest" json:"dest,omitempty"`
				Source *string `tfsdk:"source" json:"source,omitempty"`
			} `tfsdk:"subject_transforms" json:"subjectTransforms,omitempty"`
		} `tfsdk:"sources" json:"sources,omitempty"`
		Storage          *string `tfsdk:"storage" json:"storage,omitempty"`
		SubjectTransform *struct {
			Dest   *string `tfsdk:"dest" json:"dest,omitempty"`
			Source *string `tfsdk:"source" json:"source,omitempty"`
		} `tfsdk:"subject_transform" json:"subjectTransform,omitempty"`
		Subjects *[]string `tfsdk:"subjects" json:"subjects,omitempty"`
		Tls      *struct {
			ClientCert *string   `tfsdk:"client_cert" json:"clientCert,omitempty"`
			ClientKey  *string   `tfsdk:"client_key" json:"clientKey,omitempty"`
			RootCas    *[]string `tfsdk:"root_cas" json:"rootCas,omitempty"`
		} `tfsdk:"tls" json:"tls,omitempty"`
		TlsFirst *bool `tfsdk:"tls_first" json:"tlsFirst,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *JetstreamNatsIoStreamV1Beta2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_jetstream_nats_io_stream_v1beta2_manifest"
}

func (r *JetstreamNatsIoStreamV1Beta2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
						Description:         "Name of the account to which the Stream belongs.",
						MarkdownDescription: "Name of the account to which the Stream belongs.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^[^.*>]*$`), ""),
						},
					},

					"allow_direct": schema.BoolAttribute{
						Description:         "When true, allow higher performance, direct access to get individual messages.",
						MarkdownDescription: "When true, allow higher performance, direct access to get individual messages.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"allow_rollup": schema.BoolAttribute{
						Description:         "When true, allows the use of the Nats-Rollup header to replace all contents of a stream, or subject in a stream, with a single new message.",
						MarkdownDescription: "When true, allows the use of the Nats-Rollup header to replace all contents of a stream, or subject in a stream, with a single new message.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"compression": schema.StringAttribute{
						Description:         "Stream specific compression.",
						MarkdownDescription: "Stream specific compression.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("s2", "none", ""),
						},
					},

					"consumer_limits": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"inactive_threshold": schema.StringAttribute{
								Description:         "The duration of inactivity after which a consumer is considered inactive.",
								MarkdownDescription: "The duration of inactivity after which a consumer is considered inactive.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_ack_pending": schema.Int64Attribute{
								Description:         "Maximum number of outstanding unacknowledged messages.",
								MarkdownDescription: "Maximum number of outstanding unacknowledged messages.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"creds": schema.StringAttribute{
						Description:         "NATS user credentials for connecting to servers. Please make sure your controller has mounted the creds on this path.",
						MarkdownDescription: "NATS user credentials for connecting to servers. Please make sure your controller has mounted the creds on this path.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"deny_delete": schema.BoolAttribute{
						Description:         "When true, restricts the ability to delete messages from a stream via the API. Cannot be changed once set to true.",
						MarkdownDescription: "When true, restricts the ability to delete messages from a stream via the API. Cannot be changed once set to true.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"deny_purge": schema.BoolAttribute{
						Description:         "When true, restricts the ability to purge a stream via the API. Cannot be changed once set to true.",
						MarkdownDescription: "When true, restricts the ability to purge a stream via the API. Cannot be changed once set to true.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"description": schema.StringAttribute{
						Description:         "The description of the stream.",
						MarkdownDescription: "The description of the stream.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"discard": schema.StringAttribute{
						Description:         "When a Stream reach it's limits either old messages are deleted or new ones are denied.",
						MarkdownDescription: "When a Stream reach it's limits either old messages are deleted or new ones are denied.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("old", "new"),
						},
					},

					"discard_per_subject": schema.BoolAttribute{
						Description:         "Applies discard policy on a per-subject basis. Requires discard policy 'new' and 'maxMsgs' to be set.",
						MarkdownDescription: "Applies discard policy on a per-subject basis. Requires discard policy 'new' and 'maxMsgs' to be set.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"duplicate_window": schema.StringAttribute{
						Description:         "The duration window to track duplicate messages for.",
						MarkdownDescription: "The duration window to track duplicate messages for.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"first_sequence": schema.Float64Attribute{
						Description:         "Sequence number from which the Stream will start.",
						MarkdownDescription: "Sequence number from which the Stream will start.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"max_age": schema.StringAttribute{
						Description:         "Maximum age of any message in the stream, expressed in Go's time.Duration format. Empty for unlimited.",
						MarkdownDescription: "Maximum age of any message in the stream, expressed in Go's time.Duration format. Empty for unlimited.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"max_bytes": schema.Int64Attribute{
						Description:         "How big the Stream may be, when the combined stream size exceeds this old messages are removed. -1 for unlimited.",
						MarkdownDescription: "How big the Stream may be, when the combined stream size exceeds this old messages are removed. -1 for unlimited.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(-1),
						},
					},

					"max_consumers": schema.Int64Attribute{
						Description:         "How many Consumers can be defined for a given Stream. -1 for unlimited.",
						MarkdownDescription: "How many Consumers can be defined for a given Stream. -1 for unlimited.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(-1),
						},
					},

					"max_msg_size": schema.Int64Attribute{
						Description:         "The largest message that will be accepted by the Stream. -1 for unlimited.",
						MarkdownDescription: "The largest message that will be accepted by the Stream. -1 for unlimited.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(-1),
						},
					},

					"max_msgs": schema.Int64Attribute{
						Description:         "How many messages may be in a Stream, oldest messages will be removed if the Stream exceeds this size. -1 for unlimited.",
						MarkdownDescription: "How many messages may be in a Stream, oldest messages will be removed if the Stream exceeds this size. -1 for unlimited.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(-1),
						},
					},

					"max_msgs_per_subject": schema.Int64Attribute{
						Description:         "The maximum number of messages per subject.",
						MarkdownDescription: "The maximum number of messages per subject.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"metadata": schema.MapAttribute{
						Description:         "Additional Stream metadata.",
						MarkdownDescription: "Additional Stream metadata.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"mirror": schema.SingleNestedAttribute{
						Description:         "A stream mirror.",
						MarkdownDescription: "A stream mirror.",
						Attributes: map[string]schema.Attribute{
							"external_api_prefix": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"external_deliver_prefix": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"filter_subject": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
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
							},

							"opt_start_time": schema.StringAttribute{
								Description:         "Time format must be RFC3339.",
								MarkdownDescription: "Time format must be RFC3339.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"subject_transforms": schema.ListNestedAttribute{
								Description:         "List of subject transforms for this mirror.",
								MarkdownDescription: "List of subject transforms for this mirror.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"dest": schema.StringAttribute{
											Description:         "Destination subject.",
											MarkdownDescription: "Destination subject.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"source": schema.StringAttribute{
											Description:         "Source subject.",
											MarkdownDescription: "Source subject.",
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

					"mirror_direct": schema.BoolAttribute{
						Description:         "When true, enables direct access to messages from the origin stream.",
						MarkdownDescription: "When true, enables direct access to messages from the origin stream.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"name": schema.StringAttribute{
						Description:         "A unique name for the Stream.",
						MarkdownDescription: "A unique name for the Stream.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
							stringvalidator.RegexMatches(regexp.MustCompile(`^[^.*>]*$`), ""),
						},
					},

					"nkey": schema.StringAttribute{
						Description:         "NATS user NKey for connecting to servers.",
						MarkdownDescription: "NATS user NKey for connecting to servers.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"no_ack": schema.BoolAttribute{
						Description:         "Disables acknowledging messages that are received by the Stream.",
						MarkdownDescription: "Disables acknowledging messages that are received by the Stream.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"placement": schema.SingleNestedAttribute{
						Description:         "A stream's placement.",
						MarkdownDescription: "A stream's placement.",
						Attributes: map[string]schema.Attribute{
							"cluster": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tags": schema.ListAttribute{
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

					"prevent_delete": schema.BoolAttribute{
						Description:         "When true, the managed Stream will not be deleted when the resource is deleted.",
						MarkdownDescription: "When true, the managed Stream will not be deleted when the resource is deleted.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"prevent_update": schema.BoolAttribute{
						Description:         "When true, the managed Stream will not be updated when the resource is updated.",
						MarkdownDescription: "When true, the managed Stream will not be updated when the resource is updated.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"replicas": schema.Int64Attribute{
						Description:         "How many replicas to keep for each message.",
						MarkdownDescription: "How many replicas to keep for each message.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
						},
					},

					"republish": schema.SingleNestedAttribute{
						Description:         "Republish configuration of the stream.",
						MarkdownDescription: "Republish configuration of the stream.",
						Attributes: map[string]schema.Attribute{
							"destination": schema.StringAttribute{
								Description:         "Messages will be additionally published to this subject.",
								MarkdownDescription: "Messages will be additionally published to this subject.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"source": schema.StringAttribute{
								Description:         "Messages will be published from this subject to the destination subject.",
								MarkdownDescription: "Messages will be published from this subject to the destination subject.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"retention": schema.StringAttribute{
						Description:         "How messages are retained in the Stream, once this is exceeded old messages are removed.",
						MarkdownDescription: "How messages are retained in the Stream, once this is exceeded old messages are removed.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("limits", "interest", "workqueue"),
						},
					},

					"sealed": schema.BoolAttribute{
						Description:         "Seal an existing stream so no new messages may be added.",
						MarkdownDescription: "Seal an existing stream so no new messages may be added.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"servers": schema.ListAttribute{
						Description:         "A list of servers for creating stream.",
						MarkdownDescription: "A list of servers for creating stream.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"sources": schema.ListNestedAttribute{
						Description:         "A stream's sources.",
						MarkdownDescription: "A stream's sources.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"external_api_prefix": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"external_deliver_prefix": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"filter_subject": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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
								},

								"opt_start_time": schema.StringAttribute{
									Description:         "Time format must be RFC3339.",
									MarkdownDescription: "Time format must be RFC3339.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"subject_transforms": schema.ListNestedAttribute{
									Description:         "List of subject transforms for this mirror.",
									MarkdownDescription: "List of subject transforms for this mirror.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"dest": schema.StringAttribute{
												Description:         "Destination subject.",
												MarkdownDescription: "Destination subject.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"source": schema.StringAttribute{
												Description:         "Source subject.",
												MarkdownDescription: "Source subject.",
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

					"storage": schema.StringAttribute{
						Description:         "The storage backend to use for the Stream.",
						MarkdownDescription: "The storage backend to use for the Stream.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("file", "memory"),
						},
					},

					"subject_transform": schema.SingleNestedAttribute{
						Description:         "SubjectTransform is for applying a subject transform (to matching messages) when a new message is received.",
						MarkdownDescription: "SubjectTransform is for applying a subject transform (to matching messages) when a new message is received.",
						Attributes: map[string]schema.Attribute{
							"dest": schema.StringAttribute{
								Description:         "Destination subject to transform into.",
								MarkdownDescription: "Destination subject to transform into.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"source": schema.StringAttribute{
								Description:         "Source subject.",
								MarkdownDescription: "Source subject.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"subjects": schema.ListAttribute{
						Description:         "A list of subjects to consume, supports wildcards.",
						MarkdownDescription: "A list of subjects to consume, supports wildcards.",
						ElementType:         types.StringType,
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

func (r *JetstreamNatsIoStreamV1Beta2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_jetstream_nats_io_stream_v1beta2_manifest")

	var model JetstreamNatsIoStreamV1Beta2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("jetstream.nats.io/v1beta2")
	model.Kind = pointer.String("Stream")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
