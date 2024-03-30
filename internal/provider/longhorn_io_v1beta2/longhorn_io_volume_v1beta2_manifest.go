/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package longhorn_io_v1beta2

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
	_ datasource.DataSource = &LonghornIoVolumeV1Beta2Manifest{}
)

func NewLonghornIoVolumeV1Beta2Manifest() datasource.DataSource {
	return &LonghornIoVolumeV1Beta2Manifest{}
}

type LonghornIoVolumeV1Beta2Manifest struct{}

type LonghornIoVolumeV1Beta2ManifestData struct {
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
		Standby                     *bool     `tfsdk:"standby" json:"Standby,omitempty"`
		AccessMode                  *string   `tfsdk:"access_mode" json:"accessMode,omitempty"`
		BackendStoreDriver          *string   `tfsdk:"backend_store_driver" json:"backendStoreDriver,omitempty"`
		BackingImage                *string   `tfsdk:"backing_image" json:"backingImage,omitempty"`
		BackupCompressionMethod     *string   `tfsdk:"backup_compression_method" json:"backupCompressionMethod,omitempty"`
		DataEngine                  *string   `tfsdk:"data_engine" json:"dataEngine,omitempty"`
		DataLocality                *string   `tfsdk:"data_locality" json:"dataLocality,omitempty"`
		DataSource                  *string   `tfsdk:"data_source" json:"dataSource,omitempty"`
		DisableFrontend             *bool     `tfsdk:"disable_frontend" json:"disableFrontend,omitempty"`
		DiskSelector                *[]string `tfsdk:"disk_selector" json:"diskSelector,omitempty"`
		Encrypted                   *bool     `tfsdk:"encrypted" json:"encrypted,omitempty"`
		EngineImage                 *string   `tfsdk:"engine_image" json:"engineImage,omitempty"`
		FromBackup                  *string   `tfsdk:"from_backup" json:"fromBackup,omitempty"`
		Frontend                    *string   `tfsdk:"frontend" json:"frontend,omitempty"`
		Image                       *string   `tfsdk:"image" json:"image,omitempty"`
		LastAttachedBy              *string   `tfsdk:"last_attached_by" json:"lastAttachedBy,omitempty"`
		Migratable                  *bool     `tfsdk:"migratable" json:"migratable,omitempty"`
		MigrationNodeID             *string   `tfsdk:"migration_node_id" json:"migrationNodeID,omitempty"`
		NodeID                      *string   `tfsdk:"node_id" json:"nodeID,omitempty"`
		NodeSelector                *[]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
		NumberOfReplicas            *int64    `tfsdk:"number_of_replicas" json:"numberOfReplicas,omitempty"`
		OfflineReplicaRebuilding    *string   `tfsdk:"offline_replica_rebuilding" json:"offlineReplicaRebuilding,omitempty"`
		ReplicaAutoBalance          *string   `tfsdk:"replica_auto_balance" json:"replicaAutoBalance,omitempty"`
		ReplicaDiskSoftAntiAffinity *string   `tfsdk:"replica_disk_soft_anti_affinity" json:"replicaDiskSoftAntiAffinity,omitempty"`
		ReplicaSoftAntiAffinity     *string   `tfsdk:"replica_soft_anti_affinity" json:"replicaSoftAntiAffinity,omitempty"`
		ReplicaZoneSoftAntiAffinity *string   `tfsdk:"replica_zone_soft_anti_affinity" json:"replicaZoneSoftAntiAffinity,omitempty"`
		RestoreVolumeRecurringJob   *string   `tfsdk:"restore_volume_recurring_job" json:"restoreVolumeRecurringJob,omitempty"`
		RevisionCounterDisabled     *bool     `tfsdk:"revision_counter_disabled" json:"revisionCounterDisabled,omitempty"`
		Size                        *string   `tfsdk:"size" json:"size,omitempty"`
		SnapshotDataIntegrity       *string   `tfsdk:"snapshot_data_integrity" json:"snapshotDataIntegrity,omitempty"`
		SnapshotMaxCount            *int64    `tfsdk:"snapshot_max_count" json:"snapshotMaxCount,omitempty"`
		SnapshotMaxSize             *string   `tfsdk:"snapshot_max_size" json:"snapshotMaxSize,omitempty"`
		StaleReplicaTimeout         *int64    `tfsdk:"stale_replica_timeout" json:"staleReplicaTimeout,omitempty"`
		UnmapMarkSnapChainRemoved   *string   `tfsdk:"unmap_mark_snap_chain_removed" json:"unmapMarkSnapChainRemoved,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *LonghornIoVolumeV1Beta2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_longhorn_io_volume_v1beta2_manifest"
}

func (r *LonghornIoVolumeV1Beta2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Volume is where Longhorn stores volume object.",
		MarkdownDescription: "Volume is where Longhorn stores volume object.",
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
				Description:         "VolumeSpec defines the desired state of the Longhorn volume",
				MarkdownDescription: "VolumeSpec defines the desired state of the Longhorn volume",
				Attributes: map[string]schema.Attribute{
					"standby": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"access_mode": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("rwo", "rwx"),
						},
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

					"backup_compression_method": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("none", "lz4", "gzip"),
						},
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

					"data_locality": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("disabled", "best-effort", "strict-local"),
						},
					},

					"data_source": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"disable_frontend": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"disk_selector": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"encrypted": schema.BoolAttribute{
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

					"from_backup": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"frontend": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("blockdev", "iscsi", "nvmf", ""),
						},
					},

					"image": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"last_attached_by": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"migratable": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"migration_node_id": schema.StringAttribute{
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

					"node_selector": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"number_of_replicas": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"offline_replica_rebuilding": schema.StringAttribute{
						Description:         "OfflineReplicaRebuilding is used to determine if the offline replica rebuilding feature is enabled or not",
						MarkdownDescription: "OfflineReplicaRebuilding is used to determine if the offline replica rebuilding feature is enabled or not",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("ignored", "disabled", "enabled"),
						},
					},

					"replica_auto_balance": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("ignored", "disabled", "least-effort", "best-effort"),
						},
					},

					"replica_disk_soft_anti_affinity": schema.StringAttribute{
						Description:         "Replica disk soft anti affinity of the volume. Set enabled to allow replicas to be scheduled in the same disk.",
						MarkdownDescription: "Replica disk soft anti affinity of the volume. Set enabled to allow replicas to be scheduled in the same disk.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("ignored", "enabled", "disabled"),
						},
					},

					"replica_soft_anti_affinity": schema.StringAttribute{
						Description:         "Replica soft anti affinity of the volume. Set enabled to allow replicas to be scheduled on the same node.",
						MarkdownDescription: "Replica soft anti affinity of the volume. Set enabled to allow replicas to be scheduled on the same node.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("ignored", "enabled", "disabled"),
						},
					},

					"replica_zone_soft_anti_affinity": schema.StringAttribute{
						Description:         "Replica zone soft anti affinity of the volume. Set enabled to allow replicas to be scheduled in the same zone.",
						MarkdownDescription: "Replica zone soft anti affinity of the volume. Set enabled to allow replicas to be scheduled in the same zone.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("ignored", "enabled", "disabled"),
						},
					},

					"restore_volume_recurring_job": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("ignored", "enabled", "disabled"),
						},
					},

					"revision_counter_disabled": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"size": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"snapshot_data_integrity": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("ignored", "disabled", "enabled", "fast-check"),
						},
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

					"stale_replica_timeout": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"unmap_mark_snap_chain_removed": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("ignored", "disabled", "enabled"),
						},
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *LonghornIoVolumeV1Beta2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_longhorn_io_volume_v1beta2_manifest")

	var model LonghornIoVolumeV1Beta2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("longhorn.io/v1beta2")
	model.Kind = pointer.String("Volume")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
