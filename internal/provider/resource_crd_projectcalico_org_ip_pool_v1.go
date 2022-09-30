/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"gopkg.in/yaml.v3"
	"time"
)

type CrdProjectcalicoOrgIPPoolV1Resource struct{}

var (
	_ resource.Resource = (*CrdProjectcalicoOrgIPPoolV1Resource)(nil)
)

type CrdProjectcalicoOrgIPPoolV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type CrdProjectcalicoOrgIPPoolV1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		AllowedUses *[]string `tfsdk:"allowed_uses" yaml:"allowedUses,omitempty"`

		Cidr *string `tfsdk:"cidr" yaml:"cidr,omitempty"`

		DisableBGPExport *bool `tfsdk:"disable_bgp_export" yaml:"disableBGPExport,omitempty"`

		Ipip *struct {
			Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

			Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`
		} `tfsdk:"ipip" yaml:"ipip,omitempty"`

		IpipMode *string `tfsdk:"ipip_mode" yaml:"ipipMode,omitempty"`

		BlockSize *int64 `tfsdk:"block_size" yaml:"blockSize,omitempty"`

		Disabled *bool `tfsdk:"disabled" yaml:"disabled,omitempty"`

		Nat_outgoing *bool `tfsdk:"nat__outgoing" yaml:"nat-outgoing,omitempty"`

		NatOutgoing *bool `tfsdk:"nat_outgoing" yaml:"natOutgoing,omitempty"`

		NodeSelector *string `tfsdk:"node_selector" yaml:"nodeSelector,omitempty"`

		VxlanMode *string `tfsdk:"vxlan_mode" yaml:"vxlanMode,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewCrdProjectcalicoOrgIPPoolV1Resource() resource.Resource {
	return &CrdProjectcalicoOrgIPPoolV1Resource{}
}

func (r *CrdProjectcalicoOrgIPPoolV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_crd_projectcalico_org_ip_pool_v1"
}

func (r *CrdProjectcalicoOrgIPPoolV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "",
		MarkdownDescription: "",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description:         "The timestamp of the last change to this resource.",
				MarkdownDescription: "The timestamp of the last change to this resource.",
				Type:                types.Int64Type,
				Computed:            true,
				Optional:            false,
			},

			"yaml": {
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"metadata": {
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Type:                types.StringType,
						Required:            true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.RequiresReplace(),
						},
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
					},

					"labels": {
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.LabelValidator(),
						},
					},
					"annotations": {
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.AnnotationValidator(),
						},
					},
				}),
			},

			"api_version": {
				Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"kind": {
				Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"spec": {
				Description:         "IPPoolSpec contains the specification for an IPPool resource.",
				MarkdownDescription: "IPPoolSpec contains the specification for an IPPool resource.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"allowed_uses": {
						Description:         "AllowedUse controls what the IP pool will be used for.  If not specified or empty, defaults to ['Tunnel', 'Workload'] for back-compatibility",
						MarkdownDescription: "AllowedUse controls what the IP pool will be used for.  If not specified or empty, defaults to ['Tunnel', 'Workload'] for back-compatibility",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"cidr": {
						Description:         "The pool CIDR.",
						MarkdownDescription: "The pool CIDR.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"disable_bgp_export": {
						Description:         "Disable exporting routes from this IP Pool's CIDR over BGP. [Default: false]",
						MarkdownDescription: "Disable exporting routes from this IP Pool's CIDR over BGP. [Default: false]",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ipip": {
						Description:         "Deprecated: this field is only used for APIv1 backwards compatibility. Setting this field is not allowed, this field is for internal use only.",
						MarkdownDescription: "Deprecated: this field is only used for APIv1 backwards compatibility. Setting this field is not allowed, this field is for internal use only.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"enabled": {
								Description:         "When enabled is true, ipip tunneling will be used to deliver packets to destinations within this pool.",
								MarkdownDescription: "When enabled is true, ipip tunneling will be used to deliver packets to destinations within this pool.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"mode": {
								Description:         "The IPIP mode.  This can be one of 'always' or 'cross-subnet'.  A mode of 'always' will also use IPIP tunneling for routing to destination IP addresses within this pool.  A mode of 'cross-subnet' will only use IPIP tunneling when the destination node is on a different subnet to the originating node.  The default value (if not specified) is 'always'.",
								MarkdownDescription: "The IPIP mode.  This can be one of 'always' or 'cross-subnet'.  A mode of 'always' will also use IPIP tunneling for routing to destination IP addresses within this pool.  A mode of 'cross-subnet' will only use IPIP tunneling when the destination node is on a different subnet to the originating node.  The default value (if not specified) is 'always'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ipip_mode": {
						Description:         "Contains configuration for IPIP tunneling for this pool. If not specified, then this is defaulted to 'Never' (i.e. IPIP tunneling is disabled).",
						MarkdownDescription: "Contains configuration for IPIP tunneling for this pool. If not specified, then this is defaulted to 'Never' (i.e. IPIP tunneling is disabled).",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"block_size": {
						Description:         "The block size to use for IP address assignments from this pool. Defaults to 26 for IPv4 and 122 for IPv6.",
						MarkdownDescription: "The block size to use for IP address assignments from this pool. Defaults to 26 for IPv4 and 122 for IPv6.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"disabled": {
						Description:         "When disabled is true, Calico IPAM will not assign addresses from this pool.",
						MarkdownDescription: "When disabled is true, Calico IPAM will not assign addresses from this pool.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"nat__outgoing": {
						Description:         "Deprecated: this field is only used for APIv1 backwards compatibility. Setting this field is not allowed, this field is for internal use only.",
						MarkdownDescription: "Deprecated: this field is only used for APIv1 backwards compatibility. Setting this field is not allowed, this field is for internal use only.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"nat_outgoing": {
						Description:         "When nat-outgoing is true, packets sent from Calico networked containers in this pool to destinations outside of this pool will be masqueraded.",
						MarkdownDescription: "When nat-outgoing is true, packets sent from Calico networked containers in this pool to destinations outside of this pool will be masqueraded.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"node_selector": {
						Description:         "Allows IPPool to allocate for a specific node by label selector.",
						MarkdownDescription: "Allows IPPool to allocate for a specific node by label selector.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"vxlan_mode": {
						Description:         "Contains configuration for VXLAN tunneling for this pool. If not specified, then this is defaulted to 'Never' (i.e. VXLAN tunneling is disabled).",
						MarkdownDescription: "Contains configuration for VXLAN tunneling for this pool. If not specified, then this is defaulted to 'Never' (i.e. VXLAN tunneling is disabled).",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *CrdProjectcalicoOrgIPPoolV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_crd_projectcalico_org_ip_pool_v1")

	var state CrdProjectcalicoOrgIPPoolV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CrdProjectcalicoOrgIPPoolV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("crd.projectcalico.org/v1")
	goModel.Kind = utilities.Ptr("IPPool")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *CrdProjectcalicoOrgIPPoolV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_crd_projectcalico_org_ip_pool_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *CrdProjectcalicoOrgIPPoolV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_crd_projectcalico_org_ip_pool_v1")

	var state CrdProjectcalicoOrgIPPoolV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CrdProjectcalicoOrgIPPoolV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("crd.projectcalico.org/v1")
	goModel.Kind = utilities.Ptr("IPPool")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *CrdProjectcalicoOrgIPPoolV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_crd_projectcalico_org_ip_pool_v1")
	// NO-OP: Terraform removes the state automatically for us
}
