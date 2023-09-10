/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package crd_projectcalico_org_v1

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
	_ datasource.DataSource              = &CrdProjectcalicoOrgIppoolV1DataSource{}
	_ datasource.DataSourceWithConfigure = &CrdProjectcalicoOrgIppoolV1DataSource{}
)

func NewCrdProjectcalicoOrgIppoolV1DataSource() datasource.DataSource {
	return &CrdProjectcalicoOrgIppoolV1DataSource{}
}

type CrdProjectcalicoOrgIppoolV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type CrdProjectcalicoOrgIppoolV1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		AllowedUses      *[]string `tfsdk:"allowed_uses" json:"allowedUses,omitempty"`
		BlockSize        *int64    `tfsdk:"block_size" json:"blockSize,omitempty"`
		Cidr             *string   `tfsdk:"cidr" json:"cidr,omitempty"`
		DisableBGPExport *bool     `tfsdk:"disable_bgp_export" json:"disableBGPExport,omitempty"`
		Disabled         *bool     `tfsdk:"disabled" json:"disabled,omitempty"`
		Ipip             *struct {
			Enabled *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
			Mode    *string `tfsdk:"mode" json:"mode,omitempty"`
		} `tfsdk:"ipip" json:"ipip,omitempty"`
		IpipMode     *string `tfsdk:"ipip_mode" json:"ipipMode,omitempty"`
		NatOutgoing  *bool   `tfsdk:"nat_outgoing" json:"natOutgoing,omitempty"`
		NodeSelector *string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
		VxlanMode    *string `tfsdk:"vxlan_mode" json:"vxlanMode,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CrdProjectcalicoOrgIppoolV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_crd_projectcalico_org_ip_pool_v1"
}

func (r *CrdProjectcalicoOrgIppoolV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "IPPoolSpec contains the specification for an IPPool resource.",
				MarkdownDescription: "IPPoolSpec contains the specification for an IPPool resource.",
				Attributes: map[string]schema.Attribute{
					"allowed_uses": schema.ListAttribute{
						Description:         "AllowedUse controls what the IP pool will be used for.  If not specified or empty, defaults to ['Tunnel', 'Workload'] for back-compatibility",
						MarkdownDescription: "AllowedUse controls what the IP pool will be used for.  If not specified or empty, defaults to ['Tunnel', 'Workload'] for back-compatibility",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"block_size": schema.Int64Attribute{
						Description:         "The block size to use for IP address assignments from this pool. Defaults to 26 for IPv4 and 122 for IPv6.",
						MarkdownDescription: "The block size to use for IP address assignments from this pool. Defaults to 26 for IPv4 and 122 for IPv6.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"cidr": schema.StringAttribute{
						Description:         "The pool CIDR.",
						MarkdownDescription: "The pool CIDR.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"disable_bgp_export": schema.BoolAttribute{
						Description:         "Disable exporting routes from this IP Pool's CIDR over BGP. [Default: false]",
						MarkdownDescription: "Disable exporting routes from this IP Pool's CIDR over BGP. [Default: false]",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"disabled": schema.BoolAttribute{
						Description:         "When disabled is true, Calico IPAM will not assign addresses from this pool.",
						MarkdownDescription: "When disabled is true, Calico IPAM will not assign addresses from this pool.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"ipip": schema.SingleNestedAttribute{
						Description:         "Deprecated: this field is only used for APIv1 backwards compatibility. Setting this field is not allowed, this field is for internal use only.",
						MarkdownDescription: "Deprecated: this field is only used for APIv1 backwards compatibility. Setting this field is not allowed, this field is for internal use only.",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "When enabled is true, ipip tunneling will be used to deliver packets to destinations within this pool.",
								MarkdownDescription: "When enabled is true, ipip tunneling will be used to deliver packets to destinations within this pool.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"mode": schema.StringAttribute{
								Description:         "The IPIP mode.  This can be one of 'always' or 'cross-subnet'.  A mode of 'always' will also use IPIP tunneling for routing to destination IP addresses within this pool.  A mode of 'cross-subnet' will only use IPIP tunneling when the destination node is on a different subnet to the originating node.  The default value (if not specified) is 'always'.",
								MarkdownDescription: "The IPIP mode.  This can be one of 'always' or 'cross-subnet'.  A mode of 'always' will also use IPIP tunneling for routing to destination IP addresses within this pool.  A mode of 'cross-subnet' will only use IPIP tunneling when the destination node is on a different subnet to the originating node.  The default value (if not specified) is 'always'.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"ipip_mode": schema.StringAttribute{
						Description:         "Contains configuration for IPIP tunneling for this pool. If not specified, then this is defaulted to 'Never' (i.e. IPIP tunneling is disabled).",
						MarkdownDescription: "Contains configuration for IPIP tunneling for this pool. If not specified, then this is defaulted to 'Never' (i.e. IPIP tunneling is disabled).",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"nat_outgoing": schema.BoolAttribute{
						Description:         "When natOutgoing is true, packets sent from Calico networked containers in this pool to destinations outside of this pool will be masqueraded.",
						MarkdownDescription: "When natOutgoing is true, packets sent from Calico networked containers in this pool to destinations outside of this pool will be masqueraded.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"node_selector": schema.StringAttribute{
						Description:         "Allows IPPool to allocate for a specific node by label selector.",
						MarkdownDescription: "Allows IPPool to allocate for a specific node by label selector.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"vxlan_mode": schema.StringAttribute{
						Description:         "Contains configuration for VXLAN tunneling for this pool. If not specified, then this is defaulted to 'Never' (i.e. VXLAN tunneling is disabled).",
						MarkdownDescription: "Contains configuration for VXLAN tunneling for this pool. If not specified, then this is defaulted to 'Never' (i.e. VXLAN tunneling is disabled).",
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
	}
}

func (r *CrdProjectcalicoOrgIppoolV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *CrdProjectcalicoOrgIppoolV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_crd_projectcalico_org_ip_pool_v1")

	var data CrdProjectcalicoOrgIppoolV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "crd.projectcalico.org", Version: "v1", Resource: "IPPool"}).
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

	var readResponse CrdProjectcalicoOrgIppoolV1DataSourceData
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
	data.Kind = pointer.String("IPPool")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
