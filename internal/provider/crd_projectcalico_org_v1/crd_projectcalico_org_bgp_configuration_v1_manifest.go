/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package crd_projectcalico_org_v1

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
	_ datasource.DataSource = &CrdProjectcalicoOrgBgpconfigurationV1Manifest{}
)

func NewCrdProjectcalicoOrgBgpconfigurationV1Manifest() datasource.DataSource {
	return &CrdProjectcalicoOrgBgpconfigurationV1Manifest{}
}

type CrdProjectcalicoOrgBgpconfigurationV1Manifest struct{}

type CrdProjectcalicoOrgBgpconfigurationV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

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

func (r *CrdProjectcalicoOrgBgpconfigurationV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_crd_projectcalico_org_bgp_configuration_v1_manifest"
}

func (r *CrdProjectcalicoOrgBgpconfigurationV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "BGPConfiguration contains the configuration for any BGP routing.",
		MarkdownDescription: "BGPConfiguration contains the configuration for any BGP routing.",
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
				Description:         "BGPConfigurationSpec contains the values of the BGP configuration.",
				MarkdownDescription: "BGPConfigurationSpec contains the values of the BGP configuration.",
				Attributes: map[string]schema.Attribute{
					"as_number": schema.Int64Attribute{
						Description:         "ASNumber is the default AS number used by a node. [Default: 64512]",
						MarkdownDescription: "ASNumber is the default AS number used by a node. [Default: 64512]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bind_mode": schema.StringAttribute{
						Description:         "BindMode indicates whether to listen for BGP connections on all addresses (None) or only on the node's canonical IP address Node.Spec.BGP.IPvXAddress (NodeIP). Default behaviour is to listen for BGP connections on all addresses.",
						MarkdownDescription: "BindMode indicates whether to listen for BGP connections on all addresses (None) or only on the node's canonical IP address Node.Spec.BGP.IPvXAddress (NodeIP). Default behaviour is to listen for BGP connections on all addresses.",
						Required:            false,
						Optional:            true,
						Computed:            false,
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
									Optional:            true,
									Computed:            false,
								},

								"value": schema.StringAttribute{
									Description:         "Value must be of format 'aa:nn' or 'aa:nn:mm'. For standard community use 'aa:nn' format, where 'aa' and 'nn' are 16 bit number. For large community use 'aa:nn:mm' format, where 'aa', 'nn' and 'mm' are 32 bit number. Where, 'aa' is an AS Number, 'nn' and 'mm' are per-AS identifier.",
									MarkdownDescription: "Value must be of format 'aa:nn' or 'aa:nn:mm'. For standard community use 'aa:nn' format, where 'aa' and 'nn' are 16 bit number. For large community use 'aa:nn:mm' format, where 'aa', 'nn' and 'mm' are 32 bit number. Where, 'aa' is an AS Number, 'nn' and 'mm' are per-AS identifier.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+):(\d+)$|^(\d+):(\d+):(\d+)$`), ""),
									},
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"ignored_interfaces": schema.ListAttribute{
						Description:         "IgnoredInterfaces indicates the network interfaces that needs to be excluded when reading device routes.",
						MarkdownDescription: "IgnoredInterfaces indicates the network interfaces that needs to be excluded when reading device routes.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"listen_port": schema.Int64Attribute{
						Description:         "ListenPort is the port where BGP protocol should listen. Defaults to 179",
						MarkdownDescription: "ListenPort is the port where BGP protocol should listen. Defaults to 179",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
							int64validator.AtMost(65535),
						},
					},

					"log_severity_screen": schema.StringAttribute{
						Description:         "LogSeverityScreen is the log severity above which logs are sent to the stdout. [Default: INFO]",
						MarkdownDescription: "LogSeverityScreen is the log severity above which logs are sent to the stdout. [Default: INFO]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"node_mesh_max_restart_time": schema.StringAttribute{
						Description:         "Time to allow for software restart for node-to-mesh peerings.  When specified, this is configured as the graceful restart timeout.  When not specified, the BIRD default of 120s is used. This field can only be set on the default BGPConfiguration instance and requires that NodeMesh is enabled",
						MarkdownDescription: "Time to allow for software restart for node-to-mesh peerings.  When specified, this is configured as the graceful restart timeout.  When not specified, the BIRD default of 120s is used. This field can only be set on the default BGPConfiguration instance and requires that NodeMesh is enabled",
						Required:            false,
						Optional:            true,
						Computed:            false,
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
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"optional": schema.BoolAttribute{
										Description:         "Specify whether the Secret or its key must be defined",
										MarkdownDescription: "Specify whether the Secret or its key must be defined",
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"node_to_node_mesh_enabled": schema.BoolAttribute{
						Description:         "NodeToNodeMeshEnabled sets whether full node to node BGP mesh is enabled. [Default: true]",
						MarkdownDescription: "NodeToNodeMeshEnabled sets whether full node to node BGP mesh is enabled. [Default: true]",
						Required:            false,
						Optional:            true,
						Computed:            false,
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
									Optional:            true,
									Computed:            false,
								},

								"communities": schema.ListAttribute{
									Description:         "Communities can be list of either community names already defined in 'Specs.Communities' or community value of format 'aa:nn' or 'aa:nn:mm'. For standard community use 'aa:nn' format, where 'aa' and 'nn' are 16 bit number. For large community use 'aa:nn:mm' format, where 'aa', 'nn' and 'mm' are 32 bit number. Where,'aa' is an AS Number, 'nn' and 'mm' are per-AS identifier.",
									MarkdownDescription: "Communities can be list of either community names already defined in 'Specs.Communities' or community value of format 'aa:nn' or 'aa:nn:mm'. For standard community use 'aa:nn' format, where 'aa' and 'nn' are 16 bit number. For large community use 'aa:nn:mm' format, where 'aa', 'nn' and 'mm' are 32 bit number. Where,'aa' is an AS Number, 'nn' and 'mm' are per-AS identifier.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
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
									Optional:            true,
									Computed:            false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
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
									Optional:            true,
									Computed:            false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
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
									Optional:            true,
									Computed:            false,
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

func (r *CrdProjectcalicoOrgBgpconfigurationV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_crd_projectcalico_org_bgp_configuration_v1_manifest")

	var model CrdProjectcalicoOrgBgpconfigurationV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("crd.projectcalico.org/v1")
	model.Kind = pointer.String("BGPConfiguration")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
