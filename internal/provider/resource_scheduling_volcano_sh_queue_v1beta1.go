/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"gopkg.in/yaml.v3"
	"time"
)

type SchedulingVolcanoShQueueV1Beta1Resource struct{}

var (
	_ resource.Resource = (*SchedulingVolcanoShQueueV1Beta1Resource)(nil)
)

type SchedulingVolcanoShQueueV1Beta1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type SchedulingVolcanoShQueueV1Beta1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		Affinity *struct {
			NodeGroupAffinity *struct {
				PreferredDuringSchedulingIgnoredDuringExecution *[]string `tfsdk:"preferred_during_scheduling_ignored_during_execution" yaml:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`

				RequiredDuringSchedulingIgnoredDuringExecution *[]string `tfsdk:"required_during_scheduling_ignored_during_execution" yaml:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
			} `tfsdk:"node_group_affinity" yaml:"nodeGroupAffinity,omitempty"`

			NodeGroupAntiAffinity *struct {
				PreferredDuringSchedulingIgnoredDuringExecution *[]string `tfsdk:"preferred_during_scheduling_ignored_during_execution" yaml:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`

				RequiredDuringSchedulingIgnoredDuringExecution *[]string `tfsdk:"required_during_scheduling_ignored_during_execution" yaml:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
			} `tfsdk:"node_group_anti_affinity" yaml:"nodeGroupAntiAffinity,omitempty"`
		} `tfsdk:"affinity" yaml:"affinity,omitempty"`

		Capability *map[string]string `tfsdk:"capability" yaml:"capability,omitempty"`

		ExtendClusters *[]struct {
			Capacity *map[string]string `tfsdk:"capacity" yaml:"capacity,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Weight *int64 `tfsdk:"weight" yaml:"weight,omitempty"`
		} `tfsdk:"extend_clusters" yaml:"extendClusters,omitempty"`

		Guarantee *struct {
			Resource *map[string]string `tfsdk:"resource" yaml:"resource,omitempty"`
		} `tfsdk:"guarantee" yaml:"guarantee,omitempty"`

		Reclaimable *bool `tfsdk:"reclaimable" yaml:"reclaimable,omitempty"`

		Type *string `tfsdk:"type" yaml:"type,omitempty"`

		Weight *int64 `tfsdk:"weight" yaml:"weight,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewSchedulingVolcanoShQueueV1Beta1Resource() resource.Resource {
	return &SchedulingVolcanoShQueueV1Beta1Resource{}
}

func (r *SchedulingVolcanoShQueueV1Beta1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_scheduling_volcano_sh_queue_v1beta1"
}

func (r *SchedulingVolcanoShQueueV1Beta1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Queue is a queue of PodGroup.",
		MarkdownDescription: "Queue is a queue of PodGroup.",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description:         "The timestamp of the last change to this resource.",
				MarkdownDescription: "The timestamp of the last change to this resource.",
				Type:                types.Int64Type,
				Computed:            true,
				Optional:            false,
			},

			"yaml": {
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"metadata": {
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Type:                types.StringType,
						Required:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
					},

					"labels": {
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.LabelValidator(),
						},
					},
					"annotations": {
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.AnnotationValidator(),
						},
					},
				}),
			},

			"api_version": {
				Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"kind": {
				Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"spec": {
				Description:         "Specification of the desired behavior of the queue. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status",
				MarkdownDescription: "Specification of the desired behavior of the queue. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"affinity": {
						Description:         "If specified, the pod owned by the queue will be scheduled with constraint",
						MarkdownDescription: "If specified, the pod owned by the queue will be scheduled with constraint",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"node_group_affinity": {
								Description:         "Describes nodegroup affinity scheduling rules for the queue(e.g. putting pods of the queue in the nodes of the nodegroup)",
								MarkdownDescription: "Describes nodegroup affinity scheduling rules for the queue(e.g. putting pods of the queue in the nodes of the nodegroup)",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"preferred_during_scheduling_ignored_during_execution": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"required_during_scheduling_ignored_during_execution": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"node_group_anti_affinity": {
								Description:         "Describes nodegroup anti-affinity scheduling rules for the queue(e.g. avoid putting pods of the queue in the nodes of the nodegroup).",
								MarkdownDescription: "Describes nodegroup anti-affinity scheduling rules for the queue(e.g. avoid putting pods of the queue in the nodes of the nodegroup).",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"preferred_during_scheduling_ignored_during_execution": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"required_during_scheduling_ignored_during_execution": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"capability": {
						Description:         "ResourceList is a set of (resource name, quantity) pairs.",
						MarkdownDescription: "ResourceList is a set of (resource name, quantity) pairs.",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"extend_clusters": {
						Description:         "extendCluster indicate the jobs in this Queue will be dispatched to these clusters.",
						MarkdownDescription: "extendCluster indicate the jobs in this Queue will be dispatched to these clusters.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"capacity": {
								Description:         "ResourceList is a set of (resource name, quantity) pairs.",
								MarkdownDescription: "ResourceList is a set of (resource name, quantity) pairs.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"weight": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"guarantee": {
						Description:         "Guarantee indicate configuration about resource reservation",
						MarkdownDescription: "Guarantee indicate configuration about resource reservation",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"resource": {
								Description:         "The amount of cluster resource reserved for queue. Just set either 'percentage' or 'resource'",
								MarkdownDescription: "The amount of cluster resource reserved for queue. Just set either 'percentage' or 'resource'",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"reclaimable": {
						Description:         "Reclaimable indicate whether the queue can be reclaimed by other queue",
						MarkdownDescription: "Reclaimable indicate whether the queue can be reclaimed by other queue",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"type": {
						Description:         "Type define the type of queue",
						MarkdownDescription: "Type define the type of queue",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"weight": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *SchedulingVolcanoShQueueV1Beta1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_scheduling_volcano_sh_queue_v1beta1")

	var state SchedulingVolcanoShQueueV1Beta1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel SchedulingVolcanoShQueueV1Beta1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("scheduling.volcano.sh/v1beta1")
	goModel.Kind = utilities.Ptr("Queue")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *SchedulingVolcanoShQueueV1Beta1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_scheduling_volcano_sh_queue_v1beta1")
	// NO-OP: All data is already in Terraform state
}

func (r *SchedulingVolcanoShQueueV1Beta1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_scheduling_volcano_sh_queue_v1beta1")

	var state SchedulingVolcanoShQueueV1Beta1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel SchedulingVolcanoShQueueV1Beta1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("scheduling.volcano.sh/v1beta1")
	goModel.Kind = utilities.Ptr("Queue")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *SchedulingVolcanoShQueueV1Beta1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_scheduling_volcano_sh_queue_v1beta1")
	// NO-OP: Terraform removes the state automatically for us
}
