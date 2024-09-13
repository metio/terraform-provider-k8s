/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package karpenter_sh_v1beta1

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
	_ datasource.DataSource = &KarpenterShNodeClaimV1Beta1Manifest{}
)

func NewKarpenterShNodeClaimV1Beta1Manifest() datasource.DataSource {
	return &KarpenterShNodeClaimV1Beta1Manifest{}
}

type KarpenterShNodeClaimV1Beta1Manifest struct{}

type KarpenterShNodeClaimV1Beta1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Kubelet *struct {
			ClusterDNS                  *[]string          `tfsdk:"cluster_dns" json:"clusterDNS,omitempty"`
			CpuCFSQuota                 *bool              `tfsdk:"cpu_cfs_quota" json:"cpuCFSQuota,omitempty"`
			EvictionHard                *map[string]string `tfsdk:"eviction_hard" json:"evictionHard,omitempty"`
			EvictionMaxPodGracePeriod   *int64             `tfsdk:"eviction_max_pod_grace_period" json:"evictionMaxPodGracePeriod,omitempty"`
			EvictionSoft                *map[string]string `tfsdk:"eviction_soft" json:"evictionSoft,omitempty"`
			EvictionSoftGracePeriod     *map[string]string `tfsdk:"eviction_soft_grace_period" json:"evictionSoftGracePeriod,omitempty"`
			ImageGCHighThresholdPercent *int64             `tfsdk:"image_gc_high_threshold_percent" json:"imageGCHighThresholdPercent,omitempty"`
			ImageGCLowThresholdPercent  *int64             `tfsdk:"image_gc_low_threshold_percent" json:"imageGCLowThresholdPercent,omitempty"`
			KubeReserved                *map[string]string `tfsdk:"kube_reserved" json:"kubeReserved,omitempty"`
			MaxPods                     *int64             `tfsdk:"max_pods" json:"maxPods,omitempty"`
			PodsPerCore                 *int64             `tfsdk:"pods_per_core" json:"podsPerCore,omitempty"`
			SystemReserved              *map[string]string `tfsdk:"system_reserved" json:"systemReserved,omitempty"`
		} `tfsdk:"kubelet" json:"kubelet,omitempty"`
		NodeClassRef *struct {
			ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
			Name       *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"node_class_ref" json:"nodeClassRef,omitempty"`
		Requirements *[]struct {
			Key       *string   `tfsdk:"key" json:"key,omitempty"`
			MinValues *int64    `tfsdk:"min_values" json:"minValues,omitempty"`
			Operator  *string   `tfsdk:"operator" json:"operator,omitempty"`
			Values    *[]string `tfsdk:"values" json:"values,omitempty"`
		} `tfsdk:"requirements" json:"requirements,omitempty"`
		Resources *struct {
			Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
		} `tfsdk:"resources" json:"resources,omitempty"`
		StartupTaints *[]struct {
			Effect    *string `tfsdk:"effect" json:"effect,omitempty"`
			Key       *string `tfsdk:"key" json:"key,omitempty"`
			TimeAdded *string `tfsdk:"time_added" json:"timeAdded,omitempty"`
			Value     *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"startup_taints" json:"startupTaints,omitempty"`
		Taints *[]struct {
			Effect    *string `tfsdk:"effect" json:"effect,omitempty"`
			Key       *string `tfsdk:"key" json:"key,omitempty"`
			TimeAdded *string `tfsdk:"time_added" json:"timeAdded,omitempty"`
			Value     *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"taints" json:"taints,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KarpenterShNodeClaimV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_karpenter_sh_node_claim_v1beta1_manifest"
}

func (r *KarpenterShNodeClaimV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "NodeClaim is the Schema for the NodeClaims API",
		MarkdownDescription: "NodeClaim is the Schema for the NodeClaims API",
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
				Description:         "NodeClaimSpec describes the desired state of the NodeClaim",
				MarkdownDescription: "NodeClaimSpec describes the desired state of the NodeClaim",
				Attributes: map[string]schema.Attribute{
					"kubelet": schema.SingleNestedAttribute{
						Description:         "Kubelet defines args to be used when configuring kubelet on provisioned nodes. They are a subset of the upstream types, recognizing not all options may be supported. Wherever possible, the types and names should reflect the upstream kubelet types.",
						MarkdownDescription: "Kubelet defines args to be used when configuring kubelet on provisioned nodes. They are a subset of the upstream types, recognizing not all options may be supported. Wherever possible, the types and names should reflect the upstream kubelet types.",
						Attributes: map[string]schema.Attribute{
							"cluster_dns": schema.ListAttribute{
								Description:         "clusterDNS is a list of IP addresses for the cluster DNS server. Note that not all providers may use all addresses.",
								MarkdownDescription: "clusterDNS is a list of IP addresses for the cluster DNS server. Note that not all providers may use all addresses.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cpu_cfs_quota": schema.BoolAttribute{
								Description:         "CPUCFSQuota enables CPU CFS quota enforcement for containers that specify CPU limits.",
								MarkdownDescription: "CPUCFSQuota enables CPU CFS quota enforcement for containers that specify CPU limits.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"eviction_hard": schema.MapAttribute{
								Description:         "EvictionHard is the map of signal names to quantities that define hard eviction thresholds",
								MarkdownDescription: "EvictionHard is the map of signal names to quantities that define hard eviction thresholds",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"eviction_max_pod_grace_period": schema.Int64Attribute{
								Description:         "EvictionMaxPodGracePeriod is the maximum allowed grace period (in seconds) to use when terminating pods in response to soft eviction thresholds being met.",
								MarkdownDescription: "EvictionMaxPodGracePeriod is the maximum allowed grace period (in seconds) to use when terminating pods in response to soft eviction thresholds being met.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"eviction_soft": schema.MapAttribute{
								Description:         "EvictionSoft is the map of signal names to quantities that define soft eviction thresholds",
								MarkdownDescription: "EvictionSoft is the map of signal names to quantities that define soft eviction thresholds",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"eviction_soft_grace_period": schema.MapAttribute{
								Description:         "EvictionSoftGracePeriod is the map of signal names to quantities that define grace periods for each eviction signal",
								MarkdownDescription: "EvictionSoftGracePeriod is the map of signal names to quantities that define grace periods for each eviction signal",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image_gc_high_threshold_percent": schema.Int64Attribute{
								Description:         "ImageGCHighThresholdPercent is the percent of disk usage after which image garbage collection is always run. The percent is calculated by dividing this field value by 100, so this field must be between 0 and 100, inclusive. When specified, the value must be greater than ImageGCLowThresholdPercent.",
								MarkdownDescription: "ImageGCHighThresholdPercent is the percent of disk usage after which image garbage collection is always run. The percent is calculated by dividing this field value by 100, so this field must be between 0 and 100, inclusive. When specified, the value must be greater than ImageGCLowThresholdPercent.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
									int64validator.AtMost(100),
								},
							},

							"image_gc_low_threshold_percent": schema.Int64Attribute{
								Description:         "ImageGCLowThresholdPercent is the percent of disk usage before which image garbage collection is never run. Lowest disk usage to garbage collect to. The percent is calculated by dividing this field value by 100, so the field value must be between 0 and 100, inclusive. When specified, the value must be less than imageGCHighThresholdPercent",
								MarkdownDescription: "ImageGCLowThresholdPercent is the percent of disk usage before which image garbage collection is never run. Lowest disk usage to garbage collect to. The percent is calculated by dividing this field value by 100, so the field value must be between 0 and 100, inclusive. When specified, the value must be less than imageGCHighThresholdPercent",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
									int64validator.AtMost(100),
								},
							},

							"kube_reserved": schema.MapAttribute{
								Description:         "KubeReserved contains resources reserved for Kubernetes system components.",
								MarkdownDescription: "KubeReserved contains resources reserved for Kubernetes system components.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_pods": schema.Int64Attribute{
								Description:         "MaxPods is an override for the maximum number of pods that can run on a worker node instance.",
								MarkdownDescription: "MaxPods is an override for the maximum number of pods that can run on a worker node instance.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
								},
							},

							"pods_per_core": schema.Int64Attribute{
								Description:         "PodsPerCore is an override for the number of pods that can run on a worker node instance based on the number of cpu cores. This value cannot exceed MaxPods, so, if MaxPods is a lower value, that value will be used.",
								MarkdownDescription: "PodsPerCore is an override for the number of pods that can run on a worker node instance based on the number of cpu cores. This value cannot exceed MaxPods, so, if MaxPods is a lower value, that value will be used.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
								},
							},

							"system_reserved": schema.MapAttribute{
								Description:         "SystemReserved contains resources reserved for OS system daemons and kernel memory.",
								MarkdownDescription: "SystemReserved contains resources reserved for OS system daemons and kernel memory.",
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

					"node_class_ref": schema.SingleNestedAttribute{
						Description:         "NodeClassRef is a reference to an object that defines provider specific configuration",
						MarkdownDescription: "NodeClassRef is a reference to an object that defines provider specific configuration",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "API version of the referent",
								MarkdownDescription: "API version of the referent",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'",
								MarkdownDescription: "Kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referent; More info: http://kubernetes.io/docs/user-guide/identifiers#names",
								MarkdownDescription: "Name of the referent; More info: http://kubernetes.io/docs/user-guide/identifiers#names",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"requirements": schema.ListNestedAttribute{
						Description:         "Requirements are layered with GetLabels and applied to every node.",
						MarkdownDescription: "Requirements are layered with GetLabels and applied to every node.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"key": schema.StringAttribute{
									Description:         "The label key that the selector applies to.",
									MarkdownDescription: "The label key that the selector applies to.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(316),
										stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*(\/))?([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9]$`), ""),
									},
								},

								"min_values": schema.Int64Attribute{
									Description:         "This field is ALPHA and can be dropped or replaced at any time MinValues is the minimum number of unique values required to define the flexibility of the specific requirement.",
									MarkdownDescription: "This field is ALPHA and can be dropped or replaced at any time MinValues is the minimum number of unique values required to define the flexibility of the specific requirement.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.Int64{
										int64validator.AtLeast(1),
										int64validator.AtMost(50),
									},
								},

								"operator": schema.StringAttribute{
									Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
									MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist", "Gt", "Lt"),
									},
								},

								"values": schema.ListAttribute{
									Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
									MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
									ElementType:         types.StringType,
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

					"resources": schema.SingleNestedAttribute{
						Description:         "Resources models the resource requirements for the NodeClaim to launch",
						MarkdownDescription: "Resources models the resource requirements for the NodeClaim to launch",
						Attributes: map[string]schema.Attribute{
							"requests": schema.MapAttribute{
								Description:         "Requests describes the minimum required resources for the NodeClaim to launch",
								MarkdownDescription: "Requests describes the minimum required resources for the NodeClaim to launch",
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

					"startup_taints": schema.ListNestedAttribute{
						Description:         "StartupTaints are taints that are applied to nodes upon startup which are expected to be removed automatically within a short period of time, typically by a DaemonSet that tolerates the taint. These are commonly used by daemonsets to allow initialization and enforce startup ordering. StartupTaints are ignored for provisioning purposes in that pods are not required to tolerate a StartupTaint in order to have nodes provisioned for them.",
						MarkdownDescription: "StartupTaints are taints that are applied to nodes upon startup which are expected to be removed automatically within a short period of time, typically by a DaemonSet that tolerates the taint. These are commonly used by daemonsets to allow initialization and enforce startup ordering. StartupTaints are ignored for provisioning purposes in that pods are not required to tolerate a StartupTaint in order to have nodes provisioned for them.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"effect": schema.StringAttribute{
									Description:         "Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.",
									MarkdownDescription: "Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("NoSchedule", "PreferNoSchedule", "NoExecute"),
									},
								},

								"key": schema.StringAttribute{
									Description:         "Required. The taint key to be applied to a node.",
									MarkdownDescription: "Required. The taint key to be applied to a node.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*(\/))?([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9]$`), ""),
									},
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
									Validators: []validator.String{
										stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*(\/))?([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9]$`), ""),
									},
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"taints": schema.ListNestedAttribute{
						Description:         "Taints will be applied to the NodeClaim's node.",
						MarkdownDescription: "Taints will be applied to the NodeClaim's node.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"effect": schema.StringAttribute{
									Description:         "Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.",
									MarkdownDescription: "Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("NoSchedule", "PreferNoSchedule", "NoExecute"),
									},
								},

								"key": schema.StringAttribute{
									Description:         "Required. The taint key to be applied to a node.",
									MarkdownDescription: "Required. The taint key to be applied to a node.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*(\/))?([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9]$`), ""),
									},
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
									Validators: []validator.String{
										stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*(\/))?([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9]$`), ""),
									},
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
	}
}

func (r *KarpenterShNodeClaimV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_karpenter_sh_node_claim_v1beta1_manifest")

	var model KarpenterShNodeClaimV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("karpenter.sh/v1beta1")
	model.Kind = pointer.String("NodeClaim")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
