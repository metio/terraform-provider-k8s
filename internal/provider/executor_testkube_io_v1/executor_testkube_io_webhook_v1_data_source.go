/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package executor_testkube_io_v1

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
	_ datasource.DataSource              = &ExecutorTestkubeIoWebhookV1DataSource{}
	_ datasource.DataSourceWithConfigure = &ExecutorTestkubeIoWebhookV1DataSource{}
)

func NewExecutorTestkubeIoWebhookV1DataSource() datasource.DataSource {
	return &ExecutorTestkubeIoWebhookV1DataSource{}
}

type ExecutorTestkubeIoWebhookV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type ExecutorTestkubeIoWebhookV1DataSourceData struct {
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
		Events                   *[]string          `tfsdk:"events" json:"events,omitempty"`
		Headers                  *map[string]string `tfsdk:"headers" json:"headers,omitempty"`
		PayloadObjectField       *string            `tfsdk:"payload_object_field" json:"payloadObjectField,omitempty"`
		PayloadTemplate          *string            `tfsdk:"payload_template" json:"payloadTemplate,omitempty"`
		PayloadTemplateReference *string            `tfsdk:"payload_template_reference" json:"payloadTemplateReference,omitempty"`
		Selector                 *string            `tfsdk:"selector" json:"selector,omitempty"`
		Uri                      *string            `tfsdk:"uri" json:"uri,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ExecutorTestkubeIoWebhookV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_executor_testkube_io_webhook_v1"
}

func (r *ExecutorTestkubeIoWebhookV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Webhook is the Schema for the webhooks API",
		MarkdownDescription: "Webhook is the Schema for the webhooks API",
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
				Description:         "WebhookSpec defines the desired state of Webhook",
				MarkdownDescription: "WebhookSpec defines the desired state of Webhook",
				Attributes: map[string]schema.Attribute{
					"events": schema.ListAttribute{
						Description:         "Events declare list if events on which webhook should be called",
						MarkdownDescription: "Events declare list if events on which webhook should be called",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"headers": schema.MapAttribute{
						Description:         "webhook headers (golang template supported)",
						MarkdownDescription: "webhook headers (golang template supported)",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"payload_object_field": schema.StringAttribute{
						Description:         "will load the generated payload for notification inside the object",
						MarkdownDescription: "will load the generated payload for notification inside the object",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"payload_template": schema.StringAttribute{
						Description:         "golang based template for notification payload",
						MarkdownDescription: "golang based template for notification payload",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"payload_template_reference": schema.StringAttribute{
						Description:         "name of the template resource",
						MarkdownDescription: "name of the template resource",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"selector": schema.StringAttribute{
						Description:         "Labels to filter for tests and test suites",
						MarkdownDescription: "Labels to filter for tests and test suites",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"uri": schema.StringAttribute{
						Description:         "Uri is address where webhook should be made (golang template supported)",
						MarkdownDescription: "Uri is address where webhook should be made (golang template supported)",
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

func (r *ExecutorTestkubeIoWebhookV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *ExecutorTestkubeIoWebhookV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_executor_testkube_io_webhook_v1")

	var data ExecutorTestkubeIoWebhookV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "executor.testkube.io", Version: "v1", Resource: "webhooks"}).
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

	var readResponse ExecutorTestkubeIoWebhookV1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("executor.testkube.io/v1")
	data.Kind = pointer.String("Webhook")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
