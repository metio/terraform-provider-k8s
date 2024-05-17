/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package groupsnapshot_storage_k8s_io_v1alpha1

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
	_ datasource.DataSource = &GroupsnapshotStorageK8SIoVolumeGroupSnapshotContentV1Alpha1Manifest{}
)

func NewGroupsnapshotStorageK8SIoVolumeGroupSnapshotContentV1Alpha1Manifest() datasource.DataSource {
	return &GroupsnapshotStorageK8SIoVolumeGroupSnapshotContentV1Alpha1Manifest{}
}

type GroupsnapshotStorageK8SIoVolumeGroupSnapshotContentV1Alpha1Manifest struct{}

type GroupsnapshotStorageK8SIoVolumeGroupSnapshotContentV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		DeletionPolicy *string `tfsdk:"deletion_policy" json:"deletionPolicy,omitempty"`
		Driver         *string `tfsdk:"driver" json:"driver,omitempty"`
		Source         *struct {
			GroupSnapshotHandles *struct {
				VolumeGroupSnapshotHandle *string   `tfsdk:"volume_group_snapshot_handle" json:"volumeGroupSnapshotHandle,omitempty"`
				VolumeSnapshotHandles     *[]string `tfsdk:"volume_snapshot_handles" json:"volumeSnapshotHandles,omitempty"`
			} `tfsdk:"group_snapshot_handles" json:"groupSnapshotHandles,omitempty"`
			VolumeHandles *[]string `tfsdk:"volume_handles" json:"volumeHandles,omitempty"`
		} `tfsdk:"source" json:"source,omitempty"`
		VolumeGroupSnapshotClassName *string `tfsdk:"volume_group_snapshot_class_name" json:"volumeGroupSnapshotClassName,omitempty"`
		VolumeGroupSnapshotRef       *struct {
			ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
			Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
			Name            *string `tfsdk:"name" json:"name,omitempty"`
			Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
			ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
			Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
		} `tfsdk:"volume_group_snapshot_ref" json:"volumeGroupSnapshotRef,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *GroupsnapshotStorageK8SIoVolumeGroupSnapshotContentV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_groupsnapshot_storage_k8s_io_volume_group_snapshot_content_v1alpha1_manifest"
}

func (r *GroupsnapshotStorageK8SIoVolumeGroupSnapshotContentV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "VolumeGroupSnapshotContent represents the actual 'on-disk' group snapshot objectin the underlying storage system",
		MarkdownDescription: "VolumeGroupSnapshotContent represents the actual 'on-disk' group snapshot objectin the underlying storage system",
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
				Description:         "Spec defines properties of a VolumeGroupSnapshotContent created by the underlying storage system.Required.",
				MarkdownDescription: "Spec defines properties of a VolumeGroupSnapshotContent created by the underlying storage system.Required.",
				Attributes: map[string]schema.Attribute{
					"deletion_policy": schema.StringAttribute{
						Description:         "DeletionPolicy determines whether this VolumeGroupSnapshotContent and thephysical group snapshot on the underlying storage system should be deletedwhen the bound VolumeGroupSnapshot is deleted.Supported values are 'Retain' and 'Delete'.'Retain' means that the VolumeGroupSnapshotContent and its physical groupsnapshot on underlying storage system are kept.'Delete' means that the VolumeGroupSnapshotContent and its physical groupsnapshot on underlying storage system are deleted.For dynamically provisioned group snapshots, this field will automaticallybe filled in by the CSI snapshotter sidecar with the 'DeletionPolicy' fielddefined in the corresponding VolumeGroupSnapshotClass.For pre-existing snapshots, users MUST specify this field when creating theVolumeGroupSnapshotContent object.Required.",
						MarkdownDescription: "DeletionPolicy determines whether this VolumeGroupSnapshotContent and thephysical group snapshot on the underlying storage system should be deletedwhen the bound VolumeGroupSnapshot is deleted.Supported values are 'Retain' and 'Delete'.'Retain' means that the VolumeGroupSnapshotContent and its physical groupsnapshot on underlying storage system are kept.'Delete' means that the VolumeGroupSnapshotContent and its physical groupsnapshot on underlying storage system are deleted.For dynamically provisioned group snapshots, this field will automaticallybe filled in by the CSI snapshotter sidecar with the 'DeletionPolicy' fielddefined in the corresponding VolumeGroupSnapshotClass.For pre-existing snapshots, users MUST specify this field when creating theVolumeGroupSnapshotContent object.Required.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Delete", "Retain"),
						},
					},

					"driver": schema.StringAttribute{
						Description:         "Driver is the name of the CSI driver used to create the physical group snapshot onthe underlying storage system.This MUST be the same as the name returned by the CSI GetPluginName() call forthat driver.Required.",
						MarkdownDescription: "Driver is the name of the CSI driver used to create the physical group snapshot onthe underlying storage system.This MUST be the same as the name returned by the CSI GetPluginName() call forthat driver.Required.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"source": schema.SingleNestedAttribute{
						Description:         "Source specifies whether the snapshot is (or should be) dynamically provisionedor already exists, and just requires a Kubernetes object representation.This field is immutable after creation.Required.",
						MarkdownDescription: "Source specifies whether the snapshot is (or should be) dynamically provisionedor already exists, and just requires a Kubernetes object representation.This field is immutable after creation.Required.",
						Attributes: map[string]schema.Attribute{
							"group_snapshot_handles": schema.SingleNestedAttribute{
								Description:         "GroupSnapshotHandles specifies the CSI 'group_snapshot_id' of a pre-existinggroup snapshot and a list of CSI 'snapshot_id' of pre-existing snapshotson the underlying storage system for which a Kubernetes objectrepresentation was (or should be) created.This field is immutable.",
								MarkdownDescription: "GroupSnapshotHandles specifies the CSI 'group_snapshot_id' of a pre-existinggroup snapshot and a list of CSI 'snapshot_id' of pre-existing snapshotson the underlying storage system for which a Kubernetes objectrepresentation was (or should be) created.This field is immutable.",
								Attributes: map[string]schema.Attribute{
									"volume_group_snapshot_handle": schema.StringAttribute{
										Description:         "VolumeGroupSnapshotHandle specifies the CSI 'group_snapshot_id' of a pre-existinggroup snapshot on the underlying storage system for which a Kubernetes objectrepresentation was (or should be) created.This field is immutable.Required.",
										MarkdownDescription: "VolumeGroupSnapshotHandle specifies the CSI 'group_snapshot_id' of a pre-existinggroup snapshot on the underlying storage system for which a Kubernetes objectrepresentation was (or should be) created.This field is immutable.Required.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"volume_snapshot_handles": schema.ListAttribute{
										Description:         "VolumeSnapshotHandles is a list of CSI 'snapshot_id' of pre-existingsnapshots on the underlying storage system for which Kubernetes objectsrepresentation were (or should be) created.This field is immutable.Required.",
										MarkdownDescription: "VolumeSnapshotHandles is a list of CSI 'snapshot_id' of pre-existingsnapshots on the underlying storage system for which Kubernetes objectsrepresentation were (or should be) created.This field is immutable.Required.",
										ElementType:         types.StringType,
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"volume_handles": schema.ListAttribute{
								Description:         "VolumeHandles is a list of volume handles on the backend to be snapshottedtogether. It is specified for dynamic provisioning of the VolumeGroupSnapshot.This field is immutable.",
								MarkdownDescription: "VolumeHandles is a list of volume handles on the backend to be snapshottedtogether. It is specified for dynamic provisioning of the VolumeGroupSnapshot.This field is immutable.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"volume_group_snapshot_class_name": schema.StringAttribute{
						Description:         "VolumeGroupSnapshotClassName is the name of the VolumeGroupSnapshotClass fromwhich this group snapshot was (or will be) created.Note that after provisioning, the VolumeGroupSnapshotClass may be deleted orrecreated with different set of values, and as such, should not be referencedpost-snapshot creation.For dynamic provisioning, this field must be set.This field may be unset for pre-provisioned snapshots.",
						MarkdownDescription: "VolumeGroupSnapshotClassName is the name of the VolumeGroupSnapshotClass fromwhich this group snapshot was (or will be) created.Note that after provisioning, the VolumeGroupSnapshotClass may be deleted orrecreated with different set of values, and as such, should not be referencedpost-snapshot creation.For dynamic provisioning, this field must be set.This field may be unset for pre-provisioned snapshots.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"volume_group_snapshot_ref": schema.SingleNestedAttribute{
						Description:         "VolumeGroupSnapshotRef specifies the VolumeGroupSnapshot object to which thisVolumeGroupSnapshotContent object is bound.VolumeGroupSnapshot.Spec.VolumeGroupSnapshotContentName field must reference tothis VolumeGroupSnapshotContent's name for the bidirectional binding to be valid.For a pre-existing VolumeGroupSnapshotContent object, name and namespace of theVolumeGroupSnapshot object MUST be provided for binding to happen.This field is immutable after creation.Required.",
						MarkdownDescription: "VolumeGroupSnapshotRef specifies the VolumeGroupSnapshot object to which thisVolumeGroupSnapshotContent object is bound.VolumeGroupSnapshot.Spec.VolumeGroupSnapshotContentName field must reference tothis VolumeGroupSnapshotContent's name for the bidirectional binding to be valid.For a pre-existing VolumeGroupSnapshotContent object, name and namespace of theVolumeGroupSnapshot object MUST be provided for binding to happen.This field is immutable after creation.Required.",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "API version of the referent.",
								MarkdownDescription: "API version of the referent.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"field_path": schema.StringAttribute{
								Description:         "If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.",
								MarkdownDescription: "If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								MarkdownDescription: "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
								MarkdownDescription: "Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resource_version": schema.StringAttribute{
								Description:         "Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
								MarkdownDescription: "Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"uid": schema.StringAttribute{
								Description:         "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
								MarkdownDescription: "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *GroupsnapshotStorageK8SIoVolumeGroupSnapshotContentV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_groupsnapshot_storage_k8s_io_volume_group_snapshot_content_v1alpha1_manifest")

	var model GroupsnapshotStorageK8SIoVolumeGroupSnapshotContentV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("groupsnapshot.storage.k8s.io/v1alpha1")
	model.Kind = pointer.String("VolumeGroupSnapshotContent")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
