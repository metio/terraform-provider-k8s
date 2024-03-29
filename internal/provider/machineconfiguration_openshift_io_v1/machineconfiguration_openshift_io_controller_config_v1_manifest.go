/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package machineconfiguration_openshift_io_v1

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
	_ datasource.DataSource = &MachineconfigurationOpenshiftIoControllerConfigV1Manifest{}
)

func NewMachineconfigurationOpenshiftIoControllerConfigV1Manifest() datasource.DataSource {
	return &MachineconfigurationOpenshiftIoControllerConfigV1Manifest{}
}

type MachineconfigurationOpenshiftIoControllerConfigV1Manifest struct{}

type MachineconfigurationOpenshiftIoControllerConfigV1ManifestData struct {
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
		AdditionalTrustBundle          *string `tfsdk:"additional_trust_bundle" json:"additionalTrustBundle,omitempty"`
		BaseOSContainerImage           *string `tfsdk:"base_os_container_image" json:"baseOSContainerImage,omitempty"`
		BaseOSExtensionsContainerImage *string `tfsdk:"base_os_extensions_container_image" json:"baseOSExtensionsContainerImage,omitempty"`
		CloudProviderCAData            *string `tfsdk:"cloud_provider_ca_data" json:"cloudProviderCAData,omitempty"`
		CloudProviderConfig            *string `tfsdk:"cloud_provider_config" json:"cloudProviderConfig,omitempty"`
		ClusterDNSIP                   *string `tfsdk:"cluster_dnsip" json:"clusterDNSIP,omitempty"`
		Dns                            *struct {
			ApiVersion *string            `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Kind       *string            `tfsdk:"kind" json:"kind,omitempty"`
			Metadata   *map[string]string `tfsdk:"metadata" json:"metadata,omitempty"`
			Spec       *struct {
				BaseDomain *string `tfsdk:"base_domain" json:"baseDomain,omitempty"`
				Platform   *struct {
					Aws *struct {
						PrivateZoneIAMRole *string `tfsdk:"private_zone_iam_role" json:"privateZoneIAMRole,omitempty"`
					} `tfsdk:"aws" json:"aws,omitempty"`
					Type *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"platform" json:"platform,omitempty"`
				PrivateZone *struct {
					Id   *string            `tfsdk:"id" json:"id,omitempty"`
					Tags *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
				} `tfsdk:"private_zone" json:"privateZone,omitempty"`
				PublicZone *struct {
					Id   *string            `tfsdk:"id" json:"id,omitempty"`
					Tags *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
				} `tfsdk:"public_zone" json:"publicZone,omitempty"`
			} `tfsdk:"spec" json:"spec,omitempty"`
			Status *map[string]string `tfsdk:"status" json:"status,omitempty"`
		} `tfsdk:"dns" json:"dns,omitempty"`
		EtcdDiscoveryDomain     *string `tfsdk:"etcd_discovery_domain" json:"etcdDiscoveryDomain,omitempty"`
		ImageRegistryBundleData *[]struct {
			Data *string `tfsdk:"data" json:"data,omitempty"`
			File *string `tfsdk:"file" json:"file,omitempty"`
		} `tfsdk:"image_registry_bundle_data" json:"imageRegistryBundleData,omitempty"`
		ImageRegistryBundleUserData *[]struct {
			Data *string `tfsdk:"data" json:"data,omitempty"`
			File *string `tfsdk:"file" json:"file,omitempty"`
		} `tfsdk:"image_registry_bundle_user_data" json:"imageRegistryBundleUserData,omitempty"`
		Images *map[string]string `tfsdk:"images" json:"images,omitempty"`
		Infra  *struct {
			ApiVersion *string            `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Kind       *string            `tfsdk:"kind" json:"kind,omitempty"`
			Metadata   *map[string]string `tfsdk:"metadata" json:"metadata,omitempty"`
			Spec       *struct {
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
			Status *struct {
				ApiServerInternalURI   *string `tfsdk:"api_server_internal_uri" json:"apiServerInternalURI,omitempty"`
				ApiServerURL           *string `tfsdk:"api_server_url" json:"apiServerURL,omitempty"`
				ControlPlaneTopology   *string `tfsdk:"control_plane_topology" json:"controlPlaneTopology,omitempty"`
				CpuPartitioning        *string `tfsdk:"cpu_partitioning" json:"cpuPartitioning,omitempty"`
				EtcdDiscoveryDomain    *string `tfsdk:"etcd_discovery_domain" json:"etcdDiscoveryDomain,omitempty"`
				InfrastructureName     *string `tfsdk:"infrastructure_name" json:"infrastructureName,omitempty"`
				InfrastructureTopology *string `tfsdk:"infrastructure_topology" json:"infrastructureTopology,omitempty"`
				Platform               *string `tfsdk:"platform" json:"platform,omitempty"`
				PlatformStatus         *struct {
					AlibabaCloud *struct {
						Region          *string `tfsdk:"region" json:"region,omitempty"`
						ResourceGroupID *string `tfsdk:"resource_group_id" json:"resourceGroupID,omitempty"`
						ResourceTags    *[]struct {
							Key   *string `tfsdk:"key" json:"key,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"resource_tags" json:"resourceTags,omitempty"`
					} `tfsdk:"alibaba_cloud" json:"alibabaCloud,omitempty"`
					Aws *struct {
						Region       *string `tfsdk:"region" json:"region,omitempty"`
						ResourceTags *[]struct {
							Key   *string `tfsdk:"key" json:"key,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"resource_tags" json:"resourceTags,omitempty"`
						ServiceEndpoints *[]struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
							Url  *string `tfsdk:"url" json:"url,omitempty"`
						} `tfsdk:"service_endpoints" json:"serviceEndpoints,omitempty"`
					} `tfsdk:"aws" json:"aws,omitempty"`
					Azure *struct {
						ArmEndpoint              *string `tfsdk:"arm_endpoint" json:"armEndpoint,omitempty"`
						CloudName                *string `tfsdk:"cloud_name" json:"cloudName,omitempty"`
						NetworkResourceGroupName *string `tfsdk:"network_resource_group_name" json:"networkResourceGroupName,omitempty"`
						ResourceGroupName        *string `tfsdk:"resource_group_name" json:"resourceGroupName,omitempty"`
						ResourceTags             *[]struct {
							Key   *string `tfsdk:"key" json:"key,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"resource_tags" json:"resourceTags,omitempty"`
					} `tfsdk:"azure" json:"azure,omitempty"`
					Baremetal *struct {
						ApiServerInternalIP  *string   `tfsdk:"api_server_internal_ip" json:"apiServerInternalIP,omitempty"`
						ApiServerInternalIPs *[]string `tfsdk:"api_server_internal_i_ps" json:"apiServerInternalIPs,omitempty"`
						IngressIP            *string   `tfsdk:"ingress_ip" json:"ingressIP,omitempty"`
						IngressIPs           *[]string `tfsdk:"ingress_i_ps" json:"ingressIPs,omitempty"`
						MachineNetworks      *[]string `tfsdk:"machine_networks" json:"machineNetworks,omitempty"`
						NodeDNSIP            *string   `tfsdk:"node_dnsip" json:"nodeDNSIP,omitempty"`
					} `tfsdk:"baremetal" json:"baremetal,omitempty"`
					EquinixMetal *struct {
						ApiServerInternalIP *string `tfsdk:"api_server_internal_ip" json:"apiServerInternalIP,omitempty"`
						IngressIP           *string `tfsdk:"ingress_ip" json:"ingressIP,omitempty"`
					} `tfsdk:"equinix_metal" json:"equinixMetal,omitempty"`
					External *struct {
						CloudControllerManager *struct {
							State *string `tfsdk:"state" json:"state,omitempty"`
						} `tfsdk:"cloud_controller_manager" json:"cloudControllerManager,omitempty"`
					} `tfsdk:"external" json:"external,omitempty"`
					Gcp *struct {
						ProjectID *string `tfsdk:"project_id" json:"projectID,omitempty"`
						Region    *string `tfsdk:"region" json:"region,omitempty"`
					} `tfsdk:"gcp" json:"gcp,omitempty"`
					Ibmcloud *struct {
						CisInstanceCRN    *string `tfsdk:"cis_instance_crn" json:"cisInstanceCRN,omitempty"`
						DnsInstanceCRN    *string `tfsdk:"dns_instance_crn" json:"dnsInstanceCRN,omitempty"`
						Location          *string `tfsdk:"location" json:"location,omitempty"`
						ProviderType      *string `tfsdk:"provider_type" json:"providerType,omitempty"`
						ResourceGroupName *string `tfsdk:"resource_group_name" json:"resourceGroupName,omitempty"`
						ServiceEndpoints  *[]struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
							Url  *string `tfsdk:"url" json:"url,omitempty"`
						} `tfsdk:"service_endpoints" json:"serviceEndpoints,omitempty"`
					} `tfsdk:"ibmcloud" json:"ibmcloud,omitempty"`
					Kubevirt *struct {
						ApiServerInternalIP *string `tfsdk:"api_server_internal_ip" json:"apiServerInternalIP,omitempty"`
						IngressIP           *string `tfsdk:"ingress_ip" json:"ingressIP,omitempty"`
					} `tfsdk:"kubevirt" json:"kubevirt,omitempty"`
					Nutanix *struct {
						ApiServerInternalIP  *string   `tfsdk:"api_server_internal_ip" json:"apiServerInternalIP,omitempty"`
						ApiServerInternalIPs *[]string `tfsdk:"api_server_internal_i_ps" json:"apiServerInternalIPs,omitempty"`
						IngressIP            *string   `tfsdk:"ingress_ip" json:"ingressIP,omitempty"`
						IngressIPs           *[]string `tfsdk:"ingress_i_ps" json:"ingressIPs,omitempty"`
					} `tfsdk:"nutanix" json:"nutanix,omitempty"`
					Openstack *struct {
						ApiServerInternalIP  *string   `tfsdk:"api_server_internal_ip" json:"apiServerInternalIP,omitempty"`
						ApiServerInternalIPs *[]string `tfsdk:"api_server_internal_i_ps" json:"apiServerInternalIPs,omitempty"`
						CloudName            *string   `tfsdk:"cloud_name" json:"cloudName,omitempty"`
						IngressIP            *string   `tfsdk:"ingress_ip" json:"ingressIP,omitempty"`
						IngressIPs           *[]string `tfsdk:"ingress_i_ps" json:"ingressIPs,omitempty"`
						LoadBalancer         *struct {
							Type *string `tfsdk:"type" json:"type,omitempty"`
						} `tfsdk:"load_balancer" json:"loadBalancer,omitempty"`
						MachineNetworks *[]string `tfsdk:"machine_networks" json:"machineNetworks,omitempty"`
						NodeDNSIP       *string   `tfsdk:"node_dnsip" json:"nodeDNSIP,omitempty"`
					} `tfsdk:"openstack" json:"openstack,omitempty"`
					Ovirt *struct {
						ApiServerInternalIP  *string   `tfsdk:"api_server_internal_ip" json:"apiServerInternalIP,omitempty"`
						ApiServerInternalIPs *[]string `tfsdk:"api_server_internal_i_ps" json:"apiServerInternalIPs,omitempty"`
						IngressIP            *string   `tfsdk:"ingress_ip" json:"ingressIP,omitempty"`
						IngressIPs           *[]string `tfsdk:"ingress_i_ps" json:"ingressIPs,omitempty"`
						NodeDNSIP            *string   `tfsdk:"node_dnsip" json:"nodeDNSIP,omitempty"`
					} `tfsdk:"ovirt" json:"ovirt,omitempty"`
					Powervs *struct {
						CisInstanceCRN   *string `tfsdk:"cis_instance_crn" json:"cisInstanceCRN,omitempty"`
						DnsInstanceCRN   *string `tfsdk:"dns_instance_crn" json:"dnsInstanceCRN,omitempty"`
						Region           *string `tfsdk:"region" json:"region,omitempty"`
						ResourceGroup    *string `tfsdk:"resource_group" json:"resourceGroup,omitempty"`
						ServiceEndpoints *[]struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
							Url  *string `tfsdk:"url" json:"url,omitempty"`
						} `tfsdk:"service_endpoints" json:"serviceEndpoints,omitempty"`
						Zone *string `tfsdk:"zone" json:"zone,omitempty"`
					} `tfsdk:"powervs" json:"powervs,omitempty"`
					Type    *string `tfsdk:"type" json:"type,omitempty"`
					Vsphere *struct {
						ApiServerInternalIP  *string   `tfsdk:"api_server_internal_ip" json:"apiServerInternalIP,omitempty"`
						ApiServerInternalIPs *[]string `tfsdk:"api_server_internal_i_ps" json:"apiServerInternalIPs,omitempty"`
						IngressIP            *string   `tfsdk:"ingress_ip" json:"ingressIP,omitempty"`
						IngressIPs           *[]string `tfsdk:"ingress_i_ps" json:"ingressIPs,omitempty"`
						MachineNetworks      *[]string `tfsdk:"machine_networks" json:"machineNetworks,omitempty"`
						NodeDNSIP            *string   `tfsdk:"node_dnsip" json:"nodeDNSIP,omitempty"`
					} `tfsdk:"vsphere" json:"vsphere,omitempty"`
				} `tfsdk:"platform_status" json:"platformStatus,omitempty"`
			} `tfsdk:"status" json:"status,omitempty"`
		} `tfsdk:"infra" json:"infra,omitempty"`
		InternalRegistryPullSecret *string `tfsdk:"internal_registry_pull_secret" json:"internalRegistryPullSecret,omitempty"`
		IpFamilies                 *string `tfsdk:"ip_families" json:"ipFamilies,omitempty"`
		KubeAPIServerServingCAData *string `tfsdk:"kube_api_server_serving_ca_data" json:"kubeAPIServerServingCAData,omitempty"`
		Network                    *struct {
			MtuMigration *struct {
				Machine *struct {
					From *int64 `tfsdk:"from" json:"from,omitempty"`
					To   *int64 `tfsdk:"to" json:"to,omitempty"`
				} `tfsdk:"machine" json:"machine,omitempty"`
				Network *struct {
					From *int64 `tfsdk:"from" json:"from,omitempty"`
					To   *int64 `tfsdk:"to" json:"to,omitempty"`
				} `tfsdk:"network" json:"network,omitempty"`
			} `tfsdk:"mtu_migration" json:"mtuMigration,omitempty"`
		} `tfsdk:"network" json:"network,omitempty"`
		NetworkType *string `tfsdk:"network_type" json:"networkType,omitempty"`
		OsImageURL  *string `tfsdk:"os_image_url" json:"osImageURL,omitempty"`
		Platform    *string `tfsdk:"platform" json:"platform,omitempty"`
		Proxy       *struct {
			HttpProxy  *string `tfsdk:"http_proxy" json:"httpProxy,omitempty"`
			HttpsProxy *string `tfsdk:"https_proxy" json:"httpsProxy,omitempty"`
			NoProxy    *string `tfsdk:"no_proxy" json:"noProxy,omitempty"`
		} `tfsdk:"proxy" json:"proxy,omitempty"`
		PullSecret *struct {
			ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
			Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
			Name            *string `tfsdk:"name" json:"name,omitempty"`
			Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
			ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
			Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
		} `tfsdk:"pull_secret" json:"pullSecret,omitempty"`
		ReleaseImage *string `tfsdk:"release_image" json:"releaseImage,omitempty"`
		RootCAData   *string `tfsdk:"root_ca_data" json:"rootCAData,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *MachineconfigurationOpenshiftIoControllerConfigV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_machineconfiguration_openshift_io_controller_config_v1_manifest"
}

func (r *MachineconfigurationOpenshiftIoControllerConfigV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ControllerConfig describes configuration for MachineConfigController. This is currently only used to drive the MachineConfig objects generated by the TemplateController.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		MarkdownDescription: "ControllerConfig describes configuration for MachineConfigController. This is currently only used to drive the MachineConfig objects generated by the TemplateController.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
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
				Description:         "ControllerConfigSpec is the spec for ControllerConfig resource.",
				MarkdownDescription: "ControllerConfigSpec is the spec for ControllerConfig resource.",
				Attributes: map[string]schema.Attribute{
					"additional_trust_bundle": schema.StringAttribute{
						Description:         "additionalTrustBundle is a certificate bundle that will be added to the nodes trusted certificate store.",
						MarkdownDescription: "additionalTrustBundle is a certificate bundle that will be added to the nodes trusted certificate store.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							validators.Base64Validator(),
						},
					},

					"base_os_container_image": schema.StringAttribute{
						Description:         "BaseOSContainerImage is the new-format container image for operating system updates.",
						MarkdownDescription: "BaseOSContainerImage is the new-format container image for operating system updates.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"base_os_extensions_container_image": schema.StringAttribute{
						Description:         "BaseOSExtensionsContainerImage is the matching extensions container for the new-format container",
						MarkdownDescription: "BaseOSExtensionsContainerImage is the matching extensions container for the new-format container",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"cloud_provider_ca_data": schema.StringAttribute{
						Description:         "cloudProvider specifies the cloud provider CA data",
						MarkdownDescription: "cloudProvider specifies the cloud provider CA data",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							validators.Base64Validator(),
						},
					},

					"cloud_provider_config": schema.StringAttribute{
						Description:         "cloudProviderConfig is the configuration for the given cloud provider",
						MarkdownDescription: "cloudProviderConfig is the configuration for the given cloud provider",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"cluster_dnsip": schema.StringAttribute{
						Description:         "clusterDNSIP is the cluster DNS IP address",
						MarkdownDescription: "clusterDNSIP is the cluster DNS IP address",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"dns": schema.SingleNestedAttribute{
						Description:         "dns holds the cluster dns details",
						MarkdownDescription: "dns holds the cluster dns details",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
								MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"metadata": schema.MapAttribute{
								Description:         "metadata is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
								MarkdownDescription: "metadata is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"spec": schema.SingleNestedAttribute{
								Description:         "spec holds user settable values for configuration",
								MarkdownDescription: "spec holds user settable values for configuration",
								Attributes: map[string]schema.Attribute{
									"base_domain": schema.StringAttribute{
										Description:         "baseDomain is the base domain of the cluster. All managed DNS records will be sub-domains of this base.  For example, given the base domain 'openshift.example.com', an API server DNS record may be created for 'cluster-api.openshift.example.com'.  Once set, this field cannot be changed.",
										MarkdownDescription: "baseDomain is the base domain of the cluster. All managed DNS records will be sub-domains of this base.  For example, given the base domain 'openshift.example.com', an API server DNS record may be created for 'cluster-api.openshift.example.com'.  Once set, this field cannot be changed.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"platform": schema.SingleNestedAttribute{
										Description:         "platform holds configuration specific to the underlying infrastructure provider for DNS. When omitted, this means the user has no opinion and the platform is left to choose reasonable defaults. These defaults are subject to change over time.",
										MarkdownDescription: "platform holds configuration specific to the underlying infrastructure provider for DNS. When omitted, this means the user has no opinion and the platform is left to choose reasonable defaults. These defaults are subject to change over time.",
										Attributes: map[string]schema.Attribute{
											"aws": schema.SingleNestedAttribute{
												Description:         "aws contains DNS configuration specific to the Amazon Web Services cloud provider.",
												MarkdownDescription: "aws contains DNS configuration specific to the Amazon Web Services cloud provider.",
												Attributes: map[string]schema.Attribute{
													"private_zone_iam_role": schema.StringAttribute{
														Description:         "privateZoneIAMRole contains the ARN of an IAM role that should be assumed when performing operations on the cluster's private hosted zone specified in the cluster DNS config. When left empty, no role should be assumed.",
														MarkdownDescription: "privateZoneIAMRole contains the ARN of an IAM role that should be assumed when performing operations on the cluster's private hosted zone specified in the cluster DNS config. When left empty, no role should be assumed.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.RegexMatches(regexp.MustCompile(`^arn:(aws|aws-cn|aws-us-gov):iam::[0-9]{12}:role\/.*$`), ""),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"type": schema.StringAttribute{
												Description:         "type is the underlying infrastructure provider for the cluster. Allowed values: '', 'AWS'.  Individual components may not support all platforms, and must handle unrecognized platforms with best-effort defaults.",
												MarkdownDescription: "type is the underlying infrastructure provider for the cluster. Allowed values: '', 'AWS'.  Individual components may not support all platforms, and must handle unrecognized platforms with best-effort defaults.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("", "AWS", "Azure", "BareMetal", "GCP", "Libvirt", "OpenStack", "None", "VSphere", "oVirt", "IBMCloud", "KubeVirt", "EquinixMetal", "PowerVS", "AlibabaCloud", "Nutanix", "External"),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"private_zone": schema.SingleNestedAttribute{
										Description:         "privateZone is the location where all the DNS records that are only available internally to the cluster exist.  If this field is nil, no private records should be created.  Once set, this field cannot be changed.",
										MarkdownDescription: "privateZone is the location where all the DNS records that are only available internally to the cluster exist.  If this field is nil, no private records should be created.  Once set, this field cannot be changed.",
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description:         "id is the identifier that can be used to find the DNS hosted zone.  on AWS zone can be fetched using 'ID' as id in [1] on Azure zone can be fetched using 'ID' as a pre-determined name in [2], on GCP zone can be fetched using 'ID' as a pre-determined name in [3].  [1]: https://docs.aws.amazon.com/cli/latest/reference/route53/get-hosted-zone.html#options [2]: https://docs.microsoft.com/en-us/cli/azure/network/dns/zone?view=azure-cli-latest#az-network-dns-zone-show [3]: https://cloud.google.com/dns/docs/reference/v1/managedZones/get",
												MarkdownDescription: "id is the identifier that can be used to find the DNS hosted zone.  on AWS zone can be fetched using 'ID' as id in [1] on Azure zone can be fetched using 'ID' as a pre-determined name in [2], on GCP zone can be fetched using 'ID' as a pre-determined name in [3].  [1]: https://docs.aws.amazon.com/cli/latest/reference/route53/get-hosted-zone.html#options [2]: https://docs.microsoft.com/en-us/cli/azure/network/dns/zone?view=azure-cli-latest#az-network-dns-zone-show [3]: https://cloud.google.com/dns/docs/reference/v1/managedZones/get",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"tags": schema.MapAttribute{
												Description:         "tags can be used to query the DNS hosted zone.  on AWS, resourcegroupstaggingapi [1] can be used to fetch a zone using 'Tags' as tag-filters,  [1]: https://docs.aws.amazon.com/cli/latest/reference/resourcegroupstaggingapi/get-resources.html#options",
												MarkdownDescription: "tags can be used to query the DNS hosted zone.  on AWS, resourcegroupstaggingapi [1] can be used to fetch a zone using 'Tags' as tag-filters,  [1]: https://docs.aws.amazon.com/cli/latest/reference/resourcegroupstaggingapi/get-resources.html#options",
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

									"public_zone": schema.SingleNestedAttribute{
										Description:         "publicZone is the location where all the DNS records that are publicly accessible to the internet exist.  If this field is nil, no public records should be created.  Once set, this field cannot be changed.",
										MarkdownDescription: "publicZone is the location where all the DNS records that are publicly accessible to the internet exist.  If this field is nil, no public records should be created.  Once set, this field cannot be changed.",
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description:         "id is the identifier that can be used to find the DNS hosted zone.  on AWS zone can be fetched using 'ID' as id in [1] on Azure zone can be fetched using 'ID' as a pre-determined name in [2], on GCP zone can be fetched using 'ID' as a pre-determined name in [3].  [1]: https://docs.aws.amazon.com/cli/latest/reference/route53/get-hosted-zone.html#options [2]: https://docs.microsoft.com/en-us/cli/azure/network/dns/zone?view=azure-cli-latest#az-network-dns-zone-show [3]: https://cloud.google.com/dns/docs/reference/v1/managedZones/get",
												MarkdownDescription: "id is the identifier that can be used to find the DNS hosted zone.  on AWS zone can be fetched using 'ID' as id in [1] on Azure zone can be fetched using 'ID' as a pre-determined name in [2], on GCP zone can be fetched using 'ID' as a pre-determined name in [3].  [1]: https://docs.aws.amazon.com/cli/latest/reference/route53/get-hosted-zone.html#options [2]: https://docs.microsoft.com/en-us/cli/azure/network/dns/zone?view=azure-cli-latest#az-network-dns-zone-show [3]: https://cloud.google.com/dns/docs/reference/v1/managedZones/get",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"tags": schema.MapAttribute{
												Description:         "tags can be used to query the DNS hosted zone.  on AWS, resourcegroupstaggingapi [1] can be used to fetch a zone using 'Tags' as tag-filters,  [1]: https://docs.aws.amazon.com/cli/latest/reference/resourcegroupstaggingapi/get-resources.html#options",
												MarkdownDescription: "tags can be used to query the DNS hosted zone.  on AWS, resourcegroupstaggingapi [1] can be used to fetch a zone using 'Tags' as tag-filters,  [1]: https://docs.aws.amazon.com/cli/latest/reference/resourcegroupstaggingapi/get-resources.html#options",
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
								Required: true,
								Optional: false,
								Computed: false,
							},

							"status": schema.MapAttribute{
								Description:         "status holds observed values from the cluster. They may not be overridden.",
								MarkdownDescription: "status holds observed values from the cluster. They may not be overridden.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"etcd_discovery_domain": schema.StringAttribute{
						Description:         "etcdDiscoveryDomain is deprecated, use Infra.Status.EtcdDiscoveryDomain instead",
						MarkdownDescription: "etcdDiscoveryDomain is deprecated, use Infra.Status.EtcdDiscoveryDomain instead",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image_registry_bundle_data": schema.ListNestedAttribute{
						Description:         "imageRegistryBundleData is the ImageRegistryData",
						MarkdownDescription: "imageRegistryBundleData is the ImageRegistryData",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"data": schema.StringAttribute{
									Description:         "data holds the contents of the bundle that will be written to the file location",
									MarkdownDescription: "data holds the contents of the bundle that will be written to the file location",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										validators.Base64Validator(),
									},
								},

								"file": schema.StringAttribute{
									Description:         "file holds the name of the file where the bundle will be written to disk",
									MarkdownDescription: "file holds the name of the file where the bundle will be written to disk",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"image_registry_bundle_user_data": schema.ListNestedAttribute{
						Description:         "imageRegistryBundleUserData is Image Registry Data provided by the user",
						MarkdownDescription: "imageRegistryBundleUserData is Image Registry Data provided by the user",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"data": schema.StringAttribute{
									Description:         "data holds the contents of the bundle that will be written to the file location",
									MarkdownDescription: "data holds the contents of the bundle that will be written to the file location",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										validators.Base64Validator(),
									},
								},

								"file": schema.StringAttribute{
									Description:         "file holds the name of the file where the bundle will be written to disk",
									MarkdownDescription: "file holds the name of the file where the bundle will be written to disk",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"images": schema.MapAttribute{
						Description:         "images is map of images that are used by the controller to render templates under ./templates/",
						MarkdownDescription: "images is map of images that are used by the controller to render templates under ./templates/",
						ElementType:         types.StringType,
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"infra": schema.SingleNestedAttribute{
						Description:         "infra holds the infrastructure details",
						MarkdownDescription: "infra holds the infrastructure details",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
								MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"metadata": schema.MapAttribute{
								Description:         "metadata is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
								MarkdownDescription: "metadata is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
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

							"status": schema.SingleNestedAttribute{
								Description:         "status holds observed values from the cluster. They may not be overridden.",
								MarkdownDescription: "status holds observed values from the cluster. They may not be overridden.",
								Attributes: map[string]schema.Attribute{
									"api_server_internal_uri": schema.StringAttribute{
										Description:         "apiServerInternalURL is a valid URI with scheme 'https', address and optionally a port (defaulting to 443).  apiServerInternalURL can be used by components like kubelets, to contact the Kubernetes API server using the infrastructure provider rather than Kubernetes networking.",
										MarkdownDescription: "apiServerInternalURL is a valid URI with scheme 'https', address and optionally a port (defaulting to 443).  apiServerInternalURL can be used by components like kubelets, to contact the Kubernetes API server using the infrastructure provider rather than Kubernetes networking.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"api_server_url": schema.StringAttribute{
										Description:         "apiServerURL is a valid URI with scheme 'https', address and optionally a port (defaulting to 443).  apiServerURL can be used by components like the web console to tell users where to find the Kubernetes API.",
										MarkdownDescription: "apiServerURL is a valid URI with scheme 'https', address and optionally a port (defaulting to 443).  apiServerURL can be used by components like the web console to tell users where to find the Kubernetes API.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"control_plane_topology": schema.StringAttribute{
										Description:         "controlPlaneTopology expresses the expectations for operands that normally run on control nodes. The default is 'HighlyAvailable', which represents the behavior operators have in a 'normal' cluster. The 'SingleReplica' mode will be used in single-node deployments and the operators should not configure the operand for highly-available operation The 'External' mode indicates that the control plane is hosted externally to the cluster and that its components are not visible within the cluster.",
										MarkdownDescription: "controlPlaneTopology expresses the expectations for operands that normally run on control nodes. The default is 'HighlyAvailable', which represents the behavior operators have in a 'normal' cluster. The 'SingleReplica' mode will be used in single-node deployments and the operators should not configure the operand for highly-available operation The 'External' mode indicates that the control plane is hosted externally to the cluster and that its components are not visible within the cluster.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("HighlyAvailable", "SingleReplica", "External"),
										},
									},

									"cpu_partitioning": schema.StringAttribute{
										Description:         "cpuPartitioning expresses if CPU partitioning is a currently enabled feature in the cluster. CPU Partitioning means that this cluster can support partitioning workloads to specific CPU Sets. Valid values are 'None' and 'AllNodes'. When omitted, the default value is 'None'. The default value of 'None' indicates that no nodes will be setup with CPU partitioning. The 'AllNodes' value indicates that all nodes have been setup with CPU partitioning, and can then be further configured via the PerformanceProfile API.",
										MarkdownDescription: "cpuPartitioning expresses if CPU partitioning is a currently enabled feature in the cluster. CPU Partitioning means that this cluster can support partitioning workloads to specific CPU Sets. Valid values are 'None' and 'AllNodes'. When omitted, the default value is 'None'. The default value of 'None' indicates that no nodes will be setup with CPU partitioning. The 'AllNodes' value indicates that all nodes have been setup with CPU partitioning, and can then be further configured via the PerformanceProfile API.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("None", "AllNodes"),
										},
									},

									"etcd_discovery_domain": schema.StringAttribute{
										Description:         "etcdDiscoveryDomain is the domain used to fetch the SRV records for discovering etcd servers and clients. For more info: https://github.com/etcd-io/etcd/blob/329be66e8b3f9e2e6af83c123ff89297e49ebd15/Documentation/op-guide/clustering.md#dns-discovery deprecated: as of 4.7, this field is no longer set or honored.  It will be removed in a future release.",
										MarkdownDescription: "etcdDiscoveryDomain is the domain used to fetch the SRV records for discovering etcd servers and clients. For more info: https://github.com/etcd-io/etcd/blob/329be66e8b3f9e2e6af83c123ff89297e49ebd15/Documentation/op-guide/clustering.md#dns-discovery deprecated: as of 4.7, this field is no longer set or honored.  It will be removed in a future release.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"infrastructure_name": schema.StringAttribute{
										Description:         "infrastructureName uniquely identifies a cluster with a human friendly name. Once set it should not be changed. Must be of max length 27 and must have only alphanumeric or hyphen characters.",
										MarkdownDescription: "infrastructureName uniquely identifies a cluster with a human friendly name. Once set it should not be changed. Must be of max length 27 and must have only alphanumeric or hyphen characters.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"infrastructure_topology": schema.StringAttribute{
										Description:         "infrastructureTopology expresses the expectations for infrastructure services that do not run on control plane nodes, usually indicated by a node selector for a 'role' value other than 'master'. The default is 'HighlyAvailable', which represents the behavior operators have in a 'normal' cluster. The 'SingleReplica' mode will be used in single-node deployments and the operators should not configure the operand for highly-available operation NOTE: External topology mode is not applicable for this field.",
										MarkdownDescription: "infrastructureTopology expresses the expectations for infrastructure services that do not run on control plane nodes, usually indicated by a node selector for a 'role' value other than 'master'. The default is 'HighlyAvailable', which represents the behavior operators have in a 'normal' cluster. The 'SingleReplica' mode will be used in single-node deployments and the operators should not configure the operand for highly-available operation NOTE: External topology mode is not applicable for this field.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("HighlyAvailable", "SingleReplica"),
										},
									},

									"platform": schema.StringAttribute{
										Description:         "platform is the underlying infrastructure provider for the cluster.  Deprecated: Use platformStatus.type instead.",
										MarkdownDescription: "platform is the underlying infrastructure provider for the cluster.  Deprecated: Use platformStatus.type instead.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("", "AWS", "Azure", "BareMetal", "GCP", "Libvirt", "OpenStack", "None", "VSphere", "oVirt", "IBMCloud", "KubeVirt", "EquinixMetal", "PowerVS", "AlibabaCloud", "Nutanix", "External"),
										},
									},

									"platform_status": schema.SingleNestedAttribute{
										Description:         "platformStatus holds status information specific to the underlying infrastructure provider.",
										MarkdownDescription: "platformStatus holds status information specific to the underlying infrastructure provider.",
										Attributes: map[string]schema.Attribute{
											"alibaba_cloud": schema.SingleNestedAttribute{
												Description:         "AlibabaCloud contains settings specific to the Alibaba Cloud infrastructure provider.",
												MarkdownDescription: "AlibabaCloud contains settings specific to the Alibaba Cloud infrastructure provider.",
												Attributes: map[string]schema.Attribute{
													"region": schema.StringAttribute{
														Description:         "region specifies the region for Alibaba Cloud resources created for the cluster.",
														MarkdownDescription: "region specifies the region for Alibaba Cloud resources created for the cluster.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9A-Za-z-]+$`), ""),
														},
													},

													"resource_group_id": schema.StringAttribute{
														Description:         "resourceGroupID is the ID of the resource group for the cluster.",
														MarkdownDescription: "resourceGroupID is the ID of the resource group for the cluster.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.RegexMatches(regexp.MustCompile(`^(rg-[0-9A-Za-z]+)?$`), ""),
														},
													},

													"resource_tags": schema.ListNestedAttribute{
														Description:         "resourceTags is a list of additional tags to apply to Alibaba Cloud resources created for the cluster.",
														MarkdownDescription: "resourceTags is a list of additional tags to apply to Alibaba Cloud resources created for the cluster.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "key is the key of the tag.",
																	MarkdownDescription: "key is the key of the tag.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtLeast(1),
																		stringvalidator.LengthAtMost(128),
																	},
																},

																"value": schema.StringAttribute{
																	Description:         "value is the value of the tag.",
																	MarkdownDescription: "value is the value of the tag.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtLeast(1),
																		stringvalidator.LengthAtMost(128),
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

											"aws": schema.SingleNestedAttribute{
												Description:         "AWS contains settings specific to the Amazon Web Services infrastructure provider.",
												MarkdownDescription: "AWS contains settings specific to the Amazon Web Services infrastructure provider.",
												Attributes: map[string]schema.Attribute{
													"region": schema.StringAttribute{
														Description:         "region holds the default AWS region for new AWS resources created by the cluster.",
														MarkdownDescription: "region holds the default AWS region for new AWS resources created by the cluster.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"resource_tags": schema.ListNestedAttribute{
														Description:         "resourceTags is a list of additional tags to apply to AWS resources created for the cluster. See https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html for information on tagging AWS resources. AWS supports a maximum of 50 tags per resource. OpenShift reserves 25 tags for its use, leaving 25 tags available for the user.",
														MarkdownDescription: "resourceTags is a list of additional tags to apply to AWS resources created for the cluster. See https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html for information on tagging AWS resources. AWS supports a maximum of 50 tags per resource. OpenShift reserves 25 tags for its use, leaving 25 tags available for the user.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "key is the key of the tag",
																	MarkdownDescription: "key is the key of the tag",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtLeast(1),
																		stringvalidator.LengthAtMost(128),
																		stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9A-Za-z_.:/=+-@]+$`), ""),
																	},
																},

																"value": schema.StringAttribute{
																	Description:         "value is the value of the tag. Some AWS service do not support empty values. Since tags are added to resources in many services, the length of the tag value must meet the requirements of all services.",
																	MarkdownDescription: "value is the value of the tag. Some AWS service do not support empty values. Since tags are added to resources in many services, the length of the tag value must meet the requirements of all services.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtLeast(1),
																		stringvalidator.LengthAtMost(256),
																		stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9A-Za-z_.:/=+-@]+$`), ""),
																	},
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"service_endpoints": schema.ListNestedAttribute{
														Description:         "ServiceEndpoints list contains custom endpoints which will override default service endpoint of AWS Services. There must be only one ServiceEndpoint for a service.",
														MarkdownDescription: "ServiceEndpoints list contains custom endpoints which will override default service endpoint of AWS Services. There must be only one ServiceEndpoint for a service.",
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

											"azure": schema.SingleNestedAttribute{
												Description:         "Azure contains settings specific to the Azure infrastructure provider.",
												MarkdownDescription: "Azure contains settings specific to the Azure infrastructure provider.",
												Attributes: map[string]schema.Attribute{
													"arm_endpoint": schema.StringAttribute{
														Description:         "armEndpoint specifies a URL to use for resource management in non-soverign clouds such as Azure Stack.",
														MarkdownDescription: "armEndpoint specifies a URL to use for resource management in non-soverign clouds such as Azure Stack.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"cloud_name": schema.StringAttribute{
														Description:         "cloudName is the name of the Azure cloud environment which can be used to configure the Azure SDK with the appropriate Azure API endpoints. If empty, the value is equal to 'AzurePublicCloud'.",
														MarkdownDescription: "cloudName is the name of the Azure cloud environment which can be used to configure the Azure SDK with the appropriate Azure API endpoints. If empty, the value is equal to 'AzurePublicCloud'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("", "AzurePublicCloud", "AzureUSGovernmentCloud", "AzureChinaCloud", "AzureGermanCloud", "AzureStackCloud"),
														},
													},

													"network_resource_group_name": schema.StringAttribute{
														Description:         "networkResourceGroupName is the Resource Group for network resources like the Virtual Network and Subnets used by the cluster. If empty, the value is same as ResourceGroupName.",
														MarkdownDescription: "networkResourceGroupName is the Resource Group for network resources like the Virtual Network and Subnets used by the cluster. If empty, the value is same as ResourceGroupName.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"resource_group_name": schema.StringAttribute{
														Description:         "resourceGroupName is the Resource Group for new Azure resources created for the cluster.",
														MarkdownDescription: "resourceGroupName is the Resource Group for new Azure resources created for the cluster.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"resource_tags": schema.ListNestedAttribute{
														Description:         "resourceTags is a list of additional tags to apply to Azure resources created for the cluster. See https://docs.microsoft.com/en-us/rest/api/resources/tags for information on tagging Azure resources. Due to limitations on Automation, Content Delivery Network, DNS Azure resources, a maximum of 15 tags may be applied. OpenShift reserves 5 tags for internal use, allowing 10 tags for user configuration.",
														MarkdownDescription: "resourceTags is a list of additional tags to apply to Azure resources created for the cluster. See https://docs.microsoft.com/en-us/rest/api/resources/tags for information on tagging Azure resources. Due to limitations on Automation, Content Delivery Network, DNS Azure resources, a maximum of 15 tags may be applied. OpenShift reserves 5 tags for internal use, allowing 10 tags for user configuration.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "key is the key part of the tag. A tag key can have a maximum of 128 characters and cannot be empty. Key must begin with a letter, end with a letter, number or underscore, and must contain only alphanumeric characters and the following special characters '_ . -'.",
																	MarkdownDescription: "key is the key part of the tag. A tag key can have a maximum of 128 characters and cannot be empty. Key must begin with a letter, end with a letter, number or underscore, and must contain only alphanumeric characters and the following special characters '_ . -'.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtLeast(1),
																		stringvalidator.LengthAtMost(128),
																		stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z]([0-9A-Za-z_.-]*[0-9A-Za-z_])?$`), ""),
																	},
																},

																"value": schema.StringAttribute{
																	Description:         "value is the value part of the tag. A tag value can have a maximum of 256 characters and cannot be empty. Value must contain only alphanumeric characters and the following special characters '_ + , - . / : ; < = > ? @'.",
																	MarkdownDescription: "value is the value part of the tag. A tag value can have a maximum of 256 characters and cannot be empty. Value must contain only alphanumeric characters and the following special characters '_ + , - . / : ; < = > ? @'.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtLeast(1),
																		stringvalidator.LengthAtMost(256),
																		stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9A-Za-z_.=+-@]+$`), ""),
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

											"baremetal": schema.SingleNestedAttribute{
												Description:         "BareMetal contains settings specific to the BareMetal platform.",
												MarkdownDescription: "BareMetal contains settings specific to the BareMetal platform.",
												Attributes: map[string]schema.Attribute{
													"api_server_internal_ip": schema.StringAttribute{
														Description:         "apiServerInternalIP is an IP address to contact the Kubernetes API server that can be used by components inside the cluster, like kubelets using the infrastructure rather than Kubernetes networking. It is the IP that the Infrastructure.status.apiServerInternalURI points to. It is the IP for a self-hosted load balancer in front of the API servers.  Deprecated: Use APIServerInternalIPs instead.",
														MarkdownDescription: "apiServerInternalIP is an IP address to contact the Kubernetes API server that can be used by components inside the cluster, like kubelets using the infrastructure rather than Kubernetes networking. It is the IP that the Infrastructure.status.apiServerInternalURI points to. It is the IP for a self-hosted load balancer in front of the API servers.  Deprecated: Use APIServerInternalIPs instead.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"api_server_internal_i_ps": schema.ListAttribute{
														Description:         "apiServerInternalIPs are the IP addresses to contact the Kubernetes API server that can be used by components inside the cluster, like kubelets using the infrastructure rather than Kubernetes networking. These are the IPs for a self-hosted load balancer in front of the API servers. In dual stack clusters this list contains two IPs otherwise only one.",
														MarkdownDescription: "apiServerInternalIPs are the IP addresses to contact the Kubernetes API server that can be used by components inside the cluster, like kubelets using the infrastructure rather than Kubernetes networking. These are the IPs for a self-hosted load balancer in front of the API servers. In dual stack clusters this list contains two IPs otherwise only one.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ingress_ip": schema.StringAttribute{
														Description:         "ingressIP is an external IP which routes to the default ingress controller. The IP is a suitable target of a wildcard DNS record used to resolve default route host names.  Deprecated: Use IngressIPs instead.",
														MarkdownDescription: "ingressIP is an external IP which routes to the default ingress controller. The IP is a suitable target of a wildcard DNS record used to resolve default route host names.  Deprecated: Use IngressIPs instead.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ingress_i_ps": schema.ListAttribute{
														Description:         "ingressIPs are the external IPs which route to the default ingress controller. The IPs are suitable targets of a wildcard DNS record used to resolve default route host names. In dual stack clusters this list contains two IPs otherwise only one.",
														MarkdownDescription: "ingressIPs are the external IPs which route to the default ingress controller. The IPs are suitable targets of a wildcard DNS record used to resolve default route host names. In dual stack clusters this list contains two IPs otherwise only one.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"machine_networks": schema.ListAttribute{
														Description:         "machineNetworks are IP networks used to connect all the OpenShift cluster nodes.",
														MarkdownDescription: "machineNetworks are IP networks used to connect all the OpenShift cluster nodes.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"node_dnsip": schema.StringAttribute{
														Description:         "nodeDNSIP is the IP address for the internal DNS used by the nodes. Unlike the one managed by the DNS operator, 'NodeDNSIP' provides name resolution for the nodes themselves. There is no DNS-as-a-service for BareMetal deployments. In order to minimize necessary changes to the datacenter DNS, a DNS service is hosted as a static pod to serve those hostnames to the nodes in the cluster.",
														MarkdownDescription: "nodeDNSIP is the IP address for the internal DNS used by the nodes. Unlike the one managed by the DNS operator, 'NodeDNSIP' provides name resolution for the nodes themselves. There is no DNS-as-a-service for BareMetal deployments. In order to minimize necessary changes to the datacenter DNS, a DNS service is hosted as a static pod to serve those hostnames to the nodes in the cluster.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"equinix_metal": schema.SingleNestedAttribute{
												Description:         "EquinixMetal contains settings specific to the Equinix Metal infrastructure provider.",
												MarkdownDescription: "EquinixMetal contains settings specific to the Equinix Metal infrastructure provider.",
												Attributes: map[string]schema.Attribute{
													"api_server_internal_ip": schema.StringAttribute{
														Description:         "apiServerInternalIP is an IP address to contact the Kubernetes API server that can be used by components inside the cluster, like kubelets using the infrastructure rather than Kubernetes networking. It is the IP that the Infrastructure.status.apiServerInternalURI points to. It is the IP for a self-hosted load balancer in front of the API servers.",
														MarkdownDescription: "apiServerInternalIP is an IP address to contact the Kubernetes API server that can be used by components inside the cluster, like kubelets using the infrastructure rather than Kubernetes networking. It is the IP that the Infrastructure.status.apiServerInternalURI points to. It is the IP for a self-hosted load balancer in front of the API servers.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ingress_ip": schema.StringAttribute{
														Description:         "ingressIP is an external IP which routes to the default ingress controller. The IP is a suitable target of a wildcard DNS record used to resolve default route host names.",
														MarkdownDescription: "ingressIP is an external IP which routes to the default ingress controller. The IP is a suitable target of a wildcard DNS record used to resolve default route host names.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"external": schema.SingleNestedAttribute{
												Description:         "External contains settings specific to the generic External infrastructure provider.",
												MarkdownDescription: "External contains settings specific to the generic External infrastructure provider.",
												Attributes: map[string]schema.Attribute{
													"cloud_controller_manager": schema.SingleNestedAttribute{
														Description:         "cloudControllerManager contains settings specific to the external Cloud Controller Manager (a.k.a. CCM or CPI). When omitted, new nodes will be not tainted and no extra initialization from the cloud controller manager is expected.",
														MarkdownDescription: "cloudControllerManager contains settings specific to the external Cloud Controller Manager (a.k.a. CCM or CPI). When omitted, new nodes will be not tainted and no extra initialization from the cloud controller manager is expected.",
														Attributes: map[string]schema.Attribute{
															"state": schema.StringAttribute{
																Description:         "state determines whether or not an external Cloud Controller Manager is expected to be installed within the cluster. https://kubernetes.io/docs/tasks/administer-cluster/running-cloud-controller/#running-cloud-controller-manager  Valid values are 'External', 'None' and omitted. When set to 'External', new nodes will be tainted as uninitialized when created, preventing them from running workloads until they are initialized by the cloud controller manager. When omitted or set to 'None', new nodes will be not tainted and no extra initialization from the cloud controller manager is expected.",
																MarkdownDescription: "state determines whether or not an external Cloud Controller Manager is expected to be installed within the cluster. https://kubernetes.io/docs/tasks/administer-cluster/running-cloud-controller/#running-cloud-controller-manager  Valid values are 'External', 'None' and omitted. When set to 'External', new nodes will be tainted as uninitialized when created, preventing them from running workloads until they are initialized by the cloud controller manager. When omitted or set to 'None', new nodes will be not tainted and no extra initialization from the cloud controller manager is expected.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("", "External", "None"),
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

											"gcp": schema.SingleNestedAttribute{
												Description:         "GCP contains settings specific to the Google Cloud Platform infrastructure provider.",
												MarkdownDescription: "GCP contains settings specific to the Google Cloud Platform infrastructure provider.",
												Attributes: map[string]schema.Attribute{
													"project_id": schema.StringAttribute{
														Description:         "resourceGroupName is the Project ID for new GCP resources created for the cluster.",
														MarkdownDescription: "resourceGroupName is the Project ID for new GCP resources created for the cluster.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"region": schema.StringAttribute{
														Description:         "region holds the region for new GCP resources created for the cluster.",
														MarkdownDescription: "region holds the region for new GCP resources created for the cluster.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"ibmcloud": schema.SingleNestedAttribute{
												Description:         "IBMCloud contains settings specific to the IBMCloud infrastructure provider.",
												MarkdownDescription: "IBMCloud contains settings specific to the IBMCloud infrastructure provider.",
												Attributes: map[string]schema.Attribute{
													"cis_instance_crn": schema.StringAttribute{
														Description:         "CISInstanceCRN is the CRN of the Cloud Internet Services instance managing the DNS zone for the cluster's base domain",
														MarkdownDescription: "CISInstanceCRN is the CRN of the Cloud Internet Services instance managing the DNS zone for the cluster's base domain",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"dns_instance_crn": schema.StringAttribute{
														Description:         "DNSInstanceCRN is the CRN of the DNS Services instance managing the DNS zone for the cluster's base domain",
														MarkdownDescription: "DNSInstanceCRN is the CRN of the DNS Services instance managing the DNS zone for the cluster's base domain",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"location": schema.StringAttribute{
														Description:         "Location is where the cluster has been deployed",
														MarkdownDescription: "Location is where the cluster has been deployed",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"provider_type": schema.StringAttribute{
														Description:         "ProviderType indicates the type of cluster that was created",
														MarkdownDescription: "ProviderType indicates the type of cluster that was created",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"resource_group_name": schema.StringAttribute{
														Description:         "ResourceGroupName is the Resource Group for new IBMCloud resources created for the cluster.",
														MarkdownDescription: "ResourceGroupName is the Resource Group for new IBMCloud resources created for the cluster.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"service_endpoints": schema.ListNestedAttribute{
														Description:         "serviceEndpoints is a list of custom endpoints which will override the default service endpoints of an IBM Cloud service. These endpoints are consumed by components within the cluster to reach the respective IBM Cloud Services.",
														MarkdownDescription: "serviceEndpoints is a list of custom endpoints which will override the default service endpoints of an IBM Cloud service. These endpoints are consumed by components within the cluster to reach the respective IBM Cloud Services.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "name is the name of the IBM Cloud service. Possible values are: CIS, COS, DNSServices, GlobalSearch, GlobalTagging, HyperProtect, IAM, KeyProtect, ResourceController, ResourceManager, or VPC. For example, the IBM Cloud Private IAM service could be configured with the service 'name' of 'IAM' and 'url' of 'https://private.iam.cloud.ibm.com' Whereas the IBM Cloud Private VPC service for US South (Dallas) could be configured with the service 'name' of 'VPC' and 'url' of 'https://us.south.private.iaas.cloud.ibm.com'",
																	MarkdownDescription: "name is the name of the IBM Cloud service. Possible values are: CIS, COS, DNSServices, GlobalSearch, GlobalTagging, HyperProtect, IAM, KeyProtect, ResourceController, ResourceManager, or VPC. For example, the IBM Cloud Private IAM service could be configured with the service 'name' of 'IAM' and 'url' of 'https://private.iam.cloud.ibm.com' Whereas the IBM Cloud Private VPC service for US South (Dallas) could be configured with the service 'name' of 'VPC' and 'url' of 'https://us.south.private.iaas.cloud.ibm.com'",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("CIS", "COS", "DNSServices", "GlobalSearch", "GlobalTagging", "HyperProtect", "IAM", "KeyProtect", "ResourceController", "ResourceManager", "VPC"),
																	},
																},

																"url": schema.StringAttribute{
																	Description:         "url is fully qualified URI with scheme https, that overrides the default generated endpoint for a client. This must be provided and cannot be empty.",
																	MarkdownDescription: "url is fully qualified URI with scheme https, that overrides the default generated endpoint for a client. This must be provided and cannot be empty.",
																	Required:            true,
																	Optional:            false,
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

											"kubevirt": schema.SingleNestedAttribute{
												Description:         "Kubevirt contains settings specific to the kubevirt infrastructure provider.",
												MarkdownDescription: "Kubevirt contains settings specific to the kubevirt infrastructure provider.",
												Attributes: map[string]schema.Attribute{
													"api_server_internal_ip": schema.StringAttribute{
														Description:         "apiServerInternalIP is an IP address to contact the Kubernetes API server that can be used by components inside the cluster, like kubelets using the infrastructure rather than Kubernetes networking. It is the IP that the Infrastructure.status.apiServerInternalURI points to. It is the IP for a self-hosted load balancer in front of the API servers.",
														MarkdownDescription: "apiServerInternalIP is an IP address to contact the Kubernetes API server that can be used by components inside the cluster, like kubelets using the infrastructure rather than Kubernetes networking. It is the IP that the Infrastructure.status.apiServerInternalURI points to. It is the IP for a self-hosted load balancer in front of the API servers.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ingress_ip": schema.StringAttribute{
														Description:         "ingressIP is an external IP which routes to the default ingress controller. The IP is a suitable target of a wildcard DNS record used to resolve default route host names.",
														MarkdownDescription: "ingressIP is an external IP which routes to the default ingress controller. The IP is a suitable target of a wildcard DNS record used to resolve default route host names.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"nutanix": schema.SingleNestedAttribute{
												Description:         "Nutanix contains settings specific to the Nutanix infrastructure provider.",
												MarkdownDescription: "Nutanix contains settings specific to the Nutanix infrastructure provider.",
												Attributes: map[string]schema.Attribute{
													"api_server_internal_ip": schema.StringAttribute{
														Description:         "apiServerInternalIP is an IP address to contact the Kubernetes API server that can be used by components inside the cluster, like kubelets using the infrastructure rather than Kubernetes networking. It is the IP that the Infrastructure.status.apiServerInternalURI points to. It is the IP for a self-hosted load balancer in front of the API servers.  Deprecated: Use APIServerInternalIPs instead.",
														MarkdownDescription: "apiServerInternalIP is an IP address to contact the Kubernetes API server that can be used by components inside the cluster, like kubelets using the infrastructure rather than Kubernetes networking. It is the IP that the Infrastructure.status.apiServerInternalURI points to. It is the IP for a self-hosted load balancer in front of the API servers.  Deprecated: Use APIServerInternalIPs instead.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"api_server_internal_i_ps": schema.ListAttribute{
														Description:         "apiServerInternalIPs are the IP addresses to contact the Kubernetes API server that can be used by components inside the cluster, like kubelets using the infrastructure rather than Kubernetes networking. These are the IPs for a self-hosted load balancer in front of the API servers. In dual stack clusters this list contains two IPs otherwise only one.",
														MarkdownDescription: "apiServerInternalIPs are the IP addresses to contact the Kubernetes API server that can be used by components inside the cluster, like kubelets using the infrastructure rather than Kubernetes networking. These are the IPs for a self-hosted load balancer in front of the API servers. In dual stack clusters this list contains two IPs otherwise only one.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ingress_ip": schema.StringAttribute{
														Description:         "ingressIP is an external IP which routes to the default ingress controller. The IP is a suitable target of a wildcard DNS record used to resolve default route host names.  Deprecated: Use IngressIPs instead.",
														MarkdownDescription: "ingressIP is an external IP which routes to the default ingress controller. The IP is a suitable target of a wildcard DNS record used to resolve default route host names.  Deprecated: Use IngressIPs instead.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ingress_i_ps": schema.ListAttribute{
														Description:         "ingressIPs are the external IPs which route to the default ingress controller. The IPs are suitable targets of a wildcard DNS record used to resolve default route host names. In dual stack clusters this list contains two IPs otherwise only one.",
														MarkdownDescription: "ingressIPs are the external IPs which route to the default ingress controller. The IPs are suitable targets of a wildcard DNS record used to resolve default route host names. In dual stack clusters this list contains two IPs otherwise only one.",
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

											"openstack": schema.SingleNestedAttribute{
												Description:         "OpenStack contains settings specific to the OpenStack infrastructure provider.",
												MarkdownDescription: "OpenStack contains settings specific to the OpenStack infrastructure provider.",
												Attributes: map[string]schema.Attribute{
													"api_server_internal_ip": schema.StringAttribute{
														Description:         "apiServerInternalIP is an IP address to contact the Kubernetes API server that can be used by components inside the cluster, like kubelets using the infrastructure rather than Kubernetes networking. It is the IP that the Infrastructure.status.apiServerInternalURI points to. It is the IP for a self-hosted load balancer in front of the API servers.  Deprecated: Use APIServerInternalIPs instead.",
														MarkdownDescription: "apiServerInternalIP is an IP address to contact the Kubernetes API server that can be used by components inside the cluster, like kubelets using the infrastructure rather than Kubernetes networking. It is the IP that the Infrastructure.status.apiServerInternalURI points to. It is the IP for a self-hosted load balancer in front of the API servers.  Deprecated: Use APIServerInternalIPs instead.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"api_server_internal_i_ps": schema.ListAttribute{
														Description:         "apiServerInternalIPs are the IP addresses to contact the Kubernetes API server that can be used by components inside the cluster, like kubelets using the infrastructure rather than Kubernetes networking. These are the IPs for a self-hosted load balancer in front of the API servers. In dual stack clusters this list contains two IPs otherwise only one.",
														MarkdownDescription: "apiServerInternalIPs are the IP addresses to contact the Kubernetes API server that can be used by components inside the cluster, like kubelets using the infrastructure rather than Kubernetes networking. These are the IPs for a self-hosted load balancer in front of the API servers. In dual stack clusters this list contains two IPs otherwise only one.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"cloud_name": schema.StringAttribute{
														Description:         "cloudName is the name of the desired OpenStack cloud in the client configuration file ('clouds.yaml').",
														MarkdownDescription: "cloudName is the name of the desired OpenStack cloud in the client configuration file ('clouds.yaml').",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ingress_ip": schema.StringAttribute{
														Description:         "ingressIP is an external IP which routes to the default ingress controller. The IP is a suitable target of a wildcard DNS record used to resolve default route host names.  Deprecated: Use IngressIPs instead.",
														MarkdownDescription: "ingressIP is an external IP which routes to the default ingress controller. The IP is a suitable target of a wildcard DNS record used to resolve default route host names.  Deprecated: Use IngressIPs instead.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ingress_i_ps": schema.ListAttribute{
														Description:         "ingressIPs are the external IPs which route to the default ingress controller. The IPs are suitable targets of a wildcard DNS record used to resolve default route host names. In dual stack clusters this list contains two IPs otherwise only one.",
														MarkdownDescription: "ingressIPs are the external IPs which route to the default ingress controller. The IPs are suitable targets of a wildcard DNS record used to resolve default route host names. In dual stack clusters this list contains two IPs otherwise only one.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"load_balancer": schema.SingleNestedAttribute{
														Description:         "loadBalancer defines how the load balancer used by the cluster is configured.",
														MarkdownDescription: "loadBalancer defines how the load balancer used by the cluster is configured.",
														Attributes: map[string]schema.Attribute{
															"type": schema.StringAttribute{
																Description:         "type defines the type of load balancer used by the cluster on OpenStack platform which can be a user-managed or openshift-managed load balancer that is to be used for the OpenShift API and Ingress endpoints. When set to OpenShiftManagedDefault the static pods in charge of API and Ingress traffic load-balancing defined in the machine config operator will be deployed. When set to UserManaged these static pods will not be deployed and it is expected that the load balancer is configured out of band by the deployer. When omitted, this means no opinion and the platform is left to choose a reasonable default. The default value is OpenShiftManagedDefault.",
																MarkdownDescription: "type defines the type of load balancer used by the cluster on OpenStack platform which can be a user-managed or openshift-managed load balancer that is to be used for the OpenShift API and Ingress endpoints. When set to OpenShiftManagedDefault the static pods in charge of API and Ingress traffic load-balancing defined in the machine config operator will be deployed. When set to UserManaged these static pods will not be deployed and it is expected that the load balancer is configured out of band by the deployer. When omitted, this means no opinion and the platform is left to choose a reasonable default. The default value is OpenShiftManagedDefault.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("OpenShiftManagedDefault", "UserManaged"),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"machine_networks": schema.ListAttribute{
														Description:         "machineNetworks are IP networks used to connect all the OpenShift cluster nodes.",
														MarkdownDescription: "machineNetworks are IP networks used to connect all the OpenShift cluster nodes.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"node_dnsip": schema.StringAttribute{
														Description:         "nodeDNSIP is the IP address for the internal DNS used by the nodes. Unlike the one managed by the DNS operator, 'NodeDNSIP' provides name resolution for the nodes themselves. There is no DNS-as-a-service for OpenStack deployments. In order to minimize necessary changes to the datacenter DNS, a DNS service is hosted as a static pod to serve those hostnames to the nodes in the cluster.",
														MarkdownDescription: "nodeDNSIP is the IP address for the internal DNS used by the nodes. Unlike the one managed by the DNS operator, 'NodeDNSIP' provides name resolution for the nodes themselves. There is no DNS-as-a-service for OpenStack deployments. In order to minimize necessary changes to the datacenter DNS, a DNS service is hosted as a static pod to serve those hostnames to the nodes in the cluster.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"ovirt": schema.SingleNestedAttribute{
												Description:         "Ovirt contains settings specific to the oVirt infrastructure provider.",
												MarkdownDescription: "Ovirt contains settings specific to the oVirt infrastructure provider.",
												Attributes: map[string]schema.Attribute{
													"api_server_internal_ip": schema.StringAttribute{
														Description:         "apiServerInternalIP is an IP address to contact the Kubernetes API server that can be used by components inside the cluster, like kubelets using the infrastructure rather than Kubernetes networking. It is the IP that the Infrastructure.status.apiServerInternalURI points to. It is the IP for a self-hosted load balancer in front of the API servers.  Deprecated: Use APIServerInternalIPs instead.",
														MarkdownDescription: "apiServerInternalIP is an IP address to contact the Kubernetes API server that can be used by components inside the cluster, like kubelets using the infrastructure rather than Kubernetes networking. It is the IP that the Infrastructure.status.apiServerInternalURI points to. It is the IP for a self-hosted load balancer in front of the API servers.  Deprecated: Use APIServerInternalIPs instead.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"api_server_internal_i_ps": schema.ListAttribute{
														Description:         "apiServerInternalIPs are the IP addresses to contact the Kubernetes API server that can be used by components inside the cluster, like kubelets using the infrastructure rather than Kubernetes networking. These are the IPs for a self-hosted load balancer in front of the API servers. In dual stack clusters this list contains two IPs otherwise only one.",
														MarkdownDescription: "apiServerInternalIPs are the IP addresses to contact the Kubernetes API server that can be used by components inside the cluster, like kubelets using the infrastructure rather than Kubernetes networking. These are the IPs for a self-hosted load balancer in front of the API servers. In dual stack clusters this list contains two IPs otherwise only one.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ingress_ip": schema.StringAttribute{
														Description:         "ingressIP is an external IP which routes to the default ingress controller. The IP is a suitable target of a wildcard DNS record used to resolve default route host names.  Deprecated: Use IngressIPs instead.",
														MarkdownDescription: "ingressIP is an external IP which routes to the default ingress controller. The IP is a suitable target of a wildcard DNS record used to resolve default route host names.  Deprecated: Use IngressIPs instead.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ingress_i_ps": schema.ListAttribute{
														Description:         "ingressIPs are the external IPs which route to the default ingress controller. The IPs are suitable targets of a wildcard DNS record used to resolve default route host names. In dual stack clusters this list contains two IPs otherwise only one.",
														MarkdownDescription: "ingressIPs are the external IPs which route to the default ingress controller. The IPs are suitable targets of a wildcard DNS record used to resolve default route host names. In dual stack clusters this list contains two IPs otherwise only one.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"node_dnsip": schema.StringAttribute{
														Description:         "deprecated: as of 4.6, this field is no longer set or honored.  It will be removed in a future release.",
														MarkdownDescription: "deprecated: as of 4.6, this field is no longer set or honored.  It will be removed in a future release.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"powervs": schema.SingleNestedAttribute{
												Description:         "PowerVS contains settings specific to the Power Systems Virtual Servers infrastructure provider.",
												MarkdownDescription: "PowerVS contains settings specific to the Power Systems Virtual Servers infrastructure provider.",
												Attributes: map[string]schema.Attribute{
													"cis_instance_crn": schema.StringAttribute{
														Description:         "CISInstanceCRN is the CRN of the Cloud Internet Services instance managing the DNS zone for the cluster's base domain",
														MarkdownDescription: "CISInstanceCRN is the CRN of the Cloud Internet Services instance managing the DNS zone for the cluster's base domain",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"dns_instance_crn": schema.StringAttribute{
														Description:         "DNSInstanceCRN is the CRN of the DNS Services instance managing the DNS zone for the cluster's base domain",
														MarkdownDescription: "DNSInstanceCRN is the CRN of the DNS Services instance managing the DNS zone for the cluster's base domain",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"region": schema.StringAttribute{
														Description:         "region holds the default Power VS region for new Power VS resources created by the cluster.",
														MarkdownDescription: "region holds the default Power VS region for new Power VS resources created by the cluster.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"resource_group": schema.StringAttribute{
														Description:         "resourceGroup is the resource group name for new IBMCloud resources created for a cluster. The resource group specified here will be used by cluster-image-registry-operator to set up a COS Instance in IBMCloud for the cluster registry. More about resource groups can be found here: https://cloud.ibm.com/docs/account?topic=account-rgs. When omitted, the image registry operator won't be able to configure storage, which results in the image registry cluster operator not being in an available state.",
														MarkdownDescription: "resourceGroup is the resource group name for new IBMCloud resources created for a cluster. The resource group specified here will be used by cluster-image-registry-operator to set up a COS Instance in IBMCloud for the cluster registry. More about resource groups can be found here: https://cloud.ibm.com/docs/account?topic=account-rgs. When omitted, the image registry operator won't be able to configure storage, which results in the image registry cluster operator not being in an available state.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtMost(40),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9-_ ]+$`), ""),
														},
													},

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

													"zone": schema.StringAttribute{
														Description:         "zone holds the default zone for the new Power VS resources created by the cluster. Note: Currently only single-zone OCP clusters are supported",
														MarkdownDescription: "zone holds the default zone for the new Power VS resources created by the cluster. Note: Currently only single-zone OCP clusters are supported",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"type": schema.StringAttribute{
												Description:         "type is the underlying infrastructure provider for the cluster. This value controls whether infrastructure automation such as service load balancers, dynamic volume provisioning, machine creation and deletion, and other integrations are enabled. If None, no infrastructure automation is enabled. Allowed values are 'AWS', 'Azure', 'BareMetal', 'GCP', 'Libvirt', 'OpenStack', 'VSphere', 'oVirt', 'EquinixMetal', 'PowerVS', 'AlibabaCloud', 'Nutanix' and 'None'. Individual components may not support all platforms, and must handle unrecognized platforms as None if they do not support that platform.  This value will be synced with to the 'status.platform' and 'status.platformStatus.type'. Currently this value cannot be changed once set.",
												MarkdownDescription: "type is the underlying infrastructure provider for the cluster. This value controls whether infrastructure automation such as service load balancers, dynamic volume provisioning, machine creation and deletion, and other integrations are enabled. If None, no infrastructure automation is enabled. Allowed values are 'AWS', 'Azure', 'BareMetal', 'GCP', 'Libvirt', 'OpenStack', 'VSphere', 'oVirt', 'EquinixMetal', 'PowerVS', 'AlibabaCloud', 'Nutanix' and 'None'. Individual components may not support all platforms, and must handle unrecognized platforms as None if they do not support that platform.  This value will be synced with to the 'status.platform' and 'status.platformStatus.type'. Currently this value cannot be changed once set.",
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
													"api_server_internal_ip": schema.StringAttribute{
														Description:         "apiServerInternalIP is an IP address to contact the Kubernetes API server that can be used by components inside the cluster, like kubelets using the infrastructure rather than Kubernetes networking. It is the IP that the Infrastructure.status.apiServerInternalURI points to. It is the IP for a self-hosted load balancer in front of the API servers.  Deprecated: Use APIServerInternalIPs instead.",
														MarkdownDescription: "apiServerInternalIP is an IP address to contact the Kubernetes API server that can be used by components inside the cluster, like kubelets using the infrastructure rather than Kubernetes networking. It is the IP that the Infrastructure.status.apiServerInternalURI points to. It is the IP for a self-hosted load balancer in front of the API servers.  Deprecated: Use APIServerInternalIPs instead.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"api_server_internal_i_ps": schema.ListAttribute{
														Description:         "apiServerInternalIPs are the IP addresses to contact the Kubernetes API server that can be used by components inside the cluster, like kubelets using the infrastructure rather than Kubernetes networking. These are the IPs for a self-hosted load balancer in front of the API servers. In dual stack clusters this list contains two IPs otherwise only one.",
														MarkdownDescription: "apiServerInternalIPs are the IP addresses to contact the Kubernetes API server that can be used by components inside the cluster, like kubelets using the infrastructure rather than Kubernetes networking. These are the IPs for a self-hosted load balancer in front of the API servers. In dual stack clusters this list contains two IPs otherwise only one.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ingress_ip": schema.StringAttribute{
														Description:         "ingressIP is an external IP which routes to the default ingress controller. The IP is a suitable target of a wildcard DNS record used to resolve default route host names.  Deprecated: Use IngressIPs instead.",
														MarkdownDescription: "ingressIP is an external IP which routes to the default ingress controller. The IP is a suitable target of a wildcard DNS record used to resolve default route host names.  Deprecated: Use IngressIPs instead.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ingress_i_ps": schema.ListAttribute{
														Description:         "ingressIPs are the external IPs which route to the default ingress controller. The IPs are suitable targets of a wildcard DNS record used to resolve default route host names. In dual stack clusters this list contains two IPs otherwise only one.",
														MarkdownDescription: "ingressIPs are the external IPs which route to the default ingress controller. The IPs are suitable targets of a wildcard DNS record used to resolve default route host names. In dual stack clusters this list contains two IPs otherwise only one.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"machine_networks": schema.ListAttribute{
														Description:         "machineNetworks are IP networks used to connect all the OpenShift cluster nodes.",
														MarkdownDescription: "machineNetworks are IP networks used to connect all the OpenShift cluster nodes.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"node_dnsip": schema.StringAttribute{
														Description:         "nodeDNSIP is the IP address for the internal DNS used by the nodes. Unlike the one managed by the DNS operator, 'NodeDNSIP' provides name resolution for the nodes themselves. There is no DNS-as-a-service for vSphere deployments. In order to minimize necessary changes to the datacenter DNS, a DNS service is hosted as a static pod to serve those hostnames to the nodes in the cluster.",
														MarkdownDescription: "nodeDNSIP is the IP address for the internal DNS used by the nodes. Unlike the one managed by the DNS operator, 'NodeDNSIP' provides name resolution for the nodes themselves. There is no DNS-as-a-service for vSphere deployments. In order to minimize necessary changes to the datacenter DNS, a DNS service is hosted as a static pod to serve those hostnames to the nodes in the cluster.",
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

					"internal_registry_pull_secret": schema.StringAttribute{
						Description:         "internalRegistryPullSecret is the pull secret for the internal registry, used by rpm-ostree to pull images from the internal registry if present",
						MarkdownDescription: "internalRegistryPullSecret is the pull secret for the internal registry, used by rpm-ostree to pull images from the internal registry if present",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							validators.Base64Validator(),
						},
					},

					"ip_families": schema.StringAttribute{
						Description:         "ipFamilies indicates the IP families in use by the cluster network",
						MarkdownDescription: "ipFamilies indicates the IP families in use by the cluster network",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"kube_api_server_serving_ca_data": schema.StringAttribute{
						Description:         "kubeAPIServerServingCAData managed Kubelet to API Server Cert... Rotated automatically",
						MarkdownDescription: "kubeAPIServerServingCAData managed Kubelet to API Server Cert... Rotated automatically",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							validators.Base64Validator(),
						},
					},

					"network": schema.SingleNestedAttribute{
						Description:         "Network contains additional network related information",
						MarkdownDescription: "Network contains additional network related information",
						Attributes: map[string]schema.Attribute{
							"mtu_migration": schema.SingleNestedAttribute{
								Description:         "MTUMigration contains the MTU migration configuration.",
								MarkdownDescription: "MTUMigration contains the MTU migration configuration.",
								Attributes: map[string]schema.Attribute{
									"machine": schema.SingleNestedAttribute{
										Description:         "Machine contains MTU migration configuration for the machine's uplink.",
										MarkdownDescription: "Machine contains MTU migration configuration for the machine's uplink.",
										Attributes: map[string]schema.Attribute{
											"from": schema.Int64Attribute{
												Description:         "From is the MTU to migrate from.",
												MarkdownDescription: "From is the MTU to migrate from.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"to": schema.Int64Attribute{
												Description:         "To is the MTU to migrate to.",
												MarkdownDescription: "To is the MTU to migrate to.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"network": schema.SingleNestedAttribute{
										Description:         "Network contains MTU migration configuration for the default network.",
										MarkdownDescription: "Network contains MTU migration configuration for the default network.",
										Attributes: map[string]schema.Attribute{
											"from": schema.Int64Attribute{
												Description:         "From is the MTU to migrate from.",
												MarkdownDescription: "From is the MTU to migrate from.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"to": schema.Int64Attribute{
												Description:         "To is the MTU to migrate to.",
												MarkdownDescription: "To is the MTU to migrate to.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
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
						Required: true,
						Optional: false,
						Computed: false,
					},

					"network_type": schema.StringAttribute{
						Description:         "networkType holds the type of network the cluster is using XXX: this is temporary and will be dropped as soon as possible in favor of a better support to start network related services the proper way. Nobody is also changing this once the cluster is up and running the first time, so, disallow regeneration if this changes.",
						MarkdownDescription: "networkType holds the type of network the cluster is using XXX: this is temporary and will be dropped as soon as possible in favor of a better support to start network related services the proper way. Nobody is also changing this once the cluster is up and running the first time, so, disallow regeneration if this changes.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"os_image_url": schema.StringAttribute{
						Description:         "OSImageURL is the old-format container image that contains the OS update payload.",
						MarkdownDescription: "OSImageURL is the old-format container image that contains the OS update payload.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"platform": schema.StringAttribute{
						Description:         "platform is deprecated, use Infra.Status.PlatformStatus.Type instead",
						MarkdownDescription: "platform is deprecated, use Infra.Status.PlatformStatus.Type instead",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"proxy": schema.SingleNestedAttribute{
						Description:         "proxy holds the current proxy configuration for the nodes",
						MarkdownDescription: "proxy holds the current proxy configuration for the nodes",
						Attributes: map[string]schema.Attribute{
							"http_proxy": schema.StringAttribute{
								Description:         "httpProxy is the URL of the proxy for HTTP requests.",
								MarkdownDescription: "httpProxy is the URL of the proxy for HTTP requests.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"https_proxy": schema.StringAttribute{
								Description:         "httpsProxy is the URL of the proxy for HTTPS requests.",
								MarkdownDescription: "httpsProxy is the URL of the proxy for HTTPS requests.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"no_proxy": schema.StringAttribute{
								Description:         "noProxy is a comma-separated list of hostnames and/or CIDRs for which the proxy should not be used.",
								MarkdownDescription: "noProxy is a comma-separated list of hostnames and/or CIDRs for which the proxy should not be used.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"pull_secret": schema.SingleNestedAttribute{
						Description:         "pullSecret is the default pull secret that needs to be installed on all machines.",
						MarkdownDescription: "pullSecret is the default pull secret that needs to be installed on all machines.",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "API version of the referent.",
								MarkdownDescription: "API version of the referent.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"field_path": schema.StringAttribute{
								Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
								MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								MarkdownDescription: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
								MarkdownDescription: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resource_version": schema.StringAttribute{
								Description:         "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
								MarkdownDescription: "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"uid": schema.StringAttribute{
								Description:         "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
								MarkdownDescription: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"release_image": schema.StringAttribute{
						Description:         "releaseImage is the image used when installing the cluster",
						MarkdownDescription: "releaseImage is the image used when installing the cluster",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"root_ca_data": schema.StringAttribute{
						Description:         "rootCAData specifies the root CA data",
						MarkdownDescription: "rootCAData specifies the root CA data",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							validators.Base64Validator(),
						},
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *MachineconfigurationOpenshiftIoControllerConfigV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_machineconfiguration_openshift_io_controller_config_v1_manifest")

	var model MachineconfigurationOpenshiftIoControllerConfigV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("machineconfiguration.openshift.io/v1")
	model.Kind = pointer.String("ControllerConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
