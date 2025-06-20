/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package cilium_io_v2alpha1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &CiliumIoCiliumBgpnodeConfigOverrideV2Alpha1Manifest{}
)

func NewCiliumIoCiliumBgpnodeConfigOverrideV2Alpha1Manifest() datasource.DataSource {
	return &CiliumIoCiliumBgpnodeConfigOverrideV2Alpha1Manifest{}
}

type CiliumIoCiliumBgpnodeConfigOverrideV2Alpha1Manifest struct{}

type CiliumIoCiliumBgpnodeConfigOverrideV2Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		BgpInstances *[]struct {
			LocalASN  *int64  `tfsdk:"local_asn" json:"localASN,omitempty"`
			LocalPort *int64  `tfsdk:"local_port" json:"localPort,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Peers     *[]struct {
				LocalAddress *string `tfsdk:"local_address" json:"localAddress,omitempty"`
				LocalPort    *int64  `tfsdk:"local_port" json:"localPort,omitempty"`
				Name         *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"peers" json:"peers,omitempty"`
			RouterID *string `tfsdk:"router_id" json:"routerID,omitempty"`
		} `tfsdk:"bgp_instances" json:"bgpInstances,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CiliumIoCiliumBgpnodeConfigOverrideV2Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_cilium_io_cilium_bgp_node_config_override_v2alpha1_manifest"
}

func (r *CiliumIoCiliumBgpnodeConfigOverrideV2Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CiliumBGPNodeConfigOverride specifies configuration overrides for a CiliumBGPNodeConfig. It allows fine-tuning of BGP behavior on a per-node basis. For the override to be effective, the names in CiliumBGPNodeConfigOverride and CiliumBGPNodeConfig must match exactly. This matching ensures that specific node configurations are applied correctly and only where intended.",
		MarkdownDescription: "CiliumBGPNodeConfigOverride specifies configuration overrides for a CiliumBGPNodeConfig. It allows fine-tuning of BGP behavior on a per-node basis. For the override to be effective, the names in CiliumBGPNodeConfigOverride and CiliumBGPNodeConfig must match exactly. This matching ensures that specific node configurations are applied correctly and only where intended.",
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
				Description:         "Spec is the specification of the desired behavior of the CiliumBGPNodeConfigOverride.",
				MarkdownDescription: "Spec is the specification of the desired behavior of the CiliumBGPNodeConfigOverride.",
				Attributes: map[string]schema.Attribute{
					"bgp_instances": schema.ListNestedAttribute{
						Description:         "BGPInstances is a list of BGP instances to override.",
						MarkdownDescription: "BGPInstances is a list of BGP instances to override.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"local_asn": schema.Int64Attribute{
									Description:         "LocalASN is the ASN to use for this BGP instance.",
									MarkdownDescription: "LocalASN is the ASN to use for this BGP instance.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.Int64{
										int64validator.AtLeast(1),
										int64validator.AtMost(4.294967295e+09),
									},
								},

								"local_port": schema.Int64Attribute{
									Description:         "LocalPort is port to use for this BGP instance.",
									MarkdownDescription: "LocalPort is port to use for this BGP instance.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name is the name of the BGP instance for which the configuration is overridden.",
									MarkdownDescription: "Name is the name of the BGP instance for which the configuration is overridden.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(255),
									},
								},

								"peers": schema.ListNestedAttribute{
									Description:         "Peers is a list of peer configurations to override.",
									MarkdownDescription: "Peers is a list of peer configurations to override.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"local_address": schema.StringAttribute{
												Description:         "LocalAddress is the IP address to use for connecting to this peer.",
												MarkdownDescription: "LocalAddress is the IP address to use for connecting to this peer.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`((^\s*((([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]))\s*$)|(^\s*((([0-9A-Fa-f]{1,4}:){7}([0-9A-Fa-f]{1,4}|:))|(([0-9A-Fa-f]{1,4}:){6}(:[0-9A-Fa-f]{1,4}|((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){1,2})|:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){1,3})|((:[0-9A-Fa-f]{1,4})?:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){1,4})|((:[0-9A-Fa-f]{1,4}){0,2}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){1,5})|((:[0-9A-Fa-f]{1,4}){0,3}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){1,6})|((:[0-9A-Fa-f]{1,4}){0,4}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(:(((:[0-9A-Fa-f]{1,4}){1,7})|((:[0-9A-Fa-f]{1,4}){0,5}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:)))(%.+)?\s*$))`), ""),
												},
											},

											"local_port": schema.Int64Attribute{
												Description:         "LocalPort is source port to use for connecting to this peer.",
												MarkdownDescription: "LocalPort is source port to use for connecting to this peer.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name is the name of the peer for which the configuration is overridden.",
												MarkdownDescription: "Name is the name of the peer for which the configuration is overridden.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
													stringvalidator.LengthAtMost(255),
												},
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"router_id": schema.StringAttribute{
									Description:         "RouterID is BGP router id to use for this instance. It must be unique across all BGP instances.",
									MarkdownDescription: "RouterID is BGP router id to use for this instance. It must be unique across all BGP instances.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *CiliumIoCiliumBgpnodeConfigOverrideV2Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_cilium_io_cilium_bgp_node_config_override_v2alpha1_manifest")

	var model CiliumIoCiliumBgpnodeConfigOverrideV2Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("cilium.io/v2alpha1")
	model.Kind = pointer.String("CiliumBGPNodeConfigOverride")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
