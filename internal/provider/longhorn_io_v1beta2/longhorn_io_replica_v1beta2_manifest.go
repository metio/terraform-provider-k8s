/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package longhorn_io_v1beta2

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
	_ datasource.DataSource = &LonghornIoReplicaV1Beta2Manifest{}
)

func NewLonghornIoReplicaV1Beta2Manifest() datasource.DataSource {
	return &LonghornIoReplicaV1Beta2Manifest{}
}

type LonghornIoReplicaV1Beta2Manifest struct{}

type LonghornIoReplicaV1Beta2ManifestData struct {
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
		Active                           *bool   `tfsdk:"active" json:"active,omitempty"`
		BackendStoreDriver               *string `tfsdk:"backend_store_driver" json:"backendStoreDriver,omitempty"`
		BackingImage                     *string `tfsdk:"backing_image" json:"backingImage,omitempty"`
		DataDirectoryName                *string `tfsdk:"data_directory_name" json:"dataDirectoryName,omitempty"`
		DataEngine                       *string `tfsdk:"data_engine" json:"dataEngine,omitempty"`
		DesireState                      *string `tfsdk:"desire_state" json:"desireState,omitempty"`
		DiskID                           *string `tfsdk:"disk_id" json:"diskID,omitempty"`
		DiskPath                         *string `tfsdk:"disk_path" json:"diskPath,omitempty"`
		EngineImage                      *string `tfsdk:"engine_image" json:"engineImage,omitempty"`
		EngineName                       *string `tfsdk:"engine_name" json:"engineName,omitempty"`
		EvictionRequested                *bool   `tfsdk:"eviction_requested" json:"evictionRequested,omitempty"`
		FailedAt                         *string `tfsdk:"failed_at" json:"failedAt,omitempty"`
		HardNodeAffinity                 *string `tfsdk:"hard_node_affinity" json:"hardNodeAffinity,omitempty"`
		HealthyAt                        *string `tfsdk:"healthy_at" json:"healthyAt,omitempty"`
		Image                            *string `tfsdk:"image" json:"image,omitempty"`
		LastFailedAt                     *string `tfsdk:"last_failed_at" json:"lastFailedAt,omitempty"`
		LastHealthyAt                    *string `tfsdk:"last_healthy_at" json:"lastHealthyAt,omitempty"`
		LogRequested                     *bool   `tfsdk:"log_requested" json:"logRequested,omitempty"`
		NodeID                           *string `tfsdk:"node_id" json:"nodeID,omitempty"`
		RebuildRetryCount                *int64  `tfsdk:"rebuild_retry_count" json:"rebuildRetryCount,omitempty"`
		RevisionCounterDisabled          *bool   `tfsdk:"revision_counter_disabled" json:"revisionCounterDisabled,omitempty"`
		SalvageRequested                 *bool   `tfsdk:"salvage_requested" json:"salvageRequested,omitempty"`
		SnapshotMaxCount                 *int64  `tfsdk:"snapshot_max_count" json:"snapshotMaxCount,omitempty"`
		SnapshotMaxSize                  *string `tfsdk:"snapshot_max_size" json:"snapshotMaxSize,omitempty"`
		UnmapMarkDiskChainRemovedEnabled *bool   `tfsdk:"unmap_mark_disk_chain_removed_enabled" json:"unmapMarkDiskChainRemovedEnabled,omitempty"`
		VolumeName                       *string `tfsdk:"volume_name" json:"volumeName,omitempty"`
		VolumeSize                       *string `tfsdk:"volume_size" json:"volumeSize,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *LonghornIoReplicaV1Beta2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_longhorn_io_replica_v1beta2_manifest"
}

func (r *LonghornIoReplicaV1Beta2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Replica is where Longhorn stores replica object.",
		MarkdownDescription: "Replica is where Longhorn stores replica object.",
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
				Description:         "ReplicaSpec defines the desired state of the Longhorn replica",
				MarkdownDescription: "ReplicaSpec defines the desired state of the Longhorn replica",
				Attributes: map[string]schema.Attribute{
					"active": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"backend_store_driver": schema.StringAttribute{
						Description:         "Deprecated: Replaced by field 'dataEngine'.",
						MarkdownDescription: "Deprecated: Replaced by field 'dataEngine'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"backing_image": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"data_directory_name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"data_engine": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("v1", "v2"),
						},
					},

					"desire_state": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"disk_id": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"disk_path": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"engine_image": schema.StringAttribute{
						Description:         "Deprecated: Replaced by field 'image'.",
						MarkdownDescription: "Deprecated: Replaced by field 'image'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"engine_name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"eviction_requested": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"failed_at": schema.StringAttribute{
						Description:         "FailedAt is set when a running replica fails or when a running engine is unable to use a replica for any reason. FailedAt indicates the time the failure occurred. When FailedAt is set, a replica is likely to have useful (though possibly stale) data. A replica with FailedAt set must be rebuilt from a non-failed replica (or it can be used in a salvage if all replicas are failed). FailedAt is cleared before a rebuild or salvage. FailedAt may be later than the corresponding entry in an engine's replicaTransitionTimeMap because it is set when the volume controller acknowledges the change.",
						MarkdownDescription: "FailedAt is set when a running replica fails or when a running engine is unable to use a replica for any reason. FailedAt indicates the time the failure occurred. When FailedAt is set, a replica is likely to have useful (though possibly stale) data. A replica with FailedAt set must be rebuilt from a non-failed replica (or it can be used in a salvage if all replicas are failed). FailedAt is cleared before a rebuild or salvage. FailedAt may be later than the corresponding entry in an engine's replicaTransitionTimeMap because it is set when the volume controller acknowledges the change.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"hard_node_affinity": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"healthy_at": schema.StringAttribute{
						Description:         "HealthyAt is set the first time a replica becomes read/write in an engine after creation or rebuild. HealthyAt indicates the time the last successful rebuild occurred. When HealthyAt is set, a replica is likely to have useful (though possibly stale) data. HealthyAt is cleared before a rebuild. HealthyAt may be later than the corresponding entry in an engine's replicaTransitionTimeMap because it is set when the volume controller acknowledges the change.",
						MarkdownDescription: "HealthyAt is set the first time a replica becomes read/write in an engine after creation or rebuild. HealthyAt indicates the time the last successful rebuild occurred. When HealthyAt is set, a replica is likely to have useful (though possibly stale) data. HealthyAt is cleared before a rebuild. HealthyAt may be later than the corresponding entry in an engine's replicaTransitionTimeMap because it is set when the volume controller acknowledges the change.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"last_failed_at": schema.StringAttribute{
						Description:         "LastFailedAt is always set at the same time as FailedAt. Unlike FailedAt, LastFailedAt is never cleared. LastFailedAt is not a reliable indicator of the state of a replica's data. For example, a replica with LastFailedAt may already be healthy and in use again. However, because it is never cleared, it can be compared to LastHealthyAt to help prevent dangerous replica deletion in some corner cases. LastFailedAt may be later than the corresponding entry in an engine's replicaTransitionTimeMap because it is set when the volume controller acknowledges the change.",
						MarkdownDescription: "LastFailedAt is always set at the same time as FailedAt. Unlike FailedAt, LastFailedAt is never cleared. LastFailedAt is not a reliable indicator of the state of a replica's data. For example, a replica with LastFailedAt may already be healthy and in use again. However, because it is never cleared, it can be compared to LastHealthyAt to help prevent dangerous replica deletion in some corner cases. LastFailedAt may be later than the corresponding entry in an engine's replicaTransitionTimeMap because it is set when the volume controller acknowledges the change.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"last_healthy_at": schema.StringAttribute{
						Description:         "LastHealthyAt is set every time a replica becomes read/write in an engine. Unlike HealthyAt, LastHealthyAt is never cleared. LastHealthyAt is not a reliable indicator of the state of a replica's data. For example, a replica with LastHealthyAt set may be in the middle of a rebuild. However, because it is never cleared, it can be compared to LastFailedAt to help prevent dangerous replica deletion in some corner cases. LastHealthyAt may be later than the corresponding entry in an engine's replicaTransitionTimeMap because it is set when the volume controller acknowledges the change.",
						MarkdownDescription: "LastHealthyAt is set every time a replica becomes read/write in an engine. Unlike HealthyAt, LastHealthyAt is never cleared. LastHealthyAt is not a reliable indicator of the state of a replica's data. For example, a replica with LastHealthyAt set may be in the middle of a rebuild. However, because it is never cleared, it can be compared to LastFailedAt to help prevent dangerous replica deletion in some corner cases. LastHealthyAt may be later than the corresponding entry in an engine's replicaTransitionTimeMap because it is set when the volume controller acknowledges the change.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"log_requested": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"node_id": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"rebuild_retry_count": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"revision_counter_disabled": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"salvage_requested": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"snapshot_max_count": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"snapshot_max_size": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"unmap_mark_disk_chain_removed_enabled": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"volume_name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"volume_size": schema.StringAttribute{
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
	}
}

func (r *LonghornIoReplicaV1Beta2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_longhorn_io_replica_v1beta2_manifest")

	var model LonghornIoReplicaV1Beta2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("longhorn.io/v1beta2")
	model.Kind = pointer.String("Replica")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
