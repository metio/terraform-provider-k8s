/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package getambassador_io_v3alpha1

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
	_ datasource.DataSource              = &GetambassadorIoLogServiceV3Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &GetambassadorIoLogServiceV3Alpha1DataSource{}
)

func NewGetambassadorIoLogServiceV3Alpha1DataSource() datasource.DataSource {
	return &GetambassadorIoLogServiceV3Alpha1DataSource{}
}

type GetambassadorIoLogServiceV3Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type GetambassadorIoLogServiceV3Alpha1DataSourceData struct {
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
		Ambassador_id *[]string `tfsdk:"ambassador_id" json:"ambassador_id,omitempty"`
		Driver        *string   `tfsdk:"driver" json:"driver,omitempty"`
		Driver_config *struct {
			Additional_log_headers *[]struct {
				During_request  *bool   `tfsdk:"during_request" json:"during_request,omitempty"`
				During_response *bool   `tfsdk:"during_response" json:"during_response,omitempty"`
				During_trailer  *bool   `tfsdk:"during_trailer" json:"during_trailer,omitempty"`
				Header_name     *string `tfsdk:"header_name" json:"header_name,omitempty"`
			} `tfsdk:"additional_log_headers" json:"additional_log_headers,omitempty"`
		} `tfsdk:"driver_config" json:"driver_config,omitempty"`
		Flush_interval_byte_size *int64  `tfsdk:"flush_interval_byte_size" json:"flush_interval_byte_size,omitempty"`
		Flush_interval_time      *int64  `tfsdk:"flush_interval_time" json:"flush_interval_time,omitempty"`
		Grpc                     *bool   `tfsdk:"grpc" json:"grpc,omitempty"`
		Protocol_version         *string `tfsdk:"protocol_version" json:"protocol_version,omitempty"`
		Service                  *string `tfsdk:"service" json:"service,omitempty"`
		Stats_name               *string `tfsdk:"stats_name" json:"stats_name,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *GetambassadorIoLogServiceV3Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_getambassador_io_log_service_v3alpha1"
}

func (r *GetambassadorIoLogServiceV3Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "LogService is the Schema for the logservices API",
		MarkdownDescription: "LogService is the Schema for the logservices API",
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
				Description:         "LogServiceSpec defines the desired state of LogService",
				MarkdownDescription: "LogServiceSpec defines the desired state of LogService",
				Attributes: map[string]schema.Attribute{
					"ambassador_id": schema.ListAttribute{
						Description:         "AmbassadorID declares which Ambassador instances should pay attention to this resource. If no value is provided, the default is:  ambassador_id: - 'default'",
						MarkdownDescription: "AmbassadorID declares which Ambassador instances should pay attention to this resource. If no value is provided, the default is:  ambassador_id: - 'default'",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"driver": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"driver_config": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"additional_log_headers": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"during_request": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"during_response": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"during_trailer": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"header_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
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

					"flush_interval_byte_size": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"flush_interval_time": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"grpc": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"protocol_version": schema.StringAttribute{
						Description:         "ProtocolVersion is the envoy api transport protocol version",
						MarkdownDescription: "ProtocolVersion is the envoy api transport protocol version",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"service": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"stats_name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
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

func (r *GetambassadorIoLogServiceV3Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *GetambassadorIoLogServiceV3Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_getambassador_io_log_service_v3alpha1")

	var data GetambassadorIoLogServiceV3Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "getambassador.io", Version: "v3alpha1", Resource: "logservices"}).
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

	var readResponse GetambassadorIoLogServiceV3Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("getambassador.io/v3alpha1")
	data.Kind = pointer.String("LogService")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
