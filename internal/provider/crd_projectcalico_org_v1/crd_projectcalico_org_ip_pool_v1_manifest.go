/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package crd_projectcalico_org_v1

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
	_ datasource.DataSource = &CrdProjectcalicoOrgIppoolV1Manifest{}
)

func NewCrdProjectcalicoOrgIppoolV1Manifest() datasource.DataSource {
	return &CrdProjectcalicoOrgIppoolV1Manifest{}
}

type CrdProjectcalicoOrgIppoolV1Manifest struct{}

type CrdProjectcalicoOrgIppoolV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

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

func (r *CrdProjectcalicoOrgIppoolV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_crd_projectcalico_org_ip_pool_v1_manifest"
}

func (r *CrdProjectcalicoOrgIppoolV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
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
				Description:         "IPPoolSpec contains the specification for an IPPool resource.",
				MarkdownDescription: "IPPoolSpec contains the specification for an IPPool resource.",
				Attributes: map[string]schema.Attribute{
					"allowed_uses": schema.ListAttribute{
						Description:         "AllowedUse controls what the IP pool will be used for.  If not specified or empty, defaults to ['Tunnel', 'Workload'] for back-compatibility",
						MarkdownDescription: "AllowedUse controls what the IP pool will be used for.  If not specified or empty, defaults to ['Tunnel', 'Workload'] for back-compatibility",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"block_size": schema.Int64Attribute{
						Description:         "The block size to use for IP address assignments from this pool. Defaults to 26 for IPv4 and 122 for IPv6.",
						MarkdownDescription: "The block size to use for IP address assignments from this pool. Defaults to 26 for IPv4 and 122 for IPv6.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"cidr": schema.StringAttribute{
						Description:         "The pool CIDR.",
						MarkdownDescription: "The pool CIDR.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"disable_bgp_export": schema.BoolAttribute{
						Description:         "Disable exporting routes from this IP Pool's CIDR over BGP. [Default: false]",
						MarkdownDescription: "Disable exporting routes from this IP Pool's CIDR over BGP. [Default: false]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"disabled": schema.BoolAttribute{
						Description:         "When disabled is true, Calico IPAM will not assign addresses from this pool.",
						MarkdownDescription: "When disabled is true, Calico IPAM will not assign addresses from this pool.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ipip": schema.SingleNestedAttribute{
						Description:         "Deprecated: this field is only used for APIv1 backwards compatibility. Setting this field is not allowed, this field is for internal use only.",
						MarkdownDescription: "Deprecated: this field is only used for APIv1 backwards compatibility. Setting this field is not allowed, this field is for internal use only.",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "When enabled is true, ipip tunneling will be used to deliver packets to destinations within this pool.",
								MarkdownDescription: "When enabled is true, ipip tunneling will be used to deliver packets to destinations within this pool.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"mode": schema.StringAttribute{
								Description:         "The IPIP mode.  This can be one of 'always' or 'cross-subnet'.  A mode of 'always' will also use IPIP tunneling for routing to destination IP addresses within this pool.  A mode of 'cross-subnet' will only use IPIP tunneling when the destination node is on a different subnet to the originating node.  The default value (if not specified) is 'always'.",
								MarkdownDescription: "The IPIP mode.  This can be one of 'always' or 'cross-subnet'.  A mode of 'always' will also use IPIP tunneling for routing to destination IP addresses within this pool.  A mode of 'cross-subnet' will only use IPIP tunneling when the destination node is on a different subnet to the originating node.  The default value (if not specified) is 'always'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"ipip_mode": schema.StringAttribute{
						Description:         "Contains configuration for IPIP tunneling for this pool. If not specified, then this is defaulted to 'Never' (i.e. IPIP tunneling is disabled).",
						MarkdownDescription: "Contains configuration for IPIP tunneling for this pool. If not specified, then this is defaulted to 'Never' (i.e. IPIP tunneling is disabled).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"nat_outgoing": schema.BoolAttribute{
						Description:         "When natOutgoing is true, packets sent from Calico networked containers in this pool to destinations outside of this pool will be masqueraded.",
						MarkdownDescription: "When natOutgoing is true, packets sent from Calico networked containers in this pool to destinations outside of this pool will be masqueraded.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"node_selector": schema.StringAttribute{
						Description:         "Allows IPPool to allocate for a specific node by label selector.",
						MarkdownDescription: "Allows IPPool to allocate for a specific node by label selector.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"vxlan_mode": schema.StringAttribute{
						Description:         "Contains configuration for VXLAN tunneling for this pool. If not specified, then this is defaulted to 'Never' (i.e. VXLAN tunneling is disabled).",
						MarkdownDescription: "Contains configuration for VXLAN tunneling for this pool. If not specified, then this is defaulted to 'Never' (i.e. VXLAN tunneling is disabled).",
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

func (r *CrdProjectcalicoOrgIppoolV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_crd_projectcalico_org_ip_pool_v1_manifest")

	var model CrdProjectcalicoOrgIppoolV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("crd.projectcalico.org/v1")
	model.Kind = pointer.String("IPPool")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
