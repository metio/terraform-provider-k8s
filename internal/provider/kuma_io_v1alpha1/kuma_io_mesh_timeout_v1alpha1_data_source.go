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
	_ datasource.DataSource              = &KumaIoMeshTimeoutV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &KumaIoMeshTimeoutV1Alpha1DataSource{}
)

func NewKumaIoMeshTimeoutV1Alpha1DataSource() datasource.DataSource {
	return &KumaIoMeshTimeoutV1Alpha1DataSource{}
}

type KumaIoMeshTimeoutV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type KumaIoMeshTimeoutV1Alpha1DataSourceData struct {
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
		From *[]struct {
			Default *struct {
				ConnectionTimeout *string `tfsdk:"connection_timeout" json:"connectionTimeout,omitempty"`
				Http              *struct {
					MaxConnectionDuration *string `tfsdk:"max_connection_duration" json:"maxConnectionDuration,omitempty"`
					MaxStreamDuration     *string `tfsdk:"max_stream_duration" json:"maxStreamDuration,omitempty"`
					RequestTimeout        *string `tfsdk:"request_timeout" json:"requestTimeout,omitempty"`
					StreamIdleTimeout     *string `tfsdk:"stream_idle_timeout" json:"streamIdleTimeout,omitempty"`
				} `tfsdk:"http" json:"http,omitempty"`
				IdleTimeout *string `tfsdk:"idle_timeout" json:"idleTimeout,omitempty"`
			} `tfsdk:"default" json:"default,omitempty"`
			TargetRef *struct {
				Kind *string            `tfsdk:"kind" json:"kind,omitempty"`
				Mesh *string            `tfsdk:"mesh" json:"mesh,omitempty"`
				Name *string            `tfsdk:"name" json:"name,omitempty"`
				Tags *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
			} `tfsdk:"target_ref" json:"targetRef,omitempty"`
		} `tfsdk:"from" json:"from,omitempty"`
		TargetRef *struct {
			Kind *string            `tfsdk:"kind" json:"kind,omitempty"`
			Mesh *string            `tfsdk:"mesh" json:"mesh,omitempty"`
			Name *string            `tfsdk:"name" json:"name,omitempty"`
			Tags *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
		} `tfsdk:"target_ref" json:"targetRef,omitempty"`
		To *[]struct {
			Default *struct {
				ConnectionTimeout *string `tfsdk:"connection_timeout" json:"connectionTimeout,omitempty"`
				Http              *struct {
					MaxConnectionDuration *string `tfsdk:"max_connection_duration" json:"maxConnectionDuration,omitempty"`
					MaxStreamDuration     *string `tfsdk:"max_stream_duration" json:"maxStreamDuration,omitempty"`
					RequestTimeout        *string `tfsdk:"request_timeout" json:"requestTimeout,omitempty"`
					StreamIdleTimeout     *string `tfsdk:"stream_idle_timeout" json:"streamIdleTimeout,omitempty"`
				} `tfsdk:"http" json:"http,omitempty"`
				IdleTimeout *string `tfsdk:"idle_timeout" json:"idleTimeout,omitempty"`
			} `tfsdk:"default" json:"default,omitempty"`
			TargetRef *struct {
				Kind *string            `tfsdk:"kind" json:"kind,omitempty"`
				Mesh *string            `tfsdk:"mesh" json:"mesh,omitempty"`
				Name *string            `tfsdk:"name" json:"name,omitempty"`
				Tags *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
			} `tfsdk:"target_ref" json:"targetRef,omitempty"`
		} `tfsdk:"to" json:"to,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KumaIoMeshTimeoutV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kuma_io_mesh_timeout_v1alpha1"
}

func (r *KumaIoMeshTimeoutV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "Spec is the specification of the Kuma MeshTimeout resource.",
				MarkdownDescription: "Spec is the specification of the Kuma MeshTimeout resource.",
				Attributes: map[string]schema.Attribute{
					"from": schema.ListNestedAttribute{
						Description:         "From list makes a match between clients and corresponding configurations",
						MarkdownDescription: "From list makes a match between clients and corresponding configurations",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"default": schema.SingleNestedAttribute{
									Description:         "Default is a configuration specific to the group of clients referenced in 'targetRef'",
									MarkdownDescription: "Default is a configuration specific to the group of clients referenced in 'targetRef'",
									Attributes: map[string]schema.Attribute{
										"connection_timeout": schema.StringAttribute{
											Description:         "ConnectionTimeout specifies the amount of time proxy will wait for an TCP connection to be established. Default value is 5 seconds. Cannot be set to 0.",
											MarkdownDescription: "ConnectionTimeout specifies the amount of time proxy will wait for an TCP connection to be established. Default value is 5 seconds. Cannot be set to 0.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"http": schema.SingleNestedAttribute{
											Description:         "Http provides configuration for HTTP specific timeouts",
											MarkdownDescription: "Http provides configuration for HTTP specific timeouts",
											Attributes: map[string]schema.Attribute{
												"max_connection_duration": schema.StringAttribute{
													Description:         "MaxConnectionDuration is the time after which a connection will be drained and/or closed, starting from when it was first established. Setting this timeout to 0 will disable it. Disabled by default.",
													MarkdownDescription: "MaxConnectionDuration is the time after which a connection will be drained and/or closed, starting from when it was first established. Setting this timeout to 0 will disable it. Disabled by default.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"max_stream_duration": schema.StringAttribute{
													Description:         "MaxStreamDuration is the maximum time that a stream’s lifetime will span. Setting this timeout to 0 will disable it. Disabled by default.",
													MarkdownDescription: "MaxStreamDuration is the maximum time that a stream’s lifetime will span. Setting this timeout to 0 will disable it. Disabled by default.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"request_timeout": schema.StringAttribute{
													Description:         "RequestTimeout The amount of time that proxy will wait for the entire request to be received. The timer is activated when the request is initiated, and is disarmed when the last byte of the request is sent, OR when the response is initiated. Setting this timeout to 0 will disable it. Default is 15s.",
													MarkdownDescription: "RequestTimeout The amount of time that proxy will wait for the entire request to be received. The timer is activated when the request is initiated, and is disarmed when the last byte of the request is sent, OR when the response is initiated. Setting this timeout to 0 will disable it. Default is 15s.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"stream_idle_timeout": schema.StringAttribute{
													Description:         "StreamIdleTimeout is the amount of time that proxy will allow a stream to exist with no activity. Setting this timeout to 0 will disable it. Default is 30m",
													MarkdownDescription: "StreamIdleTimeout is the amount of time that proxy will allow a stream to exist with no activity. Setting this timeout to 0 will disable it. Default is 30m",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"idle_timeout": schema.StringAttribute{
											Description:         "IdleTimeout is defined as the period in which there are no bytes sent or received on connection Setting this timeout to 0 will disable it. Be cautious when disabling it because it can lead to connection leaking. Default value is 1h.",
											MarkdownDescription: "IdleTimeout is defined as the period in which there are no bytes sent or received on connection Setting this timeout to 0 will disable it. Be cautious when disabling it because it can lead to connection leaking. Default value is 1h.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"target_ref": schema.SingleNestedAttribute{
									Description:         "TargetRef is a reference to the resource that represents a group of clients.",
									MarkdownDescription: "TargetRef is a reference to the resource that represents a group of clients.",
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

					"to": schema.ListNestedAttribute{
						Description:         "To list makes a match between the consumed services and corresponding configurations",
						MarkdownDescription: "To list makes a match between the consumed services and corresponding configurations",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"default": schema.SingleNestedAttribute{
									Description:         "Default is a configuration specific to the group of destinations referenced in 'targetRef'",
									MarkdownDescription: "Default is a configuration specific to the group of destinations referenced in 'targetRef'",
									Attributes: map[string]schema.Attribute{
										"connection_timeout": schema.StringAttribute{
											Description:         "ConnectionTimeout specifies the amount of time proxy will wait for an TCP connection to be established. Default value is 5 seconds. Cannot be set to 0.",
											MarkdownDescription: "ConnectionTimeout specifies the amount of time proxy will wait for an TCP connection to be established. Default value is 5 seconds. Cannot be set to 0.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"http": schema.SingleNestedAttribute{
											Description:         "Http provides configuration for HTTP specific timeouts",
											MarkdownDescription: "Http provides configuration for HTTP specific timeouts",
											Attributes: map[string]schema.Attribute{
												"max_connection_duration": schema.StringAttribute{
													Description:         "MaxConnectionDuration is the time after which a connection will be drained and/or closed, starting from when it was first established. Setting this timeout to 0 will disable it. Disabled by default.",
													MarkdownDescription: "MaxConnectionDuration is the time after which a connection will be drained and/or closed, starting from when it was first established. Setting this timeout to 0 will disable it. Disabled by default.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"max_stream_duration": schema.StringAttribute{
													Description:         "MaxStreamDuration is the maximum time that a stream’s lifetime will span. Setting this timeout to 0 will disable it. Disabled by default.",
													MarkdownDescription: "MaxStreamDuration is the maximum time that a stream’s lifetime will span. Setting this timeout to 0 will disable it. Disabled by default.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"request_timeout": schema.StringAttribute{
													Description:         "RequestTimeout The amount of time that proxy will wait for the entire request to be received. The timer is activated when the request is initiated, and is disarmed when the last byte of the request is sent, OR when the response is initiated. Setting this timeout to 0 will disable it. Default is 15s.",
													MarkdownDescription: "RequestTimeout The amount of time that proxy will wait for the entire request to be received. The timer is activated when the request is initiated, and is disarmed when the last byte of the request is sent, OR when the response is initiated. Setting this timeout to 0 will disable it. Default is 15s.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"stream_idle_timeout": schema.StringAttribute{
													Description:         "StreamIdleTimeout is the amount of time that proxy will allow a stream to exist with no activity. Setting this timeout to 0 will disable it. Default is 30m",
													MarkdownDescription: "StreamIdleTimeout is the amount of time that proxy will allow a stream to exist with no activity. Setting this timeout to 0 will disable it. Default is 30m",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"idle_timeout": schema.StringAttribute{
											Description:         "IdleTimeout is defined as the period in which there are no bytes sent or received on connection Setting this timeout to 0 will disable it. Be cautious when disabling it because it can lead to connection leaking. Default value is 1h.",
											MarkdownDescription: "IdleTimeout is defined as the period in which there are no bytes sent or received on connection Setting this timeout to 0 will disable it. Be cautious when disabling it because it can lead to connection leaking. Default value is 1h.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"target_ref": schema.SingleNestedAttribute{
									Description:         "TargetRef is a reference to the resource that represents a group of destinations.",
									MarkdownDescription: "TargetRef is a reference to the resource that represents a group of destinations.",
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

func (r *KumaIoMeshTimeoutV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *KumaIoMeshTimeoutV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_kuma_io_mesh_timeout_v1alpha1")

	var data KumaIoMeshTimeoutV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "kuma.io", Version: "v1alpha1", Resource: "MeshTimeout"}).
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

	var readResponse KumaIoMeshTimeoutV1Alpha1DataSourceData
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
	data.Kind = pointer.String("MeshTimeout")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
