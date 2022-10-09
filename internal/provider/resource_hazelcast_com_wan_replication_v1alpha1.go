/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

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

type HazelcastComWanReplicationV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*HazelcastComWanReplicationV1Alpha1Resource)(nil)
)

type HazelcastComWanReplicationV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type HazelcastComWanReplicationV1Alpha1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		Acknowledgement *struct {
			Timeout *int64 `tfsdk:"timeout" yaml:"timeout,omitempty"`

			Type *string `tfsdk:"type" yaml:"type,omitempty"`
		} `tfsdk:"acknowledgement" yaml:"acknowledgement,omitempty"`

		Batch *struct {
			MaximumDelay *int64 `tfsdk:"maximum_delay" yaml:"maximumDelay,omitempty"`

			Size *int64 `tfsdk:"size" yaml:"size,omitempty"`
		} `tfsdk:"batch" yaml:"batch,omitempty"`

		Endpoints *string `tfsdk:"endpoints" yaml:"endpoints,omitempty"`

		Queue *struct {
			Capacity *int64 `tfsdk:"capacity" yaml:"capacity,omitempty"`

			FullBehavior *string `tfsdk:"full_behavior" yaml:"fullBehavior,omitempty"`
		} `tfsdk:"queue" yaml:"queue,omitempty"`

		Resources *[]struct {
			Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"resources" yaml:"resources,omitempty"`

		TargetClusterName *string `tfsdk:"target_cluster_name" yaml:"targetClusterName,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewHazelcastComWanReplicationV1Alpha1Resource() resource.Resource {
	return &HazelcastComWanReplicationV1Alpha1Resource{}
}

func (r *HazelcastComWanReplicationV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_hazelcast_com_wan_replication_v1alpha1"
}

func (r *HazelcastComWanReplicationV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "WanReplication is the Schema for the wanreplications API",
		MarkdownDescription: "WanReplication is the Schema for the wanreplications API",
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

					"namespace": {
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Type:                types.StringType,
						Optional:            true,
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
				Description:         "WanReplicationSpec defines the desired state of WanReplication",
				MarkdownDescription: "WanReplicationSpec defines the desired state of WanReplication",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"acknowledgement": {
						Description:         "Acknowledgement is the configuration for the condition when the next batch of WAN events are sent.",
						MarkdownDescription: "Acknowledgement is the configuration for the condition when the next batch of WAN events are sent.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"timeout": {
								Description:         "Timeout represents the time the source cluster waits for the acknowledgement. After timeout, the events will be resent.",
								MarkdownDescription: "Timeout represents the time the source cluster waits for the acknowledgement. After timeout, the events will be resent.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"type": {
								Description:         "Type represents how a batch of replication events is considered successfully replicated.",
								MarkdownDescription: "Type represents how a batch of replication events is considered successfully replicated.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("ACK_ON_OPERATION_COMPLETE", "ACK_ON_RECEIPT"),
								},
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"batch": {
						Description:         "Batch is the configuration for WAN events batch.",
						MarkdownDescription: "Batch is the configuration for WAN events batch.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"maximum_delay": {
								Description:         "MaximumDelay represents the maximum delay in milliseconds. If the batch size is not reached, the events will be sent after the maximum delay.",
								MarkdownDescription: "MaximumDelay represents the maximum delay in milliseconds. If the batch size is not reached, the events will be sent after the maximum delay.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"size": {
								Description:         "Size represents the maximum batch size.",
								MarkdownDescription: "Size represents the maximum batch size.",

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

					"endpoints": {
						Description:         "Endpoints is the target cluster endpoints.",
						MarkdownDescription: "Endpoints is the target cluster endpoints.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.LengthAtLeast(1),
						},
					},

					"queue": {
						Description:         "Queue is the configuration for WAN events queue.",
						MarkdownDescription: "Queue is the configuration for WAN events queue.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"capacity": {
								Description:         "Capacity is the total capacity of WAN queue.",
								MarkdownDescription: "Capacity is the total capacity of WAN queue.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"full_behavior": {
								Description:         "FullBehavior represents the behavior of the new arrival when the queue is full.",
								MarkdownDescription: "FullBehavior represents the behavior of the new arrival when the queue is full.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("DISCARD_AFTER_MUTATION", "THROW_EXCEPTION", "THROW_EXCEPTION_ONLY_IF_REPLICATION_ACTIVE"),
								},
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"resources": {
						Description:         "Resources is the list of custom resources to which WAN replication applies.",
						MarkdownDescription: "Resources is the list of custom resources to which WAN replication applies.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"kind": {
								Description:         "Kind is the kind of custom resource to which WAN replication applies.",
								MarkdownDescription: "Kind is the kind of custom resource to which WAN replication applies.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("Map", "Hazelcast"),
								},
							},

							"name": {
								Description:         "Name is the name of custom resource to which WAN replication applies.",
								MarkdownDescription: "Name is the name of custom resource to which WAN replication applies.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.LengthAtLeast(1),
								},
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"target_cluster_name": {
						Description:         "ClusterName is the clusterName field of the target Hazelcast resource.",
						MarkdownDescription: "ClusterName is the clusterName field of the target Hazelcast resource.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.LengthAtLeast(1),
						},
					},
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *HazelcastComWanReplicationV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_hazelcast_com_wan_replication_v1alpha1")

	var state HazelcastComWanReplicationV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel HazelcastComWanReplicationV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("hazelcast.com/v1alpha1")
	goModel.Kind = utilities.Ptr("WanReplication")

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

func (r *HazelcastComWanReplicationV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_hazelcast_com_wan_replication_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *HazelcastComWanReplicationV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_hazelcast_com_wan_replication_v1alpha1")

	var state HazelcastComWanReplicationV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel HazelcastComWanReplicationV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("hazelcast.com/v1alpha1")
	goModel.Kind = utilities.Ptr("WanReplication")

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

func (r *HazelcastComWanReplicationV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_hazelcast_com_wan_replication_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
