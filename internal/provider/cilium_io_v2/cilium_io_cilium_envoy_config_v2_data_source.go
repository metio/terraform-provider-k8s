/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package cilium_io_v2

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
	_ datasource.DataSource              = &CiliumIoCiliumEnvoyConfigV2DataSource{}
	_ datasource.DataSourceWithConfigure = &CiliumIoCiliumEnvoyConfigV2DataSource{}
)

func NewCiliumIoCiliumEnvoyConfigV2DataSource() datasource.DataSource {
	return &CiliumIoCiliumEnvoyConfigV2DataSource{}
}

type CiliumIoCiliumEnvoyConfigV2DataSource struct {
	kubernetesClient dynamic.Interface
}

type CiliumIoCiliumEnvoyConfigV2DataSourceData struct {
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
		BackendServices *[]struct {
			Name      *string   `tfsdk:"name" json:"name,omitempty"`
			Namespace *string   `tfsdk:"namespace" json:"namespace,omitempty"`
			Number    *[]string `tfsdk:"number" json:"number,omitempty"`
		} `tfsdk:"backend_services" json:"backendServices,omitempty"`
		Resources *[]map[string]string `tfsdk:"resources" json:"resources,omitempty"`
		Services  *[]struct {
			Listener  *string `tfsdk:"listener" json:"listener,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"services" json:"services,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CiliumIoCiliumEnvoyConfigV2DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_cilium_io_cilium_envoy_config_v2"
}

func (r *CiliumIoCiliumEnvoyConfigV2DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"backend_services": schema.ListNestedAttribute{
						Description:         "BackendServices specifies Kubernetes services whose backends are automatically synced to Envoy using EDS.  Traffic for these services is not forwarded to an Envoy listener. This allows an Envoy listener load balance traffic to these backends while normal Cilium service load balancing takes care of balancing traffic for these services at the same time.",
						MarkdownDescription: "BackendServices specifies Kubernetes services whose backends are automatically synced to Envoy using EDS.  Traffic for these services is not forwarded to an Envoy listener. This allows an Envoy listener load balance traffic to these backends while normal Cilium service load balancing takes care of balancing traffic for these services at the same time.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name is the name of a destination Kubernetes service that identifies traffic to be redirected.",
									MarkdownDescription: "Name is the name of a destination Kubernetes service that identifies traffic to be redirected.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace is the Kubernetes service namespace. In CiliumEnvoyConfig namespace defaults to the namespace of the CEC, In CiliumClusterwideEnvoyConfig namespace defaults to 'default'.",
									MarkdownDescription: "Namespace is the Kubernetes service namespace. In CiliumEnvoyConfig namespace defaults to the namespace of the CEC, In CiliumClusterwideEnvoyConfig namespace defaults to 'default'.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"number": schema.ListAttribute{
									Description:         "Port is the port number, which can be used for filtering in case of underlying is exposing multiple port numbers.",
									MarkdownDescription: "Port is the port number, which can be used for filtering in case of underlying is exposing multiple port numbers.",
									ElementType:         types.StringType,
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

					"resources": schema.ListAttribute{
						Description:         "Envoy xDS resources, a list of the following Envoy resource types: type.googleapis.com/envoy.config.listener.v3.Listener, type.googleapis.com/envoy.config.route.v3.RouteConfiguration, type.googleapis.com/envoy.config.cluster.v3.Cluster, type.googleapis.com/envoy.config.endpoint.v3.ClusterLoadAssignment, and type.googleapis.com/envoy.extensions.transport_sockets.tls.v3.Secret.",
						MarkdownDescription: "Envoy xDS resources, a list of the following Envoy resource types: type.googleapis.com/envoy.config.listener.v3.Listener, type.googleapis.com/envoy.config.route.v3.RouteConfiguration, type.googleapis.com/envoy.config.cluster.v3.Cluster, type.googleapis.com/envoy.config.endpoint.v3.ClusterLoadAssignment, and type.googleapis.com/envoy.extensions.transport_sockets.tls.v3.Secret.",
						ElementType:         types.MapType{ElemType: types.StringType},
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"services": schema.ListNestedAttribute{
						Description:         "Services specifies Kubernetes services for which traffic is forwarded to an Envoy listener for L7 load balancing. Backends of these services are automatically synced to Envoy usign EDS.",
						MarkdownDescription: "Services specifies Kubernetes services for which traffic is forwarded to an Envoy listener for L7 load balancing. Backends of these services are automatically synced to Envoy usign EDS.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"listener": schema.StringAttribute{
									Description:         "Listener specifies the name of the Envoy listener the service traffic is redirected to. The listener must be specified in the Envoy 'resources' of the same CiliumEnvoyConfig.  If omitted, the first listener specified in 'resources' is used.",
									MarkdownDescription: "Listener specifies the name of the Envoy listener the service traffic is redirected to. The listener must be specified in the Envoy 'resources' of the same CiliumEnvoyConfig.  If omitted, the first listener specified in 'resources' is used.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"name": schema.StringAttribute{
									Description:         "Name is the name of a destination Kubernetes service that identifies traffic to be redirected.",
									MarkdownDescription: "Name is the name of a destination Kubernetes service that identifies traffic to be redirected.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace is the Kubernetes service namespace. In CiliumEnvoyConfig namespace this is overridden to the namespace of the CEC, In CiliumClusterwideEnvoyConfig namespace defaults to 'default'.",
									MarkdownDescription: "Namespace is the Kubernetes service namespace. In CiliumEnvoyConfig namespace this is overridden to the namespace of the CEC, In CiliumClusterwideEnvoyConfig namespace defaults to 'default'.",
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
		},
	}
}

func (r *CiliumIoCiliumEnvoyConfigV2DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *CiliumIoCiliumEnvoyConfigV2DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_cilium_io_cilium_envoy_config_v2")

	var data CiliumIoCiliumEnvoyConfigV2DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "cilium.io", Version: "v2", Resource: "ciliumenvoyconfigs"}).
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

	var readResponse CiliumIoCiliumEnvoyConfigV2DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("cilium.io/v2")
	data.Kind = pointer.String("CiliumEnvoyConfig")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
