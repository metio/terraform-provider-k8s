/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package hazelcast_com_v1alpha1

import (
	"context"
	"fmt"
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
	_ datasource.DataSource = &HazelcastComWanReplicationV1Alpha1Manifest{}
)

func NewHazelcastComWanReplicationV1Alpha1Manifest() datasource.DataSource {
	return &HazelcastComWanReplicationV1Alpha1Manifest{}
}

type HazelcastComWanReplicationV1Alpha1Manifest struct{}

type HazelcastComWanReplicationV1Alpha1ManifestData struct {
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
		Acknowledgement *struct {
			Timeout *int64  `tfsdk:"timeout" json:"timeout,omitempty"`
			Type    *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"acknowledgement" json:"acknowledgement,omitempty"`
		Batch *struct {
			MaximumDelay *int64 `tfsdk:"maximum_delay" json:"maximumDelay,omitempty"`
			Size         *int64 `tfsdk:"size" json:"size,omitempty"`
		} `tfsdk:"batch" json:"batch,omitempty"`
		Endpoints *string `tfsdk:"endpoints" json:"endpoints,omitempty"`
		Queue     *struct {
			Capacity     *int64  `tfsdk:"capacity" json:"capacity,omitempty"`
			FullBehavior *string `tfsdk:"full_behavior" json:"fullBehavior,omitempty"`
		} `tfsdk:"queue" json:"queue,omitempty"`
		Resources *[]struct {
			Kind *string `tfsdk:"kind" json:"kind,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"resources" json:"resources,omitempty"`
		TargetClusterName *string `tfsdk:"target_cluster_name" json:"targetClusterName,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *HazelcastComWanReplicationV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_hazelcast_com_wan_replication_v1alpha1_manifest"
}

func (r *HazelcastComWanReplicationV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "WanReplication is the Schema for the wanreplications API",
		MarkdownDescription: "WanReplication is the Schema for the wanreplications API",
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
				Description:         "WanReplicationSpec defines the desired state of WanReplication",
				MarkdownDescription: "WanReplicationSpec defines the desired state of WanReplication",
				Attributes: map[string]schema.Attribute{
					"acknowledgement": schema.SingleNestedAttribute{
						Description:         "Acknowledgement is the configuration for the condition when the next batch of WAN events are sent.",
						MarkdownDescription: "Acknowledgement is the configuration for the condition when the next batch of WAN events are sent.",
						Attributes: map[string]schema.Attribute{
							"timeout": schema.Int64Attribute{
								Description:         "Timeout represents the time in milliseconds the source cluster waits for the acknowledgement. After timeout, the events will be resent.",
								MarkdownDescription: "Timeout represents the time in milliseconds the source cluster waits for the acknowledgement. After timeout, the events will be resent.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "Type represents how a batch of replication events is considered successfully replicated.",
								MarkdownDescription: "Type represents how a batch of replication events is considered successfully replicated.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("ACK_ON_OPERATION_COMPLETE", "ACK_ON_RECEIPT"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"batch": schema.SingleNestedAttribute{
						Description:         "Batch is the configuration for WAN events batch.",
						MarkdownDescription: "Batch is the configuration for WAN events batch.",
						Attributes: map[string]schema.Attribute{
							"maximum_delay": schema.Int64Attribute{
								Description:         "MaximumDelay represents the maximum delay in milliseconds. If the batch size is not reached, the events will be sent after the maximum delay.",
								MarkdownDescription: "MaximumDelay represents the maximum delay in milliseconds. If the batch size is not reached, the events will be sent after the maximum delay.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"size": schema.Int64Attribute{
								Description:         "Size represents the maximum batch size.",
								MarkdownDescription: "Size represents the maximum batch size.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"endpoints": schema.StringAttribute{
						Description:         "Endpoints is the target cluster comma separated endpoint list .",
						MarkdownDescription: "Endpoints is the target cluster comma separated endpoint list .",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
						},
					},

					"queue": schema.SingleNestedAttribute{
						Description:         "Queue is the configuration for WAN events queue.",
						MarkdownDescription: "Queue is the configuration for WAN events queue.",
						Attributes: map[string]schema.Attribute{
							"capacity": schema.Int64Attribute{
								Description:         "Capacity is the total capacity of WAN queue.",
								MarkdownDescription: "Capacity is the total capacity of WAN queue.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"full_behavior": schema.StringAttribute{
								Description:         "FullBehavior represents the behavior of the new arrival when the queue is full.",
								MarkdownDescription: "FullBehavior represents the behavior of the new arrival when the queue is full.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("DISCARD_AFTER_MUTATION", "THROW_EXCEPTION", "THROW_EXCEPTION_ONLY_IF_REPLICATION_ACTIVE"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"resources": schema.ListNestedAttribute{
						Description:         "Resources is the list of custom resources to which WAN replication applies.",
						MarkdownDescription: "Resources is the list of custom resources to which WAN replication applies.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"kind": schema.StringAttribute{
									Description:         "Kind is the kind of custom resource to which WAN replication applies.",
									MarkdownDescription: "Kind is the kind of custom resource to which WAN replication applies.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("Map", "Hazelcast"),
									},
								},

								"name": schema.StringAttribute{
									Description:         "Name is the name of custom resource to which WAN replication applies.",
									MarkdownDescription: "Name is the name of custom resource to which WAN replication applies.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
									},
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"target_cluster_name": schema.StringAttribute{
						Description:         "ClusterName is the clusterName field of the target Hazelcast resource.",
						MarkdownDescription: "ClusterName is the clusterName field of the target Hazelcast resource.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
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

func (r *HazelcastComWanReplicationV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_hazelcast_com_wan_replication_v1alpha1_manifest")

	var model HazelcastComWanReplicationV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("hazelcast.com/v1alpha1")
	model.Kind = pointer.String("WanReplication")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
