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
	_ datasource.DataSource = &VeleroIoDataDownloadV2Alpha1Manifest{}
)

func NewVeleroIoDataDownloadV2Alpha1Manifest() datasource.DataSource {
	return &VeleroIoDataDownloadV2Alpha1Manifest{}
}

type VeleroIoDataDownloadV2Alpha1Manifest struct{}

type VeleroIoDataDownloadV2Alpha1ManifestData struct {
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
		BackupStorageLocation *string            `tfsdk:"backup_storage_location" json:"backupStorageLocation,omitempty"`
		Cancel                *bool              `tfsdk:"cancel" json:"cancel,omitempty"`
		DataMoverConfig       *map[string]string `tfsdk:"data_mover_config" json:"dataMoverConfig,omitempty"`
		Datamover             *string            `tfsdk:"datamover" json:"datamover,omitempty"`
		OperationTimeout      *string            `tfsdk:"operation_timeout" json:"operationTimeout,omitempty"`
		SnapshotID            *string            `tfsdk:"snapshot_id" json:"snapshotID,omitempty"`
		SourceNamespace       *string            `tfsdk:"source_namespace" json:"sourceNamespace,omitempty"`
		TargetVolume          *struct {
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Pv        *string `tfsdk:"pv" json:"pv,omitempty"`
			Pvc       *string `tfsdk:"pvc" json:"pvc,omitempty"`
		} `tfsdk:"target_volume" json:"targetVolume,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *VeleroIoDataDownloadV2Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_velero_io_data_download_v2alpha1_manifest"
}

func (r *VeleroIoDataDownloadV2Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "DataDownload acts as the protocol between data mover plugins and data mover controller for the datamover restore operation",
		MarkdownDescription: "DataDownload acts as the protocol between data mover plugins and data mover controller for the datamover restore operation",
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
				Description:         "DataDownloadSpec is the specification for a DataDownload.",
				MarkdownDescription: "DataDownloadSpec is the specification for a DataDownload.",
				Attributes: map[string]schema.Attribute{
					"backup_storage_location": schema.StringAttribute{
						Description:         "BackupStorageLocation is the name of the backup storage location where the backup repository is stored.",
						MarkdownDescription: "BackupStorageLocation is the name of the backup storage location where the backup repository is stored.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"cancel": schema.BoolAttribute{
						Description:         "Cancel indicates request to cancel the ongoing DataDownload. It can be set when the DataDownload is in InProgress phase",
						MarkdownDescription: "Cancel indicates request to cancel the ongoing DataDownload. It can be set when the DataDownload is in InProgress phase",
						Required:            false,
						Optional:            true,
						Computed:            false,
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

					"snapshot_id": schema.StringAttribute{
						Description:         "SnapshotID is the ID of the Velero backup snapshot to be restored from.",
						MarkdownDescription: "SnapshotID is the ID of the Velero backup snapshot to be restored from.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"source_namespace": schema.StringAttribute{
						Description:         "SourceNamespace is the original namespace where the volume is backed up from. It may be different from SourcePVC's namespace if namespace is remapped during restore.",
						MarkdownDescription: "SourceNamespace is the original namespace where the volume is backed up from. It may be different from SourcePVC's namespace if namespace is remapped during restore.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"target_volume": schema.SingleNestedAttribute{
						Description:         "TargetVolume is the information of the target PVC and PV.",
						MarkdownDescription: "TargetVolume is the information of the target PVC and PV.",
						Attributes: map[string]schema.Attribute{
							"namespace": schema.StringAttribute{
								Description:         "Namespace is the target namespace",
								MarkdownDescription: "Namespace is the target namespace",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"pv": schema.StringAttribute{
								Description:         "PV is the name of the target PV that is created by Velero restore",
								MarkdownDescription: "PV is the name of the target PV that is created by Velero restore",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"pvc": schema.StringAttribute{
								Description:         "PVC is the name of the target PVC that is created by Velero restore",
								MarkdownDescription: "PVC is the name of the target PVC that is created by Velero restore",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
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

func (r *VeleroIoDataDownloadV2Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_velero_io_data_download_v2alpha1_manifest")

	var model VeleroIoDataDownloadV2Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("velero.io/v2alpha1")
	model.Kind = pointer.String("DataDownload")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
