/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kuma_io_v1alpha1

import (
	"context"
	"fmt"
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
	_ datasource.DataSource = &KumaIoMeshTraceV1Alpha1Manifest{}
)

func NewKumaIoMeshTraceV1Alpha1Manifest() datasource.DataSource {
	return &KumaIoMeshTraceV1Alpha1Manifest{}
}

type KumaIoMeshTraceV1Alpha1Manifest struct{}

type KumaIoMeshTraceV1Alpha1ManifestData struct {
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
			Kind       *string            `tfsdk:"kind" json:"kind,omitempty"`
			Mesh       *string            `tfsdk:"mesh" json:"mesh,omitempty"`
			Name       *string            `tfsdk:"name" json:"name,omitempty"`
			ProxyTypes *[]string          `tfsdk:"proxy_types" json:"proxyTypes,omitempty"`
			Tags       *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
		} `tfsdk:"target_ref" json:"targetRef,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KumaIoMeshTraceV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kuma_io_mesh_trace_v1alpha1_manifest"
}

func (r *KumaIoMeshTraceV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "Spec is the specification of the Kuma MeshTrace resource.",
				MarkdownDescription: "Spec is the specification of the Kuma MeshTrace resource.",
				Attributes: map[string]schema.Attribute{
					"default": schema.SingleNestedAttribute{
						Description:         "MeshTrace configuration.",
						MarkdownDescription: "MeshTrace configuration.",
						Attributes: map[string]schema.Attribute{
							"backends": schema.ListNestedAttribute{
								Description:         "A one element array of backend definition.Envoy allows configuring only 1 backend, so the natural way ofrepresenting that would be just one object. Unfortunately due to thereasons explained in MADR 009-tracing-policy this has to be a one elementarray for now.",
								MarkdownDescription: "A one element array of backend definition.Envoy allows configuring only 1 backend, so the natural way ofrepresenting that would be just one object. Unfortunately due to thereasons explained in MADR 009-tracing-policy this has to be a one elementarray for now.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"datadog": schema.SingleNestedAttribute{
											Description:         "Datadog backend configuration.",
											MarkdownDescription: "Datadog backend configuration.",
											Attributes: map[string]schema.Attribute{
												"split_service": schema.BoolAttribute{
													Description:         "Determines if datadog service name should be split based on trafficdirection and destination. For example, with 'splitService: true' and a'backend' service that communicates with a couple of databases, you wouldget service names like 'backend_INBOUND', 'backend_OUTBOUND_db1', and'backend_OUTBOUND_db2' in Datadog.",
													MarkdownDescription: "Determines if datadog service name should be split based on trafficdirection and destination. For example, with 'splitService: true' and a'backend' service that communicates with a couple of databases, you wouldget service names like 'backend_INBOUND', 'backend_OUTBOUND_db1', and'backend_OUTBOUND_db2' in Datadog.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"url": schema.StringAttribute{
													Description:         "Address of Datadog collector, only host and port are allowed (no paths,fragments etc.)",
													MarkdownDescription: "Address of Datadog collector, only host and port are allowed (no paths,fragments etc.)",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"open_telemetry": schema.SingleNestedAttribute{
											Description:         "OpenTelemetry backend configuration.",
											MarkdownDescription: "OpenTelemetry backend configuration.",
											Attributes: map[string]schema.Attribute{
												"endpoint": schema.StringAttribute{
													Description:         "Address of OpenTelemetry collector.",
													MarkdownDescription: "Address of OpenTelemetry collector.",
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

										"type": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Zipkin", "Datadog", "OpenTelemetry"),
											},
										},

										"zipkin": schema.SingleNestedAttribute{
											Description:         "Zipkin backend configuration.",
											MarkdownDescription: "Zipkin backend configuration.",
											Attributes: map[string]schema.Attribute{
												"api_version": schema.StringAttribute{
													Description:         "Version of the API.https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/trace/v3/zipkin.proto#L66",
													MarkdownDescription: "Version of the API.https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/trace/v3/zipkin.proto#L66",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("httpJson", "httpProto"),
													},
												},

												"shared_span_context": schema.BoolAttribute{
													Description:         "Determines whether client and server spans will share the same spancontext.https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/trace/v3/zipkin.proto#L63",
													MarkdownDescription: "Determines whether client and server spans will share the same spancontext.https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/trace/v3/zipkin.proto#L63",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"trace_id128bit": schema.BoolAttribute{
													Description:         "Generate 128bit traces.",
													MarkdownDescription: "Generate 128bit traces.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"url": schema.StringAttribute{
													Description:         "Address of Zipkin collector.",
													MarkdownDescription: "Address of Zipkin collector.",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"sampling": schema.SingleNestedAttribute{
								Description:         "Sampling configuration.Sampling is the process by which a decision is made on whether toprocess/export a span or not.",
								MarkdownDescription: "Sampling configuration.Sampling is the process by which a decision is made on whether toprocess/export a span or not.",
								Attributes: map[string]schema.Attribute{
									"client": schema.StringAttribute{
										Description:         "Target percentage of requests that will be force traced if the'x-client-trace-id' header is set. Mirror of client_sampling in Envoyhttps://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/filter/network/http_connection_manager/v2/http_connection_manager.proto#L127-L133Either int or decimal represented as string.",
										MarkdownDescription: "Target percentage of requests that will be force traced if the'x-client-trace-id' header is set. Mirror of client_sampling in Envoyhttps://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/filter/network/http_connection_manager/v2/http_connection_manager.proto#L127-L133Either int or decimal represented as string.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"overall": schema.StringAttribute{
										Description:         "Target percentage of requests will be tracedafter all other sampling checks have been applied (client, force tracing,random sampling). This field functions as an upper limit on the totalconfigured sampling rate. For instance, setting client_sampling to 100%but overall_sampling to 1% will result in only 1% of client requests withthe appropriate headers to be force traced. Mirror ofoverall_sampling in Envoyhttps://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/filter/network/http_connection_manager/v2/http_connection_manager.proto#L142-L150Either int or decimal represented as string.",
										MarkdownDescription: "Target percentage of requests will be tracedafter all other sampling checks have been applied (client, force tracing,random sampling). This field functions as an upper limit on the totalconfigured sampling rate. For instance, setting client_sampling to 100%but overall_sampling to 1% will result in only 1% of client requests withthe appropriate headers to be force traced. Mirror ofoverall_sampling in Envoyhttps://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/filter/network/http_connection_manager/v2/http_connection_manager.proto#L142-L150Either int or decimal represented as string.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"random": schema.StringAttribute{
										Description:         "Target percentage of requests that will be randomly selected for tracegeneration, if not requested by the client or not forced.Mirror of random_sampling in Envoyhttps://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/filter/network/http_connection_manager/v2/http_connection_manager.proto#L135-L140Either int or decimal represented as string.",
										MarkdownDescription: "Target percentage of requests that will be randomly selected for tracegeneration, if not requested by the client or not forced.Mirror of random_sampling in Envoyhttps://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/filter/network/http_connection_manager/v2/http_connection_manager.proto#L135-L140Either int or decimal represented as string.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"tags": schema.ListNestedAttribute{
								Description:         "Custom tags configuration. You can add custom tags to traces based onheaders or literal values.",
								MarkdownDescription: "Custom tags configuration. You can add custom tags to traces based onheaders or literal values.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"header": schema.SingleNestedAttribute{
											Description:         "Tag taken from a header.",
											MarkdownDescription: "Tag taken from a header.",
											Attributes: map[string]schema.Attribute{
												"default": schema.StringAttribute{
													Description:         "Default value to use if header is missing.If the default is missing and there is no value the tag will not beincluded.",
													MarkdownDescription: "Default value to use if header is missing.If the default is missing and there is no value the tag will not beincluded.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the header.",
													MarkdownDescription: "Name of the header.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"literal": schema.StringAttribute{
											Description:         "Tag taken from literal value.",
											MarkdownDescription: "Tag taken from literal value.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name of the tag.",
											MarkdownDescription: "Name of the tag.",
											Required:            true,
											Optional:            false,
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

					"target_ref": schema.SingleNestedAttribute{
						Description:         "TargetRef is a reference to the resource the policy takes an effect on.The resource could be either a real store object or virtual resourcedefined inplace.",
						MarkdownDescription: "TargetRef is a reference to the resource the policy takes an effect on.The resource could be either a real store object or virtual resourcedefined inplace.",
						Attributes: map[string]schema.Attribute{
							"kind": schema.StringAttribute{
								Description:         "Kind of the referenced resource",
								MarkdownDescription: "Kind of the referenced resource",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Mesh", "MeshSubset", "MeshGateway", "MeshService", "MeshServiceSubset", "MeshHTTPRoute"),
								},
							},

							"mesh": schema.StringAttribute{
								Description:         "Mesh is reserved for future use to identify cross mesh resources.",
								MarkdownDescription: "Mesh is reserved for future use to identify cross mesh resources.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referenced resource. Can only be used with kinds: 'MeshService','MeshServiceSubset' and 'MeshGatewayRoute'",
								MarkdownDescription: "Name of the referenced resource. Can only be used with kinds: 'MeshService','MeshServiceSubset' and 'MeshGatewayRoute'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"proxy_types": schema.ListAttribute{
								Description:         "ProxyTypes specifies the data plane types that are subject to the policy. When not specified,all data plane types are targeted by the policy.",
								MarkdownDescription: "ProxyTypes specifies the data plane types that are subject to the policy. When not specified,all data plane types are targeted by the policy.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tags": schema.MapAttribute{
								Description:         "Tags used to select a subset of proxies by tags. Can only be used with kinds'MeshSubset' and 'MeshServiceSubset'",
								MarkdownDescription: "Tags used to select a subset of proxies by tags. Can only be used with kinds'MeshSubset' and 'MeshServiceSubset'",
								ElementType:         types.StringType,
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
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *KumaIoMeshTraceV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kuma_io_mesh_trace_v1alpha1_manifest")

	var model KumaIoMeshTraceV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("kuma.io/v1alpha1")
	model.Kind = pointer.String("MeshTrace")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
