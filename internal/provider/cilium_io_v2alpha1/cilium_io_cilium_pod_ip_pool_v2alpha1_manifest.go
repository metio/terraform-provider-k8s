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
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &CiliumIoCiliumPodIppoolV2Alpha1Manifest{}
)

func NewCiliumIoCiliumPodIppoolV2Alpha1Manifest() datasource.DataSource {
	return &CiliumIoCiliumPodIppoolV2Alpha1Manifest{}
}

type CiliumIoCiliumPodIppoolV2Alpha1Manifest struct{}

type CiliumIoCiliumPodIppoolV2Alpha1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Ipv4 *struct {
			Cidrs    *[]string `tfsdk:"cidrs" json:"cidrs,omitempty"`
			MaskSize *int64    `tfsdk:"mask_size" json:"maskSize,omitempty"`
		} `tfsdk:"ipv4" json:"ipv4,omitempty"`
		Ipv6 *struct {
			Cidrs    *[]string `tfsdk:"cidrs" json:"cidrs,omitempty"`
			MaskSize *int64    `tfsdk:"mask_size" json:"maskSize,omitempty"`
		} `tfsdk:"ipv6" json:"ipv6,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CiliumIoCiliumPodIppoolV2Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_cilium_io_cilium_pod_ip_pool_v2alpha1_manifest"
}

func (r *CiliumIoCiliumPodIppoolV2Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CiliumPodIPPool defines an IP pool that can be used for pooled IPAM (i.e. the multi-pool IPAM mode).",
		MarkdownDescription: "CiliumPodIPPool defines an IP pool that can be used for pooled IPAM (i.e. the multi-pool IPAM mode).",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"ipv4": schema.SingleNestedAttribute{
						Description:         "IPv4 specifies the IPv4 CIDRs and mask sizes of the pool",
						MarkdownDescription: "IPv4 specifies the IPv4 CIDRs and mask sizes of the pool",
						Attributes: map[string]schema.Attribute{
							"cidrs": schema.ListAttribute{
								Description:         "CIDRs is a list of IPv4 CIDRs that are part of the pool.",
								MarkdownDescription: "CIDRs is a list of IPv4 CIDRs that are part of the pool.",
								ElementType:         types.StringType,
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"mask_size": schema.Int64Attribute{
								Description:         "MaskSize is the mask size of the pool.",
								MarkdownDescription: "MaskSize is the mask size of the pool.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(32),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"ipv6": schema.SingleNestedAttribute{
						Description:         "IPv6 specifies the IPv6 CIDRs and mask sizes of the pool",
						MarkdownDescription: "IPv6 specifies the IPv6 CIDRs and mask sizes of the pool",
						Attributes: map[string]schema.Attribute{
							"cidrs": schema.ListAttribute{
								Description:         "CIDRs is a list of IPv6 CIDRs that are part of the pool.",
								MarkdownDescription: "CIDRs is a list of IPv6 CIDRs that are part of the pool.",
								ElementType:         types.StringType,
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"mask_size": schema.Int64Attribute{
								Description:         "MaskSize is the mask size of the pool.",
								MarkdownDescription: "MaskSize is the mask size of the pool.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(128),
								},
							},
						},
						Required: false,
						Optional: true,
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

func (r *CiliumIoCiliumPodIppoolV2Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_cilium_io_cilium_pod_ip_pool_v2alpha1_manifest")

	var model CiliumIoCiliumPodIppoolV2Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("cilium.io/v2alpha1")
	model.Kind = pointer.String("CiliumPodIPPool")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
