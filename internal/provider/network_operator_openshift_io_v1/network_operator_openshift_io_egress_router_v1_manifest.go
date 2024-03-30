/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package network_operator_openshift_io_v1

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
	_ datasource.DataSource = &NetworkOperatorOpenshiftIoEgressRouterV1Manifest{}
)

func NewNetworkOperatorOpenshiftIoEgressRouterV1Manifest() datasource.DataSource {
	return &NetworkOperatorOpenshiftIoEgressRouterV1Manifest{}
}

type NetworkOperatorOpenshiftIoEgressRouterV1Manifest struct{}

type NetworkOperatorOpenshiftIoEgressRouterV1ManifestData struct {
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
		Addresses *[]struct {
			Gateway *string `tfsdk:"gateway" json:"gateway,omitempty"`
			Ip      *string `tfsdk:"ip" json:"ip,omitempty"`
		} `tfsdk:"addresses" json:"addresses,omitempty"`
		Mode             *string `tfsdk:"mode" json:"mode,omitempty"`
		NetworkInterface *struct {
			Macvlan *struct {
				Master *string `tfsdk:"master" json:"master,omitempty"`
				Mode   *string `tfsdk:"mode" json:"mode,omitempty"`
			} `tfsdk:"macvlan" json:"macvlan,omitempty"`
		} `tfsdk:"network_interface" json:"networkInterface,omitempty"`
		Redirect *struct {
			FallbackIP    *string `tfsdk:"fallback_ip" json:"fallbackIP,omitempty"`
			RedirectRules *[]struct {
				DestinationIP *string `tfsdk:"destination_ip" json:"destinationIP,omitempty"`
				Port          *int64  `tfsdk:"port" json:"port,omitempty"`
				Protocol      *string `tfsdk:"protocol" json:"protocol,omitempty"`
				TargetPort    *int64  `tfsdk:"target_port" json:"targetPort,omitempty"`
			} `tfsdk:"redirect_rules" json:"redirectRules,omitempty"`
		} `tfsdk:"redirect" json:"redirect,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *NetworkOperatorOpenshiftIoEgressRouterV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_network_operator_openshift_io_egress_router_v1_manifest"
}

func (r *NetworkOperatorOpenshiftIoEgressRouterV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "EgressRouter is a feature allowing the user to define an egress router that acts as a bridge between pods and external systems. The egress router runs a service that redirects egress traffic originating from a pod or a group of pods to a remote external system or multiple destinations as per configuration.  It is consumed by the cluster-network-operator. More specifically, given an EgressRouter CR with <name>, the CNO will create and manage: - A service called <name> - An egress pod called <name> - A NAD called <name>  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).  EgressRouter is a single egressrouter pod configuration object.",
		MarkdownDescription: "EgressRouter is a feature allowing the user to define an egress router that acts as a bridge between pods and external systems. The egress router runs a service that redirects egress traffic originating from a pod or a group of pods to a remote external system or multiple destinations as per configuration.  It is consumed by the cluster-network-operator. More specifically, given an EgressRouter CR with <name>, the CNO will create and manage: - A service called <name> - An egress pod called <name> - A NAD called <name>  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).  EgressRouter is a single egressrouter pod configuration object.",
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
				Description:         "Specification of the desired egress router.",
				MarkdownDescription: "Specification of the desired egress router.",
				Attributes: map[string]schema.Attribute{
					"addresses": schema.ListNestedAttribute{
						Description:         "List of IP addresses to configure on the pod's secondary interface.",
						MarkdownDescription: "List of IP addresses to configure on the pod's secondary interface.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"gateway": schema.StringAttribute{
									Description:         "IP address of the next-hop gateway, if it cannot be automatically determined. Can be IPv4 or IPv6.",
									MarkdownDescription: "IP address of the next-hop gateway, if it cannot be automatically determined. Can be IPv4 or IPv6.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ip": schema.StringAttribute{
									Description:         "IP is the address to configure on the router's interface. Can be IPv4 or IPv6.",
									MarkdownDescription: "IP is the address to configure on the router's interface. Can be IPv4 or IPv6.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"mode": schema.StringAttribute{
						Description:         "Mode depicts the mode that is used for the egress router. The default mode is 'Redirect' and is the only supported mode currently.",
						MarkdownDescription: "Mode depicts the mode that is used for the egress router. The default mode is 'Redirect' and is the only supported mode currently.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Redirect"),
						},
					},

					"network_interface": schema.SingleNestedAttribute{
						Description:         "Specification of interface to create/use. The default is macvlan. Currently only macvlan is supported.",
						MarkdownDescription: "Specification of interface to create/use. The default is macvlan. Currently only macvlan is supported.",
						Attributes: map[string]schema.Attribute{
							"macvlan": schema.SingleNestedAttribute{
								Description:         "Arguments specific to the interfaceType macvlan",
								MarkdownDescription: "Arguments specific to the interfaceType macvlan",
								Attributes: map[string]schema.Attribute{
									"master": schema.StringAttribute{
										Description:         "Name of the master interface. Need not be specified if it can be inferred from the IP address.",
										MarkdownDescription: "Name of the master interface. Need not be specified if it can be inferred from the IP address.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"mode": schema.StringAttribute{
										Description:         "Mode depicts the mode that is used for the macvlan interface; one of Bridge|Private|VEPA|Passthru. The default mode is 'Bridge'.",
										MarkdownDescription: "Mode depicts the mode that is used for the macvlan interface; one of Bridge|Private|VEPA|Passthru. The default mode is 'Bridge'.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("Bridge", "Private", "VEPA", "Passthru"),
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

					"redirect": schema.SingleNestedAttribute{
						Description:         "Redirect represents the configuration parameters specific to redirect mode.",
						MarkdownDescription: "Redirect represents the configuration parameters specific to redirect mode.",
						Attributes: map[string]schema.Attribute{
							"fallback_ip": schema.StringAttribute{
								Description:         "FallbackIP specifies the remote destination's IP address. Can be IPv4 or IPv6. If no redirect rules are specified, all traffic from the router are redirected to this IP. If redirect rules are specified, then any connections on any other port (undefined in the rules) on the router will be redirected to this IP. If redirect rules are specified and no fallback IP is provided, connections on other ports will simply be rejected.",
								MarkdownDescription: "FallbackIP specifies the remote destination's IP address. Can be IPv4 or IPv6. If no redirect rules are specified, all traffic from the router are redirected to this IP. If redirect rules are specified, then any connections on any other port (undefined in the rules) on the router will be redirected to this IP. If redirect rules are specified and no fallback IP is provided, connections on other ports will simply be rejected.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"redirect_rules": schema.ListNestedAttribute{
								Description:         "List of L4RedirectRules that define the DNAT redirection from the pod to the destination in redirect mode.",
								MarkdownDescription: "List of L4RedirectRules that define the DNAT redirection from the pod to the destination in redirect mode.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"destination_ip": schema.StringAttribute{
											Description:         "IP specifies the remote destination's IP address. Can be IPv4 or IPv6.",
											MarkdownDescription: "IP specifies the remote destination's IP address. Can be IPv4 or IPv6.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"port": schema.Int64Attribute{
											Description:         "Port is the port number to which clients should send traffic to be redirected.",
											MarkdownDescription: "Port is the port number to which clients should send traffic to be redirected.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(1),
												int64validator.AtMost(65535),
											},
										},

										"protocol": schema.StringAttribute{
											Description:         "Protocol can be TCP, SCTP or UDP.",
											MarkdownDescription: "Protocol can be TCP, SCTP or UDP.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("TCP", "UDP", "SCTP"),
											},
										},

										"target_port": schema.Int64Attribute{
											Description:         "TargetPort allows specifying the port number on the remote destination to which the traffic gets redirected to. If unspecified, the value from 'Port' is used.",
											MarkdownDescription: "TargetPort allows specifying the port number on the remote destination to which the traffic gets redirected to. If unspecified, the value from 'Port' is used.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(1),
												int64validator.AtMost(65535),
											},
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
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *NetworkOperatorOpenshiftIoEgressRouterV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_network_operator_openshift_io_egress_router_v1_manifest")

	var model NetworkOperatorOpenshiftIoEgressRouterV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("network.operator.openshift.io/v1")
	model.Kind = pointer.String("EgressRouter")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
