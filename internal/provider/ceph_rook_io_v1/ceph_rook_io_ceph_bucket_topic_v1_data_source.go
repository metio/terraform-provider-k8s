/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package ceph_rook_io_v1

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
	_ datasource.DataSource              = &CephRookIoCephBucketTopicV1DataSource{}
	_ datasource.DataSourceWithConfigure = &CephRookIoCephBucketTopicV1DataSource{}
)

func NewCephRookIoCephBucketTopicV1DataSource() datasource.DataSource {
	return &CephRookIoCephBucketTopicV1DataSource{}
}

type CephRookIoCephBucketTopicV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type CephRookIoCephBucketTopicV1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

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

func (r *CephRookIoCephBucketTopicV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_ceph_rook_io_ceph_bucket_topic_v1"
}

func (r *CephRookIoCephBucketTopicV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
										Optional:            false,
										Computed:            true,
									},

									"disable_verify_ssl": schema.BoolAttribute{
										Description:         "Indicate whether the server certificate is validated by the client or not",
										MarkdownDescription: "Indicate whether the server certificate is validated by the client or not",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"exchange": schema.StringAttribute{
										Description:         "Name of the exchange that is used to route messages based on topics",
										MarkdownDescription: "Name of the exchange that is used to route messages based on topics",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"uri": schema.StringAttribute{
										Description:         "The URI of the AMQP endpoint to push notification to",
										MarkdownDescription: "The URI of the AMQP endpoint to push notification to",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"http": schema.SingleNestedAttribute{
								Description:         "Spec of HTTP endpoint",
								MarkdownDescription: "Spec of HTTP endpoint",
								Attributes: map[string]schema.Attribute{
									"disable_verify_ssl": schema.BoolAttribute{
										Description:         "Indicate whether the server certificate is validated by the client or not",
										MarkdownDescription: "Indicate whether the server certificate is validated by the client or not",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"send_cloud_events": schema.BoolAttribute{
										Description:         "Send the notifications with the CloudEvents header: https://github.com/cloudevents/spec/blob/main/cloudevents/adapters/aws-s3.md Supported for Ceph Quincy (v17) or newer.",
										MarkdownDescription: "Send the notifications with the CloudEvents header: https://github.com/cloudevents/spec/blob/main/cloudevents/adapters/aws-s3.md Supported for Ceph Quincy (v17) or newer.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"uri": schema.StringAttribute{
										Description:         "The URI of the HTTP endpoint to push notification to",
										MarkdownDescription: "The URI of the HTTP endpoint to push notification to",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"kafka": schema.SingleNestedAttribute{
								Description:         "Spec of Kafka endpoint",
								MarkdownDescription: "Spec of Kafka endpoint",
								Attributes: map[string]schema.Attribute{
									"ack_level": schema.StringAttribute{
										Description:         "The ack level required for this topic (none/broker)",
										MarkdownDescription: "The ack level required for this topic (none/broker)",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"disable_verify_ssl": schema.BoolAttribute{
										Description:         "Indicate whether the server certificate is validated by the client or not",
										MarkdownDescription: "Indicate whether the server certificate is validated by the client or not",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"uri": schema.StringAttribute{
										Description:         "The URI of the Kafka endpoint to push notification to",
										MarkdownDescription: "The URI of the Kafka endpoint to push notification to",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"use_ssl": schema.BoolAttribute{
										Description:         "Indicate whether to use SSL when communicating with the broker",
										MarkdownDescription: "Indicate whether to use SSL when communicating with the broker",
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

					"object_store_name": schema.StringAttribute{
						Description:         "The name of the object store on which to define the topic",
						MarkdownDescription: "The name of the object store on which to define the topic",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"object_store_namespace": schema.StringAttribute{
						Description:         "The namespace of the object store on which to define the topic",
						MarkdownDescription: "The namespace of the object store on which to define the topic",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"opaque_data": schema.StringAttribute{
						Description:         "Data which is sent in each event",
						MarkdownDescription: "Data which is sent in each event",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"persistent": schema.BoolAttribute{
						Description:         "Indication whether notifications to this endpoint are persistent or not",
						MarkdownDescription: "Indication whether notifications to this endpoint are persistent or not",
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

func (r *CephRookIoCephBucketTopicV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *CephRookIoCephBucketTopicV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_ceph_rook_io_ceph_bucket_topic_v1")

	var data CephRookIoCephBucketTopicV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "ceph.rook.io", Version: "v1", Resource: "cephbuckettopics"}).
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

	var readResponse CephRookIoCephBucketTopicV1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("ceph.rook.io/v1")
	data.Kind = pointer.String("CephBucketTopic")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
