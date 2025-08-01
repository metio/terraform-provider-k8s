/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package cilium_io_v2alpha1

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
	_ datasource.DataSource = &CiliumIoCiliumGatewayClassConfigV2Alpha1Manifest{}
)

func NewCiliumIoCiliumGatewayClassConfigV2Alpha1Manifest() datasource.DataSource {
	return &CiliumIoCiliumGatewayClassConfigV2Alpha1Manifest{}
}

type CiliumIoCiliumGatewayClassConfigV2Alpha1Manifest struct{}

type CiliumIoCiliumGatewayClassConfigV2Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Description *string `tfsdk:"description" json:"description,omitempty"`
		Service     *struct {
			AllocateLoadBalancerNodePorts  *bool     `tfsdk:"allocate_load_balancer_node_ports" json:"allocateLoadBalancerNodePorts,omitempty"`
			ExternalTrafficPolicy          *string   `tfsdk:"external_traffic_policy" json:"externalTrafficPolicy,omitempty"`
			IpFamilies                     *[]string `tfsdk:"ip_families" json:"ipFamilies,omitempty"`
			IpFamilyPolicy                 *string   `tfsdk:"ip_family_policy" json:"ipFamilyPolicy,omitempty"`
			LoadBalancerClass              *string   `tfsdk:"load_balancer_class" json:"loadBalancerClass,omitempty"`
			LoadBalancerSourceRanges       *[]string `tfsdk:"load_balancer_source_ranges" json:"loadBalancerSourceRanges,omitempty"`
			LoadBalancerSourceRangesPolicy *string   `tfsdk:"load_balancer_source_ranges_policy" json:"loadBalancerSourceRangesPolicy,omitempty"`
			TrafficDistribution            *string   `tfsdk:"traffic_distribution" json:"trafficDistribution,omitempty"`
			Type                           *string   `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"service" json:"service,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CiliumIoCiliumGatewayClassConfigV2Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_cilium_io_cilium_gateway_class_config_v2alpha1_manifest"
}

func (r *CiliumIoCiliumGatewayClassConfigV2Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CiliumGatewayClassConfig is a Kubernetes third-party resource which is used to configure Gateways owned by GatewayClass.",
		MarkdownDescription: "CiliumGatewayClassConfig is a Kubernetes third-party resource which is used to configure Gateways owned by GatewayClass.",
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
				Description:         "Spec is a human-readable of a GatewayClass configuration.",
				MarkdownDescription: "Spec is a human-readable of a GatewayClass configuration.",
				Attributes: map[string]schema.Attribute{
					"description": schema.StringAttribute{
						Description:         "Description helps describe a GatewayClass configuration with more details.",
						MarkdownDescription: "Description helps describe a GatewayClass configuration with more details.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(64),
						},
					},

					"service": schema.SingleNestedAttribute{
						Description:         "Service specifies the configuration for the generated Service. Note that not all fields from upstream Service.Spec are supported",
						MarkdownDescription: "Service specifies the configuration for the generated Service. Note that not all fields from upstream Service.Spec are supported",
						Attributes: map[string]schema.Attribute{
							"allocate_load_balancer_node_ports": schema.BoolAttribute{
								Description:         "Sets the Service.Spec.AllocateLoadBalancerNodePorts in generated Service objects to the given value.",
								MarkdownDescription: "Sets the Service.Spec.AllocateLoadBalancerNodePorts in generated Service objects to the given value.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"external_traffic_policy": schema.StringAttribute{
								Description:         "Sets the Service.Spec.ExternalTrafficPolicy in generated Service objects to the given value.",
								MarkdownDescription: "Sets the Service.Spec.ExternalTrafficPolicy in generated Service objects to the given value.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ip_families": schema.ListAttribute{
								Description:         "Sets the Service.Spec.IPFamilies in generated Service objects to the given value.",
								MarkdownDescription: "Sets the Service.Spec.IPFamilies in generated Service objects to the given value.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ip_family_policy": schema.StringAttribute{
								Description:         "Sets the Service.Spec.IPFamilyPolicy in generated Service objects to the given value.",
								MarkdownDescription: "Sets the Service.Spec.IPFamilyPolicy in generated Service objects to the given value.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"load_balancer_class": schema.StringAttribute{
								Description:         "Sets the Service.Spec.LoadBalancerClass in generated Service objects to the given value.",
								MarkdownDescription: "Sets the Service.Spec.LoadBalancerClass in generated Service objects to the given value.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"load_balancer_source_ranges": schema.ListAttribute{
								Description:         "Sets the Service.Spec.LoadBalancerSourceRanges in generated Service objects to the given value.",
								MarkdownDescription: "Sets the Service.Spec.LoadBalancerSourceRanges in generated Service objects to the given value.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"load_balancer_source_ranges_policy": schema.StringAttribute{
								Description:         "LoadBalancerSourceRangesPolicy defines the policy for the LoadBalancerSourceRanges if the incoming traffic is allowed or denied.",
								MarkdownDescription: "LoadBalancerSourceRangesPolicy defines the policy for the LoadBalancerSourceRanges if the incoming traffic is allowed or denied.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Allow", "Deny"),
								},
							},

							"traffic_distribution": schema.StringAttribute{
								Description:         "Sets the Service.Spec.TrafficDistribution in generated Service objects to the given value.",
								MarkdownDescription: "Sets the Service.Spec.TrafficDistribution in generated Service objects to the given value.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "Sets the Service.Spec.Type in generated Service objects to the given value. Only LoadBalancer and NodePort are supported.",
								MarkdownDescription: "Sets the Service.Spec.Type in generated Service objects to the given value. Only LoadBalancer and NodePort are supported.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("LoadBalancer", "NodePort"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *CiliumIoCiliumGatewayClassConfigV2Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_cilium_io_cilium_gateway_class_config_v2alpha1_manifest")

	var model CiliumIoCiliumGatewayClassConfigV2Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("cilium.io/v2alpha1")
	model.Kind = pointer.String("CiliumGatewayClassConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
