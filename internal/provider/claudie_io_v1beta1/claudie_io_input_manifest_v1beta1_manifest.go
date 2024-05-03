/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package claudie_io_v1beta1

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
	_ datasource.DataSource = &ClaudieIoInputManifestV1Beta1Manifest{}
)

func NewClaudieIoInputManifestV1Beta1Manifest() datasource.DataSource {
	return &ClaudieIoInputManifestV1Beta1Manifest{}
}

type ClaudieIoInputManifestV1Beta1Manifest struct{}

type ClaudieIoInputManifestV1Beta1ManifestData struct {
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
		Kubernetes *struct {
			Clusters *[]struct {
				Name    *string `tfsdk:"name" json:"name,omitempty"`
				Network *string `tfsdk:"network" json:"network,omitempty"`
				Pools   *struct {
					Compute *[]string `tfsdk:"compute" json:"compute,omitempty"`
					Control *[]string `tfsdk:"control" json:"control,omitempty"`
				} `tfsdk:"pools" json:"pools,omitempty"`
				Version *string `tfsdk:"version" json:"version,omitempty"`
			} `tfsdk:"clusters" json:"clusters,omitempty"`
		} `tfsdk:"kubernetes" json:"kubernetes,omitempty"`
		LoadBalancers *struct {
			Clusters *[]struct {
				Dns *struct {
					DnsZone  *string `tfsdk:"dns_zone" json:"dnsZone,omitempty"`
					Hostname *string `tfsdk:"hostname" json:"hostname,omitempty"`
					Provider *string `tfsdk:"provider" json:"provider,omitempty"`
				} `tfsdk:"dns" json:"dns,omitempty"`
				Name        *string   `tfsdk:"name" json:"name,omitempty"`
				Pools       *[]string `tfsdk:"pools" json:"pools,omitempty"`
				Roles       *[]string `tfsdk:"roles" json:"roles,omitempty"`
				TargetedK8s *string   `tfsdk:"targeted_k8s" json:"targetedK8s,omitempty"`
			} `tfsdk:"clusters" json:"clusters,omitempty"`
			Roles *[]struct {
				Name        *string   `tfsdk:"name" json:"name,omitempty"`
				Port        *int64    `tfsdk:"port" json:"port,omitempty"`
				Protocol    *string   `tfsdk:"protocol" json:"protocol,omitempty"`
				TargetPools *[]string `tfsdk:"target_pools" json:"targetPools,omitempty"`
				TargetPort  *int64    `tfsdk:"target_port" json:"targetPort,omitempty"`
			} `tfsdk:"roles" json:"roles,omitempty"`
		} `tfsdk:"load_balancers" json:"loadBalancers,omitempty"`
		NodePools *struct {
			Dynamic *[]struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Autoscaler  *struct {
					Max *int64 `tfsdk:"max" json:"max,omitempty"`
					Min *int64 `tfsdk:"min" json:"min,omitempty"`
				} `tfsdk:"autoscaler" json:"autoscaler,omitempty"`
				Count       *int64             `tfsdk:"count" json:"count,omitempty"`
				Image       *string            `tfsdk:"image" json:"image,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				MachineSpec *struct {
					CpuCount *int64 `tfsdk:"cpu_count" json:"cpuCount,omitempty"`
					Memory   *int64 `tfsdk:"memory" json:"memory,omitempty"`
				} `tfsdk:"machine_spec" json:"machineSpec,omitempty"`
				Name         *string `tfsdk:"name" json:"name,omitempty"`
				ProviderSpec *struct {
					Name   *string `tfsdk:"name" json:"name,omitempty"`
					Region *string `tfsdk:"region" json:"region,omitempty"`
					Zone   *string `tfsdk:"zone" json:"zone,omitempty"`
				} `tfsdk:"provider_spec" json:"providerSpec,omitempty"`
				ServerType      *string `tfsdk:"server_type" json:"serverType,omitempty"`
				StorageDiskSize *int64  `tfsdk:"storage_disk_size" json:"storageDiskSize,omitempty"`
				Taints          *[]struct {
					Effect    *string `tfsdk:"effect" json:"effect,omitempty"`
					Key       *string `tfsdk:"key" json:"key,omitempty"`
					TimeAdded *string `tfsdk:"time_added" json:"timeAdded,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"taints" json:"taints,omitempty"`
			} `tfsdk:"dynamic" json:"dynamic,omitempty"`
			Static *[]struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				Name        *string            `tfsdk:"name" json:"name,omitempty"`
				Nodes       *[]struct {
					Endpoint  *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
					SecretRef *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
					Username *string `tfsdk:"username" json:"username,omitempty"`
				} `tfsdk:"nodes" json:"nodes,omitempty"`
				Taints *[]struct {
					Effect    *string `tfsdk:"effect" json:"effect,omitempty"`
					Key       *string `tfsdk:"key" json:"key,omitempty"`
					TimeAdded *string `tfsdk:"time_added" json:"timeAdded,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"taints" json:"taints,omitempty"`
			} `tfsdk:"static" json:"static,omitempty"`
		} `tfsdk:"node_pools" json:"nodePools,omitempty"`
		Providers *[]struct {
			Name         *string `tfsdk:"name" json:"name,omitempty"`
			ProviderType *string `tfsdk:"provider_type" json:"providerType,omitempty"`
			SecretRef    *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
		} `tfsdk:"providers" json:"providers,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ClaudieIoInputManifestV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_claudie_io_input_manifest_v1beta1_manifest"
}

func (r *ClaudieIoInputManifestV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "InputManifest is a definition of the user's infrastructure. It contains cloud provider specification, nodepool specification, Kubernetes and loadbalancer clusters.",
		MarkdownDescription: "InputManifest is a definition of the user's infrastructure. It contains cloud provider specification, nodepool specification, Kubernetes and loadbalancer clusters.",
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
				Description:         "Specification of the desired behaviour of the InputManifest",
				MarkdownDescription: "Specification of the desired behaviour of the InputManifest",
				Attributes: map[string]schema.Attribute{
					"kubernetes": schema.SingleNestedAttribute{
						Description:         "Kubernetes list of Kubernetes cluster this manifest will manage.",
						MarkdownDescription: "Kubernetes list of Kubernetes cluster this manifest will manage.",
						Attributes: map[string]schema.Attribute{
							"clusters": schema.ListNestedAttribute{
								Description:         "List of Kubernetes clusters Claudie will create.",
								MarkdownDescription: "List of Kubernetes clusters Claudie will create.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name of the Kubernetes cluster. Each cluster will have a random hash appended to the name, so the whole name will be of format <name>-<hash>.",
											MarkdownDescription: "Name of the Kubernetes cluster. Each cluster will have a random hash appended to the name, so the whole name will be of format <name>-<hash>.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"network": schema.StringAttribute{
											Description:         "Network range for the VPN of the cluster. The value should be defined in format A.B.C.D/mask.",
											MarkdownDescription: "Network range for the VPN of the cluster. The value should be defined in format A.B.C.D/mask.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"pools": schema.SingleNestedAttribute{
											Description:         "List of nodepool names this cluster will use.",
											MarkdownDescription: "List of nodepool names this cluster will use.",
											Attributes: map[string]schema.Attribute{
												"compute": schema.ListAttribute{
													Description:         "List of nodepool names, that will represent compute nodes.",
													MarkdownDescription: "List of nodepool names, that will represent compute nodes.",
													ElementType:         types.StringType,
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"control": schema.ListAttribute{
													Description:         "List of nodepool names, that will represent control plane nodes.",
													MarkdownDescription: "List of nodepool names, that will represent control plane nodes.",
													ElementType:         types.StringType,
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"version": schema.StringAttribute{
											Description:         "Version should be defined in format vX.Y. In terms of supported versions of Kubernetes, Claudie follows kubeone releases and their supported versions. The current kubeone version used in Claudie is 1.5. To see the list of supported versions, please refer to kubeone documentation. https://docs.kubermatic.com/kubeone/v1.5/architecture/compatibility/supported-versions/#supported-kubernetes-versions",
											MarkdownDescription: "Version should be defined in format vX.Y. In terms of supported versions of Kubernetes, Claudie follows kubeone releases and their supported versions. The current kubeone version used in Claudie is 1.5. To see the list of supported versions, please refer to kubeone documentation. https://docs.kubermatic.com/kubeone/v1.5/architecture/compatibility/supported-versions/#supported-kubernetes-versions",
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

					"load_balancers": schema.SingleNestedAttribute{
						Description:         "LoadBalancers list of loadbalancer clusters the Kubernetes clusters may use.",
						MarkdownDescription: "LoadBalancers list of loadbalancer clusters the Kubernetes clusters may use.",
						Attributes: map[string]schema.Attribute{
							"clusters": schema.ListNestedAttribute{
								Description:         "A list of load balancers clusters.",
								MarkdownDescription: "A list of load balancers clusters.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"dns": schema.SingleNestedAttribute{
											Description:         "Specification of the loadbalancer's DNS record.",
											MarkdownDescription: "Specification of the loadbalancer's DNS record.",
											Attributes: map[string]schema.Attribute{
												"dns_zone": schema.StringAttribute{
													Description:         "DNS zone inside of which the records will be created. GCP/AWS/OCI/Azure/Cloudflare/Hetzner DNS zone is accepted",
													MarkdownDescription: "DNS zone inside of which the records will be created. GCP/AWS/OCI/Azure/Cloudflare/Hetzner DNS zone is accepted",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"hostname": schema.StringAttribute{
													Description:         "Custom hostname for your A record. If left empty, the hostname will be a random hash.",
													MarkdownDescription: "Custom hostname for your A record. If left empty, the hostname will be a random hash.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"provider": schema.StringAttribute{
													Description:         "Name of provider to be used for creating an A record entry in defined DNS zone.",
													MarkdownDescription: "Name of provider to be used for creating an A record entry in defined DNS zone.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"name": schema.StringAttribute{
											Description:         "Name of the loadbalancer.",
											MarkdownDescription: "Name of the loadbalancer.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"pools": schema.ListAttribute{
											Description:         "List of nodepool names this loadbalancer will use. Remember, that nodepools defined in nodepools are only 'blueprints'. The actual nodepool will be created once referenced here.",
											MarkdownDescription: "List of nodepool names this loadbalancer will use. Remember, that nodepools defined in nodepools are only 'blueprints'. The actual nodepool will be created once referenced here.",
											ElementType:         types.StringType,
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"roles": schema.ListAttribute{
											Description:         "List of roles the loadbalancer uses.",
											MarkdownDescription: "List of roles the loadbalancer uses.",
											ElementType:         types.StringType,
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"targeted_k8s": schema.StringAttribute{
											Description:         "Name of the Kubernetes cluster targeted by this loadbalancer.",
											MarkdownDescription: "Name of the Kubernetes cluster targeted by this loadbalancer.",
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

							"roles": schema.ListNestedAttribute{
								Description:         "List of roles loadbalancers use to forward the traffic. Single role can be used in multiple loadbalancer clusters.",
								MarkdownDescription: "List of roles loadbalancers use to forward the traffic. Single role can be used in multiple loadbalancer clusters.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name of the role. Used as a reference in clusters.",
											MarkdownDescription: "Name of the role. Used as a reference in clusters.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"port": schema.Int64Attribute{
											Description:         "Port of the incoming traffic on the loadbalancer.",
											MarkdownDescription: "Port of the incoming traffic on the loadbalancer.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"protocol": schema.StringAttribute{
											Description:         "Protocol of the rule. Allowed values are: tcp, udp.",
											MarkdownDescription: "Protocol of the rule. Allowed values are: tcp, udp.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("tcp", "udp"),
											},
										},

										"target_pools": schema.ListAttribute{
											Description:         "Defines nodepools of the targeted K8s cluster, from which nodes will be used for loadbalancing.",
											MarkdownDescription: "Defines nodepools of the targeted K8s cluster, from which nodes will be used for loadbalancing.",
											ElementType:         types.StringType,
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"target_port": schema.Int64Attribute{
											Description:         "Port where loadbalancer forwards the traffic.",
											MarkdownDescription: "Port where loadbalancer forwards the traffic.",
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

					"node_pools": schema.SingleNestedAttribute{
						Description:         "NodePool is a map of dynamic nodepools and static nodepools which will be used to form kubernetes or loadbalancer clusters.",
						MarkdownDescription: "NodePool is a map of dynamic nodepools and static nodepools which will be used to form kubernetes or loadbalancer clusters.",
						Attributes: map[string]schema.Attribute{
							"dynamic": schema.ListNestedAttribute{
								Description:         "Dynamic nodepools define nodepools dynamically created by Claudie.",
								MarkdownDescription: "Dynamic nodepools define nodepools dynamically created by Claudie.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "User defined annotations for this nodepool.",
											MarkdownDescription: "User defined annotations for this nodepool.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"autoscaler": schema.SingleNestedAttribute{
											Description:         "Autoscaler configuration for this nodepool. Mutually exclusive with count.",
											MarkdownDescription: "Autoscaler configuration for this nodepool. Mutually exclusive with count.",
											Attributes: map[string]schema.Attribute{
												"max": schema.Int64Attribute{
													Description:         "Maximum number of nodes in nodepool.",
													MarkdownDescription: "Maximum number of nodes in nodepool.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"min": schema.Int64Attribute{
													Description:         "Minimum number of nodes in nodepool.",
													MarkdownDescription: "Minimum number of nodes in nodepool.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"count": schema.Int64Attribute{
											Description:         "Number of the nodes in the nodepool. Mutually exclusive with autoscaler.",
											MarkdownDescription: "Number of the nodes in the nodepool. Mutually exclusive with autoscaler.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"image": schema.StringAttribute{
											Description:         "OS image of the machine. Currently, only Ubuntu 22.04 AMD64 images are supported.",
											MarkdownDescription: "OS image of the machine. Currently, only Ubuntu 22.04 AMD64 images are supported.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"labels": schema.MapAttribute{
											Description:         "User defined labels for this nodepool.",
											MarkdownDescription: "User defined labels for this nodepool.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"machine_spec": schema.SingleNestedAttribute{
											Description:         "MachineSpec further describe the properties of the selected server type.",
											MarkdownDescription: "MachineSpec further describe the properties of the selected server type.",
											Attributes: map[string]schema.Attribute{
												"cpu_count": schema.Int64Attribute{
													Description:         "CpuCount specifies the number of CPU cores to be used.",
													MarkdownDescription: "CpuCount specifies the number of CPU cores to be used.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"memory": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"name": schema.StringAttribute{
											Description:         "Name of the nodepool. Each nodepool will have a random hash appended to the name, so the whole name will be of format <name>-<hash>.",
											MarkdownDescription: "Name of the nodepool. Each nodepool will have a random hash appended to the name, so the whole name will be of format <name>-<hash>.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"provider_spec": schema.SingleNestedAttribute{
											Description:         "Collection of provider data to be used while creating the nodepool.",
											MarkdownDescription: "Collection of provider data to be used while creating the nodepool.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of the provider instance specified in providers",
													MarkdownDescription: "Name of the provider instance specified in providers",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"region": schema.StringAttribute{
													Description:         "Region of the nodepool.",
													MarkdownDescription: "Region of the nodepool.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"zone": schema.StringAttribute{
													Description:         "Zone of the nodepool.",
													MarkdownDescription: "Zone of the nodepool.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"server_type": schema.StringAttribute{
											Description:         "Type of the machines in the nodepool. Currently, only AMD64 machines are supported.",
											MarkdownDescription: "Type of the machines in the nodepool. Currently, only AMD64 machines are supported.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"storage_disk_size": schema.Int64Attribute{
											Description:         "Size of the storage disk on the nodes in the nodepool in GB. The OS disk is created automatically with predefined size of 100GB for kubernetes nodes and 50GB for Loadbalancer nodes. The value must be either -1 (no disk is created), or >= 50. If no value is specified, 50 is used.",
											MarkdownDescription: "Size of the storage disk on the nodes in the nodepool in GB. The OS disk is created automatically with predefined size of 100GB for kubernetes nodes and 50GB for Loadbalancer nodes. The value must be either -1 (no disk is created), or >= 50. If no value is specified, 50 is used.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"taints": schema.ListNestedAttribute{
											Description:         "User defined taints for this nodepool.",
											MarkdownDescription: "User defined taints for this nodepool.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"effect": schema.StringAttribute{
														Description:         "Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.",
														MarkdownDescription: "Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"key": schema.StringAttribute{
														Description:         "Required. The taint key to be applied to a node.",
														MarkdownDescription: "Required. The taint key to be applied to a node.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"time_added": schema.StringAttribute{
														Description:         "TimeAdded represents the time at which the taint was added. It is only written for NoExecute taints.",
														MarkdownDescription: "TimeAdded represents the time at which the taint was added. It is only written for NoExecute taints.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															validators.DateTime64Validator(),
														},
													},

													"value": schema.StringAttribute{
														Description:         "The taint value corresponding to the taint key.",
														MarkdownDescription: "The taint value corresponding to the taint key.",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"static": schema.ListNestedAttribute{
								Description:         "Static nodepools define nodepools of already existing nodes.",
								MarkdownDescription: "Static nodepools define nodepools of already existing nodes.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"labels": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name of the nodepool.",
											MarkdownDescription: "Name of the nodepool.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"nodes": schema.ListNestedAttribute{
											Description:         "List of static nodes for a particular static nodepool.",
											MarkdownDescription: "List of static nodes for a particular static nodepool.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"endpoint": schema.StringAttribute{
														Description:         "Endpoint under which Claudie will access this node.",
														MarkdownDescription: "Endpoint under which Claudie will access this node.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"secret_ref": schema.SingleNestedAttribute{
														Description:         "Secret reference to the private key of the node.",
														MarkdownDescription: "Secret reference to the private key of the node.",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "name is unique within a namespace to reference a secret resource.",
																MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "namespace defines the space within which the secret name must be unique.",
																MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: true,
														Optional: false,
														Computed: false,
													},

													"username": schema.StringAttribute{
														Description:         "Username with root access. Used in SSH connection also.",
														MarkdownDescription: "Username with root access. Used in SSH connection also.",
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

										"taints": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"effect": schema.StringAttribute{
														Description:         "Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.",
														MarkdownDescription: "Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"key": schema.StringAttribute{
														Description:         "Required. The taint key to be applied to a node.",
														MarkdownDescription: "Required. The taint key to be applied to a node.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"time_added": schema.StringAttribute{
														Description:         "TimeAdded represents the time at which the taint was added. It is only written for NoExecute taints.",
														MarkdownDescription: "TimeAdded represents the time at which the taint was added. It is only written for NoExecute taints.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															validators.DateTime64Validator(),
														},
													},

													"value": schema.StringAttribute{
														Description:         "The taint value corresponding to the taint key.",
														MarkdownDescription: "The taint value corresponding to the taint key.",
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

					"providers": schema.ListNestedAttribute{
						Description:         "Providers list of defined cloud provider configuration that will be used while infrastructure provisioning.",
						MarkdownDescription: "Providers list of defined cloud provider configuration that will be used while infrastructure provisioning.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name is the name of the provider specification. It has to be unique across all providers.",
									MarkdownDescription: "Name is the name of the provider specification. It has to be unique across all providers.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(32),
									},
								},

								"provider_type": schema.StringAttribute{
									Description:         "ProviderType type of a provider. A list of available providers can be found at https://docs.claudie.io/v0.3.2/input-manifest/providers/aws/",
									MarkdownDescription: "ProviderType type of a provider. A list of available providers can be found at https://docs.claudie.io/v0.3.2/input-manifest/providers/aws/",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("gcp", "hetzner", "aws", "oci", "azure", "cloudflare", "hetznerdns", "genesiscloud"),
									},
								},

								"secret_ref": schema.SingleNestedAttribute{
									Description:         "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
									MarkdownDescription: "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "name is unique within a namespace to reference a secret resource.",
											MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"namespace": schema.StringAttribute{
											Description:         "namespace defines the space within which the secret name must be unique.",
											MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
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

func (r *ClaudieIoInputManifestV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_claudie_io_input_manifest_v1beta1_manifest")

	var model ClaudieIoInputManifestV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("claudie.io/v1beta1")
	model.Kind = pointer.String("InputManifest")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
