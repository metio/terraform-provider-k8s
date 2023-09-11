/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kuma_io_v1alpha1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"net/http"
)

var (
	_ datasource.DataSource              = &KumaIoMeshTraceV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &KumaIoMeshTraceV1Alpha1DataSource{}
)

func NewKumaIoMeshTraceV1Alpha1DataSource() datasource.DataSource {
	return &KumaIoMeshTraceV1Alpha1DataSource{}
}

type KumaIoMeshTraceV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type KumaIoMeshTraceV1Alpha1DataSourceData struct {
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
		Default *struct {
			Backends *[]struct {
				Datadog *struct {
					SplitService *bool   `tfsdk:"split_service" json:"splitService,omitempty"`
					Url          *string `tfsdk:"url" json:"url,omitempty"`
				} `tfsdk:"datadog" json:"datadog,omitempty"`
				OpenTelemetry *struct {
					Endpoint *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
				} `tfsdk:"open_telemetry" json:"openTelemetry,omitempty"`
				Type   *string `tfsdk:"type" json:"type,omitempty"`
				Zipkin *struct {
					ApiVersion        *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
					SharedSpanContext *bool   `tfsdk:"shared_span_context" json:"sharedSpanContext,omitempty"`
					TraceId128bit     *bool   `tfsdk:"trace_id128bit" json:"traceId128bit,omitempty"`
					Url               *string `tfsdk:"url" json:"url,omitempty"`
				} `tfsdk:"zipkin" json:"zipkin,omitempty"`
			} `tfsdk:"backends" json:"backends,omitempty"`
			Sampling *struct {
				Client  *string `tfsdk:"client" json:"client,omitempty"`
				Overall *string `tfsdk:"overall" json:"overall,omitempty"`
				Random  *string `tfsdk:"random" json:"random,omitempty"`
			} `tfsdk:"sampling" json:"sampling,omitempty"`
			Tags *[]struct {
				Header *struct {
					Default *string `tfsdk:"default" json:"default,omitempty"`
					Name    *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"header" json:"header,omitempty"`
				Literal *string `tfsdk:"literal" json:"literal,omitempty"`
				Name    *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"tags" json:"tags,omitempty"`
		} `tfsdk:"default" json:"default,omitempty"`
		TargetRef *struct {
			Kind *string            `tfsdk:"kind" json:"kind,omitempty"`
			Mesh *string            `tfsdk:"mesh" json:"mesh,omitempty"`
			Name *string            `tfsdk:"name" json:"name,omitempty"`
			Tags *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
		} `tfsdk:"target_ref" json:"targetRef,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KumaIoMeshTraceV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kuma_io_mesh_trace_v1alpha1"
}

func (r *KumaIoMeshTraceV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
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
				Description:         "Spec is the specification of the Kuma MeshTrace resource.",
				MarkdownDescription: "Spec is the specification of the Kuma MeshTrace resource.",
				Attributes: map[string]schema.Attribute{
					"default": schema.SingleNestedAttribute{
						Description:         "MeshTrace configuration.",
						MarkdownDescription: "MeshTrace configuration.",
						Attributes: map[string]schema.Attribute{
							"backends": schema.ListNestedAttribute{
								Description:         "A one element array of backend definition. Envoy allows configuring only 1 backend, so the natural way of representing that would be just one object. Unfortunately due to the reasons explained in MADR 009-tracing-policy this has to be a one element array for now.",
								MarkdownDescription: "A one element array of backend definition. Envoy allows configuring only 1 backend, so the natural way of representing that would be just one object. Unfortunately due to the reasons explained in MADR 009-tracing-policy this has to be a one element array for now.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"datadog": schema.SingleNestedAttribute{
											Description:         "Datadog backend configuration.",
											MarkdownDescription: "Datadog backend configuration.",
											Attributes: map[string]schema.Attribute{
												"split_service": schema.BoolAttribute{
													Description:         "Determines if datadog service name should be split based on traffic direction and destination. For example, with 'splitService: true' and a 'backend' service that communicates with a couple of databases, you would get service names like 'backend_INBOUND', 'backend_OUTBOUND_db1', and 'backend_OUTBOUND_db2' in Datadog. Default: false",
													MarkdownDescription: "Determines if datadog service name should be split based on traffic direction and destination. For example, with 'splitService: true' and a 'backend' service that communicates with a couple of databases, you would get service names like 'backend_INBOUND', 'backend_OUTBOUND_db1', and 'backend_OUTBOUND_db2' in Datadog. Default: false",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"url": schema.StringAttribute{
													Description:         "Address of Datadog collector, only host and port are allowed (no paths, fragments etc.)",
													MarkdownDescription: "Address of Datadog collector, only host and port are allowed (no paths, fragments etc.)",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"open_telemetry": schema.SingleNestedAttribute{
											Description:         "OpenTelemetry backend configuration.",
											MarkdownDescription: "OpenTelemetry backend configuration.",
											Attributes: map[string]schema.Attribute{
												"endpoint": schema.StringAttribute{
													Description:         "Address of OpenTelemetry collector.",
													MarkdownDescription: "Address of OpenTelemetry collector.",
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
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"zipkin": schema.SingleNestedAttribute{
											Description:         "Zipkin backend configuration.",
											MarkdownDescription: "Zipkin backend configuration.",
											Attributes: map[string]schema.Attribute{
												"api_version": schema.StringAttribute{
													Description:         "Version of the API. values: httpJson, httpProto. Default: httpJson see https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/trace/v3/zipkin.proto#L66",
													MarkdownDescription: "Version of the API. values: httpJson, httpProto. Default: httpJson see https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/trace/v3/zipkin.proto#L66",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"shared_span_context": schema.BoolAttribute{
													Description:         "Determines whether client and server spans will share the same span context. Default: true. https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/trace/v3/zipkin.proto#L63",
													MarkdownDescription: "Determines whether client and server spans will share the same span context. Default: true. https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/trace/v3/zipkin.proto#L63",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"trace_id128bit": schema.BoolAttribute{
													Description:         "Generate 128bit traces. Default: false",
													MarkdownDescription: "Generate 128bit traces. Default: false",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"url": schema.StringAttribute{
													Description:         "Address of Zipkin collector.",
													MarkdownDescription: "Address of Zipkin collector.",
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
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"sampling": schema.SingleNestedAttribute{
								Description:         "Sampling configuration. Sampling is the process by which a decision is made on whether to process/export a span or not.",
								MarkdownDescription: "Sampling configuration. Sampling is the process by which a decision is made on whether to process/export a span or not.",
								Attributes: map[string]schema.Attribute{
									"client": schema.StringAttribute{
										Description:         "Target percentage of requests that will be force traced if the 'x-client-trace-id' header is set. Default: 100% Mirror of client_sampling in Envoy https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/filter/network/http_connection_manager/v2/http_connection_manager.proto#L127-L133 Either int or decimal represented as string.",
										MarkdownDescription: "Target percentage of requests that will be force traced if the 'x-client-trace-id' header is set. Default: 100% Mirror of client_sampling in Envoy https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/filter/network/http_connection_manager/v2/http_connection_manager.proto#L127-L133 Either int or decimal represented as string.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"overall": schema.StringAttribute{
										Description:         "Target percentage of requests will be traced after all other sampling checks have been applied (client, force tracing, random sampling). This field functions as an upper limit on the total configured sampling rate. For instance, setting client_sampling to 100% but overall_sampling to 1% will result in only 1% of client requests with the appropriate headers to be force traced. Default: 100% Mirror of overall_sampling in Envoy https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/filter/network/http_connection_manager/v2/http_connection_manager.proto#L142-L150 Either int or decimal represented as string.",
										MarkdownDescription: "Target percentage of requests will be traced after all other sampling checks have been applied (client, force tracing, random sampling). This field functions as an upper limit on the total configured sampling rate. For instance, setting client_sampling to 100% but overall_sampling to 1% will result in only 1% of client requests with the appropriate headers to be force traced. Default: 100% Mirror of overall_sampling in Envoy https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/filter/network/http_connection_manager/v2/http_connection_manager.proto#L142-L150 Either int or decimal represented as string.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"random": schema.StringAttribute{
										Description:         "Target percentage of requests that will be randomly selected for trace generation, if not requested by the client or not forced. Default: 100% Mirror of random_sampling in Envoy https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/filter/network/http_connection_manager/v2/http_connection_manager.proto#L135-L140 Either int or decimal represented as string.",
										MarkdownDescription: "Target percentage of requests that will be randomly selected for trace generation, if not requested by the client or not forced. Default: 100% Mirror of random_sampling in Envoy https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/filter/network/http_connection_manager/v2/http_connection_manager.proto#L135-L140 Either int or decimal represented as string.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"tags": schema.ListNestedAttribute{
								Description:         "Custom tags configuration. You can add custom tags to traces based on headers or literal values.",
								MarkdownDescription: "Custom tags configuration. You can add custom tags to traces based on headers or literal values.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"header": schema.SingleNestedAttribute{
											Description:         "Tag taken from a header.",
											MarkdownDescription: "Tag taken from a header.",
											Attributes: map[string]schema.Attribute{
												"default": schema.StringAttribute{
													Description:         "Default value to use if header is missing. If the default is missing and there is no value the tag will not be included.",
													MarkdownDescription: "Default value to use if header is missing. If the default is missing and there is no value the tag will not be included.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the header.",
													MarkdownDescription: "Name of the header.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"literal": schema.StringAttribute{
											Description:         "Tag taken from literal value.",
											MarkdownDescription: "Tag taken from literal value.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"name": schema.StringAttribute{
											Description:         "Name of the tag.",
											MarkdownDescription: "Name of the tag.",
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

					"target_ref": schema.SingleNestedAttribute{
						Description:         "TargetRef is a reference to the resource the policy takes an effect on. The resource could be either a real store object or virtual resource defined inplace.",
						MarkdownDescription: "TargetRef is a reference to the resource the policy takes an effect on. The resource could be either a real store object or virtual resource defined inplace.",
						Attributes: map[string]schema.Attribute{
							"kind": schema.StringAttribute{
								Description:         "Kind of the referenced resource",
								MarkdownDescription: "Kind of the referenced resource",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"mesh": schema.StringAttribute{
								Description:         "Mesh is reserved for future use to identify cross mesh resources.",
								MarkdownDescription: "Mesh is reserved for future use to identify cross mesh resources.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referenced resource. Can only be used with kinds: 'MeshService', 'MeshServiceSubset' and 'MeshGatewayRoute'",
								MarkdownDescription: "Name of the referenced resource. Can only be used with kinds: 'MeshService', 'MeshServiceSubset' and 'MeshGatewayRoute'",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"tags": schema.MapAttribute{
								Description:         "Tags used to select a subset of proxies by tags. Can only be used with kinds 'MeshSubset' and 'MeshServiceSubset'",
								MarkdownDescription: "Tags used to select a subset of proxies by tags. Can only be used with kinds 'MeshSubset' and 'MeshServiceSubset'",
								ElementType:         types.StringType,
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
		},
	}
}

func (r *KumaIoMeshTraceV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *KumaIoMeshTraceV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_kuma_io_mesh_trace_v1alpha1")

	var data KumaIoMeshTraceV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "kuma.io", Version: "v1alpha1", Resource: "meshtraces"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		var statusError *k8sErrors.StatusError
		if errors.As(err, &statusError) {
			if statusError.Status().Code == http.StatusNotFound {
				response.Diagnostics.AddError(
					"Unable to find resource",
					fmt.Sprintf("The requested resource cannot be found. "+
						"Make sure that it does exist in your cluster and you have set the correct name and namespace configured.\n\n"+
						"Namespace: %s\n"+
						"Name: %s", data.Metadata.Namespace, data.Metadata.Name),
				)
				return
			}
		} else {
			response.Diagnostics.AddError(
				"Unable to GET resource",
				fmt.Sprintf("An unexpected error occurred while reading the resource. "+
					"Please report this issue to the provider developers.\n\n"+
					"GET Error (%T): %s", err, err.Error()),
			)
		}
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

	var readResponse KumaIoMeshTraceV1Alpha1DataSourceData
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

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("kuma.io/v1alpha1")
	data.Kind = pointer.String("MeshTrace")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
