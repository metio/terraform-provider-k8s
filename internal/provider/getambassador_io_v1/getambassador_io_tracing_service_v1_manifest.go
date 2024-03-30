/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package getambassador_io_v1

import (
	"context"
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
	_ datasource.DataSource = &GetambassadorIoTracingServiceV1Manifest{}
)

func NewGetambassadorIoTracingServiceV1Manifest() datasource.DataSource {
	return &GetambassadorIoTracingServiceV1Manifest{}
}

type GetambassadorIoTracingServiceV1Manifest struct{}

type GetambassadorIoTracingServiceV1ManifestData struct {
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
		Ambassador_id *[]string `tfsdk:"ambassador_id" json:"ambassador_id,omitempty"`
		Config        *struct {
			Access_token_file          *string   `tfsdk:"access_token_file" json:"access_token_file,omitempty"`
			Collector_cluster          *string   `tfsdk:"collector_cluster" json:"collector_cluster,omitempty"`
			Collector_endpoint         *string   `tfsdk:"collector_endpoint" json:"collector_endpoint,omitempty"`
			Collector_endpoint_version *string   `tfsdk:"collector_endpoint_version" json:"collector_endpoint_version,omitempty"`
			Collector_hostname         *string   `tfsdk:"collector_hostname" json:"collector_hostname,omitempty"`
			Service_name               *string   `tfsdk:"service_name" json:"service_name,omitempty"`
			Shared_span_context        *bool     `tfsdk:"shared_span_context" json:"shared_span_context,omitempty"`
			Trace_id_128bit            *bool     `tfsdk:"trace_id_128bit" json:"trace_id_128bit,omitempty"`
			V3PropagationModes         *[]string `tfsdk:"v3_propagation_modes" json:"v3PropagationModes,omitempty"`
		} `tfsdk:"config" json:"config,omitempty"`
		Driver   *string `tfsdk:"driver" json:"driver,omitempty"`
		Sampling *struct {
			Client  *int64 `tfsdk:"client" json:"client,omitempty"`
			Overall *int64 `tfsdk:"overall" json:"overall,omitempty"`
			Random  *int64 `tfsdk:"random" json:"random,omitempty"`
		} `tfsdk:"sampling" json:"sampling,omitempty"`
		Service      *string   `tfsdk:"service" json:"service,omitempty"`
		Tag_headers  *[]string `tfsdk:"tag_headers" json:"tag_headers,omitempty"`
		V3CustomTags *[]struct {
			Environment *struct {
				Default_value *string `tfsdk:"default_value" json:"default_value,omitempty"`
				Name          *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"environment" json:"environment,omitempty"`
			Literal *struct {
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"literal" json:"literal,omitempty"`
			Request_header *struct {
				Default_value *string `tfsdk:"default_value" json:"default_value,omitempty"`
				Name          *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"request_header" json:"request_header,omitempty"`
			Tag *string `tfsdk:"tag" json:"tag,omitempty"`
		} `tfsdk:"v3_custom_tags" json:"v3CustomTags,omitempty"`
		V3StatsName *string `tfsdk:"v3_stats_name" json:"v3StatsName,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *GetambassadorIoTracingServiceV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_getambassador_io_tracing_service_v1_manifest"
}

func (r *GetambassadorIoTracingServiceV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "TracingService is the Schema for the tracingservices API",
		MarkdownDescription: "TracingService is the Schema for the tracingservices API",
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
				Description:         "TracingServiceSpec defines the desired state of TracingService",
				MarkdownDescription: "TracingServiceSpec defines the desired state of TracingService",
				Attributes: map[string]schema.Attribute{
					"ambassador_id": schema.ListAttribute{
						Description:         "AmbassadorID declares which Ambassador instances should pay attention to this resource.  May either be a string or a list of strings.  If no value is provided, the default is:  ambassador_id: - 'default'",
						MarkdownDescription: "AmbassadorID declares which Ambassador instances should pay attention to this resource.  May either be a string or a list of strings.  If no value is provided, the default is:  ambassador_id: - 'default'",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"config": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"access_token_file": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"collector_cluster": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"collector_endpoint": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"collector_endpoint_version": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("HTTP_JSON_V1", "HTTP_JSON", "HTTP_PROTO"),
								},
							},

							"collector_hostname": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"shared_span_context": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"trace_id_128bit": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"v3_propagation_modes": schema.ListAttribute{
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

					"driver": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("lightstep", "zipkin", "datadog", "opentelemetry"),
						},
					},

					"sampling": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"client": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"overall": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"random": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"service": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"tag_headers": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"v3_custom_tags": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"environment": schema.SingleNestedAttribute{
									Description:         "Environment explicitly specifies the protocol stack to set up. Exactly one of Literal, Environment or Header must be supplied.",
									MarkdownDescription: "Environment explicitly specifies the protocol stack to set up. Exactly one of Literal, Environment or Header must be supplied.",
									Attributes: map[string]schema.Attribute{
										"default_value": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"literal": schema.SingleNestedAttribute{
									Description:         "Literal explicitly specifies the protocol stack to set up. Exactly one of Literal, Environment or Header must be supplied.",
									MarkdownDescription: "Literal explicitly specifies the protocol stack to set up. Exactly one of Literal, Environment or Header must be supplied.",
									Attributes: map[string]schema.Attribute{
										"value": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"request_header": schema.SingleNestedAttribute{
									Description:         "Header explicitly specifies the protocol stack to set up. Exactly one of Literal, Environment or Header must be supplied.",
									MarkdownDescription: "Header explicitly specifies the protocol stack to set up. Exactly one of Literal, Environment or Header must be supplied.",
									Attributes: map[string]schema.Attribute{
										"default_value": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"tag": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"v3_stats_name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
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

func (r *GetambassadorIoTracingServiceV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_getambassador_io_tracing_service_v1_manifest")

	var model GetambassadorIoTracingServiceV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("getambassador.io/v1")
	model.Kind = pointer.String("TracingService")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
