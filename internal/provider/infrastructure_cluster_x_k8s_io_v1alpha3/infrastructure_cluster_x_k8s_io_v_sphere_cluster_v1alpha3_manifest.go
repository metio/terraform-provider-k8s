/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package infrastructure_cluster_x_k8s_io_v1alpha3

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
	_ datasource.DataSource = &InfrastructureClusterXK8SIoVsphereClusterV1Alpha3Manifest{}
)

func NewInfrastructureClusterXK8SIoVsphereClusterV1Alpha3Manifest() datasource.DataSource {
	return &InfrastructureClusterXK8SIoVsphereClusterV1Alpha3Manifest{}
}

type InfrastructureClusterXK8SIoVsphereClusterV1Alpha3Manifest struct{}

type InfrastructureClusterXK8SIoVsphereClusterV1Alpha3ManifestData struct {
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
		CloudProviderConfiguration *struct {
			Disk *struct {
				ScsiControllerType *string `tfsdk:"scsi_controller_type" json:"scsiControllerType,omitempty"`
			} `tfsdk:"disk" json:"disk,omitempty"`
			Global *struct {
				ApiBindPort       *string `tfsdk:"api_bind_port" json:"apiBindPort,omitempty"`
				ApiDisable        *bool   `tfsdk:"api_disable" json:"apiDisable,omitempty"`
				CaFile            *string `tfsdk:"ca_file" json:"caFile,omitempty"`
				Datacenters       *string `tfsdk:"datacenters" json:"datacenters,omitempty"`
				Insecure          *bool   `tfsdk:"insecure" json:"insecure,omitempty"`
				Password          *string `tfsdk:"password" json:"password,omitempty"`
				Port              *string `tfsdk:"port" json:"port,omitempty"`
				RoundTripperCount *int64  `tfsdk:"round_tripper_count" json:"roundTripperCount,omitempty"`
				SecretName        *string `tfsdk:"secret_name" json:"secretName,omitempty"`
				SecretNamespace   *string `tfsdk:"secret_namespace" json:"secretNamespace,omitempty"`
				SecretsDirectory  *string `tfsdk:"secrets_directory" json:"secretsDirectory,omitempty"`
				ServiceAccount    *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				Thumbprint        *string `tfsdk:"thumbprint" json:"thumbprint,omitempty"`
				Username          *string `tfsdk:"username" json:"username,omitempty"`
			} `tfsdk:"global" json:"global,omitempty"`
			Labels *struct {
				Region *string `tfsdk:"region" json:"region,omitempty"`
				Zone   *string `tfsdk:"zone" json:"zone,omitempty"`
			} `tfsdk:"labels" json:"labels,omitempty"`
			Network *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"network" json:"network,omitempty"`
			ProviderConfig *struct {
				Cloud *struct {
					ControllerImage *string            `tfsdk:"controller_image" json:"controllerImage,omitempty"`
					ExtraArgs       *map[string]string `tfsdk:"extra_args" json:"extraArgs,omitempty"`
				} `tfsdk:"cloud" json:"cloud,omitempty"`
				Storage *struct {
					AttacherImage       *string `tfsdk:"attacher_image" json:"attacherImage,omitempty"`
					ControllerImage     *string `tfsdk:"controller_image" json:"controllerImage,omitempty"`
					LivenessProbeImage  *string `tfsdk:"liveness_probe_image" json:"livenessProbeImage,omitempty"`
					MetadataSyncerImage *string `tfsdk:"metadata_syncer_image" json:"metadataSyncerImage,omitempty"`
					NodeDriverImage     *string `tfsdk:"node_driver_image" json:"nodeDriverImage,omitempty"`
					ProvisionerImage    *string `tfsdk:"provisioner_image" json:"provisionerImage,omitempty"`
					RegistrarImage      *string `tfsdk:"registrar_image" json:"registrarImage,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"provider_config" json:"providerConfig,omitempty"`
			VirtualCenter *struct {
				Datacenters       *string `tfsdk:"datacenters" json:"datacenters,omitempty"`
				Password          *string `tfsdk:"password" json:"password,omitempty"`
				Port              *string `tfsdk:"port" json:"port,omitempty"`
				RoundTripperCount *int64  `tfsdk:"round_tripper_count" json:"roundTripperCount,omitempty"`
				Thumbprint        *string `tfsdk:"thumbprint" json:"thumbprint,omitempty"`
				Username          *string `tfsdk:"username" json:"username,omitempty"`
			} `tfsdk:"virtual_center" json:"virtualCenter,omitempty"`
			Workspace *struct {
				Datacenter   *string `tfsdk:"datacenter" json:"datacenter,omitempty"`
				Datastore    *string `tfsdk:"datastore" json:"datastore,omitempty"`
				Folder       *string `tfsdk:"folder" json:"folder,omitempty"`
				ResourcePool *string `tfsdk:"resource_pool" json:"resourcePool,omitempty"`
				Server       *string `tfsdk:"server" json:"server,omitempty"`
			} `tfsdk:"workspace" json:"workspace,omitempty"`
		} `tfsdk:"cloud_provider_configuration" json:"cloudProviderConfiguration,omitempty"`
		ControlPlaneEndpoint *struct {
			Host *string `tfsdk:"host" json:"host,omitempty"`
			Port *int64  `tfsdk:"port" json:"port,omitempty"`
		} `tfsdk:"control_plane_endpoint" json:"controlPlaneEndpoint,omitempty"`
		IdentityRef *struct {
			Kind *string `tfsdk:"kind" json:"kind,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"identity_ref" json:"identityRef,omitempty"`
		Insecure        *bool `tfsdk:"insecure" json:"insecure,omitempty"`
		LoadBalancerRef *struct {
			ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
			Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
			Name            *string `tfsdk:"name" json:"name,omitempty"`
			Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
			ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
			Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
		} `tfsdk:"load_balancer_ref" json:"loadBalancerRef,omitempty"`
		Server     *string `tfsdk:"server" json:"server,omitempty"`
		Thumbprint *string `tfsdk:"thumbprint" json:"thumbprint,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *InfrastructureClusterXK8SIoVsphereClusterV1Alpha3Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_infrastructure_cluster_x_k8s_io_v_sphere_cluster_v1alpha3_manifest"
}

func (r *InfrastructureClusterXK8SIoVsphereClusterV1Alpha3Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "VSphereCluster is the Schema for the vsphereclusters APIDeprecated: This type will be removed in one of the next releases.",
		MarkdownDescription: "VSphereCluster is the Schema for the vsphereclusters APIDeprecated: This type will be removed in one of the next releases.",
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
				Description:         "VSphereClusterSpec defines the desired state of VSphereCluster.",
				MarkdownDescription: "VSphereClusterSpec defines the desired state of VSphereCluster.",
				Attributes: map[string]schema.Attribute{
					"cloud_provider_configuration": schema.SingleNestedAttribute{
						Description:         "CloudProviderConfiguration holds the cluster-wide configuration for the vSphere cloud provider.Deprecated: will be removed in v1alpha4.",
						MarkdownDescription: "CloudProviderConfiguration holds the cluster-wide configuration for the vSphere cloud provider.Deprecated: will be removed in v1alpha4.",
						Attributes: map[string]schema.Attribute{
							"disk": schema.SingleNestedAttribute{
								Description:         "Disk is the vSphere cloud provider's disk configuration.",
								MarkdownDescription: "Disk is the vSphere cloud provider's disk configuration.",
								Attributes: map[string]schema.Attribute{
									"scsi_controller_type": schema.StringAttribute{
										Description:         "SCSIControllerType defines SCSI controller to be used.",
										MarkdownDescription: "SCSIControllerType defines SCSI controller to be used.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"global": schema.SingleNestedAttribute{
								Description:         "Global is the vSphere cloud provider's global configuration.",
								MarkdownDescription: "Global is the vSphere cloud provider's global configuration.",
								Attributes: map[string]schema.Attribute{
									"api_bind_port": schema.StringAttribute{
										Description:         "APIBindPort configures the vSphere cloud controller manager API port.Defaults to 43001.",
										MarkdownDescription: "APIBindPort configures the vSphere cloud controller manager API port.Defaults to 43001.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"api_disable": schema.BoolAttribute{
										Description:         "APIDisable disables the vSphere cloud controller manager API.Defaults to true.",
										MarkdownDescription: "APIDisable disables the vSphere cloud controller manager API.Defaults to true.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ca_file": schema.StringAttribute{
										Description:         "CAFile Specifies the path to a CA certificate in PEM format.If not configured, the system's CA certificates will be used.",
										MarkdownDescription: "CAFile Specifies the path to a CA certificate in PEM format.If not configured, the system's CA certificates will be used.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"datacenters": schema.StringAttribute{
										Description:         "Datacenters is a CSV string of the datacenters in which VMs are located.",
										MarkdownDescription: "Datacenters is a CSV string of the datacenters in which VMs are located.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"insecure": schema.BoolAttribute{
										Description:         "Insecure is a flag that disables TLS peer verification.",
										MarkdownDescription: "Insecure is a flag that disables TLS peer verification.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"password": schema.StringAttribute{
										Description:         "Password is the password used to access a vSphere endpoint.",
										MarkdownDescription: "Password is the password used to access a vSphere endpoint.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"port": schema.StringAttribute{
										Description:         "Port is the port on which the vSphere endpoint is listening.Defaults to 443.",
										MarkdownDescription: "Port is the port on which the vSphere endpoint is listening.Defaults to 443.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"round_tripper_count": schema.Int64Attribute{
										Description:         "RoundTripperCount specifies the SOAP round tripper count(retries = RoundTripper - 1)",
										MarkdownDescription: "RoundTripperCount specifies the SOAP round tripper count(retries = RoundTripper - 1)",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"secret_name": schema.StringAttribute{
										Description:         "SecretName is the name of the Kubernetes secret in which the vSpherecredentials are located.",
										MarkdownDescription: "SecretName is the name of the Kubernetes secret in which the vSpherecredentials are located.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"secret_namespace": schema.StringAttribute{
										Description:         "SecretNamespace is the namespace for SecretName.",
										MarkdownDescription: "SecretNamespace is the namespace for SecretName.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"secrets_directory": schema.StringAttribute{
										Description:         "SecretsDirectory is a directory in which secrets may be found. Thismay used in the event that:1. It is not desirable to use the K8s API to watch changes to secrets2. The cloud controller manager is not running in a K8s environment,   such as DC/OS. For example, the container storage interface (CSI) is   container orcehstrator (CO) agnostic, and should support non-K8s COs.Defaults to /etc/cloud/credentials.",
										MarkdownDescription: "SecretsDirectory is a directory in which secrets may be found. Thismay used in the event that:1. It is not desirable to use the K8s API to watch changes to secrets2. The cloud controller manager is not running in a K8s environment,   such as DC/OS. For example, the container storage interface (CSI) is   container orcehstrator (CO) agnostic, and should support non-K8s COs.Defaults to /etc/cloud/credentials.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"service_account": schema.StringAttribute{
										Description:         "ServiceAccount is the Kubernetes service account used to launch the cloudcontroller manager.Defaults to cloud-controller-manager.",
										MarkdownDescription: "ServiceAccount is the Kubernetes service account used to launch the cloudcontroller manager.Defaults to cloud-controller-manager.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"thumbprint": schema.StringAttribute{
										Description:         "Thumbprint is the cryptographic thumbprint of the vSphere endpoint'scertificate.",
										MarkdownDescription: "Thumbprint is the cryptographic thumbprint of the vSphere endpoint'scertificate.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"username": schema.StringAttribute{
										Description:         "Username is the username used to access a vSphere endpoint.",
										MarkdownDescription: "Username is the username used to access a vSphere endpoint.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"labels": schema.SingleNestedAttribute{
								Description:         "Labels is the vSphere cloud provider's zone and region configuration.",
								MarkdownDescription: "Labels is the vSphere cloud provider's zone and region configuration.",
								Attributes: map[string]schema.Attribute{
									"region": schema.StringAttribute{
										Description:         "Region is the region in which VMs are created/located.",
										MarkdownDescription: "Region is the region in which VMs are created/located.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"zone": schema.StringAttribute{
										Description:         "Zone is the zone in which VMs are created/located.",
										MarkdownDescription: "Zone is the zone in which VMs are created/located.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"network": schema.SingleNestedAttribute{
								Description:         "Network is the vSphere cloud provider's network configuration.",
								MarkdownDescription: "Network is the vSphere cloud provider's network configuration.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name is the name of the network to which VMs are connected.",
										MarkdownDescription: "Name is the name of the network to which VMs are connected.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"provider_config": schema.SingleNestedAttribute{
								Description:         "CPIProviderConfig contains extra information used to configure thevSphere cloud provider.",
								MarkdownDescription: "CPIProviderConfig contains extra information used to configure thevSphere cloud provider.",
								Attributes: map[string]schema.Attribute{
									"cloud": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"controller_image": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"extra_args": schema.MapAttribute{
												Description:         "ExtraArgs passes through extra arguments to the cloud provider.The arguments here are passed to the cloud provider daemonset specification",
												MarkdownDescription: "ExtraArgs passes through extra arguments to the cloud provider.The arguments here are passed to the cloud provider daemonset specification",
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

									"storage": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"attacher_image": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"controller_image": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"liveness_probe_image": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"metadata_syncer_image": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"node_driver_image": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"provisioner_image": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"registrar_image": schema.StringAttribute{
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"virtual_center": schema.SingleNestedAttribute{
								Description:         "VCenter is a list of vCenter configurations.",
								MarkdownDescription: "VCenter is a list of vCenter configurations.",
								Attributes: map[string]schema.Attribute{
									"datacenters": schema.StringAttribute{
										Description:         "Datacenters is a CSV string of the datacenters in which VMs are located.",
										MarkdownDescription: "Datacenters is a CSV string of the datacenters in which VMs are located.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"password": schema.StringAttribute{
										Description:         "Password is the password used to access a vSphere endpoint.",
										MarkdownDescription: "Password is the password used to access a vSphere endpoint.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"port": schema.StringAttribute{
										Description:         "Port is the port on which the vSphere endpoint is listening.Defaults to 443.",
										MarkdownDescription: "Port is the port on which the vSphere endpoint is listening.Defaults to 443.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"round_tripper_count": schema.Int64Attribute{
										Description:         "RoundTripperCount specifies the SOAP round tripper count(retries = RoundTripper - 1)",
										MarkdownDescription: "RoundTripperCount specifies the SOAP round tripper count(retries = RoundTripper - 1)",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"thumbprint": schema.StringAttribute{
										Description:         "Thumbprint is the cryptographic thumbprint of the vSphere endpoint'scertificate.",
										MarkdownDescription: "Thumbprint is the cryptographic thumbprint of the vSphere endpoint'scertificate.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"username": schema.StringAttribute{
										Description:         "Username is the username used to access a vSphere endpoint.",
										MarkdownDescription: "Username is the username used to access a vSphere endpoint.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"workspace": schema.SingleNestedAttribute{
								Description:         "Workspace is the vSphere cloud provider's workspace configuration.",
								MarkdownDescription: "Workspace is the vSphere cloud provider's workspace configuration.",
								Attributes: map[string]schema.Attribute{
									"datacenter": schema.StringAttribute{
										Description:         "Datacenter is the datacenter in which VMs are created/located.",
										MarkdownDescription: "Datacenter is the datacenter in which VMs are created/located.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"datastore": schema.StringAttribute{
										Description:         "Datastore is the datastore in which VMs are created/located.",
										MarkdownDescription: "Datastore is the datastore in which VMs are created/located.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"folder": schema.StringAttribute{
										Description:         "Folder is the folder in which VMs are created/located.",
										MarkdownDescription: "Folder is the folder in which VMs are created/located.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"resource_pool": schema.StringAttribute{
										Description:         "ResourcePool is the resource pool in which VMs are created/located.",
										MarkdownDescription: "ResourcePool is the resource pool in which VMs are created/located.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"server": schema.StringAttribute{
										Description:         "Server is the IP address or FQDN of the vSphere endpoint.",
										MarkdownDescription: "Server is the IP address or FQDN of the vSphere endpoint.",
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

					"identity_ref": schema.SingleNestedAttribute{
						Description:         "IdentityRef is a reference to either a Secret or VSphereClusterIdentity that containsthe identity to use when reconciling the cluster.",
						MarkdownDescription: "IdentityRef is a reference to either a Secret or VSphereClusterIdentity that containsthe identity to use when reconciling the cluster.",
						Attributes: map[string]schema.Attribute{
							"kind": schema.StringAttribute{
								Description:         "Kind of the identity. Can either be VSphereClusterIdentity or Secret",
								MarkdownDescription: "Kind of the identity. Can either be VSphereClusterIdentity or Secret",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("VSphereClusterIdentity", "Secret"),
								},
							},

							"name": schema.StringAttribute{
								Description:         "Name of the identity.",
								MarkdownDescription: "Name of the identity.",
								Required:            true,
								Optional:            false,
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

					"insecure": schema.BoolAttribute{
						Description:         "Insecure is a flag that controls whether to validate thevSphere server's certificate.Deprecated: will be removed in v1alpha4.",
						MarkdownDescription: "Insecure is a flag that controls whether to validate thevSphere server's certificate.Deprecated: will be removed in v1alpha4.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"load_balancer_ref": schema.SingleNestedAttribute{
						Description:         "LoadBalancerRef may be used to enable a control plane load balancerfor this cluster.When a LoadBalancerRef is provided, the VSphereCluster.Status.Ready fieldwill not be true until the referenced resource is Status.Ready and has anon-empty Status.Address value.Deprecated: will be removed in v1alpha4.",
						MarkdownDescription: "LoadBalancerRef may be used to enable a control plane load balancerfor this cluster.When a LoadBalancerRef is provided, the VSphereCluster.Status.Ready fieldwill not be true until the referenced resource is Status.Ready and has anon-empty Status.Address value.Deprecated: will be removed in v1alpha4.",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "API version of the referent.",
								MarkdownDescription: "API version of the referent.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"field_path": schema.StringAttribute{
								Description:         "If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.",
								MarkdownDescription: "If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								MarkdownDescription: "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
								MarkdownDescription: "Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resource_version": schema.StringAttribute{
								Description:         "Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
								MarkdownDescription: "Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"uid": schema.StringAttribute{
								Description:         "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
								MarkdownDescription: "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"server": schema.StringAttribute{
						Description:         "Server is the address of the vSphere endpoint.",
						MarkdownDescription: "Server is the address of the vSphere endpoint.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"thumbprint": schema.StringAttribute{
						Description:         "Thumbprint is the colon-separated SHA-1 checksum of the given vCenter server's host certificateWhen provided, Insecure should not be set to true",
						MarkdownDescription: "Thumbprint is the colon-separated SHA-1 checksum of the given vCenter server's host certificateWhen provided, Insecure should not be set to true",
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

func (r *InfrastructureClusterXK8SIoVsphereClusterV1Alpha3Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_infrastructure_cluster_x_k8s_io_v_sphere_cluster_v1alpha3_manifest")

	var model InfrastructureClusterXK8SIoVsphereClusterV1Alpha3ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("infrastructure.cluster.x-k8s.io/v1alpha3")
	model.Kind = pointer.String("VSphereCluster")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
