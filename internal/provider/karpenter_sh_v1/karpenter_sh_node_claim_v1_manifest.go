/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package karpenter_sh_v1

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
	_ datasource.DataSource = &KarpenterShNodeClaimV1Manifest{}
)

func NewKarpenterShNodeClaimV1Manifest() datasource.DataSource {
	return &KarpenterShNodeClaimV1Manifest{}
}

type KarpenterShNodeClaimV1Manifest struct{}

type KarpenterShNodeClaimV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		ExpireAfter  *string `tfsdk:"expire_after" json:"expireAfter,omitempty"`
		NodeClassRef *struct {
			Group *string `tfsdk:"group" json:"group,omitempty"`
			Kind  *string `tfsdk:"kind" json:"kind,omitempty"`
			Name  *string `tfsdk:"name" json:"name,omitempty"`
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
		TerminationGracePeriod *string `tfsdk:"termination_grace_period" json:"terminationGracePeriod,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KarpenterShNodeClaimV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_karpenter_sh_node_claim_v1_manifest"
}

func (r *KarpenterShNodeClaimV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
					"expire_after": schema.StringAttribute{
						Description:         "ExpireAfter is the duration the controller will wait before terminating a node, measured from when the node is created. This is useful to implement features like eventually consistent node upgrade, memory leak protection, and disruption testing.",
						MarkdownDescription: "ExpireAfter is the duration the controller will wait before terminating a node, measured from when the node is created. This is useful to implement features like eventually consistent node upgrade, memory leak protection, and disruption testing.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(([0-9]+(s|m|h))+)|(Never)$`), ""),
						},
					},

					"node_class_ref": schema.SingleNestedAttribute{
						Description:         "NodeClassRef is a reference to an object that defines provider specific configuration",
						MarkdownDescription: "NodeClassRef is a reference to an object that defines provider specific configuration",
						Attributes: map[string]schema.Attribute{
							"group": schema.StringAttribute{
								Description:         "API version of the referent",
								MarkdownDescription: "API version of the referent",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^[^/]*$`), ""),
								},
							},

							"kind": schema.StringAttribute{
								Description:         "Kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'",
								MarkdownDescription: "Kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'",
								Required:            true,
								Optional:            false,
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

					"termination_grace_period": schema.StringAttribute{
						Description:         "TerminationGracePeriod is the maximum duration the controller will wait before forcefully deleting the pods on a node, measured from when deletion is first initiated. Warning: this feature takes precedence over a Pod's terminationGracePeriodSeconds value, and bypasses any blocked PDBs or the karpenter.sh/do-not-disrupt annotation. This field is intended to be used by cluster administrators to enforce that nodes can be cycled within a given time period. When set, drifted nodes will begin draining even if there are pods blocking eviction. Draining will respect PDBs and the do-not-disrupt annotation until the TGP is reached. Karpenter will preemptively delete pods so their terminationGracePeriodSeconds align with the node's terminationGracePeriod. If a pod would be terminated without being granted its full terminationGracePeriodSeconds prior to the node timeout, that pod will be deleted at T = node timeout - pod terminationGracePeriodSeconds. The feature can also be used to allow maximum time limits for long-running jobs which can delay node termination with preStop hooks. If left undefined, the controller will wait indefinitely for pods to be drained.",
						MarkdownDescription: "TerminationGracePeriod is the maximum duration the controller will wait before forcefully deleting the pods on a node, measured from when deletion is first initiated. Warning: this feature takes precedence over a Pod's terminationGracePeriodSeconds value, and bypasses any blocked PDBs or the karpenter.sh/do-not-disrupt annotation. This field is intended to be used by cluster administrators to enforce that nodes can be cycled within a given time period. When set, drifted nodes will begin draining even if there are pods blocking eviction. Draining will respect PDBs and the do-not-disrupt annotation until the TGP is reached. Karpenter will preemptively delete pods so their terminationGracePeriodSeconds align with the node's terminationGracePeriod. If a pod would be terminated without being granted its full terminationGracePeriodSeconds prior to the node timeout, that pod will be deleted at T = node timeout - pod terminationGracePeriodSeconds. The feature can also be used to allow maximum time limits for long-running jobs which can delay node termination with preStop hooks. If left undefined, the controller will wait indefinitely for pods to be drained.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(s|m|h))+$`), ""),
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

func (r *KarpenterShNodeClaimV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_karpenter_sh_node_claim_v1_manifest")

	var model KarpenterShNodeClaimV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("karpenter.sh/v1")
	model.Kind = pointer.String("NodeClaim")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
