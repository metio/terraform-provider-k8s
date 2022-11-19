/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"gopkg.in/yaml.v3"
	"time"
)

type CephRookIoCephBlockPoolV1Resource struct{}

var (
	_ resource.Resource = (*CephRookIoCephBlockPoolV1Resource)(nil)
)

type CephRookIoCephBlockPoolV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type CephRookIoCephBlockPoolV1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		CompressionMode *string `tfsdk:"compression_mode" yaml:"compressionMode,omitempty"`

		CrushRoot *string `tfsdk:"crush_root" yaml:"crushRoot,omitempty"`

		DeviceClass *string `tfsdk:"device_class" yaml:"deviceClass,omitempty"`

		EnableRBDStats *bool `tfsdk:"enable_rbd_stats" yaml:"enableRBDStats,omitempty"`

		ErasureCoded *struct {
			Algorithm *string `tfsdk:"algorithm" yaml:"algorithm,omitempty"`

			CodingChunks *int64 `tfsdk:"coding_chunks" yaml:"codingChunks,omitempty"`

			DataChunks *int64 `tfsdk:"data_chunks" yaml:"dataChunks,omitempty"`
		} `tfsdk:"erasure_coded" yaml:"erasureCoded,omitempty"`

		FailureDomain *string `tfsdk:"failure_domain" yaml:"failureDomain,omitempty"`

		Mirroring *struct {
			Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

			Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`

			Peers *struct {
				SecretNames *[]string `tfsdk:"secret_names" yaml:"secretNames,omitempty"`
			} `tfsdk:"peers" yaml:"peers,omitempty"`

			SnapshotSchedules *[]struct {
				Interval *string `tfsdk:"interval" yaml:"interval,omitempty"`

				Path *string `tfsdk:"path" yaml:"path,omitempty"`

				StartTime *string `tfsdk:"start_time" yaml:"startTime,omitempty"`
			} `tfsdk:"snapshot_schedules" yaml:"snapshotSchedules,omitempty"`
		} `tfsdk:"mirroring" yaml:"mirroring,omitempty"`

		Name *string `tfsdk:"name" yaml:"name,omitempty"`

		Parameters utilities.Dynamic `tfsdk:"parameters" yaml:"parameters,omitempty"`

		Quotas *struct {
			MaxBytes *int64 `tfsdk:"max_bytes" yaml:"maxBytes,omitempty"`

			MaxObjects *int64 `tfsdk:"max_objects" yaml:"maxObjects,omitempty"`

			MaxSize *string `tfsdk:"max_size" yaml:"maxSize,omitempty"`
		} `tfsdk:"quotas" yaml:"quotas,omitempty"`

		Replicated *struct {
			HybridStorage *struct {
				PrimaryDeviceClass *string `tfsdk:"primary_device_class" yaml:"primaryDeviceClass,omitempty"`

				SecondaryDeviceClass *string `tfsdk:"secondary_device_class" yaml:"secondaryDeviceClass,omitempty"`
			} `tfsdk:"hybrid_storage" yaml:"hybridStorage,omitempty"`

			ReplicasPerFailureDomain *int64 `tfsdk:"replicas_per_failure_domain" yaml:"replicasPerFailureDomain,omitempty"`

			RequireSafeReplicaSize *bool `tfsdk:"require_safe_replica_size" yaml:"requireSafeReplicaSize,omitempty"`

			Size *int64 `tfsdk:"size" yaml:"size,omitempty"`

			SubFailureDomain *string `tfsdk:"sub_failure_domain" yaml:"subFailureDomain,omitempty"`

			TargetSizeRatio utilities.DynamicNumber `tfsdk:"target_size_ratio" yaml:"targetSizeRatio,omitempty"`
		} `tfsdk:"replicated" yaml:"replicated,omitempty"`

		StatusCheck *struct {
			Mirror *struct {
				Disabled *bool `tfsdk:"disabled" yaml:"disabled,omitempty"`

				Interval *string `tfsdk:"interval" yaml:"interval,omitempty"`

				Timeout *string `tfsdk:"timeout" yaml:"timeout,omitempty"`
			} `tfsdk:"mirror" yaml:"mirror,omitempty"`
		} `tfsdk:"status_check" yaml:"statusCheck,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewCephRookIoCephBlockPoolV1Resource() resource.Resource {
	return &CephRookIoCephBlockPoolV1Resource{}
}

func (r *CephRookIoCephBlockPoolV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ceph_rook_io_ceph_block_pool_v1"
}

func (r *CephRookIoCephBlockPoolV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "CephBlockPool represents a Ceph Storage Pool",
		MarkdownDescription: "CephBlockPool represents a Ceph Storage Pool",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description:         "The timestamp of the last change to this resource.",
				MarkdownDescription: "The timestamp of the last change to this resource.",
				Type:                types.Int64Type,
				Computed:            true,
				Optional:            false,
			},

			"yaml": {
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"metadata": {
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Type:                types.StringType,
						Required:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
					},

					"namespace": {
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Type:                types.StringType,
						Optional:            true,
					},

					"labels": {
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.LabelValidator(),
						},
					},
					"annotations": {
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.AnnotationValidator(),
						},
					},
				}),
			},

			"api_version": {
				Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"kind": {
				Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"spec": {
				Description:         "NamedBlockPoolSpec allows a block pool to be created with a non-default name. This is more specific than the NamedPoolSpec so we get schema validation on the allowed pool names that can be specified.",
				MarkdownDescription: "NamedBlockPoolSpec allows a block pool to be created with a non-default name. This is more specific than the NamedPoolSpec so we get schema validation on the allowed pool names that can be specified.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"compression_mode": {
						Description:         "DEPRECATED: use Parameters instead, e.g., Parameters['compression_mode'] = 'force' The inline compression mode in Bluestore OSD to set to (options are: none, passive, aggressive, force) Do NOT set a default value for kubebuilder as this will override the Parameters",
						MarkdownDescription: "DEPRECATED: use Parameters instead, e.g., Parameters['compression_mode'] = 'force' The inline compression mode in Bluestore OSD to set to (options are: none, passive, aggressive, force) Do NOT set a default value for kubebuilder as this will override the Parameters",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("none", "passive", "aggressive", "force", ""),
						},
					},

					"crush_root": {
						Description:         "The root of the crush hierarchy utilized by the pool",
						MarkdownDescription: "The root of the crush hierarchy utilized by the pool",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"device_class": {
						Description:         "The device class the OSD should set to for use in the pool",
						MarkdownDescription: "The device class the OSD should set to for use in the pool",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable_rbd_stats": {
						Description:         "EnableRBDStats is used to enable gathering of statistics for all RBD images in the pool",
						MarkdownDescription: "EnableRBDStats is used to enable gathering of statistics for all RBD images in the pool",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"erasure_coded": {
						Description:         "The erasure code settings",
						MarkdownDescription: "The erasure code settings",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"algorithm": {
								Description:         "The algorithm for erasure coding",
								MarkdownDescription: "The algorithm for erasure coding",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"coding_chunks": {
								Description:         "Number of coding chunks per object in an erasure coded storage pool (required for erasure-coded pool type). This is the number of OSDs that can be lost simultaneously before data cannot be recovered.",
								MarkdownDescription: "Number of coding chunks per object in an erasure coded storage pool (required for erasure-coded pool type). This is the number of OSDs that can be lost simultaneously before data cannot be recovered.",

								Type: types.Int64Type,

								Required: true,
								Optional: false,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),
								},
							},

							"data_chunks": {
								Description:         "Number of data chunks per object in an erasure coded storage pool (required for erasure-coded pool type). The number of chunks required to recover an object when any single OSD is lost is the same as dataChunks so be aware that the larger the number of data chunks, the higher the cost of recovery.",
								MarkdownDescription: "Number of data chunks per object in an erasure coded storage pool (required for erasure-coded pool type). The number of chunks required to recover an object when any single OSD is lost is the same as dataChunks so be aware that the larger the number of data chunks, the higher the cost of recovery.",

								Type: types.Int64Type,

								Required: true,
								Optional: false,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),
								},
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"failure_domain": {
						Description:         "The failure domain: osd/host/(region or zone if available) - technically also any type in the crush map",
						MarkdownDescription: "The failure domain: osd/host/(region or zone if available) - technically also any type in the crush map",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"mirroring": {
						Description:         "The mirroring settings",
						MarkdownDescription: "The mirroring settings",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"enabled": {
								Description:         "Enabled whether this pool is mirrored or not",
								MarkdownDescription: "Enabled whether this pool is mirrored or not",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"mode": {
								Description:         "Mode is the mirroring mode: either pool or image",
								MarkdownDescription: "Mode is the mirroring mode: either pool or image",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"peers": {
								Description:         "Peers represents the peers spec",
								MarkdownDescription: "Peers represents the peers spec",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"secret_names": {
										Description:         "SecretNames represents the Kubernetes Secret names to add rbd-mirror or cephfs-mirror peers",
										MarkdownDescription: "SecretNames represents the Kubernetes Secret names to add rbd-mirror or cephfs-mirror peers",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"snapshot_schedules": {
								Description:         "SnapshotSchedules is the scheduling of snapshot for mirrored images/pools",
								MarkdownDescription: "SnapshotSchedules is the scheduling of snapshot for mirrored images/pools",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"interval": {
										Description:         "Interval represent the periodicity of the snapshot.",
										MarkdownDescription: "Interval represent the periodicity of the snapshot.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"path": {
										Description:         "Path is the path to snapshot, only valid for CephFS",
										MarkdownDescription: "Path is the path to snapshot, only valid for CephFS",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"start_time": {
										Description:         "StartTime indicates when to start the snapshot",
										MarkdownDescription: "StartTime indicates when to start the snapshot",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"name": {
						Description:         "The desired name of the pool if different from the CephBlockPool CR name.",
						MarkdownDescription: "The desired name of the pool if different from the CephBlockPool CR name.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("device_health_metrics", ".nfs", ".mgr"),
						},
					},

					"parameters": {
						Description:         "Parameters is a list of properties to enable on a given pool",
						MarkdownDescription: "Parameters is a list of properties to enable on a given pool",

						Type: utilities.DynamicType{},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"quotas": {
						Description:         "The quota settings",
						MarkdownDescription: "The quota settings",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"max_bytes": {
								Description:         "MaxBytes represents the quota in bytes Deprecated in favor of MaxSize",
								MarkdownDescription: "MaxBytes represents the quota in bytes Deprecated in favor of MaxSize",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"max_objects": {
								Description:         "MaxObjects represents the quota in objects",
								MarkdownDescription: "MaxObjects represents the quota in objects",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"max_size": {
								Description:         "MaxSize represents the quota in bytes as a string",
								MarkdownDescription: "MaxSize represents the quota in bytes as a string",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9]+[\.]?[0-9]*([KMGTPE]i|[kMGTPE])?$`), ""),
								},
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"replicated": {
						Description:         "The replication settings",
						MarkdownDescription: "The replication settings",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"hybrid_storage": {
								Description:         "HybridStorage represents hybrid storage tier settings",
								MarkdownDescription: "HybridStorage represents hybrid storage tier settings",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"primary_device_class": {
										Description:         "PrimaryDeviceClass represents high performance tier (for example SSD or NVME) for Primary OSD",
										MarkdownDescription: "PrimaryDeviceClass represents high performance tier (for example SSD or NVME) for Primary OSD",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.LengthAtLeast(1),
										},
									},

									"secondary_device_class": {
										Description:         "SecondaryDeviceClass represents low performance tier (for example HDDs) for remaining OSDs",
										MarkdownDescription: "SecondaryDeviceClass represents low performance tier (for example HDDs) for remaining OSDs",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.LengthAtLeast(1),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"replicas_per_failure_domain": {
								Description:         "ReplicasPerFailureDomain the number of replica in the specified failure domain",
								MarkdownDescription: "ReplicasPerFailureDomain the number of replica in the specified failure domain",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(1),
								},
							},

							"require_safe_replica_size": {
								Description:         "RequireSafeReplicaSize if false allows you to set replica 1",
								MarkdownDescription: "RequireSafeReplicaSize if false allows you to set replica 1",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"size": {
								Description:         "Size - Number of copies per object in a replicated storage pool, including the object itself (required for replicated pool type)",
								MarkdownDescription: "Size - Number of copies per object in a replicated storage pool, including the object itself (required for replicated pool type)",

								Type: types.Int64Type,

								Required: true,
								Optional: false,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),
								},
							},

							"sub_failure_domain": {
								Description:         "SubFailureDomain the name of the sub-failure domain",
								MarkdownDescription: "SubFailureDomain the name of the sub-failure domain",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"target_size_ratio": {
								Description:         "TargetSizeRatio gives a hint (%) to Ceph in terms of expected consumption of the total cluster capacity",
								MarkdownDescription: "TargetSizeRatio gives a hint (%) to Ceph in terms of expected consumption of the total cluster capacity",

								Type: utilities.DynamicNumberType{},

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"status_check": {
						Description:         "The mirroring statusCheck",
						MarkdownDescription: "The mirroring statusCheck",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"mirror": {
								Description:         "HealthCheckSpec represents the health check of an object store bucket",
								MarkdownDescription: "HealthCheckSpec represents the health check of an object store bucket",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"disabled": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"interval": {
										Description:         "Interval is the internal in second or minute for the health check to run like 60s for 60 seconds",
										MarkdownDescription: "Interval is the internal in second or minute for the health check to run like 60s for 60 seconds",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"timeout": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}, nil
}

func (r *CephRookIoCephBlockPoolV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_ceph_rook_io_ceph_block_pool_v1")

	var state CephRookIoCephBlockPoolV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CephRookIoCephBlockPoolV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("ceph.rook.io/v1")
	goModel.Kind = utilities.Ptr("CephBlockPool")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *CephRookIoCephBlockPoolV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_ceph_rook_io_ceph_block_pool_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *CephRookIoCephBlockPoolV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_ceph_rook_io_ceph_block_pool_v1")

	var state CephRookIoCephBlockPoolV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CephRookIoCephBlockPoolV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("ceph.rook.io/v1")
	goModel.Kind = utilities.Ptr("CephBlockPool")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *CephRookIoCephBlockPoolV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_ceph_rook_io_ceph_block_pool_v1")
	// NO-OP: Terraform removes the state automatically for us
}
