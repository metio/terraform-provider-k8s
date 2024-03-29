/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package beegfs_csi_netapp_com_v1

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
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &BeegfsCsiNetappComBeegfsDriverV1Manifest{}
)

func NewBeegfsCsiNetappComBeegfsDriverV1Manifest() datasource.DataSource {
	return &BeegfsCsiNetappComBeegfsDriverV1Manifest{}
}

type BeegfsCsiNetappComBeegfsDriverV1Manifest struct{}

type BeegfsCsiNetappComBeegfsDriverV1ManifestData struct {
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
		ContainerImageOverrides *struct {
			BeegfsCsiDriver *struct {
				Image *string `tfsdk:"image" json:"image,omitempty"`
				Tag   *string `tfsdk:"tag" json:"tag,omitempty"`
			} `tfsdk:"beegfs_csi_driver" json:"beegfsCsiDriver,omitempty"`
			CsiNodeDriverRegistrar *struct {
				Image *string `tfsdk:"image" json:"image,omitempty"`
				Tag   *string `tfsdk:"tag" json:"tag,omitempty"`
			} `tfsdk:"csi_node_driver_registrar" json:"csiNodeDriverRegistrar,omitempty"`
			CsiProvisioner *struct {
				Image *string `tfsdk:"image" json:"image,omitempty"`
				Tag   *string `tfsdk:"tag" json:"tag,omitempty"`
			} `tfsdk:"csi_provisioner" json:"csiProvisioner,omitempty"`
			LivenessProbe *struct {
				Image *string `tfsdk:"image" json:"image,omitempty"`
				Tag   *string `tfsdk:"tag" json:"tag,omitempty"`
			} `tfsdk:"liveness_probe" json:"livenessProbe,omitempty"`
		} `tfsdk:"container_image_overrides" json:"containerImageOverrides,omitempty"`
		ContainerResourceOverrides *struct {
			ControllerBeegfs *struct {
				Claims *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"controller_beegfs" json:"controllerBeegfs,omitempty"`
			ControllerCsiProvisioner *struct {
				Claims *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"controller_csi_provisioner" json:"controllerCsiProvisioner,omitempty"`
			NodeBeegfs *struct {
				Claims *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"node_beegfs" json:"nodeBeegfs,omitempty"`
			NodeDriverRegistrar *struct {
				Claims *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"node_driver_registrar" json:"nodeDriverRegistrar,omitempty"`
			NodeLivenessProbe *struct {
				Claims *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"node_liveness_probe" json:"nodeLivenessProbe,omitempty"`
		} `tfsdk:"container_resource_overrides" json:"containerResourceOverrides,omitempty"`
		LogLevel                      *int64 `tfsdk:"log_level" json:"logLevel,omitempty"`
		NodeAffinityControllerService *struct {
			PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
				Preference *struct {
					MatchExpressions *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
					MatchFields *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_fields" json:"matchFields,omitempty"`
				} `tfsdk:"preference" json:"preference,omitempty"`
				Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
			} `tfsdk:"preferred_during_scheduling_ignored_during_execution" json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
			RequiredDuringSchedulingIgnoredDuringExecution *struct {
				NodeSelectorTerms *[]struct {
					MatchExpressions *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
					MatchFields *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_fields" json:"matchFields,omitempty"`
				} `tfsdk:"node_selector_terms" json:"nodeSelectorTerms,omitempty"`
			} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
		} `tfsdk:"node_affinity_controller_service" json:"nodeAffinityControllerService,omitempty"`
		NodeAffinityNodeService *struct {
			PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
				Preference *struct {
					MatchExpressions *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
					MatchFields *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_fields" json:"matchFields,omitempty"`
				} `tfsdk:"preference" json:"preference,omitempty"`
				Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
			} `tfsdk:"preferred_during_scheduling_ignored_during_execution" json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
			RequiredDuringSchedulingIgnoredDuringExecution *struct {
				NodeSelectorTerms *[]struct {
					MatchExpressions *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
					MatchFields *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_fields" json:"matchFields,omitempty"`
				} `tfsdk:"node_selector_terms" json:"nodeSelectorTerms,omitempty"`
			} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
		} `tfsdk:"node_affinity_node_service" json:"nodeAffinityNodeService,omitempty"`
		PluginConfig *struct {
			Config *struct {
				BeegfsClientConf   *map[string]string `tfsdk:"beegfs_client_conf" json:"beegfsClientConf,omitempty"`
				ConnInterfaces     *[]string          `tfsdk:"conn_interfaces" json:"connInterfaces,omitempty"`
				ConnNetFilter      *[]string          `tfsdk:"conn_net_filter" json:"connNetFilter,omitempty"`
				ConnRDMAInterfaces *[]string          `tfsdk:"conn_rdma_interfaces" json:"connRDMAInterfaces,omitempty"`
				ConnTcpOnlyFilter  *[]string          `tfsdk:"conn_tcp_only_filter" json:"connTcpOnlyFilter,omitempty"`
			} `tfsdk:"config" json:"config,omitempty"`
			FileSystemSpecificConfigs *[]struct {
				Config *struct {
					BeegfsClientConf   *map[string]string `tfsdk:"beegfs_client_conf" json:"beegfsClientConf,omitempty"`
					ConnInterfaces     *[]string          `tfsdk:"conn_interfaces" json:"connInterfaces,omitempty"`
					ConnNetFilter      *[]string          `tfsdk:"conn_net_filter" json:"connNetFilter,omitempty"`
					ConnRDMAInterfaces *[]string          `tfsdk:"conn_rdma_interfaces" json:"connRDMAInterfaces,omitempty"`
					ConnTcpOnlyFilter  *[]string          `tfsdk:"conn_tcp_only_filter" json:"connTcpOnlyFilter,omitempty"`
				} `tfsdk:"config" json:"config,omitempty"`
				SysMgmtdHost *string `tfsdk:"sys_mgmtd_host" json:"sysMgmtdHost,omitempty"`
			} `tfsdk:"file_system_specific_configs" json:"fileSystemSpecificConfigs,omitempty"`
			NodeSpecificConfigs *[]struct {
				Config *struct {
					BeegfsClientConf   *map[string]string `tfsdk:"beegfs_client_conf" json:"beegfsClientConf,omitempty"`
					ConnInterfaces     *[]string          `tfsdk:"conn_interfaces" json:"connInterfaces,omitempty"`
					ConnNetFilter      *[]string          `tfsdk:"conn_net_filter" json:"connNetFilter,omitempty"`
					ConnRDMAInterfaces *[]string          `tfsdk:"conn_rdma_interfaces" json:"connRDMAInterfaces,omitempty"`
					ConnTcpOnlyFilter  *[]string          `tfsdk:"conn_tcp_only_filter" json:"connTcpOnlyFilter,omitempty"`
				} `tfsdk:"config" json:"config,omitempty"`
				FileSystemSpecificConfigs *[]struct {
					Config *struct {
						BeegfsClientConf   *map[string]string `tfsdk:"beegfs_client_conf" json:"beegfsClientConf,omitempty"`
						ConnInterfaces     *[]string          `tfsdk:"conn_interfaces" json:"connInterfaces,omitempty"`
						ConnNetFilter      *[]string          `tfsdk:"conn_net_filter" json:"connNetFilter,omitempty"`
						ConnRDMAInterfaces *[]string          `tfsdk:"conn_rdma_interfaces" json:"connRDMAInterfaces,omitempty"`
						ConnTcpOnlyFilter  *[]string          `tfsdk:"conn_tcp_only_filter" json:"connTcpOnlyFilter,omitempty"`
					} `tfsdk:"config" json:"config,omitempty"`
					SysMgmtdHost *string `tfsdk:"sys_mgmtd_host" json:"sysMgmtdHost,omitempty"`
				} `tfsdk:"file_system_specific_configs" json:"fileSystemSpecificConfigs,omitempty"`
				NodeList *[]string `tfsdk:"node_list" json:"nodeList,omitempty"`
			} `tfsdk:"node_specific_configs" json:"nodeSpecificConfigs,omitempty"`
		} `tfsdk:"plugin_config" json:"pluginConfig,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *BeegfsCsiNetappComBeegfsDriverV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_beegfs_csi_netapp_com_beegfs_driver_v1_manifest"
}

func (r *BeegfsCsiNetappComBeegfsDriverV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Deploys the BeeGFS CSI driver",
		MarkdownDescription: "Deploys the BeeGFS CSI driver",
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
				Description:         "BeegfsDriverSpec defines the desired state of BeegfsDriver",
				MarkdownDescription: "BeegfsDriverSpec defines the desired state of BeegfsDriver",
				Attributes: map[string]schema.Attribute{
					"container_image_overrides": schema.SingleNestedAttribute{
						Description:         "A structure that allows for default container images and tags to be overridden. Use it in air-gapped networks,networks with private registry mirrors, or to pin a particular container version. Unless otherwise noted, versionsother than the default are not supported.",
						MarkdownDescription: "A structure that allows for default container images and tags to be overridden. Use it in air-gapped networks,networks with private registry mirrors, or to pin a particular container version. Unless otherwise noted, versionsother than the default are not supported.",
						Attributes: map[string]schema.Attribute{
							"beegfs_csi_driver": schema.SingleNestedAttribute{
								Description:         "Defaults to ghcr.io/thinkparq/beegfs-csi-driver:<the operator version>.",
								MarkdownDescription: "Defaults to ghcr.io/thinkparq/beegfs-csi-driver:<the operator version>.",
								Attributes: map[string]schema.Attribute{
									"image": schema.StringAttribute{
										Description:         "A combination of registry and image (e.g. registry.k8s.io/csi-provisioner or ghcr.io/thinkparq/beegfs-csi-driver).",
										MarkdownDescription: "A combination of registry and image (e.g. registry.k8s.io/csi-provisioner or ghcr.io/thinkparq/beegfs-csi-driver).",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tag": schema.StringAttribute{
										Description:         "A tag (e.g. v2.2.2 or latest).",
										MarkdownDescription: "A tag (e.g. v2.2.2 or latest).",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"csi_node_driver_registrar": schema.SingleNestedAttribute{
								Description:         "Defaults to registry.k8s.io/sig-storage/csi-node-driver-registrar:<the most current version at operator release>.",
								MarkdownDescription: "Defaults to registry.k8s.io/sig-storage/csi-node-driver-registrar:<the most current version at operator release>.",
								Attributes: map[string]schema.Attribute{
									"image": schema.StringAttribute{
										Description:         "A combination of registry and image (e.g. registry.k8s.io/csi-provisioner or ghcr.io/thinkparq/beegfs-csi-driver).",
										MarkdownDescription: "A combination of registry and image (e.g. registry.k8s.io/csi-provisioner or ghcr.io/thinkparq/beegfs-csi-driver).",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tag": schema.StringAttribute{
										Description:         "A tag (e.g. v2.2.2 or latest).",
										MarkdownDescription: "A tag (e.g. v2.2.2 or latest).",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"csi_provisioner": schema.SingleNestedAttribute{
								Description:         "Defaults to registry.k8s.io/sig-storage/csi-provisioner:<the most current version at operator release>.",
								MarkdownDescription: "Defaults to registry.k8s.io/sig-storage/csi-provisioner:<the most current version at operator release>.",
								Attributes: map[string]schema.Attribute{
									"image": schema.StringAttribute{
										Description:         "A combination of registry and image (e.g. registry.k8s.io/csi-provisioner or ghcr.io/thinkparq/beegfs-csi-driver).",
										MarkdownDescription: "A combination of registry and image (e.g. registry.k8s.io/csi-provisioner or ghcr.io/thinkparq/beegfs-csi-driver).",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tag": schema.StringAttribute{
										Description:         "A tag (e.g. v2.2.2 or latest).",
										MarkdownDescription: "A tag (e.g. v2.2.2 or latest).",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"liveness_probe": schema.SingleNestedAttribute{
								Description:         "Defaults to registry.k8s.io/sig-storage/livenessprobe:<the most current version at operator release>.",
								MarkdownDescription: "Defaults to registry.k8s.io/sig-storage/livenessprobe:<the most current version at operator release>.",
								Attributes: map[string]schema.Attribute{
									"image": schema.StringAttribute{
										Description:         "A combination of registry and image (e.g. registry.k8s.io/csi-provisioner or ghcr.io/thinkparq/beegfs-csi-driver).",
										MarkdownDescription: "A combination of registry and image (e.g. registry.k8s.io/csi-provisioner or ghcr.io/thinkparq/beegfs-csi-driver).",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tag": schema.StringAttribute{
										Description:         "A tag (e.g. v2.2.2 or latest).",
										MarkdownDescription: "A tag (e.g. v2.2.2 or latest).",
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

					"container_resource_overrides": schema.SingleNestedAttribute{
						Description:         "The ContainerResourceOverrides allow for customization of the container resource limits and requests.Each container has default requests and limits for both cpu and memory resources. Only explicitly definedoverrides will be applied, otherwise the default values will be used. For example, if the cpu limit for thecontroller's beegfs container is the only resource with an override set, only the controller's beegfs containercpu limit setting will be overridden. Every other value will use the default setting. Storage resources are notused by the BeeGFS CSI driver. Any storage resource values configured will be ignored.",
						MarkdownDescription: "The ContainerResourceOverrides allow for customization of the container resource limits and requests.Each container has default requests and limits for both cpu and memory resources. Only explicitly definedoverrides will be applied, otherwise the default values will be used. For example, if the cpu limit for thecontroller's beegfs container is the only resource with an override set, only the controller's beegfs containercpu limit setting will be overridden. Every other value will use the default setting. Storage resources are notused by the BeeGFS CSI driver. Any storage resource values configured will be ignored.",
						Attributes: map[string]schema.Attribute{
							"controller_beegfs": schema.SingleNestedAttribute{
								Description:         "The resource specifications for the beegfs container of the BeeGFS driver controller pod.The default values for requests are (cpu: 100m, memory: 16Mi).The default values for limits are (cpu: None, memory: 256Mi).",
								MarkdownDescription: "The resource specifications for the beegfs container of the BeeGFS driver controller pod.The default values for requests are (cpu: 100m, memory: 16Mi).The default values for limits are (cpu: None, memory: 256Mi).",
								Attributes: map[string]schema.Attribute{
									"claims": schema.ListNestedAttribute{
										Description:         "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
										MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
													MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
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

									"limits": schema.MapAttribute{
										Description:         "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"requests": schema.MapAttribute{
										Description:         "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

							"controller_csi_provisioner": schema.SingleNestedAttribute{
								Description:         "The resource specifications for the csi-provisioner container of the BeeGFS driver controller pod.The default values for requests are (cpu: 80m, memory: 24Mi)The default values for limits are (cpu: None, memory 256Mi)",
								MarkdownDescription: "The resource specifications for the csi-provisioner container of the BeeGFS driver controller pod.The default values for requests are (cpu: 80m, memory: 24Mi)The default values for limits are (cpu: None, memory 256Mi)",
								Attributes: map[string]schema.Attribute{
									"claims": schema.ListNestedAttribute{
										Description:         "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
										MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
													MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
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

									"limits": schema.MapAttribute{
										Description:         "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"requests": schema.MapAttribute{
										Description:         "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

							"node_beegfs": schema.SingleNestedAttribute{
								Description:         "The resource specifications for the beegfs container of the BeeGFS driver node pod.The default values for requests are (cpu: 100m, memory: 20Mi)The default values for limits are (cpu: None, memory: 128Mi)",
								MarkdownDescription: "The resource specifications for the beegfs container of the BeeGFS driver node pod.The default values for requests are (cpu: 100m, memory: 20Mi)The default values for limits are (cpu: None, memory: 128Mi)",
								Attributes: map[string]schema.Attribute{
									"claims": schema.ListNestedAttribute{
										Description:         "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
										MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
													MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
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

									"limits": schema.MapAttribute{
										Description:         "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"requests": schema.MapAttribute{
										Description:         "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

							"node_driver_registrar": schema.SingleNestedAttribute{
								Description:         "The resource specifications for the node-driver-registrar container of the BeeGFS driver node pod.The default values for requests are (cpu: 80m, memory: 10Mi)The default values for limits are (cpu: None, memory 128Mi)",
								MarkdownDescription: "The resource specifications for the node-driver-registrar container of the BeeGFS driver node pod.The default values for requests are (cpu: 80m, memory: 10Mi)The default values for limits are (cpu: None, memory 128Mi)",
								Attributes: map[string]schema.Attribute{
									"claims": schema.ListNestedAttribute{
										Description:         "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
										MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
													MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
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

									"limits": schema.MapAttribute{
										Description:         "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"requests": schema.MapAttribute{
										Description:         "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

							"node_liveness_probe": schema.SingleNestedAttribute{
								Description:         "The resource specifications for the liveness-probe container of the BeeGFS driver node pod.The default values for requests are (cpu: 60m, memory: 20Mi)The default values for limits are (cpu: None, memory: 128Mi)",
								MarkdownDescription: "The resource specifications for the liveness-probe container of the BeeGFS driver node pod.The default values for requests are (cpu: 60m, memory: 20Mi)The default values for limits are (cpu: None, memory: 128Mi)",
								Attributes: map[string]schema.Attribute{
									"claims": schema.ListNestedAttribute{
										Description:         "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
										MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
													MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
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

									"limits": schema.MapAttribute{
										Description:         "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"requests": schema.MapAttribute{
										Description:         "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

					"log_level": schema.Int64Attribute{
						Description:         "The logging level of deployed containers expressed as an integer from 0 (low detail) to 5 (high detail). 0only logs errors. 3 logs most RPC requests/responses and some detail about driver actions. 5 logs all RPCrequests/responses, including redundant/frequently occurring ones. Empty defaults to level 3.",
						MarkdownDescription: "The logging level of deployed containers expressed as an integer from 0 (low detail) to 5 (high detail). 0only logs errors. 3 logs most RPC requests/responses and some detail about driver actions. 5 logs all RPCrequests/responses, including redundant/frequently occurring ones. Empty defaults to level 3.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
							int64validator.AtMost(5),
						},
					},

					"node_affinity_controller_service": schema.SingleNestedAttribute{
						Description:         "The controller service consists of a single Pod. It preferably runs on an infrastructure/master node, but therunning node must have the beegfs-utils and beegfs-client packages installed. E.g.'preferred: node-role.kubernetes.io/master Exists' and/or 'required: node.openshift.io/os_id NotIn rhcos'.",
						MarkdownDescription: "The controller service consists of a single Pod. It preferably runs on an infrastructure/master node, but therunning node must have the beegfs-utils and beegfs-client packages installed. E.g.'preferred: node-role.kubernetes.io/master Exists' and/or 'required: node.openshift.io/os_id NotIn rhcos'.",
						Attributes: map[string]schema.Attribute{
							"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
								Description:         "The scheduler will prefer to schedule pods to nodes that satisfythe affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node matches the corresponding matchExpressions; thenode(s) with the highest sum are the most preferred.",
								MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfythe affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node matches the corresponding matchExpressions; thenode(s) with the highest sum are the most preferred.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"preference": schema.SingleNestedAttribute{
											Description:         "A node selector term, associated with the corresponding weight.",
											MarkdownDescription: "A node selector term, associated with the corresponding weight.",
											Attributes: map[string]schema.Attribute{
												"match_expressions": schema.ListNestedAttribute{
													Description:         "A list of node selector requirements by node's labels.",
													MarkdownDescription: "A list of node selector requirements by node's labels.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The label key that the selector applies to.",
																MarkdownDescription: "The label key that the selector applies to.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"operator": schema.StringAttribute{
																Description:         "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																MarkdownDescription: "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"values": schema.ListAttribute{
																Description:         "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
																MarkdownDescription: "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
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

												"match_fields": schema.ListNestedAttribute{
													Description:         "A list of node selector requirements by node's fields.",
													MarkdownDescription: "A list of node selector requirements by node's fields.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The label key that the selector applies to.",
																MarkdownDescription: "The label key that the selector applies to.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"operator": schema.StringAttribute{
																Description:         "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																MarkdownDescription: "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"values": schema.ListAttribute{
																Description:         "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
																MarkdownDescription: "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"weight": schema.Int64Attribute{
											Description:         "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
											MarkdownDescription: "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
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

							"required_during_scheduling_ignored_during_execution": schema.SingleNestedAttribute{
								Description:         "If the affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to an update), the systemmay or may not try to eventually evict the pod from its node.",
								MarkdownDescription: "If the affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to an update), the systemmay or may not try to eventually evict the pod from its node.",
								Attributes: map[string]schema.Attribute{
									"node_selector_terms": schema.ListNestedAttribute{
										Description:         "Required. A list of node selector terms. The terms are ORed.",
										MarkdownDescription: "Required. A list of node selector terms. The terms are ORed.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"match_expressions": schema.ListNestedAttribute{
													Description:         "A list of node selector requirements by node's labels.",
													MarkdownDescription: "A list of node selector requirements by node's labels.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The label key that the selector applies to.",
																MarkdownDescription: "The label key that the selector applies to.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"operator": schema.StringAttribute{
																Description:         "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																MarkdownDescription: "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"values": schema.ListAttribute{
																Description:         "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
																MarkdownDescription: "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
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

												"match_fields": schema.ListNestedAttribute{
													Description:         "A list of node selector requirements by node's fields.",
													MarkdownDescription: "A list of node selector requirements by node's fields.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The label key that the selector applies to.",
																MarkdownDescription: "The label key that the selector applies to.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"operator": schema.StringAttribute{
																Description:         "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																MarkdownDescription: "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"values": schema.ListAttribute{
																Description:         "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
																MarkdownDescription: "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"node_affinity_node_service": schema.SingleNestedAttribute{
						Description:         "The node service consists of one Pod running on each eligible node. It runs on every node expected to host aworkload that requires BeeGFS. Running nodes must have the beegfs-utils and beegfs-client packages installed.E.g. 'required: node.openshift.io/os_id NotIn rhcos'.",
						MarkdownDescription: "The node service consists of one Pod running on each eligible node. It runs on every node expected to host aworkload that requires BeeGFS. Running nodes must have the beegfs-utils and beegfs-client packages installed.E.g. 'required: node.openshift.io/os_id NotIn rhcos'.",
						Attributes: map[string]schema.Attribute{
							"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
								Description:         "The scheduler will prefer to schedule pods to nodes that satisfythe affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node matches the corresponding matchExpressions; thenode(s) with the highest sum are the most preferred.",
								MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfythe affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node matches the corresponding matchExpressions; thenode(s) with the highest sum are the most preferred.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"preference": schema.SingleNestedAttribute{
											Description:         "A node selector term, associated with the corresponding weight.",
											MarkdownDescription: "A node selector term, associated with the corresponding weight.",
											Attributes: map[string]schema.Attribute{
												"match_expressions": schema.ListNestedAttribute{
													Description:         "A list of node selector requirements by node's labels.",
													MarkdownDescription: "A list of node selector requirements by node's labels.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The label key that the selector applies to.",
																MarkdownDescription: "The label key that the selector applies to.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"operator": schema.StringAttribute{
																Description:         "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																MarkdownDescription: "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"values": schema.ListAttribute{
																Description:         "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
																MarkdownDescription: "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
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

												"match_fields": schema.ListNestedAttribute{
													Description:         "A list of node selector requirements by node's fields.",
													MarkdownDescription: "A list of node selector requirements by node's fields.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The label key that the selector applies to.",
																MarkdownDescription: "The label key that the selector applies to.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"operator": schema.StringAttribute{
																Description:         "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																MarkdownDescription: "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"values": schema.ListAttribute{
																Description:         "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
																MarkdownDescription: "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"weight": schema.Int64Attribute{
											Description:         "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
											MarkdownDescription: "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
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

							"required_during_scheduling_ignored_during_execution": schema.SingleNestedAttribute{
								Description:         "If the affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to an update), the systemmay or may not try to eventually evict the pod from its node.",
								MarkdownDescription: "If the affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to an update), the systemmay or may not try to eventually evict the pod from its node.",
								Attributes: map[string]schema.Attribute{
									"node_selector_terms": schema.ListNestedAttribute{
										Description:         "Required. A list of node selector terms. The terms are ORed.",
										MarkdownDescription: "Required. A list of node selector terms. The terms are ORed.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"match_expressions": schema.ListNestedAttribute{
													Description:         "A list of node selector requirements by node's labels.",
													MarkdownDescription: "A list of node selector requirements by node's labels.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The label key that the selector applies to.",
																MarkdownDescription: "The label key that the selector applies to.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"operator": schema.StringAttribute{
																Description:         "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																MarkdownDescription: "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"values": schema.ListAttribute{
																Description:         "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
																MarkdownDescription: "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
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

												"match_fields": schema.ListNestedAttribute{
													Description:         "A list of node selector requirements by node's fields.",
													MarkdownDescription: "A list of node selector requirements by node's fields.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The label key that the selector applies to.",
																MarkdownDescription: "The label key that the selector applies to.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"operator": schema.StringAttribute{
																Description:         "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																MarkdownDescription: "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"values": schema.ListAttribute{
																Description:         "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
																MarkdownDescription: "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"plugin_config": schema.SingleNestedAttribute{
						Description:         "The top level configuration structure containing default configuration (applied to all file systems on all nodes),file system specific configuration, and node specific configuration. Fields from node and file system specificconfigurations override fields from the default configuration. Often not required.",
						MarkdownDescription: "The top level configuration structure containing default configuration (applied to all file systems on all nodes),file system specific configuration, and node specific configuration. Fields from node and file system specificconfigurations override fields from the default configuration. Often not required.",
						Attributes: map[string]schema.Attribute{
							"config": schema.SingleNestedAttribute{
								Description:         "The primary configuration structure containing all of the custom configuration (beegfs-client.conf keys/values andadditional CSI driver specific fields) associated with a single BeeGFS file system except for sysMgmtdHost, which isspecified elsewhere. WARNING: This structure includes a beegfsClientConf field. This field may not be rendered inform view by OpenShift or other graphical interfaces, but it can be critical in some environments. Add or modify itin YAML view.",
								MarkdownDescription: "The primary configuration structure containing all of the custom configuration (beegfs-client.conf keys/values andadditional CSI driver specific fields) associated with a single BeeGFS file system except for sysMgmtdHost, which isspecified elsewhere. WARNING: This structure includes a beegfsClientConf field. This field may not be rendered inform view by OpenShift or other graphical interfaces, but it can be critical in some environments. Add or modify itin YAML view.",
								Attributes: map[string]schema.Attribute{
									"beegfs_client_conf": schema.MapAttribute{
										Description:         "A map of additional key value pairs matching key value pairs in the beegfs-client.conf file. Seebeegfs-client.conf for more details. Values MUST be specified as strings, even if they appear to be integers orbooleans (e.g. '8000', not 8000 and 'true', not true).",
										MarkdownDescription: "A map of additional key value pairs matching key value pairs in the beegfs-client.conf file. Seebeegfs-client.conf for more details. Values MUST be specified as strings, even if they appear to be integers orbooleans (e.g. '8000', not 8000 and 'true', not true).",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"conn_interfaces": schema.ListAttribute{
										Description:         "A list of interfaces the BeeGFS client service can communicate over (e.g. 'ib0' or 'eth0'). Often not required.See beegfs-client.conf for more details.",
										MarkdownDescription: "A list of interfaces the BeeGFS client service can communicate over (e.g. 'ib0' or 'eth0'). Often not required.See beegfs-client.conf for more details.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"conn_net_filter": schema.ListAttribute{
										Description:         "A list of subnets the BeeGFS client service can use for outgoing communication (e.g. '10.10.10.10/24'). Oftennot required. See beegfs-client.conf for more details.",
										MarkdownDescription: "A list of subnets the BeeGFS client service can use for outgoing communication (e.g. '10.10.10.10/24'). Oftennot required. See beegfs-client.conf for more details.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"conn_rdma_interfaces": schema.ListAttribute{
										Description:         "A list of interfaces the BeeGFS client will use for outbound RDMA connections. This is used in supportof the BeeGFS multi-rail feature. This feature does not depend on or use the connInterfaces parameter.This feature requires the BeeGFS client version 7.3.0 or later.",
										MarkdownDescription: "A list of interfaces the BeeGFS client will use for outbound RDMA connections. This is used in supportof the BeeGFS multi-rail feature. This feature does not depend on or use the connInterfaces parameter.This feature requires the BeeGFS client version 7.3.0 or later.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"conn_tcp_only_filter": schema.ListAttribute{
										Description:         "A list of subnets in which RDMA communication cannot or should not be established (e.g. '10.10.10.11/24').Often not required. See beegfs-client.conf for more details.",
										MarkdownDescription: "A list of subnets in which RDMA communication cannot or should not be established (e.g. '10.10.10.11/24').Often not required. See beegfs-client.conf for more details.",
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

							"file_system_specific_configs": schema.ListNestedAttribute{
								Description:         "A list of file system specific configurations that override the default configuration for specific file systems.",
								MarkdownDescription: "A list of file system specific configurations that override the default configuration for specific file systems.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"config": schema.SingleNestedAttribute{
											Description:         "The primary configuration structure containing all of the custom configuration (beegfs-client.conf keys/values andadditional CSI driver specific fields) associated with a single BeeGFS file system except for sysMgmtdHost, which isspecified elsewhere. WARNING: This structure includes a beegfsClientConf field. This field may not be rendered inform view by OpenShift or other graphical interfaces, but it can be critical in some environments. Add or modify itin YAML view.",
											MarkdownDescription: "The primary configuration structure containing all of the custom configuration (beegfs-client.conf keys/values andadditional CSI driver specific fields) associated with a single BeeGFS file system except for sysMgmtdHost, which isspecified elsewhere. WARNING: This structure includes a beegfsClientConf field. This field may not be rendered inform view by OpenShift or other graphical interfaces, but it can be critical in some environments. Add or modify itin YAML view.",
											Attributes: map[string]schema.Attribute{
												"beegfs_client_conf": schema.MapAttribute{
													Description:         "A map of additional key value pairs matching key value pairs in the beegfs-client.conf file. Seebeegfs-client.conf for more details. Values MUST be specified as strings, even if they appear to be integers orbooleans (e.g. '8000', not 8000 and 'true', not true).",
													MarkdownDescription: "A map of additional key value pairs matching key value pairs in the beegfs-client.conf file. Seebeegfs-client.conf for more details. Values MUST be specified as strings, even if they appear to be integers orbooleans (e.g. '8000', not 8000 and 'true', not true).",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"conn_interfaces": schema.ListAttribute{
													Description:         "A list of interfaces the BeeGFS client service can communicate over (e.g. 'ib0' or 'eth0'). Often not required.See beegfs-client.conf for more details.",
													MarkdownDescription: "A list of interfaces the BeeGFS client service can communicate over (e.g. 'ib0' or 'eth0'). Often not required.See beegfs-client.conf for more details.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"conn_net_filter": schema.ListAttribute{
													Description:         "A list of subnets the BeeGFS client service can use for outgoing communication (e.g. '10.10.10.10/24'). Oftennot required. See beegfs-client.conf for more details.",
													MarkdownDescription: "A list of subnets the BeeGFS client service can use for outgoing communication (e.g. '10.10.10.10/24'). Oftennot required. See beegfs-client.conf for more details.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"conn_rdma_interfaces": schema.ListAttribute{
													Description:         "A list of interfaces the BeeGFS client will use for outbound RDMA connections. This is used in supportof the BeeGFS multi-rail feature. This feature does not depend on or use the connInterfaces parameter.This feature requires the BeeGFS client version 7.3.0 or later.",
													MarkdownDescription: "A list of interfaces the BeeGFS client will use for outbound RDMA connections. This is used in supportof the BeeGFS multi-rail feature. This feature does not depend on or use the connInterfaces parameter.This feature requires the BeeGFS client version 7.3.0 or later.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"conn_tcp_only_filter": schema.ListAttribute{
													Description:         "A list of subnets in which RDMA communication cannot or should not be established (e.g. '10.10.10.11/24').Often not required. See beegfs-client.conf for more details.",
													MarkdownDescription: "A list of subnets in which RDMA communication cannot or should not be established (e.g. '10.10.10.11/24').Often not required. See beegfs-client.conf for more details.",
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

										"sys_mgmtd_host": schema.StringAttribute{
											Description:         "The sysMgmtdHost used by the BeeGFS client service to make initial contact with the BeeGFS mgmtd service.",
											MarkdownDescription: "The sysMgmtdHost used by the BeeGFS client service to make initial contact with the BeeGFS mgmtd service.",
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

							"node_specific_configs": schema.ListNestedAttribute{
								Description:         "A list of node specific configurations that override file system specific configurations and the defaultconfiguration on specific nodes.",
								MarkdownDescription: "A list of node specific configurations that override file system specific configurations and the defaultconfiguration on specific nodes.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"config": schema.SingleNestedAttribute{
											Description:         "The primary configuration structure containing all of the custom configuration (beegfs-client.conf keys/values andadditional CSI driver specific fields) associated with a single BeeGFS file system except for sysMgmtdHost, which isspecified elsewhere. WARNING: This structure includes a beegfsClientConf field. This field may not be rendered inform view by OpenShift or other graphical interfaces, but it can be critical in some environments. Add or modify itin YAML view.",
											MarkdownDescription: "The primary configuration structure containing all of the custom configuration (beegfs-client.conf keys/values andadditional CSI driver specific fields) associated with a single BeeGFS file system except for sysMgmtdHost, which isspecified elsewhere. WARNING: This structure includes a beegfsClientConf field. This field may not be rendered inform view by OpenShift or other graphical interfaces, but it can be critical in some environments. Add or modify itin YAML view.",
											Attributes: map[string]schema.Attribute{
												"beegfs_client_conf": schema.MapAttribute{
													Description:         "A map of additional key value pairs matching key value pairs in the beegfs-client.conf file. Seebeegfs-client.conf for more details. Values MUST be specified as strings, even if they appear to be integers orbooleans (e.g. '8000', not 8000 and 'true', not true).",
													MarkdownDescription: "A map of additional key value pairs matching key value pairs in the beegfs-client.conf file. Seebeegfs-client.conf for more details. Values MUST be specified as strings, even if they appear to be integers orbooleans (e.g. '8000', not 8000 and 'true', not true).",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"conn_interfaces": schema.ListAttribute{
													Description:         "A list of interfaces the BeeGFS client service can communicate over (e.g. 'ib0' or 'eth0'). Often not required.See beegfs-client.conf for more details.",
													MarkdownDescription: "A list of interfaces the BeeGFS client service can communicate over (e.g. 'ib0' or 'eth0'). Often not required.See beegfs-client.conf for more details.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"conn_net_filter": schema.ListAttribute{
													Description:         "A list of subnets the BeeGFS client service can use for outgoing communication (e.g. '10.10.10.10/24'). Oftennot required. See beegfs-client.conf for more details.",
													MarkdownDescription: "A list of subnets the BeeGFS client service can use for outgoing communication (e.g. '10.10.10.10/24'). Oftennot required. See beegfs-client.conf for more details.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"conn_rdma_interfaces": schema.ListAttribute{
													Description:         "A list of interfaces the BeeGFS client will use for outbound RDMA connections. This is used in supportof the BeeGFS multi-rail feature. This feature does not depend on or use the connInterfaces parameter.This feature requires the BeeGFS client version 7.3.0 or later.",
													MarkdownDescription: "A list of interfaces the BeeGFS client will use for outbound RDMA connections. This is used in supportof the BeeGFS multi-rail feature. This feature does not depend on or use the connInterfaces parameter.This feature requires the BeeGFS client version 7.3.0 or later.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"conn_tcp_only_filter": schema.ListAttribute{
													Description:         "A list of subnets in which RDMA communication cannot or should not be established (e.g. '10.10.10.11/24').Often not required. See beegfs-client.conf for more details.",
													MarkdownDescription: "A list of subnets in which RDMA communication cannot or should not be established (e.g. '10.10.10.11/24').Often not required. See beegfs-client.conf for more details.",
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

										"file_system_specific_configs": schema.ListNestedAttribute{
											Description:         "A list of file system specific configurations that override the default configuration for specific file systemson these nodes.",
											MarkdownDescription: "A list of file system specific configurations that override the default configuration for specific file systemson these nodes.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"config": schema.SingleNestedAttribute{
														Description:         "The primary configuration structure containing all of the custom configuration (beegfs-client.conf keys/values andadditional CSI driver specific fields) associated with a single BeeGFS file system except for sysMgmtdHost, which isspecified elsewhere. WARNING: This structure includes a beegfsClientConf field. This field may not be rendered inform view by OpenShift or other graphical interfaces, but it can be critical in some environments. Add or modify itin YAML view.",
														MarkdownDescription: "The primary configuration structure containing all of the custom configuration (beegfs-client.conf keys/values andadditional CSI driver specific fields) associated with a single BeeGFS file system except for sysMgmtdHost, which isspecified elsewhere. WARNING: This structure includes a beegfsClientConf field. This field may not be rendered inform view by OpenShift or other graphical interfaces, but it can be critical in some environments. Add or modify itin YAML view.",
														Attributes: map[string]schema.Attribute{
															"beegfs_client_conf": schema.MapAttribute{
																Description:         "A map of additional key value pairs matching key value pairs in the beegfs-client.conf file. Seebeegfs-client.conf for more details. Values MUST be specified as strings, even if they appear to be integers orbooleans (e.g. '8000', not 8000 and 'true', not true).",
																MarkdownDescription: "A map of additional key value pairs matching key value pairs in the beegfs-client.conf file. Seebeegfs-client.conf for more details. Values MUST be specified as strings, even if they appear to be integers orbooleans (e.g. '8000', not 8000 and 'true', not true).",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"conn_interfaces": schema.ListAttribute{
																Description:         "A list of interfaces the BeeGFS client service can communicate over (e.g. 'ib0' or 'eth0'). Often not required.See beegfs-client.conf for more details.",
																MarkdownDescription: "A list of interfaces the BeeGFS client service can communicate over (e.g. 'ib0' or 'eth0'). Often not required.See beegfs-client.conf for more details.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"conn_net_filter": schema.ListAttribute{
																Description:         "A list of subnets the BeeGFS client service can use for outgoing communication (e.g. '10.10.10.10/24'). Oftennot required. See beegfs-client.conf for more details.",
																MarkdownDescription: "A list of subnets the BeeGFS client service can use for outgoing communication (e.g. '10.10.10.10/24'). Oftennot required. See beegfs-client.conf for more details.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"conn_rdma_interfaces": schema.ListAttribute{
																Description:         "A list of interfaces the BeeGFS client will use for outbound RDMA connections. This is used in supportof the BeeGFS multi-rail feature. This feature does not depend on or use the connInterfaces parameter.This feature requires the BeeGFS client version 7.3.0 or later.",
																MarkdownDescription: "A list of interfaces the BeeGFS client will use for outbound RDMA connections. This is used in supportof the BeeGFS multi-rail feature. This feature does not depend on or use the connInterfaces parameter.This feature requires the BeeGFS client version 7.3.0 or later.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"conn_tcp_only_filter": schema.ListAttribute{
																Description:         "A list of subnets in which RDMA communication cannot or should not be established (e.g. '10.10.10.11/24').Often not required. See beegfs-client.conf for more details.",
																MarkdownDescription: "A list of subnets in which RDMA communication cannot or should not be established (e.g. '10.10.10.11/24').Often not required. See beegfs-client.conf for more details.",
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

													"sys_mgmtd_host": schema.StringAttribute{
														Description:         "The sysMgmtdHost used by the BeeGFS client service to make initial contact with the BeeGFS mgmtd service.",
														MarkdownDescription: "The sysMgmtdHost used by the BeeGFS client service to make initial contact with the BeeGFS mgmtd service.",
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

										"node_list": schema.ListAttribute{
											Description:         "The list of nodes this configuration should be applied on. Each entry is the hostname of the node or the nameassigned to the node by the container orchestrator (e.g. 'node1' or 'cluster05-node03').",
											MarkdownDescription: "The list of nodes this configuration should be applied on. Each entry is the hostname of the node or the nameassigned to the node by the container orchestrator (e.g. 'node1' or 'cluster05-node03').",
											ElementType:         types.StringType,
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *BeegfsCsiNetappComBeegfsDriverV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_beegfs_csi_netapp_com_beegfs_driver_v1_manifest")

	var model BeegfsCsiNetappComBeegfsDriverV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("beegfs.csi.netapp.com/v1")
	model.Kind = pointer.String("BeegfsDriver")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
