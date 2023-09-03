/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package ceph_rook_io_v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	k8sTypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"regexp"
	"strings"
)

var (
	_ resource.Resource                = &CephRookIoCephBlockPoolV1Resource{}
	_ resource.ResourceWithConfigure   = &CephRookIoCephBlockPoolV1Resource{}
	_ resource.ResourceWithImportState = &CephRookIoCephBlockPoolV1Resource{}
)

func NewCephRookIoCephBlockPoolV1Resource() resource.Resource {
	return &CephRookIoCephBlockPoolV1Resource{}
}

type CephRookIoCephBlockPoolV1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type CephRookIoCephBlockPoolV1ResourceData struct {
	ID             types.String `tfsdk:"id" json:"-"`
	ForceConflicts types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager   types.String `tfsdk:"field_manager" json:"-"`
	WaitFor        types.List   `tfsdk:"wait_for" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
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

func (r *CephRookIoCephBlockPoolV1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_ceph_rook_io_ceph_block_pool_v1"
}

func (r *CephRookIoCephBlockPoolV1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CephBlockPool represents a Ceph Storage Pool",
		MarkdownDescription: "CephBlockPool represents a Ceph Storage Pool",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"force_conflicts": schema.BoolAttribute{
				Description:         "If 'true', server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "If `true`, server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"field_manager": schema.BoolAttribute{
				Description:         "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"wait_for": schema.ListNestedAttribute{
				Description:         "Wait for specific conditions after create/update of resources.",
				MarkdownDescription: "Wait for specific conditions after create/update of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"jsonpath": schema.StringAttribute{
							Description:         "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							MarkdownDescription: "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							Required:            true,
							Optional:            false,
							Computed:            false,
						},
						"value": schema.StringAttribute{
							Description:         "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							MarkdownDescription: "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							Required:            false,
							Optional:            true,
							Computed:            true,
						},
						"timeout": schema.StringAttribute{
							Description:         "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							MarkdownDescription: "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("30s"),
						},
					},
				},
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},

					"labels": schema.MapAttribute{
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            true,
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
						Computed:            true,
						Validators: []validator.Map{
							validators.AnnotationValidator(),
						},
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "NamedBlockPoolSpec allows a block pool to be created with a non-default name. This is more specific than the NamedPoolSpec so we get schema validation on the allowed pool names that can be specified.",
				MarkdownDescription: "NamedBlockPoolSpec allows a block pool to be created with a non-default name. This is more specific than the NamedPoolSpec so we get schema validation on the allowed pool names that can be specified.",
				Attributes: map[string]schema.Attribute{
					"compression_mode": schema.StringAttribute{
						Description:         "DEPRECATED: use Parameters instead, e.g., Parameters['compression_mode'] = 'force' The inline compression mode in Bluestore OSD to set to (options are: none, passive, aggressive, force) Do NOT set a default value for kubebuilder as this will override the Parameters",
						MarkdownDescription: "DEPRECATED: use Parameters instead, e.g., Parameters['compression_mode'] = 'force' The inline compression mode in Bluestore OSD to set to (options are: none, passive, aggressive, force) Do NOT set a default value for kubebuilder as this will override the Parameters",
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
								Description:         "Number of coding chunks per object in an erasure coded storage pool (required for erasure-coded pool type). This is the number of OSDs that can be lost simultaneously before data cannot be recovered.",
								MarkdownDescription: "Number of coding chunks per object in an erasure coded storage pool (required for erasure-coded pool type). This is the number of OSDs that can be lost simultaneously before data cannot be recovered.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
								},
							},

							"data_chunks": schema.Int64Attribute{
								Description:         "Number of data chunks per object in an erasure coded storage pool (required for erasure-coded pool type). The number of chunks required to recover an object when any single OSD is lost is the same as dataChunks so be aware that the larger the number of data chunks, the higher the cost of recovery.",
								MarkdownDescription: "Number of data chunks per object in an erasure coded storage pool (required for erasure-coded pool type). The number of chunks required to recover an object when any single OSD is lost is the same as dataChunks so be aware that the larger the number of data chunks, the higher the cost of recovery.",
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
							stringvalidator.OneOf("device_health_metrics", ".nfs", ".mgr"),
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
								Description:         "MaxBytes represents the quota in bytes Deprecated in favor of MaxSize",
								MarkdownDescription: "MaxBytes represents the quota in bytes Deprecated in favor of MaxSize",
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

func (r *CephRookIoCephBlockPoolV1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *dynamic.DynamicClient, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *CephRookIoCephBlockPoolV1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_ceph_rook_io_ceph_block_pool_v1")

	var model CephRookIoCephBlockPoolV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("ceph.rook.io/v1")
	model.Kind = pointer.String("CephBlockPool")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.Resource(k8sSchema.GroupVersionResource{Group: "ceph.rook.io", Version: "v1", Resource: "CephBlockPool"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while creating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse CephRookIoCephBlockPoolV1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *CephRookIoCephBlockPoolV1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_ceph_rook_io_ceph_block_pool_v1")

	var data CephRookIoCephBlockPoolV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "ceph.rook.io", Version: "v1", Resource: "CephBlockPool"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse CephRookIoCephBlockPoolV1ResourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *CephRookIoCephBlockPoolV1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_ceph_rook_io_ceph_block_pool_v1")

	var model CephRookIoCephBlockPoolV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("ceph.rook.io/v1")
	model.Kind = pointer.String("CephBlockPool")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.Resource(k8sSchema.GroupVersionResource{Group: "ceph.rook.io", Version: "v1", Resource: "CephBlockPool"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while updating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse CephRookIoCephBlockPoolV1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *CephRookIoCephBlockPoolV1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_ceph_rook_io_ceph_block_pool_v1")

	var data CephRookIoCephBlockPoolV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "ceph.rook.io", Version: "v1", Resource: "CephBlockPool"}).
		Namespace(data.Metadata.Namespace).
		Delete(ctx, data.Metadata.Name, meta.DeleteOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to DELETE resource",
			"An unexpected error occurred while deleting the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"DELETE Error: "+err.Error(),
		)
		return
	}
}

func (r *CephRookIoCephBlockPoolV1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	idParts := strings.Split(request.ID, "/")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'namespace/name' Got: '%q'", request.ID),
		)
		return
	}

	namespace := idParts[0]
	name := idParts[1]
	tflog.Trace(ctx, "parsed import ID", map[string]interface{}{
		"namespace": namespace,
		"name":      name,
	})
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("namespace"), namespace)...)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("name"), name)...)
}
