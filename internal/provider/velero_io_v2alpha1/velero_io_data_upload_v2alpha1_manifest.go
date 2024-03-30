/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package velero_io_v2alpha1

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
	_ datasource.DataSource = &VeleroIoDataUploadV2Alpha1Manifest{}
)

func NewVeleroIoDataUploadV2Alpha1Manifest() datasource.DataSource {
	return &VeleroIoDataUploadV2Alpha1Manifest{}
}

type VeleroIoDataUploadV2Alpha1Manifest struct{}

type VeleroIoDataUploadV2Alpha1ManifestData struct {
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
		BackupStorageLocation *string `tfsdk:"backup_storage_location" json:"backupStorageLocation,omitempty"`
		Cancel                *bool   `tfsdk:"cancel" json:"cancel,omitempty"`
		CsiSnapshot           *struct {
			SnapshotClass  *string `tfsdk:"snapshot_class" json:"snapshotClass,omitempty"`
			StorageClass   *string `tfsdk:"storage_class" json:"storageClass,omitempty"`
			VolumeSnapshot *string `tfsdk:"volume_snapshot" json:"volumeSnapshot,omitempty"`
		} `tfsdk:"csi_snapshot" json:"csiSnapshot,omitempty"`
		DataMoverConfig  *map[string]string `tfsdk:"data_mover_config" json:"dataMoverConfig,omitempty"`
		Datamover        *string            `tfsdk:"datamover" json:"datamover,omitempty"`
		OperationTimeout *string            `tfsdk:"operation_timeout" json:"operationTimeout,omitempty"`
		SnapshotType     *string            `tfsdk:"snapshot_type" json:"snapshotType,omitempty"`
		SourceNamespace  *string            `tfsdk:"source_namespace" json:"sourceNamespace,omitempty"`
		SourcePVC        *string            `tfsdk:"source_pvc" json:"sourcePVC,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *VeleroIoDataUploadV2Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_velero_io_data_upload_v2alpha1_manifest"
}

func (r *VeleroIoDataUploadV2Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "DataUpload acts as the protocol between data mover plugins and data mover controller for the datamover backup operation",
		MarkdownDescription: "DataUpload acts as the protocol between data mover plugins and data mover controller for the datamover backup operation",
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
				Description:         "DataUploadSpec is the specification for a DataUpload.",
				MarkdownDescription: "DataUploadSpec is the specification for a DataUpload.",
				Attributes: map[string]schema.Attribute{
					"backup_storage_location": schema.StringAttribute{
						Description:         "BackupStorageLocation is the name of the backup storage location where the backup repository is stored.",
						MarkdownDescription: "BackupStorageLocation is the name of the backup storage location where the backup repository is stored.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"cancel": schema.BoolAttribute{
						Description:         "Cancel indicates request to cancel the ongoing DataUpload. It can be set when the DataUpload is in InProgress phase",
						MarkdownDescription: "Cancel indicates request to cancel the ongoing DataUpload. It can be set when the DataUpload is in InProgress phase",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"csi_snapshot": schema.SingleNestedAttribute{
						Description:         "If SnapshotType is CSI, CSISnapshot provides the information of the CSI snapshot.",
						MarkdownDescription: "If SnapshotType is CSI, CSISnapshot provides the information of the CSI snapshot.",
						Attributes: map[string]schema.Attribute{
							"snapshot_class": schema.StringAttribute{
								Description:         "SnapshotClass is the name of the snapshot class that the volume snapshot is created with",
								MarkdownDescription: "SnapshotClass is the name of the snapshot class that the volume snapshot is created with",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"storage_class": schema.StringAttribute{
								Description:         "StorageClass is the name of the storage class of the PVC that the volume snapshot is created from",
								MarkdownDescription: "StorageClass is the name of the storage class of the PVC that the volume snapshot is created from",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"volume_snapshot": schema.StringAttribute{
								Description:         "VolumeSnapshot is the name of the volume snapshot to be backed up",
								MarkdownDescription: "VolumeSnapshot is the name of the volume snapshot to be backed up",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"data_mover_config": schema.MapAttribute{
						Description:         "DataMoverConfig is for data-mover-specific configuration fields.",
						MarkdownDescription: "DataMoverConfig is for data-mover-specific configuration fields.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"datamover": schema.StringAttribute{
						Description:         "DataMover specifies the data mover to be used by the backup. If DataMover is '' or 'velero', the built-in data mover will be used.",
						MarkdownDescription: "DataMover specifies the data mover to be used by the backup. If DataMover is '' or 'velero', the built-in data mover will be used.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"operation_timeout": schema.StringAttribute{
						Description:         "OperationTimeout specifies the time used to wait internal operations, before returning error as timeout.",
						MarkdownDescription: "OperationTimeout specifies the time used to wait internal operations, before returning error as timeout.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"snapshot_type": schema.StringAttribute{
						Description:         "SnapshotType is the type of the snapshot to be backed up.",
						MarkdownDescription: "SnapshotType is the type of the snapshot to be backed up.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"source_namespace": schema.StringAttribute{
						Description:         "SourceNamespace is the original namespace where the volume is backed up from. It is the same namespace for SourcePVC and CSI namespaced objects.",
						MarkdownDescription: "SourceNamespace is the original namespace where the volume is backed up from. It is the same namespace for SourcePVC and CSI namespaced objects.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"source_pvc": schema.StringAttribute{
						Description:         "SourcePVC is the name of the PVC which the snapshot is taken for.",
						MarkdownDescription: "SourcePVC is the name of the PVC which the snapshot is taken for.",
						Required:            true,
						Optional:            false,
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

func (r *VeleroIoDataUploadV2Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_velero_io_data_upload_v2alpha1_manifest")

	var model VeleroIoDataUploadV2Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("velero.io/v2alpha1")
	model.Kind = pointer.String("DataUpload")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
