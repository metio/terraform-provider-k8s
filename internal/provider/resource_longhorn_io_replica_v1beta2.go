/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

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

type LonghornIoReplicaV1Beta2Resource struct{}

var (
	_ resource.Resource = (*LonghornIoReplicaV1Beta2Resource)(nil)
)

type LonghornIoReplicaV1Beta2TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type LonghornIoReplicaV1Beta2GoModel struct {
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
		Active *bool `tfsdk:"active" yaml:"active,omitempty"`

		BackingImage *string `tfsdk:"backing_image" yaml:"backingImage,omitempty"`

		BaseImage *string `tfsdk:"base_image" yaml:"baseImage,omitempty"`

		DataDirectoryName *string `tfsdk:"data_directory_name" yaml:"dataDirectoryName,omitempty"`

		DataPath *string `tfsdk:"data_path" yaml:"dataPath,omitempty"`

		DesireState *string `tfsdk:"desire_state" yaml:"desireState,omitempty"`

		DiskID *string `tfsdk:"disk_id" yaml:"diskID,omitempty"`

		DiskPath *string `tfsdk:"disk_path" yaml:"diskPath,omitempty"`

		EngineImage *string `tfsdk:"engine_image" yaml:"engineImage,omitempty"`

		EngineName *string `tfsdk:"engine_name" yaml:"engineName,omitempty"`

		FailedAt *string `tfsdk:"failed_at" yaml:"failedAt,omitempty"`

		HardNodeAffinity *string `tfsdk:"hard_node_affinity" yaml:"hardNodeAffinity,omitempty"`

		HealthyAt *string `tfsdk:"healthy_at" yaml:"healthyAt,omitempty"`

		LogRequested *bool `tfsdk:"log_requested" yaml:"logRequested,omitempty"`

		NodeID *string `tfsdk:"node_id" yaml:"nodeID,omitempty"`

		RebuildRetryCount *int64 `tfsdk:"rebuild_retry_count" yaml:"rebuildRetryCount,omitempty"`

		RevisionCounterDisabled *bool `tfsdk:"revision_counter_disabled" yaml:"revisionCounterDisabled,omitempty"`

		SalvageRequested *bool `tfsdk:"salvage_requested" yaml:"salvageRequested,omitempty"`

		VolumeName *string `tfsdk:"volume_name" yaml:"volumeName,omitempty"`

		VolumeSize *string `tfsdk:"volume_size" yaml:"volumeSize,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewLonghornIoReplicaV1Beta2Resource() resource.Resource {
	return &LonghornIoReplicaV1Beta2Resource{}
}

func (r *LonghornIoReplicaV1Beta2Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_longhorn_io_replica_v1beta2"
}

func (r *LonghornIoReplicaV1Beta2Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Replica is where Longhorn stores replica object.",
		MarkdownDescription: "Replica is where Longhorn stores replica object.",
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
				Description:         "ReplicaSpec defines the desired state of the Longhorn replica",
				MarkdownDescription: "ReplicaSpec defines the desired state of the Longhorn replica",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"active": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"backing_image": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"base_image": {
						Description:         "Deprecated. Rename to BackingImage",
						MarkdownDescription: "Deprecated. Rename to BackingImage",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"data_directory_name": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"data_path": {
						Description:         "Deprecated",
						MarkdownDescription: "Deprecated",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"desire_state": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"disk_id": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"disk_path": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"engine_image": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"engine_name": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"failed_at": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"hard_node_affinity": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"healthy_at": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"log_requested": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"node_id": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"rebuild_retry_count": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"revision_counter_disabled": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"salvage_requested": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"volume_name": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"volume_size": {
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
		},
	}, nil
}

func (r *LonghornIoReplicaV1Beta2Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_longhorn_io_replica_v1beta2")

	var state LonghornIoReplicaV1Beta2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel LonghornIoReplicaV1Beta2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("longhorn.io/v1beta2")
	goModel.Kind = utilities.Ptr("Replica")

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

func (r *LonghornIoReplicaV1Beta2Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_longhorn_io_replica_v1beta2")
	// NO-OP: All data is already in Terraform state
}

func (r *LonghornIoReplicaV1Beta2Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_longhorn_io_replica_v1beta2")

	var state LonghornIoReplicaV1Beta2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel LonghornIoReplicaV1Beta2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("longhorn.io/v1beta2")
	goModel.Kind = utilities.Ptr("Replica")

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

func (r *LonghornIoReplicaV1Beta2Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_longhorn_io_replica_v1beta2")
	// NO-OP: Terraform removes the state automatically for us
}
