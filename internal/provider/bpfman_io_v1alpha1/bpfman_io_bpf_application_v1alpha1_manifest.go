/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package bpfman_io_v1alpha1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/boolvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
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
	_ datasource.DataSource = &BpfmanIoBpfApplicationV1Alpha1Manifest{}
)

func NewBpfmanIoBpfApplicationV1Alpha1Manifest() datasource.DataSource {
	return &BpfmanIoBpfApplicationV1Alpha1Manifest{}
}

type BpfmanIoBpfApplicationV1Alpha1Manifest struct{}

type BpfmanIoBpfApplicationV1Alpha1ManifestData struct {
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
		ByteCode *struct {
			Image *struct {
				ImagePullPolicy *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
				ImagePullSecret *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"image_pull_secret" json:"imagePullSecret,omitempty"`
				Url *string `tfsdk:"url" json:"url,omitempty"`
			} `tfsdk:"image" json:"image,omitempty"`
			Path *string `tfsdk:"path" json:"path,omitempty"`
		} `tfsdk:"byte_code" json:"byteCode,omitempty"`
		GlobalData       *map[string]string `tfsdk:"global_data" json:"globalData,omitempty"`
		MapOwnerSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"map_owner_selector" json:"mapOwnerSelector,omitempty"`
		NodeSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
		Programs *[]struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
			Tc   *struct {
				Links *[]struct {
					Direction         *string `tfsdk:"direction" json:"direction,omitempty"`
					InterfaceSelector *struct {
						Interfaces                *[]string `tfsdk:"interfaces" json:"interfaces,omitempty"`
						InterfacesDiscoveryConfig *struct {
							AllowedInterfaces      *[]string `tfsdk:"allowed_interfaces" json:"allowedInterfaces,omitempty"`
							ExcludeInterfaces      *[]string `tfsdk:"exclude_interfaces" json:"excludeInterfaces,omitempty"`
							InterfaceAutoDiscovery *bool     `tfsdk:"interface_auto_discovery" json:"interfaceAutoDiscovery,omitempty"`
						} `tfsdk:"interfaces_discovery_config" json:"interfacesDiscoveryConfig,omitempty"`
						PrimaryNodeInterface *bool `tfsdk:"primary_node_interface" json:"primaryNodeInterface,omitempty"`
					} `tfsdk:"interface_selector" json:"interfaceSelector,omitempty"`
					NetworkNamespaces *struct {
						Pods *struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
						} `tfsdk:"pods" json:"pods,omitempty"`
					} `tfsdk:"network_namespaces" json:"networkNamespaces,omitempty"`
					Priority  *int64    `tfsdk:"priority" json:"priority,omitempty"`
					ProceedOn *[]string `tfsdk:"proceed_on" json:"proceedOn,omitempty"`
				} `tfsdk:"links" json:"links,omitempty"`
			} `tfsdk:"tc" json:"tc,omitempty"`
			Tcx *struct {
				Links *[]struct {
					Direction         *string `tfsdk:"direction" json:"direction,omitempty"`
					InterfaceSelector *struct {
						Interfaces                *[]string `tfsdk:"interfaces" json:"interfaces,omitempty"`
						InterfacesDiscoveryConfig *struct {
							AllowedInterfaces      *[]string `tfsdk:"allowed_interfaces" json:"allowedInterfaces,omitempty"`
							ExcludeInterfaces      *[]string `tfsdk:"exclude_interfaces" json:"excludeInterfaces,omitempty"`
							InterfaceAutoDiscovery *bool     `tfsdk:"interface_auto_discovery" json:"interfaceAutoDiscovery,omitempty"`
						} `tfsdk:"interfaces_discovery_config" json:"interfacesDiscoveryConfig,omitempty"`
						PrimaryNodeInterface *bool `tfsdk:"primary_node_interface" json:"primaryNodeInterface,omitempty"`
					} `tfsdk:"interface_selector" json:"interfaceSelector,omitempty"`
					NetworkNamespaces *struct {
						Pods *struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
						} `tfsdk:"pods" json:"pods,omitempty"`
					} `tfsdk:"network_namespaces" json:"networkNamespaces,omitempty"`
					Priority *int64 `tfsdk:"priority" json:"priority,omitempty"`
				} `tfsdk:"links" json:"links,omitempty"`
			} `tfsdk:"tcx" json:"tcx,omitempty"`
			Type   *string `tfsdk:"type" json:"type,omitempty"`
			Uprobe *struct {
				Links *[]struct {
					Containers *struct {
						ContainerNames *[]string `tfsdk:"container_names" json:"containerNames,omitempty"`
						Pods           *struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
						} `tfsdk:"pods" json:"pods,omitempty"`
					} `tfsdk:"containers" json:"containers,omitempty"`
					Function *string `tfsdk:"function" json:"function,omitempty"`
					Offset   *int64  `tfsdk:"offset" json:"offset,omitempty"`
					Pid      *int64  `tfsdk:"pid" json:"pid,omitempty"`
					Target   *string `tfsdk:"target" json:"target,omitempty"`
				} `tfsdk:"links" json:"links,omitempty"`
			} `tfsdk:"uprobe" json:"uprobe,omitempty"`
			Uretprobe *struct {
				Links *[]struct {
					Containers *struct {
						ContainerNames *[]string `tfsdk:"container_names" json:"containerNames,omitempty"`
						Pods           *struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
						} `tfsdk:"pods" json:"pods,omitempty"`
					} `tfsdk:"containers" json:"containers,omitempty"`
					Function *string `tfsdk:"function" json:"function,omitempty"`
					Offset   *int64  `tfsdk:"offset" json:"offset,omitempty"`
					Pid      *int64  `tfsdk:"pid" json:"pid,omitempty"`
					Target   *string `tfsdk:"target" json:"target,omitempty"`
				} `tfsdk:"links" json:"links,omitempty"`
			} `tfsdk:"uretprobe" json:"uretprobe,omitempty"`
			Xdp *struct {
				Links *[]struct {
					InterfaceSelector *struct {
						Interfaces                *[]string `tfsdk:"interfaces" json:"interfaces,omitempty"`
						InterfacesDiscoveryConfig *struct {
							AllowedInterfaces      *[]string `tfsdk:"allowed_interfaces" json:"allowedInterfaces,omitempty"`
							ExcludeInterfaces      *[]string `tfsdk:"exclude_interfaces" json:"excludeInterfaces,omitempty"`
							InterfaceAutoDiscovery *bool     `tfsdk:"interface_auto_discovery" json:"interfaceAutoDiscovery,omitempty"`
						} `tfsdk:"interfaces_discovery_config" json:"interfacesDiscoveryConfig,omitempty"`
						PrimaryNodeInterface *bool `tfsdk:"primary_node_interface" json:"primaryNodeInterface,omitempty"`
					} `tfsdk:"interface_selector" json:"interfaceSelector,omitempty"`
					NetworkNamespaces *struct {
						Pods *struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
						} `tfsdk:"pods" json:"pods,omitempty"`
					} `tfsdk:"network_namespaces" json:"networkNamespaces,omitempty"`
					Priority  *int64    `tfsdk:"priority" json:"priority,omitempty"`
					ProceedOn *[]string `tfsdk:"proceed_on" json:"proceedOn,omitempty"`
				} `tfsdk:"links" json:"links,omitempty"`
			} `tfsdk:"xdp" json:"xdp,omitempty"`
		} `tfsdk:"programs" json:"programs,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *BpfmanIoBpfApplicationV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_bpfman_io_bpf_application_v1alpha1_manifest"
}

func (r *BpfmanIoBpfApplicationV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "BpfApplication is the schema for the namespace scoped BPF Applications API. This API allows applications to use bpfman to load and attach one or more eBPF programs on a Kubernetes cluster. The bpfApplication.status field reports the overall status of the BpfApplication CRD. A given BpfApplication CRD can result in loading and attaching multiple eBPF programs on multiple nodes, so this status is just a summary. More granular per-node status details can be found in the corresponding BpfApplicationState CRD that bpfman creates for each node.",
		MarkdownDescription: "BpfApplication is the schema for the namespace scoped BPF Applications API. This API allows applications to use bpfman to load and attach one or more eBPF programs on a Kubernetes cluster. The bpfApplication.status field reports the overall status of the BpfApplication CRD. A given BpfApplication CRD can result in loading and attaching multiple eBPF programs on multiple nodes, so this status is just a summary. More granular per-node status details can be found in the corresponding BpfApplicationState CRD that bpfman creates for each node.",
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
				Description:         "spec defines the desired state of the BpfApplication. The BpfApplication describes the set of one or more namespace scoped eBPF programs that should be loaded for a given application and attributes for how they should be loaded. eBPF programs that are grouped together under the same BpfApplication instance can share maps and global data between the eBPF programs loaded on the same Kubernetes Node.",
				MarkdownDescription: "spec defines the desired state of the BpfApplication. The BpfApplication describes the set of one or more namespace scoped eBPF programs that should be loaded for a given application and attributes for how they should be loaded. eBPF programs that are grouped together under the same BpfApplication instance can share maps and global data between the eBPF programs loaded on the same Kubernetes Node.",
				Attributes: map[string]schema.Attribute{
					"byte_code": schema.SingleNestedAttribute{
						Description:         "bytecode is a required field and configures where the eBPF program's bytecode should be loaded from. The image must contain one or more eBPF programs.",
						MarkdownDescription: "bytecode is a required field and configures where the eBPF program's bytecode should be loaded from. The image must contain one or more eBPF programs.",
						Attributes: map[string]schema.Attribute{
							"image": schema.SingleNestedAttribute{
								Description:         "image is an optional field and used to specify details on how to retrieve an eBPF program packaged in a OCI container image from a given registry.",
								MarkdownDescription: "image is an optional field and used to specify details on how to retrieve an eBPF program packaged in a OCI container image from a given registry.",
								Attributes: map[string]schema.Attribute{
									"image_pull_policy": schema.StringAttribute{
										Description:         "pullPolicy is an optional field that describes a policy for if/when to pull a bytecode image. Defaults to IfNotPresent. Allowed values are: Always, IfNotPresent and Never When set to Always, the given image will be pulled even if the image is already present on the node. When set to IfNotPresent, the given image will only be pulled if it is not present on the node. When set to Never, the given image will never be pulled and must be loaded on the node by some other means.",
										MarkdownDescription: "pullPolicy is an optional field that describes a policy for if/when to pull a bytecode image. Defaults to IfNotPresent. Allowed values are: Always, IfNotPresent and Never When set to Always, the given image will be pulled even if the image is already present on the node. When set to IfNotPresent, the given image will only be pulled if it is not present on the node. When set to Never, the given image will never be pulled and must be loaded on the node by some other means.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("Always", "Never", "IfNotPresent"),
										},
									},

									"image_pull_secret": schema.SingleNestedAttribute{
										Description:         "imagePullSecret is an optional field and indicates the secret which contains the credentials to access the image repository.",
										MarkdownDescription: "imagePullSecret is an optional field and indicates the secret which contains the credentials to access the image repository.",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "name is a required field and is the name of the secret which contains the credentials to access the image repository.",
												MarkdownDescription: "name is a required field and is the name of the secret which contains the credentials to access the image repository.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "namespace is a required field and is the namespace of the secret which contains the credentials to access the image repository.",
												MarkdownDescription: "namespace is a required field and is the namespace of the secret which contains the credentials to access the image repository.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"url": schema.StringAttribute{
										Description:         "url is a required field and is a valid container image URL used to reference a remote bytecode image. url must not be an empty string, must not exceed 525 characters in length and must be a valid URL.",
										MarkdownDescription: "url is a required field and is a valid container image URL used to reference a remote bytecode image. url must not be an empty string, must not exceed 525 characters in length and must be a valid URL.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtMost(525),
											stringvalidator.RegexMatches(regexp.MustCompile(`[a-zA-Z0-9_][a-zA-Z0-9._-]{0,127}`), ""),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
								Validators: []validator.Object{
									objectvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("path")),
								},
							},

							"path": schema.StringAttribute{
								Description:         "path is an optional field and used to specify a bytecode object file via filepath on a Kubernetes node.",
								MarkdownDescription: "path is an optional field and used to specify a bytecode object file via filepath on a Kubernetes node.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(/[^/\0]+)+/?$`), ""),
									stringvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("image")),
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"global_data": schema.MapAttribute{
						Description:         "globalData is an optional field that allows the user to set global variables when the program is loaded. This allows the same compiled bytecode to be deployed by different BPF Applications to behave differently based on globalData configuration values. It uses an array of raw bytes. This is a very low level primitive. The caller is responsible for formatting the byte string appropriately considering such things as size, endianness, alignment and packing of data structures.",
						MarkdownDescription: "globalData is an optional field that allows the user to set global variables when the program is loaded. This allows the same compiled bytecode to be deployed by different BPF Applications to behave differently based on globalData configuration values. It uses an array of raw bytes. This is a very low level primitive. The caller is responsible for formatting the byte string appropriately considering such things as size, endianness, alignment and packing of data structures.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"map_owner_selector": schema.SingleNestedAttribute{
						Description:         "mapOwnerSelector is an optional field used to share maps across applications. eBPF programs loaded with the same ClusterBpfApplication or BpfApplication instance do not need to use this field. This label selector allows maps from a different ClusterBpfApplication or BpfApplication instance to be used by this instance. TODO: mapOwnerSelector is currently not supported due to recent code rework.",
						MarkdownDescription: "mapOwnerSelector is an optional field used to share maps across applications. eBPF programs loaded with the same ClusterBpfApplication or BpfApplication instance do not need to use this field. This label selector allows maps from a different ClusterBpfApplication or BpfApplication instance to be used by this instance. TODO: mapOwnerSelector is currently not supported due to recent code rework.",
						Attributes: map[string]schema.Attribute{
							"match_expressions": schema.ListNestedAttribute{
								Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "key is the label key that the selector applies to.",
											MarkdownDescription: "key is the label key that the selector applies to.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"operator": schema.StringAttribute{
											Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"values": schema.ListAttribute{
											Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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

							"match_labels": schema.MapAttribute{
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

					"node_selector": schema.SingleNestedAttribute{
						Description:         "nodeSelector is a required field and allows the user to specify which Kubernetes nodes to deploy the eBPF programs. To select all nodes use standard metav1.LabelSelector semantics and make it empty.",
						MarkdownDescription: "nodeSelector is a required field and allows the user to specify which Kubernetes nodes to deploy the eBPF programs. To select all nodes use standard metav1.LabelSelector semantics and make it empty.",
						Attributes: map[string]schema.Attribute{
							"match_expressions": schema.ListNestedAttribute{
								Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "key is the label key that the selector applies to.",
											MarkdownDescription: "key is the label key that the selector applies to.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"operator": schema.StringAttribute{
											Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"values": schema.ListAttribute{
											Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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

							"match_labels": schema.MapAttribute{
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

					"programs": schema.ListNestedAttribute{
						Description:         "programs is a required field and is the list of eBPF programs in a BPF Application CRD that should be loaded in kernel memory. At least one entry is required. eBPF programs in this list will be loaded on the system based the nodeSelector. Even if an eBPF program is loaded in kernel memory, it cannot be triggered until an attachment point is provided. The different program types have different ways of attaching. The attachment points can be added at creation time or modified (added or removed) at a later time to activate or deactivate the eBPF program as desired. CAUTION: When programs are added or removed from the list, that requires all programs in the list to be reloaded, which could be temporarily service effecting. For this reason, modifying the list is currently not allowed.",
						MarkdownDescription: "programs is a required field and is the list of eBPF programs in a BPF Application CRD that should be loaded in kernel memory. At least one entry is required. eBPF programs in this list will be loaded on the system based the nodeSelector. Even if an eBPF program is loaded in kernel memory, it cannot be triggered until an attachment point is provided. The different program types have different ways of attaching. The attachment points can be added at creation time or modified (added or removed) at a later time to activate or deactivate the eBPF program as desired. CAUTION: When programs are added or removed from the list, that requires all programs in the list to be reloaded, which could be temporarily service effecting. For this reason, modifying the list is currently not allowed.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "name is a required field and is the name of the function that is the entry point for the eBPF program. name must not be an empty string, must not exceed 64 characters in length, must start with alpha characters and must only contain alphanumeric characters.",
									MarkdownDescription: "name is a required field and is the name of the function that is the entry point for the eBPF program. name must not be an empty string, must not exceed 64 characters in length, must start with alpha characters and must only contain alphanumeric characters.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(64),
										stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]+.`), ""),
									},
								},

								"tc": schema.SingleNestedAttribute{
									Description:         "tc is an optional field, but required when the type field is set to TC. tc defines the desired state of the application's TC programs. TC programs are attached to network devices (interfaces). The program can be attached on either packet ingress or egress, so the program will be called on every incoming or outgoing packet seen by the network device. The TC attachment point is in Linux's Traffic Control (tc) subsystem, which is after the Linux kernel has allocated an sk_buff. TCX is newer implementation of TC with enhanced performance and better support for running multiple programs on a given network device. This makes TC useful for packet classification actions.",
									MarkdownDescription: "tc is an optional field, but required when the type field is set to TC. tc defines the desired state of the application's TC programs. TC programs are attached to network devices (interfaces). The program can be attached on either packet ingress or egress, so the program will be called on every incoming or outgoing packet seen by the network device. The TC attachment point is in Linux's Traffic Control (tc) subsystem, which is after the Linux kernel has allocated an sk_buff. TCX is newer implementation of TC with enhanced performance and better support for running multiple programs on a given network device. This makes TC useful for packet classification actions.",
									Attributes: map[string]schema.Attribute{
										"links": schema.ListNestedAttribute{
											Description:         "links is an optional field and is the list of attachment points to which the TC program should be attached. The TC program is loaded in kernel memory when the BPF Application CRD is created and the selected Kubernetes nodes are active. The TC program will not be triggered until the program has also been attached to an attachment point described in this list. Items may be added or removed from the list at any point, causing the TC program to be attached or detached. The attachment point for a TC program is a network interface (or device). The interface can be specified by name, by allowing bpfman to discover each interface, or by setting the primaryNodeInterface flag, which instructs bpfman to use the primary interface of a Kubernetes node. Optionally, the TC program can also be installed into a set of network namespaces.",
											MarkdownDescription: "links is an optional field and is the list of attachment points to which the TC program should be attached. The TC program is loaded in kernel memory when the BPF Application CRD is created and the selected Kubernetes nodes are active. The TC program will not be triggered until the program has also been attached to an attachment point described in this list. Items may be added or removed from the list at any point, causing the TC program to be attached or detached. The attachment point for a TC program is a network interface (or device). The interface can be specified by name, by allowing bpfman to discover each interface, or by setting the primaryNodeInterface flag, which instructs bpfman to use the primary interface of a Kubernetes node. Optionally, the TC program can also be installed into a set of network namespaces.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"direction": schema.StringAttribute{
														Description:         "direction is a required field and specifies the direction of traffic. Allowed values are: Ingress, Egress When set to Ingress, the TC program is triggered when packets are received by the interface. When set to Egress, the TC program is triggered when packets are to be transmitted by the interface.",
														MarkdownDescription: "direction is a required field and specifies the direction of traffic. Allowed values are: Ingress, Egress When set to Ingress, the TC program is triggered when packets are received by the interface. When set to Egress, the TC program is triggered when packets are to be transmitted by the interface.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("Ingress", "Egress"),
														},
													},

													"interface_selector": schema.SingleNestedAttribute{
														Description:         "interfaceSelector is a required field and is used to determine the network interface (or interfaces) the TC program is attached. Interface list is set by providing a list of interface names, enabling auto discovery, or setting the primaryNodeInterface flag, but only one option is allowed.",
														MarkdownDescription: "interfaceSelector is a required field and is used to determine the network interface (or interfaces) the TC program is attached. Interface list is set by providing a list of interface names, enabling auto discovery, or setting the primaryNodeInterface flag, but only one option is allowed.",
														Attributes: map[string]schema.Attribute{
															"interfaces": schema.ListAttribute{
																Description:         "interfaces is an optional field and is a list of network interface names to attach the eBPF program. The interface names in the list are case-sensitive.",
																MarkdownDescription: "interfaces is an optional field and is a list of network interface names to attach the eBPF program. The interface names in the list are case-sensitive.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.List{
																	listvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("interfaces_discovery_config"), path.MatchRelative().AtParent().AtName("primary_node_interface")),
																},
															},

															"interfaces_discovery_config": schema.SingleNestedAttribute{
																Description:         "interfacesDiscoveryConfig is an optional field that is used to control if and how to automatically discover interfaces. If the agent should automatically discover and attach eBPF programs to interfaces, use the fields under interfacesDiscoveryConfig to control what is allow and excluded from discovery.",
																MarkdownDescription: "interfacesDiscoveryConfig is an optional field that is used to control if and how to automatically discover interfaces. If the agent should automatically discover and attach eBPF programs to interfaces, use the fields under interfacesDiscoveryConfig to control what is allow and excluded from discovery.",
																Attributes: map[string]schema.Attribute{
																	"allowed_interfaces": schema.ListAttribute{
																		Description:         "allowedInterfaces is an optional field that contains a list of interface names that are allowed to be discovered. If empty, the agent will fetch all the interfaces in the system, excepting the ones listed in excludeInterfaces. if non-empty, only entries in the list will be considered for discovery. If an entry enclosed by slashes, such as '/br-/' or '/veth*/', then the entry is considered as a regular expression for matching. Otherwise, the interface names in the list are case-sensitive. This field is only taken into consideration if interfaceAutoDiscovery is set to true.",
																		MarkdownDescription: "allowedInterfaces is an optional field that contains a list of interface names that are allowed to be discovered. If empty, the agent will fetch all the interfaces in the system, excepting the ones listed in excludeInterfaces. if non-empty, only entries in the list will be considered for discovery. If an entry enclosed by slashes, such as '/br-/' or '/veth*/', then the entry is considered as a regular expression for matching. Otherwise, the interface names in the list are case-sensitive. This field is only taken into consideration if interfaceAutoDiscovery is set to true.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"exclude_interfaces": schema.ListAttribute{
																		Description:         "excludeInterfaces is an optional field that contains a list of interface names that are excluded from interface discovery. The interface names in the list are case-sensitive. By default, the list contains the loopback interface, 'lo'. This field is only taken into consideration if interfaceAutoDiscovery is set to true.",
																		MarkdownDescription: "excludeInterfaces is an optional field that contains a list of interface names that are excluded from interface discovery. The interface names in the list are case-sensitive. By default, the list contains the loopback interface, 'lo'. This field is only taken into consideration if interfaceAutoDiscovery is set to true.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"interface_auto_discovery": schema.BoolAttribute{
																		Description:         "interfaceAutoDiscovery is an optional field. When enabled, the agent monitors the creation and deletion of interfaces and automatically attached eBPF programs to the newly discovered interfaces. CAUTION: This has the potential to attach a given eBPF program to a large number of interfaces. Use with caution.",
																		MarkdownDescription: "interfaceAutoDiscovery is an optional field. When enabled, the agent monitors the creation and deletion of interfaces and automatically attached eBPF programs to the newly discovered interfaces. CAUTION: This has the potential to attach a given eBPF program to a large number of interfaces. Use with caution.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
																Validators: []validator.Object{
																	objectvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("interfaces"), path.MatchRelative().AtParent().AtName("primary_node_interface")),
																},
															},

															"primary_node_interface": schema.BoolAttribute{
																Description:         "primaryNodeInterface is and optional field and indicates to attach the eBPF program to the primary interface on the Kubernetes node. Only 'true' is accepted.",
																MarkdownDescription: "primaryNodeInterface is and optional field and indicates to attach the eBPF program to the primary interface on the Kubernetes node. Only 'true' is accepted.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Bool{
																	boolvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("interfaces"), path.MatchRelative().AtParent().AtName("interfaces_discovery_config")),
																},
															},
														},
														Required: true,
														Optional: false,
														Computed: false,
													},

													"network_namespaces": schema.SingleNestedAttribute{
														Description:         "networkNamespaces is a required field that identifies the set of network namespaces in which to attach the eBPF program.",
														MarkdownDescription: "networkNamespaces is a required field that identifies the set of network namespaces in which to attach the eBPF program.",
														Attributes: map[string]schema.Attribute{
															"pods": schema.SingleNestedAttribute{
																Description:         "pods is a required field and indicates the target pods. To select all pods use the standard metav1.LabelSelector semantics and make it empty.",
																MarkdownDescription: "pods is a required field and indicates the target pods. To select all pods use the standard metav1.LabelSelector semantics and make it empty.",
																Attributes: map[string]schema.Attribute{
																	"match_expressions": schema.ListNestedAttribute{
																		Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																		MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"key": schema.StringAttribute{
																					Description:         "key is the label key that the selector applies to.",
																					MarkdownDescription: "key is the label key that the selector applies to.",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"operator": schema.StringAttribute{
																					Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																					MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"values": schema.ListAttribute{
																					Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																					MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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

																	"match_labels": schema.MapAttribute{
																		Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																		MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
														},
														Required: true,
														Optional: false,
														Computed: false,
													},

													"priority": schema.Int64Attribute{
														Description:         "priority is an optional field and determines the execution order of the TC program relative to other TC programs attached to the same attachment point. It must be a value between 0 and 1000, where lower values indicate higher precedence. For TC programs on the same attachment point with the same direction and priority, the most recently attached program has a lower precedence. If not provided, priority will default to 1000.",
														MarkdownDescription: "priority is an optional field and determines the execution order of the TC program relative to other TC programs attached to the same attachment point. It must be a value between 0 and 1000, where lower values indicate higher precedence. For TC programs on the same attachment point with the same direction and priority, the most recently attached program has a lower precedence. If not provided, priority will default to 1000.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(0),
															int64validator.AtMost(1000),
														},
													},

													"proceed_on": schema.ListAttribute{
														Description:         "proceedOn is an optional field and allows the user to call other TC programs in a chain, or not call the next program in a chain based on the exit code of a TC program. Allowed values, which are the possible exit codes from a TC eBPF program, are: UnSpec, OK, ReClassify, Shot, Pipe, Stolen, Queued, Repeat, ReDirect, Trap, DispatcherReturn Multiple values are supported. Default is OK, Pipe and DispatcherReturn. So using the default values, if a TC program returns Pipe, the next TC program in the chain will be called. If a TC program returns Stolen, the next TC program in the chain will NOT be called.",
														MarkdownDescription: "proceedOn is an optional field and allows the user to call other TC programs in a chain, or not call the next program in a chain based on the exit code of a TC program. Allowed values, which are the possible exit codes from a TC eBPF program, are: UnSpec, OK, ReClassify, Shot, Pipe, Stolen, Queued, Repeat, ReDirect, Trap, DispatcherReturn Multiple values are supported. Default is OK, Pipe and DispatcherReturn. So using the default values, if a TC program returns Pipe, the next TC program in the chain will be called. If a TC program returns Stolen, the next TC program in the chain will NOT be called.",
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
									Required: false,
									Optional: true,
									Computed: false,
								},

								"tcx": schema.SingleNestedAttribute{
									Description:         "tcx is an optional field, but required when the type field is set to TCX. tcx defines the desired state of the application's TCX programs. TCX programs are attached to network devices (interfaces). The program can be attached on either packet ingress or egress, so the program will be called on every incoming or outgoing packet seen by the network device. The TCX attachment point is in Linux's Traffic Control (tc) subsystem, which is after the Linux kernel has allocated an sk_buff. This makes TCX useful for packet classification actions. TCX is a newer implementation of TC with enhanced performance and better support for running multiple programs on a given network device.",
									MarkdownDescription: "tcx is an optional field, but required when the type field is set to TCX. tcx defines the desired state of the application's TCX programs. TCX programs are attached to network devices (interfaces). The program can be attached on either packet ingress or egress, so the program will be called on every incoming or outgoing packet seen by the network device. The TCX attachment point is in Linux's Traffic Control (tc) subsystem, which is after the Linux kernel has allocated an sk_buff. This makes TCX useful for packet classification actions. TCX is a newer implementation of TC with enhanced performance and better support for running multiple programs on a given network device.",
									Attributes: map[string]schema.Attribute{
										"links": schema.ListNestedAttribute{
											Description:         "links is an optional field and is the list of attachment points to which the TCX program should be attached. The TCX program is loaded in kernel memory when the BPF Application CRD is created and the selected Kubernetes nodes are active. The TCX program will not be triggered until the program has also been attached to an attachment point described in this list. Items may be added or removed from the list at any point, causing the TCX program to be attached or detached. The attachment point for a TCX program is a network interface (or device). The interface can be specified by name, by allowing bpfman to discover each interface, or by setting the primaryNodeInterface flag, which instructs bpfman to use the primary interface of a Kubernetes node. Optionally, the TCX program can also be installed into a set of network namespaces.",
											MarkdownDescription: "links is an optional field and is the list of attachment points to which the TCX program should be attached. The TCX program is loaded in kernel memory when the BPF Application CRD is created and the selected Kubernetes nodes are active. The TCX program will not be triggered until the program has also been attached to an attachment point described in this list. Items may be added or removed from the list at any point, causing the TCX program to be attached or detached. The attachment point for a TCX program is a network interface (or device). The interface can be specified by name, by allowing bpfman to discover each interface, or by setting the primaryNodeInterface flag, which instructs bpfman to use the primary interface of a Kubernetes node. Optionally, the TCX program can also be installed into a set of network namespaces.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"direction": schema.StringAttribute{
														Description:         "direction is a required field and specifies the direction of traffic. Allowed values are: Ingress, Egress When set to Ingress, the TC program is triggered when packets are received by the interface. When set to Egress, the TC program is triggered when packets are to be transmitted by the interface.",
														MarkdownDescription: "direction is a required field and specifies the direction of traffic. Allowed values are: Ingress, Egress When set to Ingress, the TC program is triggered when packets are received by the interface. When set to Egress, the TC program is triggered when packets are to be transmitted by the interface.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("Ingress", "Egress"),
														},
													},

													"interface_selector": schema.SingleNestedAttribute{
														Description:         "interfaceSelector is a required field and is used to determine the network interface (or interfaces) the TCX program is attached. Interface list is set by providing a list of interface names, enabling auto discovery, or setting the primaryNodeInterface flag, but only one option is allowed.",
														MarkdownDescription: "interfaceSelector is a required field and is used to determine the network interface (or interfaces) the TCX program is attached. Interface list is set by providing a list of interface names, enabling auto discovery, or setting the primaryNodeInterface flag, but only one option is allowed.",
														Attributes: map[string]schema.Attribute{
															"interfaces": schema.ListAttribute{
																Description:         "interfaces is an optional field and is a list of network interface names to attach the eBPF program. The interface names in the list are case-sensitive.",
																MarkdownDescription: "interfaces is an optional field and is a list of network interface names to attach the eBPF program. The interface names in the list are case-sensitive.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.List{
																	listvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("interfaces_discovery_config"), path.MatchRelative().AtParent().AtName("primary_node_interface")),
																},
															},

															"interfaces_discovery_config": schema.SingleNestedAttribute{
																Description:         "interfacesDiscoveryConfig is an optional field that is used to control if and how to automatically discover interfaces. If the agent should automatically discover and attach eBPF programs to interfaces, use the fields under interfacesDiscoveryConfig to control what is allow and excluded from discovery.",
																MarkdownDescription: "interfacesDiscoveryConfig is an optional field that is used to control if and how to automatically discover interfaces. If the agent should automatically discover and attach eBPF programs to interfaces, use the fields under interfacesDiscoveryConfig to control what is allow and excluded from discovery.",
																Attributes: map[string]schema.Attribute{
																	"allowed_interfaces": schema.ListAttribute{
																		Description:         "allowedInterfaces is an optional field that contains a list of interface names that are allowed to be discovered. If empty, the agent will fetch all the interfaces in the system, excepting the ones listed in excludeInterfaces. if non-empty, only entries in the list will be considered for discovery. If an entry enclosed by slashes, such as '/br-/' or '/veth*/', then the entry is considered as a regular expression for matching. Otherwise, the interface names in the list are case-sensitive. This field is only taken into consideration if interfaceAutoDiscovery is set to true.",
																		MarkdownDescription: "allowedInterfaces is an optional field that contains a list of interface names that are allowed to be discovered. If empty, the agent will fetch all the interfaces in the system, excepting the ones listed in excludeInterfaces. if non-empty, only entries in the list will be considered for discovery. If an entry enclosed by slashes, such as '/br-/' or '/veth*/', then the entry is considered as a regular expression for matching. Otherwise, the interface names in the list are case-sensitive. This field is only taken into consideration if interfaceAutoDiscovery is set to true.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"exclude_interfaces": schema.ListAttribute{
																		Description:         "excludeInterfaces is an optional field that contains a list of interface names that are excluded from interface discovery. The interface names in the list are case-sensitive. By default, the list contains the loopback interface, 'lo'. This field is only taken into consideration if interfaceAutoDiscovery is set to true.",
																		MarkdownDescription: "excludeInterfaces is an optional field that contains a list of interface names that are excluded from interface discovery. The interface names in the list are case-sensitive. By default, the list contains the loopback interface, 'lo'. This field is only taken into consideration if interfaceAutoDiscovery is set to true.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"interface_auto_discovery": schema.BoolAttribute{
																		Description:         "interfaceAutoDiscovery is an optional field. When enabled, the agent monitors the creation and deletion of interfaces and automatically attached eBPF programs to the newly discovered interfaces. CAUTION: This has the potential to attach a given eBPF program to a large number of interfaces. Use with caution.",
																		MarkdownDescription: "interfaceAutoDiscovery is an optional field. When enabled, the agent monitors the creation and deletion of interfaces and automatically attached eBPF programs to the newly discovered interfaces. CAUTION: This has the potential to attach a given eBPF program to a large number of interfaces. Use with caution.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
																Validators: []validator.Object{
																	objectvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("interfaces"), path.MatchRelative().AtParent().AtName("primary_node_interface")),
																},
															},

															"primary_node_interface": schema.BoolAttribute{
																Description:         "primaryNodeInterface is and optional field and indicates to attach the eBPF program to the primary interface on the Kubernetes node. Only 'true' is accepted.",
																MarkdownDescription: "primaryNodeInterface is and optional field and indicates to attach the eBPF program to the primary interface on the Kubernetes node. Only 'true' is accepted.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Bool{
																	boolvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("interfaces"), path.MatchRelative().AtParent().AtName("interfaces_discovery_config")),
																},
															},
														},
														Required: true,
														Optional: false,
														Computed: false,
													},

													"network_namespaces": schema.SingleNestedAttribute{
														Description:         "networkNamespaces is a required field that identifies the set of network namespaces in which to attach the eBPF program.",
														MarkdownDescription: "networkNamespaces is a required field that identifies the set of network namespaces in which to attach the eBPF program.",
														Attributes: map[string]schema.Attribute{
															"pods": schema.SingleNestedAttribute{
																Description:         "pods is a required field and indicates the target pods. To select all pods use the standard metav1.LabelSelector semantics and make it empty.",
																MarkdownDescription: "pods is a required field and indicates the target pods. To select all pods use the standard metav1.LabelSelector semantics and make it empty.",
																Attributes: map[string]schema.Attribute{
																	"match_expressions": schema.ListNestedAttribute{
																		Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																		MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"key": schema.StringAttribute{
																					Description:         "key is the label key that the selector applies to.",
																					MarkdownDescription: "key is the label key that the selector applies to.",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"operator": schema.StringAttribute{
																					Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																					MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"values": schema.ListAttribute{
																					Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																					MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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

																	"match_labels": schema.MapAttribute{
																		Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																		MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
														},
														Required: true,
														Optional: false,
														Computed: false,
													},

													"priority": schema.Int64Attribute{
														Description:         "priority is an optional field and determines the execution order of the TCX program relative to other TCX programs attached to the same attachment point. It must be a value between 0 and 1000, where lower values indicate higher precedence. For TCX programs on the same attachment point with the same direction and priority, the most recently attached program has a lower precedence. If not provided, priority will default to 1000.",
														MarkdownDescription: "priority is an optional field and determines the execution order of the TCX program relative to other TCX programs attached to the same attachment point. It must be a value between 0 and 1000, where lower values indicate higher precedence. For TCX programs on the same attachment point with the same direction and priority, the most recently attached program has a lower precedence. If not provided, priority will default to 1000.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(0),
															int64validator.AtMost(1000),
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
									Description:         "type is a required field used to specify the type of the eBPF program. Allowed values are: TC, TCX, UProbe, URetProbe, XDP When set to TC, the eBPF program can attach to network devices (interfaces). The program can be attached on either packet ingress or egress, so the program will be called on every incoming or outgoing packet seen by the network device. When using the TC program type, the tc field is required. See tc for more details on TC programs. When set to TCX, the eBPF program can attach to network devices (interfaces). The program can be attached on either packet ingress or egress, so the program will be called on every incoming or outgoing packet seen by the network device. When using the TCX program type, the tcx field is required. See tcx for more details on TCX programs. When set to UProbe, the program can attach in user-space. The UProbe is attached to a binary, library or function name, and optionally an offset in the code. When using the UProbe program type, the uprobe field is required. See uprobe for more details on UProbe programs. When set to URetProbe, the program can attach in user-space. The URetProbe is attached to the return of a binary, library or function name, and optionally an offset in the code. When using the URetProbe program type, the uretprobe field is required. See uretprobe for more details on URetProbe programs. When set to XDP, the eBPF program can attach to network devices (interfaces) and will be called on every incoming packet received by the network device. When using the XDP program type, the xdp field is required. See xdp for more details on XDP programs.",
									MarkdownDescription: "type is a required field used to specify the type of the eBPF program. Allowed values are: TC, TCX, UProbe, URetProbe, XDP When set to TC, the eBPF program can attach to network devices (interfaces). The program can be attached on either packet ingress or egress, so the program will be called on every incoming or outgoing packet seen by the network device. When using the TC program type, the tc field is required. See tc for more details on TC programs. When set to TCX, the eBPF program can attach to network devices (interfaces). The program can be attached on either packet ingress or egress, so the program will be called on every incoming or outgoing packet seen by the network device. When using the TCX program type, the tcx field is required. See tcx for more details on TCX programs. When set to UProbe, the program can attach in user-space. The UProbe is attached to a binary, library or function name, and optionally an offset in the code. When using the UProbe program type, the uprobe field is required. See uprobe for more details on UProbe programs. When set to URetProbe, the program can attach in user-space. The URetProbe is attached to the return of a binary, library or function name, and optionally an offset in the code. When using the URetProbe program type, the uretprobe field is required. See uretprobe for more details on URetProbe programs. When set to XDP, the eBPF program can attach to network devices (interfaces) and will be called on every incoming packet received by the network device. When using the XDP program type, the xdp field is required. See xdp for more details on XDP programs.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("XDP", "TC", "TCX", "UProbe", "URetProbe"),
									},
								},

								"uprobe": schema.SingleNestedAttribute{
									Description:         "uprobe is an optional field, but required when the type field is set to UProbe. uprobe defines the desired state of the application's UProbe programs. UProbe programs are user-space probes. A target must be provided, which is the library name or absolute path to a binary or library where the probe is attached. Optionally, a function name can also be provided to provide finer granularity on where the probe is attached. They can be attached at any point in the binary, library or function using the optional offset field. However, caution must be taken when using the offset, ensuring the offset is still in the desired bytecode.",
									MarkdownDescription: "uprobe is an optional field, but required when the type field is set to UProbe. uprobe defines the desired state of the application's UProbe programs. UProbe programs are user-space probes. A target must be provided, which is the library name or absolute path to a binary or library where the probe is attached. Optionally, a function name can also be provided to provide finer granularity on where the probe is attached. They can be attached at any point in the binary, library or function using the optional offset field. However, caution must be taken when using the offset, ensuring the offset is still in the desired bytecode.",
									Attributes: map[string]schema.Attribute{
										"links": schema.ListNestedAttribute{
											Description:         "links is an optional field and is the list of attachment points to which the UProbe or URetProbe program should be attached. The eBPF program is loaded in kernel memory when the BPF Application CRD is created and the selected Kubernetes nodes are active. The eBPF program will not be triggered until the program has also been attached to an attachment point described in this list. Items may be added or removed from the list at any point, causing the eBPF program to be attached or detached. The attachment point for a UProbe and URetProbe program is a user-space binary or function. By default, the eBPF program is triggered at the entry of the attachment point, but the attachment point can be adjusted using an optional function name and/or offset. Optionally, the eBPF program can be installed in a set of containers or limited to a specified PID.",
											MarkdownDescription: "links is an optional field and is the list of attachment points to which the UProbe or URetProbe program should be attached. The eBPF program is loaded in kernel memory when the BPF Application CRD is created and the selected Kubernetes nodes are active. The eBPF program will not be triggered until the program has also been attached to an attachment point described in this list. Items may be added or removed from the list at any point, causing the eBPF program to be attached or detached. The attachment point for a UProbe and URetProbe program is a user-space binary or function. By default, the eBPF program is triggered at the entry of the attachment point, but the attachment point can be adjusted using an optional function name and/or offset. Optionally, the eBPF program can be installed in a set of containers or limited to a specified PID.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"containers": schema.SingleNestedAttribute{
														Description:         "containers is an optional field that identifies the set of containers in which to attach the UProbe or URetProbe program. If containers is not specified, the eBPF program will be attached in the bpfman container. uprobe.",
														MarkdownDescription: "containers is an optional field that identifies the set of containers in which to attach the UProbe or URetProbe program. If containers is not specified, the eBPF program will be attached in the bpfman container. uprobe.",
														Attributes: map[string]schema.Attribute{
															"container_names": schema.ListAttribute{
																Description:         "containerNames is an optional field and is a list of container names in a pod to attach the eBPF program. If no names are specified, all containers in the pod are selected.",
																MarkdownDescription: "containerNames is an optional field and is a list of container names in a pod to attach the eBPF program. If no names are specified, all containers in the pod are selected.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"pods": schema.SingleNestedAttribute{
																Description:         "pods is a required field and indicates the target pods. To select all pods use the standard metav1.LabelSelector semantics and make it empty.",
																MarkdownDescription: "pods is a required field and indicates the target pods. To select all pods use the standard metav1.LabelSelector semantics and make it empty.",
																Attributes: map[string]schema.Attribute{
																	"match_expressions": schema.ListNestedAttribute{
																		Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																		MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"key": schema.StringAttribute{
																					Description:         "key is the label key that the selector applies to.",
																					MarkdownDescription: "key is the label key that the selector applies to.",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"operator": schema.StringAttribute{
																					Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																					MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"values": schema.ListAttribute{
																					Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																					MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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

																	"match_labels": schema.MapAttribute{
																		Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																		MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
														},
														Required: true,
														Optional: false,
														Computed: false,
													},

													"function": schema.StringAttribute{
														Description:         "function is an optional field and specifies the name of a user-space function to attach the UProbe or URetProbe program. If not provided, the eBPF program will be triggered on the entry of the target. function must not be an empty string, must not exceed 64 characters in length, must start with alpha characters and must only contain alphanumeric characters.",
														MarkdownDescription: "function is an optional field and specifies the name of a user-space function to attach the UProbe or URetProbe program. If not provided, the eBPF program will be triggered on the entry of the target. function must not be an empty string, must not exceed 64 characters in length, must start with alpha characters and must only contain alphanumeric characters.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.LengthAtMost(64),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]+.`), ""),
														},
													},

													"offset": schema.Int64Attribute{
														Description:         "offset is an optional field and the value is added to the address of the attachment point function. If not provided, offset defaults to 0.",
														MarkdownDescription: "offset is an optional field and the value is added to the address of the attachment point function. If not provided, offset defaults to 0.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"pid": schema.Int64Attribute{
														Description:         "pid is an optional field and if provided, limits the execution of the UProbe or URetProbe to the provided process identification number (PID). If pid is not provided, the UProbe or URetProbe executes for all PIDs.",
														MarkdownDescription: "pid is an optional field and if provided, limits the execution of the UProbe or URetProbe to the provided process identification number (PID). If pid is not provided, the UProbe or URetProbe executes for all PIDs.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"target": schema.StringAttribute{
														Description:         "target is a required field and is the user-space library name or the absolute path to a binary or library.",
														MarkdownDescription: "target is a required field and is the user-space library name or the absolute path to a binary or library.",
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

								"uretprobe": schema.SingleNestedAttribute{
									Description:         "uretprobe is an optional field, but required when the type field is set to URetProbe. uretprobe defines the desired state of the application's URetProbe programs. URetProbe programs are user-space probes. A target must be provided, which is the library name or absolute path to a binary or library where the probe is attached. Optionally, a function name can also be provided to provide finer granularity on where the probe is attached. They are attached to the return point of the binary, library or function, but can be set anywhere using the optional offset field. However, caution must be taken when using the offset, ensuring the offset is still in the desired bytecode.",
									MarkdownDescription: "uretprobe is an optional field, but required when the type field is set to URetProbe. uretprobe defines the desired state of the application's URetProbe programs. URetProbe programs are user-space probes. A target must be provided, which is the library name or absolute path to a binary or library where the probe is attached. Optionally, a function name can also be provided to provide finer granularity on where the probe is attached. They are attached to the return point of the binary, library or function, but can be set anywhere using the optional offset field. However, caution must be taken when using the offset, ensuring the offset is still in the desired bytecode.",
									Attributes: map[string]schema.Attribute{
										"links": schema.ListNestedAttribute{
											Description:         "links is an optional field and is the list of attachment points to which the UProbe or URetProbe program should be attached. The eBPF program is loaded in kernel memory when the BPF Application CRD is created and the selected Kubernetes nodes are active. The eBPF program will not be triggered until the program has also been attached to an attachment point described in this list. Items may be added or removed from the list at any point, causing the eBPF program to be attached or detached. The attachment point for a UProbe and URetProbe program is a user-space binary or function. By default, the eBPF program is triggered at the entry of the attachment point, but the attachment point can be adjusted using an optional function name and/or offset. Optionally, the eBPF program can be installed in a set of containers or limited to a specified PID.",
											MarkdownDescription: "links is an optional field and is the list of attachment points to which the UProbe or URetProbe program should be attached. The eBPF program is loaded in kernel memory when the BPF Application CRD is created and the selected Kubernetes nodes are active. The eBPF program will not be triggered until the program has also been attached to an attachment point described in this list. Items may be added or removed from the list at any point, causing the eBPF program to be attached or detached. The attachment point for a UProbe and URetProbe program is a user-space binary or function. By default, the eBPF program is triggered at the entry of the attachment point, but the attachment point can be adjusted using an optional function name and/or offset. Optionally, the eBPF program can be installed in a set of containers or limited to a specified PID.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"containers": schema.SingleNestedAttribute{
														Description:         "containers is an optional field that identifies the set of containers in which to attach the UProbe or URetProbe program. If containers is not specified, the eBPF program will be attached in the bpfman container. uprobe.",
														MarkdownDescription: "containers is an optional field that identifies the set of containers in which to attach the UProbe or URetProbe program. If containers is not specified, the eBPF program will be attached in the bpfman container. uprobe.",
														Attributes: map[string]schema.Attribute{
															"container_names": schema.ListAttribute{
																Description:         "containerNames is an optional field and is a list of container names in a pod to attach the eBPF program. If no names are specified, all containers in the pod are selected.",
																MarkdownDescription: "containerNames is an optional field and is a list of container names in a pod to attach the eBPF program. If no names are specified, all containers in the pod are selected.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"pods": schema.SingleNestedAttribute{
																Description:         "pods is a required field and indicates the target pods. To select all pods use the standard metav1.LabelSelector semantics and make it empty.",
																MarkdownDescription: "pods is a required field and indicates the target pods. To select all pods use the standard metav1.LabelSelector semantics and make it empty.",
																Attributes: map[string]schema.Attribute{
																	"match_expressions": schema.ListNestedAttribute{
																		Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																		MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"key": schema.StringAttribute{
																					Description:         "key is the label key that the selector applies to.",
																					MarkdownDescription: "key is the label key that the selector applies to.",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"operator": schema.StringAttribute{
																					Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																					MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"values": schema.ListAttribute{
																					Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																					MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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

																	"match_labels": schema.MapAttribute{
																		Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																		MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
														},
														Required: true,
														Optional: false,
														Computed: false,
													},

													"function": schema.StringAttribute{
														Description:         "function is an optional field and specifies the name of a user-space function to attach the UProbe or URetProbe program. If not provided, the eBPF program will be triggered on the entry of the target. function must not be an empty string, must not exceed 64 characters in length, must start with alpha characters and must only contain alphanumeric characters.",
														MarkdownDescription: "function is an optional field and specifies the name of a user-space function to attach the UProbe or URetProbe program. If not provided, the eBPF program will be triggered on the entry of the target. function must not be an empty string, must not exceed 64 characters in length, must start with alpha characters and must only contain alphanumeric characters.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.LengthAtMost(64),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]+.`), ""),
														},
													},

													"offset": schema.Int64Attribute{
														Description:         "offset is an optional field and the value is added to the address of the attachment point function. If not provided, offset defaults to 0.",
														MarkdownDescription: "offset is an optional field and the value is added to the address of the attachment point function. If not provided, offset defaults to 0.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"pid": schema.Int64Attribute{
														Description:         "pid is an optional field and if provided, limits the execution of the UProbe or URetProbe to the provided process identification number (PID). If pid is not provided, the UProbe or URetProbe executes for all PIDs.",
														MarkdownDescription: "pid is an optional field and if provided, limits the execution of the UProbe or URetProbe to the provided process identification number (PID). If pid is not provided, the UProbe or URetProbe executes for all PIDs.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"target": schema.StringAttribute{
														Description:         "target is a required field and is the user-space library name or the absolute path to a binary or library.",
														MarkdownDescription: "target is a required field and is the user-space library name or the absolute path to a binary or library.",
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

								"xdp": schema.SingleNestedAttribute{
									Description:         "xdp is an optional field, but required when the type field is set to XDP. xdp defines the desired state of the application's XDP programs. XDP program can be attached to network devices (interfaces) and will be called on every incoming packet received by the network device. The XDP attachment point is just after the packet has been received off the wire, but before the Linux kernel has allocated an sk_buff, which is used to pass the packet through the kernel networking stack.",
									MarkdownDescription: "xdp is an optional field, but required when the type field is set to XDP. xdp defines the desired state of the application's XDP programs. XDP program can be attached to network devices (interfaces) and will be called on every incoming packet received by the network device. The XDP attachment point is just after the packet has been received off the wire, but before the Linux kernel has allocated an sk_buff, which is used to pass the packet through the kernel networking stack.",
									Attributes: map[string]schema.Attribute{
										"links": schema.ListNestedAttribute{
											Description:         "links is an optional field and is the list of attachment points to which the XDP program should be attached. The XDP program is loaded in kernel memory when the BPF Application CRD is created and the selected Kubernetes nodes are active. The XDP program will not be triggered until the program has also been attached to an attachment point described in this list. Items may be added or removed from the list at any point, causing the XDP program to be attached or detached. The attachment point for a XDP program is a network interface (or device). The interface can be specified by name, by allowing bpfman to discover each interface, or by setting the primaryNodeInterface flag, which instructs bpfman to use the primary interface of a Kubernetes node.",
											MarkdownDescription: "links is an optional field and is the list of attachment points to which the XDP program should be attached. The XDP program is loaded in kernel memory when the BPF Application CRD is created and the selected Kubernetes nodes are active. The XDP program will not be triggered until the program has also been attached to an attachment point described in this list. Items may be added or removed from the list at any point, causing the XDP program to be attached or detached. The attachment point for a XDP program is a network interface (or device). The interface can be specified by name, by allowing bpfman to discover each interface, or by setting the primaryNodeInterface flag, which instructs bpfman to use the primary interface of a Kubernetes node.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"interface_selector": schema.SingleNestedAttribute{
														Description:         "interfaceSelector is a required field and is used to determine the network interface (or interfaces) the XDP program is attached. Interface list is set by providing a list of interface names, enabling auto discovery, or setting the primaryNodeInterface flag, but only one option is allowed.",
														MarkdownDescription: "interfaceSelector is a required field and is used to determine the network interface (or interfaces) the XDP program is attached. Interface list is set by providing a list of interface names, enabling auto discovery, or setting the primaryNodeInterface flag, but only one option is allowed.",
														Attributes: map[string]schema.Attribute{
															"interfaces": schema.ListAttribute{
																Description:         "interfaces is an optional field and is a list of network interface names to attach the eBPF program. The interface names in the list are case-sensitive.",
																MarkdownDescription: "interfaces is an optional field and is a list of network interface names to attach the eBPF program. The interface names in the list are case-sensitive.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.List{
																	listvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("interfaces_discovery_config"), path.MatchRelative().AtParent().AtName("primary_node_interface")),
																},
															},

															"interfaces_discovery_config": schema.SingleNestedAttribute{
																Description:         "interfacesDiscoveryConfig is an optional field that is used to control if and how to automatically discover interfaces. If the agent should automatically discover and attach eBPF programs to interfaces, use the fields under interfacesDiscoveryConfig to control what is allow and excluded from discovery.",
																MarkdownDescription: "interfacesDiscoveryConfig is an optional field that is used to control if and how to automatically discover interfaces. If the agent should automatically discover and attach eBPF programs to interfaces, use the fields under interfacesDiscoveryConfig to control what is allow and excluded from discovery.",
																Attributes: map[string]schema.Attribute{
																	"allowed_interfaces": schema.ListAttribute{
																		Description:         "allowedInterfaces is an optional field that contains a list of interface names that are allowed to be discovered. If empty, the agent will fetch all the interfaces in the system, excepting the ones listed in excludeInterfaces. if non-empty, only entries in the list will be considered for discovery. If an entry enclosed by slashes, such as '/br-/' or '/veth*/', then the entry is considered as a regular expression for matching. Otherwise, the interface names in the list are case-sensitive. This field is only taken into consideration if interfaceAutoDiscovery is set to true.",
																		MarkdownDescription: "allowedInterfaces is an optional field that contains a list of interface names that are allowed to be discovered. If empty, the agent will fetch all the interfaces in the system, excepting the ones listed in excludeInterfaces. if non-empty, only entries in the list will be considered for discovery. If an entry enclosed by slashes, such as '/br-/' or '/veth*/', then the entry is considered as a regular expression for matching. Otherwise, the interface names in the list are case-sensitive. This field is only taken into consideration if interfaceAutoDiscovery is set to true.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"exclude_interfaces": schema.ListAttribute{
																		Description:         "excludeInterfaces is an optional field that contains a list of interface names that are excluded from interface discovery. The interface names in the list are case-sensitive. By default, the list contains the loopback interface, 'lo'. This field is only taken into consideration if interfaceAutoDiscovery is set to true.",
																		MarkdownDescription: "excludeInterfaces is an optional field that contains a list of interface names that are excluded from interface discovery. The interface names in the list are case-sensitive. By default, the list contains the loopback interface, 'lo'. This field is only taken into consideration if interfaceAutoDiscovery is set to true.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"interface_auto_discovery": schema.BoolAttribute{
																		Description:         "interfaceAutoDiscovery is an optional field. When enabled, the agent monitors the creation and deletion of interfaces and automatically attached eBPF programs to the newly discovered interfaces. CAUTION: This has the potential to attach a given eBPF program to a large number of interfaces. Use with caution.",
																		MarkdownDescription: "interfaceAutoDiscovery is an optional field. When enabled, the agent monitors the creation and deletion of interfaces and automatically attached eBPF programs to the newly discovered interfaces. CAUTION: This has the potential to attach a given eBPF program to a large number of interfaces. Use with caution.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
																Validators: []validator.Object{
																	objectvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("interfaces"), path.MatchRelative().AtParent().AtName("primary_node_interface")),
																},
															},

															"primary_node_interface": schema.BoolAttribute{
																Description:         "primaryNodeInterface is and optional field and indicates to attach the eBPF program to the primary interface on the Kubernetes node. Only 'true' is accepted.",
																MarkdownDescription: "primaryNodeInterface is and optional field and indicates to attach the eBPF program to the primary interface on the Kubernetes node. Only 'true' is accepted.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Bool{
																	boolvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("interfaces"), path.MatchRelative().AtParent().AtName("interfaces_discovery_config")),
																},
															},
														},
														Required: true,
														Optional: false,
														Computed: false,
													},

													"network_namespaces": schema.SingleNestedAttribute{
														Description:         "networkNamespaces is a required field that identifies the set of network namespaces in which to attach the eBPF program.",
														MarkdownDescription: "networkNamespaces is a required field that identifies the set of network namespaces in which to attach the eBPF program.",
														Attributes: map[string]schema.Attribute{
															"pods": schema.SingleNestedAttribute{
																Description:         "pods is a required field and indicates the target pods. To select all pods use the standard metav1.LabelSelector semantics and make it empty.",
																MarkdownDescription: "pods is a required field and indicates the target pods. To select all pods use the standard metav1.LabelSelector semantics and make it empty.",
																Attributes: map[string]schema.Attribute{
																	"match_expressions": schema.ListNestedAttribute{
																		Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																		MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"key": schema.StringAttribute{
																					Description:         "key is the label key that the selector applies to.",
																					MarkdownDescription: "key is the label key that the selector applies to.",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"operator": schema.StringAttribute{
																					Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																					MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"values": schema.ListAttribute{
																					Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																					MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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

																	"match_labels": schema.MapAttribute{
																		Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																		MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
														},
														Required: true,
														Optional: false,
														Computed: false,
													},

													"priority": schema.Int64Attribute{
														Description:         "priority is an optional field and determines the execution order of the XDP program relative to other XDP programs attached to the same attachment point. It must be a value between 0 and 1000, where lower values indicate higher precedence. For XDP programs on the same attachment point with the same priority, the most recently attached program has a lower precedence. If not provided, priority will default to 1000.",
														MarkdownDescription: "priority is an optional field and determines the execution order of the XDP program relative to other XDP programs attached to the same attachment point. It must be a value between 0 and 1000, where lower values indicate higher precedence. For XDP programs on the same attachment point with the same priority, the most recently attached program has a lower precedence. If not provided, priority will default to 1000.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(0),
															int64validator.AtMost(1000),
														},
													},

													"proceed_on": schema.ListAttribute{
														Description:         "proceedOn is an optional field and allows the user to call other XDP programs in a chain, or not call the next program in a chain based on the exit code of an XDP program. Allowed values, which are the possible exit codes from an XDP eBPF program, are: Aborted, Drop, Pass, TX, ReDirect,DispatcherReturn Multiple values are supported. Default is Pass and DispatcherReturn. So using the default values, if an XDP program returns Pass, the next XDP program in the chain will be called. If an XDP program returns Drop, the next XDP program in the chain will NOT be called.",
														MarkdownDescription: "proceedOn is an optional field and allows the user to call other XDP programs in a chain, or not call the next program in a chain based on the exit code of an XDP program. Allowed values, which are the possible exit codes from an XDP eBPF program, are: Aborted, Drop, Pass, TX, ReDirect,DispatcherReturn Multiple values are supported. Default is Pass and DispatcherReturn. So using the default values, if an XDP program returns Pass, the next XDP program in the chain will be called. If an XDP program returns Drop, the next XDP program in the chain will NOT be called.",
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
		},
	}
}

func (r *BpfmanIoBpfApplicationV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_bpfman_io_bpf_application_v1alpha1_manifest")

	var model BpfmanIoBpfApplicationV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("bpfman.io/v1alpha1")
	model.Kind = pointer.String("BpfApplication")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
