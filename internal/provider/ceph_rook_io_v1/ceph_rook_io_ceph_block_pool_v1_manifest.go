/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package ceph_rook_io_v1

import (
	"context"
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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &CephRookIoCephBlockPoolV1Manifest{}
)

func NewCephRookIoCephBlockPoolV1Manifest() datasource.DataSource {
	return &CephRookIoCephBlockPoolV1Manifest{}
}

type CephRookIoCephBlockPoolV1Manifest struct{}

type CephRookIoCephBlockPoolV1ManifestData struct {
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
		Application     *string `tfsdk:"application" json:"application,omitempty"`
		CompressionMode *string `tfsdk:"compression_mode" json:"compressionMode,omitempty"`
		CrushRoot       *string `tfsdk:"crush_root" json:"crushRoot,omitempty"`
		DeviceClass     *string `tfsdk:"device_class" json:"deviceClass,omitempty"`
		EnableRBDStats  *bool   `tfsdk:"enable_rbd_stats" json:"enableRBDStats,omitempty"`
		ErasureCoded    *struct {
			Algorithm    *string `tfsdk:"algorithm" json:"algorithm,omitempty"`
			CodingChunks *int64  `tfsdk:"coding_chunks" json:"codingChunks,omitempty"`
			DataChunks   *int64  `tfsdk:"data_chunks" json:"dataChunks,omitempty"`
		} `tfsdk:"erasure_coded" json:"erasureCoded,omitempty"`
		FailureDomain *string `tfsdk:"failure_domain" json:"failureDomain,omitempty"`
		Mirroring     *struct {
			Enabled *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
			Mode    *string `tfsdk:"mode" json:"mode,omitempty"`
			Peers   *struct {
				SecretNames *[]string `tfsdk:"secret_names" json:"secretNames,omitempty"`
			} `tfsdk:"peers" json:"peers,omitempty"`
			SnapshotSchedules *[]struct {
				Interval  *string `tfsdk:"interval" json:"interval,omitempty"`
				Path      *string `tfsdk:"path" json:"path,omitempty"`
				StartTime *string `tfsdk:"start_time" json:"startTime,omitempty"`
			} `tfsdk:"snapshot_schedules" json:"snapshotSchedules,omitempty"`
		} `tfsdk:"mirroring" json:"mirroring,omitempty"`
		Name       *string            `tfsdk:"name" json:"name,omitempty"`
		Parameters *map[string]string `tfsdk:"parameters" json:"parameters,omitempty"`
		Quotas     *struct {
			MaxBytes   *int64  `tfsdk:"max_bytes" json:"maxBytes,omitempty"`
			MaxObjects *int64  `tfsdk:"max_objects" json:"maxObjects,omitempty"`
			MaxSize    *string `tfsdk:"max_size" json:"maxSize,omitempty"`
		} `tfsdk:"quotas" json:"quotas,omitempty"`
		Replicated *struct {
			HybridStorage *struct {
				PrimaryDeviceClass   *string `tfsdk:"primary_device_class" json:"primaryDeviceClass,omitempty"`
				SecondaryDeviceClass *string `tfsdk:"secondary_device_class" json:"secondaryDeviceClass,omitempty"`
			} `tfsdk:"hybrid_storage" json:"hybridStorage,omitempty"`
			ReplicasPerFailureDomain *int64   `tfsdk:"replicas_per_failure_domain" json:"replicasPerFailureDomain,omitempty"`
			RequireSafeReplicaSize   *bool    `tfsdk:"require_safe_replica_size" json:"requireSafeReplicaSize,omitempty"`
			Size                     *int64   `tfsdk:"size" json:"size,omitempty"`
			SubFailureDomain         *string  `tfsdk:"sub_failure_domain" json:"subFailureDomain,omitempty"`
			TargetSizeRatio          *float64 `tfsdk:"target_size_ratio" json:"targetSizeRatio,omitempty"`
		} `tfsdk:"replicated" json:"replicated,omitempty"`
		StatusCheck *struct {
			Mirror *struct {
				Disabled *bool   `tfsdk:"disabled" json:"disabled,omitempty"`
				Interval *string `tfsdk:"interval" json:"interval,omitempty"`
				Timeout  *string `tfsdk:"timeout" json:"timeout,omitempty"`
			} `tfsdk:"mirror" json:"mirror,omitempty"`
		} `tfsdk:"status_check" json:"statusCheck,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CephRookIoCephBlockPoolV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_ceph_rook_io_ceph_block_pool_v1_manifest"
}

func (r *CephRookIoCephBlockPoolV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CephBlockPool represents a Ceph Storage Pool",
		MarkdownDescription: "CephBlockPool represents a Ceph Storage Pool",
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
				Description:         "NamedBlockPoolSpec allows a block pool to be created with a non-default name.This is more specific than the NamedPoolSpec so we get schema validation on theallowed pool names that can be specified.",
				MarkdownDescription: "NamedBlockPoolSpec allows a block pool to be created with a non-default name.This is more specific than the NamedPoolSpec so we get schema validation on theallowed pool names that can be specified.",
				Attributes: map[string]schema.Attribute{
					"application": schema.StringAttribute{
						Description:         "The application name to set on the pool. Only expected to be set for rgw pools.",
						MarkdownDescription: "The application name to set on the pool. Only expected to be set for rgw pools.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"compression_mode": schema.StringAttribute{
						Description:         "DEPRECATED: use Parameters instead, e.g., Parameters['compression_mode'] = 'force'The inline compression mode in Bluestore OSD to set to (options are: none, passive, aggressive, force)Do NOT set a default value for kubebuilder as this will override the Parameters",
						MarkdownDescription: "DEPRECATED: use Parameters instead, e.g., Parameters['compression_mode'] = 'force'The inline compression mode in Bluestore OSD to set to (options are: none, passive, aggressive, force)Do NOT set a default value for kubebuilder as this will override the Parameters",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("none", "passive", "aggressive", "force", ""),
						},
					},

					"crush_root": schema.StringAttribute{
						Description:         "The root of the crush hierarchy utilized by the pool",
						MarkdownDescription: "The root of the crush hierarchy utilized by the pool",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"device_class": schema.StringAttribute{
						Description:         "The device class the OSD should set to for use in the pool",
						MarkdownDescription: "The device class the OSD should set to for use in the pool",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_rbd_stats": schema.BoolAttribute{
						Description:         "EnableRBDStats is used to enable gathering of statistics for all RBD images in the pool",
						MarkdownDescription: "EnableRBDStats is used to enable gathering of statistics for all RBD images in the pool",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"erasure_coded": schema.SingleNestedAttribute{
						Description:         "The erasure code settings",
						MarkdownDescription: "The erasure code settings",
						Attributes: map[string]schema.Attribute{
							"algorithm": schema.StringAttribute{
								Description:         "The algorithm for erasure coding",
								MarkdownDescription: "The algorithm for erasure coding",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"coding_chunks": schema.Int64Attribute{
								Description:         "Number of coding chunks per object in an erasure coded storage pool (required for erasure-coded pool type).This is the number of OSDs that can be lost simultaneously before data cannot be recovered.",
								MarkdownDescription: "Number of coding chunks per object in an erasure coded storage pool (required for erasure-coded pool type).This is the number of OSDs that can be lost simultaneously before data cannot be recovered.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
								},
							},

							"data_chunks": schema.Int64Attribute{
								Description:         "Number of data chunks per object in an erasure coded storage pool (required for erasure-coded pool type).The number of chunks required to recover an object when any single OSD is lost is the sameas dataChunks so be aware that the larger the number of data chunks, the higher the cost of recovery.",
								MarkdownDescription: "Number of data chunks per object in an erasure coded storage pool (required for erasure-coded pool type).The number of chunks required to recover an object when any single OSD is lost is the sameas dataChunks so be aware that the larger the number of data chunks, the higher the cost of recovery.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"failure_domain": schema.StringAttribute{
						Description:         "The failure domain: osd/host/(region or zone if available) - technically also any type in the crush map",
						MarkdownDescription: "The failure domain: osd/host/(region or zone if available) - technically also any type in the crush map",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"mirroring": schema.SingleNestedAttribute{
						Description:         "The mirroring settings",
						MarkdownDescription: "The mirroring settings",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "Enabled whether this pool is mirrored or not",
								MarkdownDescription: "Enabled whether this pool is mirrored or not",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"mode": schema.StringAttribute{
								Description:         "Mode is the mirroring mode: either pool or image",
								MarkdownDescription: "Mode is the mirroring mode: either pool or image",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"peers": schema.SingleNestedAttribute{
								Description:         "Peers represents the peers spec",
								MarkdownDescription: "Peers represents the peers spec",
								Attributes: map[string]schema.Attribute{
									"secret_names": schema.ListAttribute{
										Description:         "SecretNames represents the Kubernetes Secret names to add rbd-mirror or cephfs-mirror peers",
										MarkdownDescription: "SecretNames represents the Kubernetes Secret names to add rbd-mirror or cephfs-mirror peers",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"snapshot_schedules": schema.ListNestedAttribute{
								Description:         "SnapshotSchedules is the scheduling of snapshot for mirrored images/pools",
								MarkdownDescription: "SnapshotSchedules is the scheduling of snapshot for mirrored images/pools",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"interval": schema.StringAttribute{
											Description:         "Interval represent the periodicity of the snapshot.",
											MarkdownDescription: "Interval represent the periodicity of the snapshot.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"path": schema.StringAttribute{
											Description:         "Path is the path to snapshot, only valid for CephFS",
											MarkdownDescription: "Path is the path to snapshot, only valid for CephFS",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"start_time": schema.StringAttribute{
											Description:         "StartTime indicates when to start the snapshot",
											MarkdownDescription: "StartTime indicates when to start the snapshot",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"name": schema.StringAttribute{
						Description:         "The desired name of the pool if different from the CephBlockPool CR name.",
						MarkdownDescription: "The desired name of the pool if different from the CephBlockPool CR name.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf(".rgw.root", ".nfs", ".mgr"),
						},
					},

					"parameters": schema.MapAttribute{
						Description:         "Parameters is a list of properties to enable on a given pool",
						MarkdownDescription: "Parameters is a list of properties to enable on a given pool",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"quotas": schema.SingleNestedAttribute{
						Description:         "The quota settings",
						MarkdownDescription: "The quota settings",
						Attributes: map[string]schema.Attribute{
							"max_bytes": schema.Int64Attribute{
								Description:         "MaxBytes represents the quota in bytesDeprecated in favor of MaxSize",
								MarkdownDescription: "MaxBytes represents the quota in bytesDeprecated in favor of MaxSize",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_objects": schema.Int64Attribute{
								Description:         "MaxObjects represents the quota in objects",
								MarkdownDescription: "MaxObjects represents the quota in objects",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_size": schema.StringAttribute{
								Description:         "MaxSize represents the quota in bytes as a string",
								MarkdownDescription: "MaxSize represents the quota in bytes as a string",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9]+[\.]?[0-9]*([KMGTPE]i|[kMGTPE])?$`), ""),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"replicated": schema.SingleNestedAttribute{
						Description:         "The replication settings",
						MarkdownDescription: "The replication settings",
						Attributes: map[string]schema.Attribute{
							"hybrid_storage": schema.SingleNestedAttribute{
								Description:         "HybridStorage represents hybrid storage tier settings",
								MarkdownDescription: "HybridStorage represents hybrid storage tier settings",
								Attributes: map[string]schema.Attribute{
									"primary_device_class": schema.StringAttribute{
										Description:         "PrimaryDeviceClass represents high performance tier (for example SSD or NVME) for Primary OSD",
										MarkdownDescription: "PrimaryDeviceClass represents high performance tier (for example SSD or NVME) for Primary OSD",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(1),
										},
									},

									"secondary_device_class": schema.StringAttribute{
										Description:         "SecondaryDeviceClass represents low performance tier (for example HDDs) for remaining OSDs",
										MarkdownDescription: "SecondaryDeviceClass represents low performance tier (for example HDDs) for remaining OSDs",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(1),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"replicas_per_failure_domain": schema.Int64Attribute{
								Description:         "ReplicasPerFailureDomain the number of replica in the specified failure domain",
								MarkdownDescription: "ReplicasPerFailureDomain the number of replica in the specified failure domain",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
								},
							},

							"require_safe_replica_size": schema.BoolAttribute{
								Description:         "RequireSafeReplicaSize if false allows you to set replica 1",
								MarkdownDescription: "RequireSafeReplicaSize if false allows you to set replica 1",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"size": schema.Int64Attribute{
								Description:         "Size - Number of copies per object in a replicated storage pool, including the object itself (required for replicated pool type)",
								MarkdownDescription: "Size - Number of copies per object in a replicated storage pool, including the object itself (required for replicated pool type)",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
								},
							},

							"sub_failure_domain": schema.StringAttribute{
								Description:         "SubFailureDomain the name of the sub-failure domain",
								MarkdownDescription: "SubFailureDomain the name of the sub-failure domain",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"target_size_ratio": schema.Float64Attribute{
								Description:         "TargetSizeRatio gives a hint (%) to Ceph in terms of expected consumption of the total cluster capacity",
								MarkdownDescription: "TargetSizeRatio gives a hint (%) to Ceph in terms of expected consumption of the total cluster capacity",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"status_check": schema.SingleNestedAttribute{
						Description:         "The mirroring statusCheck",
						MarkdownDescription: "The mirroring statusCheck",
						Attributes: map[string]schema.Attribute{
							"mirror": schema.SingleNestedAttribute{
								Description:         "HealthCheckSpec represents the health check of an object store bucket",
								MarkdownDescription: "HealthCheckSpec represents the health check of an object store bucket",
								Attributes: map[string]schema.Attribute{
									"disabled": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"interval": schema.StringAttribute{
										Description:         "Interval is the internal in second or minute for the health check to run like 60s for 60 seconds",
										MarkdownDescription: "Interval is the internal in second or minute for the health check to run like 60s for 60 seconds",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"timeout": schema.StringAttribute{
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

func (r *CephRookIoCephBlockPoolV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_ceph_rook_io_ceph_block_pool_v1_manifest")

	var model CephRookIoCephBlockPoolV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("ceph.rook.io/v1")
	model.Kind = pointer.String("CephBlockPool")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
