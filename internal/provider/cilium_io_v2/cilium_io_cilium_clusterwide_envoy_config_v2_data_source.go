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
	_ datasource.DataSource              = &CiliumIoCiliumClusterwideEnvoyConfigV2DataSource{}
	_ datasource.DataSourceWithConfigure = &CiliumIoCiliumClusterwideEnvoyConfigV2DataSource{}
)

func NewCiliumIoCiliumClusterwideEnvoyConfigV2DataSource() datasource.DataSource {
	return &CiliumIoCiliumClusterwideEnvoyConfigV2DataSource{}
}

type CiliumIoCiliumClusterwideEnvoyConfigV2DataSource struct {
	kubernetesClient dynamic.Interface
}

type CiliumIoCiliumClusterwideEnvoyConfigV2DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
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

func (r *CiliumIoCiliumClusterwideEnvoyConfigV2DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_cilium_io_cilium_clusterwide_envoy_config_v2"
}

func (r *CiliumIoCiliumClusterwideEnvoyConfigV2DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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

func (r *CiliumIoCiliumClusterwideEnvoyConfigV2DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *CiliumIoCiliumClusterwideEnvoyConfigV2DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_cilium_io_cilium_clusterwide_envoy_config_v2")

	var data CiliumIoCiliumClusterwideEnvoyConfigV2DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "cilium.io", Version: "v2", Resource: "CiliumClusterwideEnvoyConfig"}).
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

	var readResponse CiliumIoCiliumClusterwideEnvoyConfigV2DataSourceData
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

	data.ID = types.StringValue(data.Metadata.Name)
	data.ApiVersion = pointer.String("cilium.io/v2")
	data.Kind = pointer.String("CiliumClusterwideEnvoyConfig")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
