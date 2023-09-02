/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package chaos_mesh_org_v1alpha1

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
	_ datasource.DataSource              = &ChaosMeshOrgPodNetworkChaosV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &ChaosMeshOrgPodNetworkChaosV1Alpha1DataSource{}
)

func NewChaosMeshOrgPodNetworkChaosV1Alpha1DataSource() datasource.DataSource {
	return &ChaosMeshOrgPodNetworkChaosV1Alpha1DataSource{}
}

type ChaosMeshOrgPodNetworkChaosV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type ChaosMeshOrgPodNetworkChaosV1Alpha1DataSourceData struct {
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
		Ipsets *[]struct {
			CidrAndPorts *[]struct {
				Cidr *string `tfsdk:"cidr" json:"cidr,omitempty"`
				Port *int64  `tfsdk:"port" json:"port,omitempty"`
			} `tfsdk:"cidr_and_ports" json:"cidrAndPorts,omitempty"`
			Cidrs     *[]string `tfsdk:"cidrs" json:"cidrs,omitempty"`
			IpsetType *string   `tfsdk:"ipset_type" json:"ipsetType,omitempty"`
			Name      *string   `tfsdk:"name" json:"name,omitempty"`
			SetNames  *[]string `tfsdk:"set_names" json:"setNames,omitempty"`
			Source    *string   `tfsdk:"source" json:"source,omitempty"`
		} `tfsdk:"ipsets" json:"ipsets,omitempty"`
		Iptables *[]struct {
			Device    *string   `tfsdk:"device" json:"device,omitempty"`
			Direction *string   `tfsdk:"direction" json:"direction,omitempty"`
			Ipsets    *[]string `tfsdk:"ipsets" json:"ipsets,omitempty"`
			Name      *string   `tfsdk:"name" json:"name,omitempty"`
			Source    *string   `tfsdk:"source" json:"source,omitempty"`
		} `tfsdk:"iptables" json:"iptables,omitempty"`
		Tcs *[]struct {
			Bandwidth *struct {
				Buffer   *int64  `tfsdk:"buffer" json:"buffer,omitempty"`
				Limit    *int64  `tfsdk:"limit" json:"limit,omitempty"`
				Minburst *int64  `tfsdk:"minburst" json:"minburst,omitempty"`
				Peakrate *int64  `tfsdk:"peakrate" json:"peakrate,omitempty"`
				Rate     *string `tfsdk:"rate" json:"rate,omitempty"`
			} `tfsdk:"bandwidth" json:"bandwidth,omitempty"`
			Corrupt *struct {
				Correlation *string `tfsdk:"correlation" json:"correlation,omitempty"`
				Corrupt     *string `tfsdk:"corrupt" json:"corrupt,omitempty"`
			} `tfsdk:"corrupt" json:"corrupt,omitempty"`
			Delay *struct {
				Correlation *string `tfsdk:"correlation" json:"correlation,omitempty"`
				Jitter      *string `tfsdk:"jitter" json:"jitter,omitempty"`
				Latency     *string `tfsdk:"latency" json:"latency,omitempty"`
				Reorder     *struct {
					Correlation *string `tfsdk:"correlation" json:"correlation,omitempty"`
					Gap         *int64  `tfsdk:"gap" json:"gap,omitempty"`
					Reorder     *string `tfsdk:"reorder" json:"reorder,omitempty"`
				} `tfsdk:"reorder" json:"reorder,omitempty"`
			} `tfsdk:"delay" json:"delay,omitempty"`
			Device    *string `tfsdk:"device" json:"device,omitempty"`
			Duplicate *struct {
				Correlation *string `tfsdk:"correlation" json:"correlation,omitempty"`
				Duplicate   *string `tfsdk:"duplicate" json:"duplicate,omitempty"`
			} `tfsdk:"duplicate" json:"duplicate,omitempty"`
			Ipset *string `tfsdk:"ipset" json:"ipset,omitempty"`
			Loss  *struct {
				Correlation *string `tfsdk:"correlation" json:"correlation,omitempty"`
				Loss        *string `tfsdk:"loss" json:"loss,omitempty"`
			} `tfsdk:"loss" json:"loss,omitempty"`
			Source *string `tfsdk:"source" json:"source,omitempty"`
			Type   *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"tcs" json:"tcs,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ChaosMeshOrgPodNetworkChaosV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_chaos_mesh_org_pod_network_chaos_v1alpha1"
}

func (r *ChaosMeshOrgPodNetworkChaosV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "PodNetworkChaos is the Schema for the PodNetworkChaos API",
		MarkdownDescription: "PodNetworkChaos is the Schema for the PodNetworkChaos API",
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
				Description:         "Spec defines the behavior of a pod chaos experiment",
				MarkdownDescription: "Spec defines the behavior of a pod chaos experiment",
				Attributes: map[string]schema.Attribute{
					"ipsets": schema.ListNestedAttribute{
						Description:         "The ipset on the pod",
						MarkdownDescription: "The ipset on the pod",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"cidr_and_ports": schema.ListNestedAttribute{
									Description:         "The contents of ipset. Only available when IPSetType is NetPortIPSet.",
									MarkdownDescription: "The contents of ipset. Only available when IPSetType is NetPortIPSet.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"cidr": schema.StringAttribute{
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
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"cidrs": schema.ListAttribute{
									Description:         "The contents of ipset. Only available when IPSetType is NetIPSet.",
									MarkdownDescription: "The contents of ipset. Only available when IPSetType is NetIPSet.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"ipset_type": schema.StringAttribute{
									Description:         "IPSetType represents the type of IP set",
									MarkdownDescription: "IPSetType represents the type of IP set",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"name": schema.StringAttribute{
									Description:         "The name of ipset",
									MarkdownDescription: "The name of ipset",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"set_names": schema.ListAttribute{
									Description:         "The contents of ipset. Only available when IPSetType is SetIPSet.",
									MarkdownDescription: "The contents of ipset. Only available when IPSetType is SetIPSet.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"source": schema.StringAttribute{
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

					"iptables": schema.ListNestedAttribute{
						Description:         "The iptables rules on the pod",
						MarkdownDescription: "The iptables rules on the pod",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"device": schema.StringAttribute{
									Description:         "Device represents the network device to be affected.",
									MarkdownDescription: "Device represents the network device to be affected.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"direction": schema.StringAttribute{
									Description:         "The block direction of this iptables rule",
									MarkdownDescription: "The block direction of this iptables rule",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"ipsets": schema.ListAttribute{
									Description:         "The name of related ipset",
									MarkdownDescription: "The name of related ipset",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"name": schema.StringAttribute{
									Description:         "The name of iptables chain",
									MarkdownDescription: "The name of iptables chain",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"source": schema.StringAttribute{
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

					"tcs": schema.ListNestedAttribute{
						Description:         "The tc rules on the pod",
						MarkdownDescription: "The tc rules on the pod",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"bandwidth": schema.SingleNestedAttribute{
									Description:         "Bandwidth represents the detail about bandwidth control action",
									MarkdownDescription: "Bandwidth represents the detail about bandwidth control action",
									Attributes: map[string]schema.Attribute{
										"buffer": schema.Int64Attribute{
											Description:         "Buffer is the maximum amount of bytes that tokens can be available for instantaneously.",
											MarkdownDescription: "Buffer is the maximum amount of bytes that tokens can be available for instantaneously.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"limit": schema.Int64Attribute{
											Description:         "Limit is the number of bytes that can be queued waiting for tokens to become available.",
											MarkdownDescription: "Limit is the number of bytes that can be queued waiting for tokens to become available.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"minburst": schema.Int64Attribute{
											Description:         "Minburst specifies the size of the peakrate bucket. For perfect accuracy, should be set to the MTU of the interface.  If a peakrate is needed, but some burstiness is acceptable, this size can be raised. A 3000 byte minburst allows around 3mbit/s of peakrate, given 1000 byte packets.",
											MarkdownDescription: "Minburst specifies the size of the peakrate bucket. For perfect accuracy, should be set to the MTU of the interface.  If a peakrate is needed, but some burstiness is acceptable, this size can be raised. A 3000 byte minburst allows around 3mbit/s of peakrate, given 1000 byte packets.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"peakrate": schema.Int64Attribute{
											Description:         "Peakrate is the maximum depletion rate of the bucket. The peakrate does not need to be set, it is only necessary if perfect millisecond timescale shaping is required.",
											MarkdownDescription: "Peakrate is the maximum depletion rate of the bucket. The peakrate does not need to be set, it is only necessary if perfect millisecond timescale shaping is required.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"rate": schema.StringAttribute{
											Description:         "Rate is the speed knob. Allows bit, kbit, mbit, gbit, tbit, bps, kbps, mbps, gbps, tbps unit. bps means bytes per second.",
											MarkdownDescription: "Rate is the speed knob. Allows bit, kbit, mbit, gbit, tbit, bps, kbps, mbps, gbps, tbps unit. bps means bytes per second.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"corrupt": schema.SingleNestedAttribute{
									Description:         "Corrupt represents the detail about corrupt action",
									MarkdownDescription: "Corrupt represents the detail about corrupt action",
									Attributes: map[string]schema.Attribute{
										"correlation": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"corrupt": schema.StringAttribute{
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

								"delay": schema.SingleNestedAttribute{
									Description:         "Delay represents the detail about delay action",
									MarkdownDescription: "Delay represents the detail about delay action",
									Attributes: map[string]schema.Attribute{
										"correlation": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"jitter": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"latency": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"reorder": schema.SingleNestedAttribute{
											Description:         "ReorderSpec defines details of packet reorder.",
											MarkdownDescription: "ReorderSpec defines details of packet reorder.",
											Attributes: map[string]schema.Attribute{
												"correlation": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"gap": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"reorder": schema.StringAttribute{
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
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"device": schema.StringAttribute{
									Description:         "Device represents the network device to be affected.",
									MarkdownDescription: "Device represents the network device to be affected.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"duplicate": schema.SingleNestedAttribute{
									Description:         "DuplicateSpec represents the detail about loss action",
									MarkdownDescription: "DuplicateSpec represents the detail about loss action",
									Attributes: map[string]schema.Attribute{
										"correlation": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"duplicate": schema.StringAttribute{
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

								"ipset": schema.StringAttribute{
									Description:         "The name of target ipset",
									MarkdownDescription: "The name of target ipset",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"loss": schema.SingleNestedAttribute{
									Description:         "Loss represents the detail about loss action",
									MarkdownDescription: "Loss represents the detail about loss action",
									Attributes: map[string]schema.Attribute{
										"correlation": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"loss": schema.StringAttribute{
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

								"source": schema.StringAttribute{
									Description:         "The name and namespace of the source network chaos",
									MarkdownDescription: "The name and namespace of the source network chaos",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"type": schema.StringAttribute{
									Description:         "The type of traffic control",
									MarkdownDescription: "The type of traffic control",
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

func (r *ChaosMeshOrgPodNetworkChaosV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *ChaosMeshOrgPodNetworkChaosV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_chaos_mesh_org_pod_network_chaos_v1alpha1")

	var data ChaosMeshOrgPodNetworkChaosV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "chaos-mesh.org", Version: "v1alpha1", Resource: "PodNetworkChaos"}).
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

	var readResponse ChaosMeshOrgPodNetworkChaosV1Alpha1DataSourceData
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
	data.ApiVersion = pointer.String("chaos-mesh.org/v1alpha1")
	data.Kind = pointer.String("PodNetworkChaos")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}