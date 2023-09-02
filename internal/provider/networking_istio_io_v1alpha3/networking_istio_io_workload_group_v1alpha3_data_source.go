/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package networking_istio_io_v1alpha3

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
	_ datasource.DataSource              = &NetworkingIstioIoWorkloadGroupV1Alpha3DataSource{}
	_ datasource.DataSourceWithConfigure = &NetworkingIstioIoWorkloadGroupV1Alpha3DataSource{}
)

func NewNetworkingIstioIoWorkloadGroupV1Alpha3DataSource() datasource.DataSource {
	return &NetworkingIstioIoWorkloadGroupV1Alpha3DataSource{}
}

type NetworkingIstioIoWorkloadGroupV1Alpha3DataSource struct {
	kubernetesClient dynamic.Interface
}

type NetworkingIstioIoWorkloadGroupV1Alpha3DataSourceData struct {
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
		Metadata *struct {
			Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		} `tfsdk:"metadata" json:"metadata,omitempty"`
		Probe *struct {
			Exec *struct {
				Command *[]string `tfsdk:"command" json:"command,omitempty"`
			} `tfsdk:"exec" json:"exec,omitempty"`
			FailureThreshold *int64 `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
			HttpGet          *struct {
				Host        *string `tfsdk:"host" json:"host,omitempty"`
				HttpHeaders *[]struct {
					Name  *string `tfsdk:"name" json:"name,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"http_headers" json:"httpHeaders,omitempty"`
				Path   *string `tfsdk:"path" json:"path,omitempty"`
				Port   *int64  `tfsdk:"port" json:"port,omitempty"`
				Scheme *string `tfsdk:"scheme" json:"scheme,omitempty"`
			} `tfsdk:"http_get" json:"httpGet,omitempty"`
			InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" json:"initialDelaySeconds,omitempty"`
			PeriodSeconds       *int64 `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
			SuccessThreshold    *int64 `tfsdk:"success_threshold" json:"successThreshold,omitempty"`
			TcpSocket           *struct {
				Host *string `tfsdk:"host" json:"host,omitempty"`
				Port *int64  `tfsdk:"port" json:"port,omitempty"`
			} `tfsdk:"tcp_socket" json:"tcpSocket,omitempty"`
			TimeoutSeconds *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
		} `tfsdk:"probe" json:"probe,omitempty"`
		Template *struct {
			Address        *string            `tfsdk:"address" json:"address,omitempty"`
			Labels         *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Locality       *string            `tfsdk:"locality" json:"locality,omitempty"`
			Network        *string            `tfsdk:"network" json:"network,omitempty"`
			Ports          *map[string]string `tfsdk:"ports" json:"ports,omitempty"`
			ServiceAccount *string            `tfsdk:"service_account" json:"serviceAccount,omitempty"`
			Weight         *int64             `tfsdk:"weight" json:"weight,omitempty"`
		} `tfsdk:"template" json:"template,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *NetworkingIstioIoWorkloadGroupV1Alpha3DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_networking_istio_io_workload_group_v1alpha3"
}

func (r *NetworkingIstioIoWorkloadGroupV1Alpha3DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "Describes a collection of workload instances. See more details at: https://istio.io/docs/reference/config/networking/workload-group.html",
				MarkdownDescription: "Describes a collection of workload instances. See more details at: https://istio.io/docs/reference/config/networking/workload-group.html",
				Attributes: map[string]schema.Attribute{
					"metadata": schema.SingleNestedAttribute{
						Description:         "Metadata that will be used for all corresponding 'WorkloadEntries'.",
						MarkdownDescription: "Metadata that will be used for all corresponding 'WorkloadEntries'.",
						Attributes: map[string]schema.Attribute{
							"annotations": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"labels": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
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

					"probe": schema.SingleNestedAttribute{
						Description:         "'ReadinessProbe' describes the configuration the user must provide for healthchecking on their workload.",
						MarkdownDescription: "'ReadinessProbe' describes the configuration the user must provide for healthchecking on their workload.",
						Attributes: map[string]schema.Attribute{
							"exec": schema.SingleNestedAttribute{
								Description:         "Health is determined by how the command that is executed exited.",
								MarkdownDescription: "Health is determined by how the command that is executed exited.",
								Attributes: map[string]schema.Attribute{
									"command": schema.ListAttribute{
										Description:         "Command to run.",
										MarkdownDescription: "Command to run.",
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

							"failure_threshold": schema.Int64Attribute{
								Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded.",
								MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"http_get": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"host": schema.StringAttribute{
										Description:         "Host name to connect to, defaults to the pod IP.",
										MarkdownDescription: "Host name to connect to, defaults to the pod IP.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"http_headers": schema.ListNestedAttribute{
										Description:         "Headers the proxy will pass on to make the request.",
										MarkdownDescription: "Headers the proxy will pass on to make the request.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"value": schema.StringAttribute{
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

									"path": schema.StringAttribute{
										Description:         "Path to access on the HTTP server.",
										MarkdownDescription: "Path to access on the HTTP server.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"port": schema.Int64Attribute{
										Description:         "Port on which the endpoint lives.",
										MarkdownDescription: "Port on which the endpoint lives.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"scheme": schema.StringAttribute{
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

							"initial_delay_seconds": schema.Int64Attribute{
								Description:         "Number of seconds after the container has started before readiness probes are initiated.",
								MarkdownDescription: "Number of seconds after the container has started before readiness probes are initiated.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"period_seconds": schema.Int64Attribute{
								Description:         "How often (in seconds) to perform the probe.",
								MarkdownDescription: "How often (in seconds) to perform the probe.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"success_threshold": schema.Int64Attribute{
								Description:         "Minimum consecutive successes for the probe to be considered successful after having failed.",
								MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"tcp_socket": schema.SingleNestedAttribute{
								Description:         "Health is determined by if the proxy is able to connect.",
								MarkdownDescription: "Health is determined by if the proxy is able to connect.",
								Attributes: map[string]schema.Attribute{
									"host": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"port": schema.Int64Attribute{
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

							"timeout_seconds": schema.Int64Attribute{
								Description:         "Number of seconds after which the probe times out.",
								MarkdownDescription: "Number of seconds after which the probe times out.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"template": schema.SingleNestedAttribute{
						Description:         "Template to be used for the generation of 'WorkloadEntry' resources that belong to this 'WorkloadGroup'.",
						MarkdownDescription: "Template to be used for the generation of 'WorkloadEntry' resources that belong to this 'WorkloadGroup'.",
						Attributes: map[string]schema.Attribute{
							"address": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"labels": schema.MapAttribute{
								Description:         "One or more labels associated with the endpoint.",
								MarkdownDescription: "One or more labels associated with the endpoint.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"locality": schema.StringAttribute{
								Description:         "The locality associated with the endpoint.",
								MarkdownDescription: "The locality associated with the endpoint.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"network": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"ports": schema.MapAttribute{
								Description:         "Set of ports associated with the endpoint.",
								MarkdownDescription: "Set of ports associated with the endpoint.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"service_account": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"weight": schema.Int64Attribute{
								Description:         "The load balancing weight associated with the endpoint.",
								MarkdownDescription: "The load balancing weight associated with the endpoint.",
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

func (r *NetworkingIstioIoWorkloadGroupV1Alpha3DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *NetworkingIstioIoWorkloadGroupV1Alpha3DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_networking_istio_io_workload_group_v1alpha3")

	var data NetworkingIstioIoWorkloadGroupV1Alpha3DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "networking.istio.io", Version: "v1alpha3", Resource: "WorkloadGroup"}).
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

	var readResponse NetworkingIstioIoWorkloadGroupV1Alpha3DataSourceData
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
	data.ApiVersion = pointer.String("networking.istio.io/v1alpha3")
	data.Kind = pointer.String("WorkloadGroup")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
