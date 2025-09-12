/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package jetstream_nats_io_v1beta1

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
	_ datasource.DataSource = &JetstreamNatsIoStreamTemplateV1Beta1Manifest{}
)

func NewJetstreamNatsIoStreamTemplateV1Beta1Manifest() datasource.DataSource {
	return &JetstreamNatsIoStreamTemplateV1Beta1Manifest{}
}

type JetstreamNatsIoStreamTemplateV1Beta1Manifest struct{}

type JetstreamNatsIoStreamTemplateV1Beta1ManifestData struct {
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
		Discard         *string   `tfsdk:"discard" json:"discard,omitempty"`
		DuplicateWindow *string   `tfsdk:"duplicate_window" json:"duplicateWindow,omitempty"`
		MaxAge          *string   `tfsdk:"max_age" json:"maxAge,omitempty"`
		MaxBytes        *int64    `tfsdk:"max_bytes" json:"maxBytes,omitempty"`
		MaxConsumers    *int64    `tfsdk:"max_consumers" json:"maxConsumers,omitempty"`
		MaxMsgSize      *int64    `tfsdk:"max_msg_size" json:"maxMsgSize,omitempty"`
		MaxMsgs         *int64    `tfsdk:"max_msgs" json:"maxMsgs,omitempty"`
		MaxStreams      *int64    `tfsdk:"max_streams" json:"maxStreams,omitempty"`
		Name            *string   `tfsdk:"name" json:"name,omitempty"`
		NoAck           *bool     `tfsdk:"no_ack" json:"noAck,omitempty"`
		Replicas        *int64    `tfsdk:"replicas" json:"replicas,omitempty"`
		Retention       *string   `tfsdk:"retention" json:"retention,omitempty"`
		Storage         *string   `tfsdk:"storage" json:"storage,omitempty"`
		Subjects        *[]string `tfsdk:"subjects" json:"subjects,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *JetstreamNatsIoStreamTemplateV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_jetstream_nats_io_stream_template_v1beta1_manifest"
}

func (r *JetstreamNatsIoStreamTemplateV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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

					"duplicate_window": schema.StringAttribute{
						Description:         "The duration window to track duplicate messages for.",
						MarkdownDescription: "The duration window to track duplicate messages for.",
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

					"max_streams": schema.Int64Attribute{
						Description:         "The maximum number of Streams this Template can create, -1 for unlimited.",
						MarkdownDescription: "The maximum number of Streams this Template can create, -1 for unlimited.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(-1),
						},
					},

					"name": schema.StringAttribute{
						Description:         "A unique name for the Stream Template.",
						MarkdownDescription: "A unique name for the Stream Template.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
							stringvalidator.RegexMatches(regexp.MustCompile(`^[^.*>]*$`), ""),
						},
					},

					"no_ack": schema.BoolAttribute{
						Description:         "Disables acknowledging messages that are received by the Stream.",
						MarkdownDescription: "Disables acknowledging messages that are received by the Stream.",
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

					"subjects": schema.ListAttribute{
						Description:         "A list of subjects to consume, supports wildcards.",
						MarkdownDescription: "A list of subjects to consume, supports wildcards.",
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
	}
}

func (r *JetstreamNatsIoStreamTemplateV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_jetstream_nats_io_stream_template_v1beta1_manifest")

	var model JetstreamNatsIoStreamTemplateV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("jetstream.nats.io/v1beta1")
	model.Kind = pointer.String("StreamTemplate")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
