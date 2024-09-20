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
	_ datasource.DataSource = &InfrastructureClusterXK8SIoIbmvpcclusterV1Beta2Manifest{}
)

func NewInfrastructureClusterXK8SIoIbmvpcclusterV1Beta2Manifest() datasource.DataSource {
	return &InfrastructureClusterXK8SIoIbmvpcclusterV1Beta2Manifest{}
}

type InfrastructureClusterXK8SIoIbmvpcclusterV1Beta2Manifest struct{}

type InfrastructureClusterXK8SIoIbmvpcclusterV1Beta2ManifestData struct {
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
		ControlPlaneLoadBalancer *struct {
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
		} `tfsdk:"control_plane_load_balancer" json:"controlPlaneLoadBalancer,omitempty"`
		Image *struct {
			CosBucket       *string `tfsdk:"cos_bucket" json:"cosBucket,omitempty"`
			CosBucketRegion *string `tfsdk:"cos_bucket_region" json:"cosBucketRegion,omitempty"`
			CosInstance     *string `tfsdk:"cos_instance" json:"cosInstance,omitempty"`
			CosObject       *string `tfsdk:"cos_object" json:"cosObject,omitempty"`
			Crn             *string `tfsdk:"crn" json:"crn,omitempty"`
			Name            *string `tfsdk:"name" json:"name,omitempty"`
			OperatingSystem *string `tfsdk:"operating_system" json:"operatingSystem,omitempty"`
			ResourceGroup   *struct {
				Id   *string `tfsdk:"id" json:"id,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"resource_group" json:"resourceGroup,omitempty"`
		} `tfsdk:"image" json:"image,omitempty"`
		Network *struct {
			ControlPlaneSubnets *[]struct {
				Cidr *string `tfsdk:"cidr" json:"cidr,omitempty"`
				Id   *string `tfsdk:"id" json:"id,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
				Zone *string `tfsdk:"zone" json:"zone,omitempty"`
			} `tfsdk:"control_plane_subnets" json:"controlPlaneSubnets,omitempty"`
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
			ResourceGroup *struct {
				Id   *string `tfsdk:"id" json:"id,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"resource_group" json:"resourceGroup,omitempty"`
			SecurityGroups *[]struct {
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
			} `tfsdk:"security_groups" json:"securityGroups,omitempty"`
			Vpc *struct {
				Id   *string `tfsdk:"id" json:"id,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"vpc" json:"vpc,omitempty"`
			WorkerSubnets *[]struct {
				Cidr *string `tfsdk:"cidr" json:"cidr,omitempty"`
				Id   *string `tfsdk:"id" json:"id,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
				Zone *string `tfsdk:"zone" json:"zone,omitempty"`
			} `tfsdk:"worker_subnets" json:"workerSubnets,omitempty"`
		} `tfsdk:"network" json:"network,omitempty"`
		Region        *string `tfsdk:"region" json:"region,omitempty"`
		ResourceGroup *string `tfsdk:"resource_group" json:"resourceGroup,omitempty"`
		Vpc           *string `tfsdk:"vpc" json:"vpc,omitempty"`
		Zone          *string `tfsdk:"zone" json:"zone,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *InfrastructureClusterXK8SIoIbmvpcclusterV1Beta2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_infrastructure_cluster_x_k8s_io_ibmvpc_cluster_v1beta2_manifest"
}

func (r *InfrastructureClusterXK8SIoIbmvpcclusterV1Beta2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "IBMVPCCluster is the Schema for the ibmvpcclusters API.",
		MarkdownDescription: "IBMVPCCluster is the Schema for the ibmvpcclusters API.",
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
				Description:         "IBMVPCClusterSpec defines the desired state of IBMVPCCluster.",
				MarkdownDescription: "IBMVPCClusterSpec defines the desired state of IBMVPCCluster.",
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

					"control_plane_load_balancer": schema.SingleNestedAttribute{
						Description:         "ControlPlaneLoadBalancer is optional configuration for customizing control plane behavior. Use this for legacy support, use Network.LoadBalancers for the extended VPC support.",
						MarkdownDescription: "ControlPlaneLoadBalancer is optional configuration for customizing control plane behavior. Use this for legacy support, use Network.LoadBalancers for the extended VPC support.",
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
											Description:         "protocol defines the protocol to use for the VPC Load Balancer Listener. Will default to TCP protocol if not specified.",
											MarkdownDescription: "protocol defines the protocol to use for the VPC Load Balancer Listener. Will default to TCP protocol if not specified.",
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
								Description:         "securityGroups defines the Security Groups to attach to the load balancer. Security Groups defined here are expected to already exist when the load balancer is reconciled (these do not get created when reconciling the load balancer).",
								MarkdownDescription: "securityGroups defines the Security Groups to attach to the load balancer. Security Groups defined here are expected to already exist when the load balancer is reconciled (these do not get created when reconciling the load balancer).",
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
								Description:         "subnets defines the VPC Subnets to attach to the load balancer. Subnets defiens here are expected to already exist when the load balancer is reconciled (these do not get created when reconciling the load balancer).",
								MarkdownDescription: "subnets defines the VPC Subnets to attach to the load balancer. Subnets defiens here are expected to already exist when the load balancer is reconciled (these do not get created when reconciling the load balancer).",
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"image": schema.SingleNestedAttribute{
						Description:         "image represents the Image details used for the cluster.",
						MarkdownDescription: "image represents the Image details used for the cluster.",
						Attributes: map[string]schema.Attribute{
							"cos_bucket": schema.StringAttribute{
								Description:         "cosBucket is the name of the IBM Cloud COS Bucket containing the source of the image, if necessary.",
								MarkdownDescription: "cosBucket is the name of the IBM Cloud COS Bucket containing the source of the image, if necessary.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cos_bucket_region": schema.StringAttribute{
								Description:         "cosBucketRegion is the COS region the bucket is in.",
								MarkdownDescription: "cosBucketRegion is the COS region the bucket is in.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cos_instance": schema.StringAttribute{
								Description:         "cosInstance is the name of the IBM Cloud COS Instance containing the source of the image, if necessary.",
								MarkdownDescription: "cosInstance is the name of the IBM Cloud COS Instance containing the source of the image, if necessary.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cos_object": schema.StringAttribute{
								Description:         "cosObject is the name of a IBM Cloud COS Object used as the source of the image, if necessary.",
								MarkdownDescription: "cosObject is the name of a IBM Cloud COS Object used as the source of the image, if necessary.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"crn": schema.StringAttribute{
								Description:         "crn is the IBM Cloud CRN of the existing VPC Custom Image.",
								MarkdownDescription: "crn is the IBM Cloud CRN of the existing VPC Custom Image.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "name is the name of the desired VPC Custom Image.",
								MarkdownDescription: "name is the name of the desired VPC Custom Image.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(63),
									stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`), ""),
								},
							},

							"operating_system": schema.StringAttribute{
								Description:         "operatingSystem is the Custom Image's Operating System name.",
								MarkdownDescription: "operatingSystem is the Custom Image's Operating System name.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resource_group": schema.SingleNestedAttribute{
								Description:         "resourceGroup is the Resource Group to create the Custom Image in.",
								MarkdownDescription: "resourceGroup is the Resource Group to create the Custom Image in.",
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description:         "id defines the IBM Cloud Resource ID.",
										MarkdownDescription: "id defines the IBM Cloud Resource ID.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "name defines the IBM Cloud Resource Name.",
										MarkdownDescription: "name defines the IBM Cloud Resource Name.",
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

					"network": schema.SingleNestedAttribute{
						Description:         "network represents the VPC network to use for the cluster.",
						MarkdownDescription: "network represents the VPC network to use for the cluster.",
						Attributes: map[string]schema.Attribute{
							"control_plane_subnets": schema.ListNestedAttribute{
								Description:         "controlPlaneSubnets is a set of Subnet's which define the Control Plane subnets.",
								MarkdownDescription: "controlPlaneSubnets is a set of Subnet's which define the Control Plane subnets.",
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

							"load_balancers": schema.ListNestedAttribute{
								Description:         "loadBalancers is a set of VPC Load Balancer definitions to use for the cluster.",
								MarkdownDescription: "loadBalancers is a set of VPC Load Balancer definitions to use for the cluster.",
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
														Description:         "protocol defines the protocol to use for the VPC Load Balancer Listener. Will default to TCP protocol if not specified.",
														MarkdownDescription: "protocol defines the protocol to use for the VPC Load Balancer Listener. Will default to TCP protocol if not specified.",
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
											Description:         "securityGroups defines the Security Groups to attach to the load balancer. Security Groups defined here are expected to already exist when the load balancer is reconciled (these do not get created when reconciling the load balancer).",
											MarkdownDescription: "securityGroups defines the Security Groups to attach to the load balancer. Security Groups defined here are expected to already exist when the load balancer is reconciled (these do not get created when reconciling the load balancer).",
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
											Description:         "subnets defines the VPC Subnets to attach to the load balancer. Subnets defiens here are expected to already exist when the load balancer is reconciled (these do not get created when reconciling the load balancer).",
											MarkdownDescription: "subnets defines the VPC Subnets to attach to the load balancer. Subnets defiens here are expected to already exist when the load balancer is reconciled (these do not get created when reconciling the load balancer).",
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

							"resource_group": schema.SingleNestedAttribute{
								Description:         "resourceGroup is the Resource Group containing all of the newtork resources. This can be different than the Resource Group containing the remaining cluster resources.",
								MarkdownDescription: "resourceGroup is the Resource Group containing all of the newtork resources. This can be different than the Resource Group containing the remaining cluster resources.",
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description:         "id defines the IBM Cloud Resource ID.",
										MarkdownDescription: "id defines the IBM Cloud Resource ID.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "name defines the IBM Cloud Resource Name.",
										MarkdownDescription: "name defines the IBM Cloud Resource Name.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"security_groups": schema.ListNestedAttribute{
								Description:         "securityGroups is a set of VPCSecurityGroup's which define the VPC Security Groups that manage traffic within and out of the VPC.",
								MarkdownDescription: "securityGroups is a set of VPCSecurityGroup's which define the VPC Security Groups that manage traffic within and out of the VPC.",
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
														Description:         "destination is a VPCSecurityGroupRulePrototype which defines the destination of outbound traffic for the Security Group Rule. Only used when direction is VPCSecurityGroupRuleDirectionOutbound.",
														MarkdownDescription: "destination is a VPCSecurityGroupRulePrototype which defines the destination of outbound traffic for the Security Group Rule. Only used when direction is VPCSecurityGroupRuleDirectionOutbound.",
														Attributes: map[string]schema.Attribute{
															"icmp_code": schema.Int64Attribute{
																Description:         "icmpCode is the ICMP code for the Rule. Only used when Protocol is VPCSecurityGroupRuleProtocolIcmp.",
																MarkdownDescription: "icmpCode is the ICMP code for the Rule. Only used when Protocol is VPCSecurityGroupRuleProtocolIcmp.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"icmp_type": schema.Int64Attribute{
																Description:         "icmpType is the ICMP type for the Rule. Only used when Protocol is VPCSecurityGroupRuleProtocolIcmp.",
																MarkdownDescription: "icmpType is the ICMP type for the Rule. Only used when Protocol is VPCSecurityGroupRuleProtocolIcmp.",
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
																Description:         "remotes is a set of VPCSecurityGroupRuleRemote's that define the traffic allowed by the Rule's remote. Specifying multiple VPCSecurityGroupRuleRemote's creates a unique Security Group Rule with the shared Protocol, PortRange, etc. This allows for easier management of Security Group Rule's for sets of CIDR's, IP's, etc.",
																MarkdownDescription: "remotes is a set of VPCSecurityGroupRuleRemote's that define the traffic allowed by the Rule's remote. Specifying multiple VPCSecurityGroupRuleRemote's creates a unique Security Group Rule with the shared Protocol, PortRange, etc. This allows for easier management of Security Group Rule's for sets of CIDR's, IP's, etc.",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"address": schema.StringAttribute{
																			Description:         " address is the address to use for the remote's destination/source. Only used when remoteType is VPCSecurityGroupRuleRemoteTypeAddress.",
																			MarkdownDescription: " address is the address to use for the remote's destination/source. Only used when remoteType is VPCSecurityGroupRuleRemoteTypeAddress.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"cidr_subnet_name": schema.StringAttribute{
																			Description:         "cidrSubnetName is the name of the VPC Subnet to retrieve the CIDR from, to use for the remote's destination/source. Only used when remoteType is VPCSecurityGroupRuleRemoteTypeCIDR.",
																			MarkdownDescription: "cidrSubnetName is the name of the VPC Subnet to retrieve the CIDR from, to use for the remote's destination/source. Only used when remoteType is VPCSecurityGroupRuleRemoteTypeCIDR.",
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
																			Description:         "securityGroupName is the name of the VPC Security Group to use for the remote's destination/source. Only used when remoteType is VPCSecurityGroupRuleRemoteTypeSG",
																			MarkdownDescription: "securityGroupName is the name of the VPC Security Group to use for the remote's destination/source. Only used when remoteType is VPCSecurityGroupRuleRemoteTypeSG",
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
														Description:         "source is a VPCSecurityGroupRulePrototype which defines the source of inbound traffic for the Security Group Rule. Only used when direction is VPCSecurityGroupRuleDirectionInbound.",
														MarkdownDescription: "source is a VPCSecurityGroupRulePrototype which defines the source of inbound traffic for the Security Group Rule. Only used when direction is VPCSecurityGroupRuleDirectionInbound.",
														Attributes: map[string]schema.Attribute{
															"icmp_code": schema.Int64Attribute{
																Description:         "icmpCode is the ICMP code for the Rule. Only used when Protocol is VPCSecurityGroupRuleProtocolIcmp.",
																MarkdownDescription: "icmpCode is the ICMP code for the Rule. Only used when Protocol is VPCSecurityGroupRuleProtocolIcmp.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"icmp_type": schema.Int64Attribute{
																Description:         "icmpType is the ICMP type for the Rule. Only used when Protocol is VPCSecurityGroupRuleProtocolIcmp.",
																MarkdownDescription: "icmpType is the ICMP type for the Rule. Only used when Protocol is VPCSecurityGroupRuleProtocolIcmp.",
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
																Description:         "remotes is a set of VPCSecurityGroupRuleRemote's that define the traffic allowed by the Rule's remote. Specifying multiple VPCSecurityGroupRuleRemote's creates a unique Security Group Rule with the shared Protocol, PortRange, etc. This allows for easier management of Security Group Rule's for sets of CIDR's, IP's, etc.",
																MarkdownDescription: "remotes is a set of VPCSecurityGroupRuleRemote's that define the traffic allowed by the Rule's remote. Specifying multiple VPCSecurityGroupRuleRemote's creates a unique Security Group Rule with the shared Protocol, PortRange, etc. This allows for easier management of Security Group Rule's for sets of CIDR's, IP's, etc.",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"address": schema.StringAttribute{
																			Description:         " address is the address to use for the remote's destination/source. Only used when remoteType is VPCSecurityGroupRuleRemoteTypeAddress.",
																			MarkdownDescription: " address is the address to use for the remote's destination/source. Only used when remoteType is VPCSecurityGroupRuleRemoteTypeAddress.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"cidr_subnet_name": schema.StringAttribute{
																			Description:         "cidrSubnetName is the name of the VPC Subnet to retrieve the CIDR from, to use for the remote's destination/source. Only used when remoteType is VPCSecurityGroupRuleRemoteTypeCIDR.",
																			MarkdownDescription: "cidrSubnetName is the name of the VPC Subnet to retrieve the CIDR from, to use for the remote's destination/source. Only used when remoteType is VPCSecurityGroupRuleRemoteTypeCIDR.",
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
																			Description:         "securityGroupName is the name of the VPC Security Group to use for the remote's destination/source. Only used when remoteType is VPCSecurityGroupRuleRemoteTypeSG",
																			MarkdownDescription: "securityGroupName is the name of the VPC Security Group to use for the remote's destination/source. Only used when remoteType is VPCSecurityGroupRuleRemoteTypeSG",
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

							"vpc": schema.SingleNestedAttribute{
								Description:         "vpc defines the IBM Cloud VPC for extended VPC Infrastructure support.",
								MarkdownDescription: "vpc defines the IBM Cloud VPC for extended VPC Infrastructure support.",
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
								Required: false,
								Optional: true,
								Computed: false,
							},

							"worker_subnets": schema.ListNestedAttribute{
								Description:         "workerSubnets is a set of Subnet's which define the Worker subnets.",
								MarkdownDescription: "workerSubnets is a set of Subnet's which define the Worker subnets.",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"region": schema.StringAttribute{
						Description:         "The IBM Cloud Region the cluster lives in.",
						MarkdownDescription: "The IBM Cloud Region the cluster lives in.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"resource_group": schema.StringAttribute{
						Description:         "The VPC resources should be created under the resource group.",
						MarkdownDescription: "The VPC resources should be created under the resource group.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"vpc": schema.StringAttribute{
						Description:         "The Name of VPC.",
						MarkdownDescription: "The Name of VPC.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"zone": schema.StringAttribute{
						Description:         "The Name of availability zone.",
						MarkdownDescription: "The Name of availability zone.",
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

func (r *InfrastructureClusterXK8SIoIbmvpcclusterV1Beta2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_infrastructure_cluster_x_k8s_io_ibmvpc_cluster_v1beta2_manifest")

	var model InfrastructureClusterXK8SIoIbmvpcclusterV1Beta2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("infrastructure.cluster.x-k8s.io/v1beta2")
	model.Kind = pointer.String("IBMVPCCluster")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
