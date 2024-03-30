/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package ceph_rook_io_v1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
	_ datasource.DataSource = &CephRookIoCephFilesystemSubVolumeGroupV1Manifest{}
)

func NewCephRookIoCephFilesystemSubVolumeGroupV1Manifest() datasource.DataSource {
	return &CephRookIoCephFilesystemSubVolumeGroupV1Manifest{}
}

type CephRookIoCephFilesystemSubVolumeGroupV1Manifest struct{}

type CephRookIoCephFilesystemSubVolumeGroupV1ManifestData struct {
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
		FilesystemName *string `tfsdk:"filesystem_name" json:"filesystemName,omitempty"`
		Name           *string `tfsdk:"name" json:"name,omitempty"`
		Pinning        *struct {
			Distributed *int64   `tfsdk:"distributed" json:"distributed,omitempty"`
			Export      *int64   `tfsdk:"export" json:"export,omitempty"`
			Random      *float64 `tfsdk:"random" json:"random,omitempty"`
		} `tfsdk:"pinning" json:"pinning,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CephRookIoCephFilesystemSubVolumeGroupV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_ceph_rook_io_ceph_filesystem_sub_volume_group_v1_manifest"
}

func (r *CephRookIoCephFilesystemSubVolumeGroupV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CephFilesystemSubVolumeGroup represents a Ceph Filesystem SubVolumeGroup",
		MarkdownDescription: "CephFilesystemSubVolumeGroup represents a Ceph Filesystem SubVolumeGroup",
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
				Description:         "Spec represents the specification of a Ceph Filesystem SubVolumeGroup",
				MarkdownDescription: "Spec represents the specification of a Ceph Filesystem SubVolumeGroup",
				Attributes: map[string]schema.Attribute{
					"filesystem_name": schema.StringAttribute{
						Description:         "FilesystemName is the name of Ceph Filesystem SubVolumeGroup volume name. Typically it's the name of the CephFilesystem CR. If not coming from the CephFilesystem CR, it can be retrieved from the list of Ceph Filesystem volumes with 'ceph fs volume ls'. To learn more about Ceph Filesystem abstractions see https://docs.ceph.com/en/latest/cephfs/fs-volumes/#fs-volumes-and-subvolumes",
						MarkdownDescription: "FilesystemName is the name of Ceph Filesystem SubVolumeGroup volume name. Typically it's the name of the CephFilesystem CR. If not coming from the CephFilesystem CR, it can be retrieved from the list of Ceph Filesystem volumes with 'ceph fs volume ls'. To learn more about Ceph Filesystem abstractions see https://docs.ceph.com/en/latest/cephfs/fs-volumes/#fs-volumes-and-subvolumes",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"name": schema.StringAttribute{
						Description:         "The name of the subvolume group. If not set, the default is the name of the subvolumeGroup CR.",
						MarkdownDescription: "The name of the subvolume group. If not set, the default is the name of the subvolumeGroup CR.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"pinning": schema.SingleNestedAttribute{
						Description:         "Pinning configuration of CephFilesystemSubVolumeGroup, reference https://docs.ceph.com/en/latest/cephfs/fs-volumes/#pinning-subvolumes-and-subvolume-groups only one out of (export, distributed, random) can be set at a time",
						MarkdownDescription: "Pinning configuration of CephFilesystemSubVolumeGroup, reference https://docs.ceph.com/en/latest/cephfs/fs-volumes/#pinning-subvolumes-and-subvolume-groups only one out of (export, distributed, random) can be set at a time",
						Attributes: map[string]schema.Attribute{
							"distributed": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
									int64validator.AtMost(1),
								},
							},

							"export": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(-1),
									int64validator.AtMost(256),
								},
							},

							"random": schema.Float64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Float64{
									float64validator.AtLeast(0),
									float64validator.AtMost(1),
								},
							},
						},
						Required: false,
						Optional: true,
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

func (r *CephRookIoCephFilesystemSubVolumeGroupV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_ceph_rook_io_ceph_filesystem_sub_volume_group_v1_manifest")

	var model CephRookIoCephFilesystemSubVolumeGroupV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("ceph.rook.io/v1")
	model.Kind = pointer.String("CephFilesystemSubVolumeGroup")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
