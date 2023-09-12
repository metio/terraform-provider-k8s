/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package infrastructure_cluster_x_k8s_io_v1alpha3

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	k8sTypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"strings"
	"time"
)

var (
	_ resource.Resource                = &InfrastructureClusterXK8SIoVsphereClusterV1Alpha3Resource{}
	_ resource.ResourceWithConfigure   = &InfrastructureClusterXK8SIoVsphereClusterV1Alpha3Resource{}
	_ resource.ResourceWithImportState = &InfrastructureClusterXK8SIoVsphereClusterV1Alpha3Resource{}
)

func NewInfrastructureClusterXK8SIoVsphereClusterV1Alpha3Resource() resource.Resource {
	return &InfrastructureClusterXK8SIoVsphereClusterV1Alpha3Resource{}
}

type InfrastructureClusterXK8SIoVsphereClusterV1Alpha3Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type InfrastructureClusterXK8SIoVsphereClusterV1Alpha3ResourceData struct {
	ID                  types.String `tfsdk:"id" json:"-"`
	ForceConflicts      types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager        types.String `tfsdk:"field_manager" json:"-"`
	DeletionPropagation types.String `tfsdk:"deletion_propagation" json:"-"`
	WaitForUpsert       types.List   `tfsdk:"wait_for_upsert" json:"-"`
	WaitForDelete       types.Object `tfsdk:"wait_for_delete" json:"-"`

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

func (r *InfrastructureClusterXK8SIoVsphereClusterV1Alpha3Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_infrastructure_cluster_x_k8s_io_v_sphere_cluster_v1alpha3"
}

func (r *InfrastructureClusterXK8SIoVsphereClusterV1Alpha3Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "VSphereCluster is the Schema for the vsphereclusters API  Deprecated: This type will be removed in one of the next releases.",
		MarkdownDescription: "VSphereCluster is the Schema for the vsphereclusters API  Deprecated: This type will be removed in one of the next releases.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"force_conflicts": schema.BoolAttribute{
				Description:         "If 'true', server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "If `true`, server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"field_manager": schema.StringAttribute{
				Description:         "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},

			"deletion_propagation": schema.StringAttribute{
				Description:         "Decides if a deletion will propagate to the dependents of the object, and how the garbage collector will handle the propagation.",
				MarkdownDescription: "Decides if a deletion will propagate to the dependents of the object, and how the garbage collector will handle the propagation.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("Orphan", "Background", "Foreground"),
				},
			},

			"wait_for_upsert": schema.ListNestedAttribute{
				Description:         "Wait for specific conditions after create/update of resources.",
				MarkdownDescription: "Wait for specific conditions after create/update of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"jsonpath": schema.StringAttribute{
							Description:         "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							MarkdownDescription: "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							Required:            true,
							Optional:            false,
							Computed:            false,
						},
						"value": schema.StringAttribute{
							Description:         "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							MarkdownDescription: "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							Required:            false,
							Optional:            true,
							Computed:            true,
						},
						"timeout": schema.Int64Attribute{
							Description:         "The number of seconds to wait before giving up. Zero means check once and don't wait.",
							MarkdownDescription: "The number of seconds to wait before giving up. Zero means check once and don't wait.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             int64default.StaticInt64(30),
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
						},
						"poll_interval": schema.Int64Attribute{
							Description:         "The number of seconds to wait before checking again.",
							MarkdownDescription: "The number of seconds to wait before checking again.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             int64default.StaticInt64(5),
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
						},
					},
				},
			},

			"wait_for_delete": schema.SingleNestedAttribute{
				Description:         "Wait for deletion of resources.",
				MarkdownDescription: "Wait for deletion of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"timeout": schema.Int64Attribute{
						Description:         "The number of seconds to wait before giving up. Zero means check once and don't wait.",
						MarkdownDescription: "The number of seconds to wait before giving up. Zero means check once and don't wait.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             int64default.StaticInt64(30),
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},
					"poll_interval": schema.Int64Attribute{
						Description:         "The number of seconds to wait before checking again.",
						MarkdownDescription: "The number of seconds to wait before checking again.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             int64default.StaticInt64(5),
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},
				},
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},

					"labels": schema.MapAttribute{
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            true,
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
						Computed:            true,
						Validators: []validator.Map{
							validators.AnnotationValidator(),
						},
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "VSphereClusterSpec defines the desired state of VSphereCluster",
				MarkdownDescription: "VSphereClusterSpec defines the desired state of VSphereCluster",
				Attributes: map[string]schema.Attribute{
					"cloud_provider_configuration": schema.SingleNestedAttribute{
						Description:         "CloudProviderConfiguration holds the cluster-wide configuration for the DEPRECATED: will be removed in v1alpha4 vSphere cloud provider.",
						MarkdownDescription: "CloudProviderConfiguration holds the cluster-wide configuration for the DEPRECATED: will be removed in v1alpha4 vSphere cloud provider.",
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
										Description:         "APIBindPort configures the vSphere cloud controller manager API port. Defaults to 43001.",
										MarkdownDescription: "APIBindPort configures the vSphere cloud controller manager API port. Defaults to 43001.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"api_disable": schema.BoolAttribute{
										Description:         "APIDisable disables the vSphere cloud controller manager API. Defaults to true.",
										MarkdownDescription: "APIDisable disables the vSphere cloud controller manager API. Defaults to true.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ca_file": schema.StringAttribute{
										Description:         "CAFile Specifies the path to a CA certificate in PEM format. If not configured, the system's CA certificates will be used.",
										MarkdownDescription: "CAFile Specifies the path to a CA certificate in PEM format. If not configured, the system's CA certificates will be used.",
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
										Description:         "Port is the port on which the vSphere endpoint is listening. Defaults to 443.",
										MarkdownDescription: "Port is the port on which the vSphere endpoint is listening. Defaults to 443.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"round_tripper_count": schema.Int64Attribute{
										Description:         "RoundTripperCount specifies the SOAP round tripper count (retries = RoundTripper - 1)",
										MarkdownDescription: "RoundTripperCount specifies the SOAP round tripper count (retries = RoundTripper - 1)",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"secret_name": schema.StringAttribute{
										Description:         "SecretName is the name of the Kubernetes secret in which the vSphere credentials are located.",
										MarkdownDescription: "SecretName is the name of the Kubernetes secret in which the vSphere credentials are located.",
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
										Description:         "SecretsDirectory is a directory in which secrets may be found. This may used in the event that: 1. It is not desirable to use the K8s API to watch changes to secrets 2. The cloud controller manager is not running in a K8s environment, such as DC/OS. For example, the container storage interface (CSI) is container orcehstrator (CO) agnostic, and should support non-K8s COs. Defaults to /etc/cloud/credentials.",
										MarkdownDescription: "SecretsDirectory is a directory in which secrets may be found. This may used in the event that: 1. It is not desirable to use the K8s API to watch changes to secrets 2. The cloud controller manager is not running in a K8s environment, such as DC/OS. For example, the container storage interface (CSI) is container orcehstrator (CO) agnostic, and should support non-K8s COs. Defaults to /etc/cloud/credentials.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"service_account": schema.StringAttribute{
										Description:         "ServiceAccount is the Kubernetes service account used to launch the cloud controller manager. Defaults to cloud-controller-manager.",
										MarkdownDescription: "ServiceAccount is the Kubernetes service account used to launch the cloud controller manager. Defaults to cloud-controller-manager.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"thumbprint": schema.StringAttribute{
										Description:         "Thumbprint is the cryptographic thumbprint of the vSphere endpoint's certificate.",
										MarkdownDescription: "Thumbprint is the cryptographic thumbprint of the vSphere endpoint's certificate.",
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
								Description:         "CPIProviderConfig contains extra information used to configure the vSphere cloud provider.",
								MarkdownDescription: "CPIProviderConfig contains extra information used to configure the vSphere cloud provider.",
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
												Description:         "ExtraArgs passes through extra arguments to the cloud provider. The arguments here are passed to the cloud provider daemonset specification",
												MarkdownDescription: "ExtraArgs passes through extra arguments to the cloud provider. The arguments here are passed to the cloud provider daemonset specification",
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
										Description:         "Port is the port on which the vSphere endpoint is listening. Defaults to 443.",
										MarkdownDescription: "Port is the port on which the vSphere endpoint is listening. Defaults to 443.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"round_tripper_count": schema.Int64Attribute{
										Description:         "RoundTripperCount specifies the SOAP round tripper count (retries = RoundTripper - 1)",
										MarkdownDescription: "RoundTripperCount specifies the SOAP round tripper count (retries = RoundTripper - 1)",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"thumbprint": schema.StringAttribute{
										Description:         "Thumbprint is the cryptographic thumbprint of the vSphere endpoint's certificate.",
										MarkdownDescription: "Thumbprint is the cryptographic thumbprint of the vSphere endpoint's certificate.",
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
						Description:         "IdentityRef is a reference to either a Secret or VSphereClusterIdentity that contains the identity to use when reconciling the cluster.",
						MarkdownDescription: "IdentityRef is a reference to either a Secret or VSphereClusterIdentity that contains the identity to use when reconciling the cluster.",
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
						Description:         "Insecure is a flag that controls whether or not to validate the vSphere server's certificate. DEPRECATED: will be removed in v1alpha4",
						MarkdownDescription: "Insecure is a flag that controls whether or not to validate the vSphere server's certificate. DEPRECATED: will be removed in v1alpha4",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"load_balancer_ref": schema.SingleNestedAttribute{
						Description:         "LoadBalancerRef may be used to enable a control plane load balancer for this cluster. When a LoadBalancerRef is provided, the VSphereCluster.Status.Ready field will not be true until the referenced resource is Status.Ready and has a non-empty Status.Address value. DEPRECATED: will be removed in v1alpha4",
						MarkdownDescription: "LoadBalancerRef may be used to enable a control plane load balancer for this cluster. When a LoadBalancerRef is provided, the VSphereCluster.Status.Ready field will not be true until the referenced resource is Status.Ready and has a non-empty Status.Address value. DEPRECATED: will be removed in v1alpha4",
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

					"server": schema.StringAttribute{
						Description:         "Server is the address of the vSphere endpoint.",
						MarkdownDescription: "Server is the address of the vSphere endpoint.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"thumbprint": schema.StringAttribute{
						Description:         "Thumbprint is the colon-separated SHA-1 checksum of the given vCenter server's host certificate When provided, Insecure should not be set to true",
						MarkdownDescription: "Thumbprint is the colon-separated SHA-1 checksum of the given vCenter server's host certificate When provided, Insecure should not be set to true",
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

func (r *InfrastructureClusterXK8SIoVsphereClusterV1Alpha3Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedResourceDataError(request.ProviderData))
	}
}

func (r *InfrastructureClusterXK8SIoVsphereClusterV1Alpha3Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_infrastructure_cluster_x_k8s_io_v_sphere_cluster_v1alpha3")

	var model InfrastructureClusterXK8SIoVsphereClusterV1Alpha3ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("infrastructure.cluster.x-k8s.io/v1alpha3")
	model.Kind = pointer.String("VSphereCluster")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonMarshalError(err))
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "infrastructure.cluster.x-k8s.io", Version: "v1alpha3", Resource: "vsphereclusters"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.Append(utilities.PatchError(err))
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse InfrastructureClusterXK8SIoVsphereClusterV1Alpha3ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec
	if model.ForceConflicts.IsUnknown() {
		model.ForceConflicts = types.BoolNull()
	}
	if model.FieldManager.IsUnknown() {
		model.FieldManager = types.StringNull()
	}
	if model.DeletionPropagation.IsUnknown() {
		model.DeletionPropagation = types.StringNull()
	}
	if model.WaitForUpsert.IsUnknown() {
		model.WaitForUpsert = types.ListNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"jsonpath":      types.StringType,
				"value":         types.StringType,
				"timeout":       types.Int64Type,
				"poll_interval": types.Int64Type,
			},
		})
	}
	if model.WaitForDelete.IsUnknown() {
		model.WaitForDelete = types.ObjectNull(map[string]attr.Type{
			"timeout":       types.Int64Type,
			"poll_interval": types.Int64Type,
		})
	}

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *InfrastructureClusterXK8SIoVsphereClusterV1Alpha3Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_infrastructure_cluster_x_k8s_io_v_sphere_cluster_v1alpha3")

	var data InfrastructureClusterXK8SIoVsphereClusterV1Alpha3ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "infrastructure.cluster.x-k8s.io", Version: "v1alpha3", Resource: "vsphereclusters"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetNamespacedResourceError(err, data.Metadata.Name, data.Metadata.Namespace))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse InfrastructureClusterXK8SIoVsphereClusterV1Alpha3ResourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec
	if data.ForceConflicts.IsUnknown() {
		data.ForceConflicts = types.BoolNull()
	}
	if data.FieldManager.IsUnknown() {
		data.FieldManager = types.StringNull()
	}
	if data.DeletionPropagation.IsUnknown() {
		data.DeletionPropagation = types.StringNull()
	}
	if data.WaitForUpsert.IsUnknown() {
		data.WaitForUpsert = types.ListNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"jsonpath":      types.StringType,
				"value":         types.StringType,
				"timeout":       types.Int64Type,
				"poll_interval": types.Int64Type,
			},
		})
	}
	if data.WaitForDelete.IsUnknown() {
		data.WaitForDelete = types.ObjectNull(map[string]attr.Type{
			"timeout":       types.Int64Type,
			"poll_interval": types.Int64Type,
		})
	}

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *InfrastructureClusterXK8SIoVsphereClusterV1Alpha3Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_infrastructure_cluster_x_k8s_io_v_sphere_cluster_v1alpha3")

	var model InfrastructureClusterXK8SIoVsphereClusterV1Alpha3ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("infrastructure.cluster.x-k8s.io/v1alpha3")
	model.Kind = pointer.String("VSphereCluster")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonMarshalError(err))
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "infrastructure.cluster.x-k8s.io", Version: "v1alpha3", Resource: "vsphereclusters"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.Append(utilities.PatchError(err))
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse InfrastructureClusterXK8SIoVsphereClusterV1Alpha3ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *InfrastructureClusterXK8SIoVsphereClusterV1Alpha3Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_infrastructure_cluster_x_k8s_io_v_sphere_cluster_v1alpha3")

	var data InfrastructureClusterXK8SIoVsphereClusterV1Alpha3ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	deleteOptions := meta.DeleteOptions{}
	if !data.DeletionPropagation.IsNull() && !data.DeletionPropagation.IsUnknown() {
		deleteOptions.PropagationPolicy = utilities.MapDeletionPropagation(data.DeletionPropagation.ValueString())
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "infrastructure.cluster.x-k8s.io", Version: "v1alpha3", Resource: "vsphereclusters"}).
		Namespace(data.Metadata.Namespace).
		Delete(ctx, data.Metadata.Name, deleteOptions)
	if utilities.IsDeletionError(err) {
		response.Diagnostics.Append(utilities.DeleteError(err))
		return
	}

	if !data.WaitForDelete.IsNull() && !data.WaitForDelete.IsUnknown() {
		timeout := utilities.DetermineTimeout(data.WaitForDelete.Attributes())
		pollInterval := utilities.DeterminePollInterval(data.WaitForDelete.Attributes())

		startTime := time.Now()
		for {
			_, err := r.kubernetesClient.
				Resource(k8sSchema.GroupVersionResource{Group: "infrastructure.cluster.x-k8s.io", Version: "v1alpha3", Resource: "vsphereclusters"}).
				Namespace(data.Metadata.Namespace).
				Get(ctx, data.Metadata.Name, meta.GetOptions{})
			if utilities.IsNotFound(err) || timeout.Milliseconds() == 0 {
				break
			}
			if time.Now().After(startTime.Add(timeout)) {
				response.Diagnostics.Append(utilities.WaitTimeoutExceeded())
				return
			}
			time.Sleep(pollInterval)
		}
	}
}

func (r *InfrastructureClusterXK8SIoVsphereClusterV1Alpha3Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	idParts := strings.Split(request.ID, "/")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'namespace/name' Got: '%q'", request.ID),
		)
		return
	}

	namespace := idParts[0]
	name := idParts[1]
	tflog.Trace(ctx, "parsed import ID", map[string]interface{}{
		"namespace": namespace,
		"name":      name,
	})
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("namespace"), namespace)...)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("name"), name)...)
}
