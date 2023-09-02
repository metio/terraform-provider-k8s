/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package ceph_rook_io_v1

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &CephRookIoCephBucketTopicV1Manifest{}
)

func NewCephRookIoCephBucketTopicV1Manifest() datasource.DataSource {
	return &CephRookIoCephBucketTopicV1Manifest{}
}

type CephRookIoCephBucketTopicV1Manifest struct{}

type CephRookIoCephBucketTopicV1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
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
		Endpoint *struct {
			Amqp *struct {
				AckLevel         *string `tfsdk:"ack_level" json:"ackLevel,omitempty"`
				DisableVerifySSL *bool   `tfsdk:"disable_verify_ssl" json:"disableVerifySSL,omitempty"`
				Exchange         *string `tfsdk:"exchange" json:"exchange,omitempty"`
				Uri              *string `tfsdk:"uri" json:"uri,omitempty"`
			} `tfsdk:"amqp" json:"amqp,omitempty"`
			Http *struct {
				DisableVerifySSL *bool   `tfsdk:"disable_verify_ssl" json:"disableVerifySSL,omitempty"`
				SendCloudEvents  *bool   `tfsdk:"send_cloud_events" json:"sendCloudEvents,omitempty"`
				Uri              *string `tfsdk:"uri" json:"uri,omitempty"`
			} `tfsdk:"http" json:"http,omitempty"`
			Kafka *struct {
				AckLevel         *string `tfsdk:"ack_level" json:"ackLevel,omitempty"`
				DisableVerifySSL *bool   `tfsdk:"disable_verify_ssl" json:"disableVerifySSL,omitempty"`
				Uri              *string `tfsdk:"uri" json:"uri,omitempty"`
				UseSSL           *bool   `tfsdk:"use_ssl" json:"useSSL,omitempty"`
			} `tfsdk:"kafka" json:"kafka,omitempty"`
		} `tfsdk:"endpoint" json:"endpoint,omitempty"`
		ObjectStoreName      *string `tfsdk:"object_store_name" json:"objectStoreName,omitempty"`
		ObjectStoreNamespace *string `tfsdk:"object_store_namespace" json:"objectStoreNamespace,omitempty"`
		OpaqueData           *string `tfsdk:"opaque_data" json:"opaqueData,omitempty"`
		Persistent           *bool   `tfsdk:"persistent" json:"persistent,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CephRookIoCephBucketTopicV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_ceph_rook_io_ceph_bucket_topic_v1_manifest"
}

func (r *CephRookIoCephBucketTopicV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CephBucketTopic represents a Ceph Object Topic for Bucket Notifications",
		MarkdownDescription: "CephBucketTopic represents a Ceph Object Topic for Bucket Notifications",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

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
				Description:         "BucketTopicSpec represent the spec of a Bucket Topic",
				MarkdownDescription: "BucketTopicSpec represent the spec of a Bucket Topic",
				Attributes: map[string]schema.Attribute{
					"endpoint": schema.SingleNestedAttribute{
						Description:         "Contains the endpoint spec of the topic",
						MarkdownDescription: "Contains the endpoint spec of the topic",
						Attributes: map[string]schema.Attribute{
							"amqp": schema.SingleNestedAttribute{
								Description:         "Spec of AMQP endpoint",
								MarkdownDescription: "Spec of AMQP endpoint",
								Attributes: map[string]schema.Attribute{
									"ack_level": schema.StringAttribute{
										Description:         "The ack level required for this topic (none/broker/routeable)",
										MarkdownDescription: "The ack level required for this topic (none/broker/routeable)",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("none", "broker", "routeable"),
										},
									},

									"disable_verify_ssl": schema.BoolAttribute{
										Description:         "Indicate whether the server certificate is validated by the client or not",
										MarkdownDescription: "Indicate whether the server certificate is validated by the client or not",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"exchange": schema.StringAttribute{
										Description:         "Name of the exchange that is used to route messages based on topics",
										MarkdownDescription: "Name of the exchange that is used to route messages based on topics",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(1),
										},
									},

									"uri": schema.StringAttribute{
										Description:         "The URI of the AMQP endpoint to push notification to",
										MarkdownDescription: "The URI of the AMQP endpoint to push notification to",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(1),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"http": schema.SingleNestedAttribute{
								Description:         "Spec of HTTP endpoint",
								MarkdownDescription: "Spec of HTTP endpoint",
								Attributes: map[string]schema.Attribute{
									"disable_verify_ssl": schema.BoolAttribute{
										Description:         "Indicate whether the server certificate is validated by the client or not",
										MarkdownDescription: "Indicate whether the server certificate is validated by the client or not",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"send_cloud_events": schema.BoolAttribute{
										Description:         "Send the notifications with the CloudEvents header: https://github.com/cloudevents/spec/blob/main/cloudevents/adapters/aws-s3.md Supported for Ceph Quincy (v17) or newer.",
										MarkdownDescription: "Send the notifications with the CloudEvents header: https://github.com/cloudevents/spec/blob/main/cloudevents/adapters/aws-s3.md Supported for Ceph Quincy (v17) or newer.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"uri": schema.StringAttribute{
										Description:         "The URI of the HTTP endpoint to push notification to",
										MarkdownDescription: "The URI of the HTTP endpoint to push notification to",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(1),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"kafka": schema.SingleNestedAttribute{
								Description:         "Spec of Kafka endpoint",
								MarkdownDescription: "Spec of Kafka endpoint",
								Attributes: map[string]schema.Attribute{
									"ack_level": schema.StringAttribute{
										Description:         "The ack level required for this topic (none/broker)",
										MarkdownDescription: "The ack level required for this topic (none/broker)",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("none", "broker"),
										},
									},

									"disable_verify_ssl": schema.BoolAttribute{
										Description:         "Indicate whether the server certificate is validated by the client or not",
										MarkdownDescription: "Indicate whether the server certificate is validated by the client or not",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"uri": schema.StringAttribute{
										Description:         "The URI of the Kafka endpoint to push notification to",
										MarkdownDescription: "The URI of the Kafka endpoint to push notification to",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(1),
										},
									},

									"use_ssl": schema.BoolAttribute{
										Description:         "Indicate whether to use SSL when communicating with the broker",
										MarkdownDescription: "Indicate whether to use SSL when communicating with the broker",
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
						Required: true,
						Optional: false,
						Computed: false,
					},

					"object_store_name": schema.StringAttribute{
						Description:         "The name of the object store on which to define the topic",
						MarkdownDescription: "The name of the object store on which to define the topic",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
						},
					},

					"object_store_namespace": schema.StringAttribute{
						Description:         "The namespace of the object store on which to define the topic",
						MarkdownDescription: "The namespace of the object store on which to define the topic",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
						},
					},

					"opaque_data": schema.StringAttribute{
						Description:         "Data which is sent in each event",
						MarkdownDescription: "Data which is sent in each event",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"persistent": schema.BoolAttribute{
						Description:         "Indication whether notifications to this endpoint are persistent or not",
						MarkdownDescription: "Indication whether notifications to this endpoint are persistent or not",
						Required:            false,
						Optional:            true,
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

func (r *CephRookIoCephBucketTopicV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_ceph_rook_io_ceph_bucket_topic_v1_manifest")

	var model CephRookIoCephBucketTopicV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("ceph.rook.io/v1")
	model.Kind = pointer.String("CephBucketTopic")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"YAML Error: "+err.Error(),
		)
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
