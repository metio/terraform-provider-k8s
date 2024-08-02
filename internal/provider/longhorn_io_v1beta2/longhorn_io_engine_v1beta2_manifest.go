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
	_ datasource.DataSource = &LonghornIoEngineV1Beta2Manifest{}
)

func NewLonghornIoEngineV1Beta2Manifest() datasource.DataSource {
	return &LonghornIoEngineV1Beta2Manifest{}
}

type LonghornIoEngineV1Beta2Manifest struct{}

type LonghornIoEngineV1Beta2ManifestData struct {
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
		Active                           *bool              `tfsdk:"active" json:"active,omitempty"`
		BackendStoreDriver               *string            `tfsdk:"backend_store_driver" json:"backendStoreDriver,omitempty"`
		BackupVolume                     *string            `tfsdk:"backup_volume" json:"backupVolume,omitempty"`
		DataEngine                       *string            `tfsdk:"data_engine" json:"dataEngine,omitempty"`
		DesireState                      *string            `tfsdk:"desire_state" json:"desireState,omitempty"`
		DisableFrontend                  *bool              `tfsdk:"disable_frontend" json:"disableFrontend,omitempty"`
		EngineImage                      *string            `tfsdk:"engine_image" json:"engineImage,omitempty"`
		Frontend                         *string            `tfsdk:"frontend" json:"frontend,omitempty"`
		Image                            *string            `tfsdk:"image" json:"image,omitempty"`
		LogRequested                     *bool              `tfsdk:"log_requested" json:"logRequested,omitempty"`
		NodeID                           *string            `tfsdk:"node_id" json:"nodeID,omitempty"`
		ReplicaAddressMap                *map[string]string `tfsdk:"replica_address_map" json:"replicaAddressMap,omitempty"`
		RequestedBackupRestore           *string            `tfsdk:"requested_backup_restore" json:"requestedBackupRestore,omitempty"`
		RequestedDataSource              *string            `tfsdk:"requested_data_source" json:"requestedDataSource,omitempty"`
		RevisionCounterDisabled          *bool              `tfsdk:"revision_counter_disabled" json:"revisionCounterDisabled,omitempty"`
		SalvageRequested                 *bool              `tfsdk:"salvage_requested" json:"salvageRequested,omitempty"`
		SnapshotMaxCount                 *int64             `tfsdk:"snapshot_max_count" json:"snapshotMaxCount,omitempty"`
		SnapshotMaxSize                  *string            `tfsdk:"snapshot_max_size" json:"snapshotMaxSize,omitempty"`
		UnmapMarkSnapChainRemovedEnabled *bool              `tfsdk:"unmap_mark_snap_chain_removed_enabled" json:"unmapMarkSnapChainRemovedEnabled,omitempty"`
		UpgradedReplicaAddressMap        *map[string]string `tfsdk:"upgraded_replica_address_map" json:"upgradedReplicaAddressMap,omitempty"`
		VolumeName                       *string            `tfsdk:"volume_name" json:"volumeName,omitempty"`
		VolumeSize                       *string            `tfsdk:"volume_size" json:"volumeSize,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *LonghornIoEngineV1Beta2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_longhorn_io_engine_v1beta2_manifest"
}

func (r *LonghornIoEngineV1Beta2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Engine is where Longhorn stores engine object.",
		MarkdownDescription: "Engine is where Longhorn stores engine object.",
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
				Description:         "EngineSpec defines the desired state of the Longhorn engine",
				MarkdownDescription: "EngineSpec defines the desired state of the Longhorn engine",
				Attributes: map[string]schema.Attribute{
					"active": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"backend_store_driver": schema.StringAttribute{
						Description:         "Deprecated:Replaced by field 'dataEngine'.",
						MarkdownDescription: "Deprecated:Replaced by field 'dataEngine'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"backup_volume": schema.StringAttribute{
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

					"disable_frontend": schema.BoolAttribute{
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

					"replica_address_map": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"requested_backup_restore": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"requested_data_source": schema.StringAttribute{
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

					"unmap_mark_snap_chain_removed_enabled": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"upgraded_replica_address_map": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
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

func (r *LonghornIoEngineV1Beta2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_longhorn_io_engine_v1beta2_manifest")

	var model LonghornIoEngineV1Beta2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("longhorn.io/v1beta2")
	model.Kind = pointer.String("Engine")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
