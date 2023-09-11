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
	_ datasource.DataSource              = &CiliumIoCiliumLocalRedirectPolicyV2DataSource{}
	_ datasource.DataSourceWithConfigure = &CiliumIoCiliumLocalRedirectPolicyV2DataSource{}
)

func NewCiliumIoCiliumLocalRedirectPolicyV2DataSource() datasource.DataSource {
	return &CiliumIoCiliumLocalRedirectPolicyV2DataSource{}
}

type CiliumIoCiliumLocalRedirectPolicyV2DataSource struct {
	kubernetesClient dynamic.Interface
}

type CiliumIoCiliumLocalRedirectPolicyV2DataSourceData struct {
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
		Description     *string `tfsdk:"description" json:"description,omitempty"`
		RedirectBackend *struct {
			LocalEndpointSelector *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"local_endpoint_selector" json:"localEndpointSelector,omitempty"`
			ToPorts *[]struct {
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Port     *string `tfsdk:"port" json:"port,omitempty"`
				Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
			} `tfsdk:"to_ports" json:"toPorts,omitempty"`
		} `tfsdk:"redirect_backend" json:"redirectBackend,omitempty"`
		RedirectFrontend *struct {
			AddressMatcher *struct {
				Ip      *string `tfsdk:"ip" json:"ip,omitempty"`
				ToPorts *[]struct {
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Port     *string `tfsdk:"port" json:"port,omitempty"`
					Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
				} `tfsdk:"to_ports" json:"toPorts,omitempty"`
			} `tfsdk:"address_matcher" json:"addressMatcher,omitempty"`
			ServiceMatcher *struct {
				Namespace   *string `tfsdk:"namespace" json:"namespace,omitempty"`
				ServiceName *string `tfsdk:"service_name" json:"serviceName,omitempty"`
				ToPorts     *[]struct {
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Port     *string `tfsdk:"port" json:"port,omitempty"`
					Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
				} `tfsdk:"to_ports" json:"toPorts,omitempty"`
			} `tfsdk:"service_matcher" json:"serviceMatcher,omitempty"`
		} `tfsdk:"redirect_frontend" json:"redirectFrontend,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CiliumIoCiliumLocalRedirectPolicyV2DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_cilium_io_cilium_local_redirect_policy_v2"
}

func (r *CiliumIoCiliumLocalRedirectPolicyV2DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CiliumLocalRedirectPolicy is a Kubernetes Custom Resource that contains a specification to redirect traffic locally within a node.",
		MarkdownDescription: "CiliumLocalRedirectPolicy is a Kubernetes Custom Resource that contains a specification to redirect traffic locally within a node.",
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
				Description:         "Spec is the desired behavior of the local redirect policy.",
				MarkdownDescription: "Spec is the desired behavior of the local redirect policy.",
				Attributes: map[string]schema.Attribute{
					"description": schema.StringAttribute{
						Description:         "Description can be used by the creator of the policy to describe the purpose of this policy.",
						MarkdownDescription: "Description can be used by the creator of the policy to describe the purpose of this policy.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"redirect_backend": schema.SingleNestedAttribute{
						Description:         "RedirectBackend specifies backend configuration to redirect traffic to. It can not be empty.",
						MarkdownDescription: "RedirectBackend specifies backend configuration to redirect traffic to. It can not be empty.",
						Attributes: map[string]schema.Attribute{
							"local_endpoint_selector": schema.SingleNestedAttribute{
								Description:         "LocalEndpointSelector selects node local pod(s) where traffic is redirected to.",
								MarkdownDescription: "LocalEndpointSelector selects node local pod(s) where traffic is redirected to.",
								Attributes: map[string]schema.Attribute{
									"match_expressions": schema.ListNestedAttribute{
										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "key is the label key that the selector applies to.",
													MarkdownDescription: "key is the label key that the selector applies to.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"operator": schema.StringAttribute{
													Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
													MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"values": schema.ListAttribute{
													Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
													MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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

									"match_labels": schema.MapAttribute{
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

							"to_ports": schema.ListNestedAttribute{
								Description:         "ToPorts is a list of L4 ports with protocol of node local pod(s) where traffic is redirected to. When multiple ports are specified, the ports must be named.",
								MarkdownDescription: "ToPorts is a list of L4 ports with protocol of node local pod(s) where traffic is redirected to. When multiple ports are specified, the ports must be named.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name is a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",
											MarkdownDescription: "Name is a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"port": schema.StringAttribute{
											Description:         "Port is an L4 port number. The string will be strictly parsed as a single uint16.",
											MarkdownDescription: "Port is an L4 port number. The string will be strictly parsed as a single uint16.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"protocol": schema.StringAttribute{
											Description:         "Protocol is the L4 protocol. Accepted values: 'TCP', 'UDP'",
											MarkdownDescription: "Protocol is the L4 protocol. Accepted values: 'TCP', 'UDP'",
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

					"redirect_frontend": schema.SingleNestedAttribute{
						Description:         "RedirectFrontend specifies frontend configuration to redirect traffic from. It can not be empty.",
						MarkdownDescription: "RedirectFrontend specifies frontend configuration to redirect traffic from. It can not be empty.",
						Attributes: map[string]schema.Attribute{
							"address_matcher": schema.SingleNestedAttribute{
								Description:         "AddressMatcher is a tuple {IP, port, protocol} that matches traffic to be redirected.",
								MarkdownDescription: "AddressMatcher is a tuple {IP, port, protocol} that matches traffic to be redirected.",
								Attributes: map[string]schema.Attribute{
									"ip": schema.StringAttribute{
										Description:         "IP is a destination ip address for traffic to be redirected.  Example: When it is set to '169.254.169.254', traffic destined to '169.254.169.254' is redirected.",
										MarkdownDescription: "IP is a destination ip address for traffic to be redirected.  Example: When it is set to '169.254.169.254', traffic destined to '169.254.169.254' is redirected.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"to_ports": schema.ListNestedAttribute{
										Description:         "ToPorts is a list of destination L4 ports with protocol for traffic to be redirected. When multiple ports are specified, the ports must be named.  Example: When set to Port: '53' and Protocol: UDP, traffic destined to port '53' with UDP protocol is redirected.",
										MarkdownDescription: "ToPorts is a list of destination L4 ports with protocol for traffic to be redirected. When multiple ports are specified, the ports must be named.  Example: When set to Port: '53' and Protocol: UDP, traffic destined to port '53' with UDP protocol is redirected.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name is a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",
													MarkdownDescription: "Name is a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"port": schema.StringAttribute{
													Description:         "Port is an L4 port number. The string will be strictly parsed as a single uint16.",
													MarkdownDescription: "Port is an L4 port number. The string will be strictly parsed as a single uint16.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"protocol": schema.StringAttribute{
													Description:         "Protocol is the L4 protocol. Accepted values: 'TCP', 'UDP'",
													MarkdownDescription: "Protocol is the L4 protocol. Accepted values: 'TCP', 'UDP'",
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

							"service_matcher": schema.SingleNestedAttribute{
								Description:         "ServiceMatcher specifies Kubernetes service and port that matches traffic to be redirected.",
								MarkdownDescription: "ServiceMatcher specifies Kubernetes service and port that matches traffic to be redirected.",
								Attributes: map[string]schema.Attribute{
									"namespace": schema.StringAttribute{
										Description:         "Namespace is the Kubernetes service namespace. The service namespace must match the namespace of the parent Local Redirect Policy.  For Cluster-wide Local Redirect Policy, this can be any namespace.",
										MarkdownDescription: "Namespace is the Kubernetes service namespace. The service namespace must match the namespace of the parent Local Redirect Policy.  For Cluster-wide Local Redirect Policy, this can be any namespace.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"service_name": schema.StringAttribute{
										Description:         "Name is the name of a destination Kubernetes service that identifies traffic to be redirected. The service type needs to be ClusterIP.  Example: When this field is populated with 'serviceName:myService', all the traffic destined to the cluster IP of this service at the (specified) service port(s) will be redirected.",
										MarkdownDescription: "Name is the name of a destination Kubernetes service that identifies traffic to be redirected. The service type needs to be ClusterIP.  Example: When this field is populated with 'serviceName:myService', all the traffic destined to the cluster IP of this service at the (specified) service port(s) will be redirected.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"to_ports": schema.ListNestedAttribute{
										Description:         "ToPorts is a list of destination service L4 ports with protocol for traffic to be redirected. If not specified, traffic for all the service ports will be redirected. When multiple ports are specified, the ports must be named.",
										MarkdownDescription: "ToPorts is a list of destination service L4 ports with protocol for traffic to be redirected. If not specified, traffic for all the service ports will be redirected. When multiple ports are specified, the ports must be named.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name is a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",
													MarkdownDescription: "Name is a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"port": schema.StringAttribute{
													Description:         "Port is an L4 port number. The string will be strictly parsed as a single uint16.",
													MarkdownDescription: "Port is an L4 port number. The string will be strictly parsed as a single uint16.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"protocol": schema.StringAttribute{
													Description:         "Protocol is the L4 protocol. Accepted values: 'TCP', 'UDP'",
													MarkdownDescription: "Protocol is the L4 protocol. Accepted values: 'TCP', 'UDP'",
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

func (r *CiliumIoCiliumLocalRedirectPolicyV2DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *CiliumIoCiliumLocalRedirectPolicyV2DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_cilium_io_cilium_local_redirect_policy_v2")

	var data CiliumIoCiliumLocalRedirectPolicyV2DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "cilium.io", Version: "v2", Resource: "ciliumlocalredirectpolicies"}).
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

	var readResponse CiliumIoCiliumLocalRedirectPolicyV2DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("cilium.io/v2")
	data.Kind = pointer.String("CiliumLocalRedirectPolicy")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
