/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package infrastructure_cluster_x_k8s_io_v1beta2

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
	_ datasource.DataSource = &InfrastructureClusterXK8SIoIbmpowerVsclusterV1Beta2Manifest{}
)

func NewInfrastructureClusterXK8SIoIbmpowerVsclusterV1Beta2Manifest() datasource.DataSource {
	return &InfrastructureClusterXK8SIoIbmpowerVsclusterV1Beta2Manifest{}
}

type InfrastructureClusterXK8SIoIbmpowerVsclusterV1Beta2Manifest struct{}

type InfrastructureClusterXK8SIoIbmpowerVsclusterV1Beta2ManifestData struct {
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
		ControlPlaneEndpoint *struct {
			Host *string `tfsdk:"host" json:"host,omitempty"`
			Port *int64  `tfsdk:"port" json:"port,omitempty"`
		} `tfsdk:"control_plane_endpoint" json:"controlPlaneEndpoint,omitempty"`
		CosInstance *struct {
			BucketName   *string `tfsdk:"bucket_name" json:"bucketName,omitempty"`
			BucketRegion *string `tfsdk:"bucket_region" json:"bucketRegion,omitempty"`
			Name         *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"cos_instance" json:"cosInstance,omitempty"`
		DhcpServer *struct {
			Cidr      *string `tfsdk:"cidr" json:"cidr,omitempty"`
			DnsServer *string `tfsdk:"dns_server" json:"dnsServer,omitempty"`
			Id        *string `tfsdk:"id" json:"id,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Snat      *bool   `tfsdk:"snat" json:"snat,omitempty"`
		} `tfsdk:"dhcp_server" json:"dhcpServer,omitempty"`
		Ignition *struct {
			Version *string `tfsdk:"version" json:"version,omitempty"`
		} `tfsdk:"ignition" json:"ignition,omitempty"`
		LoadBalancers *[]struct {
			AdditionalListeners *[]struct {
				DefaultPoolName *string `tfsdk:"default_pool_name" json:"defaultPoolName,omitempty"`
				Port            *int64  `tfsdk:"port" json:"port,omitempty"`
				Protocol        *string `tfsdk:"protocol" json:"protocol,omitempty"`
			} `tfsdk:"additional_listeners" json:"additionalListeners,omitempty"`
			BackendPools *[]struct {
				Algorithm     *string `tfsdk:"algorithm" json:"algorithm,omitempty"`
				HealthMonitor *struct {
					Delay   *int64  `tfsdk:"delay" json:"delay,omitempty"`
					Port    *int64  `tfsdk:"port" json:"port,omitempty"`
					Retries *int64  `tfsdk:"retries" json:"retries,omitempty"`
					Timeout *int64  `tfsdk:"timeout" json:"timeout,omitempty"`
					Type    *string `tfsdk:"type" json:"type,omitempty"`
					UrlPath *string `tfsdk:"url_path" json:"urlPath,omitempty"`
				} `tfsdk:"health_monitor" json:"healthMonitor,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
			} `tfsdk:"backend_pools" json:"backendPools,omitempty"`
			Id             *string `tfsdk:"id" json:"id,omitempty"`
			Name           *string `tfsdk:"name" json:"name,omitempty"`
			Public         *bool   `tfsdk:"public" json:"public,omitempty"`
			SecurityGroups *[]struct {
				Id   *string `tfsdk:"id" json:"id,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"security_groups" json:"securityGroups,omitempty"`
			Subnets *[]struct {
				Id   *string `tfsdk:"id" json:"id,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"subnets" json:"subnets,omitempty"`
		} `tfsdk:"load_balancers" json:"loadBalancers,omitempty"`
		Network *struct {
			Id    *string `tfsdk:"id" json:"id,omitempty"`
			Name  *string `tfsdk:"name" json:"name,omitempty"`
			Regex *string `tfsdk:"regex" json:"regex,omitempty"`
		} `tfsdk:"network" json:"network,omitempty"`
		ResourceGroup *struct {
			Id    *string `tfsdk:"id" json:"id,omitempty"`
			Name  *string `tfsdk:"name" json:"name,omitempty"`
			Regex *string `tfsdk:"regex" json:"regex,omitempty"`
		} `tfsdk:"resource_group" json:"resourceGroup,omitempty"`
		ServiceInstance *struct {
			Id    *string `tfsdk:"id" json:"id,omitempty"`
			Name  *string `tfsdk:"name" json:"name,omitempty"`
			Regex *string `tfsdk:"regex" json:"regex,omitempty"`
		} `tfsdk:"service_instance" json:"serviceInstance,omitempty"`
		ServiceInstanceID *string `tfsdk:"service_instance_id" json:"serviceInstanceID,omitempty"`
		TransitGateway    *struct {
			GlobalRouting *bool   `tfsdk:"global_routing" json:"globalRouting,omitempty"`
			Id            *string `tfsdk:"id" json:"id,omitempty"`
			Name          *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"transit_gateway" json:"transitGateway,omitempty"`
		Vpc *struct {
			Id     *string `tfsdk:"id" json:"id,omitempty"`
			Name   *string `tfsdk:"name" json:"name,omitempty"`
			Region *string `tfsdk:"region" json:"region,omitempty"`
		} `tfsdk:"vpc" json:"vpc,omitempty"`
		VpcSecurityGroups *[]struct {
			Id    *string `tfsdk:"id" json:"id,omitempty"`
			Name  *string `tfsdk:"name" json:"name,omitempty"`
			Rules *[]struct {
				Action      *string `tfsdk:"action" json:"action,omitempty"`
				Destination *struct {
					IcmpCode  *int64 `tfsdk:"icmp_code" json:"icmpCode,omitempty"`
					IcmpType  *int64 `tfsdk:"icmp_type" json:"icmpType,omitempty"`
					PortRange *struct {
						MaximumPort *int64 `tfsdk:"maximum_port" json:"maximumPort,omitempty"`
						MinimumPort *int64 `tfsdk:"minimum_port" json:"minimumPort,omitempty"`
					} `tfsdk:"port_range" json:"portRange,omitempty"`
					Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
					Remotes  *[]struct {
						Address           *string `tfsdk:"address" json:"address,omitempty"`
						CidrSubnetName    *string `tfsdk:"cidr_subnet_name" json:"cidrSubnetName,omitempty"`
						RemoteType        *string `tfsdk:"remote_type" json:"remoteType,omitempty"`
						SecurityGroupName *string `tfsdk:"security_group_name" json:"securityGroupName,omitempty"`
					} `tfsdk:"remotes" json:"remotes,omitempty"`
				} `tfsdk:"destination" json:"destination,omitempty"`
				Direction       *string `tfsdk:"direction" json:"direction,omitempty"`
				SecurityGroupID *string `tfsdk:"security_group_id" json:"securityGroupID,omitempty"`
				Source          *struct {
					IcmpCode  *int64 `tfsdk:"icmp_code" json:"icmpCode,omitempty"`
					IcmpType  *int64 `tfsdk:"icmp_type" json:"icmpType,omitempty"`
					PortRange *struct {
						MaximumPort *int64 `tfsdk:"maximum_port" json:"maximumPort,omitempty"`
						MinimumPort *int64 `tfsdk:"minimum_port" json:"minimumPort,omitempty"`
					} `tfsdk:"port_range" json:"portRange,omitempty"`
					Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
					Remotes  *[]struct {
						Address           *string `tfsdk:"address" json:"address,omitempty"`
						CidrSubnetName    *string `tfsdk:"cidr_subnet_name" json:"cidrSubnetName,omitempty"`
						RemoteType        *string `tfsdk:"remote_type" json:"remoteType,omitempty"`
						SecurityGroupName *string `tfsdk:"security_group_name" json:"securityGroupName,omitempty"`
					} `tfsdk:"remotes" json:"remotes,omitempty"`
				} `tfsdk:"source" json:"source,omitempty"`
			} `tfsdk:"rules" json:"rules,omitempty"`
			Tags *[]string `tfsdk:"tags" json:"tags,omitempty"`
		} `tfsdk:"vpc_security_groups" json:"vpcSecurityGroups,omitempty"`
		VpcSubnets *[]struct {
			Cidr *string `tfsdk:"cidr" json:"cidr,omitempty"`
			Id   *string `tfsdk:"id" json:"id,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
			Zone *string `tfsdk:"zone" json:"zone,omitempty"`
		} `tfsdk:"vpc_subnets" json:"vpcSubnets,omitempty"`
		Zone *string `tfsdk:"zone" json:"zone,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *InfrastructureClusterXK8SIoIbmpowerVsclusterV1Beta2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_infrastructure_cluster_x_k8s_io_ibm_power_vs_cluster_v1beta2_manifest"
}

func (r *InfrastructureClusterXK8SIoIbmpowerVsclusterV1Beta2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "IBMPowerVSCluster is the Schema for the ibmpowervsclusters API.",
		MarkdownDescription: "IBMPowerVSCluster is the Schema for the ibmpowervsclusters API.",
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
				Description:         "IBMPowerVSClusterSpec defines the desired state of IBMPowerVSCluster.",
				MarkdownDescription: "IBMPowerVSClusterSpec defines the desired state of IBMPowerVSCluster.",
				Attributes: map[string]schema.Attribute{
					"control_plane_endpoint": schema.SingleNestedAttribute{
						Description:         "ControlPlaneEndpoint represents the endpoint used to communicate with the control plane.",
						MarkdownDescription: "ControlPlaneEndpoint represents the endpoint used to communicate with the control plane.",
						Attributes: map[string]schema.Attribute{
							"host": schema.StringAttribute{
								Description:         "The hostname on which the API server is serving.",
								MarkdownDescription: "The hostname on which the API server is serving.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "The port on which the API server is serving.",
								MarkdownDescription: "The port on which the API server is serving.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"cos_instance": schema.SingleNestedAttribute{
						Description:         "cosInstance contains options to configure a supporting IBM Cloud COS bucket for thiscluster - currently used for nodes requiring Ignition(https://coreos.github.io/ignition/) for bootstrapping (requiresBootstrapFormatIgnition feature flag to be enabled).when powervs.cluster.x-k8s.io/create-infra=true annotation is set on IBMPowerVSCluster resource and Ignition is set, then1. CosInstance.Name should be set not setting will result in webhook error.2. CosInstance.BucketName should be set not setting will result in webhook error.3. CosInstance.BucketRegion should be set not setting will result in webhook error.",
						MarkdownDescription: "cosInstance contains options to configure a supporting IBM Cloud COS bucket for thiscluster - currently used for nodes requiring Ignition(https://coreos.github.io/ignition/) for bootstrapping (requiresBootstrapFormatIgnition feature flag to be enabled).when powervs.cluster.x-k8s.io/create-infra=true annotation is set on IBMPowerVSCluster resource and Ignition is set, then1. CosInstance.Name should be set not setting will result in webhook error.2. CosInstance.BucketName should be set not setting will result in webhook error.3. CosInstance.BucketRegion should be set not setting will result in webhook error.",
						Attributes: map[string]schema.Attribute{
							"bucket_name": schema.StringAttribute{
								Description:         "bucketName is IBM cloud COS bucket name",
								MarkdownDescription: "bucketName is IBM cloud COS bucket name",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"bucket_region": schema.StringAttribute{
								Description:         "bucketRegion is IBM cloud COS bucket region",
								MarkdownDescription: "bucketRegion is IBM cloud COS bucket region",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "name defines name of IBM cloud COS instance to be created.when IBMPowerVSCluster.Ignition is set",
								MarkdownDescription: "name defines name of IBM cloud COS instance to be created.when IBMPowerVSCluster.Ignition is set",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(3),
									stringvalidator.LengthAtMost(63),
									stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9][a-z0-9.-]{1,61}[a-z0-9]$`), ""),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"dhcp_server": schema.SingleNestedAttribute{
						Description:         "dhcpServer is contains the configuration to be used while creating a new DHCP server in PowerVS workspace.when the field is omitted, CLUSTER_NAME will be used as DHCPServer.Name and DHCP server will be created.it will automatically create network with name DHCPSERVER<DHCPServer.Name>_Private in PowerVS workspace.",
						MarkdownDescription: "dhcpServer is contains the configuration to be used while creating a new DHCP server in PowerVS workspace.when the field is omitted, CLUSTER_NAME will be used as DHCPServer.Name and DHCP server will be created.it will automatically create network with name DHCPSERVER<DHCPServer.Name>_Private in PowerVS workspace.",
						Attributes: map[string]schema.Attribute{
							"cidr": schema.StringAttribute{
								Description:         "Optional cidr for DHCP private network",
								MarkdownDescription: "Optional cidr for DHCP private network",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"dns_server": schema.StringAttribute{
								Description:         "Optional DNS Server for DHCP service",
								MarkdownDescription: "Optional DNS Server for DHCP service",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"id": schema.StringAttribute{
								Description:         "Optional id of the existing DHCPServer",
								MarkdownDescription: "Optional id of the existing DHCPServer",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Optional name of DHCP Service. Only alphanumeric characters and dashes are allowed.",
								MarkdownDescription: "Optional name of DHCP Service. Only alphanumeric characters and dashes are allowed.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"snat": schema.BoolAttribute{
								Description:         "Optional indicates if SNAT will be enabled for DHCP service",
								MarkdownDescription: "Optional indicates if SNAT will be enabled for DHCP service",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"ignition": schema.SingleNestedAttribute{
						Description:         "Ignition defined options related to the bootstrapping systems where Ignition is used.",
						MarkdownDescription: "Ignition defined options related to the bootstrapping systems where Ignition is used.",
						Attributes: map[string]schema.Attribute{
							"version": schema.StringAttribute{
								Description:         "Version defines which version of Ignition will be used to generate bootstrap data.",
								MarkdownDescription: "Version defines which version of Ignition will be used to generate bootstrap data.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("2.3", "2.4", "3.0", "3.1", "3.2", "3.3", "3.4"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"load_balancers": schema.ListNestedAttribute{
						Description:         "loadBalancers is optional configuration for configuring loadbalancers to control plane or data plane nodes.when omitted system will create a default public loadbalancer with name CLUSTER_NAME-loadbalancer.when specified a vpc loadbalancer will be created and controlPlaneEndpoint will be set with associated hostname of loadbalancer.ControlPlaneEndpoint will be set with associated hostname of public loadbalancer.when LoadBalancers[].ID is set, its expected that there exist a loadbalancer with ID or else system will give error.when LoadBalancers[].Name is set, system will first check for loadbalancer with Name, if not exist system will create new loadbalancer.For each loadbalancer a default backed pool and front listener will be configured with port 6443.",
						MarkdownDescription: "loadBalancers is optional configuration for configuring loadbalancers to control plane or data plane nodes.when omitted system will create a default public loadbalancer with name CLUSTER_NAME-loadbalancer.when specified a vpc loadbalancer will be created and controlPlaneEndpoint will be set with associated hostname of loadbalancer.ControlPlaneEndpoint will be set with associated hostname of public loadbalancer.when LoadBalancers[].ID is set, its expected that there exist a loadbalancer with ID or else system will give error.when LoadBalancers[].Name is set, system will first check for loadbalancer with Name, if not exist system will create new loadbalancer.For each loadbalancer a default backed pool and front listener will be configured with port 6443.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"additional_listeners": schema.ListNestedAttribute{
									Description:         "AdditionalListeners sets the additional listeners for the control plane load balancer.",
									MarkdownDescription: "AdditionalListeners sets the additional listeners for the control plane load balancer.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"default_pool_name": schema.StringAttribute{
												Description:         "defaultPoolName defines the name of a VPC Load Balancer Backend Pool to use for the VPC Load Balancer Listener.",
												MarkdownDescription: "defaultPoolName defines the name of a VPC Load Balancer Backend Pool to use for the VPC Load Balancer Listener.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
													stringvalidator.LengthAtMost(63),
													stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`), ""),
												},
											},

											"port": schema.Int64Attribute{
												Description:         "Port sets the port for the additional listener.",
												MarkdownDescription: "Port sets the port for the additional listener.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
													int64validator.AtMost(65535),
												},
											},

											"protocol": schema.StringAttribute{
												Description:         "protocol defines the protocol to use for the VPC Load Balancer Listener.Will default to TCP protocol if not specified.",
												MarkdownDescription: "protocol defines the protocol to use for the VPC Load Balancer Listener.Will default to TCP protocol if not specified.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("http", "https", "tcp", "udp"),
												},
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"backend_pools": schema.ListNestedAttribute{
									Description:         "backendPools defines the load balancer's backend pools.",
									MarkdownDescription: "backendPools defines the load balancer's backend pools.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"algorithm": schema.StringAttribute{
												Description:         "algorithm defines the load balancing algorithm to use.",
												MarkdownDescription: "algorithm defines the load balancing algorithm to use.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("least_connections", "round_robin", "weighted_round_robin"),
												},
											},

											"health_monitor": schema.SingleNestedAttribute{
												Description:         "healthMonitor defines the backend pool's health monitor.",
												MarkdownDescription: "healthMonitor defines the backend pool's health monitor.",
												Attributes: map[string]schema.Attribute{
													"delay": schema.Int64Attribute{
														Description:         "delay defines the seconds to wait between health checks.",
														MarkdownDescription: "delay defines the seconds to wait between health checks.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(2),
															int64validator.AtMost(60),
														},
													},

													"port": schema.Int64Attribute{
														Description:         "port defines the port to perform health monitoring on.",
														MarkdownDescription: "port defines the port to perform health monitoring on.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(1),
															int64validator.AtMost(65535),
														},
													},

													"retries": schema.Int64Attribute{
														Description:         "retries defines the max retries for health check.",
														MarkdownDescription: "retries defines the max retries for health check.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(1),
															int64validator.AtMost(10),
														},
													},

													"timeout": schema.Int64Attribute{
														Description:         "timeout defines the seconds to wait for a health check response.",
														MarkdownDescription: "timeout defines the seconds to wait for a health check response.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(1),
															int64validator.AtMost(59),
														},
													},

													"type": schema.StringAttribute{
														Description:         "type defines the protocol used for health checks.",
														MarkdownDescription: "type defines the protocol used for health checks.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("http", "https", "tcp"),
														},
													},

													"url_path": schema.StringAttribute{
														Description:         "urlPath defines the URL to use for health monitoring.",
														MarkdownDescription: "urlPath defines the URL to use for health monitoring.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.RegexMatches(regexp.MustCompile(`^\/(([a-zA-Z0-9-._~!$&'()*+,;=:@]|%[a-fA-F0-9]{2})+(\/([a-zA-Z0-9-._~!$&'()*+,;=:@]|%[a-fA-F0-9]{2})*)*)?(\\?([a-zA-Z0-9-._~!$&'()*+,;=:@\/?]|%[a-fA-F0-9]{2})*)?$`), ""),
														},
													},
												},
												Required: true,
												Optional: false,
												Computed: false,
											},

											"name": schema.StringAttribute{
												Description:         "name defines the name of the Backend Pool.",
												MarkdownDescription: "name defines the name of the Backend Pool.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
													stringvalidator.LengthAtMost(63),
													stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`), ""),
												},
											},

											"protocol": schema.StringAttribute{
												Description:         "protocol defines the protocol to use for the Backend Pool.",
												MarkdownDescription: "protocol defines the protocol to use for the Backend Pool.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("http", "https", "tcp", "udp"),
												},
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"id": schema.StringAttribute{
									Description:         "id of the loadbalancer",
									MarkdownDescription: "id of the loadbalancer",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(64),
										stringvalidator.RegexMatches(regexp.MustCompile(`^[-0-9a-z_]+$`), ""),
									},
								},

								"name": schema.StringAttribute{
									Description:         "Name sets the name of the VPC load balancer.",
									MarkdownDescription: "Name sets the name of the VPC load balancer.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(63),
										stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`), ""),
									},
								},

								"public": schema.BoolAttribute{
									Description:         "public indicates that load balancer is public or private",
									MarkdownDescription: "public indicates that load balancer is public or private",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"security_groups": schema.ListNestedAttribute{
									Description:         "securityGroups defines the Security Groups to attach to the load balancer.Security Groups defined here are expected to already exist when the load balancer is reconciled (these do not get created when reconciling the load balancer).",
									MarkdownDescription: "securityGroups defines the Security Groups to attach to the load balancer.Security Groups defined here are expected to already exist when the load balancer is reconciled (these do not get created when reconciling the load balancer).",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description:         "id of the resource.",
												MarkdownDescription: "id of the resource.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
												},
											},

											"name": schema.StringAttribute{
												Description:         "name of the resource.",
												MarkdownDescription: "name of the resource.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
												},
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"subnets": schema.ListNestedAttribute{
									Description:         "subnets defines the VPC Subnets to attach to the load balancer.Subnets defiens here are expected to already exist when the load balancer is reconciled (these do not get created when reconciling the load balancer).",
									MarkdownDescription: "subnets defines the VPC Subnets to attach to the load balancer.Subnets defiens here are expected to already exist when the load balancer is reconciled (these do not get created when reconciling the load balancer).",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description:         "id of the resource.",
												MarkdownDescription: "id of the resource.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
												},
											},

											"name": schema.StringAttribute{
												Description:         "name of the resource.",
												MarkdownDescription: "name of the resource.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
												},
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"network": schema.SingleNestedAttribute{
						Description:         "Network is the reference to the Network to use for this cluster.when the field is omitted, A DHCP service will be created in the Power VS workspace and its private network will be used.the DHCP service created network will have the following name format1. in the case of DHCPServer.Name is not set the name will be DHCPSERVER<CLUSTER_NAME>_Private.2. if DHCPServer.Name is set the name will be DHCPSERVER<DHCPServer.Name>_Private.when Network.ID is set, its expected that there exist a network in PowerVS workspace with id or else system will give error.when Network.Name is set, system will first check for network with Name in PowerVS workspace, if not exist network will be created by DHCP service.Network.RegEx is not yet supported and system will ignore the value.",
						MarkdownDescription: "Network is the reference to the Network to use for this cluster.when the field is omitted, A DHCP service will be created in the Power VS workspace and its private network will be used.the DHCP service created network will have the following name format1. in the case of DHCPServer.Name is not set the name will be DHCPSERVER<CLUSTER_NAME>_Private.2. if DHCPServer.Name is set the name will be DHCPSERVER<DHCPServer.Name>_Private.when Network.ID is set, its expected that there exist a network in PowerVS workspace with id or else system will give error.when Network.Name is set, system will first check for network with Name in PowerVS workspace, if not exist network will be created by DHCP service.Network.RegEx is not yet supported and system will ignore the value.",
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Description:         "ID of resource",
								MarkdownDescription: "ID of resource",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
							},

							"name": schema.StringAttribute{
								Description:         "Name of resource",
								MarkdownDescription: "Name of resource",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
							},

							"regex": schema.StringAttribute{
								Description:         "Regular expression to match resource,In case of multiple resources matches the provided regular expression the first matched resource will be selected",
								MarkdownDescription: "Regular expression to match resource,In case of multiple resources matches the provided regular expression the first matched resource will be selected",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"resource_group": schema.SingleNestedAttribute{
						Description:         "resourceGroup name under which the resources will be created.when powervs.cluster.x-k8s.io/create-infra=true annotation is set on IBMPowerVSCluster resource,1. it is expected to set the ResourceGroup.Name, not setting will result in webhook error.ResourceGroup.ID and ResourceGroup.Regex is not yet supported and system will ignore the value.",
						MarkdownDescription: "resourceGroup name under which the resources will be created.when powervs.cluster.x-k8s.io/create-infra=true annotation is set on IBMPowerVSCluster resource,1. it is expected to set the ResourceGroup.Name, not setting will result in webhook error.ResourceGroup.ID and ResourceGroup.Regex is not yet supported and system will ignore the value.",
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Description:         "ID of resource",
								MarkdownDescription: "ID of resource",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
							},

							"name": schema.StringAttribute{
								Description:         "Name of resource",
								MarkdownDescription: "Name of resource",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
							},

							"regex": schema.StringAttribute{
								Description:         "Regular expression to match resource,In case of multiple resources matches the provided regular expression the first matched resource will be selected",
								MarkdownDescription: "Regular expression to match resource,In case of multiple resources matches the provided regular expression the first matched resource will be selected",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"service_instance": schema.SingleNestedAttribute{
						Description:         "serviceInstance is the reference to the Power VS server workspace on which the server instance(VM) will be created.Power VS server workspace is a container for all Power VS instances at a specific geographic region.serviceInstance can be created via IBM Cloud catalog or CLI.supported serviceInstance identifier in PowerVSResource are Name and ID and that can be obtained from IBM Cloud UI or IBM Cloud cli.More detail about Power VS service instance.https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-creating-power-virtual-serverwhen omitted system will dynamically create the service instance with name CLUSTER_NAME-serviceInstance.when ServiceInstance.ID is set, its expected that there exist a service instance in PowerVS workspace with id or else system will give error.when ServiceInstance.Name is set, system will first check for service instance with Name in PowerVS workspace, if not exist system will create new instance.if there are more than one service instance exist with the ServiceInstance.Name in given Zone, installation fails with an error. Use ServiceInstance.ID in those situations to use the specific service instance.ServiceInstance.Regex is not yet supported not yet supported and system will ignore the value.",
						MarkdownDescription: "serviceInstance is the reference to the Power VS server workspace on which the server instance(VM) will be created.Power VS server workspace is a container for all Power VS instances at a specific geographic region.serviceInstance can be created via IBM Cloud catalog or CLI.supported serviceInstance identifier in PowerVSResource are Name and ID and that can be obtained from IBM Cloud UI or IBM Cloud cli.More detail about Power VS service instance.https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-creating-power-virtual-serverwhen omitted system will dynamically create the service instance with name CLUSTER_NAME-serviceInstance.when ServiceInstance.ID is set, its expected that there exist a service instance in PowerVS workspace with id or else system will give error.when ServiceInstance.Name is set, system will first check for service instance with Name in PowerVS workspace, if not exist system will create new instance.if there are more than one service instance exist with the ServiceInstance.Name in given Zone, installation fails with an error. Use ServiceInstance.ID in those situations to use the specific service instance.ServiceInstance.Regex is not yet supported not yet supported and system will ignore the value.",
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Description:         "ID of resource",
								MarkdownDescription: "ID of resource",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
							},

							"name": schema.StringAttribute{
								Description:         "Name of resource",
								MarkdownDescription: "Name of resource",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
							},

							"regex": schema.StringAttribute{
								Description:         "Regular expression to match resource,In case of multiple resources matches the provided regular expression the first matched resource will be selected",
								MarkdownDescription: "Regular expression to match resource,In case of multiple resources matches the provided regular expression the first matched resource will be selected",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"service_instance_id": schema.StringAttribute{
						Description:         "ServiceInstanceID is the id of the power cloud instance where the vsi instance will get deployed.Deprecated: use ServiceInstance instead",
						MarkdownDescription: "ServiceInstanceID is the id of the power cloud instance where the vsi instance will get deployed.Deprecated: use ServiceInstance instead",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"transit_gateway": schema.SingleNestedAttribute{
						Description:         "transitGateway contains information about IBM Cloud TransitGatewayIBM Cloud TransitGateway helps in establishing network connectivity between IBM Cloud Power VS and VPC infrastructuremore information about TransitGateway can be found here https://www.ibm.com/products/transit-gateway.when TransitGateway.ID is set, its expected that there exist a TransitGateway with ID or else system will give error.when TransitGateway.Name is set, system will first check for TransitGateway with Name, if not exist system will create new TransitGateway.",
						MarkdownDescription: "transitGateway contains information about IBM Cloud TransitGatewayIBM Cloud TransitGateway helps in establishing network connectivity between IBM Cloud Power VS and VPC infrastructuremore information about TransitGateway can be found here https://www.ibm.com/products/transit-gateway.when TransitGateway.ID is set, its expected that there exist a TransitGateway with ID or else system will give error.when TransitGateway.Name is set, system will first check for TransitGateway with Name, if not exist system will create new TransitGateway.",
						Attributes: map[string]schema.Attribute{
							"global_routing": schema.BoolAttribute{
								Description:         "globalRouting indicates whether to set global routing true or not while creating the transit gateway.set this field to true only when PowerVS and VPC are from different regions, if they are same it's suggested to use local routing by setting the field to false.when the field is omitted,  based on PowerVS region (region associated with IBMPowerVSCluster.Spec.Zone) and VPC region(IBMPowerVSCluster.Spec.VPC.Region) system will decide whether to enable globalRouting or not.",
								MarkdownDescription: "globalRouting indicates whether to set global routing true or not while creating the transit gateway.set this field to true only when PowerVS and VPC are from different regions, if they are same it's suggested to use local routing by setting the field to false.when the field is omitted,  based on PowerVS region (region associated with IBMPowerVSCluster.Spec.Zone) and VPC region(IBMPowerVSCluster.Spec.VPC.Region) system will decide whether to enable globalRouting or not.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"id": schema.StringAttribute{
								Description:         "id of resource.",
								MarkdownDescription: "id of resource.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "name of resource.",
								MarkdownDescription: "name of resource.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(63),
									stringvalidator.RegexMatches(regexp.MustCompile(`^([a-zA-Z]|[a-zA-Z][-_a-zA-Z0-9]*[a-zA-Z0-9])$`), ""),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"vpc": schema.SingleNestedAttribute{
						Description:         "vpc contains information about IBM Cloud VPC resources.when omitted system will dynamically create the VPC with name CLUSTER_NAME-vpc.when VPC.ID is set, its expected that there exist a VPC with ID or else system will give error.when VPC.Name is set, system will first check for VPC with Name, if not exist system will create new VPC.when powervs.cluster.x-k8s.io/create-infra=true annotation is set on IBMPowerVSCluster resource,1. it is expected to set the VPC.Region, not setting will result in webhook error.",
						MarkdownDescription: "vpc contains information about IBM Cloud VPC resources.when omitted system will dynamically create the VPC with name CLUSTER_NAME-vpc.when VPC.ID is set, its expected that there exist a VPC with ID or else system will give error.when VPC.Name is set, system will first check for VPC with Name, if not exist system will create new VPC.when powervs.cluster.x-k8s.io/create-infra=true annotation is set on IBMPowerVSCluster resource,1. it is expected to set the VPC.Region, not setting will result in webhook error.",
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Description:         "id of resource.",
								MarkdownDescription: "id of resource.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(64),
									stringvalidator.RegexMatches(regexp.MustCompile(`^[-0-9a-z_]+$`), ""),
								},
							},

							"name": schema.StringAttribute{
								Description:         "name of resource.",
								MarkdownDescription: "name of resource.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(63),
									stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`), ""),
								},
							},

							"region": schema.StringAttribute{
								Description:         "region of IBM Cloud VPC.when powervs.cluster.x-k8s.io/create-infra=true annotation is set on IBMPowerVSCluster resource,it is expected to set the region, not setting will result in webhook error.",
								MarkdownDescription: "region of IBM Cloud VPC.when powervs.cluster.x-k8s.io/create-infra=true annotation is set on IBMPowerVSCluster resource,it is expected to set the region, not setting will result in webhook error.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"vpc_security_groups": schema.ListNestedAttribute{
						Description:         "VPCSecurityGroups to attach it to the VPC resource",
						MarkdownDescription: "VPCSecurityGroups to attach it to the VPC resource",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description:         "id of the Security Group.",
									MarkdownDescription: "id of the Security Group.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "name of the Security Group.",
									MarkdownDescription: "name of the Security Group.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"rules": schema.ListNestedAttribute{
									Description:         "rules are the Security Group Rules for the Security Group.",
									MarkdownDescription: "rules are the Security Group Rules for the Security Group.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"action": schema.StringAttribute{
												Description:         "action defines whether to allow or deny traffic defined by the Security Group Rule.",
												MarkdownDescription: "action defines whether to allow or deny traffic defined by the Security Group Rule.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("allow", "deny"),
												},
											},

											"destination": schema.SingleNestedAttribute{
												Description:         "destination is a VPCSecurityGroupRulePrototype which defines the destination of outbound traffic for the Security Group Rule.Only used when direction is VPCSecurityGroupRuleDirectionOutbound.",
												MarkdownDescription: "destination is a VPCSecurityGroupRulePrototype which defines the destination of outbound traffic for the Security Group Rule.Only used when direction is VPCSecurityGroupRuleDirectionOutbound.",
												Attributes: map[string]schema.Attribute{
													"icmp_code": schema.Int64Attribute{
														Description:         "icmpCode is the ICMP code for the Rule.Only used when Protocol is VPCSecurityGroupRuleProtocolIcmp.",
														MarkdownDescription: "icmpCode is the ICMP code for the Rule.Only used when Protocol is VPCSecurityGroupRuleProtocolIcmp.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"icmp_type": schema.Int64Attribute{
														Description:         "icmpType is the ICMP type for the Rule.Only used when Protocol is VPCSecurityGroupRuleProtocolIcmp.",
														MarkdownDescription: "icmpType is the ICMP type for the Rule.Only used when Protocol is VPCSecurityGroupRuleProtocolIcmp.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"port_range": schema.SingleNestedAttribute{
														Description:         "portRange is a range of ports allowed for the Rule's remote.",
														MarkdownDescription: "portRange is a range of ports allowed for the Rule's remote.",
														Attributes: map[string]schema.Attribute{
															"maximum_port": schema.Int64Attribute{
																Description:         "maximumPort is the inclusive upper range of ports.",
																MarkdownDescription: "maximumPort is the inclusive upper range of ports.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(1),
																	int64validator.AtMost(65535),
																},
															},

															"minimum_port": schema.Int64Attribute{
																Description:         "minimumPort is the inclusive lower range of ports.",
																MarkdownDescription: "minimumPort is the inclusive lower range of ports.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(1),
																	int64validator.AtMost(65535),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"protocol": schema.StringAttribute{
														Description:         "protocol defines the traffic protocol used for the Security Group Rule.",
														MarkdownDescription: "protocol defines the traffic protocol used for the Security Group Rule.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("all", "icmp", "tcp", "udp"),
														},
													},

													"remotes": schema.ListNestedAttribute{
														Description:         "remotes is a set of VPCSecurityGroupRuleRemote's that define the traffic allowed by the Rule's remote.Specifying multiple VPCSecurityGroupRuleRemote's creates a unique Security Group Rule with the shared Protocol, PortRange, etc.This allows for easier management of Security Group Rule's for sets of CIDR's, IP's, etc.",
														MarkdownDescription: "remotes is a set of VPCSecurityGroupRuleRemote's that define the traffic allowed by the Rule's remote.Specifying multiple VPCSecurityGroupRuleRemote's creates a unique Security Group Rule with the shared Protocol, PortRange, etc.This allows for easier management of Security Group Rule's for sets of CIDR's, IP's, etc.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"address": schema.StringAttribute{
																	Description:         " address is the address to use for the remote's destination/source.Only used when remoteType is VPCSecurityGroupRuleRemoteTypeAddress.",
																	MarkdownDescription: " address is the address to use for the remote's destination/source.Only used when remoteType is VPCSecurityGroupRuleRemoteTypeAddress.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"cidr_subnet_name": schema.StringAttribute{
																	Description:         "cidrSubnetName is the name of the VPC Subnet to retrieve the CIDR from, to use for the remote's destination/source.Only used when remoteType is VPCSecurityGroupRuleRemoteTypeCIDR.",
																	MarkdownDescription: "cidrSubnetName is the name of the VPC Subnet to retrieve the CIDR from, to use for the remote's destination/source.Only used when remoteType is VPCSecurityGroupRuleRemoteTypeCIDR.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"remote_type": schema.StringAttribute{
																	Description:         "remoteType defines the type of filter to define for the remote's destination/source.",
																	MarkdownDescription: "remoteType defines the type of filter to define for the remote's destination/source.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("any", "cidr", "address", "sg"),
																	},
																},

																"security_group_name": schema.StringAttribute{
																	Description:         "securityGroupName is the name of the VPC Security Group to use for the remote's destination/source.Only used when remoteType is VPCSecurityGroupRuleRemoteTypeSG",
																	MarkdownDescription: "securityGroupName is the name of the VPC Security Group to use for the remote's destination/source.Only used when remoteType is VPCSecurityGroupRuleRemoteTypeSG",
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
												Required: false,
												Optional: true,
												Computed: false,
											},

											"direction": schema.StringAttribute{
												Description:         "direction defines whether the traffic is inbound or outbound for the Security Group Rule.",
												MarkdownDescription: "direction defines whether the traffic is inbound or outbound for the Security Group Rule.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("inbound", "outbound"),
												},
											},

											"security_group_id": schema.StringAttribute{
												Description:         "securityGroupID is the ID of the Security Group for the Security Group Rule.",
												MarkdownDescription: "securityGroupID is the ID of the Security Group for the Security Group Rule.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"source": schema.SingleNestedAttribute{
												Description:         "source is a VPCSecurityGroupRulePrototype which defines the source of inbound traffic for the Security Group Rule.Only used when direction is VPCSecurityGroupRuleDirectionInbound.",
												MarkdownDescription: "source is a VPCSecurityGroupRulePrototype which defines the source of inbound traffic for the Security Group Rule.Only used when direction is VPCSecurityGroupRuleDirectionInbound.",
												Attributes: map[string]schema.Attribute{
													"icmp_code": schema.Int64Attribute{
														Description:         "icmpCode is the ICMP code for the Rule.Only used when Protocol is VPCSecurityGroupRuleProtocolIcmp.",
														MarkdownDescription: "icmpCode is the ICMP code for the Rule.Only used when Protocol is VPCSecurityGroupRuleProtocolIcmp.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"icmp_type": schema.Int64Attribute{
														Description:         "icmpType is the ICMP type for the Rule.Only used when Protocol is VPCSecurityGroupRuleProtocolIcmp.",
														MarkdownDescription: "icmpType is the ICMP type for the Rule.Only used when Protocol is VPCSecurityGroupRuleProtocolIcmp.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"port_range": schema.SingleNestedAttribute{
														Description:         "portRange is a range of ports allowed for the Rule's remote.",
														MarkdownDescription: "portRange is a range of ports allowed for the Rule's remote.",
														Attributes: map[string]schema.Attribute{
															"maximum_port": schema.Int64Attribute{
																Description:         "maximumPort is the inclusive upper range of ports.",
																MarkdownDescription: "maximumPort is the inclusive upper range of ports.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(1),
																	int64validator.AtMost(65535),
																},
															},

															"minimum_port": schema.Int64Attribute{
																Description:         "minimumPort is the inclusive lower range of ports.",
																MarkdownDescription: "minimumPort is the inclusive lower range of ports.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(1),
																	int64validator.AtMost(65535),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"protocol": schema.StringAttribute{
														Description:         "protocol defines the traffic protocol used for the Security Group Rule.",
														MarkdownDescription: "protocol defines the traffic protocol used for the Security Group Rule.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("all", "icmp", "tcp", "udp"),
														},
													},

													"remotes": schema.ListNestedAttribute{
														Description:         "remotes is a set of VPCSecurityGroupRuleRemote's that define the traffic allowed by the Rule's remote.Specifying multiple VPCSecurityGroupRuleRemote's creates a unique Security Group Rule with the shared Protocol, PortRange, etc.This allows for easier management of Security Group Rule's for sets of CIDR's, IP's, etc.",
														MarkdownDescription: "remotes is a set of VPCSecurityGroupRuleRemote's that define the traffic allowed by the Rule's remote.Specifying multiple VPCSecurityGroupRuleRemote's creates a unique Security Group Rule with the shared Protocol, PortRange, etc.This allows for easier management of Security Group Rule's for sets of CIDR's, IP's, etc.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"address": schema.StringAttribute{
																	Description:         " address is the address to use for the remote's destination/source.Only used when remoteType is VPCSecurityGroupRuleRemoteTypeAddress.",
																	MarkdownDescription: " address is the address to use for the remote's destination/source.Only used when remoteType is VPCSecurityGroupRuleRemoteTypeAddress.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"cidr_subnet_name": schema.StringAttribute{
																	Description:         "cidrSubnetName is the name of the VPC Subnet to retrieve the CIDR from, to use for the remote's destination/source.Only used when remoteType is VPCSecurityGroupRuleRemoteTypeCIDR.",
																	MarkdownDescription: "cidrSubnetName is the name of the VPC Subnet to retrieve the CIDR from, to use for the remote's destination/source.Only used when remoteType is VPCSecurityGroupRuleRemoteTypeCIDR.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"remote_type": schema.StringAttribute{
																	Description:         "remoteType defines the type of filter to define for the remote's destination/source.",
																	MarkdownDescription: "remoteType defines the type of filter to define for the remote's destination/source.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("any", "cidr", "address", "sg"),
																	},
																},

																"security_group_name": schema.StringAttribute{
																	Description:         "securityGroupName is the name of the VPC Security Group to use for the remote's destination/source.Only used when remoteType is VPCSecurityGroupRuleRemoteTypeSG",
																	MarkdownDescription: "securityGroupName is the name of the VPC Security Group to use for the remote's destination/source.Only used when remoteType is VPCSecurityGroupRuleRemoteTypeSG",
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
												Required: false,
												Optional: true,
												Computed: false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"tags": schema.ListAttribute{
									Description:         "tags are tags to add to the Security Group.",
									MarkdownDescription: "tags are tags to add to the Security Group.",
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

					"vpc_subnets": schema.ListNestedAttribute{
						Description:         "vpcSubnets contains information about IBM Cloud VPC Subnet resources.when omitted system will create the subnets in all the zone corresponding to VPC.Region, with name CLUSTER_NAME-vpcsubnet-ZONE_NAME.possible values can be found here https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-creating-power-virtual-server.when VPCSubnets[].ID is set, its expected that there exist a subnet with ID or else system will give error.when VPCSubnets[].Zone is not set, a random zone is picked from available zones of VPC.Region.when VPCSubnets[].Name is not set, system will set name as CLUSTER_NAME-vpcsubnet-INDEX.if subnet with name VPCSubnets[].Name not found, system will create new subnet in VPCSubnets[].Zone.",
						MarkdownDescription: "vpcSubnets contains information about IBM Cloud VPC Subnet resources.when omitted system will create the subnets in all the zone corresponding to VPC.Region, with name CLUSTER_NAME-vpcsubnet-ZONE_NAME.possible values can be found here https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-creating-power-virtual-server.when VPCSubnets[].ID is set, its expected that there exist a subnet with ID or else system will give error.when VPCSubnets[].Zone is not set, a random zone is picked from available zones of VPC.Region.when VPCSubnets[].Name is not set, system will set name as CLUSTER_NAME-vpcsubnet-INDEX.if subnet with name VPCSubnets[].Name not found, system will create new subnet in VPCSubnets[].Zone.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"cidr": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"id": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(64),
										stringvalidator.RegexMatches(regexp.MustCompile(`^[-0-9a-z_]+$`), ""),
									},
								},

								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(63),
										stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`), ""),
									},
								},

								"zone": schema.StringAttribute{
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

					"zone": schema.StringAttribute{
						Description:         "zone is the name of Power VS zone where the cluster will be createdpossible values can be found here https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-creating-power-virtual-server.when powervs.cluster.x-k8s.io/create-infra=true annotation is set on IBMPowerVSCluster resource,1. it is expected to set the zone, not setting will result in webhook error.2. the zone should have PER capabilities, or else system will give error.",
						MarkdownDescription: "zone is the name of Power VS zone where the cluster will be createdpossible values can be found here https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-creating-power-virtual-server.when powervs.cluster.x-k8s.io/create-infra=true annotation is set on IBMPowerVSCluster resource,1. it is expected to set the zone, not setting will result in webhook error.2. the zone should have PER capabilities, or else system will give error.",
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

func (r *InfrastructureClusterXK8SIoIbmpowerVsclusterV1Beta2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_infrastructure_cluster_x_k8s_io_ibm_power_vs_cluster_v1beta2_manifest")

	var model InfrastructureClusterXK8SIoIbmpowerVsclusterV1Beta2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("infrastructure.cluster.x-k8s.io/v1beta2")
	model.Kind = pointer.String("IBMPowerVSCluster")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
