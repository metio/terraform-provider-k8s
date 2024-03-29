/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package submariner_io_v1alpha1

import (
	"context"
	"fmt"
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
	_ datasource.DataSource = &SubmarinerIoSubmarinerV1Alpha1Manifest{}
)

func NewSubmarinerIoSubmarinerV1Alpha1Manifest() datasource.DataSource {
	return &SubmarinerIoSubmarinerV1Alpha1Manifest{}
}

type SubmarinerIoSubmarinerV1Alpha1Manifest struct{}

type SubmarinerIoSubmarinerV1Alpha1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
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
		AirGappedDeployment      *bool   `tfsdk:"air_gapped_deployment" json:"airGappedDeployment,omitempty"`
		Broker                   *string `tfsdk:"broker" json:"broker,omitempty"`
		BrokerK8sApiServer       *string `tfsdk:"broker_k8s_api_server" json:"brokerK8sApiServer,omitempty"`
		BrokerK8sApiServerToken  *string `tfsdk:"broker_k8s_api_server_token" json:"brokerK8sApiServerToken,omitempty"`
		BrokerK8sCA              *string `tfsdk:"broker_k8s_ca" json:"brokerK8sCA,omitempty"`
		BrokerK8sInsecure        *bool   `tfsdk:"broker_k8s_insecure" json:"brokerK8sInsecure,omitempty"`
		BrokerK8sRemoteNamespace *string `tfsdk:"broker_k8s_remote_namespace" json:"brokerK8sRemoteNamespace,omitempty"`
		BrokerK8sSecret          *string `tfsdk:"broker_k8s_secret" json:"brokerK8sSecret,omitempty"`
		CableDriver              *string `tfsdk:"cable_driver" json:"cableDriver,omitempty"`
		CeIPSecDebug             *bool   `tfsdk:"ce_ip_sec_debug" json:"ceIPSecDebug,omitempty"`
		CeIPSecForceUDPEncaps    *bool   `tfsdk:"ce_ip_sec_force_udp_encaps" json:"ceIPSecForceUDPEncaps,omitempty"`
		CeIPSecIKEPort           *int64  `tfsdk:"ce_ip_sec_ike_port" json:"ceIPSecIKEPort,omitempty"`
		CeIPSecNATTPort          *int64  `tfsdk:"ce_ip_sec_natt_port" json:"ceIPSecNATTPort,omitempty"`
		CeIPSecPSK               *string `tfsdk:"ce_ip_sec_psk" json:"ceIPSecPSK,omitempty"`
		CeIPSecPSKSecret         *string `tfsdk:"ce_ip_sec_psk_secret" json:"ceIPSecPSKSecret,omitempty"`
		CeIPSecPreferredServer   *bool   `tfsdk:"ce_ip_sec_preferred_server" json:"ceIPSecPreferredServer,omitempty"`
		ClusterCIDR              *string `tfsdk:"cluster_cidr" json:"clusterCIDR,omitempty"`
		ClusterID                *string `tfsdk:"cluster_id" json:"clusterID,omitempty"`
		ColorCodes               *string `tfsdk:"color_codes" json:"colorCodes,omitempty"`
		ConnectionHealthCheck    *struct {
			Enabled            *bool  `tfsdk:"enabled" json:"enabled,omitempty"`
			IntervalSeconds    *int64 `tfsdk:"interval_seconds" json:"intervalSeconds,omitempty"`
			MaxPacketLossCount *int64 `tfsdk:"max_packet_loss_count" json:"maxPacketLossCount,omitempty"`
		} `tfsdk:"connection_health_check" json:"connectionHealthCheck,omitempty"`
		CoreDNSCustomConfig *struct {
			ConfigMapName *string `tfsdk:"config_map_name" json:"configMapName,omitempty"`
			Namespace     *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"core_dns_custom_config" json:"coreDNSCustomConfig,omitempty"`
		CustomDomains           *[]string          `tfsdk:"custom_domains" json:"customDomains,omitempty"`
		Debug                   *bool              `tfsdk:"debug" json:"debug,omitempty"`
		GlobalCIDR              *string            `tfsdk:"global_cidr" json:"globalCIDR,omitempty"`
		HaltOnCertificateError  *bool              `tfsdk:"halt_on_certificate_error" json:"haltOnCertificateError,omitempty"`
		ImageOverrides          *map[string]string `tfsdk:"image_overrides" json:"imageOverrides,omitempty"`
		LoadBalancerEnabled     *bool              `tfsdk:"load_balancer_enabled" json:"loadBalancerEnabled,omitempty"`
		Namespace               *string            `tfsdk:"namespace" json:"namespace,omitempty"`
		NatEnabled              *bool              `tfsdk:"nat_enabled" json:"natEnabled,omitempty"`
		NodeSelector            *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
		Repository              *string            `tfsdk:"repository" json:"repository,omitempty"`
		ServiceCIDR             *string            `tfsdk:"service_cidr" json:"serviceCIDR,omitempty"`
		ServiceDiscoveryEnabled *bool              `tfsdk:"service_discovery_enabled" json:"serviceDiscoveryEnabled,omitempty"`
		Tolerations             *[]struct {
			Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
			Key               *string `tfsdk:"key" json:"key,omitempty"`
			Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
			TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
			Value             *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tolerations" json:"tolerations,omitempty"`
		Version *string `tfsdk:"version" json:"version,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SubmarinerIoSubmarinerV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_submariner_io_submariner_v1alpha1_manifest"
}

func (r *SubmarinerIoSubmarinerV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Submariner is the Schema for the submariners API.",
		MarkdownDescription: "Submariner is the Schema for the submariners API.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
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
				Description:         "SubmarinerSpec defines the desired state of Submariner.",
				MarkdownDescription: "SubmarinerSpec defines the desired state of Submariner.",
				Attributes: map[string]schema.Attribute{
					"air_gapped_deployment": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"broker": schema.StringAttribute{
						Description:         "Type of broker (must be 'k8s').",
						MarkdownDescription: "Type of broker (must be 'k8s').",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"broker_k8s_api_server": schema.StringAttribute{
						Description:         "The broker API URL.",
						MarkdownDescription: "The broker API URL.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"broker_k8s_api_server_token": schema.StringAttribute{
						Description:         "The broker API Token.",
						MarkdownDescription: "The broker API Token.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"broker_k8s_ca": schema.StringAttribute{
						Description:         "The broker certificate authority.",
						MarkdownDescription: "The broker certificate authority.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"broker_k8s_insecure": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"broker_k8s_remote_namespace": schema.StringAttribute{
						Description:         "The Broker namespace.",
						MarkdownDescription: "The Broker namespace.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"broker_k8s_secret": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"cable_driver": schema.StringAttribute{
						Description:         "Cable driver implementation - any of [libreswan, wireguard, vxlan].",
						MarkdownDescription: "Cable driver implementation - any of [libreswan, wireguard, vxlan].",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ce_ip_sec_debug": schema.BoolAttribute{
						Description:         "Enable logging IPsec debugging information.",
						MarkdownDescription: "Enable logging IPsec debugging information.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"ce_ip_sec_force_udp_encaps": schema.BoolAttribute{
						Description:         "Force UDP encapsulation for IPsec.",
						MarkdownDescription: "Force UDP encapsulation for IPsec.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ce_ip_sec_ike_port": schema.Int64Attribute{
						Description:         "The IPsec IKE port (500 usually).",
						MarkdownDescription: "The IPsec IKE port (500 usually).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ce_ip_sec_natt_port": schema.Int64Attribute{
						Description:         "The IPsec NAT traversal port (4500 usually).",
						MarkdownDescription: "The IPsec NAT traversal port (4500 usually).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ce_ip_sec_psk": schema.StringAttribute{
						Description:         "The IPsec Pre-Shared Key which must be identical in all route agents across the cluster.",
						MarkdownDescription: "The IPsec Pre-Shared Key which must be identical in all route agents across the cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ce_ip_sec_psk_secret": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ce_ip_sec_preferred_server": schema.BoolAttribute{
						Description:         "Enable this cluster as a preferred server for data-plane connections.",
						MarkdownDescription: "Enable this cluster as a preferred server for data-plane connections.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"cluster_cidr": schema.StringAttribute{
						Description:         "The cluster CIDR.",
						MarkdownDescription: "The cluster CIDR.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"cluster_id": schema.StringAttribute{
						Description:         "The cluster ID used to identify the tunnels.",
						MarkdownDescription: "The cluster ID used to identify the tunnels.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"color_codes": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"connection_health_check": schema.SingleNestedAttribute{
						Description:         "The gateway connection health check.",
						MarkdownDescription: "The gateway connection health check.",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "Enable the connection health check.",
								MarkdownDescription: "Enable the connection health check.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"interval_seconds": schema.Int64Attribute{
								Description:         "The interval at which health check pings are sent.",
								MarkdownDescription: "The interval at which health check pings are sent.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_packet_loss_count": schema.Int64Attribute{
								Description:         "The maximum number of packets lost at which the health checker will mark the connection as down.",
								MarkdownDescription: "The maximum number of packets lost at which the health checker will mark the connection as down.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"core_dns_custom_config": schema.SingleNestedAttribute{
						Description:         "Name of the custom CoreDNS configmap to configure forwarding to Lighthouse. It should be in <namespace>/<name> format where <namespace> is optional and defaults to kube-system.",
						MarkdownDescription: "Name of the custom CoreDNS configmap to configure forwarding to Lighthouse. It should be in <namespace>/<name> format where <namespace> is optional and defaults to kube-system.",
						Attributes: map[string]schema.Attribute{
							"config_map_name": schema.StringAttribute{
								Description:         "Name of the custom CoreDNS configmap.",
								MarkdownDescription: "Name of the custom CoreDNS configmap.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace of the custom CoreDNS configmap.",
								MarkdownDescription: "Namespace of the custom CoreDNS configmap.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"custom_domains": schema.ListAttribute{
						Description:         "List of domains to use for multi-cluster service discovery.",
						MarkdownDescription: "List of domains to use for multi-cluster service discovery.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"debug": schema.BoolAttribute{
						Description:         "Enable operator debugging.",
						MarkdownDescription: "Enable operator debugging.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"global_cidr": schema.StringAttribute{
						Description:         "The Global CIDR super-net range for allocating GlobalCIDRs to each cluster.",
						MarkdownDescription: "The Global CIDR super-net range for allocating GlobalCIDRs to each cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"halt_on_certificate_error": schema.BoolAttribute{
						Description:         "Halt on certificate error (so the pod gets restarted).",
						MarkdownDescription: "Halt on certificate error (so the pod gets restarted).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image_overrides": schema.MapAttribute{
						Description:         "Override component images.",
						MarkdownDescription: "Override component images.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"load_balancer_enabled": schema.BoolAttribute{
						Description:         "Enable automatic Load Balancer in front of the gateways.",
						MarkdownDescription: "Enable automatic Load Balancer in front of the gateways.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"namespace": schema.StringAttribute{
						Description:         "The namespace in which to deploy the submariner operator.",
						MarkdownDescription: "The namespace in which to deploy the submariner operator.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"nat_enabled": schema.BoolAttribute{
						Description:         "Enable NAT between clusters.",
						MarkdownDescription: "Enable NAT between clusters.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"node_selector": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"repository": schema.StringAttribute{
						Description:         "The image repository.",
						MarkdownDescription: "The image repository.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"service_cidr": schema.StringAttribute{
						Description:         "The service CIDR.",
						MarkdownDescription: "The service CIDR.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"service_discovery_enabled": schema.BoolAttribute{
						Description:         "Enable support for Service Discovery (Lighthouse).",
						MarkdownDescription: "Enable support for Service Discovery (Lighthouse).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tolerations": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"effect": schema.StringAttribute{
									Description:         "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
									MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"key": schema.StringAttribute{
									Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
									MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"operator": schema.StringAttribute{
									Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
									MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"toleration_seconds": schema.Int64Attribute{
									Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
									MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"value": schema.StringAttribute{
									Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
									MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
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

					"version": schema.StringAttribute{
						Description:         "The image tag.",
						MarkdownDescription: "The image tag.",
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

func (r *SubmarinerIoSubmarinerV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_submariner_io_submariner_v1alpha1_manifest")

	var model SubmarinerIoSubmarinerV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("submariner.io/v1alpha1")
	model.Kind = pointer.String("Submariner")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
