/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package config_openshift_io_v1

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
	_ datasource.DataSource = &ConfigOpenshiftIoInfrastructureV1Manifest{}
)

func NewConfigOpenshiftIoInfrastructureV1Manifest() datasource.DataSource {
	return &ConfigOpenshiftIoInfrastructureV1Manifest{}
}

type ConfigOpenshiftIoInfrastructureV1Manifest struct{}

type ConfigOpenshiftIoInfrastructureV1ManifestData struct {
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
		CloudConfig *struct {
			Key  *string `tfsdk:"key" json:"key,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"cloud_config" json:"cloudConfig,omitempty"`
		PlatformSpec *struct {
			AlibabaCloud *map[string]string `tfsdk:"alibaba_cloud" json:"alibabaCloud,omitempty"`
			Aws          *struct {
				ServiceEndpoints *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
					Url  *string `tfsdk:"url" json:"url,omitempty"`
				} `tfsdk:"service_endpoints" json:"serviceEndpoints,omitempty"`
			} `tfsdk:"aws" json:"aws,omitempty"`
			Azure     *map[string]string `tfsdk:"azure" json:"azure,omitempty"`
			Baremetal *struct {
				ApiServerInternalIPs *[]string `tfsdk:"api_server_internal_i_ps" json:"apiServerInternalIPs,omitempty"`
				IngressIPs           *[]string `tfsdk:"ingress_i_ps" json:"ingressIPs,omitempty"`
				MachineNetworks      *[]string `tfsdk:"machine_networks" json:"machineNetworks,omitempty"`
			} `tfsdk:"baremetal" json:"baremetal,omitempty"`
			EquinixMetal *map[string]string `tfsdk:"equinix_metal" json:"equinixMetal,omitempty"`
			External     *struct {
				PlatformName *string `tfsdk:"platform_name" json:"platformName,omitempty"`
			} `tfsdk:"external" json:"external,omitempty"`
			Gcp      *map[string]string `tfsdk:"gcp" json:"gcp,omitempty"`
			Ibmcloud *map[string]string `tfsdk:"ibmcloud" json:"ibmcloud,omitempty"`
			Kubevirt *map[string]string `tfsdk:"kubevirt" json:"kubevirt,omitempty"`
			Nutanix  *struct {
				FailureDomains *[]struct {
					Cluster *struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
						Type *string `tfsdk:"type" json:"type,omitempty"`
						Uuid *string `tfsdk:"uuid" json:"uuid,omitempty"`
					} `tfsdk:"cluster" json:"cluster,omitempty"`
					Name    *string `tfsdk:"name" json:"name,omitempty"`
					Subnets *[]struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
						Type *string `tfsdk:"type" json:"type,omitempty"`
						Uuid *string `tfsdk:"uuid" json:"uuid,omitempty"`
					} `tfsdk:"subnets" json:"subnets,omitempty"`
				} `tfsdk:"failure_domains" json:"failureDomains,omitempty"`
				PrismCentral *struct {
					Address *string `tfsdk:"address" json:"address,omitempty"`
					Port    *int64  `tfsdk:"port" json:"port,omitempty"`
				} `tfsdk:"prism_central" json:"prismCentral,omitempty"`
				PrismElements *[]struct {
					Endpoint *struct {
						Address *string `tfsdk:"address" json:"address,omitempty"`
						Port    *int64  `tfsdk:"port" json:"port,omitempty"`
					} `tfsdk:"endpoint" json:"endpoint,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"prism_elements" json:"prismElements,omitempty"`
			} `tfsdk:"nutanix" json:"nutanix,omitempty"`
			Openstack *struct {
				ApiServerInternalIPs *[]string `tfsdk:"api_server_internal_i_ps" json:"apiServerInternalIPs,omitempty"`
				IngressIPs           *[]string `tfsdk:"ingress_i_ps" json:"ingressIPs,omitempty"`
				MachineNetworks      *[]string `tfsdk:"machine_networks" json:"machineNetworks,omitempty"`
			} `tfsdk:"openstack" json:"openstack,omitempty"`
			Ovirt   *map[string]string `tfsdk:"ovirt" json:"ovirt,omitempty"`
			Powervs *struct {
				ServiceEndpoints *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
					Url  *string `tfsdk:"url" json:"url,omitempty"`
				} `tfsdk:"service_endpoints" json:"serviceEndpoints,omitempty"`
			} `tfsdk:"powervs" json:"powervs,omitempty"`
			Type    *string `tfsdk:"type" json:"type,omitempty"`
			Vsphere *struct {
				ApiServerInternalIPs *[]string `tfsdk:"api_server_internal_i_ps" json:"apiServerInternalIPs,omitempty"`
				FailureDomains       *[]struct {
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Region   *string `tfsdk:"region" json:"region,omitempty"`
					Server   *string `tfsdk:"server" json:"server,omitempty"`
					Topology *struct {
						ComputeCluster *string   `tfsdk:"compute_cluster" json:"computeCluster,omitempty"`
						Datacenter     *string   `tfsdk:"datacenter" json:"datacenter,omitempty"`
						Datastore      *string   `tfsdk:"datastore" json:"datastore,omitempty"`
						Folder         *string   `tfsdk:"folder" json:"folder,omitempty"`
						Networks       *[]string `tfsdk:"networks" json:"networks,omitempty"`
						ResourcePool   *string   `tfsdk:"resource_pool" json:"resourcePool,omitempty"`
						Template       *string   `tfsdk:"template" json:"template,omitempty"`
					} `tfsdk:"topology" json:"topology,omitempty"`
					Zone *string `tfsdk:"zone" json:"zone,omitempty"`
				} `tfsdk:"failure_domains" json:"failureDomains,omitempty"`
				IngressIPs      *[]string `tfsdk:"ingress_i_ps" json:"ingressIPs,omitempty"`
				MachineNetworks *[]string `tfsdk:"machine_networks" json:"machineNetworks,omitempty"`
				NodeNetworking  *struct {
					External *struct {
						ExcludeNetworkSubnetCidr *[]string `tfsdk:"exclude_network_subnet_cidr" json:"excludeNetworkSubnetCidr,omitempty"`
						Network                  *string   `tfsdk:"network" json:"network,omitempty"`
						NetworkSubnetCidr        *[]string `tfsdk:"network_subnet_cidr" json:"networkSubnetCidr,omitempty"`
					} `tfsdk:"external" json:"external,omitempty"`
					Internal *struct {
						ExcludeNetworkSubnetCidr *[]string `tfsdk:"exclude_network_subnet_cidr" json:"excludeNetworkSubnetCidr,omitempty"`
						Network                  *string   `tfsdk:"network" json:"network,omitempty"`
						NetworkSubnetCidr        *[]string `tfsdk:"network_subnet_cidr" json:"networkSubnetCidr,omitempty"`
					} `tfsdk:"internal" json:"internal,omitempty"`
				} `tfsdk:"node_networking" json:"nodeNetworking,omitempty"`
				Vcenters *[]struct {
					Datacenters *[]string `tfsdk:"datacenters" json:"datacenters,omitempty"`
					Port        *int64    `tfsdk:"port" json:"port,omitempty"`
					Server      *string   `tfsdk:"server" json:"server,omitempty"`
				} `tfsdk:"vcenters" json:"vcenters,omitempty"`
			} `tfsdk:"vsphere" json:"vsphere,omitempty"`
		} `tfsdk:"platform_spec" json:"platformSpec,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ConfigOpenshiftIoInfrastructureV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_config_openshift_io_infrastructure_v1_manifest"
}

func (r *ConfigOpenshiftIoInfrastructureV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Infrastructure holds cluster-wide information about Infrastructure.  The canonical name is 'cluster'  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		MarkdownDescription: "Infrastructure holds cluster-wide information about Infrastructure.  The canonical name is 'cluster'  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
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
				Description:         "spec holds user settable values for configuration",
				MarkdownDescription: "spec holds user settable values for configuration",
				Attributes: map[string]schema.Attribute{
					"cloud_config": schema.SingleNestedAttribute{
						Description:         "cloudConfig is a reference to a ConfigMap containing the cloud provider configuration file. This configuration file is used to configure the Kubernetes cloud provider integration when using the built-in cloud provider integration or the external cloud controller manager. The namespace for this config map is openshift-config.  cloudConfig should only be consumed by the kube_cloud_config controller. The controller is responsible for using the user configuration in the spec for various platforms and combining that with the user provided ConfigMap in this field to create a stitched kube cloud config. The controller generates a ConfigMap 'kube-cloud-config' in 'openshift-config-managed' namespace with the kube cloud config is stored in 'cloud.conf' key. All the clients are expected to use the generated ConfigMap only.",
						MarkdownDescription: "cloudConfig is a reference to a ConfigMap containing the cloud provider configuration file. This configuration file is used to configure the Kubernetes cloud provider integration when using the built-in cloud provider integration or the external cloud controller manager. The namespace for this config map is openshift-config.  cloudConfig should only be consumed by the kube_cloud_config controller. The controller is responsible for using the user configuration in the spec for various platforms and combining that with the user provided ConfigMap in this field to create a stitched kube cloud config. The controller generates a ConfigMap 'kube-cloud-config' in 'openshift-config-managed' namespace with the kube cloud config is stored in 'cloud.conf' key. All the clients are expected to use the generated ConfigMap only.",
						Attributes: map[string]schema.Attribute{
							"key": schema.StringAttribute{
								Description:         "Key allows pointing to a specific key/value inside of the configmap.  This is useful for logical file references.",
								MarkdownDescription: "Key allows pointing to a specific key/value inside of the configmap.  This is useful for logical file references.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"platform_spec": schema.SingleNestedAttribute{
						Description:         "platformSpec holds desired information specific to the underlying infrastructure provider.",
						MarkdownDescription: "platformSpec holds desired information specific to the underlying infrastructure provider.",
						Attributes: map[string]schema.Attribute{
							"alibaba_cloud": schema.MapAttribute{
								Description:         "AlibabaCloud contains settings specific to the Alibaba Cloud infrastructure provider.",
								MarkdownDescription: "AlibabaCloud contains settings specific to the Alibaba Cloud infrastructure provider.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"aws": schema.SingleNestedAttribute{
								Description:         "AWS contains settings specific to the Amazon Web Services infrastructure provider.",
								MarkdownDescription: "AWS contains settings specific to the Amazon Web Services infrastructure provider.",
								Attributes: map[string]schema.Attribute{
									"service_endpoints": schema.ListNestedAttribute{
										Description:         "serviceEndpoints list contains custom endpoints which will override default service endpoint of AWS Services. There must be only one ServiceEndpoint for a service.",
										MarkdownDescription: "serviceEndpoints list contains custom endpoints which will override default service endpoint of AWS Services. There must be only one ServiceEndpoint for a service.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "name is the name of the AWS service. The list of all the service names can be found at https://docs.aws.amazon.com/general/latest/gr/aws-service-information.html This must be provided and cannot be empty.",
													MarkdownDescription: "name is the name of the AWS service. The list of all the service names can be found at https://docs.aws.amazon.com/general/latest/gr/aws-service-information.html This must be provided and cannot be empty.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9-]+$`), ""),
													},
												},

												"url": schema.StringAttribute{
													Description:         "url is fully qualified URI with scheme https, that overrides the default generated endpoint for a client. This must be provided and cannot be empty.",
													MarkdownDescription: "url is fully qualified URI with scheme https, that overrides the default generated endpoint for a client. This must be provided and cannot be empty.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`^https://`), ""),
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

							"azure": schema.MapAttribute{
								Description:         "Azure contains settings specific to the Azure infrastructure provider.",
								MarkdownDescription: "Azure contains settings specific to the Azure infrastructure provider.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"baremetal": schema.SingleNestedAttribute{
								Description:         "BareMetal contains settings specific to the BareMetal platform.",
								MarkdownDescription: "BareMetal contains settings specific to the BareMetal platform.",
								Attributes: map[string]schema.Attribute{
									"api_server_internal_i_ps": schema.ListAttribute{
										Description:         "apiServerInternalIPs are the IP addresses to contact the Kubernetes API server that can be used by components inside the cluster, like kubelets using the infrastructure rather than Kubernetes networking. These are the IPs for a self-hosted load balancer in front of the API servers. In dual stack clusters this list contains two IP addresses, one from IPv4 family and one from IPv6. In single stack clusters a single IP address is expected. When omitted, values from the status.apiServerInternalIPs will be used. Once set, the list cannot be completely removed (but its second entry can).",
										MarkdownDescription: "apiServerInternalIPs are the IP addresses to contact the Kubernetes API server that can be used by components inside the cluster, like kubelets using the infrastructure rather than Kubernetes networking. These are the IPs for a self-hosted load balancer in front of the API servers. In dual stack clusters this list contains two IP addresses, one from IPv4 family and one from IPv6. In single stack clusters a single IP address is expected. When omitted, values from the status.apiServerInternalIPs will be used. Once set, the list cannot be completely removed (but its second entry can).",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ingress_i_ps": schema.ListAttribute{
										Description:         "ingressIPs are the external IPs which route to the default ingress controller. The IPs are suitable targets of a wildcard DNS record used to resolve default route host names. In dual stack clusters this list contains two IP addresses, one from IPv4 family and one from IPv6. In single stack clusters a single IP address is expected. When omitted, values from the status.ingressIPs will be used. Once set, the list cannot be completely removed (but its second entry can).",
										MarkdownDescription: "ingressIPs are the external IPs which route to the default ingress controller. The IPs are suitable targets of a wildcard DNS record used to resolve default route host names. In dual stack clusters this list contains two IP addresses, one from IPv4 family and one from IPv6. In single stack clusters a single IP address is expected. When omitted, values from the status.ingressIPs will be used. Once set, the list cannot be completely removed (but its second entry can).",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"machine_networks": schema.ListAttribute{
										Description:         "machineNetworks are IP networks used to connect all the OpenShift cluster nodes. Each network is provided in the CIDR format and should be IPv4 or IPv6, for example '10.0.0.0/8' or 'fd00::/8'.",
										MarkdownDescription: "machineNetworks are IP networks used to connect all the OpenShift cluster nodes. Each network is provided in the CIDR format and should be IPv4 or IPv6, for example '10.0.0.0/8' or 'fd00::/8'.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"equinix_metal": schema.MapAttribute{
								Description:         "EquinixMetal contains settings specific to the Equinix Metal infrastructure provider.",
								MarkdownDescription: "EquinixMetal contains settings specific to the Equinix Metal infrastructure provider.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"external": schema.SingleNestedAttribute{
								Description:         "ExternalPlatformType represents generic infrastructure provider. Platform-specific components should be supplemented separately.",
								MarkdownDescription: "ExternalPlatformType represents generic infrastructure provider. Platform-specific components should be supplemented separately.",
								Attributes: map[string]schema.Attribute{
									"platform_name": schema.StringAttribute{
										Description:         "PlatformName holds the arbitrary string representing the infrastructure provider name, expected to be set at the installation time. This field is solely for informational and reporting purposes and is not expected to be used for decision-making.",
										MarkdownDescription: "PlatformName holds the arbitrary string representing the infrastructure provider name, expected to be set at the installation time. This field is solely for informational and reporting purposes and is not expected to be used for decision-making.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"gcp": schema.MapAttribute{
								Description:         "GCP contains settings specific to the Google Cloud Platform infrastructure provider.",
								MarkdownDescription: "GCP contains settings specific to the Google Cloud Platform infrastructure provider.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ibmcloud": schema.MapAttribute{
								Description:         "IBMCloud contains settings specific to the IBMCloud infrastructure provider.",
								MarkdownDescription: "IBMCloud contains settings specific to the IBMCloud infrastructure provider.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kubevirt": schema.MapAttribute{
								Description:         "Kubevirt contains settings specific to the kubevirt infrastructure provider.",
								MarkdownDescription: "Kubevirt contains settings specific to the kubevirt infrastructure provider.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"nutanix": schema.SingleNestedAttribute{
								Description:         "Nutanix contains settings specific to the Nutanix infrastructure provider.",
								MarkdownDescription: "Nutanix contains settings specific to the Nutanix infrastructure provider.",
								Attributes: map[string]schema.Attribute{
									"failure_domains": schema.ListNestedAttribute{
										Description:         "failureDomains configures failure domains information for the Nutanix platform. When set, the failure domains defined here may be used to spread Machines across prism element clusters to improve fault tolerance of the cluster.",
										MarkdownDescription: "failureDomains configures failure domains information for the Nutanix platform. When set, the failure domains defined here may be used to spread Machines across prism element clusters to improve fault tolerance of the cluster.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"cluster": schema.SingleNestedAttribute{
													Description:         "cluster is to identify the cluster (the Prism Element under management of the Prism Central), in which the Machine's VM will be created. The cluster identifier (uuid or name) can be obtained from the Prism Central console or using the prism_central API.",
													MarkdownDescription: "cluster is to identify the cluster (the Prism Element under management of the Prism Central), in which the Machine's VM will be created. The cluster identifier (uuid or name) can be obtained from the Prism Central console or using the prism_central API.",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "name is the resource name in the PC. It cannot be empty if the type is Name.",
															MarkdownDescription: "name is the resource name in the PC. It cannot be empty if the type is Name.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"type": schema.StringAttribute{
															Description:         "type is the identifier type to use for this resource.",
															MarkdownDescription: "type is the identifier type to use for this resource.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("UUID", "Name"),
															},
														},

														"uuid": schema.StringAttribute{
															Description:         "uuid is the UUID of the resource in the PC. It cannot be empty if the type is UUID.",
															MarkdownDescription: "uuid is the UUID of the resource in the PC. It cannot be empty if the type is UUID.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},

												"name": schema.StringAttribute{
													Description:         "name defines the unique name of a failure domain. Name is required and must be at most 64 characters in length. It must consist of only lower case alphanumeric characters and hyphens (-). It must start and end with an alphanumeric character. This value is arbitrary and is used to identify the failure domain within the platform.",
													MarkdownDescription: "name defines the unique name of a failure domain. Name is required and must be at most 64 characters in length. It must consist of only lower case alphanumeric characters and hyphens (-). It must start and end with an alphanumeric character. This value is arbitrary and is used to identify the failure domain within the platform.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
														stringvalidator.LengthAtMost(64),
														stringvalidator.RegexMatches(regexp.MustCompile(`[a-z0-9]([-a-z0-9]*[a-z0-9])?`), ""),
													},
												},

												"subnets": schema.ListNestedAttribute{
													Description:         "subnets holds a list of identifiers (one or more) of the cluster's network subnets for the Machine's VM to connect to. The subnet identifiers (uuid or name) can be obtained from the Prism Central console or using the prism_central API.",
													MarkdownDescription: "subnets holds a list of identifiers (one or more) of the cluster's network subnets for the Machine's VM to connect to. The subnet identifiers (uuid or name) can be obtained from the Prism Central console or using the prism_central API.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "name is the resource name in the PC. It cannot be empty if the type is Name.",
																MarkdownDescription: "name is the resource name in the PC. It cannot be empty if the type is Name.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"type": schema.StringAttribute{
																Description:         "type is the identifier type to use for this resource.",
																MarkdownDescription: "type is the identifier type to use for this resource.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("UUID", "Name"),
																},
															},

															"uuid": schema.StringAttribute{
																Description:         "uuid is the UUID of the resource in the PC. It cannot be empty if the type is UUID.",
																MarkdownDescription: "uuid is the UUID of the resource in the PC. It cannot be empty if the type is UUID.",
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"prism_central": schema.SingleNestedAttribute{
										Description:         "prismCentral holds the endpoint address and port to access the Nutanix Prism Central. When a cluster-wide proxy is installed, by default, this endpoint will be accessed via the proxy. Should you wish for communication with this endpoint not to be proxied, please add the endpoint to the proxy spec.noProxy list.",
										MarkdownDescription: "prismCentral holds the endpoint address and port to access the Nutanix Prism Central. When a cluster-wide proxy is installed, by default, this endpoint will be accessed via the proxy. Should you wish for communication with this endpoint not to be proxied, please add the endpoint to the proxy spec.noProxy list.",
										Attributes: map[string]schema.Attribute{
											"address": schema.StringAttribute{
												Description:         "address is the endpoint address (DNS name or IP address) of the Nutanix Prism Central or Element (cluster)",
												MarkdownDescription: "address is the endpoint address (DNS name or IP address) of the Nutanix Prism Central or Element (cluster)",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtMost(256),
												},
											},

											"port": schema.Int64Attribute{
												Description:         "port is the port number to access the Nutanix Prism Central or Element (cluster)",
												MarkdownDescription: "port is the port number to access the Nutanix Prism Central or Element (cluster)",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
													int64validator.AtMost(65535),
												},
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"prism_elements": schema.ListNestedAttribute{
										Description:         "prismElements holds one or more endpoint address and port data to access the Nutanix Prism Elements (clusters) of the Nutanix Prism Central. Currently we only support one Prism Element (cluster) for an OpenShift cluster, where all the Nutanix resources (VMs, subnets, volumes, etc.) used in the OpenShift cluster are located. In the future, we may support Nutanix resources (VMs, etc.) spread over multiple Prism Elements (clusters) of the Prism Central.",
										MarkdownDescription: "prismElements holds one or more endpoint address and port data to access the Nutanix Prism Elements (clusters) of the Nutanix Prism Central. Currently we only support one Prism Element (cluster) for an OpenShift cluster, where all the Nutanix resources (VMs, subnets, volumes, etc.) used in the OpenShift cluster are located. In the future, we may support Nutanix resources (VMs, etc.) spread over multiple Prism Elements (clusters) of the Prism Central.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"endpoint": schema.SingleNestedAttribute{
													Description:         "endpoint holds the endpoint address and port data of the Prism Element (cluster). When a cluster-wide proxy is installed, by default, this endpoint will be accessed via the proxy. Should you wish for communication with this endpoint not to be proxied, please add the endpoint to the proxy spec.noProxy list.",
													MarkdownDescription: "endpoint holds the endpoint address and port data of the Prism Element (cluster). When a cluster-wide proxy is installed, by default, this endpoint will be accessed via the proxy. Should you wish for communication with this endpoint not to be proxied, please add the endpoint to the proxy spec.noProxy list.",
													Attributes: map[string]schema.Attribute{
														"address": schema.StringAttribute{
															Description:         "address is the endpoint address (DNS name or IP address) of the Nutanix Prism Central or Element (cluster)",
															MarkdownDescription: "address is the endpoint address (DNS name or IP address) of the Nutanix Prism Central or Element (cluster)",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtMost(256),
															},
														},

														"port": schema.Int64Attribute{
															Description:         "port is the port number to access the Nutanix Prism Central or Element (cluster)",
															MarkdownDescription: "port is the port number to access the Nutanix Prism Central or Element (cluster)",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(1),
																int64validator.AtMost(65535),
															},
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},

												"name": schema.StringAttribute{
													Description:         "name is the name of the Prism Element (cluster). This value will correspond with the cluster field configured on other resources (eg Machines, PVCs, etc).",
													MarkdownDescription: "name is the name of the Prism Element (cluster). This value will correspond with the cluster field configured on other resources (eg Machines, PVCs, etc).",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtMost(256),
													},
												},
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"openstack": schema.SingleNestedAttribute{
								Description:         "OpenStack contains settings specific to the OpenStack infrastructure provider.",
								MarkdownDescription: "OpenStack contains settings specific to the OpenStack infrastructure provider.",
								Attributes: map[string]schema.Attribute{
									"api_server_internal_i_ps": schema.ListAttribute{
										Description:         "apiServerInternalIPs are the IP addresses to contact the Kubernetes API server that can be used by components inside the cluster, like kubelets using the infrastructure rather than Kubernetes networking. These are the IPs for a self-hosted load balancer in front of the API servers. In dual stack clusters this list contains two IP addresses, one from IPv4 family and one from IPv6. In single stack clusters a single IP address is expected. When omitted, values from the status.apiServerInternalIPs will be used. Once set, the list cannot be completely removed (but its second entry can).",
										MarkdownDescription: "apiServerInternalIPs are the IP addresses to contact the Kubernetes API server that can be used by components inside the cluster, like kubelets using the infrastructure rather than Kubernetes networking. These are the IPs for a self-hosted load balancer in front of the API servers. In dual stack clusters this list contains two IP addresses, one from IPv4 family and one from IPv6. In single stack clusters a single IP address is expected. When omitted, values from the status.apiServerInternalIPs will be used. Once set, the list cannot be completely removed (but its second entry can).",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ingress_i_ps": schema.ListAttribute{
										Description:         "ingressIPs are the external IPs which route to the default ingress controller. The IPs are suitable targets of a wildcard DNS record used to resolve default route host names. In dual stack clusters this list contains two IP addresses, one from IPv4 family and one from IPv6. In single stack clusters a single IP address is expected. When omitted, values from the status.ingressIPs will be used. Once set, the list cannot be completely removed (but its second entry can).",
										MarkdownDescription: "ingressIPs are the external IPs which route to the default ingress controller. The IPs are suitable targets of a wildcard DNS record used to resolve default route host names. In dual stack clusters this list contains two IP addresses, one from IPv4 family and one from IPv6. In single stack clusters a single IP address is expected. When omitted, values from the status.ingressIPs will be used. Once set, the list cannot be completely removed (but its second entry can).",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"machine_networks": schema.ListAttribute{
										Description:         "machineNetworks are IP networks used to connect all the OpenShift cluster nodes. Each network is provided in the CIDR format and should be IPv4 or IPv6, for example '10.0.0.0/8' or 'fd00::/8'.",
										MarkdownDescription: "machineNetworks are IP networks used to connect all the OpenShift cluster nodes. Each network is provided in the CIDR format and should be IPv4 or IPv6, for example '10.0.0.0/8' or 'fd00::/8'.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"ovirt": schema.MapAttribute{
								Description:         "Ovirt contains settings specific to the oVirt infrastructure provider.",
								MarkdownDescription: "Ovirt contains settings specific to the oVirt infrastructure provider.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"powervs": schema.SingleNestedAttribute{
								Description:         "PowerVS contains settings specific to the IBM Power Systems Virtual Servers infrastructure provider.",
								MarkdownDescription: "PowerVS contains settings specific to the IBM Power Systems Virtual Servers infrastructure provider.",
								Attributes: map[string]schema.Attribute{
									"service_endpoints": schema.ListNestedAttribute{
										Description:         "serviceEndpoints is a list of custom endpoints which will override the default service endpoints of a Power VS service.",
										MarkdownDescription: "serviceEndpoints is a list of custom endpoints which will override the default service endpoints of a Power VS service.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "name is the name of the Power VS service. Few of the services are IAM - https://cloud.ibm.com/apidocs/iam-identity-token-api ResourceController - https://cloud.ibm.com/apidocs/resource-controller/resource-controller Power Cloud - https://cloud.ibm.com/apidocs/power-cloud",
													MarkdownDescription: "name is the name of the Power VS service. Few of the services are IAM - https://cloud.ibm.com/apidocs/iam-identity-token-api ResourceController - https://cloud.ibm.com/apidocs/resource-controller/resource-controller Power Cloud - https://cloud.ibm.com/apidocs/power-cloud",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9-]+$`), ""),
													},
												},

												"url": schema.StringAttribute{
													Description:         "url is fully qualified URI with scheme https, that overrides the default generated endpoint for a client. This must be provided and cannot be empty.",
													MarkdownDescription: "url is fully qualified URI with scheme https, that overrides the default generated endpoint for a client. This must be provided and cannot be empty.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`^https://`), ""),
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

							"type": schema.StringAttribute{
								Description:         "type is the underlying infrastructure provider for the cluster. This value controls whether infrastructure automation such as service load balancers, dynamic volume provisioning, machine creation and deletion, and other integrations are enabled. If None, no infrastructure automation is enabled. Allowed values are 'AWS', 'Azure', 'BareMetal', 'GCP', 'Libvirt', 'OpenStack', 'VSphere', 'oVirt', 'KubeVirt', 'EquinixMetal', 'PowerVS', 'AlibabaCloud', 'Nutanix' and 'None'. Individual components may not support all platforms, and must handle unrecognized platforms as None if they do not support that platform.",
								MarkdownDescription: "type is the underlying infrastructure provider for the cluster. This value controls whether infrastructure automation such as service load balancers, dynamic volume provisioning, machine creation and deletion, and other integrations are enabled. If None, no infrastructure automation is enabled. Allowed values are 'AWS', 'Azure', 'BareMetal', 'GCP', 'Libvirt', 'OpenStack', 'VSphere', 'oVirt', 'KubeVirt', 'EquinixMetal', 'PowerVS', 'AlibabaCloud', 'Nutanix' and 'None'. Individual components may not support all platforms, and must handle unrecognized platforms as None if they do not support that platform.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("", "AWS", "Azure", "BareMetal", "GCP", "Libvirt", "OpenStack", "None", "VSphere", "oVirt", "IBMCloud", "KubeVirt", "EquinixMetal", "PowerVS", "AlibabaCloud", "Nutanix", "External"),
								},
							},

							"vsphere": schema.SingleNestedAttribute{
								Description:         "VSphere contains settings specific to the VSphere infrastructure provider.",
								MarkdownDescription: "VSphere contains settings specific to the VSphere infrastructure provider.",
								Attributes: map[string]schema.Attribute{
									"api_server_internal_i_ps": schema.ListAttribute{
										Description:         "apiServerInternalIPs are the IP addresses to contact the Kubernetes API server that can be used by components inside the cluster, like kubelets using the infrastructure rather than Kubernetes networking. These are the IPs for a self-hosted load balancer in front of the API servers. In dual stack clusters this list contains two IP addresses, one from IPv4 family and one from IPv6. In single stack clusters a single IP address is expected. When omitted, values from the status.apiServerInternalIPs will be used. Once set, the list cannot be completely removed (but its second entry can).",
										MarkdownDescription: "apiServerInternalIPs are the IP addresses to contact the Kubernetes API server that can be used by components inside the cluster, like kubelets using the infrastructure rather than Kubernetes networking. These are the IPs for a self-hosted load balancer in front of the API servers. In dual stack clusters this list contains two IP addresses, one from IPv4 family and one from IPv6. In single stack clusters a single IP address is expected. When omitted, values from the status.apiServerInternalIPs will be used. Once set, the list cannot be completely removed (but its second entry can).",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"failure_domains": schema.ListNestedAttribute{
										Description:         "failureDomains contains the definition of region, zone and the vCenter topology. If this is omitted failure domains (regions and zones) will not be used.",
										MarkdownDescription: "failureDomains contains the definition of region, zone and the vCenter topology. If this is omitted failure domains (regions and zones) will not be used.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "name defines the arbitrary but unique name of a failure domain.",
													MarkdownDescription: "name defines the arbitrary but unique name of a failure domain.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
														stringvalidator.LengthAtMost(256),
													},
												},

												"region": schema.StringAttribute{
													Description:         "region defines the name of a region tag that will be attached to a vCenter datacenter. The tag category in vCenter must be named openshift-region.",
													MarkdownDescription: "region defines the name of a region tag that will be attached to a vCenter datacenter. The tag category in vCenter must be named openshift-region.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
														stringvalidator.LengthAtMost(80),
													},
												},

												"server": schema.StringAttribute{
													Description:         "server is the fully-qualified domain name or the IP address of the vCenter server. ---",
													MarkdownDescription: "server is the fully-qualified domain name or the IP address of the vCenter server. ---",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
														stringvalidator.LengthAtMost(255),
													},
												},

												"topology": schema.SingleNestedAttribute{
													Description:         "Topology describes a given failure domain using vSphere constructs",
													MarkdownDescription: "Topology describes a given failure domain using vSphere constructs",
													Attributes: map[string]schema.Attribute{
														"compute_cluster": schema.StringAttribute{
															Description:         "computeCluster the absolute path of the vCenter cluster in which virtual machine will be located. The absolute path is of the form /<datacenter>/host/<cluster>. The maximum length of the path is 2048 characters.",
															MarkdownDescription: "computeCluster the absolute path of the vCenter cluster in which virtual machine will be located. The absolute path is of the form /<datacenter>/host/<cluster>. The maximum length of the path is 2048 characters.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtMost(2048),
																stringvalidator.RegexMatches(regexp.MustCompile(`^/.*?/host/.*?`), ""),
															},
														},

														"datacenter": schema.StringAttribute{
															Description:         "datacenter is the name of vCenter datacenter in which virtual machines will be located. The maximum length of the datacenter name is 80 characters.",
															MarkdownDescription: "datacenter is the name of vCenter datacenter in which virtual machines will be located. The maximum length of the datacenter name is 80 characters.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtMost(80),
															},
														},

														"datastore": schema.StringAttribute{
															Description:         "datastore is the absolute path of the datastore in which the virtual machine is located. The absolute path is of the form /<datacenter>/datastore/<datastore> The maximum length of the path is 2048 characters.",
															MarkdownDescription: "datastore is the absolute path of the datastore in which the virtual machine is located. The absolute path is of the form /<datacenter>/datastore/<datastore> The maximum length of the path is 2048 characters.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtMost(2048),
																stringvalidator.RegexMatches(regexp.MustCompile(`^/.*?/datastore/.*?`), ""),
															},
														},

														"folder": schema.StringAttribute{
															Description:         "folder is the absolute path of the folder where virtual machines are located. The absolute path is of the form /<datacenter>/vm/<folder>. The maximum length of the path is 2048 characters.",
															MarkdownDescription: "folder is the absolute path of the folder where virtual machines are located. The absolute path is of the form /<datacenter>/vm/<folder>. The maximum length of the path is 2048 characters.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtMost(2048),
																stringvalidator.RegexMatches(regexp.MustCompile(`^/.*?/vm/.*?`), ""),
															},
														},

														"networks": schema.ListAttribute{
															Description:         "networks is the list of port group network names within this failure domain. Currently, we only support a single interface per RHCOS virtual machine. The available networks (port groups) can be listed using 'govc ls 'network/*'' The single interface should be the absolute path of the form /<datacenter>/network/<portgroup>.",
															MarkdownDescription: "networks is the list of port group network names within this failure domain. Currently, we only support a single interface per RHCOS virtual machine. The available networks (port groups) can be listed using 'govc ls 'network/*'' The single interface should be the absolute path of the form /<datacenter>/network/<portgroup>.",
															ElementType:         types.StringType,
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"resource_pool": schema.StringAttribute{
															Description:         "resourcePool is the absolute path of the resource pool where virtual machines will be created. The absolute path is of the form /<datacenter>/host/<cluster>/Resources/<resourcepool>. The maximum length of the path is 2048 characters.",
															MarkdownDescription: "resourcePool is the absolute path of the resource pool where virtual machines will be created. The absolute path is of the form /<datacenter>/host/<cluster>/Resources/<resourcepool>. The maximum length of the path is 2048 characters.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtMost(2048),
																stringvalidator.RegexMatches(regexp.MustCompile(`^/.*?/host/.*?/Resources.*`), ""),
															},
														},

														"template": schema.StringAttribute{
															Description:         "template is the full inventory path of the virtual machine or template that will be cloned when creating new machines in this failure domain. The maximum length of the path is 2048 characters.  When omitted, the template will be calculated by the control plane machineset operator based on the region and zone defined in VSpherePlatformFailureDomainSpec. For example, for zone=zonea, region=region1, and infrastructure name=test, the template path would be calculated as /<datacenter>/vm/test-rhcos-region1-zonea.",
															MarkdownDescription: "template is the full inventory path of the virtual machine or template that will be cloned when creating new machines in this failure domain. The maximum length of the path is 2048 characters.  When omitted, the template will be calculated by the control plane machineset operator based on the region and zone defined in VSpherePlatformFailureDomainSpec. For example, for zone=zonea, region=region1, and infrastructure name=test, the template path would be calculated as /<datacenter>/vm/test-rhcos-region1-zonea.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtLeast(1),
																stringvalidator.LengthAtMost(2048),
																stringvalidator.RegexMatches(regexp.MustCompile(`^/.*?/vm/.*?`), ""),
															},
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},

												"zone": schema.StringAttribute{
													Description:         "zone defines the name of a zone tag that will be attached to a vCenter cluster. The tag category in vCenter must be named openshift-zone.",
													MarkdownDescription: "zone defines the name of a zone tag that will be attached to a vCenter cluster. The tag category in vCenter must be named openshift-zone.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
														stringvalidator.LengthAtMost(80),
													},
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"ingress_i_ps": schema.ListAttribute{
										Description:         "ingressIPs are the external IPs which route to the default ingress controller. The IPs are suitable targets of a wildcard DNS record used to resolve default route host names. In dual stack clusters this list contains two IP addresses, one from IPv4 family and one from IPv6. In single stack clusters a single IP address is expected. When omitted, values from the status.ingressIPs will be used. Once set, the list cannot be completely removed (but its second entry can).",
										MarkdownDescription: "ingressIPs are the external IPs which route to the default ingress controller. The IPs are suitable targets of a wildcard DNS record used to resolve default route host names. In dual stack clusters this list contains two IP addresses, one from IPv4 family and one from IPv6. In single stack clusters a single IP address is expected. When omitted, values from the status.ingressIPs will be used. Once set, the list cannot be completely removed (but its second entry can).",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"machine_networks": schema.ListAttribute{
										Description:         "machineNetworks are IP networks used to connect all the OpenShift cluster nodes. Each network is provided in the CIDR format and should be IPv4 or IPv6, for example '10.0.0.0/8' or 'fd00::/8'.",
										MarkdownDescription: "machineNetworks are IP networks used to connect all the OpenShift cluster nodes. Each network is provided in the CIDR format and should be IPv4 or IPv6, for example '10.0.0.0/8' or 'fd00::/8'.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"node_networking": schema.SingleNestedAttribute{
										Description:         "nodeNetworking contains the definition of internal and external network constraints for assigning the node's networking. If this field is omitted, networking defaults to the legacy address selection behavior which is to only support a single address and return the first one found.",
										MarkdownDescription: "nodeNetworking contains the definition of internal and external network constraints for assigning the node's networking. If this field is omitted, networking defaults to the legacy address selection behavior which is to only support a single address and return the first one found.",
										Attributes: map[string]schema.Attribute{
											"external": schema.SingleNestedAttribute{
												Description:         "external represents the network configuration of the node that is externally routable.",
												MarkdownDescription: "external represents the network configuration of the node that is externally routable.",
												Attributes: map[string]schema.Attribute{
													"exclude_network_subnet_cidr": schema.ListAttribute{
														Description:         "excludeNetworkSubnetCidr IP addresses in subnet ranges will be excluded when selecting the IP address from the VirtualMachine's VM for use in the status.addresses fields. ---",
														MarkdownDescription: "excludeNetworkSubnetCidr IP addresses in subnet ranges will be excluded when selecting the IP address from the VirtualMachine's VM for use in the status.addresses fields. ---",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"network": schema.StringAttribute{
														Description:         "network VirtualMachine's VM Network names that will be used to when searching for status.addresses fields. Note that if internal.networkSubnetCIDR and external.networkSubnetCIDR are not set, then the vNIC associated to this network must only have a single IP address assigned to it. The available networks (port groups) can be listed using 'govc ls 'network/*''",
														MarkdownDescription: "network VirtualMachine's VM Network names that will be used to when searching for status.addresses fields. Note that if internal.networkSubnetCIDR and external.networkSubnetCIDR are not set, then the vNIC associated to this network must only have a single IP address assigned to it. The available networks (port groups) can be listed using 'govc ls 'network/*''",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"network_subnet_cidr": schema.ListAttribute{
														Description:         "networkSubnetCidr IP address on VirtualMachine's network interfaces included in the fields' CIDRs that will be used in respective status.addresses fields. ---",
														MarkdownDescription: "networkSubnetCidr IP address on VirtualMachine's network interfaces included in the fields' CIDRs that will be used in respective status.addresses fields. ---",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"internal": schema.SingleNestedAttribute{
												Description:         "internal represents the network configuration of the node that is routable only within the cluster.",
												MarkdownDescription: "internal represents the network configuration of the node that is routable only within the cluster.",
												Attributes: map[string]schema.Attribute{
													"exclude_network_subnet_cidr": schema.ListAttribute{
														Description:         "excludeNetworkSubnetCidr IP addresses in subnet ranges will be excluded when selecting the IP address from the VirtualMachine's VM for use in the status.addresses fields. ---",
														MarkdownDescription: "excludeNetworkSubnetCidr IP addresses in subnet ranges will be excluded when selecting the IP address from the VirtualMachine's VM for use in the status.addresses fields. ---",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"network": schema.StringAttribute{
														Description:         "network VirtualMachine's VM Network names that will be used to when searching for status.addresses fields. Note that if internal.networkSubnetCIDR and external.networkSubnetCIDR are not set, then the vNIC associated to this network must only have a single IP address assigned to it. The available networks (port groups) can be listed using 'govc ls 'network/*''",
														MarkdownDescription: "network VirtualMachine's VM Network names that will be used to when searching for status.addresses fields. Note that if internal.networkSubnetCIDR and external.networkSubnetCIDR are not set, then the vNIC associated to this network must only have a single IP address assigned to it. The available networks (port groups) can be listed using 'govc ls 'network/*''",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"network_subnet_cidr": schema.ListAttribute{
														Description:         "networkSubnetCidr IP address on VirtualMachine's network interfaces included in the fields' CIDRs that will be used in respective status.addresses fields. ---",
														MarkdownDescription: "networkSubnetCidr IP address on VirtualMachine's network interfaces included in the fields' CIDRs that will be used in respective status.addresses fields. ---",
														ElementType:         types.StringType,
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

									"vcenters": schema.ListNestedAttribute{
										Description:         "vcenters holds the connection details for services to communicate with vCenter. Currently, only a single vCenter is supported. ---",
										MarkdownDescription: "vcenters holds the connection details for services to communicate with vCenter. Currently, only a single vCenter is supported. ---",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"datacenters": schema.ListAttribute{
													Description:         "The vCenter Datacenters in which the RHCOS vm guests are located. This field will be used by the Cloud Controller Manager. Each datacenter listed here should be used within a topology.",
													MarkdownDescription: "The vCenter Datacenters in which the RHCOS vm guests are located. This field will be used by the Cloud Controller Manager. Each datacenter listed here should be used within a topology.",
													ElementType:         types.StringType,
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"port": schema.Int64Attribute{
													Description:         "port is the TCP port that will be used to communicate to the vCenter endpoint. When omitted, this means the user has no opinion and it is up to the platform to choose a sensible default, which is subject to change over time.",
													MarkdownDescription: "port is the TCP port that will be used to communicate to the vCenter endpoint. When omitted, this means the user has no opinion and it is up to the platform to choose a sensible default, which is subject to change over time.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(1),
														int64validator.AtMost(32767),
													},
												},

												"server": schema.StringAttribute{
													Description:         "server is the fully-qualified domain name or the IP address of the vCenter server. ---",
													MarkdownDescription: "server is the fully-qualified domain name or the IP address of the vCenter server. ---",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtMost(255),
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

func (r *ConfigOpenshiftIoInfrastructureV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_config_openshift_io_infrastructure_v1_manifest")

	var model ConfigOpenshiftIoInfrastructureV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("config.openshift.io/v1")
	model.Kind = pointer.String("Infrastructure")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
