/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package infrastructure_cluster_x_k8s_io_v1beta2

import (
	"context"
	"fmt"
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
	_ datasource.DataSource = &InfrastructureClusterXK8SIoIbmpowerVsclusterTemplateV1Beta2Manifest{}
)

func NewInfrastructureClusterXK8SIoIbmpowerVsclusterTemplateV1Beta2Manifest() datasource.DataSource {
	return &InfrastructureClusterXK8SIoIbmpowerVsclusterTemplateV1Beta2Manifest{}
}

type InfrastructureClusterXK8SIoIbmpowerVsclusterTemplateV1Beta2Manifest struct{}

type InfrastructureClusterXK8SIoIbmpowerVsclusterTemplateV1Beta2ManifestData struct {
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
		Template *struct {
			Metadata *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"metadata" json:"metadata,omitempty"`
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
						Port *int64 `tfsdk:"port" json:"port,omitempty"`
					} `tfsdk:"additional_listeners" json:"additionalListeners,omitempty"`
					Id     *string `tfsdk:"id" json:"id,omitempty"`
					Name   *string `tfsdk:"name" json:"name,omitempty"`
					Public *bool   `tfsdk:"public" json:"public,omitempty"`
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
					Id   *string `tfsdk:"id" json:"id,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"transit_gateway" json:"transitGateway,omitempty"`
				Vpc *struct {
					Id     *string `tfsdk:"id" json:"id,omitempty"`
					Name   *string `tfsdk:"name" json:"name,omitempty"`
					Region *string `tfsdk:"region" json:"region,omitempty"`
				} `tfsdk:"vpc" json:"vpc,omitempty"`
				VpcSubnets *[]struct {
					Cidr *string `tfsdk:"cidr" json:"cidr,omitempty"`
					Id   *string `tfsdk:"id" json:"id,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
					Zone *string `tfsdk:"zone" json:"zone,omitempty"`
				} `tfsdk:"vpc_subnets" json:"vpcSubnets,omitempty"`
				Zone *string `tfsdk:"zone" json:"zone,omitempty"`
			} `tfsdk:"spec" json:"spec,omitempty"`
		} `tfsdk:"template" json:"template,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *InfrastructureClusterXK8SIoIbmpowerVsclusterTemplateV1Beta2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_infrastructure_cluster_x_k8s_io_ibm_power_vs_cluster_template_v1beta2_manifest"
}

func (r *InfrastructureClusterXK8SIoIbmpowerVsclusterTemplateV1Beta2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "IBMPowerVSClusterTemplate is the schema for IBM Power VS Kubernetes Cluster Templates.",
		MarkdownDescription: "IBMPowerVSClusterTemplate is the schema for IBM Power VS Kubernetes Cluster Templates.",
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
				Description:         "IBMPowerVSClusterTemplateSpec defines the desired state of IBMPowerVSClusterTemplate.",
				MarkdownDescription: "IBMPowerVSClusterTemplateSpec defines the desired state of IBMPowerVSClusterTemplate.",
				Attributes: map[string]schema.Attribute{
					"template": schema.SingleNestedAttribute{
						Description:         "IBMPowerVSClusterTemplateResource describes the data needed to create an IBMPowerVSCluster from a template.",
						MarkdownDescription: "IBMPowerVSClusterTemplateResource describes the data needed to create an IBMPowerVSCluster from a template.",
						Attributes: map[string]schema.Attribute{
							"metadata": schema.SingleNestedAttribute{
								Description:         "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
								MarkdownDescription: "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
										MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.MapAttribute{
										Description:         "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
										MarkdownDescription: "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
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
										Description:         "cosInstance contains options to configure a supporting IBM Cloud COS bucket for this cluster - currently used for nodes requiring Ignition (https://coreos.github.io/ignition/) for bootstrapping (requires BootstrapFormatIgnition feature flag to be enabled). when powervs.cluster.x-k8s.io/create-infra=true annotation is set on IBMPowerVSCluster resource and Ignition is set, then 1. CosInstance.Name should be set not setting will result in webhook error. 2. CosInstance.BucketName should be set not setting will result in webhook error. 3. CosInstance.BucketRegion should be set not setting will result in webhook error.",
										MarkdownDescription: "cosInstance contains options to configure a supporting IBM Cloud COS bucket for this cluster - currently used for nodes requiring Ignition (https://coreos.github.io/ignition/) for bootstrapping (requires BootstrapFormatIgnition feature flag to be enabled). when powervs.cluster.x-k8s.io/create-infra=true annotation is set on IBMPowerVSCluster resource and Ignition is set, then 1. CosInstance.Name should be set not setting will result in webhook error. 2. CosInstance.BucketName should be set not setting will result in webhook error. 3. CosInstance.BucketRegion should be set not setting will result in webhook error.",
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
												Description:         "name defines name of IBM cloud COS instance to be created. when IBMPowerVSCluster.Ignition is set",
												MarkdownDescription: "name defines name of IBM cloud COS instance to be created. when IBMPowerVSCluster.Ignition is set",
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
										Description:         "dhcpServer is contains the configuration to be used while creating a new DHCP server in PowerVS workspace. when the field is omitted, CLUSTER_NAME will be used as DHCPServer.Name and DHCP server will be created. it will automatically create network with name DHCPSERVER<DHCPServer.Name>_Private in PowerVS workspace.",
										MarkdownDescription: "dhcpServer is contains the configuration to be used while creating a new DHCP server in PowerVS workspace. when the field is omitted, CLUSTER_NAME will be used as DHCPServer.Name and DHCP server will be created. it will automatically create network with name DHCPSERVER<DHCPServer.Name>_Private in PowerVS workspace.",
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
										Description:         "loadBalancers is optional configuration for configuring loadbalancers to control plane or data plane nodes. when omitted system will create a default public loadbalancer with name CLUSTER_NAME-loadbalancer. when specified a vpc loadbalancer will be created and controlPlaneEndpoint will be set with associated hostname of loadbalancer. ControlPlaneEndpoint will be set with associated hostname of public loadbalancer. when LoadBalancers[].ID is set, its expected that there exist a loadbalancer with ID or else system will give error. when LoadBalancers[].Name is set, system will first check for loadbalancer with Name, if not exist system will create new loadbalancer. For each loadbalancer a default backed pool and front listener will be configured with port 6443.",
										MarkdownDescription: "loadBalancers is optional configuration for configuring loadbalancers to control plane or data plane nodes. when omitted system will create a default public loadbalancer with name CLUSTER_NAME-loadbalancer. when specified a vpc loadbalancer will be created and controlPlaneEndpoint will be set with associated hostname of loadbalancer. ControlPlaneEndpoint will be set with associated hostname of public loadbalancer. when LoadBalancers[].ID is set, its expected that there exist a loadbalancer with ID or else system will give error. when LoadBalancers[].Name is set, system will first check for loadbalancer with Name, if not exist system will create new loadbalancer. For each loadbalancer a default backed pool and front listener will be configured with port 6443.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"additional_listeners": schema.ListNestedAttribute{
													Description:         "AdditionalListeners sets the additional listeners for the control plane load balancer.",
													MarkdownDescription: "AdditionalListeners sets the additional listeners for the control plane load balancer.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
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
												},

												"name": schema.StringAttribute{
													Description:         "Name sets the name of the VPC load balancer.",
													MarkdownDescription: "Name sets the name of the VPC load balancer.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtMost(63),
													},
												},

												"public": schema.BoolAttribute{
													Description:         "public indicates that load balancer is public or private",
													MarkdownDescription: "public indicates that load balancer is public or private",
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

									"network": schema.SingleNestedAttribute{
										Description:         "Network is the reference to the Network to use for this cluster. when the field is omitted, A DHCP service will be created in the Power VS workspace and its private network will be used. the DHCP service created network will have the following name format 1. in the case of DHCPServer.Name is not set the name will be DHCPSERVER<CLUSTER_NAME>_Private. 2. if DHCPServer.Name is set the name will be DHCPSERVER<DHCPServer.Name>_Private. when Network.ID is set, its expected that there exist a network in PowerVS workspace with id or else system will give error. when Network.Name is set, system will first check for network with Name in PowerVS workspace, if not exist network will be created by DHCP service. Network.RegEx is not yet supported and system will ignore the value.",
										MarkdownDescription: "Network is the reference to the Network to use for this cluster. when the field is omitted, A DHCP service will be created in the Power VS workspace and its private network will be used. the DHCP service created network will have the following name format 1. in the case of DHCPServer.Name is not set the name will be DHCPSERVER<CLUSTER_NAME>_Private. 2. if DHCPServer.Name is set the name will be DHCPSERVER<DHCPServer.Name>_Private. when Network.ID is set, its expected that there exist a network in PowerVS workspace with id or else system will give error. when Network.Name is set, system will first check for network with Name in PowerVS workspace, if not exist network will be created by DHCP service. Network.RegEx is not yet supported and system will ignore the value.",
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
												Description:         "Regular expression to match resource, In case of multiple resources matches the provided regular expression the first matched resource will be selected",
												MarkdownDescription: "Regular expression to match resource, In case of multiple resources matches the provided regular expression the first matched resource will be selected",
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
										Description:         "resourceGroup name under which the resources will be created. when powervs.cluster.x-k8s.io/create-infra=true annotation is set on IBMPowerVSCluster resource, 1. it is expected to set the ResourceGroup.Name, not setting will result in webhook error. ServiceInstance.ID and ServiceInstance.Regex is not yet supported and system will ignore the value.",
										MarkdownDescription: "resourceGroup name under which the resources will be created. when powervs.cluster.x-k8s.io/create-infra=true annotation is set on IBMPowerVSCluster resource, 1. it is expected to set the ResourceGroup.Name, not setting will result in webhook error. ServiceInstance.ID and ServiceInstance.Regex is not yet supported and system will ignore the value.",
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
												Description:         "Regular expression to match resource, In case of multiple resources matches the provided regular expression the first matched resource will be selected",
												MarkdownDescription: "Regular expression to match resource, In case of multiple resources matches the provided regular expression the first matched resource will be selected",
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
										Description:         "serviceInstance is the reference to the Power VS server workspace on which the server instance(VM) will be created. Power VS server workspace is a container for all Power VS instances at a specific geographic region. serviceInstance can be created via IBM Cloud catalog or CLI. supported serviceInstance identifier in PowerVSResource are Name and ID and that can be obtained from IBM Cloud UI or IBM Cloud cli. More detail about Power VS service instance. https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-creating-power-virtual-server when omitted system will dynamically create the service instance with name CLUSTER_NAME-serviceInstance. when ServiceInstance.ID is set, its expected that there exist a service instance in PowerVS workspace with id or else system will give error. when ServiceInstance.Name is set, system will first check for service instance with Name in PowerVS workspace, if not exist system will create new instance. ServiceInstance.Regex is not yet supported not yet supported and system will ignore the value.",
										MarkdownDescription: "serviceInstance is the reference to the Power VS server workspace on which the server instance(VM) will be created. Power VS server workspace is a container for all Power VS instances at a specific geographic region. serviceInstance can be created via IBM Cloud catalog or CLI. supported serviceInstance identifier in PowerVSResource are Name and ID and that can be obtained from IBM Cloud UI or IBM Cloud cli. More detail about Power VS service instance. https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-creating-power-virtual-server when omitted system will dynamically create the service instance with name CLUSTER_NAME-serviceInstance. when ServiceInstance.ID is set, its expected that there exist a service instance in PowerVS workspace with id or else system will give error. when ServiceInstance.Name is set, system will first check for service instance with Name in PowerVS workspace, if not exist system will create new instance. ServiceInstance.Regex is not yet supported not yet supported and system will ignore the value.",
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
												Description:         "Regular expression to match resource, In case of multiple resources matches the provided regular expression the first matched resource will be selected",
												MarkdownDescription: "Regular expression to match resource, In case of multiple resources matches the provided regular expression the first matched resource will be selected",
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
										Description:         "ServiceInstanceID is the id of the power cloud instance where the vsi instance will get deployed. Deprecated: use ServiceInstance instead",
										MarkdownDescription: "ServiceInstanceID is the id of the power cloud instance where the vsi instance will get deployed. Deprecated: use ServiceInstance instead",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"transit_gateway": schema.SingleNestedAttribute{
										Description:         "transitGateway contains information about IBM Cloud TransitGateway IBM Cloud TransitGateway helps in establishing network connectivity between IBM Cloud Power VS and VPC infrastructure more information about TransitGateway can be found here https://www.ibm.com/products/transit-gateway. when TransitGateway.ID is set, its expected that there exist a TransitGateway with ID or else system will give error. when TransitGateway.Name is set, system will first check for TransitGateway with Name, if not exist system will create new TransitGateway.",
										MarkdownDescription: "transitGateway contains information about IBM Cloud TransitGateway IBM Cloud TransitGateway helps in establishing network connectivity between IBM Cloud Power VS and VPC infrastructure more information about TransitGateway can be found here https://www.ibm.com/products/transit-gateway. when TransitGateway.ID is set, its expected that there exist a TransitGateway with ID or else system will give error. when TransitGateway.Name is set, system will first check for TransitGateway with Name, if not exist system will create new TransitGateway.",
										Attributes: map[string]schema.Attribute{
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
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"vpc": schema.SingleNestedAttribute{
										Description:         "vpc contains information about IBM Cloud VPC resources. when omitted system will dynamically create the VPC with name CLUSTER_NAME-vpc. when VPC.ID is set, its expected that there exist a VPC with ID or else system will give error. when VPC.Name is set, system will first check for VPC with Name, if not exist system will create new VPC. when powervs.cluster.x-k8s.io/create-infra=true annotation is set on IBMPowerVSCluster resource, 1. it is expected to set the VPC.Region, not setting will result in webhook error.",
										MarkdownDescription: "vpc contains information about IBM Cloud VPC resources. when omitted system will dynamically create the VPC with name CLUSTER_NAME-vpc. when VPC.ID is set, its expected that there exist a VPC with ID or else system will give error. when VPC.Name is set, system will first check for VPC with Name, if not exist system will create new VPC. when powervs.cluster.x-k8s.io/create-infra=true annotation is set on IBMPowerVSCluster resource, 1. it is expected to set the VPC.Region, not setting will result in webhook error.",
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description:         "id of resource.",
												MarkdownDescription: "id of resource.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
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
												},
											},

											"region": schema.StringAttribute{
												Description:         "region of IBM Cloud VPC. when powervs.cluster.x-k8s.io/create-infra=true annotation is set on IBMPowerVSCluster resource, it is expected to set the region, not setting will result in webhook error.",
												MarkdownDescription: "region of IBM Cloud VPC. when powervs.cluster.x-k8s.io/create-infra=true annotation is set on IBMPowerVSCluster resource, it is expected to set the region, not setting will result in webhook error.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"vpc_subnets": schema.ListNestedAttribute{
										Description:         "vpcSubnets contains information about IBM Cloud VPC Subnet resources. when omitted system will create the subnets in all the zone corresponding to VPC.Region, with name CLUSTER_NAME-vpcsubnet-ZONE_NAME. possible values can be found here https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-creating-power-virtual-server. when VPCSubnets[].ID is set, its expected that there exist a subnet with ID or else system will give error. when VPCSubnets[].Zone is not set, a random zone is picked from available zones of VPC.Region. when VPCSubnets[].Name is not set, system will set name as CLUSTER_NAME-vpcsubnet-INDEX. if subnet with name VPCSubnets[].Name not found, system will create new subnet in VPCSubnets[].Zone.",
										MarkdownDescription: "vpcSubnets contains information about IBM Cloud VPC Subnet resources. when omitted system will create the subnets in all the zone corresponding to VPC.Region, with name CLUSTER_NAME-vpcsubnet-ZONE_NAME. possible values can be found here https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-creating-power-virtual-server. when VPCSubnets[].ID is set, its expected that there exist a subnet with ID or else system will give error. when VPCSubnets[].Zone is not set, a random zone is picked from available zones of VPC.Region. when VPCSubnets[].Name is not set, system will set name as CLUSTER_NAME-vpcsubnet-INDEX. if subnet with name VPCSubnets[].Name not found, system will create new subnet in VPCSubnets[].Zone.",
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
												},

												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
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
										Description:         "zone is the name of Power VS zone where the cluster will be created possible values can be found here https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-creating-power-virtual-server. when powervs.cluster.x-k8s.io/create-infra=true annotation is set on IBMPowerVSCluster resource, 1. it is expected to set the zone, not setting will result in webhook error. 2. the zone should have PER capabilities, or else system will give error.",
										MarkdownDescription: "zone is the name of Power VS zone where the cluster will be created possible values can be found here https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-creating-power-virtual-server. when powervs.cluster.x-k8s.io/create-infra=true annotation is set on IBMPowerVSCluster resource, 1. it is expected to set the zone, not setting will result in webhook error. 2. the zone should have PER capabilities, or else system will give error.",
										Required:            false,
										Optional:            true,
										Computed:            false,
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
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *InfrastructureClusterXK8SIoIbmpowerVsclusterTemplateV1Beta2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_infrastructure_cluster_x_k8s_io_ibm_power_vs_cluster_template_v1beta2_manifest")

	var model InfrastructureClusterXK8SIoIbmpowerVsclusterTemplateV1Beta2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("infrastructure.cluster.x-k8s.io/v1beta2")
	model.Kind = pointer.String("IBMPowerVSClusterTemplate")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
