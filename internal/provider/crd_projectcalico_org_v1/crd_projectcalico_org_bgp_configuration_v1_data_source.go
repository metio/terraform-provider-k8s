/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package crd_projectcalico_org_v1

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
	_ datasource.DataSource              = &CrdProjectcalicoOrgBgpconfigurationV1DataSource{}
	_ datasource.DataSourceWithConfigure = &CrdProjectcalicoOrgBgpconfigurationV1DataSource{}
)

func NewCrdProjectcalicoOrgBgpconfigurationV1DataSource() datasource.DataSource {
	return &CrdProjectcalicoOrgBgpconfigurationV1DataSource{}
}

type CrdProjectcalicoOrgBgpconfigurationV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type CrdProjectcalicoOrgBgpconfigurationV1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		AsNumber    *int64  `tfsdk:"as_number" json:"asNumber,omitempty"`
		BindMode    *string `tfsdk:"bind_mode" json:"bindMode,omitempty"`
		Communities *[]struct {
			Name  *string `tfsdk:"name" json:"name,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"communities" json:"communities,omitempty"`
		IgnoredInterfaces      *[]string `tfsdk:"ignored_interfaces" json:"ignoredInterfaces,omitempty"`
		ListenPort             *int64    `tfsdk:"listen_port" json:"listenPort,omitempty"`
		LogSeverityScreen      *string   `tfsdk:"log_severity_screen" json:"logSeverityScreen,omitempty"`
		NodeMeshMaxRestartTime *string   `tfsdk:"node_mesh_max_restart_time" json:"nodeMeshMaxRestartTime,omitempty"`
		NodeMeshPassword       *struct {
			SecretKeyRef *struct {
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
		} `tfsdk:"node_mesh_password" json:"nodeMeshPassword,omitempty"`
		NodeToNodeMeshEnabled *bool `tfsdk:"node_to_node_mesh_enabled" json:"nodeToNodeMeshEnabled,omitempty"`
		PrefixAdvertisements  *[]struct {
			Cidr        *string   `tfsdk:"cidr" json:"cidr,omitempty"`
			Communities *[]string `tfsdk:"communities" json:"communities,omitempty"`
		} `tfsdk:"prefix_advertisements" json:"prefixAdvertisements,omitempty"`
		ServiceClusterIPs *[]struct {
			Cidr *string `tfsdk:"cidr" json:"cidr,omitempty"`
		} `tfsdk:"service_cluster_i_ps" json:"serviceClusterIPs,omitempty"`
		ServiceExternalIPs *[]struct {
			Cidr *string `tfsdk:"cidr" json:"cidr,omitempty"`
		} `tfsdk:"service_external_i_ps" json:"serviceExternalIPs,omitempty"`
		ServiceLoadBalancerIPs *[]struct {
			Cidr *string `tfsdk:"cidr" json:"cidr,omitempty"`
		} `tfsdk:"service_load_balancer_i_ps" json:"serviceLoadBalancerIPs,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CrdProjectcalicoOrgBgpconfigurationV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_crd_projectcalico_org_bgp_configuration_v1"
}

func (r *CrdProjectcalicoOrgBgpconfigurationV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "BGPConfiguration contains the configuration for any BGP routing.",
		MarkdownDescription: "BGPConfiguration contains the configuration for any BGP routing.",
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
				Description:         "BGPConfigurationSpec contains the values of the BGP configuration.",
				MarkdownDescription: "BGPConfigurationSpec contains the values of the BGP configuration.",
				Attributes: map[string]schema.Attribute{
					"as_number": schema.Int64Attribute{
						Description:         "ASNumber is the default AS number used by a node. [Default: 64512]",
						MarkdownDescription: "ASNumber is the default AS number used by a node. [Default: 64512]",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"bind_mode": schema.StringAttribute{
						Description:         "BindMode indicates whether to listen for BGP connections on all addresses (None) or only on the node's canonical IP address Node.Spec.BGP.IPvXAddress (NodeIP). Default behaviour is to listen for BGP connections on all addresses.",
						MarkdownDescription: "BindMode indicates whether to listen for BGP connections on all addresses (None) or only on the node's canonical IP address Node.Spec.BGP.IPvXAddress (NodeIP). Default behaviour is to listen for BGP connections on all addresses.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"communities": schema.ListNestedAttribute{
						Description:         "Communities is a list of BGP community values and their arbitrary names for tagging routes.",
						MarkdownDescription: "Communities is a list of BGP community values and their arbitrary names for tagging routes.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name given to community value.",
									MarkdownDescription: "Name given to community value.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"value": schema.StringAttribute{
									Description:         "Value must be of format 'aa:nn' or 'aa:nn:mm'. For standard community use 'aa:nn' format, where 'aa' and 'nn' are 16 bit number. For large community use 'aa:nn:mm' format, where 'aa', 'nn' and 'mm' are 32 bit number. Where, 'aa' is an AS Number, 'nn' and 'mm' are per-AS identifier.",
									MarkdownDescription: "Value must be of format 'aa:nn' or 'aa:nn:mm'. For standard community use 'aa:nn' format, where 'aa' and 'nn' are 16 bit number. For large community use 'aa:nn:mm' format, where 'aa', 'nn' and 'mm' are 32 bit number. Where, 'aa' is an AS Number, 'nn' and 'mm' are per-AS identifier.",
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

					"ignored_interfaces": schema.ListAttribute{
						Description:         "IgnoredInterfaces indicates the network interfaces that needs to be excluded when reading device routes.",
						MarkdownDescription: "IgnoredInterfaces indicates the network interfaces that needs to be excluded when reading device routes.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"listen_port": schema.Int64Attribute{
						Description:         "ListenPort is the port where BGP protocol should listen. Defaults to 179",
						MarkdownDescription: "ListenPort is the port where BGP protocol should listen. Defaults to 179",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"log_severity_screen": schema.StringAttribute{
						Description:         "LogSeverityScreen is the log severity above which logs are sent to the stdout. [Default: INFO]",
						MarkdownDescription: "LogSeverityScreen is the log severity above which logs are sent to the stdout. [Default: INFO]",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"node_mesh_max_restart_time": schema.StringAttribute{
						Description:         "Time to allow for software restart for node-to-mesh peerings.  When specified, this is configured as the graceful restart timeout.  When not specified, the BIRD default of 120s is used. This field can only be set on the default BGPConfiguration instance and requires that NodeMesh is enabled",
						MarkdownDescription: "Time to allow for software restart for node-to-mesh peerings.  When specified, this is configured as the graceful restart timeout.  When not specified, the BIRD default of 120s is used. This field can only be set on the default BGPConfiguration instance and requires that NodeMesh is enabled",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"node_mesh_password": schema.SingleNestedAttribute{
						Description:         "Optional BGP password for full node-to-mesh peerings. This field can only be set on the default BGPConfiguration instance and requires that NodeMesh is enabled",
						MarkdownDescription: "Optional BGP password for full node-to-mesh peerings. This field can only be set on the default BGPConfiguration instance and requires that NodeMesh is enabled",
						Attributes: map[string]schema.Attribute{
							"secret_key_ref": schema.SingleNestedAttribute{
								Description:         "Selects a key of a secret in the node pod's namespace.",
								MarkdownDescription: "Selects a key of a secret in the node pod's namespace.",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "The key of the secret to select from.  Must be a valid secret key.",
										MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"name": schema.StringAttribute{
										Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"optional": schema.BoolAttribute{
										Description:         "Specify whether the Secret or its key must be defined",
										MarkdownDescription: "Specify whether the Secret or its key must be defined",
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

					"node_to_node_mesh_enabled": schema.BoolAttribute{
						Description:         "NodeToNodeMeshEnabled sets whether full node to node BGP mesh is enabled. [Default: true]",
						MarkdownDescription: "NodeToNodeMeshEnabled sets whether full node to node BGP mesh is enabled. [Default: true]",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"prefix_advertisements": schema.ListNestedAttribute{
						Description:         "PrefixAdvertisements contains per-prefix advertisement configuration.",
						MarkdownDescription: "PrefixAdvertisements contains per-prefix advertisement configuration.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"cidr": schema.StringAttribute{
									Description:         "CIDR for which properties should be advertised.",
									MarkdownDescription: "CIDR for which properties should be advertised.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"communities": schema.ListAttribute{
									Description:         "Communities can be list of either community names already defined in 'Specs.Communities' or community value of format 'aa:nn' or 'aa:nn:mm'. For standard community use 'aa:nn' format, where 'aa' and 'nn' are 16 bit number. For large community use 'aa:nn:mm' format, where 'aa', 'nn' and 'mm' are 32 bit number. Where,'aa' is an AS Number, 'nn' and 'mm' are per-AS identifier.",
									MarkdownDescription: "Communities can be list of either community names already defined in 'Specs.Communities' or community value of format 'aa:nn' or 'aa:nn:mm'. For standard community use 'aa:nn' format, where 'aa' and 'nn' are 16 bit number. For large community use 'aa:nn:mm' format, where 'aa', 'nn' and 'mm' are 32 bit number. Where,'aa' is an AS Number, 'nn' and 'mm' are per-AS identifier.",
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

					"service_cluster_i_ps": schema.ListNestedAttribute{
						Description:         "ServiceClusterIPs are the CIDR blocks from which service cluster IPs are allocated. If specified, Calico will advertise these blocks, as well as any cluster IPs within them.",
						MarkdownDescription: "ServiceClusterIPs are the CIDR blocks from which service cluster IPs are allocated. If specified, Calico will advertise these blocks, as well as any cluster IPs within them.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"cidr": schema.StringAttribute{
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

					"service_external_i_ps": schema.ListNestedAttribute{
						Description:         "ServiceExternalIPs are the CIDR blocks for Kubernetes Service External IPs. Kubernetes Service ExternalIPs will only be advertised if they are within one of these blocks.",
						MarkdownDescription: "ServiceExternalIPs are the CIDR blocks for Kubernetes Service External IPs. Kubernetes Service ExternalIPs will only be advertised if they are within one of these blocks.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"cidr": schema.StringAttribute{
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

					"service_load_balancer_i_ps": schema.ListNestedAttribute{
						Description:         "ServiceLoadBalancerIPs are the CIDR blocks for Kubernetes Service LoadBalancer IPs. Kubernetes Service status.LoadBalancer.Ingress IPs will only be advertised if they are within one of these blocks.",
						MarkdownDescription: "ServiceLoadBalancerIPs are the CIDR blocks for Kubernetes Service LoadBalancer IPs. Kubernetes Service status.LoadBalancer.Ingress IPs will only be advertised if they are within one of these blocks.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"cidr": schema.StringAttribute{
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
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *CrdProjectcalicoOrgBgpconfigurationV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *CrdProjectcalicoOrgBgpconfigurationV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_crd_projectcalico_org_bgp_configuration_v1")

	var data CrdProjectcalicoOrgBgpconfigurationV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "crd.projectcalico.org", Version: "v1", Resource: "bgpconfigurations"}).
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

	var readResponse CrdProjectcalicoOrgBgpconfigurationV1DataSourceData
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
	data.ApiVersion = pointer.String("crd.projectcalico.org/v1")
	data.Kind = pointer.String("BGPConfiguration")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
