/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package snapshot_storage_k8s_io_v1

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
	_ datasource.DataSource = &SnapshotStorageK8SIoVolumeSnapshotClassV1Manifest{}
)

func NewSnapshotStorageK8SIoVolumeSnapshotClassV1Manifest() datasource.DataSource {
	return &SnapshotStorageK8SIoVolumeSnapshotClassV1Manifest{}
}

type SnapshotStorageK8SIoVolumeSnapshotClassV1Manifest struct{}

type SnapshotStorageK8SIoVolumeSnapshotClassV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	DeletionPolicy *string            `tfsdk:"deletion_policy" json:"deletionPolicy,omitempty"`
	Driver         *string            `tfsdk:"driver" json:"driver,omitempty"`
	Parameters     *map[string]string `tfsdk:"parameters" json:"parameters,omitempty"`
}

func (r *SnapshotStorageK8SIoVolumeSnapshotClassV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_snapshot_storage_k8s_io_volume_snapshot_class_v1_manifest"
}

func (r *SnapshotStorageK8SIoVolumeSnapshotClassV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "VolumeSnapshotClass specifies parameters that a underlying storage system uses when creating a volume snapshot. A specific VolumeSnapshotClass is used by specifying its name in a VolumeSnapshot object. VolumeSnapshotClasses are non-namespaced",
		MarkdownDescription: "VolumeSnapshotClass specifies parameters that a underlying storage system uses when creating a volume snapshot. A specific VolumeSnapshotClass is used by specifying its name in a VolumeSnapshot object. VolumeSnapshotClasses are non-namespaced",
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

			"deletion_policy": schema.StringAttribute{
				Description:         "deletionPolicy determines whether a VolumeSnapshotContent created through the VolumeSnapshotClass should be deleted when its bound VolumeSnapshot is deleted. Supported values are 'Retain' and 'Delete'. 'Retain' means that the VolumeSnapshotContent and its physical snapshot on underlying storage system are kept. 'Delete' means that the VolumeSnapshotContent and its physical snapshot on underlying storage system are deleted. Required.",
				MarkdownDescription: "deletionPolicy determines whether a VolumeSnapshotContent created through the VolumeSnapshotClass should be deleted when its bound VolumeSnapshot is deleted. Supported values are 'Retain' and 'Delete'. 'Retain' means that the VolumeSnapshotContent and its physical snapshot on underlying storage system are kept. 'Delete' means that the VolumeSnapshotContent and its physical snapshot on underlying storage system are deleted. Required.",
				Required:            true,
				Optional:            false,
				Computed:            false,
				Validators: []validator.String{
					stringvalidator.OneOf("Delete", "Retain"),
				},
			},

			"driver": schema.StringAttribute{
				Description:         "driver is the name of the storage driver that handles this VolumeSnapshotClass. Required.",
				MarkdownDescription: "driver is the name of the storage driver that handles this VolumeSnapshotClass. Required.",
				Required:            true,
				Optional:            false,
				Computed:            false,
			},

			"parameters": schema.MapAttribute{
				Description:         "parameters is a key-value map with storage driver specific parameters for creating snapshots. These values are opaque to Kubernetes.",
				MarkdownDescription: "parameters is a key-value map with storage driver specific parameters for creating snapshots. These values are opaque to Kubernetes.",
				ElementType:         types.StringType,
				Required:            false,
				Optional:            true,
				Computed:            false,
			},
		},
	}
}

func (r *SnapshotStorageK8SIoVolumeSnapshotClassV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_snapshot_storage_k8s_io_volume_snapshot_class_v1_manifest")

	var model SnapshotStorageK8SIoVolumeSnapshotClassV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("snapshot.storage.k8s.io/v1")
	model.Kind = pointer.String("VolumeSnapshotClass")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
