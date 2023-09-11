/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package cilium_io_v2alpha1

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
	_ datasource.DataSource              = &CiliumIoCiliumEndpointSliceV2Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &CiliumIoCiliumEndpointSliceV2Alpha1DataSource{}
)

func NewCiliumIoCiliumEndpointSliceV2Alpha1DataSource() datasource.DataSource {
	return &CiliumIoCiliumEndpointSliceV2Alpha1DataSource{}
}

type CiliumIoCiliumEndpointSliceV2Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type CiliumIoCiliumEndpointSliceV2Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Endpoints *[]struct {
		Encryption *struct {
			Key *int64 `tfsdk:"key" json:"key,omitempty"`
		} `tfsdk:"encryption" json:"encryption,omitempty"`
		Id          *int64  `tfsdk:"id" json:"id,omitempty"`
		Name        *string `tfsdk:"name" json:"name,omitempty"`
		Named_ports *[]struct {
			Name     *string `tfsdk:"name" json:"name,omitempty"`
			Port     *int64  `tfsdk:"port" json:"port,omitempty"`
			Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
		} `tfsdk:"named_ports" json:"named-ports,omitempty"`
		Networking *struct {
			Addressing *[]struct {
				Ipv4 *string `tfsdk:"ipv4" json:"ipv4,omitempty"`
				Ipv6 *string `tfsdk:"ipv6" json:"ipv6,omitempty"`
			} `tfsdk:"addressing" json:"addressing,omitempty"`
			Node *string `tfsdk:"node" json:"node,omitempty"`
		} `tfsdk:"networking" json:"networking,omitempty"`
	} `tfsdk:"endpoints" json:"endpoints,omitempty"`
	Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
}

func (r *CiliumIoCiliumEndpointSliceV2Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_cilium_io_cilium_endpoint_slice_v2alpha1"
}

func (r *CiliumIoCiliumEndpointSliceV2Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CiliumEndpointSlice contains a group of CoreCiliumendpoints.",
		MarkdownDescription: "CiliumEndpointSlice contains a group of CoreCiliumendpoints.",
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

			"endpoints": schema.ListNestedAttribute{
				Description:         "Endpoints is a list of coreCEPs packed in a CiliumEndpointSlice",
				MarkdownDescription: "Endpoints is a list of coreCEPs packed in a CiliumEndpointSlice",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"encryption": schema.SingleNestedAttribute{
							Description:         "EncryptionSpec defines the encryption relevant configuration of a node.",
							MarkdownDescription: "EncryptionSpec defines the encryption relevant configuration of a node.",
							Attributes: map[string]schema.Attribute{
								"key": schema.Int64Attribute{
									Description:         "Key is the index to the key to use for encryption or 0 if encryption is disabled.",
									MarkdownDescription: "Key is the index to the key to use for encryption or 0 if encryption is disabled.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},
							},
							Required: false,
							Optional: false,
							Computed: true,
						},

						"id": schema.Int64Attribute{
							Description:         "IdentityID is the numeric identity of the endpoint",
							MarkdownDescription: "IdentityID is the numeric identity of the endpoint",
							Required:            false,
							Optional:            false,
							Computed:            true,
						},

						"name": schema.StringAttribute{
							Description:         "Name indicate as CiliumEndpoint name.",
							MarkdownDescription: "Name indicate as CiliumEndpoint name.",
							Required:            false,
							Optional:            false,
							Computed:            true,
						},

						"named_ports": schema.ListNestedAttribute{
							Description:         "NamedPorts List of named Layer 4 port and protocol pairs which will be used in Network Policy specs.  swagger:model NamedPorts",
							MarkdownDescription: "NamedPorts List of named Layer 4 port and protocol pairs which will be used in Network Policy specs.  swagger:model NamedPorts",
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Optional layer 4 port name",
										MarkdownDescription: "Optional layer 4 port name",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"port": schema.Int64Attribute{
										Description:         "Layer 4 port number",
										MarkdownDescription: "Layer 4 port number",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"protocol": schema.StringAttribute{
										Description:         "Layer 4 protocol Enum: [TCP UDP SCTP ICMP ICMPV6 ANY]",
										MarkdownDescription: "Layer 4 protocol Enum: [TCP UDP SCTP ICMP ICMPV6 ANY]",
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

						"networking": schema.SingleNestedAttribute{
							Description:         "EndpointNetworking is the addressing information of an endpoint.",
							MarkdownDescription: "EndpointNetworking is the addressing information of an endpoint.",
							Attributes: map[string]schema.Attribute{
								"addressing": schema.ListNestedAttribute{
									Description:         "IP4/6 addresses assigned to this Endpoint",
									MarkdownDescription: "IP4/6 addresses assigned to this Endpoint",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"ipv4": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"ipv6": schema.StringAttribute{
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

								"node": schema.StringAttribute{
									Description:         "NodeIP is the IP of the node the endpoint is running on. The IP must be reachable between nodes.",
									MarkdownDescription: "NodeIP is the IP of the node the endpoint is running on. The IP must be reachable between nodes.",
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

			"namespace": schema.StringAttribute{
				Description:         "Namespace indicate as CiliumEndpointSlice namespace. All the CiliumEndpoints within the same namespace are put together in CiliumEndpointSlice.",
				MarkdownDescription: "Namespace indicate as CiliumEndpointSlice namespace. All the CiliumEndpoints within the same namespace are put together in CiliumEndpointSlice.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},
		},
	}
}

func (r *CiliumIoCiliumEndpointSliceV2Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *CiliumIoCiliumEndpointSliceV2Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_cilium_io_cilium_endpoint_slice_v2alpha1")

	var data CiliumIoCiliumEndpointSliceV2Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "cilium.io", Version: "v2alpha1", Resource: "ciliumendpointslices"}).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		var statusError *k8sErrors.StatusError
		if errors.As(err, &statusError) {
			if statusError.Status().Code == http.StatusNotFound {
				response.Diagnostics.AddError(
					"Unable to find resource",
					fmt.Sprintf("The requested resource cannot be found. "+
						"Make sure that it does exist in your cluster and you have set the correct name configured.\n\n"+
						"Name: %s", data.Metadata.Name),
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

	var readResponse CiliumIoCiliumEndpointSliceV2Alpha1DataSourceData
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
	data.ApiVersion = pointer.String("cilium.io/v2alpha1")
	data.Kind = pointer.String("CiliumEndpointSlice")
	data.Metadata = readResponse.Metadata
	data.Endpoints = readResponse.Endpoints
	data.Namespace = readResponse.Namespace

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
