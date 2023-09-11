/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kuma_io_v1alpha1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	k8sTypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"strings"
)

var (
	_ resource.Resource                = &KumaIoMeshTraceV1Alpha1Resource{}
	_ resource.ResourceWithConfigure   = &KumaIoMeshTraceV1Alpha1Resource{}
	_ resource.ResourceWithImportState = &KumaIoMeshTraceV1Alpha1Resource{}
)

func NewKumaIoMeshTraceV1Alpha1Resource() resource.Resource {
	return &KumaIoMeshTraceV1Alpha1Resource{}
}

type KumaIoMeshTraceV1Alpha1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type KumaIoMeshTraceV1Alpha1ResourceData struct {
	ID             types.String `tfsdk:"id" json:"-"`
	ForceConflicts types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager   types.String `tfsdk:"field_manager" json:"-"`
	WaitFor        types.List   `tfsdk:"wait_for" json:"-"`

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

func (r *KumaIoMeshTraceV1Alpha1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kuma_io_mesh_trace_v1alpha1"
}

func (r *KumaIoMeshTraceV1Alpha1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
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

			"force_conflicts": schema.BoolAttribute{
				Description:         "If 'true', server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "If `true`, server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"field_manager": schema.BoolAttribute{
				Description:         "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"wait_for": schema.ListNestedAttribute{
				Description:         "Wait for specific conditions after create/update of resources.",
				MarkdownDescription: "Wait for specific conditions after create/update of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"jsonpath": schema.StringAttribute{
							Description:         "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							MarkdownDescription: "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							Required:            true,
							Optional:            false,
							Computed:            false,
						},
						"value": schema.StringAttribute{
							Description:         "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							MarkdownDescription: "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							Required:            false,
							Optional:            true,
							Computed:            true,
						},
						"timeout": schema.StringAttribute{
							Description:         "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							MarkdownDescription: "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("30s"),
						},
					},
				},
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},

					"labels": schema.MapAttribute{
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            true,
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
						Computed:            true,
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
													Optional:            true,
													Computed:            false,
												},

												"url": schema.StringAttribute{
													Description:         "Address of Datadog collector, only host and port are allowed (no paths, fragments etc.)",
													MarkdownDescription: "Address of Datadog collector, only host and port are allowed (no paths, fragments etc.)",
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
													Description:         "Version of the API. values: httpJson, httpProto. Default: httpJson see https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/trace/v3/zipkin.proto#L66",
													MarkdownDescription: "Version of the API. values: httpJson, httpProto. Default: httpJson see https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/trace/v3/zipkin.proto#L66",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("httpJson", "httpProto"),
													},
												},

												"shared_span_context": schema.BoolAttribute{
													Description:         "Determines whether client and server spans will share the same span context. Default: true. https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/trace/v3/zipkin.proto#L63",
													MarkdownDescription: "Determines whether client and server spans will share the same span context. Default: true. https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/trace/v3/zipkin.proto#L63",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"trace_id128bit": schema.BoolAttribute{
													Description:         "Generate 128bit traces. Default: false",
													MarkdownDescription: "Generate 128bit traces. Default: false",
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
								Description:         "Sampling configuration. Sampling is the process by which a decision is made on whether to process/export a span or not.",
								MarkdownDescription: "Sampling configuration. Sampling is the process by which a decision is made on whether to process/export a span or not.",
								Attributes: map[string]schema.Attribute{
									"client": schema.StringAttribute{
										Description:         "Target percentage of requests that will be force traced if the 'x-client-trace-id' header is set. Default: 100% Mirror of client_sampling in Envoy https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/filter/network/http_connection_manager/v2/http_connection_manager.proto#L127-L133 Either int or decimal represented as string.",
										MarkdownDescription: "Target percentage of requests that will be force traced if the 'x-client-trace-id' header is set. Default: 100% Mirror of client_sampling in Envoy https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/filter/network/http_connection_manager/v2/http_connection_manager.proto#L127-L133 Either int or decimal represented as string.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"overall": schema.StringAttribute{
										Description:         "Target percentage of requests will be traced after all other sampling checks have been applied (client, force tracing, random sampling). This field functions as an upper limit on the total configured sampling rate. For instance, setting client_sampling to 100% but overall_sampling to 1% will result in only 1% of client requests with the appropriate headers to be force traced. Default: 100% Mirror of overall_sampling in Envoy https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/filter/network/http_connection_manager/v2/http_connection_manager.proto#L142-L150 Either int or decimal represented as string.",
										MarkdownDescription: "Target percentage of requests will be traced after all other sampling checks have been applied (client, force tracing, random sampling). This field functions as an upper limit on the total configured sampling rate. For instance, setting client_sampling to 100% but overall_sampling to 1% will result in only 1% of client requests with the appropriate headers to be force traced. Default: 100% Mirror of overall_sampling in Envoy https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/filter/network/http_connection_manager/v2/http_connection_manager.proto#L142-L150 Either int or decimal represented as string.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"random": schema.StringAttribute{
										Description:         "Target percentage of requests that will be randomly selected for trace generation, if not requested by the client or not forced. Default: 100% Mirror of random_sampling in Envoy https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/filter/network/http_connection_manager/v2/http_connection_manager.proto#L135-L140 Either int or decimal represented as string.",
										MarkdownDescription: "Target percentage of requests that will be randomly selected for trace generation, if not requested by the client or not forced. Default: 100% Mirror of random_sampling in Envoy https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/filter/network/http_connection_manager/v2/http_connection_manager.proto#L135-L140 Either int or decimal represented as string.",
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
						Description:         "TargetRef is a reference to the resource the policy takes an effect on. The resource could be either a real store object or virtual resource defined inplace.",
						MarkdownDescription: "TargetRef is a reference to the resource the policy takes an effect on. The resource could be either a real store object or virtual resource defined inplace.",
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
								Description:         "Name of the referenced resource. Can only be used with kinds: 'MeshService', 'MeshServiceSubset' and 'MeshGatewayRoute'",
								MarkdownDescription: "Name of the referenced resource. Can only be used with kinds: 'MeshService', 'MeshServiceSubset' and 'MeshGatewayRoute'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tags": schema.MapAttribute{
								Description:         "Tags used to select a subset of proxies by tags. Can only be used with kinds 'MeshSubset' and 'MeshServiceSubset'",
								MarkdownDescription: "Tags used to select a subset of proxies by tags. Can only be used with kinds 'MeshSubset' and 'MeshServiceSubset'",
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

func (r *KumaIoMeshTraceV1Alpha1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *dynamic.DynamicClient, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *KumaIoMeshTraceV1Alpha1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_kuma_io_mesh_trace_v1alpha1")

	var model KumaIoMeshTraceV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("kuma.io/v1alpha1")
	model.Kind = pointer.String("MeshTrace")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "kuma.io", Version: "v1alpha1", Resource: "meshtraces"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while creating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse KumaIoMeshTraceV1Alpha1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *KumaIoMeshTraceV1Alpha1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kuma_io_mesh_trace_v1alpha1")

	var data KumaIoMeshTraceV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "kuma.io", Version: "v1alpha1", Resource: "meshtraces"}).
		Namespace(data.Metadata.Namespace).
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

	var readResponse KumaIoMeshTraceV1Alpha1ResourceData
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

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *KumaIoMeshTraceV1Alpha1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_kuma_io_mesh_trace_v1alpha1")

	var model KumaIoMeshTraceV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("kuma.io/v1alpha1")
	model.Kind = pointer.String("MeshTrace")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "kuma.io", Version: "v1alpha1", Resource: "meshtraces"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while updating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse KumaIoMeshTraceV1Alpha1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *KumaIoMeshTraceV1Alpha1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_kuma_io_mesh_trace_v1alpha1")

	var data KumaIoMeshTraceV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "kuma.io", Version: "v1alpha1", Resource: "meshtraces"}).
		Namespace(data.Metadata.Namespace).
		Delete(ctx, data.Metadata.Name, meta.DeleteOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to DELETE resource",
			"An unexpected error occurred while deleting the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"DELETE Error: "+err.Error(),
		)
		return
	}
}

func (r *KumaIoMeshTraceV1Alpha1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	idParts := strings.Split(request.ID, "/")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'namespace/name' Got: '%q'", request.ID),
		)
		return
	}

	namespace := idParts[0]
	name := idParts[1]
	tflog.Trace(ctx, "parsed import ID", map[string]interface{}{
		"namespace": namespace,
		"name":      name,
	})
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("namespace"), namespace)...)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("name"), name)...)
}
